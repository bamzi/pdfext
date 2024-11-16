package transform

import (
	_a "fmt"
	_f "math"

	_ad "github.com/bamzi/pdfext/common"
)

func (_c Matrix) Round(precision float64) Matrix {
	for _adb := range _c {
		_c[_adb] = _f.Round(_c[_adb]/precision) * precision
	}
	return _c
}
func ShearMatrix(x, y float64) Matrix { return NewMatrix(1, y, x, 1, 0, 0) }
func NewMatrixFromTransforms(xScale, yScale, theta, tx, ty float64) Matrix {
	return IdentityMatrix().Scale(xScale, yScale).Rotate(theta).Translate(tx, ty)
}
func (_ef *Matrix) Set(a, b, c, d, tx, ty float64) {
	_ef[0], _ef[1] = a, b
	_ef[3], _ef[4] = c, d
	_ef[6], _ef[7] = tx, ty
	_ef.clampRange()
}
func (_ga *Matrix) Concat(b Matrix) {
	*_ga = Matrix{b[0]*_ga[0] + b[1]*_ga[3], b[0]*_ga[1] + b[1]*_ga[4], 0, b[3]*_ga[0] + b[4]*_ga[3], b[3]*_ga[1] + b[4]*_ga[4], 0, b[6]*_ga[0] + b[7]*_ga[3] + _ga[6], b[6]*_ga[1] + b[7]*_ga[4] + _ga[7], 1}
	_ga.clampRange()
}
func (_d Matrix) String() string {
	_de, _gb, _cg, _fe, _ac, _ec := _d[0], _d[1], _d[3], _d[4], _d[6], _d[7]
	return _a.Sprintf("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d", _de, _gb, _cg, _fe, _ac, _ec)
}
func (_ff *Matrix) Clone() Matrix                      { return NewMatrix(_ff[0], _ff[1], _ff[3], _ff[4], _ff[6], _ff[7]) }
func (_da Matrix) Singular() bool                      { return _f.Abs(_da[0]*_da[4]-_da[1]*_da[3]) < _ag }
func (_fd Matrix) Scale(xScale, yScale float64) Matrix { return _fd.Mult(ScaleMatrix(xScale, yScale)) }
func (_df Matrix) Translation() (float64, float64)     { return _df[6], _df[7] }
func (_g Matrix) Identity() bool {
	return _g[0] == 1 && _g[1] == 0 && _g[2] == 0 && _g[3] == 0 && _g[4] == 1 && _g[5] == 0 && _g[6] == 0 && _g[7] == 0 && _g[8] == 1
}
func (_fee Matrix) Rotate(theta float64) Matrix { return _fee.Mult(RotationMatrix(theta)) }
func RotationMatrix(angle float64) Matrix {
	_fb := _f.Cos(angle)
	_aa := _f.Sin(angle)
	return NewMatrix(_fb, _aa, -_aa, _fb, 0, 0)
}
func (_fc Matrix) ScalingFactorY() float64  { return _f.Hypot(_fc[3], _fc[4]) }
func (_eca Point) Distance(b Point) float64 { return _f.Hypot(_eca.X-b.X, _eca.Y-b.Y) }
func (_add Matrix) Inverse() (Matrix, bool) {
	_gae, _gc := _add[0], _add[1]
	_gaa, _fg := _add[3], _add[4]
	_aac, _gd := _add[6], _add[7]
	_ea := _gae*_fg - _gc*_gaa
	if _f.Abs(_ea) < _bfcd {
		return Matrix{}, false
	}
	_bb, _bf := _fg/_ea, -_gc/_ea
	_dbd, _ab := -_gaa/_ea, _gae/_ea
	_bfc := -(_bb*_aac + _dbd*_gd)
	_edd := -(_bf*_aac + _ab*_gd)
	return NewMatrix(_bb, _bf, _dbd, _ab, _bfc, _edd), true
}
func (_dc Matrix) Transform(x, y float64) (float64, float64) {
	_db := x*_dc[0] + y*_dc[3] + _dc[6]
	_ce := x*_dc[1] + y*_dc[4] + _dc[7]
	return _db, _ce
}
func (_bfca Point) String() string {
	return _a.Sprintf("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029", _bfca.X, _bfca.Y)
}

type Matrix [9]float64

const _ag = 1e-10

func (_fbda Matrix) Unrealistic() bool {
	_dcc, _gaef, _gab, _ecg := _f.Abs(_fbda[0]), _f.Abs(_fbda[1]), _f.Abs(_fbda[3]), _f.Abs(_fbda[4])
	_ba := _dcc > _ace && _ecg > _ace
	_bg := _gaef > _ace && _gab > _ace
	return !(_ba || _bg)
}
func (_be Point) Displace(delta Point) Point { return Point{_be.X + delta.X, _be.Y + delta.Y} }
func ScaleMatrix(x, y float64) Matrix        { return NewMatrix(x, 0, 0, y, 0, 0) }

type Point struct {
	X float64
	Y float64
}

func (_bd Point) Interpolate(b Point, t float64) Point {
	return Point{X: (1-t)*_bd.X + t*b.X, Y: (1-t)*_bd.Y + t*b.Y}
}
func (_cb *Point) Set(x, y float64) { _cb.X, _cb.Y = x, y }
func (_ee *Point) Transform(a, b, c, d, tx, ty float64) {
	_cd := NewMatrix(a, b, c, d, tx, ty)
	_ee.transformByMatrix(_cd)
}

const _af = 1e9

func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

const _ace = 1e-6

func TranslationMatrix(tx, ty float64) Matrix { return NewMatrix(1, 0, 0, 1, tx, ty) }
func (_eag Point) Rotate(theta float64) Point {
	_febb := _f.Hypot(_eag.X, _eag.Y)
	_fcf := _f.Atan2(_eag.Y, _eag.X)
	_ded, _cgg := _f.Sincos(_fcf + theta/180.0*_f.Pi)
	return Point{_febb * _cgg, _febb * _ded}
}
func IdentityMatrix() Matrix                      { return NewMatrix(1, 0, 0, 1, 0, 0) }
func (_fda *Point) transformByMatrix(_gdg Matrix) { _fda.X, _fda.Y = _gdg.Transform(_fda.X, _fda.Y) }
func (_dff *Matrix) clampRange() {
	for _fa, _gaaa := range _dff {
		if _gaaa > _af {
			_ad.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gaaa, _af)
			_dff[_fa] = _af
		} else if _gaaa < -_af {
			_ad.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gaaa, -_af)
			_dff[_fa] = -_af
		}
	}
}
func (_dee *Matrix) Shear(x, y float64) { _dee.Concat(ShearMatrix(x, y)) }
func (_feb Matrix) Angle() float64 {
	_bc := _f.Atan2(-_feb[1], _feb[0])
	if _bc < 0.0 {
		_bc += 2 * _f.Pi
	}
	return _bc / _f.Pi * 180.0
}
func (_fbd Matrix) Mult(b Matrix) Matrix {
	_fbd.Concat(b)
	return _fbd
}

const _bfcd = 1.0e-6

func (_deg Matrix) Translate(tx, ty float64) Matrix { return _deg.Mult(TranslationMatrix(tx, ty)) }
func (_b Matrix) ScalingFactorX() float64           { return _f.Hypot(_b[0], _b[1]) }
func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	_ed := Matrix{a, b, 0, c, d, 0, tx, ty, 1}
	_ed.clampRange()
	return _ed
}
