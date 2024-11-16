// Package extractor is used for quickly extracting PDF content through a simple interface.
// Currently offers functionality for extracting textual content.
package extractor

import (
	_aa "bytes"
	_dg "errors"
	_d "fmt"
	_f "image"
	_ab "image/color"
	_e "io"
	_bb "math"
	_af "reflect"
	_gb "regexp"
	_gd "sort"
	_ge "strings"
	_ae "unicode"
	_a "unicode/utf8"

	_b "github.com/bamzi/pdfext/common"
	_aed "github.com/bamzi/pdfext/contentstream"
	_gbc "github.com/bamzi/pdfext/core"
	_fe "github.com/bamzi/pdfext/internal/license"
	_abg "github.com/bamzi/pdfext/internal/textencoding"
	_ef "github.com/bamzi/pdfext/internal/transform"
	_bg "github.com/bamzi/pdfext/model"
	_ca "golang.org/x/image/draw"
	_g "golang.org/x/text/unicode/norm"
)

func (_dafe *shapesState) devicePoint(_eeag, _dda float64) _ef.Point {
	_cgfg := _dafe._gbag.Mult(_dafe._bdab)
	_eeag, _dda = _cgfg.Transform(_eeag, _dda)
	return _ef.NewPoint(_eeag, _dda)
}

const _degf = 1.0 / 1000.0

func _bbf(_dbdd []*textLine) map[float64][]*textLine {
	_gd.Slice(_dbdd, func(_bgec, _eafe int) bool { return _dbdd[_bgec]._ggfc < _dbdd[_eafe]._ggfc })
	_dagf := map[float64][]*textLine{}
	for _, _bdbe := range _dbdd {
		_baag := _dcde(_bdbe)
		_baag = _bb.Round(_baag)
		_dagf[_baag] = append(_dagf[_baag], _bdbe)
	}
	return _dagf
}

// ExtractPageText returns the text contents of `e` (an Extractor for a page) as a PageText.
// TODO(peterwilliams97): The stats complicate this function signature and aren't very useful.
//
//	Replace with a function like Extract() (*PageText, error)
func (_ebff *Extractor) ExtractPageText() (*PageText, int, int, error) {
	_gacb, _abdf, _ccbf, _adgb := _ebff.extractPageText(_ebff._eca, _ebff._feg, _ef.IdentityMatrix(), 0, false)
	if _adgb != nil && _adgb != _bg.ErrColorOutOfRange {
		return nil, 0, 0, _adgb
	}
	if _ebff._dff != nil {
		_gacb._gfee._ddcf = _ebff._dff.UseSimplerExtractionProcess
	}
	_gacb.computeViews()
	//------------------- fmt.Println("---v check ")
	// _adgb = _fccc(_gacb)
	if _adgb != nil {
		return nil, 0, 0, _adgb
	}
	if _ebff._dff != nil {
		if _ebff._dff.ApplyCropBox && _ebff._fcb != nil {
			_gacb.ApplyArea(*_ebff._fcb)
		}
		_gacb._gfee._abdc = _ebff._dff.DisableDocumentTags
	}
	return _gacb, _abdf, _ccbf, nil
}
func _dfde(_aefc structElement) []structElement {
	_ageg := []structElement{}
	for _, _edcfe := range _aefc._fbcba {
		for _, _cadfd := range _edcfe._fbcba {
			for _, _eacc := range _cadfd._fbcba {
				if _eacc._adfb == "\u004c" {
					_ageg = append(_ageg, _eacc)
				}
			}
		}
	}
	return _ageg
}

// NewFromContents creates a new extractor from contents and page resources.
func NewFromContents(contents string, resources *_bg.PdfPageResources) (*Extractor, error) {
	const _eeg = "\u0065x\u0074\u0072\u0061\u0063t\u006f\u0072\u002e\u004e\u0065w\u0046r\u006fm\u0043\u006f\u006e\u0074\u0065\u006e\u0074s"
	_fbf := &Extractor{_eca: contents, _feg: resources, _deg: map[string]fontEntry{}, _gdfd: map[string]textResult{}}
	_fe.TrackUse(_eeg)
	return _fbf, nil
}
func (_cage *textObject) getCurrentFont() *_bg.PdfFont {
	_ceegg := _cage._egf._aede
	if _ceegg == nil {
		_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020U\u0073\u0069\u006e\u0067\u0020d\u0065\u0066a\u0075\u006c\u0074\u002e")
		return _bg.DefaultFont()
	}
	return _ceegg
}
func (_dfcgg *textPara) isAtom() *textTable {
	_fdggd := _dfcgg
	_cfee := _dfcgg._gbfc
	_eafa := _dfcgg._adef
	if _cfee.taken() || _eafa.taken() {
		return nil
	}
	_dbceg := _cfee._adef
	if _dbceg.taken() || _dbceg != _eafa._gbfc {
		return nil
	}
	return _dbfbe(_fdggd, _cfee, _eafa, _dbceg)
}
func _adccd(_cffa *wordBag, _eeda int) *textLine {
	_dgacf := _cffa.firstWord(_eeda)
	_ffbd := textLine{PdfRectangle: _dgacf.PdfRectangle, _gdbd: _dgacf._dafae, _ggfc: _dgacf._ccee}
	_ffbd.pullWord(_cffa, _dgacf, _eeda)
	return &_ffbd
}
func (_afadc *textTable) putComposite(_bfggb, _cdagb int, _egfb paraList, _ccgb _bg.PdfRectangle) {
	if len(_egfb) == 0 {
		_b.Log.Error("\u0074\u0065xt\u0054\u0061\u0062l\u0065\u0029\u0020\u0070utC\u006fmp\u006f\u0073\u0069\u0074\u0065\u003a\u0020em\u0070\u0074\u0079\u0020\u0070\u0061\u0072a\u0073")
		return
	}
	_gddfb := compositeCell{PdfRectangle: _ccgb, paraList: _egfb}
	if _eeeg {
		_d.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0070\u0075\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u003c\u002d\u0025\u0073\u000a", _bfggb, _cdagb, _gddfb.String())
	}
	_gddfb.updateBBox()
	_afadc._fccfa[_caabd(_bfggb, _cdagb)] = _gddfb
}
func _faef(_dcfe *Extractor, _ecd *_bg.PdfPageResources, _deeg _aed.GraphicsState, _cecac *textState, _fgg *stateStack) *textObject {
	return &textObject{_eadb: _dcfe, _fgf: _ecd, _dba: _deeg, _fbad: _fgg, _egf: _cecac, _gbcf: _ef.IdentityMatrix(), _gbe: _ef.IdentityMatrix()}
}
func (_agfdc gridTile) numBorders() int {
	_faga := 0
	if _agfdc._acef {
		_faga++
	}
	if _agfdc._bgdga {
		_faga++
	}
	if _agfdc._dbbba {
		_faga++
	}
	if _agfdc._acffg {
		_faga++
	}
	return _faga
}

// TableInfo gets table information of the textmark `tm`.
func (_gefb *TextMark) TableInfo() (*TextTable, [][]int) {
	if !_gefb._fafb {
		return nil, nil
	}
	_ffdf := _gefb._aeaf
	_eaag := _ffdf.getCellInfo(*_gefb)
	return _ffdf, _eaag
}

// BBox returns the smallest axis-aligned rectangle that encloses all the TextMarks in `ma`.
func (_cefbb *TextMarkArray) BBox() (_bg.PdfRectangle, bool) {
	var _edgg _bg.PdfRectangle
	_cfef := false
	for _, _bfgba := range _cefbb._egbe {
		if _bfgba.Meta || _degb(_bfgba.Text) {
			continue
		}
		if _cfef {
			_edgg = _egbb(_edgg, _bfgba.BBox)
		} else {
			_edgg = _bfgba.BBox
			_cfef = true
		}
	}
	return _edgg, _cfef
}
func (_degcg *textWord) toTextMarks(_dbcef *int) []TextMark {
	var _accb []TextMark
	for _, _bfcbb := range _degcg._ecdf {
		_accb = _cefe(_accb, _dbcef, _bfcbb.ToTextMark())
	}
	return _accb
}
func (_eagg *textObject) setTextMatrix(_bgg []float64) {
	if len(_bgg) != 6 {
		_b.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006c\u0065\u006e\u0028\u0066\u0029\u0020\u0021\u003d\u0020\u0036\u0020\u0028\u0025\u0064\u0029", len(_bgg))
		return
	}
	_cee, _ebcf, _cgfb, _gga, _accc, _gfbd := _bgg[0], _bgg[1], _bgg[2], _bgg[3], _bgg[4], _bgg[5]
	_eagg._gbcf = _ef.NewMatrix(_cee, _ebcf, _cgfb, _gga, _accc, _gfbd)
	_eagg._gbe = _eagg._gbcf
}
func (_fabaa *wordBag) sort() {
	for _, _ebdcb := range _fabaa._cccf {
		_gd.Slice(_ebdcb, func(_gcdg, _fecc int) bool { return _ccfa(_ebdcb[_gcdg], _ebdcb[_fecc]) < 0 })
	}
}
func _baagc(_agca, _ffcag *textPara) bool {
	if _agca._ecee || _ffcag._ecee {
		return true
	}
	return _gacaa(_agca.depth() - _ffcag.depth())
}
func (_bgbg *textObject) nextLine() { _bgbg.moveLP(0, -_bgbg._egf._fba) }
func _ddgea(_ebga int, _eggdd map[int][]float64) ([]int, int) {
	_eaee := make([]int, _ebga)
	_eadeb := 0
	for _faaa := 0; _faaa < _ebga; _faaa++ {
		_eaee[_faaa] = _eadeb
		_eadeb += len(_eggdd[_faaa]) + 1
	}
	return _eaee, _eadeb
}
func (_fega *textPara) getListLines() []*textLine {
	var _egdb []*textLine
	_dcad := _agfa(_fega._gdc)
	for _, _gbcge := range _fega._gdc {
		_dcac := _gbcge._edee[0]._edac[0]
		if _gbggf(_dcac) {
			_egdb = append(_egdb, _gbcge)
		}
	}
	_egdb = append(_egdb, _dcad...)
	return _egdb
}
func (_aceb *subpath) isQuadrilateral() bool {
	if len(_aceb._agea) < 4 || len(_aceb._agea) > 5 {
		return false
	}
	if len(_aceb._agea) == 5 {
		_fbdga := _aceb._agea[0]
		_egddg := _aceb._agea[4]
		if _fbdga.X != _egddg.X || _fbdga.Y != _egddg.Y {
			return false
		}
	}
	return true
}
func (_cggc *shapesState) drawRectangle(_adbb, _bdfa, _afbg, _efcb float64) {
	if _fbfd {
		_ebfd := _cggc.devicePoint(_adbb, _bdfa)
		_edge := _cggc.devicePoint(_adbb+_afbg, _bdfa+_efcb)
		_afd := _bg.PdfRectangle{Llx: _ebfd.X, Lly: _ebfd.Y, Urx: _edge.X, Ury: _edge.Y}
		_b.Log.Info("d\u0072a\u0077\u0052\u0065\u0063\u0074\u0061\u006e\u0067l\u0065\u003a\u0020\u00256.\u0032\u0066", _afd)
	}
	_cggc.newSubPath()
	_cggc.moveTo(_adbb, _bdfa)
	_cggc.lineTo(_adbb+_afbg, _bdfa)
	_cggc.lineTo(_adbb+_afbg, _bdfa+_efcb)
	_cggc.lineTo(_adbb, _bdfa+_efcb)
	_cggc.closePath()
}

type textLine struct {
	_bg.PdfRectangle
	_ggfc float64
	_edee []*textWord
	_gdbd float64
}

func (_gecgf *structTreeRoot) parseStructTreeRoot(_aagc _gbc.PdfObject) {
	if _aagc != nil {
		_cefc, _cdfd := _gbc.GetDict(_aagc)
		if !_cdfd {
			_b.Log.Debug("\u0070\u0061\u0072s\u0065\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u003a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
		}
		K := _cefc.Get("\u004b")
		_bcg := _cefc.Get("\u0054\u0079\u0070\u0065").String()
		var _ecb *_gbc.PdfObjectArray
		switch _gebd := K.(type) {
		case *_gbc.PdfObjectArray:
			_ecb = _gebd
		case *_gbc.PdfObjectReference:
			_ecb = _gbc.MakeArray(K)
		}
		_febfc := []structElement{}
		for _, _ecbb := range _ecb.Elements() {
			_aeac := &structElement{}
			_aeac.parseStructElement(_ecbb)
			_febfc = append(_febfc, *_aeac)
		}
		_gecgf._dgeg = _febfc
		_gecgf._dgba = _bcg
	}
}
func (_fegaf *structElement) parseStructElement(_gdag _gbc.PdfObject) {
	_gdaa, _dadb := _gbc.GetDict(_gdag)
	if !_dadb {
		_b.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0053\u0074\u0072u\u0063\u0074\u0045le\u006d\u0065\u006e\u0074\u003a\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u002e")
		return
	}
	_degd := _gdaa.Get("\u0053")
	_dbfd := _gdaa.Get("\u0050\u0067")
	_caea := ""
	if _degd != nil {
		_caea = _degd.String()
	}
	_eaaf := _gdaa.Get("\u004b")
	_fegaf._adfb = _caea
	_fegaf._dbdae = _dbfd
	switch _efgggc := _eaaf.(type) {
	case *_gbc.PdfObjectInteger:
		_fegaf._adfb = _caea
		_fegaf._agec = int64(*_efgggc)
		_fegaf._dbdae = _dbfd
	case *_gbc.PdfObjectReference:
		_cfdf := *_gbc.MakeArray(_efgggc)
		var _gccab int64 = -1
		_fegaf._agec = _gccab
		if _cfdf.Len() == 1 {
			_dcdg := _cfdf.Elements()[0]
			_eccdf, _bacd := _dcdg.(*_gbc.PdfObjectInteger)
			if _bacd {
				_gccab = int64(*_eccdf)
				_fegaf._agec = _gccab
				_fegaf._adfb = _caea
				_fegaf._dbdae = _dbfd
				return
			}
		}
		_fgag := []structElement{}
		for _, _gbee := range _cfdf.Elements() {
			_fgbb, _ffgg := _gbee.(*_gbc.PdfObjectInteger)
			if _ffgg {
				_gccab = int64(*_fgbb)
				_fegaf._agec = _gccab
				_fegaf._adfb = _caea
			} else {
				_adbbd := &structElement{}
				_adbbd.parseStructElement(_gbee)
				_fgag = append(_fgag, *_adbbd)
			}
			_gccab = -1
		}
		_fegaf._fbcba = _fgag
	case *_gbc.PdfObjectArray:
		_dfbaa := _eaaf.(*_gbc.PdfObjectArray)
		var _fgfb int64 = -1
		_fegaf._agec = _fgfb
		if _dfbaa.Len() == 1 {
			_fcea := _dfbaa.Elements()[0]
			_fedgc, _ded := _fcea.(*_gbc.PdfObjectInteger)
			if _ded {
				_fgfb = int64(*_fedgc)
				_fegaf._agec = _fgfb
				_fegaf._adfb = _caea
				_fegaf._dbdae = _dbfd
				return
			}
		}
		_ffcf := []structElement{}
		for _, _cedd := range _dfbaa.Elements() {
			_cedb, _ecefb := _cedd.(*_gbc.PdfObjectInteger)
			if _ecefb {
				_fgfb = int64(*_cedb)
				_fegaf._agec = _fgfb
				_fegaf._adfb = _caea
				_fegaf._dbdae = _dbfd
			} else {
				_ffca := &structElement{}
				_ffca.parseStructElement(_cedd)
				_ffcf = append(_ffcf, *_ffca)
			}
			_fgfb = -1
		}
		_fegaf._fbcba = _ffcf
	}
}
func (_bfgbf *wordBag) makeRemovals() map[int]map[*textWord]struct{} {
	_aedee := make(map[int]map[*textWord]struct{}, len(_bfgbf._cccf))
	for _fbea := range _bfgbf._cccf {
		_aedee[_fbea] = make(map[*textWord]struct{})
	}
	return _aedee
}
func (_dccf compositeCell) hasLines(_abbcg []*textLine) bool {
	for _befcd, _abdb := range _abbcg {
		_aecaa := _adcg(_dccf.PdfRectangle, _abdb.PdfRectangle)
		if _eeeg {
			_d.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u005e\u005e\u005e\u0069\u006e\u0074\u0065\u0072\u0073e\u0063t\u0073\u003d\u0025\u0074\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u000a", _aecaa, _befcd, len(_abbcg))
			_d.Printf("\u0020\u0020\u0020\u0020  \u005e\u005e\u005e\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025s\u000a", _dccf)
			_d.Printf("\u0020 \u0020 \u0020\u0020\u0020\u006c\u0069\u006e\u0065\u003d\u0025\u0073\u000a", _abdb)
		}
		if _aecaa {
			return true
		}
	}
	return false
}
func (_gbfe paraList) findTables(_gcgbd []gridTiling) []*textTable {
	_gbfe.addNeighbours()
	_gd.Slice(_gbfe, func(_fefeb, _edde int) bool { return _cada(_gbfe[_fefeb], _gbfe[_edde]) < 0 })
	var _gegb []*textTable
	if _dgaa {
		_cbbb := _gbfe.findGridTables(_gcgbd)
		_gegb = append(_gegb, _cbbb...)
	}
	if _bbgf {
		_febef := _gbfe.findTextTables()
		_gegb = append(_gegb, _febef...)
	}
	return _gegb
}

// String returns a description of `v`.
func (_aeeg *ruling) String() string {
	if _aeeg._gcafd == _bgeag {
		return "\u004e\u004f\u0054\u0020\u0052\u0055\u004c\u0049\u004e\u0047"
	}
	_cdeb, _decfc := "\u0078", "\u0079"
	if _aeeg._gcafd == _bfcbd {
		_cdeb, _decfc = "\u0079", "\u0078"
	}
	_fcgd := ""
	if _aeeg._gcec != 0.0 {
		_fcgd = _d.Sprintf(" \u0077\u0069\u0064\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _aeeg._gcec)
	}
	return _d.Sprintf("\u0025\u00310\u0073\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0020\u0025\u0076\u0025\u0073", _aeeg._gcafd, _cdeb, _aeeg._gbaff, _decfc, _aeeg._fabe, _aeeg._ccba, _aeeg._ccba-_aeeg._fabe, _aeeg._cebe, _aeeg.Color, _fcgd)
}

type markKind int

// String returns a human readable description of `path`.
func (_fabag *subpath) String() string {
	_ggcd := _fabag._agea
	_ccce := len(_ggcd)
	if _ccce <= 5 {
		return _d.Sprintf("\u0025d\u003a\u0020\u0025\u0036\u002e\u0032f", _ccce, _ggcd)
	}
	return _d.Sprintf("\u0025d\u003a\u0020\u0025\u0036.\u0032\u0066\u0020\u0025\u0036.\u0032f\u0020.\u002e\u002e\u0020\u0025\u0036\u002e\u0032f", _ccce, _ggcd[0], _ggcd[1], _ggcd[_ccce-1])
}

type stateStack []*textState

func (_caf *textObject) checkOp(_ebbb *_aed.ContentStreamOperation, _efd int, _ggb bool) (_cga bool, _dgac error) {
	if _caf == nil {
		var _bafd []_gbc.PdfObject
		if _efd > 0 {
			_bafd = _ebbb.Params
			if len(_bafd) > _efd {
				_bafd = _bafd[:_efd]
			}
		}
		_b.Log.Debug("\u0025\u0023q \u006f\u0070\u0065r\u0061\u006e\u0064\u0020out\u0073id\u0065\u0020\u0074\u0065\u0078\u0074\u002e p\u0061\u0072\u0061\u006d\u0073\u003d\u0025+\u0076", _ebbb.Operand, _bafd)
	}
	if _efd >= 0 {
		if len(_ebbb.Params) != _efd {
			if _ggb {
				_dgac = _dg.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
			}
			_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _ebbb.Operand, _efd, len(_ebbb.Params), _ebbb.Params)
			return false, _dgac
		}
	}
	return true, nil
}
func _gebg(_abgaa []*textMark, _eegadg _bg.PdfRectangle) []*textWord {
	var _fdcba []*textWord
	var _febec *textWord
	if _fcee {
		_b.Log.Info("\u006d\u0061\u006beT\u0065\u0078\u0074\u0057\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0064\u0020\u006d\u0061\u0072\u006b\u0073", len(_abgaa))
	}
	_afgdc := func() {
		if _febec != nil {
			_gcgba := _febec.computeText()
			if !_degb(_gcgba) {
				_febec._edac = _gcgba
				_fdcba = append(_fdcba, _febec)
				if _fcee {
					_b.Log.Info("\u0061\u0064\u0064Ne\u0077\u0057\u006f\u0072\u0064\u003a\u0020\u0025\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", len(_fdcba)-1, _febec.String())
					for _cdef, _bgaa := range _febec._ecdf {
						_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cdef, _bgaa.String())
					}
				}
			}
			_febec = nil
		}
	}
	for _, _dffca := range _abgaa {
		if _afgc && _febec != nil && len(_febec._ecdf) > 0 {
			_fbgag := _febec._ecdf[len(_febec._ecdf)-1]
			_fafde, _gabf := _bdge(_dffca._aaag)
			_gdbgb, _cebbb := _bdge(_fbgag._aaag)
			if _gabf && !_cebbb && _fbgag.inDiacriticArea(_dffca) {
				_febec.addDiacritic(_fafde)
				continue
			}
			if _cebbb && !_gabf && _dffca.inDiacriticArea(_fbgag) {
				_febec._ecdf = _febec._ecdf[:len(_febec._ecdf)-1]
				_febec.appendMark(_dffca, _eegadg)
				_febec.addDiacritic(_gdbgb)
				continue
			}
		}
		_acbb := _degb(_dffca._aaag)
		if _acbb {
			_afgdc()
			continue
		}
		if _febec == nil && !_acbb {
			_febec = _dcacd([]*textMark{_dffca}, _eegadg)
			continue
		}
		_gabgd := _febec._dafae
		_cbccg := _bb.Abs(_gafde(_eegadg, _dffca)-_febec._ccee) / _gabgd
		_gfbed := _gefae(_dffca, _febec) / _gabgd
		if _gfbed >= _gebc || !(-_fbadd <= _gfbed && _cbccg <= _gadf) {
			_afgdc()
			_febec = _dcacd([]*textMark{_dffca}, _eegadg)
			continue
		}
		_febec.appendMark(_dffca, _eegadg)
	}
	_afgdc()
	return _fdcba
}
func _cgef(_aaeg []*textLine, _dadc string) string {
	var _fffgc _ge.Builder
	_bafdb := 0.0
	for _adfd, _deaa := range _aaeg {
		_bbgc := _deaa.text()
		_ecbf := _deaa._ggfc
		if _adfd < len(_aaeg)-1 {
			_bafdb = _aaeg[_adfd+1]._ggfc
		} else {
			_bafdb = 0.0
		}
		_fffgc.WriteString(_dadc)
		_fffgc.WriteString(_bbgc)
		if _bafdb != _ecbf {
			_fffgc.WriteString("\u000a")
		} else {
			_fffgc.WriteString("\u0020")
		}
	}
	return _fffgc.String()
}

var (
	_gfg = _dg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_adc = _dg.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func (_eaab *imageExtractContext) extractFormImages(_afc *_gbc.PdfObjectName, _gfcc _aed.GraphicsState, _edd *_bg.PdfPageResources) error {
	_dfg, _dfag := _edd.GetXObjectFormByName(*_afc)
	if _dfag != nil {
		return _dfag
	}
	if _dfg == nil {
		return nil
	}
	_aab, _dfag := _dfg.GetContentStream()
	if _dfag != nil {
		return _dfag
	}
	_cdgg := _dfg.Resources
	if _cdgg == nil {
		_cdgg = _edd
	}
	_dfag = _eaab.extractContentStreamImages(string(_aab), _cdgg)
	if _dfag != nil {
		return _dfag
	}
	_eaab._da++
	return nil
}
func (_cbcgdc gridTiling) complete() bool {
	for _, _fbcec := range _cbcgdc._egff {
		for _, _acaca := range _fbcec {
			if !_acaca.complete() {
				return false
			}
		}
	}
	return true
}
func _gag(_abeg *wordBag, _gaacd *textWord, _abab float64) bool {
	return _abeg.Urx <= _gaacd.Llx && _gaacd.Llx < _abeg.Urx+_abab
}
func (_efe *shapesState) newSubPath() {
	_efe.clearPath()
	if _fbfd {
		_b.Log.Info("\u006e\u0065\u0077\u0053\u0075\u0062\u0050\u0061\u0074h\u003a\u0020\u0025\u0073", _efe)
	}
}
func (_ebfa *subpath) last() _ef.Point { return _ebfa._agea[len(_ebfa._agea)-1] }
func (_ebadg *textTable) getComposite(_bebge, _daga int) (paraList, _bg.PdfRectangle) {
	_cfceg, _badfd := _ebadg._fccfa[_caabd(_bebge, _daga)]
	if _eeeg {
		_d.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0067\u0065\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u002d\u003e\u0025\u0073\u000a", _bebge, _daga, _cfceg.String())
	}
	if !_badfd {
		return nil, _bg.PdfRectangle{}
	}
	return _cfceg.parasBBox()
}
func (_dfdad *textTable) compositeRowCorridors() map[int][]float64 {
	_afbgd := make(map[int][]float64, _dfdad._egbbf)
	if _eeeg {
		_b.Log.Info("c\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0052\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073:\u0020h\u003d\u0025\u0064", _dfdad._egbbf)
	}
	for _ebcgf := 1; _ebcgf < _dfdad._egbbf; _ebcgf++ {
		var _bcbaf []compositeCell
		for _dbdac := 0; _dbdac < _dfdad._ddcega; _dbdac++ {
			if _ddfcc, _afbfe := _dfdad._fccfa[_caabd(_dbdac, _ebcgf)]; _afbfe {
				_bcbaf = append(_bcbaf, _ddfcc)
			}
		}
		if len(_bcbaf) == 0 {
			continue
		}
		_cbdf := _gdgea(_bcbaf)
		_afbgd[_ebcgf] = _cbdf
		if _eeeg {
			_d.Printf("\u0020\u0020\u0020\u0025\u0032\u0064\u003a\u0020\u00256\u002e\u0032\u0066\u000a", _ebcgf, _cbdf)
		}
	}
	return _afbgd
}
func _fcgb(_fdfa func(*wordBag, *textWord, float64) bool, _dafd float64) func(*wordBag, *textWord) bool {
	return func(_cfcf *wordBag, _ageaa *textWord) bool { return _fdfa(_cfcf, _ageaa, _dafd) }
}
func (_ffd *textObject) showTextAdjusted(_bcf *_gbc.PdfObjectArray, _gaf int, _ceab string) error {
	_dfgbg := false
	for _, _deab := range _bcf.Elements() {
		switch _deab.(type) {
		case *_gbc.PdfObjectFloat, *_gbc.PdfObjectInteger:
			_agbd, _ccde := _gbc.GetNumberAsFloat(_deab)
			if _ccde != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078t\u0041\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0042\u0061\u0064\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0061\u006c\u0020a\u0072\u0067\u002e\u0020\u006f\u003d\u0025\u0073\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _deab, _bcf)
				return _ccde
			}
			_dgdc, _adcc := -_agbd*0.001*_ffd._egf._ebab, 0.0
			if _dfgbg {
				_adcc, _dgdc = _dgdc, _adcc
			}
			_dcfa := _bcca(_ef.Point{X: _dgdc, Y: _adcc})
			_ffd._gbcf.Concat(_dcfa)
		case *_gbc.PdfObjectString:
			_ced := _gbc.TraceToDirectObject(_deab)
			_aeca, _ddgg := _gbc.GetStringBytes(_ced)
			if !_ddgg {
				_b.Log.Trace("s\u0068\u006f\u0077\u0054\u0065\u0078\u0074\u0041\u0064j\u0075\u0073\u0074\u0065\u0064\u003a\u0020Ba\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u002e\u0020o=\u0025\u0073 \u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _deab, _bcf)
				return _gbc.ErrTypeError
			}
			_ffd.renderText(_ced, _aeca, _gaf, _ceab)
		default:
			_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078\u0074A\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0055\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0028%T\u0029\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _deab, _bcf)
			return _gbc.ErrTypeError
		}
	}
	return nil
}
func (_babad rulingList) findPrimSec(_fdbg, _cgfab float64) *ruling {
	for _, _babd := range _babad {
		if _gacaa(_babd._gbaff-_fdbg) && _babd._fabe-_beac <= _cgfab && _cgfab <= _babd._ccba+_beac {
			return _babd
		}
	}
	return nil
}
func (_fccg gridTile) contains(_bebc _bg.PdfRectangle) bool {
	if _fccg.numBorders() < 3 {
		return false
	}
	if _fccg._acef && _bebc.Llx < _fccg.Llx-_edef {
		return false
	}
	if _fccg._bgdga && _bebc.Urx > _fccg.Urx+_edef {
		return false
	}
	if _fccg._dbbba && _bebc.Lly < _fccg.Lly-_edef {
		return false
	}
	if _fccg._acffg && _bebc.Ury > _fccg.Ury+_edef {
		return false
	}
	return true
}

// Marks returns the TextMark collection for a page. It represents all the text on the page.
func (_dbdgb PageText) Marks() *TextMarkArray { return &TextMarkArray{_egbe: _dbdgb._dcge} }
func (_dfbd *textObject) moveLP(_bafc, _ceeg float64) {
	_dfbd._gbe.Concat(_ef.NewMatrix(1, 0, 0, 1, _bafc, _ceeg))
	_dfbd._gbcf = _dfbd._gbe
}
func _dgcb(_eaffc, _fgab _ef.Point) bool {
	_egdc := _bb.Abs(_eaffc.X - _fgab.X)
	_egbdb := _bb.Abs(_eaffc.Y - _fgab.Y)
	return _fcge(_egdc, _egbdb)
}

// String returns a string describing the current state of the textState stack.
func (_dag *stateStack) String() string {
	_ddfb := []string{_d.Sprintf("\u002d\u002d\u002d\u002d f\u006f\u006e\u0074\u0020\u0073\u0074\u0061\u0063\u006b\u003a\u0020\u0025\u0064", len(*_dag))}
	for _deef, _edae := range *_dag {
		_eaad := "\u003c\u006e\u0069l\u003e"
		if _edae != nil {
			_eaad = _edae.String()
		}
		_ddfb = append(_ddfb, _d.Sprintf("\u0009\u0025\u0032\u0064\u003a\u0020\u0025\u0073", _deef, _eaad))
	}
	return _ge.Join(_ddfb, "\u000a")
}
func (_aadc paraList) reorder(_baafba []int) {
	_fabf := make(paraList, len(_aadc))
	for _eebd, _fbdg := range _baafba {
		_fabf[_eebd] = _aadc[_fbdg]
	}
	copy(_aadc, _fabf)
}
func _ddceg(_ddbd map[float64][]*textLine) []float64 {
	_ffafg := []float64{}
	for _dgcf := range _ddbd {
		_ffafg = append(_ffafg, _dgcf)
	}
	_gd.Float64s(_ffafg)
	return _ffafg
}
func (_cbcc *wordBag) maxDepth() float64               { return _cbcc._cfb - _cbcc.Lly }
func (_febf *textObject) moveText(_bfb, _gfbb float64) { _febf.moveLP(_bfb, _gfbb) }
func (_gcc *TextMarkArray) getTextMarkAtOffset(_bbg int) *TextMark {
	for _, _fffb := range _gcc._egbe {
		if _fffb.Offset == _bbg {
			return &_fffb
		}
	}
	return nil
}
func (_bedg gridTile) complete() bool { return _bedg.numBorders() == 4 }
func (_ddc *textObject) moveTextSetLeading(_adf, _geg float64) {
	_ddc._egf._fba = -_geg
	_ddc.moveLP(_adf, _geg)
}

const (
	_baf  = "\u0045\u0052R\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_febe = "\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043a\u006e\u0027\u0074 g\u0065\u0074\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002c\u0020\u0066\u006fn\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064"
	_dbdg = "\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_edfe = "E\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066o\u006e\u0074\u0020\u0066ou\u006e\u0064"
)

func _ccfa(_cfgc, _feaf bounded) float64 { return _cfgc.bbox().Llx - _feaf.bbox().Llx }
func (_egbgd rulingList) sort()          { _gd.Slice(_egbgd, _egbgd.comp) }
func _agfa(_fccfc []*textLine) []*textLine {
	_fbcbc := []*textLine{}
	for _, _cfaa := range _fccfc {
		_fbbf := _cfaa.text()
		_gfabe := _cfabe.Find([]byte(_fbbf))
		if _gfabe != nil {
			_fbcbc = append(_fbcbc, _cfaa)
		}
	}
	return _fbcbc
}
func _cefe(_fcdce []TextMark, _dggf *int, _fbcc TextMark) []TextMark {
	_fbcc.Offset = *_dggf
	_fcdce = append(_fcdce, _fbcc)
	*_dggf += len(_fbcc.Text)
	return _fcdce
}
func (_bdage *textTable) emptyCompositeColumn(_fdaee int) bool {
	for _gcdee := 0; _gcdee < _bdage._egbbf; _gcdee++ {
		if _ceage, _efed := _bdage._fccfa[_caabd(_fdaee, _gcdee)]; _efed {
			if len(_ceage.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_geed *wordBag) blocked(_beae *textWord) bool {
	if _beae.Urx < _geed.Llx {
		_adaa := _bbgg(_beae.PdfRectangle)
		_febfg := _ccbc(_geed.PdfRectangle)
		if _geed._ecdb.blocks(_adaa, _febfg) {
			if _bbae {
				_b.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0078\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _beae, _geed)
			}
			return true
		}
	} else if _geed.Urx < _beae.Llx {
		_ccdc := _bbgg(_geed.PdfRectangle)
		_cabe := _ccbc(_beae.PdfRectangle)
		if _geed._ecdb.blocks(_ccdc, _cabe) {
			if _bbae {
				_b.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0078\u2192\u0020:\u0020\u0025\u0073\u0020%s", _beae, _geed)
			}
			return true
		}
	}
	if _beae.Ury < _geed.Lly {
		_dgae := _dcgec(_beae.PdfRectangle)
		_ebfg := _dgge(_geed.PdfRectangle)
		if _geed._ebfag.blocks(_dgae, _ebfg) {
			if _bbae {
				_b.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0079\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _beae, _geed)
			}
			return true
		}
	} else if _geed.Ury < _beae.Lly {
		_aedgf := _dcgec(_geed.PdfRectangle)
		_adbca := _dgge(_beae.PdfRectangle)
		if _geed._ebfag.blocks(_aedgf, _adbca) {
			if _bbae {
				_b.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0079\u2192\u0020:\u0020\u0025\u0073\u0020%s", _beae, _geed)
			}
			return true
		}
	}
	return false
}

// GetContentStreamOps returns the contentStreamOps field of `pt`.
func (_eddb *PageText) GetContentStreamOps() *_aed.ContentStreamOperations { return _eddb._dbcd }
func (_bgagf *wordBag) removeWord(_gcae *textWord, _gbcfb int) {
	_eefcg := _bgagf._cccf[_gbcfb]
	_eefcg = _eaggd(_eefcg, _gcae)
	if len(_eefcg) == 0 {
		delete(_bgagf._cccf, _gbcfb)
	} else {
		_bgagf._cccf[_gbcfb] = _eefcg
	}
}
func (_fcdc *wordBag) firstReadingIndex(_cgge int) int {
	_eebf := _fcdc.firstWord(_cgge)._dafae
	_ffcbb := float64(_cgge+1) * _dfbad
	_adae := _ffcbb + _debb*_eebf
	_aaeb := _cgge
	for _, _cgea := range _fcdc.depthBand(_ffcbb, _adae) {
		if _ccfa(_fcdc.firstWord(_cgea), _fcdc.firstWord(_aaeb)) < 0 {
			_aaeb = _cgea
		}
	}
	return _aaeb
}
func _dcde(_gfgf *textLine) float64 { return _gfgf._edee[0].Llx }
func _bfge(_efgfe, _cbcgd _ef.Point, _fbbac _ab.Color) (*ruling, bool) {
	_ageec := lineRuling{_abfed: _efgfe, _fgcb: _cbcgd, _gadeg: _bafag(_efgfe, _cbcgd), Color: _fbbac}
	if _ageec._gadeg == _bgeag {
		return nil, false
	}
	return _ageec.asRuling()
}
func (_bgffb *wordBag) removeDuplicates() {
	if _bfab {
		_b.Log.Info("r\u0065m\u006f\u0076\u0065\u0044\u0075\u0070\u006c\u0069c\u0061\u0074\u0065\u0073: \u0025\u0071", _bgffb.text())
	}
	for _, _cbcg := range _bgffb.depthIndexes() {
		if len(_bgffb._cccf[_cbcg]) == 0 {
			continue
		}
		_eaff := _bgffb._cccf[_cbcg][0]
		_gecc := _ebee * _eaff._dafae
		_cddf := _eaff._ccee
		for _, _ffdd := range _bgffb.depthBand(_cddf, _cddf+_gecc) {
			_effc := map[*textWord]struct{}{}
			_ecgg := _bgffb._cccf[_ffdd]
			for _, _fccfb := range _ecgg {
				if _, _cgfed := _effc[_fccfb]; _cgfed {
					continue
				}
				for _, _gbafb := range _ecgg {
					if _, _adadc := _effc[_gbafb]; _adadc {
						continue
					}
					if _gbafb != _fccfb && _gbafb._edac == _fccfb._edac && _bb.Abs(_gbafb.Llx-_fccfb.Llx) < _gecc && _bb.Abs(_gbafb.Urx-_fccfb.Urx) < _gecc && _bb.Abs(_gbafb.Lly-_fccfb.Lly) < _gecc && _bb.Abs(_gbafb.Ury-_fccfb.Ury) < _gecc {
						_effc[_gbafb] = struct{}{}
					}
				}
			}
			if len(_effc) > 0 {
				_efcc := 0
				for _, _bcdg := range _ecgg {
					if _, _gdfc := _effc[_bcdg]; !_gdfc {
						_ecgg[_efcc] = _bcdg
						_efcc++
					}
				}
				_bgffb._cccf[_ffdd] = _ecgg[:len(_ecgg)-len(_effc)]
				if len(_bgffb._cccf[_ffdd]) == 0 {
					delete(_bgffb._cccf, _ffdd)
				}
			}
		}
	}
}

var _fdad = TextMark{Text: "\u005b\u0058\u005d", Original: "\u0020", Meta: true, FillColor: _ab.White, StrokeColor: _ab.White}

func (_aeadb rulingList) aligned() bool {
	if len(_aeadb) < 2 {
		return false
	}
	_cgfa := make(map[*ruling]int)
	_cgfa[_aeadb[0]] = 0
	for _, _bfce := range _aeadb[1:] {
		_gedb := false
		for _dbce := range _cgfa {
			if _bfce.gridIntersecting(_dbce) {
				_cgfa[_dbce]++
				_gedb = true
				break
			}
		}
		if !_gedb {
			_cgfa[_bfce] = 0
		}
	}
	_egee := 0
	for _, _begd := range _cgfa {
		if _begd == 0 {
			_egee++
		}
	}
	_gdda := float64(_egee) / float64(len(_aeadb))
	_fbeb := _gdda <= 1.0-_ddee
	if _gfab {
		_b.Log.Info("\u0061\u006c\u0069\u0067\u006e\u0065\u0064\u003d\u0025\u0074\u0020\u0075\u006em\u0061\u0074\u0063\u0068\u0065\u0064=\u0025\u002e\u0032\u0066\u003d\u0025\u0064\u002f\u0025\u0064\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _fbeb, _gdda, _egee, len(_aeadb), _aeadb.String())
	}
	return _fbeb
}
func (_fgbe lineRuling) yMean() float64 { return 0.5 * (_fgbe._abfed.Y + _fgbe._fgcb.Y) }
func _egdbf(_gfcd *textLine) bool {
	_cfgcg := true
	_bedf := -1
	for _, _afdf := range _gfcd._edee {
		for _, _dfed := range _afdf._ecdf {
			_bfde := _dfed._bcab
			if _bedf == -1 {
				_bedf = _bfde
			} else {
				if _bedf != _bfde {
					_cfgcg = false
					break
				}
			}
		}
	}
	return _cfgcg
}
func _fgc(_gdfgb, _gggg bounded) float64 { return _bggc(_gdfgb) - _bggc(_gggg) }

// BidiText represents a bidi text organized in its visual order
// with base direction of the text.
type BidiText struct {
	_feb string
	_cae string
}

func _cgb(_eefc *_aed.ContentStreamOperation) (float64, error) {
	if len(_eefc.Params) != 1 {
		_adfac := _dg.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
		_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _eefc.Operand, 1, len(_eefc.Params), _eefc.Params)
		return 0.0, _adfac
	}
	return _gbc.GetNumberAsFloat(_eefc.Params[0])
}
func (_ddga *textTable) getDown() paraList {
	_fafe := make(paraList, _ddga._ddcega)
	for _bfbbb := 0; _bfbbb < _ddga._ddcega; _bfbbb++ {
		_cgbg := _ddga.get(_bfbbb, _ddga._egbbf-1)._adef
		if _cgbg.taken() {
			return nil
		}
		_fafe[_bfbbb] = _cgbg
	}
	for _gaaga := 0; _gaaga < _ddga._ddcega-1; _gaaga++ {
		if _fafe[_gaaga]._gbfc != _fafe[_gaaga+1] {
			return nil
		}
	}
	return _fafe
}

// String returns a description of `k`.
func (_gdcb markKind) String() string {
	_ccfba, _fbgg := _dgeda[_gdcb]
	if !_fbgg {
		return _d.Sprintf("\u004e\u006f\u0074\u0020\u0061\u0020\u006d\u0061\u0072k\u003a\u0020\u0025\u0064", _gdcb)
	}
	return _ccfba
}
func (_bdff *shapesState) stroke(_cca *[]pathSection) {
	_ace := pathSection{_cdc: _bdff._dgdd, Color: _bdff._eadd.getStrokeColor()}
	*_cca = append(*_cca, _ace)
	if _gfab {
		_d.Printf("\u0020 \u0020\u0020S\u0054\u0052\u004fK\u0045\u003a\u0020\u0025\u0064\u0020\u0073t\u0072\u006f\u006b\u0065\u0073\u0020s\u0073\u003d\u0025\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d%\u002b\u0076\u0020\u0025\u0036\u002e\u0032\u0066\u000a", len(*_cca), _bdff, _bdff._eadd.getStrokeColor(), _ace.bbox())
		if _fafd {
			for _egbf, _bge := range _bdff._dgdd {
				_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _egbf, _bge)
				if _egbf == 10 {
					break
				}
			}
		}
	}
}
func (_ffaf *textLine) appendWord(_ababe *textWord) {
	_ffaf._edee = append(_ffaf._edee, _ababe)
	_ffaf.PdfRectangle = _egbb(_ffaf.PdfRectangle, _ababe.PdfRectangle)
	if _ababe._dafae > _ffaf._gdbd {
		_ffaf._gdbd = _ababe._dafae
	}
	if _ababe._ccee > _ffaf._ggfc {
		_ffaf._ggfc = _ababe._ccee
	}
}
func (_fbeaa *wordBag) minDepth() float64 { return _fbeaa._cfb - (_fbeaa.Ury - _fbeaa._aced) }
func _cffg(_eff float64, _bfbb int) int {
	if _bfbb == 0 {
		_bfbb = 1
	}
	_daffb := float64(_bfbb)
	return int(_bb.Round(_eff/_daffb) * _daffb)
}

type textTable struct {
	_bg.PdfRectangle
	_ddcega, _egbbf int
	_gfagc          bool
	_bcec           map[uint64]*textPara
	_fccfa          map[uint64]compositeCell
}

func (_abce rulingList) primMinMax() (float64, float64) {
	_gaag, _daad := _abce[0]._gbaff, _abce[0]._gbaff
	for _, _gcea := range _abce[1:] {
		if _gcea._gbaff < _gaag {
			_gaag = _gcea._gbaff
		} else if _gcea._gbaff > _daad {
			_daad = _gcea._gbaff
		}
	}
	return _gaag, _daad
}

const _fbfa = 10

func _ebeeg(_bfae _bg.PdfRectangle, _geea []*textLine) *textPara {
	return &textPara{PdfRectangle: _bfae, _gdc: _geea}
}
func _bggc(_gbd bounded) float64 { return -_gbd.bbox().Lly }

type fontEntry struct {
	_cbaf *_bg.PdfFont
	_afgb int64
}
type gridTiling struct {
	_bg.PdfRectangle
	_cdfg []float64
	_fgcd []float64
	_egff map[float64]map[float64]gridTile
}

func _efee(_fafbg *list, _fbbb *string) string {
	_gaee := _ge.Split(_fafbg._fccf, "\u000a")
	_dcaf := &_ge.Builder{}
	for _, _cdcce := range _gaee {
		if _cdcce != "" {
			_dcaf.WriteString(*_fbbb)
			_dcaf.WriteString(_cdcce)
			_dcaf.WriteString("\u000a")
		}
	}
	return _dcaf.String()
}

type rulingKind int

// ImageMark represents an image drawn on a page and its position in device coordinates.
// All coordinates are in device coordinates.
type ImageMark struct {
	Image *_bg.Image

	// Dimensions of the image as displayed in the PDF.
	Width  float64
	Height float64

	// Position of the image in PDF coordinates (lower left corner).
	X float64
	Y float64

	// Angle in degrees, if rotated.
	Angle float64
}

func _fcge(_dgca, _fcfc float64) bool { return _dgca/_bb.Max(_cbab, _fcfc) < _baaee }
func (_ccgfb *textTable) growTable() {
	_abba := func(_aage paraList) {
		_ccgfb._egbbf++
		for _eabcc := 0; _eabcc < _ccgfb._ddcega; _eabcc++ {
			_cfccg := _aage[_eabcc]
			_ccgfb.put(_eabcc, _ccgfb._egbbf-1, _cfccg)
		}
	}
	_fcadc := func(_bbeac paraList) {
		_ccgfb._ddcega++
		for _degc := 0; _degc < _ccgfb._egbbf; _degc++ {
			_fdge := _bbeac[_degc]
			_ccgfb.put(_ccgfb._ddcega-1, _degc, _fdge)
		}
	}
	if _dcfc {
		_ccgfb.log("\u0067r\u006f\u0077\u0054\u0061\u0062\u006ce")
	}
	for _edgd := 0; ; _edgd++ {
		_ecedg := false
		_becga := _ccgfb.getDown()
		_ggga := _ccgfb.getRight()
		if _dcfc {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _edgd, _ccgfb)
			_d.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0020\u0064\u006f\u0077\u006e\u003d\u0025\u0073\u000a", _becga)
			_d.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0072\u0069\u0067\u0068\u0074\u003d\u0025\u0073\u000a", _ggga)
		}
		if _becga != nil && _ggga != nil {
			_agae := _becga[len(_becga)-1]
			if !_agae.taken() && _agae == _ggga[len(_ggga)-1] {
				_abba(_becga)
				if _ggga = _ccgfb.getRight(); _ggga != nil {
					_fcadc(_ggga)
					_ccgfb.put(_ccgfb._ddcega-1, _ccgfb._egbbf-1, _agae)
				}
				_ecedg = true
			}
		}
		if !_ecedg && _becga != nil {
			_abba(_becga)
			_ecedg = true
		}
		if !_ecedg && _ggga != nil {
			_fcadc(_ggga)
			_ecedg = true
		}
		if !_ecedg {
			break
		}
	}
}
func (_baddd rulingList) tidied(_egbge string) rulingList {
	_ceege := _baddd.removeDuplicates()
	_ceege.log("\u0075n\u0069\u0071\u0075\u0065\u0073")
	_eddf := _ceege.snapToGroups()
	if _eddf == nil {
		return nil
	}
	_eddf.sort()
	if _gfab {
		_b.Log.Info("\u0074\u0069\u0064i\u0065\u0064\u003a\u0020\u0025\u0071\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u0075\u006e\u0069\u0071\u0075\u0065\u0073\u003d\u0025\u0064\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0064\u003d\u0025\u0064", _egbge, len(_baddd), len(_ceege), len(_eddf))
	}
	_eddf.log("\u0063o\u0061\u006c\u0065\u0073\u0063\u0065d")
	return _eddf
}
func (_faeb lineRuling) xMean() float64 { return 0.5 * (_faeb._abfed.X + _faeb._fgcb.X) }

// ToTextMark returns the public view of `tm`.
func (_dbdb *textMark) ToTextMark() TextMark {
	return TextMark{Text: _dbdb._aaag, Original: _dbdb._aca, BBox: _dbdb._beceb, Font: _dbdb._bebg, FontSize: _dbdb._cfgcd, FillColor: _dbdb._ddbdc, StrokeColor: _dbdb._dgad, Orientation: _dbdb._ecefa, DirectObject: _dbdb._afabc, ObjString: _dbdb._gabe, Tw: _dbdb.Tw, Th: _dbdb.Th, Tc: _dbdb._deed, Index: _dbdb._fcbe}
}
func (_bfgb *textObject) showText(_aecf _gbc.PdfObject, _gege []byte, _edc int, _cdf string) error {
	return _bfgb.renderText(_aecf, _gege, _edc, _cdf)
}
func (_ccca *wordBag) absorb(_eee *wordBag) {
	_gfbe := _eee.makeRemovals()
	for _efggg, _efabc := range _eee._cccf {
		for _, _ceda := range _efabc {
			_ccca.pullWord(_ceda, _efggg, _gfbe)
		}
	}
	_eee.applyRemovals(_gfbe)
}

// ToText returns the page text as a single string.
// Deprecated: This function is deprecated and will be removed in a future major version. Please use
// Text() instead.
func (_bce PageText) ToText() string { return _bce.Text() }
func (_dgfd *textObject) newTextMark(_cacb string, _dfda _ef.Matrix, _ccgeg _ef.Point, _gcg float64, _bgdg *_bg.PdfFont, _gaff float64, _bbce, _fbeab _ab.Color, _faag _gbc.PdfObject, _daeg []string, _eccdc int, _dbafe int) (textMark, bool) {
	_fcacc := _dfda.Angle()
	_dcbc := _cffg(_fcacc, _efdfe)
	var _egbd float64
	if _dcbc%180 != 90 {
		_egbd = _dfda.ScalingFactorY()
	} else {
		_egbd = _dfda.ScalingFactorX()
	}
	_adcf := _cfg(_dfda)
	_baga := _bg.PdfRectangle{Llx: _adcf.X, Lly: _adcf.Y, Urx: _ccgeg.X, Ury: _ccgeg.Y}
	switch _dcbc % 360 {
	case 90:
		_baga.Urx -= _egbd
	case 180:
		_baga.Ury -= _egbd
	case 270:
		_baga.Urx += _egbd
	case 0:
		_baga.Ury += _egbd
	default:
		_dcbc = 0
		_baga.Ury += _egbd
	}
	if _baga.Llx > _baga.Urx {
		_baga.Llx, _baga.Urx = _baga.Urx, _baga.Llx
	}
	if _baga.Lly > _baga.Ury {
		_baga.Lly, _baga.Ury = _baga.Ury, _baga.Lly
	}
	_aggd := true
	if _dgfd._eadb._ac.Width() > 0 {
		_cbcb, _fceae := _cbee(_baga, _dgfd._eadb._ac)
		if !_fceae {
			_aggd = false
			_b.Log.Debug("\u0054\u0065\u0078\u0074\u0020m\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0070a\u0067\u0065\u002e\u0020\u0062\u0062\u006f\u0078\u003d\u0025\u0067\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u0067\u0020\u0074\u0065\u0078\u0074\u003d\u0025q", _baga, _dgfd._eadb._ac, _cacb)
		}
		_baga = _cbcb
	}
	_gfgfb := _baga
	_fdfb := _dgfd._eadb._ac
	switch _dcbc % 360 {
	case 90:
		_fdfb.Urx, _fdfb.Ury = _fdfb.Ury, _fdfb.Urx
		_gfgfb = _bg.PdfRectangle{Llx: _fdfb.Urx - _baga.Ury, Urx: _fdfb.Urx - _baga.Lly, Lly: _baga.Llx, Ury: _baga.Urx}
	case 180:
		_gfgfb = _bg.PdfRectangle{Llx: _fdfb.Urx - _baga.Llx, Urx: _fdfb.Urx - _baga.Urx, Lly: _fdfb.Ury - _baga.Lly, Ury: _fdfb.Ury - _baga.Ury}
	case 270:
		_fdfb.Urx, _fdfb.Ury = _fdfb.Ury, _fdfb.Urx
		_gfgfb = _bg.PdfRectangle{Llx: _baga.Ury, Urx: _baga.Lly, Lly: _fdfb.Ury - _baga.Llx, Ury: _fdfb.Ury - _baga.Urx}
	}
	if _gfgfb.Llx > _gfgfb.Urx {
		_gfgfb.Llx, _gfgfb.Urx = _gfgfb.Urx, _gfgfb.Llx
	}
	if _gfgfb.Lly > _gfgfb.Ury {
		_gfgfb.Lly, _gfgfb.Ury = _gfgfb.Ury, _gfgfb.Lly
	}
	_gdagb := textMark{_aaag: _cacb, PdfRectangle: _gfgfb, _beceb: _baga, _bebg: _bgdg, _cfgcd: _egbd, _deed: _gaff, _aefd: _dfda, _dadg: _ccgeg, _ecefa: _dcbc, _ddbdc: _bbce, _dgad: _fbeab, _afabc: _faag, _gabe: _daeg, Th: _dgfd._egf._ffdc, Tw: _dgfd._egf._eaed, _bcab: _dbafe, _fcbe: _eccdc}
	if _fcee {
		_b.Log.Info("n\u0065\u0077\u0054\u0065\u0078\u0074M\u0061\u0072\u006b\u003a\u0020\u0073t\u0061\u0072\u0074\u003d\u0025\u002e\u0032f\u0020\u0065\u006e\u0064\u003d\u0025\u002e\u0032\u0066\u0020%\u0073", _adcf, _ccgeg, _gdagb.String())
	}
	return _gdagb, _aggd
}
func _fccc(_eeeea *PageText) error {
	_deefa := _fe.GetLicenseKey()
	if _deefa != nil && _deefa.IsLicensed() || _abe {
		return nil
	}
	_d.Printf("\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064\u0020c\u006f\u0070\u0079\u0020\u006f\u0066\u0020\u0055\u006e\u0069P\u0044\u0046\u000a")
	_d.Println("-\u0020\u0047\u0065\u0074\u0020\u0061\u0020\u0066\u0072e\u0065\u0020\u0074\u0072\u0069\u0061\u006c l\u0069\u0063\u0065\u006es\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070s:\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	return _dg.New("\u0075\u006e\u0069\u0070d\u0066\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020c\u006fd\u0065\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0064")
}

// String returns a description of `t`.
func (_cgcf *textTable) String() string {
	return _d.Sprintf("\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0025\u0074", _cgcf._ddcega, _cgcf._egbbf, _cgcf._gfagc)
}
func (_dacc rulingList) splitSec() []rulingList {
	_gd.Slice(_dacc, func(_debd, _aecc int) bool {
		_bdcf, _aggfb := _dacc[_debd], _dacc[_aecc]
		if _bdcf._fabe != _aggfb._fabe {
			return _bdcf._fabe < _aggfb._fabe
		}
		return _bdcf._ccba < _aggfb._ccba
	})
	_dgeba := make(map[*ruling]struct{}, len(_dacc))
	_acgc := func(_ebfga *ruling) rulingList {
		_eagecc := rulingList{_ebfga}
		_dgeba[_ebfga] = struct{}{}
		for _, _cffe := range _dacc {
			if _, _caddd := _dgeba[_cffe]; _caddd {
				continue
			}
			for _, _dcaed := range _eagecc {
				if _cffe.alignsSec(_dcaed) {
					_eagecc = append(_eagecc, _cffe)
					_dgeba[_cffe] = struct{}{}
					break
				}
			}
		}
		return _eagecc
	}
	_bdgda := []rulingList{_acgc(_dacc[0])}
	for _, _agbb := range _dacc[1:] {
		if _, _cbeec := _dgeba[_agbb]; _cbeec {
			continue
		}
		_bdgda = append(_bdgda, _acgc(_agbb))
	}
	return _bdgda
}

type gridTile struct {
	_bg.PdfRectangle
	_acffg, _acef, _dbbba, _bgdga bool
}

func (_dbaa *textObject) renderText(_dad _gbc.PdfObject, _edec []byte, _bdcg int, _bcae string) error {
	if _dbaa._gea {
		_b.Log.Debug("\u0072\u0065\u006e\u0064\u0065r\u0054\u0065\u0078\u0074\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0066\u006f\u006e\u0074\u002e\u0020\u004e\u006f\u0074\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u002e")
		return nil
	}
	_adgbb := _dbaa.getCurrentFont()
	_ccc := _adgbb.BytesToCharcodes(_edec)
	_eccd, _cag, _cgbd := _adgbb.CharcodesToStrings(_ccc, _bcae)
	if _cgbd > 0 {
		_b.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _cag, _cgbd)
	}
	_dbaa._egf._gfff += _cag
	_dbaa._egf._ffbf += _cgbd
	_gfge := _dbaa._egf
	_gcf := _gfge._ebab
	_acgg := _gfge._ffdc / 100.0
	_bdd := _degf
	if _adgbb.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_bdd = 1
	}
	_bffe, _gaef := _adgbb.GetRuneMetrics(' ')
	if !_gaef {
		_bffe, _gaef = _adgbb.GetCharMetrics(32)
	}
	if !_gaef {
		_bffe, _ = _bg.DefaultFont().GetRuneMetrics(' ')
	}
	_bfe := _bffe.Wx * _bdd
	_b.Log.Trace("\u0073p\u0061\u0063e\u0057\u0069\u0064t\u0068\u003d\u0025\u002e\u0032\u0066\u0020t\u0065\u0078\u0074\u003d\u0025\u0071 \u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0066\u006f\u006et\u0053\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066", _bfe, _eccd, _adgbb, _gcf)
	_gegg := _ef.NewMatrix(_gcf*_acgg, 0, 0, _gcf, 0, _gfge._dbca)
	if _eadc {
		_b.Log.Info("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0074\u0065\u0078t\u0073\u003d\u0025\u0071", len(_ccc), _ccc, _eccd)
	}
	_b.Log.Trace("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0072\u0075\u006ee\u0073\u003d\u0025\u0071", len(_ccc), _ccc, len(_eccd))
	_ecaad := _dbaa.getFillColor()
	_bdca := _dbaa.getStrokeColor()
	for _cdbd, _bdcac := range _eccd {
		_bega := []rune(_bdcac)
		if len(_bega) == 1 && _bega[0] == '\x00' {
			continue
		}
		_bfec := _ccc[_cdbd]
		_edb := _dbaa._dba.CTM.Mult(_dbaa._gbcf).Mult(_gegg)
		_ggd := 0.0
		if len(_bega) == 1 && _bega[0] == 32 {
			_ggd = _gfge._eaed
		}
		_geab, _abfb := _adgbb.GetCharMetrics(_bfec)
		if !_abfb {
			_b.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u004e\u006f \u006d\u0065\u0074r\u0069\u0063\u0020\u0066\u006f\u0072\u0020\u0063\u006fde\u003d\u0025\u0064 \u0072\u003d0\u0078\u0025\u0030\u0034\u0078\u003d%\u002b\u0071 \u0025\u0073", _bfec, _bega, _bega, _adgbb)
			return _d.Errorf("\u006e\u006f\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073:\u0020f\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0063\u006f\u0064\u0065\u003d\u0025\u0064", _adgbb.String(), _bfec)
		}
		_cce := _ef.Point{X: _geab.Wx * _bdd, Y: _geab.Wy * _bdd}
		_aagb := _ef.Point{X: (_cce.X*_gcf + _ggd) * _acgg}
		_baff := _ef.Point{X: (_cce.X*_gcf + _gfge._fgdd + _ggd) * _acgg}
		if _eadc {
			_b.Log.Info("\u0074\u0066\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0063\u003d\u0025\u002e\u0032f\u0020t\u0077\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _gcf, _gfge._fgdd, _gfge._eaed, _acgg)
			_b.Log.Info("\u0064x\u002c\u0064\u0079\u003d%\u002e\u0033\u0066\u0020\u00740\u003d%\u002e3\u0066\u0020\u0074\u003d\u0025\u002e\u0033f", _cce, _aagb, _baff)
		}
		_ggad := _bcca(_aagb)
		_edcf := _bcca(_baff)
		_dgbb := _dbaa._dba.CTM.Mult(_dbaa._gbcf).Mult(_ggad)
		if _acgb {
			_b.Log.Info("e\u006e\u0064\u003a\u000a\tC\u0054M\u003d\u0025\u0073\u000a\u0009 \u0074\u006d\u003d\u0025\u0073\u000a"+"\u0009\u0020t\u0064\u003d\u0025s\u0020\u0078\u006c\u0061\u0074\u003d\u0025\u0073\u000a"+"\u0009t\u0064\u0030\u003d\u0025s\u000a\u0009\u0020\u0020\u2192 \u0025s\u0020x\u006c\u0061\u0074\u003d\u0025\u0073", _dbaa._dba.CTM, _dbaa._gbcf, _edcf, _cfg(_dbaa._dba.CTM.Mult(_dbaa._gbcf).Mult(_edcf)), _ggad, _dgbb, _cfg(_dgbb))
		}
		_ffc, _adcd := _dbaa.newTextMark(_abg.ExpandLigatures(_bega), _edb, _cfg(_dgbb), _bb.Abs(_bfe*_edb.ScalingFactorX()), _adgbb, _dbaa._egf._fgdd, _ecaad, _bdca, _dad, _eccd, _cdbd, _bdcg)
		if !_adcd {
			_b.Log.Debug("\u0054\u0065\u0078\u0074\u0020\u006d\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069d\u0065 \u0070\u0061\u0067\u0065\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		if _adgbb == nil {
			_b.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u002e")
		} else if _adgbb.Encoder() == nil {
			_b.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020N\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u002e\u0020\u0066o\u006et\u003d\u0025\u0073", _adgbb)
		} else {
			if _ffcc, _ffde := _adgbb.Encoder().CharcodeToRune(_bfec); _ffde {
				_ffc._aca = string(_ffcc)
			}
		}
		_b.Log.Trace("i\u003d\u0025\u0064\u0020\u0063\u006fd\u0065\u003d\u0025\u0064\u0020\u006d\u0061\u0072\u006b=\u0025\u0073\u0020t\u0072m\u003d\u0025\u0073", _cdbd, _bfec, _ffc, _edb)
		_dbaa._gaa = append(_dbaa._gaa, &_ffc)
		_dbaa._gbcf.Concat(_edcf)
	}
	return nil
}
func (_aagg rulingList) vertsHorzs() (rulingList, rulingList) {
	var _ebba, _eagf rulingList
	for _, _cgac := range _aagg {
		switch _cgac._gcafd {
		case _cege:
			_ebba = append(_ebba, _cgac)
		case _bfcbd:
			_eagf = append(_eagf, _cgac)
		}
	}
	return _ebba, _eagf
}
func (_bcfe rulingList) bbox() _bg.PdfRectangle {
	var _gcbd _bg.PdfRectangle
	if len(_bcfe) == 0 {
		_b.Log.Error("r\u0075\u006c\u0069\u006e\u0067\u004ci\u0073\u0074\u002e\u0062\u0062\u006f\u0078\u003a\u0020n\u006f\u0020\u0072u\u006ci\u006e\u0067\u0073")
		return _bg.PdfRectangle{}
	}
	if _bcfe[0]._gcafd == _bfcbd {
		_gcbd.Llx, _gcbd.Urx = _bcfe.secMinMax()
		_gcbd.Lly, _gcbd.Ury = _bcfe.primMinMax()
	} else {
		_gcbd.Llx, _gcbd.Urx = _bcfe.primMinMax()
		_gcbd.Lly, _gcbd.Ury = _bcfe.secMinMax()
	}
	return _gcbd
}
func _adfaa(_gged, _bfcg int) int {
	if _gged > _bfcg {
		return _gged
	}
	return _bfcg
}
func (_gcbg paraList) computeEBBoxes() {
	if _edaef {
		_b.Log.Info("\u0063o\u006dp\u0075\u0074\u0065\u0045\u0042\u0042\u006f\u0078\u0065\u0073\u003a")
	}
	for _, _gegge := range _gcbg {
		_gegge._dgcfe = _gegge.PdfRectangle
	}
	_ceff := _gcbg.yNeighbours(0)
	for _acccf, _dgdcf := range _gcbg {
		_eefd := _dgdcf._dgcfe
		_agcd, _ggdcc := -1.0e9, +1.0e9
		for _, _fbaa := range _ceff[_dgdcf] {
			_edggdc := _gcbg[_fbaa]._dgcfe
			if _edggdc.Urx < _eefd.Llx {
				_agcd = _bb.Max(_agcd, _edggdc.Urx)
			} else if _eefd.Urx < _edggdc.Llx {
				_ggdcc = _bb.Min(_ggdcc, _edggdc.Llx)
			}
		}
		for _bcaefc, _ccac := range _gcbg {
			_bfabe := _ccac._dgcfe
			if _acccf == _bcaefc || _bfabe.Ury > _eefd.Lly {
				continue
			}
			if _agcd <= _bfabe.Llx && _bfabe.Llx < _eefd.Llx {
				_eefd.Llx = _bfabe.Llx
			} else if _bfabe.Urx <= _ggdcc && _eefd.Urx < _bfabe.Urx {
				_eefd.Urx = _bfabe.Urx
			}
		}
		if _edaef {
			_d.Printf("\u0025\u0034\u0064\u003a %\u0036\u002e\u0032\u0066\u2192\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0071\u000a", _acccf, _dgdcf._dgcfe, _eefd, _cfbg(_dgdcf.text(), 50))
		}
		_dgdcf._dgcfe = _eefd
	}
	if _daee {
		for _, _gecf := range _gcbg {
			_gecf.PdfRectangle = _gecf._dgcfe
		}
	}
}
func _aaaa(_cfabb []int) []int {
	_bgeg := make([]int, len(_cfabb))
	for _aefef, _fegd := range _cfabb {
		_bgeg[len(_cfabb)-1-_aefef] = _fegd
	}
	return _bgeg
}
func _fdc(_cafd float64) int {
	var _ecca int
	if _cafd >= 0 {
		_ecca = int(_cafd / _dfbad)
	} else {
		_ecca = int(_cafd/_dfbad) - 1
	}
	return _ecca
}

type structElement struct {
	_adfb  string
	_fbcba []structElement
	_agec  int64
	_dbdae _gbc.PdfObject
}

func _gdbe(_bbef []*textWord, _ffab float64, _ccgdbe, _cfc rulingList) *wordBag {
	_bedb := _edbf(_bbef[0], _ffab, _ccgdbe, _cfc)
	for _, _edcg := range _bbef[1:] {
		_cbfb := _fdc(_edcg._ccee)
		_bedb._cccf[_cbfb] = append(_bedb._cccf[_cbfb], _edcg)
		_bedb.PdfRectangle = _egbb(_bedb.PdfRectangle, _edcg.PdfRectangle)
	}
	_bedb.sort()
	return _bedb
}

// String returns a string descibing `i`.
func (_aeafb gridTile) String() string {
	_adgca := func(_bcgde bool, _gedg string) string {
		if _bcgde {
			return _gedg
		}
		return "\u005f"
	}
	return _d.Sprintf("\u00256\u002e2\u0066\u0020\u0025\u0031\u0073%\u0031\u0073%\u0031\u0073\u0025\u0031\u0073", _aeafb.PdfRectangle, _adgca(_aeafb._acef, "\u004c"), _adgca(_aeafb._bgdga, "\u0052"), _adgca(_aeafb._dbbba, "\u0042"), _adgca(_aeafb._acffg, "\u0054"))
}
func _bfaaa(_fcfgc *_bg.Image, _egeb _ab.Color) _f.Image {
	_gecb, _beba := int(_fcfgc.Width), int(_fcfgc.Height)
	_gaeb := _f.NewRGBA(_f.Rect(0, 0, _gecb, _beba))
	for _gbffe := 0; _gbffe < _beba; _gbffe++ {
		for _addf := 0; _addf < _gecb; _addf++ {
			_efbgd, _agaa := _fcfgc.ColorAt(_addf, _gbffe)
			if _agaa != nil {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _addf, _gbffe)
				continue
			}
			_adaba, _bgdea, _bbeg, _ := _efbgd.RGBA()
			var _ecfag _ab.Color
			if _adaba+_bgdea+_bbeg == 0 {
				_ecfag = _ab.Transparent
			} else {
				_ecfag = _egeb
			}
			_gaeb.Set(_addf, _gbffe, _ecfag)
		}
	}
	return _gaeb
}
func (_cggef rulingList) intersections() map[int]intSet {
	var _ddfbf, _geec []int
	for _dfcg, _caab := range _cggef {
		switch _caab._gcafd {
		case _cege:
			_ddfbf = append(_ddfbf, _dfcg)
		case _bfcbd:
			_geec = append(_geec, _dfcg)
		}
	}
	if len(_ddfbf) < _fgbd+1 || len(_geec) < _abfd+1 {
		return nil
	}
	if len(_ddfbf)+len(_geec) > _bbag {
		_b.Log.Debug("\u0069\u006e\u0074\u0065\u0072\u0073e\u0063\u0074\u0069\u006f\u006e\u0073\u003a\u0020\u0054\u004f\u004f\u0020\u004d\u0041\u004e\u0059\u0020\u0072\u0075\u006ci\u006e\u0067\u0073\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u003d\u0020%\u0064 \u0078\u0020\u0025\u0064", len(_cggef), len(_ddfbf), len(_geec))
		return nil
	}
	_ggced := make(map[int]intSet, len(_ddfbf)+len(_geec))
	for _, _dccb := range _ddfbf {
		for _, _fbga := range _geec {
			if _cggef[_dccb].intersects(_cggef[_fbga]) {
				if _, _ffac := _ggced[_dccb]; !_ffac {
					_ggced[_dccb] = make(intSet)
				}
				if _, _abge := _ggced[_fbga]; !_abge {
					_ggced[_fbga] = make(intSet)
				}
				_ggced[_dccb].add(_fbga)
				_ggced[_fbga].add(_dccb)
			}
		}
	}
	return _ggced
}
func (_fab *stateStack) top() *textState {
	if _fab.empty() {
		return nil
	}
	return (*_fab)[_fab.size()-1]
}
func _bgcd(_befg []_gbc.PdfObject) (_dcgc, _acbaf float64, _cbabca error) {
	if len(_befg) != 2 {
		return 0, 0, _d.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0073\u003a \u0025\u0064", len(_befg))
	}
	_eacf, _cbabca := _gbc.GetNumbersAsFloat(_befg)
	if _cbabca != nil {
		return 0, 0, _cbabca
	}
	return _eacf[0], _eacf[1], nil
}
func _cafb(_gffe *wordBag, _cbgd *textWord, _cefbe float64) bool {
	return _cbgd.Llx < _gffe.Urx+_cefbe && _gffe.Llx-_cefbe < _cbgd.Urx
}
func _badd(_gbeb, _ged _bg.PdfRectangle) bool { return _gbeb.Lly <= _ged.Ury && _ged.Lly <= _gbeb.Ury }
func _gbaa(_dcag []*wordBag) []*wordBag {
	if len(_dcag) <= 1 {
		return _dcag
	}
	if _cgdff {
		_b.Log.Info("\u006d\u0065\u0072\u0067\u0065\u0057\u006f\u0072\u0064B\u0061\u0067\u0073\u003a")
	}
	_gd.Slice(_dcag, func(_gdba, _cdac int) bool {
		_bbea, _cfcc := _dcag[_gdba], _dcag[_cdac]
		_dgaee := _bbea.Width() * _bbea.Height()
		_gade := _cfcc.Width() * _cfcc.Height()
		if _dgaee != _gade {
			return _dgaee > _gade
		}
		if _bbea.Height() != _cfcc.Height() {
			return _bbea.Height() > _cfcc.Height()
		}
		return _gdba < _cdac
	})
	var _ecaf []*wordBag
	_gfce := make(intSet)
	for _gaac := 0; _gaac < len(_dcag); _gaac++ {
		if _gfce.has(_gaac) {
			continue
		}
		_edcgd := _dcag[_gaac]
		for _edgb := _gaac + 1; _edgb < len(_dcag); _edgb++ {
			if _gfce.has(_gaac) {
				continue
			}
			_eade := _dcag[_edgb]
			_bffec := _edcgd.PdfRectangle
			_bffec.Llx -= _edcgd._aced
			if _ffff(_bffec, _eade.PdfRectangle) {
				_edcgd.absorb(_eade)
				_gfce.add(_edgb)
			}
		}
		_ecaf = append(_ecaf, _edcgd)
	}
	if len(_dcag) != len(_ecaf)+len(_gfce) {
		_b.Log.Error("\u006d\u0065\u0072ge\u0057\u006f\u0072\u0064\u0042\u0061\u0067\u0073\u003a \u0025d\u2192%\u0064 \u0061\u0062\u0073\u006f\u0072\u0062\u0065\u0064\u003d\u0025\u0064", len(_dcag), len(_ecaf), len(_gfce))
	}
	return _ecaf
}
func _efc(_dbf []rune) BidiText {
	_eda := -1
	_gfb := false
	_cba := true
	_bd := len(_dbf)
	_gfd := make([]string, _bd)
	_fea := make([]string, _bd)
	if _bd == 0 || _gfb {
		return _ce(string(_dbf), _cba, _gfb)
	}
	_fa := 0
	for _gdf, _be := range _dbf {
		_gfd[_gdf] = string(_be)
		_df := "\u004c"
		if _be <= 0x00ff {
			_df = _de[_be]
		} else if 0x0590 <= _be && _be <= 0x05f4 {
			_df = "\u0052"
		} else if 0x0600 <= _be && _be <= 0x06ff {
			_ade := _be & 0xff
			if int(_ade) >= len(_db) {
				_b.Log.Debug("\u0042\u0069\u0064\u0069\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0055n\u0069c\u006f\u0064\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020" + string(_be))
			}
			_df = _db[_be&0xff]
		} else if (0x0700 <= _be && _be <= 0x08ac) || (0xfb50 <= _be && _be <= 0xfdff) || (0xfe70 <= _be && _be <= 0xfeff) {
			_df = "\u0041\u004c"
		}
		if _df == "\u0052" || _df == "\u0041\u004c" || _df == "\u0041\u004e" {
			_fa++
		}
		_fea[_gdf] = _df
	}
	if _fa == 0 {
		_cba = true
		return _ce(string(_dbf), _cba, false)
	}
	if _eda == -1 {
		if float64(_fa)/float64(_bd) < 0.3 && _bd > 4 {
			_cba = true
			_eda = 0
		} else {
			_cba = false
			_eda = 1
		}
	}
	var _edf []int
	for range _dbf {
		_edf = append(_edf, _eda)
	}
	_eaa := "\u004c"
	if _caed(_eda) {
		_eaa = "\u0052"
	}
	_dgd := _eaa
	_ebc := _dgd
	_dcb := _dgd
	for _ff := range _dbf {
		if _fea[_ff] == "\u004e\u0053\u004d" {
			_fea[_ff] = _dcb
		} else {
			_dcb = _fea[_ff]
		}
	}
	_dcb = _dgd
	var _beg string
	for _fc := range _dbf {
		_beg = _fea[_fc]
		if _beg == "\u0045\u004e" {
			if _dcb == "\u0041\u004c" {
				_fea[_fc] = "\u0041\u004e"
			} else {
				_fea[_fc] = "\u0045\u004e"
			}
		} else if _beg == "\u0052" || _beg == "\u004c" || _beg == "\u0041\u004c" {
			_dcb = _beg
		}
	}
	for _fb := range _dbf {
		_edfb := _fea[_fb]
		if _edfb == "\u0041\u004c" {
			_fea[_fb] = "\u0052"
		}
	}
	for _dbd := 1; _dbd < (len(_dbf) - 1); _dbd++ {
		if _fea[_dbd] == "\u0045\u0053" && _fea[_dbd-1] == "\u0045\u004e" && _fea[_dbd+1] == "\u0045\u004e" {
			_fea[_dbd] = "\u0045\u004e"
		}
		if _fea[_dbd] == "\u0043\u0053" && (_fea[_dbd-1] == "\u0045\u004e" || _fea[_dbd-1] == "\u0041\u004e") && _fea[_dbd+1] == _fea[_dbd-1] {
			_fea[_dbd] = _fea[_dbd-1]
		}
	}
	for _abf := range _dbf {
		if _fea[_abf] == "\u0045\u004e" {
			for _dgcc := _abf - 1; _dgcc >= 0; _dgcc-- {
				if _fea[_dgcc] != "\u0045\u0054" {
					break
				}
				_fea[_dgcc] = "\u0045\u004e"
			}
			for _gff := _abf + 1; _gff < _bd; _gff++ {
				if _fea[_gff] != "\u0045\u0054" {
					break
				}
				_fea[_gff] = "\u0045\u004e"
			}
		}
	}
	for _gbb := range _dbf {
		_cec := _fea[_gbb]
		if _cec == "\u0057\u0053" || _cec == "\u0045\u0053" || _cec == "\u0045\u0054" || _cec == "\u0043\u0053" {
			_fea[_gbb] = "\u004f\u004e"
		}
	}
	_dcb = "\u0073\u006f\u0072"
	for _gg := range _dbf {
		_dfa := _fea[_gg]
		if _dfa == "\u0045\u004e" {
			if _dcb == "\u004c" {
				_fea[_gg] = "\u004c"
			} else {
				_fea[_gg] = "\u0045\u004e"
			}
		} else if _dfa == "\u0052" || _dfa == "\u004c" {
			_dcb = _dfa
		}
	}
	for _aga := 0; _aga < len(_dbf); _aga++ {
		if _fea[_aga] == "\u004f\u004e" {
			_bdf := _bbe(_fea, _aga+1, "\u004f\u004e")
			_abd := _ebc
			if _aga > 0 {
				_abd = _fea[_aga-1]
			}
			_eg := _ebc
			if _bdf+1 < _bd {
				_eg = _fea[_bdf+1]
			}
			if _abd != "\u004c" {
				_abd = "\u0052"
			}
			if _eg != "\u004c" {
				_eg = "\u0052"
			}
			if _abd == _eg {
				_ad(_fea, _aga, _bdf, _abd)
			}
			_aga = _bdf - 1
		}
	}
	for _febc := range _dbf {
		if _fea[_febc] == "\u004f\u004e" {
			_fea[_febc] = _eaa
		}
	}
	for _afb := range _dbf {
		_ebd := _fea[_afb]
		if _afa(_edf[_afb]) {
			if _ebd == "\u0052" {
				_edf[_afb]++
			} else if _ebd == "\u0041\u004e" || _ebd == "\u0045\u004e" {
				_edf[_afb] += 2
			}
		} else if _ebd == "\u004c" || _ebd == "\u0041\u004e" || _ebd == "\u0045\u004e" {
			_edf[_afb]++
		}
	}
	_ec := -1
	_gbbf := 99
	var _eag int
	for _bgfc := 0; _bgfc < len(_edf); _bgfc++ {
		_eag = _edf[_bgfc]
		if _ec < _eag {
			_ec = _eag
		}
		if _gbbf > _eag && _caed(_eag) {
			_gbbf = _eag
		}
	}
	for _bab := _ec; _bab >= _gbbf; _bab-- {
		_eef := -1
		for _abgc := 0; _abgc < len(_edf); _abgc++ {
			if _edf[_abgc] < _bab {
				if _eef >= 0 {
					_dd(_gfd, _eef, _abgc)
					_eef = -1
				}
			} else if _eef < 0 {
				_eef = _abgc
			}
		}
		if _eef >= 0 {
			_dd(_gfd, _eef, len(_edf))
		}
	}
	for _dbc := 0; _dbc < len(_gfd); _dbc++ {
		_fed := _gfd[_dbc]
		if _fed == "\u003c" || _fed == "\u003e" {
			_gfd[_dbc] = ""
		}
	}
	return _ce(_ge.Join(_gfd, ""), _cba, false)
}

// PageFonts represents extracted fonts on a PDF page.
type PageFonts struct{ Fonts []Font }

func _egbb(_bddg, _gegc _bg.PdfRectangle) _bg.PdfRectangle {
	return _bg.PdfRectangle{Llx: _bb.Min(_bddg.Llx, _gegc.Llx), Lly: _bb.Min(_bddg.Lly, _gegc.Lly), Urx: _bb.Max(_bddg.Urx, _gegc.Urx), Ury: _bb.Max(_bddg.Ury, _gegc.Ury)}
}

const (
	_edaef = false
	_fcee  = false
	_dfga  = false
	_acgb  = false
	_fbfd  = false
	_eadc  = false
	_gdgf  = false
	_agba  = false
	_cgdff = false
	_gfeec = _cgdff && true
	_eddg  = _gfeec && false
	_bfab  = _cgdff && true
	_eeeg  = false
	_dcfc  = _eeeg && false
	_gbcg  = _eeeg && true
	_gfab  = false
	_fafd  = _gfab && false
	_bffed = _gfab && false
	_degg  = _gfab && true
	_bfgg  = _gfab && false
	_bbae  = _gfab && false
)

var _cdcg = []string{"\u2756", "\u27a2", "\u2713", "\u2022", "\uf0a7", "\u25a1", "\u2212", "\u25a0", "\u25aa", "\u006f"}

func (_fcad *shapesState) cubicTo(_beb, _geae, _ebdd, _dcd, _egab, _agab float64) {
	if _fbfd {
		_b.Log.Info("\u0063\u0075\u0062\u0069\u0063\u0054\u006f\u003a")
	}
	_fcad.addPoint(_egab, _agab)
}
func (_gfceac rulingList) asTiling() gridTiling {
	if _degg {
		_b.Log.Info("r\u0075\u006ci\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0063s\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u002b\u002b\u002b\u0020\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_gfceac))
	}
	for _dfbc, _eeee := range _gfceac[1:] {
		_bdea := _gfceac[_dfbc]
		if _bdea.alignsPrimary(_eeee) && _bdea.alignsSec(_eeee) {
			_b.Log.Error("a\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0044\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065 \u0072\u0075\u006c\u0069\u006e\u0067\u0073\u002e\u000a\u0009v=\u0025\u0073\u000a\t\u0077=\u0025\u0073", _eeee, _bdea)
		}
	}
	_gfceac.sortStrict()
	_gfceac.log("\u0073n\u0061\u0070\u0070\u0065\u0064")
	_ecdef, _ddea := _gfceac.vertsHorzs()
	_bgcfb := _ecdef.primaries()
	_dgce := _ddea.primaries()
	_fbge := len(_bgcfb) - 1
	_acbf := len(_dgce) - 1
	if _fbge == 0 || _acbf == 0 {
		return gridTiling{}
	}
	_gccg := _bg.PdfRectangle{Llx: _bgcfb[0], Urx: _bgcfb[_fbge], Lly: _dgce[0], Ury: _dgce[_acbf]}
	if _degg {
		_b.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0072\u0074s=\u0025\u0064", len(_ecdef))
		for _fgfd, _gggb := range _ecdef {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _fgfd, _gggb)
		}
		_b.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0068\u006f\u0072\u007as=\u0025\u0064", len(_ddea))
		for _cbfd, _acgff := range _ddea {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cbfd, _acgff)
		}
		_b.Log.Info("\u0072\u0075\u006c\u0069\u006eg\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067:\u0020\u0020\u0077\u0078\u0068\u003d\u0025\u0064\u0078\u0025\u0064\u000a\u0009\u006c\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a\u0009\u006c\u006c\u0079\u003d\u0025\u002e\u0032f", _fbge, _acbf, _bgcfb, _dgce)
	}
	_gcbe := make([]gridTile, _fbge*_acbf)
	for _cagg := _acbf - 1; _cagg >= 0; _cagg-- {
		_gddf := _dgce[_cagg]
		_bbeag := _dgce[_cagg+1]
		for _dffe := 0; _dffe < _fbge; _dffe++ {
			_fage := _bgcfb[_dffe]
			_caag := _bgcfb[_dffe+1]
			_afgaf := _ecdef.findPrimSec(_fage, _gddf)
			_faca := _ecdef.findPrimSec(_caag, _gddf)
			_deeb := _ddea.findPrimSec(_gddf, _fage)
			_afgd := _ddea.findPrimSec(_bbeag, _fage)
			_cfed := _bg.PdfRectangle{Llx: _fage, Urx: _caag, Lly: _gddf, Ury: _bbeag}
			_dege := _fefe(_cfed, _afgaf, _faca, _deeb, _afgd)
			_gcbe[_cagg*_fbge+_dffe] = _dege
			if _degg {
				_d.Printf("\u0020\u0020\u0078\u003d\u0025\u0032\u0064\u0020\u0079\u003d\u0025\u0032\u0064\u003a\u0020%\u0073 \u0025\u0036\u002e\u0032\u0066\u0020\u0078\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _dffe, _cagg, _dege.String(), _dege.Width(), _dege.Height())
			}
		}
	}
	if _degg {
		_b.Log.Info("r\u0075\u006c\u0069\u006e\u0067\u004c\u0069\u0073\u0074.\u0061\u0073\u0054\u0069\u006c\u0069\u006eg:\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0020\u0068\u006f\u0072\u0069zo\u006e\u0074a\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _gccg)
	}
	_edceff := make([]map[float64]gridTile, _acbf)
	for _dbfg := _acbf - 1; _dbfg >= 0; _dbfg-- {
		if _degg {
			_d.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _dbfg)
		}
		_edceff[_dbfg] = make(map[float64]gridTile, _fbge)
		for _gddd := 0; _gddd < _fbge; _gddd++ {
			_dbgbd := _gcbe[_dbfg*_fbge+_gddd]
			if _degg {
				_d.Printf("\u0020\u0020\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gddd, _dbgbd)
			}
			if !_dbgbd._acef {
				continue
			}
			_bfdc := _gddd
			for _geffg := _gddd + 1; !_dbgbd._bgdga && _geffg < _fbge; _geffg++ {
				_aegafa := _gcbe[_dbfg*_fbge+_geffg]
				_dbgbd.Urx = _aegafa.Urx
				_dbgbd._acffg = _dbgbd._acffg || _aegafa._acffg
				_dbgbd._dbbba = _dbgbd._dbbba || _aegafa._dbbba
				_dbgbd._bgdga = _aegafa._bgdga
				if _degg {
					_d.Printf("\u0020 \u0020%\u0034\u0064\u003a\u0020\u0025s\u0020\u2192 \u0025\u0073\u000a", _geffg, _aegafa, _dbgbd)
				}
				_bfdc = _geffg
			}
			if _degg {
				_d.Printf(" \u0020 \u0025\u0032\u0064\u0020\u002d\u0020\u0025\u0032d\u0020\u2192\u0020\u0025s\n", _gddd, _bfdc, _dbgbd)
			}
			_gddd = _bfdc
			_edceff[_dbfg][_dbgbd.Llx] = _dbgbd
		}
	}
	_cfeb := make(map[float64]map[float64]gridTile, _acbf)
	_gddda := make(map[float64]map[float64]struct{}, _acbf)
	for _gede := _acbf - 1; _gede >= 0; _gede-- {
		_edggb := _gcbe[_gede*_fbge].Lly
		_cfeb[_edggb] = make(map[float64]gridTile, _fbge)
		_gddda[_edggb] = make(map[float64]struct{}, _fbge)
	}
	if _degg {
		_b.Log.Info("\u0072u\u006c\u0069n\u0067\u004c\u0069s\u0074\u002e\u0061\u0073\u0054\u0069\u006ci\u006e\u0067\u003a\u0020\u0063\u006fa\u006c\u0065\u0073\u0063\u0065\u0020\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _gccg)
	}
	for _gbad := _acbf - 1; _gbad >= 0; _gbad-- {
		_gbae := _gcbe[_gbad*_fbge].Lly
		_daafc := _edceff[_gbad]
		if _degg {
			_d.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _gbad)
		}
		for _, _ecfc := range _ebbf(_daafc) {
			if _, _fgda := _gddda[_gbae][_ecfc]; _fgda {
				continue
			}
			_dfcf := _daafc[_ecfc]
			if _degg {
				_d.Printf(" \u0020\u0020\u0020\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _dfcf.String())
			}
			for _effb := _gbad - 1; _effb >= 0; _effb-- {
				if _dfcf._dbbba {
					break
				}
				_dfee := _edceff[_effb]
				_cega, _ccbag := _dfee[_ecfc]
				if !_ccbag {
					break
				}
				if _cega.Urx != _dfcf.Urx {
					break
				}
				_dfcf._dbbba = _cega._dbbba
				_dfcf.Lly = _cega.Lly
				if _degg {
					_d.Printf("\u0020\u0020\u0020\u0020  \u0020\u0020\u0076\u003d\u0025\u0073\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _cega.String(), _dfcf.String())
				}
				_gddda[_cega.Lly][_cega.Llx] = struct{}{}
			}
			if _gbad == 0 {
				_dfcf._dbbba = true
			}
			if _dfcf.complete() {
				_cfeb[_gbae][_ecfc] = _dfcf
			}
		}
	}
	_dcfgd := gridTiling{PdfRectangle: _gccg, _cdfg: _dbadc(_cfeb), _fgcd: _bcgc(_cfeb), _egff: _cfeb}
	_dcfgd.log("\u0043r\u0065\u0061\u0074\u0065\u0064")
	return _dcfgd
}
func (_dfeac *textPara) text() string {
	_fgcca := new(_aa.Buffer)
	_dfeac.writeText(_fgcca)
	return _fgcca.String()
}
func (_dfdf *textObject) getFont(_bffea string) (*_bg.PdfFont, error) {
	if _dfdf._eadb._deg != nil {
		_abeb, _ggdg := _dfdf.getFontDict(_bffea)
		if _ggdg != nil {
			_b.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0067\u0065\u0074\u0046\u006f\u006e\u0074:\u0020n\u0061m\u0065=\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _bffea, _ggdg.Error())
			return nil, _ggdg
		}
		_dfdf._eadb._cd++
		_bdcab, _cgcce := _dfdf._eadb._deg[_abeb.String()]
		if _cgcce {
			_bdcab._afgb = _dfdf._eadb._cd
			return _bdcab._cbaf, nil
		}
	}
	_gfbg, _ddgd := _dfdf.getFontDict(_bffea)
	if _ddgd != nil {
		return nil, _ddgd
	}
	_cab, _ddgd := _dfdf.getFontDirect(_bffea)
	if _ddgd != nil {
		return nil, _ddgd
	}
	if _dfdf._eadb._deg != nil {
		_bedd := fontEntry{_cab, _dfdf._eadb._cd}
		if len(_dfdf._eadb._deg) >= _fbfa {
			var _addb []string
			for _ffcb := range _dfdf._eadb._deg {
				_addb = append(_addb, _ffcb)
			}
			_gd.Slice(_addb, func(_adeb, _eeaa int) bool {
				return _dfdf._eadb._deg[_addb[_adeb]]._afgb < _dfdf._eadb._deg[_addb[_eeaa]]._afgb
			})
			delete(_dfdf._eadb._deg, _addb[0])
		}
		_dfdf._eadb._deg[_gfbg.String()] = _bedd
	}
	return _cab, nil
}
func (_ggae rulingList) snapToGroups() rulingList {
	_fgbce, _eaadf := _ggae.vertsHorzs()
	if len(_fgbce) > 0 {
		_fgbce = _fgbce.snapToGroupsDirection()
	}
	if len(_eaadf) > 0 {
		_eaadf = _eaadf.snapToGroupsDirection()
	}
	_cbd := append(_fgbce, _eaadf...)
	_cbd.log("\u0073\u006e\u0061p\u0054\u006f\u0047\u0072\u006f\u0075\u0070\u0073")
	return _cbd
}

type textState struct {
	_fgdd float64
	_eaed float64
	_ffdc float64
	_fba  float64
	_ebab float64
	_dge  RenderMode
	_dbca float64
	_aede *_bg.PdfFont
	_bbd  _bg.PdfRectangle
	_gfff int
	_ffbf int
}

func (_dbfc *TextMarkArray) exists(_bgag TextMark) bool {
	for _, _gfffd := range _dbfc.Elements() {
		if _af.DeepEqual(_bgag.DirectObject, _gfffd.DirectObject) && _af.DeepEqual(_bgag.BBox, _gfffd.BBox) && _gfffd.Text == _bgag.Text {
			return true
		}
	}
	return false
}

// Elements returns the TextMarks in `ma`.
func (_begb *TextMarkArray) Elements() []TextMark { return _begb._egbe }

type bounded interface{ bbox() _bg.PdfRectangle }

func (_eced *wordBag) pullWord(_bafa *textWord, _fdfg int, _fddd map[int]map[*textWord]struct{}) {
	_eced.PdfRectangle = _egbb(_eced.PdfRectangle, _bafa.PdfRectangle)
	if _bafa._dafae > _eced._aced {
		_eced._aced = _bafa._dafae
	}
	_eced._cccf[_fdfg] = append(_eced._cccf[_fdfg], _bafa)
	_fddd[_fdfg][_bafa] = struct{}{}
}
func _dbfbe(_dade, _bbggb, _dcgea, _ecbc *textPara) *textTable {
	_bbaca := &textTable{_ddcega: 2, _egbbf: 2, _bcec: make(map[uint64]*textPara, 4)}
	_bbaca.put(0, 0, _dade)
	_bbaca.put(1, 0, _bbggb)
	_bbaca.put(0, 1, _dcgea)
	_bbaca.put(1, 1, _ecbc)
	return _bbaca
}
func (_ebf *imageExtractContext) extractInlineImage(_daf *_aed.ContentStreamInlineImage, _dgag _aed.GraphicsState, _ccb *_bg.PdfPageResources) error {
	_cefb, _gac := _daf.ToImage(_ccb)
	if _gac != nil {
		return _gac
	}
	_fff, _gac := _daf.GetColorSpace(_ccb)
	if _gac != nil {
		return _gac
	}
	if _fff == nil {
		_fff = _bg.NewPdfColorspaceDeviceGray()
	}
	_bff, _gac := _fff.ImageToRGB(*_cefb)
	if _gac != nil {
		return _gac
	}
	_gba := ImageMark{Image: &_bff, Width: _dgag.CTM.ScalingFactorX(), Height: _dgag.CTM.ScalingFactorY(), Angle: _dgag.CTM.Angle()}
	_gba.X, _gba.Y = _dgag.CTM.Translation()
	_ebf._aaa = append(_ebf._aaa, _gba)
	_ebf._egbg++
	return nil
}

type textPara struct {
	_bg.PdfRectangle
	_dgcfe _bg.PdfRectangle
	_gdc   []*textLine
	_cece  *textTable
	_bfaab bool
	_ecee  bool
	_agge  *textPara
	_gbfc  *textPara
	_ggfda *textPara
	_adef  *textPara
	_bgea  []list
}

var _dgeda = map[markKind]string{_gffff: "\u0073\u0074\u0072\u006f\u006b\u0065", _aead: "\u0066\u0069\u006c\u006c", _cccg: "\u0061u\u0067\u006d\u0065\u006e\u0074"}

func _eged(_eebab map[int]intSet) []int {
	_gfgfg := make([]int, 0, len(_eebab))
	for _fgfa := range _eebab {
		_gfgfg = append(_gfgfg, _fgfa)
	}
	_gd.Ints(_gfgfg)
	return _gfgfg
}

type shapesState struct {
	_bdab _ef.Matrix
	_gbag _ef.Matrix
	_dgdd []*subpath
	_aaba bool
	_gbgg _ef.Point
	_eadd *textObject
}

const (
	_degfc = true
	_fbec  = true
	_afgc  = true
	_daee  = false
	_cebf  = false
	_dgddb = 6
	_cecf  = 3.0
	_fdgb  = 200
	_dgaa  = true
	_bbgf  = true
	_dccd  = true
	_bfac  = true
	_gbbfa = false
)

func (_fcac *subpath) clear() { *_fcac = subpath{} }
func (_fbce *wordBag) arrangeText() *textPara {
	_fbce.sort()
	if _fbec {
		_fbce.removeDuplicates()
	}
	var _cgbf []*textLine
	for _, _bfdb := range _fbce.depthIndexes() {
		for !_fbce.empty(_bfdb) {
			_fcafd := _fbce.firstReadingIndex(_bfdb)
			_ecgd := _fbce.firstWord(_fcafd)
			_bbgd := _adccd(_fbce, _fcafd)
			_fbgf := _ecgd._dafae
			if _fbgf < _gccbd {
				_fbgf = _gccbd
			}
			_cgdef := _ecgd._ccee - _gfgc*_fbgf
			_abc := _ecgd._ccee + _gfgc*_fbgf
			_ebda := _bfcc * _fbgf
			_ggfac := _bggg * _fbgf
		_eded:
			for {
				var _ebcb *textWord
				_fgdg := 0
				for _, _dafaa := range _fbce.depthBand(_cgdef, _abc) {
					_aeef := _fbce.highestWord(_dafaa, _cgdef, _abc)
					if _aeef == nil {
						continue
					}
					_fgdc := _gefae(_aeef, _bbgd._edee[len(_bbgd._edee)-1])
					if _fgdc < -_ggfac {
						break _eded
					}
					if _fgdc > _ebda {
						continue
					}
					if _ebcb != nil && _ccfa(_aeef, _ebcb) >= 0 {
						continue
					}
					_ebcb = _aeef
					_fgdg = _dafaa
				}
				if _ebcb == nil {
					break
				}
				_bbgd.pullWord(_fbce, _ebcb, _fgdg)
			}
			_bbgd.markWordBoundaries()
			_cgbf = append(_cgbf, _bbgd)
		}
	}
	if len(_cgbf) == 0 {
		return nil
	}
	_gd.Slice(_cgbf, func(_bdgfb, _aafgd int) bool { return _cfda(_cgbf[_bdgfb], _cgbf[_aafgd]) < 0 })
	_fbbe := _ebeeg(_fbce.PdfRectangle, _cgbf)
	if _cgdff {
		_b.Log.Info("\u0061\u0072\u0072an\u0067\u0065\u0054\u0065\u0078\u0074\u0020\u0021\u0021\u0021\u0020\u0070\u0061\u0072\u0061\u003d\u0025\u0073", _fbbe.String())
		if _gfeec {
			for _cgeab, _gbba := range _fbbe._gdc {
				_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cgeab, _gbba.String())
				if _eddg {
					for _eeeb, _fdee := range _gbba._edee {
						_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eeeb, _fdee.String())
						for _efcbff, _gfgbd := range _fdee._ecdf {
							_d.Printf("\u00251\u0032\u0064\u003a\u0020\u0025\u0073\n", _efcbff, _gfgbd.String())
						}
					}
				}
			}
		}
	}
	return _fbbe
}
func (_eaagg *wordBag) text() string {
	_abebf := _eaagg.allWords()
	_ffad := make([]string, len(_abebf))
	for _aega, _dgcg := range _abebf {
		_ffad[_aega] = _dgcg._edac
	}
	return _ge.Join(_ffad, "\u0020")
}
func (_bdae paraList) yNeighbours(_cbbd float64) map[*textPara][]int {
	_dagdc := make([]event, 2*len(_bdae))
	if _cbbd == 0 {
		for _ccedd, _gefd := range _bdae {
			_dagdc[2*_ccedd] = event{_gefd.Lly, true, _ccedd}
			_dagdc[2*_ccedd+1] = event{_gefd.Ury, false, _ccedd}
		}
	} else {
		for _eedg, _agceb := range _bdae {
			_dagdc[2*_eedg] = event{_agceb.Lly - _cbbd*_agceb.fontsize(), true, _eedg}
			_dagdc[2*_eedg+1] = event{_agceb.Ury + _cbbd*_agceb.fontsize(), false, _eedg}
		}
	}
	return _bdae.eventNeighbours(_dagdc)
}
func _bdef(_fadgd, _fdgd _ef.Point) bool {
	_bafge := _bb.Abs(_fadgd.X - _fdgd.X)
	_cdbgd := _bb.Abs(_fadgd.Y - _fdgd.Y)
	return _fcge(_cdbgd, _bafge)
}
func (_dffcf paraList) addNeighbours() {
	_bbfa := func(_fgbee []int, _bbad *textPara) ([]*textPara, []*textPara) {
		_bgafd := make([]*textPara, 0, len(_fgbee)-1)
		_ccfbd := make([]*textPara, 0, len(_fgbee)-1)
		for _, _addg := range _fgbee {
			_ffeb := _dffcf[_addg]
			if _ffeb.Urx <= _bbad.Llx {
				_bgafd = append(_bgafd, _ffeb)
			} else if _ffeb.Llx >= _bbad.Urx {
				_ccfbd = append(_ccfbd, _ffeb)
			}
		}
		return _bgafd, _ccfbd
	}
	_bfcce := func(_abbb []int, _bdcace *textPara) ([]*textPara, []*textPara) {
		_eec := make([]*textPara, 0, len(_abbb)-1)
		_aggb := make([]*textPara, 0, len(_abbb)-1)
		for _, _aabgd := range _abbb {
			_dadbe := _dffcf[_aabgd]
			if _dadbe.Ury <= _bdcace.Lly {
				_aggb = append(_aggb, _dadbe)
			} else if _dadbe.Lly >= _bdcace.Ury {
				_eec = append(_eec, _dadbe)
			}
		}
		return _eec, _aggb
	}
	_gcee := _dffcf.yNeighbours(_fgfcb)
	for _, _eceb := range _dffcf {
		_gdde := _gcee[_eceb]
		if len(_gdde) == 0 {
			continue
		}
		_edff, _acebd := _bbfa(_gdde, _eceb)
		if len(_edff) == 0 && len(_acebd) == 0 {
			continue
		}
		if len(_edff) > 0 {
			_gggf := _edff[0]
			for _, _fdfbd := range _edff[1:] {
				if _fdfbd.Urx >= _gggf.Urx {
					_gggf = _fdfbd
				}
			}
			for _, _ffdg := range _edff {
				if _ffdg != _gggf && _ffdg.Urx > _gggf.Llx {
					_gggf = nil
					break
				}
			}
			if _gggf != nil && _badd(_eceb.PdfRectangle, _gggf.PdfRectangle) {
				_eceb._agge = _gggf
			}
		}
		if len(_acebd) > 0 {
			_dacg := _acebd[0]
			for _, _gced := range _acebd[1:] {
				if _gced.Llx <= _dacg.Llx {
					_dacg = _gced
				}
			}
			for _, _gdgecd := range _acebd {
				if _gdgecd != _dacg && _gdgecd.Llx < _dacg.Urx {
					_dacg = nil
					break
				}
			}
			if _dacg != nil && _badd(_eceb.PdfRectangle, _dacg.PdfRectangle) {
				_eceb._gbfc = _dacg
			}
		}
	}
	_gcee = _dffcf.xNeighbours(_eead)
	for _, _eeedg := range _dffcf {
		_ebecd := _gcee[_eeedg]
		if len(_ebecd) == 0 {
			continue
		}
		_ecfa, _deaaf := _bfcce(_ebecd, _eeedg)
		if len(_ecfa) == 0 && len(_deaaf) == 0 {
			continue
		}
		if len(_deaaf) > 0 {
			_bagab := _deaaf[0]
			for _, _faacf := range _deaaf[1:] {
				if _faacf.Ury >= _bagab.Ury {
					_bagab = _faacf
				}
			}
			for _, _dafafb := range _deaaf {
				if _dafafb != _bagab && _dafafb.Ury > _bagab.Lly {
					_bagab = nil
					break
				}
			}
			if _bagab != nil && _degfb(_eeedg.PdfRectangle, _bagab.PdfRectangle) {
				_eeedg._adef = _bagab
			}
		}
		if len(_ecfa) > 0 {
			_faeba := _ecfa[0]
			for _, _gcabe := range _ecfa[1:] {
				if _gcabe.Lly <= _faeba.Lly {
					_faeba = _gcabe
				}
			}
			for _, _gbdf := range _ecfa {
				if _gbdf != _faeba && _gbdf.Lly < _faeba.Ury {
					_faeba = nil
					break
				}
			}
			if _faeba != nil && _degfb(_eeedg.PdfRectangle, _faeba.PdfRectangle) {
				_eeedg._ggfda = _faeba
			}
		}
	}
	for _, _acbcb := range _dffcf {
		if _acbcb._agge != nil && _acbcb._agge._gbfc != _acbcb {
			_acbcb._agge = nil
		}
		if _acbcb._ggfda != nil && _acbcb._ggfda._adef != _acbcb {
			_acbcb._ggfda = nil
		}
		if _acbcb._gbfc != nil && _acbcb._gbfc._agge != _acbcb {
			_acbcb._gbfc = nil
		}
		if _acbcb._adef != nil && _acbcb._adef._ggfda != _acbcb {
			_acbcb._adef = nil
		}
	}
}
func (_geaec *textMark) bbox() _bg.PdfRectangle { return _geaec.PdfRectangle }
func (_dcfdb *textObject) getStrokeColor() _ab.Color {
	return _fbbag(_dcfdb._dba.ColorspaceStroking, _dcfdb._dba.ColorStroking)
}
func _afa(_cg int) bool { return (_cg & 1) == 0 }
func (_dec *imageExtractContext) extractXObjectImage(_dcf *_gbc.PdfObjectName, _efg _aed.GraphicsState, _bcd *_bg.PdfPageResources) error {
	_deb, _ := _bcd.GetXObjectByName(*_dcf)
	if _deb == nil {
		return nil
	}
	_gdgc, _cfd := _dec._bag[_deb]
	if !_cfd {
		_cad, _gca := _bcd.GetXObjectImageByName(*_dcf)
		if _gca != nil {
			return _gca
		}
		if _cad == nil {
			return nil
		}
		_fcf, _gca := _cad.ToImage()
		if _gca != nil {
			return _gca
		}
		var _bcb _f.Image
		if _cad.Mask != nil {
			if _bcb, _gca = _ccacb(_cad.Mask, _ab.Opaque); _gca != nil {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		} else if _cad.SMask != nil {
			_bcb, _gca = _gfccg(_cad.SMask, _ab.Opaque)
			if _gca != nil {
				_b.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		}
		if _bcb != nil {
			_fgd, _abde := _fcf.ToGoImage()
			if _abde != nil {
				return _abde
			}
			_fgd = _cdbae(_fgd, _bcb)
			switch _cad.ColorSpace.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				_fcf, _abde = _bg.ImageHandling.NewGrayImageFromGoImage(_fgd)
				if _abde != nil {
					return _abde
				}
			default:
				_fcf, _abde = _bg.ImageHandling.NewImageFromGoImage(_fgd)
				if _abde != nil {
					return _abde
				}
			}
		}
		_gdgc = &cachedImage{_ffg: _fcf, _fefd: _cad.ColorSpace}
		_dec._bag[_deb] = _gdgc
	}
	_gcd := _gdgc._ffg
	_gad := _gdgc._fefd
	_dafa, _dgde := _gad.ImageToRGB(*_gcd)
	if _dgde != nil {
		return _dgde
	}
	_b.Log.Debug("@\u0044\u006f\u0020\u0043\u0054\u004d\u003a\u0020\u0025\u0073", _efg.CTM.String())
	_fee := ImageMark{Image: &_dafa, Width: _efg.CTM.ScalingFactorX(), Height: _efg.CTM.ScalingFactorY(), Angle: _efg.CTM.Angle()}
	_fee.X, _fee.Y = _efg.CTM.Translation()
	_dec._aaa = append(_dec._aaa, _fee)
	_dec._cdg++
	return nil
}
func _cbeeg(_cfba []*textLine, _gggd, _eggcf float64) []*textLine {
	var _fadg []*textLine
	for _, _ffgae := range _cfba {
		if _gggd == -1 {
			if _ffgae._ggfc > _eggcf {
				_fadg = append(_fadg, _ffgae)
			}
		} else {
			if _ffgae._ggfc > _eggcf && _ffgae._ggfc < _gggd {
				_fadg = append(_fadg, _ffgae)
			}
		}
	}
	return _fadg
}
func _abgf(_cdee []structElement, _afgg map[int][]*textLine, _eeef _gbc.PdfObject) []*list {
	_faac := []*list{}
	for _, _ggdc := range _cdee {
		_dfeg := _ggdc._fbcba
		_ddbc := int(_ggdc._agec)
		_bgbdd := _ggdc._adfb
		_cfff := []*textLine{}
		_dbbff := []*list{}
		_dadf := _ggdc._dbdae
		_fgbc, _fffd := (_dadf.(*_gbc.PdfObjectReference))
		if !_fffd {
			_b.Log.Debug("\u0066\u0061\u0069l\u0065\u0064\u0020\u006f\u0074\u0020\u0063\u0061\u0073\u0074\u0020\u0074\u006f\u0020\u002a\u0063\u006f\u0072\u0065\u002e\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _ddbc != -1 && _fgbc != nil {
			if _acde, _eedag := _afgg[_ddbc]; _eedag {
				if _cefdf, _gbfg := _eeef.(*_gbc.PdfIndirectObject); _gbfg {
					_edecf := _cefdf.PdfObjectReference
					if _af.DeepEqual(*_fgbc, _edecf) {
						_cfff = _acde
					}
				}
			}
		}
		if _dfeg != nil {
			_dbbff = _abgf(_dfeg, _afgg, _eeef)
		}
		_aegd := _egfa(_cfff, _bgbdd, _dbbff)
		_faac = append(_faac, _aegd)
	}
	return _faac
}
func (_gefa *PageText) computeViews() {
	_dfbe := _gefa.getParagraphs()
	_ggg := new(_aa.Buffer)
	_dfbe.writeText(_ggg)
	_gefa._bda = _ggg.String()
	_gefa._dcge = _dfbe.toTextMarks()
	_gefa._eedcg = _dfbe.tables()
	if _eeeg {
		_b.Log.Info("\u0063\u006f\u006dpu\u0074\u0065\u0056\u0069\u0065\u0077\u0073\u003a\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064", len(_gefa._eedcg))
	}
}
func _cdbae(_addag, _cdea _f.Image) _f.Image {
	_cfeef, _egdbc := _cdea.Bounds().Size(), _addag.Bounds().Size()
	_ebabe, _abegd := _cfeef.X, _cfeef.Y
	if _egdbc.X > _ebabe {
		_ebabe = _egdbc.X
	}
	if _egdbc.Y > _abegd {
		_abegd = _egdbc.Y
	}
	_fbfac := _f.Rect(0, 0, _ebabe, _abegd)
	if _cfeef.X != _ebabe || _cfeef.Y != _abegd {
		_cbceg := _f.NewRGBA(_fbfac)
		_ca.BiLinear.Scale(_cbceg, _fbfac, _addag, _cdea.Bounds(), _ca.Over, nil)
		_cdea = _cbceg
	}
	if _egdbc.X != _ebabe || _egdbc.Y != _abegd {
		_afdca := _f.NewRGBA(_fbfac)
		_ca.BiLinear.Scale(_afdca, _fbfac, _addag, _addag.Bounds(), _ca.Over, nil)
		_addag = _afdca
	}
	_acfgg := _f.NewRGBA(_fbfac)
	_ca.DrawMask(_acfgg, _fbfac, _addag, _f.Point{}, _cdea, _f.Point{}, _ca.Over)
	return _acfgg
}
func (_dfcc rulingList) toGrids() []rulingList {
	if _gfab {
		_b.Log.Info("t\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _dfcc)
	}
	_ddec := _dfcc.intersections()
	if _gfab {
		_b.Log.Info("\u0074\u006f\u0047r\u0069\u0064\u0073\u003a \u0076\u0065\u0063\u0073\u003d\u0025\u0064 \u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003d\u0025\u0064\u0020", len(_dfcc), len(_ddec))
		for _, _acdbc := range _eged(_ddec) {
			_d.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _acdbc, _ddec[_acdbc])
		}
	}
	_ggcfaa := make(map[int]intSet, len(_dfcc))
	for _cgfff := range _dfcc {
		_fbaf := _dfcc.connections(_ddec, _cgfff)
		if len(_fbaf) > 0 {
			_ggcfaa[_cgfff] = _fbaf
		}
	}
	if _gfab {
		_b.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0063\u006fn\u006e\u0065\u0063\u0074s=\u0025\u0064", len(_ggcfaa))
		for _, _cdag := range _eged(_ggcfaa) {
			_d.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _cdag, _ggcfaa[_cdag])
		}
	}
	_ffcfa := _dbbdd(len(_dfcc), func(_fdae, _febcg int) bool {
		_dgccg, _ebfdg := len(_ggcfaa[_fdae]), len(_ggcfaa[_febcg])
		if _dgccg != _ebfdg {
			return _dgccg > _ebfdg
		}
		return _dfcc.comp(_fdae, _febcg)
	})
	if _gfab {
		_b.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u006f\u0072d\u0065\u0072\u0069\u006eg=\u0025\u0076", _ffcfa)
	}
	_gagg := [][]int{{_ffcfa[0]}}
_ffgad:
	for _, _ffcff := range _ffcfa[1:] {
		for _caae, _edaac := range _gagg {
			for _, _fdbc := range _edaac {
				if _ggcfaa[_fdbc].has(_ffcff) {
					_gagg[_caae] = append(_edaac, _ffcff)
					continue _ffgad
				}
			}
		}
		_gagg = append(_gagg, []int{_ffcff})
	}
	if _gfab {
		_b.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020i\u0067\u0072i\u0064\u0073\u003d\u0025\u0076", _gagg)
	}
	_gd.SliceStable(_gagg, func(_gffd, _bgggc int) bool { return len(_gagg[_gffd]) > len(_gagg[_bgggc]) })
	for _, _geaed := range _gagg {
		_gd.Slice(_geaed, func(_feafa, _debce int) bool { return _dfcc.comp(_geaed[_feafa], _geaed[_debce]) })
	}
	_acadb := make([]rulingList, len(_gagg))
	for _ddde, _accg := range _gagg {
		_cdcb := make(rulingList, len(_accg))
		for _fdcga, _becc := range _accg {
			_cdcb[_fdcga] = _dfcc[_becc]
		}
		_acadb[_ddde] = _cdcb
	}
	if _gfab {
		_b.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020g\u0072\u0069d\u0073\u003d\u0025\u002b\u0076", _acadb)
	}
	var _cgad []rulingList
	for _, _afba := range _acadb {
		if _fege, _cecec := _afba.isActualGrid(); _cecec {
			_afba = _fege
			_afba = _afba.snapToGroups()
			_cgad = append(_cgad, _afba)
		}
	}
	if _gfab {
		_fgba("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0061\u0063t\u0075\u0061\u006c\u0047ri\u0064\u0073", _cgad)
		_b.Log.Info("\u0074\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0067\u0072\u0069\u0064\u0073\u003d%\u0064 \u0061\u0063\u0074\u0075\u0061\u006c\u0047\u0072\u0069\u0064\u0073\u003d\u0025\u0064", len(_acadb), len(_cgad))
	}
	return _cgad
}

var _aafg string = "\u005e\u005b\u0061\u002d\u007a\u0041\u002dZ\u005d\u0028\u005c)\u007c\u005c\u002e)\u007c\u005e[\u005c\u0064\u005d\u002b\u0028\u005c)\u007c\\.\u0029\u007c\u005e\u005c\u0028\u005b\u0061\u002d\u007a\u0041\u002d\u005a\u005d\u005c\u0029\u007c\u005e\u005c\u0028\u005b\u005c\u0064\u005d\u002b\u005c\u0029"

func (_bcage *textMark) inDiacriticArea(_aadf *textMark) bool {
	_dcab := _bcage.Llx - _aadf.Llx
	_bgeb := _bcage.Urx - _aadf.Urx
	_geeb := _bcage.Lly - _aadf.Lly
	return _bb.Abs(_dcab+_bgeb) < _bcage.Width()*_dgee && _bb.Abs(_geeb) < _bcage.Height()*_dgee
}
func _dbge(_cfca []*textLine) {
	_gd.Slice(_cfca, func(_feda, _dfbde int) bool {
		_baefa, _ffcd := _cfca[_feda], _cfca[_dfbde]
		return _baefa._ggfc < _ffcd._ggfc
	})
}
func (_gacgg *textPara) writeCellText(_bdee _e.Writer) {
	for _aeag, _eadbb := range _gacgg._gdc {
		_cabd := _eadbb.text()
		_bcga := _degfc && _eadbb.endsInHyphen() && _aeag != len(_gacgg._gdc)-1
		if _bcga {
			_cabd = _ggfa(_cabd)
		}
		_bdee.Write([]byte(_cabd))
		if !(_bcga || _aeag == len(_gacgg._gdc)-1) {
			_bdee.Write([]byte(_afbff(_eadbb._ggfc, _gacgg._gdc[_aeag+1]._ggfc)))
		}
	}
}
func (_acgbd *textLine) text() string {
	var _dabg []string
	for _, _gcbc := range _acgbd._edee {
		if _gcbc._ggdce {
			_dabg = append(_dabg, "\u0020")
		}
		_dabg = append(_dabg, _gcbc._edac)
	}
	_bcfg := _ge.Join(_dabg, "")
	_bdda := _efc([]rune(_bcfg))
	return _bdda._feb
}
func (_eedf *textLine) bbox() _bg.PdfRectangle { return _eedf.PdfRectangle }
func _fbfb(_gbcc []Font, _affc string) bool {
	for _, _eea := range _gbcc {
		if _eea.FontName == _affc {
			return true
		}
	}
	return false
}
func (_cdae paraList) extractTables(_cgcb []gridTiling) paraList {
	if _eeeg {
		_b.Log.Debug("\u0065\u0078\u0074r\u0061\u0063\u0074\u0054\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_cdae))
	}
	if len(_cdae) < _cbea {
		return _cdae
	}
	_facc := _cdae.findTables(_cgcb)
	if _eeeg {
		_b.Log.Info("c\u006f\u006d\u0062\u0069\u006e\u0065d\u0020\u0074\u0061\u0062\u006c\u0065s\u0020\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_facc))
		for _cecg, _bagcd := range _facc {
			_bagcd.log(_d.Sprintf("c\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u0020\u0025\u0064", _cecg))
		}
	}
	return _cdae.applyTables(_facc)
}

// Tables returns the tables extracted from the page.
func (_gacg PageText) Tables() []TextTable {
	if _eeeg {
		_b.Log.Info("\u0054\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_gacg._eedcg))
	}
	return _gacg._eedcg
}
func (_aadgc rectRuling) asRuling() (*ruling, bool) {
	_fcgbd := ruling{_gcafd: _aadgc._acad, Color: _aadgc.Color, _cebe: _aead}
	switch _aadgc._acad {
	case _cege:
		_fcgbd._gbaff = 0.5 * (_aadgc.Llx + _aadgc.Urx)
		_fcgbd._fabe = _aadgc.Lly
		_fcgbd._ccba = _aadgc.Ury
		_bbge, _eddcg := _aadgc.checkWidth(_aadgc.Llx, _aadgc.Urx)
		if !_eddcg {
			if _bfgg {
				_b.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067V\u0065\u0072\u0074\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _aadgc)
			}
			return nil, false
		}
		_fcgbd._gcec = _bbge
	case _bfcbd:
		_fcgbd._gbaff = 0.5 * (_aadgc.Lly + _aadgc.Ury)
		_fcgbd._fabe = _aadgc.Llx
		_fcgbd._ccba = _aadgc.Urx
		_fgagg, _ccgaa := _aadgc.checkWidth(_aadgc.Lly, _aadgc.Ury)
		if !_ccgaa {
			if _bfgg {
				_b.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067H\u006f\u0072\u007a\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _aadgc)
			}
			return nil, false
		}
		_fcgbd._gcec = _fgagg
	default:
		_b.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _aadgc._acad)
		return nil, false
	}
	return &_fcgbd, true
}
func _gdgea(_cccgg []compositeCell) []float64 {
	var _ebdge []*textLine
	_ggcee := 0
	for _, _ebec := range _cccgg {
		_ggcee += len(_ebec.paraList)
		_ebdge = append(_ebdge, _ebec.lines()...)
	}
	_gd.Slice(_ebdge, func(_aggdc, _eeagf int) bool {
		_eefef, _bfff := _ebdge[_aggdc], _ebdge[_eeagf]
		_fgcg, _deaad := _eefef._ggfc, _bfff._ggfc
		if !_gacaa(_fgcg - _deaad) {
			return _fgcg < _deaad
		}
		return _eefef.Llx < _bfff.Llx
	})
	if _eeeg {
		_d.Printf("\u0020\u0020\u0020 r\u006f\u0077\u0042\u006f\u0072\u0064\u0065\u0072\u0073:\u0020%\u0064 \u0070a\u0072\u0061\u0073\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", _ggcee, len(_ebdge))
		for _cbfdc, _ffedf := range _ebdge {
			_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _cbfdc, _ffedf)
		}
	}
	var _gbfa []float64
	_aacf := _ebdge[0]
	var _ebef [][]*textLine
	_cabg := []*textLine{_aacf}
	for _bfgde, _gfda := range _ebdge[1:] {
		if _gfda.Ury < _aacf.Lly {
			_ecea := 0.5 * (_gfda.Ury + _aacf.Lly)
			if _eeeg {
				_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u003c\u0020\u0025\u0036.\u0032f\u0020\u0062\u006f\u0072\u0064\u0065\u0072\u003d\u0025\u0036\u002e\u0032\u0066\u000a"+"\u0009\u0020\u0071\u003d\u0025\u0073\u000a\u0009\u0020p\u003d\u0025\u0073\u000a", _bfgde, _gfda.Ury, _aacf.Lly, _ecea, _aacf, _gfda)
			}
			_gbfa = append(_gbfa, _ecea)
			_ebef = append(_ebef, _cabg)
			_cabg = nil
		}
		_cabg = append(_cabg, _gfda)
		if _gfda.Lly < _aacf.Lly {
			_aacf = _gfda
		}
	}
	if len(_cabg) > 0 {
		_ebef = append(_ebef, _cabg)
	}
	if _eeeg {
		_d.Printf(" \u0020\u0020\u0020\u0020\u0020\u0020 \u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073\u003d%\u0036.\u0032\u0066\u000a", _gbfa)
	}
	if _eeeg {
		_b.Log.Info("\u0072\u006f\u0077\u003d\u0025\u0064", len(_cccgg))
		for _aefea, _bfaag := range _cccgg {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aefea, _bfaag)
		}
		_b.Log.Info("\u0067r\u006f\u0075\u0070\u0073\u003d\u0025d", len(_ebef))
		for _baeaf, _ceeb := range _ebef {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0064\u000a", _baeaf, len(_ceeb))
			for _cdece, _bedcd := range _ceeb {
				_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _cdece, _bedcd)
			}
		}
	}
	_agce := true
	for _fdbbd, _aefeb := range _ebef {
		_ggcg := true
		for _cddff, _bdfbe := range _cccgg {
			if _eeeg {
				_d.Printf("\u0020\u0020\u0020\u007e\u007e\u007e\u0067\u0072\u006f\u0075\u0070\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0063\u0065\u006cl\u0020\u0025\u0064\u0020\u006ff\u0020\u0025d\u0020\u0025\u0073\u000a", _fdbbd, len(_ebef), _cddff, len(_cccgg), _bdfbe)
			}
			if !_bdfbe.hasLines(_aefeb) {
				if _eeeg {
					_d.Printf("\u0020\u0020\u0020\u0021\u0021\u0021\u0067\u0072\u006f\u0075\u0070\u0020\u0025d\u0020\u006f\u0066\u0020\u0025\u0064 \u0063\u0065\u006c\u006c\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064 \u004f\u0055\u0054\u000a", _fdbbd, len(_ebef), _cddff, len(_cccgg))
				}
				_ggcg = false
				break
			}
		}
		if !_ggcg {
			_agce = false
			break
		}
	}
	if !_agce {
		if _eeeg {
			_b.Log.Info("\u0072\u006f\u0077\u0020\u0063o\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u0020\u0064\u006f\u006e\u0027\u0074 \u0073\u0070\u0061\u006e\u0020\u0061\u006c\u006c\u0020\u0063\u0065\u006c\u006c\u0073\u0020\u0069\u006e\u0020\u0072\u006f\u0077\u002e\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
		}
		_gbfa = nil
	}
	if _eeeg && _gbfa != nil {
		_d.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u002a\u002a*\u0072\u006f\u0077\u0043\u006f\u0072\u0072i\u0064\u006f\u0072\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u000a", _gbfa)
	}
	return _gbfa
}
func _dbadc(_fdcfc map[float64]map[float64]gridTile) []float64 {
	_ebadf := make([]float64, 0, len(_fdcfc))
	_fgabd := make(map[float64]struct{}, len(_fdcfc))
	for _, _fcdb := range _fdcfc {
		for _fecce := range _fcdb {
			if _, _cdade := _fgabd[_fecce]; _cdade {
				continue
			}
			_ebadf = append(_ebadf, _fecce)
			_fgabd[_fecce] = struct{}{}
		}
	}
	_gd.Float64s(_ebadf)
	return _ebadf
}
func (_fcfee *textTable) put(_geebg, _fagd int, _afcgb *textPara) {
	_fcfee._bcec[_caabd(_geebg, _fagd)] = _afcgb
}
func (_egffb *textWord) absorb(_abbce *textWord) {
	_egffb.PdfRectangle = _egbb(_egffb.PdfRectangle, _abbce.PdfRectangle)
	_egffb._ecdf = append(_egffb._ecdf, _abbce._ecdf...)
}
func _fgebf(_bgddb, _egdd, _bgcb float64) rulingKind {
	if _bgddb >= _bgcb && _fcge(_egdd, _bgddb) {
		return _bfcbd
	}
	if _egdd >= _bgcb && _fcge(_bgddb, _egdd) {
		return _cege
	}
	return _bgeag
}

type textObject struct {
	_eadb *Extractor
	_fgf  *_bg.PdfPageResources
	_dba  _aed.GraphicsState
	_egf  *textState
	_fbad *stateStack
	_gbcf _ef.Matrix
	_gbe  _ef.Matrix
	_gaa  []*textMark
	_gea  bool
}

func (_abff *wordBag) depthBand(_gdd, _adcdf float64) []int {
	if len(_abff._cccf) == 0 {
		return nil
	}
	return _abff.depthRange(_abff.getDepthIdx(_gdd), _abff.getDepthIdx(_adcdf))
}
func (_fcdf paraList) topoOrder() []int {
	if _agba {
		_b.Log.Info("\u0074\u006f\u0070\u006f\u004f\u0072\u0064\u0065\u0072\u003a")
	}
	_fgbcc := len(_fcdf)
	_ffgab := make([]bool, _fgbcc)
	_fffdb := make([]int, 0, _fgbcc)
	_cfcd := _fcdf.llyOrdering()
	var _egcd func(_eeba int)
	_egcd = func(_dgaag int) {
		_ffgab[_dgaag] = true
		for _feefc := 0; _feefc < _fgbcc; _feefc++ {
			if !_ffgab[_feefc] {
				if _fcdf.readBefore(_cfcd, _dgaag, _feefc) {
					_egcd(_feefc)
				}
			}
		}
		_fffdb = append(_fffdb, _dgaag)
	}
	for _baafb := 0; _baafb < _fgbcc; _baafb++ {
		if !_ffgab[_baafb] {
			_egcd(_baafb)
		}
	}
	return _aaaa(_fffdb)
}
func (_bgbe *wordBag) applyRemovals(_fgfc map[int]map[*textWord]struct{}) {
	for _eeff, _acga := range _fgfc {
		if len(_acga) == 0 {
			continue
		}
		_bdgd := _bgbe._cccf[_eeff]
		_beec := len(_bdgd) - len(_acga)
		if _beec == 0 {
			delete(_bgbe._cccf, _eeff)
			continue
		}
		_bad := make([]*textWord, _beec)
		_ccab := 0
		for _, _egag := range _bdgd {
			if _, _cadg := _acga[_egag]; !_cadg {
				_bad[_ccab] = _egag
				_ccab++
			}
		}
		_bgbe._cccf[_eeff] = _bad
	}
}
func (_agbe *compositeCell) updateBBox() {
	for _, _fgec := range _agbe.paraList {
		_agbe.PdfRectangle = _egbb(_agbe.PdfRectangle, _fgec.PdfRectangle)
	}
}
func (_egg *shapesState) quadraticTo(_aacb, _gfec, _cgg, _begbe float64) {
	if _fbfd {
		_b.Log.Info("\u0071\u0075\u0061d\u0072\u0061\u0074\u0069\u0063\u0054\u006f\u003a")
	}
	_egg.addPoint(_cgg, _begbe)
}
func (_bbaed compositeCell) parasBBox() (paraList, _bg.PdfRectangle) {
	return _bbaed.paraList, _bbaed.PdfRectangle
}
func (_acgfe paraList) findTableGrid(_ageb gridTiling) (*textTable, map[*textPara]struct{}) {
	_fdced := len(_ageb._cdfg)
	_gaadd := len(_ageb._fgcd)
	_fafeb := textTable{_gfagc: true, _ddcega: _fdced, _egbbf: _gaadd, _bcec: make(map[uint64]*textPara, _fdced*_gaadd), _fccfa: make(map[uint64]compositeCell, _fdced*_gaadd)}
	_fafeb.PdfRectangle = _ageb.PdfRectangle
	_fffda := make(map[*textPara]struct{})
	_dbcaa := int((1.0 - _ddcb) * float64(_fdced*_gaadd))
	_eadcd := 0
	if _degg {
		_b.Log.Info("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0064\u0020\u0078\u0020%\u0064", _fdced, _gaadd)
	}
	for _bceea, _degdg := range _ageb._fgcd {
		_abaf, _dfcfb := _ageb._egff[_degdg]
		if !_dfcfb {
			continue
		}
		for _gegca, _afeab := range _ageb._cdfg {
			_cdbbfc, _gbafa := _abaf[_afeab]
			if !_gbafa {
				continue
			}
			_eegad := _acgfe.inTile(_cdbbfc)
			if len(_eegad) == 0 {
				_eadcd++
				if _eadcd > _dbcaa {
					if _degg {
						_b.Log.Info("\u0021\u006e\u0075m\u0045\u006d\u0070\u0074\u0079\u003d\u0025\u0064", _eadcd)
					}
					return nil, nil
				}
			} else {
				_fafeb.putComposite(_gegca, _bceea, _eegad, _cdbbfc.PdfRectangle)
				for _, _bceb := range _eegad {
					_fffda[_bceb] = struct{}{}
				}
			}
		}
	}
	_bacc := 0
	for _dcef := 0; _dcef < _fdced; _dcef++ {
		_adfacd := _fafeb.get(_dcef, 0)
		if _adfacd == nil || !_adfacd._ecee {
			_bacc++
		}
	}
	if _bacc == 0 {
		if _degg {
			_b.Log.Info("\u0021\u006e\u0075m\u0048\u0065\u0061\u0064\u0065\u0072\u003d\u0030")
		}
		return nil, nil
	}
	_eagb := _fafeb.reduceTiling(_ageb, _dbae)
	_eagb = _eagb.subdivide()
	return _eagb, _fffda
}
func (_bca *stateStack) pop() *textState {
	if _bca.empty() {
		return nil
	}
	_dafb := *(*_bca)[len(*_bca)-1]
	*_bca = (*_bca)[:len(*_bca)-1]
	return &_dafb
}
func (_fdada *ruling) alignsSec(_ffdaa *ruling) bool {
	const _ccgg = _egba + 1.0
	return _fdada._fabe-_ccgg <= _ffdaa._ccba && _ffdaa._fabe-_ccgg <= _fdada._ccba
}
func _gbeg(_fggfd float64) float64 { return _bagc * _bb.Round(_fggfd/_bagc) }

// Text returns the text content of the `bulletLists`.
func (_adac *lists) Text() string {
	_dgbf := &_ge.Builder{}
	for _, _bgbd := range *_adac {
		_aeff := _bgbd.Text()
		_dgbf.WriteString(_aeff)
	}
	return _dgbf.String()
}

const _ggc = 20

func _afbff(_aegb, _cadc float64) string {
	_gggdg := !_gacaa(_aegb - _cadc)
	if _gggdg {
		return "\u000a"
	}
	return "\u0020"
}
func (_geb *textObject) setWordSpacing(_ccgdb float64) {
	if _geb == nil {
		return
	}
	_geb._egf._eaed = _ccgdb
}

// PageText represents the layout of text on a device page.
type PageText struct {
	_faba  []*textMark
	_bda   string
	_dcge  []TextMark
	_eedcg []TextTable
	_egdec _bg.PdfRectangle
	_fce   []pathSection
	_bdfb  []pathSection
	_gge   *_gbc.PdfObject
	_ecfd  _gbc.PdfObject
	_dbcd  *_aed.ContentStreamOperations
	_gfee  PageTextOptions
}

// String returns a description of `p`.
func (_bbaga *textPara) String() string {
	if _bbaga._ecee {
		return _d.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u005b\u0045\u004d\u0050\u0054\u0059\u005d", _bbaga.PdfRectangle)
	}
	_bcgf := ""
	if _bbaga._cece != nil {
		_bcgf = _d.Sprintf("\u005b\u0025\u0064\u0078\u0025\u0064\u005d\u0020", _bbaga._cece._ddcega, _bbaga._cece._egbbf)
	}
	return _d.Sprintf("\u0025\u0036\u002e\u0032f \u0025\u0073\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0025\u0071", _bbaga.PdfRectangle, _bcgf, len(_bbaga._gdc), _cfbg(_bbaga.text(), 50))
}
func (_ebfgd *ruling) equals(_cabdb *ruling) bool {
	return _ebfgd._gcafd == _cabdb._gcafd && _gddc(_ebfgd._gbaff, _cabdb._gbaff) && _gddc(_ebfgd._fabe, _cabdb._fabe) && _gddc(_ebfgd._ccba, _cabdb._ccba)
}

// PageImages represents extracted images on a PDF page with spatial information:
// display position and size.
type PageImages struct{ Images []ImageMark }

func (_bba *imageExtractContext) extractContentStreamImages(_gbf string, _fadc *_bg.PdfPageResources) error {
	_acf := _aed.NewContentStreamParser(_gbf)
	_gc, _cc := _acf.Parse()
	if _cc != nil {
		return _cc
	}
	if _bba._bag == nil {
		_bba._bag = map[*_gbc.PdfObjectStream]*cachedImage{}
	}
	if _bba._gfgg == nil {
		_bba._gfgg = &ImageExtractOptions{}
	}
	_bac := _aed.NewContentStreamProcessor(*_gc)
	_bac.AddHandler(_aed.HandlerConditionEnumAllOperands, "", _bba.processOperand)
	return _bac.Process(_fadc)
}
func _cdaae(_baba []TextMark, _egfd *int) []TextMark {
	_gfadc := _baba[len(_baba)-1]
	_eeec := []rune(_gfadc.Text)
	if len(_eeec) == 1 {
		_baba = _baba[:len(_baba)-1]
		_dcga := _baba[len(_baba)-1]
		*_egfd = _dcga.Offset + len(_dcga.Text)
	} else {
		_edgac := _ggfa(_gfadc.Text)
		*_egfd += len(_edgac) - len(_gfadc.Text)
		_gfadc.Text = _edgac
	}
	return _baba
}
func _faebg(_bgffbg map[int][]float64) string {
	_dccgg := _bgffa(_bgffbg)
	_cdfdd := make([]string, len(_bgffbg))
	for _befd, _decg := range _dccgg {
		_cdfdd[_befd] = _d.Sprintf("\u0025\u0064\u003a\u0020\u0025\u002e\u0032\u0066", _decg, _bgffbg[_decg])
	}
	return _d.Sprintf("\u007b\u0025\u0073\u007d", _ge.Join(_cdfdd, "\u002c\u0020"))
}

// Text returns the extracted page text.
func (_beef PageText) Text() string { return _beef._bda }
func _gfcea(_gffb []*textLine, _cbeg map[float64][]*textLine) []*list {
	_aaea := _ddceg(_cbeg)
	_fagbc := []*list{}
	if len(_aaea) == 0 {
		return _fagbc
	}
	_cadfa := _aaea[0]
	_fbac := 1
	_gceg := _cbeg[_cadfa]
	for _becg, _ggca := range _gceg {
		var _cfce float64
		_eggc := []*list{}
		_dfff := _ggca._ggfc
		_ccdg := -1.0
		if _becg < len(_gceg)-1 {
			_ccdg = _gceg[_becg+1]._ggfc
		}
		if _fbac < len(_aaea) {
			_eggc = _agef(_gffb, _cbeg, _aaea, _fbac, _dfff, _ccdg)
		}
		_cfce = _ccdg
		if len(_eggc) > 0 {
			_eefff := _eggc[0]
			if len(_eefff._eadba) > 0 {
				_cfce = _eefff._eadba[0]._ggfc
			}
		}
		_eaec := []*textLine{_ggca}
		_agdf := _ebfgb(_ggca, _gffb, _aaea, _dfff, _cfce)
		_eaec = append(_eaec, _agdf...)
		_efcbf := _egfa(_eaec, "\u0062\u0075\u006c\u006c\u0065\u0074", _eggc)
		_efcbf._fccf = _cgef(_eaec, "")
		_fagbc = append(_fagbc, _efcbf)
	}
	return _fagbc
}
func _cdaab(_abbae, _fccfac int) int {
	if _abbae < _fccfac {
		return _abbae
	}
	return _fccfac
}

// ExtractTextWithStats works like ExtractText but returns the number of characters in the output
// (`numChars`) and the number of characters that were not decoded (`numMisses`).
func (_cea *Extractor) ExtractTextWithStats() (_aad string, _aef int, _eega int, _affg error) {
	_dgf, _aef, _eega, _affg := _cea.ExtractPageText()
	if _affg != nil {
		return "", _aef, _eega, _affg
	}
	return _dgf.Text(), _aef, _eega, nil
}
func _agef(_gaca []*textLine, _fdef map[float64][]*textLine, _cgde []float64, _ggab int, _geff, _geac float64) []*list {
	_cbag := []*list{}
	_fgeg := _ggab
	_ggab = _ggab + 1
	_egbc := _cgde[_fgeg]
	_bgbef := _fdef[_egbc]
	_fac := _cbeeg(_bgbef, _geac, _geff)
	for _adge, _dcca := range _fac {
		var _daff float64
		_gfca := []*list{}
		_dbab := _dcca._ggfc
		_afec := _geac
		if _adge < len(_fac)-1 {
			_afec = _fac[_adge+1]._ggfc
		}
		if _ggab < len(_cgde) {
			_gfca = _agef(_gaca, _fdef, _cgde, _ggab, _dbab, _afec)
		}
		_daff = _afec
		if len(_gfca) > 0 {
			_efeb := _gfca[0]
			if len(_efeb._eadba) > 0 {
				_daff = _efeb._eadba[0]._ggfc
			}
		}
		_eefcgg := []*textLine{_dcca}
		_dffc := _ebfgb(_dcca, _gaca, _cgde, _dbab, _daff)
		_eefcgg = append(_eefcgg, _dffc...)
		_dbbc := _egfa(_eefcgg, "\u0062\u0075\u006c\u006c\u0065\u0074", _gfca)
		_dbbc._fccf = _cgef(_eefcgg, "")
		_cbag = append(_cbag, _dbbc)
	}
	return _cbag
}
func _cccae(_bceg _bg.PdfRectangle) rulingKind {
	_ddbg := _bceg.Width()
	_ccdd := _bceg.Height()
	if _ddbg > _ccdd {
		if _ddbg >= _feef {
			return _bfcbd
		}
	} else {
		if _ccdd >= _feef {
			return _cege
		}
	}
	return _bgeag
}
func _cfg(_dca _ef.Matrix) _ef.Point {
	_eegbf, _bgaf := _dca.Translation()
	return _ef.Point{X: _eegbf, Y: _bgaf}
}
func _fgba(_acfd string, _gaeef []rulingList) {
	_b.Log.Info("\u0024\u0024 \u0025\u0064\u0020g\u0072\u0069\u0064\u0073\u0020\u002d\u0020\u0025\u0073", len(_gaeef), _acfd)
	for _cgce, _fggg := range _gaeef {
		_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cgce, _fggg.String())
	}
}
func (_feab *textPara) bbox() _bg.PdfRectangle { return _feab.PdfRectangle }
func (_gaaag intSet) add(_gdcd int)            { _gaaag[_gdcd] = struct{}{} }
func _bafag(_gdfdg, _bcfgg _ef.Point) rulingKind {
	_gaceg := _bb.Abs(_gdfdg.X - _bcfgg.X)
	_ababc := _bb.Abs(_gdfdg.Y - _bcfgg.Y)
	return _fgebf(_gaceg, _ababc, _feef)
}
func (_eagcf *subpath) add(_debcg ..._ef.Point) { _eagcf._agea = append(_eagcf._agea, _debcg...) }
func (_ggde *wordBag) empty(_eabb int) bool     { _, _feed := _ggde._cccf[_eabb]; return !_feed }
func (_ecaa *PageFonts) extractPageResourcesToFont(_dga *_bg.PdfPageResources) error {
	if _dga.Font == nil {
		return _dg.New(_edfe)
	}
	_cef, _eac := _gbc.GetDict(_dga.Font)
	if !_eac {
		return _dg.New(_baf)
	}
	for _, _fgac := range _cef.Keys() {
		var (
			_cgf = true
			_cgc []byte
			_fgb string
		)
		_aeb, _egc := _dga.GetFontByName(_fgac)
		if !_egc {
			return _dg.New(_febe)
		}
		_eaf, _bef := _bg.NewPdfFontFromPdfObject(_aeb)
		if _bef != nil {
			return _bef
		}
		_fad := _eaf.FontDescriptor()
		_gfe := _eaf.FontDescriptor().FontName.String()
		_gfc := _eaf.Subtype()
		if _fbfb(_ecaa.Fonts, _gfe) {
			continue
		}
		if len(_eaf.ToUnicode()) == 0 {
			_cgf = false
		}
		if _fad.FontFile != nil {
			if _eefb, _ebcd := _gbc.GetStream(_fad.FontFile); _ebcd {
				_cgc, _bef = _gbc.DecodeStream(_eefb)
				if _bef != nil {
					return _bef
				}
				_fgb = _gfe + "\u002e\u0070\u0066\u0062"
			}
		} else if _fad.FontFile2 != nil {
			if _fef, _eae := _gbc.GetStream(_fad.FontFile2); _eae {
				_cgc, _bef = _gbc.DecodeStream(_fef)
				if _bef != nil {
					return _bef
				}
				_fgb = _gfe + "\u002e\u0074\u0074\u0066"
			}
		} else if _fad.FontFile3 != nil {
			if _babc, _efa := _gbc.GetStream(_fad.FontFile3); _efa {
				_cgc, _bef = _gbc.DecodeStream(_babc)
				if _bef != nil {
					return _bef
				}
				_fgb = _gfe + "\u002e\u0063\u0066\u0066"
			}
		}
		if len(_fgb) < 1 {
			_b.Log.Debug(_dbdg)
		}
		_bf := Font{FontName: _gfe, PdfFont: _eaf, IsCID: _eaf.IsCID(), IsSimple: _eaf.IsSimple(), ToUnicode: _cgf, FontType: _gfc, FontData: _cgc, FontFileName: _fgb, FontDescriptor: _fad}
		_ecaa.Fonts = append(_ecaa.Fonts, _bf)
	}
	return nil
}
func (_fbdfg paraList) merge() *textPara {
	_b.Log.Trace("\u006d\u0065\u0072\u0067\u0065:\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_fbdfg))
	if len(_fbdfg) == 0 {
		return nil
	}
	_fbdfg.sortReadingOrder()
	_egcb := _fbdfg[0].PdfRectangle
	_eaaa := _fbdfg[0]._gdc
	for _, _gcfa := range _fbdfg[1:] {
		_egcb = _egbb(_egcb, _gcfa.PdfRectangle)
		_eaaa = append(_eaaa, _gcfa._gdc...)
	}
	return _ebeeg(_egcb, _eaaa)
}
func _ebbf(_edda map[float64]gridTile) []float64 {
	_bafdc := make([]float64, 0, len(_edda))
	for _adgeg := range _edda {
		_bafdc = append(_bafdc, _adgeg)
	}
	_gd.Float64s(_bafdc)
	return _bafdc
}
func (_gecg *textLine) pullWord(_cgeeb *wordBag, _fdcd *textWord, _acda int) {
	_gecg.appendWord(_fdcd)
	_cgeeb.removeWord(_fdcd, _acda)
}

// String returns a string describing `ma`.
func (_cbce TextMarkArray) String() string {
	_acgf := len(_cbce._egbe)
	if _acgf == 0 {
		return "\u0045\u004d\u0050T\u0059"
	}
	_adfg := _cbce._egbe[0]
	_dafc := _cbce._egbe[_acgf-1]
	return _d.Sprintf("\u007b\u0054\u0045\u0058\u0054\u004d\u0041\u0052K\u0041\u0052\u0052AY\u003a\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025s\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d%\u0073\u007d", _acgf, _adfg, _dafc)
}

type cachedImage struct {
	_ffg  *_bg.Image
	_fefd _bg.PdfColorspace
}

func (_dbfeb rulingList) sortStrict() {
	_gd.Slice(_dbfeb, func(_ebbbc, _bagbc int) bool {
		_gfadf, _fbfg := _dbfeb[_ebbbc], _dbfeb[_bagbc]
		_fcggc, _gaecd := _gfadf._gcafd, _fbfg._gcafd
		if _fcggc != _gaecd {
			return _fcggc > _gaecd
		}
		_dgfde, _aedb := _gfadf._gbaff, _fbfg._gbaff
		if !_gacaa(_dgfde - _aedb) {
			return _dgfde < _aedb
		}
		_dgfde, _aedb = _gfadf._fabe, _fbfg._fabe
		if _dgfde != _aedb {
			return _dgfde < _aedb
		}
		return _gfadf._ccba < _fbfg._ccba
	})
}

const (
	_bgeag rulingKind = iota
	_bfcbd
	_cege
)

func (_acg *textObject) setFont(_ddba string, _cdd float64) error {
	if _acg == nil {
		return nil
	}
	_acg._egf._ebab = _cdd
	_ccga, _adfa := _acg.getFont(_ddba)
	if _adfa != nil {
		return _adfa
	}
	_acg._egf._aede = _ccga
	return nil
}
func (_gace *shapesState) addPoint(_bbee, _ddda float64) {
	_accdg := _gace.establishSubpath()
	_fdgg := _gace.devicePoint(_bbee, _ddda)
	if _accdg == nil {
		_gace._aaba = true
		_gace._gbgg = _fdgg
	} else {
		_accdg.add(_fdgg)
	}
}
func (_faec rulingList) primaries() []float64 {
	_bddd := make(map[float64]struct{}, len(_faec))
	for _, _fbdb := range _faec {
		_bddd[_fbdb._gbaff] = struct{}{}
	}
	_gdecg := make([]float64, len(_bddd))
	_gcda := 0
	for _gfbdf := range _bddd {
		_gdecg[_gcda] = _gfbdf
		_gcda++
	}
	_gd.Float64s(_gdecg)
	return _gdecg
}
func _caabd(_fbag, _dafdga int) uint64 { return uint64(_fbag)*0x1000000 + uint64(_dafdga) }

// RangeOffset returns the TextMarks in `ma` that overlap text[start:end] in the extracted text.
// These are tm: `start` <= tm.Offset + len(tm.Text) && tm.Offset < `end` where
// `start` and `end` are offsets in the extracted text.
// NOTE: TextMarks can contain multiple characters. e.g. "ffi" for the ﬃ ligature so the first and
// last elements of the returned TextMarkArray may only partially overlap text[start:end].
func (_caede *TextMarkArray) RangeOffset(start, end int) (*TextMarkArray, error) {
	if _caede == nil {
		return nil, _dg.New("\u006da\u003d\u003d\u006e\u0069\u006c")
	}
	if end < start {
		return nil, _d.Errorf("\u0065\u006e\u0064\u0020\u003c\u0020\u0073\u0074\u0061\u0072\u0074\u002e\u0020\u0052\u0061n\u0067\u0065\u004f\u0066\u0066\u0073\u0065\u0074\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020\u0073\u0074\u0061\u0072t=\u0025\u0064\u0020\u0065\u006e\u0064\u003d\u0025\u0064\u0020", start, end)
	}
	_ecef := len(_caede._egbe)
	if _ecef == 0 {
		return _caede, nil
	}
	if start < _caede._egbe[0].Offset {
		start = _caede._egbe[0].Offset
	}
	if end > _caede._egbe[_ecef-1].Offset+1 {
		end = _caede._egbe[_ecef-1].Offset + 1
	}
	_acdb := _gd.Search(_ecef, func(_fcfg int) bool { return _caede._egbe[_fcfg].Offset+len(_caede._egbe[_fcfg].Text)-1 >= start })
	if !(0 <= _acdb && _acdb < _ecef) {
		_bafdg := _d.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073\u0074\u0061\u0072\u0074\u003d%\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009 \u006c\u0061\u0073\u0074\u003d%\u0076", start, _acdb, _ecef, _caede._egbe[0], _caede._egbe[_ecef-1])
		return nil, _bafdg
	}
	_gaaa := _gd.Search(_ecef, func(_beaf int) bool { return _caede._egbe[_beaf].Offset > end-1 })
	if !(0 <= _gaaa && _gaaa < _ecef) {
		_dde := _d.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020r\u0061\u006e\u0067e\u002e\u0020\u0065n\u0064\u003d%\u0064\u0020\u0069\u0045\u006e\u0064=\u0025d \u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d\u0025\u0076", end, _gaaa, _ecef, _caede._egbe[0], _caede._egbe[_ecef-1])
		return nil, _dde
	}
	if _gaaa <= _acdb {
		return nil, _d.Errorf("\u0069\u0045\u006e\u0064\u0020\u003c=\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003a\u0020\u0073\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u0065\u006ed\u003d\u0025\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020i\u0045n\u0064\u003d\u0025\u0064", start, end, _acdb, _gaaa)
	}
	return &TextMarkArray{_egbe: _caede._egbe[_acdb:_gaaa]}, nil
}
func (_edca *shapesState) establishSubpath() *subpath {
	_efdfb, _ffcce := _edca.lastpointEstablished()
	if !_ffcce {
		_edca._dgdd = append(_edca._dgdd, _dce(_efdfb))
	}
	if len(_edca._dgdd) == 0 {
		return nil
	}
	_edca._aaba = false
	return _edca._dgdd[len(_edca._dgdd)-1]
}
func _ggfa(_adga string) string { _cdbde := []rune(_adga); return string(_cdbde[:len(_cdbde)-1]) }
func (_bagfa *ruling) gridIntersecting(_egcda *ruling) bool {
	return _gddc(_bagfa._fabe, _egcda._fabe) && _gddc(_bagfa._ccba, _egcda._ccba)
}

type list struct {
	_eadba []*textLine
	_dbaad string
	_bdfaa []*list
	_fccf  string
}

func (_aabc TextTable) getCellInfo(_dfbdc TextMark) [][]int {
	for _bgfa, _aedec := range _aabc.Cells {
		for _beee := range _aedec {
			_edce := &_aedec[_beee].Marks
			if _edce.exists(_dfbdc) {
				return [][]int{{_bgfa}, {_beee}}
			}
		}
	}
	return nil
}

// New returns an Extractor instance for extracting content from the input PDF page.
func New(page *_bg.PdfPage) (*Extractor, error) { return NewWithOptions(page, nil) }
func (_gccbdb rulingList) connections(_dcfab map[int]intSet, _ebad int) intSet {
	_gfeeb := make(intSet)
	_fafg := make(intSet)
	var _bgage func(int)
	_bgage = func(_aedc int) {
		if !_fafg.has(_aedc) {
			_fafg.add(_aedc)
			for _ege := range _gccbdb {
				if _dcfab[_ege].has(_aedc) {
					_gfeeb.add(_ege)
				}
			}
			for _ddfc := range _gccbdb {
				if _gfeeb.has(_ddfc) {
					_bgage(_ddfc)
				}
			}
		}
	}
	_bgage(_ebad)
	return _gfeeb
}

var _cfabe *_gb.Regexp = _gb.MustCompile(_fagb + "\u007c" + _aafg)

type rulingList []*ruling

func (_fgdb *textTable) toTextTable() TextTable {
	if _eeeg {
		_b.Log.Info("t\u006fT\u0065\u0078\u0074\u0054\u0061\u0062\u006c\u0065:\u0020\u0025\u0064\u0020x \u0025\u0064", _fgdb._ddcega, _fgdb._egbbf)
	}
	_cfbd := make([][]TableCell, _fgdb._egbbf)
	for _abbf := 0; _abbf < _fgdb._egbbf; _abbf++ {
		_cfbd[_abbf] = make([]TableCell, _fgdb._ddcega)
		for _dbff := 0; _dbff < _fgdb._ddcega; _dbff++ {
			_fegg := _fgdb.get(_dbff, _abbf)
			if _fegg == nil {
				continue
			}
			_dbge(_fegg._gdc)
			if _eeeg {
				_d.Printf("\u0025\u0034\u0064 \u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _dbff, _abbf, _fegg)
			}
			_cfbd[_abbf][_dbff].Text = _fegg.text()
			_feea := 0
			_cfbd[_abbf][_dbff].Marks._egbe = _fegg.toTextMarks(&_feea)
		}
	}
	_eaaaa := TextTable{W: _fgdb._ddcega, H: _fgdb._egbbf, Cells: _cfbd}
	_eaaaa.PdfRectangle = _fgdb.bbox()
	return _eaaaa
}
func (_egae *ruling) encloses(_dedf, _cbegd float64) bool {
	return _egae._fabe-_beac <= _dedf && _cbegd <= _egae._ccba+_beac
}

var (
	_geacg = map[rune]string{0x0060: "\u0300", 0x02CB: "\u0300", 0x0027: "\u0301", 0x00B4: "\u0301", 0x02B9: "\u0301", 0x02CA: "\u0301", 0x005E: "\u0302", 0x02C6: "\u0302", 0x007E: "\u0303", 0x02DC: "\u0303", 0x00AF: "\u0304", 0x02C9: "\u0304", 0x02D8: "\u0306", 0x02D9: "\u0307", 0x00A8: "\u0308", 0x00B0: "\u030a", 0x02DA: "\u030a", 0x02BA: "\u030b", 0x02DD: "\u030b", 0x02C7: "\u030c", 0x02C8: "\u030d", 0x0022: "\u030e", 0x02BB: "\u0312", 0x02BC: "\u0313", 0x0486: "\u0313", 0x055A: "\u0313", 0x02BD: "\u0314", 0x0485: "\u0314", 0x0559: "\u0314", 0x02D4: "\u031d", 0x02D5: "\u031e", 0x02D6: "\u031f", 0x02D7: "\u0320", 0x02B2: "\u0321", 0x00B8: "\u0327", 0x02CC: "\u0329", 0x02B7: "\u032b", 0x02CD: "\u0331", 0x005F: "\u0332", 0x204E: "\u0359"}
)

func (_ceea *textTable) isExportable() bool {
	if _ceea._gfagc {
		return true
	}
	_dcgac := func(_aggec int) bool {
		_bcbd := _ceea.get(0, _aggec)
		if _bcbd == nil {
			return false
		}
		_fgdbb := _bcbd.text()
		_bedc := _a.RuneCountInString(_fgdbb)
		_afaa := _gdac.MatchString(_fgdbb)
		return _bedc <= 1 || _afaa
	}
	for _ffbec := 0; _ffbec < _ceea._egbbf; _ffbec++ {
		if !_dcgac(_ffbec) {
			return true
		}
	}
	return false
}
func (_dafce *textObject) getFontDict(_dbbf string) (_fdafa _gbc.PdfObject, _eabf error) {
	_bacad := _dafce._fgf
	if _bacad == nil {
		_b.Log.Debug("g\u0065\u0074\u0046\u006f\u006e\u0074D\u0069\u0063\u0074\u002e\u0020\u004eo\u0020\u0072\u0065\u0073\u006f\u0075\u0072c\u0065\u0073\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025#\u0071", _dbbf)
		return nil, nil
	}
	_fdafa, _dfce := _bacad.GetFontByName(_gbc.PdfObjectName(_dbbf))
	if !_dfce {
		_b.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0067\u0065t\u0046\u006f\u006et\u0044\u0069\u0063\u0074\u003a\u0020\u0046\u006f\u006et \u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020\u006e\u0061m\u0065\u003d%\u0023\u0071", _dbbf)
		return nil, _dg.New("f\u006f\u006e\u0074\u0020no\u0074 \u0069\u006e\u0020\u0072\u0065s\u006f\u0075\u0072\u0063\u0065\u0073")
	}
	return _fdafa, nil
}
func _dcgec(_bcea _bg.PdfRectangle) *ruling {
	return &ruling{_gcafd: _bfcbd, _gbaff: _bcea.Ury, _fabe: _bcea.Llx, _ccba: _bcea.Urx}
}
func _begbef(_agee, _gcgd *textPara) bool { return _degfb(_agee._dgcfe, _gcgd._dgcfe) }
func (_aeba rulingList) blocks(_baea, _edaeb *ruling) bool {
	if _baea._fabe > _edaeb._ccba || _edaeb._fabe > _baea._ccba {
		return false
	}
	_fggfg := _bb.Max(_baea._fabe, _edaeb._fabe)
	_bgda := _bb.Min(_baea._ccba, _edaeb._ccba)
	if _baea._gbaff > _edaeb._gbaff {
		_baea, _edaeb = _edaeb, _baea
	}
	for _, _efef := range _aeba {
		if _baea._gbaff <= _efef._gbaff+_egba && _efef._gbaff <= _edaeb._gbaff+_egba && _efef._fabe <= _bgda && _fggfg <= _efef._ccba {
			return true
		}
	}
	return false
}
func (_bbac compositeCell) String() string {
	_fefge := ""
	if len(_bbac.paraList) > 0 {
		_fefge = _cfbg(_bbac.paraList.merge().text(), 50)
	}
	return _d.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0070\u0061\u0072a\u0073\u0020\u0025\u0071", _bbac.PdfRectangle, len(_bbac.paraList), _fefge)
}
func (_cadf *Extractor) extractPageText(_ccd string, _fdd *_bg.PdfPageResources, _fgae _ef.Matrix, _agd int, _caeg bool) (*PageText, int, int, error) {
	_b.Log.Trace("\u0065x\u0074\u0072\u0061\u0063t\u0050\u0061\u0067\u0065\u0054e\u0078t\u003a \u006c\u0065\u0076\u0065\u006c\u003d\u0025d", _agd)
	_dfc := &PageText{_egdec: _cadf._ac, _gge: _cadf._egb, _ecfd: _cadf._edaa}
	_ggf := _fccd(_cadf._ac)
	var _bdc stateStack
	_fdf := _faef(_cadf, _fdd, _aed.GraphicsState{}, &_ggf, &_bdc)
	_cff := shapesState{_gbag: _fgae, _bdab: _ef.IdentityMatrix(), _eadd: _fdf}
	var _dab bool
	_dfd := -1
	_eedb := ""
	if _agd > _ggc {
		_fbd := _dg.New("\u0066\u006f\u0072\u006d s\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077")
		_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0065\u0078\u0074\u0072\u0061\u0063\u0074\u0050\u0061\u0067\u0065\u0054\u0065\u0078\u0074\u002e\u0020\u0072\u0065\u0063u\u0072\u0073\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u003d\u0025\u0064 \u0065r\u0072\u003d\u0025\u0076", _agd, _fbd)
		return _dfc, _ggf._gfff, _ggf._ffbf, _fbd
	}
	_gec := _aed.NewContentStreamParser(_ccd)
	_fbg, _gab := _gec.Parse()
	if _gab != nil {
		_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020e\u0078\u0074\u0072a\u0063\u0074\u0050\u0061g\u0065\u0054\u0065\u0078\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gab)
		return _dfc, _ggf._gfff, _ggf._ffbf, _gab
	}
	_dfc._dbcd = _fbg
	_cdb := _aed.NewContentStreamProcessor(*_fbg)
	_cdb.AddHandler(_aed.HandlerConditionEnumAllOperands, "", func(_ggcf *_aed.ContentStreamOperation, _eeab _aed.GraphicsState, _gbg *_bg.PdfPageResources) error {
		_bafb := _ggcf.Operand
		if _dfga {
			_b.Log.Info("\u0026&\u0026\u0020\u006f\u0070\u003d\u0025s", _ggcf)
		}
		switch _bafb {
		case "\u0071":
			if _fbfd {
				_b.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _cff._bdab)
			}
			_bdc.push(&_ggf)
		case "\u0051":
			if !_bdc.empty() {
				_ggf = *_bdc.pop()
			}
			_cff._bdab = _eeab.CTM
			if _fbfd {
				_b.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _cff._bdab)
			}
		case "\u0042\u0044\u0043":
			_decf, _cfe := _gbc.GetDict(_ggcf.Params[1])
			if !_cfe {
				_b.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0042D\u0043\u0020\u006f\u0070\u003d\u0025\u0073 \u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _ggcf)
				return _gab
			}
			_bdg := _decf.Get("\u004d\u0043\u0049\u0044")
			if _bdg != nil {
				_fcbf, _eedc := _gbc.GetIntVal(_bdg)
				if !_eedc {
					_b.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0044C\u0020\u006f\u0070=\u0025\u0073\u002e\u0020\u0042\u0061\u0064\u0020\u006eum\u0065\u0072\u0069c\u0061\u006c \u006f\u0062\u006a\u0065\u0063\u0074.\u0020\u006f=\u0025\u0073", _ggcf, _bdg)
				}
				_dfd = _fcbf
			} else {
				_dfd = -1
			}
			_ccf := _decf.Get("\u0041\u0063\u0074\u0075\u0061\u006c\u0054\u0065\u0078\u0074")
			if _ccf != nil {
				_eedb = _ccf.String()
			}
		case "\u0045\u004d\u0043":
			_dfd = -1
			_eedb = ""
		case "\u0042\u0054":
			if _dab {
				_b.Log.Debug("\u0042\u0054\u0020\u0063\u0061\u006c\u006c\u0065\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0069n\u0020\u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
				_dfc._faba = append(_dfc._faba, _fdf._gaa...)
			}
			_dab = true
			_fgee := _eeab
			if _caeg {
				_fgee = _aed.GraphicsState{}
				_fgee.CTM = _cff._bdab
			}
			_fgee.CTM = _fgae.Mult(_fgee.CTM)
			_fdf = _faef(_cadf, _gbg, _fgee, &_ggf, &_bdc)
			_cff._eadd = _fdf
		case "\u0045\u0054":
			if !_dab {
				_b.Log.Debug("\u0045\u0054\u0020ca\u006c\u006c\u0065\u0064\u0020\u006f\u0075\u0074\u0073i\u0064e\u0020o\u0066 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_dab = false
			_dfc._faba = append(_dfc._faba, _fdf._gaa...)
			_fdf.reset()
		case "\u0054\u002a":
			_fdf.nextLine()
		case "\u0054\u0064":
			if _aae, _eefg := _fdf.checkOp(_ggcf, 2, true); !_aae {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eefg)
				return _eefg
			}
			_ddb, _fbdf, _aec := _bgcd(_ggcf.Params)
			if _aec != nil {
				return _aec
			}
			_fdf.moveText(_ddb, _fbdf)
		case "\u0054\u0044":
			if _bffd, _bfc := _fdf.checkOp(_ggcf, 2, true); !_bffd {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bfc)
				return _bfc
			}
			_cdgb, _agf, _ggfd := _bgcd(_ggcf.Params)
			if _ggfd != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ggfd)
				return _ggfd
			}
			_fdf.moveTextSetLeading(_cdgb, _agf)
		case "\u0054\u006a":
			if _fae, _afe := _fdf.checkOp(_ggcf, 1, true); !_fae {
				_b.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0054\u006a\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d%\u0076", _ggcf, _afe)
				return _afe
			}
			_ebb := _gbc.TraceToDirectObject(_ggcf.Params[0])
			_aac, _bfg := _gbc.GetStringBytes(_ebb)
			if !_bfg {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020T\u006a\u0020o\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074S\u0074\u0072\u0069\u006e\u0067\u0042\u0079\u0074\u0065\u0073\u0020\u0066a\u0069\u006c\u0065\u0064", _ggcf)
				return _gbc.ErrTypeError
			}
			return _fdf.showText(_ebb, _aac, _dfd, _eedb)
		case "\u0054\u004a":
			if _gde, _gaba := _fdf.checkOp(_ggcf, 1, true); !_gde {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gaba)
				return _gaba
			}
			_cdaa, _faf := _gbc.GetArray(_ggcf.Params[0])
			if !_faf {
				_b.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _ggcf)
				return _gab
			}
			return _fdf.showTextAdjusted(_cdaa, _dfd, _eedb)
		case "\u0027":
			if _bcc, _bae := _fdf.checkOp(_ggcf, 1, true); !_bcc {
				_b.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0027\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bae)
				return _bae
			}
			_ecf := _gbc.TraceToDirectObject(_ggcf.Params[0])
			_bed, _aebb := _gbc.GetStringBytes(_ecf)
			if !_aebb {
				_b.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020'\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _ggcf)
				return _gbc.ErrTypeError
			}
			_fdf.nextLine()
			return _fdf.showText(_ecf, _bed, _dfd, _eedb)
		case "\u0022":
			if _bea, _gda := _fdf.checkOp(_ggcf, 3, true); !_bea {
				_b.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0022\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gda)
				return _gda
			}
			_add, _gacf, _efb := _bgcd(_ggcf.Params[:2])
			if _efb != nil {
				return _efb
			}
			_age := _gbc.TraceToDirectObject(_ggcf.Params[2])
			_cgdc, _decb := _gbc.GetStringBytes(_age)
			if !_decb {
				_b.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020\"\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _ggcf)
				return _gbc.ErrTypeError
			}
			_fdf.setCharSpacing(_add)
			_fdf.setWordSpacing(_gacf)
			_fdf.nextLine()
			return _fdf.showText(_age, _cgdc, _dfd, _eedb)
		case "\u0054\u004c":
			_cgfe, _ffe := _cgb(_ggcf)
			if _ffe != nil {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004c\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ffe)
				return _ffe
			}
			_fdf.setTextLeading(_cgfe)
		case "\u0054\u0063":
			_bgb, _bfgd := _cgb(_ggcf)
			if _bfgd != nil {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0063\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bfgd)
				return _bfgd
			}
			_fdf.setCharSpacing(_bgb)
		case "\u0054\u0066":
			if _ccg, _eba := _fdf.checkOp(_ggcf, 2, true); !_ccg {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0066\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eba)
				return _eba
			}
			_gabc, _ead := _gbc.GetNameVal(_ggcf.Params[0])
			if !_ead {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _ggcf)
				return _gbc.ErrTypeError
			}
			_fda, _aea := _gbc.GetNumberAsFloat(_ggcf.Params[1])
			if !_ead {
				_b.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u0046\u006c\u006f\u0061\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065d\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ggcf, _aea)
				return _aea
			}
			_aea = _fdf.setFont(_gabc, _fda)
			_fdf._gea = _dg.Is(_aea, _gbc.ErrNotSupported)
			if _aea != nil && !_fdf._gea {
				return _aea
			}
		case "\u0054\u006d":
			if _dgdeg, _baa := _fdf.checkOp(_ggcf, 6, true); !_dgdeg {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u006d\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _baa)
				return _baa
			}
			_dcc, _dcgf := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _dcgf != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dcgf)
				return _dcgf
			}
			_fdf.setTextMatrix(_dcc)
		case "\u0054\u0072":
			if _dfgb, _ffef := _fdf.checkOp(_ggcf, 1, true); !_dfgb {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0072\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ffef)
				return _ffef
			}
			_fffe, _aee := _gbc.GetIntVal(_ggcf.Params[0])
			if !_aee {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0020\u006f\u0070\u003d\u0025\u0073 \u0047e\u0074\u0049\u006e\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _ggcf)
				return _gbc.ErrTypeError
			}
			_fdf.setTextRenderMode(_fffe)
		case "\u0054\u0073":
			if _dac, _cac := _fdf.checkOp(_ggcf, 1, true); !_dac {
				_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cac)
				return _cac
			}
			_fec, _dccg := _gbc.GetNumberAsFloat(_ggcf.Params[0])
			if _dccg != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dccg)
				return _dccg
			}
			_fdf.setTextRise(_fec)
		case "\u0054\u0077":
			if _acc, _afcg := _fdf.checkOp(_ggcf, 1, true); !_acc {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _afcg)
				return _afcg
			}
			_bagb, _afg := _gbc.GetNumberAsFloat(_ggcf.Params[0])
			if _afg != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _afg)
				return _afg
			}
			_fdf.setWordSpacing(_bagb)
		case "\u0054\u007a":
			if _fefg, _gabd := _fdf.checkOp(_ggcf, 1, true); !_fefg {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gabd)
				return _gabd
			}
			_eegb, _dgfb := _gbc.GetNumberAsFloat(_ggcf.Params[0])
			if _dgfb != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgfb)
				return _dgfb
			}
			_fdf.setHorizScaling(_eegb)
		case "\u0063\u006d":
			if !_caeg {
				_cff._bdab = _eeab.CTM
			}
			if _cff._bdab.Singular() {
				_gabcb := _ef.IdentityMatrix().Translate(_cff._bdab.Translation())
				_b.Log.Debug("S\u0069n\u0067\u0075\u006c\u0061\u0072\u0020\u0063\u0074m\u003d\u0025\u0073\u2192%s", _cff._bdab, _gabcb)
				_cff._bdab = _gabcb
			}
			if _fbfd {
				_b.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _cff._bdab)
			}
		case "\u006d":
			if len(_ggcf.Params) != 2 {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _adc)
				return nil
			}
			_fcc, _aag := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _aag != nil {
				return _aag
			}
			_cff.moveTo(_fcc[0], _fcc[1])
		case "\u006c":
			if len(_ggcf.Params) != 2 {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _adc)
				return nil
			}
			_agc, _cecc := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _cecc != nil {
				return _cecc
			}
			_cff.lineTo(_agc[0], _agc[1])
		case "\u0063":
			if len(_ggcf.Params) != 6 {
				return _adc
			}
			_egde, _ffb := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _ffb != nil {
				return _ffb
			}
			_b.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _egde)
			_cff.cubicTo(_egde[0], _egde[1], _egde[2], _egde[3], _egde[4], _egde[5])
		case "\u0076", "\u0079":
			if len(_ggcf.Params) != 4 {
				return _adc
			}
			_agdc, _feac := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _feac != nil {
				return _feac
			}
			_b.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _agdc)
			_cff.quadraticTo(_agdc[0], _agdc[1], _agdc[2], _agdc[3])
		case "\u0068":
			_cff.closePath()
		case "\u0072\u0065":
			if len(_ggcf.Params) != 4 {
				return _adc
			}
			_def, _fcce := _gbc.GetNumbersAsFloat(_ggcf.Params)
			if _fcce != nil {
				return _fcce
			}
			_cff.drawRectangle(_def[0], _def[1], _def[2], _def[3])
			_cff.closePath()
		case "\u0053":
			_cff.stroke(&_dfc._fce)
			_cff.clearPath()
		case "\u0073":
			_cff.closePath()
			_cff.stroke(&_dfc._fce)
			_cff.clearPath()
		case "\u0046":
			_cff.fill(&_dfc._bdfb)
			_cff.clearPath()
		case "\u0066", "\u0066\u002a":
			_cff.closePath()
			_cff.fill(&_dfc._bdfb)
			_cff.clearPath()
		case "\u0042", "\u0042\u002a":
			_cff.fill(&_dfc._bdfb)
			_cff.stroke(&_dfc._fce)
			_cff.clearPath()
		case "\u0062", "\u0062\u002a":
			_cff.closePath()
			_cff.fill(&_dfc._bdfb)
			_cff.stroke(&_dfc._fce)
			_cff.clearPath()
		case "\u006e":
			_cff.clearPath()
		case "\u0044\u006f":
			if len(_ggcf.Params) == 0 {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0058\u004fbj\u0065c\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0070\u0065\u0072\u0061n\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u006f\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072.\u0020\u0047\u006f\u0074\u0020\u0025\u002b\u0076\u002e", _ggcf.Params)
				return _gbc.ErrRangeError
			}
			_bdgf, _bec := _gbc.GetName(_ggcf.Params[0])
			if !_bec {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u006f\u0020\u006f\u0070e\u0072a\u0074\u006f\u0072\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006fp\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u002b\u0076\u002e", _ggcf.Params[0])
				return _gbc.ErrTypeError
			}
			_, _debc := _gbg.GetXObjectByName(*_bdgf)
			if _debc != _bg.XObjectTypeForm {
				break
			}
			_dfb, _bec := _cadf._gdfd[_bdgf.String()]
			if !_bec {
				_ecc, _dggb := _gbg.GetXObjectFormByName(*_bdgf)
				if _dggb != nil {
					_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dggb)
					return _dggb
				}
				_bgff, _dggb := _ecc.GetContentStream()
				if _dggb != nil {
					_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dggb)
					return _dggb
				}
				_affb := _ecc.Resources
				if _affb == nil {
					_affb = _gbg
				}
				_aaec := _eeab.CTM
				if _accd, _agb := _gbc.GetArray(_ecc.Matrix); _agb {
					_agg, _bee := _accd.GetAsFloat64Slice()
					if _bee != nil {
						return _bee
					}
					if len(_agg) != 6 {
						return _adc
					}
					_ddf := _ef.NewMatrix(_agg[0], _agg[1], _agg[2], _agg[3], _agg[4], _agg[5])
					_aaec = _eeab.CTM.Mult(_ddf)
				}
				_ffeg, _agfg, _dea, _dggb := _cadf.extractPageText(string(_bgff), _affb, _fgae.Mult(_aaec), _agd+1, false)
				if _dggb != nil {
					_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dggb)
					return _dggb
				}
				_dfb = textResult{*_ffeg, _agfg, _dea}
				_cadf._gdfd[_bdgf.String()] = _dfb
			}
			_cff._bdab = _eeab.CTM
			if _fbfd {
				_b.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _cff._bdab)
			}
			_dfc._faba = append(_dfc._faba, _dfb._eafb._faba...)
			_dfc._fce = append(_dfc._fce, _dfb._eafb._fce...)
			_dfc._bdfb = append(_dfc._bdfb, _dfb._eafb._bdfb...)
			_ggf._gfff += _dfb._baca
			_ggf._ffbf += _dfb._eagd
		case "\u0072\u0067", "\u0067", "\u006b", "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_fdf._dba.ColorspaceNonStroking = _eeab.ColorspaceNonStroking
			_fdf._dba.ColorNonStroking = _eeab.ColorNonStroking
		case "\u0052\u0047", "\u0047", "\u004b", "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_fdf._dba.ColorspaceStroking = _eeab.ColorspaceStroking
			_fdf._dba.ColorStroking = _eeab.ColorStroking
		}
		return nil
	})
	_gab = _cdb.Process(_fdd)
	if _cadf._dff != nil && _cadf._dff.IncludeAnnotations && !_caeg {
		for _, _ddg := range _cadf._ga {
			_ceb, _gabac := _gbc.GetDict(_ddg.AP)
			if !_gabac {
				continue
			}
			_eede, _gabac := _ceb.Get("\u004e").(*_gbc.PdfObjectStream)
			if !_gabac {
				continue
			}
			_aecd, _edg := _gbc.DecodeStream(_eede)
			if _edg != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u006f\u006e\u0020\u0064\u0065c\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _edg)
				continue
			}
			_bcbf := _eede.PdfObjectDictionary.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
			_cacc, _edg := _bg.NewPdfPageResourcesFromDict(_bcbf.(*_gbc.PdfObjectDictionary))
			if _edg != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u006f\u006e\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0076", _edg)
				continue
			}
			_ebdc := _ef.IdentityMatrix()
			_ccgd, _gabac := _eede.PdfObjectDictionary.Get("\u004d\u0061\u0074\u0072\u0069\u0078").(*_gbc.PdfObjectArray)
			if _gabac {
				_bagd, _gbaf := _ccgd.GetAsFloat64Slice()
				if _gbaf != nil {
					_b.Log.Debug("\u0045\u0072\u0072or\u0020\u006f\u006e\u0020\u0067\u0065\u0074\u0074\u0069n\u0067 \u0066l\u006fa\u0074\u0036\u0034\u0020\u0073\u006c\u0069\u0063\u0065\u003a\u0020\u0025\u0076", _gbaf)
					continue
				}
				if len(_bagd) != 6 {
					_b.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006d\u0061\u0074\u0072\u0069\u0078\u0020\u0073\u006ci\u0063\u0065\u0020l\u0065n\u0067\u0074\u0068")
					continue
				}
				_ebdc = _ef.NewMatrix(_bagd[0], _bagd[1], _bagd[2], _bagd[3], _bagd[4], _bagd[5])
			}
			_ceaa, _gabac := _cadf._adg[_eede.String()]
			if !_gabac {
				_egac, _dcfd, _efgg, _bfgc := _cadf.extractPageText(string(_aecd), _cacc, _ebdc, _agd+1, true)
				if _bfgc != nil {
					_b.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0065x\u0074\u0072a\u0063\u0074\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0065\u0078\u0074s\u003a\u0020\u0025\u0076", _bfgc)
					continue
				}
				_ceaa = textResult{*_egac, _dcfd, _efgg}
				_cadf._adg[_eede.String()] = _ceaa
			}
			_dfc._faba = append(_dfc._faba, _ceaa._eafb._faba...)
			_dfc._fce = append(_dfc._fce, _ceaa._eafb._fce...)
			_dfc._bdfb = append(_dfc._bdfb, _ceaa._eafb._bdfb...)
			_ggf._gfff += _ceaa._baca
			_ggf._ffbf += _ceaa._eagd
		}
	}
	return _dfc, _ggf._gfff, _ggf._ffbf, _gab
}
func (_acea *textTable) bbox() _bg.PdfRectangle { return _acea.PdfRectangle }

var _de = []string{"\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0053", "\u0042", "\u0053", "\u0057\u0053", "\u0042", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042", "\u0042", "\u0042", "\u0053", "\u0057\u0053", "\u004f\u004e", "\u004f\u004e", "\u0045\u0054", "\u0045\u0054", "\u0045\u0054", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u0045\u0053", "\u0043\u0053", "\u0045\u0053", "\u0043\u0053", "\u0043\u0053", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0043\u0053", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0042\u004e", "\u0043\u0053", "\u004f\u004e", "\u0045\u0054", "\u0045\u0054", "\u0045\u0054", "\u0045\u0054", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004c", "\u004f\u004e", "\u004f\u004e", "\u0042\u004e", "\u004f\u004e", "\u004f\u004e", "\u0045\u0054", "\u0045\u0054", "\u0045\u004e", "\u0045\u004e", "\u004f\u004e", "\u004c", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u0045\u004e", "\u004c", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004f\u004e", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004f\u004e", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004f\u004e", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c", "\u004c"}

func (_badg *textLine) markWordBoundaries() {
	_fcafg := _gcca * _badg._gdbd
	for _ggadb, _ddce := range _badg._edee[1:] {
		if _gefae(_ddce, _badg._edee[_ggadb]) >= _fcafg {
			_ddce._ggdce = true
		}
	}
}

type lists []*list

func (_bfd *textLine) toTextMarks(_feacd *int) []TextMark {
	var _gbage []TextMark
	for _, _fdcb := range _bfd._edee {
		if _fdcb._ggdce {
			_gbage = _eaagb(_gbage, _feacd, "\u0020")
		}
		_aadg := _fdcb.toTextMarks(_feacd)
		_gbage = append(_gbage, _aadg...)
	}
	return _gbage
}
func _dfgba(_bbgda map[int][]float64) {
	if len(_bbgda) <= 1 {
		return
	}
	_fdcda := _bgffa(_bbgda)
	if _eeeg {
		_b.Log.Info("\u0066i\u0078C\u0065\u006c\u006c\u0073\u003a \u006b\u0065y\u0073\u003d\u0025\u002b\u0076", _fdcda)
	}
	var _gfbf, _ffba int
	for _gfbf, _ffba = range _fdcda {
		if _bbgda[_ffba] != nil {
			break
		}
	}
	for _eedd, _cbdg := range _fdcda[_gfbf:] {
		_bfeg := _bbgda[_cbdg]
		if _bfeg == nil {
			continue
		}
		if _eeeg {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u006b\u0030\u003d\u0025\u0064\u0020\u006b1\u003d\u0025\u0064\u000a", _gfbf+_eedd, _ffba, _cbdg)
		}
		_dcdae := _bbgda[_cbdg]
		if _dcdae[len(_dcdae)-1] > _bfeg[0] {
			_dcdae[len(_dcdae)-1] = _bfeg[0]
			_bbgda[_ffba] = _dcdae
		}
		_ffba = _cbdg
	}
}

var _fagb string = "\u0028\u003f\u0069\u0029\u005e\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028?\u003a\u0044\u007cM\u0029\u007c\u0044\u003f\u0043{\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028\u003f\u003a\u004c\u007c\u0043\u0029\u007cL\u003f\u0058\u007b\u0030\u002c\u0033}\u0029\u0028\u0049\u0028\u003f\u003a\u0056\u007c\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u005c\u0029\u007c\u005c\u002e\u0029\u007c\u005e\u005c\u0028\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028\u003f\u003aD\u007cM\u0029\u007c\u0044\u003f\u0043\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028?\u003a\u004c\u007c\u0043\u0029\u007c\u004c?\u0058\u007b0\u002c\u0033\u007d\u0029(\u0049\u0028\u003f\u003a\u0056|\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u005c\u0029"

type pathSection struct {
	_cdc []*subpath
	_ab.Color
}

// Extractor stores and offers functionality for extracting content from PDF pages.
type Extractor struct {
	_eca  string
	_feg  *_bg.PdfPageResources
	_ac   _bg.PdfRectangle
	_fcb  *_bg.PdfRectangle
	_deg  map[string]fontEntry
	_gdfd map[string]textResult
	_adg  map[string]textResult
	_cd   int64
	_egd  int
	_dff  *Options
	_egb  *_gbc.PdfObject
	_edaa _gbc.PdfObject
	_ga   []*_bg.PdfAnnotation
}

// TextMarkArray is a collection of TextMarks.
type TextMarkArray struct{ _egbe []TextMark }

func (_ecfb paraList) log(_fggf string) {
	if !_agba {
		return
	}
	_b.Log.Info("%\u0038\u0073\u003a\u0020\u0025\u0064 \u0070\u0061\u0072\u0061\u0073\u0020=\u003d\u003d\u003d\u003d\u003d\u003d\u002d-\u002d\u002d\u002d\u002d\u002d\u003d\u003d\u003d\u003d\u003d=\u003d", _fggf, len(_ecfb))
	for _ebdg, _acac := range _ecfb {
		if _acac == nil {
			continue
		}
		_eabg := _acac.text()
		_ggbb := "\u0020\u0020"
		if _acac._cece != nil {
			_ggbb = _d.Sprintf("\u005b%\u0064\u0078\u0025\u0064\u005d", _acac._cece._ddcega, _acac._cece._egbbf)
		}
		_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0025s\u0020\u0025\u0071\u000a", _ebdg, _acac.PdfRectangle, _ggbb, _cfbg(_eabg, 50))
	}
}
func (_bcagg intSet) has(_gcgbc int) bool { _, _fedag := _bcagg[_gcgbc]; return _fedag }
func _gdfgg(_cedde []pathSection) {
	if _bagc < 0.0 {
		return
	}
	if _gfab {
		_b.Log.Info("\u0067\u0072\u0061\u006e\u0075\u006c\u0061\u0072\u0069\u007a\u0065\u003a\u0020\u0025\u0064 \u0073u\u0062\u0070\u0061\u0074\u0068\u0020\u0073\u0065\u0063\u0074\u0069\u006f\u006e\u0073", len(_cedde))
	}
	for _ebge, _fgbg := range _cedde {
		for _fdadc, _bacadb := range _fgbg._cdc {
			for _dgdae, _cgcd := range _bacadb._agea {
				_bacadb._agea[_dgdae] = _ef.Point{X: _gbeg(_cgcd.X), Y: _gbeg(_cgcd.Y)}
				if _gfab {
					_adda := _bacadb._agea[_dgdae]
					if !_fedaf(_cgcd, _adda) {
						_fgccf := _ef.Point{X: _adda.X - _cgcd.X, Y: _adda.Y - _cgcd.Y}
						_d.Printf("\u0025\u0034d \u002d\u0020\u00254\u0064\u0020\u002d\u0020%4d\u003a %\u002e\u0032\u0066\u0020\u2192\u0020\u0025.2\u0066\u0020\u0028\u0025\u0067\u0029\u000a", _ebge, _fdadc, _dgdae, _cgcd, _adda, _fgccf)
					}
				}
			}
		}
	}
}
func _gddc(_gaagb, _eedcb float64) bool { return _bb.Abs(_gaagb-_eedcb) <= _beac }
func _degfb(_ecceg, _fbcb _bg.PdfRectangle) bool {
	return _fbcb.Llx <= _ecceg.Urx && _ecceg.Llx <= _fbcb.Urx
}
func _ffbe(_afgbc []*textMark, _aefce _bg.PdfRectangle, _dbeb rulingList, _agdb []gridTiling, _faeaf bool) paraList {
	_b.Log.Trace("\u006d\u0061\u006b\u0065\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u003a \u0025\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073\u0020\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065=\u0025\u002e\u0032\u0066", len(_afgbc), _aefce)
	if len(_afgbc) == 0 {
		return nil
	}
	_cace := _gebg(_afgbc, _aefce)
	if len(_cace) == 0 {
		return nil
	}
	_dbeb.log("\u006d\u0061\u006be\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065")
	_aecg, _aagccg := _dbeb.vertsHorzs()
	_cafc := _gdbe(_cace, _aefce.Ury, _aecg, _aagccg)
	_bebfc := _ccae(_cafc, _aefce.Ury, _aecg, _aagccg)
	_bebfc = _gbaa(_bebfc)
	_ddaa := make(paraList, 0, len(_bebfc))
	for _, _addcb := range _bebfc {
		_adff := _addcb.arrangeText()
		if _adff != nil {
			_ddaa = append(_ddaa, _adff)
		}
	}
	if !_faeaf && len(_ddaa) >= _cbea {
		_ddaa = _ddaa.extractTables(_agdb)
	}
	_ddaa.sortReadingOrder()
	if !_faeaf {
		_ddaa.sortTopoOrder()
	}
	_ddaa.log("\u0073\u006f\u0072te\u0064\u0020\u0069\u006e\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0064\u0065\u0072")
	return _ddaa
}
func (_afbbb *textWord) appendMark(_acdg *textMark, _bgce _bg.PdfRectangle) {
	_afbbb._ecdf = append(_afbbb._ecdf, _acdg)
	_afbbb.PdfRectangle = _egbb(_afbbb.PdfRectangle, _acdg.PdfRectangle)
	if _acdg._cfgcd > _afbbb._dafae {
		_afbbb._dafae = _acdg._cfgcd
	}
	_afbbb._ccee = _bgce.Ury - _afbbb.PdfRectangle.Lly
}

// String returns a description of `tm`.
func (_dbcg *textMark) String() string {
	return _d.Sprintf("\u0025\u002e\u0032f \u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\u0022\u0025\u0073\u0022", _dbcg.PdfRectangle, _dbcg._cfgcd, _dbcg._aaag)
}
func (_fadca *textTable) newTablePara() *textPara {
	_caabg := _fadca.computeBbox()
	_afbb := &textPara{PdfRectangle: _caabg, _dgcfe: _caabg, _cece: _fadca}
	if _eeeg {
		_b.Log.Info("\u006e\u0065w\u0054\u0061\u0062l\u0065\u0050\u0061\u0072\u0061\u003a\u0020\u0025\u0073", _afbb)
	}
	return _afbb
}

// String returns a human readable description of `ss`.
func (_fgeeb *shapesState) String() string {
	return _d.Sprintf("\u007b\u0025\u0064\u0020su\u0062\u0070\u0061\u0074\u0068\u0073\u0020\u0066\u0072\u0065\u0073\u0068\u003d\u0025t\u007d", len(_fgeeb._dgdd), _fgeeb._aaba)
}
func _cada(_dbda, _bagf bounded) float64 {
	_fbcg := _ccfa(_dbda, _bagf)
	if !_gacaa(_fbcg) {
		return _fbcg
	}
	return _fgc(_dbda, _bagf)
}
func (_fcfgd *textPara) writeText(_abbc _e.Writer) {
	if _fcfgd._cece == nil {
		_fcfgd.writeCellText(_abbc)
		return
	}
	for _aafd := 0; _aafd < _fcfgd._cece._egbbf; _aafd++ {
		for _fbbad := 0; _fbbad < _fcfgd._cece._ddcega; _fbbad++ {
			_afeb := _fcfgd._cece.get(_fbbad, _aafd)
			if _afeb == nil {
				_abbc.Write([]byte("\u0009"))
			} else {
				_dbge(_afeb._gdc)
				_afeb.writeCellText(_abbc)
			}
			_abbc.Write([]byte("\u0020"))
		}
		if _aafd < _fcfgd._cece._egbbf-1 {
			_abbc.Write([]byte("\u000a"))
		}
	}
}
func _eaagb(_eefae []TextMark, _eagec *int, _gaad string) []TextMark {
	_cceg := _fdad
	_cceg.Text = _gaad
	return _cefe(_eefae, _eagec, _cceg)
}
func _gefae(_gcad, _acbc bounded) float64 { return _gcad.bbox().Llx - _acbc.bbox().Urx }

// List returns all the list objects detected on the page.
// It detects all the bullet point Lists from a given pdf page and builds a slice of bullet list objects.
// A given bullet list object has a tree structure.
// Each bullet point list is extracted with the text content it contains and all the sub lists found under it as children in the tree.
// The rest content of the pdf is ignored and only text in the bullet point lists are extracted.
// The list extraction is done in two ways.
// 1. If the document is tagged then the lists are extracted using the tags provided in the document.
// 2. Otherwise the bullet lists are extracted from the raw text using regex matching.
// By default the document tag is used if available.
// However this can be disabled using `DisableDocumentTags` in the `Options` object.
// Sometimes disabling document tags option might give a better bullet list extraction if the document was tagged incorrectly.
//
//	    options := &Options{
//		     DisableDocumentTags: false, // this means use document tag if available
//	    }
//	    ex, err := NewWithOptions(page, options)
//	    // handle error
//	    pageText, _, _, err := ex.ExtractPageText()
//	    // handle error
//	    lists := pageText.List()
//	    txt := lists.Text()
func (_dbad PageText) List() lists {
	_dfaf := !_dbad._gfee._abdc
	_deca := _dbad.getParagraphs()
	_fceg := true
	if _dbad._gge == nil || *_dbad._gge == nil {
		_fceg = false
	}
	_bde := _deca.list()
	if _fceg && _dfaf {
		_eaabg := _defc(&_deca)
		_afab := &structTreeRoot{}
		_afab.parseStructTreeRoot(*_dbad._gge)
		if _afab._dgeg == nil {
			_b.Log.Debug("\u004c\u0069\u0073\u0074\u003a\u0020\u0073t\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e'\u0074\u0020\u0068\u0061\u0076e\u0020\u0061\u006e\u0079\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u002c\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u002e")
			return _bde
		}
		_bde = _afab.buildList(_eaabg, _dbad._ecfd)
	}
	return _bde
}
func (_eaadd *shapesState) moveTo(_efaa, _dfe float64) {
	_eaadd._aaba = true
	_eaadd._gbgg = _eaadd.devicePoint(_efaa, _dfe)
	if _fbfd {
		_b.Log.Info("\u006d\u006fv\u0065\u0054\u006f\u003a\u0020\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u003d%.\u0032\u0066", _efaa, _dfe, _eaadd._gbgg)
	}
}
func (_fdb *textObject) setTextLeading(_fca float64) {
	if _fdb == nil {
		return
	}
	_fdb._egf._fba = _fca
}
func _aaegg(_fede []*textWord, _ebac int) []*textWord {
	_ebbag := len(_fede)
	copy(_fede[_ebac:], _fede[_ebac+1:])
	return _fede[:_ebbag-1]
}
func _edbf(_eebb *textWord, _cgda float64, _dacfg, _eadf rulingList) *wordBag {
	_baad := _fdc(_eebb._ccee)
	_cdcc := []*textWord{_eebb}
	_bgfaa := wordBag{_cccf: map[int][]*textWord{_baad: _cdcc}, PdfRectangle: _eebb.PdfRectangle, _aced: _eebb._dafae, _cfb: _cgda, _ecdb: _dacfg, _ebfag: _eadf}
	return &_bgfaa
}
func _ccbc(_egce _bg.PdfRectangle) *ruling {
	return &ruling{_gcafd: _cege, _gbaff: _egce.Llx, _fabe: _egce.Lly, _ccba: _egce.Ury}
}
func _ffff(_dcae, _fbba _bg.PdfRectangle) bool {
	return _dcae.Llx <= _fbba.Llx && _fbba.Urx <= _dcae.Urx && _dcae.Lly <= _fbba.Lly && _fbba.Ury <= _dcae.Ury
}
func (_dgeea *textTable) computeBbox() _bg.PdfRectangle {
	var _bgca _bg.PdfRectangle
	_gdfbe := false
	for _gdgec := 0; _gdgec < _dgeea._egbbf; _gdgec++ {
		for _febcgf := 0; _febcgf < _dgeea._ddcega; _febcgf++ {
			_cgffff := _dgeea.get(_febcgf, _gdgec)
			if _cgffff == nil {
				continue
			}
			if !_gdfbe {
				_bgca = _cgffff.PdfRectangle
				_gdfbe = true
			} else {
				_bgca = _egbb(_bgca, _cgffff.PdfRectangle)
			}
		}
	}
	return _bgca
}
func (_cfgg *wordBag) scanBand(_ffda string, _dbfe *wordBag, _cadd func(_fedg *wordBag, _ddab *textWord) bool, _eaea, _bgfg, _ecg float64, _edcad, _dega bool) int {
	_bbdb := _dbfe._aced
	var _gafg map[int]map[*textWord]struct{}
	if !_edcad {
		_gafg = _cfgg.makeRemovals()
	}
	_faee := _gfgc * _bbdb
	_gacc := 0
	for _, _cgdfa := range _cfgg.depthBand(_eaea-_faee, _bgfg+_faee) {
		if len(_cfgg._cccf[_cgdfa]) == 0 {
			continue
		}
		for _, _fag := range _cfgg._cccf[_cgdfa] {
			if !(_eaea-_faee <= _fag._ccee && _fag._ccee <= _bgfg+_faee) {
				continue
			}
			if !_cadd(_dbfe, _fag) {
				continue
			}
			_egga := 2.0 * _bb.Abs(_fag._dafae-_dbfe._aced) / (_fag._dafae + _dbfe._aced)
			_ccge := _bb.Max(_fag._dafae/_dbfe._aced, _dbfe._aced/_fag._dafae)
			_dddb := _bb.Min(_egga, _ccge)
			if _ecg > 0 && _dddb > _ecg {
				continue
			}
			if _dbfe.blocked(_fag) {
				continue
			}
			if !_edcad {
				_dbfe.pullWord(_fag, _cgdfa, _gafg)
			}
			_gacc++
			if !_dega {
				if _fag._ccee < _eaea {
					_eaea = _fag._ccee
				}
				if _fag._ccee > _bgfg {
					_bgfg = _fag._ccee
				}
			}
			if _edcad {
				break
			}
		}
	}
	if !_edcad {
		_cfgg.applyRemovals(_gafg)
	}
	return _gacc
}
func _bbgg(_acffa _bg.PdfRectangle) *ruling {
	return &ruling{_gcafd: _cege, _gbaff: _acffa.Urx, _fabe: _acffa.Lly, _ccba: _acffa.Ury}
}

// TableCell is a cell in a TextTable.
type TableCell struct {
	_bg.PdfRectangle

	// Text is the extracted text.
	Text string

	// Marks returns the TextMarks corresponding to the text in Text.
	Marks TextMarkArray
}

func (_aebc *wordBag) allWords() []*textWord {
	var _accf []*textWord
	for _, _dfcd := range _aebc._cccf {
		_accf = append(_accf, _dfcd...)
	}
	return _accf
}

type compositeCell struct {
	_bg.PdfRectangle
	paraList
}
type paraList []*textPara

func (_cdad *shapesState) lineTo(_eagc, _gabaa float64) {
	if _fbfd {
		_b.Log.Info("\u006c\u0069\u006eeT\u006f\u0028\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0070\u003d\u0025\u002e\u0032\u0066", _eagc, _gabaa, _cdad.devicePoint(_eagc, _gabaa))
	}
	_cdad.addPoint(_eagc, _gabaa)
}
func (_gcbf *wordBag) depthRange(_caba, _ebcg int) []int {
	var _egdf []int
	for _bebf := range _gcbf._cccf {
		if _caba <= _bebf && _bebf <= _ebcg {
			_egdf = append(_egdf, _bebf)
		}
	}
	if len(_egdf) == 0 {
		return nil
	}
	_gd.Ints(_egdf)
	return _egdf
}
func (_fgca *textLine) endsInHyphen() bool {
	_ebaf := _fgca._edee[len(_fgca._edee)-1]
	_dffb := _ebaf._edac
	_gfgbg, _dcgee := _a.DecodeLastRuneInString(_dffb)
	if _dcgee <= 0 || !_ae.Is(_ae.Hyphen, _gfgbg) {
		return false
	}
	if _ebaf._ggdce && _fafdc(_dffb) {
		return true
	}
	return _fafdc(_fgca.text())
}
func _dcacd(_dceag []*textMark, _edfbb _bg.PdfRectangle) *textWord {
	_ccedf := _dceag[0].PdfRectangle
	_ddae := _dceag[0]._cfgcd
	for _, _dccc := range _dceag[1:] {
		_ccedf = _egbb(_ccedf, _dccc.PdfRectangle)
		if _dccc._cfgcd > _ddae {
			_ddae = _dccc._cfgcd
		}
	}
	return &textWord{PdfRectangle: _ccedf, _ecdf: _dceag, _ccee: _edfbb.Ury - _ccedf.Lly, _dafae: _ddae}
}
func (_ccbd rulingList) snapToGroupsDirection() rulingList {
	_ccbd.sortStrict()
	_eagge := make(map[*ruling]rulingList, len(_ccbd))
	_ecag := _ccbd[0]
	_dfbab := func(_bgfd *ruling) { _ecag = _bgfd; _eagge[_ecag] = rulingList{_bgfd} }
	_dfbab(_ccbd[0])
	for _, _dgdb := range _ccbd[1:] {
		if _dgdb._gbaff < _ecag._gbaff-_cfaf {
			_b.Log.Error("\u0073\u006e\u0061\u0070T\u006f\u0047\u0072\u006f\u0075\u0070\u0073\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0057\u0072\u006f\u006e\u0067\u0020\u0070\u0072\u0069\u006da\u0072\u0079\u0020\u006f\u0072d\u0065\u0072\u002e\u000a\u0009\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0020\u0076\u003d\u0025\u0073", _ecag, _dgdb)
		}
		if _dgdb._gbaff > _ecag._gbaff+_egba {
			_dfbab(_dgdb)
		} else {
			_eagge[_ecag] = append(_eagge[_ecag], _dgdb)
		}
	}
	_dgec := make(map[*ruling]float64, len(_eagge))
	_dedc := make(map[*ruling]*ruling, len(_ccbd))
	for _cefbbe, _fgbeg := range _eagge {
		_dgec[_cefbbe] = _fgbeg.mergePrimary()
		for _, _fdgbb := range _fgbeg {
			_dedc[_fdgbb] = _cefbbe
		}
	}
	for _, _fcab := range _ccbd {
		_fcab._gbaff = _dgec[_dedc[_fcab]]
	}
	_cebg := make(rulingList, 0, len(_ccbd))
	for _, _bbbf := range _eagge {
		_gfbga := _bbbf.splitSec()
		for _decbf, _bcegb := range _gfbga {
			_gbcb := _bcegb.merge()
			if len(_cebg) > 0 {
				_bdbg := _cebg[len(_cebg)-1]
				if _bdbg.alignsPrimary(_gbcb) && _bdbg.alignsSec(_gbcb) {
					_b.Log.Error("\u0073\u006e\u0061\u0070\u0054\u006fG\u0072\u006f\u0075\u0070\u0073\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0044\u0075\u0070\u006ci\u0063\u0061\u0074\u0065\u0020\u0069\u003d\u0025\u0064\u000a\u0009\u0077\u003d\u0025s\u000a\t\u0076\u003d\u0025\u0073", _decbf, _bdbg, _gbcb)
					continue
				}
			}
			_cebg = append(_cebg, _gbcb)
		}
	}
	_cebg.sortStrict()
	return _cebg
}
func (_eagdd *PageText) getParagraphs() paraList {
	var _gbgc rulingList
	if _dccd {
		_gfag := _feeb(_eagdd._fce)
		_gbgc = append(_gbgc, _gfag...)
	}
	if _bfac {
		_cgcc := _dfec(_eagdd._bdfb)
		_gbgc = append(_gbgc, _cgcc...)
	}
	_gbgc, _gdee := _gbgc.toTilings()
	var _aedg paraList
	_deae := len(_eagdd._faba)
	for _cebd := 0; _cebd < 360 && _deae > 0; _cebd += 90 {
		_cbf := make([]*textMark, 0, len(_eagdd._faba)-_deae)
		for _, _aaca := range _eagdd._faba {
			if _aaca._ecefa == _cebd {
				_cbf = append(_cbf, _aaca)
			}
		}
		if len(_cbf) > 0 {
			_gdec := _ffbe(_cbf, _eagdd._egdec, _gbgc, _gdee, _eagdd._gfee._ddcf)
			_aedg = append(_aedg, _gdec...)
			_deae -= len(_cbf)
		}
	}
	return _aedg
}
func (_cgcg *textTable) reduce() *textTable {
	_adbbe := make([]int, 0, _cgcg._egbbf)
	_dbbe := make([]int, 0, _cgcg._ddcega)
	for _fbcge := 0; _fbcge < _cgcg._egbbf; _fbcge++ {
		if !_cgcg.emptyCompositeRow(_fbcge) {
			_adbbe = append(_adbbe, _fbcge)
		}
	}
	for _gcgc := 0; _gcgc < _cgcg._ddcega; _gcgc++ {
		if !_cgcg.emptyCompositeColumn(_gcgc) {
			_dbbe = append(_dbbe, _gcgc)
		}
	}
	if len(_adbbe) == _cgcg._egbbf && len(_dbbe) == _cgcg._ddcega {
		return _cgcg
	}
	_befee := textTable{_gfagc: _cgcg._gfagc, _ddcega: len(_dbbe), _egbbf: len(_adbbe), _bcec: make(map[uint64]*textPara, len(_dbbe)*len(_adbbe))}
	if _eeeg {
		_b.Log.Info("\u0072\u0065\u0064\u0075ce\u003a\u0020\u0025\u0064\u0078\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064\u0078%\u0064", _cgcg._ddcega, _cgcg._egbbf, len(_dbbe), len(_adbbe))
		_b.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _dbbe)
		_b.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _adbbe)
	}
	for _gffdb, _ebefe := range _adbbe {
		for _fgdbc, _acae := range _dbbe {
			_eaae, _fbbadd := _cgcg.getComposite(_acae, _ebefe)
			if _eaae == nil {
				continue
			}
			if _eeeg {
				_d.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _fgdbc, _gffdb, _acae, _ebefe, _cfbg(_eaae.merge().text(), 50))
			}
			_befee.putComposite(_fgdbc, _gffdb, _eaae, _fbbadd)
		}
	}
	return &_befee
}

type imageExtractContext struct {
	_aaa  []ImageMark
	_egbg int
	_cdg  int
	_da   int
	_bag  map[*_gbc.PdfObjectStream]*cachedImage
	_gfgg *ImageExtractOptions
	_ebe  bool
}

var _db = []string{"\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u004f\u004e", "\u004f\u004e", "\u0041\u004c", "\u0045\u0054", "\u0045\u0054", "\u0041\u004c", "\u0043\u0053", "\u0041\u004c", "\u004f\u004e", "\u004f\u004e", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u0041\u004c", "\u0041\u004c", "", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0041\u004e", "\u0045\u0054", "\u0041\u004e", "\u0041\u004e", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u004e\u0053\u004d", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u0041\u004e", "\u004f\u004e", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u0041\u004c", "\u0041\u004c", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004f\u004e", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u004e\u0053\u004d", "\u0041\u004c", "\u0041\u004c", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0045\u004e", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c", "\u0041\u004c"}

func _ad(_fd []string, _dgc int, _ba int, _ee string) {
	for _dc := _dgc; _dc < _ba; _dc++ {
		_fd[_dc] = _ee
	}
}

// ExtractPageImages returns the image contents of the page extractor, including data
// and position, size information for each image.
// A set of options to control page image extraction can be passed in. The options
// parameter can be nil for the default options. By default, inline stencil masks
// are not extracted.
func (_ceca *Extractor) ExtractPageImages(options *ImageExtractOptions) (*PageImages, error) {
	_eefe := &imageExtractContext{_gfgg: options}
	_efab := _eefe.extractContentStreamImages(_ceca._eca, _ceca._feg)
	if _efab != nil {
		return nil, _efab
	}
	return &PageImages{Images: _eefe._aaa}, nil
}
func _afgcc(_adfc []rulingList) (rulingList, rulingList) {
	var _cgbb rulingList
	for _, _fbbfg := range _adfc {
		_cgbb = append(_cgbb, _fbbfg...)
	}
	return _cgbb.vertsHorzs()
}
func (_gaaf paraList) sortReadingOrder() {
	_b.Log.Trace("\u0073\u006fr\u0074\u0052\u0065\u0061\u0064i\u006e\u0067\u004f\u0072\u0064e\u0072\u003a\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_gaaf))
	if len(_gaaf) <= 1 {
		return
	}
	_gaaf.computeEBBoxes()
	_gd.Slice(_gaaf, func(_dbbb, _dged int) bool { return _cfda(_gaaf[_dbbb], _gaaf[_dged]) <= 0 })
}

type subpath struct {
	_agea []_ef.Point
	_gdge bool
}

// String returns a description of `b`.
func (_abfa *wordBag) String() string {
	var _adab []string
	for _, _dbe := range _abfa.depthIndexes() {
		_bcfc := _abfa._cccf[_dbe]
		for _, _cabb := range _bcfc {
			_adab = append(_adab, _cabb._edac)
		}
	}
	return _d.Sprintf("\u0025.\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065=\u0025\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0025\u0071", _abfa.PdfRectangle, _abfa._aced, len(_adab), _adab)
}

// String returns a string describing `pt`.
func (_efgf PageText) String() string {
	_gfgb := _d.Sprintf("P\u0061\u0067\u0065\u0054ex\u0074:\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073", len(_efgf._faba))
	_bbc := []string{"\u002d" + _gfgb}
	for _, _aaf := range _efgf._faba {
		_bbc = append(_bbc, _aaf.String())
	}
	_bbc = append(_bbc, "\u002b"+_gfgb)
	return _ge.Join(_bbc, "\u000a")
}

type intSet map[int]struct{}

func (_eddd *textTable) logComposite(_fbbd string) {
	if !_eeeg {
		return
	}
	_b.Log.Info("\u007e~\u007eP\u0061\u0072\u0061\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _eddd._ddcega, _eddd._egbbf, _fbbd)
	_d.Printf("\u0025\u0035\u0073 \u007c", "")
	for _ccgaag := 0; _ccgaag < _eddd._ddcega; _ccgaag++ {
		_d.Printf("\u0025\u0033\u0064 \u007c", _ccgaag)
	}
	_d.Println("")
	_d.Printf("\u0025\u0035\u0073 \u002b", "")
	for _aafga := 0; _aafga < _eddd._ddcega; _aafga++ {
		_d.Printf("\u0025\u0033\u0073 \u002b", "\u002d\u002d\u002d")
	}
	_d.Println("")
	for _ggggb := 0; _ggggb < _eddd._egbbf; _ggggb++ {
		_d.Printf("\u0025\u0035\u0064 \u007c", _ggggb)
		for _bbgga := 0; _bbgga < _eddd._ddcega; _bbgga++ {
			_dbbd, _ := _eddd._fccfa[_caabd(_bbgga, _ggggb)].parasBBox()
			_d.Printf("\u0025\u0033\u0064 \u007c", len(_dbbd))
		}
		_d.Println("")
	}
	_b.Log.Info("\u007e~\u007eT\u0065\u0078\u0074\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _eddd._ddcega, _eddd._egbbf, _fbbd)
	_d.Printf("\u0025\u0035\u0073 \u007c", "")
	for _ggaeb := 0; _ggaeb < _eddd._ddcega; _ggaeb++ {
		_d.Printf("\u0025\u0031\u0032\u0064\u0020\u007c", _ggaeb)
	}
	_d.Println("")
	_d.Printf("\u0025\u0035\u0073 \u002b", "")
	for _cabfb := 0; _cabfb < _eddd._ddcega; _cabfb++ {
		_d.Print("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d\u002b")
	}
	_d.Println("")
	for _dgbfd := 0; _dgbfd < _eddd._egbbf; _dgbfd++ {
		_d.Printf("\u0025\u0035\u0064 \u007c", _dgbfd)
		for _gffc := 0; _gffc < _eddd._ddcega; _gffc++ {
			_fdbca, _ := _eddd._fccfa[_caabd(_gffc, _dgbfd)].parasBBox()
			_gabad := ""
			_ccbfc := _fdbca.merge()
			if _ccbfc != nil {
				_gabad = _ccbfc.text()
			}
			_gabad = _d.Sprintf("\u0025\u0071", _cfbg(_gabad, 12))
			_gabad = _gabad[1 : len(_gabad)-1]
			_d.Printf("\u0025\u0031\u0032\u0073\u0020\u007c", _gabad)
		}
		_d.Println("")
	}
}
func _gfccg(_dfcef _gbc.PdfObject, _efea _ab.Color) (_f.Image, error) {
	_babeb, _fgddb := _gbc.GetStream(_dfcef)
	if !_fgddb {
		return nil, nil
	}
	_bcgfe, _bccag := _bg.NewXObjectImageFromStream(_babeb)
	if _bccag != nil {
		return nil, _bccag
	}
	_ebeb, _bccag := _bcgfe.ToImage()
	if _bccag != nil {
		return nil, _bccag
	}
	return _bfaaa(_ebeb, _efea), nil
}
func _gbggf(_bgd byte) bool {
	for _, _cbabc := range _cdcg {
		if []byte(_cbabc)[0] == _bgd {
			return true
		}
	}
	return false
}
func (_fddb gridTiling) log(_adgd string) {
	if !_degg {
		return
	}
	_b.Log.Info("\u0074i\u006ci\u006e\u0067\u003a\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0071", len(_fddb._cdfg), len(_fddb._fgcd), _adgd)
	_d.Printf("\u0020\u0020\u0020l\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a", _fddb._cdfg)
	_d.Printf("\u0020\u0020\u0020l\u006c\u0079\u003d\u0025\u002e\u0032\u0066\u000a", _fddb._fgcd)
	for _baed, _ebbd := range _fddb._fgcd {
		_dbaef, _fceb := _fddb._egff[_ebbd]
		if !_fceb {
			continue
		}
		_d.Printf("%\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _baed, _ebbd)
		for _eebfg, _ddfe := range _fddb._cdfg {
			_dbebf, _fgebg := _dbaef[_ddfe]
			if !_fgebg {
				continue
			}
			_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eebfg, _dbebf.String())
		}
	}
}
func _dfec(_efbd []pathSection) rulingList {
	_gdfgg(_efbd)
	if _gfab {
		_b.Log.Info("\u006da\u006b\u0065\u0046\u0069l\u006c\u0052\u0075\u006c\u0069n\u0067s\u003a \u0025\u0064\u0020\u0066\u0069\u006c\u006cs", len(_efbd))
	}
	var _gagf rulingList
	for _, _ecga := range _efbd {
		for _, _aagf := range _ecga._cdc {
			if !_aagf.isQuadrilateral() {
				if _gfab {
					_b.Log.Error("!\u0069s\u0051\u0075\u0061\u0064\u0072\u0069\u006c\u0061t\u0065\u0072\u0061\u006c: \u0025\u0073", _aagf)
				}
				continue
			}
			if _cfeffa, _bfgga := _aagf.makeRectRuling(_ecga.Color); _bfgga {
				_gagf = append(_gagf, _cfeffa)
			} else {
				if _bfgg {
					_b.Log.Error("\u0021\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0073", _aagf)
				}
			}
		}
	}
	if _gfab {
		_b.Log.Info("\u006d\u0061\u006b\u0065Fi\u006c\u006c\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0073", _gagf.String())
	}
	return _gagf
}

// Len returns the number of TextMarks in `ma`.
func (_cfgd *TextMarkArray) Len() int {
	if _cfgd == nil {
		return 0
	}
	return len(_cfgd._egbe)
}
func (_bbca *textTable) depth() float64 {
	_ddcfg := 1e10
	for _cfaac := 0; _cfaac < _bbca._ddcega; _cfaac++ {
		_eege := _bbca.get(_cfaac, 0)
		if _eege == nil || _eege._ecee {
			continue
		}
		_ddcfg = _bb.Min(_ddcfg, _eege.depth())
	}
	return _ddcfg
}

var _fcgbb = map[rulingKind]string{_bgeag: "\u006e\u006f\u006e\u0065", _bfcbd: "\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _cege: "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c"}

func _ccacb(_ebacd _gbc.PdfObject, _badb _ab.Color) (_f.Image, error) {
	_adbd, _edaab := _gbc.GetStream(_ebacd)
	if !_edaab {
		return nil, nil
	}
	_gfecg, _dbdbd := _bg.NewXObjectImageFromStream(_adbd)
	if _dbdbd != nil {
		return nil, _dbdbd
	}
	_fgedc, _dbdbd := _gfecg.ToImage()
	if _dbdbd != nil {
		return nil, _dbdbd
	}
	return _cbff(_fgedc, _badb), nil
}
func (_bdcaba paraList) eventNeighbours(_fceef []event) map[*textPara][]int {
	_gd.Slice(_fceef, func(_febfgb, _dgdbf int) bool {
		_eggf, _dgeec := _fceef[_febfgb], _fceef[_dgdbf]
		_eceeg, _fcfgg := _eggf._abbcd, _dgeec._abbcd
		if _eceeg != _fcfgg {
			return _eceeg < _fcfgg
		}
		if _eggf._dcgeb != _dgeec._dcgeb {
			return _eggf._dcgeb
		}
		return _febfgb < _dgdbf
	})
	_bcdb := make(map[int]intSet)
	_edgf := make(intSet)
	for _, _feaba := range _fceef {
		if _feaba._dcgeb {
			_bcdb[_feaba._bbeaga] = make(intSet)
			for _dcaef := range _edgf {
				if _dcaef != _feaba._bbeaga {
					_bcdb[_feaba._bbeaga].add(_dcaef)
					_bcdb[_dcaef].add(_feaba._bbeaga)
				}
			}
			_edgf.add(_feaba._bbeaga)
		} else {
			_edgf.del(_feaba._bbeaga)
		}
	}
	_dfgae := map[*textPara][]int{}
	for _dfdc, _gcff := range _bcdb {
		_beed := _bdcaba[_dfdc]
		if len(_gcff) == 0 {
			_dfgae[_beed] = nil
			continue
		}
		_bfdf := make([]int, len(_gcff))
		_bcbg := 0
		for _abfc := range _gcff {
			_bfdf[_bcbg] = _abfc
			_bcbg++
		}
		_dfgae[_beed] = _bfdf
	}
	return _dfgae
}

// Options extractor options.
type Options struct {

	// DisableDocumentTags specifies whether to use the document tags during list extraction.
	DisableDocumentTags bool

	// ApplyCropBox will extract page text based on page cropbox if set to `true`.
	ApplyCropBox bool

	// UseSimplerExtractionProcess will skip topological text ordering and table processing.
	//
	// NOTE: While normally the extra processing is beneficial, it can also lead to problems when it does not work.
	// Thus it is a flag to allow the user to control this process.
	//
	// Skipping some extraction processes would also lead to the reduced processing time.
	UseSimplerExtractionProcess bool

	// IncludeAnnotations specifies whether to include annotations in the extraction process, default value is `false`.
	IncludeAnnotations bool
}

func (_gdbab intSet) del(_fdfbb int) { delete(_gdbab, _fdfbb) }
func (_ecgde paraList) findGridTables(_cbdc []gridTiling) []*textTable {
	if _eeeg {
		_b.Log.Info("\u0066i\u006e\u0064\u0047\u0072\u0069\u0064\u0054\u0061\u0062\u006c\u0065s\u003a\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073", len(_ecgde))
		for _bgbgf, _dfbf := range _ecgde {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bgbgf, _dfbf)
		}
	}
	var _gfefa []*textTable
	for _babf, _cadff := range _cbdc {
		_efgca, _fgebe := _ecgde.findTableGrid(_cadff)
		if _efgca != nil {
			_efgca.log(_d.Sprintf("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0057\u0069\u0074\u0068\u0047\u0072\u0069\u0064\u0073\u003a\u0020%\u0064", _babf))
			_gfefa = append(_gfefa, _efgca)
			_efgca.markCells()
		}
		for _gcfed := range _fgebe {
			_gcfed._bfaab = true
		}
	}
	if _eeeg {
		_b.Log.Info("\u0066i\u006e\u0064\u0047\u0072i\u0064\u0054\u0061\u0062\u006ce\u0073:\u0020%\u0064\u0020\u0074\u0061\u0062\u006c\u0065s", len(_gfefa))
	}
	return _gfefa
}
func (_cffab *textTable) compositeColCorridors() map[int][]float64 {
	_ecdegd := make(map[int][]float64, _cffab._ddcega)
	if _eeeg {
		_b.Log.Info("\u0063\u006f\u006d\u0070o\u0073\u0069\u0074\u0065\u0043\u006f\u006c\u0043\u006f\u0072r\u0069d\u006f\u0072\u0073\u003a\u0020\u0077\u003d%\u0064\u0020", _cffab._ddcega)
	}
	for _bfecd := 0; _bfecd < _cffab._ddcega; _bfecd++ {
		_ecdegd[_bfecd] = nil
	}
	return _ecdegd
}
func _dce(_gce _ef.Point) *subpath { return &subpath{_agea: []_ef.Point{_gce}} }
func (_efdf *textObject) getFillColor() _ab.Color {
	return _fbbag(_efdf._dba.ColorspaceNonStroking, _efdf._dba.ColorNonStroking)
}

// String returns a description of `w`.
func (_afadd *textWord) String() string {
	return _d.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _afadd._ccee, _afadd.PdfRectangle, _afadd._dafae, _afadd._edac)
}
func (_fcg *shapesState) closePath() {
	if _fcg._aaba {
		_fcg._dgdd = append(_fcg._dgdd, _dce(_fcg._gbgg))
		_fcg._aaba = false
	} else if len(_fcg._dgdd) == 0 {
		if _fbfd {
			_b.Log.Debug("\u0063\u006c\u006f\u0073eP\u0061\u0074\u0068\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u0020\u0070\u0061t\u0068")
		}
		_fcg._aaba = false
		return
	}
	_fcg._dgdd[len(_fcg._dgdd)-1].close()
	if _fbfd {
		_b.Log.Info("\u0063\u006c\u006f\u0073\u0065\u0050\u0061\u0074\u0068\u003a\u0020\u0025\u0073", _fcg)
	}
}
func (_cfde *stateStack) size() int                      { return len(*_cfde) }
func (_ccbff *textTable) get(_adbf, _gcba int) *textPara { return _ccbff._bcec[_caabd(_adbf, _gcba)] }
func (_cfab *subpath) removeDuplicates() {
	if len(_cfab._agea) == 0 {
		return
	}
	_afbf := []_ef.Point{_cfab._agea[0]}
	for _, _edggd := range _cfab._agea[1:] {
		if !_fedaf(_edggd, _afbf[len(_afbf)-1]) {
			_afbf = append(_afbf, _edggd)
		}
	}
	_cfab._agea = _afbf
}
func (_cgfd paraList) tables() []TextTable {
	var _becd []TextTable
	if _eeeg {
		_b.Log.Info("\u0070\u0061\u0072\u0061\u0073\u002e\u0074\u0061\u0062\u006c\u0065\u0073\u003a")
	}
	for _, _dcgfg := range _cgfd {
		_gebb := _dcgfg._cece
		if _gebb != nil && _gebb.isExportable() {
			_becd = append(_becd, _gebb.toTextTable())
		}
	}
	return _becd
}
func _ebfgb(_fadb *textLine, _efae []*textLine, _fade []float64, _aagcc, _dfea float64) []*textLine {
	_gfad := []*textLine{}
	for _, _dfca := range _efae {
		if _dfca._ggfc >= _aagcc {
			if _dfea != -1 && _dfca._ggfc < _dfea {
				if _dfca.text() != _fadb.text() {
					if _bb.Round(_dfca.Llx) < _bb.Round(_fadb.Llx) {
						break
					}
					_gfad = append(_gfad, _dfca)
				}
			} else if _dfea == -1 {
				if _dfca._ggfc == _fadb._ggfc {
					if _dfca.text() != _fadb.text() {
						_gfad = append(_gfad, _dfca)
					}
					continue
				}
				_dbed := _fgfcbe(_fadb, _efae, _fade)
				if _dbed != -1 && _dfca._ggfc <= _dbed {
					_gfad = append(_gfad, _dfca)
				}
			}
		}
	}
	return _gfad
}
func _bcgc(_dadcb map[float64]map[float64]gridTile) []float64 {
	_dcfeg := make([]float64, 0, len(_dadcb))
	for _fecb := range _dadcb {
		_dcfeg = append(_dcfeg, _fecb)
	}
	_gd.Float64s(_dcfeg)
	_fabc := len(_dcfeg)
	for _ffbcf := 0; _ffbcf < _fabc/2; _ffbcf++ {
		_dcfeg[_ffbcf], _dcfeg[_fabc-1-_ffbcf] = _dcfeg[_fabc-1-_ffbcf], _dcfeg[_ffbcf]
	}
	return _dcfeg
}

// NewWithOptions an Extractor instance for extracting content from the input PDF page with options.
func NewWithOptions(page *_bg.PdfPage, options *Options) (*Extractor, error) {
	const _gfa = "\u0065x\u0074\u0072\u0061\u0063\u0074\u006f\u0072\u002e\u004e\u0065\u0077W\u0069\u0074\u0068\u004f\u0070\u0074\u0069\u006f\u006e\u0073"
	_gee, _ada := page.GetAllContentStreams()
	if _ada != nil {
		return nil, _ada
	}
	_fge, _cgd := page.GetStructTreeRoot()
	if !_cgd {
		_b.Log.Debug("T\u0068\u0065\u0020\u0070\u0064\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0074\u0061\u0067g\u0065d\u002e\u0020\u0053\u0074r\u0075\u0063t\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0065\u0078\u0069\u0073\u0074\u002e")
	}
	_adb := page.GetContainingPdfObject()
	_cbad, _ada := page.GetMediaBox()
	if _ada != nil {
		return nil, _d.Errorf("\u0065\u0078\u0074r\u0061\u0063\u0074\u006fr\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u002e\u0020\u0025\u0076", _ada)
	}
	_gae := &Extractor{_eca: _gee, _feg: page.Resources, _ac: *_cbad, _fcb: page.CropBox, _deg: map[string]fontEntry{}, _gdfd: map[string]textResult{}, _adg: map[string]textResult{}, _dff: options, _egb: _fge, _edaa: _adb}
	if _gae._ac.Llx > _gae._ac.Urx {
		_b.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0058\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _gae._ac)
		_gae._ac.Llx, _gae._ac.Urx = _gae._ac.Urx, _gae._ac.Llx
	}
	if _gae._ac.Lly > _gae._ac.Ury {
		_b.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0059\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _gae._ac)
		_gae._ac.Lly, _gae._ac.Ury = _gae._ac.Ury, _gae._ac.Lly
	}
	if _gae._dff != nil {
		if _gae._dff.IncludeAnnotations {
			_gae._ga, _ada = page.GetAnnotations()
			if _ada != nil {
				_b.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0073: \u0025\u0076", _ada)
			}
		}
	}
	_fe.TrackUse(_gfa)
	return _gae, nil
}
func _fafdc(_edcd string) bool {
	if _a.RuneCountInString(_edcd) < _cfgcc {
		return false
	}
	_abbg, _ccdb := _a.DecodeLastRuneInString(_edcd)
	if _ccdb <= 0 || !_ae.Is(_ae.Hyphen, _abbg) {
		return false
	}
	_abbg, _ccdb = _a.DecodeLastRuneInString(_edcd[:len(_edcd)-_ccdb])
	return _ccdb > 0 && !_ae.IsSpace(_abbg)
}
func (_efgd pathSection) bbox() _bg.PdfRectangle {
	_ecce := _efgd._cdc[0]._agea[0]
	_cdbg := _bg.PdfRectangle{Llx: _ecce.X, Urx: _ecce.X, Lly: _ecce.Y, Ury: _ecce.Y}
	_gdfa := func(_cgdf _ef.Point) {
		if _cgdf.X < _cdbg.Llx {
			_cdbg.Llx = _cgdf.X
		} else if _cgdf.X > _cdbg.Urx {
			_cdbg.Urx = _cgdf.X
		}
		if _cgdf.Y < _cdbg.Lly {
			_cdbg.Lly = _cgdf.Y
		} else if _cgdf.Y > _cdbg.Ury {
			_cdbg.Ury = _cgdf.Y
		}
	}
	for _, _becf := range _efgd._cdc[0]._agea[1:] {
		_gdfa(_becf)
	}
	for _, _bbeec := range _efgd._cdc[1:] {
		for _, _gdfg := range _bbeec._agea {
			_gdfa(_gdfg)
		}
	}
	return _cdbg
}
func (_caac rulingList) comp(_eabc, _bdabd int) bool {
	_gcgb, _aegdgc := _caac[_eabc], _caac[_bdabd]
	_bcba, _gcab := _gcgb._gcafd, _aegdgc._gcafd
	if _bcba != _gcab {
		return _bcba > _gcab
	}
	if _bcba == _bgeag {
		return false
	}
	_gabdd := func(_cacf bool) bool {
		if _bcba == _bfcbd {
			return _cacf
		}
		return !_cacf
	}
	_cgca, _dafff := _gcgb._gbaff, _aegdgc._gbaff
	if _cgca != _dafff {
		return _gabdd(_cgca > _dafff)
	}
	_cgca, _dafff = _gcgb._fabe, _aegdgc._fabe
	if _cgca != _dafff {
		return _gabdd(_cgca < _dafff)
	}
	return _gabdd(_gcgb._ccba < _aegdgc._ccba)
}

// TextMark represents extracted text on a page with information regarding both textual content,
// formatting (font and size) and positioning.
// It is the smallest unit of text on a PDF page, typically a single character.
//
// getBBox() in test_text.go shows how to compute bounding boxes of substrings of extracted text.
// The following code extracts the text on PDF page `page` into `text` then finds the bounding box
// `bbox` of substring `term` in `text`.
//
//	ex, _ := New(page)
//	// handle errors
//	pageText, _, _, err := ex.ExtractPageText()
//	// handle errors
//	text := pageText.Text()
//	textMarks := pageText.Marks()
//
//		start := strings.Index(text, term)
//	 end := start + len(term)
//	 spanMarks, err := textMarks.RangeOffset(start, end)
//	 // handle errors
//	 bbox, ok := spanMarks.BBox()
//	 // handle errors
type TextMark struct {

	// Text is the extracted text.
	Text string

	// Original is the text in the PDF. It has not been decoded like `Text`.
	Original string

	// BBox is the bounding box of the text.
	BBox _bg.PdfRectangle

	// Font is the font the text was drawn with.
	Font *_bg.PdfFont

	// FontSize is the font size the text was drawn with.
	FontSize float64

	// Offset is the offset of the start of TextMark.Text in the extracted text. If you do this
	//
	//	text, textMarks := pageText.Text(), pageText.Marks()
	//	marks := textMarks.Elements()
	//
	// then marks[i].Offset is the offset of marks[i].Text in text.
	Offset int

	// Meta is set true for spaces and line breaks that we insert in the extracted text. We insert
	// spaces (line breaks) when we see characters that are over a threshold horizontal (vertical)
	//
	//	distance  apart. See wordJoiner (lineJoiner) in PageText.computeViews().
	Meta bool

	// FillColor is the fill color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	FillColor _ab.Color

	// StrokeColor is the stroke color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	StrokeColor _ab.Color

	// Orientation is the text orientation
	Orientation int

	// DirectObject is the underlying PdfObject (Text Object) that represents the visible texts. This is introduced to get
	// a simple access to the TextObject in case editing or replacment of some text is needed. E.g during redaction.
	DirectObject _gbc.PdfObject

	// ObjString is a decoded string operand of a text-showing operator. It has the same value as `Text` attribute except
	// when many glyphs are represented with the same Text Object that contains multiple length string operand in which case
	// ObjString spans more than one character string that falls in different TextMark objects.
	ObjString []string
	Tw        float64
	Th        float64
	Tc        float64
	Index     int
	_fafb     bool
	_aeaf     *TextTable
}
type textResult struct {
	_eafb PageText
	_baca int
	_eagd int
}

func (_ebgcf *wordBag) getDepthIdx(_aeg float64) int {
	_fgff := _ebgcf.depthIndexes()
	_ddcc := _fdc(_aeg)
	if _ddcc < _fgff[0] {
		return _fgff[0]
	}
	if _ddcc > _fgff[len(_fgff)-1] {
		return _fgff[len(_fgff)-1]
	}
	return _ddcc
}
func _fcda(_feedf *list) []*list {
	var _efda []*list
	for _, _fgfbb := range _feedf._bdfaa {
		switch _fgfbb._dbaad {
		case "\u004c\u0049":
			_ddgdf := _fcgc(_fgfbb)
			_bafbb := _fcda(_fgfbb)
			_aaed := _egfa(_ddgdf, "\u0062\u0075\u006c\u006c\u0065\u0074", _bafbb)
			_gcde := _cgef(_ddgdf, "")
			_aaed._fccf = _gcde
			_efda = append(_efda, _aaed)
		case "\u004c\u0042\u006fd\u0079":
			return _fcda(_fgfbb)
		case "\u004c":
			_daa := _fcda(_fgfbb)
			_efda = append(_efda, _daa...)
			return _efda
		}
	}
	return _efda
}
func _cbff(_afbfg *_bg.Image, _ebgaa _ab.Color) _f.Image {
	_cgdab, _gbegf := int(_afbfg.Width), int(_afbfg.Height)
	_bgbb := _f.NewRGBA(_f.Rect(0, 0, _cgdab, _gbegf))
	for _cgdea := 0; _cgdea < _gbegf; _cgdea++ {
		for _afgda := 0; _afgda < _cgdab; _afgda++ {
			_affaa, _dcfaa := _afbfg.ColorAt(_afgda, _cgdea)
			if _dcfaa != nil {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _afgda, _cgdea)
				continue
			}
			_fgagf, _cfaag, _ddcff, _ := _affaa.RGBA()
			var _bgef _ab.Color
			if _fgagf+_cfaag+_ddcff == 0 {
				_bgef = _ebgaa
			} else {
				_bgef = _ab.Transparent
			}
			_bgbb.Set(_afgda, _cgdea, _bgef)
		}
	}
	return _bgbb
}
func _cfbg(_dbgc string, _ddccb int) string {
	if len(_dbgc) < _ddccb {
		return _dbgc
	}
	return _dbgc[:_ddccb]
}
func (_eaba *ruling) alignsPrimary(_cgdg *ruling) bool {
	return _eaba._gcafd == _cgdg._gcafd && _bb.Abs(_eaba._gbaff-_cgdg._gbaff) < _egba*0.5
}
func _dd(_dcg []string, _ea int, _fg int) {
	for _fga, _eeb := _ea, _fg-1; _fga < _eeb; _fga, _eeb = _fga+1, _eeb-1 {
		_aff := _dcg[_fga]
		_dcg[_fga] = _dcg[_eeb]
		_dcg[_eeb] = _aff
	}
}
func _ccae(_bgad *wordBag, _gdbf float64, _ddge, _afga rulingList) []*wordBag {
	var _adfdc []*wordBag
	for _, _fagbe := range _bgad.depthIndexes() {
		_ffbcb := false
		for !_bgad.empty(_fagbe) {
			_cbb := _bgad.firstReadingIndex(_fagbe)
			_ddbe := _bgad.firstWord(_cbb)
			_fdec := _edbf(_ddbe, _gdbf, _ddge, _afga)
			_bgad.removeWord(_ddbe, _cbb)
			if _gdgf {
				_b.Log.Info("\u0066\u0069\u0072\u0073\u0074\u0057\u006f\u0072\u0064\u0020\u005e\u005e^\u005e\u0020\u0025\u0073", _ddbe.String())
			}
			for _dfdg := true; _dfdg; _dfdg = _ffbcb {
				_ffbcb = false
				_dcec := _ffed * _fdec._aced
				_daag := _cfdb * _fdec._aced
				_bdgg := _ccdef * _fdec._aced
				if _gdgf {
					_b.Log.Info("\u0070a\u0072a\u0057\u006f\u0072\u0064\u0073\u0020\u0064\u0065\u0070\u0074\u0068 \u0025\u002e\u0032\u0066 \u002d\u0020\u0025\u002e\u0032f\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061\u0044\u0065\u0070\u0074\u0068\u0047\u0061\u0070\u003d\u0025\u002e\u0032\u0066\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061R\u0065\u0061\u0064\u0069\u006e\u0067\u0047\u0061p\u003d\u0025\u002e\u0032\u0066", _fdec.minDepth(), _fdec.maxDepth(), _bdgg, _daag)
				}
				if _bgad.scanBand("\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c", _fdec, _fcgb(_cafb, 0), _fdec.minDepth()-_bdgg, _fdec.maxDepth()+_bdgg, _abffb, false, false) > 0 {
					_ffbcb = true
				}
				if _bgad.scanBand("\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _fdec, _fcgb(_cafb, _daag), _fdec.minDepth(), _fdec.maxDepth(), _dbgb, false, false) > 0 {
					_ffbcb = true
				}
				if _ffbcb {
					continue
				}
				_dbcf := _bgad.scanBand("", _fdec, _fcgb(_gag, _dcec), _fdec.minDepth(), _fdec.maxDepth(), _dcda, true, false)
				if _dbcf > 0 {
					_ggac := (_fdec.maxDepth() - _fdec.minDepth()) / _fdec._aced
					if (_dbcf > 1 && float64(_dbcf) > 0.3*_ggac) || _dbcf <= 10 {
						if _bgad.scanBand("\u006f\u0074\u0068e\u0072", _fdec, _fcgb(_gag, _dcec), _fdec.minDepth(), _fdec.maxDepth(), _dcda, false, true) > 0 {
							_ffbcb = true
						}
					}
				}
			}
			_adfdc = append(_adfdc, _fdec)
		}
	}
	return _adfdc
}
func _bcca(_dfba _ef.Point) _ef.Matrix { return _ef.TranslationMatrix(_dfba.X, _dfba.Y) }

// String returns a description of `state`.
func (_acd *textState) String() string {
	_feff := "\u005bN\u004f\u0054\u0020\u0053\u0045\u0054]"
	if _acd._aede != nil {
		_feff = _acd._aede.BaseFont()
	}
	return _d.Sprintf("\u0074\u0063\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0077\u003d\u0025\u002e\u0032\u0066 \u0074f\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071", _acd._fgdd, _acd._eaed, _acd._ebab, _feff)
}

// Append appends `mark` to the mark array.
func (_aefg *TextMarkArray) Append(mark TextMark) { _aefg._egbe = append(_aefg._egbe, mark) }
func (_ggcb *textObject) setCharSpacing(_adcb float64) {
	if _ggcb == nil {
		return
	}
	_ggcb._egf._fgdd = _adcb
	if _eadc {
		_b.Log.Info("\u0073\u0065t\u0043\u0068\u0061\u0072\u0053\u0070\u0061\u0063\u0069\u006e\u0067\u003a\u0020\u0025\u002e\u0032\u0066\u0020\u0073\u0074\u0061\u0074e=\u0025\u0073", _adcb, _ggcb._egf.String())
	}
}
func _fcgc(_ffce *list) []*textLine {
	for _, _bggf := range _ffce._bdfaa {
		switch _bggf._dbaad {
		case "\u004c\u0042\u006fd\u0079":
			if len(_bggf._eadba) != 0 {
				return _bggf._eadba
			}
			return _fcgc(_bggf)
		case "\u0053\u0070\u0061\u006e":
			return _bggf._eadba
		case "I\u006e\u006c\u0069\u006e\u0065\u0053\u0068\u0061\u0070\u0065":
			return _bggf._eadba
		}
	}
	return nil
}
func (_bfcb *textObject) setTextRenderMode(_gef int) {
	if _bfcb == nil {
		return
	}
	_bfcb._egf._dge = RenderMode(_gef)
}

// TextTable represents a table.
// Cells are ordered top-to-bottom, left-to-right.
// Cells[y] is the (0-offset) y'th row in the table.
// Cells[y][x] is the (0-offset) x'th column in the table.
type TextTable struct {
	_bg.PdfRectangle
	W, H  int
	Cells [][]TableCell
}

func (_acdaf paraList) llyOrdering() []int {
	_bafg := make([]int, len(_acdaf))
	for _ceag := range _acdaf {
		_bafg[_ceag] = _ceag
	}
	_gd.SliceStable(_bafg, func(_fdag, _gaec int) bool {
		_bgde, _gabg := _bafg[_fdag], _bafg[_gaec]
		return _acdaf[_bgde].Lly < _acdaf[_gabg].Lly
	})
	return _bafg
}
func (_cagfb *textTable) getRight() paraList {
	_dbgd := make(paraList, _cagfb._egbbf)
	for _gedc := 0; _gedc < _cagfb._egbbf; _gedc++ {
		_fggd := _cagfb.get(_cagfb._ddcega-1, _gedc)._gbfc
		if _fggd.taken() {
			return nil
		}
		_dbgd[_gedc] = _fggd
	}
	for _eggd := 0; _eggd < _cagfb._egbbf-1; _eggd++ {
		if _dbgd[_eggd]._adef != _dbgd[_eggd+1] {
			return nil
		}
	}
	return _dbgd
}
func (_gbbg *imageExtractContext) processOperand(_eed *_aed.ContentStreamOperation, _ede _aed.GraphicsState, _dbb *_bg.PdfPageResources) error {
	if _eed.Operand == "\u0042\u0049" && len(_eed.Params) == 1 {
		_cf, _cefd := _eed.Params[0].(*_aed.ContentStreamInlineImage)
		if !_cefd {
			return nil
		}
		if _acb, _befc := _gbc.GetBoolVal(_cf.ImageMask); _befc {
			if _acb && !_gbbg._gfgg.IncludeInlineStencilMasks {
				return nil
			}
		}
		return _gbbg.extractInlineImage(_cf, _ede, _dbb)
	} else if _eed.Operand == "\u0044\u006f" && len(_eed.Params) == 1 {
		_gdg, _gcb := _gbc.GetName(_eed.Params[0])
		if !_gcb {
			_b.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _gfg
		}
		_, _ega := _dbb.GetXObjectByName(*_gdg)
		switch _ega {
		case _bg.XObjectTypeImage:
			return _gbbg.extractXObjectImage(_gdg, _ede, _dbb)
		case _bg.XObjectTypeForm:
			return _gbbg.extractFormImages(_gdg, _ede, _dbb)
		}
	} else if _gbbg._ebe && (_eed.Operand == "\u0073\u0063\u006e" || _eed.Operand == "\u0053\u0043\u004e") && len(_eed.Params) == 1 {
		_abb, _ffa := _gbc.GetName(_eed.Params[0])
		if !_ffa {
			_b.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _gfg
		}
		_acff, _ffa := _dbb.GetPatternByName(*_abb)
		if !_ffa {
			_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
			return nil
		}
		if _acff.IsTiling() {
			_gfef := _acff.GetAsTilingPattern()
			_dgg, _bbb := _gfef.GetContentStream()
			if _bbb != nil {
				return _bbb
			}
			_bbb = _gbbg.extractContentStreamImages(string(_dgg), _gfef.Resources)
			if _bbb != nil {
				return _bbb
			}
		}
	} else if (_eed.Operand == "\u0063\u0073" || _eed.Operand == "\u0043\u0053") && len(_eed.Params) >= 1 {
		_gbbg._ebe = _eed.Params[0].String() == "\u0050a\u0074\u0074\u0065\u0072\u006e"
	}
	return nil
}

const (
	_abdfb markKind = iota
	_gffff
	_aead
	_cccg
)

func (_edcef paraList) list() []*list {
	var _aecga []*textLine
	var _cdec []*textLine
	for _, _afdc := range _edcef {
		_acee := _afdc.getListLines()
		_aecga = append(_aecga, _acee...)
		_cdec = append(_cdec, _afdc._gdc...)
	}
	_fcbd := _bbf(_aecga)
	_fcfga := _gfcea(_cdec, _fcbd)
	return _fcfga
}
func _eaggd(_fcfea []*textWord, _gggaf *textWord) []*textWord {
	for _gfeef, _begf := range _fcfea {
		if _begf == _gggaf {
			return _aaegg(_fcfea, _gfeef)
		}
	}
	_b.Log.Error("\u0072\u0065\u006d\u006f\u0076e\u0057\u006f\u0072\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u0073\u0020\u0064o\u0065\u0073\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", _gggaf)
	return nil
}

// String returns a description of `k`.
func (_gdbec rulingKind) String() string {
	_ebeg, _bgae := _fcgbb[_gdbec]
	if !_bgae {
		return _d.Sprintf("\u004e\u006ft\u0020\u0061\u0020r\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _gdbec)
	}
	return _ebeg
}
func _caed(_bgf int) bool                      { return (_bgf & 1) != 0 }
func _adcg(_cbcd, _agdg _bg.PdfRectangle) bool { return _degfb(_cbcd, _agdg) && _badd(_cbcd, _agdg) }

// ApplyArea processes the page text only within the specified area `bbox`.
// Each time ApplyArea is called, it updates the result set in `pt`.
// Can be called multiple times in a row with different bounding boxes.
func (_afad *PageText) ApplyArea(bbox _bg.PdfRectangle) {
	_dacf := make([]*textMark, 0, len(_afad._faba))
	for _, _fbe := range _afad._faba {
		if _adcg(_fbe.bbox(), bbox) {
			_dacf = append(_dacf, _fbe)
		}
	}
	var _cbc paraList
	_edbd := len(_dacf)
	for _agfga := 0; _agfga < 360 && _edbd > 0; _agfga += 90 {
		_dabb := make([]*textMark, 0, len(_dacf)-_edbd)
		for _, _cfa := range _dacf {
			if _cfa._ecefa == _agfga {
				_dabb = append(_dabb, _cfa)
			}
		}
		if len(_dabb) > 0 {
			_agcf := _ffbe(_dabb, _afad._egdec, nil, nil, _afad._gfee._ddcf)
			_cbc = append(_cbc, _agcf...)
			_edbd -= len(_dabb)
		}
	}
	_bede := new(_aa.Buffer)
	_cbc.writeText(_bede)
	_afad._bda = _bede.String()
	_afad._dcge = _cbc.toTextMarks()
	_afad._eedcg = _cbc.tables()
}
func _cbee(_cgee, _gdeb _bg.PdfRectangle) (_bg.PdfRectangle, bool) {
	if !_adcg(_cgee, _gdeb) {
		return _bg.PdfRectangle{}, false
	}
	return _bg.PdfRectangle{Llx: _bb.Max(_cgee.Llx, _gdeb.Llx), Urx: _bb.Min(_cgee.Urx, _gdeb.Urx), Lly: _bb.Max(_cgee.Lly, _gdeb.Lly), Ury: _bb.Min(_cgee.Ury, _gdeb.Ury)}, true
}

type lineRuling struct {
	_gadeg rulingKind
	_gfefg markKind
	_ab.Color
	_abfed, _fgcb _ef.Point
}

func (_aedea paraList) lines() []*textLine {
	var _dbfb []*textLine
	for _, _cfbf := range _aedea {
		_dbfb = append(_dbfb, _cfbf._gdc...)
	}
	return _dbfb
}

type textWord struct {
	_bg.PdfRectangle
	_ccee  float64
	_edac  string
	_ecdf  []*textMark
	_dafae float64
	_ggdce bool
}

func (_bbfc *textTable) log(_eegd string) {
	if !_eeeg {
		return
	}
	_b.Log.Info("~\u007e\u007e\u0020\u0025\u0073\u003a \u0025\u0064\u0020\u0078\u0020\u0025d\u0020\u0067\u0072\u0069\u0064\u003d\u0025t\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0025\u0036\u002e2\u0066", _eegd, _bbfc._ddcega, _bbfc._egbbf, _bbfc._gfagc, _bbfc.PdfRectangle)
	for _adacg := 0; _adacg < _bbfc._egbbf; _adacg++ {
		for _dffd := 0; _dffd < _bbfc._ddcega; _dffd++ {
			_dfffg := _bbfc.get(_dffd, _adacg)
			if _dfffg == nil {
				continue
			}
			_d.Printf("%\u0034\u0064\u0020\u00252d\u003a \u0025\u0036\u002e\u0032\u0066 \u0025\u0071\u0020\u0025\u0064\u000a", _dffd, _adacg, _dfffg.PdfRectangle, _cfbg(_dfffg.text(), 50), _a.RuneCountInString(_dfffg.text()))
		}
	}
}

type ruling struct {
	_gcafd rulingKind
	_cebe  markKind
	_ab.Color
	_gbaff float64
	_fabe  float64
	_ccba  float64
	_gcec  float64
}

func (_eeed *ruling) intersects(_abgeb *ruling) bool {
	_cagf := (_eeed._gcafd == _cege && _abgeb._gcafd == _bfcbd) || (_abgeb._gcafd == _cege && _eeed._gcafd == _bfcbd)
	_agege := func(_bgebb, _aaga *ruling) bool {
		return _bgebb._fabe-_beac <= _aaga._gbaff && _aaga._gbaff <= _bgebb._ccba+_beac
	}
	_deaf := _agege(_eeed, _abgeb)
	_fbfdd := _agege(_abgeb, _eeed)
	if _gfab {
		_d.Printf("\u0020\u0020\u0020\u0020\u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003a\u0020\u0020\u006fr\u0074\u0068\u006f\u0067\u006f\u006e\u0061l\u003d\u0025\u0074\u0020\u006f\u0031\u003d\u0025\u0074\u0020\u006f2\u003d\u0025\u0074\u0020\u2192\u0020\u0025\u0074\u000a"+"\u0020\u0020\u0020 \u0020\u0020\u0020\u0076\u003d\u0025\u0073\u000a"+" \u0020\u0020\u0020\u0020\u0020\u0077\u003d\u0025\u0073\u000a", _cagf, _deaf, _fbfdd, _cagf && _deaf && _fbfdd, _eeed, _abgeb)
	}
	return _cagf && _deaf && _fbfdd
}

type wordBag struct {
	_bg.PdfRectangle
	_aced         float64
	_ecdb, _ebfag rulingList
	_cfb          float64
	_cccf         map[int][]*textWord
}

var _abe = false

func (_fadef lineRuling) asRuling() (*ruling, bool) {
	_bgee := ruling{_gcafd: _fadef._gadeg, Color: _fadef.Color, _cebe: _gffff}
	switch _fadef._gadeg {
	case _cege:
		_bgee._gbaff = _fadef.xMean()
		_bgee._fabe = _bb.Min(_fadef._abfed.Y, _fadef._fgcb.Y)
		_bgee._ccba = _bb.Max(_fadef._abfed.Y, _fadef._fgcb.Y)
	case _bfcbd:
		_bgee._gbaff = _fadef.yMean()
		_bgee._fabe = _bb.Min(_fadef._abfed.X, _fadef._fgcb.X)
		_bgee._ccba = _bb.Max(_fadef._abfed.X, _fadef._fgcb.X)
	default:
		_b.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _fadef._gadeg)
		return nil, false
	}
	return &_bgee, true
}
func _fccd(_fdg _bg.PdfRectangle) textState {
	return textState{_ffdc: 100, _dge: RenderModeFill, _bbd: _fdg}
}

// ExtractText processes and extracts all text data in content streams and returns as a string.
// It takes into account character encodings in the PDF file, which are decoded by
// CharcodeBytesToUnicode.
// Characters that can't be decoded are replaced with MissingCodeRune ('\ufffd' = �).
func (_fbfbg *Extractor) ExtractText() (string, error) {
	_febg, _, _, _cge := _fbfbg.ExtractTextWithStats()
	return _febg, _cge
}
func _dbbdd(_aebf int, _defa func(int, int) bool) []int {
	_dbdf := make([]int, _aebf)
	for _adffb := range _dbdf {
		_dbdf[_adffb] = _adffb
	}
	_gd.Slice(_dbdf, func(_fbcf, _fegdf int) bool { return _defa(_dbdf[_fbcf], _dbdf[_fegdf]) })
	return _dbdf
}
func (_adea rulingList) mergePrimary() float64 {
	_feffd := _adea[0]._gbaff
	for _, _ccgf := range _adea[1:] {
		_feffd += _ccgf._gbaff
	}
	return _feffd / float64(len(_adea))
}
func _bbe(_dbg []string, _gf int, _eb string) int {
	_bgc := _gf
	for ; _bgc < len(_dbg); _bgc++ {
		if _dbg[_bgc] != _eb {
			return _bgc
		}
	}
	return _bgc
}

// String returns a human readable description of `s`.
func (_dabc intSet) String() string {
	var _gcbdb []int
	for _dbged := range _dabc {
		if _dabc.has(_dbged) {
			_gcbdb = append(_gcbdb, _dbged)
		}
	}
	_gd.Ints(_gcbdb)
	return _d.Sprintf("\u0025\u002b\u0076", _gcbdb)
}
func (_aaab paraList) findTextTables() []*textTable {
	var _fbeaab []*textTable
	for _, _cdfgb := range _aaab {
		if _cdfgb.taken() || _cdfgb.Width() == 0 {
			continue
		}
		_bcecc := _cdfgb.isAtom()
		if _bcecc == nil {
			continue
		}
		_bcecc.growTable()
		if _bcecc._ddcega*_bcecc._egbbf < _cbea {
			continue
		}
		_bcecc.markCells()
		_bcecc.log("\u0067\u0072\u006fw\u006e")
		_fbeaab = append(_fbeaab, _bcecc)
	}
	return _fbeaab
}
func (_beeg *textTable) markCells() {
	for _ggffb := 0; _ggffb < _beeg._egbbf; _ggffb++ {
		for _dcadf := 0; _dcadf < _beeg._ddcega; _dcadf++ {
			_fdbcg := _beeg.get(_dcadf, _ggffb)
			if _fdbcg != nil {
				_fdbcg._bfaab = true
			}
		}
	}
}
func (_faea *textObject) reset() {
	_faea._gbcf = _ef.IdentityMatrix()
	_faea._gbe = _ef.IdentityMatrix()
	_faea._gaa = nil
}
func _dceb(_cccgd float64) bool { return _bb.Abs(_cccgd) < _egba }
func (_becfd *wordBag) depthIndexes() []int {
	if len(_becfd._cccf) == 0 {
		return nil
	}
	_dae := make([]int, len(_becfd._cccf))
	_bggd := 0
	for _baae := range _becfd._cccf {
		_dae[_bggd] = _baae
		_bggd++
	}
	_gd.Ints(_dae)
	return _dae
}
func (_gbfb *textObject) setHorizScaling(_ebg float64) {
	if _gbfb == nil {
		return
	}
	_gbfb._egf._ffdc = _ebg
}
func _gacaa(_fgegg float64) bool { return _bb.Abs(_fgegg) < _cfaf }
func (_ggcbg paraList) applyTables(_cadb []*textTable) paraList {
	var _dafaf paraList
	for _, _caegd := range _cadb {
		_dafaf = append(_dafaf, _caegd.newTablePara())
	}
	for _, _fffff := range _ggcbg {
		if _fffff._bfaab {
			continue
		}
		_dafaf = append(_dafaf, _fffff)
	}
	return _dafaf
}
func (_dacfc *shapesState) clearPath() {
	_dacfc._dgdd = nil
	_dacfc._aaba = false
	if _fbfd {
		_b.Log.Info("\u0043\u004c\u0045A\u0052\u003a\u0020\u0073\u0073\u003d\u0025\u0073", _dacfc)
	}
}
func (_adbc *subpath) close() {
	if !_fedaf(_adbc._agea[0], _adbc.last()) {
		_adbc.add(_adbc._agea[0])
	}
	_adbc._gdge = true
	_adbc.removeDuplicates()
}

// RenderMode specifies the text rendering mode (Tmode), which determines whether showing text shall cause
// glyph outlines to be  stroked, filled, used as a clipping boundary, or some combination of the three.
// Stroking, filling, and clipping shall have the same effects for a text object as they do for a path object
// (see 8.5.3, "Path-Painting Operators" and 8.5.4, "Clipping Path Operators").
type RenderMode int

func (_cbbg rulingList) secMinMax() (float64, float64) {
	_afea, _ffea := _cbbg[0]._fabe, _cbbg[0]._ccba
	for _, _dgcfc := range _cbbg[1:] {
		if _dgcfc._fabe < _afea {
			_afea = _dgcfc._fabe
		}
		if _dgcfc._ccba > _ffea {
			_ffea = _dgcfc._ccba
		}
	}
	return _afea, _ffea
}
func (_bded rulingList) isActualGrid() (rulingList, bool) {
	_dagd, _abae := _bded.augmentGrid()
	if !(len(_dagd) >= _fgbd+1 && len(_abae) >= _abfd+1) {
		if _gfab {
			_b.Log.Info("\u0069s\u0041\u0063t\u0075\u0061\u006c\u0047r\u0069\u0064\u003a \u004e\u006f\u0074\u0020\u0061\u006c\u0069\u0067\u006eed\u002e\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u003c\u0020\u0025d\u0020\u0078 \u0025\u0064", len(_dagd), len(_abae), _fgbd+1, _abfd+1)
		}
		return nil, false
	}
	if _gfab {
		_b.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075a\u006c\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025t\u0020\u0026\u0020\u0025\u0074\u0020\u2192 \u0025\u0074", _bded, len(_dagd) >= 2, len(_abae) >= 2, len(_dagd) >= 2 && len(_abae) >= 2)
		for _gcbfe, _ebdgg := range _bded {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0076\u000a", _gcbfe, _ebdgg)
		}
	}
	if _gbbfa {
		_adaf, _bgbf := _dagd[0], _dagd[len(_dagd)-1]
		_babg, _fedfe := _abae[0], _abae[len(_abae)-1]
		if !(_dceb(_adaf._gbaff-_babg._fabe) && _dceb(_bgbf._gbaff-_babg._ccba) && _dceb(_babg._gbaff-_adaf._ccba) && _dceb(_fedfe._gbaff-_adaf._fabe)) {
			if _gfab {
				_b.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069d\u003a\u0020\u0020N\u006f\u0074 \u0061\u006c\u0069\u0067\u006e\u0065d\u002e\n\t\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0076\u0031\u003d\u0025\u0073\u000a\u0009\u0068\u0030\u003d\u0025\u0073\u000a\u0009\u0068\u0031\u003d\u0025\u0073", _adaf, _bgbf, _babg, _fedfe)
			}
			return nil, false
		}
	} else {
		if !_dagd.aligned() {
			if _bffed {
				_b.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0076\u0065\u0072\u0074\u0073\u002e\u0020%\u0064", len(_dagd))
			}
			return nil, false
		}
		if !_abae.aligned() {
			if _gfab {
				_b.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0068\u006f\u0072\u007a\u0073\u002e\u0020%\u0064", len(_abae))
			}
			return nil, false
		}
	}
	_fged := append(_dagd, _abae...)
	return _fged, true
}
func _bdge(_cfgcb string) (string, bool) {
	_adcgf := []rune(_cfgcb)
	if len(_adcgf) != 1 {
		return "", false
	}
	_dbef, _gdce := _geacg[_adcgf[0]]
	return _dbef, _gdce
}

// Text gets the extracted text contained in `l`.
func (_ccfb *list) Text() string {
	_ceba := &_ge.Builder{}
	_gcaf := ""
	_ecaac(_ccfb, _ceba, &_gcaf)
	return _ceba.String()
}
func (_eggde *textTable) subdivide() *textTable {
	_eggde.logComposite("\u0073u\u0062\u0064\u0069\u0076\u0069\u0064e")
	_edbb := _eggde.compositeRowCorridors()
	_dbcba := _eggde.compositeColCorridors()
	if _eeeg {
		_b.Log.Info("\u0073u\u0062\u0064i\u0076\u0069\u0064\u0065:\u000a\u0009\u0072o\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s=\u0025\u0073\u000a\t\u0063\u006fl\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073=\u0025\u0073", _faebg(_edbb), _faebg(_dbcba))
	}
	if len(_edbb) == 0 || len(_dbcba) == 0 {
		return _eggde
	}
	_dfgba(_edbb)
	_dfgba(_dbcba)
	if _eeeg {
		_b.Log.Info("\u0073\u0075\u0062\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0066\u0069\u0078\u0065\u0064\u003a\u000a\u0009r\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073\u003d\u0025\u0073\u000a\u0009\u0063\u006f\u006cC\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d\u0025\u0073", _faebg(_edbb), _faebg(_dbcba))
	}
	_fgebga, _gbcd := _ddgea(_eggde._egbbf, _edbb)
	_deea, _bggba := _ddgea(_eggde._ddcega, _dbcba)
	_degfca := make(map[uint64]*textPara, _bggba*_gbcd)
	_edecd := &textTable{PdfRectangle: _eggde.PdfRectangle, _gfagc: _eggde._gfagc, _egbbf: _gbcd, _ddcega: _bggba, _bcec: _degfca}
	if _eeeg {
		_b.Log.Info("\u0073\u0075b\u0064\u0069\u0076\u0069\u0064\u0065\u003a\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0063\u0065\u006c\u006c\u0073\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0079\u004f\u0066\u0066\u0073\u0065\u0074\u0073=\u0025\u002b\u0076\u000a"+"\u0009\u0078\u004f\u0066\u0066\u0073\u0065\u0074\u0073\u003d\u0025\u002b\u0076", _eggde._ddcega, _eggde._egbbf, _bggba, _gbcd, _faebg(_edbb), _faebg(_dbcba), _fgebga, _deea)
	}
	for _egec := 0; _egec < _eggde._egbbf; _egec++ {
		_afgag := _fgebga[_egec]
		for _dada := 0; _dada < _eggde._ddcega; _dada++ {
			_dafdg := _deea[_dada]
			if _eeeg {
				_d.Printf("\u0025\u0036\u0064\u002c %\u0032\u0064\u003a\u0020\u0078\u0030\u003d\u0025\u0064\u0020\u0079\u0030\u003d\u0025d\u000a", _dada, _egec, _dafdg, _afgag)
			}
			_bfecg, _aeeff := _eggde._fccfa[_caabd(_dada, _egec)]
			if !_aeeff {
				continue
			}
			_daca := _bfecg.split(_edbb[_egec], _dbcba[_dada])
			for _bdeae := 0; _bdeae < _daca._egbbf; _bdeae++ {
				for _aeadd := 0; _aeadd < _daca._ddcega; _aeadd++ {
					_bgbc := _daca.get(_aeadd, _bdeae)
					_edecd.put(_dafdg+_aeadd, _afgag+_bdeae, _bgbc)
					if _eeeg {
						_d.Printf("\u0025\u0038\u0064\u002c\u0020\u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _dafdg+_aeadd, _afgag+_bdeae, _bgbc)
					}
				}
			}
		}
	}
	return _edecd
}
func _cgdb(_efbeb []float64, _bbgb, _gecff float64) []float64 {
	_gfefc, _ccgff := _bbgb, _gecff
	if _ccgff < _gfefc {
		_gfefc, _ccgff = _ccgff, _gfefc
	}
	_bcfeg := make([]float64, 0, len(_efbeb)+2)
	_bcfeg = append(_bcfeg, _bbgb)
	for _, _ecfbd := range _efbeb {
		if _ecfbd <= _gfefc {
			continue
		} else if _ecfbd >= _ccgff {
			break
		}
		_bcfeg = append(_bcfeg, _ecfbd)
	}
	_bcfeg = append(_bcfeg, _gecff)
	return _bcfeg
}
func (_afed paraList) inTile(_dcea gridTile) paraList {
	var _effd paraList
	for _, _gddga := range _afed {
		if _dcea.contains(_gddga.PdfRectangle) {
			_effd = append(_effd, _gddga)
		}
	}
	if _eeeg {
		_d.Printf("\u0020 \u0020\u0069\u006e\u0054i\u006c\u0065\u003a\u0020\u0020%\u0073 \u0069n\u0073\u0069\u0064\u0065\u003d\u0025\u0064\n", _dcea, len(_effd))
		for _fdgda, _ddca := range _effd {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _fdgda, _ddca)
		}
		_d.Println("")
	}
	return _effd
}

type event struct {
	_abbcd  float64
	_dcgeb  bool
	_bbeaga int
}

// PageTextOptions holds various options available in extraction process.
type PageTextOptions struct {
	_abdc bool
	_ddcf bool
}

func (_eedcc *textTable) emptyCompositeRow(_adfe int) bool {
	for _edag := 0; _edag < _eedcc._ddcega; _edag++ {
		if _deag, _bdga := _eedcc._fccfa[_caabd(_edag, _adfe)]; _bdga {
			if len(_deag.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_abda *subpath) makeRectRuling(_gdga _ab.Color) (*ruling, bool) {
	if _bfgg {
		_b.Log.Info("\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0070\u0061\u0074\u0068\u003d\u0025\u0076", _abda)
	}
	_fecgc := _abda._agea[:4]
	_agbdf := make(map[int]rulingKind, len(_fecgc))
	for _ddggc, _bcee := range _fecgc {
		_fagg := _abda._agea[(_ddggc+1)%4]
		_agbdf[_ddggc] = _caa(_bcee, _fagg)
		if _bfgg {
			_d.Printf("\u0025\u0034\u0064: \u0025\u0073\u0020\u003d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066", _ddggc, _agbdf[_ddggc], _bcee, _fagg)
		}
	}
	if _bfgg {
		_d.Printf("\u0020\u0020\u0020\u006b\u0069\u006e\u0064\u0073\u003d\u0025\u002b\u0076\u000a", _agbdf)
	}
	var _aegaf, _fdfe []int
	for _agfd, _gbefc := range _agbdf {
		switch _gbefc {
		case _bfcbd:
			_fdfe = append(_fdfe, _agfd)
		case _cege:
			_aegaf = append(_aegaf, _agfd)
		}
	}
	if _bfgg {
		_d.Printf("\u0020\u0020 \u0068\u006f\u0072z\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_fdfe), _fdfe)
		_d.Printf("\u0020\u0020 \u0076\u0065\u0072t\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_aegaf), _aegaf)
	}
	_dbcb := (len(_fdfe) == 2 && len(_aegaf) == 2) || (len(_fdfe) == 2 && len(_aegaf) == 0 && _bdef(_fecgc[_fdfe[0]], _fecgc[_fdfe[1]])) || (len(_aegaf) == 2 && len(_fdfe) == 0 && _dgcb(_fecgc[_aegaf[0]], _fecgc[_aegaf[1]]))
	if _bfgg {
		_d.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_fdfe), len(_aegaf), _dbcb)
	}
	if !_dbcb {
		if _bfgg {
			_b.Log.Error("\u0021!\u006d\u0061\u006b\u0065R\u0065\u0063\u0074\u0052\u0075l\u0069n\u0067:\u0020\u0070\u0061\u0074\u0068\u003d\u0025v", _abda)
			_d.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_fdfe), len(_aegaf), _dbcb)
		}
		return &ruling{}, false
	}
	if len(_aegaf) == 0 {
		for _acbg, _fcba := range _agbdf {
			if _fcba != _bfcbd {
				_aegaf = append(_aegaf, _acbg)
			}
		}
	}
	if len(_fdfe) == 0 {
		for _afdfg, _ffdeg := range _agbdf {
			if _ffdeg != _cege {
				_fdfe = append(_fdfe, _afdfg)
			}
		}
	}
	if _bfgg {
		_b.Log.Info("\u006da\u006b\u0065R\u0065\u0063\u0074\u0052u\u006c\u0069\u006eg\u003a\u0020\u0068\u006f\u0072\u007a\u0073\u003d\u0025d \u0076\u0065\u0072t\u0073\u003d%\u0064\u0020\u0070\u006f\u0069\u006et\u0073\u003d%\u0064\u000a"+"\u0009\u0020\u0068o\u0072\u007a\u0073\u003d\u0025\u002b\u0076\u000a"+"\u0009\u0020\u0076e\u0072\u0074\u0073\u003d\u0025\u002b\u0076\u000a"+"\t\u0070\u006f\u0069\u006e\u0074\u0073\u003d\u0025\u002b\u0076", len(_fdfe), len(_aegaf), len(_fecgc), _fdfe, _aegaf, _fecgc)
	}
	var _dcbb, _acdaa, _cgaa, _facb _ef.Point
	if _fecgc[_fdfe[0]].Y > _fecgc[_fdfe[1]].Y {
		_cgaa, _facb = _fecgc[_fdfe[0]], _fecgc[_fdfe[1]]
	} else {
		_cgaa, _facb = _fecgc[_fdfe[1]], _fecgc[_fdfe[0]]
	}
	if _fecgc[_aegaf[0]].X > _fecgc[_aegaf[1]].X {
		_dcbb, _acdaa = _fecgc[_aegaf[0]], _fecgc[_aegaf[1]]
	} else {
		_dcbb, _acdaa = _fecgc[_aegaf[1]], _fecgc[_aegaf[0]]
	}
	_gbgd := _bg.PdfRectangle{Llx: _dcbb.X, Urx: _acdaa.X, Lly: _facb.Y, Ury: _cgaa.Y}
	if _gbgd.Llx > _gbgd.Urx {
		_gbgd.Llx, _gbgd.Urx = _gbgd.Urx, _gbgd.Llx
	}
	if _gbgd.Lly > _gbgd.Ury {
		_gbgd.Lly, _gbgd.Ury = _gbgd.Ury, _gbgd.Lly
	}
	_ffdab := rectRuling{PdfRectangle: _gbgd, _acad: _cccae(_gbgd), Color: _gdga}
	if _ffdab._acad == _bgeag {
		if _bfgg {
			_b.Log.Error("\u006da\u006b\u0065\u0052\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006eg\u003a\u0020\u006b\u0069\u006e\u0064\u003d\u006e\u0069\u006c")
		}
		return nil, false
	}
	_gcded, _afbd := _ffdab.asRuling()
	if !_afbd {
		if _bfgg {
			_b.Log.Error("\u006da\u006b\u0065\u0052\u0065c\u0074\u0052\u0075\u006c\u0069n\u0067:\u0020!\u0069\u0073\u0052\u0075\u006c\u0069\u006eg")
		}
		return nil, false
	}
	if _gfab {
		_d.Printf("\u0020\u0020\u0020\u0072\u003d\u0025\u0073\u000a", _gcded.String())
	}
	return _gcded, true
}
func (_faa *textObject) getFontDirect(_dgeb string) (*_bg.PdfFont, error) {
	_agfb, _fefa := _faa.getFontDict(_dgeb)
	if _fefa != nil {
		return nil, _fefa
	}
	_dcff, _fefa := _bg.NewPdfFontFromPdfObject(_agfb)
	if _fefa != nil {
		_b.Log.Debug("\u0067\u0065\u0074\u0046\u006f\u006e\u0074\u0044\u0069\u0072\u0065\u0063\u0074\u003a\u0020\u004e\u0065\u0077Pd\u0066F\u006f\u006e\u0074\u0046\u0072\u006f\u006d\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d%\u0023\u0071\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgeb, _fefa)
	}
	return _dcff, _fefa
}
func (_deec *textObject) setTextRise(_fbb float64) {
	if _deec == nil {
		return
	}
	_deec._egf._dbca = _fbb
}
func _fedaf(_gbbe, _dcebd _ef.Point) bool { return _gbbe.X == _dcebd.X && _gbbe.Y == _dcebd.Y }

type rectRuling struct {
	_acad rulingKind
	_dagc markKind
	_ab.Color
	_bg.PdfRectangle
}

// String returns a description of `l`.
func (_ddgf *textLine) String() string {
	return _d.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _ddgf._ggfc, _ddgf.PdfRectangle, _ddgf._gdbd, _ddgf.text())
}
func _caa(_decd, _abag _ef.Point) rulingKind {
	_cdgd := _bb.Abs(_decd.X - _abag.X)
	_gaddb := _bb.Abs(_decd.Y - _abag.Y)
	return _fgebf(_cdgd, _gaddb, _baaee)
}
func (_cebc *textPara) depth() float64 {
	if _cebc._ecee {
		return -1.0
	}
	if len(_cebc._gdc) > 0 {
		return _cebc._gdc[0]._ggfc
	}
	return _cebc._cece.depth()
}

// String returns a string describing `tm`.
func (_ddd TextMark) String() string {
	_fbgc := _ddd.BBox
	var _ggda string
	if _ddd.Font != nil {
		_ggda = _ddd.Font.String()
		if len(_ggda) > 50 {
			_ggda = _ggda[:50] + "\u002e\u002e\u002e"
		}
	}
	var _eab string
	if _ddd.Meta {
		_eab = "\u0020\u002a\u004d\u002a"
	}
	return _d.Sprintf("\u007b\u0054\u0065\u0078t\u004d\u0061\u0072\u006b\u003a\u0020\u0025\u0064\u0020%\u0071\u003d\u0025\u0030\u0032\u0078\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e2\u0066\u0029\u0020\u0028\u00256\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0025\u0073\u007d", _ddd.Offset, _ddd.Text, []rune(_ddd.Text), _fbgc.Llx, _fbgc.Lly, _fbgc.Urx, _fbgc.Ury, _ggda, _eab)
}
func (_bbeda paraList) llyRange(_fgef []int, _bfbf, _gcgg float64) []int {
	_debe := len(_bbeda)
	if _gcgg < _bbeda[_fgef[0]].Lly || _bfbf > _bbeda[_fgef[_debe-1]].Lly {
		return nil
	}
	_aefe := _gd.Search(_debe, func(_ceegd int) bool { return _bbeda[_fgef[_ceegd]].Lly >= _bfbf })
	_bebd := _gd.Search(_debe, func(_agbf int) bool { return _bbeda[_fgef[_agbf]].Lly > _gcgg })
	return _fgef[_aefe:_bebd]
}

// ImageExtractOptions contains options for controlling image extraction from
// PDF pages.
type ImageExtractOptions struct{ IncludeInlineStencilMasks bool }

func (_adfgf *textPara) toCellTextMarks(_gffed *int) []TextMark {
	var _gbgef []TextMark
	for _gccf, _ebfc := range _adfgf._gdc {
		_eagdb := _ebfc.toTextMarks(_gffed)
		_fedfc := _degfc && _ebfc.endsInHyphen() && _gccf != len(_adfgf._gdc)-1
		if _fedfc {
			_eagdb = _cdaae(_eagdb, _gffed)
		}
		_gbgef = append(_gbgef, _eagdb...)
		if !(_fedfc || _gccf == len(_adfgf._gdc)-1) {
			_gbgef = _eaagb(_gbgef, _gffed, _afbff(_ebfc._ggfc, _adfgf._gdc[_gccf+1]._ggfc))
		}
	}
	return _gbgef
}
func (_afeg *textTable) reduceTiling(_bfca gridTiling, _cfdec float64) *textTable {
	_dgcd := make([]int, 0, _afeg._egbbf)
	_dcdaed := make([]int, 0, _afeg._ddcega)
	_gbff := _bfca._cdfg
	_gbcde := _bfca._fgcd
	for _ecdg := 0; _ecdg < _afeg._egbbf; _ecdg++ {
		_bbgcf := _ecdg > 0 && _bb.Abs(_gbcde[_ecdg-1]-_gbcde[_ecdg]) < _cfdec && _afeg.emptyCompositeRow(_ecdg)
		if !_bbgcf {
			_dgcd = append(_dgcd, _ecdg)
		}
	}
	for _bedfc := 0; _bedfc < _afeg._ddcega; _bedfc++ {
		_bdeab := _bedfc < _afeg._ddcega-1 && _bb.Abs(_gbff[_bedfc+1]-_gbff[_bedfc]) < _cfdec && _afeg.emptyCompositeColumn(_bedfc)
		if !_bdeab {
			_dcdaed = append(_dcdaed, _bedfc)
		}
	}
	if len(_dgcd) == _afeg._egbbf && len(_dcdaed) == _afeg._ddcega {
		return _afeg
	}
	_decae := textTable{_gfagc: _afeg._gfagc, _ddcega: len(_dcdaed), _egbbf: len(_dgcd), _fccfa: make(map[uint64]compositeCell, len(_dcdaed)*len(_dgcd))}
	if _eeeg {
		_b.Log.Info("\u0072\u0065\u0064\u0075c\u0065\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0025d\u0078%\u0064\u0020\u002d\u003e\u0020\u0025\u0064x\u0025\u0064", _afeg._ddcega, _afeg._egbbf, len(_dcdaed), len(_dgcd))
		_b.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _dcdaed)
		_b.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _dgcd)
	}
	for _gbfga, _efaba := range _dgcd {
		for _gffgf, _ffafd := range _dcdaed {
			_affbb, _becb := _afeg.getComposite(_ffafd, _efaba)
			if len(_affbb) == 0 {
				continue
			}
			if _eeeg {
				_d.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _gffgf, _gbfga, _ffafd, _efaba, _cfbg(_affbb.merge().text(), 50))
			}
			_decae.putComposite(_gffgf, _gbfga, _affbb, _becb)
		}
	}
	return &_decae
}
func _gafde(_cabaf _bg.PdfRectangle, _fffg bounded) float64 { return _cabaf.Ury - _fffg.bbox().Lly }
func _ecaac(_bcaef *list, _bbed *_ge.Builder, _cbef *string) {
	_cddc := _efee(_bcaef, _cbef)
	_bbed.WriteString(_cddc)
	for _, _ggfe := range _bcaef._bdfaa {
		_dabe := *_cbef + "\u0020\u0020\u0020"
		_ecaac(_ggfe, _bbed, &_dabe)
	}
}
func _fbbag(_cgdbg _bg.PdfColorspace, _affa _bg.PdfColor) _ab.Color {
	if _cgdbg == nil || _affa == nil {
		return _ab.Black
	}
	_adfbe, _effe := _cgdbg.ColorToRGB(_affa)
	if _effe != nil {
		_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006fu\u006c\u0064\u0020no\u0074\u0020\u0063\u006f\u006e\u0076e\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0025\u0076\u0020\u0028\u0025\u0076)\u0020\u0074\u006f\u0020\u0052\u0047\u0042\u003a \u0025\u0073", _affa, _cgdbg, _effe)
		return _ab.Black
	}
	_bgdb, _fbfbb := _adfbe.(*_bg.PdfColorDeviceRGB)
	if !_fbfbb {
		_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0052\u0047\u0042\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0076", _adfbe)
		return _ab.Black
	}
	return _ab.NRGBA{R: uint8(_bgdb.R() * 255), G: uint8(_bgdb.G() * 255), B: uint8(_bgdb.B() * 255), A: uint8(255)}
}
func _daac(_acfg []TextMark, _fdcf *TextTable) []TextMark {
	var _ddef []TextMark
	for _, _cebb := range _acfg {
		_cebb._fafb = true
		_cebb._aeaf = _fdcf
		_ddef = append(_ddef, _cebb)
	}
	return _ddef
}
func (_cgae *shapesState) fill(_ebbe *[]pathSection) {
	_ebag := pathSection{_cdc: _cgae._dgdd, Color: _cgae._eadd.getFillColor()}
	*_ebbe = append(*_ebbe, _ebag)
	if _gfab {
		_agcb := _ebag.bbox()
		_d.Printf("\u0020 \u0020\u0020\u0046\u0049\u004c\u004c\u003a %\u0032\u0064\u0020\u0066\u0069\u006c\u006c\u0073\u0020\u0028\u0025\u0064\u0020\u006ee\u0077\u0029 \u0073\u0073\u003d%\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d\u0025\u0033\u0076\u0020\u0025\u0036\u002e\u0032f\u003d\u00256.\u0032\u0066\u0078%\u0036\u002e\u0032\u0066\u000a", len(*_ebbe), len(_ebag._cdc), _cgae, _ebag.Color, _agcb, _agcb.Width(), _agcb.Height())
		if _fafd {
			for _bcbe, _gccb := range _ebag._cdc {
				_d.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _bcbe, _gccb)
				if _bcbe == 10 {
					break
				}
			}
		}
	}
}

const (
	RenderModeStroke RenderMode = 1 << iota
	RenderModeFill
	RenderModeClip
)
const (
	_cfaf  = 1.0e-6
	_bagc  = 1.0e-4
	_efdfe = 10
	_dfbad = 6
	_gfgc  = 0.5
	_gebc  = 0.12
	_fbadd = 0.19
	_gadf  = 0.04
	_cgga  = 0.04
	_ccdef = 1.0
	_abffb = 0.04
	_gccbd = 12
	_cfdb  = 0.4
	_dbgb  = 0.7
	_ffed  = 1.0
	_dcda  = 0.1
	_bfcc  = 1.4
	_bggg  = 0.46
	_gcca  = 0.02
	_ebee  = 0.2
	_dgee  = 0.5
	_cfgcc = 4
	_debb  = 4.0
	_cbea  = 6
	_ddcb  = 0.3
	_eead  = 0.01
	_fgfcb = 0.02
	_fgbd  = 2
	_abfd  = 2
	_bbag  = 500
	_feef  = 4.0
	_dbcc  = 4.0
	_baaee = 0.05
	_cbab  = 0.1
	_beac  = 2.0
	_egba  = 2.0
	_edef  = 1.5
	_dbae  = 3.0
	_ddee  = 0.25
)

func _dgge(_bfga _bg.PdfRectangle) *ruling {
	return &ruling{_gcafd: _bfcbd, _gbaff: _bfga.Lly, _fabe: _bfga.Llx, _ccba: _bfga.Urx}
}
func _fgfcbe(_ffbc *textLine, _gddg []*textLine, _baaf []float64) float64 {
	var _febeg float64 = -1
	for _, _ffcba := range _gddg {
		if _ffcba._ggfc > _ffbc._ggfc {
			if _bb.Round(_ffcba.Llx) >= _bb.Round(_ffbc.Llx) {
				_febeg = _ffcba._ggfc
			} else {
				break
			}
		}
	}
	return _febeg
}
func (_bddaf paraList) writeText(_gcdeb _e.Writer) {
	for _gegcg, _bfaa := range _bddaf {
		if _bfaa._ecee {
			continue
		}
		_bfaa.writeText(_gcdeb)
		if _gegcg != len(_bddaf)-1 {
			if _baagc(_bfaa, _bddaf[_gegcg+1]) {
				_gcdeb.Write([]byte("\u0020"))
			} else {
				_gcdeb.Write([]byte("\u000a"))
				_gcdeb.Write([]byte("\u000a"))
			}
		}
	}
	_gcdeb.Write([]byte("\u000a"))
	_gcdeb.Write([]byte("\u000a"))
}
func _fefe(_eafd _bg.PdfRectangle, _fgcf, _gegfd, _dbdbe, _gebf *ruling) gridTile {
	_ffddg := _eafd.Llx
	_eafef := _eafd.Urx
	_cefg := _eafd.Lly
	_eaeb := _eafd.Ury
	return gridTile{PdfRectangle: _eafd, _acef: _fgcf != nil && _fgcf.encloses(_cefg, _eaeb), _bgdga: _gegfd != nil && _gegfd.encloses(_cefg, _eaeb), _dbbba: _dbdbe != nil && _dbdbe.encloses(_ffddg, _eafef), _acffg: _gebf != nil && _gebf.encloses(_ffddg, _eafef)}
}
func (_fbef paraList) sortTopoOrder()          { _cacbb := _fbef.topoOrder(); _fbef.reorder(_cacbb) }
func (_caee *textWord) bbox() _bg.PdfRectangle { return _caee.PdfRectangle }

// Font represents the font properties on a PDF page.
type Font struct {
	PdfFont *_bg.PdfFont

	// FontName represents Font Name from font properties.
	FontName string

	// FontType represents Font Subtype entry in the font dictionary inside page resources.
	// Examples : type0, Type1, MMType1, Type3, TrueType, CIDFont.
	FontType string

	// ToUnicode is true if font provides a `ToUnicode` mapping.
	ToUnicode bool

	// IsCID is true if underlying font is a composite font.
	// Composite font is represented by a font dictionary whose Subtype is `Type0`
	IsCID bool

	// IsSimple is true if font is simple font.
	// A simple font is limited to only 8 bit (255) character codes.
	IsSimple bool

	// FontData represents the raw data of the embedded font file.
	// It can have format TrueType (TTF), PostScript Font (PFB) or Compact Font Format (CCF).
	// FontData value can be indicates from `FontFile`, `FontFile2` or `FontFile3` inside Font Descriptor.
	// At most, only one of `FontFile`, `FontFile2` or `FontFile3` will be FontData value.
	FontData []byte

	// FontFileName is a name representing the font. it has format:
	// (Font Name) + (Font Type Extension), example: helvetica.ttf.
	FontFileName string

	// FontDescriptor represents metrics and other attributes inside font properties from PDF Structure (Font Descriptor).
	FontDescriptor *_bg.PdfFontDescriptor
}

func (_cgfee rulingList) toTilings() (rulingList, []gridTiling) {
	_cgfee.log("\u0074o\u0054\u0069\u006c\u0069\u006e\u0067s")
	if len(_cgfee) == 0 {
		return nil, nil
	}
	_cgfee = _cgfee.tidied("\u0061\u006c\u006c")
	_cgfee.log("\u0074\u0069\u0064\u0069\u0065\u0064")
	_effce := _cgfee.toGrids()
	_dadfb := make([]gridTiling, len(_effce))
	for _bbcb, _bcbb := range _effce {
		_dadfb[_bbcb] = _bcbb.asTiling()
	}
	return _cgfee, _dadfb
}
func (_cgdcg *wordBag) highestWord(_adad int, _adag, _fcd float64) *textWord {
	for _, _dbbfc := range _cgdcg._cccf[_adad] {
		if _adag <= _dbbfc._ccee && _dbbfc._ccee <= _fcd {
			return _dbbfc
		}
	}
	return nil
}
func (_dfdfg compositeCell) split(_gccc, _dcfg []float64) *textTable {
	_bbcd := len(_gccc) + 1
	_bggge := len(_dcfg) + 1
	if _eeeg {
		_b.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0043\u0065l\u006c\u002e\u0073\u0070l\u0069\u0074\u003a\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a\u0009\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025\u0073\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073=\u0025\u0036\u002e\u0032\u0066\u000a\t\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d%\u0036\u002e\u0032\u0066", _bggge, _bbcd, _dfdfg, _gccc, _dcfg)
		_d.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073\u000a", len(_dfdfg.paraList))
		for _fcdfb, _fdcg := range _dfdfg.paraList {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _fcdfb, _fdcg.String())
		}
		_d.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", len(_dfdfg.lines()))
		for _eggg, _efbg := range _dfdfg.lines() {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _eggg, _efbg)
		}
	}
	_gccc = _cgdb(_gccc, _dfdfg.Ury, _dfdfg.Lly)
	_dcfg = _cgdb(_dcfg, _dfdfg.Llx, _dfdfg.Urx)
	_afcc := make(map[uint64]*textPara, _bggge*_bbcd)
	_dgada := textTable{_ddcega: _bggge, _egbbf: _bbcd, _bcec: _afcc}
	_dcade := _dfdfg.paraList
	_gd.Slice(_dcade, func(_gfae, _abed int) bool {
		_fafa, _eacb := _dcade[_gfae], _dcade[_abed]
		_dgef, _aegdg := _fafa.Lly, _eacb.Lly
		if _dgef != _aegdg {
			return _dgef < _aegdg
		}
		return _fafa.Llx < _eacb.Llx
	})
	_acba := make(map[uint64]_bg.PdfRectangle, _bggge*_bbcd)
	for _gdgfb, _aabg := range _gccc[1:] {
		_faab := _gccc[_gdgfb]
		for _bedbg, _eddc := range _dcfg[1:] {
			_cedag := _dcfg[_bedbg]
			_acba[_caabd(_bedbg, _gdgfb)] = _bg.PdfRectangle{Llx: _cedag, Urx: _eddc, Lly: _aabg, Ury: _faab}
		}
	}
	if _eeeg {
		_b.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0043\u0065l\u006c\u002e\u0073\u0070\u006c\u0069\u0074\u003a\u0020\u0072e\u0063\u0074\u0073")
		_d.Printf("\u0020\u0020\u0020\u0020")
		for _cccb := 0; _cccb < _bggge; _cccb++ {
			_d.Printf("\u0025\u0033\u0030\u0064\u002c\u0020", _cccb)
		}
		_d.Println()
		for _bdec := 0; _bdec < _bbcd; _bdec++ {
			_d.Printf("\u0020\u0020\u0025\u0032\u0064\u003a", _bdec)
			for _gdff := 0; _gdff < _bggge; _gdff++ {
				_d.Printf("\u00256\u002e\u0032\u0066\u002c\u0020", _acba[_caabd(_gdff, _bdec)])
			}
			_d.Println()
		}
	}
	_cbbe := func(_ggce *textLine) (int, int) {
		for _cede := 0; _cede < _bbcd; _cede++ {
			for _fcdcd := 0; _fcdcd < _bggge; _fcdcd++ {
				if _ffff(_acba[_caabd(_fcdcd, _cede)], _ggce.PdfRectangle) {
					return _fcdcd, _cede
				}
			}
		}
		return -1, -1
	}
	_cfbc := make(map[uint64][]*textLine, _bggge*_bbcd)
	for _, _cded := range _dcade.lines() {
		_fdab, _eada := _cbbe(_cded)
		if _fdab < 0 {
			continue
		}
		_cfbc[_caabd(_fdab, _eada)] = append(_cfbc[_caabd(_fdab, _eada)], _cded)
	}
	for _fegb := 0; _fegb < len(_gccc)-1; _fegb++ {
		_eedbc := _gccc[_fegb]
		_cdbgc := _gccc[_fegb+1]
		for _beeed := 0; _beeed < len(_dcfg)-1; _beeed++ {
			_ddcd := _dcfg[_beeed]
			_cdeg := _dcfg[_beeed+1]
			_gfadg := _bg.PdfRectangle{Llx: _ddcd, Urx: _cdeg, Lly: _cdbgc, Ury: _eedbc}
			_cdfc := _cfbc[_caabd(_beeed, _fegb)]
			if len(_cdfc) == 0 {
				continue
			}
			_bdac := _ebeeg(_gfadg, _cdfc)
			_dgada.put(_beeed, _fegb, _bdac)
		}
	}
	return &_dgada
}
func (_gfaf *textWord) addDiacritic(_gdgb string) {
	_gcfg := _gfaf._ecdf[len(_gfaf._ecdf)-1]
	_gcfg._aaag += _gdgb
	_gcfg._aaag = _g.NFKC.String(_gcfg._aaag)
}
func _bgffa(_befe map[int][]float64) []int {
	_fdcbg := make([]int, len(_befe))
	_fgeff := 0
	for _bbcbb := range _befe {
		_fdcbg[_fgeff] = _bbcbb
		_fgeff++
	}
	_gd.Ints(_fdcbg)
	return _fdcbg
}
func (_baef *stateStack) push(_gbce *textState)       { _bfa := *_gbce; *_baef = append(*_baef, &_bfa) }
func (_eage *wordBag) firstWord(_geaea int) *textWord { return _eage._cccf[_geaea][0] }
func (_ecbg paraList) toTextMarks() []TextMark {
	_cffc := 0
	var _fdgba []TextMark
	for _fdgf, _ggfec := range _ecbg {
		if _ggfec._ecee {
			continue
		}
		_ebgf := _ggfec.toTextMarks(&_cffc)
		_fdgba = append(_fdgba, _ebgf...)
		if _fdgf != len(_ecbg)-1 {
			if _baagc(_ggfec, _ecbg[_fdgf+1]) {
				_fdgba = _eaagb(_fdgba, &_cffc, "\u0020")
			} else {
				_fdgba = _eaagb(_fdgba, &_cffc, "\u000a")
				_fdgba = _eaagb(_fdgba, &_cffc, "\u000a")
			}
		}
	}
	_fdgba = _eaagb(_fdgba, &_cffc, "\u000a")
	_fdgba = _eaagb(_fdgba, &_cffc, "\u000a")
	return _fdgba
}
func _degb(_faaf string) bool {
	for _, _agad := range _faaf {
		if !_ae.IsSpace(_agad) {
			return false
		}
	}
	return true
}

// ExtractFonts returns all font information from the page extractor, including
// font name, font type, the raw data of the embedded font file (if embedded), font descriptor and more.
//
// The argument `previousPageFonts` is used when trying to build a complete font catalog for multiple pages or the entire document.
// The entries from `previousPageFonts` are added to the returned result unless already included in the page, i.e. no duplicate entries.
//
// NOTE: If previousPageFonts is nil, all fonts from the page will be returned. Use it when building up a full list of fonts for a document or page range.
func (_dgb *Extractor) ExtractFonts(previousPageFonts *PageFonts) (*PageFonts, error) {
	_gdb := PageFonts{}
	_cbe := _gdb.extractPageResourcesToFont(_dgb._feg)
	if _cbe != nil {
		return nil, _cbe
	}
	if previousPageFonts != nil {
		for _, _fbc := range previousPageFonts.Fonts {
			if !_fbfb(_gdb.Fonts, _fbc.FontName) {
				_gdb.Fonts = append(_gdb.Fonts, _fbc)
			}
		}
	}
	return &PageFonts{Fonts: _gdb.Fonts}, nil
}
func (_bgeaf rulingList) log(_cgdd string) {
	if !_gfab {
		return
	}
	_b.Log.Info("\u0023\u0023\u0023\u0020\u0025\u0031\u0030\u0073\u003a\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _cgdd, _bgeaf.String())
	for _dcace, _daaf := range _bgeaf {
		_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dcace, _daaf.String())
	}
}
func (_ebcc paraList) xNeighbours(_gcbce float64) map[*textPara][]int {
	_ecggg := make([]event, 2*len(_ebcc))
	if _gcbce == 0 {
		for _bfcca, _gafef := range _ebcc {
			_ecggg[2*_bfcca] = event{_gafef.Llx, true, _bfcca}
			_ecggg[2*_bfcca+1] = event{_gafef.Urx, false, _bfcca}
		}
	} else {
		for _cbcgb, _bgge := range _ebcc {
			_ecggg[2*_cbcgb] = event{_bgge.Llx - _gcbce*_bgge.fontsize(), true, _cbcgb}
			_ecggg[2*_cbcgb+1] = event{_bgge.Urx + _gcbce*_bgge.fontsize(), false, _cbcgb}
		}
	}
	return _ebcc.eventNeighbours(_ecggg)
}
func (_abfe *shapesState) lastpointEstablished() (_ef.Point, bool) {
	if _abfe._aaba {
		return _abfe._gbgg, false
	}
	_gafe := len(_abfe._dgdd)
	if _gafe > 0 && _abfe._dgdd[_gafe-1]._gdge {
		return _abfe._dgdd[_gafe-1].last(), false
	}
	return _ef.Point{}, true
}

type structTreeRoot struct {
	_dgeg []structElement
	_dgba string
}

var _gdac = _gb.MustCompile("\u005e\u005c\u0073\u002a\u0028\u005c\u0064\u002b\u005c\u002e\u003f|\u005b\u0049\u0069\u0076\u005d\u002b\u0029\u005c\u0073\u002a\\\u0029\u003f\u0024")

func (_daeee *textWord) computeText() string {
	_dfdeg := make([]string, len(_daeee._ecdf))
	for _bgdcf, _acbgc := range _daeee._ecdf {
		_dfdeg[_bgdcf] = _acbgc._aaag
	}
	return _ge.Join(_dfdeg, "")
}
func (_efdfbg *structTreeRoot) buildList(_ffded map[int][]*textLine, _cde _gbc.PdfObject) []*list {
	if _efdfbg == nil {
		_b.Log.Debug("\u0062\u0075\u0069\u006c\u0064\u004c\u0069\u0073\u0074\u003a\u0020t\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return nil
	}
	var _egfaf *structElement
	_gdbg := []structElement{}
	if len(_efdfbg._dgeg) == 1 {
		_bdcb := _efdfbg._dgeg[0]._adfb
		if _bdcb == "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074" || _bdcb == "\u0053\u0065\u0063\u0074" || _bdcb == "\u0050\u0061\u0072\u0074" || _bdcb == "\u0044\u0069\u0076" || _bdcb == "\u0041\u0072\u0074" {
			_egfaf = &_efdfbg._dgeg[0]
		}
	} else {
		_egfaf = &structElement{_fbcba: _efdfbg._dgeg, _adfb: _efdfbg._dgba}
	}
	if _egfaf == nil {
		_b.Log.Debug("\u0062\u0075\u0069\u006cd\u004c\u0069\u0073\u0074\u003a\u0020\u0074\u006f\u0070\u0045l\u0065m\u0065\u006e\u0074\u0020\u0069\u0073\u0020n\u0069\u006c")
		return nil
	}
	for _, _efgc := range _egfaf._fbcba {
		if _efgc._adfb == "\u004c" {
			_gdbg = append(_gdbg, _efgc)
		} else if _efgc._adfb == "\u0054\u0061\u0062l\u0065" {
			_dfad := _dfde(_efgc)
			_gdbg = append(_gdbg, _dfad...)
		}
	}
	_fbde := _abgf(_gdbg, _ffded, _cde)
	var _gfdf []*list
	for _, _adgc := range _fbde {
		_gacef := _fcda(_adgc)
		_gfdf = append(_gfdf, _gacef...)
	}
	return _gfdf
}
func (_cfebe *textPara) taken() bool { return _cfebe == nil || _cfebe._bfaab }
func (_aggf rulingList) removeDuplicates() rulingList {
	if len(_aggf) == 0 {
		return nil
	}
	_aggf.sort()
	_gaab := rulingList{_aggf[0]}
	for _, _fcgg := range _aggf[1:] {
		if _fcgg.equals(_gaab[len(_gaab)-1]) {
			continue
		}
		_gaab = append(_gaab, _fcgg)
	}
	return _gaab
}
func _egfa(_cfeff []*textLine, _egaga string, _gafea []*list) *list {
	return &list{_eadba: _cfeff, _dbaad: _egaga, _bdfaa: _gafea}
}
func (_ccad *textPara) toTextMarks(_bgdd *int) []TextMark {
	if _ccad._cece == nil {
		return _ccad.toCellTextMarks(_bgdd)
	}
	var _cdbe []TextMark
	for _ffabc := 0; _ffabc < _ccad._cece._egbbf; _ffabc++ {
		for _febgg := 0; _febgg < _ccad._cece._ddcega; _febgg++ {
			_bfbd := _ccad._cece.get(_febgg, _ffabc)
			if _bfbd == nil {
				_cdbe = _eaagb(_cdbe, _bgdd, "\u0009")
			} else {
				_ffffg := _bfbd.toCellTextMarks(_bgdd)
				_cdbe = append(_cdbe, _ffffg...)
			}
			_cdbe = _eaagb(_cdbe, _bgdd, "\u0020")
		}
		if _ffabc < _ccad._cece._egbbf-1 {
			_cdbe = _eaagb(_cdbe, _bgdd, "\u000a")
		}
	}
	_fdce := _ccad._cece
	if _fdce.isExportable() {
		_fgeb := _fdce.toTextTable()
		_cdbe = _daac(_cdbe, &_fgeb)
	}
	return _cdbe
}
func (_ccfbe rulingList) augmentGrid() (rulingList, rulingList) {
	_cbgde, _gaacf := _ccfbe.vertsHorzs()
	if len(_cbgde) == 0 || len(_gaacf) == 0 {
		return _cbgde, _gaacf
	}
	_dacd, _abgfg := _cbgde, _gaacf
	_bafac := _cbgde.bbox()
	_cdbbf := _gaacf.bbox()
	if _gfab {
		_b.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0056\u003d\u0025\u0036\u002e\u0032\u0066", _bafac)
		_b.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0048\u003d\u0025\u0036\u002e\u0032\u0066", _cdbbf)
	}
	var _ebcbf, _ecdeg, _fggc, _fbdd *ruling
	if _cdbbf.Llx < _bafac.Llx-_beac {
		_ebcbf = &ruling{_cebe: _cccg, _gcafd: _cege, _gbaff: _cdbbf.Llx, _fabe: _bafac.Lly, _ccba: _bafac.Ury}
		_cbgde = append(rulingList{_ebcbf}, _cbgde...)
	}
	if _cdbbf.Urx > _bafac.Urx+_beac {
		_ecdeg = &ruling{_cebe: _cccg, _gcafd: _cege, _gbaff: _cdbbf.Urx, _fabe: _bafac.Lly, _ccba: _bafac.Ury}
		_cbgde = append(_cbgde, _ecdeg)
	}
	if _bafac.Lly < _cdbbf.Lly-_beac {
		_fggc = &ruling{_cebe: _cccg, _gcafd: _bfcbd, _gbaff: _bafac.Lly, _fabe: _cdbbf.Llx, _ccba: _cdbbf.Urx}
		_gaacf = append(rulingList{_fggc}, _gaacf...)
	}
	if _bafac.Ury > _cdbbf.Ury+_beac {
		_fbdd = &ruling{_cebe: _cccg, _gcafd: _bfcbd, _gbaff: _bafac.Ury, _fabe: _cdbbf.Llx, _ccba: _cdbbf.Urx}
		_gaacf = append(_gaacf, _fbdd)
	}
	if len(_cbgde)+len(_gaacf) == len(_ccfbe) {
		return _dacd, _abgfg
	}
	_dgda := append(_cbgde, _gaacf...)
	_ccfbe.log("u\u006e\u0061\u0075\u0067\u006d\u0065\u006e\u0074\u0065\u0064")
	_dgda.log("\u0061u\u0067\u006d\u0065\u006e\u0074\u0065d")
	return _cbgde, _gaacf
}
func (_bcfge rulingList) merge() *ruling {
	_cfedd := _bcfge[0]._gbaff
	_fbgb := _bcfge[0]._fabe
	_ddcde := _bcfge[0]._ccba
	for _, _ccfca := range _bcfge[1:] {
		_cfedd += _ccfca._gbaff
		if _ccfca._fabe < _fbgb {
			_fbgb = _ccfca._fabe
		}
		if _ccfca._ccba > _ddcde {
			_ddcde = _ccfca._ccba
		}
	}
	_bgdc := &ruling{_gcafd: _bcfge[0]._gcafd, _cebe: _bcfge[0]._cebe, Color: _bcfge[0].Color, _gbaff: _cfedd / float64(len(_bcfge)), _fabe: _fbgb, _ccba: _ddcde}
	if _bffed {
		_b.Log.Info("\u006de\u0072g\u0065\u003a\u0020\u0025\u0032d\u0020\u0076e\u0063\u0073\u0020\u0025\u0073", len(_bcfge), _bgdc)
		for _egbed, _cdccf := range _bcfge {
			_d.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _egbed, _cdccf)
		}
	}
	return _bgdc
}
func _defc(_beaca *paraList) map[int][]*textLine {
	_bfba := map[int][]*textLine{}
	for _, _fgge := range *_beaca {
		for _, _eefa := range _fgge._gdc {
			if !_egdbf(_eefa) {
				_b.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
				continue
			}
			_fgffg := _eefa._edee[0]._ecdf[0]._bcab
			_bfba[_fgffg] = append(_bfba[_fgffg], _eefa)
		}
		if _fgge._cece != nil {
			_ebae := _fgge._cece._bcec
			for _, _ebdb := range _ebae {
				for _, _fgga := range _ebdb._gdc {
					if !_egdbf(_fgga) {
						_b.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
						continue
					}
					_dbdc := _fgga._edee[0]._ecdf[0]._bcab
					_bfba[_dbdc] = append(_bfba[_dbdc], _fgga)
				}
			}
		}
	}
	return _bfba
}
func _feeb(_bcbfa []pathSection) rulingList {
	_gdfgg(_bcbfa)
	if _gfab {
		_b.Log.Info("\u006d\u0061k\u0065\u0053\u0074\u0072\u006f\u006b\u0065\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0064\u0020\u0073\u0074\u0072ok\u0065\u0073", len(_bcbfa))
	}
	var _ggff rulingList
	for _, _fcfgdb := range _bcbfa {
		for _, _ffcfc := range _fcfgdb._cdc {
			if len(_ffcfc._agea) < 2 {
				continue
			}
			_dffbg := _ffcfc._agea[0]
			for _, _cfdfe := range _ffcfc._agea[1:] {
				if _adfgfg, _fabb := _bfge(_dffbg, _cfdfe, _fcfgdb.Color); _fabb {
					_ggff = append(_ggff, _adfgfg)
				}
				_dffbg = _cfdfe
			}
		}
	}
	if _gfab {
		_b.Log.Info("m\u0061\u006b\u0065\u0053tr\u006fk\u0065\u0052\u0075\u006c\u0069n\u0067\u0073\u003a\u0020\u0025\u0073", _ggff)
	}
	return _ggff
}
func (_bfad rectRuling) checkWidth(_dfecf, _eddbe float64) (float64, bool) {
	_daage := _eddbe - _dfecf
	_caddc := _daage <= _egba
	return _daage, _caddc
}

type textMark struct {
	_bg.PdfRectangle
	_ecefa int
	_aaag  string
	_aca   string
	_bebg  *_bg.PdfFont
	_cfgcd float64
	_deed  float64
	_aefd  _ef.Matrix
	_dadg  _ef.Point
	_beceb _bg.PdfRectangle
	_ddbdc _ab.Color
	_dgad  _ab.Color
	_afabc _gbc.PdfObject
	_gabe  []string
	Tw     float64
	Th     float64
	_bcab  int
	_fcbe  int
}

func (_cdba *stateStack) empty() bool { return len(*_cdba) == 0 }
func _cfda(_ddad, _edbg bounded) float64 {
	_adebc := _fgc(_ddad, _edbg)
	if !_gacaa(_adebc) {
		return _adebc
	}
	return _ccfa(_ddad, _edbg)
}
func (_bagdg *textPara) fontsize() float64 { return _bagdg._gdc[0]._gdbd }

// String returns a human readable description of `vecs`.
func (_babe rulingList) String() string {
	if len(_babe) == 0 {
		return "\u007b \u0045\u004d\u0050\u0054\u0059\u0020}"
	}
	_edefe, _faad := _babe.vertsHorzs()
	_febge := len(_edefe)
	_edea := len(_faad)
	if _febge == 0 || _edea == 0 {
		return _d.Sprintf("\u007b%\u0064\u0020\u0078\u0020\u0025\u0064}", _febge, _edea)
	}
	_fdac := _bg.PdfRectangle{Llx: _edefe[0]._gbaff, Urx: _edefe[_febge-1]._gbaff, Lly: _faad[_edea-1]._gbaff, Ury: _faad[0]._gbaff}
	return _d.Sprintf("\u007b\u0025d\u0020\u0078\u0020%\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u007d", _febge, _edea, _fdac)
}
func (_gdfb paraList) readBefore(_gcfe []int, _edga, _ggef int) bool {
	_ffffe, _baaa := _gdfb[_edga], _gdfb[_ggef]
	if _begbef(_ffffe, _baaa) && _ffffe.Lly > _baaa.Lly {
		return true
	}
	if !(_ffffe._dgcfe.Urx < _baaa._dgcfe.Llx) {
		return false
	}
	_dgbba, _dggfb := _ffffe.Lly, _baaa.Lly
	if _dgbba > _dggfb {
		_dggfb, _dgbba = _dgbba, _dggfb
	}
	_cefa := _bb.Max(_ffffe._dgcfe.Llx, _baaa._dgcfe.Llx)
	_bebe := _bb.Min(_ffffe._dgcfe.Urx, _baaa._dgcfe.Urx)
	_gcadc := _gdfb.llyRange(_gcfe, _dgbba, _dggfb)
	for _, _gfede := range _gcadc {
		if _gfede == _edga || _gfede == _ggef {
			continue
		}
		_affd := _gdfb[_gfede]
		if _affd._dgcfe.Llx <= _bebe && _cefa <= _affd._dgcfe.Urx {
			return false
		}
	}
	return true
}
func _ce(_ag string, _ed bool, _cb bool) BidiText {
	_bc := "\u006c\u0074\u0072"
	if _cb {
		_bc = "\u0074\u0074\u0062"
	} else if !_ed {
		_bc = "\u0072\u0074\u006c"
	}
	return BidiText{_feb: _ag, _cae: _bc}
}
