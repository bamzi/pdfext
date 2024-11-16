package crypt

import (
	_g "crypto/aes"
	_abd "crypto/cipher"
	_b "crypto/md5"
	_d "crypto/rand"
	_ab "crypto/rc4"
	_a "fmt"
	_f "io"

	_fa "github.com/bamzi/pdfext/common"
	_ag "github.com/bamzi/pdfext/core/security"
)

func init() { _aba("\u0041\u0045\u0053V\u0032", _ee) }
func _bgb(_gd FilterDict) (Filter, error) {
	if _gd.Length%8 != 0 {
		return nil, _a.Errorf("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029", _gd.Length)
	}
	if _gd.Length < 5 || _gd.Length > 16 {
		if _gd.Length == 40 || _gd.Length == 64 || _gd.Length == 128 {
			_fa.Log.Debug("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029", _gd.Length)
			_gd.Length /= 8
		} else {
			return nil, _a.Errorf("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029", _gd.Length)
		}
	}
	return filterV2{_ea: _gd.Length}, nil
}

// EncryptBytes implements Filter interface.
func (filterV2) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_gg, _aga := _ab.NewCipher(okey)
	if _aga != nil {
		return nil, _aga
	}
	_fa.Log.Trace("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_gg.XORKeyStream(buf, buf)
	_fa.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2(length int) Filter {
	_fed, _gcb := _bgb(FilterDict{Length: length})
	if _gcb != nil {
		_fa.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _gcb)
		return filterV2{_ea: length}
	}
	return _fed
}

// Name implements Filter interface.
func (filterAESV2) Name() string { return "\u0041\u0045\u0053V\u0032" }
func _abf(_dd FilterDict) (Filter, error) {
	if _dd.Length == 256 {
		_fa.Log.Debug("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _dd.Length)
		_dd.Length /= 8
	}
	if _dd.Length != 0 && _dd.Length != 32 {
		return nil, _a.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _dd.Length)
	}
	return filterAESV3{}, nil
}

// KeyLength implements Filter interface.
func (filterAESV3) KeyLength() int { return 256 / 8 }

// Name implements Filter interface.
func (filterV2) Name() string { return "\u0056\u0032" }

type filterAESV3 struct{ filterAES }

func (filterAES) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_de, _fe := _g.NewCipher(okey)
	if _fe != nil {
		return nil, _fe
	}
	_fa.Log.Trace("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	const _eg = _g.BlockSize
	_fee := _eg - len(buf)%_eg
	for _cge := 0; _cge < _fee; _cge++ {
		buf = append(buf, byte(_fee))
	}
	_fa.Log.Trace("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073", len(buf))
	_fg := make([]byte, _eg+len(buf))
	_df := _fg[:_eg]
	if _, _ad := _f.ReadFull(_d.Reader, _df); _ad != nil {
		return nil, _ad
	}
	_db := _abd.NewCBCEncrypter(_de, _df)
	_db.CryptBlocks(_fg[_eg:], buf)
	buf = _fg
	_fa.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	return buf, nil
}
func init() { _aba("\u0041\u0045\u0053V\u0033", _abf) }

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2() Filter {
	_ef, _da := _ee(FilterDict{})
	if _da != nil {
		_fa.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _da)
		return filterAESV2{}
	}
	return _ef
}

// KeyLength implements Filter interface.
func (filterAESV2) KeyLength() int { return 128 / 8 }

// HandlerVersion implements Filter interface.
func (filterAESV3) HandlerVersion() (V, R int) { V, R = 5, 6; return }

type filterIdentity struct{}

// PDFVersion implements Filter interface.
func (_deb filterV2) PDFVersion() [2]int                                  { return [2]int{} }
func (filterIdentity) EncryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct {
	CFM       string
	AuthEvent _ag.AuthEvent
	Length    int
}

var (
	_egf = make(map[string]filterFunc)
)

// MakeKey implements Filter interface.
func (filterAESV3) MakeKey(_, _ uint32, ekey []byte) ([]byte, error) { return ekey, nil }

// KeyLength implements Filter interface.
func (_agf filterV2) KeyLength() int { return _agf._ea }

// Name implements Filter interface.
func (filterAESV3) Name() string { return "\u0041\u0045\u0053V\u0033" }

type filterAESV2 struct{ filterAES }

func init() { _aba("\u0056\u0032", _bgb) }

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter(d FilterDict) (Filter, error) {
	_afd, _aae := _ded(d.CFM)
	if _aae != nil {
		return nil, _aae
	}
	_cdf, _aae := _afd(d)
	if _aae != nil {
		return nil, _aae
	}
	return _cdf, nil
}
func _aba(_edg string, _gdb filterFunc) {
	if _, _geg := _egf[_edg]; _geg {
		panic("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064")
	}
	_egf[_edg] = _gdb
}
func _ee(_bb FilterDict) (Filter, error) {
	if _bb.Length == 128 {
		_fa.Log.Debug("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _bb.Length)
		_bb.Length /= 8
	}
	if _bb.Length != 0 && _bb.Length != 16 {
		return nil, _a.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _bb.Length)
	}
	return filterAESV2{}, nil
}
func (filterIdentity) KeyLength() int { return 0 }

type filterFunc func(_bag FilterDict) (Filter, error)

// HandlerVersion implements Filter interface.
func (_ba filterV2) HandlerVersion() (V, R int) { V, R = 2, 3; return }

// MakeKey implements Filter interface.
func (_af filterV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bc(objNum, genNum, ekey, false)
}
func _ded(_afe string) (filterFunc, error) {
	_faf := _egf[_afe]
	if _faf == nil {
		return nil, _a.Errorf("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071", _afe)
	}
	return _faf, nil
}

var _ Filter = filterAESV2{}

// Filter is a common interface for crypt filter methods.
type Filter interface {

	// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
	Name() string

	// KeyLength returns a length of the encryption key in bytes.
	KeyLength() int

	// PDFVersion reports the minimal version of PDF document that introduced this filter.
	PDFVersion() [2]int

	// HandlerVersion reports V and R parameters that should be used for this filter.
	HandlerVersion() (V, R int)

	// MakeKey generates a object encryption key based on file encryption key and object numbers.
	// Used only for legacy filters - AESV3 doesn't change the key for each object.
	MakeKey(_gf, _ega uint32, _bbd []byte) ([]byte, error)

	// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and encrypt data in-place.
	EncryptBytes(_eab []byte, _gee []byte) ([]byte, error)

	// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and decrypt data in-place.
	DecryptBytes(_agc []byte, _deg []byte) ([]byte, error)
}

// DecryptBytes implements Filter interface.
func (filterV2) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_cea, _dc := _ab.NewCipher(okey)
	if _dc != nil {
		return nil, _dc
	}
	_fa.Log.Trace("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_cea.XORKeyStream(buf, buf)
	_fa.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

// PDFVersion implements Filter interface.
func (filterAESV2) PDFVersion() [2]int { return [2]int{1, 5} }

// MakeKey implements Filter interface.
func (filterAESV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bc(objNum, genNum, ekey, true)
}

type filterAES struct{}

func (filterAES) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_gc, _ec := _g.NewCipher(okey)
	if _ec != nil {
		return nil, _ec
	}
	if len(buf) < 16 {
		_fa.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073", buf)
		return buf, _a.Errorf("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029", len(buf))
	}
	_cc := buf[:16]
	buf = buf[16:]
	if len(buf)%16 != 0 {
		_fa.Log.Debug("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(_cc), _cc)
		_fa.Log.Debug("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
		return buf, _a.Errorf("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029", len(buf))
	}
	_fga := _abd.NewCBCDecrypter(_gc, _cc)
	_fa.Log.Trace("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	_fa.Log.Trace("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078", len(buf), buf)
	_fga.CryptBlocks(buf, buf)
	_fa.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	if len(buf) == 0 {
		_fa.Log.Trace("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067")
		return buf, nil
	}
	_eb := int(buf[len(buf)-1])
	if _eb > len(buf) {
		_fa.Log.Debug("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029", _eb, len(buf))
		return buf, _a.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068")
	}
	buf = buf[:len(buf)-_eb]
	return buf, nil
}

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3() Filter {
	_ge, _bg := _abf(FilterDict{})
	if _bg != nil {
		_fa.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _bg)
		return filterAESV3{}
	}
	return _ge
}

var _ Filter = filterAESV3{}

func (filterIdentity) HandlerVersion() (V, R int) { return }

var _ Filter = filterV2{}

func _bc(_ce, _cd uint32, _gef []byte, _ecb bool) ([]byte, error) {
	_def := make([]byte, len(_gef)+5)
	copy(_def, _gef)
	for _adg := 0; _adg < 3; _adg++ {
		_ga := byte((_ce >> uint32(8*_adg)) & 0xff)
		_def[_adg+len(_gef)] = _ga
	}
	for _dbf := 0; _dbf < 2; _dbf++ {
		_abb := byte((_cd >> uint32(8*_dbf)) & 0xff)
		_def[_dbf+len(_gef)+3] = _abb
	}
	if _ecb {
		_def = append(_def, 0x73)
		_def = append(_def, 0x41)
		_def = append(_def, 0x6C)
		_def = append(_def, 0x54)
	}
	_ae := _b.New()
	_ae.Write(_def)
	_gb := _ae.Sum(nil)
	if len(_gef)+5 < 16 {
		return _gb[0 : len(_gef)+5], nil
	}
	return _gb, nil
}

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity() Filter                                                         { return filterIdentity{} }
func (filterIdentity) MakeKey(objNum, genNum uint32, fkey []byte) ([]byte, error) { return fkey, nil }

// HandlerVersion implements Filter interface.
func (filterAESV2) HandlerVersion() (V, R int) { V, R = 4, 4; return }

// PDFVersion implements Filter interface.
func (filterAESV3) PDFVersion() [2]int                                    { return [2]int{2, 0} }
func (filterIdentity) PDFVersion() [2]int                                 { return [2]int{} }
func (filterIdentity) Name() string                                       { return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" }
func (filterIdentity) DecryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

type filterV2 struct{ _ea int }
