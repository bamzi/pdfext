package sigutil

import (
	_d "bytes"
	_a "crypto"
	_ed "crypto/x509"
	_ab "encoding/asn1"
	_be "encoding/pem"
	_b "errors"
	_ac "fmt"
	_e "io"
	_bf "net/http"
	_g "time"

	_c "github.com/bamzi/pdfext/common"
	_ef "github.com/unidoc/timestamp"
	_ga "golang.org/x/crypto/ocsp"
)

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_acf *TimestampClient) GetEncodedToken(serverURL string, req *_ef.Request) ([]byte, error) {
	if serverURL == "" {
		return nil, _ac.Errorf("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c")
	}
	if req == nil {
		return nil, _ac.Errorf("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_egc, _fbg := req.Marshal()
	if _fbg != nil {
		return nil, _fbg
	}
	_cab, _fbg := _bf.NewRequest("\u0050\u004f\u0053\u0054", serverURL, _d.NewBuffer(_egc))
	if _fbg != nil {
		return nil, _fbg
	}
	_cab.Header.Set("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079")
	if _acf.BeforeHTTPRequest != nil {
		if _cc := _acf.BeforeHTTPRequest(_cab); _cc != nil {
			return nil, _cc
		}
	}
	_cf := _acf.HTTPClient
	if _cf == nil {
		_cf = _cabc()
	}
	_gff, _fbg := _cf.Do(_cab)
	if _fbg != nil {
		return nil, _fbg
	}
	defer _gff.Body.Close()
	_egd, _fbg := _e.ReadAll(_gff.Body)
	if _fbg != nil {
		return nil, _fbg
	}
	if _gff.StatusCode != _bf.StatusOK {
		return nil, _ac.Errorf("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064", _gff.StatusCode)
	}
	var _cee struct {
		Version _ab.RawValue
		Content _ab.RawValue
	}
	if _, _fbg = _ab.Unmarshal(_egd, &_cee); _fbg != nil {
		return nil, _fbg
	}
	return _cee.Content.FullBytes, nil
}

// NewCRLClient returns a new CRL client.
func NewCRLClient() *CRLClient { return &CRLClient{HTTPClient: _cabc()} }

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct {

	// HTTPClient is the HTTP client used to make certificate requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bf.Client
}

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient() *OCSPClient { return &OCSPClient{HTTPClient: _cabc(), Hash: _a.SHA1} }

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct {

	// HTTPClient is the HTTP client used to make CRL requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bf.Client
}

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct {

	// HTTPClient is the HTTP client used to make OCSP requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bf.Client

	// Hash is the hash function  used when constructing the OCSP
	// requests. If zero, SHA-1 will be used.
	Hash _a.Hash
}

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest(body _e.Reader, opts *_ef.RequestOptions) (*_ef.Request, error) {
	if opts == nil {
		opts = &_ef.RequestOptions{}
	}
	if opts.Hash == 0 {
		opts.Hash = _a.SHA256
	}
	if !opts.Hash.Available() {
		return nil, _ed.ErrUnsupportedAlgorithm
	}
	_cea := opts.Hash.New()
	if _, _gfc := _e.Copy(_cea, body); _gfc != nil {
		return nil, _gfc
	}
	return &_ef.Request{HashAlgorithm: opts.Hash, HashedMessage: _cea.Sum(nil), Certificates: opts.Certificates, TSAPolicyOID: opts.TSAPolicyOID, Nonce: opts.Nonce}, nil
}

// Get retrieves the certificate at the specified URL.
func (_bb *CertClient) Get(url string) (*_ed.Certificate, error) {
	if _bb.HTTPClient == nil {
		_bb.HTTPClient = _cabc()
	}
	_bbd, _ec := _bb.HTTPClient.Get(url)
	if _ec != nil {
		return nil, _ec
	}
	defer _bbd.Body.Close()
	_fb, _ec := _e.ReadAll(_bbd.Body)
	if _ec != nil {
		return nil, _ec
	}
	if _ea, _ := _be.Decode(_fb); _ea != nil {
		_fb = _ea.Bytes
	}
	_gad, _ec := _ed.ParseCertificate(_fb)
	if _ec != nil {
		return nil, _ec
	}
	return _gad, nil
}

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient() *TimestampClient { return &TimestampClient{HTTPClient: _cabc()} }

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_gc *CertClient) IsCA(cert *_ed.Certificate) bool {
	return cert.IsCA && _d.Equal(cert.RawIssuer, cert.RawSubject)
}
func _cabc() *_bf.Client { return &_bf.Client{Timeout: 5 * _g.Second} }

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_ged *CRLClient) MakeRequest(serverURL string, cert *_ed.Certificate) ([]byte, error) {
	if _ged.HTTPClient == nil {
		_ged.HTTPClient = _cabc()
	}
	if serverURL == "" {
		if len(cert.CRLDistributionPoints) == 0 {
			return nil, _b.New("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073")
		}
		serverURL = cert.CRLDistributionPoints[0]
	}
	_ba, _efc := _ged.HTTPClient.Get(serverURL)
	if _efc != nil {
		return nil, _efc
	}
	defer _ba.Body.Close()
	_gaf, _efc := _e.ReadAll(_ba.Body)
	if _efc != nil {
		return nil, _efc
	}
	if _fa, _ := _be.Decode(_gaf); _fa != nil {
		_gaf = _fa.Bytes
	}
	return _gaf, nil
}

// NewCertClient returns a new certificate client.
func NewCertClient() *CertClient { return &CertClient{HTTPClient: _cabc()} }

// GetIssuer retrieves the issuer of the provided certificate.
func (_bed *CertClient) GetIssuer(cert *_ed.Certificate) (*_ed.Certificate, error) {
	for _, _gab := range cert.IssuingCertificateURL {
		_gf, _ge := _bed.Get(_gab)
		if _ge != nil {
			_c.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076", cert.Subject.CommonName, _ge)
			continue
		}
		return _gf, nil
	}
	return nil, _ac.Errorf("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064")
}

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_gca *OCSPClient) MakeRequest(serverURL string, cert, issuer *_ed.Certificate) (*_ga.Response, []byte, error) {
	if _gca.HTTPClient == nil {
		_gca.HTTPClient = _cabc()
	}
	if serverURL == "" {
		if len(cert.OCSPServer) == 0 {
			return nil, nil, _b.New("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073")
		}
		serverURL = cert.OCSPServer[0]
	}
	_edb, _ff := _ga.CreateRequest(cert, issuer, &_ga.RequestOptions{Hash: _gca.Hash})
	if _ff != nil {
		return nil, nil, _ff
	}
	_fbf, _ff := _gca.HTTPClient.Post(serverURL, "\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074", _d.NewReader(_edb))
	if _ff != nil {
		return nil, nil, _ff
	}
	defer _fbf.Body.Close()
	_gee, _ff := _e.ReadAll(_fbf.Body)
	if _ff != nil {
		return nil, nil, _ff
	}
	if _ceb, _ := _be.Decode(_gee); _ceb != nil {
		_gee = _ceb.Bytes
	}
	_edg, _ff := _ga.ParseResponseForCert(_gee, cert, issuer)
	if _ff != nil {
		return nil, nil, _ff
	}
	return _edg, _gee, nil
}

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct {

	// HTTPClient is the HTTP client used to make timestamp requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bf.Client

	// Callbacks.
	BeforeHTTPRequest func(_ca *_bf.Request) error
}
