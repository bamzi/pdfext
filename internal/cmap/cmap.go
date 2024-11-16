package cmap

import (
	_ef "bufio"
	_ed "bytes"
	_ba "encoding/hex"
	_c "errors"
	_b "fmt"
	_eg "io"
	_bd "sort"
	_e "strconv"
	_bf "strings"
	_f "unicode/utf16"

	_cf "github.com/bamzi/pdfext/common"
	_bde "github.com/bamzi/pdfext/core"
	_cd "github.com/bamzi/pdfext/internal/cmap/bcmaps"
)

func _fcg() cmapDict { return cmapDict{Dict: map[string]cmapObject{}} }
func LoadCmapFromData(data []byte, isSimple bool) (*CMap, error) {
	_cf.Log.Trace("\u004c\u006fa\u0064\u0043\u006d\u0061\u0070\u0046\u0072\u006f\u006d\u0044\u0061\u0074\u0061\u003a\u0020\u0069\u0073\u0053\u0069\u006d\u0070\u006ce=\u0025\u0074", isSimple)
	cmap := _da(isSimple)
	cmap.cMapParser = _ffda(data)
	_ac := cmap.parse()
	if _ac != nil {
		return nil, _ac
	}
	if len(cmap._gd) == 0 {
		if cmap._cab != "" {
			return cmap, nil
		}
		_cf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0063\u006d\u0061p=\u0025\u0073", cmap)
	}
	cmap.computeInverseMappings()
	return cmap, nil
}

type cmapArray struct{ Array []cmapObject }

func (cmap *CMap) CIDToCharcode(cid CharCode) (CharCode, bool) {
	_abe, _age := cmap._afa[cid]
	return _abe, _age
}
func (cmap *CMap) StringToCID(s string) (CharCode, bool) {
	_dbe, _ff := cmap._ad[s]
	return _dbe, _ff
}
func (cmap *CMap) parseWMode() error {
	var _cdff int
	_bcde := false
	for _fcfa := 0; _fcfa < 3 && !_bcde; _fcfa++ {
		_ccd, _eec := cmap.parseObject()
		if _eec != nil {
			return _eec
		}
		switch _acc := _ccd.(type) {
		case cmapOperand:
			switch _acc.Operand {
			case "\u0064\u0065\u0066":
				_bcde = true
			default:
				_cf.Log.Error("\u0070\u0061\u0072\u0073\u0065\u0057\u004d\u006f\u0064\u0065:\u0020\u0073\u0074\u0061\u0074\u0065\u0020e\u0072\u0072\u006f\u0072\u002e\u0020\u006f\u003d\u0025\u0023\u0076", _ccd)
				return ErrBadCMap
			}
		case cmapInt:
			_cdff = int(_acc._cfgd)
		}
	}
	cmap._gf = integer{_fde: true, _gcg: _cdff}
	return nil
}
func (_daea *cMapParser) parseHexString() (cmapHexString, error) {
	_daea._deg.ReadByte()
	_dgbg := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	_bga := _ed.Buffer{}
	for {
		_daea.skipSpaces()
		_bdfd, _bgab := _daea._deg.Peek(1)
		if _bgab != nil {
			return cmapHexString{}, _bgab
		}
		if _bdfd[0] == '>' {
			_daea._deg.ReadByte()
			break
		}
		_adee, _ := _daea._deg.ReadByte()
		if _ed.IndexByte(_dgbg, _adee) >= 0 {
			_bga.WriteByte(_adee)
		}
	}
	if _bga.Len()%2 == 1 {
		_cf.Log.Debug("\u0070\u0061rs\u0065\u0048\u0065x\u0053\u0074\u0072\u0069ng:\u0020ap\u0070\u0065\u006e\u0064\u0069\u006e\u0067 '\u0030\u0027\u0020\u0074\u006f\u0020\u0025#\u0071", _bga.String())
		_bga.WriteByte('0')
	}
	_eaga := _bga.Len() / 2
	_dec, _ := _ba.DecodeString(_bga.String())
	return cmapHexString{_cceg: _eaga, _baeg: _dec}, nil
}

type cmapObject interface{}
type integer struct {
	_fde bool
	_gcg int
}

func (cmap *CMap) String() string {
	_bfee := cmap._cfb
	_ccg := []string{_b.Sprintf("\u006e\u0062\u0069\u0074\u0073\u003a\u0025\u0064", cmap._db), _b.Sprintf("\u0074y\u0070\u0065\u003a\u0025\u0064", cmap._gc)}
	if cmap._af != "" {
		_ccg = append(_ccg, _b.Sprintf("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0025\u0073", cmap._af))
	}
	if cmap._cab != "" {
		_ccg = append(_ccg, _b.Sprintf("u\u0073\u0065\u0063\u006d\u0061\u0070\u003a\u0025\u0023\u0071", cmap._cab))
	}
	_ccg = append(_ccg, _b.Sprintf("\u0073\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0025\u0073", _bfee.String()))
	if len(cmap._gd) > 0 {
		_ccg = append(_ccg, _b.Sprintf("\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u003a\u0025\u0064", len(cmap._gd)))
	}
	if len(cmap._bea) > 0 {
		_ccg = append(_ccg, _b.Sprintf("\u0063\u006fd\u0065\u0054\u006fU\u006e\u0069\u0063\u006f\u0064\u0065\u003a\u0025\u0064", len(cmap._bea)))
	}
	return _b.Sprintf("\u0043\u004d\u0041P\u007b\u0025\u0023\u0071\u0020\u0025\u0073\u007d", cmap._fb, _bf.Join(_ccg, "\u0020"))
}
func (cmap *CMap) parseBfrange() error {
	for {
		var _fdb CharCode
		_bbc, _cbd := cmap.parseObject()
		if _cbd != nil {
			if _cbd == _eg.EOF {
				break
			}
			return _cbd
		}
		switch _bffb := _bbc.(type) {
		case cmapOperand:
			if _bffb.Operand == _bddc {
				return nil
			}
			return _c.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_fdb = _cfcg(_bffb)
		default:
			return _c.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		var _abec CharCode
		_bbc, _cbd = cmap.parseObject()
		if _cbd != nil {
			if _cbd == _eg.EOF {
				break
			}
			return _cbd
		}
		switch _cdffg := _bbc.(type) {
		case cmapOperand:
			_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065\u0020\u0062\u0066r\u0061\u006e\u0067\u0065\u0020\u0074\u0072i\u0070\u006c\u0065\u0074")
			return ErrBadCMap
		case cmapHexString:
			_abec = _cfcg(_cdffg)
			if _abec > 0xffff {
				_abec = 0xffff
			}
		default:
			_cf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _bbc)
			return ErrBadCMap
		}
		_bbc, _cbd = cmap.parseObject()
		if _cbd != nil {
			if _cbd == _eg.EOF {
				break
			}
			return _cbd
		}
		switch _bba := _bbc.(type) {
		case cmapArray:
			if len(_bba.Array) != int(_abec-_fdb)+1 {
				_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0069\u0074\u0065\u006d\u0073\u0020\u0069\u006e\u0020a\u0072\u0072\u0061\u0079")
				return ErrBadCMap
			}
			for _cdgg := _fdb; _cdgg <= _abec; _cdgg++ {
				_gdcb := _bba.Array[_cdgg-_fdb]
				_ced, _agec := _gdcb.(cmapHexString)
				if !_agec {
					return _c.New("\u006e\u006f\u006e-h\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				}
				_gad := _abga(_ced)
				cmap._bea[_cdgg] = string(_gad)
			}
		case cmapHexString:
			_ffc := _abga(_bba)
			_gdgc := len(_ffc)
			for _eebf := _fdb; _eebf <= _abec; _eebf++ {
				cmap._bea[_eebf] = string(_ffc)
				if _gdgc > 0 {
					_ffc[_gdgc-1]++
				} else {
					_cf.Log.Debug("\u004e\u006f\u0020c\u006d\u0061\u0070\u0020\u0074\u0061\u0072\u0067\u0065\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0066\u006f\u0072\u0020\u0025\u0023\u0076", _eebf)
				}
				if _eebf == 1<<32-1 {
					break
				}
			}
		default:
			_cf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _bbc)
			return ErrBadCMap
		}
	}
	return nil
}
func (cmap *CMap) computeInverseMappings() {
	for _ee, _eee := range cmap._bad {
		if _cac, _ce := cmap._afa[_eee]; !_ce || (_ce && _cac > _ee) {
			cmap._afa[_eee] = _ee
		}
	}
	for _bada, _ceg := range cmap._bea {
		if _gae, _dd := cmap._ad[_ceg]; !_dd || (_dd && _gae > _bada) {
			cmap._ad[_ceg] = _bada
		}
	}
	_bd.Slice(cmap._gd, func(_fcf, _bce int) bool { return cmap._gd[_fcf].Low < cmap._gd[_bce].Low })
}
func IsPredefinedCMap(name string) bool { return _cd.AssetExists(name) }

type cmapOperand struct{ Operand string }

func (cmap *CMap) Name() string { return cmap._fb }
func _ceb(_cdf string) string {
	_fag := []rune(_cdf)
	_gaa := make([]string, len(_fag))
	for _bfc, _gde := range _fag {
		_gaa[_bfc] = _b.Sprintf("\u0025\u0030\u0034\u0078", _gde)
	}
	return _b.Sprintf("\u003c\u0025\u0073\u003e", _bf.Join(_gaa, ""))
}
func (cmap *CMap) Stream() (*_bde.PdfObjectStream, error) {
	if cmap._cge != nil {
		return cmap._cge, nil
	}
	_cde, _cfga := _bde.MakeStream(cmap.Bytes(), _bde.NewFlateEncoder())
	if _cfga != nil {
		return nil, _cfga
	}
	cmap._cge = _cde
	return cmap._cge, nil
}
func (cmap *CMap) matchCode(_beg []byte) (_aba CharCode, _bdf int, _ddb bool) {
	for _ddff := 0; _ddff < _dc; _ddff++ {
		if _ddff < len(_beg) {
			_aba = _aba<<8 | CharCode(_beg[_ddff])
			_bdf++
		}
		_ddb = cmap.inCodespace(_aba, _ddff+1)
		if _ddb {
			return _aba, _bdf, true
		}
	}
	_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063o\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0020m\u0061t\u0063\u0068\u0065\u0073\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d=\u0025\u0023\u0071\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", _beg, string(_beg), cmap)
	return 0, 0, false
}
func (_bfd *CIDSystemInfo) String() string {
	return _b.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0025\u0030\u0033\u0064", _bfd.Registry, _bfd.Ordering, _bfd.Supplement)
}
func (cmap *CMap) inCodespace(_bff CharCode, _bbe int) bool {
	for _, _cdd := range cmap._gd {
		if _cdd.Low <= _bff && _bff <= _cdd.High && _bbe == _cdd.NumBytes {
			return true
		}
	}
	return false
}

type Codespace struct {
	NumBytes int
	Low      CharCode
	High     CharCode
}
type fbRange struct {
	_g   CharCode
	_cg  CharCode
	_cgb string
}

func (cmap *CMap) CharcodeBytesToUnicode(data []byte) (string, int) {
	_bee, _bae := cmap.BytesToCharcodes(data)
	if !_bae {
		_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065s\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069n\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0064\u0061\u0074\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078]\u0020\u0063\u006d\u0061\u0070=\u0025\u0073", data, cmap)
		return "", 0
	}
	_gdg := make([]string, len(_bee))
	var _afg []CharCode
	for _eed, _fa := range _bee {
		_dga, _fg := cmap._bea[_fa]
		if !_fg {
			_afg = append(_afg, _fa)
			_dga = MissingCodeString
		}
		_gdg[_eed] = _dga
	}
	_ge := _bf.Join(_gdg, "")
	if len(_afg) > 0 {
		_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0068\u0061\u0072c\u006f\u0064\u0065\u0042y\u0074\u0065\u0073\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070\u002e\u000a"+"\u0009d\u0061t\u0061\u003d\u005b\u0025\u00200\u0032\u0078]\u003d\u0025\u0023\u0071\u000a"+"\u0009\u0063h\u0061\u0072\u0063o\u0064\u0065\u0073\u003d\u0025\u0030\u0032\u0078\u000a"+"\u0009\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003d\u0025\u0064\u0020%\u0030\u0032\u0078\u000a"+"\u0009\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u003d`\u0025\u0073\u0060\u000a"+"\u0009\u0063\u006d\u0061\u0070\u003d\u0025\u0073", data, string(data), _bee, len(_afg), _afg, _ge, cmap)
	}
	return _ge, len(_afg)
}
func (cmap *CMap) parseName() error {
	_aa := ""
	_fbe := false
	for _bfb := 0; _bfb < 20 && !_fbe; _bfb++ {
		_agc, _cfa := cmap.parseObject()
		if _cfa != nil {
			return _cfa
		}
		switch _gfe := _agc.(type) {
		case cmapOperand:
			switch _gfe.Operand {
			case "\u0064\u0065\u0066":
				_fbe = true
			default:
				_cf.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a\u0020\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076\u0020n\u0061\u006de\u003d\u0025\u0023\u0071", _agc, _aa)
				if _aa != "" {
					_aa = _b.Sprintf("\u0025\u0073\u0020%\u0073", _aa, _gfe.Operand)
				}
				_cf.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a \u0052\u0065\u0063\u006f\u0076\u0065\u0072e\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025\u0023\u0071", _aa)
			}
		case cmapName:
			_aa = _gfe.Name
		}
	}
	if !_fbe {
		_cf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0070\u0061\u0072\u0073\u0065N\u0061m\u0065:\u0020\u004e\u006f\u0020\u0064\u0065\u0066 ")
		return ErrBadCMap
	}
	cmap._fb = _aa
	return nil
}

type cmapName struct {
	Name string
}

func (_fdc *cMapParser) parseName() (cmapName, error) {
	_bgce := ""
	_bgeg := false
	for {
		_cacf, _adec := _fdc._deg.Peek(1)
		if _adec == _eg.EOF {
			break
		}
		if _adec != nil {
			return cmapName{_bgce}, _adec
		}
		if !_bgeg {
			if _cacf[0] == '/' {
				_bgeg = true
				_fdc._deg.ReadByte()
			} else {
				_cf.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u004e\u0061\u006d\u0065\u0020\u0073\u0074a\u0072t\u0069n\u0067 \u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0028\u0025\u0020\u0078\u0029", _cacf, _cacf)
				return cmapName{_bgce}, _b.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _cacf[0])
			}
		} else {
			if _bde.IsWhiteSpace(_cacf[0]) {
				break
			} else if (_cacf[0] == '/') || (_cacf[0] == '[') || (_cacf[0] == '(') || (_cacf[0] == ']') || (_cacf[0] == '<') || (_cacf[0] == '>') {
				break
			} else if _cacf[0] == '#' {
				_eag, _cfbf := _fdc._deg.Peek(3)
				if _cfbf != nil {
					return cmapName{_bgce}, _cfbf
				}
				_fdc._deg.Discard(3)
				_edc, _cfbf := _ba.DecodeString(string(_eag[1:3]))
				if _cfbf != nil {
					return cmapName{_bgce}, _cfbf
				}
				_bgce += string(_edc)
			} else {
				_dceg, _ := _fdc._deg.ReadByte()
				_bgce += string(_dceg)
			}
		}
	}
	return cmapName{_bgce}, nil
}
func (cmap *CMap) parseType() error {
	_egg := 0
	_fce := false
	for _aac := 0; _aac < 3 && !_fce; _aac++ {
		_geba, _fdf := cmap.parseObject()
		if _fdf != nil {
			return _fdf
		}
		switch _fcd := _geba.(type) {
		case cmapOperand:
			switch _fcd.Operand {
			case "\u0064\u0065\u0066":
				_fce = true
			default:
				_cf.Log.Error("\u0070\u0061r\u0073\u0065\u0054\u0079\u0070\u0065\u003a\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u006f=%\u0023\u0076", _geba)
				return ErrBadCMap
			}
		case cmapInt:
			_egg = int(_fcd._cfgd)
		}
	}
	cmap._gc = _egg
	return nil
}
func (cmap *CMap) parseCodespaceRange() error {
	for {
		_bgde, _ada := cmap.parseObject()
		if _ada != nil {
			if _ada == _eg.EOF {
				break
			}
			return _ada
		}
		_afga, _acbd := _bgde.(cmapHexString)
		if !_acbd {
			if _dcd, _badde := _bgde.(cmapOperand); _badde {
				if _dcd.Operand == _baa {
					return nil
				}
				return _c.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
			}
		}
		_bgde, _ada = cmap.parseObject()
		if _ada != nil {
			if _ada == _eg.EOF {
				break
			}
			return _ada
		}
		_egf, _acbd := _bgde.(cmapHexString)
		if !_acbd {
			return _c.New("\u006e\u006f\u006e-\u0068\u0065\u0078\u0020\u0068\u0069\u0067\u0068")
		}
		if len(_afga._baeg) != len(_egf._baeg) {
			return _c.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_bead := _cfcg(_afga)
		_abag := _cfcg(_egf)
		if _abag < _bead {
			_cf.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0061d\u0020\u0063\u006fd\u0065\u0073\u0070\u0061\u0063\u0065\u002e\u0020\u006cow\u003d\u0030\u0078%\u0030\u0032x\u0020\u0068\u0069\u0067\u0068\u003d0\u0078\u00250\u0032\u0078", _bead, _abag)
			return ErrBadCMap
		}
		_dff := _egf._cceg
		_fgf := Codespace{NumBytes: _dff, Low: _bead, High: _abag}
		cmap._gd = append(cmap._gd, _fgf)
		_cf.Log.Trace("\u0043\u006f\u0064e\u0073\u0070\u0061\u0063e\u0020\u006c\u006f\u0077\u003a\u0020\u0030x\u0025\u0058\u002c\u0020\u0068\u0069\u0067\u0068\u003a\u0020\u0030\u0078\u0025\u0058", _bead, _abag)
	}
	if len(cmap._gd) == 0 {
		_cf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u0020\u0069\u006e\u0020\u0063ma\u0070\u002e")
		return ErrBadCMap
	}
	return nil
}
func (_agd *cMapParser) parseArray() (cmapArray, error) {
	_eecg := cmapArray{}
	_eecg.Array = []cmapObject{}
	_agd._deg.ReadByte()
	for {
		_agd.skipSpaces()
		_gdd, _ebe := _agd._deg.Peek(1)
		if _ebe != nil {
			return _eecg, _ebe
		}
		if _gdd[0] == ']' {
			_agd._deg.ReadByte()
			break
		}
		_cgd, _ebe := _agd.parseObject()
		if _ebe != nil {
			return _eecg, _ebe
		}
		_eecg.Array = append(_eecg.Array, _cgd)
	}
	return _eecg, nil
}
func _abga(_bed cmapHexString) []rune {
	if len(_bed._baeg) == 1 {
		return []rune{rune(_bed._baeg[0])}
	}
	_fdg := _bed._baeg
	if len(_fdg)%2 != 0 {
		_fdg = append(_fdg, 0)
		_cf.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u0073\u002e\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0073\u0068\u0065\u0078\u003d\u0025#\u0076\u0020\u0074\u006f\u0020\u0025\u002b\u0076", _bed, _fdg)
	}
	_cgbc := len(_fdg) >> 1
	_dda := make([]uint16, _cgbc)
	for _acda := 0; _acda < _cgbc; _acda++ {
		_dda[_acda] = uint16(_fdg[_acda<<1])<<8 + uint16(_fdg[_acda<<1+1])
	}
	_fgb := _f.Decode(_dda)
	return _fgb
}
func (cmap *CMap) BytesToCharcodes(data []byte) ([]CharCode, bool) {
	var _cc []CharCode
	if cmap._db == 8 {
		for _, _gdc := range data {
			_cc = append(_cc, CharCode(_gdc))
		}
		return _cc, true
	}
	for _bgf := 0; _bgf < len(data); {
		_ffac, _dfb, _dbbe := cmap.matchCode(data[_bgf:])
		if !_dbbe {
			_cf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0074\u0020\u0069\u003d\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", _bgf, data, string(data))
			return _cc, false
		}
		_cc = append(_cc, _ffac)
		_bgf += _dfb
	}
	return _cc, true
}
func (cmap *CMap) WMode() (int, bool) { return cmap._gf._gcg, cmap._gf._fde }

type CIDSystemInfo struct {
	Registry   string
	Ordering   string
	Supplement int
}
type cmapString struct {
	String string
}

func (_fgde *cMapParser) parseDict() (cmapDict, error) {
	_cf.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_add := _fcg()
	_gdcf, _ := _fgde._deg.ReadByte()
	if _gdcf != '<' {
		return _add, ErrBadCMapDict
	}
	_gdcf, _ = _fgde._deg.ReadByte()
	if _gdcf != '<' {
		return _add, ErrBadCMapDict
	}
	for {
		_fgde.skipSpaces()
		_fee, _gfa := _fgde._deg.Peek(2)
		if _gfa != nil {
			return _add, _gfa
		}
		if (_fee[0] == '>') && (_fee[1] == '>') {
			_fgde._deg.ReadByte()
			_fgde._deg.ReadByte()
			break
		}
		_ded, _gfa := _fgde.parseName()
		_cf.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _ded.Name)
		if _gfa != nil {
			_cf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u002e\u0020\u0065\u0072r=\u0025\u0076", _gfa)
			return _add, _gfa
		}
		_fgde.skipSpaces()
		_gebe, _gfa := _fgde.parseObject()
		if _gfa != nil {
			return _add, _gfa
		}
		_add.Dict[_ded.Name] = _gebe
		_fgde.skipSpaces()
		_fee, _gfa = _fgde._deg.Peek(3)
		if _gfa != nil {
			return _add, _gfa
		}
		if string(_fee) == "\u0064\u0065\u0066" {
			_fgde._deg.Discard(3)
		}
	}
	return _add, nil
}
func _cfcg(_efed cmapHexString) CharCode {
	_bgcd := CharCode(0)
	for _, _dea := range _efed._baeg {
		_bgcd <<= 8
		_bgcd |= CharCode(_dea)
	}
	return _bgcd
}
func _afda(_dfaa string) rune { _fca := []rune(_dfaa); return _fca[len(_fca)-1] }
func (cmap *CMap) CharcodeToCID(code CharCode) (CharCode, bool) {
	_bb, _caf := cmap._bad[code]
	return _bb, _caf
}
func (_dbd *cMapParser) skipSpaces() (int, error) {
	_cee := 0
	for {
		_ceac, _afaa := _dbd._deg.Peek(1)
		if _afaa != nil {
			return 0, _afaa
		}
		if _bde.IsWhiteSpace(_ceac[0]) {
			_dbd._deg.ReadByte()
			_cee++
		} else {
			break
		}
	}
	return _cee, nil
}

const (
	_dc               = 4
	MissingCodeRune   = '\ufffd'
	MissingCodeString = string(MissingCodeRune)
)

func (cmap *CMap) CIDSystemInfo() CIDSystemInfo { return cmap._cfb }

type cmapHexString struct {
	_cceg int
	_baeg []byte
}
type CharCode uint32

func (cmap *CMap) NBits() int { return cmap._db }
func (_daf *cMapParser) parseComment() (string, error) {
	var _aag _ed.Buffer
	_, _dffc := _daf.skipSpaces()
	if _dffc != nil {
		return _aag.String(), _dffc
	}
	_gdef := true
	for {
		_gbef, _cdgd := _daf._deg.Peek(1)
		if _cdgd != nil {
			_cf.Log.Debug("p\u0061r\u0073\u0065\u0043\u006f\u006d\u006d\u0065\u006et\u003a\u0020\u0065\u0072r=\u0025\u0076", _cdgd)
			return _aag.String(), _cdgd
		}
		if _gdef && _gbef[0] != '%' {
			return _aag.String(), ErrBadCMapComment
		}
		_gdef = false
		if (_gbef[0] != '\r') && (_gbef[0] != '\n') {
			_ggd, _ := _daf._deg.ReadByte()
			_aag.WriteByte(_ggd)
		} else {
			break
		}
	}
	return _aag.String(), nil
}
func (cmap *CMap) Bytes() []byte {
	_cf.Log.Trace("\u0063\u006d\u0061\u0070.B\u0079\u0074\u0065\u0073\u003a\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", cmap.String())
	if len(cmap._dbb) > 0 {
		return cmap._dbb
	}
	cmap._dbb = []byte(_bf.Join([]string{_abg, cmap.toBfData(), _gabf}, "\u000a"))
	return cmap._dbb
}
func _caa(_bbb, _cba int) int {
	if _bbb < _cba {
		return _bbb
	}
	return _cba
}
func _da(_fd bool) *CMap {
	_dfa := 16
	if _fd {
		_dfa = 8
	}
	return &CMap{_db: _dfa, _bad: make(map[CharCode]CharCode), _afa: make(map[CharCode]CharCode), _bea: make(map[CharCode]string), _ad: make(map[string]CharCode)}
}
func (cmap *CMap) parseVersion() error {
	_ffb := ""
	_bcf := false
	for _dab := 0; _dab < 3 && !_bcf; _dab++ {
		_aaa, _fbc := cmap.parseObject()
		if _fbc != nil {
			return _fbc
		}
		switch _edd := _aaa.(type) {
		case cmapOperand:
			switch _edd.Operand {
			case "\u0064\u0065\u0066":
				_bcf = true
			default:
				_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0070\u0061\u0072\u0073\u0065\u0056e\u0072\u0073\u0069\u006f\u006e\u003a \u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076", _aaa)
				return ErrBadCMap
			}
		case cmapInt:
			_ffb = _b.Sprintf("\u0025\u0064", _edd._cfgd)
		case cmapFloat:
			_ffb = _b.Sprintf("\u0025\u0066", _edd._cdfd)
		case cmapString:
			_ffb = _edd.String
		default:
			_cf.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020p\u0061\u0072\u0073\u0065Ver\u0073io\u006e\u003a\u0020\u0042\u0061\u0064\u0020ty\u0070\u0065\u002e\u0020\u006f\u003d\u0025#\u0076", _aaa)
		}
	}
	cmap._af = _ffb
	return nil
}
func (cmap *CMap) parseBfchar() error {
	for {
		_dfc, _ffad := cmap.parseObject()
		if _ffad != nil {
			if _ffad == _eg.EOF {
				break
			}
			return _ffad
		}
		var _bcef CharCode
		switch _abd := _dfc.(type) {
		case cmapOperand:
			if _abd.Operand == _cec {
				return nil
			}
			return _c.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_bcef = _cfcg(_abd)
		default:
			return _c.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_dfc, _ffad = cmap.parseObject()
		if _ffad != nil {
			if _ffad == _eg.EOF {
				break
			}
			return _ffad
		}
		var _cdg []rune
		switch _fgd := _dfc.(type) {
		case cmapOperand:
			if _fgd.Operand == _cec {
				return nil
			}
			_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u0070\u0065\u0072\u0061\u006e\u0064\u002e\u0020\u0025\u0023\u0076", _fgd)
			return ErrBadCMap
		case cmapHexString:
			_cdg = _abga(_fgd)
		case cmapName:
			_cf.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0061\u006de\u002e \u0025\u0023\u0076", _fgd)
			_cdg = []rune{MissingCodeRune}
		default:
			_cf.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0074\u0079\u0070e\u002e \u0025\u0023\u0076", _dfc)
			return ErrBadCMap
		}
		cmap._bea[_bcef] = string(_cdg)
	}
	return nil
}

type cMapParser struct{ _deg *_ef.Reader }

func (cmap *CMap) toBfData() string {
	if len(cmap._bea) == 0 {
		return ""
	}
	_afgc := make([]CharCode, 0, len(cmap._bea))
	for _geg := range cmap._bea {
		_afgc = append(_afgc, _geg)
	}
	_bd.Slice(_afgc, func(_geb, _deb int) bool { return _afgc[_geb] < _afgc[_deb] })
	var _cgbg []charRange
	_eb := charRange{_afgc[0], _afgc[0]}
	_gg := cmap._bea[_afgc[0]]
	for _, _ebf := range _afgc[1:] {
		_dba := cmap._bea[_ebf]
		if _ebf == _eb._bfe+1 && _afda(_dba) == _afda(_gg)+1 {
			_eb._bfe = _ebf
		} else {
			_cgbg = append(_cgbg, _eb)
			_eb._ca, _eb._bfe = _ebf, _ebf
		}
		_gg = _dba
	}
	_cgbg = append(_cgbg, _eb)
	var _fab []CharCode
	var _dcef []fbRange
	for _, _adc := range _cgbg {
		if _adc._ca == _adc._bfe {
			_fab = append(_fab, _adc._ca)
		} else {
			_dcef = append(_dcef, fbRange{_g: _adc._ca, _cg: _adc._bfe, _cgb: cmap._bea[_adc._ca]})
		}
	}
	_cf.Log.Trace("\u0063\u0068ar\u0052\u0061\u006eg\u0065\u0073\u003d\u0025d f\u0062Ch\u0061\u0072\u0073\u003d\u0025\u0064\u0020fb\u0052\u0061\u006e\u0067\u0065\u0073\u003d%\u0064", len(_cgbg), len(_fab), len(_dcef))
	var _dbf []string
	if len(_fab) > 0 {
		_eebg := (len(_fab) + _fe - 1) / _fe
		for _gab := 0; _gab < _eebg; _gab++ {
			_eedf := _caa(len(_fab)-_gab*_fe, _fe)
			_dbf = append(_dbf, _b.Sprintf("\u0025\u0064\u0020\u0062\u0065\u0067\u0069\u006e\u0062f\u0063\u0068\u0061\u0072", _eedf))
			for _bgc := 0; _bgc < _eedf; _bgc++ {
				_dcc := _fab[_gab*_fe+_bgc]
				_dae := cmap._bea[_dcc]
				_dbf = append(_dbf, _b.Sprintf("\u003c%\u0030\u0034\u0078\u003e\u0020\u0025s", _dcc, _ceb(_dae)))
			}
			_dbf = append(_dbf, "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r")
		}
	}
	if len(_dcef) > 0 {
		_cb := (len(_dcef) + _fe - 1) / _fe
		for _def := 0; _def < _cb; _def++ {
			_cdda := _caa(len(_dcef)-_def*_fe, _fe)
			_dbf = append(_dbf, _b.Sprintf("\u0025d\u0020b\u0065\u0067\u0069\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065", _cdda))
			for _ggg := 0; _ggg < _cdda; _ggg++ {
				_cea := _dcef[_def*_fe+_ggg]
				_dbf = append(_dbf, _b.Sprintf("\u003c%\u00304\u0078\u003e\u003c\u0025\u0030\u0034\u0078\u003e\u0020\u0025\u0073", _cea._g, _cea._cg, _ceb(_cea._cgb)))
			}
			_dbf = append(_dbf, "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065")
		}
	}
	return _bf.Join(_dbf, "\u000a")
}
func (cmap *CMap) parseSystemInfo() error {
	_aeg := false
	_ddg := false
	_dgb := ""
	_edb := false
	_gag := CIDSystemInfo{}
	for _ebd := 0; _ebd < 50 && !_edb; _ebd++ {
		_ccda, _gacb := cmap.parseObject()
		if _gacb != nil {
			return _gacb
		}
		switch _ffbc := _ccda.(type) {
		case cmapDict:
			_bcb := _ffbc.Dict
			_gaee, _acb := _bcb["\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"]
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_gebd, _acb := _gaee.(cmapString)
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_gag.Registry = _gebd.String
			_gaee, _acb = _bcb["\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"]
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_gebd, _acb = _gaee.(cmapString)
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_gag.Ordering = _gebd.String
			_fbb, _acb := _bcb["\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"]
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cfgb, _acb := _fbb.(cmapInt)
			if !_acb {
				_cf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_gag.Supplement = int(_cfgb._cfgd)
			_edb = true
		case cmapOperand:
			switch _ffbc.Operand {
			case "\u0062\u0065\u0067i\u006e":
				_aeg = true
			case "\u0065\u006e\u0064":
				_edb = true
			case "\u0064\u0065\u0066":
				_ddg = false
			}
		case cmapName:
			if _aeg {
				_dgb = _ffbc.Name
				_ddg = true
			}
		case cmapString:
			if _ddg {
				switch _dgb {
				case "\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079":
					_gag.Registry = _ffbc.String
				case "\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067":
					_gag.Ordering = _ffbc.String
				}
			}
		case cmapInt:
			if _ddg {
				switch _dgb {
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_gag.Supplement = int(_ffbc._cfgd)
				}
			}
		}
	}
	if !_edb {
		_cf.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u006c\u0079")
		return ErrBadCMap
	}
	cmap._cfb = _gag
	return nil
}
func (_gdgf *cMapParser) parseNumber() (cmapObject, error) {
	_ebdf, _fdd := _bde.ParseNumber(_gdgf._deg)
	if _fdd != nil {
		return nil, _fdd
	}
	switch _bda := _ebdf.(type) {
	case *_bde.PdfObjectFloat:
		return cmapFloat{float64(*_bda)}, nil
	case *_bde.PdfObjectInteger:
		return cmapInt{int64(*_bda)}, nil
	}
	return nil, _b.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _ebdf)
}
func _faeg(_cedb cmapHexString) rune {
	_eca := _abga(_cedb)
	if _abab := len(_eca); _abab == 0 {
		_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054o\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065c\u0074\u0065\u0064\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0072u\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023\u0076", _cedb)
		return MissingCodeRune
	}
	if len(_eca) > 1 {
		_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0078\u0061\u0063\u0074\u006c\u0079\u0020\u006f\u006e\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023v\u0020\u002d\u003e\u0020\u0025#\u0076", _cedb, _eca)
	}
	return _eca[0]
}
func (cmap *CMap) CharcodeToUnicode(code CharCode) (string, bool) {
	if _cfg, _bge := cmap._bea[code]; _bge {
		return _cfg, true
	}
	return MissingCodeString, false
}

type CMap struct {
	*cMapParser
	_fb  string
	_db  int
	_gc  int
	_af  string
	_cab string
	_cfb CIDSystemInfo
	_gd  []Codespace
	_bad map[CharCode]CharCode
	_afa map[CharCode]CharCode
	_bea map[CharCode]string
	_ad  map[string]CharCode
	_dbb []byte
	_cge *_bde.PdfObjectStream
	_gf  integer
}

func _ffda(_eeg []byte) *cMapParser {
	_ecf := cMapParser{}
	_dcbc := _ed.NewBuffer(_eeg)
	_ecf._deg = _ef.NewReader(_dcbc)
	return &_ecf
}
func LoadCmapFromDataCID(data []byte) (*CMap, error) { return LoadCmapFromData(data, false) }

type cmapFloat struct{ _cdfd float64 }

func (cmap *CMap) parseCIDRange() error {
	for {
		_bfbg, _cca := cmap.parseObject()
		if _cca != nil {
			if _cca == _eg.EOF {
				break
			}
			return _cca
		}
		_ddc, _ggc := _bfbg.(cmapHexString)
		if !_ggc {
			if _bab, _eecb := _bfbg.(cmapOperand); _eecb {
				if _bab.Operand == _fea {
					return nil
				}
				return _c.New("\u0063\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0072\u0076\u0061\u006c\u0020s\u0074\u0061\u0072\u0074\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
			}
		}
		_cabb := _cfcg(_ddc)
		_bfbg, _cca = cmap.parseObject()
		if _cca != nil {
			if _cca == _eg.EOF {
				break
			}
			return _cca
		}
		_fgg, _ggc := _bfbg.(cmapHexString)
		if !_ggc {
			return _c.New("\u0063\u0069d\u0020\u0069\u006e\u0074e\u0072\u0076a\u006c\u0020\u0065\u006e\u0064\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_ddc._baeg) != len(_fgg._baeg) {
			return _c.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_gaf := _cfcg(_fgg)
		if _cabb > _gaf {
			_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0049\u0044\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073t\u0061\u0072\u0074\u003d\u0030\u0078\u0025\u0030\u0032\u0078\u0020\u0065\u006e\u0064=\u0030x\u0025\u0030\u0032\u0078", _cabb, _gaf)
			return ErrBadCMap
		}
		_bfbg, _cca = cmap.parseObject()
		if _cca != nil {
			if _cca == _eg.EOF {
				break
			}
			return _cca
		}
		_ccab, _ggc := _bfbg.(cmapInt)
		if !_ggc {
			return _c.New("\u0063\u0069\u0064\u0020\u0073t\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		if _ccab._cfgd < 0 {
			return _c.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0063\u0069\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065")
		}
		_eef := _ccab._cfgd
		for _fge := _cabb; _fge <= _gaf; _fge++ {
			cmap._bad[_fge] = CharCode(_eef)
			_eef++
		}
		_cf.Log.Trace("C\u0049\u0044\u0020\u0072\u0061\u006eg\u0065\u003a\u0020\u003c\u0030\u0078\u0025\u0058\u003e \u003c\u0030\u0078%\u0058>\u0020\u0025\u0064", _cabb, _gaf, _ccab._cfgd)
	}
	return nil
}

type cmapInt struct{ _cfgd int64 }
type charRange struct {
	_ca  CharCode
	_bfe CharCode
}

func (_egcb *cMapParser) parseOperand() (cmapOperand, error) {
	_eae := cmapOperand{}
	_bdb := _ed.Buffer{}
	for {
		_dgf, _defb := _egcb._deg.Peek(1)
		if _defb != nil {
			if _defb == _eg.EOF {
				break
			}
			return _eae, _defb
		}
		if _bde.IsDelimiter(_dgf[0]) {
			break
		}
		if _bde.IsWhiteSpace(_dgf[0]) {
			break
		}
		_cbe, _ := _egcb._deg.ReadByte()
		_bdb.WriteByte(_cbe)
	}
	if _bdb.Len() == 0 {
		return _eae, _b.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	_eae.Operand = _bdb.String()
	return _eae, nil
}
func LoadPredefinedCMap(name string) (*CMap, error) {
	cmap, _bca := _ab(name)
	if _bca != nil {
		return nil, _bca
	}
	if cmap._cab == "" {
		cmap.computeInverseMappings()
		return cmap, nil
	}
	_ede, _bca := _ab(cmap._cab)
	if _bca != nil {
		return nil, _bca
	}
	for _ag, _cdc := range _ede._bad {
		if _, _afc := cmap._bad[_ag]; !_afc {
			cmap._bad[_ag] = _cdc
		}
	}
	cmap._gd = append(cmap._gd, _ede._gd...)
	cmap.computeInverseMappings()
	return cmap, nil
}

var (
	ErrBadCMap        = _c.New("\u0062\u0061\u0064\u0020\u0063\u006d\u0061\u0070")
	ErrBadCMapComment = _c.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
	ErrBadCMapDict    = _c.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
)

func NewToUnicodeCMap(codeToRune map[CharCode]rune) *CMap {
	_egd := make(map[CharCode]string, len(codeToRune))
	for _ae, _de := range codeToRune {
		_egd[_ae] = string(_de)
	}
	cmap := &CMap{_fb: "\u0041d\u006fb\u0065\u002d\u0049\u0064\u0065n\u0074\u0069t\u0079\u002d\u0055\u0043\u0053", _gc: 2, _db: 16, _cfb: CIDSystemInfo{Registry: "\u0041\u0064\u006fb\u0065", Ordering: "\u0055\u0043\u0053", Supplement: 0}, _gd: []Codespace{{Low: 0, High: 0xffff}}, _bea: _egd, _ad: make(map[string]CharCode, len(codeToRune)), _bad: make(map[CharCode]CharCode, len(codeToRune)), _afa: make(map[CharCode]CharCode, len(codeToRune))}
	cmap.computeInverseMappings()
	return cmap
}
func NewCIDSystemInfo(obj _bde.PdfObject) (_a CIDSystemInfo, _be error) {
	_ga, _fc := _bde.GetDict(obj)
	if !_fc {
		return CIDSystemInfo{}, _bde.ErrTypeError
	}
	_df, _fc := _bde.GetStringVal(_ga.Get("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"))
	if !_fc {
		return CIDSystemInfo{}, _bde.ErrTypeError
	}
	_eda, _fc := _bde.GetStringVal(_ga.Get("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"))
	if !_fc {
		return CIDSystemInfo{}, _bde.ErrTypeError
	}
	_bc, _fc := _bde.GetIntVal(_ga.Get("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"))
	if !_fc {
		return CIDSystemInfo{}, _bde.ErrTypeError
	}
	return CIDSystemInfo{Registry: _df, Ordering: _eda, Supplement: _bc}, nil
}

const (
	_cfe  = "\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"
	_ggce = "\u0062e\u0067\u0069\u006e\u0063\u006d\u0061p"
	_agf  = "\u0065n\u0064\u0063\u006d\u0061\u0070"
	_dgae = "\u0062\u0065\u0067\u0069nc\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0072\u0061\u006e\u0067\u0065"
	_baa  = "\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065"
	_aeb  = "b\u0065\u0067\u0069\u006e\u0062\u0066\u0063\u0068\u0061\u0072"
	_cec  = "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r"
	_dfbb = "\u0062\u0065\u0067i\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_bddc = "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_cece = "\u0062\u0065\u0067\u0069\u006e\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_fea  = "e\u006e\u0064\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_dbcb = "\u0075s\u0065\u0063\u006d\u0061\u0070"
	_geca = "\u0057\u004d\u006fd\u0065"
	_bcff = "\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065"
	_bcfb = "\u0043\u004d\u0061\u0070\u0054\u0079\u0070\u0065"
	_ecc  = "C\u004d\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e"
)

type cmapDict struct{ Dict map[string]cmapObject }

func (cmap *CMap) parse() error {
	var _dbc cmapObject
	for {
		_cda, _badd := cmap.parseObject()
		if _badd != nil {
			if _badd == _eg.EOF {
				break
			}
			_cf.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u004d\u0061\u0070\u003a\u0020\u0025\u0076", _badd)
			return _badd
		}
		switch _ged := _cda.(type) {
		case cmapOperand:
			_acd := _ged
			switch _acd.Operand {
			case _dgae:
				_cebg := cmap.parseCodespaceRange()
				if _cebg != nil {
					return _cebg
				}
			case _cece:
				_ggf := cmap.parseCIDRange()
				if _ggf != nil {
					return _ggf
				}
			case _aeb:
				_abf := cmap.parseBfchar()
				if _abf != nil {
					return _abf
				}
			case _dfbb:
				_bdc := cmap.parseBfrange()
				if _bdc != nil {
					return _bdc
				}
			case _dbcb:
				if _dbc == nil {
					_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u0073\u0065\u0063m\u0061\u0070\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f \u0061\u0072\u0067")
					return ErrBadCMap
				}
				_gee, _edeb := _dbc.(cmapName)
				if !_edeb {
					_cf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0075\u0073\u0065\u0063\u006d\u0061\u0070\u0020\u0061\u0072\u0067\u0020\u006eo\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u0020\u0025\u0023\u0076", _dbc)
					return ErrBadCMap
				}
				cmap._cab = _gee.Name
			case _cfe:
				_ggfb := cmap.parseSystemInfo()
				if _ggfb != nil {
					return _ggfb
				}
			}
		case cmapName:
			_gec := _ged
			switch _gec.Name {
			case _cfe:
				_gegc := cmap.parseSystemInfo()
				if _gegc != nil {
					return _gegc
				}
			case _bcff:
				_dee := cmap.parseName()
				if _dee != nil {
					return _dee
				}
			case _bcfb:
				_ffd := cmap.parseType()
				if _ffd != nil {
					return _ffd
				}
			case _ecc:
				_bcc := cmap.parseVersion()
				if _bcc != nil {
					return _bcc
				}
			case _geca:
				if _badd = cmap.parseWMode(); _badd != nil {
					return _badd
				}
			}
		}
		_dbc = _cda
	}
	return nil
}
func _ab(_dg string) (*CMap, error) {
	_aef, _dbg := _cd.Asset(_dg)
	if _dbg != nil {
		return nil, _dbg
	}
	return LoadCmapFromDataCID(_aef)
}

const (
	_fe   = 100
	_abg  = "\u000a\u002f\u0043\u0049\u0044\u0049\u006e\u0069\u0074\u0020\u002f\u0050\u0072\u006fc\u0053\u0065\u0074\u0020\u0066\u0069\u006e\u0064\u0072es\u006fu\u0072c\u0065 \u0062\u0065\u0067\u0069\u006e\u000a\u0031\u0032\u0020\u0064\u0069\u0063\u0074\u0020\u0062\u0065\u0067\u0069n\u000a\u0062\u0065\u0067\u0069\u006e\u0063\u006d\u0061\u0070\n\u002f\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u003c\u003c\u0020\u002f\u0052\u0065\u0067\u0069\u0073t\u0072\u0079\u0020\u0028\u0041\u0064\u006f\u0062\u0065\u0029\u0020\u002f\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0028\u0055\u0043\u0053)\u0020\u002f\u0053\u0075\u0070p\u006c\u0065\u006d\u0065\u006et\u0020\u0030\u0020\u003e\u003e\u0020\u0064\u0065\u0066\u000a\u002f\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u002f\u0041\u0064\u006f\u0062\u0065-\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0055\u0043\u0053\u0020\u0064\u0065\u0066\u000a\u002fC\u004d\u0061\u0070\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0064\u0065\u0066\u000a\u0031\u0020\u0062\u0065\u0067\u0069\u006e\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063e\u0072\u0061n\u0067\u0065\n\u003c\u0030\u0030\u0030\u0030\u003e\u0020<\u0046\u0046\u0046\u0046\u003e\u000a\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065\u000a"
	_gabf = "\u0065\u006e\u0064\u0063\u006d\u0061\u0070\u000a\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u0063ur\u0072e\u006e\u0074\u0064\u0069\u0063\u0074\u0020\u002f\u0043\u004d\u0061\u0070 \u0064\u0065\u0066\u0069\u006e\u0065\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0070\u006fp\u000a\u0065\u006e\u0064\u000a\u0065\u006e\u0064\u000a"
)

func (cmap *CMap) Type() int { return cmap._gc }
func (_efg *cMapParser) parseString() (cmapString, error) {
	_efg._deg.ReadByte()
	_dbec := _ed.Buffer{}
	_bccb := 1
	for {
		_bbd, _cdbd := _efg._deg.Peek(1)
		if _cdbd != nil {
			return cmapString{_dbec.String()}, _cdbd
		}
		if _bbd[0] == '\\' {
			_efg._deg.ReadByte()
			_fae, _badc := _efg._deg.ReadByte()
			if _badc != nil {
				return cmapString{_dbec.String()}, _badc
			}
			if _bde.IsOctalDigit(_fae) {
				_eecf, _bced := _efg._deg.Peek(2)
				if _bced != nil {
					return cmapString{_dbec.String()}, _bced
				}
				var _fbbd []byte
				_fbbd = append(_fbbd, _fae)
				for _, _ebg := range _eecf {
					if _bde.IsOctalDigit(_ebg) {
						_fbbd = append(_fbbd, _ebg)
					} else {
						break
					}
				}
				_efg._deg.Discard(len(_fbbd) - 1)
				_cf.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _fbbd)
				_bfda, _bced := _e.ParseUint(string(_fbbd), 8, 32)
				if _bced != nil {
					return cmapString{_dbec.String()}, _bced
				}
				_dbec.WriteByte(byte(_bfda))
				continue
			}
			switch _fae {
			case 'n':
				_dbec.WriteByte('\n')
			case 'r':
				_dbec.WriteByte('\r')
			case 't':
				_dbec.WriteByte('\t')
			case 'b':
				_dbec.WriteByte('\b')
			case 'f':
				_dbec.WriteByte('\f')
			case '(':
				_dbec.WriteByte('(')
			case ')':
				_dbec.WriteByte(')')
			case '\\':
				_dbec.WriteByte('\\')
			}
			continue
		} else if _bbd[0] == '(' {
			_bccb++
		} else if _bbd[0] == ')' {
			_bccb--
			if _bccb == 0 {
				_efg._deg.ReadByte()
				break
			}
		}
		_cecb, _ := _efg._deg.ReadByte()
		_dbec.WriteByte(_cecb)
	}
	return cmapString{_dbec.String()}, nil
}
func (_ddcb *cMapParser) parseObject() (cmapObject, error) {
	_ddcb.skipSpaces()
	for {
		_bbf, _fbg := _ddcb._deg.Peek(2)
		if _fbg != nil {
			return nil, _fbg
		}
		if _bbf[0] == '%' {
			_ddcb.parseComment()
			_ddcb.skipSpaces()
			continue
		} else if _bbf[0] == '/' {
			_ebfd, _aed := _ddcb.parseName()
			return _ebfd, _aed
		} else if _bbf[0] == '(' {
			_ade, _daaa := _ddcb.parseString()
			return _ade, _daaa
		} else if _bbf[0] == '[' {
			_beed, _abc := _ddcb.parseArray()
			return _beed, _abc
		} else if (_bbf[0] == '<') && (_bbf[1] == '<') {
			_dac, _ddd := _ddcb.parseDict()
			return _dac, _ddd
		} else if _bbf[0] == '<' {
			_gfd, _cbc := _ddcb.parseHexString()
			return _gfd, _cbc
		} else if _bde.IsDecimalDigit(_bbf[0]) || (_bbf[0] == '-' && _bde.IsDecimalDigit(_bbf[1])) {
			_ffaf, _agce := _ddcb.parseNumber()
			if _agce != nil {
				return nil, _agce
			}
			return _ffaf, nil
		} else {
			_cfc, _dfg := _ddcb.parseOperand()
			if _dfg != nil {
				return nil, _dfg
			}
			return _cfc, nil
		}
	}
}
