package classer

import (
	_d "image"
	_db "math"

	_dbe "github.com/bamzi/pdfext/common"
	_e "github.com/bamzi/pdfext/internal/jbig2/basic"
	_g "github.com/bamzi/pdfext/internal/jbig2/bitmap"
	_ff "github.com/bamzi/pdfext/internal/jbig2/errors"
)

func DefaultSettings() Settings { _ebeb := &Settings{}; _ebeb.SetDefault(); return *_ebeb }

const (
	MaxConnCompWidth = 350
	MaxCharCompWidth = 350
	MaxWordCompWidth = 1000
	MaxCompHeight    = 120
)

type Method int

func (_aac *Classer) getULCorners(_bc *_g.Bitmap, _ggfc *_g.Boxes) error {
	const _ggb = "\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _bc == nil {
		return _ff.Error(_ggb, "\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _ggfc == nil {
		return _ff.Error(_ggb, "\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	if _aac.PtaUL == nil {
		_aac.PtaUL = &_g.Points{}
	}
	_dd := len(*_ggfc)
	var (
		_ac, _fc, _fe, _c   int
		_ce, _ae, _fee, _cd float32
		_ea                 error
		_gbd                *_d.Rectangle
		_cf                 *_g.Bitmap
		_dg                 _d.Point
	)
	for _dc := 0; _dc < _dd; _dc++ {
		_ac = _aac.BaseIndex + _dc
		if _ce, _ae, _ea = _aac.CentroidPoints.GetGeometry(_ac); _ea != nil {
			return _ff.Wrap(_ea, _ggb, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073")
		}
		if _fc, _ea = _aac.ClassIDs.Get(_ac); _ea != nil {
			return _ff.Wrap(_ea, _ggb, "\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074")
		}
		if _fee, _cd, _ea = _aac.CentroidPointsTemplates.GetGeometry(_fc); _ea != nil {
			return _ff.Wrap(_ea, _ggb, "\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_fde := _fee - _ce
		_ee := _cd - _ae
		if _fde >= 0 {
			_fe = int(_fde + 0.5)
		} else {
			_fe = int(_fde - 0.5)
		}
		if _ee >= 0 {
			_c = int(_ee + 0.5)
		} else {
			_c = int(_ee - 0.5)
		}
		if _gbd, _ea = _ggfc.Get(_dc); _ea != nil {
			return _ff.Wrap(_ea, _ggb, "")
		}
		_agf, _eb := _gbd.Min.X, _gbd.Min.Y
		_cf, _ea = _aac.UndilatedTemplates.GetBitmap(_fc)
		if _ea != nil {
			return _ff.Wrap(_ea, _ggb, "\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029")
		}
		_dg, _ea = _ffa(_bc, _agf, _eb, _fe, _c, _cf)
		if _ea != nil {
			return _ff.Wrap(_ea, _ggb, "")
		}
		_aac.PtaUL.AddPoint(float32(_agf-_fe+_dg.X), float32(_eb-_c+_dg.Y))
	}
	return nil
}

type Classer struct {
	BaseIndex               int
	Settings                Settings
	ComponentsNumber        *_e.IntSlice
	TemplateAreas           *_e.IntSlice
	Widths                  map[int]int
	Heights                 map[int]int
	NumberOfClasses         int
	ClassInstances          *_g.BitmapsArray
	UndilatedTemplates      *_g.Bitmaps
	DilatedTemplates        *_g.Bitmaps
	TemplatesSize           _e.IntsMap
	FgTemplates             *_e.NumSlice
	CentroidPoints          *_g.Points
	CentroidPointsTemplates *_g.Points
	ClassIDs                *_e.IntSlice
	ComponentPageNumbers    *_e.IntSlice
	PtaUL                   *_g.Points
	PtaLL                   *_g.Points
}

func (_fgd *Classer) classifyRankHouseNonOne(_afb *_g.Boxes, _eea, _gfe, _gaf *_g.Bitmaps, _fbb *_g.Points, _dgd *_e.NumSlice, _fgg int) (_ffff error) {
	const _gae = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_gfb, _edb, _dfa, _dada       float32
		_egaf, _daaa, _bcgd           int
		_ecgb, _acc, _adf, _cce, _eec *_g.Bitmap
		_dgdg, _afe                   bool
	)
	_ecf := _g.MakePixelSumTab8()
	for _bbc := 0; _bbc < len(_eea.Values); _bbc++ {
		if _acc, _ffff = _gfe.GetBitmap(_bbc); _ffff != nil {
			return _ff.Wrap(_ffff, _gae, "b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _egaf, _ffff = _dgd.GetInt(_bbc); _ffff != nil {
			_dbe.Log.Trace("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bbc, _ffff)
		}
		if _adf, _ffff = _gaf.GetBitmap(_bbc); _ffff != nil {
			return _ff.Wrap(_ffff, _gae, "b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _gfb, _edb, _ffff = _fbb.GetGeometry(_bbc); _ffff != nil {
			return _ff.Wrapf(_ffff, _gae, "\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079")
		}
		_fcd := len(_fgd.UndilatedTemplates.Values)
		_dgdg = false
		_eedf := _dcc(_fgd, _acc)
		for _bcgd = _eedf.Next(); _bcgd > -1; {
			if _cce, _ffff = _fgd.UndilatedTemplates.GetBitmap(_bcgd); _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _daaa, _ffff = _fgd.FgTemplates.GetInt(_bcgd); _ffff != nil {
				_dbe.Log.Trace("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bcgd, _ffff)
			}
			if _eec, _ffff = _fgd.DilatedTemplates.GetBitmap(_bcgd); _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _dfa, _dada, _ffff = _fgd.CentroidPointsTemplates.GetGeometry(_bcgd); _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d")
			}
			_afe, _ffff = _g.RankHausTest(_acc, _adf, _cce, _eec, _gfb-_dfa, _edb-_dada, MaxDiffWidth, MaxDiffHeight, _egaf, _daaa, float32(_fgd.Settings.RankHaus), _ecf)
			if _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "")
			}
			if _afe {
				_dgdg = true
				if _ffff = _fgd.ClassIDs.Add(_bcgd); _ffff != nil {
					return _ff.Wrap(_ffff, _gae, "")
				}
				if _ffff = _fgd.ComponentPageNumbers.Add(_fgg); _ffff != nil {
					return _ff.Wrap(_ffff, _gae, "")
				}
				if _fgd.Settings.KeepClassInstances {
					_def, _eca := _fgd.ClassInstances.GetBitmaps(_bcgd)
					if _eca != nil {
						return _ff.Wrap(_eca, _gae, "\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029")
					}
					if _ecgb, _eca = _eea.GetBitmap(_bbc); _eca != nil {
						return _ff.Wrap(_eca, _gae, "\u0070i\u0078\u0061\u005b\u0069\u005d")
					}
					_def.Values = append(_def.Values, _ecgb)
					_egbg, _eca := _afb.Get(_bbc)
					if _eca != nil {
						return _ff.Wrap(_eca, _gae, "b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
					}
					_def.Boxes = append(_def.Boxes, _egbg)
				}
				break
			}
		}
		if !_dgdg {
			if _ffff = _fgd.ClassIDs.Add(_fcd); _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _ffff = _fgd.ComponentPageNumbers.Add(_fgg); _ffff != nil {
				return _ff.Wrap(_ffff, _gae, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bcbg := &_g.Bitmaps{}
			_ecgb = _eea.Values[_bbc]
			_bcbg.AddBitmap(_ecgb)
			_agb, _fdf := _ecgb.Width, _ecgb.Height
			_fgd.TemplatesSize.Add(uint64(_agb)*uint64(_fdf), _fcd)
			_babc, _fcde := _afb.Get(_bbc)
			if _fcde != nil {
				return _ff.Wrap(_fcde, _gae, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bcbg.AddBox(_babc)
			_fgd.ClassInstances.AddBitmaps(_bcbg)
			_fgd.CentroidPointsTemplates.AddPoint(_gfb, _edb)
			_fgd.UndilatedTemplates.AddBitmap(_acc)
			_fgd.DilatedTemplates.AddBitmap(_adf)
			_fgd.FgTemplates.AddInt(_egaf)
		}
	}
	_fgd.NumberOfClasses = len(_fgd.UndilatedTemplates.Values)
	return nil
}

const (
	RankHaus Method = iota
	Correlation
)

func (_fd *Classer) ComputeLLCorners() (_de error) {
	const _gd = "\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _fd.PtaUL == nil {
		return _ff.Error(_gd, "\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	_fbg := len(*_fd.PtaUL)
	_fd.PtaLL = &_g.Points{}
	var (
		_fbf, _b float32
		_bg, _gg int
		_be      *_g.Bitmap
	)
	for _ba := 0; _ba < _fbg; _ba++ {
		_fbf, _b, _de = _fd.PtaUL.GetGeometry(_ba)
		if _de != nil {
			_dbe.Log.Debug("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _de)
			return _ff.Wrap(_de, _gd, "\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_bg, _de = _fd.ClassIDs.Get(_ba)
		if _de != nil {
			_dbe.Log.Debug("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _de)
			return _ff.Wrap(_de, _gd, "\u0043l\u0061\u0073\u0073\u0049\u0044")
		}
		_be, _de = _fd.UndilatedTemplates.GetBitmap(_bg)
		if _de != nil {
			_dbe.Log.Debug("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _de)
			return _ff.Wrap(_de, _gd, "\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_gg = _be.Height
		_fd.PtaLL.AddPoint(_fbf, _b+float32(_gg))
	}
	return nil
}
func (_ecd *Classer) classifyCorrelation(_bdd *_g.Boxes, _efd *_g.Bitmaps, _dbd int) error {
	const _cc = "\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e"
	if _bdd == nil {
		return _ff.Error(_cc, "\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _efd == nil {
		return _ff.Error(_cc, "\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_egb := len(_efd.Values)
	if _egb == 0 {
		_dbe.Log.Debug("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079")
		return nil
	}
	var (
		_ccf, _cdd *_g.Bitmap
		_bde       error
	)
	_gf := &_g.Bitmaps{Values: make([]*_g.Bitmap, _egb)}
	for _dga, _fba := range _efd.Values {
		_cdd, _bde = _fba.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _bde != nil {
			return _ff.Wrap(_bde, _cc, "")
		}
		_gf.Values[_dga] = _cdd
	}
	_ga := _ecd.FgTemplates
	_egd := _g.MakePixelSumTab8()
	_gfc := _g.MakePixelCentroidTab8()
	_ge := make([]int, _egb)
	_ffb := make([][]int, _egb)
	_dgg := _g.Points(make([]_g.Point, _egb))
	_bdde := &_dgg
	var (
		_df, _gba       int
		_ffd, _ca, _fed int
		_agec, _eaff    int
		_gea            byte
	)
	for _bb, _fedf := range _gf.Values {
		_ffb[_bb] = make([]int, _fedf.Height)
		_df = 0
		_gba = 0
		_ca = (_fedf.Height - 1) * _fedf.RowStride
		_ffd = 0
		for _eaff = _fedf.Height - 1; _eaff >= 0; _eaff, _ca = _eaff-1, _ca-_fedf.RowStride {
			_ffb[_bb][_eaff] = _ffd
			_fed = 0
			for _agec = 0; _agec < _fedf.RowStride; _agec++ {
				_gea = _fedf.Data[_ca+_agec]
				_fed += _egd[_gea]
				_df += _gfc[_gea] + _agec*8*_egd[_gea]
			}
			_ffd += _fed
			_gba += _fed * _eaff
		}
		_ge[_bb] = _ffd
		if _ffd > 0 {
			(*_bdde)[_bb] = _g.Point{X: float32(_df) / float32(_ffd), Y: float32(_gba) / float32(_ffd)}
		} else {
			(*_bdde)[_bb] = _g.Point{X: float32(_fedf.Width) / float32(2), Y: float32(_fedf.Height) / float32(2)}
		}
	}
	if _bde = _ecd.CentroidPoints.Add(_bdde); _bde != nil {
		return _ff.Wrap(_bde, _cc, "\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064")
	}
	var (
		_dad, _cfa, _ffdb      int
		_aeg                   float64
		_ade, _cdb, _fca, _ccd float32
		_eba, _faa             _g.Point
		_ffg                   bool
		_ddb                   *similarTemplatesFinder
		_fab                   int
		_acf                   *_g.Bitmap
		_dba                   *_d.Rectangle
		_cg                    *_g.Bitmaps
	)
	for _fab, _cdd = range _gf.Values {
		_cfa = _ge[_fab]
		if _ade, _cdb, _bde = _bdde.GetGeometry(_fab); _bde != nil {
			return _ff.Wrap(_bde, _cc, "\u0070t\u0061\u0020\u002d\u0020\u0069")
		}
		_ffg = false
		_aabb := len(_ecd.UndilatedTemplates.Values)
		_ddb = _dcc(_ecd, _cdd)
		for _gcd := _ddb.Next(); _gcd > -1; {
			if _acf, _bde = _ecd.UndilatedTemplates.GetBitmap(_gcd); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032")
			}
			if _ffdb, _bde = _ga.GetInt(_gcd); _bde != nil {
				_dbe.Log.Trace("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bde)
			}
			if _fca, _ccd, _bde = _ecd.CentroidPointsTemplates.GetGeometry(_gcd); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020")
			}
			if _ecd.Settings.WeightFactor > 0.0 {
				if _dad, _bde = _ecd.TemplateAreas.Get(_gcd); _bde != nil {
					_dbe.Log.Trace("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076", _bde)
				}
				_aeg = _ecd.Settings.Thresh + (1.0-_ecd.Settings.Thresh)*_ecd.Settings.WeightFactor*float64(_ffdb)/float64(_dad)
			} else {
				_aeg = _ecd.Settings.Thresh
			}
			_agfa, _bbe := _g.CorrelationScoreThresholded(_cdd, _acf, _cfa, _ffdb, _eba.X-_faa.X, _eba.Y-_faa.Y, MaxDiffWidth, MaxDiffHeight, _egd, _ffb[_fab], float32(_aeg))
			if _bbe != nil {
				return _ff.Wrap(_bbe, _cc, "")
			}
			if _aaf {
				var (
					_cee, _eafd float64
					_adc, _bag  int
				)
				_cee, _bbe = _g.CorrelationScore(_cdd, _acf, _cfa, _ffdb, _ade-_fca, _cdb-_ccd, MaxDiffWidth, MaxDiffHeight, _egd)
				if _bbe != nil {
					return _ff.Wrap(_bbe, _cc, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_eafd, _bbe = _g.CorrelationScoreSimple(_cdd, _acf, _cfa, _ffdb, _ade-_fca, _cdb-_ccd, MaxDiffWidth, MaxDiffHeight, _egd)
				if _bbe != nil {
					return _ff.Wrap(_bbe, _cc, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_adc = int(_db.Sqrt(_cee * float64(_cfa) * float64(_ffdb)))
				_bag = int(_db.Sqrt(_eafd * float64(_cfa) * float64(_ffdb)))
				if (_cee >= _aeg) != (_eafd >= _aeg) {
					return _ff.Errorf(_cc, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066", _adc, _cee, _cee >= _aeg, _bag, _eafd, _eafd >= _aeg, _cee-_eafd)
				}
				if _cee >= _aeg != _agfa {
					return _ff.Errorf(_cc, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076", _cee, _cee*float64(_cfa)*float64(_ffdb), _adc, _aeg, float32(_aeg)*float32(_cfa)*float32(_ffdb), _agfa)
				}
			}
			if _agfa {
				_ffg = true
				if _bbe = _ecd.ClassIDs.Add(_gcd); _bbe != nil {
					return _ff.Wrap(_bbe, _cc, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _bbe = _ecd.ComponentPageNumbers.Add(_dbd); _bbe != nil {
					return _ff.Wrap(_bbe, _cc, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _ecd.Settings.KeepClassInstances {
					if _ccf, _bbe = _efd.GetBitmap(_fab); _bbe != nil {
						return _ff.Wrap(_bbe, _cc, "\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069")
					}
					if _cg, _bbe = _ecd.ClassInstances.GetBitmaps(_gcd); _bbe != nil {
						return _ff.Wrap(_bbe, _cc, "K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073")
					}
					_cg.AddBitmap(_ccf)
					if _dba, _bbe = _bdd.Get(_fab); _bbe != nil {
						return _ff.Wrap(_bbe, _cc, "\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073")
					}
					_cg.AddBox(_dba)
				}
				break
			}
		}
		if !_ffg {
			if _bde = _ecd.ClassIDs.Add(_aabb); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _bde = _ecd.ComponentPageNumbers.Add(_dbd); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cg = &_g.Bitmaps{}
			if _ccf, _bde = _efd.GetBitmap(_fab); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cg.AddBitmap(_ccf)
			_ccfd, _ebea := _ccf.Width, _ccf.Height
			_bdg := uint64(_ebea) * uint64(_ccfd)
			_ecd.TemplatesSize.Add(_bdg, _aabb)
			if _dba, _bde = _bdd.Get(_fab); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cg.AddBox(_dba)
			_ecd.ClassInstances.AddBitmaps(_cg)
			_ecd.CentroidPointsTemplates.AddPoint(_ade, _cdb)
			_ecd.FgTemplates.AddInt(_cfa)
			_ecd.UndilatedTemplates.AddBitmap(_ccf)
			_dad = (_cdd.Width - 2*JbAddedPixels) * (_cdd.Height - 2*JbAddedPixels)
			if _bde = _ecd.TemplateAreas.Add(_dad); _bde != nil {
				return _ff.Wrap(_bde, _cc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
		}
	}
	_ecd.NumberOfClasses = len(_ecd.UndilatedTemplates.Values)
	return nil
}
func Init(settings Settings) (*Classer, error) {
	const _ef = "\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074"
	_a := &Classer{Settings: settings, Widths: map[int]int{}, Heights: map[int]int{}, TemplatesSize: _e.IntsMap{}, TemplateAreas: &_e.IntSlice{}, ComponentPageNumbers: &_e.IntSlice{}, ClassIDs: &_e.IntSlice{}, ComponentsNumber: &_e.IntSlice{}, CentroidPoints: &_g.Points{}, CentroidPointsTemplates: &_g.Points{}, UndilatedTemplates: &_g.Bitmaps{}, DilatedTemplates: &_g.Bitmaps{}, ClassInstances: &_g.BitmapsArray{}, FgTemplates: &_e.NumSlice{}}
	if _ec := _a.Settings.Validate(); _ec != nil {
		return nil, _ff.Wrap(_ec, _ef, "")
	}
	return _a, nil
}
func _ffa(_eaf *_g.Bitmap, _gc, _ggfa, _aab, _bab int, _bcc *_g.Bitmap) (_ecg _d.Point, _gdb error) {
	const _bf = "\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg"
	if _eaf == nil {
		return _ecg, _ff.Error(_bf, "\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	if _bcc == nil {
		return _ecg, _ff.Error(_bf, "t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_age, _deb := _bcc.Width, _bcc.Height
	_fea, _eafg := _gc-_aab-JbAddedPixels, _ggfa-_bab-JbAddedPixels
	_dbe.Log.Trace("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027", _gc, _ggfa, _age, _deb, _fea, _eafg)
	_agg, _gdb := _g.Rect(_fea, _eafg, _age, _deb)
	if _gdb != nil {
		return _ecg, _ff.Wrap(_gdb, _bf, "")
	}
	_ad, _, _gdb := _eaf.ClipRectangle(_agg)
	if _gdb != nil {
		_dbe.Log.Error("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076", _agg)
		return _ecg, _ff.Wrap(_gdb, _bf, "")
	}
	_aag := _g.New(_ad.Width, _ad.Height)
	_dbc := _db.MaxInt32
	var _ed, _bgg, _fa, _ebe, _daa int
	for _ed = -1; _ed <= 1; _ed++ {
		for _bgg = -1; _bgg <= 1; _bgg++ {
			if _, _gdb = _g.Copy(_aag, _ad); _gdb != nil {
				return _ecg, _ff.Wrap(_gdb, _bf, "")
			}
			if _gdb = _aag.RasterOperation(_bgg, _ed, _age, _deb, _g.PixSrcXorDst, _bcc, 0, 0); _gdb != nil {
				return _ecg, _ff.Wrap(_gdb, _bf, "")
			}
			_fa = _aag.CountPixels()
			if _fa < _dbc {
				_ebe = _bgg
				_daa = _ed
				_dbc = _fa
			}
		}
	}
	_ecg.X = _ebe
	_ecg.Y = _daa
	return _ecg, nil
}
func _dcc(_bgb *Classer, _dca *_g.Bitmap) *similarTemplatesFinder {
	return &similarTemplatesFinder{Width: _dca.Width, Height: _dca.Height, Classer: _bgb}
}

const JbAddedPixels = 6

func (_fce *Classer) classifyRankHaus(_adg *_g.Boxes, _af *_g.Bitmaps, _gge int) error {
	const _ega = "\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073"
	if _adg == nil {
		return _ff.Error(_ega, "\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _af == nil {
		return _ff.Error(_ega, "\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_gcg := len(_af.Values)
	if _gcg == 0 {
		return _ff.Error(_ega, "e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073")
	}
	_ffaf := _af.CountPixels()
	_dab := _fce.Settings.SizeHaus
	_dfe := _g.SelCreateBrick(_dab, _dab, _dab/2, _dab/2, _g.SelHit)
	_cde := &_g.Bitmaps{Values: make([]*_g.Bitmap, _gcg)}
	_fcg := &_g.Bitmaps{Values: make([]*_g.Bitmap, _gcg)}
	var (
		_fdg, _bcb, _gef *_g.Bitmap
		_dgc             error
	)
	for _fbge := 0; _fbge < _gcg; _fbge++ {
		_fdg, _dgc = _af.GetBitmap(_fbge)
		if _dgc != nil {
			return _ff.Wrap(_dgc, _ega, "")
		}
		_bcb, _dgc = _fdg.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _dgc != nil {
			return _ff.Wrap(_dgc, _ega, "")
		}
		_gef, _dgc = _g.Dilate(nil, _bcb, _dfe)
		if _dgc != nil {
			return _ff.Wrap(_dgc, _ega, "")
		}
		_cde.Values[_gcg] = _bcb
		_fcg.Values[_gcg] = _gef
	}
	_afd, _dgc := _g.Centroids(_cde.Values)
	if _dgc != nil {
		return _ff.Wrap(_dgc, _ega, "")
	}
	if _dgc = _afd.Add(_fce.CentroidPoints); _dgc != nil {
		_dbe.Log.Trace("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064")
	}
	if _fce.Settings.RankHaus == 1.0 {
		_dgc = _fce.classifyRankHouseOne(_adg, _af, _cde, _fcg, _afd, _gge)
	} else {
		_dgc = _fce.classifyRankHouseNonOne(_adg, _af, _cde, _fcg, _afd, _ffaf, _gge)
	}
	if _dgc != nil {
		return _ff.Wrap(_dgc, _ega, "")
	}
	return nil
}

const (
	MaxDiffWidth  = 2
	MaxDiffHeight = 2
)

type similarTemplatesFinder struct {
	Classer        *Classer
	Width          int
	Height         int
	Index          int
	CurrentNumbers []int
	N              int
}

func (_ggc *similarTemplatesFinder) Next() int {
	var (
		_cge, _gff, _bcac, _gfbb int
		_cgc                     bool
		_bfe                     *_g.Bitmap
		_geaa                    error
	)
	for {
		if _ggc.Index >= 25 {
			return -1
		}
		_gff = _ggc.Width + TwoByTwoWalk[2*_ggc.Index]
		_cge = _ggc.Height + TwoByTwoWalk[2*_ggc.Index+1]
		if _cge < 1 || _gff < 1 {
			_ggc.Index++
			continue
		}
		if len(_ggc.CurrentNumbers) == 0 {
			_ggc.CurrentNumbers, _cgc = _ggc.Classer.TemplatesSize.GetSlice(uint64(_gff) * uint64(_cge))
			if !_cgc {
				_ggc.Index++
				continue
			}
			_ggc.N = 0
		}
		_bcac = len(_ggc.CurrentNumbers)
		for ; _ggc.N < _bcac; _ggc.N++ {
			_gfbb = _ggc.CurrentNumbers[_ggc.N]
			_bfe, _geaa = _ggc.Classer.DilatedTemplates.GetBitmap(_gfbb)
			if _geaa != nil {
				_dbe.Log.Debug("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020")
				return 0
			}
			if _bfe.Width-2*JbAddedPixels == _gff && _bfe.Height-2*JbAddedPixels == _cge {
				return _gfbb
			}
		}
		_ggc.Index++
		_ggc.CurrentNumbers = nil
	}
}
func (_bcg *Classer) verifyMethod(_bd Method) error {
	if _bd != RankHaus && _bd != Correlation {
		return _ff.Error("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064", "\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064")
	}
	return nil
}

var _aaf bool

func (_eeb Settings) Validate() error {
	const _fga = "\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065"
	if _eeb.Thresh < 0.4 || _eeb.Thresh > 0.98 {
		return _ff.Error(_fga, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d")
	}
	if _eeb.WeightFactor < 0.0 || _eeb.WeightFactor > 1.0 {
		return _ff.Error(_fga, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _eeb.RankHaus < 0.5 || _eeb.RankHaus > 1.0 {
		return _ff.Error(_fga, "\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _eeb.SizeHaus < 1 || _eeb.SizeHaus > 10 {
		return _ff.Error(_fga, "\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]")
	}
	switch _eeb.Components {
	case _g.ComponentConn, _g.ComponentCharacters, _g.ComponentWords:
	default:
		return _ff.Error(_fga, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et")
	}
	return nil
}
func (_ag *Classer) AddPage(inputPage *_g.Bitmap, pageNumber int, method Method) (_aa error) {
	const _fb = "\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_ag.Widths[pageNumber] = inputPage.Width
	_ag.Heights[pageNumber] = inputPage.Height
	if _aa = _ag.verifyMethod(method); _aa != nil {
		return _ff.Wrap(_aa, _fb, "")
	}
	_fbd, _ab, _aa := inputPage.GetComponents(_ag.Settings.Components, _ag.Settings.MaxCompWidth, _ag.Settings.MaxCompHeight)
	if _aa != nil {
		return _ff.Wrap(_aa, _fb, "")
	}
	_dbe.Log.Debug("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _fbd)
	if _aa = _ag.addPageComponents(inputPage, _ab, _fbd, pageNumber, method); _aa != nil {
		return _ff.Wrap(_aa, _fb, "")
	}
	return nil
}
func (_cbg *Settings) SetDefault() {
	if _cbg.MaxCompWidth == 0 {
		switch _cbg.Components {
		case _g.ComponentConn:
			_cbg.MaxCompWidth = MaxConnCompWidth
		case _g.ComponentCharacters:
			_cbg.MaxCompWidth = MaxCharCompWidth
		case _g.ComponentWords:
			_cbg.MaxCompWidth = MaxWordCompWidth
		}
	}
	if _cbg.MaxCompHeight == 0 {
		_cbg.MaxCompHeight = MaxCompHeight
	}
	if _cbg.Thresh == 0.0 {
		_cbg.Thresh = 0.9
	}
	if _cbg.WeightFactor == 0.0 {
		_cbg.WeightFactor = 0.75
	}
	if _cbg.RankHaus == 0.0 {
		_cbg.RankHaus = 0.97
	}
	if _cbg.SizeHaus == 0 {
		_cbg.SizeHaus = 2
	}
}
func (_ecde *Classer) classifyRankHouseOne(_aeb *_g.Boxes, _baa, _ggg, _fffd *_g.Bitmaps, _gggd *_g.Points, _cb int) (_aad error) {
	const _gbe = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_cag, _ccde, _gfcg, _bca      float32
		_dge                          int
		_bbb, _acb, _cbc, _cgd, _ebeg *_g.Bitmap
		_dgea, _bcbc                  bool
	)
	for _ecc := 0; _ecc < len(_baa.Values); _ecc++ {
		_acb = _ggg.Values[_ecc]
		_cbc = _fffd.Values[_ecc]
		_cag, _ccde, _aad = _gggd.GetGeometry(_ecc)
		if _aad != nil {
			return _ff.Wrapf(_aad, _gbe, "\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_eed := len(_ecde.UndilatedTemplates.Values)
		_dgea = false
		_cda := _dcc(_ecde, _acb)
		for _dge = _cda.Next(); _dge > -1; {
			_cgd, _aad = _ecde.UndilatedTemplates.GetBitmap(_dge)
			if _aad != nil {
				return _ff.Wrap(_aad, _gbe, "\u0062\u006d\u0033")
			}
			_ebeg, _aad = _ecde.DilatedTemplates.GetBitmap(_dge)
			if _aad != nil {
				return _ff.Wrap(_aad, _gbe, "\u0062\u006d\u0034")
			}
			_gfcg, _bca, _aad = _ecde.CentroidPointsTemplates.GetGeometry(_dge)
			if _aad != nil {
				return _ff.Wrap(_aad, _gbe, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073")
			}
			_bcbc, _aad = _g.HausTest(_acb, _cbc, _cgd, _ebeg, _cag-_gfcg, _ccde-_bca, MaxDiffWidth, MaxDiffHeight)
			if _aad != nil {
				return _ff.Wrap(_aad, _gbe, "")
			}
			if _bcbc {
				_dgea = true
				if _aad = _ecde.ClassIDs.Add(_dge); _aad != nil {
					return _ff.Wrap(_aad, _gbe, "")
				}
				if _aad = _ecde.ComponentPageNumbers.Add(_cb); _aad != nil {
					return _ff.Wrap(_aad, _gbe, "")
				}
				if _ecde.Settings.KeepClassInstances {
					_bec, _ecdf := _ecde.ClassInstances.GetBitmaps(_dge)
					if _ecdf != nil {
						return _ff.Wrap(_ecdf, _gbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_bbb, _ecdf = _baa.GetBitmap(_ecc)
					if _ecdf != nil {
						return _ff.Wrap(_ecdf, _gbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_bec.AddBitmap(_bbb)
					_abc, _ecdf := _aeb.Get(_ecc)
					if _ecdf != nil {
						return _ff.Wrap(_ecdf, _gbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_bec.AddBox(_abc)
				}
				break
			}
		}
		if !_dgea {
			if _aad = _ecde.ClassIDs.Add(_eed); _aad != nil {
				return _ff.Wrap(_aad, _gbe, "")
			}
			if _aad = _ecde.ComponentPageNumbers.Add(_cb); _aad != nil {
				return _ff.Wrap(_aad, _gbe, "")
			}
			_bfb := &_g.Bitmaps{}
			_bbb, _aad = _baa.GetBitmap(_ecc)
			if _aad != nil {
				return _ff.Wrap(_aad, _gbe, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bfb.Values = append(_bfb.Values, _bbb)
			_ced, _cfg := _bbb.Width, _bbb.Height
			_ecde.TemplatesSize.Add(uint64(_cfg)*uint64(_ced), _eed)
			_eef, _fda := _aeb.Get(_ecc)
			if _fda != nil {
				return _ff.Wrap(_fda, _gbe, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bfb.AddBox(_eef)
			_ecde.ClassInstances.AddBitmaps(_bfb)
			_ecde.CentroidPointsTemplates.AddPoint(_cag, _ccde)
			_ecde.UndilatedTemplates.AddBitmap(_acb)
			_ecde.DilatedTemplates.AddBitmap(_cbc)
		}
	}
	return nil
}

var TwoByTwoWalk = []int{0, 0, 0, 1, -1, 0, 0, -1, 1, 0, -1, 1, 1, 1, -1, -1, 1, -1, 0, -2, 2, 0, 0, 2, -2, 0, -1, -2, 1, -2, 2, -1, 2, 1, 1, 2, -1, 2, -2, 1, -2, -1, -2, -2, 2, -2, 2, 2, -2, 2}

func (_agd *Classer) addPageComponents(_eg *_g.Bitmap, _fg *_g.Boxes, _dec *_g.Bitmaps, _da int, _daf Method) error {
	const _abb = "\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s"
	if _eg == nil {
		return _ff.Error(_abb, "\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065")
	}
	if _fg == nil || _dec == nil || len(*_fg) == 0 {
		_dbe.Log.Trace("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064", _eg)
		return nil
	}
	var _gb error
	switch _daf {
	case RankHaus:
		_gb = _agd.classifyRankHaus(_fg, _dec, _da)
	case Correlation:
		_gb = _agd.classifyCorrelation(_fg, _dec, _da)
	default:
		_dbe.Log.Debug("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027", _daf)
		return _ff.Error(_abb, "\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
	if _gb != nil {
		return _ff.Wrap(_gb, _abb, "")
	}
	if _gb = _agd.getULCorners(_eg, _fg); _gb != nil {
		return _ff.Wrap(_gb, _abb, "")
	}
	_ggf := len(*_fg)
	_agd.BaseIndex += _ggf
	if _gb = _agd.ComponentsNumber.Add(_ggf); _gb != nil {
		return _ff.Wrap(_gb, _abb, "")
	}
	return nil
}

type Settings struct {
	MaxCompWidth       int
	MaxCompHeight      int
	SizeHaus           int
	RankHaus           float64
	Thresh             float64
	WeightFactor       float64
	KeepClassInstances bool
	Components         _g.Component
	Method             Method
}
