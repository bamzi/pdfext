package textshaping

import (
	_a "strings"

	_c "github.com/unidoc/garabic"
	_e "golang.org/x/text/unicode/bidi"
)

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape(text string) (string, error) {
	_bg := _e.Paragraph{}
	_bg.SetString(text)
	_d, _ba := _bg.Order()
	if _ba != nil {
		return "", _ba
	}
	for _bb := 0; _bb < _d.NumRuns(); _bb++ {
		_be := _d.Run(_bb)
		_ee := _be.String()
		if _be.Direction() == _e.RightToLeft {
			var (
				_bgf = _c.Shape(_ee)
				_de  = []rune(_bgf)
				_cb  = make([]rune, len(_de))
			)
			_g := 0
			for _ed := len(_de) - 1; _ed >= 0; _ed-- {
				_cb[_g] = _de[_ed]
				_g++
			}
			_ee = string(_cb)
			text = _a.Replace(text, _a.TrimSpace(_be.String()), _ee, 1)
		}
	}
	return text, nil
}
