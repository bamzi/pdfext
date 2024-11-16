package sampling

import (
	_b "io"

	_ff "github.com/bamzi/pdfext/internal/bitwise"
	_e "github.com/bamzi/pdfext/internal/imageutil"
)

func (_aa *Reader) ReadSample() (uint32, error) {
	if _aa._fgc == _aa._a.Height {
		return 0, _b.EOF
	}
	_be, _dc := _aa._c.ReadBits(byte(_aa._a.BitsPerComponent))
	if _dc != nil {
		return 0, _dc
	}
	_aa._bg--
	if _aa._bg == 0 {
		_aa._bg = _aa._a.ColorComponents
		_aa._fg++
	}
	if _aa._fg == _aa._a.Width {
		if _aa._ad {
			_aa._c.ConsumeRemainingBits()
		}
		_aa._fg = 0
		_aa._fgc++
	}
	return uint32(_be), nil
}

type SampleWriter interface {
	WriteSample(_bb uint32) error
	WriteSamples(_db []uint32) error
}

func ResampleBytes(data []byte, bitsPerSample int) []uint32 {
	var _fgcg []uint32
	_cgb := bitsPerSample
	var _ed uint32
	var _af byte
	_ce := 0
	_ae := 0
	_gc := 0
	for _gc < len(data) {
		if _ce > 0 {
			_ef := _ce
			if _cgb < _ef {
				_ef = _cgb
			}
			_ed = (_ed << uint(_ef)) | uint32(_af>>uint(8-_ef))
			_ce -= _ef
			if _ce > 0 {
				_af = _af << uint(_ef)
			} else {
				_af = 0
			}
			_cgb -= _ef
			if _cgb == 0 {
				_fgcg = append(_fgcg, _ed)
				_cgb = bitsPerSample
				_ed = 0
				_ae++
			}
		} else {
			_fe := data[_gc]
			_gc++
			_bgd := 8
			if _cgb < _bgd {
				_bgd = _cgb
			}
			_ce = 8 - _bgd
			_ed = (_ed << uint(_bgd)) | uint32(_fe>>uint(_ce))
			if _bgd < 8 {
				_af = _fe << uint(_bgd)
			}
			_cgb -= _bgd
			if _cgb == 0 {
				_fgcg = append(_fgcg, _ed)
				_cgb = bitsPerSample
				_ed = 0
				_ae++
			}
		}
	}
	for _ce >= bitsPerSample {
		_afc := _ce
		if _cgb < _afc {
			_afc = _cgb
		}
		_ed = (_ed << uint(_afc)) | uint32(_af>>uint(8-_afc))
		_ce -= _afc
		if _ce > 0 {
			_af = _af << uint(_afc)
		} else {
			_af = 0
		}
		_cgb -= _afc
		if _cgb == 0 {
			_fgcg = append(_fgcg, _ed)
			_cgb = bitsPerSample
			_ed = 0
			_ae++
		}
	}
	return _fgcg
}
func (_cg *Reader) ReadSamples(samples []uint32) (_g error) {
	for _ab := 0; _ab < len(samples); _ab++ {
		samples[_ab], _g = _cg.ReadSample()
		if _g != nil {
			return _g
		}
	}
	return nil
}

type Reader struct {
	_a             _e.ImageBase
	_c             *_ff.Reader
	_fg, _fgc, _bg int
	_ad            bool
}

func NewWriter(img _e.ImageBase) *Writer {
	return &Writer{_bac: _ff.NewWriterMSB(img.Data), _gf: img, _ac: img.ColorComponents, _ee: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
func NewReader(img _e.ImageBase) *Reader {
	return &Reader{_c: _ff.NewReader(img.Data), _a: img, _bg: img.ColorComponents, _ad: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
func ResampleUint32(data []uint32, bitsPerInputSample int, bitsPerOutputSample int) []uint32 {
	var _fb []uint32
	_ec := bitsPerOutputSample
	var _gca uint32
	var _eff uint32
	_ecf := 0
	_dcd := 0
	_bf := 0
	for _bf < len(data) {
		if _ecf > 0 {
			_fbg := _ecf
			if _ec < _fbg {
				_fbg = _ec
			}
			_gca = (_gca << uint(_fbg)) | (_eff >> uint(bitsPerInputSample-_fbg))
			_ecf -= _fbg
			if _ecf > 0 {
				_eff = _eff << uint(_fbg)
			} else {
				_eff = 0
			}
			_ec -= _fbg
			if _ec == 0 {
				_fb = append(_fb, _gca)
				_ec = bitsPerOutputSample
				_gca = 0
				_dcd++
			}
		} else {
			_gcab := data[_bf]
			_bf++
			_ecb := bitsPerInputSample
			if _ec < _ecb {
				_ecb = _ec
			}
			_ecf = bitsPerInputSample - _ecb
			_gca = (_gca << uint(_ecb)) | (_gcab >> uint(_ecf))
			if _ecb < bitsPerInputSample {
				_eff = _gcab << uint(_ecb)
			}
			_ec -= _ecb
			if _ec == 0 {
				_fb = append(_fb, _gca)
				_ec = bitsPerOutputSample
				_gca = 0
				_dcd++
			}
		}
	}
	for _ecf >= bitsPerOutputSample {
		_ba := _ecf
		if _ec < _ba {
			_ba = _ec
		}
		_gca = (_gca << uint(_ba)) | (_eff >> uint(bitsPerInputSample-_ba))
		_ecf -= _ba
		if _ecf > 0 {
			_eff = _eff << uint(_ba)
		} else {
			_eff = 0
		}
		_ec -= _ba
		if _ec == 0 {
			_fb = append(_fb, _gca)
			_ec = bitsPerOutputSample
			_gca = 0
			_dcd++
		}
	}
	if _ec > 0 && _ec < bitsPerOutputSample {
		_gca <<= uint(_ec)
		_fb = append(_fb, _gca)
	}
	return _fb
}
func (_dbc *Writer) WriteSamples(samples []uint32) error {
	for _ca := 0; _ca < len(samples); _ca++ {
		if _da := _dbc.WriteSample(samples[_ca]); _da != nil {
			return _da
		}
	}
	return nil
}

type SampleReader interface {
	ReadSample() (uint32, error)
	ReadSamples(_d []uint32) error
}
type Writer struct {
	_gf      _e.ImageBase
	_bac     *_ff.Writer
	_ge, _ac int
	_ee      bool
}

func (_fbc *Writer) WriteSample(sample uint32) error {
	if _, _ede := _fbc._bac.WriteBits(uint64(sample), _fbc._gf.BitsPerComponent); _ede != nil {
		return _ede
	}
	_fbc._ac--
	if _fbc._ac == 0 {
		_fbc._ac = _fbc._gf.ColorComponents
		_fbc._ge++
	}
	if _fbc._ge == _fbc._gf.Width {
		if _fbc._ee {
			_fbc._bac.FinishByte()
		}
		_fbc._ge = 0
	}
	return nil
}
