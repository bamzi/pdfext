package sanitize

import (
	_f "github.com/bamzi/pdfext/common"
	_e "github.com/bamzi/pdfext/core"
)

// New returns a new sanitizer object.
func New(opts SanitizationOpts) *Sanitizer { return &Sanitizer{_d: opts} }

// GetSuspiciousObjects returns a count of each detected suspicious object.
func (_gfd *Sanitizer) GetSuspiciousObjects() map[string]int { return _gfd._ec }
func (_dd *Sanitizer) processObjects(_b []_e.PdfObject) ([]_e.PdfObject, error) {
	_ge := []_e.PdfObject{}
	_gb := _dd._d
	for _, _a := range _b {
		switch _aa := _a.(type) {
		case *_e.PdfIndirectObject:
			_ed, _edf := _e.GetDict(_aa)
			if _edf {
				if _c, _af := _e.GetName(_ed.Get("\u0054\u0079\u0070\u0065")); _af && *_c == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _ddf := _e.GetIndirect(_ed.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _ddf && _gb.OpenAction {
						_ed.Remove("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")
					}
				} else if _fa, _ba := _e.GetName(_ed.Get("\u0053")); _ba {
					switch *_fa {
					case "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074":
						if _gb.JavaScript {
							if _ca, _bd := _e.GetStream(_ed.Get("\u004a\u0053")); _bd {
								_cc := []byte{}
								_bc, _da := _e.MakeStream(_cc, nil)
								if _da == nil {
									*_ca = *_bc
								}
							}
							_f.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0055\u0052\u0049":
						if _gb.URI {
							_f.Log.Debug("\u0055\u0052\u0049\u0020ac\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070\u0070\u0065\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054\u006f":
						if _gb.GoTo {
							_f.Log.Debug("G\u004fT\u004f\u0020\u0061\u0063\u0074\u0069\u006f\u006e \u0073\u006b\u0069\u0070pe\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054o\u0052":
						if _gb.GoToR {
							_f.Log.Debug("R\u0065\u006d\u006f\u0074\u0065\u0020G\u006f\u0054\u004f\u0020\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u006bi\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u004c\u0061\u0075\u006e\u0063\u0068":
						if _gb.Launch {
							_f.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn":
						if _cf, _bg := _e.GetStream(_ed.Get("\u004a\u0053")); _bg {
							_bcf := []byte{}
							_bce, _cfa := _e.MakeStream(_bcf, nil)
							if _cfa == nil {
								*_cf = *_bce
							}
						}
					}
				} else if _gc := _ed.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _gc != nil && _gb.JavaScript {
					continue
				} else if _bceb, _gge := _e.GetName(_ed.Get("\u0054\u0079\u0070\u0065")); _gge && *_bceb == "\u0041\u006e\u006eo\u0074" && _gb.JavaScript {
					if _gbg, _df := _e.GetIndirect(_ed.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _df {
						if _fb, _ac := _e.GetDict(_gbg.PdfObject); _ac {
							if _ecd, _ad := _e.GetDict(_fb.Get("\u0041\u0041")); _ad {
								_ab, _db := _e.GetIndirect(_ecd.Get("\u004b"))
								if _db {
									if _abc, _gf := _e.GetDict(_ab.PdfObject); _gf {
										if _bde, _dg := _e.GetName(_abc.Get("\u0053")); _dg && *_bde == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
											_abc.Clear()
										} else if _ff := _ecd.Get("\u0046"); _ff != nil {
											if _dc, _bdb := _e.GetIndirect(_ff); _bdb {
												if _fe, _bdc := _e.GetDict(_dc.PdfObject); _bdc {
													if _afb, _cfc := _e.GetName(_fe.Get("\u0053")); _cfc && *_afb == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
														_fe.Clear()
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case *_e.PdfObjectStream:
			_f.Log.Debug("\u0070d\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073t\u0072e\u0061m\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _aa)
		case *_e.PdfObjectStreams:
			_f.Log.Debug("\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0074\u0079\u0070e\u0020\u0025\u0054", _aa)
		default:
			_f.Log.Debug("u\u006e\u006b\u006e\u006fwn\u0020p\u0064\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0025\u0054", _aa)
		}
		_ge = append(_ge, _a)
	}
	_dd.analyze(_ge)
	return _ge, nil
}

// SanitizationOpts specifies the objects to be removed during sanitization.
type SanitizationOpts struct {

	// JavaScript specifies wether JavaScript action should be removed. JavaScript Actions, section 12.6.4.16 of PDF32000_2008
	JavaScript bool

	// URI specifies if URI actions should be removed. 12.6.4.7 URI Actions, PDF32000_2008.
	URI bool

	// GoToR removes remote GoTo actions. 12.6.4.3 Remote Go-To Actions, PDF32000_2008.
	GoToR bool

	// GoTo specifies wether GoTo actions should be removed. 12.6.4.2 Go-To Actions, PDF32000_2008.
	GoTo bool

	// RenditionJS enables removing of `JS` entry from a Rendition Action.
	// The `JS` entry has a value of text string or stream containing a JavaScript script that shall be executed when the action is triggered.
	// 12.6.4.13 Rendition Actions Table 214, PDF32000_2008.
	RenditionJS bool

	// OpenAction removes OpenAction entry from the document catalog.
	OpenAction bool

	// Launch specifies wether Launch Action should be removed.
	// A launch action launches an application or opens or prints a document.
	// 12.6.4.5 Launch Actions, PDF32000_2008.
	Launch bool
}

// Sanitizer represents a sanitizer object.
// It implements the Optimizer interface to access the objects field from the writer.
type Sanitizer struct {
	_d  SanitizationOpts
	_ec map[string]int
}

func (_dgd *Sanitizer) analyze(_dgb []_e.PdfObject) {
	_fba := map[string]int{}
	for _, _daf := range _dgb {
		switch _dbe := _daf.(type) {
		case *_e.PdfIndirectObject:
			_gbgg, _eeb := _e.GetDict(_dbe.PdfObject)
			if _eeb {
				if _fbb, _dca := _e.GetName(_gbgg.Get("\u0054\u0079\u0070\u0065")); _dca && *_fbb == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _ae := _e.GetIndirect(_gbgg.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _ae {
						_fba["\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"]++
					}
				} else if _fc, _fd := _e.GetName(_gbgg.Get("\u0053")); _fd {
					_gee := _fc.String()
					if _gee == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" || _gee == "\u0055\u0052\u0049" || _gee == "\u0047\u006f\u0054\u006f" || _gee == "\u0047\u006f\u0054o\u0052" || _gee == "\u004c\u0061\u0075\u006e\u0063\u0068" {
						_fba[_gee]++
					} else if _gee == "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn" {
						if _, _fg := _e.GetStream(_gbgg.Get("\u004a\u0053")); _fg {
							_fba[_gee]++
						}
					}
				} else if _geeg := _gbgg.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _geeg != nil {
					_fba["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
				} else if _ce, _cg := _e.GetIndirect(_gbgg.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _cg {
					if _ea, _dab := _e.GetDict(_ce.PdfObject); _dab {
						if _ccb, _aeb := _e.GetDict(_ea.Get("\u0041\u0041")); _aeb {
							_gd := _ccb.Get("\u004b")
							_aag, _aec := _e.GetIndirect(_gd)
							if _aec {
								if _dff, _abe := _e.GetDict(_aag.PdfObject); _abe {
									if _ddc, _cgb := _e.GetName(_dff.Get("\u0053")); _cgb && *_ddc == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
										_fba["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else if _, _eg := _e.GetString(_dff.Get("\u004a\u0053")); _eg {
										_fba["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else {
										_faa := _ccb.Get("\u0046")
										if _faa != nil {
											_bgd, _de := _e.GetIndirect(_faa)
											if _de {
												if _bcd, _feb := _e.GetDict(_bgd.PdfObject); _feb {
													if _geegd, _ag := _e.GetName(_bcd.Get("\u0053")); _ag {
														_fdd := _geegd.String()
														_fba[_fdd]++
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	_dgd._ec = _fba
}

// Optimize optimizes `objects` and returns updated list of objects.
func (_ee *Sanitizer) Optimize(objects []_e.PdfObject) ([]_e.PdfObject, error) {
	return _ee.processObjects(objects)
}
