// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson

import (
	_b "encoding/json"
	_c "io"
	_e "os"

	_cc "github.com/bamzi/pdfext/common"
	_ca "github.com/bamzi/pdfext/core"
	_g "github.com/bamzi/pdfext/model"
)

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile(filePath string) (*FieldData, error) {
	_ef, _d := _e.Open(filePath)
	if _d != nil {
		return nil, _d
	}
	defer _ef.Close()
	return LoadFromJSON(_ef)
}

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON(r _c.Reader) (*FieldData, error) {
	var _fa FieldData
	_ea := _b.NewDecoder(r).Decode(&_fa._cf)
	if _ea != nil {
		return nil, _ea
	}
	return &_fa, nil
}

type fieldValue struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	ImageValue *_g.Image `json:"-"`

	// Options lists allowed values if present.
	Options []string `json:"options,omitempty"`
}

// FieldValues implements model.FieldValueProvider interface.
func (_cg *FieldData) FieldValues() (map[string]_ca.PdfObject, error) {
	_fg := make(map[string]_ca.PdfObject)
	for _, _ab := range _cg._cf {
		if len(_ab.Value) > 0 {
			_fg[_ab.Name] = _ca.MakeString(_ab.Value)
		}
	}
	return _fg, nil
}

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile(filePath string) (*FieldData, error) {
	_egc, _dgb := _e.Open(filePath)
	if _dgb != nil {
		return nil, _dgb
	}
	defer _egc.Close()
	return LoadFromPDF(_egc)
}

// FieldImageValues implements model.FieldImageProvider interface.
func (_ad *FieldData) FieldImageValues() (map[string]*_g.Image, error) {
	_cbc := make(map[string]*_g.Image)
	for _, _ga := range _ad._cf {
		if _ga.ImageValue != nil {
			_cbc[_ga.Name] = _ga.ImageValue
		}
	}
	return _cbc, nil
}

// SetImage assign model.Image to a specific field identified by fieldName.
func (_bc *FieldData) SetImage(fieldName string, img *_g.Image, opt []string) error {
	_ed := fieldValue{Name: fieldName, ImageValue: img, Options: opt}
	_bc._cf = append(_bc._cf, _ed)
	return nil
}

// JSON returns the field data as a string in JSON format.
func (_dfd FieldData) JSON() (string, error) {
	_fac, _fd := _b.MarshalIndent(_dfd._cf, "", "\u0020\u0020\u0020\u0020")
	return string(_fac), _fd
}

// FieldData represents form field data loaded from JSON file.
type FieldData struct{ _cf []fieldValue }

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF(rs _c.ReadSeeker) (*FieldData, error) {
	_gd, _gb := _g.NewPdfReader(rs)
	if _gb != nil {
		return nil, _gb
	}
	if _gd.AcroForm == nil {
		return nil, nil
	}
	var _efc []fieldValue
	_eg := _gd.AcroForm.AllFields()
	for _, _cb := range _eg {
		var _gcd []string
		_dc := make(map[string]struct{})
		_eb, _fb := _cb.FullName()
		if _fb != nil {
			return nil, _fb
		}
		if _cbg, _de := _cb.V.(*_ca.PdfObjectString); _de {
			_efc = append(_efc, fieldValue{Name: _eb, Value: _cbg.Decoded()})
			continue
		}
		var _fc string
		for _, _deb := range _cb.Annotations {
			_a, _bfd := _ca.GetName(_deb.AS)
			if _bfd {
				_fc = _a.String()
			}
			_gbb, _ce := _ca.GetDict(_deb.AP)
			if !_ce {
				continue
			}
			_cd, _ := _ca.GetDict(_gbb.Get("\u004e"))
			for _, _ge := range _cd.Keys() {
				_dg := _ge.String()
				if _, _dcd := _dc[_dg]; !_dcd {
					_gcd = append(_gcd, _dg)
					_dc[_dg] = struct{}{}
				}
			}
			_cbb, _ := _ca.GetDict(_gbb.Get("\u0044"))
			for _, _df := range _cbb.Keys() {
				_egg := _df.String()
				if _, _afb := _dc[_egg]; !_afb {
					_gcd = append(_gcd, _egg)
					_dc[_egg] = struct{}{}
				}
			}
		}
		_afba := fieldValue{Name: _eb, Value: _fc, Options: _gcd}
		_efc = append(_efc, _afba)
	}
	_cda := FieldData{_cf: _efc}
	return &_cda, nil
}

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_gaa *FieldData) SetImageFromFile(fieldName string, imagePath string, opt []string) error {
	_ag, _abb := _e.Open(imagePath)
	if _abb != nil {
		return _abb
	}
	defer _ag.Close()
	_fge, _abb := _g.ImageHandling.Read(_ag)
	if _abb != nil {
		_cc.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _abb)
		return _abb
	}
	return _gaa.SetImage(fieldName, _fge, opt)
}
