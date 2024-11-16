package bitwise

import (
	_a "encoding/binary"
	_g "errors"
	_db "fmt"
	_c "io"

	_ga "github.com/bamzi/pdfext/common"
	_gd "github.com/bamzi/pdfext/internal/jbig2/errors"
)

func (_bag *Writer) writeBit(_ffgg uint8) error {
	if len(_bag._da)-1 < _bag._eead {
		return _c.EOF
	}
	_gdf := _bag._cfce
	if _bag._gde {
		_gdf = 7 - _bag._cfce
	}
	_bag._da[_bag._eead] |= byte(uint16(_ffgg<<_gdf) & 0xff)
	_bag._cfce++
	if _bag._cfce == 8 {
		_bag._eead++
		_bag._cfce = 0
	}
	return nil
}
func (_eff *Reader) ConsumeRemainingBits() (uint64, error) {
	if _eff._aed != 0 {
		return _eff.ReadBits(_eff._aed)
	}
	return 0, nil
}
func (_effc *Reader) ReadBool() (bool, error) { return _effc.readBool() }
func (_agd *Reader) NewPartialReader(offset, length int, relative bool) (*Reader, error) {
	if offset < 0 {
		return nil, _g.New("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065")
	}
	if relative {
		offset = _agd._cga._ce + offset
	}
	if length > 0 {
		_fgf := len(_agd._cga._ffg)
		if relative {
			_fgf = _agd._cga._gbb
		}
		if offset+length > _fgf {
			return nil, _db.Errorf("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064", offset, length, offset+length, _agd._cga._gbb)
		}
	}
	if length < 0 {
		_fcc := len(_agd._cga._ffg)
		if relative {
			_fcc = _agd._cga._gbb
		}
		length = _fcc - offset
	}
	return &Reader{_cga: readerSource{_ffg: _agd._cga._ffg, _gbb: length, _ce: offset}}, nil
}
func BufferedMSB() *BufferedWriter { return &BufferedWriter{_cd: true} }
func (_eb *BufferedWriter) Write(d []byte) (int, error) {
	_eb.expandIfNeeded(len(d))
	if _eb._gaf == 0 {
		return _eb.writeFullBytes(d), nil
	}
	return _eb.writeShiftedBytes(d), nil
}
func (_cgf *Reader) Reset() {
	_cgf._dbdb = _cgf._cab
	_cgf._aed = _cgf._abd
	_cgf._egc = _cgf._gdc
	_cgf._fd = _cgf._fdg
}
func (_fac *Writer) WriteBits(bits uint64, number int) (_ebg int, _eae error) {
	const _eef = "\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073"
	if number < 0 || number > 64 {
		return 0, _gd.Errorf(_eef, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	if number == 0 {
		return 0, nil
	}
	_cdc := number / 8
	if _cdc > 0 {
		_fbcf := number - _cdc*8
		for _ceb := _cdc - 1; _ceb >= 0; _ceb-- {
			_bcg := byte((bits >> uint(_ceb*8+_fbcf)) & 0xff)
			if _eae = _fac.WriteByte(_bcg); _eae != nil {
				return _ebg, _gd.Wrapf(_eae, _eef, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _cdc-_ceb+1)
			}
		}
		number -= _cdc * 8
		if number == 0 {
			return _cdc, nil
		}
	}
	var _gfa int
	for _ggfd := 0; _ggfd < number; _ggfd++ {
		if _fac._gde {
			_gfa = int((bits >> uint(number-1-_ggfd)) & 0x1)
		} else {
			_gfa = int(bits & 0x1)
			bits >>= 1
		}
		if _eae = _fac.WriteBit(_gfa); _eae != nil {
			return _ebg, _gd.Wrapf(_eae, _eef, "\u0062i\u0074\u003a\u0020\u0025\u0064", _ggfd)
		}
	}
	return _cdc, nil
}

var _ _c.ByteWriter = &BufferedWriter{}

func (_dg *Reader) Seek(offset int64, whence int) (int64, error) {
	_dg._cc = -1
	_dg._aed = 0
	_dg._egc = 0
	_dg._fd = 0
	var _ecg int64
	switch whence {
	case _c.SeekStart:
		_ecg = offset
	case _c.SeekCurrent:
		_ecg = _dg._dbdb + offset
	case _c.SeekEnd:
		_ecg = int64(_dg._cga._gbb) + offset
	default:
		return 0, _g.New("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	if _ecg < 0 {
		return 0, _g.New("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e")
	}
	_dg._dbdb = _ecg
	_dg._aed = 0
	return _ecg, nil
}
func (_cfa *BufferedWriter) writeFullBytes(_bcb []byte) int {
	_fe := copy(_cfa._cf[_cfa.fullOffset():], _bcb)
	_cfa._ef += _fe
	return _fe
}
func (_cdgg *Writer) Write(p []byte) (int, error) {
	if len(p) > _cdgg.byteCapacity() {
		return 0, _c.EOF
	}
	for _, _eac := range p {
		if _agb := _cdgg.writeByte(_eac); _agb != nil {
			return 0, _agb
		}
	}
	return len(p), nil
}
func NewWriter(data []byte) *Writer { return &Writer{_da: data} }
func (_ab *BufferedWriter) fullOffset() int {
	_ae := _ab._ef
	if _ab._gaf != 0 {
		_ae++
	}
	return _ae
}

type readerSource struct {
	_ffg []byte
	_ce  int
	_gbb int
}

const (
	_e  = 64
	_gc = int(^uint(0) >> 1)
)

type Reader struct {
	_cga  readerSource
	_egc  byte
	_aed  byte
	_dbdb int64
	_fd   int
	_cc   int
	_cab  int64
	_abd  byte
	_gdc  byte
	_fdg  int
}

func (_eaca *Writer) WriteByte(c byte) error { return _eaca.writeByte(c) }
func (_gag *BufferedWriter) writeByte(_bbd byte) {
	switch {
	case _gag._gaf == 0:
		_gag._cf[_gag._ef] = _bbd
		_gag._ef++
	case _gag._cd:
		_gag._cf[_gag._ef] |= _bbd >> _gag._gaf
		_gag._ef++
		_gag._cf[_gag._ef] = byte(uint16(_bbd) << (8 - _gag._gaf) & 0xff)
	default:
		_gag._cf[_gag._ef] |= byte(uint16(_bbd) << _gag._gaf & 0xff)
		_gag._ef++
		_gag._cf[_gag._ef] = _bbd >> (8 - _gag._gaf)
	}
}
func (_bab *BufferedWriter) byteCapacity() int {
	_efc := len(_bab._cf) - _bab._ef
	if _bab._gaf != 0 {
		_efc--
	}
	return _efc
}
func (_gca *BufferedWriter) WriteBits(bits uint64, number int) (_dbd int, _ea error) {
	const _eab = "\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s"
	if number < 0 || number > 64 {
		return 0, _gd.Errorf(_eab, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	_ba := number / 8
	if _ba > 0 {
		_gcf := number - _ba*8
		for _cb := _ba - 1; _cb >= 0; _cb-- {
			_ca := byte((bits >> uint(_cb*8+_gcf)) & 0xff)
			if _ea = _gca.WriteByte(_ca); _ea != nil {
				return _dbd, _gd.Wrapf(_ea, _eab, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _ba-_cb+1)
			}
		}
		number -= _ba * 8
		if number == 0 {
			return _ba, nil
		}
	}
	var _cfb int
	for _bac := 0; _bac < number; _bac++ {
		if _gca._cd {
			_cfb = int((bits >> uint(number-1-_bac)) & 0x1)
		} else {
			_cfb = int(bits & 0x1)
			bits >>= 1
		}
		if _ea = _gca.WriteBit(_cfb); _ea != nil {
			return _dbd, _gd.Wrapf(_ea, _eab, "\u0062i\u0074\u003a\u0020\u0025\u0064", _bac)
		}
	}
	return _ba, nil
}
func (_ebc *Writer) UseMSB() bool { return _ebc._gde }
func (_gbd *Writer) FinishByte() {
	if _gbd._cfce == 0 {
		return
	}
	_gbd._cfce = 0
	_gbd._eead++
}

var (
	_ _c.Reader     = &Reader{}
	_ _c.ByteReader = &Reader{}
	_ _c.Seeker     = &Reader{}
	_ StreamReader  = &Reader{}
)

func (_bb *BufferedWriter) SkipBits(skip int) error {
	if skip == 0 {
		return nil
	}
	_cg := int(_bb._gaf) + skip
	if _cg >= 0 && _cg < 8 {
		_bb._gaf = uint8(_cg)
		return nil
	}
	_cg = int(_bb._gaf) + _bb._ef*8 + skip
	if _cg < 0 {
		return _gd.Errorf("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_eg := _cg / 8
	_bbc := _cg % 8
	_bb._gaf = uint8(_bbc)
	if _gda := _eg - _bb._ef; _gda > 0 && len(_bb._cf)-1 < _eg {
		if _bb._gaf != 0 {
			_gda++
		}
		_bb.expandIfNeeded(_gda)
	}
	_bb._ef = _eg
	return nil
}
func (_ggda *Reader) AbsolutePosition() int64 { return _ggda._dbdb + int64(_ggda._cga._ce) }
func NewReader(data []byte) *Reader {
	return &Reader{_cga: readerSource{_ffg: data, _gbb: len(data), _ce: 0}}
}

var _ BinaryWriter = &BufferedWriter{}

func (_bc *BufferedWriter) WriteByte(bt byte) error {
	if _bc._ef > len(_bc._cf)-1 || (_bc._ef == len(_bc._cf)-1 && _bc._gaf != 0) {
		_bc.expandIfNeeded(1)
	}
	_bc.writeByte(bt)
	return nil
}
func (_cdg *BufferedWriter) expandIfNeeded(_fc int) {
	if !_cdg.tryGrowByReslice(_fc) {
		_cdg.grow(_fc)
	}
}
func (_gec *Reader) ReadBit() (_cce int, _bbb error) {
	_ead, _bbb := _gec.readBool()
	if _bbb != nil {
		return 0, _bbb
	}
	if _ead {
		_cce = 1
	}
	return _cce, nil
}
func NewWriterMSB(data []byte) *Writer { return &Writer{_da: data, _gde: true} }
func (_dbba *Writer) Data() []byte     { return _dbba._da }

type BitWriter interface {
	WriteBit(_gb int) error
	WriteBits(_ag uint64, _abe int) (_ggd int, _ffb error)
	FinishByte()
	SkipBits(_dea int) error
}

func (_gcb *Reader) RelativePosition() int64 { return _gcb._dbdb }
func (_eee *Reader) Read(p []byte) (_efaa int, _ggdg error) {
	if _eee._aed == 0 {
		return _eee.read(p)
	}
	for ; _efaa < len(p); _efaa++ {
		if p[_efaa], _ggdg = _eee.readUnalignedByte(); _ggdg != nil {
			return 0, _ggdg
		}
	}
	return _efaa, nil
}
func (_dd *Reader) ReadByte() (byte, error) {
	if _dd._aed == 0 {
		return _dd.readBufferByte()
	}
	return _dd.readUnalignedByte()
}
func (_bge *Writer) ResetBit() { _bge._cfce = 0 }
func (_aeg *Reader) ReadBits(n byte) (_afc uint64, _bca error) {
	if n < _aeg._aed {
		_bgd := _aeg._aed - n
		_afc = uint64(_aeg._egc >> _bgd)
		_aeg._egc &= 1<<_bgd - 1
		_aeg._aed = _bgd
		return _afc, nil
	}
	if n > _aeg._aed {
		if _aeg._aed > 0 {
			_afc = uint64(_aeg._egc)
			n -= _aeg._aed
		}
		for n >= 8 {
			_fcf, _fa := _aeg.readBufferByte()
			if _fa != nil {
				return 0, _fa
			}
			_afc = _afc<<8 + uint64(_fcf)
			n -= 8
		}
		if n > 0 {
			if _aeg._egc, _bca = _aeg.readBufferByte(); _bca != nil {
				return 0, _bca
			}
			_df := 8 - n
			_afc = _afc<<n + uint64(_aeg._egc>>_df)
			_aeg._egc &= 1<<_df - 1
			_aeg._aed = _df
		} else {
			_aeg._aed = 0
		}
		return _afc, nil
	}
	_aeg._aed = 0
	return uint64(_aeg._egc), nil
}
func (_ccg *Reader) read(_edf []byte) (int, error) {
	if _ccg._dbdb >= int64(_ccg._cga._gbb) {
		return 0, _c.EOF
	}
	_ccg._cc = -1
	_dc := copy(_edf, _ccg._cga._ffg[(int64(_ccg._cga._ce)+_ccg._dbdb):(_ccg._cga._ce+_ccg._cga._gbb)])
	_ccg._dbdb += int64(_dc)
	return _dc, nil
}
func (_bcf *Writer) SkipBits(skip int) error {
	const _aa = "\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073"
	if skip == 0 {
		return nil
	}
	_fcb := int(_bcf._cfce) + skip
	if _fcb >= 0 && _fcb < 8 {
		_bcf._cfce = uint8(_fcb)
		return nil
	}
	_fcb = int(_bcf._cfce) + _bcf._eead*8 + skip
	if _fcb < 0 {
		return _gd.Errorf(_aa, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_fcg := _fcb / 8
	_ded := _fcb % 8
	_ga.Log.Trace("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073")
	_ga.Log.Trace("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027", _bcf._cfce, _bcf._eead, int(_bcf._cfce)+(_bcf._eead)*8, len(_bcf._da), cap(_bcf._da))
	_ga.Log.Trace("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027", skip, _fcb, _ded)
	_bcf._cfce = uint8(_ded)
	if _ad := _fcg - _bcf._eead; _ad > 0 && len(_bcf._da)-1 < _fcg {
		_ga.Log.Trace("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064", _ad)
		return _gd.Errorf(_aa, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bcf._eead = _fcg
	_ga.Log.Trace("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027", _bcf._cfce, _bcf._eead)
	return nil
}
func (_b *BufferedWriter) Len() int { return _b.byteCapacity() }
func (_dcg *Reader) readUnalignedByte() (_gac byte, _fff error) {
	_ccf := _dcg._aed
	_gac = _dcg._egc << (8 - _ccf)
	_dcg._egc, _fff = _dcg.readBufferByte()
	if _fff != nil {
		return 0, _fff
	}
	_gac |= _dcg._egc >> _ccf
	_dcg._egc &= 1<<_ccf - 1
	return _gac, nil
}

type StreamReader interface {
	_c.Reader
	_c.ByteReader
	_c.Seeker
	Align() byte
	BitPosition() int
	Mark()
	Length() uint64
	ReadBit() (int, error)
	ReadBits(_dbb byte) (uint64, error)
	ReadBool() (bool, error)
	ReadUint32() (uint32, error)
	Reset()
	AbsolutePosition() int64
}

func (_gfbc *Reader) readBool() (_eda bool, _egcf error) {
	if _gfbc._aed == 0 {
		_gfbc._egc, _egcf = _gfbc.readBufferByte()
		if _egcf != nil {
			return false, _egcf
		}
		_eda = (_gfbc._egc & 0x80) != 0
		_gfbc._egc, _gfbc._aed = _gfbc._egc&0x7f, 7
		return _eda, nil
	}
	_gfbc._aed--
	_eda = (_gfbc._egc & (1 << _gfbc._aed)) != 0
	_gfbc._egc &= 1<<_gfbc._aed - 1
	return _eda, nil
}
func (_ed *Reader) AbsoluteLength() uint64 { return uint64(len(_ed._cga._ffg)) }
func (_af *BufferedWriter) Data() []byte   { return _af._cf }
func (_ebf *BufferedWriter) WriteBit(bit int) error {
	if bit != 1 && bit != 0 {
		return _gd.Errorf("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064", bit)
	}
	if len(_ebf._cf)-1 < _ebf._ef {
		_ebf.expandIfNeeded(1)
	}
	_gce := _ebf._gaf
	if _ebf._cd {
		_gce = 7 - _ebf._gaf
	}
	_ebf._cf[_ebf._ef] |= byte(uint16(bit<<_gce) & 0xff)
	_ebf._gaf++
	if _ebf._gaf == 8 {
		_ebf._ef++
		_ebf._gaf = 0
	}
	return nil
}

type BufferedWriter struct {
	_cf  []byte
	_gaf uint8
	_ef  int
	_cd  bool
}

func (_gfb *Reader) BitPosition() int       { return int(_gfb._aed) }
func (_cad *Reader) Align() (_feg byte)     { _feg = _cad._aed; _cad._aed = 0; return _feg }
func (_cfg *BufferedWriter) ResetBitIndex() { _cfg._gaf = 0 }
func (_cac *Reader) Length() uint64         { return uint64(_cac._cga._gbb) }
func (_cfcf *Reader) Mark() {
	_cfcf._cab = _cfcf._dbdb
	_cfcf._abd = _cfcf._aed
	_cfcf._gdc = _cfcf._egc
	_cfcf._fdg = _cfcf._fd
}
func (_gf *BufferedWriter) grow(_ff int) {
	if _gf._cf == nil && _ff < _e {
		_gf._cf = make([]byte, _ff, _e)
		return
	}
	_caa := len(_gf._cf)
	if _gf._gaf != 0 {
		_caa++
	}
	_gg := cap(_gf._cf)
	switch {
	case _ff <= _gg/2-_caa:
		_ga.Log.Trace("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027", len(_gf._cf), cap(_gf._cf), _ff)
		_ga.Log.Trace("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027", _gg, _caa)
		copy(_gf._cf, _gf._cf[_gf.fullOffset():])
	case _gg > _gc-_gg-_ff:
		_ga.Log.Error("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065")
		return
	default:
		_de := make([]byte, 2*_gg+_ff)
		copy(_de, _gf._cf)
		_gf._cf = _de
	}
	_gf._cf = _gf._cf[:_caa+_ff]
}
func (_cfc *BufferedWriter) Reset() { _cfc._cf = _cfc._cf[:0]; _cfc._ef = 0; _cfc._gaf = 0 }

var _ _c.Writer = &BufferedWriter{}

func (_bg *BufferedWriter) writeShiftedBytes(_gab []byte) int {
	for _, _bf := range _gab {
		_bg.writeByte(_bf)
	}
	return len(_gab)
}

var _ BinaryWriter = &Writer{}

func (_egb *Reader) ReadUint32() (uint32, error) {
	_ced := make([]byte, 4)
	_, _eea := _egb.Read(_ced)
	if _eea != nil {
		return 0, _eea
	}
	return _a.BigEndian.Uint32(_ced), nil
}
func (_efa *BufferedWriter) tryGrowByReslice(_fg int) bool {
	if _fb := len(_efa._cf); _fg <= cap(_efa._cf)-_fb {
		_efa._cf = _efa._cf[:_fb+_fg]
		return true
	}
	return false
}
func (_aga *Writer) writeByte(_ggff byte) error {
	if _aga._eead > len(_aga._da)-1 {
		return _c.EOF
	}
	if _aga._eead == len(_aga._da)-1 && _aga._cfce != 0 {
		return _c.EOF
	}
	if _aga._cfce == 0 {
		_aga._da[_aga._eead] = _ggff
		_aga._eead++
		return nil
	}
	if _aga._gde {
		_aga._da[_aga._eead] |= _ggff >> _aga._cfce
		_aga._eead++
		_aga._da[_aga._eead] = byte(uint16(_ggff) << (8 - _aga._cfce) & 0xff)
	} else {
		_aga._da[_aga._eead] |= byte(uint16(_ggff) << _aga._cfce & 0xff)
		_aga._eead++
		_aga._da[_aga._eead] = _ggff >> (8 - _aga._cfce)
	}
	return nil
}
func (_gaba *Reader) readBufferByte() (byte, error) {
	if _gaba._dbdb >= int64(_gaba._cga._gbb) {
		return 0, _c.EOF
	}
	_gaba._cc = -1
	_gdcd := _gaba._cga._ffg[int64(_gaba._cga._ce)+_gaba._dbdb]
	_gaba._dbdb++
	_gaba._fd = int(_gdcd)
	return _gdcd, nil
}

type BinaryWriter interface {
	BitWriter
	_c.Writer
	_c.ByteWriter
	Data() []byte
}

func (_deb *Writer) WriteBit(bit int) error {
	switch bit {
	case 0, 1:
		return _deb.writeBit(uint8(bit))
	}
	return _gd.Error("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065")
}
func (_cgag *Writer) byteCapacity() int {
	_cdb := len(_cgag._da) - _cgag._eead
	if _cgag._cfce != 0 {
		_cdb--
	}
	return _cdb
}
func (_afg *BufferedWriter) FinishByte() {
	if _afg._gaf == 0 {
		return
	}
	_afg._gaf = 0
	_afg._ef++
}

type Writer struct {
	_da   []byte
	_cfce uint8
	_eead int
	_gde  bool
}
