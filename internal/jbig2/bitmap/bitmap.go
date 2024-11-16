package bitmap

import (
	_df "encoding/binary"
	_fb "image"
	_c "math"
	_a "sort"
	_e "strings"
	_da "testing"

	_ea "github.com/bamzi/pdfext/common"
	_bg "github.com/bamzi/pdfext/internal/bitwise"
	_ed "github.com/bamzi/pdfext/internal/imageutil"
	_b "github.com/bamzi/pdfext/internal/jbig2/basic"
	_d "github.com/bamzi/pdfext/internal/jbig2/errors"
	_dc "github.com/stretchr/testify/require"
)

func (_gged *BitmapsArray) AddBox(box *_fb.Rectangle) { _gged.Boxes = append(_gged.Boxes, box) }
func TstGetScaledSymbol(t *_da.T, sm *Bitmap, scale ...int) *Bitmap {
	if len(scale) == 0 {
		return sm
	}
	if scale[0] == 1 {
		return sm
	}
	_ddgc, _ecagd := MorphSequence(sm, MorphProcess{Operation: MopReplicativeBinaryExpansion, Arguments: scale})
	_dc.NoError(t, _ecagd)
	return _ddgc
}
func _gcc(_fgdg uint, _bffc byte) byte { return _bffc >> _fgdg << _fgdg }
func (_ddga *Bitmap) equivalent(_cdb *Bitmap) bool {
	if _ddga == _cdb {
		return true
	}
	if !_ddga.SizesEqual(_cdb) {
		return false
	}
	_gcef := _ecfg(_ddga, _cdb, CmbOpXor)
	_dggfc := _ddga.countPixels()
	_dafba := int(0.25 * float32(_dggfc))
	if _gcef.thresholdPixelSum(_dafba) {
		return false
	}
	var (
		_fda  [9][9]int
		_afae [18][9]int
		_dadd [9][18]int
		_gdbc int
		_eeag int
	)
	_edc := 9
	_dedf := _ddga.Height / _edc
	_cfe := _ddga.Width / _edc
	_gde, _cfbe := _dedf/2, _cfe/2
	if _dedf < _cfe {
		_gde = _cfe / 2
		_cfbe = _dedf / 2
	}
	_bcb := float64(_gde) * float64(_cfbe) * _c.Pi
	_ebdd := int(float64(_dedf*_cfe/2) * 0.9)
	_dbfa := int(float64(_cfe*_dedf/2) * 0.9)
	for _dfge := 0; _dfge < _edc; _dfge++ {
		_deda := _cfe*_dfge + _gdbc
		var _cba int
		if _dfge == _edc-1 {
			_gdbc = 0
			_cba = _ddga.Width
		} else {
			_cba = _deda + _cfe
			if ((_ddga.Width - _gdbc) % _edc) > 0 {
				_gdbc++
				_cba++
			}
		}
		for _eege := 0; _eege < _edc; _eege++ {
			_gagf := _dedf*_eege + _eeag
			var _eccb int
			if _eege == _edc-1 {
				_eeag = 0
				_eccb = _ddga.Height
			} else {
				_eccb = _gagf + _dedf
				if (_ddga.Height-_eeag)%_edc > 0 {
					_eeag++
					_eccb++
				}
			}
			var _dbbd, _bbb, _gceca, _abc int
			_cbg := (_deda + _cba) / 2
			_ggd := (_gagf + _eccb) / 2
			for _ccef := _deda; _ccef < _cba; _ccef++ {
				for _cgga := _gagf; _cgga < _eccb; _cgga++ {
					if _gcef.GetPixel(_ccef, _cgga) {
						if _ccef < _cbg {
							_dbbd++
						} else {
							_bbb++
						}
						if _cgga < _ggd {
							_abc++
						} else {
							_gceca++
						}
					}
				}
			}
			_fda[_dfge][_eege] = _dbbd + _bbb
			_afae[_dfge*2][_eege] = _dbbd
			_afae[_dfge*2+1][_eege] = _bbb
			_dadd[_dfge][_eege*2] = _abc
			_dadd[_dfge][_eege*2+1] = _gceca
		}
	}
	for _ccc := 0; _ccc < _edc*2-1; _ccc++ {
		for _faab := 0; _faab < (_edc - 1); _faab++ {
			var _cfabg int
			for _deag := 0; _deag < 2; _deag++ {
				for _dgf := 0; _dgf < 2; _dgf++ {
					_cfabg += _afae[_ccc+_deag][_faab+_dgf]
				}
			}
			if _cfabg > _dbfa {
				return false
			}
		}
	}
	for _egg := 0; _egg < (_edc - 1); _egg++ {
		for _gbgb := 0; _gbgb < ((_edc * 2) - 1); _gbgb++ {
			var _gafa int
			for _bcdfd := 0; _bcdfd < 2; _bcdfd++ {
				for _cfgge := 0; _cfgge < 2; _cfgge++ {
					_gafa += _dadd[_egg+_bcdfd][_gbgb+_cfgge]
				}
			}
			if _gafa > _ebdd {
				return false
			}
		}
	}
	for _affg := 0; _affg < (_edc - 2); _affg++ {
		for _badg := 0; _badg < (_edc - 2); _badg++ {
			var _dbbb, _fbgd int
			for _dbab := 0; _dbab < 3; _dbab++ {
				for _bgda := 0; _bgda < 3; _bgda++ {
					if _dbab == _bgda {
						_dbbb += _fda[_affg+_dbab][_badg+_bgda]
					}
					if (2 - _dbab) == _bgda {
						_fbgd += _fda[_affg+_dbab][_badg+_bgda]
					}
				}
			}
			if _dbbb > _dbfa || _fbgd > _dbfa {
				return false
			}
		}
	}
	for _cbfa := 0; _cbfa < (_edc - 1); _cbfa++ {
		for _aafd := 0; _aafd < (_edc - 1); _aafd++ {
			var _aacd int
			for _cbaa := 0; _cbaa < 2; _cbaa++ {
				for _feeg := 0; _feeg < 2; _feeg++ {
					_aacd += _fda[_cbfa+_cbaa][_aafd+_feeg]
				}
			}
			if float64(_aacd) > _bcb {
				return false
			}
		}
	}
	return true
}
func (_bddf *Bitmap) setBit(_dfba int) { _bddf.Data[(_dfba >> 3)] |= 0x80 >> uint(_dfba&7) }

type LocationFilter int

func (_cadc *byHeight) Less(i, j int) bool       { return _cadc.Values[i].Height < _cadc.Values[j].Height }
func (_defd *Bitmaps) AddBox(box *_fb.Rectangle) { _defd.Boxes = append(_defd.Boxes, box) }
func _ebgg(_adde *Bitmap, _babc, _cedb, _bcg, _cefb int, _edaa RasterOperator) {
	if _babc < 0 {
		_bcg += _babc
		_babc = 0
	}
	_cdfd := _babc + _bcg - _adde.Width
	if _cdfd > 0 {
		_bcg -= _cdfd
	}
	if _cedb < 0 {
		_cefb += _cedb
		_cedb = 0
	}
	_gcgfe := _cedb + _cefb - _adde.Height
	if _gcgfe > 0 {
		_cefb -= _gcgfe
	}
	if _bcg <= 0 || _cefb <= 0 {
		return
	}
	if (_babc & 7) == 0 {
		_bgab(_adde, _babc, _cedb, _bcg, _cefb, _edaa)
	} else {
		_efgb(_adde, _babc, _cedb, _bcg, _cefb, _edaa)
	}
}
func _cgg() (_ba [256]uint16) {
	for _gag := 0; _gag < 256; _gag++ {
		if _gag&0x01 != 0 {
			_ba[_gag] |= 0x3
		}
		if _gag&0x02 != 0 {
			_ba[_gag] |= 0xc
		}
		if _gag&0x04 != 0 {
			_ba[_gag] |= 0x30
		}
		if _gag&0x08 != 0 {
			_ba[_gag] |= 0xc0
		}
		if _gag&0x10 != 0 {
			_ba[_gag] |= 0x300
		}
		if _gag&0x20 != 0 {
			_ba[_gag] |= 0xc00
		}
		if _gag&0x40 != 0 {
			_ba[_gag] |= 0x3000
		}
		if _gag&0x80 != 0 {
			_ba[_gag] |= 0xc000
		}
	}
	return _ba
}
func (_cgff *Bitmap) setPadBits(_edg int) {
	_fcd := 8 - _cgff.Width%8
	if _fcd == 8 {
		return
	}
	_fceg := _cgff.Width / 8
	_bdgc := _eaec[_fcd]
	if _edg == 0 {
		_bdgc ^= _bdgc
	}
	var _bdbd int
	for _egf := 0; _egf < _cgff.Height; _egf++ {
		_bdbd = _egf*_cgff.RowStride + _fceg
		if _edg == 0 {
			_cgff.Data[_bdbd] &= _bdgc
		} else {
			_cgff.Data[_bdbd] |= _bdgc
		}
	}
}

type Color int

func _egfg(_bba int) int {
	if _bba < 0 {
		return -_bba
	}
	return _bba
}
func _cab(_bgeg, _aabd, _gdd *Bitmap) (*Bitmap, error) {
	const _egbf = "\u0073\u0075\u0062\u0074\u0072\u0061\u0063\u0074"
	if _aabd == nil {
		return nil, _d.Error(_egbf, "'\u0073\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _gdd == nil {
		return nil, _d.Error(_egbf, "'\u0073\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	var _cabd error
	switch {
	case _bgeg == _aabd:
		if _cabd = _bgeg.RasterOperation(0, 0, _aabd.Width, _aabd.Height, PixNotSrcAndDst, _gdd, 0, 0); _cabd != nil {
			return nil, _d.Wrap(_cabd, _egbf, "\u0064 \u003d\u003d\u0020\u0073\u0031")
		}
	case _bgeg == _gdd:
		if _cabd = _bgeg.RasterOperation(0, 0, _aabd.Width, _aabd.Height, PixNotSrcAndDst, _aabd, 0, 0); _cabd != nil {
			return nil, _d.Wrap(_cabd, _egbf, "\u0064 \u003d\u003d\u0020\u0073\u0032")
		}
	default:
		_bgeg, _cabd = _bce(_bgeg, _aabd)
		if _cabd != nil {
			return nil, _d.Wrap(_cabd, _egbf, "")
		}
		if _cabd = _bgeg.RasterOperation(0, 0, _aabd.Width, _aabd.Height, PixNotSrcAndDst, _gdd, 0, 0); _cabd != nil {
			return nil, _d.Wrap(_cabd, _egbf, "\u0064e\u0066\u0061\u0075\u006c\u0074")
		}
	}
	return _bgeg, nil
}

const (
	_ SizeComparison = iota
	SizeSelectIfLT
	SizeSelectIfGT
	SizeSelectIfLTE
	SizeSelectIfGTE
	SizeSelectIfEQ
)

func (_fbd *Bitmap) SetPadBits(value int) { _fbd.setPadBits(value) }
func _bce(_gcbd, _gagb *Bitmap) (*Bitmap, error) {
	if _gagb == nil {
		return nil, _d.Error("\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gagb == _gcbd {
		return _gcbd, nil
	}
	if _gcbd == nil {
		_gcbd = _gagb.createTemplate()
		copy(_gcbd.Data, _gagb.Data)
		return _gcbd, nil
	}
	_gda := _gcbd.resizeImageData(_gagb)
	if _gda != nil {
		return nil, _d.Wrap(_gda, "\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "")
	}
	_gcbd.Text = _gagb.Text
	copy(_gcbd.Data, _gagb.Data)
	return _gcbd, nil
}
func init() {
	for _cbeg := 0; _cbeg < 256; _cbeg++ {
		_cdc[_cbeg] = uint8(_cbeg&0x1) + (uint8(_cbeg>>1) & 0x1) + (uint8(_cbeg>>2) & 0x1) + (uint8(_cbeg>>3) & 0x1) + (uint8(_cbeg>>4) & 0x1) + (uint8(_cbeg>>5) & 0x1) + (uint8(_cbeg>>6) & 0x1) + (uint8(_cbeg>>7) & 0x1)
	}
}
func (_dedg *Bitmap) InverseData() { _dedg.inverseData() }
func (_fgb *Bitmap) addPadBits() (_gaa error) {
	const _efec = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0061\u0064\u0064\u0050\u0061d\u0042\u0069\u0074\u0073"
	_cfgg := _fgb.Width % 8
	if _cfgg == 0 {
		return nil
	}
	_cgfda := _fgb.Width / 8
	_cbdc := _bg.NewReader(_fgb.Data)
	_dfed := make([]byte, _fgb.Height*_fgb.RowStride)
	_aed := _bg.NewWriterMSB(_dfed)
	_dedgb := make([]byte, _cgfda)
	var (
		_cga  int
		_gebg uint64
	)
	for _cga = 0; _cga < _fgb.Height; _cga++ {
		if _, _gaa = _cbdc.Read(_dedgb); _gaa != nil {
			return _d.Wrap(_gaa, _efec, "\u0066u\u006c\u006c\u0020\u0062\u0079\u0074e")
		}
		if _, _gaa = _aed.Write(_dedgb); _gaa != nil {
			return _d.Wrap(_gaa, _efec, "\u0066\u0075\u006c\u006c\u0020\u0062\u0079\u0074\u0065\u0073")
		}
		if _gebg, _gaa = _cbdc.ReadBits(byte(_cfgg)); _gaa != nil {
			return _d.Wrap(_gaa, _efec, "\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0069\u0074\u0073")
		}
		if _gaa = _aed.WriteByte(byte(_gebg) << uint(8-_cfgg)); _gaa != nil {
			return _d.Wrap(_gaa, _efec, "\u006ca\u0073\u0074\u0020\u0062\u0079\u0074e")
		}
	}
	_fgb.Data = _aed.Data()
	return nil
}
func _gaeee(_bggef *Bitmap, _bbcg *_b.Stack, _beed, _fdeb int) (_degd *_fb.Rectangle, _adgcb error) {
	const _dbgde = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _bggef == nil {
		return nil, _d.Error(_dbgde, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _bbcg == nil {
		return nil, _d.Error(_dbgde, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_fgab, _dacd := _bggef.Width, _bggef.Height
	_gafdd := _fgab - 1
	_ffef := _dacd - 1
	if _beed < 0 || _beed > _gafdd || _fdeb < 0 || _fdeb > _ffef || !_bggef.GetPixel(_beed, _fdeb) {
		return nil, nil
	}
	var _bcgb *_fb.Rectangle
	_bcgb, _adgcb = Rect(100000, 100000, 0, 0)
	if _adgcb != nil {
		return nil, _d.Wrap(_adgcb, _dbgde, "")
	}
	if _adgcb = _geab(_bbcg, _beed, _beed, _fdeb, 1, _ffef, _bcgb); _adgcb != nil {
		return nil, _d.Wrap(_adgcb, _dbgde, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _adgcb = _geab(_bbcg, _beed, _beed, _fdeb+1, -1, _ffef, _bcgb); _adgcb != nil {
		return nil, _d.Wrap(_adgcb, _dbgde, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_bcgb.Min.X, _bcgb.Max.X = _beed, _beed
	_bcgb.Min.Y, _bcgb.Max.Y = _fdeb, _fdeb
	var (
		_fabe *fillSegment
		_eced int
	)
	for _bbcg.Len() > 0 {
		if _fabe, _adgcb = _feef(_bbcg); _adgcb != nil {
			return nil, _d.Wrap(_adgcb, _dbgde, "")
		}
		_fdeb = _fabe._aegc
		for _beed = _fabe._gebc; _beed >= 0 && _bggef.GetPixel(_beed, _fdeb); _beed-- {
			if _adgcb = _bggef.SetPixel(_beed, _fdeb, 0); _adgcb != nil {
				return nil, _d.Wrap(_adgcb, _dbgde, "")
			}
		}
		if _beed >= _fabe._gebc {
			for _beed++; _beed <= _fabe._fdeg && _beed <= _gafdd && !_bggef.GetPixel(_beed, _fdeb); _beed++ {
			}
			_eced = _beed
			if !(_beed <= _fabe._fdeg && _beed <= _gafdd) {
				continue
			}
		} else {
			_eced = _beed + 1
			if _eced < _fabe._gebc-1 {
				if _adgcb = _geab(_bbcg, _eced, _fabe._gebc-1, _fabe._aegc, -_fabe._bbcgb, _ffef, _bcgb); _adgcb != nil {
					return nil, _d.Wrap(_adgcb, _dbgde, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
				}
			}
			_beed = _fabe._gebc + 1
		}
		for {
			for ; _beed <= _gafdd && _bggef.GetPixel(_beed, _fdeb); _beed++ {
				if _adgcb = _bggef.SetPixel(_beed, _fdeb, 0); _adgcb != nil {
					return nil, _d.Wrap(_adgcb, _dbgde, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _adgcb = _geab(_bbcg, _eced, _beed-1, _fabe._aegc, _fabe._bbcgb, _ffef, _bcgb); _adgcb != nil {
				return nil, _d.Wrap(_adgcb, _dbgde, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _beed > _fabe._fdeg+1 {
				if _adgcb = _geab(_bbcg, _fabe._fdeg+1, _beed-1, _fabe._aegc, -_fabe._bbcgb, _ffef, _bcgb); _adgcb != nil {
					return nil, _d.Wrap(_adgcb, _dbgde, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _beed++; _beed <= _fabe._fdeg && _beed <= _gafdd && !_bggef.GetPixel(_beed, _fdeb); _beed++ {
			}
			_eced = _beed
			if !(_beed <= _fabe._fdeg && _beed <= _gafdd) {
				break
			}
		}
	}
	_bcgb.Max.X++
	_bcgb.Max.Y++
	return _bcgb, nil
}
func (_dacb *Bitmap) setTwoBytes(_beg int, _ecac uint16) error {
	if _beg+1 > len(_dacb.Data)-1 {
		return _d.Errorf("s\u0065\u0074\u0054\u0077\u006f\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _beg)
	}
	_dacb.Data[_beg] = byte((_ecac & 0xff00) >> 8)
	_dacb.Data[_beg+1] = byte(_ecac & 0xff)
	return nil
}
func TstFrameBitmapData() []byte { return _bfgad.Data }
func _abbg(_ededg, _cecf *Bitmap, _fcff, _caf int) (_efbd error) {
	const _daca = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0038"
	var (
		_dgee, _fceea, _bgaee, _bdbdd                             int
		_eabf, _gagfc, _ecfc, _debe, _baff, _babce, _gddd, _cedcf byte
	)
	for _dgee = 0; _dgee < _fcff; _dgee++ {
		_bgaee = _dgee * _ededg.RowStride
		_bdbdd = _dgee * _cecf.RowStride
		for _fceea = 0; _fceea < _caf; _fceea++ {
			if _eabf, _efbd = _ededg.GetByte(_bgaee + _fceea); _efbd != nil {
				return _d.Wrap(_efbd, _daca, "\u0067e\u0074 \u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _gagfc, _efbd = _cecf.GetByte(_bdbdd + _fceea); _efbd != nil {
				return _d.Wrap(_efbd, _daca, "\u0067\u0065\u0074\u0020\u006d\u0061\u0073\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _dgee > 0 {
				if _ecfc, _efbd = _ededg.GetByte(_bgaee - _ededg.RowStride + _fceea); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u0069\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
				}
				_eabf |= _ecfc | (_ecfc << 1) | (_ecfc >> 1)
				if _fceea > 0 {
					if _cedcf, _efbd = _ededg.GetByte(_bgaee - _ededg.RowStride + _fceea - 1); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "\u0069\u0020\u003e\u00200 \u0026\u0026\u0020\u006a\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
					}
					_eabf |= _cedcf << 7
				}
				if _fceea < _caf-1 {
					if _cedcf, _efbd = _ededg.GetByte(_bgaee - _ededg.RowStride + _fceea + 1); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "\u006a\u0020<\u0020\u0077\u0070l\u0020\u002d\u0020\u0031\u0020\u0062\u0079\u0074\u0065")
					}
					_eabf |= _cedcf >> 7
				}
			}
			if _fceea > 0 {
				if _debe, _efbd = _ededg.GetByte(_bgaee + _fceea - 1); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u006a\u0020\u003e \u0030")
				}
				_eabf |= _debe << 7
			}
			_eabf &= _gagfc
			if _eabf == 0 || ^_eabf == 0 {
				if _efbd = _ededg.SetByte(_bgaee+_fceea, _eabf); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u0073e\u0074t\u0069\u006e\u0067\u0020\u0065m\u0070\u0074y\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_gddd = _eabf
				_eabf = (_eabf | (_eabf >> 1) | (_eabf << 1)) & _gagfc
				if (_eabf ^ _gddd) == 0 {
					if _efbd = _ededg.SetByte(_bgaee+_fceea, _eabf); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _dgee = _fcff - 1; _dgee >= 0; _dgee-- {
		_bgaee = _dgee * _ededg.RowStride
		_bdbdd = _dgee * _cecf.RowStride
		for _fceea = _caf - 1; _fceea >= 0; _fceea-- {
			if _eabf, _efbd = _ededg.GetByte(_bgaee + _fceea); _efbd != nil {
				return _d.Wrap(_efbd, _daca, "\u0072\u0065\u0076er\u0073\u0065\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _gagfc, _efbd = _cecf.GetByte(_bdbdd + _fceea); _efbd != nil {
				return _d.Wrap(_efbd, _daca, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _dgee < _fcff-1 {
				if _baff, _efbd = _ededg.GetByte(_bgaee + _ededg.RowStride + _fceea); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u0069\u0020\u003c\u0020h\u0020\u002d\u0020\u0031\u0020\u002d\u003e\u0020\u0067\u0065t\u0020s\u006f\u0075\u0072\u0063\u0065\u0020\u0062y\u0074\u0065")
				}
				_eabf |= _baff | (_baff << 1) | _baff>>1
				if _fceea > 0 {
					if _cedcf, _efbd = _ededg.GetByte(_bgaee + _ededg.RowStride + _fceea - 1); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "\u0069\u0020\u003c h\u002d\u0031\u0020\u0026\u0020\u006a\u0020\u003e\u00200\u0020-\u003e \u0067e\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
					}
					_eabf |= _cedcf << 7
				}
				if _fceea < _caf-1 {
					if _cedcf, _efbd = _ededg.GetByte(_bgaee + _ededg.RowStride + _fceea + 1); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "\u0069\u0020\u003c\u0020\u0068\u002d\u0031\u0020\u0026\u0026\u0020\u006a\u0020\u003c\u0077\u0070\u006c\u002d\u0031\u0020\u002d\u003e\u0020\u0067e\u0074\u0020\u0073\u006f\u0075r\u0063\u0065 \u0062\u0079\u0074\u0065")
					}
					_eabf |= _cedcf >> 7
				}
			}
			if _fceea < _caf-1 {
				if _babce, _efbd = _ededg.GetByte(_bgaee + _fceea + 1); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u006a\u0020<\u0020\u0077\u0070\u006c\u0020\u002d\u0031\u0020\u002d\u003e\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020by\u0074\u0065")
				}
				_eabf |= _babce >> 7
			}
			_eabf &= _gagfc
			if _eabf == 0 || (^_eabf) == 0 {
				if _efbd = _ededg.SetByte(_bgaee+_fceea, _eabf); _efbd != nil {
					return _d.Wrap(_efbd, _daca, "\u0073e\u0074 \u006d\u0061\u0073\u006b\u0065\u0064\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_gddd = _eabf
				_eabf = (_eabf | (_eabf >> 1) | (_eabf << 1)) & _gagfc
				if (_eabf ^ _gddd) == 0 {
					if _efbd = _ededg.SetByte(_bgaee+_fceea, _eabf); _efbd != nil {
						return _d.Wrap(_efbd, _daca, "r\u0065\u0076\u0065\u0072se\u0020s\u0065\u0074\u0020\u0070\u0072e\u0076\u0020\u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	return nil
}
func TstAddSymbol(t *_da.T, bms *Bitmaps, sym *Bitmap, x *int, y int, space int) {
	bms.AddBitmap(sym)
	_cagf := _fb.Rect(*x, y, *x+sym.Width, y+sym.Height)
	bms.AddBox(&_cagf)
	*x += sym.Width + space
}

type fillSegment struct {
	_gebc  int
	_fdeg  int
	_aegc  int
	_bbcgb int
}

func (_ddfg *Bitmap) GetByte(index int) (byte, error) {
	if index > len(_ddfg.Data)-1 || index < 0 {
		return 0, _d.Errorf("\u0047e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _ddfg.Data[index], nil
}
func (_baga *ClassedPoints) SortByX() { _baga._afafe = _baga.xSortFunction(); _a.Sort(_baga) }

var _ _a.Interface = &ClassedPoints{}

func (_faad *Bitmaps) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_dedag *Bitmaps, _abfa error) {
	const _adfg = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0053\u0065\u006ce\u0063\u0074\u0042\u0079Si\u007a\u0065"
	if _faad == nil {
		return nil, _d.Error(_adfg, "\u0027\u0062\u0027 B\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _d.Errorf(_adfg, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _d.Errorf(_adfg, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", relation)
	}
	_dfdde, _abfa := _faad.makeSizeIndicator(width, height, tp, relation)
	if _abfa != nil {
		return nil, _d.Wrap(_abfa, _adfg, "")
	}
	_dedag, _abfa = _faad.selectByIndicator(_dfdde)
	if _abfa != nil {
		return nil, _d.Wrap(_abfa, _adfg, "")
	}
	return _dedag, nil
}
func MorphSequence(src *Bitmap, sequence ...MorphProcess) (*Bitmap, error) {
	return _dfda(src, sequence...)
}
func (_gfeb *Bitmap) setEightBytes(_dga int, _bead uint64) error {
	_dfgg := _gfeb.RowStride - (_dga % _gfeb.RowStride)
	if _gfeb.RowStride != _gfeb.Width>>3 {
		_dfgg--
	}
	if _dfgg >= 8 {
		return _gfeb.setEightFullBytes(_dga, _bead)
	}
	return _gfeb.setEightPartlyBytes(_dga, _dfgg, _bead)
}
func (_cdcd *ClassedPoints) GetIntXByClass(i int) (int, error) {
	const _cbfdf = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _cdcd.IntSlice.Size() {
		return 0, _d.Errorf(_cbfdf, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_cdcd.XAtIndex(i)), nil
}
func _bgbe(_agg, _ebee, _beba *Bitmap) (*Bitmap, error) {
	const _eab = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0078\u006f\u0072"
	if _ebee == nil {
		return nil, _d.Error(_eab, "'\u0062\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _beba == nil {
		return nil, _d.Error(_eab, "'\u0062\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _agg == _beba {
		return nil, _d.Error(_eab, "'\u0064\u0027\u0020\u003d\u003d\u0020\u0027\u0062\u0032\u0027")
	}
	if !_ebee.SizesEqual(_beba) {
		_ea.Log.Debug("\u0025s\u0020\u002d \u0042\u0069\u0074\u006da\u0070\u0020\u0027b\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074 e\u0071\u0075\u0061l\u0020\u0073i\u007a\u0065\u0020\u0077\u0069\u0074h\u0020\u0027b\u0032\u0027", _eab)
	}
	var _eefe error
	if _agg, _eefe = _bce(_agg, _ebee); _eefe != nil {
		return nil, _d.Wrap(_eefe, _eab, "\u0063\u0061n\u0027\u0074\u0020c\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027")
	}
	if _eefe = _agg.RasterOperation(0, 0, _agg.Width, _agg.Height, PixSrcXorDst, _beba, 0, 0); _eefe != nil {
		return nil, _d.Wrap(_eefe, _eab, "")
	}
	return _agg, nil
}

const (
	_ LocationFilter = iota
	LocSelectWidth
	LocSelectHeight
	LocSelectXVal
	LocSelectYVal
	LocSelectIfEither
	LocSelectIfBoth
)

func _fccf(_def *Bitmap, _dafbaa *Bitmap, _dgaf *Selection, _ebbef **Bitmap) (*Bitmap, error) {
	const _afad = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0031"
	if _dafbaa == nil {
		return nil, _d.Error(_afad, "\u004d\u006f\u0072\u0070\u0068\u0041\u0072\u0067\u0073\u0031\u0020'\u0073\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066i\u006e\u0065\u0064")
	}
	if _dgaf == nil {
		return nil, _d.Error(_afad, "\u004d\u006f\u0072\u0068p\u0041\u0072\u0067\u0073\u0031\u0020\u0027\u0073\u0065\u006c'\u0020n\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064")
	}
	_gbgbb, _bdbda := _dgaf.Height, _dgaf.Width
	if _gbgbb == 0 || _bdbda == 0 {
		return nil, _d.Error(_afad, "\u0073\u0065\u006c\u0065ct\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _def == nil {
		_def = _dafbaa.createTemplate()
		*_ebbef = _dafbaa
		return _def, nil
	}
	_def.Width = _dafbaa.Width
	_def.Height = _dafbaa.Height
	_def.RowStride = _dafbaa.RowStride
	_def.Color = _dafbaa.Color
	_def.Data = make([]byte, _dafbaa.RowStride*_dafbaa.Height)
	if _def == _dafbaa {
		*_ebbef = _dafbaa.Copy()
	} else {
		*_ebbef = _dafbaa
	}
	return _def, nil
}
func Centroid(bm *Bitmap, centTab, sumTab []int) (Point, error) { return bm.centroid(centTab, sumTab) }
func Blit(src *Bitmap, dst *Bitmap, x, y int, op CombinationOperator) error {
	var _gfc, _ecf int
	_bbba := src.RowStride - 1
	if x < 0 {
		_ecf = -x
		x = 0
	} else if x+src.Width > dst.Width {
		_bbba -= src.Width + x - dst.Width
	}
	if y < 0 {
		_gfc = -y
		y = 0
		_ecf += src.RowStride
		_bbba += src.RowStride
	} else if y+src.Height > dst.Height {
		_gfc = src.Height + y - dst.Height
	}
	var (
		_cag  int
		_bddb error
	)
	_dbefd := x & 0x07
	_cbdd := 8 - _dbefd
	_abebe := src.Width & 0x07
	_eaaa := _cbdd - _abebe
	_aaa := _cbdd&0x07 != 0
	_baaa := src.Width <= ((_bbba-_ecf)<<3)+_cbdd
	_gdbg := dst.GetByteIndex(x, y)
	_gdgbc := _gfc + dst.Height
	if src.Height > _gdgbc {
		_cag = _gdgbc
	} else {
		_cag = src.Height
	}
	switch {
	case !_aaa:
		_bddb = _fedd(src, dst, _gfc, _cag, _gdbg, _ecf, _bbba, op)
	case _baaa:
		_bddb = _cedg(src, dst, _gfc, _cag, _gdbg, _ecf, _bbba, _eaaa, _dbefd, _cbdd, op)
	default:
		_bddb = _age(src, dst, _gfc, _cag, _gdbg, _ecf, _bbba, _eaaa, _dbefd, _cbdd, op, _abebe)
	}
	return _bddb
}

const (
	AsymmetricMorphBC BoundaryCondition = iota
	SymmetricMorphBC
)

func (_eaab *ClassedPoints) Len() int { return _eaab.IntSlice.Size() }
func _ccgeg(_efbac *Bitmap, _acef *Bitmap, _gdfdc int) (_ecad error) {
	const _dgcf = "\u0073\u0065\u0065\u0064\u0066\u0069\u006c\u006c\u0042\u0069\u006e\u0061r\u0079\u004c\u006f\u0077"
	_abdf := _fbgc(_efbac.Height, _acef.Height)
	_cgbe := _fbgc(_efbac.RowStride, _acef.RowStride)
	switch _gdfdc {
	case 4:
		_ecad = _fged(_efbac, _acef, _abdf, _cgbe)
	case 8:
		_ecad = _abbg(_efbac, _acef, _abdf, _cgbe)
	default:
		return _d.Errorf(_dgcf, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0034\u0020\u006fr\u0020\u0038\u0020\u002d\u0020i\u0073\u003a \u0027\u0025\u0064\u0027", _gdfdc)
	}
	if _ecad != nil {
		return _d.Wrap(_ecad, _dgcf, "")
	}
	return nil
}
func _gebgfd() []int {
	_ade := make([]int, 256)
	for _cadf := 0; _cadf <= 0xff; _cadf++ {
		_efag := byte(_cadf)
		_ade[_efag] = int(_efag&0x1) + (int(_efag>>1) & 0x1) + (int(_efag>>2) & 0x1) + (int(_efag>>3) & 0x1) + (int(_efag>>4) & 0x1) + (int(_efag>>5) & 0x1) + (int(_efag>>6) & 0x1) + (int(_efag>>7) & 0x1)
	}
	return _ade
}
func (_ebc *Bitmap) GetPixel(x, y int) bool {
	_bef := _ebc.GetByteIndex(x, y)
	_eaca := _ebc.GetBitOffset(x)
	_eaac := uint(7 - _eaca)
	if _bef > len(_ebc.Data)-1 {
		_ea.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0072\u0061\u006e\u0067\u0065\u002e \u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0027\u0025\u0064'\u002c\u0020\u0062m\u003a\u0020\u0027\u0025\u0073\u0027", x, y, _ebc)
		return false
	}
	if (_ebc.Data[_bef]>>_eaac)&0x01 >= 1 {
		return true
	}
	return false
}
func _fedd(_efecf, _daedb *Bitmap, _cge, _efgc, _baba, _edbec, _bdac int, _dcgb CombinationOperator) error {
	var _decc int
	_cea := func() {
		_decc++
		_baba += _daedb.RowStride
		_edbec += _efecf.RowStride
		_bdac += _efecf.RowStride
	}
	for _decc = _cge; _decc < _efgc; _cea() {
		_gffg := _baba
		for _fggd := _edbec; _fggd <= _bdac; _fggd++ {
			_fde, _gggg := _daedb.GetByte(_gffg)
			if _gggg != nil {
				return _gggg
			}
			_gdda, _gggg := _efecf.GetByte(_fggd)
			if _gggg != nil {
				return _gggg
			}
			if _gggg = _daedb.SetByte(_gffg, _adda(_fde, _gdda, _dcgb)); _gggg != nil {
				return _gggg
			}
			_gffg++
		}
	}
	return nil
}
func TstWordBitmap(t *_da.T, scale ...int) *Bitmap {
	_fbgcb := 1
	if len(scale) > 0 {
		_fbgcb = scale[0]
	}
	_ggfd := 3
	_fafa := 9 + 7 + 15 + 2*_ggfd
	_bbga := 5 + _ggfd + 5
	_fcege := New(_fafa*_fbgcb, _bbga*_fbgcb)
	_cegc := &Bitmaps{}
	var _ddcd *int
	_ggfd *= _fbgcb
	_eceag := 0
	_ddcd = &_eceag
	_eadea := 0
	_gefb := TstDSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, _ggfd)
	_gefb = TstISymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstTSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, _ggfd)
	_gefb = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstWSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 0)
	*_ddcd = 0
	_eadea = 5*_fbgcb + _ggfd
	_gefb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, _ggfd)
	_gefb = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstESymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstVSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstESymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 1*_fbgcb)
	_gefb = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cegc, _gefb, _ddcd, _eadea, 0)
	TstWriteSymbols(t, _cegc, _fcege)
	return _fcege
}
func (_gff *Bitmap) String() string {
	var _eeg = "\u000a"
	for _bccd := 0; _bccd < _gff.Height; _bccd++ {
		var _bed string
		for _dbb := 0; _dbb < _gff.Width; _dbb++ {
			_gga := _gff.GetPixel(_dbb, _bccd)
			if _gga {
				_bed += "\u0031"
			} else {
				_bed += "\u0030"
			}
		}
		_eeg += _bed + "\u000a"
	}
	return _eeg
}
func (_afgb *Bitmap) ConnComponents(bms *Bitmaps, connectivity int) (_cda *Boxes, _gfeg error) {
	const _dabd = "B\u0069\u0074\u006d\u0061p.\u0043o\u006e\u006e\u0043\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0073"
	if _afgb == nil {
		return nil, _d.Error(_dabd, "\u0070r\u006f\u0076\u0069\u0064e\u0064\u0020\u0065\u006d\u0070t\u0079 \u0027b\u0027\u0020\u0062\u0069\u0074\u006d\u0061p")
	}
	if connectivity != 4 && connectivity != 8 {
		return nil, _d.Error(_dabd, "\u0063\u006f\u006ene\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006e\u006f\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038")
	}
	if bms == nil {
		if _cda, _gfeg = _afgb.connComponentsBB(connectivity); _gfeg != nil {
			return nil, _d.Wrap(_gfeg, _dabd, "")
		}
	} else {
		if _cda, _gfeg = _afgb.connComponentsBitmapsBB(bms, connectivity); _gfeg != nil {
			return nil, _d.Wrap(_gfeg, _dabd, "")
		}
	}
	return _cda, nil
}
func (_ddc *Bitmap) ToImage() _fb.Image {
	_gcd, _cbdaf := _ed.NewImage(_ddc.Width, _ddc.Height, 1, 1, _ddc.Data, nil, nil)
	if _cbdaf != nil {
		_ea.Log.Error("\u0043\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020j\u0062\u0069\u0067\u0032\u002e\u0042\u0069\u0074m\u0061p\u0020\u0074\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002e\u0049\u006d\u0061\u0067e\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cbdaf)
	}
	return _gcd
}
func (_fgfb *ClassedPoints) Less(i, j int) bool { return _fgfb._afafe(i, j) }
func (_cbcf *Bitmaps) AddBitmap(bm *Bitmap)     { _cbcf.Values = append(_cbcf.Values, bm) }
func (_fbfad *byWidth) Swap(i, j int) {
	_fbfad.Values[i], _fbfad.Values[j] = _fbfad.Values[j], _fbfad.Values[i]
	if _fbfad.Boxes != nil {
		_fbfad.Boxes[i], _fbfad.Boxes[j] = _fbfad.Boxes[j], _fbfad.Boxes[i]
	}
}
func _dfaef(_ffbeg *Bitmap, _ccfe, _bfbe, _agff, _feec int, _agd RasterOperator, _fag *Bitmap, _ccfg, _fbdaf int) error {
	var (
		_fadga      bool
		_bedf       bool
		_cbdb       int
		_dgdcb      int
		_aecdc      int
		_ffba       bool
		_eacg       byte
		_fecb       int
		_cecdb      int
		_cddb       int
		_ace, _gafb int
	)
	_geeg := 8 - (_ccfe & 7)
	_fgdec := _eaec[_geeg]
	_bgged := _ffbeg.RowStride*_bfbe + (_ccfe >> 3)
	_dagge := _fag.RowStride*_fbdaf + (_ccfg >> 3)
	if _agff < _geeg {
		_fadga = true
		_fgdec &= _fada[8-_geeg+_agff]
	}
	if !_fadga {
		_cbdb = (_agff - _geeg) >> 3
		if _cbdb > 0 {
			_bedf = true
			_dgdcb = _bgged + 1
			_aecdc = _dagge + 1
		}
	}
	_fecb = (_ccfe + _agff) & 7
	if !(_fadga || _fecb == 0) {
		_ffba = true
		_eacg = _fada[_fecb]
		_cecdb = _bgged + 1 + _cbdb
		_cddb = _dagge + 1 + _cbdb
	}
	switch _agd {
	case PixSrc:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = _fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotSrc:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^_fag.Data[_dagge], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = ^_fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^_fag.Data[_cddb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixSrcOrDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge]|_ffbeg.Data[_bgged], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] |= _fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb]|_ffbeg.Data[_cecdb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixSrcAndDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge]&_ffbeg.Data[_bgged], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] &= _fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb]&_ffbeg.Data[_cecdb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixSrcXorDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge]^_ffbeg.Data[_bgged], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] ^= _fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb]^_ffbeg.Data[_cecdb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotSrcOrDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^(_fag.Data[_dagge])|_ffbeg.Data[_bgged], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] |= ^(_fag.Data[_aecdc+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^(_fag.Data[_cddb])|_ffbeg.Data[_cecdb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotSrcAndDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^(_fag.Data[_dagge])&_ffbeg.Data[_bgged], _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] &= ^_fag.Data[_aecdc+_gafb]
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^(_fag.Data[_cddb])&_ffbeg.Data[_cecdb], _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixSrcOrNotDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge]|^(_ffbeg.Data[_bgged]), _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = _fag.Data[_aecdc+_gafb] | ^(_ffbeg.Data[_dgdcb+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb]|^(_ffbeg.Data[_cecdb]), _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixSrcAndNotDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], _fag.Data[_dagge]&^(_ffbeg.Data[_bgged]), _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = _fag.Data[_aecdc+_gafb] &^ (_ffbeg.Data[_dgdcb+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], _fag.Data[_cddb]&^(_ffbeg.Data[_cecdb]), _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^(_fag.Data[_dagge] | _ffbeg.Data[_bgged]), _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = ^(_fag.Data[_aecdc+_gafb] | _ffbeg.Data[_dgdcb+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^(_fag.Data[_cddb] | _ffbeg.Data[_cecdb]), _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^(_fag.Data[_dagge] & _ffbeg.Data[_bgged]), _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = ^(_fag.Data[_aecdc+_gafb] & _ffbeg.Data[_dgdcb+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^(_fag.Data[_cddb] & _ffbeg.Data[_cecdb]), _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		for _ace = 0; _ace < _feec; _ace++ {
			_ffbeg.Data[_bgged] = _efee(_ffbeg.Data[_bgged], ^(_fag.Data[_dagge] ^ _ffbeg.Data[_bgged]), _fgdec)
			_bgged += _ffbeg.RowStride
			_dagge += _fag.RowStride
		}
		if _bedf {
			for _ace = 0; _ace < _feec; _ace++ {
				for _gafb = 0; _gafb < _cbdb; _gafb++ {
					_ffbeg.Data[_dgdcb+_gafb] = ^(_fag.Data[_aecdc+_gafb] ^ _ffbeg.Data[_dgdcb+_gafb])
				}
				_dgdcb += _ffbeg.RowStride
				_aecdc += _fag.RowStride
			}
		}
		if _ffba {
			for _ace = 0; _ace < _feec; _ace++ {
				_ffbeg.Data[_cecdb] = _efee(_ffbeg.Data[_cecdb], ^(_fag.Data[_cddb] ^ _ffbeg.Data[_cecdb]), _eacg)
				_cecdb += _ffbeg.RowStride
				_cddb += _fag.RowStride
			}
		}
	default:
		_ea.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _agd)
		return _d.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0056\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004c\u006f\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func _dbeae(_cfde, _ddba int, _caed string) *Selection {
	_ebcc := &Selection{Height: _cfde, Width: _ddba, Name: _caed}
	_ebcc.Data = make([][]SelectionValue, _cfde)
	for _bfgg := 0; _bfgg < _cfde; _bfgg++ {
		_ebcc.Data[_bfgg] = make([]SelectionValue, _ddba)
	}
	return _ebcc
}
func _fgcd(_gadd, _dbgdb *Bitmap, _ccec, _bbfbg int) (*Bitmap, error) {
	const _cgb = "\u0065\u0072\u006f\u0064\u0065\u0042\u0072\u0069\u0063\u006b"
	if _dbgdb == nil {
		return nil, _d.Error(_cgb, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ccec < 1 || _bbfbg < 1 {
		return nil, _d.Error(_cgb, "\u0068\u0073\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069\u007a\u0065\u0020\u0061\u0072e\u0020\u006e\u006f\u0074\u0020\u0067\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _ccec == 1 && _bbfbg == 1 {
		_acba, _agc := _bce(_gadd, _dbgdb)
		if _agc != nil {
			return nil, _d.Wrap(_agc, _cgb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _acba, nil
	}
	if _ccec == 1 || _bbfbg == 1 {
		_gbed := SelCreateBrick(_bbfbg, _ccec, _bbfbg/2, _ccec/2, SelHit)
		_adgd, _aeec := _agbc(_gadd, _dbgdb, _gbed)
		if _aeec != nil {
			return nil, _d.Wrap(_aeec, _cgb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _adgd, nil
	}
	_ebeg := SelCreateBrick(1, _ccec, 0, _ccec/2, SelHit)
	_adfc := SelCreateBrick(_bbfbg, 1, _bbfbg/2, 0, SelHit)
	_accc, _cfefg := _agbc(nil, _dbgdb, _ebeg)
	if _cfefg != nil {
		return nil, _d.Wrap(_cfefg, _cgb, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_gadd, _cfefg = _agbc(_gadd, _accc, _adfc)
	if _cfefg != nil {
		return nil, _d.Wrap(_cfefg, _cgb, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _gadd, nil
}
func (_fgaaa *Bitmap) setFourBytes(_ebfd int, _fdad uint32) error {
	if _ebfd+3 > len(_fgaaa.Data)-1 {
		return _d.Errorf("\u0073\u0065\u0074F\u006f\u0075\u0072\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _ebfd)
	}
	_fgaaa.Data[_ebfd] = byte((_fdad & 0xff000000) >> 24)
	_fgaaa.Data[_ebfd+1] = byte((_fdad & 0xff0000) >> 16)
	_fgaaa.Data[_ebfd+2] = byte((_fdad & 0xff00) >> 8)
	_fgaaa.Data[_ebfd+3] = byte(_fdad & 0xff)
	return nil
}
func (_deec *Bitmap) inverseData() {
	if _dcfg := _deec.RasterOperation(0, 0, _deec.Width, _deec.Height, PixNotDst, nil, 0, 0); _dcfg != nil {
		_ea.Log.Debug("\u0049n\u0076\u0065\u0072\u0073e\u0020\u0064\u0061\u0074\u0061 \u0066a\u0069l\u0065\u0064\u003a\u0020\u0027\u0025\u0076'", _dcfg)
	}
	if _deec.Color == Chocolate {
		_deec.Color = Vanilla
	} else {
		_deec.Color = Chocolate
	}
}
func _eec(_bbef, _ecdg int) int {
	if _bbef > _ecdg {
		return _bbef
	}
	return _ecdg
}
func (_bged *ClassedPoints) YAtIndex(i int) float32 { return (*_bged.Points)[_bged.IntSlice[i]].Y }
func _efgd(_bgdab *Bitmap) (_eagcc *Bitmap, _gbef int, _cgeb error) {
	const _dbgb = "\u0042i\u0074\u006d\u0061\u0070.\u0077\u006f\u0072\u0064\u004da\u0073k\u0042y\u0044\u0069\u006c\u0061\u0074\u0069\u006fn"
	if _bgdab == nil {
		return nil, 0, _d.Errorf(_dbgb, "\u0027\u0073\u0027\u0020bi\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	var _cfbf, _ecccd *Bitmap
	if _cfbf, _cgeb = _bce(nil, _bgdab); _cgeb != nil {
		return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0063\u006f\u0070\u0079\u0020\u0027\u0073\u0027")
	}
	var (
		_dde         [13]int
		_bfdg, _fbgg int
	)
	_gccd := 12
	_bgfg := _b.NewNumSlice(_gccd + 1)
	_bfaa := _b.NewNumSlice(_gccd + 1)
	var _eabb *Boxes
	for _afgc := 0; _afgc <= _gccd; _afgc++ {
		if _afgc == 0 {
			if _ecccd, _cgeb = _bce(nil, _cfbf); _cgeb != nil {
				return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0066i\u0072\u0073\u0074\u0020\u0062\u006d2")
			}
		} else {
			if _ecccd, _cgeb = _dfda(_cfbf, MorphProcess{Operation: MopDilation, Arguments: []int{2, 1}}); _cgeb != nil {
				return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0064\u0069\u006ca\u0074\u0069\u006f\u006e\u0020\u0062\u006d\u0032")
			}
		}
		if _eabb, _cgeb = _ecccd.connComponentsBB(4); _cgeb != nil {
			return nil, 0, _d.Wrap(_cgeb, _dbgb, "")
		}
		_dde[_afgc] = len(*_eabb)
		_bgfg.AddInt(_dde[_afgc])
		switch _afgc {
		case 0:
			_bfdg = _dde[0]
		default:
			_fbgg = _dde[_afgc-1] - _dde[_afgc]
			_bfaa.AddInt(_fbgg)
		}
		_cfbf = _ecccd
	}
	_fgeg := true
	_gdge := 2
	var _ffde, _acdd int
	for _edfcc := 1; _edfcc < len(*_bfaa); _edfcc++ {
		if _ffde, _cgeb = _bgfg.GetInt(_edfcc); _cgeb != nil {
			return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0043\u0068\u0065\u0063ki\u006e\u0067\u0020\u0062\u0065\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0069o\u006e")
		}
		if _fgeg && _ffde < int(0.3*float32(_bfdg)) {
			_gdge = _edfcc + 1
			_fgeg = false
		}
		if _fbgg, _cgeb = _bfaa.GetInt(_edfcc); _cgeb != nil {
			return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006ea\u0044\u0069\u0066\u0066")
		}
		if _fbgg > _acdd {
			_acdd = _fbgg
		}
	}
	_fcg := _bgdab.XResolution
	if _fcg == 0 {
		_fcg = 150
	}
	if _fcg > 110 {
		_gdge++
	}
	if _gdge < 2 {
		_ea.Log.Trace("J\u0042\u0049\u0047\u0032\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0042\u0065\u0073\u0074 \u0074\u006f\u0020\u006d\u0069\u006e\u0069\u006d\u0075\u006d a\u006c\u006c\u006fw\u0061b\u006c\u0065")
		_gdge = 2
	}
	_gbef = _gdge + 1
	if _eagcc, _cgeb = _gagbb(nil, _bgdab, _gdge+1, 1); _cgeb != nil {
		return nil, 0, _d.Wrap(_cgeb, _dbgb, "\u0067\u0065\u0074\u0074in\u0067\u0020\u006d\u0061\u0073\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	return _eagcc, _gbef, nil
}
func (_fee *Bitmap) Zero() bool {
	_daag := _fee.Width / 8
	_fge := _fee.Width & 7
	var _abab byte
	if _fge != 0 {
		_abab = byte(0xff << uint(8-_fge))
	}
	var _ecce, _afb, _bcdf int
	for _afb = 0; _afb < _fee.Height; _afb++ {
		_ecce = _fee.RowStride * _afb
		for _bcdf = 0; _bcdf < _daag; _bcdf, _ecce = _bcdf+1, _ecce+1 {
			if _fee.Data[_ecce] != 0 {
				return false
			}
		}
		if _fge > 0 {
			if _fee.Data[_ecce]&_abab != 0 {
				return false
			}
		}
	}
	return true
}
func (_bfab *Bitmaps) HeightSorter() func(_acdff, _gbde int) bool {
	return func(_acbf, _fefc int) bool {
		_cbbbc := _bfab.Values[_acbf].Height < _bfab.Values[_fefc].Height
		_ea.Log.Debug("H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0076\u0020<\u0020\u0025\u0076\u0020= \u0025\u0076", _bfab.Values[_acbf].Height, _bfab.Values[_fefc].Height, _cbbbc)
		return _cbbbc
	}
}

const (
	MopDilation MorphOperation = iota
	MopErosion
	MopOpening
	MopClosing
	MopRankBinaryReduction
	MopReplicativeBinaryExpansion
	MopAddBorder
)
const (
	_ SizeSelection = iota
	SizeSelectByWidth
	SizeSelectByHeight
	SizeSelectByMaxDimension
	SizeSelectByArea
	SizeSelectByPerimeter
)

func (_ddbff *ClassedPoints) SortByY() {
	_ddbff._afafe = _ddbff.ySortFunction()
	_a.Sort(_ddbff)
}
func TstNSymbol(t *_da.T, scale ...int) *Bitmap {
	_acae, _ffec := NewWithData(4, 5, []byte{0x90, 0xD0, 0xB0, 0x90, 0x90})
	_dc.NoError(t, _ffec)
	return TstGetScaledSymbol(t, _acae, scale...)
}
func (_dfde *Bitmap) connComponentsBitmapsBB(_efd *Bitmaps, _eff int) (_adff *Boxes, _dcc error) {
	const _eeaf = "\u0063\u006f\u006enC\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0069\u0074\u006d\u0061\u0070\u0073\u0042\u0042"
	if _eff != 4 && _eff != 8 {
		return nil, _d.Error(_eeaf, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _efd == nil {
		return nil, _d.Error(_eeaf, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if len(_efd.Values) > 0 {
		return nil, _d.Error(_eeaf, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u006fn\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0042\u0069\u0074m\u0061\u0070\u0073")
	}
	if _dfde.Zero() {
		return &Boxes{}, nil
	}
	var (
		_daeg, _gdfd, _bggcd, _fcaa *Bitmap
	)
	_dfde.setPadBits(0)
	if _daeg, _dcc = _bce(nil, _dfde); _dcc != nil {
		return nil, _d.Wrap(_dcc, _eeaf, "\u0062\u006d\u0031")
	}
	if _gdfd, _dcc = _bce(nil, _dfde); _dcc != nil {
		return nil, _d.Wrap(_dcc, _eeaf, "\u0062\u006d\u0032")
	}
	_fggg := &_b.Stack{}
	_fggg.Aux = &_b.Stack{}
	_adff = &Boxes{}
	var (
		_aeded, _adb int
		_bceg        _fb.Point
		_bbf         bool
		_cbdaa       *_fb.Rectangle
	)
	for {
		if _bceg, _bbf, _dcc = _daeg.nextOnPixel(_aeded, _adb); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "")
		}
		if !_bbf {
			break
		}
		if _cbdaa, _dcc = _gegb(_daeg, _fggg, _bceg.X, _bceg.Y, _eff); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "")
		}
		if _dcc = _adff.Add(_cbdaa); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "")
		}
		if _bggcd, _dcc = _daeg.clipRectangle(_cbdaa, nil); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "\u0062\u006d\u0033")
		}
		if _fcaa, _dcc = _gdfd.clipRectangle(_cbdaa, nil); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "\u0062\u006d\u0034")
		}
		if _, _dcc = _bgbe(_bggcd, _bggcd, _fcaa); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "\u0062m\u0033\u0020\u005e\u0020\u0062\u006d4")
		}
		if _dcc = _gdfd.RasterOperation(_cbdaa.Min.X, _cbdaa.Min.Y, _cbdaa.Dx(), _cbdaa.Dy(), PixSrcXorDst, _bggcd, 0, 0); _dcc != nil {
			return nil, _d.Wrap(_dcc, _eeaf, "\u0062\u006d\u0032\u0020\u002d\u0058\u004f\u0052\u002d>\u0020\u0062\u006d\u0033")
		}
		_efd.AddBitmap(_bggcd)
		_aeded = _bceg.X
		_adb = _bceg.Y
	}
	_efd.Boxes = *_adff
	return _adff, nil
}

type Selection struct {
	Height, Width int
	Cx, Cy        int
	Name          string
	Data          [][]SelectionValue
}

func (_fab *Boxes) makeSizeIndicator(_bfa, _eebg int, _cggb LocationFilter, _fdec SizeComparison) *_b.NumSlice {
	_bgea := &_b.NumSlice{}
	var _cdcgc, _cfaa, _ega int
	for _, _abaf := range *_fab {
		_cdcgc = 0
		_cfaa, _ega = _abaf.Dx(), _abaf.Dy()
		switch _cggb {
		case LocSelectWidth:
			if (_fdec == SizeSelectIfLT && _cfaa < _bfa) || (_fdec == SizeSelectIfGT && _cfaa > _bfa) || (_fdec == SizeSelectIfLTE && _cfaa <= _bfa) || (_fdec == SizeSelectIfGTE && _cfaa >= _bfa) {
				_cdcgc = 1
			}
		case LocSelectHeight:
			if (_fdec == SizeSelectIfLT && _ega < _eebg) || (_fdec == SizeSelectIfGT && _ega > _eebg) || (_fdec == SizeSelectIfLTE && _ega <= _eebg) || (_fdec == SizeSelectIfGTE && _ega >= _eebg) {
				_cdcgc = 1
			}
		case LocSelectIfEither:
			if (_fdec == SizeSelectIfLT && (_ega < _eebg || _cfaa < _bfa)) || (_fdec == SizeSelectIfGT && (_ega > _eebg || _cfaa > _bfa)) || (_fdec == SizeSelectIfLTE && (_ega <= _eebg || _cfaa <= _bfa)) || (_fdec == SizeSelectIfGTE && (_ega >= _eebg || _cfaa >= _bfa)) {
				_cdcgc = 1
			}
		case LocSelectIfBoth:
			if (_fdec == SizeSelectIfLT && (_ega < _eebg && _cfaa < _bfa)) || (_fdec == SizeSelectIfGT && (_ega > _eebg && _cfaa > _bfa)) || (_fdec == SizeSelectIfLTE && (_ega <= _eebg && _cfaa <= _bfa)) || (_fdec == SizeSelectIfGTE && (_ega >= _eebg && _cfaa >= _bfa)) {
				_cdcgc = 1
			}
		}
		_bgea.AddInt(_cdcgc)
	}
	return _bgea
}
func Dilate(d *Bitmap, s *Bitmap, sel *Selection) (*Bitmap, error) { return _aef(d, s, sel) }

type shift int

func _cgdc(_ageda *Bitmap, _dbaga, _gdgf, _adaa, _bbegc int, _fddf RasterOperator, _gbdd *Bitmap, _eade, _fcfa int) error {
	var (
		_eadg         bool
		_efeg         bool
		_dffg         byte
		_gbf          int
		_affc         int
		_fcba         int
		_ecea         int
		_beadc        bool
		_ddbb         int
		_gffa         int
		_gdfb         int
		_dcdbe        bool
		_cffa         byte
		_ccge         int
		_cee          int
		_ebed         int
		_debg         byte
		_bbfc         int
		_edba         int
		_edcb         uint
		_feag         uint
		_cgfad        byte
		_bgde         shift
		_fdcc         bool
		_cfbg         bool
		_ccbc, _fbceb int
	)
	if _eade&7 != 0 {
		_edba = 8 - (_eade & 7)
	}
	if _dbaga&7 != 0 {
		_affc = 8 - (_dbaga & 7)
	}
	if _edba == 0 && _affc == 0 {
		_cgfad = _eaec[0]
	} else {
		if _affc > _edba {
			_edcb = uint(_affc - _edba)
		} else {
			_edcb = uint(8 - (_edba - _affc))
		}
		_feag = 8 - _edcb
		_cgfad = _eaec[_edcb]
	}
	if (_dbaga & 7) != 0 {
		_eadg = true
		_gbf = 8 - (_dbaga & 7)
		_dffg = _eaec[_gbf]
		_fcba = _ageda.RowStride*_gdgf + (_dbaga >> 3)
		_ecea = _gbdd.RowStride*_fcfa + (_eade >> 3)
		_bbfc = 8 - (_eade & 7)
		if _gbf > _bbfc {
			_bgde = _fdef
			if _adaa >= _edba {
				_fdcc = true
			}
		} else {
			_bgde = _aage
		}
	}
	if _adaa < _gbf {
		_efeg = true
		_dffg &= _fada[8-_gbf+_adaa]
	}
	if !_efeg {
		_ddbb = (_adaa - _gbf) >> 3
		if _ddbb != 0 {
			_beadc = true
			_gffa = _ageda.RowStride*_gdgf + ((_dbaga + _affc) >> 3)
			_gdfb = _gbdd.RowStride*_fcfa + ((_eade + _affc) >> 3)
		}
	}
	_ccge = (_dbaga + _adaa) & 7
	if !(_efeg || _ccge == 0) {
		_dcdbe = true
		_cffa = _fada[_ccge]
		_cee = _ageda.RowStride*_gdgf + ((_dbaga + _affc) >> 3) + _ddbb
		_ebed = _gbdd.RowStride*_fcfa + ((_eade + _affc) >> 3) + _ddbb
		if _ccge > int(_feag) {
			_cfbg = true
		}
	}
	switch _fddf {
	case PixSrc:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg, _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = _debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg, _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotSrc:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^_debg, _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = ^_debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^_debg, _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixSrcOrDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg|_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] |= _debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg|_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixSrcAndDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg&_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] &= _debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg&_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixSrcXorDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg^_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] ^= _debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg^_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotSrcOrDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^_debg|_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] |= ^_debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^_debg|_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotSrcAndDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^_debg&_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] &= ^_debg
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^_debg&_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixSrcOrNotDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg|^_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = _debg | ^_ageda.Data[_gffa+_fbceb]
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg|^_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixSrcAndNotDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], _debg&^_ageda.Data[_fcba], _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = _debg &^ _ageda.Data[_gffa+_fbceb]
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], _debg&^_ageda.Data[_cee], _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^(_debg | _ageda.Data[_fcba]), _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = ^(_debg | _ageda.Data[_gffa+_fbceb])
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^(_debg | _ageda.Data[_cee]), _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^(_debg & _ageda.Data[_fcba]), _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = ^(_debg & _ageda.Data[_gffa+_fbceb])
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^(_debg & _ageda.Data[_cee]), _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		if _eadg {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				if _bgde == _fdef {
					_debg = _gbdd.Data[_ecea] << _edcb
					if _fdcc {
						_debg = _efee(_debg, _gbdd.Data[_ecea+1]>>_feag, _cgfad)
					}
				} else {
					_debg = _gbdd.Data[_ecea] >> _feag
				}
				_ageda.Data[_fcba] = _efee(_ageda.Data[_fcba], ^(_debg ^ _ageda.Data[_fcba]), _dffg)
				_fcba += _ageda.RowStride
				_ecea += _gbdd.RowStride
			}
		}
		if _beadc {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				for _fbceb = 0; _fbceb < _ddbb; _fbceb++ {
					_debg = _efee(_gbdd.Data[_gdfb+_fbceb]<<_edcb, _gbdd.Data[_gdfb+_fbceb+1]>>_feag, _cgfad)
					_ageda.Data[_gffa+_fbceb] = ^(_debg ^ _ageda.Data[_gffa+_fbceb])
				}
				_gffa += _ageda.RowStride
				_gdfb += _gbdd.RowStride
			}
		}
		if _dcdbe {
			for _ccbc = 0; _ccbc < _bbegc; _ccbc++ {
				_debg = _gbdd.Data[_ebed] << _edcb
				if _cfbg {
					_debg = _efee(_debg, _gbdd.Data[_ebed+1]>>_feag, _cgfad)
				}
				_ageda.Data[_cee] = _efee(_ageda.Data[_cee], ^(_debg ^ _ageda.Data[_cee]), _cffa)
				_cee += _ageda.RowStride
				_ebed += _gbdd.RowStride
			}
		}
	default:
		_ea.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _fddf)
		return _d.Error("\u0072a\u0073t\u0065\u0072\u004f\u0070\u0047e\u006e\u0065r\u0061\u006c\u004c\u006f\u0077", "\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

type byHeight Bitmaps

func _feef(_ceac *_b.Stack) (_dafa *fillSegment, _cfefd error) {
	const _dbbc = "\u0070\u006f\u0070\u0046\u0069\u006c\u006c\u0053\u0065g\u006d\u0065\u006e\u0074"
	if _ceac == nil {
		return nil, _d.Error(_dbbc, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _ceac.Aux == nil {
		return nil, _d.Error(_dbbc, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	_cfbac, _cede := _ceac.Pop()
	if !_cede {
		return nil, nil
	}
	_gcded, _cede := _cfbac.(*fillSegment)
	if !_cede {
		return nil, _d.Error(_dbbc, "\u0073\u0074\u0061ck\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020c\u006fn\u0074a\u0069n\u0020\u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_dafa = &fillSegment{_gcded._gebc, _gcded._fdeg, _gcded._aegc + _gcded._bbcgb, _gcded._bbcgb}
	_ceac.Aux.Push(_gcded)
	return _dafa, nil
}
func _efee(_bbac, _geaa, _ddca byte) byte { return (_bbac &^ (_ddca)) | (_geaa & _ddca) }
func _ecfg(_feg, _fedc *Bitmap, _cebc CombinationOperator) *Bitmap {
	_cgaa := New(_feg.Width, _feg.Height)
	for _dbge := 0; _dbge < len(_cgaa.Data); _dbge++ {
		_cgaa.Data[_dbge] = _adda(_feg.Data[_dbge], _fedc.Data[_dbge], _cebc)
	}
	return _cgaa
}
func _cef(_gdef, _dcdc *Bitmap, _edbda, _faec int) (*Bitmap, error) {
	const _dfae = "d\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063\u006b"
	if _dcdc == nil {
		_ea.Log.Debug("\u0064\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063k\u0020\u0073\u006f\u0075\u0072\u0063\u0065 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
		return nil, _d.Error(_dfae, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _edbda < 1 || _faec < 1 {
		return nil, _d.Error(_dfae, "\u0068\u0053\u007a\u0069\u0065 \u0061\u006e\u0064\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u0061\u0072\u0065 \u006e\u006f\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _edbda == 1 && _faec == 1 {
		_bde, _bdaf := _bce(_gdef, _dcdc)
		if _bdaf != nil {
			return nil, _d.Wrap(_bdaf, _dfae, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _bde, nil
	}
	if _edbda == 1 || _faec == 1 {
		_ecgd := SelCreateBrick(_faec, _edbda, _faec/2, _edbda/2, SelHit)
		_fdbg, _fbfa := _aef(_gdef, _dcdc, _ecgd)
		if _fbfa != nil {
			return nil, _d.Wrap(_fbfa, _dfae, "\u0068s\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _fdbg, nil
	}
	_fcdf := SelCreateBrick(1, _edbda, 0, _edbda/2, SelHit)
	_gaea := SelCreateBrick(_faec, 1, _faec/2, 0, SelHit)
	_ceg, _efaeg := _aef(nil, _dcdc, _fcdf)
	if _efaeg != nil {
		return nil, _d.Wrap(_efaeg, _dfae, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_gdef, _efaeg = _aef(_gdef, _ceg, _gaea)
	if _efaeg != nil {
		return nil, _d.Wrap(_efaeg, _dfae, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _gdef, nil
}
func CorrelationScore(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_abdc float64, _gded error) {
	const _eece = "\u0063\u006fr\u0072\u0065\u006ca\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065"
	if bm1 == nil || bm2 == nil {
		return 0, _d.Error(_eece, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0062\u0069\u0074ma\u0070\u0073")
	}
	if tab == nil {
		return 0, _d.Error(_eece, "\u0027\u0074\u0061\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if area1 <= 0 || area2 <= 0 {
		return 0, _d.Error(_eece, "\u0061\u0072\u0065\u0061s\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067r\u0065a\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
	}
	_cdec, _afed := bm1.Width, bm1.Height
	_fgbf, _fccg := bm2.Width, bm2.Height
	_agba := _egfg(_cdec - _fgbf)
	if _agba > maxDiffW {
		return 0, nil
	}
	_gcde := _egfg(_afed - _fccg)
	if _gcde > maxDiffH {
		return 0, nil
	}
	var _cbga, _aaad int
	if delX >= 0 {
		_cbga = int(delX + 0.5)
	} else {
		_cbga = int(delX - 0.5)
	}
	if delY >= 0 {
		_aaad = int(delY + 0.5)
	} else {
		_aaad = int(delY - 0.5)
	}
	_ffb := _eec(_aaad, 0)
	_gfg := _fbgc(_fccg+_aaad, _afed)
	_dge := bm1.RowStride * _ffb
	_bbfe := bm2.RowStride * (_ffb - _aaad)
	_cbdcc := _eec(_cbga, 0)
	_fgce := _fbgc(_fgbf+_cbga, _cdec)
	_efbc := bm2.RowStride
	var _feb, _dada int
	if _cbga >= 8 {
		_feb = _cbga >> 3
		_dge += _feb
		_cbdcc -= _feb << 3
		_fgce -= _feb << 3
		_cbga &= 7
	} else if _cbga <= -8 {
		_dada = -((_cbga + 7) >> 3)
		_bbfe += _dada
		_efbc -= _dada
		_cbga += _dada << 3
	}
	if _cbdcc >= _fgce || _ffb >= _gfg {
		return 0, nil
	}
	_eafa := (_fgce + 7) >> 3
	var (
		_bace, _dcdg, _bcdb byte
		_dgdf, _eddg, _adc  int
	)
	switch {
	case _cbga == 0:
		for _adc = _ffb; _adc < _gfg; _adc, _dge, _bbfe = _adc+1, _dge+bm1.RowStride, _bbfe+bm2.RowStride {
			for _eddg = 0; _eddg < _eafa; _eddg++ {
				_bcdb = bm1.Data[_dge+_eddg] & bm2.Data[_bbfe+_eddg]
				_dgdf += tab[_bcdb]
			}
		}
	case _cbga > 0:
		if _efbc < _eafa {
			for _adc = _ffb; _adc < _gfg; _adc, _dge, _bbfe = _adc+1, _dge+bm1.RowStride, _bbfe+bm2.RowStride {
				_bace, _dcdg = bm1.Data[_dge], bm2.Data[_bbfe]>>uint(_cbga)
				_bcdb = _bace & _dcdg
				_dgdf += tab[_bcdb]
				for _eddg = 1; _eddg < _efbc; _eddg++ {
					_bace, _dcdg = bm1.Data[_dge+_eddg], (bm2.Data[_bbfe+_eddg]>>uint(_cbga))|(bm2.Data[_bbfe+_eddg-1]<<uint(8-_cbga))
					_bcdb = _bace & _dcdg
					_dgdf += tab[_bcdb]
				}
				_bace = bm1.Data[_dge+_eddg]
				_dcdg = bm2.Data[_bbfe+_eddg-1] << uint(8-_cbga)
				_bcdb = _bace & _dcdg
				_dgdf += tab[_bcdb]
			}
		} else {
			for _adc = _ffb; _adc < _gfg; _adc, _dge, _bbfe = _adc+1, _dge+bm1.RowStride, _bbfe+bm2.RowStride {
				_bace, _dcdg = bm1.Data[_dge], bm2.Data[_bbfe]>>uint(_cbga)
				_bcdb = _bace & _dcdg
				_dgdf += tab[_bcdb]
				for _eddg = 1; _eddg < _eafa; _eddg++ {
					_bace = bm1.Data[_dge+_eddg]
					_dcdg = (bm2.Data[_bbfe+_eddg] >> uint(_cbga)) | (bm2.Data[_bbfe+_eddg-1] << uint(8-_cbga))
					_bcdb = _bace & _dcdg
					_dgdf += tab[_bcdb]
				}
			}
		}
	default:
		if _eafa < _efbc {
			for _adc = _ffb; _adc < _gfg; _adc, _dge, _bbfe = _adc+1, _dge+bm1.RowStride, _bbfe+bm2.RowStride {
				for _eddg = 0; _eddg < _eafa; _eddg++ {
					_bace = bm1.Data[_dge+_eddg]
					_dcdg = bm2.Data[_bbfe+_eddg] << uint(-_cbga)
					_dcdg |= bm2.Data[_bbfe+_eddg+1] >> uint(8+_cbga)
					_bcdb = _bace & _dcdg
					_dgdf += tab[_bcdb]
				}
			}
		} else {
			for _adc = _ffb; _adc < _gfg; _adc, _dge, _bbfe = _adc+1, _dge+bm1.RowStride, _bbfe+bm2.RowStride {
				for _eddg = 0; _eddg < _eafa-1; _eddg++ {
					_bace = bm1.Data[_dge+_eddg]
					_dcdg = bm2.Data[_bbfe+_eddg] << uint(-_cbga)
					_dcdg |= bm2.Data[_bbfe+_eddg+1] >> uint(8+_cbga)
					_bcdb = _bace & _dcdg
					_dgdf += tab[_bcdb]
				}
				_bace = bm1.Data[_dge+_eddg]
				_dcdg = bm2.Data[_bbfe+_eddg] << uint(-_cbga)
				_bcdb = _bace & _dcdg
				_dgdf += tab[_bcdb]
			}
		}
	}
	_abdc = float64(_dgdf) * float64(_dgdf) / (float64(area1) * float64(area2))
	return _abdc, nil
}
func (_gebgf *Bitmap) setEightPartlyBytes(_gbcd, _dbac int, _aab uint64) (_dbba error) {
	var (
		_gbga byte
		_fadg int
	)
	const _faeg = "\u0073\u0065\u0074\u0045ig\u0068\u0074\u0050\u0061\u0072\u0074\u006c\u0079\u0042\u0079\u0074\u0065\u0073"
	for _gdbd := 1; _gdbd <= _dbac; _gdbd++ {
		_fadg = 64 - _gdbd*8
		_gbga = byte(_aab >> uint(_fadg) & 0xff)
		_ea.Log.Trace("\u0074\u0065\u006d\u0070\u003a\u0020\u0025\u0030\u0038\u0062\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a %\u0064,\u0020\u0069\u0064\u0078\u003a\u0020\u0025\u0064\u002c\u0020\u0066\u0075l\u006c\u0042\u0079\u0074\u0065\u0073\u004e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c \u0073\u0068\u0069\u0066\u0074\u003a\u0020\u0025\u0064", _gbga, _gbcd, _gbcd+_gdbd-1, _dbac, _fadg)
		if _dbba = _gebgf.SetByte(_gbcd+_gdbd-1, _gbga); _dbba != nil {
			return _d.Wrap(_dbba, _faeg, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
		}
	}
	_bbeb := _gebgf.RowStride*8 - _gebgf.Width
	if _bbeb == 0 {
		return nil
	}
	_fadg -= 8
	_gbga = byte(_aab>>uint(_fadg)&0xff) << uint(_bbeb)
	if _dbba = _gebgf.SetByte(_gbcd+_dbac, _gbga); _dbba != nil {
		return _d.Wrap(_dbba, _faeg, "\u0070\u0061\u0064\u0064\u0065\u0064")
	}
	return nil
}

type Bitmaps struct {
	Values []*Bitmap
	Boxes  []*_fb.Rectangle
}

func (_egb *Bitmap) SetDefaultPixel() {
	for _bfgb := range _egb.Data {
		_egb.Data[_bfgb] = byte(0xff)
	}
}
func _adbc(_dfggg, _geag, _baad *Bitmap, _bbc int) (*Bitmap, error) {
	const _fbdae = "\u0073\u0065\u0065\u0064\u0046\u0069\u006c\u006c\u0042i\u006e\u0061\u0072\u0079"
	if _geag == nil {
		return nil, _d.Error(_fbdae, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _baad == nil {
		return nil, _d.Error(_fbdae, "'\u006da\u0073\u006b\u0027\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _bbc != 4 && _bbc != 8 {
		return nil, _d.Error(_fbdae, "\u0063\u006f\u006en\u0065\u0063\u0074\u0069v\u0069\u0074\u0079\u0020\u006e\u006f\u0074 \u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0034\u002c\u0038\u007d")
	}
	var _ffga error
	_dfggg, _ffga = _bce(_dfggg, _geag)
	if _ffga != nil {
		return nil, _d.Wrap(_ffga, _fbdae, "\u0063o\u0070y\u0020\u0073\u006f\u0075\u0072c\u0065\u0020t\u006f\u0020\u0027\u0064\u0027")
	}
	_dcacf := _geag.createTemplate()
	_baad.setPadBits(0)
	for _fcgf := 0; _fcgf < _ceea; _fcgf++ {
		_dcacf, _ffga = _bce(_dcacf, _dfggg)
		if _ffga != nil {
			return nil, _d.Wrapf(_ffga, _fbdae, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _fcgf)
		}
		if _ffga = _ccgeg(_dfggg, _baad, _bbc); _ffga != nil {
			return nil, _d.Wrapf(_ffga, _fbdae, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _fcgf)
		}
		if _dcacf.Equals(_dfggg) {
			break
		}
	}
	return _dfggg, nil
}
func Extract(roi _fb.Rectangle, src *Bitmap) (*Bitmap, error) {
	_egcd := New(roi.Dx(), roi.Dy())
	_fgdf := roi.Min.X & 0x07
	_fdge := 8 - _fgdf
	_feaa := uint(8 - _egcd.Width&0x07)
	_gfba := src.GetByteIndex(roi.Min.X, roi.Min.Y)
	_dgcea := src.GetByteIndex(roi.Max.X-1, roi.Min.Y)
	_bae := _egcd.RowStride == _dgcea+1-_gfba
	var _ffe int
	for _edbbe := roi.Min.Y; _edbbe < roi.Max.Y; _edbbe++ {
		_ecde := _gfba
		_gadc := _ffe
		switch {
		case _gfba == _dgcea:
			_efge, _bfe := src.GetByte(_ecde)
			if _bfe != nil {
				return nil, _bfe
			}
			_efge <<= uint(_fgdf)
			_bfe = _egcd.SetByte(_gadc, _gcc(_feaa, _efge))
			if _bfe != nil {
				return nil, _bfe
			}
		case _fgdf == 0:
			for _ffea := _gfba; _ffea <= _dgcea; _ffea++ {
				_bffd, _fcce := src.GetByte(_ecde)
				if _fcce != nil {
					return nil, _fcce
				}
				_ecde++
				if _ffea == _dgcea && _bae {
					_bffd = _gcc(_feaa, _bffd)
				}
				_fcce = _egcd.SetByte(_gadc, _bffd)
				if _fcce != nil {
					return nil, _fcce
				}
				_gadc++
			}
		default:
			_gfd := _eagc(src, _egcd, uint(_fgdf), uint(_fdge), _feaa, _gfba, _dgcea, _bae, _ecde, _gadc)
			if _gfd != nil {
				return nil, _gfd
			}
		}
		_gfba += src.RowStride
		_dgcea += src.RowStride
		_ffe += _egcd.RowStride
	}
	return _egcd, nil
}
func (_feed *ClassedPoints) GroupByY() ([]*ClassedPoints, error) {
	const _fgee = "\u0043\u006c\u0061\u0073se\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0072\u006f\u0075\u0070\u0042y\u0059"
	if _bebc := _feed.validateIntSlice(); _bebc != nil {
		return nil, _d.Wrap(_bebc, _fgee, "")
	}
	if _feed.IntSlice.Size() == 0 {
		return nil, _d.Error(_fgee, "\u004e\u006f\u0020\u0063la\u0073\u0073\u0065\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_feed.SortByY()
	var (
		_dgff []*ClassedPoints
		_ebec int
	)
	_dcec := -1
	var _gbcge *ClassedPoints
	for _cccd := 0; _cccd < len(_feed.IntSlice); _cccd++ {
		_ebec = int(_feed.YAtIndex(_cccd))
		if _ebec != _dcec {
			_gbcge = &ClassedPoints{Points: _feed.Points}
			_dcec = _ebec
			_dgff = append(_dgff, _gbcge)
		}
		_gbcge.IntSlice = append(_gbcge.IntSlice, _feed.IntSlice[_cccd])
	}
	for _, _cbdf := range _dgff {
		_cbdf.SortByX()
	}
	return _dgff, nil
}
func (_fgaba *byWidth) Less(i, j int) bool { return _fgaba.Values[i].Width < _fgaba.Values[j].Width }
func (_aeagb MorphProcess) verify(_efc int, _dceb, _cebce *int) error {
	const _bbbf = "\u004d\u006f\u0072\u0070hP\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0076\u0065\u0072\u0069\u0066\u0079"
	switch _aeagb.Operation {
	case MopDilation, MopErosion, MopOpening, MopClosing:
		if len(_aeagb.Arguments) != 2 {
			return _d.Error(_bbbf, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0064\u0027\u002c\u0020\u0027\u0065\u0027\u002c \u0027\u006f\u0027\u002c\u0020\u0027\u0063\u0027\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0032\u0020\u0061r\u0067\u0075\u006d\u0065\u006et\u0073")
		}
		_dddb, _cfebc := _aeagb.getWidthHeight()
		if _dddb <= 0 || _cfebc <= 0 {
			return _d.Error(_bbbf, "O\u0070er\u0061t\u0069o\u006e\u003a\u0020\u0027\u0064'\u002c\u0020\u0027e\u0027\u002c\u0020\u0027\u006f'\u002c\u0020\u0027c\u0027\u0020\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u006f\u0074h w\u0069\u0064\u0074\u0068\u0020\u0061n\u0064\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0074\u006f\u0020b\u0065 \u003e\u003d\u0020\u0030")
		}
	case MopRankBinaryReduction:
		_cedf := len(_aeagb.Arguments)
		*_dceb += _cedf
		if _cedf < 1 || _cedf > 4 {
			return _d.Error(_bbbf, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0072\u0027\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0034\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0073")
		}
		for _bdc := 0; _bdc < _cedf; _bdc++ {
			if _aeagb.Arguments[_bdc] < 1 || _aeagb.Arguments[_bdc] > 4 {
				return _d.Error(_bbbf, "\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0052\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 \u00280\u002c\u0020\u0034\u003e")
			}
		}
	case MopReplicativeBinaryExpansion:
		if len(_aeagb.Arguments) == 0 {
			return _d.Error(_bbbf, "\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074i\u0076\u0065\u0042in\u0061\u0072\u0079\u0045\u0078\u0070a\u006e\u0073\u0069\u006f\u006e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_ccca := _aeagb.Arguments[0]
		if _ccca != 2 && _ccca != 4 && _ccca != 8 {
			return _d.Error(_bbbf, "R\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u0076\u0065\u0042\u0069\u006e\u0061\u0072\u0079\u0045\u0078\u0070\u0061\u006e\u0073\u0069\u006f\u006e\u0020m\u0075s\u0074\u0020\u0062\u0065 \u006f\u0066 \u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
		}
		*_dceb -= _eged[_ccca/4]
	case MopAddBorder:
		if len(_aeagb.Arguments) == 0 {
			return _d.Error(_bbbf, "\u0041\u0064\u0064B\u006f\u0072\u0064\u0065r\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_bgag := _aeagb.Arguments[0]
		if _efc > 0 {
			return _d.Error(_bbbf, "\u0041\u0064\u0064\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020f\u0069\u0072\u0073\u0074\u0020\u006d\u006f\u0072\u0070\u0068\u0020\u0070\u0072o\u0063\u0065\u0073\u0073")
		}
		if _bgag < 1 {
			return _d.Error(_bbbf, "\u0041\u0064\u0064\u0042o\u0072\u0064\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065 \u006co\u0077\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
		}
		*_cebce = _bgag
	}
	return nil
}
func Centroids(bms []*Bitmap) (*Points, error) {
	_befe := make([]Point, len(bms))
	_agbad := _dbag()
	_gcfeg := _gebgfd()
	var _adad error
	for _ffbe, _dcdf := range bms {
		_befe[_ffbe], _adad = _dcdf.centroid(_agbad, _gcfeg)
		if _adad != nil {
			return nil, _adad
		}
	}
	_dgab := Points(_befe)
	return &_dgab, nil
}
func (_agb *Bitmap) clearAll() error {
	return _agb.RasterOperation(0, 0, _agb.Width, _agb.Height, PixClr, nil, 0, 0)
}

const (
	CmbOpOr CombinationOperator = iota
	CmbOpAnd
	CmbOpXor
	CmbOpXNor
	CmbOpReplace
	CmbOpNot
)

func _bdg(_bcc, _ee *Bitmap, _cgfe int, _gdc []byte, _fad int) (_cbc error) {
	const _ddg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0032"
	var (
		_bfg, _abd, _gef, _gea, _dbdd, _dabc, _fae, _fbc int
		_eea, _ddb, _aae, _eag                           uint32
		_dcbg, _af                                       byte
		_edbf                                            uint16
	)
	_adg := make([]byte, 4)
	_fbg := make([]byte, 4)
	for _gef = 0; _gef < _bcc.Height-1; _gef, _gea = _gef+2, _gea+1 {
		_bfg = _gef * _bcc.RowStride
		_abd = _gea * _ee.RowStride
		for _dbdd, _dabc = 0, 0; _dbdd < _fad; _dbdd, _dabc = _dbdd+4, _dabc+1 {
			for _fae = 0; _fae < 4; _fae++ {
				_fbc = _bfg + _dbdd + _fae
				if _fbc <= len(_bcc.Data)-1 && _fbc < _bfg+_bcc.RowStride {
					_adg[_fae] = _bcc.Data[_fbc]
				} else {
					_adg[_fae] = 0x00
				}
				_fbc = _bfg + _bcc.RowStride + _dbdd + _fae
				if _fbc <= len(_bcc.Data)-1 && _fbc < _bfg+(2*_bcc.RowStride) {
					_fbg[_fae] = _bcc.Data[_fbc]
				} else {
					_fbg[_fae] = 0x00
				}
			}
			_eea = _df.BigEndian.Uint32(_adg)
			_ddb = _df.BigEndian.Uint32(_fbg)
			_aae = _eea & _ddb
			_aae |= _aae << 1
			_eag = _eea | _ddb
			_eag &= _eag << 1
			_ddb = _aae | _eag
			_ddb &= 0xaaaaaaaa
			_eea = _ddb | (_ddb << 7)
			_dcbg = byte(_eea >> 24)
			_af = byte((_eea >> 8) & 0xff)
			_fbc = _abd + _dabc
			if _fbc+1 == len(_ee.Data)-1 || _fbc+1 >= _abd+_ee.RowStride {
				if _cbc = _ee.SetByte(_fbc, _gdc[_dcbg]); _cbc != nil {
					return _d.Wrapf(_cbc, _ddg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fbc)
				}
			} else {
				_edbf = (uint16(_gdc[_dcbg]) << 8) | uint16(_gdc[_af])
				if _cbc = _ee.setTwoBytes(_fbc, _edbf); _cbc != nil {
					return _d.Wrapf(_cbc, _ddg, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fbc)
				}
				_dabc++
			}
		}
	}
	return nil
}
func _bbd(_gaf *Bitmap, _ebfc int, _eg []byte) (_dcd *Bitmap, _bea error) {
	const _cdg = "\u0072\u0065\u0064\u0075\u0063\u0065\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0032"
	if _gaf == nil {
		return nil, _d.Error(_cdg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _ebfc < 1 || _ebfc > 4 {
		return nil, _d.Error(_cdg, "\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _gaf.Height <= 1 {
		return nil, _d.Errorf(_cdg, "\u0073o\u0075\u0072c\u0065\u0020\u0068e\u0069\u0067\u0068\u0074\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027\u0020-\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", _gaf.Height)
	}
	_dcd = New(_gaf.Width/2, _gaf.Height/2)
	if _eg == nil {
		_eg = _gdgb()
	}
	_ded := _fbgc(_gaf.RowStride, 2*_dcd.RowStride)
	switch _ebfc {
	case 1:
		_bea = _bgg(_gaf, _dcd, _ebfc, _eg, _ded)
	case 2:
		_bea = _bdg(_gaf, _dcd, _ebfc, _eg, _ded)
	case 3:
		_bea = _cfd(_gaf, _dcd, _ebfc, _eg, _ded)
	case 4:
		_bea = _fdf(_gaf, _dcd, _ebfc, _eg, _ded)
	}
	if _bea != nil {
		return nil, _bea
	}
	return _dcd, nil
}
func (_fabaa *BitmapsArray) AddBitmaps(bm *Bitmaps) { _fabaa.Values = append(_fabaa.Values, bm) }

const (
	Vanilla Color = iota
	Chocolate
)

func (_edbdc *Bitmaps) Size() int { return len(_edbdc.Values) }
func (_beab *ClassedPoints) xSortFunction() func(_cdgf int, _fcbf int) bool {
	return func(_bebb, _dcad int) bool { return _beab.XAtIndex(_bebb) < _beab.XAtIndex(_dcad) }
}
func (_ebba *Bitmap) RemoveBorder(borderSize int) (*Bitmap, error) {
	if borderSize == 0 {
		return _ebba.Copy(), nil
	}
	_aecd, _bgge := _ebba.removeBorderGeneral(borderSize, borderSize, borderSize, borderSize)
	if _bgge != nil {
		return nil, _d.Wrap(_bgge, "\u0052\u0065\u006do\u0076\u0065\u0042\u006f\u0072\u0064\u0065\u0072", "")
	}
	return _aecd, nil
}
func _aeeg(_cggbg *Bitmap, _eefee, _egeg int, _bfdf, _decd int, _fcfc RasterOperator, _ffce *Bitmap, _cccag, _cabb int) error {
	var _cgc, _bcab, _bgbfg, _egcff int
	if _eefee < 0 {
		_cccag -= _eefee
		_bfdf += _eefee
		_eefee = 0
	}
	if _cccag < 0 {
		_eefee -= _cccag
		_bfdf += _cccag
		_cccag = 0
	}
	_cgc = _eefee + _bfdf - _cggbg.Width
	if _cgc > 0 {
		_bfdf -= _cgc
	}
	_bcab = _cccag + _bfdf - _ffce.Width
	if _bcab > 0 {
		_bfdf -= _bcab
	}
	if _egeg < 0 {
		_cabb -= _egeg
		_decd += _egeg
		_egeg = 0
	}
	if _cabb < 0 {
		_egeg -= _cabb
		_decd += _cabb
		_cabb = 0
	}
	_bgbfg = _egeg + _decd - _cggbg.Height
	if _bgbfg > 0 {
		_decd -= _bgbfg
	}
	_egcff = _cabb + _decd - _ffce.Height
	if _egcff > 0 {
		_decd -= _egcff
	}
	if _bfdf <= 0 || _decd <= 0 {
		return nil
	}
	var _edae error
	switch {
	case _eefee&7 == 0 && _cccag&7 == 0:
		_edae = _ddab(_cggbg, _eefee, _egeg, _bfdf, _decd, _fcfc, _ffce, _cccag, _cabb)
	case _eefee&7 == _cccag&7:
		_edae = _dfaef(_cggbg, _eefee, _egeg, _bfdf, _decd, _fcfc, _ffce, _cccag, _cabb)
	default:
		_edae = _cgdc(_cggbg, _eefee, _egeg, _bfdf, _decd, _fcfc, _ffce, _cccag, _cabb)
	}
	if _edae != nil {
		return _d.Wrap(_edae, "r\u0061\u0073\u0074\u0065\u0072\u004f\u0070\u004c\u006f\u0077", "")
	}
	return nil
}

var _adba = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x78, 0x27, 0xC2, 0x27, 0x91, 0x00, 0x22, 0x48, 0x21, 0x03, 0x24, 0x91, 0x00, 0x22, 0x48, 0x21, 0x02, 0xA4, 0x95, 0x00, 0x22, 0x48, 0x21, 0x02, 0x64, 0x9B, 0x00, 0x3C, 0x78, 0x21, 0x02, 0x27, 0x91, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func (_dbae *Boxes) Add(box *_fb.Rectangle) error {
	if _dbae == nil {
		return _d.Error("\u0042o\u0078\u0065\u0073\u002e\u0041\u0064d", "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	*_dbae = append(*_dbae, box)
	return nil
}
func CorrelationScoreThresholded(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab, downcount []int, scoreThreshold float32) (bool, error) {
	const _ffbg = "C\u006f\u0072\u0072\u0065\u006c\u0061t\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054h\u0072\u0065\u0073h\u006fl\u0064\u0065\u0064"
	if bm1 == nil {
		return false, _d.Error(_ffbg, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d1\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if bm2 == nil {
		return false, _d.Error(_ffbg, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d2\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if area1 <= 0 || area2 <= 0 {
		return false, _d.Error(_ffbg, "c\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0053\u0063\u006f\u0072\u0065\u0054\u0068re\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u002d\u0020\u0061\u0072\u0065\u0061s \u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u003e\u0020\u0030")
	}
	if downcount == nil {
		return false, _d.Error(_ffbg, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u006f\u0020\u0027\u0064\u006f\u0077\u006e\u0063\u006f\u0075\u006e\u0074\u0027")
	}
	if tab == nil {
		return false, _d.Error(_ffbg, "p\u0072\u006f\u0076\u0069de\u0064 \u006e\u0069\u006c\u0020\u0027s\u0075\u006d\u0074\u0061\u0062\u0027")
	}
	_fgaae, _acbe := bm1.Width, bm1.Height
	_efae, _ggag := bm2.Width, bm2.Height
	if _b.Abs(_fgaae-_efae) > maxDiffW {
		return false, nil
	}
	if _b.Abs(_acbe-_ggag) > maxDiffH {
		return false, nil
	}
	_bddc := int(delX + _b.Sign(delX)*0.5)
	_ebff := int(delY + _b.Sign(delY)*0.5)
	_fffe := int(_c.Ceil(_c.Sqrt(float64(scoreThreshold) * float64(area1) * float64(area2))))
	_cdba := bm2.RowStride
	_fcceb := _eec(_ebff, 0)
	_egcf := _fbgc(_ggag+_ebff, _acbe)
	_accb := bm1.RowStride * _fcceb
	_egdb := bm2.RowStride * (_fcceb - _ebff)
	var _dgac int
	if _egcf <= _acbe {
		_dgac = downcount[_egcf-1]
	}
	_fcde := _eec(_bddc, 0)
	_aeeeb := _fbgc(_efae+_bddc, _fgaae)
	var _ggdd, _caeg int
	if _bddc >= 8 {
		_ggdd = _bddc >> 3
		_accb += _ggdd
		_fcde -= _ggdd << 3
		_aeeeb -= _ggdd << 3
		_bddc &= 7
	} else if _bddc <= -8 {
		_caeg = -((_bddc + 7) >> 3)
		_egdb += _caeg
		_cdba -= _caeg
		_bddc += _caeg << 3
	}
	var (
		_abb, _aeag, _gfbf  int
		_bgbf, _ddeg, _baag byte
	)
	if _fcde >= _aeeeb || _fcceb >= _egcf {
		return false, nil
	}
	_cbfd := (_aeeeb + 7) >> 3
	switch {
	case _bddc == 0:
		for _aeag = _fcceb; _aeag < _egcf; _aeag, _accb, _egdb = _aeag+1, _accb+bm1.RowStride, _egdb+bm2.RowStride {
			for _gfbf = 0; _gfbf < _cbfd; _gfbf++ {
				_bgbf = bm1.Data[_accb+_gfbf] & bm2.Data[_egdb+_gfbf]
				_abb += tab[_bgbf]
			}
			if _abb >= _fffe {
				return true, nil
			}
			if _abg := _abb + downcount[_aeag] - _dgac; _abg < _fffe {
				return false, nil
			}
		}
	case _bddc > 0 && _cdba < _cbfd:
		for _aeag = _fcceb; _aeag < _egcf; _aeag, _accb, _egdb = _aeag+1, _accb+bm1.RowStride, _egdb+bm2.RowStride {
			_ddeg = bm1.Data[_accb]
			_baag = bm2.Data[_egdb] >> uint(_bddc)
			_bgbf = _ddeg & _baag
			_abb += tab[_bgbf]
			for _gfbf = 1; _gfbf < _cdba; _gfbf++ {
				_ddeg = bm1.Data[_accb+_gfbf]
				_baag = bm2.Data[_egdb+_gfbf]>>uint(_bddc) | bm2.Data[_egdb+_gfbf-1]<<uint(8-_bddc)
				_bgbf = _ddeg & _baag
				_abb += tab[_bgbf]
			}
			_ddeg = bm1.Data[_accb+_gfbf]
			_baag = bm2.Data[_egdb+_gfbf-1] << uint(8-_bddc)
			_bgbf = _ddeg & _baag
			_abb += tab[_bgbf]
			if _abb >= _fffe {
				return true, nil
			} else if _abb+downcount[_aeag]-_dgac < _fffe {
				return false, nil
			}
		}
	case _bddc > 0 && _cdba >= _cbfd:
		for _aeag = _fcceb; _aeag < _egcf; _aeag, _accb, _egdb = _aeag+1, _accb+bm1.RowStride, _egdb+bm2.RowStride {
			_ddeg = bm1.Data[_accb]
			_baag = bm2.Data[_egdb] >> uint(_bddc)
			_bgbf = _ddeg & _baag
			_abb += tab[_bgbf]
			for _gfbf = 1; _gfbf < _cbfd; _gfbf++ {
				_ddeg = bm1.Data[_accb+_gfbf]
				_baag = bm2.Data[_egdb+_gfbf] >> uint(_bddc)
				_baag |= bm2.Data[_egdb+_gfbf-1] << uint(8-_bddc)
				_bgbf = _ddeg & _baag
				_abb += tab[_bgbf]
			}
			if _abb >= _fffe {
				return true, nil
			} else if _abb+downcount[_aeag]-_dgac < _fffe {
				return false, nil
			}
		}
	case _cbfd < _cdba:
		for _aeag = _fcceb; _aeag < _egcf; _aeag, _accb, _egdb = _aeag+1, _accb+bm1.RowStride, _egdb+bm2.RowStride {
			for _gfbf = 0; _gfbf < _cbfd; _gfbf++ {
				_ddeg = bm1.Data[_accb+_gfbf]
				_baag = bm2.Data[_egdb+_gfbf] << uint(-_bddc)
				_baag |= bm2.Data[_egdb+_gfbf+1] >> uint(8+_bddc)
				_bgbf = _ddeg & _baag
				_abb += tab[_bgbf]
			}
			if _abb >= _fffe {
				return true, nil
			} else if _faf := _abb + downcount[_aeag] - _dgac; _faf < _fffe {
				return false, nil
			}
		}
	case _cdba >= _cbfd:
		for _aeag = _fcceb; _aeag < _egcf; _aeag, _accb, _egdb = _aeag+1, _accb+bm1.RowStride, _egdb+bm2.RowStride {
			for _gfbf = 0; _gfbf < _cbfd; _gfbf++ {
				_ddeg = bm1.Data[_accb+_gfbf]
				_baag = bm2.Data[_egdb+_gfbf] << uint(-_bddc)
				_baag |= bm2.Data[_egdb+_gfbf+1] >> uint(8+_bddc)
				_bgbf = _ddeg & _baag
				_abb += tab[_bgbf]
			}
			_ddeg = bm1.Data[_accb+_gfbf]
			_baag = bm2.Data[_egdb+_gfbf] << uint(-_bddc)
			_bgbf = _ddeg & _baag
			_abb += tab[_bgbf]
			if _abb >= _fffe {
				return true, nil
			} else if _abb+downcount[_aeag]-_dgac < _fffe {
				return false, nil
			}
		}
	}
	_fafd := float32(_abb) * float32(_abb) / (float32(area1) * float32(area2))
	if _fafd >= scoreThreshold {
		_ea.Log.Trace("\u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064\u0020\u003c\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0020\u0025\u0064\u0020\u0062\u0075\u0074\u0020\u0073c\u006f\u0072\u0065\u0020\u0025\u0066\u0020\u003e\u003d\u0020\u0073\u0063\u006fr\u0065\u0054\u0068\u0072\u0065\u0073h\u006f\u006c\u0064 \u0025\u0066", _abb, _fffe, _fafd, scoreThreshold)
	}
	return false, nil
}
func (_bfb *Bitmap) GetChocolateData() []byte {
	if _bfb.Color == Vanilla {
		_bfb.inverseData()
	}
	return _bfb.Data
}
func _eeed(_aecde, _eabg *Bitmap, _dbgc *Selection) (*Bitmap, error) {
	const _efde = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0032"
	var _fded, _cdgc int
	if _eabg == nil {
		return nil, _d.Error(_efde, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _dbgc == nil {
		return nil, _d.Error(_efde, "\u0073e\u006c \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_fded = _dbgc.Width
	_cdgc = _dbgc.Height
	if _fded == 0 || _cdgc == 0 {
		return nil, _d.Error(_efde, "\u0073\u0065\u006c\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _aecde == nil {
		return _eabg.createTemplate(), nil
	}
	if _cabce := _aecde.resizeImageData(_eabg); _cabce != nil {
		return nil, _cabce
	}
	return _aecde, nil
}
func (_dfec *ClassedPoints) Swap(i, j int) {
	_dfec.IntSlice[i], _dfec.IntSlice[j] = _dfec.IntSlice[j], _dfec.IntSlice[i]
}
func (_ccfd *Boxes) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_fcf *Boxes, _ggdc error) {
	const _ceda = "\u0042o\u0078e\u0073\u002e\u0053\u0065\u006ce\u0063\u0074B\u0079\u0053\u0069\u007a\u0065"
	if _ccfd == nil {
		return nil, _d.Error(_ceda, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_ccfd) == 0 {
		return _ccfd, nil
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _d.Errorf(_ceda, "\u0069\u006e\u0076al\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE:
	default:
		return nil, _d.Errorf(_ceda, "i\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020t\u0079\u0070\u0065:\u0020'\u0025\u0064\u0027", tp)
	}
	_bdf := _ccfd.makeSizeIndicator(width, height, tp, relation)
	_eede, _ggdc := _ccfd.selectWithIndicator(_bdf)
	if _ggdc != nil {
		return nil, _d.Wrap(_ggdc, _ceda, "")
	}
	return _eede, nil
}

type BoundaryCondition int

func _gdfg(_daegf, _ddac *Bitmap, _dbga *Selection) (*Bitmap, error) {
	const _ddbf = "c\u006c\u006f\u0073\u0065\u0042\u0069\u0074\u006d\u0061\u0070"
	var _cabc error
	if _daegf, _cabc = _eeed(_daegf, _ddac, _dbga); _cabc != nil {
		return nil, _cabc
	}
	_aeg, _cabc := _aef(nil, _ddac, _dbga)
	if _cabc != nil {
		return nil, _d.Wrap(_cabc, _ddbf, "")
	}
	if _, _cabc = _agbc(_daegf, _aeg, _dbga); _cabc != nil {
		return nil, _d.Wrap(_cabc, _ddbf, "")
	}
	return _daegf, nil
}
func _dcag(_fdfe, _fdc *Bitmap, _dffb, _bcfe int) (*Bitmap, error) {
	const _deb = "\u0063\u006c\u006f\u0073\u0065\u0053\u0061\u0066\u0065B\u0072\u0069\u0063\u006b"
	if _fdc == nil {
		return nil, _d.Error(_deb, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _dffb < 1 || _bcfe < 1 {
		return nil, _d.Error(_deb, "\u0068s\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _dffb == 1 && _bcfe == 1 {
		return _bce(_fdfe, _fdc)
	}
	if MorphBC == SymmetricMorphBC {
		_bdcb, _cdcfd := _gagbb(_fdfe, _fdc, _dffb, _bcfe)
		if _cdcfd != nil {
			return nil, _d.Wrap(_cdcfd, _deb, "\u0053\u0079m\u006d\u0065\u0074r\u0069\u0063\u004d\u006f\u0072\u0070\u0068\u0042\u0043")
		}
		return _bdcb, nil
	}
	_gfag := _eec(_dffb/2, _bcfe/2)
	_cbea := 8 * ((_gfag + 7) / 8)
	_ebeec, _babf := _fdc.AddBorder(_cbea, 0)
	if _babf != nil {
		return nil, _d.Wrapf(_babf, _deb, "\u0042\u006f\u0072\u0064\u0065\u0072\u0053\u0069\u007ae\u003a\u0020\u0025\u0064", _cbea)
	}
	var _gfcf, _daebe *Bitmap
	if _dffb == 1 || _bcfe == 1 {
		_dedfa := SelCreateBrick(_bcfe, _dffb, _bcfe/2, _dffb/2, SelHit)
		_gfcf, _babf = _gdfg(nil, _ebeec, _dedfa)
		if _babf != nil {
			return nil, _d.Wrap(_babf, _deb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
	} else {
		_gddb := SelCreateBrick(1, _dffb, 0, _dffb/2, SelHit)
		_fbad, _cggbc := _aef(nil, _ebeec, _gddb)
		if _cggbc != nil {
			return nil, _d.Wrap(_cggbc, _deb, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0064\u0069\u006c\u0061t\u0065")
		}
		_bdfb := SelCreateBrick(_bcfe, 1, _bcfe/2, 0, SelHit)
		_gfcf, _cggbc = _aef(nil, _fbad, _bdfb)
		if _cggbc != nil {
			return nil, _d.Wrap(_cggbc, _deb, "\u0072\u0065\u0067ul\u0061\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
		}
		if _, _cggbc = _agbc(_fbad, _gfcf, _gddb); _cggbc != nil {
			return nil, _d.Wrap(_cggbc, _deb, "r\u0065\u0067\u0075\u006car\u0020-\u0020\u0066\u0069\u0072\u0073t\u0020\u0065\u0072\u006f\u0064\u0065")
		}
		if _, _cggbc = _agbc(_gfcf, _fbad, _bdfb); _cggbc != nil {
			return nil, _d.Wrap(_cggbc, _deb, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0065\u0072\u006fd\u0065")
		}
	}
	if _daebe, _babf = _gfcf.RemoveBorder(_cbea); _babf != nil {
		return nil, _d.Wrap(_babf, _deb, "\u0072e\u0067\u0075\u006c\u0061\u0072")
	}
	if _fdfe == nil {
		return _daebe, nil
	}
	if _, _babf = _bce(_fdfe, _daebe); _babf != nil {
		return nil, _babf
	}
	return _fdfe, nil
}
func (_gcbdd *Bitmaps) SortByWidth() { _cadfc := (*byWidth)(_gcbdd); _a.Sort(_cadfc) }

type SizeSelection int

func _cb(_ae, _fa *Bitmap) (_be error) {
	const _edb = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0032"
	_dac := _fa.RowStride
	_dg := _ae.RowStride
	var (
		_fac                    byte
		_fg                     uint16
		_ef, _db, _cf, _cc, _ca int
	)
	for _cf = 0; _cf < _fa.Height; _cf++ {
		_ef = _cf * _dac
		_db = 2 * _cf * _dg
		for _cc = 0; _cc < _dac; _cc++ {
			_fac = _fa.Data[_ef+_cc]
			_fg = _bdaca[_fac]
			_ca = _db + _cc*2
			if _ae.RowStride != _fa.RowStride*2 && (_cc+1)*2 > _ae.RowStride {
				_be = _ae.SetByte(_ca, byte(_fg>>8))
			} else {
				_be = _ae.setTwoBytes(_ca, _fg)
			}
			if _be != nil {
				return _d.Wrap(_be, _edb, "")
			}
		}
		for _cc = 0; _cc < _dg; _cc++ {
			_ca = _db + _dg + _cc
			_fac = _ae.Data[_db+_cc]
			if _be = _ae.SetByte(_ca, _fac); _be != nil {
				return _d.Wrapf(_be, _edb, "c\u006f\u0070\u0079\u0020\u0064\u006fu\u0062\u006c\u0065\u0064\u0020\u006ci\u006e\u0065\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0042\u0079\u0074\u0065\u003a\u0020\u0027\u0025d\u0027", _db+_cc, _db+_dg+_cc)
			}
		}
	}
	return nil
}
func (_aaee *Points) Add(pt *Points) error {
	const _bdcbg = "\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0041\u0064\u0064"
	if _aaee == nil {
		return _d.Error(_bdcbg, "\u0070o\u0069n\u0074\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if pt == nil {
		return _d.Error(_bdcbg, "a\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	*_aaee = append(*_aaee, *pt...)
	return nil
}
func TstImageBitmap() *Bitmap { return _edcgg.Copy() }

type ClassedPoints struct {
	*Points
	_b.IntSlice
	_afafe func(_abgg, _fafe int) bool
}

func (_eac *Bitmap) Equals(s *Bitmap) bool {
	if len(_eac.Data) != len(s.Data) || _eac.Width != s.Width || _eac.Height != s.Height {
		return false
	}
	for _dbc := 0; _dbc < _eac.Height; _dbc++ {
		_dbg := _dbc * _eac.RowStride
		for _fadb := 0; _fadb < _eac.RowStride; _fadb++ {
			if _eac.Data[_dbg+_fadb] != s.Data[_dbg+_fadb] {
				return false
			}
		}
	}
	return true
}

type MorphProcess struct {
	Operation MorphOperation
	Arguments []int
}

func (_ccefd *ClassedPoints) validateIntSlice() error {
	const _gbae = "\u0076\u0061l\u0069\u0064\u0061t\u0065\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065"
	for _, _fbba := range _ccefd.IntSlice {
		if _fbba >= (_ccefd.Points.Size()) {
			return _d.Errorf(_gbae, "c\u006c\u0061\u0073\u0073\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006ci\u0064 \u0069\u006e\u0064\u0065x\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0064", _fbba, _ccefd.Points.Size())
		}
	}
	return nil
}
func TstImageBitmapInverseData() []byte {
	_gcbeb := _edcgg.Copy()
	_gcbeb.InverseData()
	return _gcbeb.Data
}
func (_feea Points) Size() int { return len(_feea) }
func (_bbg *Bitmap) setAll() error {
	_bab := _fdfa(_bbg, 0, 0, _bbg.Width, _bbg.Height, PixSet, nil, 0, 0)
	if _bab != nil {
		return _d.Wrap(_bab, "\u0073\u0065\u0074\u0041\u006c\u006c", "")
	}
	return nil
}
func (_fdd *Bitmap) And(s *Bitmap) (_faa *Bitmap, _gcb error) {
	const _baa = "\u0042\u0069\u0074\u006d\u0061\u0070\u002e\u0041\u006e\u0064"
	if _fdd == nil {
		return nil, _d.Error(_baa, "\u0027b\u0069t\u006d\u0061\u0070\u0020\u0027b\u0027\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	if s == nil {
		return nil, _d.Error(_baa, "\u0062\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if !_fdd.SizesEqual(s) {
		_ea.Log.Debug("\u0025\u0073\u0020-\u0020\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0069\u007a\u0065 \u0077\u0069\u0074\u0068\u0020\u0027\u0062\u0027", _baa)
	}
	if _faa, _gcb = _bce(_faa, _fdd); _gcb != nil {
		return nil, _d.Wrap(_gcb, _baa, "\u0063\u0061\u006e't\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _gcb = _faa.RasterOperation(0, 0, _faa.Width, _faa.Height, PixSrcAndDst, s, 0, 0); _gcb != nil {
		return nil, _d.Wrap(_gcb, _baa, "")
	}
	return _faa, nil
}
func TstImageBitmapData() []byte             { return _edcgg.Data }
func (_cbda *Bitmap) GetBitOffset(x int) int { return x & 0x07 }
func (_afe *Bitmap) ClipRectangle(box *_fb.Rectangle) (_face *Bitmap, _acd *_fb.Rectangle, _fdg error) {
	const _deg = "\u0043\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, nil, _d.Error(_deg, "\u0062o\u0078 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	_dfbf, _acb := _afe.Width, _afe.Height
	_ddf := _fb.Rect(0, 0, _dfbf, _acb)
	if !box.Overlaps(_ddf) {
		return nil, nil, _d.Error(_deg, "b\u006f\u0078\u0020\u0064oe\u0073n\u0027\u0074\u0020\u006f\u0076e\u0072\u006c\u0061\u0070\u0020\u0062")
	}
	_ebbb := box.Intersect(_ddf)
	_cfg, _badf := _ebbb.Min.X, _ebbb.Min.Y
	_eaa, _cgfd := _ebbb.Dx(), _ebbb.Dy()
	_face = New(_eaa, _cgfd)
	_face.Text = _afe.Text
	if _fdg = _face.RasterOperation(0, 0, _eaa, _cgfd, PixSrc, _afe, _cfg, _badf); _fdg != nil {
		return nil, nil, _d.Wrap(_fdg, _deg, "\u0050\u0069\u0078\u0053\u0072\u0063\u0020\u0074\u006f\u0020\u0063\u006ci\u0070\u0070\u0065\u0064")
	}
	_acd = &_ebbb
	return _face, _acd, nil
}
func (_cdegg *ClassedPoints) XAtIndex(i int) float32 { return (*_cdegg.Points)[_cdegg.IntSlice[i]].X }
func _fgf(_ebgd, _fabg *Bitmap, _bfgd, _dfcd int) (*Bitmap, error) {
	const _dbgge = "\u006fp\u0065\u006e\u0042\u0072\u0069\u0063k"
	if _fabg == nil {
		return nil, _d.Error(_dbgge, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _bfgd < 1 && _dfcd < 1 {
		return nil, _d.Error(_dbgge, "\u0068\u0053\u0069\u007ae \u003c\u0020\u0031\u0020\u0026\u0026\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u003c \u0031")
	}
	if _bfgd == 1 && _dfcd == 1 {
		return _fabg.Copy(), nil
	}
	if _bfgd == 1 || _dfcd == 1 {
		var _dafff error
		_dgdcg := SelCreateBrick(_dfcd, _bfgd, _dfcd/2, _bfgd/2, SelHit)
		_ebgd, _dafff = _gdfe(_ebgd, _fabg, _dgdcg)
		if _dafff != nil {
			return nil, _d.Wrap(_dafff, _dbgge, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _ebgd, nil
	}
	_bdaag := SelCreateBrick(1, _bfgd, 0, _bfgd/2, SelHit)
	_ddcg := SelCreateBrick(_dfcd, 1, _dfcd/2, 0, SelHit)
	_eaba, _edfccb := _agbc(nil, _fabg, _bdaag)
	if _edfccb != nil {
		return nil, _d.Wrap(_edfccb, _dbgge, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_ebgd, _edfccb = _agbc(_ebgd, _eaba, _ddcg)
	if _edfccb != nil {
		return nil, _d.Wrap(_edfccb, _dbgge, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	_, _edfccb = _aef(_eaba, _ebgd, _bdaag)
	if _edfccb != nil {
		return nil, _d.Wrap(_edfccb, _dbgge, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_, _edfccb = _aef(_ebgd, _eaba, _ddcg)
	if _edfccb != nil {
		return nil, _d.Wrap(_edfccb, _dbgge, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _ebgd, nil
}
func (_efa *Bitmap) ThresholdPixelSum(thresh int, tab8 []int) (_cbde bool, _aea error) {
	const _adgg = "\u0042i\u0074\u006d\u0061\u0070\u002e\u0054\u0068\u0072\u0065\u0073\u0068o\u006c\u0064\u0050\u0069\u0078\u0065\u006c\u0053\u0075\u006d"
	if tab8 == nil {
		tab8 = _gebgfd()
	}
	_ggb := _efa.Width >> 3
	_faea := _efa.Width & 7
	_cgd := byte(0xff << uint(8-_faea))
	var (
		_dda, _gfa, _bggc, _afac int
		_dbgd                    byte
	)
	for _dda = 0; _dda < _efa.Height; _dda++ {
		_bggc = _efa.RowStride * _dda
		for _gfa = 0; _gfa < _ggb; _gfa++ {
			_dbgd, _aea = _efa.GetByte(_bggc + _gfa)
			if _aea != nil {
				return false, _d.Wrap(_aea, _adgg, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
			}
			_afac += tab8[_dbgd]
		}
		if _faea != 0 {
			_dbgd, _aea = _efa.GetByte(_bggc + _gfa)
			if _aea != nil {
				return false, _d.Wrap(_aea, _adgg, "p\u0061\u0072\u0074\u0069\u0061\u006c\u0042\u0079\u0074\u0065")
			}
			_dbgd &= _cgd
			_afac += tab8[_dbgd]
		}
		if _afac > thresh {
			return true, nil
		}
	}
	return _cbde, nil
}
func (_fgbg *Bitmaps) ClipToBitmap(s *Bitmap) (*Bitmaps, error) {
	const _bfcf = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0043\u006c\u0069p\u0054\u006f\u0042\u0069tm\u0061\u0070"
	if _fgbg == nil {
		return nil, _d.Error(_bfcf, "\u0042\u0069\u0074\u006dap\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if s == nil {
		return nil, _d.Error(_bfcf, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	_dfgd := len(_fgbg.Values)
	_aeab := &Bitmaps{Values: make([]*Bitmap, _dfgd), Boxes: make([]*_fb.Rectangle, _dfgd)}
	var (
		_bddcg, _aafc *Bitmap
		_eeda         *_fb.Rectangle
		_eebbd        error
	)
	for _dgcac := 0; _dgcac < _dfgd; _dgcac++ {
		if _bddcg, _eebbd = _fgbg.GetBitmap(_dgcac); _eebbd != nil {
			return nil, _d.Wrap(_eebbd, _bfcf, "")
		}
		if _eeda, _eebbd = _fgbg.GetBox(_dgcac); _eebbd != nil {
			return nil, _d.Wrap(_eebbd, _bfcf, "")
		}
		if _aafc, _eebbd = s.clipRectangle(_eeda, nil); _eebbd != nil {
			return nil, _d.Wrap(_eebbd, _bfcf, "")
		}
		if _aafc, _eebbd = _aafc.And(_bddcg); _eebbd != nil {
			return nil, _d.Wrap(_eebbd, _bfcf, "")
		}
		_aeab.Values[_dgcac] = _aafc
		_aeab.Boxes[_dgcac] = _eeda
	}
	return _aeab, nil
}
func (_cbfg *Bitmap) Copy() *Bitmap {
	_bgd := make([]byte, len(_cbfg.Data))
	copy(_bgd, _cbfg.Data)
	return &Bitmap{Width: _cbfg.Width, Height: _cbfg.Height, RowStride: _cbfg.RowStride, Data: _bgd, Color: _cbfg.Color, Text: _cbfg.Text, BitmapNumber: _cbfg.BitmapNumber, Special: _cbfg.Special}
}
func _eagc(_cccc, _ffa *Bitmap, _efgeb, _aebc, _gefc uint, _cbfe, _eebd int, _eee bool, _abf, _cfgb int) error {
	for _dgdee := _cbfe; _dgdee < _eebd; _dgdee++ {
		if _abf+1 < len(_cccc.Data) {
			_bdaa := _dgdee+1 == _eebd
			_dbaa, _eeab := _cccc.GetByte(_abf)
			if _eeab != nil {
				return _eeab
			}
			_abf++
			_dbaa <<= _efgeb
			_bccde, _eeab := _cccc.GetByte(_abf)
			if _eeab != nil {
				return _eeab
			}
			_bccde >>= _aebc
			_gcg := _dbaa | _bccde
			if _bdaa && !_eee {
				_gcg = _gcc(_gefc, _gcg)
			}
			_eeab = _ffa.SetByte(_cfgb, _gcg)
			if _eeab != nil {
				return _eeab
			}
			_cfgb++
			if _bdaa && _eee {
				_eda, _addc := _cccc.GetByte(_abf)
				if _addc != nil {
					return _addc
				}
				_eda <<= _efgeb
				_gcg = _gcc(_gefc, _eda)
				if _addc = _ffa.SetByte(_cfgb, _gcg); _addc != nil {
					return _addc
				}
			}
			continue
		}
		_ebbe, _afgea := _cccc.GetByte(_abf)
		if _afgea != nil {
			_ea.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0061\u0074\u003a\u0020%\u0064\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0073", _abf, _afgea)
			return _afgea
		}
		_ebbe <<= _efgeb
		_abf++
		_afgea = _ffa.SetByte(_cfgb, _ebbe)
		if _afgea != nil {
			return _afgea
		}
		_cfgb++
	}
	return nil
}
func (_cce *Bitmap) GetVanillaData() []byte {
	if _cce.Color == Chocolate {
		_cce.inverseData()
	}
	return _cce.Data
}
func (_fdccc *Bitmaps) GetBitmap(i int) (*Bitmap, error) {
	const _afbc = "\u0047e\u0074\u0042\u0069\u0074\u006d\u0061p"
	if _fdccc == nil {
		return nil, _d.Error(_afbc, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if i > len(_fdccc.Values)-1 {
		return nil, _d.Errorf(_afbc, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fdccc.Values[i], nil
}

type Getter interface {
	GetBitmap() *Bitmap
}

func MakePixelCentroidTab8() []int { return _dbag() }
func (_efgg *Bitmap) nextOnPixel(_fadd, _gecf int) (_edfc _fb.Point, _bade bool, _becd error) {
	const _cfeb = "n\u0065\u0078\u0074\u004f\u006e\u0050\u0069\u0078\u0065\u006c"
	_edfc, _bade, _becd = _efgg.nextOnPixelLow(_efgg.Width, _efgg.Height, _efgg.RowStride, _fadd, _gecf)
	if _becd != nil {
		return _edfc, false, _d.Wrap(_becd, _cfeb, "")
	}
	return _edfc, _bade, nil
}
func NewClassedPoints(points *Points, classes _b.IntSlice) (*ClassedPoints, error) {
	const _deab = "\u004e\u0065w\u0043\u006c\u0061s\u0073\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073"
	if points == nil {
		return nil, _d.Error(_deab, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if classes == nil {
		return nil, _d.Error(_deab, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0063\u006c\u0061ss\u0065\u0073")
	}
	_gcga := &ClassedPoints{Points: points, IntSlice: classes}
	if _caa := _gcga.validateIntSlice(); _caa != nil {
		return nil, _d.Wrap(_caa, _deab, "")
	}
	return _gcga, nil
}
func (_aacda *Boxes) Get(i int) (*_fb.Rectangle, error) {
	const _eae = "\u0042o\u0078\u0065\u0073\u002e\u0047\u0065t"
	if _aacda == nil {
		return nil, _d.Error(_eae, "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if i > len(*_aacda)-1 {
		return nil, _d.Errorf(_eae, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return (*_aacda)[i], nil
}
func (_dgeb *ClassedPoints) ySortFunction() func(_accbf int, _dfea int) bool {
	return func(_bbeg, _cgfb int) bool { return _dgeb.YAtIndex(_bbeg) < _dgeb.YAtIndex(_cgfb) }
}
func _acbd(_ffdfc *Bitmap, _fafg *_b.Stack, _acdb, _bbab int) (_dbbg *_fb.Rectangle, _edfa error) {
	const _cfac = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _ffdfc == nil {
		return nil, _d.Error(_cfac, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _fafg == nil {
		return nil, _d.Error(_cfac, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_gddef, _dcee := _ffdfc.Width, _ffdfc.Height
	_acab := _gddef - 1
	_gccc := _dcee - 1
	if _acdb < 0 || _acdb > _acab || _bbab < 0 || _bbab > _gccc || !_ffdfc.GetPixel(_acdb, _bbab) {
		return nil, nil
	}
	_dgfe := _fb.Rect(100000, 100000, 0, 0)
	if _edfa = _geab(_fafg, _acdb, _acdb, _bbab, 1, _gccc, &_dgfe); _edfa != nil {
		return nil, _d.Wrap(_edfa, _cfac, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _edfa = _geab(_fafg, _acdb, _acdb, _bbab+1, -1, _gccc, &_dgfe); _edfa != nil {
		return nil, _d.Wrap(_edfa, _cfac, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_dgfe.Min.X, _dgfe.Max.X = _acdb, _acdb
	_dgfe.Min.Y, _dgfe.Max.Y = _bbab, _bbab
	var (
		_acf   *fillSegment
		_aeedg int
	)
	for _fafg.Len() > 0 {
		if _acf, _edfa = _feef(_fafg); _edfa != nil {
			return nil, _d.Wrap(_edfa, _cfac, "")
		}
		_bbab = _acf._aegc
		for _acdb = _acf._gebc - 1; _acdb >= 0 && _ffdfc.GetPixel(_acdb, _bbab); _acdb-- {
			if _edfa = _ffdfc.SetPixel(_acdb, _bbab, 0); _edfa != nil {
				return nil, _d.Wrap(_edfa, _cfac, "\u0031s\u0074\u0020\u0073\u0065\u0074")
			}
		}
		if _acdb >= _acf._gebc-1 {
			for {
				for _acdb++; _acdb <= _acf._fdeg+1 && _acdb <= _acab && !_ffdfc.GetPixel(_acdb, _bbab); _acdb++ {
				}
				_aeedg = _acdb
				if !(_acdb <= _acf._fdeg+1 && _acdb <= _acab) {
					break
				}
				for ; _acdb <= _acab && _ffdfc.GetPixel(_acdb, _bbab); _acdb++ {
					if _edfa = _ffdfc.SetPixel(_acdb, _bbab, 0); _edfa != nil {
						return nil, _d.Wrap(_edfa, _cfac, "\u0032n\u0064\u0020\u0073\u0065\u0074")
					}
				}
				if _edfa = _geab(_fafg, _aeedg, _acdb-1, _acf._aegc, _acf._bbcgb, _gccc, &_dgfe); _edfa != nil {
					return nil, _d.Wrap(_edfa, _cfac, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
				}
				if _acdb > _acf._fdeg {
					if _edfa = _geab(_fafg, _acf._fdeg+1, _acdb-1, _acf._aegc, -_acf._bbcgb, _gccc, &_dgfe); _edfa != nil {
						return nil, _d.Wrap(_edfa, _cfac, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
					}
				}
			}
			continue
		}
		_aeedg = _acdb + 1
		if _aeedg < _acf._gebc {
			if _edfa = _geab(_fafg, _aeedg, _acf._gebc-1, _acf._aegc, -_acf._bbcgb, _gccc, &_dgfe); _edfa != nil {
				return nil, _d.Wrap(_edfa, _cfac, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
			}
		}
		_acdb = _acf._gebc
		for {
			for ; _acdb <= _acab && _ffdfc.GetPixel(_acdb, _bbab); _acdb++ {
				if _edfa = _ffdfc.SetPixel(_acdb, _bbab, 0); _edfa != nil {
					return nil, _d.Wrap(_edfa, _cfac, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _edfa = _geab(_fafg, _aeedg, _acdb-1, _acf._aegc, _acf._bbcgb, _gccc, &_dgfe); _edfa != nil {
				return nil, _d.Wrap(_edfa, _cfac, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _acdb > _acf._fdeg {
				if _edfa = _geab(_fafg, _acf._fdeg+1, _acdb-1, _acf._aegc, -_acf._bbcgb, _gccc, &_dgfe); _edfa != nil {
					return nil, _d.Wrap(_edfa, _cfac, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _acdb++; _acdb <= _acf._fdeg+1 && _acdb <= _acab && !_ffdfc.GetPixel(_acdb, _bbab); _acdb++ {
			}
			_aeedg = _acdb
			if !(_acdb <= _acf._fdeg+1 && _acdb <= _acab) {
				break
			}
		}
	}
	_dgfe.Max.X++
	_dgfe.Max.Y++
	return &_dgfe, nil
}
func _ebf(_dgcb *Bitmap, _bec ...int) (_dgg *Bitmap, _cae error) {
	const _cde = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0043\u0061\u0073\u0063\u0061\u0064\u0065"
	if _dgcb == nil {
		return nil, _d.Error(_cde, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_bec) == 0 || len(_bec) > 4 {
		return nil, _d.Error(_cde, "t\u0068\u0065\u0072\u0065\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0074\u0020\u006cea\u0073\u0074\u0020\u006fn\u0065\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020mo\u0073\u0074 \u0034\u0020\u006c\u0065\u0076\u0065\u006c\u0073")
	}
	if _bec[0] <= 0 {
		_ea.Log.Debug("\u006c\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030 \u002d\u0020\u006e\u006f\u0020\u0072\u0065\u0064\u0075\u0063t\u0069\u006f\u006e")
		_dgg, _cae = _bce(nil, _dgcb)
		if _cae != nil {
			return nil, _d.Wrap(_cae, _cde, "l\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030")
		}
		return _dgg, nil
	}
	_dgce := _gdgb()
	_dgg = _dgcb
	for _fgd, _daf := range _bec {
		if _daf <= 0 {
			break
		}
		_dgg, _cae = _bbd(_dgg, _daf, _dgce)
		if _cae != nil {
			return nil, _d.Wrapf(_cae, _cde, "\u006c\u0065\u0076\u0065\u006c\u0025\u0064\u0020\u0072\u0065\u0064\u0075c\u0074\u0069\u006f\u006e", _fgd)
		}
	}
	return _dgg, nil
}
func TstISymbol(t *_da.T, scale ...int) *Bitmap {
	_edea, _ebfag := NewWithData(1, 5, []byte{0x80, 0x80, 0x80, 0x80, 0x80})
	_dc.NoError(t, _ebfag)
	return TstGetScaledSymbol(t, _edea, scale...)
}

type BitmapsArray struct {
	Values []*Bitmaps
	Boxes  []*_fb.Rectangle
}

func (_cbef *Bitmap) createTemplate() *Bitmap {
	return &Bitmap{Width: _cbef.Width, Height: _cbef.Height, RowStride: _cbef.RowStride, Color: _cbef.Color, Text: _cbef.Text, BitmapNumber: _cbef.BitmapNumber, Special: _cbef.Special, Data: make([]byte, len(_cbef.Data))}
}

const (
	ComponentConn Component = iota
	ComponentCharacters
	ComponentWords
)

func Copy(d, s *Bitmap) (*Bitmap, error) { return _bce(d, s) }
func _fged(_ccea, _gaca *Bitmap, _bcca, _ebae int) (_gdad error) {
	const _ebcb = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0034"
	var (
		_adbeg, _ffaf, _befec, _aeeebd                  int
		_abfe, _adcf, _aaeb, _gagc, _gada, _gdcd, _dabe byte
	)
	for _adbeg = 0; _adbeg < _bcca; _adbeg++ {
		_befec = _adbeg * _ccea.RowStride
		_aeeebd = _adbeg * _gaca.RowStride
		for _ffaf = 0; _ffaf < _ebae; _ffaf++ {
			_abfe, _gdad = _ccea.GetByte(_befec + _ffaf)
			if _gdad != nil {
				return _d.Wrap(_gdad, _ebcb, "\u0066i\u0072\u0073\u0074\u0020\u0067\u0065t")
			}
			_adcf, _gdad = _gaca.GetByte(_aeeebd + _ffaf)
			if _gdad != nil {
				return _d.Wrap(_gdad, _ebcb, "\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0067\u0065\u0074")
			}
			if _adbeg > 0 {
				_aaeb, _gdad = _ccea.GetByte(_befec - _ccea.RowStride + _ffaf)
				if _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "\u0069\u0020\u003e \u0030")
				}
				_abfe |= _aaeb
			}
			if _ffaf > 0 {
				_gagc, _gdad = _ccea.GetByte(_befec + _ffaf - 1)
				if _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "\u006a\u0020\u003e \u0030")
				}
				_abfe |= _gagc << 7
			}
			_abfe &= _adcf
			if _abfe == 0 || (^_abfe) == 0 {
				if _gdad = _ccea.SetByte(_befec+_ffaf, _abfe); _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "b\u0074\u0020\u003d\u003d 0\u0020|\u007c\u0020\u0028\u005e\u0062t\u0029\u0020\u003d\u003d\u0020\u0030")
				}
				continue
			}
			for {
				_dabe = _abfe
				_abfe = (_abfe | (_abfe >> 1) | (_abfe << 1)) & _adcf
				if (_abfe ^ _dabe) == 0 {
					if _gdad = _ccea.SetByte(_befec+_ffaf, _abfe); _gdad != nil {
						return _d.Wrap(_gdad, _ebcb, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _adbeg = _bcca - 1; _adbeg >= 0; _adbeg-- {
		_befec = _adbeg * _ccea.RowStride
		_aeeebd = _adbeg * _gaca.RowStride
		for _ffaf = _ebae - 1; _ffaf >= 0; _ffaf-- {
			if _abfe, _gdad = _ccea.GetByte(_befec + _ffaf); _gdad != nil {
				return _d.Wrap(_gdad, _ebcb, "\u0072\u0065\u0076\u0065\u0072\u0073\u0065\u0020\u0066\u0069\u0072\u0073t\u0020\u0067\u0065\u0074")
			}
			if _adcf, _gdad = _gaca.GetByte(_aeeebd + _ffaf); _gdad != nil {
				return _d.Wrap(_gdad, _ebcb, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _adbeg < _bcca-1 {
				if _gada, _gdad = _ccea.GetByte(_befec + _ccea.RowStride + _ffaf); _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "\u0072\u0065v\u0065\u0072\u0073e\u0020\u0069\u0020\u003c\u0020\u0068\u0020\u002d\u0031")
				}
				_abfe |= _gada
			}
			if _ffaf < _ebae-1 {
				if _gdcd, _gdad = _ccea.GetByte(_befec + _ffaf + 1); _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "\u0072\u0065\u0076\u0065rs\u0065\u0020\u006a\u0020\u003c\u0020\u0077\u0070\u006c\u0020\u002d\u0020\u0031")
				}
				_abfe |= _gdcd >> 7
			}
			_abfe &= _adcf
			if _abfe == 0 || (^_abfe) == 0 {
				if _gdad = _ccea.SetByte(_befec+_ffaf, _abfe); _gdad != nil {
					return _d.Wrap(_gdad, _ebcb, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006d\u0061\u0073k\u0065\u0064\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061i\u006c\u0065\u0064")
				}
				continue
			}
			for {
				_dabe = _abfe
				_abfe = (_abfe | (_abfe >> 1) | (_abfe << 1)) & _adcf
				if (_abfe ^ _dabe) == 0 {
					if _gdad = _ccea.SetByte(_befec+_ffaf, _abfe); _gdad != nil {
						return _d.Wrap(_gdad, _ebcb, "\u0072e\u0076\u0065\u0072\u0073e\u0020\u0073\u0065\u0074\u0074i\u006eg\u0020p\u0072\u0065\u0076\u0020\u0062\u0079\u0074e")
					}
					break
				}
			}
		}
	}
	return nil
}

var (
	_bdaca = _cgg()
	_degf  = _dfg()
	_fbcd  = _dea()
)

func _cfd(_dce, _fgc *Bitmap, _bdb int, _egca []byte, _ecc int) (_aebd error) {
	const _gedc = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0033"
	var (
		_fdbc, _adgc, _faed, _cad, _dedc, _eccg, _gedf, _dggf int
		_aba, _dgca, _bfga, _gdb                              uint32
		_ada, _ffc                                            byte
		_gee                                                  uint16
	)
	_gbg := make([]byte, 4)
	_daed := make([]byte, 4)
	for _faed = 0; _faed < _dce.Height-1; _faed, _cad = _faed+2, _cad+1 {
		_fdbc = _faed * _dce.RowStride
		_adgc = _cad * _fgc.RowStride
		for _dedc, _eccg = 0, 0; _dedc < _ecc; _dedc, _eccg = _dedc+4, _eccg+1 {
			for _gedf = 0; _gedf < 4; _gedf++ {
				_dggf = _fdbc + _dedc + _gedf
				if _dggf <= len(_dce.Data)-1 && _dggf < _fdbc+_dce.RowStride {
					_gbg[_gedf] = _dce.Data[_dggf]
				} else {
					_gbg[_gedf] = 0x00
				}
				_dggf = _fdbc + _dce.RowStride + _dedc + _gedf
				if _dggf <= len(_dce.Data)-1 && _dggf < _fdbc+(2*_dce.RowStride) {
					_daed[_gedf] = _dce.Data[_dggf]
				} else {
					_daed[_gedf] = 0x00
				}
			}
			_aba = _df.BigEndian.Uint32(_gbg)
			_dgca = _df.BigEndian.Uint32(_daed)
			_bfga = _aba & _dgca
			_bfga |= _bfga << 1
			_gdb = _aba | _dgca
			_gdb &= _gdb << 1
			_dgca = _bfga & _gdb
			_dgca &= 0xaaaaaaaa
			_aba = _dgca | (_dgca << 7)
			_ada = byte(_aba >> 24)
			_ffc = byte((_aba >> 8) & 0xff)
			_dggf = _adgc + _eccg
			if _dggf+1 == len(_fgc.Data)-1 || _dggf+1 >= _adgc+_fgc.RowStride {
				if _aebd = _fgc.SetByte(_dggf, _egca[_ada]); _aebd != nil {
					return _d.Wrapf(_aebd, _gedc, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _dggf)
				}
			} else {
				_gee = (uint16(_egca[_ada]) << 8) | uint16(_egca[_ffc])
				if _aebd = _fgc.setTwoBytes(_dggf, _gee); _aebd != nil {
					return _d.Wrapf(_aebd, _gedc, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _dggf)
				}
				_eccg++
			}
		}
	}
	return nil
}
func (_acfe *Bitmaps) GroupByWidth() (*BitmapsArray, error) {
	const _ffaa = "\u0047\u0072\u006fu\u0070\u0042\u0079\u0057\u0069\u0064\u0074\u0068"
	if len(_acfe.Values) == 0 {
		return nil, _d.Error(_ffaa, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_cfae := &BitmapsArray{}
	_acfe.SortByWidth()
	_bdcf := -1
	_fbgce := -1
	for _cddc := 0; _cddc < len(_acfe.Values); _cddc++ {
		_dedd := _acfe.Values[_cddc].Width
		if _dedd > _bdcf {
			_bdcf = _dedd
			_fbgce++
			_cfae.Values = append(_cfae.Values, &Bitmaps{})
		}
		_cfae.Values[_fbgce].AddBitmap(_acfe.Values[_cddc])
	}
	return _cfae, nil
}
func HausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH int) (bool, error) {
	const _gaga = "\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_egfa, _baf := p1.Width, p1.Height
	_cfga, _ecdf := p3.Width, p3.Height
	if _b.Abs(_egfa-_cfga) > maxDiffW {
		return false, nil
	}
	if _b.Abs(_baf-_ecdf) > maxDiffH {
		return false, nil
	}
	_aeef := int(delX + _b.Sign(delX)*0.5)
	_cfgab := int(delY + _b.Sign(delY)*0.5)
	var _cfbab error
	_gdbf := p1.CreateTemplate()
	if _cfbab = _gdbf.RasterOperation(0, 0, _egfa, _baf, PixSrc, p1, 0, 0); _cfbab != nil {
		return false, _d.Wrap(_cfbab, _gaga, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _cfbab = _gdbf.RasterOperation(_aeef, _cfgab, _egfa, _baf, PixNotSrcAndDst, p4, 0, 0); _cfbab != nil {
		return false, _d.Wrap(_cfbab, _gaga, "\u0021p\u0034\u0020\u0026\u0020\u0074")
	}
	if _gdbf.Zero() {
		return false, nil
	}
	if _cfbab = _gdbf.RasterOperation(_aeef, _cfgab, _cfga, _ecdf, PixSrc, p3, 0, 0); _cfbab != nil {
		return false, _d.Wrap(_cfbab, _gaga, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _cfbab = _gdbf.RasterOperation(0, 0, _cfga, _ecdf, PixNotSrcAndDst, p2, 0, 0); _cfbab != nil {
		return false, _d.Wrap(_cfbab, _gaga, "\u0021p\u0032\u0020\u0026\u0020\u0074")
	}
	return _gdbf.Zero(), nil
}
func (_fbaa MorphProcess) getWidthHeight() (_addag, _bgbd int) {
	return _fbaa.Arguments[0], _fbaa.Arguments[1]
}
func (_gafd Points) GetGeometry(i int) (_effb, _aged float32, _fbce error) {
	if i > len(_gafd)-1 {
		return 0, 0, _d.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_fdce := _gafd[i]
	return _fdce.X, _fdce.Y, nil
}
func TstRSymbol(t *_da.T, scale ...int) *Bitmap {
	_acbc, _acaea := NewWithData(4, 5, []byte{0xF0, 0x90, 0xF0, 0xA0, 0x90})
	_dc.NoError(t, _acaea)
	return TstGetScaledSymbol(t, _acbc, scale...)
}
func RasterOperation(dest *Bitmap, dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _fdfa(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func TstASymbol(t *_da.T) *Bitmap {
	t.Helper()
	_bdbc := New(6, 6)
	_dc.NoError(t, _bdbc.SetPixel(1, 0, 1))
	_dc.NoError(t, _bdbc.SetPixel(2, 0, 1))
	_dc.NoError(t, _bdbc.SetPixel(3, 0, 1))
	_dc.NoError(t, _bdbc.SetPixel(4, 0, 1))
	_dc.NoError(t, _bdbc.SetPixel(5, 1, 1))
	_dc.NoError(t, _bdbc.SetPixel(1, 2, 1))
	_dc.NoError(t, _bdbc.SetPixel(2, 2, 1))
	_dc.NoError(t, _bdbc.SetPixel(3, 2, 1))
	_dc.NoError(t, _bdbc.SetPixel(4, 2, 1))
	_dc.NoError(t, _bdbc.SetPixel(5, 2, 1))
	_dc.NoError(t, _bdbc.SetPixel(0, 3, 1))
	_dc.NoError(t, _bdbc.SetPixel(5, 3, 1))
	_dc.NoError(t, _bdbc.SetPixel(0, 4, 1))
	_dc.NoError(t, _bdbc.SetPixel(5, 4, 1))
	_dc.NoError(t, _bdbc.SetPixel(1, 5, 1))
	_dc.NoError(t, _bdbc.SetPixel(2, 5, 1))
	_dc.NoError(t, _bdbc.SetPixel(3, 5, 1))
	_dc.NoError(t, _bdbc.SetPixel(4, 5, 1))
	_dc.NoError(t, _bdbc.SetPixel(5, 5, 1))
	return _bdbc
}
func Rect(x, y, w, h int) (*_fb.Rectangle, error) {
	const _eefd = "b\u0069\u0074\u006d\u0061\u0070\u002e\u0052\u0065\u0063\u0074"
	if x < 0 {
		w += x
		x = 0
		if w <= 0 {
			return nil, _d.Errorf(_eefd, "x\u003a\u0027\u0025\u0064\u0027\u0020<\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0077\u003a \u0027\u0025\u0064'\u0020<\u003d\u0020\u0030", x, w)
		}
	}
	if y < 0 {
		h += y
		y = 0
		if h <= 0 {
			return nil, _d.Error(_eefd, "\u0079\u0020\u003c 0\u0020\u0061\u006e\u0064\u0020\u0062\u006f\u0078\u0020\u006f\u0066\u0066\u0020\u002b\u0071\u0075\u0061\u0064")
		}
	}
	_efad := _fb.Rect(x, y, x+w, y+h)
	return &_efad, nil
}
func SelCreateBrick(h, w int, cy, cx int, tp SelectionValue) *Selection {
	_cgac := _dbeae(h, w, "")
	_cgac.setOrigin(cy, cx)
	var _facc, _ebab int
	for _facc = 0; _facc < h; _facc++ {
		for _ebab = 0; _ebab < w; _ebab++ {
			_cgac.Data[_facc][_ebab] = tp
		}
	}
	return _cgac
}

type MorphOperation int

func NewWithUnpaddedData(width, height int, data []byte) (*Bitmap, error) {
	const _edd = "\u004e\u0065\u0077\u0057it\u0068\u0055\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	_ceb := _afge(width, height)
	_ceb.Data = data
	if _gfea := ((width * height) + 7) >> 3; len(data) < _gfea {
		return nil, _d.Errorf(_edd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(data), _gfea)
	}
	if _gcec := _ceb.addPadBits(); _gcec != nil {
		return nil, _d.Wrap(_gcec, _edd, "")
	}
	return _ceb, nil
}
func (_ccdc *Bitmap) nextOnPixelLow(_ggc, _abdg, _eacd, _fffb, _ecd int) (_fba _fb.Point, _cbb bool, _fbdg error) {
	const _geccc = "B\u0069\u0074\u006d\u0061p.\u006ee\u0078\u0074\u004f\u006e\u0050i\u0078\u0065\u006c\u004c\u006f\u0077"
	var (
		_gcbe int
		_ecgf byte
	)
	_fgaa := _ecd * _eacd
	_dgga := _fgaa + (_fffb / 8)
	if _ecgf, _fbdg = _ccdc.GetByte(_dgga); _fbdg != nil {
		return _fba, false, _d.Wrap(_fbdg, _geccc, "\u0078\u0053\u0074\u0061\u0072\u0074\u0020\u0061\u006e\u0064 \u0079\u0053\u0074\u0061\u0072\u0074\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	if _ecgf != 0 {
		_dec := _fffb - (_fffb % 8) + 7
		for _gcbe = _fffb; _gcbe <= _dec && _gcbe < _ggc; _gcbe++ {
			if _ccdc.GetPixel(_gcbe, _ecd) {
				_fba.X = _gcbe
				_fba.Y = _ecd
				return _fba, true, nil
			}
		}
	}
	_bgdaa := (_fffb / 8) + 1
	_gcbe = 8 * _bgdaa
	var _fgag int
	for _dgga = _fgaa + _bgdaa; _gcbe < _ggc; _dgga, _gcbe = _dgga+1, _gcbe+8 {
		if _ecgf, _fbdg = _ccdc.GetByte(_dgga); _fbdg != nil {
			return _fba, false, _d.Wrap(_fbdg, _geccc, "r\u0065\u0073\u0074\u0020of\u0020t\u0068\u0065\u0020\u006c\u0069n\u0065\u0020\u0062\u0079\u0074\u0065")
		}
		if _ecgf == 0 {
			continue
		}
		for _fgag = 0; _fgag < 8 && _gcbe < _ggc; _fgag, _gcbe = _fgag+1, _gcbe+1 {
			if _ccdc.GetPixel(_gcbe, _ecd) {
				_fba.X = _gcbe
				_fba.Y = _ecd
				return _fba, true, nil
			}
		}
	}
	for _fgae := _ecd + 1; _fgae < _abdg; _fgae++ {
		_fgaa = _fgae * _eacd
		for _dgga, _gcbe = _fgaa, 0; _gcbe < _ggc; _dgga, _gcbe = _dgga+1, _gcbe+8 {
			if _ecgf, _fbdg = _ccdc.GetByte(_dgga); _fbdg != nil {
				return _fba, false, _d.Wrap(_fbdg, _geccc, "\u0066o\u006cl\u006f\u0077\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u0073")
			}
			if _ecgf == 0 {
				continue
			}
			for _fgag = 0; _fgag < 8 && _gcbe < _ggc; _fgag, _gcbe = _fgag+1, _gcbe+1 {
				if _ccdc.GetPixel(_gcbe, _fgae) {
					_fba.X = _gcbe
					_fba.Y = _fgae
					return _fba, true, nil
				}
			}
		}
	}
	return _fba, false, nil
}
func _fdfa(_ffgbc *Bitmap, _gaag, _eafe, _eedbb, _bcdfb int, _dcfde RasterOperator, _bccb *Bitmap, _begg, _ebcd int) error {
	const _dfce = "\u0072a\u0073t\u0065\u0072\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e"
	if _ffgbc == nil {
		return _d.Error(_dfce, "\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _dcfde == PixDst {
		return nil
	}
	switch _dcfde {
	case PixClr, PixSet, PixNotDst:
		_ebgg(_ffgbc, _gaag, _eafe, _eedbb, _bcdfb, _dcfde)
		return nil
	}
	if _bccb == nil {
		_ea.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _d.Error(_dfce, "\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _afcb := _aeeg(_ffgbc, _gaag, _eafe, _eedbb, _bcdfb, _dcfde, _bccb, _begg, _ebcd); _afcb != nil {
		return _d.Wrap(_afcb, _dfce, "")
	}
	return nil
}

var MorphBC BoundaryCondition

func (_cadb *Bitmap) thresholdPixelSum(_bdde int) bool {
	var (
		_gecg int
		_eeb  uint8
		_bfbg byte
		_ggf  int
	)
	_afab := _cadb.RowStride
	_bcf := uint(_cadb.Width & 0x07)
	if _bcf != 0 {
		_eeb = uint8((0xff << (8 - _bcf)) & 0xff)
		_afab--
	}
	for _dfff := 0; _dfff < _cadb.Height; _dfff++ {
		for _ggf = 0; _ggf < _afab; _ggf++ {
			_bfbg = _cadb.Data[_dfff*_cadb.RowStride+_ggf]
			_gecg += int(_cdc[_bfbg])
		}
		if _bcf != 0 {
			_bfbg = _cadb.Data[_dfff*_cadb.RowStride+_ggf] & _eeb
			_gecg += int(_cdc[_bfbg])
		}
		if _gecg > _bdde {
			return true
		}
	}
	return false
}
func (_dece *byHeight) Swap(i, j int) {
	_dece.Values[i], _dece.Values[j] = _dece.Values[j], _dece.Values[i]
	if _dece.Boxes != nil {
		_dece.Boxes[i], _dece.Boxes[j] = _dece.Boxes[j], _dece.Boxes[i]
	}
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

func (_cgec *ClassedPoints) GetIntYByClass(i int) (int, error) {
	const _bag = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _cgec.IntSlice.Size() {
		return 0, _d.Errorf(_bag, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_cgec.YAtIndex(i)), nil
}
func NewWithData(width, height int, data []byte) (*Bitmap, error) {
	const _aadc = "N\u0065\u0077\u0057\u0069\u0074\u0068\u0044\u0061\u0074\u0061"
	_dgdc := _afge(width, height)
	_dgdc.Data = data
	if len(data) < height*_dgdc.RowStride {
		return nil, _d.Errorf(_aadc, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u003a\u0020\u0025\u0064", len(data), height*_dgdc.RowStride)
	}
	return _dgdc, nil
}
func _gdfe(_fadgc, _ggbf *Bitmap, _gcgc *Selection) (*Bitmap, error) {
	const _bee = "\u006f\u0070\u0065\u006e"
	var _beee error
	_fadgc, _beee = _eeed(_fadgc, _ggbf, _gcgc)
	if _beee != nil {
		return nil, _d.Wrap(_beee, _bee, "")
	}
	_eba, _beee := _agbc(nil, _ggbf, _gcgc)
	if _beee != nil {
		return nil, _d.Wrap(_beee, _bee, "")
	}
	_, _beee = _aef(_fadgc, _eba, _gcgc)
	if _beee != nil {
		return nil, _d.Wrap(_beee, _bee, "")
	}
	return _fadgc, nil
}

type Component int

func New(width, height int) *Bitmap {
	_gab := _afge(width, height)
	_gab.Data = make([]byte, height*_gab.RowStride)
	return _gab
}
func (_efb *Bitmap) removeBorderGeneral(_fcbd, _facag, _efga, _gae int) (*Bitmap, error) {
	const _ebfcb = "\u0072\u0065\u006d\u006fve\u0042\u006f\u0072\u0064\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _fcbd < 0 || _facag < 0 || _efga < 0 || _gae < 0 {
		return nil, _d.Error(_ebfcb, "\u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020\u0062\u0072\u006f\u0064\u0065\u0072\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u0076\u0061lu\u0065\u0073")
	}
	_eded, _gfb := _efb.Width, _efb.Height
	_ccf := _eded - _fcbd - _facag
	_gca := _gfb - _efga - _gae
	if _ccf <= 0 {
		return nil, _d.Errorf(_ebfcb, "w\u0069\u0064\u0074\u0068: \u0025d\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u0020\u0030", _ccf)
	}
	if _gca <= 0 {
		return nil, _d.Errorf(_ebfcb, "\u0068\u0065\u0069\u0067ht\u003a\u0020\u0025\u0064\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", _gca)
	}
	_ccb := New(_ccf, _gca)
	_ccb.Color = _efb.Color
	_dfef := _ccb.RasterOperation(0, 0, _ccf, _gca, PixSrc, _efb, _fcbd, _efga)
	if _dfef != nil {
		return nil, _d.Wrap(_dfef, _ebfcb, "")
	}
	return _ccb, nil
}

type Bitmap struct {
	Width, Height            int
	BitmapNumber             int
	RowStride                int
	Data                     []byte
	Color                    Color
	Special                  int
	Text                     string
	XResolution, YResolution int
}
type SelectionValue int

func TstWSymbol(t *_da.T, scale ...int) *Bitmap {
	_dcafeg, _bgeb := NewWithData(5, 5, []byte{0x88, 0x88, 0xA8, 0xD8, 0x88})
	_dc.NoError(t, _bgeb)
	return TstGetScaledSymbol(t, _dcafeg, scale...)
}
func (_bacg *Selection) setOrigin(_ebddb, _bggd int) { _bacg.Cy, _bacg.Cx = _ebddb, _bggd }
func TstWriteSymbols(t *_da.T, bms *Bitmaps, src *Bitmap) {
	for _cabg := 0; _cabg < bms.Size(); _cabg++ {
		_abafe := bms.Values[_cabg]
		_becf := bms.Boxes[_cabg]
		_feeb := src.RasterOperation(_becf.Min.X, _becf.Min.Y, _abafe.Width, _abafe.Height, PixSrc, _abafe, 0, 0)
		_dc.NoError(t, _feeb)
	}
}
func _cedg(_dbdg, _dag *Bitmap, _eggc, _fgde, _fca, _gagbd, _aaae, _agec, _eedbd, _dbabd int, _bfdb CombinationOperator) error {
	var _ecb int
	_faedf := func() { _ecb++; _fca += _dag.RowStride; _gagbd += _dbdg.RowStride; _aaae += _dbdg.RowStride }
	for _ecb = _eggc; _ecb < _fgde; _faedf() {
		var _ggba uint16
		_egd := _fca
		for _ccee := _gagbd; _ccee <= _aaae; _ccee++ {
			_bgee, _aede := _dag.GetByte(_egd)
			if _aede != nil {
				return _aede
			}
			_adfd, _aede := _dbdg.GetByte(_ccee)
			if _aede != nil {
				return _aede
			}
			_ggba = (_ggba | uint16(_adfd)) << uint(_dbabd)
			_adfd = byte(_ggba >> 8)
			if _ccee == _aaae {
				_adfd = _gcc(uint(_agec), _adfd)
			}
			if _aede = _dag.SetByte(_egd, _adda(_bgee, _adfd, _bfdb)); _aede != nil {
				return _aede
			}
			_egd++
			_ggba <<= uint(_eedbd)
		}
	}
	return nil
}
func _cbac(_efgcf *Bitmap, _ggfg int) (*Bitmap, error) {
	const _aedea = "\u0065x\u0070a\u006e\u0064\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0065"
	if _efgcf == nil {
		return nil, _d.Error(_aedea, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ggfg <= 0 {
		return nil, _d.Error(_aedea, "i\u006e\u0076\u0061\u006cid\u0020f\u0061\u0063\u0074\u006f\u0072 \u002d\u0020\u003c\u003d\u0020\u0030")
	}
	if _ggfg == 1 {
		_gba, _eeabe := _bce(nil, _efgcf)
		if _eeabe != nil {
			return nil, _d.Wrap(_eeabe, _aedea, "\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u003d\u0020\u0031")
		}
		return _gba, nil
	}
	_agcd, _ecgde := _cgf(_efgcf, _ggfg, _ggfg)
	if _ecgde != nil {
		return nil, _d.Wrap(_ecgde, _aedea, "")
	}
	return _agcd, nil
}
func TstPSymbol(t *_da.T) *Bitmap {
	t.Helper()
	_gadac := New(5, 8)
	_dc.NoError(t, _gadac.SetPixel(0, 0, 1))
	_dc.NoError(t, _gadac.SetPixel(1, 0, 1))
	_dc.NoError(t, _gadac.SetPixel(2, 0, 1))
	_dc.NoError(t, _gadac.SetPixel(3, 0, 1))
	_dc.NoError(t, _gadac.SetPixel(4, 1, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 1, 1))
	_dc.NoError(t, _gadac.SetPixel(4, 2, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 2, 1))
	_dc.NoError(t, _gadac.SetPixel(4, 3, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 3, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 4, 1))
	_dc.NoError(t, _gadac.SetPixel(1, 4, 1))
	_dc.NoError(t, _gadac.SetPixel(2, 4, 1))
	_dc.NoError(t, _gadac.SetPixel(3, 4, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 5, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 6, 1))
	_dc.NoError(t, _gadac.SetPixel(0, 7, 1))
	return _gadac
}
func (_fegc *Points) AddPoint(x, y float32)     { *_fegc = append(*_fegc, Point{x, y}) }
func (_cgge *Bitmap) GetByteIndex(x, y int) int { return y*_cgge.RowStride + (x >> 3) }
func (_bebad *Bitmaps) SortByHeight()           { _facfb := (*byHeight)(_bebad); _a.Sort(_facfb) }
func _edbfd(_afba ...MorphProcess) (_gcgf error) {
	const _cbgg = "v\u0065r\u0069\u0066\u0079\u004d\u006f\u0072\u0070\u0068P\u0072\u006f\u0063\u0065ss\u0065\u0073"
	var _cabdd, _gaead int
	for _acca, _gcfc := range _afba {
		if _gcgf = _gcfc.verify(_acca, &_cabdd, &_gaead); _gcgf != nil {
			return _d.Wrap(_gcgf, _cbgg, "")
		}
	}
	if _gaead != 0 && _cabdd != 0 {
		return _d.Error(_cbgg, "\u004d\u006f\u0072\u0070\u0068\u0020\u0073\u0065\u0071\u0075\u0065n\u0063\u0065\u0020\u002d\u0020\u0062\u006f\u0072d\u0065r\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u0065\u0074\u0020\u0072\u0065\u0064u\u0063\u0074\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0030")
	}
	return nil
}
func (_baddd Points) GetIntX(i int) (int, error) {
	if i >= len(_baddd) {
		return 0, _d.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0058", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_baddd[i].X), nil
}
func _ddab(_dgafc *Bitmap, _aegf, _ddbfd, _ecef, _affb int, _dfdd RasterOperator, _gebb *Bitmap, _bfda, _ffed int) error {
	var (
		_ebcfd        byte
		_caec         int
		_ccgb         int
		_fbda, _cfaab int
		_cfgd, _fdaf  int
	)
	_gbcc := _ecef >> 3
	_fdfee := _ecef & 7
	if _fdfee > 0 {
		_ebcfd = _fada[_fdfee]
	}
	_caec = _gebb.RowStride*_ffed + (_bfda >> 3)
	_ccgb = _dgafc.RowStride*_ddbfd + (_aegf >> 3)
	switch _dfdd {
	case PixSrc:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = _gebb.Data[_fbda]
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda], _ebcfd)
			}
		}
	case PixNotSrc:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = ^(_gebb.Data[_fbda])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^_gebb.Data[_fbda], _ebcfd)
			}
		}
	case PixSrcOrDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] |= _gebb.Data[_fbda]
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda]|_dgafc.Data[_cfaab], _ebcfd)
			}
		}
	case PixSrcAndDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] &= _gebb.Data[_fbda]
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda]&_dgafc.Data[_cfaab], _ebcfd)
			}
		}
	case PixSrcXorDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] ^= _gebb.Data[_fbda]
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda]^_dgafc.Data[_cfaab], _ebcfd)
			}
		}
	case PixNotSrcOrDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] |= ^(_gebb.Data[_fbda])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^(_gebb.Data[_fbda])|_dgafc.Data[_cfaab], _ebcfd)
			}
		}
	case PixNotSrcAndDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] &= ^(_gebb.Data[_fbda])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^(_gebb.Data[_fbda])&_dgafc.Data[_cfaab], _ebcfd)
			}
		}
	case PixSrcOrNotDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = _gebb.Data[_fbda] | ^(_dgafc.Data[_cfaab])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda]|^(_dgafc.Data[_cfaab]), _ebcfd)
			}
		}
	case PixSrcAndNotDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = _gebb.Data[_fbda] &^ (_dgafc.Data[_cfaab])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], _gebb.Data[_fbda]&^(_dgafc.Data[_cfaab]), _ebcfd)
			}
		}
	case PixNotPixSrcOrDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = ^(_gebb.Data[_fbda] | _dgafc.Data[_cfaab])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^(_gebb.Data[_fbda] | _dgafc.Data[_cfaab]), _ebcfd)
			}
		}
	case PixNotPixSrcAndDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = ^(_gebb.Data[_fbda] & _dgafc.Data[_cfaab])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^(_gebb.Data[_fbda] & _dgafc.Data[_cfaab]), _ebcfd)
			}
		}
	case PixNotPixSrcXorDst:
		for _cfgd = 0; _cfgd < _affb; _cfgd++ {
			_fbda = _caec + _cfgd*_gebb.RowStride
			_cfaab = _ccgb + _cfgd*_dgafc.RowStride
			for _fdaf = 0; _fdaf < _gbcc; _fdaf++ {
				_dgafc.Data[_cfaab] = ^(_gebb.Data[_fbda] ^ _dgafc.Data[_cfaab])
				_cfaab++
				_fbda++
			}
			if _fdfee > 0 {
				_dgafc.Data[_cfaab] = _efee(_dgafc.Data[_cfaab], ^(_gebb.Data[_fbda] ^ _dgafc.Data[_cfaab]), _ebcfd)
			}
		}
	default:
		_ea.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _dfdd)
		return _d.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004co\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_faba *Bitmaps) makeSizeIndicator(_abfc, _dcea int, _dfgeg LocationFilter, _gced SizeComparison) (_dacg *_b.NumSlice, _ggbc error) {
	const _fbdd = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u006d\u0061\u006b\u0065S\u0069z\u0065I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _faba == nil {
		return nil, _d.Error(_fbdd, "\u0062\u0069\u0074ma\u0070\u0073\u0020\u0027\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch _dfgeg {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _d.Errorf(_fbdd, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", _dfgeg)
	}
	switch _gced {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _d.Errorf(_fbdd, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", _gced)
	}
	_dacg = &_b.NumSlice{}
	var (
		_ecdag, _acde, _ffbag int
		_dcaba                *Bitmap
	)
	for _, _dcaba = range _faba.Values {
		_ecdag = 0
		_acde, _ffbag = _dcaba.Width, _dcaba.Height
		switch _dfgeg {
		case LocSelectWidth:
			if (_gced == SizeSelectIfLT && _acde < _abfc) || (_gced == SizeSelectIfGT && _acde > _abfc) || (_gced == SizeSelectIfLTE && _acde <= _abfc) || (_gced == SizeSelectIfGTE && _acde >= _abfc) || (_gced == SizeSelectIfEQ && _acde == _abfc) {
				_ecdag = 1
			}
		case LocSelectHeight:
			if (_gced == SizeSelectIfLT && _ffbag < _dcea) || (_gced == SizeSelectIfGT && _ffbag > _dcea) || (_gced == SizeSelectIfLTE && _ffbag <= _dcea) || (_gced == SizeSelectIfGTE && _ffbag >= _dcea) || (_gced == SizeSelectIfEQ && _ffbag == _dcea) {
				_ecdag = 1
			}
		case LocSelectIfEither:
			if (_gced == SizeSelectIfLT && (_acde < _abfc || _ffbag < _dcea)) || (_gced == SizeSelectIfGT && (_acde > _abfc || _ffbag > _dcea)) || (_gced == SizeSelectIfLTE && (_acde <= _abfc || _ffbag <= _dcea)) || (_gced == SizeSelectIfGTE && (_acde >= _abfc || _ffbag >= _dcea)) || (_gced == SizeSelectIfEQ && (_acde == _abfc || _ffbag == _dcea)) {
				_ecdag = 1
			}
		case LocSelectIfBoth:
			if (_gced == SizeSelectIfLT && (_acde < _abfc && _ffbag < _dcea)) || (_gced == SizeSelectIfGT && (_acde > _abfc && _ffbag > _dcea)) || (_gced == SizeSelectIfLTE && (_acde <= _abfc && _ffbag <= _dcea)) || (_gced == SizeSelectIfGTE && (_acde >= _abfc && _ffbag >= _dcea)) || (_gced == SizeSelectIfEQ && (_acde == _abfc && _ffbag == _dcea)) {
				_ecdag = 1
			}
		}
		_dacg.AddInt(_ecdag)
	}
	return _dacg, nil
}
func _cfa(_aeb, _efg *Bitmap) (_g error) {
	const _ec = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0034"
	_gc := _efg.RowStride
	_gf := _aeb.RowStride
	_bc := _efg.RowStride*4 - _aeb.RowStride
	var (
		_fd, _eca                             byte
		_fc                                   uint32
		_cbf, _gg, _cfab, _ga, _ac, _cg, _bcd int
	)
	for _cfab = 0; _cfab < _efg.Height; _cfab++ {
		_cbf = _cfab * _gc
		_gg = 4 * _cfab * _gf
		for _ga = 0; _ga < _gc; _ga++ {
			_fd = _efg.Data[_cbf+_ga]
			_fc = _degf[_fd]
			_cg = _gg + _ga*4
			if _bc != 0 && (_ga+1)*4 > _aeb.RowStride {
				for _ac = _bc; _ac > 0; _ac-- {
					_eca = byte((_fc >> uint(_ac*8)) & 0xff)
					_bcd = _cg + (_bc - _ac)
					if _g = _aeb.SetByte(_bcd, _eca); _g != nil {
						return _d.Wrapf(_g, _ec, "D\u0069\u0066\u0066\u0065\u0072\u0065n\u0074\u0020\u0072\u006f\u0077\u0073\u0074\u0072\u0069d\u0065\u0073\u002e \u004b:\u0020\u0025\u0064", _ac)
					}
				}
			} else if _g = _aeb.setFourBytes(_cg, _fc); _g != nil {
				return _d.Wrap(_g, _ec, "")
			}
			if _g = _aeb.setFourBytes(_gg+_ga*4, _degf[_efg.Data[_cbf+_ga]]); _g != nil {
				return _d.Wrap(_g, _ec, "")
			}
		}
		for _ac = 1; _ac < 4; _ac++ {
			for _ga = 0; _ga < _gf; _ga++ {
				if _g = _aeb.SetByte(_gg+_ac*_gf+_ga, _aeb.Data[_gg+_ga]); _g != nil {
					return _d.Wrapf(_g, _ec, "\u0063\u006f\u0070\u0079\u0020\u0027\u0071\u0075\u0061\u0064\u0072\u0061\u0062l\u0065\u0027\u0020\u006c\u0069\u006ee\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0079\u0074\u0065\u003a \u0027\u0025\u0064\u0027", _ac, _ga)
				}
			}
		}
	}
	return nil
}

var (
	_bfgad *Bitmap
	_edcgg *Bitmap
)

type CombinationOperator int

func _adda(_aeee, _ecda byte, _egfb CombinationOperator) byte {
	switch _egfb {
	case CmbOpOr:
		return _ecda | _aeee
	case CmbOpAnd:
		return _ecda & _aeee
	case CmbOpXor:
		return _ecda ^ _aeee
	case CmbOpXNor:
		return ^(_ecda ^ _aeee)
	case CmbOpNot:
		return ^(_ecda)
	default:
		return _ecda
	}
}
func TstDSymbol(t *_da.T, scale ...int) *Bitmap {
	_cega, _cbdg := NewWithData(4, 5, []byte{0xf0, 0x90, 0x90, 0x90, 0xE0})
	_dc.NoError(t, _cbdg)
	return TstGetScaledSymbol(t, _cega, scale...)
}
func _dfg() (_aaf [256]uint32) {
	for _gec := 0; _gec < 256; _gec++ {
		if _gec&0x01 != 0 {
			_aaf[_gec] |= 0xf
		}
		if _gec&0x02 != 0 {
			_aaf[_gec] |= 0xf0
		}
		if _gec&0x04 != 0 {
			_aaf[_gec] |= 0xf00
		}
		if _gec&0x08 != 0 {
			_aaf[_gec] |= 0xf000
		}
		if _gec&0x10 != 0 {
			_aaf[_gec] |= 0xf0000
		}
		if _gec&0x20 != 0 {
			_aaf[_gec] |= 0xf00000
		}
		if _gec&0x40 != 0 {
			_aaf[_gec] |= 0xf000000
		}
		if _gec&0x80 != 0 {
			_aaf[_gec] |= 0xf0000000
		}
	}
	return _aaf
}
func _age(_cceg, _cadbb *Bitmap, _eegf, _edbd, _gbgf, _ccce, _edeb, _dfa, _bgf, _fbf int, _fcbc CombinationOperator, _gfcc int) error {
	var _eadc int
	_gegf := func() { _eadc++; _gbgf += _cadbb.RowStride; _ccce += _cceg.RowStride; _edeb += _cceg.RowStride }
	for _eadc = _eegf; _eadc < _edbd; _gegf() {
		var _dgb uint16
		_gbcgf := _gbgf
		for _deac := _ccce; _deac <= _edeb; _deac++ {
			_ebddd, _cgfc := _cadbb.GetByte(_gbcgf)
			if _cgfc != nil {
				return _cgfc
			}
			_cfge, _cgfc := _cceg.GetByte(_deac)
			if _cgfc != nil {
				return _cgfc
			}
			_dgb = (_dgb | (uint16(_cfge) & 0xff)) << uint(_fbf)
			_cfge = byte(_dgb >> 8)
			if _cgfc = _cadbb.SetByte(_gbcgf, _adda(_ebddd, _cfge, _fcbc)); _cgfc != nil {
				return _cgfc
			}
			_gbcgf++
			_dgb <<= uint(_bgf)
			if _deac == _edeb {
				_cfge = byte(_dgb >> (8 - uint8(_fbf)))
				if _gfcc != 0 {
					_cfge = _gcc(uint(8+_dfa), _cfge)
				}
				_ebddd, _cgfc = _cadbb.GetByte(_gbcgf)
				if _cgfc != nil {
					return _cgfc
				}
				if _cgfc = _cadbb.SetByte(_gbcgf, _adda(_ebddd, _cfge, _fcbc)); _cgfc != nil {
					return _cgfc
				}
			}
		}
	}
	return nil
}
func (_bbge *Bitmaps) selectByIndexes(_cafc []int) (*Bitmaps, error) {
	_fbaae := &Bitmaps{}
	for _, _fcdg := range _cafc {
		_dfca, _dbgae := _bbge.GetBitmap(_fcdg)
		if _dbgae != nil {
			return nil, _d.Wrap(_dbgae, "\u0073e\u006ce\u0063\u0074\u0042\u0079\u0049\u006e\u0064\u0065\u0078\u0065\u0073", "")
		}
		_fbaae.AddBitmap(_dfca)
	}
	return _fbaae, nil
}
func _ce(_fbb *Bitmap, _ced *Bitmap, _bd int) (_dfe error) {
	const _fbe = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0050\u006fw\u0065\u0072\u0032\u004c\u006f\u0077"
	switch _bd {
	case 2:
		_dfe = _cb(_fbb, _ced)
	case 4:
		_dfe = _cfa(_fbb, _ced)
	case 8:
		_dfe = _edbe(_fbb, _ced)
	default:
		return _d.Error(_fbe, "\u0065\u0078p\u0061\u006e\u0073\u0069o\u006e\u0020f\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020r\u0061\u006e\u0067\u0065")
	}
	if _dfe != nil {
		_dfe = _d.Wrap(_dfe, _fbe, "")
	}
	return _dfe
}
func (_deece *Bitmaps) GetBox(i int) (*_fb.Rectangle, error) {
	const _eabd = "\u0047\u0065\u0074\u0042\u006f\u0078"
	if _deece == nil {
		return nil, _d.Error(_eabd, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074\u006d\u0061\u0070s\u0027")
	}
	if i > len(_deece.Boxes)-1 {
		return nil, _d.Errorf(_eabd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _deece.Boxes[i], nil
}
func _cgf(_ead *Bitmap, _ge, _dgc int) (*Bitmap, error) {
	const _aa = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0052\u0065p\u006c\u0069\u0063\u0061\u0074\u0065"
	if _ead == nil {
		return nil, _d.Error(_aa, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ge <= 0 || _dgc <= 0 {
		return nil, _d.Error(_aa, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _ge == _dgc {
		if _ge == 1 {
			_cgfa, _gbb := _bce(nil, _ead)
			if _gbb != nil {
				return nil, _d.Wrap(_gbb, _aa, "\u0078\u0046\u0061\u0063\u0074\u0020\u003d\u003d\u0020y\u0046\u0061\u0063\u0074")
			}
			return _cgfa, nil
		}
		if _ge == 2 || _ge == 4 || _ge == 8 {
			_dab, _cd := _fcb(_ead, _ge)
			if _cd != nil {
				return nil, _d.Wrap(_cd, _aa, "\u0078\u0046a\u0063\u0074\u0020i\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
			}
			return _dab, nil
		}
	}
	_beb := _ge * _ead.Width
	_dbd := _dgc * _ead.Height
	_efe := New(_beb, _dbd)
	_dca := _efe.RowStride
	var (
		_dae, _dcaf, _gd, _dd, _ecag int
		_dbf                         byte
		_cedc                        error
	)
	for _dcaf = 0; _dcaf < _ead.Height; _dcaf++ {
		_dae = _dgc * _dcaf * _dca
		for _gd = 0; _gd < _ead.Width; _gd++ {
			if _bff := _ead.GetPixel(_gd, _dcaf); _bff {
				_ecag = _ge * _gd
				for _dd = 0; _dd < _ge; _dd++ {
					_efe.setBit(_dae*8 + _ecag + _dd)
				}
			}
		}
		for _dd = 1; _dd < _dgc; _dd++ {
			_faca := _dae + _dd*_dca
			for _cec := 0; _cec < _dca; _cec++ {
				if _dbf, _cedc = _efe.GetByte(_dae + _cec); _cedc != nil {
					return nil, _d.Wrapf(_cedc, _aa, "\u0072\u0065\u0070\u006cic\u0061\u0074\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0027\u0025d\u0027", _dd)
				}
				if _cedc = _efe.SetByte(_faca+_cec, _dbf); _cedc != nil {
					return nil, _d.Wrap(_cedc, _aa, "\u0053\u0065\u0074\u0074in\u0067\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
				}
			}
		}
	}
	return _efe, nil
}

type Points []Point

var (
	_fada = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_eaec = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func _fcb(_dcf *Bitmap, _ebb int) (*Bitmap, error) {
	const _dee = "\u0065x\u0070a\u006e\u0064\u0042\u0069\u006ea\u0072\u0079P\u006f\u0077\u0065\u0072\u0032"
	if _dcf == nil {
		return nil, _d.Error(_dee, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ebb == 1 {
		return _bce(nil, _dcf)
	}
	if _ebb != 2 && _ebb != 4 && _ebb != 8 {
		return nil, _d.Error(_dee, "\u0066\u0061\u0063t\u006f\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069n\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_ebd := _ebb * _dcf.Width
	_ebg := _ebb * _dcf.Height
	_gb := New(_ebd, _ebg)
	var _dad error
	switch _ebb {
	case 2:
		_dad = _cb(_gb, _dcf)
	case 4:
		_dad = _cfa(_gb, _dcf)
	case 8:
		_dad = _edbe(_gb, _dcf)
	}
	if _dad != nil {
		return nil, _d.Wrap(_dad, _dee, "")
	}
	return _gb, nil
}
func TstFrameBitmap() *Bitmap { return _bfgad.Copy() }

const (
	_fdef shift = iota
	_aage
)

func _gagbb(_agee, _cfebb *Bitmap, _adbe, _agf int) (*Bitmap, error) {
	const _bcad = "\u0063\u006c\u006f\u0073\u0065\u0042\u0072\u0069\u0063\u006b"
	if _cfebb == nil {
		return nil, _d.Error(_bcad, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _adbe < 1 || _agf < 1 {
		return nil, _d.Error(_bcad, "\u0068S\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0053\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _adbe == 1 && _agf == 1 {
		return _cfebb.Copy(), nil
	}
	if _adbe == 1 || _agf == 1 {
		_cgee := SelCreateBrick(_agf, _adbe, _agf/2, _adbe/2, SelHit)
		var _eacb error
		_agee, _eacb = _gdfg(_agee, _cfebb, _cgee)
		if _eacb != nil {
			return nil, _d.Wrap(_eacb, _bcad, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _agee, nil
	}
	_cdf := SelCreateBrick(1, _adbe, 0, _adbe/2, SelHit)
	_cgab := SelCreateBrick(_agf, 1, _agf/2, 0, SelHit)
	_cddg, _effd := _aef(nil, _cfebb, _cdf)
	if _effd != nil {
		return nil, _d.Wrap(_effd, _bcad, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _agee, _effd = _aef(_agee, _cddg, _cgab); _effd != nil {
		return nil, _d.Wrap(_effd, _bcad, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _, _effd = _agbc(_cddg, _agee, _cdf); _effd != nil {
		return nil, _d.Wrap(_effd, _bcad, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	if _, _effd = _agbc(_agee, _cddg, _cgab); _effd != nil {
		return nil, _d.Wrap(_effd, _bcad, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _agee, nil
}
func (_abdb *Bitmaps) selectByIndicator(_dafbe *_b.NumSlice) (_gdbgd *Bitmaps, _bbbd error) {
	const _bfde = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u0073\u0065\u006c\u0065c\u0074B\u0079I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _abdb == nil {
		return nil, _d.Error(_bfde, "\u0027\u0062\u0027 b\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _dafbe == nil {
		return nil, _d.Error(_bfde, "'\u006e\u0061\u0027\u0020\u0069\u006ed\u0069\u0063\u0061\u0074\u006f\u0072\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	if len(_abdb.Values) == 0 {
		return _abdb, nil
	}
	if len(*_dafbe) != len(_abdb.Values) {
		return nil, _d.Errorf(_bfde, "\u006ea\u0020\u006ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064,\u0020\u0069\u0073\u0020\u0064\u0069\u0066\u0066\u0065re\u006e\u0074\u0020t\u0068\u0061n\u0020\u0062\u0069\u0074\u006d\u0061p\u0073\u003a \u0025\u0064", len(*_dafbe), len(_abdb.Values))
	}
	var _cabf, _bfdba, _dfeaa int
	for _bfdba = 0; _bfdba < len(*_dafbe); _bfdba++ {
		if _cabf, _bbbd = _dafbe.GetInt(_bfdba); _bbbd != nil {
			return nil, _d.Wrap(_bbbd, _bfde, "f\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		if _cabf == 1 {
			_dfeaa++
		}
	}
	if _dfeaa == len(_abdb.Values) {
		return _abdb, nil
	}
	_gdbgd = &Bitmaps{}
	_febd := len(_abdb.Values) == len(_abdb.Boxes)
	for _bfdba = 0; _bfdba < len(*_dafbe); _bfdba++ {
		if _cabf = int((*_dafbe)[_bfdba]); _cabf == 0 {
			continue
		}
		_gdbgd.Values = append(_gdbgd.Values, _abdb.Values[_bfdba])
		if _febd {
			_gdbgd.Boxes = append(_gdbgd.Boxes, _abdb.Boxes[_bfdba])
		}
	}
	return _gdbgd, nil
}
func (_dcg *Bitmap) SetByte(index int, v byte) error {
	if index > len(_dcg.Data)-1 || index < 0 {
		return _d.Errorf("\u0053e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", index)
	}
	_dcg.Data[index] = v
	return nil
}

type Boxes []*_fb.Rectangle

func ClipBoxToRectangle(box *_fb.Rectangle, wi, hi int) (_aee *_fb.Rectangle, _ege error) {
	const _bggg = "\u0043l\u0069p\u0042\u006f\u0078\u0054\u006fR\u0065\u0063t\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, _d.Error(_bggg, "\u0027\u0062\u006f\u0078\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if box.Min.X >= wi || box.Min.Y >= hi || box.Max.X <= 0 || box.Max.Y <= 0 {
		return nil, _d.Error(_bggg, "\u0027\u0062\u006fx'\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065")
	}
	_bgfd := *box
	_aee = &_bgfd
	if _aee.Min.X < 0 {
		_aee.Max.X += _aee.Min.X
		_aee.Min.X = 0
	}
	if _aee.Min.Y < 0 {
		_aee.Max.Y += _aee.Min.Y
		_aee.Min.Y = 0
	}
	if _aee.Max.X > wi {
		_aee.Max.X = wi
	}
	if _aee.Max.Y > hi {
		_aee.Max.Y = hi
	}
	return _aee, nil
}

type Point struct{ X, Y float32 }

func TstTSymbol(t *_da.T, scale ...int) *Bitmap {
	_ecfcf, _fgbd := NewWithData(5, 5, []byte{0xF8, 0x20, 0x20, 0x20, 0x20})
	_dc.NoError(t, _fgbd)
	return TstGetScaledSymbol(t, _ecfcf, scale...)
}
func (_eaga *Bitmap) Equivalent(s *Bitmap) bool { return _eaga.equivalent(s) }
func (_afce *Bitmap) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _fdfa(_afce, dx, dy, dw, dh, op, src, sx, sy)
}
func _fdf(_bge, _eed *Bitmap, _cbe int, _afa []byte, _gbc int) (_aag error) {
	const _ede = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0034"
	var (
		_bgb, _ccd, _edbb, _aff, _bfd, _dba, _daa, _gfe int
		_ffg, _dafd                                     uint32
		_bda, _ebe                                      byte
		_bac                                            uint16
	)
	_gcea := make([]byte, 4)
	_aad := make([]byte, 4)
	for _edbb = 0; _edbb < _bge.Height-1; _edbb, _aff = _edbb+2, _aff+1 {
		_bgb = _edbb * _bge.RowStride
		_ccd = _aff * _eed.RowStride
		for _bfd, _dba = 0, 0; _bfd < _gbc; _bfd, _dba = _bfd+4, _dba+1 {
			for _daa = 0; _daa < 4; _daa++ {
				_gfe = _bgb + _bfd + _daa
				if _gfe <= len(_bge.Data)-1 && _gfe < _bgb+_bge.RowStride {
					_gcea[_daa] = _bge.Data[_gfe]
				} else {
					_gcea[_daa] = 0x00
				}
				_gfe = _bgb + _bge.RowStride + _bfd + _daa
				if _gfe <= len(_bge.Data)-1 && _gfe < _bgb+(2*_bge.RowStride) {
					_aad[_daa] = _bge.Data[_gfe]
				} else {
					_aad[_daa] = 0x00
				}
			}
			_ffg = _df.BigEndian.Uint32(_gcea)
			_dafd = _df.BigEndian.Uint32(_aad)
			_dafd &= _ffg
			_dafd &= _dafd << 1
			_dafd &= 0xaaaaaaaa
			_ffg = _dafd | (_dafd << 7)
			_bda = byte(_ffg >> 24)
			_ebe = byte((_ffg >> 8) & 0xff)
			_gfe = _ccd + _dba
			if _gfe+1 == len(_eed.Data)-1 || _gfe+1 >= _ccd+_eed.RowStride {
				_eed.Data[_gfe] = _afa[_bda]
				if _aag = _eed.SetByte(_gfe, _afa[_bda]); _aag != nil {
					return _d.Wrapf(_aag, _ede, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _gfe)
				}
			} else {
				_bac = (uint16(_afa[_bda]) << 8) | uint16(_afa[_ebe])
				if _aag = _eed.setTwoBytes(_gfe, _bac); _aag != nil {
					return _d.Wrapf(_aag, _ede, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _gfe)
				}
				_dba++
			}
		}
	}
	return nil
}
func TstOSymbol(t *_da.T, scale ...int) *Bitmap {
	_fdba, _ceff := NewWithData(4, 5, []byte{0xF0, 0x90, 0x90, 0x90, 0xF0})
	_dc.NoError(t, _ceff)
	return TstGetScaledSymbol(t, _fdba, scale...)
}
func (_gcbf CombinationOperator) String() string {
	var _caba string
	switch _gcbf {
	case CmbOpOr:
		_caba = "\u004f\u0052"
	case CmbOpAnd:
		_caba = "\u0041\u004e\u0044"
	case CmbOpXor:
		_caba = "\u0058\u004f\u0052"
	case CmbOpXNor:
		_caba = "\u0058\u004e\u004f\u0052"
	case CmbOpReplace:
		_caba = "\u0052E\u0050\u004c\u0041\u0043\u0045"
	case CmbOpNot:
		_caba = "\u004e\u004f\u0054"
	}
	return _caba
}
func (_agea *BitmapsArray) GetBitmaps(i int) (*Bitmaps, error) {
	const _ggcd = "\u0042\u0069\u0074ma\u0070\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u0069\u0074\u006d\u0061\u0070\u0073"
	if _agea == nil {
		return nil, _d.Error(_ggcd, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_agea.Values)-1 {
		return nil, _d.Errorf(_ggcd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _agea.Values[i], nil
}
func (_gac *Bitmap) clipRectangle(_dgde, _gdgbg *_fb.Rectangle) (_cgag *Bitmap, _dcga error) {
	const _gcfd = "\u0063\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if _dgde == nil {
		return nil, _d.Error(_gcfd, "\u0070r\u006fv\u0069\u0064\u0065\u0064\u0020n\u0069\u006c \u0027\u0062\u006f\u0078\u0027")
	}
	_bbe, _adf := _gac.Width, _gac.Height
	_gbcg, _dcga := ClipBoxToRectangle(_dgde, _bbe, _adf)
	if _dcga != nil {
		_ea.Log.Warning("\u0027\u0062ox\u0027\u0020\u0064o\u0065\u0073\u006e\u0027t o\u0076er\u006c\u0061\u0070\u0020\u0062\u0069\u0074ma\u0070\u0020\u0027\u0062\u0027\u003a\u0020%\u0076", _dcga)
		return nil, nil
	}
	_gaba, _dbe := _gbcg.Min.X, _gbcg.Min.Y
	_bga, _dfd := _gbcg.Max.X-_gbcg.Min.X, _gbcg.Max.Y-_gbcg.Min.Y
	_cgag = New(_bga, _dfd)
	_cgag.Text = _gac.Text
	if _dcga = _cgag.RasterOperation(0, 0, _bga, _dfd, PixSrc, _gac, _gaba, _dbe); _dcga != nil {
		return nil, _d.Wrap(_dcga, _gcfd, "")
	}
	if _gdgbg != nil {
		*_gdgbg = *_gbcg
	}
	return _cgag, nil
}

const (
	SelDontCare SelectionValue = iota
	SelHit
	SelMiss
)

func (_bedff *Bitmaps) WidthSorter() func(_abdd, _gafc int) bool {
	return func(_fecg, _fgcb int) bool { return _bedff.Values[_fecg].Width < _bedff.Values[_fgcb].Width }
}
func DilateBrick(d, s *Bitmap, hSize, vSize int) (*Bitmap, error) { return _cef(d, s, hSize, vSize) }
func _afge(_ffgb, _ece int) *Bitmap {
	return &Bitmap{Width: _ffgb, Height: _ece, RowStride: (_ffgb + 7) >> 3}
}
func (_egebc *BitmapsArray) GetBox(i int) (*_fb.Rectangle, error) {
	const _dbed = "\u0042\u0069\u0074\u006dap\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u006f\u0078"
	if _egebc == nil {
		return nil, _d.Error(_dbed, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_egebc.Boxes)-1 {
		return nil, _d.Errorf(_dbed, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _egebc.Boxes[i], nil
}
func TstCSymbol(t *_da.T) *Bitmap {
	t.Helper()
	_gbgc := New(6, 6)
	_dc.NoError(t, _gbgc.SetPixel(1, 0, 1))
	_dc.NoError(t, _gbgc.SetPixel(2, 0, 1))
	_dc.NoError(t, _gbgc.SetPixel(3, 0, 1))
	_dc.NoError(t, _gbgc.SetPixel(4, 0, 1))
	_dc.NoError(t, _gbgc.SetPixel(0, 1, 1))
	_dc.NoError(t, _gbgc.SetPixel(5, 1, 1))
	_dc.NoError(t, _gbgc.SetPixel(0, 2, 1))
	_dc.NoError(t, _gbgc.SetPixel(0, 3, 1))
	_dc.NoError(t, _gbgc.SetPixel(0, 4, 1))
	_dc.NoError(t, _gbgc.SetPixel(5, 4, 1))
	_dc.NoError(t, _gbgc.SetPixel(1, 5, 1))
	_dc.NoError(t, _gbgc.SetPixel(2, 5, 1))
	_dc.NoError(t, _gbgc.SetPixel(3, 5, 1))
	_dc.NoError(t, _gbgc.SetPixel(4, 5, 1))
	return _gbgc
}
func (_ggbb *Bitmaps) GroupByHeight() (*BitmapsArray, error) {
	const _bdbgg = "\u0047\u0072\u006f\u0075\u0070\u0042\u0079\u0048\u0065\u0069\u0067\u0068\u0074"
	if len(_ggbb.Values) == 0 {
		return nil, _d.Error(_bdbgg, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_faedc := &BitmapsArray{}
	_ggbb.SortByHeight()
	_fffg := -1
	_gabb := -1
	for _ggfc := 0; _ggfc < len(_ggbb.Values); _ggfc++ {
		_dcafe := _ggbb.Values[_ggfc].Height
		if _dcafe > _fffg {
			_fffg = _dcafe
			_gabb++
			_faedc.Values = append(_faedc.Values, &Bitmaps{})
		}
		_faedc.Values[_gabb].AddBitmap(_ggbb.Values[_ggfc])
	}
	return _faedc, nil
}
func TstESymbol(t *_da.T, scale ...int) *Bitmap {
	_eadeb, _bacc := NewWithData(4, 5, []byte{0xF0, 0x80, 0xE0, 0x80, 0xF0})
	_dc.NoError(t, _bacc)
	return TstGetScaledSymbol(t, _eadeb, scale...)
}
func (_gecb *Bitmaps) SelectByIndexes(idx []int) (*Bitmaps, error) {
	const _cceag = "B\u0069\u0074\u006d\u0061\u0070\u0073.\u0053\u006f\u0072\u0074\u0049\u006e\u0064\u0065\u0078e\u0073\u0042\u0079H\u0065i\u0067\u0068\u0074"
	_bfce, _fbee := _gecb.selectByIndexes(idx)
	if _fbee != nil {
		return nil, _d.Wrap(_fbee, _cceag, "")
	}
	return _bfce, nil
}

type SizeComparison int

func (_aaec *Bitmap) centroid(_abbe, _dfc []int) (Point, error) {
	_cdcf := Point{}
	_aaec.setPadBits(0)
	if len(_abbe) == 0 {
		_abbe = _dbag()
	}
	if len(_dfc) == 0 {
		_dfc = _gebgfd()
	}
	var _fgec, _fcac, _dbcf, _ffdc, _dgea, _gafe int
	var _addcd byte
	for _dgea = 0; _dgea < _aaec.Height; _dgea++ {
		_ebgb := _aaec.RowStride * _dgea
		_ffdc = 0
		for _gafe = 0; _gafe < _aaec.RowStride; _gafe++ {
			_addcd = _aaec.Data[_ebgb+_gafe]
			if _addcd != 0 {
				_ffdc += _dfc[_addcd]
				_fgec += _abbe[_addcd] + _gafe*8*_dfc[_addcd]
			}
		}
		_dbcf += _ffdc
		_fcac += _ffdc * _dgea
	}
	if _dbcf != 0 {
		_cdcf.X = float32(_fgec) / float32(_dbcf)
		_cdcf.Y = float32(_fcac) / float32(_dbcf)
	}
	return _cdcf, nil
}
func _bgab(_bada *Bitmap, _fbdge, _acac int, _dcab, _edeg int, _ebgf RasterOperator) {
	var (
		_ccefa        int
		_cbcg         byte
		_fbdgg, _bfbc int
		_defe         int
	)
	_gbbb := _dcab >> 3
	_bceb := _dcab & 7
	if _bceb > 0 {
		_cbcg = _fada[_bceb]
	}
	_ccefa = _bada.RowStride*_acac + (_fbdge >> 3)
	switch _ebgf {
	case PixClr:
		for _fbdgg = 0; _fbdgg < _edeg; _fbdgg++ {
			_defe = _ccefa + _fbdgg*_bada.RowStride
			for _bfbc = 0; _bfbc < _gbbb; _bfbc++ {
				_bada.Data[_defe] = 0x0
				_defe++
			}
			if _bceb > 0 {
				_bada.Data[_defe] = _efee(_bada.Data[_defe], 0x0, _cbcg)
			}
		}
	case PixSet:
		for _fbdgg = 0; _fbdgg < _edeg; _fbdgg++ {
			_defe = _ccefa + _fbdgg*_bada.RowStride
			for _bfbc = 0; _bfbc < _gbbb; _bfbc++ {
				_bada.Data[_defe] = 0xff
				_defe++
			}
			if _bceb > 0 {
				_bada.Data[_defe] = _efee(_bada.Data[_defe], 0xff, _cbcg)
			}
		}
	case PixNotDst:
		for _fbdgg = 0; _fbdgg < _edeg; _fbdgg++ {
			_defe = _ccefa + _fbdgg*_bada.RowStride
			for _bfbc = 0; _bfbc < _gbbb; _bfbc++ {
				_bada.Data[_defe] = ^_bada.Data[_defe]
				_defe++
			}
			if _bceb > 0 {
				_bada.Data[_defe] = _efee(_bada.Data[_defe], ^_bada.Data[_defe], _cbcg)
			}
		}
	}
}
func (_eceff *Selection) findMaxTranslations() (_ebeb, _edac, _aega, _gefd int) {
	for _ggad := 0; _ggad < _eceff.Height; _ggad++ {
		for _debgd := 0; _debgd < _eceff.Width; _debgd++ {
			if _eceff.Data[_ggad][_debgd] == SelHit {
				_ebeb = _eec(_ebeb, _eceff.Cx-_debgd)
				_edac = _eec(_edac, _eceff.Cy-_ggad)
				_aega = _eec(_aega, _debgd-_eceff.Cx)
				_gefd = _eec(_gefd, _ggad-_eceff.Cy)
			}
		}
	}
	return _ebeb, _edac, _aega, _gefd
}

const _ceea = 5000

func (_ddfab Points) XSorter() func(_befg, _adcc int) bool {
	return func(_beec, _fec int) bool { return _ddfab[_beec].X < _ddfab[_fec].X }
}
func _dbag() []int {
	_efgf := make([]int, 256)
	_efgf[0] = 0
	_efgf[1] = 7
	var _bdbg int
	for _bdbg = 2; _bdbg < 4; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-2] + 6
	}
	for _bdbg = 4; _bdbg < 8; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-4] + 5
	}
	for _bdbg = 8; _bdbg < 16; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-8] + 4
	}
	for _bdbg = 16; _bdbg < 32; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-16] + 3
	}
	for _bdbg = 32; _bdbg < 64; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-32] + 2
	}
	for _bdbg = 64; _bdbg < 128; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-64] + 1
	}
	for _bdbg = 128; _bdbg < 256; _bdbg++ {
		_efgf[_bdbg] = _efgf[_bdbg-128]
	}
	return _efgf
}

type RasterOperator int

func (_fgg *Bitmap) AddBorder(borderSize, val int) (*Bitmap, error) {
	if borderSize == 0 {
		return _fgg.Copy(), nil
	}
	_gabd, _fbcf := _fgg.addBorderGeneral(borderSize, borderSize, borderSize, borderSize, val)
	if _fbcf != nil {
		return nil, _d.Wrap(_fbcf, "\u0041d\u0064\u0042\u006f\u0072\u0064\u0065r", "")
	}
	return _gabd, nil
}
func (_ccad Points) YSorter() func(_efba, _ggeb int) bool {
	return func(_dgbf, _gecge int) bool { return _ccad[_dgbf].Y < _ccad[_gecge].Y }
}
func (_daagd Points) Get(i int) (Point, error) {
	if i > len(_daagd)-1 {
		return Point{}, _d.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _daagd[i], nil
}
func (_eedb *Bitmap) AddBorderGeneral(left, right, top, bot int, val int) (*Bitmap, error) {
	return _eedb.addBorderGeneral(left, right, top, bot, val)
}
func CorrelationScoreSimple(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_cbbb float64, _dcfc error) {
	const _gfab = "\u0043\u006f\u0072\u0072el\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0053\u0069\u006d\u0070l\u0065"
	if bm1 == nil || bm2 == nil {
		return _cbbb, _d.Error(_gfab, "n\u0069l\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0073 \u0070\u0072\u006f\u0076id\u0065\u0064")
	}
	if tab == nil {
		return _cbbb, _d.Error(_gfab, "\u0074\u0061\u0062\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if area1 == 0 || area2 == 0 {
		return _cbbb, _d.Error(_gfab, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0061\u0072e\u0061\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u0020\u0030")
	}
	_gfdd, _bfdbc := bm1.Width, bm1.Height
	_fefe, _cbgd := bm2.Width, bm2.Height
	if _egfg(_gfdd-_fefe) > maxDiffW {
		return 0, nil
	}
	if _egfg(_bfdbc-_cbgd) > maxDiffH {
		return 0, nil
	}
	var _gcfdc, _abff int
	if delX >= 0 {
		_gcfdc = int(delX + 0.5)
	} else {
		_gcfdc = int(delX - 0.5)
	}
	if delY >= 0 {
		_abff = int(delY + 0.5)
	} else {
		_abff = int(delY - 0.5)
	}
	_daeb := bm1.createTemplate()
	if _dcfc = _daeb.RasterOperation(_gcfdc, _abff, _fefe, _cbgd, PixSrc, bm2, 0, 0); _dcfc != nil {
		return _cbbb, _d.Wrap(_dcfc, _gfab, "\u0062m\u0032 \u0074\u006f\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _dcfc = _daeb.RasterOperation(0, 0, _gfdd, _bfdbc, PixSrcAndDst, bm1, 0, 0); _dcfc != nil {
		return _cbbb, _d.Wrap(_dcfc, _gfab, "b\u006d\u0031\u0020\u0061\u006e\u0064\u0020\u0062\u006d\u0054")
	}
	_deef := _daeb.countPixels()
	_cbbb = float64(_deef) * float64(_deef) / (float64(area1) * float64(area2))
	return _cbbb, nil
}
func (_afd *Bitmap) addBorderGeneral(_fbde, _daff, _ebbf, _gad int, _ddfa int) (*Bitmap, error) {
	const _fcee = "\u0061\u0064d\u0042\u006f\u0072d\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _fbde < 0 || _daff < 0 || _ebbf < 0 || _gad < 0 {
		return nil, _d.Error(_fcee, "n\u0065\u0067\u0061\u0074iv\u0065 \u0062\u006f\u0072\u0064\u0065r\u0020\u0061\u0064\u0064\u0065\u0064")
	}
	_daaga, _ffdf := _afd.Width, _afd.Height
	_gfaa := _daaga + _fbde + _daff
	_eegb := _ffdf + _ebbf + _gad
	_aac := New(_gfaa, _eegb)
	_aac.Color = _afd.Color
	_dafb := PixClr
	if _ddfa > 0 {
		_dafb = PixSet
	}
	_geg := _aac.RasterOperation(0, 0, _fbde, _eegb, _dafb, nil, 0, 0)
	if _geg != nil {
		return nil, _d.Wrap(_geg, _fcee, "\u006c\u0065\u0066\u0074")
	}
	_geg = _aac.RasterOperation(_gfaa-_daff, 0, _daff, _eegb, _dafb, nil, 0, 0)
	if _geg != nil {
		return nil, _d.Wrap(_geg, _fcee, "\u0072\u0069\u0067h\u0074")
	}
	_geg = _aac.RasterOperation(0, 0, _gfaa, _ebbf, _dafb, nil, 0, 0)
	if _geg != nil {
		return nil, _d.Wrap(_geg, _fcee, "\u0074\u006f\u0070")
	}
	_geg = _aac.RasterOperation(0, _eegb-_gad, _gfaa, _gad, _dafb, nil, 0, 0)
	if _geg != nil {
		return nil, _d.Wrap(_geg, _fcee, "\u0062\u006f\u0074\u0074\u006f\u006d")
	}
	_geg = _aac.RasterOperation(_fbde, _ebbf, _daaga, _ffdf, PixSrc, _afd, 0, 0)
	if _geg != nil {
		return nil, _d.Wrap(_geg, _fcee, "\u0063\u006f\u0070\u0079")
	}
	return _aac, nil
}
func (_dcac *Bitmap) setEightFullBytes(_add int, _ffcf uint64) error {
	if _add+7 > len(_dcac.Data)-1 {
		return _d.Error("\u0073\u0065\u0074\u0045\u0069\u0067\u0068\u0074\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dcac.Data[_add] = byte((_ffcf & 0xff00000000000000) >> 56)
	_dcac.Data[_add+1] = byte((_ffcf & 0xff000000000000) >> 48)
	_dcac.Data[_add+2] = byte((_ffcf & 0xff0000000000) >> 40)
	_dcac.Data[_add+3] = byte((_ffcf & 0xff00000000) >> 32)
	_dcac.Data[_add+4] = byte((_ffcf & 0xff000000) >> 24)
	_dcac.Data[_add+5] = byte((_ffcf & 0xff0000) >> 16)
	_dcac.Data[_add+6] = byte((_ffcf & 0xff00) >> 8)
	_dcac.Data[_add+7] = byte(_ffcf & 0xff)
	return nil
}
func (_acddc *Bitmaps) String() string {
	_geabb := _e.Builder{}
	for _, _fgfe := range _acddc.Values {
		_geabb.WriteString(_fgfe.String())
		_geabb.WriteRune('\n')
	}
	return _geabb.String()
}
func _dea() (_cff [256]uint64) {
	for _ff := 0; _ff < 256; _ff++ {
		if _ff&0x01 != 0 {
			_cff[_ff] |= 0xff
		}
		if _ff&0x02 != 0 {
			_cff[_ff] |= 0xff00
		}
		if _ff&0x04 != 0 {
			_cff[_ff] |= 0xff0000
		}
		if _ff&0x08 != 0 {
			_cff[_ff] |= 0xff000000
		}
		if _ff&0x10 != 0 {
			_cff[_ff] |= 0xff00000000
		}
		if _ff&0x20 != 0 {
			_cff[_ff] |= 0xff0000000000
		}
		if _ff&0x40 != 0 {
			_cff[_ff] |= 0xff000000000000
		}
		if _ff&0x80 != 0 {
			_cff[_ff] |= 0xff00000000000000
		}
	}
	return _cff
}
func (_cdcg *Bitmap) CreateTemplate() *Bitmap { return _cdcg.createTemplate() }
func TstVSymbol(t *_da.T, scale ...int) *Bitmap {
	_dbfe, _fgcgg := NewWithData(5, 5, []byte{0x88, 0x88, 0x88, 0x50, 0x20})
	_dc.NoError(t, _fgcgg)
	return TstGetScaledSymbol(t, _dbfe, scale...)
}
func (_gceaa *Bitmap) SetPixel(x, y int, pixel byte) error {
	_abeb := _gceaa.GetByteIndex(x, y)
	if _abeb > len(_gceaa.Data)-1 {
		return _d.Errorf("\u0053\u0065\u0074\u0050\u0069\u0078\u0065\u006c", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", _abeb)
	}
	_gcfe := _gceaa.GetBitOffset(x)
	_cfb := uint(7 - _gcfe)
	_fed := _gceaa.Data[_abeb]
	var _ffgbg byte
	if pixel == 1 {
		_ffgbg = _fed | (pixel & 0x01 << _cfb)
	} else {
		_ffgbg = _fed &^ (1 << _cfb)
	}
	_gceaa.Data[_abeb] = _ffgbg
	return nil
}
func (_eace *Bitmap) connComponentsBB(_cfef int) (_gbd *Boxes, _aedcd error) {
	const _cggbb = "\u0042\u0069\u0074ma\u0070\u002e\u0063\u006f\u006e\u006e\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0042"
	if _cfef != 4 && _cfef != 8 {
		return nil, _d.Error(_cggbb, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _eace.Zero() {
		return &Boxes{}, nil
	}
	_eace.setPadBits(0)
	_aeed, _aedcd := _bce(nil, _eace)
	if _aedcd != nil {
		return nil, _d.Wrap(_aedcd, _cggbb, "\u0062\u006d\u0031")
	}
	_cffb := &_b.Stack{}
	_cffb.Aux = &_b.Stack{}
	_gbd = &Boxes{}
	var (
		_bdff, _beaf int
		_begc        _fb.Point
		_dbdc        bool
		_eccc        *_fb.Rectangle
	)
	for {
		if _begc, _dbdc, _aedcd = _aeed.nextOnPixel(_beaf, _bdff); _aedcd != nil {
			return nil, _d.Wrap(_aedcd, _cggbb, "")
		}
		if !_dbdc {
			break
		}
		if _eccc, _aedcd = _gegb(_aeed, _cffb, _begc.X, _begc.Y, _cfef); _aedcd != nil {
			return nil, _d.Wrap(_aedcd, _cggbb, "")
		}
		if _aedcd = _gbd.Add(_eccc); _aedcd != nil {
			return nil, _d.Wrap(_aedcd, _cggbb, "")
		}
		_beaf = _begc.X
		_bdff = _begc.Y
	}
	return _gbd, nil
}
func _edbe(_dcb, _de *Bitmap) (_bca error) {
	const _fce = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0038"
	_gce := _de.RowStride
	_eb := _dcb.RowStride
	var _bf, _fga, _ab, _ggg, _bb int
	for _ab = 0; _ab < _de.Height; _ab++ {
		_bf = _ab * _gce
		_fga = 8 * _ab * _eb
		for _ggg = 0; _ggg < _gce; _ggg++ {
			if _bca = _dcb.setEightBytes(_fga+_ggg*8, _fbcd[_de.Data[_bf+_ggg]]); _bca != nil {
				return _d.Wrap(_bca, _fce, "")
			}
		}
		for _bb = 1; _bb < 8; _bb++ {
			for _ggg = 0; _ggg < _eb; _ggg++ {
				if _bca = _dcb.SetByte(_fga+_bb*_eb+_ggg, _dcb.Data[_fga+_ggg]); _bca != nil {
					return _d.Wrap(_bca, _fce, "")
				}
			}
		}
	}
	return nil
}

type byWidth Bitmaps

func (_ecee *Bitmap) countPixels() int {
	var (
		_dff   int
		_aecdg uint8
		_edf   byte
		_eaf   int
	)
	_bacd := _ecee.RowStride
	_dbef := uint(_ecee.Width & 0x07)
	if _dbef != 0 {
		_aecdg = uint8((0xff << (8 - _dbef)) & 0xff)
		_bacd--
	}
	for _afc := 0; _afc < _ecee.Height; _afc++ {
		for _eaf = 0; _eaf < _bacd; _eaf++ {
			_edf = _ecee.Data[_afc*_ecee.RowStride+_eaf]
			_dff += int(_cdc[_edf])
		}
		if _dbef != 0 {
			_dff += int(_cdc[_ecee.Data[_afc*_ecee.RowStride+_eaf]&_aecdg])
		}
	}
	return _dff
}

var _eged = [5]int{1, 2, 3, 0, 4}

func _gdgb() (_dfb []byte) {
	_dfb = make([]byte, 256)
	for _cdd := 0; _cdd < 256; _cdd++ {
		_egce := byte(_cdd)
		_dfb[_egce] = (_egce & 0x01) | ((_egce & 0x04) >> 1) | ((_egce & 0x10) >> 2) | ((_egce & 0x40) >> 3) | ((_egce & 0x02) << 3) | ((_egce & 0x08) << 2) | ((_egce & 0x20) << 1) | (_egce & 0x80)
	}
	return _dfb
}
func (_gdff *Boxes) selectWithIndicator(_affe *_b.NumSlice) (_addf *Boxes, _bgc error) {
	const _aaac = "\u0042o\u0078\u0065\u0073\u002es\u0065\u006c\u0065\u0063\u0074W\u0069t\u0068I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _gdff == nil {
		return nil, _d.Error(_aaac, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _affe == nil {
		return nil, _d.Error(_aaac, "\u0027\u006ea\u0027\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_affe) != len(*_gdff) {
		return nil, _d.Error(_aaac, "\u0062\u006f\u0078\u0065\u0073\u0020\u0027\u0062\u0027\u0020\u0068\u0061\u0073\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020s\u0069\u007a\u0065\u0020\u0074h\u0061\u006e \u0027\u006e\u0061\u0027")
	}
	var _aedc, _ggaa int
	for _abdgg := 0; _abdgg < len(*_affe); _abdgg++ {
		if _aedc, _bgc = _affe.GetInt(_abdgg); _bgc != nil {
			return nil, _d.Wrap(_bgc, _aaac, "\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u0067\u0020c\u006f\u0075\u006e\u0074")
		}
		if _aedc == 1 {
			_ggaa++
		}
	}
	if _ggaa == len(*_gdff) {
		return _gdff, nil
	}
	_bbbe := Boxes{}
	for _agbg := 0; _agbg < len(*_affe); _agbg++ {
		_aedc = int((*_affe)[_agbg])
		if _aedc == 0 {
			continue
		}
		_bbbe = append(_bbbe, (*_gdff)[_agbg])
	}
	_addf = &_bbbe
	return _addf, nil
}
func (_bfdd *Bitmap) resizeImageData(_fef *Bitmap) error {
	if _fef == nil {
		return _d.Error("\u0072e\u0073i\u007a\u0065\u0049\u006d\u0061\u0067\u0065\u0044\u0061\u0074\u0061", "\u0073r\u0063 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _bfdd.SizesEqual(_fef) {
		return nil
	}
	_bfdd.Data = make([]byte, len(_fef.Data))
	_bfdd.Width = _fef.Width
	_bfdd.Height = _fef.Height
	_bfdd.RowStride = _fef.RowStride
	return nil
}
func init() {
	const _ebfa = "\u0062\u0069\u0074\u006dap\u0073\u002e\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0061\u0074\u0069o\u006e"
	_bfgad = New(50, 40)
	var _dfgb error
	_bfgad, _dfgb = _bfgad.AddBorder(2, 1)
	if _dfgb != nil {
		panic(_d.Wrap(_dfgb, _ebfa, "f\u0072\u0061\u006d\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
	_edcgg, _dfgb = NewWithData(50, 22, _adba)
	if _dfgb != nil {
		panic(_d.Wrap(_dfgb, _ebfa, "i\u006d\u0061\u0067\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
}
func MakePixelSumTab8() []int { return _gebgfd() }

var _cdc [256]uint8

func (_fe *Bitmap) CountPixels() int { return _fe.countPixels() }
func _dfda(_dbgg *Bitmap, _egeb ...MorphProcess) (_cdeg *Bitmap, _abcf error) {
	const _abge = "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065"
	if _dbgg == nil {
		return nil, _d.Error(_abge, "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065 \u0073\u006f\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if len(_egeb) == 0 {
		return nil, _d.Error(_abge, "m\u006f\u0072\u0070\u0068\u0053\u0065q\u0075\u0065\u006e\u0063\u0065\u002c \u0073\u0065\u0071\u0075\u0065\u006e\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _abcf = _edbfd(_egeb...); _abcf != nil {
		return nil, _d.Wrap(_abcf, _abge, "")
	}
	var _fadeb, _bfc, _cabcd int
	_cdeg = _dbgg.Copy()
	for _, _cdedb := range _egeb {
		switch _cdedb.Operation {
		case MopDilation:
			_fadeb, _bfc = _cdedb.getWidthHeight()
			_cdeg, _abcf = DilateBrick(nil, _cdeg, _fadeb, _bfc)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopErosion:
			_fadeb, _bfc = _cdedb.getWidthHeight()
			_cdeg, _abcf = _fgcd(nil, _cdeg, _fadeb, _bfc)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopOpening:
			_fadeb, _bfc = _cdedb.getWidthHeight()
			_cdeg, _abcf = _fgf(nil, _cdeg, _fadeb, _bfc)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopClosing:
			_fadeb, _bfc = _cdedb.getWidthHeight()
			_cdeg, _abcf = _dcag(nil, _cdeg, _fadeb, _bfc)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopRankBinaryReduction:
			_cdeg, _abcf = _ebf(_cdeg, _cdedb.Arguments...)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopReplicativeBinaryExpansion:
			_cdeg, _abcf = _cbac(_cdeg, _cdedb.Arguments[0])
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		case MopAddBorder:
			_cabcd = _cdedb.Arguments[0]
			_cdeg, _abcf = _cdeg.AddBorder(_cabcd, 0)
			if _abcf != nil {
				return nil, _d.Wrap(_abcf, _abge, "")
			}
		default:
			return nil, _d.Error(_abge, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u006fr\u0070\u0068\u004f\u0070\u0065\u0072\u0061ti\u006f\u006e\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u006f t\u0068\u0065 \u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065")
		}
	}
	if _cabcd > 0 {
		_cdeg, _abcf = _cdeg.RemoveBorder(_cabcd)
		if _abcf != nil {
			return nil, _d.Wrap(_abcf, _abge, "\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u003e\u0020\u0030")
		}
	}
	return _cdeg, nil
}
func _gegb(_cgfdg *Bitmap, _gace *_b.Stack, _bcbd, _facf, _bffa int) (_fbec *_fb.Rectangle, _aacg error) {
	const _eddd = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _cgfdg == nil {
		return nil, _d.Error(_eddd, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _gace == nil {
		return nil, _d.Error(_eddd, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	switch _bffa {
	case 4:
		if _fbec, _aacg = _gaeee(_cgfdg, _gace, _bcbd, _facf); _aacg != nil {
			return nil, _d.Wrap(_aacg, _eddd, "")
		}
		return _fbec, nil
	case 8:
		if _fbec, _aacg = _acbd(_cgfdg, _gace, _bcbd, _facf); _aacg != nil {
			return nil, _d.Wrap(_aacg, _eddd, "")
		}
		return _fbec, nil
	default:
		return nil, _d.Errorf(_eddd, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u0069\u0073 \u006eo\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038\u003a\u0020\u0027\u0025\u0064\u0027", _bffa)
	}
}
func _geab(_eceb *_b.Stack, _edcf, _decg, _fcfd, _fdada, _eacdg int, _edcg *_fb.Rectangle) (_dfdc error) {
	const _adcg = "\u0070\u0075\u0073\u0068\u0046\u0069\u006c\u006c\u0053\u0065\u0067m\u0065\u006e\u0074\u0042\u006f\u0075\u006e\u0064\u0069\u006eg\u0042\u006f\u0078"
	if _eceb == nil {
		return _d.Error(_adcg, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _edcg == nil {
		return _d.Error(_adcg, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0069\u006da\u0067e\u002e\u0052\u0065\u0063\u0074\u0061\u006eg\u006c\u0065")
	}
	_edcg.Min.X = _b.Min(_edcg.Min.X, _edcf)
	_edcg.Max.X = _b.Max(_edcg.Max.X, _decg)
	_edcg.Min.Y = _b.Min(_edcg.Min.Y, _fcfd)
	_edcg.Max.Y = _b.Max(_edcg.Max.Y, _fcfd)
	if !(_fcfd+_fdada >= 0 && _fcfd+_fdada <= _eacdg) {
		return nil
	}
	if _eceb.Aux == nil {
		return _d.Error(_adcg, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	var _acdf *fillSegment
	_beda, _ggbaa := _eceb.Aux.Pop()
	if _ggbaa {
		if _acdf, _ggbaa = _beda.(*fillSegment); !_ggbaa {
			return _d.Error(_adcg, "a\u0075\u0078\u0053\u0074\u0061\u0063k\u0020\u0064\u0061\u0074\u0061\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061 \u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065n\u0074")
		}
	} else {
		_acdf = &fillSegment{}
	}
	_acdf._gebc = _edcf
	_acdf._fdeg = _decg
	_acdf._aegc = _fcfd
	_acdf._bbcgb = _fdada
	_eceb.Push(_acdf)
	return nil
}
func (_dfbc Points) GetIntY(i int) (int, error) {
	if i >= len(_dfbc) {
		return 0, _d.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0059", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_dfbc[i].Y), nil
}
func (_daggb *Bitmaps) CountPixels() *_b.NumSlice {
	_gbab := &_b.NumSlice{}
	for _, _gegc := range _daggb.Values {
		_gbab.AddInt(_gegc.CountPixels())
	}
	return _gbab
}
func (_fdbf *Bitmap) RemoveBorderGeneral(left, right, top, bot int) (*Bitmap, error) {
	return _fdbf.removeBorderGeneral(left, right, top, bot)
}
func (_ebda *Bitmap) GetComponents(components Component, maxWidth, maxHeight int) (_egab *Bitmaps, _fade *Boxes, _bgff error) {
	const _afdf = "B\u0069t\u006d\u0061\u0070\u002e\u0047\u0065\u0074\u0043o\u006d\u0070\u006f\u006een\u0074\u0073"
	if _ebda == nil {
		return nil, nil, _d.Error(_afdf, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0042\u0069\u0074\u006da\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064\u002e")
	}
	switch components {
	case ComponentConn, ComponentCharacters, ComponentWords:
	default:
		return nil, nil, _d.Error(_afdf, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074s\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072")
	}
	if _ebda.Zero() {
		_fade = &Boxes{}
		_egab = &Bitmaps{}
		return _egab, _fade, nil
	}
	switch components {
	case ComponentConn:
		_egab = &Bitmaps{}
		if _fade, _bgff = _ebda.ConnComponents(_egab, 8); _bgff != nil {
			return nil, nil, _d.Wrap(_bgff, _afdf, "\u006e\u006f \u0070\u0072\u0065p\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentCharacters:
		_begb, _geee := MorphSequence(_ebda, MorphProcess{Operation: MopClosing, Arguments: []int{1, 6}})
		if _geee != nil {
			return nil, nil, _d.Wrap(_geee, _afdf, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _ea.Log.IsLogLevel(_ea.LogLevelTrace) {
			_ea.Log.Trace("\u0043o\u006d\u0070o\u006e\u0065\u006e\u0074C\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074ma\u0070\u0020\u0061f\u0074\u0065r\u0020\u0063\u006c\u006f\u0073\u0069n\u0067\u003a \u0025\u0073", _begb.String())
		}
		_eebb := &Bitmaps{}
		_fade, _geee = _begb.ConnComponents(_eebb, 8)
		if _geee != nil {
			return nil, nil, _d.Wrap(_geee, _afdf, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _ea.Log.IsLogLevel(_ea.LogLevelTrace) {
			_ea.Log.Trace("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0043\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0020a\u0066\u0074\u0065\u0072\u0020\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076i\u0074y\u003a\u0020\u0025\u0073", _eebb.String())
		}
		if _egab, _geee = _eebb.ClipToBitmap(_ebda); _geee != nil {
			return nil, nil, _d.Wrap(_geee, _afdf, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentWords:
		_bccdf := 1
		var _gfac *Bitmap
		switch {
		case _ebda.XResolution <= 200:
			_gfac = _ebda
		case _ebda.XResolution <= 400:
			_bccdf = 2
			_gfac, _bgff = _ebf(_ebda, 1, 0, 0, 0)
			if _bgff != nil {
				return nil, nil, _d.Wrap(_bgff, _afdf, "w\u006f\u0072\u0064\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0020\u002d \u0078\u0072\u0065s\u003c=\u0034\u0030\u0030")
			}
		default:
			_bccdf = 4
			_gfac, _bgff = _ebf(_ebda, 1, 1, 0, 0)
			if _bgff != nil {
				return nil, nil, _d.Wrap(_bgff, _afdf, "\u0077\u006f\u0072\u0064 \u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073 \u002d \u0078\u0072\u0065\u0073\u0020\u003e\u00204\u0030\u0030")
			}
		}
		_gcfa, _, _egaf := _efgd(_gfac)
		if _egaf != nil {
			return nil, nil, _d.Wrap(_egaf, _afdf, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_gccb, _egaf := _cbac(_gcfa, _bccdf)
		if _egaf != nil {
			return nil, nil, _d.Wrap(_egaf, _afdf, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_dagg := &Bitmaps{}
		if _fade, _egaf = _gccb.ConnComponents(_dagg, 4); _egaf != nil {
			return nil, nil, _d.Wrap(_egaf, _afdf, "\u0077\u006f\u0072\u0064\u0020\u0070r\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002c\u0020\u0063\u006f\u006en\u0065\u0063\u0074\u0020\u0065\u0078\u0070a\u006e\u0064\u0065\u0064")
		}
		if _egab, _egaf = _dagg.ClipToBitmap(_ebda); _egaf != nil {
			return nil, nil, _d.Wrap(_egaf, _afdf, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
	}
	_egab, _bgff = _egab.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _bgff != nil {
		return nil, nil, _d.Wrap(_bgff, _afdf, "")
	}
	_fade, _bgff = _fade.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _bgff != nil {
		return nil, nil, _d.Wrap(_bgff, _afdf, "")
	}
	return _egab, _fade, nil
}
func TstWordBitmapWithSpaces(t *_da.T, scale ...int) *Bitmap {
	_daagg := 1
	if len(scale) > 0 {
		_daagg = scale[0]
	}
	_acdfa := 3
	_dbce := 9 + 7 + 15 + 2*_acdfa + 2*_acdfa
	_cbfde := 5 + _acdfa + 5 + 2*_acdfa
	_edcgf := New(_dbce*_daagg, _cbfde*_daagg)
	_beea := &Bitmaps{}
	var _gabbg *int
	_acdfa *= _daagg
	_caad := _acdfa
	_gabbg = &_caad
	_egfc := _acdfa
	_cabac := TstDSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, _acdfa)
	_cabac = TstISymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstTSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, _acdfa)
	_cabac = TstNSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstWSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 0)
	*_gabbg = _acdfa
	_egfc = 5*_daagg + _acdfa
	_cabac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstRSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, _acdfa)
	_cabac = TstNSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstESymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstVSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstESymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 1*_daagg)
	_cabac = TstRSymbol(t, scale...)
	TstAddSymbol(t, _beea, _cabac, _gabbg, _egfc, 0)
	TstWriteSymbols(t, _beea, _edcgf)
	return _edcgf
}
func (_cdab *byHeight) Len() int { return len(_cdab.Values) }
func _agbc(_dcfd, _bffg *Bitmap, _bfec *Selection) (*Bitmap, error) {
	const _ecbd = "\u0065\u0072\u006fd\u0065"
	var (
		_dfffd error
		_dbea  *Bitmap
	)
	_dcfd, _dfffd = _fccf(_dcfd, _bffg, _bfec, &_dbea)
	if _dfffd != nil {
		return nil, _d.Wrap(_dfffd, _ecbd, "")
	}
	if _dfffd = _dcfd.setAll(); _dfffd != nil {
		return nil, _d.Wrap(_dfffd, _ecbd, "")
	}
	var _gefg SelectionValue
	for _bbfb := 0; _bbfb < _bfec.Height; _bbfb++ {
		for _egde := 0; _egde < _bfec.Width; _egde++ {
			_gefg = _bfec.Data[_bbfb][_egde]
			if _gefg == SelHit {
				_dfffd = _fdfa(_dcfd, _bfec.Cx-_egde, _bfec.Cy-_bbfb, _bffg.Width, _bffg.Height, PixSrcAndDst, _dbea, 0, 0)
				if _dfffd != nil {
					return nil, _d.Wrap(_dfffd, _ecbd, "")
				}
			}
		}
	}
	if MorphBC == SymmetricMorphBC {
		return _dcfd, nil
	}
	_ebdf, _egbg, _gfga, _dfad := _bfec.findMaxTranslations()
	if _ebdf > 0 {
		if _dfffd = _dcfd.RasterOperation(0, 0, _ebdf, _bffg.Height, PixClr, nil, 0, 0); _dfffd != nil {
			return nil, _d.Wrap(_dfffd, _ecbd, "\u0078\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _gfga > 0 {
		if _dfffd = _dcfd.RasterOperation(_bffg.Width-_gfga, 0, _gfga, _bffg.Height, PixClr, nil, 0, 0); _dfffd != nil {
			return nil, _d.Wrap(_dfffd, _ecbd, "\u0078\u006e\u0020\u003e\u0020\u0030")
		}
	}
	if _egbg > 0 {
		if _dfffd = _dcfd.RasterOperation(0, 0, _bffg.Width, _egbg, PixClr, nil, 0, 0); _dfffd != nil {
			return nil, _d.Wrap(_dfffd, _ecbd, "\u0079\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _dfad > 0 {
		if _dfffd = _dcfd.RasterOperation(0, _bffg.Height-_dfad, _bffg.Width, _dfad, PixClr, nil, 0, 0); _dfffd != nil {
			return nil, _d.Wrap(_dfffd, _ecbd, "\u0079\u006e\u0020\u003e\u0020\u0030")
		}
	}
	return _dcfd, nil
}
func CombineBytes(oldByte, newByte byte, op CombinationOperator) byte {
	return _adda(oldByte, newByte, op)
}
func RankHausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH, area1, area3 int, rank float32, tab8 []int) (_afga bool, _bddg error) {
	const _fbfc = "\u0052\u0061\u006ek\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_efdc, _cbag := p1.Width, p1.Height
	_gge, _fgaf := p3.Width, p3.Height
	if _b.Abs(_efdc-_gge) > maxDiffW {
		return false, nil
	}
	if _b.Abs(_cbag-_fgaf) > maxDiffH {
		return false, nil
	}
	_afgag := int(float32(area1)*(1.0-rank) + 0.5)
	_ccg := int(float32(area3)*(1.0-rank) + 0.5)
	var _aeda, _daad int
	if delX >= 0 {
		_aeda = int(delX + 0.5)
	} else {
		_aeda = int(delX - 0.5)
	}
	if delY >= 0 {
		_daad = int(delY + 0.5)
	} else {
		_daad = int(delY - 0.5)
	}
	_gfbb := p1.CreateTemplate()
	if _bddg = _gfbb.RasterOperation(0, 0, _efdc, _cbag, PixSrc, p1, 0, 0); _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _bddg = _gfbb.RasterOperation(_aeda, _daad, _efdc, _cbag, PixNotSrcAndDst, p4, 0, 0); _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "\u0074 \u0026\u0020\u0021\u0070\u0034")
	}
	_afga, _bddg = _gfbb.ThresholdPixelSum(_afgag, tab8)
	if _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0031")
	}
	if _afga {
		return false, nil
	}
	if _bddg = _gfbb.RasterOperation(_aeda, _daad, _gge, _fgaf, PixSrc, p3, 0, 0); _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _bddg = _gfbb.RasterOperation(0, 0, _gge, _fgaf, PixNotSrcAndDst, p2, 0, 0); _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "\u0074 \u0026\u0020\u0021\u0070\u0032")
	}
	_afga, _bddg = _gfbb.ThresholdPixelSum(_ccg, tab8)
	if _bddg != nil {
		return false, _d.Wrap(_bddg, _fbfc, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0033")
	}
	return !_afga, nil
}
func _fbgc(_ecca, _eeae int) int {
	if _ecca < _eeae {
		return _ecca
	}
	return _eeae
}
func _efgb(_bbaa *Bitmap, _daegd, _gccg int, _aafe, _aada int, _gdde RasterOperator) {
	var (
		_bgae  bool
		_ageeb bool
		_gaee  int
		_eaaf  int
		_bcee  int
		_dbbac int
		_bdae  bool
		_fbfb  byte
	)
	_bccbe := 8 - (_daegd & 7)
	_begf := _eaec[_bccbe]
	_ffac := _bbaa.RowStride*_gccg + (_daegd >> 3)
	if _aafe < _bccbe {
		_bgae = true
		_begf &= _fada[8-_bccbe+_aafe]
	}
	if !_bgae {
		_gaee = (_aafe - _bccbe) >> 3
		if _gaee != 0 {
			_ageeb = true
			_eaaf = _ffac + 1
		}
	}
	_bcee = (_daegd + _aafe) & 7
	if !(_bgae || _bcee == 0) {
		_bdae = true
		_fbfb = _fada[_bcee]
		_dbbac = _ffac + 1 + _gaee
	}
	var _cfgc, _ffca int
	switch _gdde {
	case PixClr:
		for _cfgc = 0; _cfgc < _aada; _cfgc++ {
			_bbaa.Data[_ffac] = _efee(_bbaa.Data[_ffac], 0x0, _begf)
			_ffac += _bbaa.RowStride
		}
		if _ageeb {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				for _ffca = 0; _ffca < _gaee; _ffca++ {
					_bbaa.Data[_eaaf+_ffca] = 0x0
				}
				_eaaf += _bbaa.RowStride
			}
		}
		if _bdae {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				_bbaa.Data[_dbbac] = _efee(_bbaa.Data[_dbbac], 0x0, _fbfb)
				_dbbac += _bbaa.RowStride
			}
		}
	case PixSet:
		for _cfgc = 0; _cfgc < _aada; _cfgc++ {
			_bbaa.Data[_ffac] = _efee(_bbaa.Data[_ffac], 0xff, _begf)
			_ffac += _bbaa.RowStride
		}
		if _ageeb {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				for _ffca = 0; _ffca < _gaee; _ffca++ {
					_bbaa.Data[_eaaf+_ffca] = 0xff
				}
				_eaaf += _bbaa.RowStride
			}
		}
		if _bdae {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				_bbaa.Data[_dbbac] = _efee(_bbaa.Data[_dbbac], 0xff, _fbfb)
				_dbbac += _bbaa.RowStride
			}
		}
	case PixNotDst:
		for _cfgc = 0; _cfgc < _aada; _cfgc++ {
			_bbaa.Data[_ffac] = _efee(_bbaa.Data[_ffac], ^_bbaa.Data[_ffac], _begf)
			_ffac += _bbaa.RowStride
		}
		if _ageeb {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				for _ffca = 0; _ffca < _gaee; _ffca++ {
					_bbaa.Data[_eaaf+_ffca] = ^(_bbaa.Data[_eaaf+_ffca])
				}
				_eaaf += _bbaa.RowStride
			}
		}
		if _bdae {
			for _cfgc = 0; _cfgc < _aada; _cfgc++ {
				_bbaa.Data[_dbbac] = _efee(_bbaa.Data[_dbbac], ^_bbaa.Data[_dbbac], _fbfb)
				_dbbac += _bbaa.RowStride
			}
		}
	}
}
func _aef(_dacbd *Bitmap, _afaf *Bitmap, _cadbf *Selection) (*Bitmap, error) {
	var (
		_gfdde *Bitmap
		_egbd  error
	)
	_dacbd, _egbd = _fccf(_dacbd, _afaf, _cadbf, &_gfdde)
	if _egbd != nil {
		return nil, _egbd
	}
	if _egbd = _dacbd.clearAll(); _egbd != nil {
		return nil, _egbd
	}
	var _fbca SelectionValue
	for _gadg := 0; _gadg < _cadbf.Height; _gadg++ {
		for _fcef := 0; _fcef < _cadbf.Width; _fcef++ {
			_fbca = _cadbf.Data[_gadg][_fcef]
			if _fbca == SelHit {
				if _egbd = _dacbd.RasterOperation(_fcef-_cadbf.Cx, _gadg-_cadbf.Cy, _afaf.Width, _afaf.Height, PixSrcOrDst, _gfdde, 0, 0); _egbd != nil {
					return nil, _egbd
				}
			}
		}
	}
	return _dacbd, nil
}
func (_cgfab *byWidth) Len() int { return len(_cgfab.Values) }
func (_bdd *Bitmap) SizesEqual(s *Bitmap) bool {
	if _bdd == s {
		return true
	}
	if _bdd.Width != s.Width || _bdd.Height != s.Height {
		return false
	}
	return true
}
func (_ffd *Bitmap) GetUnpaddedData() ([]byte, error) {
	_badd := uint(_ffd.Width & 0x07)
	if _badd == 0 {
		return _ffd.Data, nil
	}
	_fcc := _ffd.Width * _ffd.Height
	if _fcc%8 != 0 {
		_fcc >>= 3
		_fcc++
	} else {
		_fcc >>= 3
	}
	_cac := make([]byte, _fcc)
	_eddb := _bg.NewWriterMSB(_cac)
	const _aec = "\u0047e\u0074U\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	for _gecc := 0; _gecc < _ffd.Height; _gecc++ {
		for _gbe := 0; _gbe < _ffd.RowStride; _gbe++ {
			_ddge := _ffd.Data[_gecc*_ffd.RowStride+_gbe]
			if _gbe != _ffd.RowStride-1 {
				_fea := _eddb.WriteByte(_ddge)
				if _fea != nil {
					return nil, _d.Wrap(_fea, _aec, "")
				}
				continue
			}
			for _fgcg := uint(0); _fgcg < _badd; _fgcg++ {
				_abe := _eddb.WriteBit(int(_ddge >> (7 - _fgcg) & 0x01))
				if _abe != nil {
					return nil, _d.Wrap(_abe, _aec, "")
				}
			}
		}
	}
	return _cac, nil
}
func _bgg(_gage, _fdb *Bitmap, _gdg int, _aca []byte, _acc int) (_gcf error) {
	const _dgd = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0031"
	var (
		_cffd, _cbd, _ged, _ebdc, _cecd, _ecg, _ad, _ecab int
		_ag, _bad                                         uint32
		_geb, _cca                                        byte
		_cdgd                                             uint16
	)
	_egc := make([]byte, 4)
	_ddd := make([]byte, 4)
	for _ged = 0; _ged < _gage.Height-1; _ged, _ebdc = _ged+2, _ebdc+1 {
		_cffd = _ged * _gage.RowStride
		_cbd = _ebdc * _fdb.RowStride
		for _cecd, _ecg = 0, 0; _cecd < _acc; _cecd, _ecg = _cecd+4, _ecg+1 {
			for _ad = 0; _ad < 4; _ad++ {
				_ecab = _cffd + _cecd + _ad
				if _ecab <= len(_gage.Data)-1 && _ecab < _cffd+_gage.RowStride {
					_egc[_ad] = _gage.Data[_ecab]
				} else {
					_egc[_ad] = 0x00
				}
				_ecab = _cffd + _gage.RowStride + _cecd + _ad
				if _ecab <= len(_gage.Data)-1 && _ecab < _cffd+(2*_gage.RowStride) {
					_ddd[_ad] = _gage.Data[_ecab]
				} else {
					_ddd[_ad] = 0x00
				}
			}
			_ag = _df.BigEndian.Uint32(_egc)
			_bad = _df.BigEndian.Uint32(_ddd)
			_bad |= _ag
			_bad |= _bad << 1
			_bad &= 0xaaaaaaaa
			_ag = _bad | (_bad << 7)
			_geb = byte(_ag >> 24)
			_cca = byte((_ag >> 8) & 0xff)
			_ecab = _cbd + _ecg
			if _ecab+1 == len(_fdb.Data)-1 || _ecab+1 >= _cbd+_fdb.RowStride {
				_fdb.Data[_ecab] = _aca[_geb]
			} else {
				_cdgd = (uint16(_aca[_geb]) << 8) | uint16(_aca[_cca])
				if _gcf = _fdb.setTwoBytes(_ecab, _cdgd); _gcf != nil {
					return _d.Wrapf(_gcf, _dgd, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ecab)
				}
				_ecg++
			}
		}
	}
	return nil
}
