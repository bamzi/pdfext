package mmr

import (
	_bg "errors"
	_d "fmt"
	_b "io"

	_da "github.com/bamzi/pdfext/common"
	_e "github.com/bamzi/pdfext/internal/bitwise"
	_c "github.com/bamzi/pdfext/internal/jbig2/bitmap"
)

func (_bbe *runData) uncompressGetNextCodeLittleEndian() (int, error) {
	_eff := _bbe._ffdd - _bbe._fad
	if _eff < 0 || _eff > 24 {
		_dab := (_bbe._ffdd >> 3) - _bbe._cef
		if _dab >= _bbe._dace {
			_dab += _bbe._cef
			if _adf := _bbe.fillBuffer(_dab); _adf != nil {
				return 0, _adf
			}
			_dab -= _bbe._cef
		}
		_bbff := (uint32(_bbe._agfc[_dab]&0xFF) << 16) | (uint32(_bbe._agfc[_dab+1]&0xFF) << 8) | (uint32(_bbe._agfc[_dab+2] & 0xFF))
		_gea := uint32(_bbe._ffdd & 7)
		_bbff <<= _gea
		_bbe._eabg = int(_bbff)
	} else {
		_cbb := _bbe._fad & 7
		_eaab := 7 - _cbb
		if _eff <= _eaab {
			_bbe._eabg <<= uint(_eff)
		} else {
			_ga := (_bbe._fad >> 3) + 3 - _bbe._cef
			if _ga >= _bbe._dace {
				_ga += _bbe._cef
				if _gbb := _bbe.fillBuffer(_ga); _gbb != nil {
					return 0, _gbb
				}
				_ga -= _bbe._cef
			}
			_cbb = 8 - _cbb
			for {
				_bbe._eabg <<= uint(_cbb)
				_bbe._eabg |= int(uint(_bbe._agfc[_ga]) & 0xFF)
				_eff -= _cbb
				_ga++
				_cbb = 8
				if !(_eff >= 8) {
					break
				}
			}
			_bbe._eabg <<= uint(_eff)
		}
	}
	_bbe._fad = _bbe._ffdd
	return _bbe._eabg, nil
}
func (_cfg *runData) uncompressGetCode(_cbc []*code) (*code, error) {
	return _cfg.uncompressGetCodeLittleEndian(_cbc)
}
func (_ceg *Decoder) uncompress1d(_eaf *runData, _bcf []int, _bfb int) (int, error) {
	var (
		_dag  = true
		_bd   int
		_eaa  *code
		_ggf  int
		_eedb error
	)
_bfae:
	for _bd < _bfb {
	_efg:
		for {
			if _dag {
				_eaa, _eedb = _eaf.uncompressGetCode(_ceg._ffd)
				if _eedb != nil {
					return 0, _eedb
				}
			} else {
				_eaa, _eedb = _eaf.uncompressGetCode(_ceg._ab)
				if _eedb != nil {
					return 0, _eedb
				}
			}
			_eaf._ffdd += _eaa._cd
			if _eaa._f < 0 {
				break _bfae
			}
			_bd += _eaa._f
			if _eaa._f < 64 {
				_dag = !_dag
				_bcf[_ggf] = _bd
				_ggf++
				break _efg
			}
		}
	}
	if _bcf[_ggf] != _bfb {
		_bcf[_ggf] = _bfb
	}
	_dga := EOL
	if _eaa != nil && _eaa._f != EOL {
		_dga = _ggf
	}
	return _dga, nil
}
func _eeg(_edag *_e.Reader) (*runData, error) {
	_cab := &runData{_edfb: _edag, _ffdd: 0, _fad: 1}
	_abe := _ea(_ba(_cca, int(_edag.Length())), _cgc)
	_cab._agfc = make([]byte, _abe)
	if _feb := _cab.fillBuffer(0); _feb != nil {
		if _feb == _b.EOF {
			_cab._agfc = make([]byte, 10)
			_da.Log.Debug("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076", _feb)
		} else {
			return nil, _feb
		}
	}
	return _cab, nil
}

const (
	_bb mmrCode = iota
	_baa
	_fb
	_gb
	_cdf
	_edb
	_bf
	_dg
	_egg
	_cc
	_be
)

var (
	_dc  = [][3]int{{4, 0x1, int(_bb)}, {3, 0x1, int(_baa)}, {1, 0x1, int(_fb)}, {3, 0x3, int(_gb)}, {6, 0x3, int(_cdf)}, {7, 0x3, int(_edb)}, {3, 0x2, int(_bf)}, {6, 0x2, int(_dg)}, {7, 0x2, int(_egg)}, {10, 0xf, int(_cc)}, {12, 0xf, int(_be)}, {12, 0x1, int(EOL)}}
	_ge  = [][3]int{{4, 0x07, 2}, {4, 0x08, 3}, {4, 0x0B, 4}, {4, 0x0C, 5}, {4, 0x0E, 6}, {4, 0x0F, 7}, {5, 0x12, 128}, {5, 0x13, 8}, {5, 0x14, 9}, {5, 0x1B, 64}, {5, 0x07, 10}, {5, 0x08, 11}, {6, 0x17, 192}, {6, 0x18, 1664}, {6, 0x2A, 16}, {6, 0x2B, 17}, {6, 0x03, 13}, {6, 0x34, 14}, {6, 0x35, 15}, {6, 0x07, 1}, {6, 0x08, 12}, {7, 0x13, 26}, {7, 0x17, 21}, {7, 0x18, 28}, {7, 0x24, 27}, {7, 0x27, 18}, {7, 0x28, 24}, {7, 0x2B, 25}, {7, 0x03, 22}, {7, 0x37, 256}, {7, 0x04, 23}, {7, 0x08, 20}, {7, 0xC, 19}, {8, 0x12, 33}, {8, 0x13, 34}, {8, 0x14, 35}, {8, 0x15, 36}, {8, 0x16, 37}, {8, 0x17, 38}, {8, 0x1A, 31}, {8, 0x1B, 32}, {8, 0x02, 29}, {8, 0x24, 53}, {8, 0x25, 54}, {8, 0x28, 39}, {8, 0x29, 40}, {8, 0x2A, 41}, {8, 0x2B, 42}, {8, 0x2C, 43}, {8, 0x2D, 44}, {8, 0x03, 30}, {8, 0x32, 61}, {8, 0x33, 62}, {8, 0x34, 63}, {8, 0x35, 0}, {8, 0x36, 320}, {8, 0x37, 384}, {8, 0x04, 45}, {8, 0x4A, 59}, {8, 0x4B, 60}, {8, 0x5, 46}, {8, 0x52, 49}, {8, 0x53, 50}, {8, 0x54, 51}, {8, 0x55, 52}, {8, 0x58, 55}, {8, 0x59, 56}, {8, 0x5A, 57}, {8, 0x5B, 58}, {8, 0x64, 448}, {8, 0x65, 512}, {8, 0x67, 640}, {8, 0x68, 576}, {8, 0x0A, 47}, {8, 0x0B, 48}, {9, 0x01, _fd}, {9, 0x98, 1472}, {9, 0x99, 1536}, {9, 0x9A, 1600}, {9, 0x9B, 1728}, {9, 0xCC, 704}, {9, 0xCD, 768}, {9, 0xD2, 832}, {9, 0xD3, 896}, {9, 0xD4, 960}, {9, 0xD5, 1024}, {9, 0xD6, 1088}, {9, 0xD7, 1152}, {9, 0xD8, 1216}, {9, 0xD9, 1280}, {9, 0xDA, 1344}, {9, 0xDB, 1408}, {10, 0x01, _fd}, {11, 0x01, _fd}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}}
	_deg = [][3]int{{2, 0x02, 3}, {2, 0x03, 2}, {3, 0x02, 1}, {3, 0x03, 4}, {4, 0x02, 6}, {4, 0x03, 5}, {5, 0x03, 7}, {6, 0x04, 9}, {6, 0x05, 8}, {7, 0x04, 10}, {7, 0x05, 11}, {7, 0x07, 12}, {8, 0x04, 13}, {8, 0x07, 14}, {9, 0x01, _fd}, {9, 0x18, 15}, {10, 0x01, _fd}, {10, 0x17, 16}, {10, 0x18, 17}, {10, 0x37, 0}, {10, 0x08, 18}, {10, 0x0F, 64}, {11, 0x01, _fd}, {11, 0x17, 24}, {11, 0x18, 25}, {11, 0x28, 23}, {11, 0x37, 22}, {11, 0x67, 19}, {11, 0x68, 20}, {11, 0x6C, 21}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}, {12, 0x24, 52}, {12, 0x27, 55}, {12, 0x28, 56}, {12, 0x2B, 59}, {12, 0x2C, 60}, {12, 0x33, 320}, {12, 0x34, 384}, {12, 0x35, 448}, {12, 0x37, 53}, {12, 0x38, 54}, {12, 0x52, 50}, {12, 0x53, 51}, {12, 0x54, 44}, {12, 0x55, 45}, {12, 0x56, 46}, {12, 0x57, 47}, {12, 0x58, 57}, {12, 0x59, 58}, {12, 0x5A, 61}, {12, 0x5B, 256}, {12, 0x64, 48}, {12, 0x65, 49}, {12, 0x66, 62}, {12, 0x67, 63}, {12, 0x68, 30}, {12, 0x69, 31}, {12, 0x6A, 32}, {12, 0x6B, 33}, {12, 0x6C, 40}, {12, 0x6D, 41}, {12, 0xC8, 128}, {12, 0xC9, 192}, {12, 0xCA, 26}, {12, 0xCB, 27}, {12, 0xCC, 28}, {12, 0xCD, 29}, {12, 0xD2, 34}, {12, 0xD3, 35}, {12, 0xD4, 36}, {12, 0xD5, 37}, {12, 0xD6, 38}, {12, 0xD7, 39}, {12, 0xDA, 42}, {12, 0xDB, 43}, {13, 0x4A, 640}, {13, 0x4B, 704}, {13, 0x4C, 768}, {13, 0x4D, 832}, {13, 0x52, 1280}, {13, 0x53, 1344}, {13, 0x54, 1408}, {13, 0x55, 1472}, {13, 0x5A, 1536}, {13, 0x5B, 1600}, {13, 0x64, 1664}, {13, 0x65, 1728}, {13, 0x6C, 512}, {13, 0x6D, 576}, {13, 0x72, 896}, {13, 0x73, 960}, {13, 0x74, 1024}, {13, 0x75, 1088}, {13, 0x76, 1152}, {13, 0x77, 1216}}
)

func New(r *_e.Reader, width, height int, dataOffset, dataLength int64) (*Decoder, error) {
	_eed := &Decoder{_gf: width, _bfa: height}
	_ffb, _dcg := r.NewPartialReader(int(dataOffset), int(dataLength), false)
	if _dcg != nil {
		return nil, _dcg
	}
	_cg, _dcg := _eeg(_ffb)
	if _dcg != nil {
		return nil, _dcg
	}
	_, _dcg = r.Seek(_ffb.RelativePosition(), _b.SeekCurrent)
	if _dcg != nil {
		return nil, _dcg
	}
	_eed._eef = _cg
	if _dd := _eed.initTables(); _dd != nil {
		return nil, _dd
	}
	return _eed, nil
}
func (_deb *Decoder) fillBitmap(_ecd *_c.Bitmap, _ad int, _eggb []int, _add int) error {
	var _eagf byte
	_cde := 0
	_gfd := _ecd.GetByteIndex(_cde, _ad)
	for _aed := 0; _aed < _add; _aed++ {
		_ef := byte(1)
		_acb := _eggb[_aed]
		if (_aed & 1) == 0 {
			_ef = 0
		}
		for _cde < _acb {
			_eagf = (_eagf << 1) | _ef
			_cde++
			if (_cde & 7) == 0 {
				if _cfe := _ecd.SetByte(_gfd, _eagf); _cfe != nil {
					return _cfe
				}
				_gfd++
				_eagf = 0
			}
		}
	}
	if (_cde & 7) != 0 {
		_eagf <<= uint(8 - (_cde & 7))
		if _df := _ecd.SetByte(_gfd, _eagf); _df != nil {
			return _df
		}
	}
	return nil
}

type runData struct {
	_edfb *_e.Reader
	_ffdd int
	_fad  int
	_eabg int
	_agfc []byte
	_cef  int
	_dace int
}

func (_ag *code) String() string {
	return _d.Sprintf("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _ag._cd, _ag._eg, _ag._f)
}
func (_ae *Decoder) createLittleEndianTable(_eab [][3]int) ([]*code, error) {
	_gg := make([]*code, _eda+1)
	for _bc := 0; _bc < len(_eab); _bc++ {
		_dce := _ff(_eab[_bc])
		if _dce._cd <= _ee {
			_gfa := _ee - _dce._cd
			_cdd := _dce._eg << uint(_gfa)
			for _ec := (1 << uint(_gfa)) - 1; _ec >= 0; _ec-- {
				_ca := _cdd | _ec
				_gg[_ca] = _dce
			}
		} else {
			_bcd := _dce._eg >> uint(_dce._cd-_ee)
			if _gg[_bcd] == nil {
				var _cda = _ff([3]int{})
				_cda._ed = make([]*code, _bbfb+1)
				_gg[_bcd] = _cda
			}
			if _dce._cd <= _ee+_bbf {
				_eefe := _ee + _bbf - _dce._cd
				_db := (_dce._eg << uint(_eefe)) & _bbfb
				_gg[_bcd]._cb = true
				for _ac := (1 << uint(_eefe)) - 1; _ac >= 0; _ac-- {
					_gg[_bcd]._ed[_db|_ac] = _dce
				}
			} else {
				return nil, _bg.New("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072")
			}
		}
	}
	return _gg, nil
}

const (
	_cgc int  = 1024 << 7
	_cca int  = 3
	_efd uint = 24
)

func _ff(_ce [3]int) *code   { return &code{_cd: _ce[0], _eg: _ce[1], _f: _ce[2]} }
func (_gdb *runData) align() { _gdb._ffdd = ((_gdb._ffdd + 7) >> 3) << 3 }
func (_ccc *Decoder) uncompress2d(_fbb *runData, _aa []int, _dbb int, _dgb []int, _dcgc int) (int, error) {
	var (
		_ffe  int
		_edf  int
		_dff  int
		_degb = true
		_fef  error
		_abg  *code
	)
	_aa[_dbb] = _dcgc
	_aa[_dbb+1] = _dcgc
	_aa[_dbb+2] = _dcgc + 1
	_aa[_dbb+3] = _dcgc + 1
_agf:
	for _dff < _dcgc {
		_abg, _fef = _fbb.uncompressGetCode(_ccc._dee)
		if _fef != nil {
			return EOL, nil
		}
		if _abg == nil {
			_fbb._ffdd++
			break _agf
		}
		_fbb._ffdd += _abg._cd
		switch mmrCode(_abg._f) {
		case _fb:
			_dff = _aa[_ffe]
		case _gb:
			_dff = _aa[_ffe] + 1
		case _bf:
			_dff = _aa[_ffe] - 1
		case _baa:
			for {
				var _dagc []*code
				if _degb {
					_dagc = _ccc._ffd
				} else {
					_dagc = _ccc._ab
				}
				_abg, _fef = _fbb.uncompressGetCode(_dagc)
				if _fef != nil {
					return 0, _fef
				}
				if _abg == nil {
					break _agf
				}
				_fbb._ffdd += _abg._cd
				if _abg._f < 64 {
					if _abg._f < 0 {
						_dgb[_edf] = _dff
						_edf++
						_abg = nil
						break _agf
					}
					_dff += _abg._f
					_dgb[_edf] = _dff
					_edf++
					break
				}
				_dff += _abg._f
			}
			_dbf := _dff
		_ccd:
			for {
				var _gec []*code
				if !_degb {
					_gec = _ccc._ffd
				} else {
					_gec = _ccc._ab
				}
				_abg, _fef = _fbb.uncompressGetCode(_gec)
				if _fef != nil {
					return 0, _fef
				}
				if _abg == nil {
					break _agf
				}
				_fbb._ffdd += _abg._cd
				if _abg._f < 64 {
					if _abg._f < 0 {
						_dgb[_edf] = _dff
						_edf++
						break _agf
					}
					_dff += _abg._f
					if _dff < _dcgc || _dff != _dbf {
						_dgb[_edf] = _dff
						_edf++
					}
					break _ccd
				}
				_dff += _abg._f
			}
			for _dff < _dcgc && _aa[_ffe] <= _dff {
				_ffe += 2
			}
			continue _agf
		case _bb:
			_ffe++
			_dff = _aa[_ffe]
			_ffe++
			continue _agf
		case _cdf:
			_dff = _aa[_ffe] + 2
		case _dg:
			_dff = _aa[_ffe] - 2
		case _edb:
			_dff = _aa[_ffe] + 3
		case _egg:
			_dff = _aa[_ffe] - 3
		default:
			if _fbb._ffdd == 12 && _abg._f == EOL {
				_fbb._ffdd = 0
				if _, _fef = _ccc.uncompress1d(_fbb, _aa, _dcgc); _fef != nil {
					return 0, _fef
				}
				_fbb._ffdd++
				if _, _fef = _ccc.uncompress1d(_fbb, _dgb, _dcgc); _fef != nil {
					return 0, _fef
				}
				_dac, _abc := _ccc.uncompress1d(_fbb, _aa, _dcgc)
				if _abc != nil {
					return EOF, _abc
				}
				_fbb._ffdd++
				return _dac, nil
			}
			_dff = _dcgc
			continue _agf
		}
		if _dff <= _dcgc {
			_degb = !_degb
			_dgb[_edf] = _dff
			_edf++
			if _ffe > 0 {
				_ffe--
			} else {
				_ffe++
			}
			for _dff < _dcgc && _aa[_ffe] <= _dff {
				_ffe += 2
			}
		}
	}
	if _dgb[_edf] != _dcgc {
		_dgb[_edf] = _dcgc
	}
	if _abg == nil {
		return EOL, nil
	}
	return _edf, nil
}
func (_eag *Decoder) detectAndSkipEOL() error {
	for {
		_gbc, _edab := _eag._eef.uncompressGetCode(_eag._dee)
		if _edab != nil {
			return _edab
		}
		if _gbc != nil && _gbc._f == EOL {
			_eag._eef._ffdd += _gbc._cd
		} else {
			return nil
		}
	}
}
func _ba(_de, _g int) int {
	if _de < _g {
		return _g
	}
	return _de
}

const (
	EOF   = -3
	_fd   = -2
	EOL   = -1
	_ee   = 8
	_eda  = (1 << _ee) - 1
	_bbf  = 5
	_bbfb = (1 << _bbf) - 1
)

type code struct {
	_cd int
	_eg int
	_f  int
	_ed []*code
	_cb bool
}

func (_eb *Decoder) initTables() (_eggf error) {
	if _eb._ffd == nil {
		_eb._ffd, _eggf = _eb.createLittleEndianTable(_ge)
		if _eggf != nil {
			return
		}
		_eb._ab, _eggf = _eb.createLittleEndianTable(_deg)
		if _eggf != nil {
			return
		}
		_eb._dee, _eggf = _eb.createLittleEndianTable(_dc)
		if _eggf != nil {
			return
		}
	}
	return nil
}
func (_fff *Decoder) UncompressMMR() (_fe *_c.Bitmap, _fbf error) {
	_fe = _c.New(_fff._gf, _fff._bfa)
	_cgg := make([]int, _fe.Width+5)
	_bfe := make([]int, _fe.Width+5)
	_bfe[0] = _fe.Width
	_dcf := 1
	var _gd int
	for _cf := 0; _cf < _fe.Height; _cf++ {
		_gd, _fbf = _fff.uncompress2d(_fff._eef, _bfe, _dcf, _cgg, _fe.Width)
		if _fbf != nil {
			return nil, _fbf
		}
		if _gd == EOF {
			break
		}
		if _gd > 0 {
			_fbf = _fff.fillBitmap(_fe, _cf, _cgg, _gd)
			if _fbf != nil {
				return nil, _fbf
			}
		}
		_bfe, _cgg = _cgg, _bfe
		_dcf = _gd
	}
	if _fbf = _fff.detectAndSkipEOL(); _fbf != nil {
		return nil, _fbf
	}
	_fff._eef.align()
	return _fe, nil
}
func _ea(_agd, _fa int) int {
	if _agd > _fa {
		return _fa
	}
	return _agd
}
func (_ede *runData) uncompressGetCodeLittleEndian(_agg []*code) (*code, error) {
	_gfdc, _dbe := _ede.uncompressGetNextCodeLittleEndian()
	if _dbe != nil {
		_da.Log.Debug("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076", _dbe)
		return nil, _dbe
	}
	_gfdc &= 0xffffff
	_ddg := _gfdc >> (_efd - _ee)
	_ece := _agg[_ddg]
	if _ece != nil && _ece._cb {
		_ddg = (_gfdc >> (_efd - _ee - _bbf)) & _bbfb
		_ece = _ece._ed[_ddg]
	}
	return _ece, nil
}
func (_agc *runData) fillBuffer(_ecb int) error {
	_agc._cef = _ecb
	_, _fc := _agc._edfb.Seek(int64(_ecb), _b.SeekStart)
	if _fc != nil {
		if _fc == _b.EOF {
			_da.Log.Debug("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046")
			_agc._dace = -1
		} else {
			return _fc
		}
	}
	if _fc == nil {
		_agc._dace, _fc = _agc._edfb.Read(_agc._agfc)
		if _fc != nil {
			if _fc == _b.EOF {
				_da.Log.Trace("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046")
				_agc._dace = -1
			} else {
				return _fc
			}
		}
	}
	if _agc._dace > -1 && _agc._dace < 3 {
		for _agc._dace < 3 {
			_adde, _gc := _agc._edfb.ReadByte()
			if _gc != nil {
				if _gc == _b.EOF {
					_agc._agfc[_agc._dace] = 0
				} else {
					return _gc
				}
			} else {
				_agc._agfc[_agc._dace] = _adde & 0xFF
			}
			_agc._dace++
		}
	}
	_agc._dace -= 3
	if _agc._dace < 0 {
		_agc._agfc = make([]byte, len(_agc._agfc))
		_agc._dace = len(_agc._agfc) - 3
	}
	return nil
}

type Decoder struct {
	_gf, _bfa int
	_eef      *runData
	_ffd      []*code
	_ab       []*code
	_dee      []*code
}
type mmrCode int
