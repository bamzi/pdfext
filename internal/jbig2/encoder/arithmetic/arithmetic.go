package arithmetic

import (
	_a "bytes"
	_ba "io"

	_e "github.com/bamzi/pdfext/common"
	_d "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_db "github.com/bamzi/pdfext/internal/jbig2/errors"
)

type codingContext struct {
	_fe []byte
	_ce []byte
}

func (_ca *Encoder) renormalize() {
	for {
		_ca._adc <<= 1
		_ca._g <<= 1
		_ca._cg--
		if _ca._cg == 0 {
			_ca.byteOut()
		}
		if (_ca._adc & 0x8000) != 0 {
			break
		}
	}
}

type Encoder struct {
	_g       uint32
	_adc     uint16
	_cg, _fc uint8
	_gb      int
	_gg      int
	_cfd     [][]byte
	_ade     []byte
	_cb      int
	_egb     *codingContext
	_bg      [13]*codingContext
	_bga     *codingContext
}

func (_bgg *Encoder) Init() {
	_bgg._egb = _feb(_bbf)
	_bgg._adc = 0x8000
	_bgg._g = 0
	_bgg._cg = 12
	_bgg._gb = -1
	_bgg._fc = 0
	_bgg._cb = 0
	_bgg._ade = make([]byte, _dacd)
	for _bad := 0; _bad < len(_bgg._bg); _bad++ {
		_bgg._bg[_bad] = _feb(512)
	}
	_bgg._bga = nil
}

const (
	_bbf  = 65536
	_dacd = 20 * 1024
)

func New() *Encoder { _ef := &Encoder{}; _ef.Init(); return _ef }
func (_gaaa *Encoder) encodeInteger(_eb Class, _eec int) error {
	const _cbg = "E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072"
	if _eec > 2000000000 || _eec < -2000000000 {
		return _db.Errorf(_cbg, "\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", _eec)
	}
	_bbb := _gaaa._bg[_eb]
	_bbd := uint32(1)
	var _afde int
	for ; ; _afde++ {
		if _cf[_afde]._ac <= _eec && _cf[_afde]._ad >= _eec {
			break
		}
	}
	if _eec < 0 {
		_eec = -_eec
	}
	_eec -= int(_cf[_afde]._bd)
	_fad := _cf[_afde]._f
	for _aaf := uint8(0); _aaf < _cf[_afde]._ea; _aaf++ {
		_bag := _fad & 1
		if _feg := _gaaa.encodeBit(_bbb, _bbd, _bag); _feg != nil {
			return _db.Wrap(_feg, _cbg, "")
		}
		_fad >>= 1
		if _bbd&0x100 > 0 {
			_bbd = (((_bbd << 1) | uint32(_bag)) & 0x1ff) | 0x100
		} else {
			_bbd = (_bbd << 1) | uint32(_bag)
		}
	}
	_eec <<= 32 - _cf[_afde]._eg
	for _fba := uint8(0); _fba < _cf[_afde]._eg; _fba++ {
		_be := uint8((uint32(_eec) & 0x80000000) >> 31)
		if _aea := _gaaa.encodeBit(_bbb, _bbd, _be); _aea != nil {
			return _db.Wrap(_aea, _cbg, "\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064")
		}
		_eec <<= 1
		if _bbd&0x100 != 0 {
			_bbd = (((_bbd << 1) | uint32(_be)) & 0x1ff) | 0x100
		} else {
			_bbd = (_bbd << 1) | uint32(_be)
		}
	}
	return nil
}
func (_gcd *Encoder) byteOut() {
	if _gcd._fc == 0xff {
		_gcd.rBlock()
		return
	}
	if _gcd._g < 0x8000000 {
		_gcd.lBlock()
		return
	}
	_gcd._fc++
	if _gcd._fc != 0xff {
		_gcd.lBlock()
		return
	}
	_gcd._g &= 0x7ffffff
	_gcd.rBlock()
}

const _efd = 0x9b25

func (_gfgc *Encoder) encodeIAID(_dbe, _gbfb int) error {
	if _gfgc._bga == nil {
		_gfgc._bga = _feb(1 << uint(_dbe))
	}
	_gaf := uint32(1<<uint32(_dbe+1)) - 1
	_gbfb <<= uint(32 - _dbe)
	_ffbd := uint32(1)
	for _efa := 0; _efa < _dbe; _efa++ {
		_gfgg := _ffbd & _gaf
		_gag := uint8((uint32(_gbfb) & 0x80000000) >> 31)
		if _df := _gfgc.encodeBit(_gfgc._bga, _gfgg, _gag); _df != nil {
			return _df
		}
		_ffbd = (_ffbd << 1) | uint32(_gag)
		_gbfb <<= 1
	}
	return nil
}
func (_ead *Encoder) Reset() {
	_ead._adc = 0x8000
	_ead._g = 0
	_ead._cg = 12
	_ead._gb = -1
	_ead._fc = 0
	_ead._bga = nil
	_ead._egb = _feb(_bbf)
}

var _cf = []intEncRangeS{{0, 3, 0, 2, 0, 2}, {-1, -1, 9, 4, 0, 0}, {-3, -2, 5, 3, 2, 1}, {4, 19, 2, 3, 4, 4}, {-19, -4, 3, 3, 4, 4}, {20, 83, 6, 4, 20, 6}, {-83, -20, 7, 4, 20, 6}, {84, 339, 14, 5, 84, 8}, {-339, -84, 15, 5, 84, 8}, {340, 4435, 30, 6, 340, 12}, {-4435, -340, 31, 6, 340, 12}, {4436, 2000000000, 62, 6, 4436, 32}, {-2000000000, -4436, 63, 6, 4436, 32}}

func (_edg *Encoder) codeLPS(_dba *codingContext, _bffc uint32, _bgd uint16, _efdb byte) {
	_edg._adc -= _bgd
	if _edg._adc < _bgd {
		_edg._g += uint32(_bgd)
	} else {
		_edg._adc = _bgd
	}
	if _ceg[_efdb]._aba == 1 {
		_dba.flipMps(_bffc)
	}
	_dba._fe[_bffc] = _ceg[_efdb]._efda
	_edg.renormalize()
}
func (_cab *Encoder) setBits() {
	_cfb := _cab._g + uint32(_cab._adc)
	_cab._g |= 0xffff
	if _cab._g >= _cfb {
		_cab._g -= 0x8000
	}
}

const (
	IAAI Class = iota
	IADH
	IADS
	IADT
	IADW
	IAEX
	IAFS
	IAIT
	IARDH
	IARDW
	IARDX
	IARDY
	IARI
)

func (_ed *codingContext) flipMps(_bb uint32) { _ed._ce[_bb] = 1 - _ed._ce[_bb] }
func (_daa *Encoder) WriteTo(w _ba.Writer) (int64, error) {
	const _bfa = "\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f"
	var _aed int64
	for _edd, _bbe := range _daa._cfd {
		_aca, _eff := w.Write(_bbe)
		if _eff != nil {
			return 0, _db.Wrapf(_eff, _bfa, "\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b", _edd)
		}
		_aed += int64(_aca)
	}
	_daa._ade = _daa._ade[:_daa._cb]
	_fce, _gc := w.Write(_daa._ade)
	if _gc != nil {
		return 0, _db.Wrap(_gc, _bfa, "\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073")
	}
	_aed += int64(_fce)
	return _aed, nil
}
func (_bc *Encoder) Final() { _bc.flush() }

var _ceg = []state{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

func (_gd *Encoder) code0(_dg *codingContext, _aag uint32, _aac uint16, _bfg byte) {
	if _dg.mps(_aag) == 0 {
		_gd.codeMPS(_dg, _aag, _aac, _bfg)
	} else {
		_gd.codeLPS(_dg, _aag, _aac, _bfg)
	}
}
func (_bf *Encoder) EncodeInteger(proc Class, value int) (_eed error) {
	_e.Log.Trace("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027", value, proc)
	if _eed = _bf.encodeInteger(proc, value); _eed != nil {
		return _db.Wrap(_eed, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072", "")
	}
	return nil
}
func (_egf *Encoder) EncodeIAID(symbolCodeLength, value int) (_af error) {
	_e.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", symbolCodeLength, value)
	if _af = _egf.encodeIAID(symbolCodeLength, value); _af != nil {
		return _db.Wrap(_af, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044", "")
	}
	return nil
}
func (_gbc *Encoder) dataSize() int { return _dacd*len(_gbc._cfd) + _gbc._cb }
func (_gfg *Encoder) EncodeOOB(proc Class) (_fa error) {
	_e.Log.Trace("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027", proc)
	if _fa = _gfg.encodeOOB(proc); _fa != nil {
		return _db.Wrap(_fa, "\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB", "")
	}
	return nil
}
func (_efbg *Encoder) emit() {
	if _efbg._cb == _dacd {
		_efbg._cfd = append(_efbg._cfd, _efbg._ade)
		_efbg._ade = make([]byte, _dacd)
		_efbg._cb = 0
	}
	_efbg._ade[_efbg._cb] = _efbg._fc
	_efbg._cb++
}
func (_gga *Encoder) code1(_afg *codingContext, _efb uint32, _efg uint16, _eae byte) {
	if _afg.mps(_efb) == 1 {
		_gga.codeMPS(_afg, _efb, _efg, _eae)
	} else {
		_gga.codeLPS(_afg, _efb, _efg, _eae)
	}
}
func (_abf *Encoder) codeMPS(_aedd *codingContext, _cefb uint32, _ecc uint16, _gbf byte) {
	_abf._adc -= _ecc
	if _abf._adc&0x8000 != 0 {
		_abf._g += uint32(_ecc)
		return
	}
	if _abf._adc < _ecc {
		_abf._adc = _ecc
	} else {
		_abf._g += uint32(_ecc)
	}
	_aedd._fe[_cefb] = _ceg[_gbf]._gca
	_abf.renormalize()
}
func (_ace *Encoder) Refine(iTemp, iTarget *_d.Bitmap, ox, oy int) error {
	for _bcf := 0; _bcf < iTarget.Height; _bcf++ {
		var _bdg int
		_badb := _bcf + oy
		var (
			_acb, _gad, _ec, _fb, _bac  uint16
			_ag, _bge, _cbd, _ccc, _ecg byte
		)
		if _badb >= 1 && (_badb-1) < iTemp.Height {
			_ag = iTemp.Data[(_badb-1)*iTemp.RowStride]
		}
		if _badb >= 0 && _badb < iTemp.Height {
			_bge = iTemp.Data[_badb*iTemp.RowStride]
		}
		if _badb >= -1 && _badb+1 < iTemp.Height {
			_cbd = iTemp.Data[(_badb+1)*iTemp.RowStride]
		}
		if _bcf >= 1 {
			_ccc = iTarget.Data[(_bcf-1)*iTarget.RowStride]
		}
		_ecg = iTarget.Data[_bcf*iTarget.RowStride]
		_ffe := uint(6 + ox)
		_acb = uint16(_ag >> _ffe)
		_gad = uint16(_bge >> _ffe)
		_ec = uint16(_cbd >> _ffe)
		_fb = uint16(_ccc >> 6)
		_dac := uint(2 - ox)
		_ag <<= _dac
		_bge <<= _dac
		_cbd <<= _dac
		_ccc <<= 2
		for _bdg = 0; _bdg < iTarget.Width; _bdg++ {
			_cgf := (_acb << 10) | (_gad << 7) | (_ec << 4) | (_fb << 1) | _bac
			_dcb := _ecg >> 7
			_aaa := _ace.encodeBit(_ace._egb, uint32(_cgf), _dcb)
			if _aaa != nil {
				return _aaa
			}
			_acb <<= 1
			_gad <<= 1
			_ec <<= 1
			_fb <<= 1
			_acb |= uint16(_ag >> 7)
			_gad |= uint16(_bge >> 7)
			_ec |= uint16(_cbd >> 7)
			_fb |= uint16(_ccc >> 7)
			_bac = uint16(_dcb)
			_bff := _bdg % 8
			_bgea := _bdg/8 + 1
			if _bff == 5+ox {
				_ag, _bge, _cbd = 0, 0, 0
				if _bgea < iTemp.RowStride && _badb >= 1 && (_badb-1) < iTemp.Height {
					_ag = iTemp.Data[(_badb-1)*iTemp.RowStride+_bgea]
				}
				if _bgea < iTemp.RowStride && _badb >= 0 && _badb < iTemp.Height {
					_bge = iTemp.Data[_badb*iTemp.RowStride+_bgea]
				}
				if _bgea < iTemp.RowStride && _badb >= -1 && (_badb+1) < iTemp.Height {
					_cbd = iTemp.Data[(_badb+1)*iTemp.RowStride+_bgea]
				}
			} else {
				_ag <<= 1
				_bge <<= 1
				_cbd <<= 1
			}
			if _bff == 5 && _bcf >= 1 {
				_ccc = 0
				if _bgea < iTarget.RowStride {
					_ccc = iTarget.Data[(_bcf-1)*iTarget.RowStride+_bgea]
				}
			} else {
				_ccc <<= 1
			}
			if _bff == 7 {
				_ecg = 0
				if _bgea < iTarget.RowStride {
					_ecg = iTarget.Data[_bcf*iTarget.RowStride+_bgea]
				}
			} else {
				_ecg <<= 1
			}
			_acb &= 7
			_gad &= 7
			_ec &= 7
			_fb &= 7
		}
	}
	return nil
}
func _feb(_dc int) *codingContext {
	return &codingContext{_fe: make([]byte, _dc), _ce: make([]byte, _dc)}
}
func (_faec *Encoder) flush() {
	_faec.setBits()
	_faec._g <<= _faec._cg
	_faec.byteOut()
	_faec._g <<= _faec._cg
	_faec.byteOut()
	_faec.emit()
	if _faec._fc != 0xff {
		_faec._gb++
		_faec._fc = 0xff
		_faec.emit()
	}
	_faec._gb++
	_faec._fc = 0xac
	_faec._gb++
	_faec.emit()
}
func (_ffb *Encoder) encodeBit(_abb *codingContext, _bcb uint32, _eedf uint8) error {
	const _afd = "\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074"
	_ffb._gg++
	if _bcb >= uint32(len(_abb._fe)) {
		return _db.Errorf(_afd, "\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027", _bcb)
	}
	_ccd := _abb._fe[_bcb]
	_gbg := _abb.mps(_bcb)
	_dae := _ceg[_ccd]._eee
	_e.Log.Trace("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064", _ffb._gg, _eedf, _ccd, _gbg, _dae, _ffb._adc, _ffb._g, _ffb._cg, _ffb._fc, _ffb._gb)
	if _eedf == 0 {
		_ffb.code0(_abb, _bcb, _dae, _ccd)
	} else {
		_ffb.code1(_abb, _bcb, _dae, _ccd)
	}
	return nil
}

var _ _ba.WriterTo = &Encoder{}

type state struct {
	_eee        uint16
	_gca, _efda uint8
	_aba        uint8
}

func (_fcee *Encoder) rBlock() {
	if _fcee._gb >= 0 {
		_fcee.emit()
	}
	_fcee._gb++
	_fcee._fc = uint8(_fcee._g >> 20)
	_fcee._g &= 0xfffff
	_fcee._cg = 7
}

type intEncRangeS struct {
	_ac, _ad int
	_f, _ea  uint8
	_bd      uint16
	_eg      uint8
}

func (_bgc *Encoder) lBlock() {
	if _bgc._gb >= 0 {
		_bgc.emit()
	}
	_bgc._gb++
	_bgc._fc = uint8(_bgc._g >> 19)
	_bgc._g &= 0x7ffff
	_bgc._cg = 8
}
func (_ae *Encoder) DataSize() int { return _ae.dataSize() }

type Class int

func (_dgb *Encoder) encodeOOB(_eef Class) error {
	_fae := _dgb._bg[_eef]
	_bfb := _dgb.encodeBit(_fae, 1, 1)
	if _bfb != nil {
		return _bfb
	}
	_bfb = _dgb.encodeBit(_fae, 3, 0)
	if _bfb != nil {
		return _bfb
	}
	_bfb = _dgb.encodeBit(_fae, 6, 0)
	if _bfb != nil {
		return _bfb
	}
	_bfb = _dgb.encodeBit(_fae, 12, 0)
	if _bfb != nil {
		return _bfb
	}
	return nil
}
func (_fcc *Encoder) Flush() {
	_fcc._cb = 0
	_fcc._cfd = nil
	_fcc._gb = -1
}
func (_cef *codingContext) mps(_de uint32) int { return int(_cef._ce[_de]) }
func (_ee *Encoder) EncodeBitmap(bm *_d.Bitmap, duplicateLineRemoval bool) error {
	_e.Log.Trace("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s", bm.Width, bm.Height, bm)
	var (
		_cbf, _da        uint8
		_ga, _febe, _aeb uint16
		_aa, _gaa, _ab   byte
		_gf, _dcd, _egg  int
		_ff, _eaf        []byte
	)
	for _cc := 0; _cc < bm.Height; _cc++ {
		_aa, _gaa = 0, 0
		if _cc >= 2 {
			_aa = bm.Data[(_cc-2)*bm.RowStride]
		}
		if _cc >= 1 {
			_gaa = bm.Data[(_cc-1)*bm.RowStride]
			if duplicateLineRemoval {
				_dcd = _cc * bm.RowStride
				_ff = bm.Data[_dcd : _dcd+bm.RowStride]
				_egg = (_cc - 1) * bm.RowStride
				_eaf = bm.Data[_egg : _egg+bm.RowStride]
				if _a.Equal(_ff, _eaf) {
					_da = _cbf ^ 1
					_cbf = 1
				} else {
					_da = _cbf
					_cbf = 0
				}
			}
		}
		if duplicateLineRemoval {
			if _fea := _ee.encodeBit(_ee._egb, _efd, _da); _fea != nil {
				return _fea
			}
			if _cbf != 0 {
				continue
			}
		}
		_ab = bm.Data[_cc*bm.RowStride]
		_ga = uint16(_aa >> 5)
		_febe = uint16(_gaa >> 4)
		_aa <<= 3
		_gaa <<= 4
		_aeb = 0
		for _gf = 0; _gf < bm.Width; _gf++ {
			_dbg := uint32(_ga<<11 | _febe<<4 | _aeb)
			_aeg := (_ab & 0x80) >> 7
			_ge := _ee.encodeBit(_ee._egb, _dbg, _aeg)
			if _ge != nil {
				return _ge
			}
			_ga <<= 1
			_febe <<= 1
			_aeb <<= 1
			_ga |= uint16((_aa & 0x80) >> 7)
			_febe |= uint16((_gaa & 0x80) >> 7)
			_aeb |= uint16(_aeg)
			_fg := _gf % 8
			_edc := _gf/8 + 1
			if _fg == 4 && _cc >= 2 {
				_aa = 0
				if _edc < bm.RowStride {
					_aa = bm.Data[(_cc-2)*bm.RowStride+_edc]
				}
			} else {
				_aa <<= 1
			}
			if _fg == 3 && _cc >= 1 {
				_gaa = 0
				if _edc < bm.RowStride {
					_gaa = bm.Data[(_cc-1)*bm.RowStride+_edc]
				}
			} else {
				_gaa <<= 1
			}
			if _fg == 7 {
				_ab = 0
				if _edc < bm.RowStride {
					_ab = bm.Data[_cc*bm.RowStride+_edc]
				}
			} else {
				_ab <<= 1
			}
			_ga &= 31
			_febe &= 127
			_aeb &= 15
		}
	}
	return nil
}
func (_c Class) String() string {
	switch _c {
	case IAAI:
		return "\u0049\u0041\u0041\u0049"
	case IADH:
		return "\u0049\u0041\u0044\u0048"
	case IADS:
		return "\u0049\u0041\u0044\u0053"
	case IADT:
		return "\u0049\u0041\u0044\u0054"
	case IADW:
		return "\u0049\u0041\u0044\u0057"
	case IAEX:
		return "\u0049\u0041\u0045\u0058"
	case IAFS:
		return "\u0049\u0041\u0046\u0053"
	case IAIT:
		return "\u0049\u0041\u0049\u0054"
	case IARDH:
		return "\u0049\u0041\u0052D\u0048"
	case IARDW:
		return "\u0049\u0041\u0052D\u0057"
	case IARDX:
		return "\u0049\u0041\u0052D\u0058"
	case IARDY:
		return "\u0049\u0041\u0052D\u0059"
	case IARI:
		return "\u0049\u0041\u0052\u0049"
	default:
		return "\u0055N\u004b\u004e\u004f\u0057\u004e"
	}
}
