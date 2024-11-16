package graphic2d

import (
	_f "image/color"
	_cg "math"
)

func _gb(_gaf, _ae float64) bool { return _cg.Abs(_gaf-_ae) <= _edf }
func (_eee Point) Interpolate(q Point, t float64) Point {
	return Point{(1-t)*_eee.X + t*q.X, (1-t)*_eee.Y + t*q.Y}
}

type Point struct{ X, Y float64 }

func _dba(_ced, _ga, _ec, _adf, _ee, _fb float64) (float64, float64) {
	_eef, _gdf := _cg.Sincos(_fb)
	_ddg, _geg := _cg.Sincos(_ec)
	_b := _adf + _ced*_gdf*_geg - _ga*_eef*_ddg
	_bf := _ee + _ced*_gdf*_ddg + _ga*_eef*_geg
	return _b, _bf
}
func (_bde Point) Add(q Point) Point { return Point{_bde.X + q.X, _bde.Y + q.Y} }

var ColorMap = map[string]_f.RGBA{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e": _f.RGBA{0xf0, 0xf8, 0xff, 0xff}, "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065": _f.RGBA{0xfa, 0xeb, 0xd7, 0xff}, "\u0061\u0071\u0075\u0061": _f.RGBA{0x00, 0xff, 0xff, 0xff}, "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": _f.RGBA{0x7f, 0xff, 0xd4, 0xff}, "\u0061\u007a\u0075r\u0065": _f.RGBA{0xf0, 0xff, 0xff, 0xff}, "\u0062\u0065\u0069g\u0065": _f.RGBA{0xf5, 0xf5, 0xdc, 0xff}, "\u0062\u0069\u0073\u0071\u0075\u0065": _f.RGBA{0xff, 0xe4, 0xc4, 0xff}, "\u0062\u006c\u0061c\u006b": _f.RGBA{0x00, 0x00, 0x00, 0xff}, "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064": _f.RGBA{0xff, 0xeb, 0xcd, 0xff}, "\u0062\u006c\u0075\u0065": _f.RGBA{0x00, 0x00, 0xff, 0xff}, "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074": _f.RGBA{0x8a, 0x2b, 0xe2, 0xff}, "\u0062\u0072\u006fw\u006e": _f.RGBA{0xa5, 0x2a, 0x2a, 0xff}, "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd": _f.RGBA{0xde, 0xb8, 0x87, 0xff}, "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e": _f.RGBA{0x5f, 0x9e, 0xa0, 0xff}, "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065": _f.RGBA{0x7f, 0xff, 0x00, 0xff}, "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e": _f.RGBA{0xd2, 0x69, 0x1e, 0xff}, "\u0063\u006f\u0072a\u006c": _f.RGBA{0xff, 0x7f, 0x50, 0xff}, "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065": _f.RGBA{0x64, 0x95, 0xed, 0xff}, "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b": _f.RGBA{0xff, 0xf8, 0xdc, 0xff}, "\u0063r\u0069\u006d\u0073\u006f\u006e": _f.RGBA{0xdc, 0x14, 0x3c, 0xff}, "\u0063\u0079\u0061\u006e": _f.RGBA{0x00, 0xff, 0xff, 0xff}, "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065": _f.RGBA{0x00, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e": _f.RGBA{0x00, 0x8b, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": _f.RGBA{0xb8, 0x86, 0x0b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079": _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n": _f.RGBA{0x00, 0x64, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079": _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi": _f.RGBA{0xbd, 0xb7, 0x6b, 0xff}, "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061": _f.RGBA{0x8b, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e": _f.RGBA{0x55, 0x6b, 0x2f, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065": _f.RGBA{0xff, 0x8c, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064": _f.RGBA{0x99, 0x32, 0xcc, 0xff}, "\u0064a\u0072\u006b\u0072\u0065\u0064": _f.RGBA{0x8b, 0x00, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e": _f.RGBA{0xe9, 0x96, 0x7a, 0xff}, "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x8f, 0xbc, 0x8f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": _f.RGBA{0x48, 0x3d, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079": _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079": _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _f.RGBA{0x00, 0xce, 0xd1, 0xff}, "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074": _f.RGBA{0x94, 0x00, 0xd3, 0xff}, "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b": _f.RGBA{0xff, 0x14, 0x93, 0xff}, "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065": _f.RGBA{0x00, 0xbf, 0xff, 0xff}, "\u0064i\u006d\u0067\u0072\u0061\u0079": _f.RGBA{0x69, 0x69, 0x69, 0xff}, "\u0064i\u006d\u0067\u0072\u0065\u0079": _f.RGBA{0x69, 0x69, 0x69, 0xff}, "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065": _f.RGBA{0x1e, 0x90, 0xff, 0xff}, "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k": _f.RGBA{0xb2, 0x22, 0x22, 0xff}, "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065": _f.RGBA{0xff, 0xfa, 0xf0, 0xff}, "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x22, 0x8b, 0x22, 0xff}, "\u0066u\u0063\u0068\u0073\u0069\u0061": _f.RGBA{0xff, 0x00, 0xff, 0xff}, "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o": _f.RGBA{0xdc, 0xdc, 0xdc, 0xff}, "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065": _f.RGBA{0xf8, 0xf8, 0xff, 0xff}, "\u0067\u006f\u006c\u0064": _f.RGBA{0xff, 0xd7, 0x00, 0xff}, "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd": _f.RGBA{0xda, 0xa5, 0x20, 0xff}, "\u0067\u0072\u0061\u0079": _f.RGBA{0x80, 0x80, 0x80, 0xff}, "\u0067\u0072\u0065e\u006e": _f.RGBA{0x00, 0x80, 0x00, 0xff}, "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077": _f.RGBA{0xad, 0xff, 0x2f, 0xff}, "\u0067\u0072\u0065\u0079": _f.RGBA{0x80, 0x80, 0x80, 0xff}, "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077": _f.RGBA{0xf0, 0xff, 0xf0, 0xff}, "\u0068o\u0074\u0070\u0069\u006e\u006b": _f.RGBA{0xff, 0x69, 0xb4, 0xff}, "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d": _f.RGBA{0xcd, 0x5c, 0x5c, 0xff}, "\u0069\u006e\u0064\u0069\u0067\u006f": _f.RGBA{0x4b, 0x00, 0x82, 0xff}, "\u0069\u0076\u006fr\u0079": _f.RGBA{0xff, 0xff, 0xf0, 0xff}, "\u006b\u0068\u0061k\u0069": _f.RGBA{0xf0, 0xe6, 0x8c, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072": _f.RGBA{0xe6, 0xe6, 0xfa, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068": _f.RGBA{0xff, 0xf0, 0xf5, 0xff}, "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n": _f.RGBA{0x7c, 0xfc, 0x00, 0xff}, "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e": _f.RGBA{0xff, 0xfa, 0xcd, 0xff}, "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e": _f.RGBA{0xad, 0xd8, 0xe6, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c": _f.RGBA{0xf0, 0x80, 0x80, 0xff}, "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n": _f.RGBA{0xe0, 0xff, 0xff, 0xff}, "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077": _f.RGBA{0xfa, 0xfa, 0xd2, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y": _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x90, 0xee, 0x90, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y": _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}, "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek": _f.RGBA{0xff, 0xb6, 0xc1, 0xff}, "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e": _f.RGBA{0xff, 0xa0, 0x7a, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x20, 0xb2, 0xaa, 0xff}, "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065": _f.RGBA{0x87, 0xce, 0xfa, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079": _f.RGBA{0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079": _f.RGBA{0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065": _f.RGBA{0xb0, 0xc4, 0xde, 0xff}, "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077": _f.RGBA{0xff, 0xff, 0xe0, 0xff}, "\u006c\u0069\u006d\u0065": _f.RGBA{0x00, 0xff, 0x00, 0xff}, "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n": _f.RGBA{0x32, 0xcd, 0x32, 0xff}, "\u006c\u0069\u006ee\u006e": _f.RGBA{0xfa, 0xf0, 0xe6, 0xff}, "\u006da\u0067\u0065\u006e\u0074\u0061": _f.RGBA{0xff, 0x00, 0xff, 0xff}, "\u006d\u0061\u0072\u006f\u006f\u006e": _f.RGBA{0x80, 0x00, 0x00, 0xff}, "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": _f.RGBA{0x66, 0xcd, 0xaa, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065": _f.RGBA{0x00, 0x00, 0xcd, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064": _f.RGBA{0xba, 0x55, 0xd3, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065": _f.RGBA{0x93, 0x70, 0xdb, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e": _f.RGBA{0x3c, 0xb3, 0x71, 0xff}, "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": _f.RGBA{0x7b, 0x68, 0xee, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e": _f.RGBA{0x00, 0xfa, 0x9a, 0xff}, "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _f.RGBA{0x48, 0xd1, 0xcc, 0xff}, "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": _f.RGBA{0xc7, 0x15, 0x85, 0xff}, "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065": _f.RGBA{0x19, 0x19, 0x70, 0xff}, "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m": _f.RGBA{0xf5, 0xff, 0xfa, 0xff}, "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e": _f.RGBA{0xff, 0xe4, 0xe1, 0xff}, "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e": _f.RGBA{0xff, 0xe4, 0xb5, 0xff}, "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065": _f.RGBA{0xff, 0xde, 0xad, 0xff}, "\u006e\u0061\u0076\u0079": _f.RGBA{0x00, 0x00, 0x80, 0xff}, "\u006fl\u0064\u006c\u0061\u0063\u0065": _f.RGBA{0xfd, 0xf5, 0xe6, 0xff}, "\u006f\u006c\u0069v\u0065": _f.RGBA{0x80, 0x80, 0x00, 0xff}, "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b": _f.RGBA{0x6b, 0x8e, 0x23, 0xff}, "\u006f\u0072\u0061\u006e\u0067\u0065": _f.RGBA{0xff, 0xa5, 0x00, 0xff}, "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d": _f.RGBA{0xff, 0x45, 0x00, 0xff}, "\u006f\u0072\u0063\u0068\u0069\u0064": _f.RGBA{0xda, 0x70, 0xd6, 0xff}, "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": _f.RGBA{0xee, 0xe8, 0xaa, 0xff}, "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n": _f.RGBA{0x98, 0xfb, 0x98, 0xff}, "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _f.RGBA{0xaf, 0xee, 0xee, 0xff}, "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": _f.RGBA{0xdb, 0x70, 0x93, 0xff}, "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070": _f.RGBA{0xff, 0xef, 0xd5, 0xff}, "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f": _f.RGBA{0xff, 0xda, 0xb9, 0xff}, "\u0070\u0065\u0072\u0075": _f.RGBA{0xcd, 0x85, 0x3f, 0xff}, "\u0070\u0069\u006e\u006b": _f.RGBA{0xff, 0xc0, 0xcb, 0xff}, "\u0070\u006c\u0075\u006d": _f.RGBA{0xdd, 0xa0, 0xdd, 0xff}, "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065": _f.RGBA{0xb0, 0xe0, 0xe6, 0xff}, "\u0070\u0075\u0072\u0070\u006c\u0065": _f.RGBA{0x80, 0x00, 0x80, 0xff}, "\u0072\u0065\u0064": _f.RGBA{0xff, 0x00, 0x00, 0xff}, "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n": _f.RGBA{0xbc, 0x8f, 0x8f, 0xff}, "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e": _f.RGBA{0x41, 0x69, 0xe1, 0xff}, "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e": _f.RGBA{0x8b, 0x45, 0x13, 0xff}, "\u0073\u0061\u006c\u006d\u006f\u006e": _f.RGBA{0xfa, 0x80, 0x72, 0xff}, "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e": _f.RGBA{0xf4, 0xa4, 0x60, 0xff}, "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x2e, 0x8b, 0x57, 0xff}, "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c": _f.RGBA{0xff, 0xf5, 0xee, 0xff}, "\u0073\u0069\u0065\u006e\u006e\u0061": _f.RGBA{0xa0, 0x52, 0x2d, 0xff}, "\u0073\u0069\u006c\u0076\u0065\u0072": _f.RGBA{0xc0, 0xc0, 0xc0, 0xff}, "\u0073k\u0079\u0062\u006c\u0075\u0065": _f.RGBA{0x87, 0xce, 0xeb, 0xff}, "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e": _f.RGBA{0x6a, 0x5a, 0xcd, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y": _f.RGBA{0x70, 0x80, 0x90, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y": _f.RGBA{0x70, 0x80, 0x90, 0xff}, "\u0073\u006e\u006f\u0077": _f.RGBA{0xff, 0xfa, 0xfa, 0xff}, "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x00, 0xff, 0x7f, 0xff}, "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e": _f.RGBA{0x46, 0x82, 0xb4, 0xff}, "\u0074\u0061\u006e": _f.RGBA{0xd2, 0xb4, 0x8c, 0xff}, "\u0074\u0065\u0061\u006c": _f.RGBA{0x00, 0x80, 0x80, 0xff}, "\u0074h\u0069\u0073\u0074\u006c\u0065": _f.RGBA{0xd8, 0xbf, 0xd8, 0xff}, "\u0074\u006f\u006d\u0061\u0074\u006f": _f.RGBA{0xff, 0x63, 0x47, 0xff}, "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e": _f.RGBA{0x40, 0xe0, 0xd0, 0xff}, "\u0076\u0069\u006f\u006c\u0065\u0074": _f.RGBA{0xee, 0x82, 0xee, 0xff}, "\u0077\u0068\u0065a\u0074": _f.RGBA{0xf5, 0xde, 0xb3, 0xff}, "\u0077\u0068\u0069t\u0065": _f.RGBA{0xff, 0xff, 0xff, 0xff}, "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065": _f.RGBA{0xf5, 0xf5, 0xf5, 0xff}, "\u0079\u0065\u006c\u006c\u006f\u0077": _f.RGBA{0xff, 0xff, 0x00, 0xff}, "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e": _f.RGBA{0x9a, 0xcd, 0x32, 0xff}}
var Names = []string{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e", "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065", "\u0061\u0071\u0075\u0061", "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u0061\u007a\u0075r\u0065", "\u0062\u0065\u0069g\u0065", "\u0062\u0069\u0073\u0071\u0075\u0065", "\u0062\u006c\u0061c\u006b", "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064", "\u0062\u006c\u0075\u0065", "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074", "\u0062\u0072\u006fw\u006e", "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd", "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e", "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065", "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e", "\u0063\u006f\u0072a\u006c", "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065", "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b", "\u0063r\u0069\u006d\u0073\u006f\u006e", "\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079", "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n", "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079", "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi", "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061", "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065", "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064", "\u0064a\u0072\u006b\u0072\u0065\u0064", "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e", "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079", "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074", "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b", "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u0064i\u006d\u0067\u0072\u0061\u0079", "\u0064i\u006d\u0067\u0072\u0065\u0079", "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065", "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k", "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065", "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e", "\u0066u\u0063\u0068\u0073\u0069\u0061", "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o", "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065", "\u0067\u006f\u006c\u0064", "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd", "\u0067\u0072\u0061\u0079", "\u0067\u0072\u0065e\u006e", "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077", "\u0067\u0072\u0065\u0079", "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077", "\u0068o\u0074\u0070\u0069\u006e\u006b", "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d", "\u0069\u006e\u0064\u0069\u0067\u006f", "\u0069\u0076\u006fr\u0079", "\u006b\u0068\u0061k\u0069", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068", "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n", "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e", "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e", "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c", "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n", "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y", "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y", "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek", "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e", "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065", "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077", "\u006c\u0069\u006d\u0065", "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n", "\u006c\u0069\u006ee\u006e", "\u006da\u0067\u0065\u006e\u0074\u0061", "\u006d\u0061\u0072\u006f\u006f\u006e", "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064", "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065", "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m", "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e", "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e", "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065", "\u006e\u0061\u0076\u0079", "\u006fl\u0064\u006c\u0061\u0063\u0065", "\u006f\u006c\u0069v\u0065", "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b", "\u006f\u0072\u0061\u006e\u0067\u0065", "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d", "\u006f\u0072\u0063\u0068\u0069\u0064", "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n", "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070", "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f", "\u0070\u0065\u0072\u0075", "\u0070\u0069\u006e\u006b", "\u0070\u006c\u0075\u006d", "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065", "\u0070\u0075\u0072\u0070\u006c\u0065", "\u0072\u0065\u0064", "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n", "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e", "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e", "\u0073\u0061\u006c\u006d\u006f\u006e", "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e", "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c", "\u0073\u0069\u0065\u006e\u006e\u0061", "\u0073\u0069\u006c\u0076\u0065\u0072", "\u0073k\u0079\u0062\u006c\u0075\u0065", "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y", "\u0073\u006e\u006f\u0077", "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e", "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e", "\u0074\u0061\u006e", "\u0074\u0065\u0061\u006c", "\u0074h\u0069\u0073\u0074\u006c\u0065", "\u0074\u006f\u006d\u0061\u0074\u006f", "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e", "\u0076\u0069\u006f\u006c\u0065\u0074", "\u0077\u0068\u0065a\u0074", "\u0077\u0068\u0069t\u0065", "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065", "\u0079\u0065\u006c\u006c\u006f\u0077", "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e"}

const _edf = 1e-10

func (_bfe Point) Mul(f float64) Point { return Point{f * _bfe.X, f * _bfe.Y} }

var (
	Aliceblue            = _f.RGBA{0xf0, 0xf8, 0xff, 0xff}
	Antiquewhite         = _f.RGBA{0xfa, 0xeb, 0xd7, 0xff}
	Aqua                 = _f.RGBA{0x00, 0xff, 0xff, 0xff}
	Aquamarine           = _f.RGBA{0x7f, 0xff, 0xd4, 0xff}
	Azure                = _f.RGBA{0xf0, 0xff, 0xff, 0xff}
	Beige                = _f.RGBA{0xf5, 0xf5, 0xdc, 0xff}
	Bisque               = _f.RGBA{0xff, 0xe4, 0xc4, 0xff}
	Black                = _f.RGBA{0x00, 0x00, 0x00, 0xff}
	Blanchedalmond       = _f.RGBA{0xff, 0xeb, 0xcd, 0xff}
	Blue                 = _f.RGBA{0x00, 0x00, 0xff, 0xff}
	Blueviolet           = _f.RGBA{0x8a, 0x2b, 0xe2, 0xff}
	Brown                = _f.RGBA{0xa5, 0x2a, 0x2a, 0xff}
	Burlywood            = _f.RGBA{0xde, 0xb8, 0x87, 0xff}
	Cadetblue            = _f.RGBA{0x5f, 0x9e, 0xa0, 0xff}
	Chartreuse           = _f.RGBA{0x7f, 0xff, 0x00, 0xff}
	Chocolate            = _f.RGBA{0xd2, 0x69, 0x1e, 0xff}
	Coral                = _f.RGBA{0xff, 0x7f, 0x50, 0xff}
	Cornflowerblue       = _f.RGBA{0x64, 0x95, 0xed, 0xff}
	Cornsilk             = _f.RGBA{0xff, 0xf8, 0xdc, 0xff}
	Crimson              = _f.RGBA{0xdc, 0x14, 0x3c, 0xff}
	Cyan                 = _f.RGBA{0x00, 0xff, 0xff, 0xff}
	Darkblue             = _f.RGBA{0x00, 0x00, 0x8b, 0xff}
	Darkcyan             = _f.RGBA{0x00, 0x8b, 0x8b, 0xff}
	Darkgoldenrod        = _f.RGBA{0xb8, 0x86, 0x0b, 0xff}
	Darkgray             = _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkgreen            = _f.RGBA{0x00, 0x64, 0x00, 0xff}
	Darkgrey             = _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkkhaki            = _f.RGBA{0xbd, 0xb7, 0x6b, 0xff}
	Darkmagenta          = _f.RGBA{0x8b, 0x00, 0x8b, 0xff}
	Darkolivegreen       = _f.RGBA{0x55, 0x6b, 0x2f, 0xff}
	Darkorange           = _f.RGBA{0xff, 0x8c, 0x00, 0xff}
	Darkorchid           = _f.RGBA{0x99, 0x32, 0xcc, 0xff}
	Darkred              = _f.RGBA{0x8b, 0x00, 0x00, 0xff}
	Darksalmon           = _f.RGBA{0xe9, 0x96, 0x7a, 0xff}
	Darkseagreen         = _f.RGBA{0x8f, 0xbc, 0x8f, 0xff}
	Darkslateblue        = _f.RGBA{0x48, 0x3d, 0x8b, 0xff}
	Darkslategray        = _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkslategrey        = _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkturquoise        = _f.RGBA{0x00, 0xce, 0xd1, 0xff}
	Darkviolet           = _f.RGBA{0x94, 0x00, 0xd3, 0xff}
	Deeppink             = _f.RGBA{0xff, 0x14, 0x93, 0xff}
	Deepskyblue          = _f.RGBA{0x00, 0xbf, 0xff, 0xff}
	Dimgray              = _f.RGBA{0x69, 0x69, 0x69, 0xff}
	Dimgrey              = _f.RGBA{0x69, 0x69, 0x69, 0xff}
	Dodgerblue           = _f.RGBA{0x1e, 0x90, 0xff, 0xff}
	Firebrick            = _f.RGBA{0xb2, 0x22, 0x22, 0xff}
	Floralwhite          = _f.RGBA{0xff, 0xfa, 0xf0, 0xff}
	Forestgreen          = _f.RGBA{0x22, 0x8b, 0x22, 0xff}
	Fuchsia              = _f.RGBA{0xff, 0x00, 0xff, 0xff}
	Gainsboro            = _f.RGBA{0xdc, 0xdc, 0xdc, 0xff}
	Ghostwhite           = _f.RGBA{0xf8, 0xf8, 0xff, 0xff}
	Gold                 = _f.RGBA{0xff, 0xd7, 0x00, 0xff}
	Goldenrod            = _f.RGBA{0xda, 0xa5, 0x20, 0xff}
	Gray                 = _f.RGBA{0x80, 0x80, 0x80, 0xff}
	Green                = _f.RGBA{0x00, 0x80, 0x00, 0xff}
	Greenyellow          = _f.RGBA{0xad, 0xff, 0x2f, 0xff}
	Grey                 = _f.RGBA{0x80, 0x80, 0x80, 0xff}
	Honeydew             = _f.RGBA{0xf0, 0xff, 0xf0, 0xff}
	Hotpink              = _f.RGBA{0xff, 0x69, 0xb4, 0xff}
	Indianred            = _f.RGBA{0xcd, 0x5c, 0x5c, 0xff}
	Indigo               = _f.RGBA{0x4b, 0x00, 0x82, 0xff}
	Ivory                = _f.RGBA{0xff, 0xff, 0xf0, 0xff}
	Khaki                = _f.RGBA{0xf0, 0xe6, 0x8c, 0xff}
	Lavender             = _f.RGBA{0xe6, 0xe6, 0xfa, 0xff}
	Lavenderblush        = _f.RGBA{0xff, 0xf0, 0xf5, 0xff}
	Lawngreen            = _f.RGBA{0x7c, 0xfc, 0x00, 0xff}
	Lemonchiffon         = _f.RGBA{0xff, 0xfa, 0xcd, 0xff}
	Lightblue            = _f.RGBA{0xad, 0xd8, 0xe6, 0xff}
	Lightcoral           = _f.RGBA{0xf0, 0x80, 0x80, 0xff}
	Lightcyan            = _f.RGBA{0xe0, 0xff, 0xff, 0xff}
	Lightgoldenrodyellow = _f.RGBA{0xfa, 0xfa, 0xd2, 0xff}
	Lightgray            = _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightgreen           = _f.RGBA{0x90, 0xee, 0x90, 0xff}
	Lightgrey            = _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightpink            = _f.RGBA{0xff, 0xb6, 0xc1, 0xff}
	Lightsalmon          = _f.RGBA{0xff, 0xa0, 0x7a, 0xff}
	Lightseagreen        = _f.RGBA{0x20, 0xb2, 0xaa, 0xff}
	Lightskyblue         = _f.RGBA{0x87, 0xce, 0xfa, 0xff}
	Lightslategray       = _f.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightslategrey       = _f.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightsteelblue       = _f.RGBA{0xb0, 0xc4, 0xde, 0xff}
	Lightyellow          = _f.RGBA{0xff, 0xff, 0xe0, 0xff}
	Lime                 = _f.RGBA{0x00, 0xff, 0x00, 0xff}
	Limegreen            = _f.RGBA{0x32, 0xcd, 0x32, 0xff}
	Linen                = _f.RGBA{0xfa, 0xf0, 0xe6, 0xff}
	Magenta              = _f.RGBA{0xff, 0x00, 0xff, 0xff}
	Maroon               = _f.RGBA{0x80, 0x00, 0x00, 0xff}
	Mediumaquamarine     = _f.RGBA{0x66, 0xcd, 0xaa, 0xff}
	Mediumblue           = _f.RGBA{0x00, 0x00, 0xcd, 0xff}
	Mediumorchid         = _f.RGBA{0xba, 0x55, 0xd3, 0xff}
	Mediumpurple         = _f.RGBA{0x93, 0x70, 0xdb, 0xff}
	Mediumseagreen       = _f.RGBA{0x3c, 0xb3, 0x71, 0xff}
	Mediumslateblue      = _f.RGBA{0x7b, 0x68, 0xee, 0xff}
	Mediumspringgreen    = _f.RGBA{0x00, 0xfa, 0x9a, 0xff}
	Mediumturquoise      = _f.RGBA{0x48, 0xd1, 0xcc, 0xff}
	Mediumvioletred      = _f.RGBA{0xc7, 0x15, 0x85, 0xff}
	Midnightblue         = _f.RGBA{0x19, 0x19, 0x70, 0xff}
	Mintcream            = _f.RGBA{0xf5, 0xff, 0xfa, 0xff}
	Mistyrose            = _f.RGBA{0xff, 0xe4, 0xe1, 0xff}
	Moccasin             = _f.RGBA{0xff, 0xe4, 0xb5, 0xff}
	Navajowhite          = _f.RGBA{0xff, 0xde, 0xad, 0xff}
	Navy                 = _f.RGBA{0x00, 0x00, 0x80, 0xff}
	Oldlace              = _f.RGBA{0xfd, 0xf5, 0xe6, 0xff}
	Olive                = _f.RGBA{0x80, 0x80, 0x00, 0xff}
	Olivedrab            = _f.RGBA{0x6b, 0x8e, 0x23, 0xff}
	Orange               = _f.RGBA{0xff, 0xa5, 0x00, 0xff}
	Orangered            = _f.RGBA{0xff, 0x45, 0x00, 0xff}
	Orchid               = _f.RGBA{0xda, 0x70, 0xd6, 0xff}
	Palegoldenrod        = _f.RGBA{0xee, 0xe8, 0xaa, 0xff}
	Palegreen            = _f.RGBA{0x98, 0xfb, 0x98, 0xff}
	Paleturquoise        = _f.RGBA{0xaf, 0xee, 0xee, 0xff}
	Palevioletred        = _f.RGBA{0xdb, 0x70, 0x93, 0xff}
	Papayawhip           = _f.RGBA{0xff, 0xef, 0xd5, 0xff}
	Peachpuff            = _f.RGBA{0xff, 0xda, 0xb9, 0xff}
	Peru                 = _f.RGBA{0xcd, 0x85, 0x3f, 0xff}
	Pink                 = _f.RGBA{0xff, 0xc0, 0xcb, 0xff}
	Plum                 = _f.RGBA{0xdd, 0xa0, 0xdd, 0xff}
	Powderblue           = _f.RGBA{0xb0, 0xe0, 0xe6, 0xff}
	Purple               = _f.RGBA{0x80, 0x00, 0x80, 0xff}
	Red                  = _f.RGBA{0xff, 0x00, 0x00, 0xff}
	Rosybrown            = _f.RGBA{0xbc, 0x8f, 0x8f, 0xff}
	Royalblue            = _f.RGBA{0x41, 0x69, 0xe1, 0xff}
	Saddlebrown          = _f.RGBA{0x8b, 0x45, 0x13, 0xff}
	Salmon               = _f.RGBA{0xfa, 0x80, 0x72, 0xff}
	Sandybrown           = _f.RGBA{0xf4, 0xa4, 0x60, 0xff}
	Seagreen             = _f.RGBA{0x2e, 0x8b, 0x57, 0xff}
	Seashell             = _f.RGBA{0xff, 0xf5, 0xee, 0xff}
	Sienna               = _f.RGBA{0xa0, 0x52, 0x2d, 0xff}
	Silver               = _f.RGBA{0xc0, 0xc0, 0xc0, 0xff}
	Skyblue              = _f.RGBA{0x87, 0xce, 0xeb, 0xff}
	Slateblue            = _f.RGBA{0x6a, 0x5a, 0xcd, 0xff}
	Slategray            = _f.RGBA{0x70, 0x80, 0x90, 0xff}
	Slategrey            = _f.RGBA{0x70, 0x80, 0x90, 0xff}
	Snow                 = _f.RGBA{0xff, 0xfa, 0xfa, 0xff}
	Springgreen          = _f.RGBA{0x00, 0xff, 0x7f, 0xff}
	Steelblue            = _f.RGBA{0x46, 0x82, 0xb4, 0xff}
	Tan                  = _f.RGBA{0xd2, 0xb4, 0x8c, 0xff}
	Teal                 = _f.RGBA{0x00, 0x80, 0x80, 0xff}
	Thistle              = _f.RGBA{0xd8, 0xbf, 0xd8, 0xff}
	Tomato               = _f.RGBA{0xff, 0x63, 0x47, 0xff}
	Turquoise            = _f.RGBA{0x40, 0xe0, 0xd0, 0xff}
	Violet               = _f.RGBA{0xee, 0x82, 0xee, 0xff}
	Wheat                = _f.RGBA{0xf5, 0xde, 0xb3, 0xff}
	White                = _f.RGBA{0xff, 0xff, 0xff, 0xff}
	Whitesmoke           = _f.RGBA{0xf5, 0xf5, 0xf5, 0xff}
	Yellow               = _f.RGBA{0xff, 0xff, 0x00, 0xff}
	Yellowgreen          = _f.RGBA{0x9a, 0xcd, 0x32, 0xff}
)

func QuadraticToCubicBezier(startX, startY, x1, y1, x, y float64) (Point, Point) {
	_cba := Point{X: startX, Y: startY}
	_ef := Point{X: x1, Y: y1}
	_bgc := Point{X: x, Y: y}
	_bga := _cba.Interpolate(_ef, 2.0/3.0)
	_cgc := _bgc.Interpolate(_ef, 2.0/3.0)
	return _bga, _cgc
}
func EllipseToCubicBeziers(startX, startY, rx, ry, rot float64, large, sweep bool, endX, endY float64) [][4]Point {
	rx = _cg.Abs(rx)
	ry = _cg.Abs(ry)
	if rx < ry {
		rx, ry = ry, rx
		rot += 90.0
	}
	_g := _ggg(rot * _cg.Pi / 180.0)
	if _cg.Pi <= _g {
		_g -= _cg.Pi
	}
	_a, _gg, _fd, _fdd := _bg(startX, startY, rx, ry, _g, large, sweep, endX, endY)
	_gd := _cg.Pi / 2.0
	_d := int(_cg.Ceil(_cg.Abs(_fdd-_fd) / _gd))
	_gd = _cg.Abs(_fdd-_fd) / float64(_d)
	_ac := _cg.Sin(_gd) * (_cg.Sqrt(4.0+3.0*_cg.Pow(_cg.Tan(_gd/2.0), 2.0)) - 1.0) / 3.0
	if !sweep {
		_gd = -_gd
	}
	_acf := Point{X: startX, Y: startY}
	_dd, _ag := _eca(rx, ry, _g, sweep, _fd)
	_ad := Point{X: _dd, Y: _ag}
	_ace := [][4]Point{}
	for _ge := 1; _ge < _d+1; _ge++ {
		_gc := _fd + float64(_ge)*_gd
		_agb, _gee := _dba(rx, ry, _g, _a, _gg, _gc)
		_e := Point{X: _agb, Y: _gee}
		_db, _adg := _eca(rx, ry, _g, sweep, _gc)
		_ca := Point{X: _db, Y: _adg}
		_ce := _acf.Add(_ad.Mul(_ac))
		_de := _e.Sub(_ca.Mul(_ac))
		_ace = append(_ace, [4]Point{_acf, _ce, _de, _e})
		_ad = _ca
		_acf = _e
	}
	return _ace
}
func (_cfc Point) Sub(q Point) Point { return Point{_cfc.X - q.X, _cfc.Y - q.Y} }
func _bg(_ba, _gf, _bb, _fbg, _def float64, _cf, _fbgd bool, _bba, _cfb float64) (float64, float64, float64, float64) {
	if _gb(_ba, _bba) && _gb(_gf, _cfb) {
		return _ba, _gf, 0.0, 0.0
	}
	_eb, _ebe := _cg.Sincos(_def)
	_be := _ebe*(_ba-_bba)/2.0 + _eb*(_gf-_cfb)/2.0
	_bc := -_eb*(_ba-_bba)/2.0 + _ebe*(_gf-_cfb)/2.0
	_cb := _be*_be/_bb/_bb + _bc*_bc/_fbg/_fbg
	if _cb > 1.0 {
		_bb *= _cg.Sqrt(_cb)
		_fbg *= _cg.Sqrt(_cb)
	}
	_bef := (_bb*_bb*_fbg*_fbg - _bb*_bb*_bc*_bc - _fbg*_fbg*_be*_be) / (_bb*_bb*_bc*_bc + _fbg*_fbg*_be*_be)
	if _bef < 0.0 {
		_bef = 0.0
	}
	_bgd := _cg.Sqrt(_bef)
	if _cf == _fbgd {
		_bgd = -_bgd
	}
	_fe := _bgd * _bb * _bc / _fbg
	_ecb := _bgd * -_fbg * _be / _bb
	_fde := _ebe*_fe - _eb*_ecb + (_ba+_bba)/2.0
	_bfb := _eb*_fe + _ebe*_ecb + (_gf+_cfb)/2.0
	_eea := (_be - _fe) / _bb
	_aa := (_bc - _ecb) / _fbg
	_caf := -(_be + _fe) / _bb
	_bae := -(_bc + _ecb) / _fbg
	_bd := _cg.Acos(_eea / _cg.Sqrt(_eea*_eea+_aa*_aa))
	if _aa < 0.0 {
		_bd = -_bd
	}
	_bd = _ggg(_bd)
	_fec := (_eea*_caf + _aa*_bae) / _cg.Sqrt((_eea*_eea+_aa*_aa)*(_caf*_caf+_bae*_bae))
	_fec = _cg.Min(1.0, _cg.Max(-1.0, _fec))
	_cc := _cg.Acos(_fec)
	if _eea*_bae-_aa*_caf < 0.0 {
		_cc = -_cc
	}
	if !_fbgd && _cc > 0.0 {
		_cc -= 2.0 * _cg.Pi
	} else if _fbgd && _cc < 0.0 {
		_cc += 2.0 * _cg.Pi
	}
	return _fde, _bfb, _bd, _bd + _cc
}
func _ggg(_fc float64) float64 {
	_fc = _cg.Mod(_fc, 2.0*_cg.Pi)
	if _fc < 0.0 {
		_fc += 2.0 * _cg.Pi
	}
	return _fc
}
func _eca(_df, _ed, _fed float64, _deb bool, _gcc float64) (float64, float64) {
	_aac, _fg := _cg.Sincos(_gcc)
	_ceg, _ff := _cg.Sincos(_fed)
	_cgg := -_df*_aac*_ff - _ed*_fg*_ceg
	_fba := -_df*_aac*_ceg + _ed*_fg*_ff
	if !_deb {
		return -_cgg, -_fba
	}
	return _cgg, _fba
}
