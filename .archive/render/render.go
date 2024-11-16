package render

import (
	_e "errors"
	_bbe "fmt"
	_d "image"
	_fg "image/color"
	_ga "image/draw"
	_eb "image/jpeg"
	_bb "image/png"
	_ge "math"
	_b "os"
	_f "path/filepath"
	_ec "strings"

	_bd "github.com/adrg/sysfont"
	_bf "github.com/bamzi/pdfext/annotator"
	_fa "github.com/bamzi/pdfext/common"
	_aa "github.com/bamzi/pdfext/contentstream"
	_a "github.com/bamzi/pdfext/contentstream/draw"
	_dg "github.com/bamzi/pdfext/core"
	_bc "github.com/bamzi/pdfext/internal/license"
	_fae "github.com/bamzi/pdfext/internal/transform"
	_dc "github.com/bamzi/pdfext/model"
	_ad "github.com/bamzi/pdfext/render/internal/context"
	_fe "github.com/bamzi/pdfext/render/internal/context/imagerender"
	_ece "golang.org/x/image/draw"
)

var (
	_ca = _e.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_da = _e.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func _ffeb(_cbab *_dc.Image, _bda _fg.Color) _d.Image {
	_def, _gggb := int(_cbab.Width), int(_cbab.Height)
	_cacg := _d.NewRGBA(_d.Rect(0, 0, _def, _gggb))
	for _bdae := 0; _bdae < _gggb; _bdae++ {
		for _fagb := 0; _fagb < _def; _fagb++ {
			_bfbd, _fgde := _cbab.ColorAt(_fagb, _bdae)
			if _fgde != nil {
				_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _fagb, _bdae)
				continue
			}
			_gcad, _fcg, _cacc, _ := _bfbd.RGBA()
			var _cae _fg.Color
			if _gcad+_fcg+_cacc == 0 {
				_cae = _bda
			} else {
				_cae = _fg.Transparent
			}
			_cacg.Set(_fagb, _bdae, _cae)
		}
	}
	return _cacg
}

type renderer struct{ _aab float64 }

const (
	ShadingTypeFunctionBased PdfShadingType = 1
	ShadingTypeAxial         PdfShadingType = 2
	ShadingTypeRadial        PdfShadingType = 3
	ShadingTypeFreeForm      PdfShadingType = 4
	ShadingTypeLatticeForm   PdfShadingType = 5
	ShadingTypeCoons         PdfShadingType = 6
	ShadingTypeTensorProduct PdfShadingType = 7
)

func _gbe(_adg, _bcdd, _gace float64) _a.BoundingBox {
	return _a.Path{Points: []_a.Point{_a.NewPoint(0, 0).Rotate(_gace), _a.NewPoint(_adg, 0).Rotate(_gace), _a.NewPoint(0, _bcdd).Rotate(_gace), _a.NewPoint(_adg, _bcdd).Rotate(_gace)}}.GetBoundingBox()
}
func _cggc(_gefb, _ddc _d.Image) _d.Image {
	_dfg, _fcea := _ddc.Bounds().Size(), _gefb.Bounds().Size()
	_eed, _gfde := _dfg.X, _dfg.Y
	if _fcea.X > _eed {
		_eed = _fcea.X
	}
	if _fcea.Y > _gfde {
		_gfde = _fcea.Y
	}
	_abff := _d.Rect(0, 0, _eed, _gfde)
	if _dfg.X != _eed || _dfg.Y != _gfde {
		_ebga := _d.NewRGBA(_abff)
		_ece.BiLinear.Scale(_ebga, _abff, _gefb, _ddc.Bounds(), _ece.Over, nil)
		_ddc = _ebga
	}
	if _fcea.X != _eed || _fcea.Y != _gfde {
		_fgbd := _d.NewRGBA(_abff)
		_ece.BiLinear.Scale(_fgbd, _abff, _gefb, _gefb.Bounds(), _ece.Over, nil)
		_gefb = _fgbd
	}
	_gbf := _d.NewRGBA(_abff)
	_ece.DrawMask(_gbf, _abff, _gefb, _d.Point{}, _ddc, _d.Point{}, _ece.Over)
	return _gbf
}
func _cba(_fedf string, _fag _d.Image) error {
	_cfc, _ddb := _b.Create(_fedf)
	if _ddb != nil {
		return _ddb
	}
	defer _cfc.Close()
	return _bb.Encode(_cfc, _fag)
}
func (_degf renderer) processLinearShading(_afc _ad.Context, _abef *_dc.PdfShading) (_ad.Gradient, *_dg.PdfObjectArray, error) {
	_edb := _abef.GetContext().(*_dc.PdfShadingType2)
	if len(_edb.Function) == 0 {
		return nil, nil, _e.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_caca, _afce := _edb.Coords.ToFloat64Array()
	if _afce != nil {
		return nil, nil, _e.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_abf := _abef.ColorSpace
	_eabe, _ebfe := _afc.Matrix().Transform(_caca[0], _caca[1])
	_cca, _aedd := _afc.Matrix().Transform(_caca[2], _caca[3])
	_ccf := _fe.NewLinearGradient(_eabe, _ebfe, _cca, _aedd)
	_bff := _dg.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	for _, _fde := range _caca {
		if _fde > 1 {
			_bff = _edb.Coords
			break
		}
	}
	if _fdfc, _afdb := _edb.Function[0].(*_dc.PdfFunctionType2); _afdb {
		_ccf, _afce = _dad(_ccf, _fdfc, _abf, 1.0, true)
	} else if _caae, _bdf := _edb.Function[0].(*_dc.PdfFunctionType3); _bdf {
		_fce := append([]float64{0}, _caae.Bounds...)
		_fce = append(_fce, 1.0)
		_ccf, _afce = _gbb(_ccf, _caae, _abf, _fce)
	}
	return _ccf, _bff, _afce
}

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct {
	renderer

	// OutputWidth represents the width of the rendered images in pixels.
	// The heights of the output images are calculated based on the selected
	// width and the original height of each rendered page.
	OutputWidth int
}

func (_bgbb renderer) processShading(_cgf _ad.Context, _dagb *_dc.PdfShading) (_ad.Gradient, *_dg.PdfObjectArray, error) {
	_bgac := int64(*_dagb.ShadingType)
	if _bgac == int64(ShadingTypeAxial) {
		return _bgbb.processLinearShading(_cgf, _dagb)
	} else if _bgac == int64(ShadingTypeRadial) {
		return _bgbb.processRadialShading(_cgf, _dagb)
	} else {
		_fa.Log.Debug(_bbe.Sprintf("\u0050r\u006f\u0063e\u0073\u0073\u0069n\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0074\u0079\u0070e\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074 \u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064", _bgac))
	}
	return nil, nil, nil
}
func _gcc(_cde string, _aee _d.Image, _efa int) error {
	_aedb, _cfg := _b.Create(_cde)
	if _cfg != nil {
		return _cfg
	}
	defer _aedb.Close()
	return _eb.Encode(_aedb, _aee, &_eb.Options{Quality: _efa})
}
func _ffef(_afad _dg.PdfObject, _gffe _fg.Color) (_d.Image, error) {
	_gcdd, _fgbf := _dg.GetStream(_afad)
	if !_fgbf {
		return nil, nil
	}
	_dbgb, _caef := _dc.NewXObjectImageFromStream(_gcdd)
	if _caef != nil {
		return nil, _caef
	}
	_gcaf, _caef := _dbgb.ToImage()
	if _caef != nil {
		return nil, _caef
	}
	return _ffeb(_gcaf, _gffe), nil
}
func (_bbb renderer) processGradient(_db _ad.Context, _fbcf *_aa.ContentStreamOperation, _bga *_dc.PdfPageResources, _dba *_dg.PdfObjectName) (_ad.Gradient, error) {
	if _ccg, _afgf := _bga.GetPatternByName(*_dba); _afgf && _ccg.IsShading() {
		_adaa := _ccg.GetAsShadingPattern().Shading
		_cac, _, _fcbf := _bbb.processShading(_db, _adaa)
		if _fcbf != nil {
			return nil, _fcbf
		}
		return _cac, nil
	}
	return nil, nil
}
func (_bead renderer) processRadialShading(_dfbb _ad.Context, _gcab *_dc.PdfShading) (_ad.Gradient, *_dg.PdfObjectArray, error) {
	_bacg := _gcab.GetContext().(*_dc.PdfShadingType3)
	if len(_bacg.Function) == 0 {
		return nil, nil, _e.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_dgc, _ecge := _bacg.Coords.ToFloat64Array()
	if _ecge != nil {
		return nil, nil, _e.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_gfe := _gcab.ColorSpace
	_bgcb := _dg.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	var _cef, _ega, _cdcg, _bbfg, _fbcd, _gadg float64
	_cef, _ega = _dfbb.Matrix().Transform(_dgc[0], _dgc[1])
	_cdcg, _bbfg = _dfbb.Matrix().Transform(_dgc[3], _dgc[4])
	_fbcd, _ = _dfbb.Matrix().Transform(_dgc[2], 0)
	_gadg, _ = _dfbb.Matrix().Transform(_dgc[5], 0)
	_ggfb, _ := _dfbb.Matrix().Translation()
	_fbcd -= _ggfb
	_gadg -= _ggfb
	for _fdd, _bfbe := range _dgc {
		if _fdd == 2 || _fdd == 5 {
			continue
		}
		if _bfbe > 1.0 {
			_fcef := _ge.Min(_cef-_fbcd, _cdcg-_gadg)
			_aegc := _ge.Min(_ega-_fbcd, _bbfg-_gadg)
			_ecbd := _ge.Max(_cef+_fbcd, _cdcg+_gadg)
			_gdfa := _ge.Max(_ega+_fbcd, _bbfg+_gadg)
			_fbac := _ecbd - _fcef
			_eec := _aegc - _gdfa
			_bgcb = _dg.MakeArrayFromFloats([]float64{_fcef, _aegc, _fbac, _eec})
			break
		}
	}
	_bfbf := _fe.NewRadialGradient(_cef, _ega, _fbcd, _cdcg, _bbfg, _gadg)
	if _fbeb, _fac := _bacg.Function[0].(*_dc.PdfFunctionType2); _fac {
		_bfbf, _ecge = _dad(_bfbf, _fbeb, _gfe, 1.0, true)
	} else if _acg, _dbaf := _bacg.Function[0].(*_dc.PdfFunctionType3); _dbaf {
		_cfda := append([]float64{0}, _acg.Bounds...)
		_cfda = append(_cfda, 1.0)
		_bfbf, _ecge = _gbb(_bfbf, _acg, _gfe, _cfda)
	}
	if _ecge != nil {
		return nil, nil, _ecge
	}
	return _bfbf, _bgcb, nil
}
func _ceg(_dgcb _dg.PdfObject, _affa _fg.Color) (_d.Image, error) {
	_dbe, _dgf := _dg.GetStream(_dgcb)
	if !_dgf {
		return nil, nil
	}
	_geac, _ageg := _dc.NewXObjectImageFromStream(_dbe)
	if _ageg != nil {
		return nil, _ageg
	}
	_fcfg, _ageg := _geac.ToImage()
	if _ageg != nil {
		return nil, _ageg
	}
	return _gfee(_fcfg, _affa), nil
}

// Render converts the specified PDF page into an image, flattens annotations by default and returns the result.
func (_fgb *ImageDevice) Render(page *_dc.PdfPage) (_d.Image, error) {
	return _fgb.RenderWithOpts(page, false)
}

// PdfShadingType defines PDF shading types.
// Source: PDF32000_2008.pdf. Chapter 8.7.4.5
type PdfShadingType int64

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_gad *ImageDevice) RenderToPath(page *_dc.PdfPage, outputPath string) error {
	_cec, _bfd := _gad.Render(page)
	if _bfd != nil {
		return _bfd
	}
	_aad := _ec.ToLower(_f.Ext(outputPath))
	if _aad == "" {
		return _e.New("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065")
	}
	switch _aad {
	case "\u002e\u0070\u006e\u0067":
		return _cba(outputPath, _cec)
	case "\u002e\u006a\u0070\u0067", "\u002e\u006a\u0070e\u0067":
		return _gcc(outputPath, _cec, 100)
	}
	return _bbe.Errorf("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _aad)
}
func _gbb(_gdg _ad.Gradient, _fbec *_dc.PdfFunctionType3, _fcd _dc.PdfColorspace, _gag []float64) (_ad.Gradient, error) {
	var _egd error
	for _efb := 0; _efb < len(_fbec.Functions); _efb++ {
		if _gdff, _abc := _fbec.Functions[_efb].(*_dc.PdfFunctionType2); _abc {
			_gdg, _egd = _dad(_gdg, _gdff, _fcd, _gag[_efb+1], _efb == 0)
			if _egd != nil {
				return nil, _egd
			}
		}
	}
	return _gdg, nil
}
func (_gg renderer) renderPage(_fgc _ad.Context, _ac *_dc.PdfPage, _bdc _fae.Matrix, _ebg bool) error {
	if !_ebg {
		_gce := _dc.FieldFlattenOpts{AnnotFilterFunc: func(_cge *_dc.PdfAnnotation) bool {
			switch _cge.GetContext().(type) {
			case *_dc.PdfAnnotationLine:
				return true
			case *_dc.PdfAnnotationSquare:
				return true
			case *_dc.PdfAnnotationCircle:
				return true
			case *_dc.PdfAnnotationPolygon:
				return true
			case *_dc.PdfAnnotationPolyLine:
				return true
			}
			return false
		}}
		_ea := _bf.FieldAppearance{}
		_ff := _ac.FlattenFieldsWithOpts(_ea, &_gce)
		if _ff != nil {
			_fa.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0064u\u0072\u0069n\u0067\u0020\u0061\u006e\u006e\u006f\u0074\u0061t\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0074\u0074\u0065\u006e\u0069n\u0067\u0020\u0025\u0076", _ff)
		}
	}
	_aaf, _abg := _ac.GetAllContentStreams()
	if _abg != nil {
		return _abg
	}
	if _fb := _bdc; !_fb.Identity() {
		_aaf = _bbe.Sprintf("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073", _fb[0], _fb[1], _fb[3], _fb[4], _fb[6], _fb[7], _aaf)
	}
	_fgc.Translate(0, float64(_fgc.Height()))
	_fgc.Scale(1, -1)
	_fgc.Push()
	_fgc.SetRGBA(1, 1, 1, 1)
	_fgc.DrawRectangle(0, 0, float64(_fgc.Width()), float64(_fgc.Height()))
	_fgc.Fill()
	_fgc.Pop()
	_fgc.SetLineWidth(1.0)
	_fgc.SetRGBA(0, 0, 0, 1)
	return _gg.renderContentStream(_fgc, _aaf, _ac.Resources)
}
func (_ebb renderer) renderContentStream(_ffe _ad.Context, _faee string, _bac *_dc.PdfPageResources) error {
	_cbg, _aabb := _aa.NewContentStreamParser(_faee).Parse()
	if _aabb != nil {
		return _aabb
	}
	_bgf := _ffe.TextState()
	_bgf.GlobalScale = _ebb._aab
	_fbf := map[string]*_ad.TextFont{}
	_de := _bd.NewFinder(&_bd.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	var _ged *_aa.ContentStreamOperation
	_ecb := _aa.NewContentStreamProcessor(*_cbg)
	_ecb.AddHandler(_aa.HandlerConditionEnumAllOperands, "", func(_bbf *_aa.ContentStreamOperation, _cc _aa.GraphicsState, _ade *_dc.PdfPageResources) error {
		_fa.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073", _bbf.Operand)
		switch _bbf.Operand {
		case "\u0071":
			_ffe.Push()
		case "\u0051":
			_ffe.Pop()
			_bgf = _ffe.TextState()
		case "\u0063\u006d":
			if len(_bbf.Params) != 6 {
				return _da
			}
			_fdg, _cgg := _dg.GetNumbersAsFloat(_bbf.Params)
			if _cgg != nil {
				return _cgg
			}
			_abd := _fae.NewMatrix(_fdg[0], _fdg[1], _fdg[2], _fdg[3], _fdg[4], _fdg[5])
			_fa.Log.Debug("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076", _abd)
			_ffe.SetMatrix(_ffe.Matrix().Mult(_abd))
		case "\u0077":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_gcb, _dcf := _dg.GetNumbersAsFloat(_bbf.Params)
			if _dcf != nil {
				return _dcf
			}
			_ffe.SetLineWidth(_gcb[0])
		case "\u004a":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_cdc, _fbe := _dg.GetIntVal(_bbf.Params[0])
			if !_fbe {
				return _ca
			}
			switch _cdc {
			case 0:
				_ffe.SetLineCap(_ad.LineCapButt)
			case 1:
				_ffe.SetLineCap(_ad.LineCapRound)
			case 2:
				_ffe.SetLineCap(_ad.LineCapSquare)
			default:
				_fa.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064", _cdc)
				return _da
			}
		case "\u006a":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_dd, _fafe := _dg.GetIntVal(_bbf.Params[0])
			if !_fafe {
				return _ca
			}
			switch _dd {
			case 0:
				_ffe.SetLineJoin(_ad.LineJoinBevel)
			case 1:
				_ffe.SetLineJoin(_ad.LineJoinRound)
			case 2:
				_ffe.SetLineJoin(_ad.LineJoinBevel)
			default:
				_fa.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064", _dd)
				return _da
			}
		case "\u004d":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_ecec, _cf := _dg.GetNumbersAsFloat(_bbf.Params)
			if _cf != nil {
				return _cf
			}
			_ = _ecec
			_fa.Log.Debug("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d")
		case "\u0064":
			if len(_bbf.Params) != 2 {
				return _da
			}
			_afd, _fgcb := _dg.GetArray(_bbf.Params[0])
			if !_fgcb {
				return _ca
			}
			_adc, _fgcb := _dg.GetIntVal(_bbf.Params[1])
			if !_fgcb {
				_, _ee := _dg.GetFloatVal(_bbf.Params[1])
				if !_ee {
					return _ca
				}
			}
			_ag, _bgg := _dg.GetNumbersAsFloat(_afd.Elements())
			if _bgg != nil {
				return _bgg
			}
			_ffe.SetDash(_ag...)
			_ = _adc
			_fa.Log.Debug("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064")
		case "\u0072\u0069":
			_fa.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0069":
			_fa.Log.Debug("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0067\u0073":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_fbd, _eda := _dg.GetName(_bbf.Params[0])
			if !_eda {
				return _ca
			}
			if _fbd == nil {
				return _da
			}
			_bfc, _eda := _ade.GetExtGState(*_fbd)
			if !_eda {
				_fa.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073", *_fbd)
				return _e.New("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064")
			}
			_ecbf, _eda := _dg.GetDict(_bfc)
			if !_eda {
				_fa.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074")
				return _ca
			}
			_fa.Log.Debug("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _ecbf.String())
			_cdb := _ecbf.Get("\u0063\u0061")
			if _cdb != nil {
				_fc, _dge := _dg.GetNumberAsFloat(_cdb)
				if _dge == nil {
					_ebf, _ace := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
					if _ace != nil {
						_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ace)
						return _ace
					}
					_dgd, _bgd := _ebf.(*_dc.PdfColorDeviceRGB)
					if !_bgd {
						_fa.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
						return _ace
					}
					_ffe.SetFillRGBA(_dgd.R(), _dgd.G(), _dgd.B(), _fc)
				}
			}
		case "\u006d":
			if len(_bbf.Params) != 2 {
				_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _da)
				return nil
			}
			_aed, _afg := _dg.GetNumbersAsFloat(_bbf.Params)
			if _afg != nil {
				return _afg
			}
			_fa.Log.Debug("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076", _aed)
			_ffe.NewSubPath()
			_ffe.MoveTo(_aed[0], _aed[1])
		case "\u006c":
			if len(_bbf.Params) != 2 {
				_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _da)
				return nil
			}
			_age, _gca := _dg.GetNumbersAsFloat(_bbf.Params)
			if _gca != nil {
				return _gca
			}
			_ffe.LineTo(_age[0], _age[1])
		case "\u0063":
			if len(_bbf.Params) != 6 {
				return _da
			}
			_be, _fbg := _dg.GetNumbersAsFloat(_bbf.Params)
			if _fbg != nil {
				return _fbg
			}
			_fa.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _be)
			_ffe.CubicTo(_be[0], _be[1], _be[2], _be[3], _be[4], _be[5])
		case "\u0076", "\u0079":
			if len(_bbf.Params) != 4 {
				return _da
			}
			_faef, _dae := _dg.GetNumbersAsFloat(_bbf.Params)
			if _dae != nil {
				return _dae
			}
			_fa.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _faef)
			_ffe.QuadraticTo(_faef[0], _faef[1], _faef[2], _faef[3])
		case "\u0068":
			_ffe.ClosePath()
			_ffe.NewSubPath()
		case "\u0072\u0065":
			if len(_bbf.Params) != 4 {
				return _da
			}
			_aabe, _ggf := _dg.GetNumbersAsFloat(_bbf.Params)
			if _ggf != nil {
				return _ggf
			}
			_ffe.DrawRectangle(_aabe[0], _aabe[1], _aabe[2], _aabe[3])
			_ffe.NewSubPath()
		case "\u0053":
			_bfb, _gddc := _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
			if _gddc != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gddc)
				return _gddc
			}
			_ebd, _gec := _bfb.(*_dc.PdfColorDeviceRGB)
			if !_gec {
				_fa.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _gddc
			}
			_ffe.SetRGBA(_ebd.R(), _ebd.G(), _ebd.B(), 1)
			_ffe.Stroke()
		case "\u0073":
			_edc, _acc := _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
			if _acc != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _acc)
				return _acc
			}
			_bbd, _gef := _edc.(*_dc.PdfColorDeviceRGB)
			if !_gef {
				_fa.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _acc
			}
			_ffe.ClosePath()
			_ffe.NewSubPath()
			_ffe.SetRGBA(_bbd.R(), _bbd.G(), _bbd.B(), 1)
			_ffe.Stroke()
		case "\u0066", "\u0046":
			_eea, _bcd := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _bcd != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bcd)
				return _bcd
			}
			switch _fbfc := _eea.(type) {
			case *_dc.PdfColorDeviceRGB:
				_ffe.SetRGBA(_fbfc.R(), _fbfc.G(), _fbfc.B(), 1)
				_ffe.SetFillRule(_ad.FillRuleWinding)
				_ffe.Fill()
			case *_dc.PdfColorPattern:
				_ffe.Fill()
			}
			_fa.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
		case "\u0066\u002a":
			_aaa, _bce := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _bce != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bce)
				return _bce
			}
			_gea, _bcef := _aaa.(*_dc.PdfColorDeviceRGB)
			if !_bcef {
				_fa.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _bce
			}
			_ffe.SetRGBA(_gea.R(), _gea.G(), _gea.B(), 1)
			_ffe.SetFillRule(_ad.FillRuleEvenOdd)
			_ffe.Fill()
		case "\u0042":
			_bgb, _cggd := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _cggd != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cggd)
				return _cggd
			}
			switch _gdb := _bgb.(type) {
			case *_dc.PdfColorDeviceRGB:
				_ffe.SetRGBA(_gdb.R(), _gdb.G(), _gdb.B(), 1)
				_ffe.SetFillRule(_ad.FillRuleWinding)
				_ffe.FillPreserve()
				_bgb, _cggd = _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
				if _cggd != nil {
					_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cggd)
					return _cggd
				}
				if _cdf, _dfb := _bgb.(*_dc.PdfColorDeviceRGB); _dfb {
					_ffe.SetRGBA(_cdf.R(), _cdf.G(), _cdf.B(), 1)
					_ffe.Stroke()
				}
			case *_dc.PdfColorPattern:
				_ffe.SetFillRule(_ad.FillRuleWinding)
				_ffe.Fill()
				_ffe.StrokePattern()
			}
		case "\u0042\u002a":
			_fca, _feb := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _feb != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _feb)
				return _feb
			}
			switch _add := _fca.(type) {
			case *_dc.PdfColorDeviceRGB:
				_ffe.SetRGBA(_add.R(), _add.G(), _add.B(), 1)
				_ffe.SetFillRule(_ad.FillRuleEvenOdd)
				_ffe.FillPreserve()
				_fca, _feb = _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
				if _feb != nil {
					_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _feb)
					return _feb
				}
				if _bcde, _acd := _fca.(*_dc.PdfColorDeviceRGB); _acd {
					_ffe.SetRGBA(_bcde.R(), _bcde.G(), _bcde.B(), 1)
					_ffe.Stroke()
				}
			case *_dc.PdfColorPattern:
				_ffe.SetFillRule(_ad.FillRuleEvenOdd)
				_ffe.Fill()
				_ffe.StrokePattern()
			}
		case "\u0062":
			_ffe.ClosePath()
			_adcc, _cfe := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _cfe != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cfe)
				return _cfe
			}
			switch _dcfg := _adcc.(type) {
			case *_dc.PdfColorDeviceRGB:
				_ffe.SetRGBA(_dcfg.R(), _dcfg.G(), _dcfg.B(), 1)
				_ffe.NewSubPath()
				_ffe.SetFillRule(_ad.FillRuleWinding)
				_ffe.FillPreserve()
				_adcc, _cfe = _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
				if _cfe != nil {
					_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cfe)
					return _cfe
				}
				if _aec, _geaa := _adcc.(*_dc.PdfColorDeviceRGB); _geaa {
					_ffe.SetRGBA(_aec.R(), _aec.G(), _aec.B(), 1)
					_ffe.Stroke()
				}
			case *_dc.PdfColorPattern:
				_ffe.NewSubPath()
				_ffe.SetFillRule(_ad.FillRuleWinding)
				_ffe.Fill()
				_ffe.StrokePattern()
			}
		case "\u0062\u002a":
			_ffe.ClosePath()
			_cfd, _gcd := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _gcd != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcd)
				return _gcd
			}
			switch _cad := _cfd.(type) {
			case *_dc.PdfColorDeviceRGB:
				_ffe.SetRGBA(_cad.R(), _cad.G(), _cad.B(), 1)
				_ffe.NewSubPath()
				_ffe.SetFillRule(_ad.FillRuleEvenOdd)
				_ffe.FillPreserve()
				_cfd, _gcd = _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
				if _gcd != nil {
					_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcd)
					return _gcd
				}
				if _ege, _dce := _cfd.(*_dc.PdfColorDeviceRGB); _dce {
					_ffe.SetRGBA(_ege.R(), _ege.G(), _ege.B(), 1)
					_ffe.Stroke()
				}
			case *_dc.PdfColorPattern:
				_ffe.NewSubPath()
				_ffe.SetFillRule(_ad.FillRuleEvenOdd)
				_ffe.Fill()
				_ffe.StrokePattern()
			}
		case "\u006e":
			_ffe.ClearPath()
		case "\u0057":
			_ffe.SetFillRule(_ad.FillRuleWinding)
			_ffe.ClipPreserve()
		case "\u0057\u002a":
			_ffe.SetFillRule(_ad.FillRuleEvenOdd)
			_ffe.ClipPreserve()
		case "\u0072\u0067":
			_fga, _dfbc := _cc.ColorNonStroking.(*_dc.PdfColorDeviceRGB)
			if !_dfbc {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_ffe.SetFillRGBA(_fga.R(), _fga.G(), _fga.B(), 1)
		case "\u0052\u0047":
			_cdcb, _fdf := _cc.ColorStroking.(*_dc.PdfColorDeviceRGB)
			if !_fdf {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_ffe.SetStrokeRGBA(_cdcb.R(), _cdcb.G(), _cdcb.B(), 1)
		case "\u006b":
			_bbed, _eeaa := _cc.ColorNonStroking.(*_dc.PdfColorDeviceCMYK)
			if !_eeaa {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_ecf, _bef := _cc.ColorspaceNonStroking.ColorToRGB(_bbed)
			if _bef != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_ffg, _eeaa := _ecf.(*_dc.PdfColorDeviceRGB)
			if !_eeaa {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ecf)
				return nil
			}
			_ffe.SetFillRGBA(_ffg.R(), _ffg.G(), _ffg.B(), 1)
		case "\u004b":
			_befe, _dggc := _cc.ColorStroking.(*_dc.PdfColorDeviceCMYK)
			if !_dggc {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_fbb, _afda := _cc.ColorspaceStroking.ColorToRGB(_befe)
			if _afda != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_deg, _dggc := _fbb.(*_dc.PdfColorDeviceRGB)
			if !_dggc {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fbb)
				return nil
			}
			_ffe.SetStrokeRGBA(_deg.R(), _deg.G(), _deg.B(), 1)
		case "\u0067":
			_eeg, _dda := _cc.ColorNonStroking.(*_dc.PdfColorDeviceGray)
			if !_dda {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_efd, _bfdb := _cc.ColorspaceNonStroking.ColorToRGB(_eeg)
			if _bfdb != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_abdf, _dda := _efd.(*_dc.PdfColorDeviceRGB)
			if !_dda {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _efd)
				return nil
			}
			_ffe.SetFillRGBA(_abdf.R(), _abdf.G(), _abdf.B(), 1)
		case "\u0047":
			_fdga, _dcd := _cc.ColorStroking.(*_dc.PdfColorDeviceGray)
			if !_dcd {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_bdd, _adb := _cc.ColorspaceStroking.ColorToRGB(_fdga)
			if _adb != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_egeb, _dcd := _bdd.(*_dc.PdfColorDeviceRGB)
			if !_dcd {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bdd)
				return nil
			}
			_ffe.SetStrokeRGBA(_egeb.R(), _egeb.G(), _egeb.B(), 1)
		case "\u0063\u0073":
			if len(_bbf.Params) > 0 {
				if _cgb, _bgc := _dg.GetName(_bbf.Params[0]); _bgc && _cgb.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_fgd, _dgaf := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _dgaf != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_cbe, _fdb := _fgd.(*_dc.PdfColorDeviceRGB)
			if !_fdb {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fgd)
				return nil
			}
			_ffe.SetFillRGBA(_cbe.R(), _cbe.G(), _cbe.B(), 1)
		case "\u0073\u0063":
			_dfc, _fec := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _fec != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_dab, _agee := _dfc.(*_dc.PdfColorDeviceRGB)
			if !_agee {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dfc)
				return nil
			}
			_ffe.SetFillRGBA(_dab.R(), _dab.G(), _dab.B(), 1)
		case "\u0073\u0063\u006e":
			if len(_bbf.Params) > 0 && len(_ged.Params) > 0 {
				if _dggg, _efde := _dg.GetName(_ged.Params[0]); _efde && _dggg.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _cgc, _efe := _dg.GetName(_bbf.Params[0]); _efe {
						_bea, _gb := _ebb.processGradient(_ffe, _bbf, _ade, _cgc)
						if _gb != nil {
							_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _gb)
							break
						}
						if _bea == nil {
							_fa.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_ffe.SetFillStyle(_bea)
						_ffe.SetStrokeStyle(_bea)
						break
					}
				}
			}
			_eag, _egc := _cc.ColorspaceNonStroking.ColorToRGB(_cc.ColorNonStroking)
			if _egc != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorNonStroking)
				return nil
			}
			_gf, _dgb := _eag.(*_dc.PdfColorDeviceRGB)
			if !_dgb {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _eag)
				return nil
			}
			_ffe.SetFillRGBA(_gf.R(), _gf.G(), _gf.B(), 1)
		case "\u0043\u0053":
			if len(_bbf.Params) > 0 {
				if _fee, _acb := _dg.GetName(_bbf.Params[0]); _acb && _fee.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_caf, _ecg := _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
			if _ecg != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_afe, _gda := _caf.(*_dc.PdfColorDeviceRGB)
			if !_gda {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _caf)
				return nil
			}
			_ffe.SetStrokeRGBA(_afe.R(), _afe.G(), _afe.B(), 1)
		case "\u0053\u0043":
			_fda, _ede := _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
			if _ede != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_dff, _ddab := _fda.(*_dc.PdfColorDeviceRGB)
			if !_ddab {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fda)
				return nil
			}
			_ffe.SetStrokeRGBA(_dff.R(), _dff.G(), _dff.B(), 1)
		case "\u0053\u0043\u004e":
			if len(_bbf.Params) > 0 && len(_ged.Params) > 0 {
				if _dfa, _gfd := _dg.GetName(_ged.Params[0]); _gfd && _dfa.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _gga, _afga := _dg.GetName(_bbf.Params[0]); _afga {
						_fcf, _gac := _ebb.processGradient(_ffe, _bbf, _ade, _gga)
						if _gac != nil {
							_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _gac)
							break
						}
						if _fcf == nil {
							_fa.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_ffe.SetFillStyle(_fcf)
						_ffe.SetStrokeStyle(_fcf)
						break
					}
				}
			}
			_ffa, _ggd := _cc.ColorspaceStroking.ColorToRGB(_cc.ColorStroking)
			if _ggd != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cc.ColorStroking)
				return nil
			}
			_aeg, _cbd := _ffa.(*_dc.PdfColorDeviceRGB)
			if !_cbd {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ffa)
				return nil
			}
			_ffe.SetStrokeRGBA(_aeg.R(), _aeg.G(), _aeg.B(), 1)
		case "\u0073\u0068":
			if len(_bbf.Params) != 1 {
				_fa.Log.Debug("\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0020\u0070\u0061r\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
				break
			}
			_adcf, _cadg := _dg.GetName(_bbf.Params[0])
			if !_cadg {
				_fa.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020n\u0061\u006d\u0065")
				break
			}
			_fed, _cadg := _ade.GetShadingByName(*_adcf)
			if !_cadg {
				_fa.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020d\u0061\u0074\u0061")
				break
			}
			_dcdb, _fcb, _acea := _ebb.processShading(_ffe, _fed)
			if _acea != nil {
				_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006e\u0067\u0020d\u0061\u0074a\u003a\u0020\u0025\u0076", _acea)
				break
			}
			if _dcdb == nil {
				_fa.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
				break
			}
			_bee, _acea := _fcb.ToFloat64Array()
			if _acea != nil {
				_fa.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073: \u0025\u0076", _acea)
				break
			}
			_ffe.DrawRectangle(_bee[0], _bee[1], _bee[2], _bee[3])
			_ffe.NewSubPath()
			_ffe.SetFillStyle(_dcdb)
			_ffe.SetStrokeStyle(_dcdb)
			_ffe.Fill()
		case "\u0044\u006f":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_cfa, _eac := _dg.GetName(_bbf.Params[0])
			if !_eac {
				return _ca
			}
			_, _afff := _ade.GetXObjectByName(*_cfa)
			switch _afff {
			case _dc.XObjectTypeImage:
				_fa.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _cfa.String())
				_gdf, _fecg := _ade.GetXObjectImageByName(*_cfa)
				if _fecg != nil {
					return _fecg
				}
				_gdfe, _fecg := _gdf.ToImage()
				if _fecg != nil {
					_fa.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _fecg)
					return nil
				}
				if _baa := _gdf.ColorSpace; _baa != nil {
					var _aabc bool
					switch _baa.(type) {
					case *_dc.PdfColorspaceSpecialIndexed:
						_aabc = true
					}
					if _aabc {
						if _bdb, _ffc := _baa.ImageToRGB(*_gdfe); _ffc != nil {
							_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
						} else {
							_gdfe = &_bdb
						}
					}
				}
				_bad := _ffe.FillPattern().ColorAt(0, 0)
				var _cbdd _d.Image
				if _gdf.Mask != nil {
					if _cbdd, _fecg = _ffef(_gdf.Mask, _bad); _fecg != nil {
						_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				} else if _gdf.SMask != nil {
					if _cbdd, _fecg = _ceg(_gdf.SMask, _bad); _fecg != nil {
						_fa.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				}
				var _fbc _d.Image
				if _eege, _ := _dg.GetBoolVal(_gdf.ImageMask); _eege {
					_fbc = _ffeb(_gdfe, _bad)
				} else {
					_fbc, _fecg = _gdfe.ToGoImage()
					if _fecg != nil {
						_fa.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _fecg)
						return nil
					}
				}
				if _cbdd != nil {
					_fbc = _cggc(_fbc, _cbdd)
				}
				_baf := _fbc.Bounds()
				_ffe.Push()
				_ffe.Scale(1.0/float64(_baf.Dx()), -1.0/float64(_baf.Dy()))
				_ffe.DrawImageAnchored(_fbc, 0, 0, 0, 1)
				_ffe.Pop()
			case _dc.XObjectTypeForm:
				_fa.Log.Debug("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073", _cfa.String())
				_gcag, _afb := _ade.GetXObjectFormByName(*_cfa)
				if _afb != nil {
					return _afb
				}
				_eab, _afb := _gcag.GetContentStream()
				if _afb != nil {
					return _afb
				}
				_fge := _gcag.Resources
				if _fge == nil {
					_fge = _ade
				}
				_ffe.Push()
				if _gcag.Matrix != nil {
					_ccb, _baff := _dg.GetArray(_gcag.Matrix)
					if !_baff {
						return _ca
					}
					_ecd, _cfeb := _dg.GetNumbersAsFloat(_ccb.Elements())
					if _cfeb != nil {
						return _cfeb
					}
					if len(_ecd) != 6 {
						return _da
					}
					_ddg := _fae.NewMatrix(_ecd[0], _ecd[1], _ecd[2], _ecd[3], _ecd[4], _ecd[5])
					_ffe.SetMatrix(_ffe.Matrix().Mult(_ddg))
				}
				if _gcag.BBox != nil {
					_fgcd, _bcg := _dg.GetArray(_gcag.BBox)
					if !_bcg {
						return _ca
					}
					_gcee, _ecbe := _dg.GetNumbersAsFloat(_fgcd.Elements())
					if _ecbe != nil {
						return _ecbe
					}
					if len(_gcee) != 4 {
						_fa.Log.Debug("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064", len(_gcee))
						return _da
					}
					_ffe.DrawRectangle(_gcee[0], _gcee[1], _gcee[2]-_gcee[0], _gcee[3]-_gcee[1])
					_ffe.SetRGBA(1, 0, 0, 1)
					_ffe.Clip()
				} else {
					_fa.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d")
				}
				_afb = _ebb.renderContentStream(_ffe, string(_eab), _fge)
				if _afb != nil {
					return _afb
				}
				_ffe.Pop()
			}
		case "\u0042\u0049":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_ada, _gff := _bbf.Params[0].(*_aa.ContentStreamInlineImage)
			if !_gff {
				return nil
			}
			_caa, _ddd := _ada.ToImage(_ade)
			if _ddd != nil {
				_fa.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ddd)
				return nil
			}
			_edeg, _ddd := _caa.ToGoImage()
			if _ddd != nil {
				_fa.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ddd)
				return nil
			}
			_fdff := _edeg.Bounds()
			_ffe.Push()
			_ffe.Scale(1.0/float64(_fdff.Dx()), -1.0/float64(_fdff.Dy()))
			_ffe.DrawImageAnchored(_edeg, 0, 0, 0, 1)
			_ffe.Pop()
		case "\u0042\u0054":
			_bgf.Reset()
		case "\u0045\u0054":
			_bgf.Reset()
		case "\u0054\u0072":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_dcee, _dabg := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _dabg != nil {
				return _dabg
			}
			_bgf.Tr = _ad.TextRenderingMode(_dcee)
		case "\u0054\u004c":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_abdc, _bafg := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _bafg != nil {
				return _bafg
			}
			_bgf.Tl = _abdc
		case "\u0054\u0063":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_gedf, _ced := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _ced != nil {
				return _ced
			}
			_fa.Log.Debug("\u0054\u0063\u003a\u0020\u0025\u0076", _gedf)
			_bgf.Tc = _gedf
		case "\u0054\u0077":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_cce, _dde := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _dde != nil {
				return _dde
			}
			_fa.Log.Debug("\u0054\u0077\u003a\u0020\u0025\u0076", _cce)
			_bgf.Tw = _cce
		case "\u0054\u007a":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_eba, _agg := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _agg != nil {
				return _agg
			}
			_bgf.Th = _eba
		case "\u0054\u0073":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_daa, _ecdf := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _ecdf != nil {
				return _ecdf
			}
			_bgf.Ts = _daa
		case "\u0054\u0064":
			if len(_bbf.Params) != 2 {
				return _da
			}
			_cfag, _egef := _dg.GetNumbersAsFloat(_bbf.Params)
			if _egef != nil {
				return _egef
			}
			_fa.Log.Debug("\u0054\u0064\u003a\u0020\u0025\u0076", _cfag)
			_bgf.ProcTd(_cfag[0], _cfag[1])
		case "\u0054\u0044":
			if len(_bbf.Params) != 2 {
				return _da
			}
			_badb, _ccbd := _dg.GetNumbersAsFloat(_bbf.Params)
			if _ccbd != nil {
				return _ccbd
			}
			_fa.Log.Debug("\u0054\u0044\u003a\u0020\u0025\u0076", _badb)
			_bgf.ProcTD(_badb[0], _badb[1])
		case "\u0054\u002a":
			_bgf.ProcTStar()
		case "\u0054\u006d":
			if len(_bbf.Params) != 6 {
				return _da
			}
			_ggg, _cdg := _dg.GetNumbersAsFloat(_bbf.Params)
			if _cdg != nil {
				return _cdg
			}
			_fa.Log.Debug("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076", _ggg)
			_bgf.ProcTm(_ggg[0], _ggg[1], _ggg[2], _ggg[3], _ggg[4], _ggg[5])
		case "\u0027":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_gge, _aada := _dg.GetStringBytes(_bbf.Params[0])
			if !_aada {
				return _ca
			}
			_fa.Log.Debug("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gge))
			_bgf.ProcQ(_gge, _ffe)
		case "\u0022":
			if len(_bbf.Params) != 3 {
				return _da
			}
			_fba, _bfe := _dg.GetNumberAsFloat(_bbf.Params[0])
			if _bfe != nil {
				return _bfe
			}
			_fdfa, _bfe := _dg.GetNumberAsFloat(_bbf.Params[1])
			if _bfe != nil {
				return _bfe
			}
			_fgbb, _deb := _dg.GetStringBytes(_bbf.Params[2])
			if !_deb {
				return _ca
			}
			_bgf.ProcDQ(_fgbb, _fba, _fdfa, _ffe)
		case "\u0054\u006a":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_abe, _cbda := _dg.GetStringBytes(_bbf.Params[0])
			if !_cbda {
				return _ca
			}
			_fa.Log.Debug("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060", string(_abe))
			_bgf.ProcTj(_abe, _ffe)
		case "\u0054\u004a":
			if len(_bbf.Params) != 1 {
				return _da
			}
			_afa, _bca := _dg.GetArray(_bbf.Params[0])
			if !_bca {
				_fa.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _afa)
				return _ca
			}
			_fa.Log.Debug("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076", _afa)
			for _, _fgcc := range _afa.Elements() {
				switch _caag := _fgcc.(type) {
				case *_dg.PdfObjectString:
					if _caag != nil {
						_bgf.ProcTj(_caag.Bytes(), _ffe)
					}
				case *_dg.PdfObjectFloat, *_dg.PdfObjectInteger:
					_ddac, _fgf := _dg.GetNumberAsFloat(_caag)
					if _fgf == nil {
						_bgf.Translate(-_ddac*0.001*_bgf.Tf.Size*_bgf.Th/100.0, 0)
					}
				}
			}
		case "\u0054\u0066":
			if len(_bbf.Params) != 2 {
				return _da
			}
			_fa.Log.Debug("\u0025\u0023\u0076", _bbf.Params)
			_fdfb, _adda := _dg.GetName(_bbf.Params[0])
			if !_adda || _fdfb == nil {
				_fa.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bbf.Params[0])
				return _ca
			}
			_fa.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _fdfb.String())
			_cgbf, _efg := _dg.GetNumberAsFloat(_bbf.Params[1])
			if _efg != nil {
				_fa.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bbf.Params[1])
				return _ca
			}
			_fa.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076", _cgbf)
			_bcc, _bdbc := _ade.GetFontByName(*_fdfb)
			if !_bdbc {
				_fa.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _fdfb.String())
				return _e.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
			}
			_fa.Log.Debug("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054", _bcc)
			_aedg, _adda := _dg.GetDict(_bcc)
			if !_adda {
				_fa.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074")
				return _ca
			}
			_ccbe, _efg := _dc.NewPdfFontFromPdfObject(_aedg)
			if _efg != nil {
				_fa.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
				return _efg
			}
			_faeg := _ccbe.BaseFont()
			if _faeg == "" {
				_faeg = _fdfb.String()
			}
			_bdg, _adda := _fbf[_faeg]
			if !_adda {
				_bdg, _efg = _ad.NewTextFont(_ccbe, _cgbf)
				if _efg != nil {
					_fa.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _efg)
				}
			}
			if _bdg == nil {
				if len(_faeg) > 7 && _faeg[6] == '+' {
					_faeg = _faeg[7:]
				}
				_ded := []string{_faeg, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
				for _, _dag := range _ded {
					_fa.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _dag)
					if _bdg, _adda = _fbf[_dag]; _adda {
						break
					}
					_ggff := _de.Match(_dag)
					if _ggff == nil {
						_fa.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _dag)
						continue
					}
					_bdg, _efg = _ad.NewTextFontFromPath(_ggff.Filename, _cgbf)
					if _efg != nil {
						_fa.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _ggff.Filename)
						continue
					}
					_fa.Log.Debug("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029", _faeg, _ggff.Name, _ggff.Filename)
					_fbf[_dag] = _bdg
					break
				}
			}
			if _bdg == nil {
				_fa.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074")
				return _e.New("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074")
			}
			_bgf.ProcTf(_bdg.WithSize(_cgbf, _ccbe))
		case "\u0042\u004d\u0043", "\u0042\u0044\u0043":
		case "\u0045\u004d\u0043":
		default:
			_fa.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073", _bbf.Operand)
		}
		_ged = _bbf
		return nil
	})
	_aabb = _ecb.Process(_bac)
	if _aabb != nil {
		return _aabb
	}
	return nil
}
func _dad(_fecf _ad.Gradient, _aca *_dc.PdfFunctionType2, _dbg _dc.PdfColorspace, _adee float64, _edce bool) (_ad.Gradient, error) {
	switch _dbg.(type) {
	case *_dc.PdfColorspaceDeviceRGB:
		if len(_aca.C0) != 3 || len(_aca.C1) != 3 {
			return nil, _e.New("\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0052\u0047\u0042\u0020\u0063o\u006co\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
		}
		_cgcb := _aca.C0
		_edee := _aca.C1
		if _edce {
			_fecf.AddColorStop(0.0, _fg.RGBA{R: uint8(_cgcb[0] * 255), G: uint8(_cgcb[1] * 255), B: uint8(_cgcb[2] * 255), A: 255})
		}
		_fecf.AddColorStop(_adee, _fg.RGBA{R: uint8(_edee[0] * 255), G: uint8(_edee[1] * 255), B: uint8(_edee[2] * 255), A: 255})
	case *_dc.PdfColorspaceDeviceCMYK:
		if len(_aca.C0) != 4 || len(_aca.C1) != 4 {
			return nil, _e.New("\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072 \u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
		}
		_fdfg := _aca.C0
		_ecfg := _aca.C1
		if _edce {
			_fecf.AddColorStop(0.0, _fg.CMYK{C: uint8(_fdfg[0] * 255), M: uint8(_fdfg[1] * 255), Y: uint8(_fdfg[2] * 255), K: uint8(_fdfg[3] * 255)})
		}
		_fecf.AddColorStop(_adee, _fg.CMYK{C: uint8(_ecfg[0] * 255), M: uint8(_ecfg[1] * 255), Y: uint8(_ecfg[2] * 255), K: uint8(_ecfg[3] * 255)})
	default:
		return nil, _bbe.Errorf("u\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072 \u0073\u0070\u0061c\u0065:\u0020\u0025\u0073", _dbg.String())
	}
	return _fecf, nil
}

// NewImageDevice returns a new image device.
func NewImageDevice() *ImageDevice {
	const _dga = "r\u0065\u006e\u0064\u0065r.\u004ee\u0077\u0049\u006d\u0061\u0067e\u0044\u0065\u0076\u0069\u0063\u0065"
	_bc.TrackUse(_dga)
	return &ImageDevice{}
}
func _gfee(_fab *_dc.Image, _agb _fg.Color) _d.Image {
	_ceb, _aga := int(_fab.Width), int(_fab.Height)
	_bed := _d.NewRGBA(_d.Rect(0, 0, _ceb, _aga))
	for _afbd := 0; _afbd < _aga; _afbd++ {
		for _cgef := 0; _cgef < _ceb; _cgef++ {
			_eca, _ffb := _fab.ColorAt(_cgef, _afbd)
			if _ffb != nil {
				_fa.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _cgef, _afbd)
				continue
			}
			_gfeb, _cag, _gdc, _ := _eca.RGBA()
			var _ffcc _fg.Color
			if _gfeb+_cag+_gdc == 0 {
				_ffcc = _fg.Transparent
			} else {
				_ffcc = _agb
			}
			_bed.Set(_cgef, _afbd, _ffcc)
		}
	}
	return _bed
}

// RenderWithOpts converts the specified PDF page into an image, optionally flattens annotations and returns the result.
func (_ef *ImageDevice) RenderWithOpts(page *_dc.PdfPage, skipFlattening bool) (_d.Image, error) {
	_c, _ae := page.GetMediaBox()
	if _ae != nil {
		return nil, _ae
	}
	_c.Normalize()
	_bg := page.CropBox
	var _dgg, _ab float64
	if _bg != nil {
		_bg.Normalize()
		_dgg, _ab = _bg.Width(), _bg.Height()
	}
	_fd := page.Rotate
	_cg, _ed, _efc, _cd := _c.Llx, _c.Lly, _c.Width(), _c.Height()
	_bfg := _fae.IdentityMatrix()
	if _fd != nil && *_fd%360 != 0 && *_fd%90 == 0 {
		_gd := -float64(*_fd)
		_eg := _gbe(_efc, _cd, _gd)
		_bfg = _bfg.Translate((_eg.Width-_efc)/2+_efc/2, (_eg.Height-_cd)/2+_cd/2).Rotate(_gd*_ge.Pi/180).Translate(-_efc/2, -_cd/2)
		_efc, _cd = _eg.Width, _eg.Height
		if _bg != nil {
			_dcg := _gbe(_dgg, _ab, _gd)
			_dgg, _ab = _dcg.Width, _dcg.Height
		}
	}
	if _cg != 0 || _ed != 0 {
		_bfg = _bfg.Translate(-_cg, -_ed)
	}
	_ef._aab = 1.0
	if _ef.OutputWidth != 0 {
		_ba := _efc
		if _bg != nil {
			_ba = _dgg
		}
		_ef._aab = float64(_ef.OutputWidth) / _ba
		_efc, _cd, _dgg, _ab = _efc*_ef._aab, _cd*_ef._aab, _dgg*_ef._aab, _ab*_ef._aab
		_bfg = _fae.ScaleMatrix(_ef._aab, _ef._aab).Mult(_bfg)
	}
	_df := _fe.NewContext(int(_efc), int(_cd))
	if _cb := _ef.renderPage(_df, page, _bfg, skipFlattening); _cb != nil {
		return nil, _cb
	}
	_af := _df.Image()
	if _bg != nil {
		_ce, _aff := (_bg.Llx-_cg)*_ef._aab, (_bg.Lly-_ed)*_ef._aab
		_faf := _d.Rect(0, 0, int(_dgg), int(_ab))
		_gdd := _d.Pt(int(_ce), int(_cd-_aff-_ab))
		_gc := _d.NewRGBA(_faf)
		_ga.Draw(_gc, _faf, _af, _gdd, _ga.Src)
		_af = _gc
	}
	return _af, nil
}
