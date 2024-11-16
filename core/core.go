// Package core defines and implements the primitive PDF object types in golang, and provides functionality
// for parsing those from a PDF file stream. This includes I/O handling, cross references, repairs, encryption,
// encoding and other core capabilities.
package core

import (
	_ga "bufio"
	_bg "bytes"
	_eg "compress/lzw"
	_e "compress/zlib"
	_ab "crypto/md5"
	_dg "crypto/rand"
	_ae "encoding/hex"
	_f "errors"
	_ee "fmt"
	_a "image"
	_cb "image/color"
	_beg "image/jpeg"
	_bb "io"
	_c "os"
	_fa "reflect"
	_ba "regexp"
	_be "sort"
	_d "strconv"
	_cc "strings"
	_g "sync"
	_ag "time"
	_gb "unicode"

	_eb "github.com/bamzi/pdfext/common"
	_cbd "github.com/bamzi/pdfext/core/security"
	_gbe "github.com/bamzi/pdfext/core/security/crypt"
	_bad "github.com/bamzi/pdfext/internal/ccittfax"
	_eeb "github.com/bamzi/pdfext/internal/imageutil"
	_gd "github.com/bamzi/pdfext/internal/jbig2"
	_ec "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_gae "github.com/bamzi/pdfext/internal/jbig2/decoder"
	_ce "github.com/bamzi/pdfext/internal/jbig2/document"
	_eca "github.com/bamzi/pdfext/internal/jbig2/errors"
	_dd "github.com/bamzi/pdfext/internal/strutils"
	_cg "golang.org/x/image/tiff/lzw"
	_cf "golang.org/x/xerrors"
)

// GetNameVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetNameVal(obj PdfObject) (_dbfa string, _gfgdbf bool) {
	_agec, _gfgdbf := TraceToDirectObject(obj).(*PdfObjectName)
	if _gfgdbf {
		return string(*_agec), true
	}
	return
}

// DecodeBytes decodes a slice of JBIG2 encoded bytes and returns the results.
func (_dgadg *JBIG2Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	return _gd.DecodeBytes(encoded, _gae.Parameters{}, _dgadg.Globals)
}

// DecodeBytes decodes a slice of ASCII encoded bytes and returns the result.
func (_bdf *ASCIIHexEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_afcb := _bg.NewReader(encoded)
	var _cfda []byte
	for {
		_face, _cabg := _afcb.ReadByte()
		if _cabg != nil {
			return nil, _cabg
		}
		if _face == '>' {
			break
		}
		if IsWhiteSpace(_face) {
			continue
		}
		if (_face >= 'a' && _face <= 'f') || (_face >= 'A' && _face <= 'F') || (_face >= '0' && _face <= '9') {
			_cfda = append(_cfda, _face)
		} else {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0061\u0073\u0063\u0069\u0069 \u0068\u0065\u0078\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072 \u0028\u0025\u0063\u0029", _face)
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0073\u0063\u0069\u0069\u0020\u0068e\u0078 \u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0028\u0025\u0063\u0029", _face)
		}
	}
	if len(_cfda)%2 == 1 {
		_cfda = append(_cfda, '0')
	}
	_eb.Log.Trace("\u0049\u006e\u0062\u006f\u0075\u006e\u0064\u0020\u0025\u0073", _cfda)
	_dfbf := make([]byte, _ae.DecodedLen(len(_cfda)))
	_, _afaa := _ae.Decode(_dfbf, _cfda)
	if _afaa != nil {
		return nil, _afaa
	}
	return _dfbf, nil
}
func (_adfb *offsetReader) Read(p []byte) (_bcdgb int, _dedfe error) { return _adfb._bfdd.Read(p) }
func _ddee(_ceaae int) int {
	if _ceaae < 0 {
		return -_ceaae
	}
	return _ceaae
}

// GetEncryptObj returns the PdfIndirectObject which has information about the PDFs encryption details.
func (_gbcgb *PdfParser) GetEncryptObj() *PdfIndirectObject { return _gbcgb._gbcg }
func _ed(_gff PdfObject) (int64, int64, error) {
	if _ac, _agf := _gff.(*PdfIndirectObject); _agf {
		return _ac.ObjectNumber, _ac.GenerationNumber, nil
	}
	if _gac, _acd := _gff.(*PdfObjectStream); _acd {
		return _gac.ObjectNumber, _gac.GenerationNumber, nil
	}
	return 0, 0, _f.New("\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u002f\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062je\u0063\u0074")
}
func _fcee(_eebb, _ecac, _bgegb uint8) uint8 {
	_dcbe := int(_bgegb)
	_gdddd := int(_ecac) - _dcbe
	_bddd := int(_eebb) - _dcbe
	_dcbe = _gged(_gdddd + _bddd)
	_gdddd = _gged(_gdddd)
	_bddd = _gged(_bddd)
	if _gdddd <= _bddd && _gdddd <= _dcbe {
		return _eebb
	} else if _bddd <= _dcbe {
		return _ecac
	}
	return _bgegb
}
func (_caa *PdfParser) parseDetailedHeader() (_aeba error) {
	_caa._cdea.Seek(0, _bb.SeekStart)
	_caa._dedbc = _ga.NewReader(_caa._cdea)
	_cgb := 20
	_dgc := make([]byte, _cgb)
	var (
		_bfc bool
		_bdd int
	)
	for {
		_dcf, _fff := _caa._dedbc.ReadByte()
		if _fff != nil {
			if _fff == _bb.EOF {
				break
			} else {
				return _fff
			}
		}
		if IsDecimalDigit(_dcf) && _dgc[_cgb-1] == '.' && IsDecimalDigit(_dgc[_cgb-2]) && _dgc[_cgb-3] == '-' && _dgc[_cgb-4] == 'F' && _dgc[_cgb-5] == 'D' && _dgc[_cgb-6] == 'P' && _dgc[_cgb-7] == '%' {
			_caa._fdff = Version{Major: int(_dgc[_cgb-2] - '0'), Minor: int(_dcf - '0')}
			_caa._eggc._gcc = _bdd - 7
			_bfc = true
			break
		}
		_bdd++
		_dgc = append(_dgc[1:_cgb], _dcf)
	}
	if !_bfc {
		return _ee.Errorf("n\u006f \u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066ou\u006e\u0064")
	}
	_abd, _aeba := _caa._dedbc.ReadByte()
	if _aeba == _bb.EOF {
		return _ee.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0050d\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _aeba != nil {
		return _aeba
	}
	_caa._eggc._aced = _abd == '\n'
	_abd, _aeba = _caa._dedbc.ReadByte()
	if _aeba != nil {
		return _ee.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _aeba)
	}
	if _abd != '%' {
		return nil
	}
	_bagc := make([]byte, 4)
	_, _aeba = _caa._dedbc.Read(_bagc)
	if _aeba != nil {
		return _ee.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _aeba)
	}
	_caa._eggc._affb = [4]byte{_bagc[0], _bagc[1], _bagc[2], _bagc[3]}
	return nil
}

// GetObjectStreams returns the *PdfObjectStreams represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetObjectStreams(obj PdfObject) (_afac *PdfObjectStreams, _dggf bool) {
	_afac, _dggf = obj.(*PdfObjectStreams)
	return _afac, _dggf
}

// MakeStream creates an PdfObjectStream with specified contents and encoding. If encoding is nil, then raw encoding
// will be used (i.e. no encoding applied).
func MakeStream(contents []byte, encoder StreamEncoder) (*PdfObjectStream, error) {
	_aaaba := &PdfObjectStream{}
	if encoder == nil {
		encoder = NewRawEncoder()
	}
	_aaaba.PdfObjectDictionary = encoder.MakeStreamDict()
	_gfgdb, _bfede := encoder.EncodeBytes(contents)
	if _bfede != nil {
		return nil, _bfede
	}
	_aaaba.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_gfgdb))))
	_aaaba.Stream = _gfgdb
	return _aaaba, nil
}

// PdfObjectStreams represents the primitive PDF object streams.
// 7.5.7 Object Streams (page 45).
type PdfObjectStreams struct {
	PdfObjectReference
	_fagda []PdfObject
}

// GetFileOffset returns the current file offset, accounting for buffered position.
func (_cfcf *PdfParser) GetFileOffset() int64 {
	_gfe, _ := _cfcf._cdea.Seek(0, _bb.SeekCurrent)
	_gfe -= int64(_cfcf._dedbc.Buffered())
	return _gfe
}

// NewParser creates a new parser for a PDF file via ReadSeeker. Loads the cross reference stream and trailer.
// An error is returned on failure.
func NewParser(rs _bb.ReadSeeker) (*PdfParser, error) {
	_egae := &PdfParser{_cdea: rs, ObjCache: make(objectCache), _daad: map[int64]bool{}, _bagd: make([]int64, 0), _bddb: make(map[*PdfParser]*PdfParser)}
	_beef, _cdab, _gccc := _egae.parsePdfVersion()
	if _gccc != nil {
		_eb.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _gccc)
		return nil, _gccc
	}
	_egae._fdff.Major = _beef
	_egae._fdff.Minor = _cdab
	if _egae._cbeb, _gccc = _egae.loadXrefs(); _gccc != nil {
		_eb.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _gccc)
		return nil, _gccc
	}
	_eb.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _egae._cbeb)
	_eecd, _gccc := _egae.parseLinearizedDictionary()
	if _gccc != nil {
		return nil, _gccc
	}
	if _eecd != nil {
		_egae._aggcf, _gccc = _egae.checkLinearizedInformation(_eecd)
		if _gccc != nil {
			return nil, _gccc
		}
	}
	if len(_egae._bfba.ObjectMap) == 0 {
		return nil, _ee.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	_egae._aegg = len(_egae._bagd)
	if _egae._aggcf && _egae._aegg != 0 {
		_egae._aegg--
	}
	_egae._gfgdg = make([]*PdfParser, _egae._aegg)
	return _egae, nil
}
func _cfdbd(_gfebg string) (int, int, error) {
	_cfabe := _ffdaa.FindStringSubmatch(_gfebg)
	if len(_cfabe) < 3 {
		return 0, 0, _f.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_aadgb, _ := _d.Atoi(_cfabe[1])
	_dece, _ := _d.Atoi(_cfabe[2])
	return _aadgb, _dece, nil
}

// HeaderCommentBytes gets the header comment bytes.
func (_afae ParserMetadata) HeaderCommentBytes() [4]byte { return _afae._affb }

// RegisterCustomStreamEncoder register a custom encoder handler for certain filter.
func RegisterCustomStreamEncoder(filterName string, customStreamEncoder StreamEncoder) {
	_eedgb.Store(filterName, customStreamEncoder)
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_aacdd *RunLengthEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_dgag := _bg.NewReader(data)
	var _efgg []byte
	var _eafc []byte
	_ecfb, _dfbbf := _dgag.ReadByte()
	if _dfbbf == _bb.EOF {
		return []byte{}, nil
	} else if _dfbbf != nil {
		return nil, _dfbbf
	}
	_egc := 1
	for {
		_fcae, _aaf := _dgag.ReadByte()
		if _aaf == _bb.EOF {
			break
		} else if _aaf != nil {
			return nil, _aaf
		}
		if _fcae == _ecfb {
			if len(_eafc) > 0 {
				_eafc = _eafc[:len(_eafc)-1]
				if len(_eafc) > 0 {
					_efgg = append(_efgg, byte(len(_eafc)-1))
					_efgg = append(_efgg, _eafc...)
				}
				_egc = 1
				_eafc = []byte{}
			}
			_egc++
			if _egc >= 127 {
				_efgg = append(_efgg, byte(257-_egc), _ecfb)
				_egc = 0
			}
		} else {
			if _egc > 0 {
				if _egc == 1 {
					_eafc = []byte{_ecfb}
				} else {
					_efgg = append(_efgg, byte(257-_egc), _ecfb)
				}
				_egc = 0
			}
			_eafc = append(_eafc, _fcae)
			if len(_eafc) >= 127 {
				_efgg = append(_efgg, byte(len(_eafc)-1))
				_efgg = append(_efgg, _eafc...)
				_eafc = []byte{}
			}
		}
		_ecfb = _fcae
	}
	if len(_eafc) > 0 {
		_efgg = append(_efgg, byte(len(_eafc)-1))
		_efgg = append(_efgg, _eafc...)
	} else if _egc > 0 {
		_efgg = append(_efgg, byte(257-_egc), _ecfb)
	}
	_efgg = append(_efgg, 128)
	return _efgg, nil
}

// Append appends PdfObject(s) to the streams.
func (_aadg *PdfObjectStreams) Append(objects ...PdfObject) {
	if _aadg == nil {
		_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0020-\u0020\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0074\u006f\u0020\u0061p\u0070\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073")
		return
	}
	_aadg._fagda = append(_aadg._fagda, objects...)
}

// NewMultiEncoder returns a new instance of MultiEncoder.
func NewMultiEncoder() *MultiEncoder {
	_bfce := MultiEncoder{}
	_bfce._abeab = []StreamEncoder{}
	return &_bfce
}

type objectCache map[int]PdfObject

func (_fdg *PdfCrypt) securityHandler() _cbd.StdHandler {
	if _fdg._adfe.R >= 5 {
		return _cbd.NewHandlerR6()
	}
	return _cbd.NewHandlerR4(_fdg._ffc, _fdg._ffe.Length)
}
func (_cece *PdfCrypt) authenticate(_egf []byte) (bool, error) {
	_cece._agg = false
	_ecgb := _cece.securityHandler()
	_fdfd, _fga, _cdcf := _ecgb.Authenticate(&_cece._adfe, _egf)
	if _cdcf != nil {
		return false, _cdcf
	} else if _fga == 0 || len(_fdfd) == 0 {
		return false, nil
	}
	_cece._agg = true
	_cece._cbcg = _fdfd
	return true, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_eef *RunLengthEncoder) MakeStreamDict() *PdfObjectDictionary {
	_begb := MakeDict()
	_begb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_eef.GetFilterName()))
	return _begb
}

// UpdateParams updates the parameter values of the encoder.
func (_baea *RawEncoder) UpdateParams(params *PdfObjectDictionary) {}

// UpdateParams updates the parameter values of the encoder.
// Implements StreamEncoder interface.
func (_cbe *JBIG2Encoder) UpdateParams(params *PdfObjectDictionary) {
	_ccgef, _afca := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _afca == nil {
		_cbe.BitsPerComponent = int(_ccgef)
	}
	_fcdb, _afca := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _afca == nil {
		_cbe.Width = int(_fcdb)
	}
	_aefc, _afca := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _afca == nil {
		_cbe.Height = int(_aefc)
	}
	_gedb, _afca := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _afca == nil {
		_cbe.ColorComponents = int(_gedb)
	}
}
func (_ebbe *PdfParser) parseObject() (PdfObject, error) {
	_eb.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_ebbe.skipSpaces()
	for {
		_efge, _cbbgf := _ebbe._dedbc.Peek(2)
		if _cbbgf != nil {
			if _cbbgf != _bb.EOF || len(_efge) == 0 {
				return nil, _cbbgf
			}
			if len(_efge) == 1 {
				_efge = append(_efge, ' ')
			}
		}
		_eb.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_efge))
		if _efge[0] == '/' {
			_abgc, _ccaa := _ebbe.parseName()
			_eb.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _abgc)
			return &_abgc, _ccaa
		} else if _efge[0] == '(' {
			_eb.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_bcge, _bbffc := _ebbe.parseString()
			return _bcge, _bbffc
		} else if _efge[0] == '[' {
			_eb.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_fbfg, _beca := _ebbe.parseArray()
			return _fbfg, _beca
		} else if (_efge[0] == '<') && (_efge[1] == '<') {
			_eb.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			_dccff, _agcb := _ebbe.ParseDict()
			return _dccff, _agcb
		} else if _efge[0] == '<' {
			_eb.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			_efdg, _aefb := _ebbe.parseHexString()
			return _efdg, _aefb
		} else if _efge[0] == '%' {
			_ebbe.readComment()
			_ebbe.skipSpaces()
		} else {
			_eb.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_efge, _ = _ebbe._dedbc.Peek(15)
			_eced := string(_efge)
			_eb.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _eced)
			if (len(_eced) > 3) && (_eced[:4] == "\u006e\u0075\u006c\u006c") {
				_ecda, _cbaa := _ebbe.parseNull()
				return &_ecda, _cbaa
			} else if (len(_eced) > 4) && (_eced[:5] == "\u0066\u0061\u006cs\u0065") {
				_ccafd, _bedf := _ebbe.parseBool()
				return &_ccafd, _bedf
			} else if (len(_eced) > 3) && (_eced[:4] == "\u0074\u0072\u0075\u0065") {
				_cade, _cbaf := _ebbe.parseBool()
				return &_cade, _cbaf
			}
			_dbcf := _fgaa.FindStringSubmatch(_eced)
			if len(_dbcf) > 1 {
				_efge, _ = _ebbe._dedbc.ReadBytes('R')
				_eb.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_efge[:]))
				_cbdf, _fbge := _fdae(string(_efge))
				_cbdf._dcge = _ebbe
				return &_cbdf, _fbge
			}
			_abeffg := _gdga.FindStringSubmatch(_eced)
			if len(_abeffg) > 1 {
				_eb.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				_dgee, _eceg := _ebbe.parseNumber()
				return _dgee, _eceg
			}
			_abeffg = _baed.FindStringSubmatch(_eced)
			if len(_abeffg) > 1 {
				_eb.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_eb.Log.Trace("\u0025\u0020\u0073", _abeffg)
				_bddc, _gfeg := _ebbe.parseNumber()
				return _bddc, _gfeg
			}
			_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _eced)
			return nil, _f.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}

// EncodeBytes JPX encodes the passed in slice of bytes.
func (_cebg *JPXEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _cebg.GetFilterName())
	return data, ErrNoJPXDecode
}
func _gaf(_cbfg *_cbd.StdEncryptDict, _bac *PdfObjectDictionary) error {
	R, _dce := _bac.Get("\u0052").(*PdfObjectInteger)
	if !_dce {
		return _f.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0052")
	}
	if *R < 2 || *R > 6 {
		return _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052 \u0028\u0025\u0064\u0029", *R)
	}
	_cbfg.R = int(*R)
	O, _dce := _bac.GetString("\u004f")
	if !_dce {
		return _f.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u004f")
	}
	if _cbfg.R == 5 || _cbfg.R == 6 {
		if len(O) < 48 {
			return _ee.Errorf("\u004c\u0065\u006e\u0067th\u0028\u004f\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(O))
		}
	} else if len(O) != 32 {
		return _ee.Errorf("L\u0065n\u0067\u0074\u0068\u0028\u004f\u0029\u0020\u0021=\u0020\u0033\u0032\u0020(%\u0064\u0029", len(O))
	}
	_cbfg.O = []byte(O)
	U, _dce := _bac.GetString("\u0055")
	if !_dce {
		return _f.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0055")
	}
	if _cbfg.R == 5 || _cbfg.R == 6 {
		if len(U) < 48 {
			return _ee.Errorf("\u004c\u0065\u006e\u0067th\u0028\u0055\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(U))
		}
	} else if len(U) != 32 {
		_eb.Log.Debug("\u0057\u0061r\u006e\u0069\u006e\u0067\u003a\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0028\u0055\u0029\u0020\u0021\u003d\u0020\u0033\u0032\u0020(%\u0064\u0029", len(U))
	}
	_cbfg.U = []byte(U)
	if _cbfg.R >= 5 {
		OE, _agd := _bac.GetString("\u004f\u0045")
		if !_agd {
			return _f.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u004f\u0045")
		} else if len(OE) != 32 {
			return _ee.Errorf("L\u0065\u006e\u0067\u0074h(\u004fE\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(OE))
		}
		_cbfg.OE = []byte(OE)
		UE, _agd := _bac.GetString("\u0055\u0045")
		if !_agd {
			return _f.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u0055\u0045")
		} else if len(UE) != 32 {
			return _ee.Errorf("L\u0065\u006e\u0067\u0074h(\u0055E\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(UE))
		}
		_cbfg.UE = []byte(UE)
	}
	P, _dce := _bac.Get("\u0050").(*PdfObjectInteger)
	if !_dce {
		return _f.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074 \u0064\u0069\u0063t\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0061\u0074\u0074\u0072")
	}
	_cbfg.P = _cbd.Permissions(*P)
	if _cbfg.R == 6 {
		Perms, _bbb := _bac.GetString("\u0050\u0065\u0072m\u0073")
		if !_bbb {
			return _f.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0050\u0065\u0072\u006d\u0073")
		} else if len(Perms) != 16 {
			return _ee.Errorf("\u004ce\u006e\u0067\u0074\u0068\u0028\u0050\u0065\u0072\u006d\u0073\u0029 \u0021\u003d\u0020\u0031\u0036\u0020\u0028\u0025\u0064\u0029", len(Perms))
		}
		_cbfg.Perms = []byte(Perms)
	}
	if _cga, _bcd := _bac.Get("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061").(*PdfObjectBool); _bcd {
		_cbfg.EncryptMetadata = bool(*_cga)
	} else {
		_cbfg.EncryptMetadata = true
	}
	return nil
}

// Seek implementation of Seek interface.
func (_adgf *limitedReadSeeker) Seek(offset int64, whence int) (int64, error) {
	var _fgab int64
	switch whence {
	case _bb.SeekStart:
		_fgab = offset
	case _bb.SeekCurrent:
		_agef, _ffega := _adgf._efag.Seek(0, _bb.SeekCurrent)
		if _ffega != nil {
			return 0, _ffega
		}
		_fgab = _agef + offset
	case _bb.SeekEnd:
		_fgab = _adgf._befa + offset
	}
	if _gegc := _adgf.getError(_fgab); _gegc != nil {
		return 0, _gegc
	}
	if _, _dccbc := _adgf._efag.Seek(_fgab, _bb.SeekStart); _dccbc != nil {
		return 0, _dccbc
	}
	return _fgab, nil
}

// DCTEncoder provides a DCT (JPG) encoding/decoding functionality for images.
type DCTEncoder struct {
	ColorComponents  int
	BitsPerComponent int
	Width            int
	Height           int
	Quality          int
	Decode           []float64
}

// PdfParser parses a PDF file and provides access to the object structure of the PDF.
type PdfParser struct {
	_fdff    Version
	_cdea    _bb.ReadSeeker
	_dedbc   *_ga.Reader
	_ggdf    int64
	_bfba    XrefTable
	_cgfg    int64
	_gfggg   *xrefType
	_fggc    objectStreams
	_cbeb    *PdfObjectDictionary
	_eccc    *PdfCrypt
	_gbcg    *PdfIndirectObject
	_egcga   bool
	ObjCache objectCache
	_cefg    map[int]bool
	_daad    map[int64]bool
	_eggc    ParserMetadata
	_aega    bool
	_bagd    []int64
	_aegg    int
	_aggcf   bool
	_ggfc    int64
	_bddb    map[*PdfParser]*PdfParser
	_gfgdg   []*PdfParser
}

func _gadd(_bagb, _bbfe, _cdaa int) error {
	if _bbfe < 0 || _bbfe > _bagb {
		return _f.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0061\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	if _cdaa < _bbfe {
		return _f.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0073\u006c\u0069\u0063e\u0020i\u006ed\u0065\u0078\u0020\u0062\u0020\u003c\u0020a")
	}
	if _cdaa > _bagb {
		return _f.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0062\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	return nil
}

// ParserMetadata gets the pdf parser metadata.
func (_cae *PdfParser) ParserMetadata() (ParserMetadata, error) {
	if !_cae._aega {
		return ParserMetadata{}, _ee.Errorf("\u0070\u0061\u0072\u0073\u0065r\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0072\u006be\u0064\u0020\u0066\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0064\u0065\u0074\u0061\u0069\u006c\u0065\u0064\u0020\u006d\u0065\u0074\u0061\u0064\u0061\u0074a")
	}
	return _cae._eggc, nil
}

// PdfIndirectObject represents the primitive PDF indirect object.
type PdfIndirectObject struct {
	PdfObjectReference
	PdfObject
}

// UpdateParams updates the parameter values of the encoder.
func (_adfa *FlateEncoder) UpdateParams(params *PdfObjectDictionary) {
	_ecfg, _afea := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _afea == nil {
		_adfa.Predictor = int(_ecfg)
	}
	_fcfa, _afea := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _afea == nil {
		_adfa.BitsPerComponent = int(_fcfa)
	}
	_bcdg, _afea := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _afea == nil {
		_adfa.Columns = int(_bcdg)
	}
	_gbg, _afea := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _afea == nil {
		_adfa.Colors = int(_gbg)
	}
}
func (_gcf *PdfCrypt) isEncrypted(_febg PdfObject) bool {
	_, _bbd := _gcf._edc[_febg]
	if _bbd {
		_eb.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0065\u006e\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	_eb.Log.Trace("\u004e\u006f\u0074\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_bacf *LZWEncoder) MakeStreamDict() *PdfObjectDictionary {
	_edg := MakeDict()
	_edg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_bacf.GetFilterName()))
	_faa := _bacf.MakeDecodeParams()
	if _faa != nil {
		_edg.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _faa)
	}
	_edg.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(int64(_bacf.EarlyChange)))
	return _edg
}

// WriteString outputs the object as it is to be written to file.
func (_ccefc *PdfObjectName) WriteString() string {
	var _bdded _bg.Buffer
	if len(*_ccefc) > 127 {
		_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u0061\u006d\u0065\u0020t\u006fo\u0020l\u006f\u006e\u0067\u0020\u0028\u0025\u0073)", *_ccefc)
	}
	_bdded.WriteString("\u002f")
	for _bbdd := 0; _bbdd < len(*_ccefc); _bbdd++ {
		_eead := (*_ccefc)[_bbdd]
		if !IsPrintable(_eead) || _eead == '#' || IsDelimiter(_eead) {
			_bdded.WriteString(_ee.Sprintf("\u0023\u0025\u002e2\u0078", _eead))
		} else {
			_bdded.WriteByte(_eead)
		}
	}
	return _bdded.String()
}

// DecodeStream decodes a JBIG2 encoded stream and returns the result as a slice of bytes.
func (_bcgfb *JBIG2Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _bcgfb.DecodeBytes(streamObj.Stream)
}

// DecodeBytes decodes a slice of Flate encoded bytes and returns the result.
func (_dbbe *FlateEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_eb.Log.Trace("\u0046\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065\u0020b\u0079\u0074\u0065\u0073")
	if len(encoded) == 0 {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0046\u006c\u0061\u0074\u0065 e\u006ec\u006f\u0064\u0065\u0064\u0020\u0062\u0075\u0066\u0066\u0065\u0072\u002e \u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0062y\u0074\u0065\u0020\u0073\u006c\u0069\u0063\u0065\u002e")
		return []byte{}, nil
	}
	_fge := _bg.NewReader(encoded)
	_bebc, _adg := _e.NewReader(_fge)
	if _adg != nil {
		_eb.Log.Debug("\u0044e\u0063o\u0064\u0069\u006e\u0067\u0020e\u0072\u0072o\u0072\u0020\u0025\u0076\u000a", _adg)
		_eb.Log.Debug("\u0053t\u0072e\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(encoded), encoded)
		return nil, _adg
	}
	defer _bebc.Close()
	var _feae _bg.Buffer
	_feae.ReadFrom(_bebc)
	return _feae.Bytes(), nil
}

// Decrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_bga *PdfCrypt) Decrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _bga.isDecrypted(obj) {
		return nil
	}
	switch _dbgf := obj.(type) {
	case *PdfIndirectObject:
		_bga._ge[_dbgf] = true
		_eb.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _dbgf.ObjectNumber, _dbgf.GenerationNumber)
		_gdfa := _dbgf.ObjectNumber
		_fec := _dbgf.GenerationNumber
		_bfg := _bga.Decrypt(_dbgf.PdfObject, _gdfa, _fec)
		if _bfg != nil {
			return _bfg
		}
		return nil
	case *PdfObjectStream:
		_bga._ge[_dbgf] = true
		_ggb := _dbgf.PdfObjectDictionary
		if _bga._adfe.R != 5 {
			if _gga, _fceb := _ggb.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _fceb && *_gga == "\u0058\u0052\u0065\u0066" {
				return nil
			}
		}
		_dabc := _dbgf.ObjectNumber
		_aff := _dbgf.GenerationNumber
		_eb.Log.Trace("\u0044e\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _dabc, _aff)
		_fedd := _dedd
		if _bga._ffe.V >= 4 {
			_fedd = _bga._badb
			_eb.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _bga._badb)
			if _def, _eeg := _ggb.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _eeg {
				if _acdb, _fgag := GetName(_def.Get(0)); _fgag {
					if *_acdb == "\u0043\u0072\u0079p\u0074" {
						_fedd = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _gee, _ccf := _ggb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _ccf {
							if _ffb, _efg := _gee.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _efg {
								if _, _efca := _bga._ace[string(*_ffb)]; _efca {
									_eb.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_ffb)
									_fedd = string(*_ffb)
								}
							}
						}
					}
				}
			}
			_eb.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _fedd)
			if _fedd == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_ceg := _bga.Decrypt(_ggb, _dabc, _aff)
		if _ceg != nil {
			return _ceg
		}
		_ccg, _ceg := _bga.makeKey(_fedd, uint32(_dabc), uint32(_aff), _bga._cbcg)
		if _ceg != nil {
			return _ceg
		}
		_dbgf.Stream, _ceg = _bga.decryptBytes(_dbgf.Stream, _fedd, _ccg)
		if _ceg != nil {
			return _ceg
		}
		_ggb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_dbgf.Stream))))
		return nil
	case *PdfObjectString:
		_eb.Log.Trace("\u0044e\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_ebag := _dedd
		if _bga._ffe.V >= 4 {
			_eb.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bga._df)
			if _bga._df == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_ebag = _bga._df
		}
		_aef, _bgeb := _bga.makeKey(_ebag, uint32(parentObjNum), uint32(parentGenNum), _bga._cbcg)
		if _bgeb != nil {
			return _bgeb
		}
		_abfc := _dbgf.Str()
		_cef := make([]byte, len(_abfc))
		for _efed := 0; _efed < len(_abfc); _efed++ {
			_cef[_efed] = _abfc[_efed]
		}
		if len(_cef) > 0 {
			_eb.Log.Trace("\u0044e\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _cef, _cef)
			_cef, _bgeb = _bga.decryptBytes(_cef, _ebag, _aef)
			if _bgeb != nil {
				return _bgeb
			}
		}
		_dbgf._cbdg = string(_cef)
		return nil
	case *PdfObjectArray:
		for _, _gdcb := range _dbgf.Elements() {
			_afe := _bga.Decrypt(_gdcb, parentObjNum, parentGenNum)
			if _afe != nil {
				return _afe
			}
		}
		return nil
	case *PdfObjectDictionary:
		_dafe := false
		if _cada := _dbgf.Get("\u0054\u0079\u0070\u0065"); _cada != nil {
			_ebcg, _gfae := _cada.(*PdfObjectName)
			if _gfae && *_ebcg == "\u0053\u0069\u0067" {
				_dafe = true
			}
		}
		for _, _abad := range _dbgf.Keys() {
			_cfa := _dbgf.Get(_abad)
			if _dafe && string(_abad) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_abad) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_abad) != "\u0050\u0072\u0065\u0076" && string(_abad) != "\u004c\u0061\u0073\u0074" {
				_acb := _bga.Decrypt(_cfa, parentObjNum, parentGenNum)
				if _acb != nil {
					return _acb
				}
			}
		}
		return nil
	}
	return nil
}

// GetFilterName returns the name of the encoding filter.
func (_egg *LZWEncoder) GetFilterName() string { return StreamEncodingFilterNameLZW }

// String returns a string representation of the *PdfObjectString.
func (_eggb *PdfObjectString) String() string { return _eggb._cbdg }
func (_dfd *PdfCrypt) newEncryptDict() *PdfObjectDictionary {
	_dfb := MakeDict()
	_dfb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064"))
	_dfb.Set("\u0056", MakeInteger(int64(_dfd._ffe.V)))
	_dfb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_dfd._ffe.Length)))
	return _dfb
}
func (_bgag *PdfParser) parseLinearizedDictionary() (*PdfObjectDictionary, error) {
	_cdfgf, _cdgc := _bgag._cdea.Seek(0, _bb.SeekEnd)
	if _cdgc != nil {
		return nil, _cdgc
	}
	var _caeg int64
	var _cgda int64 = 2048
	for _caeg < _cdfgf-4 {
		if _cdfgf <= (_cgda + _caeg) {
			_cgda = _cdfgf - _caeg
		}
		_, _dcd := _bgag._cdea.Seek(_caeg, _bb.SeekStart)
		if _dcd != nil {
			return nil, _dcd
		}
		_efce := make([]byte, _cgda)
		_, _dcd = _bgag._cdea.Read(_efce)
		if _dcd != nil {
			return nil, _dcd
		}
		_eb.Log.Trace("\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0066i\u0072\u0073\u0074\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0022\u0025\u0073\u0022", string(_efce))
		_ecef := _ffdaa.FindAllStringIndex(string(_efce), -1)
		if _ecef != nil {
			_ebgde := _ecef[0]
			_eb.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _ecef)
			_, _gaaf := _bgag._cdea.Seek(int64(_ebgde[0]), _bb.SeekStart)
			if _gaaf != nil {
				return nil, _gaaf
			}
			_bgag._dedbc = _ga.NewReader(_bgag._cdea)
			_cgaa, _gaaf := _bgag.ParseIndirectObject()
			if _gaaf != nil {
				return nil, nil
			}
			if _daaa, _ecdg := GetIndirect(_cgaa); _ecdg {
				if _gbbc, _adeb := GetDict(_daaa.PdfObject); _adeb {
					if _cdaf := _gbbc.Get("\u004c\u0069\u006e\u0065\u0061\u0072\u0069\u007a\u0065\u0064"); _cdaf != nil {
						return _gbbc, nil
					}
					return nil, nil
				}
			}
			return nil, nil
		}
		_caeg += _cgda - 4
	}
	return nil, _f.New("\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064")
}

// GetDict returns the *PdfObjectDictionary represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetDict(obj PdfObject) (_cfbe *PdfObjectDictionary, _efcc bool) {
	_cfbe, _efcc = TraceToDirectObject(obj).(*PdfObjectDictionary)
	return _cfbe, _efcc
}

var _bfda = _ba.MustCompile("\u0025\u0025\u0045\u004f\u0046\u003f")

func _bdfa(_adee *PdfObjectStream, _afeae *PdfObjectDictionary) (*JBIG2Encoder, error) {
	const _dgefc = "\u006ee\u0077\u004a\u0042\u0049G\u0032\u0044\u0065\u0063\u006fd\u0065r\u0046r\u006f\u006d\u0053\u0074\u0072\u0065\u0061m"
	_dedg := NewJBIG2Encoder()
	_fede := _adee.PdfObjectDictionary
	if _fede == nil {
		return _dedg, nil
	}
	if _afeae == nil {
		_gfdd := _fede.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		if _gfdd != nil {
			switch _dgfae := _gfdd.(type) {
			case *PdfObjectDictionary:
				_afeae = _dgfae
			case *PdfObjectArray:
				if _dgfae.Len() == 1 {
					if _dada, _baeag := GetDict(_dgfae.Get(0)); _baeag {
						_afeae = _dada
					}
				}
			default:
				_eb.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065P\u0061\u0072\u0061\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0025\u0023\u0076", _gfdd)
				return nil, _eca.Errorf(_dgefc, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050a\u0072m\u0073\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _dgfae)
			}
		}
	}
	if _afeae == nil {
		return _dedg, nil
	}
	_dedg.UpdateParams(_afeae)
	_aebca, _abbd := GetStream(_afeae.Get("\u004a\u0042\u0049G\u0032\u0047\u006c\u006f\u0062\u0061\u006c\u0073"))
	if !_abbd {
		return _dedg, nil
	}
	var _debb error
	_dedg.Globals, _debb = _gd.DecodeGlobals(_aebca.Stream)
	if _debb != nil {
		_debb = _eca.Wrap(_debb, _dgefc, "\u0063\u006f\u0072\u0072u\u0070\u0074\u0065\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0065n\u0063\u006f\u0064\u0065\u0064\u0020\u0064a\u0074\u0061")
		_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _debb)
		return nil, _debb
	}
	return _dedg, nil
}
func (_ca *PdfParser) lookupObjectViaOS(_eea int, _ecc int) (PdfObject, error) {
	var _badg *_bg.Reader
	var _cec objectStream
	var _dc bool
	_cec, _dc = _ca._fggc[_eea]
	if !_dc {
		_ccc, _dcg := _ca.LookupByNumber(_eea)
		if _dcg != nil {
			_eb.Log.Debug("\u004d\u0069ss\u0069\u006e\u0067 \u006f\u0062\u006a\u0065ct \u0073tr\u0065\u0061\u006d\u0020\u0077\u0069\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u0020%\u0064", _eea)
			return nil, _dcg
		}
		_cbde, _gg := _ccc.(*PdfObjectStream)
		if !_gg {
			return nil, _f.New("i\u006e\u0076\u0061\u006cid\u0020o\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _ca._eccc != nil && !_ca._eccc.isDecrypted(_cbde) {
			return nil, _f.New("\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0064\u0065\u0063r\u0079\u0070\u0074\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d")
		}
		_gdg := _cbde.PdfObjectDictionary
		_eb.Log.Trace("\u0073o\u0020\u0064\u003a\u0020\u0025\u0073\n", _gdg.String())
		_da, _gg := _gdg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
		if !_gg {
			_eb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0061\u006c\u0077\u0061\u0079\u0073\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054\u0079\u0070\u0065")
			return nil, _f.New("\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020T\u0079\u0070\u0065")
		}
		if _cc.ToLower(string(*_da)) != "\u006f\u0062\u006a\u0073\u0074\u006d" {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0074\u0079\u0070\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0077\u0061\u0079\u0073 \u0062\u0065\u0020\u004f\u0062\u006a\u0053\u0074\u006d\u0020\u0021")
			return nil, _f.New("\u006f\u0062\u006a\u0065c\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074y\u0070e\u0020\u0021\u003d\u0020\u004f\u0062\u006aS\u0074\u006d")
		}
		N, _gg := _gdg.Get("\u004e").(*PdfObjectInteger)
		if !_gg {
			return nil, _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020i\u006e\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_cbf, _gg := _gdg.Get("\u0046\u0069\u0072s\u0074").(*PdfObjectInteger)
		if !_gg {
			return nil, _f.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0046\u0069\u0072\u0073\u0074\u0020i\u006e \u0073t\u0072e\u0061\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_eb.Log.Trace("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _da, *N)
		_fe, _dcg := DecodeStream(_cbde)
		if _dcg != nil {
			return nil, _dcg
		}
		_eb.Log.Trace("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _fe)
		_ef := _ca.GetFileOffset()
		defer func() { _ca.SetFileOffset(_ef) }()
		_badg = _bg.NewReader(_fe)
		_ca._dedbc = _ga.NewReader(_badg)
		_eb.Log.Trace("\u0050a\u0072s\u0069\u006e\u0067\u0020\u006ff\u0066\u0073e\u0074\u0020\u006d\u0061\u0070")
		_fef := map[int]int64{}
		for _gf := 0; _gf < int(*N); _gf++ {
			_ca.skipSpaces()
			_dgd, _eebg := _ca.parseNumber()
			if _eebg != nil {
				return nil, _eebg
			}
			_gag, _bef := _dgd.(*PdfObjectInteger)
			if !_bef {
				return nil, _f.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_ca.skipSpaces()
			_dgd, _eebg = _ca.parseNumber()
			if _eebg != nil {
				return nil, _eebg
			}
			_aa, _bef := _dgd.(*PdfObjectInteger)
			if !_bef {
				return nil, _f.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_eb.Log.Trace("\u006f\u0062j\u0020\u0025\u0064 \u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0025\u0064", *_gag, *_aa)
			_fef[int(*_gag)] = int64(*_cbf + *_aa)
		}
		_cec = objectStream{N: int(*N), _cgg: _fe, _faf: _fef}
		_ca._fggc[_eea] = _cec
	} else {
		_dda := _ca.GetFileOffset()
		defer func() { _ca.SetFileOffset(_dda) }()
		_badg = _bg.NewReader(_cec._cgg)
		_ca._dedbc = _ga.NewReader(_badg)
	}
	_ebc := _cec._faf[_ecc]
	_eb.Log.Trace("\u0041\u0043\u0054\u0055AL\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u003d\u0020%\u0064", _ecc, _ebc)
	_badg.Seek(_ebc, _bb.SeekStart)
	_ca._dedbc = _ga.NewReader(_badg)
	_dge, _ := _ca._dedbc.Peek(100)
	_eb.Log.Trace("\u004f\u0042\u004a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_dge))
	_db, _fb := _ca.parseObject()
	if _fb != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0072\u0065\u0061\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _fb)
		return nil, _fb
	}
	if _db == nil {
		return nil, _f.New("o\u0062\u006a\u0065\u0063t \u0063a\u006e\u006e\u006f\u0074\u0020b\u0065\u0020\u006e\u0075\u006c\u006c")
	}
	_cfe := PdfIndirectObject{}
	_cfe.ObjectNumber = int64(_ecc)
	_cfe.PdfObject = _db
	_cfe._dcge = _ca
	return &_cfe, nil
}

var _bfff = _ba.MustCompile("\u0073t\u0061r\u0074\u0078\u003f\u0072\u0065f\u005c\u0073*\u0028\u005c\u0064\u002b\u0029")

// GetIntVal returns the int value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetIntVal(obj PdfObject) (_gcgb int, _bcce bool) {
	_cgdga, _bcce := TraceToDirectObject(obj).(*PdfObjectInteger)
	if _bcce && _cgdga != nil {
		return int(*_cgdga), true
	}
	return 0, false
}

// UpdateParams updates the parameter values of the encoder.
func (_ddab *LZWEncoder) UpdateParams(params *PdfObjectDictionary) {
	_ebcb, _dbf := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _dbf == nil {
		_ddab.Predictor = int(_ebcb)
	}
	_ffdc, _dbf := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _dbf == nil {
		_ddab.BitsPerComponent = int(_ffdc)
	}
	_fac, _dbf := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _dbf == nil {
		_ddab.Columns = int(_fac)
	}
	_bcaf, _dbf := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _dbf == nil {
		_ddab.Colors = int(_bcaf)
	}
	_gfcd, _dbf := GetNumberAsInt64(params.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065"))
	if _dbf == nil {
		_ddab.EarlyChange = int(_gfcd)
	}
}
func _gged(_gdea int) int {
	_gfff := _gdea >> (_fddb - 1)
	return (_gdea ^ _gfff) - _gfff
}

// GetFilterName returns the name of the encoding filter.
func (_dgbe *RunLengthEncoder) GetFilterName() string { return StreamEncodingFilterNameRunLength }

// LZWEncoder provides LZW encoding/decoding functionality.
type LZWEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Colors  int

	// LZW algorithm setting.
	EarlyChange int
}

const (
	DefaultJPEGQuality = 75
)

func (_dde *FlateEncoder) postDecodePredict(_efgf []byte) ([]byte, error) {
	if _dde.Predictor > 1 {
		if _dde.Predictor == 2 {
			_eb.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_eb.Log.Trace("\u0043\u006f\u006c\u006f\u0072\u0073\u003a\u0020\u0025\u0064", _dde.Colors)
			_gaab := _dde.Columns * _dde.Colors
			if _gaab < 1 {
				return []byte{}, nil
			}
			_daef := len(_efgf) / _gaab
			if len(_efgf)%_gaab != 0 {
				_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_efgf), _gaab)
			}
			if _gaab%_dde.Colors != 0 {
				return nil, _ee.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _gaab, _dde.Colors)
			}
			if _gaab > len(_efgf) {
				_eb.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _gaab, len(_efgf))
				return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_eb.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_efgf), _efgf)
			_ebd := _bg.NewBuffer(nil)
			for _ccfd := 0; _ccfd < _daef; _ccfd++ {
				_ggad := _efgf[_gaab*_ccfd : _gaab*(_ccfd+1)]
				for _cggde := _dde.Colors; _cggde < _gaab; _cggde++ {
					_ggad[_cggde] += _ggad[_cggde-_dde.Colors]
				}
				_ebd.Write(_ggad)
			}
			_dccde := _ebd.Bytes()
			_eb.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_dccde), _dccde)
			return _dccde, nil
		} else if _dde.Predictor >= 10 && _dde.Predictor <= 15 {
			_eb.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_abdf := _dde.Columns*_dde.Colors + 1
			_egbf := len(_efgf) / _abdf
			if len(_efgf)%_abdf != 0 {
				return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_efgf), _abdf)
			}
			if _abdf > len(_efgf) {
				_eb.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _abdf, len(_efgf))
				return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_acga := _bg.NewBuffer(nil)
			_eb.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _dde.Columns)
			_eb.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_efgf), _abdf, _egbf)
			_dcea := make([]byte, _abdf)
			for _agdf := 0; _agdf < _abdf; _agdf++ {
				_dcea[_agdf] = 0
			}
			_geff := _dde.Colors
			for _fcgc := 0; _fcgc < _egbf; _fcgc++ {
				_gaee := _efgf[_abdf*_fcgc : _abdf*(_fcgc+1)]
				_afeb := _gaee[0]
				switch _afeb {
				case _cafb:
				case _ebeb:
					for _age := 1 + _geff; _age < _abdf; _age++ {
						_gaee[_age] += _gaee[_age-_geff]
					}
				case _ggef:
					for _ffed := 1; _ffed < _abdf; _ffed++ {
						_gaee[_ffed] += _dcea[_ffed]
					}
				case _bae:
					for _adcb := 1; _adcb < _geff+1; _adcb++ {
						_gaee[_adcb] += _dcea[_adcb] / 2
					}
					for _faee := _geff + 1; _faee < _abdf; _faee++ {
						_gaee[_faee] += byte((int(_gaee[_faee-_geff]) + int(_dcea[_faee])) / 2)
					}
				case _cecg:
					for _bebe := 1; _bebe < _abdf; _bebe++ {
						var _fgagg, _bgc, _egee byte
						_bgc = _dcea[_bebe]
						if _bebe >= _geff+1 {
							_fgagg = _gaee[_bebe-_geff]
							_egee = _dcea[_bebe-_geff]
						}
						_gaee[_bebe] += _fcee(_fgagg, _bgc, _egee)
					}
				default:
					_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029\u0020\u0040\u0072o\u0077\u0020\u0025\u0064", _afeb, _fcgc)
					return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _afeb)
				}
				copy(_dcea, _gaee)
				_acga.Write(_gaee[1:])
			}
			_bfgg := _acga.Bytes()
			return _bfgg, nil
		} else {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _dde.Predictor)
			return nil, _ee.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _dde.Predictor)
		}
	}
	return _efgf, nil
}

// MakeStringFromBytes creates an PdfObjectString from a byte array.
// This is more natural than MakeString as `data` is usually not utf-8 encoded.
func MakeStringFromBytes(data []byte) *PdfObjectString { return MakeString(string(data)) }

// GetFilterName returns the name of the encoding filter.
func (_ecd *JBIG2Encoder) GetFilterName() string { return StreamEncodingFilterNameJBIG2 }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_cbbd *JPXEncoder) MakeDecodeParams() PdfObject { return nil }

// Set sets the PdfObject at index i of the streams. An error is returned if the index is outside bounds.
func (_fgbf *PdfObjectStreams) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_fgbf._fagda) {
		return _f.New("\u004f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_fgbf._fagda[i] = obj
	return nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_fgagd *DCTEncoder) MakeDecodeParams() PdfObject { return nil }

type limitedReadSeeker struct {
	_efag _bb.ReadSeeker
	_befa int64
}

// ParseDict reads and parses a PDF dictionary object enclosed with '<<' and '>>'
func (_bfbe *PdfParser) ParseDict() (*PdfObjectDictionary, error) {
	_eb.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_dfg := MakeDict()
	_dfg._eddefe = _bfbe
	_ddeg, _ := _bfbe._dedbc.ReadByte()
	if _ddeg != '<' {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_ddeg, _ = _bfbe._dedbc.ReadByte()
	if _ddeg != '<' {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_bfbe.skipSpaces()
		_bfbe.skipComments()
		_fddd, _bgef := _bfbe._dedbc.Peek(2)
		if _bgef != nil {
			return nil, _bgef
		}
		_eb.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_fddd), string(_fddd))
		if (_fddd[0] == '>') && (_fddd[1] == '>') {
			_eb.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_bfbe._dedbc.ReadByte()
			_bfbe._dedbc.ReadByte()
			break
		}
		_eb.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_cabd, _bgef := _bfbe.parseName()
		_eb.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _cabd)
		if _bgef != nil {
			_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _bgef)
			return nil, _bgef
		}
		if len(_cabd) > 4 && _cabd[len(_cabd)-4:] == "\u006e\u0075\u006c\u006c" {
			_abbb := _cabd[0 : len(_cabd)-4]
			_eb.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _cabd)
			_eb.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _abbb)
			_bfbe.skipSpaces()
			_ddce, _ := _bfbe._dedbc.Peek(1)
			if _ddce[0] == '/' {
				_dfg.Set(_abbb, MakeNull())
				continue
			}
		}
		_bfbe.skipSpaces()
		_gfbc, _bgef := _bfbe.parseObject()
		if _bgef != nil {
			return nil, _bgef
		}
		_dfg.Set(_cabd, _gfbc)
		if _eb.Log.IsLogLevel(_eb.LogLevelTrace) {
			_eb.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _cabd, _gfbc.String())
		}
	}
	_eb.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _dfg, nil
}
func _gfad(_cce *_gbe.FilterDict, _aac *PdfObjectDictionary) error {
	if _bgg, _gdc := _aac.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _gdc {
		if _efa := string(*_bgg); _efa != "C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072" {
			_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020C\u0046\u0020\u0064ic\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074e\u0072\u0029", _efa)
		}
	}
	_ded, _daa := _aac.Get("\u0043\u0046\u004d").(*PdfObjectName)
	if !_daa {
		return _ee.Errorf("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0028\u004e\u006f\u006e\u0065\u0029")
	}
	_cce.CFM = string(*_ded)
	if _fg, _efb := _aac.Get("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et").(*PdfObjectName); _efb {
		_cce.AuthEvent = _cbd.AuthEvent(*_fg)
	} else {
		_cce.AuthEvent = _cbd.EventDocOpen
	}
	if _dbbg, _eaa := _aac.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _eaa {
		_cce.Length = int(*_dbbg)
	}
	return nil
}

// Keys returns the list of keys in the dictionary.
// If `d` is nil returns a nil slice.
func (_defdc *PdfObjectDictionary) Keys() []PdfObjectName {
	if _defdc == nil {
		return nil
	}
	return _defdc._begbb
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_dabg *RawEncoder) MakeDecodeParams() PdfObject { return nil }

// String returns a string describing `d`.
func (_bafc *PdfObjectDictionary) String() string {
	var _fcfde _cc.Builder
	_fcfde.WriteString("\u0044\u0069\u0063t\u0028")
	for _, _adca := range _bafc._begbb {
		_gfda := _bafc._abec[_adca]
		_fcfde.WriteString("\u0022" + _adca.String() + "\u0022\u003a\u0020")
		_fcfde.WriteString(_gfda.String())
		_fcfde.WriteString("\u002c\u0020")
	}
	_fcfde.WriteString("\u0029")
	return _fcfde.String()
}

// MakeArrayFromFloats creates an PdfObjectArray from a slice of float64s, where each array element is an
// PdfObjectFloat.
func MakeArrayFromFloats(vals []float64) *PdfObjectArray {
	_adaab := MakeArray()
	for _, _cabde := range vals {
		_adaab.Append(MakeFloat(_cabde))
	}
	return _adaab
}

const (
	XrefTypeTableEntry   xrefType = iota
	XrefTypeObjectStream xrefType = iota
)

// IsOctalDigit checks if a character can be part of an octal digit string.
func IsOctalDigit(c byte) bool { return '0' <= c && c <= '7' }

var _gdga = _ba.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_cfge *PdfParser) repairLocateXref() (int64, error) {
	_dfbfb := int64(1000)
	_cfge._cdea.Seek(-_dfbfb, _bb.SeekCurrent)
	_cbbc, _ffec := _cfge._cdea.Seek(0, _bb.SeekCurrent)
	if _ffec != nil {
		return 0, _ffec
	}
	_eaae := make([]byte, _dfbfb)
	_cfge._cdea.Read(_eaae)
	_aebbd := _ecfbb.FindAllStringIndex(string(_eaae), -1)
	if len(_aebbd) < 1 {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0065\u0070a\u0069\u0072\u003a\u0020\u0078\u0072\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021")
		return 0, _f.New("\u0072\u0065\u0070\u0061ir\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_fcba := int64(_aebbd[len(_aebbd)-1][0])
	_bcda := _cbbc + _fcba
	return _bcda, nil
}

// DecodeBytes decodes a slice of LZW encoded bytes and returns the result.
func (_ged *LZWEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _faeeg _bg.Buffer
	_acbg := _bg.NewReader(encoded)
	var _agge _bb.ReadCloser
	if _ged.EarlyChange == 1 {
		_agge = _cg.NewReader(_acbg, _cg.MSB, 8)
	} else {
		_agge = _eg.NewReader(_acbg, _eg.MSB, 8)
	}
	defer _agge.Close()
	if _, _ffg := _faeeg.ReadFrom(_agge); _ffg != nil {
		if _ffg != _bb.ErrUnexpectedEOF || _faeeg.Len() == 0 {
			return nil, _ffg
		}
		_eb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u004c\u005a\u0057\u0020\u0064\u0065\u0063\u006f\u0064i\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076\u002e \u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062e \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _ffg)
	}
	return _faeeg.Bytes(), nil
}

// GetArray returns the *PdfObjectArray represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetArray(obj PdfObject) (_afgc *PdfObjectArray, _dbfb bool) {
	_afgc, _dbfb = TraceToDirectObject(obj).(*PdfObjectArray)
	return _afgc, _dbfb
}

// EncodeBytes encodes data into ASCII85 encoded format.
func (_acfe *ASCII85Encoder) EncodeBytes(data []byte) ([]byte, error) {
	var _ebf _bg.Buffer
	for _bceb := 0; _bceb < len(data); _bceb += 4 {
		_gacdf := data[_bceb]
		_efgd := 1
		_abga := byte(0)
		if _bceb+1 < len(data) {
			_abga = data[_bceb+1]
			_efgd++
		}
		_dcb := byte(0)
		if _bceb+2 < len(data) {
			_dcb = data[_bceb+2]
			_efgd++
		}
		_bcdb := byte(0)
		if _bceb+3 < len(data) {
			_bcdb = data[_bceb+3]
			_efgd++
		}
		_dccb := (uint32(_gacdf) << 24) | (uint32(_abga) << 16) | (uint32(_dcb) << 8) | uint32(_bcdb)
		if _dccb == 0 {
			_ebf.WriteByte('z')
		} else {
			_gdgff := _acfe.base256Tobase85(_dccb)
			for _, _cacc := range _gdgff[:_efgd+1] {
				_ebf.WriteByte(_cacc + '!')
			}
		}
	}
	_ebf.WriteString("\u007e\u003e")
	return _ebf.Bytes(), nil
}

// NewFlateEncoder makes a new flate encoder with default parameters, predictor 1 and bits per component 8.
func NewFlateEncoder() *FlateEncoder {
	_bea := &FlateEncoder{}
	_bea.Predictor = 1
	_bea.BitsPerComponent = 8
	_bea.Colors = 1
	_bea.Columns = 1
	return _bea
}
func (_ggbga *PdfParser) parsePdfVersion() (int, int, error) {
	var _efbce int64 = 20
	_effdc := make([]byte, _efbce)
	_ggbga._cdea.Seek(0, _bb.SeekStart)
	_ggbga._cdea.Read(_effdc)
	var _cbbdd error
	var _ecbcd, _ebgg int
	if _beab := _bffg.FindStringSubmatch(string(_effdc)); len(_beab) < 3 {
		if _ecbcd, _ebgg, _cbbdd = _ggbga.seekPdfVersionTopDown(); _cbbdd != nil {
			_eb.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _cbbdd
		}
		_ggbga._cdea, _cbbdd = _aggd(_ggbga._cdea, _ggbga.GetFileOffset()-8)
		if _cbbdd != nil {
			return 0, 0, _cbbdd
		}
	} else {
		if _ecbcd, _cbbdd = _d.Atoi(_beab[1]); _cbbdd != nil {
			return 0, 0, _cbbdd
		}
		if _ebgg, _cbbdd = _d.Atoi(_beab[2]); _cbbdd != nil {
			return 0, 0, _cbbdd
		}
		_ggbga.SetFileOffset(0)
	}
	_ggbga._dedbc = _ga.NewReader(_ggbga._cdea)
	_eb.Log.Debug("\u0050\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _ecbcd, _ebgg)
	return _ecbcd, _ebgg, nil
}

// UpdateParams updates the parameter values of the encoder.
func (_fbcd *JPXEncoder) UpdateParams(params *PdfObjectDictionary) {}

// GetInt returns the *PdfObjectBool object that is represented by a PdfObject either directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetInt(obj PdfObject) (_addc *PdfObjectInteger, _gcedg bool) {
	_addc, _gcedg = TraceToDirectObject(obj).(*PdfObjectInteger)
	return _addc, _gcedg
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set.  Some other parameters are generated elsewhere.
func (_ffbdb *DCTEncoder) MakeStreamDict() *PdfObjectDictionary {
	_cdfe := MakeDict()
	_cdfe.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_ffbdb.GetFilterName()))
	return _cdfe
}
func (_bcdge *PdfParser) parseNull() (PdfObjectNull, error) {
	_, _aaaea := _bcdge._dedbc.Discard(4)
	return PdfObjectNull{}, _aaaea
}

// NewRawEncoder returns a new instace of RawEncoder.
func NewRawEncoder() *RawEncoder { return &RawEncoder{} }
func (_efe *PdfParser) lookupByNumberWrapper(_fbb int, _bc bool) (PdfObject, bool, error) {
	_aec, _gfa, _dab := _efe.lookupByNumber(_fbb, _bc)
	if _dab != nil {
		return nil, _gfa, _dab
	}
	if !_gfa && _efe._eccc != nil && _efe._eccc._agg && !_efe._eccc.isDecrypted(_aec) {
		_dca := _efe._eccc.Decrypt(_aec, 0, 0)
		if _dca != nil {
			return nil, _gfa, _dca
		}
	}
	return _aec, _gfa, nil
}
func (_eeeag *PdfParser) parseString() (*PdfObjectString, error) {
	_eeeag._dedbc.ReadByte()
	var _bcbgf _bg.Buffer
	_begcc := 1
	for {
		_bfceb, _eada := _eeeag._dedbc.Peek(1)
		if _eada != nil {
			return MakeString(_bcbgf.String()), _eada
		}
		if _bfceb[0] == '\\' {
			_eeeag._dedbc.ReadByte()
			_abcg, _bcgg := _eeeag._dedbc.ReadByte()
			if _bcgg != nil {
				return MakeString(_bcbgf.String()), _bcgg
			}
			if IsOctalDigit(_abcg) {
				_fagb, _cegdg := _eeeag._dedbc.Peek(2)
				if _cegdg != nil {
					return MakeString(_bcbgf.String()), _cegdg
				}
				var _dgbf []byte
				_dgbf = append(_dgbf, _abcg)
				for _, _fggdd := range _fagb {
					if IsOctalDigit(_fggdd) {
						_dgbf = append(_dgbf, _fggdd)
					} else {
						break
					}
				}
				_eeeag._dedbc.Discard(len(_dgbf) - 1)
				_eb.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _dgbf)
				_egbgc, _cegdg := _d.ParseUint(string(_dgbf), 8, 32)
				if _cegdg != nil {
					return MakeString(_bcbgf.String()), _cegdg
				}
				_bcbgf.WriteByte(byte(_egbgc))
				continue
			}
			switch _abcg {
			case 'n':
				_bcbgf.WriteRune('\n')
			case 'r':
				_bcbgf.WriteRune('\r')
			case 't':
				_bcbgf.WriteRune('\t')
			case 'b':
				_bcbgf.WriteRune('\b')
			case 'f':
				_bcbgf.WriteRune('\f')
			case '(':
				_bcbgf.WriteRune('(')
			case ')':
				_bcbgf.WriteRune(')')
			case '\\':
				_bcbgf.WriteRune('\\')
			}
			continue
		} else if _bfceb[0] == '(' {
			_begcc++
		} else if _bfceb[0] == ')' {
			_begcc--
			if _begcc == 0 {
				_eeeag._dedbc.ReadByte()
				break
			}
		}
		_cfbf, _ := _eeeag._dedbc.ReadByte()
		_bcbgf.WriteByte(_cfbf)
	}
	return MakeString(_bcbgf.String()), nil
}
func (_becc *JBIG2Encoder) encodeImage(_cdeb _a.Image) ([]byte, error) {
	const _begbd = "e\u006e\u0063\u006f\u0064\u0065\u0049\u006d\u0061\u0067\u0065"
	_affd, _abbg := GoImageToJBIG2(_cdeb, JB2ImageAutoThreshold)
	if _abbg != nil {
		return nil, _eca.Wrap(_abbg, _begbd, "\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069m\u0061g\u0065\u0020\u0074\u006f\u0020\u006a\u0062\u0069\u0067\u0032\u0020\u0069\u006d\u0067")
	}
	if _abbg = _becc.AddPageImage(_affd, &_becc.DefaultPageSettings); _abbg != nil {
		return nil, _eca.Wrap(_abbg, _begbd, "")
	}
	return _becc.Encode()
}

var _bcc = _ba.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064\u002b)\u005c\u0073\u002a\u0024")

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_gbfe *JBIG2Encoder) MakeStreamDict() *PdfObjectDictionary {
	_eegc := MakeDict()
	_eegc.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_gbfe.GetFilterName()))
	return _eegc
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_bgbe *RunLengthEncoder) MakeDecodeParams() PdfObject { return nil }

// EncodeImage encodes 'img' golang image.Image into jbig2 encoded bytes document using default encoder settings.
func (_efedb *JBIG2Encoder) EncodeImage(img _a.Image) ([]byte, error) { return _efedb.encodeImage(img) }

// GetObjectNums returns a sorted list of object numbers of the PDF objects in the file.
func (_gabc *PdfParser) GetObjectNums() []int {
	var _eafd []int
	for _, _gdffd := range _gabc._bfba.ObjectMap {
		_eafd = append(_eafd, _gdffd.ObjectNumber)
	}
	_be.Ints(_eafd)
	return _eafd
}
func (_gde *PdfCrypt) decryptBytes(_gfc []byte, _ffce string, _ffeg []byte) ([]byte, error) {
	_eb.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_dgba, _abc := _gde._ace[_ffce]
	if !_abc {
		return nil, _ee.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _ffce)
	}
	return _dgba.DecryptBytes(_gfc, _ffeg)
}

// String returns a string describing `streams`.
func (_bccb *PdfObjectStreams) String() string {
	return _ee.Sprintf("\u004f\u0062j\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0064", _bccb.ObjectNumber)
}

// UpdateParams updates the parameter values of the encoder.
func (_gfd *ASCIIHexEncoder) UpdateParams(params *PdfObjectDictionary) {}

// EncodeBytes DCT encodes the passed in slice of bytes.
func (_dfdb *DCTEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _dgef _a.Image
	if _dfdb.ColorComponents == 1 && _dfdb.BitsPerComponent == 8 {
		_dgef = &_a.Gray{Rect: _a.Rect(0, 0, _dfdb.Width, _dfdb.Height), Pix: data, Stride: _eeb.BytesPerLine(_dfdb.Width, _dfdb.BitsPerComponent, _dfdb.ColorComponents)}
	} else {
		var _eab error
		_dgef, _eab = _eeb.NewImage(_dfdb.Width, _dfdb.Height, _dfdb.BitsPerComponent, _dfdb.ColorComponents, data, nil, nil)
		if _eab != nil {
			return nil, _eab
		}
	}
	_abaf := _beg.Options{}
	_abaf.Quality = _dfdb.Quality
	var _gbebd _bg.Buffer
	if _bdg := _beg.Encode(&_gbebd, _dgef, &_abaf); _bdg != nil {
		return nil, _bdg
	}
	return _gbebd.Bytes(), nil
}

// GetFilterName returns the name of the encoding filter.
func (_bgba *ASCII85Encoder) GetFilterName() string { return StreamEncodingFilterNameASCII85 }

// JPXEncoder implements JPX encoder/decoder (dummy, for now)
// FIXME: implement
type JPXEncoder struct{}

// WriteString outputs the object as it is to be written to file.
func (_affae *PdfObjectFloat) WriteString() string {
	return _d.FormatFloat(float64(*_affae), 'f', -1, 64)
}

// GetPreviousRevisionReadSeeker returns ReadSeeker for the previous version of the Pdf document.
func (_bddbd *PdfParser) GetPreviousRevisionReadSeeker() (_bb.ReadSeeker, error) {
	if _eeff := _bddbd.seekToEOFMarker(_bddbd._ggdf - _cafga); _eeff != nil {
		return nil, _eeff
	}
	_eaaf, _gbgd := _bddbd._cdea.Seek(0, _bb.SeekCurrent)
	if _gbgd != nil {
		return nil, _gbgd
	}
	_eaaf += _cafga
	return _gbeba(_bddbd._cdea, _eaaf)
}

// Validate validates the page settings for the JBIG2 encoder.
func (_cdbed JBIG2EncoderSettings) Validate() error {
	const _ebde = "\u0076a\u006ci\u0064\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065\u0072"
	if _cdbed.Threshold < 0 || _cdbed.Threshold > 1.0 {
		return _eca.Errorf(_ebde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064\u0020\u0076a\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0076\u0027 \u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061n\u0067\u0065\u0020\u005b\u0030\u002e0\u002c\u0020\u0031.\u0030\u005d", _cdbed.Threshold)
	}
	if _cdbed.ResolutionX < 0 {
		return _eca.Errorf(_ebde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0078\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _cdbed.ResolutionX)
	}
	if _cdbed.ResolutionY < 0 {
		return _eca.Errorf(_ebde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0079\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _cdbed.ResolutionY)
	}
	if _cdbed.DefaultPixelValue != 0 && _cdbed.DefaultPixelValue != 1 {
		return _eca.Errorf(_ebde, "de\u0066\u0061u\u006c\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072 \u0074\u0068\u0065\u0020\u0062\u0069\u0074\u003a \u007b0\u002c\u0031}", _cdbed.DefaultPixelValue)
	}
	if _cdbed.Compression != JB2Generic {
		return _eca.Errorf(_ebde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0063\u006fm\u0070\u0072\u0065\u0073s\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	}
	return nil
}

// UpdateParams updates the parameter values of the encoder.
func (_cafbe *DCTEncoder) UpdateParams(params *PdfObjectDictionary) {
	_cdb, _ggag := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _ggag == nil {
		_cafbe.ColorComponents = int(_cdb)
	}
	_gead, _ggag := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _ggag == nil {
		_cafbe.BitsPerComponent = int(_gead)
	}
	_aaad, _ggag := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _ggag == nil {
		_cafbe.Width = int(_aaad)
	}
	_bcgf, _ggag := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _ggag == nil {
		_cafbe.Height = int(_bcgf)
	}
	_eag, _ggag := GetNumberAsInt64(params.Get("\u0051u\u0061\u006c\u0069\u0074\u0079"))
	if _ggag == nil {
		_cafbe.Quality = int(_eag)
	}
	_dffc, _dged := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _dged {
		_cafbe.Decode, _ggag = _dffc.ToFloat64Array()
		if _ggag != nil {
			_eb.Log.Error("F\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020de\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006eto\u0020\u0061r\u0072\u0061\u0079\u0073\u003a\u0020\u0025\u0076", _ggag)
		}
	}
}

// NewJBIG2Encoder creates a new JBIG2Encoder.
func NewJBIG2Encoder() *JBIG2Encoder { return &JBIG2Encoder{_eabf: _ce.InitEncodeDocument(false)} }
func _fefgg(_adce *PdfObjectStream) (*MultiEncoder, error) {
	_egfb := NewMultiEncoder()
	_aceaf := _adce.PdfObjectDictionary
	if _aceaf == nil {
		return _egfb, nil
	}
	var _ggc *PdfObjectDictionary
	var _gcfa []PdfObject
	_bfd := _aceaf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
	if _bfd != nil {
		_gbab, _bdc := _bfd.(*PdfObjectDictionary)
		if _bdc {
			_ggc = _gbab
		}
		_cbcf, _gcde := _bfd.(*PdfObjectArray)
		if _gcde {
			for _, _gdcbd := range _cbcf.Elements() {
				_gdcbd = TraceToDirectObject(_gdcbd)
				if _dbcb, _gcgg := _gdcbd.(*PdfObjectDictionary); _gcgg {
					_gcfa = append(_gcfa, _dbcb)
				} else {
					_gcfa = append(_gcfa, MakeDict())
				}
			}
		}
	}
	_bfd = _aceaf.Get("\u0046\u0069\u006c\u0074\u0065\u0072")
	if _bfd == nil {
		return nil, _ee.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_cdcd, _fcab := _bfd.(*PdfObjectArray)
	if !_fcab {
		return nil, _ee.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _aggc, _gafb := range _cdcd.Elements() {
		_bdacb, _fafea := _gafb.(*PdfObjectName)
		if !_fafea {
			return nil, _ee.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _fedc PdfObject
		if _ggc != nil {
			_fedc = _ggc
		} else {
			if len(_gcfa) > 0 {
				if _aggc >= len(_gcfa) {
					return nil, _ee.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_fedc = _gcfa[_aggc]
			}
		}
		var _fbf *PdfObjectDictionary
		if _cadfc, _ecge := _fedc.(*PdfObjectDictionary); _ecge {
			_fbf = _cadfc
		}
		_eb.Log.Trace("\u004e\u0065\u0078t \u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073\u002c \u0064p\u003a \u0025v\u002c\u0020\u0064\u0050\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u0076", *_bdacb, _fedc, _fbf)
		if *_bdacb == StreamEncodingFilterNameFlate {
			_bdgb, _ade := _fcd(_adce, _fbf)
			if _ade != nil {
				return nil, _ade
			}
			_egfb.AddEncoder(_bdgb)
		} else if *_bdacb == StreamEncodingFilterNameLZW {
			_gfgg, _ggab := _bagf(_adce, _fbf)
			if _ggab != nil {
				return nil, _ggab
			}
			_egfb.AddEncoder(_gfgg)
		} else if *_bdacb == StreamEncodingFilterNameASCIIHex {
			_fagd := NewASCIIHexEncoder()
			_egfb.AddEncoder(_fagd)
		} else if *_bdacb == StreamEncodingFilterNameASCII85 {
			_cdfb := NewASCII85Encoder()
			_egfb.AddEncoder(_cdfb)
		} else if *_bdacb == StreamEncodingFilterNameDCT {
			_baee, _dbce := _efbb(_adce, _egfb)
			if _dbce != nil {
				return nil, _dbce
			}
			_egfb.AddEncoder(_baee)
			_eb.Log.Trace("A\u0064d\u0065\u0064\u0020\u0044\u0043\u0054\u0020\u0065n\u0063\u006f\u0064\u0065r.\u002e\u002e")
			_eb.Log.Trace("\u004du\u006ct\u0069\u0020\u0065\u006e\u0063o\u0064\u0065r\u003a\u0020\u0025\u0023\u0076", _egfb)
		} else if *_bdacb == StreamEncodingFilterNameCCITTFax {
			_bafb, _deaf := _cgbba(_adce, _fbf)
			if _deaf != nil {
				return nil, _deaf
			}
			_egfb.AddEncoder(_bafb)
		} else {
			_eb.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_bdacb)
			return nil, _ee.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _egfb, nil
}

// PdfVersion returns version of the PDF file.
func (_cdde *PdfParser) PdfVersion() Version { return _cdde._fdff }

// RunLengthEncoder represents Run length encoding.
type RunLengthEncoder struct{}

// String returns a string describing `ref`.
func (_bcgc *PdfObjectReference) String() string {
	return _ee.Sprintf("\u0052\u0065\u0066\u0028\u0025\u0064\u0020\u0025\u0064\u0029", _bcgc.ObjectNumber, _bcgc.GenerationNumber)
}
func _aggd(_dedb _bb.ReadSeeker, _faffb int64) (*offsetReader, error) {
	_caad := &offsetReader{_bfdd: _dedb, _cbca: _faffb}
	_, _dgfg := _caad.Seek(0, _bb.SeekStart)
	return _caad, _dgfg
}

// DecodeStream decodes a DCT encoded stream and returns the result as a
// slice of bytes.
func (_bega *DCTEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _bega.DecodeBytes(streamObj.Stream)
}
func (_bag *PdfCrypt) saveCryptFilters(_bda *PdfObjectDictionary) error {
	if _bag._ffe.V < 4 {
		return _f.New("\u0063\u0061\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020V\u003e\u003d\u0034")
	}
	_bde := MakeDict()
	_bda.Set("\u0043\u0046", _bde)
	for _fdf, _ebe := range _bag._ace {
		if _fdf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			continue
		}
		_fca := _acag(_ebe, "")
		_bde.Set(PdfObjectName(_fdf), _fca)
	}
	_bda.Set("\u0053\u0074\u0072\u0046", MakeName(_bag._df))
	_bda.Set("\u0053\u0074\u006d\u0046", MakeName(_bag._badb))
	return nil
}

var _bgbf = _f.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")

// MultiEncoder supports serial encoding.
type MultiEncoder struct{ _abeab []StreamEncoder }

// GetFilterName returns the name of the encoding filter.
func (_edag *DCTEncoder) GetFilterName() string { return StreamEncodingFilterNameDCT }

// Set sets the PdfObject at index i of the array. An error is returned if the index is outside bounds.
func (_bfgdf *PdfObjectArray) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_bfgdf._aagc) {
		return _f.New("\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_bfgdf._aagc[i] = obj
	return nil
}
func _dcca(_gcee _eeb.Image) *JBIG2Image {
	_cfgc := _gcee.Base()
	return &JBIG2Image{Data: _cfgc.Data, Width: _cfgc.Width, Height: _cfgc.Height, HasPadding: true}
}
func (_ffeb *limitedReadSeeker) getError(_efab int64) error {
	switch {
	case _efab < 0:
		return _ee.Errorf("\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0065\u0067\u0061\u0074\u0069\u0076e\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064", _efab)
	case _efab > _ffeb._befa:
		return _ee.Errorf("u\u006e\u0065\u0078\u0070ec\u0074e\u0064\u0020\u006f\u0066\u0066s\u0065\u0074\u003a\u0020\u0025\u0064", _efab)
	}
	return nil
}

// PdfCryptNewEncrypt makes the document crypt handler based on a specified crypt filter.
func PdfCryptNewEncrypt(cf _gbe.Filter, userPass, ownerPass []byte, perm _cbd.Permissions) (*PdfCrypt, *EncryptInfo, error) {
	_cag := &PdfCrypt{_edc: make(map[PdfObject]bool), _ace: make(cryptFilters), _adfe: _cbd.StdEncryptDict{P: perm, EncryptMetadata: true}}
	var _dbbd Version
	if cf != nil {
		_bba := cf.PDFVersion()
		_dbbd.Major, _dbbd.Minor = _bba[0], _bba[1]
		V, R := cf.HandlerVersion()
		_cag._ffe.V = V
		_cag._adfe.R = R
		_cag._ffe.Length = cf.KeyLength() * 8
	}
	const (
		_beb = _dedd
	)
	_cag._ace[_beb] = cf
	if _cag._ffe.V >= 4 {
		_cag._badb = _beb
		_cag._df = _beb
	}
	_aba := _cag.newEncryptDict()
	_ea := _ab.Sum([]byte(_ag.Now().Format(_ag.RFC850)))
	_ead := string(_ea[:])
	_beee := make([]byte, 100)
	_dg.Read(_beee)
	_ea = _ab.Sum(_beee)
	_abg := string(_ea[:])
	_eb.Log.Trace("\u0052\u0061\u006e\u0064\u006f\u006d\u0020\u0062\u003a\u0020\u0025\u0020\u0078", _beee)
	_eb.Log.Trace("\u0047\u0065\u006e\u0020\u0049\u0064\u0020\u0030\u003a\u0020\u0025\u0020\u0078", _ead)
	_cag._ffc = _ead
	_cbb := _cag.generateParams(userPass, ownerPass)
	if _cbb != nil {
		return nil, nil, _cbb
	}
	_eeba(&_cag._adfe, _aba)
	if _cag._ffe.V >= 4 {
		if _cace := _cag.saveCryptFilters(_aba); _cace != nil {
			return nil, nil, _cace
		}
	}
	return _cag, &EncryptInfo{Version: _dbbd, Encrypt: _aba, ID0: _ead, ID1: _abg}, nil
}
func (_babb *PdfParser) loadXrefs() (*PdfObjectDictionary, error) {
	_babb._bfba.ObjectMap = make(map[int]XrefObject)
	_babb._fggc = make(objectStreams)
	_cgae, _gebg := _babb._cdea.Seek(0, _bb.SeekEnd)
	if _gebg != nil {
		return nil, _gebg
	}
	_eb.Log.Trace("\u0066s\u0069\u007a\u0065\u003a\u0020\u0025d", _cgae)
	_babb._ggdf = _cgae
	_gebg = _babb.seekToEOFMarker(_cgae)
	if _gebg != nil {
		_eb.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0073\u0065\u0065\u006b\u0020\u0074\u006f\u0020\u0065\u006f\u0066\u0020\u006d\u0061\u0072\u006b\u0065\u0072: \u0025\u0076", _gebg)
		return nil, _gebg
	}
	_bafa, _gebg := _babb._cdea.Seek(0, _bb.SeekCurrent)
	if _gebg != nil {
		return nil, _gebg
	}
	var _aeag int64 = 64
	_ddae := _bafa - _aeag
	if _ddae < 0 {
		_ddae = 0
	}
	_, _gebg = _babb._cdea.Seek(_ddae, _bb.SeekStart)
	if _gebg != nil {
		return nil, _gebg
	}
	_ceda := make([]byte, _aeag)
	_, _gebg = _babb._cdea.Read(_ceda)
	if _gebg != nil {
		_eb.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u003a\u0020\u0025\u0076", _gebg)
		return nil, _gebg
	}
	_defe := _bfff.FindStringSubmatch(string(_ceda))
	if len(_defe) < 2 {
		_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020n\u006f\u0074\u0020f\u006fu\u006e\u0064\u0021")
		return nil, _f.New("\u0073\u0074\u0061\u0072tx\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	if len(_defe) > 2 {
		_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u004du\u006c\u0074\u0069\u0070\u006c\u0065\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029\u0021", _ceda)
		return nil, _f.New("m\u0075\u006c\u0074\u0069\u0070\u006ce\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065f\u0020\u0065\u006et\u0072i\u0065\u0073\u003f")
	}
	_ggba, _ := _d.ParseInt(_defe[1], 10, 64)
	_eb.Log.Trace("\u0073t\u0061r\u0074\u0078\u0072\u0065\u0066\u0020\u0061\u0074\u0020\u0025\u0064", _ggba)
	if _ggba > _cgae {
		_eb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u006f\u0066f\u0073e\u0074 \u006fu\u0074\u0073\u0069\u0064\u0065\u0020\u006f\u0066\u0020\u0066\u0069\u006c\u0065")
		_eb.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0072e\u0070\u0061\u0069\u0072")
		_ggba, _gebg = _babb.repairLocateXref()
		if _gebg != nil {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0065\u0070\u0061\u0069\u0072\u0020\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0066\u0061\u0069\u006c\u0065\u0064 \u0028\u0025\u0073\u0029")
			return nil, _gebg
		}
	}
	_babb._cdea.Seek(_ggba, _bb.SeekStart)
	_babb._dedbc = _ga.NewReader(_babb._cdea)
	_dcacd, _gebg := _babb.parseXref()
	if _gebg != nil {
		return nil, _gebg
	}
	_eece := _dcacd.Get("\u0058R\u0065\u0066\u0053\u0074\u006d")
	if _eece != nil {
		_fgce, _eggdb := _eece.(*PdfObjectInteger)
		if !_eggdb {
			return nil, _f.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0021=\u0020\u0069\u006e\u0074")
		}
		_, _gebg = _babb.parseXrefStream(_fgce)
		if _gebg != nil {
			return nil, _gebg
		}
	}
	var _cbcff []int64
	_dgfgb := func(_fcea int64, _gbadc []int64) bool {
		for _, _dfcdd := range _gbadc {
			if _dfcdd == _fcea {
				return true
			}
		}
		return false
	}
	_eece = _dcacd.Get("\u0050\u0072\u0065\u0076")
	for _eece != nil {
		_gdgg, _aaee := _eece.(*PdfObjectInteger)
		if !_aaee {
			_eb.Log.Debug("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020P\u0072\u0065\u0076\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _eece)
			return _dcacd, nil
		}
		_dccfb := *_gdgg
		_eb.Log.Trace("\u0041\u006eot\u0068\u0065\u0072 \u0050\u0072\u0065\u0076 xr\u0065f \u0074\u0061\u0062\u006c\u0065\u0020\u006fbj\u0065\u0063\u0074\u0020\u0061\u0074\u0020%\u0064", _dccfb)
		_babb._cdea.Seek(int64(_dccfb), _bb.SeekStart)
		_babb._dedbc = _ga.NewReader(_babb._cdea)
		_fcaa, _adbc := _babb.parseXref()
		if _adbc != nil {
			_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0045\u0072\u0072\u006f\u0072\u0020-\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069n\u0067\u0020\u0061\u006e\u006f\u0074\u0068\u0065\u0072\u0020\u0028\u0050re\u0076\u0029\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
			_eb.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0069n\u0075\u0065\u0020\u0062\u0079\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067\u0020\u0069\u0074")
			break
		}
		_babb._bagd = append(_babb._bagd, int64(_dccfb))
		_eece = _fcaa.Get("\u0050\u0072\u0065\u0076")
		if _eece != nil {
			_fdeb := *(_eece.(*PdfObjectInteger))
			if _dgfgb(int64(_fdeb), _cbcff) {
				_eb.Log.Debug("\u0050\u0072ev\u0065\u006e\u0074i\u006e\u0067\u0020\u0063irc\u0075la\u0072\u0020\u0078\u0072\u0065\u0066\u0020re\u0066\u0065\u0072\u0065\u006e\u0063\u0069n\u0067")
				break
			}
			_cbcff = append(_cbcff, int64(_fdeb))
		}
	}
	return _dcacd, nil
}
func _gfac(_cfc XrefTable) {
	_eb.Log.Debug("\u003dX\u003d\u0058\u003d\u0058\u003d")
	_eb.Log.Debug("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_ddb := 0
	for _, _bab := range _cfc.ObjectMap {
		_eb.Log.Debug("i\u002b\u0031\u003a\u0020\u0025\u0064 \u0028\u006f\u0062\u006a\u0020\u006eu\u006d\u003a\u0020\u0025\u0064\u0020\u0067e\u006e\u003a\u0020\u0025\u0064\u0029\u0020\u002d\u003e\u0020%\u0064", _ddb+1, _bab.ObjectNumber, _bab.Generation, _bab.Offset)
		_ddb++
	}
}

type offsetReader struct {
	_bfdd _bb.ReadSeeker
	_cbca int64
}

// XrefObject defines a cross reference entry which is a map between object number (with generation number) and the
// location of the actual object, either as a file offset (xref table entry), or as a location within an xref
// stream object (xref object stream).
type XrefObject struct {
	XType        xrefType
	ObjectNumber int
	Generation   int

	// For normal xrefs (defined by OFFSET)
	Offset int64

	// For xrefs to object streams.
	OsObjNumber int
	OsObjIndex  int
}

// HasInvalidHexRunes implements core.ParserMetadata interface.
func (_aeg ParserMetadata) HasInvalidHexRunes() bool { return _aeg._gca }
func _fdfc(_ggfd uint, _adae, _aedc float64) float64 {
	return (_adae + (float64(_ggfd) * (_aedc - _adae) / 255)) * 255
}

// Set sets the dictionary's key -> val mapping entry. Overwrites if key already set.
func (_ecbfc *PdfObjectDictionary) Set(key PdfObjectName, val PdfObject) {
	_ecbfc.setWithLock(key, val, true)
}

// DecodeBytes decodes a slice of JPX encoded bytes and returns the result.
func (_ffag *JPXEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _ffag.GetFilterName())
	return encoded, ErrNoJPXDecode
}
func (_edde *PdfParser) checkPostEOFData() error {
	const _bfbg = "\u0025\u0025\u0045O\u0046"
	_, _gaad := _edde._cdea.Seek(-int64(len([]byte(_bfbg)))-1, _bb.SeekEnd)
	if _gaad != nil {
		return _gaad
	}
	_dbad := make([]byte, len([]byte(_bfbg))+1)
	_, _gaad = _edde._cdea.Read(_dbad)
	if _gaad != nil {
		if _gaad != _bb.EOF {
			return _gaad
		}
	}
	if string(_dbad) == _bfbg || string(_dbad) == _bfbg+"\u000a" {
		_edde._eggc._fcag = true
	}
	return nil
}

// Update updates multiple keys and returns the dictionary back so can be used in a chained fashion.
func (_eafe *PdfObjectDictionary) Update(objmap map[string]PdfObject) *PdfObjectDictionary {
	_eafe._bfdf.Lock()
	defer _eafe._bfdf.Unlock()
	for _fgcc, _agbg := range objmap {
		_eafe.setWithLock(PdfObjectName(_fgcc), _agbg, false)
	}
	return _eafe
}

type objectStream struct {
	N    int
	_cgg []byte
	_faf map[int]int64
}

// ReadAtLeast reads at least n bytes into slice p.
// Returns the number of bytes read (should always be == n), and an error on failure.
func (_baba *PdfParser) ReadAtLeast(p []byte, n int) (int, error) {
	_afagf := n
	_gbefe := 0
	_ceced := 0
	for _afagf > 0 {
		_aceg, _ebab := _baba._dedbc.Read(p[_gbefe:])
		if _ebab != nil {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _aceg, _ceced, _ebab.Error())
			return _gbefe, _f.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_ceced++
		_gbefe += _aceg
		_afagf -= _aceg
	}
	return _gbefe, nil
}

const (
	_cafb = 0
	_ebeb = 1
	_ggef = 2
	_bae  = 3
	_cecg = 4
)

func (_dggb *PdfParser) xrefNextObjectOffset(_bfced int64) int64 {
	_ceceg := int64(0)
	if len(_dggb._bfba.ObjectMap) == 0 {
		return 0
	}
	if len(_dggb._bfba._bd) == 0 {
		_dbge := 0
		for _, _fadfc := range _dggb._bfba.ObjectMap {
			if _fadfc.Offset > 0 {
				_dbge++
			}
		}
		if _dbge == 0 {
			return 0
		}
		_dggb._bfba._bd = make([]XrefObject, _dbge)
		_gdag := 0
		for _, _gfdf := range _dggb._bfba.ObjectMap {
			if _gfdf.Offset > 0 {
				_dggb._bfba._bd[_gdag] = _gfdf
				_gdag++
			}
		}
		_be.Slice(_dggb._bfba._bd, func(_bagde, _gaeec int) bool { return _dggb._bfba._bd[_bagde].Offset < _dggb._bfba._bd[_gaeec].Offset })
	}
	_ceeb := _be.Search(len(_dggb._bfba._bd), func(_eeebd int) bool { return _dggb._bfba._bd[_eeebd].Offset >= _bfced })
	if _ceeb < len(_dggb._bfba._bd) {
		_ceceg = _dggb._bfba._bd[_ceeb].Offset
	}
	return _ceceg
}

// GetFilterName returns the name of the encoding filter.
func (_bcbb *ASCIIHexEncoder) GetFilterName() string { return StreamEncodingFilterNameASCIIHex }

var _baed = _ba.MustCompile("\u005e\u005b\\\u002b\u002d\u002e\u005d*\u0028\u005b0\u002d\u0039\u002e\u005d\u002b\u0029\u005b\u0065E\u005d\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func _eeba(_adfg *_cbd.StdEncryptDict, _fd *PdfObjectDictionary) {
	_fd.Set("\u0052", MakeInteger(int64(_adfg.R)))
	_fd.Set("\u0050", MakeInteger(int64(_adfg.P)))
	_fd.Set("\u004f", MakeStringFromBytes(_adfg.O))
	_fd.Set("\u0055", MakeStringFromBytes(_adfg.U))
	if _adfg.R >= 5 {
		_fd.Set("\u004f\u0045", MakeStringFromBytes(_adfg.OE))
		_fd.Set("\u0055\u0045", MakeStringFromBytes(_adfg.UE))
		_fd.Set("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", MakeBool(_adfg.EncryptMetadata))
		if _adfg.R > 5 {
			_fd.Set("\u0050\u0065\u0072m\u0073", MakeStringFromBytes(_adfg.Perms))
		}
	}
}

const (
	StreamEncodingFilterNameFlate     = "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameLZW       = "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameDCT       = "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameRunLength = "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameASCIIHex  = "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameASCII85   = "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameCCITTFax  = "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameJBIG2     = "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameJPX       = "\u004aP\u0058\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameRaw       = "\u0052\u0061\u0077"
)

// WriteString outputs the object as it is to be written to file.
func (_dfdg *PdfObjectBool) WriteString() string {
	if *_dfdg {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// HeaderPosition gets the file header position.
func (_gce ParserMetadata) HeaderPosition() int { return _gce._gcc }

var _fgaa = _ba.MustCompile("\u005e\\\u0073\u002a\u005b\u002d]\u002a\u0028\u005c\u0064\u002b)\u005cs\u002b(\u005c\u0064\u002b\u0029\u005c\u0073\u002bR")

// Decrypt attempts to decrypt the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful, false otherwise.
// An error is returned when there is a problem with decrypting.
func (_ecfe *PdfParser) Decrypt(password []byte) (bool, error) {
	if _ecfe._eccc == nil {
		return false, _f.New("\u0063\u0068\u0065\u0063k \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u0072s\u0074")
	}
	_dabgb, _fbgc := _ecfe._eccc.authenticate(password)
	if _fbgc != nil {
		return false, _fbgc
	}
	if !_dabgb {
		_dabgb, _fbgc = _ecfe._eccc.authenticate([]byte(""))
	}
	return _dabgb, _fbgc
}

// MakeArrayFromIntegers64 creates an PdfObjectArray from a slice of int64s, where each array element
// is an PdfObjectInteger.
func MakeArrayFromIntegers64(vals []int64) *PdfObjectArray {
	_baff := MakeArray()
	for _, _gdffg := range vals {
		_baff.Append(MakeInteger(_gdffg))
	}
	return _baff
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_bbcg *JPXEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }

// PdfObjectReference represents the primitive PDF reference object.
type PdfObjectReference struct {
	_dcge            *PdfParser
	ObjectNumber     int64
	GenerationNumber int64
}

// EqualObjects returns true if `obj1` and `obj2` have the same contents.
//
// NOTE: It is a good idea to flatten obj1 and obj2 with FlattenObject before calling this function
// so that contents, rather than references, can be compared.
func EqualObjects(obj1, obj2 PdfObject) bool { return _dcbcc(obj1, obj2, 0) }

// DecodeStream decodes a JPX encoded stream and returns the result as a
// slice of bytes.
func (_baagb *JPXEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _baagb.GetFilterName())
	return streamObj.Stream, ErrNoJPXDecode
}
func (_dggc *PdfParser) traceStreamLength(_cgdgc PdfObject) (PdfObject, error) {
	_bcgb, _abbe := _cgdgc.(*PdfObjectReference)
	if _abbe {
		_ggbbgf, _acbd := _dggc._daad[_bcgb.ObjectNumber]
		if _acbd && _ggbbgf {
			_eb.Log.Debug("\u0053t\u0072\u0065a\u006d\u0020\u004c\u0065n\u0067\u0074\u0068 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 u\u006e\u0072\u0065s\u006f\u006cv\u0065\u0064\u0020\u0028\u0069\u006cl\u0065\u0067a\u006c\u0029")
			return nil, _f.New("\u0069\u006c\u006c\u0065ga\u006c\u0020\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006fo\u0070")
		}
		_dggc._daad[_bcgb.ObjectNumber] = true
	}
	_afd, _geed := _dggc.Resolve(_cgdgc)
	if _geed != nil {
		return nil, _geed
	}
	_eb.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _afd)
	if _abbe {
		_dggc._daad[_bcgb.ObjectNumber] = false
	}
	return _afd, nil
}
func _gbeba(_ceaf _bb.ReadSeeker, _cgac int64) (*limitedReadSeeker, error) {
	_, _aedf := _ceaf.Seek(0, _bb.SeekStart)
	if _aedf != nil {
		return nil, _aedf
	}
	return &limitedReadSeeker{_efag: _ceaf, _befa: _cgac}, nil
}

// DecodeStream decodes a FlateEncoded stream object and give back decoded bytes.
func (_cggc *FlateEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_eb.Log.Trace("\u0046l\u0061t\u0065\u0044\u0065\u0063\u006fd\u0065\u0020s\u0074\u0072\u0065\u0061\u006d")
	_eb.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _cggc.Predictor)
	if _cggc.BitsPerComponent != 8 {
		return nil, _ee.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003d\u0025\u0064\u0020\u0028\u006f\u006e\u006c\u0079\u0020\u0038\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0029", _cggc.BitsPerComponent)
	}
	_gbee, _dad := _cggc.DecodeBytes(streamObj.Stream)
	if _dad != nil {
		return nil, _dad
	}
	_gbee, _dad = _cggc.postDecodePredict(_gbee)
	if _dad != nil {
		return nil, _dad
	}
	return _gbee, nil
}

// MakeInteger creates a PdfObjectInteger from an int64.
func MakeInteger(val int64) *PdfObjectInteger { _abacd := PdfObjectInteger(val); return &_abacd }

var _dffca = []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")

// GetRevision returns PdfParser for the specific version of the Pdf document.
func (_ddbe *PdfParser) GetRevision(revisionNumber int) (*PdfParser, error) {
	_gbec := _ddbe._aegg
	if _gbec == revisionNumber {
		return _ddbe, nil
	}
	if _gbec < revisionNumber {
		return nil, _f.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0072\u0065\u0076\u0069\u0073i\u006fn\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	if _ddbe._gfgdg[revisionNumber] != nil {
		return _ddbe._gfgdg[revisionNumber], nil
	}
	_eade := _ddbe
	for ; _gbec > revisionNumber; _gbec-- {
		_cfgg, _abdg := _eade.GetPreviousRevisionParser()
		if _abdg != nil {
			return nil, _abdg
		}
		_ddbe._gfgdg[_gbec-1] = _cfgg
		_ddbe._bddb[_eade] = _cfgg
		_eade = _cfgg
	}
	return _eade, nil
}

// IsAuthenticated returns true if the PDF has already been authenticated for accessing.
func (_acgd *PdfParser) IsAuthenticated() bool { return _acgd._eccc._agg }

// NewLZWEncoder makes a new LZW encoder with default parameters.
func NewLZWEncoder() *LZWEncoder {
	_dgfa := &LZWEncoder{}
	_dgfa.Predictor = 1
	_dgfa.BitsPerComponent = 8
	_dgfa.Colors = 1
	_dgfa.Columns = 1
	_dgfa.EarlyChange = 1
	return _dgfa
}

// MakeDictMap creates a PdfObjectDictionary initialized from a map of keys to values.
func MakeDictMap(objmap map[string]PdfObject) *PdfObjectDictionary {
	_ecbf := MakeDict()
	return _ecbf.Update(objmap)
}

var _cdfg = _ba.MustCompile("\u005c\u0073\u002a\u0078\u0072\u0065\u0066\u005c\u0073\u002a")

// WriteString outputs the object as it is to be written to file.
func (_dbgbd *PdfObjectNull) WriteString() string { return "\u006e\u0075\u006c\u006c" }

// MakeNull creates an PdfObjectNull.
func MakeNull() *PdfObjectNull { _egaeb := PdfObjectNull{}; return &_egaeb }

// JBIG2CompressionType defines the enum compression type used by the JBIG2Encoder.
type JBIG2CompressionType int

// EncodeBytes encodes the passed in slice of bytes by passing it through the
// EncodeBytes method of the underlying encoders.
func (_fdab *MultiEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_aebc := data
	var _gebc error
	for _feaf := len(_fdab._abeab) - 1; _feaf >= 0; _feaf-- {
		_fadf := _fdab._abeab[_feaf]
		_aebc, _gebc = _fadf.EncodeBytes(_aebc)
		if _gebc != nil {
			return nil, _gebc
		}
	}
	return _aebc, nil
}

const JB2ImageAutoThreshold = -1.0

// Get returns the PdfObject corresponding to the specified key.
// Returns a nil value if the key is not set.
func (_dgac *PdfObjectDictionary) Get(key PdfObjectName) PdfObject {
	_dgac._bfdf.Lock()
	defer _dgac._bfdf.Unlock()
	_cdbec, _ffggd := _dgac._abec[key]
	if !_ffggd {
		return nil
	}
	return _cdbec
}

// GetParser returns the parser for lazy-loading or compare references.
func (_gfce *PdfObjectReference) GetParser() *PdfParser { return _gfce._dcge }

var _bffg = _ba.MustCompile("\u0025P\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func _fcd(_gba *PdfObjectStream, _bbff *PdfObjectDictionary) (*FlateEncoder, error) {
	_dfbb := NewFlateEncoder()
	_cagb := _gba.PdfObjectDictionary
	if _cagb == nil {
		return _dfbb, nil
	}
	_dfbb._gaade = _acgde(_cagb)
	if _bbff == nil {
		_ggbbg := TraceToDirectObject(_cagb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		switch _gagb := _ggbbg.(type) {
		case *PdfObjectArray:
			if _gagb.Len() != 1 {
				_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020\u0028\u0025\u0064\u0029", _gagb.Len())
				return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			if _dff, _fded := GetDict(_gagb.Get(0)); _fded {
				_bbff = _dff
			}
		case *PdfObjectDictionary:
			_bbff = _gagb
		case *PdfObjectNull, nil:
		default:
			_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _ggbbg)
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _bbff == nil {
		return _dfbb, nil
	}
	_eb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _bbff.String())
	_afg := _bbff.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _afg == nil {
		_eb.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_ccec, _gagf := _afg.(*PdfObjectInteger)
		if !_gagf {
			_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _afg)
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_dfbb.Predictor = int(*_ccec)
	}
	_afg = _bbff.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _afg != nil {
		_dedf, _gfb := _afg.(*PdfObjectInteger)
		if !_gfb {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_dfbb.BitsPerComponent = int(*_dedf)
	}
	if _dfbb.Predictor > 1 {
		_dfbb.Columns = 1
		_afg = _bbff.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _afg != nil {
			_eebc, _ceede := _afg.(*PdfObjectInteger)
			if !_ceede {
				return nil, _ee.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_dfbb.Columns = int(*_eebc)
		}
		_dfbb.Colors = 1
		_afg = _bbff.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _afg != nil {
			_egeb, _cdg := _afg.(*PdfObjectInteger)
			if !_cdg {
				return nil, _ee.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_dfbb.Colors = int(*_egeb)
		}
	}
	return _dfbb, nil
}
func (_eecde *PdfParser) resolveReference(_babc *PdfObjectReference) (PdfObject, bool, error) {
	_bfbgc, _ecedd := _eecde.ObjCache[int(_babc.ObjectNumber)]
	if _ecedd {
		return _bfbgc, true, nil
	}
	_abdc, _babbg := _eecde.LookupByReference(*_babc)
	if _babbg != nil {
		return nil, false, _babbg
	}
	_eecde.ObjCache[int(_babc.ObjectNumber)] = _abdc
	return _abdc, false, nil
}

// Encrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_ddag *PdfCrypt) Encrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _ddag.isEncrypted(obj) {
		return nil
	}
	switch _geg := obj.(type) {
	case *PdfIndirectObject:
		_ddag._edc[_geg] = true
		_eb.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _geg.ObjectNumber, _geg.GenerationNumber)
		_dcgd := _geg.ObjectNumber
		_ddd := _geg.GenerationNumber
		_ggbb := _ddag.Encrypt(_geg.PdfObject, _dcgd, _ddd)
		if _ggbb != nil {
			return _ggbb
		}
		return nil
	case *PdfObjectStream:
		_ddag._edc[_geg] = true
		_adaf := _geg.PdfObjectDictionary
		if _efdd, _ecf := _adaf.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _ecf && *_efdd == "\u0058\u0052\u0065\u0066" {
			return nil
		}
		_fdbb := _geg.ObjectNumber
		_daec := _geg.GenerationNumber
		_eb.Log.Trace("\u0045n\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _fdbb, _daec)
		_afc := _dedd
		if _ddag._ffe.V >= 4 {
			_afc = _ddag._badb
			_eb.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _ddag._badb)
			if _fde, _fea := _adaf.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _fea {
				if _befe, _dccc := GetName(_fde.Get(0)); _dccc {
					if *_befe == "\u0043\u0072\u0079p\u0074" {
						_afc = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _bdbd, _eff := _adaf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _eff {
							if _beba, _cgga := _bdbd.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _cgga {
								if _, _bgd := _ddag._ace[string(*_beba)]; _bgd {
									_eb.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_beba)
									_afc = string(*_beba)
								}
							}
						}
					}
				}
			}
			_eb.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _afc)
			if _afc == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_debe := _ddag.Encrypt(_geg.PdfObjectDictionary, _fdbb, _daec)
		if _debe != nil {
			return _debe
		}
		_abac, _debe := _ddag.makeKey(_afc, uint32(_fdbb), uint32(_daec), _ddag._cbcg)
		if _debe != nil {
			return _debe
		}
		_geg.Stream, _debe = _ddag.encryptBytes(_geg.Stream, _afc, _abac)
		if _debe != nil {
			return _debe
		}
		_adaf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_geg.Stream))))
		return nil
	case *PdfObjectString:
		_eb.Log.Trace("\u0045n\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_dafg := _dedd
		if _ddag._ffe.V >= 4 {
			_eb.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _ddag._df)
			if _ddag._df == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_dafg = _ddag._df
		}
		_fdee, _abbc := _ddag.makeKey(_dafg, uint32(parentObjNum), uint32(parentGenNum), _ddag._cbcg)
		if _abbc != nil {
			return _abbc
		}
		_cdae := _geg.Str()
		_ecga := make([]byte, len(_cdae))
		for _gffe := 0; _gffe < len(_cdae); _gffe++ {
			_ecga[_gffe] = _cdae[_gffe]
		}
		_eb.Log.Trace("\u0045n\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _ecga, _ecga)
		_ecga, _abbc = _ddag.encryptBytes(_ecga, _dafg, _fdee)
		if _abbc != nil {
			return _abbc
		}
		_geg._cbdg = string(_ecga)
		return nil
	case *PdfObjectArray:
		for _, _aefg := range _geg.Elements() {
			_dggg := _ddag.Encrypt(_aefg, parentObjNum, parentGenNum)
			if _dggg != nil {
				return _dggg
			}
		}
		return nil
	case *PdfObjectDictionary:
		_gef := false
		if _febc := _geg.Get("\u0054\u0079\u0070\u0065"); _febc != nil {
			_affa, _ffd := _febc.(*PdfObjectName)
			if _ffd && *_affa == "\u0053\u0069\u0067" {
				_gef = true
			}
		}
		for _, _dag := range _geg.Keys() {
			_dbgb := _geg.Get(_dag)
			if _gef && string(_dag) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_dag) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_dag) != "\u0050\u0072\u0065\u0076" && string(_dag) != "\u004c\u0061\u0073\u0074" {
				_gge := _ddag.Encrypt(_dbgb, parentObjNum, parentGenNum)
				if _gge != nil {
					return _gge
				}
			}
		}
		return nil
	}
	return nil
}

// ParseIndirectObject parses an indirect object from the input stream. Can also be an object stream.
// Returns the indirect object (*PdfIndirectObject) or the stream object (*PdfObjectStream).
func (_agfd *PdfParser) ParseIndirectObject() (PdfObject, error) {
	_dcgdf := PdfIndirectObject{}
	_dcgdf._dcge = _agfd
	_eb.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_aag, _egdg := _agfd._dedbc.Peek(20)
	if _egdg != nil {
		if _egdg != _bb.EOF {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return &_dcgdf, _egdg
		}
	}
	_eb.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_aag))
	_fecf := _ffdaa.FindStringSubmatchIndex(string(_aag))
	if len(_fecf) < 6 {
		if _egdg == _bb.EOF {
			return nil, _egdg
		}
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_aag))
		return &_dcgdf, _f.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_agfd._dedbc.Discard(_fecf[0])
	_eb.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _fecf)
	_eddef := _fecf[1] - _fecf[0]
	_agcea := make([]byte, _eddef)
	_, _egdg = _agfd.ReadAtLeast(_agcea, _eddef)
	if _egdg != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _egdg)
		return nil, _egdg
	}
	_eb.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _agcea)
	_gdab := _ffdaa.FindStringSubmatch(string(_agcea))
	if len(_gdab) < 3 {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_agcea))
		return &_dcgdf, _f.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_gcaa, _ := _d.Atoi(_gdab[1])
	_geab, _ := _d.Atoi(_gdab[2])
	_dcgdf.ObjectNumber = int64(_gcaa)
	_dcgdf.GenerationNumber = int64(_geab)
	for {
		_begee, _gcbf := _agfd._dedbc.Peek(2)
		if _gcbf != nil {
			return &_dcgdf, _gcbf
		}
		_eb.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_begee), string(_begee))
		if IsWhiteSpace(_begee[0]) {
			_agfd.skipSpaces()
		} else if _begee[0] == '%' {
			_agfd.skipComments()
		} else if (_begee[0] == '<') && (_begee[1] == '<') {
			_eb.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_dcgdf.PdfObject, _gcbf = _agfd.ParseDict()
			_eb.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _gcbf)
			if _gcbf != nil {
				return &_dcgdf, _gcbf
			}
			_eb.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_begee[0] == '/') || (_begee[0] == '(') || (_begee[0] == '[') || (_begee[0] == '<') {
			_dcgdf.PdfObject, _gcbf = _agfd.parseObject()
			if _gcbf != nil {
				return &_dcgdf, _gcbf
			}
			_eb.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else if _begee[0] == ']' {
			_eb.Log.Debug("\u0057\u0041\u0052\u004e\u0049N\u0047\u003a\u0020\u0027\u005d\u0027 \u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u006eo\u0074\u0020\u0062\u0065i\u006e\u0067\u0020\u0075\u0073\u0065d\u0020\u0061\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006e\u0064\u0069n\u0067\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			_agfd._dedbc.Discard(1)
		} else {
			if _begee[0] == 'e' {
				_acdc, _ffef := _agfd.readTextLine()
				if _ffef != nil {
					return nil, _ffef
				}
				if len(_acdc) >= 6 && _acdc[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _begee[0] == 's' {
				_begee, _ = _agfd._dedbc.Peek(10)
				if string(_begee[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_fgfg := 6
					if len(_begee) > 6 {
						if IsWhiteSpace(_begee[_fgfg]) && _begee[_fgfg] != '\r' && _begee[_fgfg] != '\n' {
							_eb.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0050\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_agfd._eggc._caf = true
							_fgfg++
						}
						if _begee[_fgfg] == '\r' {
							_fgfg++
							if _begee[_fgfg] == '\n' {
								_fgfg++
							}
						} else if _begee[_fgfg] == '\n' {
							_fgfg++
						} else {
							_agfd._eggc._caf = true
						}
					}
					_agfd._dedbc.Discard(_fgfg)
					_dcbc, _defc := _dcgdf.PdfObject.(*PdfObjectDictionary)
					if !_defc {
						return nil, _f.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_eb.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _dcbc)
					_eega, _gbgc := _agfd.traceStreamLength(_dcbc.Get("\u004c\u0065\u006e\u0067\u0074\u0068"))
					if _gbgc != nil {
						_eb.Log.Debug("\u0046\u0061\u0069l\u0020\u0074\u006f\u0020t\u0072\u0061\u0063\u0065\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076", _gbgc)
						return nil, _gbgc
					}
					_eb.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _eega)
					_ffgg, _cebcc := _eega.(*PdfObjectInteger)
					if !_cebcc {
						return nil, _f.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_ebaa := *_ffgg
					if _ebaa < 0 {
						return nil, _f.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					_debbe := _agfd.GetFileOffset()
					_fgeg := _agfd.xrefNextObjectOffset(_debbe)
					if _debbe+int64(_ebaa) > _fgeg && _fgeg > _debbe {
						_eb.Log.Debug("E\u0078\u0070\u0065\u0063te\u0064 \u0065\u006e\u0064\u0069\u006eg\u0020\u0061\u0074\u0020\u0025\u0064", _debbe+int64(_ebaa))
						_eb.Log.Debug("\u004e\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020\u0061t\u0020\u0025\u0064", _fgeg)
						_eegcc := _fgeg - _debbe - 17
						if _eegcc < 0 {
							return nil, _f.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020go\u0069\u006e\u0067 \u0070\u0061s\u0074\u0020\u0062\u006f\u0075\u006ed\u0061\u0072i\u0065\u0073")
						}
						_eb.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u0020c\u006f\u0072\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020%\u0064\u002e\u002e\u002e", _eegcc)
						_ebaa = PdfObjectInteger(_eegcc)
						_dcbc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(_eegcc))
					}
					if int64(_ebaa) > _agfd._ggdf {
						_eb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _f.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_dadf := make([]byte, _ebaa)
					_, _gbgc = _agfd.ReadAtLeast(_dadf, int(_ebaa))
					if _gbgc != nil {
						_eb.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_dadf), _dadf)
						_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gbgc)
						return nil, _gbgc
					}
					_dgabc := PdfObjectStream{}
					_dgabc.Stream = _dadf
					_dgabc.PdfObjectDictionary = _dcgdf.PdfObject.(*PdfObjectDictionary)
					_dgabc.ObjectNumber = _dcgdf.ObjectNumber
					_dgabc.GenerationNumber = _dcgdf.GenerationNumber
					_dgabc.PdfObjectReference._dcge = _agfd
					_agfd.skipSpaces()
					_agfd._dedbc.Discard(9)
					_agfd.skipSpaces()
					return &_dgabc, nil
				}
			}
			_dcgdf.PdfObject, _gcbf = _agfd.parseObject()
			if _dcgdf.PdfObject == nil {
				_eb.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
				_dcgdf.PdfObject = MakeNull()
			}
			return &_dcgdf, _gcbf
		}
	}
	if _dcgdf.PdfObject == nil {
		_eb.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
		_dcgdf.PdfObject = MakeNull()
	}
	_eb.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_dcgdf, nil
}

// Elements returns a slice of the PdfObject elements in the array.
func (_eeedf *PdfObjectArray) Elements() []PdfObject {
	if _eeedf == nil {
		return nil
	}
	return _eeedf._aagc
}

// Encode encodes previously prepare jbig2 document and stores it as the byte slice.
func (_fcdg *JBIG2Encoder) Encode() (_fcad []byte, _gbb error) {
	const _eafa = "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	if _fcdg._eabf == nil {
		return nil, _eca.Errorf(_eafa, "\u0064\u006f\u0063u\u006d\u0065\u006e\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0064a\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_fcdg._eabf.FullHeaders = _fcdg.DefaultPageSettings.FileMode
	_fcad, _gbb = _fcdg._eabf.Encode()
	if _gbb != nil {
		return nil, _eca.Wrap(_gbb, _eafa, "")
	}
	return _fcad, nil
}

// SetIfNotNil sets the dictionary's key -> val mapping entry -IF- val is not nil.
// Note that we take care to perform a type switch.  Otherwise if we would supply a nil value
// of another type, e.g. (PdfObjectArray*)(nil), then it would not be a PdfObject(nil) and thus
// would get set.
func (_cdcce *PdfObjectDictionary) SetIfNotNil(key PdfObjectName, val PdfObject) {
	if val != nil {
		switch _gaeag := val.(type) {
		case *PdfObjectName:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectDictionary:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectStream:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectString:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectNull:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectInteger:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectArray:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectBool:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectFloat:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfObjectReference:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		case *PdfIndirectObject:
			if _gaeag != nil {
				_cdcce.Set(key, val)
			}
		default:
			_eb.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", val)
		}
	}
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on the current encoder settings.
func (_cefc *JBIG2Encoder) MakeDecodeParams() PdfObject { return MakeDict() }

// NewEncoderFromStream creates a StreamEncoder based on the stream's dictionary.
func NewEncoderFromStream(streamObj *PdfObjectStream) (StreamEncoder, error) {
	_ecbd := TraceToDirectObject(streamObj.PdfObjectDictionary.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if _ecbd == nil {
		return NewRawEncoder(), nil
	}
	if _, _fceed := _ecbd.(*PdfObjectNull); _fceed {
		return NewRawEncoder(), nil
	}
	_aefba, _bcea := _ecbd.(*PdfObjectName)
	if !_bcea {
		_aaabf, _fgabe := _ecbd.(*PdfObjectArray)
		if !_fgabe {
			return nil, _ee.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _aaabf.Len() == 0 {
			return NewRawEncoder(), nil
		}
		if _aaabf.Len() != 1 {
			_geada, _cadee := _fefgg(streamObj)
			if _cadee != nil {
				_eb.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _cadee)
				return nil, _cadee
			}
			_eb.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _geada)
			return _geada, nil
		}
		_ecbd = _aaabf.Get(0)
		_aefba, _fgabe = _ecbd.(*PdfObjectName)
		if !_fgabe {
			return nil, _ee.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	if _aggcg, _adfed := _eedgb.Load(_aefba.String()); _adfed {
		return _aggcg.(StreamEncoder), nil
	}
	switch *_aefba {
	case StreamEncodingFilterNameFlate:
		return _fcd(streamObj, nil)
	case StreamEncodingFilterNameLZW:
		return _bagf(streamObj, nil)
	case StreamEncodingFilterNameDCT:
		return _efbb(streamObj, nil)
	case StreamEncodingFilterNameRunLength:
		return _cbce(streamObj, nil)
	case StreamEncodingFilterNameASCIIHex:
		return NewASCIIHexEncoder(), nil
	case StreamEncodingFilterNameASCII85, "\u0041\u0038\u0035":
		return NewASCII85Encoder(), nil
	case StreamEncodingFilterNameCCITTFax:
		return _cgbba(streamObj, nil)
	case StreamEncodingFilterNameJBIG2:
		return _bdfa(streamObj, nil)
	case StreamEncodingFilterNameJPX:
		return NewJPXEncoder(), nil
	}
	_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006fd\u0021")
	return nil, _ee.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064i\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0028\u0025\u0073\u0029", *_aefba)
}

// NewCCITTFaxEncoder makes a new CCITTFax encoder.
func NewCCITTFaxEncoder() *CCITTFaxEncoder { return &CCITTFaxEncoder{Columns: 1728, EndOfBlock: true} }

var _gdgf = []PdfObjectName{"\u0056", "\u0052", "\u004f", "\u0055", "\u0050"}

// PdfObject is an interface which all primitive PDF objects must implement.
type PdfObject interface {

	// String outputs a string representation of the primitive (for debugging).
	String() string

	// WriteString outputs the PDF primitive as written to file as expected by the standard.
	// TODO(dennwc): it should return a byte slice, or accept a writer
	WriteString() string
}

// Len returns the number of elements in the array.
func (_eadad *PdfObjectArray) Len() int {
	if _eadad == nil {
		return 0
	}
	return len(_eadad._aagc)
}

// MakeString creates an PdfObjectString from a string.
// NOTE: PDF does not use utf-8 string encoding like Go so `s` will often not be a utf-8 encoded
// string.
func MakeString(s string) *PdfObjectString { _ddggc := PdfObjectString{_cbdg: s}; return &_ddggc }
func _bfcd(_cacd PdfObject, _gefc int) PdfObject {
	if _gefc > _gegcg {
		_eb.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _gegcg)
		return MakeNull()
	}
	switch _gfbd := _cacd.(type) {
	case *PdfIndirectObject:
		_cacd = _bfcd((*_gfbd).PdfObject, _gefc+1)
	case *PdfObjectArray:
		for _gdaa, _dceeb := range (*_gfbd)._aagc {
			(*_gfbd)._aagc[_gdaa] = _bfcd(_dceeb, _gefc+1)
		}
	case *PdfObjectDictionary:
		for _debg, _agca := range (*_gfbd)._abec {
			(*_gfbd)._abec[_debg] = _bfcd(_agca, _gefc+1)
		}
		_be.Slice((*_gfbd)._begbb, func(_dace, _ebge int) bool { return (*_gfbd)._begbb[_dace] < (*_gfbd)._begbb[_ebge] })
	}
	return _cacd
}
func _cgbba(_dfdf *PdfObjectStream, _cgge *PdfObjectDictionary) (*CCITTFaxEncoder, error) {
	_bcag := NewCCITTFaxEncoder()
	_ecbc := _dfdf.PdfObjectDictionary
	if _ecbc == nil {
		return _bcag, nil
	}
	if _cgge == nil {
		_cabf := TraceToDirectObject(_ecbc.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _cabf != nil {
			switch _beec := _cabf.(type) {
			case *PdfObjectDictionary:
				_cgge = _beec
			case *PdfObjectArray:
				if _beec.Len() == 1 {
					if _fefe, _ecca := GetDict(_beec.Get(0)); _ecca {
						_cgge = _fefe
					}
				}
			default:
				_eb.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _cabf)
				return nil, _f.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
		if _cgge == nil {
			_eb.Log.Error("\u0044\u0065c\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064 %\u0023\u0076", _cabf)
			return nil, _f.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _bgade, _begaf := GetNumberAsInt64(_cgge.Get("\u004b")); _begaf == nil {
		_bcag.K = int(_bgade)
	}
	if _aabb, _facd := GetNumberAsInt64(_cgge.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _facd == nil {
		_bcag.Columns = int(_aabb)
	} else {
		_bcag.Columns = 1728
	}
	if _cabc, _cfae := GetNumberAsInt64(_cgge.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _cfae == nil {
		_bcag.BlackIs1 = _cabc > 0
	} else {
		if _cdgd, _aade := GetBoolVal(_cgge.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _aade {
			_bcag.BlackIs1 = _cdgd
		} else {
			if _ccdf, _gfg := GetArray(_cgge.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _gfg {
				_cefd, _cagdb := _ccdf.ToIntegerArray()
				if _cagdb == nil {
					_bcag.BlackIs1 = _cefd[0] == 1 && _cefd[1] == 0
				}
			}
		}
	}
	if _aecg, _afec := GetNumberAsInt64(_cgge.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _afec == nil {
		_bcag.EncodedByteAlign = _aecg > 0
	} else {
		if _fcbg, _ccab := GetBoolVal(_cgge.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _ccab {
			_bcag.EncodedByteAlign = _fcbg
		}
	}
	if _aegd, _fbef := GetNumberAsInt64(_cgge.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _fbef == nil {
		_bcag.EndOfLine = _aegd > 0
	} else {
		if _eeab, _gdce := GetBoolVal(_cgge.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _gdce {
			_bcag.EndOfLine = _eeab
		}
	}
	if _fgea, _agb := GetNumberAsInt64(_cgge.Get("\u0052\u006f\u0077\u0073")); _agb == nil {
		_bcag.Rows = int(_fgea)
	}
	_bcag.EndOfBlock = true
	if _aecc, _fcagb := GetNumberAsInt64(_cgge.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _fcagb == nil {
		_bcag.EndOfBlock = _aecc > 0
	} else {
		if _fgbd, _dgab := GetBoolVal(_cgge.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _dgab {
			_bcag.EndOfBlock = _fgbd
		}
	}
	if _fgd, _gefd := GetNumberAsInt64(_cgge.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _gefd != nil {
		_bcag.DamagedRowsBeforeError = int(_fgd)
	}
	_eb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _cgge.String())
	return _bcag, nil
}

// GetFilterName returns the name of the encoding filter.
func (_bec *CCITTFaxEncoder) GetFilterName() string { return StreamEncodingFilterNameCCITTFax }

// JBIG2Encoder implements both jbig2 encoder and the decoder. The encoder allows to encode
// provided images (best used document scans) in multiple way. By default it uses single page generic
// encoder. It allows to store lossless data as a single segment.
// In order to store multiple image pages use the 'FileMode' which allows to store more pages within single jbig2 document.
// WIP: In order to obtain better compression results the encoder would allow to encode the input in a
// lossy or lossless way with a component (symbol) mode. It divides the image into components.
// Then checks if any component is 'similar' to the others and maps them together. The symbol classes are stored
// in the dictionary. Then the encoder creates text regions which uses the related symbol classes to fill it's space.
// The similarity is defined by the 'Threshold' variable (default: 0.95). The less the value is, the more components
// matches to single class, thus the compression is better, but the result might become lossy.
type JBIG2Encoder struct {

	// These values are required to be set for the 'EncodeBytes' method.
	// ColorComponents defines the number of color components for provided image.
	ColorComponents int

	// BitsPerComponent is the number of bits that stores per color component
	BitsPerComponent int

	// Width is the width of the image to encode
	Width int

	// Height is the height of the image to encode.
	Height int
	_eabf  *_ce.Document

	// Globals are the JBIG2 global segments.
	Globals _gd.Globals

	// IsChocolateData defines if the data is encoded such that
	// binary data '1' means black and '0' white.
	// otherwise the data is called vanilla.
	// Naming convention taken from: 'https://en.wikipedia.org/wiki/Binary_image#Interpretation'
	IsChocolateData bool

	// DefaultPageSettings are the settings parameters used by the jbig2 encoder.
	DefaultPageSettings JBIG2EncoderSettings
}

// NewRunLengthEncoder makes a new run length encoder
func NewRunLengthEncoder() *RunLengthEncoder { return &RunLengthEncoder{} }

// String returns the state of the bool as "true" or "false".
func (_deg *PdfObjectBool) String() string {
	if *_deg {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}
func (_gcbg *PdfParser) checkLinearizedInformation(_caace *PdfObjectDictionary) (bool, error) {
	var _fbed error
	_gcbg._ggfc, _fbed = GetNumberAsInt64(_caace.Get("\u004c"))
	if _fbed != nil {
		return false, _fbed
	}
	_fbed = _gcbg.seekToEOFMarker(_gcbg._ggfc)
	switch _fbed {
	case nil:
		return true, nil
	case _bgbf:
		return false, nil
	default:
		return false, _fbed
	}
}

// DecodeStream decodes a LZW encoded stream and returns the result as a
// slice of bytes.
func (_abef *LZWEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_eb.Log.Trace("\u004c\u005a\u0057 \u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067")
	_eb.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _abef.Predictor)
	_cgd, _agce := _abef.DecodeBytes(streamObj.Stream)
	if _agce != nil {
		return nil, _agce
	}
	_eb.Log.Trace("\u0020\u0049\u004e\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(streamObj.Stream), streamObj.Stream)
	_eb.Log.Trace("\u004f\u0055\u0054\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(_cgd), _cgd)
	if _abef.Predictor > 1 {
		if _abef.Predictor == 2 {
			_eb.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_beaa := _abef.Columns * _abef.Colors
			if _beaa < 1 {
				return []byte{}, nil
			}
			_gbfg := len(_cgd) / _beaa
			if len(_cgd)%_beaa != 0 {
				_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_cgd), _beaa)
			}
			if _beaa%_abef.Colors != 0 {
				return nil, _ee.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _beaa, _abef.Colors)
			}
			if _beaa > len(_cgd) {
				_eb.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _beaa, len(_cgd))
				return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_eb.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_cgd), _cgd)
			_aee := _bg.NewBuffer(nil)
			for _acf := 0; _acf < _gbfg; _acf++ {
				_ffbd := _cgd[_beaa*_acf : _beaa*(_acf+1)]
				for _gceb := _abef.Colors; _gceb < _beaa; _gceb++ {
					_ffbd[_gceb] = byte(int(_ffbd[_gceb]+_ffbd[_gceb-_abef.Colors]) % 256)
				}
				_aee.Write(_ffbd)
			}
			_cdd := _aee.Bytes()
			_eb.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_cdd), _cdd)
			return _cdd, nil
		} else if _abef.Predictor >= 10 && _abef.Predictor <= 15 {
			_eb.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_ddda := _abef.Columns*_abef.Colors + 1
			if _ddda < 1 {
				return []byte{}, nil
			}
			_afad := len(_cgd) / _ddda
			if len(_cgd)%_ddda != 0 {
				return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_cgd), _ddda)
			}
			if _ddda > len(_cgd) {
				_eb.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _ddda, len(_cgd))
				return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_eeaa := _bg.NewBuffer(nil)
			_eb.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _abef.Columns)
			_eb.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_cgd), _ddda, _afad)
			_dcee := make([]byte, _ddda)
			for _fbg := 0; _fbg < _ddda; _fbg++ {
				_dcee[_fbg] = 0
			}
			for _gabb := 0; _gabb < _afad; _gabb++ {
				_dcfc := _cgd[_ddda*_gabb : _ddda*(_gabb+1)]
				_ced := _dcfc[0]
				switch _ced {
				case 0:
				case 1:
					for _cfg := 2; _cfg < _ddda; _cfg++ {
						_dcfc[_cfg] = byte(int(_dcfc[_cfg]+_dcfc[_cfg-1]) % 256)
					}
				case 2:
					for _dfe := 1; _dfe < _ddda; _dfe++ {
						_dcfc[_dfe] = byte(int(_dcfc[_dfe]+_dcee[_dfe]) % 256)
					}
				default:
					_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0066i\u006c\u0074\u0065\u0072\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _ced)
					return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _ced)
				}
				for _fegg := 0; _fegg < _ddda; _fegg++ {
					_dcee[_fegg] = _dcfc[_fegg]
				}
				_eeaa.Write(_dcfc[1:])
			}
			_cfce := _eeaa.Bytes()
			return _cfce, nil
		} else {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _abef.Predictor)
			return nil, _ee.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _abef.Predictor)
		}
	}
	return _cgd, nil
}
func (_geeg *PdfParser) skipComments() error {
	if _, _cdbb := _geeg.skipSpaces(); _cdbb != nil {
		return _cdbb
	}
	_fbcb := true
	for {
		_fab, _fdcf := _geeg._dedbc.Peek(1)
		if _fdcf != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _fdcf.Error())
			return _fdcf
		}
		if _fbcb && _fab[0] != '%' {
			return nil
		}
		_fbcb = false
		if (_fab[0] != '\r') && (_fab[0] != '\n') {
			_geeg._dedbc.ReadByte()
		} else {
			break
		}
	}
	return _geeg.skipComments()
}

// UpdateParams updates the parameter values of the encoder.
func (_bfeeg *MultiEncoder) UpdateParams(params *PdfObjectDictionary) {
	for _, _ffaf := range _bfeeg._abeab {
		_ffaf.UpdateParams(params)
	}
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_dac *ASCIIHexEncoder) MakeDecodeParams() PdfObject { return nil }
func (_dgcf *ASCII85Encoder) base256Tobase85(_eegd uint32) [5]byte {
	_fcfe := [5]byte{0, 0, 0, 0, 0}
	_ddcfg := _eegd
	for _abea := 0; _abea < 5; _abea++ {
		_cfca := uint32(1)
		for _fdcd := 0; _fdcd < 4-_abea; _fdcd++ {
			_cfca *= 85
		}
		_cdbe := _ddcfg / _cfca
		_ddcfg = _ddcfg % _cfca
		_fcfe[_abea] = byte(_cdbe)
	}
	return _fcfe
}

// HasInvalidSubsectionHeader implements core.ParserMetadata interface.
func (_gbfd ParserMetadata) HasInvalidSubsectionHeader() bool { return _gbfd._bgf }

// StreamEncoder represents the interface for all PDF stream encoders.
type StreamEncoder interface {
	GetFilterName() string
	MakeDecodeParams() PdfObject
	MakeStreamDict() *PdfObjectDictionary
	UpdateParams(_aggb *PdfObjectDictionary)
	EncodeBytes(_dgf []byte) ([]byte, error)
	DecodeBytes(_ecfc []byte) ([]byte, error)
	DecodeStream(_daeb *PdfObjectStream) ([]byte, error)
}

func (_bgbdf *PdfParser) parseBool() (PdfObjectBool, error) {
	_cded, _dcacf := _bgbdf._dedbc.Peek(4)
	if _dcacf != nil {
		return PdfObjectBool(false), _dcacf
	}
	if (len(_cded) >= 4) && (string(_cded[:4]) == "\u0074\u0072\u0075\u0065") {
		_bgbdf._dedbc.Discard(4)
		return PdfObjectBool(true), nil
	}
	_cded, _dcacf = _bgbdf._dedbc.Peek(5)
	if _dcacf != nil {
		return PdfObjectBool(false), _dcacf
	}
	if (len(_cded) >= 5) && (string(_cded[:5]) == "\u0066\u0061\u006cs\u0065") {
		_bgbdf._dedbc.Discard(5)
		return PdfObjectBool(false), nil
	}
	return PdfObjectBool(false), _f.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// Version represents a version of a PDF standard.
type Version struct {
	Major int
	Minor int
}

func (_bcfg *PdfParser) readTextLine() (string, error) {
	var _adfbf _bg.Buffer
	for {
		_abed, _ceedd := _bcfg._dedbc.Peek(1)
		if _ceedd != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ceedd.Error())
			return _adfbf.String(), _ceedd
		}
		if (_abed[0] != '\r') && (_abed[0] != '\n') {
			_bcad, _ := _bcfg._dedbc.ReadByte()
			_adfbf.WriteByte(_bcad)
		} else {
			break
		}
	}
	return _adfbf.String(), nil
}

// IsFloatDigit checks if a character can be a part of a float number string.
func IsFloatDigit(c byte) bool { return ('0' <= c && c <= '9') || c == '.' }

// HasOddLengthHexStrings checks if the document has odd length hexadecimal strings.
func (_ebb ParserMetadata) HasOddLengthHexStrings() bool { return _ebb._cdff }

var _eedgb _g.Map

// GetAccessPermissions returns the PDF access permissions as an AccessPermissions object.
func (_eecc *PdfCrypt) GetAccessPermissions() _cbd.Permissions { return _eecc._adfe.P }

// GetNumbersAsFloat converts a list of pdf objects representing floats or integers to a slice of
// float64 values.
func GetNumbersAsFloat(objects []PdfObject) (_gfde []float64, _gcbac error) {
	for _, _efff := range objects {
		_cdfdc, _bgagf := GetNumberAsFloat(_efff)
		if _bgagf != nil {
			return nil, _bgagf
		}
		_gfde = append(_gfde, _cdfdc)
	}
	return _gfde, nil
}

// WriteString outputs the object as it is to be written to file.
func (_gec *PdfObjectReference) WriteString() string {
	var _ceaa _cc.Builder
	_ceaa.WriteString(_d.FormatInt(_gec.ObjectNumber, 10))
	_ceaa.WriteString("\u0020")
	_ceaa.WriteString(_d.FormatInt(_gec.GenerationNumber, 10))
	_ceaa.WriteString("\u0020\u0052")
	return _ceaa.String()
}

// ToInt64Slice returns a slice of all array elements as an int64 slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_cffa *PdfObjectArray) ToInt64Slice() ([]int64, error) {
	var _fecb []int64
	for _, _eebd := range _cffa.Elements() {
		if _dgfaeb, _ddeb := _eebd.(*PdfObjectInteger); _ddeb {
			_fecb = append(_fecb, int64(*_dgfaeb))
		} else {
			return nil, ErrTypeError
		}
	}
	return _fecb, nil
}

// DecodeStream returns the passed in stream as a slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_gaac *RawEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return streamObj.Stream, nil
}

// MakeEncodedString creates a PdfObjectString with encoded content, which can be either
// UTF-16BE or PDFDocEncoding depending on whether `utf16BE` is true or false respectively.
func MakeEncodedString(s string, utf16BE bool) *PdfObjectString {
	if utf16BE {
		var _egda _bg.Buffer
		_egda.Write([]byte{0xFE, 0xFF})
		_egda.WriteString(_dd.StringToUTF16(s))
		return &PdfObjectString{_cbdg: _egda.String(), _cdca: true}
	}
	return &PdfObjectString{_cbdg: string(_dd.StringToPDFDocEncoding(s)), _cdca: false}
}
func (_aeb *PdfCrypt) isDecrypted(_feb PdfObject) bool {
	_, _bbg := _aeb._ge[_feb]
	if _bbg {
		_eb.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0064\u0065\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	switch _ede := _feb.(type) {
	case *PdfObjectStream:
		if _aeb._adfe.R != 5 {
			if _acea, _ccb := _ede.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _ccb && *_acea == "\u0058\u0052\u0065\u0066" {
				return true
			}
		}
	case *PdfIndirectObject:
		if _, _bbg = _aeb._eda[int(_ede.ObjectNumber)]; _bbg {
			return true
		}
		switch _fed := _ede.PdfObject.(type) {
		case *PdfObjectDictionary:
			_aaae := true
			for _, _bcbg := range _gdgf {
				if _fed.Get(_bcbg) == nil {
					_aaae = false
					break
				}
			}
			if _aaae {
				return true
			}
		}
	}
	_eb.Log.Trace("\u004e\u006f\u0074\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// MakeLazy create temporary file for stream to reduce memory usage.
// It can be used for creating PDF with many images.
// Temporary files are removed automatically when Write/WriteToFile is called for creator object.
func (_cadc *PdfObjectStream) MakeLazy() error {
	if _cadc.Lazy {
		return nil
	}
	_cbcb, _eedb := _c.CreateTemp("", "\u0078o\u0062\u006a\u0065\u0063\u0074")
	if _eedb != nil {
		return _eedb
	}
	defer _cbcb.Close()
	_, _eedb = _cbcb.Write(_cadc.Stream)
	if _eedb != nil {
		return _eedb
	}
	_cadc.Lazy = true
	_cadc.Stream = nil
	_cadc.TempFile = _cbcb.Name()
	return nil
}

// Resolve resolves the reference and returns the indirect or stream object.
// If the reference cannot be resolved, a *PdfObjectNull object is returned.
func (_cdcc *PdfObjectReference) Resolve() PdfObject {
	if _cdcc._dcge == nil {
		return MakeNull()
	}
	_bbaa, _, _bbe := _cdcc._dcge.resolveReference(_cdcc)
	if _bbe != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065\u0073\u006f\u006cv\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065r\u0065n\u0063\u0065\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074", _bbe)
		return MakeNull()
	}
	if _bbaa == nil {
		_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0072\u0065\u0073ol\u0076\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065:\u0020\u006ei\u006c\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0061\u0020nu\u006c\u006c\u0020o\u0062\u006a\u0065\u0063\u0074")
		return MakeNull()
	}
	return _bbaa
}

// GetXrefTable returns the PDFs xref table.
func (_fefgd *PdfParser) GetXrefTable() XrefTable { return _fefgd._bfba }

type cryptFilters map[string]_gbe.Filter

// WriteString outputs the object as it is to be written to file.
func (_gbce *PdfObjectDictionary) WriteString() string {
	var _bfbbf _cc.Builder
	_bfbbf.WriteString("\u003c\u003c")
	for _, _gcaf := range _gbce._begbb {
		_cbag := _gbce._abec[_gcaf]
		_bfbbf.WriteString(_gcaf.WriteString())
		_bfbbf.WriteString("\u0020")
		_bfbbf.WriteString(_cbag.WriteString())
	}
	_bfbbf.WriteString("\u003e\u003e")
	return _bfbbf.String()
}
func (_aagg *PdfObjectFloat) String() string { return _ee.Sprintf("\u0025\u0066", *_aagg) }

// ASCII85Encoder implements ASCII85 encoder/decoder.
type ASCII85Encoder struct{}

// NewCompliancePdfParser creates a new PdfParser that will parse input reader with the focus on extracting more metadata, which
// might affect performance of the regular PdfParser this function.
func NewCompliancePdfParser(rs _bb.ReadSeeker) (_fad *PdfParser, _bfb error) {
	_fad = &PdfParser{_cdea: rs, ObjCache: make(objectCache), _daad: map[int64]bool{}, _aega: true, _bddb: make(map[*PdfParser]*PdfParser)}
	if _bfb = _fad.parseDetailedHeader(); _bfb != nil {
		return nil, _bfb
	}
	if _fad._cbeb, _bfb = _fad.loadXrefs(); _bfb != nil {
		_eb.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _bfb)
		return nil, _bfb
	}
	_eb.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _fad._cbeb)
	if len(_fad._bfba.ObjectMap) == 0 {
		return nil, _ee.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	return _fad, nil
}

// IsNullObject returns true if `obj` is a PdfObjectNull.
func IsNullObject(obj PdfObject) bool {
	_, _bcadcd := TraceToDirectObject(obj).(*PdfObjectNull)
	return _bcadcd
}

// String returns a string representation of `name`.
func (_ecab *PdfObjectName) String() string { return string(*_ecab) }

// AddEncoder adds the passed in encoder to the underlying encoder slice.
func (_agda *MultiEncoder) AddEncoder(encoder StreamEncoder) {
	_agda._abeab = append(_agda._abeab, encoder)
}

// ParserMetadata is the parser based metadata information about document.
// The data here could be used on document verification.
type ParserMetadata struct {
	_gcc  int
	_aced bool
	_affb [4]byte
	_fcag bool
	_cdff bool
	_gca  bool
	_caf  bool
	_bgf  bool
	_bgb  bool
}

// NewParserFromString is used for testing purposes.
func NewParserFromString(txt string) *PdfParser {
	_acbgg := _bg.NewReader([]byte(txt))
	_fegc := &PdfParser{ObjCache: objectCache{}, _cdea: _acbgg, _dedbc: _ga.NewReader(_acbgg), _ggdf: int64(len(txt)), _daad: map[int64]bool{}, _bddb: make(map[*PdfParser]*PdfParser)}
	_fegc._bfba.ObjectMap = make(map[int]XrefObject)
	return _fegc
}

// WriteString outputs the object as it is to be written to file.
func (_efee *PdfObjectInteger) WriteString() string { return _d.FormatInt(int64(*_efee), 10) }

// Append appends PdfObject(s) to the array.
func (_cgeb *PdfObjectArray) Append(objects ...PdfObject) {
	if _cgeb == nil {
		_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0020\u002d\u0020\u0041\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0061\u0070p\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020a\u0072\u0072\u0061\u0079")
		return
	}
	_cgeb._aagc = append(_cgeb._aagc, objects...)
}

// GetPreviousRevisionParser returns PdfParser for the previous version of the Pdf document.
func (_dgfe *PdfParser) GetPreviousRevisionParser() (*PdfParser, error) {
	if _dgfe._aegg == 0 {
		return nil, _f.New("\u0074\u0068\u0069\u0073 i\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0072\u0065\u0076\u0069\u0073\u0069o\u006e")
	}
	if _geda, _gaeef := _dgfe._bddb[_dgfe]; _gaeef {
		return _geda, nil
	}
	_gbbgg, _fgbbg := _dgfe.GetPreviousRevisionReadSeeker()
	if _fgbbg != nil {
		return nil, _fgbbg
	}
	_gada, _fgbbg := NewParser(_gbbgg)
	_gada._bddb = _dgfe._bddb
	if _fgbbg != nil {
		return nil, _fgbbg
	}
	_dgfe._bddb[_dgfe] = _gada
	return _gada, nil
}

// Elements returns a slice of the PdfObject elements in the array.
// Preferred over accessing the array directly as type may be changed in future major versions (v3).
func (_ccefcg *PdfObjectStreams) Elements() []PdfObject {
	if _ccefcg == nil {
		return nil
	}
	return _ccefcg._fagda
}

// Bytes returns the PdfObjectString content as a []byte array.
func (_eaacd *PdfObjectString) Bytes() []byte { return []byte(_eaacd._cbdg) }

// WriteString outputs the object as it is to be written to file.
func (_gbaf *PdfObjectArray) WriteString() string {
	var _defed _cc.Builder
	_defed.WriteString("\u005b")
	for _dagbf, _cfdd := range _gbaf.Elements() {
		_defed.WriteString(_cfdd.WriteString())
		if _dagbf < (_gbaf.Len() - 1) {
			_defed.WriteString("\u0020")
		}
	}
	_defed.WriteString("\u005d")
	return _defed.String()
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_cab *ASCIIHexEncoder) MakeStreamDict() *PdfObjectDictionary {
	_ffdd := MakeDict()
	_ffdd.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_cab.GetFilterName()))
	return _ffdd
}

// HasEOLAfterHeader gets information if there is a EOL after the version header.
func (_eee ParserMetadata) HasEOLAfterHeader() bool { return _eee._aced }

// CCITTFaxEncoder implements Group3 and Group4 facsimile (fax) encoder/decoder.
type CCITTFaxEncoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}

// IsDecimalDigit checks if the character is a part of a decimal number string.
func IsDecimalDigit(c byte) bool { return '0' <= c && c <= '9' }

// GetStream returns the *PdfObjectStream represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetStream(obj PdfObject) (_fddc *PdfObjectStream, _fdfg bool) {
	obj = ResolveReference(obj)
	_fddc, _fdfg = obj.(*PdfObjectStream)
	return _fddc, _fdfg
}

// Len returns the number of elements in the streams.
func (_bfbf *PdfObjectStreams) Len() int {
	if _bfbf == nil {
		return 0
	}
	return len(_bfbf._fagda)
}

const (
	JB2Generic JBIG2CompressionType = iota
	JB2SymbolCorrelation
	JB2SymbolRankHaus
)

// DecodeBytes decodes the CCITTFax encoded image data.
func (_faed *CCITTFaxEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_ggdgf, _edbc := _bad.NewDecoder(encoded, _bad.DecodeOptions{Columns: _faed.Columns, Rows: _faed.Rows, K: _faed.K, EncodedByteAligned: _faed.EncodedByteAlign, BlackIsOne: _faed.BlackIs1, EndOfBlock: _faed.EndOfBlock, EndOfLine: _faed.EndOfLine, DamagedRowsBeforeError: _faed.DamagedRowsBeforeError})
	if _edbc != nil {
		return nil, _edbc
	}
	_gebf, _edbc := _bb.ReadAll(_ggdgf)
	if _edbc != nil {
		return nil, _edbc
	}
	return _gebf, nil
}

// String returns a string describing `null`.
func (_dcad *PdfObjectNull) String() string { return "\u006e\u0075\u006c\u006c" }

// EncodeBytes implements support for LZW encoding.  Currently not supporting predictors (raw compressed data only).
// Only supports the Early change = 1 algorithm (compress/lzw) as the other implementation
// does not have a write method.
// TODO: Consider refactoring compress/lzw to allow both.
func (_geffb *LZWEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _geffb.Predictor != 1 {
		return nil, _ee.Errorf("\u004c\u005aW \u0050\u0072\u0065d\u0069\u0063\u0074\u006fr =\u00201 \u006f\u006e\u006c\u0079\u0020\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0079e\u0074")
	}
	if _geffb.EarlyChange == 1 {
		return nil, _ee.Errorf("\u004c\u005a\u0057\u0020\u0045\u0061\u0072\u006c\u0079\u0020\u0043\u0068\u0061n\u0067\u0065\u0020\u003d\u0020\u0030 \u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0079\u0065\u0074")
	}
	var _bgac _bg.Buffer
	_dgcg := _eg.NewWriter(&_bgac, _eg.MSB, 8)
	_dgcg.Write(data)
	_dgcg.Close()
	return _bgac.Bytes(), nil
}
func _dfca(_bdea PdfObject) (*float64, error) {
	switch _edab := _bdea.(type) {
	case *PdfObjectFloat:
		_dddb := float64(*_edab)
		return &_dddb, nil
	case *PdfObjectInteger:
		_gbfb := float64(*_edab)
		return &_gbfb, nil
	case *PdfObjectNull:
		return nil, nil
	}
	return nil, ErrNotANumber
}

// EncodeJBIG2Image encodes 'img' into jbig2 encoded bytes stream, using default encoder settings.
func (_ecfa *JBIG2Encoder) EncodeJBIG2Image(img *JBIG2Image) ([]byte, error) {
	const _aafc = "c\u006f\u0072\u0065\u002eEn\u0063o\u0064\u0065\u004a\u0042\u0049G\u0032\u0049\u006d\u0061\u0067\u0065"
	if _fbda := _ecfa.AddPageImage(img, &_ecfa.DefaultPageSettings); _fbda != nil {
		return nil, _eca.Wrap(_fbda, _aafc, "")
	}
	return _ecfa.Encode()
}

// GetAsFloat64Slice returns the array as []float64 slice.
// Returns an error if not entirely numeric (only PdfObjectIntegers, PdfObjectFloats).
func (_bafbd *PdfObjectArray) GetAsFloat64Slice() ([]float64, error) {
	var _cfdb []float64
	for _, _cdfeb := range _bafbd.Elements() {
		_cabe, _cgeg := GetNumberAsFloat(TraceToDirectObject(_cdfeb))
		if _cgeg != nil {
			return nil, _ee.Errorf("\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0075m\u0062\u0065\u0072")
		}
		_cfdb = append(_cfdb, _cabe)
	}
	return _cfdb, nil
}
func (_adfc *PdfCrypt) encryptBytes(_abb []byte, _aabg string, _ddg []byte) ([]byte, error) {
	_eb.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_abfa, _gaa := _adfc._ace[_aabg]
	if !_gaa {
		return nil, _ee.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _aabg)
	}
	return _abfa.EncryptBytes(_abb, _ddg)
}

// DrawableImage is same as golang image/draw's Image interface that allow drawing images.
type DrawableImage interface {
	ColorModel() _cb.Model
	Bounds() _a.Rectangle
	At(_beaf, _ecb int) _cb.Color
	Set(_bfbb, _cdgb int, _agdfa _cb.Color)
}

// GetTrailer returns the PDFs trailer dictionary. The trailer dictionary is typically the starting point for a PDF,
// referencing other key objects that are important in the document structure.
func (_efcd *PdfParser) GetTrailer() *PdfObjectDictionary { return _efcd._cbeb }

// Merge merges in key/values from another dictionary. Overwriting if has same keys.
// The mutated dictionary (d) is returned in order to allow method chaining.
func (_febe *PdfObjectDictionary) Merge(another *PdfObjectDictionary) *PdfObjectDictionary {
	if another != nil {
		for _, _afgdc := range another.Keys() {
			_eabe := another.Get(_afgdc)
			_febe.Set(_afgdc, _eabe)
		}
	}
	return _febe
}

// IsEncrypted checks if the document is encrypted. A bool flag is returned indicating the result.
// First time when called, will check if the Encrypt dictionary is accessible through the trailer dictionary.
// If encrypted, prepares a crypt datastructure which can be used to authenticate and decrypt the document.
// On failure, an error is returned.
func (_eagb *PdfParser) IsEncrypted() (bool, error) {
	if _eagb._eccc != nil {
		return true, nil
	} else if _eagb._cbeb == nil {
		return false, nil
	}
	_eb.Log.Trace("\u0043\u0068\u0065c\u006b\u0069\u006e\u0067 \u0065\u006e\u0063\u0072\u0079\u0070\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021")
	_cebb := _eagb._cbeb.Get("\u0045n\u0063\u0072\u0079\u0070\u0074")
	if _cebb == nil {
		return false, nil
	}
	_eb.Log.Trace("\u0049\u0073\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0021")
	var (
		_dbd *PdfObjectDictionary
	)
	switch _cdad := _cebb.(type) {
	case *PdfObjectDictionary:
		_dbd = _cdad
	case *PdfObjectReference:
		_eb.Log.Trace("\u0030\u003a\u0020\u004c\u006f\u006f\u006b\u0020\u0075\u0070\u0020\u0072e\u0066\u0020\u0025\u0071", _cdad)
		_dfcdde, _becg := _eagb.LookupByReference(*_cdad)
		_eb.Log.Trace("\u0031\u003a\u0020%\u0071", _dfcdde)
		if _becg != nil {
			return false, _becg
		}
		_ebbc, _accgg := _dfcdde.(*PdfIndirectObject)
		if !_accgg {
			_eb.Log.Debug("E\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			return false, _f.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ebdd, _accgg := _ebbc.PdfObject.(*PdfObjectDictionary)
		_eagb._gbcg = _ebbc
		_eb.Log.Trace("\u0032\u003a\u0020%\u0071", _ebdd)
		if !_accgg {
			return false, _f.New("\u0074\u0072a\u0069\u006c\u0065\u0072 \u0045\u006ec\u0072\u0079\u0070\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		_dbd = _ebdd
	case *PdfObjectNull:
		_eb.Log.Debug("\u0045\u006e\u0063\u0072\u0079\u0070\u0074 \u0069\u0073\u0020a\u0020\u006e\u0075l\u006c\u0020o\u0062\u006a\u0065\u0063\u0074\u002e \u0046il\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u002e")
		return false, nil
	default:
		return false, _ee.Errorf("u\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065: \u0025\u0054", _cdad)
	}
	_aabd, _gdff := PdfCryptNewDecrypt(_eagb, _dbd, _eagb._cbeb)
	if _gdff != nil {
		return false, _gdff
	}
	for _, _cfea := range []string{"\u0045n\u0063\u0072\u0079\u0070\u0074"} {
		_dgeee := _eagb._cbeb.Get(PdfObjectName(_cfea))
		if _dgeee == nil {
			continue
		}
		switch _fcadb := _dgeee.(type) {
		case *PdfObjectReference:
			_aabd._eda[int(_fcadb.ObjectNumber)] = struct{}{}
		case *PdfIndirectObject:
			_aabd._ge[_fcadb] = true
			_aabd._eda[int(_fcadb.ObjectNumber)] = struct{}{}
		}
	}
	_eagb._eccc = _aabd
	_eb.Log.Trace("\u0043\u0072\u0079\u0070\u0074\u0065\u0072\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0025\u0062", _aabd)
	return true, nil
}

// GetNumberAsInt64 returns the contents of `obj` as an int64 if it is an integer or float, or an
// error if it isn't. This is for cases where expecting an integer, but some implementations
// actually store the number in a floating point format.
func GetNumberAsInt64(obj PdfObject) (int64, error) {
	switch _egbgg := obj.(type) {
	case *PdfObjectFloat:
		_eb.Log.Debug("\u004e\u0075m\u0062\u0065\u0072\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u0073\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0077\u0061s\u0020\u0073\u0074\u006f\u0072\u0065\u0064\u0020\u0061\u0073\u0020\u0066\u006c\u006fa\u0074\u0020(\u0074\u0079\u0070\u0065 \u0063\u0061\u0073\u0074\u0069n\u0067\u0020\u0075\u0073\u0065\u0064\u0029")
		return int64(*_egbgg), nil
	case *PdfObjectInteger:
		return int64(*_egbgg), nil
	case *PdfObjectReference:
		_eddee := TraceToDirectObject(obj)
		return GetNumberAsInt64(_eddee)
	case *PdfIndirectObject:
		return GetNumberAsInt64(_egbgg.PdfObject)
	}
	return 0, ErrNotANumber
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_acaeb *CCITTFaxEncoder) MakeDecodeParams() PdfObject {
	_ceec := MakeDict()
	_ceec.Set("\u004b", MakeInteger(int64(_acaeb.K)))
	_ceec.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_acaeb.Columns)))
	if _acaeb.BlackIs1 {
		_ceec.Set("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031", MakeBool(_acaeb.BlackIs1))
	}
	if _acaeb.EncodedByteAlign {
		_ceec.Set("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e", MakeBool(_acaeb.EncodedByteAlign))
	}
	if _acaeb.EndOfLine && _acaeb.K >= 0 {
		_ceec.Set("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee", MakeBool(_acaeb.EndOfLine))
	}
	if _acaeb.Rows != 0 && !_acaeb.EndOfBlock {
		_ceec.Set("\u0052\u006f\u0077\u0073", MakeInteger(int64(_acaeb.Rows)))
	}
	if !_acaeb.EndOfBlock {
		_ceec.Set("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b", MakeBool(_acaeb.EndOfBlock))
	}
	if _acaeb.DamagedRowsBeforeError != 0 {
		_ceec.Set("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072", MakeInteger(int64(_acaeb.DamagedRowsBeforeError)))
	}
	return _ceec
}

// MakeIndirectObject creates an PdfIndirectObject with a specified direct object PdfObject.
func MakeIndirectObject(obj PdfObject) *PdfIndirectObject {
	_adbd := &PdfIndirectObject{}
	_adbd.PdfObject = obj
	return _adbd
}

// NewJPXEncoder returns a new instance of JPXEncoder.
func NewJPXEncoder() *JPXEncoder { return &JPXEncoder{} }

// ParseNumber parses a numeric objects from a buffered stream.
// Section 7.3.3.
// Integer or Float.
//
// An integer shall be written as one or more decimal digits optionally
// preceded by a sign. The value shall be interpreted as a signed
// decimal integer and shall be converted to an integer object.
//
// A real value shall be written as one or more decimal digits with an
// optional sign and a leading, trailing, or embedded PERIOD (2Eh)
// (decimal point). The value shall be interpreted as a real number
// and shall be converted to a real object.
//
// Regarding exponential numbers: 7.3.3 Numeric Objects:
// A conforming writer shall not use the PostScript syntax for numbers
// with non-decimal radices (such as 16#FFFE) or in exponential format
// (such as 6.02E23).
// Nonetheless, we sometimes get numbers with exponential format, so
// we will support it in the reader (no confusion with other types, so
// no compromise).
func ParseNumber(buf *_ga.Reader) (PdfObject, error) {
	_dbde := false
	_aacg := true
	var _abda _bg.Buffer
	for {
		if _eb.Log.IsLogLevel(_eb.LogLevelTrace) {
			_eb.Log.Trace("\u0050\u0061\u0072\u0073in\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0022\u0025\u0073\u0022", _abda.String())
		}
		_cffff, _cgdb := buf.Peek(1)
		if _cgdb == _bb.EOF {
			break
		}
		if _cgdb != nil {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0025\u0073", _cgdb)
			return nil, _cgdb
		}
		if _aacg && (_cffff[0] == '-' || _cffff[0] == '+') {
			_cgdcc, _ := buf.ReadByte()
			_abda.WriteByte(_cgdcc)
			_aacg = false
		} else if IsDecimalDigit(_cffff[0]) {
			_bccgd, _ := buf.ReadByte()
			_abda.WriteByte(_bccgd)
		} else if _cffff[0] == '.' {
			_bace, _ := buf.ReadByte()
			_abda.WriteByte(_bace)
			_dbde = true
		} else if _cffff[0] == 'e' || _cffff[0] == 'E' {
			_fgaae, _ := buf.ReadByte()
			_abda.WriteByte(_fgaae)
			_dbde = true
			_aacg = true
		} else {
			break
		}
	}
	var _dedbd PdfObject
	if _dbde {
		_egbff, _cgcc := _d.ParseFloat(_abda.String(), 64)
		if _cgcc != nil {
			_eb.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025v\u0020\u0065\u0072\u0072\u003d\u0025v\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0030\u002e\u0020\u004fu\u0074\u0070u\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074", _abda.String(), _cgcc)
			_egbff = 0.0
		}
		_dcba := PdfObjectFloat(_egbff)
		_dedbd = &_dcba
	} else {
		_eddf, _feafa := _d.ParseInt(_abda.String(), 10, 64)
		if _feafa != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0025\u0076\u0020\u0065\u0072\u0072\u003d%\u0076\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 \u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006ec\u006f\u0072r\u0065c\u0074", _abda.String(), _feafa)
			_eddf = 0
		}
		_cgfb := PdfObjectInteger(_eddf)
		_dedbd = &_cgfb
	}
	return _dedbd, nil
}

// MakeFloat creates an PdfObjectFloat from a float64.
func MakeFloat(val float64) *PdfObjectFloat { _cgbf := PdfObjectFloat(val); return &_cgbf }

// DecodeStream decodes a multi-encoded stream by passing it through the
// DecodeStream method of the underlying encoders.
func (_gbea *MultiEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gbea.DecodeBytes(streamObj.Stream)
}

// Clear resets the array to an empty state.
func (_egfe *PdfObjectArray) Clear() { _egfe._aagc = []PdfObject{} }

// DecodeBytes decodes a byte slice from Run length encoding.
//
// 7.4.5 RunLengthDecode Filter
// The RunLengthDecode filter decodes data that has been encoded in a simple byte-oriented format based on run length.
// The encoded data shall be a sequence of runs, where each run shall consist of a length byte followed by 1 to 128
// bytes of data. If the length byte is in the range 0 to 127, the following length + 1 (1 to 128) bytes shall be
// copied literally during decompression. If length is in the range 129 to 255, the following single byte shall be
// copied 257 - length (2 to 128) times during decompression. A length value of 128 shall denote EOD.
func (_dgad *RunLengthEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_feaa := _bg.NewReader(encoded)
	var _aacd []byte
	for {
		_eagc, _fgb := _feaa.ReadByte()
		if _fgb != nil {
			return nil, _fgb
		}
		if _eagc > 128 {
			_add, _ccef := _feaa.ReadByte()
			if _ccef != nil {
				return nil, _ccef
			}
			for _aefgb := 0; _aefgb < 257-int(_eagc); _aefgb++ {
				_aacd = append(_aacd, _add)
			}
		} else if _eagc < 128 {
			for _ccge := 0; _ccge < int(_eagc)+1; _ccge++ {
				_ebgf, _geb := _feaa.ReadByte()
				if _geb != nil {
					return nil, _geb
				}
				_aacd = append(_aacd, _ebgf)
			}
		} else {
			break
		}
	}
	return _aacd, nil
}
func (_dgadc *PdfParser) parseHexString() (*PdfObjectString, error) {
	_dgadc._dedbc.ReadByte()
	var _cggcf _bg.Buffer
	for {
		_edceg, _bgbdb := _dgadc._dedbc.Peek(1)
		if _bgbdb != nil {
			return MakeString(""), _bgbdb
		}
		if _edceg[0] == '>' {
			_dgadc._dedbc.ReadByte()
			break
		}
		_bgdg, _ := _dgadc._dedbc.ReadByte()
		if _dgadc._aega {
			if _bg.IndexByte(_dffca, _bgdg) == -1 {
				_dgadc._eggc._gca = true
			}
		}
		if !IsWhiteSpace(_bgdg) {
			_cggcf.WriteByte(_bgdg)
		}
	}
	if _cggcf.Len()%2 == 1 {
		_dgadc._eggc._cdff = true
		_cggcf.WriteRune('0')
	}
	_baaf, _ := _ae.DecodeString(_cggcf.String())
	return MakeHexString(string(_baaf)), nil
}
func _dcbcc(_gecc, _aeaad PdfObject, _gaba int) bool {
	if _gaba > _gegcg {
		_eb.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _gegcg)
		return false
	}
	if _gecc == nil && _aeaad == nil {
		return true
	} else if _gecc == nil || _aeaad == nil {
		return false
	}
	if _fa.TypeOf(_gecc) != _fa.TypeOf(_aeaad) {
		return false
	}
	switch _fcgg := _gecc.(type) {
	case *PdfObjectNull, *PdfObjectReference:
		return true
	case *PdfObjectName:
		return *_fcgg == *(_aeaad.(*PdfObjectName))
	case *PdfObjectString:
		return *_fcgg == *(_aeaad.(*PdfObjectString))
	case *PdfObjectInteger:
		return *_fcgg == *(_aeaad.(*PdfObjectInteger))
	case *PdfObjectBool:
		return *_fcgg == *(_aeaad.(*PdfObjectBool))
	case *PdfObjectFloat:
		return *_fcgg == *(_aeaad.(*PdfObjectFloat))
	case *PdfIndirectObject:
		return _dcbcc(TraceToDirectObject(_gecc), TraceToDirectObject(_aeaad), _gaba+1)
	case *PdfObjectArray:
		_cefgc := _aeaad.(*PdfObjectArray)
		if len((*_fcgg)._aagc) != len((*_cefgc)._aagc) {
			return false
		}
		for _bdgf, _dcfcf := range (*_fcgg)._aagc {
			if !_dcbcc(_dcfcf, (*_cefgc)._aagc[_bdgf], _gaba+1) {
				return false
			}
		}
		return true
	case *PdfObjectDictionary:
		_deed := _aeaad.(*PdfObjectDictionary)
		_aafb, _fdabf := (*_fcgg)._abec, (*_deed)._abec
		if len(_aafb) != len(_fdabf) {
			return false
		}
		for _bedc, _dfaa := range _aafb {
			_afgaf, _ggea := _fdabf[_bedc]
			if !_ggea || !_dcbcc(_dfaa, _afgaf, _gaba+1) {
				return false
			}
		}
		return true
	case *PdfObjectStream:
		_adgfe := _aeaad.(*PdfObjectStream)
		return _dcbcc((*_fcgg).PdfObjectDictionary, (*_adgfe).PdfObjectDictionary, _gaba+1)
	default:
		_eb.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", _gecc)
	}
	return false
}

// DecodeBytes decodes a slice of DCT encoded bytes and returns the result.
func (_cfd *DCTEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_gega := _bg.NewReader(encoded)
	_cgc, _bgdc := _beg.Decode(_gega)
	if _bgdc != nil {
		_eb.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bgdc)
		return nil, _bgdc
	}
	_begd := _cgc.Bounds()
	var _efbc = make([]byte, _begd.Dx()*_begd.Dy()*_cfd.ColorComponents*_cfd.BitsPerComponent/8)
	_egfc := 0
	switch _cfd.ColorComponents {
	case 1:
		_eege := []float64{_cfd.Decode[0], _cfd.Decode[1]}
		for _dccg := _begd.Min.Y; _dccg < _begd.Max.Y; _dccg++ {
			for _fdd := _begd.Min.X; _fdd < _begd.Max.X; _fdd++ {
				_fefa := _cgc.At(_fdd, _dccg)
				if _cfd.BitsPerComponent == 16 {
					_faea, _gacc := _fefa.(_cb.Gray16)
					if !_gacc {
						return nil, _f.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_abeff := _fdfc(uint(_faea.Y>>8), _eege[0], _eege[1])
					_bfca := _fdfc(uint(_faea.Y), _eege[0], _eege[1])
					_efbc[_egfc] = byte(_abeff)
					_egfc++
					_efbc[_egfc] = byte(_bfca)
					_egfc++
				} else {
					_aebae, _fdfa := _fefa.(_cb.Gray)
					if !_fdfa {
						return nil, _f.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_efbc[_egfc] = byte(_fdfc(uint(_aebae.Y), _eege[0], _eege[1]))
					_egfc++
				}
			}
		}
	case 3:
		_ebcd := []float64{_cfd.Decode[0], _cfd.Decode[1]}
		_dedea := []float64{_cfd.Decode[2], _cfd.Decode[3]}
		_gbgg := []float64{_cfd.Decode[4], _cfd.Decode[5]}
		for _ggd := _begd.Min.Y; _ggd < _begd.Max.Y; _ggd++ {
			for _befb := _begd.Min.X; _befb < _begd.Max.X; _befb++ {
				_dee := _cgc.At(_befb, _ggd)
				if _cfd.BitsPerComponent == 16 {
					_cgdg, _aabge := _dee.(_cb.RGBA64)
					if !_aabge {
						return nil, _f.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_gaag := _fdfc(uint(_cgdg.R>>8), _ebcd[0], _ebcd[1])
					_gddg := _fdfc(uint(_cgdg.R), _ebcd[0], _ebcd[1])
					_bff := _fdfc(uint(_cgdg.G>>8), _dedea[0], _dedea[1])
					_gdda := _fdfc(uint(_cgdg.G), _dedea[0], _dedea[1])
					_cddf := _fdfc(uint(_cgdg.B>>8), _gbgg[0], _gbgg[1])
					_cffc := _fdfc(uint(_cgdg.B), _gbgg[0], _gbgg[1])
					_efbc[_egfc] = byte(_gaag)
					_egfc++
					_efbc[_egfc] = byte(_gddg)
					_egfc++
					_efbc[_egfc] = byte(_bff)
					_egfc++
					_efbc[_egfc] = byte(_gdda)
					_egfc++
					_efbc[_egfc] = byte(_cddf)
					_egfc++
					_efbc[_egfc] = byte(_cffc)
					_egfc++
				} else {
					_bebd, _eebe := _dee.(_cb.RGBA)
					if _eebe {
						_gacd := _fdfc(uint(_bebd.R), _ebcd[0], _ebcd[1])
						_defd := _fdfc(uint(_bebd.G), _dedea[0], _dedea[1])
						_edb := _fdfc(uint(_bebd.B), _gbgg[0], _gbgg[1])
						_efbc[_egfc] = byte(_gacd)
						_egfc++
						_efbc[_egfc] = byte(_defd)
						_egfc++
						_efbc[_egfc] = byte(_edb)
						_egfc++
					} else {
						_beeg, _aea := _dee.(_cb.YCbCr)
						if !_aea {
							return nil, _f.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
						}
						_dedc, _ggfb, _ggac, _ := _beeg.RGBA()
						_efeg := _fdfc(uint(_dedc>>8), _ebcd[0], _ebcd[1])
						_gbca := _fdfc(uint(_ggfb>>8), _dedea[0], _dedea[1])
						_bfge := _fdfc(uint(_ggac>>8), _gbgg[0], _gbgg[1])
						_efbc[_egfc] = byte(_efeg)
						_egfc++
						_efbc[_egfc] = byte(_gbca)
						_egfc++
						_efbc[_egfc] = byte(_bfge)
						_egfc++
					}
				}
			}
		}
	case 4:
		_egd := []float64{_cfd.Decode[0], _cfd.Decode[1]}
		_ddfe := []float64{_cfd.Decode[2], _cfd.Decode[3]}
		_deeb := []float64{_cfd.Decode[4], _cfd.Decode[5]}
		_dcfb := []float64{_cfd.Decode[6], _cfd.Decode[7]}
		for _gaea := _begd.Min.Y; _gaea < _begd.Max.Y; _gaea++ {
			for _fcef := _begd.Min.X; _fcef < _begd.Max.X; _fcef++ {
				_egbb := _cgc.At(_fcef, _gaea)
				_bcde, _efaf := _egbb.(_cb.CMYK)
				if !_efaf {
					return nil, _f.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
				}
				_gaeb := 255 - _fdfc(uint(_bcde.C), _egd[0], _egd[1])
				_eac := 255 - _fdfc(uint(_bcde.M), _ddfe[0], _ddfe[1])
				_bgbg := 255 - _fdfc(uint(_bcde.Y), _deeb[0], _deeb[1])
				_adad := 255 - _fdfc(uint(_bcde.K), _dcfb[0], _dcfb[1])
				_efbc[_egfc] = byte(_gaeb)
				_egfc++
				_efbc[_egfc] = byte(_eac)
				_egfc++
				_efbc[_egfc] = byte(_bgbg)
				_egfc++
				_efbc[_egfc] = byte(_adad)
				_egfc++
			}
		}
	}
	return _efbc, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_cedg *MultiEncoder) MakeDecodeParams() PdfObject {
	if len(_cedg._abeab) == 0 {
		return nil
	}
	if len(_cedg._abeab) == 1 {
		return _cedg._abeab[0].MakeDecodeParams()
	}
	_eafg := MakeArray()
	_eed := true
	for _, _dfdbf := range _cedg._abeab {
		_dccf := _dfdbf.MakeDecodeParams()
		if _dccf == nil {
			_eafg.Append(MakeNull())
		} else {
			_eed = false
			_eafg.Append(_dccf)
		}
	}
	if _eed {
		return nil
	}
	return _eafg
}

// GoImageToJBIG2 creates a binary image on the base of 'i' golang image.Image.
// If the image is not a black/white image then the function converts provided input into
// JBIG2Image with 1bpp. For non grayscale images the function performs the conversion to the grayscale temp image.
// Then it checks the value of the gray image value if it's within bounds of the black white threshold.
// This 'bwThreshold' value should be in range (0.0, 1.0). The threshold checks if the grayscale pixel (uint) value
// is greater or smaller than 'bwThreshold' * 255. Pixels inside the range will be white, and the others will be black.
// If the 'bwThreshold' is equal to -1.0 - JB2ImageAutoThreshold then it's value would be set on the base of
// it's histogram using Triangle method. For more information go to:
//
//	https://www.mathworks.com/matlabcentral/fileexchange/28047-gray-image-thresholding-using-the-triangle-method
func GoImageToJBIG2(i _a.Image, bwThreshold float64) (*JBIG2Image, error) {
	const _gbbg = "\u0047\u006f\u0049\u006d\u0061\u0067\u0065\u0054\u006fJ\u0042\u0049\u0047\u0032"
	if i == nil {
		return nil, _eca.Error(_gbbg, "i\u006d\u0061\u0067\u0065 '\u0069'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	var (
		_gddb uint8
		_geeb _eeb.Image
		_effe error
	)
	if bwThreshold == JB2ImageAutoThreshold {
		_geeb, _effe = _eeb.MonochromeConverter.Convert(i)
	} else if bwThreshold > 1.0 || bwThreshold < 0.0 {
		return nil, _eca.Error(_gbbg, "p\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074h\u0072\u0065\u0073\u0068\u006f\u006c\u0064 i\u0073\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u0020\u0072\u0061\u006ege\u0020\u007b0\u002e\u0030\u002c\u0020\u0031\u002e\u0030\u007d")
	} else {
		_gddb = uint8(255 * bwThreshold)
		_geeb, _effe = _eeb.MonochromeThresholdConverter(_gddb).Convert(i)
	}
	if _effe != nil {
		return nil, _effe
	}
	return _dcca(_geeb), nil
}

// MakeStreamDict make a new instance of an encoding dictionary for a stream object.
func (_gcg *ASCII85Encoder) MakeStreamDict() *PdfObjectDictionary {
	_gfbf := MakeDict()
	_gfbf.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_gcg.GetFilterName()))
	return _gfbf
}

// ReadBytesAt reads byte content at specific offset and length within the PDF.
func (_cfgf *PdfParser) ReadBytesAt(offset, len int64) ([]byte, error) {
	_cafd := _cfgf.GetFileOffset()
	_, _ddgg := _cfgf._cdea.Seek(offset, _bb.SeekStart)
	if _ddgg != nil {
		return nil, _ddgg
	}
	_afgb := make([]byte, len)
	_, _ddgg = _bb.ReadAtLeast(_cfgf._cdea, _afgb, int(len))
	if _ddgg != nil {
		return nil, _ddgg
	}
	_cfgf.SetFileOffset(_cafd)
	return _afgb, nil
}

// GetFilterArray returns the names of the underlying encoding filters in an array that
// can be used as /Filter entry.
func (_bafe *MultiEncoder) GetFilterArray() *PdfObjectArray {
	_dgbd := make([]PdfObject, len(_bafe._abeab))
	for _dcef, _acfec := range _bafe._abeab {
		_dgbd[_dcef] = MakeName(_acfec.GetFilterName())
	}
	return MakeArray(_dgbd...)
}
func (_fba *PdfParser) getNumbersOfUpdatedObjects(_cged *PdfParser) ([]int, error) {
	if _cged == nil {
		return nil, _f.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_defgd := _cged._ggdf
	_fefae := make([]int, 0)
	_gddag := make(map[int]interface{})
	_ebad := make(map[int]int64)
	for _faeae, _egfg := range _fba._bfba.ObjectMap {
		if _egfg.Offset == 0 {
			if _egfg.OsObjNumber != 0 {
				if _bdbg, _gdae := _fba._bfba.ObjectMap[_egfg.OsObjNumber]; _gdae {
					_gddag[_egfg.OsObjNumber] = struct{}{}
					_ebad[_faeae] = _bdbg.Offset
				} else {
					return nil, _f.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0078r\u0065\u0066\u0020\u0074ab\u006c\u0065")
				}
			}
		} else {
			_ebad[_faeae] = _egfg.Offset
		}
	}
	for _adef, _efba := range _ebad {
		if _, _cgcf := _gddag[_adef]; _cgcf {
			continue
		}
		if _efba > _defgd {
			_fefae = append(_fefae, _adef)
		}
	}
	return _fefae, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_eaac *LZWEncoder) MakeDecodeParams() PdfObject {
	if _eaac.Predictor > 1 {
		_dbbc := MakeDict()
		_dbbc.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_eaac.Predictor)))
		if _eaac.BitsPerComponent != 8 {
			_dbbc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_eaac.BitsPerComponent)))
		}
		if _eaac.Columns != 1 {
			_dbbc.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_eaac.Columns)))
		}
		if _eaac.Colors != 1 {
			_dbbc.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_eaac.Colors)))
		}
		return _dbbc
	}
	return nil
}
func (_dgg *PdfCrypt) makeKey(_afa string, _cggb, _gage uint32, _edf []byte) ([]byte, error) {
	_edd, _fdb := _dgg._ace[_afa]
	if !_fdb {
		return nil, _ee.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _afa)
	}
	return _edd.MakeKey(_cggb, _gage, _edf)
}

// TraceToDirectObject traces a PdfObject to a direct object.  For example direct objects contained
// in indirect objects (can be double referenced even).
func TraceToDirectObject(obj PdfObject) PdfObject {
	if _ccgeg, _aacdg := obj.(*PdfObjectReference); _aacdg {
		obj = _ccgeg.Resolve()
	}
	_efda, _gafbe := obj.(*PdfIndirectObject)
	_gfbca := 0
	for _gafbe {
		obj = _efda.PdfObject
		_efda, _gafbe = GetIndirect(obj)
		_gfbca++
		if _gfbca > _gegcg {
			_eb.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0061\u0063\u0065\u0020\u0064\u0065p\u0074\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u0062\u0065\u0079\u006fn\u0064\u0020\u0025\u0064\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0067oi\u006e\u0067\u0020\u0064\u0065\u0065\u0070\u0065\u0072\u0021", _gegcg)
			return nil
		}
	}
	return obj
}

// GetCrypter returns the PdfCrypt instance which has information about the PDFs encryption.
func (_decf *PdfParser) GetCrypter() *PdfCrypt { return _decf._eccc }
func _acgde(_cggf *PdfObjectDictionary) (_gffd *_eeb.ImageBase) {
	var (
		_bgfb  *PdfObjectInteger
		_ffgdf bool
	)
	if _bgfb, _ffgdf = _cggf.Get("\u0057\u0069\u0064t\u0068").(*PdfObjectInteger); _ffgdf {
		_gffd = &_eeb.ImageBase{Width: int(*_bgfb)}
	} else {
		return nil
	}
	if _bgfb, _ffgdf = _cggf.Get("\u0048\u0065\u0069\u0067\u0068\u0074").(*PdfObjectInteger); _ffgdf {
		_gffd.Height = int(*_bgfb)
	}
	if _bgfb, _ffgdf = _cggf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074").(*PdfObjectInteger); _ffgdf {
		_gffd.BitsPerComponent = int(*_bgfb)
	}
	if _bgfb, _ffgdf = _cggf.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073").(*PdfObjectInteger); _ffgdf {
		_gffd.ColorComponents = int(*_bgfb)
	}
	return _gffd
}

// PdfObjectDictionary represents the primitive PDF dictionary/map object.
type PdfObjectDictionary struct {
	_abec   map[PdfObjectName]PdfObject
	_begbb  []PdfObjectName
	_bfdf   *_g.Mutex
	_eddefe *PdfParser
}

// ToFloat64Array returns a slice of all elements in the array as a float64 slice.  An error is
// returned if the array contains non-numeric objects (each element can be either PdfObjectInteger
// or PdfObjectFloat).
func (_gfbb *PdfObjectArray) ToFloat64Array() ([]float64, error) {
	var _ffbc []float64
	for _, _abca := range _gfbb.Elements() {
		switch _cdbf := _abca.(type) {
		case *PdfObjectInteger:
			_ffbc = append(_ffbc, float64(*_cdbf))
		case *PdfObjectFloat:
			_ffbc = append(_ffbc, float64(*_cdbf))
		default:
			return nil, ErrTypeError
		}
	}
	return _ffbc, nil
}

// IsHexadecimal checks if the PdfObjectString contains Hexadecimal data.
func (_fbeg *PdfObjectString) IsHexadecimal() bool { return _fbeg._cdca }

// PdfObjectArray represents the primitive PDF array object.
type PdfObjectArray struct{ _aagc []PdfObject }

func (_bce *PdfCrypt) checkAccessRights(_dga []byte) (bool, _cbd.Permissions, error) {
	_bdb := _bce.securityHandler()
	_gdbg, _abe, _cfff := _bdb.Authenticate(&_bce._adfe, _dga)
	if _cfff != nil {
		return false, 0, _cfff
	} else if _abe == 0 || len(_gdbg) == 0 {
		return false, 0, nil
	}
	return true, _abe, nil
}
func (_gafd *PdfParser) seekPdfVersionTopDown() (int, int, error) {
	_gafd._cdea.Seek(0, _bb.SeekStart)
	_gafd._dedbc = _ga.NewReader(_gafd._cdea)
	_cbeg := 20
	_daggg := make([]byte, _cbeg)
	for {
		_gaeae, _feec := _gafd._dedbc.ReadByte()
		if _feec != nil {
			if _feec == _bb.EOF {
				break
			} else {
				return 0, 0, _feec
			}
		}
		if IsDecimalDigit(_gaeae) && _daggg[_cbeg-1] == '.' && IsDecimalDigit(_daggg[_cbeg-2]) && _daggg[_cbeg-3] == '-' && _daggg[_cbeg-4] == 'F' && _daggg[_cbeg-5] == 'D' && _daggg[_cbeg-6] == 'P' {
			_dgff := int(_daggg[_cbeg-2] - '0')
			_gggg := int(_gaeae - '0')
			return _dgff, _gggg, nil
		}
		_daggg = append(_daggg[1:_cbeg], _gaeae)
	}
	return 0, 0, _f.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// DecodeGlobals decodes 'encoded' byte stream and returns their Globally defined segments ('Globals').
func (_gggb *JBIG2Encoder) DecodeGlobals(encoded []byte) (_gd.Globals, error) {
	return _gd.DecodeGlobals(encoded)
}

// DecodeStream implements ASCII hex decoding.
func (_agfc *ASCIIHexEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _agfc.DecodeBytes(streamObj.Stream)
}

// NewDCTEncoder makes a new DCT encoder with default parameters.
func NewDCTEncoder() *DCTEncoder {
	_dfdc := &DCTEncoder{}
	_dfdc.ColorComponents = 3
	_dfdc.BitsPerComponent = 8
	_dfdc.Quality = DefaultJPEGQuality
	_dfdc.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	return _dfdc
}

// HasInvalidSeparationAfterXRef implements core.ParserMetadata interface.
func (_bded ParserMetadata) HasInvalidSeparationAfterXRef() bool { return _bded._bgb }

// JBIG2Image is the image structure used by the jbig2 encoder. Its Data must be in a
// 1 bit per component and 1 component per pixel (1bpp). In order to create binary image
// use GoImageToJBIG2 function. If the image data contains the row bytes padding set the HasPadding to true.
type JBIG2Image struct {

	// Width and Height defines the image boundaries.
	Width, Height int

	// Data is the byte slice data for the input image
	Data []byte

	// HasPadding is the attribute that defines if the last byte of the data in the row contains
	// 0 bits padding.
	HasPadding bool
}

// MakeArrayFromIntegers creates an PdfObjectArray from a slice of ints, where each array element is
// an PdfObjectInteger.
func MakeArrayFromIntegers(vals []int) *PdfObjectArray {
	_eedcf := MakeArray()
	for _, _fadc := range vals {
		_eedcf.Append(MakeInteger(int64(_fadc)))
	}
	return _eedcf
}

// PdfObjectStream represents the primitive PDF Object stream.
type PdfObjectStream struct {
	PdfObjectReference
	*PdfObjectDictionary
	Stream   []byte
	Lazy     bool
	TempFile string
}

// GetStringBytes is like GetStringVal except that it returns the string as a []byte.
// It is for convenience.
func GetStringBytes(obj PdfObject) (_bcdf []byte, _accc bool) {
	_edffg, _accc := TraceToDirectObject(obj).(*PdfObjectString)
	if _accc {
		return _edffg.Bytes(), true
	}
	return
}
func (_caada *PdfParser) inspect() (map[string]int, error) {
	_eb.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0053P\u0045\u0043\u0054\u0020\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d")
	_eb.Log.Trace("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_badbd := map[string]int{}
	_cdgcf := 0
	_fbff := 0
	var _cdee []int
	for _cbafb := range _caada._bfba.ObjectMap {
		_cdee = append(_cdee, _cbafb)
	}
	_be.Ints(_cdee)
	_defba := 0
	for _, _feddb := range _cdee {
		_agefd := _caada._bfba.ObjectMap[_feddb]
		if _agefd.ObjectNumber == 0 {
			continue
		}
		_cdgcf++
		_eb.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
		_eb.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0075\u0070\u0020\u006f\u0062\u006a\u0065c\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0025\u0064", _agefd.ObjectNumber)
		_ecgg, _dgaf := _caada.LookupByNumber(_agefd.ObjectNumber)
		if _dgaf != nil {
			_eb.Log.Trace("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u006c\u006f\u006f\u006b\u0075p\u0020\u006f\u0062\u006a\u0020\u0025\u0064 \u0028\u0025\u0073\u0029", _agefd.ObjectNumber, _dgaf)
			_fbff++
			continue
		}
		_eb.Log.Trace("\u006fb\u006a\u003a\u0020\u0025\u0073", _ecgg)
		_cdag, _gfgb := _ecgg.(*PdfIndirectObject)
		if _gfgb {
			_eb.Log.Trace("\u0049N\u0044 \u004f\u004f\u0042\u004a\u0020\u0025\u0064\u003a\u0020\u0025\u0073", _agefd.ObjectNumber, _cdag)
			_degg, _gdece := _cdag.PdfObject.(*PdfObjectDictionary)
			if _gdece {
				if _effc, _dfaf := _degg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _dfaf {
					_cfec := string(*_effc)
					_eb.Log.Trace("\u002d\u002d\u002d\u003e\u0020\u004f\u0062\u006a\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _cfec)
					_, _fedcf := _badbd[_cfec]
					if _fedcf {
						_badbd[_cfec]++
					} else {
						_badbd[_cfec] = 1
					}
				} else if _cgaag, _cfeg := _degg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065").(*PdfObjectName); _cfeg {
					_geebd := string(*_cgaag)
					_eb.Log.Trace("-\u002d-\u003e\u0020\u004f\u0062\u006a\u0020\u0073\u0075b\u0074\u0079\u0070\u0065: \u0025\u0073", _geebd)
					_, _efbcg := _badbd[_geebd]
					if _efbcg {
						_badbd[_geebd]++
					} else {
						_badbd[_geebd] = 1
					}
				}
				if _eabfa, _aecda := _degg.Get("\u0053").(*PdfObjectName); _aecda && *_eabfa == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
					_, _aebfa := _badbd["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]
					if _aebfa {
						_badbd["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
					} else {
						_badbd["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"] = 1
					}
				}
			}
		} else if _gfbba, _beed := _ecgg.(*PdfObjectStream); _beed {
			if _gcbfg, _egbge := _gfbba.PdfObjectDictionary.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _egbge {
				_eb.Log.Trace("\u002d\u002d\u003e\u0020\u0053\u0074\u0072\u0065\u0061\u006d\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073", *_gcbfg)
				_decbb := string(*_gcbfg)
				_badbd[_decbb]++
			}
		} else {
			_dgcc, _cagg := _ecgg.(*PdfObjectDictionary)
			if _cagg {
				_gbeae, _cedc := _dgcc.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
				if _cedc {
					_ecae := string(*_gbeae)
					_eb.Log.Trace("\u002d-\u002d \u006f\u0062\u006a\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _ecae)
					_badbd[_ecae]++
				}
			}
			_eb.Log.Trace("\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u004f\u0042\u004a\u0020\u0025d\u003a\u0020\u0025\u0073", _agefd.ObjectNumber, _ecgg)
		}
		_defba++
	}
	_eb.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0045\u004fF\u0020\u0049\u004e\u0053\u0050\u0045\u0043T\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d")
	_eb.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	_eb.Log.Trace("\u004f\u0062j\u0065\u0063\u0074 \u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064", _cdgcf)
	_eb.Log.Trace("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u006f\u006b\u0075p\u003a\u0020\u0025\u0064", _fbff)
	for _egca, _bcdea := range _badbd {
		_eb.Log.Trace("\u0025\u0073\u003a\u0020\u0025\u0064", _egca, _bcdea)
	}
	_eb.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	if len(_caada._bfba.ObjectMap) < 1 {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0068\u0069\u0073 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074 \u0069s\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0021\u0029")
		return nil, _ee.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0028\u0078r\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u0029")
	}
	_fdcc, _gcafc := _badbd["\u0046\u006f\u006e\u0074"]
	if !_gcafc || _fdcc < 2 {
		_eb.Log.Trace("\u0054\u0068\u0069s \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073 \u0070r\u006fb\u0061\u0062\u006c\u0079\u0020\u0073\u0063\u0061\u006e\u006e\u0065\u0064\u0021")
	} else {
		_eb.Log.Trace("\u0054\u0068\u0069\u0073\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0066o\u0072\u0020\u0065\u0078\u0074r\u0061\u0063t\u0069\u006f\u006e\u0021")
	}
	return _badbd, nil
}

// UpdateParams updates the parameter values of the encoder.
func (_baae *CCITTFaxEncoder) UpdateParams(params *PdfObjectDictionary) {
	if _cabgg, _ebeg := GetNumberAsInt64(params.Get("\u004b")); _ebeg == nil {
		_baae.K = int(_cabgg)
	}
	if _ebfb, _cea := GetNumberAsInt64(params.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _cea == nil {
		_baae.Columns = int(_ebfb)
	} else if _ebfb, _cea = GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068")); _cea == nil {
		_baae.Columns = int(_ebfb)
	}
	if _cacb, _baaa := GetNumberAsInt64(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _baaa == nil {
		_baae.BlackIs1 = _cacb > 0
	} else {
		if _bfe, _bed := GetBoolVal(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _bed {
			_baae.BlackIs1 = _bfe
		} else {
			if _eeed, _bbag := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _bbag {
				_ddcc, _babf := _eeed.ToIntegerArray()
				if _babf == nil {
					_baae.BlackIs1 = _ddcc[0] == 1 && _ddcc[1] == 0
				}
			}
		}
	}
	if _abeb, _ddabf := GetNumberAsInt64(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _ddabf == nil {
		_baae.EncodedByteAlign = _abeb > 0
	} else {
		if _acde, _cfab := GetBoolVal(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _cfab {
			_baae.EncodedByteAlign = _acde
		}
	}
	if _fcde, _bcee := GetNumberAsInt64(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _bcee == nil {
		_baae.EndOfLine = _fcde > 0
	} else {
		if _dadb, _cefeg := GetBoolVal(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _cefeg {
			_baae.EndOfLine = _dadb
		}
	}
	if _eaef, _ggdg := GetNumberAsInt64(params.Get("\u0052\u006f\u0077\u0073")); _ggdg == nil {
		_baae.Rows = int(_eaef)
	} else if _eaef, _ggdg = GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _ggdg == nil {
		_baae.Rows = int(_eaef)
	}
	if _abcd, _bgeg := GetNumberAsInt64(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _bgeg == nil {
		_baae.EndOfBlock = _abcd > 0
	} else {
		if _becd, _ffda := GetBoolVal(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _ffda {
			_baae.EndOfBlock = _becd
		}
	}
	if _fdea, _bdeg := GetNumberAsInt64(params.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _bdeg != nil {
		_baae.DamagedRowsBeforeError = int(_fdea)
	}
}

// AddPageImage adds the page with the image 'img' to the encoder context in order to encode it jbig2 document.
// The 'settings' defines what encoding type should be used by the encoder.
func (_ggg *JBIG2Encoder) AddPageImage(img *JBIG2Image, settings *JBIG2EncoderSettings) (_gfcf error) {
	const _dbed = "\u004a\u0042\u0049\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065\u0049m\u0061\u0067\u0065"
	if _ggg == nil {
		return _eca.Error(_dbed, "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u0069\u0073 n\u0069\u006c")
	}
	if settings == nil {
		settings = &_ggg.DefaultPageSettings
	}
	if _ggg._eabf == nil {
		_ggg._eabf = _ce.InitEncodeDocument(settings.FileMode)
	}
	if _gfcf = settings.Validate(); _gfcf != nil {
		return _eca.Wrap(_gfcf, _dbed, "")
	}
	_afgd, _gfcf := img.toBitmap()
	if _gfcf != nil {
		return _eca.Wrap(_gfcf, _dbed, "")
	}
	switch settings.Compression {
	case JB2Generic:
		if _gfcf = _ggg._eabf.AddGenericPage(_afgd, settings.DuplicatedLinesRemoval); _gfcf != nil {
			return _eca.Wrap(_gfcf, _dbed, "")
		}
	case JB2SymbolCorrelation:
		return _eca.Error(_dbed, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return _eca.Error(_dbed, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return _eca.Error(_dbed, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return nil
}
func (_egad *PdfParser) skipSpaces() (int, error) {
	_cafdb := 0
	for {
		_ggbd, _ggeg := _egad._dedbc.ReadByte()
		if _ggeg != nil {
			return 0, _ggeg
		}
		if IsWhiteSpace(_ggbd) {
			_cafdb++
		} else {
			_egad._dedbc.UnreadByte()
			break
		}
	}
	return _cafdb, nil
}

// Decoded returns the PDFDocEncoding or UTF-16BE decoded string contents.
// UTF-16BE is applied when the first two bytes are 0xFE, 0XFF, otherwise decoding of
// PDFDocEncoding is performed.
func (_beff *PdfObjectString) Decoded() string {
	if _beff == nil {
		return ""
	}
	_deafe := []byte(_beff._cbdg)
	if len(_deafe) >= 2 && _deafe[0] == 0xFE && _deafe[1] == 0xFF {
		return _dd.UTF16ToString(_deafe[2:])
	}
	return _dd.PDFDocEncodingToString(_deafe)
}

// SetFileOffset sets the file to an offset position and resets buffer.
func (_gbbge *PdfParser) SetFileOffset(offset int64) {
	if offset < 0 {
		offset = 0
	}
	_gbbge._cdea.Seek(offset, _bb.SeekStart)
	_gbbge._dedbc = _ga.NewReader(_gbbge._cdea)
}

// MakeArray creates an PdfObjectArray from a list of PdfObjects.
func MakeArray(objects ...PdfObject) *PdfObjectArray { return &PdfObjectArray{_aagc: objects} }
func _gcb(_gdf int) cryptFilters                     { return cryptFilters{_dedd: _gbe.NewFilterV2(_gdf)} }
func (_deab *PdfObjectInteger) String() string       { return _ee.Sprintf("\u0025\u0064", *_deab) }

// GetFilterName returns the name of the encoding filter.
func (_cebe *RawEncoder) GetFilterName() string { return StreamEncodingFilterNameRaw }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_fcg *FlateEncoder) MakeDecodeParams() PdfObject {
	if _fcg.Predictor > 1 {
		_cebc := MakeDict()
		_cebc.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_fcg.Predictor)))
		if _fcg.BitsPerComponent != 8 {
			_cebc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_fcg.BitsPerComponent)))
		}
		if _fcg.Columns != 1 {
			_cebc.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_fcg.Columns)))
		}
		if _fcg.Colors != 1 {
			_cebc.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_fcg.Colors)))
		}
		return _cebc
	}
	return nil
}

// IsPrintable checks if a character is printable.
// Regular characters that are outside the range EXCLAMATION MARK(21h)
// (!) to TILDE (7Eh) (~) should be written using the hexadecimal notation.
func IsPrintable(c byte) bool { return 0x21 <= c && c <= 0x7E }
func (_bfa *PdfCrypt) generateParams(_bcbe, _ddf []byte) error {
	_cefe := _bfa.securityHandler()
	_cbcga, _baf := _cefe.GenerateParams(&_bfa._adfe, _ddf, _bcbe)
	if _baf != nil {
		return _baf
	}
	_bfa._cbcg = _cbcga
	return nil
}

// GetFilterName returns the name of the encoding filter.
func (_daebd *JPXEncoder) GetFilterName() string { return StreamEncodingFilterNameJPX }

// GetString is a helper for Get that returns a string value.
// Returns false if the key is missing or a value is not a string.
func (_fafgc *PdfObjectDictionary) GetString(key PdfObjectName) (string, bool) {
	_agbc := _fafgc.Get(key)
	if _agbc == nil {
		return "", false
	}
	_bade, _cafge := _agbc.(*PdfObjectString)
	if !_cafge {
		return "", false
	}
	return _bade.Str(), true
}
func (_dfad *PdfParser) parseXrefTable() (*PdfObjectDictionary, error) {
	var _daeab *PdfObjectDictionary
	_bdab, _cbae := _dfad.readTextLine()
	if _cbae != nil {
		return nil, _cbae
	}
	if _dfad._aega && _cc.Count(_cc.TrimPrefix(_bdab, "\u0078\u0072\u0065\u0066"), "\u0020") > 0 {
		_dfad._eggc._bgb = true
	}
	_eb.Log.Trace("\u0078\u0072\u0065\u0066 f\u0069\u0072\u0073\u0074\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _bdab)
	_edeb := -1
	_cfcd := 0
	_egdb := false
	_becb := ""
	for {
		_dfad.skipSpaces()
		_, _abdd := _dfad._dedbc.Peek(1)
		if _abdd != nil {
			return nil, _abdd
		}
		_bdab, _abdd = _dfad.readTextLine()
		if _abdd != nil {
			return nil, _abdd
		}
		_acegb := _bcc.FindStringSubmatch(_bdab)
		if len(_acegb) == 0 {
			_cffcf := len(_becb) > 0
			_becb += _bdab + "\u000a"
			if _cffcf {
				_acegb = _bcc.FindStringSubmatch(_becb)
			}
		}
		if len(_acegb) == 3 {
			if _dfad._aega && !_dfad._eggc._bgf {
				var (
					_gabg bool
					_eedd int
				)
				for _, _bbce := range _bdab {
					if _gb.IsDigit(_bbce) {
						if _gabg {
							break
						}
						continue
					}
					if !_gabg {
						_gabg = true
					}
					_eedd++
				}
				if _eedd > 1 {
					_dfad._eggc._bgf = true
				}
			}
			_defbc, _ := _d.Atoi(_acegb[1])
			_befg, _ := _d.Atoi(_acegb[2])
			_edeb = _defbc
			_cfcd = _befg
			_egdb = true
			_becb = ""
			_eb.Log.Trace("\u0078r\u0065\u0066 \u0073\u0075\u0062s\u0065\u0063\u0074\u0069\u006f\u006e\u003a \u0066\u0069\u0072\u0073\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0064\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _edeb, _cfcd)
			continue
		}
		_gbbd := _abeba.FindStringSubmatch(_bdab)
		if len(_gbbd) == 4 {
			if !_egdb {
				_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0058r\u0065\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006fr\u006da\u0074\u0021\u000a")
				return nil, _f.New("\u0078\u0072\u0065\u0066 i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_fafg, _ := _d.ParseInt(_gbbd[1], 10, 64)
			_adaac, _ := _d.Atoi(_gbbd[2])
			_fdca := _gbbd[3]
			_becb = ""
			if _cc.ToLower(_fdca) == "\u006e" && _fafg > 1 {
				_cggba, _bbgb := _dfad._bfba.ObjectMap[_edeb]
				if !_bbgb || _adaac > _cggba.Generation {
					_bcgd := XrefObject{ObjectNumber: _edeb, XType: XrefTypeTableEntry, Offset: _fafg, Generation: _adaac}
					_dfad._bfba.ObjectMap[_edeb] = _bcgd
				}
			}
			_edeb++
			continue
		}
		if (len(_bdab) > 6) && (_bdab[:7] == "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_eb.Log.Trace("\u0046o\u0075n\u0064\u0020\u0074\u0072\u0061i\u006c\u0065r\u0020\u002d\u0020\u0025\u0073", _bdab)
			if len(_bdab) > 9 {
				_fggca := _dfad.GetFileOffset()
				_dfad.SetFileOffset(_fggca - int64(len(_bdab)) + 7)
			}
			_dfad.skipSpaces()
			_dfad.skipComments()
			_eb.Log.Trace("R\u0065\u0061\u0064\u0069ng\u0020t\u0072\u0061\u0069\u006c\u0065r\u0020\u0064\u0069\u0063\u0074\u0021")
			_eb.Log.Trace("\u0070\u0065\u0065\u006b\u003a\u0020\u0022\u0025\u0073\u0022", _bdab)
			_daeab, _abdd = _dfad.ParseDict()
			_eb.Log.Trace("\u0045O\u0046\u0020\u0072\u0065a\u0064\u0069\u006e\u0067\u0020t\u0072a\u0069l\u0065\u0072\u0020\u0064\u0069\u0063\u0074!")
			if _abdd != nil {
				_eb.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _abdd)
				return nil, _abdd
			}
			break
		}
		if _bdab == "\u0025\u0025\u0045O\u0046" {
			_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006e\u0064 \u006f\u0066\u0020\u0066\u0069\u006c\u0065 -\u0020\u0074\u0072\u0061i\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066ou\u006e\u0064 \u002d\u0020\u0065\u0072\u0072\u006f\u0072\u0021")
			return nil, _f.New("\u0065\u006e\u0064 \u006f\u0066\u0020\u0066i\u006c\u0065\u0020\u002d\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_eb.Log.Trace("\u0078\u0072\u0065\u0066\u0020\u006d\u006f\u0072\u0065 \u003a\u0020\u0025\u0073", _bdab)
	}
	_eb.Log.Trace("\u0045\u004f\u0046 p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0021")
	if _dfad._gfggg == nil {
		_begde := XrefTypeTableEntry
		_dfad._gfggg = &_begde
	}
	return _daeab, nil
}

// EncodeStream encodes the stream data using the encoded specified by the stream's dictionary.
func EncodeStream(streamObj *PdfObjectStream) error {
	_eb.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_dceae, _gccce := NewEncoderFromStream(streamObj)
	if _gccce != nil {
		_eb.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _gccce)
		return _gccce
	}
	if _gcbc, _defee := _dceae.(*LZWEncoder); _defee {
		_gcbc.EarlyChange = 0
		streamObj.PdfObjectDictionary.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(0))
	}
	_eb.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u000a", _dceae)
	_gedf, _gccce := _dceae.EncodeBytes(streamObj.Stream)
	if _gccce != nil {
		_eb.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _gccce)
		return _gccce
	}
	streamObj.Stream = _gedf
	streamObj.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_gedf))))
	return nil
}

// DecodeBytes decodes byte array with ASCII85. 5 ASCII characters -> 4 raw binary bytes
func (_cega *ASCII85Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _ece []byte
	_eb.Log.Trace("\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0020\u0044e\u0063\u006f\u0064\u0065")
	_aada := 0
	_cgab := false
	for _aada < len(encoded) && !_cgab {
		_gbef := [5]byte{0, 0, 0, 0, 0}
		_dgfb := 0
		_bgbd := 0
		_gbeg := 4
		for _bgbd < 5+_dgfb {
			if _aada+_bgbd == len(encoded) {
				break
			}
			_effd := encoded[_aada+_bgbd]
			if IsWhiteSpace(_effd) {
				_dgfb++
				_bgbd++
				continue
			} else if _effd == '~' && _aada+_bgbd+1 < len(encoded) && encoded[_aada+_bgbd+1] == '>' {
				_gbeg = (_bgbd - _dgfb) - 1
				if _gbeg < 0 {
					_gbeg = 0
				}
				_cgab = true
				break
			} else if _effd >= '!' && _effd <= 'u' {
				_effd -= '!'
			} else if _effd == 'z' && _bgbd-_dgfb == 0 {
				_gbeg = 4
				_bgbd++
				break
			} else {
				_eb.Log.Error("\u0046\u0061i\u006c\u0065\u0064\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020co\u0064\u0065")
				return nil, _f.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u0064\u0065\u0020e\u006e\u0063\u006f\u0075\u006e\u0074\u0065\u0072\u0065\u0064")
			}
			_gbef[_bgbd-_dgfb] = _effd
			_bgbd++
		}
		_aada += _bgbd
		for _ecfgg := _gbeg + 1; _ecfgg < 5; _ecfgg++ {
			_gbef[_ecfgg] = 84
		}
		_afag := uint32(_gbef[0])*85*85*85*85 + uint32(_gbef[1])*85*85*85 + uint32(_gbef[2])*85*85 + uint32(_gbef[3])*85 + uint32(_gbef[4])
		_aafe := []byte{byte((_afag >> 24) & 0xff), byte((_afag >> 16) & 0xff), byte((_afag >> 8) & 0xff), byte(_afag & 0xff)}
		_ece = append(_ece, _aafe[:_gbeg]...)
	}
	_eb.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0065\u006e\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", encoded)
	_eb.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0064\u0065\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", _ece)
	return _ece, nil
}

type objectStreams map[int]objectStream

// GetIndirect returns the *PdfIndirectObject represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetIndirect(obj PdfObject) (_abgac *PdfIndirectObject, _gcbe bool) {
	obj = ResolveReference(obj)
	_abgac, _gcbe = obj.(*PdfIndirectObject)
	return _abgac, _gcbe
}

// PdfObjectNull represents the primitive PDF null object.
type PdfObjectNull struct{}

// ToIntegerArray returns a slice of all array elements as an int slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_eedf *PdfObjectArray) ToIntegerArray() ([]int, error) {
	var _eeae []int
	for _, _geef := range _eedf.Elements() {
		if _cgffa, _eegdg := _geef.(*PdfObjectInteger); _eegdg {
			_eeae = append(_eeae, int(*_cgffa))
		} else {
			return nil, ErrTypeError
		}
	}
	return _eeae, nil
}
func _bagf(_gab *PdfObjectStream, _fedf *PdfObjectDictionary) (*LZWEncoder, error) {
	_fafb := NewLZWEncoder()
	_ddc := _gab.PdfObjectDictionary
	if _ddc == nil {
		return _fafb, nil
	}
	if _fedf == nil {
		_eaf := TraceToDirectObject(_ddc.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _eaf != nil {
			if _ffcg, _eae := _eaf.(*PdfObjectDictionary); _eae {
				_fedf = _ffcg
			} else if _dbfg, _ggbg := _eaf.(*PdfObjectArray); _ggbg {
				if _dbfg.Len() == 1 {
					if _cggab, _cgf := GetDict(_dbfg.Get(0)); _cgf {
						_fedf = _cggab
					}
				}
			}
			if _fedf == nil {
				_eb.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _eaf)
				return nil, _ee.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
	}
	_bddg := _ddc.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _bddg != nil {
		_dede, _gbac := _bddg.(*PdfObjectInteger)
		if !_gbac {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _bddg)
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_dede != 0 && *_dede != 1 {
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_fafb.EarlyChange = int(*_dede)
	} else {
		_fafb.EarlyChange = 1
	}
	if _fedf == nil {
		return _fafb, nil
	}
	if _cagd, _affe := GetIntVal(_fedf.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")); _affe {
		if _cagd == 0 || _cagd == 1 {
			_fafb.EarlyChange = _cagd
		} else {
			_eb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020E\u0061\u0072\u006c\u0079\u0043\u0068\u0061n\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020%\u0064", _cagd)
		}
	}
	_bddg = _fedf.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _bddg != nil {
		_baeb, _gcdd := _bddg.(*PdfObjectInteger)
		if !_gcdd {
			_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _bddg)
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_fafb.Predictor = int(*_baeb)
	}
	_bddg = _fedf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bddg != nil {
		_afga, _edda := _bddg.(*PdfObjectInteger)
		if !_edda {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_fafb.BitsPerComponent = int(*_afga)
	}
	if _fafb.Predictor > 1 {
		_fafb.Columns = 1
		_bddg = _fedf.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _bddg != nil {
			_caeb, _gabd := _bddg.(*PdfObjectInteger)
			if !_gabd {
				return nil, _ee.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_fafb.Columns = int(*_caeb)
		}
		_fafb.Colors = 1
		_bddg = _fedf.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _bddg != nil {
			_bgad, _cde := _bddg.(*PdfObjectInteger)
			if !_cde {
				return nil, _ee.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_fafb.Colors = int(*_bgad)
		}
	}
	_eb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _fedf.String())
	return _fafb, nil
}

type encryptDict struct {
	Filter    string
	V         int
	SubFilter string
	Length    int
	StmF      string
	StrF      string
	EFF       string
	CF        map[string]_gbe.FilterDict
}

var _ecfbb = _ba.MustCompile("\u005b\\\u0072\u005c\u006e\u005d\u005c\u0073\u002a\u0028\u0078\u0072\u0065f\u0029\u005c\u0073\u002a\u005b\u005c\u0072\u005c\u006e\u005d")

func (_bgde *PdfObjectDictionary) setWithLock(_ccdfd PdfObjectName, _dgec PdfObject, _aeeff bool) {
	if _aeeff {
		_bgde._bfdf.Lock()
		defer _bgde._bfdf.Unlock()
	}
	_, _ccac := _bgde._abec[_ccdfd]
	if !_ccac {
		_bgde._begbb = append(_bgde._begbb, _ccdfd)
	}
	_bgde._abec[_ccdfd] = _dgec
}
func (_cfeab *PdfParser) repairRebuildXrefsTopDown() (*XrefTable, error) {
	if _cfeab._egcga {
		return nil, _ee.Errorf("\u0072\u0065\u0070\u0061\u0069\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_cfeab._egcga = true
	_cfeab._cdea.Seek(0, _bb.SeekStart)
	_cfeab._dedbc = _ga.NewReader(_cfeab._cdea)
	_ecbfe := 20
	_dcag := make([]byte, _ecbfe)
	_badbb := XrefTable{}
	_badbb.ObjectMap = make(map[int]XrefObject)
	for {
		_gbafb, _cadb := _cfeab._dedbc.ReadByte()
		if _cadb != nil {
			if _cadb == _bb.EOF {
				break
			} else {
				return nil, _cadb
			}
		}
		if _gbafb == 'j' && _dcag[_ecbfe-1] == 'b' && _dcag[_ecbfe-2] == 'o' && IsWhiteSpace(_dcag[_ecbfe-3]) {
			_dfdca := _ecbfe - 4
			for IsWhiteSpace(_dcag[_dfdca]) && _dfdca > 0 {
				_dfdca--
			}
			if _dfdca == 0 || !IsDecimalDigit(_dcag[_dfdca]) {
				continue
			}
			for IsDecimalDigit(_dcag[_dfdca]) && _dfdca > 0 {
				_dfdca--
			}
			if _dfdca == 0 || !IsWhiteSpace(_dcag[_dfdca]) {
				continue
			}
			for IsWhiteSpace(_dcag[_dfdca]) && _dfdca > 0 {
				_dfdca--
			}
			if _dfdca == 0 || !IsDecimalDigit(_dcag[_dfdca]) {
				continue
			}
			for IsDecimalDigit(_dcag[_dfdca]) && _dfdca > 0 {
				_dfdca--
			}
			if _dfdca == 0 {
				continue
			}
			_dabd := _cfeab.GetFileOffset() - int64(_ecbfe-_dfdca)
			_ggce := append(_dcag[_dfdca+1:], _gbafb)
			_fbfb, _fbfgc, _bdebc := _cfdbd(string(_ggce))
			if _bdebc != nil {
				_eb.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u003a\u0020\u0025\u0076", _bdebc)
				return nil, _bdebc
			}
			if _edad, _ggabg := _badbb.ObjectMap[_fbfb]; !_ggabg || _edad.Generation < _fbfgc {
				_facf := XrefObject{}
				_facf.XType = XrefTypeTableEntry
				_facf.ObjectNumber = _fbfb
				_facf.Generation = _fbfgc
				_facf.Offset = _dabd
				_badbb.ObjectMap[_fbfb] = _facf
			}
		}
		_dcag = append(_dcag[1:_ecbfe], _gbafb)
	}
	_cfeab._cefg = nil
	return &_badbb, nil
}

// GetXrefType returns the type of the first xref object (table or stream).
func (_cfb *PdfParser) GetXrefType() *xrefType { return _cfb._gfggg }

// String returns a descriptive information string about the encryption method used.
func (_cggd *PdfCrypt) String() string {
	if _cggd == nil {
		return ""
	}
	_deb := _cggd._ffe.Filter + "\u0020\u002d\u0020"
	if _cggd._ffe.V == 0 {
		_deb += "\u0055\u006e\u0064\u006fcu\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0061\u006c\u0067\u006f\u0072\u0069\u0074h\u006d"
	} else if _cggd._ffe.V == 1 {
		_deb += "\u0052\u0043\u0034:\u0020\u0034\u0030\u0020\u0062\u0069\u0074\u0073"
	} else if _cggd._ffe.V == 2 {
		_deb += _ee.Sprintf("\u0052\u0043\u0034:\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073", _cggd._ffe.Length)
	} else if _cggd._ffe.V == 3 {
		_deb += "U\u006e\u0070\u0075\u0062li\u0073h\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d"
	} else if _cggd._ffe.V >= 4 {
		_deb += _ee.Sprintf("\u0053\u0074r\u0065\u0061\u006d\u0020f\u0069\u006ct\u0065\u0072\u003a\u0020\u0025\u0073\u0020\u002d \u0053\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074\u0065r\u003a\u0020\u0025\u0073", _cggd._badb, _cggd._df)
		_deb += "\u003b\u0020C\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0073\u003a"
		for _abgg, _dae := range _cggd._ace {
			_deb += _ee.Sprintf("\u0020\u002d\u0020\u0025\u0073\u003a\u0020\u0025\u0073 \u0028\u0025\u0064\u0029", _abgg, _dae.Name(), _dae.KeyLength())
		}
	}
	_dbg := _cggd.GetAccessPermissions()
	_deb += _ee.Sprintf("\u0020\u002d\u0020\u0025\u0023\u0076", _dbg)
	return _deb
}
func (_beea *PdfParser) repairSeekXrefMarker() error {
	_gfee, _ebca := _beea._cdea.Seek(0, _bb.SeekEnd)
	if _ebca != nil {
		return _ebca
	}
	_adfgf := _ba.MustCompile("\u005cs\u0078\u0072\u0065\u0066\u005c\u0073*")
	var _dfac int64
	var _bccg int64 = 1000
	for _dfac < _gfee {
		if _gfee <= (_bccg + _dfac) {
			_bccg = _gfee - _dfac
		}
		_, _gaafc := _beea._cdea.Seek(-_dfac-_bccg, _bb.SeekEnd)
		if _gaafc != nil {
			return _gaafc
		}
		_fdfdd := make([]byte, _bccg)
		_beea._cdea.Read(_fdfdd)
		_eb.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0078\u0072\u0065\u0066\u0020\u003a\u0020\u0022\u0025\u0073\u0022", string(_fdfdd))
		_abgb := _adfgf.FindAllStringIndex(string(_fdfdd), -1)
		if _abgb != nil {
			_ebea := _abgb[len(_abgb)-1]
			_eb.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _abgb)
			_beea._cdea.Seek(-_dfac-_bccg+int64(_ebea[0]), _bb.SeekEnd)
			_beea._dedbc = _ga.NewReader(_beea._cdea)
			for {
				_dacb, _efagd := _beea._dedbc.Peek(1)
				if _efagd != nil {
					return _efagd
				}
				_eb.Log.Trace("\u0042\u003a\u0020\u0025\u0064\u0020\u0025\u0063", _dacb[0], _dacb[0])
				if !IsWhiteSpace(_dacb[0]) {
					break
				}
				_beea._dedbc.Discard(1)
			}
			return nil
		}
		_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_dfac += _bccg
	}
	_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074a\u0062\u006c\u0065\u0020\u006d\u0061r\u006b\u0065\u0072\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u002e")
	return _f.New("\u0078r\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020")
}

var _abeba = _ba.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064+\u0029\u005c\u0073\u002b\u0028\u005b\u006e\u0066\u005d\u0029\\\u0073\u002a\u0024")

// Read implementation of Read interface.
func (_fdfdc *limitedReadSeeker) Read(p []byte) (_baaag int, _gfeb error) {
	_dccgc, _gfeb := _fdfdc._efag.Seek(0, _bb.SeekCurrent)
	if _gfeb != nil {
		return 0, _gfeb
	}
	_accg := _fdfdc._befa - _dccgc
	if _accg == 0 {
		return 0, _bb.EOF
	}
	if _dgfd := int64(len(p)); _dgfd < _accg {
		_accg = _dgfd
	}
	_adea := make([]byte, _accg)
	_baaag, _gfeb = _fdfdc._efag.Read(_adea)
	copy(p, _adea)
	return _baaag, _gfeb
}

// MakeDict creates and returns an empty PdfObjectDictionary.
func MakeDict() *PdfObjectDictionary {
	_abgab := &PdfObjectDictionary{}
	_abgab._abec = map[PdfObjectName]PdfObject{}
	_abgab._begbb = []PdfObjectName{}
	_abgab._bfdf = &_g.Mutex{}
	return _abgab
}

// EncodeBytes ASCII encodes the passed in slice of bytes.
func (_edgf *ASCIIHexEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _bfcab _bg.Buffer
	for _, _dddf := range data {
		_bfcab.WriteString(_ee.Sprintf("\u0025\u002e\u0032X\u0020", _dddf))
	}
	_bfcab.WriteByte('>')
	return _bfcab.Bytes(), nil
}

// DecodeStream decodes the stream containing CCITTFax encoded image data.
func (_caea *CCITTFaxEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _caea.DecodeBytes(streamObj.Stream)
}

// String returns a string describing `stream`.
func (_ebfe *PdfObjectStream) String() string {
	return _ee.Sprintf("O\u0062j\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0025\u0064: \u0025\u0073", _ebfe.ObjectNumber, _ebfe.PdfObjectDictionary)
}

// GetBoolVal returns the bool value within a *PdObjectBool represented by an PdfObject interface directly or indirectly.
// If the PdfObject does not represent a bool value, a default value of false is returned (found = false also).
func GetBoolVal(obj PdfObject) (_fdaf bool, _fggf bool) {
	_dbcg, _fggf := TraceToDirectObject(obj).(*PdfObjectBool)
	if _fggf {
		return bool(*_dbcg), true
	}
	return false, false
}
func _efbb(_ffgd *PdfObjectStream, _cegd *MultiEncoder) (*DCTEncoder, error) {
	_dafb := NewDCTEncoder()
	_acae := _ffgd.PdfObjectDictionary
	if _acae == nil {
		return _dafb, nil
	}
	_ddde := _ffgd.Stream
	if _cegd != nil {
		_aedd, _dbcd := _cegd.DecodeBytes(_ddde)
		if _dbcd != nil {
			return nil, _dbcd
		}
		_ddde = _aedd
	}
	_eadf := _bg.NewReader(_ddde)
	_ccd, _gcff := _beg.DecodeConfig(_eadf)
	if _gcff != nil {
		_eb.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _gcff)
		return nil, _gcff
	}
	switch _ccd.ColorModel {
	case _cb.RGBAModel:
		_dafb.BitsPerComponent = 8
		_dafb.ColorComponents = 3
		_dafb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _cb.RGBA64Model:
		_dafb.BitsPerComponent = 16
		_dafb.ColorComponents = 3
		_dafb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _cb.GrayModel:
		_dafb.BitsPerComponent = 8
		_dafb.ColorComponents = 1
		_dafb.Decode = []float64{0.0, 1.0}
	case _cb.Gray16Model:
		_dafb.BitsPerComponent = 16
		_dafb.ColorComponents = 1
		_dafb.Decode = []float64{0.0, 1.0}
	case _cb.CMYKModel:
		_dafb.BitsPerComponent = 8
		_dafb.ColorComponents = 4
		_dafb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _cb.YCbCrModel:
		_dafb.BitsPerComponent = 8
		_dafb.ColorComponents = 3
		_dafb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	default:
		return nil, _f.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_dafb.Width = _ccd.Width
	_dafb.Height = _ccd.Height
	_eb.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _dafb)
	_dafb.Quality = DefaultJPEGQuality
	_cbbg, _fgc := GetArray(_acae.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _fgc {
		_ddcf, _gadf := _cbbg.ToFloat64Array()
		if _gadf != nil {
			return _dafb, _gadf
		}
		_dafb.Decode = _ddcf
	}
	return _dafb, nil
}

// GetFilterName returns the name of the encoding filter.
func (_bcg *FlateEncoder) GetFilterName() string { return StreamEncodingFilterNameFlate }

// CheckAccessRights checks access rights and permissions for a specified password. If either user/owner password is
// specified, full rights are granted, otherwise the access rights are specified by the Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_dagbd *PdfParser) CheckAccessRights(password []byte) (bool, _cbd.Permissions, error) {
	if _dagbd._eccc == nil {
		return true, _cbd.PermOwner, nil
	}
	return _dagbd._eccc.checkAccessRights(password)
}

// Get returns the i-th element of the array or nil if out of bounds (by index).
func (_edec *PdfObjectArray) Get(i int) PdfObject {
	if _edec == nil || i >= len(_edec._aagc) || i < 0 {
		return nil
	}
	return _edec._aagc[i]
}

// PdfCryptNewDecrypt makes the document crypt handler based on the encryption dictionary
// and trailer dictionary. Returns an error on failure to process.
func PdfCryptNewDecrypt(parser *PdfParser, ed, trailer *PdfObjectDictionary) (*PdfCrypt, error) {
	_ceed := &PdfCrypt{_agg: false, _ge: make(map[PdfObject]bool), _edc: make(map[PdfObject]bool), _eda: make(map[int]struct{}), _eadg: parser}
	_fefg, _fcac := ed.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectName)
	if !_fcac {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0043\u0072\u0079\u0070\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u0066\u0069\u0065\u006c\u0064\u0021")
		return _ceed, _f.New("r\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0063\u0072\u0079\u0070\u0074 \u0066\u0069\u0065\u006c\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
	}
	if *_fefg != "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064" {
		_eb.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020(%\u0073\u0029", *_fefg)
		return _ceed, _f.New("\u0075n\u0073u\u0070\u0070\u006f\u0072\u0074e\u0064\u0020F\u0069\u006c\u0074\u0065\u0072")
	}
	_ceed._ffe.Filter = string(*_fefg)
	if _cca, _bcae := ed.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r").(*PdfObjectString); _bcae {
		_ceed._ffe.SubFilter = _cca.Str()
		_eb.Log.Debug("\u0055s\u0069n\u0067\u0020\u0073\u0075\u0062f\u0069\u006ct\u0065\u0072\u0020\u0025\u0073", _cca)
	}
	if L, _dfc := ed.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _dfc {
		if (*L % 8) != 0 {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0065\u006ec\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
			return _ceed, _f.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0065\u006e\u0063\u0072y\u0070t\u0069o\u006e\u0020\u006c\u0065\u006e\u0067\u0074h")
		}
		_ceed._ffe.Length = int(*L)
	} else {
		_ceed._ffe.Length = 40
	}
	_ceed._ffe.V = 0
	if _gbeb, _dccd := ed.Get("\u0056").(*PdfObjectInteger); _dccd {
		V := int(*_gbeb)
		_ceed._ffe.V = V
		if V >= 1 && V <= 2 {
			_ceed._ace = _gcb(_ceed._ffe.Length)
		} else if V >= 4 && V <= 5 {
			if _bge := _ceed.loadCryptFilters(ed); _bge != nil {
				return _ceed, _bge
			}
		} else {
			_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0061lg\u006f\u0020\u0056 \u003d \u0025\u0064", V)
			return _ceed, _f.New("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d")
		}
	}
	if _debf := _gaf(&_ceed._adfe, ed); _debf != nil {
		return _ceed, _debf
	}
	_gdgd := ""
	if _baag, _cff := trailer.Get("\u0049\u0044").(*PdfObjectArray); _cff && _baag.Len() >= 1 {
		_bacb, _cdc := GetString(_baag.Get(0))
		if !_cdc {
			return _ceed, _f.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0074r\u0061\u0069l\u0065\u0072\u0020\u0049\u0044")
		}
		_gdgd = _bacb.Str()
	} else {
		_eb.Log.Debug("\u0054\u0072ai\u006c\u0065\u0072 \u0049\u0044\u0020\u0061rra\u0079 m\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0021")
	}
	_ceed._ffc = _gdgd
	return _ceed, nil
}
func (_efc *PdfCrypt) loadCryptFilters(_gdb *PdfObjectDictionary) error {
	_efc._ace = cryptFilters{}
	_ceb := _gdb.Get("\u0043\u0046")
	_ceb = TraceToDirectObject(_ceb)
	if _eba, _fcf := _ceb.(*PdfObjectReference); _fcf {
		_ggf, _bf := _efc._eadg.LookupByReference(*_eba)
		if _bf != nil {
			_eb.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0075\u0070\u0020\u0043\u0046\u0020\u0072\u0065\u0066\u0065\u0072en\u0063\u0065")
			return _bf
		}
		_ceb = TraceToDirectObject(_ggf)
	}
	_fbbg, _feg := _ceb.(*PdfObjectDictionary)
	if !_feg {
		_eb.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0043\u0046\u002c \u0074\u0079\u0070\u0065: \u0025\u0054", _ceb)
		return _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0046")
	}
	for _, _ebgd := range _fbbg.Keys() {
		_gbf := _fbbg.Get(_ebgd)
		if _ebce, _abae := _gbf.(*PdfObjectReference); _abae {
			_daf, _gddd := _efc._eadg.LookupByReference(*_ebce)
			if _gddd != nil {
				_eb.Log.Debug("\u0045\u0072ro\u0072\u0020\u006co\u006f\u006b\u0075\u0070 up\u0020di\u0063\u0074\u0069\u006f\u006e\u0061\u0072y \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065")
				return _gddd
			}
			_gbf = TraceToDirectObject(_daf)
		}
		_af, _cbfe := _gbf.(*PdfObjectDictionary)
		if !_cbfe {
			return _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006e \u0043\u0046\u0020\u0028\u006e\u0061\u006d\u0065\u0020\u0025\u0073\u0029\u0020-\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0062\u0075\u0074\u0020\u0025\u0054", _ebgd, _gbf)
		}
		if _ebgd == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u002d\u0020\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006f\u0076\u0065\u0072\u0077r\u0069\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0069d\u0065\u006e\u0074\u0069\u0074\u0079\u0020\u0066\u0069\u006c\u0074\u0065\u0072 \u002d\u0020\u0054\u0072\u0079\u0069n\u0067\u0020\u006ee\u0078\u0074")
			continue
		}
		var _fafe _gbe.FilterDict
		if _ecag := _gfad(&_fafe, _af); _ecag != nil {
			return _ecag
		}
		_abfg, _aaa := _gbe.NewFilter(_fafe)
		if _aaa != nil {
			return _aaa
		}
		_efc._ace[string(_ebgd)] = _abfg
	}
	_efc._ace["\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"] = _gbe.NewIdentity()
	_efc._df = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _dbc, _fegf := _gdb.Get("\u0053\u0074\u0072\u0046").(*PdfObjectName); _fegf {
		if _, _aed := _efc._ace[string(*_dbc)]; !_aed {
			return _ee.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u0072\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_dbc)
		}
		_efc._df = string(*_dbc)
	}
	_efc._badb = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _eec, _efd := _gdb.Get("\u0053\u0074\u006d\u0046").(*PdfObjectName); _efd {
		if _, _cdf := _efc._ace[string(*_eec)]; !_cdf {
			return _ee.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u006d\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_eec)
		}
		_efc._badb = string(*_eec)
	}
	return nil
}

// EncodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_bdec *RawEncoder) EncodeBytes(data []byte) ([]byte, error) { return data, nil }
func (_gfga *PdfParser) parseArray() (*PdfObjectArray, error) {
	_bgfe := MakeArray()
	_gfga._dedbc.ReadByte()
	for {
		_gfga.skipSpaces()
		_gaca, _fcdf := _gfga._dedbc.Peek(1)
		if _fcdf != nil {
			return _bgfe, _fcdf
		}
		if _gaca[0] == ']' {
			_gfga._dedbc.ReadByte()
			break
		}
		_gdec, _fcdf := _gfga.parseObject()
		if _fcdf != nil {
			return _bgfe, _fcdf
		}
		_bgfe.Append(_gdec)
	}
	return _bgfe, nil
}

// IsWhiteSpace checks if byte represents a white space character.
func IsWhiteSpace(ch byte) bool {
	if (ch == 0x00) || (ch == 0x09) || (ch == 0x0A) || (ch == 0x0C) || (ch == 0x0D) || (ch == 0x20) {
		return true
	}
	return false
}

// String returns the PDF version as a string. Implements interface fmt.Stringer.
func (_fgbc Version) String() string {
	return _ee.Sprintf("\u00250\u0064\u002e\u0025\u0030\u0064", _fgbc.Major, _fgbc.Minor)
}

// PdfObjectString represents the primitive PDF string object.
type PdfObjectString struct {
	_cbdg string
	_cdca bool
}

func (_adc *PdfParser) lookupByNumber(_ada int, _bcb bool) (PdfObject, bool, error) {
	_agc, _gc := _adc.ObjCache[_ada]
	if _gc {
		_eb.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0063a\u0063\u0068\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0025\u0064", _ada)
		return _agc, false, nil
	}
	if _adc._cefg == nil {
		_adc._cefg = map[int]bool{}
	}
	if _adc._cefg[_ada] {
		_eb.Log.Debug("ER\u0052\u004f\u0052\u003a\u0020\u004c\u006fok\u0075\u0070\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0069\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0069\u006e\u0020\u0070\u0072\u006f\u0067\u0072\u0065\u0073\u0073\u0020\u002d\u0020\u0072\u0065c\u0075\u0072\u0073\u0069\u0076\u0065 \u006c\u006f\u006f\u006b\u0075\u0070\u0020\u0061\u0074t\u0065m\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064", _ada)
		return nil, false, _f.New("\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006f\u006f\u006b\u0075p\u0020a\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064")
	}
	_adc._cefg[_ada] = true
	defer delete(_adc._cefg, _ada)
	_gad, _gc := _adc._bfba.ObjectMap[_ada]
	if !_gc {
		_eb.Log.Trace("\u0055\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u006c\u006f\u0063\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006e\u0020\u0078\u0072\u0065\u0066\u0073\u0021 \u002d\u0020\u0052\u0065\u0074u\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		var _egb PdfObjectNull
		return &_egb, false, nil
	}
	_eb.Log.Trace("L\u006fo\u006b\u0075\u0070\u0020\u006f\u0062\u006a\u0020n\u0075\u006d\u0062\u0065r \u0025\u0064", _ada)
	if _gad.XType == XrefTypeTableEntry {
		_eb.Log.Trace("\u0078r\u0065f\u006f\u0062\u006a\u0020\u006fb\u006a\u0020n\u0075\u006d\u0020\u0025\u0064", _gad.ObjectNumber)
		_eb.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u0067e\u006e\u0020\u0025\u0064", _gad.Generation)
		_eb.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0064", _gad.Offset)
		_adc._cdea.Seek(_gad.Offset, _bb.SeekStart)
		_adc._dedbc = _ga.NewReader(_adc._cdea)
		_ff, _ecg := _adc.ParseIndirectObject()
		if _ecg != nil {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006ce\u0064\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029", _ecg)
			if _bcb {
				_eb.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0072\u0065\u0070\u0061\u0069r\u0020\u0078\u0072\u0065\u0066\u0073\u0020\u0028\u0074\u006f\u0070\u0020d\u006f\u0077\u006e\u0029")
				_cad, _dbb := _adc.repairRebuildXrefsTopDown()
				if _dbb != nil {
					_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020r\u0065\u0070\u0061\u0069\u0072\u0020\u0028\u0025\u0073\u0029", _dbb)
					return nil, false, _dbb
				}
				_adc._bfba = *_cad
				return _adc.lookupByNumber(_ada, false)
			}
			return nil, false, _ecg
		}
		if _bcb {
			_baa, _, _ := _ed(_ff)
			if int(_baa) != _ada {
				_eb.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0078\u0072\u0065\u0066s\u003a \u0052e\u0062\u0075\u0069\u006c\u0064\u0069\u006eg")
				_adf := _adc.rebuildXrefTable()
				if _adf != nil {
					return nil, false, _adf
				}
				_adc.ObjCache = objectCache{}
				return _adc.lookupByNumberWrapper(_ada, false)
			}
		}
		_eb.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006f\u0062\u006a")
		_adc.ObjCache[_ada] = _ff
		return _ff, false, nil
	} else if _gad.XType == XrefTypeObjectStream {
		_eb.Log.Trace("\u0078r\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0021")
		_eb.Log.Trace("\u003e\u004c\u006f\u0061\u0064\u0020\u0076\u0069\u0061\u0020\u004f\u0053\u0021")
		_eb.Log.Trace("\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0069\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020%\u0064\u002f\u0025\u0064", _gad.OsObjNumber, _gad.OsObjIndex)
		if _gad.OsObjNumber == _ada {
			_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0043i\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065f\u0065\u0072\u0065n\u0063e\u0021\u003f\u0021")
			return nil, true, _f.New("\u0078\u0072\u0065f \u0063\u0069\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _, _fc := _adc._bfba.ObjectMap[_gad.OsObjNumber]; _fc {
			_dgb, _bca := _adc.lookupObjectViaOS(_gad.OsObjNumber, _ada)
			if _bca != nil {
				_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069n\u0067\u0020\u0045\u0052\u0052\u0020\u0028\u0025\u0073\u0029", _bca)
				return nil, true, _bca
			}
			_eb.Log.Trace("\u003c\u004c\u006f\u0061\u0064\u0065\u0064\u0020\u0076i\u0061\u0020\u004f\u0053")
			_adc.ObjCache[_ada] = _dgb
			if _adc._eccc != nil {
				_adc._eccc._ge[_dgb] = true
			}
			return _dgb, true, nil
		}
		_eb.Log.Debug("\u003f\u003f\u0020\u0042\u0065\u006c\u006f\u006eg\u0073\u0020\u0074o \u0061\u0020\u006e\u006f\u006e\u002dc\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002e.\u002e\u0021")
		return nil, true, _f.New("\u006f\u0073\u0020\u0062\u0065\u006c\u006fn\u0067\u0073\u0020t\u006f\u0020\u0061\u0020n\u006f\u006e\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, false, _f.New("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0078\u0072\u0065\u0066 \u0074\u0079\u0070\u0065")
}

// String returns a string describing `array`.
func (_bagff *PdfObjectArray) String() string {
	_ddccd := "\u005b"
	for _cccf, _fegcc := range _bagff.Elements() {
		_ddccd += _fegcc.String()
		if _cccf < (_bagff.Len() - 1) {
			_ddccd += "\u002c\u0020"
		}
	}
	_ddccd += "\u005d"
	return _ddccd
}
func _eaeg() string { return _eb.Version }

// PdfObjectInteger represents the primitive PDF integer numerical object.
type PdfObjectInteger int64

// Clear resets the dictionary to an empty state.
func (_beaed *PdfObjectDictionary) Clear() {
	_beaed._begbb = []PdfObjectName{}
	_beaed._abec = map[PdfObjectName]PdfObject{}
	_beaed._bfdf = &_g.Mutex{}
}

const _dedd = "\u0053\u0074\u0064C\u0046"

// MakeObjectStreams creates an PdfObjectStreams from a list of PdfObjects.
func MakeObjectStreams(objects ...PdfObject) *PdfObjectStreams {
	return &PdfObjectStreams{_fagda: objects}
}

// PdfObjectName represents the primitive PDF name object.
type PdfObjectName string

func (_geffg *PdfParser) parseNumber() (PdfObject, error) { return ParseNumber(_geffg._dedbc) }

// GetXrefOffset returns the offset of the xref table.
func (_dacd *PdfParser) GetXrefOffset() int64 { return _dacd._cgfg }

// SetImage sets the image base for given flate encoder.
func (_aad *FlateEncoder) SetImage(img *_eeb.ImageBase) { _aad._gaade = img }

// EncodeBytes encodes the image data using either Group3 or Group4 CCITT facsimile (fax) encoding.
// `data` is expected to be 1 color component, 1 bit per component. It is also valid to provide 8 BPC, 1 CC image like
// a standard go image Gray data.
func (_becda *CCITTFaxEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _dcab _eeb.Gray
	switch len(data) {
	case _becda.Rows * _becda.Columns:
		_ega, _egbfd := _eeb.NewImage(_becda.Columns, _becda.Rows, 8, 1, data, nil, nil)
		if _egbfd != nil {
			return nil, _egbfd
		}
		_dcab = _ega.(_eeb.Gray)
	case (_becda.Columns * _becda.Rows) + 7>>3:
		_baef, _fag := _eeb.NewImage(_becda.Columns, _becda.Rows, 1, 1, data, nil, nil)
		if _fag != nil {
			return nil, _fag
		}
		_dfcd := _baef.(*_eeb.Monochrome)
		if _fag = _dfcd.AddPadding(); _fag != nil {
			return nil, _fag
		}
		_dcab = _dfcd
	default:
		if len(data) < _eeb.BytesPerLine(_becda.Columns, 1, 1)*_becda.Rows {
			return nil, _f.New("p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020i\u006e\u0070\u0075t\u0020d\u0061\u0074\u0061")
		}
		_cegc, _ccgd := _eeb.NewImage(_becda.Columns, _becda.Rows, 1, 1, data, nil, nil)
		if _ccgd != nil {
			return nil, _ccgd
		}
		_gdbd := _cegc.(*_eeb.Monochrome)
		_dcab = _gdbd
	}
	_dfee := make([][]byte, _becda.Rows)
	for _eeea := 0; _eeea < _becda.Rows; _eeea++ {
		_dbeg := make([]byte, _becda.Columns)
		for _ebaf := 0; _ebaf < _becda.Columns; _ebaf++ {
			_bfee := _dcab.GrayAt(_ebaf, _eeea)
			_dbeg[_ebaf] = _bfee.Y >> 7
		}
		_dfee[_eeea] = _dbeg
	}
	_feba := &_bad.Encoder{K: _becda.K, Columns: _becda.Columns, EndOfLine: _becda.EndOfLine, EndOfBlock: _becda.EndOfBlock, BlackIs1: _becda.BlackIs1, DamagedRowsBeforeError: _becda.DamagedRowsBeforeError, Rows: _becda.Rows, EncodedByteAlign: _becda.EncodedByteAlign}
	return _feba.Encode(_dfee), nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_daca *CCITTFaxEncoder) MakeStreamDict() *PdfObjectDictionary {
	_egbg := MakeDict()
	_egbg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_daca.GetFilterName()))
	_egbg.SetIfNotNil("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _daca.MakeDecodeParams())
	return _egbg
}

// UpdateParams updates the parameter values of the encoder.
func (_cadf *RunLengthEncoder) UpdateParams(params *PdfObjectDictionary) {}

// SetPredictor sets the predictor function.  Specify the number of columns per row.
// The columns indicates the number of samples per row.
// Used for grouping data together for compression.
func (_cafg *FlateEncoder) SetPredictor(columns int) { _cafg.Predictor = 11; _cafg.Columns = columns }
func (_gced *PdfParser) readComment() (string, error) {
	var _aabbc _bg.Buffer
	_, _ccfg := _gced.skipSpaces()
	if _ccfg != nil {
		return _aabbc.String(), _ccfg
	}
	_efgc := true
	for {
		_adff, _efbbb := _gced._dedbc.Peek(1)
		if _efbbb != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _efbbb.Error())
			return _aabbc.String(), _efbbb
		}
		if _efgc && _adff[0] != '%' {
			return _aabbc.String(), _f.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_efgc = false
		if (_adff[0] != '\r') && (_adff[0] != '\n') {
			_dfa, _ := _gced._dedbc.ReadByte()
			_aabbc.WriteByte(_dfa)
		} else {
			break
		}
	}
	return _aabbc.String(), nil
}

const _cafga = 6

func (_bgae *offsetReader) Seek(offset int64, whence int) (int64, error) {
	if whence == _bb.SeekStart {
		offset += _bgae._cbca
	}
	_ffdf, _aaab := _bgae._bfdd.Seek(offset, whence)
	if _aaab != nil {
		return _ffdf, _aaab
	}
	if whence == _bb.SeekCurrent {
		_ffdf -= _bgae._cbca
	}
	if _ffdf < 0 {
		return 0, _f.New("\u0063\u006f\u0072\u0065\u002eo\u0066\u0066\u0073\u0065\u0074\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053e\u0065\u006b\u003a\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	return _ffdf, nil
}

// GetRevisionNumber returns the current version of the Pdf document.
func (_bddgb *PdfParser) GetRevisionNumber() int { return _bddgb._aegg }

// PdfObjectFloat represents the primitive PDF floating point numerical object.
type PdfObjectFloat float64

// LookupByNumber looks up a PdfObject by object number.  Returns an error on failure.
func (_abf *PdfParser) LookupByNumber(objNumber int) (PdfObject, error) {
	_ad, _, _adb := _abf.lookupByNumberWrapper(objNumber, true)
	return _ad, _adb
}

// UpdateParams updates the parameter values of the encoder.
func (_cgbb *ASCII85Encoder) UpdateParams(params *PdfObjectDictionary) {}

// DecodeImages decodes the page images from the jbig2 'encoded' data input.
// The jbig2 document may contain multiple pages, thus the function can return multiple
// images. The images order corresponds to the page number.
func (_gfgd *JBIG2Encoder) DecodeImages(encoded []byte) ([]_a.Image, error) {
	const _cba = "\u004aB\u0049\u0047\u0032\u0045n\u0063\u006f\u0064\u0065\u0072.\u0044e\u0063o\u0064\u0065\u0049\u006d\u0061\u0067\u0065s"
	_cgff, _gda := _gae.Decode(encoded, _gae.Parameters{}, _gfgd.Globals.ToDocumentGlobals())
	if _gda != nil {
		return nil, _eca.Wrap(_gda, _cba, "")
	}
	_bbad, _gda := _cgff.PageNumber()
	if _gda != nil {
		return nil, _eca.Wrap(_gda, _cba, "")
	}
	_afed := []_a.Image{}
	var _bege _a.Image
	for _egab := 1; _egab <= _bbad; _egab++ {
		_bege, _gda = _cgff.DecodePageImage(_egab)
		if _gda != nil {
			return nil, _eca.Wrapf(_gda, _cba, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _egab)
		}
		_afed = append(_afed, _bege)
	}
	return _afed, nil
}

// MakeHexString creates an PdfObjectString from a string intended for output as a hexadecimal string.
func MakeHexString(s string) *PdfObjectString {
	_fdcb := PdfObjectString{_cbdg: s, _cdca: true}
	return &_fdcb
}

// ResolveReferencesDeep recursively traverses through object `o`, looking up and replacing
// references with indirect objects.
// Optionally a map of already deep-resolved objects can be provided via `traversed`. The `traversed` map
// is updated while traversing the objects to avoid traversing same objects multiple times.
func ResolveReferencesDeep(o PdfObject, traversed map[PdfObject]struct{}) error {
	if traversed == nil {
		traversed = map[PdfObject]struct{}{}
	}
	return _ccbb(o, 0, traversed)
}

// GetBool returns the *PdfObjectBool object that is represented by a PdfObject directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetBool(obj PdfObject) (_aebe *PdfObjectBool, _abddg bool) {
	_aebe, _abddg = TraceToDirectObject(obj).(*PdfObjectBool)
	return _aebe, _abddg
}

// WriteString outputs the object as it is to be written to file.
func (_ccgdb *PdfObjectStream) WriteString() string {
	var _ecec _cc.Builder
	_ecec.WriteString(_d.FormatInt(_ccgdb.ObjectNumber, 10))
	_ecec.WriteString("\u0020\u0030\u0020\u0052")
	return _ecec.String()
}

// DecodeBytes decodes a multi-encoded slice of bytes by passing it through the
// DecodeBytes method of the underlying encoders.
func (_ffffb *MultiEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_dcfcg := encoded
	var _egcg error
	for _, _gegf := range _ffffb._abeab {
		_eb.Log.Trace("\u004du\u006c\u0074i\u0020\u0045\u006e\u0063o\u0064\u0065\u0072 \u0044\u0065\u0063\u006f\u0064\u0065\u003a\u0020\u0041pp\u006c\u0079\u0069n\u0067\u0020F\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0076 \u0025\u0054", _gegf, _gegf)
		_dcfcg, _egcg = _gegf.DecodeBytes(_dcfcg)
		if _egcg != nil {
			return nil, _egcg
		}
	}
	return _dcfcg, nil
}

// XrefTable represents the cross references in a PDF, i.e. the table of objects and information
// where to access within the PDF file.
type XrefTable struct {
	ObjectMap map[int]XrefObject
	_bd       []XrefObject
}

var _ffdaa = _ba.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

// GetName returns the *PdfObjectName represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetName(obj PdfObject) (_cgfcf *PdfObjectName, _bbcd bool) {
	_cgfcf, _bbcd = TraceToDirectObject(obj).(*PdfObjectName)
	return _cgfcf, _bbcd
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_ffdb *MultiEncoder) MakeStreamDict() *PdfObjectDictionary {
	_aeeb := MakeDict()
	_aeeb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", _ffdb.GetFilterArray())
	for _, _fbd := range _ffdb._abeab {
		_bdga := _fbd.MakeStreamDict()
		for _, _dceaa := range _bdga.Keys() {
			_gcebd := _bdga.Get(_dceaa)
			if _dceaa != "\u0046\u0069\u006c\u0074\u0065\u0072" && _dceaa != "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073" {
				_aeeb.Set(_dceaa, _gcebd)
			}
		}
	}
	_defb := _ffdb.MakeDecodeParams()
	if _defb != nil {
		_aeeb.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _defb)
	}
	return _aeeb
}

// HasDataAfterEOF checks if there is some data after EOF marker.
func (_fcfd ParserMetadata) HasDataAfterEOF() bool { return _fcfd._fcag }

// WriteString outputs the object as it is to be written to file.
func (_daaag *PdfObjectStreams) WriteString() string {
	var _fada _cc.Builder
	_fada.WriteString(_d.FormatInt(_daaag.ObjectNumber, 10))
	_fada.WriteString("\u0020\u0030\u0020\u0052")
	return _fada.String()
}

// ASCIIHexEncoder implements ASCII hex encoder/decoder.
type ASCIIHexEncoder struct{}

func _acag(_ddad _gbe.Filter, _cee _cbd.AuthEvent) *PdfObjectDictionary {
	if _cee == "" {
		_cee = _cbd.EventDocOpen
	}
	_aab := MakeDict()
	_aab.Set("\u0054\u0079\u0070\u0065", MakeName("C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072"))
	_aab.Set("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et", MakeName(string(_cee)))
	_aab.Set("\u0043\u0046\u004d", MakeName(_ddad.Name()))
	_aab.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_ddad.KeyLength())))
	return _aab
}

type xrefType int

// Remove removes an element specified by key.
func (_eeac *PdfObjectDictionary) Remove(key PdfObjectName) {
	_dagg := -1
	for _edfb, _aeggf := range _eeac._begbb {
		if _aeggf == key {
			_dagg = _edfb
			break
		}
	}
	if _dagg >= 0 {
		_eeac._begbb = append(_eeac._begbb[:_dagg], _eeac._begbb[_dagg+1:]...)
		delete(_eeac._abec, key)
	}
}

// LookupByReference looks up a PdfObject by a reference.
func (_cd *PdfParser) LookupByReference(ref PdfObjectReference) (PdfObject, error) {
	_eb.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0075\u0070\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", ref.String())
	return _cd.LookupByNumber(int(ref.ObjectNumber))
}
func (_bebde *JBIG2Image) toBitmap() (_ccgg *_ec.Bitmap, _faab error) {
	const _acc = "\u004a\u0042\u0049\u00472I\u006d\u0061\u0067\u0065\u002e\u0074\u006f\u0042\u0069\u0074\u006d\u0061\u0070"
	if _bebde.Data == nil {
		return nil, _eca.Error(_acc, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _bebde.Width == 0 || _bebde.Height == 0 {
		return nil, _eca.Error(_acc, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if _bebde.HasPadding {
		_ccgg, _faab = _ec.NewWithData(_bebde.Width, _bebde.Height, _bebde.Data)
	} else {
		_ccgg, _faab = _ec.NewWithUnpaddedData(_bebde.Width, _bebde.Height, _bebde.Data)
	}
	if _faab != nil {
		return nil, _eca.Wrap(_faab, _acc, "")
	}
	return _ccgg, nil
}

// MakeName creates a PdfObjectName from a string.
func MakeName(s string) *PdfObjectName { _baeaa := PdfObjectName(s); return &_baeaa }

// Inspect analyzes the document object structure. Returns a map of object types (by name) with the instance count
// as value.
func (_bffff *PdfParser) Inspect() (map[string]int, error) { return _bffff.inspect() }
func (_ecagg *PdfParser) rebuildXrefTable() error {
	_bgcb := XrefTable{}
	_bgcb.ObjectMap = map[int]XrefObject{}
	_cfdaf := make([]int, 0, len(_ecagg._bfba.ObjectMap))
	for _aga := range _ecagg._bfba.ObjectMap {
		_cfdaf = append(_cfdaf, _aga)
	}
	_be.Ints(_cfdaf)
	for _, _dgdc := range _cfdaf {
		_gffg := _ecagg._bfba.ObjectMap[_dgdc]
		_dbcc, _, _bddga := _ecagg.lookupByNumberWrapper(_dgdc, false)
		if _bddga != nil {
			_eb.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f l\u006f\u006f\u006b\u0020\u0075\u0070\u0020ob\u006a\u0065\u0063\u0074\u0020\u0028\u0025s\u0029", _bddga)
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u006fm\u0070\u006c\u0065\u0074\u0065\u006c\u0079\u0020\u0062\u0072\u006f\u006b\u0065\u006e\u0020\u002d\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0074\u006f \u0072\u0065\u0070\u0061\u0069r\u0020")
			_bcagf, _bfgc := _ecagg.repairRebuildXrefsTopDown()
			if _bfgc != nil {
				_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0072\u0065\u0062\u0075\u0069l\u0064\u0020\u0072\u0065\u0070a\u0069\u0072 \u0028\u0025\u0073\u0029", _bfgc)
				return _bfgc
			}
			_ecagg._bfba = *_bcagf
			_eb.Log.Debug("\u0052e\u0070\u0061\u0069\u0072e\u0064\u0020\u0078\u0072\u0065f\u0020t\u0061b\u006c\u0065\u0020\u0062\u0075\u0069\u006ct")
			return nil
		}
		_eabgd, _gabbc, _bddga := _ed(_dbcc)
		if _bddga != nil {
			return _bddga
		}
		_gffg.ObjectNumber = int(_eabgd)
		_gffg.Generation = int(_gabbc)
		_bgcb.ObjectMap[int(_eabgd)] = _gffg
	}
	_ecagg._bfba = _bgcb
	_eb.Log.Debug("N\u0065w\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061b\u006c\u0065\u0020\u0062ui\u006c\u0074")
	_gfac(_ecagg._bfba)
	return nil
}

// MakeBool creates a PdfObjectBool from a bool value.
func MakeBool(val bool) *PdfObjectBool { _aafca := PdfObjectBool(val); return &_aafca }

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_ege *FlateEncoder) MakeStreamDict() *PdfObjectDictionary {
	_ffa := MakeDict()
	_ffa.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_ege.GetFilterName()))
	_fgg := _ege.MakeDecodeParams()
	if _fgg != nil {
		_ffa.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _fgg)
	}
	return _ffa
}

const _gegcg = 10

// FlateEncoder represents Flate encoding.
type FlateEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Rows    int
	Colors  int
	_gaade  *_eeb.ImageBase
}

var (
	ErrUnsupportedEncodingParameters = _f.New("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	ErrNoCCITTFaxDecode              = _f.New("\u0043\u0043I\u0054\u0054\u0046\u0061\u0078\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJBIG2Decode                 = _f.New("\u004a\u0042\u0049\u0047\u0032\u0044\u0065c\u006f\u0064\u0065 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJPXDecode                   = _f.New("\u004a\u0050\u0058\u0044\u0065c\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoPdfVersion                  = _f.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	ErrTypeError                     = _f.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	ErrRangeError                    = _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrNotSupported                  = _cf.New("\u0066\u0065\u0061t\u0075\u0072\u0065\u0020n\u006f\u0074\u0020\u0063\u0075\u0072\u0072e\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
	ErrNotANumber                    = _f.New("\u006e\u006f\u0074 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
)

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_dfbg *ASCII85Encoder) MakeDecodeParams() PdfObject { return nil }

// IsDelimiter checks if a character represents a delimiter.
func IsDelimiter(c byte) bool {
	return c == '(' || c == ')' || c == '<' || c == '>' || c == '[' || c == ']' || c == '{' || c == '}' || c == '/' || c == '%'
}

// String returns a string describing `ind`.
func (_gabbf *PdfIndirectObject) String() string {
	return _ee.Sprintf("\u0049\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0025\u0064", (*_gabbf).ObjectNumber)
}

// EncodeBytes encodes slice of bytes into JBIG2 encoding format.
// The input 'data' must be an image. In order to Decode it a user is responsible to
// load the codec ('png', 'jpg').
// Returns jbig2 single page encoded document byte slice. The encoder uses DefaultPageSettings
// to encode given image.
func (_dfbgc *JBIG2Encoder) EncodeBytes(data []byte) ([]byte, error) {
	const _gcba = "\u004aB\u0049\u0047\u0032\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002eE\u006e\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065\u0073"
	if _dfbgc.ColorComponents != 1 || _dfbgc.BitsPerComponent != 1 {
		return nil, _eca.Errorf(_gcba, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u002e\u0020\u004a\u0042\u0049G\u0032\u0020E\u006e\u0063o\u0064\u0065\u0072\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020bi\u006e\u0061\u0072\u0079\u0020\u0069\u006d\u0061\u0067e\u0073\u0020\u0064\u0061\u0074\u0061")
	}
	var (
		_cgdgb *_ec.Bitmap
		_eedc  error
	)
	_dec := (_dfbgc.Width * _dfbgc.Height) == len(data)
	if _dec {
		_cgdgb, _eedc = _ec.NewWithUnpaddedData(_dfbgc.Width, _dfbgc.Height, data)
	} else {
		_cgdgb, _eedc = _ec.NewWithData(_dfbgc.Width, _dfbgc.Height, data)
	}
	if _eedc != nil {
		return nil, _eedc
	}
	_ffga := _dfbgc.DefaultPageSettings
	if _eedc = _ffga.Validate(); _eedc != nil {
		return nil, _eca.Wrap(_eedc, _gcba, "")
	}
	if _dfbgc._eabf == nil {
		_dfbgc._eabf = _ce.InitEncodeDocument(_ffga.FileMode)
	}
	switch _ffga.Compression {
	case JB2Generic:
		if _eedc = _dfbgc._eabf.AddGenericPage(_cgdgb, _ffga.DuplicatedLinesRemoval); _eedc != nil {
			return nil, _eca.Wrap(_eedc, _gcba, "")
		}
	case JB2SymbolCorrelation:
		return nil, _eca.Error(_gcba, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return nil, _eca.Error(_gcba, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return nil, _eca.Error(_gcba, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return _dfbgc.Encode()
}

// NewASCIIHexEncoder makes a new ASCII hex encoder.
func NewASCIIHexEncoder() *ASCIIHexEncoder { _fdc := &ASCIIHexEncoder{}; return _fdc }

// WriteString outputs the object as it is to be written to file.
func (_fbgf *PdfObjectString) WriteString() string {
	var _bfcf _bg.Buffer
	if _fbgf._cdca {
		_eedg := _ae.EncodeToString(_fbgf.Bytes())
		_bfcf.WriteString("\u003c")
		_bfcf.WriteString(_eedg)
		_bfcf.WriteString("\u003e")
		return _bfcf.String()
	}
	_aeaa := map[byte]string{'\n': "\u005c\u006e", '\r': "\u005c\u0072", '\t': "\u005c\u0074", '\b': "\u005c\u0062", '\f': "\u005c\u0066", '(': "\u005c\u0028", ')': "\u005c\u0029", '\\': "\u005c\u005c"}
	_bfcf.WriteString("\u0028")
	for _ecbg := 0; _ecbg < len(_fbgf._cbdg); _ecbg++ {
		_daeff := _fbgf._cbdg[_ecbg]
		if _dggbd, _cbgg := _aeaa[_daeff]; _cbgg {
			_bfcf.WriteString(_dggbd)
		} else {
			_bfcf.WriteByte(_daeff)
		}
	}
	_bfcf.WriteString("\u0029")
	return _bfcf.String()
}

// DecodeStream decodes RunLengthEncoded stream object and give back decoded bytes.
func (_ecbb *RunLengthEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ecbb.DecodeBytes(streamObj.Stream)
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_gbc *FlateEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _gbc.Predictor != 1 && _gbc.Predictor != 11 {
		_eb.Log.Debug("E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0046\u006c\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065r\u0020P\u0072\u0065\u0064\u0069c\u0074\u006fr\u0020\u003d\u0020\u0031\u002c\u0020\u0031\u0031\u0020\u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		return nil, ErrUnsupportedEncodingParameters
	}
	if _gbc.Predictor == 11 {
		_aecd := _gbc.Columns
		_ecaf := len(data) / _aecd
		if len(data)%_aecd != 0 {
			_eb.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
			return nil, _f.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
		}
		_ddaf := _bg.NewBuffer(nil)
		_fcb := make([]byte, _aecd)
		for _eeccg := 0; _eeccg < _ecaf; _eeccg++ {
			_ebcc := data[_aecd*_eeccg : _aecd*(_eeccg+1)]
			_fcb[0] = _ebcc[0]
			for _bdeb := 1; _bdeb < _aecd; _bdeb++ {
				_fcb[_bdeb] = byte(int(_ebcc[_bdeb]-_ebcc[_bdeb-1]) % 256)
			}
			_ddaf.WriteByte(1)
			_ddaf.Write(_fcb)
		}
		data = _ddaf.Bytes()
	}
	var _gdcg _bg.Buffer
	_fda := _e.NewWriter(&_gdcg)
	_fda.Write(data)
	_fda.Close()
	return _gdcg.Bytes(), nil
}

// GetFloat returns the *PdfObjectFloat represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetFloat(obj PdfObject) (_aaeb *PdfObjectFloat, _baeed bool) {
	_aaeb, _baeed = TraceToDirectObject(obj).(*PdfObjectFloat)
	return _aaeb, _baeed
}

// GetFloatVal returns the float64 value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetFloatVal(obj PdfObject) (_dbada float64, _fadcb bool) {
	_cgfc, _fadcb := TraceToDirectObject(obj).(*PdfObjectFloat)
	if _fadcb {
		return float64(*_cgfc), true
	}
	return 0, false
}

// GetUpdatedObjects returns pdf objects which were updated from the specific version (from prevParser).
func (_dacg *PdfParser) GetUpdatedObjects(prevParser *PdfParser) (map[int64]PdfObject, error) {
	if prevParser == nil {
		return nil, _f.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_bdebb, _gbbgea := _dacg.getNumbersOfUpdatedObjects(prevParser)
	if _gbbgea != nil {
		return nil, _gbbgea
	}
	_abfb := make(map[int64]PdfObject)
	for _, _begdb := range _bdebb {
		if _dbdf, _fgba := _dacg.LookupByNumber(_begdb); _fgba == nil {
			_abfb[int64(_begdb)] = _dbdf
		} else {
			return nil, _fgba
		}
	}
	return _abfb, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_eeeb *RawEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }

// NewASCII85Encoder makes a new ASCII85 encoder.
func NewASCII85Encoder() *ASCII85Encoder { _bcf := &ASCII85Encoder{}; return _bcf }

// EncryptInfo contains an information generated by the document encrypter.
type EncryptInfo struct {
	Version

	// Encrypt is an encryption dictionary that contains all necessary parameters.
	// It should be stored in all copies of the document trailer.
	Encrypt *PdfObjectDictionary

	// ID0 and ID1 are IDs used in the trailer. Older algorithms such as RC4 uses them for encryption.
	ID0, ID1 string
}

// GetStringVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetStringVal(obj PdfObject) (_eaacf string, _afdg bool) {
	_ccbf, _afdg := TraceToDirectObject(obj).(*PdfObjectString)
	if _afdg {
		return _ccbf.Str(), true
	}
	return
}

// FlattenObject returns the contents of `obj`. In other words, `obj` with indirect objects replaced
// by their values.
// The replacements are made recursively to a depth of traceMaxDepth.
// NOTE: Dicts are sorted to make objects with same contents have the same PDF object strings.
func FlattenObject(obj PdfObject) PdfObject { return _bfcd(obj, 0) }
func (_gcebda *PdfParser) seekToEOFMarker(_agga int64) error {
	var _dfce int64
	var _bcbbe int64 = 2048
	for _dfce < _agga-4 {
		if _agga <= (_bcbbe + _dfce) {
			_bcbbe = _agga - _dfce
		}
		_, _fdfe := _gcebda._cdea.Seek(_agga-_dfce-_bcbbe, _bb.SeekStart)
		if _fdfe != nil {
			return _fdfe
		}
		_bcafc := make([]byte, _bcbbe)
		_gcebda._cdea.Read(_bcafc)
		_eb.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_bcafc))
		_cefb := _bfda.FindAllStringIndex(string(_bcafc), -1)
		if _cefb != nil {
			_gfbce := _cefb[len(_cefb)-1]
			_eb.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _cefb)
			_dbgbf := _agga - _dfce - _bcbbe + int64(_gfbce[0])
			_gcebda._cdea.Seek(_dbgbf, _bb.SeekStart)
			return nil
		}
		_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_dfce += _bcbbe - 4
	}
	_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _bgbf
}
func (_fabd *PdfParser) parseXrefStream(_adcbc *PdfObjectInteger) (*PdfObjectDictionary, error) {
	if _adcbc != nil {
		_eb.Log.Trace("\u0058\u0052\u0065f\u0053\u0074\u006d\u0020x\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0061\u0074\u0020\u0025\u0064", _adcbc)
		_fabd._cdea.Seek(int64(*_adcbc), _bb.SeekStart)
		_fabd._dedbc = _ga.NewReader(_fabd._cdea)
	}
	_gfcc := _fabd.GetFileOffset()
	_dagb, _eeca := _fabd.ParseIndirectObject()
	if _eeca != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0078\u0072\u0065\u0066\u0020\u006fb\u006a\u0065\u0063\u0074")
		return nil, _f.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072e\u0061\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	_eb.Log.Trace("\u0058R\u0065f\u0053\u0074\u006d\u0020\u006fb\u006a\u0065c\u0074\u003a\u0020\u0025\u0073", _dagb)
	_ddef, _dgfga := _dagb.(*PdfObjectStream)
	if !_dgfga {
		_eb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0058R\u0065\u0066\u0053\u0074\u006d\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067 \u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0021")
		return nil, _f.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0070\u006f\u0069\u006e\u0074i\u006e\u0067\u0020\u0074\u006f\u0020a\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_gade := _ddef.PdfObjectDictionary
	_bddgf, _dgfga := _ddef.PdfObjectDictionary.Get("\u0053\u0069\u007a\u0065").(*PdfObjectInteger)
	if !_dgfga {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0069\u007a\u0065\u0020f\u0072\u006f\u006d\u0020\u0078\u0072\u0065f\u0020\u0073\u0074\u006d")
		return nil, _f.New("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0053\u0069\u007ae\u0020\u0066\u0072\u006f\u006d\u0020\u0078\u0072\u0065\u0066 \u0073\u0074\u006d")
	}
	if int64(*_bddgf) > 8388607 {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u0053\u0069\u007a\u0065\u0020\u0065x\u0063\u0065\u0065\u0064\u0065\u0064\u0020l\u0069\u006d\u0069\u0074\u002c\u0020\u006f\u0076\u0065\u0072\u00208\u0033\u0038\u0038\u0036\u0030\u0037\u0020\u0028\u0025\u0064\u0029", *_bddgf)
		return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_eeccb := _ddef.PdfObjectDictionary.Get("\u0057")
	_efgcd, _dgfga := _eeccb.(*PdfObjectArray)
	if !_dgfga {
		return nil, _f.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0020\u0069\u006e\u0020x\u0072\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_bgfd := _efgcd.Len()
	if _bgfd != 3 {
		_eb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020\u0028\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033\u0020\u002d\u0020\u0025\u0064\u0029", _bgfd)
		return nil, _f.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065f\u0020s\u0074\u006d\u0020\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033")
	}
	var _defg []int64
	for _gbebg := 0; _gbebg < 3; _gbebg++ {
		_gacf, _fedg := GetInt(_efgcd.Get(_gbebg))
		if !_fedg {
			return nil, _f.New("i\u006e\u0076\u0061\u006cid\u0020w\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		_defg = append(_defg, int64(*_gacf))
	}
	_dddfg, _eeca := DecodeStream(_ddef)
	if _eeca != nil {
		_eb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f \u0064e\u0063o\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _eeca)
		return nil, _eeca
	}
	_ddagb := int(_defg[0])
	_dbfd := int(_defg[0] + _defg[1])
	_aede := int(_defg[0] + _defg[1] + _defg[2])
	_eggd := int(_defg[0] + _defg[1] + _defg[2])
	if _ddagb < 0 || _dbfd < 0 || _aede < 0 {
		_eb.Log.Debug("\u0045\u0072\u0072\u006fr\u0020\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u003c \u0030 \u0028\u0025\u0064\u002c\u0025\u0064\u002c%\u0064\u0029", _ddagb, _dbfd, _aede)
		return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _eggd == 0 {
		_eb.Log.Debug("\u004e\u006f\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0069\u006e\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028\u0064\u0065\u006c\u0074\u0061\u0062\u0020=\u003d\u0020\u0030\u0029")
		return _gade, nil
	}
	_beeb := len(_dddfg) / _eggd
	_eadc := 0
	_acfc := _ddef.PdfObjectDictionary.Get("\u0049\u006e\u0064e\u0078")
	var _cdfd []int
	if _acfc != nil {
		_eb.Log.Trace("\u0049n\u0064\u0065\u0078\u003a\u0020\u0025b", _acfc)
		_egfa, _ebebf := _acfc.(*PdfObjectArray)
		if !_ebebf {
			_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0029")
			return nil, _f.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0049\u006e\u0064e\u0078\u0020\u006f\u0062je\u0063\u0074")
		}
		if _egfa.Len()%2 != 0 {
			_eb.Log.Debug("\u0057\u0041\u0052\u004eI\u004e\u0047\u0020\u0046\u0061\u0069\u006c\u0075\u0072e\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020i\u006e\u0064\u0065\u0078\u0020n\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006f\u0066\u0020\u0032\u002e")
			return nil, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
		}
		_eadc = 0
		_bgdcg, _efcac := _egfa.ToIntegerArray()
		if _efcac != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078 \u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0073\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0073\u003a\u0020\u0025\u0076", _efcac)
			return nil, _efcac
		}
		for _dcff := 0; _dcff < len(_bgdcg); _dcff += 2 {
			_fgf := _bgdcg[_dcff]
			_aeef := _bgdcg[_dcff+1]
			for _edbcc := 0; _edbcc < _aeef; _edbcc++ {
				_cdfd = append(_cdfd, _fgf+_edbcc)
			}
			_eadc += _aeef
		}
	} else {
		for _bfcc := 0; _bfcc < int(*_bddgf); _bfcc++ {
			_cdfd = append(_cdfd, _bfcc)
		}
		_eadc = int(*_bddgf)
	}
	if _beeb == _eadc+1 {
		_eb.Log.Debug("\u0049n\u0063\u006f\u006d\u0070ati\u0062\u0069\u006c\u0069t\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u0076\u0065\u0072\u0061\u0067\u0065\u0020\u006f\u0066\u0020\u0031\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u002d\u0020\u0061\u0070\u0070en\u0064\u0069\u006eg\u0020\u006f\u006e\u0065\u0020-\u0020M\u0061\u0079\u0020\u006c\u0065\u0061\u0064\u0020\u0074o\u0020\u0070\u0072\u006f\u0062\u006c\u0065\u006d\u0073")
		_fbbge := _eadc - 1
		for _, _cgag := range _cdfd {
			if _cgag > _fbbge {
				_fbbge = _cgag
			}
		}
		_cdfd = append(_cdfd, _fbbge+1)
		_eadc++
	}
	if _beeb != len(_cdfd) {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020x\u0072\u0065\u0066 \u0073\u0074\u006d:\u0020\u006eu\u006d\u0020\u0065\u006e\u0074\u0072i\u0065s \u0021\u003d\u0020\u006c\u0065\u006e\u0028\u0069\u006e\u0064\u0069\u0063\u0065\u0073\u0029\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _beeb, len(_cdfd))
		return nil, _f.New("\u0078\u0072ef\u0020\u0073\u0074m\u0020\u006e\u0075\u006d en\u0074ri\u0065\u0073\u0020\u0021\u003d\u0020\u006cen\u0028\u0069\u006e\u0064\u0069\u0063\u0065s\u0029")
	}
	_eb.Log.Trace("\u004f\u0062j\u0065\u0063\u0074s\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0025\u0064", _eadc)
	_eb.Log.Trace("\u0049\u006e\u0064i\u0063\u0065\u0073\u003a\u0020\u0025\u0020\u0064", _cdfd)
	_faeag := func(_gdca []byte) int64 {
		var _dgea int64
		for _ggaca := 0; _ggaca < len(_gdca); _ggaca++ {
			_dgea += int64(_gdca[_ggaca]) * (1 << uint(8*(len(_gdca)-_ggaca-1)))
		}
		return _dgea
	}
	_eb.Log.Trace("\u0044e\u0063\u006f\u0064\u0065d\u0020\u0073\u0074\u0072\u0065a\u006d \u006ce\u006e\u0067\u0074\u0068\u003a\u0020\u0025d", len(_dddfg))
	_cdebf := 0
	for _cgbbd := 0; _cgbbd < len(_dddfg); _cgbbd += _eggd {
		_cge := _gadd(len(_dddfg), _cgbbd, _cgbbd+_ddagb)
		if _cge != nil {
			_eb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _cge)
			return nil, _cge
		}
		_gfgde := _dddfg[_cgbbd : _cgbbd+_ddagb]
		_cge = _gadd(len(_dddfg), _cgbbd+_ddagb, _cgbbd+_dbfd)
		if _cge != nil {
			_eb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _cge)
			return nil, _cge
		}
		_dafbb := _dddfg[_cgbbd+_ddagb : _cgbbd+_dbfd]
		_cge = _gadd(len(_dddfg), _cgbbd+_dbfd, _cgbbd+_aede)
		if _cge != nil {
			_eb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _cge)
			return nil, _cge
		}
		_ddgd := _dddfg[_cgbbd+_dbfd : _cgbbd+_aede]
		_gdeca := _faeag(_gfgde)
		_eacd := _faeag(_dafbb)
		_abgf := _faeag(_ddgd)
		if _defg[0] == 0 {
			_gdeca = 1
		}
		if _cdebf >= len(_cdfd) {
			_eb.Log.Debug("X\u0052\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u002d\u0020\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063e\u0073s\u0020\u0069\u006e\u0064e\u0078\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073\u0020\u002d\u0020\u0062\u0072\u0065\u0061\u006b\u0069\u006e\u0067")
			break
		}
		_cffd := _cdfd[_cdebf]
		_cdebf++
		_eb.Log.Trace("%\u0064\u002e\u0020\u0070\u0031\u003a\u0020\u0025\u0020\u0078", _cffd, _gfgde)
		_eb.Log.Trace("%\u0064\u002e\u0020\u0070\u0032\u003a\u0020\u0025\u0020\u0078", _cffd, _dafbb)
		_eb.Log.Trace("%\u0064\u002e\u0020\u0070\u0033\u003a\u0020\u0025\u0020\u0078", _cffd, _ddgd)
		_eb.Log.Trace("\u0025d\u002e \u0078\u0072\u0065\u0066\u003a \u0025\u0064 \u0025\u0064\u0020\u0025\u0064", _cffd, _gdeca, _eacd, _abgf)
		if _gdeca == 0 {
			_eb.Log.Trace("-\u0020\u0046\u0072\u0065\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020-\u0020\u0063\u0061\u006e\u0020\u0070\u0072o\u0062\u0061\u0062\u006c\u0079\u0020\u0069\u0067\u006e\u006fr\u0065")
		} else if _gdeca == 1 {
			_eb.Log.Trace("\u002d\u0020I\u006e\u0020\u0075\u0073e\u0020\u002d \u0075\u006e\u0063\u006f\u006d\u0070\u0072\u0065s\u0073\u0065\u0064\u0020\u0076\u0069\u0061\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0062", _dafbb)
			if _eacd == _gfcc {
				_eb.Log.Debug("\u0055\u0070d\u0061\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0058\u0052\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064", _cffd, _ddef.ObjectNumber)
				_cffd = int(_ddef.ObjectNumber)
			}
			if _ggfce, _bcfd := _fabd._bfba.ObjectMap[_cffd]; !_bcfd || int(_abgf) > _ggfce.Generation {
				_ccgc := XrefObject{ObjectNumber: _cffd, XType: XrefTypeTableEntry, Offset: _eacd, Generation: int(_abgf)}
				_fabd._bfba.ObjectMap[_cffd] = _ccgc
			}
		} else if _gdeca == 2 {
			_eb.Log.Trace("\u002d\u0020\u0049\u006e \u0075\u0073\u0065\u0020\u002d\u0020\u0063\u006f\u006d\u0070r\u0065s\u0073\u0065\u0064\u0020\u006f\u0062\u006ae\u0063\u0074")
			if _, _caac := _fabd._bfba.ObjectMap[_cffd]; !_caac {
				_caca := XrefObject{ObjectNumber: _cffd, XType: XrefTypeObjectStream, OsObjNumber: int(_eacd), OsObjIndex: int(_abgf)}
				_fabd._bfba.ObjectMap[_cffd] = _caca
				_eb.Log.Trace("\u0065\u006e\u0074\u0072\u0079\u003a\u0020\u0025\u002b\u0076", _caca)
			}
		} else {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0056\u0041L\u0049\u0044\u0020\u0054\u0059\u0050\u0045\u0020\u0058\u0072\u0065\u0066\u0053\u0074\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u003f\u002d\u002d\u002d\u002d\u002d\u002d-")
			continue
		}
	}
	if _fabd._gfggg == nil {
		_acge := XrefTypeObjectStream
		_fabd._gfggg = &_acge
	}
	return _gade, nil
}

// GetFilterName returns the names of the underlying encoding filters,
// separated by spaces.
// Note: This is just a string, should not be used in /Filter dictionary entry. Use GetFilterArray for that.
// TODO(v4): Refactor to GetFilter() which can be used for /Filter (either Name or Array), this can be
// renamed to String() as a pretty string to use in debugging etc.
func (_edcc *MultiEncoder) GetFilterName() string {
	_efggb := ""
	for _ggda, _addf := range _edcc._abeab {
		_efggb += _addf.GetFilterName()
		if _ggda < len(_edcc._abeab)-1 {
			_efggb += "\u0020"
		}
	}
	return _efggb
}

// DecodeStream decodes the stream data and returns the decoded data.
// An error is returned upon failure.
func DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_eb.Log.Trace("\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_eebf, _fggcf := NewEncoderFromStream(streamObj)
	if _fggcf != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fggcf)
		return nil, _fggcf
	}
	_eb.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076\u000a", _eebf)
	_eegdf, _fggcf := _eebf.DecodeStream(streamObj)
	if _fggcf != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fggcf)
		return nil, _fggcf
	}
	return _eegdf, nil
}

// DecodeStream implements ASCII85 stream decoding.
func (_ffff *ASCII85Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ffff.DecodeBytes(streamObj.Stream)
}
func _ccbb(_fcfad PdfObject, _eegg int, _fedb map[PdfObject]struct{}) error {
	_eb.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072s\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0061 \u0028\u0064\u0065\u0070\u0074\u0068\u0020=\u0020\u0025\u0064\u0029", _eegg)
	if _, _efaff := _fedb[_fcfad]; _efaff {
		_eb.Log.Trace("-\u0041\u006c\u0072\u0065ad\u0079 \u0074\u0072\u0061\u0076\u0065r\u0073\u0065\u0064\u002e\u002e\u002e")
		return nil
	}
	_fedb[_fcfad] = struct{}{}
	switch _eeef := _fcfad.(type) {
	case *PdfIndirectObject:
		_bfeg := _eeef
		_eb.Log.Trace("\u0069\u006f\u003a\u0020\u0025\u0073", _bfeg)
		_eb.Log.Trace("\u002d\u0020\u0025\u0073", _bfeg.PdfObject)
		return _ccbb(_bfeg.PdfObject, _eegg+1, _fedb)
	case *PdfObjectStream:
		_cegde := _eeef
		return _ccbb(_cegde.PdfObjectDictionary, _eegg+1, _fedb)
	case *PdfObjectDictionary:
		_aaed := _eeef
		_eb.Log.Trace("\u002d\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _aaed)
		for _, _ddadf := range _aaed.Keys() {
			_eecb := _aaed.Get(_ddadf)
			if _efbe, _bfab := _eecb.(*PdfObjectReference); _bfab {
				_ddge := _efbe.Resolve()
				_aaed.Set(_ddadf, _ddge)
				_dabe := _ccbb(_ddge, _eegg+1, _fedb)
				if _dabe != nil {
					return _dabe
				}
			} else {
				_gdgfe := _ccbb(_eecb, _eegg+1, _fedb)
				if _gdgfe != nil {
					return _gdgfe
				}
			}
		}
		return nil
	case *PdfObjectArray:
		_abfd := _eeef
		_eb.Log.Trace("-\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _abfd)
		for _dgadd, _aebcc := range _abfd.Elements() {
			if _egdf, _bede := _aebcc.(*PdfObjectReference); _bede {
				_afde := _egdf.Resolve()
				_abfd.Set(_dgadd, _afde)
				_aaea := _ccbb(_afde, _eegg+1, _fedb)
				if _aaea != nil {
					return _aaea
				}
			} else {
				_ecea := _ccbb(_aebcc, _eegg+1, _fedb)
				if _ecea != nil {
					return _ecea
				}
			}
		}
		return nil
	case *PdfObjectReference:
		_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020T\u0072\u0061\u0063\u0069\u006e\u0067\u0020\u0061\u0020r\u0065\u0066\u0065r\u0065n\u0063\u0065\u0021")
		return _f.New("\u0065r\u0072\u006f\u0072\u0020t\u0072\u0061\u0063\u0069\u006eg\u0020a\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	return nil
}

// RawEncoder implements Raw encoder/decoder (no encoding, pass through)
type RawEncoder struct{}

// WriteString outputs the object as it is to be written to file.
func (_ebae *PdfIndirectObject) WriteString() string {
	var _cgcd _cc.Builder
	_cgcd.WriteString(_d.FormatInt(_ebae.ObjectNumber, 10))
	_cgcd.WriteString("\u0020\u0030\u0020\u0052")
	return _cgcd.String()
}

const _fddb = 32 << (^uint(0) >> 63)

func (_beae *PdfParser) parseName() (PdfObjectName, error) {
	var _gbadf _bg.Buffer
	_begg := false
	for {
		_fefab, _eadfc := _beae._dedbc.Peek(1)
		if _eadfc == _bb.EOF {
			break
		}
		if _eadfc != nil {
			return PdfObjectName(_gbadf.String()), _eadfc
		}
		if !_begg {
			if _fefab[0] == '/' {
				_begg = true
				_beae._dedbc.ReadByte()
			} else if _fefab[0] == '%' {
				_beae.readComment()
				_beae.skipSpaces()
			} else {
				_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _fefab, _fefab)
				return PdfObjectName(_gbadf.String()), _ee.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _fefab[0])
			}
		} else {
			if IsWhiteSpace(_fefab[0]) {
				break
			} else if (_fefab[0] == '/') || (_fefab[0] == '[') || (_fefab[0] == '(') || (_fefab[0] == ']') || (_fefab[0] == '<') || (_fefab[0] == '>') {
				break
			} else if _fefab[0] == '#' {
				_facb, _bdece := _beae._dedbc.Peek(3)
				if _bdece != nil {
					return PdfObjectName(_gbadf.String()), _bdece
				}
				_cfcg, _bdece := _ae.DecodeString(string(_facb[1:3]))
				if _bdece != nil {
					_eb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0066o\u006c\u006co\u0077\u0069\u006e\u0067 \u0027\u0023\u0027\u002c \u0063\u006f\u006e\u0074\u0069n\u0075\u0069\u006e\u0067\u0020\u0075\u0073i\u006e\u0067\u0020\u006c\u0069t\u0065\u0072\u0061\u006c\u0020\u002d\u0020\u004f\u0075t\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074")
					_gbadf.WriteByte('#')
					_beae._dedbc.Discard(1)
					continue
				}
				_beae._dedbc.Discard(3)
				_gbadf.Write(_cfcg)
			} else {
				_dadc, _ := _beae._dedbc.ReadByte()
				_gbadf.WriteByte(_dadc)
			}
		}
	}
	return PdfObjectName(_gbadf.String()), nil
}
func (_dfcc *PdfParser) parseXref() (*PdfObjectDictionary, error) {
	_dfcc.skipSpaces()
	const _fefd = 20
	_bbae, _ := _dfcc._dedbc.Peek(_fefd)
	for _edga := 0; _edga < 2; _edga++ {
		if _dfcc._cgfg == 0 {
			_dfcc._cgfg = _dfcc.GetFileOffset()
		}
		if _ffdaa.Match(_bbae) {
			_eb.Log.Trace("\u0078\u0072e\u0066\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002e\u0020\u0050\u0072\u006f\u0062\u0061\u0062\u006c\u0079\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			_eb.Log.Debug("\u0073t\u0061r\u0074\u0069\u006e\u0067\u0020w\u0069\u0074h\u0020\u0022\u0025\u0073\u0022", string(_bbae))
			return _dfcc.parseXrefStream(nil)
		}
		if _cdfg.Match(_bbae) {
			_eb.Log.Trace("\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073e\u0063t\u0069\u006f\u006e\u0020\u0074\u0061\u0062l\u0065\u0021")
			return _dfcc.parseXrefTable()
		}
		_ceafg := _dfcc.GetFileOffset()
		if _dfcc._cgfg == 0 {
			_dfcc._cgfg = _ceafg
		}
		_dfcc.SetFileOffset(_ceafg - _fefd)
		defer _dfcc.SetFileOffset(_ceafg)
		_gfca, _ := _dfcc._dedbc.Peek(_fefd)
		_bbae = append(_gfca, _bbae...)
	}
	_eb.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u0078\u0072\u0065f\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006fr\u0020\u0073\u0074\u0072\u0065\u0061\u006d.\u0020\u0052\u0065\u0070\u0061i\u0072\u0020\u0061\u0074\u0074e\u006d\u0070\u0074\u0065\u0064\u003a\u0020\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0065\u0061\u0072\u006c\u0069\u0065\u0073\u0074\u0020x\u0072\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u0062\u006f\u0074to\u006d\u002e")
	if _dgbfb := _dfcc.repairSeekXrefMarker(); _dgbfb != nil {
		_eb.Log.Debug("\u0052e\u0070a\u0069\u0072\u0020\u0066\u0061i\u006c\u0065d\u0020\u002d\u0020\u0025\u0076", _dgbfb)
		return nil, _dgbfb
	}
	return _dfcc.parseXrefTable()
}

// Resolve resolves a PdfObject to direct object, looking up and resolving references as needed (unlike TraceToDirect).
func (_egbd *PdfParser) Resolve(obj PdfObject) (PdfObject, error) {
	_gdd, _fce := obj.(*PdfObjectReference)
	if !_fce {
		return obj, nil
	}
	_dbe := _egbd.GetFileOffset()
	defer func() { _egbd.SetFileOffset(_dbe) }()
	_bbf, _acg := _egbd.LookupByReference(*_gdd)
	if _acg != nil {
		return nil, _acg
	}
	_fbe, _ebg := _bbf.(*PdfIndirectObject)
	if !_ebg {
		return _bbf, nil
	}
	_bbf = _fbe.PdfObject
	_, _fce = _bbf.(*PdfObjectReference)
	if _fce {
		return _fbe, _f.New("\u006d\u0075lt\u0069\u0020\u0064e\u0070\u0074\u0068\u0020tra\u0063e \u0070\u006f\u0069\u006e\u0074\u0065\u0072 t\u006f\u0020\u0070\u006f\u0069\u006e\u0074e\u0072")
	}
	return _bbf, nil
}

// GetString returns the *PdfObjectString represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetString(obj PdfObject) (_acdf *PdfObjectString, _eabd bool) {
	_acdf, _eabd = TraceToDirectObject(obj).(*PdfObjectString)
	return _acdf, _eabd
}
func _cbce(_gbae *PdfObjectStream, _begc *PdfObjectDictionary) (*RunLengthEncoder, error) {
	return NewRunLengthEncoder(), nil
}

// PdfCrypt provides PDF encryption/decryption support.
// The PDF standard supports encryption of strings and streams (Section 7.6).
type PdfCrypt struct {
	_ffe  encryptDict
	_adfe _cbd.StdEncryptDict
	_ffc  string
	_cbcg []byte
	_ge   map[PdfObject]bool
	_edc  map[PdfObject]bool
	_agg  bool
	_ace  cryptFilters
	_badb string
	_df   string
	_eadg *PdfParser
	_eda  map[int]struct{}
}

// ToGoImage converts the JBIG2Image to the golang image.Image.
func (_aedg *JBIG2Image) ToGoImage() (_a.Image, error) {
	const _efgfb = "J\u0042I\u0047\u0032\u0049\u006d\u0061\u0067\u0065\u002eT\u006f\u0047\u006f\u0049ma\u0067\u0065"
	if _aedg.Data == nil {
		return nil, _eca.Error(_efgfb, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _aedg.Width == 0 || _aedg.Height == 0 {
		return nil, _eca.Error(_efgfb, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	_ddfb, _bagcb := _eeb.NewImage(_aedg.Width, _aedg.Height, 1, 1, _aedg.Data, nil, nil)
	if _bagcb != nil {
		return nil, _bagcb
	}
	return _ddfb, nil
}

// PdfObjectBool represents the primitive PDF boolean object.
type PdfObjectBool bool

// Str returns the string value of the PdfObjectString. Defined in addition to String() function to clarify that
// this function returns the underlying string directly, whereas the String function technically could include
// debug info.
func (_efafb *PdfObjectString) Str() string { return _efafb._cbdg }
func _fdae(_ggbe string) (PdfObjectReference, error) {
	_adaa := PdfObjectReference{}
	_cbg := _fgaa.FindStringSubmatch(_ggbe)
	if len(_cbg) < 3 {
		_eb.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _adaa, _f.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_abcf, _ := _d.Atoi(_cbg[1])
	_addff, _ := _d.Atoi(_cbg[2])
	_adaa.ObjectNumber = int64(_abcf)
	_adaa.GenerationNumber = int64(_addff)
	return _adaa, nil
}

// ResolveReference resolves reference if `o` is a *PdfObjectReference and returns the object referenced to.
// Otherwise returns back `o`.
func ResolveReference(obj PdfObject) PdfObject {
	if _cdaab, _fbfge := obj.(*PdfObjectReference); _fbfge {
		return _cdaab.Resolve()
	}
	return obj
}

// HasNonConformantStream implements core.ParserMetadata.
func (_fbc ParserMetadata) HasNonConformantStream() bool { return _fbc._caf }

// GetNumberAsFloat returns the contents of `obj` as a float if it is an integer or float, or an
// error if it isn't.
func GetNumberAsFloat(obj PdfObject) (float64, error) {
	switch _aebb := obj.(type) {
	case *PdfObjectFloat:
		return float64(*_aebb), nil
	case *PdfObjectInteger:
		return float64(*_aebb), nil
	case *PdfObjectReference:
		_edgb := TraceToDirectObject(obj)
		return GetNumberAsFloat(_edgb)
	case *PdfIndirectObject:
		return GetNumberAsFloat(_aebb.PdfObject)
	}
	return 0, ErrNotANumber
}

// JBIG2EncoderSettings contains the parameters and settings used by the JBIG2Encoder.
// Current version works only on JB2Generic compression.
type JBIG2EncoderSettings struct {

	// FileMode defines if the jbig2 encoder should return full jbig2 file instead of
	// shortened pdf mode. This adds the file header to the jbig2 definition.
	FileMode bool

	// Compression is the setting that defines the compression type used for encoding the page.
	Compression JBIG2CompressionType

	// DuplicatedLinesRemoval code generic region in a way such that if the lines are duplicated the encoder
	// doesn't store it twice.
	DuplicatedLinesRemoval bool

	// DefaultPixelValue is the bit value initial for every pixel in the page.
	DefaultPixelValue uint8

	// ResolutionX optional setting that defines the 'x' axis input image resolution - used for single page encoding.
	ResolutionX int

	// ResolutionY optional setting that defines the 'y' axis input image resolution - used for single page encoding.
	ResolutionY int

	// Threshold defines the threshold of the image correlation for
	// non Generic compression.
	// User only for JB2SymbolCorrelation and JB2SymbolRankHaus methods.
	// Best results in range [0.7 - 0.98] - the less the better the compression would be
	// but the more lossy.
	// Default value: 0.95
	Threshold float64
}

// DecodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_cfcab *RawEncoder) DecodeBytes(encoded []byte) ([]byte, error) { return encoded, nil }
