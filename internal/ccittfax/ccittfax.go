package ccittfax

import (
	_e "errors"
	_d "io"
	_g "math"

	_ge "github.com/bamzi/pdfext/internal/bitwise"
)

func _cbcc(_dcae [][]byte) [][]byte {
	_fgdb := make([]byte, len(_dcae[0]))
	for _gbgg := range _fgdb {
		_fgdb[_gbgg] = _bffa
	}
	_dcae = append(_dcae, []byte{})
	for _bdbb := len(_dcae) - 1; _bdbb > 0; _bdbb-- {
		_dcae[_bdbb] = _dcae[_bdbb-1]
	}
	_dcae[0] = _fgdb
	return _dcae
}
func init() {
	_a = &treeNode{_bcfg: true, _dce: _gd}
	_ba = &treeNode{_dce: _fa, _dbdde: _a}
	_ba._dgfg = _ba
	_f = &tree{_gdgd: &treeNode{}}
	if _eb := _f.fillWithNode(12, 0, _ba); _eb != nil {
		panic(_eb.Error())
	}
	if _db := _f.fillWithNode(12, 1, _a); _db != nil {
		panic(_db.Error())
	}
	_bf = &tree{_gdgd: &treeNode{}}
	for _de := 0; _de < len(_bag); _de++ {
		for _gb := 0; _gb < len(_bag[_de]); _gb++ {
			if _bd := _bf.fill(_de+2, int(_bag[_de][_gb]), int(_fee[_de][_gb])); _bd != nil {
				panic(_bd.Error())
			}
		}
	}
	if _df := _bf.fillWithNode(12, 0, _ba); _df != nil {
		panic(_df.Error())
	}
	if _gdf := _bf.fillWithNode(12, 1, _a); _gdf != nil {
		panic(_gdf.Error())
	}
	_gg = &tree{_gdgd: &treeNode{}}
	for _ca := 0; _ca < len(_ab); _ca++ {
		for _ce := 0; _ce < len(_ab[_ca]); _ce++ {
			if _gf := _gg.fill(_ca+4, int(_ab[_ca][_ce]), int(_ga[_ca][_ce])); _gf != nil {
				panic(_gf.Error())
			}
		}
	}
	if _ed := _gg.fillWithNode(12, 0, _ba); _ed != nil {
		panic(_ed.Error())
	}
	if _cb := _gg.fillWithNode(12, 1, _a); _cb != nil {
		panic(_cb.Error())
	}
	_c = &tree{_gdgd: &treeNode{}}
	if _gbf := _c.fill(4, 1, _ae); _gbf != nil {
		panic(_gbf.Error())
	}
	if _bdb := _c.fill(3, 1, _gge); _bdb != nil {
		panic(_bdb.Error())
	}
	if _fe := _c.fill(1, 1, 0); _fe != nil {
		panic(_fe.Error())
	}
	if _bfe := _c.fill(3, 3, 1); _bfe != nil {
		panic(_bfe.Error())
	}
	if _dd := _c.fill(6, 3, 2); _dd != nil {
		panic(_dd.Error())
	}
	if _ged := _c.fill(7, 3, 3); _ged != nil {
		panic(_ged.Error())
	}
	if _fg := _c.fill(3, 2, -1); _fg != nil {
		panic(_fg.Error())
	}
	if _fag := _c.fill(6, 2, -2); _fag != nil {
		panic(_fag.Error())
	}
	if _dg := _c.fill(7, 2, -3); _dg != nil {
		panic(_dg.Error())
	}
}
func _gbb(_dbf []byte, _ecc int) int {
	if _ecc >= len(_dbf) {
		return _ecc
	}
	if _ecc < -1 {
		_ecc = -1
	}
	var _agg byte
	if _ecc > -1 {
		_agg = _dbf[_ecc]
	} else {
		_agg = _bffa
	}
	_aee := _ecc + 1
	for _aee < len(_dbf) {
		if _dbf[_aee] != _agg {
			break
		}
		_aee++
	}
	return _aee
}

var _fee = [...][]uint16{{3, 2}, {1, 4}, {6, 5}, {7}, {9, 8}, {10, 11, 12}, {13, 14}, {15}, {16, 17, 0, 18, 64}, {24, 25, 23, 22, 19, 20, 21, 1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560, 52, 55, 56, 59, 60, 320, 384, 448, 53, 54, 50, 51, 44, 45, 46, 47, 57, 58, 61, 256, 48, 49, 62, 63, 30, 31, 32, 33, 40, 41, 128, 192, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 42, 43}, {640, 704, 768, 832, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 512, 576, 896, 960, 1024, 1088, 1152, 1216}}

func (_eeg *Decoder) decodeRun(_facd *tree) (int, error) {
	var _baa int
	_fcd := _facd._gdgd
	for {
		_baac, _cega := _eeg._gbfa.ReadBool()
		if _cega != nil {
			return 0, _cega
		}
		_fcd = _fcd.walk(_baac)
		if _fcd == nil {
			return 0, _e.New("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _fcd._bcfg {
			_baa += _fcd._dce
			switch {
			case _fcd._dce >= 64:
				_fcd = _facd._gdgd
			case _fcd._dce >= 0:
				return _baa, nil
			default:
				return _eeg._ddfd, nil
			}
		}
	}
}
func (_edc *Decoder) fetch() error {
	if _edc._acd == -1 {
		return nil
	}
	if _edc._cef < _edc._acd {
		return nil
	}
	_edc._acd = 0
	_fd := _edc.decodeRow()
	if _fd != nil {
		if !_e.Is(_fd, _d.EOF) {
			return _fd
		}
		if _edc._acd != 0 {
			return _fd
		}
		_edc._acd = -1
	}
	_edc._cef = 0
	return nil
}
func _cdd(_agaa, _ddg int) code {
	var _gbe code
	switch _ddg - _agaa {
	case -1:
		_gbe = _eg
	case -2:
		_gbe = _gc
	case -3:
		_gbe = _ddd
	case 0:
		_gbe = _ddf
	case 1:
		_gbe = _acf
	case 2:
		_gbe = _gec
	case 3:
		_gbe = _abe
	}
	return _gbe
}
func (_cec *Decoder) decodeRowType6() error {
	if _cec._gfd {
		_cec._gbfa.Align()
	}
	if _cec._egg {
		_cec._gbfa.Mark()
		_gbfd, _ffge := _cec.tryFetchEOL()
		if _ffge != nil {
			return _ffge
		}
		if _gbfd {
			_gbfd, _ffge = _cec.tryFetchEOL()
			if _ffge != nil {
				return _ffge
			}
			if _gbfd {
				return _d.EOF
			}
		}
		_cec._gbfa.Reset()
	}
	return _cec.decode2D()
}
func init() {
	_dgg = make(map[int]code)
	_dgg[0] = code{Code: 13<<8 | 3<<6, BitsWritten: 10}
	_dgg[1] = code{Code: 2 << (5 + 8), BitsWritten: 3}
	_dgg[2] = code{Code: 3 << (6 + 8), BitsWritten: 2}
	_dgg[3] = code{Code: 2 << (6 + 8), BitsWritten: 2}
	_dgg[4] = code{Code: 3 << (5 + 8), BitsWritten: 3}
	_dgg[5] = code{Code: 3 << (4 + 8), BitsWritten: 4}
	_dgg[6] = code{Code: 2 << (4 + 8), BitsWritten: 4}
	_dgg[7] = code{Code: 3 << (3 + 8), BitsWritten: 5}
	_dgg[8] = code{Code: 5 << (2 + 8), BitsWritten: 6}
	_dgg[9] = code{Code: 4 << (2 + 8), BitsWritten: 6}
	_dgg[10] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_dgg[11] = code{Code: 5 << (1 + 8), BitsWritten: 7}
	_dgg[12] = code{Code: 7 << (1 + 8), BitsWritten: 7}
	_dgg[13] = code{Code: 4 << 8, BitsWritten: 8}
	_dgg[14] = code{Code: 7 << 8, BitsWritten: 8}
	_dgg[15] = code{Code: 12 << 8, BitsWritten: 9}
	_dgg[16] = code{Code: 5<<8 | 3<<6, BitsWritten: 10}
	_dgg[17] = code{Code: 6 << 8, BitsWritten: 10}
	_dgg[18] = code{Code: 2 << 8, BitsWritten: 10}
	_dgg[19] = code{Code: 12<<8 | 7<<5, BitsWritten: 11}
	_dgg[20] = code{Code: 13 << 8, BitsWritten: 11}
	_dgg[21] = code{Code: 13<<8 | 4<<5, BitsWritten: 11}
	_dgg[22] = code{Code: 6<<8 | 7<<5, BitsWritten: 11}
	_dgg[23] = code{Code: 5 << 8, BitsWritten: 11}
	_dgg[24] = code{Code: 2<<8 | 7<<5, BitsWritten: 11}
	_dgg[25] = code{Code: 3 << 8, BitsWritten: 11}
	_dgg[26] = code{Code: 12<<8 | 10<<4, BitsWritten: 12}
	_dgg[27] = code{Code: 12<<8 | 11<<4, BitsWritten: 12}
	_dgg[28] = code{Code: 12<<8 | 12<<4, BitsWritten: 12}
	_dgg[29] = code{Code: 12<<8 | 13<<4, BitsWritten: 12}
	_dgg[30] = code{Code: 6<<8 | 8<<4, BitsWritten: 12}
	_dgg[31] = code{Code: 6<<8 | 9<<4, BitsWritten: 12}
	_dgg[32] = code{Code: 6<<8 | 10<<4, BitsWritten: 12}
	_dgg[33] = code{Code: 6<<8 | 11<<4, BitsWritten: 12}
	_dgg[34] = code{Code: 13<<8 | 2<<4, BitsWritten: 12}
	_dgg[35] = code{Code: 13<<8 | 3<<4, BitsWritten: 12}
	_dgg[36] = code{Code: 13<<8 | 4<<4, BitsWritten: 12}
	_dgg[37] = code{Code: 13<<8 | 5<<4, BitsWritten: 12}
	_dgg[38] = code{Code: 13<<8 | 6<<4, BitsWritten: 12}
	_dgg[39] = code{Code: 13<<8 | 7<<4, BitsWritten: 12}
	_dgg[40] = code{Code: 6<<8 | 12<<4, BitsWritten: 12}
	_dgg[41] = code{Code: 6<<8 | 13<<4, BitsWritten: 12}
	_dgg[42] = code{Code: 13<<8 | 10<<4, BitsWritten: 12}
	_dgg[43] = code{Code: 13<<8 | 11<<4, BitsWritten: 12}
	_dgg[44] = code{Code: 5<<8 | 4<<4, BitsWritten: 12}
	_dgg[45] = code{Code: 5<<8 | 5<<4, BitsWritten: 12}
	_dgg[46] = code{Code: 5<<8 | 6<<4, BitsWritten: 12}
	_dgg[47] = code{Code: 5<<8 | 7<<4, BitsWritten: 12}
	_dgg[48] = code{Code: 6<<8 | 4<<4, BitsWritten: 12}
	_dgg[49] = code{Code: 6<<8 | 5<<4, BitsWritten: 12}
	_dgg[50] = code{Code: 5<<8 | 2<<4, BitsWritten: 12}
	_dgg[51] = code{Code: 5<<8 | 3<<4, BitsWritten: 12}
	_dgg[52] = code{Code: 2<<8 | 4<<4, BitsWritten: 12}
	_dgg[53] = code{Code: 3<<8 | 7<<4, BitsWritten: 12}
	_dgg[54] = code{Code: 3<<8 | 8<<4, BitsWritten: 12}
	_dgg[55] = code{Code: 2<<8 | 7<<4, BitsWritten: 12}
	_dgg[56] = code{Code: 2<<8 | 8<<4, BitsWritten: 12}
	_dgg[57] = code{Code: 5<<8 | 8<<4, BitsWritten: 12}
	_dgg[58] = code{Code: 5<<8 | 9<<4, BitsWritten: 12}
	_dgg[59] = code{Code: 2<<8 | 11<<4, BitsWritten: 12}
	_dgg[60] = code{Code: 2<<8 | 12<<4, BitsWritten: 12}
	_dgg[61] = code{Code: 5<<8 | 10<<4, BitsWritten: 12}
	_dgg[62] = code{Code: 6<<8 | 6<<4, BitsWritten: 12}
	_dgg[63] = code{Code: 6<<8 | 7<<4, BitsWritten: 12}
	_fae = make(map[int]code)
	_fae[0] = code{Code: 53 << 8, BitsWritten: 8}
	_fae[1] = code{Code: 7 << (2 + 8), BitsWritten: 6}
	_fae[2] = code{Code: 7 << (4 + 8), BitsWritten: 4}
	_fae[3] = code{Code: 8 << (4 + 8), BitsWritten: 4}
	_fae[4] = code{Code: 11 << (4 + 8), BitsWritten: 4}
	_fae[5] = code{Code: 12 << (4 + 8), BitsWritten: 4}
	_fae[6] = code{Code: 14 << (4 + 8), BitsWritten: 4}
	_fae[7] = code{Code: 15 << (4 + 8), BitsWritten: 4}
	_fae[8] = code{Code: 19 << (3 + 8), BitsWritten: 5}
	_fae[9] = code{Code: 20 << (3 + 8), BitsWritten: 5}
	_fae[10] = code{Code: 7 << (3 + 8), BitsWritten: 5}
	_fae[11] = code{Code: 8 << (3 + 8), BitsWritten: 5}
	_fae[12] = code{Code: 8 << (2 + 8), BitsWritten: 6}
	_fae[13] = code{Code: 3 << (2 + 8), BitsWritten: 6}
	_fae[14] = code{Code: 52 << (2 + 8), BitsWritten: 6}
	_fae[15] = code{Code: 53 << (2 + 8), BitsWritten: 6}
	_fae[16] = code{Code: 42 << (2 + 8), BitsWritten: 6}
	_fae[17] = code{Code: 43 << (2 + 8), BitsWritten: 6}
	_fae[18] = code{Code: 39 << (1 + 8), BitsWritten: 7}
	_fae[19] = code{Code: 12 << (1 + 8), BitsWritten: 7}
	_fae[20] = code{Code: 8 << (1 + 8), BitsWritten: 7}
	_fae[21] = code{Code: 23 << (1 + 8), BitsWritten: 7}
	_fae[22] = code{Code: 3 << (1 + 8), BitsWritten: 7}
	_fae[23] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_fae[24] = code{Code: 40 << (1 + 8), BitsWritten: 7}
	_fae[25] = code{Code: 43 << (1 + 8), BitsWritten: 7}
	_fae[26] = code{Code: 19 << (1 + 8), BitsWritten: 7}
	_fae[27] = code{Code: 36 << (1 + 8), BitsWritten: 7}
	_fae[28] = code{Code: 24 << (1 + 8), BitsWritten: 7}
	_fae[29] = code{Code: 2 << 8, BitsWritten: 8}
	_fae[30] = code{Code: 3 << 8, BitsWritten: 8}
	_fae[31] = code{Code: 26 << 8, BitsWritten: 8}
	_fae[32] = code{Code: 27 << 8, BitsWritten: 8}
	_fae[33] = code{Code: 18 << 8, BitsWritten: 8}
	_fae[34] = code{Code: 19 << 8, BitsWritten: 8}
	_fae[35] = code{Code: 20 << 8, BitsWritten: 8}
	_fae[36] = code{Code: 21 << 8, BitsWritten: 8}
	_fae[37] = code{Code: 22 << 8, BitsWritten: 8}
	_fae[38] = code{Code: 23 << 8, BitsWritten: 8}
	_fae[39] = code{Code: 40 << 8, BitsWritten: 8}
	_fae[40] = code{Code: 41 << 8, BitsWritten: 8}
	_fae[41] = code{Code: 42 << 8, BitsWritten: 8}
	_fae[42] = code{Code: 43 << 8, BitsWritten: 8}
	_fae[43] = code{Code: 44 << 8, BitsWritten: 8}
	_fae[44] = code{Code: 45 << 8, BitsWritten: 8}
	_fae[45] = code{Code: 4 << 8, BitsWritten: 8}
	_fae[46] = code{Code: 5 << 8, BitsWritten: 8}
	_fae[47] = code{Code: 10 << 8, BitsWritten: 8}
	_fae[48] = code{Code: 11 << 8, BitsWritten: 8}
	_fae[49] = code{Code: 82 << 8, BitsWritten: 8}
	_fae[50] = code{Code: 83 << 8, BitsWritten: 8}
	_fae[51] = code{Code: 84 << 8, BitsWritten: 8}
	_fae[52] = code{Code: 85 << 8, BitsWritten: 8}
	_fae[53] = code{Code: 36 << 8, BitsWritten: 8}
	_fae[54] = code{Code: 37 << 8, BitsWritten: 8}
	_fae[55] = code{Code: 88 << 8, BitsWritten: 8}
	_fae[56] = code{Code: 89 << 8, BitsWritten: 8}
	_fae[57] = code{Code: 90 << 8, BitsWritten: 8}
	_fae[58] = code{Code: 91 << 8, BitsWritten: 8}
	_fae[59] = code{Code: 74 << 8, BitsWritten: 8}
	_fae[60] = code{Code: 75 << 8, BitsWritten: 8}
	_fae[61] = code{Code: 50 << 8, BitsWritten: 8}
	_fae[62] = code{Code: 51 << 8, BitsWritten: 8}
	_fae[63] = code{Code: 52 << 8, BitsWritten: 8}
	_bac = make(map[int]code)
	_bac[64] = code{Code: 3<<8 | 3<<6, BitsWritten: 10}
	_bac[128] = code{Code: 12<<8 | 8<<4, BitsWritten: 12}
	_bac[192] = code{Code: 12<<8 | 9<<4, BitsWritten: 12}
	_bac[256] = code{Code: 5<<8 | 11<<4, BitsWritten: 12}
	_bac[320] = code{Code: 3<<8 | 3<<4, BitsWritten: 12}
	_bac[384] = code{Code: 3<<8 | 4<<4, BitsWritten: 12}
	_bac[448] = code{Code: 3<<8 | 5<<4, BitsWritten: 12}
	_bac[512] = code{Code: 3<<8 | 12<<3, BitsWritten: 13}
	_bac[576] = code{Code: 3<<8 | 13<<3, BitsWritten: 13}
	_bac[640] = code{Code: 2<<8 | 10<<3, BitsWritten: 13}
	_bac[704] = code{Code: 2<<8 | 11<<3, BitsWritten: 13}
	_bac[768] = code{Code: 2<<8 | 12<<3, BitsWritten: 13}
	_bac[832] = code{Code: 2<<8 | 13<<3, BitsWritten: 13}
	_bac[896] = code{Code: 3<<8 | 18<<3, BitsWritten: 13}
	_bac[960] = code{Code: 3<<8 | 19<<3, BitsWritten: 13}
	_bac[1024] = code{Code: 3<<8 | 20<<3, BitsWritten: 13}
	_bac[1088] = code{Code: 3<<8 | 21<<3, BitsWritten: 13}
	_bac[1152] = code{Code: 3<<8 | 22<<3, BitsWritten: 13}
	_bac[1216] = code{Code: 119 << 3, BitsWritten: 13}
	_bac[1280] = code{Code: 2<<8 | 18<<3, BitsWritten: 13}
	_bac[1344] = code{Code: 2<<8 | 19<<3, BitsWritten: 13}
	_bac[1408] = code{Code: 2<<8 | 20<<3, BitsWritten: 13}
	_bac[1472] = code{Code: 2<<8 | 21<<3, BitsWritten: 13}
	_bac[1536] = code{Code: 2<<8 | 26<<3, BitsWritten: 13}
	_bac[1600] = code{Code: 2<<8 | 27<<3, BitsWritten: 13}
	_bac[1664] = code{Code: 3<<8 | 4<<3, BitsWritten: 13}
	_bac[1728] = code{Code: 3<<8 | 5<<3, BitsWritten: 13}
	_fgc = make(map[int]code)
	_fgc[64] = code{Code: 27 << (3 + 8), BitsWritten: 5}
	_fgc[128] = code{Code: 18 << (3 + 8), BitsWritten: 5}
	_fgc[192] = code{Code: 23 << (2 + 8), BitsWritten: 6}
	_fgc[256] = code{Code: 55 << (1 + 8), BitsWritten: 7}
	_fgc[320] = code{Code: 54 << 8, BitsWritten: 8}
	_fgc[384] = code{Code: 55 << 8, BitsWritten: 8}
	_fgc[448] = code{Code: 100 << 8, BitsWritten: 8}
	_fgc[512] = code{Code: 101 << 8, BitsWritten: 8}
	_fgc[576] = code{Code: 104 << 8, BitsWritten: 8}
	_fgc[640] = code{Code: 103 << 8, BitsWritten: 8}
	_fgc[704] = code{Code: 102 << 8, BitsWritten: 9}
	_fgc[768] = code{Code: 102<<8 | 1<<7, BitsWritten: 9}
	_fgc[832] = code{Code: 105 << 8, BitsWritten: 9}
	_fgc[896] = code{Code: 105<<8 | 1<<7, BitsWritten: 9}
	_fgc[960] = code{Code: 106 << 8, BitsWritten: 9}
	_fgc[1024] = code{Code: 106<<8 | 1<<7, BitsWritten: 9}
	_fgc[1088] = code{Code: 107 << 8, BitsWritten: 9}
	_fgc[1152] = code{Code: 107<<8 | 1<<7, BitsWritten: 9}
	_fgc[1216] = code{Code: 108 << 8, BitsWritten: 9}
	_fgc[1280] = code{Code: 108<<8 | 1<<7, BitsWritten: 9}
	_fgc[1344] = code{Code: 109 << 8, BitsWritten: 9}
	_fgc[1408] = code{Code: 109<<8 | 1<<7, BitsWritten: 9}
	_fgc[1472] = code{Code: 76 << 8, BitsWritten: 9}
	_fgc[1536] = code{Code: 76<<8 | 1<<7, BitsWritten: 9}
	_fgc[1600] = code{Code: 77 << 8, BitsWritten: 9}
	_fgc[1664] = code{Code: 24 << (2 + 8), BitsWritten: 6}
	_fgc[1728] = code{Code: 77<<8 | 1<<7, BitsWritten: 9}
	_fac = make(map[int]code)
	_fac[1792] = code{Code: 1 << 8, BitsWritten: 11}
	_fac[1856] = code{Code: 1<<8 | 4<<5, BitsWritten: 11}
	_fac[1920] = code{Code: 1<<8 | 5<<5, BitsWritten: 11}
	_fac[1984] = code{Code: 1<<8 | 2<<4, BitsWritten: 12}
	_fac[2048] = code{Code: 1<<8 | 3<<4, BitsWritten: 12}
	_fac[2112] = code{Code: 1<<8 | 4<<4, BitsWritten: 12}
	_fac[2176] = code{Code: 1<<8 | 5<<4, BitsWritten: 12}
	_fac[2240] = code{Code: 1<<8 | 6<<4, BitsWritten: 12}
	_fac[2304] = code{Code: 1<<8 | 7<<4, BitsWritten: 12}
	_fac[2368] = code{Code: 1<<8 | 12<<4, BitsWritten: 12}
	_fac[2432] = code{Code: 1<<8 | 13<<4, BitsWritten: 12}
	_fac[2496] = code{Code: 1<<8 | 14<<4, BitsWritten: 12}
	_fac[2560] = code{Code: 1<<8 | 15<<4, BitsWritten: 12}
	_dda = make(map[int]byte)
	_dda[0] = 0xFF
	_dda[1] = 0xFE
	_dda[2] = 0xFC
	_dda[3] = 0xF8
	_dda[4] = 0xF0
	_dda[5] = 0xE0
	_dda[6] = 0xC0
	_dda[7] = 0x80
	_dda[8] = 0x00
}
func _dbb(_dgf, _cfbf []byte, _defb int) int {
	_edff := _gbb(_cfbf, _defb)
	if _edff < len(_cfbf) && (_defb == -1 && _cfbf[_edff] == _bffa || _defb >= 0 && _defb < len(_dgf) && _dgf[_defb] == _cfbf[_edff] || _defb >= len(_dgf) && _dgf[_defb-1] != _cfbf[_edff]) {
		_edff = _gbb(_cfbf, _edff)
	}
	return _edff
}
func (_gdge *tree) fill(_bdeb, _gfbd, _eae int) error {
	_dbfe := _gdge._gdgd
	for _gdb := 0; _gdb < _bdeb; _gdb++ {
		_ccg := _bdeb - 1 - _gdb
		_dgc := ((_gfbd >> uint(_ccg)) & 1) != 0
		_bafd := _dbfe.walk(_dgc)
		if _bafd != nil {
			if _bafd._bcfg {
				return _e.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_dbfe = _bafd
			continue
		}
		_bafd = &treeNode{}
		if _gdb == _bdeb-1 {
			_bafd._dce = _eae
			_bafd._bcfg = true
		}
		if _gfbd == 0 {
			_bafd._ead = true
		}
		_dbfe.set(_dgc, _bafd)
		_dbfe = _bafd
	}
	return nil
}
func (_bcc *Decoder) tryFetchEOL() (bool, error) {
	_bfdad, _aaa := _bcc._gbfa.ReadBits(12)
	if _aaa != nil {
		return false, _aaa
	}
	return _bfdad == 0x1, nil
}
func NewDecoder(data []byte, options DecodeOptions) (*Decoder, error) {
	_ffb := &Decoder{_gbfa: _ge.NewReader(data), _ddfd: options.Columns, _cg: options.Rows, _aef: options.DamagedRowsBeforeError, _bfd: make([]byte, (options.Columns+7)/8), _abee: make([]int, options.Columns+2), _bb: make([]int, options.Columns+2), _gfd: options.EncodedByteAligned, _ea: options.BlackIsOne, _gcf: options.EndOfLine, _egg: options.EndOfBlock}
	switch {
	case options.K == 0:
		_ffb._dfd = _ff
		if len(data) < 20 {
			return nil, _e.New("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m")
		}
		_bde := data[:20]
		if _bde[0] != 0 || (_bde[1]>>4 != 1 && _bde[1] != 1) {
			_ffb._dfd = _dde
			_dge := (uint16(_bde[0])<<8 + uint16(_bde[1]&0xff)) >> 4
			for _aea := 12; _aea < 160; _aea++ {
				_dge = (_dge << 1) + uint16((_bde[_aea/8]>>uint16(7-(_aea%8)))&0x01)
				if _dge&0xfff == 1 {
					_ffb._dfd = _ff
					break
				}
			}
		}
	case options.K < 0:
		_ffb._dfd = _aa
	case options.K > 0:
		_ffb._dfd = _ff
		_ffb._bfb = true
	}
	switch _ffb._dfd {
	case _dde, _ff, _aa:
	default:
		return nil, _e.New("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065")
	}
	return _ffb, nil
}
func (_gacc *tree) fillWithNode(_fcg, _ddae int, _ecgb *treeNode) error {
	_bge := _gacc._gdgd
	for _aggg := 0; _aggg < _fcg; _aggg++ {
		_abb := uint(_fcg - 1 - _aggg)
		_efa := ((_ddae >> _abb) & 1) != 0
		_bea := _bge.walk(_efa)
		if _bea != nil {
			if _bea._bcfg {
				return _e.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_bge = _bea
			continue
		}
		if _aggg == _fcg-1 {
			_bea = _ecgb
		} else {
			_bea = &treeNode{}
		}
		if _ddae == 0 {
			_bea._ead = true
		}
		_bge.set(_efa, _bea)
		_bge = _bea
	}
	return nil
}
func (_efg *Decoder) decodeG32D() error {
	_efg._fgcd = _efg._ee
	_efg._bb, _efg._abee = _efg._abee, _efg._bb
	_cgg := true
	var (
		_gaa  bool
		_eee  int
		_abec error
	)
	_efg._ee = 0
_cf:
	for _eee < _efg._ddfd {
		_cfg := _c._gdgd
		for {
			_gaa, _abec = _efg._gbfa.ReadBool()
			if _abec != nil {
				return _abec
			}
			_cfg = _cfg.walk(_gaa)
			if _cfg == nil {
				continue _cf
			}
			if !_cfg._bcfg {
				continue
			}
			switch _cfg._dce {
			case _gge:
				var _ccf int
				if _cgg {
					_ccf, _abec = _efg.decodeRun(_gg)
				} else {
					_ccf, _abec = _efg.decodeRun(_bf)
				}
				if _abec != nil {
					return _abec
				}
				_eee += _ccf
				_efg._bb[_efg._ee] = _eee
				_efg._ee++
				if _cgg {
					_ccf, _abec = _efg.decodeRun(_bf)
				} else {
					_ccf, _abec = _efg.decodeRun(_gg)
				}
				if _abec != nil {
					return _abec
				}
				_eee += _ccf
				_efg._bb[_efg._ee] = _eee
				_efg._ee++
			case _ae:
				_gfda := _efg.getNextChangingElement(_eee, _cgg) + 1
				if _gfda >= _efg._fgcd {
					_eee = _efg._ddfd
				} else {
					_eee = _efg._abee[_gfda]
				}
			default:
				_fcb := _efg.getNextChangingElement(_eee, _cgg)
				if _fcb >= _efg._fgcd || _fcb == -1 {
					_eee = _efg._ddfd + _cfg._dce
				} else {
					_eee = _efg._abee[_fcb] + _cfg._dce
				}
				_efg._bb[_efg._ee] = _eee
				_efg._ee++
				_cgg = !_cgg
			}
			continue _cf
		}
	}
	return nil
}

var (
	_dgg  map[int]code
	_fae  map[int]code
	_bac  map[int]code
	_fgc  map[int]code
	_fac  map[int]code
	_dda  map[int]byte
	_ag   = code{Code: 1 << 4, BitsWritten: 12}
	_bff  = code{Code: 3 << 3, BitsWritten: 13}
	_ceg  = code{Code: 2 << 3, BitsWritten: 13}
	_cc   = code{Code: 1 << 12, BitsWritten: 4}
	_fgce = code{Code: 1 << 13, BitsWritten: 3}
	_ddf  = code{Code: 1 << 15, BitsWritten: 1}
	_eg   = code{Code: 3 << 13, BitsWritten: 3}
	_gc   = code{Code: 3 << 10, BitsWritten: 6}
	_ddd  = code{Code: 3 << 9, BitsWritten: 7}
	_acf  = code{Code: 2 << 13, BitsWritten: 3}
	_gec  = code{Code: 2 << 10, BitsWritten: 6}
	_abe  = code{Code: 2 << 9, BitsWritten: 7}
)

func (_gag *Decoder) getNextChangingElement(_efc int, _bdg bool) int {
	_abf := 0
	if !_bdg {
		_abf = 1
	}
	_abfb := int(uint32(_gag._bc)&0xFFFFFFFE) + _abf
	if _abfb > 2 {
		_abfb -= 2
	}
	if _efc == 0 {
		return _abfb
	}
	for _ccdd := _abfb; _ccdd < _gag._fgcd; _ccdd += 2 {
		if _efc < _gag._abee[_ccdd] {
			_gag._bc = _ccdd
			return _ccdd
		}
	}
	return -1
}

type code struct {
	Code        uint16
	BitsWritten int
}

func (_gede *Decoder) tryFetchRTC2D() (_fgd error) {
	_gede._gbfa.Mark()
	var _deb bool
	for _fdc := 0; _fdc < 5; _fdc++ {
		_deb, _fgd = _gede.tryFetchEOL1()
		if _fgd != nil {
			if _e.Is(_fgd, _d.EOF) {
				if _fdc == 0 {
					break
				}
				return _ebc
			}
		}
		if _deb {
			continue
		}
		if _fdc > 0 {
			return _ebc
		}
		break
	}
	if _deb {
		return _d.EOF
	}
	_gede._gbfa.Reset()
	return _fgd
}

var (
	_a   *treeNode
	_ba  *treeNode
	_bf  *tree
	_gg  *tree
	_f   *tree
	_c   *tree
	_gd  = -2000
	_fa  = -1000
	_ae  = -3000
	_gge = -4000
)

type tiffType int

func (_bdc *Decoder) Read(in []byte) (int, error) {
	if _bdc._ccd != nil {
		return 0, _bdc._ccd
	}
	_def := len(in)
	var (
		_bdd int
		_fc  int
	)
	for _def != 0 {
		if _bdc._cef >= _bdc._acd {
			if _dea := _bdc.fetch(); _dea != nil {
				_bdc._ccd = _dea
				return 0, _dea
			}
		}
		if _bdc._acd == -1 {
			return _bdd, _d.EOF
		}
		switch {
		case _def <= _bdc._acd-_bdc._cef:
			_af := _bdc._bfd[_bdc._cef : _bdc._cef+_def]
			for _, _dbg := range _af {
				if !_bdc._ea {
					_dbg = ^_dbg
				}
				in[_fc] = _dbg
				_fc++
			}
			_bdd += len(_af)
			_bdc._cef += len(_af)
			return _bdd, nil
		default:
			_ef := _bdc._bfd[_bdc._cef:]
			for _, _fgb := range _ef {
				if !_bdc._ea {
					_fgb = ^_fgb
				}
				in[_fc] = _fgb
				_fc++
			}
			_bdd += len(_ef)
			_bdc._cef += len(_ef)
			_def -= len(_ef)
		}
	}
	return _bdd, nil
}
func (_eab *Decoder) decode1D() error {
	var (
		_eaa int
		_age error
	)
	_bfdag := true
	_eab._ee = 0
	for {
		var _gdaf int
		if _bfdag {
			_gdaf, _age = _eab.decodeRun(_gg)
		} else {
			_gdaf, _age = _eab.decodeRun(_bf)
		}
		if _age != nil {
			return _age
		}
		_eaa += _gdaf
		_eab._bb[_eab._ee] = _eaa
		_eab._ee++
		_bfdag = !_bfdag
		if _eaa >= _eab._ddfd {
			break
		}
	}
	return nil
}
func (_bfed *Decoder) looseFetchEOL() (bool, error) {
	_gdg, _dfa := _bfed._gbfa.ReadBits(12)
	if _dfa != nil {
		return false, _dfa
	}
	switch _gdg {
	case 0x1:
		return true, nil
	case 0x0:
		for {
			_ggb, _ffcd := _bfed._gbfa.ReadBool()
			if _ffcd != nil {
				return false, _ffcd
			}
			if _ggb {
				return true, nil
			}
		}
	default:
		return false, nil
	}
}
func (_dca *Encoder) encodeG4(_beb [][]byte) []byte {
	_feb := make([][]byte, len(_beb))
	copy(_feb, _beb)
	_feb = _cbcc(_feb)
	var _cge []byte
	var _fcc int
	for _cdb := 1; _cdb < len(_feb); _cdb++ {
		if _dca.Rows > 0 && !_dca.EndOfBlock && _cdb == (_dca.Rows+1) {
			break
		}
		var _edcf []byte
		var _bga, _eca, _aae int
		_gecg := _fcc
		_cgfdc := -1
		for _cgfdc < len(_feb[_cdb]) {
			_bga = _gbb(_feb[_cdb], _cgfdc)
			_eca = _dbb(_feb[_cdb], _feb[_cdb-1], _cgfdc)
			_aae = _gbb(_feb[_cdb-1], _eca)
			if _aae < _bga {
				_edcf, _gecg = _edf(_edcf, _gecg, _cc)
				_cgfdc = _aae
			} else {
				if _g.Abs(float64(_eca-_bga)) > 3 {
					_edcf, _gecg, _cgfdc = _fcf(_feb[_cdb], _edcf, _gecg, _cgfdc, _bga)
				} else {
					_edcf, _gecg = _eggg(_edcf, _gecg, _bga, _eca)
					_cgfdc = _bga
				}
			}
		}
		_cge = _dca.appendEncodedRow(_cge, _edcf, _fcc)
		if _dca.EncodedByteAlign {
			_gecg = 0
		}
		_fcc = _gecg % 8
	}
	if _dca.EndOfBlock {
		_ced, _ := _bba(_fcc)
		_cge = _dca.appendEncodedRow(_cge, _ced, _fcc)
	}
	return _cge
}
func _gcfe(_fea []byte, _cgc int, _bcf code) ([]byte, int) {
	_gfbc := true
	var _gage []byte
	_gage, _cgc = _edf(nil, _cgc, _bcf)
	_ddb := 0
	var _acdc int
	for _ddb < len(_fea) {
		_acdc, _ddb = _gagf(_fea, _gfbc, _ddb)
		_gage, _cgc = _dgef(_gage, _cgc, _acdc, _gfbc)
		_gfbc = !_gfbc
	}
	return _gage, _cgc % 8
}

var _ga = [...][]uint16{{2, 3, 4, 5, 6, 7}, {128, 8, 9, 64, 10, 11}, {192, 1664, 16, 17, 13, 14, 15, 1, 12}, {26, 21, 28, 27, 18, 24, 25, 22, 256, 23, 20, 19}, {33, 34, 35, 36, 37, 38, 31, 32, 29, 53, 54, 39, 40, 41, 42, 43, 44, 30, 61, 62, 63, 0, 320, 384, 45, 59, 60, 46, 49, 50, 51, 52, 55, 56, 57, 58, 448, 512, 640, 576, 47, 48}, {1472, 1536, 1600, 1728, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216, 1280, 1344, 1408}, {}, {1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560}}

func (_gbc *treeNode) set(_da bool, _fga *treeNode) {
	if !_da {
		_gbc._dgfg = _fga
	} else {
		_gbc._dbdde = _fga
	}
}

type Encoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}

func (_bdba *Decoder) decodeRowType2() error {
	if _bdba._gfd {
		_bdba._gbfa.Align()
	}
	if _gfa := _bdba.decode1D(); _gfa != nil {
		return _gfa
	}
	return nil
}

type DecodeOptions struct {
	Columns                int
	Rows                   int
	K                      int
	EncodedByteAligned     bool
	BlackIsOne             bool
	EndOfBlock             bool
	EndOfLine              bool
	DamagedRowsBeforeError int
}

var _bag = [...][]uint16{{0x2, 0x3}, {0x2, 0x3}, {0x2, 0x3}, {0x3}, {0x4, 0x5}, {0x4, 0x5, 0x7}, {0x4, 0x7}, {0x18}, {0x17, 0x18, 0x37, 0x8, 0xf}, {0x17, 0x18, 0x28, 0x37, 0x67, 0x68, 0x6c, 0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f, 0x24, 0x27, 0x28, 0x2b, 0x2c, 0x33, 0x34, 0x35, 0x37, 0x38, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xda, 0xdb}, {0x4a, 0x4b, 0x4c, 0x4d, 0x52, 0x53, 0x54, 0x55, 0x5a, 0x5b, 0x64, 0x65, 0x6c, 0x6d, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77}}

func (_ddac tiffType) String() string {
	switch _ddac {
	case _dde:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065"
	case _ff:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034"
	case _aa:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036"
	default:
		return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d"
	}
}
func _ebb(_agc int) ([]byte, int) {
	var _ddfg []byte
	for _dcg := 0; _dcg < 6; _dcg++ {
		_ddfg, _agc = _edf(_ddfg, _agc, _ag)
	}
	return _ddfg, _agc % 8
}
func _baad(_fbc int) ([]byte, int) {
	var _cfb []byte
	for _fca := 0; _fca < 6; _fca++ {
		_cfb, _fbc = _edf(_cfb, _fbc, _bff)
	}
	return _cfb, _fbc % 8
}
func _eggg(_aaac []byte, _ebgg, _acdd, _geda int) ([]byte, int) {
	_ffab := _cdd(_acdd, _geda)
	_aaac, _ebgg = _edf(_aaac, _ebgg, _ffab)
	return _aaac, _ebgg
}
func _dgef(_gad []byte, _acfg int, _cab int, _cffg bool) ([]byte, int) {
	var (
		_gea  code
		_dbdd bool
	)
	for !_dbdd {
		_gea, _cab, _dbdd = _ebf(_cab, _cffg)
		_gad, _acfg = _edf(_gad, _acfg, _gea)
	}
	return _gad, _acfg
}
func (_bgb *Encoder) Encode(pixels [][]byte) []byte {
	if _bgb.BlackIs1 {
		_bffa = 0
		_bg = 1
	} else {
		_bffa = 1
		_bg = 0
	}
	if _bgb.K == 0 {
		return _bgb.encodeG31D(pixels)
	}
	if _bgb.K > 0 {
		return _bgb.encodeG32D(pixels)
	}
	if _bgb.K < 0 {
		return _bgb.encodeG4(pixels)
	}
	return nil
}

const (
	_ tiffType = iota
	_dde
	_ff
	_aa
)

type tree struct{ _gdgd *treeNode }

func _cbcf(_acffg []byte, _bbg int) ([]byte, int) { return _edf(_acffg, _bbg, _cc) }
func _fcf(_dbe, _cgce []byte, _gdc, _aga, _aac int) ([]byte, int, int) {
	_efee := _gbb(_dbe, _aac)
	_ggc := _aga >= 0 && _dbe[_aga] == _bffa || _aga == -1
	_cgce, _gdc = _edf(_cgce, _gdc, _fgce)
	var _ffbe int
	if _aga > -1 {
		_ffbe = _aac - _aga
	} else {
		_ffbe = _aac - _aga - 1
	}
	_cgce, _gdc = _dgef(_cgce, _gdc, _ffbe, _ggc)
	_ggc = !_ggc
	_agd := _efee - _aac
	_cgce, _gdc = _dgef(_cgce, _gdc, _agd, _ggc)
	_aga = _efee
	return _cgce, _gdc, _aga
}
func (_ffc *Decoder) decoderRowType41D() error {
	if _ffc._gfd {
		_ffc._gbfa.Align()
	}
	_ffc._gbfa.Mark()
	var (
		_faff bool
		_be   error
	)
	if _ffc._gcf {
		_faff, _be = _ffc.tryFetchEOL()
		if _be != nil {
			return _be
		}
		if !_faff {
			return _ec
		}
	} else {
		_faff, _be = _ffc.looseFetchEOL()
		if _be != nil {
			return _be
		}
	}
	if !_faff {
		_ffc._gbfa.Reset()
	}
	if _faff && _ffc._egg {
		_ffc._gbfa.Mark()
		for _abg := 0; _abg < 5; _abg++ {
			_faff, _be = _ffc.tryFetchEOL()
			if _be != nil {
				if _e.Is(_be, _d.EOF) {
					if _abg == 0 {
						break
					}
					return _ebc
				}
			}
			if _faff {
				continue
			}
			if _abg > 0 {
				return _ebc
			}
			break
		}
		if _faff {
			return _d.EOF
		}
		_ffc._gbfa.Reset()
	}
	if _be = _ffc.decode1D(); _be != nil {
		return _be
	}
	return nil
}
func _bba(_gbg int) ([]byte, int) {
	var _bdag []byte
	for _gdfg := 0; _gdfg < 2; _gdfg++ {
		_bdag, _gbg = _edf(_bdag, _gbg, _ag)
	}
	return _bdag, _gbg % 8
}
func _cdga(_gfad, _bcda []byte, _bgag int, _ddc bool) int {
	_dfaf := _gbb(_bcda, _bgag)
	if _dfaf < len(_bcda) && (_bgag == -1 && _bcda[_dfaf] == _bffa || _bgag >= 0 && _bgag < len(_gfad) && _gfad[_bgag] == _bcda[_dfaf] || _bgag >= len(_gfad) && _ddc && _bcda[_dfaf] == _bffa || _bgag >= len(_gfad) && !_ddc && _bcda[_dfaf] == _bg) {
		_dfaf = _gbb(_bcda, _dfaf)
	}
	return _dfaf
}

type treeNode struct {
	_dgfg  *treeNode
	_dbdde *treeNode
	_dce   int
	_ead   bool
	_bcfg  bool
}

func (_ecg *Encoder) encodeG31D(_edb [][]byte) []byte {
	var _cgfd []byte
	_afc := 0
	for _eabb := range _edb {
		if _ecg.Rows > 0 && !_ecg.EndOfBlock && _eabb == _ecg.Rows {
			break
		}
		_cbg, _afb := _gcfe(_edb[_eabb], _afc, _ag)
		_cgfd = _ecg.appendEncodedRow(_cgfd, _cbg, _afc)
		if _ecg.EncodedByteAlign {
			_afb = 0
		}
		_afc = _afb
	}
	if _ecg.EndOfBlock {
		_bee, _ := _ebb(_afc)
		_cgfd = _ecg.appendEncodedRow(_cgfd, _bee, _afc)
	}
	return _cgfd
}
func (_dc *Decoder) decode2D() error {
	_dc._fgcd = _dc._ee
	_dc._bb, _dc._abee = _dc._abee, _dc._bb
	_ffe := true
	var (
		_bcd bool
		_ecb int
		_fdb error
	)
	_dc._ee = 0
_cff:
	for _ecb < _dc._ddfd {
		_bbba := _c._gdgd
		for {
			_bcd, _fdb = _dc._gbfa.ReadBool()
			if _fdb != nil {
				return _fdb
			}
			_bbba = _bbba.walk(_bcd)
			if _bbba == nil {
				continue _cff
			}
			if !_bbba._bcfg {
				continue
			}
			switch _bbba._dce {
			case _gge:
				var _cbc int
				if _ffe {
					_cbc, _fdb = _dc.decodeRun(_gg)
				} else {
					_cbc, _fdb = _dc.decodeRun(_bf)
				}
				if _fdb != nil {
					return _fdb
				}
				_ecb += _cbc
				_dc._bb[_dc._ee] = _ecb
				_dc._ee++
				if _ffe {
					_cbc, _fdb = _dc.decodeRun(_bf)
				} else {
					_cbc, _fdb = _dc.decodeRun(_gg)
				}
				if _fdb != nil {
					return _fdb
				}
				_ecb += _cbc
				_dc._bb[_dc._ee] = _ecb
				_dc._ee++
			case _ae:
				_ccb := _dc.getNextChangingElement(_ecb, _ffe) + 1
				if _ccb >= _dc._fgcd {
					_ecb = _dc._ddfd
				} else {
					_ecb = _dc._abee[_ccb]
				}
			default:
				_fb := _dc.getNextChangingElement(_ecb, _ffe)
				if _fb >= _dc._fgcd || _fb == -1 {
					_ecb = _dc._ddfd + _bbba._dce
				} else {
					_ecb = _dc._abee[_fb] + _bbba._dce
				}
				_dc._bb[_dc._ee] = _ecb
				_dc._ee++
				_ffe = !_ffe
			}
			continue _cff
		}
	}
	return nil
}
func (_gfbg *Decoder) decodeRowType4() error {
	if !_gfbg._bfb {
		return _gfbg.decoderRowType41D()
	}
	if _gfbg._gfd {
		_gfbg._gbfa.Align()
	}
	_gfbg._gbfa.Mark()
	_cgd, _bab := _gfbg.tryFetchEOL()
	if _bab != nil {
		return _bab
	}
	if !_cgd && _gfbg._gcf {
		_gfbg._gda++
		if _gfbg._gda > _gfbg._aef {
			return _ec
		}
		_gfbg._gbfa.Reset()
	}
	if !_cgd {
		_gfbg._gbfa.Reset()
	}
	_gac, _bab := _gfbg._gbfa.ReadBool()
	if _bab != nil {
		return _bab
	}
	if _gac {
		if _cgd && _gfbg._egg {
			if _bab = _gfbg.tryFetchRTC2D(); _bab != nil {
				return _bab
			}
		}
		_bab = _gfbg.decode1D()
	} else {
		_bab = _gfbg.decode2D()
	}
	if _bab != nil {
		return _bab
	}
	return nil
}
func _ebf(_bce int, _fgda bool) (code, int, bool) {
	if _bce < 64 {
		if _fgda {
			return _fae[_bce], 0, true
		}
		return _dgg[_bce], 0, true
	}
	_abc := _bce / 64
	if _abc > 40 {
		return _fac[2560], _bce - 2560, false
	}
	if _abc > 27 {
		return _fac[_abc*64], _bce - _abc*64, false
	}
	if _fgda {
		return _fgc[_abc*64], _bce - _abc*64, false
	}
	return _bac[_abc*64], _bce - _abc*64, false
}
func (_faf *Decoder) decodeRow() (_cgf error) {
	if !_faf._egg && _faf._cg > 0 && _faf._cg == _faf._ebg {
		return _d.EOF
	}
	switch _faf._dfd {
	case _dde:
		_cgf = _faf.decodeRowType2()
	case _ff:
		_cgf = _faf.decodeRowType4()
	case _aa:
		_cgf = _faf.decodeRowType6()
	}
	if _cgf != nil {
		return _cgf
	}
	_gca := 0
	_gdd := true
	_faf._bc = 0
	for _bda := 0; _bda < _faf._ee; _bda++ {
		_faca := _faf._ddfd
		if _bda != _faf._ee {
			_faca = _faf._bb[_bda]
		}
		if _faca > _faf._ddfd {
			_faca = _faf._ddfd
		}
		_cga := _gca / 8
		for _gca%8 != 0 && _faca-_gca > 0 {
			var _acff byte
			if !_gdd {
				_acff = 1 << uint(7-(_gca%8))
			}
			_faf._bfd[_cga] |= _acff
			_gca++
		}
		if _gca%8 == 0 {
			_cga = _gca / 8
			var _bbb byte
			if !_gdd {
				_bbb = 0xff
			}
			for _faca-_gca > 7 {
				_faf._bfd[_cga] = _bbb
				_gca += 8
				_cga++
			}
		}
		for _faca-_gca > 0 {
			if _gca%8 == 0 {
				_faf._bfd[_cga] = 0
			}
			var _ffg byte
			if !_gdd {
				_ffg = 1 << uint(7-(_gca%8))
			}
			_faf._bfd[_cga] |= _ffg
			_gca++
		}
		_gdd = !_gdd
	}
	if _gca != _faf._ddfd {
		return _e.New("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068")
	}
	_faf._acd = (_gca + 7) / 8
	_faf._ebg++
	return nil
}

type Decoder struct {
	_ddfd int
	_cg   int
	_ebg  int
	_bfd  []byte
	_aef  int
	_bfb  bool
	_fgf  bool
	_geg  bool
	_ea   bool
	_gcf  bool
	_egg  bool
	_gfd  bool
	_acd  int
	_cef  int
	_abee []int
	_bb   []int
	_fgcd int
	_ee   int
	_gda  int
	_bc   int
	_gbfa *_ge.Reader
	_dfd  tiffType
	_ccd  error
}

func (_ggeb *Decoder) tryFetchEOL1() (bool, error) {
	_gdff, _ded := _ggeb._gbfa.ReadBits(13)
	if _ded != nil {
		return false, _ded
	}
	return _gdff == 0x3, nil
}
func _gagf(_agcc []byte, _efe bool, _ad int) (int, int) {
	_gcfa := 0
	for _ad < len(_agcc) {
		if _efe {
			if _agcc[_ad] != _bffa {
				break
			}
		} else {
			if _agcc[_ad] != _bg {
				break
			}
		}
		_gcfa++
		_ad++
	}
	return _gcfa, _ad
}

var _ab = [...][]uint16{{0x7, 0x8, 0xb, 0xc, 0xe, 0xf}, {0x12, 0x13, 0x14, 0x1b, 0x7, 0x8}, {0x17, 0x18, 0x2a, 0x2b, 0x3, 0x34, 0x35, 0x7, 0x8}, {0x13, 0x17, 0x18, 0x24, 0x27, 0x28, 0x2b, 0x3, 0x37, 0x4, 0x8, 0xc}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1a, 0x1b, 0x2, 0x24, 0x25, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x3, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x4, 0x4a, 0x4b, 0x5, 0x52, 0x53, 0x54, 0x55, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x67, 0x68, 0xa, 0xb}, {0x98, 0x99, 0x9a, 0x9b, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb}, {}, {0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f}}

func (_fead *treeNode) walk(_adf bool) *treeNode {
	if _adf {
		return _fead._dbdde
	}
	return _fead._dgfg
}
func (_gdafd *Encoder) encodeG32D(_eef [][]byte) []byte {
	var _gdgb []byte
	var _cd int
	for _dbd := 0; _dbd < len(_eef); _dbd += _gdafd.K {
		if _gdafd.Rows > 0 && !_gdafd.EndOfBlock && _dbd == _gdafd.Rows {
			break
		}
		_ffa, _aaf := _gcfe(_eef[_dbd], _cd, _bff)
		_gdgb = _gdafd.appendEncodedRow(_gdgb, _ffa, _cd)
		if _gdafd.EncodedByteAlign {
			_aaf = 0
		}
		_cd = _aaf
		for _bbf := _dbd + 1; _bbf < (_dbd+_gdafd.K) && _bbf < len(_eef); _bbf++ {
			if _gdafd.Rows > 0 && !_gdafd.EndOfBlock && _bbf == _gdafd.Rows {
				break
			}
			_gfc, _cdg := _edf(nil, _cd, _ceg)
			var _egf, _ede, _bbd int
			_cfgc := -1
			for _cfgc < len(_eef[_bbf]) {
				_egf = _gbb(_eef[_bbf], _cfgc)
				_ede = _dbb(_eef[_bbf], _eef[_bbf-1], _cfgc)
				_bbd = _gbb(_eef[_bbf-1], _ede)
				if _bbd < _egf {
					_gfc, _cdg = _cbcf(_gfc, _cdg)
					_cfgc = _bbd
				} else {
					if _g.Abs(float64(_ede-_egf)) > 3 {
						_gfc, _cdg, _cfgc = _fcf(_eef[_bbf], _gfc, _cdg, _cfgc, _egf)
					} else {
						_gfc, _cdg = _eggg(_gfc, _cdg, _egf, _ede)
						_cfgc = _egf
					}
				}
			}
			_gdgb = _gdafd.appendEncodedRow(_gdgb, _gfc, _cd)
			if _gdafd.EncodedByteAlign {
				_cdg = 0
			}
			_cd = _cdg % 8
		}
	}
	if _gdafd.EndOfBlock {
		_cad, _ := _baad(_cd)
		_gdgb = _gdafd.appendEncodedRow(_gdgb, _cad, _cd)
	}
	return _gdgb
}

var (
	_bffa byte = 1
	_bg   byte = 0
)
var (
	_ebc = _e.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043")
	_ec  = _e.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
)

func _edf(_ecbc []byte, _egb int, _faa code) ([]byte, int) {
	_cfff := 0
	for _cfff < _faa.BitsWritten {
		_deca := _egb / 8
		_eda := _egb % 8
		if _deca >= len(_ecbc) {
			_ecbc = append(_ecbc, 0)
		}
		_baaca := 8 - _eda
		_dgee := _faa.BitsWritten - _cfff
		if _baaca > _dgee {
			_baaca = _dgee
		}
		if _cfff < 8 {
			_ecbc[_deca] = _ecbc[_deca] | byte(_faa.Code>>uint(8+_eda-_cfff))&_dda[8-_baaca-_eda]
		} else {
			_ecbc[_deca] = _ecbc[_deca] | (byte(_faa.Code<<uint(_cfff-8))&_dda[8-_baaca])>>uint(_eda)
		}
		_egb += _baaca
		_cfff += _baaca
	}
	return _ecbc, _egb
}
func (_ace *Encoder) appendEncodedRow(_fgg, _baf []byte, _caf int) []byte {
	if len(_fgg) > 0 && _caf != 0 && !_ace.EncodedByteAlign {
		_fgg[len(_fgg)-1] = _fgg[len(_fgg)-1] | _baf[0]
		_fgg = append(_fgg, _baf[1:]...)
	} else {
		_fgg = append(_fgg, _baf...)
	}
	return _fgg
}
