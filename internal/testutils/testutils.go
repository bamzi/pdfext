package testutils

import (
	_b "crypto/md5"
	_ad "encoding/hex"
	_e "errors"
	_ab "fmt"
	_cg "image"
	_cb "image/png"
	_cee "io"
	_gf "os"
	_a "os/exec"
	_ce "path/filepath"
	_c "strings"
	_gg "testing"

	_cf "github.com/bamzi/pdfext/common"
	_ca "github.com/bamzi/pdfext/core"
)

func ParseIndirectObjects(rawpdf string) (map[int64]_ca.PdfObject, error) {
	_gd := _ca.NewParserFromString(rawpdf)
	_abf := map[int64]_ca.PdfObject{}
	for {
		_gge, _ecf := _gd.ParseIndirectObject()
		if _ecf != nil {
			if _ecf == _cee.EOF {
				break
			}
			return nil, _ecf
		}
		switch _ee := _gge.(type) {
		case *_ca.PdfIndirectObject:
			_abf[_ee.ObjectNumber] = _gge
		case *_ca.PdfObjectStream:
			_abf[_ee.ObjectNumber] = _gge
		}
	}
	for _, _dda := range _abf {
		_ege(_dda, _abf)
	}
	return _abf, nil
}
func CopyFile(src, dst string) error {
	_ec, _be := _gf.Open(src)
	if _be != nil {
		return _be
	}
	defer _ec.Close()
	_bb, _be := _gf.Create(dst)
	if _be != nil {
		return _be
	}
	defer _bb.Close()
	_, _be = _cee.Copy(_bb, _ec)
	return _be
}
func _ege(_bf _ca.PdfObject, _ace map[int64]_ca.PdfObject) error {
	switch _bfa := _bf.(type) {
	case *_ca.PdfIndirectObject:
		_fbe := _bfa
		_ege(_fbe.PdfObject, _ace)
	case *_ca.PdfObjectDictionary:
		_df := _bfa
		for _, _gb := range _df.Keys() {
			_ecfe := _df.Get(_gb)
			if _gc, _efc := _ecfe.(*_ca.PdfObjectReference); _efc {
				_aaa, _aae := _ace[_gc.ObjectNumber]
				if !_aae {
					return _e.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_df.Set(_gb, _aaa)
			} else {
				_ege(_ecfe, _ace)
			}
		}
	case *_ca.PdfObjectArray:
		_ba := _bfa
		for _abfe, _eb := range _ba.Elements() {
			if _fe, _fea := _eb.(*_ca.PdfObjectReference); _fea {
				_fed, _ade := _ace[_fe.ObjectNumber]
				if !_ade {
					return _e.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_ba.Set(_abfe, _fed)
			} else {
				_ege(_eb, _ace)
			}
		}
	}
	return nil
}
func CompareDictionariesDeep(d1, d2 *_ca.PdfObjectDictionary) bool {
	if len(d1.Keys()) != len(d2.Keys()) {
		_cf.Log.Debug("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", len(d1.Keys()), len(d2.Keys()))
		_cf.Log.Debug("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027", d1.WriteString(), d2.WriteString())
		return false
	}
	for _, _dg := range d1.Keys() {
		if _dg == "\u0050\u0061\u0072\u0065\u006e\u0074" {
			continue
		}
		_ag := _ca.TraceToDirectObject(d1.Get(_dg))
		_ea := _ca.TraceToDirectObject(d2.Get(_dg))
		if _ag == nil {
			_cf.Log.Debug("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		if _ea == nil {
			_cf.Log.Debug("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		switch _af := _ag.(type) {
		case *_ca.PdfObjectDictionary:
			_cca, _fba := _ea.(*_ca.PdfObjectDictionary)
			if !_fba {
				_cf.Log.Debug("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054", _ag, _ea)
				return false
			}
			if !CompareDictionariesDeep(_af, _cca) {
				return false
			}
			continue
		case *_ca.PdfObjectArray:
			_fdb, _dfa := _ea.(*_ca.PdfObjectArray)
			if !_dfa {
				_cf.Log.Debug("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				return false
			}
			if _af.Len() != _fdb.Len() {
				_cf.Log.Debug("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _af.Len(), _fdb.Len())
				return false
			}
			for _eef := 0; _eef < _af.Len(); _eef++ {
				_gga := _ca.TraceToDirectObject(_af.Get(_eef))
				_gae := _ca.TraceToDirectObject(_fdb.Get(_eef))
				if _becg, _gbd := _gga.(*_ca.PdfObjectDictionary); _gbd {
					_gbg, _gde := _gae.(*_ca.PdfObjectDictionary)
					if !_gde {
						return false
					}
					if !CompareDictionariesDeep(_becg, _gbg) {
						return false
					}
				} else {
					if _gga.WriteString() != _gae.WriteString() {
						_cf.Log.Debug("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027", _gga.WriteString(), _gae.WriteString())
						return false
					}
				}
			}
			continue
		}
		if _ag.String() != _ea.String() {
			_cf.Log.Debug("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027", _dg, _ag.String(), _ea.String())
			_cf.Log.Debug("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027", _ag, _ea)
			_cf.Log.Debug("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027", _ag, _ea)
			return false
		}
	}
	return true
}

var (
	ErrRenderNotSupported = _e.New("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m")
)

func RenderPDFToPNGs(pdfPath string, dpi int, outpathTpl string) error {
	if dpi <= 0 {
		dpi = 100
	}
	if _, _bec := _a.LookPath("\u0067\u0073"); _bec != nil {
		return ErrRenderNotSupported
	}
	return _a.Command("\u0067\u0073", "\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061", "\u002d\u006f", outpathTpl, _ab.Sprintf("\u002d\u0072\u0025\u0064", dpi), pdfPath).Run()
}
func HashFile(file string) (string, error) {
	_dc, _db := _gf.Open(file)
	if _db != nil {
		return "", _db
	}
	defer _dc.Close()
	_bg := _b.New()
	if _, _db = _cee.Copy(_bg, _dc); _db != nil {
		return "", _db
	}
	return _ad.EncodeToString(_bg.Sum(nil)), nil
}
func CompareImages(img1, img2 _cg.Image) (bool, error) {
	_ac := img1.Bounds()
	_bc := 0
	for _dd := 0; _dd < _ac.Size().X; _dd++ {
		for _ef := 0; _ef < _ac.Size().Y; _ef++ {
			_ae, _ge, _ff, _ := img1.At(_dd, _ef).RGBA()
			_fc, _fg, _ddd, _ := img2.At(_dd, _ef).RGBA()
			if _ae != _fc || _ge != _fg || _ff != _ddd {
				_bc++
			}
		}
	}
	_dbf := float64(_bc) / float64(_ac.Dx()*_ac.Dy())
	if _dbf > 0.0001 {
		_ab.Printf("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a", _dbf, _bc)
		return false, nil
	}
	return true, nil
}
func ComparePNGFiles(file1, file2 string) (bool, error) {
	_fb, _da := HashFile(file1)
	if _da != nil {
		return false, _da
	}
	_fa, _da := HashFile(file2)
	if _da != nil {
		return false, _da
	}
	if _fb == _fa {
		return true, nil
	}
	_cc, _da := ReadPNG(file1)
	if _da != nil {
		return false, _da
	}
	_cgd, _da := ReadPNG(file2)
	if _da != nil {
		return false, _da
	}
	if _cc.Bounds() != _cgd.Bounds() {
		return false, nil
	}
	return CompareImages(_cc, _cgd)
}
func ReadPNG(file string) (_cg.Image, error) {
	_cab, _ga := _gf.Open(file)
	if _ga != nil {
		return nil, _ga
	}
	defer _cab.Close()
	return _cb.Decode(_cab)
}
func RunRenderTest(t *_gg.T, pdfPath, outputDir, baselineRenderPath string, saveBaseline bool) {
	_ed := _c.TrimSuffix(_ce.Base(pdfPath), _ce.Ext(pdfPath))
	t.Run("\u0072\u0065\u006e\u0064\u0065\u0072", func(_ggc *_gg.T) {
		_fgg := _ce.Join(outputDir, _ed)
		_efe := _fgg + "\u002d%\u0064\u002e\u0070\u006e\u0067"
		if _eg := RenderPDFToPNGs(pdfPath, 0, _efe); _eg != nil {
			_ggc.Skip(_eg)
		}
		for _bcc := 1; true; _bcc++ {
			_dbc := _ab.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _fgg, _bcc)
			_ced := _ce.Join(baselineRenderPath, _ab.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _ed, _bcc))
			if _, _cgdf := _gf.Stat(_dbc); _cgdf != nil {
				break
			}
			_ggc.Logf("\u0025\u0073", _ced)
			if saveBaseline {
				_ggc.Logf("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073", _dbc, _ced)
				_cfc := CopyFile(_dbc, _ced)
				if _cfc != nil {
					_ggc.Fatalf("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _ced, _cfc)
				}
				continue
			}
			_ggc.Run(_ab.Sprintf("\u0070\u0061\u0067\u0065\u0025\u0064", _bcc), func(_fd *_gg.T) {
				_fd.Logf("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073", _dbc, _ced)
				_aed, _edb := ComparePNGFiles(_dbc, _ced)
				if _gf.IsNotExist(_edb) {
					_fd.Fatal("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
				} else if !_aed {
					_fd.Fatal("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064")
				}
			})
		}
	})
}
