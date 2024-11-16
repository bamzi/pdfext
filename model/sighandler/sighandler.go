// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler

import (
	_bb "bytes"
	_ce "crypto"
	_gc "crypto/rand"
	_de "crypto/rsa"
	_fb "crypto/x509"
	_g "crypto/x509/pkix"
	_ea "encoding/asn1"
	_b "encoding/hex"
	_c "errors"
	_cg "fmt"
	_a "hash"
	_d "math/big"
	_eb "strings"
	_e "time"

	_bd "github.com/bamzi/pdfext/common"
	_dd "github.com/bamzi/pdfext/core"
	_gg "github.com/bamzi/pdfext/model"
	_gf "github.com/bamzi/pdfext/model/mdp"
	_fg "github.com/bamzi/pdfext/model/sigutil"
	_ed "github.com/unidoc/pkcs7"
	_cef "github.com/unidoc/timestamp"
)

// RevocationInfoArchival is OIDAttributeAdobeRevocation attribute.
type RevocationInfoArchival struct {
	Crl          []_ea.RawValue `asn1:"explicit,tag:0,optional"`
	Ocsp         []_ea.RawValue `asn1:"explicit,tag:1,optional"`
	OtherRevInfo []_ea.RawValue `asn1:"explicit,tag:2,optional"`
}

// Validate validates PdfSignature.
func (_dc *etsiPAdES) Validate(sig *_gg.PdfSignature, digest _gg.Hasher) (_gg.SignatureValidationResult, error) {
	_cda := sig.Contents.Bytes()
	_fce, _geg := _ed.Parse(_cda)
	if _geg != nil {
		return _gg.SignatureValidationResult{}, _geg
	}
	_cdeb, _eab := digest.(*_bb.Buffer)
	if !_eab {
		return _gg.SignatureValidationResult{}, _cg.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_fce.Content = _cdeb.Bytes()
	if _geg = _fce.Verify(); _geg != nil {
		return _gg.SignatureValidationResult{}, _geg
	}
	_eee := false
	_dee := false
	var _eea _e.Time
	for _, _eae := range _fce.Signers {
		_def := _eae.EncryptedDigest
		var _bdeb RevocationInfoArchival
		_geg = _fce.UnmarshalSignedAttribute(_ed.OIDAttributeAdobeRevocation, &_bdeb)
		if _geg == nil {
			if len(_bdeb.Crl) > 0 {
				_dee = true
			}
			if len(_bdeb.Ocsp) > 0 {
				_eee = true
			}
		}
		for _, _fcg := range _eae.UnauthenticatedAttributes {
			if _fcg.Type.Equal(_ed.OIDAttributeTimeStampToken) {
				_bafga, _bacd := _cef.Parse(_fcg.Value.Bytes)
				if _bacd != nil {
					return _gg.SignatureValidationResult{}, _bacd
				}
				_eea = _bafga.Time
				_gceec := _bafga.HashAlgorithm.New()
				_gceec.Write(_def)
				if !_bb.Equal(_gceec.Sum(nil), _bafga.HashedMessage) {
					return _gg.SignatureValidationResult{}, _cg.Errorf("\u0048\u0061\u0073\u0068\u0020i\u006e\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u0069s\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u006b\u0063\u0073\u0037")
				}
				break
			}
		}
	}
	_bbe := _gg.SignatureValidationResult{IsSigned: true, IsVerified: true, IsCrlFound: _dee, IsOcspFound: _eee, GeneralizedTime: _eea}
	return _bbe, nil
}

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct {
	_cf        _gg.SignatureHandler
	Permission _gf.DocMDPPermission
}

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct {

	// SignatureSize is the estimated size of the signature contents in bytes.
	// If not provided, a default signature size of 4192 is used.
	// The signing process will report the model.ErrSignNotEnoughSpace error
	// if the estimated signature size is smaller than the actual size of the
	// signature.
	SignatureSize int

	// Client is the timestamp client used to make the signature request.
	// If no client is provided, a default one is used.
	Client *_fg.TimestampClient
}

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func(_dcf *_gg.PdfSignature, _gca _gg.Hasher) ([]byte, error)

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts(certificate *_fb.Certificate, signFunc SignFunc, opts *AdobeX509RSASHA1Opts) (_gg.SignatureHandler, error) {
	if opts == nil {
		opts = &AdobeX509RSASHA1Opts{}
	}
	return &adobeX509RSASHA1{_dgee: certificate, _aga: signFunc, _bf: opts.EstimateSize, _fad: opts.Algorithm}, nil
}

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level B signature handler.
func NewEtsiPAdESLevelB(privateKey *_de.PrivateKey, certificate *_fb.Certificate, caCert *_fb.Certificate) (_gg.SignatureHandler, error) {
	return &etsiPAdES{_fgc: certificate, _baa: privateKey, _eaa: caCert}, nil
}

// Sign sets the Contents fields.
func (_efc *adobePKCS7Detached) Sign(sig *_gg.PdfSignature, digest _gg.Hasher) error {
	if _efc._dcb {
		_eff := _efc._ege
		if _eff <= 0 {
			_eff = 8192
		}
		sig.Contents = _dd.MakeHexString(string(make([]byte, _eff)))
		return nil
	}
	_gdff, _cabb := digest.(*_bb.Buffer)
	if !_cabb {
		return _cg.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_efg, _deb := _ed.NewSignedData(_gdff.Bytes())
	if _deb != nil {
		return _deb
	}
	if _aaf := _efg.AddSigner(_efc._ac, _efc._aeec, _ed.SignerInfoConfig{}); _aaf != nil {
		return _aaf
	}
	_efg.Detach()
	_deba, _deb := _efg.Finish()
	if _deb != nil {
		return _deb
	}
	_cabbf := make([]byte, 8192)
	copy(_cabbf, _deba)
	sig.Contents = _dd.MakeHexString(string(_cabbf))
	return nil
}
func _gegf(_deg []byte, _ade int) (_acc []byte) {
	_dggb := len(_deg)
	if _dggb > _ade {
		_dggb = _ade
	}
	_acc = make([]byte, _ade)
	copy(_acc[len(_acc)-_dggb:], _deg)
	return
}
func (_ggb *etsiPAdES) addDss(_dac, _gad []*_fb.Certificate, _cegc *RevocationInfoArchival) (int, error) {
	_dafbe, _bca, _cbb := _ggb.buildCertChain(_dac, _gad)
	if _cbb != nil {
		return 0, _cbb
	}
	_edgc, _cbb := _ggb.getCerts(_dafbe)
	if _cbb != nil {
		return 0, _cbb
	}
	var _adb, _gade [][]byte
	if _ggb.OCSPClient != nil {
		_adb, _cbb = _ggb.getOCSPs(_dafbe, _bca)
		if _cbb != nil {
			return 0, _cbb
		}
	}
	if _ggb.CRLClient != nil {
		_gade, _cbb = _ggb.getCRLs(_dafbe)
		if _cbb != nil {
			return 0, _cbb
		}
	}
	if !_ggb._df {
		_, _cbb = _ggb._cbeg.AddCerts(_edgc)
		if _cbb != nil {
			return 0, _cbb
		}
		_, _cbb = _ggb._cbeg.AddOCSPs(_adb)
		if _cbb != nil {
			return 0, _cbb
		}
		_, _cbb = _ggb._cbeg.AddCRLs(_gade)
		if _cbb != nil {
			return 0, _cbb
		}
	}
	_afe := 0
	for _, _ge := range _gade {
		_afe += len(_ge)
		_cegc.Crl = append(_cegc.Crl, _ea.RawValue{FullBytes: _ge})
	}
	for _, _cdg := range _adb {
		_afe += len(_cdg)
		_cegc.Ocsp = append(_cegc.Ocsp, _ea.RawValue{FullBytes: _cdg})
	}
	return _afe, nil
}
func (_faf *adobeX509RSASHA1) getCertificate(_debb *_gg.PdfSignature) (*_fb.Certificate, error) {
	if _faf._dgee != nil {
		return _faf._dgee, nil
	}
	_fdd, _bdf := _debb.GetCerts()
	if _bdf != nil {
		return nil, _bdf
	}
	return _fdd[0], nil
}

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom(certificate *_fb.Certificate, signFunc SignFunc) (_gg.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dgee: certificate, _aga: signFunc}, nil
}
func (_afed *adobeX509RSASHA1) sign(_gdca *_gg.PdfSignature, _ccg _gg.Hasher, _fef bool) error {
	if !_fef {
		return _afed.Sign(_gdca, _ccg)
	}
	_ced, _dbff := _afed._dgee.PublicKey.(*_de.PublicKey)
	if !_dbff {
		return _cg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", _ced)
	}
	_fed, _afc := _ea.Marshal(make([]byte, _ced.Size()))
	if _afc != nil {
		return _afc
	}
	_gdca.Contents = _dd.MakeHexString(string(_fed))
	return nil
}

// NewDigest creates a new digest.
func (_efad *etsiPAdES) NewDigest(_ *_gg.PdfSignature) (_gg.Hasher, error) {
	return _bb.NewBuffer(nil), nil
}

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler(handler _gg.SignatureHandler, permission _gf.DocMDPPermission) (_gg.SignatureHandler, error) {
	return &DocMDPHandler{_cf: handler, Permission: permission}, nil
}
func (_gdba *etsiPAdES) buildCertChain(_agc, _edf []*_fb.Certificate) ([]*_fb.Certificate, map[string]*_fb.Certificate, error) {
	_cdc := map[string]*_fb.Certificate{}
	for _, _beg := range _agc {
		_cdc[_beg.Subject.CommonName] = _beg
	}
	_daa := _agc
	for _, _eg := range _edf {
		_ecg := _eg.Subject.CommonName
		if _, _cfe := _cdc[_ecg]; _cfe {
			continue
		}
		_cdc[_ecg] = _eg
		_daa = append(_daa, _eg)
	}
	if len(_daa) == 0 {
		return nil, nil, _gg.ErrSignNoCertificates
	}
	var _egd error
	for _gdc := _daa[0]; _gdc != nil && !_gdba.CertClient.IsCA(_gdc); {
		var _bac *_fb.Certificate
		_, _efa := _cdc[_gdc.Issuer.CommonName]
		if !_efa {
			if _bac, _egd = _gdba.CertClient.GetIssuer(_gdc); _egd != nil {
				_bd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _egd)
				break
			}
			_cdc[_gdc.Issuer.CommonName] = _bac
			_daa = append(_daa, _bac)
		} else {
			break
		}
		_gdc = _bac
	}
	return _daa, _cdc, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_gcg *adobePKCS7Detached) IsApplicable(sig *_gg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_gfc *DocMDPHandler) Validate(sig *_gg.PdfSignature, digest _gg.Hasher) (_gg.SignatureValidationResult, error) {
	return _gg.SignatureValidationResult{}, _c.New("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065")
}

// NewDigest creates a new digest.
func (_fd *DocMDPHandler) NewDigest(sig *_gg.PdfSignature) (_gg.Hasher, error) {
	return _fd._cf.NewDigest(sig)
}

// InitSignature initialises the PdfSignature.
func (_ddc *etsiPAdES) InitSignature(sig *_gg.PdfSignature) error {
	if !_ddc._aeg {
		if _ddc._fgc == nil {
			return _c.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _ddc._baa == nil {
			return _c.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_cec := *_ddc
	sig.Handler = &_cec
	sig.Filter = _dd.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _dd.MakeName("\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_fgfe, _bcf := _cec.NewDigest(sig)
	if _bcf != nil {
		return _bcf
	}
	_, _bcf = _fgfe.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	if _bcf != nil {
		return _bcf
	}
	_cec._df = true
	_bcf = _cec.Sign(sig, _fgfe)
	_cec._df = false
	return _bcf
}

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1(privateKey *_de.PrivateKey, certificate *_fb.Certificate) (_gg.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dgee: certificate, _bgc: privateKey}, nil
}

// NewEtsiPAdESLevelT creates a new Adobe.PPKLite ETSI.CAdES.detached Level T signature handler.
func NewEtsiPAdESLevelT(privateKey *_de.PrivateKey, certificate *_fb.Certificate, caCert *_fb.Certificate, certificateTimestampServerURL string) (_gg.SignatureHandler, error) {
	return &etsiPAdES{_fgc: certificate, _baa: privateKey, _eaa: caCert, _dab: certificateTimestampServerURL}, nil
}

type etsiPAdES struct {
	_baa *_de.PrivateKey
	_fgc *_fb.Certificate
	_aeg bool
	_df  bool
	_eaa *_fb.Certificate
	_dab string

	// CertClient is the client used to retrieve certificates.
	CertClient *_fg.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_fg.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_fg.CRLClient
	_bde      *_gg.PdfAppender
	_cbeg     *_gg.DSS
}

func (_cc *etsiPAdES) getCRLs(_abfd []*_fb.Certificate) ([][]byte, error) {
	_fea := make([][]byte, 0, len(_abfd))
	for _, _fbg := range _abfd {
		for _, _eefg := range _fbg.CRLDistributionPoints {
			if _cc.CertClient.IsCA(_fbg) {
				continue
			}
			_gfg, _fbcg := _cc.CRLClient.MakeRequest(_eefg, _fbg)
			if _fbcg != nil {
				_bd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _fbcg)
				continue
			}
			_fea = append(_fea, _gfg)
		}
	}
	return _fea, nil
}
func (_fdf *etsiPAdES) getOCSPs(_dbb []*_fb.Certificate, _bg map[string]*_fb.Certificate) ([][]byte, error) {
	_dbe := make([][]byte, 0, len(_dbb))
	for _, _dbd := range _dbb {
		for _, _gdb := range _dbd.OCSPServer {
			if _fdf.CertClient.IsCA(_dbd) {
				continue
			}
			_aag, _cab := _bg[_dbd.Issuer.CommonName]
			if !_cab {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _gb, _ef := _fdf.OCSPClient.MakeRequest(_gdb, _dbd, _aag)
			if _ef != nil {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _ef)
				continue
			}
			_dbe = append(_dbe, _gb)
		}
	}
	return _dbe, nil
}

// Sign adds a new reference to signature's references array.
func (_gdd *DocMDPHandler) Sign(sig *_gg.PdfSignature, digest _gg.Hasher) error {
	return _gdd._cf.Sign(sig, digest)
}

// InitSignature initialises the PdfSignature.
func (_bbg *adobeX509RSASHA1) InitSignature(sig *_gg.PdfSignature) error {
	if _bbg._dgee == nil {
		return _c.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
	}
	if _bbg._bgc == nil && _bbg._aga == nil {
		return _c.New("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_ebg := *_bbg
	sig.Handler = &_ebg
	sig.Filter = _dd.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _dd.MakeName("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031")
	sig.Cert = _dd.MakeString(string(_ebg._dgee.Raw))
	sig.Reference = nil
	_efag, _acb := _ebg.NewDigest(sig)
	if _acb != nil {
		return _acb
	}
	_efag.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _ebg.sign(sig, _efag, _bbg._bf)
}

type timestampInfo struct {
	Version        int
	Policy         _ea.RawValue
	MessageImprint struct {
		HashAlgorithm _g.AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber    _ea.RawValue
	GeneralizedTime _e.Time
}

func (_afb *etsiPAdES) getCerts(_dfa []*_fb.Certificate) ([][]byte, error) {
	_dba := make([][]byte, 0, len(_dfa))
	for _, _aee := range _dfa {
		_dba = append(_dba, _aee.Raw)
	}
	return _dba, nil
}

const _afa = _ce.SHA1

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached(privateKey *_de.PrivateKey, certificate *_fb.Certificate) (_gg.SignatureHandler, error) {
	return &adobePKCS7Detached{_ac: certificate, _aeec: privateKey}, nil
}

type adobeX509RSASHA1 struct {
	_bgc  *_de.PrivateKey
	_dgee *_fb.Certificate
	_aga  SignFunc
	_bf   bool
	_fad  _ce.Hash
}

func (_eeef *docTimeStamp) getCertificate(_gfgd *_gg.PdfSignature) (*_fb.Certificate, error) {
	_efaa, _bef := _gfgd.GetCerts()
	if _bef != nil {
		return nil, _bef
	}
	return _efaa[0], nil
}
func _gbec(_ff *_de.PublicKey, _bdcad []byte) _ce.Hash {
	_becd := _ff.Size()
	if _becd != len(_bdcad) {
		return 0
	}
	_agd := func(_bdcee *_d.Int, _adg *_de.PublicKey, _dbcf *_d.Int) *_d.Int {
		_cfed := _d.NewInt(int64(_adg.E))
		_bdcee.Exp(_dbcf, _cfed, _adg.N)
		return _bdcee
	}
	_ccc := new(_d.Int).SetBytes(_bdcad)
	_daab := _agd(new(_d.Int), _ff, _ccc)
	_egde := _gegf(_daab.Bytes(), _becd)
	if _egde[0] != 0 || _egde[1] != 1 {
		return 0
	}
	_bag := []struct {
		Hash   _ce.Hash
		Prefix []byte
	}{{Hash: _ce.SHA1, Prefix: []byte{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14}}, {Hash: _ce.SHA256, Prefix: []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}}, {Hash: _ce.SHA384, Prefix: []byte{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30}}, {Hash: _ce.SHA512, Prefix: []byte{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40}}, {Hash: _ce.RIPEMD160, Prefix: []byte{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14}}}
	for _, _ecd := range _bag {
		_aeeg := _ecd.Hash.Size()
		_cca := len(_ecd.Prefix) + _aeeg
		if _bb.Equal(_egde[_becd-_cca:_becd-_aeeg], _ecd.Prefix) {
			return _ecd.Hash
		}
	}
	return 0
}

// NewDigest creates a new digest.
func (_edd *adobeX509RSASHA1) NewDigest(sig *_gg.PdfSignature) (_gg.Hasher, error) {
	if _cdd, _fcf := _edd.getHashAlgorithm(sig); _cdd != 0 && _fcf == nil {
		return _cdd.New(), nil
	}
	return _afa.New(), nil
}

type docTimeStamp struct {
	_degd string
	_gdae _ce.Hash
	_effd int
	_dbba *_fg.TimestampClient
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ecc *docTimeStamp) IsApplicable(sig *_gg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031"
}

// Validate validates PdfSignature.
func (_cbba *adobePKCS7Detached) Validate(sig *_gg.PdfSignature, digest _gg.Hasher) (_gg.SignatureValidationResult, error) {
	_adc := sig.Contents.Bytes()
	_aab, _adaf := _ed.Parse(_adc)
	if _adaf != nil {
		return _gg.SignatureValidationResult{}, _adaf
	}
	_ggg, _eag := digest.(*_bb.Buffer)
	if !_eag {
		return _gg.SignatureValidationResult{}, _cg.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_aab.Content = _ggg.Bytes()
	if _adaf = _aab.Verify(); _adaf != nil {
		return _gg.SignatureValidationResult{}, _adaf
	}
	return _gg.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// NewEtsiPAdESLevelLT creates a new Adobe.PPKLite ETSI.CAdES.detached Level LT signature handler.
func NewEtsiPAdESLevelLT(privateKey *_de.PrivateKey, certificate *_fb.Certificate, caCert *_fb.Certificate, certificateTimestampServerURL string, appender *_gg.PdfAppender) (_gg.SignatureHandler, error) {
	_eed := appender.Reader.DSS
	if _eed == nil {
		_eed = _gg.NewDSS()
	}
	if _cag := _eed.GenerateHashMaps(); _cag != nil {
		return nil, _cag
	}
	return &etsiPAdES{_fgc: certificate, _baa: privateKey, _eaa: caCert, _dab: certificateTimestampServerURL, CertClient: _fg.NewCertClient(), OCSPClient: _fg.NewOCSPClient(), CRLClient: _fg.NewCRLClient(), _bde: appender, _cbeg: _eed}, nil
}
func (_fag *etsiPAdES) makeTimestampRequest(_cgf string, _eef []byte) (_ea.RawValue, error) {
	_bdc := _ce.SHA512.New()
	_bdc.Write(_eef)
	_aef := _bdc.Sum(nil)
	_baf := _cef.Request{HashAlgorithm: _ce.SHA512, HashedMessage: _aef, Certificates: true, Extensions: nil, ExtraExtensions: nil}
	_fagg := _fg.NewTimestampClient()
	_dbf, _ag := _fagg.GetEncodedToken(_cgf, &_baf)
	if _ag != nil {
		return _ea.NullRawValue, _ag
	}
	return _ea.RawValue{FullBytes: _dbf}, nil
}

// InitSignature initialises the PdfSignature.
func (_defc *docTimeStamp) InitSignature(sig *_gg.PdfSignature) error {
	_abfdb := *_defc
	sig.Type = _dd.MakeName("\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070")
	sig.Handler = &_abfdb
	sig.Filter = _dd.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _dd.MakeName("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031")
	sig.Reference = nil
	if _defc._effd > 0 {
		sig.Contents = _dd.MakeHexString(string(make([]byte, _defc._effd)))
	} else {
		_debd, _bff := _defc.NewDigest(sig)
		if _bff != nil {
			return _bff
		}
		_debd.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
		if _bff = _abfdb.Sign(sig, _debd); _bff != nil {
			return _bff
		}
		_defc._effd = _abfdb._effd
	}
	return nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_daf *etsiPAdES) Sign(sig *_gg.PdfSignature, digest _gg.Hasher) error {
	_aac, _bafc := digest.(*_bb.Buffer)
	if !_bafc {
		return _cg.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_fbcf, _bec := _ed.NewSignedData(_aac.Bytes())
	if _bec != nil {
		return _bec
	}
	_fbcf.SetDigestAlgorithm(_ed.OIDDigestAlgorithmSHA256)
	_ddgf := _ed.SignerInfoConfig{}
	_dad := _ce.SHA256.New()
	_dad.Write(_daf._fgc.Raw)
	var _gfa struct {
		Seq struct{ Seq struct{ Value []byte } }
	}
	_gfa.Seq.Seq.Value = _dad.Sum(nil)
	var _dfaa []*_fb.Certificate
	var _gba []*_fb.Certificate
	if _daf._eaa != nil {
		_gba = []*_fb.Certificate{_daf._eaa}
	}
	_gff := RevocationInfoArchival{Crl: []_ea.RawValue{}, Ocsp: []_ea.RawValue{}, OtherRevInfo: []_ea.RawValue{}}
	_fc := 0
	if _daf._bde != nil && len(_daf._dab) > 0 {
		_afd, _bafg := _daf.makeTimestampRequest(_daf._dab, ([]byte)(""))
		if _bafg != nil {
			return _bafg
		}
		_cde, _bafg := _cef.Parse(_afd.FullBytes)
		if _bafg != nil {
			return _bafg
		}
		_dfaa = append(_dfaa, _cde.Certificates...)
	}
	if _daf._bde != nil {
		_abfa, _dafb := _daf.addDss([]*_fb.Certificate{_daf._fgc}, _gba, &_gff)
		if _dafb != nil {
			return _dafb
		}
		_fc += _abfa
		if len(_dfaa) > 0 {
			_abfa, _dafb = _daf.addDss(_dfaa, nil, &_gff)
			if _dafb != nil {
				return _dafb
			}
			_fc += _abfa
		}
		if !_daf._df {
			_daf._bde.SetDSS(_daf._cbeg)
		}
	}
	_ddgf.ExtraSignedAttributes = append(_ddgf.ExtraSignedAttributes, _ed.Attribute{Type: _ed.OIDAttributeSigningCertificateV2, Value: _gfa}, _ed.Attribute{Type: _ed.OIDAttributeAdobeRevocation, Value: _gff})
	if _abc := _fbcf.AddSignerChainPAdES(_daf._fgc, _daf._baa, _gba, _ddgf); _abc != nil {
		return _abc
	}
	_fbcf.Detach()
	if len(_daf._dab) > 0 {
		_dge := _fbcf.GetSignedData().SignerInfos[0].EncryptedDigest
		_ebc, _fda := _daf.makeTimestampRequest(_daf._dab, _dge)
		if _fda != nil {
			return _fda
		}
		_fda = _fbcf.AddTimestampTokenToSigner(0, _ebc.FullBytes)
		if _fda != nil {
			return _fda
		}
	}
	_afde, _bec := _fbcf.Finish()
	if _bec != nil {
		return _bec
	}
	_agb := make([]byte, len(_afde)+1024*2+_fc)
	copy(_agb, _afde)
	sig.Contents = _dd.MakeHexString(string(_agb))
	if !_daf._df && _daf._cbeg != nil {
		_dad = _ce.SHA1.New()
		_dad.Write(_agb)
		_ada := _eb.ToUpper(_b.EncodeToString(_dad.Sum(nil)))
		if _ada != "" {
			_daf._cbeg.VRI[_ada] = &_gg.VRI{Cert: _daf._cbeg.Certs, OCSP: _daf._cbeg.OCSPs, CRL: _daf._cbeg.CRLs}
		}
		_daf._bde.SetDSS(_daf._cbeg)
	}
	return nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_gab *docTimeStamp) Sign(sig *_gg.PdfSignature, digest _gg.Hasher) error {
	_dfbg, _aae := _fg.NewTimestampRequest(digest.(*_bb.Buffer), &_cef.RequestOptions{Hash: _gab._gdae, Certificates: true})
	if _aae != nil {
		return _aae
	}
	_dgec := _gab._dbba
	if _dgec == nil {
		_dgec = _fg.NewTimestampClient()
	}
	_gbd, _aae := _dgec.GetEncodedToken(_gab._degd, _dfbg)
	if _aae != nil {
		return _aae
	}
	_eged := len(_gbd)
	if _gab._effd > 0 && _eged > _gab._effd {
		return _gg.ErrSignNotEnoughSpace
	}
	if _eged > 0 {
		_gab._effd = _eged + 128
	}
	if sig.Contents != nil {
		_cdbd := sig.Contents.Bytes()
		copy(_cdbd, _gbd)
		_gbd = _cdbd
	}
	sig.Contents = _dd.MakeHexString(string(_gbd))
	return nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_fac *adobeX509RSASHA1) IsApplicable(sig *_gg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031"
}

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct {

	// EstimateSize specifies whether the size of the signature contents
	// should be estimated based on the modulus size of the public key
	// extracted from the signing certificate. If set to false, a mock Sign
	// call is made in order to estimate the size of the signature contents.
	EstimateSize bool

	// Algorithm specifies the algorithm used for performing signing.
	// If not specified, defaults to SHA1.
	Algorithm _ce.Hash
}

// InitSignature initialises the PdfSignature.
func (_cgc *adobePKCS7Detached) InitSignature(sig *_gg.PdfSignature) error {
	if !_cgc._dcb {
		if _cgc._ac == nil {
			return _c.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _cgc._aeec == nil {
			return _c.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_cgcc := *_cgc
	sig.Handler = &_cgcc
	sig.Filter = _dd.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _dd.MakeName("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_baef, _gdf := _cgcc.NewDigest(sig)
	if _gdf != nil {
		return _gdf
	}
	_baef.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _cgcc.Sign(sig, _baef)
}

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp(timestampServerURL string, hashAlgorithm _ce.Hash) (_gg.SignatureHandler, error) {
	return &docTimeStamp{_degd: timestampServerURL, _gdae: hashAlgorithm}, nil
}

// NewDigest creates a new digest.
func (_eeac *adobePKCS7Detached) NewDigest(sig *_gg.PdfSignature) (_gg.Hasher, error) {
	return _bb.NewBuffer(nil), nil
}
func (_fcb *adobeX509RSASHA1) getHashAlgorithm(_ebf *_gg.PdfSignature) (_ce.Hash, error) {
	_cdb, _cbec := _fcb.getCertificate(_ebf)
	if _cbec != nil {
		if _fcb._fad != 0 {
			return _fcb._fad, nil
		}
		return _afa, _cbec
	}
	if _ebf.Contents != nil {
		_ead := _ebf.Contents.Bytes()
		var _gaf []byte
		if _, _dbc := _ea.Unmarshal(_ead, &_gaf); _dbc == nil {
			_cfg := _gbec(_cdb.PublicKey.(*_de.PublicKey), _gaf)
			if _cfg > 0 {
				return _cfg, nil
			}
		}
	}
	if _fcb._fad != 0 {
		return _fcb._fad, nil
	}
	return _afa, nil
}

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts(timestampServerURL string, hashAlgorithm _ce.Hash, opts *DocTimeStampOpts) (_gg.SignatureHandler, error) {
	if opts == nil {
		opts = &DocTimeStampOpts{}
	}
	if opts.SignatureSize <= 0 {
		opts.SignatureSize = 4192
	}
	return &docTimeStamp{_degd: timestampServerURL, _gdae: hashAlgorithm, _effd: opts.SignatureSize, _dbba: opts.Client}, nil
}

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached(signatureLen int) (_gg.SignatureHandler, error) {
	return &adobePKCS7Detached{_dcb: true, _ege: signatureLen}, nil
}

// Validate validates PdfSignature.
func (_fec *adobeX509RSASHA1) Validate(sig *_gg.PdfSignature, digest _gg.Hasher) (_gg.SignatureValidationResult, error) {
	_cegb, _bad := _fec.getCertificate(sig)
	if _bad != nil {
		return _gg.SignatureValidationResult{}, _bad
	}
	_gcgc := sig.Contents.Bytes()
	var _bfg []byte
	if _, _efgd := _ea.Unmarshal(_gcgc, &_bfg); _efgd != nil {
		return _gg.SignatureValidationResult{}, _efgd
	}
	_bbc, _dff := digest.(_a.Hash)
	if !_dff {
		return _gg.SignatureValidationResult{}, _c.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dcfb, _ := _fec.getHashAlgorithm(sig)
	if _dcfb == 0 {
		_dcfb = _afa
	}
	if _cfb := _de.VerifyPKCS1v15(_cegb.PublicKey.(*_de.PublicKey), _dcfb, _bbc.Sum(nil), _bfg); _cfb != nil {
		return _gg.SignatureValidationResult{}, _cfb
	}
	return _gg.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_agcb *adobeX509RSASHA1) Sign(sig *_gg.PdfSignature, digest _gg.Hasher) error {
	var _geb []byte
	var _bbgf error
	if _agcb._aga != nil {
		_geb, _bbgf = _agcb._aga(sig, digest)
		if _bbgf != nil {
			return _bbgf
		}
	} else {
		_eda, _acf := digest.(_a.Hash)
		if !_acf {
			return _c.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_abg := _afa
		if _agcb._fad != 0 {
			_abg = _agcb._fad
		}
		_geb, _bbgf = _de.SignPKCS1v15(_gc.Reader, _agcb._bgc, _abg, _eda.Sum(nil))
		if _bbgf != nil {
			return _bbgf
		}
	}
	_geb, _bbgf = _ea.Marshal(_geb)
	if _bbgf != nil {
		return _bbgf
	}
	sig.Contents = _dd.MakeHexString(string(_geb))
	return nil
}

// NewDigest creates a new digest.
func (_fff *docTimeStamp) NewDigest(sig *_gg.PdfSignature) (_gg.Hasher, error) {
	return _bb.NewBuffer(nil), nil
}

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_ab *DocMDPHandler) ValidateWithOpts(sig *_gg.PdfSignature, digest _gg.Hasher, params _gg.SignatureHandlerDocMDPParams) (_gg.SignatureValidationResult, error) {
	_fe, _ad := _ab._cf.Validate(sig, digest)
	if _ad != nil {
		return _fe, _ad
	}
	_ca := params.Parser
	if _ca == nil {
		return _gg.SignatureValidationResult{}, _c.New("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	if !_fe.IsVerified {
		return _fe, nil
	}
	_bbb := params.DiffPolicy
	if _bbb == nil {
		_bbb = _gf.NewDefaultDiffPolicy()
	}
	for _cbe := 0; _cbe <= _ca.GetRevisionNumber(); _cbe++ {
		_ga, _ae := _ca.GetRevision(_cbe)
		if _ae != nil {
			return _gg.SignatureValidationResult{}, _ae
		}
		_fgf := _ga.GetTrailer()
		if _fgf == nil {
			return _gg.SignatureValidationResult{}, _c.New("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
		_fa, _ec := _dd.GetDict(_fgf.Get("\u0052\u006f\u006f\u0074"))
		if !_ec {
			return _gg.SignatureValidationResult{}, _c.New("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_gfe, _ec := _dd.GetDict(_fa.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
		if !_ec {
			continue
		}
		_bae, _ec := _dd.GetArray(_gfe.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_ec {
			continue
		}
		for _, _be := range _bae.Elements() {
			_bc, _db := _dd.GetDict(_be)
			if !_db {
				continue
			}
			_da, _db := _dd.GetDict(_bc.Get("\u0056"))
			if !_db {
				continue
			}
			if _dd.EqualObjects(_da.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"), sig.Contents) {
				_fe.DiffResults, _ae = _bbb.ReviewFile(_ga, _ca, &_gf.MDPParameters{DocMDPLevel: _ab.Permission})
				if _ae != nil {
					return _gg.SignatureValidationResult{}, _ae
				}
				_fe.IsVerified = _fe.DiffResults.IsPermitted()
				return _fe, nil
			}
		}
	}
	return _gg.SignatureValidationResult{}, _c.New("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073")
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_egb *etsiPAdES) IsApplicable(sig *_gg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

// InitSignature initialization of the DocMDP signature.
func (_dde *DocMDPHandler) InitSignature(sig *_gg.PdfSignature) error {
	_bee := _dde._cf.InitSignature(sig)
	if _bee != nil {
		return _bee
	}
	sig.Handler = _dde
	if sig.Reference == nil {
		sig.Reference = _dd.MakeArray()
	}
	sig.Reference.Append(_gg.NewPdfSignatureReferenceDocMDP(_gg.NewPdfTransformParamsDocMDP(_dde.Permission)).ToPdfObject())
	return nil
}
func (_bdce *adobePKCS7Detached) getCertificate(_fae *_gg.PdfSignature) (*_fb.Certificate, error) {
	if _bdce._ac != nil {
		return _bdce._ac, nil
	}
	_bed, _deed := _fae.GetCerts()
	if _deed != nil {
		return nil, _deed
	}
	return _bed[0], nil
}

type adobePKCS7Detached struct {
	_aeec *_de.PrivateKey
	_ac   *_fb.Certificate
	_dcb  bool
	_ege  int
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gd *DocMDPHandler) IsApplicable(sig *_gg.PdfSignature) bool {
	_ddg := false
	for _, _cb := range sig.Reference.Elements() {
		if _ba, _cd := _dd.GetDict(_cb); _cd {
			if _gda, _ee := _dd.GetNameVal(_ba.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _ee {
				if _gda != "\u0044\u006f\u0063\u004d\u0044\u0050" {
					return false
				}
				if _ebb, _dg := _dd.GetDict(_ba.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _dg {
					_, _dgg := _dd.GetNumberAsInt64(_ebb.Get("\u0050"))
					if _dgg != nil {
						return false
					}
					_ddg = true
					break
				}
			}
		}
	}
	return _ddg && _gd._cf.IsApplicable(sig)
}

// Validate validates PdfSignature.
func (_fgdf *docTimeStamp) Validate(sig *_gg.PdfSignature, digest _gg.Hasher) (_gg.SignatureValidationResult, error) {
	_faga := sig.Contents.Bytes()
	_gfb, _beda := _ed.Parse(_faga)
	if _beda != nil {
		return _gg.SignatureValidationResult{}, _beda
	}
	if _beda = _gfb.Verify(); _beda != nil {
		return _gg.SignatureValidationResult{}, _beda
	}
	var _dafc timestampInfo
	_, _beda = _ea.Unmarshal(_gfb.Content, &_dafc)
	if _beda != nil {
		return _gg.SignatureValidationResult{}, _beda
	}
	_gbab, _beda := _cgd(_dafc.MessageImprint.HashAlgorithm.Algorithm)
	if _beda != nil {
		return _gg.SignatureValidationResult{}, _beda
	}
	_bgf := _gbab.New()
	_gcaa, _abfg := digest.(*_bb.Buffer)
	if !_abfg {
		return _gg.SignatureValidationResult{}, _cg.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_bgf.Write(_gcaa.Bytes())
	_afbc := _bgf.Sum(nil)
	_befc := _gg.SignatureValidationResult{IsSigned: true, IsVerified: _bb.Equal(_afbc, _dafc.MessageImprint.HashedMessage), GeneralizedTime: _dafc.GeneralizedTime}
	return _befc, nil
}
func _cgd(_fgg _ea.ObjectIdentifier) (_ce.Hash, error) {
	switch {
	case _fgg.Equal(_ed.OIDDigestAlgorithmSHA1), _fgg.Equal(_ed.OIDDigestAlgorithmECDSASHA1), _fgg.Equal(_ed.OIDDigestAlgorithmDSA), _fgg.Equal(_ed.OIDDigestAlgorithmDSASHA1), _fgg.Equal(_ed.OIDEncryptionAlgorithmRSA):
		return _ce.SHA1, nil
	case _fgg.Equal(_ed.OIDDigestAlgorithmSHA256), _fgg.Equal(_ed.OIDDigestAlgorithmECDSASHA256):
		return _ce.SHA256, nil
	case _fgg.Equal(_ed.OIDDigestAlgorithmSHA384), _fgg.Equal(_ed.OIDDigestAlgorithmECDSASHA384):
		return _ce.SHA384, nil
	case _fgg.Equal(_ed.OIDDigestAlgorithmSHA512), _fgg.Equal(_ed.OIDDigestAlgorithmECDSASHA512):
		return _ce.SHA512, nil
	}
	return _ce.Hash(0), _ed.ErrUnsupportedAlgorithm
}
