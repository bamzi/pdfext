package document

import (
	_b "encoding/binary"
	_bf "fmt"
	_bg "io"
	_be "math"
	_d "runtime/debug"

	_g "github.com/bamzi/pdfext/common"
	_f "github.com/bamzi/pdfext/internal/bitwise"
	_e "github.com/bamzi/pdfext/internal/jbig2/basic"
	_gc "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_gd "github.com/bamzi/pdfext/internal/jbig2/document/segments"
	_de "github.com/bamzi/pdfext/internal/jbig2/encoder/classer"
	_eb "github.com/bamzi/pdfext/internal/jbig2/errors"
)

func (_fa *Document) AddClassifiedPage(bm *_gc.Bitmap, method _de.Method) (_deg error) {
	const _caf = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0041\u0064d\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064P\u0061\u0067\u0065"
	if !_fa.FullHeaders && _fa.NumberOfPages != 0 {
		return _eb.Error(_caf, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	if _fa.Classer == nil {
		if _fa.Classer, _deg = _de.Init(_de.DefaultSettings()); _deg != nil {
			return _eb.Wrap(_deg, _caf, "")
		}
	}
	_cf := int(_fa.nextPageNumber())
	_cfe := &Page{Segments: []*_gd.Header{}, Bitmap: bm, Document: _fa, FinalHeight: bm.Height, FinalWidth: bm.Width, PageNumber: _cf}
	_fa.Pages[_cf] = _cfe
	switch method {
	case _de.RankHaus:
		_cfe.EncodingMethod = RankHausEM
	case _de.Correlation:
		_cfe.EncodingMethod = CorrelationEM
	}
	_cfe.AddPageInformationSegment()
	if _deg = _fa.Classer.AddPage(bm, _cf, method); _deg != nil {
		return _eb.Wrap(_deg, _caf, "")
	}
	if _fa.FullHeaders {
		_cfe.AddEndOfPageSegment()
	}
	return nil
}
func (_agf *Document) GetPage(pageNumber int) (_gd.Pager, error) {
	const _bgae = "\u0044\u006fc\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074\u0050\u0061\u0067\u0065"
	if pageNumber < 0 {
		_g.Log.Debug("\u004a\u0042\u0049\u00472\u0020\u0050\u0061\u0067\u0065\u0020\u002d\u0020\u0047e\u0074\u0050\u0061\u0067\u0065\u003a\u0020\u0025\u0064\u002e\u0020\u0050\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u006c\u006f\u0077\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0020\u0025\u0073", pageNumber, _d.Stack())
		return nil, _eb.Errorf(_bgae, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u002d\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064 \u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064", pageNumber)
	}
	if pageNumber > len(_agf.Pages) {
		_g.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _d.Stack())
		return nil, _eb.Error(_bgae, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_ff, _ffg := _agf.Pages[pageNumber]
	if !_ffg {
		_g.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _d.Stack())
		return nil, _eb.Errorf(_bgae, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _ff, nil
}
func (_age *Document) encodeSegment(_af *_gd.Header, _fag *int) error {
	const _cca = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	_af.SegmentNumber = _age.nextSegmentNumber()
	_bgd, _baac := _af.Encode(_age._ee)
	if _baac != nil {
		return _eb.Wrapf(_baac, _cca, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _af.SegmentNumber)
	}
	*_fag += _bgd
	return nil
}
func (_cec *Document) produceClassifiedPage(_egb *Page, _ged *_gd.Header) (_gb error) {
	const _bc = "p\u0072\u006f\u0064\u0075ce\u0043l\u0061\u0073\u0073\u0069\u0066i\u0065\u0064\u0050\u0061\u0067\u0065"
	var _cdf map[int]int
	_dga := _cec._ec
	_ac := []*_gd.Header{_ged}
	if len(_cec._ca[_egb.PageNumber]) > 0 {
		_cdf = map[int]int{}
		_gfa, _cae := _cec.addSymbolDictionary(_egb.PageNumber, _cec.Classer.UndilatedTemplates, _cec._ca[_egb.PageNumber], _cdf, false)
		if _cae != nil {
			return _eb.Wrap(_cae, _bc, "")
		}
		_ac = append(_ac, _gfa)
		_dga += len(_cec._ca[_egb.PageNumber])
	}
	_da := _cec._cd[_egb.PageNumber]
	_g.Log.Debug("P\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020c\u006f\u006d\u0070\u0073: \u0025\u0076", _egb.PageNumber, _da)
	_egb.addTextRegionSegment(_ac, _cec._df, _cdf, _cec._cd[_egb.PageNumber], _cec.Classer.PtaLL, _cec.Classer.UndilatedTemplates, _cec.Classer.ClassIDs, nil, _bad(_dga), len(_cec._cd[_egb.PageNumber]))
	return nil
}
func (_fffe *Page) getCombinationOperator(_debg *_gd.PageInformationSegment, _gcfc _gc.CombinationOperator) _gc.CombinationOperator {
	if _debg.CombinationOperatorOverrideAllowed() {
		return _gcfc
	}
	return _debg.CombinationOperator()
}
func _fde(_dge *Document, _fce int) *Page {
	return &Page{Document: _dge, PageNumber: _fce, Segments: []*_gd.Header{}}
}

type EncodingMethod int

func (_aab *Page) composePageBitmap() error {
	const _cge = "\u0063\u006f\u006d\u0070\u006f\u0073\u0065\u0050\u0061\u0067\u0065\u0042i\u0074\u006d\u0061\u0070"
	if _aab.PageNumber == 0 {
		return nil
	}
	_bgdd := _aab.getPageInformationSegment()
	if _bgdd == nil {
		return _eb.Error(_cge, "\u0070\u0061\u0067e \u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069o\u006e \u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_eggf, _bdd := _bgdd.GetSegmentData()
	if _bdd != nil {
		return _bdd
	}
	_dda, _bdde := _eggf.(*_gd.PageInformationSegment)
	if !_bdde {
		return _eb.Error(_cge, "\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006da\u0074\u0069\u006f\u006e \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
	}
	if _bdd = _aab.createPage(_dda); _bdd != nil {
		return _eb.Wrap(_bdd, _cge, "")
	}
	_aab.clearSegmentData()
	return nil
}
func (_daee *Page) getResolutionY() (int, error) {
	const _dcab = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0059"
	if _daee.ResolutionY != 0 {
		return _daee.ResolutionY, nil
	}
	_aae := _daee.getPageInformationSegment()
	if _aae == nil {
		return 0, _eb.Error(_dcab, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_ccca, _cdb := _aae.GetSegmentData()
	if _cdb != nil {
		return 0, _eb.Wrap(_cdb, _dcab, "")
	}
	_dea, _cfa := _ccca.(*_gd.PageInformationSegment)
	if !_cfa {
		return 0, _eb.Errorf(_dcab, "\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0027%\u0054\u0027", _ccca)
	}
	_daee.ResolutionY = _dea.ResolutionY
	return _daee.ResolutionY, nil
}

var _fc = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

func (_gfb *Document) nextPageNumber() uint32 { _gfb.NumberOfPages++; return _gfb.NumberOfPages }
func (_gde *Page) String() string {
	return _bf.Sprintf("\u0050\u0061\u0067\u0065\u0020\u0023\u0025\u0064", _gde.PageNumber)
}
func (_fcg *Page) collectPageStripes() (_bac []_gd.Segmenter, _feg error) {
	const _fca = "\u0063o\u006cl\u0065\u0063\u0074\u0050\u0061g\u0065\u0053t\u0072\u0069\u0070\u0065\u0073"
	var _dcb _gd.Segmenter
	for _, _fdec := range _fcg.Segments {
		switch _fdec.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_dcb, _feg = _fdec.GetSegmentData()
			if _feg != nil {
				return nil, _eb.Wrap(_feg, _fca, "")
			}
			_bac = append(_bac, _dcb)
		case 50:
			_dcb, _feg = _fdec.GetSegmentData()
			if _feg != nil {
				return nil, _feg
			}
			_fbag, _fccg := _dcb.(*_gd.EndOfStripe)
			if !_fccg {
				return nil, _eb.Errorf(_fca, "\u0045\u006e\u0064\u004f\u0066\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006f\u0066\u0020\u0076\u0061l\u0069\u0064\u0020\u0074\u0079p\u0065\u003a \u0027\u0025\u0054\u0027", _dcb)
			}
			_bac = append(_bac, _fbag)
			_fcg.FinalHeight = _fbag.LineNumber()
		}
	}
	return _bac, nil
}
func (_ggd *Page) AddGenericRegion(bm *_gc.Bitmap, xloc, yloc, template int, tp _gd.Type, duplicateLineRemoval bool) error {
	const _fbe = "P\u0061\u0067\u0065\u002eAd\u0064G\u0065\u006e\u0065\u0072\u0069c\u0052\u0065\u0067\u0069\u006f\u006e"
	_ffde := &_gd.GenericRegion{}
	if _cfde := _ffde.InitEncode(bm, xloc, yloc, template, duplicateLineRemoval); _cfde != nil {
		return _eb.Wrap(_cfde, _fbe, "")
	}
	_cbd := &_gd.Header{Type: _gd.TImmediateGenericRegion, PageAssociation: _ggd.PageNumber, SegmentData: _ffde}
	_ggd.Segments = append(_ggd.Segments, _cbd)
	return nil
}
func (_cb *Document) completeClassifiedPages() (_bgg error) {
	const _eg = "\u0063\u006f\u006dpl\u0065\u0074\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067\u0065\u0073"
	if _cb.Classer == nil {
		return nil
	}
	_cb._cc = make([]int, _cb.Classer.UndilatedTemplates.Size())
	for _ab := 0; _ab < _cb.Classer.ClassIDs.Size(); _ab++ {
		_dfg, _ge := _cb.Classer.ClassIDs.Get(_ab)
		if _ge != nil {
			return _eb.Wrapf(_ge, _eg, "\u0063\u006c\u0061\u0073s \u0077\u0069\u0074\u0068\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027", _ab)
		}
		_cb._cc[_dfg]++
	}
	var _ag []int
	for _ga := 0; _ga < _cb.Classer.UndilatedTemplates.Size(); _ga++ {
		if _cb.NumberOfPages == 1 || _cb._cc[_ga] > 1 {
			_ag = append(_ag, _ga)
		}
	}
	var (
		_cg  *Page
		_bee bool
	)
	for _ce, _bfd := range *_cb.Classer.ComponentPageNumbers {
		if _cg, _bee = _cb.Pages[_bfd]; !_bee {
			return _eb.Errorf(_eg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _ce)
		}
		if _cg.EncodingMethod == GenericEM {
			_g.Log.Error("\u0047\u0065\u006e\u0065\u0072\u0069c\u0020\u0070\u0061g\u0065\u0020\u0077i\u0074\u0068\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020ma\u0070\u0070\u0065\u0064\u0020\u0061\u0073\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0020\u0070\u0061\u0067\u0065", _ce)
			continue
		}
		_cb._cd[_bfd] = append(_cb._cd[_bfd], _ce)
		_ccd, _bb := _cb.Classer.ClassIDs.Get(_ce)
		if _bb != nil {
			return _eb.Wrapf(_bb, _eg, "\u006e\u006f\u0020\u0073uc\u0068\u0020\u0063\u006c\u0061\u0073\u0073\u0049\u0044\u003a\u0020\u0025\u0064", _ce)
		}
		if _cb._cc[_ccd] == 1 && _cb.NumberOfPages != 1 {
			_bd := append(_cb._ca[_bfd], _ccd)
			_cb._ca[_bfd] = _bd
		}
	}
	if _bgg = _cb.Classer.ComputeLLCorners(); _bgg != nil {
		return _eb.Wrap(_bgg, _eg, "")
	}
	if _, _bgg = _cb.addSymbolDictionary(0, _cb.Classer.UndilatedTemplates, _ag, _cb._df, false); _bgg != nil {
		return _eb.Wrap(_bgg, _eg, "")
	}
	return nil
}
func (_gdc *Document) encodeEOFHeader(_aed _f.BinaryWriter) (_fdcc int, _dgc error) {
	_gbb := &_gd.Header{SegmentNumber: _gdc.nextSegmentNumber(), Type: _gd.TEndOfFile}
	if _fdcc, _dgc = _gbb.Encode(_aed); _dgc != nil {
		return 0, _eb.Wrap(_dgc, "\u0065n\u0063o\u0064\u0065\u0045\u004f\u0046\u0048\u0065\u0061\u0064\u0065\u0072", "")
	}
	return _fdcc, nil
}
func (_dbb *Document) completeSymbols() (_cecd error) {
	const _aa = "\u0063o\u006dp\u006c\u0065\u0074\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	if _dbb.Classer == nil {
		return nil
	}
	if _dbb.Classer.UndilatedTemplates == nil {
		return _eb.Error(_aa, "\u006e\u006f t\u0065\u006d\u0070l\u0061\u0074\u0065\u0073 de\u0066in\u0065\u0064\u0020\u0066\u006f\u0072\u0020th\u0065\u0020\u0063\u006c\u0061\u0073\u0073e\u0072")
	}
	_efg := len(_dbb.Pages) == 1
	_gcf := make([]int, _dbb.Classer.UndilatedTemplates.Size())
	var _gbf int
	for _bag := 0; _bag < _dbb.Classer.ClassIDs.Size(); _bag++ {
		_gbf, _cecd = _dbb.Classer.ClassIDs.Get(_bag)
		if _cecd != nil {
			return _eb.Wrap(_cecd, _aa, "\u0063\u006c\u0061\u0073\u0073\u0020\u0049\u0044\u0027\u0073")
		}
		_gcf[_gbf]++
	}
	var _gfcg []int
	for _ebf := 0; _ebf < _dbb.Classer.UndilatedTemplates.Size(); _ebf++ {
		if _gcf[_ebf] == 0 {
			return _eb.Error(_aa, "\u006eo\u0020\u0073y\u006d\u0062\u006f\u006cs\u0020\u0069\u006es\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006fun\u0064\u0020\u0066o\u0072\u0020g\u0069\u0076\u0065\u006e\u0020\u0063l\u0061\u0073s\u003f\u0020")
		}
		if _gcf[_ebf] > 1 || _efg {
			_gfcg = append(_gfcg, _ebf)
		}
	}
	_dbb._ec = len(_gfcg)
	var _gad, _bgb int
	for _dc := 0; _dc < _dbb.Classer.ComponentPageNumbers.Size(); _dc++ {
		_gad, _cecd = _dbb.Classer.ComponentPageNumbers.Get(_dc)
		if _cecd != nil {
			return _eb.Wrapf(_cecd, _aa, "p\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027 \u006e\u006f\u0074\u0020\u0066\u006f\u0075nd\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0063\u006c\u0061\u0073\u0073\u0065r \u0070\u0061g\u0065\u006e\u0075\u006d\u0062\u0065\u0072\u0073", _dc)
		}
		_bgb, _cecd = _dbb.Classer.ClassIDs.Get(_dc)
		if _cecd != nil {
			return _eb.Wrapf(_cecd, _aa, "\u0063\u0061\u006e\u0027\u0074\u0020\u0067e\u0074\u0020\u0073y\u006d\u0062\u006f\u006c \u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072", _gad)
		}
		if _gcf[_bgb] == 1 && !_efg {
			_dbb._ca[_gad] = append(_dbb._ca[_gad], _bgb)
		}
	}
	if _cecd = _dbb.Classer.ComputeLLCorners(); _cecd != nil {
		return _eb.Wrap(_cecd, _aa, "")
	}
	return nil
}
func (_cfg *Globals) GetSegment(segmentNumber int) (*_gd.Header, error) {
	const _bbc = "\u0047l\u006fb\u0061\u006c\u0073\u002e\u0047e\u0074\u0053e\u0067\u006d\u0065\u006e\u0074"
	if _cfg == nil {
		return nil, _eb.Error(_bbc, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_cfg.Segments) == 0 {
		return nil, _eb.Error(_bbc, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	var _gff *_gd.Header
	for _, _gff = range _cfg.Segments {
		if _gff.SegmentNumber == uint32(segmentNumber) {
			break
		}
	}
	if _gff == nil {
		return nil, _eb.Error(_bbc, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _gff, nil
}
func (_bea *Document) isFileHeaderPresent() (bool, error) {
	_bea.InputStream.Mark()
	for _, _eee := range _fc {
		_bgbb, _fg := _bea.InputStream.ReadByte()
		if _fg != nil {
			return false, _fg
		}
		if _eee != _bgbb {
			_bea.InputStream.Reset()
			return false, nil
		}
	}
	_bea.InputStream.Reset()
	return true, nil
}
func (_bgbbg *Page) createNormalPage(_egf *_gd.PageInformationSegment) error {
	const _bcae = "\u0063\u0072e\u0061\u0074\u0065N\u006f\u0072\u006d\u0061\u006c\u0050\u0061\u0067\u0065"
	_bgbbg.Bitmap = _gc.New(_egf.PageBMWidth, _egf.PageBMHeight)
	if _egf.DefaultPixelValue != 0 {
		_bgbbg.Bitmap.SetDefaultPixel()
	}
	for _, _fcef := range _bgbbg.Segments {
		switch _fcef.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_g.Log.Trace("\u0047\u0065\u0074\u0074in\u0067\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0064", _fcef.SegmentNumber)
			_gea, _cffb := _fcef.GetSegmentData()
			if _cffb != nil {
				return _cffb
			}
			_fded, _gae := _gea.(_gd.Regioner)
			if !_gae {
				_g.Log.Debug("\u0053\u0065g\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069on\u0065\u0072", _gea)
				return _eb.Errorf(_bcae, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062i\u0067\u0032\u0020\u0073\u0065\u0067\u006den\u0074\u0020\u0074\u0079p\u0065\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061 R\u0065\u0067i\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _gea)
			}
			_fdac, _cffb := _fded.GetRegionBitmap()
			if _cffb != nil {
				return _eb.Wrap(_cffb, _bcae, "")
			}
			if _bgbbg.fitsPage(_egf, _fdac) {
				_bgbbg.Bitmap = _fdac
			} else {
				_bab := _fded.GetRegionInfo()
				_gdcb := _bgbbg.getCombinationOperator(_egf, _bab.CombinaionOperator)
				_cffb = _gc.Blit(_fdac, _bgbbg.Bitmap, int(_bab.XLocation), int(_bab.YLocation), _gdcb)
				if _cffb != nil {
					return _eb.Wrap(_cffb, _bcae, "")
				}
			}
		}
	}
	return nil
}

const (
	GenericEM EncodingMethod = iota
	CorrelationEM
	RankHausEM
)

func InitEncodeDocument(fullHeaders bool) *Document {
	return &Document{FullHeaders: fullHeaders, _ee: _f.BufferedMSB(), Pages: map[int]*Page{}, _ca: map[int][]int{}, _df: map[int]int{}, _cd: map[int][]int{}}
}

type Page struct {
	Segments           []*_gd.Header
	PageNumber         int
	Bitmap             *_gc.Bitmap
	FinalHeight        int
	FinalWidth         int
	ResolutionX        int
	ResolutionY        int
	IsLossless         bool
	Document           *Document
	FirstSegmentNumber int
	EncodingMethod     EncodingMethod
	BlackIsOne         bool
}

func (_acef *Page) createStripedPage(_agd *_gd.PageInformationSegment) error {
	const _cfdg = "\u0063\u0072\u0065\u0061\u0074\u0065\u0053\u0074\u0072\u0069\u0070\u0065d\u0050\u0061\u0067\u0065"
	_cfda, _agc := _acef.collectPageStripes()
	if _agc != nil {
		return _eb.Wrap(_agc, _cfdg, "")
	}
	var _eag int
	for _, _fad := range _cfda {
		if _bfdc, _fef := _fad.(*_gd.EndOfStripe); _fef {
			_eag = _bfdc.LineNumber() + 1
		} else {
			_acdc := _fad.(_gd.Regioner)
			_eefe := _acdc.GetRegionInfo()
			_gbe := _acef.getCombinationOperator(_agd, _eefe.CombinaionOperator)
			_gce, _dde := _acdc.GetRegionBitmap()
			if _dde != nil {
				return _eb.Wrap(_dde, _cfdg, "")
			}
			_dde = _gc.Blit(_gce, _acef.Bitmap, int(_eefe.XLocation), _eag, _gbe)
			if _dde != nil {
				return _eb.Wrap(_dde, _cfdg, "")
			}
		}
	}
	return nil
}
func (_gee *Page) getResolutionX() (int, error) {
	const _ecg = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0058"
	if _gee.ResolutionX != 0 {
		return _gee.ResolutionX, nil
	}
	_egbd := _gee.getPageInformationSegment()
	if _egbd == nil {
		return 0, _eb.Error(_ecg, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_bage, _aee := _egbd.GetSegmentData()
	if _aee != nil {
		return 0, _eb.Wrap(_aee, _ecg, "")
	}
	_egfg, _cfdf := _bage.(*_gd.PageInformationSegment)
	if !_cfdf {
		return 0, _eb.Errorf(_ecg, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _bage)
	}
	_gee.ResolutionX = _egfg.ResolutionX
	return _gee.ResolutionX, nil
}
func (_aff *Page) GetBitmap() (_deb *_gc.Bitmap, _dbe error) {
	_g.Log.Trace(_bf.Sprintf("\u005b\u0050\u0041G\u0045\u005d\u005b\u0023%\u0064\u005d\u0020\u0047\u0065\u0074\u0042i\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e", _aff.PageNumber))
	defer func() {
		if _dbe != nil {
			_g.Log.Trace(_bf.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074B\u0069t\u006d\u0061\u0070\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _aff.PageNumber, _dbe))
		} else {
			_g.Log.Trace(_bf.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064]\u0020\u0047\u0065\u0074\u0042\u0069\u0074m\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064", _aff.PageNumber))
		}
	}()
	if _aff.Bitmap != nil {
		return _aff.Bitmap, nil
	}
	_dbe = _aff.composePageBitmap()
	if _dbe != nil {
		return nil, _dbe
	}
	return _aff.Bitmap, nil
}
func (_fafd *Page) GetResolutionY() (int, error) { return _fafd.getResolutionY() }
func (_aba *Page) GetSegment(number int) (*_gd.Header, error) {
	const _fddd = "\u0050a\u0067e\u002e\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	for _, _ccb := range _aba.Segments {
		if _ccb.SegmentNumber == uint32(number) {
			return _ccb, nil
		}
	}
	_ccg := make([]uint32, len(_aba.Segments))
	for _fe, _ade := range _aba.Segments {
		_ccg[_fe] = _ade.SegmentNumber
	}
	return nil, _eb.Errorf(_fddd, "\u0073e\u0067\u006d\u0065n\u0074\u0020\u0077i\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u003a\u0020'%\u0064'\u002e\u0020\u004b\u006e\u006f\u0077n\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u003a \u0025\u0076", number, _aba.PageNumber, _ccg)
}
func (_dgcg *Globals) GetSegmentByIndex(index int) (*_gd.Header, error) {
	const _ccc = "\u0047l\u006f\u0062\u0061\u006cs\u002e\u0047\u0065\u0074\u0053e\u0067m\u0065n\u0074\u0042\u0079\u0049\u006e\u0064\u0065x"
	if _dgcg == nil {
		return nil, _eb.Error(_ccc, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_dgcg.Segments) == 0 {
		return nil, _eb.Error(_ccc, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	if index > len(_dgcg.Segments)-1 {
		return nil, _eb.Error(_ccc, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	return _dgcg.Segments[index], nil
}
func (_cdg *Page) AddPageInformationSegment() {
	_ebba := &_gd.PageInformationSegment{PageBMWidth: _cdg.FinalWidth, PageBMHeight: _cdg.FinalHeight, ResolutionX: _cdg.ResolutionX, ResolutionY: _cdg.ResolutionY, IsLossless: _cdg.IsLossless}
	if _cdg.BlackIsOne {
		_ebba.DefaultPixelValue = uint8(0x1)
	}
	_bed := &_gd.Header{PageAssociation: _cdg.PageNumber, SegmentDataLength: uint64(_ebba.Size()), SegmentData: _ebba, Type: _gd.TPageInformation}
	_cdg.Segments = append(_cdg.Segments, _bed)
}
func (_ffa *Page) countRegions() int {
	var _gbeb int
	for _, _fagg := range _ffa.Segments {
		switch _fagg.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_gbeb++
		}
	}
	return _gbeb
}
func (_fcc *Document) reachedEOF(_baf int64) (bool, error) {
	const _dfab = "\u0072\u0065\u0061\u0063\u0068\u0065\u0064\u0045\u004f\u0046"
	_, _fcf := _fcc.InputStream.Seek(_baf, _bg.SeekStart)
	if _fcf != nil {
		_g.Log.Debug("\u0072\u0065\u0061c\u0068\u0065\u0064\u0045\u004f\u0046\u0020\u002d\u0020\u0064\u002e\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u002e\u0053\u0065\u0065\u006b\u0020\u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fcf)
		return false, _eb.Wrap(_fcf, _dfab, "\u0069n\u0070\u0075\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020s\u0065\u0065\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_, _fcf = _fcc.InputStream.ReadBits(32)
	if _fcf == _bg.EOF {
		return true, nil
	} else if _fcf != nil {
		return false, _eb.Wrap(_fcf, _dfab, "")
	}
	return false, nil
}
func (_fdf *Document) Encode() (_bcg []byte, _efa error) {
	const _abb = "\u0044o\u0063u\u006d\u0065\u006e\u0074\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _fdd, _dbd int
	if _fdf.FullHeaders {
		if _fdd, _efa = _fdf.encodeFileHeader(_fdf._ee); _efa != nil {
			return nil, _eb.Wrap(_efa, _abb, "")
		}
	}
	var (
		_gcb bool
		_ebc *_gd.Header
		_gdd *Page
	)
	if _efa = _fdf.completeClassifiedPages(); _efa != nil {
		return nil, _eb.Wrap(_efa, _abb, "")
	}
	if _efa = _fdf.produceClassifiedPages(); _efa != nil {
		return nil, _eb.Wrap(_efa, _abb, "")
	}
	if _fdf.GlobalSegments != nil {
		for _, _ebc = range _fdf.GlobalSegments.Segments {
			if _efa = _fdf.encodeSegment(_ebc, &_fdd); _efa != nil {
				return nil, _eb.Wrap(_efa, _abb, "")
			}
		}
	}
	for _bcd := 1; _bcd <= int(_fdf.NumberOfPages); _bcd++ {
		if _gdd, _gcb = _fdf.Pages[_bcd]; !_gcb {
			return nil, _eb.Errorf(_abb, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _bcd)
		}
		for _, _ebc = range _gdd.Segments {
			if _efa = _fdf.encodeSegment(_ebc, &_fdd); _efa != nil {
				return nil, _eb.Wrap(_efa, _abb, "")
			}
		}
	}
	if _fdf.FullHeaders {
		if _dbd, _efa = _fdf.encodeEOFHeader(_fdf._ee); _efa != nil {
			return nil, _eb.Wrap(_efa, _abb, "")
		}
		_fdd += _dbd
	}
	_bcg = _fdf._ee.Data()
	if len(_bcg) != _fdd {
		_g.Log.Debug("\u0042\u0079\u0074\u0065\u0073 \u0077\u0072\u0069\u0074\u0074\u0065\u006e \u0028\u006e\u0029\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0065\u006e\u0063\u006fd\u0065\u0064\u003a\u0020\u0027\u0025d\u0027", _fdd, len(_bcg))
	}
	return _bcg, nil
}
func (_ege *Page) fitsPage(_gbd *_gd.PageInformationSegment, _aaae *_gc.Bitmap) bool {
	return _ege.countRegions() == 1 && _gbd.DefaultPixelValue == 0 && _gbd.PageBMWidth == _aaae.Width && _gbd.PageBMHeight == _aaae.Height
}
func (_egg *Document) parseFileHeader() error {
	const _dff = "\u0070a\u0072s\u0065\u0046\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_, _ebe := _egg.InputStream.Seek(8, _bg.SeekStart)
	if _ebe != nil {
		return _eb.Wrap(_ebe, _dff, "\u0069\u0064")
	}
	_, _ebe = _egg.InputStream.ReadBits(5)
	if _ebe != nil {
		return _eb.Wrap(_ebe, _dff, "\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074\u0073")
	}
	_cfd, _ebe := _egg.InputStream.ReadBit()
	if _ebe != nil {
		return _eb.Wrap(_ebe, _dff, "\u0065x\u0074e\u006e\u0064\u0065\u0064\u0020t\u0065\u006dp\u006c\u0061\u0074\u0065\u0073")
	}
	if _cfd == 1 {
		_egg.GBUseExtTemplate = true
	}
	_cfd, _ebe = _egg.InputStream.ReadBit()
	if _ebe != nil {
		return _eb.Wrap(_ebe, _dff, "\u0075\u006e\u006b\u006eow\u006e\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _cfd != 1 {
		_egg.NumberOfPagesUnknown = false
	}
	_cfd, _ebe = _egg.InputStream.ReadBit()
	if _ebe != nil {
		return _eb.Wrap(_ebe, _dff, "\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074\u0069\u006f\u006e \u0074\u0079\u0070\u0065")
	}
	_egg.OrganizationType = _gd.OrganizationType(_cfd)
	if !_egg.NumberOfPagesUnknown {
		_egg.NumberOfPages, _ebe = _egg.InputStream.ReadUint32()
		if _ebe != nil {
			return _eb.Wrap(_ebe, _dff, "\u006eu\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0067\u0065\u0073")
		}
		_egg._a = 13
	}
	return nil
}
func _bad(_ef int) int {
	_dgg := 0
	_ad := (_ef & (_ef - 1)) == 0
	_ef >>= 1
	for ; _ef != 0; _ef >>= 1 {
		_dgg++
	}
	if _ad {
		return _dgg
	}
	return _dgg + 1
}
func (_bgf *Globals) GetSymbolDictionary() (*_gd.Header, error) {
	const _eca = "G\u006c\u006f\u0062\u0061\u006c\u0073.\u0047\u0065\u0074\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079"
	if _bgf == nil {
		return nil, _eb.Error(_eca, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_bgf.Segments) == 0 {
		return nil, _eb.Error(_eca, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	for _, _acg := range _bgf.Segments {
		if _acg.Type == _gd.TSymbolDictionary {
			return _acg, nil
		}
	}
	return nil, _eb.Error(_eca, "\u0067\u006c\u006fba\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020d\u0069c\u0074i\u006fn\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func (_abed *Page) AddEndOfPageSegment() {
	_fdg := &_gd.Header{Type: _gd.TEndOfPage, PageAssociation: _abed.PageNumber}
	_abed.Segments = append(_abed.Segments, _fdg)
}
func (_gffc *Page) getWidth() (int, error) {
	const _fbagd = "\u0067\u0065\u0074\u0057\u0069\u0064\u0074\u0068"
	if _gffc.FinalWidth != 0 {
		return _gffc.FinalWidth, nil
	}
	_gbebe := _gffc.getPageInformationSegment()
	if _gbebe == nil {
		return 0, _eb.Error(_fbagd, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_eege, _dbddb := _gbebe.GetSegmentData()
	if _dbddb != nil {
		return 0, _eb.Wrap(_dbddb, _fbagd, "")
	}
	_ecae, _aeg := _eege.(*_gd.PageInformationSegment)
	if !_aeg {
		return 0, _eb.Errorf(_fbagd, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _eege)
	}
	_gffc.FinalWidth = _ecae.PageBMWidth
	return _gffc.FinalWidth, nil
}
func (_dfd *Document) AddGenericPage(bm *_gc.Bitmap, duplicateLineRemoval bool) (_ecd error) {
	const _ae = "\u0044\u006f\u0063um\u0065\u006e\u0074\u002e\u0041\u0064\u0064\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0050\u0061\u0067\u0065"
	if !_dfd.FullHeaders && _dfd.NumberOfPages != 0 {
		return _eb.Error(_ae, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	_eef := &Page{Segments: []*_gd.Header{}, Bitmap: bm, Document: _dfd, FinalHeight: bm.Height, FinalWidth: bm.Width, IsLossless: true, BlackIsOne: bm.Color == _gc.Chocolate}
	_eef.PageNumber = int(_dfd.nextPageNumber())
	_dfd.Pages[_eef.PageNumber] = _eef
	bm.InverseData()
	_eef.AddPageInformationSegment()
	if _ecd = _eef.AddGenericRegion(bm, 0, 0, 0, _gd.TImmediateGenericRegion, duplicateLineRemoval); _ecd != nil {
		return _eb.Wrap(_ecd, _ae, "")
	}
	if _dfd.FullHeaders {
		_eef.AddEndOfPageSegment()
	}
	return nil
}
func (_badd *Document) GetNumberOfPages() (uint32, error) {
	if _badd.NumberOfPagesUnknown || _badd.NumberOfPages == 0 {
		if len(_badd.Pages) == 0 {
			if _abe := _badd.mapData(); _abe != nil {
				return 0, _eb.Wrap(_abe, "\u0044o\u0063\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074N\u0075m\u0062e\u0072\u004f\u0066\u0050\u0061\u0067\u0065s", "")
			}
		}
		return uint32(len(_badd.Pages)), nil
	}
	return _badd.NumberOfPages, nil
}
func (_eaf *Document) encodeFileHeader(_fda _f.BinaryWriter) (_ace int, _bdg error) {
	const _afa = "\u0065\u006ec\u006f\u0064\u0065F\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_ace, _bdg = _fda.Write(_fc)
	if _bdg != nil {
		return _ace, _eb.Wrap(_bdg, _afa, "\u0069\u0064")
	}
	if _bdg = _fda.WriteByte(0x01); _bdg != nil {
		return _ace, _eb.Wrap(_bdg, _afa, "\u0066\u006c\u0061g\u0073")
	}
	_ace++
	_eeb := make([]byte, 4)
	_b.BigEndian.PutUint32(_eeb, _eaf.NumberOfPages)
	_eff, _bdg := _fda.Write(_eeb)
	if _bdg != nil {
		return _eff, _eb.Wrap(_bdg, _afa, "p\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ace += _eff
	return _ace, nil
}

type Document struct {
	Pages                map[int]*Page
	NumberOfPagesUnknown bool
	NumberOfPages        uint32
	GBUseExtTemplate     bool
	InputStream          *_f.Reader
	GlobalSegments       *Globals
	OrganizationType     _gd.OrganizationType
	Classer              *_de.Classer
	XRes, YRes           int
	FullHeaders          bool
	CurrentSegmentNumber uint32
	AverageTemplates     *_gc.Bitmaps
	BaseIndexes          []int
	Refinement           bool
	RefineLevel          int
	_a                   uint8
	_ee                  *_f.BufferedWriter
	EncodeGlobals        bool
	_ec                  int
	_ca                  map[int][]int
	_cd                  map[int][]int
	_cc                  []int
	_df                  map[int]int
}

func (_fbg *Page) clearSegmentData() {
	for _gdb := range _fbg.Segments {
		_fbg.Segments[_gdb].CleanSegmentData()
	}
}
func (_efe *Document) determineRandomDataOffsets(_gec []*_gd.Header, _caa uint64) {
	if _efe.OrganizationType != _gd.ORandom {
		return
	}
	for _, _cga := range _gec {
		_cga.SegmentDataStartOffset = _caa
		_caa += _cga.SegmentDataLength
	}
}
func (_bca *Page) GetHeight() (int, error) { return _bca.getHeight() }
func (_ed *Document) addSymbolDictionary(_ced int, _dd *_gc.Bitmaps, _ega []int, _ea map[int]int, _baa bool) (*_gd.Header, error) {
	const _bbe = "\u0061\u0064\u0064\u0053ym\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079"
	_fdc := &_gd.SymbolDictionary{}
	if _gac := _fdc.InitEncode(_dd, _ega, _ea, _baa); _gac != nil {
		return nil, _gac
	}
	_dae := &_gd.Header{Type: _gd.TSymbolDictionary, PageAssociation: _ced, SegmentData: _fdc}
	if _ced == 0 {
		if _ed.GlobalSegments == nil {
			_ed.GlobalSegments = &Globals{}
		}
		_ed.GlobalSegments.AddSegment(_dae)
		return _dae, nil
	}
	_dgd, _adf := _ed.Pages[_ced]
	if !_adf {
		return nil, _eb.Errorf(_bbe, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _ced)
	}
	var (
		_eeg int
		_edb *_gd.Header
	)
	for _eeg, _edb = range _dgd.Segments {
		if _edb.Type == _gd.TPageInformation {
			break
		}
	}
	_eeg++
	_dgd.Segments = append(_dgd.Segments, nil)
	copy(_dgd.Segments[_eeg+1:], _dgd.Segments[_eeg:])
	_dgd.Segments[_eeg] = _dae
	return _dae, nil
}
func (_ba *Document) produceClassifiedPages() (_fd error) {
	const _gf = "\u0070\u0072\u006f\u0064uc\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067e\u0073"
	if _ba.Classer == nil {
		return nil
	}
	var (
		_abd *Page
		_dg  bool
		_gfc *_gd.Header
	)
	for _db := 1; _db <= int(_ba.NumberOfPages); _db++ {
		if _abd, _dg = _ba.Pages[_db]; !_dg {
			return _eb.Errorf(_gf, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _db)
		}
		if _abd.EncodingMethod == GenericEM {
			continue
		}
		if _gfc == nil {
			if _gfc, _fd = _ba.GlobalSegments.GetSymbolDictionary(); _fd != nil {
				return _eb.Wrap(_fd, _gf, "")
			}
		}
		if _fd = _ba.produceClassifiedPage(_abd, _gfc); _fd != nil {
			return _eb.Wrapf(_fd, _gf, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _db)
		}
	}
	return nil
}
func DecodeDocument(input *_f.Reader, globals *Globals) (*Document, error) {
	return _edg(input, globals)
}
func (_cfed *Document) mapData() error {
	const _aaa = "\u006da\u0070\u0044\u0061\u0074\u0061"
	var (
		_acd []*_gd.Header
		_ebb int64
		_egc _gd.Type
	)
	_ede, _gab := _cfed.isFileHeaderPresent()
	if _gab != nil {
		return _eb.Wrap(_gab, _aaa, "")
	}
	if _ede {
		if _gab = _cfed.parseFileHeader(); _gab != nil {
			return _eb.Wrap(_gab, _aaa, "")
		}
		_ebb += int64(_cfed._a)
		_cfed.FullHeaders = true
	}
	var (
		_aef  *Page
		_edbd bool
	)
	for _egc != 51 && !_edbd {
		_eed, _adc := _gd.NewHeader(_cfed, _cfed.InputStream, _ebb, _cfed.OrganizationType)
		if _adc != nil {
			return _eb.Wrap(_adc, _aaa, "")
		}
		_g.Log.Trace("\u0044\u0065c\u006f\u0064\u0069\u006eg\u0020\u0073e\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c\u0020\u0054\u0079\u0070e\u003a\u0020\u0025\u0073", _eed.SegmentNumber, _eed.Type)
		_egc = _eed.Type
		if _egc != _gd.TEndOfFile {
			if _eed.PageAssociation != 0 {
				_aef = _cfed.Pages[_eed.PageAssociation]
				if _aef == nil {
					_aef = _fde(_cfed, _eed.PageAssociation)
					_cfed.Pages[_eed.PageAssociation] = _aef
					if _cfed.NumberOfPagesUnknown {
						_cfed.NumberOfPages++
					}
				}
				_aef.Segments = append(_aef.Segments, _eed)
			} else {
				_cfed.GlobalSegments.AddSegment(_eed)
			}
		}
		_acd = append(_acd, _eed)
		_ebb = _cfed.InputStream.AbsolutePosition()
		if _cfed.OrganizationType == _gd.OSequential {
			_ebb += int64(_eed.SegmentDataLength)
		}
		_edbd, _adc = _cfed.reachedEOF(_ebb)
		if _adc != nil {
			_g.Log.Debug("\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0072\u0065\u0061\u0063h\u0065\u0064\u0020\u0045\u004f\u0046\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _adc)
			return _eb.Wrap(_adc, _aaa, "")
		}
	}
	_cfed.determineRandomDataOffsets(_acd, uint64(_ebb))
	return nil
}
func (_fbb *Page) nextSegmentNumber() uint32 { return _fbb.Document.nextSegmentNumber() }
func (_fffee *Page) lastSegmentNumber() (_eadg uint32, _daf error) {
	const _fccgb = "\u006c\u0061\u0073\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	if len(_fffee.Segments) == 0 {
		return _eadg, _eb.Errorf(_fccgb, "\u006e\u006f\u0020se\u0067\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006fu\u006ed\u0020i\u006e \u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027", _fffee.PageNumber)
	}
	return _fffee.Segments[len(_fffee.Segments)-1].SegmentNumber, nil
}
func (_dfb *Page) GetWidth() (int, error)       { return _dfb.getWidth() }
func (_fba *Page) GetResolutionX() (int, error) { return _fba.getResolutionX() }
func (_bead *Page) addTextRegionSegment(_fbc []*_gd.Header, _beae, _bagf map[int]int, _bdc []int, _faf *_gc.Points, _eeba *_gc.Bitmaps, _geg *_e.IntSlice, _gedg *_gc.Boxes, _gdg, _eac int) {
	_cdgd := &_gd.TextRegion{NumberOfSymbols: uint32(_eac)}
	_cdgd.InitEncode(_beae, _bagf, _bdc, _faf, _eeba, _geg, _gedg, _bead.FinalWidth, _bead.FinalHeight, _gdg)
	_fccb := &_gd.Header{RTSegments: _fbc, SegmentData: _cdgd, PageAssociation: _bead.PageNumber, Type: _gd.TImmediateTextRegion}
	_abdc := _gd.TPageInformation
	if _bagf != nil {
		_abdc = _gd.TSymbolDictionary
	}
	var _eec int
	for ; _eec < len(_bead.Segments); _eec++ {
		if _bead.Segments[_eec].Type == _abdc {
			_eec++
			break
		}
	}
	_bead.Segments = append(_bead.Segments, nil)
	copy(_bead.Segments[_eec+1:], _bead.Segments[_eec:])
	_bead.Segments[_eec] = _fccb
}
func _edg(_cbg *_f.Reader, _fff *Globals) (*Document, error) {
	_abee := &Document{Pages: make(map[int]*Page), InputStream: _cbg, OrganizationType: _gd.OSequential, NumberOfPagesUnknown: true, GlobalSegments: _fff, _a: 9}
	if _abee.GlobalSegments == nil {
		_abee.GlobalSegments = &Globals{}
	}
	if _dfdd := _abee.mapData(); _dfdd != nil {
		return nil, _dfdd
	}
	return _abee, nil
}
func (_bff *Document) nextSegmentNumber() uint32 {
	_ceg := _bff.CurrentSegmentNumber
	_bff.CurrentSegmentNumber++
	return _ceg
}
func (_gcfg *Globals) AddSegment(segment *_gd.Header) {
	_gcfg.Segments = append(_gcfg.Segments, segment)
}
func (_ddc *Document) GetGlobalSegment(i int) (*_gd.Header, error) {
	_dfda, _bga := _ddc.GlobalSegments.GetSegment(i)
	if _bga != nil {
		return nil, _eb.Wrap(_bga, "\u0047\u0065t\u0047\u006c\u006fb\u0061\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074", "")
	}
	return _dfda, nil
}

type Globals struct{ Segments []*_gd.Header }

func (_fdb *Page) getHeight() (int, error) {
	const _ffgc = "\u0067e\u0074\u0048\u0065\u0069\u0067\u0068t"
	if _fdb.FinalHeight != 0 {
		return _fdb.FinalHeight, nil
	}
	_ead := _fdb.getPageInformationSegment()
	if _ead == nil {
		return 0, _eb.Error(_ffgc, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_dbdd, _bbcf := _ead.GetSegmentData()
	if _bbcf != nil {
		return 0, _eb.Wrap(_bbcf, _ffgc, "")
	}
	_egd, _afg := _dbdd.(*_gd.PageInformationSegment)
	if !_afg {
		return 0, _eb.Errorf(_ffgc, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _dbdd)
	}
	if _egd.PageBMHeight == _be.MaxInt32 {
		_, _bbcf = _fdb.GetBitmap()
		if _bbcf != nil {
			return 0, _eb.Wrap(_bbcf, _ffgc, "")
		}
	} else {
		_fdb.FinalHeight = _egd.PageBMHeight
	}
	return _fdb.FinalHeight, nil
}
func (_ggg *Page) Encode(w _f.BinaryWriter) (_cff int, _ebbd error) {
	const _ecf = "P\u0061\u0067\u0065\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _gaa int
	for _, _edd := range _ggg.Segments {
		if _gaa, _ebbd = _edd.Encode(w); _ebbd != nil {
			return _cff, _eb.Wrap(_ebbd, _ecf, "")
		}
		_cff += _gaa
	}
	return _cff, nil
}
func (_edc *Page) getPageInformationSegment() *_gd.Header {
	for _, _fgf := range _edc.Segments {
		if _fgf.Type == _gd.TPageInformation {
			return _fgf
		}
	}
	_g.Log.Debug("\u0050\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066o\u0072\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0073\u002e", _edc)
	return nil
}
func (_efc *Page) createPage(_ddf *_gd.PageInformationSegment) error {
	var _fac error
	if !_ddf.IsStripe || _ddf.PageBMHeight != -1 {
		_fac = _efc.createNormalPage(_ddf)
	} else {
		_fac = _efc.createStripedPage(_ddf)
	}
	return _fac
}
