package mdp

import (
	_b "errors"
	_f "fmt"

	_c "github.com/bamzi/pdfext/core"
)

// IsPermitted returns true if changes permitted.
func (_adb *DiffResults) IsPermitted() bool { return len(_adb.Errors) == 0 }
func NewDefaultDiffPolicy() DiffPolicy {
	return &defaultDiffPolicy{_d: nil, _e: &DiffResults{}, _da: 0}
}

// String returns the state of the warning.
func (_edb *DiffResult) String() string {
	return _f.Sprintf("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064", _edb.Description, _edb.Revision)
}
func _eeb(_dgd _c.PdfObject) ([]_c.PdfObject, error) {
	_dda := make([]_c.PdfObject, 0)
	if _dgd != nil {
		_fbd := _dgd
		if _eee, _cbdc := _c.GetIndirect(_dgd); _cbdc {
			_fbd = _eee.PdfObject
		}
		if _edg, _fgg := _c.GetArray(_fbd); _fgg {
			_dda = _edg.Elements()
		} else {
			return nil, _b.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
	}
	return _dda, nil
}

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_eb *defaultDiffPolicy) ReviewFile(oldParser *_c.PdfParser, newParser *_c.PdfParser, params *MDPParameters) (*DiffResults, error) {
	if oldParser.GetRevisionNumber() > newParser.GetRevisionNumber() {
		return nil, _b.New("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e")
	}
	if oldParser.GetRevisionNumber() == newParser.GetRevisionNumber() {
		if oldParser != newParser {
			return nil, _b.New("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073")
		}
		return &DiffResults{}, nil
	}
	if params == nil {
		_eb._da = NoRestrictions
	} else {
		_eb._da = params.DocMDPLevel
	}
	_ba := &DiffResults{}
	for _bg := oldParser.GetRevisionNumber() + 1; _bg <= newParser.GetRevisionNumber(); _bg++ {
		_a, _gf := newParser.GetRevision(_bg - 1)
		if _gf != nil {
			return nil, _gf
		}
		_dg, _gf := newParser.GetRevision(_bg)
		if _gf != nil {
			return nil, _gf
		}
		_bga, _gf := _eb.compareRevisions(_a, _dg)
		if _gf != nil {
			return nil, _gf
		}
		_ba.Warnings = append(_ba.Warnings, _bga.Warnings...)
		_ba.Errors = append(_ba.Errors, _bga.Errors...)
	}
	return _ba, nil
}
func (_gc *defaultDiffPolicy) compareFields(_abe int, _baa, _edd []_c.PdfObject) error {
	_bd := make(map[int64]*_c.PdfObjectDictionary)
	for _, _gg := range _baa {
		_gcc, _cd := _c.GetIndirect(_gg)
		if !_cd {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_fd, _cd := _c.GetDict(_gcc.PdfObject)
		if !_cd {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_bd[_gcc.ObjectNumber] = _fd
	}
	for _, _bgc := range _edd {
		_ae, _ef := _c.GetIndirect(_bgc)
		if !_ef {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_de, _ef := _c.GetDict(_ae.PdfObject)
		if !_ef {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		T := _de.Get("\u0054")
		if _, _afa := _gc._d[_ae.ObjectNumber]; _afa {
			switch _gc._da {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_gc._e.addWarningWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			default:
				_gc._e.addErrorWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			}
		}
		if _, _ga := _bd[_ae.ObjectNumber]; !_ga {
			switch _gc._da {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_gc._e.addWarningWithDescription(_abe, _f.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _de.Get("\u0054")))
			default:
				_gc._e.addErrorWithDescription(_abe, _f.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _de.Get("\u0054")))
			}
		} else {
			delete(_bd, _ae.ObjectNumber)
			if _, _adf := _gc._d[_ae.ObjectNumber]; _adf {
				switch _gc._da {
				case NoRestrictions, FillForms, FillFormsAndAnnots:
					_gc._e.addWarningWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _de.Get("\u0054")))
				default:
					_gc._e.addErrorWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _de.Get("\u0054")))
				}
			}
		}
		if FT, _fge := _c.GetNameVal(_de.Get("\u0046\u0054")); _fge {
			if FT == "\u0053\u0069\u0067" {
				if _gag, _feg := _c.GetIndirect(_de.Get("\u0056")); _feg {
					if _, _dee := _gc._d[_gag.ObjectNumber]; _dee {
						switch _gc._da {
						case NoRestrictions, FillForms, FillFormsAndAnnots:
							_gc._e.addWarningWithDescription(_abe, _f.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						default:
							_gc._e.addErrorWithDescription(_abe, _f.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						}
					}
				}
			}
		}
	}
	for _, _dd := range _bd {
		switch _gc._da {
		case NoRestrictions:
			_gc._e.addWarningWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _dd.Get("\u0054")))
		default:
			_gc._e.addErrorWithDescription(_abe, _f.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _dd.Get("\u0054")))
		}
	}
	return nil
}

type defaultDiffPolicy struct {
	_d  map[int64]_c.PdfObject
	_e  *DiffResults
	_da DocMDPPermission
}

func (_gfc *defaultDiffPolicy) compareRevisions(_db *_c.PdfParser, _cb *_c.PdfParser) (*DiffResults, error) {
	var _af error
	_gfc._d, _af = _cb.GetUpdatedObjects(_db)
	if _af != nil {
		return &DiffResults{}, _af
	}
	if len(_gfc._d) == 0 {
		return &DiffResults{}, nil
	}
	_cg := _cb.GetRevisionNumber()
	_ee, _bab := _c.GetIndirect(_c.ResolveReference(_db.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	_ge, _cgb := _c.GetIndirect(_c.ResolveReference(_cb.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	if !_bab || !_cgb {
		return &DiffResults{}, _b.New("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_gd, _bab := _c.GetDict(_c.ResolveReference(_ee.PdfObject))
	_bc, _cgb := _c.GetDict(_c.ResolveReference(_ge.PdfObject))
	if !_bab || !_cgb {
		return &DiffResults{}, _b.New("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
	}
	if _fb, _fbc := _c.GetIndirect(_bc.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _fbc {
		_fg, _cge := _c.GetDict(_fb)
		if !_cge {
			return &DiffResults{}, _b.New("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_fga := make([]_c.PdfObject, 0)
		if _ab, _ad := _c.GetIndirect(_gd.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _ad {
			if _cbd, _ed := _c.GetDict(_ab); _ed {
				if _geb, _aa := _c.GetArray(_cbd.Get("\u0046\u0069\u0065\u006c\u0064\u0073")); _aa {
					_fga = _geb.Elements()
				}
			}
		}
		_eg, _cge := _c.GetArray(_fg.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_cge {
			return &DiffResults{}, _b.New("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073")
		}
		if _ca := _gfc.compareFields(_cg, _fga, _eg.Elements()); _ca != nil {
			return &DiffResults{}, _ca
		}
	}
	_cc, _egd := _c.GetIndirect(_bc.Get("\u0050\u0061\u0067e\u0073"))
	if !_egd {
		return &DiffResults{}, _b.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_dbb, _egd := _c.GetIndirect(_gd.Get("\u0050\u0061\u0067e\u0073"))
	if !_egd {
		return &DiffResults{}, _b.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	if _cae := _gfc.comparePages(_cg, _dbb, _cc); _cae != nil {
		return &DiffResults{}, _cae
	}
	return _gfc._e, nil
}

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct {
	Revision    int
	Description string
}

func (_cab *DiffResults) addError(_aae *DiffResult) {
	if _cab.Errors == nil {
		_cab.Errors = make([]*DiffResult, 0)
	}
	_cab.Errors = append(_cab.Errors, _aae)
}
func (_dc *defaultDiffPolicy) comparePages(_gee int, _aec, _fgag *_c.PdfIndirectObject) error {
	if _, _bb := _dc._d[_fgag.ObjectNumber]; _bb {
		_dc._e.addErrorWithDescription(_gee, "\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064")
	}
	_bf, _aed := _c.GetDict(_fgag.PdfObject)
	_be, _eed := _c.GetDict(_aec.PdfObject)
	if !_aed || !_eed {
		return _b.New("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_gca, _aed := _c.GetArray(_bf.Get("\u004b\u0069\u0064\u0073"))
	_dbd, _eed := _c.GetArray(_be.Get("\u004b\u0069\u0064\u0073"))
	if !_aed || !_eed {
		return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_df := _gca.Len()
	if _df > _dbd.Len() {
		_df = _dbd.Len()
	}
	for _ea := 0; _ea < _df; _ea++ {
		_cage, _afb := _c.GetIndirect(_c.ResolveReference(_dbd.Get(_ea)))
		_ccf, _dbde := _c.GetIndirect(_c.ResolveReference(_gca.Get(_ea)))
		if !_afb || !_dbde {
			return _b.New("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074")
		}
		if _cage.ObjectNumber != _ccf.ObjectNumber {
			_dc._e.addErrorWithDescription(_gee, _f.Sprintf("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064", _ea))
		}
		_ec, _afb := _c.GetDict(_ccf)
		_gec, _dbde := _c.GetDict(_cage)
		if !_afb || !_dbde {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
		}
		_bdd, _abc := _eeb(_ec.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _abc != nil {
			return _abc
		}
		_ada, _abc := _eeb(_gec.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _abc != nil {
			return _abc
		}
		if _bda := _dc.compareAnnots(_gee, _ada, _bdd); _bda != nil {
			return _bda
		}
	}
	for _fec := _df + 1; _fec <= _gca.Len(); _fec++ {
		_dc._e.addErrorWithDescription(_gee, _f.Sprintf("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _fec))
	}
	for _dcb := _df + 1; _dcb <= _dbd.Len(); _dcb++ {
		_dc._e.addErrorWithDescription(_gee, _f.Sprintf("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _dcb))
	}
	return nil
}

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{ DocMDPLevel DocMDPPermission }

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64

func (_dba *DiffResults) addWarning(_baf *DiffResult) {
	if _dba.Warnings == nil {
		_dba.Warnings = make([]*DiffResult, 0)
	}
	_dba.Warnings = append(_dba.Warnings, _baf)
}
func (_fbb *DiffResults) addErrorWithDescription(_fdc int, _eff string) {
	if _fbb.Errors == nil {
		_fbb.Errors = make([]*DiffResult, 0)
	}
	_fbb.Errors = append(_fbb.Errors, &DiffResult{Revision: _fdc, Description: _eff})
}

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct {
	Warnings []*DiffResult
	Errors   []*DiffResult
}

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface {

	// ReviewFile should check the revisions of the old and new parsers
	// and evaluate the differences between the revisions.
	// Each implementation of this interface must decide
	// how to handle cases where there are multiple revisions between the old and new revisions.
	ReviewFile(_eedb *_c.PdfParser, _aab *_c.PdfParser, _gbe *MDPParameters) (*DiffResults, error)
}

func (_bbe *DiffResults) addWarningWithDescription(_cfd int, _ega string) {
	if _bbe.Warnings == nil {
		_bbe.Warnings = make([]*DiffResult, 0)
	}
	_bbe.Warnings = append(_bbe.Warnings, &DiffResult{Revision: _cfd, Description: _ega})
}

const (
	NoRestrictions     DocMDPPermission = 0
	NoChanges          DocMDPPermission = 1
	FillForms          DocMDPPermission = 2
	FillFormsAndAnnots DocMDPPermission = 3
)

func (_deee *defaultDiffPolicy) compareAnnots(_gaga int, _gbd, _egg []_c.PdfObject) error {
	_ccb := make(map[int64]*_c.PdfObjectDictionary)
	for _, _cf := range _gbd {
		_aaf, _cad := _c.GetIndirect(_cf)
		if !_cad {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cbf, _cad := _c.GetDict(_aaf.PdfObject)
		if !_cad {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_ccb[_aaf.ObjectNumber] = _cbf
	}
	for _, _bee := range _egg {
		_agd, _abg := _c.GetIndirect(_bee)
		if !_abg {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_bdb, _abg := _c.GetDict(_agd.PdfObject)
		if !_abg {
			return _b.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_gfeg, _ := _c.GetStringVal(_bdb.Get("\u0054"))
		_bgac, _ := _c.GetNameVal(_bdb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if _, _fef := _ccb[_agd.ObjectNumber]; !_fef {
			switch _deee._da {
			case NoRestrictions, FillFormsAndAnnots:
				_deee._e.addWarningWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bgac, _gfeg))
			default:
				_aecd, _fed := _c.GetDict(_agd.PdfObject)
				if !_fed {
					return _b.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
				}
				_gba, _fed := _c.GetNameVal(_aecd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
				if !_fed {
					return _b.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
				}
				if _gba == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					switch _deee._da {
					case NoRestrictions, FillFormsAndAnnots, FillForms:
						_deee._e.addWarningWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bgac, _gfeg))
					default:
						_deee._e.addErrorWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bgac, _gfeg))
					}
				} else {
					_deee._e.addErrorWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bgac, _gfeg))
				}
			}
		} else {
			delete(_ccb, _agd.ObjectNumber)
			if _fgee, _gaf := _deee._d[_agd.ObjectNumber]; _gaf {
				switch _deee._da {
				case NoRestrictions, FillFormsAndAnnots:
					_deee._e.addWarningWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bgac, _gfeg))
				default:
					_ded, _bcg := _c.GetIndirect(_fgee)
					if !_bcg {
						return _b.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_dfff, _bcg := _c.GetDict(_ded.PdfObject)
					if !_bcg {
						return _b.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_fegb, _bcg := _c.GetNameVal(_dfff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
					if !_bcg {
						return _b.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
					}
					if _fegb == "\u0057\u0069\u0064\u0067\u0065\u0074" {
						switch _deee._da {
						case NoRestrictions, FillFormsAndAnnots, FillForms:
							_deee._e.addWarningWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bgac, _gfeg))
						default:
							_deee._e.addErrorWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bgac, _gfeg))
						}
					} else {
						_deee._e.addErrorWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bgac, _gfeg))
					}
				}
			}
		}
	}
	for _, _cff := range _ccb {
		_bce, _ := _c.GetStringVal(_cff.Get("\u0054"))
		_bgb, _ := _c.GetNameVal(_cff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		switch _deee._da {
		case NoRestrictions, FillFormsAndAnnots:
			_deee._e.addWarningWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _bgb, _bce))
		default:
			_deee._e.addErrorWithDescription(_gaga, _f.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _bgb, _bce))
		}
	}
	return nil
}
