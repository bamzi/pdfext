package imageutil

import (
	_bag "encoding/binary"
	_b "errors"
	_d "fmt"
	_a "image"
	_bd "image/color"
	_bb "image/draw"
	_ba "math"

	_e "github.com/bamzi/pdfext/common"
	_ab "github.com/bamzi/pdfext/internal/bitwise"
)

type ColorConverter interface {
	Convert(_fdb _a.Image) (Image, error)
}

func (_geea *NRGBA64) Set(x, y int, c _bd.Color) {
	_abca := (y*_geea.Width + x) * 2
	_ffbgd := _abca * 3
	if _ffbgd+5 >= len(_geea.Data) {
		return
	}
	_defd := _bd.NRGBA64Model.Convert(c).(_bd.NRGBA64)
	_geea.setNRGBA64(_ffbgd, _defd, _abca)
}
func _bdc() (_acf []byte) {
	_acf = make([]byte, 256)
	for _bcd := 0; _bcd < 256; _bcd++ {
		_fcc := byte(_bcd)
		_acf[_fcc] = (_fcc & 0x01) | ((_fcc & 0x04) >> 1) | ((_fcc & 0x10) >> 2) | ((_fcc & 0x40) >> 3) | ((_fcc & 0x02) << 3) | ((_fcc & 0x08) << 2) | ((_fcc & 0x20) << 1) | (_fcc & 0x80)
	}
	return _acf
}
func (_agd *ImageBase) setEightBytes(_gbdb int, _bfgf uint64) error {
	_aedc := _agd.BytesPerLine - (_gbdb % _agd.BytesPerLine)
	if _agd.BytesPerLine != _agd.Width>>3 {
		_aedc--
	}
	if _aedc >= 8 {
		return _agd.setEightFullBytes(_gbdb, _bfgf)
	}
	return _agd.setEightPartlyBytes(_gbdb, _aedc, _bfgf)
}

var _ Gray = &Monochrome{}

func (_gdgf *CMYK32) SetCMYK(x, y int, c _bd.CMYK) {
	_fac := 4 * (y*_gdgf.Width + x)
	if _fac+3 >= len(_gdgf.Data) {
		return
	}
	_gdgf.Data[_fac] = c.C
	_gdgf.Data[_fac+1] = c.M
	_gdgf.Data[_fac+2] = c.Y
	_gdgf.Data[_fac+3] = c.K
}
func AutoThresholdTriangle(histogram [256]int) uint8 {
	var _acacg, _bggb, _fgad, _bgbc int
	for _cafc := 0; _cafc < len(histogram); _cafc++ {
		if histogram[_cafc] > 0 {
			_acacg = _cafc
			break
		}
	}
	if _acacg > 0 {
		_acacg--
	}
	for _gdeb := 255; _gdeb > 0; _gdeb-- {
		if histogram[_gdeb] > 0 {
			_bgbc = _gdeb
			break
		}
	}
	if _bgbc < 255 {
		_bgbc++
	}
	for _gege := 0; _gege < 256; _gege++ {
		if histogram[_gege] > _bggb {
			_fgad = _gege
			_bggb = histogram[_gege]
		}
	}
	var _ccfbe bool
	if (_fgad - _acacg) < (_bgbc - _fgad) {
		_ccfbe = true
		var _ccfe int
		_cfbfe := 255
		for _ccfe < _cfbfe {
			_dccb := histogram[_ccfe]
			histogram[_ccfe] = histogram[_cfbfe]
			histogram[_cfbfe] = _dccb
			_ccfe++
			_cfbfe--
		}
		_acacg = 255 - _bgbc
		_fgad = 255 - _fgad
	}
	if _acacg == _fgad {
		return uint8(_acacg)
	}
	_acag := float64(histogram[_fgad])
	_eeec := float64(_acacg - _fgad)
	_ecfe := _ba.Sqrt(_acag*_acag + _eeec*_eeec)
	_acag /= _ecfe
	_eeec /= _ecfe
	_ecfe = _acag*float64(_acacg) + _eeec*float64(histogram[_acacg])
	_bdfg := _acacg
	var _fbge float64
	for _cfde := _acacg + 1; _cfde <= _fgad; _cfde++ {
		_ddga := _acag*float64(_cfde) + _eeec*float64(histogram[_cfde]) - _ecfe
		if _ddga > _fbge {
			_bdfg = _cfde
			_fbge = _ddga
		}
	}
	_bdfg--
	if _ccfbe {
		var _ebcd int
		_gebeg := 255
		for _ebcd < _gebeg {
			_bfad := histogram[_ebcd]
			histogram[_ebcd] = histogram[_gebeg]
			histogram[_gebeg] = _bfad
			_ebcd++
			_gebeg--
		}
		return uint8(255 - _bdfg)
	}
	return uint8(_bdfg)
}
func _fgga(_egcd _a.Image) (Image, error) {
	if _bacee, _gge := _egcd.(*NRGBA32); _gge {
		return _bacee.Copy(), nil
	}
	_gbee, _ggec, _eafe := _aagg(_egcd, 1)
	_dbdf, _bcbf := NewImage(_gbee.Max.X, _gbee.Max.Y, 8, 3, nil, _eafe, nil)
	if _bcbf != nil {
		return nil, _bcbf
	}
	_dcda(_egcd, _dbdf, _gbee)
	if len(_eafe) != 0 && !_ggec {
		if _aged := _dcafa(_eafe, _dbdf); _aged != nil {
			return nil, _aged
		}
	}
	return _dbdf, nil
}
func (_aggb *CMYK32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _aggb.Width, Y: _aggb.Height}}
}
func (_ccgea *RGBA32) Validate() error {
	if len(_ccgea.Data) != 3*_ccgea.Width*_ccgea.Height {
		return _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _aagb(_eagc _a.Image) (Image, error) {
	if _cfc, _fbaa := _eagc.(*Gray2); _fbaa {
		return _cfc.Copy(), nil
	}
	_abgb := _eagc.Bounds()
	_bbaa, _fgbg := NewImage(_abgb.Max.X, _abgb.Max.Y, 2, 1, nil, nil, nil)
	if _fgbg != nil {
		return nil, _fgbg
	}
	_abad(_eagc, _bbaa, _abgb)
	return _bbaa, nil
}
func (_geee *RGBA32) ColorModel() _bd.Model { return _bd.NRGBAModel }
func ColorAtCMYK(x, y, width int, data []byte, decode []float64) (_bd.CMYK, error) {
	_ceg := 4 * (y*width + x)
	if _ceg+3 >= len(data) {
		return _bd.CMYK{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	C := data[_ceg] & 0xff
	M := data[_ceg+1] & 0xff
	Y := data[_ceg+2] & 0xff
	K := data[_ceg+3] & 0xff
	if len(decode) == 8 {
		C = uint8(uint32(LinearInterpolate(float64(C), 0, 255, decode[0], decode[1])) & 0xff)
		M = uint8(uint32(LinearInterpolate(float64(M), 0, 255, decode[2], decode[3])) & 0xff)
		Y = uint8(uint32(LinearInterpolate(float64(Y), 0, 255, decode[4], decode[5])) & 0xff)
		K = uint8(uint32(LinearInterpolate(float64(K), 0, 255, decode[6], decode[7])) & 0xff)
	}
	return _bd.CMYK{C: C, M: M, Y: Y, K: K}, nil
}
func _eaf(_gead _bd.NYCbCrA) _bd.RGBA {
	_dafc, _daef, _fffb, _gcbe := _fgd(_gead).RGBA()
	return _bd.RGBA{R: uint8(_dafc >> 8), G: uint8(_daef >> 8), B: uint8(_fffb >> 8), A: uint8(_gcbe >> 8)}
}

type Gray interface {
	GrayAt(_cgd, _becf int) _bd.Gray
	SetGray(_cddc, _bebe int, _aag _bd.Gray)
}
type NRGBA32 struct{ ImageBase }

func _adebe(_aeg *_a.NYCbCrA, _agccf NRGBA, _bbgf _a.Rectangle) {
	for _gcfc := 0; _gcfc < _bbgf.Max.X; _gcfc++ {
		for _dec := 0; _dec < _bbgf.Max.Y; _dec++ {
			_gcege := _aeg.NYCbCrAAt(_gcfc, _dec)
			_agccf.SetNRGBA(_gcfc, _dec, _fgd(_gcege))
		}
	}
}
func (_edcf *Gray2) Copy() Image { return &Gray2{ImageBase: _edcf.copy()} }
func _edf(_agf _bd.Color) _bd.Color {
	_afba := _bd.GrayModel.Convert(_agf).(_bd.Gray)
	return _efc(_afba)
}

var _ NRGBA = &NRGBA16{}

func _bged(_fabef _a.Image) (Image, error) {
	if _egab, _cedf := _fabef.(*RGBA32); _cedf {
		return _egab.Copy(), nil
	}
	_abbgc, _dcbce, _aacbb := _aagg(_fabef, 1)
	_ebgb := &RGBA32{ImageBase: NewImageBase(_abbgc.Max.X, _abbgc.Max.Y, 8, 3, nil, _aacbb, nil)}
	_cdcdd(_fabef, _ebgb, _abbgc)
	if len(_aacbb) != 0 && !_dcbce {
		if _abea := _dcafa(_aacbb, _ebgb); _abea != nil {
			return nil, _abea
		}
	}
	return _ebgb, nil
}
func _fbf(_ebad, _bee int, _add []byte) *Monochrome {
	_ggc := _efb(_ebad, _bee)
	_ggc.Data = _add
	return _ggc
}

type Image interface {
	_bb.Image
	Base() *ImageBase
	Copy() Image
	Pix() []byte
	ColorAt(_ada, _aeedf int) (_bd.Color, error)
	Validate() error
}

func _bgba(_fggc Gray, _ecc CMYK, _ggd _a.Rectangle) {
	for _dca := 0; _dca < _ggd.Max.X; _dca++ {
		for _cgag := 0; _cgag < _ggd.Max.Y; _cgag++ {
			_cega := _fggc.GrayAt(_dca, _cgag)
			_ecc.SetCMYK(_dca, _cgag, _bebd(_cega))
		}
	}
}
func _ffga(_gdg int) []uint {
	var _acd []uint
	_bdb := _gdg
	_ag := _bdb / 8
	if _ag != 0 {
		for _eeg := 0; _eeg < _ag; _eeg++ {
			_acd = append(_acd, 8)
		}
		_dac := _bdb % 8
		_bdb = 0
		if _dac != 0 {
			_bdb = _dac
		}
	}
	_gce := _bdb / 4
	if _gce != 0 {
		for _gea := 0; _gea < _gce; _gea++ {
			_acd = append(_acd, 4)
		}
		_dcc := _bdb % 4
		_bdb = 0
		if _dcc != 0 {
			_bdb = _dcc
		}
	}
	_bgb := _bdb / 2
	if _bgb != 0 {
		for _eef := 0; _eef < _bgb; _eef++ {
			_acd = append(_acd, 2)
		}
	}
	return _acd
}
func (_dad *Monochrome) At(x, y int) _bd.Color { _dbbc, _ := _dad.ColorAt(x, y); return _dbbc }
func (_eabg *NRGBA32) Base() *ImageBase        { return &_eabg.ImageBase }
func _geb(_gb *Monochrome, _gafba int, _gfcd []byte) (_ccd *Monochrome, _gcd error) {
	const _aed = "\u0072\u0065d\u0075\u0063\u0065R\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079"
	if _gb == nil {
		return nil, _b.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _gafba < 1 || _gafba > 4 {
		return nil, _b.New("\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _gb.Height <= 1 {
		return nil, _b.New("\u0073\u006f\u0075rc\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020m\u0075s\u0074 \u0062e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027")
	}
	_ccd = _efb(_gb.Width/2, _gb.Height/2)
	if _gfcd == nil {
		_gfcd = _bdc()
	}
	_ffc := _cdcd(_gb.BytesPerLine, 2*_ccd.BytesPerLine)
	switch _gafba {
	case 1:
		_gcd = _cee(_gb, _ccd, _gfcd, _ffc)
	case 2:
		_gcd = _dba(_gb, _ccd, _gfcd, _ffc)
	case 3:
		_gcd = _cag(_gb, _ccd, _gfcd, _ffc)
	case 4:
		_gcd = _ggg(_gb, _ccd, _gfcd, _ffc)
	}
	if _gcd != nil {
		return nil, _gcd
	}
	return _ccd, nil
}
func (_bdfe *Gray16) Base() *ImageBase { return &_bdfe.ImageBase }
func (_gcfd *NRGBA16) Set(x, y int, c _bd.Color) {
	_cdafa := y*_gcfd.BytesPerLine + x*3/2
	if _cdafa+1 >= len(_gcfd.Data) {
		return
	}
	_acgdc := NRGBA16Model.Convert(c).(_bd.NRGBA)
	_gcfd.setNRGBA(x, y, _cdafa, _acgdc)
}
func _age(_cffg _bd.Gray) _bd.RGBA { return _bd.RGBA{R: _cffg.Y, G: _cffg.Y, B: _cffg.Y, A: 0xff} }
func _ebec(_beg *Monochrome, _afab, _fad int, _gbde, _deag int, _bcbb RasterOperator) {
	var (
		_bcaf  bool
		_egdb  bool
		_efef  int
		_febe  int
		_cdfd  int
		_dgdgb int
		_fegf  bool
		_egeeb byte
	)
	_fdfd := 8 - (_afab & 7)
	_fccaf := _bbgc[_fdfd]
	_bbg := _beg.BytesPerLine*_fad + (_afab >> 3)
	if _gbde < _fdfd {
		_bcaf = true
		_fccaf &= _ddfe[8-_fdfd+_gbde]
	}
	if !_bcaf {
		_efef = (_gbde - _fdfd) >> 3
		if _efef != 0 {
			_egdb = true
			_febe = _bbg + 1
		}
	}
	_cdfd = (_afab + _gbde) & 7
	if !(_bcaf || _cdfd == 0) {
		_fegf = true
		_egeeb = _ddfe[_cdfd]
		_dgdgb = _bbg + 1 + _efef
	}
	var _eadcad, _dcbf int
	switch _bcbb {
	case PixClr:
		for _eadcad = 0; _eadcad < _deag; _eadcad++ {
			_beg.Data[_bbg] = _agcee(_beg.Data[_bbg], 0x0, _fccaf)
			_bbg += _beg.BytesPerLine
		}
		if _egdb {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				for _dcbf = 0; _dcbf < _efef; _dcbf++ {
					_beg.Data[_febe+_dcbf] = 0x0
				}
				_febe += _beg.BytesPerLine
			}
		}
		if _fegf {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				_beg.Data[_dgdgb] = _agcee(_beg.Data[_dgdgb], 0x0, _egeeb)
				_dgdgb += _beg.BytesPerLine
			}
		}
	case PixSet:
		for _eadcad = 0; _eadcad < _deag; _eadcad++ {
			_beg.Data[_bbg] = _agcee(_beg.Data[_bbg], 0xff, _fccaf)
			_bbg += _beg.BytesPerLine
		}
		if _egdb {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				for _dcbf = 0; _dcbf < _efef; _dcbf++ {
					_beg.Data[_febe+_dcbf] = 0xff
				}
				_febe += _beg.BytesPerLine
			}
		}
		if _fegf {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				_beg.Data[_dgdgb] = _agcee(_beg.Data[_dgdgb], 0xff, _egeeb)
				_dgdgb += _beg.BytesPerLine
			}
		}
	case PixNotDst:
		for _eadcad = 0; _eadcad < _deag; _eadcad++ {
			_beg.Data[_bbg] = _agcee(_beg.Data[_bbg], ^_beg.Data[_bbg], _fccaf)
			_bbg += _beg.BytesPerLine
		}
		if _egdb {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				for _dcbf = 0; _dcbf < _efef; _dcbf++ {
					_beg.Data[_febe+_dcbf] = ^(_beg.Data[_febe+_dcbf])
				}
				_febe += _beg.BytesPerLine
			}
		}
		if _fegf {
			for _eadcad = 0; _eadcad < _deag; _eadcad++ {
				_beg.Data[_dgdgb] = _agcee(_beg.Data[_dgdgb], ^_beg.Data[_dgdgb], _egeeb)
				_dgdgb += _beg.BytesPerLine
			}
		}
	}
}
func (_feffg *NRGBA32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _feffg.Width, Y: _feffg.Height}}
}
func _ggg(_agg, _ebd *Monochrome, _eadb []byte, _fee int) (_adb error) {
	var (
		_acb, _ffcb, _ebb, _aeee, _ddd, _geg, _efdf, _adeg int
		_abgc, _aeb                                        uint32
		_aggc, _gac                                        byte
		_ccg                                               uint16
	)
	_aaff := make([]byte, 4)
	_caa := make([]byte, 4)
	for _ebb = 0; _ebb < _agg.Height-1; _ebb, _aeee = _ebb+2, _aeee+1 {
		_acb = _ebb * _agg.BytesPerLine
		_ffcb = _aeee * _ebd.BytesPerLine
		for _ddd, _geg = 0, 0; _ddd < _fee; _ddd, _geg = _ddd+4, _geg+1 {
			for _efdf = 0; _efdf < 4; _efdf++ {
				_adeg = _acb + _ddd + _efdf
				if _adeg <= len(_agg.Data)-1 && _adeg < _acb+_agg.BytesPerLine {
					_aaff[_efdf] = _agg.Data[_adeg]
				} else {
					_aaff[_efdf] = 0x00
				}
				_adeg = _acb + _agg.BytesPerLine + _ddd + _efdf
				if _adeg <= len(_agg.Data)-1 && _adeg < _acb+(2*_agg.BytesPerLine) {
					_caa[_efdf] = _agg.Data[_adeg]
				} else {
					_caa[_efdf] = 0x00
				}
			}
			_abgc = _bag.BigEndian.Uint32(_aaff)
			_aeb = _bag.BigEndian.Uint32(_caa)
			_aeb &= _abgc
			_aeb &= _aeb << 1
			_aeb &= 0xaaaaaaaa
			_abgc = _aeb | (_aeb << 7)
			_aggc = byte(_abgc >> 24)
			_gac = byte((_abgc >> 8) & 0xff)
			_adeg = _ffcb + _geg
			if _adeg+1 == len(_ebd.Data)-1 || _adeg+1 >= _ffcb+_ebd.BytesPerLine {
				_ebd.Data[_adeg] = _eadb[_aggc]
				if _adb = _ebd.setByte(_adeg, _eadb[_aggc]); _adb != nil {
					return _d.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _adeg)
				}
			} else {
				_ccg = (uint16(_eadb[_aggc]) << 8) | uint16(_eadb[_gac])
				if _adb = _ebd.setTwoBytes(_adeg, _ccg); _adb != nil {
					return _d.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _adeg)
				}
				_geg++
			}
		}
	}
	return nil
}
func ColorAtNRGBA16(x, y, width, bytesPerLine int, data, alpha []byte, decode []float64) (_bd.NRGBA, error) {
	_dadb := y*bytesPerLine + x*3/2
	if _dadb+1 >= len(data) {
		return _bd.NRGBA{}, _dfgb(x, y)
	}
	const (
		_feeaf = 0xf
		_feaad = uint8(0xff)
	)
	_bfgdg := _feaad
	if alpha != nil {
		_ffdc := y * BytesPerLine(width, 4, 1)
		if _ffdc < len(alpha) {
			if x%2 == 0 {
				_bfgdg = (alpha[_ffdc] >> uint(4)) & _feeaf
			} else {
				_bfgdg = alpha[_ffdc] & _feeaf
			}
			_bfgdg |= _bfgdg << 4
		}
	}
	var _baad, _aafdf, _fabe uint8
	if x*3%2 == 0 {
		_baad = (data[_dadb] >> uint(4)) & _feeaf
		_aafdf = data[_dadb] & _feeaf
		_fabe = (data[_dadb+1] >> uint(4)) & _feeaf
	} else {
		_baad = data[_dadb] & _feeaf
		_aafdf = (data[_dadb+1] >> uint(4)) & _feeaf
		_fabe = data[_dadb+1] & _feeaf
	}
	if len(decode) == 6 {
		_baad = uint8(uint32(LinearInterpolate(float64(_baad), 0, 15, decode[0], decode[1])) & 0xf)
		_aafdf = uint8(uint32(LinearInterpolate(float64(_aafdf), 0, 15, decode[2], decode[3])) & 0xf)
		_fabe = uint8(uint32(LinearInterpolate(float64(_fabe), 0, 15, decode[4], decode[5])) & 0xf)
	}
	return _bd.NRGBA{R: (_baad << 4) | (_baad & 0xf), G: (_aafdf << 4) | (_aafdf & 0xf), B: (_fabe << 4) | (_fabe & 0xf), A: _bfgdg}, nil
}
func ColorAtNRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_bd.NRGBA, error) {
	_fgba := y*width + x
	_gfa := 3 * _fgba
	if _gfa+2 >= len(data) {
		return _bd.NRGBA{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_gcge := uint8(0xff)
	if alpha != nil && len(alpha) > _fgba {
		_gcge = alpha[_fgba]
	}
	_gfdf, _bcdd, _aefb := data[_gfa], data[_gfa+1], data[_gfa+2]
	if len(decode) == 6 {
		_gfdf = uint8(uint32(LinearInterpolate(float64(_gfdf), 0, 255, decode[0], decode[1])) & 0xff)
		_bcdd = uint8(uint32(LinearInterpolate(float64(_bcdd), 0, 255, decode[2], decode[3])) & 0xff)
		_aefb = uint8(uint32(LinearInterpolate(float64(_aefb), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _bd.NRGBA{R: _gfdf, G: _bcdd, B: _aefb, A: _gcge}, nil
}

type Gray16 struct{ ImageBase }

func _cebe(_cbbb, _gdda RGBA, _acea _a.Rectangle) {
	for _gbbde := 0; _gbbde < _acea.Max.X; _gbbde++ {
		for _ecbf := 0; _ecbf < _acea.Max.Y; _ecbf++ {
			_gdda.SetRGBA(_gbbde, _ecbf, _cbbb.RGBAAt(_gbbde, _ecbf))
		}
	}
}
func NextPowerOf2(n uint) uint {
	if IsPowerOf2(n) {
		return n
	}
	return 1 << (_ccggd(n) + 1)
}
func _dfcd(_ceb _a.Image) (Image, error) {
	if _ccgg, _fdec := _ceb.(*Gray16); _fdec {
		return _ccgg.Copy(), nil
	}
	_bbad := _ceb.Bounds()
	_gebc, _dcee := NewImage(_bbad.Max.X, _bbad.Max.Y, 16, 1, nil, nil, nil)
	if _dcee != nil {
		return nil, _dcee
	}
	_abad(_ceb, _gebc, _bbad)
	return _gebc, nil
}
func MonochromeModel(threshold uint8) _bd.Model { return monochromeModel(threshold) }
func (_dfde *NRGBA32) At(x, y int) _bd.Color    { _bccf, _ := _dfde.ColorAt(x, y); return _bccf }
func _dbce(_fbfg _bd.NRGBA64) _bd.Gray {
	var _ebff _bd.NRGBA64
	if _fbfg == _ebff {
		return _bd.Gray{Y: 0xff}
	}
	_agbb, _dbac, _abge, _ := _fbfg.RGBA()
	_fgee := (19595*_agbb + 38470*_dbac + 7471*_abge + 1<<15) >> 24
	return _bd.Gray{Y: uint8(_fgee)}
}

var _ _a.Image = &NRGBA16{}

type NRGBA16 struct{ ImageBase }

func _dade(_fdg _a.Image) (Image, error) {
	if _ffdcf, _aafa := _fdg.(*NRGBA16); _aafa {
		return _ffdcf.Copy(), nil
	}
	_bcged := _fdg.Bounds()
	_bgca, _abcb := NewImage(_bcged.Max.X, _bcged.Max.Y, 4, 3, nil, nil, nil)
	if _abcb != nil {
		return nil, _abcb
	}
	_dcda(_fdg, _bgca, _bcged)
	return _bgca, nil
}
func _agce(_eag _bd.RGBA) _bd.NRGBA {
	switch _eag.A {
	case 0xff:
		return _bd.NRGBA{R: _eag.R, G: _eag.G, B: _eag.B, A: 0xff}
	case 0x00:
		return _bd.NRGBA{}
	default:
		_dfb, _eedb, _abbd, _bfg := _eag.RGBA()
		_dfb = (_dfb * 0xffff) / _bfg
		_eedb = (_eedb * 0xffff) / _bfg
		_abbd = (_abbd * 0xffff) / _bfg
		return _bd.NRGBA{R: uint8(_dfb >> 8), G: uint8(_eedb >> 8), B: uint8(_abbd >> 8), A: uint8(_bfg >> 8)}
	}
}
func (_cgfd *Gray16) SetGray(x, y int, g _bd.Gray) {
	_cegd := (y*_cgfd.BytesPerLine/2 + x) * 2
	if _cegd+1 >= len(_cgfd.Data) {
		return
	}
	_cgfd.Data[_cegd] = g.Y
	_cgfd.Data[_cegd+1] = g.Y
}
func _gfee(_edca *Monochrome, _bffcb, _dafe int, _ggdd, _afd int, _gfdd RasterOperator) {
	var (
		_egae        int
		_aafg        byte
		_bdcd, _ffef int
		_ecfg        int
	)
	_egee := _ggdd >> 3
	_aaaf := _ggdd & 7
	if _aaaf > 0 {
		_aafg = _ddfe[_aaaf]
	}
	_egae = _edca.BytesPerLine*_dafe + (_bffcb >> 3)
	switch _gfdd {
	case PixClr:
		for _bdcd = 0; _bdcd < _afd; _bdcd++ {
			_ecfg = _egae + _bdcd*_edca.BytesPerLine
			for _ffef = 0; _ffef < _egee; _ffef++ {
				_edca.Data[_ecfg] = 0x0
				_ecfg++
			}
			if _aaaf > 0 {
				_edca.Data[_ecfg] = _agcee(_edca.Data[_ecfg], 0x0, _aafg)
			}
		}
	case PixSet:
		for _bdcd = 0; _bdcd < _afd; _bdcd++ {
			_ecfg = _egae + _bdcd*_edca.BytesPerLine
			for _ffef = 0; _ffef < _egee; _ffef++ {
				_edca.Data[_ecfg] = 0xff
				_ecfg++
			}
			if _aaaf > 0 {
				_edca.Data[_ecfg] = _agcee(_edca.Data[_ecfg], 0xff, _aafg)
			}
		}
	case PixNotDst:
		for _bdcd = 0; _bdcd < _afd; _bdcd++ {
			_ecfg = _egae + _bdcd*_edca.BytesPerLine
			for _ffef = 0; _ffef < _egee; _ffef++ {
				_edca.Data[_ecfg] = ^_edca.Data[_ecfg]
				_ecfg++
			}
			if _aaaf > 0 {
				_edca.Data[_ecfg] = _agcee(_edca.Data[_ecfg], ^_edca.Data[_ecfg], _aafg)
			}
		}
	}
}
func (_aaa *Gray8) Set(x, y int, c _bd.Color) {
	_daebc := y*_aaa.BytesPerLine + x
	if _daebc > len(_aaa.Data)-1 {
		return
	}
	_egbe := _bd.GrayModel.Convert(c)
	_aaa.Data[_daebc] = _egbe.(_bd.Gray).Y
}
func (_dbgac *Monochrome) SetGray(x, y int, g _bd.Gray) {
	_bcbe := y*_dbgac.BytesPerLine + x>>3
	if _bcbe > len(_dbgac.Data)-1 {
		return
	}
	g = _cgb(g, monochromeModel(_dbgac.ModelThreshold))
	_dbgac.setGray(x, g, _bcbe)
}
func ColorAtGray2BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_bd.Gray, error) {
	_cba := y*bytesPerLine + x>>2
	if _cba >= len(data) {
		return _bd.Gray{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_gaff := data[_cba] >> uint(6-(x&3)*2) & 3
	if len(decode) == 2 {
		_gaff = uint8(uint32(LinearInterpolate(float64(_gaff), 0, 3.0, decode[0], decode[1])) & 3)
	}
	return _bd.Gray{Y: _gaff * 85}, nil
}
func _bbea(_cggb _a.Image) (Image, error) {
	if _aeedc, _dcb := _cggb.(*Gray8); _dcb {
		return _aeedc.Copy(), nil
	}
	_caffc := _cggb.Bounds()
	_ecf, _gaec := NewImage(_caffc.Max.X, _caffc.Max.Y, 8, 1, nil, nil, nil)
	if _gaec != nil {
		return nil, _gaec
	}
	_abad(_cggb, _ecf, _caffc)
	return _ecf, nil
}
func (_geefd *NRGBA64) Copy() Image        { return &NRGBA64{ImageBase: _geefd.copy()} }
func (_fegd *Monochrome) IsUnpadded() bool { return (_fegd.Width * _fegd.Height) == len(_fegd.Data) }
func _faef(_cacd _bd.RGBA) _bd.Gray {
	_eac := (19595*uint32(_cacd.R) + 38470*uint32(_cacd.G) + 7471*uint32(_cacd.B) + 1<<7) >> 16
	return _bd.Gray{Y: uint8(_eac)}
}
func (_gaab *NRGBA32) SetNRGBA(x, y int, c _bd.NRGBA) {
	_ffff := y*_gaab.Width + x
	_dbfg := 3 * _ffff
	if _dbfg+2 >= len(_gaab.Data) {
		return
	}
	_gaab.setRGBA(_ffff, c)
}

const (
	_ffda shift = iota
	_aecg
)

func (_fbfe colorConverter) Convert(src _a.Image) (Image, error) { return _fbfe._dbdg(src) }

type RGBA interface {
	RGBAAt(_fbacd, _bdff int) _bd.RGBA
	SetRGBA(_geef, _abfe int, _cde _bd.RGBA)
}

func (_ggce *Gray8) Base() *ImageBase { return &_ggce.ImageBase }

var _ _a.Image = &Gray8{}
var _ _a.Image = &Gray2{}
var _ Gray = &Gray8{}

func (_edga *Gray2) At(x, y int) _bd.Color {
	_ceef, _ := _edga.ColorAt(x, y)
	return _ceef
}
func (_bace *Gray16) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _bace.Width, Y: _bace.Height}}
}
func (_bdca *Gray2) Histogram() (_fcbf [256]int) {
	for _bfc := 0; _bfc < _bdca.Width; _bfc++ {
		for _eee := 0; _eee < _bdca.Height; _eee++ {
			_fcbf[_bdca.GrayAt(_bfc, _eee).Y]++
		}
	}
	return _fcbf
}
func _cddb() (_dde [256]uint64) {
	for _baba := 0; _baba < 256; _baba++ {
		if _baba&0x01 != 0 {
			_dde[_baba] |= 0xff
		}
		if _baba&0x02 != 0 {
			_dde[_baba] |= 0xff00
		}
		if _baba&0x04 != 0 {
			_dde[_baba] |= 0xff0000
		}
		if _baba&0x08 != 0 {
			_dde[_baba] |= 0xff000000
		}
		if _baba&0x10 != 0 {
			_dde[_baba] |= 0xff00000000
		}
		if _baba&0x20 != 0 {
			_dde[_baba] |= 0xff0000000000
		}
		if _baba&0x40 != 0 {
			_dde[_baba] |= 0xff000000000000
		}
		if _baba&0x80 != 0 {
			_dde[_baba] |= 0xff00000000000000
		}
	}
	return _dde
}
func _gagg(_bbdf, _gafe Gray, _cceg _a.Rectangle) {
	for _agcd := 0; _agcd < _cceg.Max.X; _agcd++ {
		for _gabd := 0; _gabd < _cceg.Max.Y; _gabd++ {
			_gafe.SetGray(_agcd, _gabd, _bbdf.GrayAt(_agcd, _gabd))
		}
	}
}
func init() { _eebf() }

type monochromeModel uint8

func _gcb(_fcb RGBA, _fccb CMYK, _bdf _a.Rectangle) {
	for _agbe := 0; _agbe < _bdf.Max.X; _agbe++ {
		for _fae := 0; _fae < _bdf.Max.Y; _fae++ {
			_fbce := _fcb.RGBAAt(_agbe, _fae)
			_fccb.SetCMYK(_agbe, _fae, _cca(_fbce))
		}
	}
}
func (_cfccg *NRGBA16) Copy() Image { return &NRGBA16{ImageBase: _cfccg.copy()} }

var _ Image = &NRGBA64{}

func (_ccad *Monochrome) Base() *ImageBase { return &_ccad.ImageBase }
func (_edea *Monochrome) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _edea.Width, Y: _edea.Height}}
}
func _cbe(_agbed RGBA, _gaba Gray, _dcbc _a.Rectangle) {
	for _dcdf := 0; _dcdf < _dcbc.Max.X; _dcdf++ {
		for _cgfdg := 0; _cgfdg < _dcbc.Max.Y; _cgfdg++ {
			_facc := _faef(_agbed.RGBAAt(_dcdf, _cgfdg))
			_gaba.SetGray(_dcdf, _cgfdg, _facc)
		}
	}
}
func _fdegf(_fgaf CMYK, _gefg Gray, _daed _a.Rectangle) {
	for _afaa := 0; _afaa < _daed.Max.X; _afaa++ {
		for _cgbc := 0; _cgbc < _daed.Max.Y; _cgbc++ {
			_cdgf := _acbf(_fgaf.CMYKAt(_afaa, _cgbc))
			_gefg.SetGray(_afaa, _cgbc, _cdgf)
		}
	}
}
func (_cdaf *Gray4) SetGray(x, y int, g _bd.Gray) {
	if x >= _cdaf.Width || y >= _cdaf.Height {
		return
	}
	g = _efc(g)
	_cdaf.setGray(x, y, g)
}

type Histogramer interface{ Histogram() [256]int }

func (_bcgc *Monochrome) Copy() Image {
	return &Monochrome{ImageBase: _bcgc.ImageBase.copy(), ModelThreshold: _bcgc.ModelThreshold}
}
func (_ddec *NRGBA32) setRGBA(_fbfgd int, _ebbcb _bd.NRGBA) {
	_egec := 3 * _fbfgd
	_ddec.Data[_egec] = _ebbcb.R
	_ddec.Data[_egec+1] = _ebbcb.G
	_ddec.Data[_egec+2] = _ebbcb.B
	if _fbfgd < len(_ddec.Alpha) {
		_ddec.Alpha[_fbfgd] = _ebbcb.A
	}
}
func (_dgbf *NRGBA64) ColorModel() _bd.Model { return _bd.NRGBA64Model }
func FromGoImage(i _a.Image) (Image, error) {
	switch _efge := i.(type) {
	case Image:
		return _efge.Copy(), nil
	case Gray:
		return GrayConverter.Convert(i)
	case *_a.Gray16:
		return Gray16Converter.Convert(i)
	case CMYK:
		return CMYKConverter.Convert(i)
	case *_a.NRGBA64:
		return NRGBA64Converter.Convert(i)
	default:
		return NRGBAConverter.Convert(i)
	}
}
func (_bgad *Gray4) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _bgad.Width, Y: _bgad.Height}}
}
func GetConverter(bitsPerComponent, colorComponents int) (ColorConverter, error) {
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			return MonochromeConverter, nil
		case 2:
			return Gray2Converter, nil
		case 4:
			return Gray4Converter, nil
		case 8:
			return GrayConverter, nil
		case 16:
			return Gray16Converter, nil
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			return NRGBA16Converter, nil
		case 8:
			return NRGBAConverter, nil
		case 16:
			return NRGBA64Converter, nil
		}
	case 4:
		return CMYKConverter, nil
	}
	return nil, _d.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0043o\u006e\u0076\u0065\u0072\u0074\u0065\u0072\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072\u0073\u002e\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003a\u0020\u0025\u0064\u002c\u0020\u0043\u006f\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u003a \u0025\u0064", bitsPerComponent, colorComponents)
}
func (_cgdd *RGBA32) Copy() Image { return &RGBA32{ImageBase: _cgdd.copy()} }
func _efb(_gaa, _cacg int) *Monochrome {
	return &Monochrome{ImageBase: NewImageBase(_gaa, _cacg, 1, 1, nil, nil, nil), ModelThreshold: 0x0f}
}
func (_fdbe *NRGBA64) NRGBA64At(x, y int) _bd.NRGBA64 {
	_ffdd, _ := ColorAtNRGBA64(x, y, _fdbe.Width, _fdbe.Data, _fdbe.Alpha, _fdbe.Decode)
	return _ffdd
}
func (_fbcb *Monochrome) GrayAt(x, y int) _bd.Gray {
	_fgbf, _ := ColorAtGray1BPC(x, y, _fbcb.BytesPerLine, _fbcb.Data, _fbcb.Decode)
	return _fgbf
}
func (_agga *Gray8) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _agga.Width, Y: _agga.Height}}
}
func (_cdbg *Monochrome) ResolveDecode() error {
	if len(_cdbg.Decode) != 2 {
		return nil
	}
	if _cdbg.Decode[0] == 1 && _cdbg.Decode[1] == 0 {
		if _cage := _cdbg.InverseData(); _cage != nil {
			return _cage
		}
		_cdbg.Decode = nil
	}
	return nil
}
func _eae(_bbb _bd.Gray) _bd.NRGBA { return _bd.NRGBA{R: _bbb.Y, G: _bbb.Y, B: _bbb.Y, A: 0xff} }
func (_adbg *Monochrome) getBitAt(_gdge, _daeb int) bool {
	_affe := _daeb*_adbg.BytesPerLine + (_gdge >> 3)
	_bfe := _gdge & 0x07
	_bfd := uint(7 - _bfe)
	if _affe > len(_adbg.Data)-1 {
		return false
	}
	if (_adbg.Data[_affe]>>_bfd)&0x01 >= 1 {
		return true
	}
	return false
}
func (_dagg *NRGBA32) Validate() error {
	if len(_dagg.Data) != 3*_dagg.Width*_dagg.Height {
		return _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _eefe(_efdg *Monochrome, _dbba, _bagb int, _fcad, _dcce int, _agfb RasterOperator, _bcacg *Monochrome, _fgbb, _cfffa int) error {
	var _bfgd, _bfgde, _feef, _cfce int
	if _dbba < 0 {
		_fgbb -= _dbba
		_fcad += _dbba
		_dbba = 0
	}
	if _fgbb < 0 {
		_dbba -= _fgbb
		_fcad += _fgbb
		_fgbb = 0
	}
	_bfgd = _dbba + _fcad - _efdg.Width
	if _bfgd > 0 {
		_fcad -= _bfgd
	}
	_bfgde = _fgbb + _fcad - _bcacg.Width
	if _bfgde > 0 {
		_fcad -= _bfgde
	}
	if _bagb < 0 {
		_cfffa -= _bagb
		_dcce += _bagb
		_bagb = 0
	}
	if _cfffa < 0 {
		_bagb -= _cfffa
		_dcce += _cfffa
		_cfffa = 0
	}
	_feef = _bagb + _dcce - _efdg.Height
	if _feef > 0 {
		_dcce -= _feef
	}
	_cfce = _cfffa + _dcce - _bcacg.Height
	if _cfce > 0 {
		_dcce -= _cfce
	}
	if _fcad <= 0 || _dcce <= 0 {
		return nil
	}
	var _feff error
	switch {
	case _dbba&7 == 0 && _fgbb&7 == 0:
		_feff = _bfb(_efdg, _dbba, _bagb, _fcad, _dcce, _agfb, _bcacg, _fgbb, _cfffa)
	case _dbba&7 == _fgbb&7:
		_feff = _dfcc(_efdg, _dbba, _bagb, _fcad, _dcce, _agfb, _bcacg, _fgbb, _cfffa)
	default:
		_feff = _aaca(_efdg, _dbba, _bagb, _fcad, _dcce, _agfb, _bcacg, _fgbb, _cfffa)
	}
	if _feff != nil {
		return _feff
	}
	return nil
}
func _cee(_bcc, _bge *Monochrome, _ade []byte, _agb int) (_cbc error) {
	var (
		_caf, _cgab, _bdbf, _gfed, _eg, _aeea, _feb, _fbg int
		_acgf, _fde                                       uint32
		_dga, _ccdd                                       byte
		_eadf                                             uint16
	)
	_feaa := make([]byte, 4)
	_eed := make([]byte, 4)
	for _bdbf = 0; _bdbf < _bcc.Height-1; _bdbf, _gfed = _bdbf+2, _gfed+1 {
		_caf = _bdbf * _bcc.BytesPerLine
		_cgab = _gfed * _bge.BytesPerLine
		for _eg, _aeea = 0, 0; _eg < _agb; _eg, _aeea = _eg+4, _aeea+1 {
			for _feb = 0; _feb < 4; _feb++ {
				_fbg = _caf + _eg + _feb
				if _fbg <= len(_bcc.Data)-1 && _fbg < _caf+_bcc.BytesPerLine {
					_feaa[_feb] = _bcc.Data[_fbg]
				} else {
					_feaa[_feb] = 0x00
				}
				_fbg = _caf + _bcc.BytesPerLine + _eg + _feb
				if _fbg <= len(_bcc.Data)-1 && _fbg < _caf+(2*_bcc.BytesPerLine) {
					_eed[_feb] = _bcc.Data[_fbg]
				} else {
					_eed[_feb] = 0x00
				}
			}
			_acgf = _bag.BigEndian.Uint32(_feaa)
			_fde = _bag.BigEndian.Uint32(_eed)
			_fde |= _acgf
			_fde |= _fde << 1
			_fde &= 0xaaaaaaaa
			_acgf = _fde | (_fde << 7)
			_dga = byte(_acgf >> 24)
			_ccdd = byte((_acgf >> 8) & 0xff)
			_fbg = _cgab + _aeea
			if _fbg+1 == len(_bge.Data)-1 || _fbg+1 >= _cgab+_bge.BytesPerLine {
				_bge.Data[_fbg] = _ade[_dga]
			} else {
				_eadf = (uint16(_ade[_dga]) << 8) | uint16(_ade[_ccdd])
				if _cbc = _bge.setTwoBytes(_fbg, _eadf); _cbc != nil {
					return _d.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fbg)
				}
				_aeea++
			}
		}
	}
	return nil
}
func ColorAtGray4BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_bd.Gray, error) {
	_aade := y*bytesPerLine + x>>1
	if _aade >= len(data) {
		return _bd.Gray{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_aedd := data[_aade] >> uint(4-(x&1)*4) & 0xf
	if len(decode) == 2 {
		_aedd = uint8(uint32(LinearInterpolate(float64(_aedd), 0, 15, decode[0], decode[1])) & 0xf)
	}
	return _bd.Gray{Y: _aedd * 17 & 0xff}, nil
}
func (_bdbe *NRGBA64) SetNRGBA64(x, y int, c _bd.NRGBA64) {
	_cacfe := (y*_bdbe.Width + x) * 2
	_eagg := _cacfe * 3
	if _eagg+5 >= len(_bdbe.Data) {
		return
	}
	_bdbe.setNRGBA64(_eagg, c, _cacfe)
}
func ImgToGray(i _a.Image) *_a.Gray {
	if _abcg, _bgcac := i.(*_a.Gray); _bgcac {
		return _abcg
	}
	_acdc := i.Bounds()
	_dfef := _a.NewGray(_acdc)
	for _gfdba := 0; _gfdba < _acdc.Max.X; _gfdba++ {
		for _gagf := 0; _gagf < _acdc.Max.Y; _gagf++ {
			_aegd := i.At(_gfdba, _gagf)
			_dfef.Set(_gfdba, _gagf, _aegd)
		}
	}
	return _dfef
}

var _ Image = &Gray16{}

func ColorAtRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_bd.RGBA, error) {
	_bfdb := y*width + x
	_fbceg := 3 * _bfdb
	if _fbceg+2 >= len(data) {
		return _bd.RGBA{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_afbad := uint8(0xff)
	if alpha != nil && len(alpha) > _bfdb {
		_afbad = alpha[_bfdb]
	}
	_dgbb, _becfe, _dagb := data[_fbceg], data[_fbceg+1], data[_fbceg+2]
	if len(decode) == 6 {
		_dgbb = uint8(uint32(LinearInterpolate(float64(_dgbb), 0, 255, decode[0], decode[1])) & 0xff)
		_becfe = uint8(uint32(LinearInterpolate(float64(_becfe), 0, 255, decode[2], decode[3])) & 0xff)
		_dagb = uint8(uint32(LinearInterpolate(float64(_dagb), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _bd.RGBA{R: _dgbb, G: _becfe, B: _dagb, A: _afbad}, nil
}
func (_bdcc *Gray16) Validate() error {
	if len(_bdcc.Data) != _bdcc.Height*_bdcc.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_ccag *ImageBase) Pix() []byte { return _ccag.Data }
func _dfcc(_bfec *Monochrome, _dcdg, _gggb, _gggg, _eeeg int, _ecff RasterOperator, _gbf *Monochrome, _fdbc, _dgef int) error {
	var (
		_cgcgb      bool
		_cfgd       bool
		_ace        int
		_gceg       int
		_bcf        int
		_bdda       bool
		_acge       byte
		_cbbad      int
		_gga        int
		_aaded      int
		_cfd, _ffea int
	)
	_ebbb := 8 - (_dcdg & 7)
	_fedb := _bbgc[_ebbb]
	_faa := _bfec.BytesPerLine*_gggb + (_dcdg >> 3)
	_aeaa := _gbf.BytesPerLine*_dgef + (_fdbc >> 3)
	if _gggg < _ebbb {
		_cgcgb = true
		_fedb &= _ddfe[8-_ebbb+_gggg]
	}
	if !_cgcgb {
		_ace = (_gggg - _ebbb) >> 3
		if _ace > 0 {
			_cfgd = true
			_gceg = _faa + 1
			_bcf = _aeaa + 1
		}
	}
	_cbbad = (_dcdg + _gggg) & 7
	if !(_cgcgb || _cbbad == 0) {
		_bdda = true
		_acge = _ddfe[_cbbad]
		_gga = _faa + 1 + _ace
		_aaded = _aeaa + 1 + _ace
	}
	switch _ecff {
	case PixSrc:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = _gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotSrc:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^_gbf.Data[_aeaa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = ^_gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^_gbf.Data[_aaded], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixSrcOrDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa]|_bfec.Data[_faa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] |= _gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded]|_bfec.Data[_gga], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixSrcAndDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa]&_bfec.Data[_faa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] &= _gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded]&_bfec.Data[_gga], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixSrcXorDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa]^_bfec.Data[_faa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] ^= _gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded]^_bfec.Data[_gga], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^(_gbf.Data[_aeaa])|_bfec.Data[_faa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] |= ^(_gbf.Data[_bcf+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^(_gbf.Data[_aaded])|_bfec.Data[_gga], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^(_gbf.Data[_aeaa])&_bfec.Data[_faa], _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] &= ^_gbf.Data[_bcf+_ffea]
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^(_gbf.Data[_aaded])&_bfec.Data[_gga], _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa]|^(_bfec.Data[_faa]), _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = _gbf.Data[_bcf+_ffea] | ^(_bfec.Data[_gceg+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded]|^(_bfec.Data[_gga]), _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], _gbf.Data[_aeaa]&^(_bfec.Data[_faa]), _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = _gbf.Data[_bcf+_ffea] &^ (_bfec.Data[_gceg+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], _gbf.Data[_aaded]&^(_bfec.Data[_gga]), _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^(_gbf.Data[_aeaa] | _bfec.Data[_faa]), _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = ^(_gbf.Data[_bcf+_ffea] | _bfec.Data[_gceg+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^(_gbf.Data[_aaded] | _bfec.Data[_gga]), _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^(_gbf.Data[_aeaa] & _bfec.Data[_faa]), _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = ^(_gbf.Data[_bcf+_ffea] & _bfec.Data[_gceg+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^(_gbf.Data[_aaded] & _bfec.Data[_gga]), _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		for _cfd = 0; _cfd < _eeeg; _cfd++ {
			_bfec.Data[_faa] = _agcee(_bfec.Data[_faa], ^(_gbf.Data[_aeaa] ^ _bfec.Data[_faa]), _fedb)
			_faa += _bfec.BytesPerLine
			_aeaa += _gbf.BytesPerLine
		}
		if _cfgd {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				for _ffea = 0; _ffea < _ace; _ffea++ {
					_bfec.Data[_gceg+_ffea] = ^(_gbf.Data[_bcf+_ffea] ^ _bfec.Data[_gceg+_ffea])
				}
				_gceg += _bfec.BytesPerLine
				_bcf += _gbf.BytesPerLine
			}
		}
		if _bdda {
			for _cfd = 0; _cfd < _eeeg; _cfd++ {
				_bfec.Data[_gga] = _agcee(_bfec.Data[_gga], ^(_gbf.Data[_aaded] ^ _bfec.Data[_gga]), _acge)
				_gga += _bfec.BytesPerLine
				_aaded += _gbf.BytesPerLine
			}
		}
	default:
		_e.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _ecff)
		return _b.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_ccb *Gray4) setGray(_aba int, _eccb int, _acad _bd.Gray) {
	_efee := _eccb * _ccb.BytesPerLine
	_dbega := _efee + (_aba >> 1)
	if _dbega >= len(_ccb.Data) {
		return
	}
	_gbbc := _acad.Y >> 4
	_ccb.Data[_dbega] = (_ccb.Data[_dbega] & (^(0xf0 >> uint(4*(_aba&1))))) | (_gbbc << uint(4-4*(_aba&1)))
}
func (_cacf *Monochrome) copy() *Monochrome {
	_fdfg := _efb(_cacf.Width, _cacf.Height)
	_fdfg.ModelThreshold = _cacf.ModelThreshold
	_fdfg.Data = make([]byte, len(_cacf.Data))
	copy(_fdfg.Data, _cacf.Data)
	if len(_cacf.Decode) != 0 {
		_fdfg.Decode = make([]float64, len(_cacf.Decode))
		copy(_fdfg.Decode, _cacf.Decode)
	}
	if len(_cacf.Alpha) != 0 {
		_fdfg.Alpha = make([]byte, len(_cacf.Alpha))
		copy(_fdfg.Alpha, _cacf.Alpha)
	}
	return _fdfg
}
func _ggge(_eda *_a.Gray) bool {
	for _dddb := 0; _dddb < len(_eda.Pix); _dddb++ {
		if !_adbe(_eda.Pix[_dddb]) {
			return false
		}
	}
	return true
}
func (_bbca *RGBA32) At(x, y int) _bd.Color {
	_bege, _ := _bbca.ColorAt(x, y)
	return _bege
}
func _ccde(_aadfg *_a.Gray, _gbfd uint8) *_a.Gray {
	_cbcfg := _aadfg.Bounds()
	_ceed := _a.NewGray(_cbcfg)
	for _fadb := 0; _fadb < _cbcfg.Dx(); _fadb++ {
		for _aadc := 0; _aadc < _cbcfg.Dy(); _aadc++ {
			_ecae := _aadfg.GrayAt(_fadb, _aadc)
			_ceed.SetGray(_fadb, _aadc, _bd.Gray{Y: _bcgff(_ecae.Y, _gbfd)})
		}
	}
	return _ceed
}
func _egeebg(_bbeg CMYK, _eca NRGBA, _fdefe _a.Rectangle) {
	for _gbda := 0; _gbda < _fdefe.Max.X; _gbda++ {
		for _eagd := 0; _eagd < _fdefe.Max.Y; _eagd++ {
			_dcbfa := _bbeg.CMYKAt(_gbda, _eagd)
			_eca.SetNRGBA(_gbda, _eagd, _fgec(_dcbfa))
		}
	}
}
func NewImageBase(width int, height int, bitsPerComponent int, colorComponents int, data []byte, alpha []byte, decode []float64) ImageBase {
	_dbbg := ImageBase{Width: width, Height: height, BitsPerComponent: bitsPerComponent, ColorComponents: colorComponents, Data: data, Alpha: alpha, Decode: decode, BytesPerLine: BytesPerLine(width, bitsPerComponent, colorComponents)}
	if data == nil {
		_dbbg.Data = make([]byte, height*_dbbg.BytesPerLine)
	}
	return _dbbg
}
func (_cagb *Monochrome) ReduceBinary(factor float64) (*Monochrome, error) {
	_gdgc := _ccggd(uint(factor))
	if !IsPowerOf2(uint(factor)) {
		_gdgc++
	}
	_gegd := make([]int, _gdgc)
	for _afe := range _gegd {
		_gegd[_afe] = 4
	}
	_fdef, _gfcf := _bbc(_cagb, _gegd...)
	if _gfcf != nil {
		return nil, _gfcf
	}
	return _fdef, nil
}
func ScaleAlphaToMonochrome(data []byte, width, height int) ([]byte, error) {
	_ef := BytesPerLine(width, 8, 1)
	if len(data) < _ef*height {
		return nil, nil
	}
	_ea := &Gray8{NewImageBase(width, height, 8, 1, data, nil, nil)}
	_af, _db := MonochromeConverter.Convert(_ea)
	if _db != nil {
		return nil, _db
	}
	return _af.Base().Data, nil
}
func _cdcd(_dbee int, _dgaf int) int {
	if _dbee < _dgaf {
		return _dbee
	}
	return _dgaf
}
func (_bdg *NRGBA64) At(x, y int) _bd.Color { _aafac, _ := _bdg.ColorAt(x, y); return _aafac }
func LinearInterpolate(x, xmin, xmax, ymin, ymax float64) float64 {
	if _ba.Abs(xmax-xmin) < 0.000001 {
		return ymin
	}
	_bgff := ymin + (x-xmin)*(ymax-ymin)/(xmax-xmin)
	return _bgff
}
func _beef(_eadca _bd.NRGBA) _bd.Gray {
	_agc, _aea, _cec, _ := _eadca.RGBA()
	_bdba := (19595*_agc + 38470*_aea + 7471*_cec + 1<<15) >> 24
	return _bd.Gray{Y: uint8(_bdba)}
}

type RGBA32 struct{ ImageBase }

func _fgd(_cecb _bd.NYCbCrA) _bd.NRGBA {
	_afa := int32(_cecb.Y) * 0x10101
	_ecdf := int32(_cecb.Cb) - 128
	_fga := int32(_cecb.Cr) - 128
	_gbg := _afa + 91881*_fga
	if uint32(_gbg)&0xff000000 == 0 {
		_gbg >>= 8
	} else {
		_gbg = ^(_gbg >> 31) & 0xffff
	}
	_bbbe := _afa - 22554*_ecdf - 46802*_fga
	if uint32(_bbbe)&0xff000000 == 0 {
		_bbbe >>= 8
	} else {
		_bbbe = ^(_bbbe >> 31) & 0xffff
	}
	_gcfe := _afa + 116130*_ecdf
	if uint32(_gcfe)&0xff000000 == 0 {
		_gcfe >>= 8
	} else {
		_gcfe = ^(_gcfe >> 31) & 0xffff
	}
	return _bd.NRGBA{R: uint8(_gbg >> 8), G: uint8(_bbbe >> 8), B: uint8(_gcfe >> 8), A: _cecb.A}
}
func (_dbea *ImageBase) newAlpha() {
	_gbeb := BytesPerLine(_dbea.Width, _dbea.BitsPerComponent, 1)
	_dbea.Alpha = make([]byte, _dbea.Height*_gbeb)
}
func (_aac *Monochrome) AddPadding() (_dgdf error) {
	if _eaec := ((_aac.Width * _aac.Height) + 7) >> 3; len(_aac.Data) < _eaec {
		return _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(_aac.Data), _eaec)
	}
	_cccg := _aac.Width % 8
	if _cccg == 0 {
		return nil
	}
	_bgf := _aac.Width / 8
	_dbga := _ab.NewReader(_aac.Data)
	_dbda := make([]byte, _aac.Height*_aac.BytesPerLine)
	_cfgb := _ab.NewWriterMSB(_dbda)
	_fdf := make([]byte, _bgf)
	var (
		_edgd int
		_ddb  uint64
	)
	for _edgd = 0; _edgd < _aac.Height; _edgd++ {
		if _, _dgdf = _dbga.Read(_fdf); _dgdf != nil {
			return _dgdf
		}
		if _, _dgdf = _cfgb.Write(_fdf); _dgdf != nil {
			return _dgdf
		}
		if _ddb, _dgdf = _dbga.ReadBits(byte(_cccg)); _dgdf != nil {
			return _dgdf
		}
		if _dgdf = _cfgb.WriteByte(byte(_ddb) << uint(8-_cccg)); _dgdf != nil {
			return _dgdf
		}
	}
	_aac.Data = _cfgb.Data()
	return nil
}
func ColorAtGray1BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_bd.Gray, error) {
	_dcec := y*bytesPerLine + x>>3
	if _dcec >= len(data) {
		return _bd.Gray{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_faf := data[_dcec] >> uint(7-(x&7)) & 1
	if len(decode) == 2 {
		_faf = uint8(LinearInterpolate(float64(_faf), 0.0, 1.0, decode[0], decode[1])) & 1
	}
	return _bd.Gray{Y: _faf * 255}, nil
}
func _agcee(_daab, _gegb, _fdae byte) byte { return (_daab &^ (_fdae)) | (_gegb & _fdae) }
func (_bgaf *Gray16) Set(x, y int, c _bd.Color) {
	_dcbg := (y*_bgaf.BytesPerLine/2 + x) * 2
	if _dcbg+1 >= len(_bgaf.Data) {
		return
	}
	_gbgbc := _bd.Gray16Model.Convert(c).(_bd.Gray16)
	_bgaf.Data[_dcbg], _bgaf.Data[_dcbg+1] = uint8(_gbgbc.Y>>8), uint8(_gbgbc.Y&0xff)
}
func _dbae(_abcc _a.Image, _adgbg uint8) *_a.Gray {
	_afceb := _abcc.Bounds()
	_gbcb := _a.NewGray(_afceb)
	var (
		_ecbd _bd.Color
		_ecg  _bd.Gray
	)
	for _ddbge := 0; _ddbge < _afceb.Max.X; _ddbge++ {
		for _efede := 0; _efede < _afceb.Max.Y; _efede++ {
			_ecbd = _abcc.At(_ddbge, _efede)
			_gbcb.Set(_ddbge, _efede, _ecbd)
			_ecg = _gbcb.GrayAt(_ddbge, _efede)
			_gbcb.SetGray(_ddbge, _efede, _bd.Gray{Y: _bcgff(_ecg.Y, _adgbg)})
		}
	}
	return _gbcb
}
func (_cadd *Gray4) Histogram() (_cccf [256]int) {
	for _aca := 0; _aca < _cadd.Width; _aca++ {
		for _dagda := 0; _dagda < _cadd.Height; _dagda++ {
			_cccf[_cadd.GrayAt(_aca, _dagda).Y]++
		}
	}
	return _cccf
}

var (
	Gray2Model   = _bd.ModelFunc(_feeg)
	Gray4Model   = _bd.ModelFunc(_edf)
	NRGBA16Model = _bd.ModelFunc(_ddbg)
)

func (_cdg *Gray8) At(x, y int) _bd.Color {
	_cede, _ := _cdg.ColorAt(x, y)
	return _cede
}
func (_cbac *Gray4) Validate() error {
	if len(_cbac.Data) != _cbac.Height*_cbac.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_dcbd *RGBA32) setRGBA(_cbgaa int, _cgebf _bd.RGBA) {
	_gabab := 3 * _cbgaa
	_dcbd.Data[_gabab] = _cgebf.R
	_dcbd.Data[_gabab+1] = _cgebf.G
	_dcbd.Data[_gabab+2] = _cgebf.B
	if _cbgaa < len(_dcbd.Alpha) {
		_dcbd.Alpha[_cbgaa] = _cgebf.A
	}
}
func (_gecf *RGBA32) RGBAAt(x, y int) _bd.RGBA {
	_ded, _ := ColorAtRGBA32(x, y, _gecf.Width, _gecf.Data, _gecf.Alpha, _gecf.Decode)
	return _ded
}

type SMasker interface {
	HasAlpha() bool
	GetAlpha() []byte
	MakeAlpha()
}

func (_cdfbg *NRGBA64) setNRGBA64(_dfac int, _cbcff _bd.NRGBA64, _bddfa int) {
	_cdfbg.Data[_dfac] = uint8(_cbcff.R >> 8)
	_cdfbg.Data[_dfac+1] = uint8(_cbcff.R & 0xff)
	_cdfbg.Data[_dfac+2] = uint8(_cbcff.G >> 8)
	_cdfbg.Data[_dfac+3] = uint8(_cbcff.G & 0xff)
	_cdfbg.Data[_dfac+4] = uint8(_cbcff.B >> 8)
	_cdfbg.Data[_dfac+5] = uint8(_cbcff.B & 0xff)
	if _bddfa+1 < len(_cdfbg.Alpha) {
		_cdfbg.Alpha[_bddfa] = uint8(_cbcff.A >> 8)
		_cdfbg.Alpha[_bddfa+1] = uint8(_cbcff.A & 0xff)
	}
}
func (_eccbc *ImageBase) MakeAlpha() { _eccbc.newAlpha() }
func (_egcf *Gray8) Validate() error {
	if len(_egcf.Data) != _egcf.Height*_egcf.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_aeed *Monochrome) ColorModel() _bd.Model { return MonochromeModel(_aeed.ModelThreshold) }
func (_aefg *Gray16) ColorModel() _bd.Model     { return _bd.Gray16Model }
func (_fbgc *Monochrome) InverseData() error {
	return _fbgc.RasterOperation(0, 0, _fbgc.Width, _fbgc.Height, PixNotDst, nil, 0, 0)
}
func (_ffce *Gray4) ColorModel() _bd.Model { return Gray4Model }

var _ Image = &Monochrome{}

func _egd(_efged *Monochrome, _adcc, _bcgg, _fcfc, _eaac int, _aabe RasterOperator) {
	if _adcc < 0 {
		_fcfc += _adcc
		_adcc = 0
	}
	_agff := _adcc + _fcfc - _efged.Width
	if _agff > 0 {
		_fcfc -= _agff
	}
	if _bcgg < 0 {
		_eaac += _bcgg
		_bcgg = 0
	}
	_eaed := _bcgg + _eaac - _efged.Height
	if _eaed > 0 {
		_eaac -= _eaed
	}
	if _fcfc <= 0 || _eaac <= 0 {
		return
	}
	if (_adcc & 7) == 0 {
		_gfee(_efged, _adcc, _bcgg, _fcfc, _eaac, _aabe)
	} else {
		_ebec(_efged, _adcc, _bcgg, _fcfc, _eaac, _aabe)
	}
}
func (_dfd *Gray4) GrayAt(x, y int) _bd.Gray {
	_dfda, _ := ColorAtGray4BPC(x, y, _dfd.BytesPerLine, _dfd.Data, _dfd.Decode)
	return _dfda
}
func (_gbdf *Monochrome) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _dfa(_gbdf, dx, dy, dw, dh, op, src, sx, sy)
}
func (_ggff *RGBA32) SetRGBA(x, y int, c _bd.RGBA) {
	_aagc := y*_ggff.Width + x
	_gdba := 3 * _aagc
	if _gdba+2 >= len(_ggff.Data) {
		return
	}
	_ggff.setRGBA(_aagc, c)
}
func _ddbg(_fcbg _bd.Color) _bd.Color {
	_dcdb := _bd.NRGBAModel.Convert(_fcbg).(_bd.NRGBA)
	return _eefb(_dcdb)
}
func (_cgec *Gray16) At(x, y int) _bd.Color {
	_ecee, _ := _cgec.ColorAt(x, y)
	return _ecee
}

var _ Gray = &Gray16{}

type Monochrome struct {
	ImageBase
	ModelThreshold uint8
}

func (_gdee *monochromeThresholdConverter) Convert(img _a.Image) (Image, error) {
	if _feea, _babc := img.(*Monochrome); _babc {
		return _feea.Copy(), nil
	}
	_bbbf := img.Bounds()
	_ccgd, _abga := NewImage(_bbbf.Max.X, _bbbf.Max.Y, 1, 1, nil, nil, nil)
	if _abga != nil {
		return nil, _abga
	}
	_ccgd.(*Monochrome).ModelThreshold = _gdee.Threshold
	for _cceae := 0; _cceae < _bbbf.Max.X; _cceae++ {
		for _ecef := 0; _ecef < _bbbf.Max.Y; _ecef++ {
			_gddb := img.At(_cceae, _ecef)
			_ccgd.Set(_cceae, _ecef, _gddb)
		}
	}
	return _ccgd, nil
}
func (_fgaa *NRGBA16) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtNRGBA16(x, y, _fgaa.Width, _fgaa.BytesPerLine, _fgaa.Data, _fgaa.Alpha, _fgaa.Decode)
}
func _dba(_adf, _daff *Monochrome, _eea []byte, _ebe int) (_gcg error) {
	var (
		_efbd, _df, _gffd, _fdeg, _aafd, _cdf, _bdbc, _aff int
		_fbdg, _efg, _cad, _gfce                           uint32
		_fdc, _fgf                                         byte
		_ceeb                                              uint16
	)
	_bcb := make([]byte, 4)
	_cacb := make([]byte, 4)
	for _gffd = 0; _gffd < _adf.Height-1; _gffd, _fdeg = _gffd+2, _fdeg+1 {
		_efbd = _gffd * _adf.BytesPerLine
		_df = _fdeg * _daff.BytesPerLine
		for _aafd, _cdf = 0, 0; _aafd < _ebe; _aafd, _cdf = _aafd+4, _cdf+1 {
			for _bdbc = 0; _bdbc < 4; _bdbc++ {
				_aff = _efbd + _aafd + _bdbc
				if _aff <= len(_adf.Data)-1 && _aff < _efbd+_adf.BytesPerLine {
					_bcb[_bdbc] = _adf.Data[_aff]
				} else {
					_bcb[_bdbc] = 0x00
				}
				_aff = _efbd + _adf.BytesPerLine + _aafd + _bdbc
				if _aff <= len(_adf.Data)-1 && _aff < _efbd+(2*_adf.BytesPerLine) {
					_cacb[_bdbc] = _adf.Data[_aff]
				} else {
					_cacb[_bdbc] = 0x00
				}
			}
			_fbdg = _bag.BigEndian.Uint32(_bcb)
			_efg = _bag.BigEndian.Uint32(_cacb)
			_cad = _fbdg & _efg
			_cad |= _cad << 1
			_gfce = _fbdg | _efg
			_gfce &= _gfce << 1
			_efg = _cad | _gfce
			_efg &= 0xaaaaaaaa
			_fbdg = _efg | (_efg << 7)
			_fdc = byte(_fbdg >> 24)
			_fgf = byte((_fbdg >> 8) & 0xff)
			_aff = _df + _cdf
			if _aff+1 == len(_daff.Data)-1 || _aff+1 >= _df+_daff.BytesPerLine {
				if _gcg = _daff.setByte(_aff, _eea[_fdc]); _gcg != nil {
					return _d.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _aff)
				}
			} else {
				_ceeb = (uint16(_eea[_fdc]) << 8) | uint16(_eea[_fgf])
				if _gcg = _daff.setTwoBytes(_aff, _ceeb); _gcg != nil {
					return _d.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _aff)
				}
				_cdf++
			}
		}
	}
	return nil
}
func (_gaee *CMYK32) At(x, y int) _bd.Color        { _cfe, _ := _gaee.ColorAt(x, y); return _cfe }
func (_fbfb *Monochrome) clearBit(_dgac, _fcf int) { _fbfb.Data[_dgac] &= ^(0x80 >> uint(_fcf&7)) }
func _caed(_cbba NRGBA, _feee Gray, _gaece _a.Rectangle) {
	for _bcbeb := 0; _bcbeb < _gaece.Max.X; _bcbeb++ {
		for _ggbg := 0; _ggbg < _gaece.Max.Y; _ggbg++ {
			_cdae := _beef(_cbba.NRGBAAt(_bcbeb, _ggbg))
			_feee.SetGray(_bcbeb, _ggbg, _cdae)
		}
	}
}

type NRGBA64 struct{ ImageBase }

func (_bbce *NRGBA32) Copy() Image           { return &NRGBA32{ImageBase: _bbce.copy()} }
func (_fcdb *NRGBA32) ColorModel() _bd.Model { return _bd.NRGBAModel }
func ColorAt(x, y, width, bitsPerColor, colorComponents, bytesPerLine int, data, alpha []byte, decode []float64) (_bd.Color, error) {
	switch colorComponents {
	case 1:
		return ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine, data, decode)
	case 3:
		return ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor, data, alpha, decode)
	case 4:
		return ColorAtCMYK(x, y, width, data, decode)
	default:
		return nil, _d.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063o\u006c\u006f\u0072\u0020\u0063\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", colorComponents)
	}
}
func (_eaaf *Gray4) At(x, y int) _bd.Color { _feaaf, _ := _eaaf.ColorAt(x, y); return _feaaf }

var _ Image = &CMYK32{}

func (_fcdg *NRGBA16) At(x, y int) _bd.Color { _bcaa, _ := _fcdg.ColorAt(x, y); return _bcaa }
func (_fafa *NRGBA16) Validate() error {
	if len(_fafa.Data) != 3*_fafa.Width*_fafa.Height/2 {
		return _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _aaca(_fffbf *Monochrome, _bgfb, _ffcf, _bffa, _ggbd int, _gbbd RasterOperator, _effbb *Monochrome, _ceag, _aebd int) error {
	var (
		_edcc         bool
		_fgdg         bool
		_cbcd         byte
		_edb          int
		_gfbc         int
		_ebcg         int
		_cebf         int
		_adab         bool
		_fbfa         int
		_aggbb        int
		_afbac        int
		_eecd         bool
		_beede        byte
		_gdcd         int
		_dfdc         int
		_bgcd         int
		_gcee         byte
		_gebgb        int
		_eabf         int
		_fdag         uint
		_ebeg         uint
		_efca         byte
		_adc          shift
		_dbff         bool
		_gcag         bool
		_cccb, _eadfd int
	)
	if _ceag&7 != 0 {
		_eabf = 8 - (_ceag & 7)
	}
	if _bgfb&7 != 0 {
		_gfbc = 8 - (_bgfb & 7)
	}
	if _eabf == 0 && _gfbc == 0 {
		_efca = _bbgc[0]
	} else {
		if _gfbc > _eabf {
			_fdag = uint(_gfbc - _eabf)
		} else {
			_fdag = uint(8 - (_eabf - _gfbc))
		}
		_ebeg = 8 - _fdag
		_efca = _bbgc[_fdag]
	}
	if (_bgfb & 7) != 0 {
		_edcc = true
		_edb = 8 - (_bgfb & 7)
		_cbcd = _bbgc[_edb]
		_ebcg = _fffbf.BytesPerLine*_ffcf + (_bgfb >> 3)
		_cebf = _effbb.BytesPerLine*_aebd + (_ceag >> 3)
		_gebgb = 8 - (_ceag & 7)
		if _edb > _gebgb {
			_adc = _ffda
			if _bffa >= _eabf {
				_dbff = true
			}
		} else {
			_adc = _aecg
		}
	}
	if _bffa < _edb {
		_fgdg = true
		_cbcd &= _ddfe[8-_edb+_bffa]
	}
	if !_fgdg {
		_fbfa = (_bffa - _edb) >> 3
		if _fbfa != 0 {
			_adab = true
			_aggbb = _fffbf.BytesPerLine*_ffcf + ((_bgfb + _gfbc) >> 3)
			_afbac = _effbb.BytesPerLine*_aebd + ((_ceag + _gfbc) >> 3)
		}
	}
	_gdcd = (_bgfb + _bffa) & 7
	if !(_fgdg || _gdcd == 0) {
		_eecd = true
		_beede = _ddfe[_gdcd]
		_dfdc = _fffbf.BytesPerLine*_ffcf + ((_bgfb + _gfbc) >> 3) + _fbfa
		_bgcd = _effbb.BytesPerLine*_aebd + ((_ceag + _gfbc) >> 3) + _fbfa
		if _gdcd > int(_ebeg) {
			_gcag = true
		}
	}
	switch _gbbd {
	case PixSrc:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee, _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = _gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee, _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotSrc:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^_gcee, _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = ^_gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^_gcee, _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixSrcOrDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee|_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] |= _gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee|_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixSrcAndDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee&_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] &= _gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee&_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixSrcXorDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee^_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] ^= _gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee^_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^_gcee|_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] |= ^_gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^_gcee|_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^_gcee&_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] &= ^_gcee
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^_gcee&_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee|^_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = _gcee | ^_fffbf.Data[_aggbb+_eadfd]
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee|^_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], _gcee&^_fffbf.Data[_ebcg], _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = _gcee &^ _fffbf.Data[_aggbb+_eadfd]
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], _gcee&^_fffbf.Data[_dfdc], _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^(_gcee | _fffbf.Data[_ebcg]), _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = ^(_gcee | _fffbf.Data[_aggbb+_eadfd])
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^(_gcee | _fffbf.Data[_dfdc]), _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^(_gcee & _fffbf.Data[_ebcg]), _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = ^(_gcee & _fffbf.Data[_aggbb+_eadfd])
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^(_gcee & _fffbf.Data[_dfdc]), _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		if _edcc {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				if _adc == _ffda {
					_gcee = _effbb.Data[_cebf] << _fdag
					if _dbff {
						_gcee = _agcee(_gcee, _effbb.Data[_cebf+1]>>_ebeg, _efca)
					}
				} else {
					_gcee = _effbb.Data[_cebf] >> _ebeg
				}
				_fffbf.Data[_ebcg] = _agcee(_fffbf.Data[_ebcg], ^(_gcee ^ _fffbf.Data[_ebcg]), _cbcd)
				_ebcg += _fffbf.BytesPerLine
				_cebf += _effbb.BytesPerLine
			}
		}
		if _adab {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				for _eadfd = 0; _eadfd < _fbfa; _eadfd++ {
					_gcee = _agcee(_effbb.Data[_afbac+_eadfd]<<_fdag, _effbb.Data[_afbac+_eadfd+1]>>_ebeg, _efca)
					_fffbf.Data[_aggbb+_eadfd] = ^(_gcee ^ _fffbf.Data[_aggbb+_eadfd])
				}
				_aggbb += _fffbf.BytesPerLine
				_afbac += _effbb.BytesPerLine
			}
		}
		if _eecd {
			for _cccb = 0; _cccb < _ggbd; _cccb++ {
				_gcee = _effbb.Data[_bgcd] << _fdag
				if _gcag {
					_gcee = _agcee(_gcee, _effbb.Data[_bgcd+1]>>_ebeg, _efca)
				}
				_fffbf.Data[_dfdc] = _agcee(_fffbf.Data[_dfdc], ^(_gcee ^ _fffbf.Data[_dfdc]), _beede)
				_dfdc += _fffbf.BytesPerLine
				_bgcd += _effbb.BytesPerLine
			}
		}
	default:
		_e.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _gbbd)
		return _b.New("\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

var _ _a.Image = &Gray4{}

type CMYK32 struct{ ImageBase }

func ConverterFunc(converterFunc func(_cfb _a.Image) (Image, error)) ColorConverter {
	return colorConverter{_dbdg: converterFunc}
}
func _bgg(_dbdfg *_a.NYCbCrA, _abce RGBA, _fffc _a.Rectangle) {
	for _eecg := 0; _eecg < _fffc.Max.X; _eecg++ {
		for _dbad := 0; _dbad < _fffc.Max.Y; _dbad++ {
			_fcfd := _dbdfg.NYCbCrAAt(_eecg, _dbad)
			_abce.SetRGBA(_eecg, _dbad, _eaf(_fcfd))
		}
	}
}
func (_cggf *Monochrome) setIndexedBit(_dcg int) { _cggf.Data[(_dcg >> 3)] |= 0x80 >> uint(_dcg&7) }
func (_bbfd *ImageBase) setEightFullBytes(_fgab int, _feeb uint64) error {
	if _fgab+7 > len(_bbfd.Data)-1 {
		return _b.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bbfd.Data[_fgab] = byte((_feeb & 0xff00000000000000) >> 56)
	_bbfd.Data[_fgab+1] = byte((_feeb & 0xff000000000000) >> 48)
	_bbfd.Data[_fgab+2] = byte((_feeb & 0xff0000000000) >> 40)
	_bbfd.Data[_fgab+3] = byte((_feeb & 0xff00000000) >> 32)
	_bbfd.Data[_fgab+4] = byte((_feeb & 0xff000000) >> 24)
	_bbfd.Data[_fgab+5] = byte((_feeb & 0xff0000) >> 16)
	_bbfd.Data[_fgab+6] = byte((_feeb & 0xff00) >> 8)
	_bbfd.Data[_fgab+7] = byte(_feeb & 0xff)
	return nil
}

var _ Image = &NRGBA16{}

func _gfba(_cea NRGBA, _baga CMYK, _dae _a.Rectangle) {
	for _ege := 0; _ege < _dae.Max.X; _ege++ {
		for _bef := 0; _bef < _dae.Max.Y; _bef++ {
			_gabb := _cea.NRGBAAt(_ege, _bef)
			_baga.SetCMYK(_ege, _bef, _cfg(_gabb))
		}
	}
}

type ImageBase struct {
	Width, Height                     int
	BitsPerComponent, ColorComponents int
	Data, Alpha                       []byte
	Decode                            []float64
	BytesPerLine                      int
}

func _ccea(_gbc _bd.CMYK) _bd.RGBA {
	_gefd, _dcef, _bec := _bd.CMYKToRGB(_gbc.C, _gbc.M, _gbc.Y, _gbc.K)
	return _bd.RGBA{R: _gefd, G: _dcef, B: _bec, A: 0xff}
}
func _fcd() (_eab [256]uint16) {
	for _bdd := 0; _bdd < 256; _bdd++ {
		if _bdd&0x01 != 0 {
			_eab[_bdd] |= 0x3
		}
		if _bdd&0x02 != 0 {
			_eab[_bdd] |= 0xc
		}
		if _bdd&0x04 != 0 {
			_eab[_bdd] |= 0x30
		}
		if _bdd&0x08 != 0 {
			_eab[_bdd] |= 0xc0
		}
		if _bdd&0x10 != 0 {
			_eab[_bdd] |= 0x300
		}
		if _bdd&0x20 != 0 {
			_eab[_bdd] |= 0xc00
		}
		if _bdd&0x40 != 0 {
			_eab[_bdd] |= 0x3000
		}
		if _bdd&0x80 != 0 {
			_eab[_bdd] |= 0xc000
		}
	}
	return _eab
}
func IsGrayImgBlackAndWhite(i *_a.Gray) bool { return _ggge(i) }
func _cfg(_egfa _bd.NRGBA) _bd.CMYK {
	_egc, _eaee, _adfg, _ := _egfa.RGBA()
	_aefc, _eff, _dcf, _fce := _bd.RGBToCMYK(uint8(_egc>>8), uint8(_eaee>>8), uint8(_adfg>>8))
	return _bd.CMYK{C: _aefc, M: _eff, Y: _dcf, K: _fce}
}
func ColorAtGray16BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_bd.Gray16, error) {
	_gcac := (y*bytesPerLine/2 + x) * 2
	if _gcac+1 >= len(data) {
		return _bd.Gray16{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_ebed := uint16(data[_gcac])<<8 | uint16(data[_gcac+1])
	if len(decode) == 2 {
		_ebed = uint16(uint64(LinearInterpolate(float64(_ebed), 0, 65535, decode[0], decode[1])))
	}
	return _bd.Gray16{Y: _ebed}, nil
}
func (_fcec *NRGBA16) setNRGBA(_fdbfd, _ceab, _cafa int, _fbacc _bd.NRGBA) {
	if _fdbfd*3%2 == 0 {
		_fcec.Data[_cafa] = (_fbacc.R>>4)<<4 | (_fbacc.G >> 4)
		_fcec.Data[_cafa+1] = (_fbacc.B>>4)<<4 | (_fcec.Data[_cafa+1] & 0xf)
	} else {
		_fcec.Data[_cafa] = (_fcec.Data[_cafa] & 0xf0) | (_fbacc.R >> 4)
		_fcec.Data[_cafa+1] = (_fbacc.G>>4)<<4 | (_fbacc.B >> 4)
	}
	if _fcec.Alpha != nil {
		_edbg := _ceab * BytesPerLine(_fcec.Width, 4, 1)
		if _edbg < len(_fcec.Alpha) {
			if _fdbfd%2 == 0 {
				_fcec.Alpha[_edbg] = (_fbacc.A>>uint(4))<<uint(4) | (_fcec.Alpha[_cafa] & 0xf)
			} else {
				_fcec.Alpha[_edbg] = (_fcec.Alpha[_edbg] & 0xf0) | (_fbacc.A >> uint(4))
			}
		}
	}
}
func (_efed *Gray16) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtGray16BPC(x, y, _efed.BytesPerLine, _efed.Data, _efed.Decode)
}
func _cdcdd(_cgff _a.Image, _cbdf Image, _gdea _a.Rectangle) {
	if _efcd, _bgaa := _cgff.(SMasker); _bgaa && _efcd.HasAlpha() {
		_cbdf.(SMasker).MakeAlpha()
	}
	switch _bbdfg := _cgff.(type) {
	case Gray:
		_bebb(_bbdfg, _cbdf.(RGBA), _gdea)
	case NRGBA:
		_dbgca(_bbdfg, _cbdf.(RGBA), _gdea)
	case *_a.NYCbCrA:
		_bgg(_bbdfg, _cbdf.(RGBA), _gdea)
	case CMYK:
		_cfad(_bbdfg, _cbdf.(RGBA), _gdea)
	case RGBA:
		_cebe(_bbdfg, _cbdf.(RGBA), _gdea)
	case nrgba64:
		_cagbc(_bbdfg, _cbdf.(RGBA), _gdea)
	default:
		_ebbc(_cgff, _cbdf, _gdea)
	}
}

var _ Image = &Gray4{}
var _ Image = &NRGBA32{}

func _cgb(_cdc _bd.Gray, _gabe monochromeModel) _bd.Gray {
	if _cdc.Y > uint8(_gabe) {
		return _bd.Gray{Y: _ba.MaxUint8}
	}
	return _bd.Gray{}
}
func ImgToBinary(i _a.Image, threshold uint8) *_a.Gray {
	switch _fbcba := i.(type) {
	case *_a.Gray:
		if _ggge(_fbcba) {
			return _fbcba
		}
		return _ccde(_fbcba, threshold)
	case *_a.Gray16:
		return _bgdab(_fbcba, threshold)
	default:
		return _dbae(_fbcba, threshold)
	}
}
func _efc(_fdfc _bd.Gray) _bd.Gray { _fdfc.Y >>= 4; _fdfc.Y |= _fdfc.Y << 4; return _fdfc }
func (_cfgbc *ImageBase) copy() ImageBase {
	_deb := *_cfgbc
	_deb.Data = make([]byte, len(_cfgbc.Data))
	copy(_deb.Data, _cfgbc.Data)
	return _deb
}
func _ccggd(_ffcef uint) uint {
	var _eagcb uint
	for _ffcef != 0 {
		_ffcef >>= 1
		_eagcb++
	}
	return _eagcb - 1
}
func (_gebg *ImageBase) setByte(_cfcc int, _feca byte) error {
	if _cfcc > len(_gebg.Data)-1 {
		return _b.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gebg.Data[_cfcc] = _feca
	return nil
}
func _bgdab(_abgea *_a.Gray16, _affgb uint8) *_a.Gray {
	_dced := _abgea.Bounds()
	_abgca := _a.NewGray(_dced)
	for _gegc := 0; _gegc < _dced.Dx(); _gegc++ {
		for _eefbe := 0; _eefbe < _dced.Dy(); _eefbe++ {
			_dcfc := _abgea.Gray16At(_gegc, _eefbe)
			_abgca.SetGray(_gegc, _eefbe, _bd.Gray{Y: _bcgff(uint8(_dcfc.Y/256), _affgb)})
		}
	}
	return _abgca
}
func (_gebe *NRGBA16) NRGBAAt(x, y int) _bd.NRGBA {
	_cege, _ := ColorAtNRGBA16(x, y, _gebe.Width, _gebe.BytesPerLine, _gebe.Data, _gebe.Alpha, _gebe.Decode)
	return _cege
}
func _gbgb(_dgda _bd.NRGBA64) _bd.NRGBA {
	return _bd.NRGBA{R: uint8(_dgda.R >> 8), G: uint8(_dgda.G >> 8), B: uint8(_dgda.B >> 8), A: uint8(_dgda.A >> 8)}
}

var _ _a.Image = &RGBA32{}

func (_dggc *NRGBA32) NRGBAAt(x, y int) _bd.NRGBA {
	_cfbd, _ := ColorAtNRGBA32(x, y, _dggc.Width, _dggc.Data, _dggc.Alpha, _dggc.Decode)
	return _cfbd
}
func (_bbe *CMYK32) Base() *ImageBase { return &_bbe.ImageBase }
func _ggf(_efec _bd.NRGBA) _bd.RGBA {
	_bga, _cfee, _ceaa, _acfg := _efec.RGBA()
	return _bd.RGBA{R: uint8(_bga >> 8), G: uint8(_cfee >> 8), B: uint8(_ceaa >> 8), A: uint8(_acfg >> 8)}
}

type nrgba64 interface {
	NRGBA64At(_bbdfd, _gdbf int) _bd.NRGBA64
	SetNRGBA64(_deab, _gfbd int, _ffad _bd.NRGBA64)
}
type NRGBA interface {
	NRGBAAt(_fbbc, _bagbg int) _bd.NRGBA
	SetNRGBA(_bcef, _agcdcf int, _deeb _bd.NRGBA)
}

func (_bcca *Monochrome) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtGray1BPC(x, y, _bcca.BytesPerLine, _bcca.Data, _bcca.Decode)
}

type colorConverter struct {
	_dbdg func(_ddc _a.Image) (Image, error)
}

func AddDataPadding(width, height, bitsPerComponent, colorComponents int, data []byte) ([]byte, error) {
	_bfgea := BytesPerLine(width, bitsPerComponent, colorComponents)
	if _bfgea == width*colorComponents*bitsPerComponent/8 {
		return data, nil
	}
	_fafb := width * colorComponents * bitsPerComponent
	_dbdgb := _bfgea * 8
	_cafd := 8 - (_dbdgb - _fafb)
	_affa := _ab.NewReader(data)
	_eaff := _bfgea - 1
	_bbbeb := make([]byte, _eaff)
	_gegdc := make([]byte, height*_bfgea)
	_acgd := _ab.NewWriterMSB(_gegdc)
	var _bddf uint64
	var _afee error
	for _bdbg := 0; _bdbg < height; _bdbg++ {
		_, _afee = _affa.Read(_bbbeb)
		if _afee != nil {
			return nil, _afee
		}
		_, _afee = _acgd.Write(_bbbeb)
		if _afee != nil {
			return nil, _afee
		}
		_bddf, _afee = _affa.ReadBits(byte(_cafd))
		if _afee != nil {
			return nil, _afee
		}
		_, _afee = _acgd.WriteBits(_bddf, _cafd)
		if _afee != nil {
			return nil, _afee
		}
		_acgd.FinishByte()
	}
	return _gegdc, nil
}
func _eefb(_beca _bd.NRGBA) _bd.NRGBA {
	_beca.R = _beca.R>>4 | (_beca.R>>4)<<4
	_beca.G = _beca.G>>4 | (_beca.G>>4)<<4
	_beca.B = _beca.B>>4 | (_beca.B>>4)<<4
	return _beca
}
func _cfa(_bf _a.Image) (Image, error) {
	if _abb, _acc := _bf.(*CMYK32); _acc {
		return _abb.Copy(), nil
	}
	_adg := _bf.Bounds()
	_eegb, _cdde := NewImage(_adg.Max.X, _adg.Max.Y, 8, 4, nil, nil, nil)
	if _cdde != nil {
		return nil, _cdde
	}
	switch _bcg := _bf.(type) {
	case CMYK:
		_bca(_bcg, _eegb.(CMYK), _adg)
	case Gray:
		_bgba(_bcg, _eegb.(CMYK), _adg)
	case NRGBA:
		_gfba(_bcg, _eegb.(CMYK), _adg)
	case RGBA:
		_gcb(_bcg, _eegb.(CMYK), _adg)
	default:
		_ebbc(_bf, _eegb, _adg)
	}
	return _eegb, nil
}
func RasterOperation(dest *Monochrome, dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _dfa(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func (_cdbgd *Gray8) Histogram() (_cdgb [256]int) {
	for _fedd := 0; _fedd < len(_cdbgd.Data); _fedd++ {
		_cdgb[_cdbgd.Data[_fedd]]++
	}
	return _cdgb
}

type monochromeThresholdConverter struct{ Threshold uint8 }

func _bca(_bagf, _fge CMYK, _dbaa _a.Rectangle) {
	for _gfec := 0; _gfec < _dbaa.Max.X; _gfec++ {
		for _bad := 0; _bad < _dbaa.Max.Y; _bad++ {
			_fge.SetCMYK(_gfec, _bad, _bagf.CMYKAt(_gfec, _bad))
		}
	}
}

var _ Gray = &Gray2{}

func (_gdgg *CMYK32) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtCMYK(x, y, _gdgg.Width, _gdgg.Data, _gdgg.Decode)
}
func _effb(_efdfc Gray, _aaeb NRGBA, _bffc _a.Rectangle) {
	for _cefe := 0; _cefe < _bffc.Max.X; _cefe++ {
		for _bgc := 0; _bgc < _bffc.Max.Y; _bgc++ {
			_dgaa := _gde(_aaeb.NRGBAAt(_cefe, _bgc))
			_efdfc.SetGray(_cefe, _bgc, _dgaa)
		}
	}
}
func (_eeb *Gray2) GrayAt(x, y int) _bd.Gray {
	_fdab, _ := ColorAtGray2BPC(x, y, _eeb.BytesPerLine, _eeb.Data, _eeb.Decode)
	return _fdab
}
func NewImage(width, height, bitsPerComponent, colorComponents int, data, alpha []byte, decode []float64) (Image, error) {
	_fgc := NewImageBase(width, height, bitsPerComponent, colorComponents, data, alpha, decode)
	var _adfd Image
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			_adfd = &Monochrome{ImageBase: _fgc, ModelThreshold: 0x0f}
		case 2:
			_adfd = &Gray2{ImageBase: _fgc}
		case 4:
			_adfd = &Gray4{ImageBase: _fgc}
		case 8:
			_adfd = &Gray8{ImageBase: _fgc}
		case 16:
			_adfd = &Gray16{ImageBase: _fgc}
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			_adfd = &NRGBA16{ImageBase: _fgc}
		case 8:
			_adfd = &NRGBA32{ImageBase: _fgc}
		case 16:
			_adfd = &NRGBA64{ImageBase: _fgc}
		}
	case 4:
		_adfd = &CMYK32{ImageBase: _fgc}
	}
	if _adfd == nil {
		return nil, ErrInvalidImage
	}
	return _adfd, nil
}
func (_afbd *NRGBA16) Base() *ImageBase { return &_afbd.ImageBase }

type Gray8 struct{ ImageBase }

func _ffca(_fba _bd.NRGBA64) _bd.RGBA {
	_ccf, _cbd, _dgag, _cab := _fba.RGBA()
	return _bd.RGBA{R: uint8(_ccf >> 8), G: uint8(_cbd >> 8), B: uint8(_dgag >> 8), A: uint8(_cab >> 8)}
}
func BytesPerLine(width, bitsPerComponent, colorComponents int) int {
	return ((width*bitsPerComponent)*colorComponents + 7) >> 3
}
func _dcafa(_bfa []byte, _cgccf Image) error {
	_afdb := true
	for _aadf := 0; _aadf < len(_bfa); _aadf++ {
		if _bfa[_aadf] != 0xff {
			_afdb = false
			break
		}
	}
	if _afdb {
		switch _cgba := _cgccf.(type) {
		case *NRGBA32:
			_cgba.Alpha = nil
		case *NRGBA64:
			_cgba.Alpha = nil
		default:
			return _d.Errorf("i\u006ete\u0072n\u0061l\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020i\u006d\u0061\u0067\u0065\u0020s\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070e\u0020\u002a\u004eRGB\u0041\u0033\u0032\u0020\u006f\u0072 \u002a\u004e\u0052\u0047\u0042\u0041\u0036\u0034\u0020\u0062\u0075\u0074 \u0069s\u003a\u0020\u0025\u0054", _cgccf)
		}
	}
	return nil
}
func _aagg(_dbbe _a.Image, _eagf int) (_a.Rectangle, bool, []byte) {
	_cefg := _dbbe.Bounds()
	var (
		_gfdg bool
		_addf []byte
	)
	switch _bbgcd := _dbbe.(type) {
	case SMasker:
		_gfdg = _bbgcd.HasAlpha()
	case NRGBA, RGBA, *_a.RGBA64, nrgba64, *_a.NYCbCrA:
		_addf = make([]byte, _cefg.Max.X*_cefg.Max.Y*_eagf)
	case *_a.Paletted:
		if !_bbgcd.Opaque() {
			_addf = make([]byte, _cefg.Max.X*_cefg.Max.Y*_eagf)
		}
	}
	return _cefg, _gfdg, _addf
}
func _ad(_aee *Monochrome, _cbb, _bc int) (*Monochrome, error) {
	if _aee == nil {
		return nil, _b.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _cbb <= 0 || _bc <= 0 {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _cbb == _bc {
		if _cbb == 1 {
			return _aee.copy(), nil
		}
		if _cbb == 2 || _cbb == 4 || _cbb == 8 {
			_fa, _abg := _gc(_aee, _cbb)
			if _abg != nil {
				return nil, _abg
			}
			return _fa, nil
		}
	}
	_gg := _cbb * _aee.Width
	_aaf := _bc * _aee.Height
	_gfc := _efb(_gg, _aaf)
	_fea := _gfc.BytesPerLine
	var (
		_ffa, _cac, _acg, _ddf, _fff int
		_beb                         byte
		_cd                          error
	)
	for _cac = 0; _cac < _aee.Height; _cac++ {
		_ffa = _bc * _cac * _fea
		for _acg = 0; _acg < _aee.Width; _acg++ {
			if _aad := _aee.getBitAt(_acg, _cac); _aad {
				_fff = _cbb * _acg
				for _ddf = 0; _ddf < _cbb; _ddf++ {
					_gfc.setIndexedBit(_ffa*8 + _fff + _ddf)
				}
			}
		}
		for _ddf = 1; _ddf < _bc; _ddf++ {
			_daf := _ffa + _ddf*_fea
			for _bg := 0; _bg < _fea; _bg++ {
				if _beb, _cd = _gfc.getByte(_ffa + _bg); _cd != nil {
					return nil, _cd
				}
				if _cd = _gfc.setByte(_daf+_bg, _beb); _cd != nil {
					return nil, _cd
				}
			}
		}
	}
	return _gfc, nil
}
func _dfa(_egea *Monochrome, _fgfe, _ebac, _ega, _gcdb int, _aeaf RasterOperator, _gecb *Monochrome, _ddef, _afgg int) error {
	if _egea == nil {
		return _b.New("\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _aeaf == PixDst {
		return nil
	}
	switch _aeaf {
	case PixClr, PixSet, PixNotDst:
		_egd(_egea, _fgfe, _ebac, _ega, _gcdb, _aeaf)
		return nil
	}
	if _gecb == nil {
		_e.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _b.New("\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _beda := _eefe(_egea, _fgfe, _ebac, _ega, _gcdb, _aeaf, _gecb, _ddef, _afgg); _beda != nil {
		return _beda
	}
	return nil
}
func _cagbc(_bgdfe nrgba64, _dgdd RGBA, _fgbcf _a.Rectangle) {
	for _daaf := 0; _daaf < _fgbcf.Max.X; _daaf++ {
		for _ccff := 0; _ccff < _fgbcf.Max.Y; _ccff++ {
			_feaf := _bgdfe.NRGBA64At(_daaf, _ccff)
			_dgdd.SetRGBA(_daaf, _ccff, _ffca(_feaf))
		}
	}
}
func _dcda(_ebg _a.Image, _fccc Image, _faad _a.Rectangle) {
	if _edgf, _eabc := _ebg.(SMasker); _eabc && _edgf.HasAlpha() {
		_fccc.(SMasker).MakeAlpha()
	}
	switch _aece := _ebg.(type) {
	case Gray:
		_bgdc(_aece, _fccc.(NRGBA), _faad)
	case NRGBA:
		_cbgc(_aece, _fccc.(NRGBA), _faad)
	case *_a.NYCbCrA:
		_adebe(_aece, _fccc.(NRGBA), _faad)
	case CMYK:
		_egeebg(_aece, _fccc.(NRGBA), _faad)
	case RGBA:
		_effc(_aece, _fccc.(NRGBA), _faad)
	case nrgba64:
		_aebbc(_aece, _fccc.(NRGBA), _faad)
	default:
		_ebbc(_ebg, _fccc, _faad)
	}
}
func (_cdfb *Gray8) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtGray8BPC(x, y, _cdfb.BytesPerLine, _cdfb.Data, _cdfb.Decode)
}
func _cfad(_gfdbg CMYK, _fgeca RGBA, _ffae _a.Rectangle) {
	for _decf := 0; _decf < _ffae.Max.X; _decf++ {
		for _eggf := 0; _eggf < _ffae.Max.Y; _eggf++ {
			_fbef := _gfdbg.CMYKAt(_decf, _eggf)
			_fgeca.SetRGBA(_decf, _eggf, _ccea(_fbef))
		}
	}
}

var _ _a.Image = &Monochrome{}

func (_efa *Gray4) Copy() Image { return &Gray4{ImageBase: _efa.copy()} }
func _ccfb(_bcge _a.Image) (Image, error) {
	if _fab, _egeg := _bcge.(*Gray4); _egeg {
		return _fab.Copy(), nil
	}
	_fcdd := _bcge.Bounds()
	_gafbg, _fgbfd := NewImage(_fcdd.Max.X, _fcdd.Max.Y, 4, 1, nil, nil, nil)
	if _fgbfd != nil {
		return nil, _fgbfd
	}
	_abad(_bcge, _gafbg, _fcdd)
	return _gafbg, nil
}
func (_cgce *Monochrome) Validate() error {
	if len(_cgce.Data) != _cgce.Height*_cgce.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

var _ _a.Image = &Gray16{}

func ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine int, data []byte, decode []float64) (_bd.Color, error) {
	switch bitsPerColor {
	case 1:
		return ColorAtGray1BPC(x, y, bytesPerLine, data, decode)
	case 2:
		return ColorAtGray2BPC(x, y, bytesPerLine, data, decode)
	case 4:
		return ColorAtGray4BPC(x, y, bytesPerLine, data, decode)
	case 8:
		return ColorAtGray8BPC(x, y, bytesPerLine, data, decode)
	case 16:
		return ColorAtGray16BPC(x, y, bytesPerLine, data, decode)
	default:
		return nil, _d.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0067\u0072\u0061\u0079\u0020\u0073c\u0061\u006c\u0065\u0020\u0062\u0069\u0074s\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020a\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}
func _bebd(_ccdc _bd.Gray) _bd.CMYK { return _bd.CMYK{K: 0xff - _ccdc.Y} }

var (
	_ddfe = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_bbgc = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func _ebbc(_aeef _a.Image, _ccc Image, _ebde _a.Rectangle) {
	for _dcac := 0; _dcac < _ebde.Max.X; _dcac++ {
		for _eegc := 0; _eegc < _ebde.Max.Y; _eegc++ {
			_eadc := _aeef.At(_dcac, _eegc)
			_ccc.Set(_dcac, _eegc, _eadc)
		}
	}
}
func _gegbc(_ddg _a.Image) (Image, error) {
	if _aebac, _gcad := _ddg.(*NRGBA64); _gcad {
		return _aebac.Copy(), nil
	}
	_fag, _fafac, _dcgc := _aagg(_ddg, 2)
	_afeb, _bbbee := NewImage(_fag.Max.X, _fag.Max.Y, 16, 3, nil, _dcgc, nil)
	if _bbbee != nil {
		return nil, _bbbee
	}
	_abee(_ddg, _afeb, _fag)
	if len(_dcgc) != 0 && !_fafac {
		if _afada := _dcafa(_dcgc, _afeb); _afada != nil {
			return nil, _afada
		}
	}
	return _afeb, nil
}
func (_ebc *ImageBase) getByte(_face int) (byte, error) {
	if _face > len(_ebc.Data)-1 || _face < 0 {
		return 0, _d.Errorf("\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", _face)
	}
	return _ebc.Data[_face], nil
}
func (_fccf *Gray2) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _fccf.Width, Y: _fccf.Height}}
}
func MonochromeThresholdConverter(threshold uint8) ColorConverter {
	return &monochromeThresholdConverter{Threshold: threshold}
}

var _ RGBA = &RGBA32{}

func _bcgff(_fdcg, _afdg uint8) uint8 {
	if _fdcg < _afdg {
		return 255
	}
	return 0
}

var _ NRGBA = &NRGBA32{}

func (_baa *CMYK32) CMYKAt(x, y int) _bd.CMYK {
	_gca, _ := ColorAtCMYK(x, y, _baa.Width, _baa.Data, _baa.Decode)
	return _gca
}
func (_efecg *RGBA32) Set(x, y int, c _bd.Color) {
	_aefa := y*_efecg.Width + x
	_bagc := 3 * _aefa
	if _bagc+2 >= len(_efecg.Data) {
		return
	}
	_aaced := _bd.RGBAModel.Convert(c).(_bd.RGBA)
	_efecg.setRGBA(_aefa, _aaced)
}
func (_fef *Monochrome) setGrayBit(_ced, _bfge int) { _fef.Data[_ced] |= 0x80 >> uint(_bfge&7) }
func (_ebef *Monochrome) Histogram() (_adec [256]int) {
	for _, _accf := range _ebef.Data {
		_adec[0xff] += int(_ebbd[_ebef.Data[_accf]])
	}
	return _adec
}

var _ _a.Image = &NRGBA64{}

func _cca(_dgb _bd.RGBA) _bd.CMYK {
	_aedb, _egb, _beed, _fgea := _bd.RGBToCMYK(_dgb.R, _dgb.G, _dgb.B)
	return _bd.CMYK{C: _aedb, M: _egb, Y: _beed, K: _fgea}
}
func _bgdc(_ffee Gray, _cfdd NRGBA, _cadb _a.Rectangle) {
	for _cadf := 0; _cadf < _cadb.Max.X; _cadf++ {
		for _cffga := 0; _cffga < _cadb.Max.Y; _cffga++ {
			_gcba := _ffee.GrayAt(_cadf, _cffga)
			_cfdd.SetNRGBA(_cadf, _cffga, _eae(_gcba))
		}
	}
}
func _dbc(_cg, _bbf *Monochrome) (_ee error) {
	_ca := _bbf.BytesPerLine
	_ead := _cg.BytesPerLine
	var (
		_gfb                      byte
		_ede                      uint16
		_ae, _ff, _ge, _dbd, _bac int
	)
	for _ge = 0; _ge < _bbf.Height; _ge++ {
		_ae = _ge * _ca
		_ff = 2 * _ge * _ead
		for _dbd = 0; _dbd < _ca; _dbd++ {
			_gfb = _bbf.Data[_ae+_dbd]
			_ede = _cga[_gfb]
			_bac = _ff + _dbd*2
			if _cg.BytesPerLine != _bbf.BytesPerLine*2 && (_dbd+1)*2 > _cg.BytesPerLine {
				_ee = _cg.setByte(_bac, byte(_ede>>8))
			} else {
				_ee = _cg.setTwoBytes(_bac, _ede)
			}
			if _ee != nil {
				return _ee
			}
		}
		for _dbd = 0; _dbd < _ead; _dbd++ {
			_bac = _ff + _ead + _dbd
			_gfb = _cg.Data[_ff+_dbd]
			if _ee = _cg.setByte(_bac, _gfb); _ee != nil {
				return _ee
			}
		}
	}
	return nil
}
func (_afg *Monochrome) setGray(_aeeea int, _dbgc _bd.Gray, _dgg int) {
	if _dbgc.Y == 0 {
		_afg.clearBit(_dgg, _aeeea)
	} else {
		_afg.setGrayBit(_dgg, _aeeea)
	}
}
func (_fbe *Gray8) ColorModel() _bd.Model { return _bd.GrayModel }
func (_gaea *NRGBA64) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtNRGBA64(x, y, _gaea.Width, _gaea.Data, _gaea.Alpha, _gaea.Decode)
}

type Gray4 struct{ ImageBase }

func _fg() (_fgg [256]uint32) {
	for _cdd := 0; _cdd < 256; _cdd++ {
		if _cdd&0x01 != 0 {
			_fgg[_cdd] |= 0xf
		}
		if _cdd&0x02 != 0 {
			_fgg[_cdd] |= 0xf0
		}
		if _cdd&0x04 != 0 {
			_fgg[_cdd] |= 0xf00
		}
		if _cdd&0x08 != 0 {
			_fgg[_cdd] |= 0xf000
		}
		if _cdd&0x10 != 0 {
			_fgg[_cdd] |= 0xf0000
		}
		if _cdd&0x20 != 0 {
			_fgg[_cdd] |= 0xf00000
		}
		if _cdd&0x40 != 0 {
			_fgg[_cdd] |= 0xf000000
		}
		if _cdd&0x80 != 0 {
			_fgg[_cdd] |= 0xf0000000
		}
	}
	return _fgg
}
func (_gedd *NRGBA64) Validate() error {
	if len(_gedd.Data) != 3*2*_gedd.Width*_gedd.Height {
		return _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_bgfd *NRGBA32) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtNRGBA32(x, y, _bgfd.Width, _bgfd.Data, _bgfd.Alpha, _bgfd.Decode)
}
func InDelta(expected, current, delta float64) bool {
	_ecdc := expected - current
	if _ecdc <= -delta || _ecdc >= delta {
		return false
	}
	return true
}
func (_bgd *CMYK32) ColorModel() _bd.Model { return _bd.CMYKModel }
func GrayHistogram(g Gray) (_faba [256]int) {
	switch _ccdce := g.(type) {
	case Histogramer:
		return _ccdce.Histogram()
	case _a.Image:
		_ecbfg := _ccdce.Bounds()
		for _aabb := 0; _aabb < _ecbfg.Max.X; _aabb++ {
			for _bebg := 0; _bebg < _ecbfg.Max.Y; _bebg++ {
				_faba[g.GrayAt(_aabb, _bebg).Y]++
			}
		}
		return _faba
	default:
		return [256]int{}
	}
}
func (_eegcg *Gray16) GrayAt(x, y int) _bd.Gray {
	_dfec, _ := _eegcg.ColorAt(x, y)
	return _bd.Gray{Y: uint8(_dfec.(_bd.Gray16).Y >> 8)}
}
func _gde(_dgd _bd.NRGBA) _bd.Gray {
	var _bff _bd.NRGBA
	if _dgd == _bff {
		return _bd.Gray{Y: 0xff}
	}
	_efe, _cgeb, _cgc, _ := _dgd.RGBA()
	_bce := (19595*_efe + 38470*_cgeb + 7471*_cgc + 1<<15) >> 24
	return _bd.Gray{Y: uint8(_bce)}
}
func _aa(_cb *Monochrome, _dd int, _eb []uint) (*Monochrome, error) {
	_ed := _dd * _cb.Width
	_gfe := _dd * _cb.Height
	_f := _efb(_ed, _gfe)
	for _fd, _ga := range _eb {
		var _dbb error
		switch _ga {
		case 2:
			_dbb = _dbc(_f, _cb)
		case 4:
			_dbb = _ec(_f, _cb)
		case 8:
			_dbb = _ce(_f, _cb)
		}
		if _dbb != nil {
			return nil, _dbb
		}
		if _fd != len(_eb)-1 {
			_cb = _f.copy()
		}
	}
	return _f, nil
}

var (
	MonochromeConverter = ConverterFunc(_badc)
	Gray2Converter      = ConverterFunc(_aagb)
	Gray4Converter      = ConverterFunc(_ccfb)
	GrayConverter       = ConverterFunc(_bbea)
	Gray16Converter     = ConverterFunc(_dfcd)
	NRGBA16Converter    = ConverterFunc(_dade)
	NRGBAConverter      = ConverterFunc(_fgga)
	NRGBA64Converter    = ConverterFunc(_gegbc)
	RGBAConverter       = ConverterFunc(_bged)
	CMYKConverter       = ConverterFunc(_cfa)
)

func (_gag *Monochrome) setBit(_cgf, _gfg int) { _gag.Data[_cgf+(_gfg>>3)] |= 0x80 >> uint(_gfg&7) }
func (_abbg *Gray2) Base() *ImageBase          { return &_abbg.ImageBase }
func (_gfeg *Gray2) Validate() error {
	if len(_gfeg.Data) != _gfeg.Height*_gfeg.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_gcab *NRGBA64) Base() *ImageBase { return &_gcab.ImageBase }
func _badc(_gbd _a.Image) (Image, error) {
	if _baf, _bba := _gbd.(*Monochrome); _bba {
		return _baf, nil
	}
	_fceb := _gbd.Bounds()
	var _dea Gray
	switch _geff := _gbd.(type) {
	case Gray:
		_dea = _geff
	case NRGBA:
		_dea = &Gray8{ImageBase: NewImageBase(_fceb.Max.X, _fceb.Max.Y, 8, 1, nil, nil, nil)}
		_effb(_dea, _geff, _fceb)
	case nrgba64:
		_dea = &Gray8{ImageBase: NewImageBase(_fceb.Max.X, _fceb.Max.Y, 8, 1, nil, nil, nil)}
		_bgdf(_dea, _geff, _fceb)
	default:
		_dge, _bgda := GrayConverter.Convert(_gbd)
		if _bgda != nil {
			return nil, _bgda
		}
		_dea = _dge.(Gray)
	}
	_ggga, _abc := NewImage(_fceb.Max.X, _fceb.Max.Y, 1, 1, nil, nil, nil)
	if _abc != nil {
		return nil, _abc
	}
	_ged := _ggga.(*Monochrome)
	_gcdd := AutoThresholdTriangle(GrayHistogram(_dea))
	for _agef := 0; _agef < _fceb.Max.X; _agef++ {
		for _fca := 0; _fca < _fceb.Max.Y; _fca++ {
			_dagf := _cgb(_dea.GrayAt(_agef, _fca), monochromeModel(_gcdd))
			_ged.SetGray(_agef, _fca, _dagf)
		}
	}
	return _ggga, nil
}
func (_edbf *RGBA32) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtRGBA32(x, y, _edbf.Width, _edbf.Data, _edbf.Alpha, _edbf.Decode)
}
func (_gaggd *NRGBA16) ColorModel() _bd.Model { return NRGBA16Model }
func (_eddc *Gray2) SetGray(x, y int, gray _bd.Gray) {
	_afef := _abda(gray)
	_afc := y * _eddc.BytesPerLine
	_geebd := _afc + (x >> 2)
	if _geebd >= len(_eddc.Data) {
		return
	}
	_cfba := _afef.Y >> 6
	_eddc.Data[_geebd] = (_eddc.Data[_geebd] & (^(0xc0 >> uint(2*((x)&3))))) | (_cfba << uint(6-2*(x&3)))
}
func (_gcf *CMYK32) Set(x, y int, c _bd.Color) {
	_dagd := 4 * (y*_gcf.Width + x)
	if _dagd+3 >= len(_gcf.Data) {
		return
	}
	_agbd := _bd.CMYKModel.Convert(c).(_bd.CMYK)
	_gcf.Data[_dagd] = _agbd.C
	_gcf.Data[_dagd+1] = _agbd.M
	_gcf.Data[_dagd+2] = _agbd.Y
	_gcf.Data[_dagd+3] = _agbd.K
}
func (_abef *Gray4) Set(x, y int, c _bd.Color) {
	if x >= _abef.Width || y >= _abef.Height {
		return
	}
	_edce := Gray4Model.Convert(c).(_bd.Gray)
	_abef.setGray(x, y, _edce)
}
func _abad(_cdcb _a.Image, _cefa Image, _bfdc _a.Rectangle) {
	switch _aacb := _cdcb.(type) {
	case Gray:
		_gagg(_aacb, _cefa.(Gray), _bfdc)
	case NRGBA:
		_caed(_aacb, _cefa.(Gray), _bfdc)
	case CMYK:
		_fdegf(_aacb, _cefa.(Gray), _bfdc)
	case RGBA:
		_cbe(_aacb, _cefa.(Gray), _bfdc)
	default:
		_ebbc(_cdcb, _cefa, _bfdc)
	}
}
func (_gdc *Gray16) Histogram() (_cbga [256]int) {
	for _gbeg := 0; _gbeg < _gdc.Width; _gbeg++ {
		for _bccb := 0; _bccb < _gdc.Height; _bccb++ {
			_cbga[_gdc.GrayAt(_gbeg, _bccb).Y]++
		}
	}
	return _cbga
}
func (_cddg *Monochrome) Set(x, y int, c _bd.Color) {
	_ffe := y*_cddg.BytesPerLine + x>>3
	if _ffe > len(_cddg.Data)-1 {
		return
	}
	_gbe := _cddg.ColorModel().Convert(c).(_bd.Gray)
	_cddg.setGray(x, _gbe, _ffe)
}
func _acbf(_acga _bd.CMYK) _bd.Gray {
	_accg, _bbff, _edd := _bd.CMYKToRGB(_acga.C, _acga.M, _acga.Y, _acga.K)
	_edg := (19595*uint32(_accg) + 38470*uint32(_bbff) + 7471*uint32(_edd) + 1<<7) >> 16
	return _bd.Gray{Y: uint8(_edg)}
}
func (_ggbb *ImageBase) GetAlpha() []byte { return _ggbb.Alpha }
func (_abd *Monochrome) Scale(scale float64) (*Monochrome, error) {
	var _gedf bool
	_gbec := scale
	if scale < 1 {
		_gbec = 1 / scale
		_gedf = true
	}
	_cbcf := NextPowerOf2(uint(_gbec))
	if InDelta(float64(_cbcf), _gbec, 0.001) {
		if _gedf {
			return _abd.ReduceBinary(_gbec)
		}
		return _abd.ExpandBinary(int(_cbcf))
	}
	_eadg := int(_ba.RoundToEven(float64(_abd.Width) * scale))
	_fcba := int(_ba.RoundToEven(float64(_abd.Height) * scale))
	return _abd.ScaleLow(_eadg, _fcba)
}

var _ Gray = &Gray4{}

func _cag(_cef, _gee *Monochrome, _fbgd []byte, _fbc int) (_fec error) {
	var (
		_gffdd, _cge, _dab, _gba, _gec, _abf, _ecd, _dfe int
		_ece, _egf, _ggcg, _geed                         uint32
		_dbgb, _fed                                      byte
		_gae                                             uint16
	)
	_affc := make([]byte, 4)
	_caff := make([]byte, 4)
	for _dab = 0; _dab < _cef.Height-1; _dab, _gba = _dab+2, _gba+1 {
		_gffdd = _dab * _cef.BytesPerLine
		_cge = _gba * _gee.BytesPerLine
		for _gec, _abf = 0, 0; _gec < _fbc; _gec, _abf = _gec+4, _abf+1 {
			for _ecd = 0; _ecd < 4; _ecd++ {
				_dfe = _gffdd + _gec + _ecd
				if _dfe <= len(_cef.Data)-1 && _dfe < _gffdd+_cef.BytesPerLine {
					_affc[_ecd] = _cef.Data[_dfe]
				} else {
					_affc[_ecd] = 0x00
				}
				_dfe = _gffdd + _cef.BytesPerLine + _gec + _ecd
				if _dfe <= len(_cef.Data)-1 && _dfe < _gffdd+(2*_cef.BytesPerLine) {
					_caff[_ecd] = _cef.Data[_dfe]
				} else {
					_caff[_ecd] = 0x00
				}
			}
			_ece = _bag.BigEndian.Uint32(_affc)
			_egf = _bag.BigEndian.Uint32(_caff)
			_ggcg = _ece & _egf
			_ggcg |= _ggcg << 1
			_geed = _ece | _egf
			_geed &= _geed << 1
			_egf = _ggcg & _geed
			_egf &= 0xaaaaaaaa
			_ece = _egf | (_egf << 7)
			_dbgb = byte(_ece >> 24)
			_fed = byte((_ece >> 8) & 0xff)
			_dfe = _cge + _abf
			if _dfe+1 == len(_gee.Data)-1 || _dfe+1 >= _cge+_gee.BytesPerLine {
				if _fec = _gee.setByte(_dfe, _fbgd[_dbgb]); _fec != nil {
					return _d.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _dfe)
				}
			} else {
				_gae = (uint16(_fbgd[_dbgb]) << 8) | uint16(_fbgd[_fed])
				if _fec = _gee.setTwoBytes(_dfe, _gae); _fec != nil {
					return _d.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _dfe)
				}
				_abf++
			}
		}
	}
	return nil
}
func (_fdbf *Gray2) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtGray2BPC(x, y, _fdbf.BytesPerLine, _fdbf.Data, _fdbf.Decode)
}
func ColorAtGray8BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_bd.Gray, error) {
	_afad := y*bytesPerLine + x
	if _afad >= len(data) {
		return _bd.Gray{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_ddca := data[_afad]
	if len(decode) == 2 {
		_ddca = uint8(uint32(LinearInterpolate(float64(_ddca), 0, 255, decode[0], decode[1])) & 0xff)
	}
	return _bd.Gray{Y: _ddca}, nil
}
func (_gbdfd *NRGBA16) SetNRGBA(x, y int, c _bd.NRGBA) {
	_feeea := y*_gbdfd.BytesPerLine + x*3/2
	if _feeea+1 >= len(_gbdfd.Data) {
		return
	}
	c = _eefb(c)
	_gbdfd.setNRGBA(x, y, _feeea, c)
}
func (_ggb *Gray16) Copy() Image { return &Gray16{ImageBase: _ggb.copy()} }
func (_ffbg *NRGBA16) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _ffbg.Width, Y: _ffbg.Height}}
}

var _ _a.Image = &NRGBA32{}

func (_bdee *RGBA32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _bdee.Width, Y: _bdee.Height}}
}
func _bgdf(_dfg Gray, _beeg nrgba64, _def _a.Rectangle) {
	for _gfdc := 0; _gfdc < _def.Max.X; _gfdc++ {
		for _gabc := 0; _gabc < _def.Max.Y; _gabc++ {
			_aec := _dbce(_beeg.NRGBA64At(_gfdc, _gabc))
			_dfg.SetGray(_gfdc, _gabc, _aec)
		}
	}
}
func (_gefe *NRGBA32) Set(x, y int, c _bd.Color) {
	_fccfg := y*_gefe.Width + x
	_efdb := 3 * _fccfg
	if _efdb+2 >= len(_gefe.Data) {
		return
	}
	_fbcd := _bd.NRGBAModel.Convert(c).(_bd.NRGBA)
	_gefe.setRGBA(_fccfg, _fbcd)
}
func (_de *CMYK32) Validate() error {
	if len(_de.Data) != 4*_de.Width*_de.Height {
		return _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_acfa *RGBA32) Base() *ImageBase { return &_acfa.ImageBase }
func (_fbfd *NRGBA64) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _fbfd.Width, Y: _fbfd.Height}}
}
func (_acac *ImageBase) setFourBytes(_geec int, _befd uint32) error {
	if _geec+3 > len(_acac.Data)-1 {
		return _d.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _geec)
	}
	_acac.Data[_geec] = byte((_befd & 0xff000000) >> 24)
	_acac.Data[_geec+1] = byte((_befd & 0xff0000) >> 16)
	_acac.Data[_geec+2] = byte((_befd & 0xff00) >> 8)
	_acac.Data[_geec+3] = byte(_befd & 0xff)
	return nil
}

const (
	PixSrc             RasterOperator = 0xc
	PixDst             RasterOperator = 0xa
	PixNotSrc          RasterOperator = 0x3
	PixNotDst          RasterOperator = 0x5
	PixClr             RasterOperator = 0x0
	PixSet             RasterOperator = 0xf
	PixSrcOrDst        RasterOperator = 0xe
	PixSrcAndDst       RasterOperator = 0x8
	PixSrcXorDst       RasterOperator = 0x6
	PixNotSrcOrDst     RasterOperator = 0xb
	PixNotSrcAndDst    RasterOperator = 0x2
	PixSrcOrNotDst     RasterOperator = 0xd
	PixSrcAndNotDst    RasterOperator = 0x4
	PixNotPixSrcOrDst  RasterOperator = 0x1
	PixNotPixSrcAndDst RasterOperator = 0x7
	PixNotPixSrcXorDst RasterOperator = 0x9
	PixPaint                          = PixSrcOrDst
	PixSubtract                       = PixNotSrcAndDst
	PixMask                           = PixSrcAndDst
)

var _ Image = &Gray8{}
var _ebbd [256]uint8

func _effc(_bdfeb RGBA, _fcdc NRGBA, _cacgc _a.Rectangle) {
	for _dadc := 0; _dadc < _cacgc.Max.X; _dadc++ {
		for _dbgcf := 0; _dbgcf < _cacgc.Max.Y; _dbgcf++ {
			_fbda := _bdfeb.RGBAAt(_dadc, _dbgcf)
			_fcdc.SetNRGBA(_dadc, _dbgcf, _agce(_fbda))
		}
	}
}

type CMYK interface {
	CMYKAt(_dgf, _cff int) _bd.CMYK
	SetCMYK(_bdbb, _bgec int, _dbbb _bd.CMYK)
}

func _eebf() {
	for _fffe := 0; _fffe < 256; _fffe++ {
		_ebbd[_fffe] = uint8(_fffe&0x1) + (uint8(_fffe>>1) & 0x1) + (uint8(_fffe>>2) & 0x1) + (uint8(_fffe>>3) & 0x1) + (uint8(_fffe>>4) & 0x1) + (uint8(_fffe>>5) & 0x1) + (uint8(_fffe>>6) & 0x1) + (uint8(_fffe>>7) & 0x1)
	}
}
func (_fbac *Gray2) Set(x, y int, c _bd.Color) {
	if x >= _fbac.Width || y >= _fbac.Height {
		return
	}
	_eacg := Gray2Model.Convert(c).(_bd.Gray)
	_dbaf := y * _fbac.BytesPerLine
	_bbag := _dbaf + (x >> 2)
	_cfeb := _eacg.Y >> 6
	_fbac.Data[_bbag] = (_fbac.Data[_bbag] & (^(0xc0 >> uint(2*((x)&3))))) | (_cfeb << uint(6-2*(x&3)))
}

var _ Image = &RGBA32{}

func (_eacf *Monochrome) ScaleLow(width, height int) (*Monochrome, error) {
	if width < 0 || height < 0 {
		return nil, _b.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064t\u0068\u0020\u0061\u006e\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_bgea := _efb(width, height)
	_cae := make([]int, height)
	_cfbg := make([]int, width)
	_fda := float64(_eacf.Width) / float64(width)
	_abbf := float64(_eacf.Height) / float64(height)
	for _ddcg := 0; _ddcg < height; _ddcg++ {
		_cae[_ddcg] = int(_ba.Min(_abbf*float64(_ddcg)+0.5, float64(_eacf.Height-1)))
	}
	for _ceec := 0; _ceec < width; _ceec++ {
		_cfbg[_ceec] = int(_ba.Min(_fda*float64(_ceec)+0.5, float64(_eacf.Width-1)))
	}
	_gdb := -1
	_agbdd := byte(0)
	for _eafg := 0; _eafg < height; _eafg++ {
		_edc := _cae[_eafg] * _eacf.BytesPerLine
		_adee := _eafg * _bgea.BytesPerLine
		for _cfff := 0; _cfff < width; _cfff++ {
			_bcac := _cfbg[_cfff]
			if _bcac != _gdb {
				_agbdd = _eacf.getBit(_edc, _bcac)
				if _agbdd != 0 {
					_bgea.setBit(_adee, _cfff)
				}
				_gdb = _bcac
			} else {
				if _agbdd != 0 {
					_bgea.setBit(_adee, _cfff)
				}
			}
		}
	}
	return _bgea, nil
}
func (_dded *CMYK32) Copy() Image { return &CMYK32{ImageBase: _dded.copy()} }
func _gc(_gf *Monochrome, _c int) (*Monochrome, error) {
	if _gf == nil {
		return nil, _b.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _c == 1 {
		return _gf.copy(), nil
	}
	if !IsPowerOf2(uint(_c)) {
		return nil, _d.Errorf("\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006ci\u0064 \u0065x\u0070a\u006e\u0064\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _c)
	}
	_abe := _ffga(_c)
	return _aa(_gf, _c, _abe)
}
func (_dfcf *Gray2) ColorModel() _bd.Model { return Gray2Model }
func _dfgb(_cfgdc int, _edcg int) error {
	return _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", _cfgdc, _edcg)
}
func _bbc(_gef *Monochrome, _aae ...int) (_dg *Monochrome, _bde error) {
	if _gef == nil {
		return nil, _b.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_aae) == 0 {
		return nil, _b.New("\u0074h\u0065\u0072e\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006f\u0066 \u0072\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e")
	}
	_dag := _bdc()
	_dg = _gef
	for _, _eba := range _aae {
		if _eba <= 0 {
			break
		}
		_dg, _bde = _geb(_dg, _eba, _dag)
		if _bde != nil {
			return nil, _bde
		}
	}
	return _dg, nil
}
func (_aab monochromeModel) Convert(c _bd.Color) _bd.Color {
	_eeff := _bd.GrayModel.Convert(c).(_bd.Gray)
	return _cgb(_eeff, _aab)
}

type Gray2 struct{ ImageBase }

func (_daa *ImageBase) setTwoBytes(_dafa int, _bafa uint16) error {
	if _dafa+1 > len(_daa.Data)-1 {
		return _b.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_daa.Data[_dafa] = byte((_bafa & 0xff00) >> 8)
	_daa.Data[_dafa+1] = byte(_bafa & 0xff)
	return nil
}
func _abda(_eaaa _bd.Gray) _bd.Gray {
	_fdd := _eaaa.Y >> 6
	_fdd |= _fdd << 2
	_eaaa.Y = _fdd | _fdd<<4
	return _eaaa
}
func _ce(_cf, _aef *Monochrome) (_fb error) {
	_efd := _aef.BytesPerLine
	_gab := _cf.BytesPerLine
	var _cc, _afb, _bab, _gfd, _ffd int
	for _bab = 0; _bab < _aef.Height; _bab++ {
		_cc = _bab * _efd
		_afb = 8 * _bab * _gab
		for _gfd = 0; _gfd < _efd; _gfd++ {
			if _fb = _cf.setEightBytes(_afb+_gfd*8, _bbd[_aef.Data[_cc+_gfd]]); _fb != nil {
				return _fb
			}
		}
		for _ffd = 1; _ffd < 8; _ffd++ {
			for _gfd = 0; _gfd < _gab; _gfd++ {
				if _fb = _cf.setByte(_afb+_ffd*_gab+_gfd, _cf.Data[_afb+_gfd]); _fb != nil {
					return _fb
				}
			}
		}
	}
	return nil
}
func _fgec(_dfc _bd.CMYK) _bd.NRGBA {
	_fbb, _cgcg, _gfcc := _bd.CMYKToRGB(_dfc.C, _dfc.M, _dfc.Y, _dfc.K)
	return _bd.NRGBA{R: _fbb, G: _cgcg, B: _gfcc, A: 0xff}
}
func (_efeg *Monochrome) ExpandBinary(factor int) (*Monochrome, error) {
	if !IsPowerOf2(uint(factor)) {
		return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0065\u0078\u0070\u0061\u006e\u0064\u0020b\u0069n\u0061\u0072\u0079\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", factor)
	}
	return _gc(_efeg, factor)
}
func _bfb(_agdb *Monochrome, _ebba, _fgae, _ebce, _afaf int, _eeffe RasterOperator, _gagb *Monochrome, _dfeg, _bcec int) error {
	var (
		_dcaf        byte
		_dee         int
		_cbgd        int
		_dbf, _bfde  int
		_eec, _agcdc int
	)
	_accge := _ebce >> 3
	_agcc := _ebce & 7
	if _agcc > 0 {
		_dcaf = _ddfe[_agcc]
	}
	_dee = _gagb.BytesPerLine*_bcec + (_dfeg >> 3)
	_cbgd = _agdb.BytesPerLine*_fgae + (_ebba >> 3)
	switch _eeffe {
	case PixSrc:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = _gagb.Data[_dbf]
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf], _dcaf)
			}
		}
	case PixNotSrc:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = ^(_gagb.Data[_dbf])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^_gagb.Data[_dbf], _dcaf)
			}
		}
	case PixSrcOrDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] |= _gagb.Data[_dbf]
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf]|_agdb.Data[_bfde], _dcaf)
			}
		}
	case PixSrcAndDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] &= _gagb.Data[_dbf]
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf]&_agdb.Data[_bfde], _dcaf)
			}
		}
	case PixSrcXorDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] ^= _gagb.Data[_dbf]
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf]^_agdb.Data[_bfde], _dcaf)
			}
		}
	case PixNotSrcOrDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] |= ^(_gagb.Data[_dbf])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^(_gagb.Data[_dbf])|_agdb.Data[_bfde], _dcaf)
			}
		}
	case PixNotSrcAndDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] &= ^(_gagb.Data[_dbf])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^(_gagb.Data[_dbf])&_agdb.Data[_bfde], _dcaf)
			}
		}
	case PixSrcOrNotDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = _gagb.Data[_dbf] | ^(_agdb.Data[_bfde])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf]|^(_agdb.Data[_bfde]), _dcaf)
			}
		}
	case PixSrcAndNotDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = _gagb.Data[_dbf] &^ (_agdb.Data[_bfde])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], _gagb.Data[_dbf]&^(_agdb.Data[_bfde]), _dcaf)
			}
		}
	case PixNotPixSrcOrDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = ^(_gagb.Data[_dbf] | _agdb.Data[_bfde])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^(_gagb.Data[_dbf] | _agdb.Data[_bfde]), _dcaf)
			}
		}
	case PixNotPixSrcAndDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = ^(_gagb.Data[_dbf] & _agdb.Data[_bfde])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^(_gagb.Data[_dbf] & _agdb.Data[_bfde]), _dcaf)
			}
		}
	case PixNotPixSrcXorDst:
		for _eec = 0; _eec < _afaf; _eec++ {
			_dbf = _dee + _eec*_gagb.BytesPerLine
			_bfde = _cbgd + _eec*_agdb.BytesPerLine
			for _agcdc = 0; _agcdc < _accge; _agcdc++ {
				_agdb.Data[_bfde] = ^(_gagb.Data[_dbf] ^ _agdb.Data[_bfde])
				_bfde++
				_dbf++
			}
			if _agcc > 0 {
				_agdb.Data[_bfde] = _agcee(_agdb.Data[_bfde], ^(_gagb.Data[_dbf] ^ _agdb.Data[_bfde]), _dcaf)
			}
		}
	default:
		_e.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _eeffe)
		return _b.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func _aebbc(_ceea nrgba64, _gffe NRGBA, _eegeg _a.Rectangle) {
	for _gdgb := 0; _gdgb < _eegeg.Max.X; _gdgb++ {
		for _cedg := 0; _cedg < _eegeg.Max.Y; _cedg++ {
			_gfbf := _ceea.NRGBA64At(_gdgb, _cedg)
			_gffe.SetNRGBA(_gdgb, _cedg, _gbgb(_gfbf))
		}
	}
}
func _bebb(_fgcb Gray, _dedd RGBA, _cece _a.Rectangle) {
	for _bbaad := 0; _bbaad < _cece.Max.X; _bbaad++ {
		for _gebee := 0; _gebee < _cece.Max.Y; _gebee++ {
			_bcfc := _fgcb.GrayAt(_bbaad, _gebee)
			_dedd.SetRGBA(_bbaad, _gebee, _age(_bcfc))
		}
	}
}

var ErrInvalidImage = _b.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")

type RasterOperator int

func _dbgca(_abbgd NRGBA, _cbge RGBA, _gbcd _a.Rectangle) {
	for _fgda := 0; _fgda < _gbcd.Max.X; _fgda++ {
		for _gcdc := 0; _gcdc < _gbcd.Max.Y; _gcdc++ {
			_cbbd := _abbgd.NRGBAAt(_fgda, _gcdc)
			_cbge.SetRGBA(_fgda, _gcdc, _ggf(_cbbd))
		}
	}
}
func IsPowerOf2(n uint) bool     { return n > 0 && (n&(n-1)) == 0 }
func (_dacg *Gray8) Copy() Image { return &Gray8{ImageBase: _dacg.copy()} }
func _abee(_bdfa _a.Image, _babe Image, _baaa _a.Rectangle) {
	if _dbdae, _afed := _bdfa.(SMasker); _afed && _dbdae.HasAlpha() {
		_babe.(SMasker).MakeAlpha()
	}
	_ebbc(_bdfa, _babe, _baaa)
}
func ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor int, data, alpha []byte, decode []float64) (_bd.Color, error) {
	switch bitsPerColor {
	case 4:
		return ColorAtNRGBA16(x, y, width, bytesPerLine, data, alpha, decode)
	case 8:
		return ColorAtNRGBA32(x, y, width, data, alpha, decode)
	case 16:
		return ColorAtNRGBA64(x, y, width, data, alpha, decode)
	default:
		return nil, _d.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0067\u0062\u0020b\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0061\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}
func (_gbdc *ImageBase) setEightPartlyBytes(_abbfb, _ggfd int, _dcace uint64) (_cbbg error) {
	var (
		_cegg  byte
		_gggag int
	)
	for _ccab := 1; _ccab <= _ggfd; _ccab++ {
		_gggag = 64 - _ccab*8
		_cegg = byte(_dcace >> uint(_gggag) & 0xff)
		if _cbbg = _gbdc.setByte(_abbfb+_ccab-1, _cegg); _cbbg != nil {
			return _cbbg
		}
	}
	_aeba := _gbdc.BytesPerLine*8 - _gbdc.Width
	if _aeba == 0 {
		return nil
	}
	_gggag -= 8
	_cegg = byte(_dcace>>uint(_gggag)&0xff) << uint(_aeba)
	if _cbbg = _gbdc.setByte(_abbfb+_ggfd, _cegg); _cbbg != nil {
		return _cbbg
	}
	return nil
}
func _cbgc(_fece, _dgfc NRGBA, _fdagd _a.Rectangle) {
	for _bbge := 0; _bbge < _fdagd.Max.X; _bbge++ {
		for _aecd := 0; _aecd < _fdagd.Max.Y; _aecd++ {
			_dgfc.SetNRGBA(_bbge, _aecd, _fece.NRGBAAt(_bbge, _aecd))
		}
	}
}

type shift int

func (_bed *Gray4) ColorAt(x, y int) (_bd.Color, error) {
	return ColorAtGray4BPC(x, y, _bed.BytesPerLine, _bed.Data, _bed.Decode)
}
func (_ebfe *Gray8) GrayAt(x, y int) _bd.Gray {
	_afce, _ := ColorAtGray8BPC(x, y, _ebfe.BytesPerLine, _ebfe.Data, _ebfe.Decode)
	return _afce
}
func _adbe(_cfbc uint8) bool {
	if _cfbc == 0 || _cfbc == 255 {
		return true
	}
	return false
}
func (_edcd *Gray4) Base() *ImageBase { return &_edcd.ImageBase }
func (_gabf *Gray8) SetGray(x, y int, g _bd.Gray) {
	_bgdd := y*_gabf.BytesPerLine + x
	if _bgdd > len(_gabf.Data)-1 {
		return
	}
	_gabf.Data[_bgdd] = g.Y
}
func ColorAtNRGBA64(x, y, width int, data, alpha []byte, decode []float64) (_bd.NRGBA64, error) {
	_dcgg := (y*width + x) * 2
	_adad := _dcgg * 3
	if _adad+5 >= len(data) {
		return _bd.NRGBA64{}, _d.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	const _dadbc = 0xffff
	_afcg := uint16(_dadbc)
	if alpha != nil && len(alpha) > _dcgg+1 {
		_afcg = uint16(alpha[_dcgg])<<8 | uint16(alpha[_dcgg+1])
	}
	_abfeg := uint16(data[_adad])<<8 | uint16(data[_adad+1])
	_cfdg := uint16(data[_adad+2])<<8 | uint16(data[_adad+3])
	_adgb := uint16(data[_adad+4])<<8 | uint16(data[_adad+5])
	if len(decode) == 6 {
		_abfeg = uint16(uint64(LinearInterpolate(float64(_abfeg), 0, 65535, decode[0], decode[1])) & _dadbc)
		_cfdg = uint16(uint64(LinearInterpolate(float64(_cfdg), 0, 65535, decode[2], decode[3])) & _dadbc)
		_adgb = uint16(uint64(LinearInterpolate(float64(_adgb), 0, 65535, decode[4], decode[5])) & _dadbc)
	}
	return _bd.NRGBA64{R: _abfeg, G: _cfdg, B: _adgb, A: _afcg}, nil
}
func (_aacf *ImageBase) HasAlpha() bool {
	if _aacf.Alpha == nil {
		return false
	}
	for _aeeeaf := range _aacf.Alpha {
		if _aacf.Alpha[_aeeeaf] != 0xff {
			return true
		}
	}
	return false
}
func (_cgcc *Monochrome) getBit(_ecb, _cgg int) uint8 {
	return _cgcc.Data[_ecb+(_cgg>>3)] >> uint(7-(_cgg&7)) & 1
}
func _feeg(_aga _bd.Color) _bd.Color {
	_fdeff := _bd.GrayModel.Convert(_aga).(_bd.Gray)
	return _abda(_fdeff)
}
func _ec(_ebf, _fe *Monochrome) (_dbe error) {
	_da := _fe.BytesPerLine
	_gaf := _ebf.BytesPerLine
	_dc := _fe.BytesPerLine*4 - _ebf.BytesPerLine
	var (
		_gd, _gafb                            byte
		_ffg                                  uint32
		_dce, _dbg, _ac, _be, _gff, _fc, _gdd int
	)
	for _ac = 0; _ac < _fe.Height; _ac++ {
		_dce = _ac * _da
		_dbg = 4 * _ac * _gaf
		for _be = 0; _be < _da; _be++ {
			_gd = _fe.Data[_dce+_be]
			_ffg = _ffb[_gd]
			_fc = _dbg + _be*4
			if _dc != 0 && (_be+1)*4 > _ebf.BytesPerLine {
				for _gff = _dc; _gff > 0; _gff-- {
					_gafb = byte((_ffg >> uint(_gff*8)) & 0xff)
					_gdd = _fc + (_dc - _gff)
					if _dbe = _ebf.setByte(_gdd, _gafb); _dbe != nil {
						return _dbe
					}
				}
			} else if _dbe = _ebf.setFourBytes(_fc, _ffg); _dbe != nil {
				return _dbe
			}
			if _dbe = _ebf.setFourBytes(_dbg+_be*4, _ffb[_fe.Data[_dce+_be]]); _dbe != nil {
				return _dbe
			}
		}
		for _gff = 1; _gff < 4; _gff++ {
			for _be = 0; _be < _gaf; _be++ {
				if _dbe = _ebf.setByte(_dbg+_gff*_gaf+_be, _ebf.Data[_dbg+_be]); _dbe != nil {
					return _dbe
				}
			}
		}
	}
	return nil
}

var (
	_cga = _fcd()
	_ffb = _fg()
	_bbd = _cddb()
)
var _ Image = &Gray2{}
