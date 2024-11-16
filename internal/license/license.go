package license

import (
	_a "bytes"
	_ed "compress/gzip"
	_e "crypto"
	_acb "crypto/aes"
	_dga "crypto/cipher"
	_dg "crypto/hmac"
	_efc "crypto/rand"
	_ec "crypto/rsa"
	_eba "crypto/sha256"
	_bd "crypto/sha512"
	_ee "crypto/x509"
	_edg "encoding/base64"
	_eg "encoding/hex"
	_ac "encoding/json"
	_gd "encoding/pem"
	_cd "errors"
	_eb "fmt"
	_f "io"
	_gc "net"
	_gb "net/http"
	_ef "os"
	_cg "path/filepath"
	_g "sort"
	_b "strings"
	_d "sync"
	_db "time"

	_dbb "github.com/bamzi/pdfext/common"
)

const (
	_dc = "\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
	_ca = "\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
)

type meteredUsageCheckinResp struct {
	Instance      string `json:"inst"`
	Next          string `json:"next"`
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	RemainingDocs int    `json:"rd"`
	LimitDocs     bool   `json:"ld"`
}

var _fac []interface{}

func (_cgc *LicenseKey) TypeToString() string {
	if _cgc._be {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	if _cgc.Tier == LicenseTierUnlicensed {
		return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	}
	if _cgc.Tier == LicenseTierCommunity {
		return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065"
	}
	if _cgc.Tier == LicenseTierIndividual || _cgc.Tier == "\u0069\u006e\u0064i\u0065" {
		return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	}
	return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073"
}
func SetMeteredKey(apiKey string) error {
	if len(apiKey) == 0 {
		_dbb.Log.Error("\u004d\u0065\u0074\u0065\u0072e\u0064\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0041\u0050\u0049 \u004b\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079")
		_dbb.Log.Error("\u002d\u0020\u0047\u0072\u0061\u0062\u0020\u006f\u006e\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0072\u0065\u0065\u0020\u0054\u0069\u0065\u0072\u0020\u0061t\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063\u006c\u006fud\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _eb.Errorf("\u006de\u0074\u0065\u0072e\u0064\u0020\u006ci\u0063en\u0073\u0065\u0020\u0061\u0070\u0069\u0020k\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u003a\u0020\u0063\u0072\u0065\u0061\u0074\u0065 o\u006ee\u0020\u0061\u0074\u0020\u0068\u0074t\u0070\u0073\u003a\u002f\u002fc\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063.\u0069\u006f")
	}
	if _cbe != nil && (_cbe._be || _cbe.Tier != LicenseTierUnlicensed) {
		_dbb.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u0061\u006e\u006eo\u0074 \u0073\u0065\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0074\u0077\u0069c\u0065\u0020\u002d\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u006a\u0075\u0073\u0074\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069z\u0065\u0020\u006f\u006e\u0063\u0065")
		return _cd.New("\u006c\u0069\u0063en\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0061\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0073\u0065\u0074")
	}
	_aef := _gbaf()
	_aef._fdb = apiKey
	_bca, _ebd := _aef.getStatus()
	if _ebd != nil {
		return _ebd
	}
	if !_bca.Valid {
		return _cd.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_acd := &LicenseKey{_be: true, _fbc: apiKey, _agf: true}
	_cbe = _acd
	return nil
}
func (_ada *LicenseKey) IsLicensed() bool { return _ada.Tier != LicenseTierUnlicensed || _ada._be }

var _cbe = MakeUnlicensedKey()
var _gf = _db.Date(2019, 6, 6, 0, 0, 0, 0, _db.UTC)

type meteredStatusResp struct {
	Valid        bool  `json:"valid"`
	OrgCredits   int64 `json:"org_credits"`
	OrgUsed      int64 `json:"org_used"`
	OrgRemaining int64 `json:"org_remaining"`
}

func _eed(_df string, _dgb []byte) (string, error) {
	_bf, _ := _gd.Decode([]byte(_df))
	if _bf == nil {
		return "", _eb.Errorf("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	_bda, _dd := _ee.ParsePKCS1PrivateKey(_bf.Bytes)
	if _dd != nil {
		return "", _dd
	}
	_fb := _bd.New()
	_fb.Write(_dgb)
	_ae := _fb.Sum(nil)
	_acbd, _dd := _ec.SignPKCS1v15(_efc.Reader, _bda, _e.SHA512, _ae)
	if _dd != nil {
		return "", _dd
	}
	_cdc := _edg.StdEncoding.EncodeToString(_dgb)
	_cdc += "\u000a\u002b\u000a"
	_cdc += _edg.StdEncoding.EncodeToString(_acbd)
	return _cdc, nil
}
func (_fgb *meteredClient) getStatus() (meteredStatusResp, error) {
	var _bdag meteredStatusResp
	_cafc := _fgb._ddb + "\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073"
	var _fe meteredStatusForm
	_gef, _gfc := _ac.Marshal(_fe)
	if _gfc != nil {
		return _bdag, _gfc
	}
	_acc, _gfc := _dggg(_gef)
	if _gfc != nil {
		return _bdag, _gfc
	}
	_gda, _gfc := _gb.NewRequest("\u0050\u004f\u0053\u0054", _cafc, _acc)
	if _gfc != nil {
		return _bdag, _gfc
	}
	_gda.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_gda.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gda.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gda.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _fgb._fdb)
	_fbf, _gfc := _fgb._bdb.Do(_gda)
	if _gfc != nil {
		return _bdag, _gfc
	}
	defer _fbf.Body.Close()
	if _fbf.StatusCode != 200 {
		return _bdag, _eb.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _fbf.StatusCode)
	}
	_cfg, _gfc := _ggga(_fbf)
	if _gfc != nil {
		return _bdag, _gfc
	}
	_gfc = _ac.Unmarshal(_cfg, &_bdag)
	if _gfc != nil {
		return _bdag, _gfc
	}
	return _bdag, nil
}

type meteredClient struct {
	_ddb string
	_fdb string
	_bdb *_gb.Client
}

func _aaf(_eea string) (LicenseKey, error) {
	var _ga LicenseKey
	_eae, _caf := _bfa(_dc, _ca, _eea)
	if _caf != nil {
		return _ga, _caf
	}
	_bcc, _caf := _bdg(_fec, _eae)
	if _caf != nil {
		return _ga, _caf
	}
	_caf = _ac.Unmarshal(_bcc, &_ga)
	if _caf != nil {
		return _ga, _caf
	}
	_ga.CreatedAt = _db.Unix(_ga.CreatedAtInt, 0)
	if _ga.ExpiresAtInt > 0 {
		_cb := _db.Unix(_ga.ExpiresAtInt, 0)
		_ga.ExpiresAt = &_cb
	}
	return _ga, nil
}
func SetMeteredKeyUsageLogVerboseMode(val bool) { _cbe._edb = val }

var _ecd = _db.Date(2020, 1, 1, 0, 0, 0, 0, _db.UTC)

type reportState struct {
	Instance      string         `json:"inst"`
	Next          string         `json:"n"`
	Docs          int64          `json:"d"`
	NumErrors     int64          `json:"e"`
	LimitDocs     bool           `json:"ld"`
	RemainingDocs int64          `json:"rd"`
	LastReported  _db.Time       `json:"lr"`
	LastWritten   _db.Time       `json:"lw"`
	Usage         map[string]int `json:"u"`
	UsageLogs     []interface{}  `json:"ul,omitempty"`
}

var _agb map[string]int

func (_dgac defaultStateHolder) loadState(_acce string) (reportState, error) {
	_bb, _ecc := _ded()
	if _ecc != nil {
		return reportState{}, _ecc
	}
	_ecc = _ef.MkdirAll(_bb, 0777)
	if _ecc != nil {
		return reportState{}, _ecc
	}
	if len(_acce) < 20 {
		return reportState{}, _cd.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_faa := []byte(_acce)
	_aea := _bd.Sum512_256(_faa[:20])
	_fc := _eg.EncodeToString(_aea[:])
	_bdbc := _cg.Join(_bb, _fc)
	_da, _ecc := _ef.ReadFile(_bdbc)
	if _ecc != nil {
		if _ef.IsNotExist(_ecc) {
			return reportState{}, nil
		}
		_dbb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ecc)
		return reportState{}, _cd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	const _gfg = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_da, _ecc = _afdc([]byte(_gfg), _da)
	if _ecc != nil {
		return reportState{}, _ecc
	}
	var _ffg reportState
	_ecc = _ac.Unmarshal(_da, &_ffg)
	if _ecc != nil {
		_dbb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076", _ecc)
		return reportState{}, _cd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	return _ffg, nil
}

var _gbg stateLoader = defaultStateHolder{}

func (_agd *LicenseKey) getExpiryDateToCompare() _db.Time {
	if _agd.Trial {
		return _db.Now().UTC()
	}
	return _dbb.ReleasedAt
}
func (_abgd defaultStateHolder) updateState(_abd, _aabe, _geg string, _beg int, _cdb bool, _cc int, _dcb int, _faea _db.Time, _eaf map[string]int, _efcg ...interface{}) error {
	_cad, _de := _ded()
	if _de != nil {
		return _de
	}
	if len(_abd) < 20 {
		return _cd.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_ecad := []byte(_abd)
	_fgf := _bd.Sum512_256(_ecad[:20])
	_geb := _eg.EncodeToString(_fgf[:])
	_bdca := _cg.Join(_cad, _geb)
	var _fbbc reportState
	_fbbc.Docs = int64(_beg)
	_fbbc.NumErrors = int64(_dcb)
	_fbbc.LimitDocs = _cdb
	_fbbc.RemainingDocs = int64(_cc)
	_fbbc.LastWritten = _db.Now().UTC()
	_fbbc.LastReported = _faea
	_fbbc.Instance = _aabe
	_fbbc.Next = _geg
	_fbbc.Usage = _eaf
	_fbbc.UsageLogs = _efcg
	_aefa, _de := _ac.Marshal(_fbbc)
	if _de != nil {
		return _de
	}
	const _bcb = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_aefa, _de = _gfe([]byte(_bcb), _aefa)
	if _de != nil {
		return _de
	}
	_de = _ef.WriteFile(_bdca, _aefa, 0600)
	if _de != nil {
		return _de
	}
	return nil
}

var _ggg map[string]struct{}

const _dgab = "\u0055\u004e\u0049\u0050DF\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041\u0054\u0048"
const (
	LicenseTierUnlicensed = "\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	LicenseTierCommunity  = "\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y"
	LicenseTierIndividual = "\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	LicenseTierBusiness   = "\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073"
)

var _caff = _db.Date(2010, 1, 1, 0, 0, 0, 0, _db.UTC)

func _bee(_egbd, _ddg string) string {
	_gfd := []byte(_egbd)
	_egaf := _dg.New(_eba.New, _gfd)
	_egaf.Write([]byte(_ddg))
	return _edg.StdEncoding.EncodeToString(_egaf.Sum(nil))
}

type LicenseKey struct {
	LicenseId    string    `json:"license_id"`
	CustomerId   string    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Tier         string    `json:"tier"`
	CreatedAt    _db.Time  `json:"-"`
	CreatedAtInt int64     `json:"created_at"`
	ExpiresAt    *_db.Time `json:"-"`
	ExpiresAtInt int64     `json:"expires_at"`
	CreatedBy    string    `json:"created_by"`
	CreatorName  string    `json:"creator_name"`
	CreatorEmail string    `json:"creator_email"`
	UniPDF       bool      `json:"unipdf"`
	UniOffice    bool      `json:"unioffice"`
	UniHTML      bool      `json:"unihtml"`
	Trial        bool      `json:"trial"`
	_be          bool
	_fbc         string
	_agf         bool
	_edb         bool
}

func _gff() ([]string, []string, error) {
	_fbdc, _baag := _gc.Interfaces()
	if _baag != nil {
		return nil, nil, _baag
	}
	var _ega []string
	var _faca []string
	for _, _dfeb := range _fbdc {
		if _dfeb.Flags&_gc.FlagUp == 0 || _a.Equal(_dfeb.HardwareAddr, nil) {
			continue
		}
		_fda, _fgd := _dfeb.Addrs()
		if _fgd != nil {
			return nil, nil, _fgd
		}
		_aag := 0
		for _, _cbb := range _fda {
			var _dad _gc.IP
			switch _dcf := _cbb.(type) {
			case *_gc.IPNet:
				_dad = _dcf.IP
			case *_gc.IPAddr:
				_dad = _dcf.IP
			}
			if _dad.IsLoopback() {
				continue
			}
			if _dad.To4() == nil {
				continue
			}
			_faca = append(_faca, _dad.String())
			_aag++
		}
		_gcb := _dfeb.HardwareAddr.String()
		if _gcb != "" && _aag > 0 {
			_ega = append(_ega, _gcb)
		}
	}
	return _ega, _faca, nil
}
func (_gdf *LicenseKey) Validate() error {
	return nil
	if _gdf._be {
		return nil
	}
	if len(_gdf.LicenseId) < 10 {
		return _eb.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064")
	}
	if len(_gdf.CustomerId) < 10 {
		return _eb.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064")
	}
	if len(_gdf.CustomerName) < 1 {
		return _eb.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065")
	}
	if _caff.After(_gdf.CreatedAt) {
		return _eb.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gdf.ExpiresAt == nil {
		_gg := _gdf.CreatedAt.AddDate(1, 0, 0)
		if _ecd.After(_gg) {
			_gg = _ecd
		}
		_gdf.ExpiresAt = &_gg
	}
	if _gdf.CreatedAt.After(*_gdf.ExpiresAt) {
		return _eb.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074")
	}
	if _gdf.isExpired() {
		_ge := "\u0054\u0068\u0065\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064\u002e\u000a" + "\u0059o\u0075\u0020\u006d\u0061y\u0020n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0075\u0070d\u0061\u0074\u0065\u0020\u0074\u0068\u0065\u0020l\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020t\u006f\u0020\u0074\u0068\u0065\u0020\u006e\u0065\u0077\u0065s\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0072\u0020\u0079o\u0075\u0072\u0020\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074i\u006fn\u002e\u000a" + "\u0054o\u0020\u0066\u0069\u006ed y\u006f\u0075\u0072\u0020n\u0065\u0077\u0065\u0073\u0074\u0020\u006c\u0069\u0063\u0065n\u0073\u0065\u0020\u006b\u0065\u0079\u002c\u0020\u0067\u006f\u0020\u0074\u006f\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063l\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064oc\u002e\u0069\u006f \u0061\u006e\u0064\u0020\u0067o\u0020t\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0069\u0063e\u006e\u0073\u0065\u0020\u006d\u0065\u006e\u0075\u002e"
		return _eb.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0025\u0073", _ge)
	}
	if len(_gdf.CreatorName) < 1 {
		return _eb.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065")
	}
	if len(_gdf.CreatorEmail) < 1 {
		return _eb.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c")
	}
	if _gdf.CreatedAt.After(_gf) {
		if !_gdf.UniPDF {
			return _eb.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020k\u0065\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d \u0066\u006f\u0072\u0020\u0055\u006e\u0069\u0050\u0044\u0046")
		}
	}
	return nil
}
func _ffgd(_eeac *_gb.Response) (_f.ReadCloser, error) {
	var _fef error
	var _gbac _f.ReadCloser
	switch _b.ToLower(_eeac.Header.Get("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")) {
	case "\u0067\u007a\u0069\u0070":
		_gbac, _fef = _ed.NewReader(_eeac.Body)
		if _fef != nil {
			return _gbac, _fef
		}
		defer _gbac.Close()
	default:
		_gbac = _eeac.Body
	}
	return _gbac, nil
}
func _bfa(_fd string, _bfc string, _gba string) (string, error) {
	_aab := _b.Index(_gba, _fd)
	if _aab == -1 {
		return "", _eb.Errorf("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ab := _b.Index(_gba, _bfc)
	if _ab == -1 {
		return "", _eb.Errorf("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_acg := _aab + len(_fd) + 1
	return _gba[_acg : _ab-1], nil
}
func _bdg(_dfe string, _fbb string) ([]byte, error) {
	var (
		_bc int
		_fg string
	)
	for _, _fg = range []string{"\u000a\u002b\u000a", "\u000d\u000a\u002b\r\u000a", "\u0020\u002b\u0020"} {
		if _bc = _b.Index(_fbb, _fg); _bc != -1 {
			break
		}
	}
	if _bc == -1 {
		return nil, _eb.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072")
	}
	_ea := _fbb[:_bc]
	_fa := _bc + len(_fg)
	_eeda := _fbb[_fa:]
	if _ea == "" || _eeda == "" {
		return nil, _eb.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065")
	}
	_bg, _gde := _edg.StdEncoding.DecodeString(_ea)
	if _gde != nil {
		return nil, _eb.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c")
	}
	_eab, _gde := _edg.StdEncoding.DecodeString(_eeda)
	if _gde != nil {
		return nil, _eb.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_eda, _ := _gd.Decode([]byte(_dfe))
	if _eda == nil {
		return nil, _eb.Errorf("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_ad, _gde := _ee.ParsePKIXPublicKey(_eda.Bytes)
	if _gde != nil {
		return nil, _gde
	}
	_cf := _ad.(*_ec.PublicKey)
	if _cf == nil {
		return nil, _eb.Errorf("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_ce := _bd.New()
	_ce.Write(_bg)
	_aeg := _ce.Sum(nil)
	_gde = _ec.VerifyPKCS1v15(_cf, _e.SHA512, _aeg, _eab)
	if _gde != nil {
		return nil, _gde
	}
	return _bg, nil
}
func Track(docKey string, useKey string, docName string) error {
	return _agg(docKey, useKey, docName, !_cbe._agf)
}
func _afdc(_aff, _geed []byte) ([]byte, error) {
	_afb := make([]byte, _edg.URLEncoding.DecodedLen(len(_geed)))
	_eec, _aga := _edg.URLEncoding.Decode(_afb, _geed)
	if _aga != nil {
		return nil, _aga
	}
	_afb = _afb[:_eec]
	_accb, _aga := _acb.NewCipher(_aff)
	if _aga != nil {
		return nil, _aga
	}
	if len(_afb) < _acb.BlockSize {
		return nil, _cd.New("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074")
	}
	_dce := _afb[:_acb.BlockSize]
	_afb = _afb[_acb.BlockSize:]
	_ffga := _dga.NewCFBDecrypter(_accb, _dce)
	_ffga.XORKeyStream(_afb, _afb)
	return _afb, nil
}
func SetMeteredKeyPersistentCache(val bool) { _cbe._agf = val }
func _gfe(_adag, _dfdc []byte) ([]byte, error) {
	_egb, _afc := _acb.NewCipher(_adag)
	if _afc != nil {
		return nil, _afc
	}
	_dgge := make([]byte, _acb.BlockSize+len(_dfdc))
	_ebgc := _dgge[:_acb.BlockSize]
	if _, _egbf := _f.ReadFull(_efc.Reader, _ebgc); _egbf != nil {
		return nil, _egbf
	}
	_edaf := _dga.NewCFBEncrypter(_egb, _ebgc)
	_edaf.XORKeyStream(_dgge[_acb.BlockSize:], _dfdc)
	_ddbe := make([]byte, _edg.URLEncoding.EncodedLen(len(_dgge)))
	_edg.URLEncoding.Encode(_ddbe, _dgge)
	return _ddbe, nil
}

type meteredUsageCheckinForm struct {
	Instance          string         `json:"inst"`
	Next              string         `json:"next"`
	UsageNumber       int            `json:"usage_number"`
	NumFailed         int64          `json:"num_failed"`
	Hostname          string         `json:"hostname"`
	LocalIP           string         `json:"local_ip"`
	MacAddress        string         `json:"mac_address"`
	Package           string         `json:"package"`
	PackageVersion    string         `json:"package_version"`
	Usage             map[string]int `json:"u"`
	IsPersistentCache bool           `json:"is_persistent_cache"`
	Timestamp         int64          `json:"timestamp"`
	UsageLogs         []interface{}  `json:"ul,omitempty"`
}

func _ded() (string, error) {
	_fee := _b.TrimSpace(_ef.Getenv(_adc))
	if _fee == "" {
		_dbb.Log.Debug("\u0024\u0025\u0073\u0020e\u006e\u0076\u0069\u0072\u006f\u006e\u006d\u0065\u006e\u0074\u0020\u0076\u0061\u0072\u0069\u0061\u0062l\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020\u0057\u0069\u006c\u006c\u0020\u0075\u0073\u0065\u0020\u0068\u006f\u006d\u0065\u0020\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0074\u006f\u0020s\u0074\u006f\u0072\u0065\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020in\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u002e", _adc)
		_cgb := _bcbg()
		if len(_cgb) == 0 {
			return "", _eb.Errorf("r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0024\u0025\u0073\u0020\u0065\u006e\u0076\u0069\u0072\u006f\u006e\u006d\u0065\u006e\u0074\u0020\u0076\u0061r\u0069a\u0062\u006c\u0065\u0020o\u0072\u0020h\u006f\u006d\u0065\u0020\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _adc)
		}
		_fee = _cg.Join(_cgb, "\u002eu\u006e\u0069\u0064\u006f\u0063")
	}
	_cbcd := _ef.MkdirAll(_fee, 0777)
	if _cbcd != nil {
		return "", _cbcd
	}
	return _fee, nil
}
func _gce() (_gc.IP, error) {
	_fbfa, _dcd := _gc.Dial("\u0075\u0064\u0070", "\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030")
	if _dcd != nil {
		return nil, _dcd
	}
	defer _fbfa.Close()
	_gaa := _fbfa.LocalAddr().(*_gc.UDPAddr)
	return _gaa.IP, nil
}
func GetMeteredState() (MeteredStatus, error) {
	if _cbe == nil {
		return MeteredStatus{}, _cd.New("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	if !_cbe._be || len(_cbe._fbc) == 0 {
		return MeteredStatus{}, _cd.New("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_dbe, _bfaf := _gbg.loadState(_cbe._fbc)
	if _bfaf != nil {
		_dbb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfaf)
		return MeteredStatus{}, _bfaf
	}
	if _dbe.Docs > 0 {
		_adg := _agg("", "", "", true)
		if _adg != nil {
			return MeteredStatus{}, _adg
		}
	}
	_fgg.Lock()
	defer _fgg.Unlock()
	_faf := _gbaf()
	_faf._fdb = _cbe._fbc
	_eca, _bfaf := _faf.getStatus()
	if _bfaf != nil {
		return MeteredStatus{}, _bfaf
	}
	if !_eca.Valid {
		return MeteredStatus{}, _cd.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_efa := MeteredStatus{OK: true, Credits: _eca.OrgCredits, Used: _eca.OrgUsed}
	return _efa, nil
}
func GetLicenseKey() *LicenseKey {
	if _cbe == nil {
		return nil
	}
	_bff := *_cbe
	return &_bff
}
func _bcbg() string {
	_ffgg := _ef.Getenv("\u0048\u004f\u004d\u0045")
	if len(_ffgg) == 0 {
		_ffgg, _ = _ef.UserHomeDir()
	}
	return _ffgg
}

const _fec = "\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n"

func _agg(_ggd string, _abf string, _abgb string, _bgg bool) error {
	if _cbe == nil {
		return _cd.New("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079")
	}
	if !_cbe._be || len(_cbe._fbc) == 0 {
		return nil
	}
	if len(_ggd) == 0 && !_bgg {
		return _cd.New("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	_fgg.Lock()
	defer _fgg.Unlock()
	if _ggg == nil {
		_ggg = map[string]struct{}{}
	}
	if _agb == nil {
		_agb = map[string]int{}
	}
	_dfed := 0
	if len(_ggd) > 0 {
		_, _agc := _ggg[_ggd]
		if !_agc {
			_ggg[_ggd] = struct{}{}
			_dfed++
		}
		if _cbe._edb {
			_fac = append(_fac, map[string]interface{}{"\u0074\u0069\u006d\u0065": _db.Now().String(), "\u0066\u0075\u006e\u0063": _abf, "\u0072\u0065\u0066": _ggd[:8], "\u0066\u0069\u006c\u0065": _abgb, "\u0063\u006f\u0073\u0074": _dfed})
			if _agc && _dfed == 0 {
				_dbb.Log.Info("\u0025\u0073\u0020\u0052\u0065\u0066\u003a\u0020\u0025\u0073\u0020\u007c\u0020\u0025\u0073 \u007c \u004e\u006f\u0020\u0063\u0072\u0065\u0064\u0069\u0074\u0020\u0075\u0073\u0065\u0064", _db.Now().String(), _ggd[:8], _abf)
			}
		}
	}
	if _dfed == 0 && !_bgg {
		return nil
	}
	_agb[_abf]++
	_cag := _db.Now()
	_ace, _gdab := _gbg.loadState(_cbe._fbc)
	if _gdab != nil {
		_dbb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gdab)
		return _gdab
	}
	_ace.UsageLogs = append(_ace.UsageLogs, _fac...)
	if _ace.Usage == nil {
		_ace.Usage = map[string]int{}
	}
	for _cagc, _cfb := range _agb {
		if _cagc != "" {
			_ace.Usage[_cagc] += _cfb
		}
	}
	_agb = nil
	const _dfd = 24 * _db.Hour
	const _eee = 3 * 24 * _db.Hour
	if len(_ace.Instance) == 0 || _cag.Sub(_ace.LastReported) > _dfd || (_ace.LimitDocs && _ace.RemainingDocs <= _ace.Docs+int64(_dfed)) || _bgg {
		_ffd, _afd := _ef.Hostname()
		if _afd != nil {
			return _afd
		}
		_gag := _ace.Docs
		_dgg, _feb, _afd := _gff()
		if _afd != nil {
			_dbb.Log.Debug("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u006c\u006fc\u0061\u006c\u0020\u0061\u0064\u0064\u0072\u0065\u0073\u0073\u003a\u0020\u0025\u0073", _afd.Error())
			_dgg = append(_dgg, "\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce")
			_feb = append(_feb, "\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce")
		} else {
			_g.Strings(_feb)
			_g.Strings(_dgg)
			_aad, _cea := _gce()
			if _cea != nil {
				return _cea
			}
			_cge := false
			for _, _fba := range _feb {
				if _fba == _aad.String() {
					_cge = true
				}
			}
			if !_cge {
				_feb = append(_feb, _aad.String())
			}
		}
		_agbf := _gbaf()
		_agbf._fdb = _cbe._fbc
		_gag += int64(_dfed)
		_aeb := meteredUsageCheckinForm{Instance: _ace.Instance, Next: _ace.Next, UsageNumber: int(_gag), NumFailed: _ace.NumErrors, Hostname: _ffd, LocalIP: _b.Join(_feb, "\u002c\u0020"), MacAddress: _b.Join(_dgg, "\u002c\u0020"), Package: "\u0075\u006e\u0069\u0070\u0064\u0066", PackageVersion: _dbb.Version, Usage: _ace.Usage, IsPersistentCache: _cbe._agf, Timestamp: _cag.Unix()}
		if len(_dgg) == 0 {
			_aeb.MacAddress = "\u006e\u006f\u006e\u0065"
		}
		if _cbe._edb {
			_aeb.UsageLogs = _ace.UsageLogs
		}
		_ccc := int64(0)
		_fed := _ace.NumErrors
		_bcad := _cag
		_ceg := 0
		_gcd := _ace.LimitDocs
		_abgba, _afd := _agbf.checkinUsage(_aeb)
		if _afd != nil {
			if _cag.Sub(_ace.LastReported) > _eee {
				if !_abgba.Success {
					return _cd.New(_abgba.Message)
				}
				return _cd.New("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e")
			}
			_ccc = _gag
			_fed++
			_bcad = _ace.LastReported
		} else {
			_gcd = _abgba.LimitDocs
			_ceg = _abgba.RemainingDocs
			_fed = 0
		}
		if len(_abgba.Instance) == 0 {
			_abgba.Instance = _aeb.Instance
		}
		if len(_abgba.Next) == 0 {
			_abgba.Next = _aeb.Next
		}
		_afd = _gbg.updateState(_agbf._fdb, _abgba.Instance, _abgba.Next, int(_ccc), _gcd, _ceg, int(_fed), _bcad, nil)
		if _afd != nil {
			return _afd
		}
		if !_abgba.Success {
			return _eb.Errorf("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s", _abgba.Message)
		}
	} else {
		_gdab = _gbg.updateState(_cbe._fbc, _ace.Instance, _ace.Next, int(_ace.Docs)+_dfed, _ace.LimitDocs, int(_ace.RemainingDocs), int(_ace.NumErrors), _ace.LastReported, _ace.Usage, _ace.UsageLogs...)
		if _gdab != nil {
			return _gdab
		}
	}
	if _cbe._edb && len(_ggd) > 0 {
		_cfe := ""
		if _abgb != "" {
			_cfe = _eb.Sprintf("\u0046i\u006c\u0065\u0020\u0025\u0073\u0020|", _abgb)
		}
		_dbb.Log.Info("%\u0073\u0020\u007c\u0020\u0025\u0073\u0020\u0052\u0065\u0066\u003a\u0020\u0025\u0073\u0020\u007c\u0020\u0025s\u0020\u007c\u0020\u0025\u0064\u0020\u0063\u0072\u0065\u0064it\u0028\u0073\u0029 \u0075s\u0065\u0064", _cag.String(), _cfe, _ggd[:8], _abf, _dfed)
	}
	return nil
}
func (_eabf *LicenseKey) isExpired() bool {
	return _eabf.getExpiryDateToCompare().After(*_eabf.ExpiresAt)
}

const _adc = "\u0055N\u0049D\u004f\u0043\u005f\u004c\u0049C\u0045\u004eS\u0045\u005f\u0044\u0049\u0052"

func TrackUse(useKey string) {
	if _cbe == nil {
		return
	}
	if !_cbe._be || len(_cbe._fbc) == 0 {
		return
	}
	if len(useKey) == 0 {
		return
	}
	_fgg.Lock()
	defer _fgg.Unlock()
	if _agb == nil {
		_agb = map[string]int{}
	}
	_agb[useKey]++
}
func MakeUnlicensedKey() *LicenseKey {
	_ege := LicenseKey{}
	_ege.CustomerName = "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	_ege.Tier = LicenseTierUnlicensed
	_ege.CreatedAt = _db.Now().UTC()
	_ege.CreatedAtInt = _ege.CreatedAt.Unix()
	return &_ege
}

const _bbf = "U\u004eI\u0050\u0044\u0046\u005f\u0043\u0055\u0053\u0054O\u004d\u0045\u0052\u005fNA\u004d\u0045"

func (_abg *LicenseKey) ToString() string {
	if _abg._be {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	_af := _eb.Sprintf("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _abg.LicenseId)
	_af += _eb.Sprintf("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _abg.CustomerId)
	_af += _eb.Sprintf("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a", _abg.CustomerName)
	_af += _eb.Sprintf("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n", _abg.Tier)
	_af += _eb.Sprintf("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _dbb.UtcTimeFormat(_abg.CreatedAt))
	if _abg.ExpiresAt == nil {
		_af += "\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a"
	} else {
		_af += _eb.Sprintf("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _dbb.UtcTimeFormat(*_abg.ExpiresAt))
	}
	_af += _eb.Sprintf("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a", _abg.CreatorName, _abg.CreatorEmail)
	return _af
}

var _fgg = &_d.Mutex{}

func init() {
	_gdg := _ef.Getenv(_dgab)
	_fdf := _ef.Getenv(_bbf)
	if len(_gdg) == 0 || len(_fdf) == 0 {
		return
	}
	_fbfb, _cff := _ef.ReadFile(_gdg)
	if _cff != nil {
		_dbb.Log.Error("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076", _cff)
		return
	}
	_cff = SetLicenseKey(string(_fbfb), _fdf)
	if _cff != nil {
		_dbb.Log.Error("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _cff)
		return
	}
}
func _gbaf() *meteredClient {
	_fga := meteredClient{_ddb: "h\u0074\u0074\u0070\u0073\u003a\u002f/\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069d\u006f\u0063\u002ei\u006f/\u0061\u0070\u0069", _bdb: &_gb.Client{Timeout: 30 * _db.Second}}
	if _bdc := _ef.Getenv("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L"); _b.HasPrefix(_bdc, "\u0068\u0074\u0074\u0070") {
		_fga._ddb = _bdc
	}
	return &_fga
}
func SetLicenseKey(content string, customerName string) error {
	_acdb, _eaeg := _aaf(content)
	if _eaeg != nil {
		_dbb.Log.Error("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076", _eaeg)
		return _eaeg
	}
	if !_b.EqualFold(_acdb.CustomerName, customerName) {
		_dbb.Log.Error("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027", _acdb.CustomerName, customerName)
		return _eb.Errorf("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'", _acdb.CustomerName, customerName)
	}
	_eaeg = _acdb.Validate()
	if _eaeg != nil {
		_dbb.Log.Error("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _eaeg)
		return _eaeg
	}
	_cbe = &_acdb
	return nil
}
func _ggga(_edc *_gb.Response) ([]byte, error) {
	var _fab []byte
	_fca, _daf := _ffgd(_edc)
	if _daf != nil {
		return _fab, _daf
	}
	return _f.ReadAll(_fca)
}

type MeteredStatus struct {
	OK      bool
	Credits int64
	Used    int64
}
type meteredStatusForm struct{}
type stateLoader interface {
	loadState(_cbc string) (reportState, error)
	updateState(_ff, _ebg, _abeg string, _eac int, _cbd bool, _baaf int, _ddd int, _fbce _db.Time, _gfb map[string]int, _efab ...interface{}) error
}
type defaultStateHolder struct{}

func (_gee *meteredClient) checkinUsage(_fbd meteredUsageCheckinForm) (meteredUsageCheckinResp, error) {
	_fbd.Package = "\u0075\u006e\u0069\u0070\u0064\u0066"
	_fbd.PackageVersion = _dbb.Version
	var _add meteredUsageCheckinResp
	_fae := _gee._ddb + "\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e"
	_edgf, _gab := _ac.Marshal(_fbd)
	if _gab != nil {
		return _add, _gab
	}
	_agda, _gab := _dggg(_edgf)
	if _gab != nil {
		return _add, _gab
	}
	_cdg, _gab := _gb.NewRequest("\u0050\u004f\u0053\u0054", _fae, _agda)
	if _gab != nil {
		return _add, _gab
	}
	_cdg.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_cdg.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_cdg.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_cdg.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _gee._fdb)
	_baa, _gab := _gee._bdb.Do(_cdg)
	if _gab != nil {
		return _add, _gab
	}
	defer _baa.Body.Close()
	if _baa.StatusCode != 200 {
		_abe, _abb := _ggga(_baa)
		if _abb != nil {
			return _add, _abb
		}
		_abb = _ac.Unmarshal(_abe, &_add)
		if _abb != nil {
			return _add, _abb
		}
		return _add, _eb.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _baa.StatusCode)
	}
	_fbfe := _baa.Header.Get("\u0058\u002d\u0055\u0043\u002d\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065")
	_abgf := _bee(_fbd.MacAddress, string(_edgf))
	if _abgf != _fbfe {
		_dbb.Log.Error("I\u006e\u0076\u0061l\u0069\u0064\u0020\u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u002c\u0020\u0073\u0065t\u0020\u0074\u0068e\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0073\u0065\u0072\u0076e\u0072\u0020\u0074\u006f \u0068\u0074\u0074\u0070s\u003a\u002f\u002f\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069o\u002f\u0061\u0070\u0069")
		return _add, _cd.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0065\u0072\u0076\u0065\u0072 \u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065")
	}
	_gcg, _gab := _ggga(_baa)
	if _gab != nil {
		return _add, _gab
	}
	_gab = _ac.Unmarshal(_gcg, &_add)
	if _gab != nil {
		return _add, _gab
	}
	return _add, nil
}
func _dggg(_gea []byte) (_f.Reader, error) {
	_cfeb := new(_a.Buffer)
	_afg := _ed.NewWriter(_cfeb)
	_afg.Write(_gea)
	_fbae := _afg.Close()
	if _fbae != nil {
		return nil, _fbae
	}
	return _cfeb, nil
}
