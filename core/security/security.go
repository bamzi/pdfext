package security

import (
	_cc "bytes"
	_fd "crypto/aes"
	_e "crypto/cipher"
	_bc "crypto/md5"
	_fa "crypto/rand"
	_c "crypto/rc4"
	_b "crypto/sha256"
	_ef "crypto/sha512"
	_gb "encoding/binary"
	_ec "errors"
	_ff "fmt"
	_f "hash"
	_a "io"
	_ge "math"

	_d "github.com/bamzi/pdfext/common"
)

var _ StdHandler = stdHandlerR6{}

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6() StdHandler      { return stdHandlerR6{} }
func _af(_ea _e.Block) _e.BlockMode { return (*ecbDecrypter)(_efe(_ea)) }
func (_fg *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_fg._ae != 0 {
		_d.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_d.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_fg._bca.Decrypt(dst, src[:_fg._ae])
		src = src[_fg._ae:]
		dst = dst[_fg._ae:]
	}
}
func (_bdf *ecbDecrypter) BlockSize() int { return _bdf._ae }

// AuthEvent is an event type that triggers authentication.
type AuthEvent string

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct {
	R               int
	P               Permissions
	EncryptMetadata bool
	O, U            []byte
	OE, UE          []byte
	Perms           []byte
}

const (
	EventDocOpen = AuthEvent("\u0044o\u0063\u004f\u0070\u0065\u006e")
	EventEFOpen  = AuthEvent("\u0045\u0046\u004f\u0070\u0065\u006e")
)
const (
	PermOwner             = Permissions(_ge.MaxUint32)
	PermPrinting          = Permissions(1 << 2)
	PermModify            = Permissions(1 << 3)
	PermExtractGraphics   = Permissions(1 << 4)
	PermAnnotate          = Permissions(1 << 5)
	PermFillForms         = Permissions(1 << 8)
	PermDisabilityExtract = Permissions(1 << 9)
	PermRotateInsert      = Permissions(1 << 10)
	PermFullPrintQuality  = Permissions(1 << 11)
)

func (_cfgc stdHandlerR6) alg12(_fag *StdEncryptDict, _edf []byte) ([]byte, error) {
	if _ffa := _bcf("\u0061\u006c\u00671\u0032", "\u0055", 48, _fag.U); _ffa != nil {
		return nil, _ffa
	}
	if _eeg := _bcf("\u0061\u006c\u00671\u0032", "\u004f", 48, _fag.O); _eeg != nil {
		return nil, _eeg
	}
	_aabg := make([]byte, len(_edf)+8+48)
	_bgc := copy(_aabg, _edf)
	_bgc += copy(_aabg[_bgc:], _fag.O[32:40])
	_bgc += copy(_aabg[_bgc:], _fag.U[0:48])
	_bda, _ada := _cfgc.alg2b(_fag.R, _aabg, _edf, _fag.U[0:48])
	if _ada != nil {
		return nil, _ada
	}
	_bda = _bda[:32]
	if !_cc.Equal(_bda, _fag.O[:32]) {
		return nil, nil
	}
	return _bda, nil
}

// Authenticate implements StdHandler interface.
func (_aaf stdHandlerR4) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	_d.Log.Trace("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073")
	_fffa, _adf := _aaf.alg7(d, pass)
	if _adf != nil {
		return nil, 0, _adf
	}
	if _fffa != nil {
		_d.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _fffa, PermOwner, nil
	}
	_d.Log.Trace("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073")
	_fffa, _adf = _aaf.alg6(d, pass)
	if _adf != nil {
		return nil, 0, _adf
	}
	if _fffa != nil {
		_d.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _fffa, d.P, nil
	}
	return nil, 0, nil
}

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4(id0 string, length int) StdHandler { return stdHandlerR4{ID0: id0, Length: length} }
func (_fec stdHandlerR6) alg9(_cfa *StdEncryptDict, _dbe []byte, _aef []byte) error {
	if _aad := _bcf("\u0061\u006c\u0067\u0039", "\u004b\u0065\u0079", 32, _dbe); _aad != nil {
		return _aad
	}
	if _bff := _bcf("\u0061\u006c\u0067\u0039", "\u0055", 48, _cfa.U); _bff != nil {
		return _bff
	}
	var _bdc [16]byte
	if _, _eccd := _a.ReadFull(_fa.Reader, _bdc[:]); _eccd != nil {
		return _eccd
	}
	_aga := _bdc[0:8]
	_afb := _bdc[8:16]
	_dbf := _cfa.U[:48]
	_daf := make([]byte, len(_aef)+len(_aga)+len(_dbf))
	_deb := copy(_daf, _aef)
	_deb += copy(_daf[_deb:], _aga)
	_deb += copy(_daf[_deb:], _dbf)
	_bgg, _agfd := _fec.alg2b(_cfa.R, _daf, _aef, _dbf)
	if _agfd != nil {
		return _agfd
	}
	O := make([]byte, len(_bgg)+len(_aga)+len(_afb))
	_deb = copy(O, _bgg[:32])
	_deb += copy(O[_deb:], _aga)
	_deb += copy(O[_deb:], _afb)
	_cfa.O = O
	_deb = len(_aef)
	_deb += copy(_daf[_deb:], _afb)
	_bgg, _agfd = _fec.alg2b(_cfa.R, _daf, _aef, _dbf)
	if _agfd != nil {
		return _agfd
	}
	_ggd, _agfd := _beb(_bgg[:32])
	if _agfd != nil {
		return _agfd
	}
	_afed := make([]byte, _fd.BlockSize)
	_fge := _e.NewCBCEncrypter(_ggd, _afed)
	OE := make([]byte, 32)
	_fge.CryptBlocks(OE, _dbe[:32])
	_cfa.OE = OE
	return nil
}

const _de = "\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377" + "\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014" + "\251\xfe\x64\x53\x69\172"

type ecb struct {
	_bca _e.Block
	_ae  int
}
type ecbDecrypter ecb

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32

func _beb(_bbbg []byte) (_e.Block, error) {
	_fe, _fee := _fd.NewCipher(_bbbg)
	if _fee != nil {
		_d.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076", _fee)
		return nil, _fee
	}
	return _fe, nil
}
func (_egc stdHandlerR6) alg10(_gad *StdEncryptDict, _bfd []byte) error {
	if _feg := _bcf("\u0061\u006c\u00671\u0030", "\u004b\u0065\u0079", 32, _bfd); _feg != nil {
		return _feg
	}
	_ee := uint64(uint32(_gad.P)) | (_ge.MaxUint32 << 32)
	Perms := make([]byte, 16)
	_gb.LittleEndian.PutUint64(Perms[:8], _ee)
	if _gad.EncryptMetadata {
		Perms[8] = 'T'
	} else {
		Perms[8] = 'F'
	}
	copy(Perms[9:12], "\u0061\u0064\u0062")
	if _, _cdb := _a.ReadFull(_fa.Reader, Perms[12:16]); _cdb != nil {
		return _cdb
	}
	_bcfb, _bbc := _beb(_bfd[:32])
	if _bbc != nil {
		return _bbc
	}
	_acg := _bd(_bcfb)
	_acg.CryptBlocks(Perms, Perms)
	_gad.Perms = Perms[:16]
	return nil
}
func (_dce stdHandlerR6) alg11(_cce *StdEncryptDict, _aadf []byte) ([]byte, error) {
	if _bee := _bcf("\u0061\u006c\u00671\u0031", "\u0055", 48, _cce.U); _bee != nil {
		return nil, _bee
	}
	_gbdf := make([]byte, len(_aadf)+8)
	_dgg := copy(_gbdf, _aadf)
	_dgg += copy(_gbdf[_dgg:], _cce.U[32:40])
	_egcf, _ffb := _dce.alg2b(_cce.R, _gbdf, _aadf, nil)
	if _ffb != nil {
		return nil, _ffb
	}
	_egcf = _egcf[:32]
	if !_cc.Equal(_egcf, _cce.U[:32]) {
		return nil, nil
	}
	return _egcf, nil
}
func (_fdcc stdHandlerR6) alg2b(R int, _fada, _fdf, _gce []byte) ([]byte, error) {
	if R == 5 {
		return _fdc(_fada)
	}
	return _dfc(_fada, _fdf, _gce)
}

// StdHandler is an interface for standard security handlers.
type StdHandler interface {

	// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
	// It assumes that R, P and EncryptMetadata are already set.
	GenerateParams(_ca *StdEncryptDict, _gf, _ab []byte) ([]byte, error)

	// Authenticate uses encryption dictionary parameters and the password to calculate
	// the document encryption key. It also returns permissions that should be granted to a user.
	// In case of failed authentication, it returns empty key and zero permissions with no error.
	Authenticate(_dge *StdEncryptDict, _gbf []byte) ([]byte, Permissions, error)
}

// Authenticate implements StdHandler interface.
func (_ccea stdHandlerR6) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	return _ccea.alg2a(d, pass)
}

type stdHandlerR4 struct {
	Length int
	ID0    string
}

func (_ggb stdHandlerR6) alg8(_cgb *StdEncryptDict, _fda []byte, _fceb []byte) error {
	if _gfe := _bcf("\u0061\u006c\u0067\u0038", "\u004b\u0065\u0079", 32, _fda); _gfe != nil {
		return _gfe
	}
	var _eaad [16]byte
	if _, _cfc := _a.ReadFull(_fa.Reader, _eaad[:]); _cfc != nil {
		return _cfc
	}
	_egb := _eaad[0:8]
	_cfg := _eaad[8:16]
	_aac := make([]byte, len(_fceb)+len(_egb))
	_bbaa := copy(_aac, _fceb)
	copy(_aac[_bbaa:], _egb)
	_cac, _aae := _ggb.alg2b(_cgb.R, _aac, _fceb, nil)
	if _aae != nil {
		return _aae
	}
	U := make([]byte, len(_cac)+len(_egb)+len(_cfg))
	_bbaa = copy(U, _cac[:32])
	_bbaa += copy(U[_bbaa:], _egb)
	copy(U[_bbaa:], _cfg)
	_cgb.U = U
	_bbaa = len(_fceb)
	copy(_aac[_bbaa:], _cfg)
	_cac, _aae = _ggb.alg2b(_cgb.R, _aac, _fceb, nil)
	if _aae != nil {
		return _aae
	}
	_bbaf, _aae := _beb(_cac[:32])
	if _aae != nil {
		return _aae
	}
	_egbb := make([]byte, _fd.BlockSize)
	_add := _e.NewCBCEncrypter(_bbaf, _egbb)
	UE := make([]byte, 32)
	_add.CryptBlocks(UE, _fda[:32])
	_cgb.UE = UE
	return nil
}
func (_bcg stdHandlerR4) alg3Key(R int, _fac []byte) []byte {
	_geb := _bc.New()
	_afe := _bcg.paddedPass(_fac)
	_geb.Write(_afe)
	if R >= 3 {
		for _bf := 0; _bf < 50; _bf++ {
			_bfa := _geb.Sum(nil)
			_geb = _bc.New()
			_geb.Write(_bfa)
		}
	}
	_bg := _geb.Sum(nil)
	if R == 2 {
		_bg = _bg[0:5]
	} else {
		_bg = _bg[0 : _bcg.Length/8]
	}
	return _bg
}
func (_cdf stdHandlerR4) alg5(_ad []byte, _bbe []byte) ([]byte, error) {
	_dgec := _bc.New()
	_dgec.Write([]byte(_de))
	_dgec.Write([]byte(_cdf.ID0))
	_fce := _dgec.Sum(nil)
	_d.Log.Trace("\u0061\u006c\u0067\u0035")
	_d.Log.Trace("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x", _ad)
	_d.Log.Trace("\u0049D\u003a\u0020\u0025\u0020\u0078", _cdf.ID0)
	if len(_fce) != 16 {
		return nil, _ec.New("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073")
	}
	_abf, _aab := _c.NewCipher(_ad)
	if _aab != nil {
		return nil, _ec.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_dgc := make([]byte, 16)
	_abf.XORKeyStream(_dgc, _fce)
	_ddb := make([]byte, len(_ad))
	for _bde := 0; _bde < 19; _bde++ {
		for _efea := 0; _efea < len(_ad); _efea++ {
			_ddb[_efea] = _ad[_efea] ^ byte(_bde+1)
		}
		_abf, _aab = _c.NewCipher(_ddb)
		if _aab != nil {
			return nil, _ec.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
		}
		_abf.XORKeyStream(_dgc, _dgc)
		_d.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078", _bde, _ddb)
		_d.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078", _bde, _dgc)
	}
	_ead := make([]byte, 32)
	for _caf := 0; _caf < 16; _caf++ {
		_ead[_caf] = _dgc[_caf]
	}
	_, _aab = _fa.Read(_ead[16:32])
	if _aab != nil {
		return nil, _ec.New("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r")
	}
	return _ead, nil
}

// Allowed checks if a set of permissions can be granted.
func (_afa Permissions) Allowed(p2 Permissions) bool { return _afa&p2 == p2 }

type ecbEncrypter ecb

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_ebg stdHandlerR6) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	_adc := make([]byte, 32)
	if _, _bddc := _a.ReadFull(_fa.Reader, _adc); _bddc != nil {
		return nil, _bddc
	}
	d.U = nil
	d.O = nil
	d.UE = nil
	d.OE = nil
	d.Perms = nil
	if len(upass) > 127 {
		upass = upass[:127]
	}
	if len(opass) > 127 {
		opass = opass[:127]
	}
	if _cgeb := _ebg.alg8(d, _adc, upass); _cgeb != nil {
		return nil, _cgeb
	}
	if _cff := _ebg.alg9(d, _adc, opass); _cff != nil {
		return nil, _cff
	}
	if d.R == 5 {
		return _adc, nil
	}
	if _ffcf := _ebg.alg10(d, _adc); _ffcf != nil {
		return nil, _ffcf
	}
	return _adc, nil
}
func _bcf(_fab, _aa string, _cca int, _fc []byte) error {
	if len(_fc) < _cca {
		return errInvalidField{Func: _fab, Field: _aa, Exp: _cca, Got: len(_fc)}
	}
	return nil
}
func (_bfb stdHandlerR4) alg4(_ged []byte, _ded []byte) ([]byte, error) {
	_efad, _gdc := _c.NewCipher(_ged)
	if _gdc != nil {
		return nil, _ec.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_cg := []byte(_de)
	_gc := make([]byte, len(_cg))
	_efad.XORKeyStream(_gc, _cg)
	return _gc, nil
}
func (_dga stdHandlerR6) alg2a(_adg *StdEncryptDict, _afc []byte) ([]byte, Permissions, error) {
	if _gbd := _bcf("\u0061\u006c\u00672\u0061", "\u004f", 48, _adg.O); _gbd != nil {
		return nil, 0, _gbd
	}
	if _fad := _bcf("\u0061\u006c\u00672\u0061", "\u0055", 48, _adg.U); _fad != nil {
		return nil, 0, _fad
	}
	if len(_afc) > 127 {
		_afc = _afc[:127]
	}
	_bfc, _cbe := _dga.alg12(_adg, _afc)
	if _cbe != nil {
		return nil, 0, _cbe
	}
	var (
		_abe []byte
		_adb []byte
		_ed  []byte
	)
	var _bgb Permissions
	if len(_bfc) != 0 {
		_bgb = PermOwner
		_ccd := make([]byte, len(_afc)+8+48)
		_fea := copy(_ccd, _afc)
		_fea += copy(_ccd[_fea:], _adg.O[40:48])
		copy(_ccd[_fea:], _adg.U[0:48])
		_abe = _ccd
		_adb = _adg.OE
		_ed = _adg.U[0:48]
	} else {
		_bfc, _cbe = _dga.alg11(_adg, _afc)
		if _cbe == nil && len(_bfc) == 0 {
			_bfc, _cbe = _dga.alg11(_adg, []byte(""))
		}
		if _cbe != nil {
			return nil, 0, _cbe
		} else if len(_bfc) == 0 {
			return nil, 0, nil
		}
		_bgb = _adg.P
		_dgef := make([]byte, len(_afc)+8)
		_bbgf := copy(_dgef, _afc)
		copy(_dgef[_bbgf:], _adg.U[40:48])
		_abe = _dgef
		_adb = _adg.UE
		_ed = nil
	}
	if _bba := _bcf("\u0061\u006c\u00672\u0061", "\u004b\u0065\u0079", 32, _adb); _bba != nil {
		return nil, 0, _bba
	}
	_adb = _adb[:32]
	_cf, _cbe := _dga.alg2b(_adg.R, _abe, _afc, _ed)
	if _cbe != nil {
		return nil, 0, _cbe
	}
	_abd, _cbe := _fd.NewCipher(_cf[:32])
	if _cbe != nil {
		return nil, 0, _cbe
	}
	_aged := make([]byte, _fd.BlockSize)
	_dba := _e.NewCBCDecrypter(_abd, _aged)
	_fcf := make([]byte, 32)
	_dba.CryptBlocks(_fcf, _adb)
	if _adg.R == 5 {
		return _fcf, _bgb, nil
	}
	_cbe = _dga.alg13(_adg, _fcf)
	if _cbe != nil {
		return nil, 0, _cbe
	}
	return _fcf, _bgb, nil
}
func (_dc errInvalidField) Error() string {
	return _ff.Sprintf("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064", _dc.Func, _dc.Field, _dc.Exp, _dc.Got)
}
func (_gadc stdHandlerR6) alg13(_baef *StdEncryptDict, _cbb []byte) error {
	if _dea := _bcf("\u0061\u006c\u00671\u0033", "\u004b\u0065\u0079", 32, _cbb); _dea != nil {
		return _dea
	}
	if _gfec := _bcf("\u0061\u006c\u00671\u0033", "\u0050\u0065\u0072m\u0073", 16, _baef.Perms); _gfec != nil {
		return _gfec
	}
	_bfeb := make([]byte, 16)
	copy(_bfeb, _baef.Perms[:16])
	_ccff, _gcc := _fd.NewCipher(_cbb[:32])
	if _gcc != nil {
		return _gcc
	}
	_fecd := _af(_ccff)
	_fecd.CryptBlocks(_bfeb, _bfeb)
	if !_cc.Equal(_bfeb[9:12], []byte("\u0061\u0064\u0062")) {
		return _ec.New("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_fcfe := Permissions(_gb.LittleEndian.Uint32(_bfeb[0:4]))
	if _fcfe != _baef.P {
		return _ec.New("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064")
	}
	var _abee bool
	if _bfeb[8] == 'T' {
		_abee = true
	} else if _bfeb[8] == 'F' {
		_abee = false
	} else {
		return _ec.New("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _abee != _baef.EncryptMetadata {
		return _ec.New("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	return nil
}
func _bd(_cd _e.Block) _e.BlockMode { return (*ecbEncrypter)(_efe(_cd)) }
func _fdc(_egd []byte) ([]byte, error) {
	_bbbe := _b.New()
	_bbbe.Write(_egd)
	return _bbbe.Sum(nil), nil
}
func _efe(_efa _e.Block) *ecb { return &ecb{_bca: _efa, _ae: _efa.BlockSize()} }
func (_dg *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_dg._ae != 0 {
		_d.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_d.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_dg._bca.Encrypt(dst, src[:_dg._ae])
		src = src[_dg._ae:]
		dst = dst[_dg._ae:]
	}
}

type errInvalidField struct {
	Func  string
	Field string
	Exp   int
	Got   int
}

func (_eb *ecbEncrypter) BlockSize() int { return _eb._ae }

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_fcb stdHandlerR4) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	O, _eba := _fcb.alg3(d.R, upass, opass)
	if _eba != nil {
		_d.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _eba)
		return nil, _eba
	}
	d.O = O
	_d.Log.Trace("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078", O)
	_bbb := _fcb.alg2(d, upass)
	U, _eba := _fcb.alg5(_bbb, upass)
	if _eba != nil {
		_d.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _eba)
		return nil, _eba
	}
	d.U = U
	_d.Log.Trace("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078", U)
	return _bbb, nil
}
func _dfc(_eaa, _da, _bce []byte) ([]byte, error) {
	var (
		_bcd, _bae, _dgcf _f.Hash
	)
	_bcd = _b.New()
	_agc := make([]byte, 64)
	_geeb := _bcd
	_geeb.Write(_eaa)
	K := _geeb.Sum(_agc[:0])
	_dbab := make([]byte, 64*(127+64+48))
	_ffc := func(_gfc int) ([]byte, error) {
		_baa := len(_da) + len(K) + len(_bce)
		_bgf := _dbab[:_baa]
		_bcee := copy(_bgf, _da)
		_bcee += copy(_bgf[_bcee:], K[:])
		_bcee += copy(_bgf[_bcee:], _bce)
		if _bcee != _baa {
			_d.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e")
			return nil, _ec.New("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065")
		}
		K1 := _dbab[:_baa*64]
		_baf(K1, _baa)
		_cec, _agb := _beb(K[0:16])
		if _agb != nil {
			return nil, _agb
		}
		_acd := _e.NewCBCEncrypter(_cec, K[16:32])
		_acd.CryptBlocks(K1, K1)
		E := K1
		_edd := 0
		for _ebaf := 0; _ebaf < 16; _ebaf++ {
			_edd += int(E[_ebaf] % 3)
		}
		var _aaa _f.Hash
		switch _edd % 3 {
		case 0:
			_aaa = _bcd
		case 1:
			if _bae == nil {
				_bae = _ef.New384()
			}
			_aaa = _bae
		case 2:
			if _dgcf == nil {
				_dgcf = _ef.New()
			}
			_aaa = _dgcf
		}
		_aaa.Reset()
		_aaa.Write(E)
		K = _aaa.Sum(_agc[:0])
		return E, nil
	}
	for _ecc := 0; ; {
		E, _facc := _ffc(_ecc)
		if _facc != nil {
			return nil, _facc
		}
		_dag := E[len(E)-1]
		_ecc++
		if _ecc >= 64 && _dag <= uint8(_ecc-32) {
			break
		}
	}
	return K[:32], nil
}
func (_gde stdHandlerR4) alg3(R int, _dd, _ga []byte) ([]byte, error) {
	var _cag []byte
	if len(_ga) > 0 {
		_cag = _gde.alg3Key(R, _ga)
	} else {
		_cag = _gde.alg3Key(R, _dd)
	}
	_ac, _ce := _c.NewCipher(_cag)
	if _ce != nil {
		return nil, _ec.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_cbg := _gde.paddedPass(_dd)
	_dee := make([]byte, len(_cbg))
	_ac.XORKeyStream(_dee, _cbg)
	if R >= 3 {
		_df := make([]byte, len(_cag))
		for _ega := 0; _ega < 19; _ega++ {
			for _cdc := 0; _cdc < len(_cag); _cdc++ {
				_df[_cdc] = _cag[_cdc] ^ byte(_ega+1)
			}
			_db, _ag := _c.NewCipher(_df)
			if _ag != nil {
				return nil, _ec.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
			}
			_db.XORKeyStream(_dee, _dee)
		}
	}
	return _dee, nil
}
func (_ege stdHandlerR4) alg2(_gd *StdEncryptDict, _bb []byte) []byte {
	_d.Log.Trace("\u0061\u006c\u0067\u0032")
	_gdd := _ege.paddedPass(_bb)
	_bbg := _bc.New()
	_bbg.Write(_gdd)
	_bbg.Write(_gd.O)
	var _cb [4]byte
	_gb.LittleEndian.PutUint32(_cb[:], uint32(_gd.P))
	_bbg.Write(_cb[:])
	_d.Log.Trace("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x", _cb)
	_bbg.Write([]byte(_ege.ID0))
	_d.Log.Trace("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076", _gd.R, _gd.EncryptMetadata)
	if (_gd.R >= 4) && !_gd.EncryptMetadata {
		_bbg.Write([]byte{0xff, 0xff, 0xff, 0xff})
	}
	_dgf := _bbg.Sum(nil)
	if _gd.R >= 3 {
		_bbg = _bc.New()
		for _dcc := 0; _dcc < 50; _dcc++ {
			_bbg.Reset()
			_bbg.Write(_dgf[0 : _ege.Length/8])
			_dgf = _bbg.Sum(nil)
		}
	}
	if _gd.R >= 3 {
		return _dgf[0 : _ege.Length/8]
	}
	return _dgf[0:5]
}
func (_cge stdHandlerR4) alg7(_gg *StdEncryptDict, _fff []byte) ([]byte, error) {
	_ecd := _cge.alg3Key(_gg.R, _fff)
	_ba := make([]byte, len(_gg.O))
	if _gg.R == 2 {
		_cea, _bfe := _c.NewCipher(_ecd)
		if _bfe != nil {
			return nil, _ec.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
		}
		_cea.XORKeyStream(_ba, _gg.O)
	} else if _gg.R >= 3 {
		_faf := append([]byte{}, _gg.O...)
		for _gdeb := 0; _gdeb < 20; _gdeb++ {
			_aba := append([]byte{}, _ecd...)
			for _bbf := 0; _bbf < len(_ecd); _bbf++ {
				_aba[_bbf] ^= byte(19 - _gdeb)
			}
			_ccf, _cde := _c.NewCipher(_aba)
			if _cde != nil {
				return nil, _ec.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
			}
			_ccf.XORKeyStream(_ba, _faf)
			_faf = append([]byte{}, _ba...)
		}
	} else {
		return nil, _ec.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	_fcg, _age := _cge.alg6(_gg, _ba)
	if _age != nil {
		return nil, nil
	}
	return _fcg, nil
}

var _ StdHandler = stdHandlerR4{}

func (_be stdHandlerR4) alg6(_dfb *StdEncryptDict, _agf []byte) ([]byte, error) {
	var (
		_efc []byte
		_gee error
	)
	_bdd := _be.alg2(_dfb, _agf)
	if _dfb.R == 2 {
		_efc, _gee = _be.alg4(_bdd, _agf)
	} else if _dfb.R >= 3 {
		_efc, _gee = _be.alg5(_bdd, _agf)
	} else {
		return nil, _ec.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	if _gee != nil {
		return nil, _gee
	}
	_d.Log.Trace("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f", string(_efc), string(_dfb.U))
	_cdfg := _efc
	_gda := _dfb.U
	if _dfb.R >= 3 {
		if len(_cdfg) > 16 {
			_cdfg = _cdfg[0:16]
		}
		if len(_gda) > 16 {
			_gda = _gda[0:16]
		}
	}
	if !_cc.Equal(_cdfg, _gda) {
		return nil, nil
	}
	return _bdd, nil
}

type stdHandlerR6 struct{}

func _baf(_cfe []byte, _afea int) {
	_eca := _afea
	for _eca < len(_cfe) {
		copy(_cfe[_eca:], _cfe[:_eca])
		_eca *= 2
	}
}
func (stdHandlerR4) paddedPass(_fgc []byte) []byte {
	_ccad := make([]byte, 32)
	_cab := copy(_ccad, _fgc)
	for ; _cab < 32; _cab++ {
		_ccad[_cab] = _de[_cab-len(_fgc)]
	}
	return _ccad
}
