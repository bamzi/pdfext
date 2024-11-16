package fonts

import (
	_gf "bytes"
	_dg "encoding/binary"
	_c "errors"
	_f "fmt"
	_e "io"
	_d "os"
	_ab "regexp"
	_ae "sort"
	_a "strings"
	_fg "sync"

	_fb "github.com/bamzi/pdfext/common"
	_ag "github.com/bamzi/pdfext/core"
	_de "github.com/bamzi/pdfext/internal/cmap"
	_b "github.com/bamzi/pdfext/internal/textencoding"
)

var _gdd *RuneCharSafeMap
var _acde *RuneCharSafeMap

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

func NewStdFont(desc Descriptor, metrics *RuneCharSafeMap) StdFont {
	return NewStdFontWithEncoding(desc, metrics, _b.NewStandardEncoder())
}

type Descriptor struct {
	Name        StdFontName
	Family      string
	Weight      FontWeight
	Flags       uint
	BBox        [4]float64
	ItalicAngle float64
	Ascent      float64
	Descent     float64
	CapHeight   float64
	XHeight     float64
	StemV       float64
	StemH       float64
}

func (_dde StdFont) GetMetricsTable() *RuneCharSafeMap { return _dde._dc }
func NewStdFontByName(name StdFontName) (StdFont, bool) {
	_eb, _bae := _dfb.read(name)
	if !_bae {
		return StdFont{}, false
	}
	return _eb(), true
}
func _ggd() StdFont {
	_dagc.Do(_eac)
	_eaee := Descriptor{Name: TimesItalicName, Family: _bef, Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-169, -217, 1010, 883}, ItalicAngle: -15.5, Ascent: 683, Descent: -217, CapHeight: 653, XHeight: 441, StemV: 76, StemH: 32}
	return NewStdFont(_eaee, _dcd)
}
func (_cg *fontMap) write(_ed StdFontName, _ecg func() StdFont) {
	_cg.Lock()
	defer _cg.Unlock()
	_cg._bc[_ed] = _ecg
}
func (_fed *ttfParser) readByte() (_ccge uint8) {
	_dg.Read(_fed._gbca, _dg.BigEndian, &_ccge)
	return _ccge
}

const (
	CourierName            = StdFontName("\u0043o\u0075\u0072\u0069\u0065\u0072")
	CourierBoldName        = StdFontName("\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064")
	CourierObliqueName     = StdFontName("\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
	CourierBoldObliqueName = StdFontName("\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
)

func (_bcdf *ttfParser) Skip(n int) { _bcdf._gbca.Seek(int64(n), _e.SeekCurrent) }

type GlyphName = _b.GlyphName

var _baed *RuneCharSafeMap

func (_gb *fontMap) read(_ba StdFontName) (func() StdFont, bool) {
	_gb.Lock()
	defer _gb.Unlock()
	_gd, _fgg := _gb._bc[_ba]
	return _gd, _fgg
}
func (_fbd *ttfParser) ParseComponents() error {
	if _deffb := _fbd.ParseHead(); _deffb != nil {
		return _deffb
	}
	if _aac := _fbd.ParseHhea(); _aac != nil {
		return _aac
	}
	if _bb := _fbd.ParseMaxp(); _bb != nil {
		return _bb
	}
	if _gfg := _fbd.ParseHmtx(); _gfg != nil {
		return _gfg
	}
	if _, _gfc := _fbd._ggc["\u006e\u0061\u006d\u0065"]; _gfc {
		if _afga := _fbd.ParseName(); _afga != nil {
			return _afga
		}
	}
	if _, _ggf := _fbd._ggc["\u004f\u0053\u002f\u0032"]; _ggf {
		if _fdcb := _fbd.ParseOS2(); _fdcb != nil {
			return _fdcb
		}
	}
	if _, _acgg := _fbd._ggc["\u0070\u006f\u0073\u0074"]; _acgg {
		if _gdadc := _fbd.ParsePost(); _gdadc != nil {
			return _gdadc
		}
	}
	if _, _ebba := _fbd._ggc["\u0063\u006d\u0061\u0070"]; _ebba {
		if _bcb := _fbd.ParseCmap(); _bcb != nil {
			return _bcb
		}
	}
	return nil
}

var _ddeb = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 556, 722, 722, 611, 611, 611, 611, 611, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 722, 556, 556, 584, 584, 389, 975, 556, 611, 278, 280, 389, 389, 333, 333, 333, 280, 350, 556, 556, 333, 556, 556, 333, 556, 333, 333, 278, 250, 737, 556, 611, 556, 556, 743, 611, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 611, 333, 333, 333, 611, 556, 611, 556, 556, 167, 611, 611, 611, 611, 333, 584, 549, 556, 556, 333, 333, 611, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 556, 556, 278, 278, 400, 278, 584, 549, 584, 494, 278, 889, 333, 584, 611, 584, 611, 611, 611, 611, 556, 549, 611, 556, 611, 611, 611, 611, 944, 333, 611, 611, 611, 556, 834, 834, 333, 370, 365, 611, 611, 611, 556, 333, 333, 494, 889, 278, 278, 1000, 584, 584, 611, 611, 611, 474, 500, 500, 500, 278, 278, 278, 238, 389, 389, 549, 389, 389, 737, 333, 556, 556, 556, 556, 556, 556, 333, 556, 556, 278, 278, 556, 600, 333, 389, 333, 611, 556, 834, 333, 333, 1000, 556, 333, 611, 611, 611, 611, 611, 611, 611, 556, 611, 611, 556, 778, 556, 556, 556, 556, 556, 500, 500, 500, 500, 556}
var _gcb *RuneCharSafeMap

func _egf() StdFont {
	_cdab := _b.NewSymbolEncoder()
	_eec := Descriptor{Name: SymbolName, Family: string(SymbolName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-180, -293, 1090, 1010}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 85, StemH: 92}
	return NewStdFontWithEncoding(_eec, _gcd, _cdab)
}
func (_dgf *TtfType) NewEncoder() _b.TextEncoder { return _b.NewTrueTypeFontEncoder(_dgf.Chars) }
func (_ece *ttfParser) parseTTC() (TtfType, error) {
	_ece.Skip(2 * 2)
	_egd := _ece.ReadULong()
	if _egd < 1 {
		return TtfType{}, _c.New("N\u006f \u0066\u006f\u006e\u0074\u0073\u0020\u0069\u006e \u0054\u0054\u0043\u0020fi\u006c\u0065")
	}
	_ff := _ece.ReadULong()
	_, _fa := _ece._gbca.Seek(int64(_ff), _e.SeekStart)
	if _fa != nil {
		return TtfType{}, _fa
	}
	return _ece.Parse()
}

var _eae = []int16{667, 1000, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 500, 667, 667, 556, 556, 556, 556, 556, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 667, 556, 556, 469, 584, 389, 1015, 556, 556, 278, 260, 334, 334, 278, 278, 333, 260, 350, 500, 500, 333, 500, 500, 333, 556, 333, 278, 278, 250, 737, 556, 556, 556, 556, 643, 556, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 556, 278, 333, 278, 500, 556, 500, 556, 556, 167, 556, 556, 556, 611, 333, 584, 549, 556, 556, 333, 333, 556, 333, 333, 222, 278, 278, 278, 278, 278, 222, 222, 500, 500, 222, 222, 299, 222, 584, 549, 584, 471, 222, 833, 333, 584, 556, 584, 556, 556, 556, 556, 556, 549, 556, 556, 556, 556, 556, 556, 944, 333, 556, 556, 556, 556, 834, 834, 333, 370, 365, 611, 556, 556, 537, 333, 333, 476, 889, 278, 278, 1000, 584, 584, 556, 556, 611, 355, 333, 333, 333, 222, 222, 222, 191, 333, 333, 453, 333, 333, 737, 333, 500, 500, 500, 500, 500, 556, 278, 556, 556, 278, 278, 556, 600, 278, 317, 278, 556, 556, 834, 333, 333, 1000, 556, 333, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 500, 722, 500, 500, 500, 500, 556, 500, 500, 500, 500, 556}
var _dagc _fg.Once

func _gce(_bee map[string]uint32) string {
	var _bdb []string
	for _ecec := range _bee {
		_bdb = append(_bdb, _ecec)
	}
	_ae.Slice(_bdb, func(_bfc, _fda int) bool { return _bee[_bdb[_bfc]] < _bee[_bdb[_fda]] })
	_bcf := []string{_f.Sprintf("\u0054\u0072\u0075\u0065Ty\u0070\u0065\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_bee))}
	for _, _ded := range _bdb {
		_bcf = append(_bcf, _f.Sprintf("\u0009%\u0071\u0020\u0025\u0035\u0064", _ded, _bee[_ded]))
	}
	return _a.Join(_bcf, "\u000a")
}
func _acd() StdFont {
	_ddea.Do(_gff)
	_bd := Descriptor{Name: CourierName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0021, BBox: [4]float64{-23, -250, 715, 805}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_bd, _abd)
}

var _ebb = []int16{722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 556, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 389, 722, 722, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 556, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 556, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 611, 611, 611, 611, 444, 444, 444, 444, 333, 444, 667, 444, 444, 778, 444, 444, 469, 541, 500, 921, 444, 500, 278, 200, 480, 480, 333, 333, 333, 200, 350, 444, 444, 333, 444, 444, 333, 500, 333, 278, 250, 250, 760, 500, 500, 500, 500, 588, 500, 400, 333, 564, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 564, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 564, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 344, 278, 564, 549, 564, 471, 278, 778, 333, 564, 500, 564, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 453, 333, 333, 476, 833, 250, 250, 1000, 564, 564, 500, 444, 444, 408, 444, 444, 444, 333, 333, 333, 180, 333, 333, 453, 333, 333, 760, 333, 389, 389, 389, 389, 389, 500, 278, 500, 500, 278, 250, 500, 600, 278, 326, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}
var _gbc *RuneCharSafeMap

func _dfff() StdFont {
	_dfe.Do(_acg)
	_dag := Descriptor{Name: HelveticaBoldName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-170, -228, 1003, 962}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_dag, _bge)
}
func (_aef *ttfParser) parseCmapSubtable10(_aae int64) error {
	if _aef._dad.Chars == nil {
		_aef._dad.Chars = make(map[rune]GID)
	}
	_aef._gbca.Seek(int64(_aef._ggc["\u0063\u006d\u0061\u0070"])+_aae, _e.SeekStart)
	var _fff, _cdf uint32
	_gdde := _aef.ReadUShort()
	if _gdde < 8 {
		_fff = uint32(_aef.ReadUShort())
		_cdf = uint32(_aef.ReadUShort())
	} else {
		_aef.ReadUShort()
		_fff = _aef.ReadULong()
		_cdf = _aef.ReadULong()
	}
	_fb.Log.Trace("\u0070\u0061r\u0073\u0065\u0043\u006d\u0061p\u0053\u0075\u0062\u0074\u0061b\u006c\u0065\u0031\u0030\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003d\u0025\u0064", _gdde, _fff, _cdf)
	if _gdde != 0 {
		return _c.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006d\u0061p\u0020s\u0075\u0062\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	_aaed, _bbf := _aef.ReadStr(256)
	if _bbf != nil {
		return _bbf
	}
	_bbg := []byte(_aaed)
	for _dfc, _beg := range _bbg {
		_aef._dad.Chars[rune(_dfc)] = GID(_beg)
		if _beg != 0 {
			_f.Printf("\u0009\u0030\u0078\u002502\u0078\u0020\u279e\u0020\u0030\u0078\u0025\u0030\u0032\u0078\u003d\u0025\u0063\u000a", _dfc, _beg, rune(_beg))
		}
	}
	return nil
}
func (_gcad *ttfParser) parseCmapVersion(_fdf int64) error {
	_fb.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006da\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a \u006f\u0066\u0066s\u0065t\u003d\u0025\u0064", _fdf)
	if _gcad._dad.Chars == nil {
		_gcad._dad.Chars = make(map[rune]GID)
	}
	_gcad._gbca.Seek(int64(_gcad._ggc["\u0063\u006d\u0061\u0070"])+_fdf, _e.SeekStart)
	var _eecf, _gbcg uint32
	_afb := _gcad.ReadUShort()
	if _afb < 8 {
		_eecf = uint32(_gcad.ReadUShort())
		_gbcg = uint32(_gcad.ReadUShort())
	} else {
		_gcad.ReadUShort()
		_eecf = _gcad.ReadULong()
		_gbcg = _gcad.ReadULong()
	}
	_fb.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0043m\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064 \u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067u\u0061g\u0065\u003d\u0025\u0064", _afb, _eecf, _gbcg)
	switch _afb {
	case 0:
		return _gcad.parseCmapFormat0()
	case 6:
		return _gcad.parseCmapFormat6()
	case 12:
		return _gcad.parseCmapFormat12()
	default:
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063m\u0061\u0070\u0020\u0066\u006f\u0072\u006da\u0074\u003d\u0025\u0064", _afb)
		return nil
	}
}

type RuneCharSafeMap struct {
	_ga  map[rune]CharMetrics
	_gad _fg.RWMutex
}

func (_cce *TtfType) String() string {
	return _f.Sprintf("\u0046\u004fN\u0054\u005f\u0046\u0049\u004cE\u0032\u007b\u0025\u0023\u0071 \u0055\u006e\u0069\u0074\u0073\u0050\u0065\u0072\u0045\u006d\u003d\u0025\u0064\u0020\u0042\u006f\u006c\u0064\u003d\u0025\u0074\u0020\u0049\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065\u003d\u0025\u0066\u0020"+"\u0043\u0061pH\u0065\u0069\u0067h\u0074\u003d\u0025\u0064 Ch\u0061rs\u003d\u0025\u0064\u0020\u0047\u006c\u0079ph\u004e\u0061\u006d\u0065\u0073\u003d\u0025d\u007d", _cce.PostScriptName, _cce.UnitsPerEm, _cce.Bold, _cce.ItalicAngle, _cce.CapHeight, len(_cce.Chars), len(_cce.GlyphNames))
}
func init() {
	RegisterStdFont(CourierName, _acd, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u0043\u006f\u0075\u0072\u0069e\u0072\u004e\u0065\u0077", "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077")
	RegisterStdFont(CourierBoldName, _eg, "\u0043o\u0075r\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(CourierObliqueName, _beb, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(CourierBoldObliqueName, _dcg, "C\u006f\u0075\u0072\u0069er\u004ee\u0077\u002c\u0042\u006f\u006cd\u0049\u0074\u0061\u006c\u0069\u0063")
}

const (
	SymbolName       = StdFontName("\u0053\u0079\u006d\u0062\u006f\u006c")
	ZapfDingbatsName = StdFontName("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
)

var _dcgd = []GlyphName{"\u002en\u006f\u0074\u0064\u0065\u0066", "\u002e\u006e\u0075l\u006c", "\u006e\u006fn\u006d\u0061\u0072k\u0069\u006e\u0067\u0072\u0065\u0074\u0075\u0072\u006e", "\u0073\u0070\u0061c\u0065", "\u0065\u0078\u0063\u006c\u0061\u006d", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c", "\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u0069\u0067\u006e", "\u0064\u006f\u006c\u006c\u0061\u0072", "\u0070e\u0072\u0063\u0065\u006e\u0074", "\u0061m\u0070\u0065\u0072\u0073\u0061\u006ed", "q\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067\u006c\u0065", "\u0070a\u0072\u0065\u006e\u006c\u0065\u0066t", "\u0070\u0061\u0072\u0065\u006e\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0074\u0065\u0072\u0069\u0073\u006b", "\u0070\u006c\u0075\u0073", "\u0063\u006f\u006dm\u0061", "\u0068\u0079\u0070\u0068\u0065\u006e", "\u0070\u0065\u0072\u0069\u006f\u0064", "\u0073\u006c\u0061s\u0068", "\u007a\u0065\u0072\u006f", "\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0063\u006f\u006co\u006e", "\u0073e\u006d\u0069\u0063\u006f\u006c\u006fn", "\u006c\u0065\u0073\u0073", "\u0065\u0071\u0075a\u006c", "\u0067r\u0065\u0061\u0074\u0065\u0072", "\u0071\u0075\u0065\u0073\u0074\u0069\u006f\u006e", "\u0061\u0074", "\u0041", "\u0042", "\u0043", "\u0044", "\u0045", "\u0046", "\u0047", "\u0048", "\u0049", "\u004a", "\u004b", "\u004c", "\u004d", "\u004e", "\u004f", "\u0050", "\u0051", "\u0052", "\u0053", "\u0054", "\u0055", "\u0056", "\u0057", "\u0058", "\u0059", "\u005a", "b\u0072\u0061\u0063\u006b\u0065\u0074\u006c\u0065\u0066\u0074", "\u0062a\u0063\u006b\u0073\u006c\u0061\u0073h", "\u0062\u0072\u0061c\u006b\u0065\u0074\u0072\u0069\u0067\u0068\u0074", "a\u0073\u0063\u0069\u0069\u0063\u0069\u0072\u0063\u0075\u006d", "\u0075\u006e\u0064\u0065\u0072\u0073\u0063\u006f\u0072\u0065", "\u0067\u0072\u0061v\u0065", "\u0061", "\u0062", "\u0063", "\u0064", "\u0065", "\u0066", "\u0067", "\u0068", "\u0069", "\u006a", "\u006b", "\u006c", "\u006d", "\u006e", "\u006f", "\u0070", "\u0071", "\u0072", "\u0073", "\u0074", "\u0075", "\u0076", "\u0077", "\u0078", "\u0079", "\u007a", "\u0062r\u0061\u0063\u0065\u006c\u0065\u0066t", "\u0062\u0061\u0072", "\u0062\u0072\u0061\u0063\u0065\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0063\u0069\u0069\u0074\u0069\u006c\u0064\u0065", "\u0041d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0041\u0072\u0069n\u0067", "\u0043\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0045\u0061\u0063\u0075\u0074\u0065", "\u004e\u0074\u0069\u006c\u0064\u0065", "\u004fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0055d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0061\u0063\u0075\u0074\u0065", "\u0061\u0067\u0072\u0061\u0076\u0065", "a\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0074\u0069\u006c\u0064\u0065", "\u0061\u0072\u0069n\u0067", "\u0063\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0065\u0061\u0063\u0075\u0074\u0065", "\u0065\u0067\u0072\u0061\u0076\u0065", "e\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0065d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0069\u0061\u0063\u0075\u0074\u0065", "\u0069\u0067\u0072\u0061\u0076\u0065", "i\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0069d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006e\u0074\u0069\u006c\u0064\u0065", "\u006f\u0061\u0063\u0075\u0074\u0065", "\u006f\u0067\u0072\u0061\u0076\u0065", "o\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u006fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006f\u0074\u0069\u006c\u0064\u0065", "\u0075\u0061\u0063\u0075\u0074\u0065", "\u0075\u0067\u0072\u0061\u0076\u0065", "u\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0075d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0064\u0061\u0067\u0067\u0065\u0072", "\u0064\u0065\u0067\u0072\u0065\u0065", "\u0063\u0065\u006e\u0074", "\u0073\u0074\u0065\u0072\u006c\u0069\u006e\u0067", "\u0073e\u0063\u0074\u0069\u006f\u006e", "\u0062\u0075\u006c\u006c\u0065\u0074", "\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h", "\u0067\u0065\u0072\u006d\u0061\u006e\u0064\u0062\u006c\u0073", "\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0065\u0064", "\u0063o\u0070\u0079\u0072\u0069\u0067\u0068t", "\u0074r\u0061\u0064\u0065\u006d\u0061\u0072k", "\u0061\u0063\u0075t\u0065", "\u0064\u0069\u0065\u0072\u0065\u0073\u0069\u0073", "\u006e\u006f\u0074\u0065\u0071\u0075\u0061\u006c", "\u0041\u0045", "\u004f\u0073\u006c\u0061\u0073\u0068", "\u0069\u006e\u0066\u0069\u006e\u0069\u0074\u0079", "\u0070l\u0075\u0073\u006d\u0069\u006e\u0075s", "\u006ce\u0073\u0073\u0065\u0071\u0075\u0061l", "\u0067\u0072\u0065a\u0074\u0065\u0072\u0065\u0071\u0075\u0061\u006c", "\u0079\u0065\u006e", "\u006d\u0075", "p\u0061\u0072\u0074\u0069\u0061\u006c\u0064\u0069\u0066\u0066", "\u0073u\u006d\u006d\u0061\u0074\u0069\u006fn", "\u0070r\u006f\u0064\u0075\u0063\u0074", "\u0070\u0069", "\u0069\u006e\u0074\u0065\u0067\u0072\u0061\u006c", "o\u0072\u0064\u0066\u0065\u006d\u0069\u006e\u0069\u006e\u0065", "\u006f\u0072\u0064m\u0061\u0073\u0063\u0075\u006c\u0069\u006e\u0065", "\u004f\u006d\u0065g\u0061", "\u0061\u0065", "\u006f\u0073\u006c\u0061\u0073\u0068", "\u0071\u0075\u0065s\u0074\u0069\u006f\u006e\u0064\u006f\u0077\u006e", "\u0065\u0078\u0063\u006c\u0061\u006d\u0064\u006f\u0077\u006e", "\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u006e\u006f\u0074", "\u0072a\u0064\u0069\u0063\u0061\u006c", "\u0066\u006c\u006f\u0072\u0069\u006e", "a\u0070\u0070\u0072\u006f\u0078\u0065\u0071\u0075\u0061\u006c", "\u0044\u0065\u006ct\u0061", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074r\u0069\u0067\u0068\u0074", "\u0065\u006c\u006c\u0069\u0070\u0073\u0069\u0073", "\u006e\u006fn\u0062\u0072\u0065a\u006b\u0069\u006e\u0067\u0073\u0070\u0061\u0063\u0065", "\u0041\u0067\u0072\u0061\u0076\u0065", "\u0041\u0074\u0069\u006c\u0064\u0065", "\u004f\u0074\u0069\u006c\u0064\u0065", "\u004f\u0045", "\u006f\u0065", "\u0065\u006e\u0064\u0061\u0073\u0068", "\u0065\u006d\u0064\u0061\u0073\u0068", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u006c\u0065\u0066\u0074", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c\u0072\u0069\u0067\u0068\u0074", "\u0071u\u006f\u0074\u0065\u006c\u0065\u0066t", "\u0071\u0075\u006f\u0074\u0065\u0072\u0069\u0067\u0068\u0074", "\u0064\u0069\u0076\u0069\u0064\u0065", "\u006co\u007a\u0065\u006e\u0067\u0065", "\u0079d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0059d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e", "\u0063\u0075\u0072\u0072\u0065\u006e\u0063\u0079", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006c\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006cr\u0069\u0067\u0068\u0074", "\u0066\u0069", "\u0066\u006c", "\u0064a\u0067\u0067\u0065\u0072\u0064\u0062l", "\u0070\u0065\u0072\u0069\u006f\u0064\u0063\u0065\u006et\u0065\u0072\u0065\u0064", "\u0071\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067l\u0062\u0061\u0073\u0065", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u0062\u0061\u0073\u0065", "p\u0065\u0072\u0074\u0068\u006f\u0075\u0073\u0061\u006e\u0064", "A\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "E\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0041\u0061\u0063\u0075\u0074\u0065", "\u0045d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0045\u0067\u0072\u0061\u0076\u0065", "\u0049\u0061\u0063\u0075\u0074\u0065", "I\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0049d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0049\u0067\u0072\u0061\u0076\u0065", "\u004f\u0061\u0063\u0075\u0074\u0065", "O\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061\u0070\u0070l\u0065", "\u004f\u0067\u0072\u0061\u0076\u0065", "\u0055\u0061\u0063\u0075\u0074\u0065", "U\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0055\u0067\u0072\u0061\u0076\u0065", "\u0064\u006f\u0074\u006c\u0065\u0073\u0073\u0069", "\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0074\u0069\u006cd\u0065", "\u006d\u0061\u0063\u0072\u006f\u006e", "\u0062\u0072\u0065v\u0065", "\u0064o\u0074\u0061\u0063\u0063\u0065\u006et", "\u0072\u0069\u006e\u0067", "\u0063e\u0064\u0069\u006c\u006c\u0061", "\u0068\u0075\u006eg\u0061\u0072\u0075\u006d\u006c\u0061\u0075\u0074", "\u006f\u0067\u006f\u006e\u0065\u006b", "\u0063\u0061\u0072o\u006e", "\u004c\u0073\u006c\u0061\u0073\u0068", "\u006c\u0073\u006c\u0061\u0073\u0068", "\u0053\u0063\u0061\u0072\u006f\u006e", "\u0073\u0063\u0061\u0072\u006f\u006e", "\u005a\u0063\u0061\u0072\u006f\u006e", "\u007a\u0063\u0061\u0072\u006f\u006e", "\u0062r\u006f\u006b\u0065\u006e\u0062\u0061r", "\u0045\u0074\u0068", "\u0065\u0074\u0068", "\u0059\u0061\u0063\u0075\u0074\u0065", "\u0079\u0061\u0063\u0075\u0074\u0065", "\u0054\u0068\u006fr\u006e", "\u0074\u0068\u006fr\u006e", "\u006d\u0069\u006eu\u0073", "\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "o\u006e\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "t\u0077\u006f\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u0074\u0068\u0072\u0065\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u006fn\u0065\u0068\u0061\u006c\u0066", "\u006f\u006e\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072", "\u0074\u0068\u0072\u0065\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072\u0073", "\u0066\u0072\u0061n\u0063", "\u0047\u0062\u0072\u0065\u0076\u0065", "\u0067\u0062\u0072\u0065\u0076\u0065", "\u0049\u0064\u006f\u0074\u0061\u0063\u0063\u0065\u006e\u0074", "\u0053\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0073\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0043\u0061\u0063\u0075\u0074\u0065", "\u0063\u0061\u0063\u0075\u0074\u0065", "\u0043\u0063\u0061\u0072\u006f\u006e", "\u0063\u0063\u0061\u0072\u006f\u006e", "\u0064\u0063\u0072\u006f\u0061\u0074"}

func (_cedeg *ttfParser) parseCmapFormat12() error {
	_ggdc := _cedeg.ReadULong()
	_fb.Log.Trace("\u0070\u0061\u0072se\u0043\u006d\u0061\u0070\u0046\u006f\u0072\u006d\u0061t\u00312\u003a \u0025s\u0020\u006e\u0075\u006d\u0047\u0072\u006f\u0075\u0070\u0073\u003d\u0025\u0064", _cedeg._dad.String(), _ggdc)
	for _bggf := uint32(0); _bggf < _ggdc; _bggf++ {
		_cgga := _cedeg.ReadULong()
		_fec := _cedeg.ReadULong()
		_gcf := _cedeg.ReadULong()
		if _cgga > 0x0010FFFF || (0xD800 <= _cgga && _cgga <= 0xDFFF) {
			return _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		if _fec < _cgga || _fec > 0x0010FFFF || (0xD800 <= _fec && _fec <= 0xDFFF) {
			return _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		for _fffb := _cgga; _fffb <= _fec; _fffb++ {
			if _fffb > 0x10FFFF {
				_fb.Log.Debug("\u0046\u006fr\u006d\u0061\u0074\u0020\u0031\u0032\u0020\u0063\u006d\u0061\u0070\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0062\u0065\u0079\u006f\u006e\u0064\u0020\u0055\u0043\u0053\u002d\u0034")
			}
			_cedeg._dad.Chars[rune(_fffb)] = GID(_gcf)
			_gcf++
		}
	}
	return nil
}

var _dfe _fg.Once

func (_gba *ttfParser) Read32Fixed() float64 {
	_cgf := float64(_gba.ReadShort())
	_dgg := float64(_gba.ReadUShort()) / 65536.0
	return _cgf + _dgg
}
func (_deac StdFont) ToPdfObject() _ag.PdfObject {
	_gg := _ag.MakeDict()
	_gg.Set("\u0054\u0079\u0070\u0065", _ag.MakeName("\u0046\u006f\u006e\u0074"))
	_gg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ag.MakeName("\u0054\u0079\u0070e\u0031"))
	_gg.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _ag.MakeName(_deac.Name()))
	_gg.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _deac._da.ToPdfObject())
	return _ag.MakeIndirectObject(_gg)
}
func (_af StdFont) Encoder() _b.TextEncoder { return _af._da }
func _eac() {
	_ega = MakeRuneCharSafeMap(len(_cc))
	_fdg = MakeRuneCharSafeMap(len(_cc))
	_gcb = MakeRuneCharSafeMap(len(_cc))
	_dcd = MakeRuneCharSafeMap(len(_cc))
	for _cbd, _ecd := range _cc {
		_ega.Write(_ecd, CharMetrics{Wx: float64(_ebb[_cbd])})
		_fdg.Write(_ecd, CharMetrics{Wx: float64(_adb[_cbd])})
		_gcb.Write(_ecd, CharMetrics{Wx: float64(_egae[_cbd])})
		_dcd.Write(_ecd, CharMetrics{Wx: float64(_fe[_cbd])})
	}
}

type TtfType struct {
	UnitsPerEm             uint16
	PostScriptName         string
	Bold                   bool
	ItalicAngle            float64
	IsFixedPitch           bool
	TypoAscender           int16
	TypoDescender          int16
	UnderlinePosition      int16
	UnderlineThickness     int16
	Xmin, Ymin, Xmax, Ymax int16
	CapHeight              int16
	Widths                 []uint16
	Chars                  map[rune]GID
	GlyphNames             []GlyphName
}

var _fdg *RuneCharSafeMap

type FontWeight int

func init() {
	RegisterStdFont(HelveticaName, _gaa, "\u0041\u0072\u0069a\u006c")
	RegisterStdFont(HelveticaBoldName, _dfff, "\u0041\u0072\u0069\u0061\u006c\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(HelveticaObliqueName, _gab, "\u0041\u0072\u0069a\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(HelveticaBoldObliqueName, _abc, "\u0041\u0072i\u0061\u006c\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}
func (_afe *TtfType) MakeToUnicode() *_de.CMap {
	_eege := make(map[_de.CharCode]rune)
	if len(_afe.GlyphNames) == 0 {
		for _bcd := range _afe.Chars {
			_eege[_de.CharCode(_bcd)] = _bcd
		}
		return _de.NewToUnicodeCMap(_eege)
	}
	for _adf, _ddeg := range _afe.Chars {
		_cfc := _de.CharCode(_adf)
		_gda := _afe.GlyphNames[_ddeg]
		_dbb, _dbc := _b.GlyphToRune(_gda)
		if !_dbc {
			_fb.Log.Debug("\u004e\u006f \u0072\u0075\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0067\u006c\u0079\u0070h=\u0025\u0071", _adf, _gda)
			_dbb = _b.MissingCodeRune
		}
		_eege[_cfc] = _dbb
	}
	return _de.NewToUnicodeCMap(_eege)
}
func IsStdFont(name StdFontName) bool { _, _fbe := _dfb.read(name); return _fbe }
func _gaa() StdFont {
	_dfe.Do(_acg)
	_ceb := Descriptor{Name: HelveticaName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0020, BBox: [4]float64{-166, -225, 1000, 931}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_ceb, _bgd)
}
func _dcg() StdFont {
	_ddea.Do(_gff)
	_gca := Descriptor{Name: CourierBoldObliqueName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0061, BBox: [4]float64{-57, -250, 869, 801}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_gca, _acde)
}

type StdFont struct {
	_bgg Descriptor
	_dc  *RuneCharSafeMap
	_da  _b.TextEncoder
}

func (_fgge *TtfType) MakeEncoder() (_b.SimpleEncoder, error) {
	_eag := make(map[_b.CharCode]GlyphName)
	for _gdc := _b.CharCode(0); _gdc <= 256; _gdc++ {
		_edf := rune(_gdc)
		_aagb, _gdb := _fgge.Chars[_edf]
		if !_gdb {
			continue
		}
		var _ebc GlyphName
		if int(_aagb) >= 0 && int(_aagb) < len(_fgge.GlyphNames) {
			_ebc = _fgge.GlyphNames[_aagb]
		} else {
			_gfe := rune(_aagb)
			if _cgd, _edc := _b.RuneToGlyph(_gfe); _edc {
				_ebc = _cgd
			}
		}
		if _ebc != "" {
			_eag[_gdc] = _ebc
		}
	}
	if len(_eag) == 0 {
		_fb.Log.Debug("WA\u0052\u004eI\u004e\u0047\u003a\u0020\u005a\u0065\u0072\u006f\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002e\u0020\u0074\u0074\u0066=\u0025s\u0020\u0043\u0068\u0061\u0072\u0073\u003d\u005b%\u00200\u0032\u0078]", _fgge, _fgge.Chars)
	}
	return _b.NewCustomSimpleTextEncoder(_eag, nil)
}
func (_ecf *ttfParser) ParseHead() error {
	if _dbe := _ecf.Seek("\u0068\u0065\u0061\u0064"); _dbe != nil {
		return _dbe
	}
	_ecf.Skip(3 * 4)
	_eea := _ecf.ReadULong()
	if _eea != 0x5F0F3CF5 {
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0063\u006fr\u0072e\u0063\u0074\u0020\u006d\u0061\u0067\u0069\u0063\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e\u0020\u0046\u006fn\u0074\u0020\u006d\u0061\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0069\u0073\u0070\u006c\u0061\u0079\u0020\u0063\u006f\u0072\u0072\u0065\u0063t\u006c\u0079\u002e\u0020\u0025\u0073", _ecf)
	}
	_ecf.Skip(2)
	_ecf._dad.UnitsPerEm = _ecf.ReadUShort()
	_ecf.Skip(2 * 8)
	_ecf._dad.Xmin = _ecf.ReadShort()
	_ecf._dad.Ymin = _ecf.ReadShort()
	_ecf._dad.Xmax = _ecf.ReadShort()
	_ecf._dad.Ymax = _ecf.ReadShort()
	return nil
}

var _bge *RuneCharSafeMap

func (_acdg *ttfParser) ParsePost() error {
	if _gae := _acdg.Seek("\u0070\u006f\u0073\u0074"); _gae != nil {
		return _gae
	}
	_feee := _acdg.Read32Fixed()
	_acdg._dad.ItalicAngle = _acdg.Read32Fixed()
	_acdg._dad.UnderlinePosition = _acdg.ReadShort()
	_acdg._dad.UnderlineThickness = _acdg.ReadShort()
	_acdg._dad.IsFixedPitch = _acdg.ReadULong() != 0
	_acdg.ReadULong()
	_acdg.ReadULong()
	_acdg.ReadULong()
	_acdg.ReadULong()
	_fb.Log.Trace("\u0050a\u0072\u0073\u0065\u0050\u006f\u0073\u0074\u003a\u0020\u0066\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065\u003d\u0025\u0066", _feee)
	switch _feee {
	case 1.0:
		_acdg._dad.GlyphNames = _dcgd
	case 2.0:
		_ddf := int(_acdg.ReadUShort())
		_acbe := make([]int, _ddf)
		_acdg._dad.GlyphNames = make([]GlyphName, _ddf)
		_eba := -1
		for _gge := 0; _gge < _ddf; _gge++ {
			_ged := int(_acdg.ReadUShort())
			_acbe[_gge] = _ged
			if _ged <= 0x7fff && _ged > _eba {
				_eba = _ged
			}
		}
		var _geg []GlyphName
		if _eba >= len(_dcgd) {
			_geg = make([]GlyphName, _eba-len(_dcgd)+1)
			for _aaa := 0; _aaa < _eba-len(_dcgd)+1; _aaa++ {
				_gfed := int(_acdg.readByte())
				_eed, _dbf := _acdg.ReadStr(_gfed)
				if _dbf != nil {
					return _dbf
				}
				_geg[_aaa] = GlyphName(_eed)
			}
		}
		for _gbbd := 0; _gbbd < _ddf; _gbbd++ {
			_gcfb := _acbe[_gbbd]
			if _gcfb < len(_dcgd) {
				_acdg._dad.GlyphNames[_gbbd] = _dcgd[_gcfb]
			} else if _gcfb >= len(_dcgd) && _gcfb <= 32767 {
				_acdg._dad.GlyphNames[_gbbd] = _geg[_gcfb-len(_dcgd)]
			} else {
				_acdg._dad.GlyphNames[_gbbd] = "\u002e\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064"
			}
		}
	case 2.5:
		_ccb := make([]int, _acdg._bff)
		for _ddgc := 0; _ddgc < len(_ccb); _ddgc++ {
			_aaag := int(_acdg.ReadSByte())
			_ccb[_ddgc] = _ddgc + 1 + _aaag
		}
		_acdg._dad.GlyphNames = make([]GlyphName, len(_ccb))
		for _ca := 0; _ca < len(_acdg._dad.GlyphNames); _ca++ {
			_dead := _dcgd[_ccb[_ca]]
			_acdg._dad.GlyphNames[_ca] = _dead
		}
	case 3.0:
		_fb.Log.Debug("\u004e\u006f\u0020\u0050\u006f\u0073t\u0053\u0063\u0072i\u0070\u0074\u0020n\u0061\u006d\u0065\u0020\u0069\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020is\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e")
	default:
		_fb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020f\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065=\u0025\u0066", _feee)
	}
	return nil
}
func (_cd *RuneCharSafeMap) Copy() *RuneCharSafeMap {
	_ec := MakeRuneCharSafeMap(_cd.Length())
	_cd.Range(func(_bf rune, _dea CharMetrics) (_dgc bool) { _ec._ga[_bf] = _dea; return false })
	return _ec
}
func (_baf StdFont) Descriptor() Descriptor { return _baf._bgg }
func (_dega *ttfParser) ReadStr(length int) (string, error) {
	_cgag := make([]byte, length)
	_dgb, _fbc := _dega._gbca.Read(_cgag)
	if _fbc != nil {
		return "", _fbc
	} else if _dgb != length {
		return "", _f.Errorf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073", length)
	}
	return string(_cgag), nil
}

var _bgd *RuneCharSafeMap

func init() {
	RegisterStdFont(SymbolName, _egf, "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063", "S\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064", "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(ZapfDingbatsName, _ccc)
}

type Font interface {
	Encoder() _b.TextEncoder
	GetRuneMetrics(_df rune) (CharMetrics, bool)
}
type StdFontName string

var _adb = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 611, 778, 778, 778, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 778, 778, 667, 667, 667, 667, 667, 944, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 611, 778, 722, 722, 722, 722, 556, 556, 556, 556, 556, 667, 667, 667, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 1000, 722, 722, 722, 722, 667, 667, 667, 667, 500, 500, 500, 500, 333, 500, 722, 500, 500, 833, 500, 500, 581, 520, 500, 930, 500, 556, 278, 220, 394, 394, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 556, 500, 500, 672, 556, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 556, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 333, 556, 556, 278, 278, 394, 278, 570, 549, 570, 494, 278, 833, 333, 570, 556, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 300, 330, 500, 500, 556, 540, 333, 333, 494, 1000, 250, 250, 1000, 570, 570, 556, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 444, 444, 549, 444, 444, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 333, 416, 333, 556, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func (_bcaa *ttfParser) ParseHhea() error {
	if _fae := _bcaa.Seek("\u0068\u0068\u0065\u0061"); _fae != nil {
		return _fae
	}
	_bcaa.Skip(4 + 15*2)
	_bcaa._bgda = _bcaa.ReadUShort()
	return nil
}

type CharMetrics struct {
	Wx float64
	Wy float64
}

func init() {
	RegisterStdFont(TimesRomanName, _aag, "\u0054\u0069\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006de\u0073")
	RegisterStdFont(TimesBoldName, _ddb, "\u0054i\u006de\u0073\u004e\u0065\u0077\u0052o\u006d\u0061n\u002c\u0042\u006f\u006c\u0064", "\u0054\u0069\u006d\u0065\u0073\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(TimesItalicName, _ggd, "T\u0069m\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006da\u006e\u002c\u0049\u0074al\u0069\u0063", "\u0054\u0069\u006de\u0073\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(TimesBoldItalicName, _eeb, "\u0054i\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e,\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}
func (_fgc *ttfParser) ParseOS2() error {
	if _ccg := _fgc.Seek("\u004f\u0053\u002f\u0032"); _ccg != nil {
		return _ccg
	}
	_cebg := _fgc.ReadUShort()
	_fgc.Skip(4 * 2)
	_fgc.Skip(11*2 + 10 + 4*4 + 4)
	_ebbc := _fgc.ReadUShort()
	_fgc._dad.Bold = (_ebbc & 32) != 0
	_fgc.Skip(2 * 2)
	_fgc._dad.TypoAscender = _fgc.ReadShort()
	_fgc._dad.TypoDescender = _fgc.ReadShort()
	if _cebg >= 2 {
		_fgc.Skip(3*2 + 2*4 + 2)
		_fgc._dad.CapHeight = _fgc.ReadShort()
	} else {
		_fgc._dad.CapHeight = 0
	}
	return nil
}

var _ega *RuneCharSafeMap

type GID = _b.GID

func _acg() {
	_bgd = MakeRuneCharSafeMap(len(_cc))
	_bge = MakeRuneCharSafeMap(len(_cc))
	for _fdc, _fgf := range _cc {
		_bgd.Write(_fgf, CharMetrics{Wx: float64(_eae[_fdc])})
		_bge.Write(_fgf, CharMetrics{Wx: float64(_ddeb[_fdc])})
	}
	_gbc = _bgd.Copy()
	_baed = _bge.Copy()
}
func (_efae *ttfParser) parseCmapSubtable31(_gadc int64) error {
	_bac := make([]rune, 0, 8)
	_aecb := make([]rune, 0, 8)
	_dfga := make([]int16, 0, 8)
	_cbc := make([]uint16, 0, 8)
	_efae._dad.Chars = make(map[rune]GID)
	_efae._gbca.Seek(int64(_efae._ggc["\u0063\u006d\u0061\u0070"])+_gadc, _e.SeekStart)
	_cccf := _efae.ReadUShort()
	if _cccf != 4 {
		_fb.Log.Debug("u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0073\u0075\u0062t\u0061\u0062\u006c\u0065\u0020\u0066\u006fr\u006d\u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0028\u0025w\u0029", _cccf)
		return nil
	}
	_efae.Skip(2 * 2)
	_dcga := int(_efae.ReadUShort() / 2)
	_efae.Skip(3 * 2)
	for _bcfd := 0; _bcfd < _dcga; _bcfd++ {
		_aecb = append(_aecb, rune(_efae.ReadUShort()))
	}
	_efae.Skip(2)
	for _abeb := 0; _abeb < _dcga; _abeb++ {
		_bac = append(_bac, rune(_efae.ReadUShort()))
	}
	for _bggc := 0; _bggc < _dcga; _bggc++ {
		_dfga = append(_dfga, _efae.ReadShort())
	}
	_afgb, _ := _efae._gbca.Seek(int64(0), _e.SeekCurrent)
	for _gee := 0; _gee < _dcga; _gee++ {
		_cbc = append(_cbc, _efae.ReadUShort())
	}
	for _cfcb := 0; _cfcb < _dcga; _cfcb++ {
		_gbe := _bac[_cfcb]
		_aab := _aecb[_cfcb]
		_cfd := _dfga[_cfcb]
		_acb := _cbc[_cfcb]
		if _acb > 0 {
			_efae._gbca.Seek(_afgb+2*int64(_cfcb)+int64(_acb), _e.SeekStart)
		}
		for _fgbg := _gbe; _fgbg <= _aab; _fgbg++ {
			if _fgbg == 0xFFFF {
				break
			}
			var _bcda int32
			if _acb > 0 {
				_bcda = int32(_efae.ReadUShort())
				if _bcda > 0 {
					_bcda += int32(_cfd)
				}
			} else {
				_bcda = _fgbg + int32(_cfd)
			}
			if _bcda >= 65536 {
				_bcda -= 65536
			}
			if _bcda > 0 {
				_efae._dad.Chars[_fgbg] = GID(_bcda)
			}
		}
	}
	return nil
}

var _cc = []rune{'A', 'Æ', 'Á', 'Ă', 'Â', 'Ä', 'À', 'Ā', 'Ą', 'Å', 'Ã', 'B', 'C', 'Ć', 'Č', 'Ç', 'D', 'Ď', 'Đ', '∆', 'E', 'É', 'Ě', 'Ê', 'Ë', 'Ė', 'È', 'Ē', 'Ę', 'Ð', '€', 'F', 'G', 'Ğ', 'Ģ', 'H', 'I', 'Í', 'Î', 'Ï', 'İ', 'Ì', 'Ī', 'Į', 'J', 'K', 'Ķ', 'L', 'Ĺ', 'Ľ', 'Ļ', 'Ł', 'M', 'N', 'Ń', 'Ň', 'Ņ', 'Ñ', 'O', 'Œ', 'Ó', 'Ô', 'Ö', 'Ò', 'Ő', 'Ō', 'Ø', 'Õ', 'P', 'Q', 'R', 'Ŕ', 'Ř', 'Ŗ', 'S', 'Ś', 'Š', 'Ş', 'Ș', 'T', 'Ť', 'Ţ', 'Þ', 'U', 'Ú', 'Û', 'Ü', 'Ù', 'Ű', 'Ū', 'Ų', 'Ů', 'V', 'W', 'X', 'Y', 'Ý', 'Ÿ', 'Z', 'Ź', 'Ž', 'Ż', 'a', 'á', 'ă', 'â', '´', 'ä', 'æ', 'à', 'ā', '&', 'ą', 'å', '^', '~', '*', '@', 'ã', 'b', '\\', '|', '{', '}', '[', ']', '˘', '¦', '•', 'c', 'ć', 'ˇ', 'č', 'ç', '¸', '¢', 'ˆ', ':', ',', '\uf6c3', '©', '¤', 'd', '†', '‡', 'ď', 'đ', '°', '¨', '÷', '$', '˙', 'ı', 'e', 'é', 'ě', 'ê', 'ë', 'ė', 'è', '8', '…', 'ē', '—', '–', 'ę', '=', 'ð', '!', '¡', 'f', 'ﬁ', '5', 'ﬂ', 'ƒ', '4', '⁄', 'g', 'ğ', 'ģ', 'ß', '`', '>', '≥', '«', '»', '‹', '›', 'h', '˝', '-', 'i', 'í', 'î', 'ï', 'ì', 'ī', 'į', 'j', 'k', 'ķ', 'l', 'ĺ', 'ľ', 'ļ', '<', '≤', '¬', '◊', 'ł', 'm', '¯', '−', 'µ', '×', 'n', 'ń', 'ň', 'ņ', '9', '≠', 'ñ', '#', 'o', 'ó', 'ô', 'ö', 'œ', '˛', 'ò', 'ő', 'ō', '1', '½', '¼', '¹', 'ª', 'º', 'ø', 'õ', 'p', '¶', '(', ')', '∂', '%', '.', '·', '‰', '+', '±', 'q', '?', '¿', '"', '„', '“', '”', '‘', '’', '‚', '\'', 'r', 'ŕ', '√', 'ř', 'ŗ', '®', '˚', 's', 'ś', 'š', 'ş', 'ș', '§', ';', '7', '6', '/', ' ', '£', '∑', 't', 'ť', 'ţ', 'þ', '3', '¾', '³', '˜', '™', '2', '²', 'u', 'ú', 'û', 'ü', 'ù', 'ű', 'ū', '_', 'ų', 'ů', 'v', 'w', 'x', 'y', 'ý', 'ÿ', '¥', 'z', 'ź', 'ž', 'ż', '0'}

func (_acaa *ttfParser) ReadUShort() (_abf uint16) {
	_dg.Read(_acaa._gbca, _dg.BigEndian, &_abf)
	return _abf
}

var _ Font = StdFont{}

func TtfParseFile(fileStr string) (TtfType, error) {
	_adba, _daa := _d.Open(fileStr)
	if _daa != nil {
		return TtfType{}, _daa
	}
	defer _adba.Close()
	return TtfParse(_adba)
}

var _ef *RuneCharSafeMap

func (_fd *RuneCharSafeMap) Length() int {
	_fd._gad.RLock()
	defer _fd._gad.RUnlock()
	return len(_fd._ga)
}
func _ddb() StdFont {
	_dagc.Do(_eac)
	_ddg := Descriptor{Name: TimesBoldName, Family: _bef, Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 935}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 676, XHeight: 461, StemV: 139, StemH: 44}
	return NewStdFont(_ddg, _fdg)
}

var _abd *RuneCharSafeMap

func _aag() StdFont {
	_dagc.Do(_eac)
	_gabe := Descriptor{Name: TimesRomanName, Family: _bef, Weight: FontWeightRoman, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 898}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 662, XHeight: 450, StemV: 84, StemH: 28}
	return NewStdFont(_gabe, _ega)
}
func (_fde *ttfParser) parseCmapFormat6() error {
	_abde := int(_fde.ReadUShort())
	_afeb := int(_fde.ReadUShort())
	_fb.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u0046o\u0072\u006d\u0061\u0074\u0036\u003a\u0020%s\u0020\u0066\u0069\u0072s\u0074\u0043\u006f\u0064\u0065\u003d\u0025\u0064\u0020en\u0074\u0072y\u0043\u006f\u0075\u006e\u0074\u003d\u0025\u0064", _fde._dad.String(), _abde, _afeb)
	for _efea := 0; _efea < _afeb; _efea++ {
		_ddc := GID(_fde.ReadUShort())
		_fde._dad.Chars[rune(_efea+_abde)] = _ddc
	}
	return nil
}

const (
	HelveticaName            = StdFontName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	HelveticaBoldName        = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064")
	HelveticaObliqueName     = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065")
	HelveticaBoldObliqueName = StdFontName("H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065")
)

func (_eaa StdFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_ce, _cgg := _eaa._dc.Read(r)
	return _ce, _cgg
}
func (_ggda *ttfParser) ReadShort() (_cab int16) {
	_dg.Read(_ggda._gbca, _dg.BigEndian, &_cab)
	return _cab
}

var _gcd = &RuneCharSafeMap{_ga: map[rune]CharMetrics{' ': {Wx: 250}, '!': {Wx: 333}, '#': {Wx: 500}, '%': {Wx: 833}, '&': {Wx: 778}, '(': {Wx: 333}, ')': {Wx: 333}, '+': {Wx: 549}, ',': {Wx: 250}, '.': {Wx: 250}, '/': {Wx: 278}, '0': {Wx: 500}, '1': {Wx: 500}, '2': {Wx: 500}, '3': {Wx: 500}, '4': {Wx: 500}, '5': {Wx: 500}, '6': {Wx: 500}, '7': {Wx: 500}, '8': {Wx: 500}, '9': {Wx: 500}, ':': {Wx: 278}, ';': {Wx: 278}, '<': {Wx: 549}, '=': {Wx: 549}, '>': {Wx: 549}, '?': {Wx: 444}, '[': {Wx: 333}, ']': {Wx: 333}, '_': {Wx: 500}, '{': {Wx: 480}, '|': {Wx: 200}, '}': {Wx: 480}, '¬': {Wx: 713}, '°': {Wx: 400}, '±': {Wx: 549}, 'µ': {Wx: 576}, '×': {Wx: 549}, '÷': {Wx: 549}, 'ƒ': {Wx: 500}, 'Α': {Wx: 722}, 'Β': {Wx: 667}, 'Γ': {Wx: 603}, 'Ε': {Wx: 611}, 'Ζ': {Wx: 611}, 'Η': {Wx: 722}, 'Θ': {Wx: 741}, 'Ι': {Wx: 333}, 'Κ': {Wx: 722}, 'Λ': {Wx: 686}, 'Μ': {Wx: 889}, 'Ν': {Wx: 722}, 'Ξ': {Wx: 645}, 'Ο': {Wx: 722}, 'Π': {Wx: 768}, 'Ρ': {Wx: 556}, 'Σ': {Wx: 592}, 'Τ': {Wx: 611}, 'Υ': {Wx: 690}, 'Φ': {Wx: 763}, 'Χ': {Wx: 722}, 'Ψ': {Wx: 795}, 'α': {Wx: 631}, 'β': {Wx: 549}, 'γ': {Wx: 411}, 'δ': {Wx: 494}, 'ε': {Wx: 439}, 'ζ': {Wx: 494}, 'η': {Wx: 603}, 'θ': {Wx: 521}, 'ι': {Wx: 329}, 'κ': {Wx: 549}, 'λ': {Wx: 549}, 'ν': {Wx: 521}, 'ξ': {Wx: 493}, 'ο': {Wx: 549}, 'π': {Wx: 549}, 'ρ': {Wx: 549}, 'ς': {Wx: 439}, 'σ': {Wx: 603}, 'τ': {Wx: 439}, 'υ': {Wx: 576}, 'φ': {Wx: 521}, 'χ': {Wx: 549}, 'ψ': {Wx: 686}, 'ω': {Wx: 686}, 'ϑ': {Wx: 631}, 'ϒ': {Wx: 620}, 'ϕ': {Wx: 603}, 'ϖ': {Wx: 713}, '•': {Wx: 460}, '…': {Wx: 1000}, '′': {Wx: 247}, '″': {Wx: 411}, '⁄': {Wx: 167}, '€': {Wx: 750}, 'ℑ': {Wx: 686}, '℘': {Wx: 987}, 'ℜ': {Wx: 795}, 'Ω': {Wx: 768}, 'ℵ': {Wx: 823}, '←': {Wx: 987}, '↑': {Wx: 603}, '→': {Wx: 987}, '↓': {Wx: 603}, '↔': {Wx: 1042}, '↵': {Wx: 658}, '⇐': {Wx: 987}, '⇑': {Wx: 603}, '⇒': {Wx: 987}, '⇓': {Wx: 603}, '⇔': {Wx: 1042}, '∀': {Wx: 713}, '∂': {Wx: 494}, '∃': {Wx: 549}, '∅': {Wx: 823}, '∆': {Wx: 612}, '∇': {Wx: 713}, '∈': {Wx: 713}, '∉': {Wx: 713}, '∋': {Wx: 439}, '∏': {Wx: 823}, '∑': {Wx: 713}, '−': {Wx: 549}, '∗': {Wx: 500}, '√': {Wx: 549}, '∝': {Wx: 713}, '∞': {Wx: 713}, '∠': {Wx: 768}, '∧': {Wx: 603}, '∨': {Wx: 603}, '∩': {Wx: 768}, '∪': {Wx: 768}, '∫': {Wx: 274}, '∴': {Wx: 863}, '∼': {Wx: 549}, '≅': {Wx: 549}, '≈': {Wx: 549}, '≠': {Wx: 549}, '≡': {Wx: 549}, '≤': {Wx: 549}, '≥': {Wx: 549}, '⊂': {Wx: 713}, '⊃': {Wx: 713}, '⊄': {Wx: 713}, '⊆': {Wx: 713}, '⊇': {Wx: 713}, '⊕': {Wx: 768}, '⊗': {Wx: 768}, '⊥': {Wx: 658}, '⋅': {Wx: 250}, '⌠': {Wx: 686}, '⌡': {Wx: 686}, '〈': {Wx: 329}, '〉': {Wx: 329}, '◊': {Wx: 494}, '♠': {Wx: 753}, '♣': {Wx: 753}, '♥': {Wx: 753}, '♦': {Wx: 753}, '\uf6d9': {Wx: 790}, '\uf6da': {Wx: 790}, '\uf6db': {Wx: 890}, '\uf8e5': {Wx: 500}, '\uf8e6': {Wx: 603}, '\uf8e7': {Wx: 1000}, '\uf8e8': {Wx: 790}, '\uf8e9': {Wx: 790}, '\uf8ea': {Wx: 786}, '\uf8eb': {Wx: 384}, '\uf8ec': {Wx: 384}, '\uf8ed': {Wx: 384}, '\uf8ee': {Wx: 384}, '\uf8ef': {Wx: 384}, '\uf8f0': {Wx: 384}, '\uf8f1': {Wx: 494}, '\uf8f2': {Wx: 494}, '\uf8f3': {Wx: 494}, '\uf8f4': {Wx: 494}, '\uf8f5': {Wx: 686}, '\uf8f6': {Wx: 384}, '\uf8f7': {Wx: 384}, '\uf8f8': {Wx: 384}, '\uf8f9': {Wx: 384}, '\uf8fa': {Wx: 384}, '\uf8fb': {Wx: 384}, '\uf8fc': {Wx: 494}, '\uf8fd': {Wx: 494}, '\uf8fe': {Wx: 494}, '\uf8ff': {Wx: 790}}}

func (_abb *ttfParser) ParseHmtx() error {
	if _abbf := _abb.Seek("\u0068\u006d\u0074\u0078"); _abbf != nil {
		return _abbf
	}
	_abb._dad.Widths = make([]uint16, 0, 8)
	for _cgb := uint16(0); _cgb < _abb._bgda; _cgb++ {
		_abb._dad.Widths = append(_abb._dad.Widths, _abb.ReadUShort())
		_abb.Skip(2)
	}
	if _abb._bgda < _abb._bff && _abb._bgda > 0 {
		_dfba := _abb._dad.Widths[_abb._bgda-1]
		for _bad := _abb._bgda; _bad < _abb._bff; _bad++ {
			_abb._dad.Widths = append(_abb._dad.Widths, _dfba)
		}
	}
	return nil
}
func (_gc *RuneCharSafeMap) Range(f func(_dgd rune, _dfge CharMetrics) (_ee bool)) {
	_gc._gad.RLock()
	defer _gc._gad.RUnlock()
	for _ac, _cb := range _gc._ga {
		if f(_ac, _cb) {
			break
		}
	}
}
func (_fgb *ttfParser) Parse() (TtfType, error) {
	_fee, _gdg := _fgb.ReadStr(4)
	if _gdg != nil {
		return TtfType{}, _gdg
	}
	if _fee == "\u0074\u0074\u0063\u0066" {
		return _fgb.parseTTC()
	} else if _fee != "\u0000\u0001\u0000\u0000" && _fee != "\u0074\u0072\u0075\u0065" {
		_fb.Log.Debug("\u0055n\u0072\u0065c\u006f\u0067\u006ei\u007a\u0065\u0064\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066i\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u002e\u0020v\u0065\u0072\u0073\u0069\u006f\u006e\u003d\u0025\u0071", _fee)
	}
	_fgdc := int(_fgb.ReadUShort())
	_fgb.Skip(3 * 2)
	_fgb._ggc = make(map[string]uint32)
	var _ade string
	for _aec := 0; _aec < _fgdc; _aec++ {
		_ade, _gdg = _fgb.ReadStr(4)
		if _gdg != nil {
			return TtfType{}, _gdg
		}
		_fgb.Skip(4)
		_daab := _fgb.ReadULong()
		_fgb.Skip(4)
		_fgb._ggc[_ade] = _daab
	}
	_fb.Log.Trace(_gce(_fgb._ggc))
	if _gdg = _fgb.ParseComponents(); _gdg != nil {
		return TtfType{}, _gdg
	}
	return _fgb._dad, nil
}
func (_afa *ttfParser) ParseName() error {
	if _fdfd := _afa.Seek("\u006e\u0061\u006d\u0065"); _fdfd != nil {
		return _fdfd
	}
	_bab, _ := _afa._gbca.Seek(0, _e.SeekCurrent)
	_afa._dad.PostScriptName = ""
	_afa.Skip(2)
	_gdded := _afa.ReadUShort()
	_bfe := _afa.ReadUShort()
	for _bbfc := uint16(0); _bbfc < _gdded && _afa._dad.PostScriptName == ""; _bbfc++ {
		_afa.Skip(3 * 2)
		_afbd := _afa.ReadUShort()
		_bggg := _afa.ReadUShort()
		_faea := _afa.ReadUShort()
		if _afbd == 6 {
			_afa._gbca.Seek(_bab+int64(_bfe)+int64(_faea), _e.SeekStart)
			_cgba, _aefa := _afa.ReadStr(int(_bggg))
			if _aefa != nil {
				return _aefa
			}
			_cgba = _a.Replace(_cgba, "\u0000", "", -1)
			_gbbb, _aefa := _ab.Compile("\u005b\u0028\u0029\u007b\u007d\u003c\u003e\u0020\u002f%\u005b\u005c\u005d\u005d")
			if _aefa != nil {
				return _aefa
			}
			_afa._dad.PostScriptName = _gbbb.ReplaceAllString(_cgba, "")
		}
	}
	if _afa._dad.PostScriptName == "" {
		_fb.Log.Debug("\u0050a\u0072\u0073e\u004e\u0061\u006de\u003a\u0020\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u0050\u006f\u0073t\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0077\u0061\u0073\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	}
	return nil
}
func (_dff CharMetrics) String() string {
	return _f.Sprintf("<\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u003e", _dff.Wx, _dff.Wy)
}
func TtfParse(r _e.ReadSeeker) (TtfType, error) { _ced := &ttfParser{_gbca: r}; return _ced.Parse() }
func (_cede *ttfParser) ParseCmap() error {
	var _cgdd int64
	if _bdf := _cede.Seek("\u0063\u006d\u0061\u0070"); _bdf != nil {
		return _bdf
	}
	_cede.ReadUShort()
	_gcde := int(_cede.ReadUShort())
	_fffd := int64(0)
	_ebcf := int64(0)
	_efe := int64(0)
	for _aff := 0; _aff < _gcde; _aff++ {
		_ecgd := _cede.ReadUShort()
		_dab := _cede.ReadUShort()
		_cgdd = int64(_cede.ReadULong())
		if _ecgd == 3 && _dab == 1 {
			_ebcf = _cgdd
		} else if _ecgd == 3 && _dab == 10 {
			_efe = _cgdd
		} else if _ecgd == 1 && _dab == 0 {
			_fffd = _cgdd
		}
	}
	if _fffd != 0 {
		if _afc := _cede.parseCmapVersion(_fffd); _afc != nil {
			return _afc
		}
	}
	if _ebcf != 0 {
		if _geb := _cede.parseCmapSubtable31(_ebcf); _geb != nil {
			return _geb
		}
	}
	if _efe != 0 {
		if _abbg := _cede.parseCmapVersion(_efe); _abbg != nil {
			return _abbg
		}
	}
	if _ebcf == 0 && _fffd == 0 && _efe == 0 {
		_fb.Log.Debug("\u0074\u0074\u0066P\u0061\u0072\u0073\u0065\u0072\u002e\u0050\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u002e\u0020\u004e\u006f\u0020\u0033\u0031\u002c\u0020\u0031\u0030\u002c\u0020\u00331\u0030\u0020\u0074\u0061\u0062\u006c\u0065\u002e")
	}
	return nil
}
func _eg() StdFont {
	_ddea.Do(_gff)
	_ad := Descriptor{Name: CourierBoldName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0021, BBox: [4]float64{-113, -250, 749, 801}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_ad, _gdd)
}
func _gff() {
	const _fc = 600
	_abd = MakeRuneCharSafeMap(len(_cc))
	for _, _eeg := range _cc {
		_abd.Write(_eeg, CharMetrics{Wx: _fc})
	}
	_gdd = _abd.Copy()
	_acde = _abd.Copy()
	_ef = _abd.Copy()
}
func (_dd StdFont) Name() string { return string(_dd._bgg.Name) }
func _beb() StdFont {
	_ddea.Do(_gff)
	_cf := Descriptor{Name: CourierObliqueName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0061, BBox: [4]float64{-27, -250, 849, 805}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_cf, _ef)
}
func (_dee *ttfParser) ParseMaxp() error {
	if _egfe := _dee.Seek("\u006d\u0061\u0078\u0070"); _egfe != nil {
		return _egfe
	}
	_dee.Skip(4)
	_dee._bff = _dee.ReadUShort()
	return nil
}
func RegisterStdFont(name StdFontName, fnc func() StdFont, aliases ...StdFontName) {
	if _, _aca := _dfb.read(name); _aca {
		panic("\u0066o\u006e\u0074\u0020\u0061l\u0072\u0065\u0061\u0064\u0079 \u0072e\u0067i\u0073\u0074\u0065\u0072\u0065\u0064\u003a " + string(name))
	}
	_dfb.write(name, fnc)
	for _, _gbb := range aliases {
		RegisterStdFont(_gbb, fnc)
	}
}

type ttfParser struct {
	_dad  TtfType
	_gbca _e.ReadSeeker
	_ggc  map[string]uint32
	_bgda uint16
	_bff  uint16
}

func _eeb() StdFont {
	_dagc.Do(_eac)
	_db := Descriptor{Name: TimesBoldItalicName, Family: _bef, Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-200, -218, 996, 921}, ItalicAngle: -15, Ascent: 683, Descent: -217, CapHeight: 669, XHeight: 462, StemV: 121, StemH: 42}
	return NewStdFont(_db, _gcb)
}

var _ddea _fg.Once
var _dfb = &fontMap{_bc: make(map[StdFontName]func() StdFont)}

func (_ebaf *ttfParser) ReadULong() (_feg uint32) {
	_dg.Read(_ebaf._gbca, _dg.BigEndian, &_feg)
	return _feg
}
func (_eca *RuneCharSafeMap) Write(b rune, r CharMetrics) {
	_eca._gad.Lock()
	defer _eca._gad.Unlock()
	_eca._ga[b] = r
}
func (_faa *ttfParser) ReadSByte() (_baeb int8) {
	_dg.Read(_faa._gbca, _dg.BigEndian, &_baeb)
	return _baeb
}

const (
	_bef                = "\u0054\u0069\u006de\u0073"
	TimesRomanName      = StdFontName("T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e")
	TimesBoldName       = StdFontName("\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064")
	TimesItalicName     = StdFontName("\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063")
	TimesBoldItalicName = StdFontName("\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
)

func (_ccbg *ttfParser) Seek(tag string) error {
	_gddc, _fecg := _ccbg._ggc[tag]
	if !_fecg {
		return _f.Errorf("\u0074\u0061\u0062\u006ce \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0073", tag)
	}
	_ccbg._gbca.Seek(int64(_gddc), _e.SeekStart)
	return nil
}
func NewFontFile2FromPdfObject(obj _ag.PdfObject) (TtfType, error) {
	obj = _ag.TraceToDirectObject(obj)
	_dca, _gaaf := obj.(*_ag.PdfObjectStream)
	if !_gaaf {
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072e\u0061\u006d \u0028\u0025\u0054\u0029", obj)
		return TtfType{}, _ag.ErrTypeError
	}
	_afg, _eecc := _ag.DecodeStream(_dca)
	if _eecc != nil {
		return TtfType{}, _eecc
	}
	_cea := ttfParser{_gbca: _gf.NewReader(_afg)}
	return _cea.Parse()
}

var _dcd *RuneCharSafeMap

func _ccc() StdFont {
	_fbb := _b.NewZapfDingbatsEncoder()
	_bfg := Descriptor{Name: ZapfDingbatsName, Family: string(ZapfDingbatsName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-1, -143, 981, 820}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 90, StemH: 28}
	return NewStdFontWithEncoding(_bfg, _cccg, _fbb)
}
func (_bg *RuneCharSafeMap) Read(b rune) (CharMetrics, bool) {
	_bg._gad.RLock()
	defer _bg._gad.RUnlock()
	_be, _dfg := _bg._ga[b]
	return _be, _dfg
}

var _egae = []int16{667, 944, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 667, 722, 722, 722, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 667, 667, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 889, 667, 611, 611, 611, 611, 611, 611, 611, 500, 500, 500, 500, 333, 500, 722, 500, 500, 778, 500, 500, 570, 570, 500, 832, 500, 500, 278, 220, 348, 348, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 500, 500, 500, 608, 500, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 389, 389, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 382, 278, 570, 549, 606, 494, 278, 778, 333, 606, 576, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 266, 300, 500, 500, 500, 500, 333, 333, 494, 833, 250, 250, 1000, 570, 570, 500, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 389, 389, 549, 389, 389, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 366, 278, 500, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 444, 667, 500, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func _abc() StdFont {
	_dfe.Do(_acg)
	_deff := Descriptor{Name: HelveticaBoldObliqueName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-174, -228, 1114, 962}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_deff, _baed)
}
func NewStdFontWithEncoding(desc Descriptor, metrics *RuneCharSafeMap, encoder _b.TextEncoder) StdFont {
	var _ea rune = 0xA0
	if _, _bfa := metrics.Read(_ea); !_bfa {
		_fgd, _ := metrics.Read(0x20)
		metrics.Write(_ea, _fgd)
	}
	return StdFont{_bgg: desc, _dc: metrics, _da: encoder}
}

type fontMap struct {
	_fg.Mutex
	_bc map[StdFontName]func() StdFont
}

var _fe = []int16{611, 889, 611, 611, 611, 611, 611, 611, 611, 611, 611, 611, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 611, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 444, 667, 667, 556, 556, 611, 556, 556, 833, 667, 667, 667, 667, 667, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 611, 611, 611, 611, 500, 500, 500, 500, 500, 556, 556, 556, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 611, 833, 611, 556, 556, 556, 556, 556, 556, 556, 500, 500, 500, 500, 333, 500, 667, 500, 500, 778, 500, 500, 422, 541, 500, 920, 500, 500, 278, 275, 400, 400, 389, 389, 333, 275, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 760, 500, 500, 500, 500, 544, 500, 400, 333, 675, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 889, 444, 889, 500, 444, 675, 500, 333, 389, 278, 500, 500, 500, 500, 500, 167, 500, 500, 500, 500, 333, 675, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 444, 444, 278, 278, 300, 278, 675, 549, 675, 471, 278, 722, 333, 675, 500, 675, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 667, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 523, 333, 333, 476, 833, 250, 250, 1000, 675, 675, 500, 500, 500, 420, 556, 556, 556, 333, 333, 333, 214, 389, 389, 453, 389, 389, 760, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 300, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 444, 667, 444, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func _gab() StdFont {
	_dfe.Do(_acg)
	_fgdb := Descriptor{Name: HelveticaObliqueName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-170, -225, 1116, 931}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_fgdb, _gbc)
}

var _cccg = &RuneCharSafeMap{_ga: map[rune]CharMetrics{' ': {Wx: 278}, '→': {Wx: 838}, '↔': {Wx: 1016}, '↕': {Wx: 458}, '①': {Wx: 788}, '②': {Wx: 788}, '③': {Wx: 788}, '④': {Wx: 788}, '⑤': {Wx: 788}, '⑥': {Wx: 788}, '⑦': {Wx: 788}, '⑧': {Wx: 788}, '⑨': {Wx: 788}, '⑩': {Wx: 788}, '■': {Wx: 761}, '▲': {Wx: 892}, '▼': {Wx: 892}, '◆': {Wx: 788}, '●': {Wx: 791}, '◗': {Wx: 438}, '★': {Wx: 816}, '☎': {Wx: 719}, '☛': {Wx: 960}, '☞': {Wx: 939}, '♠': {Wx: 626}, '♣': {Wx: 776}, '♥': {Wx: 694}, '♦': {Wx: 595}, '✁': {Wx: 974}, '✂': {Wx: 961}, '✃': {Wx: 974}, '✄': {Wx: 980}, '✆': {Wx: 789}, '✇': {Wx: 790}, '✈': {Wx: 791}, '✉': {Wx: 690}, '✌': {Wx: 549}, '✍': {Wx: 855}, '✎': {Wx: 911}, '✏': {Wx: 933}, '✐': {Wx: 911}, '✑': {Wx: 945}, '✒': {Wx: 974}, '✓': {Wx: 755}, '✔': {Wx: 846}, '✕': {Wx: 762}, '✖': {Wx: 761}, '✗': {Wx: 571}, '✘': {Wx: 677}, '✙': {Wx: 763}, '✚': {Wx: 760}, '✛': {Wx: 759}, '✜': {Wx: 754}, '✝': {Wx: 494}, '✞': {Wx: 552}, '✟': {Wx: 537}, '✠': {Wx: 577}, '✡': {Wx: 692}, '✢': {Wx: 786}, '✣': {Wx: 788}, '✤': {Wx: 788}, '✥': {Wx: 790}, '✦': {Wx: 793}, '✧': {Wx: 794}, '✩': {Wx: 823}, '✪': {Wx: 789}, '✫': {Wx: 841}, '✬': {Wx: 823}, '✭': {Wx: 833}, '✮': {Wx: 816}, '✯': {Wx: 831}, '✰': {Wx: 923}, '✱': {Wx: 744}, '✲': {Wx: 723}, '✳': {Wx: 749}, '✴': {Wx: 790}, '✵': {Wx: 792}, '✶': {Wx: 695}, '✷': {Wx: 776}, '✸': {Wx: 768}, '✹': {Wx: 792}, '✺': {Wx: 759}, '✻': {Wx: 707}, '✼': {Wx: 708}, '✽': {Wx: 682}, '✾': {Wx: 701}, '✿': {Wx: 826}, '❀': {Wx: 815}, '❁': {Wx: 789}, '❂': {Wx: 789}, '❃': {Wx: 707}, '❄': {Wx: 687}, '❅': {Wx: 696}, '❆': {Wx: 689}, '❇': {Wx: 786}, '❈': {Wx: 787}, '❉': {Wx: 713}, '❊': {Wx: 791}, '❋': {Wx: 785}, '❍': {Wx: 873}, '❏': {Wx: 762}, '❐': {Wx: 762}, '❑': {Wx: 759}, '❒': {Wx: 759}, '❖': {Wx: 784}, '❘': {Wx: 138}, '❙': {Wx: 277}, '❚': {Wx: 415}, '❛': {Wx: 392}, '❜': {Wx: 392}, '❝': {Wx: 668}, '❞': {Wx: 668}, '❡': {Wx: 732}, '❢': {Wx: 544}, '❣': {Wx: 544}, '❤': {Wx: 910}, '❥': {Wx: 667}, '❦': {Wx: 760}, '❧': {Wx: 760}, '❶': {Wx: 788}, '❷': {Wx: 788}, '❸': {Wx: 788}, '❹': {Wx: 788}, '❺': {Wx: 788}, '❻': {Wx: 788}, '❼': {Wx: 788}, '❽': {Wx: 788}, '❾': {Wx: 788}, '❿': {Wx: 788}, '➀': {Wx: 788}, '➁': {Wx: 788}, '➂': {Wx: 788}, '➃': {Wx: 788}, '➄': {Wx: 788}, '➅': {Wx: 788}, '➆': {Wx: 788}, '➇': {Wx: 788}, '➈': {Wx: 788}, '➉': {Wx: 788}, '➊': {Wx: 788}, '➋': {Wx: 788}, '➌': {Wx: 788}, '➍': {Wx: 788}, '➎': {Wx: 788}, '➏': {Wx: 788}, '➐': {Wx: 788}, '➑': {Wx: 788}, '➒': {Wx: 788}, '➓': {Wx: 788}, '➔': {Wx: 894}, '➘': {Wx: 748}, '➙': {Wx: 924}, '➚': {Wx: 748}, '➛': {Wx: 918}, '➜': {Wx: 927}, '➝': {Wx: 928}, '➞': {Wx: 928}, '➟': {Wx: 834}, '➠': {Wx: 873}, '➡': {Wx: 828}, '➢': {Wx: 924}, '➣': {Wx: 924}, '➤': {Wx: 917}, '➥': {Wx: 930}, '➦': {Wx: 931}, '➧': {Wx: 463}, '➨': {Wx: 883}, '➩': {Wx: 836}, '➪': {Wx: 836}, '➫': {Wx: 867}, '➬': {Wx: 867}, '➭': {Wx: 696}, '➮': {Wx: 696}, '➯': {Wx: 874}, '➱': {Wx: 874}, '➲': {Wx: 760}, '➳': {Wx: 946}, '➴': {Wx: 771}, '➵': {Wx: 865}, '➶': {Wx: 771}, '➷': {Wx: 888}, '➸': {Wx: 967}, '➹': {Wx: 888}, '➺': {Wx: 831}, '➻': {Wx: 873}, '➼': {Wx: 927}, '➽': {Wx: 970}, '➾': {Wx: 918}, '\uf8d7': {Wx: 390}, '\uf8d8': {Wx: 390}, '\uf8d9': {Wx: 317}, '\uf8da': {Wx: 317}, '\uf8db': {Wx: 276}, '\uf8dc': {Wx: 276}, '\uf8dd': {Wx: 509}, '\uf8de': {Wx: 509}, '\uf8df': {Wx: 410}, '\uf8e0': {Wx: 410}, '\uf8e1': {Wx: 234}, '\uf8e2': {Wx: 234}, '\uf8e3': {Wx: 334}, '\uf8e4': {Wx: 334}}}

func MakeRuneCharSafeMap(length int) *RuneCharSafeMap {
	return &RuneCharSafeMap{_ga: make(map[rune]CharMetrics, length)}
}
func (_abcc *ttfParser) parseCmapFormat0() error {
	_cdaf, _bag := _abcc.ReadStr(256)
	if _bag != nil {
		return _bag
	}
	_gdaf := []byte(_cdaf)
	_fb.Log.Trace("\u0070a\u0072\u0073e\u0043\u006d\u0061p\u0046\u006f\u0072\u006d\u0061\u0074\u0030:\u0020\u0025\u0073\u000a\u0064\u0061t\u0061\u0053\u0074\u0072\u003d\u0025\u002b\u0071\u000a\u0064\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d", _abcc._dad.String(), _cdaf, _gdaf)
	for _fga, _ecc := range _gdaf {
		_abcc._dad.Chars[rune(_fga)] = GID(_ecc)
	}
	return nil
}
