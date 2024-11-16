package decoder

import (
	_fc "image"

	_b "github.com/bamzi/pdfext/internal/bitwise"
	_bg "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_bf "github.com/bamzi/pdfext/internal/jbig2/document"
	_e "github.com/bamzi/pdfext/internal/jbig2/errors"
)

func (_g *Decoder) DecodePage(pageNumber int) ([]byte, error) { return _g.decodePage(pageNumber) }

type Parameters struct {
	UnpaddedData bool
	Color        _bg.Color
}

func (_cc *Decoder) DecodeNextPage() ([]byte, error) {
	_cc._ea++
	_ag := _cc._ea
	return _cc.decodePage(_ag)
}
func (_ec *Decoder) PageNumber() (int, error) {
	const _gge = "\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072"
	if _ec._c == nil {
		return 0, _e.Error(_gge, "d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074")
	}
	return int(_ec._c.NumberOfPages), nil
}
func (_a *Decoder) DecodePageImage(pageNumber int) (_fc.Image, error) {
	const _fa = "\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	_cf, _gg := _a.decodePageImage(pageNumber)
	if _gg != nil {
		return nil, _e.Wrap(_gg, _fa, "")
	}
	return _cf, nil
}
func Decode(input []byte, parameters Parameters, globals *_bf.Globals) (*Decoder, error) {
	_cca := _b.NewReader(input)
	_dd, _eeg := _bf.DecodeDocument(_cca, globals)
	if _eeg != nil {
		return nil, _eeg
	}
	return &Decoder{_fd: _cca, _c: _dd, _eaf: parameters}, nil
}
func (_d *Decoder) decodePage(_ee int) ([]byte, error) {
	const _gb = "\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065"
	if _ee < 0 {
		return nil, _e.Errorf(_gb, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _ee)
	}
	if _ee > int(_d._c.NumberOfPages) {
		return nil, _e.Errorf(_gb, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _ee)
	}
	_bge, _ce := _d._c.GetPage(_ee)
	if _ce != nil {
		return nil, _e.Wrap(_ce, _gb, "")
	}
	_bgg, _ce := _bge.GetBitmap()
	if _ce != nil {
		return nil, _e.Wrap(_ce, _gb, "")
	}
	_bgg.InverseData()
	if !_d._eaf.UnpaddedData {
		return _bgg.Data, nil
	}
	return _bgg.GetUnpaddedData()
}

type Decoder struct {
	_fd  *_b.Reader
	_c   *_bf.Document
	_ea  int
	_eaf Parameters
}

func (_eg *Decoder) decodePageImage(_db int) (_fc.Image, error) {
	const _ef = "\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	if _db < 0 {
		return nil, _e.Errorf(_ef, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _db)
	}
	if _db > int(_eg._c.NumberOfPages) {
		return nil, _e.Errorf(_ef, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _db)
	}
	_agg, _fe := _eg._c.GetPage(_db)
	if _fe != nil {
		return nil, _e.Wrap(_fe, _ef, "")
	}
	_gf, _fe := _agg.GetBitmap()
	if _fe != nil {
		return nil, _e.Wrap(_fe, _ef, "")
	}
	_gf.InverseData()
	return _gf.ToImage(), nil
}
