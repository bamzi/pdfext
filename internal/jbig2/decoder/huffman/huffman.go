package huffman

import (
	_ec "errors"
	_e "fmt"
	_f "math"
	_dg "strings"

	_de "github.com/bamzi/pdfext/internal/bitwise"
	_db "github.com/bamzi/pdfext/internal/jbig2/internal"
)

func (_ac *FixedSizeTable) RootNode() *InternalNode { return _ac._ccb }

type EncodedTable struct {
	BasicTabler
	_bg *InternalNode
}

func _fe(_ge *Code) *ValueNode { return &ValueNode{_daf: _ge._bgac, _gc: _ge._fgg, _efe: _ge._be} }
func NewCode(prefixLength, rangeLength, rangeLow int32, isLowerRange bool) *Code {
	return &Code{_gag: prefixLength, _bgac: rangeLength, _fgg: rangeLow, _be: isLowerRange, _bc: -1}
}

type OutOfBandNode struct{}

func (_add *StandardTable) RootNode() *InternalNode { return _add._dbgc }
func (_efg *ValueNode) String() string {
	return _e.Sprintf("\u0025\u0064\u002f%\u0064", _efg._daf, _efg._gc)
}

type InternalNode struct {
	_daa int32
	_bad Node
	_ggg Node
}

func (_bdg *InternalNode) append(_eg *Code) (_gfb error) {
	if _eg._gag == 0 {
		return nil
	}
	_dgf := _eg._gag - 1 - _bdg._daa
	if _dgf < 0 {
		return _ec.New("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	_ad := (_eg._bc >> uint(_dgf)) & 0x1
	if _dgf == 0 {
		if _eg._bgac == -1 {
			if _ad == 1 {
				if _bdg._ggg != nil {
					return _e.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _eg)
				}
				_bdg._ggg = _eaa(_eg)
			} else {
				if _bdg._bad != nil {
					return _e.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _eg)
				}
				_bdg._bad = _eaa(_eg)
			}
		} else {
			if _ad == 1 {
				if _bdg._ggg != nil {
					return _e.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _eg)
				}
				_bdg._ggg = _fe(_eg)
			} else {
				if _bdg._bad != nil {
					return _e.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _eg)
				}
				_bdg._bad = _fe(_eg)
			}
		}
	} else {
		if _ad == 1 {
			if _bdg._ggg == nil {
				_bdg._ggg = _cad(_bdg._daa + 1)
			}
			if _gfb = _bdg._ggg.(*InternalNode).append(_eg); _gfb != nil {
				return _gfb
			}
		} else {
			if _bdg._bad == nil {
				_bdg._bad = _cad(_bdg._daa + 1)
			}
			if _gfb = _bdg._bad.(*InternalNode).append(_eg); _gfb != nil {
				return _gfb
			}
		}
	}
	return nil
}

var _ Node = &InternalNode{}

func (_ae *OutOfBandNode) Decode(r *_de.Reader) (int64, error) { return 0, _db.ErrOOB }

var _ Node = &OutOfBandNode{}

type Code struct {
	_gag  int32
	_bgac int32
	_fgg  int32
	_be   bool
	_bc   int32
}

func _eaa(_gad *Code) *OutOfBandNode { return &OutOfBandNode{} }
func NewFixedSizeTable(codeTable []*Code) (*FixedSizeTable, error) {
	_ca := &FixedSizeTable{_ccb: &InternalNode{}}
	if _abb := _ca.InitTree(codeTable); _abb != nil {
		return nil, _abb
	}
	return _ca, nil
}
func (_dd *EncodedTable) InitTree(codeTable []*Code) error {
	_bade(codeTable)
	for _, _fc := range codeTable {
		if _g := _dd._bg.append(_fc); _g != nil {
			return _g
		}
	}
	return nil
}

type FixedSizeTable struct{ _ccb *InternalNode }

var _ Node = &ValueNode{}

func (_ecf *StandardTable) String() string { return _ecf._dbgc.String() + "\u000a" }
func (_afd *InternalNode) Decode(r *_de.Reader) (int64, error) {
	_fed, _bd := r.ReadBit()
	if _bd != nil {
		return 0, _bd
	}
	if _fed == 0 {
		return _afd._bad.Decode(r)
	}
	return _afd._ggg.Decode(r)
}
func (_bf *EncodedTable) RootNode() *InternalNode { return _bf._bg }

type StandardTable struct{ _dbgc *InternalNode }

func _cad(_bdd int32) *InternalNode { return &InternalNode{_daa: _bdd} }

type ValueNode struct {
	_daf int32
	_gc  int32
	_efe bool
}

func (_af *ValueNode) Decode(r *_de.Reader) (int64, error) {
	_ff, _cbb := r.ReadBits(byte(_af._daf))
	if _cbb != nil {
		return 0, _cbb
	}
	if _af._efe {
		_ff = -_ff
	}
	return int64(_af._gc) + int64(_ff), nil
}

type Node interface {
	Decode(_cb *_de.Reader) (int64, error)
	String() string
}

func NewEncodedTable(table BasicTabler) (*EncodedTable, error) {
	_c := &EncodedTable{_bg: &InternalNode{}, BasicTabler: table}
	if _fg := _c.parseTable(); _fg != nil {
		return nil, _fg
	}
	return _c, nil
}
func (_ag *OutOfBandNode) String() string {
	return _e.Sprintf("\u0025\u0030\u00364\u0062", int64(_f.MaxInt64))
}
func (_abg *StandardTable) Decode(r *_de.Reader) (int64, error) { return _abg._dbgc.Decode(r) }

var _efga = make([]Tabler, len(_agce))

type Tabler interface {
	Decode(_feb *_de.Reader) (int64, error)
	InitTree(_aga []*Code) error
	String() string
	RootNode() *InternalNode
}

func _bade(_abca []*Code) {
	var _bef int32
	for _, _fea := range _abca {
		_bef = _gecf(_bef, _fea._gag)
	}
	_dafe := make([]int32, _bef+1)
	for _, _aa := range _abca {
		_dafe[_aa._gag]++
	}
	var _bgb int32
	_fdd := make([]int32, len(_dafe)+1)
	_dafe[0] = 0
	for _gfg := int32(1); _gfg <= int32(len(_dafe)); _gfg++ {
		_fdd[_gfg] = (_fdd[_gfg-1] + (_dafe[_gfg-1])) << 1
		_bgb = _fdd[_gfg]
		for _, _dfec := range _abca {
			if _dfec._gag == _gfg {
				_dfec._bc = _bgb
				_bgb++
			}
		}
	}
}
func (_eb *InternalNode) String() string {
	_cbf := &_dg.Builder{}
	_cbf.WriteString("\u000a")
	_eb.pad(_cbf)
	_cbf.WriteString("\u0030\u003a\u0020")
	_cbf.WriteString(_eb._bad.String() + "\u000a")
	_eb.pad(_cbf)
	_cbf.WriteString("\u0031\u003a\u0020")
	_cbf.WriteString(_eb._ggg.String() + "\u000a")
	return _cbf.String()
}

var _ Tabler = &EncodedTable{}

func (_bga *FixedSizeTable) InitTree(codeTable []*Code) error {
	_bade(codeTable)
	for _, _gg := range codeTable {
		_gf := _bga._ccb.append(_gg)
		if _gf != nil {
			return _gf
		}
	}
	return nil
}
func (_cfb *StandardTable) InitTree(codeTable []*Code) error {
	_bade(codeTable)
	for _, _gd := range codeTable {
		if _fa := _cfb._dbgc.append(_gd); _fa != nil {
			return _fa
		}
	}
	return nil
}
func (_df *EncodedTable) String() string { return _df._bg.String() + "\u000a" }
func _gecf(_dgg, _fae int32) int32 {
	if _dgg > _fae {
		return _dgg
	}
	return _fae
}
func _cef(_dee, _dc int32) string {
	var _bbe int32
	_gff := make([]rune, _dc)
	for _bddg := int32(1); _bddg <= _dc; _bddg++ {
		_bbe = _dee >> uint(_dc-_bddg) & 1
		if _bbe != 0 {
			_gff[_bddg-1] = '1'
		} else {
			_gff[_bddg-1] = '0'
		}
	}
	return string(_gff)
}
func (_ab *EncodedTable) parseTable() error {
	var (
		_cf            []*Code
		_fd, _bgd, _ed int32
		_cc            uint64
		_ba            error
	)
	_ea := _ab.StreamReader()
	_ce := _ab.HtLow()
	for _ce < _ab.HtHigh() {
		_cc, _ba = _ea.ReadBits(byte(_ab.HtPS()))
		if _ba != nil {
			return _ba
		}
		_fd = int32(_cc)
		_cc, _ba = _ea.ReadBits(byte(_ab.HtRS()))
		if _ba != nil {
			return _ba
		}
		_bgd = int32(_cc)
		_cf = append(_cf, NewCode(_fd, _bgd, _ed, false))
		_ce += 1 << uint(_bgd)
	}
	_cc, _ba = _ea.ReadBits(byte(_ab.HtPS()))
	if _ba != nil {
		return _ba
	}
	_fd = int32(_cc)
	_bgd = 32
	_ed = _ab.HtLow() - 1
	_cf = append(_cf, NewCode(_fd, _bgd, _ed, true))
	_cc, _ba = _ea.ReadBits(byte(_ab.HtPS()))
	if _ba != nil {
		return _ba
	}
	_fd = int32(_cc)
	_bgd = 32
	_ed = _ab.HtHigh()
	_cf = append(_cf, NewCode(_fd, _bgd, _ed, false))
	if _ab.HtOOB() == 1 {
		_cc, _ba = _ea.ReadBits(byte(_ab.HtPS()))
		if _ba != nil {
			return _ba
		}
		_fd = int32(_cc)
		_cf = append(_cf, NewCode(_fd, -1, -1, false))
	}
	if _ba = _ab.InitTree(_cf); _ba != nil {
		return _ba
	}
	return nil
}

var _agce = [][][]int32{{{1, 4, 0}, {2, 8, 16}, {3, 16, 272}, {3, 32, 65808}}, {{1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {6, 32, 75}, {6, -1, 0}}, {{8, 8, -256}, {1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {8, 32, -257, 999}, {7, 32, 75}, {6, -1, 0}}, {{1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {5, 32, 76}}, {{7, 8, -255}, {1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {7, 32, -256, 999}, {6, 32, 76}}, {{5, 10, -2048}, {4, 9, -1024}, {4, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {2, 7, 0}, {3, 7, 128}, {3, 8, 256}, {4, 9, 512}, {4, 10, 1024}, {6, 32, -2049, 999}, {6, 32, 2048}}, {{4, 9, -1024}, {3, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {4, 5, 0}, {5, 5, 32}, {5, 6, 64}, {4, 7, 128}, {3, 8, 256}, {3, 9, 512}, {3, 10, 1024}, {5, 32, -1025, 999}, {5, 32, 2048}}, {{8, 3, -15}, {9, 1, -7}, {8, 1, -5}, {9, 0, -3}, {7, 0, -2}, {4, 0, -1}, {2, 1, 0}, {5, 0, 2}, {6, 0, 3}, {3, 4, 4}, {6, 1, 20}, {4, 4, 22}, {4, 5, 38}, {5, 6, 70}, {5, 7, 134}, {6, 7, 262}, {7, 8, 390}, {6, 10, 646}, {9, 32, -16, 999}, {9, 32, 1670}, {2, -1, 0}}, {{8, 4, -31}, {9, 2, -15}, {8, 2, -11}, {9, 1, -7}, {7, 1, -5}, {4, 1, -3}, {3, 1, -1}, {3, 1, 1}, {5, 1, 3}, {6, 1, 5}, {3, 5, 7}, {6, 2, 39}, {4, 5, 43}, {4, 6, 75}, {5, 7, 139}, {5, 8, 267}, {6, 8, 523}, {7, 9, 779}, {6, 11, 1291}, {9, 32, -32, 999}, {9, 32, 3339}, {2, -1, 0}}, {{7, 4, -21}, {8, 0, -5}, {7, 0, -4}, {5, 0, -3}, {2, 2, -2}, {5, 0, 2}, {6, 0, 3}, {7, 0, 4}, {8, 0, 5}, {2, 6, 6}, {5, 5, 70}, {6, 5, 102}, {6, 6, 134}, {6, 7, 198}, {6, 8, 326}, {6, 9, 582}, {6, 10, 1094}, {7, 11, 2118}, {8, 32, -22, 999}, {8, 32, 4166}, {2, -1, 0}}, {{1, 0, 1}, {2, 1, 2}, {4, 0, 4}, {4, 1, 5}, {5, 1, 7}, {5, 2, 9}, {6, 2, 13}, {7, 2, 17}, {7, 3, 21}, {7, 4, 29}, {7, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {5, 0, 5}, {5, 1, 6}, {6, 1, 8}, {7, 0, 10}, {7, 1, 11}, {7, 2, 13}, {7, 3, 17}, {7, 4, 25}, {8, 5, 41}, {8, 32, 73}}, {{1, 0, 1}, {3, 0, 2}, {4, 0, 3}, {5, 0, 4}, {4, 1, 5}, {3, 3, 7}, {6, 1, 15}, {6, 2, 17}, {6, 3, 21}, {6, 4, 29}, {6, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{3, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {3, 0, 2}}, {{7, 4, -24}, {6, 2, -8}, {5, 1, -4}, {4, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {4, 0, 2}, {5, 1, 3}, {6, 2, 5}, {7, 4, 9}, {7, 32, -25, 999}, {7, 32, 25}}}

func (_da *FixedSizeTable) String() string                      { return _da._ccb.String() + "\u000a" }
func (_ga *FixedSizeTable) Decode(r *_de.Reader) (int64, error) { return _ga._ccb.Decode(r) }
func (_cd *Code) String() string {
	var _dgd string
	if _cd._bc != -1 {
		_dgd = _cef(_cd._bc, _cd._gag)
	} else {
		_dgd = "\u003f"
	}
	return _e.Sprintf("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _dgd, _cd._gag, _cd._bgac, _cd._fgg)
}
func GetStandardTable(number int) (Tabler, error) {
	if number <= 0 || number > len(_efga) {
		return nil, _ec.New("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dfea := _efga[number-1]
	if _dfea == nil {
		var _dbe error
		_dfea, _dbe = _fb(_agce[number-1])
		if _dbe != nil {
			return nil, _dbe
		}
		_efga[number-1] = _dfea
	}
	return _dfea, nil
}
func (_bfg *InternalNode) pad(_eae *_dg.Builder) {
	for _dfe := int32(0); _dfe < _bfg._daa; _dfe++ {
		_eae.WriteString("\u0020\u0020\u0020")
	}
}

type BasicTabler interface {
	HtHigh() int32
	HtLow() int32
	StreamReader() *_de.Reader
	HtPS() int32
	HtRS() int32
	HtOOB() int32
}

func (_a *EncodedTable) Decode(r *_de.Reader) (int64, error) { return _a._bg.Decode(r) }
func _fb(_bb [][]int32) (*StandardTable, error) {
	var _ega []*Code
	for _ee := 0; _ee < len(_bb); _ee++ {
		_cbbb := _bb[_ee][0]
		_agb := _bb[_ee][1]
		_agc := _bb[_ee][2]
		var _gdg bool
		if len(_bb[_ee]) > 3 {
			_gdg = true
		}
		_ega = append(_ega, NewCode(_cbbb, _agb, _agc, _gdg))
	}
	_fdc := &StandardTable{_dbgc: _cad(0)}
	if _adc := _fdc.InitTree(_ega); _adc != nil {
		return nil, _adc
	}
	return _fdc, nil
}
