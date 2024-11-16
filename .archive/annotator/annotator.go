// Package annotator provides an interface for creating annotations with appearance
// streams.  It goes beyond the models package which includes definitions of basic annotation models, in that it
// can create the appearance streams which specify the exact appearance as needed by many pdf viewers for consistent
// appearance of the annotations.
// It also contains methods for generating appearance streams for fields via widget annotations.
package annotator

import (
	_f "bytes"
	_aeg "errors"
	_gb "fmt"
	_ca "image"
	_fg "math"
	_ae "strings"
	_a "time"
	_cb "unicode"

	_g "github.com/bamzi/pdfext/common"
	_fb "github.com/bamzi/pdfext/contentstream"
	_ce "github.com/bamzi/pdfext/contentstream/draw"
	_cac "github.com/bamzi/pdfext/core"
	_b "github.com/bamzi/pdfext/creator"
	_ac "github.com/bamzi/pdfext/internal/textencoding"
	_ge "github.com/bamzi/pdfext/model"
)

// NewCheckboxField generates a new checkbox field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewCheckboxField(page *_ge.PdfPage, name string, rect []float64, opt CheckboxFieldOptions) (*_ge.PdfFieldButton, error) {
	if page == nil {
		return nil, _aeg.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _aeg.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _aeg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_ebbb, _fcdc := _ge.NewStandard14Font(_ge.ZapfDingbatsName)
	if _fcdc != nil {
		return nil, _fcdc
	}
	_agga := _ge.NewPdfField()
	_dfec := &_ge.PdfFieldButton{}
	_agga.SetContext(_dfec)
	_dfec.PdfField = _agga
	_dfec.T = _cac.MakeString(name)
	_dfec.SetType(_ge.ButtonTypeCheckbox)
	_edcd := "\u004f\u0066\u0066"
	if opt.Checked {
		_edcd = "\u0059\u0065\u0073"
	}
	_dfec.V = _cac.MakeName(_edcd)
	_gea := _ge.NewPdfAnnotationWidget()
	_gea.Rect = _cac.MakeArrayFromFloats(rect)
	_gea.P = page.ToPdfObject()
	_gea.F = _cac.MakeInteger(4)
	_gea.Parent = _dfec.ToPdfObject()
	_dcc := rect[2] - rect[0]
	_cbcb := rect[3] - rect[1]
	var _baea _f.Buffer
	_baea.WriteString("\u0071\u000a")
	_baea.WriteString("\u0030 \u0030\u0020\u0031\u0020\u0072\u0067\n")
	_baea.WriteString("\u0042\u0054\u000a")
	_baea.WriteString("\u002f\u005a\u0061D\u0062\u0020\u0031\u0032\u0020\u0054\u0066\u000a")
	_baea.WriteString("\u0045\u0054\u000a")
	_baea.WriteString("\u0051\u000a")
	_cabgg := _fb.NewContentCreator()
	_cabgg.Add_q()
	_cabgg.Add_rg(0, 0, 1)
	_cabgg.Add_BT()
	_cabgg.Add_Tf(*_cac.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_cabgg.Add_Td(0, 0)
	_cabgg.Add_ET()
	_cabgg.Add_Q()
	_bdce := _ge.NewXObjectForm()
	_bdce.SetContentStream(_cabgg.Bytes(), _cac.NewRawEncoder())
	_bdce.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _dcc, _cbcb})
	_bdce.Resources = _ge.NewPdfPageResources()
	_bdce.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _ebbb.ToPdfObject())
	_cabgg = _fb.NewContentCreator()
	_cabgg.Add_q()
	_cabgg.Add_re(0, 0, _dcc, _cbcb)
	_cabgg.Add_W().Add_n()
	_cabgg.Add_rg(0, 0, 1)
	_cabgg.Translate(0, 3.0)
	_cabgg.Add_BT()
	_cabgg.Add_Tf(*_cac.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_cabgg.Add_Td(0, 0)
	_cabgg.Add_Tj(*_cac.MakeString("\u0034"))
	_cabgg.Add_ET()
	_cabgg.Add_Q()
	_beg := _ge.NewXObjectForm()
	_beg.SetContentStream(_cabgg.Bytes(), _cac.NewRawEncoder())
	_beg.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _dcc, _cbcb})
	_beg.Resources = _ge.NewPdfPageResources()
	_beg.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _ebbb.ToPdfObject())
	_fefg := _cac.MakeDict()
	_fefg.Set("\u004f\u0066\u0066", _bdce.ToPdfObject())
	_fefg.Set("\u0059\u0065\u0073", _beg.ToPdfObject())
	_dece := _cac.MakeDict()
	_dece.Set("\u004e", _fefg)
	_gea.AP = _dece
	_gea.AS = _cac.MakeName(_edcd)
	_dfec.Annotations = append(_dfec.Annotations, _gea)
	return _dfec, nil
}

// NewTextField generates a new text field with partial name `name` at location
// specified by `rect` on given `page` and with field specific options `opt`.
func NewTextField(page *_ge.PdfPage, name string, rect []float64, opt TextFieldOptions) (*_ge.PdfFieldText, error) {
	if page == nil {
		return nil, _aeg.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _aeg.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _aeg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_bfac := _ge.NewPdfField()
	_ggfb := &_ge.PdfFieldText{}
	_bfac.SetContext(_ggfb)
	_ggfb.PdfField = _bfac
	_ggfb.T = _cac.MakeString(name)
	if opt.MaxLen > 0 {
		_ggfb.MaxLen = _cac.MakeInteger(int64(opt.MaxLen))
	}
	if len(opt.Value) > 0 {
		_ggfb.V = _cac.MakeString(opt.Value)
	}
	if opt.TextColor != "" {
		_dfed := _b.ColorRGBFromHex(opt.TextColor)
		_cdd, _gacc, _eega := _dfed.ToRGB()
		_fdgce := 12
		if opt.FontSize > 0 {
			_fdgce = opt.FontSize
		}
		_dfaf := "\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a"
		if opt.FontName != "" {
			_dfaf = opt.FontName
		}
		_egg := _gb.Sprintf("/\u0025\u0073\u0020\u0025\u0064\u0020T\u0066\u0020\u0025\u002e\u0033\u0066\u0020\u0025\u002e3\u0066\u0020\u0025.\u0033f\u0020\u0072\u0067", _dfaf, _fdgce, _cdd, _gacc, _eega)
		_ggfb.DA = _cac.MakeString(_egg)
	}
	_bfac.SetContext(_ggfb)
	_dgab := _ge.NewPdfAnnotationWidget()
	_dgab.Rect = _cac.MakeArrayFromFloats(rect)
	_dgab.P = page.ToPdfObject()
	_dgab.F = _cac.MakeInteger(4)
	_dgab.Parent = _ggfb.ToPdfObject()
	_ggfb.Annotations = append(_ggfb.Annotations, _dgab)
	return _ggfb, nil
}

// CircleAnnotationDef defines a circle annotation or ellipse at position (X, Y) and Width and Height.
// The annotation has various style parameters including Fill and Border options and Opacity.
type CircleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_ge.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_ge.PdfColorDeviceRGB
	Opacity       float64
}

// FieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for fields taking into account what value is in the field. A common use case is for generating the
// appearance stream prior to flattening fields.
//
// If `OnlyIfMissing` is true, the field appearance is generated only for fields that do not have an
// appearance stream specified.
// If `RegenerateTextFields` is true, all text fields are regenerated (even if OnlyIfMissing is true).
type FieldAppearance struct {
	OnlyIfMissing        bool
	RegenerateTextFields bool
	_de                  *AppearanceStyle
}

func _aaed(_fae RectangleAnnotationDef) (*_cac.PdfObjectDictionary, *_ge.PdfRectangle, error) {
	_cbbeg := _ge.NewXObjectForm()
	_cbbeg.Resources = _ge.NewPdfPageResources()
	_efae := ""
	if _fae.Opacity < 1.0 {
		_bedf := _cac.MakeDict()
		_bedf.Set("\u0063\u0061", _cac.MakeFloat(_fae.Opacity))
		_bedf.Set("\u0043\u0041", _cac.MakeFloat(_fae.Opacity))
		_fbac := _cbbeg.Resources.AddExtGState("\u0067\u0073\u0031", _bedf)
		if _fbac != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _fbac
		}
		_efae = "\u0067\u0073\u0031"
	}
	_cbag, _gffa, _agbe, _aabd := _cbdd(_fae, _efae)
	if _aabd != nil {
		return nil, nil, _aabd
	}
	_aabd = _cbbeg.SetContentStream(_cbag, nil)
	if _aabd != nil {
		return nil, nil, _aabd
	}
	_cbbeg.BBox = _gffa.ToPdfObject()
	_bagb := _cac.MakeDict()
	_bagb.Set("\u004e", _cbbeg.ToPdfObject())
	return _bagb, _agbe, nil
}
func _gde(_gga *InkAnnotationDef) ([]byte, *_ge.PdfRectangle, error) {
	_gfed := [][]_ce.CubicBezierCurve{}
	for _, _dfgg := range _gga.Paths {
		if _dfgg.Length() == 0 {
			continue
		}
		_eff := _dfgg.Points
		_befb, _adaec, _gfdc := _fbaa(_eff)
		if _gfdc != nil {
			return nil, nil, _gfdc
		}
		if len(_befb) != len(_adaec) {
			return nil, nil, _aeg.New("\u0049\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u006e\u0064\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0063\u006f\u006e\u0074\u0072o\u006c\u0020\u0070\u006f\u0069n\u0074")
		}
		_cfaeg := []_ce.CubicBezierCurve{}
		for _cbgf := 0; _cbgf < len(_befb); _cbgf++ {
			_cfaeg = append(_cfaeg, _ce.CubicBezierCurve{P0: _eff[_cbgf], P1: _befb[_cbgf], P2: _adaec[_cbgf], P3: _eff[_cbgf+1]})
		}
		if len(_cfaeg) > 0 {
			_gfed = append(_gfed, _cfaeg)
		}
	}
	_agbd, _feb, _abcf := _dbe(_gfed, _gga.Color, _gga.LineWidth)
	if _abcf != nil {
		return nil, nil, _abcf
	}
	return _agbd, _feb, nil
}

// NewSignatureLine returns a new signature line displayed as a part of the
// signature field appearance.
func NewSignatureLine(desc, text string) *SignatureLine {
	return &SignatureLine{Desc: desc, Text: text}
}
func _gcea(_cedf []*SignatureLine, _agfa *SignatureFieldOpts) (*_cac.PdfObjectDictionary, error) {
	if _agfa == nil {
		_agfa = NewSignatureFieldOpts()
	}
	var _fgdf error
	var _afcd *_cac.PdfObjectName
	_fbbb := _agfa.Font
	if _fbbb != nil {
		_ebd, _ := _fbbb.GetFontDescriptor()
		if _ebd != nil {
			if _egfg, _cbfe := _ebd.FontName.(*_cac.PdfObjectName); _cbfe {
				_afcd = _egfg
			}
		}
		if _afcd == nil {
			_afcd = _cac.MakeName("\u0046\u006f\u006et\u0031")
		}
	} else {
		if _fbbb, _fgdf = _ge.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a"); _fgdf != nil {
			return nil, _fgdf
		}
		_afcd = _cac.MakeName("\u0048\u0065\u006c\u0076")
	}
	_cdgf := _agfa.FontSize
	if _cdgf <= 0 {
		_cdgf = 10
	}
	if _agfa.LineHeight <= 0 {
		_agfa.LineHeight = 1
	}
	_dfga := _agfa.LineHeight * _cdgf
	_fdf, _fddb := _fbbb.GetRuneMetrics(' ')
	if !_fddb {
		return nil, _aeg.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
	}
	_ebg := _fdf.Wx
	var _dbfd float64
	var _aeged []string
	for _, _bbg := range _cedf {
		if _bbg.Text == "" {
			continue
		}
		_eaba := _bbg.Text
		if _bbg.Desc != "" {
			_eaba = _bbg.Desc + "\u003a\u0020" + _eaba
		}
		_aeged = append(_aeged, _eaba)
		var _fda float64
		for _, _gae := range _eaba {
			_eba, _baga := _fbbb.GetRuneMetrics(_gae)
			if !_baga {
				continue
			}
			_fda += _eba.Wx
		}
		if _fda > _dbfd {
			_dbfd = _fda
		}
	}
	_dbfd = _dbfd * _cdgf / 1000.0
	_bea := float64(len(_aeged)) * _dfga
	_bbca := _agfa.Image != nil
	_gef := _agfa.Rect
	if _gef == nil {
		_gef = []float64{0, 0, _dbfd, _bea}
		if _bbca {
			_gef[2] = _dbfd * 2
			_gef[3] = _bea * 2
		}
		_agfa.Rect = _gef
	}
	_abbc := _gef[2] - _gef[0]
	_edgb := _gef[3] - _gef[1]
	_ccf, _gfcd := _gef, _gef
	var _bece, _fdb float64
	if _bbca && len(_aeged) > 0 {
		if _agfa.ImagePosition <= SignatureImageRight {
			_egfd := []float64{_gef[0], _gef[1], _gef[0] + (_abbc / 2), _gef[3]}
			_ebgg := []float64{_gef[0] + (_abbc / 2), _gef[1], _gef[2], _gef[3]}
			if _agfa.ImagePosition == SignatureImageLeft {
				_ccf, _gfcd = _egfd, _ebgg
			} else {
				_ccf, _gfcd = _ebgg, _egfd
			}
		} else {
			_cda := []float64{_gef[0], _gef[1], _gef[2], _gef[1] + (_edgb / 2)}
			_cbcg := []float64{_gef[0], _gef[1] + (_edgb / 2), _gef[2], _gef[3]}
			if _agfa.ImagePosition == SignatureImageTop {
				_ccf, _gfcd = _cbcg, _cda
			} else {
				_ccf, _gfcd = _cda, _cbcg
			}
		}
	}
	_bece = _gfcd[2] - _gfcd[0]
	_fdb = _gfcd[3] - _gfcd[1]
	var _cfd float64
	if _agfa.AutoSize {
		if _dbfd > _bece || _bea > _fdb {
			_gddad := _fg.Min(_bece/_dbfd, _fdb/_bea)
			_cdgf *= _gddad
		}
		_dfga = _agfa.LineHeight * _cdgf
		_cfd += (_fdb - float64(len(_aeged))*_dfga) / 2
	}
	_dge := _fb.NewContentCreator()
	_daag := _ge.NewPdfPageResources()
	_daag.SetFontByName(*_afcd, _fbbb.ToPdfObject())
	if _agfa.BorderSize <= 0 {
		_agfa.BorderSize = 0
		_agfa.BorderColor = _ge.NewPdfColorDeviceGray(1)
	}
	_dge.Add_q()
	if _agfa.FillColor != nil {
		_dge.SetNonStrokingColor(_agfa.FillColor)
	}
	if _agfa.BorderColor != nil {
		_dge.SetStrokingColor(_agfa.BorderColor)
	}
	_dge.Add_w(_agfa.BorderSize).Add_re(_gef[0], _gef[1], _abbc, _edgb)
	if _agfa.FillColor != nil && _agfa.BorderColor != nil {
		_dge.Add_B()
	} else if _agfa.FillColor != nil {
		_dge.Add_f()
	} else if _agfa.BorderColor != nil {
		_dge.Add_S()
	}
	_dge.Add_Q()
	if _agfa.WatermarkImage != nil {
		_dadc := []float64{_gef[0], _gef[1], _gef[2], _gef[3]}
		_afce, _abge, _efb := _ebgb(_agfa.WatermarkImage, "\u0049\u006d\u0061\u0067\u0065\u0057\u0061\u0074\u0065r\u006d\u0061\u0072\u006b", _agfa, _dadc, _dge)
		if _efb != nil {
			return nil, _efb
		}
		_daag.SetXObjectImageByName(*_afce, _abge)
	}
	_dge.Add_q()
	_dge.Translate(_gfcd[0], _gfcd[3]-_dfga-_cfd)
	_dge.Add_BT()
	_afge := _fbbb.Encoder()
	for _, _adgf := range _aeged {
		var _ede []byte
		for _, _fff := range _adgf {
			if _cb.IsSpace(_fff) {
				if len(_ede) > 0 {
					_dge.SetNonStrokingColor(_agfa.TextColor).Add_Tf(*_afcd, _cdgf).Add_TL(_dfga).Add_TJ([]_cac.PdfObject{_cac.MakeStringFromBytes(_ede)}...)
					_ede = nil
				}
				_dge.Add_Tf(*_afcd, _cdgf).Add_TL(_dfga).Add_TJ([]_cac.PdfObject{_cac.MakeFloat(-_ebg)}...)
			} else {
				_ede = append(_ede, _afge.Encode(string(_fff))...)
			}
		}
		if len(_ede) > 0 {
			_dge.SetNonStrokingColor(_agfa.TextColor).Add_Tf(*_afcd, _cdgf).Add_TL(_dfga).Add_TJ([]_cac.PdfObject{_cac.MakeStringFromBytes(_ede)}...)
		}
		_dge.Add_Td(0, -_dfga)
	}
	_dge.Add_ET()
	_dge.Add_Q()
	if _bbca {
		_ddde, _dgca, _gabc := _ebgb(_agfa.Image, "\u0049\u006d\u0061\u0067\u0065\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065", _agfa, _ccf, _dge)
		if _gabc != nil {
			return nil, _gabc
		}
		_daag.SetXObjectImageByName(*_ddde, _dgca)
	}
	_efc := _ge.NewXObjectForm()
	_efc.Resources = _daag
	_efc.BBox = _cac.MakeArrayFromFloats(_gef)
	_efc.SetContentStream(_dge.Bytes(), _daf())
	_gaa := _cac.MakeDict()
	_gaa.Set("\u004e", _efc.ToPdfObject())
	return _gaa, nil
}
func _ccd(_dfdd *_ge.PdfField, _bff, _bdb float64, _bda string, _ceab AppearanceStyle, _ecdd *_fb.ContentStreamOperations, _feef *_ge.PdfPageResources, _gcfa *_cac.PdfObjectDictionary) (*_ge.XObjectForm, error) {
	_faa := _ge.NewPdfPageResources()
	_ceaa, _ecbg := _bff, _bdb
	_eda := _fb.NewContentCreator()
	if _ceab.BorderSize > 0 {
		_abb(_eda, _ceab, _bff, _bdb)
	}
	if _ceab.DrawAlignmentReticle {
		_gac := _ceab
		_gac.BorderSize = 0.2
		_bgg(_eda, _gac, _bff, _bdb)
	}
	_eda.Add_BMC("\u0054\u0078")
	_eda.Add_q()
	_eda.Add_BT()
	_bff, _bdb = _ceab.applyRotation(_gcfa, _bff, _bdb, _eda)
	_gdda, _gcce, _bbc := _ceab.processDA(_dfdd, _ecdd, _feef, _faa, _eda)
	if _bbc != nil {
		return nil, _bbc
	}
	_decbd := _gdda.Font
	_bdac := _gdda.Size
	_bgdd := _cac.MakeName(_gdda.Name)
	_eac := _bdac == 0
	if _eac && _gcce {
		_bdac = _bdb * _ceab.AutoFontSizeFraction
	}
	_bga := _decbd.Encoder()
	if _bga == nil {
		_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_bga = _ac.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	if len(_bda) == 0 {
		return nil, nil
	}
	_dfe := _gec
	if _ceab.MarginLeft != nil {
		_dfe = *_ceab.MarginLeft
	}
	_cce := 0.0
	if _bga != nil {
		for _, _fdge := range _bda {
			_cdg, _cfe := _decbd.GetRuneMetrics(_fdge)
			if !_cfe {
				_g.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _fdge)
				continue
			}
			_cce += _cdg.Wx
		}
		_bda = string(_bga.Encode(_bda))
	}
	if _bdac == 0 || _eac && _cce > 0 && _dfe+_cce*_bdac/1000.0 > _bff {
		_bdac = 0.95 * 1000.0 * (_bff - _dfe) / _cce
	}
	_ddf := 1.0 * _bdac
	_cbge := 2.0
	{
		_gfe := _ddf
		if _eac && _cbge+_gfe > _bdb {
			_bdac = 0.95 * (_bdb - _cbge)
			_ddf = 1.0 * _bdac
			_gfe = _ddf
		}
		if _bdb > _gfe {
			_cbge = (_bdb - _gfe) / 2.0
			_cbge += 1.50
		}
	}
	_eda.Add_Tf(*_bgdd, _bdac)
	_eda.Add_Td(_dfe, _cbge)
	_eda.Add_Tj(*_cac.MakeString(_bda))
	_eda.Add_ET()
	_eda.Add_Q()
	_eda.Add_EMC()
	_fcd := _ge.NewXObjectForm()
	_fcd.Resources = _faa
	_fcd.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _ceaa, _ecbg})
	_fcd.SetContentStream(_eda.Bytes(), _daf())
	return _fcd, nil
}
func _edbe(_aedd LineAnnotationDef, _gddb string) ([]byte, *_ge.PdfRectangle, *_ge.PdfRectangle, error) {
	_gdfd := _ce.Line{X1: 0, Y1: 0, X2: _aedd.X2 - _aedd.X1, Y2: _aedd.Y2 - _aedd.Y1, LineColor: _aedd.LineColor, Opacity: _aedd.Opacity, LineWidth: _aedd.LineWidth, LineEndingStyle1: _aedd.LineEndingStyle1, LineEndingStyle2: _aedd.LineEndingStyle2}
	_acd, _bfge, _abdbg := _gdfd.Draw(_gddb)
	if _abdbg != nil {
		return nil, nil, nil, _abdbg
	}
	_aeab := &_ge.PdfRectangle{}
	_aeab.Llx = _aedd.X1 + _bfge.Llx
	_aeab.Lly = _aedd.Y1 + _bfge.Lly
	_aeab.Urx = _aedd.X1 + _bfge.Urx
	_aeab.Ury = _aedd.Y1 + _bfge.Ury
	return _acd, _bfge, _aeab, nil
}
func _ebgb(_bbe _ca.Image, _bdd string, _bagd *SignatureFieldOpts, _faac []float64, _dfdde *_fb.ContentCreator) (*_cac.PdfObjectName, *_ge.XObjectImage, error) {
	_abf, _gedd := _ge.DefaultImageHandler{}.NewImageFromGoImage(_bbe)
	if _gedd != nil {
		return nil, nil, _gedd
	}
	_egce, _gedd := _ge.NewXObjectImageFromImage(_abf, nil, _bagd.Encoder)
	if _gedd != nil {
		return nil, nil, _gedd
	}
	_aaab, _bbdb := float64(*_egce.Width), float64(*_egce.Height)
	_agbf := _faac[2] - _faac[0]
	_dfab := _faac[3] - _faac[1]
	if _bagd.AutoSize {
		_egd := _fg.Min(_agbf/_aaab, _dfab/_bbdb)
		_aaab *= _egd
		_bbdb *= _egd
		_faac[0] = _faac[0] + (_agbf / 2) - (_aaab / 2)
		_faac[1] = _faac[1] + (_dfab / 2) - (_bbdb / 2)
	}
	var _ecde *_cac.PdfObjectName
	if _ggba, _badc := _cac.GetName(_egce.Name); _badc {
		_ecde = _ggba
	} else {
		_ecde = _cac.MakeName(_bdd)
	}
	if _dfdde != nil {
		_dfdde.Add_q().Translate(_faac[0], _faac[1]).Scale(_aaab, _bbdb).Add_Do(*_ecde).Add_Q()
	} else {
		return nil, nil, _aeg.New("\u0043\u006f\u006e\u0074en\u0074\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u0075l\u006c")
	}
	return _ecde, _egce, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_cace FieldAppearance) Style() AppearanceStyle {
	if _cace._de != nil {
		return *_cace._de
	}
	_dd := _gec
	return AppearanceStyle{AutoFontSizeFraction: 0.65, CheckmarkRune: '✔', BorderSize: 0.0, BorderColor: _ge.NewPdfColorDeviceGray(0), FillColor: _ge.NewPdfColorDeviceGray(1), MultilineLineHeight: 1.2, MultilineVAlignMiddle: false, DrawAlignmentReticle: false, AllowMK: true, MarginLeft: &_dd}
}
func _ggc(_bbac *_ge.PdfAnnotationWidget, _gdb *_ge.PdfFieldText, _adba *_ge.PdfPageResources, _ddd AppearanceStyle) (*_cac.PdfObjectDictionary, error) {
	_dee := _ge.NewPdfPageResources()
	_cf, _acf := _cac.GetArray(_bbac.Rect)
	if !_acf {
		return nil, _aeg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_eg, _cc := _ge.NewPdfRectangle(*_cf)
	if _cc != nil {
		return nil, _cc
	}
	_dgb, _bfb := _eg.Width(), _eg.Height()
	_ffef, _gcf := _dgb, _bfb
	_agb := true
	_dec := _ge.NewXObjectForm()
	_dec.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _ffef, _gcf})
	if _bbac.AP != nil {
		if _fcc, _ddb := _cac.GetDict(_bbac.AP); _ddb && _fcc != nil {
			_fcf := _cac.TraceToDirectObject(_fcc.Get("\u004e"))
			switch _dff := _fcf.(type) {
			case *_cac.PdfObjectStream:
				_ecb, _ef := _cac.DecodeStream(_dff)
				if _ef != nil {
					_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u0020\u0073\u0074r\u0065\u0061m\u003a\u0020\u0025\u0076", _ef.Error())
					break
				}
				_ab, _ef := _fb.NewContentStreamParser(string(_ecb)).Parse()
				if _ef != nil {
					_g.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0075n\u0061\u0062l\u0065\u0020\u0070\u0061\u0072\u0073\u0065\u0020c\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061m\u003a\u0020\u0025\u0076", _ef.Error())
					break
				}
				_fgcd := _fb.NewContentStreamProcessor(*_ab)
				_fgcd.AddHandler(_fb.HandlerConditionEnumAllOperands, "", func(_cfg *_fb.ContentStreamOperation, _ged _fb.GraphicsState, _acg *_ge.PdfPageResources) error {
					if _cfg.Operand == "\u0054\u006a" || _cfg.Operand == "\u0054\u004a" {
						if len(_cfg.Params) == 1 {
							if _dcf, _gfb := _cac.GetString(_cfg.Params[0]); _gfb {
								_agb = _ae.TrimSpace(_dcf.Str()) == ""
							}
							return _fb.ErrEarlyExit
						}
						return nil
					}
					return nil
				})
				_fgcd.Process(_dee)
				if !_agb {
					if _bgb, _add := _cac.GetDict(_dff.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")); _add {
						_dee, _ef = _ge.NewPdfPageResourcesFromDict(_bgb)
						if _ef != nil {
							return nil, _ef
						}
					}
					if _eae, _ga := _cac.GetArray(_dff.Get("\u004d\u0061\u0074\u0072\u0069\u0078")); _ga {
						_dec.Matrix = _eae
					}
					_dec.SetContentStream(_ecb, _daf())
				}
			}
		}
	}
	if _agb {
		_aaa, _edf := _cac.GetDict(_bbac.MK)
		if _edf {
			_age, _ := _cac.GetDict(_bbac.BS)
			_bgbd := _ddd.applyAppearanceCharacteristics(_aaa, _age, nil)
			if _bgbd != nil {
				return nil, _bgbd
			}
		}
		_dffc, _dcd := _fb.NewContentStreamParser(_fde(_gdb.PdfField)).Parse()
		if _dcd != nil {
			return nil, _dcd
		}
		_gcc := _fb.NewContentCreator()
		if _ddd.BorderSize > 0 {
			_abb(_gcc, _ddd, _dgb, _bfb)
		}
		if _ddd.DrawAlignmentReticle {
			_gce := _ddd
			_gce.BorderSize = 0.2
			_bgg(_gcc, _gce, _dgb, _bfb)
		}
		_gcc.Add_BMC("\u0054\u0078")
		_gcc.Add_q()
		_dgb, _bfb = _ddd.applyRotation(_aaa, _dgb, _bfb, _gcc)
		_gcc.Add_BT()
		_cbg, _eeb, _dcd := _ddd.processDA(_gdb.PdfField, _dffc, _adba, _dee, _gcc)
		if _dcd != nil {
			return nil, _dcd
		}
		_fgcg := _cbg.Font
		_agf := _cbg.Size
		_ecd := _cac.MakeName(_cbg.Name)
		if _gdb.Flags().Has(_ge.FieldFlagMultiline) && _gdb.MaxLen != nil {
			_g.Log.Debug("\u004c\u006f\u006f\u006b\u0020\u0066\u006f\u0072\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072 \u004e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
			if _ecdf, _bfe, _fge := _gbdb(_bbac.PdfAnnotation.AP, _adba); _fge {
				_ecd = _ecdf
				_agf = _bfe
				_eeb = true
			}
		}
		_aege := _agf == 0
		if _aege && _eeb {
			_agf = _bfb * _ddd.AutoFontSizeFraction
		}
		_fee := _fgcg.Encoder()
		if _fee == nil {
			_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			_fee = _ac.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
		}
		_dce, _dcd := _fgcg.GetFontDescriptor()
		if _dcd != nil {
			_g.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
		}
		var _bdg string
		if _fbg, _fdd := _cac.GetString(_gdb.V); _fdd {
			_bdg = _fbg.Decoded()
		}
		if len(_bdg) == 0 {
			return nil, nil
		}
		_egf := []string{_bdg}
		_abd := false
		if _gdb.Flags().Has(_ge.FieldFlagMultiline) {
			_abd = true
			_bdg = _ae.Replace(_bdg, "\u000d\u000a", "\u000a", -1)
			_bdg = _ae.Replace(_bdg, "\u000d", "\u000a", -1)
			_egf = _ae.Split(_bdg, "\u000a")
		}
		_cea := make([]string, len(_egf))
		copy(_cea, _egf)
		_bfc := _ddd.MultilineLineHeight
		_bgf := 0.0
		_abdb := 0
		if _fee != nil {
			for _agf >= 0 {
				_fbf := make([]string, len(_egf))
				copy(_fbf, _egf)
				_fca := make([]string, len(_cea))
				copy(_fca, _cea)
				_bgf = 0.0
				_abdb = 0
				_fdg := len(_fbf)
				_edgd := 0
				for _edgd < _fdg {
					var _aab float64
					_deec := -1
					_gdg := _gec
					if _ddd.MarginLeft != nil {
						_gdg = *_ddd.MarginLeft
					}
					for _bac, _ccb := range _fbf[_edgd] {
						if _ccb == ' ' {
							_deec = _bac
						}
						_gedc, _acc := _fgcg.GetRuneMetrics(_ccb)
						if !_acc {
							_g.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _ccb)
							continue
						}
						_aab = _gdg
						_gdg += _gedc.Wx
						if _abd && !_aege && _agf*_gdg/1000.0 > _dgb {
							_gba := _bac
							_dga := _bac
							if _deec > 0 {
								_gba = _deec + 1
								_dga = _deec
							}
							_edfg := _fbf[_edgd][_gba:]
							_ead := _fca[_edgd][_gba:]
							if _edgd < len(_fbf)-1 {
								_fbf = append(_fbf[:_edgd+1], _fbf[_edgd:]...)
								_fbf[_edgd+1] = _edfg
								_fca = append(_fca[:_edgd+1], _fca[_edgd:]...)
								_fca[_edgd+1] = _ead
							} else {
								_fbf = append(_fbf, _edfg)
								_fca = append(_fca, _ead)
							}
							_fbf[_edgd] = _fbf[_edgd][0:_dga]
							_fca[_edgd] = _fca[_edgd][0:_dga]
							_fdg++
							_gdg = _aab
							break
						}
					}
					if _gdg > _bgf {
						_bgf = _gdg
					}
					_fbf[_edgd] = string(_fee.Encode(_fbf[_edgd]))
					if len(_fbf[_edgd]) > 0 {
						_abdb++
					}
					_edgd++
				}
				_aac := _agf
				if _abdb > 1 {
					_aac *= _bfc
				}
				_eadf := float64(_abdb) * _aac
				if _aege || _eadf <= _bfb {
					_egf = _fbf
					_cea = _fca
					break
				}
				_agf--
			}
		}
		_aae := _gec
		if _ddd.MarginLeft != nil {
			_aae = *_ddd.MarginLeft
		}
		if _agf == 0 || _aege && _bgf > 0 && _aae+_bgf*_agf/1000.0 > _dgb {
			_agf = 0.95 * 1000.0 * (_dgb - _aae) / _bgf
		}
		_fec := _ba
		{
			if _eab, _cbbe := _cac.GetIntVal(_gdb.Q); _cbbe {
				switch _eab {
				case 0:
					_fec = _ba
				case 1:
					_fec = _adf
				case 2:
					_fec = _cec
				default:
					_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0071\u0075\u0061\u0064\u0064\u0069\u006e\u0067\u003a\u0020%\u0064\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u006c\u0065ft\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074", _eab)
				}
			}
		}
		_fcfg := _agf
		if _abd && _abdb > 1 {
			_fcfg = _bfc * _agf
		}
		var _ded float64
		if _dce != nil {
			_ded, _dcd = _dce.GetCapHeight()
			if _dcd != nil {
				_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _dcd)
			}
		}
		if int(_ded) <= 0 {
			_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
			_ded = 1000
		}
		_agaa := _ded / 1000.0 * _agf
		_dad := 0.0
		{
			_aed := float64(_abdb) * _fcfg
			if _aege && _dad+_aed > _bfb {
				_agf = 0.95 * (_bfb - _dad) / float64(_abdb)
				_fcfg = _agf
				if _abd && _abdb > 1 {
					_fcfg = _bfc * _agf
				}
				_agaa = _ded / 1000.0 * _agf
				_aed = float64(_abdb) * _fcfg
			}
			if _bfb > _aed {
				if _abd {
					if _ddd.MultilineVAlignMiddle {
						_fdgc := (_bfb - (_aed + _agaa)) / 2.0
						_bag := _fdgc + _aed + _agaa - _fcfg
						_dad = _bag
						if _abdb > 1 {
							_dad = _dad + (_aed / _agf * float64(_abdb)) - _fcfg - _agaa
						}
						if _dad < _aed {
							_dad = (_bfb - _agaa) / 2.0
						}
					} else {
						_dad = _bfb - _fcfg
						if _dad > _agf {
							_dgbb := 0.0
							if _abd && _ddd.MultilineLineHeight > 1 && _abdb > 1 {
								_dgbb = _ddd.MultilineLineHeight - 1
							}
							_dad -= _agf * (0.5 - _dgbb)
						}
					}
				} else {
					_dad = (_bfb - _agaa) / 2.0
				}
			}
		}
		if _ddd.TextColor != nil {
			_dcfd := _ddd.TextColor
			_decb, _aef := _dcfd.(*_ge.PdfColorDeviceRGB)
			if !_aef {
				_decb = _ge.NewPdfColorDeviceRGB(0, 0, 0)
			}
			_gcc.Add_rg(_decb.R(), _decb.G(), _decb.B())
		} else {
			for _, _fdga := range *_dffc {
				if _fdga.Operand == "\u0072\u0067" || _fdga.Operand == "\u0067" {
					_gcc.AddOperand(*_fdga)
				}
			}
		}
		_gcc.Add_Tf(*_ecd, _agf)
		_gcc.Add_Td(_aae, _dad)
		_ced := _aae
		_fbbd := _aae
		for _ega, _dbd := range _egf {
			_bcc := 0.0
			for _, _cabb := range _cea[_ega] {
				_dbc, _ade := _fgcg.GetRuneMetrics(_cabb)
				if !_ade {
					continue
				}
				_bcc += _dbc.Wx
			}
			_ggfc := _bcc / 1000.0 * _agf
			_fab := _dgb - _ggfc
			var _abde float64
			switch _fec {
			case _ba:
				_abde = _ced
			case _adf:
				_abde = _fab / 2
			case _cec:
				_abde = _fab
			}
			_aae = _abde - _fbbd
			if _aae > 0.0 {
				_gcc.Add_Td(_aae, 0)
			}
			_fbbd = _abde
			_gcc.Add_Tj(*_cac.MakeString(_dbd))
			if _ega < len(_egf)-1 {
				_gcc.Add_Td(0, -_agf*_bfc)
			}
		}
		_gcc.Add_ET()
		_gcc.Add_Q()
		_gcc.Add_EMC()
		_dec.SetContentStream(_gcc.Bytes(), _daf())
	}
	_dec.Resources = _dee
	_abg := _cac.MakeDict()
	_abg.Set("\u004e", _dec.ToPdfObject())
	return _abg, nil
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_bcge ImageFieldAppearance) GenerateAppearanceDict(form *_ge.PdfAcroForm, field *_ge.PdfField, wa *_ge.PdfAnnotationWidget) (*_cac.PdfObjectDictionary, error) {
	_, _ebag := field.GetContext().(*_ge.PdfFieldButton)
	if !_ebag {
		_g.Log.Trace("C\u006f\u0075\u006c\u0064\u0020\u006fn\u006c\u0079\u0020\u0068\u0061\u006ed\u006c\u0065\u0020\u0062\u0075\u0074\u0074o\u006e\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067")
		return nil, nil
	}
	_bfd, _ffd := _cac.GetDict(wa.AP)
	if _ffd && _bcge.OnlyIfMissing {
		_g.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _bfd, nil
	}
	if form.DR == nil {
		form.DR = _ge.NewPdfPageResources()
	}
	switch _fcfe := field.GetContext().(type) {
	case *_ge.PdfFieldButton:
		if _fcfe.IsPush() {
			_adfde, _bfcc := _cdff(_fcfe, wa, _bcge.Style())
			if _bfcc != nil {
				return nil, _bfcc
			}
			return _adfde, nil
		}
	}
	return nil, nil
}

// InkAnnotationDef holds base information for constructing an ink annotation.
type InkAnnotationDef struct {

	// Paths is the array of stroked paths which compose the annotation.
	Paths []_ce.Path

	// Color is the color of the line. Default to black.
	Color *_ge.PdfColorDeviceRGB

	// LineWidth is the width of the line.
	LineWidth float64
}

// ImageFieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for attaching an image to a button field.
type ImageFieldAppearance struct {
	OnlyIfMissing bool
	_decd         *AppearanceStyle
}

func _gc(_ff CircleAnnotationDef) (*_cac.PdfObjectDictionary, *_ge.PdfRectangle, error) {
	_gg := _ge.NewXObjectForm()
	_gg.Resources = _ge.NewPdfPageResources()
	_bec := ""
	if _ff.Opacity < 1.0 {
		_fa := _cac.MakeDict()
		_fa.Set("\u0063\u0061", _cac.MakeFloat(_ff.Opacity))
		_fa.Set("\u0043\u0041", _cac.MakeFloat(_ff.Opacity))
		_cab := _gg.Resources.AddExtGState("\u0067\u0073\u0031", _fa)
		if _cab != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _cab
		}
		_bec = "\u0067\u0073\u0031"
	}
	_aa, _ggf, _db, _df := _fbb(_ff, _bec)
	if _df != nil {
		return nil, nil, _df
	}
	_df = _gg.SetContentStream(_aa, nil)
	if _df != nil {
		return nil, nil, _df
	}
	_gg.BBox = _ggf.ToPdfObject()
	_cbf := _cac.MakeDict()
	_cbf.Set("\u004e", _gg.ToPdfObject())
	return _cbf, _db, nil
}

// AppearanceStyle defines style parameters for appearance stream generation.
type AppearanceStyle struct {

	// How much of Rect height to fill when autosizing text.
	AutoFontSizeFraction float64

	// CheckmarkRune is a rune used for check mark in checkboxes (for ZapfDingbats font).
	CheckmarkRune rune
	BorderSize    float64
	BorderColor   _ge.PdfColor
	FillColor     _ge.PdfColor

	// Multiplier for lineheight for multi line text.
	MultilineLineHeight   float64
	MultilineVAlignMiddle bool

	// Visual guide checking alignment of field contents (debugging).
	DrawAlignmentReticle bool

	// Allow field MK appearance characteristics to override style settings.
	AllowMK bool

	// Fonts holds appearance styles for fonts.
	Fonts *AppearanceFontStyle

	// MarginLeft represents the amount of space to leave on the left side of
	// the form field bounding box when generating appearances (default: 2.0).
	MarginLeft *float64
	TextColor  _ge.PdfColor
}

// FormSubmitActionOptions holds options for creating a form submit button.
type FormSubmitActionOptions struct {

	// Rectangle holds the button position, size, and color.
	Rectangle _ce.Rectangle

	// Url specifies the URL where the fieds will be submitted.
	Url string

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _ge.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_ge.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be submitted.
	// This list may contain indirect object to fields or field names.
	Fields *_cac.PdfObjectArray

	// IsExclusionList specifies that the fields contain in `Fields` array would not be submitted.
	IsExclusionList bool

	// IncludeEmptyFields specifies if all fields would be submitted even though it's value is empty.
	IncludeEmptyFields bool

	// SubmitAsPDF specifies that the document shall be submitted as PDF.
	// If set then all the other flags shall be ignored.
	SubmitAsPDF bool
}

// SetStyle applies appearance `style` to `fa`.
func (_bccc *ImageFieldAppearance) SetStyle(style AppearanceStyle) { _bccc._decd = &style }

// SignatureFieldOpts represents a set of options used to configure
// an appearance widget dictionary.
type SignatureFieldOpts struct {

	// Rect represents the area the signature annotation is displayed on.
	Rect []float64

	// AutoSize specifies if the content of the appearance should be
	// scaled to fit in the annotation rectangle.
	AutoSize bool

	// Font specifies the font of the text content.
	Font *_ge.PdfFont

	// FontSize specifies the size of the text content.
	FontSize float64

	// LineHeight specifies the height of a line of text in the appearance annotation.
	LineHeight float64

	// TextColor represents the color of the text content displayed.
	TextColor _ge.PdfColor

	// FillColor represents the background color of the appearance annotation area.
	FillColor _ge.PdfColor

	// BorderSize represents border size of the appearance annotation area.
	BorderSize float64

	// BorderColor represents the border color of the appearance annotation area.
	BorderColor _ge.PdfColor

	// WatermarkImage specifies the image used as a watermark that will be rendered
	// behind the signature.
	WatermarkImage _ca.Image

	// Image represents the image used for the signature appearance.
	Image _ca.Image

	// Encoder specifies the image encoder used for image signature. Defaults to flate encoder.
	Encoder _cac.StreamEncoder

	// ImagePosition specifies the image location relative to the text signature.
	ImagePosition SignatureImagePosition
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_bfec FieldAppearance) WrapContentStream(page *_ge.PdfPage) error {
	_fbdf, _afgb := page.GetAllContentStreams()
	if _afgb != nil {
		return _afgb
	}
	_bce := _fb.NewContentStreamParser(_fbdf)
	_fcce, _afgb := _bce.Parse()
	if _afgb != nil {
		return _afgb
	}
	_fcce.WrapIfNeeded()
	_deb := []string{_fcce.String()}
	return page.SetContentStreams(_deb, _daf())
}
func _edga(_afeg *_ge.PdfAcroForm, _cbfb *_ge.PdfAnnotationWidget, _cee *_ge.PdfFieldChoice, _cfgf AppearanceStyle) (*_cac.PdfObjectDictionary, error) {
	_gdc, _efdc := _cac.GetArray(_cbfb.Rect)
	if !_efdc {
		return nil, _aeg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_eaab, _edgdb := _ge.NewPdfRectangle(*_gdc)
	if _edgdb != nil {
		return nil, _edgdb
	}
	_edcc, _bcaa := _eaab.Width(), _eaab.Height()
	_g.Log.Debug("\u0043\u0068\u006f\u0069\u0063\u0065\u002c\u0020\u0077\u0061\u0020\u0042S\u003a\u0020\u0025\u0076", _cbfb.BS)
	_addg, _edgdb := _fb.NewContentStreamParser(_fde(_cee.PdfField)).Parse()
	if _edgdb != nil {
		return nil, _edgdb
	}
	_cead, _gag := _cac.GetDict(_cbfb.MK)
	if _gag {
		_gffc, _ := _cac.GetDict(_cbfb.BS)
		_fea := _cfgf.applyAppearanceCharacteristics(_cead, _gffc, nil)
		if _fea != nil {
			return nil, _fea
		}
	}
	_fgcb := _cac.MakeDict()
	for _, _gbg := range _cee.Opt.Elements() {
		if _agg, _afeb := _cac.GetArray(_gbg); _afeb && _agg.Len() == 2 {
			_gbg = _agg.Get(1)
		}
		var _ddad string
		if _aede, _deeg := _cac.GetString(_gbg); _deeg {
			_ddad = _aede.Decoded()
		} else if _cfce, _ceag := _cac.GetName(_gbg); _ceag {
			_ddad = _cfce.String()
		} else {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004f\u0070\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u002f\u0073\u0074\u0072\u0069\u006e\u0067 \u002d\u0020\u0025\u0054", _gbg)
			return nil, _aeg.New("\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u002f\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_ddad) > 0 {
			_gdd, _cgg := _ccd(_cee.PdfField, _edcc, _bcaa, _ddad, _cfgf, _addg, _afeg.DR, _cead)
			if _cgg != nil {
				return nil, _cgg
			}
			_fgcb.Set(*_cac.MakeName(_ddad), _gdd.ToPdfObject())
		}
	}
	_cabg := _cac.MakeDict()
	_cabg.Set("\u004e", _fgcb)
	return _cabg, nil
}
func (_gcg *AppearanceStyle) processDA(_abc *_ge.PdfField, _aefg *_fb.ContentStreamOperations, _befc, _bed *_ge.PdfPageResources, _cbc *_fb.ContentCreator) (*AppearanceFont, bool, error) {
	var _aeeb *AppearanceFont
	var _dbf bool
	if _gcg.Fonts != nil {
		if _gcg.Fonts.Fallback != nil {
			_aeeb = _gcg.Fonts.Fallback
		}
		if _daef := _gcg.Fonts.FieldFallbacks; _daef != nil {
			if _bgc, _ace := _daef[_abc.PartialName()]; _ace {
				_aeeb = _bgc
			} else if _aggg, _dgdc := _abc.FullName(); _dgdc == nil {
				if _gaf, _adfd := _daef[_aggg]; _adfd {
					_aeeb = _gaf
				}
			}
		}
		if _aeeb != nil {
			_aeeb.fillName()
		}
		_dbf = _gcg.Fonts.ForceReplace
	}
	var _gfaf string
	var _fcdg float64
	var _caba bool
	if _aefg != nil {
		for _, _fbd := range *_aefg {
			if _fbd.Operand == "\u0054\u0066" && len(_fbd.Params) == 2 {
				if _gcgd, _cca := _cac.GetNameVal(_fbd.Params[0]); _cca {
					_gfaf = _gcgd
				}
				if _fbee, _cbga := _cac.GetNumberAsFloat(_fbd.Params[1]); _cbga == nil {
					_fcdg = _fbee
				}
				_caba = true
				continue
			}
			_cbc.AddOperand(*_fbd)
		}
	}
	var _cggb *AppearanceFont
	var _bcf _cac.PdfObject
	if _dbf && _aeeb != nil {
		_cggb = _aeeb
	} else {
		if _befc != nil && _gfaf != "" {
			if _ccag, _bcb := _befc.GetFontByName(*_cac.MakeName(_gfaf)); _bcb {
				if _cdbc, _geef := _ge.NewPdfFontFromPdfObject(_ccag); _geef == nil {
					_bcf = _ccag
					_cggb = &AppearanceFont{Name: _gfaf, Font: _cdbc, Size: _fcdg}
				} else {
					_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006fa\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _geef)
				}
			}
		}
		if _cggb == nil && _aeeb != nil {
			_cggb = _aeeb
		}
		if _cggb == nil {
			_fac, _addea := _ge.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
			if _addea != nil {
				return nil, false, _addea
			}
			_cggb = &AppearanceFont{Name: "\u0048\u0065\u006c\u0076", Font: _fac, Size: _fcdg}
		}
	}
	if _cggb.Size <= 0 && _gcg.Fonts != nil && _gcg.Fonts.FallbackSize > 0 {
		_cggb.Size = _gcg.Fonts.FallbackSize
	}
	_eagd := *_cac.MakeName(_cggb.Name)
	if _bcf == nil {
		_bcf = _cggb.Font.ToPdfObject()
	}
	if _befc != nil && !_befc.HasFontByName(_eagd) {
		_befc.SetFontByName(_eagd, _bcf)
	}
	if _bed != nil && !_bed.HasFontByName(_eagd) {
		_bed.SetFontByName(_eagd, _bcf)
	}
	return _cggb, _caba, nil
}
func _daf() _cac.StreamEncoder { return _cac.NewFlateEncoder() }
func _eaae(_ege *_ge.PdfPage, _aceb _ce.Rectangle, _efed string, _fddbb string, _fffc _ge.PdfColor, _gcga *_ge.PdfFont, _daeg *float64, _eccf _cac.PdfObject) (*_ge.PdfFieldButton, error) {
	_abba, _gcgdf := _aceb.X, _aceb.Y
	_ffg := _aceb.Width
	_cage := _aceb.Height
	if _aceb.FillColor == nil {
		_aceb.FillColor = _ge.NewPdfColorDeviceGray(0.7)
	}
	if _fffc == nil {
		_fffc = _ge.NewPdfColorDeviceGray(0)
	}
	if _gcga == nil {
		_fcg, _bbf := _ge.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
		if _bbf != nil {
			return nil, _bbf
		}
		_gcga = _fcg
	}
	_eebf := _ge.NewPdfField()
	_aecb := &_ge.PdfFieldButton{}
	_eebf.SetContext(_aecb)
	_aecb.PdfField = _eebf
	_aecb.T = _cac.MakeString(_efed)
	_aecb.SetType(_ge.ButtonTypePush)
	_aecb.V = _cac.MakeName("\u004f\u0066\u0066")
	_aecb.Ff = _cac.MakeInteger(4)
	_ece := _cac.MakeDict()
	_ece.Set(*_cac.MakeName("\u0043\u0041"), _cac.MakeString(_fddbb))
	_fba, _dfdc := _gcga.GetFontDescriptor()
	if _dfdc != nil {
		return nil, _dfdc
	}
	_gdaa := _cac.MakeName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	_fggd := 12.0
	if _fba != nil && _fba.FontName != nil {
		_gdaa, _ = _cac.GetName(_fba.FontName)
	}
	if _daeg != nil {
		_fggd = *_daeg
	}
	_cbe := _fb.NewContentCreator()
	_cbe.Add_q()
	_cbe.SetNonStrokingColor(_aceb.FillColor)
	_cbe.Add_re(0, 0, _ffg, _cage)
	_cbe.Add_f()
	_cbe.Add_Q()
	_cbe.Add_q()
	_cbe.Add_BT()
	_dcee := 0.0
	for _, _beab := range _fddbb {
		_eccb, _ecba := _gcga.GetRuneMetrics(_beab)
		if !_ecba {
			_g.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _beab)
			continue
		}
		_dcee += _eccb.Wx
	}
	_dcee = _dcee / 1000.0 * _fggd
	var _afgfa float64
	if _fba != nil {
		_afgfa, _dfdc = _fba.GetCapHeight()
		if _dfdc != nil {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _dfdc)
		}
	}
	if int(_afgfa) <= 0 {
		_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_afgfa = 1000
	}
	_cdc := _afgfa / 1000.0 * _fggd
	_gecag := (_cage - _cdc) / 2.0
	_cffa := (_ffg - _dcee) / 2.0
	_cbe.Add_Tf(*_gdaa, _fggd)
	_cbe.SetNonStrokingColor(_fffc)
	_cbe.Add_Td(_cffa, _gecag)
	_cbe.Add_Tj(*_cac.MakeString(_fddbb))
	_cbe.Add_ET()
	_cbe.Add_Q()
	_aecf := _ge.NewXObjectForm()
	_aecf.SetContentStream(_cbe.Bytes(), _cac.NewRawEncoder())
	_aecf.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _ffg, _cage})
	_aecf.Resources = _ge.NewPdfPageResources()
	_aecf.Resources.SetFontByName(*_gdaa, _gcga.ToPdfObject())
	_fcaf := _cac.MakeDict()
	_fcaf.Set("\u004e", _aecf.ToPdfObject())
	_eegc := _ge.NewPdfAnnotationWidget()
	_eegc.Rect = _cac.MakeArrayFromFloats([]float64{_abba, _gcgdf, _abba + _ffg, _gcgdf + _cage})
	_eegc.P = _ege.ToPdfObject()
	_eegc.F = _cac.MakeInteger(4)
	_eegc.Parent = _aecb.ToPdfObject()
	_eegc.A = _eccf
	_eegc.MK = _ece
	_eegc.AP = _fcaf
	_aecb.Annotations = append(_aecb.Annotations, _eegc)
	return _aecb, nil
}

// CreateLineAnnotation creates a line annotation object that can be added to page PDF annotations.
func CreateLineAnnotation(lineDef LineAnnotationDef) (*_ge.PdfAnnotation, error) {
	_adgg := _ge.NewPdfAnnotationLine()
	_adgg.L = _cac.MakeArrayFromFloats([]float64{lineDef.X1, lineDef.Y1, lineDef.X2, lineDef.Y2})
	_bgad := _cac.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle1 == _ce.LineEndingStyleArrow {
		_bgad = _cac.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_bbed := _cac.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle2 == _ce.LineEndingStyleArrow {
		_bbed = _cac.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_adgg.LE = _cac.MakeArray(_bgad, _bbed)
	if lineDef.Opacity < 1.0 {
		_adgg.CA = _cac.MakeFloat(lineDef.Opacity)
	}
	_gfbe, _bbfd, _dacf := lineDef.LineColor.R(), lineDef.LineColor.G(), lineDef.LineColor.B()
	_adgg.IC = _cac.MakeArrayFromFloats([]float64{_gfbe, _bbfd, _dacf})
	_adgg.C = _cac.MakeArrayFromFloats([]float64{_gfbe, _bbfd, _dacf})
	_dfda := _ge.NewBorderStyle()
	_dfda.SetBorderWidth(lineDef.LineWidth)
	_adgg.BS = _dfda.ToPdfObject()
	_beff, _deff, _cegb := _cecg(lineDef)
	if _cegb != nil {
		return nil, _cegb
	}
	_adgg.AP = _beff
	_adgg.Rect = _cac.MakeArrayFromFloats([]float64{_deff.Llx, _deff.Lly, _deff.Urx, _deff.Ury})
	return _adgg.PdfAnnotation, nil
}
func _cecg(_gbc LineAnnotationDef) (*_cac.PdfObjectDictionary, *_ge.PdfRectangle, error) {
	_aefa := _ge.NewXObjectForm()
	_aefa.Resources = _ge.NewPdfPageResources()
	_beebd := ""
	if _gbc.Opacity < 1.0 {
		_aff := _cac.MakeDict()
		_aff.Set("\u0063\u0061", _cac.MakeFloat(_gbc.Opacity))
		_dadd := _aefa.Resources.AddExtGState("\u0067\u0073\u0031", _aff)
		if _dadd != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _dadd
		}
		_beebd = "\u0067\u0073\u0031"
	}
	_cebe, _cadff, _ffda, _faae := _edbe(_gbc, _beebd)
	if _faae != nil {
		return nil, nil, _faae
	}
	_faae = _aefa.SetContentStream(_cebe, nil)
	if _faae != nil {
		return nil, nil, _faae
	}
	_aefa.BBox = _cadff.ToPdfObject()
	_fdgbg := _cac.MakeDict()
	_fdgbg.Set("\u004e", _aefa.ToPdfObject())
	return _fdgbg, _ffda, nil
}
func _ebec(_fbgf *InkAnnotationDef) (*_cac.PdfObjectDictionary, *_ge.PdfRectangle, error) {
	_afae := _ge.NewXObjectForm()
	_dded, _afbf, _adfe := _gde(_fbgf)
	if _adfe != nil {
		return nil, nil, _adfe
	}
	_adfe = _afae.SetContentStream(_dded, nil)
	if _adfe != nil {
		return nil, nil, _adfe
	}
	_afae.BBox = _afbf.ToPdfObject()
	_afae.Resources = _ge.NewPdfPageResources()
	_afae.Resources.ProcSet = _cac.MakeArray(_cac.MakeName("\u0050\u0044\u0046"))
	_bgac := _cac.MakeDict()
	_bgac.Set("\u004e", _afae.ToPdfObject())
	return _bgac, _afbf, nil
}
func (_cde *AppearanceStyle) applyRotation(_faab *_cac.PdfObjectDictionary, _dae, _fddd float64, _dgdd *_fb.ContentCreator) (float64, float64) {
	if !_cde.AllowMK {
		return _dae, _fddd
	}
	if _faab == nil {
		return _dae, _fddd
	}
	_gbb, _ := _cac.GetNumberAsFloat(_faab.Get("\u0052"))
	if _gbb == 0 {
		return _dae, _fddd
	}
	_eagg := -_gbb
	_dca := _ce.Path{Points: []_ce.Point{_ce.NewPoint(0, 0).Rotate(_eagg), _ce.NewPoint(_dae, 0).Rotate(_eagg), _ce.NewPoint(0, _fddd).Rotate(_eagg), _ce.NewPoint(_dae, _fddd).Rotate(_eagg)}}.GetBoundingBox()
	_dgdd.RotateDeg(_gbb)
	_dgdd.Translate(_dca.X, _dca.Y)
	return _dca.Width, _dca.Height
}

// FormResetActionOptions holds options for creating a form reset button.
type FormResetActionOptions struct {

	// Rectangle holds the button position, size, and color.
	Rectangle _ce.Rectangle

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _ge.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_ge.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be resetted.
	// This list may contain indirect object to fields or field names.
	Fields *_cac.PdfObjectArray

	// IsExclusionList specifies that the fields in the `Fields` array would be excluded form reset process.
	IsExclusionList bool
}

// CreateCircleAnnotation creates a circle/ellipse annotation object with appearance stream that can be added to
// page PDF annotations.
func CreateCircleAnnotation(circDef CircleAnnotationDef) (*_ge.PdfAnnotation, error) {
	_be := _ge.NewPdfAnnotationCircle()
	if circDef.BorderEnabled {
		_bb, _bd, _gf := circDef.BorderColor.R(), circDef.BorderColor.G(), circDef.BorderColor.B()
		_be.C = _cac.MakeArrayFromFloats([]float64{_bb, _bd, _gf})
		_af := _ge.NewBorderStyle()
		_af.SetBorderWidth(circDef.BorderWidth)
		_be.BS = _af.ToPdfObject()
	}
	if circDef.FillEnabled {
		_d, _ag, _fe := circDef.FillColor.R(), circDef.FillColor.G(), circDef.FillColor.B()
		_be.IC = _cac.MakeArrayFromFloats([]float64{_d, _ag, _fe})
	} else {
		_be.IC = _cac.MakeArrayFromIntegers([]int{})
	}
	if circDef.Opacity < 1.0 {
		_be.CA = _cac.MakeFloat(circDef.Opacity)
	}
	_e, _afe, _bf := _gc(circDef)
	if _bf != nil {
		return nil, _bf
	}
	_be.AP = _e
	_be.Rect = _cac.MakeArrayFromFloats([]float64{_afe.Llx, _afe.Lly, _afe.Urx, _afe.Ury})
	return _be.PdfAnnotation, nil
}

// NewFormSubmitButtonField would create a submit button in specified page according to the parameter in `FormSubmitActionOptions`.
func NewFormSubmitButtonField(page *_ge.PdfPage, opt FormSubmitActionOptions) (*_ge.PdfFieldButton, error) {
	_bcd := int64(_ageb)
	if opt.IsExclusionList {
		_bcd |= _gdge
	}
	if opt.IncludeEmptyFields {
		_bcd |= _aefdg
	}
	if opt.SubmitAsPDF {
		_bcd |= _gagg
	}
	_aec := _ge.NewPdfActionSubmitForm()
	_aec.Flags = _cac.MakeInteger(_bcd)
	_aec.F = _ge.NewPdfFilespec()
	if opt.Fields != nil {
		_aec.Fields = opt.Fields
	}
	_aec.F.F = _cac.MakeString(opt.Url)
	_aec.F.FS = _cac.MakeName("\u0055\u0052\u004c")
	_eccc, _bddc := _eaae(page, opt.Rectangle, "\u0062t\u006e\u0053\u0075\u0062\u006d\u0069t", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _aec.ToPdfObject())
	if _bddc != nil {
		return nil, _bddc
	}
	return _eccc, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_aggf ImageFieldAppearance) Style() AppearanceStyle {
	if _aggf._decd != nil {
		return *_aggf._decd
	}
	return AppearanceStyle{BorderSize: 0.0, BorderColor: _ge.NewPdfColorDeviceGray(0), FillColor: _ge.NewPdfColorDeviceGray(1), DrawAlignmentReticle: false}
}

// RectangleAnnotationDef is a rectangle defined with a specified Width and Height and a lower left corner at (X,Y).
// The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified).
type RectangleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_ge.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_ge.PdfColorDeviceRGB
	Opacity       float64
}

// CheckboxFieldOptions defines optional parameters for a checkbox field a form.
type CheckboxFieldOptions struct{ Checked bool }

func _abb(_fcba *_fb.ContentCreator, _feg AppearanceStyle, _efag, _dgd float64) {
	_fcba.Add_q().Add_re(0, 0, _efag, _dgd).Add_w(_feg.BorderSize).SetStrokingColor(_feg.BorderColor).SetNonStrokingColor(_feg.FillColor).Add_B().Add_Q()
}
func _bfad(_cgc *_ge.PdfAnnotationWidget, _aeac *_ge.PdfFieldButton, _afg *_ge.PdfPageResources, _edfd AppearanceStyle) (*_cac.PdfObjectDictionary, error) {
	_dbde, _fgb := _cac.GetArray(_cgc.Rect)
	if !_fgb {
		return nil, _aeg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_egfb, _dea := _ge.NewPdfRectangle(*_dbde)
	if _dea != nil {
		return nil, _dea
	}
	_bee, _ecg := _egfb.Width(), _egfb.Height()
	_gfg, _ceg := _bee, _ecg
	_g.Log.Debug("\u0043\u0068\u0065\u0063kb\u006f\u0078\u002c\u0020\u0077\u0061\u0020\u0042\u0053\u003a\u0020\u0025\u0076", _cgc.BS)
	_eca, _dea := _ge.NewStandard14Font("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
	if _dea != nil {
		return nil, _dea
	}
	_gbd, _baf := _cac.GetDict(_cgc.MK)
	if _baf {
		_bde, _ := _cac.GetDict(_cgc.BS)
		_gee := _edfd.applyAppearanceCharacteristics(_gbd, _bde, _eca)
		if _gee != nil {
			return nil, _gee
		}
	}
	_aca := _ge.NewXObjectForm()
	{
		_daa := _fb.NewContentCreator()
		if _edfd.BorderSize > 0 {
			_abb(_daa, _edfd, _bee, _ecg)
		}
		if _edfd.DrawAlignmentReticle {
			_gfd := _edfd
			_gfd.BorderSize = 0.2
			_bgg(_daa, _gfd, _bee, _ecg)
		}
		_bee, _ecg = _edfd.applyRotation(_gbd, _bee, _ecg, _daa)
		_bae := _edfd.AutoFontSizeFraction * _ecg
		_gdbe, _dcdd := _eca.GetRuneMetrics(_edfd.CheckmarkRune)
		if !_dcdd {
			return nil, _aeg.New("\u0067l\u0079p\u0068\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_edc := _eca.Encoder()
		_dgc := _edc.Encode(string(_edfd.CheckmarkRune))
		_fbe := _gdbe.Wx * _bae / 1000.0
		_acce := 705.0
		_fgga := _acce / 1000.0 * _bae
		_ddbc := _gec
		if _edfd.MarginLeft != nil {
			_ddbc = *_edfd.MarginLeft
		}
		_cag := 1.0
		if _fbe < _bee {
			_ddbc = (_bee - _fbe) / 2.0
		}
		if _fgga < _ecg {
			_cag = (_ecg - _fgga) / 2.0
		}
		_daa.Add_q().Add_g(0).Add_BT().Add_Tf("\u005a\u0061\u0044\u0062", _bae).Add_Td(_ddbc, _cag).Add_Tj(*_cac.MakeStringFromBytes(_dgc)).Add_ET().Add_Q()
		_aca.Resources = _ge.NewPdfPageResources()
		_aca.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _eca.ToPdfObject())
		_aca.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _gfg, _ceg})
		_aca.SetContentStream(_daa.Bytes(), _daf())
	}
	_dgge := _ge.NewXObjectForm()
	{
		_dbb := _fb.NewContentCreator()
		if _edfd.BorderSize > 0 {
			_abb(_dbb, _edfd, _bee, _ecg)
		}
		_dgge.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _gfg, _ceg})
		_dgge.SetContentStream(_dbb.Bytes(), _daf())
	}
	_fggb := _cac.PdfObjectName("\u0059\u0065\u0073")
	_fdgb, _baf := _cac.GetDict(_cgc.PdfAnnotation.AP)
	if _baf && _fdgb != nil {
		_gdba := _cac.TraceToDirectObject(_fdgb.Get("\u004e"))
		switch _dfce := _gdba.(type) {
		case *_cac.PdfObjectDictionary:
			_bbd := _dfce.Keys()
			for _, _bbb := range _bbd {
				if _bbb != "\u004f\u0066\u0066" {
					_fggb = _bbb
				}
			}
		}
	}
	_afcc := _cac.MakeDict()
	_afcc.Set("\u004f\u0066\u0066", _dgge.ToPdfObject())
	_afcc.Set(_fggb, _aca.ToPdfObject())
	_gab := _cac.MakeDict()
	_gab.Set("\u004e", _afcc)
	return _gab, nil
}

const (
	_ba  quadding = 0
	_adf quadding = 1
	_cec quadding = 2
	_gec float64  = 2.0
)

func _fbaa(_gbed []_ce.Point) (_cdcf []_ce.Point, _gffe []_ce.Point, _adeg error) {
	_aegf := len(_gbed) - 1
	if len(_gbed) < 1 {
		return nil, nil, _aeg.New("\u0041\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u0070\u006f\u0069\u006e\u0074s \u0072e\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0020\u0063\u0075\u0072\u0076\u0065\u0020\u0063\u006f\u006e\u0074r\u006f\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if _aegf == 1 {
		_aba := _ce.Point{X: (2*_gbed[0].X + _gbed[1].X) / 3, Y: (2*_gbed[0].Y + _gbed[1].Y) / 3}
		_cdcf = append(_cdcf, _aba)
		_gffe = append(_gffe, _ce.Point{X: 2*_aba.X - _gbed[0].X, Y: 2*_aba.Y - _gbed[0].Y})
		return _cdcf, _gffe, nil
	}
	_cegf := make([]float64, _aegf)
	for _dbee := 1; _dbee < _aegf-1; _dbee++ {
		_cegf[_dbee] = 4*_gbed[_dbee].X + 2*_gbed[_dbee+1].X
	}
	_cegf[0] = _gbed[0].X + 2*_gbed[1].X
	_cegf[_aegf-1] = (8*_gbed[_aegf-1].X + _gbed[_aegf].X) / 2.0
	_cegfc := _eea(_cegf)
	for _bfff := 1; _bfff < _aegf-1; _bfff++ {
		_cegf[_bfff] = 4*_gbed[_bfff].Y + 2*_gbed[_bfff+1].Y
	}
	_cegf[0] = _gbed[0].Y + 2*_gbed[1].Y
	_cegf[_aegf-1] = (8*_gbed[_aegf-1].Y + _gbed[_aegf].Y) / 2.0
	_aegc := _eea(_cegf)
	_cdcf = make([]_ce.Point, _aegf)
	_gffe = make([]_ce.Point, _aegf)
	for _deee := 0; _deee < _aegf; _deee++ {
		_cdcf[_deee] = _ce.Point{X: _cegfc[_deee], Y: _aegc[_deee]}
		if _deee < _aegf-1 {
			_gffe[_deee] = _ce.Point{X: 2*_gbed[_deee+1].X - _cegfc[_deee+1], Y: 2*_gbed[_deee+1].Y - _aegc[_deee+1]}
		} else {
			_gffe[_deee] = _ce.Point{X: (_gbed[_aegf].X + _cegfc[_aegf-1]) / 2, Y: (_gbed[_aegf].Y + _aegc[_aegf-1]) / 2}
		}
	}
	return _cdcf, _gffe, nil
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_adcd ImageFieldAppearance) WrapContentStream(page *_ge.PdfPage) error {
	_ffgd, _bfbeb := page.GetAllContentStreams()
	if _bfbeb != nil {
		return _bfbeb
	}
	_dgag := _fb.NewContentStreamParser(_ffgd)
	_eaca, _bfbeb := _dgag.Parse()
	if _bfbeb != nil {
		return _bfbeb
	}
	_eaca.WrapIfNeeded()
	_gegb := []string{_eaca.String()}
	return page.SetContentStreams(_gegb, _daf())
}

// NewSignatureField returns a new signature field with a visible appearance
// containing the specified signature lines and styled according to the
// specified options.
func NewSignatureField(signature *_ge.PdfSignature, lines []*SignatureLine, opts *SignatureFieldOpts) (*_ge.PdfFieldSignature, error) {
	if signature == nil {
		return nil, _aeg.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_cbgaa, _gbdd := _gcea(lines, opts)
	if _gbdd != nil {
		return nil, _gbdd
	}
	_dafg := _ge.NewPdfFieldSignature(signature)
	_dafg.Rect = _cac.MakeArrayFromFloats(opts.Rect)
	_dafg.AP = _cbgaa
	return _dafg, nil
}
func _fde(_cebd *_ge.PdfField) string {
	if _cebd == nil {
		return ""
	}
	_cad, _geb := _cebd.GetContext().(*_ge.PdfFieldText)
	if !_geb {
		return _fde(_cebd.Parent)
	}
	if _cad.DA != nil {
		return _cad.DA.Str()
	}
	return _fde(_cad.Parent)
}
func _dbe(_gbe [][]_ce.CubicBezierCurve, _afab *_ge.PdfColorDeviceRGB, _feaa float64) ([]byte, *_ge.PdfRectangle, error) {
	_eabf := _fb.NewContentCreator()
	_eabf.Add_q().SetStrokingColor(_afab).Add_w(_feaa)
	_gdf := _ce.NewCubicBezierPath()
	for _, _gdbf := range _gbe {
		_gdf.Curves = append(_gdf.Curves, _gdbf...)
		for _fdfg, _edgde := range _gdbf {
			if _fdfg == 0 {
				_eabf.Add_m(_edgde.P0.X, _edgde.P0.Y)
			} else {
				_eabf.Add_l(_edgde.P0.X, _edgde.P0.Y)
			}
			_eabf.Add_c(_edgde.P1.X, _edgde.P1.Y, _edgde.P2.X, _edgde.P2.Y, _edgde.P3.X, _edgde.P3.Y)
		}
	}
	_eabf.Add_S().Add_Q()
	return _eabf.Bytes(), _gdf.GetBoundingBox().ToPdfRectangle(), nil
}

// NewComboboxField generates a new combobox form field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewComboboxField(page *_ge.PdfPage, name string, rect []float64, opt ComboboxFieldOptions) (*_ge.PdfFieldChoice, error) {
	if page == nil {
		return nil, _aeg.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _aeg.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _aeg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_caca := _ge.NewPdfField()
	_cbd := &_ge.PdfFieldChoice{}
	_caca.SetContext(_cbd)
	_cbd.PdfField = _caca
	_cbd.T = _cac.MakeString(name)
	_cbd.Opt = _cac.MakeArray()
	for _, _fbdg := range opt.Choices {
		_cbd.Opt.Append(_cac.MakeString(_fbdg))
	}
	_cbd.SetFlag(_ge.FieldFlagCombo)
	_cfgfa := _ge.NewPdfAnnotationWidget()
	_cfgfa.Rect = _cac.MakeArrayFromFloats(rect)
	_cfgfa.P = page.ToPdfObject()
	_cfgfa.F = _cac.MakeInteger(4)
	_cfgfa.Parent = _cbd.ToPdfObject()
	_cbd.Annotations = append(_cbd.Annotations, _cfgfa)
	return _cbd, nil
}
func (_cd *AppearanceFont) fillName() {
	if _cd.Font == nil || _cd.Name != "" {
		return
	}
	_bba := _cd.Font.FontDescriptor()
	if _bba == nil || _bba.FontName == nil {
		return
	}
	_cd.Name = _bba.FontName.String()
}

// SetStyle applies appearance `style` to `fa`.
func (_dfc *FieldAppearance) SetStyle(style AppearanceStyle) { _dfc._de = &style }

// LineAnnotationDef defines a line between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none
// (regular line), or arrows at either end.  The line also has a specified width, color and opacity.
type LineAnnotationDef struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        *_ge.PdfColorDeviceRGB
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 _ce.LineEndingStyle
	LineEndingStyle2 _ce.LineEndingStyle
}

// SignatureLine represents a line of information in the signature field appearance.
type SignatureLine struct {
	Desc string
	Text string
}

func (_beee *AppearanceStyle) applyAppearanceCharacteristics(_befd *_cac.PdfObjectDictionary, _aee *_cac.PdfObjectDictionary, _bfg *_ge.PdfFont) error {
	if !_beee.AllowMK {
		return nil
	}
	if CA, _egc := _cac.GetString(_befd.Get("\u0043\u0041")); _egc && _bfg != nil {
		_adg := CA.Bytes()
		if len(_adg) != 0 {
			_cfb := []rune(_bfg.Encoder().Decode(_adg))
			if len(_cfb) == 1 {
				_beee.CheckmarkRune = _cfb[0]
			}
		}
	}
	if BC, _fbbg := _cac.GetArray(_befd.Get("\u0042\u0043")); _fbbg {
		_afa, _aebd := BC.ToFloat64Array()
		if _aebd != nil {
			return _aebd
		}
		switch len(_afa) {
		case 1:
			_beee.BorderColor = _ge.NewPdfColorDeviceGray(_afa[0])
		case 3:
			_beee.BorderColor = _ge.NewPdfColorDeviceRGB(_afa[0], _afa[1], _afa[2])
		case 4:
			_beee.BorderColor = _ge.NewPdfColorDeviceCMYK(_afa[0], _afa[1], _afa[2], _afa[3])
		default:
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0043\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_afa))
		}
		if _aee != nil {
			if _bbbe, _cedc := _cac.GetNumberAsFloat(_aee.Get("\u0057")); _cedc == nil {
				_beee.BorderSize = _bbbe
			}
		}
	}
	if BG, _aaae := _cac.GetArray(_befd.Get("\u0042\u0047")); _aaae {
		_ecfa, _cba := BG.ToFloat64Array()
		if _cba != nil {
			return _cba
		}
		switch len(_ecfa) {
		case 1:
			_beee.FillColor = _ge.NewPdfColorDeviceGray(_ecfa[0])
		case 3:
			_beee.FillColor = _ge.NewPdfColorDeviceRGB(_ecfa[0], _ecfa[1], _ecfa[2])
		case 4:
			_beee.FillColor = _ge.NewPdfColorDeviceCMYK(_ecfa[0], _ecfa[1], _ecfa[2], _ecfa[3])
		default:
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0047\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_ecfa))
		}
	}
	return nil
}
func _gbdb(_geca _cac.PdfObject, _cff *_ge.PdfPageResources) (*_cac.PdfObjectName, float64, bool) {
	var (
		_fbc  *_cac.PdfObjectName
		_dcae float64
		_agce bool
	)
	if _degg, _eacb := _cac.GetDict(_geca); _eacb && _degg != nil {
		_eeda := _cac.TraceToDirectObject(_degg.Get("\u004e"))
		switch _fafb := _eeda.(type) {
		case *_cac.PdfObjectStream:
			_ecc, _gccd := _cac.DecodeStream(_fafb)
			if _gccd != nil {
				_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u0020\u0073\u0074r\u0065\u0061m\u003a\u0020\u0025\u0076", _gccd.Error())
				return nil, 0, false
			}
			_fddc, _gccd := _fb.NewContentStreamParser(string(_ecc)).Parse()
			if _gccd != nil {
				_g.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0075n\u0061\u0062l\u0065\u0020\u0070\u0061\u0072\u0073\u0065\u0020c\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061m\u003a\u0020\u0025\u0076", _gccd.Error())
				return nil, 0, false
			}
			_afgf := _fb.NewContentStreamProcessor(*_fddc)
			_afgf.AddHandler(_fb.HandlerConditionEnumOperand, "\u0054\u0066", func(_cef *_fb.ContentStreamOperation, _dgfc _fb.GraphicsState, _agef *_ge.PdfPageResources) error {
				if len(_cef.Params) == 2 {
					if _afega, _gdab := _cac.GetName(_cef.Params[0]); _gdab {
						_fbc = _afega
					}
					if _aeag, _cdf := _cac.GetNumberAsFloat(_cef.Params[1]); _cdf == nil {
						_dcae = _aeag
					}
					_agce = true
					return _fb.ErrEarlyExit
				}
				return nil
			})
			_afgf.Process(_cff)
			return _fbc, _dcae, _agce
		}
	}
	return nil, 0, false
}
func _fbb(_gca CircleAnnotationDef, _ad string) ([]byte, *_ge.PdfRectangle, *_ge.PdfRectangle, error) {
	_ec := _ce.Circle{X: _gca.X, Y: _gca.Y, Width: _gca.Width, Height: _gca.Height, FillEnabled: _gca.FillEnabled, FillColor: _gca.FillColor, BorderEnabled: _gca.BorderEnabled, BorderWidth: _gca.BorderWidth, BorderColor: _gca.BorderColor, Opacity: _gca.Opacity}
	_caf, _fc, _dfa := _ec.Draw(_ad)
	if _dfa != nil {
		return nil, nil, nil, _dfa
	}
	_bc := &_ge.PdfRectangle{}
	_bc.Llx = _gca.X + _fc.Llx
	_bc.Lly = _gca.Y + _fc.Lly
	_bc.Urx = _gca.X + _fc.Urx
	_bc.Ury = _gca.Y + _fc.Ury
	return _caf, _fc, _bc, nil
}
func _cafc(_fgg *_ge.PdfAnnotationWidget, _gff *_ge.PdfFieldText, _bef *_ge.PdfPageResources, _deg AppearanceStyle) (*_cac.PdfObjectDictionary, error) {
	_dde := _ge.NewPdfPageResources()
	_bfa, _ggd := _cac.GetArray(_fgg.Rect)
	if !_ggd {
		return nil, _aeg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_adde, _adc := _ge.NewPdfRectangle(*_bfa)
	if _adc != nil {
		return nil, _adc
	}
	_gcab, _fbge := _adde.Width(), _adde.Height()
	_acfd, _eaa := _gcab, _fbge
	_efe, _accc := _cac.GetDict(_fgg.MK)
	if _accc {
		_ceb, _ := _cac.GetDict(_fgg.BS)
		_eag := _deg.applyAppearanceCharacteristics(_efe, _ceb, nil)
		if _eag != nil {
			return nil, _eag
		}
	}
	_gda, _accc := _cac.GetIntVal(_gff.MaxLen)
	if !_accc {
		return nil, _aeg.New("\u006d\u0061\u0078\u006c\u0065\u006e\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if _gda <= 0 {
		return nil, _aeg.New("\u006d\u0061\u0078\u004c\u0065\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_ddef := _gcab / float64(_gda)
	_dda, _adc := _fb.NewContentStreamParser(_fde(_gff.PdfField)).Parse()
	if _adc != nil {
		return nil, _adc
	}
	_gecg := _fb.NewContentCreator()
	if _deg.BorderSize > 0 {
		_abb(_gecg, _deg, _gcab, _fbge)
	}
	if _deg.DrawAlignmentReticle {
		_gfa := _deg
		_gfa.BorderSize = 0.2
		_bgg(_gecg, _gfa, _gcab, _fbge)
	}
	_gecg.Add_BMC("\u0054\u0078")
	_gecg.Add_q()
	_, _fbge = _deg.applyRotation(_efe, _gcab, _fbge, _gecg)
	_gecg.Add_BT()
	_aea, _egb, _adc := _deg.processDA(_gff.PdfField, _dda, _bef, _dde, _gecg)
	if _adc != nil {
		return nil, _adc
	}
	_dfg := _aea.Font
	_cacc := _cac.MakeName(_aea.Name)
	_afb := _aea.Size
	_agc := _afb == 0
	if _agc && _egb {
		_afb = _fbge * _deg.AutoFontSizeFraction
	}
	_efeb := _dfg.Encoder()
	if _efeb == nil {
		_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_efeb = _ac.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	var _bca string
	if _ada, _bad := _cac.GetString(_gff.V); _bad {
		_bca = _ada.Decoded()
	}
	_gecg.Add_Tf(*_cacc, _afb)
	var _fcb float64
	for _, _cdb := range _bca {
		_eege, _dfdg := _dfg.GetRuneMetrics(_cdb)
		if !_dfdg {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _cdb)
			continue
		}
		_ebb := _eege.Wy
		if int(_ebb) <= 0 {
			_ebb = _eege.Wx
		}
		if _ebb > _fcb {
			_fcb = _ebb
		}
	}
	if int(_fcb) == 0 {
		_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006d\u0061x\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0073\u0069\u007a\u0065\u0020- \u0075\u0073\u0069\u006e\u0067\u0020\u0031\u0030\u0030\u0030")
		_fcb = 1000
	}
	_cfc, _adc := _dfg.GetFontDescriptor()
	if _adc != nil {
		_g.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _gcff float64
	if _cfc != nil {
		_gcff, _adc = _cfc.GetCapHeight()
		if _adc != nil {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _adc)
		}
	}
	if int(_gcff) <= 0 {
		_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_gcff = 1000.0
	}
	_adae := _gcff / 1000.0 * _afb
	_ddcc := 0.0
	_cae := 1.0 * _afb * (_fcb / 1000.0)
	{
		_agbg := _cae
		if _agc && _ddcc+_agbg > _fbge {
			_afb = 0.95 * (_fbge - _ddcc)
			_adae = _gcff / 1000.0 * _afb
		}
		if _fbge > _adae {
			_ddcc = (_fbge - _adae) / 2.0
		}
	}
	_gecg.Add_Td(0, _ddcc)
	if _bdc, _dadb := _cac.GetIntVal(_gff.Q); _dadb {
		switch _bdc {
		case 2:
			if len(_bca) < _gda {
				_aeda := float64(_gda-len(_bca)) * _ddef
				_gecg.Add_Td(_aeda, 0)
			}
		}
	}
	for _dbg, _gfc := range _bca {
		_aaca := _gec
		if _deg.MarginLeft != nil {
			_aaca = *_deg.MarginLeft
		}
		_ggb := string(_gfc)
		if _efeb != nil {
			_bcg, _agd := _dfg.GetRuneMetrics(_gfc)
			if !_agd {
				_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _gfc)
				continue
			}
			_ggb = string(_efeb.Encode(_ggb))
			_efd := _afb * _bcg.Wx / 1000.0
			_edbb := (_ddef - _efd) / 2
			_aaca = _edbb
		}
		_gecg.Add_Td(_aaca, 0)
		_gecg.Add_Tj(*_cac.MakeString(_ggb))
		if _dbg != len(_bca)-1 {
			_gecg.Add_Td(_ddef-_aaca, 0)
		}
	}
	_gecg.Add_ET()
	_gecg.Add_Q()
	_gecg.Add_EMC()
	_ecf := _ge.NewXObjectForm()
	_ecf.Resources = _dde
	_ecf.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _acfd, _eaa})
	_ecf.SetContentStream(_gecg.Bytes(), _daf())
	_gfbb := _cac.MakeDict()
	_gfbb.Set("\u004e", _ecf.ToPdfObject())
	return _gfbb, nil
}

// ComboboxFieldOptions defines optional parameters for a combobox form field.
type ComboboxFieldOptions struct {

	// Choices is the list of string values that can be selected.
	Choices []string
}

// SignatureImagePosition specifies the image signature location relative to the text signature.
// If text signature is not defined, this position will be ignored.
type SignatureImagePosition int

func _eea(_gegg []float64) []float64 {
	var (
		_beb  = len(_gegg)
		_acee = make([]float64, _beb)
		_eec  = make([]float64, _beb)
	)
	_fbda := 2.0
	_acee[0] = _gegg[0] / _fbda
	for _abdd := 1; _abdd < _beb; _abdd++ {
		_eec[_abdd] = 1 / _fbda
		if _abdd < _beb-1 {
			_fbda = 4.0
		} else {
			_fbda = 3.5
		}
		_fbda -= _eec[_abdd]
		_acee[_abdd] = (_gegg[_abdd] - _acee[_abdd-1]) / _fbda
	}
	for _aeaf := 1; _aeaf < _beb; _aeaf++ {
		_acee[_beb-_aeaf-1] -= _eec[_beb-_aeaf] * _acee[_beb-_aeaf]
	}
	return _acee
}

const (
	SignatureImageLeft SignatureImagePosition = iota
	SignatureImageRight
	SignatureImageTop
	SignatureImageBottom
)

func _cdff(_face *_ge.PdfFieldButton, _dac *_ge.PdfAnnotationWidget, _ggg AppearanceStyle) (*_cac.PdfObjectDictionary, error) {
	_gbgb, _bafa := _cac.GetArray(_dac.Rect)
	if !_bafa {
		return nil, _aeg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_fbfe, _fed := _ge.NewPdfRectangle(*_gbgb)
	if _fed != nil {
		return nil, _fed
	}
	_ccg, _aced := _fbfe.Width(), _fbfe.Height()
	_bgcd := _fb.NewContentCreator()
	if _ggg.BorderSize > 0 {
		_abb(_bgcd, _ggg, _ccg, _aced)
	}
	if _ggg.DrawAlignmentReticle {
		_ggde := _ggg
		_ggde.BorderSize = 0.2
		_bgg(_bgcd, _ggde, _ccg, _aced)
	}
	_dega := _face.GetFillImage()
	_bcde, _fed := _dbgb(_ccg, _aced, _dega, _ggg)
	if _fed != nil {
		return nil, _fed
	}
	_bfdf, _gfafe := _cac.GetDict(_dac.MK)
	if _gfafe {
		_bfdf.Set("\u006c", _bcde.ToPdfObject())
	}
	_fdfd := _cac.MakeDict()
	_fdfd.Set("\u0046\u0052\u004d", _bcde.ToPdfObject())
	_fcfd := _ge.NewPdfPageResources()
	_fcfd.ProcSet = _cac.MakeArray(_cac.MakeName("\u0050\u0044\u0046"))
	_fcfd.XObject = _fdfd
	_bcged := _ccg - 2
	_eagb := _aced - 2
	_bgcd.Add_q()
	_bgcd.Add_re(1, 1, _bcged, _eagb)
	_bgcd.Add_W()
	_bgcd.Add_n()
	_bcged -= 2
	_eagb -= 2
	_bgcd.Add_q()
	_bgcd.Add_re(2, 2, _bcged, _eagb)
	_bgcd.Add_W()
	_bgcd.Add_n()
	_bgde := _fg.Min(_bcged/float64(_dega.Width), _eagb/float64(_dega.Height))
	_bgcd.Add_cm(_bgde, 0, 0, _bgde, (_ccg/2)-(float64(_dega.Width)*_bgde/2)+2, 2)
	_bgcd.Add_Do("\u0046\u0052\u004d")
	_bgcd.Add_Q()
	_bgcd.Add_Q()
	_fcgb := _ge.NewXObjectForm()
	_fcgb.FormType = _cac.MakeInteger(1)
	_fcgb.Resources = _fcfd
	_fcgb.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _ccg, _aced})
	_fcgb.Matrix = _cac.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_fcgb.SetContentStream(_bgcd.Bytes(), _daf())
	_bfbe := _cac.MakeDict()
	_bfbe.Set("\u004e", _fcgb.ToPdfObject())
	return _bfbe, nil
}

// CreateFileAttachmentAnnotation creates a file attachment annotation object that can be added to the annotation list of a PDF page.
func CreateFileAttachmentAnnotation(fileDef FileAnnotationDef) (*_ge.PdfAnnotation, error) {
	_caab := _ge.NewPdfFileSpecFromEmbeddedFile(fileDef.EmbeddedFile)
	if fileDef.Color == nil {
		fileDef.Color = _ge.NewPdfColorDeviceRGB(0.0, 0.0, 0.0)
	}
	if fileDef.Description == "" {
		fileDef.Description = fileDef.EmbeddedFile.Name
	}
	if fileDef.CreationDate == nil {
		_dfae := _a.Now()
		fileDef.CreationDate = &_dfae
	}
	if fileDef.IconName == "" {
		fileDef.IconName = "\u0050u\u0073\u0068\u0050\u0069\u006e"
	}
	_gdag, _bfgf := _ge.NewPdfDateFromTime(*fileDef.CreationDate)
	if _bfgf != nil {
		return nil, _bfgf
	}
	_agcd := _ge.NewPdfAnnotationFileAttachment()
	_agcd.FS = _caab.ToPdfObject()
	_agcd.C = _cac.MakeArrayFromFloats([]float64{fileDef.Color.R(), fileDef.Color.G(), fileDef.Color.B()})
	_agcd.Contents = _cac.MakeString(fileDef.Description)
	_agcd.CreationDate = _gdag.ToPdfObject()
	_agcd.M = _gdag.ToPdfObject()
	_agcd.Name = _cac.MakeName(fileDef.IconName)
	_agcd.Rect = _cac.MakeArrayFromFloats([]float64{fileDef.X, fileDef.Y, fileDef.X + fileDef.Width, fileDef.Y + fileDef.Height})
	_agcd.T = _cac.MakeString(fileDef.Author)
	_agcd.Subj = _cac.MakeString(fileDef.Subject)
	return _agcd.PdfAnnotation, nil
}

// NewFormResetButtonField would create a reset button in specified page according to the parameter in `FormResetActionOptions`.
func NewFormResetButtonField(page *_ge.PdfPage, opt FormResetActionOptions) (*_ge.PdfFieldButton, error) {
	_edbbe := _ge.NewPdfActionResetForm()
	_edbbe.Fields = opt.Fields
	_edbbe.Flags = _cac.MakeInteger(0)
	if opt.IsExclusionList {
		_edbbe.Flags = _cac.MakeInteger(1)
	}
	_def, _cefd := _eaae(page, opt.Rectangle, "\u0062\u0074\u006e\u0052\u0065\u0073\u0065\u0074", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _edbbe.ToPdfObject())
	if _cefd != nil {
		return nil, _cefd
	}
	return _def, nil
}
func _bgg(_fbgb *_fb.ContentCreator, _gega AppearanceStyle, _efaf, _beeb float64) {
	_fbgb.Add_q().Add_re(0, 0, _efaf, _beeb).Add_re(0, _beeb/2, _efaf, _beeb/2).Add_re(0, 0, _efaf, _beeb).Add_re(_efaf/2, 0, _efaf/2, _beeb).Add_w(_gega.BorderSize).SetStrokingColor(_gega.BorderColor).SetNonStrokingColor(_gega.FillColor).Add_B().Add_Q()
}

// AppearanceFont represents a font used for generating the appearance of a
// field in the filling/flattening process.
type AppearanceFont struct {

	// Name represents the name of the font which will be added to the
	// AcroForm resources (DR).
	Name string

	// Font represents the actual font used for the field appearance.
	Font *_ge.PdfFont

	// Size represents the size of the font used for the field appearance.
	// If the font size is 0, the value of the FallbackSize field of the
	// AppearanceFontStyle is used, if set. Otherwise, the font size is
	// calculated based on the available annotation height and on the
	// AutoFontSizeFraction field of the AppearanceStyle.
	Size float64
}

// TextFieldOptions defines optional parameter for a text field in a form.
type TextFieldOptions struct {
	MaxLen int
	Value  string

	// TextColor defines the color of the text in hex format. e.g #43fd23.
	// If it has an invalid value a #000000 (black) color is taken as default
	TextColor string

	// FontName defines the font of the text. Helvetica font is the default one.
	// It is recommended to use one of 14 standard PDF fonts.
	FontName string

	// FontSize defines the font size of the text, 12 is used by default.
	FontSize int
}

// AppearanceFontStyle defines font style characteristics for form fields,
// used in the filling/flattening process.
type AppearanceFontStyle struct {

	// Fallback represents a global font fallback, used for fields which do
	// not specify a font in their default appearance (DA). The fallback is
	// also used if there is a font specified in the DA, but it is not
	// found in the AcroForm resources (DR).
	Fallback *AppearanceFont

	// FallbackSize represents a global font size fallback used for fields
	// which do not specify a font size in their default appearance (DA).
	// The fallback size is applied only if its value is larger than zero.
	FallbackSize float64

	// FieldFallbacks defines font fallbacks for specific fields. The map keys
	// represent the names of the fields (which can be specified by their
	// partial or full names). Specific field fallback fonts take precedence
	// over the global font fallback.
	FieldFallbacks map[string]*AppearanceFont

	// ForceReplace forces the replacement of fonts in the filling/flattening
	// process, even if the default appearance (DA) specifies a valid font.
	// If no fallback font is provided, setting this field has no effect.
	ForceReplace bool
}

// CreateRectangleAnnotation creates a rectangle annotation object that can be added to page PDF annotations.
func CreateRectangleAnnotation(rectDef RectangleAnnotationDef) (*_ge.PdfAnnotation, error) {
	_faag := _ge.NewPdfAnnotationSquare()
	if rectDef.BorderEnabled {
		_fgee, _adge, _bccd := rectDef.BorderColor.R(), rectDef.BorderColor.G(), rectDef.BorderColor.B()
		_faag.C = _cac.MakeArrayFromFloats([]float64{_fgee, _adge, _bccd})
		_efafe := _ge.NewBorderStyle()
		_efafe.SetBorderWidth(rectDef.BorderWidth)
		_faag.BS = _efafe.ToPdfObject()
	}
	if rectDef.FillEnabled {
		_dbfg, _effd, _bacf := rectDef.FillColor.R(), rectDef.FillColor.G(), rectDef.FillColor.B()
		_faag.IC = _cac.MakeArrayFromFloats([]float64{_dbfg, _effd, _bacf})
	} else {
		_faag.IC = _cac.MakeArrayFromIntegers([]int{})
	}
	if rectDef.Opacity < 1.0 {
		_faag.CA = _cac.MakeFloat(rectDef.Opacity)
	}
	_fcgc, _ecea, _feaf := _aaed(rectDef)
	if _feaf != nil {
		return nil, _feaf
	}
	_faag.AP = _fcgc
	_faag.Rect = _cac.MakeArrayFromFloats([]float64{_ecea.Llx, _ecea.Lly, _ecea.Urx, _ecea.Ury})
	return _faag.PdfAnnotation, nil
}

// FileAnnotationDef holds base information for constructing an file attachment annotation.
type FileAnnotationDef struct {

	// Bounding box of the annotation.
	X      float64
	Y      float64
	Width  float64
	Height float64

	// EmbeddedFile is the file information to be attached.
	EmbeddedFile *_ge.EmbeddedFile

	// Author is the author of the attachment file.
	Author string

	// Subject is the subject of the attachment file.
	Subject string

	// Description of the file attachment that will be displayed as a comment on the PDF reader.
	Description string

	// IconName is The name of an icon that shall be used in displaying the annotation.
	// Conforming readers shall provide predefined icon appearances for at least the following standard names:
	//
	// - Graph
	// - PushPin
	// - Paperclip
	// - Tag
	//
	// Additional names may be supported as well. Default value: "PushPin".
	IconName string

	// Color is the color of the annotation.
	Color *_ge.PdfColorDeviceRGB

	// CreationDate is the date and time when the file attachment was created.
	// If not set, the current time is used.
	CreationDate *_a.Time
}

// ImageFieldOptions defines optional parameters for a push button with image attach capability form field.
type ImageFieldOptions struct {
	Image  *_ge.Image
	_fgcbg AppearanceStyle
}

func _cbdd(_dcdf RectangleAnnotationDef, _dedf string) ([]byte, *_ge.PdfRectangle, *_ge.PdfRectangle, error) {
	_ffga := _ce.Rectangle{X: 0, Y: 0, Width: _dcdf.Width, Height: _dcdf.Height, FillEnabled: _dcdf.FillEnabled, FillColor: _dcdf.FillColor, BorderEnabled: _dcdf.BorderEnabled, BorderWidth: 2 * _dcdf.BorderWidth, BorderColor: _dcdf.BorderColor, Opacity: _dcdf.Opacity}
	_gefe, _bcgg, _acgd := _ffga.Draw(_dedf)
	if _acgd != nil {
		return nil, nil, nil, _acgd
	}
	_ebaa := &_ge.PdfRectangle{}
	_ebaa.Llx = _dcdf.X + _bcgg.Llx
	_ebaa.Lly = _dcdf.Y + _bcgg.Lly
	_ebaa.Urx = _dcdf.X + _bcgg.Urx
	_ebaa.Ury = _dcdf.Y + _bcgg.Ury
	return _gefe, _bcgg, _ebaa, nil
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_ea FieldAppearance) GenerateAppearanceDict(form *_ge.PdfAcroForm, field *_ge.PdfField, wa *_ge.PdfAnnotationWidget) (*_cac.PdfObjectDictionary, error) {
	_g.Log.Trace("\u0047\u0065n\u0065\u0072\u0061\u0074e\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0044i\u0063\u0074\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u0020\u0056:\u0020\u0025\u002b\u0076", field.PartialName(), field.V)
	_, _fgc := field.GetContext().(*_ge.PdfFieldText)
	_dg, _dgg := _cac.GetDict(wa.AP)
	if _dgg && _ea.OnlyIfMissing && (!_fgc || !_ea.RegenerateTextFields) {
		_g.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _dg, nil
	}
	if form.DR == nil {
		form.DR = _ge.NewPdfPageResources()
	}
	switch _bg := field.GetContext().(type) {
	case *_ge.PdfFieldText:
		_ddc := _bg
		if _dgf := _fde(_ddc.PdfField); _dgf == "" {
			_ddc.DA = form.DA
		}
		if _ea._de != nil && _ea._de.TextColor != nil {
			_dfac := _fb.ContentStreamOperations{}
			_da := _fde(_ddc.PdfField)
			_ffa, _dfd := _fb.NewContentStreamParser(_da).Parse()
			if _dfd != nil {
				return nil, _dfd
			}
			for _, _adbb := range *_ffa {
				if _adbb.Operand == "\u0067" || _adbb.Operand == "\u0072\u0067" {
					continue
				}
				_dfac = append(_dfac, _adbb)
			}
			_ed := _ea._de.TextColor
			_fgce, _eb := _ed.(*_ge.PdfColorDeviceRGB)
			if !_eb {
				return nil, _dfd
			}
			_bgd, _cbb, _aga := _cac.MakeFloat(_fgce[0]), _cac.MakeFloat(_fgce[1]), _cac.MakeFloat(_fgce[2])
			_caa := &_fb.ContentStreamOperation{Params: []_cac.PdfObject{_bgd, _cbb, _aga}, Operand: "\u0072\u0067"}
			_dfac = append(_dfac, _caa)
			_fgca := _dfac.String()
			_fgca = _ae.Replace(_fgca, "\u000a", "\u0020", -1)
			_fgca = _ae.Trim(_fgca, "\u0020")
			_ddc.DA = _cac.MakeHexString(_fgca)
		}
		switch {
		case _ddc.Flags().Has(_ge.FieldFlagPassword):
			return nil, nil
		case _ddc.Flags().Has(_ge.FieldFlagFileSelect):
			return nil, nil
		case _ddc.Flags().Has(_ge.FieldFlagComb):
			if _ddc.MaxLen != nil {
				_afc, _ee := _cafc(wa, _ddc, form.DR, _ea.Style())
				if _ee != nil {
					return nil, _ee
				}
				return _afc, nil
			}
		}
		_aeb, _gd := _ggc(wa, _ddc, form.DR, _ea.Style())
		if _gd != nil {
			return nil, _gd
		}
		return _aeb, nil
	case *_ge.PdfFieldButton:
		_ffe := _bg
		if _ffe.IsCheckbox() {
			_dc, _edg := _bfad(wa, _ffe, form.DR, _ea.Style())
			if _edg != nil {
				return nil, _edg
			}
			return _dc, nil
		}
		_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041\u004e\u0044\u004c\u0045\u0044 \u0062u\u0074\u0074\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u002b\u0076", _ffe.GetType())
	case *_ge.PdfFieldChoice:
		_afca := _bg
		switch {
		case _afca.Flags().Has(_ge.FieldFlagCombo):
			_cacd, _eeg := _edga(form, wa, _afca, _ea.Style())
			if _eeg != nil {
				return nil, _eeg
			}
			return _cacd, nil
		default:
			_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055N\u0048\u0041\u004eD\u004c\u0045\u0044\u0020c\u0068\u006f\u0069\u0063\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0025\u0073", _afca.Flags().String())
		}
	default:
		_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041N\u0044\u004c\u0045\u0044\u0020\u0066\u0069e\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _bg)
	}
	return nil, nil
}

// NewImageField generates a new image field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewImageField(page *_ge.PdfPage, name string, rect []float64, opt ImageFieldOptions) (*_ge.PdfFieldButton, error) {
	if page == nil {
		return nil, _aeg.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _aeg.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _aeg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_bage := _ge.NewPdfField()
	_gcb := &_ge.PdfFieldButton{}
	_gcb.PdfField = _bage
	_bage.SetContext(_gcb)
	_gcb.SetType(_ge.ButtonTypePush)
	_gcb.T = _cac.MakeString(name)
	_ccef := _ge.NewPdfAnnotationWidget()
	_ccef.Rect = _cac.MakeArrayFromFloats(rect)
	_ccef.P = page.ToPdfObject()
	_ccef.F = _cac.MakeInteger(4)
	_ccef.Parent = _gcb.ToPdfObject()
	_dbbg := rect[2] - rect[0]
	_gdce := rect[3] - rect[1]
	_agfg := opt._fgcbg
	_fggc := _fb.NewContentCreator()
	if _agfg.BorderSize > 0 {
		_abb(_fggc, _agfg, _dbbg, _gdce)
	}
	if _agfg.DrawAlignmentReticle {
		_egdb := _agfg
		_egdb.BorderSize = 0.2
		_bgg(_fggc, _egdb, _dbbg, _gdce)
	}
	_cgf, _becb := _dbgb(_dbbg, _gdce, opt.Image, _agfg)
	if _becb != nil {
		return nil, _becb
	}
	_ceba, _cadf := _cac.GetDict(_ccef.MK)
	if _cadf {
		_ceba.Set("\u006c", _cgf.ToPdfObject())
	}
	_cdfe := _cac.MakeDict()
	_cdfe.Set("\u0046\u0052\u004d", _cgf.ToPdfObject())
	_fdab := _ge.NewPdfPageResources()
	_fdab.ProcSet = _cac.MakeArray(_cac.MakeName("\u0050\u0044\u0046"))
	_fdab.XObject = _cdfe
	_dgbbb := _dbbg - 2
	_aefd := _gdce - 2
	_fggc.Add_q()
	_fggc.Add_re(1, 1, _dgbbb, _aefd)
	_fggc.Add_W()
	_fggc.Add_n()
	_dgbbb -= 2
	_aefd -= 2
	_fggc.Add_q()
	_fggc.Add_re(2, 2, _dgbbb, _aefd)
	_fggc.Add_W()
	_fggc.Add_n()
	_egfa := _fg.Min(_dgbbb/float64(opt.Image.Width), _aefd/float64(opt.Image.Height))
	_fggc.Add_cm(_egfa, 0, 0, _egfa, (_dbbg/2)-(float64(opt.Image.Width)*_egfa/2)+2, 2)
	_fggc.Add_Do("\u0046\u0052\u004d")
	_fggc.Add_Q()
	_fggc.Add_Q()
	_edeg := _ge.NewXObjectForm()
	_edeg.FormType = _cac.MakeInteger(1)
	_edeg.Resources = _fdab
	_edeg.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, _dbbg, _gdce})
	_edeg.Matrix = _cac.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_edeg.SetContentStream(_fggc.Bytes(), _daf())
	_cgfe := _cac.MakeDict()
	_cgfe.Set("\u004e", _edeg.ToPdfObject())
	_ccef.AP = _cgfe
	_gcb.Annotations = append(_gcb.Annotations, _ccef)
	return _gcb, nil
}

type quadding int

func _dbgb(_badg, _gcge float64, _bbdf *_ge.Image, _ggef AppearanceStyle) (*_ge.XObjectForm, error) {
	_efbe, _eegg := _ge.NewXObjectImageFromImage(_bbdf, nil, _cac.NewFlateEncoder())
	if _eegg != nil {
		return nil, _eegg
	}
	_efbe.Decode = _cac.MakeArrayFromFloats([]float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0})
	_dcfc := _ge.NewPdfPageResources()
	_dcfc.ProcSet = _cac.MakeArray(_cac.MakeName("\u0050\u0044\u0046"), _cac.MakeName("\u0049\u006d\u0061\u0067\u0065\u0043"))
	_dcfc.SetXObjectImageByName(_cac.PdfObjectName("\u0049\u006d\u0030"), _efbe)
	_dba := _fb.NewContentCreator()
	_dba.Add_q()
	_dba.Add_cm(float64(_bbdf.Width), 0, 0, float64(_bbdf.Height), 0, 0)
	_dba.Add_Do("\u0049\u006d\u0030")
	_dba.Add_Q()
	_gfba := _ge.NewXObjectForm()
	_gfba.FormType = _cac.MakeInteger(1)
	_gfba.BBox = _cac.MakeArrayFromFloats([]float64{0, 0, float64(_bbdf.Width), float64(_bbdf.Height)})
	_gfba.Resources = _dcfc
	_gfba.SetContentStream(_dba.Bytes(), _daf())
	return _gfba, nil
}

const (
	_gdge  = 1
	_aefdg = 2
	_ageb  = 4
	_degd  = 8
	_bdeec = 16
	_afcg  = 32
	_cfa   = 64
	_dgac  = 128
	_gagg  = 256
	_gdgb  = 512
	_cdaf  = 1024
	_abfg  = 2048
	_ebe   = 4096
)

// CreateInkAnnotation creates an ink annotation object that can be added to the annotation list of a PDF page.
func CreateInkAnnotation(inkDef InkAnnotationDef) (*_ge.PdfAnnotation, error) {
	_cfae := _ge.NewPdfAnnotationInk()
	_edca := _cac.MakeArray()
	for _, _ffgf := range inkDef.Paths {
		if _ffgf.Length() == 0 {
			continue
		}
		_gbf := []float64{}
		for _, _gaeg := range _ffgf.Points {
			_gbf = append(_gbf, _gaeg.X, _gaeg.Y)
		}
		_edca.Append(_cac.MakeArrayFromFloats(_gbf))
	}
	_cfae.InkList = _edca
	if inkDef.Color == nil {
		inkDef.Color = _ge.NewPdfColorDeviceRGB(0.0, 0.0, 0.0)
	}
	_cfae.C = _cac.MakeArrayFromFloats([]float64{inkDef.Color.R(), inkDef.Color.G(), inkDef.Color.B()})
	_baa, _cbfc, _fcbg := _ebec(&inkDef)
	if _fcbg != nil {
		return nil, _fcbg
	}
	_cfae.AP = _baa
	_cfae.Rect = _cac.MakeArrayFromFloats([]float64{_cbfc.Llx, _cbfc.Lly, _cbfc.Urx, _cbfc.Ury})
	return _cfae.PdfAnnotation, nil
}

// NewSignatureFieldOpts returns a new initialized instance of options
// used to generate a signature appearance.
func NewSignatureFieldOpts() *SignatureFieldOpts {
	return &SignatureFieldOpts{Font: _ge.DefaultFont(), FontSize: 10, LineHeight: 1, AutoSize: true, TextColor: _ge.NewPdfColorDeviceGray(0), BorderColor: _ge.NewPdfColorDeviceGray(0), FillColor: _ge.NewPdfColorDeviceGray(1), Encoder: _cac.NewFlateEncoder(), ImagePosition: SignatureImageLeft}
}
