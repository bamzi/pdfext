package arithmetic

import (
	_da "fmt"
	_c "io"
	_d "strings"

	_e "github.com/bamzi/pdfext/common"
	_bc "github.com/bamzi/pdfext/internal/bitwise"
	_bce "github.com/bamzi/pdfext/internal/jbig2/internal"
)

func New(r *_bc.Reader) (*Decoder, error) {
	_cc := &Decoder{_dd: r, ContextSize: []uint32{16, 13, 10, 10}, ReferedToContextSize: []uint32{13, 10}}
	if _fa := _cc.init(); _fa != nil {
		return nil, _fa
	}
	return _cc, nil
}
func (_fc *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		_ff int
		_ed = _f[stats.cx()][0]
		_ge = int32(stats.cx())
	)
	defer func() { _fc._bd++ }()
	_fc._dg -= _ed
	if (_fc._bf >> 16) < uint64(_ed) {
		_ff = _fc.lpsExchange(stats, _ge, _ed)
		if _be := _fc.renormalize(); _be != nil {
			return 0, _be
		}
	} else {
		_fc._bf -= uint64(_ed) << 16
		if (_fc._dg & 0x8000) == 0 {
			_ff = _fc.mpsExchange(stats, _ge)
			if _fg := _fc.renormalize(); _fg != nil {
				return 0, _fg
			}
		} else {
			_ff = int(stats.getMps())
		}
	}
	return _ff, nil
}

var (
	_f = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}
)

type DecoderStats struct {
	_gd  int32
	_cad int32
	_dgf []byte
	_eb  []byte
}

func (_cb *Decoder) DecodeInt(stats *DecoderStats) (int32, error) {
	var (
		_fgg, _gb     int32
		_fb, _cg, _bb int
		_bee          error
	)
	if stats == nil {
		stats = NewStats(512, 1)
	}
	_cb._g = 1
	_cg, _bee = _cb.decodeIntBit(stats)
	if _bee != nil {
		return 0, _bee
	}
	_fb, _bee = _cb.decodeIntBit(stats)
	if _bee != nil {
		return 0, _bee
	}
	if _fb == 1 {
		_fb, _bee = _cb.decodeIntBit(stats)
		if _bee != nil {
			return 0, _bee
		}
		if _fb == 1 {
			_fb, _bee = _cb.decodeIntBit(stats)
			if _bee != nil {
				return 0, _bee
			}
			if _fb == 1 {
				_fb, _bee = _cb.decodeIntBit(stats)
				if _bee != nil {
					return 0, _bee
				}
				if _fb == 1 {
					_fb, _bee = _cb.decodeIntBit(stats)
					if _bee != nil {
						return 0, _bee
					}
					if _fb == 1 {
						_bb = 32
						_gb = 4436
					} else {
						_bb = 12
						_gb = 340
					}
				} else {
					_bb = 8
					_gb = 84
				}
			} else {
				_bb = 6
				_gb = 20
			}
		} else {
			_bb = 4
			_gb = 4
		}
	} else {
		_bb = 2
		_gb = 0
	}
	for _bdg := 0; _bdg < _bb; _bdg++ {
		_fb, _bee = _cb.decodeIntBit(stats)
		if _bee != nil {
			return 0, _bee
		}
		_fgg = (_fgg << 1) | int32(_fb)
	}
	_fgg += _gb
	if _cg == 0 {
		return _fgg, nil
	} else if _cg == 1 && _fgg > 0 {
		return -_fgg, nil
	}
	return 0, _bce.ErrOOB
}
func (_cd *Decoder) readByte() error {
	if _cd._dd.AbsolutePosition() > _cd._cf {
		if _, _ad := _cd._dd.Seek(-1, _c.SeekCurrent); _ad != nil {
			return _ad
		}
	}
	_fd, _ca := _cd._dd.ReadByte()
	if _ca != nil {
		return _ca
	}
	_cd._bg = _fd
	if _cd._bg == 0xFF {
		_gc, _df := _cd._dd.ReadByte()
		if _df != nil {
			return _df
		}
		if _gc > 0x8F {
			_cd._bf += 0xFF00
			_cd._ce = 8
			if _, _daf := _cd._dd.Seek(-2, _c.SeekCurrent); _daf != nil {
				return _daf
			}
		} else {
			_cd._bf += uint64(_gc) << 9
			_cd._ce = 7
		}
	} else {
		_fd, _ca = _cd._dd.ReadByte()
		if _ca != nil {
			return _ca
		}
		_cd._bg = _fd
		_cd._bf += uint64(_cd._bg) << 8
		_cd._ce = 8
	}
	_cd._bf &= 0xFFFFFFFFFF
	return nil
}
func (_db *Decoder) init() error {
	_db._cf = _db._dd.AbsolutePosition()
	_dc, _dgb := _db._dd.ReadByte()
	if _dgb != nil {
		_e.Log.Debug("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076", _dgb)
		return _dgb
	}
	_db._bg = _dc
	_db._bf = uint64(_dc) << 16
	if _dgb = _db.readByte(); _dgb != nil {
		return _dgb
	}
	_db._bf <<= 7
	_db._ce -= 7
	_db._dg = 0x8000
	_db._bd++
	return nil
}
func (_bbg *DecoderStats) SetIndex(index int32) { _bbg._gd = index }

type Decoder struct {
	ContextSize          []uint32
	ReferedToContextSize []uint32
	_dd                  *_bc.Reader
	_bg                  uint8
	_bf                  uint64
	_dg                  uint32
	_g                   int64
	_ce                  int32
	_bd                  int32
	_cf                  int64
}

func (_ccd *Decoder) decodeIntBit(_dbc *DecoderStats) (int, error) {
	_dbc.SetIndex(int32(_ccd._g))
	_bdf, _fcc := _ccd.DecodeBit(_dbc)
	if _fcc != nil {
		_e.Log.Debug("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _fcc)
		return _bdf, _fcc
	}
	if _ccd._g < 256 {
		_ccd._g = ((_ccd._g << uint64(1)) | int64(_bdf)) & 0x1ff
	} else {
		_ccd._g = (((_ccd._g<<uint64(1) | int64(_bdf)) & 511) | 256) & 0x1ff
	}
	return _bdf, nil
}
func (_fga *DecoderStats) Reset() {
	for _egc := 0; _egc < len(_fga._dgf); _egc++ {
		_fga._dgf[_egc] = 0
		_fga._eb[_egc] = 0
	}
}
func (_beg *DecoderStats) Copy() *DecoderStats {
	_ebd := &DecoderStats{_cad: _beg._cad, _dgf: make([]byte, _beg._cad)}
	copy(_ebd._dgf, _beg._dgf)
	return _ebd
}
func (_gg *DecoderStats) cx() byte      { return _gg._dgf[_gg._gd] }
func (_gab *DecoderStats) getMps() byte { return _gab._eb[_gab._gd] }
func (_dfe *DecoderStats) toggleMps()   { _dfe._eb[_dfe._gd] ^= 1 }
func (_cca *DecoderStats) Overwrite(dNew *DecoderStats) {
	for _ffd := 0; _ffd < len(_cca._dgf); _ffd++ {
		_cca._dgf[_ffd] = dNew._dgf[_ffd]
		_cca._eb[_ffd] = dNew._eb[_ffd]
	}
}
func (_bbc *DecoderStats) setEntry(_ea int) { _bfa := byte(_ea & 0x7f); _bbc._dgf[_bbc._gd] = _bfa }
func (_bbd *DecoderStats) String() string {
	_ebe := &_d.Builder{}
	_ebe.WriteString(_da.Sprintf("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a", len(_bbd._dgf)))
	for _de, _cec := range _bbd._dgf {
		if _cec != 0 {
			_ebe.WriteString(_da.Sprintf("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a", _de, _cec))
		}
	}
	return _ebe.String()
}
func (_gcf *Decoder) renormalize() error {
	for {
		if _gcf._ce == 0 {
			if _dcf := _gcf.readByte(); _dcf != nil {
				return _dcf
			}
		}
		_gcf._dg <<= 1
		_gcf._bf <<= 1
		_gcf._ce--
		if (_gcf._dg & 0x8000) != 0 {
			break
		}
	}
	_gcf._bf &= 0xffffffff
	return nil
}
func (_fdb *Decoder) mpsExchange(_dae *DecoderStats, _dab int32) int {
	_fag := _dae._eb[_dae._gd]
	if _fdb._dg < _f[_dab][0] {
		if _f[_dab][3] == 1 {
			_dae.toggleMps()
		}
		_dae.setEntry(int(_f[_dab][2]))
		return int(1 - _fag)
	}
	_dae.setEntry(int(_f[_dab][1]))
	return int(_fag)
}
func (_ef *Decoder) lpsExchange(_bed *DecoderStats, _eg int32, _fcb uint32) int {
	_aa := _bed.getMps()
	if _ef._dg < _fcb {
		_bed.setEntry(int(_f[_eg][1]))
		_ef._dg = _fcb
		return int(_aa)
	}
	if _f[_eg][3] == 1 {
		_bed.toggleMps()
	}
	_bed.setEntry(int(_f[_eg][2]))
	_ef._dg = _fcb
	return int(1 - _aa)
}
func (_a *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {
	_a._g = 1
	var _edc uint64
	for _edc = 0; _edc < codeLen; _edc++ {
		stats.SetIndex(int32(_a._g))
		_ccf, _ga := _a.DecodeBit(stats)
		if _ga != nil {
			return 0, _ga
		}
		_a._g = (_a._g << 1) | int64(_ccf)
	}
	_ccg := _a._g - (1 << codeLen)
	return _ccg, nil
}
func NewStats(contextSize int32, index int32) *DecoderStats {
	return &DecoderStats{_gd: index, _cad: contextSize, _dgf: make([]byte, contextSize), _eb: make([]byte, contextSize)}
}
