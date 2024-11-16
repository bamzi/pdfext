// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
//
// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
package ps

import (
	_g "bufio"
	_c "bytes"
	_a "errors"
	_b "fmt"
	_ab "io"
	_f "math"

	_d "github.com/bamzi/pdfext/common"
	_ed "github.com/bamzi/pdfext/core"
)

func (_cgc *PSBoolean) Duplicate() PSObject { _bd := PSBoolean{}; _bd.Val = _cgc.Val; return &_bd }
func (_cc *PSReal) String() string          { return _b.Sprintf("\u0025\u002e\u0035\u0066", _cc.Val) }
func (_aeg *PSOperand) and(_cgg *PSStack) error {
	_fcg, _ecb := _cgg.Pop()
	if _ecb != nil {
		return _ecb
	}
	_aaf, _ecb := _cgg.Pop()
	if _ecb != nil {
		return _ecb
	}
	if _gga, _ega := _fcg.(*PSBoolean); _ega {
		_ebd, _ee := _aaf.(*PSBoolean)
		if !_ee {
			return ErrTypeCheck
		}
		_ecb = _cgg.Push(MakeBool(_gga.Val && _ebd.Val))
		return _ecb
	}
	if _addc, _feb := _fcg.(*PSInteger); _feb {
		_bda, _bada := _aaf.(*PSInteger)
		if !_bada {
			return ErrTypeCheck
		}
		_ecb = _cgg.Push(MakeInteger(_addc.Val & _bda.Val))
		return _ecb
	}
	return ErrTypeCheck
}
func (_bebe *PSParser) skipComments() error {
	if _, _bfe := _bebe.skipSpaces(); _bfe != nil {
		return _bfe
	}
	_ebc := true
	for {
		_acfb, _bcc := _bebe._gdf.Peek(1)
		if _bcc != nil {
			_d.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _bcc.Error())
			return _bcc
		}
		if _ebc && _acfb[0] != '%' {
			return nil
		}
		_ebc = false
		if (_acfb[0] != '\r') && (_acfb[0] != '\n') {
			_bebe._gdf.ReadByte()
		} else {
			break
		}
	}
	return _bebe.skipComments()
}

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor(program *PSProgram) *PSExecutor {
	_bg := &PSExecutor{}
	_bg.Stack = NewPSStack()
	_bg._bc = program
	return _bg
}

// Exec executes the operand `op` in the state specified by `stack`.
func (_cae *PSOperand) Exec(stack *PSStack) error {
	_fgb := ErrUnsupportedOperand
	switch *_cae {
	case "\u0061\u0062\u0073":
		_fgb = _cae.abs(stack)
	case "\u0061\u0064\u0064":
		_fgb = _cae.add(stack)
	case "\u0061\u006e\u0064":
		_fgb = _cae.and(stack)
	case "\u0061\u0074\u0061\u006e":
		_fgb = _cae.atan(stack)
	case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":
		_fgb = _cae.bitshift(stack)
	case "\u0063e\u0069\u006c\u0069\u006e\u0067":
		_fgb = _cae.ceiling(stack)
	case "\u0063\u006f\u0070\u0079":
		_fgb = _cae.copy(stack)
	case "\u0063\u006f\u0073":
		_fgb = _cae.cos(stack)
	case "\u0063\u0076\u0069":
		_fgb = _cae.cvi(stack)
	case "\u0063\u0076\u0072":
		_fgb = _cae.cvr(stack)
	case "\u0064\u0069\u0076":
		_fgb = _cae.div(stack)
	case "\u0064\u0075\u0070":
		_fgb = _cae.dup(stack)
	case "\u0065\u0071":
		_fgb = _cae.eq(stack)
	case "\u0065\u0078\u0063\u0068":
		_fgb = _cae.exch(stack)
	case "\u0065\u0078\u0070":
		_fgb = _cae.exp(stack)
	case "\u0066\u006c\u006fo\u0072":
		_fgb = _cae.floor(stack)
	case "\u0067\u0065":
		_fgb = _cae.ge(stack)
	case "\u0067\u0074":
		_fgb = _cae.gt(stack)
	case "\u0069\u0064\u0069\u0076":
		_fgb = _cae.idiv(stack)
	case "\u0069\u0066":
		_fgb = _cae.ifCondition(stack)
	case "\u0069\u0066\u0065\u006c\u0073\u0065":
		_fgb = _cae.ifelse(stack)
	case "\u0069\u006e\u0064e\u0078":
		_fgb = _cae.index(stack)
	case "\u006c\u0065":
		_fgb = _cae.le(stack)
	case "\u006c\u006f\u0067":
		_fgb = _cae.log(stack)
	case "\u006c\u006e":
		_fgb = _cae.ln(stack)
	case "\u006c\u0074":
		_fgb = _cae.lt(stack)
	case "\u006d\u006f\u0064":
		_fgb = _cae.mod(stack)
	case "\u006d\u0075\u006c":
		_fgb = _cae.mul(stack)
	case "\u006e\u0065":
		_fgb = _cae.ne(stack)
	case "\u006e\u0065\u0067":
		_fgb = _cae.neg(stack)
	case "\u006e\u006f\u0074":
		_fgb = _cae.not(stack)
	case "\u006f\u0072":
		_fgb = _cae.or(stack)
	case "\u0070\u006f\u0070":
		_fgb = _cae.pop(stack)
	case "\u0072\u006f\u0075n\u0064":
		_fgb = _cae.round(stack)
	case "\u0072\u006f\u006c\u006c":
		_fgb = _cae.roll(stack)
	case "\u0073\u0069\u006e":
		_fgb = _cae.sin(stack)
	case "\u0073\u0071\u0072\u0074":
		_fgb = _cae.sqrt(stack)
	case "\u0073\u0075\u0062":
		_fgb = _cae.sub(stack)
	case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":
		_fgb = _cae.truncate(stack)
	case "\u0078\u006f\u0072":
		_fgb = _cae.xor(stack)
	}
	return _fgb
}

// Push pushes an object on top of the stack.
func (_fgbg *PSStack) Push(obj PSObject) error {
	if len(*_fgbg) > 100 {
		return ErrStackOverflow
	}
	*_fgbg = append(*_fgbg, obj)
	return nil
}
func (_fgbb *PSOperand) neg(_bfde *PSStack) error {
	_daga, _bca := _bfde.Pop()
	if _bca != nil {
		return _bca
	}
	if _bfae, _ada := _daga.(*PSReal); _ada {
		_bca = _bfde.Push(MakeReal(-_bfae.Val))
		return _bca
	} else if _bdced, _bdbg := _daga.(*PSInteger); _bdbg {
		_bca = _bfde.Push(MakeInteger(-_bdced.Val))
		return _bca
	} else {
		return ErrTypeCheck
	}
}
func (_cad *PSOperand) floor(_eda *PSStack) error {
	_feg, _efg := _eda.Pop()
	if _efg != nil {
		return _efg
	}
	if _adfg, _dfbd := _feg.(*PSReal); _dfbd {
		_efg = _eda.Push(MakeReal(_f.Floor(_adfg.Val)))
	} else if _afc, _gbb := _feg.(*PSInteger); _gbb {
		_efg = _eda.Push(MakeInteger(_afc.Val))
	} else {
		return ErrTypeCheck
	}
	return _efg
}
func (_af *PSBoolean) String() string { return _b.Sprintf("\u0025\u0076", _af.Val) }

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand(val string) *PSOperand { _dfffd := PSOperand(val); return &_dfffd }

// Parse parses the postscript and store as a program that can be executed.
func (_ffc *PSParser) Parse() (*PSProgram, error) {
	_ffc.skipSpaces()
	_fdcf, _bdgb := _ffc._gdf.Peek(2)
	if _bdgb != nil {
		return nil, _bdgb
	}
	if _fdcf[0] != '{' {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b")
	}
	_eabg, _bdgb := _ffc.parseFunction()
	if _bdgb != nil && _bdgb != _ab.EOF {
		return nil, _bdgb
	}
	return _eabg, _bdgb
}
func (_feebf *PSOperand) sin(_eeb *PSStack) error {
	_ebdea, _cbec := _eeb.PopNumberAsFloat64()
	if _cbec != nil {
		return _cbec
	}
	_gbdg := _f.Sin(_ebdea * _f.Pi / 180.0)
	_cbec = _eeb.Push(MakeReal(_gbdg))
	return _cbec
}
func (_afa *PSOperand) gt(_cfa *PSStack) error {
	_gae, _cdd := _cfa.PopNumberAsFloat64()
	if _cdd != nil {
		return _cdd
	}
	_dce, _cdd := _cfa.PopNumberAsFloat64()
	if _cdd != nil {
		return _cdd
	}
	if _f.Abs(_dce-_gae) < _fg {
		_babf := _cfa.Push(MakeBool(false))
		return _babf
	} else if _dce > _gae {
		_bdbe := _cfa.Push(MakeBool(true))
		return _bdbe
	} else {
		_eef := _cfa.Push(MakeBool(false))
		return _eef
	}
}

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser(content []byte) *PSParser {
	_ebf := PSParser{}
	_ggg := _c.NewBuffer(content)
	_ebf._gdf = _g.NewReader(_ggg)
	return &_ebf
}

// PSInteger represents an integer.
type PSInteger struct{ Val int }

func (_bfa *PSOperand) ifCondition(_fcb *PSStack) error {
	_dcg, _edge := _fcb.Pop()
	if _edge != nil {
		return _edge
	}
	_fbbb, _edge := _fcb.Pop()
	if _edge != nil {
		return _edge
	}
	_dgdg, _bea := _dcg.(*PSProgram)
	if !_bea {
		return ErrTypeCheck
	}
	_ecbg, _bea := _fbbb.(*PSBoolean)
	if !_bea {
		return ErrTypeCheck
	}
	if _ecbg.Val {
		_cadb := _dgdg.Exec(_fcb)
		return _cadb
	}
	return nil
}
func (_cbca *PSParser) parseFunction() (*PSProgram, error) {
	_ced, _ := _cbca._gdf.ReadByte()
	if _ced != '{' {
		return nil, _a.New("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_ggae := NewPSProgram()
	for {
		_cbca.skipSpaces()
		_cbca.skipComments()
		_gcd, _dgdc := _cbca._gdf.Peek(2)
		if _dgdc != nil {
			if _dgdc == _ab.EOF {
				break
			}
			return nil, _dgdc
		}
		_d.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gcd))
		if _gcd[0] == '}' {
			_d.Log.Trace("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
			_cbca._gdf.ReadByte()
			break
		} else if _gcd[0] == '{' {
			_d.Log.Trace("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!")
			_cgagb, _ffac := _cbca.parseFunction()
			if _ffac != nil {
				return nil, _ffac
			}
			_ggae.Append(_cgagb)
		} else if _ed.IsDecimalDigit(_gcd[0]) || (_gcd[0] == '-' && _ed.IsDecimalDigit(_gcd[1])) {
			_d.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_dbcd, _fcdb := _cbca.parseNumber()
			if _fcdb != nil {
				return nil, _fcdb
			}
			_ggae.Append(_dbcd)
		} else {
			_d.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_gcd, _ = _cbca._gdf.Peek(5)
			_ceeb := string(_gcd)
			_d.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _ceeb)
			if (len(_ceeb) > 4) && (_ceeb[:5] == "\u0066\u0061\u006cs\u0065") {
				_gccg, _cedf := _cbca.parseBool()
				if _cedf != nil {
					return nil, _cedf
				}
				_ggae.Append(_gccg)
			} else if (len(_ceeb) > 3) && (_ceeb[:4] == "\u0074\u0072\u0075\u0065") {
				_cacg, _cccd := _cbca.parseBool()
				if _cccd != nil {
					return nil, _cccd
				}
				_ggae.Append(_cacg)
			} else {
				_eeg, _cfe := _cbca.parseOperand()
				if _cfe != nil {
					return nil, _cfe
				}
				_ggae.Append(_eeg)
			}
		}
	}
	return _ggae, nil
}
func (_gacf *PSOperand) le(_bdd *PSStack) error {
	_deb, _egbe := _bdd.PopNumberAsFloat64()
	if _egbe != nil {
		return _egbe
	}
	_fbd, _egbe := _bdd.PopNumberAsFloat64()
	if _egbe != nil {
		return _egbe
	}
	if _f.Abs(_fbd-_deb) < _fg {
		_bae := _bdd.Push(MakeBool(true))
		return _bae
	} else if _fbd < _deb {
		_aace := _bdd.Push(MakeBool(true))
		return _aace
	} else {
		_beb := _bdd.Push(MakeBool(false))
		return _beb
	}
}
func (_gff *PSOperand) truncate(_fbac *PSStack) error {
	_aegd, _cccf := _fbac.Pop()
	if _cccf != nil {
		return _cccf
	}
	if _acg, _gafe := _aegd.(*PSReal); _gafe {
		_dfcda := int(_acg.Val)
		_cccf = _fbac.Push(MakeReal(float64(_dfcda)))
	} else if _gbcc, _aeff := _aegd.(*PSInteger); _aeff {
		_cccf = _fbac.Push(MakeInteger(_gbcc.Val))
	} else {
		return ErrTypeCheck
	}
	return _cccf
}

// PSOperand represents a Postscript operand (text string).
type PSOperand string

func (_ddfb *PSOperand) round(_ece *PSStack) error {
	_agba, _efgc := _ece.Pop()
	if _efgc != nil {
		return _efgc
	}
	if _eee, _gce := _agba.(*PSReal); _gce {
		_efgc = _ece.Push(MakeReal(_f.Floor(_eee.Val + 0.5)))
	} else if _gcc, _bgga := _agba.(*PSInteger); _bgga {
		_efgc = _ece.Push(MakeInteger(_gcc.Val))
	} else {
		return ErrTypeCheck
	}
	return _efgc
}

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_bdcg *PSStack) PopInteger() (int, error) {
	_deef, _dfdg := _bdcg.Pop()
	if _dfdg != nil {
		return 0, _dfdg
	}
	if _fab, _fdd := _deef.(*PSInteger); _fdd {
		return _fab.Val, nil
	}
	return 0, ErrTypeCheck
}
func (_bgc *PSOperand) add(_bef *PSStack) error {
	_gdb, _bce := _bef.Pop()
	if _bce != nil {
		return _bce
	}
	_aa, _bce := _bef.Pop()
	if _bce != nil {
		return _bce
	}
	_ae, _ef := _gdb.(*PSReal)
	_fef, _adf := _gdb.(*PSInteger)
	if !_ef && !_adf {
		return ErrTypeCheck
	}
	_bdba, _bad := _aa.(*PSReal)
	_bcb, _dg := _aa.(*PSInteger)
	if !_bad && !_dg {
		return ErrTypeCheck
	}
	if _adf && _dg {
		_fc := _fef.Val + _bcb.Val
		_gac := _bef.Push(MakeInteger(_fc))
		return _gac
	}
	var _add float64
	if _ef {
		_add = _ae.Val
	} else {
		_add = float64(_fef.Val)
	}
	if _bad {
		_add += _bdba.Val
	} else {
		_add += float64(_bcb.Val)
	}
	_bce = _bef.Push(MakeReal(_add))
	return _bce
}
func (_dffb *PSOperand) log(_dgab *PSStack) error {
	_afg, _dda := _dgab.PopNumberAsFloat64()
	if _dda != nil {
		return _dda
	}
	_ebdc := _f.Log10(_afg)
	_dda = _dgab.Push(MakeReal(_ebdc))
	return _dda
}

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_gec *PSStack) PopNumberAsFloat64() (float64, error) {
	_bbfe, _aggg := _gec.Pop()
	if _aggg != nil {
		return 0, _aggg
	}
	if _bbad, _eccb := _bbfe.(*PSReal); _eccb {
		return _bbad.Val, nil
	} else if _fcgc, _bcbd := _bbfe.(*PSInteger); _bcbd {
		return float64(_fcgc.Val), nil
	} else {
		return 0, ErrTypeCheck
	}
}
func (_bdag *PSOperand) idiv(_gfec *PSStack) error {
	_fffb, _befg := _gfec.Pop()
	if _befg != nil {
		return _befg
	}
	_fgfde, _befg := _gfec.Pop()
	if _befg != nil {
		return _befg
	}
	_adfa, _fega := _fffb.(*PSInteger)
	if !_fega {
		return ErrTypeCheck
	}
	if _adfa.Val == 0 {
		return ErrUndefinedResult
	}
	_dcc, _fega := _fgfde.(*PSInteger)
	if !_fega {
		return ErrTypeCheck
	}
	_eec := _dcc.Val / _adfa.Val
	_befg = _gfec.Push(MakeInteger(_eec))
	return _befg
}

// String returns a string representation of the stack.
func (_ade *PSStack) String() string {
	_gdbdg := "\u005b\u0020"
	for _, _cbb := range *_ade {
		_gdbdg += _cbb.String()
		_gdbdg += "\u0020"
	}
	_gdbdg += "\u005d"
	return _gdbdg
}
func (_bgd *PSOperand) eq(_cbg *PSStack) error {
	_gfd, _ddba := _cbg.Pop()
	if _ddba != nil {
		return _ddba
	}
	_fba, _ddba := _cbg.Pop()
	if _ddba != nil {
		return _ddba
	}
	_ac, _dfc := _gfd.(*PSBoolean)
	_dac, _dagd := _fba.(*PSBoolean)
	if _dfc || _dagd {
		var _ggfd error
		if _dfc && _dagd {
			_ggfd = _cbg.Push(MakeBool(_ac.Val == _dac.Val))
		} else {
			_ggfd = _cbg.Push(MakeBool(false))
		}
		return _ggfd
	}
	var _gdd float64
	var _ecf float64
	if _dgdb, _fae := _gfd.(*PSInteger); _fae {
		_gdd = float64(_dgdb.Val)
	} else if _aaa, _feeba := _gfd.(*PSReal); _feeba {
		_gdd = _aaa.Val
	} else {
		return ErrTypeCheck
	}
	if _fbe, _ecg := _fba.(*PSInteger); _ecg {
		_ecf = float64(_fbe.Val)
	} else if _bee, _eag := _fba.(*PSReal); _eag {
		_ecf = _bee.Val
	} else {
		return ErrTypeCheck
	}
	if _f.Abs(_ecf-_gdd) < _fg {
		_ddba = _cbg.Push(MakeBool(true))
	} else {
		_ddba = _cbg.Push(MakeBool(false))
	}
	return _ddba
}
func (_gd *PSInteger) String() string { return _b.Sprintf("\u0025\u0064", _gd.Val) }
func (_db *PSOperand) cvi(_ffad *PSStack) error {
	_cga, _fce := _ffad.Pop()
	if _fce != nil {
		return _fce
	}
	if _abe, _dba := _cga.(*PSReal); _dba {
		_fag := int(_abe.Val)
		_fce = _ffad.Push(MakeInteger(_fag))
	} else if _geg, _dc := _cga.(*PSInteger); _dc {
		_bdc := _geg.Val
		_fce = _ffad.Push(MakeInteger(_bdc))
	} else {
		return ErrTypeCheck
	}
	return _fce
}
func (_cb *PSReal) DebugString() string {
	return _b.Sprintf("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f", _cb.Val)
}

var ErrStackOverflow = _a.New("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077")

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array(objects []PSObject) ([]float64, error) {
	var _gg []float64
	for _, _gb := range objects {
		if _bf, _da := _gb.(*PSInteger); _da {
			_gg = append(_gg, float64(_bf.Val))
		} else if _bcf, _ff := _gb.(*PSReal); _ff {
			_gg = append(_gg, _bcf.Val)
		} else {
			return nil, ErrTypeCheck
		}
	}
	return _gg, nil
}
func (_bdf *PSOperand) cvr(_bcd *PSStack) error {
	_aac, _fbf := _bcd.Pop()
	if _fbf != nil {
		return _fbf
	}
	if _aag, _df := _aac.(*PSReal); _df {
		_fbf = _bcd.Push(MakeReal(_aag.Val))
	} else if _cef, _ea := _aac.(*PSInteger); _ea {
		_fbf = _bcd.Push(MakeReal(float64(_cef.Val)))
	} else {
		return ErrTypeCheck
	}
	return _fbf
}
func (_ddf *PSBoolean) DebugString() string {
	return _b.Sprintf("\u0062o\u006f\u006c\u003a\u0025\u0076", _ddf.Val)
}
func (_gda *PSOperand) copy(_cac *PSStack) error {
	_caf, _cf := _cac.PopInteger()
	if _cf != nil {
		return _cf
	}
	if _caf < 0 {
		return ErrRangeCheck
	}
	if _caf > len(*_cac) {
		return ErrRangeCheck
	}
	*_cac = append(*_cac, (*_cac)[len(*_cac)-_caf:]...)
	return nil
}

var ErrStackUnderflow = _a.New("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077")

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_fe *PSExecutor) Execute(objects []PSObject) ([]PSObject, error) {
	for _, _gf := range objects {
		_dd := _fe.Stack.Push(_gf)
		if _dd != nil {
			return nil, _dd
		}
	}
	_ga := _fe._bc.Exec(_fe.Stack)
	if _ga != nil {
		_d.Log.Debug("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ga)
		return nil, _ga
	}
	_be := []PSObject(*_fe.Stack)
	_fe.Stack.Empty()
	return _be, nil
}

// Empty empties the stack.
func (_cbce *PSStack) Empty() { *_cbce = []PSObject{} }
func (_egb *PSOperand) ifelse(_adge *PSStack) error {
	_cfc, _cdc := _adge.Pop()
	if _cdc != nil {
		return _cdc
	}
	_cdfe, _cdc := _adge.Pop()
	if _cdc != nil {
		return _cdc
	}
	_afe, _cdc := _adge.Pop()
	if _cdc != nil {
		return _cdc
	}
	_bde, _cafg := _cfc.(*PSProgram)
	if !_cafg {
		return ErrTypeCheck
	}
	_dbc, _cafg := _cdfe.(*PSProgram)
	if !_cafg {
		return ErrTypeCheck
	}
	_fbec, _cafg := _afe.(*PSBoolean)
	if !_cafg {
		return ErrTypeCheck
	}
	if _fbec.Val {
		_dfff := _dbc.Exec(_adge)
		return _dfff
	}
	_cdc = _bde.Exec(_adge)
	return _cdc
}

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger(val int) *PSInteger { _cegc := PSInteger{}; _cegc.Val = val; return &_cegc }

var ErrUndefinedResult = _a.New("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072")

// Exec executes the program, typically leaving output values on the stack.
func (_fee *PSProgram) Exec(stack *PSStack) error {
	for _, _bdb := range *_fee {
		var _ede error
		switch _gge := _bdb.(type) {
		case *PSInteger:
			_fbb := _gge
			_ede = stack.Push(_fbb)
		case *PSReal:
			_bfd := _gge
			_ede = stack.Push(_bfd)
		case *PSBoolean:
			_gcf := _gge
			_ede = stack.Push(_gcf)
		case *PSProgram:
			_gfg := _gge
			_ede = stack.Push(_gfg)
		case *PSOperand:
			_ccd := _gge
			_ede = _ccd.Exec(stack)
		default:
			return ErrTypeCheck
		}
		if _ede != nil {
			return _ede
		}
	}
	return nil
}

// Pop pops an object from the top of the stack.
func (_geb *PSStack) Pop() (PSObject, error) {
	if len(*_geb) < 1 {
		return nil, ErrStackUnderflow
	}
	_dee := (*_geb)[len(*_geb)-1]
	*_geb = (*_geb)[0 : len(*_geb)-1]
	return _dee, nil
}
func (_dag *PSOperand) bitshift(_febb *PSStack) error {
	_gfe, _ggaa := _febb.PopInteger()
	if _ggaa != nil {
		return _ggaa
	}
	_ddfg, _ggaa := _febb.PopInteger()
	if _ggaa != nil {
		return _ggaa
	}
	var _gaca int
	if _gfe >= 0 {
		_gaca = _ddfg << uint(_gfe)
	} else {
		_gaca = _ddfg >> uint(-_gfe)
	}
	_ggaa = _febb.Push(MakeInteger(_gaca))
	return _ggaa
}
func (_agb *PSOperand) div(_bdce *PSStack) error {
	_edc, _cge := _bdce.Pop()
	if _cge != nil {
		return _cge
	}
	_fgbf, _cge := _bdce.Pop()
	if _cge != nil {
		return _cge
	}
	_ecc, _bbg := _edc.(*PSReal)
	_dgd, _ggag := _edc.(*PSInteger)
	if !_bbg && !_ggag {
		return ErrTypeCheck
	}
	if _bbg && _ecc.Val == 0 {
		return ErrUndefinedResult
	}
	if _ggag && _dgd.Val == 0 {
		return ErrUndefinedResult
	}
	_aff, _adfe := _fgbf.(*PSReal)
	_cbc, _gba := _fgbf.(*PSInteger)
	if !_adfe && !_gba {
		return ErrTypeCheck
	}
	var _bbd float64
	if _adfe {
		_bbd = _aff.Val
	} else {
		_bbd = float64(_cbc.Val)
	}
	if _bbg {
		_bbd /= _ecc.Val
	} else {
		_bbd /= float64(_dgd.Val)
	}
	_cge = _bdce.Push(MakeReal(_bbd))
	return _cge
}
func (_dga *PSOperand) ge(_daef *PSStack) error {
	_gag, _cdf := _daef.PopNumberAsFloat64()
	if _cdf != nil {
		return _cdf
	}
	_dec, _cdf := _daef.PopNumberAsFloat64()
	if _cdf != nil {
		return _cdf
	}
	if _f.Abs(_dec-_gag) < _fg {
		_cacc := _daef.Push(MakeBool(true))
		return _cacc
	} else if _dec > _gag {
		_abca := _daef.Push(MakeBool(true))
		return _abca
	} else {
		_fbg := _daef.Push(MakeBool(false))
		return _fbg
	}
}
func (_eg *PSProgram) DebugString() string {
	_edg := "\u007b\u0020"
	for _, _ffa := range *_eg {
		_edg += _ffa.DebugString()
		_edg += "\u0020"
	}
	_edg += "\u007d"
	return _edg
}

// NewPSStack returns an initialized PSStack.
func NewPSStack() *PSStack { return &PSStack{} }

const _fg = 0.000001

func (_cd *PSOperand) DebugString() string {
	return _b.Sprintf("\u006fp\u003a\u0027\u0025\u0073\u0027", *_cd)
}
func (_gcab *PSOperand) String() string { return string(*_gcab) }

// PSParser is a basic Postscript parser.
type PSParser struct{ _gdf *_g.Reader }

func (_gaf *PSOperand) mul(_febc *PSStack) error {
	_gfdf, _ebdf := _febc.Pop()
	if _ebdf != nil {
		return _ebdf
	}
	_cbf, _ebdf := _febc.Pop()
	if _ebdf != nil {
		return _ebdf
	}
	_dea, _bdca := _gfdf.(*PSReal)
	_gbc, _ggab := _gfdf.(*PSInteger)
	if !_bdca && !_ggab {
		return ErrTypeCheck
	}
	_dgf, _ggaf := _cbf.(*PSReal)
	_bec, _ebad := _cbf.(*PSInteger)
	if !_ggaf && !_ebad {
		return ErrTypeCheck
	}
	if _ggab && _ebad {
		_fdbf := _gbc.Val * _bec.Val
		_ggfg := _febc.Push(MakeInteger(_fdbf))
		return _ggfg
	}
	var _agd float64
	if _bdca {
		_agd = _dea.Val
	} else {
		_agd = float64(_gbc.Val)
	}
	if _ggaf {
		_agd *= _dgf.Val
	} else {
		_agd *= float64(_bec.Val)
	}
	_ebdf = _febc.Push(MakeReal(_agd))
	return _ebdf
}

var ErrTypeCheck = _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")

func (_ca *PSProgram) Duplicate() PSObject {
	_ce := &PSProgram{}
	for _, _gca := range *_ca {
		_ce.Append(_gca.Duplicate())
	}
	return _ce
}
func (_adaf *PSOperand) pop(_eefd *PSStack) error {
	_, _ebde := _eefd.Pop()
	if _ebde != nil {
		return _ebde
	}
	return nil
}
func (_bed *PSReal) Duplicate() PSObject { _ggf := PSReal{}; _ggf.Val = _bed.Val; return &_ggf }
func (_aeeb *PSOperand) not(_edeg *PSStack) error {
	_cfg, _fbfd := _edeg.Pop()
	if _fbfd != nil {
		return _fbfd
	}
	if _dfcc, _eccg := _cfg.(*PSBoolean); _eccg {
		_fbfd = _edeg.Push(MakeBool(!_dfcc.Val))
		return _fbfd
	} else if _egf, _bfdaa := _cfg.(*PSInteger); _bfdaa {
		_fbfd = _edeg.Push(MakeInteger(^_egf.Val))
		return _fbfd
	} else {
		return ErrTypeCheck
	}
}

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject

func (_gbe *PSProgram) String() string {
	_ba := "\u007b\u0020"
	for _, _ag := range *_gbe {
		_ba += _ag.String()
		_ba += "\u0020"
	}
	_ba += "\u007d"
	return _ba
}
func (_bdg *PSOperand) abs(_ec *PSStack) error {
	_adg, _bb := _ec.Pop()
	if _bb != nil {
		return _bb
	}
	if _abc, _ffab := _adg.(*PSReal); _ffab {
		_bfda := _abc.Val
		if _bfda < 0 {
			_bb = _ec.Push(MakeReal(-_bfda))
		} else {
			_bb = _ec.Push(MakeReal(_bfda))
		}
	} else if _ggeb, _fga := _adg.(*PSInteger); _fga {
		_ecd := _ggeb.Val
		if _ecd < 0 {
			_bb = _ec.Push(MakeInteger(-_ecd))
		} else {
			_bb = _ec.Push(MakeInteger(_ecd))
		}
	} else {
		return ErrTypeCheck
	}
	return _bb
}

// PSReal represents a real number.
type PSReal struct{ Val float64 }

func (_bbf *PSOperand) atan(_daeg *PSStack) error {
	_cee, _fgg := _daeg.PopNumberAsFloat64()
	if _fgg != nil {
		return _fgg
	}
	_fgf, _fgg := _daeg.PopNumberAsFloat64()
	if _fgg != nil {
		return _fgg
	}
	if _cee == 0 {
		var _eba error
		if _fgf < 0 {
			_eba = _daeg.Push(MakeReal(270))
		} else {
			_eba = _daeg.Push(MakeReal(90))
		}
		return _eba
	}
	_fdb := _fgf / _cee
	_bab := _f.Atan(_fdb) * 180 / _f.Pi
	_fgg = _daeg.Push(MakeReal(_bab))
	return _fgg
}
func (_eac *PSOperand) ne(_gee *PSStack) error {
	_fbc := _eac.eq(_gee)
	if _fbc != nil {
		return _fbc
	}
	_fbc = _eac.not(_gee)
	return _fbc
}
func (_feeb *PSOperand) ceiling(_ddb *PSStack) error {
	_fa, _fec := _ddb.Pop()
	if _fec != nil {
		return _fec
	}
	if _fdc, _fad := _fa.(*PSReal); _fad {
		_fec = _ddb.Push(MakeReal(_f.Ceil(_fdc.Val)))
	} else if _ddg, _ceg := _fa.(*PSInteger); _ceg {
		_fec = _ddb.Push(MakeInteger(_ddg.Val))
	} else {
		_fec = ErrTypeCheck
	}
	return _fec
}

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject

// PSBoolean represents a boolean value.
type PSBoolean struct{ Val bool }

func (_aea *PSOperand) index(_bbda *PSStack) error {
	_cdcb, _gfa := _bbda.Pop()
	if _gfa != nil {
		return _gfa
	}
	_ecbgc, _gbd := _cdcb.(*PSInteger)
	if !_gbd {
		return ErrTypeCheck
	}
	if _ecbgc.Val < 0 {
		return ErrRangeCheck
	}
	if _ecbgc.Val > len(*_bbda)-1 {
		return ErrStackUnderflow
	}
	_cafb := (*_bbda)[len(*_bbda)-1-_ecbgc.Val]
	_gfa = _bbda.Push(_cafb.Duplicate())
	return _gfa
}
func (_gdbd *PSOperand) mod(_cdff *PSStack) error {
	_dfd, _baf := _cdff.Pop()
	if _baf != nil {
		return _baf
	}
	_dbd, _baf := _cdff.Pop()
	if _baf != nil {
		return _baf
	}
	_debd, _ebe := _dfd.(*PSInteger)
	if !_ebe {
		return ErrTypeCheck
	}
	if _debd.Val == 0 {
		return ErrUndefinedResult
	}
	_fcgf, _ebe := _dbd.(*PSInteger)
	if !_ebe {
		return ErrTypeCheck
	}
	_fade := _fcgf.Val % _debd.Val
	_baf = _cdff.Push(MakeInteger(_fade))
	return _baf
}

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool(val bool) *PSBoolean { _gggb := PSBoolean{}; _gggb.Val = val; return &_gggb }
func (_caa *PSOperand) exp(_aed *PSStack) error {
	_dfb, _cgag := _aed.PopNumberAsFloat64()
	if _cgag != nil {
		return _cgag
	}
	_cfd, _cgag := _aed.PopNumberAsFloat64()
	if _cgag != nil {
		return _cgag
	}
	if _f.Abs(_dfb) < 1 && _cfd < 0 {
		return ErrUndefinedResult
	}
	_gaaf := _f.Pow(_cfd, _dfb)
	_cgag = _aed.Push(MakeReal(_gaaf))
	return _cgag
}
func (_dacc *PSOperand) xor(_fbbc *PSStack) error {
	_eacg, _eea := _fbbc.Pop()
	if _eea != nil {
		return _eea
	}
	_efe, _eea := _fbbc.Pop()
	if _eea != nil {
		return _eea
	}
	if _fbcd, _ggac := _eacg.(*PSBoolean); _ggac {
		_baa, _fea := _efe.(*PSBoolean)
		if !_fea {
			return ErrTypeCheck
		}
		_eea = _fbbc.Push(MakeBool(_fbcd.Val != _baa.Val))
		return _eea
	}
	if _dcce, _cbd := _eacg.(*PSInteger); _cbd {
		_gaea, _agbg := _efe.(*PSInteger)
		if !_agbg {
			return ErrTypeCheck
		}
		_eea = _fbbc.Push(MakeInteger(_dcce.Val ^ _gaea.Val))
		return _eea
	}
	return ErrTypeCheck
}

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram() *PSProgram { return &PSProgram{} }

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct {
	Stack *PSStack
	_bc   *PSProgram
}

func (_bbae *PSParser) skipSpaces() (int, error) {
	_cebf := 0
	for {
		_efb, _cace := _bbae._gdf.Peek(1)
		if _cace != nil {
			return 0, _cace
		}
		if _ed.IsWhiteSpace(_efb[0]) {
			_bbae._gdf.ReadByte()
			_cebf++
		} else {
			break
		}
	}
	return _cebf, nil
}
func (_fff *PSOperand) dup(_bede *PSStack) error {
	_dff, _ffaa := _bede.Pop()
	if _ffaa != nil {
		return _ffaa
	}
	_ffaa = _bede.Push(_dff)
	if _ffaa != nil {
		return _ffaa
	}
	_ffaa = _bede.Push(_dff.Duplicate())
	return _ffaa
}
func (_ad *PSInteger) DebugString() string {
	return _b.Sprintf("\u0069\u006e\u0074\u003a\u0025\u0064", _ad.Val)
}
func (_fcdbf *PSParser) parseBool() (*PSBoolean, error) {
	_bge, _gffe := _fcdbf._gdf.Peek(4)
	if _gffe != nil {
		return MakeBool(false), _gffe
	}
	if (len(_bge) >= 4) && (string(_bge[:4]) == "\u0074\u0072\u0075\u0065") {
		_fcdbf._gdf.Discard(4)
		return MakeBool(true), nil
	}
	_bge, _gffe = _fcdbf._gdf.Peek(5)
	if _gffe != nil {
		return MakeBool(false), _gffe
	}
	if (len(_bge) >= 5) && (string(_bge[:5]) == "\u0066\u0061\u006cs\u0065") {
		_fcdbf._gdf.Discard(5)
		return MakeBool(false), nil
	}
	return MakeBool(false), _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// PSObject represents a postscript object.
type PSObject interface {

	// Duplicate makes a fresh copy of the PSObject.
	Duplicate() PSObject

	// DebugString returns a descriptive representation of the PSObject with more information than String()
	// for debugging purposes.
	DebugString() string

	// String returns a string representation of the PSObject.
	String() string
}

func (_deba *PSOperand) or(_ceb *PSStack) error {
	_fdcc, _dcbd := _ceb.Pop()
	if _dcbd != nil {
		return _dcbd
	}
	_bcfb, _dcbd := _ceb.Pop()
	if _dcbd != nil {
		return _dcbd
	}
	if _dfcd, _bcae := _fdcc.(*PSBoolean); _bcae {
		_acf, _fbfa := _bcfb.(*PSBoolean)
		if !_fbfa {
			return ErrTypeCheck
		}
		_dcbd = _ceb.Push(MakeBool(_dfcd.Val || _acf.Val))
		return _dcbd
	}
	if _fggd, _dbad := _fdcc.(*PSInteger); _dbad {
		_cfb, _cde := _bcfb.(*PSInteger)
		if !_cde {
			return ErrTypeCheck
		}
		_dcbd = _ceb.Push(MakeInteger(_fggd.Val | _cfb.Val))
		return _dcbd
	}
	return ErrTypeCheck
}

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_ddge *PSStack) DebugString() string {
	_fgc := "\u005b\u0020"
	for _, _aead := range *_ddge {
		_fgc += _aead.DebugString()
		_fgc += "\u0020"
	}
	_fgc += "\u005d"
	return _fgc
}
func (_aafc *PSOperand) roll(_bba *PSStack) error {
	_bac, _fbbg := _bba.Pop()
	if _fbbg != nil {
		return _fbbg
	}
	_ddbac, _fbbg := _bba.Pop()
	if _fbbg != nil {
		return _fbbg
	}
	_eeec, _eab := _bac.(*PSInteger)
	if !_eab {
		return ErrTypeCheck
	}
	_bead, _eab := _ddbac.(*PSInteger)
	if !_eab {
		return ErrTypeCheck
	}
	if _bead.Val < 0 {
		return ErrRangeCheck
	}
	if _bead.Val == 0 || _bead.Val == 1 {
		return nil
	}
	if _bead.Val > len(*_bba) {
		return ErrStackUnderflow
	}
	for _fcd := 0; _fcd < _gddg(_eeec.Val); _fcd++ {
		var _eeeb []PSObject
		_eeeb = (*_bba)[len(*_bba)-(_bead.Val) : len(*_bba)]
		if _eeec.Val > 0 {
			_bcg := _eeeb[len(_eeeb)-1]
			_eeeb = append([]PSObject{_bcg}, _eeeb[0:len(_eeeb)-1]...)
		} else {
			_eecc := _eeeb[len(_eeeb)-_bead.Val]
			_eeeb = append(_eeeb[1:], _eecc)
		}
		_fecg := append((*_bba)[0:len(*_bba)-_bead.Val], _eeeb...)
		_bba = &_fecg
	}
	return nil
}
func _gddg(_agge int) int {
	if _agge < 0 {
		return -_agge
	}
	return _agge
}
func (_dad *PSOperand) Duplicate() PSObject { _agg := *_dad; return &_agg }
func (_eae *PSOperand) sub(_afce *PSStack) error {
	_bedc, _afb := _afce.Pop()
	if _afb != nil {
		return _afb
	}
	_egac, _afb := _afce.Pop()
	if _afb != nil {
		return _afb
	}
	_aef, _gccb := _bedc.(*PSReal)
	_fdf, _ddd := _bedc.(*PSInteger)
	if !_gccb && !_ddd {
		return ErrTypeCheck
	}
	_bcbe, _dffbc := _egac.(*PSReal)
	_badf, _aedc := _egac.(*PSInteger)
	if !_dffbc && !_aedc {
		return ErrTypeCheck
	}
	if _ddd && _aedc {
		_bcfa := _badf.Val - _fdf.Val
		_cdb := _afce.Push(MakeInteger(_bcfa))
		return _cdb
	}
	var _gacc float64 = 0
	if _dffbc {
		_gacc = _bcbe.Val
	} else {
		_gacc = float64(_badf.Val)
	}
	if _gccb {
		_gacc -= _aef.Val
	} else {
		_gacc -= float64(_fdf.Val)
	}
	_afb = _afce.Push(MakeReal(_gacc))
	return _afb
}

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal(val float64) *PSReal { _gggd := PSReal{}; _gggd.Val = val; return &_gggd }
func (_egff *PSParser) parseNumber() (PSObject, error) {
	_adb, _fge := _ed.ParseNumber(_egff._gdf)
	if _fge != nil {
		return nil, _fge
	}
	switch _cege := _adb.(type) {
	case *_ed.PdfObjectFloat:
		return MakeReal(float64(*_cege)), nil
	case *_ed.PdfObjectInteger:
		return MakeInteger(int(*_cege)), nil
	}
	return nil, _b.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _adb)
}
func (_cggc *PSOperand) ln(_dcb *PSStack) error {
	_gbab, _dada := _dcb.PopNumberAsFloat64()
	if _dada != nil {
		return _dada
	}
	_gfc := _f.Log(_gbab)
	_dada = _dcb.Push(MakeReal(_gfc))
	return _dada
}
func (_cfcc *PSOperand) lt(_fca *PSStack) error {
	_cbe, _bbb := _fca.PopNumberAsFloat64()
	if _bbb != nil {
		return _bbb
	}
	_aca, _bbb := _fca.PopNumberAsFloat64()
	if _bbb != nil {
		return _bbb
	}
	if _f.Abs(_aca-_cbe) < _fg {
		_dcbc := _fca.Push(MakeBool(false))
		return _dcbc
	} else if _aca < _cbe {
		_bgdb := _fca.Push(MakeBool(true))
		return _bgdb
	} else {
		_aee := _fca.Push(MakeBool(false))
		return _aee
	}
}

// Append appends an object to the PSProgram.
func (_gc *PSProgram) Append(obj PSObject) { *_gc = append(*_gc, obj) }
func (_gced *PSOperand) sqrt(_ccc *PSStack) error {
	_bddb, _edeb := _ccc.PopNumberAsFloat64()
	if _edeb != nil {
		return _edeb
	}
	if _bddb < 0 {
		return ErrRangeCheck
	}
	_ecce := _f.Sqrt(_bddb)
	_edeb = _ccc.Push(MakeReal(_ecce))
	return _edeb
}
func (_fb *PSInteger) Duplicate() PSObject { _cg := PSInteger{}; _cg.Val = _fb.Val; return &_cg }

var ErrRangeCheck = _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")

func (_cec *PSOperand) cos(_bedb *PSStack) error {
	_de, _fgfd := _bedb.PopNumberAsFloat64()
	if _fgfd != nil {
		return _fgfd
	}
	_gaa := _f.Cos(_de * _f.Pi / 180.0)
	_fgfd = _bedb.Push(MakeReal(_gaa))
	return _fgfd
}

var ErrUnsupportedOperand = _a.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")

func (_bgg *PSOperand) exch(_ggc *PSStack) error {
	_faf, _bdga := _ggc.Pop()
	if _bdga != nil {
		return _bdga
	}
	_bbgf, _bdga := _ggc.Pop()
	if _bdga != nil {
		return _bdga
	}
	_bdga = _ggc.Push(_faf)
	if _bdga != nil {
		return _bdga
	}
	_bdga = _ggc.Push(_bbgf)
	return _bdga
}
func (_gbcd *PSParser) parseOperand() (*PSOperand, error) {
	var _adag []byte
	for {
		_bcdb, _dfbdc := _gbcd._gdf.Peek(1)
		if _dfbdc != nil {
			if _dfbdc == _ab.EOF {
				break
			}
			return nil, _dfbdc
		}
		if _ed.IsDelimiter(_bcdb[0]) {
			break
		}
		if _ed.IsWhiteSpace(_bcdb[0]) {
			break
		}
		_ggdc, _ := _gbcd._gdf.ReadByte()
		_adag = append(_adag, _ggdc)
	}
	if len(_adag) == 0 {
		return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	return MakeOperand(string(_adag)), nil
}
