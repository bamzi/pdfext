package redactor

import (
	_cd "errors"
	_d "fmt"
	_dg "io"
	_db "regexp"
	_a "sort"
	_df "strings"

	_f "github.com/bamzi/pdfext/common"
	_dd "github.com/bamzi/pdfext/contentstream"
	_dbb "github.com/bamzi/pdfext/core"
	_ab "github.com/bamzi/pdfext/creator"
	_bf "github.com/bamzi/pdfext/extractor"
	_b "github.com/bamzi/pdfext/model"
)

func _fdgc(_fdc *_bf.TextMarkArray) []*_bf.TextMarkArray {
	_edaa := _fdc.Elements()
	_cbfc := len(_edaa)
	var _eed _dbb.PdfObject
	_egb := []*_bf.TextMarkArray{}
	_edb := &_bf.TextMarkArray{}
	_aaba := -1
	for _fdff, _daeg := range _edaa {
		_adbd := _daeg.DirectObject
		_aaba = _daeg.Index
		if _adbd == nil {
			_ece := _bbb(_fdc, _fdff, _aaba)
			if _eed != nil {
				if _ece == -1 || _ece > _fdff {
					_egb = append(_egb, _edb)
					_edb = &_bf.TextMarkArray{}
				}
			}
		} else if _adbd != nil && _eed == nil {
			if _aaba == 0 && _fdff > 0 {
				_egb = append(_egb, _edb)
				_edb = &_bf.TextMarkArray{}
			}
		} else if _adbd != nil && _eed != nil {
			if _adbd != _eed {
				_egb = append(_egb, _edb)
				_edb = &_bf.TextMarkArray{}
			}
		}
		_eed = _adbd
		_edb.Append(_daeg)
		if _fdff == (_cbfc - 1) {
			_egb = append(_egb, _edb)
		}
	}
	return _egb
}

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_caeg *Redactor) Write(writer _dg.Writer) error { return _caeg._agff.Write(writer) }

type targetMap struct {
	_gab string
	_bcd [][]int
}

func _abb(_cdff *_bf.TextMarkArray) (float64, error) {
	_aff, _agaf := _cdff.BBox()
	if !_agaf {
		return 0.0, _d.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _cdff)
	}
	_gce := _abd(_cdff)
	_fbce := 0.0
	_, _bcg := _dcf(_cdff)
	_dbfc := _cdff.Elements()[_bcg]
	_bfef := _dbfc.Font
	if _gce > 0 {
		_fbce = _gcf(_bfef, _dbfc)
	}
	_acg := (_aff.Urx - _aff.Llx)
	_acg = _acg + _fbce*float64(_gce)
	Tj := _ddgd(_acg, _dbfc.FontSize, _dbfc.Th)
	return Tj, nil
}
func _abd(_fbc *_bf.TextMarkArray) int {
	_ddg := 0
	_egg := _fbc.Elements()
	if _egg[0].Text == "\u0020" {
		_ddg++
	}
	if _egg[_fbc.Len()-1].Text == "\u0020" {
		_ddg++
	}
	return _ddg
}
func _bda(_adcf _dbb.PdfObject, _caef *_b.PdfFont) (string, error) {
	_facc, _fgbf := _dbb.GetStringBytes(_adcf)
	if !_fgbf {
		return "", _dbb.ErrTypeError
	}
	_bgf := _caef.BytesToCharcodes(_facc)
	_ddd, _cgd, _bbddd := _caef.CharcodesToStrings(_bgf, "")
	if _bbddd > 0 {
		_f.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _cgd, _bbddd)
	}
	_agg := _df.Join(_ddd, "")
	return _agg, nil
}
func _dgc(_ead string, _de *_b.PdfFont) []byte {
	_aag, _aab := _de.StringToCharcodeBytes(_ead)
	if _aab != 0 {
		_f.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076", _ead, _aag)
	}
	return _aag
}
func _aabc(_cade *targetMap, _gba []int) {
	var _dcge [][]int
	for _fbf, _bef := range _cade._bcd {
		if _dabgg(_fbf, _gba) {
			continue
		}
		_dcge = append(_dcge, _bef)
	}
	_cade._bcd = _dcge
}

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{ Terms []RedactionTerm }

func _baaf(_efd string, _edde []localSpanMarks) ([]placeHolders, error) {
	_bgca := ""
	_gec := []placeHolders{}
	for _gecb, _fba := range _edde {
		_cccd := _fba._gaf
		_fde := _fba._gcfa
		_dca := _ebeb(_cccd)
		_cbf, _dac := _abb(_cccd)
		if _dac != nil {
			return nil, _dac
		}
		if _dca != _bgca {
			var _bea []int
			if _gecb == 0 && _fde != _dca {
				_agf := _df.Index(_efd, _dca)
				_bea = []int{_agf}
			} else if _gecb == len(_edde)-1 {
				_fbd := _df.LastIndex(_efd, _dca)
				_bea = []int{_fbd}
			} else {
				_bea = _abec(_efd, _dca)
			}
			_ade := placeHolders{_g: _bea, _dfa: _dca, _e: _cbf}
			_gec = append(_gec, _ade)
		}
		_bgca = _dca
	}
	return _gec, nil
}
func _cbd(_aaea []int, _ggd *_bf.TextMarkArray, _aedc string) (*_bf.TextMarkArray, matchedBBox, error) {
	_bbcb := matchedBBox{}
	_gdba := _aaea[0]
	_dec := _aaea[1]
	_cebf := len(_aedc) - len(_df.TrimLeft(_aedc, "\u0020"))
	_baafb := len(_aedc) - len(_df.TrimRight(_aedc, "\u0020\u000a"))
	_gdba = _gdba + _cebf
	_dec = _dec - _baafb
	_aedc = _df.Trim(_aedc, "\u0020\u000a")
	_cfb, _faf := _ggd.RangeOffset(_gdba, _dec)
	if _faf != nil {
		return nil, _bbcb, _faf
	}
	_gac, _bbgd := _cfb.BBox()
	if !_bbgd {
		return nil, _bbcb, _d.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _cfb)
	}
	_bbcb = matchedBBox{_dab: _aedc, _eee: _gac}
	return _cfb, _bbcb, nil
}
func _ffa(_ccg, _cfab targetMap) (bool, []int) {
	_gbd := _df.Contains(_ccg._gab, _cfab._gab)
	var _ccgc []int
	for _, _fgbd := range _ccg._bcd {
		for _agbe, _adb := range _cfab._bcd {
			if _adb[0] >= _fgbd[0] && _adb[1] <= _fgbd[1] {
				_ccgc = append(_ccgc, _agbe)
			}
		}
	}
	return _gbd, _ccgc
}
func _bbe(_fe localSpanMarks, _bc *_bf.TextMarkArray, _bad *_b.PdfFont, _baa, _bega string) ([]_dbb.PdfObject, error) {
	_cegc := _ebeb(_bc)
	Tj, _cg := _abb(_bc)
	if _cg != nil {
		return nil, _cg
	}
	_edee := len(_baa)
	_dbge := len(_cegc)
	_faa := -1
	_ddfa := _dbb.MakeFloat(Tj)
	if _cegc != _bega {
		_dfe := _fe._agae
		if _dfe == 0 {
			_faa = _df.LastIndex(_baa, _cegc)
		} else {
			_faa = _df.Index(_baa, _cegc)
		}
	} else {
		_faa = _df.Index(_baa, _cegc)
	}
	_abg := _faa + _dbge
	_dbdc := []_dbb.PdfObject{}
	if _faa == 0 && _abg == _edee {
		_dbdc = append(_dbdc, _ddfa)
	} else if _faa == 0 && _abg < _edee {
		_geg := _dgc(_baa[_abg:], _bad)
		_ffb := _dbb.MakeStringFromBytes(_geg)
		_dbdc = append(_dbdc, _ddfa, _ffb)
	} else if _faa > 0 && _abg >= _edee {
		_afg := _dgc(_baa[:_faa], _bad)
		_ec := _dbb.MakeStringFromBytes(_afg)
		_dbdc = append(_dbdc, _ec, _ddfa)
	} else if _faa > 0 && _abg < _edee {
		_dfee := _dgc(_baa[:_faa], _bad)
		_abda := _dgc(_baa[_abg:], _bad)
		_bebc := _dbb.MakeStringFromBytes(_dfee)
		_ada := _dbb.MakeString(string(_abda))
		_dbdc = append(_dbdc, _bebc, _ddfa, _ada)
	}
	return _dbdc, nil
}
func _bcc(_gbc []placeHolders) []replacement {
	_adc := []replacement{}
	for _, _dae := range _gbc {
		_daeb := _dae._g
		_bbg := _dae._dfa
		_ebgb := _dae._e
		for _, _adg := range _daeb {
			_bed := replacement{_ed: _bbg, _fa: _ebgb, _bfe: _adg}
			_adc = append(_adc, _bed)
		}
	}
	_a.Slice(_adc, func(_ebd, _dggc int) bool { return _adc[_ebd]._bfe < _adc[_dggc]._bfe })
	return _adc
}

type localSpanMarks struct {
	_gaf  *_bf.TextMarkArray
	_agae int
	_gcfa string
}

func _ebeb(_bde *_bf.TextMarkArray) string {
	_fce := ""
	for _, _cgdd := range _bde.Elements() {
		_fce += _cgdd.Text
	}
	return _fce
}
func _gcbc(_deef *matchedIndex, _agbg [][]int) (bool, [][]int) {
	_acgf := [][]int{}
	for _, _dbfg := range _agbg {
		if _deef._cfac < _dbfg[0] && _deef._eef > _dbfg[1] {
			_acgf = append(_acgf, _dbfg)
		}
	}
	return len(_acgf) > 0, _acgf
}

// Redactor represents a Redactor object.
type Redactor struct {
	_dda  *_b.PdfReader
	_gbb  *RedactionOptions
	_agff *_ab.Creator
	_caa  *RectangleProps
}

func _cdg(_ccd *_dd.ContentStreamOperations, _ce string, _dc int) error {
	_cf := _dd.ContentStreamOperations{}
	var _cfa _dd.ContentStreamOperation
	for _bb, _ef := range *_ccd {
		if _bb == _dc {
			if _ce == "\u0027" {
				_ee := _dd.ContentStreamOperation{Operand: "\u0054\u002a"}
				_cf = append(_cf, &_ee)
				_cfa.Params = _ef.Params
				_cfa.Operand = "\u0054\u006a"
				_cf = append(_cf, &_cfa)
			} else if _ce == "\u0022" {
				_cde := _ef.Params[:2]
				Tc, Tw := _cde[0], _cde[1]
				_ga := _dd.ContentStreamOperation{Params: []_dbb.PdfObject{Tc}, Operand: "\u0054\u0063"}
				_cf = append(_cf, &_ga)
				_ga = _dd.ContentStreamOperation{Params: []_dbb.PdfObject{Tw}, Operand: "\u0054\u0077"}
				_cf = append(_cf, &_ga)
				_cfa.Params = []_dbb.PdfObject{_ef.Params[2]}
				_cfa.Operand = "\u0054\u006a"
				_cf = append(_cf, &_cfa)
			}
		}
		_cf = append(_cf, _ef)
	}
	*_ccd = _cf
	return nil
}

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_gcd *Redactor) WriteToFile(outputPath string) error {
	if _cdb := _gcd._agff.WriteToFile(outputPath); _cdb != nil {
		return _d.Errorf("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065")
	}
	return nil
}
func (_gabe *regexMatcher) match(_bdgg string) ([]*matchedIndex, error) {
	_becg := _gabe._bgg.Pattern
	if _becg == nil {
		return nil, _cd.New("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064")
	}
	var (
		_gde  = _becg.FindAllStringIndex(_bdgg, -1)
		_afab []*matchedIndex
	)
	for _, _faba := range _gde {
		_afab = append(_afab, &matchedIndex{_cfac: _faba[0], _eef: _faba[1], _efa: _bdgg[_faba[0]:_faba[1]]})
	}
	return _afab, nil
}
func _beb(_def *_bf.TextMarkArray) *_b.PdfFont {
	_, _ddff := _dcf(_def)
	_egae := _def.Elements()[_ddff]
	_dbf := _egae.Font
	return _dbf
}
func _ddgd(_gfc, _afbf, _age float64) float64 {
	_age = _age / 100
	_efbe := (-1000 * _gfc) / (_afbf * _age)
	return _efbe
}
func _ace(_gfce *matchedIndex, _gcdc [][]int) []*matchedIndex {
	_cda := []*matchedIndex{}
	_dcec := _gfce._cfac
	_cfaa := _dcec
	_agfb := _gfce._efa
	_gagg := 0
	for _, _befd := range _gcdc {
		_ecb := _befd[0] - _dcec
		if _gagg >= _ecb {
			continue
		}
		_fbcb := _agfb[_gagg:_ecb]
		_gaba := &matchedIndex{_efa: _fbcb, _cfac: _cfaa, _eef: _befd[0]}
		if len(_df.TrimSpace(_fbcb)) != 0 {
			_cda = append(_cda, _gaba)
		}
		_gagg = _befd[1] - _dcec
		_cfaa = _dcec + _gagg
	}
	_bdc := _agfb[_gagg:]
	_cfea := &matchedIndex{_efa: _bdc, _cfac: _cfaa, _eef: _gfce._eef}
	if len(_df.TrimSpace(_bdc)) != 0 {
		_cda = append(_cda, _cfea)
	}
	return _cda
}
func _dabgg(_afgcf int, _bfbf []int) bool {
	for _, _ddac := range _bfbf {
		if _ddac == _afgcf {
			return true
		}
	}
	return false
}

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_ffcf *Redactor) Redact() error {
	_bag, _aged := _ffcf._dda.GetNumPages()
	if _aged != nil {
		return _d.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073")
	}
	_gegb := _ffcf._caa.FillColor
	_gfd := _ffcf._caa.BorderWidth
	_cea := _ffcf._caa.FillOpacity
	for _fdge := 1; _fdge <= _bag; _fdge++ {
		_cbgc, _bcb := _ffcf._dda.GetPage(_fdge)
		if _bcb != nil {
			return _bcb
		}
		_ged, _bcb := _bf.New(_cbgc)
		if _bcb != nil {
			return _bcb
		}
		_aaa, _, _, _bcb := _ged.ExtractPageText()
		if _bcb != nil {
			return _bcb
		}
		_abff := _aaa.GetContentStreamOps()
		_aed, _dabg, _bcb := _ffcf.redactPage(_abff, _cbgc.Resources)
		if _dabg == nil {
			_f.Log.Info("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e")
			_dabg = _abff
		}
		_eeb := _dd.ContentStreamOperation{Operand: "\u006e"}
		*_dabg = append(*_dabg, &_eeb)
		_cbgc.SetContentStreams([]string{_dabg.String()}, _dbb.NewFlateEncoder())
		if _bcb != nil {
			return _bcb
		}
		_afa, _bcb := _cbgc.GetMediaBox()
		if _bcb != nil {
			return _bcb
		}
		if _cbgc.MediaBox == nil {
			_cbgc.MediaBox = _afa
		}
		if _gega := _ffcf._agff.AddPage(_cbgc); _gega != nil {
			return _gega
		}
		_a.Slice(_aed, func(_dcg, _bae int) bool { return _aed[_dcg]._dab < _aed[_bae]._dab })
		_agfd := _afa.Ury
		for _, _gfg := range _aed {
			_ggca := _gfg._eee
			_agad := _ffcf._agff.NewRectangle(_ggca.Llx, _agfd-_ggca.Lly, _ggca.Urx-_ggca.Llx, -(_ggca.Ury - _ggca.Lly))
			_agad.SetFillColor(_gegb)
			_agad.SetBorderWidth(_gfd)
			_agad.SetFillOpacity(_cea)
			if _aae := _ffcf._agff.Draw(_agad); _aae != nil {
				return nil
			}
		}
	}
	_ffcf._agff.SetOutlineTree(_ffcf._dda.GetOutlineTree())
	return nil
}
func _bca(_bfc string) (string, [][]int) {
	_fcd := _db.MustCompile("\u005c\u006e")
	_gbfg := _fcd.FindAllStringIndex(_bfc, -1)
	_cag := _fcd.ReplaceAllString(_bfc, "\u0020")
	return _cag, _gbfg
}

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct {
	FillColor   _ab.Color
	BorderWidth float64
	FillOpacity float64
}

func _dfcb(_faaa string, _gca []replacement, _cfda *_b.PdfFont) []_dbb.PdfObject {
	_aee := []_dbb.PdfObject{}
	_ebee := 0
	_acf := _faaa
	for _gge, _gee := range _gca {
		_afgc := _gee._bfe
		_cae := _gee._fa
		_bec := _gee._ed
		_bfd := _dbb.MakeFloat(_cae)
		if _ebee > _afgc || _afgc == -1 {
			continue
		}
		_agfc := _faaa[_ebee:_afgc]
		_agfa := _dgc(_agfc, _cfda)
		_cab := _dbb.MakeStringFromBytes(_agfa)
		_aee = append(_aee, _cab)
		_aee = append(_aee, _bfd)
		_dbbaa := _afgc + len(_bec)
		_acf = _faaa[_dbbaa:]
		_ebee = _dbbaa
		if _gge == len(_gca)-1 {
			_agfa = _dgc(_acf, _cfda)
			_cab = _dbb.MakeStringFromBytes(_agfa)
			_aee = append(_aee, _cab)
		}
	}
	return _aee
}
func _bbb(_dcde *_bf.TextMarkArray, _cca int, _bcbf int) int {
	_ffcg := _dcde.Elements()
	_ccca := _cca - 1
	_afe := _cca + 1
	_aadd := -1
	if _ccca >= 0 {
		_fadg := _ffcg[_ccca]
		_afc := _fadg.ObjString
		_bcfb := len(_afc)
		_geaa := _fadg.Index
		if _geaa+1 < _bcfb {
			_aadd = _ccca
			return _aadd
		}
	}
	if _afe < len(_ffcg) {
		_dgcd := _ffcg[_afe]
		_bbgf := _dgcd.ObjString
		if _bbgf[0] != _dgcd.Text {
			_aadd = _afe
			return _aadd
		}
	}
	return _aadd
}

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{ Pattern *_db.Regexp }

func _ccdd(_eade []*targetMap) {
	for _bga, _afbd := range _eade {
		for _dge, _gea := range _eade {
			if _bga != _dge {
				_ebgf, _gag := _ffa(*_afbd, *_gea)
				if _ebgf {
					_aabc(_gea, _gag)
				}
			}
		}
	}
}
func _dcf(_cfd *_bf.TextMarkArray) (_dbb.PdfObject, int) {
	var _cfdg _dbb.PdfObject
	_fbg := -1
	for _eff, _eec := range _cfd.Elements() {
		_cfdg = _eec.DirectObject
		_fbg = _eff
		if _cfdg != nil {
			break
		}
	}
	return _cfdg, _fbg
}
func (_gbf *Redactor) redactPage(_gcb *_dd.ContentStreamOperations, _cdc *_b.PdfPageResources) ([]matchedBBox, *_dd.ContentStreamOperations, error) {
	_cdga, _abc := _bf.NewFromContents(_gcb.String(), _cdc)
	if _abc != nil {
		return nil, nil, _abc
	}
	_bfb, _, _, _abc := _cdga.ExtractPageText()
	if _abc != nil {
		return nil, nil, _abc
	}
	_gcb = _bfb.GetContentStreamOps()
	_fae := _bfb.Marks()
	_feb := _bfb.Text()
	_feb, _ddc := _bca(_feb)
	_dbbb := []matchedBBox{}
	_fcf := make(map[_dbb.PdfObject][]localSpanMarks)
	_dgff := []*targetMap{}
	for _, _gege := range _gbf._gbb.Terms {
		_dcfc, _bbca := _eegd(_gege)
		if _bbca != nil {
			return nil, nil, _bbca
		}
		_fdgf, _bbca := _dcfc.match(_feb)
		if _bbca != nil {
			return nil, nil, _bbca
		}
		_fdgf = _dgee(_fdgf, _ddc)
		_fab := _cbgg(_fdgf)
		_dgff = append(_dgff, _fab...)
	}
	_ccdd(_dgff)
	for _, _aeed := range _dgff {
		_gbbc := _aeed._gab
		_aeb := _aeed._bcd
		_ddb := []matchedBBox{}
		for _, _dccb := range _aeb {
			_gfca, _caac, _edf := _cbd(_dccb, _fae, _gbbc)
			if _edf != nil {
				return nil, nil, _edf
			}
			_geba := _fdgc(_gfca)
			for _dgfd, _dabd := range _geba {
				_cbc := localSpanMarks{_gaf: _dabd, _agae: _dgfd, _gcfa: _gbbc}
				_gfcf, _ := _dcf(_dabd)
				if _cdd, _adcfe := _fcf[_gfcf]; _adcfe {
					_fcf[_gfcf] = append(_cdd, _cbc)
				} else {
					_fcf[_gfcf] = []localSpanMarks{_cbc}
				}
			}
			_ddb = append(_ddb, _caac)
		}
		_dbbb = append(_dbbb, _ddb...)
	}
	_abc = _cc(_gcb, _fcf)
	if _abc != nil {
		return nil, nil, _abc
	}
	return _dbbb, _gcb, nil
}
func _cga(_bbcf *_dd.ContentStreamOperations, PdfObj _dbb.PdfObject) (*_dd.ContentStreamOperation, int, bool) {
	for _affb, _eda := range *_bbcf {
		_gcef := _eda.Operand
		if _gcef == "\u0054\u006a" {
			_bbdd := _dbb.TraceToDirectObject(_eda.Params[0])
			if _bbdd == PdfObj {
				return _eda, _affb, true
			}
		} else if _gcef == "\u0054\u004a" {
			_ccf, _fcb := _dbb.GetArray(_eda.Params[0])
			if !_fcb {
				return nil, _affb, _fcb
			}
			for _, _aeeb := range _ccf.Elements() {
				if _aeeb == PdfObj {
					return _eda, _affb, true
				}
			}
		} else if _gcef == "\u0022" {
			_aabb := _dbb.TraceToDirectObject(_eda.Params[2])
			if _aabb == PdfObj {
				return _eda, _affb, true
			}
		} else if _gcef == "\u0027" {
			_efb := _dbb.TraceToDirectObject(_eda.Params[0])
			if _efb == PdfObj {
				return _eda, _affb, true
			}
		}
	}
	return nil, -1, false
}
func _gcf(_ggce *_b.PdfFont, _ede _bf.TextMark) float64 {
	_dbba := 0.001
	_fdb := _ede.Th / 100
	if _ggce.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_dbba = 1
	}
	_gd, _fc := _ggce.GetRuneMetrics(' ')
	if !_fc {
		_gd, _fc = _ggce.GetCharMetrics(32)
	}
	if !_fc {
		_gd, _ = _b.DefaultFont().GetRuneMetrics(' ')
	}
	_bge := _dbba * ((_gd.Wx*_ede.FontSize + _ede.Tc + _ede.Tw) / _fdb)
	return _bge
}

type matchedIndex struct {
	_cfac int
	_eef  int
	_efa  string
}

func _dba(_ba []localSpanMarks) (map[string][]localSpanMarks, []string) {
	_be := make(map[string][]localSpanMarks)
	_ceb := []string{}
	for _, _gdb := range _ba {
		_beg := _gdb._gcfa
		if _fdg, _gaa := _be[_beg]; _gaa {
			_be[_beg] = append(_fdg, _gdb)
		} else {
			_be[_beg] = []localSpanMarks{_gdb}
			_ceb = append(_ceb, _beg)
		}
	}
	return _be, _ceb
}

type replacement struct {
	_ed  string
	_fa  float64
	_bfe int
}

func _cbgg(_cbgge []*matchedIndex) []*targetMap {
	_bbef := make(map[string][][]int)
	_bgfg := []*targetMap{}
	for _, _fdbc := range _cbgge {
		_faee := _fdbc._efa
		_bcca := []int{_fdbc._cfac, _fdbc._eef}
		if _aeea, _dbdcd := _bbef[_faee]; _dbdcd {
			_bbef[_faee] = append(_aeea, _bcca)
		} else {
			_bbef[_faee] = [][]int{_bcca}
		}
	}
	for _eebgf, _adbg := range _bbef {
		_gebg := &targetMap{_gab: _eebgf, _bcd: _adbg}
		_bgfg = append(_bgfg, _gebg)
	}
	return _bgfg
}

type matchedBBox struct {
	_eee _b.PdfRectangle
	_dab string
}

func _cc(_bg *_dd.ContentStreamOperations, _fb map[_dbb.PdfObject][]localSpanMarks) error {
	for _gc, _fbe := range _fb {
		if _gc == nil {
			continue
		}
		_gg, _ggc, _bgc := _cga(_bg, _gc)
		if !_bgc {
			_f.Log.Debug("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s", _gc, _bg)
			return nil
		}
		if _gg.Operand == "\u0054\u006a" {
			_dgf := _cebg(_gg, _gc, _fbe)
			if _dgf != nil {
				return _dgf
			}
		} else if _gg.Operand == "\u0054\u004a" {
			_eg := _aa(_gg, _gc, _fbe)
			if _eg != nil {
				return _eg
			}
		} else if _gg.Operand == "\u0027" || _gg.Operand == "\u0022" {
			_dde := _cdg(_bg, _gg.Operand, _ggc)
			if _dde != nil {
				return _dde
			}
			_dde = _cebg(_gg, _gc, _fbe)
			if _dde != nil {
				return _dde
			}
		}
	}
	return nil
}

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New(reader *_b.PdfReader, opts *RedactionOptions, rectProps *RectangleProps) *Redactor {
	if rectProps == nil {
		rectProps = RedactRectanglePropsNew()
	}
	return &Redactor{_dda: reader, _gbb: opts, _agff: _ab.New(), _caa: rectProps}
}
func _eegd(_gfcd RedactionTerm) (*regexMatcher, error) { return &regexMatcher{_bgg: _gfcd}, nil }
func _aa(_ag *_dd.ContentStreamOperation, _dga _dbb.PdfObject, _ca []localSpanMarks) error {
	_dcd, _ad := _dbb.GetArray(_ag.Params[0])
	_bbc := []_dbb.PdfObject{}
	if !_ad {
		_f.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _ag)
		return _d.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _ag)
	}
	_ddf, _eb := _dba(_ca)
	if len(_eb) == 1 {
		_dgg := _eb[0]
		_ebg := _ddf[_dgg]
		if len(_ebg) == 1 {
			_bbd := _ebg[0]
			_egd := _bbd._gaf
			_cad := _beb(_egd)
			_abf, _ge := _bda(_dga, _cad)
			if _ge != nil {
				return _ge
			}
			_cdf, _ge := _bbe(_bbd, _egd, _cad, _abf, _dgg)
			if _ge != nil {
				return _ge
			}
			for _, _cb := range _dcd.Elements() {
				if _cb == _dga {
					_bbc = append(_bbc, _cdf...)
				} else {
					_bbc = append(_bbc, _cb)
				}
			}
		} else {
			_ac := _ebg[0]._gaf
			_aga := _beb(_ac)
			_af, _abe := _bda(_dga, _aga)
			if _abe != nil {
				return _abe
			}
			_da, _abe := _baaf(_af, _ebg)
			if _abe != nil {
				return _abe
			}
			_fac := _bcc(_da)
			_afb := _dfcb(_af, _fac, _aga)
			for _, _dbd := range _dcd.Elements() {
				if _dbd == _dga {
					_bbc = append(_bbc, _afb...)
				} else {
					_bbc = append(_bbc, _dbd)
				}
			}
		}
		_ag.Params[0] = _dbb.MakeArray(_bbc...)
	} else if len(_eb) > 1 {
		_gf := _ca[0]
		_agb := _gf._gaf
		_, _ea := _dcf(_agb)
		_edc := _agb.Elements()[_ea]
		_eea := _edc.Font
		_dgfb, _ff := _bda(_dga, _eea)
		if _ff != nil {
			return _ff
		}
		_fd, _ff := _baaf(_dgfb, _ca)
		if _ff != nil {
			return _ff
		}
		_bfa := _bcc(_fd)
		_gb := _dfcb(_dgfb, _bfa, _eea)
		for _, _cbg := range _dcd.Elements() {
			if _cbg == _dga {
				_bbc = append(_bbc, _gb...)
			} else {
				_bbc = append(_bbc, _cbg)
			}
		}
		_ag.Params[0] = _dbb.MakeArray(_bbc...)
	}
	return nil
}
func _dgee(_fcg []*matchedIndex, _edaab [][]int) []*matchedIndex {
	_aabcc := []*matchedIndex{}
	for _, _bbeg := range _fcg {
		_badc, _dag := _gcbc(_bbeg, _edaab)
		if _badc {
			_gacf := _ace(_bbeg, _dag)
			_aabcc = append(_aabcc, _gacf...)
		} else {
			_aabcc = append(_aabcc, _bbeg)
		}
	}
	return _aabcc
}
func _abec(_ddec, _dcca string) []int {
	if len(_dcca) == 0 {
		return nil
	}
	var _bcf []int
	for _fag := 0; _fag < len(_ddec); {
		_fbba := _df.Index(_ddec[_fag:], _dcca)
		if _fbba < 0 {
			return _bcf
		}
		_bcf = append(_bcf, _fag+_fbba)
		_fag += _fbba + len(_dcca)
	}
	return _bcf
}

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew() *RectangleProps {
	return &RectangleProps{FillColor: _ab.ColorBlack, BorderWidth: 0.0, FillOpacity: 1.0}
}

type regexMatcher struct{ _bgg RedactionTerm }
type placeHolders struct {
	_g   []int
	_dfa string
	_e   float64
}

func _cebg(_dbg *_dd.ContentStreamOperation, _fbb _dbb.PdfObject, _dgag []localSpanMarks) error {
	var _dfc *_dbb.PdfObjectArray
	_cfg, _cfec := _dba(_dgag)
	if len(_cfec) == 1 {
		_dce := _cfec[0]
		_bd := _cfg[_dce]
		if len(_bd) == 1 {
			_fgb := _bd[0]
			_egde := _fgb._gaf
			_eggd := _beb(_egde)
			_aad, _fad := _bda(_fbb, _eggd)
			if _fad != nil {
				return _fad
			}
			_dbe, _fad := _bbe(_fgb, _egde, _eggd, _aad, _dce)
			if _fad != nil {
				return _fad
			}
			_dfc = _dbb.MakeArray(_dbe...)
		} else {
			_ebe := _bd[0]._gaf
			_ae := _beb(_ebe)
			_bgcd, _dcc := _bda(_fbb, _ae)
			if _dcc != nil {
				return _dcc
			}
			_gbg, _dcc := _baaf(_bgcd, _bd)
			if _dcc != nil {
				return _dcc
			}
			_dee := _bcc(_gbg)
			_aaf := _dfcb(_bgcd, _dee, _ae)
			_dfc = _dbb.MakeArray(_aaf...)
		}
	} else if len(_cfec) > 1 {
		_ebgc := _dgag[0]
		_ccc := _ebgc._gaf
		_, _daa := _dcf(_ccc)
		_ceg := _ccc.Elements()[_daa]
		_ggg := _ceg.Font
		_dbfa, _bdg := _bda(_fbb, _ggg)
		if _bdg != nil {
			return _bdg
		}
		_ced, _bdg := _baaf(_dbfa, _dgag)
		if _bdg != nil {
			return _bdg
		}
		_eeg := _bcc(_ced)
		_edd := _dfcb(_dbfa, _eeg, _ggg)
		_dfc = _dbb.MakeArray(_edd...)
	}
	_dbg.Params[0] = _dfc
	_dbg.Operand = "\u0054\u004a"
	return nil
}
