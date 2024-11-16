package context

import (
	_fa "errors"
	_d "image"
	_a "image/color"
	_c "strconv"
	_f "strings"

	_gb "github.com/bamzi/pdfext/core"
	_ff "github.com/bamzi/pdfext/internal/cmap"
	_e "github.com/bamzi/pdfext/internal/textencoding"
	_ge "github.com/bamzi/pdfext/internal/transform"
	_fg "github.com/bamzi/pdfext/model"
	_df "github.com/unidoc/freetype/truetype"
	_gc "golang.org/x/image/font"
)

func (_fac *TextFont) BytesToCharcodes(data []byte) []_e.CharCode {
	if _fac._fcb != nil {
		return _fac._fcb.BytesToCharcodes(data)
	}
	return _fac.Font.BytesToCharcodes(data)
}

type LineJoin int

const (
	TextRenderingModeFill TextRenderingMode = iota
	TextRenderingModeStroke
	TextRenderingModeFillStroke
	TextRenderingModeInvisible
	TextRenderingModeFillClip
	TextRenderingModeStrokeClip
	TextRenderingModeFillStrokeClip
	TextRenderingModeClip
)

func (_gbb *TextState) ProcTd(tx, ty float64) {
	_gbb.Tlm.Concat(_ge.TranslationMatrix(tx, ty))
	_gbb.Tm = _gbb.Tlm.Clone()
}

type TextFont struct {
	Font *_fg.PdfFont
	Size float64
	_dbe *_df.Font
	_fcb *_fg.PdfFont
}

func (_gbf *TextFont) CharcodeToRunes(charcode _e.CharCode) (_e.CharCode, []rune) {
	_ebc := []_e.CharCode{charcode}
	if _gbf._fcb == nil || _gbf._fcb == _gbf.Font {
		return _gbf.charcodeToRunesSimple(charcode)
	}
	_ega := _gbf._fcb.CharcodesToUnicode(_ebc)
	_dd, _ := _gbf.Font.RunesToCharcodeBytes(_ega)
	_gef := _gbf.Font.BytesToCharcodes(_dd)
	_dabf := charcode
	if len(_gef) > 0 && _gef[0] != 0 {
		_dabf = _gef[0]
	}
	if string(_ega) == string(_ff.MissingCodeRune) && _gbf._fcb.BaseFont() == _gbf.Font.BaseFont() {
		return _gbf.charcodeToRunesSimple(charcode)
	}
	return _dabf, _ega
}
func NewTextFont(font *_fg.PdfFont, size float64) (*TextFont, error) {
	_aed := font.FontDescriptor()
	if _aed == nil {
		return nil, _fa.New("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072")
	}
	_af, _dab := _gb.GetStream(_aed.FontFile2)
	if !_dab {
		return nil, _fa.New("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_gdda, _aeg := _gb.DecodeStream(_af)
	if _aeg != nil {
		return nil, _aeg
	}
	_eef, _aeg := _df.Parse(_gdda)
	if _aeg != nil {
		return nil, _aeg
	}
	_fgg := font.FontDescriptor().FontName.String()
	_bbg := len(_fgg) > 7 && _fgg[6] == '+'
	if _aed.Flags != nil {
		_efd, _fea := _c.Atoi(_aed.Flags.String())
		if _fea == nil && _efd == 32 {
			_bbg = false
		}
	}
	if !_eef.HasCmap() && (!_f.Contains(font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") || !_bbg) {
		return nil, _fa.New("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079")
	}
	return &TextFont{Font: font, Size: size, _dbe: _eef}, nil
}

type Context interface {
	Push()
	Pop()
	Matrix() _ge.Matrix
	SetMatrix(_ee _ge.Matrix)
	Translate(_aa, _cc float64)
	Scale(_da, _fe float64)
	Rotate(_bfb float64)
	MoveTo(_aae, _ac float64)
	LineTo(_adf, _eb float64)
	CubicTo(_eee, _ef, _ccc, _db, _de, _bd float64)
	QuadraticTo(_ga, _ea, _dba, _ead float64)
	NewSubPath()
	ClosePath()
	ClearPath()
	Clip()
	ClipPreserve()
	ResetClip()
	LineWidth() float64
	SetLineWidth(_cb float64)
	SetLineCap(_ec LineCap)
	SetLineJoin(_bfc LineJoin)
	SetDash(_bg ...float64)
	SetDashOffset(_bgg float64)
	Fill()
	FillPreserve()
	Stroke()
	StrokePreserve()
	SetRGBA(_gcd, _fag, _eg, _ed float64)
	SetFillRGBA(_ba, _fd, _gcc, _eed float64)
	SetFillStyle(_gag Pattern)
	SetFillRule(_gg FillRule)
	SetStrokeRGBA(_fc, _cg, _eac, _gac float64)
	SetStrokeStyle(_bb Pattern)
	FillPattern() Pattern
	StrokePattern() Pattern
	TextState() *TextState
	DrawString(_gdd string, _ag _gc.Face, _aaed, _dg float64)
	MeasureString(_gcdf string, _eeb _gc.Face) (_edf, _fb float64)
	DrawRectangle(_ffg, _fdc, _dgc, _cf float64)
	DrawImage(_aga _d.Image, _fdd, _ab int)
	DrawImageAnchored(_feb _d.Image, _faa, _bdg int, _ae, _baa float64)
	Height() int
	Width() int
}

func (_dbac *TextState) ProcTj(data []byte, ctx Context) {
	_be := _dbac.Tf.Size
	_afd := _dbac.Th / 100.0
	_bbgb := _dbac.GlobalScale
	_fcg := _ge.NewMatrix(_be*_afd, 0, 0, _be, 0, _dbac.Ts)
	_ebb := ctx.Matrix()
	_afa := _ebb.Clone().Mult(_dbac.Tm.Clone().Mult(_fcg)).ScalingFactorY()
	_ged := _dbac.Tf.NewFace(_afa)
	_age := _dbac.Tf.BytesToCharcodes(data)
	for _, _fda := range _age {
		_bad, _acb := _dbac.Tf.CharcodeToRunes(_fda)
		_ggc := string(_acb)
		if _ggc == "\u0000" {
			continue
		}
		_ccf := _ebb.Clone().Mult(_dbac.Tm.Clone().Mult(_fcg))
		_fdg := _ccf.ScalingFactorY()
		_ccf = _ccf.Scale(1/_fdg, -1/_fdg)
		if _dbac.Tr != TextRenderingModeInvisible {
			ctx.SetMatrix(_ccf)
			ctx.DrawString(_ggc, _ged, 0, 0)
			ctx.SetMatrix(_ebb)
		}
		_acf := 0.0
		if _ggc == "\u0020" {
			_acf = _dbac.Tw
		}
		_ebf, _, _daf := _dbac.Tf.GetCharMetrics(_bad)
		if _daf {
			_ebf = _ebf * 0.001 * _be
		} else {
			_ebf, _ = ctx.MeasureString(_ggc, _ged)
			_ebf = _ebf / _bbgb
		}
		_bc := (_ebf + _dbac.Tc + _acf) * _afd
		_dbac.Tm = _dbac.Tm.Mult(_ge.TranslationMatrix(_bc, 0))
	}
}
func (_eba *TextFont) charcodeToRunesSimple(_fddg _e.CharCode) (_e.CharCode, []rune) {
	_bbe := []_e.CharCode{_fddg}
	if _eba.Font.IsSimple() && _eba._dbe != nil {
		if _aab := _eba._dbe.Index(rune(_fddg)); _aab > 0 {
			return _fddg, []rune{rune(_fddg)}
		}
	}
	if _eba._dbe != nil && !_eba._dbe.HasCmap() && _f.Contains(_eba.Font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") {
		if _bgf := _eba._dbe.Index(rune(_fddg)); _bgf > 0 {
			return _fddg, []rune{rune(_fddg)}
		}
	}
	return _fddg, _eba.Font.CharcodesToUnicode(_bbe)
}
func (_deb *TextFont) WithSize(size float64, originalFont *_fg.PdfFont) *TextFont {
	return &TextFont{Font: _deb.Font, Size: size, _dbe: _deb._dbe, _fcb: originalFont}
}

type LineCap int

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

func (_adfb *TextState) ProcTStar()                    { _adfb.ProcTd(0, -_adfb.Tl) }
func (_fef *TextState) ProcQ(data []byte, ctx Context) { _fef.ProcTStar(); _fef.ProcTj(data, ctx) }

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

func NewTextState() TextState {
	return TextState{Th: 100, Tm: _ge.IdentityMatrix(), Tlm: _ge.IdentityMatrix()}
}
func (_aabb *TextState) ProcTD(tx, ty float64) { _aabb.Tl = -ty; _aabb.ProcTd(tx, ty) }
func (_eff *TextFont) GetCharMetrics(code _e.CharCode) (float64, float64, bool) {
	if _fbg, _fcc := _eff.Font.GetCharMetrics(code); _fcc && _fbg.Wx != 0 {
		return _fbg.Wx, _fbg.Wy, _fcc
	}
	if _eff._fcb == nil {
		return 0, 0, false
	}
	_debc, _cbd := _eff._fcb.GetCharMetrics(code)
	return _debc.Wx, _debc.Wy, _cbd && _debc.Wx != 0
}

type TextRenderingMode int

func (_bbb *TextState) ProcTf(font *TextFont) { _bbb.Tf = font }

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

func (_fbb *TextFont) NewFace(size float64) _gc.Face {
	return _df.NewFace(_fbb._dbe, &_df.Options{Size: size})
}
func (_fead *TextState) Translate(tx, ty float64) {
	_fead.Tm = _fead.Tm.Mult(_ge.TranslationMatrix(tx, ty))
}
func (_faf *TextState) ProcTm(a, b, c, d, e, f float64) {
	_faf.Tm = _ge.NewMatrix(a, b, c, d, e, f)
	_faf.Tlm = _faf.Tm.Clone()
}

type FillRule int
type Pattern interface{ ColorAt(_b, _gd int) _a.Color }

func (_gdc *TextState) ProcDQ(data []byte, aw, ac float64, ctx Context) {
	_gdc.Tw = aw
	_gdc.Tc = ac
	_gdc.ProcQ(data, ctx)
}

type TextState struct {
	Tc          float64
	Tw          float64
	Th          float64
	Tl          float64
	Tf          *TextFont
	Ts          float64
	Tm          _ge.Matrix
	Tlm         _ge.Matrix
	Tr          TextRenderingMode
	GlobalScale float64
}

func NewTextFontFromPath(filePath string, size float64) (*TextFont, error) {
	_bba, _cfa := _fg.NewPdfFontFromTTFFile(filePath)
	if _cfa != nil {
		return nil, _cfa
	}
	return NewTextFont(_bba, size)
}
func (_eec *TextState) Reset() { _eec.Tm = _ge.IdentityMatrix(); _eec.Tlm = _ge.IdentityMatrix() }

type Gradient interface {
	Pattern
	AddColorStop(_bf float64, _ad _a.Color)
}
