// Package model provides an interface for working with high-level objects (models) in PDF files, including
// reading and writing documents.
//
// The document structure of a PDF is constructed of a hierarchy of data models, representing a tree
// of information starting from the Document catalog (Figure 5 p. 80).
// It is based on the core package which handles core functionality such as file i/o, parsing and
// handling of primitive PDF objects (core.PdfObject).
//
// As an example of the interface, the following snippet can read the PDF and output the number of pages:
package model

import (
	_ga "bufio"
	_dd "bytes"
	_c "crypto/md5"
	_gaaa "crypto/rand"
	_bd "crypto/sha1"
	_bag "crypto/x509"
	_ef "encoding/binary"
	_egf "encoding/hex"
	_dcf "errors"
	_e "fmt"
	_ecd "hash"
	_fb "image"
	_b "image/color"
	_ "image/gif"
	_ "image/png"
	_bagf "io"
	_gg "math"
	_gaa "math/rand"
	_ccb "os"
	_dc "path/filepath"
	_ab "regexp"
	_ba "sort"
	_de "strconv"
	_cc "strings"
	_ec "sync"
	_d "time"
	_eg "unicode"
	_ae "unicode/utf8"

	_ddb "github.com/bamzi/pdfext/common"
	_eb "github.com/bamzi/pdfext/core"
	_cg "github.com/bamzi/pdfext/core/security"
	_be "github.com/bamzi/pdfext/core/security/crypt"
	_ff "github.com/bamzi/pdfext/internal/cmap"
	_df "github.com/bamzi/pdfext/internal/imageutil"
	_cf "github.com/bamzi/pdfext/internal/license"
	_bb "github.com/bamzi/pdfext/internal/sampling"
	_fc "github.com/bamzi/pdfext/internal/textencoding"
	_eggb "github.com/bamzi/pdfext/internal/timeutils"
	_ffg "github.com/bamzi/pdfext/internal/transform"
	_deb "github.com/bamzi/pdfext/internal/uuid"
	_cd "github.com/bamzi/pdfext/model/internal/docutil"
	_fg "github.com/bamzi/pdfext/model/internal/fonts"
	_bab "github.com/bamzi/pdfext/model/mdp"
	_dda "github.com/bamzi/pdfext/model/sigutil"
	_gc "github.com/bamzi/pdfext/ps"
	_egg "github.com/gabriel-vasile/mimetype"
	_dg "github.com/unidoc/pkcs7"
	_fa "github.com/unidoc/unitype"
	_db "golang.org/x/xerrors"
)

// SetContext sets the sub action (context).
func (_cfb *PdfAction) SetContext(ctx PdfModel) { _cfb._aee = ctx }

// Normalize swaps (Llx,Urx) if Urx < Llx, and (Lly,Ury) if Ury < Lly.
func (_daaff *PdfRectangle) Normalize() {
	if _daaff.Llx > _daaff.Urx {
		_daaff.Llx, _daaff.Urx = _daaff.Urx, _daaff.Llx
	}
	if _daaff.Lly > _daaff.Ury {
		_daaff.Lly, _daaff.Ury = _daaff.Ury, _daaff.Lly
	}
}
func _agedg() string {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	return _aeee
}

// ToInteger convert to an integer format.
func (_babfb *PdfColorLab) ToInteger(bits int) [3]uint32 {
	_cccf := _gg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_cccf * _babfb.L()), uint32(_cccf * _babfb.A()), uint32(_cccf * _babfb.B())}
}

// PdfActionMovie represents a movie action.
type PdfActionMovie struct {
	*PdfAction
	Annotation _eb.PdfObject
	T          _eb.PdfObject
	Operation  _eb.PdfObject
}

// ToPdfObject returns the PDF representation of the pattern.
func (_ccbga *PdfPattern) ToPdfObject() _eb.PdfObject {
	_bebb := _ccbga.getDict()
	_bebb.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_bebb.Set("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065", _eb.MakeInteger(_ccbga.PatternType))
	return _ccbga._agddd
}
func _adbb(_cdad *PdfPage) map[_eb.PdfObjectName]_eb.PdfObject {
	_gbfd := make(map[_eb.PdfObjectName]_eb.PdfObject)
	if _cdad.Resources == nil {
		return _gbfd
	}
	if _cdad.Resources.Font != nil {
		if _bea, _afec := _eb.GetDict(_cdad.Resources.Font); _afec {
			for _, _gcbc := range _bea.Keys() {
				_gbfd[_gcbc] = _bea.Get(_gcbc)
			}
		}
	}
	if _cdad.Resources.ExtGState != nil {
		if _bagg, _dbgb := _eb.GetDict(_cdad.Resources.ExtGState); _dbgb {
			for _, _cde := range _bagg.Keys() {
				_gbfd[_cde] = _bagg.Get(_cde)
			}
		}
	}
	if _cdad.Resources.XObject != nil {
		if _edce, _cgfa := _eb.GetDict(_cdad.Resources.XObject); _cgfa {
			for _, _fbbc := range _edce.Keys() {
				_gbfd[_fbbc] = _edce.Get(_fbbc)
			}
		}
	}
	if _cdad.Resources.Pattern != nil {
		if _eecf, _fcab := _eb.GetDict(_cdad.Resources.Pattern); _fcab {
			for _, _bda := range _eecf.Keys() {
				_gbfd[_bda] = _eecf.Get(_bda)
			}
		}
	}
	if _cdad.Resources.Shading != nil {
		if _dggcc, _ddde := _eb.GetDict(_cdad.Resources.Shading); _ddde {
			for _, _ddegb := range _dggcc.Keys() {
				_gbfd[_ddegb] = _dggcc.Get(_ddegb)
			}
		}
	}
	if _cdad.Resources.ProcSet != nil {
		if _deee, _cbgb := _eb.GetDict(_cdad.Resources.ProcSet); _cbgb {
			for _, _ddgf := range _deee.Keys() {
				_gbfd[_ddgf] = _deee.Get(_ddgf)
			}
		}
	}
	if _cdad.Resources.Properties != nil {
		if _dbgbf, _dbea := _eb.GetDict(_cdad.Resources.Properties); _dbea {
			for _, _dcec := range _dbgbf.Keys() {
				_gbfd[_dcec] = _dbgbf.Get(_dcec)
			}
		}
	}
	return _gbfd
}
func _edcffb(_gfgf *_eb.PdfObjectDictionary) (*PdfShadingType4, error) {
	_gcefd := PdfShadingType4{}
	_gdgca := _gfgf.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _gdgca == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dbdad, _gebd := _gdgca.(*_eb.PdfObjectInteger)
	if !_gebd {
		_ddb.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gdgca)
		return nil, _eb.ErrTypeError
	}
	_gcefd.BitsPerCoordinate = _dbdad
	_gdgca = _gfgf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _gdgca == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_dbdad, _gebd = _gdgca.(*_eb.PdfObjectInteger)
	if !_gebd {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _gdgca)
		return nil, _eb.ErrTypeError
	}
	_gcefd.BitsPerComponent = _dbdad
	_gdgca = _gfgf.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _gdgca == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dbdad, _gebd = _gdgca.(*_eb.PdfObjectInteger)
	if !_gebd {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _gdgca)
		return nil, _eb.ErrTypeError
	}
	_gcefd.BitsPerComponent = _dbdad
	_gdgca = _gfgf.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _gdgca == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_gbdba, _gebd := _gdgca.(*_eb.PdfObjectArray)
	if !_gebd {
		_ddb.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gdgca)
		return nil, _eb.ErrTypeError
	}
	_gcefd.Decode = _gbdba
	_gdgca = _gfgf.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _gdgca == nil {
		_ddb.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_gcefd.Function = []PdfFunction{}
	if _cabafe, _gbdfb := _gdgca.(*_eb.PdfObjectArray); _gbdfb {
		for _, _bedc := range _cabafe.Elements() {
			_fgcdd, _aagbd := _cccfa(_bedc)
			if _aagbd != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _aagbd)
				return nil, _aagbd
			}
			_gcefd.Function = append(_gcefd.Function, _fgcdd)
		}
	} else {
		_ccacdb, _fcege := _cccfa(_gdgca)
		if _fcege != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fcege)
			return nil, _fcege
		}
		_gcefd.Function = append(_gcefd.Function, _ccacdb)
	}
	return &_gcefd, nil
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ddfgg pdfFontType3) GetCharMetrics(code _fc.CharCode) (_fg.CharMetrics, bool) {
	if _bgaaa, _abfeb := _ddfgg._ccfdb[code]; _abfeb {
		return _fg.CharMetrics{Wx: _bgaaa}, true
	}
	if _fg.IsStdFont(_fg.StdFontName(_ddfgg._agcc)) {
		return _fg.CharMetrics{Wx: 250}, true
	}
	return _fg.CharMetrics{}, false
}

// NewPdfAnnotationPopup returns a new popup annotation.
func NewPdfAnnotationPopup() *PdfAnnotationPopup {
	_eeac := NewPdfAnnotation()
	_accf := &PdfAnnotationPopup{}
	_accf.PdfAnnotation = _eeac
	_eeac.SetContext(_accf)
	return _accf
}

// PdfColorspaceDeviceRGB represents an RGB colorspace.
type PdfColorspaceDeviceRGB struct{}

// PdfColorspaceLab is a L*, a*, b* 3 component colorspace.
type PdfColorspaceLab struct {
	WhitePoint []float64
	BlackPoint []float64
	Range      []float64
	_cbag      *_eb.PdfIndirectObject
}

// SetVersion sets the PDF version of the output file.
func (_afbdf *PdfWriter) SetVersion(majorVersion, minorVersion int) {
	_afbdf._edbbf.Major = majorVersion
	_afbdf._edbbf.Minor = minorVersion
}

// IDTree represents the ID tree dictionary where the format of the content
// is using Name Trees as described in chapter 7.9.6.
type IDTree struct {

	// Limits shall be an array of two strings, that shall specify the (lexically) least and greatest keys included in the Names array.
	Limits *_eb.PdfObjectArray

	// Names shall be an array of the form
	//
	// [ key1 value1 key2 value2 … keyn valuen]
	//
	// where each keyi shall be a string and the corresponding valuei shall be the object
	// associated with that key. The keys shall be sorted in lexical order, as described below.
	Names *_eb.PdfObjectArray

	// Kids Shall be an array of indirect references to the immediate children of this node.
	Kids []*IDTree
}

// SetPage directly sets the page object.
func (_aadab *KDict) SetPage(page *_eb.PdfIndirectObject) { _aadab.Pg = page }

// SetXObjectFormByName adds the provided XObjectForm to the page resources.
// The added XObjectForm is identified by the specified name.
func (_ecgfe *PdfPageResources) SetXObjectFormByName(keyName _eb.PdfObjectName, xform *XObjectForm) error {
	_dfdge := xform.ToPdfObject().(*_eb.PdfObjectStream)
	_befac := _ecgfe.SetXObjectByName(keyName, _dfdge)
	return _befac
}

// Outline represents a PDF outline dictionary (Table 152 - p. 376).
// Currently, the Outline object can only be used to construct PDF outlines.
type Outline struct {
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// GetCatalogMarkInfo gets catalog MarkInfo object.
func (_bggfb *PdfReader) GetCatalogMarkInfo() (_eb.PdfObject, bool) {
	if _bggfb._bagcfd == nil {
		return nil, false
	}
	_deefg := _bggfb._bagcfd.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
	return _deefg, _deefg != nil
}
func (_cefa *PdfReader) newPdfAnnotationProjectionFromDict(_bgfb *_eb.PdfObjectDictionary) (*PdfAnnotationProjection, error) {
	_degd := &PdfAnnotationProjection{}
	_dace, _fdee := _cefa.newPdfAnnotationMarkupFromDict(_bgfb)
	if _fdee != nil {
		return nil, _fdee
	}
	_degd.PdfAnnotationMarkup = _dace
	return _degd, nil
}

// GetNumComponents returns the number of color components (1 for Indexed).
func (_gfea *PdfColorspaceSpecialIndexed) GetNumComponents() int { return 1 }

var (
	StructureTypeDocument      = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074"
	StructureTypePart          = "\u0050\u0061\u0072\u0074"
	StructureTypeArticle       = "\u0041\u0072\u0074"
	StructureTypeSection       = "\u0053\u0065\u0063\u0074"
	StructureTypeDivision      = "\u0044\u0069\u0076"
	StructureTypeBlockQuote    = "\u0042\u006c\u006f\u0063\u006b\u0051\u0075\u006f\u0074\u0065"
	StructureTypeCaption       = "\u0043a\u0070\u0074\u0069\u006f\u006e"
	StructureTypeTOC           = "\u0054\u004f\u0043"
	StructureTypeTOCI          = "\u0054\u004f\u0043\u0049"
	StructureTypeIndex         = "\u0049\u006e\u0064e\u0078"
	StructureTypeNonStructural = "\u004eo\u006e\u0053\u0074\u0072\u0075\u0063t"
	StructureTypePrivate       = "\u0050r\u0069\u0076\u0061\u0074\u0065"
)
var _ pdfFont = (*pdfFontSimple)(nil)

func (_cgef *PdfReader) newPdfAnnotationCaretFromDict(_cfbg *_eb.PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	_dfdb := PdfAnnotationCaret{}
	_eca, _bbdf := _cgef.newPdfAnnotationMarkupFromDict(_cfbg)
	if _bbdf != nil {
		return nil, _bbdf
	}
	_dfdb.PdfAnnotationMarkup = _eca
	_dfdb.RD = _cfbg.Get("\u0052\u0044")
	_dfdb.Sy = _cfbg.Get("\u0053\u0079")
	return &_dfdb, nil
}

// GetIndirectObjectByNumber retrieves and returns a specific PdfObject by object number.
func (_ebdgba *PdfReader) GetIndirectObjectByNumber(number int) (_eb.PdfObject, error) {
	_bbegg, _gcfbf := _ebdgba._ebbe.LookupByNumber(number)
	return _bbegg, _gcfbf
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_fdadea *PdfShadingPattern) ToPdfObject() _eb.PdfObject {
	_fdadea.PdfPattern.ToPdfObject()
	_eegce := _fdadea.getDict()
	if _fdadea.Shading != nil {
		_eegce.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _fdadea.Shading.ToPdfObject())
	}
	if _fdadea.Matrix != nil {
		_eegce.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _fdadea.Matrix)
	}
	if _fdadea.ExtGState != nil {
		_eegce.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _fdadea.ExtGState)
	}
	return _fdadea._agddd
}

// SetKDict sets the KDict for the KValue.
func (_dbbd *KValue) SetKDict(kDict *KDict) { _dbbd.Clear(); _dbbd._ccbca = kDict }
func _cbcf(_cfge *_eb.PdfObjectDictionary) (*PdfFieldText, error) {
	_ffff := &PdfFieldText{}
	_ffff.DA, _ = _eb.GetString(_cfge.Get("\u0044\u0041"))
	_ffff.Q, _ = _eb.GetInt(_cfge.Get("\u0051"))
	_ffff.DS, _ = _eb.GetString(_cfge.Get("\u0044\u0053"))
	_ffff.RV = _cfge.Get("\u0052\u0056")
	_ffff.MaxLen, _ = _eb.GetInt(_cfge.Get("\u004d\u0061\u0078\u004c\u0065\u006e"))
	return _ffff, nil
}

// PdfAction represents an action in PDF (section 12.6 p. 412).
type PdfAction struct {
	_aee PdfModel
	Type _eb.PdfObject
	S    _eb.PdfObject
	Next _eb.PdfObject
	_dee *_eb.PdfIndirectObject
}

func (_fbabb *PdfWriter) adjustXRefAffectedVersion(_ecdac bool) {
	if _ecdac && _fbabb._edbbf.Major == 1 && _fbabb._edbbf.Minor < 5 {
		_fbabb._edbbf.Minor = 5
	}
}

// OutlineItem represents a PDF outline item dictionary (Table 153 - pp. 376 - 377).
type OutlineItem struct {
	Title   string         `json:"title"`
	Dest    OutlineDest    `json:"dest"`
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// GetContext returns a reference to the subshading entry as represented by PdfShadingType1-7.
func (_fdbfg *PdfShading) GetContext() PdfModel { return _fdbfg._ecffg }

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 4 for a CMYK32 device.
func (_gadd *PdfColorspaceDeviceCMYK) GetNumComponents() int { return 4 }

// ToPdfObject implements interface PdfModel.
func (_aeafd *PdfSignatureReference) ToPdfObject() _eb.PdfObject {
	_dbdec := _eb.MakeDict()
	_dbdec.SetIfNotNil("\u0054\u0079\u0070\u0065", _aeafd.Type)
	_dbdec.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064", _aeafd.TransformMethod)
	_dbdec.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073", _aeafd.TransformParams)
	_dbdec.SetIfNotNil("\u0044\u0061\u0074\u0061", _aeafd.Data)
	_dbdec.SetIfNotNil("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064", _aeafd.DigestMethod)
	return _dbdec
}

// PdfFunctionType4 is a Postscript calculator functions.
type PdfFunctionType4 struct {
	Domain  []float64
	Range   []float64
	Program *_gc.PSProgram
	_aabg   *_gc.PSExecutor
	_bdcf   []byte
	_gbcd   *_eb.PdfObjectStream
}

func (_dfecaa *fontFile) parseASCIIPart(_aacb []byte) error {
	if len(_aacb) < 2 || string(_aacb[:2]) != "\u0025\u0021" {
		return _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0074a\u0072\u0074\u0020\u006f\u0066\u0020\u0041S\u0043\u0049\u0049\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_acade, _gbbf, _gagd := _bgadc(_aacb)
	if _gagd != nil {
		return _gagd
	}
	_ffad := _eeaaf(_acade)
	_dfecaa._cddc = _ffad["\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"]
	if _dfecaa._cddc == "" {
		_ddb.Log.Debug("\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0020\u0068a\u0073\u0020\u006e\u006f\u0020\u002f\u0046\u006f\u006e\u0074N\u0061\u006d\u0065")
	}
	if _gbbf != "" {
		_fgebe, _dacff := _agbf(_gbbf)
		if _dacff != nil {
			return _dacff
		}
		_degfb, _dacff := _fc.NewCustomSimpleTextEncoder(_fgebe, nil)
		if _dacff != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004fR\u0020\u003a\u0055\u004e\u004b\u004e\u004f\u0057\u004e\u0020G\u004cY\u0050\u0048\u003a\u0020\u0065\u0072\u0072=\u0025\u0076", _dacff)
			return nil
		}
		_dfecaa._gggff = _degfb
	}
	return nil
}

var ImageHandling ImageHandler = DefaultImageHandler{}

// NewPdfColorspaceCalRGB returns a new CalRGB colorspace object.
func NewPdfColorspaceCalRGB() *PdfColorspaceCalRGB {
	_gacfd := &PdfColorspaceCalRGB{}
	_gacfd.BlackPoint = []float64{0.0, 0.0, 0.0}
	_gacfd.Gamma = []float64{1.0, 1.0, 1.0}
	_gacfd.Matrix = []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	return _gacfd
}
func _fbfae(_gfcf _fg.StdFont) pdfFontSimple {
	_bbad := _gfcf.Descriptor()
	return pdfFontSimple{fontCommon: fontCommon{_fgdee: "\u0054\u0079\u0070e\u0031", _agcc: _gfcf.Name()}, _gacdb: _gfcf.GetMetricsTable(), _debdb: &PdfFontDescriptor{FontName: _eb.MakeName(string(_bbad.Name)), FontFamily: _eb.MakeName(_bbad.Family), FontWeight: _eb.MakeFloat(float64(_bbad.Weight)), Flags: _eb.MakeInteger(int64(_bbad.Flags)), FontBBox: _eb.MakeArrayFromFloats(_bbad.BBox[:]), ItalicAngle: _eb.MakeFloat(_bbad.ItalicAngle), Ascent: _eb.MakeFloat(_bbad.Ascent), Descent: _eb.MakeFloat(_bbad.Descent), CapHeight: _eb.MakeFloat(_bbad.CapHeight), XHeight: _eb.MakeFloat(_bbad.XHeight), StemV: _eb.MakeFloat(_bbad.StemV), StemH: _eb.MakeFloat(_bbad.StemH)}, _gcgb: _gfcf.Encoder()}
}

// NewOutline returns a new outline instance.
func NewOutline() *Outline { return &Outline{} }

// AcroFormNeedsRepair returns true if the document contains widget annotations
// linked to fields which are not referenced in the AcroForm. The AcroForm can
// be repaired using the RepairAcroForm method of the reader.
func (_ccca *PdfReader) AcroFormNeedsRepair() (bool, error) {
	var _dgcgf []*PdfField
	if _ccca.AcroForm != nil {
		_dgcgf = _ccca.AcroForm.AllFields()
	}
	_ebggd := make(map[*PdfField]struct{}, len(_dgcgf))
	for _, _cdffe := range _dgcgf {
		_ebggd[_cdffe] = struct{}{}
	}
	for _, _decg := range _ccca.PageList {
		_dbba, _edabd := _decg.GetAnnotations()
		if _edabd != nil {
			return false, _edabd
		}
		for _, _dafbed := range _dbba {
			_facede, _ddff := _dafbed.GetContext().(*PdfAnnotationWidget)
			if !_ddff {
				continue
			}
			_afeab := _facede.Field()
			if _afeab == nil {
				return true, nil
			}
			if _, _bcbba := _ebggd[_afeab]; !_bcbba {
				return true, nil
			}
		}
	}
	return false, nil
}

// SetAction sets the PDF action for the annotation link.
func (_fcb *PdfAnnotationLink) SetAction(action *PdfAction) {
	_fcb._fbg = action
	if action == nil {
		_fcb.A = nil
	}
}

const (
	TrappedUnknown PdfInfoTrapped = "\u0055n\u006b\u006e\u006f\u0077\u006e"
	TrappedTrue    PdfInfoTrapped = "\u0054\u0072\u0075\u0065"
	TrappedFalse   PdfInfoTrapped = "\u0046\u0061\u006cs\u0065"
)

func (_bdfb *PdfReader) newPdfActionTransFromDict(_aef *_eb.PdfObjectDictionary) (*PdfActionTrans, error) {
	return &PdfActionTrans{Trans: _aef.Get("\u0054\u0072\u0061n\u0073")}, nil
}
func (_becfe *pdfFontType0) subsetRegistered() error {
	_bgda, _eeee := _becfe.DescendantFont._fdaa.(*pdfCIDFontType2)
	if !_eeee {
		_ddb.Log.Debug("\u0046\u006fnt\u0020\u006e\u006ft\u0020\u0073\u0075\u0070por\u0074ed\u0020\u0066\u006f\u0072\u0020\u0073\u0075bs\u0065\u0074\u0074\u0069\u006e\u0067\u0020%\u0054", _becfe.DescendantFont)
		return nil
	}
	if _bgda == nil {
		return nil
	}
	if _bgda._bged == nil {
		_ddb.Log.Debug("\u004d\u0069\u0073si\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return nil
	}
	if _becfe._edcff == nil {
		_ddb.Log.Debug("\u004e\u006f\u0020e\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0073\u0075\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0067\u006e\u006f\u0072\u0065\u0064")
		return nil
	}
	_daaa, _eeee := _eb.GetStream(_bgda._bged.FontFile2)
	if !_eeee {
		_ddb.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _dcf.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_gacedf, _aeac := _eb.DecodeStream(_daaa)
	if _aeac != nil {
		_ddb.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _aeac)
		return _aeac
	}
	_dbafb, _aeac := _fa.Parse(_dd.NewReader(_gacedf))
	if _aeac != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_daaa.Stream))
		return _aeac
	}
	var _dbdab []rune
	var _cbbge *_fa.Font
	switch _bafcg := _becfe._edcff.(type) {
	case *_fc.TrueTypeFontEncoder:
		_dbdab = _bafcg.RegisteredRunes()
		_cbbge, _aeac = _dbafb.SubsetKeepRunes(_dbdab)
		if _aeac != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aeac)
			return _aeac
		}
		_bafcg.SubsetRegistered()
	case *_fc.IdentityEncoder:
		_dbdab = _bafcg.RegisteredRunes()
		_fdaab := make([]_fa.GlyphIndex, len(_dbdab))
		for _abbgf, _gaffa := range _dbdab {
			_fdaab[_abbgf] = _fa.GlyphIndex(_gaffa)
		}
		_cbbge, _aeac = _dbafb.SubsetKeepIndices(_fdaab)
		if _aeac != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aeac)
			return _aeac
		}
	case _fc.SimpleEncoder:
		_eebbc := _bafcg.Charcodes()
		for _, _daaf := range _eebbc {
			_deba, _aaba := _bafcg.CharcodeToRune(_daaf)
			if !_aaba {
				_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0020\u0074\u006f \u0072\u0075\u006e\u0065\u003a\u0020\u0025\u0064", _daaf)
				continue
			}
			_dbdab = append(_dbdab, _deba)
		}
	default:
		return _e.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069\u006eg\u003a\u0020\u0025\u0054", _becfe._edcff)
	}
	var _ecab _dd.Buffer
	_aeac = _cbbge.Write(&_ecab)
	if _aeac != nil {
		_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aeac)
		return _aeac
	}
	if _becfe._bgbg != nil {
		_dgeba := make(map[_ff.CharCode]rune, len(_dbdab))
		for _, _dcdc := range _dbdab {
			_gdcfc, _ecgfd := _becfe._edcff.RuneToCharcode(_dcdc)
			if !_ecgfd {
				continue
			}
			_dgeba[_ff.CharCode(_gdcfc)] = _dcdc
		}
		_becfe._bgbg = _ff.NewToUnicodeCMap(_dgeba)
	}
	_daaa, _aeac = _eb.MakeStream(_ecab.Bytes(), _eb.NewFlateEncoder())
	if _aeac != nil {
		_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aeac)
		return _aeac
	}
	_daaa.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _eb.MakeInteger(int64(_ecab.Len())))
	if _ffeba, _eggbd := _eb.GetStream(_bgda._bged.FontFile2); _eggbd {
		*_ffeba = *_daaa
	} else {
		_bgda._bged.FontFile2 = _daaa
	}
	_adbcg := _fgeg()
	if len(_becfe._agcc) > 0 {
		_becfe._agcc = _dafbe(_becfe._agcc, _adbcg)
	}
	if len(_bgda._agcc) > 0 {
		_bgda._agcc = _dafbe(_bgda._agcc, _adbcg)
	}
	if len(_becfe._bgge) > 0 {
		_becfe._bgge = _dafbe(_becfe._bgge, _adbcg)
	}
	if _bgda._bged != nil {
		_ebba, _bfec := _eb.GetName(_bgda._bged.FontName)
		if _bfec && len(_ebba.String()) > 0 {
			_aafc := _dafbe(_ebba.String(), _adbcg)
			_bgda._bged.FontName = _eb.MakeName(_aafc)
		}
	}
	return nil
}
func (_ade *PdfReader) newPdfActionResetFormFromDict(_eff *_eb.PdfObjectDictionary) (*PdfActionResetForm, error) {
	return &PdfActionResetForm{Fields: _eff.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _eff.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// GetContext returns a reference to the subpattern entry: either PdfTilingPattern or PdfShadingPattern.
func (_fgcdcb *PdfPattern) GetContext() PdfModel { return _fgcdcb._eefgb }

// ToPdfObject returns the PDF representation of the tiling pattern.
func (_adcfe *PdfTilingPattern) ToPdfObject() _eb.PdfObject {
	_adcfe.PdfPattern.ToPdfObject()
	_ceafg := _adcfe.getDict()
	if _adcfe.PaintType != nil {
		_ceafg.Set("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e", _adcfe.PaintType)
	}
	if _adcfe.TilingType != nil {
		_ceafg.Set("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _adcfe.TilingType)
	}
	if _adcfe.BBox != nil {
		_ceafg.Set("\u0042\u0042\u006f\u0078", _adcfe.BBox.ToPdfObject())
	}
	if _adcfe.XStep != nil {
		_ceafg.Set("\u0058\u0053\u0074e\u0070", _adcfe.XStep)
	}
	if _adcfe.YStep != nil {
		_ceafg.Set("\u0059\u0053\u0074e\u0070", _adcfe.YStep)
	}
	if _adcfe.Resources != nil {
		_ceafg.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _adcfe.Resources.ToPdfObject())
	}
	if _adcfe.Matrix != nil {
		_ceafg.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _adcfe.Matrix)
	}
	return _adcfe._agddd
}

// ToPdfObject implements interface PdfModel.
func (_aaa *PdfBorderStyle) ToPdfObject() _eb.PdfObject {
	_fecg := _eb.MakeDict()
	if _aaa._gdgd != nil {
		if _bbf, _affb := _aaa._gdgd.(*_eb.PdfIndirectObject); _affb {
			_bbf.PdfObject = _fecg
		}
	}
	_fecg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0042\u006f\u0072\u0064\u0065\u0072"))
	if _aaa.W != nil {
		_fecg.Set("\u0057", _eb.MakeFloat(*_aaa.W))
	}
	if _aaa.S != nil {
		_fecg.Set("\u0053", _eb.MakeName(_aaa.S.GetPdfName()))
	}
	if _aaa.D != nil {
		_fecg.Set("\u0044", _eb.MakeArrayFromIntegers(*_aaa.D))
	}
	if _aaa._gdgd != nil {
		return _aaa._gdgd
	}
	return _fecg
}

// RemoveStructParentsKey removes the StructParents key.
func (_fagdd *PdfPage) RemoveStructParentsKey() { _fagdd.StructParents = nil }

// SetFillImage attach a model.Image to push button.
func (_acfdf *PdfFieldButton) SetFillImage(image *Image) {
	if _acfdf.IsPush() {
		_acfdf._cadc = image
	}
}
func (_gfafc *PdfFunctionType0) processSamples() error {
	_fddeb := _bb.ResampleBytes(_gfafc._dcdgc, _gfafc.BitsPerSample)
	_gfafc._bgaaab = _fddeb
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_fab *PdfAnnotationTrapNet) ToPdfObject() _eb.PdfObject {
	_fab.PdfAnnotation.ToPdfObject()
	_decc := _fab._ggf
	_bdec := _decc.PdfObject.(*_eb.PdfObjectDictionary)
	_bdec.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0054r\u0061\u0070\u004e\u0065\u0074"))
	return _decc
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_adacg *PdfColorspaceCalGray) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aefc, _eecfe := _eb.GetNumbersAsFloat(objects)
	if _eecfe != nil {
		return nil, _eecfe
	}
	return _adacg.ColorFromFloats(_aefc)
}

// NewPdfPage returns a new PDF page.
func NewPdfPage() *PdfPage {
	_ecdge := PdfPage{}
	_ecdge._aaagaa = _eb.MakeDict()
	_ecdge.Resources = NewPdfPageResources()
	_ccfgc := _eb.PdfIndirectObject{}
	_ccfgc.PdfObject = _ecdge._aaagaa
	_ecdge._efcff = &_ccfgc
	_ecdge._eaebc = *_ecdge._aaagaa
	return &_ecdge
}
func (_dceb *pdfCIDFontType2) getFontDescriptor() *PdfFontDescriptor { return _dceb._bged }

// ToPdfObject implements interface PdfModel.
func (_fbfa *PdfAnnotationWatermark) ToPdfObject() _eb.PdfObject {
	_fbfa.PdfAnnotation.ToPdfObject()
	_fdef := _fbfa._ggf
	_afg := _fdef.PdfObject.(*_eb.PdfObjectDictionary)
	_afg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k"))
	_afg.SetIfNotNil("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074", _fbfa.FixedPrint)
	return _fdef
}

var ErrColorOutOfRange = _dcf.New("\u0063o\u006co\u0072\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")

// SetPdfCreator sets the Creator attribute of the output PDF.
func SetPdfCreator(creator string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _geba = creator }

// PdfActionLaunch represents a launch action.
type PdfActionLaunch struct {
	*PdfAction
	F         *PdfFilespec
	Win       _eb.PdfObject
	Mac       _eb.PdfObject
	Unix      _eb.PdfObject
	NewWindow _eb.PdfObject
}

// SetAlpha sets the alpha layer for the image.
func (_gafad *Image) SetAlpha(alpha []byte) { _gafad._bdcab = alpha }

// PdfAnnotationPolyLine represents PolyLine annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolyLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _eb.PdfObject
	LE       _eb.PdfObject
	BS       _eb.PdfObject
	IC       _eb.PdfObject
	BE       _eb.PdfObject
	IT       _eb.PdfObject
	Measure  _eb.PdfObject
}

// HasShadingByName checks whether a shading is defined by the specified keyName.
func (_fgddc *PdfPageResources) HasShadingByName(keyName _eb.PdfObjectName) bool {
	_, _dbbac := _fgddc.GetShadingByName(keyName)
	return _dbbac
}

// ToPdfObject converts the pdfCIDFontType0 to a PDF representation.
func (_cgcag *pdfCIDFontType0) ToPdfObject() _eb.PdfObject { return _eb.MakeNull() }
func (_gcgef *PdfReader) resolveReference(_edfc *_eb.PdfObjectReference) (_eb.PdfObject, bool, error) {
	_dacgc, _bafag := _gcgef._ebbe.ObjCache[int(_edfc.ObjectNumber)]
	if !_bafag {
		_ddb.Log.Trace("R\u0065\u0061\u0064\u0065r \u004co\u006f\u006b\u0075\u0070\u0020r\u0065\u0066\u003a\u0020\u0025\u0073", _edfc)
		_cfab, _ceaab := _gcgef._ebbe.LookupByReference(*_edfc)
		if _ceaab != nil {
			return nil, false, _ceaab
		}
		_gcgef._ebbe.ObjCache[int(_edfc.ObjectNumber)] = _cfab
		return _cfab, false, nil
	}
	return _dacgc, true, nil
}
func (_beb *PdfReader) newPdfActionGotoFromDict(_gfc *_eb.PdfObjectDictionary) (*PdfActionGoTo, error) {
	return &PdfActionGoTo{D: _gfc.Get("\u0044")}, nil
}

const (
	RelationshipSource FileRelationship = iota
	RelationshipData
	RelationshipAlternative
	RelationshipSupplement
	RelationshipUnspecified
)

// PdfOutlineTreeNode contains common fields used by the outline and outline
// item objects.
type PdfOutlineTreeNode struct {
	_eeedb interface{}
	First  *PdfOutlineTreeNode
	Last   *PdfOutlineTreeNode
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_beab *PdfColorspaceSpecialPattern) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _beab.UnderlyingCS == nil {
		return nil, _dcf.New("u\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0043\u0053\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
	}
	return _beab.UnderlyingCS.ColorFromFloats(vals)
}

// GetRevision returns the specific version of the PdfReader for the current Pdf document
func (_adcdg *PdfReader) GetRevision(revisionNumber int) (*PdfReader, error) {
	_abacd := _adcdg._ebbe.GetRevisionNumber()
	if revisionNumber < 0 || revisionNumber > _abacd {
		return nil, _dcf.New("w\u0072\u006f\u006e\u0067 r\u0065v\u0069\u0073\u0069\u006f\u006e \u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if revisionNumber == _abacd {
		return _adcdg, nil
	}
	if _adcdg._dfeg[revisionNumber] != nil {
		return _adcdg._dfeg[revisionNumber], nil
	}
	_bedbc := _adcdg
	for _aaaad := _abacd - 1; _aaaad >= revisionNumber; _aaaad-- {
		_bbcgd, _caffc := _bedbc.GetPreviousRevision()
		if _caffc != nil {
			return nil, _caffc
		}
		_adcdg._dfeg[_aaaad] = _bbcgd
		_bedbc = _bbcgd
	}
	return _bedbc, nil
}

var _baffe = map[string]struct{}{"\u0054\u0069\u0074l\u0065": {}, "\u0041\u0075\u0074\u0068\u006f\u0072": {}, "\u0053u\u0062\u006a\u0065\u0063\u0074": {}, "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073": {}, "\u0043r\u0065\u0061\u0074\u006f\u0072": {}, "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072": {}, "\u0054r\u0061\u0070\u0070\u0065\u0064": {}, "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065": {}, "\u004do\u0064\u0044\u0061\u0074\u0065": {}}

// NewPdfAnnotation3D returns a new 3d annotation.
func NewPdfAnnotation3D() *PdfAnnotation3D {
	_bdfd := NewPdfAnnotation()
	_dca := &PdfAnnotation3D{}
	_dca.PdfAnnotation = _bdfd
	_bdfd.SetContext(_dca)
	return _dca
}

// PdfColorPatternType2 represents a color shading pattern type 2 (Axial).
type PdfColorPatternType2 struct {
	Color       PdfColor
	PatternName _eb.PdfObjectName
}

// String returns a human readable description of `fontfile`.
func (_ccded *fontFile) String() string {
	_agdgg := "\u005b\u004e\u006f\u006e\u0065\u005d"
	if _ccded._gggff != nil {
		_agdgg = _ccded._gggff.String()
	}
	return _e.Sprintf("\u0046O\u004e\u0054\u0046\u0049\u004c\u0045\u007b\u0025\u0023\u0071\u0020e\u006e\u0063\u006f\u0064\u0065\u0072\u003d\u0025\u0073\u007d", _ccded._cddc, _agdgg)
}

// NewPdfAnnotationLink returns a new link annotation.
func NewPdfAnnotationLink() *PdfAnnotationLink {
	_dcg := NewPdfAnnotation()
	_ggfe := &PdfAnnotationLink{}
	_ggfe.PdfAnnotation = _dcg
	_dcg.SetContext(_ggfe)
	return _ggfe
}

// GetPdfInfo returns the PDF info dictionary.
func (_fffbc *PdfReader) GetPdfInfo() (*PdfInfo, error) {
	_eegag, _accc := _fffbc.GetTrailer()
	if _accc != nil {
		return nil, _accc
	}
	var _eddc *_eb.PdfObjectDictionary
	_fecbcf := _eegag.Get("\u0049\u006e\u0066\u006f")
	switch _aadce := _fecbcf.(type) {
	case *_eb.PdfObjectReference:
		_cbdfa := _aadce
		_fecbcf, _accc = _fffbc.GetIndirectObjectByNumber(int(_cbdfa.ObjectNumber))
		_fecbcf = _eb.TraceToDirectObject(_fecbcf)
		if _accc != nil {
			return nil, _accc
		}
		_eddc, _ = _fecbcf.(*_eb.PdfObjectDictionary)
	case *_eb.PdfObjectDictionary:
		_eddc = _aadce
	}
	if _eddc == nil {
		return nil, _dcf.New("I\u006e\u0066\u006f\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0070r\u0065s\u0065\u006e\u0074")
	}
	_eefbg, _accc := NewPdfInfoFromObject(_eddc)
	if _accc != nil {
		return nil, _accc
	}
	return _eefbg, nil
}
func (_gcaga *pdfFontType0) getFontDescriptor() *PdfFontDescriptor {
	if _gcaga._bged == nil && _gcaga.DescendantFont != nil {
		return _gcaga.DescendantFont.FontDescriptor()
	}
	return _gcaga._bged
}

// NewPdfFilespecFromObj creates and returns a new PdfFilespec object.
func NewPdfFilespecFromObj(obj _eb.PdfObject) (*PdfFilespec, error) {
	_edgff := &PdfFilespec{}
	var _fgfc *_eb.PdfObjectDictionary
	if _agaa, _dcfcc := _eb.GetIndirect(obj); _dcfcc {
		_edgff._cbefe = _agaa
		_daeff, _dbee := _eb.GetDict(_agaa.PdfObject)
		if !_dbee {
			_ddb.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _eb.ErrTypeError
		}
		_fgfc = _daeff
	} else if _cddeg, _begfb := _eb.GetDict(obj); _begfb {
		_edgff._cbefe = _cddeg
		_fgfc = _cddeg
	} else {
		_ddb.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", obj)
		return nil, _eb.ErrTypeError
	}
	if _fgfc == nil {
		_ddb.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _dcf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _gceb := _fgfc.Get("\u0054\u0079\u0070\u0065"); _gceb != nil {
		_gfgg, _bcab := _gceb.(*_eb.PdfObjectName)
		if !_bcab {
			_ddb.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _gceb)
		} else {
			if *_gfgg != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
				_ddb.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070e\u0063\u0074\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063 \u0028\u0025\u0073\u0029", *_gfgg)
			}
		}
	}
	if _efeg := _fgfc.Get("\u0046\u0053"); _efeg != nil {
		_edgff.FS = _efeg
	}
	if _baeed := _fgfc.Get("\u0046"); _baeed != nil {
		_edgff.F = _baeed
	}
	if _ccdb := _fgfc.Get("\u0055\u0046"); _ccdb != nil {
		_edgff.UF = _ccdb
	}
	if _dgcb := _fgfc.Get("\u0044\u004f\u0053"); _dgcb != nil {
		_edgff.DOS = _dgcb
	}
	if _cbfc := _fgfc.Get("\u004d\u0061\u0063"); _cbfc != nil {
		_edgff.Mac = _cbfc
	}
	if _fgbgc := _fgfc.Get("\u0055\u006e\u0069\u0078"); _fgbgc != nil {
		_edgff.Unix = _fgbgc
	}
	if _aedc := _fgfc.Get("\u0049\u0044"); _aedc != nil {
		_edgff.ID = _aedc
	}
	if _dfec := _fgfc.Get("\u0056"); _dfec != nil {
		_edgff.V = _dfec
	}
	if _dagf := _fgfc.Get("\u0045\u0046"); _dagf != nil {
		_edgff.EF = _dagf
	}
	if _edged := _fgfc.Get("\u0052\u0046"); _edged != nil {
		_edgff.RF = _edged
	}
	if _bfbaf := _fgfc.Get("\u0044\u0065\u0073\u0063"); _bfbaf != nil {
		_edgff.Desc = _bfbaf
	}
	if _cbge := _fgfc.Get("\u0043\u0049"); _cbge != nil {
		_edgff.CI = _cbge
	}
	if _agad := _fgfc.Get("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070"); _agad != nil {
		_edgff.AFRelationship = _agad
	}
	return _edgff, nil
}
func (_dgfe *PdfField) inherit(_aega func(*PdfField) bool) (bool, error) {
	_fegd := map[*PdfField]bool{}
	_gbbgd := false
	_cecbf := _dgfe
	for _cecbf != nil {
		if _, _gdec := _fegd[_cecbf]; _gdec {
			return false, _dcf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		_cdaaeg := _aega(_cecbf)
		if _cdaaeg {
			_gbbgd = true
			break
		}
		_fegd[_cecbf] = true
		_cecbf = _cecbf.Parent
	}
	return _gbbgd, nil
}

// GetPreviousRevision returns the previous revision of PdfReader for the Pdf document
func (_eecca *PdfReader) GetPreviousRevision() (*PdfReader, error) {
	if _eecca._ebbe.GetRevisionNumber() == 0 {
		return nil, _dcf.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065xi\u0073\u0074")
	}
	if _caeaeg, _bgedb := _eecca._dbag[_eecca]; _bgedb {
		return _caeaeg, nil
	}
	_aebefc, _cbcfa := _eecca._ebbe.GetPreviousRevisionReadSeeker()
	if _cbcfa != nil {
		return nil, _cbcfa
	}
	_bgabg, _cbcfa := _dcgfe(_aebefc, _eecca._eacdg, _eecca._daag, "\u006do\u0064\u0065\u006c\u003aG\u0065\u0074\u0050\u0072\u0065v\u0069o\u0075s\u0052\u0065\u0076\u0069\u0073\u0069\u006fn")
	if _cbcfa != nil {
		return nil, _cbcfa
	}
	_eecca._dfeg[_eecca._ebbe.GetRevisionNumber()-1] = _bgabg
	_eecca._dbag[_eecca] = _bgabg
	_bgabg._dbag = _eecca._dbag
	return _bgabg, nil
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_ecaf *PdfColorspaceSpecialIndexed) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_agfbg, _dfcb := _eb.GetNumbersAsFloat(objects)
	if _dfcb != nil {
		return nil, _dfcb
	}
	return _ecaf.ColorFromFloats(_agfbg)
}

// PdfTransformParamsDocMDP represents a transform parameters dictionary for the DocMDP method and is used to detect
// modifications relative to a signature field that is signed by the author of a document.
// (Section 12.8.2.2, Table 254 - Entries in the DocMDP transform parameters dictionary p. 471 in PDF32000_2008).
type PdfTransformParamsDocMDP struct {
	Type *_eb.PdfObjectName
	P    *_eb.PdfObjectInteger
	V    *_eb.PdfObjectName
}

// NewPdfAnnotationProjection returns a new projection annotation.
func NewPdfAnnotationProjection() *PdfAnnotationProjection {
	_bcee := NewPdfAnnotation()
	_cbd := &PdfAnnotationProjection{}
	_cbd.PdfAnnotation = _bcee
	_cbd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_bcee.SetContext(_cbd)
	return _cbd
}
func _faeef(_agaee *PdfField, _ddbff _eb.PdfObject) {
	for _, _abdb := range _agaee.Annotations {
		_abdb.AS = _ddbff
		_abdb.ToPdfObject()
	}
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// A, B and C components of the color. The values of the elements should be
// between 0 and 1.
func (_agceg *PdfColorspaceCalRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bbbbe := vals[0]
	if _bbbbe < 0.0 || _bbbbe > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bbbbe)
		return nil, ErrColorOutOfRange
	}
	_fdgg := vals[1]
	if _fdgg < 0.0 || _fdgg > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fdgg)
		return nil, ErrColorOutOfRange
	}
	_ccbb := vals[2]
	if _ccbb < 0.0 || _ccbb > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ccbb)
		return nil, ErrColorOutOfRange
	}
	_fefad := NewPdfColorCalRGB(_bbbbe, _fdgg, _ccbb)
	return _fefad, nil
}

// SetDisplayDocTitle sets the value of the displayDocTitle flag.
func (_cafdd *ViewerPreferences) SetDisplayDocTitle(displayDocTitle bool) {
	_cafdd._bgdge = &displayDocTitle
}

// FontDescriptor returns font's PdfFontDescriptor. This may be a builtin descriptor for standard 14
// fonts but must be an explicit descriptor for other fonts.
func (_bgcb *PdfFont) FontDescriptor() *PdfFontDescriptor {
	if _bgcb.baseFields()._bged != nil {
		return _bgcb.baseFields()._bged
	}
	if _eaaea := _bgcb._fdaa.getFontDescriptor(); _eaaea != nil {
		return _eaaea
	}
	_ddb.Log.Error("\u0041\u006cl \u0066\u006f\u006et\u0073\u0020\u0068\u0061ve \u0061 D\u0065\u0073\u0063\u0072\u0069\u0070\u0074or\u002e\u0020\u0066\u006f\u006e\u0074\u003d%\u0073", _bgcb)
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_cdab *PdfAnnotationFileAttachment) ToPdfObject() _eb.PdfObject {
	_cdab.PdfAnnotation.ToPdfObject()
	_cfceb := _cdab._ggf
	_edbg := _cfceb.PdfObject.(*_eb.PdfObjectDictionary)
	_cdab.PdfAnnotationMarkup.appendToPdfDictionary(_edbg)
	_edbg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074"))
	_edbg.SetIfNotNil("\u0046\u0053", _cdab.FS)
	_edbg.SetIfNotNil("\u004e\u0061\u006d\u0065", _cdab.Name)
	return _cfceb
}

// PdfAnnotationSound represents Sound annotations.
// (Section 12.5.6.16).
type PdfAnnotationSound struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Sound _eb.PdfObject
	Name  _eb.PdfObject
}

var (
	StructureTypeParagraph       = "\u0050"
	StructureTypeHeader          = "\u0048"
	StructureTypeHeading1        = "\u0048\u0031"
	StructureTypeHeading2        = "\u0048\u0032"
	StructureTypeHeading3        = "\u0048\u0033"
	StructureTypeHeading4        = "\u0048\u0034"
	StructureTypeHeading5        = "\u0048\u0035"
	StructureTypeHeading6        = "\u0048\u0036"
	StructureTypeList            = "\u004c"
	StructureTypeListItem        = "\u004c\u0049"
	StructureTypeLabel           = "\u004c\u0062\u006c"
	StructureTypeListBody        = "\u004c\u0042\u006fd\u0079"
	StructureTypeTable           = "\u0054\u0061\u0062l\u0065"
	StructureTypeTableRow        = "\u0054\u0052"
	StructureTypeTableHeaderCell = "\u0054\u0048"
	StructureTypeTableData       = "\u0054\u0044"
	StructureTypeTableHead       = "\u0054\u0048\u0065a\u0064"
	StructureTypeTableBody       = "\u0054\u0042\u006fd\u0079"
	StructureTypeTableFooter     = "\u0054\u0046\u006fo\u0074"
)

// NewPdfColorspaceSpecialPattern returns a new pattern color.
func NewPdfColorspaceSpecialPattern() *PdfColorspaceSpecialPattern {
	return &PdfColorspaceSpecialPattern{}
}
func (_bfffb *XObjectImage) getParamsDict() *_eb.PdfObjectDictionary {
	_ccgab := _eb.MakeDict()
	_ccgab.Set("\u0057\u0069\u0064t\u0068", _eb.MakeInteger(*_bfffb.Width))
	_ccgab.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _eb.MakeInteger(*_bfffb.Height))
	_ccgab.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _eb.MakeInteger(int64(_bfffb.ColorSpace.GetNumComponents())))
	_ccgab.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _eb.MakeInteger(*_bfffb.BitsPerComponent))
	return _ccgab
}
func (_afdd *PdfAppender) replaceObject(_fegb, _ccad _eb.PdfObject) {
	switch _aedb := _fegb.(type) {
	case *_eb.PdfIndirectObject:
		_afdd._ebcbc[_ccad] = _aedb.ObjectNumber
	case *_eb.PdfObjectStream:
		_afdd._ebcbc[_ccad] = _aedb.ObjectNumber
	}
}

// Initialize initializes the PdfSignature.
func (_eedbd *PdfSignature) Initialize() error {
	if _eedbd.Handler == nil {
		return _dcf.New("\u0073\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0068\u0061\u006e\u0064\u006c\u0065r\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	return _eedbd.Handler.InitSignature(_eedbd)
}

// DSS represents a Document Security Store dictionary.
// The DSS dictionary contains both global and signature specific validation
// information. The certificates and revocation data in the `Certs`, `OCSPs`,
// and `CRLs` fields can be used to validate any signature in the document.
// Additionally, the VRI entry contains validation data per signature.
// The keys in the VRI entry are calculated as upper(hex(sha1(sig.Contents))).
// The values are VRI dictionaries containing certificates and revocation
// information used for validating a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type DSS struct {
	_fcdgc *_eb.PdfIndirectObject
	Certs  []*_eb.PdfObjectStream
	OCSPs  []*_eb.PdfObjectStream
	CRLs   []*_eb.PdfObjectStream
	VRI    map[string]*VRI
	_eaed  map[string]*_eb.PdfObjectStream
	_gccb  map[string]*_eb.PdfObjectStream
	_adcdd map[string]*_eb.PdfObjectStream
}

// PdfShadingType2 is an Axial shading.
type PdfShadingType2 struct {
	*PdfShading
	Coords   *_eb.PdfObjectArray
	Domain   *_eb.PdfObjectArray
	Function []PdfFunction
	Extend   *_eb.PdfObjectArray
}

func (_afccc *pdfFontType0) bytesToCharcodes(_defdg []byte) ([]_fc.CharCode, bool) {
	if _afccc._ebeb == nil {
		return nil, false
	}
	_gdbdad, _dbfg := _afccc._ebeb.BytesToCharcodes(_defdg)
	if !_dbfg {
		return nil, false
	}
	_ffae := make([]_fc.CharCode, len(_gdbdad))
	for _eggba, _aggf := range _gdbdad {
		_ffae[_eggba] = _fc.CharCode(_aggf)
	}
	return _ffae, true
}

// PdfColorspace interface defines the common methods of a PDF colorspace.
// The colorspace defines the data storage format for each color and color representation.
//
// Device based colorspace, specified by name
// - /DeviceGray
// - /DeviceRGB
// - /DeviceCMYK
//
// CIE based colorspace specified by [name, dictionary]
// - [/CalGray dict]
// - [/CalRGB dict]
// - [/Lab dict]
// - [/ICCBased dict]
//
// Special colorspaces
// - /Pattern
// - /Indexed
// - /Separation
// - /DeviceN
//
// Work is in progress to support all colorspaces. At the moment ICCBased color spaces fall back to the alternate
// colorspace which works OK in most cases. For full color support, will need fully featured ICC support.
type PdfColorspace interface {

	// String returns the PdfColorspace's name.
	String() string

	// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
	ImageToRGB(Image) (Image, error)

	// ColorToRGB converts a single color in a given PdfColorspace to an RGB color.
	ColorToRGB(_aagde PdfColor) (PdfColor, error)

	// GetNumComponents returns the number of components in the PdfColorspace.
	GetNumComponents() int

	// ToPdfObject returns a PdfObject representation of the PdfColorspace.
	ToPdfObject() _eb.PdfObject

	// ColorFromPdfObjects returns a PdfColor in the given PdfColorspace from an array of PdfObject where each
	// PdfObject represents a numeric value.
	ColorFromPdfObjects(_gbde []_eb.PdfObject) (PdfColor, error)

	// ColorFromFloats returns a new PdfColor based on input color components for a given PdfColorspace.
	ColorFromFloats(_dggcf []float64) (PdfColor, error)

	// DecodeArray returns the Decode array for the PdfColorSpace, i.e. the range of each component.
	DecodeArray() []float64
}

// ReplaceAcroForm replaces the acrobat form. It appends a new form to the Pdf which
// replaces the original AcroForm.
func (_dggdg *PdfAppender) ReplaceAcroForm(acroForm *PdfAcroForm) {
	if acroForm != nil {
		_dggdg.updateObjectsDeep(acroForm.ToPdfObject(), nil)
	}
	_dggdg._bddda = acroForm
}

// ToPdfObject returns the PDF representation of the function.
func (_fecbc *PdfFunctionType0) ToPdfObject() _eb.PdfObject {
	if _fecbc._bded == nil {
		_fecbc._bded = &_eb.PdfObjectStream{}
	}
	_afgcg := _eb.MakeDict()
	_afgcg.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _eb.MakeInteger(0))
	_aeeb := &_eb.PdfObjectArray{}
	for _, _eaac := range _fecbc.Domain {
		_aeeb.Append(_eb.MakeFloat(_eaac))
	}
	_afgcg.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _aeeb)
	_efdf := &_eb.PdfObjectArray{}
	for _, _fbdb := range _fecbc.Range {
		_efdf.Append(_eb.MakeFloat(_fbdb))
	}
	_afgcg.Set("\u0052\u0061\u006eg\u0065", _efdf)
	_edbgg := &_eb.PdfObjectArray{}
	for _, _gfbb := range _fecbc.Size {
		_edbgg.Append(_eb.MakeInteger(int64(_gfbb)))
	}
	_afgcg.Set("\u0053\u0069\u007a\u0065", _edbgg)
	_afgcg.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065", _eb.MakeInteger(int64(_fecbc.BitsPerSample)))
	if _fecbc.Order != 1 {
		_afgcg.Set("\u004f\u0072\u0064e\u0072", _eb.MakeInteger(int64(_fecbc.Order)))
	}
	_afgcg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_fecbc._dcdgc))))
	_fecbc._bded.Stream = _fecbc._dcdgc
	_fecbc._bded.PdfObjectDictionary = _afgcg
	return _fecbc._bded
}

// NewPdfAnnotationLine returns a new line annotation.
func NewPdfAnnotationLine() *PdfAnnotationLine {
	_cdfa := NewPdfAnnotation()
	_fcba := &PdfAnnotationLine{}
	_fcba.PdfAnnotation = _cdfa
	_fcba.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cdfa.SetContext(_fcba)
	return _fcba
}

// NewPdfAnnotationWidget returns an initialized annotation widget.
func NewPdfAnnotationWidget() *PdfAnnotationWidget {
	_cbb := NewPdfAnnotation()
	_bga := &PdfAnnotationWidget{}
	_bga.PdfAnnotation = _cbb
	_cbb.SetContext(_bga)
	return _bga
}
func _dafbe(_ccaaf, _aaeaf string) string {
	if _cc.Contains(_ccaaf, "\u002b") {
		_gdaaf := _cc.Split(_ccaaf, "\u002b")
		if len(_gdaaf) == 2 {
			_ccaaf = _gdaaf[1]
		}
	}
	return _aaeaf + "\u002b" + _ccaaf
}

// NewCompositePdfFontFromTTF loads a composite TTF font. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTF.
func NewCompositePdfFontFromTTF(r _bagf.ReadSeeker) (*PdfFont, error) {
	_eafd, _cbcd := _bagf.ReadAll(r)
	if _cbcd != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _cbcd)
		return nil, _cbcd
	}
	_dfae, _cbcd := _fg.TtfParse(_dd.NewReader(_eafd))
	if _cbcd != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067 \u0074\u0074\u0066\u0020\u0066\u006f\u006et\u003a\u0020\u0025\u0076", _cbcd)
		return nil, _cbcd
	}
	_eaedb := &pdfCIDFontType2{fontCommon: fontCommon{_fgdee: "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032"}, CIDToGIDMap: _eb.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079")}
	if len(_dfae.Widths) <= 0 {
		return nil, _dcf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_fefd := 1000.0 / float64(_dfae.UnitsPerEm)
	_eadbe := _fefd * float64(_dfae.Widths[0])
	_dfafa := make(map[rune]int)
	_bfee := make(map[_fg.GID]int)
	_fbcg := _fg.GID(len(_dfae.Widths))
	for _bdfba, _ecge := range _dfae.Chars {
		if _ecge > _fbcg-1 {
			continue
		}
		_eaggg := int(_fefd * float64(_dfae.Widths[_ecge]))
		_dfafa[_bdfba] = _eaggg
		_bfee[_ecge] = _eaggg
	}
	_eaedb._caba = _dfafa
	_eaedb.DW = _eb.MakeInteger(int64(_eadbe))
	_agec := _ccbbc(_bfee, uint16(_fbcg))
	_eaedb.W = _eb.MakeIndirectObject(_agec)
	_aedd := _eb.MakeDict()
	_aedd.Set("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067", _eb.MakeString("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"))
	_aedd.Set("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079", _eb.MakeString("\u0041\u0064\u006fb\u0065"))
	_aedd.Set("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074", _eb.MakeInteger(0))
	_eaedb.CIDSystemInfo = _aedd
	_afce := &PdfFontDescriptor{FontName: _eb.MakeName(_dfae.PostScriptName), Ascent: _eb.MakeFloat(_fefd * float64(_dfae.TypoAscender)), Descent: _eb.MakeFloat(_fefd * float64(_dfae.TypoDescender)), CapHeight: _eb.MakeFloat(_fefd * float64(_dfae.CapHeight)), FontBBox: _eb.MakeArrayFromFloats([]float64{_fefd * float64(_dfae.Xmin), _fefd * float64(_dfae.Ymin), _fefd * float64(_dfae.Xmax), _fefd * float64(_dfae.Ymax)}), ItalicAngle: _eb.MakeFloat(_dfae.ItalicAngle), MissingWidth: _eb.MakeFloat(_eadbe)}
	_fgfg, _cbcd := _eb.MakeStream(_eafd, _eb.NewFlateEncoder())
	if _cbcd != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _cbcd)
		return nil, _cbcd
	}
	_fgfg.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _eb.MakeInteger(int64(len(_eafd))))
	_afce.FontFile2 = _fgfg
	if _dfae.Bold {
		_afce.StemV = _eb.MakeInteger(120)
	} else {
		_afce.StemV = _eb.MakeInteger(70)
	}
	_gbeb := _bdcce
	if _dfae.IsFixedPitch {
		_gbeb |= _eaggc
	}
	if _dfae.ItalicAngle != 0 {
		_gbeb |= _ffbe
	}
	_afce.Flags = _eb.MakeInteger(int64(_gbeb))
	_eaedb._agcc = _dfae.PostScriptName
	_eaedb._bged = _afce
	_fbdea := pdfFontType0{fontCommon: fontCommon{_fgdee: "\u0054\u0079\u0070e\u0030", _agcc: _dfae.PostScriptName}, DescendantFont: &PdfFont{_fdaa: _eaedb}, Encoding: _eb.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048"), _edcff: _dfae.NewEncoder()}
	if len(_dfae.Chars) > 0 {
		_ebdf := make(map[_ff.CharCode]rune, len(_dfae.Chars))
		for _ecgcd, _dfef := range _dfae.Chars {
			_affbfc := _ff.CharCode(_dfef)
			if _gfefg, _feeda := _ebdf[_affbfc]; !_feeda || (_feeda && _gfefg > _ecgcd) {
				_ebdf[_affbfc] = _ecgcd
			}
		}
		_fbdea._bgbg = _ff.NewToUnicodeCMap(_ebdf)
	}
	_fgdfe := PdfFont{_fdaa: &_fbdea}
	return &_fgdfe, nil
}

// WatermarkImageOptions contains options for configuring the watermark process.
type WatermarkImageOptions struct {
	Alpha               float64
	FitToWidth          bool
	PreserveAspectRatio bool
}

// CompliancePdfReader is a wrapper over PdfReader that is used for verifying if the input Pdf document matches the
// compliance rules of standards like PDF/A.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type CompliancePdfReader struct {
	*PdfReader
	_fgafb _eb.ParserMetadata
}

func _fafe(_bgfbg *_eb.PdfObjectDictionary, _eefca *fontCommon) (*pdfFontType3, error) {
	_cebcd := _adbdb(_eefca)
	_cgcc := _bgfbg.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
	if _cgcc == nil {
		_cgcc = _eb.MakeInteger(0)
	}
	_cebcd.FirstChar = _cgcc
	_bddb, _agbe := _eb.GetIntVal(_cgcc)
	if !_agbe {
		_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _cgcc)
		return nil, _eb.ErrTypeError
	}
	_cgafa := _fc.CharCode(_bddb)
	_cgcc = _bgfbg.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
	if _cgcc == nil {
		_cgcc = _eb.MakeInteger(255)
	}
	_cebcd.LastChar = _cgcc
	_bddb, _agbe = _eb.GetIntVal(_cgcc)
	if !_agbe {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _cgcc)
		return nil, _eb.ErrTypeError
	}
	_egba := _fc.CharCode(_bddb)
	_cgcc = _bgfbg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _cgcc != nil {
		_cebcd.Resources = _cgcc
	}
	_cgcc = _bgfbg.Get("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s")
	if _cgcc == nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0068\u0061\u0072\u0050\u0072\u006f\u0063\u0073\u0020(%\u0076\u0029", _cgcc)
		return nil, _eb.ErrNotSupported
	}
	_cebcd.CharProcs = _cgcc
	_cgcc = _bgfbg.Get("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078")
	if _cgcc == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0046\u006f\u006et\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0028\u0025\u0076\u0029", _cgcc)
		return nil, _eb.ErrNotSupported
	}
	_cebcd.FontMatrix = _cgcc
	_cebcd._ccfdb = make(map[_fc.CharCode]float64)
	_cgcc = _bgfbg.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
	if _cgcc != nil {
		_cebcd.Widths = _cgcc
		_bdaea, _ffdaf := _eb.GetArray(_cgcc)
		if !_ffdaf {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _cgcc)
			return nil, _eb.ErrTypeError
		}
		_gebeg, _dbdga := _bdaea.ToFloat64Array()
		if _dbdga != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _dbdga
		}
		if len(_gebeg) != int(_egba-_cgafa+1) {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _egba-_cgafa+1, len(_gebeg))
			return nil, _eb.ErrRangeError
		}
		_ebacg, _ffdaf := _eb.GetArray(_cebcd.FontMatrix)
		if !_ffdaf {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _ebacg)
			return nil, _dbdga
		}
		_eced, _dbdga := _ebacg.ToFloat64Array()
		if _dbdga != nil {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020c\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0046o\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _dbdga
		}
		_cddfe := _ffg.NewMatrix(_eced[0], _eced[1], _eced[2], _eced[3], _eced[4], _eced[5])
		for _cbcfb, _gdeae := range _gebeg {
			_fgbc, _ := _cddfe.Transform(_gdeae, _gdeae)
			_cebcd._ccfdb[_cgafa+_fc.CharCode(_cbcfb)] = _fgbc
		}
	}
	_cebcd.Encoding = _eb.TraceToDirectObject(_bgfbg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	_baeeg := _bgfbg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _baeeg != nil {
		_cebcd._geee = _eb.TraceToDirectObject(_baeeg)
		_cefcb, _cdec := _geebd(_cebcd._geee, &_cebcd.fontCommon)
		if _cdec != nil {
			return nil, _cdec
		}
		_cebcd._bgbg = _cefcb
	}
	if _adcfg := _cebcd._bgbg; _adcfg != nil {
		_cebcd._bdde = _fc.NewCMapEncoder("", nil, _adcfg)
	} else {
		_cebcd._bdde = _fc.NewPdfDocEncoder()
	}
	return _cebcd, nil
}

// ToPdfObject implements interface PdfModel.
func (_efac *PdfActionGoTo3DView) ToPdfObject() _eb.PdfObject {
	_efac.PdfAction.ToPdfObject()
	_bge := _efac._dee
	_ca := _bge.PdfObject.(*_eb.PdfObjectDictionary)
	_ca.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeGoTo3DView)))
	_ca.SetIfNotNil("\u0054\u0041", _efac.TA)
	_ca.SetIfNotNil("\u0056", _efac.V)
	return _bge
}
func (_ccbc *PdfReader) buildNameNodes(_agca *_eb.PdfIndirectObject, _eecgf map[_eb.PdfObject]struct{}) error {
	if _agca == nil {
		return nil
	}
	if _, _afbgc := _eecgf[_agca]; _afbgc {
		_ddb.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _agca.ObjectNumber)
		return nil
	}
	_eecgf[_agca] = struct{}{}
	_fgcfg, _gccfc := _agca.PdfObject.(*_eb.PdfObjectDictionary)
	if !_gccfc {
		return _dcf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _bcgeb, _cgea := _eb.GetDict(_fgcfg.Get("\u0044\u0065\u0073t\u0073")); _cgea {
		_fgadf, _affec := _eb.GetArray(_bcgeb.Get("\u004b\u0069\u0064\u0073"))
		if !_affec {
			return _dcf.New("\u0049n\u0076\u0061\u006c\u0069d\u0020\u004b\u0069\u0064\u0073 \u0061r\u0072a\u0079\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_ddb.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _fgadf)
		for _acdgf, _acfaa := range _fgadf.Elements() {
			_adebb, _fcacb := _eb.GetIndirect(_acfaa)
			if !_fcacb {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u0068\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002d \u0028\u0025\u0073\u0029", _adebb)
				return _dcf.New("\u0063h\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006ed\u0069r\u0065c\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
			}
			_fgadf.Set(_acdgf, _adebb)
			_adgf := _ccbc.buildNameNodes(_adebb, _eecgf)
			if _adgf != nil {
				return _adgf
			}
		}
	}
	if _dbfbd, _fggeg := _eb.GetDict(_fgcfg); _fggeg {
		if !_eb.IsNullObject(_dbfbd.Get("\u004b\u0069\u0064\u0073")) {
			if _aagba, _abbcd := _eb.GetArray(_dbfbd.Get("\u004b\u0069\u0064\u0073")); _abbcd {
				for _efacf, _faedb := range _aagba.Elements() {
					if _afdaea, _eaabf := _eb.GetIndirect(_faedb); _eaabf {
						_aagba.Set(_efacf, _afdaea)
						_ffeee := _ccbc.buildNameNodes(_afdaea, _eecgf)
						if _ffeee != nil {
							return _ffeee
						}
					}
				}
			}
		}
	}
	return nil
}

// SetBoundingBox sets the bounding box in the attribute object.
func (_defdgd *KDict) SetBoundingBox(x, y, width, height float64) {
	_defdgd._aagfc = &PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}

// NewGrayImageFromGoImage creates a new grayscale unidoc Image from a golang Image.
func (_edgeda DefaultImageHandler) NewGrayImageFromGoImage(goimg _fb.Image) (*Image, error) {
	_fefadc := goimg.Bounds()
	_deceb := &Image{Width: int64(_fefadc.Dx()), Height: int64(_fefadc.Dy()), ColorComponents: 1, BitsPerComponent: 8}
	switch _adccg := goimg.(type) {
	case *_fb.Gray:
		if len(_adccg.Pix) != _fefadc.Dx()*_fefadc.Dy() {
			_dbcgg, _dgde := _df.GrayConverter.Convert(goimg)
			if _dgde != nil {
				return nil, _dgde
			}
			_deceb.Data = _dbcgg.Pix()
		} else {
			_deceb.Data = _adccg.Pix
		}
	case *_fb.Gray16:
		_deceb.BitsPerComponent = 16
		if len(_adccg.Pix) != _fefadc.Dx()*_fefadc.Dy()*2 {
			_baef, _bgfd := _df.Gray16Converter.Convert(goimg)
			if _bgfd != nil {
				return nil, _bgfd
			}
			_deceb.Data = _baef.Pix()
		} else {
			_deceb.Data = _adccg.Pix
		}
	case _df.Image:
		_ecbad := _adccg.Base()
		if _ecbad.ColorComponents == 1 {
			_deceb.BitsPerComponent = int64(_ecbad.BitsPerComponent)
			_deceb.Data = _ecbad.Data
			return _deceb, nil
		}
		_bggff, _feaf := _df.GrayConverter.Convert(goimg)
		if _feaf != nil {
			return nil, _feaf
		}
		_deceb.Data = _bggff.Pix()
	default:
		_bdbb, _gcbdd := _df.GrayConverter.Convert(goimg)
		if _gcbdd != nil {
			return nil, _gcbdd
		}
		_deceb.Data = _bdbb.Pix()
	}
	return _deceb, nil
}

// SetDocInfo sets the document /Info metadata.
// This will overwrite any globally declared document info.
func (_adaa *PdfAppender) SetDocInfo(info *PdfInfo) { _adaa._ccag = info }

// NewOutlineItem returns a new outline item instance.
func NewOutlineItem(title string, dest OutlineDest) *OutlineItem {
	return &OutlineItem{Title: title, Dest: dest}
}
func (_becf *PdfReader) newPdfActionHideFromDict(_fgc *_eb.PdfObjectDictionary) (*PdfActionHide, error) {
	return &PdfActionHide{T: _fgc.Get("\u0054"), H: _fgc.Get("\u0048")}, nil
}

// NewBorderStyle returns an initialized PdfBorderStyle.
func NewBorderStyle() *PdfBorderStyle { _fgd := &PdfBorderStyle{}; return _fgd }
func _eccc(_dgcg *_eb.PdfObjectStream) (*PdfFunctionType0, error) {
	_daecf := &PdfFunctionType0{}
	_daecf._bded = _dgcg
	_cdbaa := _dgcg.PdfObjectDictionary
	_ecbdg, _eedg := _eb.TraceToDirectObject(_cdbaa.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_eb.PdfObjectArray)
	if !_eedg {
		_ddb.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _ecbdg.Len() < 0 || _ecbdg.Len()%2 != 0 {
		_ddb.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_daecf.NumInputs = _ecbdg.Len() / 2
	_dabcc, _gecg := _ecbdg.ToFloat64Array()
	if _gecg != nil {
		return nil, _gecg
	}
	_daecf.Domain = _dabcc
	_ecbdg, _eedg = _eb.TraceToDirectObject(_cdbaa.Get("\u0052\u0061\u006eg\u0065")).(*_eb.PdfObjectArray)
	if !_eedg {
		_ddb.Log.Error("\u0052\u0061\u006e\u0067e \u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _ecbdg.Len() < 0 || _ecbdg.Len()%2 != 0 {
		return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_daecf.NumOutputs = _ecbdg.Len() / 2
	_effe, _gecg := _ecbdg.ToFloat64Array()
	if _gecg != nil {
		return nil, _gecg
	}
	_daecf.Range = _effe
	_ecbdg, _eedg = _eb.TraceToDirectObject(_cdbaa.Get("\u0053\u0069\u007a\u0065")).(*_eb.PdfObjectArray)
	if !_eedg {
		_ddb.Log.Error("\u0053i\u007ae\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_aecg, _gecg := _ecbdg.ToIntegerArray()
	if _gecg != nil {
		return nil, _gecg
	}
	if len(_aecg) != _daecf.NumInputs {
		_ddb.Log.Error("T\u0061\u0062\u006c\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063h\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072 o\u0066\u0020\u0069n\u0070u\u0074\u0073")
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_daecf.Size = _aecg
	_caeg, _eedg := _eb.TraceToDirectObject(_cdbaa.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065")).(*_eb.PdfObjectInteger)
	if !_eedg {
		_ddb.Log.Error("B\u0069\u0074\u0073\u0050\u0065\u0072S\u0061\u006d\u0070\u006c\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if *_caeg != 1 && *_caeg != 2 && *_caeg != 4 && *_caeg != 8 && *_caeg != 12 && *_caeg != 16 && *_caeg != 24 && *_caeg != 32 {
		_ddb.Log.Error("\u0042\u0069\u0074s \u0070\u0065\u0072\u0020\u0073\u0061\u006d\u0070\u006ce\u0020o\u0075t\u0073i\u0064\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064\u0029", *_caeg)
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_daecf.BitsPerSample = int(*_caeg)
	_daecf.Order = 1
	_acaag, _eedg := _eb.TraceToDirectObject(_cdbaa.Get("\u004f\u0072\u0064e\u0072")).(*_eb.PdfObjectInteger)
	if _eedg {
		if *_acaag != 1 && *_acaag != 3 {
			_ddb.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u006fr\u0064\u0065r\u0020\u0028\u0025\u0064\u0029", *_acaag)
			return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		_daecf.Order = int(*_acaag)
	}
	_ecbdg, _eedg = _eb.TraceToDirectObject(_cdbaa.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_eb.PdfObjectArray)
	if _eedg {
		_gceff, _gddga := _ecbdg.ToFloat64Array()
		if _gddga != nil {
			return nil, _gddga
		}
		_daecf.Encode = _gceff
	}
	_ecbdg, _eedg = _eb.TraceToDirectObject(_cdbaa.Get("\u0044\u0065\u0063\u006f\u0064\u0065")).(*_eb.PdfObjectArray)
	if _eedg {
		_eaec, _agfef := _ecbdg.ToFloat64Array()
		if _agfef != nil {
			return nil, _agfef
		}
		_daecf.Decode = _eaec
	}
	_cecfc, _gecg := _eb.DecodeStream(_dgcg)
	if _gecg != nil {
		return nil, _gecg
	}
	_daecf._dcdgc = _cecfc
	return _daecf, nil
}

// PdfColorspaceDeviceN represents a DeviceN color space. DeviceN color spaces are similar to Separation color
// spaces, except they can contain an arbitrary number of color components.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
        or: [/DeviceN names alternateSpace tintTransform attributes]
*/
type PdfColorspaceDeviceN struct {
	ColorantNames  *_eb.PdfObjectArray
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	Attributes     *PdfColorspaceDeviceNAttributes
	_cfef          *_eb.PdfIndirectObject
}

// AddCerts adds certificates to DSS.
func (_agfe *DSS) AddCerts(certs [][]byte) ([]*_eb.PdfObjectStream, error) {
	return _agfe.add(&_agfe.Certs, _agfe._eaed, certs)
}
func (_aaeb *PdfReader) newPdfAnnotationStampFromDict(_cadf *_eb.PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	_ecec := PdfAnnotationStamp{}
	_degea, _ggab := _aaeb.newPdfAnnotationMarkupFromDict(_cadf)
	if _ggab != nil {
		return nil, _ggab
	}
	_ecec.PdfAnnotationMarkup = _degea
	_ecec.Name = _cadf.Get("\u004e\u0061\u006d\u0065")
	return &_ecec, nil
}

// ValidateSignatures validates digital signatures in the document.
func (_edfg *PdfReader) ValidateSignatures(handlers []SignatureHandler) ([]SignatureValidationResult, error) {
	if _edfg.AcroForm == nil {
		return nil, nil
	}
	if _edfg.AcroForm.Fields == nil {
		return nil, nil
	}
	type sigFieldPair struct {
		_cgfba *PdfSignature
		_gdffd *PdfField
		_bgafg SignatureHandler
	}
	var _ddfdc []*sigFieldPair
	for _, _afebb := range _edfg.AcroForm.AllFields() {
		if _afebb.V == nil {
			continue
		}
		if _aceff, _bece := _eb.GetDict(_afebb.V); _bece {
			if _dacfa, _dbfdf := _eb.GetNameVal(_aceff.Get("\u0054\u0079\u0070\u0065")); _dbfdf && (_dacfa == "\u0053\u0069\u0067" || _dacfa == "\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070") {
				_dadbf, _cdffbf := _eb.GetIndirect(_afebb.V)
				if !_cdffbf {
					_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0020\u0069s\u0020\u006e\u0069\u006c")
					return nil, ErrTypeCheck
				}
				_cbdbg, _adgg := _edfg.newPdfSignatureFromIndirect(_dadbf)
				if _adgg != nil {
					return nil, _adgg
				}
				var _afaaf SignatureHandler
				for _, _afcab := range handlers {
					if _afcab.IsApplicable(_cbdbg) {
						_afaaf = _afcab
						break
					}
				}
				_ddfdc = append(_ddfdc, &sigFieldPair{_cgfba: _cbdbg, _gdffd: _afebb, _bgafg: _afaaf})
			}
		}
	}
	var _bbba []SignatureValidationResult
	for _, _deed := range _ddfdc {
		_faccc := SignatureValidationResult{IsSigned: true, Fields: []*PdfField{_deed._gdffd}}
		if _deed._bgafg == nil {
			_faccc.Errors = append(_faccc.Errors, "\u0068a\u006ed\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
			_bbba = append(_bbba, _faccc)
			continue
		}
		_dbaga, _fedcce := _deed._bgafg.NewDigest(_deed._cgfba)
		if _fedcce != nil {
			_faccc.Errors = append(_faccc.Errors, "\u0064\u0069\u0067e\u0073\u0074\u0020\u0065\u0072\u0072\u006f\u0072", _fedcce.Error())
			_bbba = append(_bbba, _faccc)
			continue
		}
		_dffg := _deed._cgfba.ByteRange
		if _dffg == nil {
			_faccc.Errors = append(_faccc.Errors, "\u0042\u0079\u0074\u0065\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
			_bbba = append(_bbba, _faccc)
			continue
		}
		for _eabgc := 0; _eabgc < _dffg.Len(); _eabgc = _eabgc + 2 {
			_ffdafb, _ := _eb.GetNumberAsInt64(_dffg.Get(_eabgc))
			_cefd, _ := _eb.GetIntVal(_dffg.Get(_eabgc + 1))
			if _, _ffccg := _edfg._cbeg.Seek(_ffdafb, _bagf.SeekStart); _ffccg != nil {
				return nil, _ffccg
			}
			_eaccd := make([]byte, _cefd)
			if _, _cccgb := _edfg._cbeg.Read(_eaccd); _cccgb != nil {
				return nil, _cccgb
			}
			_dbaga.Write(_eaccd)
		}
		var _dfgebe SignatureValidationResult
		if _dceg, _efag := _deed._bgafg.(SignatureHandlerDocMDP); _efag {
			_dfgebe, _fedcce = _dceg.ValidateWithOpts(_deed._cgfba, _dbaga, SignatureHandlerDocMDPParams{Parser: _edfg._ebbe})
		} else {
			_dfgebe, _fedcce = _deed._bgafg.Validate(_deed._cgfba, _dbaga)
		}
		if _fedcce != nil {
			_ddb.Log.Debug("E\u0052\u0052\u004f\u0052: \u0025v\u0020\u0028\u0025\u0054\u0029 \u002d\u0020\u0073\u006b\u0069\u0070", _fedcce, _deed._bgafg)
			_dfgebe.Errors = append(_dfgebe.Errors, _fedcce.Error())
		}
		_dfgebe.Name = _deed._cgfba.Name.Decoded()
		_dfgebe.Reason = _deed._cgfba.Reason.Decoded()
		if _deed._cgfba.M != nil {
			_ecadf, _fgceg := NewPdfDate(_deed._cgfba.M.String())
			if _fgceg != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fgceg)
				_dfgebe.Errors = append(_dfgebe.Errors, _fgceg.Error())
				continue
			}
			_dfgebe.Date = _ecadf
		}
		_dfgebe.ContactInfo = _deed._cgfba.ContactInfo.Decoded()
		_dfgebe.Location = _deed._cgfba.Location.Decoded()
		_dfgebe.Fields = _faccc.Fields
		_bbba = append(_bbba, _dfgebe)
	}
	return _bbba, nil
}

// L returns the value of the L component of the color.
func (_cdbe *PdfColorLab) L() float64 { return _cdbe[0] }

const (
	ActionTypeGoTo        PdfActionType = "\u0047\u006f\u0054\u006f"
	ActionTypeGoTo3DView  PdfActionType = "\u0047\u006f\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077"
	ActionTypeGoToE       PdfActionType = "\u0047\u006f\u0054o\u0045"
	ActionTypeGoToR       PdfActionType = "\u0047\u006f\u0054o\u0052"
	ActionTypeHide        PdfActionType = "\u0048\u0069\u0064\u0065"
	ActionTypeImportData  PdfActionType = "\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061"
	ActionTypeJavaScript  PdfActionType = "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"
	ActionTypeLaunch      PdfActionType = "\u004c\u0061\u0075\u006e\u0063\u0068"
	ActionTypeMovie       PdfActionType = "\u004d\u006f\u0076i\u0065"
	ActionTypeNamed       PdfActionType = "\u004e\u0061\u006de\u0064"
	ActionTypeRendition   PdfActionType = "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn"
	ActionTypeResetForm   PdfActionType = "\u0052e\u0073\u0065\u0074\u0046\u006f\u0072m"
	ActionTypeSetOCGState PdfActionType = "S\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065"
	ActionTypeSound       PdfActionType = "\u0053\u006f\u0075n\u0064"
	ActionTypeSubmitForm  PdfActionType = "\u0053\u0075\u0062\u006d\u0069\u0074\u0046\u006f\u0072\u006d"
	ActionTypeThread      PdfActionType = "\u0054\u0068\u0072\u0065\u0061\u0064"
	ActionTypeTrans       PdfActionType = "\u0054\u0072\u0061n\u0073"
	ActionTypeURI         PdfActionType = "\u0055\u0052\u0049"
)

// ImageToRGB converts Lab colorspace image to RGB and returns the result.
func (_caea *PdfColorspaceLab) ImageToRGB(img Image) (Image, error) {
	_dcbf := func(_cecb float64) float64 {
		if _cecb >= 6.0/29 {
			return _cecb * _cecb * _cecb
		}
		return 108.0 / 841 * (_cecb - 4.0/29.0)
	}
	_edgf := img._fedc
	if len(_edgf) != 6 {
		_ddb.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u002d\u0020\u004c\u0061\u0062\u0020\u0044e\u0063\u006f\u0064\u0065\u0020\u0072\u0061\u006e\u0067e\u0020\u0021\u003d\u0020\u0036\u002e\u002e\u002e\u0020\u0075\u0073\u0065\u0020\u005b0\u0020\u0031\u0030\u0030\u0020\u0061\u006d\u0069\u006e\u0020\u0061\u006d\u0061\u0078\u0020\u0062\u006d\u0069\u006e\u0020\u0062\u006d\u0061\u0078\u005d\u0020\u0064\u0065\u0066\u0061u\u006c\u0074\u0020\u0064\u0065\u0063\u006f\u0064\u0065 \u0061\u0072r\u0061\u0079")
		_edgf = _caea.DecodeArray()
	}
	_edec := _bb.NewReader(img.getBase())
	_debgb := _df.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, img._bdcab, img._fedc)
	_egec := _bb.NewWriter(_debgb)
	_ddgdb := _gg.Pow(2, float64(img.BitsPerComponent)) - 1
	_caad := make([]uint32, 3)
	var (
		_fbfd                                            error
		Ls, As, Bs, L, M, N, X, Y, Z, _adbd, _ceeg, _gbc float64
	)
	for {
		_fbfd = _edec.ReadSamples(_caad)
		if _fbfd == _bagf.EOF {
			break
		} else if _fbfd != nil {
			return img, _fbfd
		}
		Ls = float64(_caad[0]) / _ddgdb
		As = float64(_caad[1]) / _ddgdb
		Bs = float64(_caad[2]) / _ddgdb
		Ls = _df.LinearInterpolate(Ls, 0.0, 1.0, _edgf[0], _edgf[1])
		As = _df.LinearInterpolate(As, 0.0, 1.0, _edgf[2], _edgf[3])
		Bs = _df.LinearInterpolate(Bs, 0.0, 1.0, _edgf[4], _edgf[5])
		L = (Ls+16)/116 + As/500
		M = (Ls + 16) / 116
		N = (Ls+16)/116 - Bs/200
		X = _caea.WhitePoint[0] * _dcbf(L)
		Y = _caea.WhitePoint[1] * _dcbf(M)
		Z = _caea.WhitePoint[2] * _dcbf(N)
		_adbd = 3.240479*X + -1.537150*Y + -0.498535*Z
		_ceeg = -0.969256*X + 1.875992*Y + 0.041556*Z
		_gbc = 0.055648*X + -0.204043*Y + 1.057311*Z
		_adbd = _gg.Min(_gg.Max(_adbd, 0), 1.0)
		_ceeg = _gg.Min(_gg.Max(_ceeg, 0), 1.0)
		_gbc = _gg.Min(_gg.Max(_gbc, 0), 1.0)
		_caad[0] = uint32(_adbd * _ddgdb)
		_caad[1] = uint32(_ceeg * _ddgdb)
		_caad[2] = uint32(_gbc * _ddgdb)
		if _fbfd = _egec.WriteSamples(_caad); _fbfd != nil {
			return img, _fbfd
		}
	}
	return _ggaa(&_debgb), nil
}
func NewViewerPreferencesFromPdfObject(obj _eb.PdfObject) (*ViewerPreferences, error) {
	_ccgbgd := _eb.ResolveReference(obj)
	_bfbcg, _fbaaa := _eb.GetDict(_ccgbgd)
	if !_fbaaa {
		return nil, _e.Errorf("e\u0078\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0076i\u0065\u0077\u0065\u0072\u0020\u0070\u0072ef\u0065\u0072\u0065\u006ec\u0065\u0073\u0020\u0069\u0073\u0020\u006e\u006f\u0074 a\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_facd := NewViewerPreferences()
	if _afcg := _bfbcg.Get("H\u0069\u0064\u0065\u0054\u006f\u006f\u006c\u0062\u0061\u0072"); _afcg != nil {
		if _gceg, _caade := _eb.GetBool(_afcg); _caade {
			_facd._caafc = (*bool)(_gceg)
		}
	}
	if _bcfbb := _bfbcg.Get("H\u0069\u0064\u0065\u004d\u0065\u006e\u0075\u0062\u0061\u0072"); _bcfbb != nil {
		if _bdfe, _fefcg := _eb.GetBool(_bcfbb); _fefcg {
			_facd._gcbcb = (*bool)(_bdfe)
		}
	}
	if _edcce := _bfbcg.Get("\u0048\u0069\u0064e\u0057\u0069\u006e\u0064\u006f\u0077\u0055\u0049"); _edcce != nil {
		if _addgd, _adbdd := _eb.GetBool(_edcce); _adbdd {
			_facd._fcebae = (*bool)(_addgd)
		}
	}
	if _faggd := _bfbcg.Get("\u0046i\u0074\u0057\u0069\u006e\u0064\u006fw"); _faggd != nil {
		if _abga, _ebaf := _eb.GetBool(_faggd); _ebaf {
			_facd._edbaa = (*bool)(_abga)
		}
	}
	if _deeeg := _bfbcg.Get("\u0043\u0065\u006et\u0065\u0072\u0057\u0069\u006e\u0064\u006f\u0077"); _deeeg != nil {
		if _abdda, _edded := _eb.GetBool(_deeeg); _edded {
			_facd._dacc = (*bool)(_abdda)
		}
	}
	if _gadfb := _bfbcg.Get("\u0044i\u0073p\u006c\u0061\u0079\u0044\u006f\u0063\u0054\u0069\u0074\u006c\u0065"); _gadfb != nil {
		if _bdfbae, _degddd := _eb.GetBool(_gadfb); _degddd {
			_facd._bgdge = (*bool)(_bdfbae)
		}
	}
	if _fecdg := _bfbcg.Get("N\u006f\u006e\u0046\u0075ll\u0053c\u0072\u0065\u0065\u006e\u0050a\u0067\u0065\u004d\u006f\u0064\u0065"); _fecdg != nil {
		if _afgdf, _faaef := _eb.GetName(_fecdg); _faaef {
			_facd._cdcff = NonFullScreenPageMode(*_afgdf)
		}
	}
	if _dcbbga := _bfbcg.Get("\u0044i\u0072\u0065\u0063\u0074\u0069\u006fn"); _dcbbga != nil {
		if _gabf, _eabec := _eb.GetName(_dcbbga); _eabec {
			_facd._baeb = Direction(*_gabf)
		}
	}
	if _fefbf := _bfbcg.Get("\u0056\u0069\u0065\u0077\u0041\u0072\u0065\u0061"); _fefbf != nil {
		if _begcb, _eaggb := _eb.GetName(_fefbf); _eaggb {
			_facd._ffgda = PageBoundary(*_begcb)
		}
	}
	if _ecfd := _bfbcg.Get("\u0056\u0069\u0065\u0077\u0043\u006c\u0069\u0070"); _ecfd != nil {
		if _gedd, _eagcd := _eb.GetName(_ecfd); _eagcd {
			_facd._eadae = PageBoundary(*_gedd)
		}
	}
	if _gfeb := _bfbcg.Get("\u0050r\u0069\u006e\u0074\u0041\u0072\u0065a"); _gfeb != nil {
		if _bbaee, _ecae := _eb.GetName(_gfeb); _ecae {
			_facd._bdegd = PageBoundary(*_bbaee)
		}
	}
	if _gfbefc := _bfbcg.Get("\u0050r\u0069\u006e\u0074\u0043\u006c\u0069p"); _gfbefc != nil {
		if _fcfbb, _afaeb := _eb.GetName(_gfbefc); _afaeb {
			_facd._edef = PageBoundary(*_fcfbb)
		}
	}
	if _fgafbb := _bfbcg.Get("\u0050\u0072\u0069n\u0074\u0053\u0063\u0061\u006c\u0069\u006e\u0067"); _fgafbb != nil {
		if _dadee, _ecdfa := _eb.GetName(_fgafbb); _ecdfa {
			_facd._gefcg = PrintScaling(*_dadee)
		}
	}
	if _fbedgg := _bfbcg.Get("\u0044\u0075\u0070\u006c\u0065\u0078"); _fbedgg != nil {
		if _ddbea, _egbdg := _eb.GetName(_fbedgg); _egbdg {
			_facd._ebabd = Duplex(*_ddbea)
		}
	}
	if _acagf := _bfbcg.Get("\u0050\u0069\u0063\u006b\u0054\u0072\u0061\u0079\u0042\u0079\u0050\u0044F\u0053\u0069\u007a\u0065"); _acagf != nil {
		if _gfdcc, _egaeb := _eb.GetBool(_acagf); _egaeb {
			_facd._gggea = (*bool)(_gfdcc)
		}
	}
	if _gbgba := _bfbcg.Get("\u0050\u0072\u0069\u006e\u0074\u0050\u0061\u0067\u0065R\u0061\u006e\u0067\u0065"); _gbgba != nil {
		if _eggae, _eacbd := _eb.GetArray(_gbgba); _eacbd {
			_facd._gcgfc = make([]int, _eggae.Len())
			for _bdfccd := range _facd._gcgfc {
				if _cgddf := _eggae.Get(_bdfccd); _cgddf != nil {
					if _bacgf, _bbbaa := _eb.GetInt(_cgddf); _bbbaa {
						_facd._gcgfc[_bdfccd] = int(*_bacgf)
					}
				}
			}
		}
	}
	if _dbaca := _bfbcg.Get("\u004eu\u006d\u0043\u006f\u0070\u0069\u0065s"); _dbaca != nil {
		if _egcc, _abddd := _eb.GetInt(_dbaca); _abddd {
			_facd._bdeec = int(*_egcc)
		}
	}
	return _facd, nil
}

// ToUnicode returns the name of the font's "ToUnicode" field if there is one, or "" if there isn't.
func (_bcca *PdfFont) ToUnicode() string {
	if _bcca.baseFields()._bgbg == nil {
		return ""
	}
	return _bcca.baseFields()._bgbg.Name()
}

// EnableAll LTV enables all signatures in the PDF document.
// The signing certificate chain is extracted from each signature dictionary.
// Optionally, additional certificates can be specified through the
// `extraCerts` parameter. The LTV client attempts to build the certificate
// chain up to a trusted root by downloading any missing certificates.
func (_bafgec *LTV) EnableAll(extraCerts []*_bag.Certificate) error {
	_cbgcc := _bafgec._feba._adac.AcroForm
	for _, _dbac := range _cbgcc.AllFields() {
		_gcddc, _ := _dbac.GetContext().(*PdfFieldSignature)
		if _gcddc == nil {
			continue
		}
		_egfc := _gcddc.V
		if _cggd := _bafgec.validateSig(_egfc); _cggd != nil {
			_ddb.Log.Debug("\u0057\u0041\u0052N\u003a\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _cggd)
		}
		if _eebgb := _bafgec.Enable(_egfc, extraCerts); _eebgb != nil {
			return _eebgb
		}
	}
	return nil
}
func (_bgafa *PdfWriter) writeObjectsInStreams(_fbegg map[_eb.PdfObject]bool) error {
	for _, _babgba := range _bgafa._dcfgf {
		if _fdbcb := _fbegg[_babgba]; _fdbcb {
			continue
		}
		_cbfde := int64(0)
		switch _cdbcd := _babgba.(type) {
		case *_eb.PdfIndirectObject:
			_cbfde = _cdbcd.ObjectNumber
		case *_eb.PdfObjectStream:
			_cbfde = _cdbcd.ObjectNumber
		case *_eb.PdfObjectStreams:
			_cbfde = _cdbcd.ObjectNumber
		case *_eb.PdfObjectDictionary, *_eb.PdfObjectString:
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0073\u0075p\u0070\u006f\u0072t\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069n\u0020wr\u0069\u0074\u0065\u0072\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0054\u0020\u0028\u0074\u0079\u0070\u0065\u0020\u0025\u0054\u0029", _babgba, _cdbcd)
			return ErrTypeCheck
		}
		if _bgafa._gadcaa != nil && _babgba != _bgafa._ebdeed {
			_eage := _bgafa._gadcaa.Encrypt(_babgba, _cbfde, 0)
			if _eage != nil {
				_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067\u0020(%\u0073\u0029", _eage)
				return _eage
			}
		}
		_bgafa.writeObject(int(_cbfde), _babgba)
	}
	return nil
}
func (_aadbb *DSS) generateHashMap(_fdedd []*_eb.PdfObjectStream) (map[string]*_eb.PdfObjectStream, error) {
	_gcfg := map[string]*_eb.PdfObjectStream{}
	for _, _eeab := range _fdedd {
		_fggdd, _dddc := _eb.DecodeStream(_eeab)
		if _dddc != nil {
			return nil, _dddc
		}
		_begg, _dddc := _bceda(_fggdd)
		if _dddc != nil {
			return nil, _dddc
		}
		_gcfg[string(_begg)] = _eeab
	}
	return _gcfg, nil
}
func _ecff(_bafge _eb.PdfObject) (*PdfColorspaceICCBased, error) {
	_cfgfd := &PdfColorspaceICCBased{}
	if _bdbec, _ageb := _bafge.(*_eb.PdfIndirectObject); _ageb {
		_cfgfd._ebggg = _bdbec
	}
	_bafge = _eb.TraceToDirectObject(_bafge)
	_bfcd, _efcc := _bafge.(*_eb.PdfObjectArray)
	if !_efcc {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bfcd.Len() != 2 {
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020c\u006f\u006c\u006fr\u0073p\u0061\u0063\u0065")
	}
	_bafge = _eb.TraceToDirectObject(_bfcd.Get(0))
	_dcfd, _efcc := _bafge.(*_eb.PdfObjectName)
	if !_efcc {
		return nil, _e.Errorf("\u0049\u0043\u0043B\u0061\u0073\u0065\u0064 \u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_dcfd != "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064" {
		return nil, _e.Errorf("\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0049\u0043\u0043\u0042a\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_bafge = _bfcd.Get(1)
	_bceca, _efcc := _eb.GetStream(_bafge)
	if !_efcc {
		_ddb.Log.Error("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006e\u006f\u0074\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067\u0020t\u006f\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020%\u0054", _bafge)
		return nil, _e.Errorf("\u0049\u0043\u0043Ba\u0073\u0065\u0064\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_abab := _bceca.PdfObjectDictionary
	_bcggc, _efcc := _abab.Get("\u004e").(*_eb.PdfObjectInteger)
	if !_efcc {
		return nil, _e.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006d\u0069\u0073\u0073\u0069n\u0067\u0020\u004e\u0020\u0066\u0072\u006fm\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074")
	}
	if *_bcggc != 1 && *_bcggc != 3 && *_bcggc != 4 {
		return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061s\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u006e\u006f\u0074\u0020\u0031\u002c\u0033\u002c\u0034\u0029")
	}
	_cfgfd.N = int(*_bcggc)
	if _beda := _abab.Get("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e"); _beda != nil {
		_gcag, _ddfd := NewPdfColorspaceFromPdfObject(_beda)
		if _ddfd != nil {
			return nil, _ddfd
		}
		_cfgfd.Alternate = _gcag
	}
	if _eabb := _abab.Get("\u0052\u0061\u006eg\u0065"); _eabb != nil {
		_eabb = _eb.TraceToDirectObject(_eabb)
		_fcgb, _cedcg := _eabb.(*_eb.PdfObjectArray)
		if !_cedcg {
			return nil, _e.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0061\u006e \u0061r\u0072\u0061\u0079")
		}
		if _fcgb.Len() != 2*_cfgfd.N {
			return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0052\u0061\u006e\u0067e\u0020\u0077\u0072\u006f\u006e\u0067 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073")
		}
		_agac, _bbfd := _fcgb.GetAsFloat64Slice()
		if _bbfd != nil {
			return nil, _bbfd
		}
		_cfgfd.Range = _agac
	} else {
		_cfgfd.Range = make([]float64, 2*_cfgfd.N)
		for _bfba := 0; _bfba < _cfgfd.N; _bfba++ {
			_cfgfd.Range[2*_bfba] = 0.0
			_cfgfd.Range[2*_bfba+1] = 1.0
		}
	}
	if _adfeg := _abab.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _adfeg != nil {
		_gbbg, _dbcb := _adfeg.(*_eb.PdfObjectStream)
		if !_dbcb {
			return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u004de\u0074\u0061\u0064\u0061\u0074\u0061\u0020n\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		_cfgfd.Metadata = _gbbg
	}
	_efgcf, _bfdd := _eb.DecodeStream(_bceca)
	if _bfdd != nil {
		return nil, _bfdd
	}
	_cfgfd.Data = _efgcf
	_cfgfd._cdebc = _bceca
	return _cfgfd, nil
}

var (
	StructureTypeSpan               = "\u0053\u0070\u0061\u006e"
	StructureTypeQuote              = "\u0051\u0075\u006ft\u0065"
	StructureTypeNote               = "\u004e\u006f\u0074\u0065"
	StructureTypeReference          = "\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e"
	StructureTypeBibliography       = "\u0042\u0069\u0062\u0045\u006e\u0074\u0072\u0079"
	StructureTypeCode               = "\u0043\u006f\u0064\u0065"
	StructureTypeLink               = "\u004c\u0069\u006e\u006b"
	StructureTypeAnnot              = "\u0041\u006e\u006eo\u0074"
	StructureTypeRuby               = "\u0052\u0075\u0062\u0079"
	StructureTypeWarichu            = "\u0057a\u0072\u0069\u0063\u0068\u0075"
	StructureTypeRubyBase           = "\u0052\u0042"
	StructureTypeRubyText           = "\u0052\u0054"
	StructureTypeRubyPunctuation    = "\u0052\u0050"
	StructureTypeWarichuText        = "\u0057\u0054"
	StructureTypeWarichuPunctuation = "\u0057\u0050"
	StructureTypeFigure             = "\u0046\u0069\u0067\u0075\u0072\u0065"
	StructureTypeFormula            = "\u0046o\u0072\u006d\u0075\u006c\u0061"
	StructureTypeForm               = "\u0046\u006f\u0072\u006d"
)

// GetContentStream returns the pattern cell's content stream
func (_bbfbd *PdfTilingPattern) GetContentStream() ([]byte, error) {
	_bbcbfc, _, _fbdce := _bbfbd.GetContentStreamWithEncoder()
	return _bbcbfc, _fbdce
}

// ToOutlineTree returns a low level PdfOutlineTreeNode object, based on
// the current instance.
func (_bagdc *Outline) ToOutlineTree() *PdfOutlineTreeNode {
	return &_bagdc.ToPdfOutline().PdfOutlineTreeNode
}
func (_cgfbe PdfFont) actualFont() pdfFont {
	if _cgfbe._fdaa == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0061\u0063\u0074\u0075\u0061\u006c\u0046\u006f\u006e\u0074\u002e\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c.\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _cgfbe)
	}
	return _cgfbe._fdaa
}

// FieldValueProvider provides field values from a data source such as FDF, JSON or any other.
type FieldValueProvider interface {
	FieldValues() (map[string]_eb.PdfObject, error)
}

// GetContainingPdfObject returns the containing object for the PdfField, i.e. an indirect object
// containing the field dictionary.
func (_ffgd *PdfField) GetContainingPdfObject() _eb.PdfObject { return _ffgd._adgda }

// Encoder returns the font's text encoder.
func (_ecbf *PdfFont) Encoder() _fc.TextEncoder {
	_cbbff := _ecbf.actualFont()
	if _cbbff == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0045n\u0063\u006f\u0064er\u0020\u006e\u006f\u0074\u0020\u0069m\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066o\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d%\u0023\u0054", _ecbf._fdaa)
		return nil
	}
	return _cbbff.Encoder()
}

// ToPdfObject implements interface PdfModel.
func (_eade *PdfAnnotation3D) ToPdfObject() _eb.PdfObject {
	_eade.PdfAnnotation.ToPdfObject()
	_adfa := _eade._ggf
	_fcee := _adfa.PdfObject.(*_eb.PdfObjectDictionary)
	_fcee.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0033\u0044"))
	_fcee.SetIfNotNil("\u0033\u0044\u0044", _eade.T3DD)
	_fcee.SetIfNotNil("\u0033\u0044\u0056", _eade.T3DV)
	_fcee.SetIfNotNil("\u0033\u0044\u0041", _eade.T3DA)
	_fcee.SetIfNotNil("\u0033\u0044\u0049", _eade.T3DI)
	_fcee.SetIfNotNil("\u0033\u0044\u0042", _eade.T3DB)
	return _adfa
}

// ToPdfObject implements interface PdfModel.
func (_edg *PdfActionResetForm) ToPdfObject() _eb.PdfObject {
	_edg.PdfAction.ToPdfObject()
	_gdcf := _edg._dee
	_edb := _gdcf.PdfObject.(*_eb.PdfObjectDictionary)
	_edb.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeResetForm)))
	_edb.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _edg.Fields)
	_edb.SetIfNotNil("\u0046\u006c\u0061g\u0073", _edg.Flags)
	return _gdcf
}
func (_bcecba *PdfWriter) writeObject(_bdcbg int, _gbabb _eb.PdfObject) {
	_ddb.Log.Trace("\u0057\u0072\u0069\u0074\u0065\u0020\u006f\u0062\u006a \u0023\u0025\u0064\u000a", _bdcbg)
	if _edaba, _gbegg := _gbabb.(*_eb.PdfIndirectObject); _gbegg {
		_bcecba._aggfdb[_bdcbg] = crossReference{Type: 1, Offset: _bcecba._dfabe, Generation: _edaba.GenerationNumber}
		_geaaa := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bdcbg)
		if _efce, _adfec := _edaba.PdfObject.(*pdfSignDictionary); _adfec {
			_efce._afeabg = _bcecba._dfabe + int64(len(_geaaa))
		}
		if _edaba.PdfObject == nil {
			_ddb.Log.Debug("E\u0072\u0072\u006fr\u003a\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0027\u0073\u0020\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020b\u0065\u0020\u006e\u0069l\u0020\u002d\u0020\u0073e\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063t\u004e\u0075\u006c\u006c")
			_edaba.PdfObject = _eb.MakeNull()
		}
		_geaaa += _edaba.PdfObject.WriteString()
		_geaaa += "\u000a\u0065\u006e\u0064\u006f\u0062\u006a\u000a"
		_bcecba.writeString(_geaaa)
		return
	}
	if _adcfb, _dafece := _gbabb.(*_eb.PdfObjectStream); _dafece {
		_bcecba._aggfdb[_bdcbg] = crossReference{Type: 1, Offset: _bcecba._dfabe, Generation: _adcfb.GenerationNumber}
		_gcdeg := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bdcbg)
		_gcdeg += _adcfb.PdfObjectDictionary.WriteString()
		_gcdeg += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_bcecba.writeString(_gcdeg)
		if _adcfb.Lazy {
			_cdgbb, _ddda := _ccb.ReadFile(_adcfb.TempFile)
			if _ddda != nil {
				_ddb.Log.Info("\u0045\u0072\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006e\u0067\u0020\u006ca\u007ay\u0020\u0074\u0065\u006d\u0070\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _adcfb.TempFile)
				return
			}
			_bcecba.writeBytes(_cdgbb)
			_ccb.Remove(_adcfb.TempFile)
		} else {
			_bcecba.writeBytes(_adcfb.Stream)
		}
		_bcecba.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	if _eabba, _cdadc := _gbabb.(*_eb.PdfObjectStreams); _cdadc {
		_bcecba._aggfdb[_bdcbg] = crossReference{Type: 1, Offset: _bcecba._dfabe, Generation: _eabba.GenerationNumber}
		_fefeaf := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bdcbg)
		var _abgb []string
		var _abced string
		var _fced int64
		for _bcedd, _ggaag := range _eabba.Elements() {
			_egbeda, _cedfd := _ggaag.(*_eb.PdfIndirectObject)
			if !_cedfd {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065am\u0073 \u004e\u0020\u0025\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006es\u0020\u006e\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u0070\u0064\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0076", _bdcbg, _ggaag)
				continue
			}
			_gbbae := _egbeda.PdfObject.WriteString() + "\u0020"
			_abced = _abced + _gbbae
			_abgb = append(_abgb, _e.Sprintf("\u0025\u0064\u0020%\u0064", _egbeda.ObjectNumber, _fced))
			_bcecba._aggfdb[int(_egbeda.ObjectNumber)] = crossReference{Type: 2, ObjectNumber: _bdcbg, Index: _bcedd}
			_fced = _fced + int64(len([]byte(_gbbae)))
		}
		_ccee := _cc.Join(_abgb, "\u0020") + "\u0020"
		_bfbdgc := _eb.NewFlateEncoder()
		_ffaed := _bfbdgc.MakeStreamDict()
		_ffaed.Set(_eb.PdfObjectName("\u0054\u0079\u0070\u0065"), _eb.MakeName("\u004f\u0062\u006a\u0053\u0074\u006d"))
		_ceebg := int64(_eabba.Len())
		_ffaed.Set(_eb.PdfObjectName("\u004e"), _eb.MakeInteger(_ceebg))
		_bada := int64(len(_ccee))
		_ffaed.Set(_eb.PdfObjectName("\u0046\u0069\u0072s\u0074"), _eb.MakeInteger(_bada))
		_gbeba, _ := _bfbdgc.EncodeBytes([]byte(_ccee + _abced))
		_bdcbd := int64(len(_gbeba))
		_ffaed.Set(_eb.PdfObjectName("\u004c\u0065\u006e\u0067\u0074\u0068"), _eb.MakeInteger(_bdcbd))
		_fefeaf += _ffaed.WriteString()
		_fefeaf += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_bcecba.writeString(_fefeaf)
		_bcecba.writeBytes(_gbeba)
		_bcecba.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	_bcecba.writeString(_gbabb.WriteString())
}

// NewPdfAcroForm returns a new PdfAcroForm with an initialized container (indirect object).
func NewPdfAcroForm() *PdfAcroForm {
	return &PdfAcroForm{Fields: &[]*PdfField{}, _fbgad: _eb.MakeIndirectObject(_eb.MakeDict())}
}

// HasExtGState checks if ExtGState name is available.
func (_bbee *PdfPage) HasExtGState(name _eb.PdfObjectName) bool {
	if _bbee.Resources == nil {
		return false
	}
	if _bbee.Resources.ExtGState == nil {
		return false
	}
	_dacbg, _feefg := _eb.TraceToDirectObject(_bbee.Resources.ExtGState).(*_eb.PdfObjectDictionary)
	if !_feefg {
		_ddb.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _eb.TraceToDirectObject(_bbee.Resources.ExtGState))
		return false
	}
	_beagg := _dacbg.Get(name)
	_dafee := _beagg != nil
	return _dafee
}

var _ pdfFont = (*pdfFontType3)(nil)

func _cccfa(_dfgg _eb.PdfObject) (PdfFunction, error) {
	_dfgg = _eb.ResolveReference(_dfgg)
	if _gfdf, _geeef := _dfgg.(*_eb.PdfObjectStream); _geeef {
		_dfbce := _gfdf.PdfObjectDictionary
		_egede, _ggea := _dfbce.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_eb.PdfObjectInteger)
		if !_ggea {
			_ddb.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_egede == 0 {
			return _eccc(_gfdf)
		} else if *_egede == 4 {
			return _fcecd(_gfdf)
		} else {
			return nil, _dcf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _agdfd, _egdb := _dfgg.(*_eb.PdfIndirectObject); _egdb {
		_abfaa, _gdcaa := _agdfd.PdfObject.(*_eb.PdfObjectDictionary)
		if !_gdcaa {
			_ddb.Log.Error("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006eg\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		_dfdc, _gdcaa := _abfaa.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_eb.PdfObjectInteger)
		if !_gdcaa {
			_ddb.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_dfdc == 2 {
			return _ggcf(_agdfd)
		} else if *_dfdc == 3 {
			return _ddeca(_agdfd)
		} else {
			return nil, _dcf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _gdbe, _fcddg := _dfgg.(*_eb.PdfObjectDictionary); _fcddg {
		_ecbdc, _gdaae := _gdbe.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_eb.PdfObjectInteger)
		if !_gdaae {
			_ddb.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_ecbdc == 2 {
			return _ggcf(_gdbe)
		} else if *_ecbdc == 3 {
			return _ddeca(_gdbe)
		} else {
			return nil, _dcf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else {
		_ddb.Log.Debug("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0054\u0079\u0070\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0023\u0076", _dfgg)
		return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
}
func (_dbdadd *PdfWriter) writeXRefStreams(_ebdbe int, _bffaf int64) error {
	_efgcb := _ebdbe + 1
	_dbdadd._aggfdb[_efgcb] = crossReference{Type: 1, ObjectNumber: _efgcb, Offset: _bffaf}
	_dbbee := _dd.NewBuffer(nil)
	_fdefe := _eb.MakeArray()
	for _aeagd := 0; _aeagd <= _ebdbe; {
		for ; _aeagd <= _ebdbe; _aeagd++ {
			_bccc, _adgc := _dbdadd._aggfdb[_aeagd]
			if _adgc && (!_dbdadd._caed || _dbdadd._caed && (_bccc.Type == 1 && _bccc.Offset >= _dbdadd._dfbcf || _bccc.Type == 0)) {
				break
			}
		}
		var _bfggf int
		for _bfggf = _aeagd + 1; _bfggf <= _ebdbe; _bfggf++ {
			_efgcbg, _bacbf := _dbdadd._aggfdb[_bfggf]
			if _bacbf && (!_dbdadd._caed || _dbdadd._caed && (_efgcbg.Type == 1 && _efgcbg.Offset > _dbdadd._dfbcf)) {
				continue
			}
			break
		}
		_fdefe.Append(_eb.MakeInteger(int64(_aeagd)), _eb.MakeInteger(int64(_bfggf-_aeagd)))
		for _eabc := _aeagd; _eabc < _bfggf; _eabc++ {
			_dgbcd := _dbdadd._aggfdb[_eabc]
			switch _dgbcd.Type {
			case 0:
				_ef.Write(_dbbee, _ef.BigEndian, byte(0))
				_ef.Write(_dbbee, _ef.BigEndian, uint32(0))
				_ef.Write(_dbbee, _ef.BigEndian, uint16(0xFFFF))
			case 1:
				_ef.Write(_dbbee, _ef.BigEndian, byte(1))
				_ef.Write(_dbbee, _ef.BigEndian, uint32(_dgbcd.Offset))
				_ef.Write(_dbbee, _ef.BigEndian, uint16(_dgbcd.Generation))
			case 2:
				_ef.Write(_dbbee, _ef.BigEndian, byte(2))
				_ef.Write(_dbbee, _ef.BigEndian, uint32(_dgbcd.ObjectNumber))
				_ef.Write(_dbbee, _ef.BigEndian, uint16(_dgbcd.Index))
			}
		}
		_aeagd = _bfggf + 1
	}
	_eaaf, _baedgf := _eb.MakeStream(_dbbee.Bytes(), _eb.NewFlateEncoder())
	if _baedgf != nil {
		return _baedgf
	}
	_eaaf.ObjectNumber = int64(_efgcb)
	_eaaf.PdfObjectDictionary.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0058\u0052\u0065\u0066"))
	_eaaf.PdfObjectDictionary.Set("\u0057", _eb.MakeArray(_eb.MakeInteger(1), _eb.MakeInteger(4), _eb.MakeInteger(2)))
	_eaaf.PdfObjectDictionary.Set("\u0049\u006e\u0064e\u0078", _fdefe)
	_eaaf.PdfObjectDictionary.Set("\u0053\u0069\u007a\u0065", _eb.MakeInteger(int64(_efgcb)))
	_eaaf.PdfObjectDictionary.Set("\u0049\u006e\u0066\u006f", _dbdadd._ecagd)
	_eaaf.PdfObjectDictionary.Set("\u0052\u006f\u006f\u0074", _dbdadd._ccea)
	if _dbdadd._caed && _dbdadd._becbf > 0 {
		_eaaf.PdfObjectDictionary.Set("\u0050\u0072\u0065\u0076", _eb.MakeInteger(_dbdadd._becbf))
	}
	if _dbdadd._gadcaa != nil {
		_eaaf.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _dbdadd._ebdeed)
	}
	if _dbdadd._ddefd == nil && _dbdadd._cebgc != "" && _dbdadd._bgddbe != "" {
		_dbdadd._ddefd = _eb.MakeArray(_eb.MakeHexString(_dbdadd._cebgc), _eb.MakeHexString(_dbdadd._bgddbe))
	}
	if _dbdadd._ddefd != nil {
		_ddb.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _dbdadd._ddefd)
		_eaaf.Set("\u0049\u0044", _dbdadd._ddefd)
	}
	_dbdadd.writeObject(int(_eaaf.ObjectNumber), _eaaf)
	return nil
}

// NewPdfAnnotationRedact returns a new redact annotation.
func NewPdfAnnotationRedact() *PdfAnnotationRedact {
	_fdg := NewPdfAnnotation()
	_fde := &PdfAnnotationRedact{}
	_fde.PdfAnnotation = _fdg
	_fde.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdg.SetContext(_fde)
	return _fde
}

// ToPdfObject returns the choice field dictionary within an indirect object (container).
func (_cgeea *PdfFieldChoice) ToPdfObject() _eb.PdfObject {
	_cgeea.PdfField.ToPdfObject()
	_bgcgd := _cgeea._adgda
	_cbdc := _bgcgd.PdfObject.(*_eb.PdfObjectDictionary)
	_cbdc.Set("\u0046\u0054", _eb.MakeName("\u0043\u0068"))
	if _cgeea.Opt != nil {
		_cbdc.Set("\u004f\u0070\u0074", _cgeea.Opt)
	}
	if _cgeea.TI != nil {
		_cbdc.Set("\u0054\u0049", _cgeea.TI)
	}
	if _cgeea.I != nil {
		_cbdc.Set("\u0049", _cgeea.I)
	}
	return _bgcgd
}
func (_acea *PdfWriter) flushWriter() error {
	if _acea._ceffa == nil {
		_acea._ceffa = _acea._ccfbc.Flush()
	}
	return _acea._ceffa
}
func (_fgg *PdfReader) newPdfActionImportDataFromDict(_bce *_eb.PdfObjectDictionary) (*PdfActionImportData, error) {
	_fbf, _cdc := _dba(_bce.Get("\u0046"))
	if _cdc != nil {
		return nil, _cdc
	}
	return &PdfActionImportData{F: _fbf}, nil
}

// PdfColorspaceDeviceNAttributes contains additional information about the components of colour space that
// conforming readers may use. Conforming readers need not use the alternateSpace and tintTransform parameters,
// and may instead use a custom blending algorithms, along with other information provided in the attributes
// dictionary if present.
type PdfColorspaceDeviceNAttributes struct {
	Subtype     *_eb.PdfObjectName
	Colorants   _eb.PdfObject
	Process     _eb.PdfObject
	MixingHints _eb.PdfObject
	_dbggc      *_eb.PdfIndirectObject
}

// ToInteger convert to an integer format.
func (_cfec *PdfColorDeviceCMYK) ToInteger(bits int) [4]uint32 {
	_bgaa := _gg.Pow(2, float64(bits)) - 1
	return [4]uint32{uint32(_bgaa * _cfec.C()), uint32(_bgaa * _cfec.M()), uint32(_bgaa * _cfec.Y()), uint32(_bgaa * _cfec.K())}
}

// ToPdfObject implements interface PdfModel.
func (_gfd *PdfAnnotationCircle) ToPdfObject() _eb.PdfObject {
	_gfd.PdfAnnotation.ToPdfObject()
	_fedeb := _gfd._ggf
	_fgfd := _fedeb.PdfObject.(*_eb.PdfObjectDictionary)
	_gfd.PdfAnnotationMarkup.appendToPdfDictionary(_fgfd)
	_fgfd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0043\u0069\u0072\u0063\u006c\u0065"))
	_fgfd.SetIfNotNil("\u0042\u0053", _gfd.BS)
	_fgfd.SetIfNotNil("\u0049\u0043", _gfd.IC)
	_fgfd.SetIfNotNil("\u0042\u0045", _gfd.BE)
	_fgfd.SetIfNotNil("\u0052\u0044", _gfd.RD)
	return _fedeb
}

type fontFile struct {
	_cddc  string
	_fbca  string
	_gggff _fc.SimpleEncoder
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_ffeb *PdfColorspaceDeviceGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_faca := vals[0]
	if _faca < 0.0 || _faca > 1.0 {
		_ddb.Log.Debug("\u0049\u006eco\u006d\u0070\u0061t\u0069\u0062\u0069\u006city\u003a R\u0061\u006e\u0067\u0065\u0020\u006f\u0075ts\u0069\u0064\u0065\u0020\u005b\u0030\u002c1\u005d")
	}
	if _faca < 0.0 {
		_faca = 0.0
	} else if _faca > 1.0 {
		_faca = 1.0
	}
	return NewPdfColorDeviceGray(_faca), nil
}

// NewPdfActionSubmitForm returns a new "submit form" action.
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	_bbb := NewPdfAction()
	_gf := &PdfActionSubmitForm{}
	_gf.PdfAction = _bbb
	_bbb.SetContext(_gf)
	return _gf
}

// Items returns all children outline items.
func (_aefgg *Outline) Items() []*OutlineItem { return _aefgg.Entries }

// AppendContentBytes creates a PDF stream from `cs` and appends it to the
// array of streams specified by the pages's Contents entry.
// If `wrapContents` is true, the content stream of the page is wrapped using
// a `q/Q` operator pair, so that its state does not affect the appended
// content stream.
func (_bfgbd *PdfPage) AppendContentBytes(cs []byte, wrapContents bool) error {
	_cbcge := _bfgbd.GetContentStreamObjs()
	wrapContents = wrapContents && len(_cbcge) > 0
	_fbedf := _eb.NewFlateEncoder()
	_cafdb := _eb.MakeArray()
	if wrapContents {
		_bbfdd, _bagge := _eb.MakeStream([]byte("\u0071\u000a"), _fbedf)
		if _bagge != nil {
			return _bagge
		}
		_cafdb.Append(_bbfdd)
	}
	_cafdb.Append(_cbcge...)
	if wrapContents {
		_aaecd, _feage := _eb.MakeStream([]byte("\u000a\u0051\u000a"), _fbedf)
		if _feage != nil {
			return _feage
		}
		_cafdb.Append(_aaecd)
	}
	_dbfe, _dcfgad := _eb.MakeStream(cs, _fbedf)
	if _dcfgad != nil {
		return _dcfgad
	}
	_cafdb.Append(_dbfe)
	_bfgbd.Contents = _cafdb
	return nil
}

// ColorFromPdfObjects gets the color from a series of pdf objects (3 for rgb).
func (_bfcb *PdfColorspaceDeviceRGB) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fffbe, _dcac := _eb.GetNumbersAsFloat(objects)
	if _dcac != nil {
		return nil, _dcac
	}
	return _bfcb.ColorFromFloats(_fffbe)
}

// NewEmbeddedFile constructs a new EmbeddedFile object from supplied file.
// The file type of the file would be detected automatically.
func NewEmbeddedFile(path string) (*EmbeddedFile, error) {
	_feff, _abgc := _ccb.ReadFile(path)
	if _abgc != nil {
		return nil, _abgc
	}
	_ggfeb := _egg.Detect(_feff)
	_dgab := _c.Sum(_feff)
	_eaaa := &EmbeddedFile{Name: _dc.Base(path), Content: _feff, FileType: _ggfeb.String(), Hash: _egf.EncodeToString(_dgab[:])}
	return _eaaa, nil
}

// PdfOutputIntentType is the subtype of the given PdfOutputIntent.
type PdfOutputIntentType int

// AlphaMapFunc represents a alpha mapping function: byte -> byte. Can be used for
// thresholding the alpha channel, i.e. setting all alpha values below threshold to transparent.
type AlphaMapFunc func(_acefd byte) byte

func (_abace *PdfWriter) writeString(_dbdabb string) {
	if _abace._ceffa != nil {
		return
	}
	_caabf, _ecdgef := _abace._ccfbc.WriteString(_dbdabb)
	_abace._dfabe += int64(_caabf)
	_abace._ceffa = _ecdgef
}

// NewPdfColorspaceCalGray returns a new CalGray colorspace object.
func NewPdfColorspaceCalGray() *PdfColorspaceCalGray {
	_fafg := &PdfColorspaceCalGray{}
	_fafg.BlackPoint = []float64{0.0, 0.0, 0.0}
	_fafg.Gamma = 1
	return _fafg
}
func _egggc(_gacacb _eb.PdfObject) []*_eb.PdfObjectStream {
	if _gacacb == nil {
		return nil
	}
	_bgdbf, _afedg := _eb.GetArray(_gacacb)
	if !_afedg || _bgdbf.Len() == 0 {
		return nil
	}
	_cbadf := make([]*_eb.PdfObjectStream, 0, _bgdbf.Len())
	for _, _eedc := range _bgdbf.Elements() {
		if _cegeg, _bfccd := _eb.GetStream(_eedc); _bfccd {
			_cbadf = append(_cbadf, _cegeg)
		}
	}
	return _cbadf
}

// NewPdfActionNamed returns a new "named" action.
func NewPdfActionNamed() *PdfActionNamed {
	_gag := NewPdfAction()
	_ce := &PdfActionNamed{}
	_ce.PdfAction = _gag
	_gag.SetContext(_ce)
	return _ce
}

// PdfColor interface represents a generic color in PDF.
type PdfColor interface{}

// ToPdfObject implements interface PdfModel.
func (_ebccd *PdfAnnotationInk) ToPdfObject() _eb.PdfObject {
	_ebccd.PdfAnnotation.ToPdfObject()
	_geag := _ebccd._ggf
	_cecf := _geag.PdfObject.(*_eb.PdfObjectDictionary)
	_ebccd.PdfAnnotationMarkup.appendToPdfDictionary(_cecf)
	_cecf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0049\u006e\u006b"))
	_cecf.SetIfNotNil("\u0049n\u006b\u004c\u0069\u0073\u0074", _ebccd.InkList)
	_cecf.SetIfNotNil("\u0042\u0053", _ebccd.BS)
	return _geag
}

// NewPdfOutlineItem returns an initialized PdfOutlineItem.
func NewPdfOutlineItem() *PdfOutlineItem {
	_cfecg := &PdfOutlineItem{_efbfb: _eb.MakeIndirectObject(_eb.MakeDict())}
	_cfecg._eeedb = _cfecg
	return _cfecg
}

// FlattenFields flattens the form fields and annotations for the PDF loaded in `pdf` and makes
// non-editable.
// Looks up all widget annotations corresponding to form fields and flattens them by drawing the content
// through the content stream rather than annotations.
// References to flattened annotations will be removed from Page Annots array. For fields the AcroForm entry
// will be emptied.
// When `allannots` is true, all annotations will be flattened. Keep false if want to keep non-form related
// annotations intact.
// When `appgen` is not nil, it will be used to generate appearance streams for the field annotations.
func (_ddbf *PdfReader) FlattenFields(allannots bool, appgen FieldAppearanceGenerator) error {
	return _ddbf.flattenFieldsWithOpts(allannots, appgen, nil)
}
func _ffeeb(_faaeg *_eb.PdfObjectDictionary) (*PdfShadingType6, error) {
	_ededed := PdfShadingType6{}
	_cdacd := _faaeg.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _cdacd == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_gdfec, _bbgce := _cdacd.(*_eb.PdfObjectInteger)
	if !_bbgce {
		_ddb.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cdacd)
		return nil, _eb.ErrTypeError
	}
	_ededed.BitsPerCoordinate = _gdfec
	_cdacd = _faaeg.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _cdacd == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_gdfec, _bbgce = _cdacd.(*_eb.PdfObjectInteger)
	if !_bbgce {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _cdacd)
		return nil, _eb.ErrTypeError
	}
	_ededed.BitsPerComponent = _gdfec
	_cdacd = _faaeg.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _cdacd == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gdfec, _bbgce = _cdacd.(*_eb.PdfObjectInteger)
	if !_bbgce {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _cdacd)
		return nil, _eb.ErrTypeError
	}
	_ededed.BitsPerComponent = _gdfec
	_cdacd = _faaeg.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _cdacd == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dcebc, _bbgce := _cdacd.(*_eb.PdfObjectArray)
	if !_bbgce {
		_ddb.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cdacd)
		return nil, _eb.ErrTypeError
	}
	_ededed.Decode = _dcebc
	if _aaceg := _faaeg.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _aaceg != nil {
		_ededed.Function = []PdfFunction{}
		if _fgaec, _efeca := _aaceg.(*_eb.PdfObjectArray); _efeca {
			for _, _adce := range _fgaec.Elements() {
				_abcgf, _eebaf := _cccfa(_adce)
				if _eebaf != nil {
					_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _eebaf)
					return nil, _eebaf
				}
				_ededed.Function = append(_ededed.Function, _abcgf)
			}
		} else {
			_fgccd, _defa := _cccfa(_aaceg)
			if _defa != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _defa)
				return nil, _defa
			}
			_ededed.Function = append(_ededed.Function, _fgccd)
		}
	}
	return &_ededed, nil
}
func _eecd(_ecfb *_eb.PdfObjectDictionary) *VRI {
	_aafe, _ := _eb.GetString(_ecfb.Get("\u0054\u0055"))
	_gfedg, _ := _eb.GetString(_ecfb.Get("\u0054\u0053"))
	return &VRI{Cert: _egggc(_ecfb.Get("\u0043\u0065\u0072\u0074")), OCSP: _egggc(_ecfb.Get("\u004f\u0043\u0053\u0050")), CRL: _egggc(_ecfb.Get("\u0043\u0052\u004c")), TU: _aafe, TS: _gfedg}
}

// ToPdfObject implements interface PdfModel.
func (_cfgd *PdfActionThread) ToPdfObject() _eb.PdfObject {
	_cfgd.PdfAction.ToPdfObject()
	_cddg := _cfgd._dee
	_bba := _cddg.PdfObject.(*_eb.PdfObjectDictionary)
	_bba.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeThread)))
	if _cfgd.F != nil {
		_bba.Set("\u0046", _cfgd.F.ToPdfObject())
	}
	_bba.SetIfNotNil("\u0044", _cfgd.D)
	_bba.SetIfNotNil("\u0042", _cfgd.B)
	return _cddg
}

// AddImageResource adds an image to the XObject resources.
func (_dbcf *PdfPage) AddImageResource(name _eb.PdfObjectName, ximg *XObjectImage) error {
	var _fdab *_eb.PdfObjectDictionary
	if _dbcf.Resources.XObject == nil {
		_fdab = _eb.MakeDict()
		_dbcf.Resources.XObject = _fdab
	} else {
		var _cabaf bool
		_fdab, _cabaf = (_dbcf.Resources.XObject).(*_eb.PdfObjectDictionary)
		if !_cabaf {
			return _dcf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0078\u0072\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020\u0074\u0079p\u0065")
		}
	}
	_fdab.Set(name, ximg.ToPdfObject())
	return nil
}
func (_efd *PdfReader) newPdfAnnotationFreeTextFromDict(_ecgg *_eb.PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	_baa := PdfAnnotationFreeText{}
	_gef, _deea := _efd.newPdfAnnotationMarkupFromDict(_ecgg)
	if _deea != nil {
		return nil, _deea
	}
	_baa.PdfAnnotationMarkup = _gef
	_baa.DA = _ecgg.Get("\u0044\u0041")
	_baa.Q = _ecgg.Get("\u0051")
	_baa.RC = _ecgg.Get("\u0052\u0043")
	_baa.DS = _ecgg.Get("\u0044\u0053")
	_baa.CL = _ecgg.Get("\u0043\u004c")
	_baa.IT = _ecgg.Get("\u0049\u0054")
	_baa.BE = _ecgg.Get("\u0042\u0045")
	_baa.RD = _ecgg.Get("\u0052\u0044")
	_baa.BS = _ecgg.Get("\u0042\u0053")
	_baa.LE = _ecgg.Get("\u004c\u0045")
	return &_baa, nil
}
func (_gagef *Image) getSuitableEncoder() (_eb.StreamEncoder, error) {
	var (
		_cfbgf, _bgbff = int(_gagef.Width), int(_gagef.Height)
		_geea          = make(map[string]bool)
		_fafb          = true
		_deddf         = false
		_degc          = func() *_eb.DCTEncoder { return _eb.NewDCTEncoder() }
		_abcb          = func() *_eb.DCTEncoder { _fbbac := _eb.NewDCTEncoder(); _fbbac.BitsPerComponent = 16; return _fbbac }
	)
	for _bbgde := 0; _bbgde < _bgbff; _bbgde++ {
		for _bgddb := 0; _bgddb < _cfbgf; _bgddb++ {
			_ebgaa, _afdcd := _gagef.ColorAt(_bgddb, _bbgde)
			if _afdcd != nil {
				return nil, _afdcd
			}
			_aaaga, _abged, _cdag, _dcefe := _ebgaa.RGBA()
			if _fafb && (_aaaga != _abged || _aaaga != _cdag || _abged != _cdag) {
				_fafb = false
			}
			if !_deddf {
				switch _ebgaa.(type) {
				case _b.NRGBA:
					_deddf = _dcefe > 0
				}
			}
			_geea[_e.Sprintf("\u0025\u0064\u002c\u0025\u0064\u002c\u0025\u0064", _aaaga, _abged, _cdag)] = true
			if len(_geea) > 2 && _deddf {
				return _abcb(), nil
			}
		}
	}
	if _deddf || len(_gagef._bdcab) > 0 {
		return _eb.NewFlateEncoder(), nil
	}
	if len(_geea) <= 2 {
		_cbeca := _gagef.ConvertToBinary()
		if _cbeca != nil {
			return nil, _cbeca
		}
		return _eb.NewJBIG2Encoder(), nil
	}
	if _fafb {
		return _degc(), nil
	}
	if _gagef.ColorComponents == 1 {
		if _gagef.BitsPerComponent == 1 {
			return _eb.NewJBIG2Encoder(), nil
		} else if _gagef.BitsPerComponent == 8 {
			_bedba := _eb.NewDCTEncoder()
			_bedba.ColorComponents = 1
			return _bedba, nil
		}
	} else if _gagef.ColorComponents == 3 {
		if _gagef.BitsPerComponent == 8 {
			return _degc(), nil
		} else if _gagef.BitsPerComponent == 16 {
			return _abcb(), nil
		}
	} else if _gagef.ColorComponents == 4 {
		_afbeg := _abcb()
		_afbeg.ColorComponents = 4
		return _afbeg, nil
	}
	return _abcb(), nil
}

// NewPdfAnnotationWatermark returns a new watermark annotation.
func NewPdfAnnotationWatermark() *PdfAnnotationWatermark {
	_cdgd := NewPdfAnnotation()
	_fddc := &PdfAnnotationWatermark{}
	_fddc.PdfAnnotation = _cdgd
	_cdgd.SetContext(_fddc)
	return _fddc
}
func _eegg(_bedad _eb.PdfObject) (*PdfColorspaceSpecialPattern, error) {
	_ddb.Log.Trace("\u004e\u0065\u0077\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u0043\u0053\u0020\u0066\u0072\u006fm\u0020\u006f\u0062\u006a\u003a\u0020\u0025\u0073\u0020\u0025\u0054", _bedad.String(), _bedad)
	_bfabf := NewPdfColorspaceSpecialPattern()
	if _febbg, _dcbe := _bedad.(*_eb.PdfIndirectObject); _dcbe {
		_bfabf._fgff = _febbg
	}
	_bedad = _eb.TraceToDirectObject(_bedad)
	if _ccabe, _afbf := _bedad.(*_eb.PdfObjectName); _afbf {
		if *_ccabe != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			return nil, _e.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
		return _bfabf, nil
	}
	_egda, _bbg := _bedad.(*_eb.PdfObjectArray)
	if !_bbg {
		_ddb.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061t\u0074\u0065\u0072\u006e\u0020\u0043\u0053 \u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0023\u0076", _bedad)
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0050\u0061\u0074\u0074e\u0072n\u0020C\u0053\u0020\u006f\u0062\u006a\u0065\u0063t")
	}
	if _egda.Len() != 1 && _egda.Len() != 2 {
		_ddb.Log.Error("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079\u003a %\u0023\u0076", _egda)
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065r\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bedad = _egda.Get(0)
	if _dacg, _beec := _bedad.(*_eb.PdfObjectName); _beec {
		if *_dacg != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			_ddb.Log.Error("\u0049\u006e\u0076al\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020C\u0053 \u0061r\u0072\u0061\u0079\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0023\u0076", _dacg)
			return nil, _e.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
	}
	if _egda.Len() > 1 {
		_bedad = _egda.Get(1)
		_bedad = _eb.TraceToDirectObject(_bedad)
		_fgfdf, _beaa := NewPdfColorspaceFromPdfObject(_bedad)
		if _beaa != nil {
			return nil, _beaa
		}
		_bfabf.UnderlyingCS = _fgfdf
	}
	_ddb.Log.Trace("R\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0077i\u0074\u0068\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079in\u0067\u0020\u0063s\u003a \u0025\u0054", _bfabf.UnderlyingCS)
	return _bfabf, nil
}

// PdfSignature represents a PDF signature dictionary and is used for signing via form signature fields.
// (Section 12.8, Table 252 - Entries in a signature dictionary p. 475 in PDF32000_2008).
type PdfSignature struct {
	Handler SignatureHandler
	_cddce  *_eb.PdfIndirectObject

	// Type: Sig/DocTimeStamp
	Type         *_eb.PdfObjectName
	Filter       *_eb.PdfObjectName
	SubFilter    *_eb.PdfObjectName
	Contents     *_eb.PdfObjectString
	Cert         _eb.PdfObject
	ByteRange    *_eb.PdfObjectArray
	Reference    *_eb.PdfObjectArray
	Changes      *_eb.PdfObjectArray
	Name         *_eb.PdfObjectString
	M            *_eb.PdfObjectString
	Location     *_eb.PdfObjectString
	Reason       *_eb.PdfObjectString
	ContactInfo  *_eb.PdfObjectString
	R            *_eb.PdfObjectInteger
	V            *_eb.PdfObjectInteger
	PropBuild    *_eb.PdfObjectDictionary
	PropAuthTime *_eb.PdfObjectInteger
	PropAuthType *_eb.PdfObjectName
}

// NewPdfFontFromTTF loads a TTF font and returns a PdfFont type that can be
// used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTF.
func NewPdfFontFromTTF(r _bagf.ReadSeeker) (*PdfFont, error) {
	const _eefg = _fc.CharCode(32)
	const _ecgga = _fc.CharCode(255)
	_gfeff, _ccbf := _bagf.ReadAll(r)
	if _ccbf != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _ccbf)
		return nil, _ccbf
	}
	_aceda, _ccbf := _fg.TtfParse(_dd.NewReader(_gfeff))
	if _ccbf != nil {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0054\u0054F\u0020\u0066\u006fn\u0074:\u0020\u0025\u0076", _ccbf)
		return nil, _ccbf
	}
	_gafa := &pdfFontSimple{_cegda: make(map[_fc.CharCode]float64), fontCommon: fontCommon{_fgdee: "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065"}}
	_gafa._eccbb = _fc.NewWinAnsiEncoder()
	_gafa._agcc = _aceda.PostScriptName
	_gafa.FirstChar = _eb.MakeInteger(int64(_eefg))
	_gafa.LastChar = _eb.MakeInteger(int64(_ecgga))
	_afbd := 1000.0 / float64(_aceda.UnitsPerEm)
	if len(_aceda.Widths) <= 0 {
		return nil, _dcf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_fead := _afbd * float64(_aceda.Widths[0])
	_ccde := make([]float64, 0, _ecgga-_eefg+1)
	for _cgdge := _eefg; _cgdge <= _ecgga; _cgdge++ {
		_gcgga, _efaac := _gafa.Encoder().CharcodeToRune(_cgdge)
		if !_efaac {
			_ddb.Log.Debug("\u0052u\u006e\u0065\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u0028c\u006f\u0064\u0065\u003a\u0020\u0025\u0064)", _cgdge)
			_ccde = append(_ccde, _fead)
			continue
		}
		_gade, _abgdd := _aceda.Chars[_gcgga]
		if !_abgdd {
			_ddb.Log.Debug("R\u0075\u006e\u0065\u0020no\u0074 \u0069\u006e\u0020\u0054\u0054F\u0020\u0043\u0068\u0061\u0072\u0073")
			_ccde = append(_ccde, _fead)
			continue
		}
		_fbbbe := _afbd * float64(_aceda.Widths[_gade])
		_ccde = append(_ccde, _fbbbe)
	}
	_gafa.Widths = _eb.MakeIndirectObject(_eb.MakeArrayFromFloats(_ccde))
	if len(_ccde) < int(_ecgga-_eefg+1) {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u006f\u0066\u0020\u0077\u0069\u0064\u0074\u0068s,\u0020\u0025\u0064 \u003c \u0025\u0064", len(_ccde), 255-32+1)
		return nil, _eb.ErrRangeError
	}
	for _dccdb := _eefg; _dccdb <= _ecgga; _dccdb++ {
		_gafa._cegda[_dccdb] = _ccde[_dccdb-_eefg]
	}
	_gafa.Encoding = _eb.MakeName("\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_bcag := &PdfFontDescriptor{}
	_bcag.FontName = _eb.MakeName(_aceda.PostScriptName)
	_bcag.Ascent = _eb.MakeFloat(_afbd * float64(_aceda.TypoAscender))
	_bcag.Descent = _eb.MakeFloat(_afbd * float64(_aceda.TypoDescender))
	_bcag.CapHeight = _eb.MakeFloat(_afbd * float64(_aceda.CapHeight))
	_bcag.FontBBox = _eb.MakeArrayFromFloats([]float64{_afbd * float64(_aceda.Xmin), _afbd * float64(_aceda.Ymin), _afbd * float64(_aceda.Xmax), _afbd * float64(_aceda.Ymax)})
	_bcag.ItalicAngle = _eb.MakeFloat(_aceda.ItalicAngle)
	_bcag.MissingWidth = _eb.MakeFloat(_afbd * float64(_aceda.Widths[0]))
	_eebfb, _ccbf := _eb.MakeStream(_gfeff, _eb.NewFlateEncoder())
	if _ccbf != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _ccbf)
		return nil, _ccbf
	}
	_eebfb.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _eb.MakeInteger(int64(len(_gfeff))))
	_bcag.FontFile2 = _eebfb
	if _aceda.Bold {
		_bcag.StemV = _eb.MakeInteger(120)
	} else {
		_bcag.StemV = _eb.MakeInteger(70)
	}
	_gacac := _ebaed
	if _aceda.IsFixedPitch {
		_gacac |= _eaggc
	}
	if _aceda.ItalicAngle != 0 {
		_gacac |= _ffbe
	}
	_bcag.Flags = _eb.MakeInteger(int64(_gacac))
	_gafa._bged = _bcag
	_dgdg := &PdfFont{_fdaa: _gafa}
	return _dgdg, nil
}
func _affe(_bdfca *_eb.PdfObjectDictionary, _gaeag *fontCommon) (*pdfFontType0, error) {
	_daab, _gbae := _eb.GetArray(_bdfca.Get("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073"))
	if !_gbae {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006cid\u0020\u0044\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073\u0020\u002d\u0020\u006e\u006f\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079 \u0025\u0073", _gaeag)
		return nil, _eb.ErrRangeError
	}
	if _daab.Len() != 1 {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0041\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020(%\u0064\u0029", _daab.Len())
		return nil, _eb.ErrRangeError
	}
	_gfdb, _cfcbe := _fdce(_daab.Get(0), false)
	if _cfcbe != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065d \u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u0020\u0066\u006f\u006et\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076 \u0025\u0073", _cfcbe, _gaeag)
		return nil, _cfcbe
	}
	_caebe := _bdfca.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_daaaf := ""
	_dddcd := _fdbc(_gaeag)
	_dddcd.DescendantFont = _gfdb
	switch _abbf := _caebe.(type) {
	case *_eb.PdfObjectName:
		_daaaf, _gbae = _eb.GetNameVal(_caebe)
		if _gbae {
			if _daaaf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _daaaf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" {
				_dddcd._edcff = _fc.NewIdentityTextEncoder(_daaaf)
			} else if _ff.IsPredefinedCMap(_daaaf) {
				_dddcd._ebeb, _cfcbe = _ff.LoadPredefinedCMap(_daaaf)
				if _cfcbe != nil {
					_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _daaaf, _cfcbe)
				}
			} else {
				_ddb.Log.Debug("\u0055\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0063\u006da\u0070\u0020\u0025\u0071", _daaaf)
			}
		}
	case *_eb.PdfObjectStream:
		if _dddcd._geee == nil {
			_dcgec, _efbbg := _ff.NewCIDSystemInfo(_abbf.PdfObjectDictionary.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
			if _efbbg != nil {
				_ddb.Log.Debug("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0043\u0049D\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0020\u0025\u0076", _efbbg)
			}
			_bdga := _e.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _dcgec.Registry, _dcgec.Ordering)
			if _ff.IsPredefinedCMap(_bdga) {
				_dddcd._ebeb, _efbbg = _ff.LoadPredefinedCMap(_bdga)
				if _efbbg != nil {
					_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _bdga, _efbbg)
				}
			} else {
				_bdga = _abbf.PdfObjectDictionary.Get("\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065").String()
				_aeba, _gecb := _eb.DecodeStream(_abbf)
				if _gecb != nil {
					_ddb.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020s\u0074\u0072\u0065a\u006d:\u0020\u0025\u0076", _gecb)
					return _dddcd, _gecb
				}
				if _befc := _bdga == "\u004f\u006ee\u0042\u0079\u0074e\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u0048" || _bdga == "\u004f\u006ee\u0042\u0079\u0074e\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u0056"; _befc {
					_dddcd._ebeb, _gecb = _ff.LoadCmapFromData(_aeba, _befc)
					if _gecb != nil {
						_ddb.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u006c\u006f\u0061\u0064\u0020\u0043\u004d\u0061\u0070 \u0066\u0072\u006f\u006d\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076\u0020\u002d\u002d\u0020\u0025\u0076", _gecb, string(_aeba))
						return _dddcd, _gecb
					}
					_acaeea := make(map[_ff.CharCode]rune, 256)
					for _bdcd := 0x00; _bdcd <= 0xFF; _bdcd++ {
						_acaeea[_ff.CharCode(_bdcd)] = rune(_bdcd)
					}
					_dddcd._bgbg = _ff.NewToUnicodeCMap(_acaeea)
				}
			}
		}
	}
	if _gdafd := _gfdb.baseFields()._bgbg; _gdafd != nil {
		if _fdegc := _gdafd.Name(); _fdegc == "\u0041d\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002d\u0055\u0043\u0053\u0032" || _fdegc == "\u0041\u0064\u006f\u0062\u0065\u002d\u0047\u0042\u0031-\u0055\u0043\u0053\u0032" || _fdegc == "\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031-\u0055\u0043\u0053\u0032" || _fdegc == "\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006f\u0072\u0065\u0061\u0031-\u0055\u0043\u0053\u0032" {
			_dddcd._edcff = _fc.NewCMapEncoder(_daaaf, _dddcd._ebeb, _gdafd)
		}
	}
	return _dddcd, nil
}
func (_agfbb *PdfPage) getParentResources() (*PdfPageResources, error) {
	_bbaae := _agfbb.Parent
	for _bbaae != nil {
		_ecddb, _aagdf := _eb.GetDict(_bbaae)
		if !_aagdf {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020n\u006f\u0064\u0065")
			return nil, _dcf.New("i\u006e\u0076\u0061\u006cid\u0020p\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		if _caeae := _ecddb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _caeae != nil {
			_acfff, _gcgfb := _eb.GetDict(_caeae)
			if !_gcgfb {
				return nil, _dcf.New("i\u006e\u0076\u0061\u006cid\u0020r\u0065\u0073\u006f\u0075\u0072c\u0065\u0020\u0064\u0069\u0063\u0074")
			}
			_bgaca, _eeff := NewPdfPageResourcesFromDict(_acfff)
			if _eeff != nil {
				return nil, _eeff
			}
			return _bgaca, nil
		}
		_bbaae = _ecddb.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, nil
}

// AddKDict adds a K dictionary object to the structure tree root.
func (_eadgg *StructTreeRoot) AddKDict(k *KDict) { _eadgg.K = append(_eadgg.K, k) }

// PdfAnnotationStrikeOut represents StrikeOut annotations.
// (Section 12.5.6.10).
type PdfAnnotationStrikeOut struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _eb.PdfObject
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_aeab *PdfColorspaceICCBased) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _aeab.Alternate == nil {
		if _aeab.N == 1 {
			_badda := NewPdfColorspaceDeviceGray()
			return _badda.ColorFromFloats(vals)
		} else if _aeab.N == 3 {
			_fcfbg := NewPdfColorspaceDeviceRGB()
			return _fcfbg.ColorFromFloats(vals)
		} else if _aeab.N == 4 {
			_ceff := NewPdfColorspaceDeviceCMYK()
			return _ceff.ColorFromFloats(vals)
		} else {
			return nil, _dcf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _aeab.Alternate.ColorFromFloats(vals)
}

// PdfOutputIntent provides a means for matching the color characteristics of a PDF document with
// those of a target output device.
// Multiple PdfOutputIntents allows the production process to be customized to the expected workflow and the specific
// tools available.
type PdfOutputIntent struct {

	// Type is an optional PDF object that this dictionary describes.
	// If present, must be OutputIntent for an output intent dictionary.
	Type string

	// S defines the OutputIntent subtype which should match the standard used in given document i.e:
	// for PDF/X use PdfOutputIntentTypeX.
	S PdfOutputIntentType

	// OutputCondition is an optional field that is identifying the intended output device or production condition in
	// human-readable form. This is preferred method of defining such a string for presentation to the user.
	OutputCondition string

	// OutputConditionIdentifier is a required field identifying the intended output device or production condition in
	// human or machine-readable form. If human-readable, this string may be used
	// in lieu of an OutputCondition for presentation to the user.
	// A typical value for this entry would be the name of a production condition  maintained
	// in  an  industry-standard registry such  as the ICC Characterization Data Registry
	// If the intended production condition is not a recognized standard, the value Custom is recommended for this entry.
	// the DestOutputProfile entry defines the ICC profile, and the Info entry is used for further
	// human-readable identification.
	OutputConditionIdentifier string

	// RegistryName is an optional string field (conventionally URI) identifying the registry in which the condition
	// designated by OutputConditionIdentifier is defined.
	RegistryName string

	// Info is a required field if OutputConditionIdentifier does not specify a standard production condition.
	// A human-readable text string containing additional information  or comments about intended
	// target device or production condition.
	Info string

	// DestOutputProfile is required if OutputConditionIdentifier does not specify a standard production condition.
	// It is an ICC profile stream defining the transformation from the PDF document's source colors to output device colorants.
	DestOutputProfile []byte

	// ColorComponents is the number of color components supported by given output profile.
	ColorComponents int
	_ebbd           *_eb.PdfObjectDictionary
}

// SetPickTrayByPDFSize sets the value of the pickTrayByPDFSize flag.
func (_egdag *ViewerPreferences) SetPickTrayByPDFSize(pickTrayByPDFSize bool) {
	_egdag._gggea = &pickTrayByPDFSize
}

// NewPdfReaderWithOpts creates a new PdfReader for an input io.ReadSeeker interface
// with a ReaderOpts.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderWithOpts(rs _bagf.ReadSeeker, opts *ReaderOpts) (*PdfReader, error) {
	const _dagb = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0057\u0069\u0074\u0068O\u0070\u0074\u0073"
	return _dcgfe(rs, opts, true, _dagb)
}

// HasXObjectByName checks if has XObject resource by name.
func (_babfga *PdfPage) HasXObjectByName(name _eb.PdfObjectName) bool {
	_gcbe, _cbacb := _babfga.Resources.XObject.(*_eb.PdfObjectDictionary)
	if !_cbacb {
		return false
	}
	if _cccgf := _gcbe.Get(name); _cccgf != nil {
		return true
	}
	return false
}

// NewPdfActionImportData returns a new "import data" action.
func NewPdfActionImportData() *PdfActionImportData {
	_ccf := NewPdfAction()
	_efa := &PdfActionImportData{}
	_efa.PdfAction = _ccf
	_ccf.SetContext(_efa)
	return _efa
}

// Optimizer is the interface that performs optimization of PDF object structure for output writing.
//
// Optimize receives a slice of input `objects`, performs optimization, including removing, replacing objects and
// output the optimized slice of objects.
type Optimizer interface {
	Optimize(_dafca []_eb.PdfObject) ([]_eb.PdfObject, error)
}

// SetOpenAction sets the OpenAction in the PDF catalog.
// The value shall be either an array defining a destination (12.3.2 "Destinations" PDF32000_2008),
// or an action dictionary representing an action (12.6 "Actions" PDF32000_2008).
func (_facab *PdfWriter) SetOpenAction(dest _eb.PdfObject) error {
	if dest == nil || _eb.IsNullObject(dest) {
		return nil
	}
	_facab._dbffa.Set("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e", dest)
	return _facab.addObjects(dest)
}

// SetNonFullScreenPageMode sets the value of the nonFullScreenPageMode.
func (_ceebf *ViewerPreferences) SetNonFullScreenPageMode(nonFullScreenPageMode NonFullScreenPageMode) {
	_ceebf._cdcff = nonFullScreenPageMode
}

// XObjectImage (Table 89 in 8.9.5.1).
// Implements PdfModel interface.
type XObjectImage struct {

	// ColorSpace       PdfObject
	Width            *int64
	Height           *int64
	ColorSpace       PdfColorspace
	BitsPerComponent *int64
	Filter           _eb.StreamEncoder
	Intent           _eb.PdfObject
	ImageMask        _eb.PdfObject
	Mask             _eb.PdfObject
	Matte            _eb.PdfObject
	Decode           _eb.PdfObject
	Interpolate      _eb.PdfObject
	Alternatives     _eb.PdfObject
	SMask            _eb.PdfObject
	SMaskInData      _eb.PdfObject
	Name             _eb.PdfObject
	StructParent     _eb.PdfObject
	ID               _eb.PdfObject
	OPI              _eb.PdfObject
	Metadata         _eb.PdfObject
	OC               _eb.PdfObject
	Stream           []byte
	_gceffg          *_eb.PdfObjectStream
	_gfegb           bool
}

// StringToCharcodeBytes maps the provided string runes to charcode bytes and
// it returns the resulting slice of bytes, along with the number of runes
// which could not be converted. If the number of misses is 0, all string runes
// were successfully converted.
func (_cgaaa *PdfFont) StringToCharcodeBytes(str string) ([]byte, int) {
	return _cgaaa.RunesToCharcodeBytes([]rune(str))
}
func _cdcbe(_gccge *_eb.PdfObjectDictionary) (*PdfShadingType2, error) {
	_bbgfa := PdfShadingType2{}
	_dccgc := _gccge.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _dccgc == nil {
		_ddb.Log.Debug("R\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u003a\u0020\u0020\u0043\u006f\u006f\u0072d\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_agade, _agbgac := _dccgc.(*_eb.PdfObjectArray)
	if !_agbgac {
		_ddb.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _dccgc)
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _agade.Len() != 4 {
		_ddb.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0034\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _agade.Len())
		return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	}
	_bbgfa.Coords = _agade
	if _eagfe := _gccge.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _eagfe != nil {
		_eagfe = _eb.TraceToDirectObject(_eagfe)
		_fgae, _egbeb := _eagfe.(*_eb.PdfObjectArray)
		if !_egbeb {
			_ddb.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _eagfe)
			return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bbgfa.Domain = _fgae
	}
	_dccgc = _gccge.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _dccgc == nil {
		_ddb.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_bbgfa.Function = []PdfFunction{}
	if _fcff, _ffdfg := _dccgc.(*_eb.PdfObjectArray); _ffdfg {
		for _, _feffg := range _fcff.Elements() {
			_geccb, _cdbaae := _cccfa(_feffg)
			if _cdbaae != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cdbaae)
				return nil, _cdbaae
			}
			_bbgfa.Function = append(_bbgfa.Function, _geccb)
		}
	} else {
		_afgd, _fbgeeg := _cccfa(_dccgc)
		if _fbgeeg != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fbgeeg)
			return nil, _fbgeeg
		}
		_bbgfa.Function = append(_bbgfa.Function, _afgd)
	}
	if _fbfac := _gccge.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _fbfac != nil {
		_fbfac = _eb.TraceToDirectObject(_fbfac)
		_gccga, _gdgdbf := _fbfac.(*_eb.PdfObjectArray)
		if !_gdgdbf {
			_ddb.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _fbfac)
			return nil, _eb.ErrTypeError
		}
		if _gccga.Len() != 2 {
			_ddb.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _gccga.Len())
			return nil, ErrInvalidAttribute
		}
		_bbgfa.Extend = _gccga
	}
	return &_bbgfa, nil
}
func (_deag *PdfReader) newPdfAnnotationWidgetFromDict(_ddeg *_eb.PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	_gdaf := PdfAnnotationWidget{}
	_gdaf.H = _ddeg.Get("\u0048")
	_gdaf.MK = _ddeg.Get("\u004d\u004b")
	_gdaf.A = _ddeg.Get("\u0041")
	_gdaf.AA = _ddeg.Get("\u0041\u0041")
	_gdaf.BS = _ddeg.Get("\u0042\u0053")
	_gdaf.Parent = _ddeg.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	return &_gdaf, nil
}

// AddOCSPs adds OCSPs to DSS.
func (_feae *DSS) AddOCSPs(ocsps [][]byte) ([]*_eb.PdfObjectStream, error) {
	return _feae.add(&_feae.OCSPs, _feae._gccb, ocsps)
}

// ToPdfObject implements interface PdfModel.
func (_geg *PdfAnnotationUnderline) ToPdfObject() _eb.PdfObject {
	_geg.PdfAnnotation.ToPdfObject()
	_dbe := _geg._ggf
	_bdgd := _dbe.PdfObject.(*_eb.PdfObjectDictionary)
	_geg.PdfAnnotationMarkup.appendToPdfDictionary(_bdgd)
	_bdgd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee"))
	_bdgd.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _geg.QuadPoints)
	return _dbe
}

// GetNumComponents returns the number of color components (1 for grayscale).
func (_gdcac *PdfColorDeviceGray) GetNumComponents() int { return 1 }

// NewPdfInfoFromObject creates a new PdfInfo from the input core.PdfObject.
func NewPdfInfoFromObject(obj _eb.PdfObject) (*PdfInfo, error) {
	var _ebgb PdfInfo
	_ggcb, _baaf := obj.(*_eb.PdfObjectDictionary)
	if !_baaf {
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", obj)
	}
	for _, _fgdfd := range _ggcb.Keys() {
		switch _fgdfd {
		case "\u0054\u0069\u0074l\u0065":
			_ebgb.Title, _ = _eb.GetString(_ggcb.Get("\u0054\u0069\u0074l\u0065"))
		case "\u0041\u0075\u0074\u0068\u006f\u0072":
			_ebgb.Author, _ = _eb.GetString(_ggcb.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
		case "\u0053u\u0062\u006a\u0065\u0063\u0074":
			_ebgb.Subject, _ = _eb.GetString(_ggcb.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
		case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
			_ebgb.Keywords, _ = _eb.GetString(_ggcb.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
		case "\u0043r\u0065\u0061\u0074\u006f\u0072":
			_ebgb.Creator, _ = _eb.GetString(_ggcb.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
		case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
			_ebgb.Producer, _ = _eb.GetString(_ggcb.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
		case "\u0054r\u0061\u0070\u0070\u0065\u0064":
			_ebgb.Trapped, _ = _eb.GetName(_ggcb.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
		case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
			if _ddba, _eeef := _eb.GetString(_ggcb.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _eeef && _ddba.String() != "" {
				_bedg, _egaa := NewPdfDate(_ddba.String())
				if _egaa != nil {
					return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _egaa)
				}
				_ebgb.CreationDate = &_bedg
			}
		case "\u004do\u0064\u0044\u0061\u0074\u0065":
			if _gbda, _cdcb := _eb.GetString(_ggcb.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _cdcb && _gbda.String() != "" {
				_facbf, _egee := NewPdfDate(_gbda.String())
				if _egee != nil {
					return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _egee)
				}
				_ebgb.ModifiedDate = &_facbf
			}
		default:
			_gfbe, _ := _eb.GetString(_ggcb.Get(_fgdfd))
			if _ebgb._cbfb == nil {
				_ebgb._cbfb = _eb.MakeDict()
			}
			_ebgb._cbfb.Set(_fgdfd, _gfbe)
		}
	}
	return &_ebgb, nil
}

// GetBorderWidth returns the border style's width.
func (_eead *PdfBorderStyle) GetBorderWidth() float64 {
	if _eead.W == nil {
		return 1
	}
	return *_eead.W
}
func (_adgb *PdfColorspaceCalGray) String() string { return "\u0043a\u006c\u0047\u0072\u0061\u0079" }

// HasColorspaceByName checks if the colorspace with the specified name exists in the page resources.
func (_afcfad *PdfPageResources) HasColorspaceByName(keyName _eb.PdfObjectName) bool {
	_fbdbd, _egecg := _afcfad.GetColorspaces()
	if _egecg != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _egecg)
		return false
	}
	if _fbdbd == nil {
		return false
	}
	_, _dgfbc := _fbdbd.Colorspaces[string(keyName)]
	return _dgfbc
}

// NewPdfActionGoTo returns a new "go to" action.
func NewPdfActionGoTo() *PdfActionGoTo {
	_gac := NewPdfAction()
	_bdg := &PdfActionGoTo{}
	_bdg.PdfAction = _gac
	_gac.SetContext(_bdg)
	return _bdg
}

// NewReaderOpts generates a default `ReaderOpts` instance.
func NewReaderOpts() *ReaderOpts { return &ReaderOpts{Password: "", LazyLoad: true} }

// NewPdfWriter initializes a new PdfWriter.
func NewPdfWriter() PdfWriter {
	_aeecd := PdfWriter{}
	_aeecd._aeeda = map[_eb.PdfObject]struct{}{}
	_aeecd._dcfgf = []_eb.PdfObject{}
	_aeecd._cfggb = map[_eb.PdfObject][]*_eb.PdfObjectDictionary{}
	_aeecd._fabeca = map[_eb.PdfObject]struct{}{}
	_aeecd._edbbf.Major = 1
	_aeecd._edbbf.Minor = 3
	_degff := _eb.MakeDict()
	_gdaaed := []struct {
		_bdccef _eb.PdfObjectName
		_cggef  string
	}{{"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _eabfd()}, {"\u0043r\u0065\u0061\u0074\u006f\u0072", _cdeea()}, {"\u0041\u0075\u0074\u0068\u006f\u0072", _agedg()}, {"\u0053u\u0062\u006a\u0065\u0063\u0074", _aecd()}, {"\u0054\u0069\u0074l\u0065", _efebc()}, {"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _ebbea()}}
	for _, _dfcdg := range _gdaaed {
		if _dfcdg._cggef != "" {
			_degff.Set(_dfcdg._bdccef, _eb.MakeString(_dfcdg._cggef))
		}
	}
	if _beafa := _gaffc(); !_beafa.IsZero() {
		if _acffe, _fbfbd := NewPdfDateFromTime(_beafa); _fbfbd == nil {
			_degff.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _acffe.ToPdfObject())
		}
	}
	if _gdaea := _fdegcc(); !_gdaea.IsZero() {
		if _ddaae, _gfcgb := NewPdfDateFromTime(_gdaea); _gfcgb == nil {
			_degff.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _ddaae.ToPdfObject())
		}
	}
	_begcbg := _eb.PdfIndirectObject{}
	_begcbg.PdfObject = _degff
	_aeecd._ecagd = &_begcbg
	_aeecd.addObject(&_begcbg)
	_fdbcgc := _eb.PdfIndirectObject{}
	_bfgg := _eb.MakeDict()
	_bfgg.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0043a\u0074\u0061\u006c\u006f\u0067"))
	_fdbcgc.PdfObject = _bfgg
	_aeecd._ccea = &_fdbcgc
	_aeecd.addObject(_aeecd._ccea)
	_adbeg, _bffcf := _gaeaca("\u0077")
	if _bffcf != nil {
		_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bffcf)
	}
	_aeecd._badcb = _adbeg
	_bcbcf := _eb.PdfIndirectObject{}
	_fgbfc := _eb.MakeDict()
	_fgbfc.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0050\u0061\u0067e\u0073"))
	_fgcgd := _eb.PdfObjectArray{}
	_fgbfc.Set("\u004b\u0069\u0064\u0073", &_fgcgd)
	_fgbfc.Set("\u0043\u006f\u0075n\u0074", _eb.MakeInteger(0))
	_bcbcf.PdfObject = _fgbfc
	_aeecd._afga = &_bcbcf
	_aeecd._gbgc = map[_eb.PdfObject]struct{}{}
	_aeecd._aadfg = []*_eb.PdfIndirectObject{}
	_aeecd.addObject(_aeecd._afga)
	_bfgg.Set("\u0050\u0061\u0067e\u0073", &_bcbcf)
	_aeecd._dbffa = _bfgg
	_ddb.Log.Trace("\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0025\u0073", _fdbcgc)
	return _aeecd
}

// ToPdfObject returns a PDF object representation of the outline destination.
func (_ecgb OutlineDest) ToPdfObject() _eb.PdfObject {
	if (_ecgb.PageObj == nil && _ecgb.Page < 0) || _ecgb.Mode == "" {
		return _eb.MakeNull()
	}
	_bbdag := _eb.MakeArray()
	if _ecgb.PageObj != nil {
		_bbdag.Append(_ecgb.PageObj)
	} else {
		_bbdag.Append(_eb.MakeInteger(_ecgb.Page))
	}
	_bbdag.Append(_eb.MakeName(_ecgb.Mode))
	switch _ecgb.Mode {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		_bbdag.Append(_eb.MakeFloat(_ecgb.Y))
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		_bbdag.Append(_eb.MakeFloat(_ecgb.X))
	case "\u0058\u0059\u005a":
		_bbdag.Append(_eb.MakeFloat(_ecgb.X))
		_bbdag.Append(_eb.MakeFloat(_ecgb.Y))
		_bbdag.Append(_eb.MakeFloat(_ecgb.Zoom))
	default:
		_bbdag.Set(1, _eb.MakeName("\u0046\u0069\u0074"))
	}
	return _bbdag
}

// NewPdfAppenderWithOpts creates a new Pdf appender from a Pdf reader with options.
func NewPdfAppenderWithOpts(reader *PdfReader, opts *ReaderOpts, encryptOptions *EncryptOptions) (*PdfAppender, error) {
	_befb := &PdfAppender{_bccd: reader._cbeg, Reader: reader, _ebcb: reader._ebbe, _fdfg: reader._bcefc}
	_dfbe, _dga := _befb._bccd.Seek(0, _bagf.SeekEnd)
	if _dga != nil {
		return nil, _dga
	}
	_befb._eabd = _dfbe
	if _, _dga = _befb._bccd.Seek(0, _bagf.SeekStart); _dga != nil {
		return nil, _dga
	}
	_befb._adac, _dga = NewPdfReaderWithOpts(_befb._bccd, opts)
	if _dga != nil {
		return nil, _dga
	}
	for _, _ccga := range _befb.Reader.GetObjectNums() {
		if _befb._cfgdf < _ccga {
			_befb._cfgdf = _ccga
		}
	}
	_befb._dce = _befb._ebcb.GetXrefTable()
	_befb._dgc = _befb._ebcb.GetXrefOffset()
	_befb._fedg = append(_befb._fedg, _befb._adac.PageList...)
	_befb._accfg = make(map[_eb.PdfObject]struct{})
	_befb._ebcbc = make(map[_eb.PdfObject]int64)
	_befb._dbae = make(map[_eb.PdfObject]struct{})
	_befb._bddda = _befb._adac.AcroForm
	_befb._cedc = _befb._adac.DSS
	if opts != nil {
		_befb._cedcf = opts.Password
	}
	if encryptOptions != nil {
		_befb._fcfd = encryptOptions
	}
	return _befb, nil
}

const (
	BorderEffectNoEffect BorderEffect = iota
	BorderEffectCloudy   BorderEffect = iota
)

func (_gcfee *PdfReader) newPdfAnnotationPopupFromDict(_fcf *_eb.PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	_gfg := PdfAnnotationPopup{}
	_gfg.Parent = _fcf.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	_gfg.Open = _fcf.Get("\u004f\u0070\u0065\u006e")
	return &_gfg, nil
}

// PdfAnnotationStamp represents Stamp annotations.
// (Section 12.5.6.12).
type PdfAnnotationStamp struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Name _eb.PdfObject
}

// SetCatalogMarkInfo sets the catalog MarkInfo dictionary.
func (_cdede *PdfWriter) SetCatalogMarkInfo(info _eb.PdfObject) error {
	if info == nil {
		_cdede._dbffa.Remove("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
		return nil
	}
	if _egddg, _dgcdeb := info.(*_eb.PdfObjectReference); _dgcdeb {
		info = _egddg.Resolve()
		if info == nil {
			_cdede._dbffa.Remove("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
			return nil
		}
	}
	_cdede.addObject(info)
	_cdede._dbffa.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", info)
	return nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a grayscale device.
func (_dgad *PdfColorspaceDeviceGray) GetNumComponents() int { return 1 }
func (_dbeg *PdfWriter) setDocInfo(_aecdg _eb.PdfObject) {
	if _dbeg.hasObject(_dbeg._ecagd) {
		delete(_dbeg._aeeda, _dbeg._ecagd)
		delete(_dbeg._fabeca, _dbeg._ecagd)
		for _dbagg, _bcfbc := range _dbeg._dcfgf {
			if _bcfbc == _dbeg._ecagd {
				copy(_dbeg._dcfgf[_dbagg:], _dbeg._dcfgf[_dbagg+1:])
				_dbeg._dcfgf[len(_dbeg._dcfgf)-1] = nil
				_dbeg._dcfgf = _dbeg._dcfgf[:len(_dbeg._dcfgf)-1]
				break
			}
		}
	}
	_affbfcc := _eb.PdfIndirectObject{}
	_affbfcc.PdfObject = _aecdg
	_dbeg._ecagd = &_affbfcc
	_dbeg.addObject(&_affbfcc)
}
func (_daf *PdfReader) newPdfAnnotationUnderlineFromDict(_ceeb *_eb.PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	_eagg := PdfAnnotationUnderline{}
	_dedge, _bcgg := _daf.newPdfAnnotationMarkupFromDict(_ceeb)
	if _bcgg != nil {
		return nil, _bcgg
	}
	_eagg.PdfAnnotationMarkup = _dedge
	_eagg.QuadPoints = _ceeb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_eagg, nil
}

// PdfActionGoTo represents a GoTo action.
type PdfActionGoTo struct {
	*PdfAction
	D _eb.PdfObject
}

// PrintPageRange returns the value of the printPageRange.
func (_dffef *ViewerPreferences) PrintPageRange() []int { return _dffef._gcgfc }

// ToImage converts an object to an Image which can be transformed or saved out.
// The image data is decoded and the Image returned.
func (_bgfae *XObjectImage) ToImage() (*Image, error) {
	_edda := &Image{}
	if _bgfae.Height == nil {
		return nil, _dcf.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_edda.Height = *_bgfae.Height
	if _bgfae.Width == nil {
		return nil, _dcf.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_edda.Width = *_bgfae.Width
	if _bgfae.BitsPerComponent == nil {
		switch _bgfae.Filter.(type) {
		case *_eb.CCITTFaxEncoder, *_eb.JBIG2Encoder:
			_edda.BitsPerComponent = 1
		case *_eb.LZWEncoder, *_eb.RunLengthEncoder:
			_edda.BitsPerComponent = 8
		default:
			return nil, _dcf.New("\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006fm\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
	} else {
		_edda.BitsPerComponent = *_bgfae.BitsPerComponent
	}
	_edda.ColorComponents = _bgfae.ColorSpace.GetNumComponents()
	_bgfae._gceffg.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _eb.MakeInteger(int64(_edda.ColorComponents)))
	_cgdfd, _ebeab := _eb.DecodeStream(_bgfae._gceffg)
	if _ebeab != nil {
		return nil, _ebeab
	}
	_edda.Data = _cgdfd
	if _bgfae.Decode != nil {
		_ebcfa, _eacac := _bgfae.Decode.(*_eb.PdfObjectArray)
		if !_eacac {
			_ddb.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020D\u0065\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074")
			return nil, _dcf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_eacba, _ffede := _ebcfa.ToFloat64Array()
		if _ffede != nil {
			return nil, _ffede
		}
		switch _bgfae.ColorSpace.(type) {
		case *PdfColorspaceDeviceCMYK:
			_cegaae := _bgfae.ColorSpace.DecodeArray()
			if _cegaae[0] == _eacba[0] && _cegaae[1] == _eacba[1] && _cegaae[2] == _eacba[2] && _cegaae[3] == _eacba[3] {
				_edda._fedc = _eacba
			}
		default:
			_edda._fedc = _eacba
		}
	}
	return _edda, nil
}

// PdfFieldButton represents a button field which includes push buttons, checkboxes, and radio buttons.
type PdfFieldButton struct {
	*PdfField
	Opt   *_eb.PdfObjectArray
	_cadc *Image
}

// GenerateXObjectName generates an unused XObject name that can be used for
// adding new XObjects. Uses format XObj1, XObj2, ...
func (_gbagf *PdfPageResources) GenerateXObjectName() _eb.PdfObjectName {
	_ceccd := 1
	for {
		_cgcd := _eb.MakeName(_e.Sprintf("\u0058\u004f\u0062\u006a\u0025\u0064", _ceccd))
		if !_gbagf.HasXObjectByName(*_cgcd) {
			return *_cgcd
		}
		_ceccd++
	}
}

// DecodeArray returns the range of color component values in CalRGB colorspace.
func (_geae *PdfColorspaceCalRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// GetXObjectByName gets XObject by name.
func (_cfeca *PdfPage) GetXObjectByName(name _eb.PdfObjectName) (_eb.PdfObject, bool) {
	_bcbea, _daaad := _cfeca.Resources.XObject.(*_eb.PdfObjectDictionary)
	if !_daaad {
		return nil, false
	}
	if _fbeee := _bcbea.Get(name); _fbeee != nil {
		return _fbeee, true
	}
	return nil, false
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a CalGray device.
func (_ggfag *PdfColorspaceCalGray) GetNumComponents() int { return 1 }

// ToPdfObject implements interface PdfModel.
func (_cb *PdfActionMovie) ToPdfObject() _eb.PdfObject {
	_cb.PdfAction.ToPdfObject()
	_ee := _cb._dee
	_ggb := _ee.PdfObject.(*_eb.PdfObjectDictionary)
	_ggb.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeMovie)))
	_ggb.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e", _cb.Annotation)
	_ggb.SetIfNotNil("\u0054", _cb.T)
	_ggb.SetIfNotNil("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn", _cb.Operation)
	return _ee
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the page using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened for the page.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
func (_acac *PdfPage) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	_feaa := map[*PdfAnnotation]bool{}
	_fdege, _bbbfc := _acac.GetAnnotations()
	if _bbbfc != nil {
		return _bbbfc
	}
	_aceb := false
	for _, _fecd := range _fdege {
		if opts.AnnotFilterFunc != nil {
			_feaa[_fecd] = opts.AnnotFilterFunc(_fecd)
		} else {
			_feaa[_fecd] = true
		}
		if _feaa[_fecd] {
			_aceb = true
		}
	}
	if !_aceb {
		return nil
	}
	return _acac.flattenFieldsWithOpts(appgen, opts, _feaa)
}

// DecodeArray returns the component range values for the Separation colorspace.
func (_dcdf *PdfColorspaceSpecialSeparation) DecodeArray() []float64 { return []float64{0, 1.0} }
func _dcgfe(_dada _bagf.ReadSeeker, _gdcag *ReaderOpts, _bdceaf bool, _abedc string) (*PdfReader, error) {
	if _gdcag == nil {
		_gdcag = NewReaderOpts()
	}
	_ffed := ""
	if _eaaag, _ccdea := _dada.(*_ccb.File); _ccdea {
		_ffed = _eaaag.Name()
	}
	_ffde := *_gdcag
	_adbed := &PdfReader{_cbeg: _dada, _bcefc: map[_eb.PdfObject]struct{}{}, _affaf: _fcfc(), _cfcgdf: _gdcag.LazyLoad, _fafga: _gdcag.ComplianceMode, _daag: _bdceaf, _eacdg: &_ffde, _fbgfg: _ffed}
	var _bfff error
	// _efea, _bfff := _gaeaca("\u0072")
	// if _bfff != nil {
	// 	return nil, _bfff
	// }
	// _bfff = _cf.Track(_efea, _abedc, _adbed._fbgfg)
	// if _bfff != nil {
	// 	return nil, _bfff
	// }
	// _adbed._cadbf = _efea
	var _bafab *_eb.PdfParser
	if !_adbed._fafga {
		_bafab, _bfff = _eb.NewParser(_dada)
	} else {
		_bafab, _bfff = _eb.NewCompliancePdfParser(_dada)
	}
	if _bfff != nil {
		return nil, _bfff
	}
	_adbed._ebbe = _bafab
	_ebabc, _bfff := _adbed.IsEncrypted()
	if _bfff != nil {
		return nil, _bfff
	}
	if !_ebabc {
		_bfff = _adbed.loadStructure()
		if _bfff != nil {
			return nil, _bfff
		}
	} else if _bdceaf {
		_debf, _fgaa := _adbed.Decrypt([]byte(_gdcag.Password))
		if _fgaa != nil {
			return nil, _fgaa
		}
		if !_debf {
			return nil, _dcf.New("\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0064\u0065c\u0072\u0079\u0070\u0074\u0020\u0070\u0061\u0073\u0073w\u006f\u0072\u0064\u0020p\u0072\u006f\u0074\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u002d\u0020\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0073\u0070\u0065\u0063\u0069\u0066y\u0020\u0070\u0061s\u0073\u0020\u0074\u006f\u0020\u0044\u0065\u0063\u0072\u0079\u0070\u0074")
		}
	}
	_adbed._dbag = make(map[*PdfReader]*PdfReader)
	_adbed._dfeg = make([]*PdfReader, _bafab.GetRevisionNumber())
	return _adbed, nil
}

// GetContainingPdfObject returns the container of the pattern object (indirect object).
func (_gaebf *PdfPattern) GetContainingPdfObject() _eb.PdfObject { return _gaebf._agddd }

// NewPdfColorPatternType2 returns an empty color shading pattern type 2 (Axial).
func NewPdfColorPatternType2() *PdfColorPatternType2 {
	_cebgg := &PdfColorPatternType2{}
	return _cebgg
}
func (_eec *PdfReader) newPdfActionRenditionFromDict(_fed *_eb.PdfObjectDictionary) (*PdfActionRendition, error) {
	return &PdfActionRendition{R: _fed.Get("\u0052"), AN: _fed.Get("\u0041\u004e"), OP: _fed.Get("\u004f\u0050"), JS: _fed.Get("\u004a\u0053")}, nil
}

// SetContext set the sub annotation (context).
func (_dgcde *PdfShading) SetContext(ctx PdfModel) { _dgcde._ecffg = ctx }

// RepairAcroForm attempts to rebuild the AcroForm fields using the widget
// annotations present in the document pages. Pass nil for the opts parameter
// in order to use the default options.
// NOTE: Currently, the opts parameter is declared in order to enable adding
// future options, but passing nil will always result in the default options
// being used.
func (_fccec *PdfReader) RepairAcroForm(opts *AcroFormRepairOptions) error {
	var _geega []*PdfField
	_fdcbf := map[*_eb.PdfIndirectObject]struct{}{}
	for _, _fbbfe := range _fccec.PageList {
		_bgeg, _cafedf := _fbbfe.GetAnnotations()
		if _cafedf != nil {
			return _cafedf
		}
		for _, _dedea := range _bgeg {
			var _ccdfb *PdfField
			switch _aefea := _dedea.GetContext().(type) {
			case *PdfAnnotationWidget:
				if _aefea._bca != nil {
					_ccdfb = _aefea._bca
					break
				}
				if _gfbef, _cedfb := _eb.GetIndirect(_aefea.Parent); _cedfb {
					_ccdfb, _cafedf = _fccec.newPdfFieldFromIndirectObject(_gfbef, nil)
					if _cafedf == nil {
						break
					}
					_ddb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _gfbef, _cafedf)
				}
				if _aefea._ggf != nil {
					_ccdfb, _cafedf = _fccec.newPdfFieldFromIndirectObject(_aefea._ggf, nil)
					if _cafedf == nil {
						break
					}
					_ddb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _aefea._ggf, _cafedf)
				}
			}
			if _ccdfb == nil {
				continue
			}
			if _, _fbdad := _fdcbf[_ccdfb._adgda]; _fbdad {
				continue
			}
			_fdcbf[_ccdfb._adgda] = struct{}{}
			_geega = append(_geega, _ccdfb)
		}
	}
	if len(_geega) == 0 {
		return nil
	}
	if _fccec.AcroForm == nil {
		_fccec.AcroForm = NewPdfAcroForm()
	}
	_fccec.AcroForm.Fields = &_geega
	return nil
}

// NonFullScreenPageMode represents the document’s page mode when exiting
// full-screen mode.
type NonFullScreenPageMode string

func (_egacc *PdfWriter) updateObjectNumbers() {
	_afcfbg := _egacc.ObjNumOffset
	_fdecd := 0
	for _, _accgc := range _egacc._dcfgf {
		_ffbga := int64(_fdecd + 1 + _afcfbg)
		_agbfc := true
		if _egacc._caed {
			if _eeacd, _adfge := _egacc._ecbgf[_accgc]; _adfge {
				_ffbga = _eeacd
				_agbfc = false
			}
		}
		switch _fccgc := _accgc.(type) {
		case *_eb.PdfIndirectObject:
			_fccgc.ObjectNumber = _ffbga
			_fccgc.GenerationNumber = 0
		case *_eb.PdfObjectStream:
			_fccgc.ObjectNumber = _ffbga
			_fccgc.GenerationNumber = 0
		case *_eb.PdfObjectStreams:
			_fccgc.ObjectNumber = _ffbga
			_fccgc.GenerationNumber = 0
		case *_eb.PdfObjectReference:
			_fccgc.ObjectNumber = _ffbga
			_fccgc.GenerationNumber = 0
		case *_eb.PdfObjectDictionary, *_eb.PdfObjectString:
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u0020%\u0054\u0020\u002d\u0020\u0073\u006b\u0069p\u0070\u0069\u006e\u0067", _fccgc)
			continue
		}
		if _agbfc {
			_fdecd++
		}
	}
	_cfafe := func(_fbcec _eb.PdfObject) int64 {
		switch _deecd := _fbcec.(type) {
		case *_eb.PdfIndirectObject:
			return _deecd.ObjectNumber
		case *_eb.PdfObjectStream:
			return _deecd.ObjectNumber
		case *_eb.PdfObjectStreams:
			return _deecd.ObjectNumber
		case *_eb.PdfObjectReference:
			return _deecd.ObjectNumber
		}
		return 0
	}
	_ba.SliceStable(_egacc._dcfgf, func(_ceebb, _fbdbb int) bool { return _cfafe(_egacc._dcfgf[_ceebb]) < _cfafe(_egacc._dcfgf[_fbdbb]) })
}

// PdfAnnotationCircle represents Circle annotations.
// (Section 12.5.6.8).
type PdfAnnotationCircle struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _eb.PdfObject
	IC _eb.PdfObject
	BE _eb.PdfObject
	RD _eb.PdfObject
}

func (_efeb *PdfFont) baseFields() *fontCommon {
	if _efeb._fdaa == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0062\u0061\u0073\u0065\u0046\u0069\u0065l\u0064s\u002e \u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e")
		return nil
	}
	return _efeb._fdaa.baseFields()
}

// GetPageAsIndirectObject returns the page as a dictionary within an PdfIndirectObject.
func (_bdcdd *PdfPage) GetPageAsIndirectObject() *_eb.PdfIndirectObject { return _bdcdd._efcff }

// ToPdfObject returns a PDF object representation of the ViewerPreferences.
func (_fcbcbb *ViewerPreferences) ToPdfObject() _eb.PdfObject {
	_eceaa := _eb.MakeDict()
	if _fcbcbb._caafc != nil {
		_eceaa.Set("H\u0069\u0064\u0065\u0054\u006f\u006f\u006c\u0062\u0061\u0072", _eb.MakeBool(*_fcbcbb._caafc))
	}
	if _fcbcbb._gcbcb != nil {
		_eceaa.Set("H\u0069\u0064\u0065\u004d\u0065\u006e\u0075\u0062\u0061\u0072", _eb.MakeBool(*_fcbcbb._gcbcb))
	}
	if _fcbcbb._fcebae != nil {
		_eceaa.Set("\u0048\u0069\u0064e\u0057\u0069\u006e\u0064\u006f\u0077\u0055\u0049", _eb.MakeBool(*_fcbcbb._fcebae))
	}
	if _fcbcbb._edbaa != nil {
		_eceaa.Set("\u0046i\u0074\u0057\u0069\u006e\u0064\u006fw", _eb.MakeBool(*_fcbcbb._edbaa))
	}
	if _fcbcbb._dacc != nil {
		_eceaa.Set("\u0043\u0065\u006et\u0065\u0072\u0057\u0069\u006e\u0064\u006f\u0077", _eb.MakeBool(*_fcbcbb._dacc))
	}
	if _fcbcbb._bgdge != nil {
		_eceaa.Set("\u0044i\u0073p\u006c\u0061\u0079\u0044\u006f\u0063\u0054\u0069\u0074\u006c\u0065", _eb.MakeBool(*_fcbcbb._bgdge))
	}
	if _fcbcbb._cdcff != "" {
		_eceaa.Set("N\u006f\u006e\u0046\u0075ll\u0053c\u0072\u0065\u0065\u006e\u0050a\u0067\u0065\u004d\u006f\u0064\u0065", _eb.MakeName(string(_fcbcbb._cdcff)))
	}
	if _fcbcbb._baeb != "" {
		_eceaa.Set("\u0044i\u0072\u0065\u0063\u0074\u0069\u006fn", _eb.MakeName(string(_fcbcbb._baeb)))
	}
	if _fcbcbb._ffgda != "" {
		_eceaa.Set("\u0056\u0069\u0065\u0077\u0041\u0072\u0065\u0061", _eb.MakeName(string(_fcbcbb._ffgda)))
	}
	if _fcbcbb._eadae != "" {
		_eceaa.Set("\u0056\u0069\u0065\u0077\u0043\u006c\u0069\u0070", _eb.MakeName(string(_fcbcbb._eadae)))
	}
	if _fcbcbb._bdegd != "" {
		_eceaa.Set("\u0050r\u0069\u006e\u0074\u0041\u0072\u0065a", _eb.MakeName(string(_fcbcbb._bdegd)))
	}
	if _fcbcbb._edef != "" {
		_eceaa.Set("\u0050r\u0069\u006e\u0074\u0043\u006c\u0069p", _eb.MakeName(string(_fcbcbb._edef)))
	}
	if _fcbcbb._gefcg != "" {
		_eceaa.Set("\u0050\u0072\u0069n\u0074\u0053\u0063\u0061\u006c\u0069\u006e\u0067", _eb.MakeName(string(_fcbcbb._gefcg)))
	}
	if _fcbcbb._ebabd != "" {
		_eceaa.Set("\u0044\u0075\u0070\u006c\u0065\u0078", _eb.MakeName(string(_fcbcbb._ebabd)))
	}
	if _fcbcbb._gggea != nil {
		_eceaa.Set("\u0050\u0069\u0063\u006b\u0054\u0072\u0061\u0079\u0042\u0079\u0050\u0044F\u0053\u0069\u007a\u0065", _eb.MakeBool(*_fcbcbb._gggea))
	}
	if _fcbcbb._gcgfc != nil {
		_eceaa.Set("\u0050\u0072\u0069\u006e\u0074\u0050\u0061\u0067\u0065R\u0061\u006e\u0067\u0065", _eb.MakeArrayFromIntegers(_fcbcbb._gcgfc))
	}
	if _fcbcbb._bdeec != 0 {
		_eceaa.Set("\u004eu\u006d\u0043\u006f\u0070\u0069\u0065s", _eb.MakeInteger(int64(_fcbcbb._bdeec)))
	}
	return _eceaa
}

// PdfPageResourcesColorspaces contains the colorspace in the PdfPageResources.
// Needs to have matching name and colorspace map entry. The Names define the order.
type PdfPageResourcesColorspaces struct {
	Names       []string
	Colorspaces map[string]PdfColorspace
	_ccba       *_eb.PdfIndirectObject
}

func (_acgfg *PdfWriter) mapObjectStreams(_ebfeec bool) (map[_eb.PdfObject]bool, bool) {
	_cdgbg := make(map[_eb.PdfObject]bool)
	for _, _abbga := range _acgfg._dcfgf {
		if _dcefad, _cgdfg := _abbga.(*_eb.PdfObjectStreams); _cgdfg {
			_ebfeec = true
			for _, _babdcd := range _dcefad.Elements() {
				_cdgbg[_babdcd] = true
				if _acege, _bdfag := _babdcd.(*_eb.PdfIndirectObject); _bdfag {
					_cdgbg[_acege.PdfObject] = true
				}
			}
		}
	}
	return _cdgbg, _ebfeec
}

// PdfAnnotationCaret represents Caret annotations.
// (Section 12.5.6.11).
type PdfAnnotationCaret struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	RD _eb.PdfObject
	Sy _eb.PdfObject
}

// SubsetRegistered subsets the font to only the glyphs that have been registered by the encoder.
//
// NOTE: This only works on fonts that support subsetting. For unsupported fonts this is a no-op, although a debug
// message is emitted.  Currently supported fonts are embedded Truetype CID fonts (type 0).
//
// NOTE: Make sure to call this soon before writing (once all needed runes have been registered).
// If using package creator, use its EnableFontSubsetting method instead.
func (_aeaa *PdfFont) SubsetRegistered() error {
	switch _cfdc := _aeaa._fdaa.(type) {
	case *pdfFontType0:
		_cfddd := _cfdc.subsetRegistered()
		if _cfddd != nil {
			_ddb.Log.Debug("\u0053\u0075b\u0073\u0065\u0074 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _cfddd)
			return _cfddd
		}
		if _cfdc._eecg != nil {
			if _cfdc._edcff != nil {
				_cfdc._edcff.ToPdfObject()
			}
			_cfdc.ToPdfObject()
		}
	default:
		_ddb.Log.Debug("F\u006f\u006e\u0074\u0020\u0025\u0054 \u0064\u006f\u0065\u0073\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006fr\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069n\u0067", _cfdc)
	}
	return nil
}

// NewPdfShadingPatternType3 creates an empty shading pattern type 3 object.
func NewPdfShadingPatternType3() *PdfShadingPatternType3 {
	_efdagf := &PdfShadingPatternType3{}
	_efdagf.Matrix = _eb.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_efdagf.PdfPattern = &PdfPattern{}
	_efdagf.PdfPattern.PatternType = int64(*_eb.MakeInteger(2))
	_efdagf.PdfPattern._eefgb = _efdagf
	_efdagf.PdfPattern._agddd = _eb.MakeIndirectObject(_eb.MakeDict())
	return _efdagf
}

// NewXObjectImageFromImage creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
func NewXObjectImageFromImage(img *Image, cs PdfColorspace, encoder _eb.StreamEncoder) (*XObjectImage, error) {
	_cegfc := NewXObjectImage()
	return UpdateXObjectImageFromImage(_cegfc, img, cs, encoder)
}

// Val returns the value of the color.
func (_cfdd *PdfColorCalGray) Val() float64 { return float64(*_cfdd) }

// NewPdfColorspaceLab returns a new Lab colorspace object.
func NewPdfColorspaceLab() *PdfColorspaceLab {
	_afbg := &PdfColorspaceLab{}
	_afbg.BlackPoint = []float64{0.0, 0.0, 0.0}
	_afbg.Range = []float64{-100, 100, -100, 100}
	return _afbg
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// component PDF objects.
func (_fgcc *PdfColorspaceICCBased) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if _fgcc.Alternate == nil {
		if _fgcc.N == 1 {
			_affd := NewPdfColorspaceDeviceGray()
			return _affd.ColorFromPdfObjects(objects)
		} else if _fgcc.N == 3 {
			_gged := NewPdfColorspaceDeviceRGB()
			return _gged.ColorFromPdfObjects(objects)
		} else if _fgcc.N == 4 {
			_dbda := NewPdfColorspaceDeviceCMYK()
			return _dbda.ColorFromPdfObjects(objects)
		} else {
			return nil, _dcf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _fgcc.Alternate.ColorFromPdfObjects(objects)
}

var _eeece = map[string]string{"\u0053\u0079\u006d\u0062\u006f\u006c": "\u0053\u0079\u006d\u0062\u006f\u006c\u0045\u006e\u0063o\u0064\u0069\u006e\u0067", "\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073": "Z\u0061p\u0066\u0044\u0069\u006e\u0067\u0062\u0061\u0074s\u0045\u006e\u0063\u006fdi\u006e\u0067"}

const (
	NonFullScreenPageModeUseNone     NonFullScreenPageMode = "\u0055s\u0065\u004e\u006f\u006e\u0065"
	NonFullScreenPageModeUseOutlines NonFullScreenPageMode = "U\u0073\u0065\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"
	NonFullScreenPageModeUseThumbs   NonFullScreenPageMode = "\u0055s\u0065\u0054\u0068\u0075\u006d\u0062s"
	NonFullScreenPageModeUseOC       NonFullScreenPageMode = "\u0055\u0073\u0065O\u0043"
	DirectionL2R                     Direction             = "\u004c\u0032\u0052"
	DirectionR2L                     Direction             = "\u0052\u0032\u004c"
	PageBoundaryMediaBox             PageBoundary          = "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"
	PageBoundaryCropBox              PageBoundary          = "\u0043r\u006f\u0070\u0042\u006f\u0078"
	PageBoundaryBleedBox             PageBoundary          = "\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"
	PageBoundaryTrimBox              PageBoundary          = "\u0054r\u0069\u006d\u0042\u006f\u0078"
	PageBoundaryArtBox               PageBoundary          = "\u0041\u0072\u0074\u0042\u006f\u0078"
	PrintScalingNone                 PrintScaling          = "\u004e\u006f\u006e\u0065"
	PrintScalingAppDefault           PrintScaling          = "\u0041\u0070\u0070\u0044\u0065\u0066\u0061\u0075\u006c\u0074"
	DuplexNone                       Duplex                = "\u006e\u006f\u006e\u0065"
	DuplexSimplex                    Duplex                = "\u0053i\u006d\u0070\u006c\u0065\u0078"
	DuplexFlipShortEdge              Duplex                = "\u0044\u0075\u0070\u006cex\u0046\u006c\u0069\u0070\u0053\u0068\u006f\u0072\u0074\u0045\u0064\u0067\u0065"
	DuplexFlipLongEdge               Duplex                = "\u0044u\u0070l\u0065\u0078\u0046\u006c\u0069p\u004c\u006fn\u0067\u0045\u0064\u0067\u0065"
)

// NewPdfColorDeviceRGB returns a new PdfColorDeviceRGB based on the r,g,b component values.
func NewPdfColorDeviceRGB(r, g, b float64) *PdfColorDeviceRGB {
	_acdg := PdfColorDeviceRGB{r, g, b}
	return &_acdg
}

// ToPdfObject returns the PdfFontDescriptor as a PDF dictionary inside an indirect object.
func (_debcf *PdfFontDescriptor) ToPdfObject() _eb.PdfObject {
	_bccec := _eb.MakeDict()
	if _debcf._fdea == nil {
		_debcf._fdea = &_eb.PdfIndirectObject{}
	}
	_debcf._fdea.PdfObject = _bccec
	_bccec.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072"))
	if _debcf.FontName != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _debcf.FontName)
	}
	if _debcf.FontFamily != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079", _debcf.FontFamily)
	}
	if _debcf.FontStretch != nil {
		_bccec.Set("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068", _debcf.FontStretch)
	}
	if _debcf.FontWeight != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074", _debcf.FontWeight)
	}
	if _debcf.Flags != nil {
		_bccec.Set("\u0046\u006c\u0061g\u0073", _debcf.Flags)
	}
	if _debcf.FontBBox != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _debcf.FontBBox)
	}
	if _debcf.ItalicAngle != nil {
		_bccec.Set("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065", _debcf.ItalicAngle)
	}
	if _debcf.Ascent != nil {
		_bccec.Set("\u0041\u0073\u0063\u0065\u006e\u0074", _debcf.Ascent)
	}
	if _debcf.Descent != nil {
		_bccec.Set("\u0044e\u0073\u0063\u0065\u006e\u0074", _debcf.Descent)
	}
	if _debcf.Leading != nil {
		_bccec.Set("\u004ce\u0061\u0064\u0069\u006e\u0067", _debcf.Leading)
	}
	if _debcf.CapHeight != nil {
		_bccec.Set("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t", _debcf.CapHeight)
	}
	if _debcf.XHeight != nil {
		_bccec.Set("\u0058H\u0065\u0069\u0067\u0068\u0074", _debcf.XHeight)
	}
	if _debcf.StemV != nil {
		_bccec.Set("\u0053\u0074\u0065m\u0056", _debcf.StemV)
	}
	if _debcf.StemH != nil {
		_bccec.Set("\u0053\u0074\u0065m\u0048", _debcf.StemH)
	}
	if _debcf.AvgWidth != nil {
		_bccec.Set("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068", _debcf.AvgWidth)
	}
	if _debcf.MaxWidth != nil {
		_bccec.Set("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068", _debcf.MaxWidth)
	}
	if _debcf.MissingWidth != nil {
		_bccec.Set("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068", _debcf.MissingWidth)
	}
	if _debcf.FontFile != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065", _debcf.FontFile)
	}
	if _debcf.FontFile2 != nil {
		_bccec.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652", _debcf.FontFile2)
	}
	if _debcf.FontFile3 != nil {
		_bccec.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653", _debcf.FontFile3)
	}
	if _debcf.CharSet != nil {
		_bccec.Set("\u0043h\u0061\u0072\u0053\u0065\u0074", _debcf.CharSet)
	}
	if _debcf.Style != nil {
		_bccec.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _debcf.FontName)
	}
	if _debcf.Lang != nil {
		_bccec.Set("\u004c\u0061\u006e\u0067", _debcf.Lang)
	}
	if _debcf.FD != nil {
		_bccec.Set("\u0046\u0044", _debcf.FD)
	}
	if _debcf.CIDSet != nil {
		_bccec.Set("\u0043\u0049\u0044\u0053\u0065\u0074", _debcf.CIDSet)
	}
	return _debcf._fdea
}

// B returns the value of the blue component of the color.
func (_geeg *PdfColorDeviceRGB) B() float64 { return _geeg[2] }

// PrintArea returns the value of the printArea.
func (_bgebe *ViewerPreferences) PrintArea() PageBoundary { return _bgebe._bdegd }

// NewPdfColorPattern returns an empty color pattern.
func NewPdfColorPattern() *PdfColorPattern { _fcbfd := &PdfColorPattern{}; return _fcbfd }
func _aaccff(_faaae *PdfField) []*PdfField {
	_ggfee := []*PdfField{_faaae}
	for _, _dcfga := range _faaae.Kids {
		_ggfee = append(_ggfee, _aaccff(_dcfga)...)
	}
	return _ggfee
}

// Mask returns the uin32 bitmask for the specific flag.
func (_acbg FieldFlag) Mask() uint32 { return uint32(_acbg) }

// NewStructTreeRootFromPdfObject creates a new structure tree root from a PDF object.
func NewStructTreeRootFromPdfObject(obj _eb.PdfObject) (*StructTreeRoot, error) {
	_gbcba := _eb.ResolveReference(obj)
	_baceb, _ecagbe := _eb.GetDict(_gbcba)
	if !_ecagbe {
		return nil, _e.Errorf("\u0065\u0078\u0069\u0073\u0074\u0069\u006e\u0067 \u0073\u0074\u0072uc\u0074\u0075\u0072\u0065\u0020\u0074r\u0065\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_cbccg := &StructTreeRoot{}
	_cbccg._fggdf = _eb.MakeIndirectObject(_eb.MakeDict())
	if _cbdaf := _baceb.Get("\u004b"); _cbdaf != nil {
		_ffdg := _eb.ResolveReference(_cbdaf)
		_aeagb := _eb.ResolveReferencesDeep(_ffdg, nil)
		if _aeagb != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006fr\u0020\u0072\u0065\u0073\u006f\u006c\u0076\u0069\u006e\u0067 \u004b \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _aeagb)
		}
		_cbccg.K = []*KDict{}
		if _dfffc, _bcfdec := _eb.GetArray(_ffdg); _bcfdec {
			for _fefde := 0; _fefde < _dfffc.Len(); _fefde++ {
				_ddfae := _dfffc.Get(_fefde)
				_dcce, _ccbee := _gfgbd(_ddfae)
				if _ccbee != nil {
					return nil, _ccbee
				}
				_cbccg.K = append(_cbccg.K, _dcce)
			}
		} else {
			_cfefb, _fdbgb := _gfgbd(_ffdg)
			if _fdbgb != nil {
				return nil, _fdbgb
			}
			_cbccg.K = append(_cbccg.K, _cfefb)
		}
	}
	if _deegb := _baceb.Get("\u0049\u0044\u0054\u0072\u0065\u0065"); _deegb != nil {
		_cbccg.IDTree = _cafbb(_deegb)
	}
	if _babbc := _baceb.Get("\u0050\u0061\u0072\u0065\u006e\u0074\u0054\u0072\u0065\u0065"); _babbc != nil {
		_effad := _eb.ResolveReference(_babbc)
		if _egfgc, _afbff := _eb.GetDict(_effad); _afbff {
			_cbccg.ParentTree = _egfgc
		}
	}
	if _dggcdg := _baceb.Get("\u0050\u0061\u0072\u0065\u006e\u0074\u0054\u0072\u0065\u0065\u004e\u0065x\u0074\u004b\u0065\u0079"); _dggcdg != nil {
		_, _adfcc := _eb.GetInt(_dggcdg)
		if _adfcc {
			_ddgbfb, _dadf := _eb.GetNumberAsInt64(_dggcdg)
			if _dadf != nil {
				return nil, _dadf
			}
			_cbccg.ParentTreeNextKey = _ddgbfb
		}
	}
	if _aeabc := _baceb.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"); _aeabc != nil {
		switch _cedfbd := _aeabc.(type) {
		case *_eb.PdfIndirectObject:
			if _badbc, _dgfgd := _eb.GetDict(_cedfbd.PdfObject); _dgfgd {
				_cbccg.RoleMap = _badbc
			}
		case *_eb.PdfObjectDictionary:
			_cbccg.RoleMap = _cedfbd
		case *_eb.PdfObjectString:
			_cbccg.RoleMap = _cedfbd
		}
	}
	if _fggega := _baceb.Get("\u0043\u006c\u0061\u0073\u0073\u004d\u0061\u0070"); _fggega != nil {
		if _eadd, _dafcbg := _eb.GetDict(_fggega); _dafcbg {
			_cbccg.ClassMap = _eadd
		}
	}
	return _cbccg, nil
}

// IsPush returns true if the button field represents a push button, false otherwise.
func (_ebfff *PdfFieldButton) IsPush() bool { return _ebfff.GetType() == ButtonTypePush }

// XObjectType represents the type of an XObject.
type XObjectType int

// GetStandardApplier gets currently used StandardApplier..
func (_gddgf *PdfWriter) GetStandardApplier() StandardApplier { return _gddgf._ggbcd }
func (_fecbf *PdfPattern) getDict() *_eb.PdfObjectDictionary {
	if _effeb, _fbeec := _fecbf._agddd.(*_eb.PdfIndirectObject); _fbeec {
		_bafad, _ffdff := _effeb.PdfObject.(*_eb.PdfObjectDictionary)
		if !_ffdff {
			return nil
		}
		return _bafad
	} else if _aaebg, _baceg := _fecbf._agddd.(*_eb.PdfObjectStream); _baceg {
		return _aaebg.PdfObjectDictionary
	} else {
		_ddb.Log.Debug("\u0054r\u0079\u0069\u006e\u0067\u0020\u0074\u006f a\u0063\u0063\u0065\u0073\u0073\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063t \u0074\u0079\u0070e\u0020\u0028\u0025\u0054\u0029", _fecbf._agddd)
		return nil
	}
}

// AddCustomInfo adds a custom info into document info dictionary.
func (_ddeb *PdfInfo) AddCustomInfo(name string, value string) error {
	if _ddeb._cbfb == nil {
		_ddeb._cbfb = _eb.MakeDict()
	}
	if _, _gdea := _baffe[name]; _gdea {
		return _e.Errorf("\u0063\u0061\u006e\u006e\u006ft\u0020\u0075\u0073\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064 \u0069\u006e\u0066\u006f\u0020\u006b\u0065\u0079\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0063\u0075\u0073\u0074\u006f\u006d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u006b\u0065y", name)
	}
	_ddeb._cbfb.SetIfNotNil(*_eb.MakeName(name), _eb.MakeString(value))
	return nil
}
func (_ggc *PdfReader) newPdfActionFromIndirectObject(_bed *_eb.PdfIndirectObject) (*PdfAction, error) {
	_gdca, _aad := _bed.PdfObject.(*_eb.PdfObjectDictionary)
	if !_aad {
		return nil, _e.Errorf("\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _ggc._affaf.GetModelFromPrimitive(_gdca); model != nil {
		_ecb, _fecb := model.(*PdfAction)
		if !_fecb {
			return nil, _e.Errorf("\u0063\u0061c\u0068\u0065\u0064\u0020\u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0061\u0063ti\u006f\u006e")
		}
		return _ecb, nil
	}
	_adf := &PdfAction{}
	_adf._dee = _bed
	_ggc._affaf.Register(_gdca, _adf)
	if _gcf := _gdca.Get("\u0054\u0079\u0070\u0065"); _gcf != nil {
		_beee, _gacf := _gcf.(*_eb.PdfObjectName)
		if !_gacf {
			_ddb.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _gcf)
		} else {
			if *_beee != "\u0041\u0063\u0074\u0069\u006f\u006e" {
				_ddb.Log.Trace("\u0055\u006e\u0073u\u0073\u0070\u0065\u0063t\u0065\u0064\u0020\u0054\u0079\u0070\u0065 \u0021\u003d\u0020\u0041\u0063\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", *_beee)
			}
			_adf.Type = _beee
		}
	}
	if _afa := _gdca.Get("\u004e\u0065\u0078\u0074"); _afa != nil {
		_adf.Next = _afa
	}
	if _eef := _gdca.Get("\u0053"); _eef != nil {
		_adf.S = _eef
	}
	_dea, _aac := _adf.S.(*_eb.PdfObjectName)
	if !_aac {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _adf.S)
		return nil, _e.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0053\u0020\u006f\u0062\u006a\u0065c\u0074 \u0074y\u0070e\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _adf.S)
	}
	_acc := PdfActionType(_dea.String())
	switch _acc {
	case ActionTypeGoTo:
		_acb, _aca := _ggc.newPdfActionGotoFromDict(_gdca)
		if _aca != nil {
			return nil, _aca
		}
		_acb.PdfAction = _adf
		_adf._aee = _acb
		return _adf, nil
	case ActionTypeGoToR:
		_cfgf, _ceg := _ggc.newPdfActionGotoRFromDict(_gdca)
		if _ceg != nil {
			return nil, _ceg
		}
		_cfgf.PdfAction = _adf
		_adf._aee = _cfgf
		return _adf, nil
	case ActionTypeGoToE:
		_gdg, _adfb := _ggc.newPdfActionGotoEFromDict(_gdca)
		if _adfb != nil {
			return nil, _adfb
		}
		_gdg.PdfAction = _adf
		_adf._aee = _gdg
		return _adf, nil
	case ActionTypeLaunch:
		_bcg, _bdfg := _ggc.newPdfActionLaunchFromDict(_gdca)
		if _bdfg != nil {
			return nil, _bdfg
		}
		_bcg.PdfAction = _adf
		_adf._aee = _bcg
		return _adf, nil
	case ActionTypeThread:
		_gaga, _ggd := _ggc.newPdfActionThreadFromDict(_gdca)
		if _ggd != nil {
			return nil, _ggd
		}
		_gaga.PdfAction = _adf
		_adf._aee = _gaga
		return _adf, nil
	case ActionTypeURI:
		_cfc, _dab := _ggc.newPdfActionURIFromDict(_gdca)
		if _dab != nil {
			return nil, _dab
		}
		_cfc.PdfAction = _adf
		_adf._aee = _cfc
		return _adf, nil
	case ActionTypeSound:
		_gfe, _ede := _ggc.newPdfActionSoundFromDict(_gdca)
		if _ede != nil {
			return nil, _ede
		}
		_gfe.PdfAction = _adf
		_adf._aee = _gfe
		return _adf, nil
	case ActionTypeMovie:
		_eea, _aff := _ggc.newPdfActionMovieFromDict(_gdca)
		if _aff != nil {
			return nil, _aff
		}
		_eea.PdfAction = _adf
		_adf._aee = _eea
		return _adf, nil
	case ActionTypeHide:
		_dfd, _dbc := _ggc.newPdfActionHideFromDict(_gdca)
		if _dbc != nil {
			return nil, _dbc
		}
		_dfd.PdfAction = _adf
		_adf._aee = _dfd
		return _adf, nil
	case ActionTypeNamed:
		_afc, _ecbe := _ggc.newPdfActionNamedFromDict(_gdca)
		if _ecbe != nil {
			return nil, _ecbe
		}
		_afc.PdfAction = _adf
		_adf._aee = _afc
		return _adf, nil
	case ActionTypeSubmitForm:
		_dfb, _dge := _ggc.newPdfActionSubmitFormFromDict(_gdca)
		if _dge != nil {
			return nil, _dge
		}
		_dfb.PdfAction = _adf
		_adf._aee = _dfb
		return _adf, nil
	case ActionTypeResetForm:
		_ggcd, _fae := _ggc.newPdfActionResetFormFromDict(_gdca)
		if _fae != nil {
			return nil, _fae
		}
		_ggcd.PdfAction = _adf
		_adf._aee = _ggcd
		return _adf, nil
	case ActionTypeImportData:
		_edc, _dabf := _ggc.newPdfActionImportDataFromDict(_gdca)
		if _dabf != nil {
			return nil, _dabf
		}
		_edc.PdfAction = _adf
		_adf._aee = _edc
		return _adf, nil
	case ActionTypeSetOCGState:
		_ace, _dgb := _ggc.newPdfActionSetOCGStateFromDict(_gdca)
		if _dgb != nil {
			return nil, _dgb
		}
		_ace.PdfAction = _adf
		_adf._aee = _ace
		return _adf, nil
	case ActionTypeRendition:
		_eeaa, _ffag := _ggc.newPdfActionRenditionFromDict(_gdca)
		if _ffag != nil {
			return nil, _ffag
		}
		_eeaa.PdfAction = _adf
		_adf._aee = _eeaa
		return _adf, nil
	case ActionTypeTrans:
		_gad, _bgdg := _ggc.newPdfActionTransFromDict(_gdca)
		if _bgdg != nil {
			return nil, _bgdg
		}
		_gad.PdfAction = _adf
		_adf._aee = _gad
		return _adf, nil
	case ActionTypeGoTo3DView:
		_cdf, _ecf := _ggc.newPdfActionGoTo3DViewFromDict(_gdca)
		if _ecf != nil {
			return nil, _ecf
		}
		_cdf.PdfAction = _adf
		_adf._aee = _cdf
		return _adf, nil
	case ActionTypeJavaScript:
		_ddg, _ebg := _ggc.newPdfActionJavaScriptFromDict(_gdca)
		if _ebg != nil {
			return nil, _ebg
		}
		_ddg.PdfAction = _adf
		_adf._aee = _ddg
		return _adf, nil
	}
	_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006eg\u0020u\u006ek\u006eo\u0077\u006e\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _acc)
	return nil, nil
}
func (_gbab *PdfReader) newPdfAnnotationSoundFromDict(_dgba *_eb.PdfObjectDictionary) (*PdfAnnotationSound, error) {
	_deca := PdfAnnotationSound{}
	_bacc, _dgbe := _gbab.newPdfAnnotationMarkupFromDict(_dgba)
	if _dgbe != nil {
		return nil, _dgbe
	}
	_deca.PdfAnnotationMarkup = _bacc
	_deca.Name = _dgba.Get("\u004e\u0061\u006d\u0065")
	_deca.Sound = _dgba.Get("\u0053\u006f\u0075n\u0064")
	return &_deca, nil
}

// GetType returns the button field type which returns one of the following
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
func (_fbgaf *PdfFieldButton) GetType() ButtonType {
	_egac := ButtonTypeCheckbox
	if _fbgaf.Ff != nil {
		if (uint32(*_fbgaf.Ff) & FieldFlagPushbutton.Mask()) > 0 {
			_egac = ButtonTypePush
		} else if (uint32(*_fbgaf.Ff) & FieldFlagRadio.Mask()) > 0 {
			_egac = ButtonTypeRadio
		}
	}
	return _egac
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_gadea *PdfFunctionType2) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ddb.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aadba := []float64{0.0}
	if _gadea.C0 != nil {
		_aadba = _gadea.C0
	}
	_badcd := []float64{1.0}
	if _gadea.C1 != nil {
		_badcd = _gadea.C1
	}
	var _cbce []float64
	for _dgdbga := 0; _dgdbga < len(_aadba); _dgdbga++ {
		_abbc := _aadba[_dgdbga] + _gg.Pow(x[0], _gadea.N)*(_badcd[_dgdbga]-_aadba[_dgdbga])
		_cbce = append(_cbce, _abbc)
	}
	return _cbce, nil
}

// PdfFont represents an underlying font structure which can be of type:
// - Type0
// - Type1
// - TrueType
// etc.
type PdfFont struct{ _fdaa pdfFont }

// EmbeddedFile represents an embedded file.
type EmbeddedFile struct {
	Name         string
	Content      []byte
	FileType     string
	Description  string
	Relationship FileRelationship
	Hash         string
	CreationTime _d.Time
	ModTime      _d.Time
}

// UpdateXObjectImageFromImage creates a new XObject Image from an
// Image object `img` and default masks from xobjIn.
// The default masks are overridden if img.hasAlpha
// If `encoder` is nil, uses raw encoding (none).
func UpdateXObjectImageFromImage(xobjIn *XObjectImage, img *Image, cs PdfColorspace, encoder _eb.StreamEncoder) (*XObjectImage, error) {
	if encoder == nil {
		var _eaeff error
		encoder, _eaeff = img.getSuitableEncoder()
		if _eaeff != nil {
			_ddb.Log.Debug("F\u0061\u0069l\u0075\u0072\u0065\u0020\u006f\u006e\u0020\u0066\u0069\u006e\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0069\u0074\u0061b\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072,\u0020\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b\u0020\u0074\u006f\u0020R\u0061\u0077\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _eaeff)
			encoder = _eb.NewRawEncoder()
		}
	}
	encoder.UpdateParams(img.GetParamsDict())
	_efaf, _dbedgb := encoder.EncodeBytes(img.Data)
	if _dbedgb != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _dbedgb)
		return nil, _dbedgb
	}
	_bccbbg := NewXObjectImage()
	_gedda := img.Width
	_daedc := img.Height
	_bccbbg.Width = &_gedda
	_bccbbg.Height = &_daedc
	_bbeeda := img.BitsPerComponent
	_bccbbg.BitsPerComponent = &_bbeeda
	_bccbbg.Filter = encoder
	_bccbbg.Stream = _efaf
	if cs == nil {
		if img.ColorComponents == 1 {
			_bccbbg.ColorSpace = NewPdfColorspaceDeviceGray()
			if img.BitsPerComponent == 16 {
				switch encoder.(type) {
				case *_eb.DCTEncoder:
					_bccbbg.ColorSpace = NewPdfColorspaceDeviceRGB()
					_bbeeda = 8
					_bccbbg.BitsPerComponent = &_bbeeda
				}
			}
		} else if img.ColorComponents == 3 {
			_bccbbg.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			switch encoder.(type) {
			case *_eb.DCTEncoder:
				_bccbbg.ColorSpace = NewPdfColorspaceDeviceRGB()
			default:
				_bccbbg.ColorSpace = NewPdfColorspaceDeviceCMYK()
			}
		} else {
			return nil, _dcf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_bccbbg.ColorSpace = cs
	}
	if len(img._bdcab) != 0 {
		_ffec := NewXObjectImage()
		_ffec.Filter = encoder
		_degbfg, _agbac := encoder.EncodeBytes(img._bdcab)
		if _agbac != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _agbac)
			return nil, _agbac
		}
		_ffec.Stream = _degbfg
		_ffec.BitsPerComponent = _bccbbg.BitsPerComponent
		_ffec.Width = &img.Width
		_ffec.Height = &img.Height
		_ffec.ColorSpace = NewPdfColorspaceDeviceGray()
		_bccbbg.SMask = _ffec.ToPdfObject()
	} else {
		_bccbbg.SMask = xobjIn.SMask
		_bccbbg.ImageMask = xobjIn.ImageMask
		if _bccbbg.ColorSpace.GetNumComponents() == 1 {
			_ddbeg(_bccbbg)
		}
	}
	return _bccbbg, nil
}

// ToJBIG2Image converts current image to the core.JBIG2Image.
func (_bcbbg *Image) ToJBIG2Image() (*_eb.JBIG2Image, error) {
	_ecfe, _eefga := _bcbbg.ToGoImage()
	if _eefga != nil {
		return nil, _eefga
	}
	return _eb.GoImageToJBIG2(_ecfe, _eb.JB2ImageAutoThreshold)
}

// DecodeArray returns the range of color component values in DeviceRGB colorspace.
func (_gaec *PdfColorspaceDeviceRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// IsSimple returns true if `font` is a simple font.
func (_dced *PdfFont) IsSimple() bool { _, _fedfag := _dced._fdaa.(*pdfFontSimple); return _fedfag }

// PdfInfo holds document information that will overwrite
// document information global variables defined above.
type PdfInfo struct {
	Title        *_eb.PdfObjectString
	Author       *_eb.PdfObjectString
	Subject      *_eb.PdfObjectString
	Keywords     *_eb.PdfObjectString
	Creator      *_eb.PdfObjectString
	Producer     *_eb.PdfObjectString
	CreationDate *PdfDate
	ModifiedDate *PdfDate
	Trapped      *_eb.PdfObjectName
	_cbfb        *_eb.PdfObjectDictionary
}

// ToPdfObject implements interface PdfModel.
func (_bbdb *PdfAnnotationStamp) ToPdfObject() _eb.PdfObject {
	_bbdb.PdfAnnotation.ToPdfObject()
	_effd := _bbdb._ggf
	_eede := _effd.PdfObject.(*_eb.PdfObjectDictionary)
	_bbdb.PdfAnnotationMarkup.appendToPdfDictionary(_eede)
	_eede.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053\u0074\u0061m\u0070"))
	_eede.SetIfNotNil("\u004e\u0061\u006d\u0065", _bbdb.Name)
	return _effd
}

// GetPatternByName gets the pattern specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_fcbaaa *PdfPageResources) GetPatternByName(keyName _eb.PdfObjectName) (*PdfPattern, bool) {
	if _fcbaaa.Pattern == nil {
		return nil, false
	}
	_bfdba, _fdbdd := _eb.TraceToDirectObject(_fcbaaa.Pattern).(*_eb.PdfObjectDictionary)
	if !_fdbdd {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074t\u0065\u0072\u006e\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _fcbaaa.Pattern)
		return nil, false
	}
	if _eafde := _bfdba.Get(keyName); _eafde != nil {
		_cfbb, _aabae := _aefe(_eafde)
		if _aabae != nil {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0070\u0061\u0074\u0074\u0065\u0072n\u003a\u0020\u0025\u0076", _aabae)
			return nil, false
		}
		return _cfbb, true
	}
	return nil, false
}
func _ddeca(_bdegf _eb.PdfObject) (*PdfFunctionType3, error) {
	_fbedb := &PdfFunctionType3{}
	var _fcac *_eb.PdfObjectDictionary
	if _gdff, _egbbc := _bdegf.(*_eb.PdfIndirectObject); _egbbc {
		_ggfdg, _fdaf := _gdff.PdfObject.(*_eb.PdfObjectDictionary)
		if !_fdaf {
			return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fbedb._ddgdg = _gdff
		_fcac = _ggfdg
	} else if _eadgf, _dcgfg := _bdegf.(*_eb.PdfObjectDictionary); _dcgfg {
		_fcac = _eadgf
	} else {
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_affg, _eecdg := _eb.TraceToDirectObject(_fcac.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_eb.PdfObjectArray)
	if !_eecdg {
		_ddb.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _affg.Len() != 2 {
		_ddb.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_feecc, _fdfc := _affg.ToFloat64Array()
	if _fdfc != nil {
		return nil, _fdfc
	}
	_fbedb.Domain = _feecc
	_affg, _eecdg = _eb.TraceToDirectObject(_fcac.Get("\u0052\u0061\u006eg\u0065")).(*_eb.PdfObjectArray)
	if _eecdg {
		if _affg.Len() < 0 || _affg.Len()%2 != 0 {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_febc, _gece := _affg.ToFloat64Array()
		if _gece != nil {
			return nil, _gece
		}
		_fbedb.Range = _febc
	}
	_affg, _eecdg = _eb.TraceToDirectObject(_fcac.Get("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es")).(*_eb.PdfObjectArray)
	if !_eecdg {
		_ddb.Log.Error("\u0046\u0075\u006ect\u0069\u006f\u006e\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_fbedb.Functions = []PdfFunction{}
	for _, _cgdec := range _affg.Elements() {
		_fefe, _eggd := _cccfa(_cgdec)
		if _eggd != nil {
			return nil, _eggd
		}
		_fbedb.Functions = append(_fbedb.Functions, _fefe)
	}
	_affg, _eecdg = _eb.TraceToDirectObject(_fcac.Get("\u0042\u006f\u0075\u006e\u0064\u0073")).(*_eb.PdfObjectArray)
	if !_eecdg {
		_ddb.Log.Error("B\u006fu\u006e\u0064\u0073\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_fcbg, _fdfc := _affg.ToFloat64Array()
	if _fdfc != nil {
		return nil, _fdfc
	}
	_fbedb.Bounds = _fcbg
	if len(_fbedb.Bounds) != len(_fbedb.Functions)-1 {
		_ddb.Log.Error("B\u006f\u0075\u006e\u0064\u0073\u0020\u0028\u0025\u0064)\u0020\u0061\u006e\u0064\u0020\u006e\u0075m \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029 n\u006f\u0074 \u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067", len(_fbedb.Bounds), len(_fbedb.Functions))
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_affg, _eecdg = _eb.TraceToDirectObject(_fcac.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_eb.PdfObjectArray)
	if !_eecdg {
		_ddb.Log.Error("E\u006ec\u006f\u0064\u0065\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_eaga, _fdfc := _affg.ToFloat64Array()
	if _fdfc != nil {
		return nil, _fdfc
	}
	_fbedb.Encode = _eaga
	if len(_fbedb.Encode) != 2*len(_fbedb.Functions) {
		_ddb.Log.Error("\u004c\u0065\u006e\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0028\u0025\u0064\u0029 \u0061\u006e\u0064\u0020\u006e\u0075\u006d\u0020\u0066\u0075\u006e\u0063\u0074i\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029\u0020\u006e\u006f\u0074 m\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u0075\u0070", len(_fbedb.Encode), len(_fbedb.Functions))
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return _fbedb, nil
}

// GetNumComponents returns the number of color components (4 for CMYK32).
func (_eacd *PdfColorDeviceCMYK) GetNumComponents() int { return 4 }

// PdfAnnotationPolygon represents Polygon annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolygon struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _eb.PdfObject
	LE       _eb.PdfObject
	BS       _eb.PdfObject
	IC       _eb.PdfObject
	BE       _eb.PdfObject
	IT       _eb.PdfObject
	Measure  _eb.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_cea *PdfActionImportData) ToPdfObject() _eb.PdfObject {
	_cea.PdfAction.ToPdfObject()
	_ea := _cea._dee
	_gacd := _ea.PdfObject.(*_eb.PdfObjectDictionary)
	_gacd.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeImportData)))
	if _cea.F != nil {
		_gacd.Set("\u0046", _cea.F.ToPdfObject())
	}
	return _ea
}

// PdfAcroForm represents the AcroForm dictionary used for representation of form data in PDF.
type PdfAcroForm struct {
	Fields          *[]*PdfField
	NeedAppearances *_eb.PdfObjectBool
	SigFlags        *_eb.PdfObjectInteger
	CO              *_eb.PdfObjectArray
	DR              *PdfPageResources
	DA              *_eb.PdfObjectString
	Q               *_eb.PdfObjectInteger
	XFA             _eb.PdfObject

	// ADBEEchoSign extra objects from Adobe Acrobat, causing signature invalid if not exists.
	ADBEEchoSign _eb.PdfObject
	_fbgad       *_eb.PdfIndirectObject
	_cgedg       bool
}

// NumCopies returns the value of the numCopies.
func (_bccfc *ViewerPreferences) NumCopies() int { return _bccfc._bdeec }

// PdfShadingType3 is a Radial shading.
type PdfShadingType3 struct {
	*PdfShading
	Coords   *_eb.PdfObjectArray
	Domain   *_eb.PdfObjectArray
	Function []PdfFunction
	Extend   *_eb.PdfObjectArray
}

// ColorToRGB converts a color in Separation colorspace to RGB colorspace.
func (_dgdf *PdfColorspaceSpecialSeparation) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _dgdf.AlternateSpace == nil {
		return nil, _dcf.New("\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020c\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	return _dgdf.AlternateSpace.ColorToRGB(color)
}

// SignatureHandler interface defines the common functionality for PDF signature handlers, which
// need to be capable of validating digital signatures and signing PDF documents.
type SignatureHandler interface {

	// IsApplicable checks if a given signature dictionary `sig` is applicable for the signature handler.
	// For example a signature of type `adbe.pkcs7.detached` might not fit for a rsa.sha1 handler.
	IsApplicable(_acfae *PdfSignature) bool

	// Validate validates a PDF signature against a given digest (hash) such as that determined
	// for an input file. Returns validation results.
	Validate(_gfdge *PdfSignature, _daegf Hasher) (SignatureValidationResult, error)

	// InitSignature prepares the signature dictionary for signing. This involves setting all
	// necessary fields, and also allocating sufficient space to the Contents so that the
	// finalized signature can be inserted once the hash is calculated.
	InitSignature(_egbef *PdfSignature) error

	// NewDigest creates a new digest/hasher based on the signature dictionary and handler.
	NewDigest(_gebgf *PdfSignature) (Hasher, error)

	// Sign receives the hash `digest` (for example hash of an input file), and signs based
	// on the signature dictionary `sig` and applies the signature data to the signature
	// dictionary Contents field.
	Sign(_agfdb *PdfSignature, _ebecef Hasher) error
}

// PdfWriter handles outputing PDF content.
type PdfWriter struct {
	_ccea           *_eb.PdfIndirectObject
	_afga           *_eb.PdfIndirectObject
	_gbgc           map[_eb.PdfObject]struct{}
	_aadfg          []*_eb.PdfIndirectObject
	_dcfgf          []_eb.PdfObject
	_aeeda          map[_eb.PdfObject]struct{}
	_bggbe          []*_eb.PdfIndirectObject
	_fbgge          *PdfOutlineTreeNode
	_dbffa          *_eb.PdfObjectDictionary
	_daadd          []_eb.PdfObject
	_ecagd          *_eb.PdfIndirectObject
	_ccfbc          *_ga.Writer
	_dfabe          int64
	_ceffa          error
	_gadcaa         *_eb.PdfCrypt
	_ffcbb          *_eb.PdfObjectDictionary
	_ebdeed         *_eb.PdfIndirectObject
	_ddefd          *_eb.PdfObjectArray
	_edbbf          _eb.Version
	_bddggg         *bool
	_cfggb          map[_eb.PdfObject][]*_eb.PdfObjectDictionary
	_dcbec          *PdfAcroForm
	_agdbc          *Names
	_fbaag          Optimizer
	_ggbcd          StandardApplier
	_aggfdb         map[int]crossReference
	_gedbe          int64
	ObjNumOffset    int
	_caed           bool
	_bfdeb          _eb.XrefTable
	_becbf          int64
	_dfbcf          int64
	_ecbgf          map[_eb.PdfObject]int64
	_fabeca         map[_eb.PdfObject]struct{}
	_badcb          string
	_adeee          string
	_bbbff          []*PdfOutputIntent
	_beeaf          bool
	_cebgc, _bgddbe string
}

func _gacff(_fbce *_eb.PdfIndirectObject, _cgbb *_eb.PdfObjectDictionary) (*DSS, error) {
	if _fbce == nil {
		_fbce = _eb.MakeIndirectObject(nil)
	}
	_fbce.PdfObject = _eb.MakeDict()
	_dbge := map[string]*VRI{}
	if _cfdg, _fcbaa := _eb.GetDict(_cgbb.Get("\u0056\u0052\u0049")); _fcbaa {
		for _, _degg := range _cfdg.Keys() {
			if _fgccg, _fdbe := _eb.GetDict(_cfdg.Get(_degg)); _fdbe {
				_dbge[_cc.ToUpper(_degg.String())] = _eecd(_fgccg)
			}
		}
	}
	return &DSS{Certs: _egggc(_cgbb.Get("\u0043\u0065\u0072t\u0073")), OCSPs: _egggc(_cgbb.Get("\u004f\u0043\u0053P\u0073")), CRLs: _egggc(_cgbb.Get("\u0043\u0052\u004c\u0073")), VRI: _dbge, _fcdgc: _fbce}, nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a CalRGB device.
func (_gfgda *PdfColorspaceCalRGB) GetNumComponents() int { return 3 }

// CustomKeys returns all custom info keys as list.
func (_affdb *PdfInfo) CustomKeys() []string {
	if _affdb._cbfb == nil {
		return nil
	}
	_defdb := make([]string, len(_affdb._cbfb.Keys()))
	for _, _edbb := range _affdb._cbfb.Keys() {
		_defdb = append(_defdb, _edbb.String())
	}
	return _defdb
}

// AddRefChild adds a child reference object.
func (_gfgfa *KDict) AddRefChild(kChild *_eb.PdfObjectDictionary) {
	_gfgfa._eegge = append(_gfgfa._eegge, &KValue{_fbgaa: kChild})
}
func (_gfgdf *PdfReader) loadForms() (*PdfAcroForm, error) {
	if _gfgdf._ebbe.GetCrypter() != nil && !_gfgdf._ebbe.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_aggaa := _gfgdf._bagcfd
	_dgeee := _aggaa.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")
	if _dgeee == nil {
		return nil, nil
	}
	_acacgd, _geaab := _eb.GetIndirect(_dgeee)
	_dgeee = _eb.TraceToDirectObject(_dgeee)
	if _eb.IsNullObject(_dgeee) {
		_ddb.Log.Trace("\u0041\u0063\u0072of\u006f\u0072\u006d\u0020\u0069\u0073\u0020\u0061\u0020n\u0075l\u006c \u006fb\u006a\u0065\u0063\u0074\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029\u000a")
		return nil, nil
	}
	_aeag, _adfaf := _eb.GetDict(_dgeee)
	if !_adfaf {
		_ddb.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0041\u0063\u0072\u006fF\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _dgeee)
		_ddb.Log.Debug("\u0044\u006f\u0065\u0073 n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0073")
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0061\u0063\u0072\u006ff\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _dgeee)
	}
	_ddb.Log.Trace("\u0048\u0061\u0073\u0020\u0041\u0063\u0072\u006f\u0020f\u006f\u0072\u006d\u0073")
	_ddb.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0041\u0063r\u006ff\u006f\u0072\u006d\u0073\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	if !_gfgdf._cfcgdf {
		_fbdead := _gfgdf.traverseObjectData(_aeag)
		if _fbdead != nil {
			_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0074\u0072a\u0076\u0065\u0072\u0073\u0065\u0020\u0041\u0063\u0072\u006fFo\u0072\u006d\u0073 \u0028%\u0073\u0029", _fbdead)
			return nil, _fbdead
		}
	}
	_ecccfd, _bfcgf := _gfgdf.newPdfAcroFormFromDict(_acacgd, _aeag)
	if _bfcgf != nil {
		return nil, _bfcgf
	}
	_ecccfd._cgedg = !_geaab
	return _ecccfd, nil
}

const (
	BorderStyleSolid     BorderStyle = iota
	BorderStyleDashed    BorderStyle = iota
	BorderStyleBeveled   BorderStyle = iota
	BorderStyleInset     BorderStyle = iota
	BorderStyleUnderline BorderStyle = iota
)

// PdfAnnotationProjection represents Projection annotations.
type PdfAnnotationProjection struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
}

func (_agc *PdfReader) newPdfAnnotationRedactFromDict(_gdaa *_eb.PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	_eded := PdfAnnotationRedact{}
	_agge, _ecbd := _agc.newPdfAnnotationMarkupFromDict(_gdaa)
	if _ecbd != nil {
		return nil, _ecbd
	}
	_eded.PdfAnnotationMarkup = _agge
	_eded.QuadPoints = _gdaa.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_eded.IC = _gdaa.Get("\u0049\u0043")
	_eded.RO = _gdaa.Get("\u0052\u004f")
	_eded.OverlayText = _gdaa.Get("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074")
	_eded.Repeat = _gdaa.Get("\u0052\u0065\u0070\u0065\u0061\u0074")
	_eded.DA = _gdaa.Get("\u0044\u0041")
	_eded.Q = _gdaa.Get("\u0051")
	return &_eded, nil
}

type pdfFontType3 struct {
	fontCommon
	_effbd *_eb.PdfIndirectObject

	// These fields are specific to Type 3 fonts.
	CharProcs  _eb.PdfObject
	Encoding   _eb.PdfObject
	FontBBox   _eb.PdfObject
	FontMatrix _eb.PdfObject
	FirstChar  _eb.PdfObject
	LastChar   _eb.PdfObject
	Widths     _eb.PdfObject
	Resources  _eb.PdfObject
	_ccfdb     map[_fc.CharCode]float64
	_bdde      _fc.TextEncoder
}

// PrintClip returns the value of the printClip.
func (_dbecg *ViewerPreferences) PrintClip() PageBoundary { return _dbecg._edef }

type fontCommon struct {
	_agcc  string
	_fgdee string
	_bgge  string
	_geee  _eb.PdfObject
	_bgbg  *_ff.CMap
	_bged  *PdfFontDescriptor
	_babgg int64
}

// SetPdfCreationDate sets the CreationDate attribute of the output PDF.
func SetPdfCreationDate(creationDate _d.Time) {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	_ebffa = creationDate
}
func _ggcf(_cdfbg _eb.PdfObject) (*PdfFunctionType2, error) {
	_fegba := &PdfFunctionType2{}
	var _cbbd *_eb.PdfObjectDictionary
	if _acef, _adege := _cdfbg.(*_eb.PdfIndirectObject); _adege {
		_baccd, _cgded := _acef.PdfObject.(*_eb.PdfObjectDictionary)
		if !_cgded {
			return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fegba._efee = _acef
		_cbbd = _baccd
	} else if _cgfgcg, _cecdb := _cdfbg.(*_eb.PdfObjectDictionary); _cecdb {
		_cbbd = _cgfgcg
	} else {
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ddb.Log.Trace("\u0046U\u004e\u0043\u0032\u003a\u0020\u0025s", _cbbd.String())
	_fbgab, _eeaec := _eb.TraceToDirectObject(_cbbd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_eb.PdfObjectArray)
	if !_eeaec {
		_ddb.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _fbgab.Len() < 0 || _fbgab.Len()%2 != 0 {
		_ddb.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u0069\u006e\u0076al\u0069\u0064")
		return nil, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_fbgf, _acbag := _fbgab.ToFloat64Array()
	if _acbag != nil {
		return nil, _acbag
	}
	_fegba.Domain = _fbgf
	_fbgab, _eeaec = _eb.TraceToDirectObject(_cbbd.Get("\u0052\u0061\u006eg\u0065")).(*_eb.PdfObjectArray)
	if _eeaec {
		if _fbgab.Len() < 0 || _fbgab.Len()%2 != 0 {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_bbec, _gdgea := _fbgab.ToFloat64Array()
		if _gdgea != nil {
			return nil, _gdgea
		}
		_fegba.Range = _bbec
	}
	_fbgab, _eeaec = _eb.TraceToDirectObject(_cbbd.Get("\u0043\u0030")).(*_eb.PdfObjectArray)
	if _eeaec {
		_afeaa, _gdgg := _fbgab.ToFloat64Array()
		if _gdgg != nil {
			return nil, _gdgg
		}
		_fegba.C0 = _afeaa
	}
	_fbgab, _eeaec = _eb.TraceToDirectObject(_cbbd.Get("\u0043\u0031")).(*_eb.PdfObjectArray)
	if _eeaec {
		_fdagd, _afdcf := _fbgab.ToFloat64Array()
		if _afdcf != nil {
			return nil, _afdcf
		}
		_fegba.C1 = _fdagd
	}
	if len(_fegba.C0) != len(_fegba.C1) {
		_ddb.Log.Error("\u0043\u0030\u0020\u0061nd\u0020\u0043\u0031\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069n\u0067")
		return nil, _eb.ErrRangeError
	}
	N, _acbag := _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_cbbd.Get("\u004e")))
	if _acbag != nil {
		_ddb.Log.Error("\u004e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020o\u0072\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u002c\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _cbbd.String())
		return nil, _acbag
	}
	_fegba.N = N
	return _fegba, nil
}

// Register registers (caches) a model to primitive object relationship.
func (_ddcge *modelManager) Register(primitive _eb.PdfObject, model PdfModel) {
	_ddcge._edcg[model] = primitive
	_ddcge._fgffbd[primitive] = model
}
func _abaa(_bffb *_eb.PdfObjectDictionary) {
	_cfcec, _cdaf := _eb.GetArray(_bffb.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
	_decbac, _egff := _eb.GetIntVal(_bffb.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
	_deafg, _cdbab := _eb.GetIntVal(_bffb.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
	if _cdaf && _egff && _cdbab {
		_agaaa := _cfcec.Len()
		if _agaaa != _deafg-_decbac+1 {
			_ddb.Log.Debug("\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076\u002c\u0020\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0076", _agaaa, _deafg)
			_effgb := _eb.PdfObjectInteger(_decbac + _agaaa - 1)
			_bffb.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", &_effgb)
		}
	}
}
func (_cegde *PdfWriter) hasObject(_facac _eb.PdfObject) bool {
	_, _dacec := _cegde._aeeda[_facac]
	return _dacec
}
func _aecd() string {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	return _bgabb
}
func (_cccce *PdfWriter) writeOutputIntents() error {
	if len(_cccce._bbbff) == 0 {
		return nil
	}
	_bbbd := make([]_eb.PdfObject, len(_cccce._bbbff))
	for _ffcea, _afgce := range _cccce._bbbff {
		_accdg := _afgce.ToPdfObject()
		_bbbd[_ffcea] = _eb.MakeIndirectObject(_accdg)
	}
	_bcgga := _eb.MakeIndirectObject(_eb.MakeArray(_bbbd...))
	_cccce._dbffa.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _bcgga)
	if _bbfe := _cccce.addObjects(_bcgga); _bbfe != nil {
		return _bbfe
	}
	return nil
}
func (_ag *PdfReader) newPdfActionGotoEFromDict(_ece *_eb.PdfObjectDictionary) (*PdfActionGoToE, error) {
	_bebg, _cbe := _dba(_ece.Get("\u0046"))
	if _cbe != nil {
		return nil, _cbe
	}
	return &PdfActionGoToE{D: _ece.Get("\u0044"), NewWindow: _ece.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), T: _ece.Get("\u0054"), F: _bebg}, nil
}

// GetAttachedFiles retrieves all the attached files info and content.
func (_cagce *PdfReader) GetAttachedFiles() ([]*EmbeddedFile, error) {
	_aacg := []*EmbeddedFile{}
	_bcba, _fgeda := _cagce.GetNameDictionary()
	if _fgeda != nil {
		return nil, _fgeda
	}
	if _bcba == nil {
		return _aacg, nil
	}
	_fgbeca := _cadd(_bcba)
	if _fgbeca.EmbeddedFiles == nil {
		return nil, nil
	}
	_aaeba := _fgbeca.EmbeddedFiles.Get("\u004e\u0061\u006de\u0073")
	_dffbe, _dfge := _aaeba.(*_eb.PdfObjectArray)
	if !_dfge {
		return nil, _dcf.New("\u0049\u006e\u0076\u0061li\u0064\u0020\u004e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	for _cgddcc := 1; _cgddcc < len(_dffbe.Elements()); _cgddcc += 2 {
		if _cgddcc%2 != 0 {
			_gdcbb := _dffbe.Get(_cgddcc)
			_afadd, _edca := NewPdfFilespecFromObj(_gdcbb)
			if _edca != nil {
				return nil, _edca
			}
			_ebgce, _edca := NewEmbeddedFileFromObject(_afadd.EF)
			if _edca != nil {
				return nil, _edca
			}
			_aeca, _fedbd := _afadd.F.(*_eb.PdfObjectString)
			if _fedbd {
				_ebgce.Name = _aeca.Str()
			}
			_ebgce.Description = _afadd.Desc.WriteString()
			_ebgce.Relationship = RelationshipUnspecified
			if _afadd.AFRelationship != nil {
				switch _afadd.AFRelationship.WriteString() {
				case "\u0053\u006f\u0075\u0072\u0063\u0065":
					_ebgce.Relationship = RelationshipSource
				case "\u0044\u0061\u0074\u0061":
					_ebgce.Relationship = RelationshipData
				case "A\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065":
					_ebgce.Relationship = RelationshipAlternative
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_ebgce.Relationship = RelationshipSupplement
				default:
					_ebgce.Relationship = RelationshipUnspecified
				}
			}
			_aacg = append(_aacg, _ebgce)
		}
	}
	return _aacg, nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_bacbb *PdfSignature) GetContainingPdfObject() _eb.PdfObject { return _bacbb._cddce }
func (_ffga *DSS) add(_ecda *[]*_eb.PdfObjectStream, _dfbfe map[string]*_eb.PdfObjectStream, _afdaee [][]byte) ([]*_eb.PdfObjectStream, error) {
	_dfad := make([]*_eb.PdfObjectStream, 0, len(_afdaee))
	for _, _dgga := range _afdaee {
		_afgba, _cbgba := _bceda(_dgga)
		if _cbgba != nil {
			return nil, _cbgba
		}
		_fdge, _bbbc := _dfbfe[string(_afgba)]
		if !_bbbc {
			_fdge, _cbgba = _eb.MakeStream(_dgga, _eb.NewRawEncoder())
			if _cbgba != nil {
				return nil, _cbgba
			}
			_dfbfe[string(_afgba)] = _fdge
			*_ecda = append(*_ecda, _fdge)
		}
		_dfad = append(_dfad, _fdge)
	}
	return _dfad, nil
}

// ColorToRGB verifies that the input color is an RGB color. Method exists in
// order to satisfy the PdfColorspace interface.
func (_fgbe *PdfColorspaceDeviceRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ccgb, _afbb := color.(*PdfColorDeviceRGB)
	if !_afbb {
		_ddb.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072 \u006e\u006f\u0074\u0020\u0064\u0065\u0076\u0069\u0063\u0065 \u0052\u0047\u0042")
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return _ccgb, nil
}

// ButtonType represents the subtype of a button field, can be one of:
// - Checkbox (ButtonTypeCheckbox)
// - PushButton (ButtonTypePushButton)
// - RadioButton (ButtonTypeRadioButton)
type ButtonType int

// Encoder returns the font's text encoder.
func (_bgbc pdfFontType3) Encoder() _fc.TextEncoder { return _bgbc._bdde }

// GetContainingPdfObject returns the XObject Form's containing object (indirect object).
func (_cffbe *XObjectForm) GetContainingPdfObject() _eb.PdfObject { return _cffbe._afcag }

// Duplicate creates a duplicate page based on the current one and returns it.
func (_fdbb *PdfPage) Duplicate() *PdfPage {
	_egbdc := *_fdbb
	_egbdc._aaagaa = _eb.MakeDict()
	_egbdc._efcff = _eb.MakeIndirectObject(_egbdc._aaagaa)
	_egbdc._eaebc = *_egbdc._aaagaa
	return &_egbdc
}

// NewCompliancePdfReader creates a PdfReader or an input io.ReadSeeker that during reading will scan the files for the
// metadata details. It could be used for the PDF standard implementations like PDF/A or PDF/X.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
func NewCompliancePdfReader(rs _bagf.ReadSeeker) (*CompliancePdfReader, error) {
	const _dbfd = "\u006d\u006f\u0064\u0065l\u003a\u004e\u0065\u0077\u0043\u006f\u006d\u0070\u006c\u0069a\u006ec\u0065\u0050\u0064\u0066\u0052\u0065\u0061d\u0065\u0072"
	_agdcb, _gacdc := _dcgfe(rs, &ReaderOpts{ComplianceMode: true}, false, _dbfd)
	if _gacdc != nil {
		return nil, _gacdc
	}
	return &CompliancePdfReader{PdfReader: _agdcb}, nil
}

// PdfAnnotationHighlight represents Highlight annotations.
// (Section 12.5.6.10).
type PdfAnnotationHighlight struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _eb.PdfObject
}

// SetNameDictionary sets the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_fbefa *PdfWriter) SetNameDictionary(names _eb.PdfObject) error {
	if names == nil {
		return nil
	}
	_fbefa._agdbc = _cadd(names)
	_ddb.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u004e\u0061\u006d\u0065\u0073\u002e\u002e\u002e")
	_fbefa._dbffa.Set("\u004e\u0061\u006de\u0073", names)
	return _fbefa.addObjects(names)
}

// GetCatalogStructTreeRoot gets the catalog StructTreeRoot object.
func (_addce *PdfReader) GetCatalogStructTreeRoot() (_eb.PdfObject, bool) {
	_bcgdf := _eb.ResolveReference(_addce._bagcfd.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	if _bcgdf == nil {
		return nil, false
	}
	if !_addce._cfcgdf {
		_gddfe := _addce.traverseObjectData(_bcgdf)
		if _gddfe != nil {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065\u0064\u0020t\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _gddfe)
			return nil, false
		}
	}
	return _bcgdf, true
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_fcdbf *PdfOutlineItem) ToPdfObject() _eb.PdfObject {
	_fbbfd := _fcdbf._efbfb
	_fgcdag := _fbbfd.PdfObject.(*_eb.PdfObjectDictionary)
	_fgcdag.Set("\u0054\u0069\u0074l\u0065", _fcdbf.Title)
	if _fcdbf.A != nil {
		_fgcdag.Set("\u0041", _fcdbf.A)
	}
	if _eeede := _fgcdag.Get("\u0053\u0045"); _eeede != nil {
		_fgcdag.Remove("\u0053\u0045")
	}
	if _fcdbf.C != nil {
		_fgcdag.Set("\u0043", _fcdbf.C)
	}
	if _fcdbf.Dest != nil {
		_fgcdag.Set("\u0044\u0065\u0073\u0074", _fcdbf.Dest)
	}
	if _fcdbf.F != nil {
		_fgcdag.Set("\u0046", _fcdbf.F)
	}
	if _fcdbf.Count != nil {
		_fgcdag.Set("\u0043\u006f\u0075n\u0074", _eb.MakeInteger(*_fcdbf.Count))
	}
	if _fcdbf.Next != nil {
		_fgcdag.Set("\u004e\u0065\u0078\u0074", _fcdbf.Next.ToPdfObject())
	}
	if _fcdbf.First != nil {
		_fgcdag.Set("\u0046\u0069\u0072s\u0074", _fcdbf.First.ToPdfObject())
	}
	if _fcdbf.Prev != nil {
		_fgcdag.Set("\u0050\u0072\u0065\u0076", _fcdbf.Prev.GetContext().GetContainingPdfObject())
	}
	if _fcdbf.Last != nil {
		_fgcdag.Set("\u004c\u0061\u0073\u0074", _fcdbf.Last.GetContext().GetContainingPdfObject())
	}
	if _fcdbf.Parent != nil {
		_fgcdag.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _fcdbf.Parent.GetContext().GetContainingPdfObject())
	}
	return _fbbfd
}

// ToInteger convert to an integer format.
func (_bfaa *PdfColorCalGray) ToInteger(bits int) uint32 {
	_bae := _gg.Pow(2, float64(bits)) - 1
	return uint32(_bae * _bfaa.Val())
}
func (_aggc *PdfColorspaceSpecialSeparation) String() string {
	return "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"
}

// ToPdfObject implements interface PdfModel.
func (_fagcd *PdfTransformParamsDocMDP) ToPdfObject() _eb.PdfObject {
	_bbebg := _eb.MakeDict()
	_bbebg.SetIfNotNil("\u0054\u0079\u0070\u0065", _fagcd.Type)
	_bbebg.SetIfNotNil("\u0056", _fagcd.V)
	_bbebg.SetIfNotNil("\u0050", _fagcd.P)
	return _bbebg
}

// PdfAnnotationMarkup represents additional fields for mark-up annotations.
// (Section 12.5.6.2 p. 399).
type PdfAnnotationMarkup struct {
	T            _eb.PdfObject
	Popup        *PdfAnnotationPopup
	CA           _eb.PdfObject
	RC           _eb.PdfObject
	CreationDate _eb.PdfObject
	IRT          _eb.PdfObject
	Subj         _eb.PdfObject
	RT           _eb.PdfObject
	IT           _eb.PdfObject
	ExData       _eb.PdfObject
}

func _eaff(_cdffg _eb.PdfObject) (*fontFile, error) {
	_ddb.Log.Trace("\u006e\u0065\u0077\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0046\u0072\u006f\u006dP\u0064f\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u006f\u0062\u006a\u003d\u0025\u0073", _cdffg)
	_ddgaa := &fontFile{}
	_cdffg = _eb.TraceToDirectObject(_cdffg)
	_gcee, _ecfbd := _cdffg.(*_eb.PdfObjectStream)
	if !_ecfbd {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020F\u006f\u006et\u0046\u0069\u006c\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0028\u0025\u0054\u0029", _cdffg)
		return nil, _eb.ErrTypeError
	}
	_edbd := _gcee.PdfObjectDictionary
	_gggc, _fadfg := _eb.DecodeStream(_gcee)
	if _fadfg != nil {
		return nil, _fadfg
	}
	_ccaaa, _ecfbd := _eb.GetNameVal(_edbd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_ecfbd {
		_ddgaa._fbca = _ccaaa
		if _ccaaa == "\u0054\u0079\u0070\u0065\u0031\u0043" {
			_ddb.Log.Debug("T\u0079\u0070\u0065\u0031\u0043\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006e\u006f\u0074 s\u0075\u0070\u0070o\u0072t\u0065\u0064")
			return nil, ErrType1CFontNotSupported
		}
	}
	_fgcf, _ := _eb.GetIntVal(_edbd.Get("\u004ce\u006e\u0067\u0074\u0068\u0031"))
	_gfabe, _ := _eb.GetIntVal(_edbd.Get("\u004ce\u006e\u0067\u0074\u0068\u0032"))
	if _fgcf > len(_gggc) {
		_fgcf = len(_gggc)
	}
	if _fgcf+_gfabe > len(_gggc) {
		_gfabe = len(_gggc) - _fgcf
	}
	_fbef := _gggc[:_fgcf]
	var _fcgba []byte
	if _gfabe > 0 {
		_fcgba = _gggc[_fgcf : _fgcf+_gfabe]
	}
	if _fgcf > 0 && _gfabe > 0 {
		_ebaeb := _ddgaa.loadFromSegments(_fbef, _fcgba)
		if _ebaeb != nil {
			return nil, _ebaeb
		}
	}
	return _ddgaa, nil
}
func (_aagg *PdfFilespec) getDict() *_eb.PdfObjectDictionary {
	if _afgga, _gaced := _aagg._cbefe.(*_eb.PdfIndirectObject); _gaced {
		_defe, _gbabe := _afgga.PdfObject.(*_eb.PdfObjectDictionary)
		if !_gbabe {
			return nil
		}
		return _defe
	} else if _affbf, _gdeg := _aagg._cbefe.(*_eb.PdfObjectDictionary); _gdeg {
		return _affbf
	} else {
		_ddb.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074 \u0074\u0079p\u0065\u0020(\u0025T\u0029", _aagg._cbefe)
		return nil
	}
}

// ToPdfObject implements interface PdfModel.
func (_ebgg *PdfAnnotationProjection) ToPdfObject() _eb.PdfObject {
	_ebgg.PdfAnnotation.ToPdfObject()
	_bcc := _ebgg._ggf
	_eegd := _bcc.PdfObject.(*_eb.PdfObjectDictionary)
	_ebgg.PdfAnnotationMarkup.appendToPdfDictionary(_eegd)
	return _bcc
}

// SetNamedDestinations sets the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_aegcg *PdfWriter) SetNamedDestinations(dests _eb.PdfObject) error {
	if dests == nil {
		return nil
	}
	_ddb.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0044\u0065\u0073\u0074\u0073\u002e\u002e\u002e")
	_aegcg._dbffa.Set("\u0044\u0065\u0073t\u0073", dests)
	return _aegcg.addObjects(dests)
}

// PdfAnnotationPopup represents Popup annotations.
// (Section 12.5.6.14).
type PdfAnnotationPopup struct {
	*PdfAnnotation
	Parent _eb.PdfObject
	Open   _eb.PdfObject
}

// GetFontDescriptor returns the font descriptor for `font`.
func (_fedgd PdfFont) GetFontDescriptor() (*PdfFontDescriptor, error) {
	return _fedgd._fdaa.getFontDescriptor(), nil
}

// CharcodesToUnicode converts the character codes `charcodes` to a slice of runes.
// How it works:
//  1. Use the ToUnicode CMap if there is one.
//  2. Use the underlying font's encoding.
func (_beea *PdfFont) CharcodesToUnicode(charcodes []_fc.CharCode) []rune {
	_ddcbe, _, _ := _beea.CharcodesToUnicodeWithStats(charcodes)
	return _ddcbe
}
func (_cafbda *Names) addEmbeddedFile(_dgedb *EmbeddedFile) error {
	if _cafbda.EmbeddedFiles == nil {
		_cafbda.EmbeddedFiles = _eb.MakeDict()
		_cafbda.EmbeddedFiles.Set("\u004e\u0061\u006de\u0073", _eb.MakeArray())
	}
	_gegd := NewPdfFileSpecFromEmbeddedFile(_dgedb)
	_bcaf := _cafbda.EmbeddedFiles.Get("\u004e\u0061\u006de\u0073")
	_ecdgfd, _agfeg := _bcaf.(*_eb.PdfObjectArray)
	if !_agfeg {
		return _dcf.New("\u0049\u006e\u0076\u0061li\u0064\u0020\u004e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	type FileSpecMap struct {
		_gdcbe string
		_gbad  *PdfFilespec
	}
	_bbfc := []FileSpecMap{}
	for _gdeaeb := 0; _gdeaeb < len(_ecdgfd.Elements()); _gdeaeb += 2 {
		if _gdeaeb%2 == 0 {
			_gfbd := _ecdgfd.Get(_gdeaeb)
			if _gfbd != nil {
				_gbeee := _gfbd.(*_eb.PdfObjectString)
				_bfbb := _ecdgfd.Get(_gdeaeb + 1)
				_deefa, _ffdaec := NewPdfFilespecFromObj(_bfbb)
				if _ffdaec != nil {
					return _ffdaec
				}
				_bbfc = append(_bbfc, FileSpecMap{_gdcbe: _gbeee.String(), _gbad: _deefa})
			}
		}
	}
	_bbfc = append(_bbfc, FileSpecMap{_gdcbe: _dgedb.Name, _gbad: _gegd})
	_ba.Slice(_bbfc, func(_ffgad, _gegg int) bool { return _bbfc[_ffgad]._gdcbe < _bbfc[_gegg]._gdcbe })
	_ecdgfd = _eb.MakeArray()
	for _, _bafe := range _bbfc {
		_ecdgfd.Append(_eb.MakeString(_bafe._gdcbe))
		_ecdgfd.Append(_bafe._gbad.ToPdfObject())
	}
	_cafbda.EmbeddedFiles.Set("\u004e\u0061\u006de\u0073", _ecdgfd)
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_aec *PdfActionSubmitForm) ToPdfObject() _eb.PdfObject {
	_aec.PdfAction.ToPdfObject()
	_ac := _aec._dee
	_efc := _ac.PdfObject.(*_eb.PdfObjectDictionary)
	_efc.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeSubmitForm)))
	if _aec.F != nil {
		_efc.Set("\u0046", _aec.F.ToPdfObject())
	}
	_efc.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _aec.Fields)
	_efc.SetIfNotNil("\u0046\u006c\u0061g\u0073", _aec.Flags)
	return _ac
}

// PdfColorCalGray represents a CalGray colorspace.
type PdfColorCalGray float64

func _gbdf(_efdb *_eb.PdfObjectDictionary) (*PdfFieldButton, error) {
	_aaccf := &PdfFieldButton{}
	_aaccf.PdfField = NewPdfField()
	_aaccf.PdfField.SetContext(_aaccf)
	_aaccf.Opt, _ = _eb.GetArray(_efdb.Get("\u004f\u0070\u0074"))
	_decba := NewPdfAnnotationWidget()
	_decba.A, _ = _eb.GetDict(_efdb.Get("\u0041"))
	_decba.AP, _ = _eb.GetDict(_efdb.Get("\u0041\u0050"))
	_decba.SetContext(_aaccf)
	_aaccf.PdfField.Annotations = append(_aaccf.PdfField.Annotations, _decba)
	return _aaccf, nil
}

// PdfColorspaceSpecialSeparation is a Separation colorspace.
// At the moment the colour space is set to a Separation space, the conforming reader shall determine whether the
// device has an available colorant (e.g. dye) corresponding to the name of the requested space. If so, the conforming
// reader shall ignore the alternateSpace and tintTransform parameters; subsequent painting operations within the
// space shall apply the designated colorant directly, according to the tint values supplied.
//
// Format: [/Separation name alternateSpace tintTransform]
type PdfColorspaceSpecialSeparation struct {
	ColorantName   *_eb.PdfObjectName
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	_degb          *_eb.PdfIndirectObject
}

// Field returns the parent form field of the widget annotation, if one exists.
// NOTE: the method returns nil if the parent form field has not been parsed.
func (_bggf *PdfAnnotationWidget) Field() *PdfField { return _bggf._bca }

// GetNumComponents returns the number of color components (3 for CalRGB).
func (_gdfb *PdfColorCalRGB) GetNumComponents() int { return 3 }

// GetCatalogLanguage gets catalog Language object.
func (_fcgcf *PdfReader) GetCatalogLanguage() (_eb.PdfObject, bool) {
	if _fcgcf._bagcfd == nil {
		return nil, false
	}
	_eaecc := _fcgcf._bagcfd.Get("\u004c\u0061\u006e\u0067")
	return _eaecc, _eaecc != nil
}

// PdfActionImportData represents a importData action.
type PdfActionImportData struct {
	*PdfAction
	F *PdfFilespec
}

// UpdatePage updates the `page` in the new revision if it has changed.
func (_gccef *PdfAppender) UpdatePage(page *PdfPage) {
	_gccef.updateObjectsDeep(page.ToPdfObject(), nil)
}

// String returns a string describing the font descriptor.
func (_eaagc *PdfFontDescriptor) String() string {
	var _cdee []string
	if _eaagc.FontName != nil {
		_cdee = append(_cdee, _eaagc.FontName.String())
	}
	if _eaagc.FontFamily != nil {
		_cdee = append(_cdee, _eaagc.FontFamily.String())
	}
	if _eaagc.fontFile != nil {
		_cdee = append(_cdee, _eaagc.fontFile.String())
	}
	if _eaagc._bebgd != nil {
		_cdee = append(_cdee, _eaagc._bebgd.String())
	}
	_cdee = append(_cdee, _e.Sprintf("\u0046\u006f\u006et\u0046\u0069\u006c\u0065\u0033\u003d\u0025\u0074", _eaagc.FontFile3 != nil))
	return _e.Sprintf("\u0046\u004f\u004e\u0054_D\u0045\u0053\u0043\u0052\u0049\u0050\u0054\u004f\u0052\u007b\u0025\u0073\u007d", _cc.Join(_cdee, "\u002c\u0020"))
}

// GetSubFilter returns SubFilter value or empty string.
func (_bcfe *pdfSignDictionary) GetSubFilter() string {
	_fbdda := _bcfe.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r")
	if _fbdda == nil {
		return ""
	}
	if _aeade, _ccdag := _eb.GetNameVal(_fbdda); _ccdag {
		return _aeade
	}
	return ""
}

// PdfColorCalRGB represents a color in the Colorimetric CIE RGB colorspace.
// A, B, C components
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorCalRGB [3]float64

// ToPdfObject implements interface PdfModel.
func (_bec *PdfActionJavaScript) ToPdfObject() _eb.PdfObject {
	_bec.PdfAction.ToPdfObject()
	_cef := _bec._dee
	_bee := _cef.PdfObject.(*_eb.PdfObjectDictionary)
	_bee.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeJavaScript)))
	_bee.SetIfNotNil("\u004a\u0053", _bec.JS)
	return _cef
}

// PdfActionGoToE represents a GoToE action.
type PdfActionGoToE struct {
	*PdfAction
	F         *PdfFilespec
	D         _eb.PdfObject
	NewWindow _eb.PdfObject
	T         _eb.PdfObject
}

// ImageToRGB returns an error since an image cannot be defined in a pattern colorspace.
func (_fdeb *PdfColorspaceSpecialPattern) ImageToRGB(img Image) (Image, error) {
	_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0061\u0074\u0074\u0065\u0072n \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	return img, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0066\u006f\u0072\u0020\u0069m\u0061\u0067\u0065\u0020\u0028p\u0061\u0074t\u0065\u0072\u006e\u0029")
}

// GetContainingPdfObject implements interface PdfModel.
func (_accb *PdfSignatureReference) GetContainingPdfObject() _eb.PdfObject { return _accb._abdcb }

// SetCatalogLanguage sets the catalog language.
func (_dgcdd *PdfWriter) SetCatalogLanguage(lang _eb.PdfObject) error {
	if lang == nil {
		_dgcdd._dbffa.Remove("\u004c\u0061\u006e\u0067")
		return nil
	}
	_dgcdd.addObject(lang)
	_dgcdd._dbffa.Set("\u004c\u0061\u006e\u0067", lang)
	return nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_cffe *PdfShadingType7) ToPdfObject() _eb.PdfObject {
	_cffe.PdfShading.ToPdfObject()
	_edgc, _gfacd := _cffe.getShadingDict()
	if _gfacd != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cffe.BitsPerCoordinate != nil {
		_edgc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _cffe.BitsPerCoordinate)
	}
	if _cffe.BitsPerComponent != nil {
		_edgc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cffe.BitsPerComponent)
	}
	if _cffe.BitsPerFlag != nil {
		_edgc.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _cffe.BitsPerFlag)
	}
	if _cffe.Decode != nil {
		_edgc.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _cffe.Decode)
	}
	if _cffe.Function != nil {
		if len(_cffe.Function) == 1 {
			_edgc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cffe.Function[0].ToPdfObject())
		} else {
			_fcfbe := _eb.MakeArray()
			for _, _adacga := range _cffe.Function {
				_fcfbe.Append(_adacga.ToPdfObject())
			}
			_edgc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _fcfbe)
		}
	}
	return _cffe._cefaa
}

// PartialName returns the partial name of the field.
func (_bdff *PdfField) PartialName() string {
	_baee := ""
	if _bdff.T != nil {
		_baee = _bdff.T.Decoded()
	} else {
		_ddb.Log.Debug("\u0046\u0069el\u0064\u0020\u006di\u0073\u0073\u0069\u006eg T\u0020fi\u0065\u006c\u0064\u0020\u0028\u0069\u006eco\u006d\u0070\u0061\u0074\u0069\u0062\u006ce\u0029")
	}
	return _baee
}
func _afcc(_bgcd *_eb.PdfObjectDictionary) bool {
	for _, _geca := range _bgcd.Keys() {
		if _, _cefe := _gdba[_geca.String()]; _cefe {
			return true
		}
	}
	return false
}

// NewPdfAnnotation returns an initialized generic PDF annotation model.
func NewPdfAnnotation() *PdfAnnotation {
	_effc := &PdfAnnotation{}
	_effc._ggf = _eb.MakeIndirectObject(_eb.MakeDict())
	return _effc
}

// Reset sets the multi font encoder to its initial state.
func (_bega *MultipleFontEncoder) Reset() { _bega.CurrentFont = _bega._fcgbc[0] }

// String returns a string representation of what flags are set.
func (_cfceba FieldFlag) String() string {
	_ebec := ""
	if _cfceba == FieldFlagClear {
		_ebec = "\u0043\u006c\u0065a\u0072"
		return _ebec
	}
	if _cfceba&FieldFlagReadOnly > 0 {
		_ebec += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _cfceba&FieldFlagRequired > 0 {
		_ebec += "\u007cR\u0065\u0071\u0075\u0069\u0072\u0065d"
	}
	if _cfceba&FieldFlagNoExport > 0 {
		_ebec += "\u007cN\u006f\u0045\u0078\u0070\u006f\u0072t"
	}
	if _cfceba&FieldFlagNoToggleToOff > 0 {
		_ebec += "\u007c\u004e\u006f\u0054\u006f\u0067\u0067\u006c\u0065T\u006f\u004f\u0066\u0066"
	}
	if _cfceba&FieldFlagRadio > 0 {
		_ebec += "\u007c\u0052\u0061\u0064\u0069\u006f"
	}
	if _cfceba&FieldFlagPushbutton > 0 {
		_ebec += "|\u0050\u0075\u0073\u0068\u0062\u0075\u0074\u0074\u006f\u006e"
	}
	if _cfceba&FieldFlagRadiosInUnision > 0 {
		_ebec += "\u007c\u0052a\u0064\u0069\u006fs\u0049\u006e\u0055\u006e\u0069\u0073\u0069\u006f\u006e"
	}
	if _cfceba&FieldFlagMultiline > 0 {
		_ebec += "\u007c\u004d\u0075\u006c\u0074\u0069\u006c\u0069\u006e\u0065"
	}
	if _cfceba&FieldFlagPassword > 0 {
		_ebec += "\u007cP\u0061\u0073\u0073\u0077\u006f\u0072d"
	}
	if _cfceba&FieldFlagFileSelect > 0 {
		_ebec += "|\u0046\u0069\u006c\u0065\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _cfceba&FieldFlagDoNotScroll > 0 {
		_ebec += "\u007c\u0044\u006fN\u006f\u0074\u0053\u0063\u0072\u006f\u006c\u006c"
	}
	if _cfceba&FieldFlagComb > 0 {
		_ebec += "\u007c\u0043\u006fm\u0062"
	}
	if _cfceba&FieldFlagRichText > 0 {
		_ebec += "\u007cR\u0069\u0063\u0068\u0054\u0065\u0078t"
	}
	if _cfceba&FieldFlagDoNotSpellCheck > 0 {
		_ebec += "\u007c\u0044o\u004e\u006f\u0074S\u0070\u0065\u006c\u006c\u0043\u0068\u0065\u0063\u006b"
	}
	if _cfceba&FieldFlagCombo > 0 {
		_ebec += "\u007c\u0043\u006f\u006d\u0062\u006f"
	}
	if _cfceba&FieldFlagEdit > 0 {
		_ebec += "\u007c\u0045\u0064i\u0074"
	}
	if _cfceba&FieldFlagSort > 0 {
		_ebec += "\u007c\u0053\u006fr\u0074"
	}
	if _cfceba&FieldFlagMultiSelect > 0 {
		_ebec += "\u007c\u004d\u0075l\u0074\u0069\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _cfceba&FieldFlagCommitOnSelChange > 0 {
		_ebec += "\u007cC\u006fm\u006d\u0069\u0074\u004f\u006eS\u0065\u006cC\u0068\u0061\u006e\u0067\u0065"
	}
	return _cc.Trim(_ebec, "\u007c")
}

// ToPdfObject converts the font to a PDF representation.
func (_abdf *pdfFontType3) ToPdfObject() _eb.PdfObject {
	if _abdf._effbd == nil {
		_abdf._effbd = &_eb.PdfIndirectObject{}
	}
	_dbef := _abdf.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0033")
	_abdf._effbd.PdfObject = _dbef
	if _abdf.FirstChar != nil {
		_dbef.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _abdf.FirstChar)
	}
	if _abdf.LastChar != nil {
		_dbef.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _abdf.LastChar)
	}
	if _abdf.Widths != nil {
		_dbef.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _abdf.Widths)
	}
	if _abdf.Encoding != nil {
		_dbef.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _abdf.Encoding)
	} else if _abdf._bdde != nil {
		_efgea := _abdf._bdde.ToPdfObject()
		if _efgea != nil {
			_dbef.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _efgea)
		}
	}
	if _abdf.FontBBox != nil {
		_dbef.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _abdf.FontBBox)
	}
	if _abdf.FontMatrix != nil {
		_dbef.Set("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0069\u0072\u0078", _abdf.FontMatrix)
	}
	if _abdf.CharProcs != nil {
		_dbef.Set("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s", _abdf.CharProcs)
	}
	if _abdf.Resources != nil {
		_dbef.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _abdf.Resources)
	}
	return _abdf._effbd
}

// B returns the value of the B component of the color.
func (_aadcg *PdfColorCalRGB) B() float64 { return _aadcg[1] }

var (
	_ffee  = _ab.MustCompile("\u005cd\u002b\u0020\u0064\u0069c\u0074\u005c\u0073\u002b\u0028d\u0075p\u005cs\u002b\u0029\u003f\u0062\u0065\u0067\u0069n")
	_agffe = _ab.MustCompile("\u005e\u005cs\u002a\u002f\u0028\u005c\u0053\u002b\u003f\u0029\u005c\u0073\u002b\u0028\u002e\u002b\u003f\u0029\u005c\u0073\u002b\u0064\u0065\u0066\\s\u002a\u0024")
	_cccge = _ab.MustCompile("\u005e\u005c\u0073*\u0064\u0075\u0070\u005c\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002a\u002f\u0028\u005c\u0077\u002b\u003f\u0029\u0028\u003f\u003a\u005c\u002e\u005c\u0064\u002b)\u003f\u005c\u0073\u002b\u0070\u0075\u0074\u0024")
	_ccbg  = "\u002f\u0045\u006e\u0063od\u0069\u006e\u0067\u0020\u0032\u0035\u0036\u0020\u0061\u0072\u0072\u0061\u0079"
	_ceab  = "\u0072\u0065\u0061d\u006f\u006e\u006c\u0079\u0020\u0064\u0065\u0066"
	_dggg  = "\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0066\u0069\u006c\u0065\u0020e\u0065\u0078\u0065\u0063"
)

func (_cge *PdfReader) loadAction(_ddcd _eb.PdfObject) (*PdfAction, error) {
	if _cbgc, _gebb := _eb.GetIndirect(_ddcd); _gebb {
		_cgdg, _gedg := _cge.newPdfActionFromIndirectObject(_cbgc)
		if _gedg != nil {
			return nil, _gedg
		}
		return _cgdg, nil
	} else if !_eb.IsNullObject(_ddcd) {
		return nil, _dcf.New("\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, nil
}

// IsColored specifies if the pattern is colored.
func (_fgfe *PdfTilingPattern) IsColored() bool {
	if _fgfe.PaintType != nil && *_fgfe.PaintType == 1 {
		return true
	}
	return false
}
func (_daaba *PdfWriter) writeAcroFormFields() error {
	if _daaba._dcbec == nil {
		return nil
	}
	_ddb.Log.Trace("\u0057r\u0069t\u0069\u006e\u0067\u0020\u0061c\u0072\u006f \u0066\u006f\u0072\u006d\u0073")
	_cgda := _daaba._dcbec.ToPdfObject()
	_ddb.Log.Trace("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _cgda)
	_daaba._dbffa.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _cgda)
	_baac := _daaba.addObjects(_cgda)
	if _baac != nil {
		return _baac
	}
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_gacc *PdfAnnotationLine) ToPdfObject() _eb.PdfObject {
	_gacc.PdfAnnotation.ToPdfObject()
	_fca := _gacc._ggf
	_cda := _fca.PdfObject.(*_eb.PdfObjectDictionary)
	_gacc.PdfAnnotationMarkup.appendToPdfDictionary(_cda)
	_cda.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u004c\u0069\u006e\u0065"))
	_cda.SetIfNotNil("\u004c", _gacc.L)
	_cda.SetIfNotNil("\u0042\u0053", _gacc.BS)
	_cda.SetIfNotNil("\u004c\u0045", _gacc.LE)
	_cda.SetIfNotNil("\u0049\u0043", _gacc.IC)
	_cda.SetIfNotNil("\u004c\u004c", _gacc.LL)
	_cda.SetIfNotNil("\u004c\u004c\u0045", _gacc.LLE)
	_cda.SetIfNotNil("\u0043\u0061\u0070", _gacc.Cap)
	_cda.SetIfNotNil("\u0049\u0054", _gacc.IT)
	_cda.SetIfNotNil("\u004c\u004c\u004f", _gacc.LLO)
	_cda.SetIfNotNil("\u0043\u0050", _gacc.CP)
	_cda.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _gacc.Measure)
	_cda.SetIfNotNil("\u0043\u004f", _gacc.CO)
	return _fca
}

// PageCallback callback function used in page loading
// that could be used to modify the page content.
//
// Deprecated: will be removed in v4. Use PageProcessCallback instead.
type PageCallback func(_eege int, _gcda *PdfPage)

// SetColorspaceByName adds the provided colorspace to the page resources.
func (_bgfbgb *PdfPageResources) SetColorspaceByName(keyName _eb.PdfObjectName, cs PdfColorspace) error {
	_eedbb, _fceec := _bgfbgb.GetColorspaces()
	if _fceec != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _fceec)
		return _fceec
	}
	if _eedbb == nil {
		_eedbb = NewPdfPageResourcesColorspaces()
		_bgfbgb.SetColorSpace(_eedbb)
	}
	_eedbb.Set(keyName, cs)
	return nil
}

// VRI represents a Validation-Related Information dictionary.
// The VRI dictionary contains validation data in the form of
// certificates, OCSP and CRL information, for a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type VRI struct {
	Cert []*_eb.PdfObjectStream
	OCSP []*_eb.PdfObjectStream
	CRL  []*_eb.PdfObjectStream
	TU   *_eb.PdfObjectString
	TS   *_eb.PdfObjectString
}

// ToPdfObject converts colorspace to a PDF object. [/Indexed base hival lookup]
func (_abcd *PdfColorspaceSpecialIndexed) ToPdfObject() _eb.PdfObject {
	_bgdf := _eb.MakeArray(_eb.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"))
	_bgdf.Append(_abcd.Base.ToPdfObject())
	_bgdf.Append(_eb.MakeInteger(int64(_abcd.HiVal)))
	_bgdf.Append(_abcd.Lookup)
	if _abcd._daac != nil {
		_abcd._daac.PdfObject = _bgdf
		return _abcd._daac
	}
	return _bgdf
}

// NewPdfAnnotationUnderline returns a new text underline annotation.
func NewPdfAnnotationUnderline() *PdfAnnotationUnderline {
	_ceaf := NewPdfAnnotation()
	_febb := &PdfAnnotationUnderline{}
	_febb.PdfAnnotation = _ceaf
	_febb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ceaf.SetContext(_febb)
	return _febb
}
func (_fee *PdfReader) newPdfActionURIFromDict(_ced *_eb.PdfObjectDictionary) (*PdfActionURI, error) {
	return &PdfActionURI{URI: _ced.Get("\u0055\u0052\u0049"), IsMap: _ced.Get("\u0049\u0073\u004da\u0070")}, nil
}

// GetContainingPdfObject returns the container of the outline item (indirect object).
func (_gbbdf *PdfOutlineItem) GetContainingPdfObject() _eb.PdfObject { return _gbbdf._efbfb }
func _fcfc() *modelManager {
	_fagc := modelManager{}
	_fagc._edcg = map[PdfModel]_eb.PdfObject{}
	_fagc._fgffbd = map[_eb.PdfObject]PdfModel{}
	return &_fagc
}
func _gcbfd(_ccabf *XObjectForm) (*PdfRectangle, bool, error) {
	if _cbac, _ecgce := _ccabf.BBox.(*_eb.PdfObjectArray); _ecgce {
		_deeb, _aeef := NewPdfRectangle(*_cbac)
		if _aeef != nil {
			return nil, false, _aeef
		}
		if _gbdfe, _bbgd := _ccabf.Matrix.(*_eb.PdfObjectArray); _bbgd {
			_eefaa, _gege := _gbdfe.ToFloat64Array()
			if _gege != nil {
				return nil, false, _gege
			}
			_bbfa := _ffg.IdentityMatrix()
			if len(_eefaa) == 6 {
				_bbfa = _ffg.NewMatrix(_eefaa[0], _eefaa[1], _eefaa[2], _eefaa[3], _eefaa[4], _eefaa[5])
			}
			_deeb.Transform(_bbfa)
			return _deeb, true, nil
		}
		return _deeb, false, nil
	}
	return nil, false, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063e\u0020\u0042\u0042\u006f\u0078\u0020\u0074y\u0070\u0065")
}

// ToGray returns a PdfColorDeviceGray color based on the current RGB color.
func (_fffb *PdfColorDeviceRGB) ToGray() *PdfColorDeviceGray {
	_bcda := 0.3*_fffb.R() + 0.59*_fffb.G() + 0.11*_fffb.B()
	_bcda = _gg.Min(_gg.Max(_bcda, 0.0), 1.0)
	return NewPdfColorDeviceGray(_bcda)
}

// NewPdfFontFromPdfObject loads a PdfFont from the dictionary `fontObj`.  If there is a problem an
// error is returned.
func NewPdfFontFromPdfObject(fontObj _eb.PdfObject) (*PdfFont, error) { return _fdce(fontObj, true) }

// PdfColorDeviceGray represents a grayscale color value that shall be represented by a single number in the
// range 0.0 to 1.0 where 0.0 corresponds to black and 1.0 to white.
type PdfColorDeviceGray float64

// AddCRLs adds CRLs to DSS.
func (_agbb *DSS) AddCRLs(crls [][]byte) ([]*_eb.PdfObjectStream, error) {
	return _agbb.add(&_agbb.CRLs, _agbb._adcdd, crls)
}

// ToPdfObject implements interface PdfModel.
func (_eeba *PdfActionSetOCGState) ToPdfObject() _eb.PdfObject {
	_eeba.PdfAction.ToPdfObject()
	_cbc := _eeba._dee
	_bfa := _cbc.PdfObject.(*_eb.PdfObjectDictionary)
	_bfa.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeSetOCGState)))
	_bfa.SetIfNotNil("\u0053\u0074\u0061t\u0065", _eeba.State)
	_bfa.SetIfNotNil("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042", _eeba.PreserveRB)
	return _cbc
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_cdcd *PdfFunctionType0) Evaluate(x []float64) ([]float64, error) {
	if len(x) != _cdcd.NumInputs {
		_ddb.Log.Error("\u004eu\u006d\u0062e\u0072\u0020\u006f\u0066 \u0069\u006e\u0070u\u0074\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061tc\u0068\u0069\u006eg\u0020\u0077h\u0061\u0074\u0020\u0069\u0073\u0020n\u0065\u0065d\u0065\u0064")
		return nil, _dcf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _cdcd._bgaaab == nil {
		_efgd := _cdcd.processSamples()
		if _efgd != nil {
			return nil, _efgd
		}
	}
	_aadfd := _cdcd.Encode
	if _aadfd == nil {
		_aadfd = []float64{}
		for _fecff := 0; _fecff < len(_cdcd.Size); _fecff++ {
			_aadfd = append(_aadfd, 0)
			_aadfd = append(_aadfd, float64(_cdcd.Size[_fecff]-1))
		}
	}
	_cbdeg := _cdcd.Decode
	if _cbdeg == nil {
		_cbdeg = _cdcd.Range
	}
	_ebdbf := make([]int, len(x))
	for _cgfgc := 0; _cgfgc < len(x); _cgfgc++ {
		_bcgd := x[_cgfgc]
		_ddadd := _gg.Min(_gg.Max(_bcgd, _cdcd.Domain[2*_cgfgc]), _cdcd.Domain[2*_cgfgc+1])
		_geegd := _df.LinearInterpolate(_ddadd, _cdcd.Domain[2*_cgfgc], _cdcd.Domain[2*_cgfgc+1], _aadfd[2*_cgfgc], _aadfd[2*_cgfgc+1])
		_eefb := _gg.Min(_gg.Max(_geegd, 0), float64(_cdcd.Size[_cgfgc]-1))
		_abccg := int(_gg.Floor(_eefb + 0.5))
		if _abccg < 0 {
			_abccg = 0
		} else if _abccg > _cdcd.Size[_cgfgc] {
			_abccg = _cdcd.Size[_cgfgc] - 1
		}
		_ebdbf[_cgfgc] = _abccg
	}
	_gdcaf := _ebdbf[0]
	for _fegg := 1; _fegg < _cdcd.NumInputs; _fegg++ {
		_cgffb := _ebdbf[_fegg]
		for _afacb := 0; _afacb < _fegg; _afacb++ {
			_cgffb *= _cdcd.Size[_afacb]
		}
		_gdcaf += _cgffb
	}
	_gdcaf *= _cdcd.NumOutputs
	var _bacd []float64
	for _ebfg := 0; _ebfg < _cdcd.NumOutputs; _ebfg++ {
		_ccdgc := _gdcaf + _ebfg
		if _ccdgc >= len(_cdcd._bgaaab) {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u006e\u006ft\u0020\u0065\u006eo\u0075\u0067\u0068\u0020\u0069\u006ep\u0075\u0074\u0020sa\u006dp\u006c\u0065\u0073\u0020\u0074\u006f\u0020d\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0076\u0061lu\u0065\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			continue
		}
		_gfcfe := _cdcd._bgaaab[_ccdgc]
		_cgeec := _df.LinearInterpolate(float64(_gfcfe), 0, _gg.Pow(2, float64(_cdcd.BitsPerSample)), _cbdeg[2*_ebfg], _cbdeg[2*_ebfg+1])
		_gddcf := _gg.Min(_gg.Max(_cgeec, _cdcd.Range[2*_ebfg]), _cdcd.Range[2*_ebfg+1])
		_bacd = append(_bacd, _gddcf)
	}
	return _bacd, nil
}

// SetContext sets the specific fielddata type, e.g. would be PdfFieldButton for a button field.
func (_abae *PdfField) SetContext(ctx PdfModel) { _abae._fbedg = ctx }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the A, B and C components of the color.
func (_edff *PdfColorspaceCalRGB) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fefb, _faeaa := _eb.GetNumbersAsFloat(objects)
	if _faeaa != nil {
		return nil, _faeaa
	}
	return _edff.ColorFromFloats(_fefb)
}
func (_ecefg *PdfWriter) addObjects(_adcbd _eb.PdfObject) error {
	_ddb.Log.Trace("\u0041d\u0064i\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u0021")
	if _cabef, _ebddg := _adcbd.(*_eb.PdfIndirectObject); _ebddg {
		_ddb.Log.Trace("\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074")
		_ddb.Log.Trace("\u002d \u0025\u0073\u0020\u0028\u0025\u0070)", _adcbd, _cabef)
		_ddb.Log.Trace("\u002d\u0020\u0025\u0073", _cabef.PdfObject)
		if _ecefg.addObject(_cabef) {
			_gdcab := _ecefg.addObjects(_cabef.PdfObject)
			if _gdcab != nil {
				return _gdcab
			}
		}
		return nil
	}
	if _adda, _bgfa := _adcbd.(*_eb.PdfObjectStream); _bgfa {
		_ddb.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d")
		_ddb.Log.Trace("\u002d \u0025\u0073\u0020\u0025\u0070", _adcbd, _adcbd)
		if _ecefg.addObject(_adda) {
			_baddd := _ecefg.addObjects(_adda.PdfObjectDictionary)
			if _baddd != nil {
				return _baddd
			}
		}
		return nil
	}
	if _debbe, _fcdga := _adcbd.(*_eb.PdfObjectDictionary); _fcdga {
		_ddb.Log.Trace("\u0044\u0069\u0063\u0074")
		_ddb.Log.Trace("\u002d\u0020\u0025\u0073", _adcbd)
		for _, _faga := range _debbe.Keys() {
			_gede := _debbe.Get(_faga)
			if _fbdde, _bcgdg := _gede.(*_eb.PdfObjectReference); _bcgdg {
				_gede = _fbdde.Resolve()
				_debbe.Set(_faga, _gede)
			}
			if _faga != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _cccbd := _ecefg.addObjects(_gede); _cccbd != nil {
					return _cccbd
				}
			} else {
				if _, _bafdb := _gede.(*_eb.PdfObjectNull); _bafdb {
					continue
				}
				if _efdfe := _ecefg.hasObject(_gede); !_efdfe {
					_ddb.Log.Debug("P\u0061\u0072\u0065\u006e\u0074\u0020o\u0062\u006a\u0020\u006e\u006f\u0074 \u0061\u0064\u0064\u0065\u0064\u0020\u0079e\u0074\u0021\u0021\u0020\u0025\u0054\u0020\u0025\u0070\u0020%\u0076", _gede, _gede, _gede)
					_ecefg._cfggb[_gede] = append(_ecefg._cfggb[_gede], _debbe)
				}
			}
		}
		return nil
	}
	if _eeacb, _acecb := _adcbd.(*_eb.PdfObjectArray); _acecb {
		_ddb.Log.Trace("\u0041\u0072\u0072a\u0079")
		_ddb.Log.Trace("\u002d\u0020\u0025\u0073", _adcbd)
		if _eeacb == nil {
			return _dcf.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _acabc, _dgbg := range _eeacb.Elements() {
			if _gefdd, _cccbb := _dgbg.(*_eb.PdfObjectReference); _cccbb {
				_dgbg = _gefdd.Resolve()
				_eeacb.Set(_acabc, _dgbg)
			}
			if _gdgge := _ecefg.addObjects(_dgbg); _gdgge != nil {
				return _gdgge
			}
		}
		return nil
	}
	if _, _fcbac := _adcbd.(*_eb.PdfObjectReference); _fcbac {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0062\u0065\u0020\u0061\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u002d\u0020\u0067\u006f\u0074 \u0025\u0023\u0076\u0021", _adcbd)
		return _dcf.New("r\u0065\u0066\u0065\u0072en\u0063e\u0020\u006e\u006f\u0074\u0020a\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	return nil
}

// PdfField contains the common attributes of a form field. The context object contains the specific field data
// which can represent a button, text, choice or signature.
// The PdfField is typically not used directly, but is encapsulated by the more specific field types such as
// PdfFieldButton etc (i.e. the context attribute).
type PdfField struct {
	_fbedg       PdfModel
	_adgda       *_eb.PdfIndirectObject
	Parent       *PdfField
	Annotations  []*PdfAnnotationWidget
	Kids         []*PdfField
	FT           *_eb.PdfObjectName
	T            *_eb.PdfObjectString
	TU           *_eb.PdfObjectString
	TM           *_eb.PdfObjectString
	Ff           *_eb.PdfObjectInteger
	V            _eb.PdfObject
	DV           _eb.PdfObject
	AA           _eb.PdfObject
	VariableText *VariableText
}

// ToPdfObject return the CalGray colorspace as a PDF object (name dictionary).
func (_gdged *PdfColorspaceCalGray) ToPdfObject() _eb.PdfObject {
	_cabe := &_eb.PdfObjectArray{}
	_cabe.Append(_eb.MakeName("\u0043a\u006c\u0047\u0072\u0061\u0079"))
	_ddbdc := _eb.MakeDict()
	if _gdged.WhitePoint != nil {
		_ddbdc.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _eb.MakeArray(_eb.MakeFloat(_gdged.WhitePoint[0]), _eb.MakeFloat(_gdged.WhitePoint[1]), _eb.MakeFloat(_gdged.WhitePoint[2])))
	} else {
		_ddb.Log.Error("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006fi\u006e\u0074\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _gdged.BlackPoint != nil {
		_ddbdc.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _eb.MakeArray(_eb.MakeFloat(_gdged.BlackPoint[0]), _eb.MakeFloat(_gdged.BlackPoint[1]), _eb.MakeFloat(_gdged.BlackPoint[2])))
	}
	_ddbdc.Set("\u0047\u0061\u006dm\u0061", _eb.MakeFloat(_gdged.Gamma))
	_cabe.Append(_ddbdc)
	if _gdged._gdee != nil {
		_gdged._gdee.PdfObject = _cabe
		return _gdged._gdee
	}
	return _cabe
}

// ParserMetadata gets the parser  metadata.
func (_eeaaa *CompliancePdfReader) ParserMetadata() _eb.ParserMetadata {
	if _eeaaa._fgafb == (_eb.ParserMetadata{}) {
		_eeaaa._fgafb, _ = _eeaaa._ebbe.ParserMetadata()
	}
	return _eeaaa._fgafb
}

// String returns a string that describes `base`.
func (_efadd fontCommon) String() string {
	return _e.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0073\u007d", _efadd.coreString())
}

// NewPdfRectangle creates a PDF rectangle object based on an input array of 4 integers.
// Defining the lower left (LL) and upper right (UR) corners with
// floating point numbers.
func NewPdfRectangle(arr _eb.PdfObjectArray) (*PdfRectangle, error) {
	_edgfbd := PdfRectangle{}
	if arr.Len() != 4 {
		return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065\u0020\u0061\u0072r\u0061\u0079\u002c\u0020\u006c\u0065\u006e \u0021\u003d\u0020\u0034")
	}
	var _dfddg error
	_edgfbd.Llx, _dfddg = _eb.GetNumberAsFloat(arr.Get(0))
	if _dfddg != nil {
		return nil, _dfddg
	}
	_edgfbd.Lly, _dfddg = _eb.GetNumberAsFloat(arr.Get(1))
	if _dfddg != nil {
		return nil, _dfddg
	}
	_edgfbd.Urx, _dfddg = _eb.GetNumberAsFloat(arr.Get(2))
	if _dfddg != nil {
		return nil, _dfddg
	}
	_edgfbd.Ury, _dfddg = _eb.GetNumberAsFloat(arr.Get(3))
	if _dfddg != nil {
		return nil, _dfddg
	}
	return &_edgfbd, nil
}

// NewPdfReaderLazy creates a new PdfReader for `rs` in lazy-loading mode. The difference
// from NewPdfReader is that in lazy-loading mode, objects are only loaded into memory when needed
// rather than entire structure being loaded into memory on reader creation.
// Note that it may make sense to use the lazy-load reader when processing only parts of files,
// rather than loading entire file into memory. Example: splitting a few pages from a large PDF file.
func NewPdfReaderLazy(rs _bagf.ReadSeeker) (*PdfReader, error) {
	const _fcagd = "\u006d\u006f\u0064\u0065l:\u004e\u0065\u0077\u0050\u0064\u0066\u0052\u0065\u0061\u0064\u0065\u0072\u004c\u0061z\u0079"
	return _dcgfe(rs, &ReaderOpts{LazyLoad: true}, false, _fcagd)
}

// SetContentStreams sets the content streams based on a string array. Will make
// 1 object stream for each string and reference from the page Contents.
// Each stream will be encoded using the encoding specified by the StreamEncoder,
// if empty, will use identity encoding (raw data).
func (_ddgbc *PdfPage) SetContentStreams(cStreams []string, encoder _eb.StreamEncoder) error {
	if len(cStreams) == 0 {
		_ddgbc.Contents = nil
		return nil
	}
	if encoder == nil {
		encoder = _eb.NewRawEncoder()
	}
	var _ddfce []*_eb.PdfObjectStream
	for _, _ebggbe := range cStreams {
		_gedbf := &_eb.PdfObjectStream{}
		_egdd := encoder.MakeStreamDict()
		_cbfbb, _fgab := encoder.EncodeBytes([]byte(_ebggbe))
		if _fgab != nil {
			return _fgab
		}
		_egdd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_cbfbb))))
		_gedbf.PdfObjectDictionary = _egdd
		_gedbf.Stream = _cbfbb
		_ddfce = append(_ddfce, _gedbf)
	}
	if len(_ddfce) == 1 {
		_ddgbc.Contents = _ddfce[0]
	} else {
		_bafa := _eb.MakeArray()
		for _, _gggee := range _ddfce {
			_bafa.Append(_gggee)
		}
		_ddgbc.Contents = _bafa
	}
	return nil
}

// GetObjectNums returns the object numbers of the PDF objects in the file
// Numbered objects are either indirect objects or stream objects.
// e.g. objNums := pdfReader.GetObjectNums()
// The underlying objects can then be accessed with
// pdfReader.GetIndirectObjectByNumber(objNums[0]) for the first available object.
func (_bgdbe *PdfReader) GetObjectNums() []int { return _bgdbe._ebbe.GetObjectNums() }
func (_dffaf *PdfReader) loadDSS() (*DSS, error) {
	if _dffaf._ebbe.GetCrypter() != nil && !_dffaf._ebbe.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_gdacf := _dffaf._bagcfd.Get("\u0044\u0053\u0053")
	if _gdacf == nil {
		return nil, nil
	}
	_gcddb, _ := _eb.GetIndirect(_gdacf)
	_gdacf = _eb.TraceToDirectObject(_gdacf)
	switch _acgfc := _gdacf.(type) {
	case *_eb.PdfObjectNull:
		return nil, nil
	case *_eb.PdfObjectDictionary:
		return _gacff(_gcddb, _acgfc)
	}
	return nil, _e.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u0053\u0053 \u0065\u006e\u0074\u0072y \u0025\u0054", _gdacf)
}

// ColorToRGB converts a Lab color to an RGB color.
func (_ebag *PdfColorspaceLab) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ccgf := func(_egfbf float64) float64 {
		if _egfbf >= 6.0/29 {
			return _egfbf * _egfbf * _egfbf
		}
		return 108.0 / 841 * (_egfbf - 4.0/29.0)
	}
	_badd, _cgbg := color.(*PdfColorLab)
	if !_cgbg {
		_ddb.Log.Debug("\u0069\u006e\u0070\u0075t \u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u006c\u0061\u0062")
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	LStar := _badd.L()
	AStar := _badd.A()
	BStar := _badd.B()
	L := (LStar+16)/116 + AStar/500
	M := (LStar + 16) / 116
	N := (LStar+16)/116 - BStar/200
	X := _ebag.WhitePoint[0] * _ccgf(L)
	Y := _ebag.WhitePoint[1] * _ccgf(M)
	Z := _ebag.WhitePoint[2] * _ccgf(N)
	_aafb := 3.240479*X + -1.537150*Y + -0.498535*Z
	_dfba := -0.969256*X + 1.875992*Y + 0.041556*Z
	_acba := 0.055648*X + -0.204043*Y + 1.057311*Z
	_aafb = _gg.Min(_gg.Max(_aafb, 0), 1.0)
	_dfba = _gg.Min(_gg.Max(_dfba, 0), 1.0)
	_acba = _gg.Min(_gg.Max(_acba, 0), 1.0)
	return NewPdfColorDeviceRGB(_aafb, _dfba, _acba), nil
}

// GetCIDToGIDMapObject get the underlying CIDToGIDMap object if the font type is CIDFontType2.
func (_ebdc *PdfFont) GetCIDToGIDMapObject() _eb.PdfObject {
	_gcaf, _bbcg := _ebdc._fdaa.(*pdfCIDFontType2)
	if _bbcg {
		return _gcaf.CIDToGIDMap
	}
	return nil
}
func (_gddg *pdfFontType3) getFontDescriptor() *PdfFontDescriptor { return _gddg._bged }

// GenerateHashMaps generates DSS hashmaps for Certificates, OCSPs and CRLs to make sure they are unique.
func (_accd *DSS) GenerateHashMaps() error {
	_afcfa, _degda := _accd.generateHashMap(_accd.Certs)
	if _degda != nil {
		return _degda
	}
	_acdgg, _degda := _accd.generateHashMap(_accd.OCSPs)
	if _degda != nil {
		return _degda
	}
	_gcad, _degda := _accd.generateHashMap(_accd.CRLs)
	if _degda != nil {
		return _degda
	}
	_accd._eaed = _afcfa
	_accd._gccb = _acdgg
	_accd._adcdd = _gcad
	return nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_acdbf *PdfColorspaceSpecialPattern) ToPdfObject() _eb.PdfObject {
	if _acdbf.UnderlyingCS == nil {
		return _eb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e")
	}
	_cffc := _eb.MakeArray(_eb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_cffc.Append(_acdbf.UnderlyingCS.ToPdfObject())
	if _acdbf._fgff != nil {
		_acdbf._fgff.PdfObject = _cffc
		return _acdbf._fgff
	}
	return _cffc
}

// PdfAnnotationFileAttachment represents FileAttachment annotations.
// (Section 12.5.6.15).
type PdfAnnotationFileAttachment struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	FS   _eb.PdfObject
	Name _eb.PdfObject
}

// SetPrintArea sets the value of the printArea.
func (_daff *ViewerPreferences) SetPrintArea(printArea PageBoundary) { _daff._bdegd = printArea }

// PdfActionSetOCGState represents a SetOCGState action.
type PdfActionSetOCGState struct {
	*PdfAction
	State      _eb.PdfObject
	PreserveRB _eb.PdfObject
}

func (_fcg *PdfAnnotation) String() string {
	_eag := ""
	_dggd, _acd := _fcg.ToPdfObject().(*_eb.PdfIndirectObject)
	if _acd {
		_eag = _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _fcg._cdb, _dggd.PdfObject.String())
	}
	return _eag
}

// PdfColorspaceICCBased format [/ICCBased stream]
//
// The stream shall contain the ICC profile.
// A conforming reader shall support ICC.1:2004:10 as required by PDF 1.7, which will enable it
// to properly render all embedded ICC profiles regardless of the PDF version
//
// In the current implementation, we rely on the alternative colormap provided.
type PdfColorspaceICCBased struct {
	N         int
	Alternate PdfColorspace

	// If omitted ICC not supported: then use DeviceGray,
	// DeviceRGB or DeviceCMYK for N=1,3,4 respectively.
	Range    []float64
	Metadata *_eb.PdfObjectStream
	Data     []byte
	_ebggg   *_eb.PdfIndirectObject
	_cdebc   *_eb.PdfObjectStream
}

// GetNumComponents returns the number of color components.
func (_fdcb *PdfColorspaceICCBased) GetNumComponents() int { return _fdcb.N }

// NewPdfFieldSignature returns an initialized signature field.
func NewPdfFieldSignature(signature *PdfSignature) *PdfFieldSignature {
	_dfdaa := &PdfFieldSignature{}
	_dfdaa.PdfField = NewPdfField()
	_dfdaa.PdfField.SetContext(_dfdaa)
	_dfdaa.PdfAnnotationWidget = NewPdfAnnotationWidget()
	_dfdaa.PdfAnnotationWidget.SetContext(_dfdaa)
	_dfdaa.PdfAnnotationWidget._ggf = _dfdaa.PdfField._adgda
	_dfdaa.T = _eb.MakeString("")
	_dfdaa.F = _eb.MakeInteger(132)
	_dfdaa.V = signature
	return _dfdaa
}

// GetFillImage get attached model.Image in push button.
func (_bdaf *PdfFieldButton) GetFillImage() *Image {
	if _bdaf.IsPush() {
		return _bdaf._cadc
	}
	return nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a Lab device.
func (_dedgd *PdfColorspaceLab) GetNumComponents() int { return 3 }

// String implements interface PdfObject.
func (_egb *PdfAction) String() string {
	_fec, _dgg := _egb.ToPdfObject().(*_eb.PdfIndirectObject)
	if _dgg {
		return _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _egb._aee, _fec.PdfObject.String())
	}
	return ""
}

var (
	TabOrderRow       TabOrderType = "\u0052"
	TabOrderColumn    TabOrderType = "\u0043"
	TabOrderStructure TabOrderType = "\u0053"
)

// DefaultImageHandler is the default implementation of the ImageHandler using the standard go library.
type DefaultImageHandler struct{}

// NewXObjectFormFromStream builds the Form XObject from a stream object.
// TODO: Should this be exposed? Consider different access points.
func NewXObjectFormFromStream(stream *_eb.PdfObjectStream) (*XObjectForm, error) {
	_egdcg := &XObjectForm{}
	_egdcg._afcag = stream
	_gbgaf := *(stream.PdfObjectDictionary)
	_faacc, _bcddgd := _eb.NewEncoderFromStream(stream)
	if _bcddgd != nil {
		return nil, _bcddgd
	}
	_egdcg.Filter = _faacc
	if _gbfdf := _gbgaf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _gbfdf != nil {
		_fdfgb, _eaagd := _gbfdf.(*_eb.PdfObjectName)
		if !_eaagd {
			return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if *_fdfgb != "\u0046\u006f\u0072\u006d" {
			_ddb.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
			return nil, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
		}
	}
	if _gdce := _gbgaf.Get("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065"); _gdce != nil {
		_egdcg.FormType = _gdce
	}
	if _aefaab := _gbgaf.Get("\u0042\u0042\u006f\u0078"); _aefaab != nil {
		_egdcg.BBox = _aefaab
	}
	if _dcgae := _gbgaf.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _dcgae != nil {
		_egdcg.Matrix = _dcgae
	}
	if _bgafe := _gbgaf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _bgafe != nil {
		_bgafe = _eb.TraceToDirectObject(_bgafe)
		_accfd, _abebc := _bgafe.(*_eb.PdfObjectDictionary)
		if !_abebc {
			_ddb.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0058\u004f\u0062j\u0065c\u0074\u0020\u0046\u006f\u0072\u006d\u0020\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u002c\u0020\u0070\u006f\u0069\u006e\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _eb.ErrTypeError
		}
		_dacae, _fbada := NewPdfPageResourcesFromDict(_accfd)
		if _fbada != nil {
			_ddb.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0020\u0072\u0065\u0073\u006f\u0075rc\u0065\u0073")
			return nil, _fbada
		}
		_egdcg.Resources = _dacae
		_ddb.Log.Trace("\u0046\u006f\u0072\u006d r\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0023\u0076", _egdcg.Resources)
	}
	_egdcg.Group = _gbgaf.Get("\u0047\u0072\u006fu\u0070")
	_egdcg.Ref = _gbgaf.Get("\u0052\u0065\u0066")
	_egdcg.MetaData = _gbgaf.Get("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061")
	_egdcg.PieceInfo = _gbgaf.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o")
	_egdcg.LastModified = _gbgaf.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
	_egdcg.StructParent = _gbgaf.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_egdcg.StructParents = _gbgaf.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
	_egdcg.OPI = _gbgaf.Get("\u004f\u0050\u0049")
	_egdcg.OC = _gbgaf.Get("\u004f\u0043")
	_egdcg.Name = _gbgaf.Get("\u004e\u0061\u006d\u0065")
	_egdcg.Stream = stream.Stream
	return _egdcg, nil
}

// ToPdfObject implements interface PdfModel.
func (_acgc *PdfAnnotationPopup) ToPdfObject() _eb.PdfObject {
	_acgc.PdfAnnotation.ToPdfObject()
	_bde := _acgc._ggf
	_aegd := _bde.PdfObject.(*_eb.PdfObjectDictionary)
	_aegd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0050\u006f\u0070u\u0070"))
	_aegd.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _acgc.Parent)
	_aegd.SetIfNotNil("\u004f\u0070\u0065\u006e", _acgc.Open)
	return _bde
}

// ToPdfObject returns an indirect object containing the signature field dictionary.
func (_bbbf *PdfFieldSignature) ToPdfObject() _eb.PdfObject {
	if _bbbf.PdfAnnotationWidget != nil {
		_bbbf.PdfAnnotationWidget.ToPdfObject()
	}
	_bbbf.PdfField.ToPdfObject()
	_fedfac := _bbbf._adgda
	_bggfg := _fedfac.PdfObject.(*_eb.PdfObjectDictionary)
	_bggfg.SetIfNotNil("\u0046\u0054", _eb.MakeName("\u0053\u0069\u0067"))
	_bggfg.SetIfNotNil("\u004c\u006f\u0063\u006b", _bbbf.Lock)
	_bggfg.SetIfNotNil("\u0053\u0056", _bbbf.SV)
	if _bbbf.V != nil {
		_bggfg.SetIfNotNil("\u0056", _bbbf.V.ToPdfObject())
	}
	return _fedfac
}

// NewCompositePdfFontFromTTFFile loads a composite font from a TTF font file. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTFFile.
func NewCompositePdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_fcggd, _bcfd := _ccb.Open(filePath)
	if _bcfd != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006f\u0070\u0065\u006e\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _bcfd)
		return nil, _bcfd
	}
	defer _fcggd.Close()
	return NewCompositePdfFontFromTTF(_fcggd)
}

// PdfBorderEffect represents a PDF border effect.
type PdfBorderEffect struct {
	S *BorderEffect
	I *float64
}

// ToPdfObject returns the PDF representation of the outline tree node.
func (_eeaf *PdfOutlineTreeNode) ToPdfObject() _eb.PdfObject { return _eeaf.GetContext().ToPdfObject() }

// NewPdfPageResourcesFromDict creates and returns a new PdfPageResources object
// from the input dictionary.
func NewPdfPageResourcesFromDict(dict *_eb.PdfObjectDictionary) (*PdfPageResources, error) {
	_fcbbe := NewPdfPageResources()
	if _eabf := dict.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _eabf != nil {
		_fcbbe.ExtGState = _eabf
	}
	if _ddegc := dict.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"); _ddegc != nil && !_eb.IsNullObject(_ddegc) {
		_fcbbe.ColorSpace = _ddegc
	}
	if _dgdad := dict.Get("\u0050a\u0074\u0074\u0065\u0072\u006e"); _dgdad != nil {
		_fcbbe.Pattern = _dgdad
	}
	if _ecgba := dict.Get("\u0053h\u0061\u0064\u0069\u006e\u0067"); _ecgba != nil {
		_fcbbe.Shading = _ecgba
	}
	if _efacdb := dict.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"); _efacdb != nil {
		_fcbbe.XObject = _efacdb
	}
	if _ggde := _eb.ResolveReference(dict.Get("\u0046\u006f\u006e\u0074")); _ggde != nil {
		_fcbbe.Font = _ggde
	}
	if _eaffa := dict.Get("\u0050r\u006f\u0063\u0053\u0065\u0074"); _eaffa != nil {
		_fcbbe.ProcSet = _eaffa
	}
	if _cebda := dict.Get("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"); _cebda != nil {
		_fcbbe.Properties = _cebda
	}
	return _fcbbe, nil
}

// GetShadingByName gets the shading specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_fdebd *PdfPageResources) GetShadingByName(keyName _eb.PdfObjectName) (*PdfShading, bool) {
	if _fdebd.Shading == nil {
		return nil, false
	}
	_cbedf, _bgfca := _eb.TraceToDirectObject(_fdebd.Shading).(*_eb.PdfObjectDictionary)
	if !_bgfca {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0068\u0061d\u0069\u006e\u0067\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _fdebd.Shading)
		return nil, false
	}
	if _acgeg := _cbedf.Get(keyName); _acgeg != nil {
		_baccc, _eabge := _aggce(_acgeg)
		if _eabge != nil {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0073\u0068\u0061\u0064\u0069\u006eg\u003a\u0020\u0025\u0076", _eabge)
			return nil, false
		}
		return _baccc, true
	}
	return nil, false
}
func (_acfg *PdfAcroForm) fillImageWithAppearance(_gdgde FieldImageProvider, _dfdfc FieldAppearanceGenerator) error {
	if _acfg == nil {
		return nil
	}
	_edeg, _eddgd := _gdgde.FieldImageValues()
	if _eddgd != nil {
		return _eddgd
	}
	for _, _gdgc := range _acfg.AllFields() {
		_afcb := _gdgc.PartialName()
		_caeag, _ffdb := _edeg[_afcb]
		if !_ffdb {
			if _dgag, _bfgfc := _gdgc.FullName(); _bfgfc == nil {
				_caeag, _ffdb = _edeg[_dgag]
			}
		}
		if !_ffdb {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _afcb)
			continue
		}
		switch _gbcg := _gdgc.GetContext().(type) {
		case *PdfFieldButton:
			if _gbcg.IsPush() {
				_gbcg.SetFillImage(_caeag)
			}
		}
		if _dfdfc == nil {
			continue
		}
		for _, _fedgf := range _gdgc.Annotations {
			_ffcaa, _ffba := _dfdfc.GenerateAppearanceDict(_acfg, _gdgc, _fedgf)
			if _ffba != nil {
				return _ffba
			}
			_fedgf.AP = _ffcaa
			_fedgf.ToPdfObject()
		}
	}
	return nil
}

// GetXHeight returns the XHeight of the font `descriptor`.
func (_cgegc *PdfFontDescriptor) GetXHeight() (float64, error) {
	return _eb.GetNumberAsFloat(_cgegc.XHeight)
}

// NewPdfShadingType2 creates an empty shading type 2 dictionary.
func NewPdfShadingType2() *PdfShadingType2 {
	_dgedgb := &PdfShadingType2{}
	_dgedgb.PdfShading = &PdfShading{}
	_dgedgb.PdfShading._cefaa = _eb.MakeIndirectObject(_eb.MakeDict())
	_dgedgb.PdfShading._ecffg = _dgedgb
	return _dgedgb
}

// PdfVersion returns version of the PDF file.
func (_dfeae *PdfReader) PdfVersion() _eb.Version { return _dfeae._ebbe.PdfVersion() }

// Clear clears flag fl from the flag and returns the resulting flag.
func (_bgac FieldFlag) Clear(fl FieldFlag) FieldFlag { return FieldFlag(_bgac.Mask() &^ fl.Mask()) }

// ToPdfObject returns a PdfObject representation of PdfColorspaceDeviceNAttributes as a PdfObjectDictionary directly
// or indirectly within an indirect object container.
func (_babfbc *PdfColorspaceDeviceNAttributes) ToPdfObject() _eb.PdfObject {
	_fddbd := _eb.MakeDict()
	if _babfbc.Subtype != nil {
		_fddbd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _babfbc.Subtype)
	}
	_fddbd.SetIfNotNil("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s", _babfbc.Colorants)
	_fddbd.SetIfNotNil("\u0050r\u006f\u0063\u0065\u0073\u0073", _babfbc.Process)
	_fddbd.SetIfNotNil("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073", _babfbc.MixingHints)
	if _babfbc._dbggc != nil {
		_babfbc._dbggc.PdfObject = _fddbd
		return _babfbc._dbggc
	}
	return _fddbd
}

// SetCatalogStructTreeRoot sets the catalog struct tree root object.
func (_bdgfb *PdfWriter) SetCatalogStructTreeRoot(tree _eb.PdfObject) error {
	if tree == nil {
		_bdgfb._dbffa.Remove("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
		return nil
	}
	_bdgfb._dbffa.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", tree)
	return _bdgfb.addObjects(tree)
}

// Permissions specify a permissions dictionary (PDF 1.5).
// (Section 12.8.4, Table 258 - Entries in a permissions dictionary p. 477 in PDF32000_2008).
type Permissions struct {
	DocMDP *PdfSignature
	_cgbag *_eb.PdfObjectDictionary
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_cgcagf *PdfOutline) ToPdfObject() _eb.PdfObject {
	_cgec := _cgcagf._becfb
	_dcgea := _cgec.PdfObject.(*_eb.PdfObjectDictionary)
	_dcgea.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"))
	if _cgcagf.First != nil {
		_dcgea.Set("\u0046\u0069\u0072s\u0074", _cgcagf.First.ToPdfObject())
	}
	if _cgcagf.Last != nil {
		_dcgea.Set("\u004c\u0061\u0073\u0074", _cgcagf.Last.GetContext().GetContainingPdfObject())
	}
	if _cgcagf.Parent != nil {
		_dcgea.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cgcagf.Parent.GetContext().GetContainingPdfObject())
	}
	if _cgcagf.Count != nil {
		_dcgea.Set("\u0043\u006f\u0075n\u0074", _eb.MakeInteger(*_cgcagf.Count))
	}
	return _cgec
}

// ToPdfObject implements interface PdfModel.
func (_gfafe *PdfAnnotationLink) ToPdfObject() _eb.PdfObject {
	_gfafe.PdfAnnotation.ToPdfObject()
	_debg := _gfafe._ggf
	_fede := _debg.PdfObject.(*_eb.PdfObjectDictionary)
	_fede.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u004c\u0069\u006e\u006b"))
	if _gfafe._fbg != nil && _gfafe._fbg._aee != nil {
		_fede.Set("\u0041", _gfafe._fbg._aee.ToPdfObject())
	} else if _gfafe.A != nil {
		_fede.Set("\u0041", _gfafe.A)
	}
	_fede.SetIfNotNil("\u0044\u0065\u0073\u0074", _gfafe.Dest)
	_fede.SetIfNotNil("\u0048", _gfafe.H)
	_fede.SetIfNotNil("\u0050\u0041", _gfafe.PA)
	_fede.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _gfafe.QuadPoints)
	_fede.SetIfNotNil("\u0042\u0053", _gfafe.BS)
	return _debg
}

// ApplyStandard is used to apply changes required on the document to match the rules required by the input standard.
// The writer's content would be changed after all the document parts are already established during the Write method.
// A good example of the StandardApplier could be a PDF/A Profile (i.e.: pdfa.Profile1A). In such a case PdfWriter would
// set up all rules required by that Profile.
func (_eeeb *PdfWriter) ApplyStandard(optimizer StandardApplier) { _eeeb._ggbcd = optimizer }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_fafae *PdfShading) ToPdfObject() _eb.PdfObject {
	_ceca := _fafae._cefaa
	_eadff, _caecf := _fafae.getShadingDict()
	if _caecf != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _fafae.ShadingType != nil {
		_eadff.Set("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _fafae.ShadingType)
	}
	if _fafae.ColorSpace != nil {
		_eadff.Set("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _fafae.ColorSpace.ToPdfObject())
	}
	if _fafae.Background != nil {
		_eadff.Set("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064", _fafae.Background)
	}
	if _fafae.BBox != nil {
		_eadff.Set("\u0042\u0042\u006f\u0078", _fafae.BBox.ToPdfObject())
	}
	if _fafae.AntiAlias != nil {
		_eadff.Set("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s", _fafae.AntiAlias)
	}
	return _ceca
}

// DecodeArray returns the component range values for the Indexed colorspace.
func (_dacb *PdfColorspaceSpecialIndexed) DecodeArray() []float64 {
	return []float64{0, float64(_dacb.HiVal)}
}

// NewPdfActionSetOCGState returns a new "named" action.
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	_cdd := NewPdfAction()
	_fbd := &PdfActionSetOCGState{}
	_fbd.PdfAction = _cdd
	_cdd.SetContext(_fbd)
	return _fbd
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_bggb pdfFontType0) GetCharMetrics(code _fc.CharCode) (_fg.CharMetrics, bool) {
	if _bggb.DescendantFont == nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _bggb)
		return _fg.CharMetrics{}, false
	}
	return _bggb.DescendantFont.GetCharMetrics(code)
}

// GetContext returns the PdfField context which is the more specific field data type, e.g. PdfFieldButton
// for a button field.
func (_efdcc *PdfField) GetContext() PdfModel { return _efdcc._fbedg }
func _dcadb(_dccdc *_eb.PdfObjectDictionary, _abcf *fontCommon) (*pdfCIDFontType0, error) {
	if _abcf._fgdee != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" {
		_ddb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0030\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _abcf)
		return nil, _eb.ErrRangeError
	}
	_fcecg := _cgag(_abcf)
	_bbff, _addb := _eb.GetDict(_dccdc.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_addb {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _abcf)
		return nil, ErrRequiredAttributeMissing
	}
	_fcecg.CIDSystemInfo = _bbff
	_fcecg.DW = _dccdc.Get("\u0044\u0057")
	_fcecg.W = _dccdc.Get("\u0057")
	_fcecg.DW2 = _dccdc.Get("\u0044\u0057\u0032")
	_fcecg.W2 = _dccdc.Get("\u0057\u0032")
	_fcecg._dcfg = 1000.0
	if _dgabg, _ffddc := _eb.GetNumberAsFloat(_fcecg.DW); _ffddc == nil {
		_fcecg._dcfg = _dgabg
	}
	_eacc, _efeff := _cbfdd(_fcecg.W)
	if _efeff != nil {
		return nil, _efeff
	}
	if _eacc == nil {
		_eacc = map[_fc.CharCode]float64{}
	}
	_fcecg._efege = _eacc
	return _fcecg, nil
}
func _cdeea() string {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	if len(_geba) > 0 {
		return _geba
	}
	return "\u0055n\u0069\u0044\u006f\u0063 \u002d\u0020\u0068\u0074\u0074p\u003a/\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069o"
}

// OutlineDest represents the destination of an outline item.
// It holds the page and the position on the page an outline item points to.
type OutlineDest struct {
	PageObj *_eb.PdfIndirectObject `json:"-"`
	Page    int64                  `json:"page"`
	Mode    string                 `json:"mode"`
	X       float64                `json:"x"`
	Y       float64                `json:"y"`
	Zoom    float64                `json:"zoom"`
}

// NewPdfAnnotationScreen returns a new screen annotation.
func NewPdfAnnotationScreen() *PdfAnnotationScreen {
	_cac := NewPdfAnnotation()
	_ebgf := &PdfAnnotationScreen{}
	_ebgf.PdfAnnotation = _cac
	_cac.SetContext(_ebgf)
	return _ebgf
}

type pdfFontSimple struct {
	fontCommon
	_ffgdb *_eb.PdfIndirectObject
	_cegda map[_fc.CharCode]float64
	_eccbb _fc.TextEncoder
	_gcgb  _fc.TextEncoder
	_debdb *PdfFontDescriptor

	// Encoding is subject to limitations that are described in 9.6.6, "Character Encoding".
	// BaseFont is derived differently.
	FirstChar _eb.PdfObject
	LastChar  _eb.PdfObject
	Widths    _eb.PdfObject
	Encoding  _eb.PdfObject
	_gacdb    *_fg.RuneCharSafeMap
}

// FileRelationship represents a attachment file relationship type.
type FileRelationship int

func (_dbcg *pdfCIDFontType2) baseFields() *fontCommon { return &_dbcg.fontCommon }
func (_eagf *PdfReader) newPdfAnnotation3DFromDict(_ceeed *_eb.PdfObjectDictionary) (*PdfAnnotation3D, error) {
	_afdc := PdfAnnotation3D{}
	_afdc.T3DD = _ceeed.Get("\u0033\u0044\u0044")
	_afdc.T3DV = _ceeed.Get("\u0033\u0044\u0056")
	_afdc.T3DA = _ceeed.Get("\u0033\u0044\u0041")
	_afdc.T3DI = _ceeed.Get("\u0033\u0044\u0049")
	_afdc.T3DB = _ceeed.Get("\u0033\u0044\u0042")
	return &_afdc, nil
}

// GetContainingPdfObject returns the container of the outline (indirect object).
func (_dgfb *PdfOutline) GetContainingPdfObject() _eb.PdfObject { return _dgfb._becfb }

// BorderEffect represents a border effect (Table 167 p. 395).
type BorderEffect int

// Sign signs a specific page with a digital signature.
// The signature field parameter must have a valid signature dictionary
// specified by its V field.
func (_ddce *PdfAppender) Sign(pageNum int, field *PdfFieldSignature) error {
	if field == nil {
		return _dcf.New("\u0073\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 n\u0069\u006c")
	}
	_aefa := field.V
	if _aefa == nil {
		return _dcf.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061r\u0079 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_deeaf := pageNum - 1
	if _deeaf < 0 || _deeaf > len(_ddce._fedg)-1 {
		return _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064", pageNum)
	}
	_fdgd := _ddce.Reader.PageList[_deeaf]
	field.P = _fdgd.ToPdfObject()
	if field.T == nil || field.T.String() == "" {
		field.T = _eb.MakeString(_e.Sprintf("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072\u0065\u0020\u0025\u0064", pageNum))
	}
	_fdgd.AddAnnotation(field.PdfAnnotationWidget.PdfAnnotation)
	if _ddce._bddda == _ddce._adac.AcroForm {
		_ddce._bddda = _ddce.Reader.AcroForm
	}
	_afac := _ddce._bddda
	if _afac == nil {
		_afac = NewPdfAcroForm()
	}
	_afac.SigFlags = _eb.MakeInteger(3)
	if _afac.NeedAppearances != nil {
		_afac.NeedAppearances = nil
	}
	_fegf := append(_afac.AllFields(), field.PdfField)
	_afac.Fields = &_fegf
	_ddce.ReplaceAcroForm(_afac)
	_ddce.UpdatePage(_fdgd)
	_ddce._fedg[_deeaf] = _fdgd
	if _, _gbac := field.V.GetDocMDPPermission(); _gbac {
		_ddce._fbec = NewPermissions(field.V)
	}
	return nil
}
func _ggaa(_gcfa *_df.ImageBase) (_agfcd Image) {
	_agfcd.Width = int64(_gcfa.Width)
	_agfcd.Height = int64(_gcfa.Height)
	_agfcd.BitsPerComponent = int64(_gcfa.BitsPerComponent)
	_agfcd.ColorComponents = _gcfa.ColorComponents
	_agfcd.Data = _gcfa.Data
	_agfcd._fedc = _gcfa.Decode
	_agfcd._bdcab = _gcfa.Alpha
	return _agfcd
}

// KValue is a wrapper object to hold various type of K's children objects.
type KValue struct {
	_ccbca *KDict
	_fbgaa _eb.PdfObject
	_ddaf  *int
}

// UpdateObject marks `obj` as updated and to be included in the following revision.
func (_gdafc *PdfAppender) UpdateObject(obj _eb.PdfObject) {
	_gdafc.replaceObject(obj, obj)
	if _, _ddf := _gdafc._accfg[obj]; !_ddf {
		_gdafc._dfg = append(_gdafc._dfg, obj)
		_gdafc._accfg[obj] = struct{}{}
	}
}
func (_dcgf *PdfAppender) mergeResources(_dfdef, _cgc _eb.PdfObject, _bfb map[_eb.PdfObjectName]_eb.PdfObjectName) _eb.PdfObject {
	if _cgc == nil && _dfdef == nil {
		return nil
	}
	if _cgc == nil {
		return _dfdef
	}
	_gcce, _ebed := _eb.GetDict(_cgc)
	if !_ebed {
		return _dfdef
	}
	if _dfdef == nil {
		_befd := _eb.MakeDict()
		_befd.Merge(_gcce)
		return _cgc
	}
	_bfcf, _ebed := _eb.GetDict(_dfdef)
	if !_ebed {
		_ddb.Log.Error("\u0045\u0072\u0072or\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		_bfcf = _eb.MakeDict()
	}
	for _, _gfced := range _gcce.Keys() {
		if _fega, _abec := _bfb[_gfced]; _abec {
			_bfcf.Set(_fega, _gcce.Get(_gfced))
		} else {
			_bfcf.Set(_gfced, _gcce.Get(_gfced))
		}
	}
	return _bfcf
}
func _eebd(_eacbe *_eb.PdfObjectDictionary) (*PdfShadingType1, error) {
	_ggbdbc := PdfShadingType1{}
	if _gggga := _eacbe.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _gggga != nil {
		_gggga = _eb.TraceToDirectObject(_gggga)
		_cdac, _gdddg := _gggga.(*_eb.PdfObjectArray)
		if !_gdddg {
			_ddb.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gggga)
			return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ggbdbc.Domain = _cdac
	}
	if _fgdfc := _eacbe.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _fgdfc != nil {
		_fgdfc = _eb.TraceToDirectObject(_fgdfc)
		_bcfdd, _abccge := _fgdfc.(*_eb.PdfObjectArray)
		if !_abccge {
			_ddb.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _fgdfc)
			return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ggbdbc.Matrix = _bcfdd
	}
	_fcacc := _eacbe.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _fcacc == nil {
		_ddb.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_ggbdbc.Function = []PdfFunction{}
	if _feccc, _fddge := _fcacc.(*_eb.PdfObjectArray); _fddge {
		for _, _ccgg := range _feccc.Elements() {
			_degfba, _eacgae := _cccfa(_ccgg)
			if _eacgae != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _eacgae)
				return nil, _eacgae
			}
			_ggbdbc.Function = append(_ggbdbc.Function, _degfba)
		}
	} else {
		_aadccb, _fgafec := _cccfa(_fcacc)
		if _fgafec != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fgafec)
			return nil, _fgafec
		}
		_ggbdbc.Function = append(_ggbdbc.Function, _aadccb)
	}
	return &_ggbdbc, nil
}

// NewPdfColorCalRGB returns a new CalRBG color.
func NewPdfColorCalRGB(a, b, c float64) *PdfColorCalRGB {
	_ecdf := PdfColorCalRGB{a, b, c}
	return &_ecdf
}

// WriteToFile writes the output PDF to file.
func (_caded *PdfWriter) WriteToFile(outputFilePath string) error {
	_aacdb, _fabea := _ccb.Create(outputFilePath)
	if _fabea != nil {
		return _fabea
	}
	defer _aacdb.Close()
	return _caded.Write(_aacdb)
}

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	_bg := &PdfAction{}
	_bg._dee = _eb.MakeIndirectObject(_eb.MakeDict())
	return _bg
}
func (_fdegb *PdfWriter) AttachFile(file *EmbeddedFile) error {
	_dgeea := _fdegb._agdbc
	if _dgeea == nil {
		_dgeea = _gaefa()
	}
	_ebeee := _dgeea.addEmbeddedFile(file)
	if _ebeee != nil {
		return _ebeee
	}
	_fdegb._agdbc = _dgeea
	return nil
}

// PdfAnnotationLine represents Line annotations.
// (Section 12.5.6.7).
type PdfAnnotationLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	L       _eb.PdfObject
	BS      _eb.PdfObject
	LE      _eb.PdfObject
	IC      _eb.PdfObject
	LL      _eb.PdfObject
	LLE     _eb.PdfObject
	Cap     _eb.PdfObject
	IT      _eb.PdfObject
	LLO     _eb.PdfObject
	CP      _eb.PdfObject
	Measure _eb.PdfObject
	CO      _eb.PdfObject
}

// PdfActionGoToR represents a GoToR action.
type PdfActionGoToR struct {
	*PdfAction
	F         *PdfFilespec
	D         _eb.PdfObject
	NewWindow _eb.PdfObject
}

// PdfFunctionType0 uses a sequence of sample values (contained in a stream) to provide an approximation
// for functions whose domains and ranges are bounded. The samples are organized as an m-dimensional
// table in which each entry has n components
type PdfFunctionType0 struct {
	Domain        []float64
	Range         []float64
	NumInputs     int
	NumOutputs    int
	Size          []int
	BitsPerSample int
	Order         int
	Encode        []float64
	Decode        []float64
	_dcdgc        []byte
	_bgaaab       []uint32
	_bded         *_eb.PdfObjectStream
}

// NewStandard14Font returns the standard 14 font named `basefont` as a *PdfFont, or an error if it
// `basefont` is not one of the standard 14 font names.
func NewStandard14Font(basefont StdFontName) (*PdfFont, error) {
	_dbab, _dbdbe := _adca(basefont)
	if _dbdbe != nil {
		return nil, _dbdbe
	}
	if basefont != SymbolName && basefont != ZapfDingbatsName {
		_dbab._eccbb = _fc.NewWinAnsiEncoder()
	}
	return &PdfFont{_fdaa: &_dbab}, nil
}

// NewPdfReader returns a new PdfReader for an input io.ReadSeeker interface. Can be used to read PDF from
// memory or file. Immediately loads and traverses the PDF structure including pages and page contents (if
// not encrypted). Loads entire document structure into memory.
// Alternatively a lazy-loading reader can be created with NewPdfReaderLazy which loads only references,
// and references are loaded from disk into memory on an as-needed basis.
func NewPdfReader(rs _bagf.ReadSeeker) (*PdfReader, error) {
	const _gbgfa = "\u006do\u0064e\u006c\u003a\u004e\u0065\u0077P\u0064\u0066R\u0065\u0061\u0064\u0065\u0072"
	return _dcgfe(rs, &ReaderOpts{}, false, _gbgfa)
}

// AddOutlineTree adds outlines to a PDF file.
func (_bcagf *PdfWriter) AddOutlineTree(outlineTree *PdfOutlineTreeNode) { _bcagf._fbgge = outlineTree }

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_adccc *PdfColorspaceSpecialIndexed) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	N := _adccc.Base.GetNumComponents()
	_fbee := int(vals[0]) * N
	if _fbee < 0 || (_fbee+N-1) >= len(_adccc._efcb) {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fbee)
		return nil, ErrColorOutOfRange
	}
	_fgag := _adccc._efcb[_fbee : _fbee+N]
	var _fdbfc []float64
	for _, _ebcda := range _fgag {
		_fdbfc = append(_fdbfc, float64(_ebcda)/255.0)
	}
	_cgg, _acdc := _adccc.Base.ColorFromFloats(_fdbfc)
	if _acdc != nil {
		return nil, _acdc
	}
	return _cgg, nil
}

// EncryptionAlgorithm is used in EncryptOptions to change the default algorithm used to encrypt the document.
type EncryptionAlgorithm int

// SetEncoder sets the encoding for the underlying font.
// TODO(peterwilliams97): Change function signature to SetEncoder(encoder *textencoding.simpleEncoder).
// TODO(gunnsth): Makes sense if SetEncoder is removed from the interface fonts.Font as proposed in PR #260.
func (_bfabd *pdfFontSimple) SetEncoder(encoder _fc.TextEncoder) { _bfabd._eccbb = encoder }

// SignatureHandlerDocMDPParams describe the specific parameters for the SignatureHandlerEx
// These parameters describe how to check the difference between revisions.
// Revisions of the document get from the PdfParser.
type SignatureHandlerDocMDPParams struct {
	Parser     *_eb.PdfParser
	DiffPolicy _bab.DiffPolicy
}

// NewPdfColorspaceSpecialSeparation returns a new separation color.
func NewPdfColorspaceSpecialSeparation() *PdfColorspaceSpecialSeparation {
	_efgg := &PdfColorspaceSpecialSeparation{}
	return _efgg
}

// ToPdfObject returns the PDF representation of the function.
func (_ccdcf *PdfFunctionType4) ToPdfObject() _eb.PdfObject {
	_cdcf := _ccdcf._gbcd
	if _cdcf == nil {
		_ccdcf._gbcd = &_eb.PdfObjectStream{}
		_cdcf = _ccdcf._gbcd
	}
	_dadcf := _eb.MakeDict()
	_dadcf.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _eb.MakeInteger(4))
	_cabab := &_eb.PdfObjectArray{}
	for _, _agaef := range _ccdcf.Domain {
		_cabab.Append(_eb.MakeFloat(_agaef))
	}
	_dadcf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _cabab)
	_gbbgc := &_eb.PdfObjectArray{}
	for _, _becac := range _ccdcf.Range {
		_gbbgc.Append(_eb.MakeFloat(_becac))
	}
	_dadcf.Set("\u0052\u0061\u006eg\u0065", _gbbgc)
	if _ccdcf._bdcf == nil && _ccdcf.Program != nil {
		_ccdcf._bdcf = []byte(_ccdcf.Program.String())
	}
	_dadcf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_ccdcf._bdcf))))
	_cdcf.Stream = _ccdcf._bdcf
	_cdcf.PdfObjectDictionary = _dadcf
	return _cdcf
}

// PdfAnnotation3D represents 3D annotations.
// (Section 13.6.2).
type PdfAnnotation3D struct {
	*PdfAnnotation
	T3DD _eb.PdfObject
	T3DV _eb.PdfObject
	T3DA _eb.PdfObject
	T3DI _eb.PdfObject
	T3DB _eb.PdfObject
}

func _aeebg(_bddggd _eb.PdfObject) {
	_ddb.Log.Debug("\u006f\u0062\u006a\u003a\u0020\u0025\u0054\u0020\u0025\u0073", _bddggd, _bddggd.String())
	if _ebadf, _cfac := _bddggd.(*_eb.PdfObjectStream); _cfac {
		_ffef, _fdebf := _eb.DecodeStream(_ebadf)
		if _fdebf != nil {
			_ddb.Log.Debug("\u0045r\u0072\u006f\u0072\u003a\u0020\u0025v", _fdebf)
			return
		}
		_ddb.Log.Debug("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _ffef)
	} else if _ecefb, _eecgb := _bddggd.(*_eb.PdfIndirectObject); _eecgb {
		_ddb.Log.Debug("\u0025\u0054\u0020%\u0076", _ecefb.PdfObject, _ecefb.PdfObject)
		_ddb.Log.Debug("\u0025\u0073", _ecefb.PdfObject.String())
	}
}

// RemveTabOrder removes the tab order for the page.
func (_cbbgb *PdfPage) RemoveTabOrder() { _cbbgb.Tabs = nil }

// PdfAnnotationUnderline represents Underline annotations.
// (Section 12.5.6.10).
type PdfAnnotationUnderline struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _eb.PdfObject
}

// PdfShadingType5 is a Lattice-form Gouraud-shaded triangle mesh.
type PdfShadingType5 struct {
	*PdfShading
	BitsPerCoordinate *_eb.PdfObjectInteger
	BitsPerComponent  *_eb.PdfObjectInteger
	VerticesPerRow    *_eb.PdfObjectInteger
	Decode            *_eb.PdfObjectArray
	Function          []PdfFunction
}

// NewPdfPageResourcesColorspaces returns a new PdfPageResourcesColorspaces object.
func NewPdfPageResourcesColorspaces() *PdfPageResourcesColorspaces {
	_dbded := &PdfPageResourcesColorspaces{}
	_dbded.Names = []string{}
	_dbded.Colorspaces = map[string]PdfColorspace{}
	_dbded._ccba = &_eb.PdfIndirectObject{}
	return _dbded
}

// ReaderToWriterOpts options used to generate a PdfWriter.
type ReaderToWriterOpts struct {
	SkipAcroForm          bool
	SkipInfo              bool
	SkipNameDictionary    bool
	SkipNamedDests        bool
	SkipOCProperties      bool
	SkipOutlines          bool
	SkipPageLabels        bool
	SkipRotation          bool
	SkipMetadata          bool
	SkipMarkInfo          bool
	SkipViewerPreferences bool
	SkipLanguage          bool
	PageProcessCallback   PageProcessCallback

	// Deprecated: will be removed in v4. Use PageProcessCallback instead.
	PageCallback PageCallback
}

// GetRevisionNumber returns the version of the current Pdf document
func (_agbbb *PdfReader) GetRevisionNumber() int { return _agbbb._ebbe.GetRevisionNumber() }

// PdfActionGoTo3DView represents a GoTo3DView action.
type PdfActionGoTo3DView struct {
	*PdfAction
	TA _eb.PdfObject
	V  _eb.PdfObject
}

func (_gacbe *pdfFontSimple) getFontDescriptor() *PdfFontDescriptor {
	if _bffd := _gacbe._bged; _bffd != nil {
		return _bffd
	}
	return _gacbe._debdb
}

// String returns a string representation of the field.
func (_cbbbd *PdfField) String() string {
	if _aagdd, _efdaa := _cbbbd.ToPdfObject().(*_eb.PdfIndirectObject); _efdaa {
		return _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _cbbbd._fbedg, _aagdd.PdfObject.String())
	}
	return ""
}

// ToPdfObject convert PdfInfo to pdf object.
func (_bfde *PdfInfo) ToPdfObject() _eb.PdfObject {
	_bgaf := _eb.MakeDict()
	_bgaf.SetIfNotNil("\u0054\u0069\u0074l\u0065", _bfde.Title)
	_bgaf.SetIfNotNil("\u0041\u0075\u0074\u0068\u006f\u0072", _bfde.Author)
	_bgaf.SetIfNotNil("\u0053u\u0062\u006a\u0065\u0063\u0074", _bfde.Subject)
	_bgaf.SetIfNotNil("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _bfde.Keywords)
	_bgaf.SetIfNotNil("\u0043r\u0065\u0061\u0074\u006f\u0072", _bfde.Creator)
	_bgaf.SetIfNotNil("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _bfde.Producer)
	_bgaf.SetIfNotNil("\u0054r\u0061\u0070\u0070\u0065\u0064", _bfde.Trapped)
	if _bfde.CreationDate != nil {
		_bgaf.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _bfde.CreationDate.ToPdfObject())
	}
	if _bfde.ModifiedDate != nil {
		_bgaf.SetIfNotNil("\u004do\u0064\u0044\u0061\u0074\u0065", _bfde.ModifiedDate.ToPdfObject())
	}
	for _, _cdcc := range _bfde._cbfb.Keys() {
		_bgaf.SetIfNotNil(_cdcc, _bfde._cbfb.Get(_cdcc))
	}
	return _bgaf
}

// SetOCProperties sets the optional content properties.
func (_ccaaee *PdfWriter) SetOCProperties(ocProperties _eb.PdfObject) error {
	_dddedf := _ccaaee._dbffa
	if ocProperties != nil {
		_ddb.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u004f\u0043\u0020\u0050\u0072o\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002e\u002e\u002e")
		_dddedf.Set("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", ocProperties)
		return _ccaaee.addObjects(ocProperties)
	}
	return nil
}

// NewXObjectImageFromImageLazy creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
// If lazy is true, then lazy mode is enabled for XObject.
// Lazy mode allows to reduce memory usage with the help of temporary files.
func NewXObjectImageFromImageLazy(img *Image, cs PdfColorspace, encoder _eb.StreamEncoder, lazy bool) (*XObjectImage, error) {
	_fbfged := NewXObjectImage()
	if lazy {
		_bdbba, _gbea := UpdateXObjectImageFromImage(_fbfged, img, cs, encoder)
		if _gbea != nil {
			return nil, _gbea
		}
		_bdbba.ToPdfObject()
		_gbea = _bdbba._gceffg.MakeLazy()
		if _gbea != nil {
			return nil, _gbea
		}
		_bdbba.Stream = nil
		return _bdbba, nil
	}
	return UpdateXObjectImageFromImage(_fbfged, img, cs, encoder)
}

// ToPdfObject returns the PDF representation of the page resources.
func (_ffdea *PdfPageResources) ToPdfObject() _eb.PdfObject {
	_dcedc := _ffdea._fdada
	_dcedc.SetIfNotNil("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _ffdea.ExtGState)
	if _ffdea._cfcff != nil {
		_ffdea.ColorSpace = _ffdea._cfcff.ToPdfObject()
	}
	_dcedc.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _ffdea.ColorSpace)
	_dcedc.SetIfNotNil("\u0050a\u0074\u0074\u0065\u0072\u006e", _ffdea.Pattern)
	_dcedc.SetIfNotNil("\u0053h\u0061\u0064\u0069\u006e\u0067", _ffdea.Shading)
	_dcedc.SetIfNotNil("\u0058O\u0062\u006a\u0065\u0063\u0074", _ffdea.XObject)
	_dcedc.SetIfNotNil("\u0046\u006f\u006e\u0074", _ffdea.Font)
	_dcedc.SetIfNotNil("\u0050r\u006f\u0063\u0053\u0065\u0074", _ffdea.ProcSet)
	_dcedc.SetIfNotNil("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", _ffdea.Properties)
	return _dcedc
}

// GetParamsDict returns *core.PdfObjectDictionary with a set of basic image parameters.
func (_edbeac *Image) GetParamsDict() *_eb.PdfObjectDictionary {
	_gafd := _eb.MakeDict()
	_gafd.Set("\u0057\u0069\u0064t\u0068", _eb.MakeInteger(_edbeac.Width))
	_gafd.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _eb.MakeInteger(_edbeac.Height))
	_gafd.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _eb.MakeInteger(int64(_edbeac.ColorComponents)))
	_gafd.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _eb.MakeInteger(_edbeac.BitsPerComponent))
	return _gafd
}

// GetCapHeight returns the CapHeight of the font `descriptor`.
func (_dfdeg *PdfFontDescriptor) GetCapHeight() (float64, error) {
	return _eb.GetNumberAsFloat(_dfdeg.CapHeight)
}
func (_afbc fontCommon) coreString() string {
	_gdaca := ""
	if _afbc._bged != nil {
		_gdaca = _afbc._bged.String()
	}
	return _e.Sprintf("\u0025#\u0071\u0020%\u0023\u0071\u0020%\u0071\u0020\u006f\u0062\u006a\u003d\u0025d\u0020\u0054\u006f\u0055\u006e\u0069c\u006f\u0064\u0065\u003d\u0025\u0074\u0020\u0066\u006c\u0061\u0067s\u003d\u0030\u0078\u0025\u0030\u0078\u0020\u0025\u0073", _afbc._fgdee, _afbc._agcc, _afbc._bgge, _afbc._babgg, _afbc._geee != nil, _afbc.fontFlags(), _gdaca)
}

// IsRadio returns true if the button field represents a radio button, false otherwise.
func (_eceef *PdfFieldButton) IsRadio() bool { return _eceef.GetType() == ButtonTypeRadio }

type modelManager struct {
	_edcg   map[PdfModel]_eb.PdfObject
	_fgffbd map[_eb.PdfObject]PdfModel
}

// GetContainingPdfObject returns the page as a dictionary within an PdfIndirectObject.
func (_bdgc *PdfPage) GetContainingPdfObject() _eb.PdfObject { return _bdgc._efcff }
func (_baba *PdfReader) newPdfActionLaunchFromDict(_bbae *_eb.PdfObjectDictionary) (*PdfActionLaunch, error) {
	_dgd, _cfcg := _dba(_bbae.Get("\u0046"))
	if _cfcg != nil {
		return nil, _cfcg
	}
	return &PdfActionLaunch{Win: _bbae.Get("\u0057\u0069\u006e"), Mac: _bbae.Get("\u004d\u0061\u0063"), Unix: _bbae.Get("\u0055\u006e\u0069\u0078"), NewWindow: _bbae.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _dgd}, nil
}

// NewOutlineBookmark returns an initialized PdfOutlineItem for a given bookmark title and page.
func NewOutlineBookmark(title string, page *_eb.PdfIndirectObject) *PdfOutlineItem {
	_fgbgd := PdfOutlineItem{}
	_fgbgd._eeedb = &_fgbgd
	_fgbgd.Title = _eb.MakeString(title)
	_efcdf := _eb.MakeArray()
	_efcdf.Append(page)
	_efcdf.Append(_eb.MakeName("\u0046\u0069\u0074"))
	_fgbgd.Dest = _efcdf
	return &_fgbgd
}

// PdfAnnotationText represents Text annotations.
// (Section 12.5.6.4 p. 402).
type PdfAnnotationText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Open       _eb.PdfObject
	Name       _eb.PdfObject
	State      _eb.PdfObject
	StateModel _eb.PdfObject
}

func _cbcfd(_cecda *PdfField, _aadeb _eb.PdfObject) error {
	switch _cecda.GetContext().(type) {
	case *PdfFieldText:
		switch _agbaf := _aadeb.(type) {
		case *_eb.PdfObjectName:
			_eacee := _agbaf
			_ddb.Log.Debug("\u0055\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u003a\u0020\u0047\u006f\u0074 \u0056\u0020\u0061\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u003e\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f s\u0074\u0072\u0069\u006e\u0067\u0020\u0027\u0025\u0073\u0027", _eacee.String())
			_cecda.V = _eb.MakeEncodedString(_agbaf.String(), true)
		case *_eb.PdfObjectString:
			_cecda.V = _eb.MakeEncodedString(_agbaf.String(), true)
		default:
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0056\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u0028\u0025\u0023\u0076\u0029", _agbaf, _agbaf)
		}
	case *PdfFieldButton:
		switch _aadeb.(type) {
		case *_eb.PdfObjectName:
			if len(_aadeb.String()) > 0 {
				_cecda.V = _aadeb
				_faeef(_cecda, _aadeb)
			}
		case *_eb.PdfObjectString:
			if len(_aadeb.String()) > 0 {
				_cecda.V = _eb.MakeName(_aadeb.String())
				_faeef(_cecda, _cecda.V)
			}
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _cecda.PartialName(), _aadeb)
			_cecda.V = _aadeb
		}
	case *PdfFieldChoice:
		switch _aadeb.(type) {
		case *_eb.PdfObjectName:
			if len(_aadeb.String()) > 0 {
				_cecda.V = _eb.MakeString(_aadeb.String())
				_faeef(_cecda, _aadeb)
			}
		case *_eb.PdfObjectString:
			if len(_aadeb.String()) > 0 {
				_cecda.V = _aadeb
				_faeef(_cecda, _eb.MakeName(_aadeb.String()))
			}
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _cecda.PartialName(), _aadeb)
			_cecda.V = _aadeb
		}
	case *PdfFieldSignature:
		_ddb.Log.Debug("\u0054\u004f\u0044\u004f\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0061\u0070\u0070e\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0079\u0065\u0074\u003a\u0020\u0025\u0073\u002f\u0025v", _cecda.PartialName(), _aadeb)
	}
	return nil
}

var _fefac = false

// NewPdfAnnotationFileAttachment returns a new file attachment annotation.
func NewPdfAnnotationFileAttachment() *PdfAnnotationFileAttachment {
	_daa := NewPdfAnnotation()
	_deaf := &PdfAnnotationFileAttachment{}
	_deaf.PdfAnnotation = _daa
	_deaf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_daa.SetContext(_deaf)
	return _deaf
}

// Direction represents the predominant reading order for text.
type Direction string

func (_bbca *PdfReader) buildOutlineTree(_fgdfdff _eb.PdfObject, _afdg *PdfOutlineTreeNode, _gbcbb *PdfOutlineTreeNode, _acgde map[_eb.PdfObject]struct{}) (*PdfOutlineTreeNode, *PdfOutlineTreeNode, error) {
	if _acgde == nil {
		_acgde = map[_eb.PdfObject]struct{}{}
	}
	_acgde[_fgdfdff] = struct{}{}
	_abbed, _ceafgb := _fgdfdff.(*_eb.PdfIndirectObject)
	if !_ceafgb {
		return nil, nil, _e.Errorf("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0063\u006f\u006et\u0061\u0069\u006e\u0065r\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _fgdfdff)
	}
	_afbdg, _fbcgc := _abbed.PdfObject.(*_eb.PdfObjectDictionary)
	if !_fbcgc {
		return nil, nil, _dcf.New("\u006e\u006f\u0074 a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ddb.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0074\u0072\u0065\u0065\u003a\u0020d\u0069\u0063\u0074\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0076\u0029\u0020p\u003a\u0020\u0025\u0070", _afbdg, _abbed, _abbed)
	if _fadfa := _afbdg.Get("\u0054\u0069\u0074l\u0065"); _fadfa != nil {
		_agbbc, _ddee := _bbca.newPdfOutlineItemFromIndirectObject(_abbed)
		if _ddee != nil {
			return nil, nil, _ddee
		}
		_agbbc.Parent = _afdg
		_agbbc.Prev = _gbcbb
		_bagbe := _eb.ResolveReference(_afbdg.Get("\u0046\u0069\u0072s\u0074"))
		if _, _ebbf := _acgde[_bagbe]; _bagbe != nil && _bagbe != _abbed && !_ebbf {
			if !_eb.IsNullObject(_bagbe) {
				_eaabd, _bdcddb, _fcbd := _bbca.buildOutlineTree(_bagbe, &_agbbc.PdfOutlineTreeNode, nil, _acgde)
				if _fcbd != nil {
					_ddb.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0069\u0074\u0065\u006d\u0020\u0074\u0072\u0065\u0065\u003a \u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020n\u006f\u0064\u0065\u0020\u0063\u0068\u0069\u006c\u0064\u0072\u0065n\u002e", _fcbd)
				} else {
					_agbbc.First = _eaabd
					_agbbc.Last = _bdcddb
				}
			}
		}
		_cdca := _eb.ResolveReference(_afbdg.Get("\u004e\u0065\u0078\u0074"))
		if _, _dcace := _acgde[_cdca]; _cdca != nil && _cdca != _abbed && !_dcace {
			if !_eb.IsNullObject(_cdca) {
				_dadb, _addda, _fbeff := _bbca.buildOutlineTree(_cdca, _afdg, &_agbbc.PdfOutlineTreeNode, _acgde)
				if _fbeff != nil {
					_ddb.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065\u0020\u0066\u006f\u0072\u0020\u004ee\u0078\u0074\u0020\u006e\u006f\u0064\u0065\u003a\u0020\u0025\u0076\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006e\u006f\u0064e\u002e", _fbeff)
				} else {
					_agbbc.Next = _dadb
					return &_agbbc.PdfOutlineTreeNode, _addda, nil
				}
			}
		}
		return &_agbbc.PdfOutlineTreeNode, &_agbbc.PdfOutlineTreeNode, nil
	}
	_ecagb, _fgbge := _cffdc(_abbed)
	if _fgbge != nil {
		return nil, nil, _fgbge
	}
	_ecagb.Parent = _afdg
	if _edgab := _afbdg.Get("\u0046\u0069\u0072s\u0074"); _edgab != nil {
		_edgab = _eb.ResolveReference(_edgab)
		if _, _fcgfg := _acgde[_edgab]; _edgab != nil && _edgab != _abbed && !_fcgfg {
			_cabce := _eb.TraceToDirectObject(_edgab)
			if _, _gbadc := _cabce.(*_eb.PdfObjectNull); !_gbadc && _cabce != nil {
				_abcaa, _efbe, _fgcbd := _bbca.buildOutlineTree(_edgab, &_ecagb.PdfOutlineTreeNode, nil, _acgde)
				if _fgcbd != nil {
					_ddb.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020b\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0074\u0072\u0065\u0065\u003a\u0020\u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006e\u006f\u0064\u0065 \u0063\u0068i\u006c\u0064r\u0065n\u002e", _fgcbd)
				} else {
					_ecagb.First = _abcaa
					_ecagb.Last = _efbe
				}
			}
		}
	}
	return &_ecagb.PdfOutlineTreeNode, &_ecagb.PdfOutlineTreeNode, nil
}

// ToPdfObject implements interface PdfModel.
func (_gcfc *PdfAnnotationText) ToPdfObject() _eb.PdfObject {
	_gcfc.PdfAnnotation.ToPdfObject()
	_gace := _gcfc._ggf
	_agdc := _gace.PdfObject.(*_eb.PdfObjectDictionary)
	if _gcfc.PdfAnnotationMarkup != nil {
		_gcfc.PdfAnnotationMarkup.appendToPdfDictionary(_agdc)
	}
	_agdc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0054\u0065\u0078\u0074"))
	_agdc.SetIfNotNil("\u004f\u0070\u0065\u006e", _gcfc.Open)
	_agdc.SetIfNotNil("\u004e\u0061\u006d\u0065", _gcfc.Name)
	_agdc.SetIfNotNil("\u0053\u0074\u0061t\u0065", _gcfc.State)
	_agdc.SetIfNotNil("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c", _gcfc.StateModel)
	return _gace
}

var _ pdfFont = (*pdfCIDFontType0)(nil)

// PdfShadingType6 is a Coons patch mesh.
type PdfShadingType6 struct {
	*PdfShading
	BitsPerCoordinate *_eb.PdfObjectInteger
	BitsPerComponent  *_eb.PdfObjectInteger
	BitsPerFlag       *_eb.PdfObjectInteger
	Decode            *_eb.PdfObjectArray
	Function          []PdfFunction
}

// GetContentStreamWithEncoder returns the pattern cell's content stream and its encoder
func (_aegb *PdfTilingPattern) GetContentStreamWithEncoder() ([]byte, _eb.StreamEncoder, error) {
	_bagcf, _efbcc := _aegb._agddd.(*_eb.PdfObjectStream)
	if !_efbcc {
		_ddb.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _aegb._agddd)
		return nil, nil, _eb.ErrTypeError
	}
	_afebe, _abbde := _eb.DecodeStream(_bagcf)
	if _abbde != nil {
		_ddb.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020s\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0065\u0072\u0072\u003a\u0020\u0025\u0076", _abbde)
		return nil, nil, _abbde
	}
	_ccegb, _abbde := _eb.NewEncoderFromStream(_bagcf)
	if _abbde != nil {
		_ddb.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020f\u0069\u006e\u0064\u0069\u006e\u0067 \u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _abbde)
		return nil, nil, _abbde
	}
	return _afebe, _ccegb, nil
}

// NewPdfActionResetForm returns a new "reset form" action.
func NewPdfActionResetForm() *PdfActionResetForm {
	_bgg := NewPdfAction()
	_bf := &PdfActionResetForm{}
	_bf.PdfAction = _bgg
	_bgg.SetContext(_bf)
	return _bf
}

// ImageHandler interface implements common image loading and processing tasks.
// Implementing as an interface allows for the possibility to use non-standard libraries for faster
// loading and processing of images.
type ImageHandler interface {

	// Read any image type and load into a new Image object.
	Read(_ddfb _bagf.Reader) (*Image, error)

	// NewImageFromGoImage loads a NRGBA32 unidoc Image from a standard Go image structure.
	NewImageFromGoImage(_ccfa _fb.Image) (*Image, error)

	// NewGrayImageFromGoImage loads a grayscale unidoc Image from a standard Go image structure.
	NewGrayImageFromGoImage(_gdbcf _fb.Image) (*Image, error)

	// Compress an image.
	Compress(_cbffa *Image, _fbcba int64) (*Image, error)
}

// Duplex represents the paper handling option that shall be used when printing
// the file from the print dialog.
type Duplex string

func (_eebe fontCommon) asPdfObjectDictionary(_egag string) *_eb.PdfObjectDictionary {
	if _egag != "" && _eebe._fgdee != "" && _egag != _eebe._fgdee {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061\u0073\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020O\u0076\u0065\u0072\u0072\u0069\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0062t\u0079\u0070\u0065\u0020\u0074\u006f \u0025\u0023\u0071 \u0025\u0073", _egag, _eebe)
	} else if _egag == "" && _eebe._fgdee == "" {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061s\u0050\u0064\u0066Ob\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074y\u0070\u0065\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _eebe)
	} else if _eebe._fgdee == "" {
		_eebe._fgdee = _egag
	}
	_dddce := _eb.MakeDict()
	_dddce.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0046\u006f\u006e\u0074"))
	_dddce.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _eb.MakeName(_eebe._agcc))
	_dddce.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName(_eebe._fgdee))
	if _eebe._bged != nil {
		_dddce.Set("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072", _eebe._bged.ToPdfObject())
	}
	if _eebe._geee != nil {
		_dddce.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _eebe._geee)
	} else if _eebe._bgbg != nil {
		_caaa, _efaab := _eebe._bgbg.Stream()
		if _efaab != nil {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0067\u0065\u0074\u0020C\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0065r\u0072\u003d\u0025\u0076", _efaab)
		} else {
			_dddce.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _caaa)
		}
	}
	return _dddce
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_fdag *PdfColorspaceCalRGB) ToPdfObject() _eb.PdfObject {
	_dacd := &_eb.PdfObjectArray{}
	_dacd.Append(_eb.MakeName("\u0043\u0061\u006c\u0052\u0047\u0042"))
	_dccf := _eb.MakeDict()
	if _fdag.WhitePoint != nil {
		_bdag := _eb.MakeArray(_eb.MakeFloat(_fdag.WhitePoint[0]), _eb.MakeFloat(_fdag.WhitePoint[1]), _eb.MakeFloat(_fdag.WhitePoint[2]))
		_dccf.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _bdag)
	} else {
		_ddb.Log.Error("\u0043\u0061l\u0052\u0047\u0042\u003a \u004d\u0069s\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _fdag.BlackPoint != nil {
		_adfe := _eb.MakeArray(_eb.MakeFloat(_fdag.BlackPoint[0]), _eb.MakeFloat(_fdag.BlackPoint[1]), _eb.MakeFloat(_fdag.BlackPoint[2]))
		_dccf.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _adfe)
	}
	if _fdag.Gamma != nil {
		_dffb := _eb.MakeArray(_eb.MakeFloat(_fdag.Gamma[0]), _eb.MakeFloat(_fdag.Gamma[1]), _eb.MakeFloat(_fdag.Gamma[2]))
		_dccf.Set("\u0047\u0061\u006dm\u0061", _dffb)
	}
	if _fdag.Matrix != nil {
		_fded := _eb.MakeArray(_eb.MakeFloat(_fdag.Matrix[0]), _eb.MakeFloat(_fdag.Matrix[1]), _eb.MakeFloat(_fdag.Matrix[2]), _eb.MakeFloat(_fdag.Matrix[3]), _eb.MakeFloat(_fdag.Matrix[4]), _eb.MakeFloat(_fdag.Matrix[5]), _eb.MakeFloat(_fdag.Matrix[6]), _eb.MakeFloat(_fdag.Matrix[7]), _eb.MakeFloat(_fdag.Matrix[8]))
		_dccf.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _fded)
	}
	_dacd.Append(_dccf)
	if _fdag._cbfd != nil {
		_fdag._cbfd.PdfObject = _dacd
		return _fdag._cbfd
	}
	return _dacd
}
func _cbfdd(_cedf _eb.PdfObject) (map[_fc.CharCode]float64, error) {
	if _cedf == nil {
		return nil, nil
	}
	_ffaec, _dcgb := _eb.GetArray(_cedf)
	if !_dcgb {
		return nil, nil
	}
	_bdffff := map[_fc.CharCode]float64{}
	_gccbb := _ffaec.Len()
	for _ffcb := 0; _ffcb < _gccbb-1; _ffcb++ {
		_eddb := _eb.TraceToDirectObject(_ffaec.Get(_ffcb))
		_gcfbb, _cgebg := _eb.GetIntVal(_eddb)
		if !_cgebg {
			return nil, _e.Errorf("\u0042a\u0064\u0020\u0066\u006fn\u0074\u0020\u0057\u0020\u006fb\u006a0\u003a \u0069\u003d\u0025\u0064\u0020\u0025\u0023v", _ffcb, _eddb)
		}
		_ffcb++
		if _ffcb > _gccbb-1 {
			return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _ffaec)
		}
		_eaeef := _eb.TraceToDirectObject(_ffaec.Get(_ffcb))
		switch _eaeef.(type) {
		case *_eb.PdfObjectArray:
			_gead, _ := _eb.GetArray(_eaeef)
			if _dcadd, _bfdbg := _gead.ToFloat64Array(); _bfdbg == nil {
				for _ggfb := 0; _ggfb < len(_dcadd); _ggfb++ {
					_bdffff[_fc.CharCode(_gcfbb+_ggfb)] = _dcadd[_ggfb]
				}
			} else {
				return nil, _e.Errorf("\u0042\u0061\u0064 \u0066\u006f\u006e\u0074 \u0057\u0020\u0061\u0072\u0072\u0061\u0079 \u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _ffcb, _eaeef)
			}
		case *_eb.PdfObjectInteger:
			_fgbf, _eeed := _eb.GetIntVal(_eaeef)
			if !_eeed {
				return nil, _e.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _ffcb, _eaeef)
			}
			_ffcb++
			if _ffcb > _gccbb-1 {
				return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _ffaec)
			}
			_bacef := _ffaec.Get(_ffcb)
			_cbagb, _fcecga := _eb.GetNumberAsFloat(_bacef)
			if _fcecga != nil {
				return nil, _e.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0032\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _ffcb, _bacef)
			}
			for _gagbd := _gcfbb; _gagbd <= _fgbf; _gagbd++ {
				_bdffff[_fc.CharCode(_gagbd)] = _cbagb
			}
		default:
			return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057 \u006f\u0062\u006a\u0031\u0020\u0074\u0079p\u0065\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _ffcb, _eaeef)
		}
	}
	return _bdffff, nil
}

// PdfAnnotationRedact represents Redact annotations.
// (Section 12.5.6.23).
type PdfAnnotationRedact struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints  _eb.PdfObject
	IC          _eb.PdfObject
	RO          _eb.PdfObject
	OverlayText _eb.PdfObject
	Repeat      _eb.PdfObject
	DA          _eb.PdfObject
	Q           _eb.PdfObject
}

// GetAsShadingPattern returns a shading pattern. Check with IsShading() prior to using this.
func (_aadcb *PdfPattern) GetAsShadingPattern() *PdfShadingPattern {
	return _aadcb._eefgb.(*PdfShadingPattern)
}
func _fcecd(_eeeda *_eb.PdfObjectStream) (*PdfFunctionType4, error) {
	_edcee := &PdfFunctionType4{}
	_edcee._gbcd = _eeeda
	_adegd := _eeeda.PdfObjectDictionary
	_caebb, _cccdg := _eb.TraceToDirectObject(_adegd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_eb.PdfObjectArray)
	if !_cccdg {
		_ddb.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _dcf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _caebb.Len()%2 != 0 {
		_ddb.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_dfbde, _aeda := _caebb.ToFloat64Array()
	if _aeda != nil {
		return nil, _aeda
	}
	_edcee.Domain = _dfbde
	_caebb, _cccdg = _eb.TraceToDirectObject(_adegd.Get("\u0052\u0061\u006eg\u0065")).(*_eb.PdfObjectArray)
	if _cccdg {
		if _caebb.Len() < 0 || _caebb.Len()%2 != 0 {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_efbad, _efecgb := _caebb.ToFloat64Array()
		if _efecgb != nil {
			return nil, _efecgb
		}
		_edcee.Range = _efbad
	}
	_efdcf, _aeda := _eb.DecodeStream(_eeeda)
	if _aeda != nil {
		return nil, _aeda
	}
	_edcee._bdcf = _efdcf
	_ecdab := _gc.NewPSParser(_efdcf)
	_feegb, _aeda := _ecdab.Parse()
	if _aeda != nil {
		return nil, _aeda
	}
	_edcee.Program = _feegb
	return _edcee, nil
}
func (_dade *PdfReader) newPdfAnnotationHighlightFromDict(_dec *_eb.PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	_efe := PdfAnnotationHighlight{}
	_eab, _bffa := _dade.newPdfAnnotationMarkupFromDict(_dec)
	if _bffa != nil {
		return nil, _bffa
	}
	_efe.PdfAnnotationMarkup = _eab
	_efe.QuadPoints = _dec.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_efe, nil
}

// SetAnnotations sets the annotations list.
func (_ddac *PdfPage) SetAnnotations(annotations []*PdfAnnotation) { _ddac._dbga = annotations }

// SetNumCopies sets the value of the numCopies.
func (_bfceda *ViewerPreferences) SetNumCopies(numCopies int) { _bfceda._bdeec = numCopies }

// PdfColorspaceDeviceCMYK represents a CMYK32 colorspace.
type PdfColorspaceDeviceCMYK struct{}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the L, A and B components of the color.
func (_gdcfgd *PdfColorspaceLab) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_geaf, _dfbd := _eb.GetNumbersAsFloat(objects)
	if _dfbd != nil {
		return nil, _dfbd
	}
	return _gdcfgd.ColorFromFloats(_geaf)
}

// GetRuneMetrics returns the character metrics for the rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_dcfa pdfFontSimple) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	if _dcfa._gacdb != nil {
		_eefc, _eccdf := _dcfa._gacdb.Read(r)
		if _eccdf {
			return _eefc, true
		}
	}
	_gadbe := _dcfa.Encoder()
	if _gadbe == nil {
		_ddb.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _dcfa)
		return _fg.CharMetrics{}, false
	}
	_eeddc, _cddgd := _gadbe.RuneToCharcode(r)
	if !_cddgd {
		if r != ' ' {
			_ddb.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _dcfa)
		}
		return _fg.CharMetrics{}, false
	}
	_fddg, _bfcg := _dcfa.GetCharMetrics(_eeddc)
	return _fddg, _bfcg
}

// Y returns the value of the yellow component of the color.
func (_ebfc *PdfColorDeviceCMYK) Y() float64 { return _ebfc[2] }

// IsShading specifies if the pattern is a shading pattern.
func (_ccaab *PdfPattern) IsShading() bool { return _ccaab.PatternType == 2 }
func (_ccdec *PdfAcroForm) signatureFields() []*PdfFieldSignature {
	var _egdaf []*PdfFieldSignature
	for _, _eegda := range _ccdec.AllFields() {
		switch _gccca := _eegda.GetContext().(type) {
		case *PdfFieldSignature:
			_cgada := _gccca
			_egdaf = append(_egdaf, _cgada)
		}
	}
	return _egdaf
}

var _gdba = map[string]struct{}{"\u0046\u0054": {}, "\u004b\u0069\u0064\u0073": {}, "\u0054": {}, "\u0054\u0055": {}, "\u0054\u004d": {}, "\u0046\u0066": {}, "\u0056": {}, "\u0044\u0056": {}, "\u0041\u0041": {}, "\u0044\u0041": {}, "\u0051": {}, "\u0044\u0053": {}, "\u0052\u0056": {}}

// NewPdfFileSpecFromEmbeddedFile construct a new PdfFileSpec that contains an embedded file.
func NewPdfFileSpecFromEmbeddedFile(file *EmbeddedFile) *PdfFilespec {
	_bgfe := &PdfFilespec{}
	_bgfe._cbefe = _eb.MakeIndirectObject(_eb.MakeDict())
	_bgfe.Desc = _eb.MakeString(file.Description)
	_bgfe.EF = file.ToPdfObject()
	_bgfe.F = _eb.MakeString(file.Name)
	_bgfe.UF = _eb.MakeEncodedString(file.Name, true)
	_bbgf := "U\u006e\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064"
	switch file.Relationship {
	case RelationshipSource:
		_bbgf = "\u0053\u006f\u0075\u0072\u0063\u0065"
	case RelationshipData:
		_bbgf = "\u0044\u0061\u0074\u0061"
	case RelationshipAlternative:
		_bbgf = "A\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065"
	case RelationshipSupplement:
		_bbgf = "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"
	}
	_bgfe.AFRelationship = _eb.MakeName(_bbgf)
	return _bgfe
}

// AlphaMap performs mapping of alpha data for transformations. Allows custom filtering of alpha data etc.
func (_fedbg *Image) AlphaMap(mapFunc AlphaMapFunc) {
	for _abfb, _ccfbe := range _fedbg._bdcab {
		_fedbg._bdcab[_abfb] = mapFunc(_ccfbe)
	}
}

// GetNumComponents returns the number of input color components, i.e. that are input to the tint transform.
func (_ecdb *PdfColorspaceDeviceN) GetNumComponents() int { return _ecdb.ColorantNames.Len() }

// NewPdfOutlineTree returns an initialized PdfOutline tree.
func NewPdfOutlineTree() *PdfOutline {
	_afcfd := NewPdfOutline()
	_afcfd._eeedb = &_afcfd
	return _afcfd
}

// SetDecode sets the decode image float slice.
func (_gbcad *Image) SetDecode(decode []float64) { _gbcad._fedc = decode }

// ToPdfObject returns the PDF representation of the shading pattern.
func (_gdegg *PdfShadingPatternType2) ToPdfObject() _eb.PdfObject {
	_gdegg.PdfPattern.ToPdfObject()
	_efcg := _gdegg.getDict()
	if _gdegg.Shading != nil {
		_efcg.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _gdegg.Shading.ToPdfObject())
	}
	if _gdegg.Matrix != nil {
		_efcg.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _gdegg.Matrix)
	}
	if _gdegg.ExtGState != nil {
		_efcg.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _gdegg.ExtGState)
	}
	return _gdegg._agddd
}

// NewReaderForText makes a new PdfReader for an input PDF content string. For use in testing.
func NewReaderForText(txt string) *PdfReader {
	return &PdfReader{_bcefc: map[_eb.PdfObject]struct{}{}, _affaf: _fcfc(), _ebbe: _eb.NewParserFromString(txt)}
}
func _cafbb(_bbeec _eb.PdfObject) *IDTree {
	_bbeec = _eb.ResolveReference(_bbeec)
	_aeaab := _eb.MakeArray()
	_bbgaf := _eb.MakeArray()
	_aeff := []*IDTree{}
	if _gbfbe, _gacedc := _eb.GetDict(_bbeec); _gacedc {
		if _fecca := _gbfbe.Get("\u004e\u0061\u006de\u0073"); _fecca != nil {
			_fecca = _eb.ResolveReference(_fecca)
			if _bdaac, _ecca := _eb.GetArray(_fecca); _ecca {
				for _, _cgacg := range _bdaac.Elements() {
					_aeaab.Append(_cgacg)
				}
			}
		}
		if _cbab := _gbfbe.Get("\u004c\u0069\u006d\u0069\u0074\u0073"); _cbab != nil {
			_cbab = _eb.ResolveReference(_cbab)
			if _faefd, _gggad := _eb.GetArray(_cbab); _gggad {
				for _bgbcb := 0; _bgbcb < 2; _bgbcb++ {
					_bbgaf.Append(_faefd.Get(_bgbcb))
				}
			}
		}
		if _cfdfa := _gbfbe.Get("\u004b\u0069\u0064\u0073"); _cfdfa != nil {
			_cfdfa = _eb.ResolveReference(_cfdfa)
			if _bbdagg, _bfbgd := _eb.GetArray(_cfdfa); _bfbgd {
				for _, _adfef := range _bbdagg.Elements() {
					_adccga := _cafbb(_adfef)
					_aeff = append(_aeff, _adccga)
				}
			}
		}
	}
	_cfecgf := &IDTree{Names: _aeaab, Limits: _bbgaf}
	if len(_aeff) > 0 {
		_cfecgf.Kids = _aeff
	}
	return _cfecgf
}

// Compress is yet to be implemented.
// Should be able to compress in terms of JPEG quality parameter,
// and DPI threshold (need to know bounding area dimensions).
func (_efbc DefaultImageHandler) Compress(input *Image, quality int64) (*Image, error) {
	return input, nil
}

// NewPdfAnnotationCircle returns a new circle annotation.
func NewPdfAnnotationCircle() *PdfAnnotationCircle {
	_gcc := NewPdfAnnotation()
	_fdb := &PdfAnnotationCircle{}
	_fdb.PdfAnnotation = _gcc
	_fdb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gcc.SetContext(_fdb)
	return _fdb
}

// GetNameDictionary returns the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_gdfeac *PdfReader) GetNameDictionary() (_eb.PdfObject, error) {
	_bdcdc := _eb.ResolveReference(_gdfeac._bagcfd.Get("\u004e\u0061\u006de\u0073"))
	if _bdcdc == nil {
		return nil, nil
	}
	if !_gdfeac._cfcgdf {
		_acbac := _gdfeac.traverseObjectData(_bdcdc)
		if _acbac != nil {
			return nil, _acbac
		}
	}
	return _bdcdc, nil
}
func _cgag(_afaef *fontCommon) *pdfCIDFontType0 { return &pdfCIDFontType0{fontCommon: *_afaef} }

// ConvertToBinary converts current image into binary (bi-level) format.
// Binary images are composed of single bits per pixel (only black or white).
// If provided image has more color components, then it would be converted into binary image using
// histogram auto threshold function.
func (_dfade *Image) ConvertToBinary() error {
	if _dfade.ColorComponents == 1 && _dfade.BitsPerComponent == 1 {
		return nil
	}
	_bdaead, _cgac := _dfade.ToGoImage()
	if _cgac != nil {
		return _cgac
	}
	_febg, _cgac := _df.MonochromeConverter.Convert(_bdaead)
	if _cgac != nil {
		return _cgac
	}
	_dfade.Data = _febg.Base().Data
	_dfade._bdcab, _cgac = _df.ScaleAlphaToMonochrome(_dfade._bdcab, int(_dfade.Width), int(_dfade.Height))
	if _cgac != nil {
		return _cgac
	}
	_dfade.BitsPerComponent = 1
	_dfade.ColorComponents = 1
	_dfade._fedc = nil
	return nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_ebagd *PdfPageResourcesColorspaces) ToPdfObject() _eb.PdfObject {
	_eebad := _eb.MakeDict()
	for _, _abff := range _ebagd.Names {
		_eebad.Set(_eb.PdfObjectName(_abff), _ebagd.Colorspaces[_abff].ToPdfObject())
	}
	if _ebagd._ccba != nil {
		_ebagd._ccba.PdfObject = _eebad
		return _ebagd._ccba
	}
	return _eebad
}
func _gaffc() _d.Time { _dfbafc.Lock(); defer _dfbafc.Unlock(); return _ebffa }

// PdfShading represents a shading dictionary. There are 7 types of shading,
// indicatedby the shading type variable:
// 1: Function-based shading.
// 2: Axial shading.
// 3: Radial shading.
// 4: Free-form Gouraud-shaded triangle mesh.
// 5: Lattice-form Gouraud-shaded triangle mesh.
// 6: Coons patch mesh.
// 7: Tensor-product patch mesh.
// types 4-7 are contained in a stream object, where the dictionary is given by the stream dictionary.
type PdfShading struct {
	ShadingType *_eb.PdfObjectInteger
	ColorSpace  PdfColorspace
	Background  *_eb.PdfObjectArray
	BBox        *PdfRectangle
	AntiAlias   *_eb.PdfObjectBool
	_ecffg      PdfModel
	_cefaa      _eb.PdfObject
}

// ToPdfObject returns colorspace in a PDF object format [name stream]
func (_dggb *PdfColorspaceICCBased) ToPdfObject() _eb.PdfObject {
	_cfgb := &_eb.PdfObjectArray{}
	_cfgb.Append(_eb.MakeName("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"))
	var _cfcb *_eb.PdfObjectStream
	if _dggb._cdebc != nil {
		_cfcb = _dggb._cdebc
	} else {
		_cfcb = &_eb.PdfObjectStream{}
	}
	_aaedd := _eb.MakeDict()
	_aaedd.Set("\u004e", _eb.MakeInteger(int64(_dggb.N)))
	if _dggb.Alternate != nil {
		_aaedd.Set("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e", _dggb.Alternate.ToPdfObject())
	}
	if _dggb.Metadata != nil {
		_aaedd.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _dggb.Metadata)
	}
	if _dggb.Range != nil {
		var _cefb []_eb.PdfObject
		for _, _fccb := range _dggb.Range {
			_cefb = append(_cefb, _eb.MakeFloat(_fccb))
		}
		_aaedd.Set("\u0052\u0061\u006eg\u0065", _eb.MakeArray(_cefb...))
	}
	_aaedd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_dggb.Data))))
	_cfcb.Stream = _dggb.Data
	_cfcb.PdfObjectDictionary = _aaedd
	_cfgb.Append(_cfcb)
	if _dggb._ebggg != nil {
		_dggb._ebggg.PdfObject = _cfgb
		return _dggb._ebggg
	}
	return _cfgb
}
func (_bebc *PdfReader) newPdfActionGotoRFromDict(_bfc *_eb.PdfObjectDictionary) (*PdfActionGoToR, error) {
	_ege, _dfdg := _dba(_bfc.Get("\u0046"))
	if _dfdg != nil {
		return nil, _dfdg
	}
	return &PdfActionGoToR{D: _bfc.Get("\u0044"), NewWindow: _bfc.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _ege}, nil
}

var _aefb = _ab.MustCompile("\u005b\\\u006e\u005c\u0072\u005d\u002b")

// CharcodeBytesToUnicode converts PDF character codes `data` to a Go unicode string.
//
// 9.10 Extraction of Text Content (page 292)
// The process of finding glyph descriptions in OpenType fonts by a conforming reader shall be the following:
//   - For Type 1 fonts using “CFF” tables, the process shall be as described in 9.6.6.2, "Encodings
//     for Type 1 Fonts".
//   - For TrueType fonts using “glyf” tables, the process shall be as described in 9.6.6.4,
//     "Encodings for TrueType Fonts". Since this process sometimes produces ambiguous results,
//     conforming writers, instead of using a simple font, shall use a Type 0 font with an Identity-H
//     encoding and use the glyph indices as character codes, as described following Table 118.
func (_bdce *PdfFont) CharcodeBytesToUnicode(data []byte) (string, int, int) {
	_gcgg, _, _aedgc := _bdce.CharcodesToUnicodeWithStats(_bdce.BytesToCharcodes(data))
	_dcdgd := _fc.ExpandLigatures(_gcgg)
	return _dcdgd, _ae.RuneCountInString(_dcdgd), _aedgc
}

// NewPdfTransformParamsDocMDP create a PdfTransformParamsDocMDP with the specific permissions.
func NewPdfTransformParamsDocMDP(permission _bab.DocMDPPermission) *PdfTransformParamsDocMDP {
	return &PdfTransformParamsDocMDP{Type: _eb.MakeName("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"), P: _eb.MakeInteger(int64(permission)), V: _eb.MakeName("\u0031\u002e\u0032")}
}

// ToPdfObject converts the ID tree to a PDF object.
func (_faace *IDTree) ToPdfObject() _eb.PdfObject {
	_bgadcg := _eb.MakeDict()
	if _faace.Names != nil && _faace.Names.Len() > 0 {
		_bgadcg.Set("\u004e\u0061\u006de\u0073", _faace.Names)
		_bgadcg.Set("\u004c\u0069\u006d\u0069\u0074\u0073", _faace.Limits)
	}
	if len(_faace.Kids) > 0 {
		_cbgee := _eb.MakeArray()
		for _, _gfbdb := range _faace.Kids {
			_cbgee.Append(_gfbdb.ToPdfObject())
		}
		_gcddf := _eb.MakeDict()
		_gcddf.Set("\u004b\u0069\u0064\u0073", _cbgee)
		_bgadcg.Set("\u004b\u0069\u0064\u0073", _gcddf)
	}
	return _bgadcg
}

// GetXObjectImageByName returns the XObjectImage with the specified name from the
// page resources, if it exists.
func (_cggg *PdfPageResources) GetXObjectImageByName(keyName _eb.PdfObjectName) (*XObjectImage, error) {
	_bfcc, _dcaab := _cggg.GetXObjectByName(keyName)
	if _bfcc == nil {
		return nil, nil
	}
	if _dcaab != XObjectTypeImage {
		return nil, _dcf.New("\u006e\u006f\u0074 \u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065")
	}
	_dagadd, _gddbf := NewXObjectImageFromStream(_bfcc)
	if _gddbf != nil {
		return nil, _gddbf
	}
	return _dagadd, nil
}

// NewPdfActionJavaScript returns a new "javaScript" action.
func NewPdfActionJavaScript() *PdfActionJavaScript {
	_afe := NewPdfAction()
	_bc := &PdfActionJavaScript{}
	_bc.PdfAction = _afe
	_afe.SetContext(_bc)
	return _bc
}

// ColorToRGB converts a CalRGB color to an RGB color.
func (_dfcc *PdfColorspaceCalRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_deded, _edea := color.(*PdfColorCalRGB)
	if !_edea {
		_ddb.Log.Debug("\u0049\u006e\u0070ut\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0063\u0061\u006c\u0020\u0072\u0067\u0062")
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_agfb := _deded.A()
	_gcca := _deded.B()
	_fcbf := _deded.C()
	X := _dfcc.Matrix[0]*_gg.Pow(_agfb, _dfcc.Gamma[0]) + _dfcc.Matrix[3]*_gg.Pow(_gcca, _dfcc.Gamma[1]) + _dfcc.Matrix[6]*_gg.Pow(_fcbf, _dfcc.Gamma[2])
	Y := _dfcc.Matrix[1]*_gg.Pow(_agfb, _dfcc.Gamma[0]) + _dfcc.Matrix[4]*_gg.Pow(_gcca, _dfcc.Gamma[1]) + _dfcc.Matrix[7]*_gg.Pow(_fcbf, _dfcc.Gamma[2])
	Z := _dfcc.Matrix[2]*_gg.Pow(_agfb, _dfcc.Gamma[0]) + _dfcc.Matrix[5]*_gg.Pow(_gcca, _dfcc.Gamma[1]) + _dfcc.Matrix[8]*_gg.Pow(_fcbf, _dfcc.Gamma[2])
	_ecbc := 3.240479*X + -1.537150*Y + -0.498535*Z
	_agbag := -0.969256*X + 1.875992*Y + 0.041556*Z
	_ccfga := 0.055648*X + -0.204043*Y + 1.057311*Z
	_ecbc = _gg.Min(_gg.Max(_ecbc, 0), 1.0)
	_agbag = _gg.Min(_gg.Max(_agbag, 0), 1.0)
	_ccfga = _gg.Min(_gg.Max(_ccfga, 0), 1.0)
	return NewPdfColorDeviceRGB(_ecbc, _agbag, _ccfga), nil
}

// PdfAnnotationMovie represents Movie annotations.
// (Section 12.5.6.17).
type PdfAnnotationMovie struct {
	*PdfAnnotation
	T     _eb.PdfObject
	Movie _eb.PdfObject
	A     _eb.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_efef *PdfAnnotationFreeText) ToPdfObject() _eb.PdfObject {
	_efef.PdfAnnotation.ToPdfObject()
	_bfdb := _efef._ggf
	_aead := _bfdb.PdfObject.(*_eb.PdfObjectDictionary)
	_efef.PdfAnnotationMarkup.appendToPdfDictionary(_aead)
	_aead.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074"))
	_aead.SetIfNotNil("\u0044\u0041", _efef.DA)
	_aead.SetIfNotNil("\u0051", _efef.Q)
	_aead.SetIfNotNil("\u0052\u0043", _efef.RC)
	_aead.SetIfNotNil("\u0044\u0053", _efef.DS)
	_aead.SetIfNotNil("\u0043\u004c", _efef.CL)
	_aead.SetIfNotNil("\u0049\u0054", _efef.IT)
	_aead.SetIfNotNil("\u0042\u0045", _efef.BE)
	_aead.SetIfNotNil("\u0052\u0044", _efef.RD)
	_aead.SetIfNotNil("\u0042\u0053", _efef.BS)
	_aead.SetIfNotNil("\u004c\u0045", _efef.LE)
	return _bfdb
}

// ToPdfObject implements interface PdfModel.
func (_bbd *PdfActionNamed) ToPdfObject() _eb.PdfObject {
	_bbd.PdfAction.ToPdfObject()
	_gdb := _bbd._dee
	_eeb := _gdb.PdfObject.(*_eb.PdfObjectDictionary)
	_eeb.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeNamed)))
	_eeb.SetIfNotNil("\u004e", _bbd.N)
	return _gdb
}

// SetDate sets the `M` field of the signature.
func (_adbgc *PdfSignature) SetDate(date _d.Time, format string) {
	if format == "" {
		format = "\u0044\u003a\u003200\u0036\u0030\u0031\u0030\u0032\u0031\u0035\u0030\u0034\u0030\u0035\u002d\u0030\u0037\u0027\u0030\u0030\u0027"
	}
	_adbgc.M = _eb.MakeString(date.Format(format))
}

// NewEmbeddedFileFromContent construct a new EmbeddedFile from supplied file content.
func NewEmbeddedFileFromContent(content []byte) (*EmbeddedFile, error) {
	_acgb := _egg.Detect(content)
	_eceb := _c.Sum(content)
	_ggbb := &EmbeddedFile{Name: "\u0061\u0074\u0074\u0061\u0063\u0068\u006d\u0065\u006e\u0074", Content: content, FileType: _acgb.String(), Hash: _egf.EncodeToString(_eceb[:])}
	return _ggbb, nil
}

// DecodeArray returns an empty slice as there are no components associated with pattern colorspace.
func (_gffga *PdfColorspaceSpecialPattern) DecodeArray() []float64 { return []float64{} }

// SetFitWindow sets the value of the fitWindow flag.
func (_bcddaa *ViewerPreferences) SetFitWindow(fitWindow bool) { _bcddaa._edbaa = &fitWindow }

// SetPdfSubject sets the Subject attribute of the output PDF.
func SetPdfSubject(subject string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _bgabb = subject }

// ToPdfObject returns a PDF object representation of the outline item.
func (_ceegb *OutlineItem) ToPdfObject() _eb.PdfObject {
	_fddca, _ := _ceegb.ToPdfOutlineItem()
	return _fddca.ToPdfObject()
}

// GetPdfVersion gets the version of the PDF used within this document.
func (_ddedgd *PdfWriter) GetPdfVersion() string { return _ddedgd.getPdfVersion() }
func (_cccab *PdfWriter) setCatalogVersion() {
	_cccab._dbffa.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _eb.MakeName(_e.Sprintf("\u0025\u0064\u002e%\u0064", _cccab._edbbf.Major, _cccab._edbbf.Minor)))
}

// ToPdfObject converts the structure tree root to a PDF object.
func (_afdb *StructTreeRoot) ToPdfObject() _eb.PdfObject {
	_fegfb := _afdb._fggdf
	if _fegfb == nil {
		_fegfb = &_eb.PdfIndirectObject{}
		_fegfb.PdfObject = _eb.MakeDict()
	}
	_eefcag := _fegfb.PdfObject.(*_eb.PdfObjectDictionary)
	var _fgfga _eb.PdfObject
	if len(_afdb.K) == 1 {
		_fgfga = _eb.MakeIndirectObject(_afdb.K[0].ToPdfObject())
	} else {
		_cfcffa := _eb.MakeArray()
		for _, K := range _afdb.K {
			_cfcffa.Append(_eb.MakeIndirectObject(K.ToPdfObject()))
		}
		_fgfga = _cfcffa
	}
	var (
		_afab   = []_eb.PdfObject{}
		_adgbfc = map[_eb.PdfObject][]_eb.PdfObject{}
		_edbcf  = map[string]_eb.PdfObject{}
	)
	_bfecg(_fgfga, _fegfb, _adgbfc, _edbcf, &_afab)
	_eefcag.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	_eefcag.Set("\u004b", _fgfga)
	if _afdb.IDTree != nil {
		_eefcag.Set("\u0049\u0044\u0054\u0072\u0065\u0065", _eb.MakeIndirectObject(_afdb.IDTree.ToPdfObject()))
	} else if len(_edbcf) > 0 {
		_eaeefe := _eb.MakeArray()
		_bafec := make([]string, 0, len(_edbcf))
		for _eedbf := range _edbcf {
			_bafec = append(_bafec, _eedbf)
		}
		_ba.Strings(_bafec)
		for _, _dfccb := range _bafec {
			_eaeefe.Append(_eb.MakeString(_dfccb))
			_eaeefe.Append(_edbcf[_dfccb])
		}
		_afdb.IDTree = &IDTree{Names: _eaeefe, Limits: _eb.MakeArray(_eb.MakeString(_bafec[0]), _eb.MakeString(_bafec[len(_bafec)-1]))}
		_eefcag.Set("\u0049\u0044\u0054\u0072\u0065\u0065", _eb.MakeIndirectObject(_afdb.IDTree.ToPdfObject()))
	}
	if _afdb.ParentTree != nil {
		_eefcag.Set("\u0050\u0061\u0072\u0065\u006e\u0074\u0054\u0072\u0065\u0065", _eb.MakeIndirectObject(_afdb.ParentTree))
	} else if len(_adgbfc) > 0 || len(_afab) > 0 {
		_ccebf := _eb.MakeArray()
		_fbeg := 0
		for _aeecc, _fcafa := range _adgbfc {
			_bfeg := _eb.MakeArray()
			for _, _ebgaaa := range _fcafa {
				_bfeg.Append(_ebgaaa)
			}
			_aaaaf := _eb.MakeInteger(int64(_fbeg))
			_ccebf.Append(_aaaaf)
			_ccebf.Append(_eb.MakeIndirectObject(_bfeg))
			if _ddbfc, _ebeg := _eb.GetIndirect(_aeecc); _ebeg {
				if _ccaff, _febba := _eb.GetDict(_ddbfc.PdfObject); _febba {
					_bdcee := _ccaff.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
					if _bdcee != nil {
						if _bgegb, _daced := _eb.GetIntVal(_bdcee); _daced {
							if _fbeg < _bgegb {
								_ccaff.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _aaaaf)
							}
						}
					} else {
						_ccaff.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _aaaaf)
					}
				}
			}
			_fbeg++
		}
		_aedde := func(_fdca _eb.PdfObject, _ggdfb *_eb.PdfObjectDictionary) bool {
			if _fcad := _ggdfb.Get("\u004f\u0062\u006a"); _fcad != nil {
				if _abaad, _bdgeb := _eb.GetDict(_fcad); _bdgeb {
					_cagga := _abaad.Get("\u0054\u0079\u0070\u0065")
					_ebgggb := _abaad.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
					if _cagga != nil && _cagga.String() == "\u0041\u006e\u006eo\u0074" && _ebgggb != nil && _ebgggb.String() == "\u004c\u0069\u006e\u006b" {
						_eafe := _eb.MakeInteger(int64(_fbeg))
						_abaad.Set("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _eafe)
						_ccebf.Append(_eafe)
						_ccebf.Append(_fdca)
						_fbeg++
						return true
					}
				}
			}
			return false
		}
		for _, _cefca := range _afab {
			if _eaeed, _gfegc := _eb.GetDict(_cefca); _gfegc {
				if _gaag := _eaeed.Get("\u0053"); _gaag != nil {
					if _ddbdb, _eecbc := _eb.GetNameVal(_gaag); _eecbc && _ddbdb == StructureTypeLink {
						if _aadfef := _eaeed.Get("\u004b"); _aadfef != nil {
							if _facbfd, _dgfcc := _eb.GetArray(_aadfef); _dgfcc {
								for _, _cfgbec := range _facbfd.Elements() {
									if _bgdbb, _bcgdb := _eb.GetDict(_cfgbec); _bcgdb {
										_aedde(_cefca, _bgdbb)
									}
								}
							} else if _baab, _bgff := _eb.GetDict(_aadfef); _bgff {
								_aedde(_cefca, _baab)
							}
						}
					}
				}
			}
		}
		_afdb.ParentTree = _eb.MakeDict()
		_afdb.ParentTree.Set("\u004e\u0075\u006d\u0073", _ccebf)
		_afdb.ParentTreeNextKey = int64(_fbeg)
		_eefcag.Set("\u0050\u0061\u0072\u0065\u006e\u0074\u0054\u0072\u0065\u0065", _eb.MakeIndirectObject(_afdb.ParentTree))
	}
	_eefcag.Set("\u0050\u0061\u0072\u0065\u006e\u0074\u0054\u0072\u0065\u0065\u004e\u0065x\u0074\u004b\u0065\u0079", _eb.MakeInteger(_afdb.ParentTreeNextKey))
	if _afdb.RoleMap != nil {
		_eefcag.Set("\u0052o\u006c\u0065\u004d\u0061\u0070", _afdb.RoleMap)
	}
	if _afdb.ClassMap != nil {
		_eefcag.Set("\u0043\u006c\u0061\u0073\u0073\u004d\u0061\u0070", _afdb.ClassMap)
	}
	return _fegfb
}

// SetRefObject sets the reference object for the KValue.
func (_bbcbd *KValue) SetRefObject(refObject _eb.PdfObject) {
	_bbcbd.Clear()
	_bbcbd._fbgaa = refObject
}

// Write writes out the PDF.
func (_bceff *PdfWriter) Write(writer _bagf.Writer) error {
	_ddb.Log.Trace("\u0057r\u0069\u0074\u0065\u0028\u0029")
	if _ccbce, _fdcd := writer.(*_ccb.File); _fdcd {
		_bceff.SetFileName(_ccbce.Name())
	}
	_bdgae := _bceff.checkLicense()
	if _bdgae != nil {
		return _bdgae
	}
	if _bdgae = _bceff.writeOutlines(); _bdgae != nil {
		return _bdgae
	}
	if _bdgae = _bceff.writeAcroFormFields(); _bdgae != nil {
		return _bdgae
	}
	if _bdgae = _bceff.writeNamesDictionary(); _bdgae != nil {
		return _bdgae
	}
	_bceff.checkPendingObjects()
	if _bdgae = _bceff.writeOutputIntents(); _bdgae != nil {
		return _bdgae
	}
	_bceff.setCatalogVersion()
	_bceff.copyObjects()
	if _bdgae = _bceff.optimize(); _bdgae != nil {
		return _bdgae
	}
	if _bdgae = _bceff.optimizeDocument(); _bdgae != nil {
		return _bdgae
	}
	var _ccfea _ecd.Hash
	if _bceff._beeaf {
		_ccfea = _c.New()
		writer = _bagf.MultiWriter(_ccfea, writer)
	}
	_bceff.setWriter(writer)
	_gccab := _bceff.checkCrossReferenceStream()
	_bcgeg, _gccab := _bceff.mapObjectStreams(_gccab)
	_bceff.adjustXRefAffectedVersion(_gccab)
	_bceff.writeDocumentVersion()
	_bceff.updateObjectNumbers()
	_bceff.writeObjects()
	if _bdgae = _bceff.writeObjectsInStreams(_bcgeg); _bdgae != nil {
		return _bdgae
	}
	_gbgde := _bceff._dfabe
	var _ccagb int
	for _ecbac := range _bceff._aggfdb {
		if _ecbac > _ccagb {
			_ccagb = _ecbac
		}
	}
	if _bceff._beeaf {
		if _bdgae = _bceff.setHashIDs(_ccfea); _bdgae != nil {
			return _bdgae
		}
	}
	if _gccab {
		if _bdgae = _bceff.writeXRefStreams(_ccagb, _gbgde); _bdgae != nil {
			return _bdgae
		}
	} else {
		_bceff.writeTrailer(_ccagb)
	}
	_bceff.makeOffSetReference(_gbgde)
	if _bdgae = _bceff.flushWriter(); _bdgae != nil {
		return _bdgae
	}
	return nil
}
func (_bceea *PdfReader) newPdfFieldSignatureFromDict(_fdcg *_eb.PdfObjectDictionary) (*PdfFieldSignature, error) {
	_daece := &PdfFieldSignature{}
	_cggf, _agdff := _eb.GetIndirect(_fdcg.Get("\u0056"))
	if _agdff {
		var _eaebd error
		_daece.V, _eaebd = _bceea.newPdfSignatureFromIndirect(_cggf)
		if _eaebd != nil {
			return nil, _eaebd
		}
	}
	_daece.Lock, _ = _eb.GetIndirect(_fdcg.Get("\u004c\u006f\u0063\u006b"))
	_daece.SV, _ = _eb.GetIndirect(_fdcg.Get("\u0053\u0056"))
	return _daece, nil
}

// NewXObjectImageFromStream builds the image xobject from a stream object.
// An image dictionary is the dictionary portion of a stream object representing an image XObject.
func NewXObjectImageFromStream(stream *_eb.PdfObjectStream) (*XObjectImage, error) {
	_agedd := &XObjectImage{}
	_agedd._gceffg = stream
	_geecc := *(stream.PdfObjectDictionary)
	_aefbee, _addfe := _eb.NewEncoderFromStream(stream)
	if _addfe != nil {
		return nil, _addfe
	}
	_agedd.Filter = _aefbee
	if _fddbf := _eb.TraceToDirectObject(_geecc.Get("\u0057\u0069\u0064t\u0068")); _fddbf != nil {
		_dcbfg, _fbea := _fddbf.(*_eb.PdfObjectInteger)
		if !_fbea {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_ceeag := int64(*_dcbfg)
		_agedd.Width = &_ceeag
	} else {
		return nil, _dcf.New("\u0077\u0069\u0064\u0074\u0068\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _cbaffc := _eb.TraceToDirectObject(_geecc.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _cbaffc != nil {
		_bccfb, _adab := _cbaffc.(*_eb.PdfObjectInteger)
		if !_adab {
			return nil, _dcf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_cgddd := int64(*_bccfb)
		_agedd.Height = &_cgddd
	} else {
		return nil, _dcf.New("\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	if _daee := _eb.TraceToDirectObject(_geecc.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _daee != nil {
		_bffcb, _dffbb := NewPdfColorspaceFromPdfObject(_daee)
		if _dffbb != nil {
			return nil, _dffbb
		}
		_agedd.ColorSpace = _bffcb
	} else {
		_ddb.Log.Debug("\u0058O\u0062\u006a\u0065c\u0074\u0020\u0049m\u0061ge\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067 1\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u002d\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047r\u0061\u0079")
		_agedd.ColorSpace = NewPdfColorspaceDeviceGray()
	}
	if _fbgaef := _eb.TraceToDirectObject(_geecc.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _fbgaef != nil {
		_cbdcb, _dddad := _fbgaef.(*_eb.PdfObjectInteger)
		if !_dddad {
			return nil, _dcf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_bcfbe := int64(*_cbdcb)
		_agedd.BitsPerComponent = &_bcfbe
	}
	_agedd.Intent = _geecc.Get("\u0049\u006e\u0074\u0065\u006e\u0074")
	_agedd.ImageMask = _geecc.Get("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k")
	_agedd.Mask = _geecc.Get("\u004d\u0061\u0073\u006b")
	_agedd.Decode = _geecc.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	_agedd.Interpolate = _geecc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065")
	_agedd.Alternatives = _geecc.Get("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073")
	_agedd.SMask = _geecc.Get("\u0053\u004d\u0061s\u006b")
	_agedd.SMaskInData = _geecc.Get("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061")
	_agedd.Matte = _geecc.Get("\u004d\u0061\u0074t\u0065")
	_agedd.Name = _geecc.Get("\u004e\u0061\u006d\u0065")
	_agedd.StructParent = _geecc.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_agedd.ID = _geecc.Get("\u0049\u0044")
	_agedd.OPI = _geecc.Get("\u004f\u0050\u0049")
	_agedd.Metadata = _geecc.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	_agedd.OC = _geecc.Get("\u004f\u0043")
	_agedd.Stream = stream.Stream
	return _agedd, nil
}

// SetPdfModifiedDate sets the ModDate attribute of the output PDF.
func SetPdfModifiedDate(modifiedDate _d.Time) {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	_eegaeb = modifiedDate
}
func (_egea *PdfReader) newPdfActionThreadFromDict(_baf *_eb.PdfObjectDictionary) (*PdfActionThread, error) {
	_egc, _add := _dba(_baf.Get("\u0046"))
	if _add != nil {
		return nil, _add
	}
	return &PdfActionThread{D: _baf.Get("\u0044"), B: _baf.Get("\u0042"), F: _egc}, nil
}

// NewPdfColorspaceICCBased returns a new ICCBased colorspace object.
func NewPdfColorspaceICCBased(N int) (*PdfColorspaceICCBased, error) {
	_eabe := &PdfColorspaceICCBased{}
	if N != 1 && N != 3 && N != 4 {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u0031/\u0033\u002f\u0034\u0029")
	}
	_eabe.N = N
	return _eabe, nil
}

// A returns the value of the A component of the color.
func (_cfeg *PdfColorLab) A() float64 { return _cfeg[1] }

// SetTabOrder sets the tab order for the page.
func (_efdab *PdfPage) SetTabOrder(order TabOrderType) { _efdab.Tabs = _eb.MakeName(string(order)) }
func (_fgedf *PdfWriter) makeOffSetReference(_afdcg int64) {
	_bbaaf := _e.Sprintf("\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u000a\u0025\u0064\u000a", _afdcg)
	_fgedf.writeString(_bbaaf)
	_fgedf.writeString("\u0025\u0025\u0045\u004f\u0046\u000a")
}

// GetContainingPdfObject returns the container of the PdfAcroForm (indirect object).
func (_dccdbb *PdfAcroForm) GetContainingPdfObject() _eb.PdfObject { return _dccdbb._fbgad }

// GetContainingPdfObject returns the container of the DSS (indirect object).
func (_bdcc *DSS) GetContainingPdfObject() _eb.PdfObject { return _bdcc._fcdgc }

// GetChildren returns the children of the K dictionary object.
func (_feegd *KDict) GetChildren() []*KValue { return _feegd._eegge }
func (_bbaed *LTV) enable(_ecdbb, _ffbdg []*_bag.Certificate, _agfcda string) error {
	_cceb, _caadb, _cgfbg := _bbaed.buildCertChain(_ecdbb, _ffbdg)
	if _cgfbg != nil {
		return _cgfbg
	}
	_cbbc, _cgfbg := _bbaed.getCerts(_cceb)
	if _cgfbg != nil {
		return _cgfbg
	}
	var _bcacg, _agfbde [][]byte
	if _bbaed.OCSPClient != nil {
		_bcacg, _cgfbg = _bbaed.getOCSPs(_cceb, _caadb)
		if _cgfbg != nil {
			return _cgfbg
		}
	}
	if _bbaed.CRLClient != nil {
		_agfbde, _cgfbg = _bbaed.getCRLs(_cceb)
		if _cgfbg != nil {
			return _cgfbg
		}
	}
	_adfg := _bbaed._fadg
	_aafcf, _cgfbg := _adfg.AddCerts(_cbbc)
	if _cgfbg != nil {
		return _cgfbg
	}
	_bdba, _cgfbg := _adfg.AddOCSPs(_bcacg)
	if _cgfbg != nil {
		return _cgfbg
	}
	_gbec, _cgfbg := _adfg.AddCRLs(_agfbde)
	if _cgfbg != nil {
		return _cgfbg
	}
	if _agfcda != "" {
		_adfg.VRI[_agfcda] = &VRI{Cert: _aafcf, OCSP: _bdba, CRL: _gbec}
	}
	_bbaed._feba.SetDSS(_adfg)
	return nil
}
func (_edbea *PdfAcroForm) fill(_beaac FieldValueProvider, _dcfca FieldAppearanceGenerator) error {
	if _edbea == nil {
		return nil
	}
	_dbadc, _abaed := _beaac.FieldValues()
	if _abaed != nil {
		return _abaed
	}
	for _, _aggbe := range _edbea.AllFields() {
		_cdabc := _aggbe.PartialName()
		_befcb, _bbbe := _dbadc[_cdabc]
		if !_bbbe {
			if _bdbf, _acff := _aggbe.FullName(); _acff == nil {
				_befcb, _bbbe = _dbadc[_bdbf]
			}
		}
		if !_bbbe {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _cdabc)
			continue
		}
		if _bfae := _cbcfd(_aggbe, _befcb); _bfae != nil {
			return _bfae
		}
		if _dcfca == nil {
			continue
		}
		for _, _gcagf := range _aggbe.Annotations {
			_daabc, _fdbfd := _dcfca.GenerateAppearanceDict(_edbea, _aggbe, _gcagf)
			if _fdbfd != nil {
				return _fdbfd
			}
			_gcagf.AP = _daabc
			_gcagf.ToPdfObject()
		}
	}
	return nil
}

// NewPdfAppender creates a new Pdf appender from a Pdf reader.
func NewPdfAppender(reader *PdfReader) (*PdfAppender, error) {
	_cgba := &PdfAppender{_bccd: reader._cbeg, Reader: reader, _ebcb: reader._ebbe, _fdfg: reader._bcefc}
	_gaea, _ffb := _cgba._bccd.Seek(0, _bagf.SeekEnd)
	if _ffb != nil {
		return nil, _ffb
	}
	_cgba._eabd = _gaea
	if _, _ffb = _cgba._bccd.Seek(0, _bagf.SeekStart); _ffb != nil {
		return nil, _ffb
	}
	_cgba._adac, _ffb = NewPdfReader(_cgba._bccd)
	if _ffb != nil {
		return nil, _ffb
	}
	for _, _begf := range _cgba.Reader.GetObjectNums() {
		if _cgba._cfgdf < _begf {
			_cgba._cfgdf = _begf
		}
	}
	_cgba._dce = _cgba._ebcb.GetXrefTable()
	_cgba._dgc = _cgba._ebcb.GetXrefOffset()
	_cgba._fedg = append(_cgba._fedg, _cgba._adac.PageList...)
	_cgba._accfg = make(map[_eb.PdfObject]struct{})
	_cgba._ebcbc = make(map[_eb.PdfObject]int64)
	_cgba._dbae = make(map[_eb.PdfObject]struct{})
	_cgba._bddda = _cgba._adac.AcroForm
	_cgba._cedc = _cgba._adac.DSS
	return _cgba, nil
}

// ParsePdfObject parses input pdf object into given output intent.
func (_abcg *PdfOutputIntent) ParsePdfObject(object _eb.PdfObject) error {
	_ccbe, _ebdec := _eb.GetDict(object)
	if !_ebdec {
		_ddb.Log.Error("\u0055\u006e\u006bno\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020%\u0054 \u0066o\u0072 \u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074", object)
		return _dcf.New("\u0075\u006e\u006b\u006e\u006fw\u006e\u0020\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074")
	}
	_abcg._ebbd = _ccbe
	_abcg.Type, _ = _ccbe.GetString("\u0054\u0079\u0070\u0065")
	_fecgf, _ebdec := _ccbe.GetString("\u0053")
	if _ebdec {
		switch _fecgf {
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411":
			_abcg.S = PdfOutputIntentTypeA1
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00412":
			_abcg.S = PdfOutputIntentTypeA2
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00413":
			_abcg.S = PdfOutputIntentTypeA3
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00414":
			_abcg.S = PdfOutputIntentTypeA4
		case "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058":
			_abcg.S = PdfOutputIntentTypeX
		}
	}
	_abcg.OutputCondition, _ = _ccbe.GetString("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e")
	_abcg.OutputConditionIdentifier, _ = _ccbe.GetString("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r")
	_abcg.RegistryName, _ = _ccbe.GetString("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065")
	_abcg.Info, _ = _ccbe.GetString("\u0049\u006e\u0066\u006f")
	if _fgaff, _egde := _eb.GetStream(_ccbe.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _egde {
		_abcg.ColorComponents, _ = _eb.GetIntVal(_fgaff.Get("\u004e"))
		_fgagf, _dbgeg := _eb.DecodeStream(_fgaff)
		if _dbgeg != nil {
			return _dbgeg
		}
		_abcg.DestOutputProfile = _fgagf
	}
	return nil
}

// String returns a string that describes `font`.
func (_baced *PdfFont) String() string {
	_fbcb := ""
	if _baced._fdaa.Encoder() != nil {
		_fbcb = _baced._fdaa.Encoder().String()
	}
	return _e.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0054\u0020\u0025s\u0020\u0025\u0073\u007d", _baced._fdaa, _baced.baseFields().coreString(), _fbcb)
}

// AddChild adds a child object.
func (_bcgfd *KDict) AddChild(kv *KValue) { _bcgfd._eegge = append(_bcgfd._eegge, kv) }
func _bfcca(_bbgg *PdfPage) {
	_fgcae := _cf.GetLicenseKey()
	if _fgcae != nil && _fgcae.IsLicensed() {
		return
	}
	_gaccf := _eb.PdfObjectName("\u0055\u0046\u0031")
	if !_bbgg.Resources.HasFontByName(_gaccf) {
		_bbgg.Resources.SetFontByName(_gaccf, DefaultFont().ToPdfObject())
	}
	var _egdeb []string
	_egdeb = append(_egdeb, "\u0071")
	_egdeb = append(_egdeb, "\u0042\u0054")
	_egdeb = append(_egdeb, _e.Sprintf("\u002f%\u0073\u0020\u0031\u0034\u0020\u0054f", _gaccf.String()))
	_egdeb = append(_egdeb, "\u0031\u0020\u0030\u0020\u0030\u0020\u0072\u0067")
	_egdeb = append(_egdeb, "\u0031\u0030\u0020\u0031\u0030\u0020\u0054\u0064")
	_gadff := "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064\u0020\u0055\u006e\u0069\u0044o\u0063\u0020\u002d\u0020\u0047\u0065\u0074\u0020\u0061\u0020\u006c\u0069\u0063e\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073:/\u002f\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f"
	_egdeb = append(_egdeb, _e.Sprintf("\u0028%\u0073\u0029\u0020\u0054\u006a", _gadff))
	_egdeb = append(_egdeb, "\u0045\u0054")
	_egdeb = append(_egdeb, "\u0051")
	_efgfd := _cc.Join(_egdeb, "\u000a")
	_bbgg.AddContentStreamByString(_efgfd)
	_bbgg.ToPdfObject()
}

// PdfColorspaceCalRGB stores A, B, C components
type PdfColorspaceCalRGB struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      []float64
	Matrix     []float64
	_ccfd      *_eb.PdfObjectDictionary
	_cbfd      *_eb.PdfIndirectObject
}

// AddAnnotation appends `annot` to the list of page annotations.
func (_fgddb *PdfPage) AddAnnotation(annot *PdfAnnotation) {
	if _fgddb._dbga == nil {
		_fgddb.GetAnnotations()
	}
	_fgddb._dbga = append(_fgddb._dbga, annot)
}
func (_egbb *PdfReader) newPdfActionGoTo3DViewFromDict(_aab *_eb.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	return &PdfActionGoTo3DView{TA: _aab.Get("\u0054\u0041"), V: _aab.Get("\u0056")}, nil
}

// WriteString outputs the object as it is to be written to file.
func (_faef *PdfTransformParamsDocMDP) WriteString() string { return _faef.ToPdfObject().WriteString() }

// SetLocation sets the `Location` field of the signature.
func (_fcbbb *PdfSignature) SetLocation(location string) { _fcbbb.Location = _eb.MakeString(location) }

// Encoder returns the font's text encoder.
func (_cfgbf *pdfFontSimple) Encoder() _fc.TextEncoder {
	if _cfgbf._eccbb != nil {
		return _cfgbf._eccbb
	}
	if _cfgbf._gcgb != nil {
		return _cfgbf._gcgb
	}
	_dacdd, _ := _fc.NewSimpleTextEncoder("\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", nil)
	return _dacdd
}

// GetOCProperties returns the optional content properties PdfObject.
func (_gabbg *PdfReader) GetOCProperties() (_eb.PdfObject, error) {
	_efbg := _gabbg._bagcfd
	_ccbfa := _efbg.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_ccbfa = _eb.ResolveReference(_ccbfa)
	if !_gabbg._cfcgdf {
		_adbbc := _gabbg.traverseObjectData(_ccbfa)
		if _adbbc != nil {
			return nil, _adbbc
		}
	}
	return _ccbfa, nil
}

// SetPrintClip sets the value of the printClip.
func (_cfgfcc *ViewerPreferences) SetPrintClip(printClip PageBoundary) { _cfgfcc._edef = printClip }
func (_dcc *PdfReader) newPdfAnnotationLinkFromDict(_bbc *_eb.PdfObjectDictionary) (*PdfAnnotationLink, error) {
	_cdgg := PdfAnnotationLink{}
	_cdgg.A = _bbc.Get("\u0041")
	_cdgg.Dest = _bbc.Get("\u0044\u0065\u0073\u0074")
	_cdgg.H = _bbc.Get("\u0048")
	_cdgg.PA = _bbc.Get("\u0050\u0041")
	_cdgg.QuadPoints = _bbc.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_cdgg.BS = _bbc.Get("\u0042\u0053")
	return &_cdgg, nil
}

// NewPdfAnnotationPolygon returns a new polygon annotation.
func NewPdfAnnotationPolygon() *PdfAnnotationPolygon {
	_cfe := NewPdfAnnotation()
	_fbda := &PdfAnnotationPolygon{}
	_fbda.PdfAnnotation = _cfe
	_fbda.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cfe.SetContext(_fbda)
	return _fbda
}

// ReaderOpts defines options for creating PdfReader instances.
type ReaderOpts struct {

	// Password password of the PDF file encryption.
	// Default: empty ("").
	Password string

	// LazyLoad set if the PDF file would be loaded using lazy-loading mode.
	// Default: true.
	LazyLoad bool

	// ComplianceMode set if parsed PDF file should contain meta information for the verifiers of the compliance standards like PDF/A.
	ComplianceMode bool
}

// GetContentStreamObjs returns a slice of PDF objects containing the content
// streams of the page.
func (_ebcf *PdfPage) GetContentStreamObjs() []_eb.PdfObject {
	if _ebcf.Contents == nil {
		return nil
	}
	_edaa := _eb.TraceToDirectObject(_ebcf.Contents)
	if _dfacga, _gabcd := _edaa.(*_eb.PdfObjectArray); _gabcd {
		return _dfacga.Elements()
	}
	return []_eb.PdfObject{_edaa}
}

// B returns the value of the B component of the color.
func (_deagd *PdfColorLab) B() float64 { return _deagd[2] }

// ToPdfOutlineItem returns a low level PdfOutlineItem object,
// based on the current instance.
func (_cfad *OutlineItem) ToPdfOutlineItem() (*PdfOutlineItem, int64) {
	_abege := NewPdfOutlineItem()
	_abege.Title = _eb.MakeEncodedString(_cfad.Title, true)
	_abege.Dest = _cfad.Dest.ToPdfObject()
	var _bgaegb []*PdfOutlineItem
	var _dgbfe int64
	var _ddeda *PdfOutlineItem
	for _, _bdeed := range _cfad.Entries {
		_baffea, _fgcdcc := _bdeed.ToPdfOutlineItem()
		_baffea.Parent = &_abege.PdfOutlineTreeNode
		if _ddeda != nil {
			_ddeda.Next = &_baffea.PdfOutlineTreeNode
			_baffea.Prev = &_ddeda.PdfOutlineTreeNode
		}
		_bgaegb = append(_bgaegb, _baffea)
		_dgbfe += _fgcdcc
		_ddeda = _baffea
	}
	_ffcc := len(_bgaegb)
	_dgbfe += int64(_ffcc)
	if _ffcc > 0 {
		_abege.First = &_bgaegb[0].PdfOutlineTreeNode
		_abege.Last = &_bgaegb[_ffcc-1].PdfOutlineTreeNode
		_abege.Count = &_dgbfe
	}
	return _abege, _dgbfe
}

// GetAsTilingPattern returns a tiling pattern. Check with IsTiling() prior to using this.
func (_fbebb *PdfPattern) GetAsTilingPattern() *PdfTilingPattern {
	return _fbebb._eefgb.(*PdfTilingPattern)
}

// IsHideMenubar returns the value of the hideMenubar flag.
func (_efeea *ViewerPreferences) IsHideMenubar() bool {
	if _efeea._gcbcb == nil {
		return false
	}
	return *_efeea._gcbcb
}
func (_cgf *PdfReader) newPdfAnnotationPrinterMarkFromDict(_cdce *_eb.PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	_fcfb := PdfAnnotationPrinterMark{}
	_fcfb.MN = _cdce.Get("\u004d\u004e")
	return &_fcfb, nil
}

// ImageToRGB converts an image in CMYK32 colorspace to an RGB image.
func (_aaag *PdfColorspaceDeviceCMYK) ImageToRGB(img Image) (Image, error) {
	_ddb.Log.Trace("\u0043\u004d\u0059\u004b\u0033\u0032\u0020\u002d\u003e\u0020\u0052\u0047\u0042")
	_ddb.Log.Trace("I\u006d\u0061\u0067\u0065\u0020\u0042P\u0043\u003a\u0020\u0025\u0064\u002c \u0043\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020%\u0064", img.BitsPerComponent, img.ColorComponents)
	_ddb.Log.Trace("\u004c\u0065\u006e \u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0064", len(img.Data))
	_ddb.Log.Trace("H\u0065\u0069\u0067\u0068t:\u0020%\u0064\u002c\u0020\u0057\u0069d\u0074\u0068\u003a\u0020\u0025\u0064", img.Height, img.Width)
	_eada, _cfbe := _df.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._bdcab, img._fedc)
	if _cfbe != nil {
		return Image{}, _cfbe
	}
	_fefa, _cfbe := _df.NRGBAConverter.Convert(_eada)
	if _cfbe != nil {
		return Image{}, _cfbe
	}
	return _ggaa(_fefa.Base()), nil
}

// PdfActionTrans represents a trans action.
type PdfActionTrans struct {
	*PdfAction
	Trans _eb.PdfObject
}

func (_ecgd fontCommon) fontFlags() int {
	if _ecgd._bged == nil {
		return 0
	}
	return _ecgd._bged._eeda
}
func _agbf(_eebef string) (map[_fc.CharCode]_fc.GlyphName, error) {
	_ecag := _cc.Split(_eebef, "\u000a")
	_cgbdb := make(map[_fc.CharCode]_fc.GlyphName)
	for _, _ffebg := range _ecag {
		_bedfe := _cccge.FindStringSubmatch(_ffebg)
		if _bedfe == nil {
			continue
		}
		_bcdd, _gfcec := _bedfe[1], _bedfe[2]
		_fffba, _cged := _de.Atoi(_bcdd)
		if _cged != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0042\u0061\u0064\u0020\u0065\u006e\u0063\u006fd\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u002e \u0025\u0071", _ffebg)
			return nil, _eb.ErrTypeError
		}
		_cgbdb[_fc.CharCode(_fffba)] = _fc.GlyphName(_gfcec)
	}
	_ddb.Log.Trace("g\u0065\u0074\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0073\u003a\u0020\u006b\u0065\u0079V\u0061\u006c\u0075e\u0073=\u0025\u0023\u0076", _cgbdb)
	return _cgbdb, nil
}

// GetPageNumber returns the page number that has been assigned to the K object.
func (_daad *KDict) GetPageNumber() int64 { return _daad._dcgce }
func (_cagb *PdfColorspaceDeviceRGB) String() string {
	return "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
}

// ToPdfObject returns a PDF object representation of the outline.
func (_cfaa *Outline) ToPdfObject() _eb.PdfObject { return _cfaa.ToPdfOutline().ToPdfObject() }

// ColorAt returns the color of the image pixel specified by the x and y coordinates.
func (_ggbc *Image) ColorAt(x, y int) (_b.Color, error) {
	_edgbd := _df.BytesPerLine(int(_ggbc.Width), int(_ggbc.BitsPerComponent), _ggbc.ColorComponents)
	switch _ggbc.ColorComponents {
	case 1:
		return _df.ColorAtGrayscale(x, y, int(_ggbc.BitsPerComponent), _edgbd, _ggbc.Data, _ggbc._fedc)
	case 3:
		return _df.ColorAtNRGBA(x, y, int(_ggbc.Width), _edgbd, int(_ggbc.BitsPerComponent), _ggbc.Data, _ggbc._bdcab, _ggbc._fedc)
	case 4:
		return _df.ColorAtCMYK(x, y, int(_ggbc.Width), _ggbc.Data, _ggbc._fedc)
	}
	_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 i\u006da\u0067\u0065\u002e\u0020\u0025\u0064\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0073\u002c\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072 \u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ggbc.ColorComponents, _ggbc.BitsPerComponent)
	return nil, _dcf.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061g\u0065 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
}
func (_age *PdfAppender) updateObjectsDeep(_gcbd _eb.PdfObject, _cbea map[_eb.PdfObject]struct{}) {
	if _cbea == nil {
		_cbea = map[_eb.PdfObject]struct{}{}
	}
	if _, _ccab := _cbea[_gcbd]; _ccab || _gcbd == nil {
		return
	}
	_cbea[_gcbd] = struct{}{}
	_fcae := _eb.ResolveReferencesDeep(_gcbd, _age._fdfg)
	if _fcae != nil {
		_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fcae)
	}
	switch _agga := _gcbd.(type) {
	case *_eb.PdfIndirectObject:
		switch {
		case _agga.GetParser() == _age._adac._ebbe:
			return
		case _agga.GetParser() == _age.Reader._ebbe:
			_ccaa, _ := _age._adac.GetIndirectObjectByNumber(int(_agga.ObjectNumber))
			_ffaca, _bffaa := _ccaa.(*_eb.PdfIndirectObject)
			if _bffaa && _ffaca != nil {
				if _ffaca.PdfObject != _agga.PdfObject && _ffaca.PdfObject.WriteString() != _agga.PdfObject.WriteString() {
					if _cc.Contains(_agga.PdfObject.WriteString(), "\u002f\u0053\u0069\u0067") && _cc.Contains(_agga.PdfObject.WriteString(), "\u002f\u0053\u0075\u0062\u0074\u0079\u0070\u0065") {
						return
					}
					_age.addNewObject(_gcbd)
					_age._ebcbc[_gcbd] = _agga.ObjectNumber
				}
			}
		default:
			_age.addNewObject(_gcbd)
		}
		_age.updateObjectsDeep(_agga.PdfObject, _cbea)
	case *_eb.PdfObjectArray:
		for _, _eagga := range _agga.Elements() {
			_age.updateObjectsDeep(_eagga, _cbea)
		}
	case *_eb.PdfObjectDictionary:
		for _, _bdfbe := range _agga.Keys() {
			_age.updateObjectsDeep(_agga.Get(_bdfbe), _cbea)
		}
	case *_eb.PdfObjectStreams:
		if _agga.GetParser() != _age._adac._ebbe {
			for _, _gffe := range _agga.Elements() {
				_age.updateObjectsDeep(_gffe, _cbea)
			}
		}
	case *_eb.PdfObjectStream:
		switch {
		case _agga.GetParser() == _age._adac._ebbe:
			return
		case _agga.GetParser() == _age.Reader._ebbe:
			if _afda, _aga := _age._adac._ebbe.LookupByReference(_agga.PdfObjectReference); _aga == nil {
				var _cfbd bool
				if _dfbf, _badg := _eb.GetStream(_afda); _badg && _dd.Equal(_dfbf.Stream, _agga.Stream) {
					_cfbd = true
				}
				if _aagd, _ccfg := _eb.GetDict(_afda); _cfbd && _ccfg {
					_cfbd = _aagd.WriteString() == _agga.PdfObjectDictionary.WriteString()
				}
				if _cfbd {
					return
				}
			}
			if _agga.ObjectNumber != 0 {
				_age._ebcbc[_gcbd] = _agga.ObjectNumber
			}
		default:
			if _, _fecga := _age._accfg[_gcbd]; !_fecga {
				_age.addNewObject(_gcbd)
			}
		}
		_age.updateObjectsDeep(_agga.PdfObjectDictionary, _cbea)
	}
}

// NewPdfDateFromTime will create a PdfDate based on the given time
func NewPdfDateFromTime(timeObj _d.Time) (PdfDate, error) {
	_baefd := timeObj.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_gadca, _ := _de.ParseInt(_baefd[1:3], 10, 32)
	_bdfce, _ := _de.ParseInt(_baefd[4:6], 10, 32)
	return PdfDate{_fffdf: int64(timeObj.Year()), _aacegf: int64(timeObj.Month()), _afgfb: int64(timeObj.Day()), _bdfbf: int64(timeObj.Hour()), _aecad: int64(timeObj.Minute()), _cfdb: int64(timeObj.Second()), _eaede: _baefd[0], _fgggd: _gadca, _gadga: _bdfce}, nil
}

// GetContainingPdfObject implements model.PdfModel interface.
func (_facaa *PdfOutputIntent) GetContainingPdfObject() _eb.PdfObject { return _facaa._ebbd }

// ImageToRGB converts ICCBased colorspace image to RGB and returns the result.
func (_fad *PdfColorspaceICCBased) ImageToRGB(img Image) (Image, error) {
	if _fad.Alternate == nil {
		_ddb.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _fad.N == 1 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_dfcf := NewPdfColorspaceDeviceGray()
			return _dfcf.ImageToRGB(img)
		} else if _fad.N == 3 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return img, nil
		} else if _fad.N == 4 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_dcfe := NewPdfColorspaceDeviceCMYK()
			return _dcfe.ImageToRGB(img)
		} else {
			return img, _dcf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ddb.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _fad)
	_cege, _fgaf := _fad.Alternate.ImageToRGB(img)
	_ddb.Log.Trace("I\u0043C\u0020\u0049\u006e\u0070\u0075\u0074\u0020\u0069m\u0061\u0067\u0065\u003a %\u002b\u0076", img)
	_ddb.Log.Trace("I\u0043\u0043\u0020\u004fut\u0070u\u0074\u0020\u0069\u006d\u0061g\u0065\u003a\u0020\u0025\u002b\u0076", _cege)
	return _cege, _fgaf
}

// CheckAccessRights checks access rights and permissions for a specified password.  If either user/owner
// password is specified,  full rights are granted, otherwise the access rights are specified by the
// Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_acec *PdfReader) CheckAccessRights(password []byte) (bool, _cg.Permissions, error) {
	return _acec._ebbe.CheckAccessRights(password)
}

// ImageToRGB converts an image with samples in Separation CS to an image with samples specified in
// DeviceRGB CS.
func (_ggbd *PdfColorspaceSpecialSeparation) ImageToRGB(img Image) (Image, error) {
	_cdbd := _bb.NewReader(img.getBase())
	_bfdf := _df.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), _ggbd.AlternateSpace.GetNumComponents(), nil, img._bdcab, nil)
	_adeb := _bb.NewWriter(_bfdf)
	_fbad := _gg.Pow(2, float64(img.BitsPerComponent)) - 1
	_ddb.Log.Trace("\u0053\u0065\u0070a\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u002d\u003e\u0020\u0054\u006f\u0052\u0047\u0042\u0020\u0063o\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	_ddb.Log.Trace("\u0054i\u006et\u0054\u0072\u0061\u006e\u0073f\u006f\u0072m\u003a\u0020\u0025\u002b\u0076", _ggbd.TintTransform)
	_gddd := _ggbd.AlternateSpace.DecodeArray()
	var (
		_bgad uint32
		_edfe error
	)
	for {
		_bgad, _edfe = _cdbd.ReadSample()
		if _edfe == _bagf.EOF {
			break
		}
		if _edfe != nil {
			return img, _edfe
		}
		_eggg := float64(_bgad) / _fbad
		_dgfff, _ffce := _ggbd.TintTransform.Evaluate([]float64{_eggg})
		if _ffce != nil {
			return img, _ffce
		}
		for _ccfed, _bade := range _dgfff {
			_cgeg := _df.LinearInterpolate(_bade, _gddd[_ccfed*2], _gddd[_ccfed*2+1], 0, 1)
			if _ffce = _adeb.WriteSample(uint32(_cgeg * _fbad)); _ffce != nil {
				return img, _ffce
			}
		}
	}
	return _ggbd.AlternateSpace.ImageToRGB(_ggaa(&_bfdf))
}
func (_eebf *PdfReader) newPdfAnnotationRichMediaFromDict(_fgac *_eb.PdfObjectDictionary) (*PdfAnnotationRichMedia, error) {
	_eegc := &PdfAnnotationRichMedia{}
	_eegc.RichMediaSettings = _fgac.Get("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073")
	_eegc.RichMediaContent = _fgac.Get("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074")
	return _eegc, nil
}

// GetPdfName returns the PDF name used to indicate the border style.
// (Table 166 p. 395).
func (_bcce *BorderStyle) GetPdfName() string {
	switch *_bcce {
	case BorderStyleSolid:
		return "\u0053"
	case BorderStyleDashed:
		return "\u0044"
	case BorderStyleBeveled:
		return "\u0042"
	case BorderStyleInset:
		return "\u0049"
	case BorderStyleUnderline:
		return "\u0055"
	}
	return ""
}

// AddExtGState adds a graphics state to the XObject resources.
func (_acbbf *PdfPage) AddExtGState(name _eb.PdfObjectName, egs *_eb.PdfObjectDictionary) error {
	if _acbbf.Resources == nil {
		_acbbf.Resources = NewPdfPageResources()
	}
	if _acbbf.Resources.ExtGState == nil {
		_acbbf.Resources.ExtGState = _eb.MakeDict()
	}
	_gbfb, _aeea := _eb.TraceToDirectObject(_acbbf.Resources.ExtGState).(*_eb.PdfObjectDictionary)
	if !_aeea {
		_ddb.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _eb.TraceToDirectObject(_acbbf.Resources.ExtGState))
		return _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gbfb.Set(name, egs)
	return nil
}
func _degga(_dafbd *[]*PdfField, _gdcc FieldFilterFunc, _fgcg bool) []*PdfField {
	if _dafbd == nil {
		return nil
	}
	_aaefg := *_dafbd
	if len(*_dafbd) == 0 {
		return nil
	}
	_gabc := _aaefg[:0]
	if _gdcc == nil {
		_gdcc = func(*PdfField) bool { return true }
	}
	var _ceefa []*PdfField
	for _, _bafgc := range _aaefg {
		_dbabg := _gdcc(_bafgc)
		if _dbabg {
			_ceefa = append(_ceefa, _bafgc)
			if len(_bafgc.Kids) > 0 {
				_ceefa = append(_ceefa, _degga(&_bafgc.Kids, _gdcc, _fgcg)...)
			}
		}
		if !_fgcg || !_dbabg || len(_bafgc.Kids) > 0 {
			_gabc = append(_gabc, _bafgc)
		}
	}
	*_dafbd = _gabc
	return _ceefa
}
func (_eebc *PdfReader) newPdfAnnotationInkFromDict(_cdbg *_eb.PdfObjectDictionary) (*PdfAnnotationInk, error) {
	_cab := PdfAnnotationInk{}
	_eacg, _dfdbd := _eebc.newPdfAnnotationMarkupFromDict(_cdbg)
	if _dfdbd != nil {
		return nil, _dfdbd
	}
	_cab.PdfAnnotationMarkup = _eacg
	_cab.InkList = _cdbg.Get("\u0049n\u006b\u004c\u0069\u0073\u0074")
	_cab.BS = _cdbg.Get("\u0042\u0053")
	return &_cab, nil
}

// ToPdfObject implements interface PdfModel.
func (_adcc *PdfAnnotationPolyLine) ToPdfObject() _eb.PdfObject {
	_adcc.PdfAnnotation.ToPdfObject()
	_afb := _adcc._ggf
	_deae := _afb.PdfObject.(*_eb.PdfObjectDictionary)
	_adcc.PdfAnnotationMarkup.appendToPdfDictionary(_deae)
	_deae.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065"))
	_deae.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _adcc.Vertices)
	_deae.SetIfNotNil("\u004c\u0045", _adcc.LE)
	_deae.SetIfNotNil("\u0042\u0053", _adcc.BS)
	_deae.SetIfNotNil("\u0049\u0043", _adcc.IC)
	_deae.SetIfNotNil("\u0042\u0045", _adcc.BE)
	_deae.SetIfNotNil("\u0049\u0054", _adcc.IT)
	_deae.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _adcc.Measure)
	return _afb
}
func _ddcgg(_bgdae *_eb.PdfObjectDictionary) (*PdfTilingPattern, error) {
	_ecgcea := &PdfTilingPattern{}
	_gegdg := _bgdae.Get("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054\u0079\u0070\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_cabcf, _gdcdf := _gegdg.(*_eb.PdfObjectInteger)
	if !_gdcdf {
		_ddb.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054y\u0070\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _gegdg)
		return nil, _eb.ErrTypeError
	}
	_ecgcea.PaintType = _cabcf
	_gegdg = _bgdae.Get("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0054i\u006ci\u006e\u0067\u0054\u0079\u0070e\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dbgf, _gdcdf := _gegdg.(*_eb.PdfObjectInteger)
	if !_gdcdf {
		_ddb.Log.Debug("\u0054\u0069\u006cin\u0067\u0054\u0079\u0070\u0065\u0020\u006e\u006f\u0074 \u0061n\u0020i\u006et\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gegdg)
		return nil, _eb.ErrTypeError
	}
	_ecgcea.TilingType = _dbgf
	_gegdg = _bgdae.Get("\u0042\u0042\u006f\u0078")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0042\u0042\u006fx\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gegdg = _eb.TraceToDirectObject(_gegdg)
	_dgea, _gdcdf := _gegdg.(*_eb.PdfObjectArray)
	if !_gdcdf {
		_ddb.Log.Debug("\u0042B\u006f\u0078 \u0073\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061y\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gegdg)
		return nil, _eb.ErrTypeError
	}
	_ffbfb, _cdfce := NewPdfRectangle(*_dgea)
	if _cdfce != nil {
		_ddb.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _cdfce)
		return nil, _cdfce
	}
	_ecgcea.BBox = _ffbfb
	_gegdg = _bgdae.Get("\u0058\u0053\u0074e\u0070")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0058\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_adbeb, _cdfce := _eb.GetNumberAsFloat(_gegdg)
	if _cdfce != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0058S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _adbeb)
		return nil, _cdfce
	}
	_ecgcea.XStep = _eb.MakeFloat(_adbeb)
	_gegdg = _bgdae.Get("\u0059\u0053\u0074e\u0070")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0059\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_befgg, _cdfce := _eb.GetNumberAsFloat(_gegdg)
	if _cdfce != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0059S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _befgg)
		return nil, _cdfce
	}
	_ecgcea.YStep = _eb.MakeFloat(_befgg)
	_gegdg = _bgdae.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _gegdg == nil {
		_ddb.Log.Debug("\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bgdae, _gdcdf = _eb.TraceToDirectObject(_gegdg).(*_eb.PdfObjectDictionary)
	if !_gdcdf {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _gegdg)
	}
	_fdbbe, _cdfce := NewPdfPageResourcesFromDict(_bgdae)
	if _cdfce != nil {
		return nil, _cdfce
	}
	_ecgcea.Resources = _fdbbe
	if _ebea := _bgdae.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ebea != nil {
		_bbadb, _eacgf := _ebea.(*_eb.PdfObjectArray)
		if !_eacgf {
			_ddb.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ebea)
			return nil, _eb.ErrTypeError
		}
		_ecgcea.Matrix = _bbadb
	}
	return _ecgcea, nil
}
func _fdbc(_gbcc *fontCommon) *pdfFontType0 { return &pdfFontType0{fontCommon: *_gbcc} }

// GetContentStream returns the XObject Form's content stream.
func (_dacca *XObjectForm) GetContentStream() ([]byte, error) {
	_fgaae, _gdccb := _eb.DecodeStream(_dacca._afcag)
	if _gdccb != nil {
		return nil, _gdccb
	}
	return _fgaae, nil
}

// ToPdfObject implements interface PdfModel.
func (_def *PdfAnnotationMovie) ToPdfObject() _eb.PdfObject {
	_def.PdfAnnotation.ToPdfObject()
	_gaca := _def._ggf
	_bdgg := _gaca.PdfObject.(*_eb.PdfObjectDictionary)
	_bdgg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u004d\u006f\u0076i\u0065"))
	_bdgg.SetIfNotNil("\u0054", _def.T)
	_bdgg.SetIfNotNil("\u004d\u006f\u0076i\u0065", _def.Movie)
	_bdgg.SetIfNotNil("\u0041", _def.A)
	return _gaca
}

// GetAlphabet returns a map of the runes in `text` and their frequencies.
func GetAlphabet(text string) map[rune]int {
	_abag := map[rune]int{}
	for _, _bfdff := range text {
		_abag[_bfdff]++
	}
	return _abag
}

// NewPdfReaderFromFile creates a new PdfReader from the speficied PDF file.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderFromFile(pdfFile string, opts *ReaderOpts) (*PdfReader, *_ccb.File, error) {
	const _gbgb = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0046\u0072\u006f\u006dF\u0069\u006c\u0065"
	_effdb, _dggce := _ccb.Open(pdfFile)
	if _dggce != nil {
		return nil, nil, _dggce
	}
	_gfbae, _dggce := _dcgfe(_effdb, opts, true, _gbgb)
	if _dggce != nil {
		_effdb.Close()
		return nil, nil, _dggce
	}
	_gfbae._fbgfg = pdfFile
	return _gfbae, _effdb, nil
}
func _bfecg(_aefbe _eb.PdfObject, _cfda _eb.PdfObject, _bgdfd map[_eb.PdfObject][]_eb.PdfObject, _dedf map[string]_eb.PdfObject, _gefbf *[]_eb.PdfObject) {
	var _aggag *_eb.PdfIndirectObject
	if _fbggb, _bcaag := _eb.GetIndirect(_aefbe); _bcaag {
		_aggag = _fbggb
		_aefbe = _fbggb.PdfObject
	}
	switch _aedad := _aefbe.(type) {
	case *_eb.PdfObjectDictionary:
		if _aedad.Get("\u0053") == nil {
			return
		}
		_aedad.Set("\u0050", _cfda)
		if _fcega := _aedad.Get("\u0050\u0067"); _fcega != nil {
			if _cbdab, _gced := _eb.GetIndirect(_fcega); _gced && _cbdab != nil && _cbdab.PdfObject != nil {
				_bgdfd[_fcega] = append(_bgdfd[_fcega], _aggag)
			}
		}
		if _gacefg := _aedad.Get("\u0053"); _gacefg != nil {
			if _ggfdge, _acgdf := _eb.GetNameVal(_gacefg); _acgdf {
				if _ggfdge == StructureTypeLink {
					if _deacb := _aedad.Get("\u004b"); _deacb != nil {
						if _agfca, _dcbbg := _eb.GetArray(_deacb); _dcbbg && _agfca.Len() == 2 {
							_gfdac := false
							_acee := false
							for _, _ecde := range _agfca.Elements() {
								if _bbcegc, _cgfga := _eb.GetDict(_ecde); _cgfga {
									if _fcbfb, _gdad := _eb.GetName(_bbcegc.Get("\u0054\u0079\u0070\u0065")); _gdad && _fcbfb.String() == "\u004f\u0042\u004a\u0052" {
										_acee = true
									}
								} else if _, _cegab := _eb.GetInt(_ecde); _cegab {
									_gfdac = true
								}
							}
							if _gfdac && _acee {
								*_gefbf = append(*_gefbf, _aggag)
							}
						} else if _aabcb, _defc := _eb.GetDict(_deacb); _defc {
							if _bbaec, _aefcg := _eb.GetName(_aabcb.Get("\u0054\u0079\u0070\u0065")); _aefcg && _bbaec.String() == "\u004f\u0042\u004a\u0052" {
								*_gefbf = append(*_gefbf, _aggag)
							}
						}
					}
				}
			}
		}
		if _ffabb := _aedad.Get("\u0049\u0044"); _ffabb != nil {
			_dedf[_ffabb.String()] = _aefbe
		}
		if _babbe := _aedad.Get("\u004b"); _babbe != nil {
			_bfecg(_babbe, _aggag, _bgdfd, _dedf, _gefbf)
		}
	case *_eb.PdfObjectArray:
		for _, _bfbdb := range _aedad.Elements() {
			_bfecg(_bfbdb, _cfda, _bgdfd, _dedf, _gefbf)
		}
	default:
	}
}
func (_gebg *PdfReader) newPdfAnnotationTrapNetFromDict(_fedb *_eb.PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	_agba := PdfAnnotationTrapNet{}
	return &_agba, nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_cfcfe *PdfShadingType5) ToPdfObject() _eb.PdfObject {
	_cfcfe.PdfShading.ToPdfObject()
	_degba, _ceadb := _cfcfe.getShadingDict()
	if _ceadb != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cfcfe.BitsPerCoordinate != nil {
		_degba.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _cfcfe.BitsPerCoordinate)
	}
	if _cfcfe.BitsPerComponent != nil {
		_degba.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cfcfe.BitsPerComponent)
	}
	if _cfcfe.VerticesPerRow != nil {
		_degba.Set("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077", _cfcfe.VerticesPerRow)
	}
	if _cfcfe.Decode != nil {
		_degba.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _cfcfe.Decode)
	}
	if _cfcfe.Function != nil {
		if len(_cfcfe.Function) == 1 {
			_degba.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cfcfe.Function[0].ToPdfObject())
		} else {
			_edafe := _eb.MakeArray()
			for _, _egdeg := range _cfcfe.Function {
				_edafe.Append(_egdeg.ToPdfObject())
			}
			_degba.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _edafe)
		}
	}
	return _cfcfe._cefaa
}

// SetCenterWindow sets the value of the centerWindow flag.
func (_bfeag *ViewerPreferences) SetCenterWindow(centerWindow bool) { _bfeag._dacc = &centerWindow }

// GetCharMetrics returns the character metrics for the specified character code.  A bool flag is
// returned to indicate whether or not the entry was found in the glyph to charcode mapping.
// How it works:
//  1. Return a value the /Widths array (charWidths) if there is one.
//  2. If the font has the same name as a standard 14 font then return width=250.
//  3. Otherwise return no match and let the caller substitute a default.
func (_eeddf pdfFontSimple) GetCharMetrics(code _fc.CharCode) (_fg.CharMetrics, bool) {
	if _efab, _fbcbfc := _eeddf._cegda[code]; _fbcbfc {
		return _fg.CharMetrics{Wx: _efab}, true
	}
	if _fg.IsStdFont(_fg.StdFontName(_eeddf._agcc)) {
		return _fg.CharMetrics{Wx: 250}, true
	}
	return _fg.CharMetrics{}, false
}

// NewImageFromGoImage creates a new NRGBA32 unidoc Image from a golang Image.
// If `goimg` is grayscale (*goimage.Gray8) then calls NewGrayImageFromGoImage instead.
func (_cegee DefaultImageHandler) NewImageFromGoImage(goimg _fb.Image) (*Image, error) {
	_cfcaa, _dgbae := _df.FromGoImage(goimg)
	if _dgbae != nil {
		return nil, _dgbae
	}
	_abeg := _ggaa(_cfcaa.Base())
	return &_abeg, nil
}
func _cfgc(_cadbb _eb.PdfObject) (*PdfColorspaceCalGray, error) {
	_gabd := NewPdfColorspaceCalGray()
	if _egdc, _gadc := _cadbb.(*_eb.PdfIndirectObject); _gadc {
		_gabd._gdee = _egdc
	}
	_cadbb = _eb.TraceToDirectObject(_cadbb)
	_afcfc, _dgffa := _cadbb.(*_eb.PdfObjectArray)
	if !_dgffa {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _afcfc.Len() != 2 {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0061\u006cG\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_cadbb = _eb.TraceToDirectObject(_afcfc.Get(0))
	_dcad, _dgffa := _cadbb.(*_eb.PdfObjectName)
	if !_dgffa {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u0020\u006e\u0061m\u0065\u0020\u006e\u006f\u0074\u0020\u0061 \u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_dcad != "\u0043a\u006c\u0047\u0072\u0061\u0079" {
		return nil, _e.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0043\u0061\u006c\u0047\u0072\u0061\u0079 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_cadbb = _eb.TraceToDirectObject(_afcfc.Get(1))
	_bbcfb, _dgffa := _cadbb.(*_eb.PdfObjectDictionary)
	if !_dgffa {
		return nil, _e.Errorf("\u0043\u0061lG\u0072\u0061\u0079 \u0064\u0069\u0063\u0074 no\u0074 a\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u006f\u0062\u006a\u0065c\u0074")
	}
	_cadbb = _bbcfb.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_cadbb = _eb.TraceToDirectObject(_cadbb)
	_defd, _dgffa := _cadbb.(*_eb.PdfObjectArray)
	if !_dgffa {
		return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020W\u0068\u0069\u0074e\u0050o\u0069\u006e\u0074")
	}
	if _defd.Len() != 3 {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
	}
	_cgfg, _cbcg := _defd.GetAsFloat64Slice()
	if _cbcg != nil {
		return nil, _cbcg
	}
	_gabd.WhitePoint = _cgfg
	_cadbb = _bbcfb.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _cadbb != nil {
		_cadbb = _eb.TraceToDirectObject(_cadbb)
		_dbgd, _acda := _cadbb.(*_eb.PdfObjectArray)
		if !_acda {
			return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020B\u006c\u0061\u0063k\u0050o\u0069\u006e\u0074")
		}
		if _dbgd.Len() != 3 {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061c\u006b\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
		}
		_gdbcb, _ggdg := _dbgd.GetAsFloat64Slice()
		if _ggdg != nil {
			return nil, _ggdg
		}
		_gabd.BlackPoint = _gdbcb
	}
	_cadbb = _bbcfb.Get("\u0047\u0061\u006dm\u0061")
	if _cadbb != nil {
		_cadbb = _eb.TraceToDirectObject(_cadbb)
		_cdaa, _cegdgc := _eb.GetNumberAsFloat(_cadbb)
		if _cegdgc != nil {
			return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0067\u0061\u006d\u006d\u0061\u0020\u006e\u006ft\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_gabd.Gamma = _cdaa
	}
	return _gabd, nil
}

// WriteToFile writes the Appender output to file specified by path.
func (_fedfa *PdfAppender) WriteToFile(outputPath string) error {
	_edaf, _degeg := _ccb.Create(outputPath)
	if _degeg != nil {
		return _degeg
	}
	defer _edaf.Close()
	return _fedfa.Write(_edaf)
}
func (_facde *PdfWriter) optimizeDocument() error {
	if _facde._ggbcd == nil {
		return nil
	}
	_ebaga, _dadgc := _eb.GetDict(_facde._ecagd)
	if !_dadgc {
		return _dcf.New("\u0061\u006e\u0020in\u0066\u006f\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_beeaec := _cd.Document{ID: [2]string{_facde._cebgc, _facde._bgddbe}, Version: _facde._edbbf, Objects: _facde._dcfgf, Info: _ebaga, Crypt: _facde._gadcaa, UseHashBasedID: _facde._beeaf}
	if _dfadc := _facde._ggbcd.ApplyStandard(&_beeaec); _dfadc != nil {
		return _dfadc
	}
	_facde._cebgc, _facde._bgddbe = _beeaec.ID[0], _beeaec.ID[1]
	_facde._edbbf = _beeaec.Version
	_facde._dcfgf = _beeaec.Objects
	_facde._ecagd.PdfObject = _beeaec.Info
	_facde._beeaf = _beeaec.UseHashBasedID
	_facde._gadcaa = _beeaec.Crypt
	_gadcaad := make(map[_eb.PdfObject]struct{}, len(_facde._dcfgf))
	for _, _dabfg := range _facde._dcfgf {
		_gadcaad[_dabfg] = struct{}{}
	}
	_facde._aeeda = _gadcaad
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_dggc *PdfAnnotationPolygon) ToPdfObject() _eb.PdfObject {
	_dggc.PdfAnnotation.ToPdfObject()
	_bdb := _dggc._ggf
	_gacef := _bdb.PdfObject.(*_eb.PdfObjectDictionary)
	_dggc.PdfAnnotationMarkup.appendToPdfDictionary(_gacef)
	_gacef.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0050o\u006c\u0079\u0067\u006f\u006e"))
	_gacef.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _dggc.Vertices)
	_gacef.SetIfNotNil("\u004c\u0045", _dggc.LE)
	_gacef.SetIfNotNil("\u0042\u0053", _dggc.BS)
	_gacef.SetIfNotNil("\u0049\u0043", _dggc.IC)
	_gacef.SetIfNotNil("\u0042\u0045", _dggc.BE)
	_gacef.SetIfNotNil("\u0049\u0054", _dggc.IT)
	_gacef.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _dggc.Measure)
	return _bdb
}

// ToGoImage converts the unidoc Image to a golang Image structure.
func (_cgfab *Image) ToGoImage() (_fb.Image, error) {
	_ddb.Log.Trace("\u0043\u006f\u006e\u0076er\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u006f\u0020\u0069\u006d\u0061g\u0065")
	_adde, _abacf := _df.NewImage(int(_cgfab.Width), int(_cgfab.Height), int(_cgfab.BitsPerComponent), _cgfab.ColorComponents, _cgfab.Data, _cgfab._bdcab, _cgfab._fedc)
	if _abacf != nil {
		return nil, _abacf
	}
	return _adde, nil
}

// ToPdfObject returns the PDF representation of the function.
func (_dgggg *PdfFunctionType2) ToPdfObject() _eb.PdfObject {
	_fcfdf := _eb.MakeDict()
	_fcfdf.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _eb.MakeInteger(2))
	_baaac := &_eb.PdfObjectArray{}
	for _, _efebe := range _dgggg.Domain {
		_baaac.Append(_eb.MakeFloat(_efebe))
	}
	_fcfdf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _baaac)
	if _dgggg.Range != nil {
		_gecd := &_eb.PdfObjectArray{}
		for _, _edee := range _dgggg.Range {
			_gecd.Append(_eb.MakeFloat(_edee))
		}
		_fcfdf.Set("\u0052\u0061\u006eg\u0065", _gecd)
	}
	if _dgggg.C0 != nil {
		_cecge := &_eb.PdfObjectArray{}
		for _, _efacd := range _dgggg.C0 {
			_cecge.Append(_eb.MakeFloat(_efacd))
		}
		_fcfdf.Set("\u0043\u0030", _cecge)
	}
	if _dgggg.C1 != nil {
		_dagad := &_eb.PdfObjectArray{}
		for _, _fgcdc := range _dgggg.C1 {
			_dagad.Append(_eb.MakeFloat(_fgcdc))
		}
		_fcfdf.Set("\u0043\u0031", _dagad)
	}
	_fcfdf.Set("\u004e", _eb.MakeFloat(_dgggg.N))
	if _dgggg._efee != nil {
		_dgggg._efee.PdfObject = _fcfdf
		return _dgggg._efee
	}
	return _fcfdf
}

// FieldFlag represents form field flags. Some of the flags can apply to all types of fields whereas other
// flags are specific.
type FieldFlag uint32

// ToPdfObject converts the pdfCIDFontType2 to a PDF representation.
func (_bbaef *pdfCIDFontType2) ToPdfObject() _eb.PdfObject {
	if _bbaef._fceef == nil {
		_bbaef._fceef = &_eb.PdfIndirectObject{}
	}
	_cecc := _bbaef.baseFields().asPdfObjectDictionary("\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032")
	_bbaef._fceef.PdfObject = _cecc
	if _bbaef.CIDSystemInfo != nil {
		_cecc.Set("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f", _bbaef.CIDSystemInfo)
	}
	if _bbaef.DW != nil {
		_cecc.Set("\u0044\u0057", _bbaef.DW)
	}
	if _bbaef.DW2 != nil {
		_cecc.Set("\u0044\u0057\u0032", _bbaef.DW2)
	}
	if _bbaef.W != nil {
		_cecc.Set("\u0057", _bbaef.W)
	}
	if _bbaef.W2 != nil {
		_cecc.Set("\u0057\u0032", _bbaef.W2)
	}
	if _bbaef.CIDToGIDMap != nil {
		_cecc.Set("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070", _bbaef.CIDToGIDMap)
	}
	return _bbaef._fceef
}

// Set applies flag fl to the flag's bitmask and returns the combined flag.
func (_bdfdg FieldFlag) Set(fl FieldFlag) FieldFlag { return FieldFlag(_bdfdg.Mask() | fl.Mask()) }

const (
	ButtonTypeCheckbox ButtonType = iota
	ButtonTypePush     ButtonType = iota
	ButtonTypeRadio    ButtonType = iota
)

// SetPatternByName sets a pattern resource specified by keyName.
func (_cfdge *PdfPageResources) SetPatternByName(keyName _eb.PdfObjectName, pattern _eb.PdfObject) error {
	if _cfdge.Pattern == nil {
		_cfdge.Pattern = _eb.MakeDict()
	}
	_accfga, _befce := _eb.GetDict(_cfdge.Pattern)
	if !_befce {
		return _eb.ErrTypeError
	}
	_accfga.Set(keyName, pattern)
	return nil
}

// Width returns the width of `rect`.
func (_gcbeb *PdfRectangle) Width() float64 { return _gg.Abs(_gcbeb.Urx - _gcbeb.Llx) }

// AnnotFilterFunc represents a PDF annotation filtering function. If the function
// returns true, the annotation is kept, otherwise it is discarded.
type AnnotFilterFunc func(*PdfAnnotation) bool

func _fbcef(_gfcg *_eb.PdfObjectDictionary, _fbabg *fontCommon) (*pdfCIDFontType2, error) {
	if _fbabg._fgdee != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_ddb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0032\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _fbabg)
		return nil, _eb.ErrRangeError
	}
	_gffa := _fegae(_fbabg)
	_bbbg, _agef := _eb.GetDict(_gfcg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_agef {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _fbabg)
		return nil, ErrRequiredAttributeMissing
	}
	_gffa.CIDSystemInfo = _bbbg
	_gffa.DW = _gfcg.Get("\u0044\u0057")
	_gffa.W = _gfcg.Get("\u0057")
	_gffa.DW2 = _gfcg.Get("\u0044\u0057\u0032")
	_gffa.W2 = _gfcg.Get("\u0057\u0032")
	_gffa.CIDToGIDMap = _gfcg.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070")
	_gffa._aggd = 1000.0
	if _faggc, _fbffb := _eb.GetNumberAsFloat(_gffa.DW); _fbffb == nil {
		_gffa._aggd = _faggc
	}
	_fgafa, _ddgdc := _cbfdd(_gffa.W)
	if _ddgdc != nil {
		return nil, _ddgdc
	}
	if _fgafa == nil {
		_fgafa = map[_fc.CharCode]float64{}
	}
	_gffa._fcgg = _fgafa
	return _gffa, nil
}

// ToPdfObject implements interface PdfModel.
func (_bdf *PdfActionGoToR) ToPdfObject() _eb.PdfObject {
	_bdf.PdfAction.ToPdfObject()
	_fdc := _bdf._dee
	_fecf := _fdc.PdfObject.(*_eb.PdfObjectDictionary)
	_fecf.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeGoToR)))
	if _bdf.F != nil {
		_fecf.Set("\u0046", _bdf.F.ToPdfObject())
	}
	_fecf.SetIfNotNil("\u0044", _bdf.D)
	_fecf.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _bdf.NewWindow)
	return _fdc
}
func (_eccbc *PdfReader) newPdfSignatureReferenceFromDict(_cgegg *_eb.PdfObjectDictionary) (*PdfSignatureReference, error) {
	if _ccce, _dgge := _eccbc._affaf.GetModelFromPrimitive(_cgegg).(*PdfSignatureReference); _dgge {
		return _ccce, nil
	}
	_ebbb := &PdfSignatureReference{_abdcb: _cgegg, Data: _cgegg.Get("\u0044\u0061\u0074\u0061")}
	var _bddbc bool
	_ebbb.Type, _ = _eb.GetName(_cgegg.Get("\u0054\u0079\u0070\u0065"))
	_ebbb.TransformMethod, _bddbc = _eb.GetName(_cgegg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"))
	if !_bddbc {
		_ddb.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0072\u0061\u006e\u0073\u0066o\u0072\u006dM\u0065\u0074h\u006f\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020in\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020m\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_ebbb.TransformParams, _ = _eb.GetDict(_cgegg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"))
	_ebbb.DigestMethod, _ = _eb.GetName(_cgegg.Get("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064"))
	return _ebbb, nil
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the reader using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
// At the end of the process, the AcroForm contains all the fields which were
// not flattened. If all fields are flattened, the reader's AcroForm field
// is set to nil.
func (_affa *PdfReader) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	return _affa.flattenFieldsWithOpts(false, appgen, opts)
}

// GetFontByName gets the font specified by keyName. Returns the PdfObject which
// the entry refers to. Returns a bool value indicating whether or not the entry was found.
func (_faad *PdfPageResources) GetFontByName(keyName _eb.PdfObjectName) (_eb.PdfObject, bool) {
	if _faad.Font == nil {
		return nil, false
	}
	_adff, _ccacc := _eb.TraceToDirectObject(_faad.Font).(*_eb.PdfObjectDictionary)
	if !_ccacc {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _eb.TraceToDirectObject(_faad.Font))
		return nil, false
	}
	if _facfe := _adff.Get(keyName); _facfe != nil {
		return _facfe, true
	}
	return nil, false
}

// M returns the value of the magenta component of the color.
func (_dbfa *PdfColorDeviceCMYK) M() float64 { return _dbfa[1] }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_dcdb *PdfShadingType4) ToPdfObject() _eb.PdfObject {
	_dcdb.PdfShading.ToPdfObject()
	_fegcg, _bdaec := _dcdb.getShadingDict()
	if _bdaec != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _dcdb.BitsPerCoordinate != nil {
		_fegcg.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _dcdb.BitsPerCoordinate)
	}
	if _dcdb.BitsPerComponent != nil {
		_fegcg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _dcdb.BitsPerComponent)
	}
	if _dcdb.BitsPerFlag != nil {
		_fegcg.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _dcdb.BitsPerFlag)
	}
	if _dcdb.Decode != nil {
		_fegcg.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _dcdb.Decode)
	}
	if _dcdb.Function != nil {
		if len(_dcdb.Function) == 1 {
			_fegcg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dcdb.Function[0].ToPdfObject())
		} else {
			_dfaga := _eb.MakeArray()
			for _, _bbega := range _dcdb.Function {
				_dfaga.Append(_bbega.ToPdfObject())
			}
			_fegcg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dfaga)
		}
	}
	return _dcdb._cefaa
}

// GenerateRandomID generates a random ID for the K dictionary object.
func (_cdbec *KDict) GenerateRandomID() string {
	_abef, _effdc := _deb.NewUUID()
	if _effdc != nil {
		_ddb.Log.Debug("\u0045r\u0072\u006f\u0072\u0020g\u0065\u006e\u0065\u0072\u0061t\u0069n\u0067 \u0055\u0055\u0049\u0044\u003a\u0020\u0025v", _effdc)
	}
	if _abef != _deb.Nil {
		_dagfc := _abef.String()
		_cdbec.ID = _eb.MakeString(_dagfc)
		return _dagfc
	}
	return ""
}

// PdfColorDeviceCMYK is a CMYK32 color, where each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceCMYK [4]float64

// GetContainingPdfObject implements interface PdfModel.
func (_ffa *PdfAction) GetContainingPdfObject() _eb.PdfObject { return _ffa._dee }
func _dfgf(_gffac _eb.PdfObject) (*PdfPageResourcesColorspaces, error) {
	_bgbad := &PdfPageResourcesColorspaces{}
	if _agbga, _eagc := _gffac.(*_eb.PdfIndirectObject); _eagc {
		_bgbad._ccba = _agbga
		_gffac = _agbga.PdfObject
	}
	_gbgfd, _fgdbe := _eb.GetDict(_gffac)
	if !_fgdbe {
		return nil, _dcf.New("\u0043\u0053\u0020at\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bgbad.Names = []string{}
	_bgbad.Colorspaces = map[string]PdfColorspace{}
	for _, _cefeec := range _gbgfd.Keys() {
		_ecbegf := _gbgfd.Get(_cefeec)
		_bgbad.Names = append(_bgbad.Names, string(_cefeec))
		_gdfca, _ggfeg := NewPdfColorspaceFromPdfObject(_ecbegf)
		if _ggfeg != nil {
			return nil, _ggfeg
		}
		_bgbad.Colorspaces[string(_cefeec)] = _gdfca
	}
	return _bgbad, nil
}

// ToPdfObject implements model.PdfModel interface.
func (_debcd *PdfOutputIntent) ToPdfObject() _eb.PdfObject {
	if _debcd._ebbd == nil {
		_debcd._ebbd = _eb.MakeDict()
	}
	_ggbg := _debcd._ebbd
	if _debcd.Type != "" {
		_ggbg.Set("\u0054\u0079\u0070\u0065", _eb.MakeName(_debcd.Type))
	}
	_ggbg.Set("\u0053", _eb.MakeName(_debcd.S.String()))
	if _debcd.OutputCondition != "" {
		_ggbg.Set("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e", _eb.MakeString(_debcd.OutputCondition))
	}
	_ggbg.Set("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r", _eb.MakeString(_debcd.OutputConditionIdentifier))
	_ggbg.Set("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065", _eb.MakeString(_debcd.RegistryName))
	if _debcd.Info != "" {
		_ggbg.Set("\u0049\u006e\u0066\u006f", _eb.MakeString(_debcd.Info))
	}
	if len(_debcd.DestOutputProfile) != 0 {
		_ggfea, _edeeg := _eb.MakeStream(_debcd.DestOutputProfile, _eb.NewFlateEncoder())
		if _edeeg != nil {
			_ddb.Log.Error("\u004d\u0061\u006b\u0065\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0044\u0065s\u0074\u004f\u0075\u0074\u0070\u0075t\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _edeeg)
		}
		_ggfea.PdfObjectDictionary.Set("\u004e", _eb.MakeInteger(int64(_debcd.ColorComponents)))
		_bbea := make([]float64, _debcd.ColorComponents*2)
		for _gdcdd := 0; _gdcdd < _debcd.ColorComponents*2; _gdcdd++ {
			_afdefa := 0.0
			if _gdcdd%2 != 0 {
				_afdefa = 1.0
			}
			_bbea[_gdcdd] = _afdefa
		}
		_ggfea.PdfObjectDictionary.Set("\u0052\u0061\u006eg\u0065", _eb.MakeArrayFromFloats(_bbea))
		_ggbg.Set("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065", _ggfea)
	}
	return _ggbg
}

// Evaluate runs the function. Input is [x1 x2 x3].
func (_efged *PdfFunctionType4) Evaluate(xVec []float64) ([]float64, error) {
	if _efged._aabg == nil {
		_efged._aabg = _gc.NewPSExecutor(_efged.Program)
	}
	var _cffd []_gc.PSObject
	for _, _faag := range xVec {
		_cffd = append(_cffd, _gc.MakeReal(_faag))
	}
	_gfbac, _dceag := _efged._aabg.Execute(_cffd)
	if _dceag != nil {
		return nil, _dceag
	}
	_ebdee, _dceag := _gc.PSObjectArrayToFloat64Array(_gfbac)
	if _dceag != nil {
		return nil, _dceag
	}
	return _ebdee, nil
}
func _cbdd(_dcef _eb.PdfObject) (*_eb.PdfObjectDictionary, *fontCommon, error) {
	_gcfb := &fontCommon{}
	if _eaaab, _abfa := _dcef.(*_eb.PdfIndirectObject); _abfa {
		_gcfb._babgg = _eaaab.ObjectNumber
	}
	_gagb, _cadbd := _eb.GetDict(_dcef)
	if !_cadbd {
		_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _dcef)
		return nil, nil, ErrFontNotSupported
	}
	_bbef, _cadbd := _eb.GetNameVal(_gagb.Get("\u0054\u0079\u0070\u0065"))
	if !_cadbd {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	if _bbef != "\u0046\u006f\u006e\u0074" {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u0020\u0049\u006e\u0063\u006f\u006d\u0070\u0061t\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u003d\u0025\u0071\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0025\u0071.", _bbef, "\u0046\u006f\u006e\u0074")
		return nil, nil, _eb.ErrTypeError
	}
	_gdef, _cadbd := _eb.GetNameVal(_gagb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_cadbd {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u006f\u006e\u0074 \u0049\u006e\u0063o\u006d\u0070a\u0074\u0069\u0062\u0069\u006c\u0069t\u0079. \u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	_gcfb._fgdee = _gdef
	_cbde, _cadbd := _eb.GetNameVal(_gagb.Get("\u004e\u0061\u006d\u0065"))
	if _cadbd {
		_gcfb._bgge = _cbde
	}
	_fddf := _gagb.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _fddf != nil {
		_gcfb._geee = _eb.TraceToDirectObject(_fddf)
		_gaae, _dbadb := _geebd(_gcfb._geee, _gcfb)
		if _dbadb != nil {
			return _gagb, _gcfb, _dbadb
		}
		_gcfb._bgbg = _gaae
	} else if _gdef == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" || _gdef == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_bgec, _gaccc := _ff.NewCIDSystemInfo(_gagb.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
		if _gaccc != nil {
			return _gagb, _gcfb, _gaccc
		}
		_dadc := _e.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _bgec.Registry, _bgec.Ordering)
		if _ff.IsPredefinedCMap(_dadc) {
			_gcfb._bgbg, _gaccc = _ff.LoadPredefinedCMap(_dadc)
			if _gaccc != nil {
				_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _dadc, _gaccc)
			}
		}
	}
	_bbcfe := _gagb.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
	if _bbcfe != nil {
		_gefg, _gfba := _cbff(_bbcfe)
		if _gfba != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0042\u0061\u0064\u0020\u0066\u006f\u006et\u0020d\u0065s\u0063r\u0069\u0070\u0074\u006f\u0072\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gfba)
			return _gagb, _gcfb, _gfba
		}
		_gcfb._bged = _gefg
	}
	if _gdef != "\u0054\u0079\u0070e\u0033" {
		_fbff, _eebff := _eb.GetNameVal(_gagb.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
		if !_eebff {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u006f\u006et\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069t\u0079\u002e\u0020\u0042\u0061se\u0046\u006f\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
			return _gagb, _gcfb, ErrRequiredAttributeMissing
		}
		_gcfb._agcc = _fbff
	}
	return _gagb, _gcfb, nil
}

type pdfCIDFontType0 struct {
	fontCommon
	_dbeda *_eb.PdfIndirectObject
	_ebgbc _fc.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// (Required) Dictionary that defines the character collection of the CIDFont.
	// See Table 116.
	CIDSystemInfo *_eb.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW     _eb.PdfObject
	W      _eb.PdfObject
	DW2    _eb.PdfObject
	W2     _eb.PdfObject
	_efege map[_fc.CharCode]float64
	_dcfg  float64
}

// SetRotation sets the rotation of all pages added to writer. The rotation is
// specified in degrees and must be a multiple of 90.
// The Rotate field of individual pages has priority over the global rotation.
func (_cafeg *PdfWriter) SetRotation(rotate int64) error {
	_gdfgg, _ebffb := _eb.GetDict(_cafeg._afga)
	if !_ebffb {
		return ErrTypeCheck
	}
	_gdfgg.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _eb.MakeInteger(rotate))
	return nil
}

// SetContext sets the sub pattern (context).  Either PdfTilingPattern or PdfShadingPattern.
func (_fdceb *PdfPattern) SetContext(ctx PdfModel) { _fdceb._eefgb = ctx }
func (_feed *PdfReader) newPdfAnnotationFromIndirectObject(_dbd *_eb.PdfIndirectObject) (*PdfAnnotation, error) {
	_bddd, _gffd := _dbd.PdfObject.(*_eb.PdfObjectDictionary)
	if !_gffd {
		return nil, _e.Errorf("\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _feed._affaf.GetModelFromPrimitive(_bddd); model != nil {
		_cegg, _dega := model.(*PdfAnnotation)
		if !_dega {
			return nil, _e.Errorf("\u0063\u0061\u0063\u0068\u0065\u0064 \u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050D\u0046\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e")
		}
		return _cegg, nil
	}
	_abce := &PdfAnnotation{}
	_abce._ggf = _dbd
	_feed._affaf.Register(_bddd, _abce)
	if _afcf := _bddd.Get("\u0054\u0079\u0070\u0065"); _afcf != nil {
		_gfaf, _ggfd := _afcf.(*_eb.PdfObjectName)
		if !_ggfd {
			_ddb.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _afcf)
		} else {
			if *_gfaf != "\u0041\u006e\u006eo\u0074" {
				_ddb.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070\u0065\u0063\u0074\u0065d\u0020\u0054\u0079\u0070\u0065\u0020\u0021=\u0020\u0041\u006e\u006e\u006f\u0074\u0020\u0028\u0025\u0073\u0029", *_gfaf)
			}
		}
	}
	if _fbe := _bddd.Get("\u0052\u0065\u0063\u0074"); _fbe != nil {
		_abce.Rect = _fbe
	}
	if _dgf := _bddd.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _dgf != nil {
		_abce.Contents = _dgf
	}
	if _bef := _bddd.Get("\u0050"); _bef != nil {
		_abce.P = _bef
	}
	if _cbbb := _bddd.Get("\u004e\u004d"); _cbbb != nil {
		_abce.NM = _cbbb
	}
	if _cgdb := _bddd.Get("\u004d"); _cgdb != nil {
		_abce.M = _cgdb
	}
	if _acfe := _bddd.Get("\u0046"); _acfe != nil {
		_abce.F = _acfe
	}
	if _cacc := _bddd.Get("\u0041\u0050"); _cacc != nil {
		_abce.AP = _cacc
	}
	if _ecdd := _bddd.Get("\u0041\u0053"); _ecdd != nil {
		_abce.AS = _ecdd
	}
	if _gdcg := _bddd.Get("\u0042\u006f\u0072\u0064\u0065\u0072"); _gdcg != nil {
		_abce.Border = _gdcg
	}
	if _gdcgc := _bddd.Get("\u0043"); _gdcgc != nil {
		_abce.C = _gdcgc
	}
	if _eee := _bddd.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074"); _eee != nil {
		_abce.StructParent = _eee
	}
	if _fba := _bddd.Get("\u004f\u0043"); _fba != nil {
		_abce.OC = _fba
	}
	_gffc := _bddd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
	if _gffc == nil {
		_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u0049\u004e\u0047:\u0020\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079 \u0069s\u0073\u0075\u0065\u0020\u002d\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073u\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074\u0079p\u0065")
		_abce._cdb = nil
		return _abce, nil
	}
	_dgee, _agd := _gffc.(*_eb.PdfObjectName)
	if !_agd {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065 !\u003d\u0020n\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _gffc)
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d n\u0061\u006d\u0065 \u0028%\u0054\u0029", _gffc)
	}
	switch *_dgee {
	case "\u0054\u0065\u0078\u0074":
		_cfce, _fce := _feed.newPdfAnnotationTextFromDict(_bddd)
		if _fce != nil {
			return nil, _fce
		}
		_cfce.PdfAnnotation = _abce
		_abce._cdb = _cfce
		return _abce, nil
	case "\u004c\u0069\u006e\u006b":
		_cfcgd, _feee := _feed.newPdfAnnotationLinkFromDict(_bddd)
		if _feee != nil {
			return nil, _feee
		}
		_cfcgd.PdfAnnotation = _abce
		_abce._cdb = _cfcgd
		return _abce, nil
	case "\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074":
		_agb, _bbdc := _feed.newPdfAnnotationFreeTextFromDict(_bddd)
		if _bbdc != nil {
			return nil, _bbdc
		}
		_agb.PdfAnnotation = _abce
		_abce._cdb = _agb
		return _abce, nil
	case "\u004c\u0069\u006e\u0065":
		_ddd, _efad := _feed.newPdfAnnotationLineFromDict(_bddd)
		if _efad != nil {
			return nil, _efad
		}
		_ddd.PdfAnnotation = _abce
		_abce._cdb = _ddd
		_ddb.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041N\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0061\u006e\u006eo\u0074\u0020\u0028\u0025\u0054\u0029\u003a \u0025\u002b\u0076\u000a", _abce, _abce)
		_ddb.Log.Trace("\u004c\u0049\u004eE\u0020\u0041\u004e\u004eO\u0054\u0041\u0054\u0049\u004f\u004e\u003a \u0063\u0074\u0078\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u002b\u0076\u000a", _ddd, _ddd)
		_ddb.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041\u004e\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u0020\u004d\u0061\u0072\u006b\u0075\u0070\u003a\u0020c\u0074\u0078\u0020\u0028\u0025T\u0029\u003a \u0025\u002b\u0076\u000a", _ddd.PdfAnnotationMarkup, _ddd.PdfAnnotationMarkup)
		return _abce, nil
	case "\u0053\u0071\u0075\u0061\u0072\u0065":
		_efba, _addc := _feed.newPdfAnnotationSquareFromDict(_bddd)
		if _addc != nil {
			return nil, _addc
		}
		_efba.PdfAnnotation = _abce
		_abce._cdb = _efba
		return _abce, nil
	case "\u0043\u0069\u0072\u0063\u006c\u0065":
		_dac, _gaf := _feed.newPdfAnnotationCircleFromDict(_bddd)
		if _gaf != nil {
			return nil, _gaf
		}
		_dac.PdfAnnotation = _abce
		_abce._cdb = _dac
		return _abce, nil
	case "\u0050o\u006c\u0079\u0067\u006f\u006e":
		_gcfe, _eeg := _feed.newPdfAnnotationPolygonFromDict(_bddd)
		if _eeg != nil {
			return nil, _eeg
		}
		_gcfe.PdfAnnotation = _abce
		_abce._cdb = _gcfe
		return _abce, nil
	case "\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065":
		_cad, _fbgd := _feed.newPdfAnnotationPolyLineFromDict(_bddd)
		if _fbgd != nil {
			return nil, _fbgd
		}
		_cad.PdfAnnotation = _abce
		_abce._cdb = _cad
		return _abce, nil
	case "\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t":
		_edga, _bad := _feed.newPdfAnnotationHighlightFromDict(_bddd)
		if _bad != nil {
			return nil, _bad
		}
		_edga.PdfAnnotation = _abce
		_abce._cdb = _edga
		return _abce, nil
	case "\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee":
		_egcd, _ddae := _feed.newPdfAnnotationUnderlineFromDict(_bddd)
		if _ddae != nil {
			return nil, _ddae
		}
		_egcd.PdfAnnotation = _abce
		_abce._cdb = _egcd
		return _abce, nil
	case "\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079":
		_ffda, _ggg := _feed.newPdfAnnotationSquigglyFromDict(_bddd)
		if _ggg != nil {
			return nil, _ggg
		}
		_ffda.PdfAnnotation = _abce
		_abce._cdb = _ffda
		return _abce, nil
	case "\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t":
		_abb, _eega := _feed.newPdfAnnotationStrikeOut(_bddd)
		if _eega != nil {
			return nil, _eega
		}
		_abb.PdfAnnotation = _abce
		_abce._cdb = _abb
		return _abce, nil
	case "\u0043\u0061\u0072e\u0074":
		_fff, _fbeb := _feed.newPdfAnnotationCaretFromDict(_bddd)
		if _fbeb != nil {
			return nil, _fbeb
		}
		_fff.PdfAnnotation = _abce
		_abce._cdb = _fff
		return _abce, nil
	case "\u0053\u0074\u0061m\u0070":
		_cbfa, _ddga := _feed.newPdfAnnotationStampFromDict(_bddd)
		if _ddga != nil {
			return nil, _ddga
		}
		_cbfa.PdfAnnotation = _abce
		_abce._cdb = _cbfa
		return _abce, nil
	case "\u0049\u006e\u006b":
		_eaad, _gaab := _feed.newPdfAnnotationInkFromDict(_bddd)
		if _gaab != nil {
			return nil, _gaab
		}
		_eaad.PdfAnnotation = _abce
		_abce._cdb = _eaad
		return _abce, nil
	case "\u0050\u006f\u0070u\u0070":
		_bgb, _daec := _feed.newPdfAnnotationPopupFromDict(_bddd)
		if _daec != nil {
			return nil, _daec
		}
		_bgb.PdfAnnotation = _abce
		_abce._cdb = _bgb
		return _abce, nil
	case "\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074":
		_gage, _dfde := _feed.newPdfAnnotationFileAttachmentFromDict(_bddd)
		if _dfde != nil {
			return nil, _dfde
		}
		_gage.PdfAnnotation = _abce
		_abce._cdb = _gage
		return _abce, nil
	case "\u0053\u006f\u0075n\u0064":
		_cegd, _ceee := _feed.newPdfAnnotationSoundFromDict(_bddd)
		if _ceee != nil {
			return nil, _ceee
		}
		_cegd.PdfAnnotation = _abce
		_abce._cdb = _cegd
		return _abce, nil
	case "\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a":
		_gfb, _ecg := _feed.newPdfAnnotationRichMediaFromDict(_bddd)
		if _ecg != nil {
			return nil, _ecg
		}
		_gfb.PdfAnnotation = _abce
		_abce._cdb = _gfb
		return _abce, nil
	case "\u004d\u006f\u0076i\u0065":
		_gffg, _cfcd := _feed.newPdfAnnotationMovieFromDict(_bddd)
		if _cfcd != nil {
			return nil, _cfcd
		}
		_gffg.PdfAnnotation = _abce
		_abce._cdb = _gffg
		return _abce, nil
	case "\u0053\u0063\u0072\u0065\u0065\u006e":
		_cfd, _gdbg := _feed.newPdfAnnotationScreenFromDict(_bddd)
		if _gdbg != nil {
			return nil, _gdbg
		}
		_cfd.PdfAnnotation = _abce
		_abce._cdb = _cfd
		return _abce, nil
	case "\u0057\u0069\u0064\u0067\u0065\u0074":
		_fddd, _fgbbd := _feed.newPdfAnnotationWidgetFromDict(_bddd)
		if _fgbbd != nil {
			return nil, _fgbbd
		}
		_fddd.PdfAnnotation = _abce
		_abce._cdb = _fddd
		return _abce, nil
	case "P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b":
		_cacf, _ebcc := _feed.newPdfAnnotationPrinterMarkFromDict(_bddd)
		if _ebcc != nil {
			return nil, _ebcc
		}
		_cacf.PdfAnnotation = _abce
		_abce._cdb = _cacf
		return _abce, nil
	case "\u0054r\u0061\u0070\u004e\u0065\u0074":
		_ddbd, _beeb := _feed.newPdfAnnotationTrapNetFromDict(_bddd)
		if _beeb != nil {
			return nil, _beeb
		}
		_ddbd.PdfAnnotation = _abce
		_abce._cdb = _ddbd
		return _abce, nil
	case "\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k":
		_aae, _ged := _feed.newPdfAnnotationWatermarkFromDict(_bddd)
		if _ged != nil {
			return nil, _ged
		}
		_aae.PdfAnnotation = _abce
		_abce._cdb = _aae
		return _abce, nil
	case "\u0033\u0044":
		_dedg, _bafg := _feed.newPdfAnnotation3DFromDict(_bddd)
		if _bafg != nil {
			return nil, _bafg
		}
		_dedg.PdfAnnotation = _abce
		_abce._cdb = _dedg
		return _abce, nil
	case "\u0050\u0072\u006f\u006a\u0065\u0063\u0074\u0069\u006f\u006e":
		_bedf, _dbb := _feed.newPdfAnnotationProjectionFromDict(_bddd)
		if _dbb != nil {
			return nil, _dbb
		}
		_bedf.PdfAnnotation = _abce
		_abce._cdb = _bedf
		return _abce, nil
	case "\u0052\u0065\u0064\u0061\u0063\u0074":
		_gdbcg, _gdd := _feed.newPdfAnnotationRedactFromDict(_bddd)
		if _gdd != nil {
			return nil, _gdd
		}
		_gdbcg.PdfAnnotation = _abce
		_abce._cdb = _gdbcg
		return _abce, nil
	}
	_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069\u006fn\u003a\u0020\u0025\u0073", *_dgee)
	return nil, nil
}
func (_aea *PdfReader) newPdfAnnotationWatermarkFromDict(_adb *_eb.PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	_fggd := PdfAnnotationWatermark{}
	_fggd.FixedPrint = _adb.Get("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074")
	return &_fggd, nil
}

// GetAllContentStreams gets all the content streams for a page as one string.
func (_gcgd *PdfPage) GetAllContentStreams() (string, error) {
	_bcga, _gbbeb := _gcgd.GetContentStreams()
	if _gbbeb != nil {
		return "", _gbbeb
	}
	return _cc.Join(_bcga, "\u0020"), nil
}
func (_eaea *PdfReader) buildPageList(_eabga *_eb.PdfIndirectObject, _eddfd *_eb.PdfIndirectObject, _cfbca map[_eb.PdfObject]struct{}) error {
	if _eabga == nil {
		return nil
	}
	if _, _gega := _cfbca[_eabga]; _gega {
		_ddb.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _eabga.ObjectNumber)
		return nil
	}
	_cfbca[_eabga] = struct{}{}
	_feab, _dgac := _eabga.PdfObject.(*_eb.PdfObjectDictionary)
	if !_dgac {
		return _dcf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_gfdd, _dgac := (*_feab).Get("\u0054\u0079\u0070\u0065").(*_eb.PdfObjectName)
	if !_dgac {
		if _feab.Get("\u004b\u0069\u0064\u0073") == nil {
			return _dcf.New("\u006e\u006f\u0064\u0065 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0054\u0079p\u0065 \u0028\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0029")
		}
		_ddb.Log.Debug("ER\u0052\u004fR\u003a\u0020\u006e\u006f\u0064\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067\u0020\u0054\u0079\u0070\u0065\u002c\u0020\u0062\u0075\u0074\u0020\u0068\u0061\u0073\u0020\u004b\u0069\u0064\u0073\u002e\u0020\u0041\u0073\u0073u\u006di\u006e\u0067\u0020\u0050\u0061\u0067\u0065\u0073 \u006eo\u0064\u0065.")
		_gfdd = _eb.MakeName("\u0050\u0061\u0067e\u0073")
		_feab.Set("\u0054\u0079\u0070\u0065", _gfdd)
	}
	_ddb.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0050a\u0067\u0065\u004c\u0069\u0073\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0074y\u0070\u0065\u003a\u0020\u0025\u0073\u0020(\u0025\u002b\u0076\u0029", *_gfdd, _eabga)
	if *_gfdd == "\u0050\u0061\u0067\u0065" {
		_cedbb, _dafec := _eaea.newPdfPageFromDict(_feab)
		if _dafec != nil {
			return _dafec
		}
		_cedbb.setContainer(_eabga)
		if _eddfd != nil {
			_feab.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eddfd)
		}
		_eaea._cbaff = append(_eaea._cbaff, _eabga)
		_eaea.PageList = append(_eaea.PageList, _cedbb)
		return nil
	}
	if *_gfdd != "\u0050\u0061\u0067e\u0073" {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0063\u006fnt\u0065n\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u0021\u0020\u0028\u0025\u0073\u0029", _gfdd)
		return _dcf.New("\u0074\u0061\u0062\u006c\u0065\u0020o\u0066\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073 \u006fb\u006a\u0065\u0063\u0074")
	}
	if _eddfd != nil {
		_feab.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eddfd)
	}
	if !_eaea._cfcgdf {
		_feefe := _eaea.traverseObjectData(_eabga)
		if _feefe != nil {
			return _feefe
		}
	}
	_bbbcc, _cbegb := _eaea._ebbe.Resolve(_feab.Get("\u004b\u0069\u0064\u0073"))
	if _cbegb != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u004b\u0069\u0064\u0073\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _cbegb
	}
	var _ageec *_eb.PdfObjectArray
	_ageec, _dgac = _bbbcc.(*_eb.PdfObjectArray)
	if !_dgac {
		_agcba, _eeggd := _bbbcc.(*_eb.PdfIndirectObject)
		if !_eeggd {
			return _dcf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_ageec, _dgac = _agcba.PdfObject.(*_eb.PdfObjectArray)
		if !_dgac {
			return _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u0069\u006ed\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
	}
	_ddb.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _ageec)
	for _gbdbf, _dedcg := range _ageec.Elements() {
		_aecge, _fecea := _eb.GetIndirect(_dedcg)
		if !_fecea {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 \u006f\u0062\u006a\u0065\u0063t\u0020\u002d \u0028\u0025\u0073\u0029", _aecge)
			return _dcf.New("\u0070a\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_ageec.Set(_gbdbf, _aecge)
		_cbegb = _eaea.buildPageList(_aecge, _eabga, _cfbca)
		if _cbegb != nil {
			return _cbegb
		}
	}
	return nil
}

// String returns a string representation of PdfTransformParamsDocMDP.
func (_edfa *PdfTransformParamsDocMDP) String() string {
	return _e.Sprintf("\u0025\u0073\u0020\u0050\u003a\u0020\u0025\u0073\u0020V\u003a\u0020\u0025\u0073", _edfa.Type, _edfa.P, _edfa.V)
}

// ImageToRGB converts CalRGB colorspace image to RGB and returns the result.
func (_agedb *PdfColorspaceCalRGB) ImageToRGB(img Image) (Image, error) {
	_cdeb := _bb.NewReader(img.getBase())
	_effa := _df.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_babaf := _bb.NewWriter(_effa)
	_cebg := _gg.Pow(2, float64(img.BitsPerComponent)) - 1
	_fdbd := make([]uint32, 3)
	var (
		_gddc                                     error
		_cbef, _dgedd, _adccb, _fea, _dcbg, _gdac float64
	)
	for {
		_gddc = _cdeb.ReadSamples(_fdbd)
		if _gddc == _bagf.EOF {
			break
		} else if _gddc != nil {
			return img, _gddc
		}
		_cbef = float64(_fdbd[0]) / _cebg
		_dgedd = float64(_fdbd[1]) / _cebg
		_adccb = float64(_fdbd[2]) / _cebg
		_fea = _agedb.Matrix[0]*_gg.Pow(_cbef, _agedb.Gamma[0]) + _agedb.Matrix[3]*_gg.Pow(_dgedd, _agedb.Gamma[1]) + _agedb.Matrix[6]*_gg.Pow(_adccb, _agedb.Gamma[2])
		_dcbg = _agedb.Matrix[1]*_gg.Pow(_cbef, _agedb.Gamma[0]) + _agedb.Matrix[4]*_gg.Pow(_dgedd, _agedb.Gamma[1]) + _agedb.Matrix[7]*_gg.Pow(_adccb, _agedb.Gamma[2])
		_gdac = _agedb.Matrix[2]*_gg.Pow(_cbef, _agedb.Gamma[0]) + _agedb.Matrix[5]*_gg.Pow(_dgedd, _agedb.Gamma[1]) + _agedb.Matrix[8]*_gg.Pow(_adccb, _agedb.Gamma[2])
		_cbef = 3.240479*_fea + -1.537150*_dcbg + -0.498535*_gdac
		_dgedd = -0.969256*_fea + 1.875992*_dcbg + 0.041556*_gdac
		_adccb = 0.055648*_fea + -0.204043*_dcbg + 1.057311*_gdac
		_cbef = _gg.Min(_gg.Max(_cbef, 0), 1.0)
		_dgedd = _gg.Min(_gg.Max(_dgedd, 0), 1.0)
		_adccb = _gg.Min(_gg.Max(_adccb, 0), 1.0)
		_fdbd[0] = uint32(_cbef * _cebg)
		_fdbd[1] = uint32(_dgedd * _cebg)
		_fdbd[2] = uint32(_adccb * _cebg)
		if _gddc = _babaf.WriteSamples(_fdbd); _gddc != nil {
			return img, _gddc
		}
	}
	return _ggaa(&_effa), nil
}
func (_abad *PdfAnnotationMarkup) appendToPdfDictionary(_gdae *_eb.PdfObjectDictionary) {
	_gdae.SetIfNotNil("\u0054", _abad.T)
	if _abad.Popup != nil {
		_gdae.Set("\u0050\u006f\u0070u\u0070", _abad.Popup.ToPdfObject())
	}
	_gdae.SetIfNotNil("\u0043\u0041", _abad.CA)
	_gdae.SetIfNotNil("\u0052\u0043", _abad.RC)
	_gdae.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _abad.CreationDate)
	_gdae.SetIfNotNil("\u0049\u0052\u0054", _abad.IRT)
	_gdae.SetIfNotNil("\u0053\u0075\u0062\u006a", _abad.Subj)
	_gdae.SetIfNotNil("\u0052\u0054", _abad.RT)
	_gdae.SetIfNotNil("\u0049\u0054", _abad.IT)
	_gdae.SetIfNotNil("\u0045\u0078\u0044\u0061\u0074\u0061", _abad.ExData)
}

var _ pdfFont = (*pdfCIDFontType2)(nil)

// ToPdfObject converts the K dictionary to a PDF object.
func (_gdbca *KValue) ToPdfObject() _eb.PdfObject {
	if _gdbca._ccbca != nil {
		return _eb.MakeIndirectObject(_gdbca._ccbca.ToPdfObject())
	}
	if _gdbca._fbgaa != nil {
		return _gdbca._fbgaa
	}
	if _gdbca._ddaf != nil {
		return _eb.MakeInteger(int64(*_gdbca._ddaf))
	}
	return nil
}

// RemovePage removes a page by number.
func (_egdf *PdfAppender) RemovePage(pageNum int) {
	_gdgdb := pageNum - 1
	_egdf._fedg = append(_egdf._fedg[0:_gdgdb], _egdf._fedg[pageNum:]...)
}

// SetXObjectImageByName adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_egecgg *PdfPageResources) SetXObjectImageByName(keyName _eb.PdfObjectName, ximg *XObjectImage) error {
	_ggaf := ximg.ToPdfObject().(*_eb.PdfObjectStream)
	_egge := _egecgg.SetXObjectByName(keyName, _ggaf)
	return _egge
}

// SetViewClip sets the value of the viewClip.
func (_cfegf *ViewerPreferences) SetViewClip(viewClip PageBoundary) { _cfegf._eadae = viewClip }

const (
	RC4_128bit = EncryptionAlgorithm(iota)
	AES_128bit
	AES_256bit
)

// NewPdfActionGoToR returns a new "go to remote" action.
func NewPdfActionGoToR() *PdfActionGoToR {
	_ccg := NewPdfAction()
	_bdd := &PdfActionGoToR{}
	_bdd.PdfAction = _ccg
	_ccg.SetContext(_bdd)
	return _bdd
}

// Write writes the Appender output to io.Writer.
// It can only be called once and further invocations will result in an error.
func (_cegdg *PdfAppender) Write(w _bagf.Writer) error {
	if _cegdg._ffac {
		return _dcf.New("\u0061\u0070\u0070\u0065\u006e\u0064\u0065\u0072\u0020\u0077\u0072\u0069\u0074e\u0020\u0063\u0061\u006e\u0020\u006fn\u006c\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0076\u006f\u006b\u0065\u0064 \u006f\u006e\u0063\u0065")
	}
	_dcgd := NewPdfWriter()
	_fgcd, _fcgc := _eb.GetDict(_dcgd._afga)
	if !_fcgc {
		return _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_efec, _fcgc := _fgcd.Get("\u004b\u0069\u0064\u0073").(*_eb.PdfObjectArray)
	if !_fcgc {
		return _dcf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_fdcc, _fcgc := _fgcd.Get("\u0043\u006f\u0075n\u0074").(*_eb.PdfObjectInteger)
	if !_fcgc {
		return _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	_gfed := _cegdg._adac._ebbe
	_geab := _gfed.GetTrailer()
	if _geab == nil {
		return _dcf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_dede, _fcgc := _eb.GetIndirect(_geab.Get("\u0052\u006f\u006f\u0074"))
	if !_fcgc {
		return _dcf.New("c\u0061\u0074\u0061\u006c\u006f\u0067 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064")
	}
	_abba, _fcgc := _eb.GetDict(_dede)
	if !_fcgc {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _dede, *_geab)
		return _dcf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_ffgc := false
	for _, _cae := range _cegdg._adac.AcroForm.signatureFields() {
		if _cae.Lock != nil {
			_ffgc = true
			break
		}
	}
	if _ffgc {
		_dcgd._ccea = _dede
	}
	for _, _eae := range _abba.Keys() {
		if _dcgd._dbffa.Get(_eae) == nil {
			_ffgg := _abba.Get(_eae)
			_dcgd._dbffa.Set(_eae, _ffgg)
		}
	}
	if _cegdg._bddda != nil {
		if _cegdg._bddda._cgedg {
			if _gdcd := _eb.TraceToDirectObject(_cegdg._bddda.ToPdfObject()); !_eb.IsNullObject(_gdcd) {
				_dcgd._dbffa.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _gdcd)
				_cegdg.updateObjectsDeep(_gdcd, nil)
			} else {
				_ddb.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020t\u0072\u0061\u0063e\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0020o\u0062\u006a\u0065\u0063\u0074, \u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u002e")
			}
		} else {
			_dcgd._dbffa.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _cegdg._bddda.ToPdfObject())
			_cegdg.updateObjectsDeep(_cegdg._bddda.ToPdfObject(), nil)
		}
	}
	if _cegdg._cedc != nil {
		_cegdg.updateObjectsDeep(_cegdg._cedc.ToPdfObject(), nil)
		_dcgd._dbffa.Set("\u0044\u0053\u0053", _cegdg._cedc.GetContainingPdfObject())
	}
	if _cegdg._fbec != nil {
		_dcgd._dbffa.Set("\u0050\u0065\u0072m\u0073", _cegdg._fbec.ToPdfObject())
		_cegdg.updateObjectsDeep(_cegdg._fbec.ToPdfObject(), nil)
	}
	if _dcgd._edbbf.Major < 2 {
		_dcgd.AddExtension("\u0045\u0053\u0049\u0043", "\u0031\u002e\u0037", 5)
		_dcgd.AddExtension("\u0041\u0044\u0042\u0045", "\u0031\u002e\u0037", 8)
	}
	if _dged, _dcca := _eb.GetDict(_geab.Get("\u0049\u006e\u0066\u006f")); _dcca {
		if _edgd, _ecgf := _eb.GetDict(_dcgd._ecagd); _ecgf {
			for _, _bbaeb := range _dged.Keys() {
				if _edgd.Get(_bbaeb) == nil {
					_edgd.Set(_bbaeb, _dged.Get(_bbaeb))
				}
			}
		}
	}
	if _cegdg._ccag != nil {
		_dcgd._ecagd = _eb.MakeIndirectObject(_cegdg._ccag.ToPdfObject())
	}
	_cegdg.updateObjectsDeep(_dcgd._ecagd, nil)
	_cegdg.updateObjectsDeep(_dcgd._ccea, nil)
	_egab := false
	if len(_cegdg._adac.PageList) != len(_cegdg._fedg) {
		_egab = true
	} else {
		for _agff := range _cegdg._adac.PageList {
			switch {
			case _cegdg._fedg[_agff] == _cegdg._adac.PageList[_agff]:
			case _cegdg._fedg[_agff] == _cegdg.Reader.PageList[_agff]:
			default:
				_egab = true
			}
			if _egab {
				break
			}
		}
	}
	if _egab {
		_cegdg.updateObjectsDeep(_dcgd._afga, nil)
	} else {
		_cegdg._dbae[_dcgd._afga] = struct{}{}
	}
	_dcgd._afga.ObjectNumber = _cegdg.Reader._agfdg.ObjectNumber
	_cegdg._ebcbc[_dcgd._afga] = _cegdg.Reader._agfdg.ObjectNumber
	_ccgc := []_eb.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	for _, _bfdbe := range _cegdg._fedg {
		_efed := _bfdbe.ToPdfObject()
		*_fdcc = *_fdcc + 1
		if _ddcc, _cgfc := _efed.(*_eb.PdfIndirectObject); _cgfc && _ddcc.GetParser() == _cegdg._adac._ebbe {
			_efec.Append(&_ddcc.PdfObjectReference)
			continue
		}
		if _aaea, _acdd := _eb.GetDict(_efed); _acdd {
			_adbf, _bcge := _aaea.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_eb.PdfIndirectObject)
			for _bcge {
				_ddb.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _adbf)
				_gacb, _bfgb := _adbf.PdfObject.(*_eb.PdfObjectDictionary)
				if !_bfgb {
					return _dcf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
				}
				for _, _acbec := range _ccgc {
					_ddb.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _acbec)
					if _gfeg := _aaea.Get(_acbec); _gfeg != nil {
						_ddb.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
						if len(_bfdbe._eaebc.Keys()) > 0 && !_egab {
							_feec := _bfdbe._eaebc
							if _eeeg := _feec.Get(_acbec); _eeeg != nil {
								if _gfeg != _eeeg {
									_ddb.Log.Trace("\u0049\u006e\u0068\u0065\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u006f\u0072\u0069\u0067i\u006ea\u006c\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073\u002c\u0020\u0025\u0054", _acbec, _eeeg)
									_aaea.Set(_acbec, _eeeg)
								}
							}
						}
						continue
					}
					if _eegdd := _gacb.Get(_acbec); _eegdd != nil {
						_ddb.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _acbec)
						_aaea.Set(_acbec, _eegdd)
					}
				}
				_adbf, _bcge = _gacb.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_eb.PdfIndirectObject)
				_ddb.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _gacb.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
			}
			if _egab {
				_aaea.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _dcgd._afga)
			}
		}
		_cegdg.updateObjectsDeep(_efed, nil)
		_efec.Append(_efed)
	}
	if _, _aged := _cegdg._bccd.Seek(0, _bagf.SeekStart); _aged != nil {
		return _aged
	}
	_edcf := make(map[SignatureHandler]_bagf.Writer)
	_cbbf := _eb.MakeArray()
	for _, _bcf := range _cegdg._dfg {
		if _ddfa, _ebd := _eb.GetIndirect(_bcf); _ebd {
			if _begc, _dfgc := _ddfa.PdfObject.(*pdfSignDictionary); _dfgc {
				_afea := *_begc._bfadg
				var _gaeac error
				_edcf[_afea], _gaeac = _afea.NewDigest(_begc._abegc)
				if _gaeac != nil {
					return _gaeac
				}
				_cbbf.Append(_eb.MakeInteger(0xfffff), _eb.MakeInteger(0xfffff))
			}
		}
	}
	if _cbbf.Len() > 0 {
		_cbbf.Append(_eb.MakeInteger(0xfffff), _eb.MakeInteger(0xfffff))
	}
	for _, _ggcc := range _cegdg._dfg {
		if _eeaad, _bggd := _eb.GetIndirect(_ggcc); _bggd {
			if _ggbf, _gcec := _eeaad.PdfObject.(*pdfSignDictionary); _gcec {
				_ggbf.Set("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _cbbf)
			}
		}
	}
	_gdeb := len(_edcf) > 0
	var _effg _bagf.Reader = _cegdg._bccd
	if _gdeb {
		_adg := make([]_bagf.Writer, 0, len(_edcf))
		for _, _dbgg := range _edcf {
			_adg = append(_adg, _dbgg)
		}
		_effg = _bagf.TeeReader(_cegdg._bccd, _bagf.MultiWriter(_adg...))
	}
	_fbcf, _cfea := _bagf.Copy(w, _effg)
	if _cfea != nil {
		return _cfea
	}
	if len(_cegdg._dfg) == 0 {
		return nil
	}
	_dcgd._gedbe = _fbcf
	_dcgd.ObjNumOffset = _cegdg._cfgdf
	_dcgd._caed = true
	_dcgd._bfdeb = _cegdg._dce
	_dcgd._becbf = _cegdg._dgc
	_dcgd._dfbcf = _cegdg._eabd
	_dcgd._edbbf = _cegdg._adac.PdfVersion()
	_dcgd._ecbgf = _cegdg._ebcbc
	_dcgd._gadcaa = _cegdg._ebcb.GetCrypter()
	_dcgd._ebdeed = _cegdg._ebcb.GetEncryptObj()
	_faeg := _cegdg._ebcb.GetXrefType()
	if _faeg != nil {
		_bdeg := *_faeg == _eb.XrefTypeObjectStream
		_dcgd._bddggg = &_bdeg
	}
	_dcgd._aeeda = map[_eb.PdfObject]struct{}{}
	_dcgd._dcfgf = []_eb.PdfObject{}
	for _, _abd := range _cegdg._dfg {
		if _, _ecdg := _cegdg._dbae[_abd]; _ecdg {
			continue
		}
		_dcgd.addObject(_abd)
	}
	_eedea := w
	if _gdeb {
		_eedea = _dd.NewBuffer(nil)
	}
	if _cegdg._cedcf != "" && _dcgd._gadcaa == nil {
		_dcgd.Encrypt([]byte(_cegdg._cedcf), []byte(_cegdg._cedcf), _cegdg._fcfd)
	}
	if _egae := _geab.Get("\u0049\u0044"); _egae != nil {
		if _edcef, _bfdg := _eb.GetArray(_egae); _bfdg {
			_dcgd._ddefd = _edcef
		}
	}
	if _afad := _dcgd.Write(_eedea); _afad != nil {
		return _afad
	}
	if _gdeb {
		_edgb := _eedea.(*_dd.Buffer).Bytes()
		_gee := _eb.MakeArray()
		var _gegf []*pdfSignDictionary
		var _cbcc int64
		for _, _edgg := range _dcgd._dcfgf {
			if _aebe, _cdge := _eb.GetIndirect(_edgg); _cdge {
				if _bfe, _gebe := _aebe.PdfObject.(*pdfSignDictionary); _gebe {
					_gegf = append(_gegf, _bfe)
					_dfda := _bfe._afeabg + int64(_bfe._eegea)
					_gee.Append(_eb.MakeInteger(_cbcc), _eb.MakeInteger(_dfda-_cbcc))
					_cbcc = _bfe._afeabg + int64(_bfe._defaa)
				}
			}
		}
		_gee.Append(_eb.MakeInteger(_cbcc), _eb.MakeInteger(_fbcf+int64(len(_edgb))-_cbcc))
		_fbba := []byte(_gee.WriteString())
		for _, _ebfd := range _gegf {
			_abbd := int(_ebfd._afeabg - _fbcf)
			for _bfce := _ebfd._caeba; _bfce < _ebfd._dbcd; _bfce++ {
				_edgb[_abbd+_bfce] = ' '
			}
			_fdde := _edgb[_abbd+_ebfd._caeba : _abbd+_ebfd._dbcd]
			copy(_fdde, _fbba)
		}
		var _gbd int
		for _, _cfbdc := range _gegf {
			_bcdg := int(_cfbdc._afeabg - _fbcf)
			_ffdd := _edgb[_gbd : _bcdg+_cfbdc._eegea]
			_adcd := *_cfbdc._bfadg
			_edcf[_adcd].Write(_ffdd)
			_gbd = _bcdg + _cfbdc._defaa
		}
		for _, _fgdfg := range _gegf {
			_bgfg := _edgb[_gbd:]
			_fcaa := *_fgdfg._bfadg
			_edcf[_fcaa].Write(_bgfg)
		}
		for _, _agdcc := range _gegf {
			_eace := int(_agdcc._afeabg - _fbcf)
			_ebac := *_agdcc._bfadg
			_adfaa := _edcf[_ebac]
			if _adeg := _ebac.Sign(_agdcc._abegc, _adfaa); _adeg != nil {
				return _adeg
			}
			_agdcc._abegc.ByteRange = _gee
			_agaf := []byte(_agdcc._abegc.Contents.WriteString())
			for _acce := _agdcc._caeba; _acce < _agdcc._dbcd; _acce++ {
				_edgb[_eace+_acce] = ' '
			}
			for _bcfb := _agdcc._eegea; _bcfb < _agdcc._defaa; _bcfb++ {
				_edgb[_eace+_bcfb] = ' '
			}
			_babg := _edgb[_eace+_agdcc._caeba : _eace+_agdcc._dbcd]
			copy(_babg, _fbba)
			_babg = _edgb[_eace+_agdcc._eegea : _eace+_agdcc._defaa]
			copy(_babg, _agaf)
		}
		_bgfc := _dd.NewBuffer(_edgb)
		_, _cfea = _bagf.Copy(w, _bgfc)
		if _cfea != nil {
			return _cfea
		}
	}
	_cegdg._ffac = true
	return nil
}

// NewViewerPreferences returns a new ViewerPreferences object with
// default empty values.
func NewViewerPreferences() *ViewerPreferences { return &ViewerPreferences{} }
func _cafc(_eccgf []*_eb.PdfObjectStream) *_eb.PdfObjectArray {
	if len(_eccgf) == 0 {
		return nil
	}
	_egddd := make([]_eb.PdfObject, 0, len(_eccgf))
	for _, _efae := range _eccgf {
		_egddd = append(_egddd, _efae)
	}
	return _eb.MakeArray(_egddd...)
}

// PdfOutline represents a PDF outline dictionary (Table 152 - p. 376).
type PdfOutline struct {
	PdfOutlineTreeNode
	Parent *PdfOutlineTreeNode
	Count  *int64
	_becfb *_eb.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
func (_ad *PdfActionRendition) ToPdfObject() _eb.PdfObject {
	_ad.PdfAction.ToPdfObject()
	_cba := _ad._dee
	_eaf := _cba.PdfObject.(*_eb.PdfObjectDictionary)
	_eaf.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeRendition)))
	_eaf.SetIfNotNil("\u0052", _ad.R)
	_eaf.SetIfNotNil("\u0041\u004e", _ad.AN)
	_eaf.SetIfNotNil("\u004f\u0050", _ad.OP)
	_eaf.SetIfNotNil("\u004a\u0053", _ad.JS)
	return _cba
}

// Enable LTV enables the specified signature. The signing certificate
// chain is extracted from the signature dictionary. Optionally, additional
// certificates can be specified through the `extraCerts` parameter.
// The LTV client attempts to build the certificate chain up to a trusted root
// by downloading any missing certificates.
func (_gedge *LTV) Enable(sig *PdfSignature, extraCerts []*_bag.Certificate) error {
	if _cefef := _gedge.validateSig(sig); _cefef != nil {
		return _cefef
	}
	_gefc, _daeb := _gedge.generateVRIKey(sig)
	if _daeb != nil {
		return _daeb
	}
	if _, _acbgb := _gedge._fadg.VRI[_gefc]; _acbgb && _gedge.SkipExisting {
		return nil
	}
	_ccae, _daeb := sig.GetCerts()
	if _daeb != nil {
		return _daeb
	}
	return _gedge.enable(_ccae, extraCerts, _gefc)
}

// PdfFunctionType3 defines stitching of the subdomains of several 1-input functions to produce
// a single new 1-input function.
type PdfFunctionType3 struct {
	Domain    []float64
	Range     []float64
	Functions []PdfFunction
	Bounds    []float64
	Encode    []float64
	_ddgdg    *_eb.PdfIndirectObject
}

// DecodeArray returns the range of color component values in DeviceCMYK colorspace.
func (_fabe *PdfColorspaceDeviceCMYK) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// AddKChild adds a child K dictionary object.
func (_cdedb *KDict) AddKChild(kChild *KDict) {
	_cdedb._eegge = append(_cdedb._eegge, &KValue{_ccbca: kChild})
}

// Size returns the width and the height of the page. The method reports
// the page dimensions as displayed by a PDF viewer (i.e. page rotation is
// taken into account).
func (_gbgef *PdfPage) Size() (float64, float64, error) {
	_eegbg, _becfg := _gbgef.GetMediaBox()
	if _becfg != nil {
		return 0, 0, _becfg
	}
	_dbdge, _gdbfe := _eegbg.Width(), _eegbg.Height()
	_baaec, _becfg := _gbgef.GetRotate()
	if _becfg != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _becfg.Error())
	}
	if _egffc := _baaec; _egffc%360 != 0 && _egffc%90 == 0 {
		if _fdade := (360 + _egffc%360) % 360; _fdade == 90 || _fdade == 270 {
			_dbdge, _gdbfe = _gdbfe, _dbdge
		}
	}
	return _dbdge, _gdbfe, nil
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics = _fg.CharMetrics

// GetRefObject returns the reference object of the KValue.
func (_dcee *KValue) GetRefObject() _eb.PdfObject { return _dcee._fbgaa }

// BytesToCharcodes converts the bytes in a PDF string to character codes.
func (_bdabe *PdfFont) BytesToCharcodes(data []byte) []_fc.CharCode {
	_ddb.Log.Trace("\u0042\u0079\u0074es\u0054\u006f\u0043\u0068\u0061\u0072\u0063\u006f\u0064e\u0073:\u0020d\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", data, data)
	if _acebd, _cfgfb := _bdabe._fdaa.(*pdfFontType0); _cfgfb && _acebd._ebeb != nil {
		if _bgeb, _efca := _acebd.bytesToCharcodes(data); _efca {
			return _bgeb
		}
	}
	var (
		_edae  = make([]_fc.CharCode, 0, len(data)+len(data)%2)
		_edaff = _bdabe.baseFields()
	)
	if _edaff._bgbg != nil {
		if _bdaa, _gdgfa := _edaff._bgbg.BytesToCharcodes(data); _gdgfa {
			for _, _edgde := range _bdaa {
				_edae = append(_edae, _fc.CharCode(_edgde))
			}
			return _edae
		}
	}
	if _edaff.isCIDFont() {
		if len(data) == 1 {
			data = []byte{0, data[0]}
		}
		if len(data)%2 != 0 {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0064\u0061\u0074\u0061\u003d\u0025\u002b\u0076\u0020t\u006f\u0020\u0065\u0076\u0065n\u0020\u006ce\u006e\u0067\u0074\u0068", data)
			data = append(data, 0)
		}
		for _eedec := 0; _eedec < len(data); _eedec += 2 {
			_aegcb := uint16(data[_eedec])<<8 | uint16(data[_eedec+1])
			_edae = append(_edae, _fc.CharCode(_aegcb))
		}
	} else {
		for _, _gdfea := range data {
			_edae = append(_edae, _fc.CharCode(_gdfea))
		}
	}
	return _edae
}

// AddMCIDChild adds a child MCID object.
func (_dffeg *KDict) AddMCIDChild(mcid int) {
	_dffeg._eegge = append(_dffeg._eegge, &KValue{_ddaf: &mcid})
}

// Resample resamples the image data converting from current BitsPerComponent to a target BitsPerComponent
// value.  Sets the image's BitsPerComponent to the target value following resampling.
//
// For example, converting an 8-bit RGB image to 1-bit grayscale (common for scanned images):
//
//	// Convert RGB image to grayscale.
//	rgbColorSpace := pdf.NewPdfColorspaceDeviceRGB()
//	grayImage, err := rgbColorSpace.ImageToGray(rgbImage)
//	if err != nil {
//	  return err
//	}
//	// Resample as 1 bit.
//	grayImage.Resample(1)
func (_aaeg *Image) Resample(targetBitsPerComponent int64) {
	if _aaeg.BitsPerComponent == targetBitsPerComponent {
		return
	}
	_bgba := _aaeg.GetSamples()
	if targetBitsPerComponent < _aaeg.BitsPerComponent {
		_cbbfb := _aaeg.BitsPerComponent - targetBitsPerComponent
		for _ccacd := range _bgba {
			_bgba[_ccacd] >>= uint(_cbbfb)
		}
	} else if targetBitsPerComponent > _aaeg.BitsPerComponent {
		_agfd := targetBitsPerComponent - _aaeg.BitsPerComponent
		for _febdb := range _bgba {
			_bgba[_febdb] <<= uint(_agfd)
		}
	}
	_aaeg.BitsPerComponent = targetBitsPerComponent
	if _aaeg.BitsPerComponent < 8 {
		_aaeg.resampleLowBits(_bgba)
		return
	}
	_gbgd := _df.BytesPerLine(int(_aaeg.Width), int(_aaeg.BitsPerComponent), _aaeg.ColorComponents)
	_befe := make([]byte, _gbgd*int(_aaeg.Height))
	var (
		_cgddc, _ggdd, _bcfde, _fddcd int
		_bgdag                        uint32
	)
	for _bcfde = 0; _bcfde < int(_aaeg.Height); _bcfde++ {
		_cgddc = _bcfde * _gbgd
		_ggdd = (_bcfde+1)*_gbgd - 1
		_bcac := _bb.ResampleUint32(_bgba[_cgddc:_ggdd], int(targetBitsPerComponent), 8)
		for _fddcd, _bgdag = range _bcac {
			_befe[_fddcd+_cgddc] = byte(_bgdag)
		}
	}
	_aaeg.Data = _befe
}
func (_cecea *PdfSignature) extractChainFromPKCS7() ([]*_bag.Certificate, error) {
	_fcgef, _eadaf := _dg.Parse(_cecea.Contents.Bytes())
	if _eadaf != nil {
		return nil, _eadaf
	}
	return _fcgef.Certificates, nil
}

// GetCatalogMetadata gets the catalog defined XMP Metadata.
func (_bcaec *PdfReader) GetCatalogMetadata() (_eb.PdfObject, bool) {
	if _bcaec._bagcfd == nil {
		return nil, false
	}
	_becc := _bcaec._bagcfd.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _becc, _becc != nil
}

// PdfColorspaceSpecialPattern is a Pattern colorspace.
// Can be defined either as /Pattern or with an underlying colorspace [/Pattern cs].
type PdfColorspaceSpecialPattern struct {
	UnderlyingCS PdfColorspace
	_fgff        *_eb.PdfIndirectObject
}

// PdfColorspaceDeviceGray represents a grayscale colorspace.
type PdfColorspaceDeviceGray struct{}

// NewPdfAnnotationStrikeOut returns a new text strikeout annotation.
func NewPdfAnnotationStrikeOut() *PdfAnnotationStrikeOut {
	_gebf := NewPdfAnnotation()
	_eac := &PdfAnnotationStrikeOut{}
	_eac.PdfAnnotation = _gebf
	_eac.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gebf.SetContext(_eac)
	return _eac
}

// IsCheckbox returns true if the button field represents a checkbox, false otherwise.
func (_beac *PdfFieldButton) IsCheckbox() bool { return _beac.GetType() == ButtonTypeCheckbox }

// ToPdfObject implements interface PdfModel.
func (_edge *PdfAnnotationPrinterMark) ToPdfObject() _eb.PdfObject {
	_edge.PdfAnnotation.ToPdfObject()
	_fdcf := _edge._ggf
	_aade := _fdcf.PdfObject.(*_eb.PdfObjectDictionary)
	_aade.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b"))
	_aade.SetIfNotNil("\u004d\u004e", _edge.MN)
	return _fdcf
}

// GetMediaBox gets the inheritable media box value, either from the page
// or a higher up page/pages struct.
func (_cbbgef *PdfPage) GetMediaBox() (*PdfRectangle, error) {
	if _cbbgef.MediaBox != nil {
		return _cbbgef.MediaBox, nil
	}
	_edgbb := _cbbgef.Parent
	for _edgbb != nil {
		_bdcabd, _cfbge := _eb.GetDict(_edgbb)
		if !_cfbge {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _cegae := _bdcabd.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _cegae != nil {
			_fbfge, _adfbe := _eb.GetArray(_cegae)
			if !_adfbe {
				return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0064\u0069a\u0020\u0062\u006f\u0078")
			}
			_dggbd, _gfcea := NewPdfRectangle(*_fbfge)
			if _gfcea != nil {
				return nil, _gfcea
			}
			return _dggbd, nil
		}
		_edgbb = _bdcabd.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, _dcf.New("m\u0065\u0064\u0069\u0061 b\u006fx\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
}

// ToPdfObject returns a stream object.
func (_cdefa *XObjectForm) ToPdfObject() _eb.PdfObject {
	_cgfac := _cdefa._afcag
	_fefaa := _cgfac.PdfObjectDictionary
	if _cdefa.Filter != nil {
		_fefaa = _cdefa.Filter.MakeStreamDict()
		_cgfac.PdfObjectDictionary = _fefaa
	}
	_fefaa.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_fefaa.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0046\u006f\u0072\u006d"))
	_fefaa.SetIfNotNil("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065", _cdefa.FormType)
	_fefaa.SetIfNotNil("\u0042\u0042\u006f\u0078", _cdefa.BBox)
	_fefaa.SetIfNotNil("\u004d\u0061\u0074\u0072\u0069\u0078", _cdefa.Matrix)
	if _cdefa.Resources != nil {
		_fefaa.SetIfNotNil("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _cdefa.Resources.ToPdfObject())
	}
	_fefaa.SetIfNotNil("\u0047\u0072\u006fu\u0070", _cdefa.Group)
	_fefaa.SetIfNotNil("\u0052\u0065\u0066", _cdefa.Ref)
	_fefaa.SetIfNotNil("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061", _cdefa.MetaData)
	_fefaa.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _cdefa.PieceInfo)
	_fefaa.SetIfNotNil("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _cdefa.LastModified)
	_fefaa.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _cdefa.StructParent)
	_fefaa.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _cdefa.StructParents)
	_fefaa.SetIfNotNil("\u004f\u0050\u0049", _cdefa.OPI)
	_fefaa.SetIfNotNil("\u004f\u0043", _cdefa.OC)
	_fefaa.SetIfNotNil("\u004e\u0061\u006d\u0065", _cdefa.Name)
	_fefaa.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_cdefa.Stream))))
	_cgfac.Stream = _cdefa.Stream
	return _cgfac
}

// GetStructParentsKey returns the StructParents key.
// If not set, returns -1.
func (_aeega *PdfPage) GetStructParentsKey() int {
	if _cacca, _bbcgc := _eb.GetIntVal(_aeega.StructParents); _bbcgc {
		return _cacca
	}
	return -1
}

// Add appends an outline item as a child of the current outline item.
func (_bfdce *OutlineItem) Add(item *OutlineItem) { _bfdce.Entries = append(_bfdce.Entries, item) }

// NewPdfField returns an initialized PdfField.
func NewPdfField() *PdfField { return &PdfField{_adgda: _eb.MakeIndirectObject(_eb.MakeDict())} }

// ToWriter creates a new writer from the current reader, based on the specified options.
// If no options are provided, all reader properties are copied to the writer.
func (_fegfa *PdfReader) ToWriter(opts *ReaderToWriterOpts) (*PdfWriter, error) {
	_bdfgd := NewPdfWriter()
	_bdfgd.SetFileName(_fegfa._fbgfg)
	if opts == nil {
		opts = &ReaderToWriterOpts{}
	}
	_eedb, _bbeb := _fegfa.GetNumPages()
	if _bbeb != nil {
		_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bbeb)
		return nil, _bbeb
	}
	for _gdebd := 1; _gdebd <= _eedb; _gdebd++ {
		_gafcg, _adeeg := _fegfa.GetPage(_gdebd)
		if _adeeg != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _adeeg)
			return nil, _adeeg
		}
		if opts.PageProcessCallback != nil {
			_adeeg = opts.PageProcessCallback(_gdebd, _gafcg)
			if _adeeg != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _adeeg)
				return nil, _adeeg
			}
		} else if opts.PageCallback != nil {
			opts.PageCallback(_gdebd, _gafcg)
		}
		_adeeg = _bdfgd.AddPage(_gafcg)
		if _adeeg != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _adeeg)
			return nil, _adeeg
		}
	}
	_bdfgd._edbbf = _fegfa.PdfVersion()
	if !opts.SkipInfo {
		_cafgf, _dbgae := _fegfa.GetPdfInfo()
		if _dbgae != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dbgae)
		} else {
			_bdfgd._ecagd.PdfObject = _cafgf.ToPdfObject()
		}
	}
	if !opts.SkipMetadata {
		if _befbb := _fegfa._bagcfd.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _befbb != nil {
			if _dcadad := _bdfgd.SetCatalogMetadata(_befbb); _dcadad != nil {
				return nil, _dcadad
			}
		}
	}
	if !opts.SkipMarkInfo {
		if _ecac, _becff := _fegfa.GetCatalogMarkInfo(); _becff {
			if _deac := _bdfgd.SetCatalogMarkInfo(_ecac); _deac != nil {
				return nil, _deac
			}
		}
	}
	if !opts.SkipAcroForm {
		_gdbdd := _bdfgd.SetForms(_fegfa.AcroForm)
		if _gdbdd != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gdbdd)
			return nil, _gdbdd
		}
	}
	if !opts.SkipViewerPreferences {
		if _gcbec, _aafd := _fegfa.GetCatalogViewerPreferences(); _aafd {
			if _cfee := _bdfgd.SetCatalogViewerPreferences(_gcbec); _cfee != nil {
				return nil, _cfee
			}
		}
	}
	if !opts.SkipLanguage {
		if _ebcfd, _cddbd := _fegfa.GetCatalogLanguage(); _cddbd {
			if _bfdbee := _bdfgd.SetCatalogLanguage(_ebcfd); _bfdbee != nil {
				return nil, _bfdbee
			}
		}
	}
	if !opts.SkipOutlines {
		_bdfgd.AddOutlineTree(_fegfa.GetOutlineTree())
	}
	if !opts.SkipOCProperties {
		_aebeb, _bffec := _fegfa.GetOCProperties()
		if _bffec != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bffec)
		} else {
			_bffec = _bdfgd.SetOCProperties(_aebeb)
			if _bffec != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bffec)
			}
		}
	}
	if !opts.SkipPageLabels {
		_cbdec, _bedee := _fegfa.GetPageLabels()
		if _bedee != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bedee)
		} else {
			_bedee = _bdfgd.SetPageLabels(_cbdec)
			if _bedee != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bedee)
			}
		}
	}
	if !opts.SkipNamedDests {
		_ccgda, _cadaf := _fegfa.GetNamedDestinations()
		if _cadaf != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cadaf)
		} else {
			_cadaf = _bdfgd.SetNamedDestinations(_ccgda)
			if _cadaf != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cadaf)
			}
		}
	}
	if !opts.SkipNameDictionary {
		_agfdf, _egdcd := _fegfa.GetNameDictionary()
		if _egdcd != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _egdcd)
		} else {
			_egdcd = _bdfgd.SetNameDictionary(_agfdf)
			if _egdcd != nil {
				_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _egdcd)
			}
		}
	}
	_afccf, _fgfec := _fegfa.GetCatalogStructTreeRoot()
	if _fgfec {
		_dgbb := _bdfgd.SetCatalogStructTreeRoot(_afccf)
		if _dgbb != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgbb)
		}
	}
	if !opts.SkipRotation && _fegfa.Rotate != nil {
		if _beddb := _bdfgd.SetRotation(*_fegfa.Rotate); _beddb != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _beddb)
		}
	}
	return &_bdfgd, nil
}

// PdfShadingPatternType2 is shading patterns that will use a Type 2 shading pattern (Axial).
type PdfShadingPatternType2 struct {
	*PdfPattern
	Shading   *PdfShadingType2
	Matrix    *_eb.PdfObjectArray
	ExtGState _eb.PdfObject
}

// DefaultFont returns the default font, which is currently the built in Helvetica.
func DefaultFont() *PdfFont {
	_ebgde, _aacf := _fg.NewStdFontByName(HelveticaName)
	if !_aacf {
		panic("\u0048\u0065lv\u0065\u0074\u0069c\u0061\u0020\u0073\u0068oul\u0064 a\u006c\u0077\u0061\u0079\u0073\u0020\u0062e \u0061\u0076\u0061\u0069\u006c\u0061\u0062l\u0065")
	}
	_ffgf := _fbfae(_ebgde)
	return &PdfFont{_fdaa: &_ffgf}
}

// NewPdfShadingPatternType2 creates an empty shading pattern type 2 object.
func NewPdfShadingPatternType2() *PdfShadingPatternType2 {
	_bfddba := &PdfShadingPatternType2{}
	_bfddba.Matrix = _eb.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_bfddba.PdfPattern = &PdfPattern{}
	_bfddba.PdfPattern.PatternType = int64(*_eb.MakeInteger(2))
	_bfddba.PdfPattern._eefgb = _bfddba
	_bfddba.PdfPattern._agddd = _eb.MakeIndirectObject(_eb.MakeDict())
	return _bfddba
}

// FieldFlattenOpts defines a set of options which can be used to configure
// the field flattening process.
type FieldFlattenOpts struct {

	// FilterFunc allows filtering the form fields used in the flattening
	// process. If the filter function returns true, the field is flattened,
	// otherwise it is skipped.
	// If a non-terminal field is discarded, all of its children (the fields
	// present in the Kids array) are discarded as well.
	// Non-terminal fields are kept in the AcroForm if one or more of their
	// child fields have not been selected for flattening.
	// If a filter function is not provided, all form fields are flattened.
	FilterFunc FieldFilterFunc

	// AnnotFilterFunc allows filtering the annotations in the flattening
	// process. If the filter function returns true, the annotation is flattened,
	// otherwise it is skipped.
	AnnotFilterFunc AnnotFilterFunc
}

// AllFields returns a flattened list of all fields in the form.
func (_cadeg *PdfAcroForm) AllFields() []*PdfField {
	if _cadeg == nil {
		return nil
	}
	var _eecdc []*PdfField
	if _cadeg.Fields != nil {
		for _, _eabg := range *_cadeg.Fields {
			_eecdc = append(_eecdc, _aaccff(_eabg)...)
		}
	}
	return _eecdc
}

// PageFromIndirectObject returns the PdfPage and page number for a given indirect object.
func (_aabdd *PdfReader) PageFromIndirectObject(ind *_eb.PdfIndirectObject) (*PdfPage, int, error) {
	if len(_aabdd.PageList) != len(_aabdd._cbaff) {
		return nil, 0, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	for _degae, _ggfdgd := range _aabdd._cbaff {
		if _ggfdgd == ind {
			return _aabdd.PageList[_degae], _degae + 1, nil
		}
	}
	return nil, 0, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// PdfAnnotationWatermark represents Watermark annotations.
// (Section 12.5.6.22).
type PdfAnnotationWatermark struct {
	*PdfAnnotation
	FixedPrint _eb.PdfObject
}

func (_gcd *PdfReader) newPdfAnnotationScreenFromDict(_bcgge *_eb.PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	_cca := PdfAnnotationScreen{}
	_cca.T = _bcgge.Get("\u0054")
	_cca.MK = _bcgge.Get("\u004d\u004b")
	_cca.A = _bcgge.Get("\u0041")
	_cca.AA = _bcgge.Get("\u0041\u0041")
	return &_cca, nil
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_gegbc pdfCIDFontType0) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	return _fg.CharMetrics{Wx: _gegbc._dcfg}, true
}
func (_gffge *PdfWriter) setHashIDs(_dbacd _ecd.Hash) error {
	_dfedf := _dbacd.Sum(nil)
	if _gffge._cebgc == "" {
		_gffge._cebgc = _egf.EncodeToString(_dfedf[:8])
	}
	_gffge.setDocumentIDs(_gffge._cebgc, _egf.EncodeToString(_dfedf[8:]))
	return nil
}

// NewPdfAnnotationSquare returns a new square annotation.
func NewPdfAnnotationSquare() *PdfAnnotationSquare {
	_agf := NewPdfAnnotation()
	_gdf := &PdfAnnotationSquare{}
	_gdf.PdfAnnotation = _agf
	_gdf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_agf.SetContext(_gdf)
	return _gdf
}

// GetColorspaceByName returns the colorspace with the specified name from the page resources.
func (_bcfce *PdfPageResources) GetColorspaceByName(keyName _eb.PdfObjectName) (PdfColorspace, bool) {
	_bcbc, _agecd := _bcfce.GetColorspaces()
	if _agecd != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _agecd)
		return nil, false
	}
	if _bcbc == nil {
		return nil, false
	}
	_gdfg, _ffbec := _bcbc.Colorspaces[string(keyName)]
	if !_ffbec {
		return nil, false
	}
	return _gdfg, true
}

// CharcodesToStrings returns the unicode strings corresponding to `charcodes`.
// The int returns are the number of strings and the number of unconvereted codes.
// NOTE: The number of strings returned is equal to the number of charcodes
func (_bgacde *PdfFont) CharcodesToStrings(charcodes []_fc.CharCode, replacementText string) ([]string, int, int) {
	_dgaf := _bgacde.baseFields()
	_feaeg := make([]string, 0, len(charcodes))
	_ceffc := 0
	_ccagd := _bgacde.Encoder()
	_gcge := _dgaf._bgbg != nil && _bgacde.IsSimple() && _bgacde.Subtype() == "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" && !_cc.Contains(_dgaf._bgbg.Name(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")
	if !_gcge && _ccagd != nil {
		switch _bebe := _ccagd.(type) {
		case _fc.SimpleEncoder:
			_bgga := _bebe.BaseName()
			if _, _caadd := _bcaad[_bgga]; _caadd {
				for _, _egabb := range charcodes {
					if _ddcg, _geeb := _ccagd.CharcodeToRune(_egabb); _geeb {
						_feaeg = append(_feaeg, string(_ddcg))
					} else {
						_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _egabb, charcodes, _dgaf.isCIDFont(), _bgacde, _ccagd)
						_ceffc++
						_feaeg = append(_feaeg, _ff.MissingCodeString)
					}
				}
				return _feaeg, len(_feaeg), _ceffc
			}
		}
	}
	for _, _ebdg := range charcodes {
		if _dgaf._bgbg != nil {
			if _cefee, _edd := _dgaf._bgbg.CharcodeToUnicode(_ff.CharCode(_ebdg)); _edd {
				_fegc, _ := _ae.DecodeLastRuneInString(_cefee)
				_fffc := _ggaac(_fegc)
				if !(_fffc == "\u0043\u006f") {
					_feaeg = append(_feaeg, _cefee)
					continue
				}
				_ddb.Log.Debug("E\u0052\u0052\u004fR\u003a\u0020\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0043\u006d\u0061p\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067.\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0069\u0073\u0020m\u0061\u0070\u0070\u0065\u0064 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0063\u006f\u0064\u0065 \u0070\u006f\u0069\u006e\u0074\u0020\u0025\u0073", _ebdg, _cefee)
			}
		}
		if _ccagd != nil {
			if _gcfca, _aggb := _ccagd.CharcodeToRune(_ebdg); _aggb {
				_cadfg := _ggaac(_gcfca)
				if !(_cadfg == "\u0043\u006f") {
					_feaeg = append(_feaeg, string(_gcfca))
					continue
				}
				_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e \u0063\u006f\u0064\u0065\u003d\u0030x\u0025\u0030\u0034\u0078\u0020\u0069\u0073\u0020\u0064\u0065\u0063\u006f\u0064\u0065d\u0020\u0074o\u0020\u0061\u006e\u0020i\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0075\u006e\u0020\u0025\u0073", _ebdg, string(_gcfca))
			}
		}
		if replacementText != "" {
			_feaeg = append(_feaeg, replacementText)
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _ebdg, charcodes, _dgaf.isCIDFont(), _bgacde, _ccagd)
			_ceffc++
			_feaeg = append(_feaeg, _ff.MissingCodeString)
		}
	}
	if _ceffc != 0 {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0043\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0074\u006f\u0020u\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020i\u006ep\u0075\u0074\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0043\u0068\u0061\u0072\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073", len(charcodes), _ceffc, _bgacde)
	}
	return _feaeg, len(_feaeg), _ceffc
}

// IsHideWindowUI returns the value of the hideWindowUI flag.
func (_abacc *ViewerPreferences) IsHideWindowUI() bool {
	if _abacc._fcebae == nil {
		return false
	}
	return *_abacc._fcebae
}

// PdfAnnotationScreen represents Screen annotations.
// (Section 12.5.6.18).
type PdfAnnotationScreen struct {
	*PdfAnnotation
	T  _eb.PdfObject
	MK _eb.PdfObject
	A  _eb.PdfObject
	AA _eb.PdfObject
}

// StandardApplier is the interface that performs optimization of the whole PDF document.
// As a result an input document is being changed by the optimizer.
// The writer than takes back all it's parts and overwrites it.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardApplier interface {
	ApplyStandard(_dbec *_cd.Document) error
}

// Hasher is the interface that wraps the basic Write method.
type Hasher interface {
	Write(_bbed []byte) (_ecbgd int, _bcbdaf error)
}

// HasFontByName checks if has font resource by name.
func (_fedac *PdfPage) HasFontByName(name _eb.PdfObjectName) bool {
	_eaef, _bdbac := _fedac.Resources.Font.(*_eb.PdfObjectDictionary)
	if !_bdbac {
		return false
	}
	if _aecfb := _eaef.Get(name); _aecfb != nil {
		return true
	}
	return false
}
func _bfbdg(_eddf *_eb.PdfObjectDictionary, _feffd *fontCommon, _ccgbg _fc.TextEncoder) (*pdfFontSimple, error) {
	_aaaae := _abbaa(_feffd)
	_aaaae._gcgb = _ccgbg
	if _ccgbg == nil {
		_cega := _eddf.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
		if _cega == nil {
			_cega = _eb.MakeInteger(0)
		}
		_aaaae.FirstChar = _cega
		_gedf, _cgfdb := _eb.GetIntVal(_cega)
		if !_cgfdb {
			_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _cega)
			return nil, _eb.ErrTypeError
		}
		_bgdc := _fc.CharCode(_gedf)
		_cega = _eddf.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
		if _cega == nil {
			_cega = _eb.MakeInteger(255)
		}
		_aaaae.LastChar = _cega
		_gedf, _cgfdb = _eb.GetIntVal(_cega)
		if !_cgfdb {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _cega)
			return nil, _eb.ErrTypeError
		}
		_fecc := _fc.CharCode(_gedf)
		_aaaae._cegda = make(map[_fc.CharCode]float64)
		_cega = _eddf.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _cega != nil {
			_aaaae.Widths = _cega
			_gaebg, _cebb := _eb.GetArray(_cega)
			if !_cebb {
				_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _cega)
				return nil, _eb.ErrTypeError
			}
			_ecef, _gedcb := _gaebg.ToFloat64Array()
			if _gedcb != nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
				return nil, _gedcb
			}
			if len(_ecef) != int(_fecc-_bgdc+1) {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _fecc-_bgdc+1, len(_ecef))
				return nil, _eb.ErrRangeError
			}
			for _gedb, _dddd := range _ecef {
				_aaaae._cegda[_bgdc+_fc.CharCode(_gedb)] = _dddd
			}
		}
	}
	_aaaae.Encoding = _eb.TraceToDirectObject(_eddf.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	return _aaaae, nil
}

// NewPdfShadingType3 creates an empty shading type 3 dictionary.
func NewPdfShadingType3() *PdfShadingType3 {
	_fdgc := &PdfShadingType3{}
	_fdgc.PdfShading = &PdfShading{}
	_fdgc.PdfShading._cefaa = _eb.MakeIndirectObject(_eb.MakeDict())
	_fdgc.PdfShading._ecffg = _fdgc
	return _fdgc
}

// NewPdfActionRendition returns a new "rendition" action.
func NewPdfActionRendition() *PdfActionRendition {
	_ge := NewPdfAction()
	_geb := &PdfActionRendition{}
	_geb.PdfAction = _ge
	_ge.SetContext(_geb)
	return _geb
}

// PrintScaling returns the value of the printScaling.
func (_eadaa *ViewerPreferences) PrintScaling() PrintScaling { return _eadaa._gefcg }

// ToPdfObject converts rectangle to a PDF object.
func (_fdbcg *PdfRectangle) ToPdfObject() _eb.PdfObject {
	return _eb.MakeArray(_eb.MakeFloat(_fdbcg.Llx), _eb.MakeFloat(_fdbcg.Lly), _eb.MakeFloat(_fdbcg.Urx), _eb.MakeFloat(_fdbcg.Ury))
}

// PdfFieldChoice represents a choice field which includes scrollable list boxes and combo boxes.
type PdfFieldChoice struct {
	*PdfField
	Opt *_eb.PdfObjectArray
	TI  *_eb.PdfObjectInteger
	I   *_eb.PdfObjectArray
}

// PdfActionRendition represents a Rendition action.
type PdfActionRendition struct {
	*PdfAction
	R  _eb.PdfObject
	AN _eb.PdfObject
	OP _eb.PdfObject
	JS _eb.PdfObject
}

// Height returns the height of `rect`.
func (_cdbce *PdfRectangle) Height() float64 { return _gg.Abs(_cdbce.Ury - _cdbce.Lly) }

// DetermineColorspaceNameFromPdfObject determines PDF colorspace from a PdfObject.  Returns the colorspace name and
// an error on failure. If the colorspace was not found, will return an empty string.
func DetermineColorspaceNameFromPdfObject(obj _eb.PdfObject) (_eb.PdfObjectName, error) {
	var _gec *_eb.PdfObjectName
	var _eceg *_eb.PdfObjectArray
	if _gcgf, _agcg := obj.(*_eb.PdfIndirectObject); _agcg {
		if _ffe, _adgd := _gcgf.PdfObject.(*_eb.PdfObjectArray); _adgd {
			_eceg = _ffe
		} else if _egbe, _fdcfa := _gcgf.PdfObject.(*_eb.PdfObjectName); _fdcfa {
			_gec = _egbe
		}
	} else if _eeca, _aadb := obj.(*_eb.PdfObjectArray); _aadb {
		_eceg = _eeca
	} else if _geec, _deec := obj.(*_eb.PdfObjectName); _deec {
		_gec = _geec
	}
	if _gec != nil {
		switch *_gec {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_gec, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_gec, nil
		}
	}
	if _eceg != nil && _eceg.Len() > 0 {
		if _acbb, _fgcde := _eceg.Get(0).(*_eb.PdfObjectName); _fgcde {
			switch *_acbb {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _eceg.Len() == 1 {
					return *_acbb, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_acbb, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_acbb, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_acbb, nil
			}
		}
	}
	return "", nil
}

// NewPdfAnnotationSound returns a new sound annotation.
func NewPdfAnnotationSound() *PdfAnnotationSound {
	_fdd := NewPdfAnnotation()
	_efb := &PdfAnnotationSound{}
	_efb.PdfAnnotation = _fdd
	_efb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdd.SetContext(_efb)
	return _efb
}

// NewPdfFontFromTTFFile loads a TTF font file and returns a PdfFont type
// that can be used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTFFile.
func NewPdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_dfgd, _ebcbf := _ccb.Open(filePath)
	if _ebcbf != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020T\u0054F\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _ebcbf)
		return nil, _ebcbf
	}
	defer _dfgd.Close()
	return NewPdfFontFromTTF(_dfgd)
}

// PdfAnnotationFreeText represents FreeText annotations.
// (Section 12.5.6.6).
type PdfAnnotationFreeText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	DA _eb.PdfObject
	Q  _eb.PdfObject
	RC _eb.PdfObject
	DS _eb.PdfObject
	CL _eb.PdfObject
	IT _eb.PdfObject
	BE _eb.PdfObject
	RD _eb.PdfObject
	BS _eb.PdfObject
	LE _eb.PdfObject
}

// ToPdfOutline returns a low level PdfOutline object, based on the current
// instance.
func (_cgacbg *Outline) ToPdfOutline() *PdfOutline {
	_dbfbe := NewPdfOutline()
	var _bfbg []*PdfOutlineItem
	var _eeefa int64
	var _egbd *PdfOutlineItem
	for _, _ggee := range _cgacbg.Entries {
		_bffad, _bcffe := _ggee.ToPdfOutlineItem()
		_bffad.Parent = &_dbfbe.PdfOutlineTreeNode
		if _egbd != nil {
			_egbd.Next = &_bffad.PdfOutlineTreeNode
			_bffad.Prev = &_egbd.PdfOutlineTreeNode
		}
		_bfbg = append(_bfbg, _bffad)
		_eeefa += _bcffe
		_egbd = _bffad
	}
	_ddgbf := int64(len(_bfbg))
	_eeefa += _ddgbf
	if _ddgbf > 0 {
		_dbfbe.First = &_bfbg[0].PdfOutlineTreeNode
		_dbfbe.Last = &_bfbg[_ddgbf-1].PdfOutlineTreeNode
		_dbfbe.Count = &_eeefa
	}
	return _dbfbe
}
func (_adegg *PdfPage) setContainer(_cggfe *_eb.PdfIndirectObject) {
	_cggfe.PdfObject = _adegg._aaagaa
	_adegg._efcff = _cggfe
}
func (_dfa *PdfReader) newPdfAnnotationFileAttachmentFromDict(_dfaa *_eb.PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	_fga := PdfAnnotationFileAttachment{}
	_dcae, _gfab := _dfa.newPdfAnnotationMarkupFromDict(_dfaa)
	if _gfab != nil {
		return nil, _gfab
	}
	_fga.PdfAnnotationMarkup = _dcae
	_fga.FS = _dfaa.Get("\u0046\u0053")
	_fga.Name = _dfaa.Get("\u004e\u0061\u006d\u0065")
	return &_fga, nil
}

// FullName returns the full name of the field as in rootname.parentname.partialname.
func (_ddaa *PdfField) FullName() (string, error) {
	var _fabc _dd.Buffer
	_baga := []string{}
	if _ddaa.T != nil {
		_baga = append(_baga, _ddaa.T.Decoded())
	}
	_cfbf := map[*PdfField]bool{}
	_cfbf[_ddaa] = true
	_agab := _ddaa.Parent
	for _agab != nil {
		if _, _eedef := _cfbf[_agab]; _eedef {
			return _fabc.String(), _dcf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		if _agab.T == nil {
			return _fabc.String(), _dcf.New("\u0066\u0069el\u0064\u0020\u0070a\u0072\u0074\u0069\u0061l n\u0061me\u0020\u0028\u0054\u0029\u0020\u006e\u006ft \u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064")
		}
		_baga = append(_baga, _agab.T.Decoded())
		_cfbf[_agab] = true
		_agab = _agab.Parent
	}
	for _ddcca := len(_baga) - 1; _ddcca >= 0; _ddcca-- {
		_fabc.WriteString(_baga[_ddcca])
		if _ddcca > 0 {
			_fabc.WriteString("\u002e")
		}
	}
	return _fabc.String(), nil
}

// Decrypt decrypts the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful,
// false otherwise.
func (_ggfc *PdfReader) Decrypt(password []byte) (bool, error) {
	_bbfg, _abecc := _ggfc._ebbe.Decrypt(password)
	if _abecc != nil {
		return false, _abecc
	}
	if !_bbfg {
		return false, nil
	}
	_abecc = _ggfc.loadStructure()
	if _abecc != nil {
		_ddb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f \u006co\u0061d\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", _abecc)
		return false, _abecc
	}
	return true, nil
}

// ToPdfObject returns a *PdfIndirectObject containing a *PdfObjectArray representation of the DeviceN colorspace.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
	    or: [/DeviceN names alternateSpace tintTransform attributes]
*/
func (_beece *PdfColorspaceDeviceN) ToPdfObject() _eb.PdfObject {
	_bfed := _eb.MakeArray(_eb.MakeName("\u0044e\u0076\u0069\u0063\u0065\u004e"))
	_bfed.Append(_beece.ColorantNames)
	_bfed.Append(_beece.AlternateSpace.ToPdfObject())
	_bfed.Append(_beece.TintTransform.ToPdfObject())
	if _beece.Attributes != nil {
		_bfed.Append(_beece.Attributes.ToPdfObject())
	}
	if _beece._cfef != nil {
		_beece._cfef.PdfObject = _bfed
		return _beece._cfef
	}
	return _bfed
}

// PdfOutlineItem represents an outline item dictionary (Table 153 - pp. 376 - 377).
type PdfOutlineItem struct {
	PdfOutlineTreeNode
	Title  *_eb.PdfObjectString
	Parent *PdfOutlineTreeNode
	Prev   *PdfOutlineTreeNode
	Next   *PdfOutlineTreeNode
	Count  *int64
	Dest   _eb.PdfObject
	A      _eb.PdfObject
	SE     _eb.PdfObject
	C      _eb.PdfObject
	F      _eb.PdfObject
	_efbfb *_eb.PdfIndirectObject
}

// ColorToRGB converts gray -> rgb for a single color component.
func (_ebcd *PdfColorspaceDeviceGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_dbaf, _eacga := color.(*PdfColorDeviceGray)
	if !_eacga {
		_ddb.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006fr\u0020\u006e\u006f\u0074\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0067\u0072\u0061\u0079\u0020\u0025\u0054", color)
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return NewPdfColorDeviceRGB(float64(*_dbaf), float64(*_dbaf), float64(*_dbaf)), nil
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_dbfc *PdfWriter) SetOptimizer(optimizer Optimizer) { _dbfc._fbaag = optimizer }

// GetAction returns the PDF action for the annotation link.
func (_cdff *PdfAnnotationLink) GetAction() (*PdfAction, error) {
	if _cdff._fbg != nil {
		return _cdff._fbg, nil
	}
	if _cdff.A == nil {
		return nil, nil
	}
	if _cdff._aced == nil {
		return nil, nil
	}
	_fged, _aag := _cdff._aced.loadAction(_cdff.A)
	if _aag != nil {
		return nil, _aag
	}
	_cdff._fbg = _fged
	return _cdff._fbg, nil
}

// NonFullScreenPageMode returns the value of the nonFullScreenPageMode.
func (_befbbe *ViewerPreferences) NonFullScreenPageMode() NonFullScreenPageMode {
	return _befbbe._cdcff
}

// AddExtension adds the specified extension to the Extensions dictionary.
// See section 7.1.2 "Extensions Dictionary" (pp. 108-109 PDF32000_2008).
func (_ecgaa *PdfWriter) AddExtension(extName, baseVersion string, extLevel int) {
	_dbgff, _dcbag := _eb.GetDict(_ecgaa._dbffa.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073"))
	if !_dcbag {
		_dbgff = _eb.MakeDict()
		_ecgaa._dbffa.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073", _dbgff)
	}
	_fbggba, _dcbag := _eb.GetDict(_dbgff.Get(_eb.PdfObjectName(extName)))
	if !_dcbag {
		_fbggba = _eb.MakeDict()
		_dbgff.Set(_eb.PdfObjectName(extName), _fbggba)
	}
	if _aadgf, _ := _eb.GetNameVal(_fbggba.Get("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e")); _aadgf != baseVersion {
		_fbggba.Set("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e", _eb.MakeName(baseVersion))
	}
	if _faeeb, _ := _eb.GetIntVal(_fbggba.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c")); _faeeb != extLevel {
		_fbggba.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c", _eb.MakeInteger(int64(extLevel)))
	}
}

// PdfAnnotationInk represents Ink annotations.
// (Section 12.5.6.13).
type PdfAnnotationInk struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	InkList _eb.PdfObject
	BS      _eb.PdfObject
}

func (_aaac *PdfColorspaceDeviceCMYK) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
}

// GetXObjectFormByName returns the XObjectForm with the specified name from the
// page resources, if it exists.
func (_eafc *PdfPageResources) GetXObjectFormByName(keyName _eb.PdfObjectName) (*XObjectForm, error) {
	_fbgae, _ababb := _eafc.GetXObjectByName(keyName)
	if _fbgae == nil {
		return nil, nil
	}
	if _ababb != XObjectTypeForm {
		return nil, _dcf.New("\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006f\u0072\u006d")
	}
	_gadg, _cfeae := NewXObjectFormFromStream(_fbgae)
	if _cfeae != nil {
		return nil, _cfeae
	}
	return _gadg, nil
}
func (_dabeb *PdfWriter) copyObjects() {
	_cbbbbd := make(map[_eb.PdfObject]_eb.PdfObject)
	_aegab := make([]_eb.PdfObject, 0, len(_dabeb._dcfgf))
	_baea := make(map[_eb.PdfObject]struct{}, len(_dabeb._dcfgf))
	_gfcc := make(map[_eb.PdfObject]struct{})
	for _, _acca := range _dabeb._dcfgf {
		_gcbcc := _dabeb.copyObject(_acca, _cbbbbd, _gfcc, false)
		if _, _abffc := _gfcc[_acca]; _abffc {
			continue
		}
		_aegab = append(_aegab, _gcbcc)
		_baea[_gcbcc] = struct{}{}
	}
	_dabeb._dcfgf = _aegab
	_dabeb._aeeda = _baea
	_dabeb._ecagd = _dabeb.copyObject(_dabeb._ecagd, _cbbbbd, nil, false).(*_eb.PdfIndirectObject)
	_dabeb._ccea = _dabeb.copyObject(_dabeb._ccea, _cbbbbd, nil, false).(*_eb.PdfIndirectObject)
	if _dabeb._ebdeed != nil {
		_dabeb._ebdeed = _dabeb.copyObject(_dabeb._ebdeed, _cbbbbd, nil, false).(*_eb.PdfIndirectObject)
	}
	if _dabeb._caed {
		_dffgg := make(map[_eb.PdfObject]int64)
		for _aecgeb, _eedgf := range _dabeb._ecbgf {
			if _cgbaf, _egfab := _cbbbbd[_aecgeb]; _egfab {
				_dffgg[_cgbaf] = _eedgf
			} else {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020a\u0070\u0070\u0065n\u0064\u0020\u006d\u006fd\u0065\u0020\u002d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u0070\u0079\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070")
			}
		}
		_dabeb._ecbgf = _dffgg
	}
}

// NewPdfAnnotationStamp returns a new stamp annotation.
func NewPdfAnnotationStamp() *PdfAnnotationStamp {
	_daed := NewPdfAnnotation()
	_ccge := &PdfAnnotationStamp{}
	_ccge.PdfAnnotation = _daed
	_ccge.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_daed.SetContext(_ccge)
	return _ccge
}

// ToPdfObject implements interface PdfModel.
func (_gae *PdfActionURI) ToPdfObject() _eb.PdfObject {
	_gae.PdfAction.ToPdfObject()
	_ebe := _gae._dee
	_eggc := _ebe.PdfObject.(*_eb.PdfObjectDictionary)
	_eggc.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeURI)))
	_eggc.SetIfNotNil("\u0055\u0052\u0049", _gae.URI)
	_eggc.SetIfNotNil("\u0049\u0073\u004da\u0070", _gae.IsMap)
	return _ebe
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_dbbf *PdfShadingType3) ToPdfObject() _eb.PdfObject {
	_dbbf.PdfShading.ToPdfObject()
	_fcef, _gcfcb := _dbbf.getShadingDict()
	if _gcfcb != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _dbbf.Coords != nil {
		_fcef.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _dbbf.Coords)
	}
	if _dbbf.Domain != nil {
		_fcef.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _dbbf.Domain)
	}
	if _dbbf.Function != nil {
		if len(_dbbf.Function) == 1 {
			_fcef.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dbbf.Function[0].ToPdfObject())
		} else {
			_gbce := _eb.MakeArray()
			for _, _aebg := range _dbbf.Function {
				_gbce.Append(_aebg.ToPdfObject())
			}
			_fcef.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gbce)
		}
	}
	if _dbbf.Extend != nil {
		_fcef.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _dbbf.Extend)
	}
	return _dbbf._cefaa
}
func (_cddcd *PdfWriter) getPdfVersion() string {
	return _e.Sprintf("\u0025\u0064\u002e%\u0064", _cddcd._edbbf.Major, _cddcd._edbbf.Minor)
}

// PdfActionSubmitForm represents a submitForm action.
type PdfActionSubmitForm struct {
	*PdfAction
	F      *PdfFilespec
	Fields _eb.PdfObject
	Flags  _eb.PdfObject
}

// SetHideToolbar sets the value of the hideToolbar flag.
func (_gface *ViewerPreferences) SetHideToolbar(hideToolbar bool) { _gface._caafc = &hideToolbar }

// ToPdfObject implements interface PdfModel.
func (_gfcaa *PdfAnnotationSound) ToPdfObject() _eb.PdfObject {
	_gfcaa.PdfAnnotation.ToPdfObject()
	_eedd := _gfcaa._ggf
	_dgff := _eedd.PdfObject.(*_eb.PdfObjectDictionary)
	_gfcaa.PdfAnnotationMarkup.appendToPdfDictionary(_dgff)
	_dgff.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053\u006f\u0075n\u0064"))
	_dgff.SetIfNotNil("\u0053\u006f\u0075n\u0064", _gfcaa.Sound)
	_dgff.SetIfNotNil("\u004e\u0061\u006d\u0065", _gfcaa.Name)
	return _eedd
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_gdega *PdfShadingPatternType3) ToPdfObject() _eb.PdfObject {
	_gdega.PdfPattern.ToPdfObject()
	_bgab := _gdega.getDict()
	if _gdega.Shading != nil {
		_bgab.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _gdega.Shading.ToPdfObject())
	}
	if _gdega.Matrix != nil {
		_bgab.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _gdega.Matrix)
	}
	if _gdega.ExtGState != nil {
		_bgab.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _gdega.ExtGState)
	}
	return _gdega._agddd
}

// NewKValue creates a new K value object.
func NewKValue() *KValue { return &KValue{} }

// GetRotate gets the inheritable rotate value, either from the page
// or a higher up page/pages struct.
func (_dbedg *PdfPage) GetRotate() (int64, error) {
	if _dbedg.Rotate != nil {
		return *_dbedg.Rotate, nil
	}
	_adadb := _dbedg.Parent
	for _adadb != nil {
		_dagcg, _bdfcc := _eb.GetDict(_adadb)
		if !_bdfcc {
			return 0, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _gfdfe := _dagcg.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _gfdfe != nil {
			_eacea, _gdgfe := _eb.GetInt(_gfdfe)
			if !_gdgfe {
				return 0, _dcf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0074a\u0074\u0065\u0020\u0076al\u0075\u0065")
			}
			if _eacea != nil {
				return int64(*_eacea), nil
			}
			return 0, _dcf.New("\u0072\u006f\u0074\u0061te\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		_adadb = _dagcg.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return 0, _dcf.New("\u0072o\u0074a\u0074\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
}
func _aggce(_gdbeb _eb.PdfObject) (*PdfShading, error) {
	_baed := &PdfShading{}
	var _bcbg *_eb.PdfObjectDictionary
	if _baca, _cedcc := _eb.GetIndirect(_gdbeb); _cedcc {
		_baed._cefaa = _baca
		_bcbge, _aaccb := _baca.PdfObject.(*_eb.PdfObjectDictionary)
		if !_aaccb {
			_ddb.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _eb.ErrTypeError
		}
		_bcbg = _bcbge
	} else if _gadf, _gcgaf := _eb.GetStream(_gdbeb); _gcgaf {
		_baed._cefaa = _gadf
		_bcbg = _gadf.PdfObjectDictionary
	} else if _ebecec, _cagf := _eb.GetDict(_gdbeb); _cagf {
		_baed._cefaa = _ebecec
		_bcbg = _ebecec
	} else {
		_ddb.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", _gdbeb)
		return nil, _eb.ErrTypeError
	}
	if _bcbg == nil {
		_ddb.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _dcf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_gdbeb = _bcbg.Get("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _gdbeb == nil {
		_ddb.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073si\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gdbeb = _eb.TraceToDirectObject(_gdbeb)
	_cdffb, _fgca := _gdbeb.(*_eb.PdfObjectInteger)
	if !_fgca {
		_ddb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066o\u0072 \u0073h\u0061d\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _gdbeb)
		return nil, _eb.ErrTypeError
	}
	if *_cdffb < 1 || *_cdffb > 7 {
		_ddb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u006e\u006ft\u0020\u0031\u002d\u0037\u0020(\u0067\u006ft\u0020\u0025\u0064\u0029", *_cdffb)
		return nil, _eb.ErrTypeError
	}
	_baed.ShadingType = _cdffb
	_gdbeb = _bcbg.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")
	if _gdbeb == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0020\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065\u0020e\u006e\u0074\u0072\u0079\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gebcd, _ecbg := NewPdfColorspaceFromPdfObject(_gdbeb)
	if _ecbg != nil {
		_ddb.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065: \u0025\u0076", _ecbg)
		return nil, _ecbg
	}
	_baed.ColorSpace = _gebcd
	_gdbeb = _bcbg.Get("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064")
	if _gdbeb != nil {
		_gdbeb = _eb.TraceToDirectObject(_gdbeb)
		_ddaef, _cffag := _gdbeb.(*_eb.PdfObjectArray)
		if !_cffag {
			_ddb.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _gdbeb)
			return nil, _eb.ErrTypeError
		}
		_baed.Background = _ddaef
	}
	_gdbeb = _bcbg.Get("\u0042\u0042\u006f\u0078")
	if _gdbeb != nil {
		_gdbeb = _eb.TraceToDirectObject(_gdbeb)
		_deaca, _adacd := _gdbeb.(*_eb.PdfObjectArray)
		if !_adacd {
			_ddb.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _gdbeb)
			return nil, _eb.ErrTypeError
		}
		_aedf, _eacde := NewPdfRectangle(*_deaca)
		if _eacde != nil {
			_ddb.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _eacde)
			return nil, _eacde
		}
		_baed.BBox = _aedf
	}
	_gdbeb = _bcbg.Get("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s")
	if _gdbeb != nil {
		_gdbeb = _eb.TraceToDirectObject(_gdbeb)
		_edfb, _gcde := _gdbeb.(*_eb.PdfObjectBool)
		if !_gcde {
			_ddb.Log.Debug("A\u006e\u0074\u0069\u0041\u006c\u0069\u0061\u0073\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020ty\u0070\u0065\u002c\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020bo\u006f\u006c \u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gdbeb)
			return nil, _eb.ErrTypeError
		}
		_baed.AntiAlias = _edfb
	}
	switch *_cdffb {
	case 1:
		_aaacf, _bcbca := _eebd(_bcbg)
		if _bcbca != nil {
			return nil, _bcbca
		}
		_aaacf.PdfShading = _baed
		_baed._ecffg = _aaacf
		return _baed, nil
	case 2:
		_befdc, _eedaa := _cdcbe(_bcbg)
		if _eedaa != nil {
			return nil, _eedaa
		}
		_befdc.PdfShading = _baed
		_baed._ecffg = _befdc
		return _baed, nil
	case 3:
		_aabdf, _fdfa := _edafd(_bcbg)
		if _fdfa != nil {
			return nil, _fdfa
		}
		_aabdf.PdfShading = _baed
		_baed._ecffg = _aabdf
		return _baed, nil
	case 4:
		_cfgcg, _cdbf := _edcffb(_bcbg)
		if _cdbf != nil {
			return nil, _cdbf
		}
		_cfgcg.PdfShading = _baed
		_baed._ecffg = _cfgcg
		return _baed, nil
	case 5:
		_fbbaf, _fbbec := _edcec(_bcbg)
		if _fbbec != nil {
			return nil, _fbbec
		}
		_fbbaf.PdfShading = _baed
		_baed._ecffg = _fbbaf
		return _baed, nil
	case 6:
		_cgaec, _febaf := _ffeeb(_bcbg)
		if _febaf != nil {
			return nil, _febaf
		}
		_cgaec.PdfShading = _baed
		_baed._ecffg = _cgaec
		return _baed, nil
	case 7:
		_cggdc, _feecb := _facfg(_bcbg)
		if _feecb != nil {
			return nil, _feecb
		}
		_cggdc.PdfShading = _baed
		_baed._ecffg = _cggdc
		return _baed, nil
	}
	return nil, _dcf.New("u\u006ek\u006e\u006f\u0077\u006e\u0020\u0073\u0068\u0061d\u0069\u006e\u0067\u0020ty\u0070\u0065")
}

// GetVersion gets the document version.
func (_dfffa *PdfWriter) GetVersion() _eb.Version { return _dfffa._edbbf }

// PrintScaling represents the page scaling option that shall be selected
// when a print dialog is displayed for this document.
type PrintScaling string

// ToPdfObject converts the PdfFont object to its PDF representation.
func (_ddedf *PdfFont) ToPdfObject() _eb.PdfObject {
	if _ddedf._fdaa == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0066\u006f\u006e\u0074 \u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return _eb.MakeNull()
	}
	return _ddedf._fdaa.ToPdfObject()
}

// NewPdfActionSound returns a new "sound" action.
func NewPdfActionSound() *PdfActionSound {
	_ed := NewPdfAction()
	_fcc := &PdfActionSound{}
	_fcc.PdfAction = _ed
	_ed.SetContext(_fcc)
	return _fcc
}

// GetAnnotations returns the list of page annotations for `page`. If not loaded attempts to load the
// annotations, otherwise returns the loaded list.
func (_accde *PdfPage) GetAnnotations() ([]*PdfAnnotation, error) {
	if _accde._dbga != nil {
		return _accde._dbga, nil
	}
	if _accde.Annots == nil {
		_accde._dbga = []*PdfAnnotation{}
		return nil, nil
	}
	if _accde._dcdfd == nil {
		_accde._dbga = []*PdfAnnotation{}
		return nil, nil
	}
	_cfbgd, _egaef := _accde._dcdfd.loadAnnotations(_accde.Annots)
	if _egaef != nil {
		return nil, _egaef
	}
	if _cfbgd == nil {
		_accde._dbga = []*PdfAnnotation{}
	}
	_accde._dbga = _cfbgd
	return _accde._dbga, nil
}

// PdfFieldText represents a text field where user can enter text.
type PdfFieldText struct {
	*PdfField
	DA     *_eb.PdfObjectString
	Q      *_eb.PdfObjectInteger
	DS     *_eb.PdfObjectString
	RV     _eb.PdfObject
	MaxLen *_eb.PdfObjectInteger
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_eafdc *PdfShadingType2) ToPdfObject() _eb.PdfObject {
	_eafdc.PdfShading.ToPdfObject()
	_egdafe, _bbeef := _eafdc.getShadingDict()
	if _bbeef != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _egdafe == nil {
		_ddb.Log.Error("\u0053\u0068\u0061\u0064in\u0067\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		return nil
	}
	if _eafdc.Coords != nil {
		_egdafe.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _eafdc.Coords)
	}
	if _eafdc.Domain != nil {
		_egdafe.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _eafdc.Domain)
	}
	if _eafdc.Function != nil {
		if len(_eafdc.Function) == 1 {
			_egdafe.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _eafdc.Function[0].ToPdfObject())
		} else {
			_fbdag := _eb.MakeArray()
			for _, _babfe := range _eafdc.Function {
				_fbdag.Append(_babfe.ToPdfObject())
			}
			_egdafe.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _fbdag)
		}
	}
	if _eafdc.Extend != nil {
		_egdafe.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _eafdc.Extend)
	}
	return _eafdc._cefaa
}

// C returns the value of the cyan component of the color.
func (_cbbga *PdfColorDeviceCMYK) C() float64 { return _cbbga[0] }
func _agcd(_acgd _eb.PdfObject) (*PdfColorspaceSpecialSeparation, error) {
	_bdca := NewPdfColorspaceSpecialSeparation()
	if _gfafa, _fgdg := _acgd.(*_eb.PdfIndirectObject); _fgdg {
		_bdca._degb = _gfafa
	}
	_acgd = _eb.TraceToDirectObject(_acgd)
	_dafb, _badgb := _acgd.(*_eb.PdfObjectArray)
	if !_badgb {
		return nil, _e.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062je\u0063\u0074")
	}
	if _dafb.Len() != 4 {
		return nil, _e.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074i\u006f\u006e \u0043\u0053\u003a\u0020\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
	}
	_acgd = _dafb.Get(0)
	_ggad, _badgb := _acgd.(*_eb.PdfObjectName)
	if !_badgb {
		return nil, _e.Errorf("\u0073\u0065\u0070ar\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a \u0069n\u0076a\u006ci\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_ggad != "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e" {
		return nil, _e.Errorf("\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020w\u0072o\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	_acgd = _dafb.Get(1)
	_ggad, _badgb = _acgd.(*_eb.PdfObjectName)
	if !_badgb {
		return nil, _e.Errorf("\u0073\u0065pa\u0072\u0061\u0074i\u006f\u006e\u0020\u0043S: \u0049nv\u0061\u006c\u0069\u0064\u0020\u0063\u006flo\u0072\u0061\u006e\u0074\u0020\u006e\u0061m\u0065")
	}
	_bdca.ColorantName = _ggad
	_acgd = _dafb.Get(2)
	_cdgdd, _fgafe := NewPdfColorspaceFromPdfObject(_acgd)
	if _fgafe != nil {
		return nil, _fgafe
	}
	_bdca.AlternateSpace = _cdgdd
	_eagb, _fgafe := _cccfa(_dafb.Get(3))
	if _fgafe != nil {
		return nil, _fgafe
	}
	_bdca.TintTransform = _eagb
	return _bdca, nil
}

// SetCatalogViewerPreferences sets the catalog ViewerPreferences dictionary.
func (_efcffg *PdfWriter) SetCatalogViewerPreferences(pref _eb.PdfObject) error {
	if pref == nil {
		_efcffg._dbffa.Remove("\u0056\u0069\u0065\u0077\u0065\u0072\u0050\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0073")
		return nil
	}
	if _dbdgd, _deadg := pref.(*_eb.PdfObjectReference); _deadg {
		pref = _dbdgd.Resolve()
		if pref == nil {
			_efcffg._dbffa.Remove("\u0056\u0069\u0065\u0077\u0065\u0072\u0050\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0073")
			return nil
		}
	}
	_efcffg.addObject(pref)
	_efcffg._dbffa.Set("\u0056\u0069\u0065\u0077\u0065\u0072\u0050\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0073", pref)
	return nil
}

// StructTreeRoot represents the structure tree root dictionary.
// Reference: PDF documentation chapter 14.7 Logical Structure, table 322.
type StructTreeRoot struct {
	K                 []*KDict
	IDTree            *IDTree
	ParentTree        *_eb.PdfObjectDictionary
	ParentTreeNextKey int64
	RoleMap           _eb.PdfObject
	ClassMap          *_eb.PdfObjectDictionary
	_fggdf            *_eb.PdfIndirectObject
	_gagc             []_deb.UUID
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (_egga *PdfAction) GetContext() PdfModel {
	if _egga == nil {
		return nil
	}
	return _egga._aee
}
func (_eagd *pdfCIDFontType0) getFontDescriptor() *PdfFontDescriptor { return _eagd._bged }
func (_abda *pdfFontSimple) getFontEncoding() (_cabb string, _cdgee map[_fc.CharCode]_fc.GlyphName, _aafg error) {
	_cabb = "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"
	if _eecab, _fdagf := _eeece[_abda._agcc]; _fdagf {
		_cabb = _eecab
	} else if _abda.fontFlags()&_bdcce != 0 {
		for _gcba, _cgga := range _eeece {
			if _cc.Contains(_abda._agcc, _gcba) {
				_cabb = _cgga
				break
			}
		}
	}
	if _abda.Encoding == nil {
		return _cabb, nil, nil
	}
	switch _bfad := _abda.Encoding.(type) {
	case *_eb.PdfObjectName:
		return string(*_bfad), nil, nil
	case *_eb.PdfObjectDictionary:
		_fgdfdf, _bgdfa := _eb.GetName(_bfad.Get("\u0042\u0061\u0073e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
		if _bgdfa {
			_cabb = _fgdfdf.String()
		}
		if _ebece := _bfad.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"); _ebece != nil {
			_afccg, _efga := _eb.GetArray(_ebece)
			if !_efga {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042a\u0064\u0020\u0066on\u0074\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u003d\u0025\u002b\u0076\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073=\u0025\u0054", _bfad, _bfad.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"))
				return "", nil, _eb.ErrTypeError
			}
			_cdgee, _aafg = _fc.FromFontDifferences(_afccg)
		}
		return _cabb, _cdgee, _aafg
	default:
		_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0072\u0020\u0064\u0069\u0063t\u0020\u0028\u0025\u0054\u0029\u0020\u0025\u0073", _abda.Encoding, _abda.Encoding)
		return "", nil, _eb.ErrTypeError
	}
}

// ImageToRGB returns the passed in image. Method exists in order to satisfy
// the PdfColorspace interface.
func (_fcea *PdfColorspaceDeviceRGB) ImageToRGB(img Image) (Image, error) { return img, nil }
func (_babe *PdfReader) newPdfAnnotationSquigglyFromDict(_dde *_eb.PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	_edad := PdfAnnotationSquiggly{}
	_dege, _acag := _babe.newPdfAnnotationMarkupFromDict(_dde)
	if _acag != nil {
		return nil, _acag
	}
	_edad.PdfAnnotationMarkup = _dege
	_edad.QuadPoints = _dde.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_edad, nil
}

// NewEmbeddedFileFromObject construct a new EmbeddedFile from supplied object.
func NewEmbeddedFileFromObject(obj _eb.PdfObject) (*EmbeddedFile, error) {
	_gafe := _eb.TraceToDirectObject(obj)
	_egfd, _befa := _gafe.(*_eb.PdfObjectDictionary)
	if !_befa {
		return nil, _dcf.New("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_fcaf := _eb.TraceToDirectObject(_egfd.Get("\u0046"))
	if _fcaf == nil {
		return nil, _dcf.New("\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	}
	_dbaa, _befa := _fcaf.(*_eb.PdfObjectStream)
	if !_befa {
		return nil, _dcf.New("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_bdae := _dbaa.PdfObjectDictionary
	_eeacc := _bdae.Get("\u0050\u0061\u0072\u0061\u006d\u0073")
	if _eeacc == nil {
		return nil, _dcf.New("P\u0061\u0072\u0061\u006d\u0073\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0061\u0076\u0061i\u006ca\u0062\u006c\u0065")
	}
	_edbe, _befa := _eeacc.(*_eb.PdfObjectDictionary)
	if !_befa {
		return nil, _dcf.New("I\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0061\u006d\u0073 \u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_fdegd := ""
	_faaa := _edbe.Get("\u0043\u0068\u0065\u0063\u006b\u0053\u0075\u006d")
	if _faaa != nil {
		_fdegd = _faaa.(*_eb.PdfObjectString).Str()
	}
	_cdba, _edbge := _eb.DecodeStream(_dbaa)
	if _edbge != nil {
		return nil, _edbge
	}
	_dafd := &EmbeddedFile{Content: _cdba, Hash: _fdegd}
	return _dafd, nil
}

// ToPdfObject returns the PDF representation of the VRI dictionary.
func (_fdfge *VRI) ToPdfObject() *_eb.PdfObjectDictionary {
	_bfdc := _eb.MakeDict()
	_bfdc.SetIfNotNil(_eb.PdfObjectName("\u0043\u0065\u0072\u0074"), _cafc(_fdfge.Cert))
	_bfdc.SetIfNotNil(_eb.PdfObjectName("\u004f\u0043\u0053\u0050"), _cafc(_fdfge.OCSP))
	_bfdc.SetIfNotNil(_eb.PdfObjectName("\u0043\u0052\u004c"), _cafc(_fdfge.CRL))
	_bfdc.SetIfNotNil("\u0054\u0055", _fdfge.TU)
	_bfdc.SetIfNotNil("\u0054\u0053", _fdfge.TS)
	return _bfdc
}
func _ebbea() string {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	return _addbb
}

// GetDescent returns the Descent of the font `descriptor`.
func (_dafcb *PdfFontDescriptor) GetDescent() (float64, error) {
	return _eb.GetNumberAsFloat(_dafcb.Descent)
}

// NewStandard14FontWithEncoding returns the standard 14 font named `basefont` as a *PdfFont and
// a TextEncoder that encodes all the runes in `alphabet`, or an error if this is not possible.
// An error can occur if `basefont` is not one the standard 14 font names.
func NewStandard14FontWithEncoding(basefont StdFontName, alphabet map[rune]int) (*PdfFont, _fc.SimpleEncoder, error) {
	_badb, _degf := _adca(basefont)
	if _degf != nil {
		return nil, nil, _degf
	}
	_cded, _agcb := _badb.Encoder().(_fc.SimpleEncoder)
	if !_agcb {
		return nil, nil, _e.Errorf("\u006f\u006e\u006c\u0079\u0020s\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u002c\u0020\u0067\u006f\u0074\u0020\u0025\u0054", _badb.Encoder())
	}
	_aaca := make(map[rune]_fc.GlyphName)
	for _bgae := range alphabet {
		if _, _bebdb := _cded.RuneToCharcode(_bgae); !_bebdb {
			_, _aagef := _badb._gacdb.Read(_bgae)
			if !_aagef {
				_ddb.Log.Trace("r\u0075\u006e\u0065\u0020\u0025\u0023x\u003d\u0025\u0071\u0020\u006e\u006f\u0074\u0020\u0069n\u0020\u0074\u0068e\u0020f\u006f\u006e\u0074", _bgae, _bgae)
				continue
			}
			_gbca, _aagef := _fc.RuneToGlyph(_bgae)
			if !_aagef {
				_ddb.Log.Debug("\u006eo\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u0020\u0025\u0023\u0078\u003d\u0025\u0071", _bgae, _bgae)
				continue
			}
			if len(_aaca) >= 255 {
				return nil, nil, _dcf.New("\u0074\u006f\u006f\u0020\u006d\u0061\u006e\u0079\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020f\u006f\u0072\u0020\u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067")
			}
			_aaca[_bgae] = _gbca
		}
	}
	var (
		_fbdef []_fc.CharCode
		_cfecc []_fc.CharCode
	)
	for _cbcb := _fc.CharCode(1); _cbcb <= 0xff; _cbcb++ {
		_gbbga, _cccfe := _cded.CharcodeToRune(_cbcb)
		if !_cccfe {
			_fbdef = append(_fbdef, _cbcb)
			continue
		}
		if _, _cccfe = alphabet[_gbbga]; !_cccfe {
			_cfecc = append(_cfecc, _cbcb)
		}
	}
	_dage := append(_fbdef, _cfecc...)
	if len(_dage) < len(_aaca) {
		return nil, nil, _e.Errorf("n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u006ec\u006f\u0064\u0065\u0020\u0025\u0064\u0020ru\u006e\u0065\u0073\u002c \u0062\u0075\u0074\u0020\u0068\u0061\u0076\u0065\u0020on\u006c\u0079 \u0025\u0064\u0020\u0073\u006c\u006f\u0074\u0073", len(_aaca), len(_dage))
	}
	_fddbc := make([]rune, 0, len(_aaca))
	for _acaee := range _aaca {
		_fddbc = append(_fddbc, _acaee)
	}
	_ba.Slice(_fddbc, func(_cfba, _dgeg int) bool { return _fddbc[_cfba] < _fddbc[_dgeg] })
	_afgc := make(map[_fc.CharCode]_fc.GlyphName, len(_fddbc))
	for _, _gfdaf := range _fddbc {
		_aegc := _dage[0]
		_dage = _dage[1:]
		_afgc[_aegc] = _aaca[_gfdaf]
	}
	_cded = _fc.ApplyDifferences(_cded, _afgc)
	_badb.SetEncoder(_cded)
	return &PdfFont{_fdaa: &_badb}, _cded, nil
}

// IsCID returns true if the underlying font is CID.
func (_bfced *PdfFont) IsCID() bool { return _bfced.baseFields().isCIDFont() }

// PdfFilespec represents a file specification which can either refer to an external or embedded file.
type PdfFilespec struct {
	Type           _eb.PdfObject
	FS             _eb.PdfObject
	F              _eb.PdfObject
	UF             _eb.PdfObject
	DOS            _eb.PdfObject
	Mac            _eb.PdfObject
	Unix           _eb.PdfObject
	ID             _eb.PdfObject
	V              _eb.PdfObject
	EF             _eb.PdfObject
	RF             _eb.PdfObject
	Desc           _eb.PdfObject
	CI             _eb.PdfObject
	AFRelationship _eb.PdfObject
	_cbefe         _eb.PdfObject
}

// NewPdfAnnotationSquiggly returns a new text squiggly annotation.
func NewPdfAnnotationSquiggly() *PdfAnnotationSquiggly {
	_cegc := NewPdfAnnotation()
	_caf := &PdfAnnotationSquiggly{}
	_caf.PdfAnnotation = _cegc
	_caf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cegc.SetContext(_caf)
	return _caf
}

// PdfShadingPattern is a Shading patterns that provide a smooth transition between colors across an area to be painted,
// i.e. color(x,y) = f(x,y) at each point.
// It is a type 2 pattern (PatternType = 2).
type PdfShadingPattern struct {
	*PdfPattern
	Shading   *PdfShading
	Matrix    *_eb.PdfObjectArray
	ExtGState _eb.PdfObject
}

// PdfColorPattern represents a pattern color.
type PdfColorPattern struct {
	Color       PdfColor
	PatternName _eb.PdfObjectName
}

func (_eed *PdfReader) newPdfActionSetOCGStateFromDict(_aacc *_eb.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	return &PdfActionSetOCGState{State: _aacc.Get("\u0053\u0074\u0061t\u0065"), PreserveRB: _aacc.Get("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042")}, nil
}

// NewPdfAnnotationMovie returns a new movie annotation.
func NewPdfAnnotationMovie() *PdfAnnotationMovie {
	_bac := NewPdfAnnotation()
	_gge := &PdfAnnotationMovie{}
	_gge.PdfAnnotation = _bac
	_bac.SetContext(_gge)
	return _gge
}

// K returns the value of the key component of the color.
func (_bccf *PdfColorDeviceCMYK) K() float64 { return _bccf[3] }

// ColorFromFloats returns a new PdfColor based on input color components.
func (_fffd *PdfColorspaceDeviceN) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != _fffd.GetNumComponents() {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_feeea, _caga := _fffd.TintTransform.Evaluate(vals)
	if _caga != nil {
		return nil, _caga
	}
	_efaa, _caga := _fffd.AlternateSpace.ColorFromFloats(_feeea)
	if _caga != nil {
		return nil, _caga
	}
	return _efaa, nil
}

// DecodeArray returns the range of color component values in the ICCBased colorspace.
func (_cabga *PdfColorspaceICCBased) DecodeArray() []float64 { return _cabga.Range }

// GetDSS gets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_ggcg *PdfAppender) GetDSS() (_gebc *DSS) { return _ggcg._cedc }

// NewPdfActionTrans returns a new "trans" action.
func NewPdfActionTrans() *PdfActionTrans {
	_abc := NewPdfAction()
	_dcb := &PdfActionTrans{}
	_dcb.PdfAction = _abc
	_abc.SetContext(_dcb)
	return _dcb
}

// ImageToRGB convert an indexed image to RGB.
func (_cacde *PdfColorspaceSpecialIndexed) ImageToRGB(img Image) (Image, error) {
	N := _cacde.Base.GetNumComponents()
	if N < 1 {
		return Image{}, _e.Errorf("\u0062\u0061d \u0062\u0061\u0073e\u0020\u0063\u006f\u006cors\u0070ac\u0065\u0020\u004e\u0075\u006d\u0043\u006fmp\u006f\u006e\u0065\u006e\u0074\u0073\u003d%\u0064", N)
	}
	_gegb := _df.NewImageBase(int(img.Width), int(img.Height), 8, N, nil, img._bdcab, img._fedc)
	_ecgc := _bb.NewReader(img.getBase())
	_efda := _bb.NewWriter(_gegb)
	var (
		_efgf uint32
		_gadb int
		_adfd error
	)
	for {
		_efgf, _adfd = _ecgc.ReadSample()
		if _adfd == _bagf.EOF {
			break
		} else if _adfd != nil {
			return img, _adfd
		}
		_gadb = int(_efgf)
		_ddb.Log.Trace("\u0049\u006ed\u0065\u0078\u0065\u0064\u003a\u0020\u0069\u006e\u0064\u0065\u0078\u003d\u0025\u0064\u0020\u004e\u003d\u0025\u0064\u0020\u006c\u0075t=\u0025\u0064", _gadb, N, len(_cacde._efcb))
		if (_gadb+1)*N > len(_cacde._efcb) {
			_gadb = len(_cacde._efcb)/N - 1
			_ddb.Log.Trace("C\u006c\u0069\u0070\u0070in\u0067 \u0074\u006f\u0020\u0069\u006ed\u0065\u0078\u003a\u0020\u0025\u0064", _gadb)
			if _gadb < 0 {
				_ddb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069p\u0020\u0069\u006e\u0064\u0065\u0078.\u0020\u0049\u0073\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006ce\u0020\u0064\u0061\u006d\u0061\u0067\u0065\u0064\u003f")
				break
			}
		}
		for _dgedf := _gadb * N; _dgedf < (_gadb+1)*N; _dgedf++ {
			if _adfd = _efda.WriteSample(uint32(_cacde._efcb[_dgedf])); _adfd != nil {
				return img, _adfd
			}
		}
	}
	return _cacde.Base.ImageToRGB(_ggaa(&_gegb))
}

// GetNumPages returns the number of pages in the document.
func (_cfgbeb *PdfReader) GetNumPages() (int, error) {
	if _cfgbeb._ebbe.GetCrypter() != nil && !_cfgbeb._ebbe.IsAuthenticated() {
		return 0, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	return len(_cfgbeb._cbaff), nil
}

var _ pdfFont = (*pdfFontType0)(nil)

// SetFontByName sets the font specified by keyName to the given object.
func (_bbbbcc *PdfPageResources) SetFontByName(keyName _eb.PdfObjectName, obj _eb.PdfObject) error {
	if _bbbbcc.Font == nil {
		_bbbbcc.Font = _eb.MakeDict()
	}
	_gdfbb, _dfgeb := _eb.TraceToDirectObject(_bbbbcc.Font).(*_eb.PdfObjectDictionary)
	if !_dfgeb {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _eb.TraceToDirectObject(_bbbbcc.Font))
		return _eb.ErrTypeError
	}
	_gdfbb.Set(keyName, obj)
	return nil
}

// ToPdfObject returns the button field dictionary within an indirect object.
func (_egce *PdfFieldButton) ToPdfObject() _eb.PdfObject {
	_egce.PdfField.ToPdfObject()
	_ebfcc := _egce._adgda
	_cafd := _ebfcc.PdfObject.(*_eb.PdfObjectDictionary)
	_cafd.Set("\u0046\u0054", _eb.MakeName("\u0042\u0074\u006e"))
	if _egce.Opt != nil {
		_cafd.Set("\u004f\u0070\u0074", _egce.Opt)
	}
	return _ebfcc
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_cdaae *PdfColorspaceSpecialSeparation) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_dfdf, _bdebb := _eb.GetNumbersAsFloat(objects)
	if _bdebb != nil {
		return nil, _bdebb
	}
	return _cdaae.ColorFromFloats(_dfdf)
}
func (_dfdff *LTV) validateSig(_gbbd *PdfSignature) error {
	if _gbbd == nil || _gbbd.Contents == nil || len(_gbbd.Contents.Bytes()) == 0 {
		return _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0066\u0069\u0065l\u0064:\u0020\u0025\u0076", _gbbd)
	}
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_bagc *PdfActionGoToE) ToPdfObject() _eb.PdfObject {
	_bagc.PdfAction.ToPdfObject()
	_ebc := _bagc._dee
	_afd := _ebc.PdfObject.(*_eb.PdfObjectDictionary)
	_afd.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeGoToE)))
	if _bagc.F != nil {
		_afd.Set("\u0046", _bagc.F.ToPdfObject())
	}
	_afd.SetIfNotNil("\u0044", _bagc.D)
	_afd.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _bagc.NewWindow)
	_afd.SetIfNotNil("\u0054", _bagc.T)
	return _ebc
}

// PageProcessCallback callback function used in page loading
// that could be used to modify the page content.
//
// If an error is returned, the `ToWriter` process would fail.
//
// This callback, if defined, will take precedence over `PageCallback` callback.
type PageProcessCallback func(_ddbe int, _bcaaf *PdfPage) error

func (_acfa *PdfReader) traverseObjectData(_cbcbf _eb.PdfObject) error {
	return _eb.ResolveReferencesDeep(_cbcbf, _acfa._bcefc)
}

// ToPdfObject implements interface PdfModel.
// Note: Call the sub-annotation's ToPdfObject to set both the generic and non-generic information.
func (_ceea *PdfAnnotation) ToPdfObject() _eb.PdfObject {
	_gcb := _ceea._ggf
	_bbbb := _gcb.PdfObject.(*_eb.PdfObjectDictionary)
	_bbbb.Clear()
	_bbbb.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0041\u006e\u006eo\u0074"))
	_bbbb.SetIfNotNil("\u0052\u0065\u0063\u0074", _ceea.Rect)
	_bbbb.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _ceea.Contents)
	_bbbb.SetIfNotNil("\u0050", _ceea.P)
	_bbbb.SetIfNotNil("\u004e\u004d", _ceea.NM)
	_bbbb.SetIfNotNil("\u004d", _ceea.M)
	_bbbb.SetIfNotNil("\u0046", _ceea.F)
	_bbbb.SetIfNotNil("\u0041\u0050", _ceea.AP)
	_bbbb.SetIfNotNil("\u0041\u0053", _ceea.AS)
	_bbbb.SetIfNotNil("\u0042\u006f\u0072\u0064\u0065\u0072", _ceea.Border)
	_bbbb.SetIfNotNil("\u0043", _ceea.C)
	_bbbb.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _ceea.StructParent)
	_bbbb.SetIfNotNil("\u004f\u0043", _ceea.OC)
	return _gcb
}

// EnableByName LTV enables the signature dictionary of the PDF AcroForm
// field identified the specified name. The signing certificate chain is
// extracted from the signature dictionary. Optionally, additional certificates
// can be specified through the `extraCerts` parameter. The LTV client attempts
// to build the certificate chain up to a trusted root by downloading any
// missing certificates.
func (_feaad *LTV) EnableByName(name string, extraCerts []*_bag.Certificate) error {
	_fbge := _feaad._feba._adac.AcroForm
	for _, _egcg := range _fbge.AllFields() {
		_eaceb, _ := _egcg.GetContext().(*PdfFieldSignature)
		if _eaceb == nil {
			continue
		}
		if _ddfeg := _eaceb.PartialName(); _ddfeg != name {
			continue
		}
		return _feaad.Enable(_eaceb.V, extraCerts)
	}
	return nil
}

// SetReason sets the `Reason` field of the signature.
func (_cedcb *PdfSignature) SetReason(reason string) {
	_cedcb.Reason = _eb.MakeEncodedString(reason, true)
}

// GetPrimitiveFromModel returns the primitive object corresponding to the input `model`.
func (_baeg *modelManager) GetPrimitiveFromModel(model PdfModel) _eb.PdfObject {
	_fbbbc, _cead := _baeg._edcg[model]
	if !_cead {
		return nil
	}
	return _fbbbc
}
func (_dfaef *PdfWriter) writeBytes(_eeeec []byte) {
	if _dfaef._ceffa != nil {
		return
	}
	_ffgcd, _fgabe := _dfaef._ccfbc.Write(_eeeec)
	_dfaef._dfabe += int64(_ffgcd)
	_dfaef._ceffa = _fgabe
}

// NewPdfColorspaceFromPdfObject loads a PdfColorspace from a PdfObject.  Returns an error if there is
// a failure in loading.
func NewPdfColorspaceFromPdfObject(obj _eb.PdfObject) (PdfColorspace, error) {
	if obj == nil {
		return nil, nil
	}
	var _acdb *_eb.PdfIndirectObject
	var _gegfc *_eb.PdfObjectName
	var _fddde *_eb.PdfObjectArray
	if _dadea, _bgbf := obj.(*_eb.PdfIndirectObject); _bgbf {
		_acdb = _dadea
	}
	obj = _eb.TraceToDirectObject(obj)
	switch _ecad := obj.(type) {
	case *_eb.PdfObjectArray:
		_fddde = _ecad
	case *_eb.PdfObjectName:
		_gegfc = _ecad
	}
	if _gegfc != nil {
		switch *_gegfc {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			return NewPdfColorspaceDeviceGray(), nil
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			return NewPdfColorspaceDeviceRGB(), nil
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return NewPdfColorspaceDeviceCMYK(), nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return NewPdfColorspaceSpecialPattern(), nil
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0025\u0073", *_gegfc)
			return nil, _eggcg
		}
	}
	if _fddde != nil && _fddde.Len() > 0 {
		var _beae _eb.PdfObject = _acdb
		if _acdb == nil {
			_beae = _fddde
		}
		if _cgdd, _efcd := _eb.GetName(_fddde.Get(0)); _efcd {
			switch _cgdd.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				if _fddde.Len() == 1 {
					return NewPdfColorspaceDeviceGray(), nil
				}
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				if _fddde.Len() == 1 {
					return NewPdfColorspaceDeviceRGB(), nil
				}
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _fddde.Len() == 1 {
					return NewPdfColorspaceDeviceCMYK(), nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079":
				return _cfgc(_beae)
			case "\u0043\u0061\u006c\u0052\u0047\u0042":
				return _ffaa(_beae)
			case "\u004c\u0061\u0062":
				return _ecdgf(_beae)
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064":
				return _ecff(_beae)
			case "\u0050a\u0074\u0074\u0065\u0072\u006e":
				return _eegg(_beae)
			case "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return _dgfc(_beae)
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e":
				return _agcd(_beae)
			case "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return _gbge(_beae)
			default:
				_ddb.Log.Debug("A\u0072\u0072\u0061\u0079\u0020\u0077i\u0074\u0068\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0061m\u0065:\u0020\u0025\u0073", *_cgdd)
			}
		}
	}
	_ddb.Log.Debug("\u0050\u0044\u0046\u0020\u0046i\u006c\u0065\u0020\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", obj.String())
	return nil, ErrTypeCheck
}

// NewPdfOutputIntentFromPdfObject creates a new PdfOutputIntent from the input core.PdfObject.
func NewPdfOutputIntentFromPdfObject(object _eb.PdfObject) (*PdfOutputIntent, error) {
	_fdga := &PdfOutputIntent{}
	if _fdcfg := _fdga.ParsePdfObject(object); _fdcfg != nil {
		return nil, _fdcfg
	}
	return _fdga, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_ggdb *PdfColorspaceDeviceCMYK) ToPdfObject() _eb.PdfObject {
	return _eb.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b")
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for an RGB device.
func (_cecg *PdfColorspaceDeviceRGB) GetNumComponents() int { return 3 }

// GetPageDict converts the Page to a PDF object dictionary.
func (_added *PdfPage) GetPageDict() *_eb.PdfObjectDictionary {
	_feadg := _added._aaagaa
	_feadg.Clear()
	_feadg.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0050\u0061\u0067\u0065"))
	_feadg.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _added.Parent)
	if _added.LastModified != nil {
		_feadg.Set("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _added.LastModified.ToPdfObject())
	}
	if _added.Resources != nil {
		_feadg.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _added.Resources.ToPdfObject())
	}
	if _added.CropBox != nil {
		_feadg.Set("\u0043r\u006f\u0070\u0042\u006f\u0078", _added.CropBox.ToPdfObject())
	}
	if _added.MediaBox != nil {
		_feadg.Set("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", _added.MediaBox.ToPdfObject())
	}
	if _added.BleedBox != nil {
		_feadg.Set("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078", _added.BleedBox.ToPdfObject())
	}
	if _added.TrimBox != nil {
		_feadg.Set("\u0054r\u0069\u006d\u0042\u006f\u0078", _added.TrimBox.ToPdfObject())
	}
	if _added.ArtBox != nil {
		_feadg.Set("\u0041\u0072\u0074\u0042\u006f\u0078", _added.ArtBox.ToPdfObject())
	}
	_feadg.SetIfNotNil("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f", _added.BoxColorInfo)
	_feadg.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _added.Contents)
	if _added.Rotate != nil {
		_feadg.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _eb.MakeInteger(*_added.Rotate))
	}
	_feadg.SetIfNotNil("\u0047\u0072\u006fu\u0070", _added.Group)
	_feadg.SetIfNotNil("\u0054\u0068\u0075m\u0062", _added.Thumb)
	_feadg.SetIfNotNil("\u0042", _added.B)
	_feadg.SetIfNotNil("\u0044\u0075\u0072", _added.Dur)
	_feadg.SetIfNotNil("\u0054\u0072\u0061n\u0073", _added.Trans)
	_feadg.SetIfNotNil("\u0041\u0041", _added.AA)
	_feadg.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _added.Metadata)
	_feadg.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _added.PieceInfo)
	_feadg.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _added.StructParents)
	_feadg.SetIfNotNil("\u0049\u0044", _added.ID)
	_feadg.SetIfNotNil("\u0050\u005a", _added.PZ)
	_feadg.SetIfNotNil("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f", _added.SeparationInfo)
	_feadg.SetIfNotNil("\u0054\u0061\u0062\u0073", _added.Tabs)
	_feadg.SetIfNotNil("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064", _added.TemplateInstantiated)
	_feadg.SetIfNotNil("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s", _added.PresSteps)
	_feadg.SetIfNotNil("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074", _added.UserUnit)
	_feadg.SetIfNotNil("\u0056\u0050", _added.VP)
	if _added._dbga != nil {
		_befg := _eb.MakeArray()
		for _, _ddbdg := range _added._dbga {
			if _bafd := _ddbdg.GetContext(); _bafd != nil {
				_befg.Append(_bafd.ToPdfObject())
			} else {
				_befg.Append(_ddbdg.ToPdfObject())
			}
		}
		if _befg.Len() > 0 {
			_feadg.Set("\u0041\u006e\u006e\u006f\u0074\u0073", _befg)
		}
	} else if _added.Annots != nil {
		_feadg.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0073", _added.Annots)
	}
	return _feadg
}

// IsEncrypted returns true if the PDF file is encrypted.
func (_dgabe *PdfReader) IsEncrypted() (bool, error) { return _dgabe._ebbe.IsEncrypted() }
func _ccbbc(_fgdd map[_fg.GID]int, _cebc uint16) *_eb.PdfObjectArray {
	_dgdb := &_eb.PdfObjectArray{}
	_gdbcbg := _fg.GID(_cebc)
	for _dgdbg := _fg.GID(0); _dgdbg < _gdbcbg; {
		_dcgbd, _dfbga := _fgdd[_dgdbg]
		if !_dfbga {
			_dgdbg++
			continue
		}
		_fcded := _dgdbg
		for _ddedg := _fcded + 1; _ddedg < _gdbcbg; _ddedg++ {
			if _fbfbg, _fggga := _fgdd[_ddedg]; !_fggga || _dcgbd != _fbfbg {
				break
			}
			_fcded = _ddedg
		}
		_dgdb.Append(_eb.MakeInteger(int64(_dgdbg)))
		_dgdb.Append(_eb.MakeInteger(int64(_fcded)))
		_dgdb.Append(_eb.MakeInteger(int64(_dcgbd)))
		_dgdbg = _fcded + 1
	}
	return _dgdb
}

// FillWithAppearance populates `form` with values provided by `provider`.
// If not nil, `appGen` is used to generate appearance dictionaries for the
// field annotations, based on the specified settings. Otherwise, appearance
// generation is skipped.
// e.g.: appGen := annotator.FieldAppearance{OnlyIfMissing: true, RegenerateTextFields: true}
// NOTE: In next major version this functionality will be part of Fill. (v4)
func (_ffacfa *PdfAcroForm) FillWithAppearance(provider FieldValueProvider, appGen FieldAppearanceGenerator) error {
	_fdcfab := _ffacfa.fill(provider, appGen)
	if _fdcfab != nil {
		return _fdcfab
	}
	if _, _gdeaf := provider.(FieldImageProvider); _gdeaf {
		_fdcfab = _ffacfa.fillImageWithAppearance(provider.(FieldImageProvider), appGen)
	}
	return _fdcfab
}
func _geebd(_ddbdd _eb.PdfObject, _eeace *fontCommon) (*_ff.CMap, error) {
	_abca, _cbbba := _eb.GetStream(_ddbdd)
	if !_cbbba {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0074\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0054\u006f\u0043m\u0061\u0070\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0054\u0029", _ddbdd)
		return nil, _eb.ErrTypeError
	}
	_cfcbd, _affdg := _eb.DecodeStream(_abca)
	if _affdg != nil {
		return nil, _affdg
	}
	_afeba, _affdg := _ff.LoadCmapFromData(_cfcbd, !_eeace.isCIDFont())
	if _affdg != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u004e\u0075\u006d\u0062\u0065\u0072\u003d\u0025\u0064\u0020\u0065\u0072r=\u0025\u0076", _abca.ObjectNumber, _affdg)
	}
	return _afeba, _affdg
}
func (_aeeaf *PdfWriter) writeObjects() {
	_ddb.Log.Trace("\u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0025d\u0020\u006f\u0062\u006a", len(_aeeaf._dcfgf))
	_aeeaf._aggfdb = make(map[int]crossReference)
	_aeeaf._aggfdb[0] = crossReference{Type: 0, ObjectNumber: 0, Generation: 0xFFFF}
	if _aeeaf._bfdeb.ObjectMap != nil {
		for _bdcbb, _dfdfca := range _aeeaf._bfdeb.ObjectMap {
			if _bdcbb == 0 {
				continue
			}
			if _dfdfca.XType == _eb.XrefTypeObjectStream {
				_debgd := crossReference{Type: 2, ObjectNumber: _dfdfca.OsObjNumber, Index: _dfdfca.OsObjIndex}
				_aeeaf._aggfdb[_bdcbb] = _debgd
			}
			if _dfdfca.XType == _eb.XrefTypeTableEntry {
				_eebde := crossReference{Type: 1, ObjectNumber: _dfdfca.ObjectNumber, Offset: _dfdfca.Offset}
				_aeeaf._aggfdb[_bdcbb] = _eebde
			}
		}
	}
}

// GetContainingPdfObject gets the primitive used to parse the color space.
func (_bedd *PdfColorspaceICCBased) GetContainingPdfObject() _eb.PdfObject { return _bedd._cdebc }

// EncryptOptions represents encryption options for an output PDF.
type EncryptOptions struct {
	Permissions _cg.Permissions
	Algorithm   EncryptionAlgorithm
}

// PdfColorspaceSpecialIndexed is an indexed color space is a lookup table, where the input element
// is an index to the lookup table and the output is a color defined in the lookup table in the Base
// colorspace.
// [/Indexed base hival lookup]
type PdfColorspaceSpecialIndexed struct {
	Base   PdfColorspace
	HiVal  int
	Lookup _eb.PdfObject
	_efcb  []byte
	_daac  *_eb.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
func (_bbab *PdfAnnotationSquiggly) ToPdfObject() _eb.PdfObject {
	_bbab.PdfAnnotation.ToPdfObject()
	_ada := _bbab._ggf
	_acab := _ada.PdfObject.(*_eb.PdfObjectDictionary)
	_bbab.PdfAnnotationMarkup.appendToPdfDictionary(_acab)
	_acab.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079"))
	_acab.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _bbab.QuadPoints)
	return _ada
}

// SignatureHandlerDocMDP extends SignatureHandler with the ValidateWithOpts method for checking the DocMDP policy.
type SignatureHandlerDocMDP interface {
	SignatureHandler

	// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser
	// ValidateWithOpts shall contain Validate call
	ValidateWithOpts(_dedcc *PdfSignature, _agaca Hasher, _acgg SignatureHandlerDocMDPParams) (SignatureValidationResult, error)
}

// GetContainingPdfObject implements interface PdfModel.
func (_cgfd *PdfFilespec) GetContainingPdfObject() _eb.PdfObject { return _cgfd._cbefe }

// SetColorSpace sets `r` colorspace object to `colorspace`.
func (_fbdbc *PdfPageResources) SetColorSpace(colorspace *PdfPageResourcesColorspaces) {
	_fbdbc._cfcff = colorspace
}
func (_bfag *PdfReader) newPdfOutlineItemFromIndirectObject(_dgcgg *_eb.PdfIndirectObject) (*PdfOutlineItem, error) {
	_defda, _bfgaa := _dgcgg.PdfObject.(*_eb.PdfObjectDictionary)
	if !_bfgaa {
		return nil, _e.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_agdffa := NewPdfOutlineItem()
	_ccdcb := _defda.Get("\u0054\u0069\u0074l\u0065")
	if _ccdcb == nil {
		return nil, _e.Errorf("\u006d\u0069\u0073s\u0069\u006e\u0067\u0020\u0054\u0069\u0074\u006c\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0049\u0074\u0065\u006d\u0020\u0028r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	_abdc, _gaee := _eb.GetString(_ccdcb)
	if !_gaee {
		return nil, _e.Errorf("\u0074\u0069\u0074le\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0054\u0029", _ccdcb)
	}
	_agdffa.Title = _abdc
	if _acacg := _defda.Get("\u0043\u006f\u0075n\u0074"); _acacg != nil {
		_gbbc, _facg := _acacg.(*_eb.PdfObjectInteger)
		if !_facg {
			return nil, _e.Errorf("\u0063o\u0075\u006e\u0074\u0020n\u006f\u0074\u0020\u0061\u006e \u0069n\u0074e\u0067\u0065\u0072\u0020\u0028\u0025\u0054)", _acacg)
		}
		_gefa := int64(*_gbbc)
		_agdffa.Count = &_gefa
	}
	if _dffa := _defda.Get("\u0044\u0065\u0073\u0074"); _dffa != nil {
		_agdffa.Dest = _eb.ResolveReference(_dffa)
		if !_bfag._cfcgdf {
			_cagg := _bfag.traverseObjectData(_agdffa.Dest)
			if _cagg != nil {
				return nil, _cagg
			}
		}
	}
	if _dfdfce := _defda.Get("\u0041"); _dfdfce != nil {
		_agdffa.A = _eb.ResolveReference(_dfdfce)
		if !_bfag._cfcgdf {
			_eeadc := _bfag.traverseObjectData(_agdffa.A)
			if _eeadc != nil {
				return nil, _eeadc
			}
		}
	}
	if _gfdg := _defda.Get("\u0053\u0045"); _gfdg != nil {
		_agdffa.SE = nil
	}
	if _gfabf := _defda.Get("\u0043"); _gfabf != nil {
		_agdffa.C = _eb.ResolveReference(_gfabf)
	}
	if _bfgd := _defda.Get("\u0046"); _bfgd != nil {
		_agdffa.F = _eb.ResolveReference(_bfgd)
	}
	return _agdffa, nil
}
func _bcbe(_bgaeg _eb.PdfObject, _fedfc *PdfReader) (*OutlineDest, error) {
	_fbgee, _cbbfba := _eb.GetArray(_bgaeg)
	if !_cbbfba {
		return nil, _dcf.New("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0064\u0065\u0073t\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_degbf := _fbgee.Len()
	if _degbf < 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u006fu\u0074\u006c\u0069n\u0065\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061ti\u006f\u006e\u0020a\u0072\u0072a\u0079\u0020\u006c\u0065\u006e\u0067t\u0068\u003a \u0025\u0064", _degbf)
	}
	_efbab := &OutlineDest{Mode: "\u0046\u0069\u0074"}
	_baggc := _fbgee.Get(0)
	if _cebdg, _gccbe := _eb.GetIndirect(_baggc); _gccbe {
		if _, _fgfdd, _daafa := _fedfc.PageFromIndirectObject(_cebdg); _daafa == nil {
			_efbab.Page = int64(_fgfdd - 1)
		} else {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020g\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002b\u0076", _cebdg)
		}
		_efbab.PageObj = _cebdg
	} else if _gdaab, _dbaaa := _eb.GetIntVal(_baggc); _dbaaa {
		if _gdaab >= 0 && _gdaab < len(_fedfc.PageList) {
			_efbab.PageObj = _fedfc.PageList[_gdaab].GetPageAsIndirectObject()
		} else {
			_ddb.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _gdaab)
		}
		_efbab.Page = int64(_gdaab)
	} else {
		return nil, _e.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006f\u0075\u0074\u006cine\u0020de\u0073\u0074\u0069\u006e\u0061\u0074\u0069on\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0054", _baggc)
	}
	_ceccg, _cbbfba := _eb.GetNameVal(_fbgee.Get(1))
	if !_cbbfba {
		_ddb.Log.Debug("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006d\u0061\u0067\u006e\u0069\u0066\u0069\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _fbgee.Get(1))
		return _efbab, nil
	}
	switch _ceccg {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		if _degbf > 2 {
			_efbab.Y, _ = _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_fbgee.Get(2)))
		}
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		if _degbf > 2 {
			_efbab.X, _ = _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_fbgee.Get(2)))
		}
	case "\u0058\u0059\u005a":
		if _degbf > 4 {
			_efbab.X, _ = _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_fbgee.Get(2)))
			_efbab.Y, _ = _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_fbgee.Get(3)))
			_efbab.Zoom, _ = _eb.GetNumberAsFloat(_eb.TraceToDirectObject(_fbgee.Get(4)))
		}
	default:
		_ceccg = "\u0046\u0069\u0074"
	}
	_efbab.Mode = _ceccg
	return _efbab, nil
}

// NewPdfDate returns a new PdfDate object from a PDF date string (see 7.9.4 Dates).
// format: "D: YYYYMMDDHHmmSSOHH'mm"
func NewPdfDate(dateStr string) (PdfDate, error) {
	_gbdae, _efff := _eggb.ParsePdfTime(dateStr)
	if _efff != nil {
		return PdfDate{}, _efff
	}
	return NewPdfDateFromTime(_gbdae)
}

const (
	_ PdfOutputIntentType = iota
	PdfOutputIntentTypeA1
	PdfOutputIntentTypeA2
	PdfOutputIntentTypeA3
	PdfOutputIntentTypeA4
	PdfOutputIntentTypeX
)

// IsFitWindow returns the value of the fitWindow flag.
func (_feeec *ViewerPreferences) IsFitWindow() bool {
	if _feeec._edbaa == nil {
		return false
	}
	return *_feeec._edbaa
}

type pdfSignDictionary struct {
	*_eb.PdfObjectDictionary
	_bfadg  *SignatureHandler
	_abegc  *PdfSignature
	_afeabg int64
	_eegea  int
	_defaa  int
	_caeba  int
	_dbcd   int
}

// ColorFromPdfObjects gets the color from a series of pdf objects (4 for cmyk).
func (_ddfg *PdfColorspaceDeviceCMYK) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 4 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ebee, _bbcf := _eb.GetNumbersAsFloat(objects)
	if _bbcf != nil {
		return nil, _bbcf
	}
	return _ddfg.ColorFromFloats(_ebee)
}
func (_cgab *PdfShading) getShadingDict() (*_eb.PdfObjectDictionary, error) {
	_fcaef := _cgab._cefaa
	if _aegaf, _cbdcd := _fcaef.(*_eb.PdfIndirectObject); _cbdcd {
		_cdaaf, _ggede := _aegaf.PdfObject.(*_eb.PdfObjectDictionary)
		if !_ggede {
			return nil, _eb.ErrTypeError
		}
		return _cdaaf, nil
	} else if _gdgcf, _ffagd := _fcaef.(*_eb.PdfObjectStream); _ffagd {
		return _gdgcf.PdfObjectDictionary, nil
	} else if _dgcf, _ccage := _fcaef.(*_eb.PdfObjectDictionary); _ccage {
		return _dgcf, nil
	} else {
		_ddb.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0063\u0063\u0065s\u0073\u0020\u0073\u0068\u0061\u0064\u0069n\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079")
		return nil, _eb.ErrTypeError
	}
}
func (_cefed *pdfFontType0) baseFields() *fontCommon { return &_cefed.fontCommon }

// ImageToRGB convert 1-component grayscale data to 3-component RGB.
func (_ffab *PdfColorspaceDeviceGray) ImageToRGB(img Image) (Image, error) {
	if img.ColorComponents != 1 {
		return img, _dcf.New("\u0074\u0068e \u0070\u0072\u006fv\u0069\u0064\u0065\u0064 im\u0061ge\u0020\u0069\u0073\u0020\u006e\u006f\u0074 g\u0072\u0061\u0079\u0020\u0073\u0063\u0061l\u0065")
	}
	_faed, _gcbb := _df.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._bdcab, img._fedc)
	if _gcbb != nil {
		return img, _gcbb
	}
	_eaca, _gcbb := _df.NRGBAConverter.Convert(_faed)
	if _gcbb != nil {
		return img, _gcbb
	}
	_cddb := _ggaa(_eaca.Base())
	_ddb.Log.Trace("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079\u0020\u002d>\u0020\u0052\u0047\u0042")
	_ddb.Log.Trace("s\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", img.Data)
	_ddb.Log.Trace("\u0052G\u0042 \u0073\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", _cddb.Data)
	_ddb.Log.Trace("\u0025\u0076\u0020\u002d\u003e\u0020\u0025\u0076", img, _cddb)
	return _cddb, nil
}

// NewPdfSignature creates a new PdfSignature object.
func NewPdfSignature(handler SignatureHandler) *PdfSignature {
	_gefad := &PdfSignature{Type: _eb.MakeName("\u0053\u0069\u0067"), Handler: handler}
	_fgddee := &pdfSignDictionary{PdfObjectDictionary: _eb.MakeDict(), _bfadg: &handler, _abegc: _gefad}
	_gefad._cddce = _eb.MakeIndirectObject(_fgddee)
	return _gefad
}
func (_cfgca fontCommon) isCIDFont() bool {
	if _cfgca._fgdee == "" {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0069\u0073\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u002e\u0020\u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _cfgca)
	}
	_eeaga := false
	switch _cfgca._fgdee {
	case "\u0054\u0079\u0070e\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_eeaga = true
	}
	_ddb.Log.Trace("i\u0073\u0043\u0049\u0044\u0046\u006fn\u0074\u003a\u0020\u0069\u0073\u0043\u0049\u0044\u003d%\u0074\u0020\u0066o\u006et\u003d\u0025\u0073", _eeaga, _cfgca)
	return _eeaga
}

// PdfColorDeviceRGB represents a color in DeviceRGB colorspace with R, G, B components, where component is
// defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceRGB [3]float64

// GetMCID returns the MCID of the KValue.
func (_afbgg *KValue) GetMCID() *int { return _afbgg._ddaf }

// NewPdfActionMovie returns a new "movie" action.
func NewPdfActionMovie() *PdfActionMovie {
	_cfg := NewPdfAction()
	_da := &PdfActionMovie{}
	_da.PdfAction = _cfg
	_cfg.SetContext(_da)
	return _da
}

// SetStructParentsKey sets the StructParents key.
func (_cadef *PdfPage) SetStructParentsKey(key int) {
	if key == -1 {
		_cadef.StructParents = nil
	} else {
		_cadef.StructParents = _eb.MakeInteger(int64(key))
	}
}

// GetEncryptionMethod returns a descriptive information string about the encryption method used.
func (_dcfee *PdfReader) GetEncryptionMethod() string {
	_ebabb := _dcfee._ebbe.GetCrypter()
	return _ebabb.String()
}

// ToPdfObject converts date to a PDF string object.
func (_gcac *PdfDate) ToPdfObject() _eb.PdfObject {
	_cdeae := _e.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _gcac._fffdf, _gcac._aacegf, _gcac._afgfb, _gcac._bdfbf, _gcac._aecad, _gcac._cfdb, _gcac._eaede, _gcac._fgggd, _gcac._gadga)
	return _eb.MakeString(_cdeae)
}

// NewPdfAnnotationPrinterMark returns a new printermark annotation.
func NewPdfAnnotationPrinterMark() *PdfAnnotationPrinterMark {
	_cee := NewPdfAnnotation()
	_aba := &PdfAnnotationPrinterMark{}
	_aba.PdfAnnotation = _cee
	_cee.SetContext(_aba)
	return _aba
}

// GetTrailer returns the PDF's trailer dictionary.
func (_dfbcc *PdfReader) GetTrailer() (*_eb.PdfObjectDictionary, error) {
	_deecaa := _dfbcc._ebbe.GetTrailer()
	if _deecaa == nil {
		return nil, _dcf.New("\u0074r\u0061i\u006c\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	return _deecaa, nil
}

// GetContext returns the annotation context which contains the specific type-dependent context.
// The context represents the subannotation.
func (_fccg *PdfAnnotation) GetContext() PdfModel {
	if _fccg == nil {
		return nil
	}
	return _fccg._cdb
}

// GetContext returns the context of the outline tree node, which is either a
// *PdfOutline or a *PdfOutlineItem. The method returns nil for uninitialized
// tree nodes.
func (_cbdbd *PdfOutlineTreeNode) GetContext() PdfModel {
	if _ebbcf, _afcdf := _cbdbd._eeedb.(*PdfOutline); _afcdf {
		return _ebbcf
	}
	if _dgda, _dcgfc := _cbdbd._eeedb.(*PdfOutlineItem); _dcgfc {
		return _dgda
	}
	_ddb.Log.Debug("\u0045\u0052RO\u0052\u0020\u0049n\u0076\u0061\u006c\u0069d o\u0075tl\u0069\u006e\u0065\u0020\u0074\u0072\u0065e \u006e\u006f\u0064\u0065\u0020\u0069\u0074e\u006d")
	return nil
}
func (_gdbc *PdfReader) newPdfActionNamedFromDict(_gdcb *_eb.PdfObjectDictionary) (*PdfActionNamed, error) {
	return &PdfActionNamed{N: _gdcb.Get("\u004e")}, nil
}

type pdfFont interface {
	_fg.Font

	// ToPdfObject returns a PDF representation of the font and implements interface Model.
	ToPdfObject() _eb.PdfObject
	getFontDescriptor() *PdfFontDescriptor
	baseFields() *fontCommon
}

// ToPdfObject implements interface PdfModel.
func (_bgee *PdfAnnotationWidget) ToPdfObject() _eb.PdfObject {
	_bgee.PdfAnnotation.ToPdfObject()
	_acgf := _bgee._ggf
	_fda := _acgf.PdfObject.(*_eb.PdfObjectDictionary)
	if _bgee._eccd {
		return _acgf
	}
	_bgee._eccd = true
	_fda.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0057\u0069\u0064\u0067\u0065\u0074"))
	_fda.SetIfNotNil("\u0048", _bgee.H)
	_fda.SetIfNotNil("\u004d\u004b", _bgee.MK)
	_fda.SetIfNotNil("\u0041", _bgee.A)
	_fda.SetIfNotNil("\u0041\u0041", _bgee.AA)
	_fda.SetIfNotNil("\u0042\u0053", _bgee.BS)
	_aadg := _bgee.Parent
	if _bgee._bca != nil {
		if _bgee._bca._adgda == _bgee._ggf {
			_bgee._bca.ToPdfObject()
		}
		_aadg = _bgee._bca.GetContainingPdfObject()
	}
	if _aadg != _acgf {
		_fda.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _aadg)
	}
	_bgee._eccd = false
	return _acgf
}
func (_caac *fontFile) loadFromSegments(_ecddf, _eegac []byte) error {
	_ddb.Log.Trace("\u006c\u006f\u0061dF\u0072\u006f\u006d\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u0020\u0025\u0064", len(_ecddf), len(_eegac))
	_aggfd := _caac.parseASCIIPart(_ecddf)
	if _aggfd != nil {
		return _aggfd
	}
	_ddb.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _caac)
	if len(_eegac) == 0 {
		return nil
	}
	_ddb.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _caac)
	return nil
}

// String returns the name of the colorspace (DeviceN).
func (_fbfaf *PdfColorspaceDeviceN) String() string { return "\u0044e\u0076\u0069\u0063\u0065\u004e" }

// PdfDate represents a date, which is a PDF string of the form:
// (D:YYYYMMDDHHmmSSOHH'mm)
type PdfDate struct {
	_fffdf  int64
	_aacegf int64
	_afgfb  int64
	_bdfbf  int64
	_aecad  int64
	_cfdb   int64
	_eaede  byte
	_fgggd  int64
	_gadga  int64
}

func (_cgee *PdfColorspaceLab) String() string { return "\u004c\u0061\u0062" }
func _bceda(_fagga []byte) ([]byte, error) {
	_fdbcc := _bd.New()
	if _, _fcffb := _bagf.Copy(_fdbcc, _dd.NewReader(_fagga)); _fcffb != nil {
		return nil, _fcffb
	}
	return _fdbcc.Sum(nil), nil
}

// GetDocMDPPermission returns the DocMDP level of the restrictions
func (_gcgcf *PdfSignature) GetDocMDPPermission() (_bab.DocMDPPermission, bool) {
	for _, _fcgfb := range _gcgcf.Reference.Elements() {
		if _bgfdf, _gcefbf := _eb.GetDict(_fcgfb); _gcefbf {
			if _abfbc, _ggae := _eb.GetNameVal(_bgfdf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _ggae && _abfbc == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				if _afbgcf, _dadca := _eb.GetDict(_bgfdf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _dadca {
					if P, _eecec := _eb.GetIntVal(_afbgcf.Get("\u0050")); _eecec {
						return _bab.DocMDPPermission(P), true
					}
				}
			}
		}
	}
	return 0, false
}

// PdfBorderStyle represents a border style dictionary (12.5.4 Border Styles p. 394).
type PdfBorderStyle struct {
	W     *float64
	S     *BorderStyle
	D     *[]int
	_gdgd _eb.PdfObject
}

func (_ebggb *PdfReader) flattenFieldsWithOpts(_ffegb bool, _bgfcd FieldAppearanceGenerator, _dead *FieldFlattenOpts) error {
	if _dead == nil {
		_dead = &FieldFlattenOpts{}
	}
	var _ceebc bool
	_dcdfc := map[*PdfAnnotation]bool{}
	{
		var _afdeb []*PdfField
		_cfdfb := _ebggb.AcroForm
		if _cfdfb != nil {
			if _dead.FilterFunc != nil {
				_afdeb = _cfdfb.filteredFields(_dead.FilterFunc, true)
				_ceebc = _cfdfb.Fields != nil && len(*_cfdfb.Fields) > 0
			} else {
				_afdeb = _cfdfb.AllFields()
			}
		}
		for _, _febd := range _afdeb {
			if len(_febd.Annotations) < 1 {
				_ddb.Log.Debug("\u004e\u006f\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006ed\u0020\u0066\u006f\u0072\u003a\u0020\u0025v\u002c\u0020\u006c\u006f\u006f\u006b\u0020\u0069\u006e\u0074\u006f \u004b\u0069\u0064\u0073\u0020\u004f\u0062\u006a\u0065\u0063\u0074", _febd.PartialName())
				for _fbag, _dbgbb := range _febd.Kids {
					for _, _gdag := range _dbgbb.Annotations {
						_dcdfc[_gdag.PdfAnnotation] = _febd.V != nil
						if _dbgbb.V == nil {
							_dbgbb.V = _febd.V
						}
						if _dbgbb.T == nil {
							_dbgbb.T = _eb.MakeString(_e.Sprintf("\u0025\u0073\u0023%\u0064", _febd.PartialName(), _fbag))
						}
						if _bgfcd != nil {
							_ceacg, _bbeg := _bgfcd.GenerateAppearanceDict(_cfdfb, _dbgbb, _gdag)
							if _bbeg != nil {
								return _bbeg
							}
							_gdag.AP = _ceacg
						}
					}
				}
			}
			for _, _efecg := range _febd.Annotations {
				_dcdfc[_efecg.PdfAnnotation] = _febd.V != nil
				if _bgfcd != nil {
					_dccd, _accec := _bgfcd.GenerateAppearanceDict(_cfdfb, _febd, _efecg)
					if _accec != nil {
						return _accec
					}
					_efecg.AP = _dccd
				}
			}
		}
	}
	if _ffegb {
		for _, _ebfe := range _ebggb.PageList {
			_fgdec, _fbgdf := _ebfe.GetAnnotations()
			if _fbgdf != nil {
				return _fbgdf
			}
			for _, _gggg := range _fgdec {
				_dcdfc[_gggg] = true
			}
		}
	}
	for _, _dbde := range _ebggb.PageList {
		_aaaba := _dbde.flattenFieldsWithOpts(_bgfcd, _dead, _dcdfc)
		if _aaaba != nil {
			return _aaaba
		}
	}
	if !_ceebc {
		_ebggb.AcroForm = nil
	}
	return nil
}

// GetOptimizer returns current PDF optimizer.
func (_cdecc *PdfWriter) GetOptimizer() Optimizer { return _cdecc._fbaag }

// HasExtGState checks whether a font is defined by the specified keyName.
func (_fbacc *PdfPageResources) HasExtGState(keyName _eb.PdfObjectName) bool {
	_, _aeegb := _fbacc.GetFontByName(keyName)
	return _aeegb
}

// A returns the value of the A component of the color.
func (_gccg *PdfColorCalRGB) A() float64 { return _gccg[0] }

// SetContext sets the sub annotation (context).
func (_cgd *PdfAnnotation) SetContext(ctx PdfModel) { _cgd._cdb = ctx }
func (_gbga *PdfReader) newPdfAcroFormFromDict(_dgce *_eb.PdfIndirectObject, _cbffd *_eb.PdfObjectDictionary) (*PdfAcroForm, error) {
	_cafda := NewPdfAcroForm()
	if _dgce != nil {
		_cafda._fbgad = _dgce
		_dgce.PdfObject = _eb.MakeDict()
	}
	if _aeacf := _cbffd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"); _aeacf != nil && !_eb.IsNullObject(_aeacf) {
		_dgcd, _egdad := _eb.GetArray(_aeacf)
		if !_egdad {
			return nil, _e.Errorf("\u0066i\u0065\u006c\u0064\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _aeacf)
		}
		var _ceafe []*PdfField
		for _, _beeae := range _dgcd.Elements() {
			_cced, _ddegd := _eb.GetIndirect(_beeae)
			if !_ddegd {
				if _, _bacg := _beeae.(*_eb.PdfObjectNull); _bacg {
					_ddb.Log.Trace("\u0053k\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072 \u006e\u0075\u006c\u006c\u0020\u0066\u0069\u0065\u006c\u0064")
					continue
				}
				_ddb.Log.Debug("\u0046\u0069\u0065\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0064 \u0069\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _beeae)
				return nil, _e.Errorf("\u0066\u0069\u0065l\u0064\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_ebce, _dbeaa := _gbga.newPdfFieldFromIndirectObject(_cced, nil)
			if _dbeaa != nil {
				return nil, _dbeaa
			}
			_ddb.Log.Trace("\u0041\u0063\u0072\u006fFo\u0072\u006d\u0020\u0046\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u002b\u0076", *_ebce)
			_ceafe = append(_ceafe, _ebce)
		}
		_cafda.Fields = &_ceafe
	}
	if _cefg := _cbffd.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _cefg != nil {
		_feef, _bbbfb := _eb.GetBool(_cefg)
		if _bbbfb {
			_cafda.NeedAppearances = _feef
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u0065\u0065\u0064\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063e\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _cefg)
		}
	}
	if _abcab := _cbffd.Get("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073"); _abcab != nil {
		_bgbe, _ffbde := _eb.GetInt(_abcab)
		if _ffbde {
			_cafda.SigFlags = _bgbe
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073 \u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _abcab)
		}
	}
	if _fbfc := _cbffd.Get("\u0043\u004f"); _fbfc != nil {
		_aaggc, _bedgb := _eb.GetArray(_fbfc)
		if _bedgb {
			_cafda.CO = _aaggc
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u004f\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _fbfc)
		}
	}
	if _bcbda := _cbffd.Get("\u0044\u0052"); _bcbda != nil {
		if _baae, _gbbff := _eb.GetDict(_bcbda); _gbbff {
			_fdcbe, _ddgb := NewPdfPageResourcesFromDict(_baae)
			if _ddgb != nil {
				_ddb.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0044R\u003a\u0020\u0025\u0076", _ddgb)
				return nil, _ddgb
			}
			_cafda.DR = _fdcbe
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0052\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _bcbda)
		}
	}
	if _bceb := _cbffd.Get("\u0044\u0041"); _bceb != nil {
		_gbee, _adcag := _eb.GetString(_bceb)
		if _adcag {
			_cafda.DA = _gbee
		} else {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0041\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _bceb)
		}
	}
	if _bffda := _cbffd.Get("\u0051"); _bffda != nil {
		_ffcf, _dccag := _eb.GetInt(_bffda)
		if _dccag {
			_cafda.Q = _ffcf
		} else {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0051\u0020\u0069\u006e\u0076a\u006ci\u0064 \u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _bffda)
		}
	}
	if _cgeff := _cbffd.Get("\u0058\u0046\u0041"); _cgeff != nil {
		_cafda.XFA = _cgeff
	}
	if _addg := _cbffd.Get("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e"); _addg != nil {
		_cafda.ADBEEchoSign = _addg
	}
	_cafda.ToPdfObject()
	return _cafda, nil
}

// NewPdfAnnotationCaret returns a new caret annotation.
func NewPdfAnnotationCaret() *PdfAnnotationCaret {
	_gde := NewPdfAnnotation()
	_dae := &PdfAnnotationCaret{}
	_dae.PdfAnnotation = _gde
	_dae.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gde.SetContext(_dae)
	return _dae
}

// NewPdfAnnotationTrapNet returns a new trapnet annotation.
func NewPdfAnnotationTrapNet() *PdfAnnotationTrapNet {
	_gbb := NewPdfAnnotation()
	_gfaa := &PdfAnnotationTrapNet{}
	_gfaa.PdfAnnotation = _gbb
	_gbb.SetContext(_gfaa)
	return _gfaa
}

// ToPdfObject implements interface PdfModel.
func (_gga *PdfActionLaunch) ToPdfObject() _eb.PdfObject {
	_gga.PdfAction.ToPdfObject()
	_dbf := _gga._dee
	_feb := _dbf.PdfObject.(*_eb.PdfObjectDictionary)
	_feb.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeLaunch)))
	if _gga.F != nil {
		_feb.Set("\u0046", _gga.F.ToPdfObject())
	}
	_feb.SetIfNotNil("\u0057\u0069\u006e", _gga.Win)
	_feb.SetIfNotNil("\u004d\u0061\u0063", _gga.Mac)
	_feb.SetIfNotNil("\u0055\u006e\u0069\u0078", _gga.Unix)
	_feb.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _gga.NewWindow)
	return _dbf
}
func (_fcca Image) getBase() _df.ImageBase {
	return _df.NewImageBase(int(_fcca.Width), int(_fcca.Height), int(_fcca.BitsPerComponent), _fcca.ColorComponents, _fcca.Data, _fcca._bdcab, _fcca._fedc)
}

var (
	_dfbafc _ec.Mutex
	_aeee   = ""
	_ebffa  _d.Time
	_geba   = ""
	_addbb  = ""
	_eegaeb _d.Time
	_ccgde  = ""
	_bgabb  = ""
	_gacdf  = ""
)

func (_babbd *pdfCIDFontType0) baseFields() *fontCommon { return &_babbd.fontCommon }

// GetContainingPdfObject returns the container of the shading object (indirect object).
func (_dbca *PdfShading) GetContainingPdfObject() _eb.PdfObject { return _dbca._cefaa }

// SetContentStream updates the content stream with specified encoding.
// If encoding is null, will use the xform.Filter object or Raw encoding if not set.
func (_dbcea *XObjectForm) SetContentStream(content []byte, encoder _eb.StreamEncoder) error {
	_baafg := content
	if encoder == nil {
		if _dbcea.Filter != nil {
			encoder = _dbcea.Filter
		} else {
			encoder = _eb.NewRawEncoder()
		}
	}
	_dadeeb, _cdef := encoder.EncodeBytes(_baafg)
	if _cdef != nil {
		return _cdef
	}
	_baafg = _dadeeb
	_dbcea.Stream = _baafg
	_dbcea.Filter = encoder
	return nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_fedcc *PdfShadingType6) ToPdfObject() _eb.PdfObject {
	_fedcc.PdfShading.ToPdfObject()
	_ggddf, _bfca := _fedcc.getShadingDict()
	if _bfca != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _fedcc.BitsPerCoordinate != nil {
		_ggddf.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _fedcc.BitsPerCoordinate)
	}
	if _fedcc.BitsPerComponent != nil {
		_ggddf.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _fedcc.BitsPerComponent)
	}
	if _fedcc.BitsPerFlag != nil {
		_ggddf.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _fedcc.BitsPerFlag)
	}
	if _fedcc.Decode != nil {
		_ggddf.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _fedcc.Decode)
	}
	if _fedcc.Function != nil {
		if len(_fedcc.Function) == 1 {
			_ggddf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _fedcc.Function[0].ToPdfObject())
		} else {
			_ffbdd := _eb.MakeArray()
			for _, _cfcc := range _fedcc.Function {
				_ffbdd.Append(_cfcc.ToPdfObject())
			}
			_ggddf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ffbdd)
		}
	}
	return _fedcc._cefaa
}

// NewPdfColorPatternType3 returns an empty color shading pattern type 3 (Radial).
func NewPdfColorPatternType3() *PdfColorPatternType3 { _gbdc := &PdfColorPatternType3{}; return _gbdc }
func (_eaecd *LTV) getCerts(_agaac []*_bag.Certificate) ([][]byte, error) {
	_eefaf := make([][]byte, 0, len(_agaac))
	for _, _dfacg := range _agaac {
		_eefaf = append(_eefaf, _dfacg.Raw)
	}
	return _eefaf, nil
}

// ColorToRGB converts an Indexed color to an RGB color.
func (_dacfg *PdfColorspaceSpecialIndexed) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _dacfg.Base == nil {
		return nil, _dcf.New("\u0069\u006e\u0064\u0065\u0078\u0065d\u0020\u0062\u0061\u0073\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _dacfg.Base.ColorToRGB(color)
}

// PdfTilingPattern is a Tiling pattern that consists of repetitions of a pattern cell with defined intervals.
// It is a type 1 pattern. (PatternType = 1).
// A tiling pattern is represented by a stream object, where the stream content is
// a content stream that describes the pattern cell.
type PdfTilingPattern struct {
	*PdfPattern
	PaintType  *_eb.PdfObjectInteger
	TilingType *_eb.PdfObjectInteger
	BBox       *PdfRectangle
	XStep      *_eb.PdfObjectFloat
	YStep      *_eb.PdfObjectFloat
	Resources  *PdfPageResources
	Matrix     *_eb.PdfObjectArray
}

// XObjectForm (Table 95 in 8.10.2).
type XObjectForm struct {
	Filter        _eb.StreamEncoder
	FormType      _eb.PdfObject
	BBox          _eb.PdfObject
	Matrix        _eb.PdfObject
	Resources     *PdfPageResources
	Group         _eb.PdfObject
	Ref           _eb.PdfObject
	MetaData      _eb.PdfObject
	PieceInfo     _eb.PdfObject
	LastModified  _eb.PdfObject
	StructParent  _eb.PdfObject
	StructParents _eb.PdfObject
	OPI           _eb.PdfObject
	OC            _eb.PdfObject
	Name          _eb.PdfObject

	// Stream data.
	Stream []byte
	_afcag *_eb.PdfObjectStream
}

// ToPdfObject implements interface PdfModel.
func (_debb *EmbeddedFile) ToPdfObject() _eb.PdfObject {
	_cgefa := _eb.NewFlateEncoder()
	_cgde, _bcff := _eb.MakeStream(_debb.Content, _cgefa)
	if _bcff != nil {
		_ddb.Log.Debug("\u0046\u0061\u0069\u006c\u0065d\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0065\u006db\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _bcff)
		return nil
	}
	_cefaf := _cgde.PdfObjectDictionary
	_cefaf.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0045\u006d\u0062e\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065"))
	_cefaf.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeEncodedString("\u0074\u0065\u0078\u0074\u002f\u0070\u006c\u0061\u0069\u006e", true))
	_fcdec := _eb.MakeDict()
	_fcdec.Set("\u0043\u0068\u0065\u0063\u006b\u0053\u0075\u006d", _eb.MakeString(_debb.Hash[:]))
	_fcdec.Set("\u0053\u0069\u007a\u0065", _eb.MakeInteger(int64(len(_debb.Content))))
	_efdd := _debb.CreationTime
	if _efdd.IsZero() {
		_efdd = _d.Now()
	}
	_dbfb := _debb.ModTime
	if _dbfb.IsZero() {
		_dbfb = _efdd
	}
	_cgdf, _bcff := NewPdfDateFromTime(_efdd)
	if _bcff != nil {
		_ddb.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074o\u0020\u0063\u0072ea\u0074\u0065\u0020\u0065\u006d\u0062e\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0063\u0072\u0065\u0061\u0074i\u006f\u006e\u0020\u0064\u0061\u0074\u0065\u003a \u0025\u0076", _bcff)
		return nil
	}
	_gcdcg, _bcff := NewPdfDateFromTime(_dbfb)
	if _bcff != nil {
		_ddb.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074o\u0020\u0063\u0072ea\u0074\u0065\u0020\u0065\u006d\u0062e\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0063\u0072\u0065\u0061\u0074i\u006f\u006e\u0020\u0064\u0061\u0074\u0065\u003a \u0025\u0076", _bcff)
		return nil
	}
	_fcdec.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _cgdf.ToPdfObject())
	_fcdec.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _gcdcg.ToPdfObject())
	_cefaf.Set("\u0050\u0061\u0072\u0061\u006d\u0073", _fcdec)
	_aece := _eb.MakeDict()
	_aece.Set(*_eb.MakeName("\u0046"), _cgde)
	return _aece
}

const (
	_eaggc  = 0x00001
	_dcaf   = 0x00002
	_bdcce  = 0x00004
	_fgbece = 0x00008
	_ebaed  = 0x00020
	_ffbe   = 0x00040
	_bcdae  = 0x10000
	_dfcd   = 0x20000
	_gfegd  = 0x40000
)

// ToPdfObject converts the pdfFontSimple to its PDF representation for outputting.
func (_eggag *pdfFontSimple) ToPdfObject() _eb.PdfObject {
	if _eggag._ffgdb == nil {
		_eggag._ffgdb = &_eb.PdfIndirectObject{}
	}
	_ffbd := _eggag.baseFields().asPdfObjectDictionary("")
	_eggag._ffgdb.PdfObject = _ffbd
	if _eggag.FirstChar != nil {
		_ffbd.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _eggag.FirstChar)
	}
	if _eggag.LastChar != nil {
		_ffbd.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _eggag.LastChar)
	}
	if _eggag.Widths != nil {
		_ffbd.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _eggag.Widths)
	}
	if _eggag.Encoding != nil {
		_ffbd.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _eggag.Encoding)
	} else if _eggag._eccbb != nil {
		_adee := _eggag._eccbb.ToPdfObject()
		if _adee != nil {
			_ffbd.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _adee)
		}
	}
	return _eggag._ffgdb
}

// HasFontByName checks whether a font is defined by the specified keyName.
func (_fagee *PdfPageResources) HasFontByName(keyName _eb.PdfObjectName) bool {
	_, _cafdg := _fagee.GetFontByName(keyName)
	return _cafdg
}
func (_dfaab *PdfReader) newPdfSignatureFromIndirect(_ggga *_eb.PdfIndirectObject) (*PdfSignature, error) {
	_bbdgg, _dgegb := _ggga.PdfObject.(*_eb.PdfObjectDictionary)
	if !_dgegb {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		return nil, ErrTypeCheck
	}
	if _ccfdd, _ccaae := _dfaab._affaf.GetModelFromPrimitive(_ggga).(*PdfSignature); _ccaae {
		return _ccfdd, nil
	}
	_efacg := &PdfSignature{}
	_efacg._cddce = _ggga
	_efacg.Type, _ = _eb.GetName(_bbdgg.Get("\u0054\u0079\u0070\u0065"))
	_efacg.Filter, _dgegb = _eb.GetName(_bbdgg.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if !_dgegb {
		_ddb.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053i\u0067\u006e\u0061\u0074\u0075r\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_efacg.SubFilter, _ = _eb.GetName(_bbdgg.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r"))
	_efacg.Contents, _dgegb = _eb.GetString(_bbdgg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_dgegb {
		_ddb.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	if _bgedd, _gefbc := _eb.GetArray(_bbdgg.Get("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e")); _gefbc {
		_efacg.Reference = _eb.MakeArray()
		for _, _aaaag := range _bgedd.Elements() {
			_fggdb, _egeaa := _eb.GetDict(_aaaag)
			if !_egeaa {
				_ddb.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020R\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0061\u0074\u0065\u0064")
				return nil, ErrInvalidAttribute
			}
			_aaede, _gedce := _dfaab.newPdfSignatureReferenceFromDict(_fggdb)
			if _gedce != nil {
				return nil, _gedce
			}
			_efacg.Reference.Append(_aaede.ToPdfObject())
		}
	}
	_efacg.Cert = _bbdgg.Get("\u0043\u0065\u0072\u0074")
	_efacg.ByteRange, _ = _eb.GetArray(_bbdgg.Get("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e"))
	_efacg.Changes, _ = _eb.GetArray(_bbdgg.Get("\u0043h\u0061\u006e\u0067\u0065\u0073"))
	_efacg.Name, _ = _eb.GetString(_bbdgg.Get("\u004e\u0061\u006d\u0065"))
	_efacg.M, _ = _eb.GetString(_bbdgg.Get("\u004d"))
	_efacg.Location, _ = _eb.GetString(_bbdgg.Get("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"))
	_efacg.Reason, _ = _eb.GetString(_bbdgg.Get("\u0052\u0065\u0061\u0073\u006f\u006e"))
	_efacg.ContactInfo, _ = _eb.GetString(_bbdgg.Get("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f"))
	_efacg.R, _ = _eb.GetInt(_bbdgg.Get("\u0052"))
	_efacg.V, _ = _eb.GetInt(_bbdgg.Get("\u0056"))
	_efacg.PropBuild, _ = _eb.GetDict(_bbdgg.Get("\u0050\u0072\u006f\u0070\u005f\u0042\u0075\u0069\u006c\u0064"))
	_efacg.PropAuthTime, _ = _eb.GetInt(_bbdgg.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0069\u006d\u0065"))
	_efacg.PropAuthType, _ = _eb.GetName(_bbdgg.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0079\u0070\u0065"))
	_dfaab._affaf.Register(_ggga, _efacg)
	return _efacg, nil
}

// ToPdfObject implements interface PdfModel.
func (_gfdbg *PdfSignature) ToPdfObject() _eb.PdfObject {
	_gcceg := _gfdbg._cddce
	var _gdab *_eb.PdfObjectDictionary
	if _abdd, _abdbd := _gcceg.PdfObject.(*pdfSignDictionary); _abdbd {
		_gdab = _abdd.PdfObjectDictionary
	} else {
		_gdab = _gcceg.PdfObject.(*_eb.PdfObjectDictionary)
	}
	_gdab.SetIfNotNil("\u0054\u0079\u0070\u0065", _gfdbg.Type)
	_gdab.SetIfNotNil("\u0046\u0069\u006c\u0074\u0065\u0072", _gfdbg.Filter)
	_gdab.SetIfNotNil("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r", _gfdbg.SubFilter)
	_gdab.SetIfNotNil("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _gfdbg.ByteRange)
	_gdab.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _gfdbg.Contents)
	_gdab.SetIfNotNil("\u0043\u0065\u0072\u0074", _gfdbg.Cert)
	_gdab.SetIfNotNil("\u004e\u0061\u006d\u0065", _gfdbg.Name)
	_gdab.SetIfNotNil("\u0052\u0065\u0061\u0073\u006f\u006e", _gfdbg.Reason)
	_gdab.SetIfNotNil("\u004d", _gfdbg.M)
	_gdab.SetIfNotNil("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e", _gfdbg.Reference)
	_gdab.SetIfNotNil("\u0043h\u0061\u006e\u0067\u0065\u0073", _gfdbg.Changes)
	_gdab.SetIfNotNil("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f", _gfdbg.ContactInfo)
	_gdab.SetIfNotNil("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e", _gfdbg.Location)
	return _gcceg
}

// PdfInfoTrapped specifies pdf trapped information.
type PdfInfoTrapped string

// GetPerms returns the Permissions dictionary
func (_gabe *PdfReader) GetPerms() *Permissions { return _gabe._bdffe }
func (_adgfa *PdfWriter) seekByName(_egdcb _eb.PdfObject, _faggac []string, _bcbbgc string) ([]_eb.PdfObject, error) {
	_ddb.Log.Trace("\u0053\u0065\u0065\u006b\u0020\u0062\u0079\u0020\u006e\u0061\u006d\u0065.\u002e\u0020\u0025\u0054", _egdcb)
	var _bfdgg []_eb.PdfObject
	if _bfggg, _agfaf := _egdcb.(*_eb.PdfIndirectObject); _agfaf {
		return _adgfa.seekByName(_bfggg.PdfObject, _faggac, _bcbbgc)
	}
	if _fabd, _ccbdga := _egdcb.(*_eb.PdfObjectStream); _ccbdga {
		return _adgfa.seekByName(_fabd.PdfObjectDictionary, _faggac, _bcbbgc)
	}
	if _egedc, _bbdgb := _egdcb.(*_eb.PdfObjectDictionary); _bbdgb {
		_ddb.Log.Trace("\u0044\u0069\u0063\u0074")
		for _, _aaedc := range _egedc.Keys() {
			_ecgaad := _egedc.Get(_aaedc)
			if string(_aaedc) == _bcbbgc {
				_bfdgg = append(_bfdgg, _ecgaad)
			}
			for _, _dfgge := range _faggac {
				if string(_aaedc) == _dfgge {
					_ddb.Log.Trace("\u0046\u006f\u006c\u006c\u006f\u0077\u0020\u006b\u0065\u0079\u0020\u0025\u0073", _dfgge)
					_dbeba, _adccd := _adgfa.seekByName(_ecgaad, _faggac, _bcbbgc)
					if _adccd != nil {
						return _bfdgg, _adccd
					}
					_bfdgg = append(_bfdgg, _dbeba...)
					break
				}
			}
		}
		return _bfdgg, nil
	}
	return _bfdgg, nil
}
func (_ceeaa *PdfColorspaceSpecialIndexed) String() string {
	return "\u0049n\u0064\u0065\u0078\u0065\u0064"
}
func (_ggba *PdfReader) newPdfActionSubmitFormFromDict(_ccfe *_eb.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	_bagdg, _ecc := _dba(_ccfe.Get("\u0046"))
	if _ecc != nil {
		return nil, _ecc
	}
	return &PdfActionSubmitForm{F: _bagdg, Fields: _ccfe.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _ccfe.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// MergePageWith appends page content to source Pdf file page content.
func (_faf *PdfAppender) MergePageWith(pageNum int, page *PdfPage) error {
	_adea := pageNum - 1
	var _dbdf *PdfPage
	for _fgdf, _gdge := range _faf._fedg {
		if _fgdf == _adea {
			_dbdf = _gdge
		}
	}
	if _dbdf == nil {
		return _e.Errorf("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073o\u0075\u0072\u0063\u0065\u0020\u0064o\u0063\u0075\u006de\u006e\u0074", pageNum)
	}
	if _dbdf._efcff != nil && _dbdf._efcff.GetParser() == _faf._adac._ebbe {
		_dbdf = _dbdf.Duplicate()
		_faf._fedg[_adea] = _dbdf
	}
	page = page.Duplicate()
	_bfcca(page)
	_fgde := _adbb(_dbdf)
	_ceac := _adbb(page)
	_gdda := make(map[_eb.PdfObjectName]_eb.PdfObjectName)
	for _dcbb := range _ceac {
		if _, _dfbc := _fgde[_dcbb]; _dfbc {
			for _gfgd := 1; true; _gfgd++ {
				_fcde := _eb.PdfObjectName(string(_dcbb) + _de.Itoa(_gfgd))
				if _, _afdae := _fgde[_fcde]; !_afdae {
					_gdda[_dcbb] = _fcde
					break
				}
			}
		}
	}
	_eba, _gccc := page.GetContentStreams()
	if _gccc != nil {
		return _gccc
	}
	_fbed, _gccc := _dbdf.GetContentStreams()
	if _gccc != nil {
		return _gccc
	}
	for _bdcb, _bgdd := range _eba {
		for _ebga, _bafc := range _gdda {
			_bgdd = _cc.Replace(_bgdd, "\u002f"+string(_ebga), "\u002f"+string(_bafc), -1)
		}
		_eba[_bdcb] = _bgdd
	}
	_fbed = append(_fbed, _eba...)
	if _fbdc := _dbdf.SetContentStreams(_fbed, _eb.NewFlateEncoder()); _fbdc != nil {
		return _fbdc
	}
	_dbdf._dbga = append(_dbdf._dbga, page._dbga...)
	if _dbdf.Resources == nil {
		_dbdf.Resources = NewPdfPageResources()
	}
	if page.Resources != nil {
		_dbdf.Resources.Font = _faf.mergeResources(_dbdf.Resources.Font, page.Resources.Font, _gdda)
		_dbdf.Resources.XObject = _faf.mergeResources(_dbdf.Resources.XObject, page.Resources.XObject, _gdda)
		_dbdf.Resources.Properties = _faf.mergeResources(_dbdf.Resources.Properties, page.Resources.Properties, _gdda)
		if _dbdf.Resources.ProcSet == nil {
			_dbdf.Resources.ProcSet = page.Resources.ProcSet
		}
		_dbdf.Resources.Shading = _faf.mergeResources(_dbdf.Resources.Shading, page.Resources.Shading, _gdda)
		_dbdf.Resources.ExtGState = _faf.mergeResources(_dbdf.Resources.ExtGState, page.Resources.ExtGState, _gdda)
	}
	_bfg, _gccc := _dbdf.GetMediaBox()
	if _gccc != nil {
		return _gccc
	}
	_gffed, _gccc := page.GetMediaBox()
	if _gccc != nil {
		return _gccc
	}
	var _ececa bool
	if _bfg.Llx > _gffed.Llx {
		_bfg.Llx = _gffed.Llx
		_ececa = true
	}
	if _bfg.Lly > _gffed.Lly {
		_bfg.Lly = _gffed.Lly
		_ececa = true
	}
	if _bfg.Urx < _gffed.Urx {
		_bfg.Urx = _gffed.Urx
		_ececa = true
	}
	if _bfg.Ury < _gffed.Ury {
		_bfg.Ury = _gffed.Ury
		_ececa = true
	}
	if _ececa {
		_dbdf.MediaBox = _bfg
	}
	return nil
}

// PdfFunction interface represents the common methods of a function in PDF.
type PdfFunction interface {
	Evaluate([]float64) ([]float64, error)
	ToPdfObject() _eb.PdfObject
}

// SetViewArea sets the value of the viewArea.
func (_ggece *ViewerPreferences) SetViewArea(viewArea PageBoundary) { _ggece._ffgda = viewArea }

// BorderStyle defines border type, typically used for annotations.
type BorderStyle int

// PdfShadingType4 is a Free-form Gouraud-shaded triangle mesh.
type PdfShadingType4 struct {
	*PdfShading
	BitsPerCoordinate *_eb.PdfObjectInteger
	BitsPerComponent  *_eb.PdfObjectInteger
	BitsPerFlag       *_eb.PdfObjectInteger
	Decode            *_eb.PdfObjectArray
	Function          []PdfFunction
}

// ToPdfObject implements interface PdfModel.
func (_bage *PdfActionHide) ToPdfObject() _eb.PdfObject {
	_bage.PdfAction.ToPdfObject()
	_gdc := _bage._dee
	_bagd := _gdc.PdfObject.(*_eb.PdfObjectDictionary)
	_bagd.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeHide)))
	_bagd.SetIfNotNil("\u0054", _bage.T)
	_bagd.SetIfNotNil("\u0048", _bage.H)
	return _gdc
}

// GetContainingPdfObject implements interface PdfModel.
func (_abceg *Permissions) GetContainingPdfObject() _eb.PdfObject { return _abceg._cgbag }
func _ecdgf(_afaca _eb.PdfObject) (*PdfColorspaceLab, error) {
	_agfab := NewPdfColorspaceLab()
	if _gcbbd, _afegc := _afaca.(*_eb.PdfIndirectObject); _afegc {
		_agfab._cbag = _gcbbd
	}
	_afaca = _eb.TraceToDirectObject(_afaca)
	_fgbec, _cbgbg := _afaca.(*_eb.PdfObjectArray)
	if !_cbgbg {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _fgbec.Len() != 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_afaca = _eb.TraceToDirectObject(_fgbec.Get(0))
	_daga, _cbgbg := _afaca.(*_eb.PdfObjectName)
	if !_cbgbg {
		return nil, _e.Errorf("\u006c\u0061\u0062\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if *_daga != "\u004c\u0061\u0062" {
		return nil, _e.Errorf("n\u006ft\u0020\u0061\u0020\u004c\u0061\u0062\u0020\u0063o\u006c\u006f\u0072\u0073pa\u0063\u0065")
	}
	_afaca = _eb.TraceToDirectObject(_fgbec.Get(1))
	_aegdf, _cbgbg := _afaca.(*_eb.PdfObjectDictionary)
	if !_cbgbg {
		return nil, _e.Errorf("c\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006ev\u0061l\u0069\u0064")
	}
	_afaca = _aegdf.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_afaca = _eb.TraceToDirectObject(_afaca)
	_cbee, _cbgbg := _afaca.(*_eb.PdfObjectArray)
	if !_cbgbg {
		return nil, _e.Errorf("\u004c\u0061\u0062\u0020In\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069n\u0074")
	}
	if _cbee.Len() != 3 {
		return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
	}
	_efgc, _edcca := _cbee.GetAsFloat64Slice()
	if _edcca != nil {
		return nil, _edcca
	}
	_agfab.WhitePoint = _efgc
	_afaca = _aegdf.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _afaca != nil {
		_afaca = _eb.TraceToDirectObject(_afaca)
		_bagb, _fbaf := _afaca.(*_eb.PdfObjectArray)
		if !_fbaf {
			return nil, _e.Errorf("\u004c\u0061\u0062: \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
		}
		if _bagb.Len() != 3 {
			return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
		}
		_gfda, _aaed := _bagb.GetAsFloat64Slice()
		if _aaed != nil {
			return nil, _aaed
		}
		_agfab.BlackPoint = _gfda
	}
	_afaca = _aegdf.Get("\u0052\u0061\u006eg\u0065")
	if _afaca != nil {
		_afaca = _eb.TraceToDirectObject(_afaca)
		_fag, _dcdg := _afaca.(*_eb.PdfObjectArray)
		if !_dcdg {
			_ddb.Log.Error("\u0052\u0061n\u0067\u0065\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
			return nil, _e.Errorf("\u004ca\u0062:\u0020\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if _fag.Len() != 4 {
			_ddb.Log.Error("\u0052\u0061\u006e\u0067\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020e\u0072\u0072\u006f\u0072")
			return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0052a\u006e\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bgcg, _bffc := _fag.GetAsFloat64Slice()
		if _bffc != nil {
			return nil, _bffc
		}
		_agfab.Range = _bgcg
	}
	return _agfab, nil
}
func _eeaaf(_dffee string) map[string]string {
	_adba := _aefb.Split(_dffee, -1)
	_eebee := map[string]string{}
	for _, _ccda := range _adba {
		_bfgf := _agffe.FindStringSubmatch(_ccda)
		if _bfgf == nil {
			continue
		}
		_dccdbd, _dggde := _bfgf[1], _bfgf[2]
		_eebee[_dccdbd] = _dggde
	}
	return _eebee
}

// NewPdfAnnotationText returns a new text annotation.
func NewPdfAnnotationText() *PdfAnnotationText {
	_eda := NewPdfAnnotation()
	_gdcfg := &PdfAnnotationText{}
	_gdcfg.PdfAnnotation = _eda
	_gdcfg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_eda.SetContext(_gdcfg)
	return _gdcfg
}

// EnableChain adds the specified certificate chain and validation data (OCSP
// and CRL information) for it to the global scope of the document DSS. The
// added data is used for validating any of the signatures present in the
// document. The LTV client attempts to build the certificate chain up to a
// trusted root by downloading any missing certificates.
func (_afdef *LTV) EnableChain(chain []*_bag.Certificate) error { return _afdef.enable(nil, chain, "") }
func (_ffbgc *PdfWriter) checkCrossReferenceStream() bool {
	_ggdge := _ffbgc._edbbf.Major > 1 || (_ffbgc._edbbf.Major == 1 && _ffbgc._edbbf.Minor > 4)
	if _ffbgc._bddggg != nil {
		_ggdge = *_ffbgc._bddggg
	}
	return _ggdge
}

// SetXObjectImageByNameLazy adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_gdbcfd *PdfPageResources) SetXObjectImageByNameLazy(keyName _eb.PdfObjectName, ximg *XObjectImage, lazy bool) error {
	_cbefb := ximg.ToPdfObject().(*_eb.PdfObjectStream)
	if lazy {
		_cbefb.MakeLazy()
	}
	_aadbac := _gdbcfd.SetXObjectByName(keyName, _cbefb)
	return _aadbac
}

// SetCatalogMetadata sets the catalog metadata (XMP) stream object.
func (_ebbdd *PdfWriter) SetCatalogMetadata(meta _eb.PdfObject) error {
	if meta == nil {
		_ebbdd._dbffa.Remove("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
		return nil
	}
	_eeedg, _dbfeb := _eb.GetStream(meta)
	if !_dbfeb {
		return _dcf.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006d\u0065\u0074\u0061\u0064a\u0074\u0061\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_ebbdd.addObject(_eeedg)
	_ebbdd._dbffa.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _eeedg)
	return nil
}

// ViewClip returns the value of the viewClip.
func (_acfc *ViewerPreferences) ViewClip() PageBoundary { return _acfc._eadae }

// ToPdfObject converts the PdfPage to a dictionary within an indirect object container.
func (_dceac *PdfPage) ToPdfObject() _eb.PdfObject {
	_deebe := _dceac._efcff
	_dceac.GetPageDict()
	return _deebe
}
func (_bddab *PdfReader) loadOutlines() (*PdfOutlineTreeNode, error) {
	if _bddab._ebbe.GetCrypter() != nil && !_bddab._ebbe.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_adbad := _bddab._bagcfd
	_abcfd := _adbad.Get("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	if _abcfd == nil {
		return nil, nil
	}
	_ddb.Log.Trace("\u002d\u0048\u0061\u0073\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	_ffbeb := _eb.ResolveReference(_abcfd)
	_ddb.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0020\u0072\u006f\u006f\u0074\u003a\u0020\u0025\u0076", _ffbeb)
	if _cabddb := _eb.IsNullObject(_ffbeb); _cabddb {
		_ddb.Log.Trace("\u004f\u0075\u0074li\u006e\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069s\u0020n\u0075l\u006c \u002d\u0020\u006e\u006f\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
		return nil, nil
	}
	_afcdfb, _cffcc := _ffbeb.(*_eb.PdfIndirectObject)
	if !_cffcc {
		if _, _deefc := _eb.GetDict(_ffbeb); !_deefc {
			_ddb.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			return nil, nil
		}
		_ddb.Log.Debug("\u004f\u0075t\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u0069s\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		_afcdfb = _eb.MakeIndirectObject(_ffbeb)
	}
	_cabeb, _cffcc := _afcdfb.PdfObject.(*_eb.PdfObjectDictionary)
	if !_cffcc {
		return nil, _dcf.New("\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y")
	}
	_ddb.Log.Trace("O\u0075\u0074\u006c\u0069ne\u0020r\u006f\u006f\u0074\u0020\u0064i\u0063\u0074\u003a\u0020\u0025\u0076", _cabeb)
	_gbef, _, _egfda := _bddab.buildOutlineTree(_afcdfb, nil, nil, nil)
	if _egfda != nil {
		return nil, _egfda
	}
	_ddb.Log.Trace("\u0052\u0065\u0073\u0075\u006c\u0074\u0069\u006e\u0067\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065:\u0020\u0025\u0076", _gbef)
	return _gbef, nil
}

// PdfActionNamed represents a named action.
type PdfActionNamed struct {
	*PdfAction
	N _eb.PdfObject
}

// SetDirection sets the value of the direction.
func (_dbaag *ViewerPreferences) SetDirection(direction Direction) { _dbaag._baeb = direction }
func (_cabf *PdfColorspaceCalRGB) String() string                  { return "\u0043\u0061\u006c\u0052\u0047\u0042" }

const (
	XObjectTypeUndefined XObjectType = iota
	XObjectTypeImage
	XObjectTypeForm
	XObjectTypePS
	XObjectTypeUnknown
)

// PdfShadingType7 is a Tensor-product patch mesh.
type PdfShadingType7 struct {
	*PdfShading
	BitsPerCoordinate *_eb.PdfObjectInteger
	BitsPerComponent  *_eb.PdfObjectInteger
	BitsPerFlag       *_eb.PdfObjectInteger
	Decode            *_eb.PdfObjectArray
	Function          []PdfFunction
}

func _fgeg() string {
	_bddg := "\u0051\u0057\u0045\u0052\u0054\u0059\u0055\u0049\u004f\u0050\u0041S\u0044\u0046\u0047\u0048\u004a\u004b\u004c\u005a\u0058\u0043V\u0042\u004e\u004d"
	var _face _dd.Buffer
	for _gffef := 0; _gffef < 6; _gffef++ {
		_face.WriteRune(rune(_bddg[_gaa.Intn(len(_bddg))]))
	}
	return _face.String()
}

// GetRuneMetrics iterates through each font in the list of fonts the returns the fonts.CharMetrics from working font.
func (_ffddd *MultipleFontEncoder) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	_ecccf := _ffddd.CurrentFont
	_gcaec, _ddfc := _ecccf.GetRuneMetrics(r)
	for _bdafb := 1; _bdafb < len(_ffddd._fcgbc) && _gcaec.Wx == 0; _bdafb++ {
		_ecccf = _ffddd._fcgbc[_bdafb]
		_gcaec, _ddfc = _ecccf.GetRuneMetrics(r)
	}
	return _gcaec, _ddfc
}

// GetContainingPdfObject returns the container of the image object (indirect object).
func (_bgcad *XObjectImage) GetContainingPdfObject() _eb.PdfObject { return _bgcad._gceffg }

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_accgf pdfCIDFontType2) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	_bacb, _fefg := _accgf._caba[r]
	if !_fefg {
		_babgb, _fbcbf := _eb.GetInt(_accgf.DW)
		if !_fbcbf {
			return _fg.CharMetrics{}, false
		}
		_bacb = int(*_babgb)
	}
	return _fg.CharMetrics{Wx: float64(_bacb)}, true
}

// ToPdfObject returns the text field dictionary within an indirect object (container).
func (_fage *PdfFieldText) ToPdfObject() _eb.PdfObject {
	_fage.PdfField.ToPdfObject()
	_aded := _fage._adgda
	_dfdd := _aded.PdfObject.(*_eb.PdfObjectDictionary)
	_dfdd.Set("\u0046\u0054", _eb.MakeName("\u0054\u0078"))
	if _fage.DA != nil {
		_dfdd.Set("\u0044\u0041", _fage.DA)
	}
	if _fage.Q != nil {
		_dfdd.Set("\u0051", _fage.Q)
	}
	if _fage.DS != nil {
		_dfdd.Set("\u0044\u0053", _fage.DS)
	}
	if _fage.RV != nil {
		_dfdd.Set("\u0052\u0056", _fage.RV)
	}
	if _fage.MaxLen != nil {
		_dfdd.Set("\u004d\u0061\u0078\u004c\u0065\u006e", _fage.MaxLen)
	}
	return _aded
}
func _dgfc(_aaef _eb.PdfObject) (*PdfColorspaceSpecialIndexed, error) {
	_gbag := NewPdfColorspaceSpecialIndexed()
	if _afca, _dcfc := _aaef.(*_eb.PdfIndirectObject); _dcfc {
		_gbag._daac = _afca
	}
	_aaef = _eb.TraceToDirectObject(_aaef)
	_accg, _ebgfd := _aaef.(*_eb.PdfObjectArray)
	if !_ebgfd {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _accg.Len() != 4 {
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_aaef = _accg.Get(0)
	_becag, _ebgfd := _aaef.(*_eb.PdfObjectName)
	if !_ebgfd {
		return nil, _e.Errorf("\u0069n\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
	}
	if *_becag != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065xe\u0064\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u006e\u0061m\u0065")
	}
	_aaef = _accg.Get(1)
	_dfccd, _baaa := DetermineColorspaceNameFromPdfObject(_aaef)
	if _baaa != nil {
		return nil, _baaa
	}
	if _dfccd == "\u0049n\u0064\u0065\u0078\u0065\u0064" || _dfccd == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
		_ddb.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0049\u006e\u0064e\u0078\u0065\u0064\u002f\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043S\u0020\u0061\u0073\u0020\u0062\u0061\u0073\u0065\u0020\u0028\u0025v\u0029", _dfccd)
		return nil, _eggcg
	}
	_eefe, _baaa := NewPdfColorspaceFromPdfObject(_aaef)
	if _baaa != nil {
		return nil, _baaa
	}
	_gbag.Base = _eefe
	_aaef = _accg.Get(2)
	_aecbe, _baaa := _eb.GetNumberAsInt64(_aaef)
	if _baaa != nil {
		return nil, _baaa
	}
	if _aecbe > 255 {
		return nil, _e.Errorf("\u0069n\u0064\u0065\u0078\u0065d\u0020\u0043\u0053\u003a\u0020I\u006ev\u0061l\u0069\u0064\u0020\u0068\u0069\u0076\u0061l")
	}
	_gbag.HiVal = int(_aecbe)
	_aaef = _accg.Get(3)
	_gbag.Lookup = _aaef
	_aaef = _eb.TraceToDirectObject(_aaef)
	var _ccgd []byte
	if _efgca, _beaf := _aaef.(*_eb.PdfObjectString); _beaf {
		_ccgd = _efgca.Bytes()
		_ddb.Log.Trace("\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0020\u0064", _ccgd)
	} else if _bbaa, _fddb := _aaef.(*_eb.PdfObjectStream); _fddb {
		_ddb.Log.Trace("\u0049n\u0064e\u0078\u0065\u0064\u0020\u0073t\u0072\u0065a\u006d\u003a\u0020\u0025\u0073", _aaef.String())
		_ddb.Log.Trace("\u0045\u006e\u0063\u006fde\u0064\u0020\u0028\u0025\u0064\u0029\u0020\u003a\u0020\u0025\u0023\u0020\u0078", len(_bbaa.Stream), _bbaa.Stream)
		_fbfdg, _bggg := _eb.DecodeStream(_bbaa)
		if _bggg != nil {
			return nil, _bggg
		}
		_ddb.Log.Trace("\u0044e\u0063o\u0064\u0065\u0064\u0020\u0028%\u0064\u0029 \u003a\u0020\u0025\u0020\u0058", len(_fbfdg), _fbfdg)
		_ccgd = _fbfdg
	} else {
		_ddb.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _aaef)
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076a\u006ci\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if len(_ccgd) < _gbag.Base.GetNumComponents()*(_gbag.HiVal+1) {
		_ddb.Log.Debug("\u0050\u0044\u0046\u0020\u0049\u006e\u0063o\u006d\u0070\u0061t\u0069\u0062\u0069\u006ci\u0074\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074\u006f\u006f\u0020\u0073\u0068\u006f\u0072\u0074")
		_ddb.Log.Debug("\u0046\u0061i\u006c\u002c\u0020\u006c\u0065\u006e\u0028\u0064\u0061\u0074\u0061\u0029\u003a\u0020\u0025\u0064\u002c\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u002c\u0020\u0068\u0069\u0056\u0061\u006c\u003a\u0020\u0025\u0064", len(_ccgd), _gbag.Base.GetNumComponents(), _gbag.HiVal)
	} else {
		_ccgd = _ccgd[:_gbag.Base.GetNumComponents()*(_gbag.HiVal+1)]
	}
	_gbag._efcb = _ccgd
	return _gbag, nil
}

// ColorToRGB converts a CMYK32 color to an RGB color.
func (_bcbb *PdfColorspaceDeviceCMYK) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ggfa, _dgef := color.(*PdfColorDeviceCMYK)
	if !_dgef {
		_ddb.Log.Debug("I\u006e\u0070\u0075\u0074\u0020\u0063o\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0064e\u0076\u0069\u0063e\u0020c\u006d\u0079\u006b")
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bbdfc := _ggfa.C()
	_fccc := _ggfa.M()
	_feeg := _ggfa.Y()
	_ceec := _ggfa.K()
	_bbdfc = _bbdfc*(1-_ceec) + _ceec
	_fccc = _fccc*(1-_ceec) + _ceec
	_feeg = _feeg*(1-_ceec) + _ceec
	_dgca := 1 - _bbdfc
	_ddag := 1 - _fccc
	_faa := 1 - _feeg
	return NewPdfColorDeviceRGB(_dgca, _ddag, _faa), nil
}

// Flags returns the field flags for the field accounting for any inherited flags.
func (_ededc *PdfField) Flags() FieldFlag {
	var _ecfa FieldFlag
	_ggffd, _bffe := _ededc.inherit(func(_ecbeg *PdfField) bool {
		if _ecbeg.Ff != nil {
			_ecfa = FieldFlag(*_ecbeg.Ff)
			return true
		}
		return false
	})
	if _bffe != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0069\u006e\u0067\u0020\u0066\u006c\u0061\u0067\u0073\u0020\u0076\u0069\u0061\u0020\u0069\u006e\u0068\u0065\u0072\u0069t\u0061\u006e\u0063\u0065\u003a\u0020\u0025\u0076", _bffe)
	}
	if !_ggffd {
		_ddb.Log.Trace("N\u006f\u0020\u0066\u0069\u0065\u006cd\u0020\u0066\u006c\u0061\u0067\u0073 \u0066\u006f\u0075\u006e\u0064\u0020\u002d \u0061\u0073\u0073\u0075\u006d\u0065\u0020\u0063\u006c\u0065a\u0072")
	}
	return _ecfa
}

// ToPdfObject converts the K dictionary to a PDF object.
func (_eefed *KDict) ToPdfObject() _eb.PdfObject {
	_faddg := _eb.MakeDict()
	if _eefed.ID != nil {
		_faddg.Set("\u0049\u0044", _eefed.ID)
	}
	if _eefed.K != nil {
		_faddg.Set("\u004b", _eefed.K)
	} else if len(_eefed._eegge) > 0 {
		if len(_eefed._eegge) == 1 {
			_eefed.K = _eefed._eegge[0].ToPdfObject()
		} else {
			_adbfb := _eb.MakeArray()
			for _, _gaed := range _eefed._eegge {
				_adbfb.Append(_gaed.ToPdfObject())
			}
			_eefed.K = _adbfb
		}
		_faddg.Set("\u004b", _eefed.K)
	}
	if _eefed.S != nil {
		_faddg.Set("\u0053", _eefed.S)
	}
	if _eefed.P != nil {
		_faddg.Set("\u0050", _eefed.P)
	}
	if _eefed.Pg != nil {
		_faddg.Set("\u0050\u0067", _eefed.Pg)
	}
	if _eefed.C != nil {
		_faddg.Set("\u0043", _eefed.C)
	}
	if _eefed.R != nil {
		_faddg.Set("\u0052", _eefed.R)
	}
	if _eefed.T != nil {
		_faddg.Set("\u0054", _eefed.T)
	}
	if _eefed.Lang != nil {
		_faddg.Set("\u004c\u0061\u006e\u0067", _eefed.Lang)
	}
	if _eefed.Alt != nil {
		_faddg.Set("\u0041\u006c\u0074", _eefed.Alt)
	}
	if _eefed.E != nil {
		_faddg.Set("\u0045", _eefed.E)
	}
	if _eefed.A != nil {
		_faddg.Set("\u0041", _eefed.A)
	} else if _eefed._aagfc != nil {
		_fafaf := _eb.MakeArrayFromFloats([]float64{_eefed._aagfc.Llx, _eefed._aagfc.Lly, _eefed._aagfc.Urx, _eefed._aagfc.Ury})
		_babdcf := _eb.MakeDict()
		_babdcf.Set("\u0042\u0042\u006f\u0078", _fafaf)
		_babdcf.Set("\u004f", _eb.MakeString("\u004c\u0061\u0079\u006f\u0075\u0074"))
		_faddg.Set("\u0041", _eb.MakeIndirectObject(_babdcf))
	}
	return _faddg
}
func _cffg(_gddb *PdfAnnotation) (*XObjectForm, *PdfRectangle, error) {
	_gdbda, _deeca := _eb.GetDict(_gddb.AP)
	if !_deeca {
		return nil, nil, _dcf.New("f\u0069\u0065\u006c\u0064\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0041\u0050\u0020d\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079")
	}
	if _gdbda == nil {
		return nil, nil, nil
	}
	_gbcf, _deeca := _eb.GetArray(_gddb.Rect)
	if !_deeca || _gbcf.Len() != 4 {
		return nil, nil, _dcf.New("\u0072\u0065\u0063t\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_ffbf, _adbg := NewPdfRectangle(*_gbcf)
	if _adbg != nil {
		return nil, nil, _adbg
	}
	_gfcd := _eb.TraceToDirectObject(_gdbda.Get("\u004e"))
	switch _cdfc := _gfcd.(type) {
	case *_eb.PdfObjectStream:
		_fabgd := _cdfc
		_gbacd, _fgede := NewXObjectFormFromStream(_fabgd)
		return _gbacd, _ffbf, _fgede
	case *_eb.PdfObjectDictionary:
		_gccaa := _cdfc
		if _gccaa == nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067e\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0061\u006e\u0063\u0065\u002e\u0020\u0044\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
			return nil, nil, nil
		}
		var _ccdc _eb.PdfObject
		_cgad, _dcba := _eb.GetName(_gddb.AS)
		if _dcba {
			_ccdc = _gccaa.Get(*_cgad)
		} else {
			_caca := _gddb._ggf.PdfObject.(*_eb.PdfObjectDictionary)
			if _caca == nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u0067\u0065\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u0076\u0061\u006c\u0075\u0065\u002e")
				return nil, nil, nil
			}
			if _gded := _caca.Get("\u0056"); _gded != nil {
				_ccdc = _gccaa.Get(_eb.PdfObjectName(_gded.String()))
			}
		}
		if _ccdc == nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0041\u0053\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
			return nil, nil, nil
		}
		_bacec, _dcba := _eb.GetStream(_ccdc)
		if !_dcba {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006ce \u0074\u006f\u0020\u0061\u0063\u0063e\u0073\u0073\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0066\u006f\u0072 \u0025\u0076", _cgad)
			return nil, nil, _dcf.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		}
		_agag, _cdfag := NewXObjectFormFromStream(_bacec)
		return _agag, _ffbf, _cdfag
	}
	_ddb.Log.Debug("\u0049\u006e\u0076\u0061li\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u004e\u003a\u0020%\u0054", _gfcd)
	return nil, nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
}

// SetXObjectByName adds the XObject from the passed in stream to the page resources.
// The added XObject is identified by the specified name.
func (_gdeca *PdfPageResources) SetXObjectByName(keyName _eb.PdfObjectName, stream *_eb.PdfObjectStream) error {
	if _gdeca.XObject == nil {
		_gdeca.XObject = _eb.MakeDict()
	}
	_ebbca := _eb.TraceToDirectObject(_gdeca.XObject)
	_geed, _ebbg := _ebbca.(*_eb.PdfObjectDictionary)
	if !_ebbg {
		_ddb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0058\u004f\u0062j\u0065\u0063\u0074\u002c\u0020\u0067\u006f\u0074\u0020\u0025T\u002f\u0025\u0054", _gdeca.XObject, _ebbca)
		return _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_geed.Set(keyName, stream)
	return nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// L (range 0-100), A (range -100-100) and B (range -100-100) components of
// the color.
func (_gfef *PdfColorspaceLab) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aegdb := vals[0]
	if _aegdb < 0.0 || _aegdb > 100.0 {
		_ddb.Log.Debug("\u004c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030-\u0031\u0030\u0030\u0029", _aegdb)
		return nil, ErrColorOutOfRange
	}
	_bfga := vals[1]
	_fdgb := float64(-100)
	_fbbg := float64(100)
	if len(_gfef.Range) > 1 {
		_fdgb = _gfef.Range[0]
		_fbbg = _gfef.Range[1]
	}
	if _bfga < _fdgb || _bfga > _fbbg {
		_ddb.Log.Debug("\u0041\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _bfga, _fdgb, _fbbg)
		return nil, ErrColorOutOfRange
	}
	_bbfb := vals[2]
	_fbcd := float64(-100)
	_cfgcf := float64(100)
	if len(_gfef.Range) > 3 {
		_fbcd = _gfef.Range[2]
		_cfgcf = _gfef.Range[3]
	}
	if _bbfb < _fbcd || _bbfb > _cfgcf {
		_ddb.Log.Debug("\u0062\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _bbfb, _fbcd, _cfgcf)
		return nil, ErrColorOutOfRange
	}
	_adgbd := NewPdfColorLab(_aegdb, _bfga, _bbfb)
	return _adgbd, nil
}

type crossReference struct {
	Type int

	// Type 1
	Offset     int64
	Generation int64

	// Type 2
	ObjectNumber int
	Index        int
}

// HasXObjectByName checks if an XObject with a specified keyName is defined.
func (_efead *PdfPageResources) HasXObjectByName(keyName _eb.PdfObjectName) bool {
	_dgfg, _ := _efead.GetXObjectByName(keyName)
	return _dgfg != nil
}

// GetOutlinesFlattened returns a flattened list of tree nodes and titles.
// NOTE: for most use cases, it is recommended to use the high-level GetOutlines
// method instead, which also provides information regarding the destination
// of the outline items.
func (_dddf *PdfReader) GetOutlinesFlattened() ([]*PdfOutlineTreeNode, []string, error) {
	var _bcdda []*PdfOutlineTreeNode
	var _bgfgb []string
	var _cdcfb func(*PdfOutlineTreeNode, *[]*PdfOutlineTreeNode, *[]string, int)
	_cdcfb = func(_caae *PdfOutlineTreeNode, _gaceff *[]*PdfOutlineTreeNode, _bbceg *[]string, _cede int) {
		if _caae == nil {
			return
		}
		if _caae._eeedb == nil {
			_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020M\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006e\u006fd\u0065\u002e\u0063o\u006et\u0065\u0078\u0074")
			return
		}
		_gefd, _fbac := _caae._eeedb.(*PdfOutlineItem)
		if _fbac {
			*_gaceff = append(*_gaceff, &_gefd.PdfOutlineTreeNode)
			_dedc := _cc.Repeat("\u0020", _cede*2) + _gefd.Title.Decoded()
			*_bbceg = append(*_bbceg, _dedc)
		}
		if _caae.First != nil {
			_fdda := _cc.Repeat("\u0020", _cede*2) + "\u002b"
			*_bbceg = append(*_bbceg, _fdda)
			_cdcfb(_caae.First, _gaceff, _bbceg, _cede+1)
		}
		if _fbac && _gefd.Next != nil {
			_cdcfb(_gefd.Next, _gaceff, _bbceg, _cede)
		}
	}
	_cdcfb(_dddf._efabg, &_bcdda, &_bgfgb, 0)
	return _bcdda, _bgfgb, nil
}

// Fill populates `form` with values provided by `provider`.
func (_efbf *PdfAcroForm) Fill(provider FieldValueProvider) error { return _efbf.fill(provider, nil) }

// GetStructRoot gets the StructTreeRoot object
func (_fafbd *PdfPage) GetStructTreeRoot() (*_eb.PdfObject, bool) {
	_cffcf, _ddfge := _fafbd._dcdfd.GetCatalogStructTreeRoot()
	return &_cffcf, _ddfge
}
func _eadgga(_ggca *_eb.PdfObjectArray) (float64, error) {
	_dcbae, _gcadg := _ggca.ToFloat64Array()
	if _gcadg != nil {
		_ddb.Log.Debug("\u0042\u0061\u0064\u0020\u004d\u0061\u0074\u0074\u0065\u0020\u0061\u0072\u0072\u0061\u0079:\u0020m\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ggca, _gcadg)
	}
	switch len(_dcbae) {
	case 1:
		return _dcbae[0], nil
	case 3:
		_cgbc := PdfColorspaceDeviceRGB{}
		_fgebedc, _ggfdd := _cgbc.ColorFromFloats(_dcbae)
		if _ggfdd != nil {
			return 0.0, _ggfdd
		}
		return _fgebedc.(*PdfColorDeviceRGB).ToGray().Val(), nil
	case 4:
		_abfab := PdfColorspaceDeviceCMYK{}
		_faedg, _eabbcd := _abfab.ColorFromFloats(_dcbae)
		if _eabbcd != nil {
			return 0.0, _eabbcd
		}
		_daae, _eabbcd := _abfab.ColorToRGB(_faedg.(*PdfColorDeviceCMYK))
		if _eabbcd != nil {
			return 0.0, _eabbcd
		}
		return _daae.(*PdfColorDeviceRGB).ToGray().Val(), nil
	}
	_gcadg = _dcf.New("\u0062a\u0064 \u004d\u0061\u0074\u0074\u0065\u0020\u0063\u006f\u006c\u006f\u0072")
	_ddb.Log.Error("\u0074\u006f\u0047ra\u0079\u003a\u0020\u006d\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ggca, _gcadg)
	return 0.0, _gcadg
}

// A PdfPattern can represent a Pattern, either a tiling pattern or a shading pattern.
// Note that all patterns shall be treated as colours; a Pattern colour space shall be established with the CS or cs
// operator just like other colour spaces, and a particular pattern shall be installed as the current colour with the
// SCN or scn operator.
type PdfPattern struct {

	// Type: Pattern
	PatternType int64
	_eefgb      PdfModel
	_agddd      _eb.PdfObject
}

func _agdf(_adaf *_eb.PdfObjectDictionary) (*PdfFieldChoice, error) {
	_dcea := &PdfFieldChoice{}
	_dcea.Opt, _ = _eb.GetArray(_adaf.Get("\u004f\u0070\u0074"))
	_dcea.TI, _ = _eb.GetInt(_adaf.Get("\u0054\u0049"))
	_dcea.I, _ = _eb.GetArray(_adaf.Get("\u0049"))
	return _dcea, nil
}

// GetCustomInfo returns a custom info value for the specified name.
func (_dcge *PdfInfo) GetCustomInfo(name string) *_eb.PdfObjectString {
	var _afae *_eb.PdfObjectString
	if _dcge._cbfb == nil {
		return _afae
	}
	if _abfg, _eedf := _dcge._cbfb.Get(*_eb.MakeName(name)).(*_eb.PdfObjectString); _eedf {
		_afae = _abfg
	}
	return _afae
}

// GetNumComponents returns the number of color components (1 for Separation).
func (_ecba *PdfColorspaceSpecialSeparation) GetNumComponents() int { return 1 }

// ToPdfObject returns the PDF representation of the colorspace.
func (_cgff *PdfColorspaceDeviceRGB) ToPdfObject() _eb.PdfObject {
	return _eb.MakeName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_dedd *PdfColorspaceDeviceGray) ToPdfObject() _eb.PdfObject {
	return _eb.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079")
}
func (_cegb *PdfColorspaceDeviceGray) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
}

// NewCustomPdfOutputIntent creates a new custom PdfOutputIntent.
func NewCustomPdfOutputIntent(outputCondition, outputConditionIdentifier, info string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, Info: info, DestOutputProfile: destOutputProfile, _ebbd: _eb.MakeDict(), ColorComponents: colorComponents}
}
func _ddbeg(_bbefb *XObjectImage) error {
	if _bbefb.SMask == nil {
		return nil
	}
	_gbbcg, _cgcf := _bbefb.SMask.(*_eb.PdfObjectStream)
	if !_cgcf {
		_ddb.Log.Debug("\u0053\u004da\u0073\u006b\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0053\u0074re\u0061\u006d")
		return _eb.ErrTypeError
	}
	_cgfe := _gbbcg.PdfObjectDictionary
	_eaafe := _cgfe.Get("\u004d\u0061\u0074t\u0065")
	if _eaafe == nil {
		return nil
	}
	_gffba, _bgecg := _eadgga(_eaafe.(*_eb.PdfObjectArray))
	if _bgecg != nil {
		return _bgecg
	}
	_ffdaa := _eb.MakeArrayFromFloats([]float64{_gffba})
	_cgfe.SetIfNotNil("\u004d\u0061\u0074t\u0065", _ffdaa)
	return nil
}

// IsCenterWindow returns the value of the centerWindow flag.
func (_aebbc *ViewerPreferences) IsCenterWindow() bool {
	if _aebbc._dacc == nil {
		return false
	}
	return *_aebbc._dacc
}

// PdfActionThread represents a thread action.
type PdfActionThread struct {
	*PdfAction
	F *PdfFilespec
	D _eb.PdfObject
	B _eb.PdfObject
}

// ToPdfObject returns the PDF representation of the function.
func (_ecfba *PdfFunctionType3) ToPdfObject() _eb.PdfObject {
	_bbgfb := _eb.MakeDict()
	_bbgfb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _eb.MakeInteger(3))
	_eegaa := &_eb.PdfObjectArray{}
	for _, _bcad := range _ecfba.Domain {
		_eegaa.Append(_eb.MakeFloat(_bcad))
	}
	_bbgfb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _eegaa)
	if _ecfba.Range != nil {
		_fgbca := &_eb.PdfObjectArray{}
		for _, _bacdc := range _ecfba.Range {
			_fgbca.Append(_eb.MakeFloat(_bacdc))
		}
		_bbgfb.Set("\u0052\u0061\u006eg\u0065", _fgbca)
	}
	if _ecfba.Functions != nil {
		_gfedc := &_eb.PdfObjectArray{}
		for _, _eefcaa := range _ecfba.Functions {
			_gfedc.Append(_eefcaa.ToPdfObject())
		}
		_bbgfb.Set("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es", _gfedc)
	}
	if _ecfba.Bounds != nil {
		_gdgb := &_eb.PdfObjectArray{}
		for _, _dbeb := range _ecfba.Bounds {
			_gdgb.Append(_eb.MakeFloat(_dbeb))
		}
		_bbgfb.Set("\u0042\u006f\u0075\u006e\u0064\u0073", _gdgb)
	}
	if _ecfba.Encode != nil {
		_gcfed := &_eb.PdfObjectArray{}
		for _, _gacgf := range _ecfba.Encode {
			_gcfed.Append(_eb.MakeFloat(_gacgf))
		}
		_bbgfb.Set("\u0045\u006e\u0063\u006f\u0064\u0065", _gcfed)
	}
	if _ecfba._ddgdg != nil {
		_ecfba._ddgdg.PdfObject = _bbgfb
		return _ecfba._ddgdg
	}
	return _bbgfb
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_bcbf pdfFontType3) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	_fcda := _bcbf.Encoder()
	if _fcda == nil {
		_ddb.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _bcbf)
		return _fg.CharMetrics{}, false
	}
	_cecd, _cdggf := _fcda.RuneToCharcode(r)
	if !_cdggf {
		if r != ' ' {
			_ddb.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _bcbf)
		}
		return _fg.CharMetrics{}, false
	}
	_badc, _gcef := _bcbf.GetCharMetrics(_cecd)
	return _badc, _gcef
}

// SignatureValidationResult defines the response from the signature validation handler.
type SignatureValidationResult struct {

	// List of errors when validating the signature.
	Errors      []string
	IsSigned    bool
	IsVerified  bool
	IsTrusted   bool
	Fields      []*PdfField
	Name        string
	Date        PdfDate
	Reason      string
	Location    string
	ContactInfo string
	DiffResults *_bab.DiffResults
	IsCrlFound  bool
	IsOcspFound bool

	// GeneralizedTime is the time at which the time-stamp token has been created by the TSA (RFC 3161).
	GeneralizedTime _d.Time
}

// SetName sets the `Name` field of the signature.
func (_gdccd *PdfSignature) SetName(name string) { _gdccd.Name = _eb.MakeEncodedString(name, true) }
func (_babag *PdfReader) newPdfActionSoundFromDict(_gff *_eb.PdfObjectDictionary) (*PdfActionSound, error) {
	return &PdfActionSound{Sound: _gff.Get("\u0053\u006f\u0075n\u0064"), Volume: _gff.Get("\u0056\u006f\u006c\u0075\u006d\u0065"), Synchronous: _gff.Get("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073"), Repeat: _gff.Get("\u0052\u0065\u0070\u0065\u0061\u0074"), Mix: _gff.Get("\u004d\u0069\u0078")}, nil
}
func (_agdbd *PdfColorspaceICCBased) String() string {
	return "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"
}

// PdfActionURI represents an URI action.
type PdfActionURI struct {
	*PdfAction
	URI   _eb.PdfObject
	IsMap _eb.PdfObject
}

// SetPdfProducer sets the Producer attribute of the output PDF.
func SetPdfProducer(producer string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _ccgde = producer }

// PdfFontDescriptor specifies metrics and other attributes of a font and can refer to a FontFile
// for embedded fonts.
// 9.8 Font Descriptors (page 281)
type PdfFontDescriptor struct {
	FontName     _eb.PdfObject
	FontFamily   _eb.PdfObject
	FontStretch  _eb.PdfObject
	FontWeight   _eb.PdfObject
	Flags        _eb.PdfObject
	FontBBox     _eb.PdfObject
	ItalicAngle  _eb.PdfObject
	Ascent       _eb.PdfObject
	Descent      _eb.PdfObject
	Leading      _eb.PdfObject
	CapHeight    _eb.PdfObject
	XHeight      _eb.PdfObject
	StemV        _eb.PdfObject
	StemH        _eb.PdfObject
	AvgWidth     _eb.PdfObject
	MaxWidth     _eb.PdfObject
	MissingWidth _eb.PdfObject
	FontFile     _eb.PdfObject
	FontFile2    _eb.PdfObject
	FontFile3    _eb.PdfObject
	CharSet      _eb.PdfObject
	_eeda        int
	_acaad       float64
	*fontFile
	_bebgd *_fg.TtfType

	// Additional entries for CIDFonts
	Style  _eb.PdfObject
	Lang   _eb.PdfObject
	FD     _eb.PdfObject
	CIDSet _eb.PdfObject
	_fdea  *_eb.PdfIndirectObject
}

// Encoder iterates through the list of fonts and returns a working encoder
func (_abdbe *MultipleFontEncoder) Encoder(rn rune) (_fc.TextEncoder, bool) {
	_gcgac := _abdbe.CurrentFont
	_cbgbf := _gcgac.Encoder()
	_, _ddede := _cbgbf.RuneToCharcode(rn)
	for _affge := 1; _affge < len(_abdbe._fcgbc) && !_ddede; _affge++ {
		_gcgac = _abdbe._fcgbc[_affge]
		_cbgbf = _gcgac.Encoder()
		_, _ddede = _cbgbf.RuneToCharcode(rn)
		_abdbe.CurrentFont = _gcgac
	}
	return _cbgbf, _ddede
}

// SetHideWindowUI sets the value of the hideWindowUI flag.
func (_fcdac *ViewerPreferences) SetHideWindowUI(hideWindowUI bool) { _fcdac._fcebae = &hideWindowUI }

// PdfFieldSignature signature field represents digital signatures and optional data for authenticating
// the name of the signer and verifying document contents.
type PdfFieldSignature struct {
	*PdfField
	*PdfAnnotationWidget
	V    *PdfSignature
	Lock *_eb.PdfIndirectObject
	SV   *_eb.PdfIndirectObject
}

// SetPdfKeywords sets the Keywords attribute of the output PDF.
func SetPdfKeywords(keywords string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _addbb = keywords }
func (_fgdde *PdfReader) loadAnnotations(_fbbee _eb.PdfObject) ([]*PdfAnnotation, error) {
	_ebdd, _eaedbe := _eb.GetArray(_fbbee)
	if !_eaedbe {
		return nil, _e.Errorf("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	var _faafa []*PdfAnnotation
	for _, _dfefd := range _ebdd.Elements() {
		_dfefd = _eb.ResolveReference(_dfefd)
		if _, _gdfcf := _dfefd.(*_eb.PdfObjectNull); _gdfcf {
			continue
		}
		_aaee, _decbd := _dfefd.(*_eb.PdfObjectDictionary)
		_eaggaf, _bege := _dfefd.(*_eb.PdfIndirectObject)
		if _decbd {
			_eaggaf = &_eb.PdfIndirectObject{}
			_eaggaf.PdfObject = _aaee
		} else {
			if !_bege {
				return nil, _e.Errorf("\u0061\u006eno\u0074\u0061\u0074i\u006f\u006e\u0020\u006eot \u0069n \u0061\u006e\u0020\u0069\u006e\u0064\u0069re\u0063\u0074\u0020\u006f\u0062\u006a\u0065c\u0074")
			}
		}
		_gfeag, _acgbe := _fgdde.newPdfAnnotationFromIndirectObject(_eaggaf)
		if _acgbe != nil {
			return nil, _acgbe
		}
		switch _bcae := _gfeag.GetContext().(type) {
		case *PdfAnnotationWidget:
			for _, _dfagd := range _fgdde.AcroForm.AllFields() {
				if _dfagd._adgda == _bcae.Parent {
					_bcae._bca = _dfagd
					break
				}
			}
		}
		if _gfeag != nil {
			_faafa = append(_faafa, _gfeag)
		}
	}
	return _faafa, nil
}

// IsDisplayDocTitle returns the value of the displayDocTitle flag.
func (_abgef *ViewerPreferences) IsDisplayDocTitle() bool {
	if _abgef._bgdge == nil {
		return false
	}
	return *_abgef._bgdge
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_agdb *PdfColorspaceLab) ToPdfObject() _eb.PdfObject {
	_fgcb := _eb.MakeArray()
	_fgcb.Append(_eb.MakeName("\u004c\u0061\u0062"))
	_ccac := _eb.MakeDict()
	if _agdb.WhitePoint != nil {
		_gdcfb := _eb.MakeArray(_eb.MakeFloat(_agdb.WhitePoint[0]), _eb.MakeFloat(_agdb.WhitePoint[1]), _eb.MakeFloat(_agdb.WhitePoint[2]))
		_ccac.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _gdcfb)
	} else {
		_ddb.Log.Error("\u004c\u0061\u0062: \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057h\u0069t\u0065P\u006fi\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	if _agdb.BlackPoint != nil {
		_fgdb := _eb.MakeArray(_eb.MakeFloat(_agdb.BlackPoint[0]), _eb.MakeFloat(_agdb.BlackPoint[1]), _eb.MakeFloat(_agdb.BlackPoint[2]))
		_ccac.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _fgdb)
	}
	if _agdb.Range != nil {
		_aedba := _eb.MakeArray(_eb.MakeFloat(_agdb.Range[0]), _eb.MakeFloat(_agdb.Range[1]), _eb.MakeFloat(_agdb.Range[2]), _eb.MakeFloat(_agdb.Range[3]))
		_ccac.Set("\u0052\u0061\u006eg\u0065", _aedba)
	}
	_fgcb.Append(_ccac)
	if _agdb._cbag != nil {
		_agdb._cbag.PdfObject = _fgcb
		return _agdb._cbag
	}
	return _fgcb
}
func _facfg(_gegce *_eb.PdfObjectDictionary) (*PdfShadingType7, error) {
	_febdg := PdfShadingType7{}
	_cafef := _gegce.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _cafef == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dbbbb, _gbbebb := _cafef.(*_eb.PdfObjectInteger)
	if !_gbbebb {
		_ddb.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cafef)
		return nil, _eb.ErrTypeError
	}
	_febdg.BitsPerCoordinate = _dbbbb
	_cafef = _gegce.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _cafef == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_dbbbb, _gbbebb = _cafef.(*_eb.PdfObjectInteger)
	if !_gbbebb {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _cafef)
		return nil, _eb.ErrTypeError
	}
	_febdg.BitsPerComponent = _dbbbb
	_cafef = _gegce.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _cafef == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dbbbb, _gbbebb = _cafef.(*_eb.PdfObjectInteger)
	if !_gbbebb {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _cafef)
		return nil, _eb.ErrTypeError
	}
	_febdg.BitsPerComponent = _dbbbb
	_cafef = _gegce.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _cafef == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_abaae, _gbbebb := _cafef.(*_eb.PdfObjectArray)
	if !_gbbebb {
		_ddb.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cafef)
		return nil, _eb.ErrTypeError
	}
	_febdg.Decode = _abaae
	if _bgcbc := _gegce.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _bgcbc != nil {
		_febdg.Function = []PdfFunction{}
		if _gcabg, _fdcba := _bgcbc.(*_eb.PdfObjectArray); _fdcba {
			for _, _ebffc := range _gcabg.Elements() {
				_cccc, _faac := _cccfa(_ebffc)
				if _faac != nil {
					_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _faac)
					return nil, _faac
				}
				_febdg.Function = append(_febdg.Function, _cccc)
			}
		} else {
			_fbbcg, _cdcfe := _cccfa(_bgcbc)
			if _cdcfe != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cdcfe)
				return nil, _cdcfe
			}
			_febdg.Function = append(_febdg.Function, _fbbcg)
		}
	}
	return &_febdg, nil
}
func (_cffged *PdfWriter) copyObject(_adeea _eb.PdfObject, _bdgbc map[_eb.PdfObject]_eb.PdfObject, _cfefe map[_eb.PdfObject]struct{}, _dcfgc bool) _eb.PdfObject {
	_afdf := !_cffged._caed && _cfefe != nil
	if _fbdbdg, _cddd := _bdgbc[_adeea]; _cddd {
		if _afdf && !_dcfgc {
			delete(_cfefe, _adeea)
		}
		return _fbdbdg
	}
	if _adeea == nil {
		_fdbfa := _eb.MakeNull()
		return _fdbfa
	}
	_cccef := _adeea
	switch _cffbc := _adeea.(type) {
	case *_eb.PdfObjectArray:
		_gacag := _eb.MakeArray()
		_cccef = _gacag
		_bdgbc[_adeea] = _cccef
		for _, _gbbea := range _cffbc.Elements() {
			_gacag.Append(_cffged.copyObject(_gbbea, _bdgbc, _cfefe, _dcfgc))
		}
	case *_eb.PdfObjectStreams:
		_fgbae := &_eb.PdfObjectStreams{PdfObjectReference: _cffbc.PdfObjectReference}
		_cccef = _fgbae
		_bdgbc[_adeea] = _cccef
		for _, _cegf := range _cffbc.Elements() {
			_fgbae.Append(_cffged.copyObject(_cegf, _bdgbc, _cfefe, _dcfgc))
		}
	case *_eb.PdfObjectStream:
		_dbeeca := &_eb.PdfObjectStream{Stream: _cffbc.Stream, PdfObjectReference: _cffbc.PdfObjectReference, Lazy: _cffbc.Lazy, TempFile: _cffbc.TempFile}
		_cccef = _dbeeca
		_bdgbc[_adeea] = _cccef
		_dbeeca.PdfObjectDictionary = _cffged.copyObject(_cffbc.PdfObjectDictionary, _bdgbc, _cfefe, _dcfgc).(*_eb.PdfObjectDictionary)
	case *_eb.PdfObjectDictionary:
		var _dfbbb bool
		if _afdf && !_dcfgc {
			if _acaga, _ := _eb.GetNameVal(_cffbc.Get("\u0054\u0079\u0070\u0065")); _acaga == "\u0050\u0061\u0067\u0065" {
				_, _ffcec := _cffged._gbgc[_cffbc]
				_dcfgc = !_ffcec
				_dfbbb = _dcfgc
			}
		}
		_dggf := _eb.MakeDict()
		_cccef = _dggf
		_bdgbc[_adeea] = _cccef
		for _, _gbbab := range _cffbc.Keys() {
			_dggf.Set(_gbbab, _cffged.copyObject(_cffbc.Get(_gbbab), _bdgbc, _cfefe, _dcfgc))
		}
		if _dfbbb {
			_cccef = _eb.MakeNull()
			_dcfgc = false
		}
	case *_eb.PdfIndirectObject:
		_aedccc := &_eb.PdfIndirectObject{PdfObjectReference: _cffbc.PdfObjectReference}
		_cccef = _aedccc
		_bdgbc[_adeea] = _cccef
		_aedccc.PdfObject = _cffged.copyObject(_cffbc.PdfObject, _bdgbc, _cfefe, _dcfgc)
	case *_eb.PdfObjectString:
		_bdgfc := *_cffbc
		_cccef = &_bdgfc
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectName:
		_eagae := *_cffbc
		_cccef = &_eagae
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectNull:
		_cccef = _eb.MakeNull()
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectInteger:
		_gefbg := *_cffbc
		_cccef = &_gefbg
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectReference:
		_gcfgb := *_cffbc
		_cccef = &_gcfgb
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectFloat:
		_ebcffd := *_cffbc
		_cccef = &_ebcffd
		_bdgbc[_adeea] = _cccef
	case *_eb.PdfObjectBool:
		_aadgd := *_cffbc
		_cccef = &_aadgd
		_bdgbc[_adeea] = _cccef
	case *pdfSignDictionary:
		_dfcbg := &pdfSignDictionary{PdfObjectDictionary: _eb.MakeDict(), _bfadg: _cffbc._bfadg, _abegc: _cffbc._abegc}
		_cccef = _dfcbg
		_bdgbc[_adeea] = _cccef
		for _, _cbcfae := range _cffbc.Keys() {
			_dfcbg.Set(_cbcfae, _cffged.copyObject(_cffbc.Get(_cbcfae), _bdgbc, _cfefe, _dcfgc))
		}
	default:
		_ddb.Log.Info("\u0054\u004f\u0044\u004f\u0028\u0061\u0035\u0069\u0029\u003a\u0020\u0069\u006dp\u006c\u0065\u006d\u0065\u006e\u0074 \u0063\u006f\u0070\u0079\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0066\u006fr\u0020\u0025\u002b\u0076", _adeea)
	}
	if _afdf && _dcfgc {
		_cfefe[_adeea] = struct{}{}
	}
	return _cccef
}

// DecodeArray returns the range of color component values in CalGray colorspace.
func (_gfac *PdfColorspaceCalGray) DecodeArray() []float64 { return []float64{0.0, 1.0} }

// GetXObjectByName returns the XObject with the specified keyName and the object type.
func (_dbbe *PdfPageResources) GetXObjectByName(keyName _eb.PdfObjectName) (*_eb.PdfObjectStream, XObjectType) {
	if _dbbe.XObject == nil {
		return nil, XObjectTypeUndefined
	}
	_ddcfa, _bcage := _eb.TraceToDirectObject(_dbbe.XObject).(*_eb.PdfObjectDictionary)
	if !_bcage {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _eb.TraceToDirectObject(_dbbe.XObject))
		return nil, XObjectTypeUndefined
	}
	if _cggaa := _ddcfa.Get(keyName); _cggaa != nil {
		_fgafd, _ddgga := _eb.GetStream(_cggaa)
		if !_ddgga {
			_ddb.Log.Debug("X\u004f\u0062\u006a\u0065\u0063\u0074 \u006e\u006f\u0074\u0020\u0070\u006fi\u006e\u0074\u0069\u006e\u0067\u0020\u0074o\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020%\u0054", _cggaa)
			return nil, XObjectTypeUndefined
		}
		_beeac := _fgafd.PdfObjectDictionary
		_dcddf, _ddgga := _eb.TraceToDirectObject(_beeac.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")).(*_eb.PdfObjectName)
		if !_ddgga {
			_ddb.Log.Debug("\u0058\u004fbj\u0065\u0063\u0074 \u0053\u0075\u0062\u0074ype\u0020no\u0074\u0020\u0061\u0020\u004e\u0061\u006de,\u0020\u0064\u0069\u0063\u0074\u003a\u0020%\u0073", _beeac.String())
			return nil, XObjectTypeUndefined
		}
		if *_dcddf == "\u0049\u006d\u0061g\u0065" {
			return _fgafd, XObjectTypeImage
		} else if *_dcddf == "\u0046\u006f\u0072\u006d" {
			return _fgafd, XObjectTypeForm
		} else if *_dcddf == "\u0050\u0053" {
			return _fgafd, XObjectTypePS
		} else {
			_ddb.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0053\u0075b\u0074\u0079\u0070\u0065\u0020\u006e\u006ft\u0020\u006b\u006e\u006f\u0077\u006e\u0020\u0028\u0025\u0073\u0029", *_dcddf)
			return nil, XObjectTypeUndefined
		}
	} else {
		return nil, XObjectTypeUndefined
	}
}

// ToPdfObject implements interface PdfModel.
func (_ggec *PdfAnnotationScreen) ToPdfObject() _eb.PdfObject {
	_ggec.PdfAnnotation.ToPdfObject()
	_bcb := _ggec._ggf
	_eecb := _bcb.PdfObject.(*_eb.PdfObjectDictionary)
	_eecb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053\u0063\u0072\u0065\u0065\u006e"))
	_eecb.SetIfNotNil("\u0054", _ggec.T)
	_eecb.SetIfNotNil("\u004d\u004b", _ggec.MK)
	_eecb.SetIfNotNil("\u0041", _ggec.A)
	_eecb.SetIfNotNil("\u0041\u0041", _ggec.AA)
	return _bcb
}
func _fegae(_ddded *fontCommon) *pdfCIDFontType2 { return &pdfCIDFontType2{fontCommon: *_ddded} }

// Val returns the color value.
func (_fafa *PdfColorDeviceGray) Val() float64 { return float64(*_fafa) }

// SetBorderWidth sets the style's border width.
func (_gdfe *PdfBorderStyle) SetBorderWidth(width float64) { _gdfe.W = &width }

// Encoder returns the font's text encoder.
func (_daacf pdfCIDFontType2) Encoder() _fc.TextEncoder { return _daacf._cafea }
func (_eccg *LTV) getCRLs(_acaf []*_bag.Certificate) ([][]byte, error) {
	_eeaef := make([][]byte, 0, len(_acaf))
	for _, _cdfge := range _acaf {
		for _, _abegf := range _cdfge.CRLDistributionPoints {
			if _eccg.CertClient.IsCA(_cdfge) {
				continue
			}
			_cdda, _fbfgf := _eccg.CRLClient.MakeRequest(_abegf, _cdfge)
			if _fbfgf != nil {
				_ddb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _fbfgf)
				continue
			}
			_eeaef = append(_eeaef, _cdda)
		}
	}
	return _eeaef, nil
}

// AddPages adds pages to be appended to the end of the source PDF.
func (_cgbad *PdfAppender) AddPages(pages ...*PdfPage) {
	for _, _badf := range pages {
		_badf = _badf.Duplicate()
		_bfcca(_badf)
		_cgbad._fedg = append(_cgbad._fedg, _badf)
	}
}

// GetContentStreams returns the content stream as an array of strings.
func (_cdebd *PdfPage) GetContentStreams() ([]string, error) {
	_agegd := _cdebd.GetContentStreamObjs()
	var _adgdg []string
	for _, _bgcdg := range _agegd {
		_dcacf, _fcce := _edba(_bgcdg)
		if _fcce != nil {
			return nil, _fcce
		}
		_adgdg = append(_adgdg, _dcacf)
	}
	return _adgdg, nil
}

// NewPdfAnnotationRichMedia returns a new rich media annotation.
func NewPdfAnnotationRichMedia() *PdfAnnotationRichMedia {
	_gba := NewPdfAnnotation()
	_eaa := &PdfAnnotationRichMedia{}
	_eaa.PdfAnnotation = _gba
	_gba.SetContext(_eaa)
	return _eaa
}

// AddPage adds a page to the PDF file. The new page should be an indirect object.
func (_dddda *PdfWriter) AddPage(page *PdfPage) error {
	const _cedd = "\u006d\u006f\u0064el\u003a\u0050\u0064\u0066\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_bfcca(page)
	_egad := page.ToPdfObject()
	_ddb.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
	_ddb.Log.Trace("\u0041p\u0070\u0065\u006e\u0064i\u006e\u0067\u0020\u0074\u006f \u0070a\u0067e\u0020\u006c\u0069\u0073\u0074\u0020\u0025T", _egad)
	_agfdc, _eeafg := _eb.GetIndirect(_egad)
	if !_eeafg {
		return _dcf.New("\u0070\u0061\u0067\u0065\u0020\u0073h\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_ddb.Log.Trace("\u0025\u0073", _agfdc)
	_ddb.Log.Trace("\u0025\u0073", _agfdc.PdfObject)
	_adga, _eeafg := _eb.GetDict(_agfdc.PdfObject)
	if !_eeafg {
		return _dcf.New("\u0070\u0061\u0067e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068o\u0075l\u0064 \u0062e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_ceccf, _eeafg := _eb.GetName(_adga.Get("\u0054\u0079\u0070\u0065"))
	if !_eeafg {
		return _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054y\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u006e\u0061m\u0065\u0020\u0028%\u0054\u0029", _adga.Get("\u0054\u0079\u0070\u0065"))
	}
	if _ceccf.String() != "\u0050\u0061\u0067\u0065" {
		return _dcf.New("\u0066\u0069e\u006c\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u0050\u0061\u0067\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069re\u0064\u0029")
	}
	_ebcff := []_eb.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	_daffe, _gdgcaa := _eb.GetIndirect(_adga.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	_ddb.Log.Trace("P\u0061g\u0065\u0020\u0050\u0061\u0072\u0065\u006e\u0074:\u0020\u0025\u0054\u0020(%\u0076\u0029", _adga.Get("\u0050\u0061\u0072\u0065\u006e\u0074"), _gdgcaa)
	for _gdgcaa {
		_ddb.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _daffe)
		_eadbf, _gbcda := _eb.GetDict(_daffe.PdfObject)
		if !_gbcda {
			return _dcf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		for _, _gdbb := range _ebcff {
			_ddb.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _gdbb)
			if _adga.Get(_gdbb) != nil {
				_ddb.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
				continue
			}
			if _cbga := _eadbf.Get(_gdbb); _cbga != nil {
				_ddb.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _gdbb)
				_adga.Set(_gdbb, _cbga)
			}
		}
		_daffe, _gdgcaa = _eb.GetIndirect(_eadbf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
		_ddb.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _eadbf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	}
	_ddb.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c \u0064\u006f\u006e\u0065")
	_adga.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _dddda._afga)
	_agfdc.PdfObject = _adga
	_bdfcb, _eeafg := _eb.GetDict(_dddda._afga.PdfObject)
	if !_eeafg {
		return _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_adccce, _eeafg := _eb.GetArray(_bdfcb.Get("\u004b\u0069\u0064\u0073"))
	if !_eeafg {
		return _dcf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_adccce.Append(_agfdc)
	_dddda._gbgc[_adga] = struct{}{}
	_dddda._aadfg = append(_dddda._aadfg, _agfdc)
	_cdgf, _eeafg := _eb.GetInt(_bdfcb.Get("\u0043\u006f\u0075n\u0074"))
	if !_eeafg {
		return _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	*_cdgf = *_cdgf + 1
	if page._dcdfd == nil {
		_ffaad := _cf.Track(_dddda._badcb, _cedd, _dddda._adeee)
		if _ffaad != nil {
			return _ffaad
		}
	} else {
		_gegbg := _cf.Track(page._dcdfd._cadbf, _cedd, page._dcdfd._fbgfg)
		if _gegbg != nil {
			return _gegbg
		}
	}
	_dddda.addObject(_agfdc)
	_agagd := _dddda.addObjects(_adga)
	if _agagd != nil {
		return _agagd
	}
	return nil
}

// KDict represents a K dictionary object.
type KDict struct {

	// The structure type, a name object identifying the nature of the
	// structure element and its role within the document,
	// such as a chapter, paragraph, or footnote
	S _eb.PdfObject

	// The structure element that is the immediate parent of this one
	// in the structure hierarchy.
	P _eb.PdfObject

	// The element identifier, a byte string designating this structure element.
	ID *_eb.PdfObjectString

	// A page object representing a page on which some or all of the content
	// items designated by the K entry shall be rendered.
	Pg _eb.PdfObject

	// The children of this structure element.
	K _eb.PdfObject

	// A single attribute object or array of attribute objects associated
	// with this structure element.
	A _eb.PdfObject

	// An attribute class name or array of class names associated with
	// this structure element.
	C _eb.PdfObject

	// The current revision number of this structure element
	R *_eb.PdfObjectInteger

	// The title of the structure element, a text string representing
	// it in human-readable form.
	T *_eb.PdfObjectString

	// A language identifier specifying the natural language for all text
	// in the structure element except where overridden by
	// language specifications for nested structure elements or marked content.
	Lang *_eb.PdfObjectString

	// An alternate description of the structure element and its
	// children in human-readable form, which is useful when extracting
	// the document’s contents in support of accessibility to users with
	// disabilities or for other purposes.
	Alt *_eb.PdfObjectString

	// The expanded form of an abbreviation.
	E *_eb.PdfObjectString

	// Text that is an exact replacement for the structure element and its children.
	ActualText *_eb.PdfObjectString
	_eegge     []*KValue
	_dcgce     int64
	_aagfc     *PdfRectangle
}

// PdfFunctionType2 defines an exponential interpolation of one input value and n
// output values:
//
//	f(x) = y_0, ..., y_(n-1)
//
// y_j = C0_j + x^N * (C1_j - C0_j); for 0 <= j < n
// When N=1 ; linear interpolation between C0 and C1.
type PdfFunctionType2 struct {
	Domain []float64
	Range  []float64
	C0     []float64
	C1     []float64
	N      float64
	_efee  *_eb.PdfIndirectObject
}

// AddContentStreamByString adds content stream by string. Puts the content
// string into a stream object and points the content stream towards it.
func (_ebgc *PdfPage) AddContentStreamByString(contentStr string) error {
	_bbgfg, _cabdd := _eb.MakeStream([]byte(contentStr), _eb.NewFlateEncoder())
	if _cabdd != nil {
		return _cabdd
	}
	if _ebgc.Contents == nil {
		_ebgc.Contents = _bbgfg
	} else {
		_ddbc := _eb.TraceToDirectObject(_ebgc.Contents)
		_afba, _dfgb := _ddbc.(*_eb.PdfObjectArray)
		if !_dfgb {
			_afba = _eb.MakeArray(_ddbc)
		}
		_afba.Append(_bbgfg)
		_ebgc.Contents = _afba
	}
	return nil
}

// SetPageLabels sets the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_bafb *PdfWriter) SetPageLabels(pageLabels _eb.PdfObject) error {
	if pageLabels == nil {
		return nil
	}
	_ddb.Log.Trace("\u0053\u0065t\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006cs.\u002e\u002e")
	_bafb._dbffa.Set("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073", pageLabels)
	return _bafb.addObjects(pageLabels)
}

// SetPdfTitle sets the Title attribute of the output PDF.
func SetPdfTitle(title string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _gacdf = title }

// ToPdfObject generates a PdfObject representation of the Names struct.
func (_degeb *Names) ToPdfObject() _eb.PdfObject {
	_bdgb := _degeb._cggag
	_gecef, _eacdb := _bdgb.PdfObject.(*_eb.PdfObjectDictionary)
	if !_eacdb {
		_ddb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _bdgb
	}
	if _degeb.EmbeddedFiles != nil {
		_gecef.Set("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073", _eb.MakeIndirectObject(_degeb.EmbeddedFiles))
	}
	return _bdgb
}
func _eabfd() string {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	_edgaa := _cf.GetLicenseKey()
	if len(_ccgde) > 0 && (_edgaa.IsLicensed() || _fefac) {
		return _ccgde
	}
	return _e.Sprintf("\u0055\u006e\u0069Do\u0063\u0020\u0076\u0025\u0073\u0020\u0028\u0025\u0073)\u0020-\u0020h\u0074t\u0070\u003a\u002f\u002f\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f", _dgeac(), _edgaa.TypeToString())
}

// PdfReader represents a PDF file reader. It is a frontend to the lower level parsing mechanism and provides
// a higher level access to work with PDF structure and information, such as the page structure etc.
type PdfReader struct {
	_ebbe    *_eb.PdfParser
	_ecddba  _eb.PdfObject
	_agfdg   *_eb.PdfIndirectObject
	_bgdad   *_eb.PdfObjectDictionary
	_cbaff   []*_eb.PdfIndirectObject
	PageList []*PdfPage
	_cdbc    int
	_bagcfd  *_eb.PdfObjectDictionary
	_efabg   *PdfOutlineTreeNode
	AcroForm *PdfAcroForm
	DSS      *DSS
	Rotate   *int64
	_bdffe   *Permissions
	_dbag    map[*PdfReader]*PdfReader
	_dfeg    []*PdfReader
	_affaf   *modelManager
	_cfcgdf  bool
	_bcefc   map[_eb.PdfObject]struct{}
	_cbeg    _bagf.ReadSeeker
	_cadbf   string
	_fafga   bool
	_fbgfg   string
	_eacdg   *ReaderOpts
	_daag    bool
}

// ImageToGray returns a new grayscale image based on the passed in RGB image.
func (_aecb *PdfColorspaceDeviceRGB) ImageToGray(img Image) (Image, error) {
	if img.ColorComponents != 3 {
		return img, _dcf.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042")
	}
	_dfe, _aacd := _df.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._bdcab, img._fedc)
	if _aacd != nil {
		return img, _aacd
	}
	_eged, _aacd := _df.GrayConverter.Convert(_dfe)
	if _aacd != nil {
		return img, _aacd
	}
	return _ggaa(_eged.Base()), nil
}
func (_faae *PdfReader) loadStructure() error {
	if _faae._ebbe.GetCrypter() != nil && !_faae._ebbe.IsAuthenticated() {
		return _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_fabbd := _faae._ebbe.GetTrailer()
	if _fabbd == nil {
		return _e.Errorf("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_cadac, _fgbaa := _fabbd.Get("\u0052\u006f\u006f\u0074").(*_eb.PdfObjectReference)
	if !_fgbaa {
		return _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052\u006f\u006ft\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a \u0025\u0073\u0029", _fabbd)
	}
	_egfa, _dbagb := _faae._ebbe.LookupByReference(*_cadac)
	if _dbagb != nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0072\u006f\u006f\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _dbagb)
		return _dbagb
	}
	_efgbc, _fgbaa := _egfa.(*_eb.PdfIndirectObject)
	if !_fgbaa {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _egfa, *_fabbd)
		return _dcf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_fadc, _fgbaa := (*_efgbc).PdfObject.(*_eb.PdfObjectDictionary)
	if !_fgbaa {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0061t\u0061\u006c\u006fg\u0020(\u0025\u0073\u0029", _efgbc.PdfObject)
		return _dcf.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_ddb.Log.Trace("C\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _fadc)
	_edaef, _fgbaa := _fadc.Get("\u0050\u0061\u0067e\u0073").(*_eb.PdfObjectReference)
	if !_fgbaa {
		return _dcf.New("\u0070\u0061\u0067\u0065\u0073\u0020\u0069\u006e\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0061\u0020\u0072e\u0066\u0065r\u0065\u006e\u0063\u0065")
	}
	_eaaga, _dbagb := _faae._ebbe.LookupByReference(*_edaef)
	if _dbagb != nil {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020r\u0065\u0061\u0064 \u0070a\u0067\u0065\u0073")
		return _dbagb
	}
	_bdabg, _fgbaa := _eaaga.(*_eb.PdfIndirectObject)
	if !_fgbaa {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020P\u0061\u0067\u0065\u0073\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0069n\u0076a\u006c\u0069\u0064")
		_ddb.Log.Debug("\u006f\u0070\u003a\u0020\u0025\u0070", _bdabg)
		return _dcf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_dagcd, _fgbaa := _bdabg.PdfObject.(*_eb.PdfObjectDictionary)
	if !_fgbaa {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065c\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0073\u0029", _bdabg)
		return _dcf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_bfdgd, _fgbaa := _eb.GetInt(_dagcd.Get("\u0043\u006f\u0075n\u0074"))
	if !_fgbaa {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0050\u0061\u0067\u0065\u0073\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return _dcf.New("\u0070\u0061\u0067\u0065s \u0063\u006f\u0075\u006e\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _, _fgbaa = _eb.GetName(_dagcd.Get("\u0054\u0079\u0070\u0065")); !_fgbaa {
		_ddb.Log.Debug("\u0050\u0061\u0067\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020T\u0079\u0070\u0065\u0020\u0066\u0069\u0065\u006cd\u0020n\u006f\u0074\u0020\u0073\u0065\u0074\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0054\u0079p\u0065\u0020\u0074\u006f\u0020\u0050\u0061\u0067\u0065\u0073\u002e")
		_dagcd.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0050\u0061\u0067e\u0073"))
	}
	if _ffcfd, _fdgdg := _eb.GetInt(_dagcd.Get("\u0052\u006f\u0074\u0061\u0074\u0065")); _fdgdg {
		_bcfg := int64(*_ffcfd)
		_faae.Rotate = &_bcfg
	}
	_faae._ecddba = _cadac
	_faae._bagcfd = _fadc
	_faae._bgdad = _dagcd
	_faae._agfdg = _bdabg
	_faae._cdbc = int(*_bfdgd)
	_faae._cbaff = []*_eb.PdfIndirectObject{}
	_cccfag := map[_eb.PdfObject]struct{}{}
	_dbagb = _faae.buildPageList(_bdabg, nil, _cccfag)
	if _dbagb != nil {
		return _dbagb
	}
	_ddb.Log.Trace("\u002d\u002d\u002d")
	_ddb.Log.Trace("\u0054\u004f\u0043")
	_ddb.Log.Trace("\u0050\u0061\u0067e\u0073")
	_ddb.Log.Trace("\u0025\u0064\u003a\u0020\u0025\u0073", len(_faae._cbaff), _faae._cbaff)
	_faae._efabg, _dbagb = _faae.loadOutlines()
	if _dbagb != nil {
		_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075i\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 t\u0072\u0065\u0065 \u0028%\u0073\u0029", _dbagb)
		return _dbagb
	}
	_faae.AcroForm, _dbagb = _faae.loadForms()
	if _dbagb != nil {
		return _dbagb
	}
	_faae.DSS, _dbagb = _faae.loadDSS()
	if _dbagb != nil {
		return _dbagb
	}
	_faae._bdffe, _dbagb = _faae.loadPerms()
	if _dbagb != nil {
		return _dbagb
	}
	return nil
}

// GetPage returns the PdfPage model for the specified page number.
func (_dbbb *PdfReader) GetPage(pageNumber int) (*PdfPage, error) {
	if _dbbb._ebbe.GetCrypter() != nil && !_dbbb._ebbe.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	if len(_dbbb._cbaff) < pageNumber {
		return nil, _dcf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0028\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0074o\u006f\u0020\u0073\u0068\u006f\u0072\u0074\u0029")
	}
	_bgacdc := pageNumber - 1
	if _bgacdc < 0 {
		return nil, _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0069\u006e\u0067\u0020\u006d\u0075\u0073t\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0061\u0074\u0020\u0031")
	}
	_eecfc := _dbbb.PageList[_bgacdc]
	return _eecfc, nil
}

// ToPdfObject implements interface PdfModel.
func (_abgdb *PdfFilespec) ToPdfObject() _eb.PdfObject {
	_bfdef := _abgdb.getDict()
	_bfdef.Clear()
	_bfdef.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063"))
	_bfdef.SetIfNotNil("\u0046\u0053", _abgdb.FS)
	_bfdef.SetIfNotNil("\u0046", _abgdb.F)
	_bfdef.SetIfNotNil("\u0055\u0046", _abgdb.UF)
	_bfdef.SetIfNotNil("\u0044\u004f\u0053", _abgdb.DOS)
	_bfdef.SetIfNotNil("\u004d\u0061\u0063", _abgdb.Mac)
	_bfdef.SetIfNotNil("\u0055\u006e\u0069\u0078", _abgdb.Unix)
	_bfdef.SetIfNotNil("\u0049\u0044", _abgdb.ID)
	_bfdef.SetIfNotNil("\u0056", _abgdb.V)
	_bfdef.SetIfNotNil("\u0045\u0046", _abgdb.EF)
	_bfdef.SetIfNotNil("\u0052\u0046", _abgdb.RF)
	_bfdef.SetIfNotNil("\u0044\u0065\u0073\u0063", _abgdb.Desc)
	_bfdef.SetIfNotNil("\u0043\u0049", _abgdb.CI)
	_bfdef.SetIfNotNil("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070", _abgdb.AFRelationship)
	return _abgdb._cbefe
}

// PdfShadingPatternType3 is shading patterns that will use a Type 3 shading pattern (Radial).
type PdfShadingPatternType3 struct {
	*PdfPattern
	Shading   *PdfShadingType3
	Matrix    *_eb.PdfObjectArray
	ExtGState _eb.PdfObject
}

// GetContainingPdfObject returns the container of the resources object (indirect object).
func (_becba *PdfPageResources) GetContainingPdfObject() _eb.PdfObject { return _becba._fdada }

// FieldAppearanceGenerator generates appearance stream for a given field.
type FieldAppearanceGenerator interface {
	ContentStreamWrapper
	GenerateAppearanceDict(_dacgd *PdfAcroForm, _dbed *PdfField, _ggbde *PdfAnnotationWidget) (*_eb.PdfObjectDictionary, error)
}

// ToPdfObject sets the common field elements.
// Note: Call the more field context's ToPdfObject to set both the generic and
// non-generic information.
func (_gaeb *PdfField) ToPdfObject() _eb.PdfObject {
	_cdbdb := _gaeb._adgda
	_aeaf := _cdbdb.PdfObject.(*_eb.PdfObjectDictionary)
	_fcec := _eb.MakeArray()
	for _, _bdee := range _gaeb.Kids {
		_fcec.Append(_bdee.ToPdfObject())
	}
	for _, _cbeeg := range _gaeb.Annotations {
		if _cbeeg._ggf != _gaeb._adgda {
			_fcec.Append(_cbeeg.GetContext().ToPdfObject())
		}
	}
	if _gaeb.Parent != nil {
		_aeaf.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _gaeb.Parent.GetContainingPdfObject())
	}
	if _fcec.Len() > 0 {
		_aeaf.Set("\u004b\u0069\u0064\u0073", _fcec)
	}
	_aeaf.SetIfNotNil("\u0046\u0054", _gaeb.FT)
	_aeaf.SetIfNotNil("\u0054", _gaeb.T)
	_aeaf.SetIfNotNil("\u0054\u0055", _gaeb.TU)
	_aeaf.SetIfNotNil("\u0054\u004d", _gaeb.TM)
	_aeaf.SetIfNotNil("\u0046\u0066", _gaeb.Ff)
	_aeaf.SetIfNotNil("\u0056", _gaeb.V)
	_aeaf.SetIfNotNil("\u0044\u0056", _gaeb.DV)
	_aeaf.SetIfNotNil("\u0041\u0041", _gaeb.AA)
	if _gaeb.VariableText != nil {
		_aeaf.SetIfNotNil("\u0044\u0041", _gaeb.VariableText.DA)
		_aeaf.SetIfNotNil("\u0051", _gaeb.VariableText.Q)
		_aeaf.SetIfNotNil("\u0044\u0053", _gaeb.VariableText.DS)
		_aeaf.SetIfNotNil("\u0052\u0056", _gaeb.VariableText.RV)
	}
	return _cdbdb
}
func _gaefa() *Names { return &Names{_cggag: _eb.MakeIndirectObject(_eb.MakeDict())} }

// NewDSS returns a new DSS dictionary.
func NewDSS() *DSS {
	return &DSS{_fcdgc: _eb.MakeIndirectObject(_eb.MakeDict()), VRI: map[string]*VRI{}}
}

// NewPdfActionHide returns a new "hide" action.
func NewPdfActionHide() *PdfActionHide {
	_eggf := NewPdfAction()
	_gab := &PdfActionHide{}
	_gab.PdfAction = _eggf
	_eggf.SetContext(_gab)
	return _gab
}

// GetNumComponents returns the number of color components of the underlying
// colorspace device.
func (_bfbd *PdfColorspaceSpecialPattern) GetNumComponents() int {
	return _bfbd.UnderlyingCS.GetNumComponents()
}
func (_ceddg *PdfWriter) writeTrailer(_dgfbce int) {
	_ceddg.writeString("\u0078\u0072\u0065\u0066\u000d\u000a")
	for _eefadc := 0; _eefadc <= _dgfbce; {
		for ; _eefadc <= _dgfbce; _eefadc++ {
			_gbcbd, _addgg := _ceddg._aggfdb[_eefadc]
			if _addgg && (!_ceddg._caed || _ceddg._caed && (_gbcbd.Type == 1 && _gbcbd.Offset >= _ceddg._dfbcf || _gbcbd.Type == 0)) {
				break
			}
		}
		var _bgce int
		for _bgce = _eefadc + 1; _bgce <= _dgfbce; _bgce++ {
			_caccg, _cedbf := _ceddg._aggfdb[_bgce]
			if _cedbf && (!_ceddg._caed || _ceddg._caed && (_caccg.Type == 1 && _caccg.Offset > _ceddg._dfbcf)) {
				continue
			}
			break
		}
		_fadgc := _e.Sprintf("\u0025d\u0020\u0025\u0064\u000d\u000a", _eefadc, _bgce-_eefadc)
		_ceddg.writeString(_fadgc)
		for _cgbdc := _eefadc; _cgbdc < _bgce; _cgbdc++ {
			_edgea := _ceddg._aggfdb[_cgbdc]
			switch _edgea.Type {
			case 0:
				_fadgc = _e.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u0066\u000d\u000a", 0, 65535)
				_ceddg.writeString(_fadgc)
			case 1:
				_fadgc = _e.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u006e\u000d\u000a", _edgea.Offset, 0)
				_ceddg.writeString(_fadgc)
			}
		}
		_eefadc = _bgce + 1
	}
	_eaebdf := _eb.MakeDict()
	_eaebdf.Set("\u0049\u006e\u0066\u006f", _ceddg._ecagd)
	_eaebdf.Set("\u0052\u006f\u006f\u0074", _ceddg._ccea)
	_eaebdf.Set("\u0053\u0069\u007a\u0065", _eb.MakeInteger(int64(_dgfbce+1)))
	if _ceddg._caed && _ceddg._becbf > 0 {
		_eaebdf.Set("\u0050\u0072\u0065\u0076", _eb.MakeInteger(_ceddg._becbf))
	}
	if _ceddg._gadcaa != nil {
		_eaebdf.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _ceddg._ebdeed)
	}
	if _ceddg._ddefd == nil && _ceddg._cebgc != "" && _ceddg._bgddbe != "" {
		_ceddg._ddefd = _eb.MakeArray(_eb.MakeHexString(_ceddg._cebgc), _eb.MakeHexString(_ceddg._bgddbe))
	}
	if _ceddg._ddefd != nil {
		_eaebdf.Set("\u0049\u0044", _ceddg._ddefd)
		_ddb.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _ceddg._ddefd)
	}
	_ceddg.writeString("\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u000a")
	_ceddg.writeString(_eaebdf.WriteString())
	_ceddg.writeString("\u000a")
}

// GetSamples converts the raw byte slice into samples which are stored in a uint32 bit array.
// Each sample is represented by BitsPerComponent consecutive bits in the raw data.
// NOTE: The method resamples the image byte data before returning the result and
// this could lead to high memory usage, especially on large images. It should
// be avoided, when possible. It is recommended to access the Data field of the
// image directly or use the ColorAt method to extract individual pixels.
func (_gdefb *Image) GetSamples() []uint32 {
	_aebd := _bb.ResampleBytes(_gdefb.Data, int(_gdefb.BitsPerComponent))
	if _gdefb.BitsPerComponent < 8 {
		_aebd = _gdefb.samplesTrimPadding(_aebd)
	}
	_ebacb := int(_gdefb.Width) * int(_gdefb.Height) * _gdefb.ColorComponents
	if len(_aebd) < _ebacb {
		_ddb.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u0066\u0065\u0077\u0020\u0073a\u006d\u0070\u006c\u0065\u0073\u0020(\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064\u0029", len(_aebd), _ebacb)
		return _aebd
	} else if len(_aebd) > _ebacb {
		_ddb.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u006d\u0061\u006e\u0079\u0020s\u0061\u006d\u0070\u006c\u0065\u0073 \u0028\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078p\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064", len(_aebd), _ebacb)
		_aebd = _aebd[:_ebacb]
	}
	return _aebd
}

// NewPdfSignatureReferenceDocMDP returns PdfSignatureReference for the transformParams.
func NewPdfSignatureReferenceDocMDP(transformParams *PdfTransformParamsDocMDP) *PdfSignatureReference {
	return &PdfSignatureReference{Type: _eb.MakeName("\u0053\u0069\u0067\u0052\u0065\u0066"), TransformMethod: _eb.MakeName("\u0044\u006f\u0063\u004d\u0044\u0050"), TransformParams: transformParams.ToPdfObject()}
}

// ReplacePage replaces the original page to a new page.
func (_effb *PdfAppender) ReplacePage(pageNum int, page *PdfPage) {
	_fcgf := pageNum - 1
	for _fffe := range _effb._fedg {
		if _fffe == _fcgf {
			_ebab := page.Duplicate()
			_bfcca(_ebab)
			_effb._fedg[_fffe] = _ebab
		}
	}
}

type pdfCIDFontType2 struct {
	fontCommon
	_fceef *_eb.PdfIndirectObject
	_cafea _fc.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// Dictionary that defines the character collection of the CIDFont (required).
	// See Table 116.
	CIDSystemInfo *_eb.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW  _eb.PdfObject
	W   _eb.PdfObject
	DW2 _eb.PdfObject
	W2  _eb.PdfObject

	// CIDs to glyph indices mapping (optional).
	CIDToGIDMap _eb.PdfObject
	_fcgg       map[_fc.CharCode]float64
	_aggd       float64
	_caba       map[rune]int
}

// ToPdfObject implements interface PdfModel.
func (_gfa *PdfActionTrans) ToPdfObject() _eb.PdfObject {
	_gfa.PdfAction.ToPdfObject()
	_bgd := _gfa._dee
	_acg := _bgd.PdfObject.(*_eb.PdfObjectDictionary)
	_acg.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeTrans)))
	_acg.SetIfNotNil("\u0054\u0072\u0061n\u0073", _gfa.Trans)
	return _bgd
}

// Image interface is a basic representation of an image used in PDF.
// The colorspace is not specified, but must be known when handling the image.
type Image struct {
	Width            int64
	Height           int64
	BitsPerComponent int64
	ColorComponents  int
	Data             []byte
	_bdcab           []byte
	_fedc            []float64
}

// multiFontEncoder implements a an Encoder that holds a list of fonts provided.
type MultipleFontEncoder struct {
	_fcgbc      []*PdfFont
	CurrentFont *PdfFont
}

// NewPdfColorspaceDeviceN returns an initialized PdfColorspaceDeviceN.
func NewPdfColorspaceDeviceN() *PdfColorspaceDeviceN {
	_eebcd := &PdfColorspaceDeviceN{}
	return _eebcd
}

// Transform rectangle with the supplied matrix.
func (_ecead *PdfRectangle) Transform(transformMatrix _ffg.Matrix) {
	_ecead.Llx, _ecead.Lly = transformMatrix.Transform(_ecead.Llx, _ecead.Lly)
	_ecead.Urx, _ecead.Ury = transformMatrix.Transform(_ecead.Urx, _ecead.Ury)
	_ecead.Normalize()
}

// SetDuplex sets the value of the duplex.
func (_bbbbd *ViewerPreferences) SetDuplex(duplex Duplex) { _bbbbd._ebabd = duplex }

// Set sets the colorspace corresponding to key. Add to Names if not set.
func (_cbcdb *PdfPageResourcesColorspaces) Set(key _eb.PdfObjectName, val PdfColorspace) {
	if _, _cfcda := _cbcdb.Colorspaces[string(key)]; !_cfcda {
		_cbcdb.Names = append(_cbcdb.Names, string(key))
	}
	_cbcdb.Colorspaces[string(key)] = val
}

// AddExtGState add External Graphics State (GState). The gsDict can be specified
// either directly as a dictionary or an indirect object containing a dictionary.
func (_dgae *PdfPageResources) AddExtGState(gsName _eb.PdfObjectName, gsDict _eb.PdfObject) error {
	if _dgae.ExtGState == nil {
		_dgae.ExtGState = _eb.MakeDict()
	}
	_bbeed := _dgae.ExtGState
	_adef, _cfgfc := _eb.TraceToDirectObject(_bbeed).(*_eb.PdfObjectDictionary)
	if !_cfgfc {
		_ddb.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020e\u0072r\u006f\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u002f\u0025\u0054\u0029", _bbeed, _eb.TraceToDirectObject(_bbeed))
		return _eb.ErrTypeError
	}
	_adef.Set(gsName, gsDict)
	return nil
}

// ToInteger convert to an integer format.
func (_gedc *PdfColorDeviceGray) ToInteger(bits int) uint32 {
	_ecegb := _gg.Pow(2, float64(bits)) - 1
	return uint32(_ecegb * _gedc.Val())
}

// BaseFont returns the font's "BaseFont" field.
func (_fbbb *PdfFont) BaseFont() string { return _fbbb.baseFields()._agcc }

// NewPdfColorspaceDeviceGray returns a new grayscale colorspace.
func NewPdfColorspaceDeviceGray() *PdfColorspaceDeviceGray { return &PdfColorspaceDeviceGray{} }
func (_agdd *Image) samplesTrimPadding(_fefea []uint32) []uint32 {
	_ecfgg := _agdd.ColorComponents * int(_agdd.Width) * int(_agdd.Height)
	if len(_fefea) == _ecfgg {
		return _fefea
	}
	_ecea := make([]uint32, _ecfgg)
	_adeab := int(_agdd.Width) * _agdd.ColorComponents
	var _abcff, _bdaaa, _facce, _eafgb int
	_bccb := _df.BytesPerLine(int(_agdd.Width), int(_agdd.BitsPerComponent), _agdd.ColorComponents)
	for _abcff = 0; _abcff < int(_agdd.Height); _abcff++ {
		_bdaaa = _abcff * int(_agdd.Width)
		_facce = _abcff * _bccb
		for _eafgb = 0; _eafgb < _adeab; _eafgb++ {
			_ecea[_bdaaa+_eafgb] = _fefea[_facce+_eafgb]
		}
	}
	return _ecea
}

// ToPdfObject implements interface PdfModel.
func (_gadef *Permissions) ToPdfObject() _eb.PdfObject { return _gadef._cgbag }

// GetCharMetrics returns the char metrics for character code `code`.
func (_fabec pdfCIDFontType0) GetCharMetrics(code _fc.CharCode) (_fg.CharMetrics, bool) {
	_cafb := _fabec._dcfg
	if _dbada, _acdcd := _fabec._efege[code]; _acdcd {
		_cafb = _dbada
	}
	return _fg.CharMetrics{Wx: _cafb}, true
}

// G returns the value of the green component of the color.
func (_ffc *PdfColorDeviceRGB) G() float64 { return _ffc[1] }
func _ggaac(_ffcbg rune) string {
	for _cagcee, _fbcgf := range _eg.Categories {
		if len(_cagcee) == 2 && _eg.Is(_fbcgf, _ffcbg) {
			return _cagcee
		}
	}
	return "\u0043\u006e"
}

// NewPdfAnnotationInk returns a new ink annotation.
func NewPdfAnnotationInk() *PdfAnnotationInk {
	_feg := NewPdfAnnotation()
	_cbf := &PdfAnnotationInk{}
	_cbf.PdfAnnotation = _feg
	_cbf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_feg.SetContext(_cbf)
	return _cbf
}
func (_fgdeca *pdfFontSimple) updateStandard14Font() {
	_edgfb, _geeeb := _fgdeca.Encoder().(_fc.SimpleEncoder)
	if !_geeeb {
		_ddb.Log.Error("\u0057\u0072\u006f\u006e\u0067\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074y\u0070e\u003a\u0020\u0025\u0054\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u002e", _fgdeca.Encoder(), _fgdeca)
		return
	}
	_deggf := _edgfb.Charcodes()
	_fgdeca._cegda = make(map[_fc.CharCode]float64, len(_deggf))
	for _, _afefc := range _deggf {
		_cegaa, _ := _edgfb.CharcodeToRune(_afefc)
		_dfbeb, _ := _fgdeca._gacdb.Read(_cegaa)
		_fgdeca._cegda[_afefc] = _dfbeb.Wx
	}
}

// PdfAnnotationPrinterMark represents PrinterMark annotations.
// (Section 12.5.6.20).
type PdfAnnotationPrinterMark struct {
	*PdfAnnotation
	MN _eb.PdfObject
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_ggeab *PdfFunctionType3) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ddb.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return nil, _dcf.New("\u006e\u006f\u0074\u0020im\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
}

// SetSamples convert samples to byte-data and sets for the image.
// NOTE: The method resamples the data and this could lead to high memory usage,
// especially on large images. It should be used only when it is not possible
// to work with the image byte data directly.
func (_aeabf *Image) SetSamples(samples []uint32) {
	if _aeabf.BitsPerComponent < 8 {
		samples = _aeabf.samplesAddPadding(samples)
	}
	_bbege := _bb.ResampleUint32(samples, int(_aeabf.BitsPerComponent), 8)
	_gcab := make([]byte, len(_bbege))
	for _eecae, _bcgee := range _bbege {
		_gcab[_eecae] = byte(_bcgee)
	}
	_aeabf.Data = _gcab
}

// PdfAppender appends new PDF content to an existing PDF document via incremental updates.
type PdfAppender struct {
	_bccd  _bagf.ReadSeeker
	_ebcb  *_eb.PdfParser
	_adac  *PdfReader
	Reader *PdfReader
	_fedg  []*PdfPage
	_bddda *PdfAcroForm
	_cedc  *DSS
	_fbec  *Permissions
	_dce   _eb.XrefTable
	_dgc   int64
	_cfgdf int
	_dfg   []_eb.PdfObject
	_accfg map[_eb.PdfObject]struct{}
	_ebcbc map[_eb.PdfObject]int64
	_dbae  map[_eb.PdfObject]struct{}
	_fdfg  map[_eb.PdfObject]struct{}
	_eabd  int64
	_ffac  bool
	_cedcf string
	_fcfd  *EncryptOptions
	_ccag  *PdfInfo
}

// Clear clears the KValue.
func (_ecbaa *KValue) Clear()                 { _ecbaa._ccbca = nil; _ecbaa._fbgaa = nil; _ecbaa._ddaf = nil }
func _adbdb(_agadf *fontCommon) *pdfFontType3 { return &pdfFontType3{fontCommon: *_agadf} }
func _fdegcc() _d.Time {
	_dfbafc.Lock()
	defer _dfbafc.Unlock()
	return _eegaeb
}

// FieldFilterFunc represents a PDF field filtering function. If the function
// returns true, the PDF field is kept, otherwise it is discarded.
type FieldFilterFunc func(*PdfField) bool

func _eadcd(_dabc []byte) []byte {
	const _dddcb = 52845
	const _cefedg = 22719
	_aafcc := 55665
	for _, _gbbe := range _dabc[:4] {
		_aafcc = (int(_gbbe)+_aafcc)*_dddcb + _cefedg
	}
	_aeec := make([]byte, len(_dabc)-4)
	for _daedf, _fgcda := range _dabc[4:] {
		_aeec[_daedf] = byte(int(_fgcda) ^ _aafcc>>8)
		_aafcc = (int(_fgcda)+_aafcc)*_dddcb + _cefedg
	}
	return _aeec
}

// ColorToRGB converts a ICCBased color to an RGB color.
func (_cbed *PdfColorspaceICCBased) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _cbed.Alternate == nil {
		_ddb.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _cbed.N == 1 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_bdeb := NewPdfColorspaceDeviceGray()
			return _bdeb.ColorToRGB(color)
		} else if _cbed.N == 3 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return color, nil
		} else if _cbed.N == 4 {
			_ddb.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_abecf := NewPdfColorspaceDeviceCMYK()
			return _abecf.ColorToRGB(color)
		} else {
			return nil, _dcf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ddb.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _cbed)
	return _cbed.Alternate.ColorToRGB(color)
}
func (_cadb *PdfReader) newPdfAnnotationLineFromDict(_dff *_eb.PdfObjectDictionary) (*PdfAnnotationLine, error) {
	_acad := PdfAnnotationLine{}
	_efcf, _ead := _cadb.newPdfAnnotationMarkupFromDict(_dff)
	if _ead != nil {
		return nil, _ead
	}
	_acad.PdfAnnotationMarkup = _efcf
	_acad.L = _dff.Get("\u004c")
	_acad.BS = _dff.Get("\u0042\u0053")
	_acad.LE = _dff.Get("\u004c\u0045")
	_acad.IC = _dff.Get("\u0049\u0043")
	_acad.LL = _dff.Get("\u004c\u004c")
	_acad.LLE = _dff.Get("\u004c\u004c\u0045")
	_acad.Cap = _dff.Get("\u0043\u0061\u0070")
	_acad.IT = _dff.Get("\u0049\u0054")
	_acad.LLO = _dff.Get("\u004c\u004c\u004f")
	_acad.CP = _dff.Get("\u0043\u0050")
	_acad.Measure = _dff.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	_acad.CO = _dff.Get("\u0043\u004f")
	return &_acad, nil
}

// SetMCID sets the MCID for the KValue.
func (_cdbdg *KValue) SetMCID(mcid int) { _cdbdg.Clear(); _cdbdg._ddaf = &mcid }
func _cadd(_gddfd _eb.PdfObject) *Names {
	_fcabd := _gaefa()
	_babfg := _eb.TraceToDirectObject(_gddfd).(*_eb.PdfObjectDictionary)
	if _dcda := _babfg.Get("\u0044\u0065\u0073t\u0073"); _dcda != nil {
		_fcabd.Dests = _eb.TraceToDirectObject(_dcda).(*_eb.PdfObjectDictionary)
	}
	if _eefcc := _babfg.Get("\u0041\u0050"); _eefcc != nil {
		_fcabd.AP = _eb.TraceToDirectObject(_eefcc).(*_eb.PdfObjectDictionary)
	}
	if _gcfaa := _babfg.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _gcfaa != nil {
		_fcabd.JavaScript = _eb.TraceToDirectObject(_gcfaa).(*_eb.PdfObjectDictionary)
	}
	if _bagbg := _babfg.Get("\u0050\u0061\u0067e\u0073"); _bagbg != nil {
		_fcabd.Pages = _eb.TraceToDirectObject(_bagbg).(*_eb.PdfObjectDictionary)
	}
	if _dcada := _babfg.Get("\u0054e\u006d\u0070\u006c\u0061\u0074\u0065s"); _dcada != nil {
		_fcabd.Templates = _eb.TraceToDirectObject(_dcada).(*_eb.PdfObjectDictionary)
	}
	if _acefb := _babfg.Get("\u0049\u0044\u0053"); _acefb != nil {
		_fcabd.IDS = _eb.TraceToDirectObject(_acefb).(*_eb.PdfObjectDictionary)
	}
	if _bcdad := _babfg.Get("\u0055\u0052\u004c\u0053"); _bcdad != nil {
		_fcabd.URLS = _eb.TraceToDirectObject(_bcdad).(*_eb.PdfObjectDictionary)
	}
	if _cgacb := _babfg.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073"); _cgacb != nil {
		_fcabd.EmbeddedFiles = _eb.TraceToDirectObject(_cgacb).(*_eb.PdfObjectDictionary)
	}
	if _fgggb := _babfg.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073"); _fgggb != nil {
		_fcabd.AlternatePresentations = _eb.TraceToDirectObject(_fgggb).(*_eb.PdfObjectDictionary)
	}
	if _daeba := _babfg.Get("\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073"); _daeba != nil {
		_fcabd.Renditions = _eb.TraceToDirectObject(_daeba).(*_eb.PdfObjectDictionary)
	}
	return _fcabd
}

// NewPdfOutline returns an initialized PdfOutline.
func NewPdfOutline() *PdfOutline {
	_bagda := &PdfOutline{_becfb: _eb.MakeIndirectObject(_eb.MakeDict())}
	_bagda._eeedb = _bagda
	return _bagda
}

// StandardImplementer is an interface that defines specified PDF standards like PDF/A-1A (pdfa.Profile1A)
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardImplementer interface {
	StandardValidator
	StandardApplier

	// StandardName gets the human-readable name of the standard.
	StandardName() string
}

func (_fcadb *PdfWriter) checkPendingObjects() {
	for _gcfad, _cadab := range _fcadb._cfggb {
		if !_fcadb.hasObject(_gcfad) {
			_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u0020\u0050\u0065n\u0064\u0069\u006eg\u0020\u006f\u0062j\u0065\u0063t\u0020\u0025\u002b\u0076\u0020\u0025T\u0020(%\u0070\u0029\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0077\u0072\u0069\u0074\u0069\u006e\u0067", _gcfad, _gcfad, _gcfad)
			for _, _dfcde := range _cadab {
				for _, _eceab := range _dfcde.Keys() {
					_dcde := _dfcde.Get(_eceab)
					if _dcde == _gcfad {
						_ddb.Log.Debug("\u0050e\u006e\u0064i\u006e\u0067\u0020\u006fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0061nd\u0020\u0072\u0065p\u006c\u0061c\u0065\u0064\u0020\u0077\u0069\u0074h\u0020\u006eu\u006c\u006c")
						_dfcde.Set(_eceab, _eb.MakeNull())
						break
					}
				}
			}
		}
	}
}

// SetType sets the field button's type.  Can be one of:
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
// This sets the field's flag appropriately.
func (_gggb *PdfFieldButton) SetType(btype ButtonType) {
	_afcd := uint32(0)
	if _gggb.Ff != nil {
		_afcd = uint32(*_gggb.Ff)
	}
	switch btype {
	case ButtonTypePush:
		_afcd |= FieldFlagPushbutton.Mask()
	case ButtonTypeRadio:
		_afcd |= FieldFlagRadio.Mask()
	}
	_gggb.Ff = _eb.MakeInteger(int64(_afcd))
}

// ColorFromPdfObjects loads the color from PDF objects.
// The first objects (if present) represent the color in underlying colorspace.  The last one represents
// the name of the pattern.
func (_bdab *PdfColorspaceSpecialPattern) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) < 1 {
		return nil, _dcf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_aedg := &PdfColorPattern{}
	_gfec, _gddab := objects[len(objects)-1].(*_eb.PdfObjectName)
	if !_gddab {
		_ddb.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020a\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", objects[len(objects)-1])
		return nil, ErrTypeCheck
	}
	_aedg.PatternName = *_gfec
	if len(objects) > 1 {
		_fgec := objects[0 : len(objects)-1]
		if _bdab.UnderlyingCS == nil {
			_ddb.Log.Debug("P\u0061\u0074t\u0065\u0072\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0077\u0069\u0074\u0068\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u0020\u0062\u0075\u0074\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0063\u0073\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
			return nil, _dcf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
		}
		_efbb, _ecga := _bdab.UnderlyingCS.ColorFromPdfObjects(_fgec)
		if _ecga != nil {
			_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0076\u0069\u0061\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0063\u0073\u003a\u0020\u0025\u0076", _ecga)
			return nil, _ecga
		}
		_aedg.Color = _efbb
	}
	return _aedg, nil
}

// PdfAnnotationLink represents Link annotations.
// (Section 12.5.6.5 p. 403).
type PdfAnnotationLink struct {
	*PdfAnnotation
	A          _eb.PdfObject
	Dest       _eb.PdfObject
	H          _eb.PdfObject
	PA         _eb.PdfObject
	QuadPoints _eb.PdfObject
	BS         _eb.PdfObject
	_fbg       *PdfAction
	_aced      *PdfReader
}

// GetPageIndirectObject returns the indirect object of page for the specified page number.
func (_fgggab *PdfWriter) GetPageIndirectObject(pageNum int) (*_eb.PdfIndirectObject, error) {
	if pageNum < 0 || pageNum >= len(_fgggab._aadfg) {
		return nil, _dcf.New("\u0070a\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	return _fgggab._aadfg[pageNum], nil
}
func (_bgf *PdfReader) newPdfAnnotationPolyLineFromDict(_fgf *_eb.PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	_deef := PdfAnnotationPolyLine{}
	_caff, _aaf := _bgf.newPdfAnnotationMarkupFromDict(_fgf)
	if _aaf != nil {
		return nil, _aaf
	}
	_deef.PdfAnnotationMarkup = _caff
	_deef.Vertices = _fgf.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_deef.LE = _fgf.Get("\u004c\u0045")
	_deef.BS = _fgf.Get("\u0042\u0053")
	_deef.IC = _fgf.Get("\u0049\u0043")
	_deef.BE = _fgf.Get("\u0042\u0045")
	_deef.IT = _fgf.Get("\u0049\u0054")
	_deef.Measure = _fgf.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_deef, nil
}

// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
func (_fdac *PdfColorspaceDeviceN) ImageToRGB(img Image) (Image, error) {
	_gfecb := _bb.NewReader(img.getBase())
	_geda := _df.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, nil, img._bdcab, img._fedc)
	_ceef := _bb.NewWriter(_geda)
	_cggb := _gg.Pow(2, float64(img.BitsPerComponent)) - 1
	_cbec := _fdac.GetNumComponents()
	_babdc := make([]uint32, _cbec)
	_bgdb := make([]float64, _cbec)
	for {
		_ddfef := _gfecb.ReadSamples(_babdc)
		if _ddfef == _bagf.EOF {
			break
		} else if _ddfef != nil {
			return img, _ddfef
		}
		for _gdfc := 0; _gdfc < _cbec; _gdfc++ {
			_bgdba := float64(_babdc[_gdfc]) / _cggb
			_bgdb[_gdfc] = _bgdba
		}
		_bdfc, _ddfef := _fdac.TintTransform.Evaluate(_bgdb)
		if _ddfef != nil {
			return img, _ddfef
		}
		for _, _fdbg := range _bdfc {
			_fdbg = _gg.Min(_gg.Max(0, _fdbg), 1.0)
			if _ddfef = _ceef.WriteSample(uint32(_fdbg * _cggb)); _ddfef != nil {
				return img, _ddfef
			}
		}
	}
	return _fdac.AlternateSpace.ImageToRGB(_ggaa(&_geda))
}
func (_eegace *LTV) generateVRIKey(_aaagfc *PdfSignature) (string, error) {
	_cfbc, _cagd := _bceda(_aaagfc.Contents.Bytes())
	if _cagd != nil {
		return "", _cagd
	}
	return _cc.ToUpper(_egf.EncodeToString(_cfbc)), nil
}

// NewPdfActionThread returns a new "thread" action.
func NewPdfActionThread() *PdfActionThread {
	_ded := NewPdfAction()
	_fd := &PdfActionThread{}
	_fd.PdfAction = _ded
	_ded.SetContext(_fd)
	return _fd
}

// Encoder returns the font's text encoder.
func (_bfbe pdfCIDFontType0) Encoder() _fc.TextEncoder { return _bfbe._ebgbc }
func _gaeaca(_ceeee string) (string, error) {
	var _efgcc _dd.Buffer
	_efgcc.WriteString(_ceeee)
	_ffadd := make([]byte, 8+16)
	_bccbb := _d.Now().UTC().UnixNano()
	_ef.BigEndian.PutUint64(_ffadd, uint64(_bccbb))
	_, _ccbaf := _gaaa.Read(_ffadd[8:])
	if _ccbaf != nil {
		return "", _ccbaf
	}
	_efgcc.WriteString(_egf.EncodeToString(_ffadd))
	return _efgcc.String(), nil
}
func _fdce(_geaa _eb.PdfObject, _ffaf bool) (*PdfFont, error) {
	_fgded, _eefeg, _fdfb := _cbdd(_geaa)
	if _fgded != nil {
		_abaa(_fgded)
	}
	if _fdfb != nil {
		if _fdfb == ErrType1CFontNotSupported {
			_dfea, _bfddb := _bfbdg(_fgded, _eefeg, nil)
			if _bfddb != nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _eefeg, _bfddb)
				return nil, _fdfb
			}
			return &PdfFont{_fdaa: _dfea}, _fdfb
		}
		return nil, _fdfb
	}
	_caab := &PdfFont{}
	switch _eefeg._fgdee {
	case "\u0054\u0079\u0070e\u0030":
		if !_ffaf {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u004c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u00650\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _eefeg)
			return nil, _dcf.New("\u0063\u0079\u0063\u006cic\u0061\u006c\u0020\u0074\u0079\u0070\u0065\u0030\u0020\u006c\u006f\u0061\u0064\u0069n\u0067")
		}
		_dfccg, _cdgea := _affe(_fgded, _eefeg)
		if _cdgea != nil {
			_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0057\u0068\u0069l\u0065\u0020\u006c\u006f\u0061\u0064\u0069ng\u0020\u0054\u0079\u0070e\u0030\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066on\u0074\u003d%\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eefeg, _cdgea)
			return nil, _cdgea
		}
		_caab._fdaa = _dfccg
	case "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		var _dgcag *pdfFontSimple
		_egdfc, _abfga := _fg.NewStdFontByName(_fg.StdFontName(_eefeg._agcc))
		if _abfga {
			_ebcbcb := _fbfae(_egdfc)
			_caab._fdaa = &_ebcbcb
			_ebccde := _eb.TraceToDirectObject(_ebcbcb.ToPdfObject())
			_gbdb, _babb, _bcfc := _cbdd(_ebccde)
			if _bcfc != nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _eefeg, _ebcbcb)
				return nil, _bcfc
			}
			for _, _gfaff := range _fgded.Keys() {
				_gbdb.Set(_gfaff, _fgded.Get(_gfaff))
			}
			_dgcag, _bcfc = _bfbdg(_gbdb, _babb, _ebcbcb._gcgb)
			if _bcfc != nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _eefeg, _ebcbcb)
				return nil, _bcfc
			}
			_dgcag._cegda = _ebcbcb._cegda
			_dgcag._gacdb = _ebcbcb._gacdb
			if _dgcag._debdb == nil {
				_dgcag._debdb = _ebcbcb._debdb
			}
		} else {
			_dgcag, _fdfb = _bfbdg(_fgded, _eefeg, nil)
			if _fdfb != nil {
				_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _eefeg, _fdfb)
				return nil, _fdfb
			}
		}
		_fdfb = _dgcag.addEncoding()
		if _fdfb != nil {
			return nil, _fdfb
		}
		if _abfga {
			_dgcag.updateStandard14Font()
		}
		if _abfga && _dgcag._eccbb == nil && _dgcag._gcgb == nil {
			_ddb.Log.Error("\u0073\u0069\u006d\u0070\u006c\u0065\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _dgcag)
			_ddb.Log.Error("\u0066n\u0074\u003d\u0025\u002b\u0076", _egdfc)
		}
		if len(_dgcag._cegda) == 0 {
			_ddb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u006f\u0020\u0077\u0069d\u0074h\u0073.\u0020\u0066\u006f\u006e\u0074\u003d\u0025s", _dgcag)
		}
		_caab._fdaa = _dgcag
	case "\u0054\u0079\u0070e\u0033":
		_gggd, _aagb := _fafe(_fgded, _eefeg)
		if _aagb != nil {
			_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020W\u0068\u0069\u006c\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0074y\u0070\u0065\u0033\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _aagb)
			return nil, _aagb
		}
		_caab._fdaa = _gggd
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030":
		_deage, _bcgf := _dcadb(_fgded, _eefeg)
		if _bcgf != nil {
			_ddb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0057\u0068i\u006c\u0065\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069d \u0066\u006f\u006et\u0020\u0074y\u0070\u0065\u0030\u0020\u0066\u006fn\u0074\u003a \u0025\u0076", _bcgf)
			return nil, _bcgf
		}
		_caab._fdaa = _deage
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_cece, _cdeaf := _fbcef(_fgded, _eefeg)
		if _cdeaf != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0057\u0068\u0069l\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069\u0064\u0020f\u006f\u006e\u0074\u0020\u0074yp\u0065\u0032\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eefeg, _cdeaf)
			return nil, _cdeaf
		}
		_caab._fdaa = _cece
	default:
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073", _eefeg)
		return nil, _e.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0066\u006f\u006e\u0074\u0020\u0074y\u0070\u0065\u003a\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _eefeg)
	}
	return _caab, nil
}

// ToPdfObject implements interface PdfModel.
func (_abe *PdfAction) ToPdfObject() _eb.PdfObject {
	_fe := _abe._dee
	_gb := _fe.PdfObject.(*_eb.PdfObjectDictionary)
	_gb.Clear()
	_gb.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0041\u0063\u0074\u0069\u006f\u006e"))
	_gb.SetIfNotNil("\u0053", _abe.S)
	_gb.SetIfNotNil("\u004e\u0065\u0078\u0074", _abe.Next)
	return _fe
}

// FieldImageProvider provides fields images for specified fields.
type FieldImageProvider interface {
	FieldImageValues() (map[string]*Image, error)
}

// PdfAnnotationWidget represents Widget annotations.
// Note: Widget annotations are used to display form fields.
// (Section 12.5.6.19).
type PdfAnnotationWidget struct {
	*PdfAnnotation
	H      _eb.PdfObject
	MK     _eb.PdfObject
	A      _eb.PdfObject
	AA     _eb.PdfObject
	BS     _eb.PdfObject
	Parent _eb.PdfObject
	_bca   *PdfField
	_eccd  bool
}

func (_ggge *PdfReader) newPdfFieldFromIndirectObject(_babaa *_eb.PdfIndirectObject, _ffeg *PdfField) (*PdfField, error) {
	if _cagc, _dabe := _ggge._affaf.GetModelFromPrimitive(_babaa).(*PdfField); _dabe {
		return _cagc, nil
	}
	_fdec, _gcbf := _eb.GetDict(_babaa)
	if !_gcbf {
		return nil, _e.Errorf("\u0050\u0064f\u0046\u0069\u0065\u006c\u0064 \u0069\u006e\u0064\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_cgbd := NewPdfField()
	_cgbd._adgda = _babaa
	_cgbd._adgda.PdfObject = _fdec
	if _bbda, _ccaf := _eb.GetName(_fdec.Get("\u0046\u0054")); _ccaf {
		_cgbd.FT = _bbda
	}
	if _ffeg != nil {
		_cgbd.Parent = _ffeg
	}
	_cgbd.T, _ = _fdec.Get("\u0054").(*_eb.PdfObjectString)
	_cgbd.TU, _ = _fdec.Get("\u0054\u0055").(*_eb.PdfObjectString)
	_cgbd.TM, _ = _fdec.Get("\u0054\u004d").(*_eb.PdfObjectString)
	_cgbd.Ff, _ = _fdec.Get("\u0046\u0066").(*_eb.PdfObjectInteger)
	_cgbd.V = _fdec.Get("\u0056")
	_cgbd.DV = _fdec.Get("\u0044\u0056")
	_cgbd.AA = _fdec.Get("\u0041\u0041")
	if DA := _fdec.Get("\u0044\u0041"); DA != nil {
		DA, _ := _eb.GetString(DA)
		_cgbd.VariableText = &VariableText{DA: DA}
		Q, _ := _fdec.Get("\u0051").(*_eb.PdfObjectInteger)
		DS, _ := _fdec.Get("\u0044\u0053").(*_eb.PdfObjectString)
		RV := _fdec.Get("\u0052\u0056")
		_cgbd.VariableText.Q = Q
		_cgbd.VariableText.DS = DS
		_cgbd.VariableText.RV = RV
	}
	_becb := _cgbd.FT
	if _becb == nil && _ffeg != nil {
		_becb = _ffeg.FT
	}
	if _becb != nil {
		switch *_becb {
		case "\u0054\u0078":
			_aecc, _feag := _cbcf(_fdec)
			if _feag != nil {
				return nil, _feag
			}
			_aecc.PdfField = _cgbd
			_cgbd._fbedg = _aecc
		case "\u0043\u0068":
			_cebd, _ddfgd := _agdf(_fdec)
			if _ddfgd != nil {
				return nil, _ddfgd
			}
			_cebd.PdfField = _cgbd
			_cgbd._fbedg = _cebd
		case "\u0042\u0074\u006e":
			_gbcb, _eeag := _gbdf(_fdec)
			if _eeag != nil {
				return nil, _eeag
			}
			_gbcb.PdfField = _cgbd
			_cgbd._fbedg = _gbcb
		case "\u0053\u0069\u0067":
			_dbdg, _dadg := _ggge.newPdfFieldSignatureFromDict(_fdec)
			if _dadg != nil {
				return nil, _dadg
			}
			_dbdg.PdfField = _cgbd
			_cgbd._fbedg = _dbdg
		default:
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", *_cgbd.FT)
			return nil, _dcf.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079p\u0065")
		}
	}
	if _aafea, _bceef := _eb.GetName(_fdec.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bceef {
		if *_aafea == "\u0057\u0069\u0064\u0067\u0065\u0074" {
			_fbbe, _eebb := _ggge.newPdfAnnotationFromIndirectObject(_babaa)
			if _eebb != nil {
				return nil, _eebb
			}
			_bfea, _aadd := _fbbe.GetContext().(*PdfAnnotationWidget)
			if !_aadd {
				return nil, _dcf.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0067e\u0074 \u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006fn")
			}
			_bfea._bca = _cgbd
			_bfea.Parent = _cgbd._adgda
			_cgbd.Annotations = append(_cgbd.Annotations, _bfea)
			return _cgbd, nil
		}
	}
	_fcdeg := true
	if _dbafc, _dafc := _eb.GetArray(_fdec.Get("\u004b\u0069\u0064\u0073")); _dafc {
		_gffca := make([]*_eb.PdfIndirectObject, 0, _dbafc.Len())
		for _, _ebbc := range _dbafc.Elements() {
			_affda, _cdea := _eb.GetIndirect(_ebbc)
			if !_cdea {
				_gccd, _cbfe := _eb.GetStream(_ebbc)
				if _cbfe && _gccd.PdfObjectDictionary != nil {
					_aefg, _fefc := _eb.GetNameVal(_gccd.Get("\u0054\u0079\u0070\u0065"))
					if _fefc && _aefg == "\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061" {
						_ddb.Log.Debug("E\u0052RO\u0052:\u0020f\u006f\u0072\u006d\u0020\u0066i\u0065\u006c\u0064 \u004b\u0069\u0064\u0073\u0020a\u0072\u0072\u0061y\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0069n\u0076\u0061\u006cid \u004d\u0065\u0074\u0061\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e")
						continue
					}
				}
				return nil, _dcf.New("n\u006f\u0074\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0028\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065\u006cd\u0029")
			}
			_cdabe, _eaae := _eb.GetDict(_affda)
			if !_eaae {
				return nil, ErrTypeCheck
			}
			if _fcdeg {
				_fcdeg = !_afcc(_cdabe)
			}
			_gffca = append(_gffca, _affda)
		}
		for _, _gbe := range _gffca {
			if _fcdeg {
				_aecf, _dfed := _ggge.newPdfAnnotationFromIndirectObject(_gbe)
				if _dfed != nil {
					_ddb.Log.Debug("\u0045r\u0072\u006fr\u0020\u006c\u006fa\u0064\u0069\u006e\u0067\u0020\u0077\u0069d\u0067\u0065\u0074\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0072 \u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _dfed)
					return nil, _dfed
				}
				_dfab, _ebgff := _aecf._cdb.(*PdfAnnotationWidget)
				if !_ebgff {
					return nil, ErrTypeCheck
				}
				_dfab._bca = _cgbd
				_cgbd.Annotations = append(_cgbd.Annotations, _dfab)
			} else {
				_afgg, _ffdf := _ggge.newPdfFieldFromIndirectObject(_gbe, _cgbd)
				if _ffdf != nil {
					_ddb.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0068\u0069\u006c\u0064\u0020\u0066\u0069\u0065\u006c\u0064: \u0025\u0076", _ffdf)
					return nil, _ffdf
				}
				_cgbd.Kids = append(_cgbd.Kids, _afgg)
			}
		}
	}
	return _cgbd, nil
}

// ViewerPreferences represents the viewer preferences of a PDF document.
type ViewerPreferences struct {
	_caafc  *bool
	_gcbcb  *bool
	_fcebae *bool
	_edbaa  *bool
	_dacc   *bool
	_bgdge  *bool
	_cdcff  NonFullScreenPageMode
	_baeb   Direction
	_ffgda  PageBoundary
	_eadae  PageBoundary
	_bdegd  PageBoundary
	_edef   PageBoundary
	_gefcg  PrintScaling
	_ebabd  Duplex
	_gggea  *bool
	_gcgfc  []int
	_bdeec  int
}

const (
	FieldFlagClear             FieldFlag = 0
	FieldFlagReadOnly          FieldFlag = 1
	FieldFlagRequired          FieldFlag = (1 << 1)
	FieldFlagNoExport          FieldFlag = (2 << 1)
	FieldFlagNoToggleToOff     FieldFlag = (1 << 14)
	FieldFlagRadio             FieldFlag = (1 << 15)
	FieldFlagPushbutton        FieldFlag = (1 << 16)
	FieldFlagRadiosInUnision   FieldFlag = (1 << 25)
	FieldFlagMultiline         FieldFlag = (1 << 12)
	FieldFlagPassword          FieldFlag = (1 << 13)
	FieldFlagFileSelect        FieldFlag = (1 << 20)
	FieldFlagDoNotScroll       FieldFlag = (1 << 23)
	FieldFlagComb              FieldFlag = (1 << 24)
	FieldFlagRichText          FieldFlag = (1 << 26)
	FieldFlagDoNotSpellCheck   FieldFlag = (1 << 22)
	FieldFlagCombo             FieldFlag = (1 << 17)
	FieldFlagEdit              FieldFlag = (1 << 18)
	FieldFlagSort              FieldFlag = (1 << 19)
	FieldFlagMultiSelect       FieldFlag = (1 << 21)
	FieldFlagCommitOnSelChange FieldFlag = (1 << 27)
)

// SetFileName sets the pdf writer file name for metered usage tracker.
func (_afcgd *PdfWriter) SetFileName(name string) { _afcgd._adeee = name }

// DecodeArray returns the range of color component values in DeviceGray colorspace.
func (_bbcb *PdfColorspaceDeviceGray) DecodeArray() []float64 { return []float64{0, 1.0} }

// GetKDict returns the KDict of the KValue.
func (_eegbd *KValue) GetKDict() *KDict { return _eegbd._ccbca }

// NewPdfColorLab returns a new Lab color.
func NewPdfColorLab(l, a, b float64) *PdfColorLab { _gaad := PdfColorLab{l, a, b}; return &_gaad }

// SetFlag sets the flag for the field.
func (_bageg *PdfField) SetFlag(flag FieldFlag) { _bageg.Ff = _eb.MakeInteger(int64(flag)) }

// Subtype returns the font's "Subtype" field.
func (_fceab *PdfFont) Subtype() string {
	_fadfd := _fceab.baseFields()._fgdee
	if _fbbf, _fceae := _fceab._fdaa.(*pdfFontType0); _fceae {
		_fadfd = _fadfd + "\u003a" + _fbbf.DescendantFont.Subtype()
	}
	return _fadfd
}

// NewPdfAnnotationPolyLine returns a new polyline annotation.
func NewPdfAnnotationPolyLine() *PdfAnnotationPolyLine {
	_fdf := NewPdfAnnotation()
	_afef := &PdfAnnotationPolyLine{}
	_afef.PdfAnnotation = _fdf
	_afef.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdf.SetContext(_afef)
	return _afef
}

// SetDocInfo set document info.
// This will overwrite any globally declared document info.
func (_affbfb *PdfWriter) SetDocInfo(info *PdfInfo) { _affbfb.setDocInfo(info.ToPdfObject()) }

// RunesToCharcodeBytes maps the provided runes to charcode bytes and it
// returns the resulting slice of bytes, along with the number of runes which
// could not be converted. If the number of misses is 0, all runes were
// successfully converted.
func (_fgedg *PdfFont) RunesToCharcodeBytes(data []rune) ([]byte, int) {
	var _cadfc []_fc.TextEncoder
	var _adeaf _fc.CMapEncoder
	if _ebda := _fgedg.baseFields()._bgbg; _ebda != nil {
		_adeaf = _fc.NewCMapEncoder("", nil, _ebda)
	}
	_bdege := _fgedg.Encoder()
	if _bdege != nil {
		switch _fgeb := _bdege.(type) {
		case _fc.SimpleEncoder:
			_abbe := _fgeb.BaseName()
			if _, _ebeec := _bcaad[_abbe]; _ebeec {
				_cadfc = append(_cadfc, _bdege)
			}
		}
	}
	if len(_cadfc) == 0 {
		if _fgedg.baseFields()._bgbg != nil {
			_cadfc = append(_cadfc, _adeaf)
		}
		if _bdege != nil {
			_cadfc = append(_cadfc, _bdege)
		}
	}
	var _bccg _dd.Buffer
	var _eagff int
	for _, _agbd := range data {
		var _cafa bool
		for _, _cdgec := range _cadfc {
			if _bdfff := _cdgec.Encode(string(_agbd)); len(_bdfff) > 0 {
				_bccg.Write(_bdfff)
				_cafa = true
				break
			}
		}
		if !_cafa {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020f\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u006d\u0061\u0070\u0020\u0072\u0075\u006e\u0065\u0020\u0060\u0025\u002b\u0071\u0060\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065", _agbd)
			_eagff++
		}
	}
	if _eagff != 0 {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0061\u006cl\u0020\u0072\u0075\u006e\u0065\u0073\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072c\u006fd\u0065\u0073\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0052\u0075\u006e\u0065\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073 \u0065\u006e\u0063\u006f\u0064e\u0072\u0073\u003d\u0025\u002b\u0076", len(data), _eagff, _fgedg, _cadfc)
	}
	return _bccg.Bytes(), _eagff
}

// Read reads an image and loads into a new Image object with an RGB
// colormap and 8 bits per component.
func (_ecgge DefaultImageHandler) Read(reader _bagf.Reader) (*Image, error) {
	_fgddf, _, _edbga := _fb.Decode(reader)
	if _edbga != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _edbga)
		return nil, _edbga
	}
	return _ecgge.NewImageFromGoImage(_fgddf)
}

// HasPatternByName checks whether a pattern object is defined by the specified keyName.
func (_cffge *PdfPageResources) HasPatternByName(keyName _eb.PdfObjectName) bool {
	_, _cffca := _cffge.GetPatternByName(keyName)
	return _cffca
}

// Duplex returns the value of the duplex.
func (_adfgc *ViewerPreferences) Duplex() Duplex { return _adfgc._ebabd }
func _edcec(_acgff *_eb.PdfObjectDictionary) (*PdfShadingType5, error) {
	_eabab := PdfShadingType5{}
	_ecacg := _acgff.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _ecacg == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_afaa, _dafgd := _ecacg.(*_eb.PdfObjectInteger)
	if !_dafgd {
		_ddb.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ecacg)
		return nil, _eb.ErrTypeError
	}
	_eabab.BitsPerCoordinate = _afaa
	_ecacg = _acgff.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ecacg == nil {
		_ddb.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_afaa, _dafgd = _ecacg.(*_eb.PdfObjectInteger)
	if !_dafgd {
		_ddb.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _ecacg)
		return nil, _eb.ErrTypeError
	}
	_eabab.BitsPerComponent = _afaa
	_ecacg = _acgff.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077")
	if _ecacg == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071u\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0056\u0065\u0072\u0074\u0069c\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077")
		return nil, ErrRequiredAttributeMissing
	}
	_afaa, _dafgd = _ecacg.(*_eb.PdfObjectInteger)
	if !_dafgd {
		_ddb.Log.Debug("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006et\u0065\u0067\u0065\u0072\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _ecacg)
		return nil, _eb.ErrTypeError
	}
	_eabab.VerticesPerRow = _afaa
	_ecacg = _acgff.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _ecacg == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fgcff, _dafgd := _ecacg.(*_eb.PdfObjectArray)
	if !_dafgd {
		_ddb.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ecacg)
		return nil, _eb.ErrTypeError
	}
	_eabab.Decode = _fgcff
	if _gaeg := _acgff.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _gaeg != nil {
		_eabab.Function = []PdfFunction{}
		if _dfbaf, _ffeef := _gaeg.(*_eb.PdfObjectArray); _ffeef {
			for _, _gbebd := range _dfbaf.Elements() {
				_bgdcf, _ffgfe := _cccfa(_gbebd)
				if _ffgfe != nil {
					_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ffgfe)
					return nil, _ffgfe
				}
				_eabab.Function = append(_eabab.Function, _bgdcf)
			}
		} else {
			_baffb, _bdgf := _cccfa(_gaeg)
			if _bdgf != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bdgf)
				return nil, _bdgf
			}
			_eabab.Function = append(_eabab.Function, _baffb)
		}
	}
	return &_eabab, nil
}

// Insert adds a top level outline item in the outline,
// at the specified index.
func (_cbcab *Outline) Insert(index uint, item *OutlineItem) {
	_aebef := uint(len(_cbcab.Entries))
	if index > _aebef {
		index = _aebef
	}
	_cbcab.Entries = append(_cbcab.Entries[:index], append([]*OutlineItem{item}, _cbcab.Entries[index:]...)...)
}
func (_ccgfc *PdfPage) flattenFieldsWithOpts(_afeb FieldAppearanceGenerator, _deggb *FieldFlattenOpts, _ccaba map[*PdfAnnotation]bool) error {
	var _gcga []*PdfAnnotation
	if _afeb != nil {
		if _bceg := _afeb.WrapContentStream(_ccgfc); _bceg != nil {
			return _bceg
		}
	}
	_bcecb, _fgffb := _ccgfc.GetAnnotations()
	if _fgffb != nil {
		return _fgffb
	}
	for _, _efge := range _bcecb {
		_bgacd, _aeceg := _ccaba[_efge]
		if !_aeceg && _deggb.AnnotFilterFunc != nil {
			if _, _gfcedd := _efge.GetContext().(*PdfAnnotationWidget); !_gfcedd {
				_aeceg = _deggb.AnnotFilterFunc(_efge)
			}
		}
		if !_aeceg {
			_gcga = append(_gcga, _efge)
			continue
		}
		switch _efge.GetContext().(type) {
		case *PdfAnnotationPopup:
			continue
		case *PdfAnnotationLink:
			continue
		case *PdfAnnotationProjection:
			continue
		}
		_ddef, _gdbd, _ceffe := _cffg(_efge)
		if _ceffe != nil {
			if !_bgacd {
				_ddb.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0056\u0020\u002d\u003e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075t\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065am\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072")
				continue
			}
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d,\u0020\u0065\u0072\u0072\u0020\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006bi\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072", _ceffe)
			continue
		}
		if _ddef == nil {
			continue
		}
		_acge := _ccgfc.Resources.GenerateXObjectName()
		_ccgfc.Resources.SetXObjectFormByName(_acge, _ddef)
		_adacb, _bebd, _ceffe := _gcbfd(_ddef)
		if _ceffe != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0061\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u004d\u0061\u0074\u0072\u0069\u0078\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0078\u0066\u006f\u0072\u006d\u0020\u0062\u0062\u006f\u0078\u0020\u0061\u0064\u006a\u0075\u0073t\u006d\u0065\u006e\u0074\u003a \u0025\u0076", _ceffe)
		} else {
			_deeff := _ffg.IdentityMatrix()
			_deeff = _deeff.Translate(-_adacb.Llx, -_adacb.Lly)
			if _bebd {
				_gcdf := 0.0
				if _adacb.Width() > 0 {
					_gcdf = _gdbd.Width() / _adacb.Width()
				}
				_cccd := 0.0
				if _adacb.Height() > 0 {
					_cccd = _gdbd.Height() / _adacb.Height()
				}
				_deeff = _deeff.Scale(_gcdf, _cccd)
			}
			_gdbd.Transform(_deeff)
		}
		_ebabg := _gg.Min(_gdbd.Llx, _gdbd.Urx)
		_cceg := _gg.Min(_gdbd.Lly, _gdbd.Ury)
		var _bdccf []string
		_bdccf = append(_bdccf, "\u0071")
		_bdccf = append(_bdccf, _e.Sprintf("\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020%\u002e6\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0063\u006d", 1.0, 0.0, 0.0, 1.0, _ebabg, _cceg))
		_bdccf = append(_bdccf, _e.Sprintf("\u002f\u0025\u0073\u0020\u0044\u006f", _acge.String()))
		_bdccf = append(_bdccf, "\u0051")
		_eegae := _cc.Join(_bdccf, "\u000a")
		_ceffe = _ccgfc.AppendContentStream(_eegae)
		if _ceffe != nil {
			return _ceffe
		}
		if _ddef.Resources != nil {
			_ccdg, _dcdd := _eb.GetDict(_ddef.Resources.Font)
			if _dcdd {
				for _, _cbba := range _ccdg.Keys() {
					if !_ccgfc.Resources.HasFontByName(_cbba) {
						_ccgfc.Resources.SetFontByName(_cbba, _ccdg.Get(_cbba))
					}
				}
			}
		}
	}
	if len(_gcga) > 0 {
		_ccgfc._dbga = _gcga
	} else {
		_ccgfc._dbga = []*PdfAnnotation{}
	}
	return nil
}

// ColorToRGB only converts color used with uncolored patterns (defined in underlying colorspace).  Does not go into the
// pattern objects and convert those.  If that is desired, needs to be done separately.  See for example
// grayscale conversion example in unidoc-examples repo.
func (_gbba *PdfColorspaceSpecialPattern) ColorToRGB(color PdfColor) (PdfColor, error) {
	_gaff, _ceggb := color.(*PdfColorPattern)
	if !_ceggb {
		_ddb.Log.Debug("\u0043\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0070a\u0074\u0074\u0065\u0072\u006e\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", color)
		return nil, ErrTypeCheck
	}
	if _gaff.Color == nil {
		return color, nil
	}
	if _gbba.UnderlyingCS == nil {
		return nil, _dcf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	return _gbba.UnderlyingCS.ColorToRGB(_gaff.Color)
}

// NewPermissions returns a new permissions object.
func NewPermissions(docMdp *PdfSignature) *Permissions {
	_aadaa := Permissions{}
	_aadaa.DocMDP = docMdp
	_ebdeec := _eb.MakeDict()
	_ebdeec.Set("\u0044\u006f\u0063\u004d\u0044\u0050", docMdp.ToPdfObject())
	_aadaa._cgbag = _ebdeec
	return &_aadaa
}

// PdfShadingType1 is a Function-based shading.
type PdfShadingType1 struct {
	*PdfShading
	Domain   *_eb.PdfObjectArray
	Matrix   *_eb.PdfObjectArray
	Function []PdfFunction
}

// ToPdfObject implements interface PdfModel.
func (_egfg *PdfAnnotationRedact) ToPdfObject() _eb.PdfObject {
	_egfg.PdfAnnotation.ToPdfObject()
	_cdgdg := _egfg._ggf
	_beg := _cdgdg.PdfObject.(*_eb.PdfObjectDictionary)
	_egfg.PdfAnnotationMarkup.appendToPdfDictionary(_beg)
	_beg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0052\u0065\u0064\u0061\u0063\u0074"))
	_beg.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _egfg.QuadPoints)
	_beg.SetIfNotNil("\u0049\u0043", _egfg.IC)
	_beg.SetIfNotNil("\u0052\u004f", _egfg.RO)
	_beg.SetIfNotNil("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074", _egfg.OverlayText)
	_beg.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _egfg.Repeat)
	_beg.SetIfNotNil("\u0044\u0041", _egfg.DA)
	_beg.SetIfNotNil("\u0051", _egfg.Q)
	return _cdgdg
}
func (_bdge *PdfReader) newPdfPageFromDict(_gdaag *_eb.PdfObjectDictionary) (*PdfPage, error) {
	_cace := NewPdfPage()
	_cace._aaagaa = _gdaag
	_cace._eaebc = *_gdaag
	_acgdb := *_gdaag
	_ebagb, _ggabbb := _acgdb.Get("\u0054\u0079\u0070\u0065").(*_eb.PdfObjectName)
	if !_ggabbb {
		return nil, _dcf.New("\u006d\u0069ss\u0069\u006e\u0067/\u0069\u006e\u0076\u0061lid\u0020Pa\u0067\u0065\u0020\u0064\u0069\u0063\u0074io\u006e\u0061\u0072\u0079\u0020\u0054\u0079p\u0065")
	}
	if *_ebagb != "\u0050\u0061\u0067\u0065" {
		return nil, _dcf.New("\u0070\u0061\u0067\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u0050a\u0067\u0065")
	}
	if _accfgc := _acgdb.Get("\u0050\u0061\u0072\u0065\u006e\u0074"); _accfgc != nil {
		_cace.Parent = _accfgc
	}
	if _aeeg := _acgdb.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"); _aeeg != nil {
		_bebgb, _cfbde := _eb.GetString(_aeeg)
		if !_cfbde {
			return nil, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u004c\u0061\u0073\u0074\u004d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0021=\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		_ggbdb, _acged := NewPdfDate(_bebgb.Str())
		if _acged != nil {
			return nil, _acged
		}
		_cace.LastModified = &_ggbdb
	}
	if _cggfg := _acgdb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _cggfg != nil && !_eb.IsNullObject(_cggfg) {
		_cabbg, _fagd := _eb.GetDict(_cggfg)
		if !_fagd {
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _cggfg)
		}
		var _fgadc error
		_cace.Resources, _fgadc = NewPdfPageResourcesFromDict(_cabbg)
		if _fgadc != nil {
			return nil, _fgadc
		}
	} else {
		_aeaag, _dcbfd := _cace.getParentResources()
		if _dcbfd != nil {
			return nil, _dcbfd
		}
		if _aeaag == nil {
			_aeaag = NewPdfPageResources()
		}
		_cace.Resources = _aeaag
	}
	if _cfbaa := _acgdb.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _cfbaa != nil {
		_cgge, _cada := _eb.GetArray(_cfbaa)
		if !_cada {
			return nil, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _cbbca error
		_cace.MediaBox, _cbbca = NewPdfRectangle(*_cgge)
		if _cbbca != nil {
			return nil, _cbbca
		}
	}
	if _bdda := _acgdb.Get("\u0043r\u006f\u0070\u0042\u006f\u0078"); _bdda != nil {
		_cedg, _efdag := _eb.GetArray(_bdda)
		if !_efdag {
			return nil, _dcf.New("\u0070a\u0067\u0065\u0020\u0043r\u006f\u0070\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _ebecg error
		_cace.CropBox, _ebecg = NewPdfRectangle(*_cedg)
		if _ebecg != nil {
			return nil, _ebecg
		}
	}
	if _adegf := _acgdb.Get("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"); _adegf != nil {
		_ggcde, _aaad := _eb.GetArray(_adegf)
		if !_aaad {
			return nil, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u0042\u006c\u0065\u0065\u0064\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _fabb error
		_cace.BleedBox, _fabb = NewPdfRectangle(*_ggcde)
		if _fabb != nil {
			return nil, _fabb
		}
	}
	if _fgebed := _acgdb.Get("\u0054r\u0069\u006d\u0042\u006f\u0078"); _fgebed != nil {
		_aebed, _cggdf := _eb.GetArray(_fgebed)
		if !_cggdf {
			return nil, _dcf.New("\u0070a\u0067\u0065\u0020\u0054r\u0069\u006d\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _fabcc error
		_cace.TrimBox, _fabcc = NewPdfRectangle(*_aebed)
		if _fabcc != nil {
			return nil, _fabcc
		}
	}
	if _ccbdg := _acgdb.Get("\u0041\u0072\u0074\u0042\u006f\u0078"); _ccbdg != nil {
		_dcefa, _cfdcg := _eb.GetArray(_ccbdg)
		if !_cfdcg {
			return nil, _dcf.New("\u0070a\u0067\u0065\u0020\u0041\u0072\u0074\u0042\u006f\u0078\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		var _fbdf error
		_cace.ArtBox, _fbdf = NewPdfRectangle(*_dcefa)
		if _fbdf != nil {
			return nil, _fbdf
		}
	}
	if _cebbg := _acgdb.Get("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f"); _cebbg != nil {
		_cace.BoxColorInfo = _cebbg
	}
	if _degdd := _acgdb.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _degdd != nil {
		_cace.Contents = _degdd
	}
	if _cfaab := _acgdb.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _cfaab != nil {
		_eabdg, _aaaab := _eb.GetNumberAsInt64(_cfaab)
		if _aaaab != nil {
			return nil, _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0067e\u0020\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_cace.Rotate = &_eabdg
	}
	if _gfgge := _acgdb.Get("\u0047\u0072\u006fu\u0070"); _gfgge != nil {
		_cace.Group = _gfgge
	}
	if _ceaa := _acgdb.Get("\u0054\u0068\u0075m\u0062"); _ceaa != nil {
		_cace.Thumb = _ceaa
	}
	if _fgfgc := _acgdb.Get("\u0042"); _fgfgc != nil {
		_cace.B = _fgfgc
	}
	if _cadbc := _acgdb.Get("\u0044\u0075\u0072"); _cadbc != nil {
		_cace.Dur = _cadbc
	}
	if _gagbf := _acgdb.Get("\u0054\u0072\u0061n\u0073"); _gagbf != nil {
		_cace.Trans = _gagbf
	}
	if _dfag := _acgdb.Get("\u0041\u0041"); _dfag != nil {
		_cace.AA = _dfag
	}
	if _gecc := _acgdb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _gecc != nil {
		_cace.Metadata = _gecc
	}
	if _cgdgd := _acgdb.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o"); _cgdgd != nil {
		_cace.PieceInfo = _cgdgd
	}
	if _fcdcg := _acgdb.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073"); _fcdcg != nil {
		_cace.StructParents = _fcdcg
	}
	if _dfga := _acgdb.Get("\u0049\u0044"); _dfga != nil {
		_cace.ID = _dfga
	}
	if _bbegc := _acgdb.Get("\u0050\u005a"); _bbegc != nil {
		_cace.PZ = _bbegc
	}
	if _dgabbc := _acgdb.Get("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f"); _dgabbc != nil {
		_cace.SeparationInfo = _dgabbc
	}
	if _aaga := _acgdb.Get("\u0054\u0061\u0062\u0073"); _aaga != nil {
		_cace.Tabs = _aaga
	}
	if _eefad := _acgdb.Get("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064"); _eefad != nil {
		_cace.TemplateInstantiated = _eefad
	}
	if _bcde := _acgdb.Get("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s"); _bcde != nil {
		_cace.PresSteps = _bcde
	}
	if _fcfde := _acgdb.Get("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074"); _fcfde != nil {
		_cace.UserUnit = _fcfde
	}
	if _egfde := _acgdb.Get("\u0056\u0050"); _egfde != nil {
		_cace.VP = _egfde
	}
	if _cbead := _acgdb.Get("\u0041\u006e\u006e\u006f\u0074\u0073"); _cbead != nil {
		_cace.Annots = _cbead
	}
	_cace._dcdfd = _bdge
	return _cace, nil
}

// PdfColorPatternType3 represents a color shading pattern type 3 (Radial).
type PdfColorPatternType3 struct {
	Color       PdfColor
	PatternName _eb.PdfObjectName
}

func _dfaf(_bced _eb.PdfObject) (*PdfBorderStyle, error) {
	_edab := &PdfBorderStyle{}
	_edab._gdgd = _bced
	var _dded *_eb.PdfObjectDictionary
	_bced = _eb.TraceToDirectObject(_bced)
	_dded, _abbg := _bced.(*_eb.PdfObjectDictionary)
	if !_abbg {
		return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	if _adbe := _dded.Get("\u0054\u0079\u0070\u0065"); _adbe != nil {
		_gfdc, _afgb := _adbe.(*_eb.PdfObjectName)
		if !_afgb {
			_ddb.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u0077\u0069th\u0020\u0054\u0079\u0070e\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061me\u0020\u006fb\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0054", _adbe)
		} else {
			if *_gfdc != "\u0042\u006f\u0072\u0064\u0065\u0072" {
				_ddb.Log.Debug("W\u0061\u0072\u006e\u0069\u006e\u0067,\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020B\u006f\u0072\u0064e\u0072:\u0020\u0025\u0073", *_gfdc)
			}
		}
	}
	if _ffagf := _dded.Get("\u0057"); _ffagf != nil {
		_dbg, _gcfcg := _eb.GetNumberAsFloat(_ffagf)
		if _gcfcg != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006fr \u0072\u0065\u0074\u0072\u0069\u0065\u0076\u0069\u006e\u0067\u0020\u0057\u003a\u0020%\u0076", _gcfcg)
			return nil, _gcfcg
		}
		_edab.W = &_dbg
	}
	if _ddgd := _dded.Get("\u0053"); _ddgd != nil {
		_babf, _cdfg := _ddgd.(*_eb.PdfObjectName)
		if !_cdfg {
			return nil, _dcf.New("\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u0053\u0020\u006e\u006ft\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		var _agce BorderStyle
		switch *_babf {
		case "\u0053":
			_agce = BorderStyleSolid
		case "\u0044":
			_agce = BorderStyleDashed
		case "\u0042":
			_agce = BorderStyleBeveled
		case "\u0049":
			_agce = BorderStyleInset
		case "\u0055":
			_agce = BorderStyleUnderline
		default:
			_ddb.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020s\u0074\u0079\u006c\u0065\u0020n\u0061\u006d\u0065\u0020\u0025\u0073", *_babf)
			return nil, _dcf.New("\u0073\u0074\u0079\u006ce \u0074\u0079\u0070\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065c\u006b")
		}
		_edab.S = &_agce
	}
	if _ceb := _dded.Get("\u0044"); _ceb != nil {
		_bebcb, _fcbb := _ceb.(*_eb.PdfObjectArray)
		if !_fcbb {
			_ddb.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044\u0020\u0064a\u0073\u0068\u0020\u006e\u006f\u0074\u0020a\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0054", _ceb)
			return nil, _dcf.New("\u0062o\u0072\u0064\u0065\u0072 \u0044\u0020\u0074\u0079\u0070e\u0020c\u0068e\u0063\u006b\u0020\u0065\u0072\u0072\u006fr")
		}
		_eaab, _edfd := _bebcb.ToIntegerArray()
		if _edfd != nil {
			_ddb.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044 \u0050\u0072\u006fbl\u0065\u006d\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u006e\u0074\u0065\u0067e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u003a \u0025\u0076", _edfd)
			return nil, _edfd
		}
		_edab.D = &_eaab
	}
	return _edab, nil
}
func (_cabgg *PdfWriter) setDocumentIDs(_cdgdf, _cdaeg string) {
	_cabgg._ddefd = _eb.MakeArray(_eb.MakeHexString(_cdgdf), _eb.MakeHexString(_cdaeg))
}
func _aefe(_eacb _eb.PdfObject) (*PdfPattern, error) {
	_fabf := &PdfPattern{}
	var _ffadg *_eb.PdfObjectDictionary
	if _bbgc, _ccbef := _eb.GetIndirect(_eacb); _ccbef {
		_fabf._agddd = _bbgc
		_ddcea, _gaada := _bbgc.PdfObject.(*_eb.PdfObjectDictionary)
		if !_gaada {
			_ddb.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0028g\u006f\u0074\u0020%\u0054\u0029", _bbgc.PdfObject)
			return nil, _eb.ErrTypeError
		}
		_ffadg = _ddcea
	} else if _fcfe, _gbffd := _eb.GetStream(_eacb); _gbffd {
		_fabf._agddd = _fcfe
		_ffadg = _fcfe.PdfObjectDictionary
	} else {
		_ddb.Log.Debug("\u0050a\u0074\u0074e\u0072\u006e\u0020\u006eo\u0074\u0020\u0061n\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 o\u0062\u006a\u0065c\u0074\u0020o\u0072\u0020\u0073\u0074\u0072\u0065a\u006d\u002e \u0025\u0054", _eacb)
		return nil, _eb.ErrTypeError
	}
	_dagab := _ffadg.Get("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065")
	if _dagab == nil {
		_ddb.Log.Debug("\u0050\u0064\u0066\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069n\u0067\u0020\u0050\u0061\u0074t\u0065\u0072n\u0054\u0079\u0070\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_gdbgg, _caee := _dagab.(*_eb.PdfObjectInteger)
	if !_caee {
		_ddb.Log.Debug("\u0050\u0061tt\u0065\u0072\u006e \u0074\u0079\u0070\u0065 no\u0074 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067er\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _dagab)
		return nil, _eb.ErrTypeError
	}
	if *_gdbgg != 1 && *_gdbgg != 2 {
		_ddb.Log.Debug("\u0050\u0061\u0074\u0074e\u0072\u006e\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d \u0031/\u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", *_gdbgg)
		return nil, _eb.ErrRangeError
	}
	_fabf.PatternType = int64(*_gdbgg)
	switch *_gdbgg {
	case 1:
		_gdga, _abgeb := _ddcgg(_ffadg)
		if _abgeb != nil {
			return nil, _abgeb
		}
		_gdga.PdfPattern = _fabf
		_fabf._eefgb = _gdga
		return _fabf, nil
	case 2:
		_gegc, _agfcb := _dfgag(_ffadg)
		if _agfcb != nil {
			return nil, _agfcb
		}
		_gegc.PdfPattern = _fabf
		_fabf._eefgb = _gegc
		return _fabf, nil
	}
	return nil, _dcf.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
}

// NewStandardPdfOutputIntent creates a new standard PdfOutputIntent.
func NewStandardPdfOutputIntent(outputCondition, outputConditionIdentifier, registryName string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, RegistryName: registryName, DestOutputProfile: destOutputProfile, ColorComponents: colorComponents, _ebbd: _eb.MakeDict()}
}

// ToPdfObject implements interface PdfModel.
func (_gddf *PdfAnnotationSquare) ToPdfObject() _eb.PdfObject {
	_gddf.PdfAnnotation.ToPdfObject()
	_dfac := _gddf._ggf
	_gdgf := _dfac.PdfObject.(*_eb.PdfObjectDictionary)
	if _gddf.PdfAnnotationMarkup != nil {
		_gddf.PdfAnnotationMarkup.appendToPdfDictionary(_gdgf)
	}
	_gdgf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053\u0071\u0075\u0061\u0072\u0065"))
	_gdgf.SetIfNotNil("\u0042\u0053", _gddf.BS)
	_gdgf.SetIfNotNil("\u0049\u0043", _gddf.IC)
	_gdgf.SetIfNotNil("\u0042\u0045", _gddf.BE)
	_gdgf.SetIfNotNil("\u0052\u0044", _gddf.RD)
	return _dfac
}

// SetPrintScaling sets the value of the printScaling.
func (_eagab *ViewerPreferences) SetPrintScaling(printScaling PrintScaling) {
	_eagab._gefcg = printScaling
}

// PdfPageResources is a Page resources model.
// Implements PdfModel.
type PdfPageResources struct {
	ExtGState  _eb.PdfObject
	ColorSpace _eb.PdfObject
	Pattern    _eb.PdfObject
	Shading    _eb.PdfObject
	XObject    _eb.PdfObject
	Font       _eb.PdfObject
	ProcSet    _eb.PdfObject
	Properties _eb.PdfObject
	_fdada     *_eb.PdfObjectDictionary
	_cfcff     *PdfPageResourcesColorspaces
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_fadf *PdfColorspaceSpecialSeparation) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cgfb := vals[0]
	_gaecc := []float64{_cgfb}
	_baff, _cbca := _fadf.TintTransform.Evaluate(_gaecc)
	if _cbca != nil {
		_ddb.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0065: \u0025\u0076", _cbca)
		_ddb.Log.Trace("\u0054\u0069\u006e\u0074 t\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _fadf.TintTransform)
		return nil, _cbca
	}
	_ddb.Log.Trace("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0043\u006f\u006c\u006fr\u0046\u0072\u006f\u006d\u0046\u006c\u006f\u0061\u0074\u0073\u0028\u0025\u002bv\u0029\u0020\u006f\u006e\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061te\u0053\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0023\u0076", _baff, _fadf.AlternateSpace)
	_ddcb, _cbca := _fadf.AlternateSpace.ColorFromFloats(_baff)
	if _cbca != nil {
		_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u002c\u0020\u0066a\u0069\u006c\u0065d \u0074\u006f\u0020\u0065\u0076\u0061l\u0075\u0061\u0074\u0065\u0020\u0069\u006e\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u003a \u0025\u0076", _cbca)
		return nil, _cbca
	}
	return _ddcb, nil
}

// PdfAnnotation represents an annotation in PDF (section 12.5 p. 389).
type PdfAnnotation struct {
	_cdb         PdfModel
	Rect         _eb.PdfObject
	Contents     _eb.PdfObject
	P            _eb.PdfObject
	NM           _eb.PdfObject
	M            _eb.PdfObject
	F            _eb.PdfObject
	AP           _eb.PdfObject
	AS           _eb.PdfObject
	Border       _eb.PdfObject
	C            _eb.PdfObject
	StructParent _eb.PdfObject
	OC           _eb.PdfObject
	_ggf         *_eb.PdfIndirectObject
}

// GetOutlines returns a high-level Outline object, based on the outline tree
// of the reader.
func (_begaa *PdfReader) GetOutlines() (*Outline, error) {
	if _begaa == nil {
		return nil, _dcf.New("\u0063\u0061n\u006e\u006f\u0074\u0020c\u0072\u0065a\u0074\u0065\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u006e\u0069\u006c\u0020\u0072e\u0061\u0064\u0065\u0072")
	}
	_bdfa := _begaa.GetOutlineTree()
	if _bdfa == nil {
		return nil, _dcf.New("\u0074\u0068\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u0064\u006f\u0065\u0073\u0020n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065")
	}
	var _bggdd func(_fcdf *PdfOutlineTreeNode, _bcfa *[]*OutlineItem)
	_bggdd = func(_feaea *PdfOutlineTreeNode, _dbfgg *[]*OutlineItem) {
		if _feaea == nil {
			return
		}
		if _feaea._eeedb == nil {
			_ddb.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020m\u0069\u0073\u0073\u0069ng \u006fut\u006c\u0069\u006e\u0065\u0020\u0065\u006etr\u0079\u0020\u0063\u006f\u006e\u0074\u0065x\u0074")
			return
		}
		var _aggbg *OutlineItem
		if _efeba, _gccac := _feaea._eeedb.(*PdfOutlineItem); _gccac {
			_dabfa := _efeba.Dest
			if (_dabfa == nil || _eb.IsNullObject(_dabfa)) && _efeba.A != nil {
				if _gcebe, _fgcdg := _eb.GetDict(_efeba.A); _fgcdg {
					if _cdaef, _befba := _eb.GetArray(_gcebe.Get("\u0044")); _befba {
						_dabfa = _cdaef
					} else {
						_fgce, _agabe := _eb.GetString(_gcebe.Get("\u0044"))
						if !_agabe {
							return
						}
						_cfgbe, _agabe := _begaa._bagcfd.Get("\u004e\u0061\u006de\u0073").(*_eb.PdfObjectReference)
						if !_agabe {
							return
						}
						_bbdd, _ffgb := _begaa._ebbe.LookupByReference(*_cfgbe)
						if _ffgb != nil {
							_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0028\u0025\u0073\u0029", _ffgb.Error())
							return
						}
						_adcb, _agabe := _bbdd.(*_eb.PdfIndirectObject)
						if !_agabe {
							return
						}
						_caec := map[_eb.PdfObject]struct{}{}
						_ffgb = _begaa.buildNameNodes(_adcb, _caec)
						if _ffgb != nil {
							_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006ea\u006d\u0065\u0020\u006e\u006fd\u0065\u0073 \u0028\u0025\u0073\u0029", _ffgb.Error())
							return
						}
						for _gcgcd := range _caec {
							_gfbeg, _acegd := _eb.GetDict(_gcgcd)
							if !_acegd {
								continue
							}
							_befad, _acegd := _eb.GetArray(_gfbeg.Get("\u004e\u0061\u006de\u0073"))
							if !_acegd {
								continue
							}
							for _gcgbe, _ebdgb := range _befad.Elements() {
								switch _ebdgb.(type) {
								case *_eb.PdfObjectString:
									if _ebdgb.String() == _fgce.String() {
										if _fadd := _befad.Get(_gcgbe + 1); _fadd != nil {
											if _effea, _bgef := _eb.GetDict(_fadd); _bgef {
												_dabfa = _effea.Get("\u0044")
												break
											}
										}
									}
								}
							}
						}
					}
				}
			}
			var _ddceaf OutlineDest
			if _dabfa != nil && !_eb.IsNullObject(_dabfa) {
				if _afgfa, _cdeecc := _bcbe(_dabfa, _begaa); _cdeecc == nil {
					_ddceaf = *_afgfa
				} else {
					_ddb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020p\u0061\u0072\u0073\u0065\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065\u0073\u0074\u0020\u0028\u0025\u0076\u0029\u003a\u0020\u0025\u0076", _dabfa, _cdeecc)
				}
			}
			_aggbg = NewOutlineItem(_efeba.Title.Decoded(), _ddceaf)
			*_dbfgg = append(*_dbfgg, _aggbg)
			if _efeba.Next != nil {
				_bggdd(_efeba.Next, _dbfgg)
			}
		}
		if _feaea.First != nil {
			if _aggbg != nil {
				_dbfgg = &_aggbg.Entries
			}
			_bggdd(_feaea.First, _dbfgg)
		}
	}
	_dfbdef := NewOutline()
	_bggdd(_bdfa, &_dfbdef.Entries)
	return _dfbdef, nil
}

// Inspect inspects the object types, subtypes and content in the PDF file returning a map of
// object type to number of instances of each.
func (_daafe *PdfReader) Inspect() (map[string]int, error) { return _daafe._ebbe.Inspect() }

// Encoder returns the font's text encoder.
func (_edgeg pdfFontType0) Encoder() _fc.TextEncoder { return _edgeg._edcff }

// GetCerts returns the signature certificate chain.
func (_dfcfcd *PdfSignature) GetCerts() ([]*_bag.Certificate, error) {
	var _fdaba []func() ([]*_bag.Certificate, error)
	switch _cgddg, _ := _eb.GetNameVal(_dfcfcd.SubFilter); _cgddg {
	case "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064", "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064":
		_fdaba = append(_fdaba, _dfcfcd.extractChainFromPKCS7, _dfcfcd.extractChainFromCert)
	case "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031":
		_fdaba = append(_fdaba, _dfcfcd.extractChainFromCert)
	case "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031":
		_fdaba = append(_fdaba, _dfcfcd.extractChainFromPKCS7)
	default:
		return nil, _e.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020S\u0075b\u0046i\u006c\u0074\u0065\u0072\u003a\u0020\u0025s", _cgddg)
	}
	for _, _dgdgc := range _fdaba {
		_bdafg, _baedg := _dgdgc()
		if _baedg != nil {
			return nil, _baedg
		}
		if len(_bdafg) > 0 {
			return _bdafg, nil
		}
	}
	return nil, ErrSignNoCertificates
}
func (_cbffag *PdfWriter) writeDocumentVersion() {
	if _cbffag._caed {
		_cbffag.writeString("\u000a")
	} else {
		_cbffag.writeString(_e.Sprintf("\u0025\u0025\u0050D\u0046\u002d\u0025\u0064\u002e\u0025\u0064\u000a", _cbffag._edbbf.Major, _cbffag._edbbf.Minor))
		_cbffag.writeString("\u0025\u00e2\u00e3\u00cf\u00d3\u000a")
	}
}

var _bcaad = map[string]struct{}{"\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067": {}, "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}}

// NewPdfColorspaceDeviceCMYK returns a new CMYK32 colorspace object.
func NewPdfColorspaceDeviceCMYK() *PdfColorspaceDeviceCMYK { return &PdfColorspaceDeviceCMYK{} }
func (_eabbc *PdfWriter) writeOutlines() error {
	if _eabbc._fbgge == nil {
		return nil
	}
	_ddb.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0054\u0072\u0065\u0065\u003a\u0020\u0025\u002b\u0076", _eabbc._fbgge)
	_gffbf := _eabbc._fbgge.ToPdfObject()
	_ddb.Log.Trace("\u004fu\u0074\u006c\u0069\u006e\u0065\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0025\u0054\u002c\u0020\u0070\u003a\u0025\u0070\u0029", _gffbf, _gffbf, _gffbf)
	_eabbc._dbffa.Set("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073", _gffbf)
	_gdcfbc := _eabbc.addObjects(_gffbf)
	if _gdcfbc != nil {
		return _gdcfbc
	}
	return nil
}

// ColorToRGB converts a DeviceN color to an RGB color.
func (_fdgbf *PdfColorspaceDeviceN) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _fdgbf.AlternateSpace == nil {
		return nil, _dcf.New("\u0044\u0065\u0076\u0069\u0063\u0065N\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _fdgbf.AlternateSpace.ColorToRGB(color)
}

// GetNumComponents returns the number of color components (1 for CalGray).
func (_bbabe *PdfColorCalGray) GetNumComponents() int { return 1 }
func (_ggbad *Image) samplesAddPadding(_ddbdf []uint32) []uint32 {
	_cgcaf := _df.BytesPerLine(int(_ggbad.Width), int(_ggbad.BitsPerComponent), _ggbad.ColorComponents) * (8 / int(_ggbad.BitsPerComponent))
	_gceeb := _cgcaf * int(_ggbad.Height)
	if len(_ddbdf) == _gceeb {
		return _ddbdf
	}
	_bcddg := make([]uint32, _gceeb)
	_eaeea := int(_ggbad.Width) * _ggbad.ColorComponents
	for _aadcc := 0; _aadcc < int(_ggbad.Height); _aadcc++ {
		_fdfd := _aadcc * int(_ggbad.Width)
		_aeebc := _aadcc * _cgcaf
		for _dfff := 0; _dfff < _eaeea; _dfff++ {
			_bcddg[_aeebc+_dfff] = _ddbdf[_fdfd+_dfff]
		}
	}
	return _bcddg
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_bfab *PdfColorspaceCalGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fbab := vals[0]
	if _fbab < 0.0 || _fbab > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fbab)
		return nil, ErrColorOutOfRange
	}
	_fcdc := NewPdfColorCalGray(_fbab)
	return _fcdc, nil
}

// ViewArea returns the value of the viewArea.
func (_beafe *ViewerPreferences) ViewArea() PageBoundary { return _beafe._ffgda }

// SetContentStream sets the pattern cell's content stream.
func (_cdeec *PdfTilingPattern) SetContentStream(content []byte, encoder _eb.StreamEncoder) error {
	_cgbbe, _eefgf := _cdeec._agddd.(*_eb.PdfObjectStream)
	if !_eefgf {
		_ddb.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cdeec._agddd)
		return _eb.ErrTypeError
	}
	if encoder == nil {
		encoder = _eb.NewRawEncoder()
	}
	_eefcb := _cgbbe.PdfObjectDictionary
	_ecabd := encoder.MakeStreamDict()
	_eefcb.Merge(_ecabd)
	_aeddb, _fgdbeg := encoder.EncodeBytes(content)
	if _fgdbeg != nil {
		return _fgdbeg
	}
	_eefcb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_aeddb))))
	_cgbbe.Stream = _aeddb
	return nil
}

// GetNumComponents returns the number of color components (3 for Lab).
func (_acae *PdfColorLab) GetNumComponents() int { return 3 }

// SetPdfAuthor sets the Author attribute of the output PDF.
func SetPdfAuthor(author string) { _dfbafc.Lock(); defer _dfbafc.Unlock(); _aeee = author }

// DecodeArray returns the component range values for the DeviceN colorspace.
// [0 1.0 0 1.0 ...] for each color component.
func (_aaaa *PdfColorspaceDeviceN) DecodeArray() []float64 {
	var _aaagf []float64
	for _eece := 0; _eece < _aaaa.GetNumComponents(); _eece++ {
		_aaagf = append(_aaagf, 0.0, 1.0)
	}
	return _aaagf
}

// SetDSS sets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_agee *PdfAppender) SetDSS(dss *DSS) {
	if dss != nil {
		_agee.updateObjectsDeep(dss.ToPdfObject(), nil)
	}
	_agee._cedc = dss
}

// GetExtGState gets the ExtGState specified by keyName. Returns a bool
// indicating whether it was found or not.
func (_bfcdd *PdfPageResources) GetExtGState(keyName _eb.PdfObjectName) (_eb.PdfObject, bool) {
	if _bfcdd.ExtGState == nil {
		return nil, false
	}
	_geef, _dcgba := _eb.TraceToDirectObject(_bfcdd.ExtGState).(*_eb.PdfObjectDictionary)
	if !_dcgba {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bfcdd.ExtGState)
		return nil, false
	}
	if _cgfgce := _geef.Get(keyName); _cgfgce != nil {
		return _cgfgce, true
	}
	return nil, false
}

// NewPdfAnnotationFreeText returns a new free text annotation.
func NewPdfAnnotationFreeText() *PdfAnnotationFreeText {
	_cbgf := NewPdfAnnotation()
	_fgbb := &PdfAnnotationFreeText{}
	_fgbb.PdfAnnotation = _cbgf
	_fgbb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cbgf.SetContext(_fgbb)
	return _fgbb
}

// IsValid checks if the given pdf output intent type is valid.
func (_dcecg PdfOutputIntentType) IsValid() bool {
	return _dcecg >= PdfOutputIntentTypeA1 && _dcecg <= PdfOutputIntentTypeX
}

// SetFilter sets compression filter. Decodes with current filter sets and
// encodes the data with the new filter.
func (_gdbddd *XObjectImage) SetFilter(encoder _eb.StreamEncoder) error {
	_aaedcb := _gdbddd.Stream
	_cfdca, _bfdee := _gdbddd.Filter.DecodeBytes(_aaedcb)
	if _bfdee != nil {
		return _bfdee
	}
	_gdbddd.Filter = encoder
	encoder.UpdateParams(_gdbddd.getParamsDict())
	_aaedcb, _bfdee = encoder.EncodeBytes(_cfdca)
	if _bfdee != nil {
		return _bfdee
	}
	_gdbddd.Stream = _aaedcb
	return nil
}

// ContentStreamWrapper wraps the Page's contentstream into q ... Q blocks.
type ContentStreamWrapper interface{ WrapContentStream(_dfeca *PdfPage) error }

// GetAscent returns the Ascent of the font `descriptor`.
func (_caeb *PdfFontDescriptor) GetAscent() (float64, error) {
	return _eb.GetNumberAsFloat(_caeb.Ascent)
}

// SetImage updates XObject Image with new image data.
func (_ddggg *XObjectImage) SetImage(img *Image, cs PdfColorspace) error {
	_ddggg.Filter.UpdateParams(img.GetParamsDict())
	_cacb, _cdfaa := _ddggg.Filter.EncodeBytes(img.Data)
	if _cdfaa != nil {
		return _cdfaa
	}
	_ddggg.Stream = _cacb
	_feded := img.Width
	_ddggg.Width = &_feded
	_cbbbf := img.Height
	_ddggg.Height = &_cbbbf
	_cefcbg := img.BitsPerComponent
	_ddggg.BitsPerComponent = &_cefcbg
	if cs == nil {
		if img.ColorComponents == 1 {
			_ddggg.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_ddggg.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_ddggg.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return _dcf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_ddggg.ColorSpace = cs
	}
	return nil
}

// LTV represents an LTV (Long-Term Validation) client. It is used to LTV
// enable signatures by adding validation and revocation data (certificate,
// OCSP and CRL information) to the DSS dictionary of a PDF document.
//
// LTV is added through the DSS by:
//   - Adding certificates, OCSP and CRL information in the global scope of the
//     DSS. The global data is used for validating any of the signatures present
//     in the document.
//   - Adding certificates, OCSP and CRL information for a single signature,
//     through an entry in the VRI dictionary of the DSS. The added data is used
//     for validating that particular signature only. This is the recommended
//     method for adding validation data for a signature. However, this is not
//     is not possible in the same revision the signature is applied. Validation
//     data for a signature is added based on the Contents entry of the signature,
//     which is known only after the revision is written. Even if the Contents
//     are known (e.g. when signing externally), updating the DSS at that point
//     would invalidate the calculated signature. As a result, if adding LTV
//     in the same revision is a requirement, use the first method.
//     See LTV.EnableChain.
//
// The client applies both methods, when possible.
//
// If `LTV.SkipExisting` is set to true (the default), validations are
// not added for signatures which are already present in the VRI entry of the
// document's DSS dictionary.
type LTV struct {

	// CertClient is the client used to retrieve certificates.
	CertClient *_dda.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_dda.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_dda.CRLClient

	// SkipExisting specifies whether existing signature validations
	// should be skipped.
	SkipExisting bool
	_feba        *PdfAppender
	_fadg        *DSS
}

func (_cddf *pdfFontType3) baseFields() *fontCommon { return &_cddf.fontCommon }
func (_ccgad *PdfWriter) checkLicense() error {
	_bcgfdd := _cf.GetLicenseKey()
	if (_bcgfdd == nil || !_bcgfdd.IsLicensed()) && !_fefac {
		_e.Printf("\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064\u0020c\u006f\u0070\u0079\u0020\u006f\u0066\u0020\u0055\u006e\u0069P\u0044\u0046\u000a")
		_e.Println("-\u0020\u0047\u0065\u0074\u0020\u0061\u0020\u0066\u0072e\u0065\u0020\u0074\u0072\u0069\u0061\u006c l\u0069\u0063\u0065\u006es\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070s:\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _dcf.New("\u0075\u006e\u0069\u0070d\u0066\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020c\u006fd\u0065\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0064")
	}
	return nil
}
func (_cce *PdfReader) newPdfAnnotationPolygonFromDict(_dabd *_eb.PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	_bgc := PdfAnnotationPolygon{}
	_cec, _efdc := _cce.newPdfAnnotationMarkupFromDict(_dabd)
	if _efdc != nil {
		return nil, _efdc
	}
	_bgc.PdfAnnotationMarkup = _cec
	_bgc.Vertices = _dabd.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_bgc.LE = _dabd.Get("\u004c\u0045")
	_bgc.BS = _dabd.Get("\u0042\u0053")
	_bgc.IC = _dabd.Get("\u0049\u0043")
	_bgc.BE = _dabd.Get("\u0042\u0045")
	_bgc.IT = _dabd.Get("\u0049\u0054")
	_bgc.Measure = _dabd.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_bgc, nil
}
func _dgeac() string { return _ddb.Version }

// PdfAnnotationTrapNet represents TrapNet annotations.
// (Section 12.5.6.21).
type PdfAnnotationTrapNet struct{ *PdfAnnotation }

func (_bdfbd *PdfReader) loadPerms() (*Permissions, error) {
	if _fgebb := _bdfbd._bagcfd.Get("\u0050\u0065\u0072m\u0073"); _fgebb != nil {
		if _egbge, _beaab := _eb.GetDict(_fgebb); _beaab {
			_abegb := _egbge.Get("\u0044\u006f\u0063\u004d\u0044\u0050")
			if _abegb == nil {
				return nil, nil
			}
			if _afdcdd, _ecddbd := _eb.GetIndirect(_abegb); _ecddbd {
				_gddfc, _dggdeb := _bdfbd.newPdfSignatureFromIndirect(_afdcdd)
				if _dggdeb != nil {
					return nil, _dggdeb
				}
				return NewPermissions(_gddfc), nil
			}
			return nil, _e.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u006f\u0063M\u0044\u0050\u0020\u0065nt\u0072\u0079")
		}
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0050\u0065\u0072\u006d\u0073\u0020\u0065\u006e\u0074\u0072\u0079")
	}
	return nil, nil
}

// ToPdfObject implements interface PdfModel.
func (_fge *PdfActionSound) ToPdfObject() _eb.PdfObject {
	_fge.PdfAction.ToPdfObject()
	_fbde := _fge._dee
	_fgb := _fbde.PdfObject.(*_eb.PdfObjectDictionary)
	_fgb.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeSound)))
	_fgb.SetIfNotNil("\u0053\u006f\u0075n\u0064", _fge.Sound)
	_fgb.SetIfNotNil("\u0056\u006f\u006c\u0075\u006d\u0065", _fge.Volume)
	_fgb.SetIfNotNil("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073", _fge.Synchronous)
	_fgb.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _fge.Repeat)
	_fgb.SetIfNotNil("\u004d\u0069\u0078", _fge.Mix)
	return _fbde
}

// PdfSignatureReference represents a PDF signature reference dictionary and is used for signing via form signature fields.
// (Section 12.8.1, Table 253 - Entries in a signature reference dictionary p. 469 in PDF32000_2008).
type PdfSignatureReference struct {
	_abdcb          *_eb.PdfObjectDictionary
	Type            *_eb.PdfObjectName
	TransformMethod *_eb.PdfObjectName
	TransformParams _eb.PdfObject
	Data            _eb.PdfObject
	DigestMethod    *_eb.PdfObjectName
}

var (
	ErrRequiredAttributeMissing = _dcf.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074t\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
	ErrInvalidAttribute         = _dcf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	ErrTypeCheck                = _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	_eggcg                      = _dcf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrEncrypted                = _dcf.New("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	ErrNoFont                   = _dcf.New("\u0066\u006fn\u0074\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	ErrFontNotSupported         = _db.Errorf("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u006fn\u0074\u0020\u0028\u0025\u0077\u0029", _eb.ErrNotSupported)
	ErrType1CFontNotSupported   = _db.Errorf("\u0054y\u0070\u00651\u0043\u0020\u0066o\u006e\u0074\u0073\u0020\u0061\u0072\u0065 \u006e\u006f\u0074\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _eb.ErrNotSupported)
	ErrType3FontNotSupported    = _db.Errorf("\u0054y\u0070\u00653\u0020\u0066\u006f\u006et\u0073\u0020\u0061r\u0065\u0020\u006e\u006f\u0074\u0020\u0063\u0075\u0072re\u006e\u0074\u006cy\u0020\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0028%\u0077\u0029", _eb.ErrNotSupported)
	ErrTTCmapNotSupported       = _db.Errorf("\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065d\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070e\u0020\u0063\u006d\u0061\u0070\u0020\u0066\u006f\u0072\u006d\u0061\u0074 \u0028\u0025\u0077\u0029", _eb.ErrNotSupported)
	ErrSignNotEnoughSpace       = _db.Errorf("\u0069\u006e\u0073\u0075\u0066\u0066\u0069c\u0069\u0065\u006et\u0020\u0073\u0070a\u0063\u0065 \u0061\u006c\u006c\u006f\u0063\u0061t\u0065d \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	ErrSignNoCertificates       = _db.Errorf("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065tri\u0065ve\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0063\u0068\u0061i\u006e")
)

// AppendContentStream adds content stream by string.  Appends to the last
// contentstream instance if many.
func (_deccf *PdfPage) AppendContentStream(contentStr string) error {
	_cdcdd, _gcdfe := _deccf.GetContentStreams()
	if _gcdfe != nil {
		return _gcdfe
	}
	if len(_cdcdd) == 0 {
		_cdcdd = []string{contentStr}
		return _deccf.SetContentStreams(_cdcdd, _eb.NewFlateEncoder())
	}
	var _gabb _dd.Buffer
	_gabb.WriteString(_cdcdd[len(_cdcdd)-1])
	_gabb.WriteString("\u000a")
	_gabb.WriteString(contentStr)
	_cdcdd[len(_cdcdd)-1] = _gabb.String()
	return _deccf.SetContentStreams(_cdcdd, _eb.NewFlateEncoder())
}

// SetSubtype sets the Subtype S for given PdfOutputIntent.
func (_cbeb *PdfOutputIntent) SetSubtype(subtype PdfOutputIntentType) error {
	if !subtype.IsValid() {
		return _dcf.New("\u0070\u0072o\u0076\u0069\u0064\u0065d\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u004f\u0075t\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0053\u0075b\u0054\u0079\u0070\u0065")
	}
	_cbeb.S = subtype
	return nil
}
func (_ffdae *PdfReader) newPdfAnnotationTextFromDict(_beca *_eb.PdfObjectDictionary) (*PdfAnnotationText, error) {
	_aeg := PdfAnnotationText{}
	_dbdb, _fac := _ffdae.newPdfAnnotationMarkupFromDict(_beca)
	if _fac != nil {
		return nil, _fac
	}
	_aeg.PdfAnnotationMarkup = _dbdb
	_aeg.Open = _beca.Get("\u004f\u0070\u0065\u006e")
	_aeg.Name = _beca.Get("\u004e\u0061\u006d\u0065")
	_aeg.State = _beca.Get("\u0053\u0074\u0061t\u0065")
	_aeg.StateModel = _beca.Get("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c")
	return &_aeg, nil
}

// Has checks if flag fl is set in flag and returns true if so, false otherwise.
func (_fdaca FieldFlag) Has(fl FieldFlag) bool { return (_fdaca.Mask() & fl.Mask()) > 0 }
func (_aada *PdfReader) lookupPageByObject(_aaeafa _eb.PdfObject) (*PdfPage, error) {
	return nil, _dcf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
func _ffaa(_cabd _eb.PdfObject) (*PdfColorspaceCalRGB, error) {
	_fbgc := NewPdfColorspaceCalRGB()
	if _dafe, _deeg := _cabd.(*_eb.PdfIndirectObject); _deeg {
		_fbgc._cbfd = _dafe
	}
	_cabd = _eb.TraceToDirectObject(_cabd)
	_ccbd, _edcc := _cabd.(*_eb.PdfObjectArray)
	if !_edcc {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ccbd.Len() != 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_cabd = _eb.TraceToDirectObject(_ccbd.Get(0))
	_cfga, _edcc := _cabd.(*_eb.PdfObjectName)
	if !_edcc {
		return nil, _e.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	if *_cfga != "\u0043\u0061\u006c\u0052\u0047\u0042" {
		return nil, _e.Errorf("\u006e\u006f\u0074 a\u0020\u0043\u0061\u006c\u0052\u0047\u0042\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_cabd = _eb.TraceToDirectObject(_ccbd.Get(1))
	_decb, _edcc := _cabd.(*_eb.PdfObjectDictionary)
	if !_edcc {
		return nil, _e.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	_cabd = _decb.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_cabd = _eb.TraceToDirectObject(_cabd)
	_ggfab, _edcc := _cabd.(*_eb.PdfObjectArray)
	if !_edcc {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050o\u0069\u006e\u0074")
	}
	if _ggfab.Len() != 3 {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0057h\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cdde, _bdfbeg := _ggfab.GetAsFloat64Slice()
	if _bdfbeg != nil {
		return nil, _bdfbeg
	}
	_fbgc.WhitePoint = _cdde
	_cabd = _decb.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _cabd != nil {
		_cabd = _eb.TraceToDirectObject(_cabd)
		_beag, _ccd := _cabd.(*_eb.PdfObjectArray)
		if !_ccd {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050o\u0069\u006e\u0074")
		}
		if _beag.Len() != 3 {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0042l\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		_bbe, _cfdf := _beag.GetAsFloat64Slice()
		if _cfdf != nil {
			return nil, _cfdf
		}
		_fbgc.BlackPoint = _bbe
	}
	_cabd = _decb.Get("\u0047\u0061\u006dm\u0061")
	if _cabd != nil {
		_cabd = _eb.TraceToDirectObject(_cabd)
		_bgeed, _ededd := _cabd.(*_eb.PdfObjectArray)
		if !_ededd {
			return nil, _e.Errorf("C\u0061\u006c\u0052\u0047B:\u0020I\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0061\u006d\u006d\u0061")
		}
		if _bgeed.Len() != 3 {
			return nil, _e.Errorf("C\u0061\u006c\u0052\u0047\u0042\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0047a\u006d\u006d\u0061 \u0061r\u0072\u0061\u0079")
		}
		_cage, _ageg := _bgeed.GetAsFloat64Slice()
		if _ageg != nil {
			return nil, _ageg
		}
		_fbgc.Gamma = _cage
	}
	_cabd = _decb.Get("\u004d\u0061\u0074\u0072\u0069\u0078")
	if _cabd != nil {
		_cabd = _eb.TraceToDirectObject(_cabd)
		_gca, _eaee := _cabd.(*_eb.PdfObjectArray)
		if !_eaee {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052GB\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004d\u0061\u0074\u0072i\u0078")
		}
		if _gca.Len() != 9 {
			_ddb.Log.Error("\u004d\u0061t\u0072\u0069\u0078 \u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _gca.String())
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052G\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u004da\u0074\u0072\u0069\u0078\u0020\u0061\u0072r\u0061\u0079")
		}
		_gfga, _baag := _gca.GetAsFloat64Slice()
		if _baag != nil {
			return nil, _baag
		}
		_fbgc.Matrix = _gfga
	}
	return _fbgc, nil
}

// PdfAnnotationSquiggly represents Squiggly annotations.
// (Section 12.5.6.10).
type PdfAnnotationSquiggly struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _eb.PdfObject
}

// ToInteger convert to an integer format.
func (_cgaf *PdfColorDeviceRGB) ToInteger(bits int) [3]uint32 {
	_bdegg := _gg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_bdegg * _cgaf.R()), uint32(_bdegg * _cgaf.G()), uint32(_bdegg * _cgaf.B())}
}
func (_ebgca *PdfSignature) extractChainFromCert() ([]*_bag.Certificate, error) {
	var _bafaa *_eb.PdfObjectArray
	switch _afcfb := _ebgca.Cert.(type) {
	case *_eb.PdfObjectString:
		_bafaa = _eb.MakeArray(_afcfb)
	case *_eb.PdfObjectArray:
		_bafaa = _afcfb
	default:
		return nil, _e.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _afcfb)
	}
	var _deaead _dd.Buffer
	for _, _cecff := range _bafaa.Elements() {
		_eggcc, _cafag := _eb.GetString(_cecff)
		if !_cafag {
			return nil, _e.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054", _cecff)
		}
		if _, _cgged := _deaead.Write(_eggcc.Bytes()); _cgged != nil {
			return nil, _cgged
		}
	}
	return _bag.ParseCertificates(_deaead.Bytes())
}

// ToPdfObject implements interface PdfModel.
func (_fef *PdfActionGoTo) ToPdfObject() _eb.PdfObject {
	_fef.PdfAction.ToPdfObject()
	_ffd := _fef._dee
	_abg := _ffd.PdfObject.(*_eb.PdfObjectDictionary)
	_abg.SetIfNotNil("\u0053", _eb.MakeName(string(ActionTypeGoTo)))
	_abg.SetIfNotNil("\u0044", _fef.D)
	return _ffd
}

// SetHideMenubar sets the value of the hideMenubar flag.
func (_ceacc *ViewerPreferences) SetHideMenubar(hideMenubar bool) { _ceacc._gcbcb = &hideMenubar }

// NewKDictionary creates a new K dictionary object.
func NewKDictionary() *KDict { return &KDict{_eegge: make([]*KValue, 0), _dcgce: -1} }

// Encrypt encrypts the output file with a specified user/owner password.
func (_edbag *PdfWriter) Encrypt(userPass, ownerPass []byte, options *EncryptOptions) error {
	_agcbac := RC4_128bit
	if options != nil {
		_agcbac = options.Algorithm
	}
	_bdaba := _cg.PermOwner
	if options != nil {
		_bdaba = options.Permissions
	}
	var _cfcgg _be.Filter
	switch _agcbac {
	case RC4_128bit:
		_cfcgg = _be.NewFilterV2(16)
	case AES_128bit:
		_cfcgg = _be.NewFilterAESV2()
	case AES_256bit:
		_cfcgg = _be.NewFilterAESV3()
	default:
		return _e.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020a\u006cg\u006fr\u0069\u0074\u0068\u006d\u003a\u0020\u0025v", options.Algorithm)
	}
	_ggdac, _dbeec, _fecda := _eb.PdfCryptNewEncrypt(_cfcgg, userPass, ownerPass, _bdaba)
	if _fecda != nil {
		return _fecda
	}
	_edbag._gadcaa = _ggdac
	if _dbeec.Major != 0 {
		_edbag.SetVersion(_dbeec.Major, _dbeec.Minor)
	}
	_edbag._ffcbb = _dbeec.Encrypt
	_edbag._cebgc, _edbag._bgddbe = _dbeec.ID0, _dbeec.ID1
	_babgbc := _eb.MakeIndirectObject(_dbeec.Encrypt)
	_edbag._ebdeed = _babgbc
	_edbag.addObject(_babgbc)
	return nil
}

// GetModelFromPrimitive returns the model corresponding to the `primitive` PdfObject.
func (_beeag *modelManager) GetModelFromPrimitive(primitive _eb.PdfObject) PdfModel {
	model, _gbfdb := _beeag._fgffbd[primitive]
	if !_gbfdb {
		return nil
	}
	return model
}

// NewPdfAnnotationHighlight returns a new text highlight annotation.
func NewPdfAnnotationHighlight() *PdfAnnotationHighlight {
	_dad := NewPdfAnnotation()
	_egbg := &PdfAnnotationHighlight{}
	_egbg.PdfAnnotation = _dad
	_egbg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dad.SetContext(_egbg)
	return _egbg
}
func (_baccf *PdfAcroForm) filteredFields(_eadf FieldFilterFunc, _fbdg bool) []*PdfField {
	if _baccf == nil {
		return nil
	}
	return _degga(_baccf.Fields, _eadf, _fbdg)
}
func (_cdaab *pdfFontSimple) addEncoding() error {
	var (
		_cgbf  string
		_efcbg map[_fc.CharCode]_fc.GlyphName
		_ababd _fc.SimpleEncoder
	)
	if _cdaab.Encoder() != nil {
		_aaggd, _gefe := _cdaab.Encoder().(_fc.SimpleEncoder)
		if _gefe && _aaggd != nil {
			_cgbf = _aaggd.BaseName()
		}
	}
	if _cdaab.Encoding != nil {
		_bccgf, _bbcbb, _fdgf := _cdaab.getFontEncoding()
		if _fdgf != nil {
			_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0073\u0065F\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0053u\u0062t\u0079\u0070\u0065\u003d\u0025\u0071\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003d\u0025\u0073 \u0028\u0025\u0054\u0029\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cdaab._agcc, _cdaab._fgdee, _cdaab.Encoding, _cdaab.Encoding, _fdgf)
			return _fdgf
		}
		if _bccgf != "" {
			_cgbf = _bccgf
		}
		_efcbg = _bbcbb
		_ababd, _fdgf = _fc.NewSimpleTextEncoder(_cgbf, _efcbg)
		if _fdgf != nil {
			return _fdgf
		}
	}
	if _ababd == nil {
		_aaeab := _cdaab._bged
		if _aaeab != nil {
			switch _cdaab._fgdee {
			case "\u0054\u0079\u0070e\u0031":
				if _aaeab.fontFile != nil && _aaeab.fontFile._gggff != nil {
					_ddb.Log.Debug("\u0055\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006et\u0046\u0069\u006c\u0065")
					_ababd = _aaeab.fontFile._gggff
				}
			case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
				if _aaeab._bebgd != nil {
					_ddb.Log.Debug("\u0055s\u0069n\u0067\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032")
					_fbfg, _gecac := _aaeab._bebgd.MakeEncoder()
					if _gecac == nil {
						_ababd = _fbfg
					}
					if _cdaab._bgbg == nil {
						_cdaab._bgbg = _aaeab._bebgd.MakeToUnicode()
					}
				}
			}
		}
	}
	if _ababd != nil {
		if _efcbg != nil {
			_ddb.Log.Trace("\u0064\u0069\u0066fe\u0072\u0065\u006e\u0063\u0065\u0073\u003d\u0025\u002b\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _efcbg, _cdaab.baseFields())
			_ababd = _fc.ApplyDifferences(_ababd, _efcbg)
		}
		_cdaab.SetEncoder(_ababd)
	}
	return nil
}

// SetImageHandler sets the image handler used by the package.
func SetImageHandler(imgHandling ImageHandler) { ImageHandling = imgHandling }

// Add appends a top level outline item to the outline.
func (_cgddb *Outline) Add(item *OutlineItem) { _cgddb.Entries = append(_cgddb.Entries, item) }

// PdfModel is a higher level PDF construct which can be collapsed into a PdfObject.
// Each PdfModel has an underlying PdfObject and vice versa (one-to-one).
// Under normal circumstances there should only be one copy of each.
// Copies can be made, but care must be taken to do it properly.
type PdfModel interface {
	ToPdfObject() _eb.PdfObject
	GetContainingPdfObject() _eb.PdfObject
}

// ColorFromFloats returns a new PdfColorDevice based on the input slice of
// color components. The slice should contain four elements representing the
// cyan, magenta, yellow and key components of the color. The values of the
// elements should be between 0 and 1.
func (_dfgcb *PdfColorspaceDeviceCMYK) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 4 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ddad := vals[0]
	if _ddad < 0.0 || _ddad > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ddad)
		return nil, ErrColorOutOfRange
	}
	_edbgd := vals[1]
	if _edbgd < 0.0 || _edbgd > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _edbgd)
		return nil, ErrColorOutOfRange
	}
	_bbdg := vals[2]
	if _bbdg < 0.0 || _bbdg > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bbdg)
		return nil, ErrColorOutOfRange
	}
	_gacg := vals[3]
	if _gacg < 0.0 || _gacg > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _gacg)
		return nil, ErrColorOutOfRange
	}
	_ggff := NewPdfColorDeviceCMYK(_ddad, _edbgd, _bbdg, _gacg)
	return _ggff, nil
}

// NewPdfActionGoToE returns a new "go to embedded" action.
func NewPdfActionGoToE() *PdfActionGoToE {
	_gd := NewPdfAction()
	_aeb := &PdfActionGoToE{}
	_aeb.PdfAction = _gd
	_gd.SetContext(_aeb)
	return _aeb
}

// GetContainingPdfObject implements interface PdfModel.
func (_gdbf *PdfAnnotation) GetContainingPdfObject() _eb.PdfObject { return _gdbf._ggf }

// ToPdfObject returns a stream object.
func (_gcccf *XObjectImage) ToPdfObject() _eb.PdfObject {
	_cdga := _gcccf._gceffg
	if _gcccf._gfegb {
		return _cdga
	}
	_egfcg := _cdga.PdfObjectDictionary
	if _gcccf.Filter != nil {
		_egfcg = _gcccf.Filter.MakeStreamDict()
		_cdga.PdfObjectDictionary = _egfcg
	}
	_egfcg.Set("\u0054\u0079\u0070\u0065", _eb.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_egfcg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0049\u006d\u0061g\u0065"))
	_egfcg.Set("\u0057\u0069\u0064t\u0068", _eb.MakeInteger(*(_gcccf.Width)))
	_egfcg.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _eb.MakeInteger(*(_gcccf.Height)))
	if _gcccf.BitsPerComponent != nil {
		_egfcg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _eb.MakeInteger(*(_gcccf.BitsPerComponent)))
	}
	if _gcccf.ColorSpace != nil {
		_egfcg.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _gcccf.ColorSpace.ToPdfObject())
	}
	_egfcg.SetIfNotNil("\u0049\u006e\u0074\u0065\u006e\u0074", _gcccf.Intent)
	_egfcg.SetIfNotNil("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k", _gcccf.ImageMask)
	_egfcg.SetIfNotNil("\u004d\u0061\u0073\u006b", _gcccf.Mask)
	_cafec := _egfcg.Get("\u0044\u0065\u0063\u006f\u0064\u0065") != nil
	if _gcccf.Decode == nil && _cafec {
		_egfcg.Remove("\u0044\u0065\u0063\u006f\u0064\u0065")
	} else if _gcccf.Decode != nil {
		_egfcg.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _gcccf.Decode)
	}
	_egfcg.SetIfNotNil("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _gcccf.Interpolate)
	_egfcg.SetIfNotNil("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073", _gcccf.Alternatives)
	_egfcg.SetIfNotNil("\u0053\u004d\u0061s\u006b", _gcccf.SMask)
	_egfcg.SetIfNotNil("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061", _gcccf.SMaskInData)
	_egfcg.SetIfNotNil("\u004d\u0061\u0074t\u0065", _gcccf.Matte)
	_egfcg.SetIfNotNil("\u004e\u0061\u006d\u0065", _gcccf.Name)
	_egfcg.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _gcccf.StructParent)
	_egfcg.SetIfNotNil("\u0049\u0044", _gcccf.ID)
	_egfcg.SetIfNotNil("\u004f\u0050\u0049", _gcccf.OPI)
	_egfcg.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _gcccf.Metadata)
	_egfcg.SetIfNotNil("\u004f\u0043", _gcccf.OC)
	_egfcg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _eb.MakeInteger(int64(len(_gcccf.Stream))))
	_cdga.Stream = _gcccf.Stream
	_gcccf._gfegb = true
	return _cdga
}

// R returns the value of the red component of the color.
func (_cedb *PdfColorDeviceRGB) R() float64 { return _cedb[0] }

// GetColorspaces loads PdfPageResourcesColorspaces from `r.ColorSpace` and returns an error if there
// is a problem loading. Once loaded, the same object is returned on multiple calls.
func (_bgcdc *PdfPageResources) GetColorspaces() (*PdfPageResourcesColorspaces, error) {
	if _bgcdc._cfcff != nil {
		return _bgcdc._cfcff, nil
	}
	if _bgcdc.ColorSpace == nil {
		return nil, nil
	}
	_cacag, _feggf := _dfgf(_bgcdc.ColorSpace)
	if _feggf != nil {
		return nil, _feggf
	}
	_bgcdc._cfcff = _cacag
	return _bgcdc._cfcff, nil
}

// PdfActionJavaScript represents a javaScript action.
type PdfActionJavaScript struct {
	*PdfAction
	JS _eb.PdfObject
}

func _adca(_beaec StdFontName) (pdfFontSimple, error) {
	_agfbd, _afdebf := _fg.NewStdFontByName(_beaec)
	if !_afdebf {
		return pdfFontSimple{}, ErrFontNotSupported
	}
	_ccdf := _fbfae(_agfbd)
	return _ccdf, nil
}

// NewXObjectForm creates a brand new XObject Form. Creates a new underlying PDF object stream primitive.
func NewXObjectForm() *XObjectForm {
	_bafgd := &XObjectForm{}
	_gagaa := &_eb.PdfObjectStream{}
	_gagaa.PdfObjectDictionary = _eb.MakeDict()
	_bafgd._afcag = _gagaa
	return _bafgd
}

// SetForms sets the Acroform for a PDF file.
func (_ebcdf *PdfWriter) SetForms(form *PdfAcroForm) error { _ebcdf._dcbec = form; return nil }

// ToPdfObject returns the PDF representation of the colorspace.
func (_cfcdb *PdfColorspaceSpecialSeparation) ToPdfObject() _eb.PdfObject {
	_eceed := _eb.MakeArray(_eb.MakeName("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"))
	_eceed.Append(_cfcdb.ColorantName)
	_eceed.Append(_cfcdb.AlternateSpace.ToPdfObject())
	_eceed.Append(_cfcdb.TintTransform.ToPdfObject())
	if _cfcdb._degb != nil {
		_cfcdb._degb.PdfObject = _eceed
		return _cfcdb._degb
	}
	return _eceed
}

// NewStructTreeRoot creates a new structure tree root dictionary.
func NewStructTreeRoot() *StructTreeRoot {
	return &StructTreeRoot{K: []*KDict{}, RoleMap: _eb.MakeDict(), ParentTreeNextKey: 0}
}
func (_egd *PdfReader) newPdfActionMovieFromDict(_fbc *_eb.PdfObjectDictionary) (*PdfActionMovie, error) {
	return &PdfActionMovie{Annotation: _fbc.Get("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e"), T: _fbc.Get("\u0054"), Operation: _fbc.Get("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn")}, nil
}
func (_cfdcc *PdfWriter) writeNamesDictionary() error {
	if _cfdcc._agdbc == nil {
		return nil
	}
	_abccgb := _cfdcc._agdbc.ToPdfObject()
	_cfdcc._dbffa.Set("\u004e\u0061\u006de\u0073", _abccgb)
	_ffbfe := _cfdcc.addObjects(_abccgb)
	if _ffbfe != nil {
		return _ffbfe
	}
	return nil
}
func (_fgba *PdfColorspaceSpecialPattern) String() string {
	return "\u0050a\u0074\u0074\u0065\u0072\u006e"
}

// GetCharMetrics returns the char metrics for character code `code`.
// How it works:
//  1. It calls the GetCharMetrics function for the underlying font, either a simple font or
//     a Type0 font. The underlying font GetCharMetrics() functions do direct charcode ➞  metrics
//     mappings.
//  2. If the underlying font's GetCharMetrics() doesn't have a CharMetrics for `code` then a
//     a CharMetrics with the FontDescriptor's /MissingWidth is returned.
//  3. If there is no /MissingWidth then a failure is returned.
//
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
// well give them 0 width. There is no need for the bool return.
//
// TODO(gunnsth): Reconsider whether needed or if can map via GlyphName.
func (_cbbe *PdfFont) GetCharMetrics(code _fc.CharCode) (CharMetrics, bool) {
	var _bedaa _fg.CharMetrics
	switch _faaf := _cbbe._fdaa.(type) {
	case *pdfFontSimple:
		if _fbgg, _bdddd := _faaf.GetCharMetrics(code); _bdddd {
			return _fbgg, _bdddd
		}
	case *pdfFontType0:
		if _gfbf, _ffca := _faaf.GetCharMetrics(code); _ffca {
			return _gfbf, _ffca
		}
	case *pdfCIDFontType0:
		if _faeaf, _ddebc := _faaf.GetCharMetrics(code); _ddebc {
			return _faeaf, _ddebc
		}
	case *pdfCIDFontType2:
		if _eafb, _gfdcb := _faaf.GetCharMetrics(code); _gfdcb {
			return _eafb, _gfdcb
		}
	case *pdfFontType3:
		if _efccg, _efde := _faaf.GetCharMetrics(code); _efde {
			return _efccg, _efde
		}
	default:
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020G\u0065\u0074\u0043h\u0061\u0072\u004de\u0074\u0072i\u0063\u0073\u0020\u006e\u006f\u0074 \u0069mp\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d\u0025\u0054\u002e", _cbbe._fdaa)
		return _bedaa, false
	}
	if _cbbbb, _dfccge := _cbbe.GetFontDescriptor(); _dfccge == nil && _cbbbb != nil {
		return _fg.CharMetrics{Wx: _cbbbb._acaad}, true
	}
	_ddb.Log.Debug("\u0047\u0065\u0074\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069c\u0073\u0020\u0066\u006f\u0072 \u0066\u006fn\u0074\u003d\u0025\u0073", _cbbe)
	return _bedaa, false
}

// NewMultipleFontEncoder returns instantiates a new *MultipleFontEncoder
func NewMultipleFontEncoder(fonts []*PdfFont) *MultipleFontEncoder {
	return &MultipleFontEncoder{_fcgbc: fonts, CurrentFont: fonts[0]}
}

// PdfActionHide represents a hide action.
type PdfActionHide struct {
	*PdfAction
	T _eb.PdfObject
	H _eb.PdfObject
}

// NewStandard14FontMustCompile returns the standard 14 font named `basefont` as a *PdfFont.
// If `basefont` is one of the 14 Standard14Font values defined above then NewStandard14FontMustCompile
// is guaranteed to succeed.
func NewStandard14FontMustCompile(basefont StdFontName) *PdfFont {
	_debd, _bbegd := NewStandard14Font(basefont)
	if _bbegd != nil {
		panic(_e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0031\u0034\u0046\u006f\u006e\u0074 \u0025\u0023\u0071", basefont))
	}
	return _debd
}

// Names represents a PDF name tree.
//
// Ref: PDF32000_2008 chapter 7.7.4.
type Names struct {
	_cggag *_eb.PdfIndirectObject

	// Dests is a name tree mapping name string to destinations.
	Dests *_eb.PdfObjectDictionary

	// AP is a name tree mapping name strings to annotation appearance streams.
	AP *_eb.PdfObjectDictionary

	// JavaScript is a name tree mapping name strings to JavaScript actions.
	JavaScript *_eb.PdfObjectDictionary

	// Pages is a name tree mapping name strings to visible pages for use in interactive forms.
	Pages *_eb.PdfObjectDictionary

	// Templates is a name tree mapping name strings to invisible (template) pages for use in interactive forms.
	Templates *_eb.PdfObjectDictionary

	// IDS is a name tree mapping digital identifies to Web Capture content sets.
	IDS *_eb.PdfObjectDictionary

	// URLS is a name tree mapping URLs to Web Capture content sets.
	URLS *_eb.PdfObjectDictionary

	// EmbeddedFiles is a name tree mapping name strings to file specifications for embedded file streams.
	EmbeddedFiles *_eb.PdfObjectDictionary

	// AlternatePresentations is a name tree mapping name strings to alternate presentations.
	AlternatePresentations *_eb.PdfObjectDictionary

	// Renditions is a name tree mapping name strings (which shall have Unicode encoding) to rendition objects.
	Renditions *_eb.PdfObjectDictionary
}

// DecodeArray returns the range of color component values in the Lab colorspace.
func (_egbed *PdfColorspaceLab) DecodeArray() []float64 {
	_gcdc := []float64{0, 100}
	if _egbed != nil && _egbed.Range != nil && len(_egbed.Range) == 4 {
		_gcdc = append(_gcdc, _egbed.Range...)
	} else {
		_gcdc = append(_gcdc, -100, 100, -100, 100)
	}
	return _gcdc
}

// GetNamedDestinations returns the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_fbfag *PdfReader) GetNamedDestinations() (_eb.PdfObject, error) {
	_ffcab := _eb.ResolveReference(_fbfag._bagcfd.Get("\u0044\u0065\u0073t\u0073"))
	if _ffcab == nil {
		return nil, nil
	}
	if !_fbfag._cfcgdf {
		_bagfb := _fbfag.traverseObjectData(_ffcab)
		if _bagfb != nil {
			return nil, _bagfb
		}
	}
	return _ffcab, nil
}
func _gfgbd(_gbefb _eb.PdfObject) (*KDict, error) {
	_agffg := _eb.ResolveReference(_gbefb)
	if _agffg == nil {
		return nil, _e.Errorf("\u004b \u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	_aafde, _badfg := _eb.GetDict(_agffg)
	if !_badfg {
		return nil, _e.Errorf("\u004b\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_fabab := &KDict{}
	if _gacegg := _aafde.Get("\u0053"); _gacegg != nil {
		_fabab.S = _gacegg
	}
	if _egdcda := _aafde.Get("\u0050"); _egdcda != nil {
		_fabab.P = _egdcda
	}
	if _fcecc := _aafde.Get("\u0049\u0044"); _fcecc != nil {
		if _fgage, _gdggd := _eb.GetString(_fcecc); _gdggd {
			_fabab.ID = _fgage
		}
	}
	if _gfefd := _aafde.Get("\u0050\u0067"); _gfefd != nil {
		_fabab.Pg = _gfefd
	}
	if _eefgc := _aafde.Get("\u004b"); _eefgc != nil {
		_fabab.K = _eefgc
		switch _dgdba := _eefgc.(type) {
		case *_eb.PdfObjectArray:
			if _cdfd, _fegcc := _eb.GetArray(_eefgc); _fegcc {
				for _, _cecbe := range _cdfd.Elements() {
					switch _cefcf := _cecbe.(type) {
					case *_eb.PdfIndirectObject:
						_eegeaf, _cfegg := _gfgbd(_cefcf)
						if _cfegg != nil {
							_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u004b\u0020\u0063\u0068\u0069\u006c\u0064:\u0020\u0025\u0076", _cfegg)
							continue
						}
						_fabab._eegge = append(_fabab._eegge, &KValue{_ccbca: _eegeaf})
					case *_eb.PdfObjectInteger:
						if _bcbde, _fedbgd := _eb.GetIntVal(_cecbe); _fedbgd {
							_fabab._eegge = append(_fabab._eegge, &KValue{_ddaf: &_bcbde})
						}
					case *_eb.PdfObjectDictionary:
						_fabab._eegge = append(_fabab._eegge, &KValue{_fbgaa: _cecbe})
					}
				}
			}
		case *_eb.PdfIndirectObject:
			_ggbag, _gafce := _gfgbd(_dgdba)
			if _gafce != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u004b\u0020\u0063\u0068\u0069\u006c\u0064:\u0020\u0025\u0076", _gafce)
			}
			if _ggbag != nil {
				_fabab._eegge = append(_fabab._eegge, &KValue{_ccbca: _ggbag})
			}
		case *_eb.PdfObjectDictionary:
			_fabab._eegge = append(_fabab._eegge, &KValue{_fbgaa: _dgdba})
		case *_eb.PdfObjectInteger:
			if _fccf, _dcedb := _eb.GetIntVal(_eefgc); _dcedb {
				_fabab._eegge = append(_fabab._eegge, &KValue{_ddaf: &_fccf})
			}
		}
	}
	if _efebd := _aafde.Get("\u0041"); _efebd != nil {
		_fabab.A = _efebd
	}
	if _bdfge := _aafde.Get("\u0043"); _bdfge != nil {
		_fabab.C = _bdfge
	}
	if _dcga := _aafde.Get("\u0052"); _dcga != nil {
		if _gefde, _cfag := _eb.GetInt(_dcga); _cfag {
			_fabab.R = _gefde
		}
	}
	if _edfbf := _aafde.Get("\u0054"); _edfbf != nil {
		if _gcbee, _caceg := _eb.GetString(_edfbf); _caceg {
			_fabab.T = _gcbee
		}
	}
	if _fffcb := _aafde.Get("\u004c\u0061\u006e\u0067"); _fffcb != nil {
		if _fgdfda, _bcfdb := _eb.GetString(_fffcb); _bcfdb {
			_fabab.Lang = _fgdfda
		}
	}
	if _ebggc := _aafde.Get("\u0041\u006c\u0074"); _ebggc != nil {
		if _abede, _cbagd := _eb.GetString(_ebggc); _cbagd {
			_fabab.Alt = _abede
		}
	}
	if _bagfe := _aafde.Get("\u0045"); _bagfe != nil {
		if _abfae, _dabdc := _eb.GetString(_bagfe); _dabdc {
			_fabab.E = _abfae
		}
	}
	return _fabab, nil
}
func _dbeff(_dacdda []byte) bool {
	if len(_dacdda) < 4 {
		return true
	}
	for _bcefe := range _dacdda[:4] {
		_agcga := rune(_bcefe)
		if !_eg.Is(_eg.ASCII_Hex_Digit, _agcga) && !_eg.IsSpace(_agcga) {
			return true
		}
	}
	return false
}

// CharcodesToUnicodeWithStats is identical to CharcodesToUnicode except it returns more statistical
// information about hits and misses from the reverse mapping process.
// NOTE: The number of runes returned may be greater than the number of charcodes.
// TODO(peterwilliams97): Deprecate in v4 and use only CharcodesToStrings()
func (_fgfdb *PdfFont) CharcodesToUnicodeWithStats(charcodes []_fc.CharCode) (_afgcd []rune, _dabgd, _aeed int) {
	_cdceb, _dabgd, _aeed := _fgfdb.CharcodesToStrings(charcodes, "")
	return []rune(_cc.Join(_cdceb, "")), _dabgd, _aeed
}

// Tab order types.
type TabOrderType string

// NewOutlineDest returns a new outline destination which can be used
// with outline items.
func NewOutlineDest(page int64, x, y float64) OutlineDest {
	return OutlineDest{Page: page, Mode: "\u0058\u0059\u005a", X: x, Y: y}
}

// WriteString outputs the object as it is to be written to file.
func (_eefbc *pdfSignDictionary) WriteString() string {
	_eefbc._eegea = 0
	_eefbc._defaa = 0
	_eefbc._caeba = 0
	_eefbc._dbcd = 0
	_bfcfc := _dd.NewBuffer(nil)
	_bfcfc.WriteString("\u003c\u003c")
	for _, _acfdc := range _eefbc.Keys() {
		_ffccf := _eefbc.Get(_acfdc)
		switch _acfdc {
		case "\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e":
			_bfcfc.WriteString(_acfdc.WriteString())
			_bfcfc.WriteString("\u0020")
			_eefbc._caeba = _bfcfc.Len()
			_bfcfc.WriteString(_ffccf.WriteString())
			_bfcfc.WriteString("\u0020")
			_eefbc._dbcd = _bfcfc.Len() - 1
		case "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
			_bfcfc.WriteString(_acfdc.WriteString())
			_bfcfc.WriteString("\u0020")
			_eefbc._eegea = _bfcfc.Len()
			_bfcfc.WriteString(_ffccf.WriteString())
			_bfcfc.WriteString("\u0020")
			_eefbc._defaa = _bfcfc.Len() - 1
		default:
			_bfcfc.WriteString(_acfdc.WriteString())
			_bfcfc.WriteString("\u0020")
			_bfcfc.WriteString(_ffccf.WriteString())
		}
	}
	_bfcfc.WriteString("\u003e\u003e")
	return _bfcfc.String()
}
func (_cfae *PdfWriter) addObject(_cabcd _eb.PdfObject) bool {
	_aaecf := _cfae.hasObject(_cabcd)
	if !_aaecf {
		_bcaafe := _eb.ResolveReferencesDeep(_cabcd, _cfae._fabeca)
		if _bcaafe != nil {
			_ddb.Log.Debug("E\u0052R\u004f\u0052\u003a\u0020\u0025\u0076\u0020\u002d \u0073\u006b\u0069\u0070pi\u006e\u0067", _bcaafe)
		}
		_cfae._dcfgf = append(_cfae._dcfgf, _cabcd)
		_cfae._aeeda[_cabcd] = struct{}{}
		return true
	}
	return false
}

// ToPdfObject implements interface PdfModel.
func (_dece *PdfAnnotationStrikeOut) ToPdfObject() _eb.PdfObject {
	_dece.PdfAnnotation.ToPdfObject()
	_fcdd := _dece._ggf
	_gea := _fcdd.PdfObject.(*_eb.PdfObjectDictionary)
	_dece.PdfAnnotationMarkup.appendToPdfDictionary(_gea)
	_gea.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t"))
	_gea.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _dece.QuadPoints)
	return _fcdd
}
func _dfgag(_daegg *_eb.PdfObjectDictionary) (*PdfShadingPattern, error) {
	_bdcea := &PdfShadingPattern{}
	_bdgeg := _daegg.Get("\u0053h\u0061\u0064\u0069\u006e\u0067")
	if _bdgeg == nil {
		_ddb.Log.Debug("\u0053h\u0061d\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dcaa, _bgcgf := _aggce(_bdgeg)
	if _bgcgf != nil {
		_ddb.Log.Debug("\u0045r\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006eg\u0020s\u0068a\u0064\u0069\u006e\u0067\u003a\u0020\u0025v", _bgcgf)
		return nil, _bgcgf
	}
	_bdcea.Shading = _dcaa
	if _afbbe := _daegg.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _afbbe != nil {
		_fagcc, _bdac := _afbbe.(*_eb.PdfObjectArray)
		if !_bdac {
			_ddb.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _afbbe)
			return nil, _eb.ErrTypeError
		}
		_bdcea.Matrix = _fagcc
	}
	if _dabce := _daegg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _dabce != nil {
		_bdcea.ExtGState = _dabce
	}
	return _bdcea, nil
}

// ToPdfObject implements interface PdfModel.
func (_cafe *PdfAnnotationHighlight) ToPdfObject() _eb.PdfObject {
	_cafe.PdfAnnotation.ToPdfObject()
	_cgb := _cafe._ggf
	_fbb := _cgb.PdfObject.(*_eb.PdfObjectDictionary)
	_cafe.PdfAnnotationMarkup.appendToPdfDictionary(_fbb)
	_fbb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t"))
	_fbb.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _cafe.QuadPoints)
	return _cgb
}

// PdfAnnotationRichMedia represents Rich Media annotations.
type PdfAnnotationRichMedia struct {
	*PdfAnnotation
	RichMediaSettings _eb.PdfObject
	RichMediaContent  _eb.PdfObject
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_fcbag *PdfColorspaceDeviceGray) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ddfe, _gdeba := _eb.GetNumbersAsFloat(objects)
	if _gdeba != nil {
		return nil, _gdeba
	}
	return _fcbag.ColorFromFloats(_ddfe)
}

// String returns string value of output intent for given type
// ISO_19005-2 6.2.3: GTS_PDFA1 value should be used for PDF/A-1, A-2 and A-3 at least
func (_cabc PdfOutputIntentType) String() string {
	switch _cabc {
	case PdfOutputIntentTypeA1:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA2:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA3:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA4:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeX:
		return "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058"
	default:
		return "\u0055N\u0044\u0045\u0046\u0049\u004e\u0045D"
	}
}

// NewPdfFilespec returns an initialized generic PDF filespec model.
func NewPdfFilespec() *PdfFilespec {
	_cgca := &PdfFilespec{}
	_cgca._cbefe = _eb.MakeIndirectObject(_eb.MakeDict())
	return _cgca
}

// PdfActionType represents an action type in PDF (section 12.6.4 p. 417).
type PdfActionType string

// SetPageNumber sets the page number.
func (_egecd *KDict) SetPageNumber(pageNumber int64) { _egecd._dcgce = pageNumber }
func _efebc() string                                 { _dfbafc.Lock(); defer _dfbafc.Unlock(); return _gacdf }
func (_ffdfb *Image) resampleLowBits(_bfbc []uint32) {
	_abeb := _df.BytesPerLine(int(_ffdfb.Width), int(_ffdfb.BitsPerComponent), _ffdfb.ColorComponents)
	_fcdgf := make([]byte, _ffdfb.ColorComponents*_abeb*int(_ffdfb.Height))
	_dfce := int(_ffdfb.BitsPerComponent) * _ffdfb.ColorComponents * int(_ffdfb.Width)
	_dfca := uint8(8)
	var (
		_bdad, _gfcdc int
		_bgca         uint32
	)
	for _gcdd := 0; _gcdd < int(_ffdfb.Height); _gcdd++ {
		_gfcdc = _gcdd * _abeb
		for _cbbbaf := 0; _cbbbaf < _dfce; _cbbbaf++ {
			_bgca = _bfbc[_bdad]
			_dfca -= uint8(_ffdfb.BitsPerComponent)
			_fcdgf[_gfcdc] |= byte(_bgca) << _dfca
			if _dfca == 0 {
				_dfca = 8
				_gfcdc++
			}
			_bdad++
		}
	}
	_ffdfb.Data = _fcdgf
}
func _dba(_cbg _eb.PdfObject) (*PdfFilespec, error) {
	if _cbg == nil {
		return nil, nil
	}
	return NewPdfFilespecFromObj(_cbg)
}

// ToPdfObject converts PdfAcroForm to a PdfObject, i.e. an indirect object containing the
// AcroForm dictionary.
func (_fbaa *PdfAcroForm) ToPdfObject() _eb.PdfObject {
	_ddec := _fbaa._fbgad
	_fdff := _ddec.PdfObject.(*_eb.PdfObjectDictionary)
	if _fbaa.Fields != nil {
		_ceae := _eb.PdfObjectArray{}
		for _, _faffc := range *_fbaa.Fields {
			_eggbg := _faffc.GetContext()
			if _eggbg != nil {
				_ceae.Append(_eggbg.ToPdfObject())
			} else {
				_ceae.Append(_faffc.ToPdfObject())
			}
		}
		_fdff.Set("\u0046\u0069\u0065\u006c\u0064\u0073", &_ceae)
	}
	if _fbaa.NeedAppearances != nil {
		_fdff.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _fbaa.NeedAppearances)
	} else {
		if _ebdbb := _fdff.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _ebdbb != nil {
			_fdff.Remove("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073")
		}
	}
	if _fbaa.SigFlags != nil {
		_fdff.Set("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073", _fbaa.SigFlags)
	}
	if _fbaa.CO != nil {
		_fdff.Set("\u0043\u004f", _fbaa.CO)
	}
	if _fbaa.DR != nil {
		_fdff.Set("\u0044\u0052", _fbaa.DR.ToPdfObject())
	}
	if _fbaa.DA != nil {
		_fdff.Set("\u0044\u0041", _fbaa.DA)
	}
	if _fbaa.Q != nil {
		_fdff.Set("\u0051", _fbaa.Q)
	}
	if _fbaa.XFA != nil {
		_fdff.Set("\u0058\u0046\u0041", _fbaa.XFA)
	}
	if _fbaa.ADBEEchoSign != nil {
		_fdff.Set("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e", _fbaa.ADBEEchoSign)
	}
	return _ddec
}

// NewPdfActionURI returns a new "Uri" action.
func NewPdfActionURI() *PdfActionURI {
	_cdg := NewPdfAction()
	_af := &PdfActionURI{}
	_af.PdfAction = _cdg
	_cdg.SetContext(_af)
	return _af
}

// PdfColorspaceCalGray represents CalGray color space.
type PdfColorspaceCalGray struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      float64
	_gdee      *_eb.PdfIndirectObject
}

func (_gdede *PdfWriter) setWriter(_dgaa _bagf.Writer) {
	_gdede._dfabe = _gdede._gedbe
	_gdede._ccfbc = _ga.NewWriter(_dgaa)
}
func (_cffgb *LTV) getOCSPs(_gffda []*_bag.Certificate, _beffb map[string]*_bag.Certificate) ([][]byte, error) {
	_eggbb := make([][]byte, 0, len(_gffda))
	for _, _aabe := range _gffda {
		for _, _eeceb := range _aabe.OCSPServer {
			if _cffgb.CertClient.IsCA(_aabe) {
				continue
			}
			_efabc, _afgf := _beffb[_aabe.Issuer.CommonName]
			if !_afgf {
				_ddb.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _adfc, _egdafa := _cffgb.OCSPClient.MakeRequest(_eeceb, _aabe, _efabc)
			if _egdafa != nil {
				_ddb.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _egdafa)
				continue
			}
			_eggbb = append(_eggbb, _adfc)
		}
	}
	return _eggbb, nil
}

// GetNumComponents returns the number of color components (3 for RGB).
func (_eadb *PdfColorDeviceRGB) GetNumComponents() int { return 3 }
func _gbge(_dbff _eb.PdfObject) (*PdfColorspaceDeviceN, error) {
	_ebb := NewPdfColorspaceDeviceN()
	if _fagg, _cggc := _dbff.(*_eb.PdfIndirectObject); _cggc {
		_ebb._cfef = _fagg
	}
	_dbff = _eb.TraceToDirectObject(_dbff)
	_gbdce, _afag := _dbff.(*_eb.PdfObjectArray)
	if !_afag {
		return nil, _e.Errorf("\u0064\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if _gbdce.Len() != 4 && _gbdce.Len() != 5 {
		return nil, _e.Errorf("\u0064\u0065\u0076ic\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0049\u006ec\u006fr\u0072e\u0063t\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_dbff = _gbdce.Get(0)
	_ccbbf, _afag := _dbff.(*_eb.PdfObjectName)
	if !_afag {
		return nil, _e.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_ccbbf != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
		return nil, _e.Errorf("\u0064\u0065v\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020na\u006d\u0065")
	}
	_dbff = _gbdce.Get(1)
	_dbff = _eb.TraceToDirectObject(_dbff)
	_gaceg, _afag := _dbff.(*_eb.PdfObjectArray)
	if !_afag {
		return nil, _e.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ebb.ColorantNames = _gaceg
	_dbff = _gbdce.Get(2)
	_afed, _adad := NewPdfColorspaceFromPdfObject(_dbff)
	if _adad != nil {
		return nil, _adad
	}
	_ebb.AlternateSpace = _afed
	_eafa, _adad := _cccfa(_gbdce.Get(3))
	if _adad != nil {
		return nil, _adad
	}
	_ebb.TintTransform = _eafa
	if _gbdce.Len() == 5 {
		_edffd, _bdage := _abf(_gbdce.Get(4))
		if _bdage != nil {
			return nil, _bdage
		}
		_ebb.Attributes = _edffd
	}
	return _ebb, nil
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_dgdc pdfFontType0) GetRuneMetrics(r rune) (_fg.CharMetrics, bool) {
	if _dgdc.DescendantFont == nil {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _dgdc)
		return _fg.CharMetrics{}, false
	}
	return _dgdc.DescendantFont.GetRuneMetrics(r)
}

// ColorToRGB converts a CalGray color to an RGB color.
func (_aaec *PdfColorspaceCalGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_efg, _dag := color.(*PdfColorCalGray)
	if !_dag {
		_ddb.Log.Debug("\u0049n\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006eo\u0074\u0020\u0063\u0061\u006c\u0020\u0067\u0072\u0061\u0079")
		return nil, _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	ANorm := _efg.Val()
	X := _aaec.WhitePoint[0] * _gg.Pow(ANorm, _aaec.Gamma)
	Y := _aaec.WhitePoint[1] * _gg.Pow(ANorm, _aaec.Gamma)
	Z := _aaec.WhitePoint[2] * _gg.Pow(ANorm, _aaec.Gamma)
	_acdf := 3.240479*X + -1.537150*Y + -0.498535*Z
	_febf := -0.969256*X + 1.875992*Y + 0.041556*Z
	_adcg := 0.055648*X + -0.204043*Y + 1.057311*Z
	_acdf = _gg.Min(_gg.Max(_acdf, 0), 1.0)
	_febf = _gg.Min(_gg.Max(_febf, 0), 1.0)
	_adcg = _gg.Min(_gg.Max(_adcg, 0), 1.0)
	return NewPdfColorDeviceRGB(_acdf, _febf, _adcg), nil
}

// IsTerminal returns true for terminal fields, false otherwise.
// Terminal fields are fields whose descendants are only widget annotations.
func (_fcddf *PdfField) IsTerminal() bool { return len(_fcddf.Kids) == 0 }

// NewLTV returns a new LTV client.
func NewLTV(appender *PdfAppender) (*LTV, error) {
	_cgae := appender.Reader.DSS
	if _cgae == nil {
		_cgae = NewDSS()
	}
	if _gaef := _cgae.GenerateHashMaps(); _gaef != nil {
		return nil, _gaef
	}
	return &LTV{CertClient: _dda.NewCertClient(), OCSPClient: _dda.NewOCSPClient(), CRLClient: _dda.NewCRLClient(), SkipExisting: true, _feba: appender, _fadg: _cgae}, nil
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action.
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	_bff := NewPdfAction()
	_egfb := &PdfActionGoTo3DView{}
	_egfb.PdfAction = _bff
	_bff.SetContext(_egfb)
	return _egfb
}
func (_cbda *pdfFontSimple) baseFields() *fontCommon { return &_cbda.fontCommon }

var (
	CourierName              = _fg.CourierName
	CourierBoldName          = _fg.CourierBoldName
	CourierObliqueName       = _fg.CourierObliqueName
	CourierBoldObliqueName   = _fg.CourierBoldObliqueName
	HelveticaName            = _fg.HelveticaName
	HelveticaBoldName        = _fg.HelveticaBoldName
	HelveticaObliqueName     = _fg.HelveticaObliqueName
	HelveticaBoldObliqueName = _fg.HelveticaBoldObliqueName
	SymbolName               = _fg.SymbolName
	ZapfDingbatsName         = _fg.ZapfDingbatsName
	TimesRomanName           = _fg.TimesRomanName
	TimesBoldName            = _fg.TimesBoldName
	TimesItalicName          = _fg.TimesItalicName
	TimesBoldItalicName      = _fg.TimesBoldItalicName
)

// NewPdfActionLaunch returns a new "launch" action.
func NewPdfActionLaunch() *PdfActionLaunch {
	_aa := NewPdfAction()
	_gda := &PdfActionLaunch{}
	_gda.PdfAction = _aa
	_aa.SetContext(_gda)
	return _gda
}

// IsTiling specifies if the pattern is a tiling pattern.
func (_egbdb *PdfPattern) IsTiling() bool { return _egbdb.PatternType == 1 }

// GetCatalogViewerPreferences gets catalog ViewerPreferences object.
func (_gdbea *PdfReader) GetCatalogViewerPreferences() (_eb.PdfObject, bool) {
	if _gdbea._bagcfd == nil {
		return nil, false
	}
	_dfcfc := _gdbea._bagcfd.Get("\u0056\u0069\u0065\u0077\u0065\u0072\u0050\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0073")
	return _dfcfc, _dfcfc != nil
}

// PageBoundary represents the name of the page boundary representing
// the visible area.
type PageBoundary string

func _edba(_fdfbg _eb.PdfObject) (string, error) {
	_fdfbg = _eb.TraceToDirectObject(_fdfbg)
	switch _dabbd := _fdfbg.(type) {
	case *_eb.PdfObjectString:
		return _dabbd.Str(), nil
	case *_eb.PdfObjectStream:
		_eedfg, _eeaafg := _eb.DecodeStream(_dabbd)
		if _eeaafg != nil {
			return "", _eeaafg
		}
		return string(_eedfg), nil
	}
	return "", _e.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _fdfbg)
}

type pdfFontType0 struct {
	fontCommon
	_eecg          *_eb.PdfIndirectObject
	_edcff         _fc.TextEncoder
	Encoding       _eb.PdfObject
	DescendantFont *PdfFont
	_ebeb          *_ff.CMap
}

// NewPdfColorspaceDeviceRGB returns a new RGB colorspace object.
func NewPdfColorspaceDeviceRGB() *PdfColorspaceDeviceRGB { return &PdfColorspaceDeviceRGB{} }

// ColorFromPdfObjects returns a new PdfColor based on input color components. The input PdfObjects should
// be numeric.
func (_ggabb *PdfColorspaceDeviceN) ColorFromPdfObjects(objects []_eb.PdfObject) (PdfColor, error) {
	if len(objects) != _ggabb.GetNumComponents() {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eeec, _cegdgcb := _eb.GetNumbersAsFloat(objects)
	if _cegdgcb != nil {
		return nil, _cegdgcb
	}
	return _ggabb.ColorFromFloats(_eeec)
}

// NewPdfColorspaceSpecialIndexed returns a new Indexed color.
func NewPdfColorspaceSpecialIndexed() *PdfColorspaceSpecialIndexed {
	return &PdfColorspaceSpecialIndexed{HiVal: 255}
}

// NewPdfColorCalGray returns a new CalGray color.
func NewPdfColorCalGray(grayVal float64) *PdfColorCalGray {
	_cbdf := PdfColorCalGray(grayVal)
	return &_cbdf
}
func _cbff(_eeefe _eb.PdfObject) (*PdfFontDescriptor, error) {
	_eaba := &PdfFontDescriptor{}
	_eeefe = _eb.ResolveReference(_eeefe)
	if _ggcga, _agadc := _eeefe.(*_eb.PdfIndirectObject); _agadc {
		_eaba._fdea = _ggcga
		_eeefe = _ggcga.PdfObject
	}
	_gcbcd, _fdaaf := _eb.GetDict(_eeefe)
	if !_fdaaf {
		_ddb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0044\u0065\u0073c\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _eeefe)
		return nil, _eb.ErrTypeError
	}
	if _cfgda := _gcbcd.Get("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"); _cfgda != nil {
		_eaba.FontName = _cfgda
	} else {
		_ddb.Log.Debug("\u0049n\u0063\u006fm\u0070\u0061\u0074\u0069b\u0069\u006c\u0069t\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u004e\u0061me\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006di\u0073\u0073i\u006e\u0067")
	}
	_afeag, _ := _eb.GetName(_eaba.FontName)
	if _cadcf := _gcbcd.Get("\u0054\u0079\u0070\u0065"); _cadcf != nil {
		_begb, _afbe := _cadcf.(*_eb.PdfObjectName)
		if !_afbe || string(*_begb) != "\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072" {
			_ddb.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072i\u0070t\u006f\u0072\u0020\u0054y\u0070\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0054\u0029\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _cadcf, _afeag, _eaba.FontName)
		}
	} else {
		_ddb.Log.Trace("\u0049\u006ec\u006f\u006d\u0070\u0061\u0074i\u0062\u0069\u006c\u0069\u0074y\u003a\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _afeag, _eaba.FontName)
	}
	_eaba.FontFamily = _gcbcd.Get("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079")
	_eaba.FontStretch = _gcbcd.Get("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068")
	_eaba.FontWeight = _gcbcd.Get("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074")
	_eaba.Flags = _gcbcd.Get("\u0046\u006c\u0061g\u0073")
	_eaba.FontBBox = _gcbcd.Get("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078")
	_eaba.ItalicAngle = _gcbcd.Get("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065")
	_eaba.Ascent = _gcbcd.Get("\u0041\u0073\u0063\u0065\u006e\u0074")
	_eaba.Descent = _gcbcd.Get("\u0044e\u0073\u0063\u0065\u006e\u0074")
	_eaba.Leading = _gcbcd.Get("\u004ce\u0061\u0064\u0069\u006e\u0067")
	_eaba.CapHeight = _gcbcd.Get("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t")
	_eaba.XHeight = _gcbcd.Get("\u0058H\u0065\u0069\u0067\u0068\u0074")
	_eaba.StemV = _gcbcd.Get("\u0053\u0074\u0065m\u0056")
	_eaba.StemH = _gcbcd.Get("\u0053\u0074\u0065m\u0048")
	_eaba.AvgWidth = _gcbcd.Get("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068")
	_eaba.MaxWidth = _gcbcd.Get("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068")
	_eaba.MissingWidth = _gcbcd.Get("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068")
	_eaba.FontFile = _gcbcd.Get("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065")
	_eaba.FontFile2 = _gcbcd.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652")
	_eaba.FontFile3 = _gcbcd.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653")
	_eaba.CharSet = _gcbcd.Get("\u0043h\u0061\u0072\u0053\u0065\u0074")
	_eaba.Style = _gcbcd.Get("\u0053\u0074\u0079l\u0065")
	_eaba.Lang = _gcbcd.Get("\u004c\u0061\u006e\u0067")
	_eaba.FD = _gcbcd.Get("\u0046\u0044")
	_eaba.CIDSet = _gcbcd.Get("\u0043\u0049\u0044\u0053\u0065\u0074")
	if _eaba.Flags != nil {
		if _fbecd, _cafg := _eb.GetIntVal(_eaba.Flags); _cafg {
			_eaba._eeda = _fbecd
		}
	}
	if _eaba.MissingWidth != nil {
		if _ggeb, _bcbd := _eb.GetNumberAsFloat(_eaba.MissingWidth); _bcbd == nil {
			_eaba._acaad = _ggeb
		}
	}
	if _eaba.FontFile != nil {
		_eeegb, _affc := _eaff(_eaba.FontFile)
		if _affc != nil {
			return _eaba, _affc
		}
		_ddb.Log.Trace("f\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u003d\u0025\u0073", _eeegb)
		_eaba.fontFile = _eeegb
	}
	if _eaba.FontFile2 != nil {
		_dgeb, _eafg := _fg.NewFontFile2FromPdfObject(_eaba.FontFile2)
		if _eafg != nil {
			return _eaba, _eafg
		}
		_ddb.Log.Trace("\u0066\u006f\u006et\u0046\u0069\u006c\u0065\u0032\u003d\u0025\u0073", _dgeb.String())
		_eaba._bebgd = &_dgeb
	}
	return _eaba, nil
}

// ToPdfObject returns the PDF representation of the DSS dictionary.
func (_ggbfd *DSS) ToPdfObject() _eb.PdfObject {
	_cdbbf := _ggbfd._fcdgc.PdfObject.(*_eb.PdfObjectDictionary)
	_cdbbf.Clear()
	_beaad := _eb.MakeDict()
	for _ffbg, _fcaeg := range _ggbfd.VRI {
		_beaad.Set(*_eb.MakeName(_ffbg), _fcaeg.ToPdfObject())
	}
	_cdbbf.SetIfNotNil("\u0043\u0065\u0072t\u0073", _cafc(_ggbfd.Certs))
	_cdbbf.SetIfNotNil("\u004f\u0043\u0053P\u0073", _cafc(_ggbfd.OCSPs))
	_cdbbf.SetIfNotNil("\u0043\u0052\u004c\u0073", _cafc(_ggbfd.CRLs))
	_cdbbf.Set("\u0056\u0052\u0049", _beaad)
	return _ggbfd._fcdgc
}

// AddFont adds a font dictionary to the Font resources.
func (_gagg *PdfPage) AddFont(name _eb.PdfObjectName, font _eb.PdfObject) error {
	if _gagg.Resources == nil {
		_gagg.Resources = NewPdfPageResources()
	}
	if _gagg.Resources.Font == nil {
		_gagg.Resources.Font = _eb.MakeDict()
	}
	_fbddc, _afced := _eb.TraceToDirectObject(_gagg.Resources.Font).(*_eb.PdfObjectDictionary)
	if !_afced {
		_ddb.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0066\u006f\u006et \u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a \u0025\u0076", _eb.TraceToDirectObject(_gagg.Resources.Font))
		return _dcf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_fbddc.Set(name, font)
	return nil
}

// PdfColorLab represents a color in the L*, a*, b* 3 component colorspace.
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorLab [3]float64

func (_gaba *PdfReader) newPdfAnnotationSquareFromDict(_aadf *_eb.PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	_ega := PdfAnnotationSquare{}
	_ecee, _gfce := _gaba.newPdfAnnotationMarkupFromDict(_aadf)
	if _gfce != nil {
		return nil, _gfce
	}
	_ega.PdfAnnotationMarkup = _ecee
	_ega.BS = _aadf.Get("\u0042\u0053")
	_ega.IC = _aadf.Get("\u0049\u0043")
	_ega.BE = _aadf.Get("\u0042\u0045")
	_ega.RD = _aadf.Get("\u0052\u0044")
	return &_ega, nil
}

// StdFontName represents name of a standard font.
type StdFontName = _fg.StdFontName

// ToInteger convert to an integer format.
func (_adbc *PdfColorCalRGB) ToInteger(bits int) [3]uint32 {
	_cdfgg := _gg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_cdfgg * _adbc.A()), uint32(_cdfgg * _adbc.B()), uint32(_cdfgg * _adbc.C())}
}

// PdfPage represents a page in a PDF document. (7.7.3.3 - Table 30).
type PdfPage struct {
	Parent               _eb.PdfObject
	LastModified         *PdfDate
	Resources            *PdfPageResources
	CropBox              *PdfRectangle
	MediaBox             *PdfRectangle
	BleedBox             *PdfRectangle
	TrimBox              *PdfRectangle
	ArtBox               *PdfRectangle
	BoxColorInfo         _eb.PdfObject
	Contents             _eb.PdfObject
	Rotate               *int64
	Group                _eb.PdfObject
	Thumb                _eb.PdfObject
	B                    _eb.PdfObject
	Dur                  _eb.PdfObject
	Trans                _eb.PdfObject
	AA                   _eb.PdfObject
	Metadata             _eb.PdfObject
	PieceInfo            _eb.PdfObject
	StructParents        _eb.PdfObject
	ID                   _eb.PdfObject
	PZ                   _eb.PdfObject
	SeparationInfo       _eb.PdfObject
	Tabs                 _eb.PdfObject
	TemplateInstantiated _eb.PdfObject
	PresSteps            _eb.PdfObject
	UserUnit             _eb.PdfObject
	VP                   _eb.PdfObject
	Annots               _eb.PdfObject
	_dbga                []*PdfAnnotation
	_aaagaa              *_eb.PdfObjectDictionary
	_efcff               *_eb.PdfIndirectObject
	_eaebc               _eb.PdfObjectDictionary
	_dcdfd               *PdfReader
}

// C returns the value of the C component of the color.
func (_ddceg *PdfColorCalRGB) C() float64 { return _ddceg[2] }

// IsHideToolbar returns the value of the hideToolbar flag.
func (_deeec *ViewerPreferences) IsHideToolbar() bool {
	if _deeec._caafc == nil {
		return false
	}
	return *_deeec._caafc
}

// ToPdfObject implements interface PdfModel.
func (_fbga *PdfAnnotationRichMedia) ToPdfObject() _eb.PdfObject {
	_fbga.PdfAnnotation.ToPdfObject()
	_gbf := _fbga._ggf
	_fdbf := _gbf.PdfObject.(*_eb.PdfObjectDictionary)
	_fdbf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a"))
	_fdbf.SetIfNotNil("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073", _fbga.RichMediaSettings)
	_fdbf.SetIfNotNil("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074", _fbga.RichMediaContent)
	return _gbf
}

// NewXObjectImage returns a new XObjectImage.
func NewXObjectImage() *XObjectImage {
	_ddgdd := &XObjectImage{}
	_gedab := &_eb.PdfObjectStream{}
	_gedab.PdfObjectDictionary = _eb.MakeDict()
	_ddgdd._gceffg = _gedab
	return _ddgdd
}
func _abbaa(_gdbfa *fontCommon) *pdfFontSimple { return &pdfFontSimple{fontCommon: *_gdbfa} }

// AddWatermarkImage adds a watermark to the page.
func (_agcbf *PdfPage) AddWatermarkImage(ximg *XObjectImage, opt WatermarkImageOptions) error {
	_deda, _gccbf := _agcbf.GetMediaBox()
	if _gccbf != nil {
		return _gccbf
	}
	_cfgg := _deda.Urx - _deda.Llx
	_dccg := _deda.Ury - _deda.Lly
	_daca := float64(*ximg.Width)
	_bdafe := (_cfgg - _daca) / 2
	if opt.FitToWidth {
		_daca = _cfgg
		_bdafe = 0
	}
	_aefaa := _dccg
	_gfbg := float64(0)
	if opt.PreserveAspectRatio {
		_aefaa = _daca * float64(*ximg.Height) / float64(*ximg.Width)
		_gfbg = (_dccg - _aefaa) / 2
	}
	if _agcbf.Resources == nil {
		_agcbf.Resources = NewPdfPageResources()
	}
	_bgag := 0
	_cfeaf := _eb.PdfObjectName(_e.Sprintf("\u0049\u006d\u0077%\u0064", _bgag))
	for _agcbf.Resources.HasXObjectByName(_cfeaf) {
		_bgag++
		_cfeaf = _eb.PdfObjectName(_e.Sprintf("\u0049\u006d\u0077%\u0064", _bgag))
	}
	_gccbf = _agcbf.AddImageResource(_cfeaf, ximg)
	if _gccbf != nil {
		return _gccbf
	}
	_bgag = 0
	_gafeb := _eb.PdfObjectName(_e.Sprintf("\u0047\u0053\u0025\u0064", _bgag))
	for _agcbf.HasExtGState(_gafeb) {
		_bgag++
		_gafeb = _eb.PdfObjectName(_e.Sprintf("\u0047\u0053\u0025\u0064", _bgag))
	}
	_edccc := _eb.MakeDict()
	_edccc.Set("\u0042\u004d", _eb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	_edccc.Set("\u0043\u0041", _eb.MakeFloat(opt.Alpha))
	_edccc.Set("\u0063\u0061", _eb.MakeFloat(opt.Alpha))
	_gccbf = _agcbf.AddExtGState(_gafeb, _edccc)
	if _gccbf != nil {
		return _gccbf
	}
	_aefge := _e.Sprintf("\u0071\u000a"+"\u002f%\u0073\u0020\u0067\u0073\u000a"+"%\u002e\u0030\u0066\u0020\u0030\u00200\u0020\u0025\u002e\u0030\u0066\u0020\u0025\u002e\u0034f\u0020\u0025\u002e4\u0066 \u0063\u006d\u000a"+"\u002f%\u0073\u0020\u0044\u006f\u000a"+"\u0051", _gafeb, _daca, _aefaa, _bdafe, _gfbg, _cfeaf)
	_agcbf.AddContentStreamByString(_aefge)
	return nil
}

// NewPdfColorDeviceGray returns a new grayscale color based on an input grayscale float value in range [0-1].
func NewPdfColorDeviceGray(grayVal float64) *PdfColorDeviceGray {
	_cga := PdfColorDeviceGray(grayVal)
	return &_cga
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// red, green and blue components of the color. The values of the elements
// should be between 0 and 1.
func (_bbcc *PdfColorspaceDeviceRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _dcf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_afeg := vals[0]
	if _afeg < 0.0 || _afeg > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _afeg)
		return nil, ErrColorOutOfRange
	}
	_faee := vals[1]
	if _faee < 0.0 || _faee > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _afeg)
		return nil, ErrColorOutOfRange
	}
	_efedd := vals[2]
	if _efedd < 0.0 || _efedd > 1.0 {
		_ddb.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _afeg)
		return nil, ErrColorOutOfRange
	}
	_agfc := NewPdfColorDeviceRGB(_afeg, _faee, _efedd)
	return _agfc, nil
}

// ToPdfObject converts the font to a PDF representation.
func (_cdae *pdfFontType0) ToPdfObject() _eb.PdfObject {
	if _cdae._eecg == nil {
		_cdae._eecg = &_eb.PdfIndirectObject{}
	}
	_gcae := _cdae.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0030")
	_cdae._eecg.PdfObject = _gcae
	if _cdae.Encoding != nil {
		_gcae.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _cdae.Encoding)
	} else if _cdae._edcff != nil {
		_gcae.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _cdae._edcff.ToPdfObject())
	}
	if _cdae.DescendantFont != nil {
		_gcae.Set("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073", _eb.MakeArray(_cdae.DescendantFont.ToPdfObject()))
	}
	return _cdae._eecg
}
func _abf(_addd _eb.PdfObject) (*PdfColorspaceDeviceNAttributes, error) {
	_fedff := &PdfColorspaceDeviceNAttributes{}
	var _facb *_eb.PdfObjectDictionary
	switch _aaefd := _addd.(type) {
	case *_eb.PdfIndirectObject:
		_fedff._dbggc = _aaefd
		var _dafg bool
		_facb, _dafg = _aaefd.PdfObject.(*_eb.PdfObjectDictionary)
		if !_dafg {
			_ddb.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
			return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	case *_eb.PdfObjectDictionary:
		_facb = _aaefd
	case *_eb.PdfObjectReference:
		_cdaad := _aaefd.Resolve()
		return _abf(_cdaad)
	default:
		_ddb.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
		return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _cbdb := _facb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _cbdb != nil {
		_edcfb, _fbcc := _eb.TraceToDirectObject(_cbdb).(*_eb.PdfObjectName)
		if !_fbcc {
			_ddb.Log.Error("\u0044\u0065vi\u0063\u0065\u004e \u0061\u0074\u0074\u0072ibu\u0074e \u0053\u0075\u0062\u0074\u0079\u0070\u0065 t\u0079\u0070\u0065\u0020\u0065\u0072\u0072o\u0072")
			return nil, _dcf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fedff.Subtype = _edcfb
	}
	if _fcag := _facb.Get("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s"); _fcag != nil {
		_fedff.Colorants = _fcag
	}
	if _fceb := _facb.Get("\u0050r\u006f\u0063\u0065\u0073\u0073"); _fceb != nil {
		_fedff.Process = _fceb
	}
	if _agdg := _facb.Get("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073"); _agdg != nil {
		_fedff.MixingHints = _agdg
	}
	return _fedff, nil
}

// GetRuneMetrics returns the char metrics for a rune.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
// well give them 0 width. There is no need for the bool return.
func (_fcdecg *PdfFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_fegac := _fcdecg.actualFont()
	if _fegac == nil {
		_ddb.Log.Debug("ER\u0052\u004fR\u003a\u0020\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u004e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020f\u006fr\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079p\u0065=\u0025\u0023T", _fcdecg._fdaa)
		return _fg.CharMetrics{}, false
	}
	if _dgcab, _ecfbg := _fegac.GetRuneMetrics(r); _ecfbg {
		return _dgcab, true
	}
	if _abcc, _fcdb := _fcdecg.GetFontDescriptor(); _fcdb == nil && _abcc != nil {
		return _fg.CharMetrics{Wx: _abcc._acaad}, true
	}
	_ddb.Log.Debug("\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061r\u004d\u0065\u0074\u0072i\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _fcdecg)
	return _fg.CharMetrics{}, false
}

// SetShadingByName sets a shading resource specified by keyName.
func (_aebaa *PdfPageResources) SetShadingByName(keyName _eb.PdfObjectName, shadingObj _eb.PdfObject) error {
	if _aebaa.Shading == nil {
		_aebaa.Shading = _eb.MakeDict()
	}
	_ffddb, _fegbab := _eb.GetDict(_aebaa.Shading)
	if !_fegbab {
		return _eb.ErrTypeError
	}
	_ffddb.Set(keyName, shadingObj)
	return nil
}

// SetPrintPageRange sets the value of the printPageRange.
func (_eacgaa *ViewerPreferences) SetPrintPageRange(printPageRange []int) {
	_eacgaa._gcgfc = printPageRange
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_dbcbb *PdfShadingType1) ToPdfObject() _eb.PdfObject {
	_dbcbb.PdfShading.ToPdfObject()
	_bfbcf, _feaba := _dbcbb.getShadingDict()
	if _feaba != nil {
		_ddb.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _dbcbb.Domain != nil {
		_bfbcf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _dbcbb.Domain)
	}
	if _dbcbb.Matrix != nil {
		_bfbcf.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _dbcbb.Matrix)
	}
	if _dbcbb.Function != nil {
		if len(_dbcbb.Function) == 1 {
			_bfbcf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dbcbb.Function[0].ToPdfObject())
		} else {
			_aagf := _eb.MakeArray()
			for _, _cbeea := range _dbcbb.Function {
				_aagf.Append(_cbeea.ToPdfObject())
			}
			_bfbcf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _aagf)
		}
	}
	return _dbcbb._cefaa
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_dbaeg pdfCIDFontType2) GetCharMetrics(code _fc.CharCode) (_fg.CharMetrics, bool) {
	if _aecfe, _bbge := _dbaeg._fcgg[code]; _bbge {
		return _fg.CharMetrics{Wx: _aecfe}, true
	}
	_fggg := rune(code)
	_facf, _dacgb := _dbaeg._caba[_fggg]
	if !_dacgb {
		_facf = int(_dbaeg._aggd)
	}
	return _fg.CharMetrics{Wx: float64(_facf)}, true
}
func _bgadc(_bbdbe []byte) (_eegbe, _dagc string, _gegbe error) {
	_ddb.Log.Trace("g\u0065\u0074\u0041\u0053CI\u0049S\u0065\u0063\u0074\u0069\u006fn\u0073\u003a\u0020\u0025\u0064\u0020", len(_bbdbe))
	_ebedb := _ffee.FindIndex(_bbdbe)
	if _ebedb == nil {
		_ddb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0067\u0065\u0074\u0041\u0053\u0043\u0049\u0049\u0053\u0065\u0063\u0074\u0069o\u006e\u0073\u002e\u0020\u004e\u006f\u0020d\u0069\u0063\u0074\u002e")
		return "", "", _eb.ErrTypeError
	}
	_ecfg := _ebedb[1]
	_cbecd := _cc.Index(string(_bbdbe[_ecfg:]), _ccbg)
	if _cbecd < 0 {
		_eegbe = string(_bbdbe[_ecfg:])
		return _eegbe, "", nil
	}
	_aadfe := _ecfg + _cbecd
	_eegbe = string(_bbdbe[_ecfg:_aadfe])
	_caccd := _aadfe
	_cbecd = _cc.Index(string(_bbdbe[_caccd:]), _ceab)
	if _cbecd < 0 {
		_ddb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0067e\u0074\u0041\u0053\u0043\u0049\u0049\u0053e\u0063\u0074\u0069\u006f\u006e\u0073\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gegbe)
		return "", "", _eb.ErrTypeError
	}
	_cgbgb := _caccd + _cbecd
	_dagc = string(_bbdbe[_caccd:_cgbgb])
	return _eegbe, _dagc, nil
}
func (_aed *PdfReader) newPdfAnnotationMarkupFromDict(_aadc *_eb.PdfObjectDictionary) (*PdfAnnotationMarkup, error) {
	_fcd := &PdfAnnotationMarkup{}
	if _cbbg := _aadc.Get("\u0054"); _cbbg != nil {
		_fcd.T = _cbbg
	}
	if _adc := _aadc.Get("\u0050\u006f\u0070u\u0070"); _adc != nil {
		_gbg, _eeae := _adc.(*_eb.PdfIndirectObject)
		if !_eeae {
			if _, _ebf := _adc.(*_eb.PdfObjectNull); !_ebf {
				return nil, _dcf.New("p\u006f\u0070\u0075\u0070\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074\u0020t\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			}
		} else {
			_agfa, _bedb := _aed.newPdfAnnotationFromIndirectObject(_gbg)
			if _bedb != nil {
				return nil, _bedb
			}
			if _agfa != nil {
				_fbfb, _cag := _agfa._cdb.(*PdfAnnotationPopup)
				if !_cag {
					return nil, _dcf.New("\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0020\u0070\u006f\u0070\u0075\u0070\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
				}
				_fcd.Popup = _fbfb
			}
		}
	}
	if _gggf := _aadc.Get("\u0043\u0041"); _gggf != nil {
		_fcd.CA = _gggf
	}
	if _ddc := _aadc.Get("\u0052\u0043"); _ddc != nil {
		_fcd.RC = _ddc
	}
	if _abgd := _aadc.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"); _abgd != nil {
		_fcd.CreationDate = _abgd
	}
	if _gcfd := _aadc.Get("\u0049\u0052\u0054"); _gcfd != nil {
		_fcd.IRT = _gcfd
	}
	if _gce := _aadc.Get("\u0053\u0075\u0062\u006a"); _gce != nil {
		_fcd.Subj = _gce
	}
	if _edf := _aadc.Get("\u0052\u0054"); _edf != nil {
		_fcd.RT = _edf
	}
	if _fcgd := _aadc.Get("\u0049\u0054"); _fcgd != nil {
		_fcd.IT = _fcgd
	}
	if _ebgfa := _aadc.Get("\u0045\u0078\u0044\u0061\u0074\u0061"); _ebgfa != nil {
		_fcd.ExData = _ebgfa
	}
	return _fcd, nil
}

// PdfAnnotationSquare represents Square annotations.
// (Section 12.5.6.8).
type PdfAnnotationSquare struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _eb.PdfObject
	IC _eb.PdfObject
	BE _eb.PdfObject
	RD _eb.PdfObject
}

// PdfActionSound represents a sound action.
type PdfActionSound struct {
	*PdfAction
	Sound       _eb.PdfObject
	Volume      _eb.PdfObject
	Synchronous _eb.PdfObject
	Repeat      _eb.PdfObject
	Mix         _eb.PdfObject
}

func (_efefb *PdfWriter) optimize() error {
	if _efefb._fbaag == nil {
		return nil
	}
	var _eeedc error
	_efefb._dcfgf, _eeedc = _efefb._fbaag.Optimize(_efefb._dcfgf)
	if _eeedc != nil {
		return _eeedc
	}
	_befdb := make(map[_eb.PdfObject]struct{}, len(_efefb._dcfgf))
	for _, _gddce := range _efefb._dcfgf {
		_befdb[_gddce] = struct{}{}
	}
	_efefb._aeeda = _befdb
	return nil
}

// AcroFormRepairOptions contains options for rebuilding the AcroForm.
type AcroFormRepairOptions struct{}

// PdfRectangle is a definition of a rectangle.
type PdfRectangle struct {
	Llx float64
	Lly float64
	Urx float64
	Ury float64
}

// NewPdfPageResources returns a new PdfPageResources object.
func NewPdfPageResources() *PdfPageResources {
	_egaec := &PdfPageResources{}
	_egaec._fdada = _eb.MakeDict()
	return _egaec
}
func (_bdc *PdfAppender) addNewObject(_adcf _eb.PdfObject) {
	if _, _eaag := _bdc._accfg[_adcf]; !_eaag {
		_bdc._dfg = append(_bdc._dfg, _adcf)
		_bdc._accfg[_adcf] = struct{}{}
	}
}

// Items returns all children outline items.
func (_fbabgd *OutlineItem) Items() []*OutlineItem { return _fbabgd.Entries }
func _edafd(_dddddc *_eb.PdfObjectDictionary) (*PdfShadingType3, error) {
	_dacgca := PdfShadingType3{}
	_ccbbg := _dddddc.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _ccbbg == nil {
		_ddb.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0043\u006f\u006f\u0072\u0064\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_adgbf, _abbef := _ccbbg.(*_eb.PdfObjectArray)
	if !_abbef {
		_ddb.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ccbbg)
		return nil, _eb.ErrTypeError
	}
	if _adgbf.Len() != 6 {
		_ddb.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0036\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _adgbf.Len())
		return nil, ErrInvalidAttribute
	}
	_dacgca.Coords = _adgbf
	if _cfgeg := _dddddc.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _cfgeg != nil {
		_cfgeg = _eb.TraceToDirectObject(_cfgeg)
		_gdacb, _eceac := _cfgeg.(*_eb.PdfObjectArray)
		if !_eceac {
			_ddb.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cfgeg)
			return nil, _eb.ErrTypeError
		}
		_dacgca.Domain = _gdacb
	}
	_ccbbg = _dddddc.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _ccbbg == nil {
		_ddb.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_dacgca.Function = []PdfFunction{}
	if _gcefb, _febbd := _ccbbg.(*_eb.PdfObjectArray); _febbd {
		for _, _fcecf := range _gcefb.Elements() {
			_eeadf, _bacce := _cccfa(_fcecf)
			if _bacce != nil {
				_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bacce)
				return nil, _bacce
			}
			_dacgca.Function = append(_dacgca.Function, _eeadf)
		}
	} else {
		_bfaba, _babff := _cccfa(_ccbbg)
		if _babff != nil {
			_ddb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _babff)
			return nil, _babff
		}
		_dacgca.Function = append(_dacgca.Function, _bfaba)
	}
	if _edgdb := _dddddc.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _edgdb != nil {
		_edgdb = _eb.TraceToDirectObject(_edgdb)
		_bgegf, _edbdd := _edgdb.(*_eb.PdfObjectArray)
		if !_edbdd {
			_ddb.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _edgdb)
			return nil, _eb.ErrTypeError
		}
		if _bgegf.Len() != 2 {
			_ddb.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _bgegf.Len())
			return nil, ErrInvalidAttribute
		}
		_dacgca.Extend = _bgegf
	}
	return &_dacgca, nil
}
func _cffdc(_agddc *_eb.PdfIndirectObject) (*PdfOutline, error) {
	_ffbge, _dacbe := _agddc.PdfObject.(*_eb.PdfObjectDictionary)
	if !_dacbe {
		return nil, _e.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_gedgb := NewPdfOutline()
	if _gbgf := _ffbge.Get("\u0054\u0079\u0070\u0065"); _gbgf != nil {
		_eebefe, _cefgf := _gbgf.(*_eb.PdfObjectName)
		if _cefgf {
			if *_eebefe != "\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073" {
				_ddb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u004f\u0075\u0074l\u0069\u006e\u0065s\u0020(\u0025\u0073\u0029", *_eebefe)
			}
		}
	}
	if _ccafc := _ffbge.Get("\u0043\u006f\u0075n\u0074"); _ccafc != nil {
		_gcff, _cffa := _eb.GetNumberAsInt64(_ccafc)
		if _cffa != nil {
			return nil, _cffa
		}
		_gedgb.Count = &_gcff
	}
	return _gedgb, nil
}
func (_bcddf *LTV) buildCertChain(_gbff, _defb []*_bag.Certificate) ([]*_bag.Certificate, map[string]*_bag.Certificate, error) {
	_gfaba := map[string]*_bag.Certificate{}
	for _, _fcbc := range _gbff {
		_gfaba[_fcbc.Subject.CommonName] = _fcbc
	}
	_ggcbg := _gbff
	for _, _eccef := range _defb {
		_eafbc := _eccef.Subject.CommonName
		if _, _ccfaf := _gfaba[_eafbc]; _ccfaf {
			continue
		}
		_gfaba[_eafbc] = _eccef
		_ggcbg = append(_ggcbg, _eccef)
	}
	if len(_ggcbg) == 0 {
		return nil, nil, ErrSignNoCertificates
	}
	var _bcabf error
	for _ebca := _ggcbg[0]; _ebca != nil && !_bcddf.CertClient.IsCA(_ebca); {
		_fece, _ffdafd := _gfaba[_ebca.Issuer.CommonName]
		if !_ffdafd {
			if _fece, _bcabf = _bcddf.CertClient.GetIssuer(_ebca); _bcabf != nil {
				_ddb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _bcabf)
				break
			}
			_gfaba[_ebca.Issuer.CommonName] = _fece
			_ggcbg = append(_ggcbg, _fece)
		}
		_ebca = _fece
	}
	return _ggcbg, _gfaba, nil
}

// GetContainingPdfObject returns the container of the outline tree node (indirect object).
func (_fafc *PdfOutlineTreeNode) GetContainingPdfObject() _eb.PdfObject {
	return _fafc.GetContext().GetContainingPdfObject()
}

// Direction returns the value of the direction.
func (_gedbb *ViewerPreferences) Direction() Direction { return _gedbb._baeb }
func (_bfd *PdfReader) newPdfAnnotationStrikeOut(_agbg *_eb.PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	_ccc := PdfAnnotationStrikeOut{}
	_eadc, _agg := _bfd.newPdfAnnotationMarkupFromDict(_agbg)
	if _agg != nil {
		return nil, _agg
	}
	_ccc.PdfAnnotationMarkup = _eadc
	_ccc.QuadPoints = _agbg.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_ccc, nil
}
func (_cbad *PdfReader) newPdfActionJavaScriptFromDict(_deg *_eb.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	return &PdfActionJavaScript{JS: _deg.Get("\u004a\u0053")}, nil
}
func (_aebc *PdfReader) newPdfAnnotationCircleFromDict(_daef *_eb.PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	_gfca := PdfAnnotationCircle{}
	_gcg, _cgeb := _aebc.newPdfAnnotationMarkupFromDict(_daef)
	if _cgeb != nil {
		return nil, _cgeb
	}
	_gfca.PdfAnnotationMarkup = _gcg
	_gfca.BS = _daef.Get("\u0042\u0053")
	_gfca.IC = _daef.Get("\u0049\u0043")
	_gfca.BE = _daef.Get("\u0042\u0045")
	_gfca.RD = _daef.Get("\u0052\u0044")
	return &_gfca, nil
}
func (_acbe *PdfReader) newPdfAnnotationMovieFromDict(_cff *_eb.PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	_aaebb := PdfAnnotationMovie{}
	_aaebb.T = _cff.Get("\u0054")
	_aaebb.Movie = _cff.Get("\u004d\u006f\u0076i\u0065")
	_aaebb.A = _cff.Get("\u0041")
	return &_aaebb, nil
}

// Insert adds an outline item as a child of the current outline item,
// at the specified index.
func (_dcff *OutlineItem) Insert(index uint, item *OutlineItem) {
	_dcfaa := uint(len(_dcff.Entries))
	if index > _dcfaa {
		index = _dcfaa
	}
	_dcff.Entries = append(_dcff.Entries[:index], append([]*OutlineItem{item}, _dcff.Entries[index:]...)...)
}

// ToGoTime returns the date in time.Time format.
func (_cgcg PdfDate) ToGoTime() _d.Time {
	_adbcb := int(_cgcg._fgggd*60*60 + _cgcg._gadga*60)
	switch _cgcg._eaede {
	case '-':
		_adbcb = -_adbcb
	case 'Z':
		_adbcb = 0
	}
	_cfagc := _e.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _cgcg._eaede, _cgcg._fgggd, _cgcg._gadga)
	_fgbed := _d.FixedZone(_cfagc, _adbcb)
	return _d.Date(int(_cgcg._fffdf), _d.Month(_cgcg._aacegf), int(_cgcg._afgfb), int(_cgcg._bdfbf), int(_cgcg._aecad), int(_cgcg._cfdb), 0, _fgbed)
}

// NewPdfColorDeviceCMYK returns a new CMYK32 color.
func NewPdfColorDeviceCMYK(c, m, y, k float64) *PdfColorDeviceCMYK {
	_aceg := PdfColorDeviceCMYK{c, m, y, k}
	return &_aceg
}
func (_agbc SignatureValidationResult) String() string {
	var _aabc _dd.Buffer
	_aabc.WriteString(_e.Sprintf("\u004ea\u006d\u0065\u003a\u0020\u0025\u0073\n", _agbc.Name))
	if _agbc.Date._fffdf > 0 {
		_aabc.WriteString(_e.Sprintf("\u0044a\u0074\u0065\u003a\u0020\u0025\u0073\n", _agbc.Date.ToGoTime().String()))
	} else {
		_aabc.WriteString("\u0044\u0061\u0074\u0065 n\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_agbc.Reason) > 0 {
		_aabc.WriteString(_e.Sprintf("R\u0065\u0061\u0073\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _agbc.Reason))
	} else {
		_aabc.WriteString("N\u006f \u0072\u0065\u0061\u0073\u006f\u006e\u0020\u0073p\u0065\u0063\u0069\u0066ie\u0064\u000a")
	}
	if len(_agbc.Location) > 0 {
		_aabc.WriteString(_e.Sprintf("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _agbc.Location))
	} else {
		_aabc.WriteString("\u004c\u006f\u0063at\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_agbc.ContactInfo) > 0 {
		_aabc.WriteString(_e.Sprintf("\u0043\u006f\u006e\u0074\u0061\u0063\u0074\u0020\u0049\u006e\u0066\u006f:\u0020\u0025\u0073\u000a", _agbc.ContactInfo))
	} else {
		_aabc.WriteString("C\u006f\u006e\u0074\u0061\u0063\u0074 \u0069\u006e\u0066\u006f\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063i\u0066i\u0065\u0064\u000a")
	}
	_aabc.WriteString(_e.Sprintf("F\u0069\u0065\u006c\u0064\u0073\u003a\u0020\u0025\u0064\u000a", len(_agbc.Fields)))
	if _agbc.IsSigned {
		_aabc.WriteString("S\u0069\u0067\u006e\u0065\u0064\u003a \u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073\u0020\u0073i\u0067n\u0065\u0064\u000a")
	} else {
		_aabc.WriteString("\u0053\u0069\u0067\u006eed\u003a\u0020\u004e\u006f\u0074\u0020\u0073\u0069\u0067\u006e\u0065\u0064\u000a")
	}
	if _agbc.IsVerified {
		_aabc.WriteString("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0076\u0061\u006c\u0069\u0064\u0061t\u0069\u006f\u006e\u003a\u0020\u0049\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u000a")
	} else {
		_aabc.WriteString("\u0053\u0069\u0067\u006e\u0061\u0074u\u0072\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e:\u0020\u0049\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u000a")
	}
	if _agbc.IsTrusted {
		_aabc.WriteString("\u0054\u0072\u0075\u0073\u0074\u0065\u0064\u003a\u0020\u0043\u0065\u0072\u0074\u0069\u0066i\u0063a\u0074\u0065\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0073\u0074\u0065\u0064\u000a")
	} else {
		_aabc.WriteString("\u0054\u0072\u0075s\u0074\u0065\u0064\u003a \u0055\u006e\u0074\u0072\u0075\u0073\u0074e\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u000a")
	}
	if !_agbc.GeneralizedTime.IsZero() {
		_aabc.WriteString(_e.Sprintf("G\u0065n\u0065\u0072\u0061\u006c\u0069\u007a\u0065\u0064T\u0069\u006d\u0065\u003a %\u0073\u000a", _agbc.GeneralizedTime.String()))
	}
	if _agbc.DiffResults != nil {
		_aabc.WriteString(_e.Sprintf("\u0064\u0069\u0066\u0066 i\u0073\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u003a\u0020\u0025v\u000a", _agbc.DiffResults.IsPermitted()))
		if len(_agbc.DiffResults.Warnings) > 0 {
			_aabc.WriteString("\u004d\u0044\u0050\u0020\u0077\u0061\u0072\u006e\u0069n\u0067\u0073\u003a\u000a")
			for _, _gbeed := range _agbc.DiffResults.Warnings {
				_aabc.WriteString(_e.Sprintf("\u0009\u0025\u0073\u000a", _gbeed))
			}
		}
		if len(_agbc.DiffResults.Errors) > 0 {
			_aabc.WriteString("\u004d\u0044\u0050 \u0065\u0072\u0072\u006f\u0072\u0073\u003a\u000a")
			for _, _efbed := range _agbc.DiffResults.Errors {
				_aabc.WriteString(_e.Sprintf("\u0009\u0025\u0073\u000a", _efbed))
			}
		}
	}
	if _agbc.IsCrlFound {
		_aabc.WriteString("R\u0065\u0076\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0061\u0074\u0061\u003a\u0020C\u0052\u004c\u0020f\u006fu\u006e\u0064\u000a")
	} else {
		_aabc.WriteString("\u0052\u0065\u0076o\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0061\u0074\u0061\u003a \u0043\u0052\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	if _agbc.IsOcspFound {
		_aabc.WriteString("\u0052\u0065\u0076\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u0066\u006f\u0075n\u0064\u000a")
	} else {
		_aabc.WriteString("\u0052\u0065\u0076\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	return _aabc.String()
}

// ImageToRGB converts image in CalGray color space to RGB (A, B, C -> X, Y, Z).
func (_bede *PdfColorspaceCalGray) ImageToRGB(img Image) (Image, error) {
	_dcgdg := _bb.NewReader(img.getBase())
	_gafc := _df.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_aebb := _bb.NewWriter(_gafc)
	_bbaeg := _gg.Pow(2, float64(img.BitsPerComponent)) - 1
	_cabg := make([]uint32, 3)
	var (
		_ebae                                uint32
		ANorm, X, Y, Z, _ffebd, _cebe, _dffe float64
		_caa                                 error
	)
	for {
		_ebae, _caa = _dcgdg.ReadSample()
		if _caa == _bagf.EOF {
			break
		} else if _caa != nil {
			return img, _caa
		}
		ANorm = float64(_ebae) / _bbaeg
		X = _bede.WhitePoint[0] * _gg.Pow(ANorm, _bede.Gamma)
		Y = _bede.WhitePoint[1] * _gg.Pow(ANorm, _bede.Gamma)
		Z = _bede.WhitePoint[2] * _gg.Pow(ANorm, _bede.Gamma)
		_ffebd = 3.240479*X + -1.537150*Y + -0.498535*Z
		_cebe = -0.969256*X + 1.875992*Y + 0.041556*Z
		_dffe = 0.055648*X + -0.204043*Y + 1.057311*Z
		_ffebd = _gg.Min(_gg.Max(_ffebd, 0), 1.0)
		_cebe = _gg.Min(_gg.Max(_cebe, 0), 1.0)
		_dffe = _gg.Min(_gg.Max(_dffe, 0), 1.0)
		_cabg[0] = uint32(_ffebd * _bbaeg)
		_cabg[1] = uint32(_cebe * _bbaeg)
		_cabg[2] = uint32(_dffe * _bbaeg)
		if _caa = _aebb.WriteSamples(_cabg); _caa != nil {
			return img, _caa
		}
	}
	return _ggaa(&_gafc), nil
}

// ToPdfObject implements interface PdfModel.
func (_bcec *PdfAnnotationCaret) ToPdfObject() _eb.PdfObject {
	_bcec.PdfAnnotation.ToPdfObject()
	_cacd := _bcec._ggf
	_fcdg := _cacd.PdfObject.(*_eb.PdfObjectDictionary)
	_bcec.PdfAnnotationMarkup.appendToPdfDictionary(_fcdg)
	_fcdg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _eb.MakeName("\u0043\u0061\u0072e\u0074"))
	_fcdg.SetIfNotNil("\u0052\u0044", _bcec.RD)
	_fcdg.SetIfNotNil("\u0053\u0079", _bcec.Sy)
	return _cacd
}

// GetPageLabels returns the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_fefbg *PdfReader) GetPageLabels() (_eb.PdfObject, error) {
	_bdbg := _eb.ResolveReference(_fefbg._bagcfd.Get("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073"))
	if _bdbg == nil {
		return nil, nil
	}
	if !_fefbg._cfcgdf {
		_dbgdg := _fefbg.traverseObjectData(_bdbg)
		if _dbgdg != nil {
			return nil, _dbgdg
		}
	}
	return _bdbg, nil
}

// GetOutlineTree returns the outline tree.
func (_ecbb *PdfReader) GetOutlineTree() *PdfOutlineTreeNode { return _ecbb._efabg }

// PickTrayByPDFSize returns the value of the pickTrayByPDFSize flag.
func (_acebf *ViewerPreferences) PickTrayByPDFSize() bool {
	if _acebf._gggea == nil {
		return false
	}
	return *_acebf._gggea
}

// VariableText contains the common attributes of a variable text.
// The VariableText is typically not used directly, but is can encapsulate by PdfField
// See section 12.7.3.3 "Variable Text" and Table 222 (pp. 434-436 PDF32000_2008).
type VariableText struct {
	DA *_eb.PdfObjectString
	Q  *_eb.PdfObjectInteger
	DS *_eb.PdfObjectString
	RV _eb.PdfObject
}

// PdfActionResetForm represents a resetForm action.
type PdfActionResetForm struct {
	*PdfAction
	Fields _eb.PdfObject
	Flags  _eb.PdfObject
}

// StandardValidator is the interface that is used for the PDF StandardImplementer validation for the PDF document.
// It is using a CompliancePdfReader which is expected to give more Metadata during reading process.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardValidator interface {

	// ValidateStandard checks if the input reader
	ValidateStandard(_gdddd *CompliancePdfReader) error
}
