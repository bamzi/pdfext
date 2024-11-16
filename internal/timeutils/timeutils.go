package timeutils

import (
	_bf "errors"
	_ff "fmt"
	_ba "regexp"
	_bg "strconv"
	_f "time"
)

var _ag = _ba.MustCompile("\u005cs\u002a\u0044\u005cs\u002a\u003a\u005cs\u002a(\\\u0064\u007b\u0034\u007d\u0029\u0028\u005cd\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064{2\u007d)\u003f\u0028\u005b\u002b\u002d\u005a]\u0029\u003f\u0028\u005c\u0064{\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032}\u0029\u003f")

func FormatPdfTime(in _f.Time) string {
	_c := in.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_g, _ := _bg.ParseInt(_c[1:3], 10, 32)
	_d, _ := _bg.ParseInt(_c[4:6], 10, 32)
	_ge := int64(in.Year())
	_gf := int64(in.Month())
	_fg := int64(in.Day())
	_baf := int64(in.Hour())
	_ga := int64(in.Minute())
	_fe := int64(in.Second())
	_bb := _c[0]
	return _ff.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _ge, _gf, _fg, _baf, _ga, _fe, _bb, _g, _d)
}
func ParsePdfTime(pdfTime string) (_f.Time, error) {
	_baff := _ag.FindAllStringSubmatch(pdfTime, 1)
	if len(_baff) < 1 {
		if len(pdfTime) > 0 && pdfTime[0] != 'D' {
			pdfTime = _ff.Sprintf("\u0044\u003a\u0025\u0073", pdfTime)
			return ParsePdfTime(pdfTime)
		}
		return _f.Time{}, _ff.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029", pdfTime)
	}
	if len(_baff[0]) != 10 {
		return _f.Time{}, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030")
	}
	_a, _ := _bg.ParseInt(_baff[0][1], 10, 32)
	_cc, _ := _bg.ParseInt(_baff[0][2], 10, 32)
	_e, _ := _bg.ParseInt(_baff[0][3], 10, 32)
	_dcd, _ := _bg.ParseInt(_baff[0][4], 10, 32)
	_fa, _ := _bg.ParseInt(_baff[0][5], 10, 32)
	_bfg, _ := _bg.ParseInt(_baff[0][6], 10, 32)
	var (
		_ca  byte
		_ea  int64
		_fag int64
	)
	_ca = '+'
	if len(_baff[0][7]) > 0 {
		if _baff[0][7] == "\u002d" {
			_ca = '-'
		} else if _baff[0][7] == "\u005a" {
			_ca = 'Z'
		}
	}
	if len(_baff[0][8]) > 0 {
		_ea, _ = _bg.ParseInt(_baff[0][8], 10, 32)
	} else {
		_ea = 0
	}
	if len(_baff[0][9]) > 0 {
		_fag, _ = _bg.ParseInt(_baff[0][9], 10, 32)
	} else {
		_fag = 0
	}
	_gd := int(_ea*60*60 + _fag*60)
	switch _ca {
	case '-':
		_gd = -_gd
	case 'Z':
		_gd = 0
	}
	_bac := _ff.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _ca, _ea, _fag)
	_gdb := _f.FixedZone(_bac, _gd)
	return _f.Date(int(_a), _f.Month(_cc), int(_e), int(_dcd), int(_fa), int(_bfg), 0, _gdb), nil
}
