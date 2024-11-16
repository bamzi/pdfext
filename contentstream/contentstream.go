// Package contentstream provides functionality for parsing and creating content streams for PDF files.
//
// For processing and manipulating content streams, it allows parse the content stream into a list of
// operands that can then be processed further for rendering or extraction of information.
// The ContentStreamProcessor offers a basic engine for processing the content stream and can be used
// to render or modify the contents.
//
// For creating content streams, see NewContentCreator.  It allows adding multiple operands and then can
// be converted to a string for embedding in a PDF file.
//
// The contentstream package uses the core and model packages.
package contentstream

import (
	_f "bufio"
	_a "bytes"
	_g "encoding/hex"
	_b "errors"
	_e "fmt"
	_bg "image/color"
	_fd "image/jpeg"
	_ce "io"
	_de "math"
	_d "regexp"
	_bb "strconv"

	_ec "github.com/bamzi/pdfext/common"
	_cg "github.com/bamzi/pdfext/core"
	_aa "github.com/bamzi/pdfext/internal/imageutil"
	_da "github.com/bamzi/pdfext/internal/transform"
	_dad "github.com/bamzi/pdfext/model"
)

// Add_Tm appends 'Tm' operand to the content stream:
// Set the text line matrix.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_efe *ContentCreator) Add_Tm(a, b, c, d, e, f float64) *ContentCreator {
	_dcc := ContentStreamOperation{}
	_dcc.Operand = "\u0054\u006d"
	_dcc.Params = _gbg([]float64{a, b, c, d, e, f})
	_efe._bf = append(_efe._bf, &_dcc)
	return _efe
}

// Add_J adds 'J' operand to the content stream: Set the line cap style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ggb *ContentCreator) Add_J(lineCapStyle string) *ContentCreator {
	_fg := ContentStreamOperation{}
	_fg.Operand = "\u004a"
	_fg.Params = _fcdg([]_cg.PdfObjectName{_cg.PdfObjectName(lineCapStyle)})
	_ggb._bf = append(_ggb._bf, &_fg)
	return _ggb
}

// IsMask checks if an image is a mask.
// The image mask entry in the image dictionary specifies that the image data shall be used as a stencil
// mask for painting in the current color. The mask data is 1bpc, grayscale.
func (_cae *ContentStreamInlineImage) IsMask() (bool, error) {
	if _cae.ImageMask != nil {
		_gdff, _cagf := _cae.ImageMask.(*_cg.PdfObjectBool)
		if !_cagf {
			_ec.Log.Debug("\u0049m\u0061\u0067\u0065\u0020\u006d\u0061\u0073\u006b\u0020\u006e\u006ft\u0020\u0061\u0020\u0062\u006f\u006f\u006c\u0065\u0061\u006e")
			return false, _b.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		return bool(*_gdff), nil
	}
	return false, nil
}

// Add_Tc appends 'Tc' operand to the content stream:
// Set character spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ebe *ContentCreator) Add_Tc(charSpace float64) *ContentCreator {
	_eac := ContentStreamOperation{}
	_eac.Operand = "\u0054\u0063"
	_eac.Params = _gbg([]float64{charSpace})
	_ebe._bf = append(_ebe._bf, &_eac)
	return _ebe
}
func (_bdbd *ContentStreamProcessor) handleCommand_g(_adde *ContentStreamOperation, _cfda *_dad.PdfPageResources) error {
	_gdcga := _dad.NewPdfColorspaceDeviceGray()
	if len(_adde.Params) != _gdcga.GetNumComponents() {
		_ec.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020p\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0067")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_adde.Params), _gdcga)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fded, _cedd := _gdcga.ColorFromPdfObjects(_adde.Params)
	if _cedd != nil {
		_ec.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0043o\u006d\u006d\u0061\u006e\u0064\u005f\u0067\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061r\u0061\u006d\u0073\u002e\u0020c\u0073\u003d\u0025\u0054\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gdcga, _adde, _cedd)
		return _cedd
	}
	_bdbd._ebdd.ColorspaceNonStroking = _gdcga
	_bdbd._ebdd.ColorNonStroking = _fded
	return nil
}

// Add_BT appends 'BT' operand to the content stream:
// Begin text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_ccb *ContentCreator) Add_BT() *ContentCreator {
	_fgfg := ContentStreamOperation{}
	_fgfg.Operand = "\u0042\u0054"
	_ccb._bf = append(_ccb._bf, &_fgfg)
	return _ccb
}
func (_ffac *ContentStreamProcessor) handleCommand_cs(_gcfd *ContentStreamOperation, _acd *_dad.PdfPageResources) error {
	if len(_gcfd.Params) < 1 {
		_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_gcfd.Params) > 1 {
		_ec.Log.Debug("\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_gafag, _afeb := _gcfd.Params[0].(*_cg.PdfObjectName)
	if !_afeb {
		_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0053\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gfb, _ffag := _ffac.getColorspace(string(*_gafag), _acd)
	if _ffag != nil {
		return _ffag
	}
	_ffac._ebdd.ColorspaceNonStroking = _gfb
	_ffff, _ffag := _ffac.getInitialColor(_gfb)
	if _ffag != nil {
		return _ffag
	}
	_ffac._ebdd.ColorNonStroking = _ffff
	return nil
}

// Add_h appends 'h' operand to the content stream:
// Close the current subpath by adding a line between the current position and the starting position.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_gdf *ContentCreator) Add_h() *ContentCreator {
	_gbfg := ContentStreamOperation{}
	_gbfg.Operand = "\u0068"
	_gdf._bf = append(_gdf._bf, &_gbfg)
	return _gdf
}

type handlerEntry struct {
	Condition HandlerConditionEnum
	Operand   string
	Handler   HandlerFunc
}

// Add_M adds 'M' operand to the content stream: Set the miter limit (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_fac *ContentCreator) Add_M(miterlimit float64) *ContentCreator {
	_cff := ContentStreamOperation{}
	_cff.Operand = "\u004d"
	_cff.Params = _gbg([]float64{miterlimit})
	_fac._bf = append(_fac._bf, &_cff)
	return _fac
}

// Add_B appends 'B' operand to the content stream:
// Fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fee *ContentCreator) Add_B() *ContentCreator {
	_fed := ContentStreamOperation{}
	_fed.Operand = "\u0042"
	_fee._bf = append(_fee._bf, &_fed)
	return _fee
}

// Add_Tj appends 'Tj' operand to the content stream:
// Show a text string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_cbc *ContentCreator) Add_Tj(textstr _cg.PdfObjectString) *ContentCreator {
	_efg := ContentStreamOperation{}
	_efg.Operand = "\u0054\u006a"
	_efg.Params = _eab([]_cg.PdfObjectString{textstr})
	_cbc._bf = append(_cbc._bf, &_efg)
	return _cbc
}
func (_beaf *ContentStreamProcessor) getInitialColor(_gad _dad.PdfColorspace) (_dad.PdfColor, error) {
	switch _dbe := _gad.(type) {
	case *_dad.PdfColorspaceDeviceGray:
		return _dad.NewPdfColorDeviceGray(0.0), nil
	case *_dad.PdfColorspaceDeviceRGB:
		return _dad.NewPdfColorDeviceRGB(0.0, 0.0, 0.0), nil
	case *_dad.PdfColorspaceDeviceCMYK:
		return _dad.NewPdfColorDeviceCMYK(0.0, 0.0, 0.0, 1.0), nil
	case *_dad.PdfColorspaceCalGray:
		return _dad.NewPdfColorCalGray(0.0), nil
	case *_dad.PdfColorspaceCalRGB:
		return _dad.NewPdfColorCalRGB(0.0, 0.0, 0.0), nil
	case *_dad.PdfColorspaceLab:
		_efeg := 0.0
		_abaa := 0.0
		_gbcf := 0.0
		if _dbe.Range[0] > 0 {
			_efeg = _dbe.Range[0]
		}
		if _dbe.Range[2] > 0 {
			_abaa = _dbe.Range[2]
		}
		return _dad.NewPdfColorLab(_efeg, _abaa, _gbcf), nil
	case *_dad.PdfColorspaceICCBased:
		if _dbe.Alternate == nil {
			_ec.Log.Trace("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020-\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0066\u0061\u006c\u006c\u0020\u0062a\u0063\u006b\u0020\u0028\u004e\u0020\u003d\u0020\u0025\u0064\u0029", _dbe.N)
			if _dbe.N == 1 {
				_ec.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079")
				return _beaf.getInitialColor(_dad.NewPdfColorspaceDeviceGray())
			} else if _dbe.N == 3 {
				_ec.Log.Trace("\u0046a\u006c\u006c\u0069\u006eg\u0020\u0062\u0061\u0063\u006b \u0074o\u0020D\u0065\u0076\u0069\u0063\u0065\u0052\u0047B")
				return _beaf.getInitialColor(_dad.NewPdfColorspaceDeviceRGB())
			} else if _dbe.N == 4 {
				_ec.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065C\u004d\u0059\u004b")
				return _beaf.getInitialColor(_dad.NewPdfColorspaceDeviceCMYK())
			} else {
				return nil, _b.New("a\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0049C\u0043")
			}
		}
		return _beaf.getInitialColor(_dbe.Alternate)
	case *_dad.PdfColorspaceSpecialIndexed:
		if _dbe.Base == nil {
			return nil, _b.New("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0062\u0061\u0073e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069f\u0069\u0065\u0064")
		}
		return _beaf.getInitialColor(_dbe.Base)
	case *_dad.PdfColorspaceSpecialSeparation:
		if _dbe.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _beaf.getInitialColor(_dbe.AlternateSpace)
	case *_dad.PdfColorspaceDeviceN:
		if _dbe.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _beaf.getInitialColor(_dbe.AlternateSpace)
	case *_dad.PdfColorspaceSpecialPattern:
		return _dad.NewPdfColorPattern(), nil
	}
	_ec.Log.Debug("Un\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0066\u006f\u0072\u0020\u0075\u006e\u006b\u006e\u006fw\u006e \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065:\u0020\u0025T", _gad)
	return nil, _b.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065")
}
func _edba(_eag _dad.PdfColorspace) bool {
	_, _fbf := _eag.(*_dad.PdfColorspaceSpecialPattern)
	return _fbf
}

// Add_f appends 'f' operand to the content stream:
// Fill the path using the nonzero winding number rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ccd *ContentCreator) Add_f() *ContentCreator {
	_gf := ContentStreamOperation{}
	_gf.Operand = "\u0066"
	_ccd._bf = append(_ccd._bf, &_gf)
	return _ccd
}
func _fada(_gef *ContentStreamInlineImage) (*_cg.DCTEncoder, error) {
	_gdcf := _cg.NewDCTEncoder()
	_gacg := _a.NewReader(_gef._cbdd)
	_fde, _cde := _fd.DecodeConfig(_gacg)
	if _cde != nil {
		_ec.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _cde)
		return nil, _cde
	}
	switch _fde.ColorModel {
	case _bg.RGBAModel:
		_gdcf.BitsPerComponent = 8
		_gdcf.ColorComponents = 3
	case _bg.RGBA64Model:
		_gdcf.BitsPerComponent = 16
		_gdcf.ColorComponents = 3
	case _bg.GrayModel:
		_gdcf.BitsPerComponent = 8
		_gdcf.ColorComponents = 1
	case _bg.Gray16Model:
		_gdcf.BitsPerComponent = 16
		_gdcf.ColorComponents = 1
	case _bg.CMYKModel:
		_gdcf.BitsPerComponent = 8
		_gdcf.ColorComponents = 4
	case _bg.YCbCrModel:
		_gdcf.BitsPerComponent = 8
		_gdcf.ColorComponents = 3
	default:
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_gdcf.Width = _fde.Width
	_gdcf.Height = _fde.Height
	_ec.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _gdcf)
	return _gdcf, nil
}

// Add_Do adds 'Do' operation to the content stream:
// Displays an XObject (image or form) specified by `name`.
//
// See section 8.8 "External Objects" and Table 87 (pp. 209-220 PDF32000_2008).
func (_ae *ContentCreator) Add_Do(name _cg.PdfObjectName) *ContentCreator {
	_fbc := ContentStreamOperation{}
	_fbc.Operand = "\u0044\u006f"
	_fbc.Params = _fcdg([]_cg.PdfObjectName{name})
	_ae._bf = append(_ae._bf, &_fbc)
	return _ae
}

// Add_B_starred appends 'B*' operand to the content stream:
// Fill and then stroke the path (even-odd rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_dee *ContentCreator) Add_B_starred() *ContentCreator {
	_gaf := ContentStreamOperation{}
	_gaf.Operand = "\u0042\u002a"
	_dee._bf = append(_dee._bf, &_gaf)
	return _dee
}

// GraphicStateStack represents a stack of GraphicsState.
type GraphicStateStack []GraphicsState

// Add_Q adds 'Q' operand to the content stream: Pops the most recently stored state from the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dbf *ContentCreator) Add_Q() *ContentCreator {
	_deb := ContentStreamOperation{}
	_deb.Operand = "\u0051"
	_dbf._bf = append(_dbf._bf, &_deb)
	return _dbf
}
func _ecde(_faf *ContentStreamInlineImage, _fdbd *_cg.PdfObjectDictionary) (*_cg.FlateEncoder, error) {
	_ggbg := _cg.NewFlateEncoder()
	if _faf._ebfd != nil {
		_ggbg.SetImage(_faf._ebfd)
	}
	if _fdbd == nil {
		_adg := _faf.DecodeParms
		if _adg != nil {
			_ebce, _caf := _cg.GetDict(_adg)
			if !_caf {
				_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _adg)
				return nil, _e.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_fdbd = _ebce
		}
	}
	if _fdbd == nil {
		return _ggbg, nil
	}
	_ec.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _fdbd.String())
	_fbe := _fdbd.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _fbe == nil {
		_ec.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_cbd, _ecf := _fbe.(*_cg.PdfObjectInteger)
		if !_ecf {
			_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _fbe)
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_ggbg.Predictor = int(*_cbd)
	}
	_fbe = _fdbd.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _fbe != nil {
		_eeeg, _cgb := _fbe.(*_cg.PdfObjectInteger)
		if !_cgb {
			_ec.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_ggbg.BitsPerComponent = int(*_eeeg)
	}
	if _ggbg.Predictor > 1 {
		_ggbg.Columns = 1
		_fbe = _fdbd.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _fbe != nil {
			_aeca, _adgb := _fbe.(*_cg.PdfObjectInteger)
			if !_adgb {
				return nil, _e.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_ggbg.Columns = int(*_aeca)
		}
		_ggbg.Colors = 1
		_fbd := _fdbd.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _fbd != nil {
			_dcag, _bbf := _fbd.(*_cg.PdfObjectInteger)
			if !_bbf {
				return nil, _e.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_ggbg.Colors = int(*_dcag)
		}
	}
	return _ggbg, nil
}
func (_gdgc *ContentStreamProcessor) handleCommand_SC(_eceg *ContentStreamOperation, _faaag *_dad.PdfPageResources) error {
	_bbbc := _gdgc._ebdd.ColorspaceStroking
	if len(_eceg.Params) != _bbbc.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_eceg.Params), _bbbc)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_agga, _fcb := _bbbc.ColorFromPdfObjects(_eceg.Params)
	if _fcb != nil {
		return _fcb
	}
	_gdgc._ebdd.ColorStroking = _agga
	return nil
}

// NewInlineImageFromImage makes a new content stream inline image object from an image.
func NewInlineImageFromImage(img _dad.Image, encoder _cg.StreamEncoder) (*ContentStreamInlineImage, error) {
	if encoder == nil {
		encoder = _cg.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_gdac := ContentStreamInlineImage{}
	if img.ColorComponents == 1 {
		_gdac.ColorSpace = _cg.MakeName("\u0047")
	} else if img.ColorComponents == 3 {
		_gdac.ColorSpace = _cg.MakeName("\u0052\u0047\u0042")
	} else if img.ColorComponents == 4 {
		_gdac.ColorSpace = _cg.MakeName("\u0043\u004d\u0059\u004b")
	} else {
		_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006db\u0065\u0072\u0020o\u0066\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006dpo\u006e\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", img.ColorComponents)
		return nil, _b.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020c\u006fl\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073")
	}
	_gdac.BitsPerComponent = _cg.MakeInteger(img.BitsPerComponent)
	_gdac.Width = _cg.MakeInteger(img.Width)
	_gdac.Height = _cg.MakeInteger(img.Height)
	_cgbb, _cac := encoder.EncodeBytes(img.Data)
	if _cac != nil {
		return nil, _cac
	}
	_gdac._cbdd = _cgbb
	_debe := encoder.GetFilterName()
	if _debe != _cg.StreamEncodingFilterNameRaw {
		_gdac.Filter = _cg.MakeName(_debe)
	}
	return &_gdac, nil
}

// AddHandler adds a new ContentStreamProcessor `handler` of type `condition` for `operand`.
func (_ecgd *ContentStreamProcessor) AddHandler(condition HandlerConditionEnum, operand string, handler HandlerFunc) {
	_fgb := handlerEntry{}
	_fgb.Condition = condition
	_fgb.Operand = operand
	_fgb.Handler = handler
	_ecgd._gdgg = append(_ecgd._gdgg, _fgb)
}
func _daca(_dgdc string) bool { _, _ffgf := _cga[_dgdc]; return _ffgf }
func (_ddd *ContentStreamParser) parseString() (*_cg.PdfObjectString, error) {
	_ddd._gfea.ReadByte()
	var _aed []byte
	_fadee := 1
	for {
		_dcfd, _ddgc := _ddd._gfea.Peek(1)
		if _ddgc != nil {
			return _cg.MakeString(string(_aed)), _ddgc
		}
		if _dcfd[0] == '\\' {
			_ddd._gfea.ReadByte()
			_cgcgd, _ecfb := _ddd._gfea.ReadByte()
			if _ecfb != nil {
				return _cg.MakeString(string(_aed)), _ecfb
			}
			if _cg.IsOctalDigit(_cgcgd) {
				_agg, _gccc := _ddd._gfea.Peek(2)
				if _gccc != nil {
					return _cg.MakeString(string(_aed)), _gccc
				}
				var _dcg []byte
				_dcg = append(_dcg, _cgcgd)
				for _, _efge := range _agg {
					if _cg.IsOctalDigit(_efge) {
						_dcg = append(_dcg, _efge)
					} else {
						break
					}
				}
				_ddd._gfea.Discard(len(_dcg) - 1)
				_ec.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _dcg)
				_fafa, _gccc := _bb.ParseUint(string(_dcg), 8, 32)
				if _gccc != nil {
					return _cg.MakeString(string(_aed)), _gccc
				}
				_aed = append(_aed, byte(_fafa))
				continue
			}
			switch _cgcgd {
			case 'n':
				_aed = append(_aed, '\n')
			case 'r':
				_aed = append(_aed, '\r')
			case 't':
				_aed = append(_aed, '\t')
			case 'b':
				_aed = append(_aed, '\b')
			case 'f':
				_aed = append(_aed, '\f')
			case '(':
				_aed = append(_aed, '(')
			case ')':
				_aed = append(_aed, ')')
			case '\\':
				_aed = append(_aed, '\\')
			}
			continue
		} else if _dcfd[0] == '(' {
			_fadee++
		} else if _dcfd[0] == ')' {
			_fadee--
			if _fadee == 0 {
				_ddd._gfea.ReadByte()
				break
			}
		}
		_bcf, _ := _ddd._gfea.ReadByte()
		_aed = append(_aed, _bcf)
	}
	return _cg.MakeString(string(_aed)), nil
}
func (_gcgb *ContentStreamProcessor) handleCommand_G(_cgba *ContentStreamOperation, _gec *_dad.PdfPageResources) error {
	_geffg := _dad.NewPdfColorspaceDeviceGray()
	if len(_cgba.Params) != _geffg.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cgba.Params), _geffg)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bedd, _aaca := _geffg.ColorFromPdfObjects(_cgba.Params)
	if _aaca != nil {
		return _aaca
	}
	_gcgb._ebdd.ColorspaceStroking = _geffg
	_gcgb._ebdd.ColorStroking = _bedd
	return nil
}

// Add_f_starred appends 'f*' operand to the content stream.
// f*: Fill the path using the even-odd rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_adc *ContentCreator) Add_f_starred() *ContentCreator {
	_afd := ContentStreamOperation{}
	_afd.Operand = "\u0066\u002a"
	_adc._bf = append(_adc._bf, &_afd)
	return _adc
}
func (_geaa *ContentStreamParser) parseNull() (_cg.PdfObjectNull, error) {
	_, _gcfb := _geaa._gfea.Discard(4)
	return _cg.PdfObjectNull{}, _gcfb
}
func (_gfgc *ContentStreamParser) parseObject() (_cfec _cg.PdfObject, _cbddb bool, _fedce error) {
	_gfgc.skipSpaces()
	for {
		_cgef, _aee := _gfgc._gfea.Peek(2)
		if _aee != nil {
			return nil, false, _aee
		}
		_ec.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_cgef))
		if _cgef[0] == '%' {
			_gfgc.skipComments()
			continue
		} else if _cgef[0] == '/' {
			_debd, _adca := _gfgc.parseName()
			_ec.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _debd)
			return &_debd, false, _adca
		} else if _cgef[0] == '(' {
			_ec.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_dcfe, _aef := _gfgc.parseString()
			return _dcfe, false, _aef
		} else if _cgef[0] == '<' && _cgef[1] != '<' {
			_ec.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0053\u0074\u0072\u0069\u006e\u0067\u0021")
			_dfa, _bbaa := _gfgc.parseHexString()
			return _dfa, false, _bbaa
		} else if _cgef[0] == '[' {
			_ec.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_daee, _gaae := _gfgc.parseArray()
			return _daee, false, _gaae
		} else if _cg.IsFloatDigit(_cgef[0]) || (_cgef[0] == '-' && _cg.IsFloatDigit(_cgef[1])) || (_cgef[0] == '+' && _cg.IsFloatDigit(_cgef[1])) {
			_ec.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_effe, _edb := _gfgc.parseNumber()
			return _effe, false, _edb
		} else if _cgef[0] == '<' && _cgef[1] == '<' {
			_cfbe, _cgcc := _gfgc.parseDict()
			return _cfbe, false, _cgcc
		} else {
			_ec.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_cgef, _ = _gfgc._gfea.Peek(5)
			_gbfd := string(_cgef)
			_ec.Log.Trace("\u0063\u006f\u006e\u0074\u0020\u0050\u0065\u0065\u006b\u0020\u0073\u0074r\u003a\u0020\u0025\u0073", _gbfd)
			if (len(_gbfd) > 3) && (_gbfd[:4] == "\u006e\u0075\u006c\u006c") {
				_agabd, _dfge := _gfgc.parseNull()
				return &_agabd, false, _dfge
			} else if (len(_gbfd) > 4) && (_gbfd[:5] == "\u0066\u0061\u006cs\u0065") {
				_bdc, _gaac := _gfgc.parseBool()
				return &_bdc, false, _gaac
			} else if (len(_gbfd) > 3) && (_gbfd[:4] == "\u0074\u0072\u0075\u0065") {
				_gffd, _daaa := _gfgc.parseBool()
				return &_gffd, false, _daaa
			}
			_fbba, _gegb := _gfgc.parseOperand()
			if _gegb != nil {
				return _fbba, false, _gegb
			}
			if len(_fbba.String()) < 1 {
				return _fbba, false, ErrInvalidOperand
			}
			return _fbba, true, nil
		}
	}
}

// Add_re appends 're' operand to the content stream:
// Append a rectangle to the current path as a complete subpath, with lower left corner (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cgc *ContentCreator) Add_re(x, y, width, height float64) *ContentCreator {
	_aadc := ContentStreamOperation{}
	_aadc.Operand = "\u0072\u0065"
	_aadc.Params = _gbg([]float64{x, y, width, height})
	_cgc._bf = append(_cgc._bf, &_aadc)
	return _cgc
}

// Add_CS appends 'CS' operand to the content stream:
// Set the current colorspace for stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cfg *ContentCreator) Add_CS(name _cg.PdfObjectName) *ContentCreator {
	_bbb := ContentStreamOperation{}
	_bbb.Operand = "\u0043\u0053"
	_bbb.Params = _fcdg([]_cg.PdfObjectName{name})
	_cfg._bf = append(_cfg._bf, &_bbb)
	return _cfg
}

// Add_Ts appends 'Ts' operand to the content stream:
// Set text rise.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_eeec *ContentCreator) Add_Ts(rise float64) *ContentCreator {
	_ddg := ContentStreamOperation{}
	_ddg.Operand = "\u0054\u0073"
	_ddg.Params = _gbg([]float64{rise})
	_eeec._bf = append(_eeec._bf, &_ddg)
	return _eeec
}

// Parse parses all commands in content stream, returning a list of operation data.
func (_eeed *ContentStreamParser) Parse() (*ContentStreamOperations, error) {
	_cafa := ContentStreamOperations{}
	for {
		_ccf := ContentStreamOperation{}
		for {
			_cge, _dbad, _eae := _eeed.parseObject()
			if _eae != nil {
				if _eae == _ce.EOF {
					return &_cafa, nil
				}
				return &_cafa, _eae
			}
			if _dbad {
				_ccf.Operand, _ = _cg.GetStringVal(_cge)
				_cafa = append(_cafa, &_ccf)
				break
			} else {
				_ccf.Params = append(_ccf.Params, _cge)
			}
		}
		if _ccf.Operand == "\u0042\u0049" {
			_bffb, _dfb := _eeed.ParseInlineImage()
			if _dfb != nil {
				return &_cafa, _dfb
			}
			_ccf.Params = append(_ccf.Params, _bffb)
		}
	}
}

// ContentCreator is a builder for PDF content streams.
type ContentCreator struct{ _bf ContentStreamOperations }

// Add_b_starred appends 'b*' operand to the content stream:
// Close, fill and then stroke the path (even-odd winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fdg *ContentCreator) Add_b_starred() *ContentCreator {
	_fga := ContentStreamOperation{}
	_fga.Operand = "\u0062\u002a"
	_fdg._bf = append(_fdg._bf, &_fga)
	return _fdg
}

var (
	ErrInvalidOperand = _b.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
	ErrEarlyExit      = _b.New("\u0074\u0065\u0072\u006di\u006e\u0061\u0074\u0065\u0020\u0070\u0072\u006f\u0063\u0065s\u0073 \u0065\u0061\u0072\u006c\u0079\u0020\u0065x\u0069\u0074")
)

func (_fdbc *ContentStreamProcessor) handleCommand_sc(_cfde *ContentStreamOperation, _fcdf *_dad.PdfPageResources) error {
	_bcc := _fdbc._ebdd.ColorspaceNonStroking
	if !_edba(_bcc) {
		if len(_cfde.Params) != _bcc.GetNumComponents() {
			_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cfde.Params), _bcc)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_bdcb, _bgd := _bcc.ColorFromPdfObjects(_cfde.Params)
	if _bgd != nil {
		return _bgd
	}
	_fdbc._ebdd.ColorNonStroking = _bdcb
	return nil
}
func (_baab *ContentStreamProcessor) handleCommand_rg(_abga *ContentStreamOperation, _fcg *_dad.PdfPageResources) error {
	_fdeg := _dad.NewPdfColorspaceDeviceRGB()
	if len(_abga.Params) != _fdeg.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_abga.Params), _fdeg)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_abee, _caba := _fdeg.ColorFromPdfObjects(_abga.Params)
	if _caba != nil {
		return _caba
	}
	_baab._ebdd.ColorspaceNonStroking = _fdeg
	_baab._ebdd.ColorNonStroking = _abee
	return nil
}

// Add_m adds 'm' operand to the content stream: Move the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_fe *ContentCreator) Add_m(x, y float64) *ContentCreator {
	_gbc := ContentStreamOperation{}
	_gbc.Operand = "\u006d"
	_gbc.Params = _gbg([]float64{x, y})
	_fe._bf = append(_fe._bf, &_gbc)
	return _fe
}

// Add_scn appends 'scn' operand to the content stream:
// Same as SC but for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_aeb *ContentCreator) Add_scn(c ...float64) *ContentCreator {
	_fce := ContentStreamOperation{}
	_fce.Operand = "\u0073\u0063\u006e"
	_fce.Params = _gbg(c)
	_aeb._bf = append(_aeb._bf, &_fce)
	return _aeb
}

// SetNonStrokingColor sets the non-stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_adcf *ContentCreator) SetNonStrokingColor(color _dad.PdfColor) *ContentCreator {
	switch _cffd := color.(type) {
	case *_dad.PdfColorDeviceGray:
		_adcf.Add_g(_cffd.Val())
	case *_dad.PdfColorDeviceRGB:
		_adcf.Add_rg(_cffd.R(), _cffd.G(), _cffd.B())
	case *_dad.PdfColorDeviceCMYK:
		_adcf.Add_k(_cffd.C(), _cffd.M(), _cffd.Y(), _cffd.K())
	case *_dad.PdfColorPatternType2:
		_adcf.Add_cs(*_cg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_adcf.Add_scn_pattern(_cffd.PatternName)
	case *_dad.PdfColorPatternType3:
		_adcf.Add_cs(*_cg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_adcf.Add_scn_pattern(_cffd.PatternName)
	default:
		_ec.Log.Debug("\u0053\u0065\u0074N\u006f\u006e\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006f\u006c\u006f\u0072\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020c\u006f\u006c\u006f\u0072\u003a\u0020\u0025\u0054", _cffd)
	}
	return _adcf
}

// Add_TD appends 'TD' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_dec *ContentCreator) Add_TD(tx, ty float64) *ContentCreator {
	_gac := ContentStreamOperation{}
	_gac.Operand = "\u0054\u0044"
	_gac.Params = _gbg([]float64{tx, ty})
	_dec._bf = append(_dec._bf, &_gac)
	return _dec
}

// Add_cm adds 'cm' operation to the content stream: Modifies the current transformation matrix (ctm)
// of the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_eb *ContentCreator) Add_cm(a, b, c, d, e, f float64) *ContentCreator {
	_dd := ContentStreamOperation{}
	_dd.Operand = "\u0063\u006d"
	_dd.Params = _gbg([]float64{a, b, c, d, e, f})
	_eb._bf = append(_eb._bf, &_dd)
	return _eb
}

const (
	HandlerConditionEnumOperand HandlerConditionEnum = iota
	HandlerConditionEnumAllOperands
)

// WrapIfNeeded wraps the entire contents within q ... Q.  If unbalanced, then adds extra Qs at the end.
// Only does if needed. Ensures that when adding new content, one start with all states
// in the default condition.
func (_aaa *ContentStreamOperations) WrapIfNeeded() *ContentStreamOperations {
	if len(*_aaa) == 0 {
		return _aaa
	}
	if _aaa.isWrapped() {
		return _aaa
	}
	*_aaa = append([]*ContentStreamOperation{{Operand: "\u0071"}}, *_aaa...)
	_gb := 0
	for _, _ef := range *_aaa {
		if _ef.Operand == "\u0071" {
			_gb++
		} else if _ef.Operand == "\u0051" {
			_gb--
		}
	}
	for _gb > 0 {
		*_aaa = append(*_aaa, &ContentStreamOperation{Operand: "\u0051"})
		_gb--
	}
	return _aaa
}

// Add_W appends 'W' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (nonzero winding rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_bfb *ContentCreator) Add_W() *ContentCreator {
	_caa := ContentStreamOperation{}
	_caa.Operand = "\u0057"
	_bfb._bf = append(_bfb._bf, &_caa)
	return _bfb
}

// Add_BMC appends 'BMC' operand to the content stream:
// Begins a marked-content sequence terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_ggf *ContentCreator) Add_BMC(tag _cg.PdfObjectName) *ContentCreator {
	_gda := ContentStreamOperation{}
	_gda.Operand = "\u0042\u004d\u0043"
	_gda.Params = _fcdg([]_cg.PdfObjectName{tag})
	_ggf._bf = append(_ggf._bf, &_gda)
	return _ggf
}

// Add_scn_pattern appends 'scn' operand to the content stream for pattern `name`:
// scn with name attribute (for pattern). Syntax: c1 ... cn name scn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fabc *ContentCreator) Add_scn_pattern(name _cg.PdfObjectName, c ...float64) *ContentCreator {
	_dcf := ContentStreamOperation{}
	_dcf.Operand = "\u0073\u0063\u006e"
	_dcf.Params = _gbg(c)
	_dcf.Params = append(_dcf.Params, _cg.MakeName(string(name)))
	_fabc._bf = append(_fabc._bf, &_dcf)
	return _fabc
}

// GetEncoder returns the encoder of the inline image.
func (_ace *ContentStreamInlineImage) GetEncoder() (_cg.StreamEncoder, error) { return _gea(_ace) }

// Push pushes `gs` on the `gsStack`.
func (_gcbd *GraphicStateStack) Push(gs GraphicsState) { *_gcbd = append(*_gcbd, gs) }

// Add_w adds 'w' operand to the content stream, which sets the line width.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_abc *ContentCreator) Add_w(lineWidth float64) *ContentCreator {
	_cfd := ContentStreamOperation{}
	_cfd.Operand = "\u0077"
	_cfd.Params = _gbg([]float64{lineWidth})
	_abc._bf = append(_abc._bf, &_cfd)
	return _abc
}
func _gbg(_acg []float64) []_cg.PdfObject {
	var _eefg []_cg.PdfObject
	for _, _cfgf := range _acg {
		_eefg = append(_eefg, _cg.MakeFloat(_cfgf))
	}
	return _eefg
}

// Add_sh appends 'sh' operand to the content stream:
// Paints the shape and colour shading described by a shading dictionary specified by `name`,
// subject to the current clipping path
//
// See section 8.7.4 "Shading Patterns" and Table 77 (p. 190 PDF32000_2008).
func (_cag *ContentCreator) Add_sh(name _cg.PdfObjectName) *ContentCreator {
	_aab := ContentStreamOperation{}
	_aab.Operand = "\u0073\u0068"
	_aab.Params = _fcdg([]_cg.PdfObjectName{name})
	_cag._bf = append(_cag._bf, &_aab)
	return _cag
}

// ContentStreamProcessor defines a data structure and methods for processing a content stream, keeping track of the
// current graphics state, and allowing external handlers to define their own functions as a part of the processing,
// for example rendering or extracting certain information.
type ContentStreamProcessor struct {
	_bfa  GraphicStateStack
	_aedb []*ContentStreamOperation
	_ebdd GraphicsState
	_gdgg []handlerEntry
	_bae  int
}

// ContentStreamOperations is a slice of ContentStreamOperations.
type ContentStreamOperations []*ContentStreamOperation

// Add_b appends 'b' operand to the content stream:
// Close, fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fade *ContentCreator) Add_b() *ContentCreator {
	_cdc := ContentStreamOperation{}
	_cdc.Operand = "\u0062"
	_fade._bf = append(_fade._bf, &_cdc)
	return _fade
}

// Add_k appends 'k' operand to the content stream:
// Same as K but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fccc *ContentCreator) Add_k(c, m, y, k float64) *ContentCreator {
	_bec := ContentStreamOperation{}
	_bec.Operand = "\u006b"
	_bec.Params = _gbg([]float64{c, m, y, k})
	_fccc._bf = append(_fccc._bf, &_bec)
	return _fccc
}

// ExtractText parses and extracts all text data in content streams and returns as a string.
// Does not take into account Encoding table, the output is simply the character codes.
//
// Deprecated: More advanced text extraction is offered in package extractor with character encoding support.
func (_gbb *ContentStreamParser) ExtractText() (string, error) {
	_be, _bad := _gbb.Parse()
	if _bad != nil {
		return "", _bad
	}
	_fa := false
	_ecb, _fc := float64(-1), float64(-1)
	_dc := ""
	for _, _ge := range *_be {
		if _ge.Operand == "\u0042\u0054" {
			_fa = true
		} else if _ge.Operand == "\u0045\u0054" {
			_fa = false
		}
		if _ge.Operand == "\u0054\u0064" || _ge.Operand == "\u0054\u0044" || _ge.Operand == "\u0054\u002a" {
			_dc += "\u000a"
		}
		if _ge.Operand == "\u0054\u006d" {
			if len(_ge.Params) != 6 {
				continue
			}
			_bac, _ged := _ge.Params[4].(*_cg.PdfObjectFloat)
			if !_ged {
				_ceb, _bag := _ge.Params[4].(*_cg.PdfObjectInteger)
				if !_bag {
					continue
				}
				_bac = _cg.MakeFloat(float64(*_ceb))
			}
			_ab, _ged := _ge.Params[5].(*_cg.PdfObjectFloat)
			if !_ged {
				_eg, _fab := _ge.Params[5].(*_cg.PdfObjectInteger)
				if !_fab {
					continue
				}
				_ab = _cg.MakeFloat(float64(*_eg))
			}
			if _fc == -1 {
				_fc = float64(*_ab)
			} else if _fc > float64(*_ab) {
				_dc += "\u000a"
				_ecb = float64(*_bac)
				_fc = float64(*_ab)
				continue
			}
			if _ecb == -1 {
				_ecb = float64(*_bac)
			} else if _ecb < float64(*_bac) {
				_dc += "\u0009"
				_ecb = float64(*_bac)
			}
		}
		if _fa && _ge.Operand == "\u0054\u004a" {
			if len(_ge.Params) < 1 {
				continue
			}
			_gcd, _ggd := _ge.Params[0].(*_cg.PdfObjectArray)
			if !_ggd {
				return "", _e.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0020\u0074y\u0070\u0065\u002c\u0020\u006e\u006f\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _ge.Params[0])
			}
			for _, _egd := range _gcd.Elements() {
				switch _cebe := _egd.(type) {
				case *_cg.PdfObjectString:
					_dc += _cebe.Str()
				case *_cg.PdfObjectFloat:
					if *_cebe < -100 {
						_dc += "\u0020"
					}
				case *_cg.PdfObjectInteger:
					if *_cebe < -100 {
						_dc += "\u0020"
					}
				}
			}
		} else if _fa && _ge.Operand == "\u0054\u006a" {
			if len(_ge.Params) < 1 {
				continue
			}
			_aad, _ecg := _ge.Params[0].(*_cg.PdfObjectString)
			if !_ecg {
				return "", _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0074\u0079p\u0065\u002c\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067 \u0028\u0025\u0054\u0029", _ge.Params[0])
			}
			_dc += _aad.Str()
		}
	}
	return _dc, nil
}

// SetStrokingColor sets the stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_ggbe *ContentCreator) SetStrokingColor(color _dad.PdfColor) *ContentCreator {
	switch _dfd := color.(type) {
	case *_dad.PdfColorDeviceGray:
		_ggbe.Add_G(_dfd.Val())
	case *_dad.PdfColorDeviceRGB:
		_ggbe.Add_RG(_dfd.R(), _dfd.G(), _dfd.B())
	case *_dad.PdfColorDeviceCMYK:
		_ggbe.Add_K(_dfd.C(), _dfd.M(), _dfd.Y(), _dfd.K())
	case *_dad.PdfColorPatternType2:
		_ggbe.Add_CS(*_cg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_ggbe.Add_SCN_pattern(_dfd.PatternName)
	case *_dad.PdfColorPatternType3:
		_ggbe.Add_CS(*_cg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_ggbe.Add_SCN_pattern(_dfd.PatternName)
	default:
		_ec.Log.Debug("\u0053\u0065\u0074\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006fl\u006f\u0072\u003a\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006fr\u003a\u0020\u0025\u0054", _dfd)
	}
	return _ggbe
}

// Add_ET appends 'ET' operand to the content stream:
// End text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_aae *ContentCreator) Add_ET() *ContentCreator {
	_dgd := ContentStreamOperation{}
	_dgd.Operand = "\u0045\u0054"
	_aae._bf = append(_aae._bf, &_dgd)
	return _aae
}

// Add_d adds 'd' operand to the content stream: Set the line dash pattern.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dea *ContentCreator) Add_d(dashArray []int64, dashPhase int64) *ContentCreator {
	_gd := ContentStreamOperation{}
	_gd.Operand = "\u0064"
	_gd.Params = []_cg.PdfObject{}
	_gd.Params = append(_gd.Params, _cg.MakeArrayFromIntegers64(dashArray))
	_gd.Params = append(_gd.Params, _cg.MakeInteger(dashPhase))
	_dea._bf = append(_dea._bf, &_gd)
	return _dea
}

// AddOperand adds a specified operand.
func (_eee *ContentCreator) AddOperand(op ContentStreamOperation) *ContentCreator {
	_eee._bf = append(_eee._bf, &op)
	return _eee
}

// Pop pops and returns the topmost GraphicsState off the `gsStack`.
func (_ebff *GraphicStateStack) Pop() GraphicsState {
	_bgg := (*_ebff)[len(*_ebff)-1]
	*_ebff = (*_ebff)[:len(*_ebff)-1]
	return _bgg
}

// Transform returns coordinates x, y transformed by the CTM.
func (_abag *GraphicsState) Transform(x, y float64) (float64, float64) {
	return _abag.CTM.Transform(x, y)
}
func (_dgc *ContentStreamInlineImage) String() string {
	_dag := _e.Sprintf("I\u006el\u0069\u006e\u0065\u0049\u006d\u0061\u0067\u0065(\u006c\u0065\u006e\u003d%d\u0029\u000a", len(_dgc._cbdd))
	if _dgc.BitsPerComponent != nil {
		_dag += "\u002d\u0020\u0042\u0050\u0043\u0020" + _dgc.BitsPerComponent.WriteString() + "\u000a"
	}
	if _dgc.ColorSpace != nil {
		_dag += "\u002d\u0020\u0043S\u0020" + _dgc.ColorSpace.WriteString() + "\u000a"
	}
	if _dgc.Decode != nil {
		_dag += "\u002d\u0020\u0044\u0020" + _dgc.Decode.WriteString() + "\u000a"
	}
	if _dgc.DecodeParms != nil {
		_dag += "\u002d\u0020\u0044P\u0020" + _dgc.DecodeParms.WriteString() + "\u000a"
	}
	if _dgc.Filter != nil {
		_dag += "\u002d\u0020\u0046\u0020" + _dgc.Filter.WriteString() + "\u000a"
	}
	if _dgc.Height != nil {
		_dag += "\u002d\u0020\u0048\u0020" + _dgc.Height.WriteString() + "\u000a"
	}
	if _dgc.ImageMask != nil {
		_dag += "\u002d\u0020\u0049M\u0020" + _dgc.ImageMask.WriteString() + "\u000a"
	}
	if _dgc.Intent != nil {
		_dag += "\u002d \u0049\u006e\u0074\u0065\u006e\u0074 " + _dgc.Intent.WriteString() + "\u000a"
	}
	if _dgc.Interpolate != nil {
		_dag += "\u002d\u0020\u0049\u0020" + _dgc.Interpolate.WriteString() + "\u000a"
	}
	if _dgc.Width != nil {
		_dag += "\u002d\u0020\u0057\u0020" + _dgc.Width.WriteString() + "\u000a"
	}
	return _dag
}

// Add_rg appends 'rg' operand to the content stream:
// Same as RG but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_agde *ContentCreator) Add_rg(r, g, b float64) *ContentCreator {
	_bff := ContentStreamOperation{}
	_bff.Operand = "\u0072\u0067"
	_bff.Params = _gbg([]float64{r, g, b})
	_agde._bf = append(_agde._bf, &_bff)
	return _agde
}

// Operations returns the list of operations.
func (_dg *ContentCreator) Operations() *ContentStreamOperations { return &_dg._bf }

// Add_SCN appends 'SCN' operand to the content stream:
// Same as SC but supports more colorspaces.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_afe *ContentCreator) Add_SCN(c ...float64) *ContentCreator {
	_bbe := ContentStreamOperation{}
	_bbe.Operand = "\u0053\u0043\u004e"
	_bbe.Params = _gbg(c)
	_afe._bf = append(_afe._bf, &_bbe)
	return _afe
}

// ToImage exports the inline image to Image which can be transformed or exported easily.
// Page resources are needed to look up colorspace information.
func (_becf *ContentStreamInlineImage) ToImage(resources *_dad.PdfPageResources) (*_dad.Image, error) {
	_ebcc, _bdae := _becf.toImageBase(resources)
	if _bdae != nil {
		return nil, _bdae
	}
	_fegc, _bdae := _gea(_becf)
	if _bdae != nil {
		return nil, _bdae
	}
	_bea, _cgfb := _cg.GetDict(_becf.DecodeParms)
	if _cgfb {
		_fegc.UpdateParams(_bea)
	}
	_ec.Log.Trace("\u0065n\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u0020\u0025\u0054", _fegc, _fegc)
	_ec.Log.Trace("\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u002b\u0076", _becf)
	_ece, _bdae := _fegc.DecodeBytes(_becf._cbdd)
	if _bdae != nil {
		return nil, _bdae
	}
	_ffad := &_dad.Image{Width: int64(_ebcc.Width), Height: int64(_ebcc.Height), BitsPerComponent: int64(_ebcc.BitsPerComponent), ColorComponents: _ebcc.ColorComponents, Data: _ece}
	if len(_ebcc.Decode) > 0 {
		for _gacd := 0; _gacd < len(_ebcc.Decode); _gacd++ {
			_ebcc.Decode[_gacd] *= float64((int(1) << uint(_ebcc.BitsPerComponent)) - 1)
		}
		_ffad.SetDecode(_ebcc.Decode)
	}
	return _ffad, nil
}
func (_gacb *ContentStreamParser) parseNumber() (_cg.PdfObject, error) {
	return _cg.ParseNumber(_gacb._gfea)
}
func _ecfc(_bcd []int64) []_cg.PdfObject {
	var _fbg []_cg.PdfObject
	for _, _ggcdg := range _bcd {
		_fbg = append(_fbg, _cg.MakeInteger(_ggcdg))
	}
	return _fbg
}
func (_afea *ContentStreamProcessor) handleCommand_cm(_faaf *ContentStreamOperation, _gedaf *_dad.PdfPageResources) error {
	if len(_faaf.Params) != 6 {
		_ec.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u006d\u003a\u0020\u0025\u0064", len(_faaf.Params))
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_edf, _dbbd := _cg.GetNumbersAsFloat(_faaf.Params)
	if _dbbd != nil {
		return _dbbd
	}
	_eaaf := _da.NewMatrix(_edf[0], _edf[1], _edf[2], _edf[3], _edf[4], _edf[5])
	_afea._ebdd.CTM.Concat(_eaaf)
	return nil
}

// Add_K appends 'K' operand to the content stream:
// Set the stroking colorspace to DeviceCMYK and sets the c,m,y,k color (0-1 each component).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ebc *ContentCreator) Add_K(c, m, y, k float64) *ContentCreator {
	_gdc := ContentStreamOperation{}
	_gdc.Operand = "\u004b"
	_gdc.Params = _gbg([]float64{c, m, y, k})
	_ebc._bf = append(_ebc._bf, &_gdc)
	return _ebc
}

// HasUnclosedQ checks if all the `q` operator is properly closed by `Q` operator.
func (_ag *ContentStreamOperations) HasUnclosedQ() bool {
	_cec := 0
	for _, _cf := range *_ag {
		if _cf.Operand == "\u0071" {
			_cec++
		} else if _cf.Operand == "\u0051" {
			_cec--
		}
	}
	return _cec != 0
}

// ContentStreamParser represents a content stream parser for parsing content streams in PDFs.
type ContentStreamParser struct{ _gfea *_f.Reader }

// Add_TJ appends 'TJ' operand to the content stream:
// Show one or more text string. Array of numbers (displacement) and strings.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_ggce *ContentCreator) Add_TJ(vals ..._cg.PdfObject) *ContentCreator {
	_aaea := ContentStreamOperation{}
	_aaea.Operand = "\u0054\u004a"
	_aaea.Params = []_cg.PdfObject{_cg.MakeArray(vals...)}
	_ggce._bf = append(_ggce._bf, &_aaea)
	return _ggce
}
func _fcdg(_fcefg []_cg.PdfObjectName) []_cg.PdfObject {
	var _ffab []_cg.PdfObject
	for _, _geaaf := range _fcefg {
		_ffab = append(_ffab, _cg.MakeName(string(_geaaf)))
	}
	return _ffab
}

// NewContentStreamProcessor returns a new ContentStreamProcessor for operations `ops`.
func NewContentStreamProcessor(ops []*ContentStreamOperation) *ContentStreamProcessor {
	_abgg := ContentStreamProcessor{}
	_abgg._bfa = GraphicStateStack{}
	_gfgf := GraphicsState{}
	_abgg._ebdd = _gfgf
	_abgg._gdgg = []handlerEntry{}
	_abgg._bae = 0
	_abgg._aedb = ops
	return &_abgg
}

// Add_Tr appends 'Tr' operand to the content stream:
// Set text rendering mode.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_fcd *ContentCreator) Add_Tr(render int64) *ContentCreator {
	_aga := ContentStreamOperation{}
	_aga.Operand = "\u0054\u0072"
	_aga.Params = _ecfc([]int64{render})
	_fcd._bf = append(_fcd._bf, &_aga)
	return _fcd
}
func (_gedag *ContentStreamParser) skipComments() error {
	if _, _fbeb := _gedag.skipSpaces(); _fbeb != nil {
		return _fbeb
	}
	_becd := true
	for {
		_cedb, _acb := _gedag._gfea.Peek(1)
		if _acb != nil {
			_ec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _acb.Error())
			return _acb
		}
		if _becd && _cedb[0] != '%' {
			return nil
		}
		_becd = false
		if (_cedb[0] != '\r') && (_cedb[0] != '\n') {
			_gedag._gfea.ReadByte()
		} else {
			break
		}
	}
	return _gedag.skipComments()
}
func _eab(_dade []_cg.PdfObjectString) []_cg.PdfObject {
	var _aadcc []_cg.PdfObject
	for _, _bebd := range _dade {
		_aadcc = append(_aadcc, _cg.MakeString(_bebd.Str()))
	}
	return _aadcc
}

// GraphicsState is a basic graphics state implementation for PDF processing.
// Initially only implementing and tracking a portion of the information specified. Easy to add more.
type GraphicsState struct {
	ColorspaceStroking    _dad.PdfColorspace
	ColorspaceNonStroking _dad.PdfColorspace
	ColorStroking         _dad.PdfColor
	ColorNonStroking      _dad.PdfColor
	CTM                   _da.Matrix
}

func _gea(_gdfg *ContentStreamInlineImage) (_cg.StreamEncoder, error) {
	if _gdfg.Filter == nil {
		return _cg.NewRawEncoder(), nil
	}
	_fbb, _ecdd := _gdfg.Filter.(*_cg.PdfObjectName)
	if !_ecdd {
		_dba, _agf := _gdfg.Filter.(*_cg.PdfObjectArray)
		if !_agf {
			return nil, _e.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _dba.Len() == 0 {
			return _cg.NewRawEncoder(), nil
		}
		if _dba.Len() != 1 {
			_gag, _eec := _gba(_gdfg)
			if _eec != nil {
				_ec.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _eec)
				return nil, _eec
			}
			_ec.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _gag)
			return _gag, nil
		}
		_gab := _dba.Get(0)
		_fbb, _agf = _gab.(*_cg.PdfObjectName)
		if !_agf {
			return nil, _e.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	switch *_fbb {
	case "\u0041\u0048\u0078", "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _cg.NewASCIIHexEncoder(), nil
	case "\u0041\u0038\u0035", "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065":
		return _cg.NewASCII85Encoder(), nil
	case "\u0044\u0043\u0054", "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e":
		return _fada(_gdfg)
	case "\u0046\u006c", "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065":
		return _ecde(_gdfg, nil)
	case "\u004c\u005a\u0057", "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e":
		return _fec(_gdfg, nil)
	case "\u0043\u0043\u0046", "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _cg.NewCCITTFaxEncoder(), nil
	case "\u0052\u004c", "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065":
		return _cg.NewRunLengthEncoder(), nil
	default:
		_ec.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0069\u006d\u0061\u0067\u0065\u0020\u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u003a\u0020\u0025\u0073", *_fbb)
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006el\u0069n\u0065 \u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
}
func (_cab *ContentStreamProcessor) handleCommand_CS(_bgcb *ContentStreamOperation, _gdcg *_dad.PdfPageResources) error {
	if len(_bgcb.Params) < 1 {
		_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_bgcb.Params) > 1 {
		_ec.Log.Debug("\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_dcgd, _gcba := _bgcb.Params[0].(*_cg.PdfObjectName)
	if !_gcba {
		_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020c\u0073\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_cccd, _eeef := _cab.getColorspace(string(*_dcgd), _gdcg)
	if _eeef != nil {
		return _eeef
	}
	_cab._ebdd.ColorspaceStroking = _cccd
	_cfca, _eeef := _cab.getInitialColor(_cccd)
	if _eeef != nil {
		return _eeef
	}
	_cab._ebdd.ColorStroking = _cfca
	return nil
}

// Bytes converts a set of content stream operations to a content stream byte presentation,
// i.e. the kind that can be stored as a PDF stream or string format.
func (_db *ContentStreamOperations) Bytes() []byte {
	var _bd _a.Buffer
	for _, _gg := range *_db {
		if _gg == nil {
			continue
		}
		if _gg.Operand == "\u0042\u0049" {
			_bd.WriteString(_gg.Operand + "\u000a")
			_bd.WriteString(_gg.Params[0].WriteString())
		} else {
			for _, _agd := range _gg.Params {
				_bd.WriteString(_agd.WriteString())
				_bd.WriteString("\u0020")
			}
			_bd.WriteString(_gg.Operand + "\u000a")
		}
	}
	return _bd.Bytes()
}

// Add_quote appends "'" operand to the content stream:
// Move to next line and show a string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_cbbb *ContentCreator) Add_quote(textstr _cg.PdfObjectString) *ContentCreator {
	_dbbg := ContentStreamOperation{}
	_dbbg.Operand = "\u0027"
	_dbbg.Params = _eab([]_cg.PdfObjectString{textstr})
	_cbbb._bf = append(_cbbb._bf, &_dbbg)
	return _cbbb
}
func (_gfc *ContentStreamParser) parseBool() (_cg.PdfObjectBool, error) {
	_cfc, _gffg := _gfc._gfea.Peek(4)
	if _gffg != nil {
		return _cg.PdfObjectBool(false), _gffg
	}
	if (len(_cfc) >= 4) && (string(_cfc[:4]) == "\u0074\u0072\u0075\u0065") {
		_gfc._gfea.Discard(4)
		return _cg.PdfObjectBool(true), nil
	}
	_cfc, _gffg = _gfc._gfea.Peek(5)
	if _gffg != nil {
		return _cg.PdfObjectBool(false), _gffg
	}
	if (len(_cfc) >= 5) && (string(_cfc[:5]) == "\u0066\u0061\u006cs\u0065") {
		_gfc._gfea.Discard(5)
		return _cg.PdfObjectBool(false), nil
	}
	return _cg.PdfObjectBool(false), _b.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// HandlerFunc is the function syntax that the ContentStreamProcessor handler must implement.
type HandlerFunc func(_eaff *ContentStreamOperation, _cedc GraphicsState, _gee *_dad.PdfPageResources) error

// Add_BDC appends 'BDC' operand to the content stream:
// Begins a marked-content sequence with an associated property list terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
// `propertyList` shall be a dictionary containing the properties of the
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_geda *ContentCreator) Add_BDC(tag _cg.PdfObjectName, propertyList map[string]_cg.PdfObject) *ContentCreator {
	_faaa := ContentStreamOperation{}
	_faaa.Operand = "\u0042\u0044\u0043"
	_faaa.Params = _fcdg([]_cg.PdfObjectName{tag})
	if len(propertyList) > 0 {
		_faaa.Params = append(_faaa.Params, _cg.MakeDictMap(propertyList))
	}
	_geda._bf = append(_geda._bf, &_faaa)
	return _geda
}

// Add_gs adds 'gs' operand to the content stream: Set the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_gcf *ContentCreator) Add_gs(dictName _cg.PdfObjectName) *ContentCreator {
	_cee := ContentStreamOperation{}
	_cee.Operand = "\u0067\u0073"
	_cee.Params = _fcdg([]_cg.PdfObjectName{dictName})
	_gcf._bf = append(_gcf._bf, &_cee)
	return _gcf
}

// Add_l adds 'l' operand to the content stream:
// Append a straight line segment from the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_dbb *ContentCreator) Add_l(x, y float64) *ContentCreator {
	_add := ContentStreamOperation{}
	_add.Operand = "\u006c"
	_add.Params = _gbg([]float64{x, y})
	_dbb._bf = append(_dbb._bf, &_add)
	return _dbb
}

// String returns `ops.Bytes()` as a string.
func (_adb *ContentStreamOperations) String() string { return string(_adb.Bytes()) }

// Add_j adds 'j' operand to the content stream: Set the line join style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cb *ContentCreator) Add_j(lineJoinStyle string) *ContentCreator {
	_fcc := ContentStreamOperation{}
	_fcc.Operand = "\u006a"
	_fcc.Params = _fcdg([]_cg.PdfObjectName{_cg.PdfObjectName(lineJoinStyle)})
	_cb._bf = append(_cb._bf, &_fcc)
	return _cb
}

// Wrap ensures that the contentstream is wrapped within a balanced q ... Q expression.
func (_aff *ContentCreator) Wrap() { _aff._bf.WrapIfNeeded() }
func (_gggf *ContentStreamParser) parseOperand() (*_cg.PdfObjectString, error) {
	var _gaa []byte
	for {
		_gaaf, _bgca := _gggf._gfea.Peek(1)
		if _bgca != nil {
			return _cg.MakeString(string(_gaa)), _bgca
		}
		if _cg.IsDelimiter(_gaaf[0]) {
			break
		}
		if _cg.IsWhiteSpace(_gaaf[0]) {
			break
		}
		_ggbaa, _ := _gggf._gfea.ReadByte()
		_gaa = append(_gaa, _ggbaa)
	}
	return _cg.MakeString(string(_gaa)), nil
}
func _fec(_eff *ContentStreamInlineImage, _bbee *_cg.PdfObjectDictionary) (*_cg.LZWEncoder, error) {
	_gdg := _cg.NewLZWEncoder()
	if _bbee == nil {
		if _eff.DecodeParms != nil {
			_aca, _fca := _cg.GetDict(_eff.DecodeParms)
			if !_fca {
				_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _eff.DecodeParms)
				return nil, _e.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_bbee = _aca
		}
	}
	if _bbee == nil {
		return _gdg, nil
	}
	_ecba := _bbee.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _ecba != nil {
		_eafg, _fge := _ecba.(*_cg.PdfObjectInteger)
		if !_fge {
			_ec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _ecba)
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_eafg != 0 && *_eafg != 1 {
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_gdg.EarlyChange = int(*_eafg)
	} else {
		_gdg.EarlyChange = 1
	}
	_ecba = _bbee.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _ecba != nil {
		_ced, _ggfc := _ecba.(*_cg.PdfObjectInteger)
		if !_ggfc {
			_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _ecba)
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_gdg.Predictor = int(*_ced)
	}
	_ecba = _bbee.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ecba != nil {
		_dac, _gcc := _ecba.(*_cg.PdfObjectInteger)
		if !_gcc {
			_ec.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_gdg.BitsPerComponent = int(*_dac)
	}
	if _gdg.Predictor > 1 {
		_gdg.Columns = 1
		_ecba = _bbee.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _ecba != nil {
			_eece, _abg := _ecba.(*_cg.PdfObjectInteger)
			if !_abg {
				return nil, _e.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_gdg.Columns = int(*_eece)
		}
		_gdg.Colors = 1
		_ecba = _bbee.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _ecba != nil {
			_egb, _gcdg := _ecba.(*_cg.PdfObjectInteger)
			if !_gcdg {
				return nil, _e.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_gdg.Colors = int(*_egb)
		}
	}
	_ec.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _bbee.String())
	return _gdg, nil
}

// NewContentCreator returns a new initialized ContentCreator.
func NewContentCreator() *ContentCreator {
	_ead := &ContentCreator{}
	_ead._bf = ContentStreamOperations{}
	return _ead
}

// Add_q adds 'q' operand to the content stream: Pushes the current graphics state on the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_fb *ContentCreator) Add_q() *ContentCreator {
	_ff := ContentStreamOperation{}
	_ff.Operand = "\u0071"
	_fb._bf = append(_fb._bf, &_ff)
	return _fb
}
func (_dgdcd *ContentStreamProcessor) handleCommand_scn(_fdccc *ContentStreamOperation, _cdea *_dad.PdfPageResources) error {
	_agfd := _dgdcd._ebdd.ColorspaceNonStroking
	if !_edba(_agfd) {
		if len(_fdccc.Params) != _agfd.GetNumComponents() {
			_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fdccc.Params), _agfd)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_agfg, _fbcc := _agfd.ColorFromPdfObjects(_fdccc.Params)
	if _fbcc != nil {
		_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0063o\u006co\u0072\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0043\u0053\u0020\u0069\u0073\u0020\u0025\u002b\u0076\u0029", _fdccc.Params, _agfd)
		return _fbcc
	}
	_dgdcd._ebdd.ColorNonStroking = _agfg
	return nil
}

var _deea = _d.MustCompile("\u005e\u002f\u007b\u0032\u002c\u007d")

// HandlerConditionEnum represents the type of operand content stream processor (handler).
// The handler may process a single specific named operand or all operands.
type HandlerConditionEnum int

func _gba(_dgf *ContentStreamInlineImage) (*_cg.MultiEncoder, error) {
	_agfa := _cg.NewMultiEncoder()
	var _dfdf *_cg.PdfObjectDictionary
	var _acf []_cg.PdfObject
	if _afb := _dgf.DecodeParms; _afb != nil {
		_gfg, _fadd := _afb.(*_cg.PdfObjectDictionary)
		if _fadd {
			_dfdf = _gfg
		}
		_aabe, _eeg := _afb.(*_cg.PdfObjectArray)
		if _eeg {
			for _, _fef := range _aabe.Elements() {
				if _bfe, _cedf := _fef.(*_cg.PdfObjectDictionary); _cedf {
					_acf = append(_acf, _bfe)
				} else {
					_acf = append(_acf, nil)
				}
			}
		}
	}
	_ebd := _dgf.Filter
	if _ebd == nil {
		return nil, _e.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_ebf, _ebfb := _ebd.(*_cg.PdfObjectArray)
	if !_ebfb {
		return nil, _e.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _bbed, _ebfg := range _ebf.Elements() {
		_abe, _agb := _ebfg.(*_cg.PdfObjectName)
		if !_agb {
			return nil, _e.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _efcf _cg.PdfObject
		if _dfdf != nil {
			_efcf = _dfdf
		} else {
			if len(_acf) > 0 {
				if _bbed >= len(_acf) {
					return nil, _e.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_efcf = _acf[_bbed]
			}
		}
		var _dcb *_cg.PdfObjectDictionary
		if _ffg, _gegf := _efcf.(*_cg.PdfObjectDictionary); _gegf {
			_dcb = _ffg
		}
		if *_abe == _cg.StreamEncodingFilterNameFlate || *_abe == "\u0046\u006c" {
			_fda, _bdg := _ecde(_dgf, _dcb)
			if _bdg != nil {
				return nil, _bdg
			}
			_agfa.AddEncoder(_fda)
		} else if *_abe == _cg.StreamEncodingFilterNameLZW {
			_ecbb, _fcaa := _fec(_dgf, _dcb)
			if _fcaa != nil {
				return nil, _fcaa
			}
			_agfa.AddEncoder(_ecbb)
		} else if *_abe == _cg.StreamEncodingFilterNameASCIIHex {
			_dbaf := _cg.NewASCIIHexEncoder()
			_agfa.AddEncoder(_dbaf)
		} else if *_abe == _cg.StreamEncodingFilterNameASCII85 || *_abe == "\u0041\u0038\u0035" {
			_bbg := _cg.NewASCII85Encoder()
			_agfa.AddEncoder(_bbg)
		} else {
			_ec.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_abe)
			return nil, _e.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _agfa, nil
}

// All returns true if `hce` is equivalent to HandlerConditionEnumAllOperands.
func (_agc HandlerConditionEnum) All() bool { return _agc == HandlerConditionEnumAllOperands }

// Process processes the entire list of operations. Maintains the graphics state that is passed to any
// handlers that are triggered during processing (either on specific operators or all).
func (_fdfd *ContentStreamProcessor) Process(resources *_dad.PdfPageResources) error {
	_fdfd._ebdd.ColorspaceStroking = _dad.NewPdfColorspaceDeviceGray()
	_fdfd._ebdd.ColorspaceNonStroking = _dad.NewPdfColorspaceDeviceGray()
	_fdfd._ebdd.ColorStroking = _dad.NewPdfColorDeviceGray(0)
	_fdfd._ebdd.ColorNonStroking = _dad.NewPdfColorDeviceGray(0)
	_fdfd._ebdd.CTM = _da.IdentityMatrix()
	for _, _gcg := range _fdfd._aedb {
		var _effa error
		switch _gcg.Operand {
		case "\u0071":
			_fdfd._bfa.Push(_fdfd._ebdd)
		case "\u0051":
			if len(_fdfd._bfa) == 0 {
				_ec.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0060\u0051\u0060\u0020\u006f\u0070e\u0072\u0061\u0074\u006f\u0072\u002e\u0020\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074\u0061\u0074\u0065 \u0073\u0074\u0061\u0063\u006b\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079.\u0020\u0053\u006bi\u0070\u0070\u0069\u006e\u0067\u002e")
				continue
			}
			_fdfd._ebdd = _fdfd._bfa.Pop()
		case "\u0043\u0053":
			_effa = _fdfd.handleCommand_CS(_gcg, resources)
		case "\u0063\u0073":
			_effa = _fdfd.handleCommand_cs(_gcg, resources)
		case "\u0053\u0043":
			_effa = _fdfd.handleCommand_SC(_gcg, resources)
		case "\u0053\u0043\u004e":
			_effa = _fdfd.handleCommand_SCN(_gcg, resources)
		case "\u0073\u0063":
			_effa = _fdfd.handleCommand_sc(_gcg, resources)
		case "\u0073\u0063\u006e":
			_effa = _fdfd.handleCommand_scn(_gcg, resources)
		case "\u0047":
			_effa = _fdfd.handleCommand_G(_gcg, resources)
		case "\u0067":
			_effa = _fdfd.handleCommand_g(_gcg, resources)
		case "\u0052\u0047":
			_effa = _fdfd.handleCommand_RG(_gcg, resources)
		case "\u0072\u0067":
			_effa = _fdfd.handleCommand_rg(_gcg, resources)
		case "\u004b":
			_effa = _fdfd.handleCommand_K(_gcg, resources)
		case "\u006b":
			_effa = _fdfd.handleCommand_k(_gcg, resources)
		case "\u0063\u006d":
			_effa = _fdfd.handleCommand_cm(_gcg, resources)
		}
		if _effa != nil {
			_ec.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073s\u006f\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u0028\u0025\u0073)\u003a\u0020\u0025\u0076", _gcg.Operand, _effa)
			_ec.Log.Debug("\u004f\u0070\u0065r\u0061\u006e\u0064\u003a\u0020\u0025\u0023\u0076", _gcg.Operand)
			return _effa
		}
		for _, _gbaa := range _fdfd._gdgg {
			var _adfg error
			if _gbaa.Condition.All() {
				_adfg = _gbaa.Handler(_gcg, _fdfd._ebdd, resources)
			} else if _gbaa.Condition.Operand() && _gcg.Operand == _gbaa.Operand {
				_adfg = _gbaa.Handler(_gcg, _fdfd._ebdd, resources)
			}
			if _adfg != nil {
				_ec.Log.Debug("P\u0072\u006f\u0063\u0065\u0073\u0073o\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0072 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _adfg)
				return _adfg
			}
		}
	}
	return nil
}
func (_cdf *ContentStreamParser) parseName() (_cg.PdfObjectName, error) {
	_facc := ""
	_fedc := false
	for {
		_ggba, _fgc := _cdf._gfea.Peek(1)
		if _fgc == _ce.EOF {
			break
		}
		if _fgc != nil {
			return _cg.PdfObjectName(_facc), _fgc
		}
		if !_fedc {
			if _ggba[0] == '/' {
				_fedc = true
				_cdf._gfea.ReadByte()
			} else {
				_ec.Log.Error("N\u0061\u006d\u0065\u0020\u0073\u0074a\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069\u0074h\u0020\u0025\u0073 \u0028%\u0020\u0078\u0029", _ggba, _ggba)
				return _cg.PdfObjectName(_facc), _e.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _ggba[0])
			}
		} else {
			if _cg.IsWhiteSpace(_ggba[0]) {
				break
			} else if (_ggba[0] == '/') || (_ggba[0] == '[') || (_ggba[0] == '(') || (_ggba[0] == ']') || (_ggba[0] == '<') || (_ggba[0] == '>') {
				break
			} else if _ggba[0] == '#' {
				_deec, _dagb := _cdf._gfea.Peek(3)
				if _dagb != nil {
					return _cg.PdfObjectName(_facc), _dagb
				}
				_cdf._gfea.Discard(3)
				_gefg, _dagb := _g.DecodeString(string(_deec[1:3]))
				if _dagb != nil {
					return _cg.PdfObjectName(_facc), _dagb
				}
				_facc += string(_gefg)
			} else {
				_dgec, _ := _cdf._gfea.ReadByte()
				_facc += string(_dgec)
			}
		}
	}
	return _cg.PdfObjectName(_facc), nil
}

// Add_c adds 'c' operand to the content stream: Append a Bezier curve to the current path from
// the current point to (x3,y3) with (x1,x1) and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_bdb *ContentCreator) Add_c(x1, y1, x2, y2, x3, y3 float64) *ContentCreator {
	_faa := ContentStreamOperation{}
	_faa.Operand = "\u0063"
	_faa.Params = _gbg([]float64{x1, y1, x2, y2, x3, y3})
	_bdb._bf = append(_bdb._bf, &_faa)
	return _bdb
}
func (_dgdcg *ContentStreamParser) parseDict() (*_cg.PdfObjectDictionary, error) {
	_ec.Log.Trace("\u0052\u0065\u0061\u0064i\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074 \u0073t\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0021")
	_gce := _cg.MakeDict()
	_acee, _ := _dgdcg._gfea.ReadByte()
	if _acee != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_acee, _ = _dgdcg._gfea.ReadByte()
	if _acee != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_dgdcg.skipSpaces()
		_fgae, _agba := _dgdcg._gfea.Peek(2)
		if _agba != nil {
			return nil, _agba
		}
		_ec.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_fgae), string(_fgae))
		if (_fgae[0] == '>') && (_fgae[1] == '>') {
			_ec.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_dgdcg._gfea.ReadByte()
			_dgdcg._gfea.ReadByte()
			break
		}
		_ec.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_agab, _agba := _dgdcg.parseName()
		_ec.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _agab)
		if _agba != nil {
			_ec.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _agba)
			return nil, _agba
		}
		if len(_agab) > 4 && _agab[len(_agab)-4:] == "\u006e\u0075\u006c\u006c" {
			_eace := _agab[0 : len(_agab)-4]
			_ec.Log.Trace("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _agab)
			_ec.Log.Trace("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _eace)
			_dgdcg.skipSpaces()
			_dcd, _ := _dgdcg._gfea.Peek(1)
			if _dcd[0] == '/' {
				_gce.Set(_eace, _cg.MakeNull())
				continue
			}
		}
		_dgdcg.skipSpaces()
		_ffea, _, _agba := _dgdcg.parseObject()
		if _agba != nil {
			return nil, _agba
		}
		_gce.Set(_agab, _ffea)
		_ec.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _agab, _ffea.String())
	}
	return _gce, nil
}
func (_fede *ContentStreamProcessor) handleCommand_K(_cgbe *ContentStreamOperation, _agge *_dad.PdfPageResources) error {
	_cfef := _dad.NewPdfColorspaceDeviceCMYK()
	if len(_cgbe.Params) != _cfef.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cgbe.Params), _cfef)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_dce, _fddb := _cfef.ColorFromPdfObjects(_cgbe.Params)
	if _fddb != nil {
		return _fddb
	}
	_fede._ebdd.ColorspaceStroking = _cfef
	_fede._ebdd.ColorStroking = _dce
	return nil
}
func _gccb(_gbbb _cg.PdfObject) (_dad.PdfColorspace, error) {
	_gdd, _fffa := _gbbb.(*_cg.PdfObjectArray)
	if !_fffa {
		_ec.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020cs\u0020\u006e\u006ft\u0020\u0069n\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025#\u0076\u0029", _gbbb)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gdd.Len() != 4 {
		_ec.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061r\u0072\u0061\u0079\u002c\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0034\u0020\u0028\u0025\u0064\u0029", _gdd.Len())
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_bebac, _fffa := _gdd.Get(0).(*_cg.PdfObjectName)
	if !_fffa {
		_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072s\u0074 \u0065\u006c\u0065\u006de\u006e\u0074 \u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0023\u0076\u0029", *_gdd)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_bebac != "\u0049" && *_bebac != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		_ec.Log.Debug("\u0045\u0072r\u006f\u0072\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0049\u0020\u0028\u0067\u006f\u0074\u003a\u0020\u0025\u0076\u0029", *_bebac)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_bebac, _fffa = _gdd.Get(1).(*_cg.PdfObjectName)
	if !_fffa {
		_ec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072a\u0079\u003a\u0020\u0025\u0023v\u0029", *_gdd)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_bebac != "\u0047" && *_bebac != "\u0052\u0047\u0042" && *_bebac != "\u0043\u004d\u0059\u004b" && *_bebac != "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" && *_bebac != "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" && *_bebac != "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		_ec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0047\u002f\u0052\u0047\u0042\u002f\u0043\u004d\u0059\u004b\u0020\u0028g\u006f\u0074\u003a\u0020\u0025v\u0029", *_bebac)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_gbfde := ""
	switch *_bebac {
	case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		_gbfde = "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
	case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		_gbfde = "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
	case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		_gbfde = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	_aebe := _cg.MakeArray(_cg.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"), _cg.MakeName(_gbfde), _gdd.Get(2), _gdd.Get(3))
	return _dad.NewPdfColorspaceFromPdfObject(_aebe)
}

// WriteString outputs the object as it is to be written to file.
func (_egdg *ContentStreamInlineImage) WriteString() string {
	var _beb _a.Buffer
	_ceg := ""
	if _egdg.BitsPerComponent != nil {
		_ceg += "\u002f\u0042\u0050C\u0020" + _egdg.BitsPerComponent.WriteString() + "\u000a"
	}
	if _egdg.ColorSpace != nil {
		_ceg += "\u002f\u0043\u0053\u0020" + _egdg.ColorSpace.WriteString() + "\u000a"
	}
	if _egdg.Decode != nil {
		_ceg += "\u002f\u0044\u0020" + _egdg.Decode.WriteString() + "\u000a"
	}
	if _egdg.DecodeParms != nil {
		_ceg += "\u002f\u0044\u0050\u0020" + _egdg.DecodeParms.WriteString() + "\u000a"
	}
	if _egdg.Filter != nil {
		_ceg += "\u002f\u0046\u0020" + _egdg.Filter.WriteString() + "\u000a"
	}
	if _egdg.Height != nil {
		_ceg += "\u002f\u0048\u0020" + _egdg.Height.WriteString() + "\u000a"
	}
	if _egdg.ImageMask != nil {
		_ceg += "\u002f\u0049\u004d\u0020" + _egdg.ImageMask.WriteString() + "\u000a"
	}
	if _egdg.Intent != nil {
		_ceg += "\u002f\u0049\u006e\u0074\u0065\u006e\u0074\u0020" + _egdg.Intent.WriteString() + "\u000a"
	}
	if _egdg.Interpolate != nil {
		_ceg += "\u002f\u0049\u0020" + _egdg.Interpolate.WriteString() + "\u000a"
	}
	if _egdg.Width != nil {
		_ceg += "\u002f\u0057\u0020" + _egdg.Width.WriteString() + "\u000a"
	}
	_beb.WriteString(_ceg)
	_beb.WriteString("\u0049\u0044\u0020")
	_beb.Write(_egdg._cbdd)
	_beb.WriteString("\u000a\u0045\u0049\u000a")
	return _beb.String()
}
func (_af *ContentStreamOperations) isWrapped() bool {
	if len(*_af) < 2 {
		return false
	}
	_dada := 0
	for _, _ba := range *_af {
		if _ba.Operand == "\u0071" {
			_dada++
		} else if _ba.Operand == "\u0051" {
			_dada--
		} else {
			if _dada < 1 {
				return false
			}
		}
	}
	return _dada == 0
}
func (_dbcb *ContentStreamProcessor) handleCommand_k(_fbde *ContentStreamOperation, _acda *_dad.PdfPageResources) error {
	_ecgg := _dad.NewPdfColorspaceDeviceCMYK()
	if len(_fbde.Params) != _ecgg.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fbde.Params), _ecgg)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_eef, _afff := _ecgg.ColorFromPdfObjects(_fbde.Params)
	if _afff != nil {
		return _afff
	}
	_dbcb._ebdd.ColorspaceNonStroking = _ecgg
	_dbcb._ebdd.ColorNonStroking = _eef
	return nil
}

// Add_Tstar appends 'T*' operand to the content stream:
// Move to the start of next line.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_bbc *ContentCreator) Add_Tstar() *ContentCreator {
	_ebg := ContentStreamOperation{}
	_ebg.Operand = "\u0054\u002a"
	_bbc._bf = append(_bbc._bf, &_ebg)
	return _bbc
}

// Add_n appends 'n' operand to the content stream:
// End the path without filling or stroking.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_dff *ContentCreator) Add_n() *ContentCreator {
	_geg := ContentStreamOperation{}
	_geg.Operand = "\u006e"
	_dff._bf = append(_dff._bf, &_geg)
	return _dff
}

// Add_Tw appends 'Tw' operand to the content stream:
// Set word spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ade *ContentCreator) Add_Tw(wordSpace float64) *ContentCreator {
	_aba := ContentStreamOperation{}
	_aba.Operand = "\u0054\u0077"
	_aba.Params = _gbg([]float64{wordSpace})
	_ade._bf = append(_ade._bf, &_aba)
	return _ade
}

// Translate applies a simple x-y translation to the transformation matrix.
func (_fdc *ContentCreator) Translate(tx, ty float64) *ContentCreator {
	return _fdc.Add_cm(1, 0, 0, 1, tx, ty)
}
func (_gfbb *ContentStreamProcessor) handleCommand_RG(_gfca *ContentStreamOperation, _gece *_dad.PdfPageResources) error {
	_gccee := _dad.NewPdfColorspaceDeviceRGB()
	if len(_gfca.Params) != _gccee.GetNumComponents() {
		_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020R\u0047")
		_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_gfca.Params), _gccee)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_cfbd, _acaa := _gccee.ColorFromPdfObjects(_gfca.Params)
	if _acaa != nil {
		return _acaa
	}
	_gfbb._ebdd.ColorspaceStroking = _gccee
	_gfbb._ebdd.ColorStroking = _cfbd
	return nil
}
func (_afeg *ContentStreamProcessor) getColorspace(_fcef string, _ddb *_dad.PdfPageResources) (_dad.PdfColorspace, error) {
	switch _fcef {
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		return _dad.NewPdfColorspaceDeviceGray(), nil
	case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		return _dad.NewPdfColorspaceDeviceRGB(), nil
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		return _dad.NewPdfColorspaceDeviceCMYK(), nil
	case "\u0050a\u0074\u0074\u0065\u0072\u006e":
		return _dad.NewPdfColorspaceSpecialPattern(), nil
	}
	if _ddb != nil {
		_gfcb, _bafg := _ddb.GetColorspaceByName(_cg.PdfObjectName(_fcef))
		if _bafg {
			return _gfcb, nil
		}
	}
	switch _fcef {
	case "\u0043a\u006c\u0047\u0072\u0061\u0079":
		return _dad.NewPdfColorspaceCalGray(), nil
	case "\u0043\u0061\u006c\u0052\u0047\u0042":
		return _dad.NewPdfColorspaceCalRGB(), nil
	case "\u004c\u0061\u0062":
		return _dad.NewPdfColorspaceLab(), nil
	}
	_ec.Log.Debug("\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063e\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u0065\u0064\u003a\u0020\u0025\u0073", _fcef)
	return nil, _e.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065:\u0020\u0025\u0073", _fcef)
}

// Add_y appends 'y' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with (x1, y1) and (x3,y3) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cfe *ContentCreator) Add_y(x1, y1, x3, y3 float64) *ContentCreator {
	_ga := ContentStreamOperation{}
	_ga.Operand = "\u0079"
	_ga.Params = _gbg([]float64{x1, y1, x3, y3})
	_cfe._bf = append(_cfe._bf, &_ga)
	return _cfe
}

// RotateDeg applies a rotation to the transformation matrix.
func (_eaf *ContentCreator) RotateDeg(angle float64) *ContentCreator {
	_cd := _de.Cos(angle * _de.Pi / 180.0)
	_deg := _de.Sin(angle * _de.Pi / 180.0)
	_bda := -_de.Sin(angle * _de.Pi / 180.0)
	_ac := _de.Cos(angle * _de.Pi / 180.0)
	return _eaf.Add_cm(_cd, _deg, _bda, _ac, 0, 0)
}

// Add_i adds 'i' operand to the content stream: Set the flatness tolerance in the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ffa *ContentCreator) Add_i(flatness float64) *ContentCreator {
	_fddc := ContentStreamOperation{}
	_fddc.Operand = "\u0069"
	_fddc.Params = _gbg([]float64{flatness})
	_ffa._bf = append(_ffa._bf, &_fddc)
	return _ffa
}

// Operand returns true if `hce` is equivalent to HandlerConditionEnumOperand.
func (_efgd HandlerConditionEnum) Operand() bool { return _efgd == HandlerConditionEnumOperand }
func (_dfg *ContentStreamParser) parseArray() (*_cg.PdfObjectArray, error) {
	_ded := _cg.MakeArray()
	_dfg._gfea.ReadByte()
	for {
		_dfg.skipSpaces()
		_cfb, _dfgd := _dfg._gfea.Peek(1)
		if _dfgd != nil {
			return _ded, _dfgd
		}
		if _cfb[0] == ']' {
			_dfg._gfea.ReadByte()
			break
		}
		_cacd, _, _dfgd := _dfg.parseObject()
		if _dfgd != nil {
			return _ded, _dfgd
		}
		_ded.Append(_cacd)
	}
	return _ded, nil
}

// Add_SCN_pattern appends 'SCN' operand to the content stream for pattern `name`:
// SCN with name attribute (for pattern). Syntax: c1 ... cn name SCN.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_aec *ContentCreator) Add_SCN_pattern(name _cg.PdfObjectName, c ...float64) *ContentCreator {
	_bba := ContentStreamOperation{}
	_bba.Operand = "\u0053\u0043\u004e"
	_bba.Params = _gbg(c)
	_bba.Params = append(_bba.Params, _cg.MakeName(string(name)))
	_aec._bf = append(_aec._bf, &_bba)
	return _aec
}

// GetColorSpace returns the colorspace of the inline image.
func (_fbec *ContentStreamInlineImage) GetColorSpace(resources *_dad.PdfPageResources) (_dad.PdfColorspace, error) {
	if _fbec.ColorSpace == nil {
		_ec.Log.Debug("\u0049\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076i\u006e\u0067\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u002c\u0020\u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u0047\u0072a\u0079")
		return _dad.NewPdfColorspaceDeviceGray(), nil
	}
	if _bbeef, _gbd := _fbec.ColorSpace.(*_cg.PdfObjectArray); _gbd {
		return _gccb(_bbeef)
	}
	_bbbg, _agdf := _fbec.ColorSpace.(*_cg.PdfObjectName)
	if !_agdf {
		_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u003b\u0025\u002bv\u0029", _fbec.ColorSpace, _fbec.ColorSpace)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_bbbg == "\u0047" || *_bbbg == "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" {
		return _dad.NewPdfColorspaceDeviceGray(), nil
	} else if *_bbbg == "\u0052\u0047\u0042" || *_bbbg == "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" {
		return _dad.NewPdfColorspaceDeviceRGB(), nil
	} else if *_bbbg == "\u0043\u004d\u0059\u004b" || *_bbbg == "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		return _dad.NewPdfColorspaceDeviceCMYK(), nil
	} else if *_bbbg == "\u0049" || *_bbbg == "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _b.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0049\u006e\u0064e\u0078 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
	} else {
		if resources.ColorSpace == nil {
			_ec.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_bbbg)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		_aac, _gcb := resources.GetColorspaceByName(*_bbbg)
		if !_gcb {
			_ec.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_bbbg)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		return _aac, nil
	}
}

// Add_G appends 'G' operand to the content stream:
// Set the stroking colorspace to DeviceGray and sets the gray level (0-1).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_faag *ContentCreator) Add_G(gray float64) *ContentCreator {
	_ecc := ContentStreamOperation{}
	_ecc.Operand = "\u0047"
	_ecc.Params = _gbg([]float64{gray})
	_faag._bf = append(_faag._bf, &_ecc)
	return _faag
}

// Add_quotes appends `"` operand to the content stream:
// Move to next line and show a string, using `aw` and `ac` as word
// and character spacing respectively.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_ggcd *ContentCreator) Add_quotes(textstr _cg.PdfObjectString, aw, ac float64) *ContentCreator {
	_ebb := ContentStreamOperation{}
	_ebb.Operand = "\u0022"
	_ebb.Params = _gbg([]float64{aw, ac})
	_ebb.Params = append(_ebb.Params, _eab([]_cg.PdfObjectString{textstr})...)
	_ggcd._bf = append(_ggcd._bf, &_ebb)
	return _ggcd
}

// ContentStreamInlineImage is a representation of an inline image in a Content stream. Everything between the BI and EI operands.
// ContentStreamInlineImage implements the core.PdfObject interface although strictly it is not a PDF object.
type ContentStreamInlineImage struct {
	BitsPerComponent _cg.PdfObject
	ColorSpace       _cg.PdfObject
	Decode           _cg.PdfObject
	DecodeParms      _cg.PdfObject
	Filter           _cg.PdfObject
	Height           _cg.PdfObject
	ImageMask        _cg.PdfObject
	Intent           _cg.PdfObject
	Interpolate      _cg.PdfObject
	Width            _cg.PdfObject
	_cbdd            []byte
	_ebfd            *_aa.ImageBase
}

// Scale applies x-y scaling to the transformation matrix.
func (_efc *ContentCreator) Scale(sx, sy float64) *ContentCreator {
	return _efc.Add_cm(sx, 0, 0, sy, 0, 0)
}

// Add_cs appends 'cs' operand to the content stream:
// Same as CS but for non-stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cfa *ContentCreator) Add_cs(name _cg.PdfObjectName) *ContentCreator {
	_fgf := ContentStreamOperation{}
	_fgf.Operand = "\u0063\u0073"
	_fgf.Params = _fcdg([]_cg.PdfObjectName{name})
	_cfa._bf = append(_cfa._bf, &_fgf)
	return _cfa
}

// Add_ri adds 'ri' operand to the content stream, which sets the color rendering intent.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_adbf *ContentCreator) Add_ri(intent _cg.PdfObjectName) *ContentCreator {
	_fdd := ContentStreamOperation{}
	_fdd.Operand = "\u0072\u0069"
	_fdd.Params = _fcdg([]_cg.PdfObjectName{intent})
	_adbf._bf = append(_adbf._bf, &_fdd)
	return _adbf
}
func (_fedg *ContentStreamProcessor) handleCommand_SCN(_fdad *ContentStreamOperation, _eaec *_dad.PdfPageResources) error {
	_beg := _fedg._ebdd.ColorspaceStroking
	if !_edba(_beg) {
		if len(_fdad.Params) != _beg.GetNumComponents() {
			_ec.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_ec.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fdad.Params), _beg)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_efd, _eecd := _beg.ColorFromPdfObjects(_fdad.Params)
	if _eecd != nil {
		return _eecd
	}
	_fedg._ebdd.ColorStroking = _efd
	return nil
}

// Add_TL appends 'TL' operand to the content stream:
// Set leading.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_dga *ContentCreator) Add_TL(leading float64) *ContentCreator {
	_cgd := ContentStreamOperation{}
	_cgd.Operand = "\u0054\u004c"
	_cgd.Params = _gbg([]float64{leading})
	_dga._bf = append(_dga._bf, &_cgd)
	return _dga
}

// Add_W_starred appends 'W*' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (even odd rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_ecd *ContentCreator) Add_W_starred() *ContentCreator {
	_dda := ContentStreamOperation{}
	_dda.Operand = "\u0057\u002a"
	_ecd._bf = append(_ecd._bf, &_dda)
	return _ecd
}

// Add_Td appends 'Td' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_deag *ContentCreator) Add_Td(tx, ty float64) *ContentCreator {
	_bc := ContentStreamOperation{}
	_bc.Operand = "\u0054\u0064"
	_bc.Params = _gbg([]float64{tx, ty})
	_deag._bf = append(_deag._bf, &_bc)
	return _deag
}
func (_bdf *ContentStreamParser) parseHexString() (*_cg.PdfObjectString, error) {
	_bdf._gfea.ReadByte()
	_eegb := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	var _abb []byte
	for {
		_bdf.skipSpaces()
		_dcbf, _cecf := _bdf._gfea.Peek(1)
		if _cecf != nil {
			return _cg.MakeString(""), _cecf
		}
		if _dcbf[0] == '>' {
			_bdf._gfea.ReadByte()
			break
		}
		_ccc, _ := _bdf._gfea.ReadByte()
		if _a.IndexByte(_eegb, _ccc) >= 0 {
			_abb = append(_abb, _ccc)
		}
	}
	if len(_abb)%2 == 1 {
		_abb = append(_abb, '0')
	}
	_cgcb, _ := _g.DecodeString(string(_abb))
	return _cg.MakeHexString(string(_cgcb)), nil
}

// Add_RG appends 'RG' operand to the content stream:
// Set the stroking colorspace to DeviceRGB and sets the r,g,b colors (0-1 each).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_eaa *ContentCreator) Add_RG(r, g, b float64) *ContentCreator {
	_ggc := ContentStreamOperation{}
	_ggc.Operand = "\u0052\u0047"
	_ggc.Params = _gbg([]float64{r, g, b})
	_eaa._bf = append(_eaa._bf, &_ggc)
	return _eaa
}

// ContentStreamOperation represents an operation in PDF contentstream which consists of
// an operand and parameters.
type ContentStreamOperation struct {
	Params  []_cg.PdfObject
	Operand string
}

// ParseInlineImage parses an inline image from a content stream, both reading its properties and binary data.
// When called, "BI" has already been read from the stream.  This function
// finishes reading through "EI" and then returns the ContentStreamInlineImage.
func (_afde *ContentStreamParser) ParseInlineImage() (*ContentStreamInlineImage, error) {
	_ccge := ContentStreamInlineImage{}
	for {
		_afde.skipSpaces()
		_degc, _fdec, _bbbd := _afde.parseObject()
		if _bbbd != nil {
			return nil, _bbbd
		}
		if !_fdec {
			_bce, _baf := _cg.GetName(_degc)
			if !_baf {
				_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _degc)
				return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _degc)
			}
			_ddgd, _gge, _fece := _afde.parseObject()
			if _fece != nil {
				return nil, _fece
			}
			if _gge {
				return nil, _e.Errorf("\u006eo\u0074\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067 \u0061\u006e\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			switch *_bce {
			case "\u0042\u0050\u0043", "\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074":
				_ccge.BitsPerComponent = _ddgd
			case "\u0043\u0053", "\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065":
				_ccge.ColorSpace = _ddgd
			case "\u0044", "\u0044\u0065\u0063\u006f\u0064\u0065":
				_ccge.Decode = _ddgd
			case "\u0044\u0050", "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073":
				_ccge.DecodeParms = _ddgd
			case "\u0046", "\u0046\u0069\u006c\u0074\u0065\u0072":
				_ccge.Filter = _ddgd
			case "\u0048", "\u0048\u0065\u0069\u0067\u0068\u0074":
				_ccge.Height = _ddgd
			case "\u0049\u004d", "\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k":
				_ccge.ImageMask = _ddgd
			case "\u0049\u006e\u0074\u0065\u006e\u0074":
				_ccge.Intent = _ddgd
			case "\u0049", "I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065":
				_ccge.Interpolate = _ddgd
			case "\u0057", "\u0057\u0069\u0064t\u0068":
				_ccge.Width = _ddgd
			case "\u004c\u0065\u006e\u0067\u0074\u0068", "\u0053u\u0062\u0074\u0079\u0070\u0065", "\u0054\u0079\u0070\u0065":
				_ec.Log.Debug("\u0049\u0067\u006e\u006fr\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0070a\u0072\u0061\u006d\u0065\u0074\u0065\u0072 \u0025\u0073", *_bce)
			default:
				return nil, _e.Errorf("\u0075\u006e\u006b\u006e\u006f\u0077n\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0020\u0025\u0073", *_bce)
			}
		}
		if _fdec {
			_dge, _geff := _degc.(*_cg.PdfObjectString)
			if !_geff {
				return nil, _e.Errorf("\u0066a\u0069\u006ce\u0064\u0020\u0074o\u0020\u0072\u0065\u0061\u0064\u0020\u0069n\u006c\u0069\u006e\u0065\u0020\u0069m\u0061\u0067\u0065\u0020\u002d\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			if _dge.Str() == "\u0045\u0049" {
				_ec.Log.Trace("\u0049n\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020f\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e\u002e\u002e")
				return &_ccge, nil
			} else if _dge.Str() == "\u0049\u0044" {
				_ec.Log.Trace("\u0049\u0044\u0020\u0073\u0074\u0061\u0072\u0074")
				_gff, _beba := _afde._gfea.Peek(1)
				if _beba != nil {
					return nil, _beba
				}
				if _cg.IsWhiteSpace(_gff[0]) {
					_afde._gfea.Discard(1)
				}
				_ccge._cbdd = []byte{}
				_ed := 0
				var _eadb []byte
				for {
					_ggfb, _gfe := _afde._gfea.ReadByte()
					if _gfe != nil {
						_ec.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u0066\u0069\u006e\u0064\u0020\u0065\u006e\u0064\u0020o\u0066\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0045\u0049\u0020\u0069\u006e\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0061\u0074a")
						return nil, _gfe
					}
					if _ed == 0 {
						if _cg.IsWhiteSpace(_ggfb) {
							_eadb = []byte{}
							_eadb = append(_eadb, _ggfb)
							_ed = 1
						} else if _ggfb == 'E' {
							_eadb = append(_eadb, _ggfb)
							_ed = 2
						} else {
							_ccge._cbdd = append(_ccge._cbdd, _ggfb)
						}
					} else if _ed == 1 {
						_eadb = append(_eadb, _ggfb)
						if _ggfb == 'E' {
							_ed = 2
						} else {
							_ccge._cbdd = append(_ccge._cbdd, _eadb...)
							_eadb = []byte{}
							if _cg.IsWhiteSpace(_ggfb) {
								_ed = 1
							} else {
								_ed = 0
							}
						}
					} else if _ed == 2 {
						_eadb = append(_eadb, _ggfb)
						if _ggfb == 'I' {
							_ed = 3
						} else {
							_ccge._cbdd = append(_ccge._cbdd, _eadb...)
							_eadb = []byte{}
							_ed = 0
						}
					} else if _ed == 3 {
						_eadb = append(_eadb, _ggfb)
						if _cg.IsWhiteSpace(_ggfb) {
							_gcce, _dacb := _afde._gfea.Peek(20)
							if _dacb != nil && _dacb != _ce.EOF {
								return nil, _dacb
							}
							_fdde := NewContentStreamParser(string(_gcce))
							_cgcg := true
							for _bgb := 0; _bgb < 3; _bgb++ {
								_faab, _feec, _fdcc := _fdde.parseObject()
								if _fdcc != nil {
									if _fdcc == _ce.EOF {
										break
									}
									_cgcg = false
									continue
								}
								if _feec && !_daca(_faab.String()) {
									_cgcg = false
									break
								}
							}
							if _cgcg {
								if len(_ccge._cbdd) > 100 {
									_ec.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078 \u002e\u002e\u002e", len(_ccge._cbdd), _ccge._cbdd[:100])
								} else {
									_ec.Log.Trace("\u0049\u006d\u0061\u0067e \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025 \u0078", len(_ccge._cbdd), _ccge._cbdd)
								}
								return &_ccge, nil
							}
						}
						_ccge._cbdd = append(_ccge._cbdd, _eadb...)
						_eadb = []byte{}
						_ed = 0
					}
				}
			}
		}
	}
}

// Add_SC appends 'SC' operand to the content stream:
// Set color for stroking operations.  Input: c1, ..., cn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cbb *ContentCreator) Add_SC(c ...float64) *ContentCreator {
	_adf := ContentStreamOperation{}
	_adf.Operand = "\u0053\u0043"
	_adf.Params = _gbg(c)
	_cbb._bf = append(_cbb._bf, &_adf)
	return _cbb
}

// Add_v appends 'v' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with the current point and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_bfd *ContentCreator) Add_v(x2, y2, x3, y3 float64) *ContentCreator {
	_cbe := ContentStreamOperation{}
	_cbe.Operand = "\u0076"
	_cbe.Params = _gbg([]float64{x2, y2, x3, y3})
	_bfd._bf = append(_bfd._bf, &_cbe)
	return _bfd
}

// Add_S appends 'S' operand to the content stream: Stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fad *ContentCreator) Add_S() *ContentCreator {
	_gga := ContentStreamOperation{}
	_gga.Operand = "\u0053"
	_fad._bf = append(_fad._bf, &_gga)
	return _fad
}

var _cga = map[string]struct{}{"\u0062": struct{}{}, "\u0042": struct{}{}, "\u0062\u002a": struct{}{}, "\u0042\u002a": struct{}{}, "\u0042\u0044\u0043": struct{}{}, "\u0042\u0049": struct{}{}, "\u0042\u004d\u0043": struct{}{}, "\u0042\u0054": struct{}{}, "\u0042\u0058": struct{}{}, "\u0063": struct{}{}, "\u0063\u006d": struct{}{}, "\u0043\u0053": struct{}{}, "\u0063\u0073": struct{}{}, "\u0064": struct{}{}, "\u0064\u0030": struct{}{}, "\u0064\u0031": struct{}{}, "\u0044\u006f": struct{}{}, "\u0044\u0050": struct{}{}, "\u0045\u0049": struct{}{}, "\u0045\u004d\u0043": struct{}{}, "\u0045\u0054": struct{}{}, "\u0045\u0058": struct{}{}, "\u0066": struct{}{}, "\u0046": struct{}{}, "\u0066\u002a": struct{}{}, "\u0047": struct{}{}, "\u0067": struct{}{}, "\u0067\u0073": struct{}{}, "\u0068": struct{}{}, "\u0069": struct{}{}, "\u0049\u0044": struct{}{}, "\u006a": struct{}{}, "\u004a": struct{}{}, "\u004b": struct{}{}, "\u006b": struct{}{}, "\u006c": struct{}{}, "\u006d": struct{}{}, "\u004d": struct{}{}, "\u004d\u0050": struct{}{}, "\u006e": struct{}{}, "\u0071": struct{}{}, "\u0051": struct{}{}, "\u0072\u0065": struct{}{}, "\u0052\u0047": struct{}{}, "\u0072\u0067": struct{}{}, "\u0072\u0069": struct{}{}, "\u0073": struct{}{}, "\u0053": struct{}{}, "\u0053\u0043": struct{}{}, "\u0073\u0063": struct{}{}, "\u0053\u0043\u004e": struct{}{}, "\u0073\u0063\u006e": struct{}{}, "\u0073\u0068": struct{}{}, "\u0054\u002a": struct{}{}, "\u0054\u0063": struct{}{}, "\u0054\u0064": struct{}{}, "\u0054\u0044": struct{}{}, "\u0054\u0066": struct{}{}, "\u0054\u006a": struct{}{}, "\u0054\u004a": struct{}{}, "\u0054\u004c": struct{}{}, "\u0054\u006d": struct{}{}, "\u0054\u0072": struct{}{}, "\u0054\u0073": struct{}{}, "\u0054\u0077": struct{}{}, "\u0054\u007a": struct{}{}, "\u0076": struct{}{}, "\u0077": struct{}{}, "\u0057": struct{}{}, "\u0057\u002a": struct{}{}, "\u0079": struct{}{}, "\u0027": struct{}{}, "\u0022": struct{}{}}

// Add_EMC appends 'EMC' operand to the content stream:
// Ends a marked-content sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_fadb *ContentCreator) Add_EMC() *ContentCreator {
	_eed := ContentStreamOperation{}
	_eed.Operand = "\u0045\u004d\u0043"
	_fadb._bf = append(_fadb._bf, &_eed)
	return _fadb
}
func (_bfg *ContentStreamInlineImage) toImageBase(_acec *_dad.PdfPageResources) (*_aa.ImageBase, error) {
	if _bfg._ebfd != nil {
		return _bfg._ebfd, nil
	}
	_ffd := _aa.ImageBase{}
	if _bfg.Height == nil {
		return nil, _b.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_cda, _gfga := _bfg.Height.(*_cg.PdfObjectInteger)
	if !_gfga {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_ffd.Height = int(*_cda)
	if _bfg.Width == nil {
		return nil, _b.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_fff, _gfga := _bfg.Width.(*_cg.PdfObjectInteger)
	if !_gfga {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064\u0074\u0068")
	}
	_ffd.Width = int(*_fff)
	_caac, _fae := _bfg.IsMask()
	if _fae != nil {
		return nil, _fae
	}
	if _caac {
		_ffd.BitsPerComponent = 1
		_ffd.ColorComponents = 1
	} else {
		if _bfg.BitsPerComponent == nil {
			_ec.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0042\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u0038")
			_ffd.BitsPerComponent = 8
		} else {
			_ccg, _fcf := _bfg.BitsPerComponent.(*_cg.PdfObjectInteger)
			if !_fcf {
				_ec.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0062\u0069\u0074\u0073 p\u0065\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0076al\u0075\u0065,\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _bfg.BitsPerComponent)
				return nil, _b.New("\u0042\u0050\u0043\u0020\u0054\u0079\u0070\u0065\u0020e\u0072\u0072\u006f\u0072")
			}
			_ffd.BitsPerComponent = int(*_ccg)
		}
		if _bfg.ColorSpace != nil {
			_cgf, _baa := _bfg.GetColorSpace(_acec)
			if _baa != nil {
				return nil, _baa
			}
			_ffd.ColorComponents = _cgf.GetNumComponents()
		} else {
			_ec.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075m\u0069\u006eg\u0020\u0031\u0020\u0063o\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			_ffd.ColorComponents = 1
		}
	}
	if _bfda, _cagg := _cg.GetArray(_bfg.Decode); _cagg {
		_ffd.Decode, _fae = _bfda.ToFloat64Array()
		if _fae != nil {
			return nil, _fae
		}
	}
	_bfg._ebfd = &_ffd
	return _bfg._ebfd, nil
}

// Add_g appends 'g' operand to the content stream:
// Same as G but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_afg *ContentCreator) Add_g(gray float64) *ContentCreator {
	_daa := ContentStreamOperation{}
	_daa.Operand = "\u0067"
	_daa.Params = _gbg([]float64{gray})
	_afg._bf = append(_afg._bf, &_daa)
	return _afg
}
func (_fdag *ContentStreamParser) skipSpaces() (int, error) {
	_gafa := 0
	for {
		_ffe, _cffe := _fdag._gfea.Peek(1)
		if _cffe != nil {
			return 0, _cffe
		}
		if _cg.IsWhiteSpace(_ffe[0]) {
			_fdag._gfea.ReadByte()
			_gafa++
		} else {
			break
		}
	}
	return _gafa, nil
}

// Add_Tf appends 'Tf' operand to the content stream:
// Set font and font size specified by font resource `fontName` and `fontSize`.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_feg *ContentCreator) Add_Tf(fontName _cg.PdfObjectName, fontSize float64) *ContentCreator {
	_dca := ContentStreamOperation{}
	_dca.Operand = "\u0054\u0066"
	_dca.Params = _fcdg([]_cg.PdfObjectName{fontName})
	_dca.Params = append(_dca.Params, _gbg([]float64{fontSize})...)
	_feg._bf = append(_feg._bf, &_dca)
	return _feg
}

// Bytes converts the content stream operations to a content stream byte presentation, i.e. the kind that can be
// stored as a PDF stream or string format.
func (_df *ContentCreator) Bytes() []byte { return _df._bf.Bytes() }

// String is same as Bytes() except returns as a string for convenience.
func (_bgc *ContentCreator) String() string { return string(_bgc._bf.Bytes()) }

// Add_Tz appends 'Tz' operand to the content stream:
// Set horizontal scaling.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_fdf *ContentCreator) Add_Tz(scale float64) *ContentCreator {
	_cgg := ContentStreamOperation{}
	_cgg.Operand = "\u0054\u007a"
	_cgg.Params = _gbg([]float64{scale})
	_fdf._bf = append(_fdf._bf, &_cgg)
	return _fdf
}

// NewContentStreamParser creates a new instance of the content stream parser from an input content
// stream string.
func NewContentStreamParser(contentStr string) *ContentStreamParser {
	_fbbf := ContentStreamParser{}
	contentStr = string(_deea.ReplaceAll([]byte(contentStr), []byte("\u002f")))
	_bbff := _a.NewBufferString(contentStr + "\u000a")
	_fbbf._gfea = _f.NewReader(_bbff)
	return &_fbbf
}

// Add_s appends 's' operand to the content stream: Close and stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ca *ContentCreator) Add_s() *ContentCreator {
	_bdab := ContentStreamOperation{}
	_bdab.Operand = "\u0073"
	_ca._bf = append(_ca._bf, &_bdab)
	return _ca
}
