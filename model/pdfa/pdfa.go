// Package pdfa provides abstraction to optimize and verify documents with respect to the PDF/A standards.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
package pdfa

import (
	_e "errors"
	_d "fmt"
	_g "image/color"
	_dg "math"
	_f "sort"
	_db "strings"
	_ff "time"

	_fg "github.com/adrg/sysfont"
	_gd "github.com/bamzi/pdfext/common"
	_gg "github.com/bamzi/pdfext/contentstream"
	_geb "github.com/bamzi/pdfext/core"
	_gba "github.com/bamzi/pdfext/internal/cmap"
	_gb "github.com/bamzi/pdfext/internal/imageutil"
	_cf "github.com/bamzi/pdfext/internal/timeutils"
	_ae "github.com/bamzi/pdfext/model"
	_bae "github.com/bamzi/pdfext/model/internal/colorprofile"
	_bag "github.com/bamzi/pdfext/model/internal/docutil"
	_dbe "github.com/bamzi/pdfext/model/internal/fonts"
	_eg "github.com/bamzi/pdfext/model/xmputil"
	_ee "github.com/bamzi/pdfext/model/xmputil/pdfaextension"
	_ab "github.com/bamzi/pdfext/model/xmputil/pdfaid"
	_ag "github.com/trimmer-io/go-xmp/models/dc"
	_ba "github.com/trimmer-io/go-xmp/models/pdf"
	_b "github.com/trimmer-io/go-xmp/models/xmp_base"
	_dd "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_ge "github.com/trimmer-io/go-xmp/models/xmp_rights"
	_c "github.com/trimmer-io/go-xmp/xmp"
)

// Part gets the PDF/A version level.
func (_afgd *profile3) Part() int { return _afgd._egbf._fgb }

// NewProfile1B creates a new Profile1B with the given options.
func NewProfile1B(options *Profile1Options) *Profile1B {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_gdef(options)
	return &Profile1B{profile1{_dbbb: *options, _bcaf: _ed()}}
}
func _bfd(_cfbe *_ae.XObjectImage, _dc imageModifications) error {
	_aefg, _cbf := _cfbe.ToImage()
	if _cbf != nil {
		return _cbf
	}
	if _dc._adgf != nil {
		_cfbe.Filter = _dc._adgf
	}
	_cagd := _geb.MakeDict()
	_cagd.Set("\u0051u\u0061\u006c\u0069\u0074\u0079", _geb.MakeInteger(100))
	_cagd.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", _geb.MakeInteger(1))
	_cfbe.Decode = nil
	if _cbf = _cfbe.SetImage(_aefg, nil); _cbf != nil {
		return _cbf
	}
	_cfbe.ToPdfObject()
	return nil
}
func _fgaag(_gefbea *_ae.CompliancePdfReader) (_dfbbf []ViolatedRule) {
	var _fcada, _bbfb, _bcggd bool
	if _gefbea.ParserMetadata().HasNonConformantStream() {
		_dfbbf = []ViolatedRule{_eef("\u0036.\u0031\u002e\u0037\u002d\u0032", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _dddg := range _gefbea.GetObjectNums() {
		_gfdb, _ := _gefbea.GetIndirectObjectByNumber(_dddg)
		if _gfdb == nil {
			continue
		}
		_cgfc, _cbgg := _geb.GetStream(_gfdb)
		if !_cbgg {
			continue
		}
		if !_fcada {
			_dead := _cgfc.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _dead == nil {
				_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_fcada = true
			} else {
				_fbade, _cfadf := _geb.GetIntVal(_dead)
				if !_cfadf {
					_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0031", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_fcada = true
				} else {
					if len(_cgfc.Stream) != _fbade {
						_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_fcada = true
					}
				}
			}
		}
		if !_bbfb {
			if _cgfc.Get("\u0046") != nil {
				_bbfb = true
				_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020'\u0046\u0027,\u0020\u0027F\u0046\u0069\u006c\u0074\u0065\u0072\u0027\u002c\u0020\u006f\u0072\u0020\u0027FD\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061m\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _cgfc.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_bbfb {
				_bbfb = true
				_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020'\u0046\u0027,\u0020\u0027F\u0046\u0069\u006c\u0074\u0065\u0072\u0027\u002c\u0020\u006f\u0072\u0020\u0027FD\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061m\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _cgfc.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_bbfb {
				_bbfb = true
				_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020'\u0046\u0027,\u0020\u0027F\u0046\u0069\u006c\u0074\u0065\u0072\u0027\u002c\u0020\u006f\u0072\u0020\u0027FD\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061m\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_bcggd {
			_ecgf, _fbcea := _geb.GetName(_geb.TraceToDirectObject(_cgfc.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_fbcea {
				continue
			}
			if *_ecgf == _geb.StreamEncodingFilterNameLZW {
				_bcggd = true
				_dfbbf = append(_dfbbf, _eef("\u0036.\u0031\u002e\u0037\u002d\u0034", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _dfbbf
}
func _cba() standardType { return standardType{_fgb: 2, _ea: "\u0041"} }
func _edddc(_ccbab *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_bebad = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0032"
		_dceac = "IS\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068a\u0074\u0020\u0061\u006c\u006c\u0020\u0065m\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0073\u0020\u0069n\u0020t\u0068e\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006fG\u0049\u0044M\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079 \u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020\u0061\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070p\u0069\u006e\u0067 f\u0072\u006f\u006d \u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079p\u0068 \u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006fr\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u002e"
	)
	var _dfcga string
	if _bcca, _dgfc := _geb.GetName(_ccbab.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dgfc {
		_dfcga = _bcca.String()
	}
	if _dfcga != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _fc
	}
	if _ccbab.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _eef(_bebad, _dceac)
	}
	return _fc
}
func _efbe(_gffc *_ae.CompliancePdfReader) (_dfgb []ViolatedRule) {
	var _bfda, _fddee bool
	_cdfe := func() bool { return _bfda && _fddee }
	for _, _caead := range _gffc.GetObjectNums() {
		_baefg, _bddbg := _gffc.GetIndirectObjectByNumber(_caead)
		if _bddbg != nil {
			_gd.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _caead, _bddbg)
			continue
		}
		_ggdef, _ccegc := _geb.GetDict(_baefg)
		if !_ccegc {
			continue
		}
		_eedg, _ccegc := _geb.GetName(_ggdef.Get("\u0054\u0079\u0070\u0065"))
		if !_ccegc {
			continue
		}
		if *_eedg != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_eebcg, _ccegc := _geb.GetName(_ggdef.Get("\u0053"))
		if !_ccegc {
			if !_bfda {
				_dfgb = append(_dfgb, _eef("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_bfda = true
				if _cdfe() {
					return _dfgb
				}
			}
			continue
		}
		switch _ae.PdfActionType(*_eebcg) {
		case _ae.ActionTypeLaunch, _ae.ActionTypeSound, _ae.ActionTypeMovie, _ae.ActionTypeResetForm, _ae.ActionTypeImportData, _ae.ActionTypeJavaScript, _ae.ActionTypeHide, _ae.ActionTypeSetOCGState, _ae.ActionTypeRendition, _ae.ActionTypeTrans, _ae.ActionTypeGoTo3DView:
			if !_bfda {
				_dfgb = append(_dfgb, _eef("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_bfda = true
				if _cdfe() {
					return _dfgb
				}
			}
			continue
		case _ae.ActionTypeNamed:
			if !_fddee {
				_egbd, _cdca := _geb.GetName(_ggdef.Get("\u004e"))
				if !_cdca {
					_dfgb = append(_dfgb, _eef("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fddee = true
					if _cdfe() {
						return _dfgb
					}
					continue
				}
				switch *_egbd {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_dfgb = append(_dfgb, _eef("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fddee = true
					if _cdfe() {
						return _dfgb
					}
					continue
				}
			}
		}
	}
	return _dfgb
}

// Error implements error interface.
func (_fbe VerificationError) Error() string {
	_af := _db.Builder{}
	_af.WriteString("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u003a\u0020")
	_af.WriteString(_d.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _fbe.ConformanceLevel, _fbe.ConformanceVariant))
	_af.WriteString("\u0020\u0056\u0069\u006f\u006c\u0061\u0074\u0065\u0064\u0020\u0072\u0075l\u0065\u0073\u003a\u0020")
	for _da, _eaa := range _fbe.ViolatedRules {
		_af.WriteString(_eaa.String())
		if _da != len(_fbe.ViolatedRules)-1 {
			_af.WriteRune('\n')
		}
	}
	return _af.String()
}

// Profile2U is the implementation of the PDF/A-2U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2U struct{ profile2 }

func _abda(_bffb *_ae.CompliancePdfReader) (_aeaga []ViolatedRule) {
	var _addd, _fbbec, _dedd, _fcbca, _bdbg, _cccb, _cbcd bool
	_baefd := func() bool { return _addd && _fbbec && _dedd && _fcbca && _bdbg && _cccb && _cbcd }
	_bgfge := func(_befg *_geb.PdfObjectDictionary) bool {
		if !_addd && _befg.Get("\u0054\u0052") != nil {
			_addd = true
			_aeaga = append(_aeaga, _eef("\u0036.\u0032\u002e\u0035\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _cffgf := _befg.Get("\u0054\u0052\u0032"); !_fbbec && _cffgf != nil {
			_gecd, _bfgd := _geb.GetName(_cffgf)
			if !_bfgd || (_bfgd && *_gecd != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_fbbec = true
				_aeaga = append(_aeaga, _eef("\u0036.\u0032\u002e\u0035\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _baefd() {
					return true
				}
			}
		}
		if !_dedd && _befg.Get("\u0048\u0054\u0050") != nil {
			_dedd = true
			_aeaga = append(_aeaga, _eef("\u0036.\u0032\u002e\u0035\u002d\u0033", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020th\u0065\u0020\u0048\u0054\u0050\u0020\u006b\u0065\u0079\u002e"))
		}
		_bfec, _dbgad := _geb.GetDict(_befg.Get("\u0048\u0054"))
		if _dbgad {
			if _deaaa := _bfec.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u0054\u0079\u0070\u0065"); !_fcbca && _deaaa != nil {
				_aefeg, _gbfa := _geb.GetInt(_deaaa)
				if !_gbfa || (_gbfa && !(*_aefeg == 1 || *_aefeg == 5)) {
					_aeaga = append(_aeaga, _eef("\u0020\u0036\u002e\u0032\u002e\u0035\u002d\u0034", "\u0041\u006c\u006c\u0020\u0068\u0061\u006c\u0066\u0074\u006f\u006e\u0065\u0073\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006ce\u0020\u0073h\u0061\u006c\u006c\u0020h\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u0035 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0048\u0061l\u0066\u0074\u006fn\u0065\u0054\u0079\u0070\u0065\u0020\u006be\u0079\u002e"))
					if _baefd() {
						return true
					}
				}
			}
			if _cdde := _bfec.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u004e\u0061\u006d\u0065"); !_bdbg && _cdde != nil {
				_bdbg = true
				_aeaga = append(_aeaga, _eef("\u0036.\u0032\u002e\u0035\u002d\u0035", "\u0048\u0061\u006c\u0066\u0074o\u006e\u0065\u0073\u0020\u0069\u006e\u0020a\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0048\u0061\u006c\u0066\u0074\u006f\u006e\u0065N\u0061\u006d\u0065\u0020\u006b\u0065y\u002e"))
				if _baefd() {
					return true
				}
			}
		}
		_, _bbbd := _fdaf(_bffb)
		var _gfcf bool
		_bfdge, _dbgad := _geb.GetDict(_befg.Get("\u0047\u0072\u006fu\u0070"))
		if _dbgad {
			_, _cddb := _geb.GetName(_bfdge.Get("\u0043\u0053"))
			if _cddb {
				_gfcf = true
			}
		}
		if _caecd := _befg.Get("\u0042\u004d"); !_cccb && !_cbcd && _caecd != nil {
			_aebab, _cefb := _geb.GetName(_caecd)
			if _cefb {
				switch _aebab.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
				default:
					_cccb = true
					_aeaga = append(_aeaga, _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u004f\u006el\u0079\u0020\u0062\u006c\u0065\u006e\u0064\u0020\u006d\u006f\u0064\u0065\u0073\u0020\u0074h\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0073\u0070\u0065c\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0049\u0053O\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a2\u0030\u0030\u0038\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075e\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0042M\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0065\u0078t\u0065\u006e\u0064\u0065\u0064\u0020\u0067\u0072\u0061\u0070\u0068\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					if _baefd() {
						return true
					}
				}
				if _aebab.String() != "\u004e\u006f\u0072\u006d\u0061\u006c" && !_bbbd && !_gfcf {
					_cbcd = true
					_aeaga = append(_aeaga, _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
					if _baefd() {
						return true
					}
				}
			}
		}
		if _, _dbgad = _geb.GetDict(_befg.Get("\u0053\u004d\u0061s\u006b")); !_cbcd && _dbgad && !_bbbd && !_gfcf {
			_cbcd = true
			_aeaga = append(_aeaga, _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
			if _baefd() {
				return true
			}
		}
		if _bgdbf := _befg.Get("\u0043\u0041"); !_cbcd && _bgdbf != nil && !_bbbd && !_gfcf {
			_beeg, _adgee := _geb.GetNumberAsFloat(_bgdbf)
			if _adgee == nil && _beeg < 1.0 {
				_cbcd = true
				_aeaga = append(_aeaga, _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _baefd() {
					return true
				}
			}
		}
		if _cbdgf := _befg.Get("\u0063\u0061"); !_cbcd && _cbdgf != nil && !_bbbd && !_gfcf {
			_faaf, _gfbg := _geb.GetNumberAsFloat(_cbdgf)
			if _gfbg == nil && _faaf < 1.0 {
				_cbcd = true
				_aeaga = append(_aeaga, _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _baefd() {
					return true
				}
			}
		}
		return false
	}
	for _, _ccge := range _bffb.PageList {
		_deec := _ccge.Resources
		if _deec == nil {
			continue
		}
		if _deec.ExtGState == nil {
			continue
		}
		_faabe, _gebb := _geb.GetDict(_deec.ExtGState)
		if !_gebb {
			continue
		}
		_cebf := _faabe.Keys()
		for _, _fddd := range _cebf {
			_edcag, _dgcda := _geb.GetDict(_faabe.Get(_fddd))
			if !_dgcda {
				continue
			}
			if _bgfge(_edcag) {
				return _aeaga
			}
		}
	}
	for _, _fdbee := range _bffb.PageList {
		_dceb := _fdbee.Resources
		if _dceb == nil {
			continue
		}
		_becc, _ddecc := _geb.GetDict(_dceb.XObject)
		if !_ddecc {
			continue
		}
		for _, _ddga := range _becc.Keys() {
			_aadg, _caea := _geb.GetStream(_becc.Get(_ddga))
			if !_caea {
				continue
			}
			_fcgb, _caea := _geb.GetDict(_aadg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_caea {
				continue
			}
			_fdbc, _caea := _geb.GetDict(_fcgb.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_caea {
				continue
			}
			for _, _ccabc := range _fdbc.Keys() {
				_ecgda, _ccgea := _geb.GetDict(_fdbc.Get(_ccabc))
				if !_ccgea {
					continue
				}
				if _bgfge(_ecgda) {
					return _aeaga
				}
			}
		}
	}
	return _aeaga
}
func _fbce(_gabg *_ae.CompliancePdfReader) (_acdeca ViolatedRule) {
	_gbaaa, _eaeg := _debed(_gabg)
	if !_eaeg {
		return _fc
	}
	_ggcf, _eaeg := _geb.GetDict(_gbaaa.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_eaeg {
		return _fc
	}
	_fggg, _eaeg := _geb.GetArray(_ggcf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_eaeg {
		return _fc
	}
	for _aebe := 0; _aebe < _fggg.Len(); _aebe++ {
		_bgdae, _bfea := _geb.GetDict(_fggg.Get(_aebe))
		if !_bfea {
			continue
		}
		if _bgdae.Get("\u0041\u0041") != nil {
			return _eef("\u0036.\u0036\u002e\u0032\u002d\u0032", "\u0041\u0020F\u0069\u0065\u006cd\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079 s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061n\u0020A\u0041\u0020\u0065\u006e\u0074\u0072y f\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069on\u0061l\u002d\u0061\u0063\u0074i\u006fn\u0073 \u0064\u0069c\u0074\u0069on\u0061\u0072\u0079\u002e")
		}
	}
	return _fc
}
func _cgad(_bagag *_ae.CompliancePdfReader) (_fdag []ViolatedRule) { return _fdag }

type pageColorspaceOptimizeFunc func(_fbcbg *_bag.Document, _afe *_bag.Page, _ddd []*_bag.Image) error

func _daf(_bgf []*_bag.Image, _ef bool) error {
	_cfa := _geb.PdfObjectName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
	if _ef {
		_cfa = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	for _, _gc := range _bgf {
		if _gc.Colorspace == _cfa {
			continue
		}
		_aeb, _fbcb := _ae.NewXObjectImageFromStream(_gc.Stream)
		if _fbcb != nil {
			return _fbcb
		}
		_dbc, _fbcb := _aeb.ToImage()
		if _fbcb != nil {
			return _fbcb
		}
		_dbd, _fbcb := _dbc.ToGoImage()
		if _fbcb != nil {
			return _fbcb
		}
		var _caa _ae.PdfColorspace
		if _ef {
			_caa = _ae.NewPdfColorspaceDeviceCMYK()
			_dbd, _fbcb = _gb.CMYKConverter.Convert(_dbd)
		} else {
			_caa = _ae.NewPdfColorspaceDeviceRGB()
			_dbd, _fbcb = _gb.NRGBAConverter.Convert(_dbd)
		}
		if _fbcb != nil {
			return _fbcb
		}
		_bgd, _aae := _dbd.(_gb.Image)
		if !_aae {
			return _e.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
		}
		_dcb := _bgd.Base()
		_efd := &_ae.Image{Width: int64(_dcb.Width), Height: int64(_dcb.Height), BitsPerComponent: int64(_dcb.BitsPerComponent), ColorComponents: _dcb.ColorComponents, Data: _dcb.Data}
		_efd.SetDecode(_dcb.Decode)
		_efd.SetAlpha(_dcb.Alpha)
		if _fbcb = _aeb.SetImage(_efd, _caa); _fbcb != nil {
			return _fbcb
		}
		_aeb.ToPdfObject()
		_gc.ColorComponents = _dcb.ColorComponents
		_gc.Colorspace = _cfa
	}
	return nil
}

var _ Profile = (*Profile1A)(nil)

func _efg(_gbefg *_bag.Document, _fcf bool) error {
	_dga, _dfe := _gbefg.GetPages()
	if !_dfe {
		return nil
	}
	for _, _bdea := range _dga {
		_fefe, _ead := _bdea.GetContents()
		if !_ead {
			continue
		}
		var _acfb *_ae.PdfPageResources
		_gcfa, _ead := _bdea.GetResources()
		if _ead {
			_acfb, _ = _ae.NewPdfPageResourcesFromDict(_gcfa)
		}
		for _acae, _bgce := range _fefe {
			_bafd, _eacb := _bgce.GetData()
			if _eacb != nil {
				continue
			}
			_aad := _gg.NewContentStreamParser(string(_bafd))
			_gad, _eacb := _aad.Parse()
			if _eacb != nil {
				continue
			}
			_bea, _eacb := _dgac(_acfb, _gad, _fcf)
			if _eacb != nil {
				return _eacb
			}
			if _bea == nil {
				continue
			}
			if _eacb = (&_fefe[_acae]).SetData(_bea); _eacb != nil {
				return _eacb
			}
		}
	}
	return nil
}
func _cc() standardType { return standardType{_fgb: 3, _ea: "\u0055"} }

// Profile3U is the implementation of the PDF/A-3U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile3U struct{ profile3 }

// NewProfile3B creates a new Profile3B with the given options.
func NewProfile3B(options *Profile3Options) *Profile3B {
	if options == nil {
		options = DefaultProfile3Options()
	}
	_aged(options)
	return &Profile3B{profile3{_ebg: *options, _egbf: _eb()}}
}

// Profile2Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile2Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _ff.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _gdfegg(_edggd *_ae.CompliancePdfReader) (_bdfaed []ViolatedRule) {
	_aaba, _gcfg := _debed(_edggd)
	if !_gcfg {
		return _bdfaed
	}
	_caacf, _gcfg := _geb.GetDict(_aaba.Get("\u0050\u0065\u0072m\u0073"))
	if !_gcfg {
		return _bdfaed
	}
	_dgfe := _caacf.Keys()
	for _, _cgdc := range _dgfe {
		if _cgdc.String() != "\u0055\u0052\u0033" && _cgdc.String() != "\u0044\u006f\u0063\u004d\u0044\u0050" {
			_bdfaed = append(_bdfaed, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "\u004e\u006f\u0020\u006b\u0065\u0079\u0073 \u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0055\u0052\u0033 \u0061n\u0064\u0020\u0044\u006f\u0063\u004dD\u0050\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u002e"))
		}
	}
	return _bdfaed
}
func _bfdgc(_gccea *_ae.CompliancePdfReader) (_ebgf []ViolatedRule) {
	_cgcd := _gccea.GetObjectNums()
	for _, _accba := range _cgcd {
		_eegeg, _ffdfg := _gccea.GetIndirectObjectByNumber(_accba)
		if _ffdfg != nil {
			continue
		}
		_fegde, _befbc := _geb.GetDict(_eegeg)
		if !_befbc {
			continue
		}
		_fecag, _befbc := _geb.GetName(_fegde.Get("\u0054\u0079\u0070\u0065"))
		if !_befbc {
			continue
		}
		if _fecag.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		_agda, _ffdfg := _ae.NewPdfFilespecFromObj(_fegde)
		if _ffdfg != nil {
			continue
		}
		if _agda.EF != nil {
			if _agda.F == nil || _agda.UF == nil {
				_ebgf = append(_ebgf, _eef("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
				break
			}
			if _agda.AFRelationship == nil {
				_ebgf = append(_ebgf, _eef("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
				break
			}
			_faga, _gdegb := _ae.NewEmbeddedFileFromObject(_agda.EF)
			if _gdegb != nil {
				continue
			}
			if _db.ToLower(_faga.FileType) != "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0070\u0064\u0066" {
				_ebgf = append(_ebgf, _eef("\u0036\u002e\u0038-\u0034", "\u0041\u006c\u006c\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064 \u0066\u0069\u006c\u0065\u0073\u0020\u0073\u0068\u006fu\u006c\u0064\u0020\u0062e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0050\u0044F\u002f\u0041\u002d1\u0020\u006f\u0072\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u002e"))
				break
			}
		}
	}
	return _ebgf
}
func _dfad(_gcb *_bag.Document, _fbfg func() _ff.Time) error {
	_dbf, _gaf := _ae.NewPdfInfoFromObject(_gcb.Info)
	if _gaf != nil {
		return _gaf
	}
	if _abd := _ebcb(_dbf, _fbfg); _abd != nil {
		return _abd
	}
	_gcb.Info = _dbf.ToPdfObject()
	return nil
}
func _eef(_bg string, _ad string) ViolatedRule { return ViolatedRule{RuleNo: _bg, Detail: _ad} }
func _eb() standardType                        { return standardType{_fgb: 3, _ea: "\u0042"} }

// Profile1B is the implementation of the PDF/A-1B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1B struct{ profile1 }

func _bdcg(_geac *_ae.CompliancePdfReader) ViolatedRule {
	if _geac.ParserMetadata().HasDataAfterEOF() {
		return _eef("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _fc
}
func _fe() standardType { return standardType{_fgb: 2, _ea: "\u0055"} }
func _cb() standardType { return standardType{_fgb: 1, _ea: "\u0041"} }
func _eggbb(_aedc *_ae.PdfFont, _ebgb *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_gadc  = "\u0036.\u0033\u002e\u0035\u002d\u0033"
		_dcdgd = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0073 \u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077i\u0074\u0068\u0069n\u0020\u0061\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020\u0043\u0049\u0044\u0053\u0065\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0079\u0069\u006eg\u0020\u0077\u0068i\u0063\u0068\u0020\u0043\u0049\u0044\u0073 \u0061\u0072e\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e \u0074\u0068\u0065\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0066\u0069l\u0065,\u0020\u0061\u0073 \u0064\u0065\u0073\u0063\u0072\u0069b\u0065\u0064 \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020\u0054ab\u006c\u0065\u0020\u0035.\u00320\u002e"
	)
	var _ggbe string
	if _abccg, _ebbe := _geb.GetName(_ebgb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ebbe {
		_ggbe = _abccg.String()
	}
	switch _ggbe {
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_faba := _aedc.FontDescriptor()
		if _faba.CIDSet == nil {
			return _eef(_gadc, _dcdgd)
		}
		return _fc
	default:
		return _fc
	}
}
func _acfc(_fdd *_bag.Document, _cbgf standardType, _bbfg *_bag.OutputIntents) error {
	var (
		_gdg  *_ae.PdfOutputIntent
		_eadf error
	)
	if _fdd.Version.Minor <= 7 {
		_gdg, _eadf = _bae.NewSRGBv2OutputIntent(_cbgf.outputIntentSubtype())
	} else {
		_gdg, _eadf = _bae.NewSRGBv4OutputIntent(_cbgf.outputIntentSubtype())
	}
	if _eadf != nil {
		return _eadf
	}
	if _eadf = _bbfg.Add(_gdg.ToPdfObject()); _eadf != nil {
		return _eadf
	}
	return nil
}

// Profile1Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile1Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A-1
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _ff.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _gbec(_bdcdf *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_eccg = "\u0036.\u0033\u002e\u0033\u002d\u0032"
		_ddcc = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0054y\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020f\u006f\u0072 \u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067,\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0020\u0064\u0069c\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006f\u0047\u0049D\u004d\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0066\u0072\u006f\u006d\u0020\u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061s d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020P\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0054a\u0062\u006c\u0065\u0020\u0035\u002e\u00313"
	)
	var _ccgf string
	if _afeec, _cagdg := _geb.GetName(_bdcdf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cagdg {
		_ccgf = _afeec.String()
	}
	if _ccgf != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _fc
	}
	if _bdcdf.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _eef(_eccg, _ddcc)
	}
	return _fc
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-2 standard.
func (_fdcc *profile2) ValidateStandard(r *_ae.CompliancePdfReader) error {
	_gaeg := VerificationError{ConformanceLevel: _fdcc._ceceg._fgb, ConformanceVariant: _fdcc._ceceg._ea}
	if _gege := _caec(r); _gege != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _gege)
	}
	if _ebcf := _fedf(r); _ebcf != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ebcf)
	}
	if _ddee := _cgd(r); _ddee != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ddee)
	}
	if _ceg := _bdcg(r); _ceg != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ceg)
	}
	if _eda := _fbgfd(r); _eda != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _eda)
	}
	if _dfbb := _abeg(r); len(_dfbb) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _dfbb...)
	}
	if _fbdc := _fgaag(r); len(_fbdc) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _fbdc...)
	}
	if _cgebg := _fgabb(r); len(_cgebg) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _cgebg...)
	}
	if _gdad := _bbdbb(r); _gdad != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _gdad)
	}
	if _fbad := _gdfegg(r); len(_fbad) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _fbad...)
	}
	if _ccbf := _ggcc(r); len(_ccbf) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ccbf...)
	}
	if _edag := _faab(r); _edag != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _edag)
	}
	if _ddacb := _bdfda(r); len(_ddacb) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ddacb...)
	}
	if _ded := _abda(r); len(_ded) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ded...)
	}
	if _agfb := _fbgb(r); _agfb != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _agfb)
	}
	if _bdfg := _cega(r); len(_bdfg) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _bdfg...)
	}
	if _fddb := _cgad(r); len(_fddb) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _fddb...)
	}
	if _bdeab := _cgce(r); _bdeab != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _bdeab)
	}
	if _ffdf := _bbff(r); len(_ffdf) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ffdf...)
	}
	if _dfaf := _bbed(r, _fdcc._ceceg); len(_dfaf) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _dfaf...)
	}
	if _dbda := _gcbe(r); len(_dbda) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _dbda...)
	}
	if _cbcae := _bggcdg(r); len(_cbcae) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _cbcae...)
	}
	if _dag := _ebdd(r); len(_dag) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _dag...)
	}
	if _cabc := _bdcf(r); _cabc != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _cabc)
	}
	if _ggdf := _efbe(r); len(_ggdf) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ggdf...)
	}
	if _ccg := _dfcge(r); _ccg != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ccg)
	}
	if _gbdg := _ddadb(r, _fdcc._ceceg, false); len(_gbdg) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _gbdg...)
	}
	if _fdcc._ceceg == _cba() {
		if _fffa := _ggefb(r); len(_fffa) != 0 {
			_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _fffa...)
		}
	}
	if _bbcd := _bfdgc(r); len(_bbcd) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _bbcd...)
	}
	if _ffgf := _cebe(r); len(_ffgf) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _ffgf...)
	}
	if _gbgc := _cdfeb(r); len(_gbgc) != 0 {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _gbgc...)
	}
	if _egaf := _beef(r); _egaf != _fc {
		_gaeg.ViolatedRules = append(_gaeg.ViolatedRules, _egaf)
	}
	if len(_gaeg.ViolatedRules) > 0 {
		_f.Slice(_gaeg.ViolatedRules, func(_ffedd, _afge int) bool {
			return _gaeg.ViolatedRules[_ffedd].RuleNo < _gaeg.ViolatedRules[_afge].RuleNo
		})
		return _gaeg
	}
	return nil
}
func _edbg(_ade *_bag.Document) (*_geb.PdfObjectDictionary, bool) {
	_eggd, _cgbe := _ade.FindCatalog()
	if !_cgbe {
		return nil, false
	}
	_fdcd, _cgbe := _geb.GetArray(_eggd.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_cgbe {
		return nil, false
	}
	if _fdcd.Len() == 0 {
		return nil, false
	}
	return _geb.GetDict(_fdcd.Get(0))
}
func (_bad *documentImages) hasOnlyDeviceRGB() bool { return _bad._gdf && !_bad._cag && !_bad._cfb }

// Profile is the model.StandardImplementer enhanced by the information about the profile conformance level.
type Profile interface {
	_ae.StandardImplementer
	Conformance() string
	Part() int
}

func _aafd(_abgg *_ae.CompliancePdfReader) ViolatedRule {
	_ccad, _edefb := _abgg.GetTrailer()
	if _edefb != nil {
		_gd.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _edefb)
		return _fc
	}
	_cbdf, _ccfe := _ccad.Get("\u0052\u006f\u006f\u0074").(*_geb.PdfObjectReference)
	if !_ccfe {
		_gd.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return _fc
	}
	_deba, _ccfe := _geb.GetDict(_geb.ResolveReference(_cbdf))
	if !_ccfe {
		_gd.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return _fc
	}
	if _deba.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073") != nil {
		return _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063u\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u004f\u0043\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	}
	return _fc
}
func _cgce(_aagfb *_ae.CompliancePdfReader) (_cced ViolatedRule) {
	for _, _fgde := range _aagfb.GetObjectNums() {
		_eecd, _edbea := _aagfb.GetIndirectObjectByNumber(_fgde)
		if _edbea != nil {
			continue
		}
		_cgeg, _cada := _geb.GetStream(_eecd)
		if !_cada {
			continue
		}
		_cdgga, _cada := _geb.GetName(_cgeg.Get("\u0054\u0079\u0070\u0065"))
		if !_cada {
			continue
		}
		if *_cdgga != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_, _cada = _geb.GetName(_cgeg.Get("\u004f\u0050\u0049"))
		if _cada {
			return _eef("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		_abdga, _cada := _geb.GetName(_cgeg.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_cada {
			continue
		}
		if *_abdga == "\u0050\u0053" {
			return _eef("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _cgeg.Get("\u0050\u0053") != nil {
			return _eef("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _cced
}
func _cgfb(_fdgc *_ae.CompliancePdfReader) (_cgdb []ViolatedRule) {
	var _fefdb, _feccb, _dcce, _cbccb, _gdddfe, _afbf, _fecad bool
	_abgd := func() bool { return _fefdb && _feccb && _dcce && _cbccb && _gdddfe && _afbf && _fecad }
	for _, _ecad := range _fdgc.PageList {
		_aacdg, _afea := _ecad.GetAnnotations()
		if _afea != nil {
			_gd.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _afea)
			continue
		}
		for _, _dbcbg := range _aacdg {
			if !_fefdb {
				switch _dbcbg.GetContext().(type) {
				case *_ae.PdfAnnotationFileAttachment, *_ae.PdfAnnotationSound, *_ae.PdfAnnotationMovie, nil:
					_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020p\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e\u0020\u0041d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020F\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063\u0068\u006de\u006e\u0074\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u0020\u0061\u006e\u0064\u0020\u004d\u006f\u0076\u0069e\u0020\u0074\u0079\u0070\u0065s \u0073ha\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fefdb = true
					if _abgd() {
						return _cgdb
					}
				}
			}
			_fgc, _cbaa := _geb.GetDict(_dbcbg.GetContainingPdfObject())
			if !_cbaa {
				continue
			}
			if !_feccb {
				_efccc, _ebab := _geb.GetFloatVal(_fgc.Get("\u0043\u0041"))
				if _ebab && _efccc != 1.0 {
					_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0031", "\u0041\u006e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073h\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0043\u0041\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u002e\u0030\u002e"))
					_feccb = true
					if _abgd() {
						return _cgdb
					}
				}
			}
			if !_dcce {
				_febd, _gfde := _geb.GetIntVal(_fgc.Get("\u0046"))
				if !(_gfde && _febd&4 == 4 && _febd&1 == 0 && _febd&2 == 0 && _febd&32 == 0) {
					_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0032", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0074\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u002e\u0020\u0054\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065 s\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020I\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u0020\u0061\u006e\u0064\u0020\u004e\u006f\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020b\u0069\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073e\u0074\u0020t\u006f\u0020\u0030\u002e"))
					_dcce = true
					if _abgd() {
						return _cgdb
					}
				}
			}
			if !_cbccb {
				_gebgf, _cbde := _geb.GetDict(_fgc.Get("\u0041\u0050"))
				if _cbde {
					_eadea := _gebgf.Get("\u004e")
					if _eadea == nil || len(_gebgf.Keys()) > 1 {
						_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_cbccb = true
						if _abgd() {
							return _cgdb
						}
						continue
					}
					_, _bfffa := _dbcbg.GetContext().(*_ae.PdfAnnotationWidget)
					if _bfffa {
						_gaea, _adfc := _geb.GetName(_fgc.Get("\u0046\u0054"))
						if _adfc && *_gaea == "\u0042\u0074\u006e" {
							if _, _cfag := _geb.GetDict(_eadea); !_cfag {
								_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_cbccb = true
								if _abgd() {
									return _cgdb
								}
								continue
							}
						}
					}
					_, _adcd := _geb.GetStream(_eadea)
					if !_adcd {
						_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_cbccb = true
						if _abgd() {
							return _cgdb
						}
						continue
					}
				}
			}
			if !_gdddfe {
				if _fgc.Get("\u0043") != nil || _fgc.Get("\u0049\u0043") != nil {
					_adfb, _eagbd := _cbfe(_fdgc)
					if !_eagbd {
						_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
						_gdddfe = true
						if _abgd() {
							return _cgdb
						}
					} else {
						_bdgf, _ggegc := _geb.GetIntVal(_adfb.Get("\u004e"))
						if !_ggegc || _bdgf != 3 {
							_cgdb = append(_cgdb, _eef("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
							_gdddfe = true
							if _abgd() {
								return _cgdb
							}
						}
					}
				}
			}
			_eegad, _gedcb := _dbcbg.GetContext().(*_ae.PdfAnnotationWidget)
			if !_gedcb {
				continue
			}
			if !_afbf {
				if _eegad.A != nil {
					_cgdb = append(_cgdb, _eef("\u0036.\u0036\u002e\u0031\u002d\u0033", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_afbf = true
					if _abgd() {
						return _cgdb
					}
				}
			}
			if !_fecad {
				if _eegad.AA != nil {
					_cgdb = append(_cgdb, _eef("\u0036.\u0036\u002e\u0032\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_fecad = true
					if _abgd() {
						return _cgdb
					}
				}
			}
		}
	}
	return _cgdb
}

// DefaultProfile2Options are the default options for the Profile2.
func DefaultProfile2Options() *Profile2Options {
	return &Profile2Options{Now: _ff.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _cagb(_aage *_geb.PdfObjectDictionary, _eaedb map[*_geb.PdfObjectStream][]byte, _caeb map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_bedfg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0034"
		_fffbg = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _acfg string
	if _fbeda, _eegce := _geb.GetName(_aage.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _eegce {
		_acfg = _fbeda.String()
	}
	if _acfg != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_gffda := _aage.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ggbf := _geb.GetName(_gffda); _ggbf {
		return _fc
	}
	_dggac, _ddece := _geb.GetStream(_gffda)
	if !_ddece {
		return _eef(_bedfg, _fffbg)
	}
	_ddfbe, _daecb := _caddc(_dggac, _eaedb, _caeb)
	if _daecb != nil {
		return _eef(_bedfg, _fffbg)
	}
	_eefb, _fdcf := _geb.GetIntVal(_dggac.Get("\u0057\u004d\u006fd\u0065"))
	_fabb, _edaba := _ddfbe.WMode()
	if _fdcf && _edaba {
		if _fabb != _eefb {
			return _eef(_bedfg, _fffbg)
		}
	}
	if (_fdcf && !_edaba) || (!_fdcf && _edaba) {
		return _eef(_bedfg, _fffbg)
	}
	return _fc
}
func _geee(_aefeb standardType, _debg *_bag.OutputIntents) error {
	_gaec, _cae := _bae.NewISOCoatedV2Gray1CBasOutputIntent(_aefeb.outputIntentSubtype())
	if _cae != nil {
		return _cae
	}
	if _cae = _debg.Add(_gaec.ToPdfObject()); _cae != nil {
		return _cae
	}
	return nil
}

// StandardName gets the name of the standard.
func (_dcbfb *profile2) StandardName() string {
	return _d.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0032\u0025s", _dcbfb._ceceg._ea)
}

// Part gets the PDF/A version level.
func (_aaf *profile1) Part() int { return _aaf._bcaf._fgb }

// NewProfile1A creates a new Profile1A with given options.
func NewProfile1A(options *Profile1Options) *Profile1A {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_gdef(options)
	return &Profile1A{profile1{_dbbb: *options, _bcaf: _cb()}}
}
func _dgdb(_dabc *_bag.Document) {
	if _dabc.ID[0] != "" && _dabc.ID[1] != "" {
		return
	}
	_dabc.UseHashBasedID = true
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-1 standard.
func (_gggc *profile1) ValidateStandard(r *_ae.CompliancePdfReader) error {
	_cgeb := VerificationError{ConformanceLevel: _gggc._bcaf._fgb, ConformanceVariant: _gggc._bcaf._ea}
	if _bbga := _ffbf(r); _bbga != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _bbga)
	}
	if _cefc := _fedf(r); _cefc != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _cefc)
	}
	if _egagg := _cgd(r); _egagg != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _egagg)
	}
	if _dee := _bdcg(r); _dee != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _dee)
	}
	if _fbec := _eebc(r); _fbec != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fbec)
	}
	if _cbeg := _adef(r); len(_cbeg) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _cbeg...)
	}
	if _acdec := _gabb(r); _acdec != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _acdec)
	}
	if _caae := _abeg(r); len(_caae) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _caae...)
	}
	if _fcc := _feff(r); len(_fcc) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fcc...)
	}
	if _cea := _cfec(r); len(_cea) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _cea...)
	}
	if _eddd := _dbbe(r); _eddd != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _eddd)
	}
	if _dfdg := _ebcg(r); len(_dfdg) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _dfdg...)
	}
	if _degd := _fgbdf(r); len(_degd) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _degd...)
	}
	if _ggga := _aafd(r); _ggga != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _ggga)
	}
	if _fddg := _fdac(r, false); len(_fddg) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fddg...)
	}
	if _caaac := _bdfc(r); len(_caaac) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _caaac...)
	}
	if _aaee := _gbcb(r); _aaee != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _aaee)
	}
	if _eabc := _gdgd(r); _eabc != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _eabc)
	}
	if _fece := _bfcb(r); _fece != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fece)
	}
	if _acdcc := _fcga(r); _acdcc != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _acdcc)
	}
	if _gdcfe := _fafd(r); _gdcfe != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _gdcfe)
	}
	if _adaa := _febb(r); len(_adaa) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _adaa...)
	}
	if _ebaf := _dgba(r, _gggc._bcaf); len(_ebaf) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _ebaf...)
	}
	if _dfddf := _ccag(r); len(_dfddf) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _dfddf...)
	}
	if _gfd := _fdbef(r); _gfd != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _gfd)
	}
	if _bfdc := _ffa(r); _bfdc != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _bfdc)
	}
	if _egd := _cgfb(r); len(_egd) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _egd...)
	}
	if _fefeg := _ebgbd(r); len(_fefeg) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fefeg...)
	}
	if _cdga := _fbce(r); _cdga != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _cdga)
	}
	if _feed := _gdagde(r); _feed != _fc {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _feed)
	}
	if _fcbe := _ffgb(r, _gggc._bcaf, false); len(_fcbe) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fcbe...)
	}
	if _gggc._bcaf == _cb() {
		if _bbaa := _fffbf(r); len(_bbaa) != 0 {
			_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _bbaa...)
		}
	}
	if _fecg := _bcbdc(r); len(_fecg) != 0 {
		_cgeb.ViolatedRules = append(_cgeb.ViolatedRules, _fecg...)
	}
	if len(_cgeb.ViolatedRules) > 0 {
		_f.Slice(_cgeb.ViolatedRules, func(_ccfd, _edec int) bool {
			return _cgeb.ViolatedRules[_ccfd].RuleNo < _cgeb.ViolatedRules[_edec].RuleNo
		})
		return _cgeb
	}
	return nil
}
func _bdcf(_ecfca *_ae.CompliancePdfReader) (_bdbfc ViolatedRule) {
	_aefag, _bdca := _debed(_ecfca)
	if !_bdca {
		return _fc
	}
	_fbfdg, _bdca := _geb.GetDict(_aefag.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_bdca {
		return _fc
	}
	_fgdg, _bdca := _geb.GetArray(_fbfdg.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_bdca {
		return _fc
	}
	for _ecdeg := 0; _ecdeg < _fgdg.Len(); _ecdeg++ {
		_caca, _gfab := _geb.GetDict(_fgdg.Get(_ecdeg))
		if !_gfab {
			continue
		}
		if _caca.Get("\u0041") != nil {
			return _eef("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
		if _caca.Get("\u0041\u0041") != nil {
			return _eef("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
	}
	return _fc
}

// XmpOptions are the options used by the optimization of the XMP metadata.
type XmpOptions struct {

	// Copyright information.
	Copyright string

	// OriginalDocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	OriginalDocumentID string

	// DocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	DocumentID string

	// InstanceID is the original document identifier.
	// By default, if this field is empty the value is set to generated UUID.
	InstanceID string

	// NewDocumentVersion is a flag that defines if a document was overwritten.
	// If the new document was created this should be true. On changing given document file, and overwriting it it should be true.
	NewDocumentVersion bool

	// MarshalIndent defines marshaling indent of the XMP metadata.
	MarshalIndent string

	// MarshalPrefix defines marshaling prefix of the XMP metadata.
	MarshalPrefix string
}

func _dcbd(_fec *_bag.Document) error {
	_cece, _dgda := _fec.GetPages()
	if !_dgda {
		return nil
	}
	for _, _bagg := range _cece {
		_bdfb, _eaba := _geb.GetArray(_bagg.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_eaba {
			continue
		}
		for _, _faec := range _bdfb.Elements() {
			_faec = _geb.ResolveReference(_faec)
			if _, _defba := _faec.(*_geb.PdfObjectNull); _defba {
				continue
			}
			_eeeb, _adcg := _geb.GetDict(_faec)
			if !_adcg {
				continue
			}
			_fecf, _ := _geb.GetIntVal(_eeeb.Get("\u0046"))
			_fecf &= ^(1 << 0)
			_fecf &= ^(1 << 1)
			_fecf &= ^(1 << 5)
			_fecf &= ^(1 << 8)
			_fecf |= 1 << 2
			_eeeb.Set("\u0046", _geb.MakeInteger(int64(_fecf)))
			_gcbb := false
			if _cfef := _eeeb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _cfef != nil {
				_daba, _edea := _geb.GetName(_cfef)
				if _edea && _daba.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_gcbb = true
					if _eeeb.Get("\u0041\u0041") != nil {
						_eeeb.Remove("\u0041\u0041")
					}
					if _eeeb.Get("\u0041") != nil {
						_eeeb.Remove("\u0041")
					}
				}
				if _edea && _daba.String() == "\u0054\u0065\u0078\u0074" {
					_egac, _ := _geb.GetIntVal(_eeeb.Get("\u0046"))
					_egac |= 1 << 3
					_egac |= 1 << 4
					_eeeb.Set("\u0046", _geb.MakeInteger(int64(_egac)))
				}
			}
			_dfdd, _adcg := _geb.GetDict(_eeeb.Get("\u0041\u0050"))
			if _adcg {
				_ddbd := _dfdd.Get("\u004e")
				if _ddbd == nil {
					continue
				}
				if len(_dfdd.Keys()) > 1 {
					_dfdd.Clear()
					_dfdd.Set("\u004e", _ddbd)
				}
				if _gcbb {
					_deae, _bca := _geb.GetName(_eeeb.Get("\u0046\u0054"))
					if _bca && *_deae == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func _fcd(_aabe *_bag.Document, _eaag *_bag.Page, _dfb []*_bag.Image) error {
	for _, _afad := range _dfb {
		if _afad.SMask == nil {
			continue
		}
		_cfcd, _ccdd := _ae.NewXObjectImageFromStream(_afad.Stream)
		if _ccdd != nil {
			return _ccdd
		}
		_eff, _ccdd := _cfcd.ToImage()
		if _ccdd != nil {
			return _ccdd
		}
		_defb, _ccdd := _eff.ToGoImage()
		if _ccdd != nil {
			return _ccdd
		}
		_ddfb, _ccdd := _gb.RGBAConverter.Convert(_defb)
		if _ccdd != nil {
			return _ccdd
		}
		_ebfe := _ddfb.Base()
		_dfdc := &_ae.Image{Width: int64(_ebfe.Width), Height: int64(_ebfe.Height), BitsPerComponent: int64(_ebfe.BitsPerComponent), ColorComponents: _ebfe.ColorComponents, Data: _ebfe.Data}
		_dfdc.SetDecode(_ebfe.Decode)
		_dfdc.SetAlpha(_ebfe.Alpha)
		if _ccdd = _cfcd.SetImage(_dfdc, nil); _ccdd != nil {
			return _ccdd
		}
		_cfcd.SMask = _geb.MakeNull()
		var _deca _geb.PdfObject
		_efgc := -1
		for _efgc, _deca = range _aabe.Objects {
			if _deca == _afad.SMask.Stream {
				break
			}
		}
		if _efgc != -1 {
			_aabe.Objects = append(_aabe.Objects[:_efgc], _aabe.Objects[_efgc+1:]...)
		}
		_afad.SMask = nil
		_cfcd.ToPdfObject()
	}
	return nil
}

// Profile3B is the implementation of the PDF/A-3B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile3B struct{ profile3 }

func _ffbf(_dcc *_ae.CompliancePdfReader) ViolatedRule {
	if _dcc.ParserMetadata().HeaderPosition() != 0 {
		return _eef("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	return _fc
}
func _gdgd(_cggf *_ae.CompliancePdfReader) (_baa ViolatedRule) {
	for _, _ffga := range _cggf.GetObjectNums() {
		_gfgb, _gbeaa := _cggf.GetIndirectObjectByNumber(_ffga)
		if _gbeaa != nil {
			continue
		}
		_gfeg, _cddc := _geb.GetStream(_gfgb)
		if !_cddc {
			continue
		}
		_feaf, _cddc := _geb.GetName(_gfeg.Get("\u0054\u0079\u0070\u0065"))
		if !_cddc {
			continue
		}
		if *_feaf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _gfeg.Get("\u0052\u0065\u0066") != nil {
			return _eef("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e")
		}
	}
	return _baa
}
func _fdaf(_deebf *_ae.CompliancePdfReader) (*_ae.PdfOutputIntent, bool) {
	_adacb, _facbb := _cbfe(_deebf)
	if !_facbb {
		return nil, false
	}
	_aaea, _dfdcf := _ae.NewPdfOutputIntentFromPdfObject(_adacb)
	if _dfdcf != nil {
		return nil, false
	}
	return _aaea, true
}
func _gdef(_ebfea *Profile1Options) {
	if _ebfea.Now == nil {
		_ebfea.Now = _ff.Now
	}
}
func _fdbef(_cbcbb *_ae.CompliancePdfReader) ViolatedRule {
	for _, _geddb := range _cbcbb.GetObjectNums() {
		_dace, _bfaeb := _cbcbb.GetIndirectObjectByNumber(_geddb)
		if _bfaeb != nil {
			continue
		}
		_cgcf, _ddfd := _geb.GetStream(_dace)
		if !_ddfd {
			continue
		}
		_bdfbc, _ddfd := _geb.GetName(_cgcf.Get("\u0054\u0079\u0070\u0065"))
		if !_ddfd {
			continue
		}
		if *_bdfbc != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _cgcf.Get("\u0053\u004d\u0061s\u006b") != nil {
			return _eef("\u0036\u002e\u0034-\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e \u0053\u004d\u0061\u0073\u006b\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _fc
}

// Conformance gets the PDF/A conformance.
func (_bda *profile1) Conformance() string { return _bda._bcaf._ea }
func _fcea(_cgg bool, _bgfb standardType) (pageColorspaceOptimizeFunc, documentColorspaceOptimizeFunc) {
	var _aecg, _efaf, _gagc bool
	_bgda := func(_fad *_bag.Document, _ebc *_bag.Page, _ega []*_bag.Image) error {
		_efaf = true
		for _, _gbgb := range _ega {
			switch _gbgb.Colorspace {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				_efaf = true
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				_aecg = true
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				_gagc = true
			}
		}
		_bgc, _afg := _ebc.GetContents()
		if !_afg {
			return nil
		}
		for _, _egfc := range _bgc {
			_ccaa, _aga := _egfc.GetData()
			if _aga != nil {
				continue
			}
			_efeg := _gg.NewContentStreamParser(string(_ccaa))
			_gcbf, _aga := _efeg.Parse()
			if _aga != nil {
				continue
			}
			for _, _gfc := range *_gcbf {
				switch _gfc.Operand {
				case "\u0047", "\u0067":
					_efaf = true
				case "\u0052\u0047", "\u0072\u0067":
					_aecg = true
				case "\u004b", "\u006b":
					_gagc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_gfc.Params) == 0 {
						continue
					}
					_deab, _fbde := _geb.GetName(_gfc.Params[0])
					if !_fbde {
						continue
					}
					switch _deab.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_aecg = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_efaf = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_gagc = true
					}
				}
			}
		}
		_ddc := _ebc.FindXObjectForms()
		for _, _fgge := range _ddc {
			_bgdg := _gg.NewContentStreamParser(string(_fgge.Stream))
			_gdfe, _ecg := _bgdg.Parse()
			if _ecg != nil {
				continue
			}
			for _, _deb := range *_gdfe {
				switch _deb.Operand {
				case "\u0047", "\u0067":
					_efaf = true
				case "\u0052\u0047", "\u0072\u0067":
					_aecg = true
				case "\u004b", "\u006b":
					_gagc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_deb.Params) == 0 {
						continue
					}
					_bcdg, _fcbf := _geb.GetName(_deb.Params[0])
					if !_fcbf {
						continue
					}
					switch _bcdg.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_aecg = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_efaf = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_gagc = true
					}
				}
			}
			_gdd, _bfc := _geb.GetArray(_ebc.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
			if !_bfc {
				return nil
			}
			for _, _agge := range _gdd.Elements() {
				_gddd, _bbde := _geb.GetDict(_agge)
				if !_bbde {
					continue
				}
				_eae := _gddd.Get("\u0043")
				if _eae == nil {
					continue
				}
				_gdce, _bbde := _geb.GetArray(_eae)
				if !_bbde {
					continue
				}
				switch _gdce.Len() {
				case 0:
				case 1:
					_efaf = true
				case 3:
					_aecg = true
				case 4:
					_gagc = true
				}
			}
		}
		return nil
	}
	_bbe := func(_gedd *_bag.Document, _cfac []*_bag.Image) error {
		_bbdd, _afgc := _gedd.FindCatalog()
		if !_afgc {
			return nil
		}
		_dbac, _afgc := _bbdd.GetOutputIntents()
		if _afgc && _dbac.Len() > 0 {
			return nil
		}
		if !_afgc {
			_dbac = _bbdd.NewOutputIntents()
		}
		if !(_aecg || _gagc || _efaf) {
			return nil
		}
		defer _bbdd.SetOutputIntents(_dbac)
		if _aecg && !_gagc && !_efaf {
			return _acfc(_gedd, _bgfb, _dbac)
		}
		if _gagc && !_aecg && !_efaf {
			return _ebee(_bgfb, _dbac)
		}
		if _efaf && !_aecg && !_gagc {
			return _geee(_bgfb, _dbac)
		}
		if (_aecg && _gagc) || (_aecg && _efaf) || (_gagc && _efaf) {
			if _ggd := _daf(_cfac, _cgg); _ggd != nil {
				return _ggd
			}
			if _eega := _efg(_gedd, _cgg); _eega != nil {
				return _eega
			}
			if _aggea := _aebf(_gedd, _cgg); _aggea != nil {
				return _aggea
			}
			if _fbbf := _bbf(_gedd, _cgg); _fbbf != nil {
				return _fbbf
			}
			if _cgg {
				return _ebee(_bgfb, _dbac)
			}
			return _acfc(_gedd, _bgfb, _dbac)
		}
		return nil
	}
	return _bgda, _bbe
}

var _ Profile = (*Profile2A)(nil)

func (_fb standardType) outputIntentSubtype() _ae.PdfOutputIntentType {
	switch _fb._fgb {
	case 1:
		return _ae.PdfOutputIntentTypeA1
	case 2:
		return _ae.PdfOutputIntentTypeA2
	case 3:
		return _ae.PdfOutputIntentTypeA3
	case 4:
		return _ae.PdfOutputIntentTypeA4
	default:
		return 0
	}
}

// Profile2B is the implementation of the PDF/A-2B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2B struct{ profile2 }

// Profile1A is the implementation of the PDF/A-1A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1A struct{ profile1 }

func _cega(_gbgbb *_ae.CompliancePdfReader) (_gfdcg []ViolatedRule) {
	var _cafd, _bgdf, _cfdf, _bgca, _cedge, _dbdg, _cafe bool
	_cace := map[*_geb.PdfObjectStream]struct{}{}
	for _, _fggeca := range _gbgbb.GetObjectNums() {
		if _cafd && _bgdf && _cedge && _cfdf && _bgca && _dbdg && _cafe {
			return _gfdcg
		}
		_abdb, _baad := _gbgbb.GetIndirectObjectByNumber(_fggeca)
		if _baad != nil {
			continue
		}
		_efcdc, _dcab := _geb.GetStream(_abdb)
		if !_dcab {
			continue
		}
		if _, _dcab = _cace[_efcdc]; _dcab {
			continue
		}
		_cace[_efcdc] = struct{}{}
		_fcefe, _dcab := _geb.GetName(_efcdc.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_dcab {
			continue
		}
		if !_bgca {
			if _efcdc.Get("\u0052\u0065\u0066") != nil {
				_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0039\u002d\u0032", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_bgca = true
			}
		}
		if _fcefe.String() == "\u0050\u0053" {
			if !_dbdg {
				_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0039\u002d\u0033", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_dbdg = true
				continue
			}
		}
		if _fcefe.String() == "\u0046\u006f\u0072\u006d" {
			if _bgdf && _cfdf && _bgca {
				continue
			}
			if !_bgdf && _efcdc.Get("\u004f\u0050\u0049") != nil {
				_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d \u0058\u004f\u0062j\u0065\u0063\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
				_bgdf = true
			}
			if !_cfdf {
				if _efcdc.Get("\u0050\u0053") != nil {
					_cfdf = true
				}
				if _cbbg := _efcdc.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _cbbg != nil && !_cfdf {
					if _bfbg, _gafc := _geb.GetName(_cbbg); _gafc && *_bfbg == "\u0050\u0053" {
						_cfdf = true
					}
				}
				if _cfdf {
					_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065y \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006cu\u0065 o\u0066 \u0050\u0053\u0020\u0061\u006e\u0064\u0020t\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _fcefe.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_cafd && _efcdc.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_cafd = true
		}
		if !_cafe && _efcdc.Get("\u004f\u0050\u0049") != nil {
			_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
			_cafe = true
		}
		if !_cedge && _efcdc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_aeecg, _gcaa := _geb.GetBool(_efcdc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _gcaa && bool(*_aeecg) {
				continue
			}
			_gfdcg = append(_gfdcg, _eef("\u0036.\u0032\u002e\u0038\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_cedge = true
		}
	}
	return _gfdcg
}

// ViolatedRule is the structure that defines violated PDF/A rule.
type ViolatedRule struct {
	RuleNo string
	Detail string
}
type profile1 struct {
	_bcaf standardType
	_dbbb Profile1Options
}

func _bbf(_afef *_bag.Document, _gcg bool) error {
	_dfd, _gfa := _afef.GetPages()
	if !_gfa {
		return nil
	}
	for _, _ccab := range _dfd {
		_dbb, _aacb := _geb.GetArray(_ccab.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_aacb {
			continue
		}
		for _, _cbee := range _dbb.Elements() {
			_edbc, _aeegd := _geb.GetDict(_cbee)
			if !_aeegd {
				continue
			}
			_gab := _edbc.Get("\u0043")
			if _gab == nil {
				continue
			}
			_bgab, _aeegd := _geb.GetArray(_gab)
			if !_aeegd {
				continue
			}
			_gdbe, _gfec := _bgab.GetAsFloat64Slice()
			if _gfec != nil {
				return _gfec
			}
			switch _bgab.Len() {
			case 0, 1:
				if _gcg {
					_edbc.Set("\u0043", _geb.MakeArrayFromIntegers([]int{1, 1, 1, 1}))
				} else {
					_edbc.Set("\u0043", _geb.MakeArrayFromIntegers([]int{1, 1, 1}))
				}
			case 3:
				if _gcg {
					_bgfe, _feac, _ffd, _gabc := _g.RGBToCMYK(uint8(_gdbe[0]*255), uint8(_gdbe[1]*255), uint8(_gdbe[2]*255))
					_edbc.Set("\u0043", _geb.MakeArrayFromFloats([]float64{float64(_bgfe) / 255, float64(_feac) / 255, float64(_ffd) / 255, float64(_gabc) / 255}))
				}
			case 4:
				if !_gcg {
					_eac, _eceg, _ggef := _g.CMYKToRGB(uint8(_gdbe[0]*255), uint8(_gdbe[1]*255), uint8(_gdbe[2]*255), uint8(_gdbe[3]*255))
					_edbc.Set("\u0043", _geb.MakeArrayFromFloats([]float64{float64(_eac) / 255, float64(_eceg) / 255, float64(_ggef) / 255}))
				}
			}
		}
	}
	return nil
}
func _bdd(_add *_bag.Document) error {
	_fef := func(_ffg *_geb.PdfObjectDictionary) error {
		if _gbe := _ffg.Get("\u0053\u004d\u0061s\u006b"); _gbe != nil {
			_ffg.Set("\u0053\u004d\u0061s\u006b", _geb.MakeName("\u004e\u006f\u006e\u0065"))
		}
		_ffgg := _ffg.Get("\u0043\u0041")
		if _ffgg != nil {
			_egb, _ffge := _geb.GetNumberAsFloat(_ffgg)
			if _ffge != nil {
				_gd.Log.Debug("\u0045x\u0074\u0047S\u0074\u0061\u0074\u0065 \u006f\u0062\u006ae\u0063\u0074\u0020\u0043\u0041\u0020\u0076\u0061\u006cue\u0020\u0069\u0073 \u006e\u006ft\u0020\u0061\u0020\u0066\u006c\u006fa\u0074\u003a \u0025\u0076", _ffge)
				_egb = 0
			}
			if _egb != 1.0 {
				_ffg.Set("\u0043\u0041", _geb.MakeFloat(1.0))
			}
		}
		_ffgg = _ffg.Get("\u0063\u0061")
		if _ffgg != nil {
			_beg, _fge := _geb.GetNumberAsFloat(_ffgg)
			if _fge != nil {
				_gd.Log.Debug("\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0027\u0063\u0061\u0027\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _fge)
				_beg = 0
			}
			if _beg != 1.0 {
				_ffg.Set("\u0063\u0061", _geb.MakeFloat(1.0))
			}
		}
		_gbg := _ffg.Get("\u0042\u004d")
		if _gbg != nil {
			_fag, _ac := _geb.GetName(_gbg)
			if !_ac {
				_gd.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_fag = _geb.MakeName("")
			}
			_def := _fag.String()
			switch _def {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
			default:
				_ffg.Set("\u0042\u004d", _geb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		_aec := _ffg.Get("\u0054\u0052")
		if _aec != nil {
			_gd.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_ffg.Remove("\u0054\u0052")
		}
		_gbga := _ffg.Get("\u0054\u0052\u0032")
		if _gbga != nil {
			_dgg := _gbga.String()
			if _dgg != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_gd.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_ffg.Set("\u0054\u0052\u0032", _geb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		return nil
	}
	_ccd, _cff := _add.GetPages()
	if !_cff {
		return nil
	}
	for _, _efa := range _ccd {
		_fbeg, _edd := _efa.GetResources()
		if !_edd {
			continue
		}
		_caf, _bdg := _geb.GetDict(_fbeg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_bdg {
			return nil
		}
		_gag := _caf.Keys()
		for _, _fbed := range _gag {
			_aba, _fdf := _geb.GetDict(_caf.Get(_fbed))
			if !_fdf {
				continue
			}
			_dab := _fef(_aba)
			if _dab != nil {
				continue
			}
		}
	}
	for _, _afb := range _ccd {
		_ecf, _gee := _afb.GetContents()
		if !_gee {
			return nil
		}
		for _, _dea := range _ecf {
			_afcg, _agg := _dea.GetData()
			if _agg != nil {
				continue
			}
			_gfe := _gg.NewContentStreamParser(string(_afcg))
			_gde, _agg := _gfe.Parse()
			if _agg != nil {
				continue
			}
			for _, _fbf := range *_gde {
				if len(_fbf.Params) == 0 {
					continue
				}
				_, _bde := _geb.GetName(_fbf.Params[0])
				if !_bde {
					continue
				}
				_ffc, _feg := _afb.GetResourcesXObject()
				if !_feg {
					continue
				}
				for _, _fagd := range _ffc.Keys() {
					_cbc, _afa := _geb.GetStream(_ffc.Get(_fagd))
					if !_afa {
						continue
					}
					_dge, _afa := _geb.GetDict(_cbc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_afa {
						continue
					}
					_efaa, _afa := _geb.GetDict(_dge.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_afa {
						continue
					}
					for _, _bcf := range _efaa.Keys() {
						_bce, _bbc := _geb.GetDict(_efaa.Get(_bcf))
						if !_bbc {
							continue
						}
						_dged := _fef(_bce)
						if _dged != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _aedda(_daed *_geb.PdfObjectDictionary, _edbec map[*_geb.PdfObjectStream][]byte, _ebbg map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_ageg  = "\u0046\u006f\u0072 \u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029\u0020\u0066\u006f\u006et \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0064 \u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043I\u0044\u0053y\u0073\u0074\u0065\u006d\u0049nf\u006f\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006f\u0066\u0020i\u0074\u0073\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0061\u006e\u0064 \u0043\u004d\u0061\u0070 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0063\u006f\u006d\u0070\u0061\u0074i\u0062\u006c\u0065\u002e\u0020\u0049\u006e\u0020o\u0074\u0068\u0065\u0072\u0020\u0077\u006f\u0072\u0064\u0073\u002c\u0020\u0074\u0068\u0065\u0020R\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004fr\u0064\u0065\u0072\u0069\u006e\u0067 \u0073\u0074\u0072i\u006e\u0067\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0066\u006f\u0072 \u0074\u0068\u0061\u0074\u0020\u0066o\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u002c\u0020u\u006el\u0065ss \u0074\u0068\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073 \u0049\u0064\u0065\u006e\u0074\u0069t\u0079\u002d\u0048\u0020o\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074y\u002dV\u002e"
		_begba = "\u0036.\u0033\u002e\u0033\u002d\u0031"
	)
	var _edaa string
	if _dbcd, _defg := _geb.GetName(_daed.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _defg {
		_edaa = _dbcd.String()
	}
	if _edaa != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_bee := _daed.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _bbdb, _fegc := _geb.GetName(_bee); _fegc {
		switch _bbdb.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _fc
		}
		_aabb, _eace := _gba.LoadPredefinedCMap(_bbdb.String())
		if _eace != nil {
			return _eef(_begba, _ageg)
		}
		_baga := _aabb.CIDSystemInfo()
		if _baga.Ordering != _baga.Registry {
			return _eef(_begba, _ageg)
		}
		return _fc
	}
	_cdcg, _cddg := _geb.GetStream(_bee)
	if !_cddg {
		return _eef(_begba, _ageg)
	}
	_faegb, _gagda := _caddc(_cdcg, _edbec, _ebbg)
	if _gagda != nil {
		return _eef(_begba, _ageg)
	}
	_dagb := _faegb.CIDSystemInfo()
	if _dagb.Ordering != _dagb.Registry {
		return _eef(_begba, _ageg)
	}
	return _fc
}

var _ Profile = (*Profile1B)(nil)

func _dgcef(_cefcg *_geb.PdfObjectDictionary, _eeggbd map[*_geb.PdfObjectStream][]byte, _bfdg map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_fdcb = "\u0036.\u0033\u002e\u0033\u002d\u0033"
		_cbag = "\u0041\u006cl \u0043\u004d\u0061\u0070\u0073\u0020\u0075\u0073e\u0064 \u0077i\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072m\u0069n\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020a\u006e\u0064\u0020\u0049\u0064\u0065\u006et\u0069\u0074\u0079-\u0056\u002c\u0020\u0073\u0068a\u006c\u006c \u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064es\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u00205\u002e\u0036\u002e\u0034\u002e"
	)
	var _fcca string
	if _dfcb, _gaag := _geb.GetName(_cefcg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gaag {
		_fcca = _dfcb.String()
	}
	if _fcca != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_efbc := _cefcg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _cdgge, _dgef := _geb.GetName(_efbc); _dgef {
		switch _cdgge.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _fc
		default:
			return _eef(_fdcb, _cbag)
		}
	}
	_beca, _cgee := _geb.GetStream(_efbc)
	if !_cgee {
		return _eef(_fdcb, _cbag)
	}
	_, _dccfd := _caddc(_beca, _eeggbd, _bfdg)
	if _dccfd != nil {
		return _eef(_fdcb, _cbag)
	}
	return _fc
}
func _bcbdc(_fbcbgg *_ae.CompliancePdfReader) (_dfac []ViolatedRule) {
	for _, _bcec := range _fbcbgg.GetObjectNums() {
		_fcfee, _fdbf := _fbcbgg.GetIndirectObjectByNumber(_bcec)
		if _fdbf != nil {
			continue
		}
		_dgdc, _bagae := _geb.GetDict(_fcfee)
		if !_bagae {
			continue
		}
		_egee, _bagae := _geb.GetName(_dgdc.Get("\u0054\u0079\u0070\u0065"))
		if !_bagae {
			continue
		}
		if _egee.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_dgbce, _bagae := _geb.GetBool(_dgdc.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if !_bagae {
			return _dfac
		}
		if bool(*_dgbce) {
			_dfac = append(_dfac, _eef("\u0036\u002e\u0039-\u0031", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
	}
	return _dfac
}
func _dabae(_cfcg *_ae.PdfFont, _ebbc *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_adcee = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0036\u002d\u0033"
		_ccaac = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _dffe string
	if _caacb, _edfad := _geb.GetName(_ebbc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _edfad {
		_dffe = _caacb.String()
	}
	if _dffe != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _fc
	}
	_cggdc := _cfcg.FontDescriptor()
	_dacg, _fcgca := _geb.GetIntVal(_cggdc.Flags)
	if !_fcgca {
		_gd.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _eef(_adcee, _ccaac)
	}
	_gegaf := (uint32(_dacg) >> 3) & 1
	_badc := _gegaf != 0
	if !_badc {
		return _fc
	}
	if _ebbc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _eef(_adcee, _ccaac)
	}
	return _fc
}

type imageModifications struct {
	_fbc  *colorspaceModification
	_adgf _geb.StreamEncoder
}

func _feff(_ggb *_ae.CompliancePdfReader) (_dgeb []ViolatedRule) {
	var _gdcbe, _ddg, _cbcc bool
	if _ggb.ParserMetadata().HasNonConformantStream() {
		_dgeb = []ViolatedRule{_eef("\u0036.\u0031\u002e\u0037\u002d\u0031", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _ccc := range _ggb.GetObjectNums() {
		_aefgb, _ := _ggb.GetIndirectObjectByNumber(_ccc)
		if _aefgb == nil {
			continue
		}
		_fada, _efb := _geb.GetStream(_aefgb)
		if !_efb {
			continue
		}
		if !_gdcbe {
			_gfbe := _fada.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _gfbe == nil {
				_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_gdcbe = true
			} else {
				_adgbc, _efff := _geb.GetIntVal(_gfbe)
				if !_efff {
					_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0032", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_gdcbe = true
				} else {
					if len(_fada.Stream) != _adgbc {
						_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_gdcbe = true
					}
				}
			}
		}
		if !_ddg {
			if _fada.Get("\u0046") != nil {
				_ddg = true
				_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _fada.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_ddg {
				_ddg = true
				_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _fada.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_ddg {
				_ddg = true
				_dgeb = append(_dgeb, _eef("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_cbcc {
			_afed, _cbgd := _geb.GetName(_geb.TraceToDirectObject(_fada.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_cbgd {
				continue
			}
			if *_afed == _geb.StreamEncodingFilterNameLZW {
				_cbcc = true
				_dgeb = append(_dgeb, _eef("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0031", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _dgeb
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-1 standard.
// Implements model.StandardApplier.
func (_dgdg *profile1) ApplyStandard(document *_bag.Document) (_fcde error) {
	_fdc(document, 4)
	if _fcde = _dfad(document, _dgdg._dbbb.Now); _fcde != nil {
		return _fcde
	}
	if _fcde = _cbbd(document); _fcde != nil {
		return _fcde
	}
	_bdb, _edff := _fcea(_dgdg._dbbb.CMYKDefaultColorSpace, _dgdg._bcaf)
	_fcde = _bffd(document, []pageColorspaceOptimizeFunc{_fcd, _bdb}, []documentColorspaceOptimizeFunc{_edff})
	if _fcde != nil {
		return _fcde
	}
	_dgdb(document)
	if _fcde = _gfcc(document, _dgdg._bcaf._fgb); _fcde != nil {
		return _fcde
	}
	if _fcde = _bgea(document); _fcde != nil {
		return _fcde
	}
	if _fcde = _feb(document); _fcde != nil {
		return _fcde
	}
	if _fcde = _bdd(document); _fcde != nil {
		return _fcde
	}
	if _fcde = _gfb(document); _fcde != nil {
		return _fcde
	}
	if _dgdg._bcaf._ea == "\u0041" {
		_fagcd(document)
	}
	if _fcde = _bdf(document, _dgdg._bcaf._fgb); _fcde != nil {
		return _fcde
	}
	if _fcde = _ce(document); _fcde != nil {
		return _fcde
	}
	if _fcdg := _abc(document, _dgdg._bcaf, _dgdg._dbbb.Xmp); _fcdg != nil {
		return _fcdg
	}
	if _dgdg._bcaf == _cb() {
		if _fcde = _gfef(document); _fcde != nil {
			return _fcde
		}
	}
	if _fcde = _bbg(document); _fcde != nil {
		return _fcde
	}
	return nil
}
func _bbg(_ebf *_bag.Document) error {
	for _, _eab := range _ebf.Objects {
		_ede, _dda := _geb.GetDict(_eab)
		if !_dda {
			continue
		}
		_egc := _ede.Get("\u0054\u0079\u0070\u0065")
		if _egc == nil {
			continue
		}
		if _fagc, _fcg := _geb.GetName(_egc); _fcg && _fagc.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_gdcb, _cfbea := _geb.GetBool(_ede.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _cfbea {
			if bool(*_gdcb) {
				_ede.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _geb.MakeBool(false))
			}
		}
		_ege := _ede.Get("\u0041")
		if _ege != nil {
			_ede.Remove("\u0041")
		}
		_dbeg, _cfbea := _geb.GetArray(_ede.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if _cfbea {
			for _cgec := 0; _cgec < _dbeg.Len(); _cgec++ {
				_bcd, _cfd := _geb.GetDict(_dbeg.Get(_cgec))
				if !_cfd {
					continue
				}
				if _bcd.Get("\u0041\u0041") != nil {
					_bcd.Remove("\u0041\u0041")
				}
			}
		}
	}
	return nil
}
func _debed(_daaab *_ae.CompliancePdfReader) (*_geb.PdfObjectDictionary, bool) {
	_bbad, _cbfg := _daaab.GetTrailer()
	if _cbfg != nil {
		_gd.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _cbfg)
		return nil, false
	}
	_aafb, _egbff := _bbad.Get("\u0052\u006f\u006f\u0074").(*_geb.PdfObjectReference)
	if !_egbff {
		_gd.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return nil, false
	}
	_ebaba, _egbff := _geb.GetDict(_geb.ResolveReference(_aafb))
	if !_egbff {
		_gd.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return nil, false
	}
	return _ebaba, true
}

// String gets a string representation of the violated rule.
func (_afd ViolatedRule) String() string {
	return _d.Sprintf("\u0025\u0073\u003a\u0020\u0025\u0073", _afd.RuleNo, _afd.Detail)
}

// NewProfile2U creates a new Profile2U with the given options.
func NewProfile2U(options *Profile2Options) *Profile2U {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_ddad(options)
	return &Profile2U{profile2{_feca: *options, _ceceg: _fe()}}
}
func _dfggc(_bgge *_geb.PdfObjectDictionary, _edbeb map[*_geb.PdfObjectStream][]byte, _fggea map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_bcdc  = "\u0046\u006f\u0072\u0020\u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070o\u0073\u0069\u0074e\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029 \u0066\u006fn\u0074\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f \u0065\u006e\u0074\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074\u0073 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0066\u006fl\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0072\u0065l\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070. \u0049\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0054\u0079\u0070\u0065\u0020\u0030 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0069\u0073\u0020I\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056\u002c\u0020\u0061\u006e\u0079\u0020v\u0061\u006c\u0075\u0065\u0073\u0020\u006f\u0066\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u002c\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u002c\u0020\u0061\u006e\u0064\u0020\u0053up\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069n\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0065\u006e\u0074r\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044F\u006f\u006e\u0074\u002e\u0020\u004f\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u0072\u0072\u0065\u0073\u0070\u006f\u006e\u0064\u0069\u006e\u0067\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0069\u006e\u0020\u0062\u006f\u0074h\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006cl\u0020\u0062\u0065\u0020i\u0064en\u0074\u0069\u0063\u0061\u006c\u002c \u0061n\u0064\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0070\u0070l\u0065\u006d\u0065\u006e\u0074 \u006b\u0065\u0079\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0043I\u0044S\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0067re\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f t\u0068\u0065\u0020\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043M\u0061p\u002e"
		_deffd = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0031"
	)
	var _bdcc string
	if _dfefe, _ffefb := _geb.GetName(_bgge.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ffefb {
		_bdcc = _dfefe.String()
	}
	if _bdcc != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_baggd := _bgge.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _efgda, _bbeg := _geb.GetName(_baggd); _bbeg {
		switch _efgda.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _fc
		}
		_bafgc, _abaeb := _gba.LoadPredefinedCMap(_efgda.String())
		if _abaeb != nil {
			return _eef(_deffd, _bcdc)
		}
		_egdfg := _bafgc.CIDSystemInfo()
		if _egdfg.Ordering != _egdfg.Registry {
			return _eef(_deffd, _bcdc)
		}
		return _fc
	}
	_fbegg, _ecde := _geb.GetStream(_baggd)
	if !_ecde {
		return _eef(_deffd, _bcdc)
	}
	_dadbd, _abde := _caddc(_fbegg, _edbeb, _fggea)
	if _abde != nil {
		return _eef(_deffd, _bcdc)
	}
	_gbca := _dadbd.CIDSystemInfo()
	if _gbca.Ordering != _gbca.Registry {
		return _eef(_deffd, _bcdc)
	}
	return _fc
}

// Part gets the PDF/A version level.
func (_cga *profile2) Part() int { return _cga._ceceg._fgb }
func _abc(_bdgg *_bag.Document, _fbfd standardType, _acf XmpOptions) error {
	_ged, _gefc := _bdgg.FindCatalog()
	if !_gefc {
		return nil
	}
	var _efc *_eg.Document
	_dcg, _gefc := _ged.GetMetadata()
	if !_gefc {
		_efc = _eg.NewDocument()
	} else {
		var _dfc error
		_efc, _dfc = _eg.LoadDocument(_dcg.Stream)
		if _dfc != nil {
			return _dfc
		}
	}
	_bgg := _eg.PdfInfoOptions{InfoDict: _bdgg.Info, PdfVersion: _d.Sprintf("\u0025\u0064\u002e%\u0064", _bdgg.Version.Major, _bdgg.Version.Minor), Copyright: _acf.Copyright, Overwrite: true}
	_dcge, _gefc := _ged.GetMarkInfo()
	if _gefc {
		_dca, _abe := _geb.GetBool(_dcge.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
		if _abe && bool(*_dca) {
			_bgg.Marked = true
		}
	}
	if _caac := _efc.SetPdfInfo(&_bgg); _caac != nil {
		return _caac
	}
	if _gdcf := _efc.SetPdfAID(_fbfd._fgb, _fbfd._ea); _gdcf != nil {
		return _gdcf
	}
	_bfdd := _eg.MediaManagementOptions{OriginalDocumentID: _acf.OriginalDocumentID, DocumentID: _acf.DocumentID, InstanceID: _acf.InstanceID, NewDocumentID: !_acf.NewDocumentVersion, ModifyComment: "O\u0070\u0074\u0069\u006d\u0069\u007ae\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u006f\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064"}
	_ccf, _gefc := _geb.GetDict(_bdgg.Info)
	if _gefc {
		if _fabf, _cca := _geb.GetString(_ccf.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _cca && _fabf.String() != "" {
			_gaa, _agf := _cf.ParsePdfTime(_fabf.String())
			if _agf != nil {
				return _d.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _agf)
			}
			_bfdd.ModifyDate = _gaa
		}
	}
	if _ddf := _efc.SetMediaManagement(&_bfdd); _ddf != nil {
		return _ddf
	}
	if _aebb := _efc.SetPdfAExtension(); _aebb != nil {
		return _aebb
	}
	_egf, _ddff := _efc.MarshalIndent(_acf.MarshalPrefix, _acf.MarshalIndent)
	if _ddff != nil {
		return _ddff
	}
	if _ace := _ged.SetMetadata(_egf); _ace != nil {
		return _ace
	}
	return nil
}
func _beac(_afda *_bag.Document) error {
	_dcbf, _aaaf := _afda.FindCatalog()
	if !_aaaf {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_gbbg, _aaaf := _geb.GetDict(_dcbf.Object.Get("\u004e\u0061\u006de\u0073"))
	if !_aaaf {
		return nil
	}
	if _gbbg.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_gbbg.Remove("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073")
	}
	return nil
}
func _fffbf(_ddccf *_ae.CompliancePdfReader) (_eddf []ViolatedRule) {
	_cfca := true
	_cacf, _dfcgc := _ddccf.GetCatalogMarkInfo()
	if !_dfcgc {
		_cfca = false
	} else {
		_bbaf, _dgefe := _geb.GetDict(_cacf)
		if _dgefe {
			_fegga, _cedg := _geb.GetBool(_bbaf.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_fegga) || !_cedg {
				_cfca = false
			}
		} else {
			_cfca = false
		}
	}
	if !_cfca {
		_eddf = append(_eddf, _eef("\u0036.\u0038\u002e\u0032\u002e\u0032\u002d1", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_dgebb, _dfcgc := _ddccf.GetCatalogStructTreeRoot()
	if !_dfcgc {
		_eddf = append(_eddf, _eef("\u0036.\u0038\u002e\u0033\u002e\u0033\u002d1", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_egde, _dfcgc := _geb.GetDict(_dgebb)
	if _dfcgc {
		_faa, _fcbd := _geb.GetName(_egde.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _fcbd {
			_dbdeb, _fadf := _geb.GetDict(_faa)
			if _fadf {
				for _, _abae := range _dbdeb.Keys() {
					_edbba := _dbdeb.Get(_abae)
					if _edbba == nil {
						_eddf = append(_eddf, _eef("\u0036.\u0038\u002e\u0033\u002e\u0034\u002d1", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _eddf
}

// NewProfile2A creates a new Profile2A with given options.
func NewProfile2A(options *Profile2Options) *Profile2A {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_ddad(options)
	return &Profile2A{profile2{_feca: *options, _ceceg: _cba()}}
}
func _dfcge(_gabe *_ae.CompliancePdfReader) (_fgcg ViolatedRule) {
	_dgca, _bgbfg := _debed(_gabe)
	if !_bgbfg {
		return _fc
	}
	if _dgca.Get("\u0041\u0041") != nil {
		return _eef("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0054h\u0065\u0020\u0064\u006fc\u0075m\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020a\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079 \u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079\u002e")
	}
	return _fc
}
func _gebd(_eggbe *_ae.PdfFont, _bccg *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_eced = "\u0036.\u0033\u002e\u0037\u002d\u0032"
		_deag = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _abbe string
	if _dacd, _cgbdf := _geb.GetName(_bccg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cgbdf {
		_abbe = _dacd.String()
	}
	if _abbe != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _fc
	}
	_fffea := _eggbe.FontDescriptor()
	_gbad, _edfc := _geb.GetIntVal(_fffea.Flags)
	if !_edfc {
		_gd.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _eef(_eced, _deag)
	}
	_dcdgg := (uint32(_gbad) >> 3) & 1
	_dbbbg := _dcdgg != 0
	if !_dbbbg {
		return _fc
	}
	if _bccg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _eef(_eced, _deag)
	}
	return _fc
}
func _cbfe(_aabd *_ae.CompliancePdfReader) (*_geb.PdfObjectDictionary, bool) {
	_fccf, _gbaf := _debed(_aabd)
	if !_gbaf {
		return nil, false
	}
	_aaaa, _gbaf := _geb.GetArray(_fccf.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_gbaf {
		return nil, false
	}
	if _aaaa.Len() == 0 {
		return nil, false
	}
	return _geb.GetDict(_aaaa.Get(0))
}

// NewProfile3U creates a new Profile3U with the given options.
func NewProfile3U(options *Profile3Options) *Profile3U {
	if options == nil {
		options = DefaultProfile3Options()
	}
	_aged(options)
	return &Profile3U{profile3{_ebg: *options, _egbf: _cc()}}
}
func _ed() standardType { return standardType{_fgb: 1, _ea: "\u0042"} }
func _faab(_efegb *_ae.CompliancePdfReader) ViolatedRule {
	for _, _fdee := range _efegb.PageList {
		_afcf, _abba := _fdee.GetContentStreams()
		if _abba != nil {
			continue
		}
		for _, _gddgg := range _afcf {
			_ccfg := _gg.NewContentStreamParser(_gddgg)
			_, _abba = _ccfg.Parse()
			if _abba != nil {
				return _eef("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _fc
}

// StandardName gets the name of the standard.
func (_efaag *profile1) StandardName() string {
	return _d.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0031\u0025s", _efaag._bcaf._ea)
}
func _bfcb(_fadb *_ae.CompliancePdfReader) (_afbbd ViolatedRule) {
	for _, _badda := range _fadb.GetObjectNums() {
		_edgb, _abbfc := _fadb.GetIndirectObjectByNumber(_badda)
		if _abbfc != nil {
			continue
		}
		_debd, _fefdf := _geb.GetStream(_edgb)
		if !_fefdf {
			continue
		}
		_fbecd, _fefdf := _geb.GetName(_debd.Get("\u0054\u0079\u0070\u0065"))
		if !_fefdf {
			continue
		}
		if *_fbecd != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_fcabg, _fefdf := _geb.GetName(_debd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_fefdf {
			continue
		}
		if *_fcabg == "\u0050\u0053" {
			return _eef("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073.")
		}
	}
	return _afbbd
}
func _adef(_affgg *_ae.CompliancePdfReader) (_fffd []ViolatedRule) {
	_gbff := _affgg.ParserMetadata()
	if _gbff.HasInvalidSubsectionHeader() {
		_fffd = append(_fffd, _eef("\u0036.\u0031\u002e\u0034\u002d\u0031", "\u006e\u0020\u0061\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0073\u0065c\u0074\u0069\u006f\u006e\u0020h\u0065a\u0064\u0065\u0072\u0020t\u0068\u0065\u0020\u0073\u0074\u0061\u0072t\u0069\u006e\u0067\u0020\u006fb\u006a\u0065\u0063\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0072\u0061n\u0067e\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0053\u0050\u0041C\u0045\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u0028\u0032\u0030\u0068\u0029\u002e"))
	}
	if _gbff.HasInvalidSeparationAfterXRef() {
		_fffd = append(_fffd, _eef("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e"))
	}
	return _fffd
}

// VerificationError is the PDF/A verification error structure, that contains all violated rules.
type VerificationError struct {

	// ViolatedRules are the rules that were violated during error verification.
	ViolatedRules []ViolatedRule

	// ConformanceLevel defines the standard on verification failed.
	ConformanceLevel int

	// ConformanceVariant is the standard variant used on verification.
	ConformanceVariant string
}

func _ca() standardType { return standardType{_fgb: 2, _ea: "\u0042"} }
func _feb(_gebf *_bag.Document) error {
	_acdf, _cfed := _gebf.GetPages()
	if !_cfed {
		return nil
	}
	for _, _adb := range _acdf {
		_bgdc, _dcadd := _geb.GetArray(_adb.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_dcadd {
			continue
		}
		for _, _agac := range _bgdc.Elements() {
			_agac = _geb.ResolveReference(_agac)
			if _, _bggb := _agac.(*_geb.PdfObjectNull); _bggb {
				continue
			}
			_agdb, _aag := _geb.GetDict(_agac)
			if !_aag {
				continue
			}
			_cddd, _ := _geb.GetIntVal(_agdb.Get("\u0046"))
			_cddd &= ^(1 << 0)
			_cddd &= ^(1 << 1)
			_cddd &= ^(1 << 5)
			_cddd |= 1 << 2
			_agdb.Set("\u0046", _geb.MakeInteger(int64(_cddd)))
			_afee := false
			if _ddb := _agdb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _ddb != nil {
				_dbcc, _affe := _geb.GetName(_ddb)
				if _affe && _dbcc.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_afee = true
					if _agdb.Get("\u0041\u0041") != nil {
						_agdb.Remove("\u0041\u0041")
					}
				}
			}
			if _agdb.Get("\u0043") != nil || _agdb.Get("\u0049\u0043") != nil {
				_ceea, _adeb := _edbg(_gebf)
				if !_adeb {
					_agdb.Remove("\u0043")
					_agdb.Remove("\u0049\u0043")
				} else {
					_ddec, _dcf := _geb.GetIntVal(_ceea.Get("\u004e"))
					if !_dcf || _ddec != 3 {
						_agdb.Remove("\u0043")
						_agdb.Remove("\u0049\u0043")
					}
				}
			}
			_cdac, _aag := _geb.GetDict(_agdb.Get("\u0041\u0050"))
			if _aag {
				_feee := _cdac.Get("\u004e")
				if _feee == nil {
					continue
				}
				if len(_cdac.Keys()) > 1 {
					_cdac.Clear()
					_cdac.Set("\u004e", _feee)
				}
				if _afee {
					_eec, _bedf := _geb.GetName(_agdb.Get("\u0046\u0054"))
					if _bedf && *_eec == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func _cgd(_gfcab *_ae.CompliancePdfReader) ViolatedRule {
	_cffd, _gbed := _gfcab.PdfReader.GetTrailer()
	if _gbed != nil {
		return _eef("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _cffd.Get("\u0049\u0044") == nil {
		return _eef("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _cffd.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _eef("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _fc
}
func _fcad(_ecgb string, _ffdd string, _cfeg string) (string, bool) {
	_ddgc := _db.Index(_ecgb, _ffdd)
	if _ddgc == -1 {
		return "", false
	}
	_ddgc += len(_ffdd)
	_degdf := _db.Index(_ecgb[_ddgc:], _cfeg)
	if _degdf == -1 {
		return "", false
	}
	_degdf = _ddgc + _degdf
	return _ecgb[_ddgc:_degdf], true
}
func _gfcc(_eaea *_bag.Document, _aeda int) error {
	_fbfc := map[*_geb.PdfObjectStream]struct{}{}
	for _, _fdfd := range _eaea.Objects {
		_fcgc, _abdd := _geb.GetStream(_fdfd)
		if !_abdd {
			continue
		}
		if _, _abdd = _fbfc[_fcgc]; _abdd {
			continue
		}
		_fbfc[_fcgc] = struct{}{}
		_edfa, _abdd := _geb.GetName(_fcgc.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_abdd {
			continue
		}
		if _fcgc.Get("\u0052\u0065\u0066") != nil {
			_fcgc.Remove("\u0052\u0065\u0066")
		}
		if _edfa.String() == "\u0050\u0053" {
			_fcgc.Remove("\u0050\u0053")
			continue
		}
		if _edfa.String() == "\u0046\u006f\u0072\u006d" {
			if _fcgc.Get("\u004f\u0050\u0049") != nil {
				_fcgc.Remove("\u004f\u0050\u0049")
			}
			if _fcgc.Get("\u0050\u0053") != nil {
				_fcgc.Remove("\u0050\u0053")
			}
			if _gda := _fcgc.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _gda != nil {
				if _fgbb, _fee := _geb.GetName(_gda); _fee && *_fgbb == "\u0050\u0053" {
					_fcgc.Remove("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032")
				}
			}
			continue
		}
		if _edfa.String() == "\u0049\u006d\u0061g\u0065" {
			_feea, _fggb := _geb.GetBool(_fcgc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _fggb && bool(*_feea) {
				_fcgc.Set("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _geb.MakeBool(false))
			}
			if _aeda == 2 {
				if _fcgc.Get("\u004f\u0050\u0049") != nil {
					_fcgc.Remove("\u004f\u0050\u0049")
				}
			}
			if _fcgc.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
				_fcgc.Remove("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073")
			}
			continue
		}
	}
	return nil
}
func _dgac(_aaa *_ae.PdfPageResources, _aefec *_gg.ContentStreamOperations, _ggeg bool) ([]byte, error) {
	var _fbcba bool
	for _, _egeg := range *_aefec {
	_agc:
		switch _egeg.Operand {
		case "\u0042\u0049":
			_ddffb, _aff := _egeg.Params[0].(*_gg.ContentStreamInlineImage)
			if !_aff {
				break
			}
			_acdb, _bafc := _ddffb.GetColorSpace(_aaa)
			if _bafc != nil {
				return nil, _bafc
			}
			switch _acdb.(type) {
			case *_ae.PdfColorspaceDeviceCMYK:
				if _ggeg {
					break _agc
				}
			case *_ae.PdfColorspaceDeviceGray:
			case *_ae.PdfColorspaceDeviceRGB:
				if !_ggeg {
					break _agc
				}
			default:
				break _agc
			}
			_fbcba = true
			_cgbec, _bafc := _ddffb.ToImage(_aaa)
			if _bafc != nil {
				return nil, _bafc
			}
			_dgbg, _bafc := _cgbec.ToGoImage()
			if _bafc != nil {
				return nil, _bafc
			}
			if _ggeg {
				_dgbg, _bafc = _gb.CMYKConverter.Convert(_dgbg)
			} else {
				_dgbg, _bafc = _gb.NRGBAConverter.Convert(_dgbg)
			}
			if _bafc != nil {
				return nil, _bafc
			}
			_beb, _aff := _dgbg.(_gb.Image)
			if !_aff {
				return nil, _e.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
			}
			_abb := _beb.Base()
			_gca := _ae.Image{Width: int64(_abb.Width), Height: int64(_abb.Height), BitsPerComponent: int64(_abb.BitsPerComponent), ColorComponents: _abb.ColorComponents, Data: _abb.Data}
			_gca.SetDecode(_abb.Decode)
			_gca.SetAlpha(_abb.Alpha)
			_bfddb, _bafc := _ddffb.GetEncoder()
			if _bafc != nil {
				_bfddb = _geb.NewFlateEncoder()
			}
			_gdcc, _bafc := _gg.NewInlineImageFromImage(_gca, _bfddb)
			if _bafc != nil {
				return nil, _bafc
			}
			_egeg.Params[0] = _gdcc
		case "\u0047", "\u0067":
			if len(_egeg.Params) != 1 {
				break
			}
			_abab, _efge := _geb.GetNumberAsFloat(_egeg.Params[0])
			if _efge != nil {
				break
			}
			if _ggeg {
				_egeg.Params = []_geb.PdfObject{_geb.MakeFloat(0), _geb.MakeFloat(0), _geb.MakeFloat(0), _geb.MakeFloat(1 - _abab)}
				_eaaf := "\u004b"
				if _egeg.Operand == "\u0067" {
					_eaaf = "\u006b"
				}
				_egeg.Operand = _eaaf
			} else {
				_egeg.Params = []_geb.PdfObject{_geb.MakeFloat(_abab), _geb.MakeFloat(_abab), _geb.MakeFloat(_abab)}
				_bcgg := "\u0052\u0047"
				if _egeg.Operand == "\u0067" {
					_bcgg = "\u0072\u0067"
				}
				_egeg.Operand = _bcgg
			}
			_fbcba = true
		case "\u0052\u0047", "\u0072\u0067":
			if !_ggeg {
				break
			}
			if len(_egeg.Params) != 3 {
				break
			}
			_cab, _gec := _geb.GetNumbersAsFloat(_egeg.Params)
			if _gec != nil {
				break
			}
			_fbcba = true
			_agcb, _eggb, _gdcfa := _cab[0], _cab[1], _cab[2]
			_ada, _eegae, _gbc, _bfb := _g.RGBToCMYK(uint8(_agcb*255), uint8(_eggb*255), uint8(255*_gdcfa))
			_egeg.Params = []_geb.PdfObject{_geb.MakeFloat(float64(_ada) / 255), _geb.MakeFloat(float64(_eegae) / 255), _geb.MakeFloat(float64(_gbc) / 255), _geb.MakeFloat(float64(_bfb) / 255)}
			_dce := "\u004b"
			if _egeg.Operand == "\u0072\u0067" {
				_dce = "\u006b"
			}
			_egeg.Operand = _dce
		case "\u004b", "\u006b":
			if _ggeg {
				break
			}
			if len(_egeg.Params) != 4 {
				break
			}
			_ffb, _ebe := _geb.GetNumbersAsFloat(_egeg.Params)
			if _ebe != nil {
				break
			}
			_acde, _dbde, _gcac, _edef := _ffb[0], _ffb[1], _ffb[2], _ffb[3]
			_badb, _ebfa, _ebffd := _g.CMYKToRGB(uint8(255*_acde), uint8(255*_dbde), uint8(255*_gcac), uint8(255*_edef))
			_egeg.Params = []_geb.PdfObject{_geb.MakeFloat(float64(_badb) / 255), _geb.MakeFloat(float64(_ebfa) / 255), _geb.MakeFloat(float64(_ebffd) / 255)}
			_daa := "\u0052\u0047"
			if _egeg.Operand == "\u006b" {
				_daa = "\u0072\u0067"
			}
			_egeg.Operand = _daa
			_fbcba = true
		}
	}
	if !_fbcba {
		return nil, nil
	}
	_agcbg := _gg.NewContentCreator()
	for _, _bfg := range *_aefec {
		_agcbg.AddOperand(*_bfg)
	}
	_adc := _agcbg.Bytes()
	return _adc, nil
}
func _fafd(_cacd *_ae.CompliancePdfReader) ViolatedRule {
	for _, _fbef := range _cacd.PageList {
		_gbfd, _dcdc := _fbef.GetContentStreams()
		if _dcdc != nil {
			continue
		}
		for _, _afbg := range _gbfd {
			_feffa := _gg.NewContentStreamParser(_afbg)
			_, _dcdc = _feffa.Parse()
			if _dcdc != nil {
				return _eef("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _fc
}
func _de(_aed []_geb.PdfObject) (*documentImages, error) {
	_aee := _geb.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_edf := make(map[*_geb.PdfObjectStream]struct{})
	_eca := make(map[_geb.PdfObject]struct{})
	var (
		_aa, _cbe, _bga bool
		_ga             []*imageInfo
		_aac            error
	)
	for _, _gdc := range _aed {
		_gf, _abg := _geb.GetStream(_gdc)
		if !_abg {
			continue
		}
		if _, _cdg := _edf[_gf]; _cdg {
			continue
		}
		_edf[_gf] = struct{}{}
		_be := _gf.PdfObjectDictionary.Get(_aee)
		_fgd, _abg := _geb.GetName(_be)
		if !_abg || string(*_fgd) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if _egg := _gf.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b"); _egg != nil {
			_eca[_egg] = struct{}{}
		}
		_fab := imageInfo{BitsPerComponent: 8, Stream: _gf}
		_fab.ColorSpace, _aac = _ae.DetermineColorspaceNameFromPdfObject(_gf.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _aac != nil {
			return nil, _aac
		}
		if _afc, _eeg := _geb.GetIntVal(_gf.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _eeg {
			_fab.BitsPerComponent = _afc
		}
		if _bb, _eag := _geb.GetIntVal(_gf.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _eag {
			_fab.Width = _bb
		}
		if _fea, _cbg := _geb.GetIntVal(_gf.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _cbg {
			_fab.Height = _fea
		}
		switch _fab.ColorSpace {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_bga = true
			_fab.ColorComponents = 1
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_aa = true
			_fab.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_cbe = true
			_fab.ColorComponents = 4
		default:
			_fab._fga = true
		}
		_ga = append(_ga, &_fab)
	}
	if len(_eca) > 0 {
		if len(_eca) == len(_ga) {
			_ga = nil
		} else {
			_eagc := make([]*imageInfo, len(_ga)-len(_eca))
			var _fd int
			for _, _aefe := range _ga {
				if _, _gef := _eca[_aefe.Stream]; _gef {
					continue
				}
				_eagc[_fd] = _aefe
				_fd++
			}
			_ga = _eagc
		}
	}
	return &documentImages{_gdf: _aa, _cag: _cbe, _cfb: _bga, _cd: _eca, _fbb: _ga}, nil
}

type documentImages struct {
	_gdf, _cag, _cfb bool
	_cd              map[_geb.PdfObject]struct{}
	_fbb             []*imageInfo
}

var _ Profile = (*Profile3U)(nil)

// NewProfile2B creates a new Profile2B with the given options.
func NewProfile2B(options *Profile2Options) *Profile2B {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_ddad(options)
	return &Profile2B{profile2{_feca: *options, _ceceg: _ca()}}
}
func _ce(_gbef *_bag.Document) error {
	_ffe, _gae := _gbef.FindCatalog()
	if !_gae {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_, _gae = _geb.GetDict(_ffe.Object.Get("\u0041\u0041"))
	if !_gae {
		return nil
	}
	_ffe.Object.Remove("\u0041\u0041")
	return nil
}
func _ffgb(_dbggb *_ae.CompliancePdfReader, _eegeb standardType, _aebg bool) (_ceef []ViolatedRule) {
	_dfdgdc, _bgbb := _debed(_dbggb)
	if !_bgbb {
		return []ViolatedRule{_eef("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_fcbc := _dfdgdc.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _fcbc == nil {
		return []ViolatedRule{_eef("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u006e\u006f\u0020\u0027\u004d\u0065\u0074\u0061d\u0061\u0074\u0061' \u006b\u0065\u0079\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u002e"), _eef("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	_efgf, _bgbb := _geb.GetStream(_fcbc)
	if !_bgbb {
		return []ViolatedRule{_eef("\u0036.\u0037\u002e\u0032\u002d\u0032", "\u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0027\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"), _eef("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	if _efgf.Get("\u0046\u0069\u006c\u0074\u0065\u0072") != nil {
		_ceef = append(_ceef, _eef("\u0036.\u0037\u002e\u0032\u002d\u0032", "M\u0065\u0074a\u0064\u0061\u0074\u0061\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u006b\u0065y\u002e"))
	}
	_egdg, _abgf := _eg.LoadDocument(_efgf.Stream)
	if _abgf != nil {
		return []ViolatedRule{_eef("\u0036.\u0037\u002e\u0039\u002d\u0031", "The\u0020\u006d\u0065\u0074a\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0061\u006e\u0064\u0020\u0077\u0065\u006c\u006c\u0020\u0066\u006f\u0072\u006de\u0064\u0020\u0050\u0044\u0046\u0041\u0045\u0078\u0074e\u006e\u0073\u0069\u006f\u006e\u0020\u0053\u0063\u0068\u0065\u006da\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073\u002e")}
	}
	_cbddg := _egdg.GetGoXmpDocument()
	var _gaada []*_c.Namespace
	for _, _aeac := range _cbddg.Namespaces() {
		switch _aeac.Name {
		case _ag.NsDc.Name, _ba.NsPDF.Name, _b.NsXmp.Name, _ge.NsXmpRights.Name, _ab.Namespace.Name, _ee.Namespace.Name, _dd.NsXmpMM.Name, _ee.FieldNS.Name, _ee.SchemaNS.Name, _ee.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_gaada = append(_gaada, _aeac)
	}
	_gfeb := true
	_bcba, _abgf := _egdg.GetPdfaExtensionSchemas()
	if _abgf == nil {
		for _, _gfgc := range _gaada {
			var _caad bool
			for _degf := range _bcba {
				if _gfgc.URI == _bcba[_degf].NamespaceURI {
					_caad = true
					break
				}
			}
			if !_caad {
				_gfeb = false
				break
			}
		}
	} else {
		_gfeb = false
	}
	if !_gfeb {
		_ceef = append(_ceef, _eef("\u0036.\u0037\u002e\u0039\u002d\u0032", "\u0050\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073 \u0073\u0070\u0065\u0063\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0058M\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0073\u0068\u0061\u006cl\u0020\u0075\u0073\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073 \u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002c\u0020\u006f\u0072\u0020\u0065\u0078\u0074\u0065ns\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u0074\u0068\u0061\u0074 \u0063\u006f\u006d\u0070\u006c\u0079\u0020\u0077\u0069\u0074h\u0020\u0058\u004d\u0050\u0020\u0053\u0070e\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002e"))
	}
	_bdff, _abgf := _dbggb.GetPdfInfo()
	if _abgf == nil {
		if !_acfcb(_bdff, _egdg) {
			_ceef = append(_ceef, _eef("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
		}
	} else if _, _gaeag := _egdg.GetMediaManagement(); _gaeag {
		_ceef = append(_ceef, _eef("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
	}
	_bbge, _bgbb := _egdg.GetPdfAID()
	if !_bgbb {
		_ceef = append(_ceef, _eef("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _bbge.Part != _eegeb._fgb {
			_ceef = append(_ceef, _eef("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _eegeb._ea == "\u0041" && _bbge.Conformance != "\u0041" {
			_ceef = append(_ceef, _eef("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		} else if _eegeb._ea == "\u0042" && (_bbge.Conformance != "\u0041" && _bbge.Conformance != "\u0042") {
			_ceef = append(_ceef, _eef("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		}
	}
	return _ceef
}
func _gfb(_dfg *_bag.Document) error {
	_bgb := map[string]*_geb.PdfObjectDictionary{}
	_dbdb := _fg.NewFinder(&_fg.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	_fff := map[_geb.PdfObject]struct{}{}
	_aefa := map[_geb.PdfObject]struct{}{}
	for _, _fda := range _dfg.Objects {
		_aeec, _fegf := _geb.GetDict(_fda)
		if !_fegf {
			continue
		}
		_aca := _aeec.Get("\u0054\u0079\u0070\u0065")
		if _aca == nil {
			continue
		}
		if _cgb, _ccb := _geb.GetName(_aca); _ccb && _cgb.String() != "\u0046\u006f\u006e\u0074" {
			continue
		}
		if _, _faf := _fff[_fda]; _faf {
			continue
		}
		_edb, _ccdb := _ae.NewPdfFontFromPdfObject(_aeec)
		if _ccdb != nil {
			_gd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
			return _ccdb
		}
		if _edb.Encoder() != nil && (_edb.Encoder().String() == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _edb.Encoder().String() == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056") {
			continue
		}
		if _edb.Subtype() == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
			_fgg := _edb.GetCIDToGIDMapObject()
			if _fgg != nil {
				continue
			}
		}
		_ggg, _ccdb := _edb.GetFontDescriptor()
		if _ccdb != nil {
			return _ccdb
		}
		if _ggg != nil && (_ggg.FontFile != nil || _ggg.FontFile2 != nil || _ggg.FontFile3 != nil) {
			continue
		}
		_efdg := _edb.BaseFont()
		if _efdg == "" {
			_ggc, _efad := _edb.GetFontDescriptor()
			if _efad != nil {
				return _d.Errorf("\u0063\u0061\u006e\u0027\u0074\u0020\u0067\u0065t\u0020\u0074\u0068e \u0066\u006f\u006e\u0074\u0020\u006ea\u006d\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u003a \u0025\u0073", _aeec.String())
			}
			_efdg = _ggc.FontName.String()
			if _efdg == "" {
				return _d.Errorf("\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0073\u0020\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069d\u0020\u002d\u0020\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u0075\u006ed\u0065\u0066\u0069n\u0065\u0064\u003a\u0020\u0025\u0073", _aeec.String())
			}
		}
		_dbg, _gbeg := _bgb[_efdg]
		if !_gbeg {
			if len(_efdg) > 7 && _efdg[6] == '+' {
				_efdg = _efdg[7:]
			}
			_cbb := []string{_efdg, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
			for _, _gbd := range _cbb {
				_gd.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _gbd)
				if _dbg, _gbeg = _bgb[_gbd]; _gbeg {
					break
				}
				_gcf := _dbdb.Match(_gbd)
				if _gcf == nil {
					_gd.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _gbd)
					continue
				}
				_fdff, _dgd := _ae.NewPdfFontFromTTFFile(_gcf.Filename)
				if _dgd != nil {
					return _dgd
				}
				_bdge := _fdff.FontDescriptor()
				if _bdge.FontFile != nil {
					if _, _gbeg = _aefa[_bdge.FontFile]; !_gbeg {
						_dfg.Objects = append(_dfg.Objects, _bdge.FontFile)
						_aefa[_bdge.FontFile] = struct{}{}
					}
				}
				if _bdge.FontFile2 != nil {
					if _, _gbeg = _aefa[_bdge.FontFile2]; !_gbeg {
						_dfg.Objects = append(_dfg.Objects, _bdge.FontFile2)
						_aefa[_bdge.FontFile2] = struct{}{}
					}
				}
				if _bdge.FontFile3 != nil {
					if _, _gbeg = _aefa[_bdge.FontFile3]; !_gbeg {
						_dfg.Objects = append(_dfg.Objects, _bdge.FontFile3)
						_aefa[_bdge.FontFile3] = struct{}{}
					}
				}
				_dad, _cfbdd := _fdff.ToPdfObject().(*_geb.PdfIndirectObject)
				if !_cfbdd {
					_gd.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
					continue
				}
				_defc, _cfbdd := _dad.PdfObject.(*_geb.PdfObjectDictionary)
				if !_cfbdd {
					_gd.Log.Debug("\u0046\u006fn\u0074\u0020\u0074\u0079p\u0065\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
					continue
				}
				_bgb[_gbd] = _defc
				_dbg = _defc
				break
			}
			if _dbg == nil {
				_gd.Log.Debug("\u004e\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006eg\u0020\u0066\u006f\u006e\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u003a\u0020\u0025\u0073", _edb.BaseFont())
				return _e.New("\u006e\u006f m\u0061\u0074\u0063h\u0069\u006e\u0067\u0020fon\u0074 f\u006f\u0075\u006e\u0064\u0020\u0069\u006e t\u0068\u0065\u0020\u0073\u0079\u0073\u0074e\u006d")
			}
		}
		for _, _abf := range _dbg.Keys() {
			_aeec.Set(_abf, _dbg.Get(_abf))
		}
		_efde := _dbg.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _efde != nil {
			if _, _gbeg = _aefa[_efde]; !_gbeg {
				_dfg.Objects = append(_dfg.Objects, _efde)
				_aefa[_efde] = struct{}{}
			}
		}
		_fff[_fda] = struct{}{}
		_ece := _aeec.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
		if _ece != nil {
			_dfg.Objects = append(_dfg.Objects, _ece)
			_aefa[_ece] = struct{}{}
		}
	}
	return nil
}
func _bbcc(_ced *_bag.Document) error {
	for _, _fbbfa := range _ced.Objects {
		_acee, _eaf := _geb.GetDict(_fbbfa)
		if !_eaf {
			continue
		}
		_gbgaf := _acee.Get("\u0054\u0079\u0070\u0065")
		if _gbgaf == nil {
			continue
		}
		if _dgf, _fafe := _geb.GetName(_gbgaf); _fafe && _dgf.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_ddfc, _cef := _geb.GetBool(_acee.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _cef && bool(*_ddfc) {
			_acee.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _geb.MakeBool(false))
		}
		if _acee.Get("\u0058\u0046\u0041") != nil {
			_acee.Remove("\u0058\u0046\u0041")
		}
	}
	_ccac, _fbabb := _ced.FindCatalog()
	if !_fbabb {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _ccac.Object.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_ccac.Object.Remove("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067")
	}
	return nil
}
func _ecfed(_fbbc *_ae.CompliancePdfReader) (_ebde []ViolatedRule) {
	_fbae := _fbbc.GetObjectNums()
	for _, _gegec := range _fbae {
		_eddbc, _beefa := _fbbc.GetIndirectObjectByNumber(_gegec)
		if _beefa != nil {
			continue
		}
		_cfdfe, _dedg := _geb.GetDict(_eddbc)
		if !_dedg {
			continue
		}
		_cfbf, _dedg := _geb.GetName(_cfdfe.Get("\u0054\u0079\u0070\u0065"))
		if !_dedg {
			continue
		}
		if _cfbf.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		_beegcc, _beefa := _ae.NewPdfFilespecFromObj(_cfdfe)
		if _beefa != nil {
			continue
		}
		if _beegcc.EF != nil {
			if _beegcc.F == nil || _beegcc.UF == nil {
				_ebde = append(_ebde, _eef("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
				break
			}
			if _beegcc.AFRelationship == nil {
				_ebde = append(_ebde, _eef("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
				break
			}
		}
	}
	return _ebde
}

var _ Profile = (*Profile3B)(nil)

func _fcfc(_cddce *_ae.PdfFont, _efdf *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_dgbb = "\u0036.\u0033\u002e\u0037\u002d\u0031"
		_gfge = "\u0041\u006cl \u006e\u006f\u006e\u002d\u0073\u0079\u006db\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0066o\u006e\u0074s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020e\u0069\u0074h\u0065\u0072\u0020\u004d\u0061\u0063\u0052\u006f\u006d\u0061\u006e\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041\u006e\u0073i\u0045n\u0063\u006f\u0064\u0069n\u0067\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072\u0020t\u0068\u0065 \u0045n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020t\u0068e\u0020\u0046o\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0072\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072 \u0074\u0068\u0065\u0020\u0042\u0061\u0073\u0065\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0064i\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0069s\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0074\u0068e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006be\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046\u006f\u006e\u0074 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u0049\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e, \u006eo\u0020n\u006f\u006e\u002d\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079p\u0065 \u0066\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0020\u0061\u0020\u0044\u0069\u0066\u0066e\u0072\u0065\u006e\u0063\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0075n\u006c\u0065s\u0073\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0074h\u0065\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u006e\u0061\u006d\u0065\u0073 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006ec\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0061\u0072\u0065\u0020\u006c\u0069\u0073\u0074\u0065\u0064 \u0069\u006e \u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065 G\u006c\u0079\u0070\u0068\u0020\u004c\u0069\u0073t\u0020\u0061\u006e\u0064\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066o\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072a\u006d\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073t\u0020\u0074\u0068\u0065\u0020\u004d\u0069\u0063\u0072o\u0073o\u0066\u0074\u0020\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0028\u0033\u002c\u0031 \u2013 P\u006c\u0061\u0074\u0066\u006f\u0072\u006d\u0020I\u0044\u003d\u0033\u002c\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067 I\u0044\u003d\u0031\u0029\u0020\u0065\u006e\u0063\u006f\u0064i\u006e\u0067 \u0069\u006e\u0020t\u0068\u0065\u0020'\u0063\u006d\u0061\u0070\u0027\u0020\u0074\u0061\u0062\u006c\u0065\u002e"
	)
	var _agbg string
	if _acfd, _fcfeg := _geb.GetName(_efdf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fcfeg {
		_agbg = _acfd.String()
	}
	if _agbg != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _fc
	}
	_bccc := _cddce.FontDescriptor()
	_bfde, _fbca := _geb.GetIntVal(_bccc.Flags)
	if !_fbca {
		_gd.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _eef(_dgbb, _gfge)
	}
	_eddc := (uint32(_bfde) >> 3) != 0
	if _eddc {
		return _fc
	}
	_gebc, _fbca := _geb.GetName(_efdf.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if !_fbca {
		return _eef(_dgbb, _gfge)
	}
	switch _gebc.String() {
	case "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067":
		return _fc
	default:
		return _eef(_dgbb, _gfge)
	}
}
func _fcga(_ccce *_ae.CompliancePdfReader) ViolatedRule { return _fc }
func _ebcg(_ccdgg *_ae.CompliancePdfReader) (_cfad []ViolatedRule) {
	_eedb := _ccdgg.GetObjectNums()
	for _, _fdfga := range _eedb {
		_dcbg, _ebaca := _ccdgg.GetIndirectObjectByNumber(_fdfga)
		if _ebaca != nil {
			continue
		}
		_egcd, _dbcb := _geb.GetDict(_dcbg)
		if !_dbcb {
			continue
		}
		_eegc, _dbcb := _geb.GetName(_egcd.Get("\u0054\u0079\u0070\u0065"))
		if !_dbcb {
			continue
		}
		if _eegc.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _egcd.Get("\u0045\u0046") != nil {
			_cfad = append(_cfad, _eef("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0031", "\u0041 \u0066\u0069\u006c\u0065 \u0073p\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046 \u0033\u002e\u0031\u0030\u002e\u0032\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0045\u0046 \u006be\u0079\u002e"))
			break
		}
	}
	_fdbe, _fceag := _debed(_ccdgg)
	if !_fceag {
		return _cfad
	}
	_cad, _fceag := _geb.GetDict(_fdbe.Get("\u004e\u0061\u006de\u0073"))
	if !_fceag {
		return _cfad
	}
	if _cad.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073") != nil {
		_cfad = append(_cfad, _eef("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0032", "\u0041\u0020\u0066i\u006c\u0065\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020d\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020PD\u0046 \u0052\u0065\u0066er\u0065\u006e\u0063\u0065\u0020\u0033\u002e6\u002e\u0033\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0045m\u0062\u0065\u0064\u0064\u0065\u0064\u0046i\u006c\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
	}
	return _cfad
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-2 standard.
// Implements model.StandardApplier.
func (_gddb *profile2) ApplyStandard(document *_bag.Document) (_dac error) {
	_fdc(document, 7)
	if _dac = _dfad(document, _gddb._feca.Now); _dac != nil {
		return _dac
	}
	if _dac = _cbbd(document); _dac != nil {
		return _dac
	}
	_gcgg, _cdcf := _fcea(_gddb._feca.CMYKDefaultColorSpace, _gddb._ceceg)
	_dac = _bffd(document, []pageColorspaceOptimizeFunc{_gcgg}, []documentColorspaceOptimizeFunc{_cdcf})
	if _dac != nil {
		return _dac
	}
	_dgdb(document)
	if _dac = _cagg(document); _dac != nil {
		return _dac
	}
	if _dac = _gfcc(document, _gddb._ceceg._fgb); _dac != nil {
		return _dac
	}
	if _dac = _dcbd(document); _dac != nil {
		return _dac
	}
	if _dac = _dabd(document); _dac != nil {
		return _dac
	}
	if _dac = _gfb(document); _dac != nil {
		return _dac
	}
	if _dac = _bbcc(document); _dac != nil {
		return _dac
	}
	if _gddb._ceceg._ea == "\u0041" {
		_fagcd(document)
	}
	if _dac = _bdf(document, _gddb._ceceg._fgb); _dac != nil {
		return _dac
	}
	if _dac = _ce(document); _dac != nil {
		return _dac
	}
	if _gbaa := _abc(document, _gddb._ceceg, _gddb._feca.Xmp); _gbaa != nil {
		return _gbaa
	}
	if _gddb._ceceg == _cba() {
		if _dac = _gfef(document); _dac != nil {
			return _dac
		}
	}
	if _dac = _gefg(document); _dac != nil {
		return _dac
	}
	if _dac = _beac(document); _dac != nil {
		return _dac
	}
	if _dac = _afbbf(document); _dac != nil {
		return _dac
	}
	return nil
}
func _acfcb(_bagfg *_ae.PdfInfo, _egeec *_eg.Document) bool {
	_aaca, _dcacg := _egeec.GetPdfInfo()
	if !_dcacg {
		return false
	}
	if _aaca.InfoDict == nil {
		return false
	}
	_cgfe, _dgbgg := _ae.NewPdfInfoFromObject(_aaca.InfoDict)
	if _dgbgg != nil {
		return false
	}
	if _bagfg.Creator != nil {
		if _cgfe.Creator == nil || _cgfe.Creator.String() != _bagfg.Creator.String() {
			return false
		}
	}
	if _bagfg.CreationDate != nil {
		if _cgfe.CreationDate == nil || !_cgfe.CreationDate.ToGoTime().Equal(_bagfg.CreationDate.ToGoTime()) {
			return false
		}
	}
	if _bagfg.ModifiedDate != nil {
		if _cgfe.ModifiedDate == nil || !_cgfe.ModifiedDate.ToGoTime().Equal(_bagfg.ModifiedDate.ToGoTime()) {
			return false
		}
	}
	if _bagfg.Producer != nil {
		if _cgfe.Producer == nil || _cgfe.Producer.String() != _bagfg.Producer.String() {
			return false
		}
	}
	if _bagfg.Keywords != nil {
		if _cgfe.Keywords == nil || _cgfe.Keywords.String() != _bagfg.Keywords.String() {
			return false
		}
	}
	if _bagfg.Trapped != nil {
		if _cgfe.Trapped == nil {
			return false
		}
		switch _bagfg.Trapped.String() {
		case "\u0054\u0072\u0075\u0065":
			if _cgfe.Trapped.String() != "\u0054\u0072\u0075\u0065" {
				return false
			}
		case "\u0046\u0061\u006cs\u0065":
			if _cgfe.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		default:
			if _cgfe.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		}
	}
	if _bagfg.Title != nil {
		if _cgfe.Title == nil || _cgfe.Title.String() != _bagfg.Title.String() {
			return false
		}
	}
	if _bagfg.Subject != nil {
		if _cgfe.Subject == nil || _cgfe.Subject.String() != _bagfg.Subject.String() {
			return false
		}
	}
	return true
}
func _fagcd(_dcac *_bag.Document) {
	_gbag, _egcb := _dcac.FindCatalog()
	if !_egcb {
		return
	}
	_fdbd, _egcb := _gbag.GetMarkInfo()
	if !_egcb {
		_fdbd = _geb.MakeDict()
	}
	_agbd, _egcb := _geb.GetBool(_fdbd.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
	if !_egcb || !bool(*_agbd) {
		_fdbd.Set("\u004d\u0061\u0072\u006b\u0065\u0064", _geb.MakeBool(true))
		_gbag.SetMarkInfo(_fdbd)
	}
}
func _ebgbd(_bgdb *_ae.CompliancePdfReader) (_bbbg []ViolatedRule) {
	var _gdab, _abbd bool
	_gdee := func() bool { return _gdab && _abbd }
	for _, _dgcd := range _bgdb.GetObjectNums() {
		_daad, _fbfcb := _bgdb.GetIndirectObjectByNumber(_dgcd)
		if _fbfcb != nil {
			_gd.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _dgcd, _fbfcb)
			continue
		}
		_acgbf, _cefe := _geb.GetDict(_daad)
		if !_cefe {
			continue
		}
		_fdgce, _cefe := _geb.GetName(_acgbf.Get("\u0054\u0079\u0070\u0065"))
		if !_cefe {
			continue
		}
		if *_fdgce != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_cfgd, _cefe := _geb.GetName(_acgbf.Get("\u0053"))
		if !_cefe {
			if !_gdab {
				_bbbg = append(_bbbg, _eef("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_gdab = true
				if _gdee() {
					return _bbbg
				}
			}
			continue
		}
		switch _ae.PdfActionType(*_cfgd) {
		case _ae.ActionTypeLaunch, _ae.ActionTypeSound, _ae.ActionTypeMovie, _ae.ActionTypeResetForm, _ae.ActionTypeImportData, _ae.ActionTypeJavaScript:
			if !_gdab {
				_bbbg = append(_bbbg, _eef("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_gdab = true
				if _gdee() {
					return _bbbg
				}
			}
			continue
		case _ae.ActionTypeNamed:
			if !_abbd {
				_eafde, _acefe := _geb.GetName(_acgbf.Get("\u004e"))
				if !_acefe {
					_bbbg = append(_bbbg, _eef("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_abbd = true
					if _gdee() {
						return _bbbg
					}
					continue
				}
				switch *_eafde {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_bbbg = append(_bbbg, _eef("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_abbd = true
					if _gdee() {
						return _bbbg
					}
					continue
				}
			}
		}
	}
	return _bbbg
}
func _gbcb(_edfga *_ae.CompliancePdfReader) (_cgdd ViolatedRule) {
	for _, _gcdg := range _edfga.GetObjectNums() {
		_defe, _bebcf := _edfga.GetIndirectObjectByNumber(_gcdg)
		if _bebcf != nil {
			continue
		}
		_ccfc, _gfgf := _geb.GetStream(_defe)
		if !_gfgf {
			continue
		}
		_dgce, _gfgf := _geb.GetName(_ccfc.Get("\u0054\u0079\u0070\u0065"))
		if !_gfgf {
			continue
		}
		if *_dgce != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_acef, _gfgf := _geb.GetName(_ccfc.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_gfgf {
			continue
		}
		if *_acef == "\u0050\u0053" {
			return _eef("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _ccfc.Get("\u0050\u0053") != nil {
			return _eef("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _cgdd
}
func _gcbe(_dddb *_ae.CompliancePdfReader) (_fegd []ViolatedRule) {
	var _fbcbd, _fcae, _gecc, _feffb, _fddbe, _ggcda, _cbea bool
	_abfe := func() bool { return _fbcbd && _fcae && _gecc && _feffb && _fddbe && _ggcda && _cbea }
	for _, _cffc := range _dddb.PageList {
		_dgcc, _geffc := _cffc.GetAnnotations()
		if _geffc != nil {
			_gd.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _geffc)
			continue
		}
		for _, _fbebg := range _dgcc {
			if !_fbcbd {
				switch _fbebg.GetContext().(type) {
				case *_ae.PdfAnnotationScreen, *_ae.PdfAnnotation3D, *_ae.PdfAnnotationSound, *_ae.PdfAnnotationMovie, nil:
					_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0031\u002d\u0031", "\u0041nn\u006f\u0074\u0061\u0074i\u006f\u006e t\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065r\u006d\u0069t\u0074\u0065\u0064\u002e\u0020\u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0033\u0044\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u0053\u0063\u0072\u0065\u0065\u006e\u0020\u0061n\u0064\u0020\u004d\u006f\u0076\u0069\u0065\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fbcbd = true
					if _abfe() {
						return _fegd
					}
				}
			}
			_egfbd, _ggab := _geb.GetDict(_fbebg.GetContainingPdfObject())
			if !_ggab {
				continue
			}
			_, _afeaf := _fbebg.GetContext().(*_ae.PdfAnnotationPopup)
			if !_afeaf && !_fcae {
				_, _cfdcd := _geb.GetIntVal(_egfbd.Get("\u0046"))
				if !_cfdcd {
					_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0045\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072i\u0065\u0073\u0020\u0077\u0068\u006fs\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0069\u0073\u0020\u0050\u006f\u0070u\u0070\u002c\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065y."))
					_fcae = true
					if _abfe() {
						return _fegd
					}
				}
			}
			if !_gecc {
				_bfccd, _cabe := _geb.GetIntVal(_egfbd.Get("\u0046"))
				if _cabe && !(_bfccd&4 == 4 && _bfccd&1 == 0 && _bfccd&2 == 0 && _bfccd&32 == 0 && _bfccd&256 == 0) {
					_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0032\u002d\u0032", "I\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073\u0068\u0061l\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020\u0049\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u002c\u0020\u0054\u006f\u0067\u0067\u006c\u0065\u004e\u006f\u0056\u0069\u0065\u0077\u002c\u0020\u0061\u006e\u0064 \u004eo\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020s\u0065\u0074\u0020t\u006f\u0020\u0030."))
					_gecc = true
					if _abfe() {
						return _fegd
					}
				}
			}
			_, _face := _fbebg.GetContext().(*_ae.PdfAnnotationText)
			if _face && !_feffb {
				_eeec, _edfe := _geb.GetIntVal(_egfbd.Get("\u0046"))
				if _edfe && !(_eeec&8 == 8 && _eeec&16 == 16) {
					_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0032\u002d\u0033", "\u0054\u0065\u0078\u0074\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069o\u006e\u0020\u0068\u0061\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006ca\u0067\u0073\u0020\u004e\u006f\u005a\u006f\u006f\u006d\u0020\u006f\u0072\u0020\u004e\u006f\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u0073\u0065t\u0020\u0074\u006f\u0020\u0030\u002e"))
					_feffb = true
					if _abfe() {
						return _fegd
					}
				}
			}
			if !_fddbe {
				_dbdad, _acbd := _geb.GetDict(_egfbd.Get("\u0041\u0050"))
				if _acbd {
					_bcfaa := _dbdad.Get("\u004e")
					if _bcfaa == nil || len(_dbdad.Keys()) > 1 {
						_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_fddbe = true
						if _abfe() {
							return _fegd
						}
						continue
					}
					_, _faea := _fbebg.GetContext().(*_ae.PdfAnnotationWidget)
					if _faea {
						_beaga, _ddbgc := _geb.GetName(_egfbd.Get("\u0046\u0054"))
						if _ddbgc && *_beaga == "\u0042\u0074\u006e" {
							if _, _cbdgb := _geb.GetDict(_bcfaa); !_cbdgb {
								_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_fddbe = true
								if _abfe() {
									return _fegd
								}
								continue
							}
						}
					}
					_, _beded := _geb.GetStream(_bcfaa)
					if !_beded {
						_fegd = append(_fegd, _eef("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_fddbe = true
						if _abfe() {
							return _fegd
						}
						continue
					}
				}
			}
			_ffgc, _bgbc := _fbebg.GetContext().(*_ae.PdfAnnotationWidget)
			if !_bgbc {
				continue
			}
			if !_ggcda {
				if _ffgc.A != nil {
					_fegd = append(_fegd, _eef("\u0036.\u0034\u002e\u0031\u002d\u0031", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_ggcda = true
					if _abfe() {
						return _fegd
					}
				}
			}
			if !_cbea {
				if _ffgc.AA != nil {
					_fegd = append(_fegd, _eef("\u0036.\u0034\u002e\u0031\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_cbea = true
					if _abfe() {
						return _fegd
					}
				}
			}
		}
	}
	return _fegd
}
func _gdagde(_fafdb *_ae.CompliancePdfReader) (_fcef ViolatedRule) {
	_gefbe, _bbce := _debed(_fafdb)
	if !_bbce {
		return _fc
	}
	if _gefbe.Get("\u0041\u0041") != nil {
		return _eef("\u0036.\u0036\u002e\u0032\u002d\u0033", "\u0054\u0068e\u0020\u0064\u006f\u0063\u0075\u006d\u0065n\u0074 \u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065n\u0074r\u0079 \u0066\u006f\u0072 \u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063\u0074i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
	}
	return _fc
}
func _gfdf(_adce *_ae.PdfFont, _fcdaa *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_fefc = "\u0036.\u0033\u002e\u0035\u002d\u0032"
		_feda = "\u0046\u006f\u0072\u0020\u0061l\u006c\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074 \u0073\u0075bs\u0065\u0074\u0073 \u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074he\u0020f\u006f\u006e\u0074\u0020\u0064\u0065s\u0063r\u0069\u0070\u0074o\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006ec\u006c\u0075\u0064e\u0020\u0061\u0020\u0043\u0068\u0061\u0072\u0053\u0065\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u006c\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072 \u006e\u0061\u006d\u0065\u0073\u0020d\u0065\u0066i\u006e\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020f\u006f\u006e\u0074\u0020s\u0075\u0062\u0073\u0065\u0074, \u0061\u0073 \u0064\u0065s\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e \u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0061\u0062\u006ce\u0020\u0035\u002e1\u0038\u002e"
	)
	var _dgbc string
	if _bacbg, _feag := _geb.GetName(_fcdaa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _feag {
		_dgbc = _bacbg.String()
	}
	if _dgbc != "\u0054\u0079\u0070e\u0031" {
		return _fc
	}
	if _dbe.IsStdFont(_dbe.StdFontName(_adce.BaseFont())) {
		return _fc
	}
	_cafge := _adce.FontDescriptor()
	if _cafge.CharSet == nil {
		return _eef(_fefc, _feda)
	}
	return _fc
}
func _fcab(_bbb *_ae.CompliancePdfReader) (_eade []ViolatedRule) {
	_gfdc, _cbbb := _debed(_bbb)
	if !_cbbb {
		return _eade
	}
	_fdaa := _eef("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e")
	_cedd, _cbbb := _geb.GetArray(_gfdc.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_cbbb {
		_eade = append(_eade, _fdaa)
		return _eade
	}
	_fgaaa := _eef("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e")
	if _cedd.Len() > 1 {
		_feaa := map[*_geb.PdfObjectDictionary]struct{}{}
		for _bgae := 0; _bgae < _cedd.Len(); _bgae++ {
			_dbbc, _cadd := _geb.GetDict(_cedd.Get(_bgae))
			if !_cadd {
				_eade = append(_eade, _fdaa)
				return _eade
			}
			if _bgae == 0 {
				_feaa[_dbbc] = struct{}{}
				continue
			}
			if _, _fggec := _feaa[_dbbc]; !_fggec {
				_eade = append(_eade, _fgaaa)
				break
			}
		}
	} else if _cedd.Len() == 0 {
		_eade = append(_eade, _fdaa)
		return _eade
	}
	_edda, _cbbb := _geb.GetDict(_cedd.Get(0))
	if !_cbbb {
		_eade = append(_eade, _fdaa)
		return _eade
	}
	if _ebed, _gabf := _geb.GetName(_edda.Get("\u0053")); !_gabf || (*_ebed) != "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411" {
		_eade = append(_eade, _fdaa)
		return _eade
	}
	_cebc, _abdda := _ae.NewPdfOutputIntentFromPdfObject(_edda)
	if _abdda != nil {
		_gd.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020i\u006et\u0065\u006e\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _abdda)
		return _eade
	}
	_ageb, _abdda := _bae.ParseHeader(_cebc.DestOutputProfile)
	if _abdda != nil {
		_gd.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066\u0061i\u006c\u0065d\u003a\u0020\u0025\u0076", _abdda)
		return _eade
	}
	if (_ageb.DeviceClass == _bae.DeviceClassPRTR || _ageb.DeviceClass == _bae.DeviceClassMNTR) && (_ageb.ColorSpace == _bae.ColorSpaceRGB || _ageb.ColorSpace == _bae.ColorSpaceCMYK || _ageb.ColorSpace == _bae.ColorSpaceGRAY) {
		return _eade
	}
	_eade = append(_eade, _fdaa)
	return _eade
}
func _afeeca(_abaf *_ae.PdfFont, _eabd *_geb.PdfObjectDictionary, _dgdga bool) ViolatedRule {
	const (
		_aadb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0034\u002d\u0031"
		_baac = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006ds\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020f\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0072e\u006e\u0064\u0065\u0072\u0069\u006eg\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020w\u0069t\u0068\u0069\u006e\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u002c \u0061\u0073\u0020\u0064\u0065\u0066\u0069n\u0065\u0064 \u0069\u006e\u0020\u0049S\u004f\u0020\u0033\u0032\u00300\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0039\u002e"
	)
	if _dgdga {
		return _fc
	}
	_ccca := _abaf.FontDescriptor()
	var _fcec string
	if _gbffg, _fcfad := _geb.GetName(_eabd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fcfad {
		_fcec = _gbffg.String()
	}
	switch _fcec {
	case "\u0054\u0079\u0070e\u0031":
		if _ccca.FontFile == nil {
			return _eef(_aadb, _baac)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _ccca.FontFile2 == nil {
			return _eef(_aadb, _baac)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _ccca.FontFile3 == nil {
			return _eef(_aadb, _baac)
		}
	}
	return _fc
}
func _afbbf(_bbgg *_bag.Document) error {
	_geec, _bcb := _bbgg.FindCatalog()
	if !_bcb {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _geec.Object.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		_geec.Object.Remove("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073")
	}
	return nil
}
func _fbgb(_eadg *_ae.CompliancePdfReader) ViolatedRule { return _fc }

var _ Profile = (*Profile2U)(nil)

func _bdfc(_eafd *_ae.CompliancePdfReader) (_cbad []ViolatedRule) {
	var _cdgg, _debe, _ccabd, _gafb, _cgga, _gaecc bool
	_fbgfe := map[*_geb.PdfObjectStream]struct{}{}
	for _, _cbcb := range _eafd.GetObjectNums() {
		if _cdgg && _debe && _cgga && _ccabd && _gafb && _gaecc {
			return _cbad
		}
		_aagfe, _gebg := _eafd.GetIndirectObjectByNumber(_cbcb)
		if _gebg != nil {
			continue
		}
		_adf, _gdddf := _geb.GetStream(_aagfe)
		if !_gdddf {
			continue
		}
		if _, _gdddf = _fbgfe[_adf]; _gdddf {
			continue
		}
		_fbgfe[_adf] = struct{}{}
		_cbbdb, _gdddf := _geb.GetName(_adf.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_gdddf {
			continue
		}
		if !_gafb {
			if _adf.Get("\u0052\u0065\u0066") != nil {
				_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_gafb = true
			}
		}
		if _cbbdb.String() == "\u0050\u0053" {
			if !_gaecc {
				_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_gaecc = true
				continue
			}
		}
		if _cbbdb.String() == "\u0046\u006f\u0072\u006d" {
			if _debe && _ccabd && _gafb {
				continue
			}
			if !_debe && _adf.Get("\u004f\u0050\u0049") != nil {
				_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0034\u002d\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0028\u0049\u006d\u0061\u0067\u0065\u0020\u006f\u0072\u0020\u0046\u006f\u0072\u006d\u0029\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u004fP\u0049\u0020\u006b\u0065\u0079\u002e"))
				_debe = true
			}
			if !_ccabd {
				if _adf.Get("\u0050\u0053") != nil {
					_ccabd = true
				}
				if _fgbfe := _adf.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _fgbfe != nil && !_ccabd {
					if _geeeg, _gedg := _geb.GetName(_fgbfe); _gedg && *_geeeg == "\u0050\u0053" {
						_ccabd = true
					}
				}
				if _ccabd {
					_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _cbbdb.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_cdgg && _adf.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0034\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_cdgg = true
		}
		if !_cgga && _adf.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_gcd, _abed := _geb.GetBool(_adf.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _abed && bool(*_gcd) {
				continue
			}
			_cbad = append(_cbad, _eef("\u0036.\u0032\u002e\u0034\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_cgga = true
		}
	}
	return _cbad
}
func _fedf(_eegg *_ae.CompliancePdfReader) ViolatedRule {
	_aecb := _eegg.ParserMetadata().HeaderCommentBytes()
	if _aecb[0] > 127 && _aecb[1] > 127 && _aecb[2] > 127 && _aecb[3] > 127 {
		return _fc
	}
	return _eef("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}
func _abeg(_ebac *_ae.CompliancePdfReader) (_efag []ViolatedRule) {
	if _ebac.ParserMetadata().HasOddLengthHexStrings() {
		_efag = append(_efag, _eef("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _ebac.ParserMetadata().HasOddLengthHexStrings() {
		_efag = append(_efag, _eef("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _efag
}
func _dbbe(_eed *_ae.CompliancePdfReader) ViolatedRule {
	for _, _ccdg := range _eed.PageList {
		_cgcc := _ccdg.GetContentStreamObjs()
		for _, _fgbf := range _cgcc {
			_fgbf = _geb.TraceToDirectObject(_fgbf)
			var _debc string
			switch _agcc := _fgbf.(type) {
			case *_geb.PdfObjectString:
				_debc = _agcc.Str()
			case *_geb.PdfObjectStream:
				_fdec, _cegc := _geb.GetName(_geb.TraceToDirectObject(_agcc.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
				if _cegc {
					if *_fdec == _geb.StreamEncodingFilterNameLZW {
						return _eef("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
					}
				}
				_dcec, _agaca := _geb.DecodeStream(_agcc)
				if _agaca != nil {
					_gd.Log.Debug("\u0045r\u0072\u003a\u0020\u0025\u0076", _agaca)
					continue
				}
				_debc = string(_dcec)
			default:
				_gd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063t\u003a\u0020\u0025\u0054", _fgbf)
				continue
			}
			_bdcb := _gg.NewContentStreamParser(_debc)
			_eged, _ddda := _bdcb.Parse()
			if _ddda != nil {
				_gd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _ddda)
				continue
			}
			for _, _bdbf := range *_eged {
				if !(_bdbf.Operand == "\u0042\u0049" && len(_bdbf.Params) == 1) {
					continue
				}
				_fgbc, _ddeb := _bdbf.Params[0].(*_gg.ContentStreamInlineImage)
				if !_ddeb {
					continue
				}
				_fdfdb, _bace := _fgbc.GetEncoder()
				if _bace != nil {
					_gd.Log.Debug("\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006ee\u0020\u0069\u006d\u0061\u0067\u0065 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bace)
					continue
				}
				if _fdfdb.GetFilterName() == _geb.StreamEncodingFilterNameLZW {
					return _eef("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
				}
			}
		}
	}
	return _fc
}
func _gabb(_adee *_ae.CompliancePdfReader) ViolatedRule { return _fc }

var _fc = ViolatedRule{}

// ApplyStandard tries to change the content of the writer to match the PDF/A-3 standard.
// Implements model.StandardApplier.
func (_dada *profile3) ApplyStandard(document *_bag.Document) (_fed error) {
	_fdc(document, 7)
	if _fed = _dfad(document, _dada._ebg.Now); _fed != nil {
		return _fed
	}
	if _fed = _cbbd(document); _fed != nil {
		return _fed
	}
	_ggad, _dfgcg := _fcea(_dada._ebg.CMYKDefaultColorSpace, _dada._egbf)
	_fed = _bffd(document, []pageColorspaceOptimizeFunc{_ggad}, []documentColorspaceOptimizeFunc{_dfgcg})
	if _fed != nil {
		return _fed
	}
	_dgdb(document)
	if _fed = _cagg(document); _fed != nil {
		return _fed
	}
	if _fed = _gfcc(document, _dada._egbf._fgb); _fed != nil {
		return _fed
	}
	if _fed = _dcbd(document); _fed != nil {
		return _fed
	}
	if _fed = _dabd(document); _fed != nil {
		return _fed
	}
	if _fed = _gfb(document); _fed != nil {
		return _fed
	}
	if _fed = _bbcc(document); _fed != nil {
		return _fed
	}
	if _dada._egbf._ea == "\u0041" {
		_fagcd(document)
	}
	if _fed = _bdf(document, _dada._egbf._fgb); _fed != nil {
		return _fed
	}
	if _fed = _ce(document); _fed != nil {
		return _fed
	}
	if _deeb := _abc(document, _dada._egbf, _dada._ebg.Xmp); _deeb != nil {
		return _deeb
	}
	if _dada._egbf == _ec() {
		if _fed = _gfef(document); _fed != nil {
			return _fed
		}
	}
	if _fed = _gefg(document); _fed != nil {
		return _fed
	}
	if _fed = _beac(document); _fed != nil {
		return _fed
	}
	if _fed = _afbbf(document); _fed != nil {
		return _fed
	}
	return nil
}
func _cfec(_gdag *_ae.CompliancePdfReader) []ViolatedRule { return nil }
func _ebcb(_gece *_ae.PdfInfo, _geea func() _ff.Time) error {
	var _dbgg *_ae.PdfDate
	if _gece.CreationDate == nil {
		_ecda, _gbf := _ae.NewPdfDateFromTime(_geea())
		if _gbf != nil {
			return _gbf
		}
		_dbgg = &_ecda
		_gece.CreationDate = _dbgg
	}
	if _gece.ModifiedDate == nil {
		if _dbgg != nil {
			_aeaf, _cbef := _ae.NewPdfDateFromTime(_geea())
			if _cbef != nil {
				return _cbef
			}
			_dbgg = &_aeaf
		}
		_gece.ModifiedDate = _dbgg
	}
	return nil
}

type standardType struct {
	_fgb int
	_ea  string
}

func _befc(_beba *_geb.PdfObjectDictionary, _bebee map[*_geb.PdfObjectStream][]byte, _fdef map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_bfae = "\u0036.\u0033\u002e\u0033\u002d\u0034"
		_acff = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _dcea string
	if _gage, _feeaf := _geb.GetName(_beba.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _feeaf {
		_dcea = _gage.String()
	}
	if _dcea != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_fbdb := _beba.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ggfe := _geb.GetName(_fbdb); _ggfe {
		return _fc
	}
	_bdgd, _gbfb := _geb.GetStream(_fbdb)
	if !_gbfb {
		return _eef(_bfae, _acff)
	}
	_gbbb, _eedbg := _caddc(_bdgd, _bebee, _fdef)
	if _eedbg != nil {
		return _eef(_bfae, _acff)
	}
	_eggf, _fddfc := _geb.GetIntVal(_bdgd.Get("\u0057\u004d\u006fd\u0065"))
	_caaaa, _dfcg := _gbbb.WMode()
	if _fddfc && _dfcg {
		if _caaaa != _eggf {
			return _eef(_bfae, _acff)
		}
	}
	if (_fddfc && !_dfcg) || (!_fddfc && _dfcg) {
		return _eef(_bfae, _acff)
	}
	return _fc
}
func _cagg(_egba *_bag.Document) error {
	_age, _cdag := _egba.FindCatalog()
	if !_cdag {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_bcfgf, _cdag := _geb.GetDict(_age.Object.Get("\u0050\u0065\u0072m\u0073"))
	if _cdag {
		_fde := _geb.MakeDict()
		_bbee := _bcfgf.Keys()
		for _, _begb := range _bbee {
			if _begb.String() == "\u0055\u0052\u0033" || _begb.String() == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				_fde.Set(_begb, _bcfgf.Get(_begb))
			}
		}
		_age.Object.Set("\u0050\u0065\u0072m\u0073", _fde)
	}
	return nil
}

var _ Profile = (*Profile2B)(nil)

func _fgabb(_afcbb *_ae.CompliancePdfReader) []ViolatedRule         { return nil }
func _ebdd(_affbe *_ae.CompliancePdfReader) (_aaaab []ViolatedRule) { return _aaaab }
func (_adg *documentImages) hasOnlyDeviceCMYK() bool                { return _adg._cag && !_adg._gdf && !_adg._cfb }
func _gedf(_ccfa *_geb.PdfObjectDictionary, _egcba map[*_geb.PdfObjectStream][]byte, _decg map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_dfefb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0033"
		_bffde = "\u0041\u006c\u006c \u0043\u004d\u0061\u0070s\u0020\u0075\u0073ed\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074 th\u006f\u0073\u0065\u0020\u006ci\u0073\u0074\u0065\u0064\u0020i\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u00320\u00300\u002d1\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0035\u002e\u0032\u002c\u0020\u0054\u0061\u0062\u006c\u0065 \u0031\u00318,\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e \u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064e\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u00300\u0030-\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0035\u002e"
	)
	var _eadee string
	if _gaee, _dbef := _geb.GetName(_ccfa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dbef {
		_eadee = _gaee.String()
	}
	if _eadee != "\u0054\u0079\u0070e\u0030" {
		return _fc
	}
	_gadec := _ccfa.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _dcggc, _fcbdg := _geb.GetName(_gadec); _fcbdg {
		switch _dcggc.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _fc
		default:
			return _eef(_dfefb, _bffde)
		}
	}
	_dacc, _cegb := _geb.GetStream(_gadec)
	if !_cegb {
		return _eef(_dfefb, _bffde)
	}
	_, _bacbge := _caddc(_dacc, _egcba, _decg)
	if _bacbge != nil {
		return _eef(_dfefb, _bffde)
	}
	return _fc
}
func _bffd(_dadd *_bag.Document, _aea []pageColorspaceOptimizeFunc, _acab []documentColorspaceOptimizeFunc) error {
	_gbae, _gbefd := _dadd.GetPages()
	if !_gbefd {
		return nil
	}
	var _ffeg []*_bag.Image
	for _gdb, _fafb := range _gbae {
		_aeag, _dcdb := _fafb.FindXObjectImages()
		if _dcdb != nil {
			return _dcdb
		}
		for _, _ggcd := range _aea {
			if _dcdb = _ggcd(_dadd, &_gbae[_gdb], _aeag); _dcdb != nil {
				return _dcdb
			}
		}
		_ffeg = append(_ffeg, _aeag...)
	}
	for _, _eefg := range _acab {
		if _efdc := _eefg(_dadd, _ffeg); _efdc != nil {
			return _efdc
		}
	}
	return nil
}

type colorspaceModification struct {
	_bd   _gb.ColorConverter
	_fbeb _ae.PdfColorspace
}

func _fadc(_gbfe *_geb.PdfObjectDictionary, _gcbfbe map[*_geb.PdfObjectStream][]byte, _adcga map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_cegf = "\u0036.\u0033\u002e\u0038\u002d\u0031"
		_ccgd = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_fedb := _gbfe.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _ceeg, _agaa := _geb.GetName(_fedb); _agaa {
		if _ceeg.String() == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _ceeg.String() == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" || _ceeg.String() == "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067" || _ceeg.String() == "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067" || _ceeg.String() == "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067" {
			return _fc
		}
	}
	_cdef, _cfba := _geb.GetStream(_gbfe.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _cfba {
		_, _dbfa := _caddc(_cdef, _gcbfbe, _adcga)
		if _dbfa != nil {
			return _eef(_cegf, _ccgd)
		}
		return _fc
	}
	_fddgg, _cfba := _geb.GetName(_gbfe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_cfba {
		return _eef(_cegf, _ccgd)
	}
	switch _fddgg.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _fc
	}
	return _eef(_cegf, _ccgd)
}
func _aged(_ffff *Profile3Options) {
	if _ffff.Now == nil {
		_ffff.Now = _ff.Now
	}
}
func _ebee(_addf standardType, _cagc *_bag.OutputIntents) error {
	_acg, _gcfb := _bae.NewCmykIsoCoatedV2OutputIntent(_addf.outputIntentSubtype())
	if _gcfb != nil {
		return _gcfb
	}
	if _gcfb = _cagc.Add(_acg.ToPdfObject()); _gcfb != nil {
		return _gcfb
	}
	return nil
}
func _bef(_acdeb, _bgde, _fcfe, _edc string) (string, bool) {
	_dggc := _db.Index(_acdeb, _bgde)
	if _dggc == -1 {
		return "", false
	}
	_geeee := _db.Index(_acdeb, _fcfe)
	if _geeee == -1 {
		return "", false
	}
	if _geeee < _dggc {
		return "", false
	}
	return _acdeb[:_dggc] + _bgde + _edc + _acdeb[_geeee:], true
}
func _befb(_daag *_geb.PdfObjectDictionary, _gddgd map[*_geb.PdfObjectStream][]byte, _bbec map[*_geb.PdfObjectStream]*_gba.CMap) ViolatedRule {
	const (
		_ccbff = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0037\u002d\u0031"
		_fecbe = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_fbcdg, _aebc := _geb.GetStream(_daag.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _aebc {
		_, _ffefe := _caddc(_fbcdg, _gddgd, _bbec)
		if _ffefe != nil {
			return _eef(_ccbff, _fecbe)
		}
		return _fc
	}
	_afaf, _aebc := _geb.GetName(_daag.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_aebc {
		return _eef(_ccbff, _fecbe)
	}
	switch _afaf.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _fc
	}
	return _eef(_ccbff, _fecbe)
}
func _ggcc(_fdcbg *_ae.CompliancePdfReader) (_baag []ViolatedRule) {
	var (
		_cffg, _fbegf, _dfcbe, _eacbd, _dabe bool
		_eeed                                func(_geb.PdfObject)
	)
	_eeed = func(_gdcca _geb.PdfObject) {
		switch _ebbdc := _gdcca.(type) {
		case *_geb.PdfObjectInteger:
			if !_cffg && (int64(*_ebbdc) > _dg.MaxInt32 || int64(*_ebbdc) < -_dg.MaxInt32) {
				_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_cffg = true
			}
		case *_geb.PdfObjectFloat:
			if !_fbegf && (_dg.Abs(float64(*_ebbdc)) > _dg.MaxFloat32) {
				_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0032", "\u0041 \u0063\u006f\u006e\u0066orm\u0069\u006e\u0067\u0020f\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0061\u006c\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0075\u0074\u0073\u0069de\u0020\u0074\u0068e\u0020\u0072\u0061\u006e\u0067e\u0020o\u0066\u0020\u002b\u002f\u002d\u0033\u002e\u0034\u00303\u0020\u0078\u0020\u0031\u0030\u005e\u0033\u0038\u002e"))
			}
		case *_geb.PdfObjectString:
			if !_dfcbe && len([]byte(_ebbdc.Str())) > 32767 {
				_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0033\u0032\u0037\u0036\u0037\u002e"))
				_dfcbe = true
			}
		case *_geb.PdfObjectName:
			if !_eacbd && len([]byte(*_ebbdc)) > 127 {
				_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_eacbd = true
			}
		case *_geb.PdfObjectArray:
			for _, _egef := range _ebbdc.Elements() {
				_eeed(_egef)
			}
			if !_dabe && (_ebbdc.Len() == 4 || _ebbdc.Len() == 5) {
				_dbfg, _ccdc := _geb.GetName(_ebbdc.Get(0))
				if !_ccdc {
					return
				}
				if *_dbfg != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_cgaf := _ebbdc.Get(1)
				_cgaf = _geb.TraceToDirectObject(_cgaf)
				_ggdb, _ccdc := _geb.GetArray(_cgaf)
				if !_ccdc {
					return
				}
				if _ggdb.Len() > 32 {
					_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0039", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006e\u0075\u006db\u0065\u0072\u0020\u006ff\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0069\u0073\u0020\u0033\u0032\u002e"))
					_dabe = true
				}
			}
		case *_geb.PdfObjectDictionary:
			_dadbf := _ebbdc.Keys()
			for _geef, _fbdd := range _dadbf {
				_eeed(&_dadbf[_geef])
				_eeed(_ebbdc.Get(_fbdd))
			}
		case *_geb.PdfObjectStream:
			_eeed(_ebbdc.PdfObjectDictionary)
		case *_geb.PdfObjectStreams:
			for _, _cbfc := range _ebbdc.Elements() {
				_eeed(_cbfc)
			}
		case *_geb.PdfObjectReference:
			_eeed(_ebbdc.Resolve())
		}
	}
	_fccd := _fdcbg.GetObjectNums()
	if len(_fccd) > 8388607 {
		_baag = append(_baag, _eef("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _bbbe := range _fccd {
		_eafc, _deabc := _fdcbg.GetIndirectObjectByNumber(_bbbe)
		if _deabc != nil {
			continue
		}
		_dgedc := _geb.TraceToDirectObject(_eafc)
		_eeed(_dgedc)
	}
	return _baag
}

type profile3 struct {
	_egbf standardType
	_ebg  Profile3Options
}

func _cbbd(_bed *_bag.Document) error {
	_cfdc, _bdfd := _bed.FindCatalog()
	if !_bdfd {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_cfdc.SetVersion()
	return nil
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-3 standard.
func (_feab *profile3) ValidateStandard(r *_ae.CompliancePdfReader) error {
	_aeba := VerificationError{ConformanceLevel: _feab._egbf._fgb, ConformanceVariant: _feab._egbf._ea}
	if _fbbe := _caec(r); _fbbe != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _fbbe)
	}
	if _gedc := _fedf(r); _gedc != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gedc)
	}
	if _gdcd := _cgd(r); _gdcd != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gdcd)
	}
	if _ggf := _bdcg(r); _ggf != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _ggf)
	}
	if _dcgg := _fbgfd(r); _dcgg != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _dcgg)
	}
	if _gefcg := _abeg(r); len(_gefcg) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gefcg...)
	}
	if _bage := _fgaag(r); len(_bage) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _bage...)
	}
	if _ebb := _fgabb(r); len(_ebb) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _ebb...)
	}
	if _cgbd := _bbdbb(r); _cgbd != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _cgbd)
	}
	if _bcdd := _gdfegg(r); len(_bcdd) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _bcdd...)
	}
	if _egcf := _ggcc(r); len(_egcf) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _egcf...)
	}
	if _gff := _faab(r); _gff != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gff)
	}
	if _agee := _bdfda(r); len(_agee) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _agee...)
	}
	if _gadb := _abda(r); len(_gadb) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gadb...)
	}
	if _edg := _fbgb(r); _edg != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _edg)
	}
	if _fefd := _cega(r); len(_fefd) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _fefd...)
	}
	if _dadb := _cgad(r); len(_dadb) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _dadb...)
	}
	if _cce := _cgce(r); _cce != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _cce)
	}
	if _baeef := _bbff(r); len(_baeef) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _baeef...)
	}
	if _gaae := _bbed(r, _feab._egbf); len(_gaae) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gaae...)
	}
	if _gbbe := _gcbe(r); len(_gbbe) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gbbe...)
	}
	if _gffd := _bggcdg(r); len(_gffd) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _gffd...)
	}
	if _ebbb := _ebdd(r); len(_ebbb) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _ebbb...)
	}
	if _fdce := _bdcf(r); _fdce != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _fdce)
	}
	if _bedd := _efbe(r); len(_bedd) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _bedd...)
	}
	if _ddef := _dfcge(r); _ddef != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _ddef)
	}
	if _aded := _ddadb(r, _feab._egbf, false); len(_aded) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _aded...)
	}
	if _feab._egbf == _ec() {
		if _badf := _ggefb(r); len(_badf) != 0 {
			_aeba.ViolatedRules = append(_aeba.ViolatedRules, _badf...)
		}
	}
	if _dfeg := _ecfed(r); len(_dfeg) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _dfeg...)
	}
	if _dbga := _cebe(r); len(_dbga) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _dbga...)
	}
	if _bacb := _cdfeb(r); len(_bacb) != 0 {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _bacb...)
	}
	if _aeddg := _beef(r); _aeddg != _fc {
		_aeba.ViolatedRules = append(_aeba.ViolatedRules, _aeddg)
	}
	if len(_aeba.ViolatedRules) > 0 {
		_f.Slice(_aeba.ViolatedRules, func(_dabb, _dfdce int) bool {
			return _aeba.ViolatedRules[_dabb].RuleNo < _aeba.ViolatedRules[_dfdce].RuleNo
		})
		return _aeba
	}
	return nil
}
func _bggcdg(_accb *_ae.CompliancePdfReader) (_dcef []ViolatedRule) {
	for _, _eeccd := range _accb.GetObjectNums() {
		_ccfb, _daggf := _accb.GetIndirectObjectByNumber(_eeccd)
		if _daggf != nil {
			continue
		}
		_fdbfc, _cggcc := _geb.GetDict(_ccfb)
		if !_cggcc {
			continue
		}
		_fadbb, _cggcc := _geb.GetName(_fdbfc.Get("\u0054\u0079\u0070\u0065"))
		if !_cggcc {
			continue
		}
		if _fadbb.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_fefege, _cggcc := _geb.GetBool(_fdbfc.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _cggcc && bool(*_fefege) {
			_dcef = append(_dcef, _eef("\u0036.\u0034\u002e\u0031\u002d\u0033", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
		if _fdbfc.Get("\u0058\u0046\u0041") != nil {
			_dcef = append(_dcef, _eef("\u0036.\u0034\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020i\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020t\u0068\u0061\u0074\u0020f\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065 \u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d \u006b\u0065\u0079\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006f\u0066 \u0061 \u0050\u0044F\u002fA\u002d\u0032\u0020\u0066ile\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u0046\u0041\u0020\u006b\u0065y."))
		}
	}
	_bagaf, _ddaeb := _debed(_accb)
	if _ddaeb && _bagaf.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_dcef = append(_dcef, _eef("\u0036.\u0034\u002e\u0032\u002d\u0032", "\u0041\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006cog\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u002e"))
	}
	return _dcef
}
func _ddad(_daab *Profile2Options) {
	if _daab.Now == nil {
		_daab.Now = _ff.Now
	}
}
func _bdfda(_ebge *_ae.CompliancePdfReader) (_edfd []ViolatedRule) {
	var _dgagb, _edca, _dgfb, _cged bool
	_efdbf := func() bool { return _dgagb && _edca && _dgfb && _cged }
	_edaga, _ffgbf := _fdaf(_ebge)
	var _gadef _bae.ProfileHeader
	if _ffgbf {
		_gadef, _ = _bae.ParseHeader(_edaga.DestOutputProfile)
	}
	_dgfa := map[_geb.PdfObject]struct{}{}
	var _eagbc func(_eagba _ae.PdfColorspace) bool
	_eagbc = func(_gccc _ae.PdfColorspace) bool {
		switch _egdga := _gccc.(type) {
		case *_ae.PdfColorspaceDeviceGray:
			if !_dgagb {
				if !_ffgbf {
					_edfd = append(_edfd, _eef("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d4", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064 \u0069\u0066\u0020\u0061\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065p\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0047\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0075r \u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u0077\u0068\u0065n \u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072a\u0079\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u002c o\u0072\u0020\u0069\u0066\u0020\u0061\u0020\u0050\u0044\u0046\u002fA\u0020\u004f\u0075tp\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002e"))
					_dgagb = true
					if _efdbf() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceDeviceRGB:
			if !_edca {
				if !_ffgbf || _gadef.ColorSpace != _bae.ColorSpaceRGB {
					_edfd = append(_edfd, _eef("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d2", "\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042\u0020\u0073\u0068\u0061\u006cl\u0020\u006f\u006e\u006c\u0079\u0020\u0062e\u0020\u0075\u0073\u0065\u0064\u0020\u0069f\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063e\u0020\u0069n\u0064\u0065\u0070e\u006e\u0064\u0065\u006et \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0052\u0047\u0042\u0020\u0063\u006fl\u006f\u0075r\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020b\u0065\u0065\u006e\u0020s\u0065\u0074 \u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020c\u006flou\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020i\u0073\u0020\u0075\u0073\u0065\u0064\u002c\u0020\u006f\u0072\u0020if\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044F\u002f\u0041\u0020\u004fut\u0070\u0075\u0074\u0049\u006e\u0074\u0065n\u0074\u0020t\u0068\u0061t\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u006e\u0020\u0052\u0047\u0042\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061\u0074io\u006e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u002e"))
					_edca = true
					if _efdbf() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceDeviceCMYK:
			if !_dgfb {
				if !_ffgbf || _gadef.ColorSpace != _bae.ColorSpaceCMYK {
					_edfd = append(_edfd, _eef("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d3", "\u0044e\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020\u0073hal\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u0066\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065\u0070\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044ef\u0061\u0075\u006c\u0074\u0043\u004d\u0059K\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061s\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u006fr \u0069\u0066\u0020\u0061\u0020\u0044e\u0076\u0069\u0063\u0065\u004e\u002d\u0062\u0061\u0073\u0065\u0064\u0020\u0044\u0065f\u0061\u0075\u006c\u0074\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006f\u0075r\u0020\u0073\u0070\u0061\u0063e\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e\u0020\u0073\u0065\u0074\u0020\u0077\u0068\u0065\u006e\u0020\u0074h\u0065\u0020\u0044\u0065\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006fu\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u0072\u0020t\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u0020\u0043\u004d\u0059\u004b\u0020d\u0065\u0073\u0074\u0069\u006e\u0061t\u0069\u006f\u006e\u0020\u0070r\u006f\u0066\u0069\u006c\u0065\u002e"))
					_dgfb = true
					if _efdbf() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceICCBased:
			if !_cged {
				_ffda, _acbb := _bae.ParseHeader(_egdga.Data)
				if _acbb != nil {
					_gd.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _acbb)
					_edfd = append(_edfd, func() ViolatedRule {
						return _eef("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e")
					}())
					_cged = true
					if _efdbf() {
						return true
					}
				}
				if !_cged {
					var _bgdag, _agec bool
					switch _ffda.DeviceClass {
					case _bae.DeviceClassPRTR, _bae.DeviceClassMNTR, _bae.DeviceClassSCNR, _bae.DeviceClassSPAC:
					default:
						_bgdag = true
					}
					switch _ffda.ColorSpace {
					case _bae.ColorSpaceRGB, _bae.ColorSpaceCMYK, _bae.ColorSpaceGRAY, _bae.ColorSpaceLAB:
					default:
						_agec = true
					}
					if _bgdag || _agec {
						_edfd = append(_edfd, _eef("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e"))
						_cged = true
						if _efdbf() {
							return true
						}
					}
				}
			}
			if _egdga.Alternate != nil {
				return _eagbc(_egdga.Alternate)
			}
		}
		return false
	}
	for _, _gddc := range _ebge.GetObjectNums() {
		_ebd, _acgd := _ebge.GetIndirectObjectByNumber(_gddc)
		if _acgd != nil {
			continue
		}
		_fdefa, _dgaa := _geb.GetStream(_ebd)
		if !_dgaa {
			continue
		}
		_bcga, _dgaa := _geb.GetName(_fdefa.Get("\u0054\u0079\u0070\u0065"))
		if !_dgaa || _bcga.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_aecd, _dgaa := _geb.GetName(_fdefa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_dgaa {
			continue
		}
		_dgfa[_fdefa] = struct{}{}
		switch _aecd.String() {
		case "\u0049\u006d\u0061g\u0065":
			_eagd, _fegb := _ae.NewXObjectImageFromStream(_fdefa)
			if _fegb != nil {
				continue
			}
			_dgfa[_fdefa] = struct{}{}
			if _eagbc(_eagd.ColorSpace) {
				return _edfd
			}
		case "\u0046\u006f\u0072\u006d":
			_feec, _efdee := _geb.GetDict(_fdefa.Get("\u0047\u0072\u006fu\u0070"))
			if !_efdee {
				continue
			}
			_bcff := _feec.Get("\u0043\u0053")
			if _bcff == nil {
				continue
			}
			_fdbb, _febcc := _ae.NewPdfColorspaceFromPdfObject(_bcff)
			if _febcc != nil {
				continue
			}
			if _eagbc(_fdbb) {
				return _edfd
			}
		}
	}
	for _, _feggc := range _ebge.PageList {
		_ddce, _aege := _feggc.GetContentStreams()
		if _aege != nil {
			continue
		}
		for _, _ggge := range _ddce {
			_cbbdc, _abdg := _gg.NewContentStreamParser(_ggge).Parse()
			if _abdg != nil {
				continue
			}
			for _, _bgdaec := range *_cbbdc {
				if len(_bgdaec.Params) > 1 {
					continue
				}
				switch _bgdaec.Operand {
				case "\u0042\u0049":
					_bcgf, _debgc := _bgdaec.Params[0].(*_gg.ContentStreamInlineImage)
					if !_debgc {
						continue
					}
					_gega, _fbcd := _bcgf.GetColorSpace(_feggc.Resources)
					if _fbcd != nil {
						continue
					}
					if _eagbc(_gega) {
						return _edfd
					}
				case "\u0044\u006f":
					_bdbd, _baef := _geb.GetName(_bgdaec.Params[0])
					if !_baef {
						continue
					}
					_facd, _cfff := _feggc.Resources.GetXObjectByName(*_bdbd)
					if _, _ccaae := _dgfa[_facd]; _ccaae {
						continue
					}
					switch _cfff {
					case _ae.XObjectTypeImage:
						_fcbg, _gcdd := _ae.NewXObjectImageFromStream(_facd)
						if _gcdd != nil {
							continue
						}
						_dgfa[_facd] = struct{}{}
						if _eagbc(_fcbg.ColorSpace) {
							return _edfd
						}
					case _ae.XObjectTypeForm:
						_bgbbg, _dfef := _geb.GetDict(_facd.Get("\u0047\u0072\u006fu\u0070"))
						if !_dfef {
							continue
						}
						_abea, _dfef := _geb.GetName(_bgbbg.Get("\u0043\u0053"))
						if !_dfef {
							continue
						}
						_bcfc, _edeg := _ae.NewPdfColorspaceFromPdfObject(_abea)
						if _edeg != nil {
							continue
						}
						_dgfa[_facd] = struct{}{}
						if _eagbc(_bcfc) {
							return _edfd
						}
					}
				}
			}
		}
	}
	return _edfd
}
func _bbff(_gdeg *_ae.CompliancePdfReader) (_caaf []ViolatedRule) {
	var _daec, _ggebg, _ffef, _cfagd, _cefd, _cdff bool
	_eagg := func() bool { return _daec && _ggebg && _ffef && _cfagd && _cefd && _cdff }
	for _, _gfega := range _gdeg.PageList {
		if _gfega.Resources == nil {
			continue
		}
		_dfdb, _cfbee := _geb.GetDict(_gfega.Resources.Font)
		if !_cfbee {
			continue
		}
		for _, _dbaa := range _dfdb.Keys() {
			_gefe, _bddf := _geb.GetDict(_dfdb.Get(_dbaa))
			if !_bddf {
				if !_daec {
					_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006e\u0064\u0020\u0066on\u0074 \u0070\u0072\u006fg\u0072\u0061\u006ds\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072mi\u006e\u0067\u0020\u0066\u0069\u006ce\u002c\u0020\u0072\u0065\u0067\u0061\u0072\u0064\u006c\u0065s\u0073\u0020\u006f\u0066\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006eg m\u006f\u0064\u0065\u0020\u0075\u0073\u0061\u0067\u0065\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0074\u0068e\u0020\u0070\u0072o\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0069\u006e \u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031:\u0032\u0030\u0030\u0038\u002c \u0039\u002e\u0036\u0020a\u006e\u0064\u0020\u0039.\u0037\u002e"))
					_daec = true
					if _eagg() {
						return _caaf
					}
				}
				continue
			}
			if _bgbfa, _cceg := _geb.GetName(_gefe.Get("\u0054\u0079\u0070\u0065")); !_daec && (!_cceg || _bgbfa.String() != "\u0046\u006f\u006e\u0074") {
				_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_daec = true
				if _eagg() {
					return _caaf
				}
			}
			_cbgb, _ceaf := _ae.NewPdfFontFromPdfObject(_gefe)
			if _ceaf != nil {
				continue
			}
			var _dff string
			if _dagg, _bffa := _geb.GetName(_gefe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bffa {
				_dff = _dagg.String()
			}
			if !_ggebg {
				switch _dff {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_ggebg = true
					_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _eagg() {
						return _caaf
					}
				}
			}
			if !_ffef {
				if _dff != "\u0054\u0079\u0070e\u0033" {
					_edge, _efceg := _geb.GetName(_gefe.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_efceg || _edge.String() == "" {
						_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_ffef = true
						if _eagg() {
							return _caaf
						}
					}
				}
			}
			if _dff != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_acdcce := _dbe.IsStdFont(_dbe.StdFontName(_cbgb.BaseFont()))
			if _acdcce {
				continue
			}
			_bfecb, _adea := _geb.GetIntVal(_gefe.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_adea && !_cfagd {
				_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_cfagd = true
				if _eagg() {
					return _caaf
				}
			}
			_bddb, _bcfd := _geb.GetIntVal(_gefe.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_bcfd && !_cefd {
				_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_cefd = true
				if _eagg() {
					return _caaf
				}
			}
			if !_cdff {
				_faafd, _gded := _geb.GetArray(_gefe.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_gded || !_adea || !_bcfd || _faafd.Len() != _bddb-_bfecb+1 {
					_caaf = append(_caaf, _eef("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_cdff = true
					if _eagg() {
						return _caaf
					}
				}
			}
		}
	}
	return _caaf
}
func _ggefb(_bbca *_ae.CompliancePdfReader) (_gcce []ViolatedRule) {
	_bfaf := true
	_eafa, _dcae := _bbca.GetCatalogMarkInfo()
	if !_dcae {
		_bfaf = false
	} else {
		_degb, _afde := _geb.GetDict(_eafa)
		if _afde {
			_abdbf, _dccfdc := _geb.GetBool(_degb.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_abdbf) || !_dccfdc {
				_bfaf = false
			}
		} else {
			_bfaf = false
		}
	}
	if !_bfaf {
		_gcce = append(_gcce, _eef("\u0036.\u0037\u002e\u0032\u002e\u0032\u002d1", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_eedfc, _dcae := _bbca.GetCatalogStructTreeRoot()
	if !_dcae {
		_gcce = append(_gcce, _eef("\u0036.\u0037\u002e\u0033\u002e\u0033\u002d1", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_eeggg, _dcae := _geb.GetDict(_eedfc)
	if _dcae {
		_dbcde, _fccac := _geb.GetName(_eeggg.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _fccac {
			_badea, _fcdag := _geb.GetDict(_dbcde)
			if _fcdag {
				for _, _bfee := range _badea.Keys() {
					_deabcb := _badea.Get(_bfee)
					if _deabcb == nil {
						_gcce = append(_gcce, _eef("\u0036.\u0037\u002e\u0033\u002e\u0034\u002d1", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _gcce
}
func _eebc(_aagf *_ae.CompliancePdfReader) ViolatedRule { return _fc }

// Validate checks if provided input document reader matches given PDF/A profile.
func Validate(d *_ae.CompliancePdfReader, profile Profile) error { return profile.ValidateStandard(d) }
func _ccag(_ebad *_ae.CompliancePdfReader) (_cbdc []ViolatedRule) {
	var _fedbc, _gdcdd, _dggcd, _cgefd, _bfcd, _fgee bool
	_gbbdc := func() bool { return _fedbc && _gdcdd && _dggcd && _cgefd && _bfcd && _fgee }
	_dfgg := func(_gaagc *_geb.PdfObjectDictionary) bool {
		if !_fedbc && _gaagc.Get("\u0054\u0052") != nil {
			_fedbc = true
			_cbdc = append(_cbdc, _eef("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _gdfeg := _gaagc.Get("\u0054\u0052\u0032"); !_gdcdd && _gdfeg != nil {
			_bfeg, _accg := _geb.GetName(_gdfeg)
			if !_accg || (_accg && *_bfeg != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_gdcdd = true
				_cbdc = append(_cbdc, _eef("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _gbbdc() {
					return true
				}
			}
		}
		if _dgge := _gaagc.Get("\u0053\u004d\u0061s\u006b"); !_dggcd && _dgge != nil {
			_abac, _abaa := _geb.GetName(_dgge)
			if !_abaa || (_abaa && *_abac != "\u004e\u006f\u006e\u0065") {
				_dggcd = true
				_cbdc = append(_cbdc, _eef("\u0036\u002e\u0034-\u0031", "\u0049\u0066\u0020\u0061\u006e \u0053\u004d\u0061\u0073\u006b\u0020\u006be\u0079\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0069\u0074s\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u004e\u006f\u006ee\u002e"))
				if _gbbdc() {
					return true
				}
			}
		}
		if _aagdc := _gaagc.Get("\u0043\u0041"); !_bfcd && _aagdc != nil {
			_cfbad, _fdgg := _geb.GetNumberAsFloat(_aagdc)
			if _fdgg == nil && _cfbad != 1.0 {
				_bfcd = true
				_cbdc = append(_cbdc, _eef("\u0036\u002e\u0034-\u0035", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0043\u0041 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _gbbdc() {
					return true
				}
			}
		}
		if _dfag := _gaagc.Get("\u0063\u0061"); !_fgee && _dfag != nil {
			_aaafg, _eecf := _geb.GetNumberAsFloat(_dfag)
			if _eecf == nil && _aaafg != 1.0 {
				_fgee = true
				_cbdc = append(_cbdc, _eef("\u0036\u002e\u0034-\u0036", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0063\u0061 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _gbbdc() {
					return true
				}
			}
		}
		if _ggec := _gaagc.Get("\u0042\u004d"); !_cgefd && _ggec != nil {
			_febff, _gdec := _geb.GetName(_ggec)
			if _gdec {
				switch _febff.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
				default:
					_cgefd = true
					_cbdc = append(_cbdc, _eef("\u0036\u002e\u0034-\u0034", "T\u0068\u0065\u0020\u0066\u006f\u006cl\u006f\u0077\u0069\u006e\u0067 \u006b\u0065y\u0073\u002c\u0020\u0069\u0066 \u0070res\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047S\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065 \u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0077n\u003a\u0020\u0042\u004d\u0020\u002d\u0020\u004e\u006f\u0072m\u0061\u006c\u0020\u006f\u0072\u0020\u0043\u006f\u006d\u0070\u0061t\u0069\u0062\u006c\u0065\u002e"))
					if _gbbdc() {
						return true
					}
				}
			}
		}
		return false
	}
	for _, _acag := range _ebad.PageList {
		_fegcf := _acag.Resources
		if _fegcf == nil {
			continue
		}
		if _fegcf.ExtGState == nil {
			continue
		}
		_eebbc, _effd := _geb.GetDict(_fegcf.ExtGState)
		if !_effd {
			continue
		}
		_gfcca := _eebbc.Keys()
		for _, _geff := range _gfcca {
			_dfdgd, _gdca := _geb.GetDict(_eebbc.Get(_geff))
			if !_gdca {
				continue
			}
			if _dfgg(_dfdgd) {
				return _cbdc
			}
		}
	}
	for _, _ffdg := range _ebad.PageList {
		_aegf := _ffdg.Resources
		if _aegf == nil {
			continue
		}
		_efdb, _gdcg := _geb.GetDict(_aegf.XObject)
		if !_gdcg {
			continue
		}
		for _, _cbdb := range _efdb.Keys() {
			_facb, _agfdg := _geb.GetStream(_efdb.Get(_cbdb))
			if !_agfdg {
				continue
			}
			_fbecf, _agfdg := _geb.GetDict(_facb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_agfdg {
				continue
			}
			_baeg, _agfdg := _geb.GetDict(_fbecf.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_agfdg {
				continue
			}
			for _, _dcgf := range _baeg.Keys() {
				_ggdd, _bade := _geb.GetDict(_baeg.Get(_dcgf))
				if !_bade {
					continue
				}
				if _dfgg(_ggdd) {
					return _cbdc
				}
			}
		}
	}
	return _cbdc
}
func _aebf(_gada *_bag.Document, _cfc bool) error {
	_cda, _egbg := _gada.GetPages()
	if !_egbg {
		return nil
	}
	for _, _acb := range _cda {
		_abee := _acb.FindXObjectForms()
		for _, _efcd := range _abee {
			_dggab, _ggcb := _ae.NewXObjectFormFromStream(_efcd)
			if _ggcb != nil {
				return _ggcb
			}
			_ebfd, _ggcb := _dggab.GetContentStream()
			if _ggcb != nil {
				return _ggcb
			}
			_eagb := _gg.NewContentStreamParser(string(_ebfd))
			_fefb, _ggcb := _eagb.Parse()
			if _ggcb != nil {
				return _ggcb
			}
			_egfb, _ggcb := _dgac(_dggab.Resources, _fefb, _cfc)
			if _ggcb != nil {
				return _ggcb
			}
			if len(_egfb) == 0 {
				continue
			}
			if _ggcb = _dggab.SetContentStream(_egfb, _geb.NewFlateEncoder()); _ggcb != nil {
				return _ggcb
			}
			_dggab.ToPdfObject()
		}
	}
	return nil
}
func _beef(_daea *_ae.CompliancePdfReader) (_cfdag ViolatedRule) {
	_beegc, _ggaf := _debed(_daea)
	if !_ggaf {
		return _fc
	}
	if _beegc.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		return _eef("\u0036\u002e\u0031\u0031\u002d\u0031", "Th\u0065\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063a\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020R\u0065q\u0075\u0069\u0072\u0065\u006d\u0065\u006e\u0074s\u0020k\u0065\u0079.")
	}
	return _fc
}
func _gefg(_bebf *_bag.Document) error {
	_eeb, _eaef := _bebf.FindCatalog()
	if !_eaef {
		return _e.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_acea, _eaef := _geb.GetDict(_eeb.Object.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_eaef {
		return nil
	}
	_faeb, _eaef := _geb.GetDict(_acea.Get("\u0044"))
	if _eaef {
		if _faeb.Get("\u0041\u0053") != nil {
			_faeb.Remove("\u0041\u0053")
		}
	}
	_gac, _eaef := _geb.GetArray(_acea.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _eaef {
		for _cabb := 0; _cabb < _gac.Len(); _cabb++ {
			_abfg, _cac := _geb.GetDict(_gac.Get(_cabb))
			if !_cac {
				continue
			}
			if _abfg.Get("\u0041\u0053") != nil {
				_abfg.Remove("\u0041\u0053")
			}
		}
	}
	return nil
}
func _febb(_bgaa *_ae.CompliancePdfReader) (_dgbgf []ViolatedRule) {
	var _acgf, _dddc, _aagg, _bagfa, _fdced, _dbdc, _dfeb bool
	_ddaa := func() bool { return _acgf && _dddc && _aagg && _bagfa && _fdced && _dbdc && _dfeb }
	for _, _cffb := range _bgaa.PageList {
		if _cffb.Resources == nil {
			continue
		}
		_dbcf, _bgff := _geb.GetDict(_cffb.Resources.Font)
		if !_bgff {
			continue
		}
		for _, _fcdb := range _dbcf.Keys() {
			_edgd, _gcbfb := _geb.GetDict(_dbcf.Get(_fcdb))
			if !_gcbfb {
				if !_acgf {
					_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006fn\u0074\u0073\u0020\u0075\u0073e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0066\u006f\u0072m\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u0020d\u0065\u0066\u0069\u006e\u0065d \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035\u002e\u0035\u002e"))
					_acgf = true
					if _ddaa() {
						return _dgbgf
					}
				}
				continue
			}
			if _fefda, _deee := _geb.GetName(_edgd.Get("\u0054\u0079\u0070\u0065")); !_acgf && (!_deee || _fefda.String() != "\u0046\u006f\u006e\u0074") {
				_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_acgf = true
				if _ddaa() {
					return _dgbgf
				}
			}
			_edaf, _bdfge := _ae.NewPdfFontFromPdfObject(_edgd)
			if _bdfge != nil {
				continue
			}
			var _ggcdg string
			if _ecfc, _aagdb := _geb.GetName(_edgd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _aagdb {
				_ggcdg = _ecfc.String()
			}
			if !_dddc {
				switch _ggcdg {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_dddc = true
					_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _ddaa() {
						return _dgbgf
					}
				}
			}
			if !_aagg {
				if _ggcdg != "\u0054\u0079\u0070e\u0033" {
					_gbbf, _gdcbg := _geb.GetName(_edgd.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_gdcbg || _gbbf.String() == "" {
						_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_aagg = true
						if _ddaa() {
							return _dgbgf
						}
					}
				}
			}
			if _ggcdg != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_ebbf := _dbe.IsStdFont(_dbe.StdFontName(_edaf.BaseFont()))
			if _ebbf {
				continue
			}
			_baca, _dcbc := _geb.GetIntVal(_edgd.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_dcbc && !_bagfa {
				_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_bagfa = true
				if _ddaa() {
					return _dgbgf
				}
			}
			_bgcc, _efcc := _geb.GetIntVal(_edgd.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_efcc && !_fdced {
				_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_fdced = true
				if _ddaa() {
					return _dgbgf
				}
			}
			if !_dbdc {
				_aggf, _fgdca := _geb.GetArray(_edgd.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_fgdca || !_dcbc || !_efcc || _aggf.Len() != _bgcc-_baca+1 {
					_dgbgf = append(_dgbgf, _eef("\u0036.\u0033\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_dbdc = true
					if _ddaa() {
						return _dgbgf
					}
				}
			}
		}
	}
	return _dgbgf
}
func _ebeee(_bcfe *_ae.PdfFont, _gfbf *_geb.PdfObjectDictionary, _fcaf bool) ViolatedRule {
	const (
		_gefb = "\u0036.\u0033\u002e\u0034\u002d\u0031"
		_eccb = "\u0054\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u0020s\u0068\u0061\u006cl\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006ce\u002c\u0020a\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0035\u002e\u0038\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0077h\u0065\u006e\u0020\u0074\u0068\u0065 \u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0065\u0078\u0063\u006cu\u0073i\u0076\u0065\u006c\u0079\u0020\u0077\u0069t\u0068\u0020\u0074\u0065\u0078\u0074\u0020\u0072e\u006ed\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020\u0033\u002e"
	)
	if _fcaf {
		return _fc
	}
	_ceab := _bcfe.FontDescriptor()
	var _agba string
	if _gadaf, _bgfeg := _geb.GetName(_gfbf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bgfeg {
		_agba = _gadaf.String()
	}
	switch _agba {
	case "\u0054\u0079\u0070e\u0031":
		if _ceab.FontFile == nil {
			return _eef(_gefb, _eccb)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _ceab.FontFile2 == nil {
			return _eef(_gefb, _eccb)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _ceab.FontFile3 == nil {
			return _eef(_gefb, _eccb)
		}
	}
	return _fc
}
func _bgea(_dcad *_bag.Document) error {
	_cee, _cdae := _dcad.GetPages()
	if !_cdae {
		return nil
	}
	for _, _eba := range _cee {
		_fgbd := _eba.FindXObjectForms()
		for _, _cbd := range _fgbd {
			_egegg, _bec := _geb.GetDict(_cbd.Get("\u0047\u0072\u006fu\u0070"))
			if _bec {
				if _afeg := _egegg.Get("\u0053"); _afeg != nil {
					_fafbg, _bead := _geb.GetName(_afeg)
					if _bead && _fafbg.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_cbd.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_gfca, _fdg := _eba.GetResourcesXObject()
		if _fdg {
			_dcdg, _cdb := _geb.GetDict(_gfca.Get("\u0047\u0072\u006fu\u0070"))
			if _cdb {
				_egag := _dcdg.Get("\u0053")
				if _egag != nil {
					_bgcg, _baee := _geb.GetName(_egag)
					if _baee && _bgcg.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_gfca.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_ecd, _cgc := _geb.GetDict(_eba.Object.Get("\u0047\u0072\u006fu\u0070"))
		if _cgc {
			_bgeg := _ecd.Get("\u0053")
			if _bgeg != nil {
				_ccdf, _bagc := _geb.GetName(_bgeg)
				if _bagc && _ccdf.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
					_eba.Object.Remove("\u0047\u0072\u006fu\u0070")
				}
			}
		}
	}
	return nil
}

// NewProfile3A creates a new Profile3A with given options.
func NewProfile3A(options *Profile3Options) *Profile3A {
	if options == nil {
		options = DefaultProfile3Options()
	}
	_aged(options)
	return &Profile3A{profile3{_ebg: *options, _egbf: _ec()}}
}

// DefaultProfile1Options are the default options for the Profile1.
func DefaultProfile1Options() *Profile1Options {
	return &Profile1Options{Now: _ff.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _gfef(_dfge *_bag.Document) error {
	_bfe, _deff := _dfge.FindCatalog()
	if !_deff {
		return nil
	}
	_, _deff = _bfe.GetStructTreeRoot()
	if !_deff {
		_aeeg := _ae.NewStructTreeRoot()
		_fae := _aeeg.ToPdfObject().(*_geb.PdfIndirectObject)
		_fcb := _fae.PdfObject.(*_geb.PdfObjectDictionary)
		_bfe.SetStructTreeRoot(_fcb)
	}
	return nil
}

type profile2 struct {
	_ceceg standardType
	_feca  Profile2Options
}

func _fdc(_adgg *_bag.Document, _fbd int) {
	if _adgg.Version.Major == 0 {
		_adgg.Version.Major = 1
	}
	if _adgg.Version.Minor < _fbd {
		_adgg.Version.Minor = _fbd
	}
}
func _cdfeb(_ceaa *_ae.CompliancePdfReader) (_fcdd []ViolatedRule) {
	_dfacb, _bcaff := _debed(_ceaa)
	if !_bcaff {
		return _fcdd
	}
	_dffd, _bcaff := _geb.GetDict(_dfacb.Get("\u004e\u0061\u006de\u0073"))
	if !_bcaff {
		return _fcdd
	}
	if _dffd.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_fcdd = append(_fcdd, _eef("\u0036\u002e\u0031\u0030\u002d\u0031", "T\u0068\u0065\u0072e\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u006e\u006f\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0050\u0072\u0065s\u0065\u006e\u0074a\u0074\u0069\u006f\u006e\u0073\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002e"))
	}
	return _fcdd
}
func (_bf standardType) String() string {
	return _d.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _bf._fgb, _bf._ea)
}

type imageInfo struct {
	ColorSpace       _geb.PdfObjectName
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_geb.PdfObjectStream
	_fga             bool
}

func _bbed(_adbb *_ae.CompliancePdfReader, _fcag standardType) (_gecb []ViolatedRule) {
	var _ggff, _febbf, _edgdb, _ffffc, _fgaeb, _gadcc, _babg bool
	_acfa := func() bool { return _ggff && _febbf && _edgdb && _ffffc && _fgaeb && _gadcc && _babg }
	_fbga := map[*_geb.PdfObjectStream]*_gba.CMap{}
	_ffag := map[*_geb.PdfObjectStream][]byte{}
	_dfgcga := map[_geb.PdfObject]*_ae.PdfFont{}
	for _, _ecegb := range _adbb.GetObjectNums() {
		_dbcdd, _egdf := _adbb.GetIndirectObjectByNumber(_ecegb)
		if _egdf != nil {
			continue
		}
		_gfcd, _gabcb := _geb.GetDict(_dbcdd)
		if !_gabcb {
			continue
		}
		_ecfb, _gabcb := _geb.GetName(_gfcd.Get("\u0054\u0079\u0070\u0065"))
		if !_gabcb {
			continue
		}
		if *_ecfb != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_eea, _egdf := _ae.NewPdfFontFromPdfObject(_gfcd)
		if _egdf != nil {
			_gd.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _egdf)
			continue
		}
		_dfgcga[_gfcd] = _eea
	}
	for _, _dcbdg := range _adbb.PageList {
		_egbe, _cfda := _dcbdg.GetContentStreams()
		if _cfda != nil {
			_gd.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _dgee := range _egbe {
			_cfee := _gg.NewContentStreamParser(_dgee)
			_effa, _cfgf := _cfee.Parse()
			if _cfgf != nil {
				_gd.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cfgf)
				continue
			}
			var _faebb bool
			for _, _gbbc := range *_effa {
				if _gbbc.Operand != "\u0054\u0072" {
					continue
				}
				if len(_gbbc.Params) != 1 {
					_gd.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_gbbc.Params))
					continue
				}
				_eded, _ccae := _geb.GetIntVal(_gbbc.Params[0])
				if !_ccae {
					_gd.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _eded == 3 {
					_faebb = true
					break
				}
			}
			for _, _dgacc := range *_effa {
				if _dgacc.Operand != "\u0054\u0066" {
					continue
				}
				if len(_dgacc.Params) != 2 {
					_gd.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_dgacc.Params))
					continue
				}
				_efea, _aeef := _geb.GetName(_dgacc.Params[0])
				if !_aeef {
					_gd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _dgacc)
					continue
				}
				_gdfda, _dgcgb := _dcbdg.Resources.GetFontByName(*_efea)
				if !_dgcgb {
					_gd.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_dabbc, _aeef := _geb.GetDict(_gdfda)
				if !_aeef {
					_gd.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_bddg, _aeef := _dfgcga[_dabbc]
				if !_aeef {
					var _beag error
					_bddg, _beag = _ae.NewPdfFontFromPdfObject(_dabbc)
					if _beag != nil {
						_gd.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _beag)
						continue
					}
					_dfgcga[_dabbc] = _bddg
				}
				if !_ggff {
					_dcdbc := _dfggc(_dabbc, _ffag, _fbga)
					if _dcdbc != _fc {
						_gecb = append(_gecb, _dcdbc)
						_ggff = true
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_febbf {
					_bcfdd := _edddc(_dabbc)
					if _bcfdd != _fc {
						_gecb = append(_gecb, _bcfdd)
						_febbf = true
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_edgdb {
					_dbbd := _gedf(_dabbc, _ffag, _fbga)
					if _dbbd != _fc {
						_gecb = append(_gecb, _dbbd)
						_edgdb = true
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_ffffc {
					_acfff := _cagb(_dabbc, _ffag, _fbga)
					if _acfff != _fc {
						_gecb = append(_gecb, _acfff)
						_ffffc = true
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_fgaeb {
					_cbaf := _afeeca(_bddg, _dabbc, _faebb)
					if _cbaf != _fc {
						_fgaeb = true
						_gecb = append(_gecb, _cbaf)
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_gadcc {
					_eagdg := _dabae(_bddg, _dabbc)
					if _eagdg != _fc {
						_gadcc = true
						_gecb = append(_gecb, _eagdg)
						if _acfa() {
							return _gecb
						}
					}
				}
				if !_babg && (_fcag._ea == "\u0041" || _fcag._ea == "\u0055") {
					_feef := _befb(_dabbc, _ffag, _fbga)
					if _feef != _fc {
						_babg = true
						_gecb = append(_gecb, _feef)
						if _acfa() {
							return _gecb
						}
					}
				}
			}
		}
	}
	return _gecb
}

type documentColorspaceOptimizeFunc func(_ebff *_bag.Document, _abgc []*_bag.Image) error

func _bdf(_dgb *_bag.Document, _cfe int) error {
	for _, _gge := range _dgb.Objects {
		_bbd, _adgb := _geb.GetDict(_gge)
		if !_adgb {
			continue
		}
		_bada := _bbd.Get("\u0054\u0079\u0070\u0065")
		if _bada == nil {
			continue
		}
		if _baf, _fce := _geb.GetName(_bada); _fce && _baf.String() != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_dcdd, _fgaa := _geb.GetName(_bbd.Get("\u0053"))
		if !_fgaa {
			continue
		}
		switch _ae.PdfActionType(*_dcdd) {
		case _ae.ActionTypeLaunch, _ae.ActionTypeSound, _ae.ActionTypeMovie, _ae.ActionTypeResetForm, _ae.ActionTypeImportData, _ae.ActionTypeJavaScript:
			_bbd.Remove("\u0053")
		case _ae.ActionTypeHide, _ae.ActionTypeSetOCGState, _ae.ActionTypeRendition, _ae.ActionTypeTrans, _ae.ActionTypeGoTo3DView:
			if _cfe == 2 {
				_bbd.Remove("\u0053")
			}
		case _ae.ActionTypeNamed:
			_dec, _acd := _geb.GetName(_bbd.Get("\u004e"))
			if !_acd {
				continue
			}
			switch *_dec {
			case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
			default:
				_bbd.Remove("\u004e")
			}
		}
	}
	return nil
}
func _fgbdf(_bab *_ae.CompliancePdfReader) (_bgfg []ViolatedRule) {
	var (
		_fdae, _egdd, _gegb, _bcfa, _fegg, _cggc, _eaed bool
		_gged                                           func(_geb.PdfObject)
	)
	_gged = func(_feggb _geb.PdfObject) {
		switch _ddecf := _feggb.(type) {
		case *_geb.PdfObjectInteger:
			if !_fdae && (int64(*_ddecf) > _dg.MaxInt32 || int64(*_ddecf) < -_dg.MaxInt32) {
				_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_fdae = true
			}
		case *_geb.PdfObjectFloat:
			if !_egdd && (_dg.Abs(float64(*_ddecf)) > 32767.0) {
				_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0032", "\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0020\u0072\u0065\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u00332\u0037\u0036\u0037.\u0030\u002e"))
			}
		case *_geb.PdfObjectString:
			if !_gegb && len([]byte(_ddecf.Str())) > 65535 {
				_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0036\u0035\u0035\u0033\u0035\u002e"))
				_gegb = true
			}
		case *_geb.PdfObjectName:
			if !_bcfa && len([]byte(*_ddecf)) > 127 {
				_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_bcfa = true
			}
		case *_geb.PdfObjectArray:
			if !_fegg && _ddecf.Len() > 8191 {
				_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0035", "\u004d\u0061\u0078\u0069\u006d\u0075m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020(\u0069\u006e\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0029\u0020\u0069s\u00208\u0031\u0039\u0031\u002e"))
				_fegg = true
			}
			for _, _ceb := range _ddecf.Elements() {
				_gged(_ceb)
			}
			if !_eaed && (_ddecf.Len() == 4 || _ddecf.Len() == 5) {
				_edbe, _dbec := _geb.GetName(_ddecf.Get(0))
				if !_dbec {
					return
				}
				if *_edbe != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_dgcg := _ddecf.Get(1)
				_dgcg = _geb.TraceToDirectObject(_dgcg)
				_eagfg, _dbec := _geb.GetArray(_dgcg)
				if !_dbec {
					return
				}
				if _eagfg.Len() > 8 {
					_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0039", "\u004d\u0061\u0078i\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0069\u0073\u0020\u0038\u002e"))
					_eaed = true
				}
			}
		case *_geb.PdfObjectDictionary:
			_ageec := _ddecf.Keys()
			if !_cggc && len(_ageec) > 4095 {
				_bgfg = append(_bgfg, _eef("\u0036.\u0031\u002e\u0031\u0032\u002d\u00311", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u0063\u0061\u0070\u0061\u0063\u0069\u0074y\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0028\u0069\u006e\u0020\u0065\u006e\u0074\u0072\u0069es\u0029\u0020\u0069\u0073\u0020\u0034\u0030\u0039\u0035\u002e"))
				_cggc = true
			}
			for _eecc, _dgcf := range _ageec {
				_gged(&_ageec[_eecc])
				_gged(_ddecf.Get(_dgcf))
			}
		case *_geb.PdfObjectStream:
			_gged(_ddecf.PdfObjectDictionary)
		case *_geb.PdfObjectStreams:
			for _, _abcc := range _ddecf.Elements() {
				_gged(_abcc)
			}
		case *_geb.PdfObjectReference:
			_gged(_ddecf.Resolve())
		}
	}
	_aegg := _bab.GetObjectNums()
	if len(_aegg) > 8388607 {
		_bgfg = append(_bgfg, _eef("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _dcca := range _aegg {
		_dcfc, _bdfad := _bab.GetIndirectObjectByNumber(_dcca)
		if _bdfad != nil {
			continue
		}
		_gce := _geb.TraceToDirectObject(_dcfc)
		_gged(_gce)
	}
	return _bgfg
}
func _ddadb(_ebbda *_ae.CompliancePdfReader, _ffea standardType, _abbeg bool) (_cfgde []ViolatedRule) {
	_bagbe, _ecef := _debed(_ebbda)
	if !_ecef {
		return []ViolatedRule{_eef("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_ddegb := _bagbe.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _ddegb == nil {
		return []ViolatedRule{_eef("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_ebbfd, _ecef := _geb.GetStream(_ddegb)
	if !_ecef {
		return []ViolatedRule{_eef("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_fddde, _ebcc := _eg.LoadDocument(_ebbfd.Stream)
	if _ebcc != nil {
		return []ViolatedRule{_eef("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d4", "\u0041\u006c\u006c\u0020\u006de\u0074\u0061\u0064a\u0074\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020i\u006e \u0074\u0068\u0065\u0020\u0050\u0044\u0046 \u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065ci\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0041\u006c\u006c\u0020c\u006fn\u0074\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0061\u006c\u006c\u0020\u0058\u004d\u0050\u0020p\u0061\u0063\u006b\u0065\u0074\u0073 \u0073h\u0061\u006c\u006c \u0062\u0065\u0020\u0077\u0065\u006c\u006c\u002d\u0066o\u0072\u006de\u0064")}
	}
	_dafd := _fddde.GetGoXmpDocument()
	var _dfde []*_c.Namespace
	for _, _dgbca := range _dafd.Namespaces() {
		switch _dgbca.Name {
		case _ag.NsDc.Name, _ba.NsPDF.Name, _b.NsXmp.Name, _ge.NsXmpRights.Name, _ab.Namespace.Name, _ee.Namespace.Name, _dd.NsXmpMM.Name, _ee.FieldNS.Name, _ee.SchemaNS.Name, _ee.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_dfde = append(_dfde, _dgbca)
	}
	_beaa := true
	_eegge, _ebcc := _fddde.GetPdfaExtensionSchemas()
	if _ebcc == nil {
		for _, _cebb := range _dfde {
			var _fefa bool
			for _bdcbd := range _eegge {
				if _cebb.URI == _eegge[_bdcbd].NamespaceURI {
					_fefa = true
					break
				}
			}
			if !_fefa {
				_beaa = false
				break
			}
		}
	} else {
		_beaa = false
	}
	if !_beaa {
		_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0032\u002e\u0033\u002d7", "\u0041\u006c\u006c\u0020\u0070\u0072\u006f\u0070e\u0072\u0074\u0069e\u0073\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072m\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0075s\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063he\u006da\u0073 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002c\u0020\u0049\u0053\u004f\u0020\u0031\u00390\u0030\u0035-\u0031\u0020\u006f\u0072\u0020\u0074h\u0069s\u0020\u0070\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u0020\u0031\u0039\u0030\u0030\u0035\u002c\u0020o\u0072\u0020\u0061\u006e\u0079\u0020e\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073c\u0068\u0065\u006das\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006fm\u0070\u006c\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0036\u002e\u0036\u002e\u0032.\u0033\u002e\u0032\u002e"))
	}
	_bece, _ecef := _fddde.GetPdfAID()
	if !_ecef {
		_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0034\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _bece.Part != _ffea._fgb {
			_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0034\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _ffea._ea == "\u0041" && _bece.Conformance != "\u0041" {
			_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _ffea._ea == "\u0055" && (_bece.Conformance != "\u0041" && _bece.Conformance != "\u0055") {
			_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _ffea._ea == "\u0042" && (_bece.Conformance != "\u0041" && _bece.Conformance != "\u0042" && _bece.Conformance != "\u0055") {
			_cfgde = append(_cfgde, _eef("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		}
	}
	return _cfgde
}
func (_dgc *documentImages) hasOnlyDeviceGray() bool { return _dgc._cfb && !_dgc._gdf && !_dgc._cag }

// Conformance gets the PDF/A conformance.
func (_fecc *profile3) Conformance() string { return _fecc._egbf._ea }

// Conformance gets the PDF/A conformance.
func (_dcbda *profile2) Conformance() string { return _dcbda._ceceg._ea }

// Profile3Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile3Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _ff.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _dabd(_fdde *_bag.Document) error {
	_beda := func(_agb *_geb.PdfObjectDictionary) error {
		if _agb.Get("\u0054\u0052") != nil {
			_gd.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_agb.Remove("\u0054\u0052")
		}
		_bafg := _agb.Get("\u0054\u0052\u0032")
		if _bafg != nil {
			_bdc := _bafg.String()
			if _bdc != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_gd.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_agb.Set("\u0054\u0052\u0032", _geb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		if _agb.Get("\u0048\u0054\u0050") != nil {
			_gd.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0073\u0020\u0048\u0054P\u0020\u006b\u0065\u0079")
			_agb.Remove("\u0048\u0054\u0050")
		}
		_fbdeb := _agb.Get("\u0042\u004d")
		if _fbdeb != nil {
			_agcg, _dggd := _geb.GetName(_fbdeb)
			if !_dggd {
				_gd.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_agcg = _geb.MakeName("")
			}
			_cbca := _agcg.String()
			switch _cbca {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
			default:
				_agb.Set("\u0042\u004d", _geb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		return nil
	}
	_cdc, _gaab := _fdde.GetPages()
	if !_gaab {
		return nil
	}
	for _, _bede := range _cdc {
		_eefd, _fac := _bede.GetResources()
		if !_fac {
			continue
		}
		_dcgef, _fffe := _geb.GetDict(_eefd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_fffe {
			return nil
		}
		_bgbf := _dcgef.Keys()
		for _, _fcda := range _bgbf {
			_cec, _fgf := _geb.GetDict(_dcgef.Get(_fcda))
			if !_fgf {
				continue
			}
			_aagd := _beda(_cec)
			if _aagd != nil {
				continue
			}
		}
	}
	for _, _bffc := range _cdc {
		_ddac, _geg := _bffc.GetContents()
		if !_geg {
			return nil
		}
		for _, _bcdgf := range _ddac {
			_bebc, _acdc := _bcdgf.GetData()
			if _acdc != nil {
				continue
			}
			_fbab := _gg.NewContentStreamParser(string(_bebc))
			_cdf, _acdc := _fbab.Parse()
			if _acdc != nil {
				continue
			}
			for _, _fbaf := range *_cdf {
				if len(_fbaf.Params) == 0 {
					continue
				}
				_, _adab := _geb.GetName(_fbaf.Params[0])
				if !_adab {
					continue
				}
				_gea, _dbdf := _bffc.GetResourcesXObject()
				if !_dbdf {
					continue
				}
				for _, _bebe := range _gea.Keys() {
					_abbf, _deffa := _geb.GetStream(_gea.Get(_bebe))
					if !_deffa {
						continue
					}
					_fgef, _deffa := _geb.GetDict(_abbf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_deffa {
						continue
					}
					_eee, _deffa := _geb.GetDict(_fgef.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_deffa {
						continue
					}
					for _, _bba := range _eee.Keys() {
						_abga, _fdb := _geb.GetDict(_eee.Get(_bba))
						if !_fdb {
							continue
						}
						_dcgeg := _beda(_abga)
						if _dcgeg != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _dgba(_efaage *_ae.CompliancePdfReader, _dbdac standardType) (_aabed []ViolatedRule) {
	var _bfff, _bggcd, _bdcd, _gaad, _ceed, _bfbe, _caed, _efbg, _fecd, _ccfdd, _ggde bool
	_gbcbb := func() bool {
		return _bfff && _bggcd && _bdcd && _gaad && _ceed && _bfbe && _caed && _efbg && _fecd && _ccfdd && _ggde
	}
	_ecfcf := map[*_geb.PdfObjectStream]*_gba.CMap{}
	_edgba := map[*_geb.PdfObjectStream][]byte{}
	_bgbe := map[_geb.PdfObject]*_ae.PdfFont{}
	for _, _eede := range _efaage.GetObjectNums() {
		_dagc, _feae := _efaage.GetIndirectObjectByNumber(_eede)
		if _feae != nil {
			continue
		}
		_egbc, _baceb := _geb.GetDict(_dagc)
		if !_baceb {
			continue
		}
		_deaa, _baceb := _geb.GetName(_egbc.Get("\u0054\u0079\u0070\u0065"))
		if !_baceb {
			continue
		}
		if *_deaa != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_caacg, _feae := _ae.NewPdfFontFromPdfObject(_egbc)
		if _feae != nil {
			_gd.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _feae)
			continue
		}
		_bgbe[_egbc] = _caacg
	}
	for _, _gcbd := range _efaage.PageList {
		_caddb, _edgg := _gcbd.GetContentStreams()
		if _edgg != nil {
			_gd.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _bbcf := range _caddb {
			_eege := _gg.NewContentStreamParser(_bbcf)
			_baeaf, _bdgc := _eege.Parse()
			if _bdgc != nil {
				_gd.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bdgc)
				continue
			}
			var _agfba bool
			for _, _fffb := range *_baeaf {
				if _fffb.Operand != "\u0054\u0072" {
					continue
				}
				if len(_fffb.Params) != 1 {
					_gd.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_fffb.Params))
					continue
				}
				_cafg, _eagbg := _geb.GetIntVal(_fffb.Params[0])
				if !_eagbg {
					_gd.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _cafg == 3 {
					_agfba = true
					break
				}
			}
			for _, _agfd := range *_baeaf {
				if _agfd.Operand != "\u0054\u0066" {
					continue
				}
				if len(_agfd.Params) != 2 {
					_gd.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_agfd.Params))
					continue
				}
				_edcg, _cfeff := _geb.GetName(_agfd.Params[0])
				if !_cfeff {
					_gd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _agfd)
					continue
				}
				_ecfe, _cdee := _gcbd.Resources.GetFontByName(*_edcg)
				if !_cdee {
					_gd.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_dbggg, _cfeff := _geb.GetDict(_ecfe)
				if !_cfeff {
					_gd.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_cgag, _cfeff := _bgbe[_dbggg]
				if !_cfeff {
					var _dcga error
					_cgag, _dcga = _ae.NewPdfFontFromPdfObject(_dbggg)
					if _dcga != nil {
						_gd.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _dcga)
						continue
					}
					_bgbe[_dbggg] = _cgag
				}
				if !_bfff {
					_cgcb := _aedda(_dbggg, _edgba, _ecfcf)
					if _cgcb != _fc {
						_aabed = append(_aabed, _cgcb)
						_bfff = true
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_bggcd {
					_bdga := _gbec(_dbggg)
					if _bdga != _fc {
						_aabed = append(_aabed, _bdga)
						_bggcd = true
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_bdcd {
					_affb := _dgcef(_dbggg, _edgba, _ecfcf)
					if _affb != _fc {
						_aabed = append(_aabed, _affb)
						_bdcd = true
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_gaad {
					_fgae := _befc(_dbggg, _edgba, _ecfcf)
					if _fgae != _fc {
						_aabed = append(_aabed, _fgae)
						_gaad = true
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_ceed {
					_dcdbd := _ebeee(_cgag, _dbggg, _agfba)
					if _dcdbd != _fc {
						_ceed = true
						_aabed = append(_aabed, _dcdbd)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_bfbe {
					_dceg := _gfdf(_cgag, _dbggg)
					if _dceg != _fc {
						_bfbe = true
						_aabed = append(_aabed, _dceg)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_caed {
					_gcda := _eggbb(_cgag, _dbggg)
					if _gcda != _fc {
						_caed = true
						_aabed = append(_aabed, _gcda)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_efbg {
					_afcb := _fcfc(_cgag, _dbggg)
					if _afcb != _fc {
						_efbg = true
						_aabed = append(_aabed, _afcb)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_fecd {
					_gfbed := _gebd(_cgag, _dbggg)
					if _gfbed != _fc {
						_fecd = true
						_aabed = append(_aabed, _gfbed)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_ccfdd {
					_baab := _eaabc(_cgag, _dbggg)
					if _baab != _fc {
						_ccfdd = true
						_aabed = append(_aabed, _baab)
						if _gbcbb() {
							return _aabed
						}
					}
				}
				if !_ggde && _dbdac._ea == "\u0041" {
					_gbcbc := _fadc(_dbggg, _edgba, _ecfcf)
					if _gbcbc != _fc {
						_ggde = true
						_aabed = append(_aabed, _gbcbc)
						if _gbcbb() {
							return _aabed
						}
					}
				}
			}
		}
	}
	return _aabed
}
func _bbdbb(_ccacd *_ae.CompliancePdfReader) ViolatedRule { return _fc }
func _ffa(_cefcd *_ae.CompliancePdfReader) ViolatedRule {
	_dfaga := map[*_geb.PdfObjectStream]struct{}{}
	for _, _fecb := range _cefcd.PageList {
		if _fecb.Resources == nil && _fecb.Contents == nil {
			continue
		}
		if _eadc := _fecb.GetPageDict(); _eadc != nil {
			_cfbef, _fdba := _geb.GetDict(_eadc.Get("\u0047\u0072\u006fu\u0070"))
			if _fdba {
				if _ddbb := _cfbef.Get("\u0053"); _ddbb != nil {
					_dfada, _gebgb := _geb.GetName(_ddbb)
					if _gebgb && _dfada.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						return _eef("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
					}
				}
			}
		}
		if _fecb.Resources != nil {
			if _edbed, _ecec := _geb.GetDict(_fecb.Resources.XObject); _ecec {
				for _, _gggcb := range _edbed.Keys() {
					_eeeg, _adgfe := _geb.GetStream(_edbed.Get(_gggcb))
					if !_adgfe {
						continue
					}
					if _, _gdgf := _dfaga[_eeeg]; _gdgf {
						continue
					}
					_dade, _adgfe := _geb.GetDict(_eeeg.Get("\u0047\u0072\u006fu\u0070"))
					if !_adgfe {
						_dfaga[_eeeg] = struct{}{}
						continue
					}
					_cbfa := _dade.Get("\u0053")
					if _cbfa != nil {
						_cebd, _agfc := _geb.GetName(_cbfa)
						if _agfc && _cebd.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _eef("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_dfaga[_eeeg] = struct{}{}
					continue
				}
			}
		}
		if _fecb.Contents != nil {
			_fbea, _cbgc := _fecb.GetContentStreams()
			if _cbgc != nil {
				continue
			}
			for _, _dccg := range _fbea {
				_cedf, _aegd := _gg.NewContentStreamParser(_dccg).Parse()
				if _aegd != nil {
					continue
				}
				for _, _egge := range *_cedf {
					if len(_egge.Params) == 0 {
						continue
					}
					_cdba, _dcbfa := _geb.GetName(_egge.Params[0])
					if !_dcbfa {
						continue
					}
					_eecbc, _bgcgg := _fecb.Resources.GetXObjectByName(*_cdba)
					if _bgcgg != _ae.XObjectTypeForm {
						continue
					}
					if _, _bcde := _dfaga[_eecbc]; _bcde {
						continue
					}
					_gddg, _dcbfa := _geb.GetDict(_eecbc.Get("\u0047\u0072\u006fu\u0070"))
					if !_dcbfa {
						_dfaga[_eecbc] = struct{}{}
						continue
					}
					_aeae := _gddg.Get("\u0053")
					if _aeae != nil {
						_cggd, _cbdcf := _geb.GetName(_aeae)
						if _cbdcf && _cggd.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _eef("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_dfaga[_eecbc] = struct{}{}
				}
			}
		}
	}
	return _fc
}
func _fdac(_ccba *_ae.CompliancePdfReader, _fbgf bool) (_bagf []ViolatedRule) {
	var _ecae, _eaab, _aadc, _cabg, _fdaea, _fgefe, _abge bool
	_bbggc := func() bool { return _ecae && _eaab && _aadc && _cabg && _fdaea && _fgefe && _abge }
	_ffcf, _cbdd := _fdaf(_ccba)
	var _afedf _bae.ProfileHeader
	if _cbdd {
		_afedf, _ = _bae.ParseHeader(_ffcf.DestOutputProfile)
	}
	var _eddb bool
	_gbde := map[_geb.PdfObject]struct{}{}
	var _ebba func(_bdfae _ae.PdfColorspace) bool
	_ebba = func(_afbc _ae.PdfColorspace) bool {
		switch _cfecg := _afbc.(type) {
		case *_ae.PdfColorspaceDeviceGray:
			if !_fgefe {
				if !_cbdd {
					_eddb = true
					_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0033\u002d\u0034", "\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0075s\u0065\u0064\u0020\u006f\u006el\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006ce\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020O\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u002e"))
					_fgefe = true
					if _bbggc() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceDeviceRGB:
			if !_cabg {
				if !_cbdd || _afedf.ColorSpace != _bae.ColorSpaceRGB {
					_eddb = true
					_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0033\u002d\u0032", "\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065 \u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074In\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020u\u0073es\u0020a\u006e\u0020\u0052\u0047\u0042\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u002e"))
					_cabg = true
					if _bbggc() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceDeviceCMYK:
			if !_fdaea {
				if !_cbdd || _afedf.ColorSpace != _bae.ColorSpaceCMYK {
					_eddb = true
					_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0033\u002d\u0033", "\u0044\u0065\u0076\u0069\u0063e\u0043\u004d\u0059\u004b \u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074h\u0065\u0020\u0066\u0069\u006ce \u0068\u0061\u0073\u0020\u0061 \u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068a\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0061\u006e \u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072\u0020s\u0070\u0061\u0063e\u002e"))
					_fdaea = true
					if _bbggc() {
						return true
					}
				}
			}
		case *_ae.PdfColorspaceICCBased:
			if !_aadc || !_abge {
				_gdefg, _gcee := _bae.ParseHeader(_cfecg.Data)
				if _gcee != nil {
					_gd.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gcee)
					_bagf = append(_bagf, func() ViolatedRule {
						return _eef("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035")
					}())
					_aadc = true
					if _bbggc() {
						return true
					}
				}
				if !_aadc {
					var _dccf, _cdbf bool
					switch _gdefg.DeviceClass {
					case _bae.DeviceClassPRTR, _bae.DeviceClassMNTR, _bae.DeviceClassSCNR, _bae.DeviceClassSPAC:
					default:
						_dccf = true
					}
					switch _gdefg.ColorSpace {
					case _bae.ColorSpaceRGB, _bae.ColorSpaceCMYK, _bae.ColorSpaceGRAY, _bae.ColorSpaceLAB:
					default:
						_cdbf = true
					}
					if _dccf || _cdbf {
						_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035"))
						_aadc = true
						if _bbggc() {
							return true
						}
					}
				}
				if !_abge {
					_ecc, _ := _geb.GetStream(_cfecg.GetContainingPdfObject())
					if _ecc.Get("\u004e") == nil || (_cfecg.N == 1 && _gdefg.ColorSpace != _bae.ColorSpaceGRAY) || (_cfecg.N == 3 && !(_gdefg.ColorSpace == _bae.ColorSpaceRGB || _gdefg.ColorSpace == _bae.ColorSpaceLAB)) || (_cfecg.N == 4 && _gdefg.ColorSpace != _bae.ColorSpaceCMYK) {
						_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0033\u002d\u0035", "\u0049\u0066\u0020a\u006e\u0020u\u006e\u0063\u0061\u006c\u0069\u0062\u0072a\u0074\u0065\u0064\u0020\u0063\u006fl\u006f\u0072 \u0073\u0070\u0061c\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u0069\u006c\u0065 \u0074\u0068\u0065\u006e \u0074\u0068\u0061\u0074 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041-\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020\u0069\u006e\u0020\u0036\u002e\u0032\u002e\u0032\u002e"))
						_abge = true
						if _bbggc() {
							return true
						}
					}
				}
			}
			if _cfecg.Alternate != nil {
				return _ebba(_cfecg.Alternate)
			}
		}
		return false
	}
	for _, _bafce := range _ccba.GetObjectNums() {
		_fcgd, _ecdg := _ccba.GetIndirectObjectByNumber(_bafce)
		if _ecdg != nil {
			continue
		}
		_egafe, _cfbb := _geb.GetStream(_fcgd)
		if !_cfbb {
			continue
		}
		_cfg, _cfbb := _geb.GetName(_egafe.Get("\u0054\u0079\u0070\u0065"))
		if !_cfbb || _cfg.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_fcaa, _cfbb := _geb.GetName(_egafe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_cfbb {
			continue
		}
		_gbde[_egafe] = struct{}{}
		switch _fcaa.String() {
		case "\u0049\u006d\u0061g\u0065":
			_gffg, _ggada := _ae.NewXObjectImageFromStream(_egafe)
			if _ggada != nil {
				continue
			}
			_gbde[_egafe] = struct{}{}
			if _ebba(_gffg.ColorSpace) {
				return _bagf
			}
		case "\u0046\u006f\u0072\u006d":
			_cbdg, _gdagd := _geb.GetDict(_egafe.Get("\u0047\u0072\u006fu\u0070"))
			if !_gdagd {
				continue
			}
			_dfda := _cbdg.Get("\u0043\u0053")
			if _dfda == nil {
				continue
			}
			_eeggb, _aceg := _ae.NewPdfColorspaceFromPdfObject(_dfda)
			if _aceg != nil {
				continue
			}
			if _ebba(_eeggb) {
				return _bagf
			}
		}
	}
	for _, _ageeca := range _ccba.PageList {
		_eddag, _fdgd := _ageeca.GetContentStreams()
		if _fdgd != nil {
			continue
		}
		for _, _fddf := range _eddag {
			_febc, _dgdf := _gg.NewContentStreamParser(_fddf).Parse()
			if _dgdf != nil {
				continue
			}
			for _, _dafe := range *_febc {
				if len(_dafe.Params) > 1 {
					continue
				}
				switch _dafe.Operand {
				case "\u0042\u0049":
					_cdcc, _cgef := _dafe.Params[0].(*_gg.ContentStreamInlineImage)
					if !_cgef {
						continue
					}
					_dfca, _ddba := _cdcc.GetColorSpace(_ageeca.Resources)
					if _ddba != nil {
						continue
					}
					if _ebba(_dfca) {
						return _bagf
					}
				case "\u0044\u006f":
					_bfcc, _cceb := _geb.GetName(_dafe.Params[0])
					if !_cceb {
						continue
					}
					_agbf, _adebd := _ageeca.Resources.GetXObjectByName(*_bfcc)
					if _, _ebbd := _gbde[_agbf]; _ebbd {
						continue
					}
					switch _adebd {
					case _ae.XObjectTypeImage:
						_dgag, _fcdf := _ae.NewXObjectImageFromStream(_agbf)
						if _fcdf != nil {
							continue
						}
						_gbde[_agbf] = struct{}{}
						if _ebba(_dgag.ColorSpace) {
							return _bagf
						}
					case _ae.XObjectTypeForm:
						_eecb, _fbege := _geb.GetDict(_agbf.Get("\u0047\u0072\u006fu\u0070"))
						if !_fbege {
							continue
						}
						_dbbf, _fbege := _geb.GetName(_eecb.Get("\u0043\u0053"))
						if !_fbege {
							continue
						}
						_gcae, _cfbg := _ae.NewPdfColorspaceFromPdfObject(_dbbf)
						if _cfbg != nil {
							continue
						}
						_gbde[_agbf] = struct{}{}
						if _ebba(_gcae) {
							return _bagf
						}
					}
				}
			}
		}
	}
	if !_eddb {
		return _bagf
	}
	if (_afedf.DeviceClass == _bae.DeviceClassPRTR || _afedf.DeviceClass == _bae.DeviceClassMNTR) && (_afedf.ColorSpace == _bae.ColorSpaceRGB || _afedf.ColorSpace == _bae.ColorSpaceCMYK || _afedf.ColorSpace == _bae.ColorSpaceGRAY) {
		return _bagf
	}
	if !_fbgf {
		return _bagf
	}
	_edbb, _febf := _debed(_ccba)
	if !_febf {
		return _bagf
	}
	_efce, _febf := _geb.GetArray(_edbb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_febf {
		_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e"), _eef("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
		return _bagf
	}
	if _efce.Len() > 1 {
		_acba := map[*_geb.PdfObjectDictionary]struct{}{}
		for _bgaba := 0; _bgaba < _efce.Len(); _bgaba++ {
			_dabg, _ecgd := _geb.GetDict(_efce.Get(_bgaba))
			if !_ecgd {
				continue
			}
			if _bgaba == 0 {
				_acba[_dabg] = struct{}{}
				continue
			}
			if _, _ecac := _acba[_dabg]; !_ecac {
				_bagf = append(_bagf, _eef("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
				break
			}
		}
	}
	return _bagf
}
func _ec() standardType { return standardType{_fgb: 3, _ea: "\u0041"} }

// DefaultProfile3Options the default options for the Profile3.
func DefaultProfile3Options() *Profile3Options {
	return &Profile3Options{Now: _ff.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}

// Profile3A is the implementation of the PDF/A-3A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile3A struct{ profile3 }

func _fbgfd(_acagg *_ae.CompliancePdfReader) ViolatedRule {
	_ecbe := _acagg.ParserMetadata()
	if _ecbe.HasInvalidSeparationAfterXRef() {
		return _eef("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _fc
}
func _caddc(_cfedg *_geb.PdfObjectStream, _eaabe map[*_geb.PdfObjectStream][]byte, _baec map[*_geb.PdfObjectStream]*_gba.CMap) (*_gba.CMap, error) {
	_cdad, _eaafg := _baec[_cfedg]
	if !_eaafg {
		var _cefeb error
		_edab, _daedc := _eaabe[_cfedg]
		if !_daedc {
			_edab, _cefeb = _geb.DecodeStream(_cfedg)
			if _cefeb != nil {
				_gd.Log.Debug("\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _cefeb)
				return nil, _cefeb
			}
			_eaabe[_cfedg] = _edab
		}
		_cdad, _cefeb = _gba.LoadCmapFromData(_edab, false)
		if _cefeb != nil {
			return nil, _cefeb
		}
		_baec[_cfedg] = _cdad
	}
	return _cdad, nil
}
func _eaabc(_dfcc *_ae.PdfFont, _bgdef *_geb.PdfObjectDictionary) ViolatedRule {
	const (
		_decb = "\u0036.\u0033\u002e\u0037\u002d\u0033"
		_gcgc = "\u0046\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0027\u0020\u0022\u0063\u006d\u0061\u0070\u0022\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0073\u0079\u006d\u0062o\u006c\u0069c\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0065\u0078\u0061\u0063\u0074\u006cy\u0020\u006f\u006ee\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u002e"
	)
	var _fcfa string
	if _dcaf, _bcbd := _geb.GetName(_bgdef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bcbd {
		_fcfa = _dcaf.String()
	}
	if _fcfa != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _fc
	}
	_gcdb := _dfcc.FontDescriptor()
	_ceeae, _aacc := _geb.GetIntVal(_gcdb.Flags)
	if !_aacc {
		_gd.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _eef(_decb, _gcgc)
	}
	_bfce := (uint32(_ceeae) >> 3) != 0
	if !_bfce {
		return _fc
	}
	return _fc
}
func _cebe(_facde *_ae.CompliancePdfReader) (_gaba []ViolatedRule) {
	_ebbdca := func(_gddca *_geb.PdfObjectDictionary, _affec *[]string, _fcfcc *[]ViolatedRule) error {
		_eeaa := _gddca.Get("\u004e\u0061\u006d\u0065")
		if _eeaa == nil || len(_eeaa.String()) == 0 {
			*_fcfcc = append(*_fcfcc, _eef("\u0036\u002e\u0039-\u0031", "\u0045\u0061\u0063\u0068\u0020o\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0063\u006fn\u0066\u0069\u0067\u0075r\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002e"))
		}
		for _, _eafb := range *_affec {
			if _eafb == _eeaa.String() {
				*_fcfcc = append(*_fcfcc, _eef("\u0036\u002e\u0039-\u0032", "\u0045\u0061\u0063\u0068\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0066\u0069\u0067\u0075\u0072a\u0074\u0069\u006fn\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020N\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002c w\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075ni\u0071\u0075\u0065 \u0061\u006d\u006f\u006e\u0067\u0073\u0074\u0020\u0061\u006c\u006c\u0020o\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006fn\u0074\u0065\u006e\u0074 \u0063\u006f\u006e\u0066\u0069\u0067u\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061\u0072\u0069\u0065\u0073\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0074\u0068e\u0020\u0050\u0044\u0046\u002fA\u002d\u0032\u0020\u0066\u0069l\u0065\u002e"))
			} else {
				*_affec = append(*_affec, _eeaa.String())
			}
		}
		if _gddca.Get("\u0041\u0053") != nil {
			*_fcfcc = append(*_fcfcc, _eef("\u0036\u002e\u0039-\u0034", "Th\u0065\u0020\u0041\u0053\u0020\u006b\u0065y \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0020\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0063\u006fn\u0066\u0069\u0067\u0075\u0072\u0061\u0074\u0069\u006fn\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
		}
		return nil
	}
	_gebe, _gcebe := _debed(_facde)
	if !_gcebe {
		return _gaba
	}
	_cdggc, _gcebe := _geb.GetDict(_gebe.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_gcebe {
		return _gaba
	}
	var _ddde []string
	_bgfega, _gcebe := _geb.GetDict(_cdggc.Get("\u0044"))
	if _gcebe {
		_ebbdca(_bgfega, &_ddde, &_gaba)
	}
	_ebcd, _gcebe := _geb.GetArray(_cdggc.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _gcebe {
		for _ebdf := 0; _ebdf < _ebcd.Len(); _ebdf++ {
			_gffe, _ffbc := _geb.GetDict(_ebcd.Get(_ebdf))
			if !_ffbc {
				continue
			}
			_ebbdca(_gffe, &_ddde, &_gaba)
		}
	}
	return _gaba
}

// StandardName gets the name of the standard.
func (_fbabe *profile3) StandardName() string {
	return _d.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0033\u0025s", _fbabe._egbf._ea)
}

// Profile2A is the implementation of the PDF/A-2A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2A struct{ profile2 }

var _ Profile = (*Profile3A)(nil)

func _caec(_cdge *_ae.CompliancePdfReader) ViolatedRule {
	if _cdge.ParserMetadata().HeaderPosition() != 0 {
		return _eef("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	if _cdge.PdfVersion().Major != 1 {
		return _eef("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	if _cdge.PdfVersion().Minor < 0 || _cdge.PdfVersion().Minor > 7 {
		return _eef("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	return _fc
}
func (_aef *documentImages) hasUncalibratedImages() bool { return _aef._gdf || _aef._cag || _aef._cfb }
