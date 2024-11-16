package basic

import _b "github.com/bamzi/pdfext/internal/jbig2/errors"

func (_fb IntsMap) Delete(key uint64) { delete(_fb, key) }
func NewIntSlice(i int) *IntSlice     { _fg := IntSlice(make([]int, i)); return &_fg }
func (_db *NumSlice) AddInt(v int)    { *_db = append(*_db, float32(v)) }

type Stack struct {
	Data []interface{}
	Aux  *Stack
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Ceil(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}
func (_fa NumSlice) GetInt(i int) (int, error) {
	const _bd = "\u0047\u0065\u0074\u0049\u006e\u0074"
	if i < 0 || i > len(_fa)-1 {
		return 0, _b.Errorf(_bd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_ad := _fa[i]
	return int(_ad + Sign(_ad)*0.5), nil
}
func (_d IntsMap) Add(key uint64, value int) { _d[key] = append(_d[key], value) }
func (_dfc *Stack) Push(v interface{})       { _dfc.Data = append(_dfc.Data, v) }
func (_dab *Stack) Len() int                 { return len(_dab.Data) }
func Sign(v float32) float32 {
	if v >= 0.0 {
		return 1.0
	}
	return -1.0
}
func (_c IntsMap) Get(key uint64) (int, bool) {
	_e, _f := _c[key]
	if !_f {
		return 0, false
	}
	if len(_e) == 0 {
		return 0, false
	}
	return _e[0], true
}
func (_adc *Stack) top() int { return len(_adc.Data) - 1 }
func (_eg *Stack) Pop() (_ggd interface{}, _dc bool) {
	_ggd, _dc = _eg.peek()
	if !_dc {
		return nil, _dc
	}
	_eg.Data = _eg.Data[:_eg.top()]
	return _ggd, true
}
func (_bg *IntSlice) Add(v int) error {
	if _bg == nil {
		return _b.Error("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064", "\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	*_bg = append(*_bg, v)
	return nil
}
func (_g NumSlice) GetIntSlice() []int {
	_dag := make([]int, len(_g))
	for _ef, _af := range _g {
		_dag[_ef] = int(_af)
	}
	return _dag
}
func (_cf NumSlice) Get(i int) (float32, error) {
	if i < 0 || i > len(_cf)-1 {
		return 0, _b.Errorf("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _cf[i], nil
}
func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

type IntsMap map[uint64][]int

func (_fd IntsMap) GetSlice(key uint64) ([]int, bool) {
	_ff, _ea := _fd[key]
	if !_ea {
		return nil, false
	}
	return _ff, true
}
func (_fga *IntSlice) Copy() *IntSlice {
	_ab := IntSlice(make([]int, len(*_fga)))
	copy(_ab, *_fga)
	return &_ab
}
func (_da IntSlice) Size() int { return len(_da) }

type IntSlice []int

func NewNumSlice(i int) *NumSlice { _fc := NumSlice(make([]float32, i)); return &_fc }
func (_fcg *Stack) peek() (interface{}, bool) {
	_gf := _fcg.top()
	if _gf == -1 {
		return nil, false
	}
	return _fcg.Data[_gf], true
}
func (_df *NumSlice) Add(v float32) { *_df = append(*_df, v) }

type NumSlice []float32

func (_dd IntSlice) Get(index int) (int, error) {
	if index > len(_dd)-1 {
		return 0, _b.Errorf("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _dd[index], nil
}
func (_gg *Stack) Peek() (_dfb interface{}, _aa bool) { return _gg.peek() }
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
