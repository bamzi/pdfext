package pdfutil

import (
	_d "github.com/bamzi/pdfext/common"
	_b "github.com/bamzi/pdfext/contentstream"
	_e "github.com/bamzi/pdfext/contentstream/draw"
	_dc "github.com/bamzi/pdfext/core"
	_de "github.com/bamzi/pdfext/model"
)

// NormalizePage performs the following operations on the passed in page:
//   - Normalize the page rotation.
//     Rotates the contents of the page according to the Rotate entry, thus
//     flattening the rotation. The Rotate entry of the page is set to nil.
//   - Normalize the media box.
//     If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//     the contents of the page are translated to (-Llx, -Lly). After
//     normalization, the media box is updated (Llx and Lly are set to 0 and
//     Urx and Ury are updated accordingly).
//   - Normalize the crop box.
//     The crop box of the page is updated based on the previous operations.
//
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage(page *_de.PdfPage) error {
	_eb, _g := page.GetMediaBox()
	if _g != nil {
		return _g
	}
	_bb, _g := page.GetRotate()
	if _g != nil {
		_d.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _g.Error())
	}
	_dg := _bb%360 != 0 && _bb%90 == 0
	_eb.Normalize()
	_ce, _bd, _cd, _cee := _eb.Llx, _eb.Lly, _eb.Width(), _eb.Height()
	_ba := _ce != 0 || _bd != 0
	if !_dg && !_ba {
		return nil
	}
	_a := func(_ec, _bf, _dgc float64) _e.BoundingBox {
		return _e.Path{Points: []_e.Point{_e.NewPoint(0, 0).Rotate(_dgc), _e.NewPoint(_ec, 0).Rotate(_dgc), _e.NewPoint(0, _bf).Rotate(_dgc), _e.NewPoint(_ec, _bf).Rotate(_dgc)}}.GetBoundingBox()
	}
	_f := _b.NewContentCreator()
	var _ece float64
	if _dg {
		_ece = -float64(_bb)
		_da := _a(_cd, _cee, _ece)
		_f.Translate((_da.Width-_cd)/2+_cd/2, (_da.Height-_cee)/2+_cee/2)
		_f.RotateDeg(_ece)
		_f.Translate(-_cd/2, -_cee/2)
		_cd, _cee = _da.Width, _da.Height
	}
	if _ba {
		_f.Translate(-_ce, -_bd)
	}
	_gd := _f.Operations()
	_cf, _g := _dc.MakeStream(_gd.Bytes(), _dc.NewFlateEncoder())
	if _g != nil {
		return _g
	}
	_cde := _dc.MakeArray(_cf)
	_cde.Append(page.GetContentStreamObjs()...)
	*_eb = _de.PdfRectangle{Urx: _cd, Ury: _cee}
	if _fa := page.CropBox; _fa != nil {
		_fa.Normalize()
		_ef, _fc, _bg, _dac := _fa.Llx-_ce, _fa.Lly-_bd, _fa.Width(), _fa.Height()
		if _dg {
			_df := _a(_bg, _dac, _ece)
			_bg, _dac = _df.Width, _df.Height
		}
		*_fa = _de.PdfRectangle{Llx: _ef, Lly: _fc, Urx: _ef + _bg, Ury: _fc + _dac}
	}
	_d.Log.Debug("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066", _ece, _gd, _eb)
	page.Contents = _cde
	page.Rotate = nil
	return nil
}
