package docutil

import (
	_d "errors"
	_a "fmt"

	_e "github.com/bamzi/pdfext/common"
	_ab "github.com/bamzi/pdfext/core"
)

type Document struct {
	ID             [2]string
	Version        _ab.Version
	Objects        []_ab.PdfObject
	Info           _ab.PdfObject
	Crypt          *_ab.PdfCrypt
	UseHashBasedID bool
}

func (_fgd Page) FindXObjectForms() []*_ab.PdfObjectStream {
	_gcd, _bad := _fgd.GetResourcesXObject()
	if !_bad {
		return nil
	}
	_bfcc := map[*_ab.PdfObjectStream]struct{}{}
	var _gcb func(_bdec *_ab.PdfObjectDictionary, _gadce map[*_ab.PdfObjectStream]struct{})
	_gcb = func(_acb *_ab.PdfObjectDictionary, _cfc map[*_ab.PdfObjectStream]struct{}) {
		for _, _geg := range _acb.Keys() {
			_ggf, _ece := _ab.GetStream(_acb.Get(_geg))
			if !_ece {
				continue
			}
			if _, _gca := _cfc[_ggf]; _gca {
				continue
			}
			_ada, _egaf := _ab.GetName(_ggf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
			if !_egaf || _ada.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_cfc[_ggf] = struct{}{}
			_edc, _egaf := _ab.GetDict(_ggf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_egaf {
				continue
			}
			_egda, _ccab := _ab.GetDict(_edc.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
			if _ccab {
				_gcb(_egda, _cfc)
			}
		}
	}
	_gcb(_gcd, _bfcc)
	var _ffb []*_ab.PdfObjectStream
	for _bffb := range _bfcc {
		_ffb = append(_ffb, _bffb)
	}
	return _ffb
}
func (_cc *Catalog) HasMetadata() bool {
	_dgf := _cc.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _dgf != nil
}

type OutputIntent struct{ Object *_ab.PdfObjectDictionary }

func (_gfe Content) GetData() ([]byte, error) {
	_ccag, _dcf := _ab.NewEncoderFromStream(_gfe.Stream)
	if _dcf != nil {
		return nil, _dcf
	}
	_dcb, _dcf := _ccag.DecodeStream(_gfe.Stream)
	if _dcf != nil {
		return nil, _dcf
	}
	return _dcb, nil
}
func (_ac *Catalog) GetMarkInfo() (*_ab.PdfObjectDictionary, bool) {
	_gf, _ecb := _ab.GetDict(_ac.Object.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"))
	return _gf, _ecb
}
func (_gac *OutputIntents) Add(oi _ab.PdfObject) error {
	_bcg, _fdf := oi.(*_ab.PdfObjectDictionary)
	if !_fdf {
		return _d.New("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _df, _eeb := _ab.GetStream(_bcg.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _eeb {
		_gac._aed.Objects = append(_gac._aed.Objects, _df)
	}
	_fg, _ba := oi.(*_ab.PdfIndirectObject)
	if !_ba {
		_fg = _ab.MakeIndirectObject(oi)
	}
	if _gac._ae == nil {
		_gac._ae = _ab.MakeArray(_fg)
	} else {
		_gac._ae.Append(_fg)
	}
	_gac._aed.Objects = append(_gac._aed.Objects, _fg)
	return nil
}
func (_acdg Page) GetContents() ([]Content, bool) {
	_faf, _cdf := _ab.GetArray(_acdg.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_cdf {
		_ega, _dff := _ab.GetStream(_acdg.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if !_dff {
			return nil, false
		}
		return []Content{{Stream: _ega, _ade: _acdg, _caf: 0}}, true
	}
	_faa := make([]Content, _faf.Len())
	for _efa, _bga := range _faf.Elements() {
		_cdfe, _ge := _ab.GetStream(_bga)
		if !_ge {
			continue
		}
		_faa[_efa] = Content{Stream: _cdfe, _ade: _acdg, _caf: _efa}
	}
	return _faa, true
}
func (_bfd *Document) AddIndirectObject(indirect *_ab.PdfIndirectObject) {
	for _, _cdg := range _bfd.Objects {
		if _cdg == indirect {
			return
		}
	}
	_bfd.Objects = append(_bfd.Objects, indirect)
}
func (_af *Catalog) SetMetadata(data []byte) error {
	_ec, _da := _ab.MakeStream(data, nil)
	if _da != nil {
		return _da
	}
	_ec.Set("\u0054\u0079\u0070\u0065", _ab.MakeName("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	_ec.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ab.MakeName("\u0058\u004d\u004c"))
	_af.Object.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _ec)
	_af._f.Objects = append(_af._f.Objects, _ec)
	return nil
}
func (_gc Page) GetResourcesXObject() (*_ab.PdfObjectDictionary, bool) {
	_adf, _bba := _gc.GetResources()
	if !_bba {
		return nil, false
	}
	return _ab.GetDict(_adf.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
}
func (_face *Content) SetData(data []byte) error {
	_bcgf, _bbgf := _ab.MakeStream(data, _ab.NewFlateEncoder())
	if _bbgf != nil {
		return _bbgf
	}
	_gegg, _gee := _ab.GetArray(_face._ade.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_gee && _face._caf == 0 {
		_face._ade.Object.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _bcgf)
	} else {
		if _bbgf = _gegg.Set(_face._caf, _bcgf); _bbgf != nil {
			return _bbgf
		}
	}
	_face._ade._dgb.Objects = append(_face._ade._dgb.Objects, _bcgf)
	return nil
}
func (_be *Catalog) GetPages() ([]Page, bool) {
	_bg, _aa := _ab.GetDict(_be.Object.Get("\u0050\u0061\u0067e\u0073"))
	if !_aa {
		return nil, false
	}
	_dg, _g := _ab.GetArray(_bg.Get("\u004b\u0069\u0064\u0073"))
	if !_g {
		return nil, false
	}
	_fc := make([]Page, _dg.Len())
	for _ee, _fe := range _dg.Elements() {
		_ed, _ef := _ab.GetDict(_fe)
		if !_ef {
			continue
		}
		_fc[_ee] = Page{Object: _ed, _cce: _ee + 1, _dgb: _be._f}
	}
	return _fc, true
}

type Catalog struct {
	Object *_ab.PdfObjectDictionary
	_f     *Document
}

func (_ce *OutputIntents) Get(i int) (OutputIntent, bool) {
	if _ce._ae == nil {
		return OutputIntent{}, false
	}
	if i >= _ce._ae.Len() {
		return OutputIntent{}, false
	}
	_bdd := _ce._ae.Get(i)
	_bb, _gfg := _ab.GetIndirect(_bdd)
	if !_gfg {
		_acd, _gad := _ab.GetDict(_bdd)
		return OutputIntent{Object: _acd}, _gad
	}
	_bddc, _dab := _ab.GetDict(_bb.PdfObject)
	return OutputIntent{Object: _bddc}, _dab
}
func (_bde *OutputIntents) Len() int { return _bde._ae.Len() }
func (_bf *Document) FindCatalog() (*Catalog, bool) {
	var _ag *_ab.PdfObjectDictionary
	for _, _fa := range _bf.Objects {
		_ffg, _gb := _ab.GetDict(_fa)
		if !_gb {
			continue
		}
		if _fcf, _bbe := _ab.GetName(_ffg.Get("\u0054\u0079\u0070\u0065")); _bbe && *_fcf == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
			_ag = _ffg
			break
		}
	}
	if _ag == nil {
		return nil, false
	}
	return &Catalog{Object: _ag, _f: _bf}, true
}
func (_gbf *Page) Number() int { return _gbf._cce }
func (_aeg Page) FindXObjectImages() ([]*Image, error) {
	_dfc, _eff := _aeg.GetResourcesXObject()
	if !_eff {
		return nil, nil
	}
	var _fge []*Image
	var _egd error
	_bbb := map[*_ab.PdfObjectStream]int{}
	_fce := map[*_ab.PdfObjectStream]struct{}{}
	var _egae int
	for _, _bea := range _dfc.Keys() {
		_fac, _gfce := _ab.GetStream(_dfc.Get(_bea))
		if !_gfce {
			continue
		}
		if _, _fgb := _bbb[_fac]; _fgb {
			continue
		}
		_befc, _cb := _ab.GetName(_fac.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_cb || _befc.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_eb := Image{BitsPerComponent: 8, Stream: _fac, Name: string(_bea)}
		if _eb.Colorspace, _egd = _fdb(_fac.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _egd != nil {
			_e.Log.Error("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073", _egd)
			continue
		}
		if _ddd, _agb := _ab.GetIntVal(_fac.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _agb {
			_eb.BitsPerComponent = _ddd
		}
		if _bbg, _afcf := _ab.GetIntVal(_fac.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _afcf {
			_eb.Width = _bbg
		}
		if _gcc, _dgg := _ab.GetIntVal(_fac.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _dgg {
			_eb.Height = _gcc
		}
		if _cga, _daf := _ab.GetStream(_fac.Get("\u0053\u004d\u0061s\u006b")); _daf {
			_eb.SMask = &ImageSMask{Image: &_eb, Stream: _cga}
			_fce[_cga] = struct{}{}
		}
		switch _eb.Colorspace {
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_eb.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_eb.ColorComponents = 1
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_eb.ColorComponents = 4
		default:
			_eb.ColorComponents = -1
		}
		_bbb[_fac] = _egae
		_fge = append(_fge, &_eb)
		_egae++
	}
	var _gbc []int
	for _, _bfc := range _fge {
		if _bfc.SMask != nil {
			_gadc, _ede := _bbb[_bfc.SMask.Stream]
			if _ede {
				_gbc = append(_gbc, _gadc)
			}
		}
	}
	_cca := make([]*Image, len(_fge)-len(_gbc))
	_egae = 0
_eab:
	for _ggb, _gcf := range _fge {
		for _, _gaa := range _gbc {
			if _ggb == _gaa {
				continue _eab
			}
		}
		_cca[_egae] = _gcf
		_egae++
	}
	return _fge, nil
}
func (_gbb *Document) GetPages() ([]Page, bool) {
	_bff, _aag := _gbb.FindCatalog()
	if !_aag {
		return nil, false
	}
	return _bff.GetPages()
}
func (_cd *Catalog) GetMetadata() (*_ab.PdfObjectStream, bool) {
	return _ab.GetStream(_cd.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
}

type Image struct {
	Name             string
	Width            int
	Height           int
	Colorspace       _ab.PdfObjectName
	ColorComponents  int
	BitsPerComponent int
	SMask            *ImageSMask
	Stream           *_ab.PdfObjectStream
}
type Page struct {
	_cce   int
	Object *_ab.PdfObjectDictionary
	_dgb   *Document
}

func (_cf *Catalog) SetOutputIntents(outputIntents *OutputIntents) {
	if _dc := _cf.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"); _dc != nil {
		for _ga, _ea := range _cf._f.Objects {
			if _ea == _dc {
				if outputIntents._eac == _dc {
					return
				}
				_cf._f.Objects = append(_cf._f.Objects[:_ga], _cf._f.Objects[_ga+1:]...)
				break
			}
		}
	}
	_ad := outputIntents._eac
	if _ad == nil {
		_ad = _ab.MakeIndirectObject(outputIntents._ae)
	}
	_cf.Object.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _ad)
	_cf._f.Objects = append(_cf._f.Objects, _ad)
}
func (_fdc *Catalog) NewOutputIntents() *OutputIntents { return &OutputIntents{_aed: _fdc._f} }

type Content struct {
	Stream *_ab.PdfObjectStream
	_caf   int
	_ade   Page
}

func (_afc *Catalog) SetStructTreeRoot(structTreeRoot _ab.PdfObject) {
	if structTreeRoot == nil {
		_afc.Object.Remove("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
		return
	}
	_afd := _ab.MakeIndirectObject(structTreeRoot)
	_afc.Object.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", _afd)
	_afc._f.Objects = append(_afc._f.Objects, _afd)
}
func (_ffd *Document) AddStream(stream *_ab.PdfObjectStream) {
	for _, _bdf := range _ffd.Objects {
		if _bdf == stream {
			return
		}
	}
	_ffd.Objects = append(_ffd.Objects, stream)
}
func (_b *Catalog) SetVersion() {
	_b.Object.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _ab.MakeName(_a.Sprintf("\u0025\u0064\u002e%\u0064", _b._f.Version.Major, _b._f.Version.Minor)))
}

type OutputIntents struct {
	_ae  *_ab.PdfObjectArray
	_aed *Document
	_eac *_ab.PdfIndirectObject
}

func (_fd *Catalog) GetStructTreeRoot() (*_ab.PdfObjectDictionary, bool) {
	return _ab.GetDict(_fd.Object.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
}
func (_bef *Catalog) GetOutputIntents() (*OutputIntents, bool) {
	_beb := _bef.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073")
	if _beb == nil {
		return nil, false
	}
	_bc, _bd := _ab.GetIndirect(_beb)
	if !_bd {
		return nil, false
	}
	_ff, _gd := _ab.GetArray(_bc.PdfObject)
	if !_gd {
		return nil, false
	}
	return &OutputIntents{_eac: _bc, _ae: _ff, _aed: _bef._f}, true
}
func _fdb(_gfc _ab.PdfObject) (_ab.PdfObjectName, error) {
	var _add *_ab.PdfObjectName
	var _ccd *_ab.PdfObjectArray
	if _gg, _dd := _gfc.(*_ab.PdfIndirectObject); _dd {
		if _fff, _fee := _gg.PdfObject.(*_ab.PdfObjectArray); _fee {
			_ccd = _fff
		} else if _adb, _eef := _gg.PdfObject.(*_ab.PdfObjectName); _eef {
			_add = _adb
		}
	} else if _dda, _cfb := _gfc.(*_ab.PdfObjectArray); _cfb {
		_ccd = _dda
	} else if _eg, _abe := _gfc.(*_ab.PdfObjectName); _abe {
		_add = _eg
	}
	if _add != nil {
		switch *_add {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_add, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_add, nil
		}
	}
	if _ccd != nil && _ccd.Len() > 0 {
		if _efgg, _ca := _ccd.Get(0).(*_ab.PdfObjectName); _ca {
			switch *_efgg {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ccd.Len() == 1 {
					return *_efgg, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_efgg, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_efgg, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_efgg, nil
			}
		}
	}
	return "", nil
}

type ImageSMask struct {
	Image  *Image
	Stream *_ab.PdfObjectStream
}

func (_cg Page) GetResources() (*_ab.PdfObjectDictionary, bool) {
	return _ab.GetDict(_cg.Object.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
}
func (_efg *Catalog) SetMarkInfo(mi _ab.PdfObject) {
	if mi == nil {
		_efg.Object.Remove("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
		return
	}
	_dgd := _ab.MakeIndirectObject(mi)
	_efg.Object.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", _dgd)
	_efg._f.Objects = append(_efg._f.Objects, _dgd)
}
