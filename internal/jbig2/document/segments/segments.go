package segments

import (
	_d "encoding/binary"
	_aa "errors"
	_dg "fmt"
	_f "image"
	_fg "io"
	_c "math"
	_a "strings"
	_gf "time"

	_af "github.com/bamzi/pdfext/common"
	_ae "github.com/bamzi/pdfext/internal/bitwise"
	_fgc "github.com/bamzi/pdfext/internal/jbig2/basic"
	_df "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_b "github.com/bamzi/pdfext/internal/jbig2/decoder/arithmetic"
	_ge "github.com/bamzi/pdfext/internal/jbig2/decoder/huffman"
	_aad "github.com/bamzi/pdfext/internal/jbig2/decoder/mmr"
	_gc "github.com/bamzi/pdfext/internal/jbig2/encoder/arithmetic"
	_fe "github.com/bamzi/pdfext/internal/jbig2/errors"
	_e "github.com/bamzi/pdfext/internal/jbig2/internal"
)

func (_beec *GenericRefinementRegion) overrideAtTemplate0(_cad, _bfa, _faa, _dce, _gcg int) int {
	if _beec._ba[0] {
		_cad &= 0xfff7
		if _beec.GrAtY[0] == 0 && int(_beec.GrAtX[0]) >= -_gcg {
			_cad |= (_dce >> uint(7-(_gcg+int(_beec.GrAtX[0]))) & 0x1) << 3
		} else {
			_cad |= _beec.getPixel(_beec.RegionBitmap, _bfa+int(_beec.GrAtX[0]), _faa+int(_beec.GrAtY[0])) << 3
		}
	}
	if _beec._ba[1] {
		_cad &= 0xefff
		if _beec.GrAtY[1] == 0 && int(_beec.GrAtX[1]) >= -_gcg {
			_cad |= (_dce >> uint(7-(_gcg+int(_beec.GrAtX[1]))) & 0x1) << 12
		} else {
			_cad |= _beec.getPixel(_beec.ReferenceBitmap, _bfa+int(_beec.GrAtX[1]), _faa+int(_beec.GrAtY[1]))
		}
	}
	return _cad
}

var _ _ge.BasicTabler = &TableSegment{}

func (_bcec *PageInformationSegment) readWidthAndHeight() error {
	_abbg, _eacb := _bcec._fgee.ReadBits(32)
	if _eacb != nil {
		return _eacb
	}
	_bcec.PageBMWidth = int(_abbg & _c.MaxInt32)
	_abbg, _eacb = _bcec._fgee.ReadBits(32)
	if _eacb != nil {
		return _eacb
	}
	_bcec.PageBMHeight = int(_abbg & _c.MaxInt32)
	return nil
}
func (_dcef *SymbolDictionary) readRefinementAtPixels(_aabf int) error {
	_dcef.SdrATX = make([]int8, _aabf)
	_dcef.SdrATY = make([]int8, _aabf)
	var (
		_abef byte
		_bgba error
	)
	for _dbbe := 0; _dbbe < _aabf; _dbbe++ {
		_abef, _bgba = _dcef._bbda.ReadByte()
		if _bgba != nil {
			return _bgba
		}
		_dcef.SdrATX[_dbbe] = int8(_abef)
		_abef, _bgba = _dcef._bbda.ReadByte()
		if _bgba != nil {
			return _bgba
		}
		_dcef.SdrATY[_dbbe] = int8(_abef)
	}
	return nil
}
func (_bacb *HalftoneRegion) computeY(_fgde, _dbba int) int {
	return _bacb.shiftAndFill(int(_bacb.HGridY) + _fgde*int(_bacb.HRegionX) - _dbba*int(_bacb.HRegionY))
}
func (_afe *SymbolDictionary) setRetainedCodingContexts(_bdde *SymbolDictionary) {
	_afe._fdbf = _bdde._fdbf
	_afe.IsHuffmanEncoded = _bdde.IsHuffmanEncoded
	_afe.UseRefinementAggregation = _bdde.UseRefinementAggregation
	_afe.SdTemplate = _bdde.SdTemplate
	_afe.SdrTemplate = _bdde.SdrTemplate
	_afe.SdATX = _bdde.SdATX
	_afe.SdATY = _bdde.SdATY
	_afe.SdrATX = _bdde.SdrATX
	_afe.SdrATY = _bdde.SdrATY
	_afe._fbcd = _bdde._fbcd
}
func (_ga *GenericRefinementRegion) getGrReference() (*_df.Bitmap, error) {
	segments := _ga._ce.RTSegments
	if len(segments) == 0 {
		return nil, _aa.New("\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0065\u0078is\u0074\u0073")
	}
	_cag, _bfg := segments[0].GetSegmentData()
	if _bfg != nil {
		return nil, _bfg
	}
	_eda, _gae := _cag.(Regioner)
	if !_gae {
		return nil, _dg.Errorf("\u0072\u0065\u0066\u0065\u0072r\u0065\u0064\u0020\u0074\u006f\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _cag)
	}
	return _eda.GetRegionBitmap()
}
func (_eebb *TextRegion) blit(_cfadc *_df.Bitmap, _dbcg int64) error {
	if _eebb.IsTransposed == 0 && (_eebb.ReferenceCorner == 2 || _eebb.ReferenceCorner == 3) {
		_eebb._bcef += int64(_cfadc.Width - 1)
	} else if _eebb.IsTransposed == 1 && (_eebb.ReferenceCorner == 0 || _eebb.ReferenceCorner == 2) {
		_eebb._bcef += int64(_cfadc.Height - 1)
	}
	_dcge := _eebb._bcef
	if _eebb.IsTransposed == 1 {
		_dcge, _dbcg = _dbcg, _dcge
	}
	switch _eebb.ReferenceCorner {
	case 0:
		_dbcg -= int64(_cfadc.Height - 1)
	case 2:
		_dbcg -= int64(_cfadc.Height - 1)
		_dcge -= int64(_cfadc.Width - 1)
	case 3:
		_dcge -= int64(_cfadc.Width - 1)
	}
	_adea := _df.Blit(_cfadc, _eebb.RegionBitmap, int(_dcge), int(_dbcg), _eebb.CombinationOperator)
	if _adea != nil {
		return _adea
	}
	if _eebb.IsTransposed == 0 && (_eebb.ReferenceCorner == 0 || _eebb.ReferenceCorner == 1) {
		_eebb._bcef += int64(_cfadc.Width - 1)
	}
	if _eebb.IsTransposed == 1 && (_eebb.ReferenceCorner == 1 || _eebb.ReferenceCorner == 3) {
		_eebb._bcef += int64(_cfadc.Height - 1)
	}
	return nil
}
func (_gdad *SymbolDictionary) setCodingStatistics() error {
	if _gdad._agfa == nil {
		_gdad._agfa = _b.NewStats(512, 1)
	}
	if _gdad._cacfe == nil {
		_gdad._cacfe = _b.NewStats(512, 1)
	}
	if _gdad._cdab == nil {
		_gdad._cdab = _b.NewStats(512, 1)
	}
	if _gdad._aebeg == nil {
		_gdad._aebeg = _b.NewStats(512, 1)
	}
	if _gdad._bgdf == nil {
		_gdad._bgdf = _b.NewStats(512, 1)
	}
	if _gdad.UseRefinementAggregation && _gdad._ecc == nil {
		_gdad._ecc = _b.NewStats(1<<uint(_gdad._efgc), 1)
		_gdad._egdc = _b.NewStats(512, 1)
		_gdad._fcbb = _b.NewStats(512, 1)
	}
	if _gdad._fbcd == nil {
		_gdad._fbcd = _b.NewStats(65536, 1)
	}
	if _gdad._fdbf == nil {
		var _cedg error
		_gdad._fdbf, _cedg = _b.New(_gdad._bbda)
		if _cedg != nil {
			return _cedg
		}
	}
	return nil
}
func (_dbca *TextRegion) readRegionFlags() error {
	var (
		_gffd  int
		_fbba  uint64
		_gdgce error
	)
	_gffd, _gdgce = _dbca._deba.ReadBit()
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.SbrTemplate = int8(_gffd)
	_fbba, _gdgce = _dbca._deba.ReadBits(5)
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.SbDsOffset = int8(_fbba)
	if _dbca.SbDsOffset > 0x0f {
		_dbca.SbDsOffset -= 0x20
	}
	_gffd, _gdgce = _dbca._deba.ReadBit()
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.DefaultPixel = int8(_gffd)
	_fbba, _gdgce = _dbca._deba.ReadBits(2)
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.CombinationOperator = _df.CombinationOperator(int(_fbba) & 0x3)
	_gffd, _gdgce = _dbca._deba.ReadBit()
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.IsTransposed = int8(_gffd)
	_fbba, _gdgce = _dbca._deba.ReadBits(2)
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.ReferenceCorner = int16(_fbba) & 0x3
	_fbba, _gdgce = _dbca._deba.ReadBits(2)
	if _gdgce != nil {
		return _gdgce
	}
	_dbca.LogSBStrips = int16(_fbba) & 0x3
	_dbca.SbStrips = 1 << uint(_dbca.LogSBStrips)
	_gffd, _gdgce = _dbca._deba.ReadBit()
	if _gdgce != nil {
		return _gdgce
	}
	if _gffd == 1 {
		_dbca.UseRefinement = true
	}
	_gffd, _gdgce = _dbca._deba.ReadBit()
	if _gdgce != nil {
		return _gdgce
	}
	if _gffd == 1 {
		_dbca.IsHuffmanEncoded = true
	}
	return nil
}
func (_begcd *SymbolDictionary) parseHeader() (_ggbe error) {
	_af.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _ggbe != nil {
			_af.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u0020\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u005b\u0050\u0041\u0052SE\u002d\u0048\u0045A\u0044\u0045R\u005d\u0020\u0066\u0061\u0069\u006ce\u0064\u002e \u0025\u0076", _ggbe)
		} else {
			_af.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _ggbe = _begcd.readRegionFlags(); _ggbe != nil {
		return _ggbe
	}
	if _ggbe = _begcd.setAtPixels(); _ggbe != nil {
		return _ggbe
	}
	if _ggbe = _begcd.setRefinementAtPixels(); _ggbe != nil {
		return _ggbe
	}
	if _ggbe = _begcd.readNumberOfExportedSymbols(); _ggbe != nil {
		return _ggbe
	}
	if _ggbe = _begcd.readNumberOfNewSymbols(); _ggbe != nil {
		return _ggbe
	}
	if _ggbe = _begcd.setInSyms(); _ggbe != nil {
		return _ggbe
	}
	if _begcd._dgeg {
		_gdgdf := _begcd.Header.RTSegments
		for _baaab := len(_gdgdf) - 1; _baaab >= 0; _baaab-- {
			if _gdgdf[_baaab].Type == 0 {
				_gaae, _dddd := _gdgdf[_baaab].SegmentData.(*SymbolDictionary)
				if !_dddd {
					_ggbe = _dg.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074:\u0020\u0025\u0076\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020S\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _gdgdf[_baaab])
					return _ggbe
				}
				if _gaae._dgeg {
					_begcd.setRetainedCodingContexts(_gaae)
				}
				break
			}
		}
	}
	if _ggbe = _begcd.checkInput(); _ggbe != nil {
		return _ggbe
	}
	return nil
}
func (_fgeb *Header) GetSegmentData() (Segmenter, error) {
	var _fedg Segmenter
	if _fgeb.SegmentData != nil {
		_fedg = _fgeb.SegmentData
	}
	if _fedg == nil {
		_dfb, _dde := _dae[_fgeb.Type]
		if !_dde {
			return nil, _dg.Errorf("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u002f\u0020\u0025\u0064\u0020\u0063\u0072e\u0061t\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020", _fgeb.Type, _fgeb.Type)
		}
		_fedg = _dfb()
		_af.Log.Trace("\u005b\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u0074\u0020\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0030\u0034\u0058", _fgeb.SegmentNumber, _fgeb.SegmentDataStartOffset)
		_ddgc, _ffge := _fgeb.subInputReader()
		if _ffge != nil {
			return nil, _ffge
		}
		if _ffb := _fedg.Init(_fgeb, _ddgc); _ffb != nil {
			_af.Log.Debug("\u0049\u006e\u0069\u0074 \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076 \u0066o\u0072\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _ffb, _fedg)
			return nil, _ffb
		}
		_fgeb.SegmentData = _fedg
	}
	return _fedg, nil
}
func (_defca *PageInformationSegment) readResolution() error {
	_bcdb, _ccb := _defca._fgee.ReadBits(32)
	if _ccb != nil {
		return _ccb
	}
	_defca.ResolutionX = int(_bcdb & _c.MaxInt32)
	_bcdb, _ccb = _defca._fgee.ReadBits(32)
	if _ccb != nil {
		return _ccb
	}
	_defca.ResolutionY = int(_bcdb & _c.MaxInt32)
	return nil
}
func (_cceb *Header) readSegmentNumber(_fbad *_ae.Reader) error {
	const _eefb = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	_cbgf := make([]byte, 4)
	_, _fgab := _fbad.Read(_cbgf)
	if _fgab != nil {
		return _fe.Wrap(_fgab, _eefb, "")
	}
	_cceb.SegmentNumber = _d.BigEndian.Uint32(_cbgf)
	return nil
}
func (_cbba *PatternDictionary) GetDictionary() ([]*_df.Bitmap, error) {
	if _cbba.Patterns != nil {
		return _cbba.Patterns, nil
	}
	if !_cbba.IsMMREncoded {
		_cbba.setGbAtPixels()
	}
	_agee := NewGenericRegion(_cbba._fdcf)
	_agee.setParametersMMR(_cbba.IsMMREncoded, _cbba.DataOffset, _cbba.DataLength, uint32(_cbba.HdpHeight), (_cbba.GrayMax+1)*uint32(_cbba.HdpWidth), _cbba.HDTemplate, false, false, _cbba.GBAtX, _cbba.GBAtY)
	_cedf, _adf := _agee.GetRegionBitmap()
	if _adf != nil {
		return nil, _adf
	}
	if _adf = _cbba.extractPatterns(_cedf); _adf != nil {
		return nil, _adf
	}
	return _cbba.Patterns, nil
}
func (_bcbd *TableSegment) StreamReader() *_ae.Reader { return _bcbd._ebef }
func (_fdgge *TableSegment) HtOOB() int32             { return _fdgge._cadd }
func (_afd *GenericRefinementRegion) decodeTemplate(_fdb, _cggc, _gec, _bgc, _edcd, _ddg, _ggb, _fbf, _agc, _dcf int, _bcbb templater) (_cgd error) {
	var (
		_bfd, _cdgg, _fde, _beg, _db int16
		_eacf, _gbd, _gdgc, _dfe     int
		_cf                          byte
	)
	if _agc >= 1 && (_agc-1) < _afd.ReferenceBitmap.Height {
		_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - _bgc)
		if _cgd != nil {
			return _cgd
		}
		_eacf = int(_cf)
	}
	if _agc >= 0 && (_agc) < _afd.ReferenceBitmap.Height {
		_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf)
		if _cgd != nil {
			return _cgd
		}
		_gbd = int(_cf)
	}
	if _agc >= -1 && (_agc+1) < _afd.ReferenceBitmap.Height {
		_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf + _bgc)
		if _cgd != nil {
			return _cgd
		}
		_gdgc = int(_cf)
	}
	_dcf++
	if _fdb >= 1 {
		_cf, _cgd = _afd.RegionBitmap.GetByte(_fbf - _gec)
		if _cgd != nil {
			return _cgd
		}
		_dfe = int(_cf)
	}
	_fbf++
	_cda := _afd.ReferenceDX % 8
	_dgab := 6 + _cda
	_fdeg := _dcf % _bgc
	if _dgab >= 0 {
		if _dgab < 8 {
			_bfd = int16(_eacf>>uint(_dgab)) & 0x07
		}
		if _dgab < 8 {
			_cdgg = int16(_gbd>>uint(_dgab)) & 0x07
		}
		if _dgab < 8 {
			_fde = int16(_gdgc>>uint(_dgab)) & 0x07
		}
		if _dgab == 6 && _fdeg > 1 {
			if _agc >= 1 && (_agc-1) < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - _bgc - 2)
				if _cgd != nil {
					return _cgd
				}
				_bfd |= int16(_cf<<2) & 0x04
			}
			if _agc >= 0 && _agc < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - 2)
				if _cgd != nil {
					return _cgd
				}
				_cdgg |= int16(_cf<<2) & 0x04
			}
			if _agc >= -1 && _agc+1 < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf + _bgc - 2)
				if _cgd != nil {
					return _cgd
				}
				_fde |= int16(_cf<<2) & 0x04
			}
		}
		if _dgab == 0 {
			_eacf = 0
			_gbd = 0
			_gdgc = 0
			if _fdeg < _bgc-1 {
				if _agc >= 1 && _agc-1 < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - _bgc)
					if _cgd != nil {
						return _cgd
					}
					_eacf = int(_cf)
				}
				if _agc >= 0 && _agc < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf)
					if _cgd != nil {
						return _cgd
					}
					_gbd = int(_cf)
				}
				if _agc >= -1 && _agc+1 < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf + _bgc)
					if _cgd != nil {
						return _cgd
					}
					_gdgc = int(_cf)
				}
			}
			_dcf++
		}
	} else {
		_bfd = int16(_eacf<<1) & 0x07
		_cdgg = int16(_gbd<<1) & 0x07
		_fde = int16(_gdgc<<1) & 0x07
		_eacf = 0
		_gbd = 0
		_gdgc = 0
		if _fdeg < _bgc-1 {
			if _agc >= 1 && _agc-1 < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - _bgc)
				if _cgd != nil {
					return _cgd
				}
				_eacf = int(_cf)
			}
			if _agc >= 0 && _agc < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf)
				if _cgd != nil {
					return _cgd
				}
				_gbd = int(_cf)
			}
			if _agc >= -1 && _agc+1 < _afd.ReferenceBitmap.Height {
				_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf + _bgc)
				if _cgd != nil {
					return _cgd
				}
				_gdgc = int(_cf)
			}
			_dcf++
		}
		_bfd |= int16((_eacf >> 7) & 0x07)
		_cdgg |= int16((_gbd >> 7) & 0x07)
		_fde |= int16((_gdgc >> 7) & 0x07)
	}
	_beg = int16(_dfe >> 6)
	_db = 0
	_fab := (2 - _cda) % 8
	_eacf <<= uint(_fab)
	_gbd <<= uint(_fab)
	_gdgc <<= uint(_fab)
	_dfe <<= 2
	var _bee int
	for _bga := 0; _bga < _cggc; _bga++ {
		_bbd := _bga & 0x07
		_degc := _bcbb.form(_bfd, _cdgg, _fde, _beg, _db)
		if _afd._bf {
			_cf, _cgd = _afd.RegionBitmap.GetByte(_afd.RegionBitmap.GetByteIndex(_bga, _fdb))
			if _cgd != nil {
				return _cgd
			}
			_afd._ab.SetIndex(int32(_afd.overrideAtTemplate0(int(_degc), _bga, _fdb, int(_cf), _bbd)))
		} else {
			_afd._ab.SetIndex(int32(_degc))
		}
		_bee, _cgd = _afd._eg.DecodeBit(_afd._ab)
		if _cgd != nil {
			return _cgd
		}
		if _cgd = _afd.RegionBitmap.SetPixel(_bga, _fdb, byte(_bee)); _cgd != nil {
			return _cgd
		}
		_bfd = ((_bfd << 1) | 0x01&int16(_eacf>>7)) & 0x07
		_cdgg = ((_cdgg << 1) | 0x01&int16(_gbd>>7)) & 0x07
		_fde = ((_fde << 1) | 0x01&int16(_gdgc>>7)) & 0x07
		_beg = ((_beg << 1) | 0x01&int16(_dfe>>7)) & 0x07
		_db = int16(_bee)
		if (_bga-int(_afd.ReferenceDX))%8 == 5 {
			_eacf = 0
			_gbd = 0
			_gdgc = 0
			if ((_bga-int(_afd.ReferenceDX))/8)+1 < _afd.ReferenceBitmap.RowStride {
				if _agc >= 1 && (_agc-1) < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf - _bgc)
					if _cgd != nil {
						return _cgd
					}
					_eacf = int(_cf)
				}
				if _agc >= 0 && _agc < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf)
					if _cgd != nil {
						return _cgd
					}
					_gbd = int(_cf)
				}
				if _agc >= -1 && (_agc+1) < _afd.ReferenceBitmap.Height {
					_cf, _cgd = _afd.ReferenceBitmap.GetByte(_dcf + _bgc)
					if _cgd != nil {
						return _cgd
					}
					_gdgc = int(_cf)
				}
			}
			_dcf++
		} else {
			_eacf <<= 1
			_gbd <<= 1
			_gdgc <<= 1
		}
		if _bbd == 5 && _fdb >= 1 {
			if ((_bga >> 3) + 1) >= _afd.RegionBitmap.RowStride {
				_dfe = 0
			} else {
				_cf, _cgd = _afd.RegionBitmap.GetByte(_fbf - _gec)
				if _cgd != nil {
					return _cgd
				}
				_dfe = int(_cf)
			}
			_fbf++
		} else {
			_dfe <<= 1
		}
	}
	return nil
}
func (_dfge *GenericRefinementRegion) parseHeader() (_fdg error) {
	_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0048e\u0061\u0064e\u0072\u002e\u002e\u002e")
	_ega := _gf.Now()
	defer func() {
		if _fdg == nil {
			_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045G\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020h\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0069\u006e\u0069\u0073\u0068id\u0020\u0069\u006e\u003a\u0020\u0025\u0064\u0020\u006e\u0073", _gf.Since(_ega).Nanoseconds())
		} else {
			_af.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0073", _fdg)
		}
	}()
	if _fdg = _dfge.RegionInfo.parseHeader(); _fdg != nil {
		return _fdg
	}
	_, _fdg = _dfge._fd.ReadBits(6)
	if _fdg != nil {
		return _fdg
	}
	_dfge.IsTPGROn, _fdg = _dfge._fd.ReadBool()
	if _fdg != nil {
		return _fdg
	}
	var _eeg int
	_eeg, _fdg = _dfge._fd.ReadBit()
	if _fdg != nil {
		return _fdg
	}
	_dfge.TemplateID = int8(_eeg)
	switch _dfge.TemplateID {
	case 0:
		_dfge.Template = _dfge._fa
		if _fdg = _dfge.readAtPixels(); _fdg != nil {
			return _fdg
		}
	case 1:
		_dfge.Template = _dfge._gd
	}
	return nil
}
func (_de *GenericRefinementRegion) Init(header *Header, r *_ae.Reader) error {
	_de._ce = header
	_de._fd = r
	_de.RegionInfo = NewRegionSegment(r)
	return _de.parseHeader()
}
func (_gebc *HalftoneRegion) computeGrayScalePlanes(_dfcf []*_df.Bitmap, _caff int) ([][]int, error) {
	_aec := make([][]int, _gebc.HGridHeight)
	for _acdg := 0; _acdg < len(_aec); _acdg++ {
		_aec[_acdg] = make([]int, _gebc.HGridWidth)
	}
	for _abdga := 0; _abdga < int(_gebc.HGridHeight); _abdga++ {
		for _bcbgf := 0; _bcbgf < int(_gebc.HGridWidth); _bcbgf += 8 {
			var _bggd int
			if _eaeg := int(_gebc.HGridWidth) - _bcbgf; _eaeg > 8 {
				_bggd = 8
			} else {
				_bggd = _eaeg
			}
			_cba := _dfcf[0].GetByteIndex(_bcbgf, _abdga)
			for _fgdf := 0; _fgdf < _bggd; _fgdf++ {
				_adca := _fgdf + _bcbgf
				_aec[_abdga][_adca] = 0
				for _aebg := 0; _aebg < _caff; _aebg++ {
					_bebd, _cabe := _dfcf[_aebg].GetByte(_cba)
					if _cabe != nil {
						return nil, _cabe
					}
					_fcfa := _bebd >> uint(7-_adca&7)
					_bagc := _fcfa & 1
					_ffaf := 1 << uint(_aebg)
					_bed := int(_bagc) * _ffaf
					_aec[_abdga][_adca] += _bed
				}
			}
		}
	}
	return _aec, nil
}
func (_bcfb *SymbolDictionary) decodeHeightClassCollectiveBitmap(_beaa int64, _fdeb, _aedb uint32) (*_df.Bitmap, error) {
	if _beaa == 0 {
		_ddcf := _df.New(int(_aedb), int(_fdeb))
		var (
			_accf byte
			_gadd error
		)
		for _fcga := 0; _fcga < len(_ddcf.Data); _fcga++ {
			_accf, _gadd = _bcfb._bbda.ReadByte()
			if _gadd != nil {
				return nil, _gadd
			}
			if _gadd = _ddcf.SetByte(_fcga, _accf); _gadd != nil {
				return nil, _gadd
			}
		}
		return _ddcf, nil
	}
	if _bcfb._fgbe == nil {
		_bcfb._fgbe = NewGenericRegion(_bcfb._bbda)
	}
	_bcfb._fgbe.setParameters(true, _bcfb._bbda.AbsolutePosition(), _beaa, _fdeb, _aedb)
	_fdbb, _egcg := _bcfb._fgbe.GetRegionBitmap()
	if _egcg != nil {
		return nil, _egcg
	}
	return _fdbb, nil
}
func (_cac *GenericRefinementRegion) decodeOptimized(_ede, _ad, _dga, _fc, _aff, _cge, _deg int) error {
	var (
		_fca error
		_gfc int
		_gcf int
	)
	_fgb := _ede - int(_cac.ReferenceDY)
	if _gbb := int(-_cac.ReferenceDX); _gbb > 0 {
		_gfc = _gbb
	}
	_geg := _cac.ReferenceBitmap.GetByteIndex(_gfc, _fgb)
	if _cac.ReferenceDX > 0 {
		_gcf = int(_cac.ReferenceDX)
	}
	_gbe := _cac.RegionBitmap.GetByteIndex(_gcf, _ede)
	switch _cac.TemplateID {
	case 0:
		_fca = _cac.decodeTemplate(_ede, _ad, _dga, _fc, _aff, _cge, _deg, _gbe, _fgb, _geg, _cac._fa)
	case 1:
		_fca = _cac.decodeTemplate(_ede, _ad, _dga, _fc, _aff, _cge, _deg, _gbe, _fgb, _geg, _cac._gd)
	}
	return _fca
}
func (_cfb *GenericRegion) InitEncode(bm *_df.Bitmap, xLoc, yLoc, template int, duplicateLineRemoval bool) error {
	const _bcg = "\u0047e\u006e\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069\u006f\u006e.\u0049\u006e\u0069\u0074\u0045\u006e\u0063\u006f\u0064\u0065"
	if bm == nil {
		return _fe.Error(_bcg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if xLoc < 0 || yLoc < 0 {
		return _fe.Error(_bcg, "\u0078\u0020\u0061\u006e\u0064\u0020\u0079\u0020\u006c\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	_cfb.Bitmap = bm
	_cfb.GBTemplate = byte(template)
	switch _cfb.GBTemplate {
	case 0:
		_cfb.GBAtX = []int8{3, -3, 2, -2}
		_cfb.GBAtY = []int8{-1, -1, -2, -2}
	case 1:
		_cfb.GBAtX = []int8{3}
		_cfb.GBAtY = []int8{-1}
	case 2, 3:
		_cfb.GBAtX = []int8{2}
		_cfb.GBAtY = []int8{-1}
	default:
		return _fe.Errorf(_bcg, "\u0070\u0072o\u0076\u0069\u0064\u0065\u0064 \u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u002c\u0032\u002c\u0033\u007d", template)
	}
	_cfb.RegionSegment = &RegionSegment{BitmapHeight: uint32(bm.Height), BitmapWidth: uint32(bm.Width), XLocation: uint32(xLoc), YLocation: uint32(yLoc)}
	_cfb.IsTPGDon = duplicateLineRemoval
	return nil
}
func (_affb *SymbolDictionary) readNumberOfNewSymbols() error {
	_dbcb, _bacg := _affb._bbda.ReadBits(32)
	if _bacg != nil {
		return _bacg
	}
	_affb.NumberOfNewSymbols = uint32(_dbcb & _c.MaxUint32)
	return nil
}
func (_cbe *Header) subInputReader() (*_ae.Reader, error) {
	_cfee := int(_cbe.SegmentDataLength)
	if _cbe.SegmentDataLength == _c.MaxInt32 {
		_cfee = -1
	}
	return _cbe.Reader.NewPartialReader(int(_cbe.SegmentDataStartOffset), _cfee, false)
}
func (_deea *SymbolDictionary) InitEncode(symbols *_df.Bitmaps, symbolList []int, symbolMap map[int]int, unborderSymbols bool) error {
	const _bdce = "S\u0079\u006d\u0062\u006f\u006c\u0044i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002eI\u006e\u0069\u0074E\u006ec\u006f\u0064\u0065"
	_deea.SdATX = []int8{3, -3, 2, -2}
	_deea.SdATY = []int8{-1, -1, -2, -2}
	_deea._daaa = symbols
	_deea._face = make([]int, len(symbolList))
	copy(_deea._face, symbolList)
	if len(_deea._face) != _deea._daaa.Size() {
		return _fe.Error(_bdce, "s\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0061\u006e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u004ci\u0073\u0074\u0020\u006f\u0066\u0020\u0064\u0069\u0066\u0066er\u0065\u006e\u0074 \u0073i\u007a\u0065")
	}
	_deea.NumberOfNewSymbols = uint32(symbols.Size())
	_deea.NumberOfExportedSymbols = uint32(symbols.Size())
	_deea._feda = symbolMap
	_deea._afcf = unborderSymbols
	return nil
}

type HalftoneRegion struct {
	_baab                *_ae.Reader
	_bfed                *Header
	DataHeaderOffset     int64
	DataHeaderLength     int64
	DataOffset           int64
	DataLength           int64
	RegionSegment        *RegionSegment
	HDefaultPixel        int8
	CombinationOperator  _df.CombinationOperator
	HSkipEnabled         bool
	HTemplate            byte
	IsMMREncoded         bool
	HGridWidth           uint32
	HGridHeight          uint32
	HGridX               int32
	HGridY               int32
	HRegionX             uint16
	HRegionY             uint16
	HalftoneRegionBitmap *_df.Bitmap
	Patterns             []*_df.Bitmap
}

func (_abaa *SymbolDictionary) String() string {
	_gaaa := &_a.Builder{}
	_gaaa.WriteString("\n\u005b\u0053\u0059\u004dBO\u004c-\u0044\u0049\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u000a")
	_gaaa.WriteString(_dg.Sprintf("\u0009-\u0020S\u0064\u0072\u0054\u0065\u006dp\u006c\u0061t\u0065\u0020\u0025\u0076\u000a", _abaa.SdrTemplate))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0025\u0076\u000a", _abaa.SdTemplate))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006eg\u0043\u006f\u006e\u0074\u0065\u0078\u0074R\u0065\u0074\u0061\u0069\u006e\u0065\u0064\u0020\u0025\u0076\u000a", _abaa._acgc))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074\u0065\u0078\u0074\u0055\u0073\u0065\u0064 \u0025\u0076\u000a", _abaa._dgeg))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _abaa.SdHuffAggInstanceSelection))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053d\u0048\u0075\u0066\u0066\u0042\u004d\u0053\u0069\u007a\u0065S\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020%\u0076\u000a", _abaa.SdHuffBMSizeSelection))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _abaa.SdHuffDecodeWidthSelection))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020Sd\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064e\u0048e\u0069g\u0068t\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _abaa.SdHuffDecodeHeightSelection))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020U\u0073\u0065\u0052\u0065f\u0069\u006e\u0065\u006d\u0065\u006e\u0074A\u0067\u0067\u0072\u0065\u0067\u0061\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _abaa.UseRefinementAggregation))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020is\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0065\u0064\u0020\u0025\u0076\u000a", _abaa.IsHuffmanEncoded))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _abaa.SdATX))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _abaa.SdATY))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _abaa.SdrATX))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _abaa.SdrATY))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u004ff\u0045\u0078\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _abaa.NumberOfExportedSymbols))
	_gaaa.WriteString(_dg.Sprintf("\u0009-\u0020\u004e\u0075\u006db\u0065\u0072\u004f\u0066\u004ee\u0077S\u0079m\u0062\u006f\u006c\u0073\u0020\u0025\u0076\n", _abaa.NumberOfNewSymbols))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u004ff\u0049\u006d\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _abaa._eaab))
	_gaaa.WriteString(_dg.Sprintf("\u0009\u002d \u006e\u0075\u006d\u0062\u0065\u0072\u004f\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0064\u0053\u0079\u006d\u0062\u006f\u006c\u0073 %\u0076\u000a", _abaa._dacc))
	return _gaaa.String()
}
func (_daf *PageInformationSegment) Size() int { return 19 }
func (_aba *GenericRegion) readGBAtPixels(_edcg int) error {
	const _bbeb = "\u0072\u0065\u0061\u0064\u0047\u0042\u0041\u0074\u0050i\u0078\u0065\u006c\u0073"
	_aba.GBAtX = make([]int8, _edcg)
	_aba.GBAtY = make([]int8, _edcg)
	for _dcd := 0; _dcd < _edcg; _dcd++ {
		_fgcd, _fbdc := _aba._abe.ReadByte()
		if _fbdc != nil {
			return _fe.Wrapf(_fbdc, _bbeb, "\u0058\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _dcd)
		}
		_aba.GBAtX[_dcd] = int8(_fgcd)
		_fgcd, _fbdc = _aba._abe.ReadByte()
		if _fbdc != nil {
			return _fe.Wrapf(_fbdc, _bbeb, "\u0059\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _dcd)
		}
		_aba.GBAtY[_dcd] = int8(_fgcd)
	}
	return nil
}
func (_baa *template1) setIndex(_dgb *_b.DecoderStats) { _dgb.SetIndex(0x080) }
func (_gfcf *SymbolDictionary) readAtPixels(_gdaa int) error {
	_gfcf.SdATX = make([]int8, _gdaa)
	_gfcf.SdATY = make([]int8, _gdaa)
	var (
		_dffg byte
		_gecg error
	)
	for _beed := 0; _beed < _gdaa; _beed++ {
		_dffg, _gecg = _gfcf._bbda.ReadByte()
		if _gecg != nil {
			return _gecg
		}
		_gfcf.SdATX[_beed] = int8(_dffg)
		_dffg, _gecg = _gfcf._bbda.ReadByte()
		if _gecg != nil {
			return _gecg
		}
		_gfcf.SdATY[_beed] = int8(_dffg)
	}
	return nil
}
func (_fda *GenericRegion) Init(h *Header, r *_ae.Reader) error {
	_fda.RegionSegment = NewRegionSegment(r)
	_fda._abe = r
	return _fda.parseHeader()
}
func (_efg *GenericRegion) setParametersMMR(_bcdd bool, _dbcf, _ece int64, _gbbd, _fbdf uint32, _dac byte, _fafd, _abfe bool, _bcc, _abac []int8) {
	_efg.DataOffset = _dbcf
	_efg.DataLength = _ece
	_efg.RegionSegment = &RegionSegment{}
	_efg.RegionSegment.BitmapHeight = _gbbd
	_efg.RegionSegment.BitmapWidth = _fbdf
	_efg.GBTemplate = _dac
	_efg.IsMMREncoded = _bcdd
	_efg.IsTPGDon = _fafd
	_efg.GBAtX = _bcc
	_efg.GBAtY = _abac
}
func (_fcfaa *PageInformationSegment) Init(h *Header, r *_ae.Reader) (_adab error) {
	_fcfaa._fgee = r
	if _adab = _fcfaa.parseHeader(); _adab != nil {
		return _fe.Wrap(_adab, "P\u0061\u0067\u0065\u0049\u006e\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065g\u006d\u0065\u006et\u002eI\u006e\u0069\u0074", "")
	}
	return nil
}
func (_addd *PageInformationSegment) readCombinationOperatorOverrideAllowed() error {
	_cgcb, _dfdag := _addd._fgee.ReadBit()
	if _dfdag != nil {
		return _dfdag
	}
	if _cgcb == 1 {
		_addd._gced = true
	}
	return nil
}
func (_fbgeb *PageInformationSegment) checkInput() error {
	if _fbgeb.PageBMHeight == _c.MaxInt32 {
		if !_fbgeb.IsStripe {
			_af.Log.Debug("P\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0053\u0065\u0067me\u006e\u0074\u002e\u0049s\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0073\u0068ou\u006c\u0064 \u0062\u0065\u0020\u0074\u0072\u0075\u0065\u002e")
		}
	}
	return nil
}
func (_bgea *Header) readHeaderLength(_dfda *_ae.Reader, _bcfd int64) {
	_bgea.HeaderLength = _dfda.AbsolutePosition() - _bcfd
}
func (_adaa *TextRegion) decodeIds() (int64, error) {
	const _dfgd = "\u0064e\u0063\u006f\u0064\u0065\u0049\u0064s"
	if _adaa.IsHuffmanEncoded {
		if _adaa.SbHuffDS == 3 {
			if _adaa._ddea == nil {
				_ddfc := 0
				if _adaa.SbHuffFS == 3 {
					_ddfc++
				}
				var _ecbc error
				_adaa._ddea, _ecbc = _adaa.getUserTable(_ddfc)
				if _ecbc != nil {
					return 0, _fe.Wrap(_ecbc, _dfgd, "")
				}
			}
			return _adaa._ddea.Decode(_adaa._deba)
		}
		_cfac, _ccfa := _ge.GetStandardTable(8 + int(_adaa.SbHuffDS))
		if _ccfa != nil {
			return 0, _fe.Wrap(_ccfa, _dfgd, "")
		}
		return _cfac.Decode(_adaa._deba)
	}
	_aefc, _agfgd := _adaa._fagf.DecodeInt(_adaa._ccac)
	if _agfgd != nil {
		return 0, _fe.Wrap(_agfgd, _dfgd, "\u0063\u0078\u0049\u0041\u0044\u0053")
	}
	return int64(_aefc), nil
}
func (_cbef *SymbolDictionary) encodeRefinementATFlags(_gbeg _ae.BinaryWriter) (_eafcb int, _bcfda error) {
	const _cfbb = "\u0065\u006e\u0063od\u0065\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if !_cbef.UseRefinementAggregation || _cbef.SdrTemplate != 0 {
		return 0, nil
	}
	for _ccea := 0; _ccea < 2; _ccea++ {
		if _bcfda = _gbeg.WriteByte(byte(_cbef.SdrATX[_ccea])); _bcfda != nil {
			return _eafcb, _fe.Wrapf(_bcfda, _cfbb, "\u0053\u0064\u0072\u0041\u0054\u0058\u005b\u0025\u0064\u005d", _ccea)
		}
		_eafcb++
		if _bcfda = _gbeg.WriteByte(byte(_cbef.SdrATY[_ccea])); _bcfda != nil {
			return _eafcb, _fe.Wrapf(_bcfda, _cfbb, "\u0053\u0064\u0072\u0041\u0054\u0059\u005b\u0025\u0064\u005d", _ccea)
		}
		_eafcb++
	}
	return _eafcb, nil
}
func (_fgec *SymbolDictionary) setRefinementAtPixels() error {
	if !_fgec.UseRefinementAggregation || _fgec.SdrTemplate != 0 {
		return nil
	}
	if _gbaa := _fgec.readRefinementAtPixels(2); _gbaa != nil {
		return _gbaa
	}
	return nil
}
func (_dcfg *HalftoneRegion) checkInput() error {
	if _dcfg.IsMMREncoded {
		if _dcfg.HTemplate != 0 {
			_af.Log.Debug("\u0048\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030", _dcfg.HTemplate)
		}
		if _dcfg.HSkipEnabled {
			_af.Log.Debug("\u0048\u0053\u006b\u0069\u0070\u0045\u006e\u0061\u0062\u006c\u0065\u0064\u0020\u0030\u0020\u0025\u0076\u0020(\u0073\u0068\u006f\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u0029", _dcfg.HSkipEnabled)
		}
	}
	return nil
}

type Type int

func (_bcba *TextRegion) getSymbols() error {
	if _bcba.Header.RTSegments != nil {
		return _bcba.initSymbols()
	}
	return nil
}

type PageInformationSegment struct {
	_fgee             *_ae.Reader
	PageBMHeight      int
	PageBMWidth       int
	ResolutionX       int
	ResolutionY       int
	_gced             bool
	_bca              _df.CombinationOperator
	_fgaa             bool
	DefaultPixelValue uint8
	_cadb             bool
	IsLossless        bool
	IsStripe          bool
	MaxStripeSize     uint16
}

func (_ffce *GenericRegion) overrideAtTemplate0b(_gfb, _bfdf, _ebca, _fadd, _edgg, _abcfd int) int {
	if _ffce.GBAtOverride[0] {
		_gfb &= 0xFFFD
		if _ffce.GBAtY[0] == 0 && _ffce.GBAtX[0] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[0]&0x1)) << 1
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[0]), _ebca+int(_ffce.GBAtY[0]))) << 1
		}
	}
	if _ffce.GBAtOverride[1] {
		_gfb &= 0xDFFF
		if _ffce.GBAtY[1] == 0 && _ffce.GBAtX[1] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[1]&0x1)) << 13
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[1]), _ebca+int(_ffce.GBAtY[1]))) << 13
		}
	}
	if _ffce.GBAtOverride[2] {
		_gfb &= 0xFDFF
		if _ffce.GBAtY[2] == 0 && _ffce.GBAtX[2] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[2]&0x1)) << 9
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[2]), _ebca+int(_ffce.GBAtY[2]))) << 9
		}
	}
	if _ffce.GBAtOverride[3] {
		_gfb &= 0xBFFF
		if _ffce.GBAtY[3] == 0 && _ffce.GBAtX[3] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[3]&0x1)) << 14
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[3]), _ebca+int(_ffce.GBAtY[3]))) << 14
		}
	}
	if _ffce.GBAtOverride[4] {
		_gfb &= 0xEFFF
		if _ffce.GBAtY[4] == 0 && _ffce.GBAtX[4] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[4]&0x1)) << 12
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[4]), _ebca+int(_ffce.GBAtY[4]))) << 12
		}
	}
	if _ffce.GBAtOverride[5] {
		_gfb &= 0xFFDF
		if _ffce.GBAtY[5] == 0 && _ffce.GBAtX[5] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[5]&0x1)) << 5
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[5]), _ebca+int(_ffce.GBAtY[5]))) << 5
		}
	}
	if _ffce.GBAtOverride[6] {
		_gfb &= 0xFFFB
		if _ffce.GBAtY[6] == 0 && _ffce.GBAtX[6] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[6]&0x1)) << 2
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[6]), _ebca+int(_ffce.GBAtY[6]))) << 2
		}
	}
	if _ffce.GBAtOverride[7] {
		_gfb &= 0xFFF7
		if _ffce.GBAtY[7] == 0 && _ffce.GBAtX[7] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[7]&0x1)) << 3
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[7]), _ebca+int(_ffce.GBAtY[7]))) << 3
		}
	}
	if _ffce.GBAtOverride[8] {
		_gfb &= 0xF7FF
		if _ffce.GBAtY[8] == 0 && _ffce.GBAtX[8] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[8]&0x1)) << 11
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[8]), _ebca+int(_ffce.GBAtY[8]))) << 11
		}
	}
	if _ffce.GBAtOverride[9] {
		_gfb &= 0xFFEF
		if _ffce.GBAtY[9] == 0 && _ffce.GBAtX[9] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[9]&0x1)) << 4
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[9]), _ebca+int(_ffce.GBAtY[9]))) << 4
		}
	}
	if _ffce.GBAtOverride[10] {
		_gfb &= 0x7FFF
		if _ffce.GBAtY[10] == 0 && _ffce.GBAtX[10] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[10]&0x1)) << 15
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[10]), _ebca+int(_ffce.GBAtY[10]))) << 15
		}
	}
	if _ffce.GBAtOverride[11] {
		_gfb &= 0xFDFF
		if _ffce.GBAtY[11] == 0 && _ffce.GBAtX[11] >= -int8(_edgg) {
			_gfb |= (_fadd >> uint(int8(_abcfd)-_ffce.GBAtX[11]&0x1)) << 10
		} else {
			_gfb |= int(_ffce.getPixel(_bfdf+int(_ffce.GBAtX[11]), _ebca+int(_ffce.GBAtY[11]))) << 10
		}
	}
	return _gfb
}
func (_dbdg *TextRegion) setContexts(_aeag *_b.DecoderStats, _gabf *_b.DecoderStats, _cfde *_b.DecoderStats, _gebfb *_b.DecoderStats, _adcg *_b.DecoderStats, _fedaf *_b.DecoderStats, _adaf *_b.DecoderStats, _cbbg *_b.DecoderStats, _cccg *_b.DecoderStats, _aace *_b.DecoderStats) {
	_dbdg._bgce = _gabf
	_dbdg._deef = _cfde
	_dbdg._ccac = _gebfb
	_dbdg._cbgce = _adcg
	_dbdg._abbd = _adaf
	_dbdg._gbgf = _cbbg
	_dbdg._bebb = _fedaf
	_dbdg._gdgaf = _cccg
	_dbdg._fddb = _aace
	_dbdg._aaag = _aeag
}
func (_egag *TextRegion) decodeRdw() (int64, error) {
	const _ccbe = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064w"
	if _egag.IsHuffmanEncoded {
		if _egag.SbHuffRDWidth == 3 {
			if _egag._agab == nil {
				var (
					_dbdbe int
					_cbdfd error
				)
				if _egag.SbHuffFS == 3 {
					_dbdbe++
				}
				if _egag.SbHuffDS == 3 {
					_dbdbe++
				}
				if _egag.SbHuffDT == 3 {
					_dbdbe++
				}
				_egag._agab, _cbdfd = _egag.getUserTable(_dbdbe)
				if _cbdfd != nil {
					return 0, _fe.Wrap(_cbdfd, _ccbe, "")
				}
			}
			return _egag._agab.Decode(_egag._deba)
		}
		_adfa, _cage := _ge.GetStandardTable(14 + int(_egag.SbHuffRDWidth))
		if _cage != nil {
			return 0, _fe.Wrap(_cage, _ccbe, "")
		}
		return _adfa.Decode(_egag._deba)
	}
	_gbf, _agabc := _egag._fagf.DecodeInt(_egag._abbd)
	if _agabc != nil {
		return 0, _fe.Wrap(_agabc, _ccbe, "")
	}
	return int64(_gbf), nil
}
func (_agg *GenericRefinementRegion) getPixel(_fbc *_df.Bitmap, _bd, _def int) int {
	if _bd < 0 || _bd >= _fbc.Width {
		return 0
	}
	if _def < 0 || _def >= _fbc.Height {
		return 0
	}
	if _fbc.GetPixel(_bd, _def) {
		return 1
	}
	return 0
}
func (_fbcc *TextRegion) computeSymbolCodeLength() error {
	if _fbcc.IsHuffmanEncoded {
		return _fbcc.symbolIDCodeLengths()
	}
	_fbcc._adfd = int8(_c.Ceil(_c.Log(float64(_fbcc.NumberOfSymbols)) / _c.Log(2)))
	return nil
}

type Segmenter interface {
	Init(_dadd *Header, _ebga *_ae.Reader) error
}

func (_aafaf *TextRegion) GetRegionBitmap() (*_df.Bitmap, error) {
	if _aafaf.RegionBitmap != nil {
		return _aafaf.RegionBitmap, nil
	}
	if !_aafaf.IsHuffmanEncoded {
		if _edgf := _aafaf.setCodingStatistics(); _edgf != nil {
			return nil, _edgf
		}
	}
	if _gbcc := _aafaf.createRegionBitmap(); _gbcc != nil {
		return nil, _gbcc
	}
	if _gbdc := _aafaf.decodeSymbolInstances(); _gbdc != nil {
		return nil, _gbdc
	}
	return _aafaf.RegionBitmap, nil
}
func (_cecg *TextRegion) decodeDT() (_agca int64, _efeg error) {
	if _cecg.IsHuffmanEncoded {
		if _cecg.SbHuffDT == 3 {
			_agca, _efeg = _cecg._cafd.Decode(_cecg._deba)
			if _efeg != nil {
				return 0, _efeg
			}
		} else {
			var _dceb _ge.Tabler
			_dceb, _efeg = _ge.GetStandardTable(11 + int(_cecg.SbHuffDT))
			if _efeg != nil {
				return 0, _efeg
			}
			_agca, _efeg = _dceb.Decode(_cecg._deba)
			if _efeg != nil {
				return 0, _efeg
			}
		}
	} else {
		var _gcda int32
		_gcda, _efeg = _cecg._fagf.DecodeInt(_cecg._bgce)
		if _efeg != nil {
			return 0, _efeg
		}
		_agca = int64(_gcda)
	}
	_agca *= int64(_cecg.SbStrips)
	return _agca, nil
}
func (_edee *Header) parse(_dabf Documenter, _dgbg *_ae.Reader, _gedfc int64, _cceg OrganizationType) (_fgge error) {
	const _fbeg = "\u0070\u0061\u0072s\u0065"
	_af.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u005b\u0050A\u0052\u0053\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _fgge != nil {
			_af.Log.Trace("\u005b\u0053\u0045GM\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045R\u005d[\u0050A\u0052S\u0045\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _fgge)
		} else {
			_af.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052]\u005bP\u0041\u0052\u0053\u0045\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		}
	}()
	_, _fgge = _dgbg.Seek(_gedfc, _fg.SeekStart)
	if _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "\u0073\u0065\u0065\u006b\u0020\u0073\u0074\u0061\u0072\u0074")
	}
	if _fgge = _edee.readSegmentNumber(_dgbg); _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "")
	}
	if _fgge = _edee.readHeaderFlags(); _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "")
	}
	var _bafc uint64
	_bafc, _fgge = _edee.readNumberOfReferredToSegments(_dgbg)
	if _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "")
	}
	_edee.RTSNumbers, _fgge = _edee.readReferredToSegmentNumbers(_dgbg, int(_bafc))
	if _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "")
	}
	_fgge = _edee.readSegmentPageAssociation(_dabf, _dgbg, _bafc, _edee.RTSNumbers...)
	if _fgge != nil {
		return _fe.Wrap(_fgge, _fbeg, "")
	}
	if _edee.Type != TEndOfFile {
		if _fgge = _edee.readSegmentDataLength(_dgbg); _fgge != nil {
			return _fe.Wrap(_fgge, _fbeg, "")
		}
	}
	_edee.readDataStartOffset(_dgbg, _cceg)
	_edee.readHeaderLength(_dgbg, _gedfc)
	_af.Log.Trace("\u0025\u0073", _edee)
	return nil
}
func (_acb *Header) CleanSegmentData() {
	if _acb.SegmentData != nil {
		_acb.SegmentData = nil
	}
}
func (_gecdd *PatternDictionary) readPatternWidthAndHeight() error {
	_ccebf, _abdge := _gecdd._fdcf.ReadByte()
	if _abdge != nil {
		return _abdge
	}
	_gecdd.HdpWidth = _ccebf
	_ccebf, _abdge = _gecdd._fdcf.ReadByte()
	if _abdge != nil {
		return _abdge
	}
	_gecdd.HdpHeight = _ccebf
	return nil
}
func (_efcd *SymbolDictionary) Encode(w _ae.BinaryWriter) (_dceac int, _cdfb error) {
	const _fgaf = "\u0053\u0079\u006dbo\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	if _efcd == nil {
		return 0, _fe.Error(_fgaf, "\u0073\u0079m\u0062\u006f\u006c\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	if _dceac, _cdfb = _efcd.encodeFlags(w); _cdfb != nil {
		return _dceac, _fe.Wrap(_cdfb, _fgaf, "")
	}
	_gccc, _cdfb := _efcd.encodeATFlags(w)
	if _cdfb != nil {
		return _dceac, _fe.Wrap(_cdfb, _fgaf, "")
	}
	_dceac += _gccc
	if _gccc, _cdfb = _efcd.encodeRefinementATFlags(w); _cdfb != nil {
		return _dceac, _fe.Wrap(_cdfb, _fgaf, "")
	}
	_dceac += _gccc
	if _gccc, _cdfb = _efcd.encodeNumSyms(w); _cdfb != nil {
		return _dceac, _fe.Wrap(_cdfb, _fgaf, "")
	}
	_dceac += _gccc
	if _gccc, _cdfb = _efcd.encodeSymbols(w); _cdfb != nil {
		return _dceac, _fe.Wrap(_cdfb, _fgaf, "")
	}
	_dceac += _gccc
	return _dceac, nil
}
func (_cbfe *Header) readNumberOfReferredToSegments(_gag *_ae.Reader) (uint64, error) {
	const _ceea = "\u0072\u0065\u0061\u0064\u004e\u0075\u006d\u0062\u0065\u0072O\u0066\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	_dbac, _ddcg := _gag.ReadBits(3)
	if _ddcg != nil {
		return 0, _fe.Wrap(_ddcg, _ceea, "\u0063\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0072\u0074\u0073")
	}
	_dbac &= 0xf
	var _gabb []byte
	if _dbac <= 4 {
		_gabb = make([]byte, 5)
		for _fcaf := 0; _fcaf <= 4; _fcaf++ {
			_ddga, _eggb := _gag.ReadBit()
			if _eggb != nil {
				return 0, _fe.Wrap(_eggb, _ceea, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_gabb[_fcaf] = byte(_ddga)
		}
	} else {
		_dbac, _ddcg = _gag.ReadBits(29)
		if _ddcg != nil {
			return 0, _ddcg
		}
		_dbac &= _c.MaxInt32
		_dag := (_dbac + 8) >> 3
		_dag <<= 3
		_gabb = make([]byte, _dag)
		var _dad uint64
		for _dad = 0; _dad < _dag; _dad++ {
			_cgccd, _gdcc := _gag.ReadBit()
			if _gdcc != nil {
				return 0, _fe.Wrap(_gdcc, _ceea, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_gabb[_dad] = byte(_cgccd)
		}
	}
	return _dbac, nil
}
func (_bfge *HalftoneRegion) combineGrayscalePlanes(_eag []*_df.Bitmap, _gecb int) error {
	_dccb := 0
	for _fcd := 0; _fcd < _eag[_gecb].Height; _fcd++ {
		for _fddc := 0; _fddc < _eag[_gecb].Width; _fddc += 8 {
			_dfcd, _cdd := _eag[_gecb+1].GetByte(_dccb)
			if _cdd != nil {
				return _cdd
			}
			_abeb, _cdd := _eag[_gecb].GetByte(_dccb)
			if _cdd != nil {
				return _cdd
			}
			_cdd = _eag[_gecb].SetByte(_dccb, _df.CombineBytes(_abeb, _dfcd, _df.CmbOpXor))
			if _cdd != nil {
				return _cdd
			}
			_dccb++
		}
	}
	return nil
}
func (_dgfb *SymbolDictionary) decodeAggregate(_ggce, _ddee uint32) error {
	var (
		_dcb   int64
		_ggeee error
	)
	if _dgfb.IsHuffmanEncoded {
		_dcb, _ggeee = _dgfb.huffDecodeRefAggNInst()
		if _ggeee != nil {
			return _ggeee
		}
	} else {
		_gbegc, _ggag := _dgfb._fdbf.DecodeInt(_dgfb._aebeg)
		if _ggag != nil {
			return _ggag
		}
		_dcb = int64(_gbegc)
	}
	if _dcb > 1 {
		return _dgfb.decodeThroughTextRegion(_ggce, _ddee, uint32(_dcb))
	} else if _dcb == 1 {
		return _dgfb.decodeRefinedSymbol(_ggce, _ddee)
	}
	return nil
}
func (_fadcf *SymbolDictionary) getUserTable(_ffca int) (_ge.Tabler, error) {
	var _deac int
	for _, _bgeag := range _fadcf.Header.RTSegments {
		if _bgeag.Type == 53 {
			if _deac == _ffca {
				_fee, _ccfb := _bgeag.GetSegmentData()
				if _ccfb != nil {
					return nil, _ccfb
				}
				_dggg := _fee.(_ge.BasicTabler)
				return _ge.NewEncodedTable(_dggg)
			}
			_deac++
		}
	}
	return nil, nil
}
func (_ebfa *HalftoneRegion) GetRegionInfo() *RegionSegment { return _ebfa.RegionSegment }
func (_effa *Header) readDataStartOffset(_ced *_ae.Reader, _dfcfd OrganizationType) {
	if _dfcfd == OSequential {
		_effa.SegmentDataStartOffset = uint64(_ced.AbsolutePosition())
	}
}
func (_bff *GenericRefinementRegion) updateOverride() error {
	if _bff.GrAtX == nil || _bff.GrAtY == nil {
		return _aa.New("\u0041\u0054\u0020\u0070\u0069\u0078\u0065\u006c\u0073\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if len(_bff.GrAtX) != len(_bff.GrAtY) {
		return _aa.New("A\u0054\u0020\u0070\u0069xe\u006c \u0069\u006e\u0063\u006f\u006es\u0069\u0073\u0074\u0065\u006e\u0074")
	}
	_bff._ba = make([]bool, len(_bff.GrAtX))
	switch _bff.TemplateID {
	case 0:
		if _bff.GrAtX[0] != -1 && _bff.GrAtY[0] != -1 {
			_bff._ba[0] = true
			_bff._bf = true
		}
		if _bff.GrAtX[1] != -1 && _bff.GrAtY[1] != -1 {
			_bff._ba[1] = true
			_bff._bf = true
		}
	case 1:
		_bff._bf = false
	}
	return nil
}
func (_begeg *TextRegion) decodeRdx() (int64, error) {
	const _babc = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064x"
	if _begeg.IsHuffmanEncoded {
		if _begeg.SbHuffRDX == 3 {
			if _begeg._bgde == nil {
				var (
					_efaf  int
					_fgede error
				)
				if _begeg.SbHuffFS == 3 {
					_efaf++
				}
				if _begeg.SbHuffDS == 3 {
					_efaf++
				}
				if _begeg.SbHuffDT == 3 {
					_efaf++
				}
				if _begeg.SbHuffRDWidth == 3 {
					_efaf++
				}
				if _begeg.SbHuffRDHeight == 3 {
					_efaf++
				}
				_begeg._bgde, _fgede = _begeg.getUserTable(_efaf)
				if _fgede != nil {
					return 0, _fe.Wrap(_fgede, _babc, "")
				}
			}
			return _begeg._bgde.Decode(_begeg._deba)
		}
		_ceaf, _cefb := _ge.GetStandardTable(14 + int(_begeg.SbHuffRDX))
		if _cefb != nil {
			return 0, _fe.Wrap(_cefb, _babc, "")
		}
		return _ceaf.Decode(_begeg._deba)
	}
	_abge, _agcce := _begeg._fagf.DecodeInt(_begeg._gdgaf)
	if _agcce != nil {
		return 0, _fe.Wrap(_agcce, _babc, "")
	}
	return int64(_abge), nil
}
func (_dgfd *PatternDictionary) checkInput() error {
	if _dgfd.HdpHeight < 1 || _dgfd.HdpWidth < 1 {
		return _aa.New("in\u0076\u0061l\u0069\u0064\u0020\u0048\u0065\u0061\u0064\u0065\u0072 \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0057\u0069\u0064\u0074\u0068\u002f\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020g\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020z\u0065\u0072o")
	}
	if _dgfd.IsMMREncoded {
		if _dgfd.HDTemplate != 0 {
			_af.Log.Debug("\u0076\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0020\u0048\u0044\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030")
		}
	}
	return nil
}
func (_efdee *SymbolDictionary) getSymbol(_dgg int) (*_df.Bitmap, error) {
	const _fbca = "\u0067e\u0074\u0053\u0079\u006d\u0062\u006fl"
	_fdag, _aee := _efdee._daaa.GetBitmap(_efdee._face[_dgg])
	if _aee != nil {
		return nil, _fe.Wrap(_aee, _fbca, "\u0063\u0061n\u0027\u0074\u0020g\u0065\u0074\u0020\u0073\u0079\u006d\u0062\u006f\u006c")
	}
	return _fdag, nil
}
func (_dgdbe *TextRegion) decodeStripT() (_ffdg int64, _fagb error) {
	if _dgdbe.IsHuffmanEncoded {
		if _dgdbe.SbHuffDT == 3 {
			if _dgdbe._cafd == nil {
				var _cafb int
				if _dgdbe.SbHuffFS == 3 {
					_cafb++
				}
				if _dgdbe.SbHuffDS == 3 {
					_cafb++
				}
				_dgdbe._cafd, _fagb = _dgdbe.getUserTable(_cafb)
				if _fagb != nil {
					return 0, _fagb
				}
			}
			_ffdg, _fagb = _dgdbe._cafd.Decode(_dgdbe._deba)
			if _fagb != nil {
				return 0, _fagb
			}
		} else {
			var _bfedd _ge.Tabler
			_bfedd, _fagb = _ge.GetStandardTable(11 + int(_dgdbe.SbHuffDT))
			if _fagb != nil {
				return 0, _fagb
			}
			_ffdg, _fagb = _bfedd.Decode(_dgdbe._deba)
			if _fagb != nil {
				return 0, _fagb
			}
		}
	} else {
		var _bdda int32
		_bdda, _fagb = _dgdbe._fagf.DecodeInt(_dgdbe._bgce)
		if _fagb != nil {
			return 0, _fagb
		}
		_ffdg = int64(_bdda)
	}
	_ffdg *= int64(-_dgdbe.SbStrips)
	return _ffdg, nil
}
func (_ecbb *TextRegion) decodeRdy() (int64, error) {
	const _eedg = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064y"
	if _ecbb.IsHuffmanEncoded {
		if _ecbb.SbHuffRDY == 3 {
			if _ecbb._ffaa == nil {
				var (
					_eggfa int
					_dgef  error
				)
				if _ecbb.SbHuffFS == 3 {
					_eggfa++
				}
				if _ecbb.SbHuffDS == 3 {
					_eggfa++
				}
				if _ecbb.SbHuffDT == 3 {
					_eggfa++
				}
				if _ecbb.SbHuffRDWidth == 3 {
					_eggfa++
				}
				if _ecbb.SbHuffRDHeight == 3 {
					_eggfa++
				}
				if _ecbb.SbHuffRDX == 3 {
					_eggfa++
				}
				_ecbb._ffaa, _dgef = _ecbb.getUserTable(_eggfa)
				if _dgef != nil {
					return 0, _fe.Wrap(_dgef, _eedg, "")
				}
			}
			return _ecbb._ffaa.Decode(_ecbb._deba)
		}
		_gfeb, _eddf := _ge.GetStandardTable(14 + int(_ecbb.SbHuffRDY))
		if _eddf != nil {
			return 0, _eddf
		}
		return _gfeb.Decode(_ecbb._deba)
	}
	_eagb, _gfg := _ecbb._fagf.DecodeInt(_ecbb._fddb)
	if _gfg != nil {
		return 0, _fe.Wrap(_gfg, _eedg, "")
	}
	return int64(_eagb), nil
}

type template1 struct{}

func (_bgda *GenericRegion) overrideAtTemplate1(_dgdc, _bfdff, _cgb, _gedb, _aef int) int {
	_dgdc &= 0x1FF7
	if _bgda.GBAtY[0] == 0 && _bgda.GBAtX[0] >= -int8(_aef) {
		_dgdc |= (_gedb >> uint(7-(int8(_aef)+_bgda.GBAtX[0])) & 0x1) << 3
	} else {
		_dgdc |= int(_bgda.getPixel(_bfdff+int(_bgda.GBAtX[0]), _cgb+int(_bgda.GBAtY[0]))) << 3
	}
	return _dgdc
}

type templater interface {
	form(_dbf, _cdaa, _aeg, _adg, _cab int16) int16
	setIndex(_egc *_b.DecoderStats)
}

func _eea(_adcd *_ae.Reader, _gda *Header) *GenericRefinementRegion {
	return &GenericRefinementRegion{_fd: _adcd, RegionInfo: NewRegionSegment(_adcd), _ce: _gda, _fa: &template0{}, _gd: &template1{}}
}
func (_gcge *Header) readSegmentPageAssociation(_ffcc Documenter, _aacg *_ae.Reader, _bdac uint64, _eaeb ...int) (_fcgg error) {
	const _dcgc = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074P\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061t\u0069\u006f\u006e"
	if !_gcge.PageAssociationFieldSize {
		_gfbe, _dbd := _aacg.ReadBits(8)
		if _dbd != nil {
			return _fe.Wrap(_dbd, _dcgc, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_gcge.PageAssociation = int(_gfbe & 0xFF)
	} else {
		_edbb, _ada := _aacg.ReadBits(32)
		if _ada != nil {
			return _fe.Wrap(_ada, _dcgc, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_gcge.PageAssociation = int(_edbb & _c.MaxInt32)
	}
	if _bdac == 0 {
		return nil
	}
	if _gcge.PageAssociation != 0 {
		_adcf, _cccf := _ffcc.GetPage(_gcge.PageAssociation)
		if _cccf != nil {
			return _fe.Wrap(_cccf, _dcgc, "\u0061s\u0073\u006f\u0063\u0069a\u0074\u0065\u0064\u0020\u0070a\u0067e\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006ed")
		}
		var _abeea int
		for _acbg := uint64(0); _acbg < _bdac; _acbg++ {
			_abeea = _eaeb[_acbg]
			_gcge.RTSegments[_acbg], _cccf = _adcf.GetSegment(_abeea)
			if _cccf != nil {
				var _cada error
				_gcge.RTSegments[_acbg], _cada = _ffcc.GetGlobalSegment(_abeea)
				if _cada != nil {
					return _fe.Wrapf(_cccf, _dcgc, "\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020s\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0020\u0061\u0074\u0020pa\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0072\u0020\u0069\u006e\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0073", _gcge.PageAssociation)
				}
			}
		}
		return nil
	}
	for _abba := uint64(0); _abba < _bdac; _abba++ {
		_gcge.RTSegments[_abba], _fcgg = _ffcc.GetGlobalSegment(_eaeb[_abba])
		if _fcgg != nil {
			return _fe.Wrapf(_fcgg, _dcgc, "\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u003a\u0020\u0027\u0025d\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _eaeb[_abba])
		}
	}
	return nil
}
func (_eaf *GenericRefinementRegion) decodeTypicalPredictedLineTemplate1(_bcb, _fcg, _fbd, _ee, _edd, _ef, _dd, _cacf, _gef int) (_dcg error) {
	var (
		_faf, _bbg int
		_cdg, _gcb int
		_abc, _ff  int
		_abca      byte
	)
	if _bcb > 0 {
		_abca, _dcg = _eaf.RegionBitmap.GetByte(_dd - _fbd)
		if _dcg != nil {
			return _dcg
		}
		_cdg = int(_abca)
	}
	if _cacf > 0 && _cacf <= _eaf.ReferenceBitmap.Height {
		_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef - _ee + _ef)
		if _dcg != nil {
			return _dcg
		}
		_gcb = int(_abca) << 2
	}
	if _cacf >= 0 && _cacf < _eaf.ReferenceBitmap.Height {
		_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef + _ef)
		if _dcg != nil {
			return _dcg
		}
		_abc = int(_abca)
	}
	if _cacf > -2 && _cacf < _eaf.ReferenceBitmap.Height-1 {
		_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef + _ee + _ef)
		if _dcg != nil {
			return _dcg
		}
		_ff = int(_abca)
	}
	_faf = ((_cdg >> 5) & 0x6) | ((_ff >> 2) & 0x30) | (_abc & 0xc0) | (_gcb & 0x200)
	_bbg = ((_ff >> 2) & 0x70) | (_abc & 0xc0) | (_gcb & 0x700)
	var _adc int
	for _ffg := 0; _ffg < _edd; _ffg = _adc {
		var (
			_ded int
			_cc  int
		)
		_adc = _ffg + 8
		if _ded = _fcg - _ffg; _ded > 8 {
			_ded = 8
		}
		_dedb := _adc < _fcg
		_aabe := _adc < _eaf.ReferenceBitmap.Width
		_fgg := _ef + 1
		if _bcb > 0 {
			_abca = 0
			if _dedb {
				_abca, _dcg = _eaf.RegionBitmap.GetByte(_dd - _fbd + 1)
				if _dcg != nil {
					return _dcg
				}
			}
			_cdg = (_cdg << 8) | int(_abca)
		}
		if _cacf > 0 && _cacf <= _eaf.ReferenceBitmap.Height {
			var _cgge int
			if _aabe {
				_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef - _ee + _fgg)
				if _dcg != nil {
					return _dcg
				}
				_cgge = int(_abca) << 2
			}
			_gcb = (_gcb << 8) | _cgge
		}
		if _cacf >= 0 && _cacf < _eaf.ReferenceBitmap.Height {
			_abca = 0
			if _aabe {
				_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef + _fgg)
				if _dcg != nil {
					return _dcg
				}
			}
			_abc = (_abc << 8) | int(_abca)
		}
		if _cacf > -2 && _cacf < (_eaf.ReferenceBitmap.Height-1) {
			_abca = 0
			if _aabe {
				_abca, _dcg = _eaf.ReferenceBitmap.GetByte(_gef + _ee + _fgg)
				if _dcg != nil {
					return _dcg
				}
			}
			_ff = (_ff << 8) | int(_abca)
		}
		for _gbaf := 0; _gbaf < _ded; _gbaf++ {
			var _eaa int
			_gadf := (_bbg >> 4) & 0x1ff
			switch _gadf {
			case 0x1ff:
				_eaa = 1
			case 0x00:
				_eaa = 0
			default:
				_eaf._ab.SetIndex(int32(_faf))
				_eaa, _dcg = _eaf._eg.DecodeBit(_eaf._ab)
				if _dcg != nil {
					return _dcg
				}
			}
			_fce := uint(7 - _gbaf)
			_cc |= _eaa << _fce
			_faf = ((_faf & 0x0d6) << 1) | _eaa | (_cdg>>_fce+5)&0x002 | ((_ff>>_fce + 2) & 0x010) | ((_abc >> _fce) & 0x040) | ((_gcb >> _fce) & 0x200)
			_bbg = ((_bbg & 0xdb) << 1) | ((_ff>>_fce + 2) & 0x010) | ((_abc >> _fce) & 0x080) | ((_gcb >> _fce) & 0x400)
		}
		_dcg = _eaf.RegionBitmap.SetByte(_dd, byte(_cc))
		if _dcg != nil {
			return _dcg
		}
		_dd++
		_gef++
	}
	return nil
}
func (_gffa *TableSegment) HtPS() int32 { return _gffa._dfdfb }
func (_bcgc *TextRegion) setParameters(_daef *_b.Decoder, _fdcc, _cgga bool, _aaed, _aedf uint32, _bdfe uint32, _fcba int8, _gbbc uint32, _becg int8, _deeag _df.CombinationOperator, _fbed int8, _ecce int16, _ecfc, _dgbd, _cfacd, _eafd, _abce, _cgdbb, _accc, _caffc, _dec, _bedgc int8, _ccgg, _cdea []int8, _adcfg []*_df.Bitmap, _ceafa int8) {
	_bcgc._fagf = _daef
	_bcgc.IsHuffmanEncoded = _fdcc
	_bcgc.UseRefinement = _cgga
	_bcgc.RegionInfo.BitmapWidth = _aaed
	_bcgc.RegionInfo.BitmapHeight = _aedf
	_bcgc.NumberOfSymbolInstances = _bdfe
	_bcgc.SbStrips = _fcba
	_bcgc.NumberOfSymbols = _gbbc
	_bcgc.DefaultPixel = _becg
	_bcgc.CombinationOperator = _deeag
	_bcgc.IsTransposed = _fbed
	_bcgc.ReferenceCorner = _ecce
	_bcgc.SbDsOffset = _ecfc
	_bcgc.SbHuffFS = _dgbd
	_bcgc.SbHuffDS = _cfacd
	_bcgc.SbHuffDT = _eafd
	_bcgc.SbHuffRDWidth = _abce
	_bcgc.SbHuffRDHeight = _cgdbb
	_bcgc.SbHuffRSize = _dec
	_bcgc.SbHuffRDX = _accc
	_bcgc.SbHuffRDY = _caffc
	_bcgc.SbrTemplate = _bedgc
	_bcgc.SbrATX = _ccgg
	_bcgc.SbrATY = _cdea
	_bcgc.Symbols = _adcfg
	_bcgc._adfd = _ceafa
}
func (_afbb *TextRegion) getUserTable(_dfgb int) (_ge.Tabler, error) {
	const _bgcc = "\u0067\u0065\u0074U\u0073\u0065\u0072\u0054\u0061\u0062\u006c\u0065"
	var _ecaa int
	for _, _eggg := range _afbb.Header.RTSegments {
		if _eggg.Type == 53 {
			if _ecaa == _dfgb {
				_dgcc, _fead := _eggg.GetSegmentData()
				if _fead != nil {
					return nil, _fead
				}
				_efad, _gaca := _dgcc.(*TableSegment)
				if !_gaca {
					_af.Log.Debug(_dg.Sprintf("\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079p\u0065\u00205\u0033\u0020\u002d\u0020\u0061\u006ed\u0020in\u0064\u0065\u0078\u003a\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _eggg.SegmentNumber))
					return nil, _fe.Error(_bgcc, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079\u0070e\u0020\u0035\u0033\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u002a\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
				}
				return _ge.NewEncodedTable(_efad)
			}
			_ecaa++
		}
	}
	return nil, nil
}
func (_gebf *PageInformationSegment) readDefaultPixelValue() error {
	_gece, _fdc := _gebf._fgee.ReadBit()
	if _fdc != nil {
		return _fdc
	}
	_gebf.DefaultPixelValue = uint8(_gece & 0xf)
	return nil
}
func (_fbbe *SymbolDictionary) decodeThroughTextRegion(_ecca, _gbce, _cefg uint32) error {
	if _fbbe._ebe == nil {
		_fbbe._ebe = _cgebd(_fbbe._bbda, nil)
		_fbbe._ebe.setContexts(_fbbe._fbcd, _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _fbbe._ecc, _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1))
	}
	if _gaag := _fbbe.setSymbolsArray(); _gaag != nil {
		return _gaag
	}
	_fbbe._ebe.setParameters(_fbbe._fdbf, _fbbe.IsHuffmanEncoded, true, _ecca, _gbce, _cefg, 1, _fbbe._eaab+_fbbe._dacc, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, _fbbe.SdrTemplate, _fbbe.SdrATX, _fbbe.SdrATY, _fbbe._bccf, _fbbe._efgc)
	return _fbbe.addSymbol(_fbbe._ebe)
}
func (_gdb *PatternDictionary) readTemplate() error {
	_cccd, _bcgg := _gdb._fdcf.ReadBits(2)
	if _bcgg != nil {
		return _bcgg
	}
	_gdb.HDTemplate = byte(_cccd)
	return nil
}
func (_aaee *TextRegion) readUseRefinement() error {
	if !_aaee.UseRefinement || _aaee.SbrTemplate != 0 {
		return nil
	}
	var (
		_cbff byte
		_gbdb error
	)
	_aaee.SbrATX = make([]int8, 2)
	_aaee.SbrATY = make([]int8, 2)
	_cbff, _gbdb = _aaee._deba.ReadByte()
	if _gbdb != nil {
		return _gbdb
	}
	_aaee.SbrATX[0] = int8(_cbff)
	_cbff, _gbdb = _aaee._deba.ReadByte()
	if _gbdb != nil {
		return _gbdb
	}
	_aaee.SbrATY[0] = int8(_cbff)
	_cbff, _gbdb = _aaee._deba.ReadByte()
	if _gbdb != nil {
		return _gbdb
	}
	_aaee.SbrATX[1] = int8(_cbff)
	_cbff, _gbdb = _aaee._deba.ReadByte()
	if _gbdb != nil {
		return _gbdb
	}
	_aaee.SbrATY[1] = int8(_cbff)
	return nil
}

type RegionSegment struct {
	_gega              *_ae.Reader
	BitmapWidth        uint32
	BitmapHeight       uint32
	XLocation          uint32
	YLocation          uint32
	CombinaionOperator _df.CombinationOperator
}

func (_dcfe *Header) pageSize() uint {
	if _dcfe.PageAssociation <= 255 {
		return 1
	}
	return 4
}
func (_dea *GenericRegion) decodeTemplate3(_deff, _dba, _abcd int, _bag, _ceg int) (_bace error) {
	const _gge = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0033"
	var (
		_dff, _cfaf int
		_eefd       int
		_eab        byte
		_dab, _cgea int
	)
	if _deff >= 1 {
		_eab, _bace = _dea.Bitmap.GetByte(_ceg)
		if _bace != nil {
			return _fe.Wrap(_bace, _gge, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_eefd = int(_eab)
	}
	_dff = (_eefd >> 1) & 0x70
	for _fcfe := 0; _fcfe < _abcd; _fcfe = _dab {
		var (
			_ebbg  byte
			_bcbbe int
		)
		_dab = _fcfe + 8
		if _adge := _dba - _fcfe; _adge > 8 {
			_bcbbe = 8
		} else {
			_bcbbe = _adge
		}
		if _deff >= 1 {
			_eefd <<= 8
			if _dab < _dba {
				_eab, _bace = _dea.Bitmap.GetByte(_ceg + 1)
				if _bace != nil {
					return _fe.Wrap(_bace, _gge, "\u0069\u006e\u006e\u0065\u0072\u0020\u002d\u0020\u006c\u0069\u006e\u0065 \u003e\u003d\u0020\u0031")
				}
				_eefd |= int(_eab)
			}
		}
		for _egf := 0; _egf < _bcbbe; _egf++ {
			if _dea._dca {
				_cfaf = _dea.overrideAtTemplate3(_dff, _fcfe+_egf, _deff, int(_ebbg), _egf)
				_dea._bfe.SetIndex(int32(_cfaf))
			} else {
				_dea._bfe.SetIndex(int32(_dff))
			}
			_cgea, _bace = _dea._caf.DecodeBit(_dea._bfe)
			if _bace != nil {
				return _fe.Wrap(_bace, _gge, "")
			}
			_ebbg |= byte(_cgea) << byte(7-_egf)
			_dff = ((_dff & 0x1f7) << 1) | _cgea | ((_eefd >> uint(8-_egf)) & 0x010)
		}
		if _eeafb := _dea.Bitmap.SetByte(_bag, _ebbg); _eeafb != nil {
			return _fe.Wrap(_eeafb, _gge, "")
		}
		_bag++
		_ceg++
	}
	return nil
}
func (_aebb *RegionSegment) parseHeader() error {
	const _bgagc = "p\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_af.Log.Trace("\u005b\u0052\u0045\u0047I\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002dH\u0045A\u0044\u0045\u0052\u005d\u0020\u0042\u0065g\u0069\u006e")
	defer func() {
		_af.Log.Trace("\u005b\u0052\u0045G\u0049\u004f\u004e\u005d[\u0050\u0041\u0052\u0053\u0045\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_agcf, _dbdb := _aebb._gega.ReadBits(32)
	if _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "\u0077\u0069\u0064t\u0068")
	}
	_aebb.BitmapWidth = uint32(_agcf & _c.MaxUint32)
	_agcf, _dbdb = _aebb._gega.ReadBits(32)
	if _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_aebb.BitmapHeight = uint32(_agcf & _c.MaxUint32)
	_agcf, _dbdb = _aebb._gega.ReadBits(32)
	if _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "\u0078\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_aebb.XLocation = uint32(_agcf & _c.MaxUint32)
	_agcf, _dbdb = _aebb._gega.ReadBits(32)
	if _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "\u0079\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_aebb.YLocation = uint32(_agcf & _c.MaxUint32)
	if _, _dbdb = _aebb._gega.ReadBits(5); _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "\u0064i\u0072\u0079\u0020\u0072\u0065\u0061d")
	}
	if _dbdb = _aebb.readCombinationOperator(); _dbdb != nil {
		return _fe.Wrap(_dbdb, _bgagc, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	return nil
}
func (_aaa *GenericRegion) updateOverrideFlags() error {
	const _ffc = "\u0075\u0070\u0064\u0061te\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _aaa.GBAtX == nil || _aaa.GBAtY == nil {
		return nil
	}
	if len(_aaa.GBAtX) != len(_aaa.GBAtY) {
		return _fe.Errorf(_ffc, "i\u006eco\u0073i\u0073t\u0065\u006e\u0074\u0020\u0041T\u0020\u0070\u0069x\u0065\u006c\u002e\u0020\u0041m\u006f\u0075\u006et\u0020\u006f\u0066\u0020\u0027\u0078\u0027\u0020\u0070\u0069\u0078e\u006c\u0073\u003a %d\u002c\u0020\u0041\u006d\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0027\u0079\u0027\u0020\u0070\u0069\u0078e\u006cs\u003a\u0020\u0025\u0064", len(_aaa.GBAtX), len(_aaa.GBAtY))
	}
	_aaa.GBAtOverride = make([]bool, len(_aaa.GBAtX))
	switch _aaa.GBTemplate {
	case 0:
		if !_aaa.UseExtTemplates {
			if _aaa.GBAtX[0] != 3 || _aaa.GBAtY[0] != -1 {
				_aaa.setOverrideFlag(0)
			}
			if _aaa.GBAtX[1] != -3 || _aaa.GBAtY[1] != -1 {
				_aaa.setOverrideFlag(1)
			}
			if _aaa.GBAtX[2] != 2 || _aaa.GBAtY[2] != -2 {
				_aaa.setOverrideFlag(2)
			}
			if _aaa.GBAtX[3] != -2 || _aaa.GBAtY[3] != -2 {
				_aaa.setOverrideFlag(3)
			}
		} else {
			if _aaa.GBAtX[0] != -2 || _aaa.GBAtY[0] != 0 {
				_aaa.setOverrideFlag(0)
			}
			if _aaa.GBAtX[1] != 0 || _aaa.GBAtY[1] != -2 {
				_aaa.setOverrideFlag(1)
			}
			if _aaa.GBAtX[2] != -2 || _aaa.GBAtY[2] != -1 {
				_aaa.setOverrideFlag(2)
			}
			if _aaa.GBAtX[3] != -1 || _aaa.GBAtY[3] != -2 {
				_aaa.setOverrideFlag(3)
			}
			if _aaa.GBAtX[4] != 1 || _aaa.GBAtY[4] != -2 {
				_aaa.setOverrideFlag(4)
			}
			if _aaa.GBAtX[5] != 2 || _aaa.GBAtY[5] != -1 {
				_aaa.setOverrideFlag(5)
			}
			if _aaa.GBAtX[6] != -3 || _aaa.GBAtY[6] != 0 {
				_aaa.setOverrideFlag(6)
			}
			if _aaa.GBAtX[7] != -4 || _aaa.GBAtY[7] != 0 {
				_aaa.setOverrideFlag(7)
			}
			if _aaa.GBAtX[8] != 2 || _aaa.GBAtY[8] != -2 {
				_aaa.setOverrideFlag(8)
			}
			if _aaa.GBAtX[9] != 3 || _aaa.GBAtY[9] != -1 {
				_aaa.setOverrideFlag(9)
			}
			if _aaa.GBAtX[10] != -2 || _aaa.GBAtY[10] != -2 {
				_aaa.setOverrideFlag(10)
			}
			if _aaa.GBAtX[11] != -3 || _aaa.GBAtY[11] != -1 {
				_aaa.setOverrideFlag(11)
			}
		}
	case 1:
		if _aaa.GBAtX[0] != 3 || _aaa.GBAtY[0] != -1 {
			_aaa.setOverrideFlag(0)
		}
	case 2:
		if _aaa.GBAtX[0] != 2 || _aaa.GBAtY[0] != -1 {
			_aaa.setOverrideFlag(0)
		}
	case 3:
		if _aaa.GBAtX[0] != 2 || _aaa.GBAtY[0] != -1 {
			_aaa.setOverrideFlag(0)
		}
	}
	return nil
}
func (_aadcg Type) String() string {
	switch _aadcg {
	case TSymbolDictionary:
		return "\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079"
	case TIntermediateTextRegion:
		return "\u0049n\u0074\u0065\u0072\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020T\u0065\u0078\u0074\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateTextRegion:
		return "I\u006d\u006d\u0065\u0064ia\u0074e\u0020\u0054\u0065\u0078\u0074 \u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateLosslessTextRegion:
		return "\u0049\u006d\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020L\u006f\u0073\u0073\u006c\u0065\u0073\u0073 \u0054\u0065\u0078\u0074\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TPatternDictionary:
		return "\u0050a\u0074t\u0065\u0072\u006e\u0020\u0044i\u0063\u0074i\u006f\u006e\u0061\u0072\u0079"
	case TIntermediateHalftoneRegion:
		return "\u0049\u006e\u0074\u0065r\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0048\u0061l\u0066t\u006f\u006e\u0065\u0020\u0052\u0065\u0067i\u006f\u006e"
	case TImmediateHalftoneRegion:
		return "\u0049m\u006d\u0065\u0064\u0069a\u0074\u0065\u0020\u0048\u0061l\u0066t\u006fn\u0065\u0020\u0052\u0065\u0067\u0069\u006fn"
	case TImmediateLosslessHalftoneRegion:
		return "\u0049\u006d\u006ded\u0069\u0061\u0074\u0065\u0020\u004c\u006f\u0073\u0073l\u0065s\u0073 \u0048a\u006c\u0066\u0074\u006f\u006e\u0065\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TIntermediateGenericRegion:
		return "I\u006e\u0074\u0065\u0072\u006d\u0065d\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006e\u0065r\u0069\u0063\u0020R\u0065g\u0069\u006f\u006e"
	case TImmediateGenericRegion:
		return "\u0049m\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006ee\u0072\u0069\u0063\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateLosslessGenericRegion:
		return "\u0049\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u0020\u004c\u006f\u0073\u0073\u006c\u0065\u0073\u0073\u0020\u0047e\u006e\u0065\u0072\u0069\u0063\u0020\u0052e\u0067\u0069\u006f\u006e"
	case TIntermediateGenericRefinementRegion:
		return "\u0049\u006e\u0074\u0065\u0072\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0020\u0052\u0065\u0066i\u006e\u0065\u006d\u0065\u006et\u0020\u0052e\u0067\u0069\u006f\u006e"
	case TImmediateGenericRefinementRegion:
		return "I\u006d\u006d\u0065\u0064\u0069\u0061t\u0065\u0020\u0047\u0065\u006e\u0065r\u0069\u0063\u0020\u0052\u0065\u0066\u0069n\u0065\u006d\u0065\u006e\u0074\u0020\u0052\u0065\u0067\u0069o\u006e"
	case TImmediateLosslessGenericRefinementRegion:
		return "\u0049m\u006d\u0065d\u0069\u0061\u0074\u0065 \u004c\u006f\u0073s\u006c\u0065\u0073\u0073\u0020\u0047\u0065\u006e\u0065ri\u0063\u0020\u0052e\u0066\u0069n\u0065\u006d\u0065\u006e\u0074\u0020R\u0065\u0067i\u006f\u006e"
	case TPageInformation:
		return "\u0050\u0061g\u0065\u0020\u0049n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e"
	case TEndOfPage:
		return "E\u006e\u0064\u0020\u004f\u0066\u0020\u0050\u0061\u0067\u0065"
	case TEndOfStrip:
		return "\u0045\u006e\u0064 \u004f\u0066\u0020\u0053\u0074\u0072\u0069\u0070"
	case TEndOfFile:
		return "E\u006e\u0064\u0020\u004f\u0066\u0020\u0046\u0069\u006c\u0065"
	case TProfiles:
		return "\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0073"
	case TTables:
		return "\u0054\u0061\u0062\u006c\u0065\u0073"
	case TExtension:
		return "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"
	case TBitmap:
		return "\u0042\u0069\u0074\u006d\u0061\u0070"
	}
	return "I\u006ev\u0061\u006c\u0069\u0064\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0020Ki\u006e\u0064"
}
func (_ddcb *SymbolDictionary) decodeNewSymbols(_fggee, _daeg uint32, _eeeb *_df.Bitmap, _fcdf, _dagf int32) error {
	if _ddcb._abcc == nil {
		_ddcb._abcc = _eea(_ddcb._bbda, nil)
		if _ddcb._fdbf == nil {
			var _eefa error
			_ddcb._fdbf, _eefa = _b.New(_ddcb._bbda)
			if _eefa != nil {
				return _eefa
			}
		}
		if _ddcb._fbcd == nil {
			_ddcb._fbcd = _b.NewStats(65536, 1)
		}
	}
	_ddcb._abcc.setParameters(_ddcb._fbcd, _ddcb._fdbf, _ddcb.SdrTemplate, _fggee, _daeg, _eeeb, _fcdf, _dagf, false, _ddcb.SdrATX, _ddcb.SdrATY)
	return _ddcb.addSymbol(_ddcb._abcc)
}
func (_bba *GenericRegion) writeGBAtPixels(_bab _ae.BinaryWriter) (_gacf int, _fbge error) {
	const _cbdf = "\u0077r\u0069t\u0065\u0047\u0042\u0041\u0074\u0050\u0069\u0078\u0065\u006c\u0073"
	if _bba.UseMMR {
		return 0, nil
	}
	_gcc := 1
	if _bba.GBTemplate == 0 {
		_gcc = 4
	} else if _bba.UseExtTemplates {
		_gcc = 12
	}
	if len(_bba.GBAtX) != _gcc {
		return 0, _fe.Errorf(_cbdf, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0058\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	if len(_bba.GBAtY) != _gcc {
		return 0, _fe.Errorf(_cbdf, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0059\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	for _aaae := 0; _aaae < _gcc; _aaae++ {
		if _fbge = _bab.WriteByte(byte(_bba.GBAtX[_aaae])); _fbge != nil {
			return _gacf, _fe.Wrap(_fbge, _cbdf, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0058")
		}
		_gacf++
		if _fbge = _bab.WriteByte(byte(_bba.GBAtY[_aaae])); _fbge != nil {
			return _gacf, _fe.Wrap(_fbge, _cbdf, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0059")
		}
		_gacf++
	}
	return _gacf, nil
}
func (_ec *EndOfStripe) parseHeader() error {
	_ag, _cg := _ec._bc.ReadBits(32)
	if _cg != nil {
		return _cg
	}
	_ec._fgce = int(_ag & _c.MaxInt32)
	return nil
}
func (_gbdac *TextRegion) symbolIDCodeLengths() error {
	var (
		_cebfb []*_ge.Code
		_ccaa  uint64
		_bcefg _ge.Tabler
		_ebab  error
	)
	for _ebbgg := 0; _ebbgg < 35; _ebbgg++ {
		_ccaa, _ebab = _gbdac._deba.ReadBits(4)
		if _ebab != nil {
			return _ebab
		}
		_dege := int(_ccaa & 0xf)
		if _dege > 0 {
			_cebfb = append(_cebfb, _ge.NewCode(int32(_dege), 0, int32(_ebbgg), false))
		}
	}
	_bcefg, _ebab = _ge.NewFixedSizeTable(_cebfb)
	if _ebab != nil {
		return _ebab
	}
	var (
		_gbbcg int64
		_efaff uint32
		_agdb  []*_ge.Code
		_fbaf  int64
	)
	for _efaff < _gbdac.NumberOfSymbols {
		_fbaf, _ebab = _bcefg.Decode(_gbdac._deba)
		if _ebab != nil {
			return _ebab
		}
		if _fbaf < 32 {
			if _fbaf > 0 {
				_agdb = append(_agdb, _ge.NewCode(int32(_fbaf), 0, int32(_efaff), false))
			}
			_gbbcg = _fbaf
			_efaff++
		} else {
			var _ddeg, _caa int64
			switch _fbaf {
			case 32:
				_ccaa, _ebab = _gbdac._deba.ReadBits(2)
				if _ebab != nil {
					return _ebab
				}
				_ddeg = 3 + int64(_ccaa)
				if _efaff > 0 {
					_caa = _gbbcg
				}
			case 33:
				_ccaa, _ebab = _gbdac._deba.ReadBits(3)
				if _ebab != nil {
					return _ebab
				}
				_ddeg = 3 + int64(_ccaa)
			case 34:
				_ccaa, _ebab = _gbdac._deba.ReadBits(7)
				if _ebab != nil {
					return _ebab
				}
				_ddeg = 11 + int64(_ccaa)
			}
			for _bdade := 0; _bdade < int(_ddeg); _bdade++ {
				if _caa > 0 {
					_agdb = append(_agdb, _ge.NewCode(int32(_caa), 0, int32(_efaff), false))
				}
				_efaff++
			}
		}
	}
	_gbdac._deba.Align()
	_gbdac._bebbb, _ebab = _ge.NewFixedSizeTable(_agdb)
	return _ebab
}
func (_baffb *TableSegment) HtRS() int32 { return _baffb._bfcc }
func (_gcfe *TextRegion) decodeDfs() (int64, error) {
	if _gcfe.IsHuffmanEncoded {
		if _gcfe.SbHuffFS == 3 {
			if _gcfe._cbgg == nil {
				var _fdgad error
				_gcfe._cbgg, _fdgad = _gcfe.getUserTable(0)
				if _fdgad != nil {
					return 0, _fdgad
				}
			}
			return _gcfe._cbgg.Decode(_gcfe._deba)
		}
		_edadf, _aaaef := _ge.GetStandardTable(6 + int(_gcfe.SbHuffFS))
		if _aaaef != nil {
			return 0, _aaaef
		}
		return _edadf.Decode(_gcfe._deba)
	}
	_begad, _gdaf := _gcfe._fagf.DecodeInt(_gcfe._deef)
	if _gdaf != nil {
		return 0, _gdaf
	}
	return int64(_begad), nil
}
func (_aafa *PageInformationSegment) readIsLossless() error {
	_cgeb, _adbe := _aafa._fgee.ReadBit()
	if _adbe != nil {
		return _adbe
	}
	if _cgeb == 1 {
		_aafa.IsLossless = true
	}
	return nil
}
func (_accd *SymbolDictionary) decodeHeightClassDeltaHeight() (int64, error) {
	if _accd.IsHuffmanEncoded {
		return _accd.decodeHeightClassDeltaHeightWithHuffman()
	}
	_gfe, _ebed := _accd._fdbf.DecodeInt(_accd._cacfe)
	if _ebed != nil {
		return 0, _ebed
	}
	return int64(_gfe), nil
}
func (_efde *PageInformationSegment) readIsStriped() error {
	_gebd, _ffga := _efde._fgee.ReadBit()
	if _ffga != nil {
		return _ffga
	}
	if _gebd == 1 {
		_efde.IsStripe = true
	}
	return nil
}
func (_deb *PatternDictionary) parseHeader() error {
	_af.Log.Trace("\u005b\u0050\u0041\u0054\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054I\u004f\u004e\u0041\u0052\u0059\u005d[\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d\u0020b\u0065\u0067\u0069\u006e")
	defer func() {
		_af.Log.Trace("\u005b\u0050\u0041T\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u005b\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_, _gdgda := _deb._fdcf.ReadBits(5)
	if _gdgda != nil {
		return _gdgda
	}
	if _gdgda = _deb.readTemplate(); _gdgda != nil {
		return _gdgda
	}
	if _gdgda = _deb.readIsMMREncoded(); _gdgda != nil {
		return _gdgda
	}
	if _gdgda = _deb.readPatternWidthAndHeight(); _gdgda != nil {
		return _gdgda
	}
	if _gdgda = _deb.readGrayMax(); _gdgda != nil {
		return _gdgda
	}
	if _gdgda = _deb.computeSegmentDataStructure(); _gdgda != nil {
		return _gdgda
	}
	return _deb.checkInput()
}

type PatternDictionary struct {
	_fdcf            *_ae.Reader
	DataHeaderOffset int64
	DataHeaderLength int64
	DataOffset       int64
	DataLength       int64
	GBAtX            []int8
	GBAtY            []int8
	IsMMREncoded     bool
	HDTemplate       byte
	HdpWidth         byte
	HdpHeight        byte
	Patterns         []*_df.Bitmap
	GrayMax          uint32
}

func (_ffa *GenericRefinementRegion) setParameters(_gaf *_b.DecoderStats, _cacg *_b.Decoder, _bcf int8, _fbg, _bfc uint32, _dcc *_df.Bitmap, _aac, _fec int32, _defc bool, _aag []int8, _cfe []int8) {
	_af.Log.Trace("\u005b\u0047\u0045NE\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052E\u0047I\u004fN\u005d \u0073\u0065\u0074\u0050\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	if _gaf != nil {
		_ffa._ab = _gaf
	}
	if _cacg != nil {
		_ffa._eg = _cacg
	}
	_ffa.TemplateID = _bcf
	_ffa.RegionInfo.BitmapWidth = _fbg
	_ffa.RegionInfo.BitmapHeight = _bfc
	_ffa.ReferenceBitmap = _dcc
	_ffa.ReferenceDX = _aac
	_ffa.ReferenceDY = _fec
	_ffa.IsTPGROn = _defc
	_ffa.GrAtX = _aag
	_ffa.GrAtY = _cfe
	_ffa.RegionBitmap = nil
	_af.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073 f\u0069\u006ei\u0073\u0068\u0065\u0064\u002e\u0020\u0025\u0073", _ffa)
}
func (_ca *EndOfStripe) LineNumber() int { return _ca._fgce }
func (_dfga *template1) form(_bge, _abdg, _dbb, _ccd, _bdc int16) int16 {
	return ((_bge & 0x02) << 8) | (_abdg << 6) | ((_dbb & 0x03) << 4) | (_ccd << 1) | _bdc
}
func (_bcag *PageInformationSegment) readRequiresAuxiliaryBuffer() error {
	_ggfc, _bcfa := _bcag._fgee.ReadBit()
	if _bcfa != nil {
		return _bcfa
	}
	if _ggfc == 1 {
		_bcag._fgaa = true
	}
	return nil
}

var (
	_gdff Segmenter
	_dae  = map[Type]func() Segmenter{TSymbolDictionary: func() Segmenter { return &SymbolDictionary{} }, TIntermediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateLosslessTextRegion: func() Segmenter { return &TextRegion{} }, TPatternDictionary: func() Segmenter { return &PatternDictionary{} }, TIntermediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateLosslessHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TIntermediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateLosslessGenericRegion: func() Segmenter { return &GenericRegion{} }, TIntermediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateLosslessGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TPageInformation: func() Segmenter { return &PageInformationSegment{} }, TEndOfPage: func() Segmenter { return _gdff }, TEndOfStrip: func() Segmenter { return &EndOfStripe{} }, TEndOfFile: func() Segmenter { return _gdff }, TProfiles: func() Segmenter { return _gdff }, TTables: func() Segmenter { return &TableSegment{} }, TExtension: func() Segmenter { return _gdff }, TBitmap: func() Segmenter { return _gdff }}
)

func (_fegc *TableSegment) Init(h *Header, r *_ae.Reader) error {
	_fegc._ebef = r
	return _fegc.parseHeader()
}
func (_cegbb *SymbolDictionary) retrieveImportSymbols() error {
	for _, _bfgb := range _cegbb.Header.RTSegments {
		if _bfgb.Type == 0 {
			_adcdb, _beef := _bfgb.GetSegmentData()
			if _beef != nil {
				return _beef
			}
			_bfgf, _cgce := _adcdb.(*SymbolDictionary)
			if !_cgce {
				return _dg.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0044\u0061\u0074a\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u003a\u0020%\u0054", _adcdb)
			}
			_ebfbg, _beef := _bfgf.GetDictionary()
			if _beef != nil {
				return _dg.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064 \u0073\u0065\u0067m\u0065\u006e\u0074 \u0077\u0069t\u0068\u0020\u0069\u006e\u0064\u0065x\u003a %\u0064\u0020\u0067\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0073", _bfgb.SegmentNumber, _beef.Error())
			}
			_cegbb._fege = append(_cegbb._fege, _ebfbg...)
			_cegbb._eaab += _bfgf.NumberOfExportedSymbols
		}
	}
	return nil
}
func (_dgfc *HalftoneRegion) grayScaleDecoding(_eee int) ([][]int, error) {
	var (
		_daa []int8
		_efd []int8
	)
	if !_dgfc.IsMMREncoded {
		_daa = make([]int8, 4)
		_efd = make([]int8, 4)
		if _dgfc.HTemplate <= 1 {
			_daa[0] = 3
		} else if _dgfc.HTemplate >= 2 {
			_daa[0] = 2
		}
		_efd[0] = -1
		_daa[1] = -3
		_efd[1] = -1
		_daa[2] = 2
		_efd[2] = -2
		_daa[3] = -2
		_efd[3] = -2
	}
	_ggaa := make([]*_df.Bitmap, _eee)
	_bcgd := NewGenericRegion(_dgfc._baab)
	_bcgd.setParametersMMR(_dgfc.IsMMREncoded, _dgfc.DataOffset, _dgfc.DataLength, _dgfc.HGridHeight, _dgfc.HGridWidth, _dgfc.HTemplate, false, _dgfc.HSkipEnabled, _daa, _efd)
	_gdfc := _eee - 1
	var _cbg error
	_ggaa[_gdfc], _cbg = _bcgd.GetRegionBitmap()
	if _cbg != nil {
		return nil, _cbg
	}
	for _gdfc > 0 {
		_gdfc--
		_bcgd.Bitmap = nil
		_ggaa[_gdfc], _cbg = _bcgd.GetRegionBitmap()
		if _cbg != nil {
			return nil, _cbg
		}
		if _cbg = _dgfc.combineGrayscalePlanes(_ggaa, _gdfc); _cbg != nil {
			return nil, _cbg
		}
	}
	return _dgfc.computeGrayScalePlanes(_ggaa, _eee)
}
func (_fbfe *SymbolDictionary) huffDecodeRefAggNInst() (int64, error) {
	if !_fbfe.SdHuffAggInstanceSelection {
		_ceab, _gbabc := _ge.GetStandardTable(1)
		if _gbabc != nil {
			return 0, _gbabc
		}
		return _ceab.Decode(_fbfe._bbda)
	}
	if _fbfe._fdgf == nil {
		var (
			_daee int
			_dedg error
		)
		if _fbfe.SdHuffDecodeHeightSelection == 3 {
			_daee++
		}
		if _fbfe.SdHuffDecodeWidthSelection == 3 {
			_daee++
		}
		if _fbfe.SdHuffBMSizeSelection == 3 {
			_daee++
		}
		_fbfe._fdgf, _dedg = _fbfe.getUserTable(_daee)
		if _dedg != nil {
			return 0, _dedg
		}
	}
	return _fbfe._fdgf.Decode(_fbfe._bbda)
}
func (_bgead *SymbolDictionary) decodeHeightClassDeltaHeightWithHuffman() (int64, error) {
	switch _bgead.SdHuffDecodeHeightSelection {
	case 0:
		_ggd, _ggcde := _ge.GetStandardTable(4)
		if _ggcde != nil {
			return 0, _ggcde
		}
		return _ggd.Decode(_bgead._bbda)
	case 1:
		_afdf, _fbab := _ge.GetStandardTable(5)
		if _fbab != nil {
			return 0, _fbab
		}
		return _afdf.Decode(_bgead._bbda)
	case 3:
		if _bgead._fcdb == nil {
			_gecdb, _ecd := _ge.GetStandardTable(0)
			if _ecd != nil {
				return 0, _ecd
			}
			_bgead._fcdb = _gecdb
		}
		return _bgead._fcdb.Decode(_bgead._bbda)
	}
	return 0, nil
}

var _ SegmentEncoder = &RegionSegment{}

func (_ebfg *TableSegment) HtHigh() int32 { return _ebfg._gafe }
func (_eaac *HalftoneRegion) parseHeader() error {
	if _edff := _eaac.RegionSegment.parseHeader(); _edff != nil {
		return _edff
	}
	_afc, _afb := _eaac._baab.ReadBit()
	if _afb != nil {
		return _afb
	}
	_eaac.HDefaultPixel = int8(_afc)
	_dced, _afb := _eaac._baab.ReadBits(3)
	if _afb != nil {
		return _afb
	}
	_eaac.CombinationOperator = _df.CombinationOperator(_dced & 0xf)
	_afc, _afb = _eaac._baab.ReadBit()
	if _afb != nil {
		return _afb
	}
	if _afc == 1 {
		_eaac.HSkipEnabled = true
	}
	_dced, _afb = _eaac._baab.ReadBits(2)
	if _afb != nil {
		return _afb
	}
	_eaac.HTemplate = byte(_dced & 0xf)
	_afc, _afb = _eaac._baab.ReadBit()
	if _afb != nil {
		return _afb
	}
	if _afc == 1 {
		_eaac.IsMMREncoded = true
	}
	_dced, _afb = _eaac._baab.ReadBits(32)
	if _afb != nil {
		return _afb
	}
	_eaac.HGridWidth = uint32(_dced & _c.MaxUint32)
	_dced, _afb = _eaac._baab.ReadBits(32)
	if _afb != nil {
		return _afb
	}
	_eaac.HGridHeight = uint32(_dced & _c.MaxUint32)
	_dced, _afb = _eaac._baab.ReadBits(32)
	if _afb != nil {
		return _afb
	}
	_eaac.HGridX = int32(_dced & _c.MaxInt32)
	_dced, _afb = _eaac._baab.ReadBits(32)
	if _afb != nil {
		return _afb
	}
	_eaac.HGridY = int32(_dced & _c.MaxInt32)
	_dced, _afb = _eaac._baab.ReadBits(16)
	if _afb != nil {
		return _afb
	}
	_eaac.HRegionX = uint16(_dced & _c.MaxUint16)
	_dced, _afb = _eaac._baab.ReadBits(16)
	if _afb != nil {
		return _afb
	}
	_eaac.HRegionY = uint16(_dced & _c.MaxUint16)
	if _afb = _eaac.computeSegmentDataStructure(); _afb != nil {
		return _afb
	}
	return _eaac.checkInput()
}
func (_ebg *GenericRegion) decodeTemplate0a(_fad, _ebb, _cbda int, _fdgb, _egcf int) (_fcb error) {
	const _gea = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0061"
	var (
		_gcbe, _dgdb int
		_gdc, _aea   int
		_egb         byte
		_fggd        int
	)
	if _fad >= 1 {
		_egb, _fcb = _ebg.Bitmap.GetByte(_egcf)
		if _fcb != nil {
			return _fe.Wrap(_fcb, _gea, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_gdc = int(_egb)
	}
	if _fad >= 2 {
		_egb, _fcb = _ebg.Bitmap.GetByte(_egcf - _ebg.Bitmap.RowStride)
		if _fcb != nil {
			return _fe.Wrap(_fcb, _gea, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_aea = int(_egb) << 6
	}
	_gcbe = (_gdc & 0xf0) | (_aea & 0x3800)
	for _eed := 0; _eed < _cbda; _eed = _fggd {
		var (
			_fbdg byte
			_bbe  int
		)
		_fggd = _eed + 8
		if _egda := _ebb - _eed; _egda > 8 {
			_bbe = 8
		} else {
			_bbe = _egda
		}
		if _fad > 0 {
			_gdc <<= 8
			if _fggd < _ebb {
				_egb, _fcb = _ebg.Bitmap.GetByte(_egcf + 1)
				if _fcb != nil {
					return _fe.Wrap(_fcb, _gea, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_gdc |= int(_egb)
			}
		}
		if _fad > 1 {
			_ecf := _egcf - _ebg.Bitmap.RowStride + 1
			_aea <<= 8
			if _fggd < _ebb {
				_egb, _fcb = _ebg.Bitmap.GetByte(_ecf)
				if _fcb != nil {
					return _fe.Wrap(_fcb, _gea, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_aea |= int(_egb) << 6
			} else {
				_aea |= 0
			}
		}
		for _abde := 0; _abde < _bbe; _abde++ {
			_fadc := uint(7 - _abde)
			if _ebg._dca {
				_dgdb = _ebg.overrideAtTemplate0a(_gcbe, _eed+_abde, _fad, int(_fbdg), _abde, int(_fadc))
				_ebg._bfe.SetIndex(int32(_dgdb))
			} else {
				_ebg._bfe.SetIndex(int32(_gcbe))
			}
			var _bfab int
			_bfab, _fcb = _ebg._caf.DecodeBit(_ebg._bfe)
			if _fcb != nil {
				return _fe.Wrap(_fcb, _gea, "")
			}
			_fbdg |= byte(_bfab) << _fadc
			_gcbe = ((_gcbe & 0x7bf7) << 1) | _bfab | ((_gdc >> _fadc) & 0x10) | ((_aea >> _fadc) & 0x800)
		}
		if _ege := _ebg.Bitmap.SetByte(_fdgb, _fbdg); _ege != nil {
			return _fe.Wrap(_ege, _gea, "")
		}
		_fdgb++
		_egcf++
	}
	return nil
}
func (_gbbe *PageInformationSegment) encodeFlags(_cefe _ae.BinaryWriter) (_edggg error) {
	const _baaad = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _edggg = _cefe.SkipBits(1); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0072\u0065\u0073e\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074")
	}
	var _cfad int
	if _gbbe.CombinationOperatorOverrideAllowed() {
		_cfad = 1
	}
	if _edggg = _cefe.WriteBit(_cfad); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u006f\u0076\u0065\u0072\u0072\u0069\u0064\u0064\u0065\u006e")
	}
	_cfad = 0
	if _gbbe._fgaa {
		_cfad = 1
	}
	if _edggg = _cefe.WriteBit(_cfad); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0075\u0078i\u006ci\u0061r\u0079\u0020\u0062\u0075\u0066\u0066\u0065r")
	}
	if _edggg = _cefe.WriteBit((int(_gbbe._bca) >> 1) & 0x01); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0063\u006f\u006d\u0062\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006f\u0070\u0065\u0072\u0061\u0074o\u0072\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0062\u0069\u0074")
	}
	if _edggg = _cefe.WriteBit(int(_gbbe._bca) & 0x01); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0062\u0069\u0074")
	}
	_cfad = int(_gbbe.DefaultPixelValue)
	if _edggg = _cefe.WriteBit(_cfad); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0064e\u0066\u0061\u0075\u006c\u0074\u0020\u0070\u0061\u0067\u0065\u0020p\u0069\u0078\u0065\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	_cfad = 0
	if _gbbe._cadb {
		_cfad = 1
	}
	if _edggg = _cefe.WriteBit(_cfad); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "\u0063\u006f\u006e\u0074ai\u006e\u0073\u0020\u0072\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074")
	}
	_cfad = 0
	if _gbbe.IsLossless {
		_cfad = 1
	}
	if _edggg = _cefe.WriteBit(_cfad); _edggg != nil {
		return _fe.Wrap(_edggg, _baaad, "p\u0061\u0067\u0065\u0020\u0069\u0073 \u0065\u0076\u0065\u006e\u0074\u0075\u0061\u006c\u006cy\u0020\u006c\u006fs\u0073l\u0065\u0073\u0073")
	}
	return nil
}
func (_dgd *GenericRegion) copyLineAbove(_fdga int) error {
	_dee := _fdga * _dgd.Bitmap.RowStride
	_gde := _dee - _dgd.Bitmap.RowStride
	for _aadc := 0; _aadc < _dgd.Bitmap.RowStride; _aadc++ {
		_afa, _bgag := _dgd.Bitmap.GetByte(_gde)
		if _bgag != nil {
			return _bgag
		}
		_gde++
		if _bgag = _dgd.Bitmap.SetByte(_dee, _afa); _bgag != nil {
			return _bgag
		}
		_dee++
	}
	return nil
}

type GenericRegion struct {
	_abe             *_ae.Reader
	DataHeaderOffset int64
	DataHeaderLength int64
	DataOffset       int64
	DataLength       int64
	RegionSegment    *RegionSegment
	UseExtTemplates  bool
	IsTPGDon         bool
	GBTemplate       byte
	IsMMREncoded     bool
	UseMMR           bool
	GBAtX            []int8
	GBAtY            []int8
	GBAtOverride     []bool
	_dca             bool
	Bitmap           *_df.Bitmap
	_caf             *_b.Decoder
	_bfe             *_b.DecoderStats
	_cca             *_aad.Decoder
}

const (
	ORandom OrganizationType = iota
	OSequential
)

func (_dddc *RegionSegment) readCombinationOperator() error {
	_afca, _bedg := _dddc._gega.ReadBits(3)
	if _bedg != nil {
		return _bedg
	}
	_dddc.CombinaionOperator = _df.CombinationOperator(_afca & 0xF)
	return nil
}
func NewGenericRegion(r *_ae.Reader) *GenericRegion {
	return &GenericRegion{RegionSegment: NewRegionSegment(r), _abe: r}
}
func (_eaegb *SymbolDictionary) readRegionFlags() error {
	var (
		_eggf uint64
		_efgg int
	)
	_, _fgfa := _eaegb._bbda.ReadBits(3)
	if _fgfa != nil {
		return _fgfa
	}
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	_eaegb.SdrTemplate = int8(_efgg)
	_eggf, _fgfa = _eaegb._bbda.ReadBits(2)
	if _fgfa != nil {
		return _fgfa
	}
	_eaegb.SdTemplate = int8(_eggf & 0xf)
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	if _efgg == 1 {
		_eaegb._acgc = true
	}
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	if _efgg == 1 {
		_eaegb._dgeg = true
	}
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	if _efgg == 1 {
		_eaegb.SdHuffAggInstanceSelection = true
	}
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	_eaegb.SdHuffBMSizeSelection = int8(_efgg)
	_eggf, _fgfa = _eaegb._bbda.ReadBits(2)
	if _fgfa != nil {
		return _fgfa
	}
	_eaegb.SdHuffDecodeWidthSelection = int8(_eggf & 0xf)
	_eggf, _fgfa = _eaegb._bbda.ReadBits(2)
	if _fgfa != nil {
		return _fgfa
	}
	_eaegb.SdHuffDecodeHeightSelection = int8(_eggf & 0xf)
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	if _efgg == 1 {
		_eaegb.UseRefinementAggregation = true
	}
	_efgg, _fgfa = _eaegb._bbda.ReadBit()
	if _fgfa != nil {
		return _fgfa
	}
	if _efgg == 1 {
		_eaegb.IsHuffmanEncoded = true
	}
	return nil
}

type SymbolDictionary struct {
	_bbda                       *_ae.Reader
	SdrTemplate                 int8
	SdTemplate                  int8
	_acgc                       bool
	_dgeg                       bool
	SdHuffAggInstanceSelection  bool
	SdHuffBMSizeSelection       int8
	SdHuffDecodeWidthSelection  int8
	SdHuffDecodeHeightSelection int8
	UseRefinementAggregation    bool
	IsHuffmanEncoded            bool
	SdATX                       []int8
	SdATY                       []int8
	SdrATX                      []int8
	SdrATY                      []int8
	NumberOfExportedSymbols     uint32
	NumberOfNewSymbols          uint32
	Header                      *Header
	_eaab                       uint32
	_fege                       []*_df.Bitmap
	_dacc                       uint32
	_dede                       []*_df.Bitmap
	_fcdb                       _ge.Tabler
	_efec                       _ge.Tabler
	_efcb                       _ge.Tabler
	_fdgf                       _ge.Tabler
	_ddac                       []*_df.Bitmap
	_bccf                       []*_df.Bitmap
	_fdbf                       *_b.Decoder
	_ebe                        *TextRegion
	_fgbe                       *GenericRegion
	_abcc                       *GenericRefinementRegion
	_fbcd                       *_b.DecoderStats
	_cacfe                      *_b.DecoderStats
	_cdab                       *_b.DecoderStats
	_aebeg                      *_b.DecoderStats
	_bgdf                       *_b.DecoderStats
	_egdc                       *_b.DecoderStats
	_fcbb                       *_b.DecoderStats
	_agfa                       *_b.DecoderStats
	_ecc                        *_b.DecoderStats
	_efgc                       int8
	_daaa                       *_df.Bitmaps
	_face                       []int
	_feda                       map[int]int
	_afcf                       bool
}
type GenericRefinementRegion struct {
	_fa             templater
	_gd             templater
	_fd             *_ae.Reader
	_ce             *Header
	RegionInfo      *RegionSegment
	IsTPGROn        bool
	TemplateID      int8
	Template        templater
	GrAtX           []int8
	GrAtY           []int8
	RegionBitmap    *_df.Bitmap
	ReferenceBitmap *_df.Bitmap
	ReferenceDX     int32
	ReferenceDY     int32
	_eg             *_b.Decoder
	_ab             *_b.DecoderStats
	_bf             bool
	_ba             []bool
}

func (_gbge *SymbolDictionary) huffDecodeBmSize() (int64, error) {
	if _gbge._efcb == nil {
		var (
			_agge int
			_ggea error
		)
		if _gbge.SdHuffDecodeHeightSelection == 3 {
			_agge++
		}
		if _gbge.SdHuffDecodeWidthSelection == 3 {
			_agge++
		}
		_gbge._efcb, _ggea = _gbge.getUserTable(_agge)
		if _ggea != nil {
			return 0, _ggea
		}
	}
	return _gbge._efcb.Decode(_gbge._bbda)
}
func (_fdcd *TextRegion) initSymbols() error {
	const _cgca = "i\u006e\u0069\u0074\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	for _, _afea := range _fdcd.Header.RTSegments {
		if _afea == nil {
			return _fe.Error(_cgca, "\u006e\u0069\u006c\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020pr\u006f\u0076\u0069\u0064\u0065\u0064 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065g\u0069\u006f\u006e\u0020\u0053\u0079\u006d\u0062o\u006c\u0073")
		}
		if _afea.Type == 0 {
			_accfc, _fcbc := _afea.GetSegmentData()
			if _fcbc != nil {
				return _fe.Wrap(_fcbc, _cgca, "")
			}
			_dfbg, _dcbb := _accfc.(*SymbolDictionary)
			if !_dcbb {
				return _fe.Error(_cgca, "\u0072e\u0066\u0065r\u0072\u0065\u0064 \u0054\u006f\u0020\u0053\u0065\u0067\u006de\u006e\u0074\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			}
			_dfbg._ecc = _fdcd._bebb
			_agfge, _fcbc := _dfbg.GetDictionary()
			if _fcbc != nil {
				return _fe.Wrap(_fcbc, _cgca, "")
			}
			_fdcd.Symbols = append(_fdcd.Symbols, _agfge...)
		}
	}
	_fdcd.NumberOfSymbols = uint32(len(_fdcd.Symbols))
	return nil
}
func (_edf *GenericRegion) decodeSLTP() (int, error) {
	switch _edf.GBTemplate {
	case 0:
		_edf._bfe.SetIndex(0x9B25)
	case 1:
		_edf._bfe.SetIndex(0x795)
	case 2:
		_edf._bfe.SetIndex(0xE5)
	case 3:
		_edf._bfe.SetIndex(0x195)
	}
	return _edf._caf.DecodeBit(_edf._bfe)
}
func (_ddfe *TextRegion) decodeSymbolInstances() error {
	_cdggg, _ecdg := _ddfe.decodeStripT()
	if _ecdg != nil {
		return _ecdg
	}
	var (
		_dbafa int64
		_bdbb  uint32
	)
	for _bdbb < _ddfe.NumberOfSymbolInstances {
		_ebbe, _fdea := _ddfe.decodeDT()
		if _fdea != nil {
			return _fdea
		}
		_cdggg += _ebbe
		var _ceabd int64
		_adbea := true
		_ddfe._bcef = 0
		for {
			if _adbea {
				_ceabd, _fdea = _ddfe.decodeDfs()
				if _fdea != nil {
					return _fdea
				}
				_dbafa += _ceabd
				_ddfe._bcef = _dbafa
				_adbea = false
			} else {
				_dfad, _dafb := _ddfe.decodeIds()
				if _aa.Is(_dafb, _e.ErrOOB) {
					break
				}
				if _dafb != nil {
					return _dafb
				}
				if _bdbb >= _ddfe.NumberOfSymbolInstances {
					break
				}
				_ddfe._bcef += _dfad + int64(_ddfe.SbDsOffset)
			}
			_cbfb, _cffa := _ddfe.decodeCurrentT()
			if _cffa != nil {
				return _cffa
			}
			_cebgf := _cdggg + _cbfb
			_fdfb, _cffa := _ddfe.decodeID()
			if _cffa != nil {
				return _cffa
			}
			_agdd, _cffa := _ddfe.decodeRI()
			if _cffa != nil {
				return _cffa
			}
			_ddfgd, _cffa := _ddfe.decodeIb(_agdd, _fdfb)
			if _cffa != nil {
				return _cffa
			}
			if _cffa = _ddfe.blit(_ddfgd, _cebgf); _cffa != nil {
				return _cffa
			}
			_bdbb++
		}
	}
	return nil
}
func (_aaf *GenericRegion) getPixel(_cdfc, _gcba int) int8 {
	if _cdfc < 0 || _cdfc >= _aaf.Bitmap.Width {
		return 0
	}
	if _gcba < 0 || _gcba >= _aaf.Bitmap.Height {
		return 0
	}
	if _aaf.Bitmap.GetPixel(_cdfc, _gcba) {
		return 1
	}
	return 0
}
func (_fcab *template0) setIndex(_fdee *_b.DecoderStats) { _fdee.SetIndex(0x100) }
func (_bdfc *TextRegion) GetRegionInfo() *RegionSegment  { return _bdfc.RegionInfo }

var _ templater = &template0{}

type template0 struct{}

func (_faac *SymbolDictionary) decodeDirectlyThroughGenericRegion(_fged, _affab uint32) error {
	if _faac._fgbe == nil {
		_faac._fgbe = NewGenericRegion(_faac._bbda)
	}
	_faac._fgbe.setParametersWithAt(false, byte(_faac.SdTemplate), false, false, _faac.SdATX, _faac.SdATY, _fged, _affab, _faac._fbcd, _faac._fdbf)
	return _faac.addSymbol(_faac._fgbe)
}
func (_bde *TextRegion) Encode(w _ae.BinaryWriter) (_cgab int, _cbcb error) {
	const _eccd = "\u0054\u0065\u0078\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u002e\u0045n\u0063\u006f\u0064\u0065"
	if _cgab, _cbcb = _bde.RegionInfo.Encode(w); _cbcb != nil {
		return _cgab, _fe.Wrap(_cbcb, _eccd, "")
	}
	var _gee int
	if _gee, _cbcb = _bde.encodeFlags(w); _cbcb != nil {
		return _cgab, _fe.Wrap(_cbcb, _eccd, "")
	}
	_cgab += _gee
	if _gee, _cbcb = _bde.encodeSymbols(w); _cbcb != nil {
		return _cgab, _fe.Wrap(_cbcb, _eccd, "")
	}
	_cgab += _gee
	return _cgab, nil
}

type Header struct {
	SegmentNumber            uint32
	Type                     Type
	RetainFlag               bool
	PageAssociation          int
	PageAssociationFieldSize bool
	RTSegments               []*Header
	HeaderLength             int64
	SegmentDataLength        uint64
	SegmentDataStartOffset   uint64
	Reader                   *_ae.Reader
	SegmentData              Segmenter
	RTSNumbers               []int
	RetainBits               []uint8
}

func NewHeader(d Documenter, r *_ae.Reader, offset int64, organizationType OrganizationType) (*Header, error) {
	_gbbg := &Header{Reader: r}
	if _bega := _gbbg.parse(d, r, offset, organizationType); _bega != nil {
		return nil, _fe.Wrap(_bega, "\u004ee\u0077\u0048\u0065\u0061\u0064\u0065r", "")
	}
	return _gbbg, nil
}
func (_aafbg *TextRegion) decodeID() (int64, error) {
	if _aafbg.IsHuffmanEncoded {
		if _aafbg._bebbb == nil {
			_aggab, _edcda := _aafbg._deba.ReadBits(byte(_aafbg._adfd))
			return int64(_aggab), _edcda
		}
		return _aafbg._bebbb.Decode(_aafbg._deba)
	}
	return _aafbg._fagf.DecodeIAID(uint64(_aafbg._adfd), _aafbg._bebb)
}
func (_ac *EndOfStripe) Init(h *Header, r *_ae.Reader) error { _ac._bc = r; return _ac.parseHeader() }
func (_edac *GenericRegion) setParameters(_degg bool, _fea, _fdgg int64, _fabb, _gedd uint32) {
	_edac.IsMMREncoded = _degg
	_edac.DataOffset = _fea
	_edac.DataLength = _fdgg
	_edac.RegionSegment.BitmapHeight = _fabb
	_edac.RegionSegment.BitmapWidth = _gedd
	_edac._cca = nil
	_edac.Bitmap = nil
}
func (_bgg *GenericRefinementRegion) decodeTypicalPredictedLine(_gaa, _fb, _cea, _edc, _be, _feb int) error {
	_ceb := _gaa - int(_bgg.ReferenceDY)
	_bgf := _bgg.ReferenceBitmap.GetByteIndex(0, _ceb)
	_cgg := _bgg.RegionBitmap.GetByteIndex(0, _gaa)
	var _gg error
	switch _bgg.TemplateID {
	case 0:
		_gg = _bgg.decodeTypicalPredictedLineTemplate0(_gaa, _fb, _cea, _edc, _be, _feb, _cgg, _ceb, _bgf)
	case 1:
		_gg = _bgg.decodeTypicalPredictedLineTemplate1(_gaa, _fb, _cea, _edc, _be, _feb, _cgg, _ceb, _bgf)
	}
	return _gg
}
func (_cecc *SymbolDictionary) addSymbol(_aeef Regioner) error {
	_abfd, _cbdfg := _aeef.GetRegionBitmap()
	if _cbdfg != nil {
		return _cbdfg
	}
	_cecc._dede[_cecc._dacc] = _abfd
	_cecc._bccf = append(_cecc._bccf, _abfd)
	_af.Log.Trace("\u005b\u0053YM\u0042\u004f\u004c \u0044\u0049\u0043\u0054ION\u0041RY\u005d\u0020\u0041\u0064\u0064\u0065\u0064 s\u0079\u006d\u0062\u006f\u006c\u003a\u0020%\u0073", _abfd)
	return nil
}
func (_gbc *PageInformationSegment) readCombinationOperator() error {
	_egeb, _efa := _gbc._fgee.ReadBits(2)
	if _efa != nil {
		return _efa
	}
	_gbc._bca = _df.CombinationOperator(int(_egeb))
	return nil
}
func (_bgcd *Header) readHeaderFlags() error {
	const _agcc = "\u0072e\u0061d\u0048\u0065\u0061\u0064\u0065\u0072\u0046\u006c\u0061\u0067\u0073"
	_gdae, _cgcc := _bgcd.Reader.ReadBit()
	if _cgcc != nil {
		return _fe.Wrap(_cgcc, _agcc, "r\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	if _gdae != 0 {
		_bgcd.RetainFlag = true
	}
	_gdae, _cgcc = _bgcd.Reader.ReadBit()
	if _cgcc != nil {
		return _fe.Wrap(_cgcc, _agcc, "\u0070\u0061g\u0065\u0020\u0061s\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e")
	}
	if _gdae != 0 {
		_bgcd.PageAssociationFieldSize = true
	}
	_eaad, _cgcc := _bgcd.Reader.ReadBits(6)
	if _cgcc != nil {
		return _fe.Wrap(_cgcc, _agcc, "\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
	}
	_bgcd.Type = Type(int(_eaad))
	return nil
}
func (_bfb *GenericRegion) String() string {
	_fgd := &_a.Builder{}
	_fgd.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_fgd.WriteString(_bfb.RegionSegment.String() + "\u000a")
	_fgd.WriteString(_dg.Sprintf("\u0009\u002d\u0020Us\u0065\u0045\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _bfb.UseExtTemplates))
	_fgd.WriteString(_dg.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0044\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _bfb.IsTPGDon))
	_fgd.WriteString(_dg.Sprintf("\u0009-\u0020G\u0042\u0054\u0065\u006d\u0070l\u0061\u0074e\u003a\u0020\u0025\u0064\u000a", _bfb.GBTemplate))
	_fgd.WriteString(_dg.Sprintf("\t\u002d \u0049\u0073\u004d\u004d\u0052\u0045\u006e\u0063o\u0064\u0065\u0064\u003a %\u0076\u000a", _bfb.IsMMREncoded))
	_fgd.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _bfb.GBAtX))
	_fgd.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _bfb.GBAtY))
	_fgd.WriteString(_dg.Sprintf("\t\u002d \u0047\u0042\u0041\u0074\u004f\u0076\u0065\u0072r\u0069\u0064\u0065\u003a %\u0076\u000a", _bfb.GBAtOverride))
	return _fgd.String()
}
func (_gdd *SymbolDictionary) decodeDifferenceWidth() (int64, error) {
	if _gdd.IsHuffmanEncoded {
		switch _gdd.SdHuffDecodeWidthSelection {
		case 0:
			_edbd, _egeg := _ge.GetStandardTable(2)
			if _egeg != nil {
				return 0, _egeg
			}
			return _edbd.Decode(_gdd._bbda)
		case 1:
			_gcede, _gfdb := _ge.GetStandardTable(3)
			if _gfdb != nil {
				return 0, _gfdb
			}
			return _gcede.Decode(_gdd._bbda)
		case 3:
			if _gdd._efec == nil {
				var _gfcd int
				if _gdd.SdHuffDecodeHeightSelection == 3 {
					_gfcd++
				}
				_cdde, _bedf := _gdd.getUserTable(_gfcd)
				if _bedf != nil {
					return 0, _bedf
				}
				_gdd._efec = _cdde
			}
			return _gdd._efec.Decode(_gdd._bbda)
		}
	} else {
		_bbgf, _fcbe := _gdd._fdbf.DecodeInt(_gdd._cdab)
		if _fcbe != nil {
			return 0, _fcbe
		}
		return int64(_bbgf), nil
	}
	return 0, nil
}
func (_egff *SymbolDictionary) Init(h *Header, r *_ae.Reader) error {
	_egff.Header = h
	_egff._bbda = r
	return _egff.parseHeader()
}

type OrganizationType uint8

func (_bcgf *Header) Encode(w _ae.BinaryWriter) (_afg int, _fede error) {
	const _gafc = "\u0048\u0065\u0061d\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065"
	var _aegb _ae.BinaryWriter
	_af.Log.Trace("\u005b\u0053\u0045G\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045\u0052\u005d[\u0045\u004e\u0043\u004f\u0044\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _fede != nil {
			_af.Log.Trace("[\u0053\u0045\u0047\u004d\u0045\u004eT\u002d\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0045\u004e\u0043\u004f\u0044E\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020%\u0076", _fede)
		} else {
			_af.Log.Trace("\u005b\u0053\u0045\u0047ME\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0025\u0076", _bcgf)
			_af.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045N\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u005b\u0045\u004e\u0043O\u0044\u0045\u005d\u0020\u0046\u0069\u006ei\u0073\u0068\u0065\u0064")
		}
	}()
	w.FinishByte()
	if _bcgf.SegmentData != nil {
		_cfg, _dbff := _bcgf.SegmentData.(SegmentEncoder)
		if !_dbff {
			return 0, _fe.Errorf(_gafc, "\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0064\u006f\u0065s\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074 \u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0045\u006e\u0063\u006f\u0064er\u0020\u0069\u006e\u0074\u0065\u0072\u0066\u0061\u0063\u0065", _bcgf.SegmentData)
		}
		_aegb = _ae.BufferedMSB()
		_afg, _fede = _cfg.Encode(_aegb)
		if _fede != nil {
			return 0, _fe.Wrap(_fede, _gafc, "")
		}
		_bcgf.SegmentDataLength = uint64(_afg)
	}
	if _bcgf.pageSize() == 4 {
		_bcgf.PageAssociationFieldSize = true
	}
	var _fegb int
	_fegb, _fede = _bcgf.writeSegmentNumber(w)
	if _fede != nil {
		return 0, _fe.Wrap(_fede, _gafc, "")
	}
	_afg += _fegb
	if _fede = _bcgf.writeFlags(w); _fede != nil {
		return _afg, _fe.Wrap(_fede, _gafc, "")
	}
	_afg++
	_fegb, _fede = _bcgf.writeReferredToCount(w)
	if _fede != nil {
		return 0, _fe.Wrap(_fede, _gafc, "")
	}
	_afg += _fegb
	_fegb, _fede = _bcgf.writeReferredToSegments(w)
	if _fede != nil {
		return 0, _fe.Wrap(_fede, _gafc, "")
	}
	_afg += _fegb
	_fegb, _fede = _bcgf.writeSegmentPageAssociation(w)
	if _fede != nil {
		return 0, _fe.Wrap(_fede, _gafc, "")
	}
	_afg += _fegb
	_fegb, _fede = _bcgf.writeSegmentDataLength(w)
	if _fede != nil {
		return 0, _fe.Wrap(_fede, _gafc, "")
	}
	_afg += _fegb
	_bcgf.HeaderLength = int64(_afg) - int64(_bcgf.SegmentDataLength)
	if _aegb != nil {
		if _, _fede = w.Write(_aegb.Data()); _fede != nil {
			return _afg, _fe.Wrap(_fede, _gafc, "\u0077r\u0069t\u0065\u0020\u0073\u0065\u0067m\u0065\u006et\u0020\u0064\u0061\u0074\u0061")
		}
	}
	return _afg, nil
}
func (_caea *TextRegion) decodeIb(_fbfgb, _ecacb int64) (*_df.Bitmap, error) {
	const _fggdc = "\u0064\u0065\u0063\u006f\u0064\u0065\u0049\u0062"
	var (
		_bafgf error
		_dfde  *_df.Bitmap
	)
	if _fbfgb == 0 {
		if int(_ecacb) > len(_caea.Symbols)-1 {
			return nil, _fe.Error(_fggdc, "\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0049\u0042\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0020\u0069\u006e\u0064\u0065x\u0020\u006f\u0075\u0074\u0020o\u0066\u0020r\u0061\u006e\u0067\u0065")
		}
		return _caea.Symbols[int(_ecacb)], nil
	}
	var _bege, _gbbb, _ccga, _dfada int64
	_bege, _bafgf = _caea.decodeRdw()
	if _bafgf != nil {
		return nil, _fe.Wrap(_bafgf, _fggdc, "")
	}
	_gbbb, _bafgf = _caea.decodeRdh()
	if _bafgf != nil {
		return nil, _fe.Wrap(_bafgf, _fggdc, "")
	}
	_ccga, _bafgf = _caea.decodeRdx()
	if _bafgf != nil {
		return nil, _fe.Wrap(_bafgf, _fggdc, "")
	}
	_dfada, _bafgf = _caea.decodeRdy()
	if _bafgf != nil {
		return nil, _fe.Wrap(_bafgf, _fggdc, "")
	}
	if _caea.IsHuffmanEncoded {
		if _, _bafgf = _caea.decodeSymInRefSize(); _bafgf != nil {
			return nil, _fe.Wrap(_bafgf, _fggdc, "")
		}
		_caea._deba.Align()
	}
	_aggc := _caea.Symbols[_ecacb]
	_gcff := uint32(_aggc.Width)
	_aabg := uint32(_aggc.Height)
	_ceaa := int32(uint32(_bege)>>1) + int32(_ccga)
	_ddec := int32(uint32(_gbbb)>>1) + int32(_dfada)
	if _caea._gefd == nil {
		_caea._gefd = _eea(_caea._deba, nil)
	}
	_caea._gefd.setParameters(_caea._aaag, _caea._fagf, _caea.SbrTemplate, _gcff+uint32(_bege), _aabg+uint32(_gbbb), _aggc, _ceaa, _ddec, false, _caea.SbrATX, _caea.SbrATY)
	_dfde, _bafgf = _caea._gefd.GetRegionBitmap()
	if _bafgf != nil {
		return nil, _fe.Wrap(_bafgf, _fggdc, "\u0067\u0072\u0066")
	}
	if _caea.IsHuffmanEncoded {
		_caea._deba.Align()
	}
	return _dfde, nil
}
func (_cbgfc *Header) writeReferredToCount(_ddbg _ae.BinaryWriter) (_fgbd int, _fgbce error) {
	const _eacc = "w\u0072i\u0074\u0065\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0043ou\u006e\u0074"
	_cbgfc.RTSNumbers = make([]int, len(_cbgfc.RTSegments))
	for _beee, _faag := range _cbgfc.RTSegments {
		_cbgfc.RTSNumbers[_beee] = int(_faag.SegmentNumber)
	}
	if len(_cbgfc.RTSNumbers) <= 4 {
		var _fdfg byte
		if len(_cbgfc.RetainBits) >= 1 {
			_fdfg = _cbgfc.RetainBits[0]
		}
		_fdfg |= byte(len(_cbgfc.RTSNumbers)) << 5
		if _fgbce = _ddbg.WriteByte(_fdfg); _fgbce != nil {
			return 0, _fe.Wrap(_fgbce, _eacc, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		return 1, nil
	}
	_adgea := uint32(len(_cbgfc.RTSNumbers))
	_ecg := make([]byte, 4+_fgc.Ceil(len(_cbgfc.RTSNumbers)+1, 8))
	_adgea |= 0x7 << 29
	_d.BigEndian.PutUint32(_ecg, _adgea)
	copy(_ecg[1:], _cbgfc.RetainBits)
	_fgbd, _fgbce = _ddbg.Write(_ecg)
	if _fgbce != nil {
		return 0, _fe.Wrap(_fgbce, _eacc, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	return _fgbd, nil
}
func (_gegg *GenericRegion) decodeTemplate0b(_fagd, _aggf, _fge int, _fcc, _gadc int) (_bcd error) {
	const _eef = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0062"
	var (
		_age, _dfec int
		_eafb, _edb int
		_ead        byte
		_abfb       int
	)
	if _fagd >= 1 {
		_ead, _bcd = _gegg.Bitmap.GetByte(_gadc)
		if _bcd != nil {
			return _fe.Wrap(_bcd, _eef, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_eafb = int(_ead)
	}
	if _fagd >= 2 {
		_ead, _bcd = _gegg.Bitmap.GetByte(_gadc - _gegg.Bitmap.RowStride)
		if _bcd != nil {
			return _fe.Wrap(_bcd, _eef, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_edb = int(_ead) << 6
	}
	_age = (_eafb & 0xf0) | (_edb & 0x3800)
	for _bda := 0; _bda < _fge; _bda = _abfb {
		var (
			_bfgd byte
			_gcbd int
		)
		_abfb = _bda + 8
		if _edfd := _aggf - _bda; _edfd > 8 {
			_gcbd = 8
		} else {
			_gcbd = _edfd
		}
		if _fagd > 0 {
			_eafb <<= 8
			if _abfb < _aggf {
				_ead, _bcd = _gegg.Bitmap.GetByte(_gadc + 1)
				if _bcd != nil {
					return _fe.Wrap(_bcd, _eef, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_eafb |= int(_ead)
			}
		}
		if _fagd > 1 {
			_edb <<= 8
			if _abfb < _aggf {
				_ead, _bcd = _gegg.Bitmap.GetByte(_gadc - _gegg.Bitmap.RowStride + 1)
				if _bcd != nil {
					return _fe.Wrap(_bcd, _eef, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_edb |= int(_ead) << 6
			}
		}
		for _fbe := 0; _fbe < _gcbd; _fbe++ {
			_ccc := uint(7 - _fbe)
			if _gegg._dca {
				_dfec = _gegg.overrideAtTemplate0b(_age, _bda+_fbe, _fagd, int(_bfgd), _fbe, int(_ccc))
				_gegg._bfe.SetIndex(int32(_dfec))
			} else {
				_gegg._bfe.SetIndex(int32(_age))
			}
			var _cae int
			_cae, _bcd = _gegg._caf.DecodeBit(_gegg._bfe)
			if _bcd != nil {
				return _fe.Wrap(_bcd, _eef, "")
			}
			_bfgd |= byte(_cae << _ccc)
			_age = ((_age & 0x7bf7) << 1) | _cae | ((_eafb >> _ccc) & 0x10) | ((_edb >> _ccc) & 0x800)
		}
		if _eddc := _gegg.Bitmap.SetByte(_fcc, _bfgd); _eddc != nil {
			return _fe.Wrap(_eddc, _eef, "")
		}
		_fcc++
		_gadc++
	}
	return nil
}
func (_ccda *PatternDictionary) computeSegmentDataStructure() error {
	_ccda.DataOffset = _ccda._fdcf.AbsolutePosition()
	_ccda.DataHeaderLength = _ccda.DataOffset - _ccda.DataHeaderOffset
	_ccda.DataLength = int64(_ccda._fdcf.AbsoluteLength()) - _ccda.DataHeaderLength
	return nil
}
func (_fgda *TextRegion) Init(header *Header, r *_ae.Reader) error {
	_fgda.Header = header
	_fgda._deba = r
	_fgda.RegionInfo = NewRegionSegment(_fgda._deba)
	return _fgda.parseHeader()
}
func _cegb(_cbf int) int {
	if _cbf == 0 {
		return 0
	}
	_cbf |= _cbf >> 1
	_cbf |= _cbf >> 2
	_cbf |= _cbf >> 4
	_cbf |= _cbf >> 8
	_cbf |= _cbf >> 16
	return (_cbf + 1) >> 1
}
func (_fbbc *Header) writeFlags(_edcga _ae.BinaryWriter) (_gdgd error) {
	const _dbae = "\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065F\u006c\u0061\u0067\u0073"
	_eafc := byte(_fbbc.Type)
	if _gdgd = _edcga.WriteByte(_eafc); _gdgd != nil {
		return _fe.Wrap(_gdgd, _dbae, "\u0077\u0072\u0069ti\u006e\u0067\u0020\u0073\u0065\u0067\u006d\u0065\u006et\u0020t\u0079p\u0065 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	if !_fbbc.RetainFlag && !_fbbc.PageAssociationFieldSize {
		return nil
	}
	if _gdgd = _edcga.SkipBits(-8); _gdgd != nil {
		return _fe.Wrap(_gdgd, _dbae, "\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0061\u0063\u006b\u0020\u0074\u0068\u0065\u0020\u0062\u0069\u0074\u0073\u0020\u0066\u0061il\u0065\u0064")
	}
	var _gcaf int
	if _fbbc.RetainFlag {
		_gcaf = 1
	}
	if _gdgd = _edcga.WriteBit(_gcaf); _gdgd != nil {
		return _fe.Wrap(_gdgd, _dbae, "\u0072\u0065\u0074\u0061in\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067\u0073")
	}
	_gcaf = 0
	if _fbbc.PageAssociationFieldSize {
		_gcaf = 1
	}
	if _gdgd = _edcga.WriteBit(_gcaf); _gdgd != nil {
		return _fe.Wrap(_gdgd, _dbae, "p\u0061\u0067\u0065\u0020as\u0073o\u0063\u0069\u0061\u0074\u0069o\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	_edcga.FinishByte()
	return nil
}
func (_fbcfd *SymbolDictionary) encodeSymbols(_fddce _ae.BinaryWriter) (_aged int, _cggd error) {
	const _dagc = "\u0065\u006e\u0063o\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c"
	_eega := _gc.New()
	_eega.Init()
	_fcag, _cggd := _fbcfd._daaa.SelectByIndexes(_fbcfd._face)
	if _cggd != nil {
		return 0, _fe.Wrap(_cggd, _dagc, "\u0069n\u0069\u0074\u0069\u0061\u006c")
	}
	_cebf := map[*_df.Bitmap]int{}
	for _dgcf, _gaff := range _fcag.Values {
		_cebf[_gaff] = _dgcf
	}
	_fcag.SortByHeight()
	var _aaec, _cfgd int
	_eaag, _cggd := _fcag.GroupByHeight()
	if _cggd != nil {
		return 0, _fe.Wrap(_cggd, _dagc, "")
	}
	for _, _gfbg := range _eaag.Values {
		_fbff := _gfbg.Values[0].Height
		_adag := _fbff - _aaec
		if _cggd = _eega.EncodeInteger(_gc.IADH, _adag); _cggd != nil {
			return 0, _fe.Wrapf(_cggd, _dagc, "\u0049\u0041\u0044\u0048\u0020\u0066\u006f\u0072\u0020\u0064\u0068\u003a \u0027\u0025\u0064\u0027", _adag)
		}
		_aaec = _fbff
		_bec, _aege := _gfbg.GroupByWidth()
		if _aege != nil {
			return 0, _fe.Wrapf(_aege, _dagc, "\u0068\u0065\u0069g\u0068\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _fbff)
		}
		var _gebb int
		for _, _fedaa := range _bec.Values {
			for _, _cecf := range _fedaa.Values {
				_dfdg := _cecf.Width
				_aeca := _dfdg - _gebb
				if _aege = _eega.EncodeInteger(_gc.IADW, _aeca); _aege != nil {
					return 0, _fe.Wrapf(_aege, _dagc, "\u0049\u0041\u0044\u0057\u0020\u0066\u006f\u0072\u0020\u0064\u0077\u003a \u0027\u0025\u0064\u0027", _aeca)
				}
				_gebb += _aeca
				if _aege = _eega.EncodeBitmap(_cecf, false); _aege != nil {
					return 0, _fe.Wrapf(_aege, _dagc, "H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u0020W\u0069\u0064\u0074\u0068: \u0025\u0064", _fbff, _dfdg)
				}
				_egbe := _cebf[_cecf]
				_fbcfd._feda[_egbe] = _cfgd
				_cfgd++
			}
		}
		if _aege = _eega.EncodeOOB(_gc.IADW); _aege != nil {
			return 0, _fe.Wrap(_aege, _dagc, "\u0049\u0041\u0044\u0057")
		}
	}
	if _cggd = _eega.EncodeInteger(_gc.IAEX, 0); _cggd != nil {
		return 0, _fe.Wrap(_cggd, _dagc, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	if _cggd = _eega.EncodeInteger(_gc.IAEX, len(_fbcfd._face)); _cggd != nil {
		return 0, _fe.Wrap(_cggd, _dagc, "\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0073\u0079m\u0062\u006f\u006c\u0073")
	}
	_eega.Final()
	_gcbc, _cggd := _eega.WriteTo(_fddce)
	if _cggd != nil {
		return 0, _fe.Wrap(_cggd, _dagc, "\u0077\u0072i\u0074\u0069\u006e\u0067 \u0065\u006ec\u006f\u0064\u0065\u0072\u0020\u0063\u006f\u006et\u0065\u0078\u0074\u0020\u0074\u006f\u0020\u0027\u0077\u0027\u0020\u0077r\u0069\u0074\u0065\u0072")
	}
	return int(_gcbc), nil
}

var (
	_ Regioner  = &TextRegion{}
	_ Segmenter = &TextRegion{}
)

func (_fff *RegionSegment) String() string {
	_begc := &_a.Builder{}
	_begc.WriteString("\u0009[\u0052E\u0047\u0049\u004f\u004e\u0020S\u0045\u0047M\u0045\u004e\u0054\u005d\u000a")
	_begc.WriteString(_dg.Sprintf("\t\u0009\u002d\u0020\u0042\u0069\u0074m\u0061\u0070\u0020\u0028\u0077\u0069d\u0074\u0068\u002c\u0020\u0068\u0065\u0069g\u0068\u0074\u0029\u0020\u005b\u0025\u0064\u0078\u0025\u0064]\u000a", _fff.BitmapWidth, _fff.BitmapHeight))
	_begc.WriteString(_dg.Sprintf("\u0009\u0009\u002d\u0020L\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0028\u0078,\u0079)\u003a\u0020\u005b\u0025\u0064\u002c\u0025d\u005d\u000a", _fff.XLocation, _fff.YLocation))
	_begc.WriteString(_dg.Sprintf("\t\u0009\u002d\u0020\u0043\u006f\u006db\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0073", _fff.CombinaionOperator))
	return _begc.String()
}
func (_acc *GenericRefinementRegion) String() string {
	_fbcf := &_a.Builder{}
	_fbcf.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_fbcf.WriteString(_acc.RegionInfo.String() + "\u000a")
	_fbcf.WriteString(_dg.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0052\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _acc.IsTPGROn))
	_fbcf.WriteString(_dg.Sprintf("\u0009-\u0020T\u0065\u006d\u0070\u006c\u0061t\u0065\u0049D\u003a\u0020\u0025\u0076\u000a", _acc.TemplateID))
	_fbcf.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _acc.GrAtX))
	_fbcf.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _acc.GrAtY))
	_fbcf.WriteString(_dg.Sprintf("\u0009-\u0020R\u0065\u0066\u0065\u0072\u0065n\u0063\u0065D\u0058\u0020\u0025\u0076\u000a", _acc.ReferenceDX))
	_fbcf.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0052ef\u0065\u0072\u0065\u006e\u0063\u0044\u0065\u0059\u003a\u0020\u0025\u0076\u000a", _acc.ReferenceDY))
	return _fbcf.String()
}
func (_bced *Header) readSegmentDataLength(_befc *_ae.Reader) (_aedd error) {
	_bced.SegmentDataLength, _aedd = _befc.ReadBits(32)
	if _aedd != nil {
		return _aedd
	}
	_bced.SegmentDataLength &= _c.MaxInt32
	return nil
}
func (_fbccb *TextRegion) decodeRI() (int64, error) {
	if !_fbccb.UseRefinement {
		return 0, nil
	}
	if _fbccb.IsHuffmanEncoded {
		_egdf, _edgfc := _fbccb._deba.ReadBit()
		return int64(_egdf), _edgfc
	}
	_ccbb, _dceg := _fbccb._fagf.DecodeInt(_fbccb._gddf)
	return int64(_ccbb), _dceg
}
func (_eedf *PageInformationSegment) CombinationOperator() _df.CombinationOperator { return _eedf._bca }
func (_fdbab *SymbolDictionary) setExportedSymbols(_efagg []int) {
	for _ccgd := uint32(0); _ccgd < _fdbab._eaab+_fdbab.NumberOfNewSymbols; _ccgd++ {
		if _efagg[_ccgd] == 1 {
			var _fccc *_df.Bitmap
			if _ccgd < _fdbab._eaab {
				_fccc = _fdbab._fege[_ccgd]
			} else {
				_fccc = _fdbab._dede[_ccgd-_fdbab._eaab]
			}
			_af.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u002d\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u0020\u0041\u0064\u0064 E\u0078\u0070\u006fr\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027%\u0073\u0027", _fccc)
			_fdbab._ddac = append(_fdbab._ddac, _fccc)
		}
	}
}
func (_cgccg *SymbolDictionary) encodeFlags(_efaa _ae.BinaryWriter) (_cedfc int, _egfe error) {
	const _ffafe = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _egfe = _efaa.SkipBits(3); _egfe != nil {
		return 0, _fe.Wrap(_egfe, _ffafe, "\u0065\u006d\u0070\u0074\u0079\u0020\u0062\u0069\u0074\u0073")
	}
	var _edffg int
	if _cgccg.SdrTemplate > 0 {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "s\u0064\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_edffg = 0
	if _cgccg.SdTemplate > 1 {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_edffg = 0
	if _cgccg.SdTemplate == 1 || _cgccg.SdTemplate == 3 {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_edffg = 0
	if _cgccg._acgc {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0063\u006f\u0064in\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0065\u0064")
	}
	_edffg = 0
	if _cgccg._dgeg {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0063\u006f\u0064\u0069ng\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0075\u0073\u0065\u0064")
	}
	_edffg = 0
	if _cgccg.SdHuffAggInstanceSelection {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0048\u0075\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074")
	}
	_edffg = int(_cgccg.SdHuffBMSizeSelection)
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0048u\u0066\u0066\u0042\u006d\u0053\u0069\u007a\u0065")
	}
	_edffg = 0
	if _cgccg.SdHuffDecodeWidthSelection > 1 {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_edffg = 0
	switch _cgccg.SdHuffDecodeWidthSelection {
	case 1, 3:
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_edffg = 0
	if _cgccg.SdHuffDecodeHeightSelection > 1 {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_edffg = 0
	switch _cgccg.SdHuffDecodeHeightSelection {
	case 1, 3:
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_edffg = 0
	if _cgccg.UseRefinementAggregation {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0052\u0065\u0066\u0041\u0067\u0067")
	}
	_edffg = 0
	if _cgccg.IsHuffmanEncoded {
		_edffg = 1
	}
	if _egfe = _efaa.WriteBit(_edffg); _egfe != nil {
		return _cedfc, _fe.Wrap(_egfe, _ffafe, "\u0073\u0064\u0048\u0075\u0066\u0066")
	}
	return 2, nil
}
func (_faaa *GenericRefinementRegion) readAtPixels() error {
	_faaa.GrAtX = make([]int8, 2)
	_faaa.GrAtY = make([]int8, 2)
	_abcf, _dfd := _faaa._fd.ReadByte()
	if _dfd != nil {
		return _dfd
	}
	_faaa.GrAtX[0] = int8(_abcf)
	_abcf, _dfd = _faaa._fd.ReadByte()
	if _dfd != nil {
		return _dfd
	}
	_faaa.GrAtY[0] = int8(_abcf)
	_abcf, _dfd = _faaa._fd.ReadByte()
	if _dfd != nil {
		return _dfd
	}
	_faaa.GrAtX[1] = int8(_abcf)
	_abcf, _dfd = _faaa._fd.ReadByte()
	if _dfd != nil {
		return _dfd
	}
	_faaa.GrAtY[1] = int8(_abcf)
	return nil
}

type Documenter interface {
	GetPage(int) (Pager, error)
	GetGlobalSegment(int) (*Header, error)
}

func (_gaaab *TextRegion) InitEncode(globalSymbolsMap, localSymbolsMap map[int]int, comps []int, inLL *_df.Points, symbols *_df.Bitmaps, classIDs *_fgc.IntSlice, boxes *_df.Boxes, width, height, symBits int) {
	_gaaab.RegionInfo = &RegionSegment{BitmapWidth: uint32(width), BitmapHeight: uint32(height)}
	_gaaab._eeaa = globalSymbolsMap
	_gaaab._gecbc = localSymbolsMap
	_gaaab._faba = comps
	_gaaab._eegg = inLL
	_gaaab._gdffg = symbols
	_gaaab._cdbe = classIDs
	_gaaab._fgbef = boxes
	_gaaab._dccc = symBits
}
func (_ecec *RegionSegment) Size() int { return 17 }
func (_cff *SymbolDictionary) decodeRefinedSymbol(_eaec, _aafb uint32) error {
	var (
		_agd         int
		_efag, _bbbg int32
	)
	if _cff.IsHuffmanEncoded {
		_afaec, _ggcd := _cff._bbda.ReadBits(byte(_cff._efgc))
		if _ggcd != nil {
			return _ggcd
		}
		_agd = int(_afaec)
		_bcddd, _ggcd := _ge.GetStandardTable(15)
		if _ggcd != nil {
			return _ggcd
		}
		_dafg, _ggcd := _bcddd.Decode(_cff._bbda)
		if _ggcd != nil {
			return _ggcd
		}
		_efag = int32(_dafg)
		_dafg, _ggcd = _bcddd.Decode(_cff._bbda)
		if _ggcd != nil {
			return _ggcd
		}
		_bbbg = int32(_dafg)
		_bcddd, _ggcd = _ge.GetStandardTable(1)
		if _ggcd != nil {
			return _ggcd
		}
		if _, _ggcd = _bcddd.Decode(_cff._bbda); _ggcd != nil {
			return _ggcd
		}
		_cff._bbda.Align()
	} else {
		_fbbg, _ggge := _cff._fdbf.DecodeIAID(uint64(_cff._efgc), _cff._ecc)
		if _ggge != nil {
			return _ggge
		}
		_agd = int(_fbbg)
		_efag, _ggge = _cff._fdbf.DecodeInt(_cff._egdc)
		if _ggge != nil {
			return _ggge
		}
		_bbbg, _ggge = _cff._fdbf.DecodeInt(_cff._fcbb)
		if _ggge != nil {
			return _ggge
		}
	}
	if _gaed := _cff.setSymbolsArray(); _gaed != nil {
		return _gaed
	}
	_dcac := _cff._bccf[_agd]
	if _baaf := _cff.decodeNewSymbols(_eaec, _aafb, _dcac, _efag, _bbbg); _baaf != nil {
		return _baaf
	}
	if _cff.IsHuffmanEncoded {
		_cff._bbda.Align()
	}
	return nil
}
func (_cced *PageInformationSegment) parseHeader() (_eedb error) {
	_af.Log.Trace("\u005b\u0050\u0061\u0067\u0065I\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067m\u0065\u006e\u0074\u005d\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u002e\u002e")
	defer func() {
		var _daab = "[\u0050\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0053\u0065gm\u0065\u006e\u0074\u005d \u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065ad\u0065\u0072 \u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064"
		if _eedb != nil {
			_daab += "\u0020\u0077\u0069t\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u0020" + _eedb.Error()
		} else {
			_daab += "\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079"
		}
		_af.Log.Trace(_daab)
	}()
	if _eedb = _cced.readWidthAndHeight(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readResolution(); _eedb != nil {
		return _eedb
	}
	_, _eedb = _cced._fgee.ReadBit()
	if _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readCombinationOperatorOverrideAllowed(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readRequiresAuxiliaryBuffer(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readCombinationOperator(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readDefaultPixelValue(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readContainsRefinement(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readIsLossless(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readIsStriped(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.readMaxStripeSize(); _eedb != nil {
		return _eedb
	}
	if _eedb = _cced.checkInput(); _eedb != nil {
		return _eedb
	}
	_af.Log.Trace("\u0025\u0073", _cced)
	return nil
}
func (_abd *GenericRefinementRegion) GetRegionInfo() *RegionSegment { return _abd.RegionInfo }
func (_adcc *PatternDictionary) readGrayMax() error {
	_dbaf, _bggbb := _adcc._fdcf.ReadBits(32)
	if _bggbb != nil {
		return _bggbb
	}
	_adcc.GrayMax = uint32(_dbaf & _c.MaxUint32)
	return nil
}
func (_cdee *SymbolDictionary) readNumberOfExportedSymbols() error {
	_eabd, _cdade := _cdee._bbda.ReadBits(32)
	if _cdade != nil {
		return _cdade
	}
	_cdee.NumberOfExportedSymbols = uint32(_eabd & _c.MaxUint32)
	return nil
}
func (_aae *SymbolDictionary) encodeNumSyms(_aadg _ae.BinaryWriter) (_deffg int, _afbg error) {
	const _cggea = "\u0065\u006e\u0063\u006f\u0064\u0065\u004e\u0075\u006d\u0053\u0079\u006d\u0073"
	_cgccf := make([]byte, 4)
	_d.BigEndian.PutUint32(_cgccf, _aae.NumberOfExportedSymbols)
	if _deffg, _afbg = _aadg.Write(_cgccf); _afbg != nil {
		return _deffg, _fe.Wrap(_afbg, _cggea, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	_d.BigEndian.PutUint32(_cgccf, _aae.NumberOfNewSymbols)
	_fdce, _afbg := _aadg.Write(_cgccf)
	if _afbg != nil {
		return _deffg, _fe.Wrap(_afbg, _cggea, "n\u0065\u0077\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	return _deffg + _fdce, nil
}
func (_gbda *GenericRegion) overrideAtTemplate0a(_gceg, _ddc, _bgd, _ddb, _aaad, _dgf int) int {
	if _gbda.GBAtOverride[0] {
		_gceg &= 0xFFEF
		if _gbda.GBAtY[0] == 0 && _gbda.GBAtX[0] >= -int8(_aaad) {
			_gceg |= (_ddb >> uint(int8(_dgf)-_gbda.GBAtX[0]&0x1)) << 4
		} else {
			_gceg |= int(_gbda.getPixel(_ddc+int(_gbda.GBAtX[0]), _bgd+int(_gbda.GBAtY[0]))) << 4
		}
	}
	if _gbda.GBAtOverride[1] {
		_gceg &= 0xFBFF
		if _gbda.GBAtY[1] == 0 && _gbda.GBAtX[1] >= -int8(_aaad) {
			_gceg |= (_ddb >> uint(int8(_dgf)-_gbda.GBAtX[1]&0x1)) << 10
		} else {
			_gceg |= int(_gbda.getPixel(_ddc+int(_gbda.GBAtX[1]), _bgd+int(_gbda.GBAtY[1]))) << 10
		}
	}
	if _gbda.GBAtOverride[2] {
		_gceg &= 0xF7FF
		if _gbda.GBAtY[2] == 0 && _gbda.GBAtX[2] >= -int8(_aaad) {
			_gceg |= (_ddb >> uint(int8(_dgf)-_gbda.GBAtX[2]&0x1)) << 11
		} else {
			_gceg |= int(_gbda.getPixel(_ddc+int(_gbda.GBAtX[2]), _bgd+int(_gbda.GBAtY[2]))) << 11
		}
	}
	if _gbda.GBAtOverride[3] {
		_gceg &= 0x7FFF
		if _gbda.GBAtY[3] == 0 && _gbda.GBAtX[3] >= -int8(_aaad) {
			_gceg |= (_ddb >> uint(int8(_dgf)-_gbda.GBAtX[3]&0x1)) << 15
		} else {
			_gceg |= int(_gbda.getPixel(_ddc+int(_gbda.GBAtX[3]), _bgd+int(_gbda.GBAtY[3]))) << 15
		}
	}
	return _gceg
}
func (_ed *GenericRefinementRegion) GetRegionBitmap() (*_df.Bitmap, error) {
	var _da error
	_af.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _da != nil {
			_af.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0047\u0065\u0074R\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006dap\u0020\u0066a\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _da)
		} else {
			_af.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _ed.RegionBitmap != nil {
		return _ed.RegionBitmap, nil
	}
	_dc := 0
	if _ed.ReferenceBitmap == nil {
		_ed.ReferenceBitmap, _da = _ed.getGrReference()
		if _da != nil {
			return nil, _da
		}
	}
	if _ed._eg == nil {
		_ed._eg, _da = _b.New(_ed._fd)
		if _da != nil {
			return nil, _da
		}
	}
	if _ed._ab == nil {
		_ed._ab = _b.NewStats(8192, 1)
	}
	_ed.RegionBitmap = _df.New(int(_ed.RegionInfo.BitmapWidth), int(_ed.RegionInfo.BitmapHeight))
	if _ed.TemplateID == 0 {
		if _da = _ed.updateOverride(); _da != nil {
			return nil, _da
		}
	}
	_egd := (_ed.RegionBitmap.Width + 7) & -8
	var _bb int
	if _ed.IsTPGROn {
		_bb = int(-_ed.ReferenceDY) * _ed.ReferenceBitmap.RowStride
	}
	_fdf := _bb + 1
	for _bg := 0; _bg < _ed.RegionBitmap.Height; _bg++ {
		if _ed.IsTPGROn {
			_gdg, _eca := _ed.decodeSLTP()
			if _eca != nil {
				return nil, _eca
			}
			_dc ^= _gdg
		}
		if _dc == 0 {
			_da = _ed.decodeOptimized(_bg, _ed.RegionBitmap.Width, _ed.RegionBitmap.RowStride, _ed.ReferenceBitmap.RowStride, _egd, _bb, _fdf)
			if _da != nil {
				return nil, _da
			}
		} else {
			_da = _ed.decodeTypicalPredictedLine(_bg, _ed.RegionBitmap.Width, _ed.RegionBitmap.RowStride, _ed.ReferenceBitmap.RowStride, _egd, _bb)
			if _da != nil {
				return nil, _da
			}
		}
	}
	return _ed.RegionBitmap, nil
}
func (_dbad *GenericRegion) setParametersWithAt(_eacg bool, _bgge byte, _aeb, _fbbd bool, _bce, _gadfg []int8, _cbb, _gdf uint32, _dfed *_b.DecoderStats, _egbc *_b.Decoder) {
	_dbad.IsMMREncoded = _eacg
	_dbad.GBTemplate = _bgge
	_dbad.IsTPGDon = _aeb
	_dbad.GBAtX = _bce
	_dbad.GBAtY = _gadfg
	_dbad.RegionSegment.BitmapHeight = _gdf
	_dbad.RegionSegment.BitmapWidth = _cbb
	_dbad._cca = nil
	_dbad.Bitmap = nil
	if _dfed != nil {
		_dbad._bfe = _dfed
	}
	if _egbc != nil {
		_dbad._caf = _egbc
	}
	_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049O\u004e\u005d\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0053\u0044\u0041t\u003a\u0020\u0025\u0073", _dbad)
}
func (_agdg *TextRegion) parseHeader() error {
	var _acga error
	_af.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069n\u0073\u002e\u002e\u002e")
	defer func() {
		if _acga != nil {
			_af.Log.Trace("\u005b\u0054\u0045\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044E\u0052\u005d\u0020\u0066\u0061i\u006c\u0065d\u002e\u0020\u0025\u0076", _acga)
		} else {
			_af.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069s\u0068\u0065\u0064\u002e")
		}
	}()
	if _acga = _agdg.RegionInfo.parseHeader(); _acga != nil {
		return _acga
	}
	if _acga = _agdg.readRegionFlags(); _acga != nil {
		return _acga
	}
	if _agdg.IsHuffmanEncoded {
		if _acga = _agdg.readHuffmanFlags(); _acga != nil {
			return _acga
		}
	}
	if _acga = _agdg.readUseRefinement(); _acga != nil {
		return _acga
	}
	if _acga = _agdg.readAmountOfSymbolInstances(); _acga != nil {
		return _acga
	}
	if _acga = _agdg.getSymbols(); _acga != nil {
		return _acga
	}
	if _acga = _agdg.computeSymbolCodeLength(); _acga != nil {
		return _acga
	}
	if _acga = _agdg.checkInput(); _acga != nil {
		return _acga
	}
	_af.Log.Trace("\u0025\u0073", _agdg.String())
	return nil
}
func (_eec *GenericRegion) Size() int { return _eec.RegionSegment.Size() + 1 + 2*len(_eec.GBAtX) }

type TextRegion struct {
	_deba                   *_ae.Reader
	RegionInfo              *RegionSegment
	SbrTemplate             int8
	SbDsOffset              int8
	DefaultPixel            int8
	CombinationOperator     _df.CombinationOperator
	IsTransposed            int8
	ReferenceCorner         int16
	LogSBStrips             int16
	UseRefinement           bool
	IsHuffmanEncoded        bool
	SbHuffRSize             int8
	SbHuffRDY               int8
	SbHuffRDX               int8
	SbHuffRDHeight          int8
	SbHuffRDWidth           int8
	SbHuffDT                int8
	SbHuffDS                int8
	SbHuffFS                int8
	SbrATX                  []int8
	SbrATY                  []int8
	NumberOfSymbolInstances uint32
	_bcef                   int64
	SbStrips                int8
	NumberOfSymbols         uint32
	RegionBitmap            *_df.Bitmap
	Symbols                 []*_df.Bitmap
	_fagf                   *_b.Decoder
	_gefd                   *GenericRefinementRegion
	_bgce                   *_b.DecoderStats
	_deef                   *_b.DecoderStats
	_ccac                   *_b.DecoderStats
	_cbgce                  *_b.DecoderStats
	_gddf                   *_b.DecoderStats
	_abbd                   *_b.DecoderStats
	_gbgf                   *_b.DecoderStats
	_bebb                   *_b.DecoderStats
	_gdgaf                  *_b.DecoderStats
	_fddb                   *_b.DecoderStats
	_aaag                   *_b.DecoderStats
	_adfd                   int8
	_bebbb                  *_ge.FixedSizeTable
	Header                  *Header
	_cbgg                   _ge.Tabler
	_ddea                   _ge.Tabler
	_cafd                   _ge.Tabler
	_agab                   _ge.Tabler
	_cagbc                  _ge.Tabler
	_bgde                   _ge.Tabler
	_ffaa                   _ge.Tabler
	_ffcfg                  _ge.Tabler
	_eeaa, _gecbc           map[int]int
	_faba                   []int
	_eegg                   *_df.Points
	_gdffg                  *_df.Bitmaps
	_cdbe                   *_fgc.IntSlice
	_abcaa, _dccc           int
	_fgbef                  *_df.Boxes
}

func (_aeaa *PageInformationSegment) readMaxStripeSize() error {
	_cfaff, _gfdfb := _aeaa._fgee.ReadBits(15)
	if _gfdfb != nil {
		return _gfdfb
	}
	_aeaa.MaxStripeSize = uint16(_cfaff & _c.MaxUint16)
	return nil
}
func (_gfd *GenericRefinementRegion) decodeTypicalPredictedLineTemplate0(_egg, _gba, _beb, _gce, _bac, _ged, _fcac, _ea, _cee int) error {
	var (
		_fdd, _aab, _gga, _fag, _ace, _bggb int
		_cb                                 byte
		_gab                                error
	)
	if _egg > 0 {
		_cb, _gab = _gfd.RegionBitmap.GetByte(_fcac - _beb)
		if _gab != nil {
			return _gab
		}
		_gga = int(_cb)
	}
	if _ea > 0 && _ea <= _gfd.ReferenceBitmap.Height {
		_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee - _gce + _ged)
		if _gab != nil {
			return _gab
		}
		_fag = int(_cb) << 4
	}
	if _ea >= 0 && _ea < _gfd.ReferenceBitmap.Height {
		_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee + _ged)
		if _gab != nil {
			return _gab
		}
		_ace = int(_cb) << 1
	}
	if _ea > -2 && _ea < _gfd.ReferenceBitmap.Height-1 {
		_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee + _gce + _ged)
		if _gab != nil {
			return _gab
		}
		_bggb = int(_cb)
	}
	_fdd = ((_gga >> 5) & 0x6) | ((_bggb >> 2) & 0x30) | (_ace & 0x180) | (_fag & 0xc00)
	var _eb int
	for _dfg := 0; _dfg < _bac; _dfg = _eb {
		var _gaec int
		_eb = _dfg + 8
		var _dfc int
		if _dfc = _gba - _dfg; _dfc > 8 {
			_dfc = 8
		}
		_cd := _eb < _gba
		_gad := _eb < _gfd.ReferenceBitmap.Width
		_gcea := _ged + 1
		if _egg > 0 {
			_cb = 0
			if _cd {
				_cb, _gab = _gfd.RegionBitmap.GetByte(_fcac - _beb + 1)
				if _gab != nil {
					return _gab
				}
			}
			_gga = (_gga << 8) | int(_cb)
		}
		if _ea > 0 && _ea <= _gfd.ReferenceBitmap.Height {
			var _eac int
			if _gad {
				_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee - _gce + _gcea)
				if _gab != nil {
					return _gab
				}
				_eac = int(_cb) << 4
			}
			_fag = (_fag << 8) | _eac
		}
		if _ea >= 0 && _ea < _gfd.ReferenceBitmap.Height {
			var _abf int
			if _gad {
				_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee + _gcea)
				if _gab != nil {
					return _gab
				}
				_abf = int(_cb) << 1
			}
			_ace = (_ace << 8) | _abf
		}
		if _ea > -2 && _ea < (_gfd.ReferenceBitmap.Height-1) {
			_cb = 0
			if _gad {
				_cb, _gab = _gfd.ReferenceBitmap.GetByte(_cee + _gce + _gcea)
				if _gab != nil {
					return _gab
				}
			}
			_bggb = (_bggb << 8) | int(_cb)
		}
		for _cace := 0; _cace < _dfc; _cace++ {
			var _gfa int
			_cef := false
			_gcd := (_fdd >> 4) & 0x1ff
			if _gcd == 0x1ff {
				_cef = true
				_gfa = 1
			} else if _gcd == 0x00 {
				_cef = true
			}
			if !_cef {
				if _gfd._bf {
					_aab = _gfd.overrideAtTemplate0(_fdd, _dfg+_cace, _egg, _gaec, _cace)
					_gfd._ab.SetIndex(int32(_aab))
				} else {
					_gfd._ab.SetIndex(int32(_fdd))
				}
				_gfa, _gab = _gfd._eg.DecodeBit(_gfd._ab)
				if _gab != nil {
					return _gab
				}
			}
			_fed := uint(7 - _cace)
			_gaec |= _gfa << _fed
			_fdd = ((_fdd & 0xdb6) << 1) | _gfa | (_gga>>_fed+5)&0x002 | ((_bggb>>_fed + 2) & 0x010) | ((_ace >> _fed) & 0x080) | ((_fag >> _fed) & 0x400)
		}
		_gab = _gfd.RegionBitmap.SetByte(_fcac, byte(_gaec))
		if _gab != nil {
			return _gab
		}
		_fcac++
		_cee++
	}
	return nil
}
func (_gegb *PatternDictionary) extractPatterns(_bgaf *_df.Bitmap) error {
	var _aacf int
	_ecbg := make([]*_df.Bitmap, _gegb.GrayMax+1)
	for _aacf <= int(_gegb.GrayMax) {
		_feag := int(_gegb.HdpWidth) * _aacf
		_acce := _f.Rect(_feag, 0, _feag+int(_gegb.HdpWidth), int(_gegb.HdpHeight))
		_bfbbb, _eadb := _df.Extract(_acce, _bgaf)
		if _eadb != nil {
			return _eadb
		}
		_ecbg[_aacf] = _bfbbb
		_aacf++
	}
	_gegb.Patterns = _ecbg
	return nil
}
func (_gefa *GenericRegion) overrideAtTemplate3(_agfg, _abed, _fbb, _dfdf, _bgfe int) int {
	_agfg &= 0x3EF
	if _gefa.GBAtY[0] == 0 && _gefa.GBAtX[0] >= -int8(_bgfe) {
		_agfg |= (_dfdf >> uint(7-(int8(_bgfe)+_gefa.GBAtX[0])) & 0x1) << 4
	} else {
		_agfg |= int(_gefa.getPixel(_abed+int(_gefa.GBAtX[0]), _fbb+int(_gefa.GBAtY[0]))) << 4
	}
	return _agfg
}
func (_agac *SymbolDictionary) GetDictionary() ([]*_df.Bitmap, error) {
	_af.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c-\u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		_af.Log.Trace("\u005b\u0053\u0059M\u0042\u004f\u004c\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		_af.Log.Trace("\u005b\u0053Y\u004d\u0042\u004f\u004c\u002dD\u0049\u0043\u0054\u0049\u004fN\u0041\u0052\u0059\u005d\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u000a\u0045\u0078\u003a\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u000a\u006e\u0065\u0077\u003a\u0027\u0025\u0073\u0027", _agac._ddac, _agac._dede)
	}()
	if _agac._ddac == nil {
		var _afae error
		if _agac.UseRefinementAggregation {
			_agac._efgc = _agac.getSbSymCodeLen()
		}
		if !_agac.IsHuffmanEncoded {
			if _afae = _agac.setCodingStatistics(); _afae != nil {
				return nil, _afae
			}
		}
		_agac._dede = make([]*_df.Bitmap, _agac.NumberOfNewSymbols)
		var _cegc []int
		if _agac.IsHuffmanEncoded && !_agac.UseRefinementAggregation {
			_cegc = make([]int, _agac.NumberOfNewSymbols)
		}
		if _afae = _agac.setSymbolsArray(); _afae != nil {
			return nil, _afae
		}
		var _eefg, _cbbd int64
		_agac._dacc = 0
		for _agac._dacc < _agac.NumberOfNewSymbols {
			_cbbd, _afae = _agac.decodeHeightClassDeltaHeight()
			if _afae != nil {
				return nil, _afae
			}
			_eefg += _cbbd
			var _cfda, _fadg uint32
			_dgfe := int64(_agac._dacc)
			for {
				var _cabb int64
				_cabb, _afae = _agac.decodeDifferenceWidth()
				if _aa.Is(_afae, _e.ErrOOB) {
					break
				}
				if _afae != nil {
					return nil, _afae
				}
				if _agac._dacc >= _agac.NumberOfNewSymbols {
					break
				}
				_cfda += uint32(_cabb)
				_fadg += _cfda
				if !_agac.IsHuffmanEncoded || _agac.UseRefinementAggregation {
					if !_agac.UseRefinementAggregation {
						_afae = _agac.decodeDirectlyThroughGenericRegion(_cfda, uint32(_eefg))
						if _afae != nil {
							return nil, _afae
						}
					} else {
						_afae = _agac.decodeAggregate(_cfda, uint32(_eefg))
						if _afae != nil {
							return nil, _afae
						}
					}
				} else if _agac.IsHuffmanEncoded && !_agac.UseRefinementAggregation {
					_cegc[_agac._dacc] = int(_cfda)
				}
				_agac._dacc++
			}
			if _agac.IsHuffmanEncoded && !_agac.UseRefinementAggregation {
				var _abeg int64
				if _agac.SdHuffBMSizeSelection == 0 {
					var _gdga _ge.Tabler
					_gdga, _afae = _ge.GetStandardTable(1)
					if _afae != nil {
						return nil, _afae
					}
					_abeg, _afae = _gdga.Decode(_agac._bbda)
					if _afae != nil {
						return nil, _afae
					}
				} else {
					_abeg, _afae = _agac.huffDecodeBmSize()
					if _afae != nil {
						return nil, _afae
					}
				}
				_agac._bbda.Align()
				var _bgb *_df.Bitmap
				_bgb, _afae = _agac.decodeHeightClassCollectiveBitmap(_abeg, uint32(_eefg), _fadg)
				if _afae != nil {
					return nil, _afae
				}
				_afae = _agac.decodeHeightClassBitmap(_bgb, _dgfe, int(_eefg), _cegc)
				if _afae != nil {
					return nil, _afae
				}
			}
		}
		_daae, _afae := _agac.getToExportFlags()
		if _afae != nil {
			return nil, _afae
		}
		_agac.setExportedSymbols(_daae)
	}
	return _agac._ddac, nil
}
func (_fac *HalftoneRegion) renderPattern(_fcff [][]int) (_bbdb error) {
	var _ggc, _ddf int
	for _gedbf := 0; _gedbf < int(_fac.HGridHeight); _gedbf++ {
		for _cdga := 0; _cdga < int(_fac.HGridWidth); _cdga++ {
			_ggc = _fac.computeX(_gedbf, _cdga)
			_ddf = _fac.computeY(_gedbf, _cdga)
			_gdeg := _fac.Patterns[_fcff[_gedbf][_cdga]]
			if _bbdb = _df.Blit(_gdeg, _fac.HalftoneRegionBitmap, _ggc+int(_fac.HGridX), _ddf+int(_fac.HGridY), _fac.CombinationOperator); _bbdb != nil {
				return _bbdb
			}
		}
	}
	return nil
}
func (_agfd *SymbolDictionary) getToExportFlags() ([]int, error) {
	var (
		_cfbd  int
		_cfec  int32
		_cfge  error
		_afge  = int32(_agfd._eaab + _agfd.NumberOfNewSymbols)
		_dgegb = make([]int, _afge)
	)
	for _eeae := int32(0); _eeae < _afge; _eeae += _cfec {
		if _agfd.IsHuffmanEncoded {
			_gcgf, _ccfc := _ge.GetStandardTable(1)
			if _ccfc != nil {
				return nil, _ccfc
			}
			_gbg, _ccfc := _gcgf.Decode(_agfd._bbda)
			if _ccfc != nil {
				return nil, _ccfc
			}
			_cfec = int32(_gbg)
		} else {
			_cfec, _cfge = _agfd._fdbf.DecodeInt(_agfd._bgdf)
			if _cfge != nil {
				return nil, _cfge
			}
		}
		if _cfec != 0 {
			if _eeae+_cfec > _afge {
				return nil, _fe.Error("\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0067\u0065\u0074T\u006f\u0045\u0078\u0070\u006f\u0072\u0074F\u006c\u0061\u0067\u0073", "\u006d\u0061\u006c\u0066\u006f\u0072m\u0065\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0064\u0061\u0074\u0061\u0020\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u002e\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u006f\u0075\u0074\u0020\u006ff\u0020r\u0061\u006e\u0067\u0065")
			}
			for _ggaf := _eeae; _ggaf < _eeae+_cfec; _ggaf++ {
				_dgegb[_ggaf] = _cfbd
			}
		}
		if _cfbd == 0 {
			_cfbd = 1
		} else {
			_cfbd = 0
		}
	}
	return _dgegb, nil
}
func (_gbae *GenericRegion) GetRegionBitmap() (_caba *_df.Bitmap, _ebc error) {
	if _gbae.Bitmap != nil {
		return _gbae.Bitmap, nil
	}
	if _gbae.IsMMREncoded {
		if _gbae._cca == nil {
			_gbae._cca, _ebc = _aad.New(_gbae._abe, int(_gbae.RegionSegment.BitmapWidth), int(_gbae.RegionSegment.BitmapHeight), _gbae.DataOffset, _gbae.DataLength)
			if _ebc != nil {
				return nil, _ebc
			}
		}
		_gbae.Bitmap, _ebc = _gbae._cca.UncompressMMR()
		return _gbae.Bitmap, _ebc
	}
	if _ebc = _gbae.updateOverrideFlags(); _ebc != nil {
		return nil, _ebc
	}
	var _fcaa int
	if _gbae._caf == nil {
		_gbae._caf, _ebc = _b.New(_gbae._abe)
		if _ebc != nil {
			return nil, _ebc
		}
	}
	if _gbae._bfe == nil {
		_gbae._bfe = _b.NewStats(65536, 1)
	}
	_gbae.Bitmap = _df.New(int(_gbae.RegionSegment.BitmapWidth), int(_gbae.RegionSegment.BitmapHeight))
	_agf := int(uint32(_gbae.Bitmap.Width+7) & (^uint32(7)))
	for _cfa := 0; _cfa < _gbae.Bitmap.Height; _cfa++ {
		if _gbae.IsTPGDon {
			var _ebf int
			_ebf, _ebc = _gbae.decodeSLTP()
			if _ebc != nil {
				return nil, _ebc
			}
			_fcaa ^= _ebf
		}
		if _fcaa == 1 {
			if _cfa > 0 {
				if _ebc = _gbae.copyLineAbove(_cfa); _ebc != nil {
					return nil, _ebc
				}
			}
		} else {
			if _ebc = _gbae.decodeLine(_cfa, _gbae.Bitmap.Width, _agf); _ebc != nil {
				return nil, _ebc
			}
		}
	}
	return _gbae.Bitmap, nil
}
func (_cagd *PatternDictionary) readIsMMREncoded() error {
	_cdac, _gada := _cagd._fdcf.ReadBit()
	if _gada != nil {
		return _gada
	}
	if _cdac != 0 {
		_cagd.IsMMREncoded = true
	}
	return nil
}

type Regioner interface {
	GetRegionBitmap() (*_df.Bitmap, error)
	GetRegionInfo() *RegionSegment
}
type EndOfStripe struct {
	_bc   *_ae.Reader
	_fgce int
}

func (_aeae *TextRegion) decodeCurrentT() (int64, error) {
	if _aeae.SbStrips != 1 {
		if _aeae.IsHuffmanEncoded {
			_dbef, _afce := _aeae._deba.ReadBits(byte(_aeae.LogSBStrips))
			return int64(_dbef), _afce
		}
		_egabe, _ggdc := _aeae._fagf.DecodeInt(_aeae._cbgce)
		if _ggdc != nil {
			return 0, _ggdc
		}
		return int64(_egabe), nil
	}
	return 0, nil
}
func (_afag *SymbolDictionary) encodeATFlags(_cede _ae.BinaryWriter) (_cbgc int, _fafe error) {
	const _eded = "\u0065\u006e\u0063\u006f\u0064\u0065\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if _afag.IsHuffmanEncoded || _afag.SdTemplate != 0 {
		return 0, nil
	}
	_bebee := 4
	if _afag.SdTemplate != 0 {
		_bebee = 1
	}
	for _gaef := 0; _gaef < _bebee; _gaef++ {
		if _fafe = _cede.WriteByte(byte(_afag.SdATX[_gaef])); _fafe != nil {
			return _cbgc, _fe.Wrapf(_fafe, _eded, "\u0053d\u0041\u0054\u0058\u005b\u0025\u0064]", _gaef)
		}
		_cbgc++
		if _fafe = _cede.WriteByte(byte(_afag.SdATY[_gaef])); _fafe != nil {
			return _cbgc, _fe.Wrapf(_fafe, _eded, "\u0053d\u0041\u0054\u0059\u005b\u0025\u0064]", _gaef)
		}
		_cbgc++
	}
	return _cbgc, nil
}
func (_degb *PageInformationSegment) CombinationOperatorOverrideAllowed() bool { return _degb._gced }
func (_acdc *PageInformationSegment) String() string {
	_cecb := &_a.Builder{}
	_cecb.WriteString("\u000a\u005b\u0050\u0041G\u0045\u002d\u0049\u004e\u0046\u004f\u0052\u004d\u0041\u0054I\u004fN\u002d\u0053\u0045\u0047\u004d\u0045\u004eT\u005d\u000a")
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d \u0042\u004d\u0048e\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u000a", _acdc.PageBMHeight))
	_cecb.WriteString(_dg.Sprintf("\u0009-\u0020B\u004d\u0057\u0069\u0064\u0074\u0068\u003a\u0020\u0025\u0064\u000a", _acdc.PageBMWidth))
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0058\u003a\u0020\u0025\u0064\u000a", _acdc.ResolutionX))
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0059\u003a\u0020\u0025\u0064\u000a", _acdc.ResolutionY))
	_cecb.WriteString(_dg.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0073\u000a", _acdc._bca))
	_cecb.WriteString(_dg.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006eO\u0070\u0065\u0072\u0061\u0074\u006f\u0072O\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025v\u000a", _acdc._gced))
	_cecb.WriteString(_dg.Sprintf("\u0009-\u0020I\u0073\u004c\u006f\u0073\u0073l\u0065\u0073s\u003a\u0020\u0025\u0076\u000a", _acdc.IsLossless))
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d\u0020R\u0065\u0071\u0075\u0069r\u0065\u0073\u0041\u0075\u0078\u0069\u006ci\u0061\u0072\u0079\u0042\u0075\u0066\u0066\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _acdc._fgaa))
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d\u0020M\u0069\u0067\u0068\u0074C\u006f\u006e\u0074\u0061\u0069\u006e\u0052e\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0076\u000a", _acdc._cadb))
	_cecb.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0049\u0073\u0053\u0074\u0072\u0069\u0070\u0065\u0064:\u0020\u0025\u0076\u000a", _acdc.IsStripe))
	_cecb.WriteString(_dg.Sprintf("\t\u002d\u0020\u004d\u0061xS\u0074r\u0069\u0070\u0065\u0053\u0069z\u0065\u003a\u0020\u0025\u0076\u000a", _acdc.MaxStripeSize))
	return _cecb.String()
}
func (_fdge *TableSegment) parseHeader() error {
	var (
		_fffg int
		_cdb  uint64
		_cgf  error
	)
	_fffg, _cgf = _fdge._ebef.ReadBit()
	if _cgf != nil {
		return _cgf
	}
	if _fffg == 1 {
		return _dg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0061\u0062\u006c\u0065 \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0064e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u002e\u0020\u0042\u002e\u0032\u002e1\u0020\u0043\u006f\u0064\u0065\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0042\u0069\u0074\u0020\u0037\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u007a\u0065\u0072\u006f\u002e\u0020\u0057a\u0073\u003a \u0025\u0064", _fffg)
	}
	if _cdb, _cgf = _fdge._ebef.ReadBits(3); _cgf != nil {
		return _cgf
	}
	_fdge._bfcc = (int32(_cdb) + 1) & 0xf
	if _cdb, _cgf = _fdge._ebef.ReadBits(3); _cgf != nil {
		return _cgf
	}
	_fdge._dfdfb = (int32(_cdb) + 1) & 0xf
	if _cdb, _cgf = _fdge._ebef.ReadBits(32); _cgf != nil {
		return _cgf
	}
	_fdge._gefag = int32(_cdb & _c.MaxInt32)
	if _cdb, _cgf = _fdge._ebef.ReadBits(32); _cgf != nil {
		return _cgf
	}
	_fdge._gafe = int32(_cdb & _c.MaxInt32)
	return nil
}

type EncodeInitializer interface{ InitEncode() }

func (_dcea *GenericRegion) Encode(w _ae.BinaryWriter) (_cafc int, _eeaf error) {
	const _fbfg = "G\u0065n\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069o\u006e\u002e\u0045\u006eco\u0064\u0065"
	if _dcea.Bitmap == nil {
		return 0, _fe.Error(_fbfg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	_efe, _eeaf := _dcea.RegionSegment.Encode(w)
	if _eeaf != nil {
		return 0, _fe.Wrap(_eeaf, _fbfg, "\u0052\u0065\u0067\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_cafc += _efe
	if _eeaf = w.SkipBits(4); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "\u0073k\u0069p\u0020\u0072\u0065\u0073\u0065r\u0076\u0065d\u0020\u0062\u0069\u0074\u0073")
	}
	var _aed int
	if _dcea.IsTPGDon {
		_aed = 1
	}
	if _eeaf = w.WriteBit(_aed); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "\u0074\u0070\u0067\u0064\u006f\u006e")
	}
	_aed = 0
	if _eeaf = w.WriteBit(int(_dcea.GBTemplate>>1) & 0x01); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "f\u0069r\u0073\u0074\u0020\u0067\u0062\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 b\u0069\u0074")
	}
	if _eeaf = w.WriteBit(int(_dcea.GBTemplate) & 0x01); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "s\u0065\u0063\u006f\u006ed \u0067b\u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u0062\u0069\u0074")
	}
	if _dcea.UseMMR {
		_aed = 1
	}
	if _eeaf = w.WriteBit(_aed); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "u\u0073\u0065\u0020\u004d\u004d\u0052\u0020\u0062\u0069\u0074")
	}
	_cafc++
	if _efe, _eeaf = _dcea.writeGBAtPixels(w); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "")
	}
	_cafc += _efe
	_dbg := _gc.New()
	if _eeaf = _dbg.EncodeBitmap(_dcea.Bitmap, _dcea.IsTPGDon); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "")
	}
	_dbg.Final()
	var _cce int64
	if _cce, _eeaf = _dbg.WriteTo(w); _eeaf != nil {
		return _cafc, _fe.Wrap(_eeaf, _fbfg, "")
	}
	_cafc += int(_cce)
	return _cafc, nil
}
func (_bbed *SymbolDictionary) decodeHeightClassBitmap(_feab *_df.Bitmap, _efcg int64, _bdadd int, _eage []int) error {
	for _gdfg := _efcg; _gdfg < int64(_bbed._dacc); _gdfg++ {
		var _ffe int
		for _gbea := _efcg; _gbea <= _gdfg-1; _gbea++ {
			_ffe += _eage[_gbea]
		}
		_ddfg := _f.Rect(_ffe, 0, _ffe+_eage[_gdfg], _bdadd)
		_bddd, _abg := _df.Extract(_ddfg, _feab)
		if _abg != nil {
			return _abg
		}
		_bbed._dede[_gdfg] = _bddd
		_bbed._bccf = append(_bbed._bccf, _bddd)
	}
	return nil
}
func (_bfdb *GenericRegion) overrideAtTemplate2(_acg, _agga, _cadg, _acfa, _efc int) int {
	_acg &= 0x3FB
	if _bfdb.GBAtY[0] == 0 && _bfdb.GBAtX[0] >= -int8(_efc) {
		_acg |= (_acfa >> uint(7-(int8(_efc)+_bfdb.GBAtX[0])) & 0x1) << 2
	} else {
		_acg |= int(_bfdb.getPixel(_agga+int(_bfdb.GBAtX[0]), _cadg+int(_bfdb.GBAtY[0]))) << 2
	}
	return _acg
}

var _ templater = &template1{}

func (_baff *PageInformationSegment) Encode(w _ae.BinaryWriter) (_cacd int, _adde error) {
	const _effe = "\u0050\u0061g\u0065\u0049\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u002e\u0045\u006eco\u0064\u0065"
	_ccce := make([]byte, 4)
	_d.BigEndian.PutUint32(_ccce, uint32(_baff.PageBMWidth))
	_cacd, _adde = w.Write(_ccce)
	if _adde != nil {
		return _cacd, _fe.Wrap(_adde, _effe, "\u0077\u0069\u0064t\u0068")
	}
	_d.BigEndian.PutUint32(_ccce, uint32(_baff.PageBMHeight))
	var _agbe int
	_agbe, _adde = w.Write(_ccce)
	if _adde != nil {
		return _agbe + _cacd, _fe.Wrap(_adde, _effe, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_cacd += _agbe
	_d.BigEndian.PutUint32(_ccce, uint32(_baff.ResolutionX))
	_agbe, _adde = w.Write(_ccce)
	if _adde != nil {
		return _agbe + _cacd, _fe.Wrap(_adde, _effe, "\u0078\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_cacd += _agbe
	_d.BigEndian.PutUint32(_ccce, uint32(_baff.ResolutionY))
	if _agbe, _adde = w.Write(_ccce); _adde != nil {
		return _agbe + _cacd, _fe.Wrap(_adde, _effe, "\u0079\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_cacd += _agbe
	if _adde = _baff.encodeFlags(w); _adde != nil {
		return _cacd, _fe.Wrap(_adde, _effe, "")
	}
	_cacd++
	if _agbe, _adde = _baff.encodeStripingInformation(w); _adde != nil {
		return _cacd, _fe.Wrap(_adde, _effe, "")
	}
	_cacd += _agbe
	return _cacd, nil
}

var _ SegmentEncoder = &GenericRegion{}

func (_gecd *HalftoneRegion) GetRegionBitmap() (*_df.Bitmap, error) {
	if _gecd.HalftoneRegionBitmap != nil {
		return _gecd.HalftoneRegionBitmap, nil
	}
	var _deggb error
	_gecd.HalftoneRegionBitmap = _df.New(int(_gecd.RegionSegment.BitmapWidth), int(_gecd.RegionSegment.BitmapHeight))
	if _gecd.Patterns == nil || (_gecd.Patterns != nil && len(_gecd.Patterns) == 0) {
		_gecd.Patterns, _deggb = _gecd.GetPatterns()
		if _deggb != nil {
			return nil, _deggb
		}
	}
	if _gecd.HDefaultPixel == 1 {
		_gecd.HalftoneRegionBitmap.SetDefaultPixel()
	}
	_ebgd := _c.Ceil(_c.Log(float64(len(_gecd.Patterns))) / _c.Log(2))
	_ccf := int(_ebgd)
	var _cdag [][]int
	_cdag, _deggb = _gecd.grayScaleDecoding(_ccf)
	if _deggb != nil {
		return nil, _deggb
	}
	if _deggb = _gecd.renderPattern(_cdag); _deggb != nil {
		return nil, _deggb
	}
	return _gecd.HalftoneRegionBitmap, nil
}
func (_egab *Header) String() string {
	_gedf := &_a.Builder{}
	_gedf.WriteString("\u000a[\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045A\u0044\u0045\u0052\u005d\u000a")
	_gedf.WriteString(_dg.Sprintf("\t\u002d\u0020\u0053\u0065gm\u0065n\u0074\u004e\u0075\u006d\u0062e\u0072\u003a\u0020\u0025\u0076\u000a", _egab.SegmentNumber))
	_gedf.WriteString(_dg.Sprintf("\u0009\u002d\u0020T\u0079\u0070\u0065\u003a\u0020\u0025\u0076\u000a", _egab.Type))
	_gedf.WriteString(_dg.Sprintf("\u0009-\u0020R\u0065\u0074\u0061\u0069\u006eF\u006c\u0061g\u003a\u0020\u0025\u0076\u000a", _egab.RetainFlag))
	_gedf.WriteString(_dg.Sprintf("\u0009\u002d\u0020Pa\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _egab.PageAssociation))
	_gedf.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0050\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074i\u006fn\u0046\u0069\u0065\u006c\u0064\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _egab.PageAssociationFieldSize))
	_gedf.WriteString("\u0009-\u0020R\u0054\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u0053\u003a\u000a")
	for _, _bdg := range _egab.RTSNumbers {
		_gedf.WriteString(_dg.Sprintf("\u0009\t\u002d\u0020\u0025\u0064\u000a", _bdg))
	}
	_gedf.WriteString(_dg.Sprintf("\t\u002d \u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0065n\u0067\u0074\u0068\u003a %\u0076\u000a", _egab.HeaderLength))
	_gedf.WriteString(_dg.Sprintf("\u0009-\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061L\u0065n\u0067\u0074\u0068\u003a\u0020\u0025\u0076\n", _egab.SegmentDataLength))
	_gedf.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074D\u0061\u0074\u0061\u0053\u0074\u0061\u0072t\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0076\u000a", _egab.SegmentDataStartOffset))
	return _gedf.String()
}
func (_ffda *SymbolDictionary) getSbSymCodeLen() int8 {
	_edcf := int8(_c.Ceil(_c.Log(float64(_ffda._eaab+_ffda.NumberOfNewSymbols)) / _c.Log(2)))
	if _ffda.IsHuffmanEncoded && _edcf < 1 {
		return 1
	}
	return _edcf
}
func (_dcdd *TextRegion) decodeRdh() (int64, error) {
	const _dcedf = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064h"
	if _dcdd.IsHuffmanEncoded {
		if _dcdd.SbHuffRDHeight == 3 {
			if _dcdd._cagbc == nil {
				var (
					_geab int
					_eafe error
				)
				if _dcdd.SbHuffFS == 3 {
					_geab++
				}
				if _dcdd.SbHuffDS == 3 {
					_geab++
				}
				if _dcdd.SbHuffDT == 3 {
					_geab++
				}
				if _dcdd.SbHuffRDWidth == 3 {
					_geab++
				}
				_dcdd._cagbc, _eafe = _dcdd.getUserTable(_geab)
				if _eafe != nil {
					return 0, _fe.Wrap(_eafe, _dcedf, "")
				}
			}
			return _dcdd._cagbc.Decode(_dcdd._deba)
		}
		_bded, _daeb := _ge.GetStandardTable(14 + int(_dcdd.SbHuffRDHeight))
		if _daeb != nil {
			return 0, _fe.Wrap(_daeb, _dcedf, "")
		}
		return _bded.Decode(_dcdd._deba)
	}
	_bdbcf, _aggcg := _dcdd._fagf.DecodeInt(_dcdd._gbgf)
	if _aggcg != nil {
		return 0, _fe.Wrap(_aggcg, _dcedf, "")
	}
	return int64(_bdbcf), nil
}
func (_edad *SymbolDictionary) setSymbolsArray() error {
	if _edad._fege == nil {
		if _gffe := _edad.retrieveImportSymbols(); _gffe != nil {
			return _gffe
		}
	}
	if _edad._bccf == nil {
		_edad._bccf = append(_edad._bccf, _edad._fege...)
	}
	return nil
}
func NewRegionSegment(r *_ae.Reader) *RegionSegment { return &RegionSegment{_gega: r} }
func (_ebbc *HalftoneRegion) GetPatterns() ([]*_df.Bitmap, error) {
	var (
		_add []*_df.Bitmap
		_ecb error
	)
	for _, _fgf := range _ebbc._bfed.RTSegments {
		var _bdba Segmenter
		_bdba, _ecb = _fgf.GetSegmentData()
		if _ecb != nil {
			_af.Log.Debug("\u0047e\u0074\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061 \u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025v", _ecb)
			return nil, _ecb
		}
		_bdf, _bfbb := _bdba.(*PatternDictionary)
		if !_bfbb {
			_ecb = _dg.Errorf("\u0072e\u006c\u0061t\u0065\u0064\u0020\u0073e\u0067\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0070at\u0074\u0065\u0072n\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u003a \u0025\u0054", _bdba)
			return nil, _ecb
		}
		var _bcge []*_df.Bitmap
		_bcge, _ecb = _bdf.GetDictionary()
		if _ecb != nil {
			_af.Log.Debug("\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ecb)
			return nil, _ecb
		}
		_add = append(_add, _bcge...)
	}
	return _add, nil
}

const (
	TSymbolDictionary                         Type = 0
	TIntermediateTextRegion                   Type = 4
	TImmediateTextRegion                      Type = 6
	TImmediateLosslessTextRegion              Type = 7
	TPatternDictionary                        Type = 16
	TIntermediateHalftoneRegion               Type = 20
	TImmediateHalftoneRegion                  Type = 22
	TImmediateLosslessHalftoneRegion          Type = 23
	TIntermediateGenericRegion                Type = 36
	TImmediateGenericRegion                   Type = 38
	TImmediateLosslessGenericRegion           Type = 39
	TIntermediateGenericRefinementRegion      Type = 40
	TImmediateGenericRefinementRegion         Type = 42
	TImmediateLosslessGenericRefinementRegion Type = 43
	TPageInformation                          Type = 48
	TEndOfPage                                Type = 49
	TEndOfStrip                               Type = 50
	TEndOfFile                                Type = 51
	TProfiles                                 Type = 52
	TTables                                   Type = 53
	TExtension                                Type = 62
	TBitmap                                   Type = 70
)

func (_dcgb *PageInformationSegment) encodeStripingInformation(_feaa _ae.BinaryWriter) (_gfdd int, _cagb error) {
	const _dbe = "\u0065n\u0063\u006f\u0064\u0065S\u0074\u0072\u0069\u0070\u0069n\u0067I\u006ef\u006f\u0072\u006d\u0061\u0074\u0069\u006fn"
	if !_dcgb.IsStripe {
		if _gfdd, _cagb = _feaa.Write([]byte{0x00, 0x00}); _cagb != nil {
			return 0, _fe.Wrap(_cagb, _dbe, "n\u006f\u0020\u0073\u0074\u0072\u0069\u0070\u0069\u006e\u0067")
		}
		return _gfdd, nil
	}
	_eddg := make([]byte, 2)
	_d.BigEndian.PutUint16(_eddg, _dcgb.MaxStripeSize|1<<15)
	if _gfdd, _cagb = _feaa.Write(_eddg); _cagb != nil {
		return 0, _fe.Wrapf(_cagb, _dbe, "\u0073\u0074\u0072i\u0070\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _dcgb.MaxStripeSize)
	}
	return _gfdd, nil
}
func (_fdegd *GenericRegion) setOverrideFlag(_fada int) {
	_fdegd.GBAtOverride[_fada] = true
	_fdegd._dca = true
}
func (_bdbc *Header) referenceSize() uint {
	switch {
	case _bdbc.SegmentNumber <= 255:
		return 1
	case _bdbc.SegmentNumber <= 65535:
		return 2
	default:
		return 4
	}
}
func (_gdegf *Header) writeSegmentPageAssociation(_bfedg _ae.BinaryWriter) (_egbf int, _aaabc error) {
	const _bcgb = "w\u0072\u0069\u0074\u0065\u0053\u0065g\u006d\u0065\u006e\u0074\u0050\u0061\u0067\u0065\u0041s\u0073\u006f\u0063i\u0061t\u0069\u006f\u006e"
	if _gdegf.pageSize() != 4 {
		if _aaabc = _bfedg.WriteByte(byte(_gdegf.PageAssociation)); _aaabc != nil {
			return 0, _fe.Wrap(_aaabc, _bcgb, "\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065\u0020\u0021\u003d\u0020\u0034")
		}
		return 1, nil
	}
	_fdae := make([]byte, 4)
	_d.BigEndian.PutUint32(_fdae, uint32(_gdegf.PageAssociation))
	if _egbf, _aaabc = _bfedg.Write(_fdae); _aaabc != nil {
		return 0, _fe.Wrap(_aaabc, _bcgb, "\u0034 \u0062y\u0074\u0065\u0020\u0070\u0061g\u0065\u0020n\u0075\u006d\u0062\u0065\u0072")
	}
	return _egbf, nil
}
func (_fdba *GenericRegion) decodeLine(_bcbg, _ccee, _ggac int) error {
	const _cec = "\u0064\u0065\u0063\u006f\u0064\u0065\u004c\u0069\u006e\u0065"
	_fcabc := _fdba.Bitmap.GetByteIndex(0, _bcbg)
	_bebe := _fcabc - _fdba.Bitmap.RowStride
	switch _fdba.GBTemplate {
	case 0:
		if !_fdba.UseExtTemplates {
			return _fdba.decodeTemplate0a(_bcbg, _ccee, _ggac, _fcabc, _bebe)
		}
		return _fdba.decodeTemplate0b(_bcbg, _ccee, _ggac, _fcabc, _bebe)
	case 1:
		return _fdba.decodeTemplate1(_bcbg, _ccee, _ggac, _fcabc, _bebe)
	case 2:
		return _fdba.decodeTemplate2(_bcbg, _ccee, _ggac, _fcabc, _bebe)
	case 3:
		return _fdba.decodeTemplate3(_bcbg, _ccee, _ggac, _fcabc, _bebe)
	}
	return _fe.Errorf(_cec, "\u0069\u006e\u0076a\u006c\u0069\u0064\u0020G\u0042\u0054\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u003a\u0020\u0025\u0064", _fdba.GBTemplate)
}
func (_ggg *RegionSegment) Encode(w _ae.BinaryWriter) (_ebce int, _cbc error) {
	const _edgb = "R\u0065g\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	_bfdfe := make([]byte, 4)
	_d.BigEndian.PutUint32(_bfdfe, _ggg.BitmapWidth)
	_ebce, _cbc = w.Write(_bfdfe)
	if _cbc != nil {
		return 0, _fe.Wrap(_cbc, _edgb, "\u0057\u0069\u0064t\u0068")
	}
	_d.BigEndian.PutUint32(_bfdfe, _ggg.BitmapHeight)
	var _bdbe int
	_bdbe, _cbc = w.Write(_bfdfe)
	if _cbc != nil {
		return 0, _fe.Wrap(_cbc, _edgb, "\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_ebce += _bdbe
	_d.BigEndian.PutUint32(_bfdfe, _ggg.XLocation)
	_bdbe, _cbc = w.Write(_bfdfe)
	if _cbc != nil {
		return 0, _fe.Wrap(_cbc, _edgb, "\u0058L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_ebce += _bdbe
	_d.BigEndian.PutUint32(_bfdfe, _ggg.YLocation)
	_bdbe, _cbc = w.Write(_bfdfe)
	if _cbc != nil {
		return 0, _fe.Wrap(_cbc, _edgb, "\u0059L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_ebce += _bdbe
	if _cbc = w.WriteByte(byte(_ggg.CombinaionOperator) & 0x07); _cbc != nil {
		return 0, _fe.Wrap(_cbc, _edgb, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	_ebce++
	return _ebce, nil
}

type Pager interface {
	GetSegment(int) (*Header, error)
	GetBitmap() (*_df.Bitmap, error)
}

func (_dabff *TextRegion) encodeSymbols(_gedc _ae.BinaryWriter) (_edga int, _fbffb error) {
	const _fceg = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	_fgfc := make([]byte, 4)
	_d.BigEndian.PutUint32(_fgfc, _dabff.NumberOfSymbols)
	if _edga, _fbffb = _gedc.Write(_fgfc); _fbffb != nil {
		return _edga, _fe.Wrap(_fbffb, _fceg, "\u004e\u0075\u006dbe\u0072\u004f\u0066\u0053\u0079\u006d\u0062\u006f\u006c\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073")
	}
	_cacge, _fbffb := _df.NewClassedPoints(_dabff._eegg, _dabff._faba)
	if _fbffb != nil {
		return 0, _fe.Wrap(_fbffb, _fceg, "")
	}
	var _dcfb, _eegb int
	_cagda := _gc.New()
	_cagda.Init()
	if _fbffb = _cagda.EncodeInteger(_gc.IADT, 0); _fbffb != nil {
		return _edga, _fe.Wrap(_fbffb, _fceg, "\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0044\u0054")
	}
	_gafa, _fbffb := _cacge.GroupByY()
	if _fbffb != nil {
		return 0, _fe.Wrap(_fbffb, _fceg, "")
	}
	for _, _bggda := range _gafa {
		_cggb := int(_bggda.YAtIndex(0))
		_geabe := _cggb - _dcfb
		if _fbffb = _cagda.EncodeInteger(_gc.IADT, _geabe); _fbffb != nil {
			return _edga, _fe.Wrap(_fbffb, _fceg, "")
		}
		var _fcfad int
		for _efda, _geca := range _bggda.IntSlice {
			switch _efda {
			case 0:
				_agfgc := int(_bggda.XAtIndex(_efda)) - _eegb
				if _fbffb = _cagda.EncodeInteger(_gc.IAFS, _agfgc); _fbffb != nil {
					return _edga, _fe.Wrap(_fbffb, _fceg, "")
				}
				_eegb += _agfgc
				_fcfad = _eegb
			default:
				_dggd := int(_bggda.XAtIndex(_efda)) - _fcfad
				if _fbffb = _cagda.EncodeInteger(_gc.IADS, _dggd); _fbffb != nil {
					return _edga, _fe.Wrap(_fbffb, _fceg, "")
				}
				_fcfad += _dggd
			}
			_cabbc, _eggdb := _dabff._cdbe.Get(_geca)
			if _eggdb != nil {
				return _edga, _fe.Wrap(_eggdb, _fceg, "")
			}
			_bcfc, _eafg := _dabff._eeaa[_cabbc]
			if !_eafg {
				_bcfc, _eafg = _dabff._gecbc[_cabbc]
				if !_eafg {
					return _edga, _fe.Errorf(_fceg, "\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027\u0025d\u0027\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064 \u0069\u006e\u0020\u0067\u006cob\u0061\u006c\u0020\u0061\u006e\u0064\u0020\u006c\u006f\u0063\u0061\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020\u006d\u0061\u0070", _cabbc)
				}
			}
			if _eggdb = _cagda.EncodeIAID(_dabff._dccc, _bcfc); _eggdb != nil {
				return _edga, _fe.Wrap(_eggdb, _fceg, "")
			}
		}
		if _fbffb = _cagda.EncodeOOB(_gc.IADS); _fbffb != nil {
			return _edga, _fe.Wrap(_fbffb, _fceg, "")
		}
	}
	_cagda.Final()
	_bdfa, _fbffb := _cagda.WriteTo(_gedc)
	if _fbffb != nil {
		return _edga, _fe.Wrap(_fbffb, _fceg, "")
	}
	_edga += int(_bdfa)
	return _edga, nil
}
func (_bagb *PageInformationSegment) readContainsRefinement() error {
	_gfdf, _fcfb := _bagb._fgee.ReadBit()
	if _fcfb != nil {
		return _fcfb
	}
	if _gfdf == 1 {
		_bagb._cadb = true
	}
	return nil
}
func (_ffbd *TextRegion) checkInput() error {
	const _fbbcd = "\u0063\u0068\u0065\u0063\u006b\u0049\u006e\u0070\u0075\u0074"
	if !_ffbd.UseRefinement {
		if _ffbd.SbrTemplate != 0 {
			_af.Log.Debug("\u0053\u0062\u0072Te\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_ffbd.SbrTemplate = 0
		}
	}
	if _ffbd.SbHuffFS == 2 || _ffbd.SbHuffRDWidth == 2 || _ffbd.SbHuffRDHeight == 2 || _ffbd.SbHuffRDX == 2 || _ffbd.SbHuffRDY == 2 {
		return _fe.Error(_fbbcd, "h\u0075\u0066\u0066\u006d\u0061\u006e \u0066\u006c\u0061\u0067\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064")
	}
	if !_ffbd.UseRefinement {
		if _ffbd.SbHuffRSize != 0 {
			_af.Log.Debug("\u0053\u0062\u0048uf\u0066\u0052\u0053\u0069\u007a\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_ffbd.SbHuffRSize = 0
		}
		if _ffbd.SbHuffRDY != 0 {
			_af.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044Y\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_ffbd.SbHuffRDY = 0
		}
		if _ffbd.SbHuffRDX != 0 {
			_af.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044X\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_ffbd.SbHuffRDX = 0
		}
		if _ffbd.SbHuffRDWidth != 0 {
			_af.Log.Debug("\u0053b\u0048\u0075\u0066\u0066R\u0044\u0057\u0069\u0064\u0074h\u0020s\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u00200")
			_ffbd.SbHuffRDWidth = 0
		}
		if _ffbd.SbHuffRDHeight != 0 {
			_af.Log.Debug("\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069g\u0068\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0030")
			_ffbd.SbHuffRDHeight = 0
		}
	}
	return nil
}
func (_edg *GenericRegion) parseHeader() (_dbc error) {
	_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052I\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0050\u0061\u0072s\u0069\u006e\u0067\u0048\u0065\u0061\u0064e\u0072\u002e\u002e\u002e")
	defer func() {
		if _dbc != nil {
			_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e]\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065r\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0077\u0069th\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u0025\u0076", _dbc)
		} else {
			_af.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049C\u002d\u0052\u0045G\u0049\u004f\u004e]\u0020\u0050a\u0072\u0073\u0069\u006e\u0067\u0048e\u0061de\u0072\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0053\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079\u002e\u002e\u002e")
		}
	}()
	var (
		_bbb int
		_cbd uint64
	)
	if _dbc = _edg.RegionSegment.parseHeader(); _dbc != nil {
		return _dbc
	}
	if _, _dbc = _edg._abe.ReadBits(3); _dbc != nil {
		return _dbc
	}
	_bbb, _dbc = _edg._abe.ReadBit()
	if _dbc != nil {
		return _dbc
	}
	if _bbb == 1 {
		_edg.UseExtTemplates = true
	}
	_bbb, _dbc = _edg._abe.ReadBit()
	if _dbc != nil {
		return _dbc
	}
	if _bbb == 1 {
		_edg.IsTPGDon = true
	}
	_cbd, _dbc = _edg._abe.ReadBits(2)
	if _dbc != nil {
		return _dbc
	}
	_edg.GBTemplate = byte(_cbd & 0xf)
	_bbb, _dbc = _edg._abe.ReadBit()
	if _dbc != nil {
		return _dbc
	}
	if _bbb == 1 {
		_edg.IsMMREncoded = true
	}
	if !_edg.IsMMREncoded {
		_cde := 1
		if _edg.GBTemplate == 0 {
			_cde = 4
			if _edg.UseExtTemplates {
				_cde = 12
			}
		}
		if _dbc = _edg.readGBAtPixels(_cde); _dbc != nil {
			return _dbc
		}
	}
	if _dbc = _edg.computeSegmentDataStructure(); _dbc != nil {
		return _dbc
	}
	_af.Log.Trace("\u0025\u0073", _edg)
	return nil
}
func (_dbada *TextRegion) String() string {
	_gaea := &_a.Builder{}
	_gaea.WriteString("\u000a[\u0054E\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u000a")
	_gaea.WriteString(_dbada.RegionInfo.String() + "\u000a")
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053br\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0025\u0076\u000a", _dbada.SbrTemplate))
	_gaea.WriteString(_dg.Sprintf("\u0009-\u0020S\u0062\u0044\u0073\u004f\u0066f\u0073\u0065t\u003a\u0020\u0025\u0076\u000a", _dbada.SbDsOffset))
	_gaea.WriteString(_dg.Sprintf("\t\u002d \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0050i\u0078\u0065\u006c\u003a %\u0076\u000a", _dbada.DefaultPixel))
	_gaea.WriteString(_dg.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0076\u000a", _dbada.CombinationOperator))
	_gaea.WriteString(_dg.Sprintf("\t\u002d \u0049\u0073\u0054\u0072\u0061\u006e\u0073\u0070o\u0073\u0065\u0064\u003a %\u0076\u000a", _dbada.IsTransposed))
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020Re\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0043\u006f\u0072\u006e\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _dbada.ReferenceCorner))
	_gaea.WriteString(_dg.Sprintf("\t\u002d\u0020\u0055\u0073eR\u0065f\u0069\u006e\u0065\u006d\u0065n\u0074\u003a\u0020\u0025\u0076\u000a", _dbada.UseRefinement))
	_gaea.WriteString(_dg.Sprintf("\u0009-\u0020\u0049\u0073\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045n\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0076\u000a", _dbada.IsHuffmanEncoded))
	if _dbada.IsHuffmanEncoded {
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _dbada.SbHuffRSize))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0059:\u0020\u0025\u0076\u000a", _dbada.SbHuffRDY))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0058:\u0020\u0025\u0076\u000a", _dbada.SbHuffRDX))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069\u0067\u0068\u0074\u003a\u0020\u0025v\u000a", _dbada.SbHuffRDHeight))
		_gaea.WriteString(_dg.Sprintf("\t\u002d\u0020\u0053\u0062Hu\u0066f\u0052\u0044\u0057\u0069\u0064t\u0068\u003a\u0020\u0025\u0076\u000a", _dbada.SbHuffRDWidth))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0054\u003a\u0020\u0025\u0076\u000a", _dbada.SbHuffDT))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0053\u003a\u0020\u0025\u0076\u000a", _dbada.SbHuffDS))
		_gaea.WriteString(_dg.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0046\u0053\u003a\u0020\u0025\u0076\u000a", _dbada.SbHuffFS))
	}
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _dbada.SbrATX))
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0059:\u0020\u0025\u0076\u000a", _dbada.SbrATY))
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020N\u0075\u006d\u0062\u0065r\u004f\u0066\u0053\u0079\u006d\u0062\u006fl\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _dbada.NumberOfSymbolInstances))
	_gaea.WriteString(_dg.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _dbada.SbrATX))
	return _gaea.String()
}
func (_aabec *SymbolDictionary) checkInput() error {
	if _aabec.SdHuffDecodeHeightSelection == 2 {
		_af.Log.Debug("\u0053\u0079\u006d\u0062\u006fl\u0020\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u0053e\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0070\u0065r\u006d\u0069\u0074\u0074\u0065\u0064", _aabec.SdHuffDecodeHeightSelection)
	}
	if _aabec.SdHuffDecodeWidthSelection == 2 {
		_af.Log.Debug("\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0057\u0069\u0064t\u0068\u0020\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061l\u0075\u0065\u0020\u006e\u006f\u0074 \u0070\u0065r\u006d\u0069t\u0074e\u0064", _aabec.SdHuffDecodeWidthSelection)
	}
	if _aabec.IsHuffmanEncoded {
		if _aabec.SdTemplate != 0 {
			_af.Log.Debug("\u0053\u0044T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e \u0030\u0029", _aabec.SdTemplate)
		}
		if !_aabec.UseRefinementAggregation {
			if !_aabec.UseRefinementAggregation {
				if _aabec._acgc {
					_af.Log.Debug("\u0049\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074ex\u0074\u0052\u0065\u0074\u0061\u0069n\u0065\u0064\u0020\u003d\u0020\u0074\u0072\u0075\u0065\u0020\u0028\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0066\u0061\u006cs\u0065\u0029")
					_aabec._acgc = false
				}
				if _aabec._dgeg {
					_af.Log.Debug("\u0069s\u0043\u006fd\u0069\u006e\u0067\u0043o\u006e\u0074\u0065x\u0074\u0055\u0073\u0065\u0064\u0020\u003d\u0020\u0074ru\u0065\u0020\u0028s\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020f\u0061\u006cs\u0065\u0029")
					_aabec._dgeg = false
				}
			}
		}
	} else {
		if _aabec.SdHuffBMSizeSelection != 0 {
			_af.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066B\u004d\u0053\u0069\u007a\u0065\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_aabec.SdHuffBMSizeSelection = 0
		}
		if _aabec.SdHuffDecodeWidthSelection != 0 {
			_af.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068\u0053\u0065\u006c\u0065\u0063\u0074\u0069o\u006e\u0020\u0073\u0068\u006fu\u006c\u0064 \u0062\u0065\u0020\u0030")
			_aabec.SdHuffDecodeWidthSelection = 0
		}
		if _aabec.SdHuffDecodeHeightSelection != 0 {
			_af.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0048e\u0069\u0067\u0068\u0074\u0053\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_aabec.SdHuffDecodeHeightSelection = 0
		}
	}
	if !_aabec.UseRefinementAggregation {
		if _aabec.SdrTemplate != 0 {
			_af.Log.Debug("\u0053\u0044\u0052\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u003d\u0020\u0025\u0064\u0020\u0028s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _aabec.SdrTemplate)
			_aabec.SdrTemplate = 0
		}
	}
	if !_aabec.IsHuffmanEncoded || !_aabec.UseRefinementAggregation {
		if _aabec.SdHuffAggInstanceSelection {
			_af.Log.Debug("\u0053d\u0048\u0075f\u0066\u0041\u0067g\u0049\u006e\u0073\u0074\u0061\u006e\u0063e\u0053\u0065\u006c\u0065\u0063\u0074i\u006f\u006e\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _aabec.SdHuffAggInstanceSelection)
		}
	}
	return nil
}
func (_baf *GenericRegion) computeSegmentDataStructure() error {
	_baf.DataOffset = _baf._abe.AbsolutePosition()
	_baf.DataHeaderLength = _baf.DataOffset - _baf.DataHeaderOffset
	_baf.DataLength = int64(_baf._abe.AbsoluteLength()) - _baf.DataHeaderLength
	return nil
}
func (_dabb *HalftoneRegion) computeSegmentDataStructure() error {
	_dabb.DataOffset = _dabb._baab.AbsolutePosition()
	_dabb.DataHeaderLength = _dabb.DataOffset - _dabb.DataHeaderOffset
	_dabb.DataLength = int64(_dabb._baab.AbsoluteLength()) - _dabb.DataHeaderLength
	return nil
}
func (_gca *HalftoneRegion) Init(hd *Header, r *_ae.Reader) error {
	_gca._baab = r
	_gca._bfed = hd
	_gca.RegionSegment = NewRegionSegment(r)
	return _gca.parseHeader()
}
func (_abaad *TextRegion) encodeFlags(_ecaf _ae.BinaryWriter) (_ccge int, _fbdga error) {
	const _bgfef = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _fbdga = _ecaf.WriteBit(int(_abaad.SbrTemplate)); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "s\u0062\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _, _fbdga = _ecaf.WriteBits(uint64(_abaad.SbDsOffset), 5); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0073\u0062\u0044\u0073\u004f\u0066\u0066\u0073\u0065\u0074")
	}
	if _fbdga = _ecaf.WriteBit(int(_abaad.DefaultPixel)); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0044\u0065\u0066a\u0075\u006c\u0074\u0050\u0069\u0078\u0065\u006c")
	}
	if _, _fbdga = _ecaf.WriteBits(uint64(_abaad.CombinationOperator), 2); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0043\u006f\u006d\u0062in\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	if _fbdga = _ecaf.WriteBit(int(_abaad.IsTransposed)); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0069\u0073\u0020\u0074\u0072\u0061\u006e\u0073\u0070\u006f\u0073\u0065\u0064")
	}
	if _, _fbdga = _ecaf.WriteBits(uint64(_abaad.ReferenceCorner), 2); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0072\u0065f\u0065\u0072\u0065n\u0063\u0065\u0020\u0063\u006f\u0072\u006e\u0065\u0072")
	}
	if _, _fbdga = _ecaf.WriteBits(uint64(_abaad.LogSBStrips), 2); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "L\u006f\u0067\u0053\u0042\u0053\u0074\u0072\u0069\u0070\u0073")
	}
	var _cbaa int
	if _abaad.UseRefinement {
		_cbaa = 1
	}
	if _fbdga = _ecaf.WriteBit(_cbaa); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "\u0075\u0073\u0065\u0020\u0072\u0065\u0066\u0069\u006ee\u006d\u0065\u006e\u0074")
	}
	_cbaa = 0
	if _abaad.IsHuffmanEncoded {
		_cbaa = 1
	}
	if _fbdga = _ecaf.WriteBit(_cbaa); _fbdga != nil {
		return _ccge, _fe.Wrap(_fbdga, _bgfef, "u\u0073\u0065\u0020\u0068\u0075\u0066\u0066\u006d\u0061\u006e")
	}
	_ccge = 2
	return _ccge, nil
}
func (_bbbf *TextRegion) setCodingStatistics() error {
	if _bbbf._bgce == nil {
		_bbbf._bgce = _b.NewStats(512, 1)
	}
	if _bbbf._deef == nil {
		_bbbf._deef = _b.NewStats(512, 1)
	}
	if _bbbf._ccac == nil {
		_bbbf._ccac = _b.NewStats(512, 1)
	}
	if _bbbf._cbgce == nil {
		_bbbf._cbgce = _b.NewStats(512, 1)
	}
	if _bbbf._gddf == nil {
		_bbbf._gddf = _b.NewStats(512, 1)
	}
	if _bbbf._abbd == nil {
		_bbbf._abbd = _b.NewStats(512, 1)
	}
	if _bbbf._gbgf == nil {
		_bbbf._gbgf = _b.NewStats(512, 1)
	}
	if _bbbf._bebb == nil {
		_bbbf._bebb = _b.NewStats(1<<uint(_bbbf._adfd), 1)
	}
	if _bbbf._gdgaf == nil {
		_bbbf._gdgaf = _b.NewStats(512, 1)
	}
	if _bbbf._fddb == nil {
		_bbbf._fddb = _b.NewStats(512, 1)
	}
	if _bbbf._fagf == nil {
		var _edfg error
		_bbbf._fagf, _edfg = _b.New(_bbbf._deba)
		if _edfg != nil {
			return _edfg
		}
	}
	return nil
}
func (_dbgg *TextRegion) decodeSymInRefSize() (int64, error) {
	const _ggad = "\u0064e\u0063o\u0064\u0065\u0053\u0079\u006dI\u006e\u0052e\u0066\u0053\u0069\u007a\u0065"
	if _dbgg.SbHuffRSize == 0 {
		_dega, _eeag := _ge.GetStandardTable(1)
		if _eeag != nil {
			return 0, _fe.Wrap(_eeag, _ggad, "")
		}
		return _dega.Decode(_dbgg._deba)
	}
	if _dbgg._ffcfg == nil {
		var (
			_adce int
			_aafe error
		)
		if _dbgg.SbHuffFS == 3 {
			_adce++
		}
		if _dbgg.SbHuffDS == 3 {
			_adce++
		}
		if _dbgg.SbHuffDT == 3 {
			_adce++
		}
		if _dbgg.SbHuffRDWidth == 3 {
			_adce++
		}
		if _dbgg.SbHuffRDHeight == 3 {
			_adce++
		}
		if _dbgg.SbHuffRDX == 3 {
			_adce++
		}
		if _dbgg.SbHuffRDY == 3 {
			_adce++
		}
		_dbgg._ffcfg, _aafe = _dbgg.getUserTable(_adce)
		if _aafe != nil {
			return 0, _fe.Wrap(_aafe, _ggad, "")
		}
	}
	_dded, _egbec := _dbgg._ffcfg.Decode(_dbgg._deba)
	if _egbec != nil {
		return 0, _fe.Wrap(_egbec, _ggad, "")
	}
	return _dded, nil
}
func (_bgga *PatternDictionary) Init(h *Header, r *_ae.Reader) error {
	_bgga._fdcf = r
	return _bgga.parseHeader()
}
func (_gfab *TableSegment) HtLow() int32 { return _gfab._gefag }
func (_cgc *GenericRefinementRegion) decodeSLTP() (int, error) {
	_cgc.Template.setIndex(_cgc._ab)
	return _cgc._eg.DecodeBit(_cgc._ab)
}
func _cgebd(_deab *_ae.Reader, _gfee *Header) *TextRegion {
	_fdef := &TextRegion{_deba: _deab, Header: _gfee, RegionInfo: NewRegionSegment(_deab)}
	return _fdef
}
func (_ebfb *Header) readReferredToSegmentNumbers(_gdea *_ae.Reader, _ffd int) ([]int, error) {
	const _cfd = "\u0072\u0065\u0061\u0064R\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053e\u0067m\u0065\u006e\u0074\u004e\u0075\u006d\u0062e\u0072\u0073"
	_ebbcc := make([]int, _ffd)
	if _ffd > 0 {
		_ebfb.RTSegments = make([]*Header, _ffd)
		var (
			_bafg uint64
			_aaab error
		)
		for _aced := 0; _aced < _ffd; _aced++ {
			_bafg, _aaab = _gdea.ReadBits(byte(_ebfb.referenceSize()) << 3)
			if _aaab != nil {
				return nil, _fe.Wrapf(_aaab, _cfd, "\u0027\u0025\u0064\u0027 \u0072\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0020\u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u0075\u006db\u0065\u0072", _aced)
			}
			_ebbcc[_aced] = int(_bafg & _c.MaxInt32)
		}
	}
	return _ebbcc, nil
}
func (_adb *GenericRegion) decodeTemplate1(_fdda, _ceeg, _eeb int, _eege, _ade int) (_cfeb error) {
	const _gac = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0031"
	var (
		_ffgg, _cga  int
		_aga, _eba   int
		_ddd         byte
		_cafcc, _bea int
	)
	if _fdda >= 1 {
		_ddd, _cfeb = _adb.Bitmap.GetByte(_ade)
		if _cfeb != nil {
			return _fe.Wrap(_cfeb, _gac, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_aga = int(_ddd)
	}
	if _fdda >= 2 {
		_ddd, _cfeb = _adb.Bitmap.GetByte(_ade - _adb.Bitmap.RowStride)
		if _cfeb != nil {
			return _fe.Wrap(_cfeb, _gac, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_eba = int(_ddd) << 5
	}
	_ffgg = ((_aga >> 1) & 0x1f8) | ((_eba >> 1) & 0x1e00)
	for _fcf := 0; _fcf < _eeb; _fcf = _cafcc {
		var (
			_eeac byte
			_bdad int
		)
		_cafcc = _fcf + 8
		if _baaa := _ceeg - _fcf; _baaa > 8 {
			_bdad = 8
		} else {
			_bdad = _baaa
		}
		if _fdda > 0 {
			_aga <<= 8
			if _cafcc < _ceeg {
				_ddd, _cfeb = _adb.Bitmap.GetByte(_ade + 1)
				if _cfeb != nil {
					return _fe.Wrap(_cfeb, _gac, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_aga |= int(_ddd)
			}
		}
		if _fdda > 1 {
			_eba <<= 8
			if _cafcc < _ceeg {
				_ddd, _cfeb = _adb.Bitmap.GetByte(_ade - _adb.Bitmap.RowStride + 1)
				if _cfeb != nil {
					return _fe.Wrap(_cfeb, _gac, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_eba |= int(_ddd) << 5
			}
		}
		for _fccf := 0; _fccf < _bdad; _fccf++ {
			if _adb._dca {
				_cga = _adb.overrideAtTemplate1(_ffgg, _fcf+_fccf, _fdda, int(_eeac), _fccf)
				_adb._bfe.SetIndex(int32(_cga))
			} else {
				_adb._bfe.SetIndex(int32(_ffgg))
			}
			_bea, _cfeb = _adb._caf.DecodeBit(_adb._bfe)
			if _cfeb != nil {
				return _fe.Wrap(_cfeb, _gac, "")
			}
			_eeac |= byte(_bea) << uint(7-_fccf)
			_aeac := uint(8 - _fccf)
			_ffgg = ((_ffgg & 0xefb) << 1) | _bea | ((_aga >> _aeac) & 0x8) | ((_eba >> _aeac) & 0x200)
		}
		if _bad := _adb.Bitmap.SetByte(_eege, _eeac); _bad != nil {
			return _fe.Wrap(_bad, _gac, "")
		}
		_eege++
		_ade++
	}
	return nil
}
func (_abgf *TextRegion) createRegionBitmap() error {
	_abgf.RegionBitmap = _df.New(int(_abgf.RegionInfo.BitmapWidth), int(_abgf.RegionInfo.BitmapHeight))
	if _abgf.DefaultPixel != 0 {
		_abgf.RegionBitmap.SetDefaultPixel()
	}
	return nil
}
func (_eff *HalftoneRegion) computeX(_cdad, _cdfa int) int {
	return _eff.shiftAndFill(int(_eff.HGridX) + _cdad*int(_eff.HRegionY) + _cdfa*int(_eff.HRegionX))
}
func (_bagf *SymbolDictionary) setAtPixels() error {
	if _bagf.IsHuffmanEncoded {
		return nil
	}
	_dfbd := 1
	if _bagf.SdTemplate == 0 {
		_dfbd = 4
	}
	if _gdfcf := _bagf.readAtPixels(_dfbd); _gdfcf != nil {
		return _gdfcf
	}
	return nil
}
func (_egcd *TextRegion) readAmountOfSymbolInstances() error {
	_acea, _caga := _egcd._deba.ReadBits(32)
	if _caga != nil {
		return _caga
	}
	_egcd.NumberOfSymbolInstances = uint32(_acea & _c.MaxUint32)
	_fef := _egcd.RegionInfo.BitmapWidth * _egcd.RegionInfo.BitmapHeight
	if _fef < _egcd.NumberOfSymbolInstances {
		_af.Log.Debug("\u004c\u0069\u006d\u0069t\u0069\u006e\u0067\u0020t\u0068\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020d\u0065\u0063\u006f\u0064e\u0064\u0020\u0073\u0079m\u0062\u006f\u006c\u0020\u0069n\u0073\u0074\u0061\u006e\u0063\u0065\u0073 \u0074\u006f\u0020\u006f\u006ee\u0020\u0070\u0065\u0072\u0020\u0070\u0069\u0078\u0065l\u0020\u0028\u0020\u0025\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0029", _fef, _egcd.NumberOfSymbolInstances)
		_egcd.NumberOfSymbolInstances = _fef
	}
	return nil
}
func (_adda *Header) writeReferredToSegments(_dge _ae.BinaryWriter) (_ecfd int, _babe error) {
	const _aegd = "\u0077\u0072\u0069te\u0052\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	var (
		_gfff uint16
		_caef uint32
	)
	_ddaf := _adda.referenceSize()
	_aeff := 1
	_bagd := make([]byte, _ddaf)
	for _, _ffcf := range _adda.RTSNumbers {
		switch _ddaf {
		case 4:
			_caef = uint32(_ffcf)
			_d.BigEndian.PutUint32(_bagd, _caef)
			_aeff, _babe = _dge.Write(_bagd)
			if _babe != nil {
				return 0, _fe.Wrap(_babe, _aegd, "u\u0069\u006e\u0074\u0033\u0032\u0020\u0073\u0069\u007a\u0065")
			}
		case 2:
			_gfff = uint16(_ffcf)
			_d.BigEndian.PutUint16(_bagd, _gfff)
			_aeff, _babe = _dge.Write(_bagd)
			if _babe != nil {
				return 0, _fe.Wrap(_babe, _aegd, "\u0075\u0069\u006e\u0074\u0031\u0036")
			}
		default:
			if _babe = _dge.WriteByte(byte(_ffcf)); _babe != nil {
				return 0, _fe.Wrap(_babe, _aegd, "\u0075\u0069\u006et\u0038")
			}
		}
		_ecfd += _aeff
	}
	return _ecfd, nil
}
func (_geb *GenericRegion) GetRegionInfo() *RegionSegment { return _geb.RegionSegment }

type TableSegment struct {
	_ebef  *_ae.Reader
	_cadd  int32
	_dfdfb int32
	_bfcc  int32
	_gefag int32
	_gafe  int32
}

func (_agb *GenericRegion) decodeTemplate2(_fba, _cfbf, _bfbg int, _defcb, _gbeb int) (_eecf error) {
	const _egdb = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0032"
	var (
		_ddgb, _abbb int
		_abee, _acf  int
		_gff         byte
		_acd, _affe  int
	)
	if _fba >= 1 {
		_gff, _eecf = _agb.Bitmap.GetByte(_gbeb)
		if _eecf != nil {
			return _fe.Wrap(_eecf, _egdb, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0031")
		}
		_abee = int(_gff)
	}
	if _fba >= 2 {
		_gff, _eecf = _agb.Bitmap.GetByte(_gbeb - _agb.Bitmap.RowStride)
		if _eecf != nil {
			return _fe.Wrap(_eecf, _egdb, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0032")
		}
		_acf = int(_gff) << 4
	}
	_ddgb = (_abee >> 3 & 0x7c) | (_acf >> 3 & 0x380)
	for _eae := 0; _eae < _bfbg; _eae = _acd {
		var (
			_gcfa byte
			_bdd  int
		)
		_acd = _eae + 8
		if _abbc := _cfbf - _eae; _abbc > 8 {
			_bdd = 8
		} else {
			_bdd = _abbc
		}
		if _fba > 0 {
			_abee <<= 8
			if _acd < _cfbf {
				_gff, _eecf = _agb.Bitmap.GetByte(_gbeb + 1)
				if _eecf != nil {
					return _fe.Wrap(_eecf, _egdb, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0030")
				}
				_abee |= int(_gff)
			}
		}
		if _fba > 1 {
			_acf <<= 8
			if _acd < _cfbf {
				_gff, _eecf = _agb.Bitmap.GetByte(_gbeb - _agb.Bitmap.RowStride + 1)
				if _eecf != nil {
					return _fe.Wrap(_eecf, _egdb, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0031")
				}
				_acf |= int(_gff) << 4
			}
		}
		for _fgbc := 0; _fgbc < _bdd; _fgbc++ {
			_bdb := uint(10 - _fgbc)
			if _agb._dca {
				_abbb = _agb.overrideAtTemplate2(_ddgb, _eae+_fgbc, _fba, int(_gcfa), _fgbc)
				_agb._bfe.SetIndex(int32(_abbb))
			} else {
				_agb._bfe.SetIndex(int32(_ddgb))
			}
			_affe, _eecf = _agb._caf.DecodeBit(_agb._bfe)
			if _eecf != nil {
				return _fe.Wrap(_eecf, _egdb, "")
			}
			_gcfa |= byte(_affe << uint(7-_fgbc))
			_ddgb = ((_ddgb & 0x1bd) << 1) | _affe | ((_abee >> _bdb) & 0x4) | ((_acf >> _bdb) & 0x80)
		}
		if _bddf := _agb.Bitmap.SetByte(_defcb, _gcfa); _bddf != nil {
			return _fe.Wrap(_bddf, _egdb, "")
		}
		_defcb++
		_gbeb++
	}
	return nil
}
func (_bfag *HalftoneRegion) shiftAndFill(_feg int) int {
	_feg >>= 8
	if _feg < 0 {
		_fgac := int(_c.Log(float64(_cegb(_feg))) / _c.Log(2))
		_bffg := 31 - _fgac
		for _dda := 1; _dda < _bffg; _dda++ {
			_feg |= 1 << uint(31-_dda)
		}
	}
	return _feg
}
func (_ecac *Header) writeSegmentNumber(_cebg _ae.BinaryWriter) (_cbag int, _ffgf error) {
	_aebe := make([]byte, 4)
	_d.BigEndian.PutUint32(_aebe, _ecac.SegmentNumber)
	if _cbag, _ffgf = _cebg.Write(_aebe); _ffgf != nil {
		return 0, _fe.Wrap(_ffgf, "\u0048e\u0061\u0064\u0065\u0072.\u0077\u0072\u0069\u0074\u0065S\u0065g\u006de\u006e\u0074\u004e\u0075\u006d\u0062\u0065r", "")
	}
	return _cbag, nil
}
func (_ffgb *template0) form(_cdf, _affa, _abb, _gbdd, _bef int16) int16 {
	return (_cdf << 10) | (_affa << 7) | (_abb << 4) | (_gbdd << 1) | _bef
}

type SegmentEncoder interface {
	Encode(_eaacb _ae.BinaryWriter) (_cdfd int, _caca error)
}

func (_eaega *SymbolDictionary) setInSyms() error {
	if _eaega.Header.RTSegments != nil {
		return _eaega.retrieveImportSymbols()
	}
	_eaega._fege = make([]*_df.Bitmap, 0)
	return nil
}
func (_caeb *Header) writeSegmentDataLength(_deeb _ae.BinaryWriter) (_dfa int, _fedc error) {
	_aaga := make([]byte, 4)
	_d.BigEndian.PutUint32(_aaga, uint32(_caeb.SegmentDataLength))
	if _dfa, _fedc = _deeb.Write(_aaga); _fedc != nil {
		return 0, _fe.Wrap(_fedc, "\u0048\u0065a\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u004c\u0065ng\u0074\u0068", "")
	}
	return _dfa, nil
}
func (_beag *TextRegion) readHuffmanFlags() error {
	var (
		_adbc  int
		_gbgd  uint64
		_bbdad error
	)
	_, _bbdad = _beag._deba.ReadBit()
	if _bbdad != nil {
		return _bbdad
	}
	_adbc, _bbdad = _beag._deba.ReadBit()
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffRSize = int8(_adbc)
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffRDY = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffRDX = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffRDHeight = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffRDWidth = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffDT = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffDS = int8(_gbgd) & 0xf
	_gbgd, _bbdad = _beag._deba.ReadBits(2)
	if _bbdad != nil {
		return _bbdad
	}
	_beag.SbHuffFS = int8(_gbgd) & 0xf
	return nil
}
func (_gegc *PatternDictionary) setGbAtPixels() {
	if _gegc.HDTemplate == 0 {
		_gegc.GBAtX = make([]int8, 4)
		_gegc.GBAtY = make([]int8, 4)
		_gegc.GBAtX[0] = -int8(_gegc.HdpWidth)
		_gegc.GBAtY[0] = 0
		_gegc.GBAtX[1] = -3
		_gegc.GBAtY[1] = -1
		_gegc.GBAtX[2] = 2
		_gegc.GBAtY[2] = -2
		_gegc.GBAtX[3] = -2
		_gegc.GBAtY[3] = -2
	} else {
		_gegc.GBAtX = []int8{-int8(_gegc.HdpWidth)}
		_gegc.GBAtY = []int8{0}
	}
}
