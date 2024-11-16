package jbig2

import (
	_d "sort"

	_c "github.com/bamzi/pdfext/internal/bitwise"
	_ec "github.com/bamzi/pdfext/internal/jbig2/decoder"
	_f "github.com/bamzi/pdfext/internal/jbig2/document"
	_cd "github.com/bamzi/pdfext/internal/jbig2/document/segments"
	_a "github.com/bamzi/pdfext/internal/jbig2/errors"
)

func DecodeBytes(encoded []byte, parameters _ec.Parameters, globals ...Globals) ([]byte, error) {
	var _ca Globals
	if len(globals) > 0 {
		_ca = globals[0]
	}
	_b, _be := _ec.Decode(encoded, parameters, _ca.ToDocumentGlobals())
	if _be != nil {
		return nil, _be
	}
	return _b.DecodeNextPage()
}

type Globals map[int]*_cd.Header

func DecodeGlobals(encoded []byte) (Globals, error) {
	const _g = "\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073"
	_gc := _c.NewReader(encoded)
	_ge, _ad := _f.DecodeDocument(_gc, nil)
	if _ad != nil {
		return nil, _a.Wrap(_ad, _g, "")
	}
	if _ge.GlobalSegments == nil || (_ge.GlobalSegments.Segments == nil) {
		return nil, _a.Error(_g, "\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_aa := Globals{}
	for _, _ff := range _ge.GlobalSegments.Segments {
		_aa[int(_ff.SegmentNumber)] = _ff
	}
	return _aa, nil
}
func (_aae Globals) ToDocumentGlobals() *_f.Globals {
	if _aae == nil {
		return nil
	}
	_ffa := []*_cd.Header{}
	for _, _fc := range _aae {
		_ffa = append(_ffa, _fc)
	}
	_d.Slice(_ffa, func(_bb, _ee int) bool { return _ffa[_bb].SegmentNumber < _ffa[_ee].SegmentNumber })
	return &_f.Globals{Segments: _ffa}
}
