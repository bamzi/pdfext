// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil

import (
	_g "errors"
	_fa "fmt"
	_c "strconv"
	_a "time"

	_bb "github.com/bamzi/pdfext/core"
	_f "github.com/bamzi/pdfext/internal/timeutils"
	_fd "github.com/bamzi/pdfext/internal/uuid"
	_b "github.com/bamzi/pdfext/model/xmputil/pdfaextension"
	_dbg "github.com/bamzi/pdfext/model/xmputil/pdfaid"
	_ce "github.com/trimmer-io/go-xmp/models/pdf"
	_ae "github.com/trimmer-io/go-xmp/models/xmp_base"
	_e "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_db "github.com/trimmer-io/go-xmp/xmp"
)

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_bd *Document) GetMediaManagement() (*MediaManagement, bool) {
	_ea := _e.FindModel(_bd._fde)
	if _ea == nil {
		return nil, false
	}
	_bce := make([]MediaManagementVersion, len(_ea.Versions))
	for _aeb, _dfd := range _ea.Versions {
		_bce[_aeb] = MediaManagementVersion{VersionID: _dfd.Version, ModifyDate: _dfd.ModifyDate.Value(), Comments: _dfd.Comments, Modifier: _dfd.Modifier}
	}
	_dg := &MediaManagement{OriginalDocumentID: GUID(_ea.OriginalDocumentID.Value()), DocumentID: GUID(_ea.DocumentID.Value()), InstanceID: GUID(_ea.InstanceID.Value()), VersionID: _ea.VersionID, Versions: _bce}
	if _ea.DerivedFrom != nil {
		_dg.DerivedFrom = &MediaManagementDerivedFrom{OriginalDocumentID: GUID(_ea.DerivedFrom.OriginalDocumentID), DocumentID: GUID(_ea.DerivedFrom.DocumentID), InstanceID: GUID(_ea.DerivedFrom.InstanceID), VersionID: _ea.DerivedFrom.VersionID}
	}
	return _dg, true
}

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_fef *Document) SetPdfAID(part int, conformance string) error {
	_gbe, _ab := _dbg.MakeModel(_fef._fde)
	if _ab != nil {
		return _ab
	}
	_gbe.Part = part
	_gbe.Conformance = conformance
	if _ddg := _gbe.SyncToXMP(_fef._fde); _ddg != nil {
		return _ddg
	}
	return nil
}

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct {
	OriginalDocumentID GUID
	DocumentID         GUID
	InstanceID         GUID
	VersionID          string
}

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct {
	InfoDict   _bb.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool

	// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
	Overwrite bool
}

// Marshal the document into xml byte stream.
func (_aef *Document) Marshal() ([]byte, error) {
	if _aef._fde.IsDirty() {
		if _fb := _aef._fde.SyncModels(); _fb != nil {
			return nil, _fb
		}
	}
	return _db.Marshal(_aef._fde)
}

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_dbc *Document) SetPdfAExtension() error {
	_bg, _cda := _b.MakeModel(_dbc._fde)
	if _cda != nil {
		return _cda
	}
	if _cda = _b.FillModel(_dbc._fde, _bg); _cda != nil {
		return _cda
	}
	if _cda = _bg.SyncToXMP(_dbc._fde); _cda != nil {
		return _cda
	}
	return nil
}

// NewDocument creates a new document without any previous xmp information.
func NewDocument() *Document { _eb := _db.NewDocument(); return &Document{_fde: _eb} }

// SetPdfInfo sets the pdf info into selected document.
func (_gg *Document) SetPdfInfo(options *PdfInfoOptions) error {
	if options == nil {
		return _g.New("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_eed, _fg := _ce.MakeModel(_gg._fde)
	if _fg != nil {
		return _fg
	}
	if options.Overwrite {
		*_eed = _ce.PDFInfo{}
	}
	if options.InfoDict != nil {
		_cc, _cgc := _bb.GetDict(options.InfoDict)
		if !_cgc {
			return _fa.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", options.InfoDict)
		}
		var _bf *_bb.PdfObjectString
		for _, _ag := range _cc.Keys() {
			switch _ag {
			case "\u0054\u0069\u0074l\u0065":
				_bf, _cgc = _bb.GetString(_cc.Get("\u0054\u0069\u0074l\u0065"))
				if _cgc {
					_eed.Title = _db.NewAltString(_bf)
				}
			case "\u0041\u0075\u0074\u0068\u006f\u0072":
				_bf, _cgc = _bb.GetString(_cc.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
				if _cgc {
					_eed.Author = _db.NewStringList(_bf.String())
				}
			case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
				_bf, _cgc = _bb.GetString(_cc.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
				if _cgc {
					_eed.Keywords = _bf.String()
				}
			case "\u0043r\u0065\u0061\u0074\u006f\u0072":
				_bf, _cgc = _bb.GetString(_cc.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
				if _cgc {
					_eed.Creator = _db.AgentName(_bf.String())
				}
			case "\u0053u\u0062\u006a\u0065\u0063\u0074":
				_bf, _cgc = _bb.GetString(_cc.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
				if _cgc {
					_eed.Subject = _db.NewAltString(_bf.String())
				}
			case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
				_bf, _cgc = _bb.GetString(_cc.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
				if _cgc {
					_eed.Producer = _db.AgentName(_bf.String())
				}
			case "\u0054r\u0061\u0070\u0070\u0065\u0064":
				_ac, _da := _bb.GetName(_cc.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
				if _da {
					switch _ac.String() {
					case "\u0054\u0072\u0075\u0065":
						_eed.Trapped = true
					case "\u0046\u0061\u006cs\u0065":
						_eed.Trapped = false
					default:
						_eed.Trapped = true
					}
				}
			case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
				if _cgcg, _cce := _bb.GetString(_cc.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _cce && _cgcg.String() != "" {
					_ad, _gb := _f.ParsePdfTime(_cgcg.String())
					if _gb != nil {
						return _fa.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _gb)
					}
					_eed.CreationDate = _db.NewDate(_ad)
				}
			case "\u004do\u0064\u0044\u0061\u0074\u0065":
				if _bc, _gba := _bb.GetString(_cc.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _gba && _bc.String() != "" {
					_gf, _cf := _f.ParsePdfTime(_bc.String())
					if _cf != nil {
						return _fa.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _cf)
					}
					_eed.ModifyDate = _db.NewDate(_gf)
				}
			}
		}
	}
	if options.PdfVersion != "" {
		_eed.PDFVersion = options.PdfVersion
	}
	if options.Marked {
		_eed.Marked = _db.Bool(options.Marked)
	}
	if options.Copyright != "" {
		_eed.Copyright = options.Copyright
	}
	if _fg = _eed.SyncToXMP(_gg._fde); _fg != nil {
		return _fg
	}
	return nil
}

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct {
	VersionID  string
	ModifyDate _a.Time
	Comments   string
	Modifier   string
}

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct {
	Part        int
	Conformance string
}

// GetPdfInfo gets the document pdf info.
func (_be *Document) GetPdfInfo() (*PdfInfo, bool) {
	_adb := PdfInfo{}
	var _aga *_bb.PdfObjectDictionary
	_de := func(_bba string, _df _bb.PdfObject) {
		if _aga == nil {
			_aga = _bb.MakeDict()
		}
		_aga.Set(_bb.PdfObjectName(_bba), _df)
	}
	_cdf, _fee := _be._fde.FindModel(_ce.NsPDF).(*_ce.PDFInfo)
	if !_fee {
		_eba, _bee := _be._fde.FindModel(_ae.NsXmp).(*_ae.XmpBase)
		if !_bee {
			return nil, false
		}
		if _eba.CreatorTool != "" {
			_de("\u0043r\u0065\u0061\u0074\u006f\u0072", _bb.MakeString(string(_eba.CreatorTool)))
		}
		if !_eba.CreateDate.IsZero() {
			_de("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _bb.MakeString(_f.FormatPdfTime(_eba.CreateDate.Value())))
		}
		if !_eba.ModifyDate.IsZero() {
			_de("\u004do\u0064\u0044\u0061\u0074\u0065", _bb.MakeString(_f.FormatPdfTime(_eba.ModifyDate.Value())))
		}
		_adb.InfoDict = _aga
		return &_adb, true
	}
	_adb.Copyright = _cdf.Copyright
	_adb.PdfVersion = _cdf.PDFVersion
	_adb.Marked = bool(_cdf.Marked)
	if len(_cdf.Title) > 0 {
		_de("\u0054\u0069\u0074l\u0065", _bb.MakeString(_cdf.Title.Default()))
	}
	if len(_cdf.Author) > 0 {
		_de("\u0041\u0075\u0074\u0068\u006f\u0072", _bb.MakeString(_cdf.Author[0]))
	}
	if _cdf.Keywords != "" {
		_de("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _bb.MakeString(_cdf.Keywords))
	}
	if len(_cdf.Subject) > 0 {
		_de("\u0053u\u0062\u006a\u0065\u0063\u0074", _bb.MakeString(_cdf.Subject.Default()))
	}
	if _cdf.Creator != "" {
		_de("\u0043r\u0065\u0061\u0074\u006f\u0072", _bb.MakeString(string(_cdf.Creator)))
	}
	if _cdf.Producer != "" {
		_de("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _bb.MakeString(string(_cdf.Producer)))
	}
	if _cdf.Trapped {
		_de("\u0054r\u0061\u0070\u0070\u0065\u0064", _bb.MakeName("\u0054\u0072\u0075\u0065"))
	}
	if !_cdf.CreationDate.IsZero() {
		_de("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _bb.MakeString(_f.FormatPdfTime(_cdf.CreationDate.Value())))
	}
	if !_cdf.ModifyDate.IsZero() {
		_de("\u004do\u0064\u0044\u0061\u0074\u0065", _bb.MakeString(_f.FormatPdfTime(_cdf.ModifyDate.Value())))
	}
	_adb.InfoDict = _aga
	return &_adb, true
}

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{ _fde *_db.Document }

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_afa *Document) SetMediaManagement(options *MediaManagementOptions) error {
	_ced, _ffa := _e.MakeModel(_afa._fde)
	if _ffa != nil {
		return _ffa
	}
	if options == nil {
		options = new(MediaManagementOptions)
	}
	_add := _e.ResourceRef{}
	switch {
	case options.DocumentID != "":
		_ced.DocumentID = _db.GUID(options.DocumentID)
	case options.NewDocumentID || _ced.DocumentID.IsZero():
		if !_ced.DocumentID.IsZero() {
			_add.DocumentID = _ced.DocumentID
		}
		_cde, _bfd := _fd.NewUUID()
		if _bfd != nil {
			return _bfd
		}
		_ced.DocumentID = _db.GUID(_cde.String())
	}
	if !_ced.InstanceID.IsZero() {
		_add.InstanceID = _ced.InstanceID
	}
	_ced.InstanceID = _db.GUID(options.InstanceID)
	if _ced.InstanceID == "" {
		_bed, _cfd := _fd.NewUUID()
		if _cfd != nil {
			return _cfd
		}
		_ced.InstanceID = _db.GUID(_bed.String())
	}
	if !_add.IsZero() {
		_ced.DerivedFrom = &_add
	}
	_ebc := options.VersionID
	if _ced.VersionID != "" {
		_acd, _efe := _c.Atoi(_ced.VersionID)
		if _efe != nil {
			_ebc = _c.Itoa(len(_ced.Versions) + 1)
		} else {
			_ebc = _c.Itoa(_acd + 1)
		}
	}
	if _ebc == "" {
		_ebc = "\u0031"
	}
	_ced.VersionID = _ebc
	if _ffa = _ced.SyncToXMP(_afa._fde); _ffa != nil {
		return _ffa
	}
	return nil
}

// GUID is a string representing a globally unique identifier.
type GUID string

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_cga *Document) GetGoXmpDocument() *_db.Document { return _cga._fde }

// PdfInfo is the xmp document pdf info.
type PdfInfo struct {
	InfoDict   _bb.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool
}

// GetPdfAID gets the pdfaid xmp metadata model.
func (_daf *Document) GetPdfAID() (*PdfAID, bool) {
	_ggc, _ga := _daf._fde.FindModel(_dbg.Namespace).(*_dbg.Model)
	if !_ga {
		return nil, false
	}
	return &PdfAID{Part: _ggc.Part, Conformance: _ggc.Conformance}, true
}

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument(stream []byte) (*Document, error) {
	_ff := _db.NewDocument()
	if _ef := _db.Unmarshal(stream, _ff); _ef != nil {
		return nil, _ef
	}
	return &Document{_fde: _ff}, nil
}

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct {

	// OriginalDocumentID  as media is imported and projects is started, an original-document ID
	// must be created to identify a new document. This identifies a document as a conceptual entity.
	OriginalDocumentID GUID

	// DocumentID when a document is copied to a new file path or converted to a new format with
	// Save As, another new document ID should usually be assigned. This identifies a general version or
	// branch of a document. You can use it to track different versions or extracted portions of a document
	// with the same original-document ID.
	DocumentID GUID

	// InstanceID to track a document’s editing history, you must assign a new instance ID
	// whenever a document is saved after any changes. This uniquely identifies an exact version of a
	// document. It is used in resource references (to identify both the document or part itself and the
	// referenced or referencing documents), and in document-history resource events (to identify the
	// document instance that resulted from the change).
	InstanceID GUID

	// DerivedFrom references the source document from which this one is derived,
	// typically through a Save As operation that changes the file name or format. It is a minimal reference;
	// missing components can be assumed to be unchanged. For example, a new version might only need
	// to specify the instance ID and version number of the previous version, or a rendition might only need
	// to specify the instance ID and rendition class of the original.
	DerivedFrom *MediaManagementDerivedFrom

	// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
	// document history, but should not be used to identify a document uniquely in any context.
	// Usually it simply works by incrementing integers 1,2,3...
	VersionID string

	// Versions is the history of the document versions along with the comments, timestamps and issuers.
	Versions []MediaManagementVersion
}

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct {

	// OriginalDocumentID  as media is imported and projects is started, an original-document ID
	// must be created to identify a new document. This identifies a document as a conceptual entity.
	// By default, this value is generated.
	OriginalDocumentID string

	// NewDocumentID is a flag which generates a new Document identifier while setting media management.
	// This value should be set to true only if the document is stored and saved as new document.
	// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
	NewDocumentID bool

	// DocumentID when a document is copied to a new file path or converted to a new format with
	// Save As, another new document ID should usually be assigned. This identifies a general version or
	// branch of a document. You can use it to track different versions or extracted portions of a document
	// with the same original-document ID.
	// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
	DocumentID string

	// InstanceID to track a document’s editing history, you must assign a new instance ID
	// whenever a document is saved after any changes. This uniquely identifies an exact version of a
	// document. It is used in resource references (to identify both the document or part itself and the
	// referenced or referencing documents), and in document-history resource events (to identify the
	// document instance that resulted from the change).
	// By default, this value is generated.
	InstanceID string

	// DerivedFrom references the source document from which this one is derived,
	// typically through a Save As operation that changes the file name or format. It is a minimal reference;
	// missing components can be assumed to be unchanged. For example, a new version might only need
	// to specify the instance ID and version number of the previous version, or a rendition might only need
	// to specify the instance ID and rendition class of the original.
	// By default, the derived from structure is filled from previous XMP metadata (if exists).
	DerivedFrom string

	// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
	// document history, but should not be used to identify a document uniquely in any context.
	// Usually it simply works by incrementing integers 1,2,3...
	// By default, this values is incremented or set to the next version number.
	VersionID string

	// ModifyComment is a comment to given modification
	ModifyComment string

	// ModifyDate is a custom modification date for the versions.
	// By default, this would be set to time.Now().
	ModifyDate _a.Time

	// Modifier is a person who did the modification.
	Modifier string
}

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_cd *Document) MarshalIndent(prefix, indent string) ([]byte, error) {
	if _cd._fde.IsDirty() {
		if _cg := _cd._fde.SyncModels(); _cg != nil {
			return nil, _cg
		}
	}
	return _db.MarshalIndent(_cd._fde, prefix, indent)
}

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_ee *Document) GetPdfaExtensionSchemas() ([]_b.Schema, error) {
	_fe := _ee._fde.FindModel(_b.Namespace)
	if _fe == nil {
		return nil, nil
	}
	_ebd, _dd := _fe.(*_b.Model)
	if !_dd {
		return nil, _fa.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054", _fe)
	}
	return _ebd.Schemas, nil
}
