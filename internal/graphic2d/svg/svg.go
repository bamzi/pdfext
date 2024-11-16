package svg

import (
	_bg "encoding/xml"
	_g "fmt"
	_eb "io"
	_a "math"
	_e "os"
	_c "strconv"
	_eg "strings"
	_f "unicode"

	_fa "github.com/bamzi/pdfext/common"
	_ad "github.com/bamzi/pdfext/contentstream"
	_af "github.com/bamzi/pdfext/contentstream/draw"
	_gd "github.com/bamzi/pdfext/core"
	_d "github.com/bamzi/pdfext/internal/graphic2d"
	_fb "github.com/bamzi/pdfext/model"
	_bb "golang.org/x/net/html/charset"
)

func _cga(_dafe string) (float64, error) {
	_dafe = _eg.TrimSpace(_dafe)
	var _edeg float64
	if _eg.HasSuffix(_dafe, "\u0025") {
		_fbb, _ffa := _c.ParseFloat(_eg.TrimSuffix(_dafe, "\u0025"), 64)
		if _ffa != nil {
			return 0, _ffa
		}
		_edeg = (_fbb * 255.0) / 100.0
	} else {
		_decd, _egbe := _c.Atoi(_dafe)
		if _egbe != nil {
			return 0, _egbe
		}
		_edeg = float64(_decd)
	}
	return _edeg, nil
}
func _ebg(_ggf *GraphicSVG, _cfg *_ad.ContentCreator, _fca *_fb.PdfPageResources) {
	_cfg.Add_q()
	_ggf.Style.toContentStream(_cfg, _fca)
	_egc, _ddd := _dgef(_ggf.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _ddd != nil {
		_fa.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _ddd)
		return
	}
	if len(_egc)%2 > 0 {
		_fa.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _gdb := 0; _gdb < len(_egc); {
		if _gdb == 0 {
			_cfg.Add_m(_egc[_gdb]*_ggf._ccf, _egc[_gdb+1]*_ggf._ccf)
		} else {
			_cfg.Add_l(_egc[_gdb]*_ggf._ccf, _egc[_gdb+1]*_ggf._ccf)
		}
		_gdb += 2
	}
	_ggf.Style.fillStroke(_cfg)
	_cfg.Add_h()
	_cfg.Add_Q()
}
func _dgef(_aegcf string) ([]float64, error) {
	_aff := -1
	var _dfe []float64
	_fdc := ' '
	for _aedc, _dfd := range _aegcf {
		if !_f.IsNumber(_dfd) && _dfd != '.' && !(_dfd == '-' && _fdc == 'e') && _dfd != 'e' {
			if _aff != -1 {
				_bccc, _adef := _cdbg(_aegcf[_aff:_aedc])
				if _adef != nil {
					return _dfe, _adef
				}
				_dfe = append(_dfe, _bccc...)
			}
			if _dfd == '-' {
				_aff = _aedc
			} else {
				_aff = -1
			}
		} else if _aff == -1 {
			_aff = _aedc
		}
		_fdc = _dfd
	}
	if _aff != -1 && _aff != len(_aegcf) {
		_egg, _fbd := _cdbg(_aegcf[_aff:])
		if _fbd != nil {
			return _dfe, _fbd
		}
		_dfe = append(_dfe, _egg...)
	}
	return _dfe, nil
}
func ParseFromFile(path string) (*GraphicSVG, error) {
	_gcb, _cgg := _e.Open(path)
	if _cgg != nil {
		return nil, _cgg
	}
	defer _gcb.Close()
	return ParseFromStream(_gcb)
}
func (_ccb *GraphicSVGStyle) toContentStream(_gafbg *_ad.ContentCreator, _bdb *_fb.PdfPageResources) {
	if _ccb == nil {
		return
	}
	if _ccb.FillColor != "" {
		var _afe, _beae, _ddg float64
		if _fgc, _egcf := _d.ColorMap[_ccb.FillColor]; _egcf {
			_bdf, _efa, _ddgb, _ := _fgc.RGBA()
			_afe, _beae, _ddg = float64(_bdf), float64(_efa), float64(_ddgb)
		} else if _eg.HasPrefix(_ccb.FillColor, "\u0072\u0067\u0062\u0028") {
			_afe, _beae, _ddg = _dcb(_ccb.FillColor)
		} else {
			_afe, _beae, _ddg = _efe(_ccb.FillColor)
		}
		_gafbg.Add_rg(_afe, _beae, _ddg)
	}
	if _ccb.FillOpacity < 1.0 {
		_bbdb := 0
		_gccbd := _gd.PdfObjectName(_g.Sprintf("\u0047\u0053\u0025\u0064", _bbdb))
		for {
			_, _fda := _bdb.GetExtGState(_gccbd)
			if !_fda {
				break
			}
			_bbdb++
			_gccbd = _gd.PdfObjectName(_g.Sprintf("\u0047\u0053\u0025\u0064", _bbdb))
		}
		_fggb := _gd.MakeDict()
		_fggb.Set("\u0063\u0061", _gd.MakeFloat(_ccb.FillOpacity))
		_bged := _bdb.AddExtGState(_gccbd, _gd.MakeIndirectObject(_fggb))
		if _bged != nil {
			_fa.Log.Debug(_bged.Error())
			return
		}
		_gafbg.Add_gs(_gccbd)
	}
	if _ccb.StrokeColor != "" {
		var _bcg, _daef, _eed float64
		if _gge, _ccfb := _d.ColorMap[_ccb.StrokeColor]; _ccfb {
			_dde, _bgcb, _edcff, _ := _gge.RGBA()
			_bcg, _daef, _eed = float64(_dde)/255.0, float64(_bgcb)/255.0, float64(_edcff)/255.0
		} else if _eg.HasPrefix(_ccb.FillColor, "\u0072\u0067\u0062\u0028") {
			_bcg, _daef, _eed = _dcb(_ccb.FillColor)
		} else {
			_bcg, _daef, _eed = _efe(_ccb.StrokeColor)
		}
		_gafbg.Add_RG(_bcg, _daef, _eed)
	}
	if _ccb.StrokeWidth > 0 {
		_gafbg.Add_w(_ccb.StrokeWidth)
	}
}

type commands struct {
	_adg  []string
	_bfae map[string]int
	_dacb string
	_dec  string
}

func (_acbe *Subpath) compare(_dffe *Subpath) bool {
	if len(_acbe.Commands) != len(_dffe.Commands) {
		return false
	}
	for _ace, _bfg := range _acbe.Commands {
		if !_bfg.compare(_dffe.Commands[_ace]) {
			return false
		}
	}
	return true
}

type pathParserError struct{ _ffb string }

func (_bbb *GraphicSVG) setDefaultScaling(_abg float64) {
	_bbb._ccf = _abg
	if _bbb.Style != nil && _bbb.Style.StrokeWidth > 0 {
		_bbb.Style.StrokeWidth = _bbb.Style.StrokeWidth * _bbb._ccf
	}
	for _, _cafee := range _bbb.Children {
		_cafee.setDefaultScaling(_abg)
	}
}
func _bdd(_ggbaa string) (_edd, _aab string) {
	if _ggbaa == "" || (_ggbaa[len(_ggbaa)-1] >= '0' && _ggbaa[len(_ggbaa)-1] <= '9') {
		return _ggbaa, ""
	}
	_edd = _ggbaa
	for _, _daec := range _fbg {
		if _eg.Contains(_edd, _daec) {
			_aab = _daec
		}
		_edd = _eg.TrimSuffix(_edd, _daec)
	}
	return
}
func _cceb() *GraphicSVGStyle {
	return &GraphicSVGStyle{FillColor: "\u00230\u0030\u0030\u0030\u0030\u0030", StrokeColor: "", StrokeWidth: 0, FillOpacity: 1.0}
}
func (_fbe *GraphicSVG) ToContentCreator(cc *_ad.ContentCreator, res *_fb.PdfPageResources, scaleX, scaleY, translateX, translateY float64) *_ad.ContentCreator {
	if _fbe.Name == "\u0073\u0076\u0067" {
		_fbe.SetScaling(scaleX, scaleY)
		cc.Add_cm(1, 0, 0, 1, translateX, translateY)
		_fbe.setDefaultScaling(_fbe._ccf)
		cc.Add_q()
		_dc := _a.Max(scaleX, scaleY)
		cc.Add_re(_fbe.ViewBox.X*_dc, _fbe.ViewBox.Y*_dc, _fbe.ViewBox.W*_dc, _fbe.ViewBox.H*_dc)
		cc.Add_W()
		cc.Add_n()
		for _, _fd := range _fbe.Children {
			_fd.ViewBox = _fbe.ViewBox
			_fd.toContentStream(cc, res)
		}
		cc.Add_Q()
		return cc
	}
	return nil
}
func (_effe *Path) compare(_dacf *Path) bool {
	if len(_effe.Subpaths) != len(_dacf.Subpaths) {
		return false
	}
	for _bbee, _defd := range _effe.Subpaths {
		if !_defd.compare(_dacf.Subpaths[_bbee]) {
			return false
		}
	}
	return true
}

type GraphicSVG struct {
	ViewBox    struct{ X, Y, W, H float64 }
	Name       string
	Attributes map[string]string
	Children   []*GraphicSVG
	Content    string
	Style      *GraphicSVGStyle
	Width      float64
	Height     float64
	_ccf       float64
}

func (_dbc *Command) isAbsolute() bool { return _dbc.Symbol == _eg.ToUpper(_dbc.Symbol) }
func ParseFromStream(source _eb.Reader) (*GraphicSVG, error) {
	_efb := _bg.NewDecoder(source)
	_efb.CharsetReader = _bb.NewReaderLabel
	_eeg, _fae := _ggbc(_efb)
	if _fae != nil {
		return nil, _fae
	}
	if _bdfc := _eeg.Decode(_efb); _bdfc != nil && _bdfc != _eb.EOF {
		return nil, _bdfc
	}
	return _eeg, nil
}
func _gbf(_ddf map[string]string, _fdae float64) (*GraphicSVGStyle, error) {
	_edg := _cceb()
	_effc, _fef := _ddf["\u0066\u0069\u006c\u006c"]
	if _fef {
		_edg.FillColor = _effc
		if _effc == "\u006e\u006f\u006e\u0065" {
			_edg.FillColor = ""
		}
	}
	_ecc, _eeac := _ddf["\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079"]
	if _eeac {
		_faa, _eda := _ecd(_ecc)
		if _eda != nil {
			return nil, _eda
		}
		_edg.FillOpacity = _faa
	}
	_fbfc, _gcbc := _ddf["\u0073\u0074\u0072\u006f\u006b\u0065"]
	if _gcbc {
		_edg.StrokeColor = _fbfc
		if _fbfc == "\u006e\u006f\u006e\u0065" {
			_edg.StrokeColor = ""
		}
	}
	_ddag, _ebgf := _ddf["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
	if _ebgf {
		_caga, _babg := _fee(_ddag, 64)
		if _babg != nil {
			return nil, _babg
		}
		_edg.StrokeWidth = _caga * _fdae
	}
	return _edg, nil
}
func _efe(_beeb string) (_bdg, _gfga, _gfggb float64) {
	if (len(_beeb) != 4 && len(_beeb) != 7) || _beeb[0] != '#' {
		_fa.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _beeb)
		return _bdg, _gfga, _gfggb
	}
	var _cbf, _bdgc, _gafe int
	if len(_beeb) == 4 {
		var _gdbe, _eggc, _effcf int
		_fgef, _agef := _g.Sscanf(_beeb, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_gdbe, &_eggc, &_effcf)
		if _agef != nil {
			_fa.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _beeb, _agef)
			return _bdg, _gfga, _gfggb
		}
		if _fgef != 3 {
			_fa.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _beeb)
			return _bdg, _gfga, _gfggb
		}
		_cbf = _gdbe*16 + _gdbe
		_bdgc = _eggc*16 + _eggc
		_gafe = _effcf*16 + _effcf
	} else {
		_edag, _gcfef := _g.Sscanf(_beeb, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_cbf, &_bdgc, &_gafe)
		if _gcfef != nil {
			_fa.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _beeb)
			return _bdg, _gfga, _gfggb
		}
		if _edag != 3 {
			_fa.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", _beeb, _edag)
			return _bdg, _gfga, _gfggb
		}
	}
	_dgec := float64(_cbf) / 255.0
	_gbc := float64(_bdgc) / 255.0
	_ada := float64(_gafe) / 255.0
	return _dgec, _gbc, _ada
}
func _bfc(_egd *GraphicSVG, _efc *_ad.ContentCreator, _feb *_fb.PdfPageResources) {
	_efc.Add_q()
	_egd.Style.toContentStream(_efc, _feb)
	_ced, _dg := _fee(_egd.Attributes["\u0078"], 64)
	if _dg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _dg.Error())
	}
	_fc, _dg := _fee(_egd.Attributes["\u0079"], 64)
	if _dg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _dg.Error())
	}
	_cebg, _dg := _fee(_egd.Attributes["\u0077\u0069\u0064t\u0068"], 64)
	if _dg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _dg.Error())
	}
	_cea, _dg := _fee(_egd.Attributes["\u0068\u0065\u0069\u0067\u0068\u0074"], 64)
	if _dg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _dg.Error())
	}
	_efc.Add_re(_ced*_egd._ccf, _fc*_egd._ccf, _cebg*_egd._ccf, _cea*_egd._ccf)
	_egd.Style.fillStroke(_efc)
	_efc.Add_Q()
}
func _eadd(_cebf string) (*Path, error) {
	_daa = _dadbe()
	_cgbg, _gfd := _ebf(_fac(_cebf))
	if _gfd != nil {
		return nil, _gfd
	}
	return _effef(_cgbg), nil
}
func _effef(_abda []*Command) *Path {
	_acbea := &Path{}
	var _aaf []*Command
	for _cgcf, _aafg := range _abda {
		switch _eg.ToLower(_aafg.Symbol) {
		case _daa._dacb:
			if len(_aaf) > 0 {
				_acbea.Subpaths = append(_acbea.Subpaths, &Subpath{_aaf})
			}
			_aaf = []*Command{_aafg}
		case _daa._dec:
			_aaf = append(_aaf, _aafg)
			_acbea.Subpaths = append(_acbea.Subpaths, &Subpath{_aaf})
			_aaf = []*Command{}
		default:
			_aaf = append(_aaf, _aafg)
			if len(_abda) == _cgcf+1 {
				_acbea.Subpaths = append(_acbea.Subpaths, &Subpath{_aaf})
			}
		}
	}
	return _acbea
}
func _egab(_bdeg _bg.StartElement) *GraphicSVG {
	_dcg := &GraphicSVG{}
	_ebb := make(map[string]string)
	for _, _eab := range _bdeg.Attr {
		_ebb[_eab.Name.Local] = _eab.Value
	}
	_dcg.Name = _bdeg.Name.Local
	_dcg.Attributes = _ebb
	_dcg._ccf = 1
	if _dcg.Name == "\u0073\u0076\u0067" {
		_fdf, _ec := _dgef(_ebb["\u0076i\u0065\u0077\u0042\u006f\u0078"])
		if _ec != nil {
			_fa.Log.Debug("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076", _ec)
			return nil
		}
		_dcg.ViewBox.X = _fdf[0]
		_dcg.ViewBox.Y = _fdf[1]
		_dcg.ViewBox.W = _fdf[2]
		_dcg.ViewBox.H = _fdf[3]
		_dcg.Width = _dcg.ViewBox.W
		_dcg.Height = _dcg.ViewBox.H
		if _eeaf, _eaa := _ebb["\u0077\u0069\u0064t\u0068"]; _eaa {
			_afae, _gccb := _fee(_eeaf, 64)
			if _gccb != nil {
				_fa.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _gccb)
				return nil
			}
			_dcg.Width = _afae
		}
		if _ccc, _cbc := _ebb["\u0068\u0065\u0069\u0067\u0068\u0074"]; _cbc {
			_eaaa, _fcff := _fee(_ccc, 64)
			if _fcff != nil {
				_fa.Log.Debug("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _fcff)
				return nil
			}
			_dcg.Height = _eaaa
		}
		if _dcg.Width > 0 && _dcg.Height > 0 {
			_dcg._ccf = _dcg.Width / _dcg.ViewBox.W
		}
	}
	return _dcg
}

const (
	_ce          = 0.72
	_ae          = 28.3464
	_bgc         = _ae / 10
	_cg          = 0.551784
	_cgc         = 96
	_ef          = 16.0
	_da  float64 = 72
)

func _ggbc(_dbdc *_bg.Decoder) (*GraphicSVG, error) {
	for {
		_fgb, _gad := _dbdc.Token()
		if _fgb == nil && _gad == _eb.EOF {
			break
		}
		if _gad != nil {
			return nil, _gad
		}
		switch _dff := _fgb.(type) {
		case _bg.StartElement:
			return _egab(_dff), nil
		}
	}
	return &GraphicSVG{}, nil
}
func _gea(_bbd *GraphicSVG, _bgg *_ad.ContentCreator, _fba *_fb.PdfPageResources) {
	_bgg.Add_q()
	_bbd.Style.toContentStream(_bgg, _fba)
	_gafb, _cbg := _fee(_bbd.Attributes["\u0078\u0031"], 64)
	if _cbg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cbg.Error())
	}
	_aef, _cbg := _fee(_bbd.Attributes["\u0079\u0031"], 64)
	if _cbg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cbg.Error())
	}
	_bec, _cbg := _fee(_bbd.Attributes["\u0078\u0032"], 64)
	if _cbg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cbg.Error())
	}
	_bfe, _cbg := _fee(_bbd.Attributes["\u0079\u0032"], 64)
	if _cbg != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cbg.Error())
	}
	_bgg.Add_m(_gafb*_bbd._ccf, _aef*_bbd._ccf)
	_bgg.Add_l(_bec*_bbd._ccf, _bfe*_bbd._ccf)
	_bbd.Style.fillStroke(_bgg)
	_bgg.Add_h()
	_bgg.Add_Q()
}
func _fac(_ged string) []token {
	var (
		_ddea []token
		_feg  string
	)
	for _, _deff := range _ged {
		_fecf := string(_deff)
		switch {
		case _daa.isCommand(_fecf):
			_ddea, _feg = _bcd(_ddea, _feg)
			_ddea = append(_ddea, token{_fecf, true})
		case _fecf == "\u002e":
			if _feg == "" {
				_feg = "\u0030"
			}
			if _eg.Contains(_feg, _fecf) {
				_ddea = append(_ddea, token{_feg, false})
				_feg = "\u0030"
			}
			fallthrough
		case _fecf >= "\u0030" && _fecf <= "\u0039" || _fecf == "\u0065":
			_feg += _fecf
		case _fecf == "\u002d":
			if _eg.HasSuffix(_feg, "\u0065") {
				_feg += _fecf
			} else {
				_ddea, _ = _bcd(_ddea, _feg)
				_feg = _fecf
			}
		default:
			_ddea, _feg = _bcd(_ddea, _feg)
		}
	}
	_ddea, _ = _bcd(_ddea, _feg)
	return _ddea
}
func _dadbe() commands {
	var _cdb = map[string]int{"\u006d": 2, "\u007a": 0, "\u006c": 2, "\u0068": 1, "\u0076": 1, "\u0063": 6, "\u0073": 4, "\u0071": 4, "\u0074": 2, "\u0061": 7}
	var _dba []string
	for _fgd := range _cdb {
		_dba = append(_dba, _fgd)
	}
	return commands{_dba, _cdb, "\u006d", "\u007a"}
}
func (_ffee *Command) compare(_gdbf *Command) bool {
	if _ffee.Symbol != _gdbf.Symbol {
		return false
	}
	for _gbdd, _caaf := range _ffee.Params {
		if _caaf != _gdbf.Params[_gbdd] {
			return false
		}
	}
	return true
}
func (_ddagc *commands) isCommand(_bae string) bool {
	for _, _fefd := range _ddagc._adg {
		if _eg.ToLower(_bae) == _fefd {
			return true
		}
	}
	return false
}
func _fee(_fcfe string, _ddge int) (float64, error) {
	_ddeb, _bgfb := _bdd(_fcfe)
	_ebc, _baba := _c.ParseFloat(_ddeb, _ddge)
	if _baba != nil {
		return 0, _baba
	}
	if _geab, _fedg := _eba[_bgfb]; _fedg {
		_ebc = _ebc * _geab
	} else {
		_ebc = _ebc * _ce
	}
	return _ebc, nil
}

type Path struct{ Subpaths []*Subpath }
type Subpath struct{ Commands []*Command }

func (_aed *GraphicSVG) toContentStream(_ddb *_ad.ContentCreator, _beac *_fb.PdfPageResources) {
	_gfgg, _cdce := _gbf(_aed.Attributes, _aed._ccf)
	if _cdce != nil {
		_fa.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _cdce)
	}
	_aed.Style = _gfgg
	switch _aed.Name {
	case "\u0070\u0061\u0074\u0068":
		_cd(_aed, _ddb, _beac)
		for _, _dda := range _aed.Children {
			_dda.toContentStream(_ddb, _beac)
		}
	case "\u0072\u0065\u0063\u0074":
		_bfc(_aed, _ddb, _beac)
		for _, _cfa := range _aed.Children {
			_cfa.toContentStream(_ddb, _beac)
		}
	case "\u0063\u0069\u0072\u0063\u006c\u0065":
		_fec(_aed, _ddb, _beac)
		for _, _abd := range _aed.Children {
			_abd.toContentStream(_ddb, _beac)
		}
	case "\u0065l\u006c\u0069\u0070\u0073\u0065":
		_ede(_aed, _ddb, _beac)
		for _, _dgd := range _aed.Children {
			_dgd.toContentStream(_ddb, _beac)
		}
	case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":
		_ebg(_aed, _ddb, _beac)
		for _, _agfa := range _aed.Children {
			_agfa.toContentStream(_ddb, _beac)
		}
	case "\u0070o\u006c\u0079\u0067\u006f\u006e":
		_db(_aed, _ddb, _beac)
		for _, _cfe := range _aed.Children {
			_cfe.toContentStream(_ddb, _beac)
		}
	case "\u006c\u0069\u006e\u0065":
		_gea(_aed, _ddb, _beac)
		for _, _abb := range _aed.Children {
			_abb.toContentStream(_ddb, _beac)
		}
	case "\u0074\u0065\u0078\u0074":
		_gfbc(_aed, _ddb, _beac)
		for _, _egaf := range _aed.Children {
			_egaf.toContentStream(_ddb, _beac)
		}
	case "\u0067":
		_abe, _cfb := _aed.Attributes["\u0066\u0069\u006c\u006c"]
		_egae, _agae := _aed.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]
		_cgfd, _bbg := _aed.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
		for _, _dddf := range _aed.Children {
			if _, _bgdc := _dddf.Attributes["\u0066\u0069\u006c\u006c"]; !_bgdc && _cfb {
				_dddf.Attributes["\u0066\u0069\u006c\u006c"] = _abe
			}
			if _, _dadb := _dddf.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]; !_dadb && _agae {
				_dddf.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"] = _egae
			}
			if _, _gaa := _dddf.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]; !_gaa && _bbg {
				_dddf.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"] = _cgfd
			}
			_dddf.toContentStream(_ddb, _beac)
		}
	}
}
func _ebf(_gefe []token) ([]*Command, error) {
	var (
		_fgf  []*Command
		_fede []float64
	)
	for _ead := len(_gefe) - 1; _ead >= 0; _ead-- {
		_edcd := _gefe[_ead]
		if _edcd._cagac {
			_afee := _daa._bfae[_eg.ToLower(_edcd._bac)]
			_ccef := len(_fede)
			if _afee == 0 && _ccef == 0 {
				_bcbg := &Command{Symbol: _edcd._bac}
				_fgf = append([]*Command{_bcbg}, _fgf...)
			} else if _afee != 0 && _ccef%_afee == 0 {
				_ggeg := _ccef / _afee
				for _efba := 0; _efba < _ggeg; _efba++ {
					_cef := _edcd._bac
					if _cef == "\u006d" && _efba < _ggeg-1 {
						_cef = "\u006c"
					}
					if _cef == "\u004d" && _efba < _ggeg-1 {
						_cef = "\u004c"
					}
					_gce := &Command{_cef, _dddb(_fede[:_afee])}
					_fgf = append([]*Command{_gce}, _fgf...)
					_fede = _fede[_afee:]
				}
			} else {
				_fge := pathParserError{"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020" + _edcd._bac}
				return nil, _fge
			}
		} else {
			_bgbb, _acba := _fee(_edcd._bac, 64)
			if _acba != nil {
				return nil, _acba
			}
			_fede = append(_fede, _bgbb)
		}
	}
	return _fgf, nil
}
func _cdbg(_cec string) (_bdea []float64, _geb error) {
	var _eebd float64
	_fdcg := 0
	_cafc := true
	for _gcef, _acbae := range _cec {
		if _acbae == '.' {
			if _cafc {
				_cafc = false
				continue
			}
			_eebd, _geb = _fee(_cec[_fdcg:_gcef], 64)
			if _geb != nil {
				return
			}
			_bdea = append(_bdea, _eebd)
			_fdcg = _gcef
		}
	}
	_eebd, _geb = _fee(_cec[_fdcg:], 64)
	if _geb != nil {
		return
	}
	_bdea = append(_bdea, _eebd)
	return
}

type token struct {
	_bac   string
	_cagac bool
}

func _dcb(_baf string) (float64, float64, float64) {
	_dgdc := _eg.TrimPrefix(_baf, "\u0072\u0067\u0062\u0028")
	_dgdc = _eg.TrimSuffix(_dgdc, "\u0029")
	_eaab := _eg.Split(_dgdc, "\u002c")
	if len(_eaab) != 3 {
		_fa.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0067\u0062\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020%\u0073", _baf)
		return 0, 0, 0
	}
	var _bad, _ddc, _fecg float64
	_bad, _gdc := _cga(_eaab[0])
	if _gdc != nil {
		_fa.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0067\u0062\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020%\u0073", _baf)
		return 0, 0, 0
	}
	_ddc, _gdc = _cga(_eaab[1])
	if _gdc != nil {
		_fa.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0067\u0062\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020%\u0073", _baf)
		return 0, 0, 0
	}
	_fecg, _gdc = _cga(_eaab[2])
	if _gdc != nil {
		_fa.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0067\u0062\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020%\u0073", _baf)
		return 0, 0, 0
	}
	_efab := _bad / 255.0
	_cbcf := _ddc / 255.0
	_ffd := _fecg / 255.0
	return _efab, _cbcf, _ffd
}
func _db(_eag *GraphicSVG, _dfc *_ad.ContentCreator, _cb *_fb.PdfPageResources) {
	_dfc.Add_q()
	_eag.Style.toContentStream(_dfc, _cb)
	_bce, _befc := _dgef(_eag.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _befc != nil {
		_fa.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _befc)
		return
	}
	if len(_bce)%2 > 0 {
		_fa.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _eeb := 0; _eeb < len(_bce); {
		if _eeb == 0 {
			_dfc.Add_m(_bce[_eeb]*_eag._ccf, _bce[_eeb+1]*_eag._ccf)
		} else {
			_dfc.Add_l(_bce[_eeb]*_eag._ccf, _bce[_eeb+1]*_eag._ccf)
		}
		_eeb += 2
	}
	_dfc.Add_l(_bce[0]*_eag._ccf, _bce[1]*_eag._ccf)
	_eag.Style.fillStroke(_dfc)
	_dfc.Add_h()
	_dfc.Add_Q()
}
func _ede(_agf *GraphicSVG, _bef *_ad.ContentCreator, _bd *_fb.PdfPageResources) {
	_bef.Add_q()
	_agf.Style.toContentStream(_bef, _bd)
	_ge, _bea := _fee(_agf.Attributes["\u0063\u0078"], 64)
	if _bea != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _bea.Error())
	}
	_cgec, _bea := _fee(_agf.Attributes["\u0063\u0079"], 64)
	if _bea != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _bea.Error())
	}
	_ceaa, _bea := _fee(_agf.Attributes["\u0072\u0078"], 64)
	if _bea != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _bea.Error())
	}
	_bab, _bea := _fee(_agf.Attributes["\u0072\u0079"], 64)
	if _bea != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _bea.Error())
	}
	_dd := _ceaa * _agf._ccf
	_gef := _bab * _agf._ccf
	_fcg := _ge * _agf._ccf
	_gaf := _cgec * _agf._ccf
	_fggd := _dd * _cg
	_fea := _gef * _cg
	_eff := _af.NewCubicBezierPath()
	_eff = _eff.AppendCurve(_af.NewCubicBezierCurve(-_dd, 0, -_dd, _fea, -_fggd, _gef, 0, _gef))
	_eff = _eff.AppendCurve(_af.NewCubicBezierCurve(0, _gef, _fggd, _gef, _dd, _fea, _dd, 0))
	_eff = _eff.AppendCurve(_af.NewCubicBezierCurve(_dd, 0, _dd, -_fea, _fggd, -_gef, 0, -_gef))
	_eff = _eff.AppendCurve(_af.NewCubicBezierCurve(0, -_gef, -_fggd, -_gef, -_dd, -_fea, -_dd, 0))
	_eff = _eff.Offset(_fcg, _gaf)
	if _agf.Style.StrokeWidth > 0 {
		_eff = _eff.Offset(_agf.Style.StrokeWidth/2, _agf.Style.StrokeWidth/2)
	}
	_af.DrawBezierPathWithCreator(_eff, _bef)
	_agf.Style.fillStroke(_bef)
	_bef.Add_h()
	_bef.Add_Q()
}
func _dddb(_dfb []float64) []float64 {
	for _bee, _eabd := 0, len(_dfb)-1; _bee < _eabd; _bee, _eabd = _bee+1, _eabd-1 {
		_dfb[_bee], _dfb[_eabd] = _dfb[_eabd], _dfb[_bee]
	}
	return _dfb
}
func _ebde(_bece string) (*_fb.PdfFont, error) {
	_acb, _bage := map[string]_fb.StdFontName{"\u0063o\u0075\u0072\u0069\u0065\u0072": _fb.CourierName, "\u0063\u006f\u0075r\u0069\u0065\u0072\u002d\u0062\u006f\u006c\u0064": _fb.CourierBoldName, "\u0063o\u0075r\u0069\u0065\u0072\u002d\u006f\u0062\u006c\u0069\u0071\u0075\u0065": _fb.CourierObliqueName, "c\u006fu\u0072\u0069\u0065\u0072\u002d\u0062\u006f\u006cd\u002d\u006f\u0062\u006ciq\u0075\u0065": _fb.CourierBoldObliqueName, "\u0068e\u006c\u0076\u0065\u0074\u0069\u0063a": _fb.HelveticaName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0062\u006f\u006c\u0064": _fb.HelveticaBoldName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u006f\u0062l\u0069\u0071\u0075\u0065": _fb.HelveticaObliqueName, "\u0068\u0065\u006c\u0076et\u0069\u0063\u0061\u002d\u0062\u006f\u006c\u0064\u002d\u006f\u0062\u006c\u0069\u0071u\u0065": _fb.HelveticaBoldObliqueName, "\u0073\u0079\u006d\u0062\u006f\u006c": _fb.SymbolName, "\u007a\u0061\u0070\u0066\u002d\u0064\u0069\u006e\u0067\u0062\u0061\u0074\u0073": _fb.ZapfDingbatsName, "\u0074\u0069\u006de\u0073": _fb.TimesRomanName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064": _fb.TimesBoldName, "\u0074\u0069\u006de\u0073\u002d\u0069\u0074\u0061\u006c\u0069\u0063": _fb.TimesItalicName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064\u002d\u0069t\u0061\u006c\u0069\u0063": _fb.TimesBoldItalicName}[_bece]
	if !_bage {
		return nil, _g.Errorf("\u0066\u006f\u006e\u0074\u002df\u0061\u006d\u0069\u006c\u0079\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u006c\u0069\u0073t", _bece)
	}
	_bde, _agd := _fb.NewStandard14Font(_acb)
	if _agd != nil {
		return nil, _agd
	}
	return _bde, nil
}
func (_gfa *GraphicSVG) SetScaling(xFactor, yFactor float64) {
	_eefb := _gfa.Width / _gfa.ViewBox.W
	_gee := _gfa.Height / _gfa.ViewBox.H
	_gfa.setDefaultScaling(_a.Max(_eefb, _gee))
	for _, _eagf := range _gfa.Children {
		_eagf.SetScaling(xFactor, yFactor)
	}
}
func ParseFromString(svgStr string) (*GraphicSVG, error) {
	return ParseFromStream(_eg.NewReader(svgStr))
}

type Command struct {
	Symbol string
	Params []float64
}

func _ccbf(_cdd float64) int { return int(_cdd + _a.Copysign(0.5, _cdd)) }
func _gab(_bgff float64, _eagg int) float64 {
	_afb := _a.Pow(10, float64(_eagg))
	return float64(_ccbf(_bgff*_afb)) / _afb
}
func _fec(_bca *GraphicSVG, _gga *_ad.ContentCreator, _ag *_fb.PdfPageResources) {
	_gga.Add_q()
	_bca.Style.toContentStream(_gga, _ag)
	_age, _gfb := _fee(_bca.Attributes["\u0063\u0078"], 64)
	if _gfb != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _gfb.Error())
	}
	_bcb, _gfb := _fee(_bca.Attributes["\u0063\u0079"], 64)
	if _gfb != nil {
		_fa.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _gfb.Error())
	}
	_cad, _gfb := _fee(_bca.Attributes["\u0072"], 64)
	if _gfb != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _gfb.Error())
	}
	_adb := _cad * _bca._ccf
	_gcfe := _cad * _bca._ccf
	_aga := _adb * _cg
	_eea := _gcfe * _cg
	_cge := _af.NewCubicBezierPath()
	_cge = _cge.AppendCurve(_af.NewCubicBezierCurve(-_adb, 0, -_adb, _eea, -_aga, _gcfe, 0, _gcfe))
	_cge = _cge.AppendCurve(_af.NewCubicBezierCurve(0, _gcfe, _aga, _gcfe, _adb, _eea, _adb, 0))
	_cge = _cge.AppendCurve(_af.NewCubicBezierCurve(_adb, 0, _adb, -_eea, _aga, -_gcfe, 0, -_gcfe))
	_cge = _cge.AppendCurve(_af.NewCubicBezierCurve(0, -_gcfe, -_aga, -_gcfe, -_adb, -_eea, -_adb, 0))
	_cge = _cge.Offset(_age*_bca._ccf, _bcb*_bca._ccf)
	if _bca.Style.StrokeWidth > 0 {
		_cge = _cge.Offset(_bca.Style.StrokeWidth/2, _bca.Style.StrokeWidth/2)
	}
	_af.DrawBezierPathWithCreator(_cge, _gga)
	_bca.Style.fillStroke(_gga)
	_gga.Add_h()
	_gga.Add_Q()
}
func _gfbc(_ebd *GraphicSVG, _ade *_ad.ContentCreator, _cab *_fb.PdfPageResources) {
	_ade.Add_BT()
	_ege, _dbd := _fee(_ebd.Attributes["\u0078"], 64)
	if _dbd != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _dbd.Error())
	}
	_fad, _dbd := _fee(_ebd.Attributes["\u0079"], 64)
	if _dbd != nil {
		_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _dbd.Error())
	}
	_cgb := _ebd.Attributes["\u0066\u0069\u006c\u006c"]
	var _cgf, _cac, _beab float64
	if _bcc, _caf := _d.ColorMap[_cgb]; _caf {
		_fcb, _ccd, _agb, _ := _bcc.RGBA()
		_cgf, _cac, _beab = float64(_fcb), float64(_ccd), float64(_agb)
	} else if _eg.HasPrefix(_cgb, "\u0072\u0067\u0062\u0028") {
		_cgf, _cac, _beab = _dcb(_cgb)
	} else {
		_cgf, _cac, _beab = _efe(_cgb)
	}
	_ade.Add_rg(_cgf, _cac, _beab)
	_cafe := _ef
	if _ebe, _ggb := _ebd.Attributes["\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae"]; _ggb {
		_cafe, _dbd = _c.ParseFloat(_ebe, 64)
		if _dbd != nil {
			_fa.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067 \u0060\u0066\u006f\u006e\u0074\u002d\u0073\u0069\u007a\u0065\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _dbd.Error())
			_cafe = _ef
		}
	}
	_cabb := _ebd._ccf * _cafe * _da / _cgc
	_agc := _gd.PdfObjectName("\u0053\u0046\u006fn\u0074")
	_bfa := _fb.DefaultFont()
	_gfbe, _bag := _ebd.Attributes["f\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069\u006c\u0079"]
	if _bag {
		if _dgg, _dge := _ebde(_gfbe); _dge == nil {
			_bfa = _dgg
			_geaa := 1
			for _cab.HasFontByName(_agc) {
				_agc = _gd.PdfObjectName("\u0053\u0046\u006fn\u0074" + _c.Itoa(_geaa))
				_geaa++
			}
		}
	}
	_dac := 0.0
	_aegc, _bag := _ebd.Attributes["t\u0065\u0078\u0074\u002d\u0061\u006e\u0063\u0068\u006f\u0072"]
	if _bag && _aegc != "\u0073\u0074\u0061r\u0074" {
		var _caa float64
		for _, _afcc := range _ebd.Content {
			_bff, _gcc := _bfa.GetRuneMetrics(_afcc)
			if !_gcc {
				_fa.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _afcc)
			}
			_caa += _bff.Wx
		}
		_caa = _caa * _cabb / 1000.0
		if _aegc == "\u006d\u0069\u0064\u0064\u006c\u0065" {
			_dac = -_caa / 2
		} else if _aegc == "\u0065\u006e\u0064" {
			_dac = -_caa
		}
	}
	_ade.Add_Tm(1, 0, 0, -1, _ege*_ebd._ccf+_dac, _fad*_ebd._ccf)
	_cab.SetFontByName(_agc, _bfa.ToPdfObject())
	_ade.Add_Tf(_agc, _cabb)
	_dad := _ebd.Content
	_eaf := _gd.MakeString(_dad)
	_ade.Add_Tj(*_eaf)
	_ade.Add_ET()
}

type GraphicSVGStyle struct {
	FillColor   string
	StrokeColor string
	StrokeWidth float64
	FillOpacity float64
}

func (_bced *GraphicSVGStyle) fillStroke(_cee *_ad.ContentCreator) {
	if _bced.FillColor != "" && _bced.StrokeColor != "" {
		_cee.Add_B()
	} else if _bced.FillColor != "" {
		_cee.Add_f()
	} else if _bced.StrokeColor != "" {
		_cee.Add_S()
	}
}
func _cd(_bc *GraphicSVG, _df *_ad.ContentCreator, _fg *_fb.PdfPageResources) {
	_df.Add_q()
	_bc.Style.toContentStream(_df, _fg)
	_be, _ba := _eadd(_bc.Attributes["\u0064"])
	if _ba != nil {
		_fa.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s", _ba.Error())
	}
	var (
		_gc, _fga = 0.0, 0.0
		_ed, _ff  = 0.0, 0.0
		_cdc      *Command
	)
	for _, _gdf := range _be.Subpaths {
		for _, _bf := range _gdf.Commands {
			switch _eg.ToLower(_bf.Symbol) {
			case "\u006d":
				_ed, _ff = _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf
				if !_bf.isAbsolute() {
					_ed, _ff = _gc+_ed-_bc.ViewBox.X, _fga+_ff-_bc.ViewBox.Y
				}
				_df.Add_m(_gab(_ed, 3), _gab(_ff, 3))
				_gc, _fga = _ed, _ff
			case "\u0063":
				_ab, _bgd, _ee, _ca, _cgd, _ffc := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf, _bf.Params[2]*_bc._ccf, _bf.Params[3]*_bc._ccf, _bf.Params[4]*_bc._ccf, _bf.Params[5]*_bc._ccf
				if !_bf.isAbsolute() {
					_ab, _bgd, _ee, _ca, _cgd, _ffc = _gc+_ab, _fga+_bgd, _gc+_ee, _fga+_ca, _gc+_cgd, _fga+_ffc
				}
				_df.Add_c(_gab(_ab, 3), _gab(_bgd, 3), _gab(_ee, 3), _gab(_ca, 3), _gab(_cgd, 3), _gab(_ffc, 3))
				_gc, _fga = _cgd, _ffc
			case "\u0073":
				_gg, _eee, _edf, _gf := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf, _bf.Params[2]*_bc._ccf, _bf.Params[3]*_bc._ccf
				if !_bf.isAbsolute() {
					_gg, _eee, _edf, _gf = _gc+_gg, _fga+_eee, _gc+_edf, _fga+_gf
				}
				_df.Add_c(_gab(_gc, 3), _gab(_fga, 3), _gab(_gg, 3), _gab(_eee, 3), _gab(_edf, 3), _gab(_gf, 3))
				_gc, _fga = _edf, _gf
			case "\u006c":
				_ga, _eef := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf
				if !_bf.isAbsolute() {
					_ga, _eef = _gc+_ga, _fga+_eef
				}
				_df.Add_l(_gab(_ga, 3), _gab(_eef, 3))
				_gc, _fga = _ga, _eef
			case "\u0068":
				_aeg := _bf.Params[0] * _bc._ccf
				if !_bf.isAbsolute() {
					_aeg = _gc + _aeg
				}
				_df.Add_l(_gab(_aeg, 3), _gab(_fga, 3))
				_gc = _aeg
			case "\u0076":
				_de := _bf.Params[0] * _bc._ccf
				if !_bf.isAbsolute() {
					_de = _fga + _de
				}
				_df.Add_l(_gab(_gc, 3), _gab(_de, 3))
				_fga = _de
			case "\u0071":
				_ffg, _bfb, _fe, _gfg := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf, _bf.Params[2]*_bc._ccf, _bf.Params[3]*_bc._ccf
				if !_bf.isAbsolute() {
					_ffg, _bfb, _fe, _gfg = _gc+_ffg, _fga+_bfb, _gc+_fe, _fga+_gfg
				}
				_fed, _ac := _d.QuadraticToCubicBezier(_gc, _fga, _ffg, _bfb, _fe, _gfg)
				_df.Add_c(_gab(_fed.X, 3), _gab(_fed.Y, 3), _gab(_ac.X, 3), _gab(_ac.Y, 3), _gab(_fe, 3), _gab(_gfg, 3))
				_gc, _fga = _fe, _gfg
			case "\u0074":
				var _aa, _ffe _d.Point
				_deb, _ea := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf
				if !_bf.isAbsolute() {
					_deb, _ea = _gc+_deb, _fga+_ea
				}
				if _cdc != nil && _eg.ToLower(_cdc.Symbol) == "\u0071" {
					_aae := _d.Point{X: _cdc.Params[0] * _bc._ccf, Y: _cdc.Params[1] * _bc._ccf}
					_fbf := _d.Point{X: _cdc.Params[2] * _bc._ccf, Y: _cdc.Params[3] * _bc._ccf}
					_fgg := _fbf.Mul(2.0).Sub(_aae)
					_aa, _ffe = _d.QuadraticToCubicBezier(_gc, _fga, _fgg.X, _fgg.Y, _deb, _ea)
				}
				_df.Add_c(_gab(_aa.X, 3), _gab(_aa.Y, 3), _gab(_ffe.X, 3), _gab(_ffe.Y, 3), _gab(_deb, 3), _gab(_ea, 3))
				_gc, _fga = _deb, _ea
			case "\u0061":
				_ffge, _def := _bf.Params[0]*_bc._ccf, _bf.Params[1]*_bc._ccf
				_fggf := _bf.Params[2]
				_cc := _bf.Params[3] > 0
				_gcf := _bf.Params[4] > 0
				_eae, _ega := _bf.Params[5]*_bc._ccf, _bf.Params[6]*_bc._ccf
				if !_bf.isAbsolute() {
					_eae, _ega = _gc+_eae, _fga+_ega
				}
				_egb := _d.EllipseToCubicBeziers(_gc, _fga, _ffge, _def, _fggf, _cc, _gcf, _eae, _ega)
				for _, _ceb := range _egb {
					_df.Add_c(_gab(_ceb[1].X, 3), _gab((_ceb[1].Y), 3), _gab((_ceb[2].X), 3), _gab((_ceb[2].Y), 3), _gab((_ceb[3].X), 3), _gab((_ceb[3].Y), 3))
				}
				_gc, _fga = _eae, _ega
			case "\u007a":
				_df.Add_h()
			}
			_cdc = _bf
		}
	}
	_bc.Style.fillStroke(_df)
	_df.Add_h()
	_df.Add_Q()
}
func _bcd(_efaa []token, _cfd string) ([]token, string) {
	if _cfd != "" {
		_efaa = append(_efaa, token{_cfd, false})
		_cfd = ""
	}
	return _efaa, _cfd
}
func _ecd(_dbg string) (float64, error) {
	_dbg = _eg.TrimSpace(_dbg)
	var _bbbe float64
	if _eg.HasSuffix(_dbg, "\u0025") {
		_gae, _cafbf := _c.ParseFloat(_eg.TrimSuffix(_dbg, "\u0025"), 64)
		if _cafbf != nil {
			return 0, _cafbf
		}
		_bbbe = _gae / 100.0
	} else {
		_bdbc, _ffbe := _c.ParseFloat(_dbg, 64)
		if _ffbe != nil {
			return 0, _ffbe
		}
		_bbbe = _bdbc
	}
	return _bbbe, nil
}

var _daa commands
var (
	_fbg = []string{"\u0063\u006d", "\u006d\u006d", "\u0070\u0078", "\u0070\u0074"}
	_eba = map[string]float64{"\u0063\u006d": _ae, "\u006d\u006d": _bgc, "\u0070\u0078": _ce, "\u0070\u0074": 1}
)

func (_cce *GraphicSVG) Decode(decoder *_bg.Decoder) error {
	for {
		_cbe, _cag := decoder.Token()
		if _cbe == nil && _cag == _eb.EOF {
			break
		}
		if _cag != nil {
			return _cag
		}
		switch _daf := _cbe.(type) {
		case _bg.StartElement:
			_edfd := _egab(_daf)
			_eaed := _edfd.Decode(decoder)
			if _eaed != nil {
				return _eaed
			}
			_cce.Children = append(_cce.Children, _edfd)
		case _bg.CharData:
			_fce := _eg.TrimSpace(string(_daf))
			if _fce != "" {
				_cce.Content = string(_daf)
			}
		case _bg.EndElement:
			if _daf.Name.Local == _cce.Name {
				return nil
			}
		}
	}
	return nil
}
func (_ggba pathParserError) Error() string { return _ggba._ffb }
