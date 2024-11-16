package unichart

import (
	_e "bytes"
	_ec "fmt"
	_bd "image/color"
	_b "io"
	_a "math"

	_g "github.com/bamzi/pdfext/common"
	_bc "github.com/bamzi/pdfext/contentstream"
	_gf "github.com/bamzi/pdfext/contentstream/draw"
	_ac "github.com/bamzi/pdfext/core"
	_df "github.com/bamzi/pdfext/model"
	_c "github.com/unidoc/unichart/render"
)

func (_acg *Renderer) Save(w _b.Writer) error {
	if w == nil {
		return nil
	}
	_, _ag := _b.Copy(w, _e.NewBuffer(_acg._ba.Bytes()))
	return _ag
}
func (_cc *Renderer) SetFillColor(color _bd.Color) {
	_cc._cd = color
	_de, _ge, _aa, _ := _fcc(color)
	_cc._ba.Add_rg(_de, _ge, _aa)
}
func (_dda *Renderer) MeasureText(text string) _c.Box {
	_ca := _dda._cf
	_fa, _fbg := _dda._gg.GetFontDescriptor()
	if _fbg != nil {
		_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	} else {
		_fgb, _fae := _fa.GetCapHeight()
		if _fae != nil {
			_g.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _fae)
		} else {
			_ca = _fgb / 1000.0 * _dda._cf
		}
	}
	var (
		_gea = 0.0
		_db  = _dda.wrapText(text)
	)
	for _, _dcd := range _db {
		if _afa := _dda.getTextWidth(_dcd); _afa > _gea {
			_gea = _afa
		}
	}
	_acf := _c.NewBox(0, 0, int(_gea), int(_ca))
	if _cce := _dda._gb; _cce != 0 {
		_acf = _acf.Corners().Rotate(_cce).Box()
	}
	return _acf
}
func (_dcfc *Renderer) Fill()            { _dcfc._ba.Add_f() }
func (_fd *Renderer) ClearTextRotation() { _fd._gb = 0 }
func (_ddc *Renderer) MoveTo(x, y int)   { _ddc._ba.Add_m(float64(x), float64(y)) }
func (_dc *Renderer) SetStrokeColor(color _bd.Color) {
	_dc._da = color
	_eaf, _ga, _gca, _ := _fcc(color)
	_dc._ba.Add_RG(_eaf, _ga, _gca)
}
func (_fe *Renderer) SetFont(font _c.Font) {
	_abb, _bcb := font.(*_df.PdfFont)
	if !_bcb {
		_g.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
		return
	}
	_efd, _bcb := _fe._be[_abb]
	if !_bcb {
		_efd = _fdg("\u0046\u006f\u006e\u0074", 1, _fe._bcg.HasFontByName)
		if _ed := _fe._bcg.SetFontByName(_efd, _abb.ToPdfObject()); _ed != nil {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073", _abb)
		}
		_fe._be[_abb] = _efd
	}
	_fe._ba.Add_Tf(_efd, _fe._cf)
	_fe._gg = _abb
}
func (_bb *Renderer) SetDPI(dpi float64)               { _bb._f = dpi }
func (_dcf *Renderer) SetStrokeWidth(width float64)    { _dcf._ee = width; _dcf._ba.Add_w(width) }
func (_bfd *Renderer) SetTextRotation(radians float64) { _bfd._gb = _dddb(-radians) }
func (_ef *Renderer) ResetStyle() {
	_ef.SetFillColor(_bd.Black)
	_ef.SetStrokeColor(_bd.Transparent)
	_ef.SetStrokeWidth(0)
	_ef.SetFont(_df.DefaultFont())
	_ef.SetFontColor(_bd.Black)
	_ef.SetFontSize(12)
	_ef.SetTextRotation(0)
}
func (_bgd *Renderer) QuadCurveTo(cx, cy, x, y int) {
	_bgd._ba.Add_v(float64(x), float64(y), float64(cx), float64(cy))
}
func (_gga *Renderer) GetDPI() float64 { return _gga._f }
func (_fbe *Renderer) Text(text string, x, y int) {
	_fbe._ba.Add_q()
	_fbe.SetFont(_fbe._gg)
	_fca, _edg, _bab, _ := _fcc(_fbe._gc)
	_fbe._ba.Add_rg(_fca, _edg, _bab)
	_fbe._ba.Translate(float64(x), float64(y)).Scale(1, -1)
	if _gaga := _fbe._gb; _gaga != 0 {
		_fbe._ba.RotateDeg(_gaga)
	}
	_fbe._ba.Add_BT().Add_TL(_fbe._cf)
	var (
		_gdc = _fbe._gg.Encoder()
		_cda = _fbe.wrapText(text)
		_gec = len(_cda)
	)
	for _cba, _ddd := range _cda {
		_fbe._ba.Add_TJ(_ac.MakeStringFromBytes(_gdc.Encode(_ddd)))
		if _cba != _gec-1 {
			_fbe._ba.Add_Tstar()
		}
	}
	_fbe._ba.Add_ET()
	_fbe._ba.Add_Q()
}
func (_abf *Renderer) wrapText(_eag string) []string {
	var (
		_afaa []string
		_eb   []rune
	)
	for _, _eff := range _eag {
		if _eff == '\n' {
			_afaa = append(_afaa, string(_eb))
			_eb = []rune{}
			continue
		}
		_eb = append(_eb, _eff)
	}
	if len(_eb) > 0 {
		_afaa = append(_afaa, string(_eb))
	}
	return _afaa
}
func (_bcd *Renderer) Circle(radius float64, x, y int) {
	_gd := radius
	if _fg := _bcd._ee; _fg != 0 {
		_gd -= _fg / 2
	}
	_dgd := _gd * 0.551784
	_ad := _gf.CubicBezierPath{Curves: []_gf.CubicBezierCurve{_gf.NewCubicBezierCurve(-_gd, 0, -_gd, _dgd, -_dgd, _gd, 0, _gd), _gf.NewCubicBezierCurve(0, _gd, _dgd, _gd, _gd, _dgd, _gd, 0), _gf.NewCubicBezierCurve(_gd, 0, _gd, -_dgd, _dgd, -_gd, 0, -_gd), _gf.NewCubicBezierCurve(0, -_gd, -_dgd, -_gd, -_gd, -_dgd, -_gd, 0)}}
	if _cff := _bcd._ee; _cff != 0 {
		_ad = _ad.Offset(_cff/2, _cff/2)
	}
	_ad = _ad.Offset(float64(x), float64(y))
	_gf.DrawBezierPathWithCreator(_ad, _bcd._ba)
}
func (_bg *Renderer) LineTo(x, y int)               { _bg._ba.Add_l(float64(x), float64(y)) }
func (_cgd *Renderer) SetFontSize(size float64)     { _cgd._cf = size }
func (_ecb *Renderer) Close()                       { _ecb._ba.Add_h() }
func (_eeg *Renderer) SetFontColor(color _bd.Color) { _eeg._gc = color }
func _egf(_bgcb _bd.Color) (uint8, uint8, uint8, uint8) {
	_ff, _eaff, _baac, _ega := _bgcb.RGBA()
	return uint8(_ff >> 8), uint8(_eaff >> 8), uint8(_baac >> 8), uint8(_ega >> 8)
}
func (_cb *Renderer) SetStrokeDashArray(dashArray []float64) {
	_deg := make([]int64, len(dashArray))
	for _dca, _dd := range dashArray {
		_deg[_dca] = int64(_dd)
	}
	_cb._ba.Add_d(_deg, 0)
}
func (_fb *Renderer) ArcTo(cx, cy int, rx, ry, startAngle, deltaAngle float64) {
	startAngle = _dddb(2.0*_a.Pi - startAngle)
	deltaAngle = _dddb(-deltaAngle)
	_eee, _bde := deltaAngle, 1
	if _a.Abs(deltaAngle) > 90.0 {
		_bde = int(_a.Ceil(_a.Abs(deltaAngle) / 90.0))
		_eee = deltaAngle / float64(_bde)
	}
	var (
		_gff = _dddg(_eee / 2)
		_gbc = _a.Abs(4.0 / 3.0 * (1.0 - _a.Cos(_gff)) / _a.Sin(_gff))
		_eca = float64(cx)
		_dg  = float64(cy)
	)
	for _bfc := 0; _bfc < _bde; _bfc++ {
		_cfg := _dddg(startAngle + float64(_bfc)*_eee)
		_gbe := _dddg(startAngle + float64(_bfc+1)*_eee)
		_eac := _a.Cos(_cfg)
		_cg := _a.Cos(_gbe)
		_aba := _a.Sin(_cfg)
		_fc := _a.Sin(_gbe)
		var _eg []float64
		if _eee > 0 {
			_eg = []float64{_eca + rx*_eac, _dg - ry*_aba, _eca + rx*(_eac-_gbc*_aba), _dg - ry*(_aba+_gbc*_eac), _eca + rx*(_cg+_gbc*_fc), _dg - ry*(_fc-_gbc*_cg), _eca + rx*_cg, _dg - ry*_fc}
		} else {
			_eg = []float64{_eca + rx*_eac, _dg - ry*_aba, _eca + rx*(_eac+_gbc*_aba), _dg - ry*(_aba-_gbc*_eac), _eca + rx*(_cg-_gbc*_fc), _dg - ry*(_fc+_gbc*_cg), _eca + rx*_cg, _dg - ry*_fc}
		}
		if _bfc == 0 {
			_fb._ba.Add_l(_eg[0], _eg[1])
		}
		_fb._ba.Add_c(_eg[2], _eg[3], _eg[4], _eg[5], _eg[6], _eg[7])
	}
}
func NewRenderer(cc *_bc.ContentCreator, res *_df.PdfPageResources) func(int, int) (_c.Renderer, error) {
	return func(_bf, _gbg int) (_c.Renderer, error) {
		_ab := &Renderer{_bdg: _bf, _af: _gbg, _f: 72, _ba: cc, _bcg: res, _be: map[*_df.PdfFont]_ac.PdfObjectName{}}
		_ab.ResetStyle()
		return _ab, nil
	}
}
func _dddb(_ageb float64) float64 { return _ageb * 180 / _a.Pi }

type Renderer struct {
	_bdg int
	_af  int
	_f   float64
	_ba  *_bc.ContentCreator
	_bcg *_df.PdfPageResources
	_cd  _bd.Color
	_da  _bd.Color
	_ee  float64
	_gg  *_df.PdfFont
	_cf  float64
	_gc  _bd.Color
	_gb  float64
	_be  map[*_df.PdfFont]_ac.PdfObjectName
}

func (_ea *Renderer) SetClassName(name string) {}
func _fdg(_agb string, _faea int, _cbc func(_ac.PdfObjectName) bool) _ac.PdfObjectName {
	_bda := _ac.PdfObjectName(_ec.Sprintf("\u0025\u0073\u0025\u0064", _agb, _faea))
	for _age := _faea; _cbc(_bda); {
		_age++
		_bda = _ac.PdfObjectName(_ec.Sprintf("\u0025\u0073\u0025\u0064", _agb, _age))
	}
	return _bda
}
func (_cgf *Renderer) getTextWidth(_eab string) float64 {
	var _gfb float64
	for _, _gece := range _eab {
		_eacg, _fag := _cgf._gg.GetRuneMetrics(_gece)
		if !_fag {
			_g.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _gece)
		}
		_gfb += _eacg.Wx
	}
	return _cgf._cf * _gfb / 1000.0
}
func _dddg(_add float64) float64   { return _add * _a.Pi / 180.0 }
func (_fba *Renderer) FillStroke() { _fba._ba.Add_B() }
func (_aca *Renderer) Stroke()     { _aca._ba.Add_S() }
func _fcc(_gfc _bd.Color) (float64, float64, float64, float64) {
	_cde, _cfd, _gcg, _ebb := _egf(_gfc)
	return float64(_cde) / 255, float64(_cfd) / 255, float64(_gcg) / 255, float64(_ebb) / 255
}
