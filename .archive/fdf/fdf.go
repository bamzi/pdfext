// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf

import (
	_bg "bufio"
	_ea "bytes"
	_b "encoding/hex"
	_dd "errors"
	_eb "fmt"
	_cc "io"
	_c "os"
	_e "regexp"
	_a "sort"
	_gg "strconv"
	_g "strings"

	_dc "github.com/bamzi/pdfext/common"
	_cb "github.com/bamzi/pdfext/core"
)

func (_ccbf *fdfParser) parseName() (_cb.PdfObjectName, error) {
	var _fcb _ea.Buffer
	_ddf := false
	for {
		_eeb, _abe := _ccbf._fae.Peek(1)
		if _abe == _cc.EOF {
			break
		}
		if _abe != nil {
			return _cb.PdfObjectName(_fcb.String()), _abe
		}
		if !_ddf {
			if _eeb[0] == '/' {
				_ddf = true
				_ccbf._fae.ReadByte()
			} else if _eeb[0] == '%' {
				_ccbf.readComment()
				_ccbf.skipSpaces()
			} else {
				_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _eeb, _eeb)
				return _cb.PdfObjectName(_fcb.String()), _eb.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _eeb[0])
			}
		} else {
			if _cb.IsWhiteSpace(_eeb[0]) {
				break
			} else if (_eeb[0] == '/') || (_eeb[0] == '[') || (_eeb[0] == '(') || (_eeb[0] == ']') || (_eeb[0] == '<') || (_eeb[0] == '>') {
				break
			} else if _eeb[0] == '#' {
				_gfd, _cca := _ccbf._fae.Peek(3)
				if _cca != nil {
					return _cb.PdfObjectName(_fcb.String()), _cca
				}
				_ccbf._fae.Discard(3)
				_ded, _cca := _b.DecodeString(string(_gfd[1:3]))
				if _cca != nil {
					return _cb.PdfObjectName(_fcb.String()), _cca
				}
				_fcb.Write(_ded)
			} else {
				_fd, _ := _ccbf._fae.ReadByte()
				_fcb.WriteByte(_fd)
			}
		}
	}
	return _cb.PdfObjectName(_fcb.String()), nil
}
func (_gdce *fdfParser) parseObject() (_cb.PdfObject, error) {
	_dc.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_gdce.skipSpaces()
	for {
		_gda, _dgga := _gdce._fae.Peek(2)
		if _dgga != nil {
			return nil, _dgga
		}
		_dc.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gda))
		if _gda[0] == '/' {
			_fde, _aed := _gdce.parseName()
			_dc.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _fde)
			return &_fde, _aed
		} else if _gda[0] == '(' {
			_dc.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			return _gdce.parseString()
		} else if _gda[0] == '[' {
			_dc.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			return _gdce.parseArray()
		} else if (_gda[0] == '<') && (_gda[1] == '<') {
			_dc.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			return _gdce.parseDict()
		} else if _gda[0] == '<' {
			_dc.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			return _gdce.parseHexString()
		} else if _gda[0] == '%' {
			_gdce.readComment()
			_gdce.skipSpaces()
		} else {
			_dc.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_gda, _ = _gdce._fae.Peek(15)
			_fgb := string(_gda)
			_dc.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _fgb)
			if (len(_fgb) > 3) && (_fgb[:4] == "\u006e\u0075\u006c\u006c") {
				_cgd, _fff := _gdce.parseNull()
				return &_cgd, _fff
			} else if (len(_fgb) > 4) && (_fgb[:5] == "\u0066\u0061\u006cs\u0065") {
				_af, _gdg := _gdce.parseBool()
				return &_af, _gdg
			} else if (len(_fgb) > 3) && (_fgb[:4] == "\u0074\u0072\u0075\u0065") {
				_eac, _dgbb := _gdce.parseBool()
				return &_eac, _dgbb
			}
			_bcd := _ccc.FindStringSubmatch(_fgb)
			if len(_bcd) > 1 {
				_gda, _ = _gdce._fae.ReadBytes('R')
				_dc.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_gda[:]))
				_cbd, _dee := _efb(string(_gda))
				return &_cbd, _dee
			}
			_dde := _fed.FindStringSubmatch(_fgb)
			if len(_dde) > 1 {
				_dc.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				return _gdce.parseNumber()
			}
			_dde = _gga.FindStringSubmatch(_fgb)
			if len(_dde) > 1 {
				_dc.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_dc.Log.Trace("\u0025\u0020\u0073", _dde)
				return _gdce.parseNumber()
			}
			_dc.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _fgb)
			return nil, _dd.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}

var _cde = _e.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

// Load loads FDF form data from `r`.
func Load(r _cc.ReadSeeker) (*Data, error) {
	_dce, _ae := _fee(r)
	if _ae != nil {
		return nil, _ae
	}
	_gc, _ae := _dce.Root()
	if _ae != nil {
		return nil, _ae
	}
	_f, _ag := _cb.GetArray(_gc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_ag {
		return nil, _dd.New("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	return &Data{_bf: _f, _ddb: _gc}, nil
}
func (_edc *fdfParser) readTextLine() (string, error) {
	var _ga _ea.Buffer
	for {
		_fc, _gf := _edc._fae.Peek(1)
		if _gf != nil {
			_dc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gf.Error())
			return _ga.String(), _gf
		}
		if (_fc[0] != '\r') && (_fc[0] != '\n') {
			_bc, _ := _edc._fae.ReadByte()
			_ga.WriteByte(_bc)
		} else {
			break
		}
	}
	return _ga.String(), nil
}

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data) FieldValues() (map[string]_cb.PdfObject, error) {
	_aeg, _ccd := fdf.FieldDictionaries()
	if _ccd != nil {
		return nil, _ccd
	}
	var _ad []string
	for _eaa := range _aeg {
		_ad = append(_ad, _eaa)
	}
	_a.Strings(_ad)
	_cd := map[string]_cb.PdfObject{}
	for _, _ec := range _ad {
		_gd := _aeg[_ec]
		_aa := _cb.TraceToDirectObject(_gd.Get("\u0056"))
		_cd[_ec] = _aa
	}
	return _cd, nil
}

type fdfParser struct {
	_bfb int
	_fag int
	_cf  map[int64]_cb.PdfObject
	_ee  _cc.ReadSeeker
	_fae *_bg.Reader
	_gcc int64
	_ebc *_cb.PdfObjectDictionary
}

func (_cfg *fdfParser) parseFdfVersion() (int, int, error) {
	_cfg._ee.Seek(0, _cc.SeekStart)
	_ccg := 20
	_ddc := make([]byte, _ccg)
	_cfg._ee.Read(_ddc)
	_def := _ca.FindStringSubmatch(string(_ddc))
	if len(_def) < 3 {
		_deb, _fcba, _fcg := _cfg.seekFdfVersionTopDown()
		if _fcg != nil {
			_dc.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _fcg
		}
		return _deb, _fcba, nil
	}
	_aad, _gdcf := _gg.Atoi(_def[1])
	if _gdcf != nil {
		return 0, 0, _gdcf
	}
	_fagc, _gdcf := _gg.Atoi(_def[2])
	if _gdcf != nil {
		return 0, 0, _gdcf
	}
	_dc.Log.Debug("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _aad, _fagc)
	return _aad, _fagc, nil
}
func (_cdd *fdfParser) readAtLeast(_ed []byte, _ccb int) (int, error) {
	_ebg := _ccb
	_fa := 0
	_feg := 0
	for _ebg > 0 {
		_ab, _adc := _cdd._fae.Read(_ed[_fa:])
		if _adc != nil {
			_dc.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _ab, _feg, _adc.Error())
			return _fa, _dd.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_feg++
		_fa += _ab
		_ebg -= _ab
	}
	return _fa, nil
}
func (_ac *fdfParser) getFileOffset() int64 {
	_gcd, _ := _ac._ee.Seek(0, _cc.SeekCurrent)
	_gcd -= int64(_ac._fae.Buffered())
	return _gcd
}

var _ccc = _e.MustCompile("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052")
var _fed = _e.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_gba *fdfParser) parseNumber() (_cb.PdfObject, error) { return _cb.ParseNumber(_gba._fae) }
func (_dcb *fdfParser) skipComments() error {
	if _, _de := _dcb.skipSpaces(); _de != nil {
		return _de
	}
	_df := true
	for {
		_ggc, _aaa := _dcb._fae.Peek(1)
		if _aaa != nil {
			_dc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _aaa.Error())
			return _aaa
		}
		if _df && _ggc[0] != '%' {
			return nil
		}
		_df = false
		if (_ggc[0] != '\r') && (_ggc[0] != '\n') {
			_dcb._fae.ReadByte()
		} else {
			break
		}
	}
	return _dcb.skipComments()
}

var _age = _e.MustCompile("\u0025\u0025\u0045O\u0046")

func (_faeg *fdfParser) parseIndirectObject() (_cb.PdfObject, error) {
	_aaab := _cb.PdfIndirectObject{}
	_dc.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_cda, _afc := _faeg._fae.Peek(20)
	if _afc != nil {
		_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
		return &_aaab, _afc
	}
	_dc.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_cda))
	_ecg := _cde.FindStringSubmatchIndex(string(_cda))
	if len(_ecg) < 6 {
		_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_cda))
		return &_aaab, _dd.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_faeg._fae.Discard(_ecg[0])
	_dc.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _ecg)
	_dcfb := _ecg[1] - _ecg[0]
	_acd := make([]byte, _dcfb)
	_, _afc = _faeg.readAtLeast(_acd, _dcfb)
	if _afc != nil {
		_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _afc)
		return nil, _afc
	}
	_dc.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _acd)
	_fega := _cde.FindStringSubmatch(string(_acd))
	if len(_fega) < 3 {
		_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_acd))
		return &_aaab, _dd.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_dda, _ := _gg.Atoi(_fega[1])
	_add, _ := _gg.Atoi(_fega[2])
	_aaab.ObjectNumber = int64(_dda)
	_aaab.GenerationNumber = int64(_add)
	for {
		_ddg, _eef := _faeg._fae.Peek(2)
		if _eef != nil {
			return &_aaab, _eef
		}
		_dc.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_ddg), string(_ddg))
		if _cb.IsWhiteSpace(_ddg[0]) {
			_faeg.skipSpaces()
		} else if _ddg[0] == '%' {
			_faeg.skipComments()
		} else if (_ddg[0] == '<') && (_ddg[1] == '<') {
			_dc.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_aaab.PdfObject, _eef = _faeg.parseDict()
			_dc.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _eef)
			if _eef != nil {
				return &_aaab, _eef
			}
			_dc.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_ddg[0] == '/') || (_ddg[0] == '(') || (_ddg[0] == '[') || (_ddg[0] == '<') {
			_aaab.PdfObject, _eef = _faeg.parseObject()
			if _eef != nil {
				return &_aaab, _eef
			}
			_dc.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else {
			if _ddg[0] == 'e' {
				_gdfc, _bddd := _faeg.readTextLine()
				if _bddd != nil {
					return nil, _bddd
				}
				if len(_gdfc) >= 6 && _gdfc[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _ddg[0] == 's' {
				_ddg, _ = _faeg._fae.Peek(10)
				if string(_ddg[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_cdb := 6
					if len(_ddg) > 6 {
						if _cb.IsWhiteSpace(_ddg[_cdb]) && _ddg[_cdb] != '\r' && _ddg[_cdb] != '\n' {
							_dc.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_cdb++
						}
						if _ddg[_cdb] == '\r' {
							_cdb++
							if _ddg[_cdb] == '\n' {
								_cdb++
							}
						} else if _ddg[_cdb] == '\n' {
							_cdb++
						}
					}
					_faeg._fae.Discard(_cdb)
					_cbgf, _gfge := _aaab.PdfObject.(*_cb.PdfObjectDictionary)
					if !_gfge {
						return nil, _dd.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_dc.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _cbgf)
					_fab, _cfcc := _cbgf.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*_cb.PdfObjectInteger)
					if !_cfcc {
						return nil, _dd.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_bdg := *_fab
					if _bdg < 0 {
						return nil, _dd.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					if int64(_bdg) > _faeg._gcc {
						_dc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _dd.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_dgf := make([]byte, _bdg)
					_, _eef = _faeg.readAtLeast(_dgf, int(_bdg))
					if _eef != nil {
						_dc.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_dgf), _dgf)
						_dc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eef)
						return nil, _eef
					}
					_cdg := _cb.PdfObjectStream{}
					_cdg.Stream = _dgf
					_cdg.PdfObjectDictionary = _aaab.PdfObject.(*_cb.PdfObjectDictionary)
					_cdg.ObjectNumber = _aaab.ObjectNumber
					_cdg.GenerationNumber = _aaab.GenerationNumber
					_faeg.skipSpaces()
					_faeg._fae.Discard(9)
					_faeg.skipSpaces()
					return &_cdg, nil
				}
			}
			_aaab.PdfObject, _eef = _faeg.parseObject()
			return &_aaab, _eef
		}
	}
	_dc.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_aaab, nil
}
func (_ecdb *fdfParser) parse() error {
	_ecdb._ee.Seek(0, _cc.SeekStart)
	_ecdb._fae = _bg.NewReader(_ecdb._ee)
	for {
		_ecdb.skipComments()
		_feec, _eeeb := _ecdb._fae.Peek(20)
		if _eeeb != nil {
			_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return _eeeb
		}
		if _g.HasPrefix(string(_feec), "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_ecdb._fae.Discard(7)
			_ecdb.skipSpaces()
			_ecdb.skipComments()
			_fbb, _ := _ecdb.parseDict()
			_ecdb._ebc = _fbb
			break
		}
		_efg := _cde.FindStringSubmatchIndex(string(_feec))
		if len(_efg) < 6 {
			_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_feec))
			return _dd.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
		}
		_bec, _eeeb := _ecdb.parseIndirectObject()
		if _eeeb != nil {
			return _eeeb
		}
		switch _fea := _bec.(type) {
		case *_cb.PdfIndirectObject:
			_ecdb._cf[_fea.ObjectNumber] = _fea
		case *_cb.PdfObjectStream:
			_ecdb._cf[_fea.ObjectNumber] = _fea
		default:
			return _dd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	}
	return nil
}
func (_cbc *fdfParser) readComment() (string, error) {
	var _beg _ea.Buffer
	_, _ade := _cbc.skipSpaces()
	if _ade != nil {
		return _beg.String(), _ade
	}
	_cfe := true
	for {
		_eg, _da := _cbc._fae.Peek(1)
		if _da != nil {
			_dc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _da.Error())
			return _beg.String(), _da
		}
		if _cfe && _eg[0] != '%' {
			return _beg.String(), _dd.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_cfe = false
		if (_eg[0] != '\r') && (_eg[0] != '\n') {
			_egg, _ := _cbc._fae.ReadByte()
			_beg.WriteByte(_egg)
		} else {
			break
		}
	}
	return _beg.String(), nil
}
func (_dfg *fdfParser) parseHexString() (*_cb.PdfObjectString, error) {
	_dfg._fae.ReadByte()
	var _aee _ea.Buffer
	for {
		_ce, _ceg := _dfg._fae.Peek(1)
		if _ceg != nil {
			return _cb.MakeHexString(""), _ceg
		}
		if _ce[0] == '>' {
			_dfg._fae.ReadByte()
			break
		}
		_ffc, _ := _dfg._fae.ReadByte()
		if !_cb.IsWhiteSpace(_ffc) {
			_aee.WriteByte(_ffc)
		}
	}
	if _aee.Len()%2 == 1 {
		_aee.WriteRune('0')
	}
	_dcee, _gdd := _b.DecodeString(_aee.String())
	if _gdd != nil {
		_dc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067", _aee.String())
		return _cb.MakeHexString(""), nil
	}
	return _cb.MakeHexString(string(_dcee)), nil
}
func (_ffe *fdfParser) trace(_dab _cb.PdfObject) _cb.PdfObject {
	switch _gbe := _dab.(type) {
	case *_cb.PdfObjectReference:
		_agc, _bee := _ffe._cf[_gbe.ObjectNumber].(*_cb.PdfIndirectObject)
		if _bee {
			return _agc.PdfObject
		}
		_dc.Log.Debug("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		return nil
	case *_cb.PdfIndirectObject:
		return _gbe.PdfObject
	}
	return _dab
}

var _gga = _e.MustCompile("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029")

func (_dgg *fdfParser) skipSpaces() (int, error) {
	_bfe := 0
	for {
		_bdb, _ba := _dgg._fae.ReadByte()
		if _ba != nil {
			return 0, _ba
		}
		if _cb.IsWhiteSpace(_bdb) {
			_bfe++
		} else {
			_dgg._fae.UnreadByte()
			break
		}
	}
	return _bfe, nil
}
func (_cg *fdfParser) parseArray() (*_cb.PdfObjectArray, error) {
	_dgb := _cb.MakeArray()
	_cg._fae.ReadByte()
	for {
		_cg.skipSpaces()
		_aac, _abf := _cg._fae.Peek(1)
		if _abf != nil {
			return _dgb, _abf
		}
		if _aac[0] == ']' {
			_cg._fae.ReadByte()
			break
		}
		_caa, _abf := _cg.parseObject()
		if _abf != nil {
			return _dgb, _abf
		}
		_dgb.Append(_caa)
	}
	return _dgb, nil
}
func (_bdbc *fdfParser) parseBool() (_cb.PdfObjectBool, error) {
	_fgf, _bddf := _bdbc._fae.Peek(4)
	if _bddf != nil {
		return _cb.PdfObjectBool(false), _bddf
	}
	if (len(_fgf) >= 4) && (string(_fgf[:4]) == "\u0074\u0072\u0075\u0065") {
		_bdbc._fae.Discard(4)
		return _cb.PdfObjectBool(true), nil
	}
	_fgf, _bddf = _bdbc._fae.Peek(5)
	if _bddf != nil {
		return _cb.PdfObjectBool(false), _bddf
	}
	if (len(_fgf) >= 5) && (string(_fgf[:5]) == "\u0066\u0061\u006cs\u0065") {
		_bdbc._fae.Discard(5)
		return _cb.PdfObjectBool(false), nil
	}
	return _cb.PdfObjectBool(false), _dd.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}
func (_gb *fdfParser) setFileOffset(_fb int64) {
	_gb._ee.Seek(_fb, _cc.SeekStart)
	_gb._fae = _bg.NewReader(_gb._ee)
}
func (_gac *fdfParser) seekFdfVersionTopDown() (int, int, error) {
	_gac._ee.Seek(0, _cc.SeekStart)
	_gac._fae = _bg.NewReader(_gac._ee)
	_feb := 20
	_cae := make([]byte, _feb)
	for {
		_ceb, _fcgc := _gac._fae.ReadByte()
		if _fcgc != nil {
			if _fcgc == _cc.EOF {
				break
			} else {
				return 0, 0, _fcgc
			}
		}
		if _cb.IsDecimalDigit(_ceb) && _cae[_feb-1] == '.' && _cb.IsDecimalDigit(_cae[_feb-2]) && _cae[_feb-3] == '-' && _cae[_feb-4] == 'F' && _cae[_feb-5] == 'D' && _cae[_feb-6] == 'P' {
			_abd := int(_cae[_feb-2] - '0')
			_ebd := int(_ceb - '0')
			return _abd, _ebd, nil
		}
		_cae = append(_cae[1:_feb], _ceb)
	}
	return 0, 0, _dd.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// Root returns the Root of the FDF document.
func (_efc *fdfParser) Root() (*_cb.PdfObjectDictionary, error) {
	if _efc._ebc != nil {
		if _fce, _bgf := _efc.trace(_efc._ebc.Get("\u0052\u006f\u006f\u0074")).(*_cb.PdfObjectDictionary); _bgf {
			if _efbg, _dcea := _efc.trace(_fce.Get("\u0046\u0044\u0046")).(*_cb.PdfObjectDictionary); _dcea {
				return _efbg, nil
			}
		}
	}
	var _ede []int64
	for _eegd := range _efc._cf {
		_ede = append(_ede, _eegd)
	}
	_a.Slice(_ede, func(_dag, _adeb int) bool { return _ede[_dag] < _ede[_adeb] })
	for _, _adca := range _ede {
		_fcd := _efc._cf[_adca]
		if _cgb, _bca := _efc.trace(_fcd).(*_cb.PdfObjectDictionary); _bca {
			if _egc, _aacg := _efc.trace(_cgb.Get("\u0046\u0044\u0046")).(*_cb.PdfObjectDictionary); _aacg {
				return _egc, nil
			}
		}
	}
	return nil, _dd.New("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data) FieldDictionaries() (map[string]*_cb.PdfObjectDictionary, error) {
	_gcg := map[string]*_cb.PdfObjectDictionary{}
	for _ggg := 0; _ggg < fdf._bf.Len(); _ggg++ {
		_ef, _bd := _cb.GetDict(fdf._bf.Get(_ggg))
		if _bd {
			_dcf, _ := _cb.GetString(_ef.Get("\u0054"))
			if _dcf != nil {
				_gcg[_dcf.Str()] = _ef
			}
		}
	}
	return _gcg, nil
}

var _ca = _e.MustCompile("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func _efb(_cef string) (_cb.PdfObjectReference, error) {
	_gdc := _cb.PdfObjectReference{}
	_agd := _ccc.FindStringSubmatch(_cef)
	if len(_agd) < 3 {
		_dc.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _gdc, _dd.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_fcc, _egd := _gg.Atoi(_agd[1])
	if _egd != nil {
		_dc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030", _agd[1])
		return _gdc, nil
	}
	_gdc.ObjectNumber = int64(_fcc)
	_daa, _egd := _gg.Atoi(_agd[2])
	if _egd != nil {
		_dc.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030", _agd[2])
		return _gdc, nil
	}
	_gdc.GenerationNumber = int64(_daa)
	return _gdc, nil
}
func (_facd *fdfParser) parseDict() (*_cb.PdfObjectDictionary, error) {
	_dc.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_abfe := _cb.MakeDict()
	_db, _ := _facd._fae.ReadByte()
	if _db != '<' {
		return nil, _dd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_db, _ = _facd._fae.ReadByte()
	if _db != '<' {
		return nil, _dd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_facd.skipSpaces()
		_facd.skipComments()
		_gfg, _gfc := _facd._fae.Peek(2)
		if _gfc != nil {
			return nil, _gfc
		}
		_dc.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_gfg), string(_gfg))
		if (_gfg[0] == '>') && (_gfg[1] == '>') {
			_dc.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_facd._fae.ReadByte()
			_facd._fae.ReadByte()
			break
		}
		_dc.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_gfb, _gfc := _facd.parseName()
		_dc.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _gfb)
		if _gfc != nil {
			_dc.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _gfc)
			return nil, _gfc
		}
		if len(_gfb) > 4 && _gfb[len(_gfb)-4:] == "\u006e\u0075\u006c\u006c" {
			_fgfc := _gfb[0 : len(_gfb)-4]
			_dc.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _gfb)
			_dc.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _fgfc)
			_facd.skipSpaces()
			_abc, _ := _facd._fae.Peek(1)
			if _abc[0] == '/' {
				_abfe.Set(_fgfc, _cb.MakeNull())
				continue
			}
		}
		_facd.skipSpaces()
		_agef, _gfc := _facd.parseObject()
		if _gfc != nil {
			return nil, _gfc
		}
		_abfe.Set(_gfb, _agef)
		_dc.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _gfb, _agef.String())
	}
	_dc.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _abfe, nil
}

// Data represents forms data format (FDF) file data.
type Data struct {
	_ddb *_cb.PdfObjectDictionary
	_bf  *_cb.PdfObjectArray
}

func (_fagb *fdfParser) seekToEOFMarker(_ddd int64) error {
	_gdf := int64(0)
	_ggaf := int64(1000)
	for _gdf < _ddd {
		if _ddd <= (_ggaf + _gdf) {
			_ggaf = _ddd - _gdf
		}
		_, _cbg := _fagb._ee.Seek(-_gdf-_ggaf, _cc.SeekEnd)
		if _cbg != nil {
			return _cbg
		}
		_dea := make([]byte, _ggaf)
		_fagb._ee.Read(_dea)
		_dc.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_dea))
		_gfa := _age.FindAllStringIndex(string(_dea), -1)
		if _gfa != nil {
			_dgbe := _gfa[len(_gfa)-1]
			_dc.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _gfa)
			_fagb._ee.Seek(-_gdf-_ggaf+int64(_dgbe[0]), _cc.SeekEnd)
			return nil
		}
		_dc.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_gdf += _ggaf
	}
	_dc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _dd.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func _fee(_cba _cc.ReadSeeker) (*fdfParser, error) {
	_cab := &fdfParser{}
	_cab._ee = _cba
	_cab._cf = map[int64]_cb.PdfObject{}
	_dfc, _cdae, _adg := _cab.parseFdfVersion()
	if _adg != nil {
		_dc.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _adg)
		return nil, _adg
	}
	_cab._bfb = _dfc
	_cab._fag = _cdae
	_adg = _cab.parse()
	return _cab, _adg
}
func _adfa(_dgbc string) (*fdfParser, error) {
	_fbd := fdfParser{}
	_ffge := []byte(_dgbc)
	_ccdg := _ea.NewReader(_ffge)
	_fbd._ee = _ccdg
	_fbd._cf = map[int64]_cb.PdfObject{}
	_cee := _bg.NewReader(_ccdg)
	_fbd._fae = _cee
	_fbd._gcc = int64(len(_dgbc))
	return &_fbd, _fbd.parse()
}
func (_fac *fdfParser) parseString() (*_cb.PdfObjectString, error) {
	_fac._fae.ReadByte()
	var _bab _ea.Buffer
	_acf := 1
	for {
		_eeg, _bdd := _fac._fae.Peek(1)
		if _bdd != nil {
			return _cb.MakeString(_bab.String()), _bdd
		}
		if _eeg[0] == '\\' {
			_fac._fae.ReadByte()
			_ece, _gdb := _fac._fae.ReadByte()
			if _gdb != nil {
				return _cb.MakeString(_bab.String()), _gdb
			}
			if _cb.IsOctalDigit(_ece) {
				_ecd, _bdbf := _fac._fae.Peek(2)
				if _bdbf != nil {
					return _cb.MakeString(_bab.String()), _bdbf
				}
				var _ffg []byte
				_ffg = append(_ffg, _ece)
				for _, _dfa := range _ecd {
					if _cb.IsOctalDigit(_dfa) {
						_ffg = append(_ffg, _dfa)
					} else {
						break
					}
				}
				_fac._fae.Discard(len(_ffg) - 1)
				_dc.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _ffg)
				_baf, _bdbf := _gg.ParseUint(string(_ffg), 8, 32)
				if _bdbf != nil {
					return _cb.MakeString(_bab.String()), _bdbf
				}
				_bab.WriteByte(byte(_baf))
				continue
			}
			switch _ece {
			case 'n':
				_bab.WriteRune('\n')
			case 'r':
				_bab.WriteRune('\r')
			case 't':
				_bab.WriteRune('\t')
			case 'b':
				_bab.WriteRune('\b')
			case 'f':
				_bab.WriteRune('\f')
			case '(':
				_bab.WriteRune('(')
			case ')':
				_bab.WriteRune(')')
			case '\\':
				_bab.WriteRune('\\')
			}
			continue
		} else if _eeg[0] == '(' {
			_acf++
		} else if _eeg[0] == ')' {
			_acf--
			if _acf == 0 {
				_fac._fae.ReadByte()
				break
			}
		}
		_fbc, _ := _fac._fae.ReadByte()
		_bab.WriteByte(_fbc)
	}
	return _cb.MakeString(_bab.String()), nil
}
func (_cbf *fdfParser) parseNull() (_cb.PdfObjectNull, error) {
	_, _aacf := _cbf._fae.Discard(4)
	return _cb.PdfObjectNull{}, _aacf
}

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath(fdfPath string) (*Data, error) {
	_bgc, _bge := _c.Open(fdfPath)
	if _bge != nil {
		return nil, _bge
	}
	defer _bgc.Close()
	return Load(_bgc)
}
