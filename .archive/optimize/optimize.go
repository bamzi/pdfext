package optimize

import (
	_a "bytes"
	_ce "crypto/md5"
	_d "errors"
	_gc "fmt"
	_fb "math"
	_cg "strings"

	_f "github.com/bamzi/pdfext/common"
	_fcg "github.com/bamzi/pdfext/contentstream"
	_dc "github.com/bamzi/pdfext/core"
	_fcf "github.com/bamzi/pdfext/extractor"
	_fd "github.com/bamzi/pdfext/internal/imageutil"
	_fc "github.com/bamzi/pdfext/internal/textencoding"
	_e "github.com/bamzi/pdfext/model"
	_gb "github.com/unidoc/unitype"
	_c "golang.org/x/image/draw"
)

// Optimize optimizes PDF objects to decrease PDF size.
func (_cdfb *CombineDuplicateDirectObjects) Optimize(objects []_dc.PdfObject) (_cac []_dc.PdfObject, _bdec error) {
	_dcab(objects)
	_eaee := make(map[string][]*_dc.PdfObjectDictionary)
	var _eeba func(_geg *_dc.PdfObjectDictionary)
	_eeba = func(_egc *_dc.PdfObjectDictionary) {
		for _, _aab := range _egc.Keys() {
			_cdfc := _egc.Get(_aab)
			if _agdd, _ggea := _cdfc.(*_dc.PdfObjectDictionary); _ggea {
				if _adeb := _agdd.Keys(); len(_adeb) == 0 {
					continue
				}
				_fcgb := _ce.New()
				_fcgb.Write([]byte(_agdd.WriteString()))
				_faa := string(_fcgb.Sum(nil))
				_eaee[_faa] = append(_eaee[_faa], _agdd)
				_eeba(_agdd)
			}
		}
	}
	for _, _bbb := range objects {
		_fag, _fdcf := _bbb.(*_dc.PdfIndirectObject)
		if !_fdcf {
			continue
		}
		if _afgc, _aead := _fag.PdfObject.(*_dc.PdfObjectDictionary); _aead {
			_eeba(_afgc)
		}
	}
	_adbaa := make([]_dc.PdfObject, 0, len(_eaee))
	_bcg := make(map[_dc.PdfObject]_dc.PdfObject)
	for _, _adgc := range _eaee {
		if len(_adgc) < 2 {
			continue
		}
		_fed := _dc.MakeDict()
		_fed.Merge(_adgc[0])
		_gcb := _dc.MakeIndirectObject(_fed)
		_adbaa = append(_adbaa, _gcb)
		for _gfbf := 0; _gfbf < len(_adgc); _gfbf++ {
			_acc := _adgc[_gfbf]
			_bcg[_acc] = _gcb
		}
	}
	_cac = make([]_dc.PdfObject, len(objects))
	copy(_cac, objects)
	_cac = append(_adbaa, _cac...)
	_beg(_cac, _bcg)
	return _cac, nil
}
func _cbdc(_gdca []_dc.PdfObject) objectStructure {
	_dce := objectStructure{}
	_dba := false
	for _, _gefg := range _gdca {
		switch _ggb := _gefg.(type) {
		case *_dc.PdfIndirectObject:
			_agddb, _acdd := _dc.GetDict(_ggb)
			if !_acdd {
				continue
			}
			_ggdd, _acdd := _dc.GetName(_agddb.Get("\u0054\u0079\u0070\u0065"))
			if !_acdd {
				continue
			}
			switch _ggdd.String() {
			case "\u0043a\u0074\u0061\u006c\u006f\u0067":
				_dce._cfe = _agddb
				_dba = true
			}
		}
		if _dba {
			break
		}
	}
	if !_dba {
		return _dce
	}
	_acaa, _abaf := _dc.GetDict(_dce._cfe.Get("\u0050\u0061\u0067e\u0073"))
	if !_abaf {
		return _dce
	}
	_dce._edfbg = _acaa
	_aabc, _abaf := _dc.GetArray(_acaa.Get("\u004b\u0069\u0064\u0073"))
	if !_abaf {
		return _dce
	}
	for _, _cafa := range _aabc.Elements() {
		_deb, _degc := _dc.GetIndirect(_cafa)
		if !_degc {
			break
		}
		_dce._fbc = append(_dce._fbc, _deb)
	}
	return _dce
}

// GetOptimizers gets the list of optimizers in chain `c`.
func (_gcd *Chain) GetOptimizers() []_e.Optimizer { return _gcd._b }

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{}
type imageInfo struct {
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_dc.PdfObjectStream
	PPI              float64
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ggef *ObjectStreams) Optimize(objects []_dc.PdfObject) (_gbec []_dc.PdfObject, _feff error) {
	_eead := &_dc.PdfObjectStreams{}
	_efeg := make([]_dc.PdfObject, 0, len(objects))
	for _, _gcbd := range objects {
		if _aac, _fbdd := _gcbd.(*_dc.PdfIndirectObject); _fbdd && _aac.GenerationNumber == 0 {
			_eead.Append(_gcbd)
		} else {
			_efeg = append(_efeg, _gcbd)
		}
	}
	if _eead.Len() == 0 {
		return _efeg, nil
	}
	_gbec = make([]_dc.PdfObject, 0, len(_efeg)+_eead.Len()+1)
	if _eead.Len() > 1 {
		_gbec = append(_gbec, _eead)
	}
	_gbec = append(_gbec, _eead.Elements()...)
	_gbec = append(_gbec, _efeg...)
	return _gbec, nil
}
func _bdc(_bag _dc.PdfObject) []content {
	if _bag == nil {
		return nil
	}
	_cce, _gg := _dc.GetArray(_bag)
	if !_gg {
		_f.Log.Debug("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		return nil
	}
	var _gdbac []content
	for _, _efg := range _cce.Elements() {
		_dag, _eeb := _dc.GetDict(_efg)
		if !_eeb {
			_f.Log.Debug("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073")
			continue
		}
		_afc, _eeb := _dc.GetDict(_dag.Get("\u0041\u0050"))
		if !_eeb {
			_f.Log.Debug("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
			continue
		}
		_bcee := _dc.TraceToDirectObject(_afc.Get("\u004e"))
		if _bcee == nil {
			_f.Log.Debug("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		var _cad *_dc.PdfObjectStream
		switch _gfc := _bcee.(type) {
		case *_dc.PdfObjectDictionary:
			_bdd, _fecc := _dc.GetName(_dag.Get("\u0041\u0053"))
			if !_fecc {
				_f.Log.Debug("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
				continue
			}
			_cad, _fecc = _dc.GetStream(_gfc.Get(*_bdd))
			if !_fecc {
				_f.Log.Debug("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg")
				continue
			}
		case *_dc.PdfObjectStream:
			_cad = _gfc
		}
		if _cad == nil {
			_f.Log.Debug("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		_gge, _adgg := _e.NewXObjectFormFromStream(_cad)
		if _adgg != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _adgg)
			continue
		}
		_dfaf, _adgg := _gge.GetContentStream()
		if _adgg != nil {
			_f.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076", _adgg)
			continue
		}
		_gdbac = append(_gdbac, content{_gba: string(_dfaf), _fbb: _gge.Resources})
	}
	return _gdbac
}
func _ga(_ecg *_dc.PdfObjectStream) error {
	_gf, _cef := _dc.DecodeStream(_ecg)
	if _cef != nil {
		return _cef
	}
	_ffc := _fcg.NewContentStreamParser(string(_gf))
	_gca, _cef := _ffc.Parse()
	if _cef != nil {
		return _cef
	}
	_gca = _bf(_gca)
	_de := _gca.Bytes()
	if len(_de) >= len(_gf) {
		return nil
	}
	_fca, _cef := _dc.MakeStream(_gca.Bytes(), _dc.NewFlateEncoder())
	if _cef != nil {
		return _cef
	}
	_ecg.Stream = _fca.Stream
	_ecg.Merge(_fca.PdfObjectDictionary)
	return nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_cdbb *CombineDuplicateStreams) Optimize(objects []_dc.PdfObject) (_agdc []_dc.PdfObject, _cged error) {
	_gfe := make(map[_dc.PdfObject]_dc.PdfObject)
	_dgg := make(map[_dc.PdfObject]struct{})
	_cagd := make(map[string][]*_dc.PdfObjectStream)
	for _, _ccae := range objects {
		if _edfb, _cfg := _ccae.(*_dc.PdfObjectStream); _cfg {
			_fcag := _ce.New()
			_fcag.Write(_edfb.Stream)
			_fcag.Write([]byte(_edfb.PdfObjectDictionary.WriteString()))
			_fde := string(_fcag.Sum(nil))
			_cagd[_fde] = append(_cagd[_fde], _edfb)
		}
	}
	for _, _bff := range _cagd {
		if len(_bff) < 2 {
			continue
		}
		_dcd := _bff[0]
		for _fgf := 1; _fgf < len(_bff); _fgf++ {
			_ebc := _bff[_fgf]
			_gfe[_ebc] = _dcd
			_dgg[_ebc] = struct{}{}
		}
	}
	_agdc = make([]_dc.PdfObject, 0, len(objects)-len(_dgg))
	for _, _dddd := range objects {
		if _, _eebf := _dgg[_dddd]; _eebf {
			continue
		}
		_agdc = append(_agdc, _dddd)
	}
	_beg(_agdc, _gfe)
	return _agdc, nil
}

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ba *Chain) Optimize(objects []_dc.PdfObject) (_fg []_dc.PdfObject, _ag error) {
	_gd := objects
	for _, _dd := range _ba._b {
		_fcge, _ad := _dd.Optimize(_gd)
		if _ad != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076", _ad)
			continue
		}
		_gd = _fcge
	}
	return _gd, nil
}
func _cga(_gcfa _dc.PdfObject, _ffg map[_dc.PdfObject]struct{}) error {
	if _dgf, _bcdd := _gcfa.(*_dc.PdfIndirectObject); _bcdd {
		_ffg[_gcfa] = struct{}{}
		_gfaf := _cga(_dgf.PdfObject, _ffg)
		if _gfaf != nil {
			return _gfaf
		}
		return nil
	}
	if _bgc, _bee := _gcfa.(*_dc.PdfObjectStream); _bee {
		_ffg[_bgc] = struct{}{}
		_cgcc := _cga(_bgc.PdfObjectDictionary, _ffg)
		if _cgcc != nil {
			return _cgcc
		}
		return nil
	}
	if _aad, _bafd := _gcfa.(*_dc.PdfObjectDictionary); _bafd {
		for _, _afg := range _aad.Keys() {
			_eaa := _aad.Get(_afg)
			_ = _eaa
			if _acb, _cgaa := _eaa.(*_dc.PdfObjectReference); _cgaa {
				_eaa = _acb.Resolve()
				_aad.Set(_afg, _eaa)
			}
			if _afg != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _aaf := _cga(_eaa, _ffg); _aaf != nil {
					return _aaf
				}
			}
		}
		return nil
	}
	if _ecaf, _cdd := _gcfa.(*_dc.PdfObjectArray); _cdd {
		if _ecaf == nil {
			return _d.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _dae, _aaa := range _ecaf.Elements() {
			if _eg, _add := _aaa.(*_dc.PdfObjectReference); _add {
				_aaa = _eg.Resolve()
				_ecaf.Set(_dae, _aaa)
			}
			if _acd := _cga(_aaa, _ffg); _acd != nil {
				return _acd
			}
		}
		return nil
	}
	return nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_bg *CleanContentstream) Optimize(objects []_dc.PdfObject) (_adf []_dc.PdfObject, _ca error) {
	_fdc := map[*_dc.PdfObjectStream]struct{}{}
	var _gac []*_dc.PdfObjectStream
	_ecd := func(_cae *_dc.PdfObjectStream) {
		if _, _df := _fdc[_cae]; !_df {
			_fdc[_cae] = struct{}{}
			_gac = append(_gac, _cae)
		}
	}
	_da := map[_dc.PdfObject]bool{}
	_daf := map[_dc.PdfObject]bool{}
	for _, _ea := range objects {
		switch _dg := _ea.(type) {
		case *_dc.PdfIndirectObject:
			switch _gae := _dg.PdfObject.(type) {
			case *_dc.PdfObjectDictionary:
				if _fcb, _cc := _dc.GetName(_gae.Get("\u0054\u0079\u0070\u0065")); !_cc || _fcb.String() != "\u0050\u0061\u0067\u0065" {
					continue
				}
				if _ee, _eb := _dc.GetStream(_gae.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _eb {
					_ecd(_ee)
				} else if _ddb, _bb := _dc.GetArray(_gae.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _bb {
					var _gaee []*_dc.PdfObjectStream
					for _, _gcf := range _ddb.Elements() {
						if _ed, _fcga := _dc.GetStream(_gcf); _fcga {
							_gaee = append(_gaee, _ed)
						}
					}
					if len(_gaee) > 0 {
						var _ffd _a.Buffer
						for _, _eda := range _gaee {
							if _caa, _bce := _dc.DecodeStream(_eda); _bce == nil {
								_ffd.Write(_caa)
							}
							_da[_eda] = true
						}
						_caf, _gdb := _dc.MakeStream(_ffd.Bytes(), _dc.NewFlateEncoder())
						if _gdb != nil {
							return nil, _gdb
						}
						_daf[_caf] = true
						_gae.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _caf)
						_ecd(_caf)
					}
				}
			}
		case *_dc.PdfObjectStream:
			if _ae, _fcba := _dc.GetName(_dg.Get("\u0054\u0079\u0070\u0065")); !_fcba || _ae.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
				continue
			}
			if _gbc, _bcb := _dc.GetName(_dg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); !_bcb || _gbc.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_ecd(_dg)
		}
	}
	for _, _cgc := range _gac {
		_ca = _ga(_cgc)
		if _ca != nil {
			return nil, _ca
		}
	}
	_adf = nil
	for _, _aea := range objects {
		if _da[_aea] {
			continue
		}
		_adf = append(_adf, _aea)
	}
	for _cge := range _daf {
		_adf = append(_adf, _cge)
	}
	return _adf, nil
}

// CleanUnusedResources represents an optimizer used to clean unused resources.
type CleanUnusedResources struct{}
type imageModifications struct {
	Scale    float64
	Encoding _dc.StreamEncoder
}

func _beg(_baa []_dc.PdfObject, _abc map[_dc.PdfObject]_dc.PdfObject) {
	if len(_abc) == 0 {
		return
	}
	for _aee, _cedc := range _baa {
		if _fgb, _abde := _abc[_cedc]; _abde {
			_baa[_aee] = _fgb
			continue
		}
		_abc[_cedc] = _cedc
		switch _cec := _cedc.(type) {
		case *_dc.PdfObjectArray:
			_fdgde := make([]_dc.PdfObject, _cec.Len())
			copy(_fdgde, _cec.Elements())
			_beg(_fdgde, _abc)
			for _gbff, _egg := range _fdgde {
				_cec.Set(_gbff, _egg)
			}
		case *_dc.PdfObjectStreams:
			_beg(_cec.Elements(), _abc)
		case *_dc.PdfObjectStream:
			_caaf := []_dc.PdfObject{_cec.PdfObjectDictionary}
			_beg(_caaf, _abc)
			_cec.PdfObjectDictionary = _caaf[0].(*_dc.PdfObjectDictionary)
		case *_dc.PdfObjectDictionary:
			_agfa := _cec.Keys()
			_daad := make([]_dc.PdfObject, len(_agfa))
			for _dbda, _bcde := range _agfa {
				_daad[_dbda] = _cec.Get(_bcde)
			}
			_beg(_daad, _abc)
			for _ccaeb, _dcb := range _agfa {
				_cec.Set(_dcb, _daad[_ccaeb])
			}
		case *_dc.PdfIndirectObject:
			_fbgd := []_dc.PdfObject{_cec.PdfObject}
			_beg(_fbgd, _abc)
			_cec.PdfObject = _fbgd[0]
		}
	}
}
func _geb(_gfb *_dc.PdfObjectDictionary) []string {
	_ebb := []string{}
	for _, _gga := range _gfb.Keys() {
		_ebb = append(_ebb, _gga.String())
	}
	return _ebb
}

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{}

func _bdb(_dgc []_dc.PdfObject) (map[_dc.PdfObject]struct{}, error) {
	_gbe := _cbdc(_dgc)
	_agd := _gbe._fbc
	_agc := make(map[_dc.PdfObject]struct{})
	_ge := _baga(_agd)
	for _, _eca := range _agd {
		_dcf, _cag := _dc.GetDict(_eca.PdfObject)
		if !_cag {
			continue
		}
		_ddd, _cag := _dc.GetDict(_dcf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_cag {
			continue
		}
		_dda := _ge["\u0058O\u0062\u006a\u0065\u0063\u0074"]
		_gfd, _cag := _dc.GetDict(_ddd.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if _cag {
			_gfca := _geb(_gfd)
			for _, _bdda := range _gfca {
				if _afd(_bdda, _dda) {
					continue
				}
				_cfd := *_dc.MakeName(_bdda)
				_gacc := _gfd.Get(_cfd)
				_agc[_gacc] = struct{}{}
				_gfd.Remove(_cfd)
				_bcc := _cga(_gacc, _agc)
				if _bcc != nil {
					_f.Log.Debug("\u0066\u0061\u0069\u006ce\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065r\u0073e\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0025\u0076", _gacc)
				}
			}
		}
		_ceg, _cag := _dc.GetDict(_ddd.Get("\u0046\u006f\u006e\u0074"))
		_efga := _ge["\u0046\u006f\u006e\u0074"]
		if _cag {
			_baf := _geb(_ceg)
			for _, _edab := range _baf {
				if _afd(_edab, _efga) {
					continue
				}
				_gdff := *_dc.MakeName(_edab)
				_bde := _ceg.Get(_gdff)
				_agc[_bde] = struct{}{}
				_ceg.Remove(_gdff)
				_gffa := _cga(_bde, _agc)
				if _gffa != nil {
					_f.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _bde)
				}
			}
		}
		_eff, _cag := _dc.GetDict(_ddd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if _cag {
			_adfe := _geb(_eff)
			_afe := _ge["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]
			for _, _agf := range _adfe {
				if _afd(_agf, _afe) {
					continue
				}
				_eeaa := *_dc.MakeName(_agf)
				_gffag := _eff.Get(_eeaa)
				_agc[_gffag] = struct{}{}
				_eff.Remove(_eeaa)
				_fdg := _cga(_gffag, _agc)
				if _fdg != nil {
					_f.Log.Debug("\u0066\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _gffag)
				}
			}
		}
	}
	return _agc, nil
}
func _baga(_gea []*_dc.PdfIndirectObject) map[string][]string {
	_ccc := map[string][]string{}
	for _, _cfde := range _gea {
		_cgga, _acf := _dc.GetDict(_cfde.PdfObject)
		if !_acf {
			continue
		}
		_fdgd := _cgga.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
		_ddf := _dc.TraceToDirectObject(_fdgd)
		_aba := ""
		if _ggf, _adb := _ddf.(*_dc.PdfObjectArray); _adb {
			var _bfcc []string
			for _, _fgg := range _ggf.Elements() {
				_fbg, _ccg := _adda(_fgg)
				if _ccg != nil {
					continue
				}
				_bfcc = append(_bfcc, _fbg)
			}
			_aba = _cg.Join(_bfcc, "\u0020")
		}
		if _abe, _fff := _ddf.(*_dc.PdfObjectStream); _fff {
			_ede, _bfd := _dc.DecodeStream(_abe)
			if _bfd != nil {
				continue
			}
			_aba = string(_ede)
		}
		_agdf := _fcg.NewContentStreamParser(_aba)
		_ecgd, _cbf := _agdf.Parse()
		if _cbf != nil {
			continue
		}
		for _, _fffe := range *_ecgd {
			_beb := _fffe.Operand
			_acfc := _fffe.Params
			switch _beb {
			case "\u0044\u006f":
				_bfde := _acfc[0].String()
				if _, _bdfg := _ccc["\u0058O\u0062\u006a\u0065\u0063\u0074"]; !_bdfg {
					_ccc["\u0058O\u0062\u006a\u0065\u0063\u0074"] = []string{_bfde}
				} else {
					_ccc["\u0058O\u0062\u006a\u0065\u0063\u0074"] = append(_ccc["\u0058O\u0062\u006a\u0065\u0063\u0074"], _bfde)
				}
			case "\u0054\u0066":
				_afeg := _acfc[0].String()
				if _, _gec := _ccc["\u0046\u006f\u006e\u0074"]; !_gec {
					_ccc["\u0046\u006f\u006e\u0074"] = []string{_afeg}
				} else {
					_ccc["\u0046\u006f\u006e\u0074"] = append(_ccc["\u0046\u006f\u006e\u0074"], _afeg)
				}
			case "\u0067\u0073":
				_eebb := _acfc[0].String()
				if _, _fef := _ccc["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]; !_fef {
					_ccc["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = []string{_eebb}
				} else {
					_ccc["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = append(_ccc["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"], _eebb)
				}
			}
		}
	}
	return _ccc
}

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{}

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{ _b []_e.Optimizer }

// Optimize optimizes PDF objects to decrease PDF size.
func (_gfbfd *Image) Optimize(objects []_dc.PdfObject) (_deeg []_dc.PdfObject, _gfcag error) {
	if _gfbfd.ImageQuality <= 0 {
		return objects, nil
	}
	_effa := _ffag(objects)
	if len(_effa) == 0 {
		return objects, nil
	}
	_feb := make(map[_dc.PdfObject]_dc.PdfObject)
	_egd := make(map[_dc.PdfObject]struct{})
	for _, _edbe := range _effa {
		_dfb := _edbe.Stream.Get("\u0053\u004d\u0061s\u006b")
		_egd[_dfb] = struct{}{}
	}
	for _fab, _dbge := range _effa {
		_ddc := _dbge.Stream
		if _, _gad := _egd[_ddc]; _gad {
			continue
		}
		_gdfe, _gfg := _e.NewXObjectImageFromStream(_ddc)
		if _gfg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gfg)
			continue
		}
		switch _gdfe.Filter.(type) {
		case *_dc.JBIG2Encoder:
			continue
		case *_dc.CCITTFaxEncoder:
			continue
		}
		_gfgd, _gfg := _gdfe.ToImage()
		if _gfg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gfg)
			continue
		}
		_gbcf := _dc.NewDCTEncoder()
		_gbcf.ColorComponents = _gfgd.ColorComponents
		_gbcf.Quality = _gfbfd.ImageQuality
		_gbcf.BitsPerComponent = _dbge.BitsPerComponent
		_gbcf.Width = _dbge.Width
		_gbcf.Height = _dbge.Height
		_gcg, _gfg := _gbcf.EncodeBytes(_gfgd.Data)
		if _gfg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gfg)
			continue
		}
		var _adcd _dc.StreamEncoder
		_adcd = _gbcf
		{
			_eeee := _dc.NewFlateEncoder()
			_aef := _dc.NewMultiEncoder()
			_aef.AddEncoder(_eeee)
			_aef.AddEncoder(_gbcf)
			_ecge, _fbdg := _aef.EncodeBytes(_gfgd.Data)
			if _fbdg != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _fbdg)
				continue
			}
			if len(_ecge) < len(_gcg) {
				_f.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029", len(_gcg), len(_ecge), len(_ddc.Stream))
				_gcg = _ecge
				_adcd = _aef
			}
		}
		_abee := len(_ddc.Stream)
		if _abee < len(_gcg) {
			continue
		}
		_cbc := &_dc.PdfObjectStream{Stream: _gcg}
		_cbc.PdfObjectReference = _ddc.PdfObjectReference
		_cbc.PdfObjectDictionary = _dc.MakeDict()
		_cbc.Merge(_ddc.PdfObjectDictionary)
		_cbc.Merge(_adcd.MakeStreamDict())
		_cbc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dc.MakeInteger(int64(len(_gcg))))
		_feb[_ddc] = _cbc
		_effa[_fab].Stream = _cbc
	}
	_deeg = make([]_dc.PdfObject, len(objects))
	copy(_deeg, objects)
	_beg(_deeg, _feb)
	return _deeg, nil
}
func _dcab(_faeb []_dc.PdfObject) {
	for _cace, _ecde := range _faeb {
		switch _dcbc := _ecde.(type) {
		case *_dc.PdfIndirectObject:
			_dcbc.ObjectNumber = int64(_cace + 1)
			_dcbc.GenerationNumber = 0
		case *_dc.PdfObjectStream:
			_dcbc.ObjectNumber = int64(_cace + 1)
			_dcbc.GenerationNumber = 0
		case *_dc.PdfObjectStreams:
			_dcbc.ObjectNumber = int64(_cace + 1)
			_dcbc.GenerationNumber = 0
		}
	}
}

type content struct {
	_gba string
	_fbb *_e.PdfPageResources
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_def *CombineIdenticalIndirectObjects) Optimize(objects []_dc.PdfObject) (_afdg []_dc.PdfObject, _eed error) {
	_dcab(objects)
	_dafd := make(map[_dc.PdfObject]_dc.PdfObject)
	_gaecc := make(map[_dc.PdfObject]struct{})
	_gaag := make(map[string][]*_dc.PdfIndirectObject)
	for _, _dbb := range objects {
		_cda, _agad := _dbb.(*_dc.PdfIndirectObject)
		if !_agad {
			continue
		}
		if _bagc, _bga := _cda.PdfObject.(*_dc.PdfObjectDictionary); _bga {
			if _cgfe, _fdga := _bagc.Get("\u0054\u0079\u0070\u0065").(*_dc.PdfObjectName); _fdga && *_cgfe == "\u0050\u0061\u0067\u0065" {
				continue
			}
			if _ega := _bagc.Keys(); len(_ega) == 0 {
				continue
			}
			_dbg := _ce.New()
			_dbg.Write([]byte(_bagc.WriteString()))
			_cgfg := string(_dbg.Sum(nil))
			_gaag[_cgfg] = append(_gaag[_cgfg], _cda)
		}
	}
	for _, _dad := range _gaag {
		if len(_dad) < 2 {
			continue
		}
		_deab := _dad[0]
		for _faca := 1; _faca < len(_dad); _faca++ {
			_ffa := _dad[_faca]
			_dafd[_ffa] = _deab
			_gaecc[_ffa] = struct{}{}
		}
	}
	_afdg = make([]_dc.PdfObject, 0, len(objects)-len(_gaecc))
	for _, _fafc := range objects {
		if _, _egaa := _gaecc[_fafc]; _egaa {
			continue
		}
		_afdg = append(_afdg, _fafc)
	}
	_beg(_afdg, _dafd)
	return _afdg, nil
}

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{}

func _dea(_bge *_dc.PdfObjectStream, _ccb []rune, _cdfg []_gb.GlyphIndex) error {
	_bge, _bgee := _dc.GetStream(_bge)
	if !_bgee {
		_f.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _d.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_gfac, _bgg := _dc.DecodeStream(_bge)
	if _bgg != nil {
		_f.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bgg)
		return _bgg
	}
	_cgce, _bgg := _gb.Parse(_a.NewReader(_gfac))
	if _bgg != nil {
		_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_bge.Stream))
		return _bgg
	}
	_bcdg := _cdfg
	if len(_ccb) > 0 {
		_ded := _cgce.LookupRunes(_ccb)
		_bcdg = append(_bcdg, _ded...)
	}
	_cgce, _bgg = _cgce.SubsetKeepIndices(_bcdg)
	if _bgg != nil {
		_f.Log.Debug("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v", _bgg)
		return _bgg
	}
	var _dbd _a.Buffer
	_bgg = _cgce.Write(&_dbd)
	if _bgg != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bgg)
		return _bgg
	}
	if _dbd.Len() > len(_gfac) {
		_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
		return nil
	}
	_eab, _bgg := _dc.MakeStream(_dbd.Bytes(), _dc.NewFlateEncoder())
	if _bgg != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bgg)
		return _bgg
	}
	*_bge = *_eab
	_bge.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _dc.MakeInteger(int64(_dbd.Len())))
	return nil
}

// Options describes PDF optimization parameters.
type Options struct {
	CombineDuplicateStreams         bool
	CombineDuplicateDirectObjects   bool
	ImageUpperPPI                   float64
	ImageQuality                    int
	UseObjectStreams                bool
	CombineIdenticalIndirectObjects bool
	CompressStreams                 bool
	CleanFonts                      bool
	SubsetFonts                     bool
	CleanContentstream              bool
	CleanUnusedResources            bool
}

// Optimize implements Optimizer interface.
func (_efgc *CleanUnusedResources) Optimize(objects []_dc.PdfObject) (_deg []_dc.PdfObject, _gff error) {
	_ddbf, _gff := _bdb(objects)
	if _gff != nil {
		return nil, _gff
	}
	_cba := []_dc.PdfObject{}
	for _, _dca := range objects {
		_, _cgg := _ddbf[_dca]
		if _cgg {
			continue
		}
		_cba = append(_cba, _dca)
	}
	return _cba, nil
}

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ ImageUpperPPI float64 }

func _bf(_fa *_fcg.ContentStreamOperations) *_fcg.ContentStreamOperations {
	if _fa == nil {
		return nil
	}
	_ff := _fcg.ContentStreamOperations{}
	for _, _gcc := range *_fa {
		switch _gcc.Operand {
		case "\u0042\u0044\u0043", "\u0042\u004d\u0043", "\u0045\u004d\u0043":
			continue
		case "\u0054\u006d":
			if len(_gcc.Params) == 6 {
				if _cf, _bc := _dc.GetNumbersAsFloat(_gcc.Params); _bc == nil {
					if _cf[0] == 1 && _cf[1] == 0 && _cf[2] == 0 && _cf[3] == 1 {
						_gcc = &_fcg.ContentStreamOperation{Params: []_dc.PdfObject{_gcc.Params[4], _gcc.Params[5]}, Operand: "\u0054\u0064"}
					}
				}
			}
		}
		_ff = append(_ff, _gcc)
	}
	return &_ff
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fcfbb *ImagePPI) Optimize(objects []_dc.PdfObject) (_gcde []_dc.PdfObject, _cbbf error) {
	if _fcfbb.ImageUpperPPI <= 0 {
		return objects, nil
	}
	_cabe := _ffag(objects)
	if len(_cabe) == 0 {
		return objects, nil
	}
	_efbc := make(map[_dc.PdfObject]struct{})
	for _, _bae := range _cabe {
		_eabe := _bae.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")
		_efbc[_eabe] = struct{}{}
	}
	_aaaf := make(map[*_dc.PdfObjectStream]*imageInfo)
	for _, _fbdb := range _cabe {
		_aaaf[_fbdb.Stream] = _fbdb
	}
	var _afgg *_dc.PdfObjectDictionary
	for _, _faeg := range objects {
		if _deabb, _deega := _dc.GetDict(_faeg); _afgg == nil && _deega {
			if _caeb, _gegg := _dc.GetName(_deabb.Get("\u0054\u0079\u0070\u0065")); _gegg && *_caeb == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
				_afgg = _deabb
			}
		}
	}
	if _afgg == nil {
		return objects, nil
	}
	_dadc, _gfdd := _dc.GetDict(_afgg.Get("\u0050\u0061\u0067e\u0073"))
	if !_gfdd {
		return objects, nil
	}
	_ecca, _ecf := _dc.GetArray(_dadc.Get("\u004b\u0069\u0064\u0073"))
	if !_ecf {
		return objects, nil
	}
	for _, _gee := range _ecca.Elements() {
		_agadg := make(map[string]*imageInfo)
		_cead, _cabf := _dc.GetDict(_gee)
		if !_cabf {
			continue
		}
		_gab, _ := _dccfg(_cead.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if len(_gab) == 0 {
			continue
		}
		_degb, _ffgg := _dc.GetDict(_cead.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_ffgg {
			continue
		}
		_ceac, _fad := _e.NewPdfPageResourcesFromDict(_degb)
		if _fad != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076", _fad)
			continue
		}
		_bbbg, _gbde := _dc.GetDict(_degb.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if !_gbde {
			continue
		}
		_fggg := _bbbg.Keys()
		for _, _acca := range _fggg {
			if _ddab, _fecd := _dc.GetStream(_bbbg.Get(_acca)); _fecd {
				if _afb, _dfcb := _aaaf[_ddab]; _dfcb {
					_agadg[string(_acca)] = _afb
				}
			}
		}
		_acdc := _fcg.NewContentStreamParser(_gab)
		_feec, _fad := _acdc.Parse()
		if _fad != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _fad)
			continue
		}
		_cbcc := _fcg.NewContentStreamProcessor(*_feec)
		_cbcc.AddHandler(_fcg.HandlerConditionEnumAllOperands, "", func(_efe *_fcg.ContentStreamOperation, _gcca _fcg.GraphicsState, _bba *_e.PdfPageResources) error {
			switch _efe.Operand {
			case "\u0044\u006f":
				if len(_efe.Params) != 1 {
					_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031")
					return nil
				}
				_aeaa, _agb := _dc.GetName(_efe.Params[0])
				if !_agb {
					_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
					return nil
				}
				if _dcga, _bfa := _agadg[string(*_aeaa)]; _bfa {
					_dege := _gcca.CTM.ScalingFactorX()
					_fdf := _gcca.CTM.ScalingFactorY()
					_bfcf, _agdg := _dege/72.0, _fdf/72.0
					_bbee, _defc := float64(_dcga.Width)/_bfcf, float64(_dcga.Height)/_agdg
					if _bfcf == 0 || _agdg == 0 {
						_bbee = 72.0
						_defc = 72.0
					}
					_dcga.PPI = _fb.Max(_dcga.PPI, _bbee)
					_dcga.PPI = _fb.Max(_dcga.PPI, _defc)
				}
			}
			return nil
		})
		_fad = _cbcc.Process(_ceac)
		if _fad != nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076", _fad)
			continue
		}
	}
	for _, _fbbe := range _cabe {
		if _, _cddb := _efbc[_fbbe.Stream]; _cddb {
			continue
		}
		if _fbbe.PPI <= _fcfbb.ImageUpperPPI {
			continue
		}
		_gdg, _beec := _e.NewXObjectImageFromStream(_fbbe.Stream)
		if _beec != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _beec)
			continue
		}
		var _bef imageModifications
		_bef.Scale = _fcfbb.ImageUpperPPI / _fbbe.PPI
		if _fbbe.BitsPerComponent == 1 && _fbbe.ColorComponents == 1 {
			_fcc := _fb.Round(_fbbe.PPI / _fcfbb.ImageUpperPPI)
			_gfbg := _fd.NextPowerOf2(uint(_fcc))
			if _fd.InDelta(float64(_gfbg), 1/_bef.Scale, 0.3) {
				_bef.Scale = float64(1) / float64(_gfbg)
			}
			if _, _gaed := _gdg.Filter.(*_dc.JBIG2Encoder); !_gaed {
				_bef.Encoding = _dc.NewJBIG2Encoder()
			}
		}
		if _beec = _dgd(_gdg, _bef); _beec != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _beec)
			continue
		}
		_bef.Encoding = nil
		if _dbgg, _eeaf := _dc.GetStream(_fbbe.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")); _eeaf {
			_gabd, _eeceg := _e.NewXObjectImageFromStream(_dbgg)
			if _eeceg != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _eeceg)
				continue
			}
			if _eeceg = _dgd(_gabd, _bef); _eeceg != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _eeceg)
				continue
			}
		}
	}
	return objects, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_dfa *CleanFonts) Optimize(objects []_dc.PdfObject) (_eec []_dc.PdfObject, _cb error) {
	var _aec map[*_dc.PdfObjectStream]struct{}
	if _dfa.Subset {
		var _fdcg error
		_aec, _fdcg = _dcg(objects)
		if _fdcg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076", _fdcg)
			return nil, _fdcg
		}
	}
	for _, _cca := range objects {
		_fcfb, _eece := _dc.GetStream(_cca)
		if !_eece {
			continue
		}
		if _, _cgf := _aec[_fcfb]; _cgf {
			continue
		}
		_eaf, _dfae := _dc.NewEncoderFromStream(_fcfb)
		if _dfae != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067", _dfae)
			continue
		}
		_ecc, _dfae := _eaf.DecodeStream(_fcfb)
		if _dfae != nil {
			_f.Log.Debug("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067", _dfae)
			continue
		}
		if len(_ecc) < 4 {
			continue
		}
		_fec := string(_ecc[:4])
		if _fec == "\u004f\u0054\u0054\u004f" {
			continue
		}
		if _fec != "\u0000\u0001\u0000\u0000" && _fec != "\u0074\u0072\u0075\u0065" {
			continue
		}
		_gdc, _dfae := _gb.Parse(_a.NewReader(_ecc))
		if _dfae != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _dfae)
			continue
		}
		_dfae = _gdc.Optimize()
		if _dfae != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067", _dfae)
			continue
		}
		var _faf _a.Buffer
		_dfae = _gdc.Write(&_faf)
		if _dfae != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _dfae)
			continue
		}
		if _faf.Len() > len(_ecc) {
			_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
			continue
		}
		_gdba, _dfae := _dc.MakeStream(_faf.Bytes(), _dc.NewFlateEncoder())
		if _dfae != nil {
			continue
		}
		*_fcfb = *_gdba
		_fcfb.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _dc.MakeInteger(int64(_faf.Len())))
	}
	return objects, nil
}
func _adda(_cgff _dc.PdfObject) (string, error) {
	_ade := _dc.TraceToDirectObject(_cgff)
	switch _gaa := _ade.(type) {
	case *_dc.PdfObjectString:
		return _gaa.Str(), nil
	case *_dc.PdfObjectStream:
		_ggc, _eebd := _dc.DecodeStream(_gaa)
		if _eebd != nil {
			return "", _eebd
		}
		return string(_ggc), nil
	}
	return "", _gc.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _ade)
}

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{}

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct {

	// Subset embedded fonts if encountered (if true).
	// Otherwise attempts to reduce the font program.
	Subset bool
}

// Append appends optimizers to the chain.
func (_be *Chain) Append(optimizers ..._e.Optimizer) { _be._b = append(_be._b, optimizers...) }
func _dccfg(_aafa _dc.PdfObject) (_bca string, _ddac []_dc.PdfObject) {
	var _dcag _a.Buffer
	switch _beaf := _aafa.(type) {
	case *_dc.PdfIndirectObject:
		_ddac = append(_ddac, _beaf)
		_aafa = _beaf.PdfObject
	}
	switch _cbbb := _aafa.(type) {
	case *_dc.PdfObjectStream:
		if _aae, _eeg := _dc.DecodeStream(_cbbb); _eeg == nil {
			_dcag.Write(_aae)
			_ddac = append(_ddac, _cbbb)
		}
	case *_dc.PdfObjectArray:
		for _, _ffcf := range _cbbb.Elements() {
			switch _gbga := _ffcf.(type) {
			case *_dc.PdfObjectStream:
				if _bec, _bbdd := _dc.DecodeStream(_gbga); _bbdd == nil {
					_dcag.Write(_bec)
					_ddac = append(_ddac, _gbga)
				}
			}
		}
	}
	return _dcag.String(), _ddac
}
func _dgd(_ddca *_e.XObjectImage, _efb imageModifications) error {
	_bagg, _bbfc := _ddca.ToImage()
	if _bbfc != nil {
		return _bbfc
	}
	if _efb.Scale != 0 {
		_bagg, _bbfc = _fdeb(_bagg, _efb.Scale)
		if _bbfc != nil {
			return _bbfc
		}
	}
	if _efb.Encoding != nil {
		_ddca.Filter = _efb.Encoding
	}
	_ddca.Decode = nil
	switch _fea := _ddca.Filter.(type) {
	case *_dc.FlateEncoder:
		if _fea.Predictor != 1 && _fea.Predictor != 11 {
			_fea.Predictor = 1
		}
	}
	if _bbfc = _ddca.SetImage(_bagg, nil); _bbfc != nil {
		_f.Log.Debug("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _bbfc)
		return _bbfc
	}
	_ddca.ToPdfObject()
	return nil
}
func _afd(_fee string, _cafg []string) bool {
	for _, _ggg := range _cafg {
		if _fee == _ggg {
			return true
		}
	}
	return false
}
func _fdeb(_aegg *_e.Image, _ccad float64) (*_e.Image, error) {
	_cafgg, _bgf := _aegg.ToGoImage()
	if _bgf != nil {
		return nil, _bgf
	}
	var _ced _fd.Image
	_eaea, _bcf := _cafgg.(*_fd.Monochrome)
	if _bcf {
		if _bgf = _eaea.ResolveDecode(); _bgf != nil {
			return nil, _bgf
		}
		_ced, _bgf = _eaea.Scale(_ccad)
		if _bgf != nil {
			return nil, _bgf
		}
	} else {
		_ccbf := int(_fb.RoundToEven(float64(_aegg.Width) * _ccad))
		_adfc := int(_fb.RoundToEven(float64(_aegg.Height) * _ccad))
		_ced, _bgf = _fd.NewImage(_ccbf, _adfc, int(_aegg.BitsPerComponent), _aegg.ColorComponents, nil, nil, nil)
		if _bgf != nil {
			return nil, _bgf
		}
		_c.CatmullRom.Scale(_ced, _ced.Bounds(), _cafgg, _cafgg.Bounds(), _c.Over, &_c.Options{})
	}
	_fcbg := _ced.Base()
	_adff := &_e.Image{Width: int64(_fcbg.Width), Height: int64(_fcbg.Height), BitsPerComponent: int64(_fcbg.BitsPerComponent), ColorComponents: _fcbg.ColorComponents, Data: _fcbg.Data}
	_adff.SetDecode(_fcbg.Decode)
	_adff.SetAlpha(_fcbg.Alpha)
	return _adff, nil
}

type objectStructure struct {
	_cfe   *_dc.PdfObjectDictionary
	_edfbg *_dc.PdfObjectDictionary
	_fbc   []*_dc.PdfIndirectObject
}

func _ffag(_egb []_dc.PdfObject) []*imageInfo {
	_fdgad := _dc.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_fcgbb := make(map[*_dc.PdfObjectStream]struct{})
	var _dab []*imageInfo
	for _, _accc := range _egb {
		_eeae, _edad := _dc.GetStream(_accc)
		if !_edad {
			continue
		}
		if _, _feg := _fcgbb[_eeae]; _feg {
			continue
		}
		_fcgbb[_eeae] = struct{}{}
		_aeb := _eeae.PdfObjectDictionary.Get(_fdgad)
		_ddda, _edad := _dc.GetName(_aeb)
		if !_edad || string(*_ddda) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_dbfb := &imageInfo{Stream: _eeae, BitsPerComponent: 8}
		if _dccf, _gef := _dc.GetIntVal(_eeae.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _gef {
			_dbfb.BitsPerComponent = _dccf
		}
		if _cbb, _dfe := _dc.GetIntVal(_eeae.Get("\u0057\u0069\u0064t\u0068")); _dfe {
			_dbfb.Width = _cbb
		}
		if _dfc, _fae := _dc.GetIntVal(_eeae.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _fae {
			_dbfb.Height = _dfc
		}
		_deaa, _fggb := _e.NewPdfColorspaceFromPdfObject(_eeae.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _fggb != nil {
			_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fggb)
			continue
		}
		if _deaa == nil {
			_ffdd, _dbbe := _dc.GetName(_eeae.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
			if _dbbe {
				switch _ffdd.String() {
				case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065", "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":
					_deaa = _e.NewPdfColorspaceDeviceGray()
					_dbfb.BitsPerComponent = 1
				}
			}
		}
		switch _cea := _deaa.(type) {
		case *_e.PdfColorspaceDeviceRGB:
			_dbfb.ColorComponents = 3
		case *_e.PdfColorspaceDeviceGray:
			_dbfb.ColorComponents = 1
		default:
			_f.Log.Debug("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070", _cea)
			continue
		}
		_dab = append(_dab, _dbfb)
	}
	return _dab
}
func _dcg(_db []_dc.PdfObject) (_edg map[*_dc.PdfObjectStream]struct{}, _cfc error) {
	_edg = map[*_dc.PdfObjectStream]struct{}{}
	_aga := map[*_e.PdfFont]struct{}{}
	_ebd := _cbdc(_db)
	for _, _age := range _ebd._fbc {
		_gdf, _caab := _dc.GetDict(_age.PdfObject)
		if !_caab {
			continue
		}
		_edae, _caab := _dc.GetDict(_gdf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_caab {
			continue
		}
		_fac, _ := _dccfg(_gdf.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		_bdf, _eae := _e.NewPdfPageResourcesFromDict(_edae)
		if _eae != nil {
			return nil, _eae
		}
		_gbd := []content{{_gba: _fac, _fbb: _bdf}}
		_bda := _bdc(_gdf.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _bda != nil {
			_gbd = append(_gbd, _bda...)
		}
		for _, _aeg := range _gbd {
			_aeaf, _bcd := _fcf.NewFromContents(_aeg._gba, _aeg._fbb)
			if _bcd != nil {
				return nil, _bcd
			}
			_bfc, _, _, _bcd := _aeaf.ExtractPageText()
			if _bcd != nil {
				return nil, _bcd
			}
			for _, _gbg := range _bfc.Marks().Elements() {
				if _gbg.Font == nil {
					continue
				}
				if _, _gde := _aga[_gbg.Font]; !_gde {
					_aga[_gbg.Font] = struct{}{}
				}
			}
		}
	}
	_cd := map[*_dc.PdfObjectStream][]*_e.PdfFont{}
	for _ef := range _aga {
		_fbd := _ef.FontDescriptor()
		if _fbd == nil || _fbd.FontFile2 == nil {
			continue
		}
		_cdf, _edb := _dc.GetStream(_fbd.FontFile2)
		if !_edb {
			continue
		}
		_cd[_cdf] = append(_cd[_cdf], _ef)
	}
	for _ac := range _cd {
		var _dac []rune
		var _cfae []_gb.GlyphIndex
		for _, _bbd := range _cd[_ac] {
			switch _bab := _bbd.Encoder().(type) {
			case *_fc.IdentityEncoder:
				_adg := _bab.RegisteredRunes()
				_dee := make([]_gb.GlyphIndex, len(_adg))
				for _edf, _ace := range _adg {
					_dee[_edf] = _gb.GlyphIndex(_ace)
				}
				_cfae = append(_cfae, _dee...)
			case *_fc.TrueTypeFontEncoder:
				_bfe := _bab.RegisteredRunes()
				_dac = append(_dac, _bfe...)
			case _fc.SimpleEncoder:
				_gccef := _bab.Charcodes()
				for _, _dcc := range _gccef {
					_af, _cab := _bab.CharcodeToRune(_dcc)
					if !_cab {
						_f.Log.Debug("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064", _dcc)
						continue
					}
					_dac = append(_dac, _af)
				}
			}
		}
		_cfc = _dea(_ac, _dac, _cfae)
		if _cfc != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _cfc)
			return nil, _cfc
		}
		_edg[_ac] = struct{}{}
	}
	return _edg, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fgc *CompressStreams) Optimize(objects []_dc.PdfObject) (_bbf []_dc.PdfObject, _ebe error) {
	_bbf = make([]_dc.PdfObject, len(objects))
	copy(_bbf, objects)
	for _, _gcae := range objects {
		_ebfe, _fgce := _dc.GetStream(_gcae)
		if !_fgce {
			continue
		}
		if _eccc := _ebfe.Get("\u0046\u0069\u006c\u0074\u0065\u0072"); _eccc != nil {
			if _, _gfeg := _dc.GetName(_eccc); _gfeg {
				continue
			}
			if _abf, _dcgc := _dc.GetArray(_eccc); _dcgc && _abf.Len() > 0 {
				continue
			}
		}
		_cabc := _dc.NewFlateEncoder()
		var _gdbb []byte
		_gdbb, _ebe = _cabc.EncodeBytes(_ebfe.Stream)
		if _ebe != nil {
			return _bbf, _ebe
		}
		_eabc := _cabc.MakeStreamDict()
		if len(_gdbb)+len(_eabc.WriteString()) < len(_ebfe.Stream) {
			_ebfe.Stream = _gdbb
			_ebfe.PdfObjectDictionary.Merge(_eabc)
			_ebfe.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dc.MakeInteger(int64(len(_ebfe.Stream))))
		}
	}
	return _bbf, nil
}

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ ImageQuality int }

// New creates a optimizers chain from options.
func New(options Options) *Chain {
	_dge := new(Chain)
	if options.CleanFonts || options.SubsetFonts {
		_dge.Append(&CleanFonts{Subset: options.SubsetFonts})
	}
	if options.CleanContentstream {
		_dge.Append(new(CleanContentstream))
	}
	if options.ImageUpperPPI > 0 {
		_bfed := new(ImagePPI)
		_bfed.ImageUpperPPI = options.ImageUpperPPI
		_dge.Append(_bfed)
	}
	if options.ImageQuality > 0 {
		_aaae := new(Image)
		_aaae.ImageQuality = options.ImageQuality
		_dge.Append(_aaae)
	}
	if options.CombineDuplicateDirectObjects {
		_dge.Append(new(CombineDuplicateDirectObjects))
	}
	if options.CombineDuplicateStreams {
		_dge.Append(new(CombineDuplicateStreams))
	}
	if options.CombineIdenticalIndirectObjects {
		_dge.Append(new(CombineIdenticalIndirectObjects))
	}
	if options.UseObjectStreams {
		_dge.Append(new(ObjectStreams))
	}
	if options.CompressStreams {
		_dge.Append(new(CompressStreams))
	}
	if options.CleanUnusedResources {
		_dge.Append(new(CleanUnusedResources))
	}
	return _dge
}
