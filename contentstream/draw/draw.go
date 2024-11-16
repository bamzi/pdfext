// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw

import (
	_f "fmt"
	_g "math"

	_gb "github.com/bamzi/pdfext/contentstream"
	_gc "github.com/bamzi/pdfext/core"
	_c "github.com/bamzi/pdfext/internal/transform"
	_ff "github.com/bamzi/pdfext/model"
)

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_fg Path) RemovePoint(number int) Path {
	if number < 1 || number > len(_fg.Points) {
		return _fg
	}
	_df := number - 1
	_fg.Points = append(_fg.Points[:_df], _fg.Points[_df+1:]...)
	return _fg
}

// AppendPoint adds the specified point to the path.
func (_bfe Path) AppendPoint(point Point) Path { _bfe.Points = append(_bfe.Points, point); return _bfe }

// GetBoundingBox returns the bounding box of the path.
func (_ed Path) GetBoundingBox() BoundingBox {
	_ac := BoundingBox{}
	_fea := 0.0
	_cdc := 0.0
	_fca := 0.0
	_cg := 0.0
	for _de, _cf := range _ed.Points {
		if _de == 0 {
			_fea = _cf.X
			_cdc = _cf.X
			_fca = _cf.Y
			_cg = _cf.Y
			continue
		}
		if _cf.X < _fea {
			_fea = _cf.X
		}
		if _cf.X > _cdc {
			_cdc = _cf.X
		}
		if _cf.Y < _fca {
			_fca = _cf.Y
		}
		if _cf.Y > _cg {
			_cg = _cf.Y
		}
	}
	_ac.X = _fea
	_ac.Y = _fca
	_ac.Width = _cdc - _fea
	_ac.Height = _cg - _fca
	return _ac
}

// Length returns the number of points in the path.
func (_ga Path) Length() int { return len(_ga.Points) }

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineColor _ff.PdfColor
	Opacity   float64
	LineWidth float64
	LineStyle LineStyle
	DashArray []int64
	DashPhase int64
}

// NewPath returns a new empty path.
func NewPath() Path { return Path{} }

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int

// AddVector adds vector to a point.
func (_ead Point) AddVector(v Vector) Point { _ead.X += v.Dx; _ead.Y += v.Dy; return _ead }

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{ Points []Point }

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_dg Point) Add(dx, dy float64) Point { _dg.X += dx; _dg.Y += dy; return _dg }

// Copy returns a clone of the path.
func (_a Path) Copy() Path { _fc := Path{}; _fc.Points = append(_fc.Points, _a.Points...); return _fc }

// Offset shifts the path with the specified offsets.
func (_ec Path) Offset(offX, offY float64) Path {
	for _fcd, _cd := range _ec.Points {
		_ec.Points[_fcd] = _cd.Add(offX, offY)
	}
	return _ec
}

// AppendCurve appends the specified Bezier curve to the path.
func (_be CubicBezierPath) AppendCurve(curve CubicBezierCurve) CubicBezierPath {
	_be.Curves = append(_be.Curves, curve)
	return _be
}

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_caf BoundingBox) ToPdfRectangle() *_ff.PdfRectangle {
	return &_ff.PdfRectangle{Llx: _caf.X, Lly: _caf.Y, Urx: _caf.X + _caf.Width, Ury: _caf.Y + _caf.Height}
}

// Draw draws the composite curve polygon and marked the content using the specified marked content id.
// A graphics state name can be specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty ("").
//
// If mcid is nil, no marked content is added.
//
// Returns the content stream as a byte array and the bounding box of the polygon.
func (_cb CurvePolygon) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	_acd := _gb.NewContentCreator()
	if mcid != nil {
		_acd.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_acd.Add_q()
	_cb.FillEnabled = _cb.FillEnabled && _cb.FillColor != nil
	if _cb.FillEnabled {
		_acd.SetNonStrokingColor(_cb.FillColor)
	}
	_cb.BorderEnabled = _cb.BorderEnabled && _cb.BorderColor != nil
	if _cb.BorderEnabled {
		_acd.SetStrokingColor(_cb.BorderColor)
		_acd.Add_w(_cb.BorderWidth)
	}
	if len(gsName) > 1 {
		_acd.Add_gs(_gc.PdfObjectName(gsName))
	}
	_ede := NewCubicBezierPath()
	for _, _gcg := range _cb.Rings {
		for _ae, _gdf := range _gcg {
			if _ae == 0 {
				_acd.Add_m(_gdf.P0.X, _gdf.P0.Y)
			} else {
				_acd.Add_l(_gdf.P0.X, _gdf.P0.Y)
			}
			_acd.Add_c(_gdf.P1.X, _gdf.P1.Y, _gdf.P2.X, _gdf.P2.Y, _gdf.P3.X, _gdf.P3.Y)
			_ede = _ede.AppendCurve(_gdf)
		}
		_acd.Add_h()
	}
	if _cb.FillEnabled && _cb.BorderEnabled {
		_acd.Add_B()
	} else if _cb.FillEnabled {
		_acd.Add_f()
	} else if _cb.BorderEnabled {
		_acd.Add_S()
	}
	_acd.Add_Q()
	if mcid != nil {
		_acd.Add_EMC()
	}
	return _acd.Bytes(), _ede.GetBoundingBox().ToPdfRectangle(), nil
}

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve(x0, y0, x1, y1, x2, y2, x3, y3 float64) CubicBezierCurve {
	_d := CubicBezierCurve{}
	_d.P0 = NewPoint(x0, y0)
	_d.P1 = NewPoint(x1, y1)
	_d.P2 = NewPoint(x2, y2)
	_d.P3 = NewPoint(x3, y3)
	return _d
}

// Draw draws the rectangle and marked the content using the specified marked content id.
// A graphics state can be specified for setting additional properties (e.g. opacity).
// Otherwise pass an empty string for the `gsName` parameter.
//
// If `mcid` is nil, no marked content is added.
//
// The method returns the content stream as a byte array and the bounding box of the shape.
func (_ddg Rectangle) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	_dgd := _gb.NewContentCreator()
	if mcid != nil {
		_dgd.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_dgd.Add_q()
	if _ddg.FillEnabled {
		_dgd.SetNonStrokingColor(_ddg.FillColor)
	}
	if _ddg.BorderEnabled {
		_dgd.SetStrokingColor(_ddg.BorderColor)
		_dgd.Add_w(_ddg.BorderWidth)
	}
	if len(gsName) > 1 {
		_dgd.Add_gs(_gc.PdfObjectName(gsName))
	}
	var (
		_gcbf, _cedc = _ddg.X, _ddg.Y
		_aeg, _cab   = _ddg.Width, _ddg.Height
		_fde         = _g.Abs(_ddg.BorderRadiusTopLeft)
		_dda         = _g.Abs(_ddg.BorderRadiusTopRight)
		_fbaf        = _g.Abs(_ddg.BorderRadiusBottomLeft)
		_caa         = _g.Abs(_ddg.BorderRadiusBottomRight)
		_bce         = 0.4477
	)
	_aea := Path{Points: []Point{{X: _gcbf + _aeg - _caa, Y: _cedc}, {X: _gcbf + _aeg, Y: _cedc + _cab - _dda}, {X: _gcbf + _fde, Y: _cedc + _cab}, {X: _gcbf, Y: _cedc + _fbaf}}}
	_fcde := [][7]float64{{_caa, _gcbf + _aeg - _caa*_bce, _cedc, _gcbf + _aeg, _cedc + _caa*_bce, _gcbf + _aeg, _cedc + _caa}, {_dda, _gcbf + _aeg, _cedc + _cab - _dda*_bce, _gcbf + _aeg - _dda*_bce, _cedc + _cab, _gcbf + _aeg - _dda, _cedc + _cab}, {_fde, _gcbf + _fde*_bce, _cedc + _cab, _gcbf, _cedc + _cab - _fde*_bce, _gcbf, _cedc + _cab - _fde}, {_fbaf, _gcbf, _cedc + _fbaf*_bce, _gcbf + _fbaf*_bce, _cedc, _gcbf + _fbaf, _cedc}}
	_dgd.Add_m(_gcbf+_fbaf, _cedc)
	for _ad := 0; _ad < 4; _ad++ {
		_fcc := _aea.Points[_ad]
		_dgd.Add_l(_fcc.X, _fcc.Y)
		_bgf := _fcde[_ad]
		if _eef := _bgf[0]; _eef != 0 {
			_dgd.Add_c(_bgf[1], _bgf[2], _bgf[3], _bgf[4], _bgf[5], _bgf[6])
		}
	}
	_dgd.Add_h()
	if _ddg.FillEnabled && _ddg.BorderEnabled {
		_dgd.Add_B()
	} else if _ddg.FillEnabled {
		_dgd.Add_f()
	} else if _ddg.BorderEnabled {
		_dgd.Add_S()
	}
	_dgd.Add_Q()
	if mcid != nil {
		_dgd.Add_EMC()
	}
	return _dgd.Bytes(), _aea.GetBoundingBox().ToPdfRectangle(), nil
}

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_cdb Line) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	_agg, _ggcc := _cdb.X1, _cdb.X2
	_ggcf, _bde := _cdb.Y1, _cdb.Y2
	_dcbb := _bde - _ggcf
	_cgc := _ggcc - _agg
	_eeb := _g.Atan2(_dcbb, _cgc)
	L := _g.Sqrt(_g.Pow(_cgc, 2.0) + _g.Pow(_dcbb, 2.0))
	_ge := _cdb.LineWidth
	_cbf := _g.Pi
	_db := 1.0
	if _cgc < 0 {
		_db *= -1.0
	}
	if _dcbb < 0 {
		_db *= -1.0
	}
	VsX := _db * (-_ge / 2 * _g.Cos(_eeb+_cbf/2))
	VsY := _db * (-_ge/2*_g.Sin(_eeb+_cbf/2) + _ge*_g.Sin(_eeb+_cbf/2))
	V1X := VsX + _ge/2*_g.Cos(_eeb+_cbf/2)
	V1Y := VsY + _ge/2*_g.Sin(_eeb+_cbf/2)
	V2X := VsX + _ge/2*_g.Cos(_eeb+_cbf/2) + L*_g.Cos(_eeb)
	V2Y := VsY + _ge/2*_g.Sin(_eeb+_cbf/2) + L*_g.Sin(_eeb)
	V3X := VsX + _ge/2*_g.Cos(_eeb+_cbf/2) + L*_g.Cos(_eeb) + _ge*_g.Cos(_eeb-_cbf/2)
	V3Y := VsY + _ge/2*_g.Sin(_eeb+_cbf/2) + L*_g.Sin(_eeb) + _ge*_g.Sin(_eeb-_cbf/2)
	V4X := VsX + _ge/2*_g.Cos(_eeb-_cbf/2)
	V4Y := VsY + _ge/2*_g.Sin(_eeb-_cbf/2)
	_cda := NewPath()
	_cda = _cda.AppendPoint(NewPoint(V1X, V1Y))
	_cda = _cda.AppendPoint(NewPoint(V2X, V2Y))
	_cda = _cda.AppendPoint(NewPoint(V3X, V3Y))
	_cda = _cda.AppendPoint(NewPoint(V4X, V4Y))
	_fcb := _cdb.LineEndingStyle1
	_eae := _cdb.LineEndingStyle2
	_geb := 3 * _ge
	_eebe := 3 * _ge
	_ebf := (_eebe - _ge) / 2
	if _eae == LineEndingStyleArrow {
		_ace := _cda.GetPointNumber(2)
		_eee := NewVectorPolar(_geb, _eeb+_cbf)
		_daa := _ace.AddVector(_eee)
		_bb := NewVectorPolar(_eebe/2, _eeb+_cbf/2)
		_daf := NewVectorPolar(_geb, _eeb)
		_dac := NewVectorPolar(_ebf, _eeb+_cbf/2)
		_eefc := _daa.AddVector(_dac)
		_dgdc := _daf.Add(_bb.Flip())
		_cge := _eefc.AddVector(_dgdc)
		_gba := _bb.Scale(2).Flip().Add(_dgdc.Flip())
		_cfd := _cge.AddVector(_gba)
		_gfd := _daa.AddVector(NewVectorPolar(_ge, _eeb-_cbf/2))
		_ecd := NewPath()
		_ecd = _ecd.AppendPoint(_cda.GetPointNumber(1))
		_ecd = _ecd.AppendPoint(_daa)
		_ecd = _ecd.AppendPoint(_eefc)
		_ecd = _ecd.AppendPoint(_cge)
		_ecd = _ecd.AppendPoint(_cfd)
		_ecd = _ecd.AppendPoint(_gfd)
		_ecd = _ecd.AppendPoint(_cda.GetPointNumber(4))
		_cda = _ecd
	}
	if _fcb == LineEndingStyleArrow {
		_gbb := _cda.GetPointNumber(1)
		_ggd := _cda.GetPointNumber(_cda.Length())
		_af := NewVectorPolar(_ge/2, _eeb+_cbf+_cbf/2)
		_dcfb := _gbb.AddVector(_af)
		_cc := NewVectorPolar(_geb, _eeb).Add(NewVectorPolar(_eebe/2, _eeb+_cbf/2))
		_bda := _dcfb.AddVector(_cc)
		_aa := NewVectorPolar(_ebf, _eeb-_cbf/2)
		_fbc := _bda.AddVector(_aa)
		_ddb := NewVectorPolar(_geb, _eeb)
		_ggb := _ggd.AddVector(_ddb)
		_edb := NewVectorPolar(_ebf, _eeb+_cbf+_cbf/2)
		_bedf := _ggb.AddVector(_edb)
		_bea := _dcfb
		_gdd := NewPath()
		_gdd = _gdd.AppendPoint(_dcfb)
		_gdd = _gdd.AppendPoint(_bda)
		_gdd = _gdd.AppendPoint(_fbc)
		for _, _cbb := range _cda.Points[1 : len(_cda.Points)-1] {
			_gdd = _gdd.AppendPoint(_cbb)
		}
		_gdd = _gdd.AppendPoint(_ggb)
		_gdd = _gdd.AppendPoint(_bedf)
		_gdd = _gdd.AppendPoint(_bea)
		_cda = _gdd
	}
	_fgf := _gb.NewContentCreator()
	_fgf.Add_q().SetNonStrokingColor(_cdb.LineColor)
	if len(gsName) > 1 {
		_fgf.Add_gs(_gc.PdfObjectName(gsName))
	}
	_cda = _cda.Offset(_cdb.X1, _cdb.Y1)
	_dce := _cda.GetBoundingBox()
	DrawPathWithCreator(_cda, _fgf)
	if _cdb.LineStyle == LineStyleDashed {
		_fgf.Add_d([]int64{1, 1}, 0).Add_S().Add_f().Add_Q()
	} else {
		_fgf.Add_f().Add_Q()
	}
	return _fgf.Bytes(), _dce.ToPdfRectangle(), nil
}

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_ee Rectangle) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _ee.MarkedDraw(gsName, nil)
}

// NewPoint returns a new point with the coordinates x, y.
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath() CubicBezierPath {
	_cae := CubicBezierPath{}
	_cae.Curves = []CubicBezierCurve{}
	return _cae
}

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_eab Point) Rotate(theta float64) Point {
	_ecb := _c.NewPoint(_eab.X, _eab.Y).Rotate(theta)
	return NewPoint(_ecb.X, _ecb.Y)
}

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_da Circle) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _da.MarkedDraw(gsName, nil)
}

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_ef Path) GetPointNumber(number int) Point {
	if number < 1 || number > len(_ef.Points) {
		return Point{}
	}
	return _ef.Points[number-1]
}

// Draw draws the circle and marked the content using the specified marked content id.
// Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
//
// If mcid is nil, no marked content is added.
//
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_deg Circle) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	_dcf := _deg.Width / 2
	_dgg := _deg.Height / 2
	if _deg.BorderEnabled {
		_dcf -= _deg.BorderWidth / 2
		_dgg -= _deg.BorderWidth / 2
	}
	_fd := 0.551784
	_bdc := _dcf * _fd
	_cde := _dgg * _fd
	_gcb := NewCubicBezierPath()
	_gcb = _gcb.AppendCurve(NewCubicBezierCurve(-_dcf, 0, -_dcf, _cde, -_bdc, _dgg, 0, _dgg))
	_gcb = _gcb.AppendCurve(NewCubicBezierCurve(0, _dgg, _bdc, _dgg, _dcf, _cde, _dcf, 0))
	_gcb = _gcb.AppendCurve(NewCubicBezierCurve(_dcf, 0, _dcf, -_cde, _bdc, -_dgg, 0, -_dgg))
	_gcb = _gcb.AppendCurve(NewCubicBezierCurve(0, -_dgg, -_bdc, -_dgg, -_dcf, -_cde, -_dcf, 0))
	_gcb = _gcb.Offset(_dcf, _dgg)
	if _deg.BorderEnabled {
		_gcb = _gcb.Offset(_deg.BorderWidth/2, _deg.BorderWidth/2)
	}
	if _deg.X != 0 || _deg.Y != 0 {
		_gcb = _gcb.Offset(_deg.X, _deg.Y)
	}
	_dd := _gb.NewContentCreator()
	if mcid != nil {
		_dd.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_dd.Add_q()
	if _deg.FillEnabled {
		_dd.SetNonStrokingColor(_deg.FillColor)
	}
	if _deg.BorderEnabled {
		_dd.SetStrokingColor(_deg.BorderColor)
		_dd.Add_w(_deg.BorderWidth)
	}
	if len(gsName) > 1 {
		_dd.Add_gs(_gc.PdfObjectName(gsName))
	}
	DrawBezierPathWithCreator(_gcb, _dd)
	_dd.Add_h()
	if _deg.FillEnabled && _deg.BorderEnabled {
		_dd.Add_B()
	} else if _deg.FillEnabled {
		_dd.Add_f()
	} else if _deg.BorderEnabled {
		_dd.Add_S()
	}
	_dd.Add_Q()
	if mcid != nil {
		_dd.Add_EMC()
	}
	_dde := _gcb.GetBoundingBox()
	if _deg.BorderEnabled {
		_dde.Height += _deg.BorderWidth
		_dde.Width += _deg.BorderWidth
		_dde.X -= _deg.BorderWidth / 2
		_dde.Y -= _deg.BorderWidth / 2
	}
	return _dd.Bytes(), _dde.ToPdfRectangle(), nil
}

const (
	LineStyleSolid  LineStyle = 0
	LineStyleDashed LineStyle = 1
)

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator(path Path, creator *_gb.ContentCreator) {
	for _cdg, _cafb := range path.Points {
		if _cdg == 0 {
			creator.Add_m(_cafb.X, _cafb.Y)
		} else {
			creator.Add_l(_cafb.X, _cafb.Y)
		}
	}
}

// Draw draws the basic line to PDF and marked the content using the specified marked content id.
// Generates the content stream which can be used in page contents or appearance stream of annotation.
//
// If mcid is nil, no marked content is added.
//
// Returns the stream content, XForm bounding box (local), bounding box and an error if one occurred.
func (_dfd BasicLine) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	_efc := NewPath()
	_efc = _efc.AppendPoint(NewPoint(_dfd.X1, _dfd.Y1))
	_efc = _efc.AppendPoint(NewPoint(_dfd.X2, _dfd.Y2))
	_abb := _gb.NewContentCreator()
	if mcid != nil {
		_abb.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_abb.Add_q().Add_w(_dfd.LineWidth).SetStrokingColor(_dfd.LineColor)
	if _dfd.LineStyle == LineStyleDashed {
		if _dfd.DashArray == nil {
			_dfd.DashArray = []int64{1, 1}
		}
		_abb.Add_d(_dfd.DashArray, _dfd.DashPhase)
	}
	if len(gsName) > 1 {
		_abb.Add_gs(_gc.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_efc, _abb)
	_abb.Add_S().Add_Q()
	if mcid != nil {
		_abb.Add_EMC()
	}
	return _abb.Bytes(), _efc.GetBoundingBox().ToPdfRectangle(), nil
}

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_ce CubicBezierCurve) AddOffsetXY(offX, offY float64) CubicBezierCurve {
	_ce.P0.X += offX
	_ce.P1.X += offX
	_ce.P2.X += offX
	_ce.P3.X += offX
	_ce.P0.Y += offY
	_ce.P1.Y += offY
	_ce.P2.Y += offY
	_ce.P3.Y += offY
	return _ce
}

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct {
	Curves      []CubicBezierCurve
	BorderWidth float64
	BorderColor _ff.PdfColor
	FillEnabled bool
	FillColor   _ff.PdfColor
}

// Point represents a two-dimensional point.
type Point struct {
	X float64
	Y float64
}

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{ Curves []CubicBezierCurve }

func (_efe Point) String() string {
	return _f.Sprintf("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029", _efe.X, _efe.Y)
}

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_bgb Rectangle) ToPdfRectangle() *_ff.PdfRectangle {
	return &_ff.PdfRectangle{Llx: _bgb.X, Lly: _bgb.Y, Urx: _bgb.X + _bgb.Width, Ury: _bgb.Y + _bgb.Height}
}

// Vector represents a two-dimensional vector.
type Vector struct {
	Dx float64
	Dy float64
}

// Scale scales the vector by the specified factor.
func (_bdf Vector) Scale(factor float64) Vector {
	_bccf := _bdf.Magnitude()
	_bca := _bdf.GetPolarAngle()
	_bdf.Dx = factor * _bccf * _g.Cos(_bca)
	_bdf.Dy = factor * _bccf * _g.Sin(_bca)
	return _bdf
}

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_ffa Vector) GetPolarAngle() float64 { return _g.Atan2(_ffa.Dy, _ffa.Dx) }

// Magnitude returns the magnitude of the vector.
func (_bff Vector) Magnitude() float64 { return _g.Sqrt(_g.Pow(_bff.Dx, 2.0) + _g.Pow(_bff.Dy, 2.0)) }

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_faf BasicLine) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _faf.MarkedDraw(gsName, nil)
}

// Draw draws the polygon and marked the content using the specified marked content id.
// A graphics state name can be specified for setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty ("").
//
// If mcid is nil, no marked content is added.
//
// Returns the content stream as a byte array and the polygon bounding box.
func (_bg Polygon) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	_gd := _gb.NewContentCreator()
	if mcid != nil {
		_gd.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_gd.Add_q()
	_bg.FillEnabled = _bg.FillEnabled && _bg.FillColor != nil
	if _bg.FillEnabled {
		_gd.SetNonStrokingColor(_bg.FillColor)
	}
	_bg.BorderEnabled = _bg.BorderEnabled && _bg.BorderColor != nil
	if _bg.BorderEnabled {
		_gd.SetStrokingColor(_bg.BorderColor)
		_gd.Add_w(_bg.BorderWidth)
	}
	if len(gsName) > 1 {
		_gd.Add_gs(_gc.PdfObjectName(gsName))
	}
	_gca := NewPath()
	for _, _gcaa := range _bg.Points {
		for _abg, _fbe := range _gcaa {
			_gca = _gca.AppendPoint(_fbe)
			if _abg == 0 {
				_gd.Add_m(_fbe.X, _fbe.Y)
			} else {
				_gd.Add_l(_fbe.X, _fbe.Y)
			}
		}
		_gd.Add_h()
	}
	if _bg.FillEnabled && _bg.BorderEnabled {
		_gd.Add_B()
	} else if _bg.FillEnabled {
		_gd.Add_f()
	} else if _bg.BorderEnabled {
		_gd.Add_S()
	}
	_gd.Add_Q()
	if mcid != nil {
		_gd.Add_EMC()
	}
	return _gd.Bytes(), _gca.GetBoundingBox().ToPdfRectangle(), nil
}

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct {

	// Position and size properties.
	X      float64
	Y      float64
	Width  float64
	Height float64

	// Fill properties.
	FillEnabled bool
	FillColor   _ff.PdfColor

	// Border properties.
	BorderEnabled           bool
	BorderColor             _ff.PdfColor
	BorderWidth             float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64

	// Shape opacity (0-1 interval).
	Opacity float64
}

// Draw draws the polyline and marked the content using the specified marked content id..
// A graphics state name can be specified for setting the polyline properties (e.g. setting the opacity).
// Otherwise leave empty ("").
//
// If mcid is nil, no marked content is added.
//
// Returns the content stream as a byte array and the polyline bounding box.
func (_fac Polyline) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	if _fac.LineColor == nil {
		_fac.LineColor = _ff.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_faca := NewPath()
	for _, _ebg := range _fac.Points {
		_faca = _faca.AppendPoint(_ebg)
	}
	_gfc := _gb.NewContentCreator()
	if mcid != nil {
		_gfc.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_gfc.Add_q().SetStrokingColor(_fac.LineColor).Add_w(_fac.LineWidth)
	if len(gsName) > 1 {
		_gfc.Add_gs(_gc.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_faca, _gfc)
	_gfc.Add_S()
	_gfc.Add_Q()
	if mcid != nil {
		_gfc.Add_EMC()
	}
	return _gfc.Bytes(), _faca.GetBoundingBox().ToPdfRectangle(), nil
}

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct {
	P0 Point
	P1 Point
	P2 Point
	P3 Point
}

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar(length float64, theta float64) Vector {
	_abf := Vector{}
	_abf.Dx = length * _g.Cos(theta)
	_abf.Dy = length * _g.Sin(theta)
	return _abf
}

const (
	LineEndingStyleNone  LineEndingStyle = 0
	LineEndingStyleArrow LineEndingStyle = 1
	LineEndingStyleButt  LineEndingStyle = 2
)

// Flip changes the sign of the vector: -vector.
func (_eba Vector) Flip() Vector {
	_edea := _eba.Magnitude()
	_bfd := _eba.GetPolarAngle()
	_eba.Dx = _edea * _g.Cos(_bfd+_g.Pi)
	_eba.Dy = _edea * _g.Sin(_bfd+_g.Pi)
	return _eba
}

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_gcc CurvePolygon) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _gcc.MarkedDraw(gsName, nil)
}

// FlipY flips the sign of the Dy component of the vector.
func (_fdef Vector) FlipY() Vector { _fdef.Dy = -_fdef.Dy; return _fdef }

// Copy returns a clone of the Bezier path.
func (_fa CubicBezierPath) Copy() CubicBezierPath {
	_eb := CubicBezierPath{}
	_eb.Curves = append(_eb.Curves, _fa.Curves...)
	return _eb
}

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween(a Point, b Point) Vector {
	_ggbe := Vector{}
	_ggbe.Dx = b.X - a.X
	_ggbe.Dy = b.Y - a.Y
	return _ggbe
}

// GetBounds returns the bounding box of the Bezier curve.
func (_gg CubicBezierCurve) GetBounds() _ff.PdfRectangle {
	_e := _gg.P0.X
	_ca := _gg.P0.X
	_ea := _gg.P0.Y
	_gbf := _gg.P0.Y
	for _dc := 0.0; _dc <= 1.0; _dc += 0.001 {
		Rx := _gg.P0.X*_g.Pow(1-_dc, 3) + _gg.P1.X*3*_dc*_g.Pow(1-_dc, 2) + _gg.P2.X*3*_g.Pow(_dc, 2)*(1-_dc) + _gg.P3.X*_g.Pow(_dc, 3)
		Ry := _gg.P0.Y*_g.Pow(1-_dc, 3) + _gg.P1.Y*3*_dc*_g.Pow(1-_dc, 2) + _gg.P2.Y*3*_g.Pow(_dc, 2)*(1-_dc) + _gg.P3.Y*_g.Pow(_dc, 3)
		if Rx < _e {
			_e = Rx
		}
		if Rx > _ca {
			_ca = Rx
		}
		if Ry < _ea {
			_ea = Ry
		}
		if Ry > _gbf {
			_gbf = Ry
		}
	}
	_bd := _ff.PdfRectangle{}
	_bd.Llx = _e
	_bd.Lly = _ea
	_bd.Urx = _ca
	_bd.Ury = _gbf
	return _bd
}

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_acf PolyBezierCurve) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _acf.MarkedDraw(gsName, nil)
}

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct {
	Points        [][]Point
	FillEnabled   bool
	FillColor     _ff.PdfColor
	BorderEnabled bool
	BorderColor   _ff.PdfColor
	BorderWidth   float64
}

// Draw draws the composite Bezier curve and marked the content using the specified marked content id.
// A graphics state name can be specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty ("").
//
// If mcid is nil, no marked content is added.
//
// Returns the content stream as a byte array and the curve bounding box.
func (_ced PolyBezierCurve) MarkedDraw(gsName string, mcid *int64) ([]byte, *_ff.PdfRectangle, error) {
	if _ced.BorderColor == nil {
		_ced.BorderColor = _ff.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_fba := NewCubicBezierPath()
	for _, _dcb := range _ced.Curves {
		_fba = _fba.AppendCurve(_dcb)
	}
	_ab := _gb.NewContentCreator()
	if mcid != nil {
		_ab.Add_BDC(*_gc.MakeName(_ff.StructureTypeFigure), map[string]_gc.PdfObject{"\u004d\u0043\u0049\u0044": _gc.MakeInteger(*mcid)})
	}
	_ab.Add_q()
	_ced.FillEnabled = _ced.FillEnabled && _ced.FillColor != nil
	if _ced.FillEnabled {
		_ab.SetNonStrokingColor(_ced.FillColor)
	}
	_ab.SetStrokingColor(_ced.BorderColor)
	_ab.Add_w(_ced.BorderWidth)
	if len(gsName) > 1 {
		_ab.Add_gs(_gc.PdfObjectName(gsName))
	}
	for _bcc, _gce := range _fba.Curves {
		if _bcc == 0 {
			_ab.Add_m(_gce.P0.X, _gce.P0.Y)
		} else {
			_ab.Add_l(_gce.P0.X, _gce.P0.Y)
		}
		_ab.Add_c(_gce.P1.X, _gce.P1.Y, _gce.P2.X, _gce.P2.Y, _gce.P3.X, _gce.P3.Y)
	}
	if _ced.FillEnabled {
		_ab.Add_h()
		_ab.Add_B()
	} else {
		_ab.Add_S()
	}
	_ab.Add_Q()
	if mcid != nil {
		_ab.Add_EMC()
	}
	return _ab.Bytes(), _fba.GetBoundingBox().ToPdfRectangle(), nil
}

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _ff.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _ff.PdfColor
	Opacity       float64
}

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct {
	Rings         [][]CubicBezierCurve
	FillEnabled   bool
	FillColor     _ff.PdfColor
	BorderEnabled bool
	BorderColor   _ff.PdfColor
	BorderWidth   float64
}

// Add adds the specified vector to the current one and returns the result.
func (_dfc Vector) Add(other Vector) Vector { _dfc.Dx += other.Dx; _dfc.Dy += other.Dy; return _dfc }

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Rotate rotates the vector by the specified angle.
func (_egc Vector) Rotate(phi float64) Vector {
	_def := _egc.Magnitude()
	_facad := _egc.GetPolarAngle()
	return NewVectorPolar(_def, _facad+phi)
}

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_cdaa Polyline) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _cdaa.MarkedDraw(gsName, nil)
}

// Offset shifts the Bezier path with the specified offsets.
func (_bf CubicBezierPath) Offset(offX, offY float64) CubicBezierPath {
	for _fe, _feg := range _bf.Curves {
		_bf.Curves[_fe] = _feg.AddOffsetXY(offX, offY)
	}
	return _bf
}

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        _ff.PdfColor
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 LineEndingStyle
	LineEndingStyle2 LineEndingStyle
	LineStyle        LineStyle
}

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct {
	Points    []Point
	LineColor _ff.PdfColor
	LineWidth float64
}

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator(bpath CubicBezierPath, creator *_gb.ContentCreator) {
	for _bceb, _bbe := range bpath.Curves {
		if _bceb == 0 {
			creator.Add_m(_bbe.P0.X, _bbe.P0.Y)
		}
		creator.Add_c(_bbe.P1.X, _bbe.P1.Y, _bbe.P2.X, _bbe.P2.Y, _bbe.P3.X, _bbe.P3.Y)
	}
}

// FlipX flips the sign of the Dx component of the vector.
func (_ddba Vector) FlipX() Vector { _ddba.Dx = -_ddba.Dx; return _ddba }

// GetBoundingBox returns the bounding box of the Bezier path.
func (_bc CubicBezierPath) GetBoundingBox() Rectangle {
	_fb := Rectangle{}
	_bed := 0.0
	_fbd := 0.0
	_ffd := 0.0
	_gf := 0.0
	for _bcf, _bec := range _bc.Curves {
		_fab := _bec.GetBounds()
		if _bcf == 0 {
			_bed = _fab.Llx
			_fbd = _fab.Urx
			_ffd = _fab.Lly
			_gf = _fab.Ury
			continue
		}
		if _fab.Llx < _bed {
			_bed = _fab.Llx
		}
		if _fab.Urx > _fbd {
			_fbd = _fab.Urx
		}
		if _fab.Lly < _ffd {
			_ffd = _fab.Lly
		}
		if _fab.Ury > _gf {
			_gf = _fab.Ury
		}
	}
	_fb.X = _bed
	_fb.Y = _ffd
	_fb.Width = _fbd - _bed
	_fb.Height = _gf - _ffd
	return _fb
}

// LineStyle refers to how the line will be created.
type LineStyle int

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector(dx, dy float64) Vector { _afc := Vector{}; _afc.Dx = dx; _afc.Dy = dy; return _afc }

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_ecbe Polygon) Draw(gsName string) ([]byte, *_ff.PdfRectangle, error) {
	return _ecbe.MarkedDraw(gsName, nil)
}
