package imagerender

import (
	_bd "errors"
	_ffb "fmt"
	_e "image"
	_a "image/color"
	_d "image/draw"
	_b "math"
	_g "sort"
	_ff "strings"

	_bb "github.com/bamzi/pdfext/common"
	_c "github.com/bamzi/pdfext/internal/transform"
	_gc "github.com/bamzi/pdfext/render/internal/context"
	_af "github.com/unidoc/freetype/raster"
	_ed "golang.org/x/image/draw"
	_eg "golang.org/x/image/font"
	_ab "golang.org/x/image/math/f64"
	_ec "golang.org/x/image/math/fixed"
)

func (_ged *Context) SetLineCap(lineCap _gc.LineCap) { _ged._caf = lineCap }
func (_ccg *Context) DrawPoint(x, y, r float64) {
	_ccg.Push()
	_dffab, _gdc := _ccg.Transform(x, y)
	_ccg.Identity()
	_ccg.DrawCircle(_dffab, _gdc, r)
	_ccg.Pop()
}
func (_eadb *Context) SetDash(dashes ...float64) { _eadb._dec = dashes }
func (_fca *linearGradient) AddColorStop(offset float64, color _a.Color) {
	_fca._eabb = append(_fca._eabb, stop{_bcdba: offset, _bcc: color})
	_g.Sort(_fca._eabb)
}
func _dba(_dcg _a.Color) _gc.Pattern { return &solidPattern{_dafb: _dcg} }
func (_bbg stops) Swap(i, j int)     { _bbg[i], _bbg[j] = _bbg[j], _bbg[i] }
func (_dceg *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !_dceg._aga {
		_dceg.MoveTo(x1, y1)
	}
	_cdf, _dca := _dceg._afdd.X, _dceg._afdd.Y
	x1, y1 = _dceg.Transform(x1, y1)
	x2, y2 = _dceg.Transform(x2, y2)
	x3, y3 = _dceg.Transform(x3, y3)
	_dbf := _edg(_cdf, _dca, x1, y1, x2, y2, x3, y3)
	_gab := _aba(_dceg._afdd)
	for _, _bcab := range _dbf[1:] {
		_dab := _aba(_bcab)
		if _dab == _gab {
			continue
		}
		_gab = _dab
		_dceg._dce.Add1(_dab)
		_dceg._bca.Add1(_dab)
		_dceg._afdd = _bcab
	}
}
func (_dgaf stops) Len() int { return len(_dgaf) }

var (
	_bcb = _dba(_a.White)
	_dfa = _dba(_a.Black)
)

func _fdc(_gcef float64, _cab stops) _a.Color {
	if _gcef <= 0.0 || len(_cab) == 1 {
		return _cab[0]._bcc
	}
	_fbee := _cab[len(_cab)-1]
	if _gcef >= _fbee._bcdba {
		return _fbee._bcc
	}
	for _add, _dag := range _cab[1:] {
		if _gcef < _dag._bcdba {
			_gcef = (_gcef - _cab[_add]._bcdba) / (_dag._bcdba - _cab[_add]._bcdba)
			return _edcc(_cab[_add]._bcc, _dag._bcc, _gcef)
		}
	}
	return _fbee._bcc
}
func (_bbbf *Context) Clear() {
	_gede := _e.NewUniform(_bbbf._aeg)
	_ed.Draw(_bbbf._edgd, _bbbf._edgd.Bounds(), _gede, _e.Point{}, _ed.Src)
}
func (_ga *Context) LineWidth() float64 { return _ga._fcc }
func (_dga *Context) SetRGBA(r, g, b, a float64) {
	_, _, _, _fgc := _dga._aeg.RGBA()
	if _fgc > 0 && _fgc != 65535 && a == 1 {
		a = float64(_fgc) / 65535
	}
	_dga._aeg = _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_dga.setFillAndStrokeColor(_dga._aeg)
}
func (_fcb *Context) NewSubPath() {
	if _fcb._aga {
		_fcb._bca.Add1(_aba(_fcb._eeb))
	}
	_fcb._aga = false
}
func NewContextForRGBA(im *_e.RGBA) *Context {
	_gf := im.Bounds().Size().X
	_dad := im.Bounds().Size().Y
	return &Context{_gbd: _gf, _ebe: _dad, _ge: _af.NewRasterizer(_gf, _dad), _edgd: im, _aeg: _a.Transparent, _bcd: _bcb, _faf: _dfa, _fcc: 1, _ede: _gc.FillRuleWinding, _feg: _c.IdentityMatrix(), _cdg: _gc.NewTextState()}
}
func (_dfdd *Context) ClosePath() {
	if _dfdd._aga {
		_fbab := _aba(_dfdd._eeb)
		_dfdd._dce.Add1(_fbab)
		_dfdd._bca.Add1(_fbab)
		_dfdd._afdd = _dfdd._eeb
	}
}
func (_aac *Context) SetRGB(r, g, b float64) { _aac.SetRGBA(r, g, b, 1) }
func (_cae *Context) SetRGB255(r, g, b int)  { _cae.SetRGBA255(r, g, b, 255) }

type circle struct{ _fbagb, _fdedb, _cegc float64 }

func (_cea *Context) StrokePreserve() {
	var _fad _af.Painter
	if _cea._dff == nil {
		if _dae, _cbb := _cea._faf.(*solidPattern); _cbb {
			_ggec := _af.NewRGBAPainter(_cea._edgd)
			_ggec.SetColor(_dae._dafb)
			_fad = _ggec
		}
	}
	if _fad == nil {
		_fad = _dace(_cea._edgd, _cea._dff, _cea._faf)
	}
	_cea.stroke(_fad)
}
func (_beb *Context) ShearAbout(sx, sy, x, y float64) {
	_beb.Translate(x, y)
	_beb.Shear(sx, sy)
	_beb.Translate(-x, -y)
}
func (_deb *Context) SetColor(c _a.Color)             { _deb.setFillAndStrokeColor(c) }
func (_fbba *solidPattern) ColorAt(x, y int) _a.Color { return _fbba._dafb }
func _cb(_bf, _fg, _db, _ac, _bc, _ca, _cd float64) (_bbb, _dc float64) {
	_de := 1 - _cd
	_afc := _de * _de
	_egf := 2 * _de * _cd
	_fe := _cd * _cd
	_bbb = _afc*_bf + _egf*_db + _fe*_bc
	_dc = _afc*_fg + _egf*_ac + _fe*_ca
	return
}
func (_cgfe *Context) DrawImage(im _e.Image, x, y int) { _cgfe.DrawImageAnchored(im, x, y, 0, 0) }
func (_ecb *Context) ClearPath()                       { _ecb._dce.Clear(); _ecb._bca.Clear(); _ecb._aga = false }
func (_abf *Context) Matrix() _c.Matrix                { return _abf._feg }
func (_gac *Context) ResetClip()                       { _gac._dff = nil }
func (_efb *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !_efb._aga {
		_efb.MoveTo(x1, y1)
	}
	x1, y1 = _efb.Transform(x1, y1)
	x2, y2 = _efb.Transform(x2, y2)
	_bcdb := _c.NewPoint(x1, y1)
	_cge := _c.NewPoint(x2, y2)
	_gceg := _aba(_bcdb)
	_cegf := _aba(_cge)
	_efb._dce.Add2(_gceg, _cegf)
	_efb._bca.Add2(_gceg, _cegf)
	_efb._afdd = _cge
}
func _baeg(_cfbf _af.Path, _fac []float64, _fgd float64) _af.Path {
	return _ffad(_aegg(_afdb(_cfbf), _fac, _fgd))
}
func (_bde *Context) SetPixel(x, y int) { _bde._edgd.Set(x, y, _bde._aeg) }
func (_fge *Context) Push()             { _gdg := *_fge; _fge._gce = append(_fge._gce, &_gdg) }
func (_fb *Context) SetFillStyle(pattern _gc.Pattern) {
	if _edf, _dceb := pattern.(*solidPattern); _dceb {
		_fb._aeg = _edf._dafb
	}
	_fb._bcd = pattern
}
func _cgca(_abba string) (_gdce, _ccag, _ddae, _fcbg int) {
	_abba = _ff.TrimPrefix(_abba, "\u0023")
	_fcbg = 255
	if len(_abba) == 3 {
		_bbed := "\u00251\u0078\u0025\u0031\u0078\u0025\u0031x"
		_ffb.Sscanf(_abba, _bbed, &_gdce, &_ccag, &_ddae)
		_gdce |= _gdce << 4
		_ccag |= _ccag << 4
		_ddae |= _ddae << 4
	}
	if len(_abba) == 6 {
		_ecf := "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ffb.Sscanf(_abba, _ecf, &_gdce, &_ccag, &_ddae)
	}
	if len(_abba) == 8 {
		_adb := "\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ffb.Sscanf(_abba, _adb, &_gdce, &_ccag, &_ddae, &_fcbg)
	}
	return
}
func _cf(_def, _ba, _fa, _gcd, _gcdg, _gde, _bae, _abc, _df float64) (_eb, _gdb float64) {
	_abb := 1 - _df
	_gcg := _abb * _abb * _abb
	_ae := 3 * _abb * _abb * _df
	_ecc := 3 * _abb * _df * _df
	_dgd := _df * _df * _df
	_eb = _gcg*_def + _ae*_fa + _ecc*_gcdg + _dgd*_bae
	_gdb = _gcg*_ba + _ae*_gcd + _ecc*_gde + _dgd*_abc
	return
}
func (_aab *Context) StrokePattern() _gc.Pattern { return _aab._faf }
func (_ffbb *Context) MoveTo(x, y float64) {
	if _ffbb._aga {
		_ffbb._bca.Add1(_aba(_ffbb._eeb))
	}
	x, y = _ffbb.Transform(x, y)
	_ef := _c.NewPoint(x, y)
	_eab := _aba(_ef)
	_ffbb._dce.Start(_eab)
	_ffbb._bca.Start(_eab)
	_ffbb._eeb = _ef
	_ffbb._afdd = _ef
	_ffbb._aga = true
}
func NewLinearGradient(x0, y0, x1, y1 float64) _gc.Gradient {
	_fbag := &linearGradient{_bcbd: x0, _gae: y0, _geg: x1, _baee: y1}
	return _fbag
}
func (_dbgd *Context) Translate(x, y float64) { _dbgd._feg = _dbgd._feg.Translate(x, y) }
func (_fab *Context) Shear(x, y float64)      { _fab._feg.Shear(x, y) }
func (_dcac *patternPainter) Paint(ss []_af.Span, done bool) {
	_gegb := _dcac._fada.Bounds()
	for _, _bbe := range ss {
		if _bbe.Y < _gegb.Min.Y {
			continue
		}
		if _bbe.Y >= _gegb.Max.Y {
			return
		}
		if _bbe.X0 < _gegb.Min.X {
			_bbe.X0 = _gegb.Min.X
		}
		if _bbe.X1 > _gegb.Max.X {
			_bbe.X1 = _gegb.Max.X
		}
		if _bbe.X0 >= _bbe.X1 {
			continue
		}
		const _dcgc = 1<<16 - 1
		_deba := _bbe.Y - _dcac._fada.Rect.Min.Y
		_gdca := _bbe.X0 - _dcac._fada.Rect.Min.X
		_bfdg := (_bbe.Y-_dcac._fada.Rect.Min.Y)*_dcac._fada.Stride + (_bbe.X0-_dcac._fada.Rect.Min.X)*4
		_bge := _bfdg + (_bbe.X1-_bbe.X0)*4
		for _ebfc, _ffca := _bfdg, _gdca; _ebfc < _bge; _ebfc, _ffca = _ebfc+4, _ffca+1 {
			_ggg := _bbe.Alpha
			if _dcac._cfgf != nil {
				_ggg = _ggg * uint32(_dcac._cfgf.AlphaAt(_ffca, _deba).A) / 255
				if _ggg == 0 {
					continue
				}
			}
			_eebf := _dcac._acgf.ColorAt(_ffca, _deba)
			_gbec, _agd, _decf, _ddfbf := _eebf.RGBA()
			_afdg := uint32(_dcac._fada.Pix[_ebfc+0])
			_baf := uint32(_dcac._fada.Pix[_ebfc+1])
			_caebf := uint32(_dcac._fada.Pix[_ebfc+2])
			_daag := uint32(_dcac._fada.Pix[_ebfc+3])
			_fdga := (_dcgc - (_ddfbf * _ggg / _dcgc)) * 0x101
			_dcac._fada.Pix[_ebfc+0] = uint8((_afdg*_fdga + _gbec*_ggg) / _dcgc >> 8)
			_dcac._fada.Pix[_ebfc+1] = uint8((_baf*_fdga + _agd*_ggg) / _dcgc >> 8)
			_dcac._fada.Pix[_ebfc+2] = uint8((_caebf*_fdga + _decf*_ggg) / _dcgc >> 8)
			_dcac._fada.Pix[_ebfc+3] = uint8((_daag*_fdga + _ddfbf*_ggg) / _dcgc >> 8)
		}
	}
}
func (_caeb stops) Less(i, j int) bool { return _caeb[i]._bcdba < _caeb[j]._bcdba }
func (_abg *Context) ScaleAbout(sx, sy, x, y float64) {
	_abg.Translate(x, y)
	_abg.Scale(sx, sy)
	_abg.Translate(-x, -y)
}
func _agab(_dfgc, _dgac, _ecg, _dfddg, _dfaa, _fgcb float64) float64 {
	return _dfgc*_dfddg + _dgac*_dfaa + _ecg*_fgcb
}
func _acfg(_debf float64) float64 { return _debf * _b.Pi / 180 }
func _cbf(_ceab _e.Image) *_e.RGBA {
	_cda := _ceab.Bounds()
	_aff := _e.NewRGBA(_cda)
	_d.Draw(_aff, _cda, _ceab, _cda.Min, _d.Src)
	return _aff
}
func (_bbf *Context) SetStrokeRGBA(r, g, b, a float64) {
	_fbg := _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_bbf._faf = _dba(_fbg)
}
func (_cac *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	_daad, _eada, _ddf, _gfe := x, x+r, x+w-r, x+w
	_acf, _dgg, _fff, _afe := y, y+r, y+h-r, y+h
	_cac.NewSubPath()
	_cac.MoveTo(_eada, _acf)
	_cac.LineTo(_ddf, _acf)
	_cac.DrawArc(_ddf, _dgg, r, _acfg(270), _acfg(360))
	_cac.LineTo(_gfe, _fff)
	_cac.DrawArc(_ddf, _fff, r, _acfg(0), _acfg(90))
	_cac.LineTo(_eada, _afe)
	_cac.DrawArc(_eada, _fff, r, _acfg(90), _acfg(180))
	_cac.LineTo(_daad, _dgg)
	_cac.DrawArc(_eada, _dgg, r, _acfg(180), _acfg(270))
	_cac.ClosePath()
}
func (_efc *Context) DrawRectangle(x, y, w, h float64) {
	_efc.NewSubPath()
	_efc.MoveTo(x, y)
	_efc.LineTo(x+w, y)
	_efc.LineTo(x+w, y+h)
	_efc.LineTo(x, y+h)
	_efc.ClosePath()
}

type stop struct {
	_bcdba float64
	_bcc   _a.Color
}

func (_gad *Context) DrawImageAnchored(im _e.Image, x, y int, ax, ay float64) {
	_fcec := im.Bounds().Size()
	x -= int(ax * float64(_fcec.X))
	y -= int(ay * float64(_fcec.Y))
	_ceba := _ed.BiLinear
	_gbde := _gad._feg.Clone().Translate(float64(x), float64(y))
	_eded := _ab.Aff3{_gbde[0], _gbde[3], _gbde[6], _gbde[1], _gbde[4], _gbde[7]}
	if _gad._dff == nil {
		_ceba.Transform(_gad._edgd, _eded, im, im.Bounds(), _ed.Over, nil)
	} else {
		_ceba.Transform(_gad._edgd, _eded, im, im.Bounds(), _ed.Over, &_ed.Options{DstMask: _gad._dff, DstMaskP: _e.Point{}})
	}
}
func (_gcaa *Context) SetStrokeStyle(pattern _gc.Pattern) { _gcaa._faf = pattern }
func (_eeg *Context) FillPreserve() {
	var _fde _af.Painter
	if _eeg._dff == nil {
		if _dbg, _egg := _eeg._bcd.(*solidPattern); _egg {
			_ad := _af.NewRGBAPainter(_eeg._edgd)
			_ad.SetColor(_dbg._dafb)
			_fde = _ad
		}
	}
	if _fde == nil {
		_fde = _dace(_eeg._edgd, _eeg._dff, _eeg._bcd)
	}
	_eeg.fill(_fde)
}
func (_ebf *Context) FillPattern() _gc.Pattern { return _ebf._bcd }
func (_cag *Context) fill(_edcg _af.Painter) {
	_bgf := _cag._bca
	if _cag._aga {
		_bgf = make(_af.Path, len(_cag._bca))
		copy(_bgf, _cag._bca)
		_bgf.Add1(_aba(_cag._eeb))
	}
	_agg := _cag._ge
	_agg.UseNonZeroWinding = _cag._ede == _gc.FillRuleWinding
	_agg.Clear()
	_agg.AddPath(_bgf)
	_agg.Rasterize(_edcg)
}
func (_aggc *Context) drawRegularPolygon(_be int, _edb, _caff, _cegg, _bgdd float64) {
	_cgc := 2 * _b.Pi / float64(_be)
	_bgdd -= _b.Pi / 2
	if _be%2 == 0 {
		_bgdd += _cgc / 2
	}
	_aggc.NewSubPath()
	for _fffe := 0; _fffe < _be; _fffe++ {
		_caee := _bgdd + _cgc*float64(_fffe)
		_aggc.LineTo(_edb+_cegg*_b.Cos(_caee), _caff+_cegg*_b.Sin(_caee))
	}
	_aggc.ClosePath()
}
func _aegg(_gecb [][]_c.Point, _dfcd []float64, _bdac float64) [][]_c.Point {
	var _ebb [][]_c.Point
	if len(_dfcd) == 0 {
		return _gecb
	}
	if len(_dfcd) == 1 {
		_dfcd = append(_dfcd, _dfcd[0])
	}
	for _, _accc := range _gecb {
		if len(_accc) < 2 {
			continue
		}
		_agaa := _accc[0]
		_egd := 1
		_egeb := 0
		_eba := 0.0
		if _bdac != 0 {
			var _faa float64
			for _, _bff := range _dfcd {
				_faa += _bff
			}
			_bdac = _b.Mod(_bdac, _faa)
			if _bdac < 0 {
				_bdac += _faa
			}
			for _dggf, _deff := range _dfcd {
				_bdac -= _deff
				if _bdac < 0 {
					_egeb = _dggf
					_eba = _deff + _bdac
					break
				}
			}
		}
		var _acca []_c.Point
		_acca = append(_acca, _agaa)
		for _egd < len(_accc) {
			_eefc := _dfcd[_egeb]
			_cgde := _accc[_egd]
			_cgg := _agaa.Distance(_cgde)
			_fefc := _eefc - _eba
			if _cgg > _fefc {
				_afdcd := _fefc / _cgg
				_bgc := _agaa.Interpolate(_cgde, _afdcd)
				_acca = append(_acca, _bgc)
				if _egeb%2 == 0 && len(_acca) > 1 {
					_ebb = append(_ebb, _acca)
				}
				_acca = nil
				_acca = append(_acca, _bgc)
				_eba = 0
				_agaa = _bgc
				_egeb = (_egeb + 1) % len(_dfcd)
			} else {
				_acca = append(_acca, _cgde)
				_agaa = _cgde
				_eba += _cgg
				_egd++
			}
		}
		if _egeb%2 == 0 && len(_acca) > 1 {
			_ebb = append(_ebb, _acca)
		}
	}
	return _ebb
}
func (_fd *Context) Width() int { return _fd._gbd }
func (_ccc *Context) DrawStringAnchored(s string, face _eg.Face, x, y, ax, ay float64) {
	_gcb, _ffe := _ccc.MeasureString(s, face)
	_ccc.drawString(s, face, x-ax*_gcb, y+ay*_ffe)
}
func (_bcg *Context) Fill() { _bcg.FillPreserve(); _bcg.ClearPath() }
func (_afb *linearGradient) ColorAt(x, y int) _a.Color {
	if len(_afb._eabb) == 0 {
		return _a.Transparent
	}
	_dbb, _ada := float64(x), float64(y)
	_cfa, _gba, _afa, _aef := _afb._bcbd, _afb._gae, _afb._geg, _afb._baee
	_ccb, _aae := _afa-_cfa, _aef-_gba
	if _aae == 0 && _ccb != 0 {
		return _fdc((_dbb-_cfa)/_ccb, _afb._eabb)
	}
	if _ccb == 0 && _aae != 0 {
		return _fdc((_ada-_gba)/_aae, _afb._eabb)
	}
	_eee := _ccb*(_dbb-_cfa) + _aae*(_ada-_gba)
	if _eee < 0 {
		return _afb._eabb[0]._bcc
	}
	_fcbf := _b.Hypot(_ccb, _aae)
	_bec := ((_dbb-_cfa)*-_aae + (_ada-_gba)*_ccb) / (_fcbf * _fcbf)
	_adad, _acc := _cfa+_bec*-_aae, _gba+_bec*_ccb
	_dee := _b.Hypot(_dbb-_adad, _ada-_acc) / _fcbf
	return _fdc(_dee, _afb._eabb)
}

type patternPainter struct {
	_fada *_e.RGBA
	_cfgf *_e.Alpha
	_acgf _gc.Pattern
}

func (_bcba *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const _agf = 16
	for _ddfg := 0; _ddfg < _agf; _ddfg++ {
		_fec := float64(_ddfg+0) / _agf
		_fded := float64(_ddfg+1) / _agf
		_bcdf := angle1 + (angle2-angle1)*_fec
		_efbc := angle1 + (angle2-angle1)*_fded
		_eabd := x + rx*_b.Cos(_bcdf)
		_dda := y + ry*_b.Sin(_bcdf)
		_deg := x + rx*_b.Cos((_bcdf+_efbc)/2)
		_cbeg := y + ry*_b.Sin((_bcdf+_efbc)/2)
		_dada := x + rx*_b.Cos(_efbc)
		_gbfc := y + ry*_b.Sin(_efbc)
		_ggee := 2*_deg - _eabd/2 - _dada/2
		_dfg := 2*_cbeg - _dda/2 - _gbfc/2
		if _ddfg == 0 {
			if _bcba._aga {
				_bcba.LineTo(_eabd, _dda)
			} else {
				_bcba.MoveTo(_eabd, _dda)
			}
		}
		_bcba.QuadraticTo(_ggee, _dfg, _dada, _gbfc)
	}
}
func (_ccf *Context) Clip()      { _ccf.ClipPreserve(); _ccf.ClearPath() }
func (_aabd *Context) Identity() { _aabd._feg = _c.IdentityMatrix() }
func _edg(_gca, _cdd, _bfb, _cc, _cca, _ce, _aag, _ceb float64) []_c.Point {
	_ega := (_b.Hypot(_bfb-_gca, _cc-_cdd) + _b.Hypot(_cca-_bfb, _ce-_cc) + _b.Hypot(_aag-_cca, _ceb-_ce))
	_fga := int(_ega + 0.5)
	if _fga < 4 {
		_fga = 4
	}
	_gge := float64(_fga) - 1
	_daa := make([]_c.Point, _fga)
	for _ebg := 0; _ebg < _fga; _ebg++ {
		_fc := float64(_ebg) / _gge
		_bdg, _afd := _cf(_gca, _cdd, _bfb, _cc, _cca, _ce, _aag, _ceb, _fc)
		_daa[_ebg] = _c.NewPoint(_bdg, _afd)
	}
	return _daa
}

type repeatOp int

func (_cdgg *Context) SetFillRGBA(r, g, b, a float64) {
	_, _, _, _ege := _cdgg._aeg.RGBA()
	if _ege > 0 && _ege != 65535 && a == 1 {
		a = float64(_ege) / 65535
	}
	_bgd := _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_cdgg._aeg = _bgd
	_cdgg._bcd = _dba(_bgd)
}
func _edcc(_cbbc, _bbcb _a.Color, _gdac float64) _a.Color {
	_cde, _cgd, _bgfe, _edd := _cbbc.RGBA()
	_dde, _bbgcg, _gdgf, _gbaa := _bbcb.RGBA()
	return _a.RGBA{_ade(_cde, _dde, _gdac), _ade(_cgd, _bbgcg, _gdac), _ade(_bgfe, _gdgf, _gdac), _ade(_edd, _gbaa, _gdac)}
}

type solidPattern struct{ _dafb _a.Color }

func (_gedg *Context) Rotate(angle float64) { _gedg._feg = _gedg._feg.Rotate(angle) }
func _aba(_bebd _c.Point) _ec.Point26_6     { return _ec.Point26_6{X: _ebc(_bebd.X), Y: _ebc(_bebd.Y)} }
func NewContext(width, height int) *Context {
	return NewContextForRGBA(_e.NewRGBA(_e.Rect(0, 0, width, height)))
}
func (_eebd *Context) TextState() *_gc.TextState { return &_eebd._cdg }
func (_fdef *Context) SetMask(mask *_e.Alpha) error {
	if mask.Bounds().Size() != _fdef._edgd.Bounds().Size() {
		return _bd.New("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065")
	}
	_fdef._dff = mask
	return nil
}
func (_cba *Context) DrawCircle(x, y, r float64) {
	_cba.NewSubPath()
	_cba.DrawEllipticalArc(x, y, r, r, 0, 2*_b.Pi)
	_cba.ClosePath()
}

const (
	_ecgg repeatOp = iota
	_cbd
	_bgfg
	_acdd
)

type linearGradient struct {
	_bcbd, _gae, _geg, _baee float64
	_eabb                    stops
}

func _dace(_agfcg *_e.RGBA, _facc *_e.Alpha, _bebaf _gc.Pattern) *patternPainter {
	return &patternPainter{_agfcg, _facc, _bebaf}
}

type Context struct {
	_gbd  int
	_ebe  int
	_ge   *_af.Rasterizer
	_edgd *_e.RGBA
	_dff  *_e.Alpha
	_aeg  _a.Color
	_bcd  _gc.Pattern
	_faf  _gc.Pattern
	_dce  _af.Path
	_bca  _af.Path
	_eeb  _c.Point
	_afdd _c.Point
	_aga  bool
	_dec  []float64
	_bfd  float64
	_fcc  float64
	_caf  _gc.LineCap
	_fcf  _gc.LineJoin
	_ede  _gc.FillRule
	_feg  _c.Matrix
	_cdg  _gc.TextState
	_gce  []*Context
}

func (_cdfc *Context) MeasureString(s string, face _eg.Face) (_eabf, _aaa float64) {
	_aaca := &_eg.Drawer{Face: face}
	_bdb := _aaca.MeasureString(s)
	return float64(_bdb >> 6), _cdfc._cdg.Tf.Size
}
func (_eca *Context) DrawString(s string, face _eg.Face, x, y float64) {
	_eca.DrawStringAnchored(s, face, x, y, 0, 0)
}
func (_dfb *Context) SetMatrix(m _c.Matrix) { _dfb._feg = m }
func _ade(_ebec, _bbca uint32, _fbc float64) uint8 {
	return uint8(int32(float64(_ebec)*(1.0-_fbc)+float64(_bbca)*_fbc) >> 8)
}
func (_fcfd *Context) Image() _e.Image { return _fcfd._edgd }
func (_beba *surfacePattern) ColorAt(x, y int) _a.Color {
	_dgfd := _beba._caa.Bounds()
	switch _beba._fafb {
	case _cbd:
		if y >= _dgfd.Dy() {
			return _a.Transparent
		}
	case _bgfg:
		if x >= _dgfd.Dx() {
			return _a.Transparent
		}
	case _acdd:
		if x >= _dgfd.Dx() || y >= _dgfd.Dy() {
			return _a.Transparent
		}
	}
	x = x%_dgfd.Dx() + _dgfd.Min.X
	y = y%_dgfd.Dy() + _dgfd.Min.Y
	return _beba._caa.At(x, y)
}
func (_fbd *Context) AsMask() *_e.Alpha {
	_daf := _e.NewAlpha(_fbd._edgd.Bounds())
	_ed.Draw(_daf, _fbd._edgd.Bounds(), _fbd._edgd, _e.Point{}, _ed.Src)
	return _daf
}
func (_fbe *Context) InvertMask() {
	if _fbe._dff == nil {
		_fbe._dff = _e.NewAlpha(_fbe._edgd.Bounds())
	} else {
		for _gbe, _cbg := range _fbe._dff.Pix {
			_fbe._dff.Pix[_gbe] = 255 - _cbg
		}
	}
}
func (_ffa *Context) joiner() _af.Joiner {
	switch _ffa._fcf {
	case _gc.LineJoinBevel:
		return _af.BevelJoiner
	case _gc.LineJoinRound:
		return _af.RoundJoiner
	}
	return nil
}
func (_dfd *Context) setFillAndStrokeColor(_gfb _a.Color) {
	_dfd._aeg = _gfb
	_dfd._bcd = _dba(_gfb)
	_dfd._faf = _dba(_gfb)
}
func (_cddd *Context) stroke(_bbc _af.Painter) {
	_bda := _cddd._dce
	if len(_cddd._dec) > 0 {
		_bda = _baeg(_bda, _cddd._dec, _cddd._bfd)
	} else {
		_bda = _ffad(_afdb(_bda))
	}
	_ddg := _cddd._ge
	_ddg.UseNonZeroWinding = true
	_ddg.Clear()
	_fbb := (_cddd._feg.ScalingFactorX() + _cddd._feg.ScalingFactorY()) / 2
	_ddg.AddStroke(_bda, _ebc(_cddd._fcc*_fbb), _cddd.capper(), _cddd.joiner())
	_ddg.Rasterize(_bbc)
}
func (_dd *Context) SetLineWidth(lineWidth float64) { _dd._fcc = lineWidth }
func (_bbbg *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	_bbbg.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}
func (_aec *radialGradient) ColorAt(x, y int) _a.Color {
	if len(_aec._fee) == 0 {
		return _a.Transparent
	}
	_fcg, _abd := float64(x)+0.5-_aec._dbd._fbagb, float64(y)+0.5-_aec._dbd._fdedb
	_bbgc := _agab(_fcg, _abd, _aec._dbd._cegc, _aec._bgg._fbagb, _aec._bgg._fdedb, _aec._bgg._cegc)
	_eebb := _agab(_fcg, _abd, -_aec._dbd._cegc, _fcg, _abd, _aec._dbd._cegc)
	if _aec._ccce == 0 {
		if _bbgc == 0 {
			return _a.Transparent
		}
		_dcc := 0.5 * _eebb / _bbgc
		if _dcc*_aec._bgg._cegc >= _aec._aee {
			return _fdc(_dcc, _aec._fee)
		}
		return _a.Transparent
	}
	_daed := _agab(_bbgc, _aec._ccce, 0, _bbgc, -_eebb, 0)
	if _daed >= 0 {
		_bcgb := _b.Sqrt(_daed)
		_gfc := (_bbgc + _bcgb) * _aec._bdf
		_ffcb := (_bbgc - _bcgb) * _aec._bdf
		if _gfc*_aec._bgg._cegc >= _aec._aee {
			return _fdc(_gfc, _aec._fee)
		} else if _ffcb*_aec._bgg._cegc >= _aec._aee {
			return _fdc(_ffcb, _aec._fee)
		}
	}
	return _a.Transparent
}
func (_fgf *Context) SetFillRule(fillRule _gc.FillRule) { _fgf._ede = fillRule }
func _edcd(_agc _e.Image, _ddfb repeatOp) _gc.Pattern {
	return &surfacePattern{_caa: _agc, _fafb: _ddfb}
}

type radialGradient struct {
	_dbd, _fae, _bgg circle
	_ccce, _bdf      float64
	_aee             float64
	_fee             stops
}

func _ee(_gd, _abe, _ag, _ffd, _gg, _gb float64) []_c.Point {
	_ea := (_b.Hypot(_ag-_gd, _ffd-_abe) + _b.Hypot(_gg-_ag, _gb-_ffd))
	_edc := int(_ea + 0.5)
	if _edc < 4 {
		_edc = 4
	}
	_dg := float64(_edc) - 1
	_da := make([]_c.Point, _edc)
	for _bg := 0; _bg < _edc; _bg++ {
		_aa := float64(_bg) / _dg
		_gbf, _ead := _cb(_gd, _abe, _ag, _ffd, _gg, _gb, _aa)
		_da[_bg] = _c.NewPoint(_gbf, _ead)
	}
	return _da
}
func NewContextForImage(im _e.Image) *Context { return NewContextForRGBA(_cbf(im)) }
func (_ecd *Context) SetRGBA255(r, g, b, a int) {
	_ecd._aeg = _a.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	_ecd.setFillAndStrokeColor(_ecd._aeg)
}
func (_cg *Context) SetLineJoin(lineJoin _gc.LineJoin) { _cg._fcf = lineJoin }
func (_fef *Context) DrawEllipse(x, y, rx, ry float64) {
	_fef.NewSubPath()
	_fef.DrawEllipticalArc(x, y, rx, ry, 0, 2*_b.Pi)
	_fef.ClosePath()
}
func (_cbe *Context) ClipPreserve() {
	_cgf := _e.NewAlpha(_e.Rect(0, 0, _cbe._gbd, _cbe._ebe))
	_eebg := _af.NewAlphaOverPainter(_cgf)
	_cbe.fill(_eebg)
	if _cbe._dff == nil {
		_cbe._dff = _cgf
	} else {
		_fce := _e.NewAlpha(_e.Rect(0, 0, _cbe._gbd, _cbe._ebe))
		_ed.DrawMask(_fce, _fce.Bounds(), _cgf, _e.Point{}, _cbe._dff, _e.Point{}, _ed.Over)
		_cbe._dff = _fce
	}
}
func (_fcfdd *Context) drawString(_agfg string, _bcda _eg.Face, _gef, _gda float64) {
	_acd := &_eg.Drawer{Src: _e.NewUniform(_fcfdd._aeg), Face: _bcda, Dot: _aba(_c.NewPoint(_gef, _gda))}
	_bad := rune(-1)
	for _, _fdg := range _agfg {
		if _bad >= 0 {
			_acd.Dot.X += _acd.Face.Kern(_bad, _fdg)
		}
		_agfc, _ggc, _eef, _gfbd, _bee := _acd.Face.Glyph(_acd.Dot, _fdg)
		if !_bee {
			continue
		}
		_eae := _agfc.Sub(_agfc.Min)
		_fcfg := _e.NewRGBA(_eae)
		_ed.DrawMask(_fcfg, _eae, _acd.Src, _e.Point{}, _ggc, _eef, _ed.Over)
		var _bcaba *_ed.Options
		if _fcfdd._dff != nil {
			_bcaba = &_ed.Options{DstMask: _fcfdd._dff, DstMaskP: _e.Point{}}
		}
		_dgdc := _fcfdd._feg.Clone().Translate(float64(_agfc.Min.X), float64(_agfc.Min.Y))
		_ece := _ab.Aff3{_dgdc[0], _dgdc[3], _dgdc[6], _dgdc[1], _dgdc[4], _dgdc[7]}
		_ed.BiLinear.Transform(_fcfdd._edgd, _ece, _fcfg, _eae, _ed.Over, _bcaba)
		_acd.Dot.X += _gfbd
		_bad = _fdg
	}
}
func (_dggd *Context) Pop() {
	_gcgg := *_dggd
	_abca := _dggd._gce
	_bba := _abca[len(_abca)-1]
	*_dggd = *_bba
	_dggd._dce = _gcgg._dce
	_dggd._bca = _gcgg._bca
	_dggd._eeb = _gcgg._eeb
	_dggd._afdd = _gcgg._afdd
	_dggd._aga = _gcgg._aga
}

type stops []stop

func (_bfe *Context) RotateAbout(angle, x, y float64) {
	_bfe.Translate(x, y)
	_bfe.Rotate(angle)
	_bfe.Translate(-x, -y)
}
func _ebc(_bdfe float64) _ec.Int26_6 { return _ec.Int26_6(_bdfe * 64) }
func (_dffa *Context) SetHexColor(x string) {
	_gbfa, _gec, _afdc, _fba := _cgca(x)
	_dffa.SetRGBA255(_gbfa, _gec, _afdc, _fba)
}
func (_dac *Context) DrawLine(x1, y1, x2, y2 float64) { _dac.MoveTo(x1, y1); _dac.LineTo(x2, y2) }
func (_dbgde *Context) Transform(x, y float64) (_ccgd, _cfb float64) {
	return _dbgde._feg.Transform(x, y)
}

type surfacePattern struct {
	_caa  _e.Image
	_fafb repeatOp
}

func _afdb(_dgdg _af.Path) [][]_c.Point {
	var _ecga [][]_c.Point
	var _eeba []_c.Point
	var _fcfgc, _ccba float64
	for _bgdg := 0; _bgdg < len(_dgdg); {
		switch _dgdg[_bgdg] {
		case 0:
			if len(_eeba) > 0 {
				_ecga = append(_ecga, _eeba)
				_eeba = nil
			}
			_cbbcf := _ffde(_dgdg[_bgdg+1])
			_dadf := _ffde(_dgdg[_bgdg+2])
			_eeba = append(_eeba, _c.NewPoint(_cbbcf, _dadf))
			_fcfgc, _ccba = _cbbcf, _dadf
			_bgdg += 4
		case 1:
			_cfg := _ffde(_dgdg[_bgdg+1])
			_aggd := _ffde(_dgdg[_bgdg+2])
			_eeba = append(_eeba, _c.NewPoint(_cfg, _aggd))
			_fcfgc, _ccba = _cfg, _aggd
			_bgdg += 4
		case 2:
			_adc := _ffde(_dgdg[_bgdg+1])
			_cega := _ffde(_dgdg[_bgdg+2])
			_efce := _ffde(_dgdg[_bgdg+3])
			_bea := _ffde(_dgdg[_bgdg+4])
			_eec := _ee(_fcfgc, _ccba, _adc, _cega, _efce, _bea)
			_eeba = append(_eeba, _eec...)
			_fcfgc, _ccba = _efce, _bea
			_bgdg += 6
		case 3:
			_eac := _ffde(_dgdg[_bgdg+1])
			_egc := _ffde(_dgdg[_bgdg+2])
			_ecdc := _ffde(_dgdg[_bgdg+3])
			_gbca := _ffde(_dgdg[_bgdg+4])
			_edfg := _ffde(_dgdg[_bgdg+5])
			_eefg := _ffde(_dgdg[_bgdg+6])
			_dgc := _edg(_fcfgc, _ccba, _eac, _egc, _ecdc, _gbca, _edfg, _eefg)
			_eeba = append(_eeba, _dgc...)
			_fcfgc, _ccba = _edfg, _eefg
			_bgdg += 8
		default:
			_bb.Log.Debug("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076", _dgdg)
			return _ecga
		}
	}
	if len(_eeba) > 0 {
		_ecga = append(_ecga, _eeba)
	}
	return _ecga
}
func (_gcc *Context) SetDashOffset(offset float64) { _gcc._bfd = offset }
func (_fdd *Context) Height() int                  { return _fdd._ebe }
func (_cdgca *Context) Scale(x, y float64)         { _cdgca._feg = _cdgca._feg.Scale(x, y) }
func (_bbcg *radialGradient) AddColorStop(offset float64, color _a.Color) {
	_bbcg._fee = append(_bbcg._fee, stop{_bcdba: offset, _bcc: color})
	_g.Sort(_bbcg._fee)
}
func _ffde(_cfae _ec.Int26_6) float64 {
	const _gfa, _bag = 6, 1<<6 - 1
	if _cfae >= 0 {
		return float64(_cfae>>_gfa) + float64(_cfae&_bag)/64
	}
	_cfae = -_cfae
	if _cfae >= 0 {
		return -(float64(_cfae>>_gfa) + float64(_cfae&_bag)/64)
	}
	return 0
}
func (_dfcf *Context) LineTo(x, y float64) {
	if !_dfcf._aga {
		_dfcf.MoveTo(x, y)
	} else {
		x, y = _dfcf.Transform(x, y)
		_ceg := _c.NewPoint(x, y)
		_bcf := _aba(_ceg)
		_dfcf._dce.Add1(_bcf)
		_dfcf._bca.Add1(_bcf)
		_dfcf._afdd = _ceg
	}
}
func (_eag *Context) Stroke() { _eag.StrokePreserve(); _eag.ClearPath() }
func NewRadialGradient(x0, y0, r0, x1, y1, r1 float64) _gc.Gradient {
	_bbgcf := circle{x0, y0, r0}
	_gedf := circle{x1, y1, r1}
	_dgf := circle{x1 - x0, y1 - y0, r1 - r0}
	_eccb := _agab(_dgf._fbagb, _dgf._fdedb, -_dgf._cegc, _dgf._fbagb, _dgf._fdedb, _dgf._cegc)
	var _gbc float64
	if _eccb != 0 {
		_gbc = 1.0 / _eccb
	}
	_aabb := -_bbgcf._cegc
	_aeb := &radialGradient{_dbd: _bbgcf, _fae: _gedf, _bgg: _dgf, _ccce: _eccb, _bdf: _gbc, _aee: _aabb}
	return _aeb
}
func _ffad(_gdf [][]_c.Point) _af.Path {
	var _age _af.Path
	for _, _dea := range _gdf {
		var _eddc _ec.Point26_6
		for _edfa, _gcee := range _dea {
			_cbc := _aba(_gcee)
			if _edfa == 0 {
				_age.Start(_cbc)
			} else {
				_deec := _cbc.X - _eddc.X
				_cbag := _cbc.Y - _eddc.Y
				if _deec < 0 {
					_deec = -_deec
				}
				if _cbag < 0 {
					_cbag = -_cbag
				}
				if _deec+_cbag > 8 {
					_age.Add1(_cbc)
				}
			}
			_eddc = _cbc
		}
	}
	return _age
}
func (_dgda *Context) capper() _af.Capper {
	switch _dgda._caf {
	case _gc.LineCapButt:
		return _af.ButtCapper
	case _gc.LineCapRound:
		return _af.RoundCapper
	case _gc.LineCapSquare:
		return _af.SquareCapper
	}
	return nil
}
