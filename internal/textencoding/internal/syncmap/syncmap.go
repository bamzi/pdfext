package syncmap

import _g "sync"

func (_gae *ByteRuneMap) Read(b byte) (rune, bool) {
	_gae._e.RLock()
	defer _gae._e.RUnlock()
	_dc, _a := _gae._ga[b]
	return _dc, _a
}
func (_bbf *RuneUint16Map) RangeDelete(f func(_bdc rune, _ffa uint16) (_gaeb bool, _gba bool)) {
	_bbf._gd.Lock()
	defer _bbf._gd.Unlock()
	for _cg, _dgb := range _bbf._ba {
		_fabe, _ced := f(_cg, _dgb)
		if _fabe {
			delete(_bbf._ba, _cg)
		}
		if _ced {
			break
		}
	}
}
func (_f *RuneByteMap) Read(r rune) (byte, bool) {
	_f._dec.RLock()
	defer _f._dec.RUnlock()
	_af, _gf := _f._ded[r]
	return _af, _gf
}
func (_geb *RuneStringMap) Read(r rune) (string, bool) {
	_geb._ag.RLock()
	defer _geb._ag.RUnlock()
	_aab, _db := _geb._fab[r]
	return _aab, _db
}

type StringRuneMap struct {
	_efge map[string]rune
	_ea   _g.RWMutex
}
type RuneSet struct {
	_ege map[rune]struct{}
	_fa  _g.RWMutex
}

func (_gb *RuneStringMap) Length() int {
	_gb._ag.RLock()
	defer _gb._ag.RUnlock()
	return len(_gb._fab)
}
func (_ge *RuneByteMap) Range(f func(_aa rune, _gec byte) (_eg bool)) {
	_ge._dec.RLock()
	defer _ge._dec.RUnlock()
	for _fb, _fd := range _ge._ded {
		if f(_fb, _fd) {
			break
		}
	}
}
func MakeRuneSet(length int) *RuneSet { return &RuneSet{_ege: make(map[rune]struct{}, length)} }
func (_fda *RuneUint16Map) Length() int {
	_fda._gd.RLock()
	defer _fda._gd.RUnlock()
	return len(_fda._ba)
}
func (_ae *RuneStringMap) Range(f func(_ed rune, _ace string) (_eeg bool)) {
	_ae._ag.RLock()
	defer _ae._ag.RUnlock()
	for _ecc, _fbc := range _ae._fab {
		if f(_ecc, _fbc) {
			break
		}
	}
}

type RuneStringMap struct {
	_fab map[rune]string
	_ag  _g.RWMutex
}
type ByteRuneMap struct {
	_ga map[byte]rune
	_e  _g.RWMutex
}

func (_bg *RuneByteMap) Write(r rune, b byte) {
	_bg._dec.Lock()
	defer _bg._dec.Unlock()
	_bg._ded[r] = b
}
func (_daa *StringsMap) Write(g1, g2 string) {
	_daa._babb.Lock()
	defer _daa._babb.Unlock()
	_daa._dcc[g1] = g2
}
func MakeRuneByteMap(length int) *RuneByteMap {
	_ee := make(map[rune]byte, length)
	return &RuneByteMap{_ded: _ee}
}
func NewByteRuneMap(m map[byte]rune) *ByteRuneMap { return &ByteRuneMap{_ga: m} }
func (_daf *RuneSet) Exists(r rune) bool {
	_daf._fa.RLock()
	defer _daf._fa.RUnlock()
	_, _df := _daf._ege[r]
	return _df
}
func (_ff *RuneStringMap) Write(r rune, s string) {
	_ff._ag.Lock()
	defer _ff._ag.Unlock()
	_ff._fab[r] = s
}
func (_fbg *StringRuneMap) Range(f func(_faf string, _gaca rune) (_bc bool)) {
	_fbg._ea.RLock()
	defer _fbg._ea.RUnlock()
	for _fbga, _dga := range _fbg._efge {
		if f(_fbga, _dga) {
			break
		}
	}
}

type RuneUint16Map struct {
	_ba map[rune]uint16
	_gd _g.RWMutex
}

func (_afc *RuneUint16Map) Delete(r rune) {
	_afc._gd.Lock()
	defer _afc._gd.Unlock()
	delete(_afc._ba, r)
}
func (_gg *ByteRuneMap) Range(f func(_de byte, _dd rune) (_c bool)) {
	_gg._e.RLock()
	defer _gg._e.RUnlock()
	for _eb, _da := range _gg._ga {
		if f(_eb, _da) {
			break
		}
	}
}

type StringsMap struct {
	_dcc  map[string]string
	_babb _g.RWMutex
}

func MakeByteRuneMap(length int) *ByteRuneMap { return &ByteRuneMap{_ga: make(map[byte]rune, length)} }
func (_cf *StringsMap) Range(f func(_bgg, _fg string) (_dca bool)) {
	_cf._babb.RLock()
	defer _cf._babb.RUnlock()
	for _bgd, _bdb := range _cf._dcc {
		if f(_bgd, _bdb) {
			break
		}
	}
}
func (_eccg *StringsMap) Copy() *StringsMap {
	_eccg._babb.RLock()
	defer _eccg._babb.RUnlock()
	_bac := map[string]string{}
	for _gaf, _fgc := range _eccg._dcc {
		_bac[_gaf] = _fgc
	}
	return &StringsMap{_dcc: _bac}
}
func (_b *ByteRuneMap) Length() int {
	_b._e.RLock()
	defer _b._e.RUnlock()
	return len(_b._ga)
}
func NewStringsMap(tuples []StringsTuple) *StringsMap {
	_def := map[string]string{}
	for _, _efgb := range tuples {
		_def[_efgb.Key] = _efgb.Value
	}
	return &StringsMap{_dcc: _def}
}
func (_ce *RuneSet) Range(f func(_bb rune) (_ec bool)) {
	_ce._fa.RLock()
	defer _ce._fa.RUnlock()
	for _aaa := range _ce._ege {
		if f(_aaa) {
			break
		}
	}
}
func NewRuneStringMap(m map[rune]string) *RuneStringMap { return &RuneStringMap{_fab: m} }
func (_gac *RuneSet) Write(r rune) {
	_gac._fa.Lock()
	defer _gac._fa.Unlock()
	_gac._ege[r] = struct{}{}
}

type RuneByteMap struct {
	_ded map[rune]byte
	_dec _g.RWMutex
}

func (_dcb *RuneSet) Length() int { _dcb._fa.RLock(); defer _dcb._fa.RUnlock(); return len(_dcb._ege) }
func (_eeb *RuneByteMap) Length() int {
	_eeb._dec.RLock()
	defer _eeb._dec.RUnlock()
	return len(_eeb._ded)
}
func (_dda *StringRuneMap) Read(g string) (rune, bool) {
	_dda._ea.RLock()
	defer _dda._ea.RUnlock()
	_gda, _fff := _dda._efge[g]
	return _gda, _fff
}
func (_ef *ByteRuneMap) Write(b byte, r rune) {
	_ef._e.Lock()
	defer _ef._e.Unlock()
	_ef._ga[b] = r
}
func (_afd *StringRuneMap) Write(g string, r rune) {
	_afd._ea.Lock()
	defer _afd._ea.Unlock()
	_afd._efge[g] = r
}

type StringsTuple struct{ Key, Value string }

func (_dg *RuneUint16Map) Write(r rune, g uint16) {
	_dg._gd.Lock()
	defer _dg._gd.Unlock()
	_dg._ba[r] = g
}
func NewStringRuneMap(m map[string]rune) *StringRuneMap { return &StringRuneMap{_efge: m} }
func (_fba *StringRuneMap) Length() int {
	_fba._ea.RLock()
	defer _fba._ea.RUnlock()
	return len(_fba._efge)
}
func MakeRuneUint16Map(length int) *RuneUint16Map {
	return &RuneUint16Map{_ba: make(map[rune]uint16, length)}
}
func (_bab *RuneUint16Map) Range(f func(_efg rune, _gaec uint16) (_edg bool)) {
	_bab._gd.RLock()
	defer _bab._gd.RUnlock()
	for _acf, _bd := range _bab._ba {
		if f(_acf, _bd) {
			break
		}
	}
}
func (_ecf *RuneUint16Map) Read(r rune) (uint16, bool) {
	_ecf._gd.RLock()
	defer _ecf._gd.RUnlock()
	_ab, _aeb := _ecf._ba[r]
	return _ab, _aeb
}
func (_fafg *StringsMap) Read(g string) (string, bool) {
	_fafg._babb.RLock()
	defer _fafg._babb.RUnlock()
	_efd, _eebe := _fafg._dcc[g]
	return _efd, _eebe
}
