// Package creator is used for quickly generating pages and content with a simple interface.
// It is built on top of the model package to provide access to the most common
// operations such as creating text and image reports and manipulating existing pages.
package creator

import (
	_c "bytes"
	_gf "encoding/xml"
	_bd "errors"
	_f "fmt"
	_e "image"
	_gg "io"
	_ac "log"
	_fc "math"
	_b "os"
	_d "regexp"
	_fd "sort"
	_fg "strconv"
	_eg "strings"
	_g "text/template"
	_fe "unicode"

	_fec "github.com/bamzi/pdfext/common"
	_dg "github.com/bamzi/pdfext/contentstream"
	_gga "github.com/bamzi/pdfext/contentstream/draw"
	_bc "github.com/bamzi/pdfext/core"
	_dd "github.com/bamzi/pdfext/internal/graphic2d/svg"
	_ce "github.com/bamzi/pdfext/internal/integrations/unichart"
	_aa "github.com/bamzi/pdfext/internal/license"
	_de "github.com/bamzi/pdfext/internal/transform"
	_fgd "github.com/bamzi/pdfext/model"
	_bb "github.com/gorilla/i18n/linebreak"
	_cc "github.com/unidoc/unichart/render"
	_fa "golang.org/x/text/unicode/bidi"
)

// Draw draws the drawable d on the block.
// Note that the drawable must not wrap, i.e. only return one block. Otherwise an error is returned.
func (_bgc *Block) Draw(d Drawable) error {
	_cecb := DrawContext{}
	_cecb.Width = _bgc._da
	_cecb.Height = _bgc._ace
	_cecb.PageWidth = _bgc._da
	_cecb.PageHeight = _bgc._ace
	_cecb.X = 0
	_cecb.Y = 0
	_fac, _, _adf := d.GeneratePageBlocks(_cecb)
	if _adf != nil {
		return _adf
	}
	if len(_fac) != 1 {
		return ErrContentNotFit
	}
	for _, _bbg := range _fac {
		if _bfa := _bgc.mergeBlocks(_bbg); _bfa != nil {
			return _bfa
		}
	}
	return nil
}

// GetIndent get the cell's left indent.
func (_eaea *TableCell) GetIndent() float64 { return _eaea._fceba }

const (
	FitModeNone FitMode = iota
	FitModeFillWidth
)

// Scale block by specified factors in the x and y directions.
func (_efa *Block) Scale(sx, sy float64) {
	_afb := _dg.NewContentCreator().Scale(sx, sy).Operations()
	*_efa._cb = append(*_afb, *_efa._cb...)
	_efa._cb.WrapIfNeeded()
	_efa._da *= sx
	_efa._ace *= sy
}

// NewParagraph creates a new text paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_aafc *Creator) NewParagraph(text string) *Paragraph { return _cdbe(text, _aafc.NewTextStyle()) }
func (_gaef *templateProcessor) parseTable(_debga *templateNode) (interface{}, error) {
	var _bccgd int64
	for _, _ddedf := range _debga._adfdg.Attr {
		_eeaba := _ddedf.Value
		switch _facef := _ddedf.Name.Local; _facef {
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			_bccgd = _gaef.parseInt64Attr(_facef, _eeaba)
		}
	}
	if _bccgd <= 0 {
		_gaef.nodeLogDebug(_debga, "\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006eu\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0074\u0061\u0062\u006ce\u0020\u0063\u006f\u006cu\u006d\u006e\u0073\u003a\u0020\u0025\u0064\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0031\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020m\u0061\u0079\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _bccgd)
		_bccgd = 1
	}
	_aacc := _gaef.creator.NewTable(int(_bccgd))
	for _, _ggag := range _debga._adfdg.Attr {
		_eaddd := _ggag.Value
		switch _egabg := _ggag.Name.Local; _egabg {
		case "\u0063\u006f\u006c\u0075\u006d\u006e\u002d\u0077\u0069\u0064\u0074\u0068\u0073":
			_aacc.SetColumnWidths(_gaef.parseFloatArray(_egabg, _eaddd)...)
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fbcdf := _gaef.parseMarginAttr(_egabg, _eaddd)
			_aacc.SetMargins(_fbcdf.Left, _fbcdf.Right, _fbcdf.Top, _fbcdf.Bottom)
		case "\u0078":
			_aacc.SetPos(_gaef.parseFloatAttr(_egabg, _eaddd), _aacc._cbdba)
		case "\u0079":
			_aacc.SetPos(_aacc._cbgfe, _gaef.parseFloatAttr(_egabg, _eaddd))
		case "\u0068\u0065a\u0064\u0065\u0072-\u0073\u0074\u0061\u0072\u0074\u002d\u0072\u006f\u0077":
			_aacc._fefac = int(_gaef.parseInt64Attr(_egabg, _eaddd))
		case "\u0068\u0065\u0061\u0064\u0065\u0072\u002d\u0065\u006ed\u002d\u0072\u006f\u0077":
			_aacc._gefbf = int(_gaef.parseInt64Attr(_egabg, _eaddd))
		case "\u0065n\u0061b\u006c\u0065\u002d\u0072\u006f\u0077\u002d\u0077\u0072\u0061\u0070":
			_aacc.EnableRowWrap(_gaef.parseBoolAttr(_egabg, _eaddd))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_aacc.EnablePageWrap(_gaef.parseBoolAttr(_egabg, _eaddd))
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			break
		default:
			_gaef.nodeLogDebug(_debga, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0074a\u0062\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _egabg)
		}
	}
	if _aacc._fefac != 0 && _aacc._gefbf != 0 {
		_gdcac := _aacc.SetHeaderRows(_aacc._fefac, _aacc._gefbf)
		if _gdcac != nil {
			_gaef.nodeLogDebug(_debga, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0073\u0065t t\u0061bl\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0072\u006f\u0077\u0073\u003a\u0020\u0025v\u002e", _gdcac)
		}
	} else {
		_aacc._fefac = 0
		_aacc._gefbf = 0
	}
	return _aacc, nil
}

// ToPdfShadingPattern generates a new model.PdfShadingPatternType3 object.
func (_fddb *RadialShading) ToPdfShadingPattern() *_fgd.PdfShadingPatternType3 {
	_abgfb, _fbce, _fbae := _fddb._dafd._feac.ToRGB()
	_gdbba := _fddb.shadingModel()
	_gdbba.PdfShading.Background = _bc.MakeArrayFromFloats([]float64{_abgfb, _fbce, _fbae})
	_dbaa := _fgd.NewPdfShadingPatternType3()
	_dbaa.Shading = _gdbba
	return _dbaa
}

// SetLineMargins sets the margins for all new lines of the table of contents.
func (_ffede *TOC) SetLineMargins(left, right, top, bottom float64) {
	_bedgea := &_ffede._becec
	_bedgea.Left = left
	_bedgea.Right = right
	_bedgea.Top = top
	_bedgea.Bottom = bottom
}

// ScaleToHeight scales the Block to a specified height, maintaining the same aspect ratio.
func (_ged *Block) ScaleToHeight(h float64) { _bfg := h / _ged._ace; _ged.Scale(_bfg, _bfg) }

// SetMarkedContentID sets the marked content id for the paragraph.
func (_befbd *Paragraph) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_befbd._dege = &mcid
	_ggdg := _fgd.NewKDictionary()
	_ggdg.S = _bc.MakeName("\u0050")
	_ggdg.K = _bc.MakeInteger(mcid)
	return _ggdg
}
func _cbcc(_bfcgg, _agfdd, _efeb int) []int {
	_fadgf := []int{}
	for _gfgc := _bfcgg; _gfgc <= _efeb; _gfgc += _agfdd {
		_fadgf = append(_fadgf, _gfgc)
	}
	return _fadgf
}
func (_aefce *Invoice) generateTotalBlocks(_gcgd DrawContext) ([]*Block, DrawContext, error) {
	_dgge := _bagfg(4)
	_dgge.SetMargins(0, 0, 10, 10)
	_bgad := [][2]*InvoiceCell{_aefce._beff}
	_bgad = append(_bgad, _aefce._aegd...)
	_bgad = append(_bgad, _aefce._gdba)
	for _, _gafe := range _bgad {
		_ggge, _ebgdg := _gafe[0], _gafe[1]
		if _ebgdg.Value == "" {
			continue
		}
		_dgge.SkipCells(2)
		_caed := _dgge.NewCell()
		_caed.SetBackgroundColor(_ggge.BackgroundColor)
		_caed.SetHorizontalAlignment(_ebgdg.Alignment)
		_aefce.setCellBorder(_caed, _ggge)
		_gcf := _ddge(_ggge.TextStyle)
		_gcf.SetMargins(0, 0, 2, 1)
		_gcf.Append(_ggge.Value)
		_caed.SetContent(_gcf)
		_caed = _dgge.NewCell()
		_caed.SetBackgroundColor(_ebgdg.BackgroundColor)
		_caed.SetHorizontalAlignment(_ebgdg.Alignment)
		_aefce.setCellBorder(_caed, _ggge)
		_gcf = _ddge(_ebgdg.TextStyle)
		_gcf.SetMargins(0, 0, 2, 1)
		_gcf.Append(_ebgdg.Value)
		_caed.SetContent(_gcf)
	}
	return _dgge.GeneratePageBlocks(_gcgd)
}

// SetFitMode sets the fit mode of the image.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_dbdef *Image) SetFitMode(fitMode FitMode) { _dbdef._cba = fitMode }
func _gadda(_dcde, _aaec interface{}) (interface{}, error) {
	_afeg, _beeffc := _gaegc(_dcde)
	if _beeffc != nil {
		return nil, _beeffc
	}
	switch _deeb := _afeg.(type) {
	case int64:
		_fbbcb, _daab := _gaegc(_aaec)
		if _daab != nil {
			return nil, _daab
		}
		switch _cfagg := _fbbcb.(type) {
		case int64:
			return _deeb + _cfagg, nil
		case float64:
			return float64(_deeb) + _cfagg, nil
		}
	case float64:
		_eagec, _ebcca := _gaegc(_aaec)
		if _ebcca != nil {
			return nil, _ebcca
		}
		switch _dbbe := _eagec.(type) {
		case int64:
			return _deeb + float64(_dbbe), nil
		case float64:
			return _deeb + _dbbe, nil
		}
	}
	return nil, _f.Errorf("\u0066\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0025\u0076\u0020\u0061\u006e\u0064\u0020\u0025\u0076", _dcde, _aaec)
}

// SetShowLinks sets visibility of links for the TOC lines.
func (_gffe *TOC) SetShowLinks(showLinks bool) { _gffe._ggabc = showLinks }

// SetColor sets the line color.
func (_deea *Curve) SetColor(col Color) { _deea._afgg = col }

// SkipCells skips over a specified number of cells in the table.
func (_fbffb *Table) SkipCells(num int) {
	if num < 0 {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _aaggb := 0; _aaggb < num; _aaggb++ {
		_fbffb.NewCell()
	}
}
func (_deeef *templateProcessor) parseFitModeAttr(_fgdf, _aabbe string) FitMode {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u0069\u0074\u0020\u006do\u0064\u0065\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _fgdf, _aabbe)
	_dbad := map[string]FitMode{"\u006e\u006f\u006e\u0065": FitModeNone, "\u0066\u0069\u006c\u006c\u002d\u0077\u0069\u0064\u0074\u0068": FitModeFillWidth}[_aabbe]
	return _dbad
}

// GetMargins returns the margins of the chart (left, right, top, bottom).
func (_bbbb *Chart) GetMargins() (float64, float64, float64, float64) {
	return _bbbb._bgfg.Left, _bbbb._bgfg.Right, _bbbb._bgfg.Top, _bbbb._bgfg.Bottom
}

// TableCell defines a table cell which can contain a Drawable as content.
type TableCell struct {
	_cbgd        Color
	_gaddf       _gga.LineStyle
	_faab        CellBorderStyle
	_ebgcd       Color
	_caef        float64
	_bggf        CellBorderStyle
	_aegdf       Color
	_eabab       float64
	_cfdbdc      CellBorderStyle
	_fcbbf       Color
	_adec        float64
	_bgcce       CellBorderStyle
	_gefce       Color
	_cddebf      float64
	_deef, _bacg int
	_afddb       int
	_ecddcb      int
	_aafg        VectorDrawable
	_bbbff       CellHorizontalAlignment
	_bdfgg       CellVerticalAlignment
	_fceba       float64
	_abdcb       *Table
}

// Width returns the Block's width.
func (_bce *Block) Width() float64 { return _bce._da }

// SetMakedContentID sets the marked content id for the table.
func (_egfca *Table) SetMarkedContentID(mcid int64) *_fgd.KDict { return nil }

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_bbbfbf *LinearShading) SetAntiAlias(enable bool) { _bbbfbf._dfae.SetAntiAlias(enable) }

// SetMargins sets the Block's left, right, top, bottom, margins.
func (_ffd *Block) SetMargins(left, right, top, bottom float64) {
	_ffd._aag.Left = left
	_ffd._aag.Right = right
	_ffd._aag.Top = top
	_ffd._aag.Bottom = bottom
}

// Angle returns the block rotation angle in degrees.
func (_dfg *Block) Angle() float64 { return _dfg._df }

// GeneratePageBlocks draws the filled curve on page blocks.
func (_dcfg *FilledCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bccc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_dfda, _, _agfb := _dcfg.draw(_bccc, "")
	if _agfb != nil {
		return nil, ctx, _agfb
	}
	_agfb = _bccc.addContentsByString(string(_dfda))
	if _agfb != nil {
		return nil, ctx, _agfb
	}
	return []*Block{_bccc}, ctx, nil
}
func _fdga(_fafb []*ColorPoint) *LinearShading {
	return &LinearShading{_dfae: &shading{_feac: ColorWhite, _ccea: false, _fece: []bool{false, false}, _bcaca: _fafb}, _gafag: &_fgd.PdfRectangle{}}
}

// SetLinePageStyle sets the style for the page part of all new lines
// of the table of contents.
func (_dccc *TOC) SetLinePageStyle(style TextStyle) { _dccc._caeab = style }

// GetMargins returns the margins of the TOC line: left, right, top, bottom.
func (_aeae *TOCLine) GetMargins() (float64, float64, float64, float64) {
	_beegg := &_aeae._afgdd._ccddg
	return _aeae._afage, _beegg.Right, _beegg.Top, _beegg.Bottom
}
func (_gbce *InvoiceAddress) fmtLine(_fcfa, _aefb string, _gbbcb bool) string {
	if _gbbcb {
		_aefb = ""
	}
	return _f.Sprintf("\u0025\u0073\u0025s\u000a", _aefb, _fcfa)
}

// Line defines a line between point 1 (X1, Y1) and point 2 (X2, Y2).
// The line width, color, style (solid or dashed) and opacity can be
// configured. Implements the Drawable interface.
type Line struct {
	_gbcgde float64
	_gcfa   float64
	_cfge   float64
	_ceef   float64
	_ebgca  Color
	_cedaf  _gga.LineStyle
	_eabf   float64
	_cgecc  []int64
	_dcebc  int64
	_aebf   float64
	_babe   Positioning
	_efdg   FitMode
	_fggd   Margins
	_fbeb   *int64
}

// NoteStyle returns the style properties used to render the content of the
// invoice note sections.
func (_afab *Invoice) NoteStyle() TextStyle { return _afab._dgad }

// SetMarkedContentID sets the marked content id for the chart.
func (_gbcbd *Chart) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_gbcbd._aed = &mcid
	_cdc := _fgd.NewKDictionary()
	_cdc.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_cdc.K = _bc.MakeInteger(mcid)
	return _cdc
}

// SetLogo sets the logo of the invoice.
func (_ecfg *Invoice) SetLogo(logo *Image) { _ecfg._gdfdf = logo }

// NewImageFromFile creates an Image from a file.
func (_ccfa *Creator) NewImageFromFile(path string) (*Image, error) { return _gbggc(path) }

// ScaleToWidth sets the graphic svg scaling factor with the given width.
func (_gceeb *GraphicSVG) ScaleToWidth(w float64) {
	_dbab := _gceeb._dgccb.Height / _gceeb._dgccb.Width
	_gceeb._dgccb.Width = w
	_gceeb._dgccb.Height = w * _dbab
	_gceeb._dgccb.SetScaling(_dbab, _dbab)
}

// Width is not used as the division component is designed to fill all the
// available space, depending on the context. Returns 0.
func (_fcdd *Division) Width() float64 { return 0 }
func (_adfg *FilledCurve) draw(_ageca *Block, _bdfc string) ([]byte, *_fgd.PdfRectangle, error) {
	_bedg := _gga.NewCubicBezierPath()
	for _, _bgbe := range _adfg._acce {
		_bedg = _bedg.AppendCurve(_bgbe)
	}
	creator := _dg.NewContentCreator()
	if _adfg._deca != nil {
		creator.Add_BDC(*_bc.MakeName(_fgd.StructureTypeFigure), map[string]_bc.PdfObject{"\u004d\u0043\u0049\u0044": _bc.MakeInteger(*_adfg._deca)})
	}
	creator.Add_q()
	if _adfg.FillEnabled && _adfg._ggea != nil {
		_bbgca := _cfcee(_adfg._ggea)
		_afbf := _cacbe(_ageca, _bbgca, _adfg._ggea, func() Rectangle {
			_cbeg := _gga.NewCubicBezierPath()
			for _, _deafg := range _adfg._acce {
				_cbeg = _cbeg.AppendCurve(_deafg)
			}
			_ecdb := _cbeg.GetBoundingBox()
			if _adfg.BorderEnabled {
				_ecdb.Height += _adfg.BorderWidth
				_ecdb.Width += _adfg.BorderWidth
				_ecdb.X -= _adfg.BorderWidth / 2
				_ecdb.Y -= _adfg.BorderWidth / 2
			}
			return Rectangle{_defb: _ecdb.X, _cdgf: _ecdb.Y, _bcede: _ecdb.Width, _cggg: _ecdb.Height}
		})
		if _afbf != nil {
			return nil, nil, _afbf
		}
		creator.SetNonStrokingColor(_bbgca)
	}
	if _adfg.BorderEnabled {
		if _adfg._ccec != nil {
			creator.SetStrokingColor(_cfcee(_adfg._ccec))
		}
		creator.Add_w(_adfg.BorderWidth)
	}
	if len(_bdfc) > 1 {
		creator.Add_gs(_bc.PdfObjectName(_bdfc))
	}
	_gga.DrawBezierPathWithCreator(_bedg, creator)
	creator.Add_h()
	if _adfg.FillEnabled && _adfg.BorderEnabled {
		creator.Add_B()
	} else if _adfg.FillEnabled {
		creator.Add_f()
	} else if _adfg.BorderEnabled {
		creator.Add_S()
	}
	creator.Add_Q()
	if _adfg._deca != nil {
		creator.Add_EMC()
	}
	_ggda := _bedg.GetBoundingBox()
	if _adfg.BorderEnabled {
		_ggda.Height += _adfg.BorderWidth
		_ggda.Width += _adfg.BorderWidth
		_ggda.X -= _adfg.BorderWidth / 2
		_ggda.Y -= _adfg.BorderWidth / 2
	}
	_geag := &_fgd.PdfRectangle{}
	_geag.Llx = _ggda.X
	_geag.Lly = _ggda.Y
	_geag.Urx = _ggda.X + _ggda.Width
	_geag.Ury = _ggda.Y + _ggda.Height
	return creator.Bytes(), _geag, nil
}

// MoveY moves the drawing context to absolute position y.
func (_dffe *Creator) MoveY(y float64) { _dffe._eee.Y = y }
func (_eabcaf *templateProcessor) parseParagraph(_fbbg *templateNode, _eebd *Paragraph) (interface{}, error) {
	if _eebd == nil {
		_eebd = _eabcaf.creator.NewParagraph("")
	}
	for _, _efdda := range _fbbg._adfdg.Attr {
		_cebec := _efdda.Value
		switch _eebf := _efdda.Name.Local; _eebf {
		case "\u0066\u006f\u006e\u0074":
			_eebd.SetFont(_eabcaf.parseFontAttr(_eebf, _cebec))
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_eebd.SetFontSize(_eabcaf.parseFloatAttr(_eebf, _cebec))
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_eebd.SetTextAlignment(_eabcaf.parseTextAlignmentAttr(_eebf, _cebec))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_eebd.SetLineHeight(_eabcaf.parseFloatAttr(_eebf, _cebec))
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_eebd.SetEnableWrap(_eabcaf.parseBoolAttr(_eebf, _cebec))
		case "\u0063\u006f\u006co\u0072":
			_eebd.SetColor(_eabcaf.parseColorAttr(_eebf, _cebec))
		case "\u0078":
			_eebd.SetPos(_eabcaf.parseFloatAttr(_eebf, _cebec), _eebd._ggfb)
		case "\u0079":
			_eebd.SetPos(_eebd._gddg, _eabcaf.parseFloatAttr(_eebf, _cebec))
		case "\u0061\u006e\u0067l\u0065":
			_eebd.SetAngle(_eabcaf.parseFloatAttr(_eebf, _cebec))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_gecfe := _eabcaf.parseMarginAttr(_eebf, _cebec)
			_eebd.SetMargins(_gecfe.Left, _gecfe.Right, _gecfe.Top, _gecfe.Bottom)
		case "\u006da\u0078\u002d\u006c\u0069\u006e\u0065s":
			_eebd.SetMaxLines(int(_eabcaf.parseInt64Attr(_eebf, _cebec)))
		default:
			_eabcaf.nodeLogDebug(_fbbg, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020pa\u0072\u0061\u0067\u0072\u0061\u0070h\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _eebf)
		}
	}
	return _eebd, nil
}

// SetPageMargins sets the page margins: left, right, top, bottom.
// The default page margins are 10% of document width.
func (_fca *Creator) SetPageMargins(left, right, top, bottom float64) {
	_fca._cbb.Left = left
	_fca._cbb.Right = right
	_fca._cbb.Top = top
	_fca._cbb.Bottom = bottom
}
func _gbfdd(_bbce, _bdda TextStyle) *Invoice {
	_gccbd := &Invoice{_dgfaa: "\u0049N\u0056\u004f\u0049\u0043\u0045", _eabc: "\u002c\u0020", _cgce: _bbce, _gfdfe: _bdda}
	_gccbd._ddaa = &InvoiceAddress{Separator: _gccbd._eabc}
	_gccbd._eeeb = &InvoiceAddress{Heading: "\u0042i\u006c\u006c\u0020\u0074\u006f", Separator: _gccbd._eabc}
	_faaea := ColorRGBFrom8bit(245, 245, 245)
	_agaa := ColorRGBFrom8bit(155, 155, 155)
	_gccbd._adaa = _bdda
	_gccbd._adaa.Color = _agaa
	_gccbd._adaa.FontSize = 20
	_gccbd._efdd = _bbce
	_gccbd._geef = _bdda
	_gccbd._dgad = _bbce
	_gccbd._bgbeg = _bdda
	_gccbd._gca = _gccbd.NewCellProps()
	_gccbd._gca.BackgroundColor = _faaea
	_gccbd._gca.TextStyle = _bdda
	_gccbd._gbfg = _gccbd.NewCellProps()
	_gccbd._gbfg.TextStyle = _bdda
	_gccbd._gbfg.BackgroundColor = _faaea
	_gccbd._gbfg.BorderColor = _faaea
	_gccbd._baae = _gccbd.NewCellProps()
	_gccbd._baae.BorderColor = _faaea
	_gccbd._baae.BorderSides = []CellBorderSide{CellBorderSideBottom}
	_gccbd._baae.Alignment = CellHorizontalAlignmentRight
	_gccbd._afcea = _gccbd.NewCellProps()
	_gccbd._afcea.Alignment = CellHorizontalAlignmentRight
	_gccbd._bcab = [2]*InvoiceCell{_gccbd.newCell("\u0049\u006e\u0076\u006f\u0069\u0063\u0065\u0020\u006eu\u006d\u0062\u0065\u0072", _gccbd._gca), _gccbd.newCell("", _gccbd._gca)}
	_gccbd._cbfg = [2]*InvoiceCell{_gccbd.newCell("\u0044\u0061\u0074\u0065", _gccbd._gca), _gccbd.newCell("", _gccbd._gca)}
	_gccbd._dfge = [2]*InvoiceCell{_gccbd.newCell("\u0044\u0075\u0065\u0020\u0044\u0061\u0074\u0065", _gccbd._gca), _gccbd.newCell("", _gccbd._gca)}
	_gccbd._beff = [2]*InvoiceCell{_gccbd.newCell("\u0053\u0075\u0062\u0074\u006f\u0074\u0061\u006c", _gccbd._afcea), _gccbd.newCell("", _gccbd._afcea)}
	_gdbbb := _gccbd._afcea
	_gdbbb.TextStyle = _bdda
	_gdbbb.BackgroundColor = _faaea
	_gdbbb.BorderColor = _faaea
	_gccbd._gdba = [2]*InvoiceCell{_gccbd.newCell("\u0054\u006f\u0074a\u006c", _gdbbb), _gccbd.newCell("", _gdbbb)}
	_gccbd._bbgb = [2]string{"\u004e\u006f\u0074e\u0073", ""}
	_gccbd._affc = [2]string{"T\u0065r\u006d\u0073\u0020\u0061\u006e\u0064\u0020\u0063o\u006e\u0064\u0069\u0074io\u006e\u0073", ""}
	_gccbd._gffd = []*InvoiceCell{_gccbd.newColumn("D\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", CellHorizontalAlignmentLeft), _gccbd.newColumn("\u0051\u0075\u0061\u006e\u0074\u0069\u0074\u0079", CellHorizontalAlignmentRight), _gccbd.newColumn("\u0055\u006e\u0069\u0074\u0020\u0070\u0072\u0069\u0063\u0065", CellHorizontalAlignmentRight), _gccbd.newColumn("\u0041\u006d\u006f\u0075\u006e\u0074", CellHorizontalAlignmentRight)}
	return _gccbd
}

// PageFinalizeFunctionArgs holds the input arguments provided to the page
// finalize callback function which can be set using Creator.PageFinalize.
type PageFinalizeFunctionArgs struct {
	PageNum    int
	PageWidth  float64
	PageHeight float64
	TOCPages   int
	TotalPages int
}

func _bbdg(_bcfed *Creator, _gcedc string, _afabf []byte, _aecfa *TemplateOptions, _daecg componentRenderer) *templateProcessor {
	if _aecfa == nil {
		_aecfa = &TemplateOptions{}
	}
	_aecfa.init()
	if _daecg == nil {
		_daecg = _bcfed
	}
	return &templateProcessor{creator: _bcfed, _bcfcf: _afabf, _cbcec: _aecfa, _ccfe: _daecg, _edcgd: _gcedc}
}

// FooterFunctionArgs holds the input arguments to a footer drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FooterFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// SetMargins sets the Chapter margins: left, right, top, bottom.
// Typically not needed as the creator's page margins are used.
func (_dee *Chapter) SetMargins(left, right, top, bottom float64) {
	_dee._dbe.Left = left
	_dee._dbe.Right = right
	_dee._dbe.Top = top
	_dee._dbe.Bottom = bottom
}

// Scale scales the ellipse dimensions by the specified factors.
func (_gggf *Ellipse) Scale(xFactor, yFactor float64) {
	_gggf._eefg = xFactor * _gggf._eefg
	_gggf._beb = yFactor * _gggf._beb
}

// SetLink makes the line an internal link.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
func (_dbgff *TOCLine) SetLink(page int64, x, y float64) {
	_dbgff._beed = x
	_dbgff._cdbff = y
	_dbgff._ebaab = page
	_ecfd := _dbgff._afgdd._fbegf.Color
	_dbgff.Number.Style.Color = _ecfd
	_dbgff.Title.Style.Color = _ecfd
	_dbgff.Separator.Style.Color = _ecfd
	_dbgff.Page.Style.Color = _ecfd
}

// GeneratePageBlocks generates the page blocks for the Division component.
// Multiple blocks are generated if the contents wrap over multiple pages.
func (_abfd *Division) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_bdcc []*Block
		_daea bool
		_bdg  error
		_dfb  = _abfd._fbec.IsRelative()
		_cgca = _abfd._ddfg.Top
	)
	if _dfb && !_abfd._defa && !_abfd._fdb {
		_edcd := _abfd.ctxHeight(ctx.Width)
		if _edcd > ctx.Height-_abfd._ddfg.Top && _edcd <= ctx.PageHeight-ctx.Margins.Top-ctx.Margins.Bottom {
			if _bdcc, ctx, _bdg = _gbcfe().GeneratePageBlocks(ctx); _bdg != nil {
				return nil, ctx, _bdg
			}
			_daea = true
			_cgca = 0
		}
	}
	_aggc := ctx
	_daee := ctx
	if _dfb {
		ctx.X += _abfd._ddfg.Left
		ctx.Y += _cgca
		ctx.Width -= _abfd._ddfg.Left + _abfd._ddfg.Right
		ctx.Height -= _cgca
		_daee = ctx
		ctx.X += _abfd._gbga.Left
		ctx.Y += _abfd._gbga.Top
		ctx.Width -= _abfd._gbga.Left + _abfd._gbga.Right
		ctx.Height -= _abfd._gbga.Top
		ctx.Margins.Top += _abfd._gbga.Top
		ctx.Margins.Bottom += _abfd._gbga.Bottom
		ctx.Margins.Left += _abfd._ddfg.Left + _abfd._gbga.Left
		ctx.Margins.Right += _abfd._ddfg.Right + _abfd._gbga.Right
	}
	ctx.Inline = _abfd._fdb
	_deaa := ctx
	_cabec := ctx
	var _agce float64
	for _, _ffca := range _abfd._gebd {
		if ctx.Inline {
			if (ctx.X-_deaa.X)+_ffca.Width() <= ctx.Width {
				ctx.Y = _cabec.Y
				ctx.Height = _cabec.Height
			} else {
				ctx.X = _deaa.X
				ctx.Width = _deaa.Width
				_cabec.Y += _agce
				_cabec.Height -= _agce
				_agce = 0
			}
		}
		_ddc, _gedb, _beec := _ffca.GeneratePageBlocks(ctx)
		if _beec != nil {
			_fec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006eg\u0020p\u0061\u0067\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073\u003a\u0020\u0025\u0076", _beec)
			return nil, ctx, _beec
		}
		if len(_ddc) < 1 {
			continue
		}
		if len(_bdcc) > 0 {
			_bdcc[len(_bdcc)-1].mergeBlocks(_ddc[0])
			_bdcc = append(_bdcc, _ddc[1:]...)
		} else {
			if _ebfe := _ddc[0]._cb; _ebfe == nil || len(*_ebfe) == 0 {
				_daea = true
			}
			_bdcc = append(_bdcc, _ddc[0:]...)
		}
		if ctx.Inline {
			if ctx.Page != _gedb.Page {
				_deaa.Y = ctx.Margins.Top
				_deaa.Height = ctx.PageHeight - ctx.Margins.Top
				_cabec.Y = _deaa.Y
				_cabec.Height = _deaa.Height
				_agce = _gedb.Height - _deaa.Height
			} else {
				if _bgbd := ctx.Height - _gedb.Height; _bgbd > _agce {
					_agce = _bgbd
				}
			}
		} else {
			_gedb.X = ctx.X
		}
		ctx = _gedb
	}
	if len(_abfd._gebd) == 0 {
		_bbff := NewBlock(ctx.Width, 0)
		_bdcc = append(_bdcc, _bbff)
	}
	ctx.Inline = _aggc.Inline
	ctx.Margins = _aggc.Margins
	if _dfb {
		ctx.X = _aggc.X
		ctx.Width = _aggc.Width
		ctx.Y += _abfd._gbga.Bottom
		ctx.Height -= _abfd._gbga.Bottom
	}
	if _abfd._agecb != nil {
		_bdcc, _bdg = _abfd.drawBackground(_bdcc, _daee, ctx, _daea)
		if _bdg != nil {
			return nil, ctx, _bdg
		}
	}
	if _abfd._fbec.IsAbsolute() {
		return _bdcc, _aggc, nil
	}
	ctx.Y += _abfd._ddfg.Bottom
	ctx.Height -= _abfd._ddfg.Bottom
	return _bdcc, ctx, nil
}

// SetPageLabels adds the specified page labels to the PDF file generated
// by the creator. See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
// NOTE: for existing PDF files, the page label ranges object can be obtained
// using the model.PDFReader's GetPageLabels method.
func (_abee *Creator) SetPageLabels(pageLabels _bc.PdfObject) { _abee._deg = pageLabels }

// Invoice represents a configurable invoice template.
type Invoice struct {
	_dgfaa string
	_gdfdf *Image
	_eeeb  *InvoiceAddress
	_ddaa  *InvoiceAddress
	_eabc  string
	_bcab  [2]*InvoiceCell
	_cbfg  [2]*InvoiceCell
	_dfge  [2]*InvoiceCell
	_afff  [][2]*InvoiceCell
	_gffd  []*InvoiceCell
	_agece [][]*InvoiceCell
	_beff  [2]*InvoiceCell
	_gdba  [2]*InvoiceCell
	_aegd  [][2]*InvoiceCell
	_bbgb  [2]string
	_affc  [2]string
	_fcfd  [][2]string
	_cgce  TextStyle
	_gfdfe TextStyle
	_adaa  TextStyle
	_efdd  TextStyle
	_geef  TextStyle
	_dgad  TextStyle
	_bgbeg TextStyle
	_gca   InvoiceCellProps
	_gbfg  InvoiceCellProps
	_baae  InvoiceCellProps
	_afcea InvoiceCellProps
	_fefb  Positioning
}

func (_aege *TableCell) width(_bbfc []float64, _gbfaa float64) float64 {
	_dfadf := float64(0.0)
	for _cggea := 0; _cggea < _aege._ecddcb; _cggea++ {
		_dfadf += _bbfc[_aege._bacg+_cggea-1]
	}
	return _dfadf * _gbfaa
}

const (
	TextRenderingModeFill TextRenderingMode = iota
	TextRenderingModeStroke
	TextRenderingModeFillStroke
	TextRenderingModeInvisible
	TextRenderingModeFillClip
	TextRenderingModeStrokeClip
	TextRenderingModeFillStrokeClip
	TextRenderingModeClip
)

// SetAngle would set the angle at which the gradient is rendered.
//
// The default angle would be 0 where the gradient would be rendered from left to right side.
func (_cfecdd *LinearShading) SetAngle(angle float64) { _cfecdd._cgeb = angle }

// EnablePageWrap controls whether the table is wrapped across pages.
// If disabled, the table is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the table is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
func (_ecage *Table) EnablePageWrap(enable bool)         { _ecage._bcagb = enable }
func (_fdd *Creator) setActivePage(_dgfgg *_fgd.PdfPage) { _fdd._eda = _dgfgg }

// BorderWidth returns the border width of the rectangle.
func (_cfgec *Rectangle) BorderWidth() float64 { return _cfgec._caeg }
func _bcgf(_gfbca string) (*GraphicSVG, error) {
	_ebga, _cbdc := _dd.ParseFromString(_gfbca)
	if _cbdc != nil {
		return nil, _cbdc
	}
	return _cdfg(_ebga)
}

type templateNode struct {
	_bedcd interface{}
	_adfdg _gf.StartElement
	_cefd  *templateNode
	_fcebb int
	_gefcf int
	_agdga int64
}

func _efbef(_agdae *_gf.Decoder) (int, int) { return 0, 0 }
func _fbcdc(_bedcc *templateProcessor, _abde *templateNode) (interface{}, error) {
	return _bedcc.parseDivision(_abde)
}

// MoveDown moves the drawing context down by relative displacement dy (negative goes up).
func (_bfgb *Creator) MoveDown(dy float64) { _bfgb._eee.Y += dy }

// Width returns the width of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_gfgg *Rectangle) Width() float64 { return _gfgg._bcede }

// SetFillColor sets the fill color.
func (_dcgb *Polygon) SetFillColor(color Color) {
	_dcgb._gcbdc = color
	_dcgb._acdc.FillColor = _cfcee(color)
}

// GeneratePageBlocks draws the composite curve polygon on a new block
// representing the page. Implements the Drawable interface.
func (_cfdb *CurvePolygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gdg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_bafb, _adcag := _gdg.setOpacity(_cfdb._fafc, _cfdb._bcefc)
	if _adcag != nil {
		return nil, ctx, _adcag
	}
	_ccd := _cfdb._dgc
	_ccd.FillEnabled = _ccd.FillColor != nil
	_ccd.BorderEnabled = _ccd.BorderColor != nil && _ccd.BorderWidth > 0
	var (
		_eggf = ctx.PageHeight
		_ffbg = _ccd.Rings
		_cbee = make([][]_gga.CubicBezierCurve, 0, len(_ccd.Rings))
	)
	_gdcg := _fgd.PdfRectangle{}
	if len(_ffbg) > 0 && len(_ffbg[0]) > 0 {
		_caac := _ffbg[0][0]
		_caac.P0.Y = _eggf - _caac.P0.Y
		_caac.P1.Y = _eggf - _caac.P1.Y
		_caac.P2.Y = _eggf - _caac.P2.Y
		_caac.P3.Y = _eggf - _caac.P3.Y
		_gdcg = _caac.GetBounds()
	}
	for _, _adee := range _ffbg {
		_cccg := make([]_gga.CubicBezierCurve, 0, len(_adee))
		for _, _egdb := range _adee {
			_acee := _egdb
			_acee.P0.Y = _eggf - _acee.P0.Y
			_acee.P1.Y = _eggf - _acee.P1.Y
			_acee.P2.Y = _eggf - _acee.P2.Y
			_acee.P3.Y = _eggf - _acee.P3.Y
			_cccg = append(_cccg, _acee)
			_dgdde := _acee.GetBounds()
			_gdcg.Llx = _fc.Min(_gdcg.Llx, _dgdde.Llx)
			_gdcg.Lly = _fc.Min(_gdcg.Lly, _dgdde.Lly)
			_gdcg.Urx = _fc.Max(_gdcg.Urx, _dgdde.Urx)
			_gdcg.Ury = _fc.Max(_gdcg.Ury, _dgdde.Ury)
		}
		_cbee = append(_cbee, _cccg)
	}
	_ccd.Rings = _cbee
	defer func() { _ccd.Rings = _ffbg }()
	if _ccd.FillEnabled {
		_gfba := _cacbe(_gdg, _cfdb._dgc.FillColor, _cfdb._eeea, func() Rectangle {
			return Rectangle{_defb: _gdcg.Llx, _cdgf: _gdcg.Lly, _bcede: _gdcg.Width(), _cggg: _gdcg.Height()}
		})
		if _gfba != nil {
			return nil, ctx, _gfba
		}
	}
	_ggbe, _, _adcag := _ccd.MarkedDraw(_bafb, _cfdb._becd)
	if _adcag != nil {
		return nil, ctx, _adcag
	}
	if _adcag = _gdg.addContentsByString(string(_ggbe)); _adcag != nil {
		return nil, ctx, _adcag
	}
	return []*Block{_gdg}, ctx, nil
}

// NewChart creates a new creator drawable based on the provided
// unichart chart component.
func NewChart(chart _cc.ChartRenderable) *Chart { return _aeee(chart) }

// SetMargins sets the Paragraph's margins.
func (_fadg *StyledParagraph) SetMargins(left, right, top, bottom float64) {
	_fadg._ccddg.Left = left
	_fadg._ccddg.Right = right
	_fadg._ccddg.Top = top
	_fadg._ccddg.Bottom = bottom
}

// GetHorizontalAlignment returns the horizontal alignment of the image.
func (_abfa *Image) GetHorizontalAlignment() HorizontalAlignment { return _abfa._fbbb }

// SetOpacity sets the opacity of the line (0-1).
func (_baaf *Line) SetOpacity(opacity float64) { _baaf._eabf = opacity }

// GeneratePageBlocks draws the ellipse on a new block representing the page.
func (_afge *Ellipse) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_adfe  []*Block
		_deagd = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_dde   = ctx
	)
	_efgc := _afge._fdbc.IsRelative()
	if _efgc {
		_afge.applyFitMode(ctx.Width)
		ctx.X += _afge._ddbd.Left
		ctx.Y += _afge._ddbd.Top
		ctx.Width -= _afge._ddbd.Left + _afge._ddbd.Right
		ctx.Height -= _afge._ddbd.Top + _afge._ddbd.Bottom
		if _afge._beb > ctx.Height {
			_adfe = append(_adfe, _deagd)
			_deagd = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_aaed := ctx
			_aaed.Y = ctx.Margins.Top + _afge._ddbd.Top
			_aaed.X = ctx.Margins.Left + _afge._ddbd.Left
			_aaed.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _afge._ddbd.Top - _afge._ddbd.Bottom
			_aaed.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _afge._ddbd.Left - _afge._ddbd.Right
			ctx = _aaed
		}
	} else {
		ctx.X = _afge._adbc - _afge._eefg/2
		ctx.Y = _afge._ddaee - _afge._beb/2
	}
	_effc := _gga.Circle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _afge._beb, Width: _afge._eefg, Height: _afge._beb, BorderWidth: _afge._gfbg, Opacity: 1.0}
	if _afge._baacc != nil {
		_effc.FillEnabled = true
		_bgacd := _cfcee(_afge._baacc)
		_dbcg := _cacbe(_deagd, _bgacd, _afge._baacc, func() Rectangle {
			return Rectangle{_defb: _effc.X, _cdgf: _effc.Y, _bcede: _effc.Width, _cggg: _effc.Height}
		})
		if _dbcg != nil {
			return nil, ctx, _dbcg
		}
		_effc.FillColor = _bgacd
	}
	if _afge._add != nil {
		_effc.BorderEnabled = false
		if _afge._gfbg > 0 {
			_effc.BorderEnabled = true
		}
		_effc.BorderColor = _cfcee(_afge._add)
		_effc.BorderWidth = _afge._gfbg
	}
	_aecb, _gceg := _deagd.setOpacity(_afge._fcda, _afge._fbcf)
	if _gceg != nil {
		return nil, ctx, _gceg
	}
	_deead, _, _gceg := _effc.MarkedDraw(_aecb, _afge._adde)
	if _gceg != nil {
		return nil, ctx, _gceg
	}
	_gceg = _deagd.addContentsByString(string(_deead))
	if _gceg != nil {
		return nil, ctx, _gceg
	}
	if _efgc {
		ctx.X = _dde.X
		ctx.Width = _dde.Width
		ctx.Y += _afge._beb + _afge._ddbd.Bottom
		ctx.Height -= _afge._beb
	} else {
		ctx = _dde
	}
	_adfe = append(_adfe, _deagd)
	return _adfe, ctx, nil
}

// SetDashPattern sets the dash pattern of the line.
// NOTE: the dash pattern is taken into account only if the style of the
// line is set to dashed.
func (_fecgg *Line) SetDashPattern(dashArray []int64, dashPhase int64) {
	_fecgg._cgecc = dashArray
	_fecgg._dcebc = dashPhase
}

// SetColPosition sets cell column position.
func (_bceda *TableCell) SetColPosition(col int) { _bceda._bacg = col }

// PolyBezierCurve represents a composite curve that is the result of joining
// multiple cubic Bezier curves.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type PolyBezierCurve struct {
	_ffcd  *_gga.PolyBezierCurve
	_ffafb float64
	_addf  float64
	_bbaef Color
	_daef  *int64
}

// SetLanguage sets the language identifier that will be stored inside document catalog.
func (_dba *Creator) SetLanguage(language string) { _dba._eeae = language }
func _eec(_agfd, _bgaa *_fgd.PdfPageResources) error {
	_eed, _ := _agfd.GetColorspaces()
	if _eed != nil && len(_eed.Colorspaces) > 0 {
		for _faef, _abbd := range _eed.Colorspaces {
			_eddb := *_bc.MakeName(_faef)
			if _bgaa.HasColorspaceByName(_eddb) {
				continue
			}
			_cgd := _bgaa.SetColorspaceByName(_eddb, _abbd)
			if _cgd != nil {
				return _cgd
			}
		}
	}
	return nil
}

// Height returns Image's document height.
func (_bfd *Image) Height() float64 { return _bfd._fegb }

const (
	TextAlignmentLeft TextAlignment = iota
	TextAlignmentRight
	TextAlignmentCenter
	TextAlignmentJustify
)

// SetNoteHeadingStyle sets the style properties used to render the heading
// of the invoice note sections.
func (_ggefe *Invoice) SetNoteHeadingStyle(style TextStyle) { _ggefe._bgbeg = style }

// FillOpacity returns the fill opacity of the rectangle (0-1).
func (_edagc *Rectangle) FillOpacity() float64 { return _edagc._afdae }
func (_dddda *Creator) initContext() {
	_dddda._eee.X = _dddda._cbb.Left
	_dddda._eee.Y = _dddda._cbb.Top
	_dddda._eee.Width = _dddda._gbgg - _dddda._cbb.Right - _dddda._cbb.Left
	_dddda._eee.Height = _dddda._agfdf - _dddda._cbb.Bottom - _dddda._cbb.Top
	_dddda._eee.PageHeight = _dddda._agfdf
	_dddda._eee.PageWidth = _dddda._gbgg
	_dddda._eee.Margins = _dddda._cbb
	_dddda._eee._cffd = _dddda.UnsupportedCharacterReplacement
}

// SetBorder sets the cell's border style.
func (_begd *TableCell) SetBorder(side CellBorderSide, style CellBorderStyle, width float64) {
	if style == CellBorderStyleSingle && side == CellBorderSideAll {
		_begd._faab = CellBorderStyleSingle
		_begd._caef = width
		_begd._bggf = CellBorderStyleSingle
		_begd._eabab = width
		_begd._cfdbdc = CellBorderStyleSingle
		_begd._adec = width
		_begd._bgcce = CellBorderStyleSingle
		_begd._cddebf = width
	} else if style == CellBorderStyleDouble && side == CellBorderSideAll {
		_begd._faab = CellBorderStyleDouble
		_begd._caef = width
		_begd._bggf = CellBorderStyleDouble
		_begd._eabab = width
		_begd._cfdbdc = CellBorderStyleDouble
		_begd._adec = width
		_begd._bgcce = CellBorderStyleDouble
		_begd._cddebf = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideLeft {
		_begd._faab = style
		_begd._caef = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideBottom {
		_begd._bggf = style
		_begd._eabab = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideRight {
		_begd._cfdbdc = style
		_begd._adec = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideTop {
		_begd._bgcce = style
		_begd._cddebf = width
	}
}

// SetWidthRight sets border width for right.
func (_bfc *border) SetWidthRight(bw float64) { _bfc._gebg = bw }

// BuyerAddress returns the buyer address used in the invoice template.
func (_gafa *Invoice) BuyerAddress() *InvoiceAddress { return _gafa._eeeb }

// SetBorderWidth sets the border width.
func (_gddb *CurvePolygon) SetBorderWidth(borderWidth float64) { _gddb._dgc.BorderWidth = borderWidth }

// GetMargins returns the margins of the graphic svg (left, right, top, bottom).
func (_agef *GraphicSVG) GetMargins() (float64, float64, float64, float64) {
	return _agef._aagce.Left, _agef._aagce.Right, _agef._aagce.Top, _agef._aagce.Bottom
}
func (_egcbbg *templateProcessor) parseMarginAttr(_dgbcf, _cegeg string) Margins {
	_fec.Log.Debug("\u0050\u0061r\u0073\u0069\u006e\u0067 \u006d\u0061r\u0067\u0069\u006e\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _dgbcf, _cegeg)
	_acdcc := Margins{}
	switch _defed := _eg.Fields(_cegeg); len(_defed) {
	case 1:
		_acdcc.Top, _ = _fg.ParseFloat(_defed[0], 64)
		_acdcc.Bottom = _acdcc.Top
		_acdcc.Left = _acdcc.Top
		_acdcc.Right = _acdcc.Top
	case 2:
		_acdcc.Top, _ = _fg.ParseFloat(_defed[0], 64)
		_acdcc.Bottom = _acdcc.Top
		_acdcc.Left, _ = _fg.ParseFloat(_defed[1], 64)
		_acdcc.Right = _acdcc.Left
	case 3:
		_acdcc.Top, _ = _fg.ParseFloat(_defed[0], 64)
		_acdcc.Left, _ = _fg.ParseFloat(_defed[1], 64)
		_acdcc.Right = _acdcc.Left
		_acdcc.Bottom, _ = _fg.ParseFloat(_defed[2], 64)
	case 4:
		_acdcc.Top, _ = _fg.ParseFloat(_defed[0], 64)
		_acdcc.Right, _ = _fg.ParseFloat(_defed[1], 64)
		_acdcc.Bottom, _ = _fg.ParseFloat(_defed[2], 64)
		_acdcc.Left, _ = _fg.ParseFloat(_defed[3], 64)
	}
	return _acdcc
}

// Terms returns the terms and conditions section of the invoice as a
// title-content pair.
func (_dadaf *Invoice) Terms() (string, string) { return _dadaf._affc[0], _dadaf._affc[1] }

// SetTextOverflow controls the behavior of paragraph text which
// does not fit in the available space.
func (_cced *StyledParagraph) SetTextOverflow(textOverflow TextOverflow) { _cced._ecedg = textOverflow }
func (_caacf *templateProcessor) parseTableCell(_fcea *templateNode) (interface{}, error) {
	if _fcea._cefd == nil {
		_caacf.nodeLogError(_fcea, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gdbeeg
	}
	_eegf, _afcd := _fcea._cefd._bedcd.(*Table)
	if !_afcd {
		_caacf.nodeLogError(_fcea, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020\u006e\u006f\u0074\u0020a\u0020\u0074a\u0062\u006c\u0065\u002e", _fcea._cefd._bedcd)
		return nil, _gdbeeg
	}
	var _ffbc, _dabef int64
	for _, _abgbb := range _fcea._adfdg.Attr {
		_gaed := _abgbb.Value
		switch _efdcd := _abgbb.Name.Local; _efdcd {
		case "\u0063o\u006c\u0073\u0070\u0061\u006e":
			_ffbc = _caacf.parseInt64Attr(_efdcd, _gaed)
		case "\u0072o\u0077\u0073\u0070\u0061\u006e":
			_dabef = _caacf.parseInt64Attr(_efdcd, _gaed)
		}
	}
	if _ffbc <= 0 {
		_ffbc = 1
	}
	if _dabef <= 0 {
		_dabef = 1
	}
	_cebde := _eegf.MultiCell(int(_dabef), int(_ffbc))
	for _, _adcca := range _fcea._adfdg.Attr {
		_cfae := _adcca.Value
		switch _bffab := _adcca.Name.Local; _bffab {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_cebde.SetIndent(_caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0061\u006c\u0069g\u006e":
			_cebde.SetHorizontalAlignment(_caacf.parseCellAlignmentAttr(_bffab, _cfae))
		case "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u002da\u006c\u0069\u0067\u006e":
			_cebde.SetVerticalAlignment(_caacf.parseCellVerticalAlignmentAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065":
			_cebde.SetSideBorderStyle(CellBorderSideAll, _caacf.parseCellBorderStyleAttr(_bffab, _cfae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0073\u0074\u0079\u006c\u0065\u002d\u0074\u006f\u0070":
			_cebde.SetSideBorderStyle(CellBorderSideTop, _caacf.parseCellBorderStyleAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0073\u0074\u0079\u006c\u0065\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cebde.SetSideBorderStyle(CellBorderSideBottom, _caacf.parseCellBorderStyleAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065-\u006c\u0065\u0066\u0074":
			_cebde.SetSideBorderStyle(CellBorderSideLeft, _caacf.parseCellBorderStyleAttr(_bffab, _cfae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0073\u0074y\u006c\u0065-\u0072\u0069\u0067\u0068\u0074":
			_cebde.SetSideBorderStyle(CellBorderSideRight, _caacf.parseCellBorderStyleAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_cebde.SetSideBorderWidth(CellBorderSideAll, _caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0077\u0069\u0064\u0074\u0068\u002d\u0074\u006f\u0070":
			_cebde.SetSideBorderWidth(CellBorderSideTop, _caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cebde.SetSideBorderWidth(CellBorderSideBottom, _caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068-\u006c\u0065\u0066\u0074":
			_cebde.SetSideBorderWidth(CellBorderSideLeft, _caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0077\u0069d\u0074\u0068-\u0072\u0069\u0067\u0068\u0074":
			_cebde.SetSideBorderWidth(CellBorderSideRight, _caacf.parseFloatAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_cebde.SetSideBorderColor(CellBorderSideAll, _caacf.parseColorAttr(_bffab, _cfae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0063\u006f\u006c\u006f\u0072\u002d\u0074\u006f\u0070":
			_cebde.SetSideBorderColor(CellBorderSideTop, _caacf.parseColorAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0063\u006f\u006c\u006f\u0072\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cebde.SetSideBorderColor(CellBorderSideBottom, _caacf.parseColorAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072-\u006c\u0065\u0066\u0074":
			_cebde.SetSideBorderColor(CellBorderSideLeft, _caacf.parseColorAttr(_bffab, _cfae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0063\u006fl\u006f\u0072-\u0072\u0069\u0067\u0068\u0074":
			_cebde.SetSideBorderColor(CellBorderSideRight, _caacf.parseColorAttr(_bffab, _cfae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006c\u0069\u006e\u0065\u002ds\u0074\u0079\u006c\u0065":
			_cebde.SetBorderLineStyle(_caacf.parseLineStyleAttr(_bffab, _cfae))
		case "\u0062\u0061c\u006b\u0067\u0072o\u0075\u006e\u0064\u002d\u0063\u006f\u006c\u006f\u0072":
			_cebde.SetBackgroundColor(_caacf.parseColorAttr(_bffab, _cfae))
		case "\u0063o\u006c\u0073\u0070\u0061\u006e", "\u0072o\u0077\u0073\u0070\u0061\u006e":
			break
		default:
			_caacf.nodeLogDebug(_fcea, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _bffab)
		}
	}
	return _cebde, nil
}

// Positioning returns the type of positioning the rectangle is set to use.
func (_bff *Rectangle) Positioning() Positioning { return _bff._dfff }

// CurCol returns the currently active cell's column number.
func (_gcac *Table) CurCol() int { _bdffd := (_gcac._ecgdc-1)%(_gcac._afacb) + 1; return _bdffd }

// SetEnableWrap sets the line wrapping enabled flag.
func (_gfcba *StyledParagraph) SetEnableWrap(enableWrap bool) {
	_gfcba._fegca = enableWrap
	_gfcba._fegde = false
}
func (_aceec *templateProcessor) processGradientColorPair(_befgb []string) (_agbf []Color, _edfb []float64) {
	for _, _bfgbb := range _befgb {
		var (
			_cebbe = _eg.Fields(_bfgbb)
			_bgfd  = len(_cebbe)
		)
		if _bgfd == 0 {
			continue
		}
		_eedge := ""
		if _bgfd > 1 {
			_eedge = _eg.TrimSpace(_cebbe[1])
		}
		_aefea := -1.0
		if _eg.HasSuffix(_eedge, "\u0025") {
			_ebac, _ddbcd := _fg.ParseFloat(_eedge[:len(_eedge)-1], 64)
			if _ddbcd != nil {
				_fec.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0069n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ddbcd)
			}
			_aefea = _ebac / 100.0
		}
		_bcedbb := _aceec.parseColor(_eg.TrimSpace(_cebbe[0]))
		if _bcedbb != nil {
			_agbf = append(_agbf, _bcedbb)
			_edfb = append(_edfb, _aefea)
		}
	}
	if len(_agbf) != len(_edfb) {
		_fec.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u006c\u0069\u006e\u0065\u0061\u0072\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0064\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0021")
		return nil, nil
	}
	_gdbde := -1
	_beged := 0.0
	for _aeba, _dbdbc := range _edfb {
		if _dbdbc == -1.0 {
			if _aeba == 0 {
				_dbdbc = 0.0
				_edfb[_aeba] = 0.0
				continue
			}
			_gdbde++
			if _aeba < len(_edfb)-1 {
				continue
			} else {
				_dbdbc = 1.0
				_edfb[_aeba] = 1.0
			}
		}
		_aeafg := _gdbde + 1
		for _aadaf := _aeba - _gdbde; _aadaf < _aeba; _aadaf++ {
			_edfb[_aadaf] = _beged + (float64(_aadaf) * (_dbdbc - _beged) / float64(_aeafg))
		}
		_beged = _dbdbc
		_gdbde = -1
	}
	return _agbf, _edfb
}

// Width returns the current page width.
func (_gdbe *Creator) Width() float64 { return _gdbe._gbgg }
func (_eaae *Creator) newPage() *_fgd.PdfPage {
	_gade := _fgd.NewPdfPage()
	_ccbdc := _eaae._ccae[0]
	_dfaa := _eaae._ccae[1]
	_feed := _fgd.PdfRectangle{Llx: 0, Lly: 0, Urx: _ccbdc, Ury: _dfaa}
	_gade.MediaBox = &_feed
	_eaae._gbgg = _ccbdc
	_eaae._agfdf = _dfaa
	_eaae.initContext()
	return _gade
}

// NewList creates a new list.
func (_gda *Creator) NewList() *List { return _cfbf(_gda.NewTextStyle()) }

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_bbeb *LinearShading) SetBackgroundColor(backgroundColor Color) {
	_bbeb._dfae.SetBackgroundColor(backgroundColor)
}

// SetBorderOpacity sets the border opacity.
func (_ddae *CurvePolygon) SetBorderOpacity(opacity float64) { _ddae._bcefc = opacity }

// TOC represents a table of contents component.
// It consists of a paragraph heading and a collection of
// table of contents lines.
// The representation of a table of contents line is as follows:
//
//	[number] [title]      [separator] [page]
//
// e.g.: Chapter1 Introduction ........... 1
type TOC struct {
	_defae  *StyledParagraph
	_ccfec  []*TOCLine
	_cggbgd TextStyle
	_fefdb  TextStyle
	_fadge  TextStyle
	_caeab  TextStyle
	_begag  string
	_eaadf  float64
	_becec  Margins
	_gabag  Positioning
	_eddd   TextStyle
	_ggabc  bool
}

func _fegbd(_bbgg map[string]interface{}, _ecebba ...interface{}) (map[string]interface{}, error) {
	_dedfc := len(_ecebba)
	if _dedfc%2 != 0 {
		return nil, _bc.ErrRangeError
	}
	for _egggc := 0; _egggc < _dedfc; _egggc += 2 {
		_fefd, _ccbg := _ecebba[_egggc].(string)
		if !_ccbg {
			return nil, _bc.ErrTypeError
		}
		_bbgg[_fefd] = _ecebba[_egggc+1]
	}
	return _bbgg, nil
}

// Finalize renders all blocks to the creator pages. In addition, it takes care
// of adding headers and footers, as well as generating the front page,
// table of contents and outlines.
// Finalize is automatically called before writing the document out. Calling the
// method manually can be useful when adding external pages to the creator,
// using the AddPage method, as it renders all creator blocks to the added
// pages, without having to write the document out.
// NOTE: TOC and outlines are generated only if the AddTOC and AddOutlines
// fields of the creator are set to true (enabled by default). Furthermore, TOCs
// and outlines without content are skipped. TOC and outline content is
// added automatically when using the chapter component. TOCs and outlines can
// also be set externally, using the SetTOC and SetOutlineTree methods.
// Finalize should only be called once, after all draw calls have taken place,
// as it will return immediately if the creator instance has been finalized.
func (_fgfg *Creator) Finalize() error {
	if _fgfg._ddfc {
		return nil
	}
	_dbb := len(_fgfg._egfb)
	_bcd := 0
	if _fgfg._cee != nil {
		_dbce := *_fgfg
		_fgfg._egfb = nil
		_fgfg._eda = nil
		_fgfg.initContext()
		_gcbc := FrontpageFunctionArgs{PageNum: 1, TotalPages: _dbb}
		_fgfg._cee(_gcbc)
		_bcd += len(_fgfg._egfb)
		_fgfg._egfb = _dbce._egfb
		_fgfg._eda = _dbce._eda
	}
	if _fgfg.AddTOC {
		_fgfg.initContext()
		_fgfg._eee.Page = _bcd + 1
		if _fgfg.CustomTOC && _fgfg._ggd != nil {
			_dcff := *_fgfg
			_fgfg._egfb = nil
			_fgfg._eda = nil
			if _gcdd := _fgfg._ggd(_fgfg._aac); _gcdd != nil {
				return _gcdd
			}
			_bcd += len(_fgfg._egfb)
			_fgfg._egfb = _dcff._egfb
			_fgfg._eda = _dcff._eda
		} else {
			if _fgfg._ggd != nil {
				if _aafe := _fgfg._ggd(_fgfg._aac); _aafe != nil {
					return _aafe
				}
			}
			_eade, _, _cfa := _fgfg._aac.GeneratePageBlocks(_fgfg._eee)
			if _cfa != nil {
				_fec.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073: \u0025\u0076", _cfa)
				return _cfa
			}
			_bcd += len(_eade)
		}
		_bcbe := _fgfg._aac.Lines()
		for _, _bdd := range _bcbe {
			_ccbec, _cebb := _fg.Atoi(_bdd.Page.Text)
			if _cebb != nil {
				continue
			}
			_bdd.Page.Text = _fg.Itoa(_ccbec + _bcd)
			_bdd._ebaab += int64(_bcd)
		}
	}
	_dcffa := false
	var _dafeb []*_fgd.PdfPage
	if _fgfg._cee != nil {
		_bgga := *_fgfg
		_fgfg._egfb = nil
		_fgfg._eda = nil
		_daeb := FrontpageFunctionArgs{PageNum: 1, TotalPages: _dbb}
		_fgfg._cee(_daeb)
		_dbb += len(_fgfg._egfb)
		_dafeb = _fgfg._egfb
		_fgfg._egfb = append(_fgfg._egfb, _bgga._egfb...)
		_fgfg._eda = _bgga._eda
		_dcffa = true
	}
	var _gec []*_fgd.PdfPage
	if _fgfg.AddTOC {
		_fgfg.initContext()
		if _fgfg.CustomTOC && _fgfg._ggd != nil {
			_gccc := *_fgfg
			_fgfg._egfb = nil
			_fgfg._eda = nil
			if _egcc := _fgfg._ggd(_fgfg._aac); _egcc != nil {
				_fec.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _egcc)
				return _egcc
			}
			_gec = _fgfg._egfb
			_dbb += len(_gec)
			_fgfg._egfb = _gccc._egfb
			_fgfg._eda = _gccc._eda
		} else {
			if _fgfg._ggd != nil {
				if _dgac := _fgfg._ggd(_fgfg._aac); _dgac != nil {
					_fec.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _dgac)
					return _dgac
				}
			}
			_cfb, _, _ := _fgfg._aac.GeneratePageBlocks(_fgfg._eee)
			for _, _fdee := range _cfb {
				_fdee.SetPos(0, 0)
				_dbb++
				_gceb := _fgfg.newPage()
				_gec = append(_gec, _gceb)
				_fgfg.setActivePage(_gceb)
				_fgfg.Draw(_fdee)
			}
		}
		if _dcffa {
			_cacg := _dafeb
			_eccf := _fgfg._egfb[len(_dafeb):]
			_fgfg._egfb = append([]*_fgd.PdfPage{}, _cacg...)
			_fgfg._egfb = append(_fgfg._egfb, _gec...)
			_fgfg._egfb = append(_fgfg._egfb, _eccf...)
		} else {
			_fgfg._egfb = append(_gec, _fgfg._egfb...)
		}
	}
	if _fgfg._cbba != nil && _fgfg.AddOutlines {
		var _bdcg func(_gcddf *_fgd.OutlineItem)
		_bdcg = func(_aecdd *_fgd.OutlineItem) {
			_aecdd.Dest.Page += int64(_bcd)
			if _fbb := int(_aecdd.Dest.Page); _fbb >= 0 && _fbb < len(_fgfg._egfb) {
				_aecdd.Dest.PageObj = _fgfg._egfb[_fbb].GetPageAsIndirectObject()
			} else {
				_fec.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _fbb)
			}
			_aecdd.Dest.Y = _fgfg._agfdf - _aecdd.Dest.Y
			_aefc := _aecdd.Items()
			for _, _caab := range _aefc {
				_bdcg(_caab)
			}
		}
		_aggb := _fgfg._cbba.Items()
		for _, _dcaa := range _aggb {
			_bdcg(_dcaa)
		}
		if _fgfg.AddTOC {
			var _eggb int
			if _dcffa {
				_eggb = len(_dafeb)
			}
			_dggc := _fgd.NewOutlineDest(int64(_eggb), 0, _fgfg._agfdf)
			if _eggb >= 0 && _eggb < len(_fgfg._egfb) {
				_dggc.PageObj = _fgfg._egfb[_eggb].GetPageAsIndirectObject()
			} else {
				_fec.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _eggb)
			}
			_fgfg._cbba.Insert(0, _fgd.NewOutlineItem("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073", _dggc))
		}
	}
	for _efag, _cgc := range _fgfg._egfb {
		_fgfg.setActivePage(_cgc)
		if _fgfg._bbad != nil {
			_fbg, _ccca, _dabd := _cgc.Size()
			if _dabd != nil {
				return _dabd
			}
			_gadc := PageFinalizeFunctionArgs{PageNum: _efag + 1, PageWidth: _fbg, PageHeight: _ccca, TOCPages: len(_gec), TotalPages: _dbb}
			if _ggf := _fgfg._bbad(_gadc); _ggf != nil {
				_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0070\u0061\u0067\u0065\u0020\u0066\u0069\u006e\u0061\u006c\u0069\u007a\u0065 \u0063\u0061\u006c\u006c\u0062\u0061\u0063k\u003a\u0020\u0025\u0076", _ggf)
				return _ggf
			}
		}
		if _fgfg._cgf != nil {
			_abed := NewBlock(_fgfg._gbgg, _fgfg._cbb.Top)
			_bae := HeaderFunctionArgs{PageNum: _efag + 1, TotalPages: _dbb}
			_fgfg._cgf(_abed, _bae)
			_abed.SetPos(0, 0)
			if _aggg := _fgfg.Draw(_abed); _aggg != nil {
				_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0068e\u0061\u0064\u0065\u0072\u003a\u0020\u0025v", _aggg)
				return _aggg
			}
		}
		if _fgfg._aedb != nil {
			_feca := NewBlock(_fgfg._gbgg, _fgfg._cbb.Bottom)
			_bfae := FooterFunctionArgs{PageNum: _efag + 1, TotalPages: _dbb}
			_fgfg._aedb(_feca, _bfae)
			_feca.SetPos(0, _fgfg._agfdf-_feca._ace)
			if _fgcd := _fgfg.Draw(_feca); _fgcd != nil {
				_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0066o\u006f\u0074\u0065\u0072\u003a\u0020\u0025v", _fgcd)
				return _fgcd
			}
		}
		_acfb, _baf := _fgfg._agfdc[_cgc]
		if _eca, _baed := _fgfg._feab[_cgc]; _baed {
			if _baf {
				_acfb.transformBlock(_eca)
			}
			if _bggg := _eca.drawToPage(_cgc); _bggg != nil {
				_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020%\u0064\u0020\u0062\u006c\u006f\u0063\u006bs\u003a\u0020\u0025\u0076", _efag+1, _bggg)
				return _bggg
			}
		}
		if _baf {
			if _bedb := _acfb.transformPage(_cgc); _bedb != nil {
				_fec.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0074\u0072\u0061\u006e\u0073f\u006f\u0072\u006d\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0076", _bedb)
				return _bedb
			}
		}
	}
	_fgfg._ddfc = true
	return nil
}

// SetBorderWidth sets the border width of the rectangle.
func (_adge *Rectangle) SetBorderWidth(bw float64) { _adge._caeg = bw }
func (_cbad *listItem) ctxHeight(_efccg float64) float64 {
	var _eacf float64
	switch _bbcd := _cbad._cdeda.(type) {
	case *Paragraph:
		if _bbcd._ecae {
			_bbcd.SetWidth(_efccg - _bbcd._eebg.Horizontal())
		}
		_eacf = _bbcd.Height() + _bbcd._eebg.Vertical()
		_eacf += 0.5 * _bbcd._acdge * _bbcd._fbcg
	case *StyledParagraph:
		if _bbcd._fegca {
			_bbcd.SetWidth(_efccg - _bbcd._ccddg.Horizontal())
		}
		_eacf = _bbcd.Height() + _bbcd._ccddg.Vertical()
		_eacf += 0.5 * _bbcd.getTextHeight()
	case *List:
		_ddeb := _efccg - _cbad._edee.Width() - _bbcd._gacb.Horizontal() - _bbcd._aedbb
		_eacf = _bbcd.ctxHeight(_ddeb) + _bbcd._gacb.Vertical()
	case *Image:
		_eacf = _bbcd.Height() + _bbcd._cdbc.Vertical()
	case *Division:
		_gfde := _efccg - _cbad._edee.Width() - _bbcd._ddfg.Horizontal()
		_eacf = _bbcd.ctxHeight(_gfde) + _bbcd._ddfg.Vertical()
	case *Table:
		_bfdf := _efccg - _cbad._edee.Width() - _bbcd._gbcea.Horizontal()
		_bbcd.updateRowHeights(_bfdf)
		_eacf = _bbcd.Height() + _bbcd._gbcea.Vertical()
	default:
		_eacf = _cbad._cdeda.Height()
	}
	return _eacf
}
func (_gag rgbColor) ToRGB() (float64, float64, float64) { return _gag._ead, _gag._ddd, _gag._decg }

// SetInline sets the inline mode of the division.
func (_ebdd *Division) SetInline(inline bool) { _ebdd._fdb = inline }
func (_bcgfd *templateProcessor) run() error {
	_gdadg := _gf.NewDecoder(_c.NewReader(_bcgfd._bcfcf))
	var _afabfd *templateNode
	for {
		_dfebg, _ccccc := _gdadg.Token()
		if _ccccc != nil {
			if _ccccc == _gg.EOF {
				return nil
			}
			return _ccccc
		}
		if _dfebg == nil {
			break
		}
		_gbgaf, _aedae := _efbef(_gdadg)
		_ggegb := _gdadg.InputOffset()
		switch _ffba := _dfebg.(type) {
		case _gf.StartElement:
			_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006eg\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0074\u0061\u0067\u003a\u0020`\u0025\u0073\u0060\u002e", _ffba.Name.Local)
			_gccad, _dcag := _faba[_ffba.Name.Local]
			if !_dcag {
				if _bcgfd._edcgd == "" {
					if _gbgaf != 0 {
						_fec.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0064\u003a\u0025\u0064\u005d", _ffba.Name.Local, _gbgaf, _aedae)
					} else {
						_fec.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070i\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u002e\u0020\u005b%\u0064\u005d", _ffba.Name.Local, _ggegb)
					}
				} else {
					if _gbgaf != 0 {
						_fec.Log.Debug("\u0055\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0074e\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u003a\u0025d\u005d", _ffba.Name.Local, _bcgfd._edcgd, _gbgaf, _aedae)
					} else {
						_fec.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u005d", _ffba.Name.Local, _bcgfd._edcgd, _ggegb)
					}
				}
				continue
			}
			_afabfd = &templateNode{_adfdg: _ffba, _cefd: _afabfd, _fcebb: _gbgaf, _gefcf: _aedae, _agdga: _ggegb}
			if _beac := _gccad._cacf; _beac != nil {
				_afabfd._bedcd, _ccccc = _beac(_bcgfd, _afabfd)
				if _ccccc != nil {
					return _ccccc
				}
			}
		case _gf.EndElement:
			_fec.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0065\u006e\u0064\u0020\u0074\u0061\u0067\u003a\u0020\u0060\u0025\u0073\u0060\u002e", _ffba.Name.Local)
			if _afabfd != nil {
				if _afabfd._bedcd != nil {
					if _bbge := _bcgfd.renderNode(_afabfd); _bbge != nil {
						return _bbge
					}
				}
				_afabfd = _afabfd._cefd
			}
		case _gf.CharData:
			if _afabfd != nil && _afabfd._bedcd != nil {
				if _bdad := _bcgfd.addNodeText(_afabfd, string(_ffba)); _bdad != nil {
					return _bdad
				}
			}
		case _gf.Comment:
			_fec.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u003a\u0020\u0060\u0025\u0073\u0060\u002e", string(_ffba))
		}
	}
	return nil
}

// ColorGrayFromArithmetic creates a Color from a grayscale value (0-1).
// Example:
//
//	gray := ColorGrayFromArithmetic(0.7)
func ColorGrayFromArithmetic(g float64) Color { return grayColor{g} }

// NewCurvePolygon creates a new curve polygon.
func (_gbfea *Creator) NewCurvePolygon(rings [][]_gga.CubicBezierCurve) *CurvePolygon {
	return _eedd(rings)
}
func (_gefb *StyledParagraph) getTextLineWidth(_bbcbf []*TextChunk) float64 {
	var _adfc float64
	_eged := len(_bbcbf)
	for _bgfe, _fdaed := range _bbcbf {
		_dgebd := &_fdaed.Style
		_eede := len(_fdaed.Text)
		for _dcdgg, _dddfd := range _fdaed.Text {
			if _dddfd == '\u000A' {
				continue
			}
			_edfe, _egccc := _dgebd.Font.GetRuneMetrics(_dddfd)
			if !_egccc {
				_fec.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _dddfd)
				return -1
			}
			_adfc += _dgebd.FontSize * _edfe.Wx * _dgebd.horizontalScale()
			if _dddfd != ' ' && (_bgfe != _eged-1 || _dcdgg != _eede-1) {
				_adfc += _dgebd.CharSpacing * 1000.0
			}
		}
	}
	return _adfc
}

// SetMargins sets the margins of the rectangle.
// NOTE: rectangle margins are only applied if relative positioning is used.
func (_gacf *Rectangle) SetMargins(left, right, top, bottom float64) {
	_gacf._gagb.Left = left
	_gacf._gagb.Right = right
	_gacf._gagb.Top = top
	_gacf._gagb.Bottom = bottom
}
func _eedgb(_cgbb *templateProcessor, _eedbd *templateNode) (interface{}, error) {
	return _cgbb.parseList(_eedbd)
}

// SetLineLevelOffset sets the amount of space an indentation level occupies
// for all new lines of the table of contents.
func (_bafbe *TOC) SetLineLevelOffset(levelOffset float64) { _bafbe._eaadf = levelOffset }
func (_befed *Invoice) newCell(_fgff string, _aadbf InvoiceCellProps) *InvoiceCell {
	return &InvoiceCell{_aadbf, _fgff}
}

// SetMargins sets the margins of the line.
// NOTE: line margins are only applied if relative positioning is used.
func (_eaaa *Line) SetMargins(left, right, top, bottom float64) {
	_eaaa._fggd.Left = left
	_eaaa._fggd.Right = right
	_eaaa._fggd.Top = top
	_eaaa._fggd.Bottom = bottom
}

// List represents a list of items.
// The representation of a list item is as follows:
//
//	[marker] [content]
//
// e.g.:        • This is the content of the item.
// The supported components to add content to list items are:
// - Paragraph
// - StyledParagraph
// - List
type List struct {
	_ffegf []*listItem
	_gacb  Margins
	_edaa  TextChunk
	_aedbb float64
	_ebfa  bool
	_bdgca Positioning
	_fedb  TextStyle
}

// SetEnableWrap sets the line wrapping enabled flag.
func (_gdff *Paragraph) SetEnableWrap(enableWrap bool) {
	_gdff._ecae = enableWrap
	_gdff._dcefd = false
}

// Positioning returns the type of positioning the ellipse is set to use.
func (_cea *Ellipse) Positioning() Positioning { return _cea._fdbc }

type fontMetrics struct {
	_geede float64
	_dfeba float64
	_feeec float64
	_egeb  float64
}

// Margins represents page margins or margins around an element.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// SetFillOpacity sets the fill opacity.
func (_gfeae *PolyBezierCurve) SetFillOpacity(opacity float64) { _gfeae._ffafb = opacity }
func (_bcedb *pageTransformations) transformPage(_befb *_fgd.PdfPage) error {
	if _cabe := _bcedb.applyFlip(_befb); _cabe != nil {
		return _cabe
	}
	return nil
}

// Width returns Image's document width.
func (_gaag *Image) Width() float64 { return _gaag._ecfc }

// AddPage adds the specified page to the creator.
// NOTE: If the page has a Rotate flag, the creator will take care of
// transforming the contents to maintain the correct orientation.
func (_gcce *Creator) AddPage(page *_fgd.PdfPage) error {
	_eegcd, _acba := _gcce.wrapPageIfNeeded(page)
	if _acba != nil {
		return _acba
	}
	if _eegcd != nil {
		page = _eegcd
	}
	_afad, _acba := page.GetMediaBox()
	if _acba != nil {
		_fec.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0070\u0061g\u0065\u0020\u006d\u0065\u0064\u0069\u0061\u0062\u006f\u0078\u003a\u0020\u0025\u0076", _acba)
		return _acba
	}
	_afad.Normalize()
	_dbg, _ffgc := _afad.Llx, _afad.Lly
	_acg := _afad
	if _abbf := page.CropBox; _abbf != nil && *_abbf != *_afad {
		_abbf.Normalize()
		_dbg, _ffgc = _abbf.Llx, _abbf.Lly
		_acg = _abbf
	}
	_efad := _de.IdentityMatrix()
	_eega, _acba := page.GetRotate()
	if _acba != nil {
		_fec.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _acba.Error())
	}
	_gbba := _eega%360 != 0 && _eega%90 == 0
	if _gbba {
		_ecf := float64((360 + _eega%360) % 360)
		if _ecf == 90 {
			_efad = _efad.Translate(_acg.Width(), 0)
		} else if _ecf == 180 {
			_efad = _efad.Translate(_acg.Width(), _acg.Height())
		} else if _ecf == 270 {
			_efad = _efad.Translate(0, _acg.Height())
		}
		_efad = _efad.Mult(_de.RotationMatrix(_ecf * _fc.Pi / 180))
		_efad = _efad.Round(0.000001)
		_eeb := _gdbda(_acg, _efad)
		_acg = _eeb
		_acg.Normalize()
	}
	if _dbg != 0 || _ffgc != 0 {
		_efad = _de.TranslationMatrix(_dbg, _ffgc).Mult(_efad)
	}
	if !_efad.Identity() {
		_efad = _efad.Round(0.000001)
		_gcce._agfdc[page] = &pageTransformations{_acb: &_efad}
	}
	_gcce._gbgg = _acg.Width()
	_gcce._agfdf = _acg.Height()
	_gcce.initContext()
	_gcce._egfb = append(_gcce._egfb, page)
	_gcce._eee.Page++
	return nil
}

// SetStyleBottom sets border style for bottom side.
func (_bcb *border) SetStyleBottom(style CellBorderStyle) { _bcb._fagg = style }

// GeneratePageBlocks draw graphic svg into block.
func (_dgeg *GraphicSVG) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fbdg := ctx
	_gdde := _dgeg._ddcg.IsRelative()
	var _bgccd []*Block
	if _gdde {
		_cdfc := 1.0
		_dfed := _dgeg._aagce.Top
		if _dgeg._dgccb.Height > ctx.Height-_dgeg._aagce.Top {
			_bgccd = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _cdcfb error
			if _, ctx, _cdcfb = _gbcfe().GeneratePageBlocks(ctx); _cdcfb != nil {
				return nil, ctx, _cdcfb
			}
			_dfed = 0
		}
		ctx.X += _dgeg._aagce.Left + _cdfc
		ctx.Y += _dfed
		ctx.Width -= _dgeg._aagce.Left + _dgeg._aagce.Right + 2*_cdfc
		ctx.Height -= _dfed
	} else {
		ctx.X = _dgeg._gddbc
		ctx.Y = _dgeg._bgfb
	}
	_dag := _dg.NewContentCreator()
	_dag.Translate(0, ctx.PageHeight)
	_dag.Scale(1, -1)
	_dag.Translate(ctx.X, ctx.Y)
	_dgba := _dgeg._dgccb.Width / _dgeg._dgccb.ViewBox.W
	_fcc := _dgeg._dgccb.Height / _dgeg._dgccb.ViewBox.H
	_eafa := 0.0
	_fbff := 0.0
	if _gdde {
		_eafa = _dgeg._gddbc - (_dgeg._dgccb.ViewBox.X * _fc.Max(_dgba, _fcc))
		_fbff = _dgeg._bgfb - (_dgeg._dgccb.ViewBox.Y * _fc.Max(_dgba, _fcc))
	}
	_dbdb := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _dgeg._bdae != nil {
		_dag.Add_BDC(*_bc.MakeName(_fgd.StructureTypeFigure), map[string]_bc.PdfObject{"\u004d\u0043\u0049\u0044": _bc.MakeInteger(*_dgeg._bdae)})
	}
	_dgeg._dgccb.ToContentCreator(_dag, _dbdb._dgd, _dgba, _fcc, _eafa, _fbff)
	if _dgeg._bdae != nil {
		_dag.Add_EMC()
	}
	if _accc := _dbdb.addContentsByString(_dag.String()); _accc != nil {
		return nil, ctx, _accc
	}
	if _gdde {
		_ggee := _dgeg.Height() + _dgeg._aagce.Bottom
		ctx.Y += _ggee
		ctx.Height -= _ggee
	} else {
		ctx = _fbdg
	}
	_bgccd = append(_bgccd, _dbdb)
	return _bgccd, ctx, nil
}

// SetHeight sets the Image's document height to specified h.
func (_gcgc *Image) SetHeight(h float64) { _gcgc._fegb = h }

// MoveX moves the drawing context to absolute position x.
func (_cgba *Creator) MoveX(x float64) { _cgba._eee.X = x }
func (_ffce *Image) makeXObject() error {
	_bgbb, _decbc := _fgd.NewXObjectImageFromImageLazy(_ffce._cgdfb, nil, _ffce._fdag, _ffce._faeda)
	if _decbc != nil {
		_fec.Log.Error("\u0046\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061t\u0065 \u0078o\u0062j\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _decbc)
		return _decbc
	}
	_ffce._gacc = _bgbb
	return nil
}

// Height returns the height of the chart.
func (_dfac *Chart) Height() float64 { return float64(_dfac._gff.Height()) }

// ColorCMYKFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorCMYKFromArithmetic(1.0, 0.0, 1.0, 0.0)
func ColorCMYKFromArithmetic(c, m, y, k float64) Color {
	return cmykColor{_ffea: _fc.Max(_fc.Min(c, 1.0), 0.0), _cbg: _fc.Max(_fc.Min(m, 1.0), 0.0), _egb: _fc.Max(_fc.Min(y, 1.0), 0.0), _gbfd: _fc.Max(_fc.Min(k, 1.0), 0.0)}
}

// SetLineHeight sets the line height (1.0 default).
func (_eeagd *Paragraph) SetLineHeight(lineheight float64) { _eeagd._fbcg = lineheight }
func _dgbbf(_fecdg float64, _adbb float64) float64         { return _fc.Round(_fecdg/_adbb) * _adbb }

// SetMarkedContentID sets the marked content ID for the chapter.
func (_ebc *Chapter) SetMarkedContentID(id int64) *_fgd.KDict { return nil }

// NewGraphicSVGFromFile creates a graphic SVG from a file.
func NewGraphicSVGFromFile(path string) (*GraphicSVG, error) { return _bded(path) }

// Width returns the cell's width based on the input draw context.
func (_abdab *TableCell) Width(ctx DrawContext) float64 {
	_cfccd := float64(0.0)
	for _dgbebb := 0; _dgbebb < _abdab._ecddcb; _dgbebb++ {
		_cfccd += _abdab._abdcb._cddfg[_abdab._bacg+_dgbebb-1]
	}
	_agfc := ctx.Width * _cfccd
	return _agfc
}

// PageSize represents the page size as a 2 element array representing the width and height in PDF document units (points).
type PageSize [2]float64

func (_ebgac *templateProcessor) parseColorAttr(_fgfac, _ceada string) Color {
	_fec.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0063\u006f\u006c\u006fr a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _fgfac, _ceada)
	_ceada = _eg.TrimSpace(_ceada)
	if _eg.HasPrefix(_ceada, "\u006c\u0069n\u0065\u0061\u0072-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _eg.HasSuffix(_ceada, "\u0029") && len(_ceada) > 17 {
		return _ebgac.parseLinearGradientAttr(_ebgac.creator, _ceada)
	}
	if _eg.HasPrefix(_ceada, "\u0072\u0061d\u0069\u0061\u006c-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _eg.HasSuffix(_ceada, "\u0029") && len(_ceada) > 17 {
		return _ebgac.parseRadialGradientAttr(_ebgac.creator, _ceada)
	}
	if _cgad := _ebgac.parseColor(_ceada); _cgad != nil {
		return _cgad
	}
	return ColorBlack
}

// Scale sets the scale ratio with `X` factor and `Y` factor for the graphic svg.
func (_bbef *GraphicSVG) Scale(xFactor, yFactor float64) {
	_bbef._dgccb.Width = xFactor * _bbef._dgccb.Width
	_bbef._dgccb.Height = yFactor * _bbef._dgccb.Height
	_bbef._dgccb.SetScaling(xFactor, yFactor)
}
func (_gdea *Table) resetColumnWidths() {
	_gdea._cddfg = []float64{}
	_ecgg := float64(1.0) / float64(_gdea._afacb)
	for _feaf := 0; _feaf < _gdea._afacb; _feaf++ {
		_gdea._cddfg = append(_gdea._cddfg, _ecgg)
	}
}

// Height returns the height of the list.
func (_ebgfa *List) Height() float64 {
	var _acag float64
	for _, _bdcb := range _ebgfa._ffegf {
		_acag += _bdcb.ctxHeight(_ebgfa.Width())
	}
	return _acag
}

// NewCellProps returns the default properties of an invoice cell.
func (_beag *Invoice) NewCellProps() InvoiceCellProps {
	_dcgd := ColorRGBFrom8bit(255, 255, 255)
	return InvoiceCellProps{TextStyle: _beag._cgce, Alignment: CellHorizontalAlignmentLeft, BackgroundColor: _dcgd, BorderColor: _dcgd, BorderWidth: 1, BorderSides: []CellBorderSide{CellBorderSideAll}}
}

// SetMarkedContentID sets marked content ID.
func (_afbb *border) SetMarkedContentID(id int64) *_fgd.KDict { return nil }
func (_beg *Division) drawBackground(_ccdf []*Block, _adcc, _ecge DrawContext, _bfge bool) ([]*Block, error) {
	_aaeg := len(_ccdf)
	if _aaeg == 0 || _beg._agecb == nil {
		return _ccdf, nil
	}
	_daa := make([]*Block, 0, len(_ccdf))
	for _dddc, _gdgg := range _ccdf {
		var (
			_aacg = _beg._agecb.BorderRadiusTopLeft
			_bcda = _beg._agecb.BorderRadiusTopRight
			_dgdb = _beg._agecb.BorderRadiusBottomLeft
			_dedf = _beg._agecb.BorderRadiusBottomRight
		)
		_acbd := _adcc
		_acbd.Page += _dddc
		if _dddc == 0 {
			if _bfge {
				_daa = append(_daa, _gdgg)
				continue
			}
			if _aaeg == 1 {
				_acbd.Height = _ecge.Y - _adcc.Y
			}
		} else {
			_acbd.X = _acbd.Margins.Left + _beg._ddfg.Left
			_acbd.Y = _acbd.Margins.Top
			_acbd.Width = _acbd.PageWidth - _acbd.Margins.Left - _acbd.Margins.Right - _beg._ddfg.Left - _beg._ddfg.Right
			if _dddc == _aaeg-1 {
				_acbd.Height = _ecge.Y - _acbd.Margins.Top - _beg._ddfg.Top
			} else {
				_acbd.Height = _acbd.PageHeight - _acbd.Margins.Top - _acbd.Margins.Bottom
			}
			if !_bfge {
				_aacg = 0
				_bcda = 0
			}
		}
		if _aaeg > 1 && _dddc != _aaeg-1 {
			_dgdb = 0
			_dedf = 0
		}
		_fdbe := _aggd(_acbd.X, _acbd.Y, _acbd.Width, _acbd.Height)
		_fdbe.SetFillColor(_beg._agecb.FillColor)
		_fdbe.SetBorderColor(_beg._agecb.BorderColor)
		_fdbe.SetBorderWidth(_beg._agecb.BorderSize)
		_fdbe.SetBorderRadius(_aacg, _bcda, _dgdb, _dedf)
		_cfec, _, _deag := _fdbe.GeneratePageBlocks(_acbd)
		if _deag != nil {
			return nil, _deag
		}
		if len(_cfec) == 0 {
			continue
		}
		_ddcf := _cfec[0]
		if _deag = _ddcf.mergeBlocks(_gdgg); _deag != nil {
			return nil, _deag
		}
		_daa = append(_daa, _ddcf)
	}
	return _daa, nil
}

// SetText replaces all the text of the paragraph with the specified one.
func (_fcff *StyledParagraph) SetText(text string) *TextChunk {
	_fcff.Reset()
	return _fcff.Append(text)
}
func _ddge(_gagfb TextStyle) *StyledParagraph {
	return &StyledParagraph{_ecec: []*TextChunk{}, _fgbg: _gagfb, _fbegf: _ebea(_gagfb.Font), _fgef: 1.0, _abff: TextAlignmentLeft, _fegca: true, _fegde: true, _eadc: false, _ffab: 0, _bagff: 1, _cebbd: 1, _fceb: PositionRelative, _cccab: ""}
}
func (_ggad cmykColor) ToRGB() (float64, float64, float64) {
	_affe := _ggad._gbfd
	return 1 - (_ggad._ffea*(1-_affe) + _affe), 1 - (_ggad._cbg*(1-_affe) + _affe), 1 - (_ggad._egb*(1-_affe) + _affe)
}
func (_fdab *templateProcessor) parseCellAlignmentAttr(_dagfbf, _fffea string) CellHorizontalAlignment {
	_fec.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020c\u0065\u006c\u006c\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _dagfbf, _fffea)
	_fbcc := map[string]CellHorizontalAlignment{"\u006c\u0065\u0066\u0074": CellHorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": CellHorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": CellHorizontalAlignmentRight}[_fffea]
	return _fbcc
}

// GeneratePageBlocks implements drawable interface.
func (_ccbe *border) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_eef := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_edde := _ccbe._bad
	_feg := ctx.PageHeight - _ccbe._abda
	if _ccbe._aff != nil {
		_bac := _gga.Rectangle{Opacity: 1.0, X: _ccbe._bad, Y: ctx.PageHeight - _ccbe._abda - _ccbe._aeg, Height: _ccbe._aeg, Width: _ccbe._cbec}
		_bac.FillEnabled = true
		_gcc := _cfcee(_ccbe._aff)
		_bbc := _cacbe(_eef, _gcc, _ccbe._aff, func() Rectangle {
			return Rectangle{_defb: _bac.X, _cdgf: _bac.Y, _bcede: _bac.Width, _cggg: _bac.Height}
		})
		if _bbc != nil {
			return nil, ctx, _bbc
		}
		_bac.FillColor = _gcc
		_bac.BorderEnabled = false
		_fga, _, _bbc := _bac.Draw("")
		if _bbc != nil {
			return nil, ctx, _bbc
		}
		_bbc = _eef.addContentsByString(string(_fga))
		if _bbc != nil {
			return nil, ctx, _bbc
		}
	}
	_agcag := _ccbe._dedg
	_fdae := _ccbe._fcf
	_ggaa := _ccbe._aca
	_cgda := _ccbe._gebg
	_gea := _ccbe._dedg
	if _ccbe._cggf == CellBorderStyleDouble {
		_gea += 2 * _agcag
	}
	_abg := _ccbe._fcf
	if _ccbe._fagg == CellBorderStyleDouble {
		_abg += 2 * _fdae
	}
	_ebe := _ccbe._aca
	if _ccbe._ffbb == CellBorderStyleDouble {
		_ebe += 2 * _ggaa
	}
	_fege := _ccbe._gebg
	if _ccbe._cfc == CellBorderStyleDouble {
		_fege += 2 * _cgda
	}
	_cff := (_gea - _ebe) / 2
	_aea := (_gea - _fege) / 2
	_dabe := (_abg - _ebe) / 2
	_gcbd := (_abg - _fege) / 2
	if _ccbe._dedg != 0 {
		_aafb := _edde
		_dce := _feg
		if _ccbe._cggf == CellBorderStyleDouble {
			_dce -= _agcag
			_cde := _gga.BasicLine{LineColor: _cfcee(_ccbe._dfd), Opacity: 1.0, LineWidth: _ccbe._dedg, LineStyle: _ccbe.LineStyle, X1: _aafb - _gea/2 + _cff, Y1: _dce + 2*_agcag, X2: _aafb + _gea/2 - _aea + _ccbe._cbec, Y2: _dce + 2*_agcag}
			_dcef, _, _aefe := _cde.Draw("")
			if _aefe != nil {
				return nil, ctx, _aefe
			}
			_aefe = _eef.addContentsByString(string(_dcef))
			if _aefe != nil {
				return nil, ctx, _aefe
			}
		}
		_afa := _gga.BasicLine{LineWidth: _ccbe._dedg, Opacity: 1.0, LineColor: _cfcee(_ccbe._dfd), LineStyle: _ccbe.LineStyle, X1: _aafb - _gea/2 + _cff + (_ebe - _ccbe._aca), Y1: _dce, X2: _aafb + _gea/2 - _aea + _ccbe._cbec - (_fege - _ccbe._gebg), Y2: _dce}
		_dcefa, _, _ede := _afa.Draw("")
		if _ede != nil {
			return nil, ctx, _ede
		}
		_ede = _eef.addContentsByString(string(_dcefa))
		if _ede != nil {
			return nil, ctx, _ede
		}
	}
	if _ccbe._fcf != 0 {
		_ddg := _edde
		_dfeb := _feg - _ccbe._aeg
		if _ccbe._fagg == CellBorderStyleDouble {
			_dfeb += _fdae
			_dcfa := _gga.BasicLine{LineWidth: _ccbe._fcf, Opacity: 1.0, LineColor: _cfcee(_ccbe._dcf), LineStyle: _ccbe.LineStyle, X1: _ddg - _abg/2 + _dabe, Y1: _dfeb - 2*_fdae, X2: _ddg + _abg/2 - _gcbd + _ccbe._cbec, Y2: _dfeb - 2*_fdae}
			_edg, _, _gccb := _dcfa.Draw("")
			if _gccb != nil {
				return nil, ctx, _gccb
			}
			_gccb = _eef.addContentsByString(string(_edg))
			if _gccb != nil {
				return nil, ctx, _gccb
			}
		}
		_ffec := _gga.BasicLine{LineWidth: _ccbe._fcf, Opacity: 1.0, LineColor: _cfcee(_ccbe._dcf), LineStyle: _ccbe.LineStyle, X1: _ddg - _abg/2 + _dabe + (_ebe - _ccbe._aca), Y1: _dfeb, X2: _ddg + _abg/2 - _gcbd + _ccbe._cbec - (_fege - _ccbe._gebg), Y2: _dfeb}
		_fcfb, _, _cfg := _ffec.Draw("")
		if _cfg != nil {
			return nil, ctx, _cfg
		}
		_cfg = _eef.addContentsByString(string(_fcfb))
		if _cfg != nil {
			return nil, ctx, _cfg
		}
	}
	if _ccbe._aca != 0 {
		_cfce := _edde
		_aaef := _feg
		if _ccbe._ffbb == CellBorderStyleDouble {
			_cfce += _ggaa
			_dac := _gga.BasicLine{LineWidth: _ccbe._aca, Opacity: 1.0, LineColor: _cfcee(_ccbe._bba), LineStyle: _ccbe.LineStyle, X1: _cfce - 2*_ggaa, Y1: _aaef + _ebe/2 + _cff, X2: _cfce - 2*_ggaa, Y2: _aaef - _ebe/2 - _dabe - _ccbe._aeg}
			_bgb, _, _gaaf := _dac.Draw("")
			if _gaaf != nil {
				return nil, ctx, _gaaf
			}
			_gaaf = _eef.addContentsByString(string(_bgb))
			if _gaaf != nil {
				return nil, ctx, _gaaf
			}
		}
		_ddgf := _gga.BasicLine{LineWidth: _ccbe._aca, Opacity: 1.0, LineColor: _cfcee(_ccbe._bba), LineStyle: _ccbe.LineStyle, X1: _cfce, Y1: _aaef + _ebe/2 + _cff - (_gea - _ccbe._dedg), X2: _cfce, Y2: _aaef - _ebe/2 - _dabe - _ccbe._aeg + (_abg - _ccbe._fcf)}
		_fffc, _, _caee := _ddgf.Draw("")
		if _caee != nil {
			return nil, ctx, _caee
		}
		_caee = _eef.addContentsByString(string(_fffc))
		if _caee != nil {
			return nil, ctx, _caee
		}
	}
	if _ccbe._gebg != 0 {
		_fba := _edde + _ccbe._cbec
		_dgb := _feg
		if _ccbe._cfc == CellBorderStyleDouble {
			_fba -= _cgda
			_cdac := _gga.BasicLine{LineWidth: _ccbe._gebg, Opacity: 1.0, LineColor: _cfcee(_ccbe._gdbb), LineStyle: _ccbe.LineStyle, X1: _fba + 2*_cgda, Y1: _dgb + _fege/2 + _aea, X2: _fba + 2*_cgda, Y2: _dgb - _fege/2 - _gcbd - _ccbe._aeg}
			_bcag, _, _ggg := _cdac.Draw("")
			if _ggg != nil {
				return nil, ctx, _ggg
			}
			_ggg = _eef.addContentsByString(string(_bcag))
			if _ggg != nil {
				return nil, ctx, _ggg
			}
		}
		_aaa := _gga.BasicLine{LineWidth: _ccbe._gebg, Opacity: 1.0, LineColor: _cfcee(_ccbe._gdbb), LineStyle: _ccbe.LineStyle, X1: _fba, Y1: _dgb + _fege/2 + _aea - (_gea - _ccbe._dedg), X2: _fba, Y2: _dgb - _fege/2 - _gcbd - _ccbe._aeg + (_abg - _ccbe._fcf)}
		_cfgd, _, _fggag := _aaa.Draw("")
		if _fggag != nil {
			return nil, ctx, _fggag
		}
		_fggag = _eef.addContentsByString(string(_cfgd))
		if _fggag != nil {
			return nil, ctx, _fggag
		}
	}
	return []*Block{_eef}, ctx, nil
}

// Creator is a wrapper around functionality for creating PDF reports and/or adding new
// content onto imported PDF pages, etc.
type Creator struct {

	// Errors keeps error messages that should not interrupt pdf processing and to be checked later.
	Errors []error

	// UnsupportedCharacterReplacement is character that will be used to replace unsupported glyph.
	// The value will be passed to drawing context.
	UnsupportedCharacterReplacement rune
	_egfb                           []*_fgd.PdfPage
	_feab                           map[*_fgd.PdfPage]*Block
	_agfdc                          map[*_fgd.PdfPage]*pageTransformations
	_eda                            *_fgd.PdfPage
	_ccae                           PageSize
	_eee                            DrawContext
	_cbb                            Margins
	_gbgg, _agfdf                   float64
	_dca                            int
	_cee                            func(_eegc FrontpageFunctionArgs)
	_ggd                            func(_beeg *TOC) error
	_cgf                            func(_abeb *Block, _fffd HeaderFunctionArgs)
	_aedb                           func(_dfcg *Block, _fad FooterFunctionArgs)
	_bbad                           func(_gad PageFinalizeFunctionArgs) error
	_gacd                           func(_ebf *_fgd.PdfWriter) error
	_ddfc                           bool

	// Controls whether a table of contents will be generated.
	AddTOC bool

	// CustomTOC specifies if the TOC is rendered by the user.
	// When the `CustomTOC` field is set to `true`, the default TOC component is not rendered.
	// Instead the TOC is drawn by the user, in the callback provided to
	// the `Creator.CreateTableOfContents` method.
	// If `CustomTOC` is set to `false`, the callback provided to
	// `Creator.CreateTableOfContents` customizes the style of the automatically generated TOC component.
	CustomTOC bool
	_aac      *TOC

	// Controls whether outlines will be generated.
	AddOutlines bool
	_cbba       *_fgd.Outline
	_egdf       *_fgd.PdfOutlineTreeNode
	_ecc        *_fgd.PdfAcroForm
	_deg        _bc.PdfObject
	_caec       _fgd.Optimizer
	_dad        []*_fgd.PdfFont
	_dff        *_fgd.PdfFont
	_ceb        *_fgd.PdfFont
	_cfgc       *_fgd.StructTreeRoot
	_bbcg       *_fgd.ViewerPreferences
	_eeae       string
}

func (_eegbg *templateProcessor) parseHorizontalAlignmentAttr(_adaaf, _ceeac string) HorizontalAlignment {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020a\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029.", _adaaf, _ceeac)
	_gbeef := map[string]HorizontalAlignment{"\u006c\u0065\u0066\u0074": HorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": HorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": HorizontalAlignmentRight}[_ceeac]
	return _gbeef
}

// Number returns the invoice number description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_gaab *Invoice) Number() (*InvoiceCell, *InvoiceCell) { return _gaab._bcab[0], _gaab._bcab[1] }

var _faba = map[string]*templateTag{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _cecbb}, "\u0074\u0065\u0078\u0074\u002d\u0063\u0068\u0075\u006e\u006b": &templateTag{_geca: map[string]struct{}{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": struct{}{}}, _cacf: _bgdfa}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _fbcdc}, "\u0074\u0061\u0062l\u0065": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _egeag}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": &templateTag{_geca: map[string]struct{}{"\u0074\u0061\u0062l\u0065": struct{}{}}, _cacf: _bfdcf}, "\u006c\u0069\u006e\u0065": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _ggadb}, "\u0072e\u0063\u0074\u0061\u006e\u0067\u006ce": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _bcfa}, "\u0065l\u006c\u0069\u0070\u0073\u0065": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _gbcbg}, "\u0069\u006d\u0061g\u0065": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _dcgcc}, "\u0063h\u0061\u0070\u0074\u0065\u0072": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _bbdbf}, "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067": &templateTag{_geca: map[string]struct{}{"\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _fgba}, "\u0063\u0068\u0061r\u0074": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _gfgfe}, "\u0070\u0061\u0067\u0065\u002d\u0062\u0072\u0065\u0061\u006b": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}}, _cacf: _gfcbf}, "\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064": &templateTag{_geca: map[string]struct{}{"\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}}, _cacf: _dbffa}, "\u006c\u0069\u0073\u0074": &templateTag{_geca: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": struct{}{}, "\u0062\u006c\u006fc\u006b": struct{}{}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": struct{}{}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": struct{}{}, "\u0063h\u0061\u0070\u0074\u0065\u0072": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _eedgb}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": &templateTag{_geca: map[string]struct{}{"\u006c\u0069\u0073\u0074": struct{}{}}, _cacf: _egce}, "l\u0069\u0073\u0074\u002d\u006d\u0061\u0072\u006b\u0065\u0072": &templateTag{_geca: map[string]struct{}{"\u006c\u0069\u0073\u0074": struct{}{}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": struct{}{}}, _cacf: _accfg}}

func _dgcef(_bfaad ...interface{}) (map[string]interface{}, error) {
	_bbgd := len(_bfaad)
	if _bbgd%2 != 0 {
		_fec.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u006d\u0061\u0070\u003a\u0020\u0025\u0064\u002e", _bbgd)
		return nil, _bc.ErrRangeError
	}
	_dgcec := map[string]interface{}{}
	for _bgfcg := 0; _bgfcg < _bbgd; _bgfcg += 2 {
		_aebgb, _fdcbd := _bfaad[_bgfcg].(string)
		if !_fdcbd {
			_fec.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006d\u0061\u0070 \u006b\u0065\u0079\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u002e\u0020\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u002e", _bfaad[_bgfcg])
			return nil, _bc.ErrTypeError
		}
		_dgcec[_aebgb] = _bfaad[_bgfcg+1]
	}
	return _dgcec, nil
}

// PageBreak represents a page break for a chapter.
type PageBreak struct{}

// NewLinearGradientColor creates a linear gradient color that could act as a color in other components.
func (_dgff *Creator) NewLinearGradientColor(colorPoints []*ColorPoint) *LinearShading {
	return _fdga(colorPoints)
}

// GetMargins returns the margins of the ellipse: left, right, top, bottom.
func (_fgac *Ellipse) GetMargins() (float64, float64, float64, float64) {
	return _fgac._ddbd.Left, _fgac._ddbd.Right, _fgac._ddbd.Top, _fgac._ddbd.Bottom
}

// The Image type is used to draw an image onto PDF.
type Image struct {
	_gacc         *_fgd.XObjectImage
	_cgdfb        *_fgd.Image
	_aabe         string
	_fbbf         float64
	_ecfc, _fegb  float64
	_bgcad, _bgff float64
	_accd         Positioning
	_fbbb         HorizontalAlignment
	_afggb        float64
	_dbccc        float64
	_ddea         float64
	_cdbc         Margins
	_bbdb, _cegba float64
	_fdag         _bc.StreamEncoder
	_cba          FitMode
	_faeda        bool
	_gebca        *int64
}

// Table allows organizing content in an rows X columns matrix, which can spawn across multiple pages.
type Table struct {
	_begg          int
	_afacb         int
	_ecgdc         int
	_cddfg         []float64
	_gecdg         []float64
	_edefd         float64
	_efbe          []*TableCell
	_aedca         []int
	_cgbfc         Positioning
	_cbgfe, _cbdba float64
	_gbcea         Margins
	_ddcb          bool
	_fefac         int
	_gefbf         int
	_fgaag         bool
	_bcagb         bool
	_eadda         bool
}

func (_fgdeg *Table) getLastCellFromCol(_ccdbf int) (int, *TableCell) {
	for _cdfb := len(_fgdeg._efbe) - 1; _cdfb >= 0; _cdfb-- {
		if _fgdeg._efbe[_cdfb]._bacg == _ccdbf {
			return _cdfb, _fgdeg._efbe[_cdfb]
		}
	}
	return 0, nil
}

// ColorPoint is a pair of Color and a relative point where the color
// would be rendered.
type ColorPoint struct {
	_afddd Color
	_ffag  float64
}

func _cbdce(_dcaeg int64, _bdagg, _dcafc, _gdceg float64) *_fgd.PdfAnnotation {
	_aadf := _fgd.NewPdfAnnotationLink()
	_cbgg := _fgd.NewBorderStyle()
	_cbgg.SetBorderWidth(0)
	_aadf.BS = _cbgg.ToPdfObject()
	if _dcaeg < 0 {
		_dcaeg = 0
	}
	_aadf.Dest = _bc.MakeArray(_bc.MakeInteger(_dcaeg), _bc.MakeName("\u0058\u0059\u005a"), _bc.MakeFloat(_bdagg), _bc.MakeFloat(_dcafc), _bc.MakeFloat(_gdceg))
	return _aadf.PdfAnnotation
}

// GeneratePageBlocks generates a page break block.
func (_degg *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dfecg := []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y), NewBlock(ctx.PageWidth, ctx.PageHeight)}
	ctx.Page++
	_abeeb := ctx
	_abeeb.Y = ctx.Margins.Top
	_abeeb.X = ctx.Margins.Left
	_abeeb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	_abeeb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = _abeeb
	return _dfecg, ctx, nil
}

// SetMargins sets the margins of the ellipse.
// NOTE: ellipse margins are only applied if relative positioning is used.
func (_bbcc *Ellipse) SetMargins(left, right, top, bottom float64) {
	_bbcc._ddbd.Left = left
	_bbcc._ddbd.Right = right
	_bbcc._ddbd.Top = top
	_bbcc._ddbd.Bottom = bottom
}
func _eedd(_eadg [][]_gga.CubicBezierCurve) *CurvePolygon {
	return &CurvePolygon{_dgc: &_gga.CurvePolygon{Rings: _eadg}, _fafc: 1.0, _bcefc: 1.0}
}

// ColorRGBFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorRGBFromArithmetic(0.0, 1.0, 0.0)
func ColorRGBFromArithmetic(r, g, b float64) Color {
	return rgbColor{_ead: _fc.Max(_fc.Min(r, 1.0), 0.0), _ddd: _fc.Max(_fc.Min(g, 1.0), 0.0), _decg: _fc.Max(_fc.Min(b, 1.0), 0.0)}
}

// NewPolyBezierCurve creates a new composite Bezier (polybezier) curve.
func (_ffbe *Creator) NewPolyBezierCurve(curves []_gga.CubicBezierCurve) *PolyBezierCurve {
	return _ccfcc(curves)
}

// IsAbsolute checks if the positioning is absolute.
func (_deda Positioning) IsAbsolute() bool { return _deda == PositionAbsolute }

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_bdaafg *RadialShading) SetBoundingBox(x, y, width, height float64) {
	_bdaafg._bfgf = &_fgd.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_bede *shading) SetAntiAlias(enable bool) { _bede._ccea = enable }

// SetFillColor sets the fill color for the path.
func (_aeag *FilledCurve) SetFillColor(color Color) { _aeag._ggea = color }

// GeneratePageBlocks generate the Page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages.
func (_fea *Chapter) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_aeb := ctx
	if _fea._bda.IsRelative() {
		ctx.X += _fea._dbe.Left
		ctx.Y += _fea._dbe.Top
		ctx.Width -= _fea._dbe.Left + _fea._dbe.Right
		ctx.Height -= _fea._dbe.Top
	}
	_dcd, _eeag, _afbd := _fea._gfdf.GeneratePageBlocks(ctx)
	if _afbd != nil {
		return _dcd, ctx, _afbd
	}
	ctx = _eeag
	_ebee := ctx.X
	_fegg := ctx.Y - _fea._gfdf.Height()
	_cfcc := int64(ctx.Page)
	_aecd := _fea.headingNumber()
	_cfeb := _fea.headingText()
	if _fea._aga {
		_cge := _fea._edc.Add(_aecd, _fea._faed, _fg.FormatInt(_cfcc, 10), _fea._eedg)
		if _fea._edc._ggabc {
			_cge.SetLink(_cfcc, _ebee, _fegg)
		}
	}
	if _fea._bcaf == nil {
		_fea._bcaf = _fgd.NewOutlineItem(_cfeb, _fgd.NewOutlineDest(_cfcc-1, _ebee, _fegg))
		if _fea._cdee != nil {
			_fea._cdee._bcaf.Add(_fea._bcaf)
		} else {
			_fea._agbcd.Add(_fea._bcaf)
		}
	} else {
		_cgeg := &_fea._bcaf.Dest
		_cgeg.Page = _cfcc - 1
		_cgeg.X = _ebee
		_cgeg.Y = _fegg
	}
	for _, _defc := range _fea._agbc {
		_bfe, _fef, _bgcc := _defc.GeneratePageBlocks(ctx)
		if _bgcc != nil {
			return _dcd, ctx, _bgcc
		}
		if len(_bfe) < 1 {
			continue
		}
		_dcd[len(_dcd)-1].mergeBlocks(_bfe[0])
		_dcd = append(_dcd, _bfe[1:]...)
		ctx = _fef
	}
	if _fea._bda.IsRelative() {
		ctx.X = _aeb.X
	}
	if _fea._bda.IsAbsolute() {
		return _dcd, _aeb, nil
	}
	return _dcd, ctx, nil
}
func (_gffga *templateProcessor) parseDivision(_bgbde *templateNode) (interface{}, error) {
	_cfff := _gffga.creator.NewDivision()
	for _, _fbgca := range _bgbde._adfdg.Attr {
		_eefd := _fbgca.Value
		switch _bdge := _fbgca.Name.Local; _bdge {
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_cfff.EnablePageWrap(_gffga.parseBoolAttr(_bdge, _eefd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_aadcf := _gffga.parseMarginAttr(_bdge, _eefd)
			_cfff.SetMargins(_aadcf.Left, _aadcf.Right, _aadcf.Top, _aadcf.Bottom)
		case "\u0070a\u0064\u0064\u0069\u006e\u0067":
			_eaga := _gffga.parseMarginAttr(_bdge, _eefd)
			_cfff.SetPadding(_eaga.Left, _eaga.Right, _eaga.Top, _eaga.Bottom)
		default:
			_gffga.nodeLogDebug(_bgbde, "U\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0064\u0069\u0076\u0069\u0073\u0069on\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025s`\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bdge)
		}
	}
	return _cfff, nil
}
func _bbbf(_gbda VectorDrawable, _eabe float64) float64 {
	switch _ccdd := _gbda.(type) {
	case *Paragraph:
		if _ccdd._ecae {
			_ccdd.SetWidth(_eabe - _ccdd._eebg.Left - _ccdd._eebg.Right)
		}
		return _ccdd.Height() + _ccdd._eebg.Top + _ccdd._eebg.Bottom
	case *StyledParagraph:
		if _ccdd._fegca {
			_ccdd.SetWidth(_eabe - _ccdd._ccddg.Left - _ccdd._ccddg.Right)
		}
		return _ccdd.Height() + _ccdd._ccddg.Top + _ccdd._ccddg.Bottom
	case *Image:
		_ccdd.applyFitMode(_eabe)
		return _ccdd.Height() + _ccdd._cdbc.Top + _ccdd._cdbc.Bottom
	case *Rectangle:
		_ccdd.applyFitMode(_eabe)
		return _ccdd.Height() + _ccdd._gagb.Top + _ccdd._gagb.Bottom + _ccdd._caeg
	case *Ellipse:
		_ccdd.applyFitMode(_eabe)
		return _ccdd.Height() + _ccdd._ddbd.Top + _ccdd._ddbd.Bottom
	case *Division:
		return _ccdd.ctxHeight(_eabe) + _ccdd._ddfg.Top + _ccdd._ddfg.Bottom + _ccdd._gbga.Top + _ccdd._gbga.Bottom
	case *Table:
		_ccdd.updateRowHeights(_eabe - _ccdd._gbcea.Left - _ccdd._gbcea.Right)
		return _ccdd.Height() + _ccdd._gbcea.Top + _ccdd._gbcea.Bottom
	case *List:
		return _ccdd.ctxHeight(_eabe) + _ccdd._gacb.Top + _ccdd._gacb.Bottom
	case marginDrawable:
		_, _, _bgca, _fffg := _ccdd.GetMargins()
		return _ccdd.Height() + _bgca + _fffg
	default:
		return _ccdd.Height()
	}
}

// SetAlternateText sets the alternate text for the image.
func (_cgeaf *Image) SetAlternateText(text string) { _cgeaf._aabe = text }

// Crop crops the Image to the specified bounds.
func (_acde *Image) Crop(x0, y0, x1, y1 int) {
	_cfeg, _acge := _acde._cgdfb.ToGoImage()
	if _acge != nil {
		_ac.Fatalf("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074o\u0020\u0047\u006f\u0020\u0049m\u0061\u0067e\u003a\u0020\u0025\u0076", _acge)
	}
	var _ebec _e.Image
	_cafd := _e.Rect(x0, y0, x1, y1)
	if _dbfae := _cafd.Intersect(_cfeg.Bounds()); !_cafd.Empty() {
		_beab := _e.NewRGBA(_e.Rect(0, 0, _cafd.Dx(), _cafd.Dy()))
		for _ffacc := _dbfae.Min.Y; _ffacc < _dbfae.Max.Y; _ffacc++ {
			for _cdgb := _dbfae.Min.X; _cdgb < _dbfae.Max.X; _cdgb++ {
				_beab.Set(_cdgb-_dbfae.Min.X, _ffacc-_dbfae.Min.Y, _cfeg.At(_cdgb, _ffacc))
			}
		}
		_ebec = _beab
	} else {
		_ebec = &_e.RGBA{}
	}
	_efee, _acge := _fgd.ImageHandling.NewImageFromGoImage(_ebec)
	if _acge != nil {
		_ac.Fatalf("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0066\u0072\u006fm\u0020\u0047\u006f\u0020\u0049m\u0061\u0067e\u003a\u0020\u0025\u0076", _acge)
	}
	_bcad := float64(_efee.Width)
	_gdac := float64(_efee.Height)
	_acde._cgdfb = _efee
	_acde._bgcad = _bcad
	_acde._bgff = _gdac
	_acde._ecfc = _bcad
	_acde._fegb = _gdac
}

// Draw processes the specified Drawable widget and generates blocks that can
// be rendered to the output document. The generated blocks can span over one
// or more pages. Additional pages are added if the contents go over the current
// page. Each generated block is assigned to the creator page it will be
// rendered to. In order to render the generated blocks to the creator pages,
// call Finalize, Write or WriteToFile.
func (_ebd *Creator) Draw(d Drawable) error {
	if _ebd.getActivePage() == nil {
		_ebd.NewPage()
	}
	_dccaf, _ccaf, _ggdb := d.GeneratePageBlocks(_ebd._eee)
	if _ggdb != nil {
		return _ggdb
	}
	if len(_ccaf._agd) > 0 {
		_ebd.Errors = append(_ebd.Errors, _ccaf._agd...)
	}
	for _ffgf, _abf := range _dccaf {
		if _ffgf > 0 {
			_ebd.NewPage()
		}
		_gbfe := _ebd.getActivePage()
		if _faa, _ddbcb := _ebd._feab[_gbfe]; _ddbcb {
			if _fbaa := _faa.mergeBlocks(_abf); _fbaa != nil {
				return _fbaa
			}
			if _fegd := _eec(_abf._dgd, _faa._dgd); _fegd != nil {
				return _fegd
			}
		} else {
			_ebd._feab[_gbfe] = _abf
		}
	}
	_ebd._eee.X = _ccaf.X
	_ebd._eee.Y = _ccaf.Y
	_ebd._eee.Height = _ccaf.PageHeight - _ccaf.Y - _ccaf.Margins.Bottom
	return nil
}

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_dbabd *Image) SetPos(x, y float64) {
	_dbabd._accd = PositionAbsolute
	_dbabd._afggb = x
	_dbabd._dbccc = y
}
func (_aafae *templateProcessor) parseEllipse(_edbg *templateNode) (interface{}, error) {
	_gdfgg := _aafae.creator.NewEllipse(0, 0, 0, 0)
	for _, _fcgc := range _edbg._adfdg.Attr {
		_egead := _fcgc.Value
		switch _bcdbb := _fcgc.Name.Local; _bcdbb {
		case "\u0063\u0078":
			_gdfgg._adbc = _aafae.parseFloatAttr(_bcdbb, _egead)
		case "\u0063\u0079":
			_gdfgg._ddaee = _aafae.parseFloatAttr(_bcdbb, _egead)
		case "\u0077\u0069\u0064t\u0068":
			_gdfgg.SetWidth(_aafae.parseFloatAttr(_bcdbb, _egead))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_gdfgg.SetHeight(_aafae.parseFloatAttr(_bcdbb, _egead))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_gdfgg.SetFillColor(_aafae.parseColorAttr(_bcdbb, _egead))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_gdfgg.SetFillOpacity(_aafae.parseFloatAttr(_bcdbb, _egead))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_gdfgg.SetBorderColor(_aafae.parseColorAttr(_bcdbb, _egead))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_gdfgg.SetBorderOpacity(_aafae.parseFloatAttr(_bcdbb, _egead))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_gdfgg.SetBorderWidth(_aafae.parseFloatAttr(_bcdbb, _egead))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_gdfgg.SetPositioning(_aafae.parsePositioningAttr(_bcdbb, _egead))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_gdfgg.SetFitMode(_aafae.parseFitModeAttr(_bcdbb, _egead))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_gafea := _aafae.parseMarginAttr(_bcdbb, _egead)
			_gdfgg.SetMargins(_gafea.Left, _gafea.Right, _gafea.Top, _gafea.Bottom)
		default:
			_aafae.nodeLogDebug(_edbg, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u006c\u0069\u0070\u0073\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bcdbb)
		}
	}
	return _gdfgg, nil
}

// ToPdfShadingPattern generates a new model.PdfShadingPatternType2 object.
func (_cgcaa *LinearShading) ToPdfShadingPattern() *_fgd.PdfShadingPatternType2 {
	_bccd, _bdfg, _gbdf := _cgcaa._dfae._feac.ToRGB()
	_fddc := _cgcaa.shadingModel()
	_fddc.PdfShading.Background = _bc.MakeArrayFromFloats([]float64{_bccd, _bdfg, _gbdf})
	_gbfcf := _fgd.NewPdfShadingPatternType2()
	_gbfcf.Shading = _fddc
	return _gbfcf
}

// SetSideBorderColor sets the cell's side border color.
func (_cgab *TableCell) SetSideBorderColor(side CellBorderSide, col Color) {
	switch side {
	case CellBorderSideAll:
		_cgab._gefce = col
		_cgab._aegdf = col
		_cgab._ebgcd = col
		_cgab._fcbbf = col
	case CellBorderSideTop:
		_cgab._gefce = col
	case CellBorderSideBottom:
		_cgab._aegdf = col
	case CellBorderSideLeft:
		_cgab._ebgcd = col
	case CellBorderSideRight:
		_cgab._fcbbf = col
	}
}
func _bggbd(_adfcb *_fgd.PdfAnnotation) *_fgd.PdfAnnotation {
	if _adfcb == nil {
		return nil
	}
	var _dbdd *_fgd.PdfAnnotation
	switch _fccg := _adfcb.GetContext().(type) {
	case *_fgd.PdfAnnotationLink:
		if _cabgc := _ddgga(_fccg); _cabgc != nil {
			_dbdd = _cabgc.PdfAnnotation
		}
	}
	return _dbdd
}

// SetLanguageIdentifier sets the language identifier for the paragraph.
func (_daec *Paragraph) SetLanguageIdentifier(id string) { _daec._eadae = id }
func (_eb *Block) translate(_ebb, _cfd float64) {
	_fag := _dg.NewContentCreator().Translate(_ebb, -_cfd).Operations()
	*_eb._cb = append(*_fag, *_eb._cb...)
	_eb._cb.WrapIfNeeded()
}

// Ellipse defines an ellipse with a center at (xc,yc) and a specified width and height.  The ellipse can have a colored
// fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Ellipse struct {
	_adbc  float64
	_ddaee float64
	_eefg  float64
	_beb   float64
	_fdbc  Positioning
	_baacc Color
	_fcda  float64
	_add   Color
	_gfbg  float64
	_fbcf  float64
	_ddbd  Margins
	_eddea FitMode
	_adde  *int64
}

func (_fabe *Invoice) drawSection(_ebag, _febcf string) []*StyledParagraph {
	var _ffcg []*StyledParagraph
	if _ebag != "" {
		_cfaa := _ddge(_fabe._bgbeg)
		_cfaa.SetMargins(0, 0, 0, 5)
		_cfaa.Append(_ebag)
		_ffcg = append(_ffcg, _cfaa)
	}
	if _febcf != "" {
		_ggbeee := _ddge(_fabe._dgad)
		_ggbeee.Append(_febcf)
		_ffcg = append(_ffcg, _ggbeee)
	}
	return _ffcg
}

// SetMarkedContentID sets marked content ID.
func (_adb *Curve) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_adb._afbbc = &mcid
	_cegb := _fgd.NewKDictionary()
	_cegb.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_cegb.K = _bc.MakeInteger(mcid)
	return _cegb
}

// BorderColor returns the border color of the rectangle.
func (_gagf *Rectangle) BorderColor() Color { return _gagf._bedf }
func _gbcfe() *PageBreak                    { return &PageBreak{} }

// SetMarkedContentID sets marked content ID.
func (_bbgc *CurvePolygon) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_bbgc._becd = &mcid
	_fdad := _fgd.NewKDictionary()
	_fdad.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_fdad.K = _bc.MakeInteger(mcid)
	return _fdad
}
func _bbdbf(_gaeae *templateProcessor, _gdcde *templateNode) (interface{}, error) {
	return _gaeae.parseChapter(_gdcde)
}

// BorderOpacity returns the border opacity of the ellipse (0-1).
func (_bacc *Ellipse) BorderOpacity() float64 { return _bacc._fbcf }

// SetLineHeight sets the line height (1.0 default).
func (_cefe *StyledParagraph) SetLineHeight(lineheight float64) { _cefe._fgef = lineheight }

// SetAngle sets the rotation angle of the text.
func (_gfcb *Paragraph) SetAngle(angle float64) { _gfcb._cadb = angle }
func (_gecg *templateProcessor) parseLineStyleAttr(_dcebd, _ceee string) _gga.LineStyle {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _dcebd, _ceee)
	_cgga := map[string]_gga.LineStyle{"\u0073\u006f\u006ci\u0064": _gga.LineStyleSolid, "\u0064\u0061\u0073\u0068\u0065\u0064": _gga.LineStyleDashed}[_ceee]
	return _cgga
}

// Lines returns all the rows of the invoice line items table.
func (_gddbd *Invoice) Lines() [][]*InvoiceCell { return _gddbd._agece }

// Drawable is a widget that can be used to draw with the Creator.
type Drawable interface {

	// GeneratePageBlocks draw onto blocks representing Page contents. As the content can wrap over many pages, multiple
	// templates are returned, one per Page.  The function also takes a draw context containing information
	// where to draw (if relative positioning) and the available height to draw on accounting for Margins etc.
	GeneratePageBlocks(_ecef DrawContext) ([]*Block, DrawContext, error)

	// SetMarkedContentID sets the marked content id for the drawable.
	SetMarkedContentID(_eeaef int64) *_fgd.KDict
}

func _afae(_bafe *Block, _caagc *Paragraph, _cfecd DrawContext) (DrawContext, error) {
	_cbed := 1
	_bdag := _bc.PdfObjectName("\u0046\u006f\u006e\u0074" + _fg.Itoa(_cbed))
	for _bafe._dgd.HasFontByName(_bdag) {
		_cbed++
		_bdag = _bc.PdfObjectName("\u0046\u006f\u006e\u0074" + _fg.Itoa(_cbed))
	}
	_bdac := _bafe._dgd.SetFontByName(_bdag, _caagc._gbca.ToPdfObject())
	if _bdac != nil {
		return _cfecd, _bdac
	}
	_caagc.wrapText()
	_cbdb := _dg.NewContentCreator()
	_cbdb.Add_q()
	_bbcb := _cfecd.PageHeight - _cfecd.Y - _caagc._acdge*_caagc._fbcg
	_cbdb.Translate(_cfecd.X, _bbcb)
	if _caagc._cadb != 0 {
		_cbdb.RotateDeg(_caagc._cadb)
	}
	_cgdb := _cfcee(_caagc._bbfd)
	_bdac = _cacbe(_bafe, _cgdb, _caagc._bbfd, func() Rectangle {
		return Rectangle{_defb: _cfecd.X, _cdgf: _bbcb, _bcede: _caagc.getMaxLineWidth() / 1000.0, _cggg: _caagc.Height()}
	})
	if _bdac != nil {
		return _cfecd, _bdac
	}
	_cbdb.Add_BT()
	_fdgbf := map[string]_bc.PdfObject{}
	if _caagc._dege != nil {
		_fdgbf["\u004d\u0043\u0049\u0044"] = _bc.MakeInteger(*_caagc._dege)
	}
	if _caagc._eadae != "" {
		_fdgbf["\u004c\u0061\u006e\u0067"] = _bc.MakeString(_caagc._eadae)
	}
	if len(_fdgbf) > 0 {
		_cbdb.Add_BDC(*_bc.MakeName(_fgd.StructureTypeParagraph), _fdgbf)
	}
	_cbdb.SetNonStrokingColor(_cgdb).Add_Tf(_bdag, _caagc._acdge).Add_TL(_caagc._acdge * _caagc._fbcg)
	for _bbaeb, _fgbe := range _caagc._ebdc {
		if _bbaeb != 0 {
			_cbdb.Add_Tstar()
		}
		_ecgeg := []rune(_fgbe)
		_fffdd := 0.0
		_bbfb := 0
		for _baaeg, _afdg := range _ecgeg {
			if _afdg == ' ' {
				_bbfb++
				continue
			}
			if _afdg == '\u000A' {
				continue
			}
			_bgfgc, _dfaab := _caagc._gbca.GetRuneMetrics(_afdg)
			if !_dfaab {
				_fec.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006e\u0065\u0020\u0069=\u0025\u0064\u0020\u0072\u0075\u006e\u0065=\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0069n\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0025\u0073", _baaeg, _afdg, _afdg, _caagc._gbca.BaseFont(), _caagc._gbca.Subtype())
				return _cfecd, _bd.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
			}
			_fffdd += _caagc._acdge * _bgfgc.Wx
		}
		var _edgc []_bc.PdfObject
		_fbge, _ddde := _caagc._gbca.GetRuneMetrics(' ')
		if !_ddde {
			return _cfecd, _bd.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
		}
		_edag := _fbge.Wx
		switch _caagc._cebg {
		case TextAlignmentJustify:
			if _bbfb > 0 && _bbaeb < len(_caagc._ebdc)-1 {
				_edag = (_caagc._gcbf*1000.0 - _fffdd) / float64(_bbfb) / _caagc._acdge
			}
		case TextAlignmentCenter:
			_cecef := _fffdd + float64(_bbfb)*_edag*_caagc._acdge
			_gbdg := (_caagc._gcbf*1000.0 - _cecef) / 2 / _caagc._acdge
			_edgc = append(_edgc, _bc.MakeFloat(-_gbdg))
		case TextAlignmentRight:
			_bfad := _fffdd + float64(_bbfb)*_edag*_caagc._acdge
			_eefge := (_caagc._gcbf*1000.0 - _bfad) / _caagc._acdge
			_edgc = append(_edgc, _bc.MakeFloat(-_eefge))
		}
		_dddcc := _caagc._gbca.Encoder()
		var _cbdd []byte
		for _, _dddf := range _ecgeg {
			if _dddf == '\u000A' {
				continue
			}
			if _dddf == ' ' {
				if len(_cbdd) > 0 {
					_edgc = append(_edgc, _bc.MakeStringFromBytes(_cbdd))
					_cbdd = nil
				}
				_edgc = append(_edgc, _bc.MakeFloat(-_edag))
			} else {
				if _, _ffbd := _dddcc.RuneToCharcode(_dddf); !_ffbd {
					_bdac = UnsupportedRuneError{Message: _f.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _dddf, _dddf), Rune: _dddf}
					_cfecd._agd = append(_cfecd._agd, _bdac)
					_fec.Log.Debug(_bdac.Error())
					if _cfecd._cffd <= 0 {
						continue
					}
					_dddf = _cfecd._cffd
				}
				_cbdd = append(_cbdd, _dddcc.Encode(string(_dddf))...)
			}
		}
		if len(_cbdd) > 0 {
			_edgc = append(_edgc, _bc.MakeStringFromBytes(_cbdd))
		}
		_cbdb.Add_TJ(_edgc...)
	}
	if len(_fdgbf) > 0 {
		_cbdb.Add_EMC()
	}
	_cbdb.Add_ET()
	_cbdb.Add_Q()
	_bgfc := _cbdb.Operations()
	_bgfc.WrapIfNeeded()
	_bafe.addContents(_bgfc)
	if _caagc._dbgab.IsRelative() {
		_egdd := _caagc.Height()
		_cfecd.Y += _egdd
		_cfecd.Height -= _egdd
		if _cfecd.Inline {
			_cfecd.X += _caagc.Width() + _caagc._eebg.Right
		}
	}
	return _cfecd, nil
}

// Context returns the current drawing context.
func (_edef *Creator) Context() DrawContext { return _edef._eee }

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_febgg *LinearShading) AddShadingResource(block *Block) (_abgc _bc.PdfObjectName, _dgagd error) {
	_daggf := 1
	_abgc = _bc.PdfObjectName("\u0053\u0068" + _fg.Itoa(_daggf))
	for block._dgd.HasShadingByName(_abgc) {
		_daggf++
		_abgc = _bc.PdfObjectName("\u0053\u0068" + _fg.Itoa(_daggf))
	}
	if _gbfa := block._dgd.SetShadingByName(_abgc, _febgg.shadingModel().ToPdfObject()); _gbfa != nil {
		return "", _gbfa
	}
	return _abgc, nil
}

const (
	AnchorBottomLeft AnchorPoint = iota
	AnchorBottomRight
	AnchorTopLeft
	AnchorTopRight
	AnchorCenter
	AnchorLeft
	AnchorRight
	AnchorTop
	AnchorBottom
)

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_ccab *Paragraph) SetTextAlignment(align TextAlignment) { _ccab._cebg = align }
func _edfc(_bcbeb []_gga.Point) *Polyline {
	return &Polyline{_cdggc: &_gga.Polyline{Points: _bcbeb, LineColor: _fgd.NewPdfColorDeviceRGB(0, 0, 0), LineWidth: 1.0}, _ggaf: 1.0}
}
func _cecbb(_afca *templateProcessor, _agfgc *templateNode) (interface{}, error) {
	return _afca.parseStyledParagraph(_agfgc)
}
func (_egegg *List) markerWidth() float64 {
	var _eaaae float64
	for _, _ceaa := range _egegg._ffegf {
		_bdcgd := _ddge(_egegg._fedb)
		_bdcgd.SetEnableWrap(false)
		_bdcgd.SetTextAlignment(TextAlignmentRight)
		_bdcgd.Append(_ceaa._edee.Text).Style = _ceaa._edee.Style
		_bfce := _bdcgd.getTextWidth() / 1000.0
		if _eaaae < _bfce {
			_eaaae = _bfce
		}
	}
	return _eaaae
}

// GetMargins returns the margins of the rectangle: left, right, top, bottom.
func (_dbgg *Rectangle) GetMargins() (float64, float64, float64, float64) {
	return _dbgg._gagb.Left, _dbgg._gagb.Right, _dbgg._gagb.Top, _dbgg._gagb.Bottom
}

// SetNoteStyle sets the style properties used to render the content of the
// invoice note sections.
func (_aedf *Invoice) SetNoteStyle(style TextStyle) { _aedf._dgad = style }
func (_cfaab *Invoice) generateNoteBlocks(_fffdg DrawContext) ([]*Block, DrawContext, error) {
	_feedg := _dcbc()
	_eedga := append([][2]string{_cfaab._bbgb, _cfaab._affc}, _cfaab._fcfd...)
	for _, _fbbfe := range _eedga {
		if _fbbfe[1] != "" {
			_abef := _cfaab.drawSection(_fbbfe[0], _fbbfe[1])
			for _, _bfdc := range _abef {
				_feedg.Add(_bfdc)
			}
			_aede := _ddge(_cfaab._cgce)
			_aede.SetMargins(0, 0, 10, 0)
			_feedg.Add(_aede)
		}
	}
	return _feedg.GeneratePageBlocks(_fffdg)
}

// CellHorizontalAlignment defines the table cell's horizontal alignment.
type CellHorizontalAlignment int
type listItem struct {
	_cdeda VectorDrawable
	_edee  TextChunk
}

// SetNumber sets the number of the invoice.
func (_fcde *Invoice) SetNumber(number string) (*InvoiceCell, *InvoiceCell) {
	_fcde._bcab[1].Value = number
	return _fcde._bcab[0], _fcde._bcab[1]
}

// Total returns the invoice total description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_eaca *Invoice) Total() (*InvoiceCell, *InvoiceCell) { return _eaca._gdba[0], _eaca._gdba[1] }
func (_egece *Paragraph) getTextWidth() float64 {
	_fffb := 0.0
	for _, _aaaf := range _egece._egea {
		if _aaaf == '\u000A' {
			continue
		}
		_ffcad, _eadgg := _egece._gbca.GetRuneMetrics(_aaaf)
		if !_eadgg {
			_fec.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _aaaf, _aaaf)
			return -1
		}
		_fffb += _egece._acdge * _ffcad.Wx
	}
	return _fffb
}

// CellBorderSide defines the table cell's border side.
type CellBorderSide int

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_dbcd *RadialShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }

// SetAddressHeadingStyle sets the style properties used to render the
// heading of the invoice address sections.
func (_cgecb *Invoice) SetAddressHeadingStyle(style TextStyle) { _cgecb._geef = style }

// GetMargins returns the Block's margins: left, right, top, bottom.
func (_deb *Block) GetMargins() (float64, float64, float64, float64) {
	return _deb._aag.Left, _deb._aag.Right, _deb._aag.Top, _deb._aag.Bottom
}

// Logo returns the logo of the invoice.
func (_dcfe *Invoice) Logo() *Image { return _dcfe._gdfdf }

// GeneratePageBlocks generates the page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages. Implements the Drawable interface.
func (_afag *StyledParagraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gcfd := ctx
	var _ddbee []*Block
	_cfbeg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _afag._fceb.IsRelative() {
		ctx.X += _afag._ccddg.Left
		ctx.Y += _afag._ccddg.Top
		ctx.Width -= _afag._ccddg.Left + _afag._ccddg.Right
		ctx.Height -= _afag._ccddg.Top
		_afag.SetWidth(ctx.Width)
	} else {
		if int(_afag._ffcfd) <= 0 {
			_afag.SetWidth(_afag.getTextWidth() / 1000.0)
		}
		ctx.X = _afag._abeag
		ctx.Y = _afag._agdg
	}
	if _afag._beccd != nil {
		_afag._beccd(_afag, ctx)
	}
	if _acbgd := _afag.wrapText(); _acbgd != nil {
		return nil, ctx, _acbgd
	}
	_adaaca := _afag._bdfce
	_dacb := 0
	for {
		_bbdec, _dcgbf, _bcdfg := _dcad(_cfbeg, _afag, _adaaca, ctx)
		if _bcdfg != nil {
			_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcdfg)
			return nil, ctx, _bcdfg
		}
		ctx = _bbdec
		_ddbee = append(_ddbee, _cfbeg)
		if _adaaca = _dcgbf; len(_dcgbf) == 0 {
			break
		}
		if len(_dcgbf) == _dacb {
			return nil, ctx, _bd.New("\u006e\u006f\u0074\u0020\u0065\u006e\u006f\u0075\u0067\u0068 \u0073\u0070\u0061\u0063\u0065\u0020\u0066o\u0072\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068")
		}
		_cfbeg = NewBlock(ctx.PageWidth, ctx.PageHeight)
		ctx.Page++
		_bbdec = ctx
		_bbdec.Y = ctx.Margins.Top
		_bbdec.X = ctx.Margins.Left + _afag._ccddg.Left
		_bbdec.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
		_bbdec.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _afag._ccddg.Left - _afag._ccddg.Right
		ctx = _bbdec
		_dacb = len(_dcgbf)
	}
	if _afag._fceb.IsRelative() {
		ctx.Y += _afag._ccddg.Bottom
		ctx.Height -= _afag._ccddg.Bottom
		if !ctx.Inline {
			ctx.X = _gcfd.X
			ctx.Width = _gcfd.Width
		}
		return _ddbee, ctx, nil
	}
	return _ddbee, _gcfd, nil
}

// SetCoords sets the center coordinates of the ellipse.
func (_dgafg *Ellipse) SetCoords(xc, yc float64) { _dgafg._adbc = xc; _dgafg._ddaee = yc }

// NewPolyline creates a new polyline.
func (_gaf *Creator) NewPolyline(points []_gga.Point) *Polyline { return _edfc(points) }

// SetHeaderRows turns the selected table rows into headers that are repeated
// for every page the table spans. startRow and endRow are inclusive.
func (_gba *Table) SetHeaderRows(startRow, endRow int) error {
	if startRow <= 0 {
		return _bd.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020r\u006f\u0077\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if endRow <= 0 {
		return _bd.New("\u0068\u0065a\u0064\u0065\u0072\u0020e\u006e\u0064 \u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if startRow > endRow {
		return _bd.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0072\u006f\u0077\u0020\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065 \u0065\u006e\u0064\u0020\u0072o\u0077")
	}
	_gba._ddcb = true
	_gba._fefac = startRow
	_gba._gefbf = endRow
	return nil
}
func (_afdfe *templateProcessor) parseLine(_ecgec *templateNode) (interface{}, error) {
	_bbdd := _afdfe.creator.NewLine(0, 0, 0, 0)
	for _, _dace := range _ecgec._adfdg.Attr {
		_bgfec := _dace.Value
		switch _ccdfd := _dace.Name.Local; _ccdfd {
		case "\u0078\u0031":
			_bbdd._gbcgde = _afdfe.parseFloatAttr(_ccdfd, _bgfec)
		case "\u0079\u0031":
			_bbdd._gcfa = _afdfe.parseFloatAttr(_ccdfd, _bgfec)
		case "\u0078\u0032":
			_bbdd._cfge = _afdfe.parseFloatAttr(_ccdfd, _bgfec)
		case "\u0079\u0032":
			_bbdd._ceef = _afdfe.parseFloatAttr(_ccdfd, _bgfec)
		case "\u0074h\u0069\u0063\u006b\u006e\u0065\u0073s":
			_bbdd.SetLineWidth(_afdfe.parseFloatAttr(_ccdfd, _bgfec))
		case "\u0063\u006f\u006co\u0072":
			_bbdd.SetColor(_afdfe.parseColorAttr(_ccdfd, _bgfec))
		case "\u0073\u0074\u0079l\u0065":
			_bbdd.SetStyle(_afdfe.parseLineStyleAttr(_ccdfd, _bgfec))
		case "\u0064\u0061\u0073\u0068\u002d\u0061\u0072\u0072\u0061\u0079":
			_bbdd.SetDashPattern(_afdfe.parseInt64Array(_ccdfd, _bgfec), _bbdd._dcebc)
		case "\u0064\u0061\u0073\u0068\u002d\u0070\u0068\u0061\u0073\u0065":
			_bbdd.SetDashPattern(_bbdd._cgecc, _afdfe.parseInt64Attr(_ccdfd, _bgfec))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_bbdd.SetOpacity(_afdfe.parseFloatAttr(_ccdfd, _bgfec))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_bbdd.SetPositioning(_afdfe.parsePositioningAttr(_ccdfd, _bgfec))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_bbdd.SetFitMode(_afdfe.parseFitModeAttr(_ccdfd, _bgfec))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cbadb := _afdfe.parseMarginAttr(_ccdfd, _bgfec)
			_bbdd.SetMargins(_cbadb.Left, _cbadb.Right, _cbadb.Top, _cbadb.Bottom)
		default:
			_afdfe.nodeLogDebug(_ecgec, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u006e\u0065 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _ccdfd)
		}
	}
	return _bbdd, nil
}
func (_ccfg *templateProcessor) parseChart(_efcdg *templateNode) (interface{}, error) {
	var _accga string
	for _, _abagg := range _efcdg._adfdg.Attr {
		_adcef := _abagg.Value
		switch _degb := _abagg.Name.Local; _degb {
		case "\u0073\u0072\u0063":
			_accga = _adcef
		}
	}
	if _accga == "" {
		_ccfg.nodeLogError(_efcdg, "\u0043\u0068\u0061\u0072\u0074\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _gedbe
	}
	_gfeg, _daaf := _ccfg._cbcec.ChartMap[_accga]
	if !_daaf {
		_ccfg.nodeLogError(_efcdg, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0063ha\u0072\u0074\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _accga)
		return nil, _gedbe
	}
	_eagdf := NewChart(_gfeg)
	for _, _ddgef := range _efcdg._adfdg.Attr {
		_fbbe := _ddgef.Value
		switch _agcd := _ddgef.Name.Local; _agcd {
		case "\u0078":
			_eagdf.SetPos(_ccfg.parseFloatAttr(_agcd, _fbbe), _eagdf._gbcb)
		case "\u0079":
			_eagdf.SetPos(_eagdf._dbf, _ccfg.parseFloatAttr(_agcd, _fbbe))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cfebb := _ccfg.parseMarginAttr(_agcd, _fbbe)
			_eagdf.SetMargins(_cfebb.Left, _cfebb.Right, _cfebb.Top, _cfebb.Bottom)
		case "\u0077\u0069\u0064t\u0068":
			_eagdf._gff.SetWidth(int(_ccfg.parseFloatAttr(_agcd, _fbbe)))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_eagdf._gff.SetHeight(int(_ccfg.parseFloatAttr(_agcd, _fbbe)))
		case "\u0073\u0072\u0063":
			break
		default:
			_ccfg.nodeLogDebug(_efcdg, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0063h\u0061\u0072\u0074\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _agcd)
		}
	}
	return _eagdf, nil
}

// SetMarkedContentID sets the marked content id for the list.
func (_agcea *List) SetMarkedContentID(id int64) *_fgd.KDict { return nil }

// FillOpacity returns the fill opacity of the ellipse (0-1).
func (_dgca *Ellipse) FillOpacity() float64 { return _dgca._fcda }

// MoveTo moves the drawing context to absolute coordinates (x, y).
func (_dcbae *Creator) MoveTo(x, y float64) { _dcbae._eee.X = x; _dcbae._eee.Y = y }
func (_baca *templateProcessor) parseTextRenderingModeAttr(_fceee, _gdcbfa string) TextRenderingMode {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u006d\u006f\u0064e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fceee, _gdcbfa)
	_bggaf := map[string]TextRenderingMode{"\u0066\u0069\u006c\u006c": TextRenderingModeFill, "\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeStroke, "f\u0069\u006c\u006c\u002d\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeFillStroke, "\u0069n\u0076\u0069\u0073\u0069\u0062\u006ce": TextRenderingModeInvisible, "\u0066i\u006c\u006c\u002d\u0063\u006c\u0069p": TextRenderingModeFillClip, "s\u0074\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeStrokeClip, "\u0066\u0069l\u006c\u002d\u0073t\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeFillStrokeClip, "\u0063\u006c\u0069\u0070": TextRenderingModeClip}[_gdcbfa]
	return _bggaf
}

// Add adds a VectorDrawable to the Division container.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Rectangle
// - *Ellipse
// - *Line
// - *Table
// - *Division
// - *List
func (_eggg *Division) Add(d VectorDrawable) error {
	switch _bgdf := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Rectangle, *Ellipse, *Line, *Table, *Division, *List:
	case containerDrawable:
		_dbcf, _ggef := _bgdf.ContainerComponent(_eggg)
		if _ggef != nil {
			return _ggef
		}
		_gdaf, _dfga := _dbcf.(VectorDrawable)
		if !_dfga {
			return _f.Errorf("\u0072\u0065\u0073\u0075\u006ct\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u002d\u0020\u0025\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0056\u0065c\u0074\u006f\u0072\u0044\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020i\u006e\u0074\u0065\u0072\u0066\u0061c\u0065", _dbcf)
		}
		d = _gdaf
	default:
		return _bd.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0044\u0069\u0076\u0069\u0073i\u006f\u006e")
	}
	_eggg._gebd = append(_eggg._gebd, d)
	return nil
}

// SetBorderOpacity sets the border opacity.
func (_fggcf *Polygon) SetBorderOpacity(opacity float64) { _fggcf._bgde = opacity }

// Height returns the current page height.
func (_efg *Creator) Height() float64 { return _efg._agfdf }

// Fit fits the chunk into the specified bounding box, cropping off the
// remainder in a new chunk, if it exceeds the specified dimensions.
// NOTE: The method assumes a line height of 1.0. In order to account for other
// line height values, the passed in height must be divided by the line height:
// height = height / lineHeight
func (_ecaab *TextChunk) Fit(width, height float64) (*TextChunk, error) {
	_ddeba, _abgfbb := _ecaab.Wrap(width)
	if _abgfbb != nil {
		return nil, _abgfbb
	}
	_gcddc := int(height / _ecaab.Style.FontSize)
	if _gcddc >= len(_ddeba) {
		return nil, nil
	}
	_facfd := "\u000a"
	_ecaab.Text = _eg.Replace(_eg.Join(_ddeba[:_gcddc], "\u0020"), _facfd+"\u0020", _facfd, -1)
	_afdc := _eg.Replace(_eg.Join(_ddeba[_gcddc:], "\u0020"), _facfd+"\u0020", _facfd, -1)
	return NewTextChunk(_afdc, _ecaab.Style), nil
}
func _aggdb(_dbebc *_fgd.PdfFont, _eaad float64) *fontMetrics {
	_geace := &fontMetrics{}
	if _dbebc == nil {
		_fec.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0069s\u0020\u006e\u0069\u006c")
		return _geace
	}
	_beefc, _agefb := _dbebc.GetFontDescriptor()
	if _agefb != nil {
		_fec.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0067\u0065t\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063ri\u0070\u0074\u006fr\u003a \u0025\u0076", _agefb)
		return _geace
	}
	if _geace._geede, _agefb = _beefc.GetCapHeight(); _agefb != nil {
		_fec.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _agefb)
	}
	if int(_geace._geede) <= 0 {
		_fec.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0043\u0061p\u0020\u0048\u0065ig\u0068\u0074\u0020\u006e\u006f\u0074 \u0061\u0076\u0061\u0069\u006c\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065\u0074t\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u00310\u0030\u0030")
		_geace._geede = 1000
	}
	_geace._geede *= _eaad / 1000.0
	if _geace._dfeba, _agefb = _beefc.GetXHeight(); _agefb != nil {
		_fec.Log.Trace("\u0057\u0041R\u004e\u003a\u0020\u0055n\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0078\u002d\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _agefb)
	}
	_geace._dfeba *= _eaad / 1000.0
	if _geace._feeec, _agefb = _beefc.GetAscent(); _agefb != nil {
		_fec.Log.Trace("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0061\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _agefb)
	}
	_geace._feeec *= _eaad / 1000.0
	if _geace._egeb, _agefb = _beefc.GetDescent(); _agefb != nil {
		_fec.Log.Trace("\u0057\u0041RN\u003a\u0020\u0055n\u0061\u0062\u006c\u0065 to\u0020ge\u0074\u0020\u0066\u006f\u006e\u0074\u0020de\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _agefb)
	}
	_geace._egeb *= _eaad / 1000.0
	return _geace
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_bcdg *Invoice) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dbef := ctx
	_agbd := []func(_dbaba DrawContext) ([]*Block, DrawContext, error){_bcdg.generateHeaderBlocks, _bcdg.generateInformationBlocks, _bcdg.generateLineBlocks, _bcdg.generateTotalBlocks, _bcdg.generateNoteBlocks}
	var _ebbd []*Block
	for _, _bfeb := range _agbd {
		_bege, _dcfd, _ceceg := _bfeb(ctx)
		if _ceceg != nil {
			return _ebbd, ctx, _ceceg
		}
		if len(_ebbd) == 0 {
			_ebbd = _bege
		} else if len(_bege) > 0 {
			_ebbd[len(_ebbd)-1].mergeBlocks(_bege[0])
			_ebbd = append(_ebbd, _bege[1:]...)
		}
		ctx = _dcfd
	}
	if _bcdg._fefb.IsRelative() {
		ctx.X = _dbef.X
	}
	if _bcdg._fefb.IsAbsolute() {
		return _ebbd, _dbef, nil
	}
	return _ebbd, ctx, nil
}
func (_ccede *templateProcessor) parseInt64Array(_bgeea, _eeec string) []int64 {
	_fec.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0069\u006e\u0074\u0036\u0034\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _bgeea, _eeec)
	_aecddf := _eg.Fields(_eeec)
	_gcbfc := make([]int64, 0, len(_aecddf))
	for _, _dfcgc := range _aecddf {
		_dcaee, _ := _fg.ParseInt(_dfcgc, 10, 64)
		_gcbfc = append(_gcbfc, _dcaee)
	}
	return _gcbfc
}

// NewChapter creates a new chapter with the specified title as the heading.
func (_dedc *Creator) NewChapter(title string) *Chapter {
	_dedc._dca++
	_gaeg := _dedc.NewTextStyle()
	_gaeg.FontSize = 16
	return _dgg(nil, _dedc._aac, _dedc._cbba, title, _dedc._dca, _gaeg)
}

var _fbf = _d.MustCompile("\u005c\u0064\u002b")

// AddColorStop add color stop info for rendering gradient color.
func (_acdce *LinearShading) AddColorStop(color Color, point float64) {
	_acdce._dfae.AddColorStop(color, point)
}

// SetLineTitleStyle sets the style for the title part of all new lines
// of the table of contents.
func (_gbecc *TOC) SetLineTitleStyle(style TextStyle) { _gbecc._fefdb = style }

// SetSellerAddress sets the seller address of the invoice.
func (_bdaae *Invoice) SetSellerAddress(address *InvoiceAddress) { _bdaae._ddaa = address }
func (_bbfda *StyledParagraph) wrapWordChunks() {
	if !_bbfda._eadc {
		return
	}
	var (
		_bdeb  []*TextChunk
		_cddeb *_fgd.PdfFont
	)
	for _, _cfdda := range _bbfda._ecec {
		_eacbf := []rune(_cfdda.Text)
		if _cddeb == nil {
			_cddeb = _cfdda.Style.Font
		}
		_decge := _cfdda._abefd
		_bdffb := _cfdda.VerticalAlignment
		if len(_bdeb) > 0 {
			if len(_eacbf) == 1 && _fe.IsPunct(_eacbf[0]) && _cfdda.Style.Font == _cddeb {
				_dagbd := []rune(_bdeb[len(_bdeb)-1].Text)
				_bdeb[len(_bdeb)-1].Text = string(append(_dagbd, _eacbf[0]))
				continue
			} else {
				_, _gacbf := _fg.Atoi(_cfdda.Text)
				if _gacbf == nil {
					_gbdaf := []rune(_bdeb[len(_bdeb)-1].Text)
					_bbecf := len(_gbdaf)
					if _bbecf >= 2 {
						_, _eccde := _fg.Atoi(string(_gbdaf[_bbecf-2]))
						if _eccde == nil && _fe.IsPunct(_gbdaf[_bbecf-1]) {
							_bdeb[len(_bdeb)-1].Text = string(append(_gbdaf, _eacbf...))
							continue
						}
					}
				}
			}
		}
		_bcfd, _ffagf := _feff(_cfdda.Text)
		if _ffagf != nil {
			_fec.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0062\u0072\u0065\u0061\u006b\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020w\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0076", _ffagf)
			_bcfd = []string{_cfdda.Text}
		}
		for _, _adaac := range _bcfd {
			_ffef := NewTextChunk(_adaac, _cfdda.Style)
			_ffef._abefd = _bggbd(_decge)
			_ffef.VerticalAlignment = _bdffb
			_bdeb = append(_bdeb, _ffef)
		}
		_cddeb = _cfdda.Style.Font
	}
	if len(_bdeb) > 0 {
		_bbfda._ecec = _bdeb
	}
}

// Width returns the width of the graphic svg.
func (_ddbdf *GraphicSVG) Width() float64 { return _ddbdf._dgccb.Width }

// SetPos sets absolute positioning with specified coordinates.
func (_eaeeg *Paragraph) SetPos(x, y float64) {
	_eaeeg._dbgab = PositionAbsolute
	_eaeeg._gddg = x
	_eaeeg._ggfb = y
}

// ScaleToWidth scales the ellipse to the specified width. The height of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_gcg *Ellipse) ScaleToWidth(w float64) {
	_fdebc := _gcg._beb / _gcg._eefg
	_gcg._eefg = w
	_gcg._beb = w * _fdebc
}

// SetFillColor sets the fill color.
func (_fdade *PolyBezierCurve) SetFillColor(color Color) {
	_fdade._bbaef = color
	_fdade._ffcd.FillColor = _cfcee(color)
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_gddge *TOC) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dgfcc := ctx
	_dbacf, ctx, _egdg := _gddge._defae.GeneratePageBlocks(ctx)
	if _egdg != nil {
		return _dbacf, ctx, _egdg
	}
	for _, _baeac := range _gddge._ccfec {
		_ccdde := _baeac._ebaab
		if !_gddge._ggabc {
			_baeac._ebaab = 0
		}
		_bcfce, _gbeb, _fbbag := _baeac.GeneratePageBlocks(ctx)
		_baeac._ebaab = _ccdde
		if _fbbag != nil {
			return _dbacf, ctx, _fbbag
		}
		if len(_bcfce) < 1 {
			continue
		}
		_dbacf[len(_dbacf)-1].mergeBlocks(_bcfce[0])
		_dbacf = append(_dbacf, _bcfce[1:]...)
		ctx = _gbeb
	}
	if _gddge._gabag.IsRelative() {
		ctx.X = _dgfcc.X
	}
	if _gddge._gabag.IsAbsolute() {
		return _dbacf, _dgfcc, nil
	}
	return _dbacf, ctx, nil
}
func _cgcd(_acgf _e.Image) (*Image, error) {
	_dgfff, _dfef := _fgd.ImageHandling.NewImageFromGoImage(_acgf)
	if _dfef != nil {
		return nil, _dfef
	}
	return _eddeb(_dgfff)
}
func _egeag(_afdb *templateProcessor, _fabc *templateNode) (interface{}, error) {
	return _afdb.parseTable(_fabc)
}

// SetMarkedContentID sets marked content ID.
func (_ebgd *FilledCurve) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_ebgd._deca = &mcid
	_fed := _fgd.NewKDictionary()
	_fed.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_fed.K = _bc.MakeInteger(mcid)
	return _fed
}

// Width returns the width of the line.
// NOTE: Depending on the fit mode the line is set to use, its width may be
// calculated at runtime (e.g. when using FitModeFillWidth).
func (_dedb *Line) Width() float64 { return _fc.Abs(_dedb._cfge - _dedb._gbcgde) }

// Width returns the width of the chart. In relative positioning mode,
// all the available context width is used at render time.
func (_aaea *Chart) Width() float64 { return float64(_aaea._gff.Width()) }

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_aacf *RadialShading) SetAntiAlias(enable bool) { _aacf._dafd.SetAntiAlias(enable) }

var PPMM = float64(72 * 1.0 / 25.4)

// Height returns the height of the division, assuming all components are
// stacked on top of each other.
func (_bbd *Division) Height() float64 {
	var _egga float64
	for _, _dccb := range _bbd._gebd {
		switch _fafce := _dccb.(type) {
		case marginDrawable:
			_, _, _faae, _abfe := _fafce.GetMargins()
			_egga += _fafce.Height() + _faae + _abfe
		default:
			_egga += _fafce.Height()
		}
	}
	return _egga
}
func _dgg(_gegf *Chapter, _edec *TOC, _egfc *_fgd.Outline, _effb string, _dfa int, _eddg TextStyle) *Chapter {
	var _dcc uint = 1
	if _gegf != nil {
		_dcc = _gegf._eedg + 1
	}
	_gefc := &Chapter{_gbf: _dfa, _faed: _effb, _aee: true, _aga: true, _cdee: _gegf, _edc: _edec, _agbcd: _egfc, _agbc: []Drawable{}, _eedg: _dcc}
	_ggb := _cdbe(_gefc.headingText(), _eddg)
	_ggb.SetFont(_eddg.Font)
	_ggb.SetFontSize(_eddg.FontSize)
	_gefc._gfdf = _ggb
	return _gefc
}

// SetPositioning sets the positioning of the ellipse (absolute or relative).
func (_egcb *Ellipse) SetPositioning(position Positioning) { _egcb._fdbc = position }

// Division is a container component which can wrap across multiple pages.
// Currently supported drawable components:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
//
// The component stacking behavior is vertical, where the drawables are drawn
// on top of each other.
type Division struct {
	_gebd  []VectorDrawable
	_fbec  Positioning
	_ddfg  Margins
	_gbga  Margins
	_fdb   bool
	_defa  bool
	_agecb *Background
}

// SetFillColor sets the fill color of the ellipse.
func (_cfceg *Ellipse) SetFillColor(col Color) { _cfceg._baacc = col }

// MoveRight moves the drawing context right by relative displacement dx (negative goes left).
func (_aecda *Creator) MoveRight(dx float64) { _aecda._eee.X += dx }

// Curve represents a cubic Bezier curve with a control point.
type Curve struct {
	_aecc  float64
	_egda  float64
	_bgaca float64
	_afbc  float64
	_cdf   float64
	_adca  float64
	_afgg  Color
	_befg  float64
	_afbbc *int64
}

func (_ebfac *templateProcessor) parseStyledParagraph(_aebfb *templateNode) (interface{}, error) {
	_ebagd := _ebfac.creator.NewStyledParagraph()
	for _, _efbd := range _aebfb._adfdg.Attr {
		_bdgg := _efbd.Value
		switch _abeff := _efbd.Name.Local; _abeff {
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_ebagd.SetTextAlignment(_ebfac.parseTextAlignmentAttr(_abeff, _bdgg))
		case "\u0076\u0065\u0072\u0074ic\u0061\u006c\u002d\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_ebagd.SetTextVerticalAlignment(_ebfac.parseTextVerticalAlignmentAttr(_abeff, _bdgg))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_ebagd.SetLineHeight(_ebfac.parseFloatAttr(_abeff, _bdgg))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_eabca := _ebfac.parseMarginAttr(_abeff, _bdgg)
			_ebagd.SetMargins(_eabca.Left, _eabca.Right, _eabca.Top, _eabca.Bottom)
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_ebagd.SetEnableWrap(_ebfac.parseBoolAttr(_abeff, _bdgg))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0077\u006f\u0072\u0064\u002d\u0077\u0072\u0061\u0070":
			_ebagd.EnableWordWrap(_ebfac.parseBoolAttr(_abeff, _bdgg))
		case "\u0074\u0065\u0078\u0074\u002d\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077":
			_ebagd.SetTextOverflow(_ebfac.parseTextOverflowAttr(_abeff, _bdgg))
		case "\u0078":
			_ebagd.SetPos(_ebfac.parseFloatAttr(_abeff, _bdgg), _ebagd._agdg)
		case "\u0079":
			_ebagd.SetPos(_ebagd._abeag, _ebfac.parseFloatAttr(_abeff, _bdgg))
		case "\u0061\u006e\u0067l\u0065":
			_ebagd.SetAngle(_ebfac.parseFloatAttr(_abeff, _bdgg))
		default:
			_ebfac.nodeLogDebug(_aebfb, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0073\u0074\u0079l\u0065\u0064\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0060\u0025\u0073`.\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _abeff)
		}
	}
	return _ebagd, nil
}
func _dcad(_edcgc *Block, _eadd *StyledParagraph, _aadaa [][]*TextChunk, _efgg DrawContext) (DrawContext, [][]*TextChunk, error) {
	_cfdbd := 1
	_aeeec := _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
	for _edcgc._dgd.HasFontByName(_aeeec) {
		_cfdbd++
		_aeeec = _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
	}
	_cggc := _edcgc._dgd.SetFontByName(_aeeec, _eadd._fgbg.Font.ToPdfObject())
	if _cggc != nil {
		return _efgg, nil, _cggc
	}
	_cfdbd++
	_fgefa := _aeeec
	_adcf := _eadd._fgbg.FontSize
	_gdce := _eadd._fceb.IsRelative()
	var _bbaf [][]_bc.PdfObjectName
	var _ddbeg [][]*TextChunk
	var _geedb float64
	for _ddgaa, _cdbfg := range _aadaa {
		var _fafcdb []_bc.PdfObjectName
		var _gcegd float64
		if len(_cdbfg) > 0 {
			_gcegd = _cdbfg[0].Style.FontSize
		}
		for _, _daga := range _cdbfg {
			_fdgde := _daga.Style
			if _daga.Text != "" && _fdgde.FontSize > _gcegd {
				_gcegd = _fdgde.FontSize
			}
			if _gcegd > _efgg.PageHeight {
				return _efgg, nil, _bd.New("\u0050\u0061\u0072\u0061\u0067\u0072a\u0070\u0068\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020\u006ca\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0070\u0061\u0067\u0065 \u0068e\u0069\u0067\u0068\u0074")
			}
			_aeeec = _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
			_dacd := _edcgc._dgd.SetFontByName(_aeeec, _fdgde.Font.ToPdfObject())
			if _dacd != nil {
				return _efgg, nil, _dacd
			}
			_fafcdb = append(_fafcdb, _aeeec)
			_cfdbd++
		}
		_gcegd *= _eadd._fgef
		if _gdce && _geedb+_gcegd > _efgg.Height {
			_ddbeg = _aadaa[_ddgaa:]
			_aadaa = _aadaa[:_ddgaa]
			break
		}
		_geedb += _gcegd
		_bbaf = append(_bbaf, _fafcdb)
	}
	_cbdbd, _dabea, _dgbcc := _eadd.getLineMetrics(0)
	_gdfbf, _bccce := _cbdbd*_eadd._fgef, _dabea*_eadd._fgef
	if len(_aadaa) == 0 {
		return _efgg, _ddbeg, nil
	}
	_aagd := _dg.NewContentCreator()
	_aagd.Add_q()
	_ggab := _bccce
	if _eadd._ddec == TextVerticalAlignmentCenter {
		_ggab = _dabea + (_cbdbd+_dgbcc-_dabea)/2 + (_bccce-_dabea)/2
	}
	_gaddd := _efgg.PageHeight - _efgg.Y - _ggab
	_aagd.Translate(_efgg.X, _gaddd)
	_gbcac := _gaddd
	if _eadd._ffab != 0 {
		_aagd.RotateDeg(_eadd._ffab)
	}
	if _eadd._ecedg == TextOverflowHidden {
		_aagd.Add_re(0, -_geedb+_gdfbf+1, _eadd._ffcfd, _geedb).Add_W().Add_n()
	}
	_aagd.Add_BT()
	_egaf := map[string]_bc.PdfObject{}
	if _eadd._eeba != nil {
		_egaf["\u004d\u0043\u0049\u0044"] = _bc.MakeInteger(*_eadd._eeba)
	}
	if _eadd._cccab != "" {
		_egaf["\u004c\u0061\u006e\u0067"] = _bc.MakeString(_eadd._cccab)
	}
	if len(_egaf) > 0 {
		_aagd.Add_BDC(*_bc.MakeName(_fgd.StructureTypeParagraph), _egaf)
	}
	var _aefee []*_gga.BasicLine
	for _ddff, _efca := range _aadaa {
		_ffcfb := _efgg.X
		var _dggf float64
		if len(_efca) > 0 {
			_dggf = _efca[0].Style.FontSize
		}
		_cbdbd, _, _dgbcc = _eadd.getLineMetrics(_ddff)
		_bccce = (_cbdbd + _dgbcc)
		for _, _dcbbg := range _efca {
			_gabg := &_dcbbg.Style
			if _dcbbg.Text != "" && _gabg.FontSize > _dggf {
				_dggf = _gabg.FontSize
			}
			if _bccce > _dggf {
				_dggf = _bccce
			}
		}
		if _ddff != 0 {
			_aagd.Add_TD(0, -_dggf*_eadd._fgef)
			_gbcac -= _dggf * _eadd._fgef
		}
		_debg := _ddff == len(_aadaa)-1
		var (
			_agfg  float64
			_aeada float64
			_ebffg *fontMetrics
			_agfbc float64
			_befa  uint
		)
		var _ggac []float64
		for _, _dabb := range _efca {
			_ebbb := &_dabb.Style
			if _ebbb.FontSize > _aeada {
				_aeada = _ebbb.FontSize
				_ebffg = _aggdb(_dabb.Style.Font, _ebbb.FontSize)
			}
			if _bccce > _aeada {
				_aeada = _bccce
			}
			_ecdbb, _eadge := _ebbb.Font.GetRuneMetrics(' ')
			if _ecdbb.Wx == 0 && _ebbb.MultiFont != nil {
				_ecdbb, _eadge = _ebbb.MultiFont.GetRuneMetrics(' ')
				_ebbb.MultiFont.Reset()
			}
			if !_eadge {
				return _efgg, nil, _bd.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
			}
			var _edbfg uint
			var _ggcaf float64
			_aade := len(_dabb.Text)
			for _fefa, _ddfeb := range _dabb.Text {
				if _ddfeb == ' ' {
					_edbfg++
					continue
				}
				if _ddfeb == '\u000A' {
					continue
				}
				_agda, _adgf := _ebbb.Font.GetRuneMetrics(_ddfeb)
				if _agda.Wx == 0 && _ebbb.MultiFont != nil {
					_agda, _adgf = _ebbb.MultiFont.GetRuneMetrics(' ')
					_ebbb.MultiFont.Reset()
				}
				if !_adgf {
					_fec.Log.Debug("\u0055\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006ee\u0020%\u0076\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u000a", _ddfeb)
					return _efgg, nil, _bd.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
				}
				_ggcaf += _ebbb.FontSize * _agda.Wx * _ebbb.horizontalScale()
				if _fefa != _aade-1 {
					_ggcaf += _ebbb.CharSpacing * 1000.0
				}
			}
			_ggac = append(_ggac, _ggcaf)
			_agfg += _ggcaf
			_agfbc += float64(_edbfg) * _ecdbb.Wx * _ebbb.FontSize * _ebbb.horizontalScale()
			_befa += _edbfg
		}
		_aeada *= _eadd._fgef
		var _ccfca []_bc.PdfObject
		_dccbe := _eadd._ffcfd * 1000.0
		if _eadd._abff == TextAlignmentJustify {
			if _befa > 0 && !_debg {
				_agfbc = (_dccbe - _agfg) / float64(_befa) / _adcf
			}
		} else if _eadd._abff == TextAlignmentCenter {
			_gcgg := (_dccbe - _agfg - _agfbc) / 2
			_dadf := _gcgg / _adcf
			_ccfca = append(_ccfca, _bc.MakeFloat(-_dadf))
			_ffcfb += _gcgg / 1000.0
		} else if _eadd._abff == TextAlignmentRight {
			_dcbe := (_dccbe - _agfg - _agfbc)
			_gggd := _dcbe / _adcf
			_ccfca = append(_ccfca, _bc.MakeFloat(-_gggd))
			_ffcfb += _dcbe / 1000.0
		}
		if len(_ccfca) > 0 {
			_aagd.Add_Tf(_fgefa, _adcf).Add_TL(_adcf * _eadd._fgef).Add_TJ(_ccfca...)
		}
		_baedf := 0.0
		for _dccde, _acca := range _efca {
			_fgfc := &_acca.Style
			_fdgf := _fgefa
			_dccef := _adcf
			_gadcc := _fgfc.OutlineColor != nil
			_aaad := _fgfc.HorizontalScaling != DefaultHorizontalScaling
			_bedc := _fgfc.OutlineSize != 1
			if _bedc {
				_aagd.Add_w(_fgfc.OutlineSize)
			}
			_efcaf := _fgfc.RenderingMode != TextRenderingModeFill
			if _efcaf {
				_aagd.Add_Tr(int64(_fgfc.RenderingMode))
			}
			_edcb := _fgfc.CharSpacing != 0
			if _edcb {
				_aagd.Add_Tc(_fgfc.CharSpacing)
			}
			_gdeg := _fgfc.TextRise != 0
			if _gdeg {
				_aagd.Add_Ts(_fgfc.TextRise)
			}
			if _acca.VerticalAlignment != TextVerticalAlignmentBaseline {
				_beeb := _aggdb(_acca.Style.Font, _fgfc.FontSize)
				switch _acca.VerticalAlignment {
				case TextVerticalAlignmentCenter:
					_baedf = _ebffg._dfeba/2 - _beeb._dfeba/2
				case TextVerticalAlignmentBottom:
					_baedf = _ebffg._egeb - _beeb._egeb
				case TextVerticalAlignmentTop:
					_baedf = _dabea - _fgfc.FontSize
				}
				if _baedf != 0.0 {
					_aagd.Translate(0, _baedf)
				}
			}
			if _eadd._abff != TextAlignmentJustify || _debg {
				_ecbc, _gbbf := _fgfc.Font.GetRuneMetrics(' ')
				if !_gbbf {
					return _efgg, nil, _bd.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
				}
				_fdgf = _bbaf[_ddff][_dccde]
				_dccef = _fgfc.FontSize
				_agfbc = _ecbc.Wx * _fgfc.horizontalScale()
			}
			_gcgf := _fgfc.Font.Encoder()
			var _agdee []byte
			var _ggcc bool
			_bcea := _fgfc.Font
			for _, _afbe := range _acca.Text {
				if _afbe == '\u000A' {
					continue
				}
				if _afbe == ' ' {
					if len(_agdee) > 0 {
						if _gadcc {
							_aagd.SetStrokingColor(_cfcee(_fgfc.OutlineColor))
						}
						if _aaad {
							_aagd.Add_Tz(_fgfc.HorizontalScaling)
						}
						_gaabb := _bbaf[_ddff][_dccde]
						if _ggcc {
							_gaabb = _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
							_fffdc := _edcgc._dgd.SetFontByName(_gaabb, _bcea.ToPdfObject())
							if _fffdc != nil {
								return _efgg, nil, _fffdc
							}
							_cfdbd++
							_ggcc = false
							_gcgf = _fgfc.Font.Encoder()
						}
						_aagd.SetNonStrokingColor(_cfcee(_fgfc.Color)).Add_Tf(_gaabb, _fgfc.FontSize).Add_TJ([]_bc.PdfObject{_bc.MakeStringFromBytes(_agdee)}...)
						_agdee = nil
					}
					if _aaad {
						_aagd.Add_Tz(DefaultHorizontalScaling)
					}
					_aagd.Add_Tf(_fdgf, _dccef).Add_TJ([]_bc.PdfObject{_bc.MakeFloat(-_agfbc)}...)
					_ggac[_dccde] += _agfbc * _dccef
				} else {
					if _, _gfff := _gcgf.RuneToCharcode(_afbe); !_gfff {
						if _fgfc.MultiFont != nil {
							_dbda, _gcebf := _fgfc.MultiFont.Encoder(_afbe)
							if _gcebf {
								if len(_agdee) != 0 {
									_bggd := _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
									_cgcc := _edcgc._dgd.SetFontByName(_fdgf, _bcea.ToPdfObject())
									if _cgcc != nil {
										return _efgg, nil, _cgcc
									}
									_aagd.SetNonStrokingColor(_cfcee(_fgfc.Color)).Add_Tf(_bggd, _fgfc.FontSize).Add_TJ([]_bc.PdfObject{_bc.MakeStringFromBytes(_agdee)}...)
									_cfdbd++
									_agdee = nil
								}
								_gcgf = _dbda
								_ggcc = true
								_bcea = _fgfc.MultiFont.CurrentFont
							}
						} else {
							_cggc = UnsupportedRuneError{Message: _f.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _afbe, _afbe), Rune: _afbe}
							_efgg._agd = append(_efgg._agd, _cggc)
							_fec.Log.Debug(_cggc.Error())
							if _efgg._cffd <= 0 {
								continue
							}
							_afbe = _efgg._cffd
						}
					}
					_bagc := _gcgf.Encode(string(_afbe))
					_agdee = append(_agdee, _bagc...)
				}
				if _fgfc.MultiFont != nil {
					_fgfc.MultiFont.Reset()
				}
			}
			if len(_agdee) > 0 {
				if _gadcc {
					_aagd.SetStrokingColor(_cfcee(_fgfc.OutlineColor))
				}
				if _aaad {
					_aagd.Add_Tz(_fgfc.HorizontalScaling)
				}
				_gdad := _bbaf[_ddff][_dccde]
				if _ggcc {
					_gdad = _bc.PdfObjectName(_f.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cfdbd))
					_cecegf := _edcgc._dgd.SetFontByName(_gdad, _bcea.ToPdfObject())
					if _cecegf != nil {
						return _efgg, nil, _cecegf
					}
					_cfdbd++
					_ggcc = false
				}
				_aagd.SetNonStrokingColor(_cfcee(_fgfc.Color)).Add_Tf(_gdad, _fgfc.FontSize).Add_TJ([]_bc.PdfObject{_bc.MakeStringFromBytes(_agdee)}...)
			}
			_gacca := _ggac[_dccde] / 1000.0
			if _fgfc.Underline {
				_caca := _fgfc.UnderlineStyle.Color
				if _caca == nil {
					_caca = _acca.Style.Color
				}
				_cfgbe, _dcdgf, _geafc := _caca.ToRGB()
				_aaded := _ffcfb - _efgg.X
				_gbcee := _gbcac - _gaddd + _fgfc.TextRise - _fgfc.UnderlineStyle.Offset
				_aefee = append(_aefee, &_gga.BasicLine{X1: _aaded, Y1: _gbcee, X2: _aaded + _gacca, Y2: _gbcee, LineWidth: _acca.Style.UnderlineStyle.Thickness, LineColor: _fgd.NewPdfColorDeviceRGB(_cfgbe, _dcdgf, _geafc)})
			}
			if _acca._abefd != nil {
				var _cfdg *_bc.PdfObjectArray
				if !_acca._ebcfe {
					switch _bgfee := _acca._abefd.GetContext().(type) {
					case *_fgd.PdfAnnotationLink:
						_cfdg = _bc.MakeArray()
						_bgfee.Rect = _cfdg
						_aabde, _cfgbb := _bgfee.Dest.(*_bc.PdfObjectArray)
						if _cfgbb && _aabde.Len() == 5 {
							_dcda, _cgfa := _aabde.Get(1).(*_bc.PdfObjectName)
							if _cgfa && _dcda.String() == "\u0058\u0059\u005a" {
								_fdeea, _bdacf := _bc.GetNumberAsFloat(_aabde.Get(3))
								if _bdacf == nil {
									_aabde.Set(3, _bc.MakeFloat(_efgg.PageHeight-_fdeea))
								}
							}
						}
					}
					_acca._ebcfe = true
				}
				if _cfdg != nil {
					_ddgg := _gga.NewPoint(_ffcfb-_efgg.X, _gbcac+_fgfc.TextRise-_gaddd).Rotate(_eadd._ffab)
					_ddgg.X += _efgg.X
					_ddgg.Y += _gaddd
					_ggdbe, _dfcb, _gccgd, _adcfa := _ffcgd(_gacca, _aeada, _eadd._ffab)
					_ddgg.X += _ggdbe
					_ddgg.Y += _dfcb
					_cfdg.Clear()
					_cfdg.Append(_bc.MakeFloat(_ddgg.X))
					_cfdg.Append(_bc.MakeFloat(_ddgg.Y))
					_cfdg.Append(_bc.MakeFloat(_ddgg.X + _gccgd))
					_cfdg.Append(_bc.MakeFloat(_ddgg.Y + _adcfa))
				}
				_edcgc.AddAnnotation(_acca._abefd)
			}
			_ffcfb += _gacca
			if _bedc {
				_aagd.Add_w(1.0)
			}
			if _gadcc {
				_aagd.Add_RG(0.0, 0.0, 0.0)
			}
			if _efcaf {
				_aagd.Add_Tr(int64(TextRenderingModeFill))
			}
			if _edcb {
				_aagd.Add_Tc(0)
			}
			if _gdeg {
				_aagd.Add_Ts(0)
			}
			if _aaad {
				_aagd.Add_Tz(DefaultHorizontalScaling)
			}
			if _baedf != 0.0 {
				_aagd.Translate(0, -_baedf)
				_baedf = 0.0
			}
		}
	}
	if len(_egaf) > 0 {
		_aagd.Add_EMC()
	}
	_aagd.Add_ET()
	for _, _cbfac := range _aefee {
		_aagd.SetStrokingColor(_cbfac.LineColor).Add_w(_cbfac.LineWidth).Add_m(_cbfac.X1, _cbfac.Y1).Add_l(_cbfac.X2, _cbfac.Y2).Add_s()
	}
	_aagd.Add_Q()
	_beaba := _aagd.Operations()
	_beaba.WrapIfNeeded()
	_edcgc.addContents(_beaba)
	if _gdce {
		_gdbee := _geedb
		_efgg.Y += _gdbee
		_efgg.Height -= _gdbee
		if _efgg.Inline {
			_efgg.X += _eadd.Width() + _eadd._ccddg.Right
		}
	}
	return _efgg, _ddbeg, nil
}
func _dgbe(_cfddd, _afdd, _ebfg, _gced, _bddc, _gdd float64) *Curve {
	_eaff := &Curve{}
	_eaff._aecc = _cfddd
	_eaff._egda = _afdd
	_eaff._bgaca = _ebfg
	_eaff._afbc = _gced
	_eaff._cdf = _bddc
	_eaff._adca = _gdd
	_eaff._afgg = ColorBlack
	_eaff._befg = 1.0
	return _eaff
}

// ConvertToBinary converts current image data into binary (Bi-level image) format.
// If provided image is RGB or GrayScale the function converts it into binary image
// using histogram auto threshold method.
func (_becf *Image) ConvertToBinary() error { return _becf._cgdfb.ConvertToBinary() }

type marginDrawable interface {
	VectorDrawable
	GetMargins() (float64, float64, float64, float64)
}

// SetPos sets the Block's positioning to absolute mode with the specified coordinates.
func (_ca *Block) SetPos(x, y float64) { _ca._ad = PositionAbsolute; _ca._bf = x; _ca._cbe = y }

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

// SetBorderColor sets the border color.
func (_bbdbb *Polygon) SetBorderColor(color Color) { _bbdbb._acdc.BorderColor = _cfcee(color) }

// SetFitMode sets the fit mode of the ellipse.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_fdc *Ellipse) SetFitMode(fitMode FitMode) { _fdc._eddea = fitMode }

// TOC returns the table of contents component of the creator.
func (_ceda *Creator) TOC() *TOC { return _ceda._aac }
func (_edeb *Creator) getActivePage() *_fgd.PdfPage {
	if _edeb._eda == nil {
		if len(_edeb._egfb) == 0 {
			return nil
		}
		return _edeb._egfb[len(_edeb._egfb)-1]
	}
	return _edeb._eda
}

// SetStyle sets the style for all the line components: number, title,
// separator, page.
func (_acbfd *TOCLine) SetStyle(style TextStyle) {
	_acbfd.Number.Style = style
	_acbfd.Title.Style = style
	_acbfd.Separator.Style = style
	_acbfd.Page.Style = style
}

// Lazy gets the lazy mode for the image.
func (_gffb *Image) Lazy() bool { return _gffb._faeda }

// Height returns the height of the graphic svg.
func (_dbec *GraphicSVG) Height() float64 { return _dbec._dgccb.Height }

// SetBackground sets the background properties of the component.
func (_gebgg *Division) SetBackground(background *Background) { _gebgg._agecb = background }
func _bdgfg(_bffa *Creator, _ggbdc _gg.Reader, _ddab interface{}, _dfca *TemplateOptions, _gccdd componentRenderer) error {
	if _bffa == nil {
		_fec.Log.Error("\u0043\u0072\u0065a\u0074\u006f\u0072\u0020i\u006e\u0073\u0074\u0061\u006e\u0063\u0065 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return _cbgfd
	}
	_fccb := ""
	if _ffgfe, _geeff := _ggbdc.(*_b.File); _geeff {
		_fccb = _ffgfe.Name()
	}
	_bebed := _c.NewBuffer(nil)
	if _, _gaaa := _gg.Copy(_bebed, _ggbdc); _gaaa != nil {
		return _gaaa
	}
	_eegeb := _g.FuncMap{"\u0064\u0069\u0063\u0074": _dgcef, "\u0061\u0064\u0064": _gadda, "\u0061\u0072\u0072a\u0079": _faad, "\u0065\u0078\u0074\u0065\u006e\u0064\u0044\u0069\u0063\u0074": _fegbd, "\u006da\u006b\u0065\u0053\u0065\u0071": _cbcc}
	if _dfca != nil && _dfca.HelperFuncMap != nil {
		for _fbde, _ceaf := range _dfca.HelperFuncMap {
			if _, _dgce := _eegeb[_fbde]; _dgce {
				_fec.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074 \u006f\u0076\u0065r\u0072\u0069\u0064e\u0020\u0062\u0075\u0069\u006c\u0074\u002d\u0069\u006e\u0020`\u0025\u0073\u0060\u0020\u0068el\u0070\u0065\u0072\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _fbde)
				continue
			}
			_eegeb[_fbde] = _ceaf
		}
	}
	_adgfd, _gaae := _g.New("").Funcs(_eegeb).Parse(_bebed.String())
	if _gaae != nil {
		return _gaae
	}
	if _dfca != nil && _dfca.SubtemplateMap != nil {
		for _agdgf, _gfgdb := range _dfca.SubtemplateMap {
			if _agdgf == "" {
				_fec.Log.Debug("\u0053\u0075\u0062\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067.\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
				continue
			}
			if _gfgdb == nil {
				_fec.Log.Debug("S\u0075\u0062t\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u0061\u006e\u006eo\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079 \u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e")
				continue
			}
			_bcdb := _c.NewBuffer(nil)
			if _, _ggcff := _gg.Copy(_bcdb, _gfgdb); _ggcff != nil {
				return _ggcff
			}
			if _, _gaeaa := _adgfd.New(_agdgf).Parse(_bcdb.String()); _gaeaa != nil {
				return _gaeaa
			}
		}
	}
	_bebed.Reset()
	if _ccgcf := _adgfd.Execute(_bebed, _ddab); _ccgcf != nil {
		return _ccgcf
	}
	return _bbdg(_bffa, _fccb, _bebed.Bytes(), _dfca, _gccdd).run()
}

// SetMarkedContentID sets marked content ID.
func (_aggbd *GraphicSVG) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_aggbd._bdae = &mcid
	_bcga := _fgd.NewKDictionary()
	_bcga.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_bcga.K = _bc.MakeInteger(mcid)
	return _bcga
}

// SetBorderRadius sets the radius of the rectangle corners.
func (_abge *Rectangle) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_abge._ggeb = topLeft
	_abge._cbbc = topRight
	_abge._ggfa = bottomLeft
	_abge._gcefb = bottomRight
}

// Chapter is used to arrange multiple drawables (paragraphs, images, etc) into a single section.
// The concept is the same as a book or a report chapter.
type Chapter struct {
	_gbf         int
	_faed        string
	_gfdf        *Paragraph
	_agbc        []Drawable
	_bdcd        int
	_aee         bool
	_aga         bool
	_bda         Positioning
	_cdgg, _gdfg float64
	_dbe         Margins
	_cdee        *Chapter
	_edc         *TOC
	_agbcd       *_fgd.Outline
	_bcaf        *_fgd.OutlineItem
	_eedg        uint
}

// GetMargins returns the left, right, top, bottom Margins.
func (_adgd *Table) GetMargins() (float64, float64, float64, float64) {
	return _adgd._gbcea.Left, _adgd._gbcea.Right, _adgd._gbcea.Top, _adgd._gbcea.Bottom
}

// FitMode defines resizing options of an object inside a container.
type FitMode int

// SetColorTop sets border color for top.
func (_bbab *border) SetColorTop(col Color) { _bbab._dfd = col }

// SetPadding sets the padding of the component. The padding represents
// inner margins which are applied around the contents of the division.
// The background of the component is not affected by its padding.
func (_ggga *Division) SetPadding(left, right, top, bottom float64) {
	_ggga._gbga.Left = left
	_ggga._gbga.Right = right
	_ggga._gbga.Top = top
	_ggga._gbga.Bottom = bottom
}

// SetMargins sets the margins of the component. The margins are applied
// around the division.
func (_deaf *Division) SetMargins(left, right, top, bottom float64) {
	_deaf._ddfg.Left = left
	_deaf._ddfg.Right = right
	_deaf._ddfg.Top = top
	_deaf._ddfg.Bottom = bottom
}
func (_cebf *RadialShading) shadingModel() *_fgd.PdfShadingType3 {
	_ggbf, _fcdc, _bddfc := _cebf._dafd._feac.ToRGB()
	var _dfbf _gga.Point
	switch _cebf._effe {
	case AnchorBottomLeft:
		_dfbf = _gga.Point{X: _cebf._bfgf.Llx, Y: _cebf._bfgf.Lly}
	case AnchorBottomRight:
		_dfbf = _gga.Point{X: _cebf._bfgf.Urx, Y: _cebf._bfgf.Ury - _cebf._bfgf.Height()}
	case AnchorTopLeft:
		_dfbf = _gga.Point{X: _cebf._bfgf.Llx, Y: _cebf._bfgf.Lly + _cebf._bfgf.Height()}
	case AnchorTopRight:
		_dfbf = _gga.Point{X: _cebf._bfgf.Urx, Y: _cebf._bfgf.Ury}
	case AnchorLeft:
		_dfbf = _gga.Point{X: _cebf._bfgf.Llx, Y: _cebf._bfgf.Lly + _cebf._bfgf.Height()/2}
	case AnchorTop:
		_dfbf = _gga.Point{X: _cebf._bfgf.Llx + _cebf._bfgf.Width()/2, Y: _cebf._bfgf.Ury}
	case AnchorRight:
		_dfbf = _gga.Point{X: _cebf._bfgf.Urx, Y: _cebf._bfgf.Lly + _cebf._bfgf.Height()/2}
	case AnchorBottom:
		_dfbf = _gga.Point{X: _cebf._bfgf.Urx + _cebf._bfgf.Width()/2, Y: _cebf._bfgf.Lly}
	default:
		_dfbf = _gga.NewPoint(_cebf._bfgf.Llx+_cebf._bfgf.Width()/2, _cebf._bfgf.Lly+_cebf._bfgf.Height()/2)
	}
	_adggb := _cebf._aeeeb
	_accg := _cebf._babf
	_fegc := _dfbf.X + _cebf._eccda
	_eaage := _dfbf.Y + _cebf._dagc
	if _adggb == -1.0 {
		_adggb = 0.0
	}
	if _accg == -1.0 {
		var _facf []float64
		_ebgee := _fc.Pow(_fegc-_cebf._bfgf.Llx, 2) + _fc.Pow(_eaage-_cebf._bfgf.Lly, 2)
		_facf = append(_facf, _fc.Abs(_ebgee))
		_dgcg := _fc.Pow(_fegc-_cebf._bfgf.Llx, 2) + _fc.Pow(_cebf._bfgf.Lly+_cebf._bfgf.Height()-_eaage, 2)
		_facf = append(_facf, _fc.Abs(_dgcg))
		_dgffe := _fc.Pow(_cebf._bfgf.Urx-_fegc, 2) + _fc.Pow(_eaage-_cebf._bfgf.Ury-_cebf._bfgf.Height(), 2)
		_facf = append(_facf, _fc.Abs(_dgffe))
		_cdcda := _fc.Pow(_cebf._bfgf.Urx-_fegc, 2) + _fc.Pow(_cebf._bfgf.Ury-_eaage, 2)
		_facf = append(_facf, _fc.Abs(_cdcda))
		_fd.Slice(_facf, func(_dccbf, _becg int) bool { return _dccbf > _becg })
		_accg = _fc.Sqrt(_facf[0])
	}
	_cace := &_fgd.PdfRectangle{Llx: _fegc - _accg, Lly: _eaage - _accg, Urx: _fegc + _accg, Ury: _eaage + _accg}
	_dgcfd := _fgd.NewPdfShadingType3()
	_dgcfd.PdfShading.ShadingType = _bc.MakeInteger(3)
	_dgcfd.PdfShading.ColorSpace = _fgd.NewPdfColorspaceDeviceRGB()
	_dgcfd.PdfShading.Background = _bc.MakeArrayFromFloats([]float64{_ggbf, _fcdc, _bddfc})
	_dgcfd.PdfShading.BBox = _cace
	_dgcfd.PdfShading.AntiAlias = _bc.MakeBool(_cebf._dafd._ccea)
	_dgcfd.Coords = _bc.MakeArrayFromFloats([]float64{_fegc, _eaage, _adggb, _fegc, _eaage, _accg})
	_dgcfd.Domain = _bc.MakeArrayFromFloats([]float64{0.0, 1.0})
	_dgcfd.Extend = _bc.MakeArray(_bc.MakeBool(_cebf._dafd._fece[0]), _bc.MakeBool(_cebf._dafd._fece[1]))
	_dgcfd.Function = _cebf._dafd.generatePdfFunctions()
	return _dgcfd
}
func (_bcbc *Rectangle) applyFitMode(_cfbg float64) {
	_cfbg -= _bcbc._gagb.Left + _bcbc._gagb.Right + _bcbc._caeg
	switch _bcbc._aaab {
	case FitModeFillWidth:
		_bcbc.ScaleToWidth(_cfbg)
	}
}

// SetVerticalAlignment set the cell's vertical alignment of content.
// Can be one of:
// - CellHorizontalAlignmentTop
// - CellHorizontalAlignmentMiddle
// - CellHorizontalAlignmentBottom
func (_becff *TableCell) SetVerticalAlignment(valign CellVerticalAlignment) { _becff._bdfgg = valign }

// SetMaxLines sets the maximum number of lines before the paragraph
// text is truncated.
func (_ecfb *Paragraph) SetMaxLines(maxLines int) { _ecfb._dacad = maxLines; _ecfb.wrapText() }
func (_ffc *Block) addContentsByString(_ga string) error {
	_cda := _dg.NewContentStreamParser(_ga)
	_bca, _ade := _cda.Parse()
	if _ade != nil {
		return _ade
	}
	_ffc._cb.WrapIfNeeded()
	_bca.WrapIfNeeded()
	*_ffc._cb = append(*_ffc._cb, *_bca...)
	return nil
}

// SetShowNumbering sets a flag to indicate whether or not to show chapter numbers as part of title.
func (_bada *Chapter) SetShowNumbering(show bool) {
	_bada._aee = show
	_bada._gfdf.SetText(_bada.headingText())
}

// InvoiceAddress contains contact information that can be displayed
// in an invoice. It is used for the seller and buyer information in the
// invoice template.
type InvoiceAddress struct {
	Heading string
	Name    string
	Street  string
	Street2 string
	Zip     string
	City    string
	State   string
	Country string
	Phone   string
	Email   string

	// Separator defines the separator between different address components,
	// such as the city, state and zip code. It defaults to ", " when the
	// field is an empty string.
	Separator string

	// If enabled, the Phone field label (`Phone: `) is not displayed.
	HidePhoneLabel bool

	// If enabled, the Email field label (`Email: `) is not displayed.
	HideEmailLabel bool
}

// SetPositioning sets the positioning of the rectangle (absolute or relative).
func (_adda *Rectangle) SetPositioning(position Positioning) { _adda._dfff = position }

// GeneratePageBlocks draws the polygon on a new block representing the page.
// Implements the Drawable interface.
func (_cdca *Polygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cfdc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_dbeab, _dead := _cfdc.setOpacity(_cdca._afdf, _cdca._bgde)
	if _dead != nil {
		return nil, ctx, _dead
	}
	_gabf := _cdca._acdc
	_gabf.FillEnabled = _gabf.FillColor != nil
	_gabf.BorderEnabled = _gabf.BorderColor != nil && _gabf.BorderWidth > 0
	_dafaf := _gabf.Points
	_bedd := _fgd.PdfRectangle{}
	_ecfe := false
	for _fgfd := range _dafaf {
		for _baafe := range _dafaf[_fgfd] {
			_gbgac := &_dafaf[_fgfd][_baafe]
			_gbgac.Y = ctx.PageHeight - _gbgac.Y
			if !_ecfe {
				_bedd.Llx = _gbgac.X
				_bedd.Lly = _gbgac.Y
				_bedd.Urx = _gbgac.X
				_bedd.Ury = _gbgac.Y
				_ecfe = true
			} else {
				_bedd.Llx = _fc.Min(_bedd.Llx, _gbgac.X)
				_bedd.Lly = _fc.Min(_bedd.Lly, _gbgac.Y)
				_bedd.Urx = _fc.Max(_bedd.Urx, _gbgac.X)
				_bedd.Ury = _fc.Max(_bedd.Ury, _gbgac.Y)
			}
		}
	}
	if _gabf.FillEnabled {
		_gcge := _cacbe(_cfdc, _cdca._acdc.FillColor, _cdca._gcbdc, func() Rectangle {
			return Rectangle{_defb: _bedd.Llx, _cdgf: _bedd.Lly, _bcede: _bedd.Width(), _cggg: _bedd.Height()}
		})
		if _gcge != nil {
			return nil, ctx, _gcge
		}
	}
	_dgbg, _, _dead := _gabf.MarkedDraw(_dbeab, _cdca._fdbg)
	if _dead != nil {
		return nil, ctx, _dead
	}
	if _dead = _cfdc.addContentsByString(string(_dgbg)); _dead != nil {
		return nil, ctx, _dead
	}
	return []*Block{_cfdc}, ctx, nil
}

// ScaleToHeight scale Image to a specified height h, maintaining the aspect ratio.
func (_abdg *Image) ScaleToHeight(h float64) {
	_bgcfc := _abdg._ecfc / _abdg._fegb
	_abdg._fegb = h
	_abdg._ecfc = h * _bgcfc
}
func (_eeeag *templateProcessor) parseBackground(_geaa *templateNode) (interface{}, error) {
	_bfcgb := &Background{}
	for _, _dbeb := range _geaa._adfdg.Attr {
		_gaefc := _dbeb.Value
		switch _cacgg := _dbeb.Name.Local; _cacgg {
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_bfcgb.FillColor = _eeeag.parseColorAttr(_cacgg, _gaefc)
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_bfcgb.BorderColor = _eeeag.parseColorAttr(_cacgg, _gaefc)
		case "b\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0069\u007a\u0065":
			_bfcgb.BorderSize = _eeeag.parseFloatAttr(_cacgg, _gaefc)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_gedgd, _ecbcb, _bbebe, _fbea := _eeeag.parseBorderRadiusAttr(_cacgg, _gaefc)
			_bfcgb.SetBorderRadius(_gedgd, _ecbcb, _fbea, _bbebe)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_bfcgb.BorderRadiusTopLeft = _eeeag.parseFloatAttr(_cacgg, _gaefc)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_bfcgb.BorderRadiusTopRight = _eeeag.parseFloatAttr(_cacgg, _gaefc)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_bfcgb.BorderRadiusBottomLeft = _eeeag.parseFloatAttr(_cacgg, _gaefc)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_bfcgb.BorderRadiusBottomRight = _eeeag.parseFloatAttr(_cacgg, _gaefc)
		default:
			_eeeag.nodeLogDebug(_geaa, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _cacgg)
		}
	}
	return _bfcgb, nil
}

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_abfg *RadialShading) AddPatternResource(block *Block) (_ebade _bc.PdfObjectName, _eegg error) {
	_abag := 1
	_fdcc := _bc.PdfObjectName("\u0050" + _fg.Itoa(_abag))
	for block._dgd.HasPatternByName(_fdcc) {
		_abag++
		_fdcc = _bc.PdfObjectName("\u0050" + _fg.Itoa(_abag))
	}
	if _dagb := block._dgd.SetPatternByName(_fdcc, _abfg.ToPdfShadingPattern().ToPdfObject()); _dagb != nil {
		return "", _dagb
	}
	return _fdcc, nil
}

// AnchorPoint defines anchor point where the center position of the radial gradient would be calculated.
type AnchorPoint int

func (_baba *TOCLine) getLineLink() *_fgd.PdfAnnotation {
	if _baba._ebaab <= 0 {
		return nil
	}
	return _cbdce(_baba._ebaab-1, _baba._beed, _baba._cdbff, 0)
}

// SkipRows skips over a specified number of rows in the table.
func (_fbaba *Table) SkipRows(num int) {
	_bagd := num*_fbaba._afacb - 1
	if _bagd < 0 {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _eege := 0; _eege < _bagd; _eege++ {
		_fbaba.NewCell()
	}
}

// NewEllipse creates a new ellipse with the center at (`xc`, `yc`),
// having the specified width and height.
// NOTE: In relative positioning mode, `xc` and `yc` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the ellipse is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_gcdc *Creator) NewEllipse(xc, yc, width, height float64) *Ellipse {
	return _becb(xc, yc, width, height)
}
func _bcfa(_gggc *templateProcessor, _dgadf *templateNode) (interface{}, error) {
	return _gggc.parseRectangle(_dgadf)
}

// FitMode returns the fit mode of the ellipse.
func (_dffed *Ellipse) FitMode() FitMode { return _dffed._eddea }

// SetPageSize sets the Creator's page size.  Pages that are added after this will be created with
// this Page size.
// Does not affect pages already created.
//
// Common page sizes are defined as constants.
// Examples:
// 1. c.SetPageSize(creator.PageSizeA4)
// 2. c.SetPageSize(creator.PageSizeA3)
// 3. c.SetPageSize(creator.PageSizeLegal)
// 4. c.SetPageSize(creator.PageSizeLetter)
//
// For custom sizes: Use the PPMM (points per mm) and PPI (points per inch) when defining those based on
// physical page sizes:
//
// Examples:
// 1. 10x15 sq. mm: SetPageSize(PageSize{10*creator.PPMM, 15*creator.PPMM}) where PPMM is points per mm.
// 2. 3x2 sq. inches: SetPageSize(PageSize{3*creator.PPI, 2*creator.PPI}) where PPI is points per inch.
func (_cdbf *Creator) SetPageSize(size PageSize) {
	_cdbf._ccae = size
	_cdbf._gbgg = size[0]
	_cdbf._agfdf = size[1]
	_gbee := 0.1 * _cdbf._gbgg
	_cdbf._cbb.Left = _gbee
	_cdbf._cbb.Right = _gbee
	_cdbf._cbb.Top = _gbee
	_cdbf._cbb.Bottom = _gbee
}

// SkipOver skips over a specified number of rows and cols.
func (_dbgdcg *Table) SkipOver(rows, cols int) {
	_aabb := rows*_dbgdcg._afacb + cols - 1
	if _aabb < 0 {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _dcae := 0; _dcae < _aabb; _dcae++ {
		_dbgdcg.NewCell()
	}
}

// Title returns the title of the invoice.
func (_fabf *Invoice) Title() string { return _fabf._dgfaa }

// Reset removes all the text chunks the paragraph contains.
func (_agae *StyledParagraph) Reset() { _agae._ecec = []*TextChunk{} }

// FrontpageFunctionArgs holds the input arguments to a front page drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FrontpageFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// SetWidth sets the width of the ellipse.
func (_eegcb *Ellipse) SetWidth(width float64) { _eegcb._eefg = width }
func (_abcc *StyledParagraph) wrapText() error { return _abcc.wrapChunks(true) }
func _gbg(_ccb *_dg.ContentStreamOperations, _cbc *_fgd.PdfPageResources, _geb *_dg.ContentStreamOperations, _bfag *_fgd.PdfPageResources) error {
	_edb := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	_afda := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	_fabd := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	_bfb := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	_fae := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	_dgde := map[_bc.PdfObjectName]_bc.PdfObjectName{}
	for _, _cce := range *_geb {
		switch _cce.Operand {
		case "\u0044\u006f":
			if len(_cce.Params) == 1 {
				if _egf, _agb := _cce.Params[0].(*_bc.PdfObjectName); _agb {
					if _, _gcd := _edb[*_egf]; !_gcd {
						var _edd _bc.PdfObjectName
						_acc, _ := _bfag.GetXObjectByName(*_egf)
						if _acc != nil {
							_edd = *_egf
							for {
								_db, _ := _cbc.GetXObjectByName(_edd)
								if _db == nil || _db == _acc {
									break
								}
								_edd = *_bc.MakeName(_fbe(_edd.String()))
							}
						}
						_cbc.SetXObjectByName(_edd, _acc)
						_edb[*_egf] = _edd
					}
					_bdf := _edb[*_egf]
					_cce.Params[0] = &_bdf
				}
			}
		case "\u0054\u0066":
			if len(_cce.Params) == 2 {
				if _dfec, _ffcf := _cce.Params[0].(*_bc.PdfObjectName); _ffcf {
					if _, _bbb := _afda[*_dfec]; !_bbb {
						_ba, _ced := _bfag.GetFontByName(*_dfec)
						_cgg := *_dfec
						if _ced && _ba != nil {
							_cgg = _aef(_dfec.String(), _ba, _cbc)
						}
						_cbc.SetFontByName(_cgg, _ba)
						_afda[*_dfec] = _cgg
					}
					_bgf := _afda[*_dfec]
					_cce.Params[0] = &_bgf
				}
			}
		case "\u0043\u0053", "\u0063\u0073":
			if len(_cce.Params) == 1 {
				if _aded, _aaf := _cce.Params[0].(*_bc.PdfObjectName); _aaf {
					if _, _afc := _fabd[*_aded]; !_afc {
						var _fgg _bc.PdfObjectName
						_gde, _bee := _bfag.GetColorspaceByName(*_aded)
						if _bee {
							_fgg = *_aded
							for {
								_dc, _cca := _cbc.GetColorspaceByName(_fgg)
								if !_cca || _gde == _dc {
									break
								}
								_fgg = *_bc.MakeName(_fbe(_fgg.String()))
							}
							_cbc.SetColorspaceByName(_fgg, _gde)
							_fabd[*_aded] = _fgg
						} else {
							_fec.Log.Debug("C\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064")
						}
					}
					if _agf, _dbc := _fabd[*_aded]; _dbc {
						_cce.Params[0] = &_agf
					} else {
						_fec.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", *_aded)
					}
				}
			}
		case "\u0053\u0043\u004e", "\u0073\u0063\u006e":
			if len(_cce.Params) == 1 {
				if _eaa, _gee := _cce.Params[0].(*_bc.PdfObjectName); _gee {
					if _, _gbcf := _bfb[*_eaa]; !_gbcf {
						var _aad _bc.PdfObjectName
						_abd, _gef := _bfag.GetPatternByName(*_eaa)
						if _gef {
							_aad = *_eaa
							for {
								_fdf, _ecde := _cbc.GetPatternByName(_aad)
								if !_ecde || _fdf == _abd {
									break
								}
								_aad = *_bc.MakeName(_fbe(_aad.String()))
							}
							_ded := _cbc.SetPatternByName(_aad, _abd.ToPdfObject())
							if _ded != nil {
								return _ded
							}
							_bfb[*_eaa] = _aad
						}
					}
					if _fde, _ccg := _bfb[*_eaa]; _ccg {
						_cce.Params[0] = &_fde
					}
				}
			}
		case "\u0073\u0068":
			if len(_cce.Params) == 1 {
				if _debd, _agca := _cce.Params[0].(*_bc.PdfObjectName); _agca {
					if _, _eaag := _fae[*_debd]; !_eaag {
						var _fee _bc.PdfObjectName
						_bgd, _fbc := _bfag.GetShadingByName(*_debd)
						if _fbc {
							_fee = *_debd
							for {
								_geec, _ffcc := _cbc.GetShadingByName(_fee)
								if !_ffcc || _bgd == _geec {
									break
								}
								_fee = *_bc.MakeName(_fbe(_fee.String()))
							}
							_fgcf := _cbc.SetShadingByName(_fee, _bgd.ToPdfObject())
							if _fgcf != nil {
								_fec.Log.Debug("E\u0052\u0052\u004f\u0052 S\u0065t\u0020\u0073\u0068\u0061\u0064i\u006e\u0067\u003a\u0020\u0025\u0076", _fgcf)
								return _fgcf
							}
							_fae[*_debd] = _fee
						} else {
							_fec.Log.Debug("\u0053\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
						}
					}
					if _bcac, _gaa := _fae[*_debd]; _gaa {
						_cce.Params[0] = &_bcac
					} else {
						_fec.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020S\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0073 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064", *_debd)
					}
				}
			}
		case "\u0067\u0073":
			if len(_cce.Params) == 1 {
				if _bed, _aae := _cce.Params[0].(*_bc.PdfObjectName); _aae {
					if _, _ffb := _dgde[*_bed]; !_ffb {
						var _dcg _bc.PdfObjectName
						_eab, _gdf := _bfag.GetExtGState(*_bed)
						if _gdf {
							_dcg = *_bed
							for {
								_dgf, _age := _cbc.GetExtGState(_dcg)
								if !_age || _eab == _dgf {
									break
								}
								_dcg = *_bc.MakeName(_fbe(_dcg.String()))
							}
						}
						_cbc.AddExtGState(_dcg, _eab)
						_dgde[*_bed] = _dcg
					}
					_fgce := _dgde[*_bed]
					_cce.Params[0] = &_fgce
				}
			}
		}
		*_ccb = append(*_ccb, _cce)
	}
	return nil
}

// ScaleToWidth scales the Block to a specified width, maintaining the same aspect ratio.
func (_bef *Block) ScaleToWidth(w float64) { _gbd := w / _bef._da; _bef.Scale(_gbd, _gbd) }

// Margins returns the margins of the list: left, right, top, bottom.
func (_geab *List) Margins() (float64, float64, float64, float64) {
	return _geab._gacb.Left, _geab._gacb.Right, _geab._gacb.Top, _geab._gacb.Bottom
}
func (_dcab *Table) sortCells() {
	_fd.Slice(_dcab._efbe, func(_gedbb, _fbgf int) bool {
		_fbcd := _dcab._efbe[_gedbb]._deef
		_effec := _dcab._efbe[_fbgf]._deef
		if _fbcd < _effec {
			return true
		}
		if _fbcd > _effec {
			return false
		}
		return _dcab._efbe[_gedbb]._bacg < _dcab._efbe[_fbgf]._bacg
	})
}

// SetLazy sets the lazy mode for the image.
func (_abba *Image) SetLazy(lazy bool) { _abba._faeda = lazy }

// NewGraphicSVGFromString creates a graphic SVG from a SVG string.
func NewGraphicSVGFromString(svgStr string) (*GraphicSVG, error) { return _bcgf(svgStr) }
func (_gcad *templateProcessor) parseBoolAttr(_fdddd, _edage string) bool {
	_fec.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0062\u006f\u006f\u006c\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _fdddd, _edage)
	_badcd, _ := _fg.ParseBool(_edage)
	return _edage == "" || _badcd
}

// HeaderFunctionArgs holds the input arguments to a header drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type HeaderFunctionArgs struct {
	PageNum    int
	TotalPages int
}

func (_fcgb *templateProcessor) addNodeText(_cdbb *templateNode, _agaac string) error {
	_ecadg := _cdbb._bedcd
	if _ecadg == nil {
		return nil
	}
	switch _fdbcc := _ecadg.(type) {
	case *TextChunk:
		_fdbcc.Text = _agaac
	case *Paragraph:
		switch _cdbb._adfdg.Name.Local {
		case "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067":
			if _cdbb._cefd != nil {
				if _eagdc, _afbg := _cdbb._cefd._bedcd.(*Chapter); _afbg {
					_eagdc._faed = _agaac
					_fdbcc.SetText(_eagdc.headingText())
				}
			}
		default:
			_fdbcc.SetText(_agaac)
		}
	}
	return nil
}

// Add appends a new item to the list.
// The supported components are: *Paragraph, *StyledParagraph, *Division, *Image, *Table, and *List.
// Returns the marker used for the newly added item. The returned marker
// object can be used to change the text and style of the marker for the
// current item.
func (_eagga *List) Add(item VectorDrawable) (*TextChunk, error) {
	_ffaf := &listItem{_cdeda: item, _edee: _eagga._edaa}
	switch _fadb := item.(type) {
	case *Paragraph:
	case *StyledParagraph:
	case *List:
		if _fadb._ebfa {
			_fadb._aedbb = 15
		}
	case *Division:
	case *Image:
	case *Table:
	default:
		return nil, _bd.New("\u0074\u0068i\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u006c\u0069\u0073\u0074")
	}
	_eagga._ffegf = append(_eagga._ffegf, _ffaf)
	return &_ffaf._edee, nil
}

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_dgbeb *RadialShading) SetExtends(start bool, end bool) { _dgbeb._dafd.SetExtends(start, end) }

// Chart represents a chart drawable.
// It is used to render unichart chart components using a creator instance.
type Chart struct {
	_gff  _cc.ChartRenderable
	_cggb Positioning
	_dbf  float64
	_gbcb float64
	_bgfg Margins
	_aed  *int64
}

// Marker returns the marker used for the list items.
// The marker instance can be used the change the text and the style
// of newly added list items.
func (_gecc *List) Marker() *TextChunk { return &_gecc._edaa }

// SetWidth sets the width of the rectangle.
func (_cgceg *Rectangle) SetWidth(width float64) { _cgceg._bcede = width }

// SetOpacity sets opacity for Image.
func (_eeff *Image) SetOpacity(opacity float64) { _eeff._ddea = opacity }

// NewRectangle creates a new rectangle with the left corner at (`x`, `y`),
// having the specified width and height.
// NOTE: In relative positioning mode, `x` and `y` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the rectangle is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_fdeb *Creator) NewRectangle(x, y, width, height float64) *Rectangle {
	return _aggd(x, y, width, height)
}

// Columns returns all the columns in the invoice line items table.
func (_gbea *Invoice) Columns() []*InvoiceCell { return _gbea._gffd }

// ColorGrayFrom8bit creates a Color from 8-bit (0-255) gray values.
// Example:
//
//	gray := ColorGrayFrom8bit(255, 0, 0)
func ColorGrayFrom8bit(g byte) Color { return grayColor{float64(g) / 255.0} }

const (
	TextOverflowVisible TextOverflow = iota
	TextOverflowHidden
)

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_cegd *StyledParagraph) GetMargins() (float64, float64, float64, float64) {
	return _cegd._ccddg.Left, _cegd._ccddg.Right, _cegd._ccddg.Top, _cegd._ccddg.Bottom
}
func _eddeb(_ebgc *_fgd.Image) (*Image, error) {
	_efdb := float64(_ebgc.Width)
	_dacg := float64(_ebgc.Height)
	return &Image{_cgdfb: _ebgc, _bgcad: _efdb, _bgff: _dacg, _ecfc: _efdb, _fegb: _dacg, _fbbf: 0, _ddea: 1.0, _accd: PositionRelative}, nil
}

// SetColumnWidths sets the fractional column widths.
// Each width should be in the range 0-1 and is a fraction of the table width.
// The number of width inputs must match number of columns, otherwise an error is returned.
func (_ecddc *Table) SetColumnWidths(widths ...float64) error {
	if len(widths) != _ecddc._afacb {
		_fec.Log.Debug("M\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0077\u0069\u0064\u0074\u0068\u0073\u0020\u0061nd\u0020\u0063\u006fl\u0075m\u006e\u0073")
		return _bd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_ecddc._cddfg = widths
	return nil
}

// SetHorizontalAlignment sets the cell's horizontal alignment of content.
// Can be one of:
// - CellHorizontalAlignmentLeft
// - CellHorizontalAlignmentCenter
// - CellHorizontalAlignmentRight
func (_ddfcb *TableCell) SetHorizontalAlignment(halign CellHorizontalAlignment) {
	_ddfcb._bbbff = halign
}

// GeneratePageBlocks draws the chart onto a block.
func (_cgec *Chart) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gdc := ctx
	_gfag := _cgec._cggb.IsRelative()
	var _dcb []*Block
	if _gfag {
		_fdg := 1.0
		_dccf := _cgec._bgfg.Top
		if float64(_cgec._gff.Height()) > ctx.Height-_cgec._bgfg.Top {
			_dcb = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _gge error
			if _, ctx, _gge = _gbcfe().GeneratePageBlocks(ctx); _gge != nil {
				return nil, ctx, _gge
			}
			_dccf = 0
		}
		ctx.X += _cgec._bgfg.Left + _fdg
		ctx.Y += _dccf
		ctx.Width -= _cgec._bgfg.Left + _cgec._bgfg.Right + 2*_fdg
		ctx.Height -= _dccf
		_cgec._gff.SetWidth(int(ctx.Width))
	} else {
		ctx.X = _cgec._dbf
		ctx.Y = _cgec._gbcb
	}
	_dgfg := _dg.NewContentCreator()
	if _cgec._aed != nil {
		_dgfg.Add_BDC(*_bc.MakeName(_fgd.StructureTypeFigure), map[string]_bc.PdfObject{"\u004d\u0043\u0049\u0044": _bc.MakeInteger(*_cgec._aed)})
	}
	_dgfg.Translate(0, ctx.PageHeight)
	_dgfg.Scale(1, -1)
	_dgfg.Translate(ctx.X, ctx.Y)
	_eedf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_cgec._gff.Render(_ce.NewRenderer(_dgfg, _eedf._dgd), nil)
	if _cgec._aed != nil {
		_dgfg.Add_EMC()
	}
	if _aaeb := _eedf.addContentsByString(_dgfg.String()); _aaeb != nil {
		return nil, ctx, _aaeb
	}
	if _gfag {
		_afga := _cgec.Height() + _cgec._bgfg.Bottom
		ctx.Y += _afga
		ctx.Height -= _afga
	} else {
		ctx = _gdc
	}
	_dcb = append(_dcb, _eedf)
	return _dcb, ctx, nil
}
func (_adfeg *TableCell) cloneProps(_bbdbd VectorDrawable) *TableCell {
	_fgfe := *_adfeg
	_fgfe._aafg = _bbdbd
	return &_fgfe
}

const (
	CellVerticalAlignmentTop CellVerticalAlignment = iota
	CellVerticalAlignmentMiddle
	CellVerticalAlignmentBottom
)

func (_cebe *Paragraph) getTextMetrics() (_bfcg, _cfcg, _ggdab float64) {
	_bcdgb := _aggdb(_cebe._gbca, _cebe._acdge)
	if _bcdgb._geede > _bfcg {
		_bfcg = _bcdgb._geede
	}
	if _bcdgb._egeb < _ggdab {
		_ggdab = _bcdgb._egeb
	}
	if _bcfg := _cebe._acdge; _bcfg > _cfcg {
		_cfcg = _bcfg
	}
	return _bfcg, _cfcg, _ggdab
}

// SetMarkedContentID sets the marked content ID for the paragraph.
func (_dded *StyledParagraph) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_dded._eeba = &mcid
	_gabfb := _fgd.NewKDictionary()
	_gabfb.S = _bc.MakeName("\u0050")
	_gabfb.K = _bc.MakeInteger(mcid)
	return _gabfb
}

// FillColor returns the fill color of the ellipse.
func (_ddce *Ellipse) FillColor() Color { return _ddce._baacc }
func (_fgag *List) ctxHeight(_gdcga float64) float64 {
	_gdcga -= _fgag._aedbb
	var _fbac float64
	for _, _adad := range _fgag._ffegf {
		_fbac += _adad.ctxHeight(_gdcga)
	}
	return _fbac
}

// NewColumn returns a new column for the line items invoice table.
func (_adfaa *Invoice) NewColumn(description string) *InvoiceCell {
	return _adfaa.newColumn(description, CellHorizontalAlignmentLeft)
}

// Text sets the text content of the Paragraph.
func (_cafe *Paragraph) Text() string { return _cafe._egea }
func (_fgge *Table) updateRowHeights(_dbcfa float64) {
	for _, _cacea := range _fgge._efbe {
		_fbabb := _cacea.width(_fgge._cddfg, _dbcfa)
		_gabbb := _cacea.height(_fbabb)
		_gcccg := _fgge._gecdg[_cacea._deef+_cacea._afddb-2]
		if _cacea._afddb > 1 {
			_abbfd := 0.0
			_aggge := _fgge._gecdg[_cacea._deef-1 : (_cacea._deef + _cacea._afddb - 1)]
			for _, _bgdgf := range _aggge {
				_abbfd += _bgdgf
			}
			if _gabbb <= _abbfd {
				continue
			}
		}
		if _gabbb > _gcccg {
			_efeg := _gabbb / float64(_cacea._afddb)
			if _efeg > _gcccg {
				for _acbed := 1; _acbed <= _cacea._afddb; _acbed++ {
					if _efeg > _fgge._gecdg[_cacea._deef+_acbed-2] {
						_fgge._gecdg[_cacea._deef+_acbed-2] = _efeg
					}
				}
			}
		}
	}
}

// DrawFooter sets a function to draw a footer on created output pages.
func (_dae *Creator) DrawFooter(drawFooterFunc func(_cgb *Block, _beda FooterFunctionArgs)) {
	_dae._aedb = drawFooterFunc
}

// SetLevel sets the indentation level of the TOC line.
func (_addfcg *TOCLine) SetLevel(level uint) {
	_addfcg._gfec = level
	_addfcg._afgdd._ccddg.Left = _addfcg._afage + float64(_addfcg._gfec-1)*_addfcg._ebgcdc
}
func _cgff(_ccac float64, _aedd float64, _dfacb float64, _dddfb float64, _bcbebd []*ColorPoint) *RadialShading {
	return &RadialShading{_dafd: &shading{_feac: ColorWhite, _ccea: false, _fece: []bool{false, false}, _bcaca: _bcbebd}, _eccda: _ccac, _dagc: _aedd, _aeeeb: _dfacb, _babf: _dddfb, _effe: AnchorCenter}
}

// Height returns the height of the ellipse.
func (_cgcb *Ellipse) Height() float64 { return _cgcb._beb }

// TextOverflow determines the behavior of paragraph text which does
// not fit in the available space.
type TextOverflow int

// GeneratePageBlocks draws the composite Bezier curve on a new block
// representing the page. Implements the Drawable interface.
func (_afdad *PolyBezierCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dgfc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ggefa, _eccg := _dgfc.setOpacity(_afdad._ffafb, _afdad._addf)
	if _eccg != nil {
		return nil, ctx, _eccg
	}
	_cabd := _afdad._ffcd
	_cabd.FillEnabled = _cabd.FillColor != nil
	var (
		_fagc = ctx.PageHeight
		_fdgd = _cabd.Curves
		_ddbe = make([]_gga.CubicBezierCurve, 0, len(_cabd.Curves))
	)
	_dffa := _fgd.PdfRectangle{}
	for _becef := range _cabd.Curves {
		_acbe := _fdgd[_becef]
		_acbe.P0.Y = _fagc - _acbe.P0.Y
		_acbe.P1.Y = _fagc - _acbe.P1.Y
		_acbe.P2.Y = _fagc - _acbe.P2.Y
		_acbe.P3.Y = _fagc - _acbe.P3.Y
		_ddbe = append(_ddbe, _acbe)
		_aabg := _acbe.GetBounds()
		if _becef == 0 {
			_dffa = _aabg
		} else {
			_dffa.Llx = _fc.Min(_dffa.Llx, _aabg.Llx)
			_dffa.Lly = _fc.Min(_dffa.Lly, _aabg.Lly)
			_dffa.Urx = _fc.Max(_dffa.Urx, _aabg.Urx)
			_dffa.Ury = _fc.Max(_dffa.Ury, _aabg.Ury)
		}
	}
	_cabd.Curves = _ddbe
	defer func() { _cabd.Curves = _fdgd }()
	if _cabd.FillEnabled {
		_gacdc := _cacbe(_dgfc, _afdad._ffcd.FillColor, _afdad._bbaef, func() Rectangle {
			return Rectangle{_defb: _dffa.Llx, _cdgf: _dffa.Lly, _bcede: _dffa.Width(), _cggg: _dffa.Height()}
		})
		if _gacdc != nil {
			return nil, ctx, _gacdc
		}
	}
	_abgfg, _, _eccg := _cabd.MarkedDraw(_ggefa, _afdad._daef)
	if _eccg != nil {
		return nil, ctx, _eccg
	}
	if _eccg = _dgfc.addContentsByString(string(_abgfg)); _eccg != nil {
		return nil, ctx, _eccg
	}
	return []*Block{_dgfc}, ctx, nil
}

// SetLineWidth sets the line width.
func (_aebcb *Line) SetLineWidth(width float64) { _aebcb._aebf = width }

// GeneratePageBlocks draws the block contents on a template Page block.
// Implements the Drawable interface.
func (_fcd *Block) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ag := _de.IdentityMatrix()
	_afd, _dab := _fcd.Width(), _fcd.Height()
	if _fcd._ad.IsRelative() {
		_ag = _ag.Translate(ctx.X, ctx.PageHeight-ctx.Y-_dab)
	} else {
		_ag = _ag.Translate(_fcd._bf, ctx.PageHeight-_fcd._cbe-_dab)
	}
	_acecc := _dab
	if _fcd._df != 0 {
		_ag = _ag.Translate(_afd/2, _dab/2).Rotate(_fcd._df*_fc.Pi/180.0).Translate(-_afd/2, -_dab/2)
		_, _acecc = _fcd.RotatedSize()
	}
	if _fcd._ad.IsRelative() {
		ctx.Y += _acecc
	}
	_ea := _dg.NewContentCreator()
	_ea.Add_cm(_ag[0], _ag[1], _ag[3], _ag[4], _ag[6], _ag[7])
	_egd := _fcd.duplicate()
	_bg := append(*_ea.Operations(), *_egd._cb...)
	_bg.WrapIfNeeded()
	_egd._cb = &_bg
	for _, _efb := range _fcd._ed {
		_fff, _bcce := _bc.GetArray(_efb.Rect)
		if !_bcce || _fff.Len() != 4 {
			_fec.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0069\u006e\u0076\u0061\u006ci\u0064 \u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065l\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _efb.Rect)
			continue
		}
		_agc, _ffe := _fgd.NewPdfRectangle(*_fff)
		if _ffe != nil {
			_fec.Log.Debug("\u0057A\u0052N\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0070\u0061\u0072\u0073e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002e", _ffe)
			continue
		}
		_agc.Transform(_ag)
		_efb.Rect = _agc.ToPdfObject()
	}
	return []*Block{_egd}, ctx, nil
}

// SetPositioning sets the positioning of the line (absolute or relative).
func (_bcde *Line) SetPositioning(positioning Positioning) { _bcde._babe = positioning }
func (_dddb *Image) applyFitMode(_adbe float64) {
	_adbe -= _dddb._cdbc.Left + _dddb._cdbc.Right
	switch _dddb._cba {
	case FitModeFillWidth:
		_dddb.ScaleToWidth(_adbe)
	}
}
func (_geefb *templateProcessor) parseCellBorderStyleAttr(_gdcf, _gfaed string) CellBorderStyle {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020c\u0065\u006c\u006c b\u006f\u0072\u0064\u0065\u0072\u0020s\u0074\u0079\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025s\u0029\u002e", _gdcf, _gfaed)
	_acda := map[string]CellBorderStyle{"\u006e\u006f\u006e\u0065": CellBorderStyleNone, "\u0073\u0069\u006e\u0067\u006c\u0065": CellBorderStyleSingle, "\u0064\u006f\u0075\u0062\u006c\u0065": CellBorderStyleDouble}[_gfaed]
	return _acda
}

// LineWidth returns the width of the line.
func (_gbcbb *Line) LineWidth() float64 { return _gbcbb._aebf }

// DrawHeader sets a function to draw a header on created output pages.
func (_eagb *Creator) DrawHeader(drawHeaderFunc func(_gdca *Block, _dafe HeaderFunctionArgs)) {
	_eagb._cgf = drawHeaderFunc
}

// SetSideBorderStyle sets the cell's side border style.
func (_gagg *TableCell) SetSideBorderStyle(side CellBorderSide, style CellBorderStyle) {
	switch side {
	case CellBorderSideAll:
		_gagg._bgcce = style
		_gagg._bggf = style
		_gagg._faab = style
		_gagg._cfdbdc = style
	case CellBorderSideTop:
		_gagg._bgcce = style
	case CellBorderSideBottom:
		_gagg._bggf = style
	case CellBorderSideLeft:
		_gagg._faab = style
	case CellBorderSideRight:
		_gagg._cfdbdc = style
	}
}

// SetWidth set the Image's document width to specified w. This does not change the raw image data, i.e.
// no actual scaling of data is performed. That is handled by the PDF viewer.
func (_ccbdca *Image) SetWidth(w float64) { _ccbdca._ecfc = w }

// ScaleToWidth scale Image to a specified width w, maintaining the aspect ratio.
func (_ggca *Image) ScaleToWidth(w float64) {
	_bfcad := _ggca._fegb / _ggca._ecfc
	_ggca._ecfc = w
	_ggca._fegb = w * _bfcad
}

// ColorGrayFromHex converts color hex code to gray color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ff -> white.
func ColorGrayFromHex(hexStr string) Color {
	_fge := grayColor{}
	if (len(hexStr) != 2 && len(hexStr) != 3) || hexStr[0] != '#' {
		_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _fge
	}
	var _edga int
	if len(hexStr) == 2 {
		var _cga int
		_fcg, _caf := _f.Sscanf(hexStr, "\u0023\u0025\u0031\u0078", &_cga)
		if _caf != nil {
			_fec.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _caf)
			return _fge
		}
		if _fcg != 1 {
			_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fge
		}
		_edga = _cga*16 + _cga
	} else {
		_dda, _cafa := _f.Sscanf(hexStr, "\u0023\u0025\u0032\u0078", &_edga)
		if _cafa != nil {
			_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fge
		}
		if _dda != 1 {
			_fec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _dda)
			return _fge
		}
	}
	_fge._dbdf = float64(_edga) / 255.0
	return _fge
}

// NewTextChunk returns a new text chunk instance.
func NewTextChunk(text string, style TextStyle) *TextChunk {
	return &TextChunk{Text: text, Style: style, VerticalAlignment: TextVerticalAlignmentBaseline}
}

// Polyline represents a slice of points that are connected as straight lines.
// Implements the Drawable interface and can be rendered using the Creator.
type Polyline struct {
	_cdggc *_gga.Polyline
	_ggaf  float64
	_egfd  *int64
}

// TotalLines returns all the rows in the invoice totals table as
// description-value cell pairs.
func (_fgaa *Invoice) TotalLines() [][2]*InvoiceCell {
	_ggbdb := [][2]*InvoiceCell{_fgaa._beff}
	_ggbdb = append(_ggbdb, _fgaa._aegd...)
	return append(_ggbdb, _fgaa._gdba)
}

// IsRelative checks if the positioning is relative.
func (_gbe Positioning) IsRelative() bool { return _gbe == PositionRelative }

// Add adds a new Drawable to the chapter.
// Currently supported Drawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Table
// - *Division
// - *List
// - *Rectangle
// - *Ellipse
// - *Line
// - *Block,
// - *PageBreak
// - *Chapter
func (_aadb *Chapter) Add(d Drawable) error {
	if Drawable(_aadb) == d {
		_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0061\u0064\u0064\u0020\u0069\u0074\u0073\u0065\u006c\u0066")
		return _bd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	switch _becc := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line, *Block, *PageBreak, *Chapter:
		_aadb._agbc = append(_aadb._agbc, d)
	case containerDrawable:
		_ecdd, _fagb := _becc.ContainerComponent(_aadb)
		if _fagb != nil {
			return _fagb
		}
		_aadb._agbc = append(_aadb._agbc, _ecdd)
	default:
		_fec.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u003a\u0020\u0025\u0054", d)
		return _bd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return nil
}
func _gdbda(_feggg *_fgd.PdfRectangle, _beagd _de.Matrix) *_fgd.PdfRectangle {
	var _gedbf _fgd.PdfRectangle
	_gedbf.Llx, _gedbf.Lly = _beagd.Transform(_feggg.Llx, _feggg.Lly)
	_gedbf.Urx, _gedbf.Ury = _beagd.Transform(_feggg.Urx, _feggg.Ury)
	_gedbf.Normalize()
	return &_gedbf
}
func (_ebgaa *templateProcessor) parseAttrPropList(_ecbe string) map[string]string {
	_eegbe := _eg.Fields(_ecbe)
	if len(_eegbe) == 0 {
		return nil
	}
	_aaca := map[string]string{}
	for _, _ccee := range _eegbe {
		_fdbea := _bedfe.FindStringSubmatch(_ccee)
		if len(_fdbea) < 3 {
			continue
		}
		_acgbf, _bbdbae := _eg.TrimSpace(_fdbea[1]), _fdbea[2]
		if _acgbf == "" {
			continue
		}
		_aaca[_acgbf] = _bbdbae
	}
	return _aaca
}

// SetMarkedContentID sets marked content ID.
func (_agga *TOC) SetMarkedContentID(mcid int64) *_fgd.KDict { return nil }

// InsertColumn inserts a column in the line items table at the specified index.
func (_cdga *Invoice) InsertColumn(index uint, description string) *InvoiceCell {
	_bafbf := uint(len(_cdga._gffd))
	if index > _bafbf {
		index = _bafbf
	}
	_ggbd := _cdga.NewColumn(description)
	_cdga._gffd = append(_cdga._gffd[:index], append([]*InvoiceCell{_ggbd}, _cdga._gffd[index:]...)...)
	return _ggbd
}
func (_fffdb *shading) generatePdfFunctions() []_fgd.PdfFunction {
	if len(_fffdb._bcaca) == 0 {
		return nil
	} else if len(_fffdb._bcaca) <= 2 {
		_cdcd, _daecd, _bbece := _fffdb._bcaca[0]._afddd.ToRGB()
		_gbceg, _cbgb, _bcaa := _fffdb._bcaca[len(_fffdb._bcaca)-1]._afddd.ToRGB()
		return []_fgd.PdfFunction{&_fgd.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_cdcd, _daecd, _bbece}, C1: []float64{_gbceg, _cbgb, _bcaa}}}
	} else {
		_cfgg := []_fgd.PdfFunction{}
		_adeba := []float64{}
		for _gdgd := 0; _gdgd < len(_fffdb._bcaca)-1; _gdgd++ {
			_gdgc, _efddg, _fdeg := _fffdb._bcaca[_gdgd]._afddd.ToRGB()
			_dcea, _dbgfd, _gabe := _fffdb._bcaca[_gdgd+1]._afddd.ToRGB()
			_fdegb := &_fgd.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_gdgc, _efddg, _fdeg}, C1: []float64{_dcea, _dbgfd, _gabe}}
			_cfgg = append(_cfgg, _fdegb)
			if _gdgd > 0 {
				_adeba = append(_adeba, _fffdb._bcaca[_gdgd]._ffag)
			}
		}
		_aeaf := []float64{}
		for range _cfgg {
			_aeaf = append(_aeaf, []float64{0.0, 1.0}...)
		}
		return []_fgd.PdfFunction{&_fgd.PdfFunctionType3{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, Functions: _cfgg, Bounds: _adeba, Encode: _aeaf}}
	}
}

// Block contains a portion of PDF Page contents. It has a width and a position and can
// be placed anywhere on a Page.  It can even contain a whole Page, and is used in the creator
// where each Drawable object can output one or more blocks, each representing content for separate pages
// (typically needed when Page breaks occur).
type Block struct {
	_cb       *_dg.ContentStreamOperations
	_dgd      *_fgd.PdfPageResources
	_ad       Positioning
	_bf, _cbe float64
	_da       float64
	_ace      float64
	_df       float64
	_aag      Margins
	_ed       []*_fgd.PdfAnnotation
}

// SetWidth sets line width.
func (_bfbg *Curve) SetWidth(width float64) { _bfbg._befg = width }

// SetSubtotal sets the subtotal of the invoice.
func (_ggadf *Invoice) SetSubtotal(value string) { _ggadf._beff[1].Value = value }
func (_dage *Invoice) newColumn(_ddeg string, _acbb CellHorizontalAlignment) *InvoiceCell {
	_ddbad := &InvoiceCell{_dage._gbfg, _ddeg}
	_ddbad.Alignment = _acbb
	return _ddbad
}

// NewCell makes a new cell and inserts it into the table at the current position.
func (_eggab *Table) NewCell() *TableCell { return _eggab.MultiCell(1, 1) }

// NewPage adds a new Page to the Creator and sets as the active Page.
func (_dbeg *Creator) NewPage() *_fgd.PdfPage {
	_acea := _dbeg.newPage()
	_dbeg._egfb = append(_dbeg._egfb, _acea)
	_dbeg._eee.Page++
	return _acea
}

// ColorRGBFrom8bit creates a Color from 8-bit (0-255) r,g,b values.
// Example:
//
//	red := ColorRGBFrom8Bit(255, 0, 0)
func ColorRGBFrom8bit(r, g, b byte) Color {
	return rgbColor{_ead: float64(r) / 255.0, _ddd: float64(g) / 255.0, _decg: float64(b) / 255.0}
}

// SetBorderWidth sets the border width of the ellipse.
func (_geeb *Ellipse) SetBorderWidth(bw float64) { _geeb._gfbg = bw }

// SetBorderColor sets border color of the rectangle.
func (_accf *Rectangle) SetBorderColor(col Color) { _accf._bedf = col }

// Paragraph represents text drawn with a specified font and can wrap across lines and pages.
// By default, it occupies the available width in the drawing context.
type Paragraph struct {
	_egea         string
	_gbca         *_fgd.PdfFont
	_acdge        float64
	_fbcg         float64
	_bbfd         Color
	_cebg         TextAlignment
	_ecae         bool
	_gcbf         float64
	_dacad        int
	_dcefd        bool
	_cadb         float64
	_eebg         Margins
	_dbgab        Positioning
	_gddg         float64
	_ggfb         float64
	_gccbb, _eacb float64
	_ebdc         []string
	_dege         *int64
	_eadae        string
}

// CreateFrontPage sets a function to generate a front Page.
func (_fbd *Creator) CreateFrontPage(genFrontPageFunc func(_caa FrontpageFunctionArgs)) {
	_fbd._cee = genFrontPageFunc
}

// SetFillColor sets the fill color.
func (_abgf *CurvePolygon) SetFillColor(color Color) {
	_abgf._eeea = color
	_abgf._dgc.FillColor = _cfcee(color)
}

// Level returns the indentation level of the TOC line.
func (_addd *TOCLine) Level() uint                  { return _addd._gfec }
func _bdab(_gbgd Color, _fcaag float64) *ColorPoint { return &ColorPoint{_afddd: _gbgd, _ffag: _fcaag} }
func (_fdfb *templateProcessor) nodeLogError(_aebbf *templateNode, _ebfc string, _cadd ...interface{}) {
	_fec.Log.Error(_fdfb.getNodeErrorLocation(_aebbf, _ebfc, _cadd...))
}

// Subtotal returns the invoice subtotal description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_cbfd *Invoice) Subtotal() (*InvoiceCell, *InvoiceCell) { return _cbfd._beff[0], _cbfd._beff[1] }
func (_bdce *templateProcessor) parseRectangle(_ebgcac *templateNode) (interface{}, error) {
	_dgcfa := _bdce.creator.NewRectangle(0, 0, 0, 0)
	for _, _ceeb := range _ebgcac._adfdg.Attr {
		_gaaga := _ceeb.Value
		switch _bcbb := _ceeb.Name.Local; _bcbb {
		case "\u0078":
			_dgcfa._defb = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0079":
			_dgcfa._cdgf = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0077\u0069\u0064t\u0068":
			_dgcfa.SetWidth(_bdce.parseFloatAttr(_bcbb, _gaaga))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_dgcfa.SetHeight(_bdce.parseFloatAttr(_bcbb, _gaaga))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_dgcfa.SetFillColor(_bdce.parseColorAttr(_bcbb, _gaaga))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_dgcfa.SetFillOpacity(_bdce.parseFloatAttr(_bcbb, _gaaga))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_dgcfa.SetBorderColor(_bdce.parseColorAttr(_bcbb, _gaaga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_dgcfa.SetBorderOpacity(_bdce.parseFloatAttr(_bcbb, _gaaga))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_dgcfa.SetBorderWidth(_bdce.parseFloatAttr(_bcbb, _gaaga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_fgfa, _dgbbe, _bafc, _affee := _bdce.parseBorderRadiusAttr(_bcbb, _gaaga)
			_dgcfa.SetBorderRadius(_fgfa, _dgbbe, _affee, _bafc)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_dgcfa._ggeb = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_dgcfa._cbbc = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_dgcfa._ggfa = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_dgcfa._gcefb = _bdce.parseFloatAttr(_bcbb, _gaaga)
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_dgcfa.SetPositioning(_bdce.parsePositioningAttr(_bcbb, _gaaga))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_dgcfa.SetFitMode(_bdce.parseFitModeAttr(_bcbb, _gaaga))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_adcbf := _bdce.parseMarginAttr(_bcbb, _gaaga)
			_dgcfa.SetMargins(_adcbf.Left, _adcbf.Right, _adcbf.Top, _adcbf.Bottom)
		default:
			_bdce.nodeLogDebug(_ebgcac, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020re\u0063\u0074\u0061\u006e\u0067\u006ce\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _bcbb)
		}
	}
	return _dgcfa, nil
}

type componentRenderer interface{ Draw(_ececb Drawable) error }

// SetCompactMode sets the compact mode flag for this table.
//
// By enabling compact mode, table cell that contains Paragraph/StyleParagraph
// would not add extra height when calculating it's height.
//
// The default value is false.
func (_fgefg *Table) SetCompactMode(enable bool) { _fgefg._eadda = enable }

// SetColumns overwrites any columns in the line items table. This should be
// called before AddLine.
func (_eeab *Invoice) SetColumns(cols []*InvoiceCell) { _eeab._gffd = cols }

// SetFillOpacity sets the fill opacity.
func (_bcff *CurvePolygon) SetFillOpacity(opacity float64) { _bcff._fafc = opacity }
func (_caea *templateProcessor) parseInt64Attr(_aagff, _fgcdc string) int64 {
	_fec.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0069\u006e\u0074\u00364 a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _aagff, _fgcdc)
	_fdbcb, _ := _fg.ParseInt(_fgcdc, 10, 64)
	return _fdbcb
}
func (_dedcc *Table) clone() *Table {
	_ddcfda := *_dedcc
	_ddcfda._gecdg = make([]float64, len(_dedcc._gecdg))
	copy(_ddcfda._gecdg, _dedcc._gecdg)
	_ddcfda._cddfg = make([]float64, len(_dedcc._cddfg))
	copy(_ddcfda._cddfg, _dedcc._cddfg)
	_ddcfda._efbe = make([]*TableCell, 0, len(_dedcc._efbe))
	for _, _ggeg := range _dedcc._efbe {
		_fcfec := *_ggeg
		_fcfec._abdcb = &_ddcfda
		_ddcfda._efbe = append(_ddcfda._efbe, &_fcfec)
	}
	return &_ddcfda
}
func _cfbf(_geaga TextStyle) *List {
	return &List{_edaa: TextChunk{Text: "\u2022\u0020", Style: _geaga}, _aedbb: 0, _ebfa: true, _bdgca: PositionRelative, _fedb: _geaga}
}
func (_ec *Block) addContents(_ae *_dg.ContentStreamOperations) {
	_ec._cb.WrapIfNeeded()
	_ae.WrapIfNeeded()
	*_ec._cb = append(*_ec._cb, *_ae...)
}

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_baag *LinearShading) AddPatternResource(block *Block) (_efea _bc.PdfObjectName, _acbg error) {
	_bbaeag := 1
	_deaad := _bc.PdfObjectName("\u0050" + _fg.Itoa(_bbaeag))
	for block._dgd.HasPatternByName(_deaad) {
		_bbaeag++
		_deaad = _bc.PdfObjectName("\u0050" + _fg.Itoa(_bbaeag))
	}
	if _cege := block._dgd.SetPatternByName(_deaad, _baag.ToPdfShadingPattern().ToPdfObject()); _cege != nil {
		return "", _cege
	}
	return _deaad, nil
}

// SetNotes sets the notes section of the invoice.
func (_gcea *Invoice) SetNotes(title, content string) { _gcea._bbgb = [2]string{title, content} }

// Lines returns all the lines the table of contents has.
func (_afabb *TOC) Lines() []*TOCLine { return _afabb._ccfec }
func (_cdaca *TOCLine) prepareParagraph(_gaafa *StyledParagraph, _gfeaf DrawContext) {
	_ddcgb := _cdaca.Title.Text
	if _cdaca.Number.Text != "" {
		_ddcgb = "\u0020" + _ddcgb
	}
	_ddcgb += "\u0020"
	_cbgc := _cdaca.Page.Text
	if _cbgc != "" {
		_cbgc = "\u0020" + _cbgc
	}
	_gaafa._ecec = []*TextChunk{{Text: _cdaca.Number.Text, Style: _cdaca.Number.Style, _abefd: _cdaca.getLineLink()}, {Text: _ddcgb, Style: _cdaca.Title.Style, _abefd: _cdaca.getLineLink()}, {Text: _cbgc, Style: _cdaca.Page.Style, _abefd: _cdaca.getLineLink()}}
	_gaafa.wrapText()
	_cdbgf := len(_gaafa._bdfce)
	if _cdbgf == 0 {
		return
	}
	_dageb := _gfeaf.Width*1000 - _gaafa.getTextLineWidth(_gaafa._bdfce[_cdbgf-1])
	_aaddf := _gaafa.getTextLineWidth([]*TextChunk{&_cdaca.Separator})
	_febb := int(_dageb / _aaddf)
	_feaa := _eg.Repeat(_cdaca.Separator.Text, _febb)
	_ecgc := _cdaca.Separator.Style
	_fffgf := _gaafa.Insert(2, _feaa)
	_fffgf.Style = _ecgc
	_fffgf._abefd = _cdaca.getLineLink()
	_dageb = _dageb - float64(_febb)*_aaddf
	if _dageb > 500 {
		_dabbf, _gcaf := _ecgc.Font.GetRuneMetrics(' ')
		if _gcaf && _dageb > _dabbf.Wx {
			_efga := int(_dageb / _dabbf.Wx)
			if _efga > 0 {
				_ffgfb := _ecgc
				_ffgfb.FontSize = 1
				_fffgf = _gaafa.Insert(2, _eg.Repeat("\u0020", _efga))
				_fffgf.Style = _ffgfb
				_fffgf._abefd = _cdaca.getLineLink()
			}
		}
	}
}
func (_ffabc *templateProcessor) nodeError(_cccc *templateNode, _gddgg string, _eeeaf ...interface{}) error {
	return _f.Errorf("\u0025\u0073", _ffabc.getNodeErrorLocation(_cccc, _gddgg, _eeeaf...))
}

// SetPos sets the position of the graphic svg to the specified coordinates.
// This method sets the graphic svg to use absolute positioning.
func (_daeba *GraphicSVG) SetPos(x, y float64) {
	_daeba._ddcg = PositionAbsolute
	_daeba._gddbc = x
	_daeba._bgfb = y
}
func (_gedca *templateProcessor) parseListMarker(_bccda *templateNode) (interface{}, error) {
	if _bccda._cefd == nil {
		_gedca.nodeLogError(_bccda, "\u004c\u0069\u0073\u0074\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0063a\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _gdbeeg
	}
	var _cbddf *TextChunk
	switch _gefaa := _bccda._cefd._bedcd.(type) {
	case *List:
		_cbddf = &_gefaa._edaa
	case *listItem:
		_cbddf = &_gefaa._edee
	default:
		_gedca.nodeLogError(_bccda, "\u0025\u0076 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0066\u006f\u0072\u0020\u006c\u0069\u0073\u0074\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e", _gefaa)
		return nil, _gdbeeg
	}
	if _, _ggae := _gedca.parseTextChunk(_bccda, _cbddf); _ggae != nil {
		_gedca.nodeLogError(_bccda, "\u0043\u006f\u0075ld\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072\u0073e\u0020l\u0069s\u0074 \u006d\u0061\u0072\u006b\u0065\u0072\u003a\u0020\u0060\u0025\u0076\u0060\u002e", _ggae)
		return nil, nil
	}
	return _cbddf, nil
}
func _bfdcf(_bdeee *templateProcessor, _cdcb *templateNode) (interface{}, error) {
	return _bdeee.parseTableCell(_cdcb)
}

// SetColorBottom sets border color for bottom.
func (_befe *border) SetColorBottom(col Color) { _befe._dcf = col }
func (_adadc *templateProcessor) parseLinkAttr(_cbagf, _abab string) *_fgd.PdfAnnotation {
	_abab = _eg.TrimSpace(_abab)
	if _eg.HasPrefix(_abab, "\u0075\u0072\u006c(\u0027") && _eg.HasSuffix(_abab, "\u0027\u0029") && len(_abab) > 7 {
		return _fecd(_abab[5 : len(_abab)-2])
	}
	if _eg.HasPrefix(_abab, "\u0070\u0061\u0067e\u0028") && _eg.HasSuffix(_abab, "\u0029") && len(_abab) > 6 {
		var (
			_bggbg error
			_cbage int64
			_daagd float64
			_gece  float64
			_fdbfb = 1.0
			_cabgb = _eg.Split(_abab[5:len(_abab)-1], "\u002c")
		)
		_cbage, _bggbg = _fg.ParseInt(_eg.TrimSpace(_cabgb[0]), 10, 64)
		if _bggbg != nil {
			_fec.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _bggbg)
			return nil
		}
		if len(_cabgb) >= 2 {
			_daagd, _bggbg = _fg.ParseFloat(_eg.TrimSpace(_cabgb[1]), 64)
			if _bggbg != nil {
				_fec.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0058\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _bggbg)
				return nil
			}
		}
		if len(_cabgb) >= 3 {
			_gece, _bggbg = _fg.ParseFloat(_eg.TrimSpace(_cabgb[2]), 64)
			if _bggbg != nil {
				_fec.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0059\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _bggbg)
				return nil
			}
		}
		if len(_cabgb) >= 4 {
			_fdbfb, _bggbg = _fg.ParseFloat(_eg.TrimSpace(_cabgb[3]), 64)
			if _bggbg != nil {
				_fec.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u007a\u006f\u006f\u006d\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _bggbg)
				return nil
			}
		}
		return _cbdce(_cbage-1, _daagd, _gece, _fdbfb)
	}
	return nil
}

// SetBorderColor sets the border color of the ellipse.
func (_beef *Ellipse) SetBorderColor(col Color) { _beef._add = col }
func _dbca() *listItem                          { return &listItem{} }

// AddColorStop add color stop info for rendering gradient color.
func (_feda *RadialShading) AddColorStop(color Color, point float64) {
	_feda._dafd.AddColorStop(color, point)
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_dec *Block) DrawTemplate(c *Creator, r _gg.Reader, data interface{}, options *TemplateOptions) error {
	return _bdgfg(c, r, data, options, _dec)
}

// SetBorderOpacity sets the border opacity of the rectangle.
func (_gdcgb *Rectangle) SetBorderOpacity(opacity float64) { _gdcgb._dbff = opacity }

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_cgffg *RadialShading) SetBackgroundColor(backgroundColor Color) {
	_cgffg._dafd.SetBackgroundColor(backgroundColor)
}
func _bfaa(_bcef, _fdfd, _fagf, _facd float64) *border {
	_ebg := &border{}
	_ebg._bad = _bcef
	_ebg._abda = _fdfd
	_ebg._cbec = _fagf
	_ebg._aeg = _facd
	_ebg._dfd = ColorBlack
	_ebg._dcf = ColorBlack
	_ebg._bba = ColorBlack
	_ebg._gdbb = ColorBlack
	_ebg._dedg = 0
	_ebg._fcf = 0
	_ebg._aca = 0
	_ebg._gebg = 0
	_ebg.LineStyle = _gga.LineStyleSolid
	return _ebg
}
func (_ebeef *templateProcessor) parseRadialGradientAttr(creator *Creator, _gfbgb string) Color {
	_ggbg := ColorBlack
	if _gfbgb == "" {
		return _ggbg
	}
	var (
		_eecg   error
		_fcfecb = 0.0
		_cfecc  = 0.0
		_gcbcf  = -1.0
		_febcc  = _eg.Split(_gfbgb[16:len(_gfbgb)-1], "\u002c")
	)
	_dfaff := _eg.Fields(_febcc[0])
	if len(_dfaff) == 2 && _eg.TrimSpace(_dfaff[0])[0] != '#' {
		_fcfecb, _eecg = _fg.ParseFloat(_dfaff[0], 64)
		if _eecg != nil {
			_fec.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0058\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _eecg)
		}
		_cfecc, _eecg = _fg.ParseFloat(_dfaff[1], 64)
		if _eecg != nil {
			_fec.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0059\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _eecg)
		}
		_febcc = _febcc[1:]
	}
	_efdcf := _eg.TrimSpace(_febcc[0])
	if _efdcf[0] != '#' {
		_gcbcf, _eecg = _fg.ParseFloat(_efdcf, 64)
		if _eecg != nil {
			_fec.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0072\u0061\u0064\u0069\u0061l\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0073\u0069\u007ae\u003a\u0020\u0025\u0076", _eecg)
		}
		_febcc = _febcc[1:]
	}
	_gbede, _eccgce := _ebeef.processGradientColorPair(_febcc)
	if _gbede == nil || _eccgce == nil {
		return _ggbg
	}
	_fdgbb := creator.NewRadialGradientColor(_fcfecb, _cfecc, 0, _gcbcf, []*ColorPoint{})
	for _bage := 0; _bage < len(_gbede); _bage++ {
		_fdgbb.AddColorStop(_gbede[_bage], _eccgce[_bage])
	}
	return _fdgbb
}

// Width is not used. Not used as a Table element is designed to fill into
// available width depending on the context. Returns 0.
func (_fgfdc *Table) Width() float64 { return 0 }
func (_cbeeb *templateProcessor) parsePositioningAttr(_decf, _dgdee string) Positioning {
	_fec.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0069\u006e\u0067\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _decf, _dgdee)
	_aadd := map[string]Positioning{"\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065": PositionRelative, "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065": PositionAbsolute}[_dgdee]
	return _aadd
}

// SetAnnotation sets a annotation on a TextChunk.
func (_ccedg *TextChunk) SetAnnotation(annotation *_fgd.PdfAnnotation) { _ccedg._abefd = annotation }

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_aceaag *LinearShading) SetExtends(start bool, end bool) { _aceaag._dfae.SetExtends(start, end) }

// NewCell returns a new invoice table cell.
func (_bdff *Invoice) NewCell(value string) *InvoiceCell {
	return _bdff.newCell(value, _bdff.NewCellProps())
}

// Height returns the height of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_fbeg *Rectangle) Height() float64 { return _fbeg._cggg }

// RotatedSize returns the width and height of the rotated block.
func (_bcf *Block) RotatedSize() (float64, float64) {
	_, _, _cec, _cf := _ffcgd(_bcf._da, _bcf._ace, _bcf._df)
	return _cec, _cf
}

// RotateDeg rotates the current active page by angle degrees.  An error is returned on failure,
// which can be if there is no currently active page, or the angleDeg is not a multiple of 90 degrees.
func (_fgdd *Creator) RotateDeg(angleDeg int64) error {
	_ggc := _fgdd.getActivePage()
	if _ggc == nil {
		_fec.Log.Debug("F\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0065\u003a\u0020\u006e\u006f\u0020p\u0061\u0067\u0065\u0020\u0063\u0075\u0072\u0072\u0065\u006etl\u0079\u0020\u0061c\u0074i\u0076\u0065")
		return _bd.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	if angleDeg%90 != 0 {
		_fec.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067e\u0020\u0072\u006f\u0074\u0061\u0074\u0069on\u0020\u0061\u006e\u0067l\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006dul\u0074\u0069p\u006c\u0065\u0020\u006f\u0066\u0020\u0039\u0030")
		return _bd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	var _dbfa int64
	if _ggc.Rotate != nil {
		_dbfa = *(_ggc.Rotate)
	}
	_dbfa += angleDeg
	_ggc.Rotate = &_dbfa
	return nil
}

// SetLineColor sets the line color.
func (_efgcc *Polyline) SetLineColor(color Color) { _efgcc._cdggc.LineColor = _cfcee(color) }

// Scale scales Image by a constant factor, both width and height.
func (_gdbd *Image) Scale(xFactor, yFactor float64) {
	_gdbd._ecfc = xFactor * _gdbd._ecfc
	_gdbd._fegb = yFactor * _gdbd._fegb
}

type border struct {
	_bad      float64
	_abda     float64
	_cbec     float64
	_aeg      float64
	_aff      Color
	_bba      Color
	_aca      float64
	_dcf      Color
	_fcf      float64
	_gdbb     Color
	_gebg     float64
	_dfd      Color
	_dedg     float64
	LineStyle _gga.LineStyle
	_ffbb     CellBorderStyle
	_cfc      CellBorderStyle
	_cggf     CellBorderStyle
	_fagg     CellBorderStyle
}

// GeneratePageBlocks draws the curve onto page blocks.
func (_cbbd *Curve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_adebb := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_bbaea := _dg.NewContentCreator()
	if _cbbd._afbbc != nil {
		_bbaea.Add_BDC(*_bc.MakeName(_fgd.StructureTypeFigure), map[string]_bc.PdfObject{"\u004d\u0043\u0049\u0044": _bc.MakeInteger(*_cbbd._afbbc)})
	}
	_bbaea.Add_q().Add_w(_cbbd._befg).SetStrokingColor(_cfcee(_cbbd._afgg)).Add_m(_cbbd._aecc, ctx.PageHeight-_cbbd._egda).Add_v(_cbbd._bgaca, ctx.PageHeight-_cbbd._afbc, _cbbd._cdf, ctx.PageHeight-_cbbd._adca).Add_S().Add_Q()
	if _cbbd._afbbc != nil {
		_bbaea.Add_EMC()
	}
	_eegd := _adebb.addContentsByString(_bbaea.String())
	if _eegd != nil {
		return nil, ctx, _eegd
	}
	return []*Block{_adebb}, ctx, nil
}

// PageFinalize sets a function to be called for each page before finalization
// (i.e. the last stage of page processing before they get written out).
// The callback function allows final touch-ups for each page, and it
// provides information that might not be known at other stages of designing
// the document (e.g. the total number of pages). Unlike the header/footer
// functions, which are limited to the top/bottom margins of the page, the
// finalize function can be used draw components anywhere on the current page.
func (_fcab *Creator) PageFinalize(pageFinalizeFunc func(_dceb PageFinalizeFunctionArgs) error) {
	_fcab._bbad = pageFinalizeFunc
}

type shading struct {
	_feac  Color
	_ccea  bool
	_fece  []bool
	_bcaca []*ColorPoint
}

func _accfg(_fbffbd *templateProcessor, _bceg *templateNode) (interface{}, error) {
	return _fbffbd.parseListMarker(_bceg)
}

// NewRadialGradientColor creates a radial gradient color that could act as a color in other componenents.
// Note: The innerRadius must be smaller than outerRadius for the circle to render properly.
func (_badf *Creator) NewRadialGradientColor(x float64, y float64, innerRadius float64, outerRadius float64, colorPoints []*ColorPoint) *RadialShading {
	return _cgff(x, y, innerRadius, outerRadius, colorPoints)
}

// SetIndent sets the cell's left indent.
func (_fgddb *TableCell) SetIndent(indent float64) { _fgddb._fceba = indent }

// SetLineSeparator sets the separator for all new lines of the table of contents.
func (_aaeab *TOC) SetLineSeparator(separator string) { _aaeab._begag = separator }

// Margins returns the margins of the component.
func (_dgcc *Division) Margins() (_acga, _baacg, _ege, _cece float64) {
	return _dgcc._ddfg.Left, _dgcc._ddfg.Right, _dgcc._ddfg.Top, _dgcc._ddfg.Bottom
}
func (_bgee *TableCell) height(_daaa float64) float64 {
	var _efead float64
	switch _dfcbf := _bgee._aafg.(type) {
	case *Paragraph:
		if _dfcbf._ecae {
			_dfcbf.SetWidth(_daaa - _bgee._fceba - _dfcbf._eebg.Left - _dfcbf._eebg.Right)
		}
		_efead = _dfcbf.Height() + _dfcbf._eebg.Top + _dfcbf._eebg.Bottom
		if !_bgee._abdcb._eadda {
			_efead += (0.5 * _dfcbf._acdge * _dfcbf._fbcg)
		}
	case *StyledParagraph:
		if _dfcbf._fegca {
			_dfcbf.SetWidth(_daaa - _bgee._fceba - _dfcbf._ccddg.Left - _dfcbf._ccddg.Right)
		}
		_efead = _dfcbf.Height() + _dfcbf._ccddg.Top + _dfcbf._ccddg.Bottom
		if !_bgee._abdcb._eadda {
			_efead += (0.5 * _dfcbf.getTextHeight())
		}
	case *Image:
		_dfcbf.applyFitMode(_daaa - _bgee._fceba)
		_efead = _dfcbf.Height() + _dfcbf._cdbc.Top + _dfcbf._cdbc.Bottom
	case *Table:
		_dfcbf.updateRowHeights(_daaa - _bgee._fceba - _dfcbf._gbcea.Left - _dfcbf._gbcea.Right)
		_efead = _dfcbf.Height() + _dfcbf._gbcea.Top + _dfcbf._gbcea.Bottom
	case *List:
		_efead = _dfcbf.ctxHeight(_daaa-_bgee._fceba) + _dfcbf._gacb.Top + _dfcbf._gacb.Bottom
	case *Division:
		_efead = _dfcbf.ctxHeight(_daaa-_bgee._fceba) + _dfcbf._ddfg.Top + _dfcbf._ddfg.Bottom + _dfcbf._gbga.Top + _dfcbf._gbga.Bottom
	case *Chart:
		_efead = _dfcbf.Height() + _dfcbf._bgfg.Top + _dfcbf._bgfg.Bottom
	case *Rectangle:
		_dfcbf.applyFitMode(_daaa - _bgee._fceba)
		_efead = _dfcbf.Height() + _dfcbf._gagb.Top + _dfcbf._gagb.Bottom + _dfcbf._caeg
	case *Ellipse:
		_dfcbf.applyFitMode(_daaa - _bgee._fceba)
		_efead = _dfcbf.Height() + _dfcbf._ddbd.Top + _dfcbf._ddbd.Bottom
	case *Line:
		_efead = _dfcbf.Height() + _dfcbf._fggd.Top + _dfcbf._fggd.Bottom
	}
	return _efead
}
func _fbdb(_cfac []byte) (*Image, error) {
	_dfdb := _c.NewReader(_cfac)
	_gccg, _gbbe := _fgd.ImageHandling.Read(_dfdb)
	if _gbbe != nil {
		_fec.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _gbbe)
		return nil, _gbbe
	}
	return _eddeb(_gccg)
}

// SetAngle sets the rotation angle in degrees.
func (_af *Block) SetAngle(angleDeg float64) { _af._df = angleDeg }

// GetCoords returns the center coordinates of ellipse (`xc`, `yc`).
func (_dbcc *Ellipse) GetCoords() (float64, float64) { return _dbcc._adbc, _dbcc._ddaee }

// SetFillColor sets the fill color of the rectangle.
func (_afaf *Rectangle) SetFillColor(col Color) { _afaf._dgab = col }
func _fgba(_efda *templateProcessor, _egac *templateNode) (interface{}, error) {
	return _efda.parseChapterHeading(_egac)
}

// AddSubtable copies the cells of the subtable in the table, starting with the
// specified position. The table row and column indices are 1-based, which
// makes the position of the first cell of the first row of the table 1,1.
// The table is automatically extended if the subtable exceeds its columns.
// This can happen when the subtable has more columns than the table or when
// one or more columns of the subtable starting from the specified position
// exceed the last column of the table.
func (_dgbb *Table) AddSubtable(row, col int, subtable *Table) {
	for _, _egaee := range subtable._efbe {
		_fedd := &TableCell{}
		*_fedd = *_egaee
		_fedd._abdcb = _dgbb
		_fedd._bacg += col - 1
		if _aeegc := _dgbb._afacb - (_fedd._bacg - 1); _aeegc < _fedd._ecddcb {
			_dgbb._afacb += _fedd._ecddcb - _aeegc
			_dgbb.resetColumnWidths()
			_fec.Log.Debug("\u0054a\u0062l\u0065\u003a\u0020\u0073\u0075\u0062\u0074\u0061\u0062\u006c\u0065 \u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0061\u0062\u006c\u0065\u002e\u0020\u0045\u0078\u0070\u0061\u006e\u0064\u0069\u006e\u0067\u0020\u0074\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0025\u0064\u0020\u0063\u006fl\u0075\u006d\u006e\u0073\u002e", _dgbb._afacb)
		}
		_fedd._deef += row - 1
		_dgfaf := subtable._gecdg[_egaee._deef-1]
		if _fedd._deef > _dgbb._begg {
			for _fedd._deef > _dgbb._begg {
				_dgbb._begg++
				_dgbb._gecdg = append(_dgbb._gecdg, _dgbb._edefd)
			}
			_dgbb._gecdg[_fedd._deef-1] = _dgfaf
		} else {
			_dgbb._gecdg[_fedd._deef-1] = _fc.Max(_dgbb._gecdg[_fedd._deef-1], _dgfaf)
		}
		_dgbb._efbe = append(_dgbb._efbe, _fedd)
	}
	_dgbb.sortCells()
}

// InfoLines returns all the rows in the invoice information table as
// description-value cell pairs.
func (_gaad *Invoice) InfoLines() [][2]*InvoiceCell {
	_agad := [][2]*InvoiceCell{_gaad._bcab, _gaad._cbfg, _gaad._dfge}
	return append(_agad, _gaad._afff...)
}

// SetPos sets absolute positioning with specified coordinates.
func (_ecga *StyledParagraph) SetPos(x, y float64) {
	_ecga._fceb = PositionAbsolute
	_ecga._abeag = x
	_ecga._agdg = y
}

// Positioning returns the type of positioning the line is set to use.
func (_bddg *Line) Positioning() Positioning { return _bddg._babe }
func _afafb(_dadc, _gggef, _eace string, _beea uint, _cfecca TextStyle) *TOCLine {
	return _abgd(TextChunk{Text: _dadc, Style: _cfecca}, TextChunk{Text: _gggef, Style: _cfecca}, TextChunk{Text: _eace, Style: _cfecca}, _beea, _cfecca)
}

// GetCoords returns the upper left corner coordinates of the rectangle (`x`, `y`).
func (_bged *Rectangle) GetCoords() (float64, float64) { return _bged._defb, _bged._cdgf }
func _ddaea(_aacaf *_fgd.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 0), Font: _aacaf, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

// ScaleToHeight scales the rectangle to the specified height. The width of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_adff *Rectangle) ScaleToHeight(h float64) {
	_fafaa := _adff._bcede / _adff._cggg
	_adff._cggg = h
	_adff._bcede = h * _fafaa
}

// SetDate sets the date of the invoice.
func (_dada *Invoice) SetDate(date string) (*InvoiceCell, *InvoiceCell) {
	_dada._cbfg[1].Value = date
	return _dada._cbfg[0], _dada._cbfg[1]
}

// AddressStyle returns the style properties used to render the content of
// the invoice address sections.
func (_gecb *Invoice) AddressStyle() TextStyle { return _gecb._efdd }

// AddLine adds a new line with the provided style to the table of contents.
func (_gccae *TOC) AddLine(line *TOCLine) *TOCLine {
	if line == nil {
		return nil
	}
	_gccae._ccfec = append(_gccae._ccfec, line)
	return line
}

// GeneratePageBlocks generates the page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages. Implements the Drawable interface.
func (_cgef *Paragraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_feeb := ctx
	var _gdga []*Block
	_dgeb := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _cgef._dbgab.IsRelative() {
		ctx.X += _cgef._eebg.Left
		ctx.Y += _cgef._eebg.Top
		ctx.Width -= _cgef._eebg.Left + _cgef._eebg.Right
		ctx.Height -= _cgef._eebg.Top
		_cgef.SetWidth(ctx.Width)
		if _cgef.Height() > ctx.Height {
			_gdga = append(_gdga, _dgeb)
			_dgeb = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_afde := ctx
			_afde.Y = ctx.Margins.Top
			_afde.X = ctx.Margins.Left + _cgef._eebg.Left
			_afde.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
			_afde.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _cgef._eebg.Left - _cgef._eebg.Right
			ctx = _afde
		}
	} else {
		if int(_cgef._gcbf) <= 0 {
			_cgef.SetWidth(_cgef.getTextWidth())
		}
		ctx.X = _cgef._gddg
		ctx.Y = _cgef._ggfb
	}
	ctx, _fbab := _afae(_dgeb, _cgef, ctx)
	if _fbab != nil {
		_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fbab)
		return nil, ctx, _fbab
	}
	_gdga = append(_gdga, _dgeb)
	if _cgef._dbgab.IsRelative() {
		ctx.Y += _cgef._eebg.Bottom
		ctx.Height -= _cgef._eebg.Bottom
		if !ctx.Inline {
			ctx.X = _feeb.X
			ctx.Width = _feeb.Width
		}
		return _gdga, ctx, nil
	}
	return _gdga, _feeb, nil
}

// SetBorderLineStyle sets border style (currently dashed or plain).
func (_cgcf *TableCell) SetBorderLineStyle(style _gga.LineStyle) { _cgcf._gaddf = style }

// Add adds a new line with the default style to the table of contents.
func (_ggcaa *TOC) Add(number, title, page string, level uint) *TOCLine {
	_aeec := _ggcaa.AddLine(_abgd(TextChunk{Text: number, Style: _ggcaa._cggbgd}, TextChunk{Text: title, Style: _ggcaa._fefdb}, TextChunk{Text: page, Style: _ggcaa._caeab}, level, _ggcaa._eddd))
	if _aeec == nil {
		return nil
	}
	_dbgga := &_ggcaa._becec
	_aeec.SetMargins(_dbgga.Left, _dbgga.Right, _dbgga.Top, _dbgga.Bottom)
	_aeec.SetLevelOffset(_ggcaa._eaadf)
	_aeec.Separator.Text = _ggcaa._begag
	_aeec.Separator.Style = _ggcaa._fadge
	return _aeec
}

// SetRowHeight sets the height for a specified row.
func (_abedc *Table) SetRowHeight(row int, h float64) error {
	if row < 1 || row > len(_abedc._gecdg) {
		return _bd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_abedc._gecdg[row-1] = h
	return nil
}
func (_beagb *StyledParagraph) getTextHeight() float64 {
	var _adac float64
	for _, _aaae := range _beagb._ecec {
		_gcgecb := _aaae.Style.FontSize * _beagb._fgef
		if _gcgecb > _adac {
			_adac = _gcgecb
		}
	}
	return _adac
}

// TemplateOptions contains options and resources to use when rendering
// a template with a Creator instance.
// All the resources in the map fields can be referenced by their
// name/key in the template which is rendered using the options instance.
type TemplateOptions struct {

	// HelperFuncMap is used to define functions which can be accessed
	// inside the rendered templates by their assigned names.
	HelperFuncMap _g.FuncMap

	// SubtemplateMap contains templates which can be rendered alongside
	// the main template. They can be accessed using their assigned names
	// in the main template or in the other subtemplates.
	// Subtemplates defined inside the subtemplates specified in the map
	// can be accessed directly.
	// All resources available to the main template are also available
	// to the subtemplates.
	SubtemplateMap map[string]_gg.Reader

	// FontMap contains pre-loaded fonts which can be accessed
	// inside the rendered templates by their assigned names.
	FontMap map[string]*_fgd.PdfFont

	// ImageMap contains pre-loaded images which can be accessed
	// inside the rendered templates by their assigned names.
	ImageMap map[string]*_fgd.Image

	// ColorMap contains colors which can be accessed
	// inside the rendered templates by their assigned names.
	ColorMap map[string]Color

	// ChartMap contains charts which can be accessed
	// inside the rendered templates by their assigned names.
	ChartMap map[string]_cc.ChartRenderable
}

// GraphicSVG represents a drawable graphic SVG.
// It is used to render the graphic SVG components using a creator instance.
type GraphicSVG struct {
	_dgccb *_dd.GraphicSVG
	_ddcg  Positioning
	_gddbc float64
	_bgfb  float64
	_aagce Margins
	_bdae  *int64
}

func _ggadb(_accca *templateProcessor, _cccbc *templateNode) (interface{}, error) {
	return _accca.parseLine(_cccbc)
}
func (_fgcea *templateProcessor) parseLinearGradientAttr(creator *Creator, _fbba string) Color {
	_bbdef := ColorBlack
	if _fbba == "" {
		return _bbdef
	}
	_gagfe := creator.NewLinearGradientColor([]*ColorPoint{})
	_gagfe.SetExtends(true, true)
	var (
		_cfaae = _eg.Split(_fbba[16:len(_fbba)-1], "\u002c")
		_gbef  = _eg.TrimSpace(_cfaae[0])
	)
	if _eg.HasSuffix(_gbef, "\u0064\u0065\u0067") {
		_afgcc, _gecdb := _fg.ParseFloat(_gbef[:len(_gbef)-3], 64)
		if _gecdb != nil {
			_fec.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0061\u006e\u0067\u006ce\u003a\u0020\u0025\u0076", _gecdb)
		} else {
			_gagfe.SetAngle(_afgcc)
		}
		_cfaae = _cfaae[1:]
	}
	_ceggb, _fgcg := _fgcea.processGradientColorPair(_cfaae)
	if _ceggb == nil || _fgcg == nil {
		return _bbdef
	}
	for _bccff := 0; _bccff < len(_ceggb); _bccff++ {
		_gagfe.AddColorStop(_ceggb[_bccff], _fgcg[_bccff])
	}
	return _gagfe
}

// NewImageFromData creates an Image from image data.
func (_bbaa *Creator) NewImageFromData(data []byte) (*Image, error) { return _fbdb(data) }

// Height returns the height of the Paragraph. The height is calculated based on the input text and
// how it is wrapped within the container. Does not include Margins.
func (_dcfec *Paragraph) Height() float64 {
	_dcfec.wrapText()
	return float64(len(_dcfec._ebdc)) * _dcfec._fbcg * _dcfec._acdge
}

// SetColor sets the color of the Paragraph text.
//
// Example:
//
//  1. p := NewParagraph("Red paragraph")
//     // Set to red color with a hex code:
//     p.SetColor(creator.ColorRGBFromHex("#ff0000"))
//
//  2. Make Paragraph green with 8-bit rgb values (0-255 each component)
//     p.SetColor(creator.ColorRGBFrom8bit(0, 255, 0)
//
//  3. Make Paragraph blue with arithmetic (0-1) rgb components.
//     p.SetColor(creator.ColorRGBFromArithmetic(0, 0, 1.0)
func (_adag *Paragraph) SetColor(col Color) { _adag._bbfd = col }

// SetBorderWidth sets the border width.
func (_daebd *PolyBezierCurve) SetBorderWidth(borderWidth float64) {
	_daebd._ffcd.BorderWidth = borderWidth
}

// StyledParagraph represents text drawn with a specified font and can wrap across lines and pages.
// By default occupies the available width in the drawing context.
type StyledParagraph struct {
	_ecec  []*TextChunk
	_fgbg  TextStyle
	_fbegf TextStyle
	_abff  TextAlignment
	_ddec  TextVerticalAlignment
	_fgef  float64
	_fegca bool
	_ffcfd float64
	_eadc  bool
	_fegde bool
	_ecedg TextOverflow
	_ffab  float64
	_ccddg Margins
	_fceb  Positioning
	_abeag float64
	_agdg  float64
	_bagff float64
	_cebbd float64
	_bdfce [][]*TextChunk
	_beccd func(_agcc *StyledParagraph, _ddecb DrawContext)
	_eeba  *int64
	_cccab string
}

func (_bab *Chapter) headingNumber() string {
	var _febc string
	if _bab._aee {
		if _bab._gbf != 0 {
			_febc = _fg.Itoa(_bab._gbf) + "\u002e"
		}
		if _bab._cdee != nil {
			_ccc := _bab._cdee.headingNumber()
			if _ccc != "" {
				_febc = _ccc + _febc
			}
		}
	}
	return _febc
}

// DrawContext defines the drawing context. The DrawContext is continuously used and updated when
// drawing the page contents in relative mode.  Keeps track of current X, Y position, available
// height as well as other page parameters such as margins and dimensions.
type DrawContext struct {

	// Current page number.
	Page int

	// Current position.  In a relative positioning mode, a drawable will be placed at these coordinates.
	X, Y float64

	// Context dimensions.  Available width and height (on current page).
	Width, Height float64

	// Page Margins.
	Margins Margins

	// Absolute Page size, widths and height.
	PageWidth  float64
	PageHeight float64

	// Controls whether the components are stacked horizontally
	Inline bool
	_cffd  rune
	_agd   []error
}

// Width returns the width of the ellipse.
func (_ccddc *Ellipse) Width() float64 { return _ccddc._eefg }

// SetLineSeparatorStyle sets the style for the separator part of all new
// lines of the table of contents.
func (_aaegg *TOC) SetLineSeparatorStyle(style TextStyle) { _aaegg._fadge = style }

// Indent returns the left offset of the list when nested into another list.
func (_caae *List) Indent() float64 { return _caae._aedbb }
func (_bdbf *pageTransformations) applyFlip(_cddc *_fgd.PdfPage) error {
	_gfe, _dadb := _bdbf._bbed, _bdbf._fcge
	if !_gfe && !_dadb {
		return nil
	}
	if _cddc == nil {
		return _bd.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_bcfc, _dge := _cddc.GetMediaBox()
	if _dge != nil {
		return _dge
	}
	_fgcc, _dcga := _bcfc.Width(), _bcfc.Height()
	_ada, _dge := _cddc.GetRotate()
	if _dge != nil {
		_fec.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _dge.Error())
	}
	if _dcgc := _ada%360 != 0 && _ada%90 == 0; _dcgc {
		if _cedb := (360 + _ada%360) % 360; _cedb == 90 || _cedb == 270 {
			_gfe, _dadb = _dadb, _gfe
		}
	}
	_fegf, _gcba := 1.0, 0.0
	if _gfe {
		_fegf, _gcba = -1.0, -_fgcc
	}
	_eaf, _bbae := 1.0, 0.0
	if _dadb {
		_eaf, _bbae = -1.0, -_dcga
	}
	_bgge := _dg.NewContentCreator().Scale(_fegf, _eaf).Translate(_gcba, _bbae)
	_fgb, _dge := _bc.MakeStream(_bgge.Bytes(), _bc.NewFlateEncoder())
	if _dge != nil {
		return _dge
	}
	_eeac := _bc.MakeArray(_fgb)
	_eeac.Append(_cddc.GetContentStreamObjs()...)
	_cddc.Contents = _eeac
	return nil
}

// AddTextItem appends a new item with the specified text to the list.
// The method creates a styled paragraph with the specified text and returns
// it so that the item style can be customized.
// The method also returns the marker used for the newly added item.
// The marker object can be used to change the text and style of the marker
// for the current item.
func (_eedff *List) AddTextItem(text string) (*StyledParagraph, *TextChunk, error) {
	_acdg := _ddge(_eedff._fedb)
	_acdg.Append(text)
	_gecd, _gfca := _eedff.Add(_acdg)
	return _acdg, _gecd, _gfca
}
func (_cecea *templateProcessor) nodeLogDebug(_fccaa *templateNode, _acac string, _badcg ...interface{}) {
	_fec.Log.Debug(_cecea.getNodeErrorLocation(_fccaa, _acac, _badcg...))
}
func _fecd(_adae string) *_fgd.PdfAnnotation {
	_acdeb := _fgd.NewPdfAnnotationLink()
	_dgage := _fgd.NewBorderStyle()
	_dgage.SetBorderWidth(0)
	_acdeb.BS = _dgage.ToPdfObject()
	_bgccf := _fgd.NewPdfActionURI()
	_bgccf.URI = _bc.MakeString(_adae)
	_acdeb.SetAction(_bgccf.PdfAction)
	return _acdeb.PdfAnnotation
}

// GetHeading returns the chapter heading paragraph. Used to give access to address style: font, sizing etc.
func (_cdab *Chapter) GetHeading() *Paragraph { return _cdab._gfdf }
func _gfcbf(_cgfab *templateProcessor, _bbfa *templateNode) (interface{}, error) {
	return _cgfab.parsePageBreak(_bbfa)
}
func (_cbecc *templateProcessor) parseList(_cgde *templateNode) (interface{}, error) {
	_deccc := _cbecc.creator.NewList()
	for _, _gbgafd := range _cgde._adfdg.Attr {
		_gabec := _gbgafd.Value
		switch _aadbb := _gbgafd.Name.Local; _aadbb {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_deccc.SetIndent(_cbecc.parseFloatAttr(_aadbb, _gabec))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_ebeb := _cbecc.parseMarginAttr(_aadbb, _gabec)
			_deccc.SetMargins(_ebeb.Left, _ebeb.Right, _ebeb.Top, _ebeb.Bottom)
		default:
			_cbecc.nodeLogDebug(_cgde, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u0073\u0074 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aadbb)
		}
	}
	return _deccc, nil
}

// AddInternalLink adds a new internal link to the paragraph.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
// The zoom of the destination page is controlled with the zoom
// parameter. Pass in 0 to keep the current zoom value.
func (_dggcb *StyledParagraph) AddInternalLink(text string, page int64, x, y, zoom float64) *TextChunk {
	_eeaee := NewTextChunk(text, _dggcb._fbegf)
	_eeaee._abefd = _cbdce(page-1, x, y, zoom)
	return _dggcb.appendChunk(_eeaee)
}

// SetDueDate sets the due date of the invoice.
func (_bag *Invoice) SetDueDate(dueDate string) (*InvoiceCell, *InvoiceCell) {
	_bag._dfge[1].Value = dueDate
	return _bag._dfge[0], _bag._dfge[1]
}
func (_cacb *Image) rotatedSize() (float64, float64) {
	_baeb := _cacb._ecfc
	_gfgd := _cacb._fegb
	_adga := _cacb._fbbf
	if _adga == 0 {
		return _baeb, _gfgd
	}
	_fccd := _gga.Path{Points: []_gga.Point{_gga.NewPoint(0, 0).Rotate(_adga), _gga.NewPoint(_baeb, 0).Rotate(_adga), _gga.NewPoint(0, _gfgd).Rotate(_adga), _gga.NewPoint(_baeb, _gfgd).Rotate(_adga)}}.GetBoundingBox()
	return _fccd.Width, _fccd.Height
}

// Width returns the width of the Paragraph.
func (_ggdba *Paragraph) Width() float64 {
	if _ggdba._ecae && int(_ggdba._gcbf) > 0 {
		return _ggdba._gcbf
	}
	return _ggdba.getTextWidth() / 1000.0
}

// Date returns the invoice date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_efdc *Invoice) Date() (*InvoiceCell, *InvoiceCell) { return _efdc._cbfg[0], _efdc._cbfg[1] }
func (_bdc *Block) transform(_ffa _de.Matrix) {
	_cae := _dg.NewContentCreator().Add_cm(_ffa[0], _ffa[1], _ffa[3], _ffa[4], _ffa[6], _ffa[7]).Operations()
	*_bdc._cb = append(*_cae, *_bdc._cb...)
	_bdc._cb.WrapIfNeeded()
}
func _gfgfe(_edagf *templateProcessor, _daebe *templateNode) (interface{}, error) {
	return _edagf.parseChart(_daebe)
}

// SetSideBorderWidth sets the cell's side border width.
func (_bdagd *TableCell) SetSideBorderWidth(side CellBorderSide, width float64) {
	switch side {
	case CellBorderSideAll:
		_bdagd._cddebf = width
		_bdagd._eabab = width
		_bdagd._caef = width
		_bdagd._adec = width
	case CellBorderSideTop:
		_bdagd._cddebf = width
	case CellBorderSideBottom:
		_bdagd._eabab = width
	case CellBorderSideLeft:
		_bdagd._caef = width
	case CellBorderSideRight:
		_bdagd._adec = width
	}
}
func (_caad *templateProcessor) getNodeErrorLocation(_cgfad *templateNode, _daeff string, _beecc ...interface{}) string {
	_cggbg := _f.Sprintf(_daeff, _beecc...)
	_aaafc := _f.Sprintf("\u0025\u0064", _cgfad._agdga)
	if _cgfad._fcebb != 0 {
		_aaafc = _f.Sprintf("\u0025\u0064\u003a%\u0064", _cgfad._fcebb, _cgfad._gefcf)
	}
	if _caad._edcgd != "" {
		return _f.Sprintf("\u0025\u0073\u0020\u005b\u0025\u0073\u003a\u0025\u0073\u005d", _cggbg, _caad._edcgd, _aaafc)
	}
	return _f.Sprintf("\u0025s\u0020\u005b\u0025\u0073\u005d", _cggbg, _aaafc)
}

// NewTOC creates a new table of contents.
func (_eaeg *Creator) NewTOC(title string) *TOC {
	_eaba := _eaeg.NewTextStyle()
	_eaba.Font = _eaeg._ceb
	return _cffe(title, _eaeg.NewTextStyle(), _eaba)
}

// SetWidthTop sets border width for top.
func (_gfgf *border) SetWidthTop(bw float64) { _gfgf._dedg = bw }
func (_gdab *StyledParagraph) getLineMetrics(_dccg int) (_aagcd, _dadd, _ddda float64) {
	if _gdab._bdfce == nil || (_gdab._bdfce != nil && len(_gdab._bdfce) == 0) {
		_gdab.wrapText()
	}
	if _dccg < 0 || _dccg > len(_gdab._bdfce)-1 {
		_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020p\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u006c\u0069\u006e\u0065 \u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002e\u0020\u0052\u0065tu\u0072\u006e\u0069\u006e\u0067\u0020\u0030\u002c\u0020\u0030", _dccg)
		return 0, 0, 0
	}
	_caff := _gdab._bdfce[_dccg]
	for _, _cgag := range _caff {
		_dfbb := _aggdb(_cgag.Style.Font, _cgag.Style.FontSize)
		if _dfbb._geede > _aagcd {
			_aagcd = _dfbb._geede
		}
		if _dfbb._egeb < _ddda {
			_ddda = _dfbb._egeb
		}
		if _aegc := _cgag.Style.FontSize; _aegc > _dadd {
			_dadd = _aegc
		}
	}
	return _aagcd, _dadd, _ddda
}

const (
	CellBorderStyleNone CellBorderStyle = iota
	CellBorderStyleSingle
	CellBorderStyleDouble
)

// ColorCMYKFrom8bit creates a Color from c,m,y,k values (0-100).
// Example:
//
//	red := ColorCMYKFrom8Bit(0, 100, 100, 0)
func ColorCMYKFrom8bit(c, m, y, k byte) Color {
	return cmykColor{_ffea: _fc.Min(float64(c), 100) / 100.0, _cbg: _fc.Min(float64(m), 100) / 100.0, _egb: _fc.Min(float64(y), 100) / 100.0, _gbfd: _fc.Min(float64(k), 100) / 100.0}
}

// TextRenderingMode determines whether showing text shall cause glyph
// outlines to be stroked, filled, used as a clipping boundary, or some
// combination of the three.
// See section 9.3 "Text State Parameters and Operators" and
// Table 106 (pp. 254-255 PDF32000_2008).
type TextRenderingMode int

func _gaegc(_eacef interface{}) (interface{}, error) {
	switch _agced := _eacef.(type) {
	case uint8:
		return int64(_agced), nil
	case int8:
		return int64(_agced), nil
	case uint16:
		return int64(_agced), nil
	case int16:
		return int64(_agced), nil
	case uint32:
		return int64(_agced), nil
	case int32:
		return int64(_agced), nil
	case uint64:
		return int64(_agced), nil
	case int64:
		return _agced, nil
	case int:
		return int64(_agced), nil
	case float32:
		return float64(_agced), nil
	case float64:
		return _agced, nil
	}
	return nil, _f.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0076\u0061\u006c\u0075\u0065\u002c\u0020\u0025\u0076\u0020\u0069\u0073 \u006e\u006f\u0074\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072", _eacef)
}

// SetAngle sets Image rotation angle in degrees.
func (_adfec *Image) SetAngle(angle float64) { _adfec._fbbf = angle }

const (
	TextVerticalAlignmentBaseline TextVerticalAlignment = iota
	TextVerticalAlignmentCenter
	TextVerticalAlignmentBottom
	TextVerticalAlignmentTop
)

// DueDate returns the invoice due date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_dafca *Invoice) DueDate() (*InvoiceCell, *InvoiceCell) {
	return _dafca._dfge[0], _dafca._dfge[1]
}

// SetFillOpacity sets the fill opacity of the ellipse.
func (_bfbf *Ellipse) SetFillOpacity(opacity float64) { _bfbf._fcda = opacity }

// Cols returns the total number of columns the table has.
func (_fcdad *Table) Cols() int { return _fcdad._afacb }

// TextDecorationLineStyle represents the style of lines used to decorate
// a text chunk (e.g. underline).
type TextDecorationLineStyle struct {

	// Color represents the color of the line (default: the color of the text).
	Color Color

	// Offset represents the vertical offset of the line (default: 1).
	Offset float64

	// Thickness represents the thickness of the line (default: 1).
	Thickness float64
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_cbf *Creator) DrawTemplate(r _gg.Reader, data interface{}, options *TemplateOptions) error {
	return _bdgfg(_cbf, r, data, options, _cbf)
}

// AppendCurve appends a Bezier curve to the filled curve.
func (_aeda *FilledCurve) AppendCurve(curve _gga.CubicBezierCurve) *FilledCurve {
	_aeda._acce = append(_aeda._acce, curve)
	return _aeda
}

type cmykColor struct{ _ffea, _cbg, _egb, _gbfd float64 }

// SetTOC sets the table of content component of the creator.
// This method should be used when building a custom table of contents.
func (_agac *Creator) SetTOC(toc *TOC) {
	if toc == nil {
		return
	}
	_agac._aac = toc
}

// SetMargins sets the Paragraph's margins.
func (_gaec *Paragraph) SetMargins(left, right, top, bottom float64) {
	_gaec._eebg.Left = left
	_gaec._eebg.Right = right
	_gaec._eebg.Top = top
	_gaec._eebg.Bottom = bottom
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width,
// i.e. the width the text can extend to prior to wrapping over to next line.
func (_dgda *StyledParagraph) SetWidth(width float64) { _dgda._ffcfd = width; _dgda.wrapText() }
func (_cfbe *Line) computeCoords(_cdgbb DrawContext) (_gfee, _bcbd, _fdcb, _fbda float64) {
	_gfee = _cdgbb.X
	_fdcb = _gfee + _cfbe._cfge - _cfbe._gbcgde
	_acfc := _cfbe._aebf
	if _cfbe._gbcgde == _cfbe._cfge {
		_acfc /= 2
	}
	if _cfbe._gcfa < _cfbe._ceef {
		_bcbd = _cdgbb.PageHeight - _cdgbb.Y - _acfc
		_fbda = _bcbd - _cfbe._ceef + _cfbe._gcfa
	} else {
		_fbda = _cdgbb.PageHeight - _cdgbb.Y - _acfc
		_bcbd = _fbda - _cfbe._gcfa + _cfbe._ceef
	}
	switch _cfbe._efdg {
	case FitModeFillWidth:
		_fdcb = _gfee + _cdgbb.Width
	}
	return _gfee, _bcbd, _fdcb, _fbda
}

// Color interface represents colors in the PDF creator.
type Color interface {
	ToRGB() (float64, float64, float64)
}

// NewPolygon creates a new polygon.
func (_aacd *Creator) NewPolygon(points [][]_gga.Point) *Polygon { return _dedgg(points) }

// AddInfo is used to append a piece of invoice information in the template
// information table.
func (_ceff *Invoice) AddInfo(description, value string) (*InvoiceCell, *InvoiceCell) {
	_gedd := [2]*InvoiceCell{_ceff.newCell(description, _ceff._gca), _ceff.newCell(value, _ceff._gca)}
	_ceff._afff = append(_ceff._afff, _gedd)
	return _gedd[0], _gedd[1]
}

// DrawWithContext draws the Block using the specified drawing context.
func (_def *Block) DrawWithContext(d Drawable, ctx DrawContext) error {
	_ecd, _, _gac := d.GeneratePageBlocks(ctx)
	if _gac != nil {
		return _gac
	}
	if len(_ecd) != 1 {
		return ErrContentNotFit
	}
	for _, _dfe := range _ecd {
		if _cfe := _def.mergeBlocks(_dfe); _cfe != nil {
			return _cfe
		}
	}
	return nil
}
func (_dccae *templateProcessor) parseFloatAttr(_gddbde, _gadeg string) float64 {
	_fec.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0066\u006c\u006f\u0061t a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _gddbde, _gadeg)
	_fdbda, _ := _fg.ParseFloat(_gadeg, 64)
	return _fdbda
}

var (
	_bedfe  = _d.MustCompile("\u0028[\u005cw\u002d\u005d\u002b\u0029\u005c(\u0027\u0028.\u002b\u0029\u0027\u005c\u0029")
	_cbgfd  = _bd.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u0072\u0065a\u0074\u006f\u0072\u0020\u0069\u006e\u0073t\u0061\u006e\u0063\u0065")
	_gdbeeg = _bd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020p\u0061\u0072\u0065\u006e\u0074\u0020\u006eo\u0064\u0065")
	_dggef  = _bd.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020c\u0068\u0069\u006cd\u0020n\u006f\u0064\u0065")
	_gedbe  = _bd.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070l\u0061t\u0065 \u0072\u0065\u0073\u006f\u0075\u0072\u0063e")
)

func _bagfg(_bcdc int) *Table {
	_ddcfd := &Table{_afacb: _bcdc, _edefd: 10.0, _cddfg: []float64{}, _gecdg: []float64{}, _efbe: []*TableCell{}, _aedca: make([]int, _bcdc), _bcagb: true}
	_ddcfd.resetColumnWidths()
	return _ddcfd
}

// SetBorderColor sets the cell's border color.
func (_ccef *TableCell) SetBorderColor(col Color) {
	_ccef._ebgcd = col
	_ccef._aegdf = col
	_ccef._fcbbf = col
	_ccef._gefce = col
}
func (_adgb *Block) setOpacity(_eea float64, _gb float64) (string, error) {
	if (_eea < 0 || _eea >= 1.0) && (_gb < 0 || _gb >= 1.0) {
		return "", nil
	}
	_ge := 0
	_gc := _f.Sprintf("\u0047\u0053\u0025\u0064", _ge)
	for _adgb._dgd.HasExtGState(_bc.PdfObjectName(_gc)) {
		_ge++
		_gc = _f.Sprintf("\u0047\u0053\u0025\u0064", _ge)
	}
	_fab := _bc.MakeDict()
	if _eea >= 0 && _eea < 1.0 {
		_fab.Set("\u0063\u0061", _bc.MakeFloat(_eea))
	}
	if _gb >= 0 && _gb < 1.0 {
		_fab.Set("\u0043\u0041", _bc.MakeFloat(_gb))
	}
	_gce := _adgb._dgd.AddExtGState(_bc.PdfObjectName(_gc), _fab)
	if _gce != nil {
		return "", _gce
	}
	return _gc, nil
}
func _ffcgd(_cbaa, _fffa, _cgcg float64) (_bbafb, _ffff, _fddf, _fafbd float64) {
	if _cgcg == 0 {
		return 0, 0, _cbaa, _fffa
	}
	_ggccg := _gga.Path{Points: []_gga.Point{_gga.NewPoint(0, 0).Rotate(_cgcg), _gga.NewPoint(_cbaa, 0).Rotate(_cgcg), _gga.NewPoint(0, _fffa).Rotate(_cgcg), _gga.NewPoint(_cbaa, _fffa).Rotate(_cgcg)}}.GetBoundingBox()
	return _ggccg.X, _ggccg.Y, _ggccg.Width, _ggccg.Height
}

var (
	ColorBlack  = ColorRGBFromArithmetic(0, 0, 0)
	ColorWhite  = ColorRGBFromArithmetic(1, 1, 1)
	ColorRed    = ColorRGBFromArithmetic(1, 0, 0)
	ColorGreen  = ColorRGBFromArithmetic(0, 1, 0)
	ColorBlue   = ColorRGBFromArithmetic(0, 0, 1)
	ColorYellow = ColorRGBFromArithmetic(1, 1, 0)
)

// Style returns the style of the line.
func (_eccd *Line) Style() _gga.LineStyle { return _eccd._cedaf }

// Rows returns the total number of rows the table has.
func (_ddbg *Table) Rows() int { return _ddbg._begg }

// SetOutlineTree adds the specified outline tree to the PDF file generated
// by the creator. Adding an external outline tree disables the automatic
// generation of outlines done by the creator for the relevant components.
func (_eceb *Creator) SetOutlineTree(outlineTree *_fgd.PdfOutlineTreeNode) { _eceb._egdf = outlineTree }
func _dbffa(_cbgbc *templateProcessor, _ccag *templateNode) (interface{}, error) {
	return _cbgbc.parseBackground(_ccag)
}

// InvoiceCellProps holds all style properties for an invoice cell.
type InvoiceCellProps struct {
	TextStyle       TextStyle
	Alignment       CellHorizontalAlignment
	BackgroundColor Color
	BorderColor     Color
	BorderWidth     float64
	BorderSides     []CellBorderSide
}

func (_ecggg *templateProcessor) parseListItem(_eaageb *templateNode) (interface{}, error) {
	if _eaageb._cefd == nil {
		_ecggg.nodeLogError(_eaageb, "\u004c\u0069\u0073t\u0020\u0069\u0074\u0065m\u0020\u0070\u0061\u0072\u0065\u006e\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gdbeeg
	}
	_gabc, _cbeb := _eaageb._cefd._bedcd.(*List)
	if !_cbeb {
		_ecggg.nodeLogError(_eaageb, "\u004c\u0069s\u0074\u0020\u0069\u0074\u0065\u006d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u004cis\u0074\u002e")
		return nil, _gdbeeg
	}
	_bgbg := _dbca()
	_bgbg._edee = _gabc._edaa
	return _bgbg, nil
}

// SetTotal sets the total of the invoice.
func (_eggba *Invoice) SetTotal(value string) { _eggba._gdba[1].Value = value }

// SetLineOpacity sets the line opacity.
func (_cebga *Polyline) SetLineOpacity(opacity float64) { _cebga._ggaf = opacity }

// SetMargins sets the margins TOC line.
func (_adbdb *TOCLine) SetMargins(left, right, top, bottom float64) {
	_adbdb._afage = left
	_aagge := &_adbdb._afgdd._ccddg
	_aagge.Left = _adbdb._afage + float64(_adbdb._gfec-1)*_adbdb._ebgcdc
	_aagge.Right = right
	_aagge.Top = top
	_aagge.Bottom = bottom
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_ggec *StyledParagraph) SetTextAlignment(align TextAlignment) { _ggec._abff = align }
func _cacbe(_abbag *Block, _fgbd _fgd.PdfColor, _ceffa Color, _bbga func() Rectangle) error {
	switch _bagf := _fgbd.(type) {
	case *_fgd.PdfColorPatternType2:
		_beee, _ebbdg := _ceffa.(*LinearShading)
		if !_ebbdg {
			return _f.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u004c\u0069\u006e\u0065\u0061\u0072\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_egab := _bbga()
		_beee.SetBoundingBox(_egab._defb, _egab._cdgf, _egab._bcede, _egab._cggg)
		_febge, _cgaf := _beee.AddPatternResource(_abbag)
		if _cgaf != nil {
			return _f.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _cgaf)
		}
		_bagf.PatternName = _febge
	case *_fgd.PdfColorPatternType3:
		_aegae, _dbgf := _ceffa.(*RadialShading)
		if !_dbgf {
			return _f.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0052\u0061\u0064\u0069\u0061\u006c\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_bbbfb := _bbga()
		_aegae.SetBoundingBox(_bbbfb._defb, _bbbfb._cdgf, _bbbfb._bcede, _bbbfb._cggg)
		_agea, _ffae := _aegae.AddPatternResource(_abbag)
		if _ffae != nil {
			return _f.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _ffae)
		}
		_bagf.PatternName = _agea
	}
	return nil
}

// SetAngle sets the rotation angle of the text.
func (_bdffc *StyledParagraph) SetAngle(angle float64) { _bdffc._ffab = angle }

// GetMargins returns the margins of the line: left, right, top, bottom.
func (_dcee *Line) GetMargins() (float64, float64, float64, float64) {
	return _dcee._fggd.Left, _dcee._fggd.Right, _dcee._fggd.Top, _dcee._fggd.Bottom
}

// BorderOpacity returns the border opacity of the rectangle (0-1).
func (_gdag *Rectangle) BorderOpacity() float64 { return _gdag._dbff }

// SetColorRight sets border color for right.
func (_gaba *border) SetColorRight(col Color) { _gaba._gdbb = col }

// SetHeight sets the height of the rectangle.
func (_dgdbc *Rectangle) SetHeight(height float64) { _dgdbc._cggg = height }
func _fbe(_gcb string) string {
	_fgga := _fbf.FindAllString(_gcb, -1)
	if len(_fgga) == 0 {
		_gcb = _gcb + "\u0030"
	} else {
		_afg, _dbd := _fg.Atoi(_fgga[len(_fgga)-1])
		if _dbd != nil {
			_fec.Log.Debug("\u0045r\u0072\u006f\u0072 \u0063\u006f\u006ev\u0065rt\u0069\u006e\u0067\u0020\u0064\u0069\u0067i\u0074\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006e\u0061\u006de,\u0020f\u0061\u006c\u006c\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0062a\u0073\u0069\u0063\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a \u0025\u0076", _dbd)
			_gcb = _gcb + "\u0030"
		} else {
			_afg++
			_geg := _eg.LastIndex(_gcb, _fgga[len(_fgga)-1])
			if _geg == -1 {
				_gcb = _f.Sprintf("\u0025\u0073\u0025\u0064", _gcb[:len(_gcb)-1], _afg)
			} else {
				_gcb = _gcb[:_geg] + _fg.Itoa(_afg)
			}
		}
	}
	return _gcb
}
func _eeccc(_cdbbd float64, _gfce int) float64 {
	_gfbce := _fc.Pow10(_gfce)
	return _fc.Round(_gfbce*_cdbbd) / _gfbce
}

// SetLevelOffset sets the amount of space an indentation level occupies.
func (_dgegd *TOCLine) SetLevelOffset(levelOffset float64) {
	_dgegd._ebgcdc = levelOffset
	_dgegd._afgdd._ccddg.Left = _dgegd._afage + float64(_dgegd._gfec-1)*_dgegd._ebgcdc
}
func _aef(_ccbc string, _gdb _bc.PdfObject, _gfg *_fgd.PdfPageResources) _bc.PdfObjectName {
	_agec := _eg.TrimRightFunc(_eg.TrimSpace(_ccbc), func(_gae rune) bool { return _fe.IsNumber(_gae) })
	if _agec == "" {
		_agec = "\u0046\u006f\u006e\u0074"
	}
	_egfe := 0
	_afe := _bc.PdfObjectName(_ccbc)
	for {
		_fgda, _cgdf := _gfg.GetFontByName(_afe)
		if !_cgdf || _fgda == _gdb {
			break
		}
		_egfe++
		_afe = _bc.PdfObjectName(_f.Sprintf("\u0025\u0073\u0025\u0064", _agec, _egfe))
	}
	return _afe
}
func (_gffg *Table) wrapRow(_eefaf int, _gfbd DrawContext, _abeec float64) (bool, error) {
	if !_gffg._fgaag {
		return false, nil
	}
	var (
		_eddf  = _gffg._efbe[_eefaf]
		_bbefa = -1
		_dbgdc []*TableCell
		_gdfde float64
		_gbfda bool
		_dcffc = make([]float64, 0, len(_gffg._cddfg))
	)
	_ecea := func(_deed *TableCell, _dgccd VectorDrawable, _gfgeg bool) *TableCell {
		_abgca := *_deed
		_abgca._aafg = _dgccd
		if _gfgeg {
			_abgca._deef++
		}
		return &_abgca
	}
	_dfged := func(_febf int, _dgbgc VectorDrawable) {
		var _agcg float64 = -1
		if _dgbgc == nil {
			if _defea := _dcffc[_febf-_eefaf]; _defea > _gfbd.Height {
				_dgbgc = _gffg._efbe[_febf]._aafg
				_gffg._efbe[_febf]._aafg = nil
				_dcffc[_febf-_eefaf] = 0
				_agcg = _defea
			}
		}
		_bccee := _ecea(_gffg._efbe[_febf], _dgbgc, true)
		_dbgdc = append(_dbgdc, _bccee)
		if _agcg < 0 {
			_agcg = _bccee.height(_gfbd.Width)
		}
		if _agcg > _gdfde {
			_gdfde = _agcg
		}
	}
	for _gefa := _eefaf; _gefa < len(_gffg._efbe); _gefa++ {
		_gbag := _gffg._efbe[_gefa]
		if _eddf._deef != _gbag._deef {
			_bbefa = _gefa
			break
		}
		_gfbd.Width = _gbag.width(_gffg._cddfg, _abeec)
		_fagad := _gbag.height(_gfbd.Width)
		var _ccdb VectorDrawable
		switch _gfgga := _gbag._aafg.(type) {
		case *StyledParagraph:
			if _fagad > _gfbd.Height {
				_beeec := _gfbd
				_beeec.Height = _fc.Floor(_gfbd.Height - _gfgga._ccddg.Top - _gfgga._ccddg.Bottom - 0.5*_gfgga.getTextHeight())
				_fcgd, _ecbg, _dacaf := _gfgga.split(_beeec)
				if _dacaf != nil {
					return false, _dacaf
				}
				if _fcgd != nil && _ecbg != nil {
					_gfgga = _fcgd
					_gbag = _ecea(_gbag, _fcgd, false)
					_gffg._efbe[_gefa] = _gbag
					_ccdb = _ecbg
					_gbfda = true
				}
				_fagad = _gbag.height(_gfbd.Width)
			}
		case *Division:
			if _fagad > _gfbd.Height {
				_dfcea := _gfbd
				_dfcea.Height = _fc.Floor(_gfbd.Height - _gfgga._ddfg.Top - _gfgga._ddfg.Bottom)
				_ageb, _bgcfd := _gfgga.split(_dfcea)
				if _ageb != nil && _bgcfd != nil {
					_gfgga = _ageb
					_gbag = _ecea(_gbag, _ageb, false)
					_gffg._efbe[_gefa] = _gbag
					_ccdb = _bgcfd
					_gbfda = true
					if _ageb._agecb != nil {
						_ageb._agecb.BorderRadiusBottomLeft = 0
						_ageb._agecb.BorderRadiusBottomRight = 0
					}
					if _bgcfd._agecb != nil {
						_bgcfd._agecb.BorderRadiusTopLeft = 0
						_bgcfd._agecb.BorderRadiusTopRight = 0
					}
					_fagad = _gbag.height(_gfbd.Width)
				}
			}
		case *List:
			if _fagad > _gfbd.Height {
				_bcfe := _gfbd
				_bcfe.Height = _fc.Floor(_gfbd.Height - _gfgga._gacb.Vertical())
				_ddfbc, _bfdb := _gfgga.split(_bcfe)
				if _ddfbc != nil {
					_gfgga = _ddfbc
					_gbag = _ecea(_gbag, _ddfbc, false)
					_gffg._efbe[_gefa] = _gbag
				}
				if _bfdb != nil {
					_ccdb = _bfdb
					_gbfda = true
				}
				_fagad = _gbag.height(_gfbd.Width)
			}
		}
		_dcffc = append(_dcffc, _fagad)
		if _gbfda {
			if _dbgdc == nil {
				_dbgdc = make([]*TableCell, 0, len(_gffg._cddfg))
				for _bgag := _eefaf; _bgag < _gefa; _bgag++ {
					_dfged(_bgag, nil)
				}
			}
			_dfged(_gefa, _ccdb)
		}
	}
	var _babfb float64
	for _, _ebcf := range _dcffc {
		if _ebcf > _babfb {
			_babfb = _ebcf
		}
	}
	if _gbfda && _babfb < _gfbd.Height {
		if _bbefa < 0 {
			_bbefa = len(_gffg._efbe)
		}
		_bcba := _gffg._efbe[_bbefa-1]._deef + _gffg._efbe[_bbefa-1]._afddb - 1
		for _ceag := _bbefa; _ceag < len(_gffg._efbe); _ceag++ {
			_gffg._efbe[_ceag]._deef++
		}
		_gffg._efbe = append(_gffg._efbe[:_bbefa], append(_dbgdc, _gffg._efbe[_bbefa:]...)...)
		_gffg._gecdg = append(_gffg._gecdg[:_bcba], append([]float64{_gdfde}, _gffg._gecdg[_bcba:]...)...)
		_gffg._gecdg[_eddf._deef+_eddf._afddb-2] = _babfb
	}
	return _gbfda, nil
}
func (_bdaf *templateProcessor) parseImage(_bfcd *templateNode) (interface{}, error) {
	var _cebbf string
	for _, _egba := range _bfcd._adfdg.Attr {
		_cccf := _egba.Value
		switch _bffd := _egba.Name.Local; _bffd {
		case "\u0073\u0072\u0063":
			_cebbf = _cccf
		}
	}
	_gadec, _cagd := _bdaf.loadImageFromSrc(_cebbf)
	if _cagd != nil {
		return nil, _cagd
	}
	for _, _fbgg := range _bfcd._adfdg.Attr {
		_adffb := _fbgg.Value
		switch _afbcb := _fbgg.Name.Local; _afbcb {
		case "\u0061\u006c\u0069g\u006e":
			_gadec.SetHorizontalAlignment(_bdaf.parseHorizontalAlignmentAttr(_afbcb, _adffb))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_gadec.SetOpacity(_bdaf.parseFloatAttr(_afbcb, _adffb))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_gcfdg := _bdaf.parseMarginAttr(_afbcb, _adffb)
			_gadec.SetMargins(_gcfdg.Left, _gcfdg.Right, _gcfdg.Top, _gcfdg.Bottom)
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_gadec.SetFitMode(_bdaf.parseFitModeAttr(_afbcb, _adffb))
		case "\u0078":
			_gadec.SetPos(_bdaf.parseFloatAttr(_afbcb, _adffb), _gadec._dbccc)
		case "\u0079":
			_gadec.SetPos(_gadec._afggb, _bdaf.parseFloatAttr(_afbcb, _adffb))
		case "\u0077\u0069\u0064t\u0068":
			_gadec.SetWidth(_bdaf.parseFloatAttr(_afbcb, _adffb))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_gadec.SetHeight(_bdaf.parseFloatAttr(_afbcb, _adffb))
		case "\u0061\u006e\u0067l\u0065":
			_gadec.SetAngle(_bdaf.parseFloatAttr(_afbcb, _adffb))
		case "\u0073\u0072\u0063":
			break
		default:
			_bdaf.nodeLogDebug(_bfcd, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0069m\u0061\u0067\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _afbcb)
		}
	}
	return _gadec, nil
}

// Height returns the height of the Paragraph. The height is calculated based on the input text and how it is wrapped
// within the container. Does not include Margins.
func (_faaa *StyledParagraph) Height() float64 {
	_faaa.wrapText()
	var _eecc float64
	for _, _aagec := range _faaa._bdfce {
		var _ddaeg float64
		for _, _aeca := range _aagec {
			_bdbec := _faaa._fgef * _aeca.Style.FontSize
			if _bdbec > _ddaeg {
				_ddaeg = _bdbec
			}
		}
		_eecc += _ddaeg
	}
	return _eecc
}

// SetMarkedContentID sets marked content ID.
func (_gbed *Polyline) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_gbed._egfd = &mcid
	_eaab := _fgd.NewKDictionary()
	_eaab.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_eaab.K = _bc.MakeInteger(mcid)
	return _eaab
}
func (_geea *StyledParagraph) appendChunk(_baaa *TextChunk) *TextChunk {
	_geea._ecec = append(_geea._ecec, _baaa)
	_geea.wrapText()
	return _baaa
}

// Horizontal returns total horizontal (left + right) margin.
func (_aagg *Margins) Horizontal() float64 { return _aagg.Left + _aagg.Right }

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_bbde *shading) SetBackgroundColor(backgroundColor Color) { _bbde._feac = backgroundColor }

type templateProcessor struct {
	creator *Creator
	_bcfcf  []byte
	_cbcec  *TemplateOptions
	_ccfe   componentRenderer
	_edcgd  string
}

// GetCoords returns the (x1, y1), (x2, y2) points defining the Line.
func (_bgcb *Line) GetCoords() (float64, float64, float64, float64) {
	return _bgcb._gbcgde, _bgcb._gcfa, _bgcb._cfge, _bgcb._ceef
}

// DashPattern returns the dash pattern of the line.
func (_gfab *Line) DashPattern() (_dbded []int64, _dede int64) { return _gfab._cgecc, _gfab._dcebc }
func (_eagd *Invoice) generateInformationBlocks(_daca DrawContext) ([]*Block, DrawContext, error) {
	_dafg := _ddge(_eagd._cgce)
	_dafg.SetMargins(0, 0, 0, 20)
	_bdgd := _eagd.drawAddress(_eagd._ddaa)
	_bdgd = append(_bdgd, _dafg)
	_bdgd = append(_bdgd, _eagd.drawAddress(_eagd._eeeb)...)
	_faaee := _dcbc()
	for _, _ebab := range _bdgd {
		_faaee.Add(_ebab)
	}
	_fcfdf := _eagd.drawInformation()
	_ccff := _bagfg(2)
	_ccff.SetMargins(0, 0, 25, 0)
	_faea := _ccff.NewCell()
	_faea.SetIndent(0)
	_faea.SetContent(_faaee)
	_faea = _ccff.NewCell()
	_faea.SetContent(_fcfdf)
	return _ccff.GeneratePageBlocks(_daca)
}
func (_effcg *StyledParagraph) getTextWidth() float64 {
	var _gdgcd float64
	_cadg := len(_effcg._ecec)
	for _aggca, _cedba := range _effcg._ecec {
		_cfag := &_cedba.Style
		_fdaf := len(_cedba.Text)
		for _eeef, _gfaa := range _cedba.Text {
			if _gfaa == '\u000A' {
				continue
			}
			_edebg, _gcead := _cfag.Font.GetRuneMetrics(_gfaa)
			if !_gcead {
				_fec.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _gfaa)
				return -1
			}
			_gdgcd += _cfag.FontSize * _edebg.Wx * _cfag.horizontalScale()
			if _gfaa != ' ' && (_aggca != _cadg-1 || _eeef != _fdaf-1) {
				_gdgcd += _cfag.CharSpacing * 1000.0
			}
		}
	}
	return _gdgcd
}

// SetBorderWidth sets the border width.
func (_efade *Polygon) SetBorderWidth(borderWidth float64) { _efade._acdc.BorderWidth = borderWidth }

// Write output of creator to io.Writer interface.
func (_baedd *Creator) Write(ws _gg.Writer) error {
	if _eddbb := _baedd.Finalize(); _eddbb != nil {
		return _eddbb
	}
	_abc := ""
	if _gfed, _aaga := ws.(*_b.File); _aaga {
		_abc = _gfed.Name()
	}
	_ceea := _fgd.NewPdfWriter()
	_ceea.SetOptimizer(_baedd._caec)
	_ceea.SetFileName(_abc)
	if _baedd._ecc != nil {
		_abebcb := _ceea.SetForms(_baedd._ecc)
		if _abebcb != nil {
			_fec.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _abebcb)
			return _abebcb
		}
	}
	if _baedd._egdf != nil {
		_ceea.AddOutlineTree(_baedd._egdf)
	} else if _baedd._cbba != nil && _baedd.AddOutlines {
		_ceea.AddOutlineTree(&_baedd._cbba.ToPdfOutline().PdfOutlineTreeNode)
	}
	if _baedd._deg != nil {
		if _abebd := _ceea.SetPageLabels(_baedd._deg); _abebd != nil {
			_fec.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020C\u006f\u0075\u006c\u0064 no\u0074 s\u0065\u0074\u0020\u0070\u0061\u0067\u0065 l\u0061\u0062\u0065\u006c\u0073\u003a\u0020%\u0076", _abebd)
			return _abebd
		}
	}
	if _baedd._dad != nil {
		for _, _ffac := range _baedd._dad {
			_dabc := _ffac.SubsetRegistered()
			if _dabc != nil {
				_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020s\u0075\u0062\u0073\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _dabc)
				return _dabc
			}
		}
	}
	if _baedd._gacd != nil {
		_geed := _baedd._gacd(&_ceea)
		if _geed != nil {
			_fec.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _geed)
			return _geed
		}
	}
	for _fffe, _aecdc := range _baedd._egfb {
		_fcabd := _ceea.AddPage(_aecdc)
		if _fcabd != nil {
			_fec.Log.Error("\u0046\u0061\u0069\u006ced\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0050\u0061\u0067\u0065\u003a\u0020%\u0076", _fcabd)
			return _fcabd
		}
		if _baedd._cfgc != nil {
			_baac := _baedd._cfgc.K
			_ecg, _ebbg := _ceea.GetPageIndirectObject(_fffe)
			if _ebbg != nil {
				_fec.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043\u006fu\u006c\u0064\u0020n\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061ge\u0020\u0069\u006ed\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074 \u0025\u0076", _ebbg)
			}
			var _egfcb func(_abdf *_fgd.KDict)
			_egfcb = func(_fgbf *_fgd.KDict) {
				if _fgbf == nil {
					return
				}
				if _fgbf.GetPageNumber()-1 == int64(_fffe) {
					_fgbf.SetPage(_ecg)
				}
				for _, _eefb := range _fgbf.GetChildren() {
					if _efc := _eefb.GetKDict(); _efc != nil {
						_egfcb(_efc)
					}
				}
			}
			for _, _aagc := range _baac {
				_egfcb(_aagc)
			}
		}
	}
	if _baedd._cfgc != nil {
		if _bfca := _ceea.SetCatalogStructTreeRoot(_baedd._cfgc.ToPdfObject()); _bfca != nil {
			_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0073\u0065\u0074 \u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006ft\u003a\u0020\u0025\u0076", _bfca)
			return _bfca
		}
	}
	if _baedd._bbcg != nil {
		if _cfceea := _ceea.SetCatalogViewerPreferences(_baedd._bbcg.ToPdfObject()); _cfceea != nil {
			_fec.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074\u0020\u0056\u0069\u0065\u0077\u0065\u0072\u0050\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076", _cfceea)
			return _cfceea
		}
	}
	if _baedd._eeae != "" {
		if _ebfd := _ceea.SetCatalogLanguage(_bc.MakeString(_baedd._eeae)); _ebfd != nil {
			_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0073\u0065t\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _ebfd)
			return _ebfd
		}
	}
	_aega := _ceea.Write(ws)
	if _aega != nil {
		return _aega
	}
	return nil
}

// EnableFontSubsetting enables font subsetting for `font` when the creator output is written to file.
// Embeds only the subset of the runes/glyphs that are actually used to display the file.
// Subsetting can reduce the size of fonts significantly.
func (_ega *Creator) EnableFontSubsetting(font *_fgd.PdfFont) { _ega._dad = append(_ega._dad, font) }
func (_ffeab *Ellipse) applyFitMode(_gababc float64) {
	_gababc -= _ffeab._ddbd.Left + _ffeab._ddbd.Right
	switch _ffeab._eddea {
	case FitModeFillWidth:
		_ffeab.ScaleToWidth(_gababc)
	}
}

// SetHeading sets the text and the style of the heading of the TOC component.
func (_caefe *TOC) SetHeading(text string, style TextStyle) {
	_babg := _caefe.Heading()
	_babg.Reset()
	_afgeb := _babg.Append(text)
	_afgeb.Style = style
}

// SetBorderOpacity sets the border opacity of the ellipse.
func (_eead *Ellipse) SetBorderOpacity(opacity float64) { _eead._fbcf = opacity }

// EnableWordWrap sets the paragraph word wrap flag.
func (_eccgc *StyledParagraph) EnableWordWrap(val bool) { _eccgc._eadc = val }

const (
	DefaultHorizontalScaling = 100
)

// RadialShading holds information that will be used to render a radial shading.
type RadialShading struct {
	_dafd  *shading
	_bfgf  *_fgd.PdfRectangle
	_effe  AnchorPoint
	_eccda float64
	_dagc  float64
	_aeeeb float64
	_babf  float64
}

// SetLineNumberStyle sets the style for the numbers part of all new lines
// of the table of contents.
func (_fadcf *TOC) SetLineNumberStyle(style TextStyle) { _fadcf._cggbgd = style }

// SetTitle sets the title of the invoice.
func (_edcg *Invoice) SetTitle(title string) { _edcg._dgfaa = title }

// SetCoords sets the upper left corner coordinates of the rectangle.
func (_bddd *Rectangle) SetCoords(x, y float64) { _bddd._defb = x; _bddd._cdgf = y }
func _cffe(_efeda string, _cbdgc, _gcbcg TextStyle) *TOC {
	_fcae := _gcbcg
	_fcae.FontSize = 14
	_acgbc := _ddge(_fcae)
	_acgbc.SetEnableWrap(true)
	_acgbc.SetTextAlignment(TextAlignmentLeft)
	_acgbc.SetMargins(0, 0, 0, 5)
	_gbdgd := _acgbc.Append(_efeda)
	_gbdgd.Style = _fcae
	return &TOC{_defae: _acgbc, _ccfec: []*TOCLine{}, _cggbgd: _cbdgc, _fefdb: _cbdgc, _fadge: _cbdgc, _caeab: _cbdgc, _begag: "\u002e", _eaadf: 10, _becec: Margins{0, 0, 2, 2}, _gabag: PositionRelative, _eddd: _cbdgc, _ggabc: true}
}

// ColorRGBFromHex converts color hex code to rgb color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ffffff -> (1,1,1) white.
func ColorRGBFromHex(hexStr string) Color {
	_ccbd := rgbColor{}
	if (len(hexStr) != 4 && len(hexStr) != 7) || hexStr[0] != '#' {
		_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _ccbd
	}
	var _dddd, _cbcf, _fgde int
	if len(hexStr) == 4 {
		var _eada, _ecdec, _ddf int
		_dcba, _bcg := _f.Sscanf(hexStr, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_eada, &_ecdec, &_ddf)
		if _bcg != nil {
			_fec.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _bcg)
			return _ccbd
		}
		if _dcba != 3 {
			_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _ccbd
		}
		_dddd = _eada*16 + _eada
		_cbcf = _ecdec*16 + _ecdec
		_fgde = _ddf*16 + _ddf
	} else {
		_ebed, _cffb := _f.Sscanf(hexStr, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_dddd, &_cbcf, &_fgde)
		if _cffb != nil {
			_fec.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _ccbd
		}
		if _ebed != 3 {
			_fec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _ebed)
			return _ccbd
		}
	}
	_ccbeg := float64(_dddd) / 255.0
	_fbfe := float64(_cbcf) / 255.0
	_efaf := float64(_fgde) / 255.0
	_ccbd._ead = _ccbeg
	_ccbd._ddd = _fbfe
	_ccbd._decg = _efaf
	return _ccbd
}

// Height returns the Block's height.
func (_faf *Block) Height() float64 { return _faf._ace }
func (_ceab *templateProcessor) parseChapterHeading(_bddgg *templateNode) (interface{}, error) {
	if _bddgg._cefd == nil {
		_ceab.nodeLogError(_bddgg, "\u0043\u0068a\u0070\u0074\u0065\u0072 \u0068\u0065a\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _gdbeeg
	}
	_faedb, _ebgfc := _bddgg._cefd._bedcd.(*Chapter)
	if !_ebgfc {
		_ceab.nodeLogError(_bddgg, "\u0043h\u0061\u0070t\u0065\u0072\u0020h\u0065\u0061\u0064\u0069\u006e\u0067\u0020p\u0061\u0072\u0065\u006e\u0074\u0020(\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u002e", _bddgg._cefd._bedcd)
		return nil, _gdbeeg
	}
	_accef := _faedb.GetHeading()
	if _, _cbdbf := _ceab.parseParagraph(_bddgg, _accef); _cbdbf != nil {
		return nil, _cbdbf
	}
	return _accef, nil
}

// CellBorderStyle defines the table cell's border style.
type CellBorderStyle int

// CreateTableOfContents sets a function to generate table of contents.
func (_ffg *Creator) CreateTableOfContents(genTOCFunc func(_dgaa *TOC) error) { _ffg._ggd = genTOCFunc }

// SetText sets the text content of the Paragraph.
func (_bdfae *Paragraph) SetText(text string) { _bdfae._egea = text }

// SetMarkedContentID sets the marked content identifier.
func (_ecab *Division) SetMarkedContentID(id int64) *_fgd.KDict { return nil }

// SetMargins sets the margins of the chart component.
func (_bgg *Chart) SetMargins(left, right, top, bottom float64) {
	_bgg._bgfg.Left = left
	_bgg._bgfg.Right = right
	_bgg._bgfg.Top = top
	_bgg._bgfg.Bottom = bottom
}

// SetBackgroundColor sets the cell's background color.
func (_cgfg *TableCell) SetBackgroundColor(col Color) { _cgfg._cbgd = col }
func _egddc(_gcff *_b.File) ([]*_fgd.PdfPage, error) {
	_fgedg, _fcece := _fgd.NewPdfReader(_gcff)
	if _fcece != nil {
		return nil, _fcece
	}
	_gacce, _fcece := _fgedg.GetNumPages()
	if _fcece != nil {
		return nil, _fcece
	}
	var _dggfe []*_fgd.PdfPage
	for _fbafc := 0; _fbafc < _gacce; _fbafc++ {
		_aafef, _ecdgc := _fgedg.GetPage(_fbafc + 1)
		if _ecdgc != nil {
			return nil, _ecdgc
		}
		_dggfe = append(_dggfe, _aafef)
	}
	return _dggfe, nil
}

// NewBlockFromPage creates a Block from a PDF Page.  Useful for loading template pages as blocks
// from a PDF document and additional content with the creator.
func NewBlockFromPage(page *_fgd.PdfPage) (*Block, error) {
	_egc := &Block{}
	_ef, _ee := page.GetAllContentStreams()
	if _ee != nil {
		return nil, _ee
	}
	_be := _dg.NewContentStreamParser(_ef)
	_acec, _ee := _be.Parse()
	if _ee != nil {
		return nil, _ee
	}
	_acec.WrapIfNeeded()
	_egc._cb = _acec
	if page.Resources != nil {
		_egc._dgd = page.Resources
	} else {
		_egc._dgd = _fgd.NewPdfPageResources()
	}
	_adg, _ee := page.GetMediaBox()
	if _ee != nil {
		return nil, _ee
	}
	if _adg.Llx != 0 || _adg.Lly != 0 {
		_egc.translate(-_adg.Llx, _adg.Lly)
	}
	_egc._da = _adg.Urx - _adg.Llx
	_egc._ace = _adg.Ury - _adg.Lly
	if page.Rotate != nil {
		_egc._df = -float64(*page.Rotate)
	}
	return _egc, nil
}

// NoteHeadingStyle returns the style properties used to render the heading of
// the invoice note sections.
func (_egeg *Invoice) NoteHeadingStyle() TextStyle { return _egeg._bgbeg }

// SetPos sets the position of the chart to the specified coordinates.
// This method sets the chart to use absolute positioning.
func (_bcefb *Chart) SetPos(x, y float64) {
	_bcefb._cggb = PositionAbsolute
	_bcefb._dbf = x
	_bcefb._gbcb = y
}
func (_geefd *StyledParagraph) getMaxLineWidth() float64 {
	if _geefd._bdfce == nil || (_geefd._bdfce != nil && len(_geefd._bdfce) == 0) {
		_geefd.wrapText()
	}
	var _geac float64
	for _, _fcabb := range _geefd._bdfce {
		_fcca := _geefd.getTextLineWidth(_fcabb)
		if _fcca > _geac {
			_geac = _fcca
		}
	}
	return _geac
}
func (_egg *Block) SetMarkedContentID(id int64) *_fgd.KDict { return nil }

// SetHorizontalAlignment sets the horizontal alignment of the image.
func (_ccfc *Image) SetHorizontalAlignment(alignment HorizontalAlignment) { _ccfc._fbbb = alignment }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_bccg *List) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _fcded float64
	var _gbec []*StyledParagraph
	for _, _dadba := range _bccg._ffegf {
		_bece := _ddge(_bccg._fedb)
		_bece.SetEnableWrap(false)
		_bece.SetTextAlignment(TextAlignmentRight)
		_bece.Append(_dadba._edee.Text).Style = _dadba._edee.Style
		_dfee := _bece.getTextWidth() / 1000.0 / ctx.Width
		if _fcded < _dfee {
			_fcded = _dfee
		}
		_gbec = append(_gbec, _bece)
	}
	_dcbaf := _bagfg(2)
	_dcbaf.SetColumnWidths(_fcded, 1-_fcded)
	_dcbaf.SetMargins(_bccg._gacb.Left+_bccg._aedbb, _bccg._gacb.Right, _bccg._gacb.Top, _bccg._gacb.Bottom)
	_dcbaf.EnableRowWrap(true)
	for _fbbbe, _effg := range _bccg._ffegf {
		_dggad := _dcbaf.NewCell()
		_dggad.SetIndent(0)
		_dggad.SetContent(_gbec[_fbbbe])
		_dggad = _dcbaf.NewCell()
		_dggad.SetIndent(0)
		_dggad.SetContent(_effg._cdeda)
	}
	return _dcbaf.GeneratePageBlocks(ctx)
}
func (_addeb *StyledParagraph) split(_fegdb DrawContext) (_ggada, _caaee *StyledParagraph, _aafa error) {
	if _aafa = _addeb.wrapChunks(false); _aafa != nil {
		return nil, nil, _aafa
	}
	if len(_addeb._bdfce) == 1 && _addeb._fgef > _fegdb.Height {
		return _addeb, nil, nil
	}
	_fbad := func(_ggfdc []*TextChunk, _bfbd []*TextChunk) []*TextChunk {
		if len(_bfbd) == 0 {
			return _ggfdc
		}
		_cacge := len(_ggfdc)
		if _cacge == 0 {
			return append(_ggfdc, _bfbd...)
		}
		if _ggfdc[_cacge-1].Style == _bfbd[0].Style {
			_ggfdc[_cacge-1].Text += _bfbd[0].Text
		} else {
			_ggfdc = append(_ggfdc, _bfbd[0])
		}
		return append(_ggfdc, _bfbd[1:]...)
	}
	_ffed := func(_feee *StyledParagraph, _cecf []*TextChunk) *StyledParagraph {
		if len(_cecf) == 0 {
			return nil
		}
		_eefa := *_feee
		_eefa._ecec = _cecf
		return &_eefa
	}
	var (
		_bbca  float64
		_ggaff []*TextChunk
		_dgfb  []*TextChunk
	)
	for _, _fgbga := range _addeb._bdfce {
		var _edff float64
		_cccb := make([]*TextChunk, 0, len(_fgbga))
		for _, _aefd := range _fgbga {
			if _abbaa := _aefd.Style.FontSize; _abbaa > _edff {
				_edff = _abbaa
			}
			_cccb = append(_cccb, _aefd.clone())
		}
		_edff *= _addeb._fgef
		if _addeb._fceb.IsRelative() {
			if _bbca+_edff > _fegdb.Height {
				_dgfb = _fbad(_dgfb, _cccb)
			} else {
				_ggaff = _fbad(_ggaff, _cccb)
			}
		}
		_bbca += _edff
	}
	_addeb._bdfce = nil
	if len(_dgfb) == 0 {
		return _addeb, nil, nil
	}
	return _ffed(_addeb, _ggaff), _ffed(_addeb, _dgfb), nil
}
func (_gdcbf *templateProcessor) parseChapter(_faagb *templateNode) (interface{}, error) {
	_acbga := _gdcbf.creator.NewChapter
	if _faagb._cefd != nil {
		if _acbfa, _dgfd := _faagb._cefd._bedcd.(*Chapter); _dgfd {
			_acbga = _acbfa.NewSubchapter
		}
	}
	_gaecc := _acbga("")
	for _, _eeded := range _faagb._adfdg.Attr {
		_ffdfd := _eeded.Value
		switch _fgdb := _eeded.Name.Local; _fgdb {
		case "\u0073\u0068\u006f\u0077\u002d\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067":
			_gaecc.SetShowNumbering(_gdcbf.parseBoolAttr(_fgdb, _ffdfd))
		case "\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u002d\u0069n\u002d\u0074\u006f\u0063":
			_gaecc.SetIncludeInTOC(_gdcbf.parseBoolAttr(_fgdb, _ffdfd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_dfdc := _gdcbf.parseMarginAttr(_fgdb, _ffdfd)
			_gaecc.SetMargins(_dfdc.Left, _dfdc.Right, _dfdc.Top, _dfdc.Bottom)
		default:
			_gdcbf.nodeLogDebug(_faagb, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _fgdb)
		}
	}
	return _gaecc, nil
}
func (_faefd *List) split(_adcb DrawContext) (_bdbc, _fgab *List) {
	var (
		_ecefb        float64
		_abbb, _eaaga []*listItem
	)
	_bfcc := _adcb.Width - _faefd._gacb.Horizontal() - _faefd._aedbb - _faefd.markerWidth()
	_fcac := _faefd.markerWidth()
	for _fbbc, _ebae := range _faefd._ffegf {
		_afcg := _ebae.ctxHeight(_bfcc)
		_ecefb += _afcg
		if _ecefb <= _adcb.Height {
			_abbb = append(_abbb, _ebae)
		} else {
			switch _eebe := _ebae._cdeda.(type) {
			case *List:
				_ccad := _adcb
				_ccad.Height = _fc.Floor(_afcg - (_ecefb - _adcb.Height))
				_fafcd, _fgccb := _eebe.split(_ccad)
				if _fafcd != nil {
					_cecd := _dbca()
					_cecd._edee = _ebae._edee
					_cecd._cdeda = _fafcd
					_abbb = append(_abbb, _cecd)
				}
				if _fgccb != nil {
					_gafg := _eebe._edaa.Style.FontSize
					_adbd, _cfbc := _eebe._edaa.Style.Font.GetRuneMetrics(' ')
					if _cfbc {
						_gafg = _eebe._edaa.Style.FontSize * _adbd.Wx * _eebe._edaa.Style.horizontalScale() / 1000.0
					}
					_cdgd := _eg.Repeat("\u0020", int(_fcac/_gafg))
					_dbead := _dbca()
					_dbead._edee = *NewTextChunk(_cdgd, _eebe._edaa.Style)
					_dbead._cdeda = _fgccb
					_eaaga = append(_eaaga, _dbead)
					_eaaga = append(_eaaga, _faefd._ffegf[_fbbc+1:]...)
				}
			default:
				_eaaga = _faefd._ffegf[_fbbc:]
			}
			if len(_eaaga) > 0 {
				break
			}
		}
	}
	if len(_abbb) > 0 {
		_bdbc = _cfbf(_faefd._fedb)
		*_bdbc = *_faefd
		_bdbc._ffegf = _abbb
	}
	if len(_eaaga) > 0 {
		_fgab = _cfbf(_faefd._fedb)
		*_fgab = *_faefd
		_fgab._ffegf = _eaaga
	}
	return _bdbc, _fgab
}
func _cdfg(_abfc *_dd.GraphicSVG) (*GraphicSVG, error) {
	return &GraphicSVG{_dgccb: _abfc, _ddcg: PositionRelative, _aagce: Margins{Top: 10, Bottom: 10}}, nil
}

// SetPdfWriterAccessFunc sets a PdfWriter access function/hook.
// Exposes the PdfWriter just prior to writing the PDF.  Can be used to encrypt the output PDF, etc.
//
// Example of encrypting with a user/owner password "password"
// Prior to calling c.WriteFile():
//
//	c.SetPdfWriterAccessFunc(func(w *model.PdfWriter) error {
//		userPass := []byte("password")
//		ownerPass := []byte("password")
//		err := w.Encrypt(userPass, ownerPass, nil)
//		return err
//	})
func (_gcee *Creator) SetPdfWriterAccessFunc(pdfWriterAccessFunc func(_bgcf *_fgd.PdfWriter) error) {
	_gcee._gacd = pdfWriterAccessFunc
}

// SetTextVerticalAlignment sets the vertical alignment of the text within the
// bounds of the styled paragraph.
//
// Note: Currently Styled Paragraph doesn't support TextVerticalAlignmentBottom
// as that option only used for aligning text chunks.
//
// In order to change the vertical alignment of individual text chunks, use TextChunk.VerticalAlignment.
func (_abeee *StyledParagraph) SetTextVerticalAlignment(align TextVerticalAlignment) {
	_abeee._ddec = align
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_ceec *Paragraph) GetMargins() (float64, float64, float64, float64) {
	return _ceec._eebg.Left, _ceec._eebg.Right, _ceec._eebg.Top, _ceec._eebg.Bottom
}

// SetMarkedContentID sets the marked content identifier.
func (_cabg *Polygon) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_cabg._fdbg = &mcid
	_fdac := _fgd.NewKDictionary()
	_fdac.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_fdac.K = _bc.MakeInteger(mcid)
	return _fdac
}

// SetMargins sets the margins for the Image (in relative mode): left, right, top, bottom.
func (_agdd *Image) SetMargins(left, right, top, bottom float64) {
	_agdd._cdbc.Left = left
	_agdd._cdbc.Right = right
	_agdd._cdbc.Top = top
	_agdd._cdbc.Bottom = bottom
}
func _edebe(_egec *Block, _ddba *Image, _gedg DrawContext) (DrawContext, error) {
	_face := _gedg
	_eaee := 1
	_adccd := _bc.PdfObjectName(_f.Sprintf("\u0049\u006d\u0067%\u0064", _eaee))
	for _egec._dgd.HasXObjectByName(_adccd) {
		_eaee++
		_adccd = _bc.PdfObjectName(_f.Sprintf("\u0049\u006d\u0067%\u0064", _eaee))
	}
	_bggac := _egec._dgd.SetXObjectImageByNameLazy(_adccd, _ddba._gacc, _ddba._faeda)
	if _bggac != nil {
		return _gedg, _bggac
	}
	_dagf := 0
	_cfgde := _bc.PdfObjectName(_f.Sprintf("\u0047\u0053\u0025\u0064", _dagf))
	for _egec._dgd.HasExtGState(_cfgde) {
		_dagf++
		_cfgde = _bc.PdfObjectName(_f.Sprintf("\u0047\u0053\u0025\u0064", _dagf))
	}
	_gadd := _bc.MakeDict()
	_gadd.Set("\u0042\u004d", _bc.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	if _ddba._ddea < 1.0 {
		_gadd.Set("\u0043\u0041", _bc.MakeFloat(_ddba._ddea))
		_gadd.Set("\u0063\u0061", _bc.MakeFloat(_ddba._ddea))
	}
	_bggac = _egec._dgd.AddExtGState(_cfgde, _bc.MakeIndirectObject(_gadd))
	if _bggac != nil {
		return _gedg, _bggac
	}
	_ffdc := _ddba.Width()
	_gcbad := _ddba.Height()
	_, _aagaa := _ddba.rotatedSize()
	_egcg := _gedg.X
	_fggc := _gedg.PageHeight - _gedg.Y - _gcbad
	if _ddba._accd.IsRelative() {
		_fggc -= (_aagaa - _gcbad) / 2
		switch _ddba._fbbb {
		case HorizontalAlignmentCenter:
			_egcg += (_gedg.Width - _ffdc) / 2
		case HorizontalAlignmentRight:
			_egcg = _gedg.PageWidth - _gedg.Margins.Right - _ddba._cdbc.Right - _ffdc
		}
	}
	_begf := _ddba._fbbf
	_bgeg := _dg.NewContentCreator()
	if _ddba._gebca != nil {
		_bgeg.Add_BDC(*_bc.MakeName(_fgd.StructureTypeFigure), map[string]_bc.PdfObject{"\u004d\u0043\u0049\u0044": _bc.MakeInteger(*_ddba._gebca)})
	}
	_bgeg.Add_gs(_cfgde)
	_bgeg.Translate(_egcg, _fggc)
	if _begf != 0 {
		_bgeg.Translate(_ffdc/2, _gcbad/2)
		_bgeg.RotateDeg(_begf)
		_bgeg.Translate(-_ffdc/2, -_gcbad/2)
	}
	_bgeg.Scale(_ffdc, _gcbad).Add_Do(_adccd)
	if _ddba._gebca != nil {
		_bgeg.Add_EMC()
	}
	_fdfc := _bgeg.Operations()
	_fdfc.WrapIfNeeded()
	_egec.addContents(_fdfc)
	if _ddba._accd.IsRelative() {
		_gedg.Y += _aagaa
		_gedg.Height -= _aagaa
		return _gedg, nil
	}
	return _face, nil
}

// SetFillColor sets background color for border.
func (_fce *border) SetFillColor(col Color) { _fce._aff = col }

// Inline returns whether the inline mode of the division is active.
func (_ffcb *Division) Inline() bool { return _ffcb._fdb }

// SetBorderColor sets the border color for the path.
func (_dcce *FilledCurve) SetBorderColor(color Color) { _dcce._ccec = color }

// AddAnnotation adds an annotation to the current block.
// The annotation will be added to the page the block will be rendered on.
func (_acf *Block) AddAnnotation(annotation *_fgd.PdfAnnotation) {
	for _, _dga := range _acf._ed {
		if _dga == annotation {
			return
		}
	}
	_acf._ed = append(_acf._ed, annotation)
}

// SetBorderRadius sets the radius of the background corners.
func (_ff *Background) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_ff.BorderRadiusTopLeft = topLeft
	_ff.BorderRadiusTopRight = topRight
	_ff.BorderRadiusBottomLeft = bottomLeft
	_ff.BorderRadiusBottomRight = bottomRight
}

// ScaleToHeight scales the ellipse to the specified height. The width of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_bge *Ellipse) ScaleToHeight(h float64) {
	_babc := _bge._eefg / _bge._beb
	_bge._beb = h
	_bge._eefg = h * _babc
}

var (
	ErrContentNotFit = _bd.New("\u0043\u0061\u006e\u006e\u006ft\u0020\u0066\u0069\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020i\u006e\u0074\u006f\u0020\u0061\u006e\u0020\u0065\u0078\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0073\u0070\u0061\u0063\u0065")
)

// Padding returns the padding of the component.
func (_cbd *Division) Padding() (_affed, _efd, _fecg, _abac float64) {
	return _cbd._gbga.Left, _cbd._gbga.Right, _cbd._gbga.Top, _cbd._gbga.Bottom
}

// MultiRowCell makes a new cell with the specified row span and inserts it
// into the table at the current position.
func (_ffcab *Table) MultiRowCell(rowspan int) *TableCell { return _ffcab.MultiCell(rowspan, 1) }

type rgbColor struct{ _ead, _ddd, _decg float64 }

const (
	PositionRelative Positioning = iota
	PositionAbsolute
)

// SetRowPosition sets cell row position.
func (_bega *TableCell) SetRowPosition(row int) { _bega._deef = row }

// TextAlignment options for paragraph.
type TextAlignment int

// NewPageBreak create a new page break.
func (_gabb *Creator) NewPageBreak() *PageBreak { return _gbcfe() }
func (_dfde *Creator) wrapPageIfNeeded(_fgf *_fgd.PdfPage) (*_fgd.PdfPage, error) {
	_dccfe, _dgaf := _fgf.GetAllContentStreams()
	if _dgaf != nil {
		return nil, _dgaf
	}
	_cbef := _dg.NewContentStreamParser(_dccfe)
	_febcd, _dgaf := _cbef.Parse()
	if _dgaf != nil {
		return nil, _dgaf
	}
	if !_febcd.HasUnclosedQ() {
		return nil, nil
	}
	_febcd.WrapIfNeeded()
	_aebd, _dgaf := _bc.MakeStream(_febcd.Bytes(), _bc.NewFlateEncoder())
	if _dgaf != nil {
		return nil, _dgaf
	}
	_fgf.Contents = _bc.MakeArray(_aebd)
	return _fgf, nil
}

// GetCoords returns coordinates of border.
func (_bbfe *border) GetCoords() (float64, float64) { return _bbfe._bad, _bbfe._abda }

// Length calculates and returns the length of the line.
func (_aafca *Line) Length() float64 {
	return _fc.Sqrt(_fc.Pow(_aafca._cfge-_aafca._gbcgde, 2.0) + _fc.Pow(_aafca._ceef-_aafca._gcfa, 2.0))
}

// Width returns the width of the specified text chunk.
func (_fbfg *TextChunk) Width() float64 {
	var (
		_gcbg  float64
		_edbag = _fbfg.Style
	)
	for _, _dfedf := range _fbfg.Text {
		_dbfd, _bedde := _edbag.Font.GetRuneMetrics(_dfedf)
		if !_bedde {
			_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _dfedf, _dfedf, _edbag.Font.BaseFont(), _edbag.Font.Subtype())
			_fec.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _edbag.Font)
			_fec.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _edbag.Font.Encoder())
		}
		_egde := _edbag.FontSize * _dbfd.Wx
		_dbed := _egde
		if _dfedf != ' ' {
			_dbed = _egde + _edbag.CharSpacing*1000.0
		}
		_gcbg += _dbed
	}
	return _gcbg / 1000.0
}
func (_efdcb *Table) moveToNextAvailableCell() int {
	_bedee := (_efdcb._ecgdc-1)%(_efdcb._afacb) + 1
	for {
		if _bedee-1 >= len(_efdcb._aedca) {
			if _efdcb._aedca[0] == 0 {
				return _bedee
			}
			_bedee = 1
		} else if _efdcb._aedca[_bedee-1] == 0 {
			return _bedee
		}
		_efdcb._ecgdc++
		_efdcb._aedca[_bedee-1]--
		_bedee++
	}
}
func _aeee(_dgfa _cc.ChartRenderable) *Chart {
	return &Chart{_gff: _dgfa, _cggb: PositionRelative, _bgfg: Margins{Top: 10, Bottom: 10}}
}

// NewStyledTOCLine creates a new table of contents line with the provided style.
func (_gbfc *Creator) NewStyledTOCLine(number, title, page TextChunk, level uint, style TextStyle) *TOCLine {
	return _abgd(number, title, page, level, style)
}

// SetLanguageIdentifier sets the language identifier for the paragraph.
func (_baea *StyledParagraph) SetLanguageIdentifier(id string) { _baea._cccab = id }

// MultiCell makes a new cell with the specified row span and col span
// and inserts it into the table at the current position.
func (_abega *Table) MultiCell(rowspan, colspan int) *TableCell {
	_abega._ecgdc++
	_aeea := (_abega.moveToNextAvailableCell()-1)%(_abega._afacb) + 1
	_efgce := (_abega._ecgdc-1)/_abega._afacb + 1
	for _efgce > _abega._begg {
		_abega._begg++
		_abega._gecdg = append(_abega._gecdg, _abega._edefd)
	}
	_bedge := &TableCell{}
	_bedge._deef = _efgce
	_bedge._bacg = _aeea
	_bedge._fceba = 5
	_bedge._faab = CellBorderStyleNone
	_bedge._gaddf = _gga.LineStyleSolid
	_bedge._bbbff = CellHorizontalAlignmentLeft
	_bedge._bdfgg = CellVerticalAlignmentTop
	_bedge._caef = 0
	_bedge._eabab = 0
	_bedge._adec = 0
	_bedge._cddebf = 0
	_fgbef := ColorBlack
	_bedge._ebgcd = _fgbef
	_bedge._aegdf = _fgbef
	_bedge._fcbbf = _fgbef
	_bedge._gefce = _fgbef
	if rowspan < 1 {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061t\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0072\u006f\u0077s\u0070\u0061n\u0020\u0074o\u00201\u002e", rowspan)
		rowspan = 1
	}
	_abca := _abega._begg - (_bedge._deef - 1)
	if rowspan > _abca {
		_fec.Log.Debug("\u0054\u0061b\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070\u0061\u006e\u0020\u0028\u0025d\u0029\u0020\u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072o\u0077\u0073 \u0028\u0025\u0064\u0029.\u0020\u0041\u0064\u0064\u0069n\u0067\u0020\u0072\u006f\u0077\u0073\u002e", rowspan, _abca)
		_abega._begg += rowspan - 1
		for _eecag := 0; _eecag <= rowspan-_abca; _eecag++ {
			_abega._gecdg = append(_abega._gecdg, _abega._edefd)
		}
	}
	for _gbedg := 0; _gbedg < colspan && _aeea+_gbedg-1 < len(_abega._aedca); _gbedg++ {
		_abega._aedca[_aeea+_gbedg-1] = rowspan - 1
	}
	_bedge._afddb = rowspan
	if colspan < 1 {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0063\u006f\u006c\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061n\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006cs\u0070\u0061n\u0020\u0074o\u00201\u002e", colspan)
		colspan = 1
	}
	_ffdf := _abega._afacb - (_bedge._bacg - 1)
	if colspan > _ffdf {
		_fec.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0065\u006c\u006c\u0020\u0063o\u006c\u0073\u0070\u0061\u006e\u0020\u0028\u0025\u0064\u0029\u0020\u0065\u0078\u0063\u0065\u0065\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072\u006f\u0077\u0020\u0063\u006f\u006c\u0073\u0020\u0028\u0025d\u0029\u002e\u0020\u0041\u0064\u006a\u0075\u0073\u0074\u0069\u006e\u0067 \u0063\u006f\u006c\u0073\u0070\u0061n\u002e", colspan, _ffdf)
		colspan = _ffdf
	}
	_bedge._ecddcb = colspan
	_abega._ecgdc += colspan - 1
	_abega._efbe = append(_abega._efbe, _bedge)
	_bedge._abdcb = _abega
	return _bedge
}

// FitMode returns the fit mode of the image.
func (_aeeg *Image) FitMode() FitMode { return _aeeg._cba }

// Wrap wraps the text of the chunk into lines based on its style and the
// specified width.
func (_fgbcd *TextChunk) Wrap(width float64) ([]string, error) {
	if int(width) <= 0 {
		return []string{_fgbcd.Text}, nil
	}
	var _fdfg []string
	var _dgdeeb []rune
	var _addfc float64
	var _ccbea []float64
	_dbba := _fgbcd.Style
	_cecdb := _eedbc(_fgbcd.Text)
	for _, _fcgdf := range _fgbcd.Text {
		if _fcgdf == '\u000A' {
			_accb := _cgeac(string(_dgdeeb), _cecdb)
			_fdfg = append(_fdfg, _eg.TrimRightFunc(_accb, _fe.IsSpace)+string(_fcgdf))
			_dgdeeb = nil
			_addfc = 0
			_ccbea = nil
			continue
		}
		_affd := _fcgdf == ' '
		_gbceeb, _egcba := _dbba.Font.GetRuneMetrics(_fcgdf)
		if !_egcba {
			_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _fcgdf, _fcgdf, _dbba.Font.BaseFont(), _dbba.Font.Subtype())
			_fec.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _dbba.Font)
			_fec.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _dbba.Font.Encoder())
			return nil, _bd.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
		_edad := _dbba.FontSize * _gbceeb.Wx
		_bdfb := _edad
		if !_affd {
			_bdfb = _edad + _dbba.CharSpacing*1000.0
		}
		if _addfc+_edad > width*1000.0 {
			_adaaa := -1
			if !_affd {
				for _cabag := len(_dgdeeb) - 1; _cabag >= 0; _cabag-- {
					if _dgdeeb[_cabag] == ' ' {
						_adaaa = _cabag
						break
					}
				}
			}
			_agaef := string(_dgdeeb)
			if _adaaa > 0 {
				_agaef = string(_dgdeeb[0 : _adaaa+1])
				_dgdeeb = append(_dgdeeb[_adaaa+1:], _fcgdf)
				_ccbea = append(_ccbea[_adaaa+1:], _bdfb)
				_addfc = 0
				for _, _gegfa := range _ccbea {
					_addfc += _gegfa
				}
			} else {
				if _affd {
					_dgdeeb = []rune{}
					_ccbea = []float64{}
					_addfc = 0
				} else {
					_dgdeeb = []rune{_fcgdf}
					_ccbea = []float64{_bdfb}
					_addfc = _bdfb
				}
			}
			_agaef = _cgeac(_agaef, _cecdb)
			_fdfg = append(_fdfg, _eg.TrimRightFunc(_agaef, _fe.IsSpace))
		} else {
			_dgdeeb = append(_dgdeeb, _fcgdf)
			_addfc += _bdfb
			_ccbea = append(_ccbea, _bdfb)
		}
	}
	if len(_dgdeeb) > 0 {
		_faeg := string(_dgdeeb)
		_faeg = _cgeac(_faeg, _cecdb)
		_fdfg = append(_fdfg, _faeg)
	}
	return _fdfg, nil
}
func (_adfd *Invoice) drawAddress(_cgfe *InvoiceAddress) []*StyledParagraph {
	var _bgdg []*StyledParagraph
	if _cgfe.Heading != "" {
		_dadafd := _ddge(_adfd._geef)
		_dadafd.SetMargins(0, 0, 0, 7)
		_dadafd.Append(_cgfe.Heading)
		_bgdg = append(_bgdg, _dadafd)
	}
	_abfb := _ddge(_adfd._efdd)
	_abfb.SetLineHeight(1.2)
	_cbag := _cgfe.Separator
	if _cbag == "" {
		_cbag = _adfd._eabc
	}
	_bdcf := _cgfe.City
	if _cgfe.State != "" {
		if _bdcf != "" {
			_bdcf += _cbag
		}
		_bdcf += _cgfe.State
	}
	if _cgfe.Zip != "" {
		if _bdcf != "" {
			_bdcf += _cbag
		}
		_bdcf += _cgfe.Zip
	}
	if _cgfe.Name != "" {
		_abfb.Append(_cgfe.Name + "\u000a")
	}
	if _cgfe.Street != "" {
		_abfb.Append(_cgfe.Street + "\u000a")
	}
	if _cgfe.Street2 != "" {
		_abfb.Append(_cgfe.Street2 + "\u000a")
	}
	if _bdcf != "" {
		_abfb.Append(_bdcf + "\u000a")
	}
	if _cgfe.Country != "" {
		_abfb.Append(_cgfe.Country + "\u000a")
	}
	_fafe := _ddge(_adfd._efdd)
	_fafe.SetLineHeight(1.2)
	_fafe.SetMargins(0, 0, 7, 0)
	if _cgfe.Phone != "" {
		_fafe.Append(_cgfe.fmtLine(_cgfe.Phone, "\u0050h\u006f\u006e\u0065\u003a\u0020", _cgfe.HidePhoneLabel))
	}
	if _cgfe.Email != "" {
		_fafe.Append(_cgfe.fmtLine(_cgfe.Email, "\u0045m\u0061\u0069\u006c\u003a\u0020", _cgfe.HideEmailLabel))
	}
	_bgdg = append(_bgdg, _abfb, _fafe)
	return _bgdg
}

// SetMarkedContentID sets the marked content ID for the image.
func (_acdb *Image) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_acdb._gebca = &mcid
	_fbfd := _fgd.NewKDictionary()
	_fbfd.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_fbfd.K = _bc.MakeInteger(mcid)
	return _fbfd
}

// SetFitMode sets the fit mode of the rectangle.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_afdff *Rectangle) SetFitMode(fitMode FitMode) { _afdff._aaab = fitMode }
func (_cbfge *Invoice) drawInformation() *Table {
	_dcaf := _bagfg(2)
	_dcgaa := append([][2]*InvoiceCell{_cbfge._bcab, _cbfge._cbfg, _cbfge._dfge}, _cbfge._afff...)
	for _, _fdebg := range _dcgaa {
		_cdacg, _dgga := _fdebg[0], _fdebg[1]
		if _dgga.Value == "" {
			continue
		}
		_fffef := _dcaf.NewCell()
		_fffef.SetBackgroundColor(_cdacg.BackgroundColor)
		_cbfge.setCellBorder(_fffef, _cdacg)
		_aegb := _ddge(_cdacg.TextStyle)
		_aegb.Append(_cdacg.Value)
		_aegb.SetMargins(0, 0, 2, 1)
		_fffef.SetContent(_aegb)
		_fffef = _dcaf.NewCell()
		_fffef.SetBackgroundColor(_dgga.BackgroundColor)
		_cbfge.setCellBorder(_fffef, _dgga)
		_aegb = _ddge(_dgga.TextStyle)
		_aegb.Append(_dgga.Value)
		_aegb.SetMargins(0, 0, 2, 1)
		_fffef.SetContent(_aegb)
	}
	return _dcaf
}
func _ebea(_fgfdg *_fgd.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 238), Font: _fgfdg, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}
func (_eggc *templateProcessor) parseTextChunk(_bfcf *templateNode, _ccgce *TextChunk) (interface{}, error) {
	if _bfcf._cefd == nil {
		_eggc.nodeLogError(_bfcf, "\u0054\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gdbeeg
	}
	var (
		_bgadc = _eggc.creator.NewTextStyle()
		_baeg  bool
	)
	for _, _baccf := range _bfcf._adfdg.Attr {
		if _baccf.Name.Local == "\u006c\u0069\u006e\u006b" {
			_ceegc, _caedd := _bfcf._cefd._bedcd.(*StyledParagraph)
			if !_caedd {
				_eggc.nodeLogError(_bfcf, "\u004c\u0069\u006e\u006b \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u006f\u006e\u006c\u0079\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0070\u0061\u0072\u0061\u0067r\u0061\u0070\u0068\u0027\u0073\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e")
				_baeg = true
			} else {
				_bgadc = _ceegc._fbegf
			}
			break
		}
	}
	if _ccgce == nil {
		_ccgce = NewTextChunk("", _bgadc)
	}
	for _, _eggbdb := range _bfcf._adfdg.Attr {
		_gdcag := _eggbdb.Value
		switch _cfbfg := _eggbdb.Name.Local; _cfbfg {
		case "\u0063\u006f\u006co\u0072":
			_ccgce.Style.Color = _eggc.parseColorAttr(_cfbfg, _gdcag)
		case "\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_ccgce.Style.OutlineColor = _eggc.parseColorAttr(_cfbfg, _gdcag)
		case "\u0066\u006f\u006e\u0074":
			_ccgce.Style.Font = _eggc.parseFontAttr(_cfbfg, _gdcag)
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_ccgce.Style.FontSize = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u006f\u0075\u0074l\u0069\u006e\u0065\u002d\u0073\u0069\u007a\u0065":
			_ccgce.Style.OutlineSize = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u002d\u0073\u0070a\u0063\u0069\u006e\u0067":
			_ccgce.Style.CharSpacing = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u0068o\u0072i\u007a\u006f\u006e\u0074\u0061l\u002d\u0073c\u0061\u006c\u0069\u006e\u0067":
			_ccgce.Style.HorizontalScaling = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067-\u006d\u006f\u0064\u0065":
			_ccgce.Style.RenderingMode = _eggc.parseTextRenderingModeAttr(_cfbfg, _gdcag)
		case "\u0075n\u0064\u0065\u0072\u006c\u0069\u006ee":
			_ccgce.Style.Underline = _eggc.parseBoolAttr(_cfbfg, _gdcag)
		case "\u0075n\u0064e\u0072\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_ccgce.Style.UnderlineStyle.Color = _eggc.parseColorAttr(_cfbfg, _gdcag)
		case "\u0075\u006ed\u0065\u0072\u006ci\u006e\u0065\u002d\u006f\u0066\u0066\u0073\u0065\u0074":
			_ccgce.Style.UnderlineStyle.Offset = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u0075\u006e\u0064\u0065rl\u0069\u006e\u0065\u002d\u0074\u0068\u0069\u0063\u006b\u006e\u0065\u0073\u0073":
			_ccgce.Style.UnderlineStyle.Thickness = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		case "\u006c\u0069\u006e\u006b":
			if !_baeg {
				_ccgce._abefd = _eggc.parseLinkAttr(_cfbfg, _gdcag)
			}
		case "\u0074e\u0078\u0074\u002d\u0072\u0069\u0073e":
			_ccgce.Style.TextRise = _eggc.parseFloatAttr(_cfbfg, _gdcag)
		default:
			_eggc.nodeLogDebug(_bfcf, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _cfbfg)
		}
	}
	return _ccgce, nil
}

// CellVerticalAlignment defines the table cell's vertical alignment.
type CellVerticalAlignment int

// SetColor sets the line color. Use ColorRGBFromHex, ColorRGBFrom8bit or
// ColorRGBFromArithmetic to create the color object.
func (_cedag *Line) SetColor(color Color) { _cedag._ebgca = color }
func (_aceeg *templateProcessor) parseTextOverflowAttr(_cfgfg, _bafa string) TextOverflow {
	_fec.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0074e\u0078\u0074\u0020o\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _cfgfg, _bafa)
	_eaaea := map[string]TextOverflow{"\u0076i\u0073\u0069\u0062\u006c\u0065": TextOverflowVisible, "\u0068\u0069\u0064\u0064\u0065\u006e": TextOverflowHidden}[_bafa]
	return _eaaea
}

// SetContent sets the cell's content.  The content is a VectorDrawable, i.e.
// a Drawable with a known height and width.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Table
// - *Division
// - *List
// - *Rectangle
// - *Ellipse
// - *Line
func (_dcbea *TableCell) SetContent(vd VectorDrawable) error {
	switch _cfcef := vd.(type) {
	case *Paragraph:
		if _cfcef._dcefd {
			_cfcef._ecae = true
		}
		_dcbea._aafg = vd
	case *StyledParagraph:
		if _cfcef._fegde {
			_cfcef._fegca = true
		}
		_dcbea._aafg = vd
	case *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line:
		_dcbea._aafg = vd
	default:
		_fec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0054", vd)
		return _bc.ErrTypeError
	}
	return nil
}
func _gbggc(_efed string) (*Image, error) {
	_badc, _cdebe := _b.Open(_efed)
	if _cdebe != nil {
		return nil, _cdebe
	}
	defer _badc.Close()
	_dgef, _cdebe := _fgd.ImageHandling.Read(_badc)
	if _cdebe != nil {
		_fec.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _cdebe)
		return nil, _cdebe
	}
	return _eddeb(_dgef)
}
func _aggd(_adgg, _dcbb, _faga, _ebcbd float64) *Rectangle {
	return &Rectangle{_defb: _adgg, _cdgf: _dcbb, _bcede: _faga, _cggg: _ebcbd, _dfff: PositionAbsolute, _afdae: 1.0, _bedf: ColorBlack, _caeg: 1.0, _dbff: 1.0}
}

// GetOptimizer returns current PDF optimizer.
func (_abebc *Creator) GetOptimizer() _fgd.Optimizer { return _abebc._caec }

// SetStyle sets the style of the line (solid or dashed).
func (_bbedd *Line) SetStyle(style _gga.LineStyle) { _bbedd._cedaf = style }

// Background contains properties related to the background of a component.
type Background struct {
	FillColor               Color
	BorderColor             Color
	BorderSize              float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64
}

// GeneratePageBlocks draws the line on a new block representing the page.
// Implements the Drawable interface.
func (_fdbf *Line) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_affb          []*Block
		_cdea          = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_ebaa          = ctx
		_adbg, _bafde  = _fdbf._gbcgde, ctx.PageHeight - _fdbf._gcfa
		_bggef, _eggbd = _fdbf._cfge, ctx.PageHeight - _fdbf._ceef
	)
	_fcec := _fdbf._babe.IsRelative()
	if _fcec {
		ctx.X += _fdbf._fggd.Left
		ctx.Y += _fdbf._fggd.Top
		ctx.Width -= _fdbf._fggd.Left + _fdbf._fggd.Right
		ctx.Height -= _fdbf._fggd.Top + _fdbf._fggd.Bottom
		_adbg, _bafde, _bggef, _eggbd = _fdbf.computeCoords(ctx)
		if _fdbf.Height() > ctx.Height {
			_affb = append(_affb, _cdea)
			_cdea = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_afac := ctx
			_afac.Y = ctx.Margins.Top + _fdbf._fggd.Top
			_afac.X = ctx.Margins.Left + _fdbf._fggd.Left
			_afac.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _fdbf._fggd.Top - _fdbf._fggd.Bottom
			_afac.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _fdbf._fggd.Left - _fdbf._fggd.Right
			ctx = _afac
			_adbg, _bafde, _bggef, _eggbd = _fdbf.computeCoords(ctx)
		}
	}
	_ecag := _gga.BasicLine{X1: _adbg, Y1: _bafde, X2: _bggef, Y2: _eggbd, LineColor: _cfcee(_fdbf._ebgca), Opacity: _fdbf._eabf, LineWidth: _fdbf._aebf, LineStyle: _fdbf._cedaf, DashArray: _fdbf._cgecc, DashPhase: _fdbf._dcebc}
	_efbf, _ebdf := _cdea.setOpacity(1.0, _fdbf._eabf)
	if _ebdf != nil {
		return nil, ctx, _ebdf
	}
	_ccdfa, _, _ebdf := _ecag.MarkedDraw(_efbf, _fdbf._fbeb)
	if _ebdf != nil {
		return nil, ctx, _ebdf
	}
	if _ebdf = _cdea.addContentsByString(string(_ccdfa)); _ebdf != nil {
		return nil, ctx, _ebdf
	}
	if _fcec {
		ctx.X = _ebaa.X
		ctx.Width = _ebaa.Width
		_gccd := _fdbf.Height()
		ctx.Y += _gccd + _fdbf._fggd.Bottom
		ctx.Height -= _gccd
	} else {
		ctx = _ebaa
	}
	_affb = append(_affb, _cdea)
	return _affb, ctx, nil
}

// SetStyleTop sets border style for top side.
func (_ece *border) SetStyleTop(style CellBorderStyle) { _ece._cggf = style }
func _eedbc(_caaf string) bool {
	_cfga := func(_aadbg rune) bool { return _aadbg == '\u000A' }
	_bbaeg := _eg.TrimFunc(_caaf, _cfga)
	_gfcab := _fa.Paragraph{}
	_, _dfgf := _gfcab.SetString(_bbaeg)
	if _dfgf != nil {
		return true
	}
	_agcgf, _dfgf := _gfcab.Order()
	if _dfgf != nil {
		return true
	}
	if _agcgf.NumRuns() < 1 {
		return true
	}
	return _gfcab.IsLeftToRight()
}

// SetMakeredContentID sets the marked content identifier for the ellipse.
func (_adcd *Ellipse) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_adcd._adde = &mcid
	_cbea := _fgd.NewKDictionary()
	_cbea.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_cbea.K = _bc.MakeInteger(mcid)
	return _cbea
}

// SetFont sets the Paragraph's font.
func (_fdff *Paragraph) SetFont(font *_fgd.PdfFont) { _fdff._gbca = font }
func _gbcbg(_dggcg *templateProcessor, _adgff *templateNode) (interface{}, error) {
	return _dggcg.parseEllipse(_adgff)
}

// SetBorderColor sets the border color.
func (_afce *CurvePolygon) SetBorderColor(color Color) { _afce._dgc.BorderColor = _cfcee(color) }

// SetEncoder sets the encoding/compression mechanism for the image.
func (_edba *Image) SetEncoder(encoder _bc.StreamEncoder) { _edba._fdag = encoder }
func _becb(_aebg, _cef, _ebef, _bde float64) *Ellipse {
	return &Ellipse{_adbc: _aebg, _ddaee: _cef, _eefg: _ebef, _beb: _bde, _fdbc: PositionAbsolute, _fcda: 1.0, _add: ColorBlack, _gfbg: 1.0, _fbcf: 1.0}
}

// GeneratePageBlocks draws the polyline on a new block representing the page.
// Implements the Drawable interface.
func (_eece *Polyline) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fcaa := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_cfgf, _gecbc := _fcaa.setOpacity(_eece._ggaf, _eece._ggaf)
	if _gecbc != nil {
		return nil, ctx, _gecbc
	}
	_bcgb := _eece._cdggc.Points
	for _cgbd := range _bcgb {
		_bbdba := &_bcgb[_cgbd]
		_bbdba.Y = ctx.PageHeight - _bbdba.Y
	}
	_cddb, _, _gecbc := _eece._cdggc.MarkedDraw(_cfgf, _eece._egfd)
	if _gecbc != nil {
		return nil, ctx, _gecbc
	}
	if _gecbc = _fcaa.addContentsByString(string(_cddb)); _gecbc != nil {
		return nil, ctx, _gecbc
	}
	return []*Block{_fcaa}, ctx, nil
}
func (_cded *Division) ctxHeight(_efe float64) float64 {
	_efe -= _cded._ddfg.Left + _cded._ddfg.Right + _cded._gbga.Left + _cded._gbga.Right
	var _bbec float64
	for _, _bfgee := range _cded._gebd {
		_bbec += _bbbf(_bfgee, _efe)
	}
	return _bbec
}
func _dcbc() *Division { return &Division{_defa: true} }

// NewStyledParagraph creates a new styled paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_cdcf *Creator) NewStyledParagraph() *StyledParagraph { return _ddge(_cdcf.NewTextStyle()) }

// NewLine creates a new line between (x1, y1) to (x2, y2),
// using default attributes.
// NOTE: In relative positioning mode, `x1` and `y1` are calculated using the
// current context and `x2`, `y2` are used only to calculate the position of
// the second point in relation to the first one (used just as a measurement
// of size). Furthermore, when the fit mode is set to fill the context width,
// `x2` is set to the right edge coordinate of the context.
func (_aggf *Creator) NewLine(x1, y1, x2, y2 float64) *Line { return _ddfb(x1, y1, x2, y2) }

// NewTOCLine creates a new table of contents line with the default style.
func (_geaf *Creator) NewTOCLine(number, title, page string, level uint) *TOCLine {
	return _afafb(number, title, page, level, _geaf.NewTextStyle())
}
func (_gbecd *TextStyle) horizontalScale() float64 { return _gbecd.HorizontalScaling / 100 }
func _egce(_debb *templateProcessor, _gfaag *templateNode) (interface{}, error) {
	return _debb.parseListItem(_gfaag)
}
func _abgd(_ebecd, _ecdde, _gace TextChunk, _ebdab uint, _ddaba TextStyle) *TOCLine {
	_gcfad := _ddge(_ddaba)
	_gcfad.SetEnableWrap(true)
	_gcfad.SetTextAlignment(TextAlignmentLeft)
	_gcfad.SetMargins(0, 0, 2, 2)
	_debf := &TOCLine{_afgdd: _gcfad, Number: _ebecd, Title: _ecdde, Page: _gace, Separator: TextChunk{Text: "\u002e", Style: _ddaba}, _afage: 0, _gfec: _ebdab, _ebgcdc: 10, _afdaf: PositionRelative}
	_gcfad._ccddg.Left = _debf._afage + float64(_debf._gfec-1)*_debf._ebgcdc
	_gcfad._beccd = _debf.prepareParagraph
	return _debf
}

// SetMarkedContentID sets the marked content ID.
func (_defe *PageBreak) SetMarkedContentID(id int64) *_fgd.KDict { return nil }

// SetTerms sets the terms and conditions section of the invoice.
func (_aedc *Invoice) SetTerms(title, content string) { _aedc._affc = [2]string{title, content} }

// WriteToFile writes the Creator output to file specified by path.
func (_eae *Creator) WriteToFile(outputPath string) error {
	_gebc, _gdfb := _b.Create(outputPath)
	if _gdfb != nil {
		return _gdfb
	}
	defer _gebc.Close()
	return _eae.Write(_gebc)
}
func _cfcee(_acfd Color) _fgd.PdfColor {
	if _acfd == nil {
		_acfd = ColorBlack
	}
	switch _aadc := _acfd.(type) {
	case grayColor:
		return _fgd.NewPdfColorDeviceGray(_aadc._dbdf)
	case cmykColor:
		return _fgd.NewPdfColorDeviceCMYK(_aadc._ffea, _aadc._cbg, _aadc._egb, _aadc._gbfd)
	case *LinearShading:
		return _fgd.NewPdfColorPatternType2()
	case *RadialShading:
		return _fgd.NewPdfColorPatternType3()
	}
	return _fgd.NewPdfColorDeviceRGB(_acfd.ToRGB())
}

// Opacity returns the opacity of the line.
func (_eaac *Line) Opacity() float64 { return _eaac._eabf }
func _ccfcc(_fbdf []_gga.CubicBezierCurve) *PolyBezierCurve {
	return &PolyBezierCurve{_ffcd: &_gga.PolyBezierCurve{Curves: _fbdf, BorderColor: _fgd.NewPdfColorDeviceRGB(0, 0, 0), BorderWidth: 1.0}, _ffafb: 1.0, _addf: 1.0}
}

// Append adds a new text chunk to the paragraph.
func (_daad *StyledParagraph) Append(text string) *TextChunk {
	_abbda := NewTextChunk(text, _daad._fgbg)
	return _daad.appendChunk(_abbda)
}
func _cdbe(_eage string, _eeaf TextStyle) *Paragraph {
	_abdb := &Paragraph{_egea: _eage, _gbca: _eeaf.Font, _acdge: _eeaf.FontSize, _fbcg: 1.0, _ecae: true, _dcefd: true, _cebg: TextAlignmentLeft, _cadb: 0, _gccbb: 1, _eacb: 1, _dbgab: PositionRelative, _eadae: ""}
	_abdb.SetColor(_eeaf.Color)
	return _abdb
}

// AddExternalLink adds a new external link to the paragraph.
// The text parameter represents the text that is displayed and the url
// parameter sets the destionation of the link.
func (_ccadd *StyledParagraph) AddExternalLink(text, url string) *TextChunk {
	_gcfe := NewTextChunk(text, _ccadd._fbegf)
	_gcfe._abefd = _fecd(url)
	return _ccadd.appendChunk(_gcfe)
}

// SetColorLeft sets border color for left.
func (_fda *border) SetColorLeft(col Color) { _fda._bba = col }

// GetMargins returns the Chapter's margin: left, right, top, bottom.
func (_caba *Chapter) GetMargins() (float64, float64, float64, float64) {
	return _caba._dbe.Left, _caba._dbe.Right, _caba._dbe.Top, _caba._dbe.Bottom
}

// BorderColor returns the border color of the ellipse.
func (_gfea *Ellipse) BorderColor() Color { return _gfea._add }

// EnablePageWrap controls whether the division is wrapped across pages.
// If disabled, the division is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the division is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
// Currently, page wrapping can only be disabled for vertical divisions.
func (_gbbc *Division) EnablePageWrap(enable bool) { _gbbc._defa = enable }

// Sections returns the custom content sections of the invoice as
// title-content pairs.
func (_fafd *Invoice) Sections() [][2]string { return _fafd._fcfd }

// AddColorStop add color stop information for rendering gradient.
func (_cfbgc *shading) AddColorStop(color Color, point float64) {
	_cfbgc._bcaca = append(_cfbgc._bcaca, _bdab(color, point))
}

// Height returns the height of the line.
func (_dgbc *Line) Height() float64 {
	_edce := _dgbc._aebf
	if _dgbc._gbcgde == _dgbc._cfge {
		_edce /= 2
	}
	return _fc.Abs(_dgbc._ceef-_dgbc._gcfa) + _edce
}

// NewSubchapter creates a new child chapter with the specified title.
func (_bfbc *Chapter) NewSubchapter(title string) *Chapter {
	_eeg := _ddaea(_bfbc._gfdf._gbca)
	_eeg.FontSize = 14
	_bfbc._bdcd++
	_gfb := _dgg(_bfbc, _bfbc._edc, _bfbc._agbcd, title, _bfbc._bdcd, _eeg)
	_bfbc.Add(_gfb)
	return _gfb
}
func (_baafa *TemplateOptions) init() {
	if _baafa.SubtemplateMap == nil {
		_baafa.SubtemplateMap = map[string]_gg.Reader{}
	}
	if _baafa.FontMap == nil {
		_baafa.FontMap = map[string]*_fgd.PdfFont{}
	}
	if _baafa.ImageMap == nil {
		_baafa.ImageMap = map[string]*_fgd.Image{}
	}
	if _baafa.ColorMap == nil {
		_baafa.ColorMap = map[string]Color{}
	}
	if _baafa.ChartMap == nil {
		_baafa.ChartMap = map[string]_cc.ChartRenderable{}
	}
}

// Flip flips the active page on the specified axes.
// If `flipH` is true, the page is flipped horizontally. Similarly, if `flipV`
// is true, the page is flipped vertically. If both are true, the page is
// flipped both horizontally and vertically.
// NOTE: the flip transformations are applied when the creator is finalized,
// which is at write time in most cases.
func (_ffbbd *Creator) Flip(flipH, flipV bool) error {
	_bgdd := _ffbbd.getActivePage()
	if _bgdd == nil {
		return _bd.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_cgea, _dbgb := _ffbbd._agfdc[_bgdd]
	if !_dbgb {
		_cgea = &pageTransformations{}
		_ffbbd._agfdc[_bgdd] = _cgea
	}
	_cgea._bbed = flipH
	_cgea._fcge = flipV
	return nil
}
func (_gbcga *TextChunk) clone() *TextChunk {
	_gfga := *_gbcga
	_gfga._abefd = _bggbd(_gbcga._abefd)
	return &_gfga
}
func (_gbc *Block) duplicate() *Block {
	_bcc := &Block{}
	*_bcc = *_gbc
	_acd := _dg.ContentStreamOperations{}
	_acd = append(_acd, *_gbc._cb...)
	_bcc._cb = &_acd
	return _bcc
}
func (_cgbdg *templateProcessor) parseColor(_gbfee string) Color {
	if _gbfee == "" {
		return nil
	}
	_gfdd, _cdcab := _cgbdg._cbcec.ColorMap[_gbfee]
	if _cdcab {
		return _gfdd
	}
	if _gbfee[0] == '#' {
		return ColorRGBFromHex(_gbfee)
	}
	return nil
}

// Vertical returns total vertical (top + bottom) margin.
func (_aecf *Margins) Vertical() float64 { return _aecf.Bottom + _aecf.Top }

// Link returns link information for this line.
func (_daecc *TOCLine) Link() (_defac int64, _fded, _ccffc float64) {
	return _daecc._ebaab, _daecc._beed, _daecc._cdbff
}

// SetFontSize sets the font size in document units (points).
func (_aead *Paragraph) SetFontSize(fontSize float64) { _aead._acdge = fontSize }

// New creates a new instance of the PDF Creator.
func New() *Creator {
	const _daf = "c\u0072\u0065\u0061\u0074\u006f\u0072\u002e\u004e\u0065\u0077"
	_cdeb := &Creator{}
	_cdeb._egfb = []*_fgd.PdfPage{}
	_cdeb._feab = map[*_fgd.PdfPage]*Block{}
	_cdeb._agfdc = map[*_fgd.PdfPage]*pageTransformations{}
	_cdeb.SetPageSize(PageSizeLetter)
	_bcgg := 0.1 * _cdeb._gbgg
	_cdeb._cbb.Left = _bcgg
	_cdeb._cbb.Right = _bcgg
	_cdeb._cbb.Top = _bcgg
	_cdeb._cbb.Bottom = _bcgg
	var _egca error
	_cdeb._dff, _egca = _fgd.NewStandard14Font(_fgd.HelveticaName)
	if _egca != nil {
		_cdeb._dff = _fgd.DefaultFont()
	}
	_cdeb._ceb, _egca = _fgd.NewStandard14Font(_fgd.HelveticaBoldName)
	if _egca != nil {
		_cdeb._dff = _fgd.DefaultFont()
	}
	_cdeb._aac = _cdeb.NewTOC("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073")
	_cdeb.AddOutlines = true
	_cdeb._cbba = _fgd.NewOutline()
	_aa.TrackUse(_daf)
	return _cdeb
}

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_bdfd *RadialShading) AddShadingResource(block *Block) (_adgc _bc.PdfObjectName, _dafec error) {
	_afgb := 1
	_adgc = _bc.PdfObjectName("\u0053\u0068" + _fg.Itoa(_afgb))
	for block._dgd.HasShadingByName(_adgc) {
		_afgb++
		_adgc = _bc.PdfObjectName("\u0053\u0068" + _fg.Itoa(_afgb))
	}
	if _facee := block._dgd.SetShadingByName(_adgc, _bdfd.shadingModel().ToPdfObject()); _facee != nil {
		return "", _facee
	}
	return _adgc, nil
}

// EnableRowWrap controls whether rows are wrapped across pages.
// NOTE: Currently, row wrapping is supported for rows using StyledParagraphs.
func (_bdfdc *Table) EnableRowWrap(enable bool) { _bdfdc._fgaag = enable }

// ScaleToHeight sets the graphic svg scaling factor with the given height.
func (_aaff *GraphicSVG) ScaleToHeight(h float64) {
	_bade := _aaff._dgccb.Width / _aaff._dgccb.Height
	_aaff._dgccb.Height = h
	_aaff._dgccb.Width = h * _bade
	_aaff._dgccb.SetScaling(_bade, _bade)
}

// SetFitMode sets the fit mode of the line.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_cdgga *Line) SetFitMode(fitMode FitMode) { _cdgga._efdg = fitMode }
func (_bdgc *Division) split(_febg DrawContext) (_gaeb, _bbcgg *Division) {
	var (
		_efcc        float64
		_deab, _fefg []VectorDrawable
	)
	_ebdb := _febg.Width - _bdgc._ddfg.Left - _bdgc._ddfg.Right - _bdgc._gbga.Left - _bdgc._gbga.Right
	for _ecb, _cfgb := range _bdgc._gebd {
		_efcc += _bbbf(_cfgb, _ebdb)
		if _efcc < _febg.Height {
			_deab = append(_deab, _cfgb)
		} else {
			_fefg = _bdgc._gebd[_ecb:]
			break
		}
	}
	if len(_deab) > 0 {
		_gaeb = _dcbc()
		*_gaeb = *_bdgc
		_gaeb._gebd = _deab
		if _bdgc._agecb != nil {
			_gaeb._agecb = &Background{}
			*_gaeb._agecb = *_bdgc._agecb
		}
	}
	if len(_fefg) > 0 {
		_bbcgg = _dcbc()
		*_bbcgg = *_bdgc
		_bbcgg._gebd = _fefg
		if _bdgc._agecb != nil {
			_bbcgg._agecb = &Background{}
			*_bbcgg._agecb = *_bdgc._agecb
		}
	}
	return _gaeb, _bbcgg
}

// FitMode returns the fit mode of the rectangle.
func (_edgb *Rectangle) FitMode() FitMode { return _edgb._aaab }
func (_aeegcb *templateProcessor) parseTextAlignmentAttr(_fcbg, _gdcgbc string) TextAlignment {
	_fec.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020t\u0065\u0078\u0074\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fcbg, _gdcgbc)
	_bffabd := map[string]TextAlignment{"\u006c\u0065\u0066\u0074": TextAlignmentLeft, "\u0072\u0069\u0067h\u0074": TextAlignmentRight, "\u0063\u0065\u006e\u0074\u0065\u0072": TextAlignmentCenter, "\u006au\u0073\u0074\u0069\u0066\u0079": TextAlignmentJustify}[_gdcgbc]
	return _bffabd
}

// LevelOffset returns the amount of space an indentation level occupies.
func (_dddea *TOCLine) LevelOffset() float64 { return _dddea._ebgcdc }
func _bded(_eba string) (*GraphicSVG, error) {
	_abgg, _ggfd := _dd.ParseFromFile(_eba)
	if _ggfd != nil {
		return nil, _ggfd
	}
	return _cdfg(_abgg)
}

// ScaleToWidth scales the rectangle to the specified width. The height of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_eeca *Rectangle) ScaleToWidth(w float64) {
	_bbccc := _eeca._cggg / _eeca._bcede
	_eeca._bcede = w
	_eeca._cggg = w * _bbccc
}

// SetStyleLeft sets border style for left side.
func (_decb *border) SetStyleLeft(style CellBorderStyle) { _decb._ffbb = style }

// FitMode returns the fit mode of the line.
func (_bggb *Line) FitMode() FitMode { return _bggb._efdg }

var PPI float64 = 72

// GetMargins returns the Image's margins: left, right, top, bottom.
func (_bgce *Image) GetMargins() (float64, float64, float64, float64) {
	return _bgce._cdbc.Left, _bgce._cdbc.Right, _bgce._cdbc.Top, _bgce._cdbc.Bottom
}

// Scale scales the rectangle dimensions by the specified factors.
func (_bddf *Rectangle) Scale(xFactor, yFactor float64) {
	_bddf._bcede = xFactor * _bddf._bcede
	_bddf._cggg = yFactor * _bddf._cggg
}

// SetWidth sets the Paragraph width. This is essentially the wrapping width, i.e. the width the
// text can extend to prior to wrapping over to next line.
func (_gcgdd *Paragraph) SetWidth(width float64) { _gcgdd._gcbf = width; _gcgdd.wrapText() }

// AppendColumn appends a column to the line items table.
func (_cfegc *Invoice) AppendColumn(description string) *InvoiceCell {
	_gbcff := _cfegc.NewColumn(description)
	_cfegc._gffd = append(_cfegc._gffd, _gbcff)
	return _gbcff
}

type templateTag struct {
	_geca map[string]struct{}
	_cacf func(*templateProcessor, *templateNode) (interface{}, error)
}

func (_aeeed *Invoice) generateHeaderBlocks(_aebc DrawContext) ([]*Block, DrawContext, error) {
	_cafg := _ddge(_aeeed._adaa)
	_cafg.SetEnableWrap(true)
	_cafg.Append(_aeeed._dgfaa)
	_efcd := _bagfg(2)
	if _aeeed._gdfdf != nil {
		_bdbe := _efcd.NewCell()
		_bdbe.SetHorizontalAlignment(CellHorizontalAlignmentLeft)
		_bdbe.SetVerticalAlignment(CellVerticalAlignmentMiddle)
		_bdbe.SetIndent(0)
		_bdbe.SetContent(_aeeed._gdfdf)
		_aeeed._gdfdf.ScaleToHeight(_cafg.Height() + 20)
	} else {
		_efcd.SkipCells(1)
	}
	_fcce := _efcd.NewCell()
	_fcce.SetHorizontalAlignment(CellHorizontalAlignmentRight)
	_fcce.SetVerticalAlignment(CellVerticalAlignmentMiddle)
	_fcce.SetContent(_cafg)
	return _efcd.GeneratePageBlocks(_aebc)
}

// SetLineStyle sets the style for all the line components: number, title,
// separator, page. The style is applied only for new lines added to the
// TOC component.
func (_dgcgd *TOC) SetLineStyle(style TextStyle) {
	_dgcgd.SetLineNumberStyle(style)
	_dgcgd.SetLineTitleStyle(style)
	_dgcgd.SetLineSeparatorStyle(style)
	_dgcgd.SetLinePageStyle(style)
}

// SetWidthBottom sets border width for bottom.
func (_eabb *border) SetWidthBottom(bw float64) { _eabb._fcf = bw }

// SetBuyerAddress sets the buyer address of the invoice.
func (_dccd *Invoice) SetBuyerAddress(address *InvoiceAddress) { _dccd._eeeb = address }

// AddLine appends a new line to the invoice line items table.
func (_fdfdb *Invoice) AddLine(values ...string) []*InvoiceCell {
	_gcef := len(_fdfdb._gffd)
	var _dgbd []*InvoiceCell
	for _beae, _bebg := range values {
		_dfebd := _fdfdb.newCell(_bebg, _fdfdb._baae)
		if _beae < _gcef {
			_dfebd.Alignment = _fdfdb._gffd[_beae].Alignment
		}
		_dgbd = append(_dgbd, _dfebd)
	}
	_fdfdb._agece = append(_fdfdb._agece, _dgbd)
	return _dgbd
}

// SetBorderOpacity sets the border opacity.
func (_ebaed *PolyBezierCurve) SetBorderOpacity(opacity float64) { _ebaed._addf = opacity }

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_dafc *Creator) SetOptimizer(optimizer _fgd.Optimizer) { _dafc._caec = optimizer }

// String implements error interface.
func (_ecebb UnsupportedRuneError) Error() string { return _ecebb.Message }

// SetMargins sets the margins of the graphic svg component.
func (_dfce *GraphicSVG) SetMargins(left, right, top, bottom float64) {
	_dfce._aagce.Left = left
	_dfce._aagce.Right = right
	_dfce._aagce.Top = top
	_dfce._aagce.Bottom = bottom
}

// Heading returns the heading component of the table of contents.
func (_fddddf *TOC) Heading() *StyledParagraph { return _fddddf._defae }

// SetStyleRight sets border style for right side.
func (_bdb *border) SetStyleRight(style CellBorderStyle) { _bdb._cfc = style }
func (_abdad *templateProcessor) parseTextVerticalAlignmentAttr(_cgcbf, _cdae string) TextVerticalAlignment {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _cgcbf, _cdae)
	_gdeb := map[string]TextVerticalAlignment{"\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065": TextVerticalAlignmentBaseline, "\u0063\u0065\u006e\u0074\u0065\u0072": TextVerticalAlignmentCenter}[_cdae]
	return _gdeb
}

const (
	CellBorderSideLeft CellBorderSide = iota
	CellBorderSideRight
	CellBorderSideTop
	CellBorderSideBottom
	CellBorderSideAll
)

// NewFilledCurve returns a instance of filled curve.
func (_fcdf *Creator) NewFilledCurve() *FilledCurve { return _aagf() }
func (_cgae *templateProcessor) loadImageFromSrc(_aecag string) (*Image, error) {
	if _aecag == "" {
		_fec.Log.Error("\u0049\u006d\u0061\u0067\u0065\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _gedbe
	}
	_beffe := _eg.Split(_aecag, "\u002c")
	for _, _fbggg := range _beffe {
		_fbggg = _eg.TrimSpace(_fbggg)
		if _fbggg == "" {
			continue
		}
		_bgacf, _bafce := _cgae._cbcec.ImageMap[_fbggg]
		if _bafce {
			return _eddeb(_bgacf)
		}
		if _eedaa := _cgae.parseAttrPropList(_fbggg); len(_eedaa) > 0 {
			if _dffec, _bafg := _eedaa["\u0070\u0061\u0074\u0068"]; _bafg {
				if _cgdffd, _fcdeb := _gbggc(_dffec); _fcdeb != nil {
					_fec.Log.Debug("\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0060\u0025\u0073\u0060\u003a\u0020\u0025\u0076\u002e", _dffec, _fcdeb)
				} else {
					return _cgdffd, nil
				}
			}
		}
	}
	_fec.Log.Error("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0069ma\u0067\u0065\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _aecag)
	return nil, _gedbe
}

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_ebda *LinearShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }

// SetHeight sets the height of the ellipse.
func (_aecfd *Ellipse) SetHeight(height float64) { _aecfd._beb = height }

// SetExtends specifies whether to extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_gbedf *shading) SetExtends(start bool, end bool) { _gbedf._fece = []bool{start, end} }

// SetStructTreeRoot sets the structure tree root to be appended in the document that will be created.
func (_fbaf *Creator) SetStructTreeRoot(structTreeRoot *_fgd.StructTreeRoot) {
	_fbaf._cfgc = structTreeRoot
}
func (_agg *Block) mergeBlocks(_dgdd *Block) error {
	_bced := _gbg(_agg._cb, _agg._dgd, _dgdd._cb, _dgdd._dgd)
	if _bced != nil {
		return _bced
	}
	for _, _aba := range _dgdd._ed {
		_agg.AddAnnotation(_aba)
	}
	return nil
}
func (_bbcab *templateProcessor) parseFontAttr(_bgddc, _dedd string) *_fgd.PdfFont {
	_fec.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _bgddc, _dedd)
	_ecdg := _bbcab.creator._dff
	if _dedd == "" {
		return _ecdg
	}
	_gccfb := _eg.Split(_dedd, "\u002c")
	for _, _eaeb := range _gccfb {
		_eaeb = _eg.TrimSpace(_eaeb)
		if _eaeb == "" {
			continue
		}
		_gffff, _cagb := _bbcab._cbcec.FontMap[_dedd]
		if _cagb {
			return _gffff
		}
		_dbegd, _cagb := map[string]_fgd.StdFontName{"\u0063o\u0075\u0072\u0069\u0065\u0072": _fgd.CourierName, "\u0063\u006f\u0075r\u0069\u0065\u0072\u002d\u0062\u006f\u006c\u0064": _fgd.CourierBoldName, "\u0063o\u0075r\u0069\u0065\u0072\u002d\u006f\u0062\u006c\u0069\u0071\u0075\u0065": _fgd.CourierObliqueName, "c\u006fu\u0072\u0069\u0065\u0072\u002d\u0062\u006f\u006cd\u002d\u006f\u0062\u006ciq\u0075\u0065": _fgd.CourierBoldObliqueName, "\u0068e\u006c\u0076\u0065\u0074\u0069\u0063a": _fgd.HelveticaName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0062\u006f\u006c\u0064": _fgd.HelveticaBoldName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u006f\u0062l\u0069\u0071\u0075\u0065": _fgd.HelveticaObliqueName, "\u0068\u0065\u006c\u0076et\u0069\u0063\u0061\u002d\u0062\u006f\u006c\u0064\u002d\u006f\u0062\u006c\u0069\u0071u\u0065": _fgd.HelveticaBoldObliqueName, "\u0073\u0079\u006d\u0062\u006f\u006c": _fgd.SymbolName, "\u007a\u0061\u0070\u0066\u002d\u0064\u0069\u006e\u0067\u0062\u0061\u0074\u0073": _fgd.ZapfDingbatsName, "\u0074\u0069\u006de\u0073": _fgd.TimesRomanName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064": _fgd.TimesBoldName, "\u0074\u0069\u006de\u0073\u002d\u0069\u0074\u0061\u006c\u0069\u0063": _fgd.TimesItalicName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064\u002d\u0069t\u0061\u006c\u0069\u0063": _fgd.TimesBoldItalicName}[_dedd]
		if _cagb {
			if _daeab, _egdff := _fgd.NewStandard14Font(_dbegd); _egdff == nil {
				return _daeab
			}
		}
		if _effbb := _bbcab.parseAttrPropList(_eaeb); len(_effbb) > 0 {
			if _eadf, _ebdbf := _effbb["\u0070\u0061\u0074\u0068"]; _ebdbf {
				_aefg := _fgd.NewPdfFontFromTTFFile
				if _dgcge, _dabeaf := _effbb["\u0074\u0079\u0070\u0065"]; _dabeaf && _dgcge == "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e" {
					_aefg = _fgd.NewCompositePdfFontFromTTFFile
				}
				if _gceeg, _cada := _aefg(_eadf); _cada != nil {
					_fec.Log.Debug("\u0043\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0060\u0025\u0073\u0060\u003a %\u0076\u002e", _eadf, _cada)
				} else {
					return _gceeg
				}
			}
		}
	}
	return _ecdg
}

// SetLineWidth sets the line width.
func (_eacaf *Polyline) SetLineWidth(lineWidth float64) { _eacaf._cdggc.LineWidth = lineWidth }

// SetMarkedContentID sets the marked content id for the line.
func (_fafa *Line) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_fafa._fbeb = &mcid
	_gada := _fgd.NewKDictionary()
	_gada.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_gada.K = _bc.MakeInteger(mcid)
	return _gada
}

// NewTable create a new Table with a specified number of columns.
func (_gfbc *Creator) NewTable(cols int) *Table { return _bagfg(cols) }

// SetMargins sets the Table's left, right, top, bottom margins.
func (_facff *Table) SetMargins(left, right, top, bottom float64) {
	_facff._gbcea.Left = left
	_facff._gbcea.Right = right
	_facff._gbcea.Top = top
	_facff._gbcea.Bottom = bottom
}
func _cgeac(_cceb string, _gfdg bool) string {
	_dgfgb := _cceb
	if _dgfgb == "" {
		return ""
	}
	_ecfdc := _fa.Paragraph{}
	_, _acdcd := _ecfdc.SetString(_cceb)
	if _acdcd != nil {
		return _dgfgb
	}
	_cagdf, _acdcd := _ecfdc.Order()
	if _acdcd != nil {
		return _dgfgb
	}
	_bbfcf := _cagdf.NumRuns()
	_dbbbe := make([]string, _bbfcf)
	for _bgdga := 0; _bgdga < _cagdf.NumRuns(); _bgdga++ {
		_dbdbgc := _cagdf.Run(_bgdga)
		_gbfb := _dbdbgc.String()
		if _dbdbgc.Direction() == _fa.RightToLeft {
			_gbfb = _fa.ReverseString(_gbfb)
		}
		if _gfdg {
			_dbbbe[_bgdga] = _gbfb
		} else {
			_dbbbe[_bbfcf-1] = _gbfb
		}
		_bbfcf--
	}
	if len(_dbbbe) != _cagdf.NumRuns() {
		return _cceb
	}
	_dgfgb = _eg.Join(_dbbbe, "")
	return _dgfgb
}
func _bgdfa(_beafe *templateProcessor, _dgfae *templateNode) (interface{}, error) {
	return _beafe.parseTextChunk(_dgfae, nil)
}

type pageTransformations struct {
	_acb  *_de.Matrix
	_bbed bool
	_fcge bool
}

// SetIndent sets the left offset of the list when nested into another list.
func (_effcd *List) SetIndent(indent float64) { _effcd._aedbb = indent; _effcd._ebfa = false }

// FillColor returns the fill color of the rectangle.
func (_bcfcd *Rectangle) FillColor() Color { return _bcfcd._dgab }

// SetAddressStyle sets the style properties used to render the content of
// the invoice address sections.
func (_cfddg *Invoice) SetAddressStyle(style TextStyle) { _cfddg._efdd = style }
func _aagf() *FilledCurve {
	_egcbb := FilledCurve{}
	_egcbb._acce = []_gga.CubicBezierCurve{}
	return &_egcbb
}

// SetViewerPreferences sets the viewer preferences for the PDF document.
func (_baa *Creator) SetViewerPreferences(viewerPreferences *_fgd.ViewerPreferences) {
	_baa._bbcg = viewerPreferences
}
func _dedgg(_bcge [][]_gga.Point) *Polygon {
	return &Polygon{_acdc: &_gga.Polygon{Points: _bcge}, _afdf: 1.0, _bgde: 1.0}
}

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_gfda *LinearShading) SetBoundingBox(x, y, width, height float64) {
	_gfda._gafag = &_fgd.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}
func (_edf *Paragraph) getMaxLineWidth() float64 {
	if _edf._ebdc == nil || (_edf._ebdc != nil && len(_edf._ebdc) == 0) {
		_edf.wrapText()
	}
	var _ecaa float64
	for _, _abea := range _edf._ebdc {
		_faca := _edf.getTextLineWidth(_abea)
		if _faca > _ecaa {
			_ecaa = _faca
		}
	}
	return _ecaa
}

const (
	CellHorizontalAlignmentLeft CellHorizontalAlignment = iota
	CellHorizontalAlignmentCenter
	CellHorizontalAlignmentRight
)

// SetMarkedContentID sets the marked content ID.
func (_gead *Rectangle) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_gead._gbdag = &mcid
	_ebff := _fgd.NewKDictionary()
	_ebff.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_ebff.K = _bc.MakeInteger(mcid)
	return _ebff
}

// NewDivision returns a new Division container component.
func (_bfeg *Creator) NewDivision() *Division { return _dcbc() }

// TextStyle is a collection of properties that can be assigned to a text chunk.
type TextStyle struct {

	// Color represents the color of the text.
	Color Color

	// OutlineColor represents the color of the text outline.
	OutlineColor Color

	// MultiFont represents an encoder that accepts multiple fonts and selects the correct font for encoding.
	MultiFont *_fgd.MultipleFontEncoder

	// Font represents the font the text will use.
	Font *_fgd.PdfFont

	// FontSize represents the size of the font.
	FontSize float64

	// OutlineSize represents the thickness of the text outline.
	OutlineSize float64

	// CharSpacing represents the character spacing.
	CharSpacing float64

	// HorizontalScaling represents the percentage to horizontally scale
	// characters by (default: 100). Values less than 100 will result in
	// narrower text while values greater than 100 will result in wider text.
	HorizontalScaling float64

	// RenderingMode represents the rendering mode.
	RenderingMode TextRenderingMode

	// Underline specifies if the text chunk is underlined.
	Underline bool

	// UnderlineStyle represents the style of the line used to underline text.
	UnderlineStyle TextDecorationLineStyle

	// TextRise specifies a vertical adjustment for text. It is useful for
	// drawing subscripts/superscripts. A positive text rise value will
	// produce superscript text, while a negative one will result in
	// subscript text.
	TextRise float64
}

// SetAnchor set gradient position anchor.
// Default to center.
func (_gdcb *RadialShading) SetAnchor(anchor AnchorPoint) { _gdcb._effe = anchor }

// SetMargins sets the margins of the paragraph.
func (_fcbb *List) SetMargins(left, right, top, bottom float64) {
	_fcbb._gacb.Left = left
	_fcbb._gacb.Right = right
	_fcbb._gacb.Top = top
	_fcbb._gacb.Bottom = bottom
}
func (_ecgd *Invoice) setCellBorder(_dedfa *TableCell, _beaf *InvoiceCell) {
	for _, _eedb := range _beaf.BorderSides {
		_dedfa.SetBorder(_eedb, CellBorderStyleSingle, _beaf.BorderWidth)
	}
	_dedfa.SetBorderColor(_beaf.BorderColor)
}

// NewBlock creates a new Block with specified width and height.
func NewBlock(width float64, height float64) *Block {
	_bbf := &Block{}
	_bbf._cb = &_dg.ContentStreamOperations{}
	_bbf._dgd = _fgd.NewPdfPageResources()
	_bbf._da = width
	_bbf._ace = height
	return _bbf
}

// SetMarkedContentID sets the marked content ID.
func (_dcdg *PolyBezierCurve) SetMarkedContentID(mcid int64) *_fgd.KDict {
	_dcdg._daef = &mcid
	_cafb := _fgd.NewKDictionary()
	_cafb.S = _bc.MakeName(_fgd.StructureTypeFigure)
	_cafb.K = _bc.MakeInteger(mcid)
	return _cafb
}

// Insert adds a new text chunk at the specified position in the paragraph.
func (_ccde *StyledParagraph) Insert(index uint, text string) *TextChunk {
	_bacd := uint(len(_ccde._ecec))
	if index > _bacd {
		index = _bacd
	}
	_fcfe := NewTextChunk(text, _ccde._fgbg)
	_ccde._ecec = append(_ccde._ecec[:index], append([]*TextChunk{_fcfe}, _ccde._ecec[index:]...)...)
	_ccde.wrapText()
	return _fcfe
}

// TextChunk represents a chunk of text along with a particular style.
type TextChunk struct {

	// The text that is being rendered in the PDF.
	Text string

	// The style of the text being rendered.
	Style  TextStyle
	_abefd *_fgd.PdfAnnotation
	_ebcfe bool

	// The vertical alignment of the text chunk.
	VerticalAlignment TextVerticalAlignment
}

// Positioning represents the positioning type for drawing creator components (relative/absolute).
type Positioning int

// BorderWidth returns the border width of the ellipse.
func (_ggaaf *Ellipse) BorderWidth() float64 { return _ggaaf._gfbg }
func (_dfcf *StyledParagraph) wrapChunks(_abdgb bool) error {
	if !_dfcf._fegca || int(_dfcf._ffcfd) <= 0 {
		_dfcf._bdfce = [][]*TextChunk{_dfcf._ecec}
		return nil
	}
	if _dfcf._eadc {
		_dfcf.wrapWordChunks()
	}
	_dfcf._bdfce = [][]*TextChunk{}
	var _dcfdd []*TextChunk
	var _gadb float64
	_agbb := _fe.IsSpace
	if !_abdgb {
		_agbb = func(rune) bool { return false }
	}
	_eaaf := _dgbbf(_dfcf._ffcfd*1000.0, 0.000001)
	for _, _gcdf := range _dfcf._ecec {
		_dcge := _gcdf.Style
		_ddgab := _gcdf._abefd
		_faff := _gcdf.VerticalAlignment
		var (
			_dceaf []rune
			_bdbd  []float64
		)
		_cead := _eedbc(_gcdf.Text)
		for _, _fadf := range _gcdf.Text {
			if _fadf == '\u000A' {
				if !_abdgb {
					_dceaf = append(_dceaf, _fadf)
				}
				_dcfdd = append(_dcfdd, &TextChunk{Text: _eg.TrimRightFunc(string(_dceaf), _agbb), Style: _dcge, _abefd: _bggbd(_ddgab), VerticalAlignment: _faff})
				_dfcf._bdfce = append(_dfcf._bdfce, _dcfdd)
				_dcfdd = nil
				_gadb = 0
				_dceaf = nil
				_bdbd = nil
				continue
			}
			_gddc := _fadf == ' '
			_cbefb, _cedbc := _dcge.Font.GetRuneMetrics(_fadf)
			if _cbefb.Wx == 0 && _dcge.MultiFont != nil || _dcge.MultiFont != nil && !_cedbc {
				_cbefb, _cedbc = _dcge.MultiFont.GetRuneMetrics(_fadf)
			}
			if !_cedbc {
				_fec.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _fadf)
				return _bd.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
			}
			_gfcae := _dcge.FontSize * _cbefb.Wx * _dcge.horizontalScale()
			_afdaa := _gfcae
			if !_gddc {
				_afdaa = _gfcae + _dcge.CharSpacing*1000.0
			}
			if _gadb+_gfcae > _eaaf {
				_bgffc := -1
				if !_gddc {
					for _bddcd := len(_dceaf) - 1; _bddcd >= 0; _bddcd-- {
						if _dceaf[_bddcd] == ' ' {
							_bgffc = _bddcd
							break
						}
					}
				}
				if _dfcf._eadc {
					_aafeb := len(_dcfdd)
					if _aafeb > 0 {
						_dcfdd[_aafeb-1].Text = _eg.TrimRightFunc(_dcfdd[_aafeb-1].Text, _agbb)
						_dfcf._bdfce = append(_dfcf._bdfce, _dcfdd)
						_dcfdd = []*TextChunk{}
					}
					_dceaf = append(_dceaf, _fadf)
					_bdbd = append(_bdbd, _afdaa)
					if _bgffc >= 0 {
						_dceaf = _dceaf[_bgffc+1:]
						_bdbd = _bdbd[_bgffc+1:]
					}
					_gadb = 0
					for _, _fgec := range _bdbd {
						_gadb += _fgec
					}
					if _gadb > _eaaf {
						_bdffe := string(_dceaf[:len(_dceaf)-1])
						_bdffe = _cgeac(_bdffe, _cead)
						if !_abdgb && _gddc {
							_bdffe += "\u0020"
						}
						_dcfdd = append(_dcfdd, &TextChunk{Text: _eg.TrimRightFunc(_bdffe, _agbb), Style: _dcge, _abefd: _bggbd(_ddgab), VerticalAlignment: _faff})
						_dfcf._bdfce = append(_dfcf._bdfce, _dcfdd)
						_dcfdd = []*TextChunk{}
						_dceaf = []rune{_fadf}
						_bdbd = []float64{_afdaa}
						_gadb = _afdaa
					}
					continue
				}
				_cegg := string(_dceaf)
				if _bgffc >= 0 {
					_cegg = string(_dceaf[0 : _bgffc+1])
					_dceaf = _dceaf[_bgffc+1:]
					_dceaf = append(_dceaf, _fadf)
					_bdbd = _bdbd[_bgffc+1:]
					_bdbd = append(_bdbd, _afdaa)
					_gadb = 0
					for _, _cfda := range _bdbd {
						_gadb += _cfda
					}
				} else {
					if _gddc {
						_gadb = 0
						_dceaf = []rune{}
						_bdbd = []float64{}
					} else {
						_gadb = _afdaa
						_dceaf = []rune{_fadf}
						_bdbd = []float64{_afdaa}
					}
				}
				_cegg = _cgeac(_cegg, _cead)
				if !_abdgb && _gddc {
					_cegg += "\u0020"
				}
				_dcfdd = append(_dcfdd, &TextChunk{Text: _eg.TrimRightFunc(_cegg, _agbb), Style: _dcge, _abefd: _bggbd(_ddgab), VerticalAlignment: _faff})
				_dfcf._bdfce = append(_dfcf._bdfce, _dcfdd)
				_dcfdd = []*TextChunk{}
			} else {
				_gadb += _afdaa
				_dceaf = append(_dceaf, _fadf)
				_bdbd = append(_bdbd, _afdaa)
			}
		}
		if len(_dceaf) > 0 {
			_bgdb := _cgeac(string(_dceaf), _cead)
			_dcfdd = append(_dcfdd, &TextChunk{Text: _bgdb, Style: _dcge, _abefd: _bggbd(_ddgab), VerticalAlignment: _faff})
		}
	}
	if len(_dcfdd) > 0 {
		_dfcf._bdfce = append(_dfcf._bdfce, _dcfdd)
	}
	return nil
}
func (_eeccd *templateProcessor) parseBorderRadiusAttr(_ecfgc, _cdbcb string) (_eeeda, _cade, _ccbb, _bgbf float64) {
	_fec.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0062o\u0072\u0064\u0065r\u0020\u0072\u0061\u0064\u0069\u0075\u0073\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _ecfgc, _cdbcb)
	switch _caegg := _eg.Fields(_cdbcb); len(_caegg) {
	case 1:
		_eeeda, _ = _fg.ParseFloat(_caegg[0], 64)
		_cade = _eeeda
		_ccbb = _eeeda
		_bgbf = _eeeda
	case 2:
		_eeeda, _ = _fg.ParseFloat(_caegg[0], 64)
		_ccbb = _eeeda
		_cade, _ = _fg.ParseFloat(_caegg[1], 64)
		_bgbf = _cade
	case 3:
		_eeeda, _ = _fg.ParseFloat(_caegg[0], 64)
		_cade, _ = _fg.ParseFloat(_caegg[1], 64)
		_bgbf = _cade
		_ccbb, _ = _fg.ParseFloat(_caegg[2], 64)
	case 4:
		_eeeda, _ = _fg.ParseFloat(_caegg[0], 64)
		_cade, _ = _fg.ParseFloat(_caegg[1], 64)
		_ccbb, _ = _fg.ParseFloat(_caegg[2], 64)
		_bgbf, _ = _fg.ParseFloat(_caegg[3], 64)
	}
	return _eeeda, _cade, _ccbb, _bgbf
}

// Notes returns the notes section of the invoice as a title-content pair.
func (_ccga *Invoice) Notes() (string, string) { return _ccga._bbgb[0], _ccga._bbgb[1] }

// TitleStyle returns the style properties used to render the invoice title.
func (_dbcb *Invoice) TitleStyle() TextStyle { return _dbcb._adaa }
func _dcgcc(_ffgcg *templateProcessor, _dcefg *templateNode) (interface{}, error) {
	return _ffgcg.parseImage(_dcefg)
}

// GeneratePageBlocks generate the Page blocks. Draws the Image on a block, implementing the Drawable interface.
func (_fgcdd *Image) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	if _fgcdd._gacc == nil {
		if _dgefe := _fgcdd.makeXObject(); _dgefe != nil {
			return nil, ctx, _dgefe
		}
	}
	var _ceae []*Block
	_gbcg := ctx
	_daed := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _fgcdd._accd.IsRelative() {
		_fgcdd.applyFitMode(ctx.Width)
		ctx.X += _fgcdd._cdbc.Left
		ctx.Y += _fgcdd._cdbc.Top
		ctx.Width -= _fgcdd._cdbc.Left + _fgcdd._cdbc.Right
		ctx.Height -= _fgcdd._cdbc.Top + _fgcdd._cdbc.Bottom
		if _fgcdd._fegb > ctx.Height {
			_ceae = append(_ceae, _daed)
			_daed = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_ebge := ctx
			_ebge.Y = ctx.Margins.Top + _fgcdd._cdbc.Top
			_ebge.X = ctx.Margins.Left + _fgcdd._cdbc.Left
			_ebge.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _fgcdd._cdbc.Top - _fgcdd._cdbc.Bottom
			_ebge.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _fgcdd._cdbc.Left - _fgcdd._cdbc.Right
			ctx = _ebge
		}
	} else {
		ctx.X = _fgcdd._afggb
		ctx.Y = _fgcdd._dbccc
	}
	ctx, _gegfg := _edebe(_daed, _fgcdd, ctx)
	if _gegfg != nil {
		return nil, ctx, _gegfg
	}
	_ceae = append(_ceae, _daed)
	if _fgcdd._accd.IsAbsolute() {
		ctx = _gbcg
	} else {
		ctx.X = _gbcg.X
		ctx.Width = _gbcg.Width
		ctx.Y += _fgcdd._cdbc.Bottom
	}
	return _ceae, ctx, nil
}

// SetTitleStyle sets the style properties of the invoice title.
func (_ccfd *Invoice) SetTitleStyle(style TextStyle) { _ccfd._adaa = style }

// AddTotalLine adds a new line in the invoice totals table.
func (_debdg *Invoice) AddTotalLine(desc, value string) (*InvoiceCell, *InvoiceCell) {
	_dbea := &InvoiceCell{_debdg._afcea, desc}
	_gaagg := &InvoiceCell{_debdg._afcea, value}
	_debdg._aegd = append(_debdg._aegd, [2]*InvoiceCell{_dbea, _gaagg})
	return _dbea, _gaagg
}

// SetFillOpacity sets the fill opacity of the rectangle.
func (_dbcaf *Rectangle) SetFillOpacity(opacity float64) { _dbcaf._afdae = opacity }
func _faad(_ecaae ...interface{}) []interface{}          { return _ecaae }

// GetRowHeight returns the height of the specified row.
func (_eeeab *Table) GetRowHeight(row int) (float64, error) {
	if row < 1 || row > len(_eeeab._gecdg) {
		return 0, _bd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	return _eeeab._gecdg[row-1], nil
}
func (_gfaga *templateProcessor) parsePageBreak(_bcgea *templateNode) (interface{}, error) {
	return _gbcfe(), nil
}

// UnsupportedRuneError is an error that occurs when there is unsupported glyph being used.
type UnsupportedRuneError struct {
	Message string
	Rune    rune
}

// Width is not used. The list component is designed to fill into the available
// width depending on the context. Returns 0.
func (_caag *List) Width() float64 { return 0 }
func _ddgga(_eccdc *_fgd.PdfAnnotationLink) *_fgd.PdfAnnotationLink {
	if _eccdc == nil {
		return nil
	}
	_eedag := _fgd.NewPdfAnnotationLink()
	_eedag.BS = _eccdc.BS
	_eedag.A = _eccdc.A
	if _fecgd, _aefde := _eccdc.GetAction(); _aefde == nil && _fecgd != nil {
		_eedag.SetAction(_fecgd)
	}
	if _gacff, _fadfb := _eccdc.Dest.(*_bc.PdfObjectArray); _fadfb {
		_eedag.Dest = _bc.MakeArray(_gacff.Elements()...)
	}
	return _eedag
}

// NewImageFromGoImage creates an Image from a go image.Image data structure.
func (_ffge *Creator) NewImageFromGoImage(goimg _e.Image) (*Image, error) { return _cgcd(goimg) }

// NewCurve returns new instance of Curve between points (x1,y1) and (x2, y2) with control point (cx,cy).
func (_acae *Creator) NewCurve(x1, y1, cx, cy, x2, y2 float64) *Curve {
	return _dgbe(x1, y1, cx, cy, x2, y2)
}

// NewInvoice returns an instance of an empty invoice.
func (_adc *Creator) NewInvoice() *Invoice {
	_eeed := _adc.NewTextStyle()
	_eeed.Font = _adc._ceb
	return _gbfdd(_adc.NewTextStyle(), _eeed)
}
func (_afcf *templateProcessor) renderNode(_fddd *templateNode) error {
	_fdbga := _fddd._bedcd
	if _fdbga == nil {
		return nil
	}
	_abaga := _fddd._adfdg.Name.Local
	_aacde, _dbee := _faba[_abaga]
	if !_dbee {
		_afcf.nodeLogDebug(_fddd, "I\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e \u0053\u006b\u0069p\u0070i\u006e\u0067\u002e", _abaga)
		return nil
	}
	var _faaf interface{}
	if _fddd._cefd != nil && _fddd._cefd._bedcd != nil {
		_fbgea := _fddd._cefd._adfdg.Name.Local
		if _, _dbee = _aacde._geca[_fbgea]; !_dbee {
			_afcf.nodeLogDebug(_fddd, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _fbgea, _abaga)
			return _gdbeeg
		}
		_faaf = _fddd._cefd._bedcd
	} else {
		_acgg := "\u0063r\u0065\u0061\u0074\u006f\u0072"
		switch _afcf._ccfe.(type) {
		case *Block:
			_acgg = "\u0062\u006c\u006fc\u006b"
		}
		if _, _dbee = _aacde._geca[_acgg]; !_dbee {
			_afcf.nodeLogDebug(_fddd, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _acgg, _abaga)
			return _gdbeeg
		}
		_faaf = _afcf._ccfe
	}
	switch _agbba := _faaf.(type) {
	case componentRenderer:
		_bdcec, _ggcbb := _fdbga.(Drawable)
		if !_ggcbb {
			_afcf.nodeLogError(_fddd, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u002e", _abaga, _fdbga)
			return _dggef
		}
		_bcca := _agbba.Draw(_bdcec)
		if _bcca != nil {
			return _afcf.nodeError(_fddd, "\u0043\u0061\u006en\u006f\u0074\u0020\u0064r\u0061\u0077\u0073\u0020\u0074\u0061\u0067 \u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u0073\u002e", _abaga, _fdbga, _bcca)
		}
	case *Division:
		switch _afged := _fdbga.(type) {
		case *Background:
			_agbba.SetBackground(_afged)
		case VectorDrawable:
			_egfa := _agbba.Add(_afged)
			if _egfa != nil {
				return _afcf.nodeError(_fddd, "\u0043a\u006e\u006eo\u0074\u0020\u0061d\u0064\u0020\u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054)\u0020\u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0044\u0069\u0076i\u0073\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u002e", _abaga, _fdbga, _egfa)
			}
		}
	case *TableCell:
		_afcff, _dfcfg := _fdbga.(VectorDrawable)
		if !_dfcfg {
			_afcf.nodeLogError(_fddd, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029 \u0069\u0073\u0020\u006e\u006f\u0074 \u0061\u0020\u0076\u0065\u0063\u0074\u006f\u0072\u0020\u0064\u0072\u0061\u0077a\u0062\u006c\u0065\u002e", _abaga, _fdbga)
			return _dggef
		}
		_gecdgg := _agbba.SetContent(_afcff)
		if _gecdgg != nil {
			return _afcf.nodeError(_fddd, "C\u0061\u006e\u006e\u006f\u0074\u0020\u0061\u0064\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073> \u0028\u0025\u0054\u0029 \u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0074\u0061bl\u0065\u0020c\u0065\u006c\u006c\u003a\u0020\u0025\u0073\u002e", _abaga, _fdbga, _gecdgg)
		}
	case *StyledParagraph:
		_ccdc, _cgggg := _fdbga.(*TextChunk)
		if !_cgggg {
			_afcf.nodeLogError(_fddd, "\u0054\u0061\u0067 <\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020i\u0073 \u006eo\u0074 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e", _abaga, _fdbga)
			return _dggef
		}
		_agbba.appendChunk(_ccdc)
	case *Chapter:
		switch _ddaac := _fdbga.(type) {
		case *Chapter:
			return nil
		case *Paragraph:
			if _fddd._adfdg.Name.Local == "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067" {
				return nil
			}
			_bgade := _agbba.Add(_ddaac)
			if _bgade != nil {
				return _afcf.nodeError(_fddd, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _abaga, _fdbga, _bgade)
			}
		case Drawable:
			_feggc := _agbba.Add(_ddaac)
			if _feggc != nil {
				return _afcf.nodeError(_fddd, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _abaga, _fdbga, _feggc)
			}
		}
	case *List:
		switch _gdbbg := _fdbga.(type) {
		case *TextChunk:
		case *listItem:
			_agbba._ffegf = append(_agbba._ffegf, _gdbbg)
		default:
			_afcf.nodeLogError(_fddd, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073>\u0020\u0028\u0025T\u0029\u0020\u0069\u0073 \u006e\u006f\u0074\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u0074\u0065\u006d\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _abaga, _fdbga)
		}
	case *listItem:
		switch _dfbe := _fdbga.(type) {
		case *TextChunk:
		case *StyledParagraph:
			_agbba._cdeda = _dfbe
		case *List:
			if _dfbe._ebfa {
				_dfbe._aedbb = 15
			}
			_agbba._cdeda = _dfbe
		case *Image:
			_agbba._cdeda = _dfbe
		case *Division:
			_agbba._cdeda = _dfbe
		case *Table:
			_agbba._cdeda = _dfbe
		default:
			_afcf.nodeLogError(_fddd, "\u0054\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u002e", _abaga, _fdbga)
			return _dggef
		}
	}
	return nil
}

// InvoiceCell represents any cell belonging to a table from the invoice
// template. The main tables are the invoice information table, the line
// items table and totals table. Contains the text value of the cell and
// the style properties of the cell.
type InvoiceCell struct {
	InvoiceCellProps
	Value string
}

func (_dbgd *Paragraph) getTextLineWidth(_ggcb string) float64 {
	var _cgdff float64
	for _, _ecdbe := range _ggcb {
		if _ecdbe == '\u000A' {
			continue
		}
		_efbb, _cbcg := _dbgd._gbca.GetRuneMetrics(_ecdbe)
		if !_cbcg {
			_fec.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _ecdbe, _ecdbe)
			return -1
		}
		_cgdff += _dbgd._acdge * _efbb.Wx
	}
	return _cgdff
}
func (_efeec *templateProcessor) parseCellVerticalAlignmentAttr(_agcf, _dgfdg string) CellVerticalAlignment {
	_fec.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u0065\u006c\u006c\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _agcf, _dgfdg)
	_dddca := map[string]CellVerticalAlignment{"\u0074\u006f\u0070": CellVerticalAlignmentTop, "\u006d\u0069\u0064\u0064\u006c\u0065": CellVerticalAlignmentMiddle, "\u0062\u006f\u0074\u0074\u006f\u006d": CellVerticalAlignmentBottom}[_dgfdg]
	return _dddca
}

// GeneratePageBlocks generates the table page blocks. Multiple blocks are
// generated if the contents wrap over multiple pages.
// Implements the Drawable interface.
func (_fged *Table) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_eggd := _fged
	if _fged._fgaag {
		_eggd = _fged.clone()
	}
	return _ecdbc(_eggd, ctx)
}

// NewImage create a new image from a unidoc image (model.Image).
func (_cbfa *Creator) NewImage(img *_fgd.Image) (*Image, error) { return _eddeb(img) }

// SetBorderColor sets the border color.
func (_dbdc *PolyBezierCurve) SetBorderColor(color Color) { _dbdc._ffcd.BorderColor = _cfcee(color) }
func _ddfb(_gceef, _cbbe, _bfcb, _aabd float64) *Line {
	return &Line{_gbcgde: _gceef, _gcfa: _cbbe, _cfge: _bfcb, _ceef: _aabd, _ebgca: ColorBlack, _eabf: 1.0, _aebf: 1.0, _cgecc: []int64{1, 1}, _babe: PositionAbsolute}
}

// Height returns the total height of all rows.
func (_cccabe *Table) Height() float64 {
	_bccb := float64(0.0)
	for _, _dfcc := range _cccabe._gecdg {
		_bccb += _dfcc
	}
	return _bccb
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_aaeaa *TOCLine) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_afcfb := ctx
	_ggged, ctx, _agddb := _aaeaa._afgdd.GeneratePageBlocks(ctx)
	if _agddb != nil {
		return _ggged, ctx, _agddb
	}
	if _aaeaa._afdaf.IsRelative() {
		ctx.X = _afcfb.X
	}
	if _aaeaa._afdaf.IsAbsolute() {
		return _ggged, _afcfb, nil
	}
	return _ggged, ctx, nil
}

// SetWidthLeft sets border width for left.
func (_feb *border) SetWidthLeft(bw float64) { _feb._aca = bw }

// FilledCurve represents a closed path of Bezier curves with a border and fill.
type FilledCurve struct {
	_acce         []_gga.CubicBezierCurve
	FillEnabled   bool
	_ggea         Color
	BorderEnabled bool
	BorderWidth   float64
	_ccec         Color
	_deca         *int64
}

// SetFillOpacity sets the fill opacity.
func (_agde *Polygon) SetFillOpacity(opacity float64) { _agde._afdf = opacity }
func (_gggb *Paragraph) wrapText() error {
	if !_gggb._ecae || int(_gggb._gcbf) <= 0 {
		_gggb._ebdc = []string{_gggb._egea}
		return nil
	}
	_affedd := NewTextChunk(_gggb._egea, TextStyle{Font: _gggb._gbca, FontSize: _gggb._acdge})
	_ebad, _bdaaf := _affedd.Wrap(_gggb._gcbf)
	if _bdaaf != nil {
		return _bdaaf
	}
	if _gggb._dacad > 0 && len(_ebad) > _gggb._dacad {
		_ebad = _ebad[:_gggb._dacad]
	}
	_gggb._ebdc = _ebad
	return nil
}

// Color returns the color of the line.
func (_ffccf *Line) Color() Color { return _ffccf._ebgca }

// GeneratePageBlocks draws the rectangle on a new block representing the page. Implements the Drawable interface.
func (_cag *Rectangle) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_abcf  []*Block
		_deee  = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_effbg = ctx
		_fgfb  = _cag._caeg / 2
	)
	_gcccf := _cag._dfff.IsRelative()
	if _gcccf {
		_cag.applyFitMode(ctx.Width)
		ctx.X += _cag._gagb.Left + _fgfb
		ctx.Y += _cag._gagb.Top + _fgfb
		ctx.Width -= _cag._gagb.Left + _cag._gagb.Right
		ctx.Height -= _cag._gagb.Top + _cag._gagb.Bottom
		if _cag._cggg > ctx.Height {
			_abcf = append(_abcf, _deee)
			_deee = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_fdbcg := ctx
			_fdbcg.Y = ctx.Margins.Top + _cag._gagb.Top + _fgfb
			_fdbcg.X = ctx.Margins.Left + _cag._gagb.Left + _fgfb
			_fdbcg.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _cag._gagb.Top - _cag._gagb.Bottom
			_fdbcg.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _cag._gagb.Left - _cag._gagb.Right
			ctx = _fdbcg
		}
	} else {
		ctx.X = _cag._defb
		ctx.Y = _cag._cdgf
	}
	_dgag := _gga.Rectangle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _cag._cggg, Width: _cag._bcede, Height: _cag._cggg, BorderRadiusTopLeft: _cag._ggeb, BorderRadiusTopRight: _cag._cbbc, BorderRadiusBottomLeft: _cag._ggfa, BorderRadiusBottomRight: _cag._gcefb, Opacity: 1.0}
	if _cag._dgab != nil {
		_dgag.FillEnabled = true
		_gcca := _cfcee(_cag._dgab)
		_dfgd := _cacbe(_deee, _gcca, _cag._dgab, func() Rectangle {
			return Rectangle{_defb: _dgag.X, _cdgf: _dgag.Y, _bcede: _dgag.Width, _cggg: _dgag.Height}
		})
		if _dfgd != nil {
			return nil, ctx, _dfgd
		}
		_dgag.FillColor = _gcca
	}
	if _cag._bedf != nil && _cag._caeg > 0 {
		_dgag.BorderEnabled = true
		_dgag.BorderColor = _cfcee(_cag._bedf)
		_dgag.BorderWidth = _cag._caeg
	}
	_fcdb, _dagg := _deee.setOpacity(_cag._afdae, _cag._dbff)
	if _dagg != nil {
		return nil, ctx, _dagg
	}
	_eabfe, _, _dagg := _dgag.MarkedDraw(_fcdb, _cag._gbdag)
	if _dagg != nil {
		return nil, ctx, _dagg
	}
	if _dagg = _deee.addContentsByString(string(_eabfe)); _dagg != nil {
		return nil, ctx, _dagg
	}
	if _gcccf {
		ctx.X = _effbg.X
		ctx.Width = _effbg.Width
		_cgge := _cag._cggg + _fgfb
		ctx.Y += _cgge + _cag._gagb.Bottom
		ctx.Height -= _cgge
	} else {
		ctx = _effbg
	}
	_abcf = append(_abcf, _deee)
	return _abcf, ctx, nil
}

// Rectangle defines a rectangle with upper left corner at (x,y) and a specified width and height.  The rectangle
// can have a colored fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Rectangle struct {
	_defb  float64
	_cdgf  float64
	_bcede float64
	_cggg  float64
	_dfff  Positioning
	_dgab  Color
	_afdae float64
	_bedf  Color
	_caeg  float64
	_dbff  float64
	_ggeb  float64
	_cbbc  float64
	_ggfa  float64
	_gcefb float64
	_gagb  Margins
	_aaab  FitMode
	_gbdag *int64
}

var (
	PageSizeA3     = PageSize{297 * PPMM, 420 * PPMM}
	PageSizeA4     = PageSize{210 * PPMM, 297 * PPMM}
	PageSizeA5     = PageSize{148 * PPMM, 210 * PPMM}
	PageSizeLetter = PageSize{8.5 * PPI, 11 * PPI}
	PageSizeLegal  = PageSize{8.5 * PPI, 14 * PPI}
)

// SetMarkedContentID sets marked content ID.
func (_fgbb *Invoice) SetMarkedContentID(id int64) *_fgd.KDict { return nil }
func (_cgdfc *pageTransformations) transformBlock(_bfab *Block) {
	if _cgdfc._acb != nil {
		_bfab.transform(*_cgdfc._acb)
	}
}

// TextVerticalAlignment controls the vertical position of the text
// in a styled paragraph.
type TextVerticalAlignment int

// Width returns the width of the Paragraph.
func (_fede *StyledParagraph) Width() float64 {
	if _fede._fegca && int(_fede._ffcfd) > 0 {
		return _fede._ffcfd
	}
	return _fede.getTextWidth() / 1000.0
}

// CurRow returns the currently active cell's row number.
func (_gccda *Table) CurRow() int { _dcbafe := (_gccda._ecgdc-1)/_gccda._afacb + 1; return _dcbafe }

// LinearShading holds data for rendering a linear shading gradient.
type LinearShading struct {
	_dfae  *shading
	_gafag *_fgd.PdfRectangle
	_cgeb  float64
}

// Polygon represents a polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Polygon struct {
	_acdc  *_gga.Polygon
	_afdf  float64
	_bgde  float64
	_gcbdc Color
	_fdbg  *int64
}

func _ecdbc(_cccbb *Table, _bbac DrawContext) ([]*Block, DrawContext, error) {
	var _ecad []*Block
	_eegb := NewBlock(_bbac.PageWidth, _bbac.PageHeight)
	_cccbb.updateRowHeights(_bbac.Width - _cccbb._gbcea.Left - _cccbb._gbcea.Right)
	_cacbed := _cccbb._gbcea.Top
	if _cccbb._cgbfc.IsRelative() && !_cccbb._bcagb {
		_fbgc := _cccbb.Height()
		if _fbgc > _bbac.Height-_cccbb._gbcea.Top && _fbgc <= _bbac.PageHeight-_bbac.Margins.Top-_bbac.Margins.Bottom {
			_ecad = []*Block{NewBlock(_bbac.PageWidth, _bbac.PageHeight-_bbac.Y)}
			var _afbed error
			if _, _bbac, _afbed = _gbcfe().GeneratePageBlocks(_bbac); _afbed != nil {
				return nil, _bbac, _afbed
			}
			_cacbed = 0
		}
	}
	_afddc := _bbac
	if _cccbb._cgbfc.IsAbsolute() {
		_bbac.X = _cccbb._cbgfe
		_bbac.Y = _cccbb._cbdba
	} else {
		_bbac.X += _cccbb._gbcea.Left
		_bbac.Y += _cacbed
		_bbac.Width -= _cccbb._gbcea.Left + _cccbb._gbcea.Right
		_bbac.Height -= _cacbed
	}
	_ggcfd := _bbac.Width
	_dbdbg := _bbac.X
	_fdfa := _bbac.Y
	_gecf := _bbac.Height
	_cbdg := 0
	_bdeg, _dccba := -1, -1
	if _cccbb._ddcb {
		for _ggfg, _ebcbg := range _cccbb._efbe {
			if _ebcbg._deef < _cccbb._fefac {
				continue
			}
			if _ebcbg._deef > _cccbb._gefbf {
				break
			}
			if _bdeg < 0 {
				_bdeg = _ggfg
			}
			_dccba = _ggfg
		}
	}
	if _gacbe := _cccbb.wrapContent(_bbac); _gacbe != nil {
		return nil, _bbac, _gacbe
	}
	_cccbb.updateRowHeights(_bbac.Width - _cccbb._gbcea.Left - _cccbb._gbcea.Right)
	var (
		_badb   bool
		_bgccde int
		_abdc   int
		_cedg   bool
		_fdbd   int
		_gaccg  error
	)
	for _fgdea := 0; _fgdea < len(_cccbb._efbe); _fgdea++ {
		_efdcc := _cccbb._efbe[_fgdea]
		if _cgefc, _ebabb := _cccbb.getLastCellFromCol(_efdcc._bacg); _cgefc == _fgdea {
			if (_ebabb._deef + _ebabb._afddb - 1) < _cccbb._begg {
				for _ggbc := _efdcc._deef; _ggbc < _cccbb._begg; _ggbc++ {
					_aebfe := &TableCell{}
					_aebfe._deef = _ggbc + 1
					_aebfe._afddb = 1
					_aebfe._bacg = _efdcc._bacg
					_cccbb._efbe = append(_cccbb._efbe, _aebfe)
				}
			}
		}
		_badfe := _efdcc.width(_cccbb._cddfg, _ggcfd)
		_baec := float64(0.0)
		for _dffee := 0; _dffee < _efdcc._bacg-1; _dffee++ {
			_baec += _cccbb._cddfg[_dffee] * _ggcfd
		}
		_dbggc := float64(0.0)
		for _cbde := _cbdg; _cbde < _efdcc._deef-1; _cbde++ {
			_dbggc += _cccbb._gecdg[_cbde]
		}
		_bbac.Height = _gecf - _dbggc
		_eafg := float64(0.0)
		for _cabc := 0; _cabc < _efdcc._afddb; _cabc++ {
			_eafg += _cccbb._gecdg[_efdcc._deef+_cabc-1]
		}
		_fcffa := _cedg && _efdcc._deef != _fdbd
		_fdbd = _efdcc._deef
		if _fcffa || _eafg > _bbac.Height {
			if _cccbb._fgaag && !_cedg {
				_cedg, _gaccg = _cccbb.wrapRow(_fgdea, _bbac, _ggcfd)
				if _gaccg != nil {
					return nil, _bbac, _gaccg
				}
				if _cedg {
					_fgdea--
					continue
				}
			}
			_ecad = append(_ecad, _eegb)
			_eegb = NewBlock(_bbac.PageWidth, _bbac.PageHeight)
			_dbdbg = _bbac.Margins.Left + _cccbb._gbcea.Left
			_fdfa = _bbac.Margins.Top
			_bbac.Height = _bbac.PageHeight - _bbac.Margins.Top - _bbac.Margins.Bottom
			_bbac.Page++
			_gecf = _bbac.Height
			_cbdg = _efdcc._deef - 1
			_dbggc = 0
			_cedg = false
			if _cccbb._ddcb && _bdeg >= 0 {
				_bgccde = _fgdea
				_fgdea = _bdeg - 1
				_abdc = _cbdg
				_cbdg = _cccbb._fefac - 1
				_badb = true
				if _efdcc._afddb > (_cccbb._begg-_fdbd) || (_efdcc._afddb > 1 && _fgdea < 0) {
					_fec.Log.Debug("\u0054a\u0062\u006ce\u0020\u0068\u0065a\u0064\u0065\u0072\u0020\u0072\u006f\u0077s\u0070\u0061\u006e\u0020\u0065\u0078c\u0065\u0065\u0064\u0073\u0020\u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u002e")
					_badb = false
					_bdeg, _dccba = -1, -1
				}
				continue
			}
			if _fcffa {
				_fgdea--
				continue
			}
		}
		_bbac.Width = _badfe
		_bbac.X = _dbdbg + _baec
		_bbac.Y = _fdfa + _dbggc
		if _eafg > _bbac.PageHeight-_bbac.Margins.Top-_bbac.Margins.Bottom {
			_eafg = _bbac.PageHeight - _bbac.Margins.Top - _bbac.Margins.Bottom
		}
		_ebcc := _bfaa(_bbac.X, _bbac.Y, _badfe, _eafg)
		if _efdcc._cbgd != nil {
			_ebcc.SetFillColor(_efdcc._cbgd)
		}
		_ebcc.LineStyle = _efdcc._gaddf
		_ebcc._ffbb = _efdcc._faab
		_ebcc._cfc = _efdcc._cfdbdc
		_ebcc._cggf = _efdcc._bgcce
		_ebcc._fagg = _efdcc._bggf
		if _efdcc._ebgcd != nil {
			_ebcc.SetColorLeft(_efdcc._ebgcd)
		}
		if _efdcc._aegdf != nil {
			_ebcc.SetColorBottom(_efdcc._aegdf)
		}
		if _efdcc._fcbbf != nil {
			_ebcc.SetColorRight(_efdcc._fcbbf)
		}
		if _efdcc._gefce != nil {
			_ebcc.SetColorTop(_efdcc._gefce)
		}
		_ebcc.SetWidthBottom(_efdcc._eabab)
		_ebcc.SetWidthLeft(_efdcc._caef)
		_ebcc.SetWidthRight(_efdcc._adec)
		_ebcc.SetWidthTop(_efdcc._cddebf)
		_gadf := NewBlock(_eegb._da, _eegb._ace)
		_ggdf := _eegb.Draw(_ebcc)
		if _ggdf != nil {
			_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ggdf)
		}
		if _efdcc._aafg != nil {
			_bdddb := _efdcc._aafg.Width()
			_adce := _efdcc._aafg.Height()
			_cebd := 0.0
			switch _begge := _efdcc._aafg.(type) {
			case *Paragraph:
				if _begge._ecae {
					_bdddb = _begge.getMaxLineWidth() / 1000.0
				}
				_ecgf, _abeg, _ := _begge.getTextMetrics()
				_fgbc, _geabg := _ecgf*_begge._fbcg, _abeg*_begge._fbcg
				_adce = _adce - _geabg + _fgbc
				_cebd += _fgbc - _geabg
				_aeagc := 0.5
				if _cccbb._eadda {
					_aeagc = 0.3
				}
				switch _efdcc._bdfgg {
				case CellVerticalAlignmentTop:
					_cebd += _fgbc * _aeagc
				case CellVerticalAlignmentBottom:
					_cebd -= _fgbc * _aeagc
				}
				_bdddb += _begge._eebg.Left + _begge._eebg.Right
				_adce += _begge._eebg.Top + _begge._eebg.Bottom
			case *StyledParagraph:
				if _begge._fegca {
					_bdddb = _begge.getMaxLineWidth() / 1000.0
				}
				_gdgga, _dbdbe, _efde := _begge.getLineMetrics(0)
				_gefd, _fdbge := _gdgga*_begge._fgef, _dbdbe*_begge._fgef
				if _begge._ddec == TextVerticalAlignmentCenter {
					_cebd = _fdbge - (_dbdbe + (_gdgga+_efde-_dbdbe)/2 + (_fdbge-_dbdbe)/2)
				}
				if len(_begge._bdfce) == 1 {
					_adce = _gefd
				} else {
					_adce = _adce - _fdbge + _gefd
				}
				_cebd += _gefd - _fdbge
				switch _efdcc._bdfgg {
				case CellVerticalAlignmentTop:
					_cebd += _gefd * 0.5
				case CellVerticalAlignmentBottom:
					_cebd -= _gefd * 0.5
				}
				_bdddb += _begge._ccddg.Left + _begge._ccddg.Right
				_adce += _begge._ccddg.Top + _begge._ccddg.Bottom
			case *Table:
				_bdddb = _badfe
			case *List:
				_bdddb = _badfe
			case *Division:
				_bdddb = _badfe
			case *Chart:
				_bdddb = _badfe
			case *Line:
				_adce += _begge._fggd.Top + _begge._fggd.Bottom
				_cebd -= _begge.Height() / 2
			case *Image:
				_bdddb += _begge._cdbc.Left + _begge._cdbc.Right
				_adce += _begge._cdbc.Top + _begge._cdbc.Bottom
			}
			switch _efdcc._bbbff {
			case CellHorizontalAlignmentLeft:
				_bbac.X += _efdcc._fceba
				_bbac.Width -= _efdcc._fceba
			case CellHorizontalAlignmentCenter:
				if _dfaf := _badfe - _bdddb; _dfaf > 0 {
					_bbac.X += _dfaf / 2
					_bbac.Width -= _dfaf / 2
				}
			case CellHorizontalAlignmentRight:
				if _badfe > _bdddb {
					_bbac.X = _bbac.X + _badfe - _bdddb - _efdcc._fceba
					_bbac.Width -= _efdcc._fceba
				}
			}
			_gadddc := _bbac.Y
			_fbbcd := _bbac.Height
			_bbac.Y += _cebd
			switch _efdcc._bdfgg {
			case CellVerticalAlignmentTop:
			case CellVerticalAlignmentMiddle:
				if _fegfd := _eafg - _adce; _fegfd > 0 {
					_bbac.Y += _fegfd / 2
					_bbac.Height -= _fegfd / 2
				}
			case CellVerticalAlignmentBottom:
				if _eafg > _adce {
					_bbac.Y = _bbac.Y + _eafg - _adce
					_bbac.Height = _eafg
				}
			}
			_decgc := _eegb.DrawWithContext(_efdcc._aafg, _bbac)
			if _decgc != nil {
				if _bd.Is(_decgc, ErrContentNotFit) && !_fcffa {
					_eegb = _gadf
					_fcffa = true
					_fgdea--
					continue
				}
				_fec.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _decgc)
			}
			_bbac.Y = _gadddc
			_bbac.Height = _fbbcd
		}
		_bbac.Y += _eafg
		_bbac.Height -= _eafg
		if _badb && _fgdea+1 > _dccba {
			_fdfa += _dbggc + _eafg
			_gecf -= _eafg + _dbggc
			_cbdg = _abdc
			_fgdea = _bgccde - 1
			_badb = false
		}
	}
	_ecad = append(_ecad, _eegb)
	if _cccbb._cgbfc.IsAbsolute() {
		return _ecad, _afddc, nil
	}
	_bbac.X = _afddc.X
	_bbac.Width = _afddc.Width
	_bbac.Y += _cccbb._gbcea.Bottom
	_bbac.Height -= _cccbb._gbcea.Bottom
	return _ecad, _bbac, nil
}
func (_ggcf *LinearShading) shadingModel() *_fgd.PdfShadingType2 {
	_gedc := _gga.NewPoint(_ggcf._gafag.Llx+_ggcf._gafag.Width()/2, _ggcf._gafag.Lly+_ggcf._gafag.Height()/2)
	_edge := _gga.NewPoint(_ggcf._gafag.Llx, _ggcf._gafag.Lly+_ggcf._gafag.Height()/2).Add(-_gedc.X, -_gedc.Y).Rotate(_ggcf._cgeb).Add(_gedc.X, _gedc.Y)
	_edge = _gga.NewPoint(_fc.Max(_fc.Min(_edge.X, _ggcf._gafag.Urx), _ggcf._gafag.Llx), _fc.Max(_fc.Min(_edge.Y, _ggcf._gafag.Ury), _ggcf._gafag.Lly))
	_effgc := _gga.NewPoint(_ggcf._gafag.Urx, _ggcf._gafag.Lly+_ggcf._gafag.Height()/2).Add(-_gedc.X, -_gedc.Y).Rotate(_ggcf._cgeb).Add(_gedc.X, _gedc.Y)
	_effgc = _gga.NewPoint(_fc.Min(_fc.Max(_effgc.X, _ggcf._gafag.Llx), _ggcf._gafag.Urx), _fc.Min(_fc.Max(_effgc.Y, _ggcf._gafag.Lly), _ggcf._gafag.Ury))
	_ggbea := _fgd.NewPdfShadingType2()
	_ggbea.PdfShading.ShadingType = _bc.MakeInteger(2)
	_ggbea.PdfShading.ColorSpace = _fgd.NewPdfColorspaceDeviceRGB()
	_ggbea.PdfShading.AntiAlias = _bc.MakeBool(_ggcf._dfae._ccea)
	_ggbea.Coords = _bc.MakeArrayFromFloats([]float64{_edge.X, _edge.Y, _effgc.X, _effgc.Y})
	_ggbea.Extend = _bc.MakeArray(_bc.MakeBool(_ggcf._dfae._fece[0]), _bc.MakeBool(_ggcf._dfae._fece[1]))
	_ggbea.Function = _ggcf._dfae.generatePdfFunctions()
	return _ggbea
}

// SetPos sets the Table's positioning to absolute mode and specifies the upper-left corner
// coordinates as (x,y).
// Note that this is only sensible to use when the table does not wrap over multiple pages.
// TODO: Should be able to set width too (not just based on context/relative positioning mode).
func (_dedgd *Table) SetPos(x, y float64) {
	_dedgd._cgbfc = PositionAbsolute
	_dedgd._cbgfe = x
	_dedgd._cbdba = y
}
func (_gaea *Invoice) generateLineBlocks(_fegbb DrawContext) ([]*Block, DrawContext, error) {
	_edbf := _bagfg(len(_gaea._gffd))
	_edbf.SetMargins(0, 0, 25, 0)
	for _, _eagg := range _gaea._gffd {
		_gfcg := _ddge(_eagg.TextStyle)
		_gfcg.SetMargins(0, 0, 1, 0)
		_gfcg.Append(_eagg.Value)
		_dfadg := _edbf.NewCell()
		_dfadg.SetHorizontalAlignment(_eagg.Alignment)
		_dfadg.SetBackgroundColor(_eagg.BackgroundColor)
		_gaea.setCellBorder(_dfadg, _eagg)
		_dfadg.SetContent(_gfcg)
	}
	for _, _afbbg := range _gaea._agece {
		for _, _bafd := range _afbbg {
			_cbfgg := _ddge(_bafd.TextStyle)
			_cbfgg.SetMargins(0, 0, 3, 2)
			_cbfgg.Append(_bafd.Value)
			_bebc := _edbf.NewCell()
			_bebc.SetHorizontalAlignment(_bafd.Alignment)
			_bebc.SetBackgroundColor(_bafd.BackgroundColor)
			_gaea.setCellBorder(_bebc, _bafd)
			_bebc.SetContent(_cbfgg)
		}
	}
	return _edbf.GeneratePageBlocks(_fegbb)
}

// VectorDrawable is a Drawable with a specified width and height.
type VectorDrawable interface {
	Drawable

	// Width returns the width of the Drawable.
	Width() float64

	// Height returns the height of the Drawable.
	Height() float64
}

func (_gdfd grayColor) ToRGB() (float64, float64, float64) {
	return _gdfd._dbdf, _gdfd._dbdf, _gdfd._dbdf
}

// AddressHeadingStyle returns the style properties used to render the
// heading of the invoice address sections.
func (_edbaf *Invoice) AddressHeadingStyle() TextStyle { return _edbaf._gfdfe }
func _feff(_gcega string) ([]string, error) {
	var (
		_egbd  []string
		_gaeda []rune
	)
	for _, _dbaf := range _gcega {
		if _dbaf == '\u000A' {
			if len(_gaeda) > 0 {
				_egbd = append(_egbd, string(_gaeda))
			}
			_egbd = append(_egbd, string(_dbaf))
			_gaeda = nil
			continue
		}
		_gaeda = append(_gaeda, _dbaf)
	}
	if len(_gaeda) > 0 {
		_egbd = append(_egbd, string(_gaeda))
	}
	var _bdde []string
	for _, _gafc := range _egbd {
		_egceb := []rune(_gafc)
		_ccgg := _bb.NewScanner(_egceb)
		var _dece []rune
		for _fbebd := 0; _fbebd < len(_egceb); _fbebd++ {
			_, _ccba, _dccbd := _ccgg.Next()
			if _dccbd != nil {
				return nil, _dccbd
			}
			if _ccba == _bb.BreakProhibited || _fe.IsSpace(_egceb[_fbebd]) {
				_dece = append(_dece, _egceb[_fbebd])
				if _fe.IsSpace(_egceb[_fbebd]) {
					_bdde = append(_bdde, string(_dece))
					_dece = []rune{}
				}
				continue
			} else {
				if len(_dece) > 0 {
					_bdde = append(_bdde, string(_dece))
				}
				_dece = []rune{_egceb[_fbebd]}
			}
		}
		if len(_dece) > 0 {
			_bdde = append(_bdde, string(_dece))
		}
	}
	return _bdde, nil
}

// SellerAddress returns the seller address used in the invoice template.
func (_dafa *Invoice) SellerAddress() *InvoiceAddress { return _dafa._ddaa }

// MultiColCell makes a new cell with the specified column span and inserts it
// into the table at the current position.
func (_fdca *Table) MultiColCell(colspan int) *TableCell { return _fdca.MultiCell(1, colspan) }

// SetForms adds an Acroform to a PDF file.  Sets the specified form for writing.
func (_fdgb *Creator) SetForms(form *_fgd.PdfAcroForm) error { _fdgb._ecc = form; return nil }

// NewColorPoint creates a new color and point object for use in the gradient rendering process.
func NewColorPoint(color Color, point float64) *ColorPoint { return _bdab(color, point) }
func (_aceee *templateProcessor) parseFloatArray(_fcddf, _gedba string) []float64 {
	_fec.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0066\u006c\u006f\u0061\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fcddf, _gedba)
	_dcdggf := _eg.Fields(_gedba)
	_gefg := make([]float64, 0, len(_dcdggf))
	for _, _dafdd := range _dcdggf {
		_fedeg, _ := _fg.ParseFloat(_dafdd, 64)
		_gefg = append(_gefg, _fedeg)
	}
	return _gefg
}
func (_ebgf *Chapter) headingText() string {
	_ceg := _ebgf._faed
	if _dbde := _ebgf.headingNumber(); _dbde != "" {
		_ceg = _f.Sprintf("\u0025\u0073\u0020%\u0073", _dbde, _ceg)
	}
	return _ceg
}
func (_fb *Block) drawToPage(_gfd *_fgd.PdfPage) error {
	_bga := &_dg.ContentStreamOperations{}
	if _gfd.Resources == nil {
		_gfd.Resources = _fgd.NewPdfPageResources()
	}
	_ab := _gbg(_bga, _gfd.Resources, _fb._cb, _fb._dgd)
	if _ab != nil {
		return _ab
	}
	if _ab = _eec(_fb._dgd, _gfd.Resources); _ab != nil {
		return _ab
	}
	if _ab = _gfd.AppendContentBytes(_bga.Bytes(), true); _ab != nil {
		return _ab
	}
	for _, _ddb := range _fb._ed {
		_gfd.AddAnnotation(_ddb)
	}
	return nil
}

type containerDrawable interface {
	Drawable

	// ContainerComponent checks if the component is allowed to be added into provided 'container' and returns
	// preprocessed copy of itself. If the component is not changed it is allowed to return itself in a callback way.
	// If the component is not compatible with provided container this method should return an error.
	ContainerComponent(_dedfb Drawable) (Drawable, error)
}

func (_cfde *Table) wrapContent(_bbbfbd DrawContext) error {
	if _cfde._fgaag {
		return nil
	}
	_cfde.sortCells()
	_dgdab := func(_ddgc *TableCell, _aedbe int, _acgb int, _geee int) (_adfde int) {
		if _geee < 1 {
			return -1
		}
		_dagfb := 0
		for _cedd := _acgb + 1; _cedd < len(_cfde._efbe)-1; _cedd++ {
			_dedgdb := _cfde._efbe[_cedd]
			if _dedgdb._deef == _geee && _dagfb != _acgb {
				_dagfb = _cedd
				if (_dedgdb._bacg < _ddgc._bacg && _cfde._afacb > _dedgdb._bacg) || _ddgc._bacg < _cfde._afacb {
					continue
				}
				break
			}
		}
		_faac := float64(0.0)
		for _efba := 0; _efba < _ddgc._afddb; _efba++ {
			_faac += _cfde._gecdg[_ddgc._deef+_efba-1]
		}
		_fagada := _ddgc.width(_cfde._cddfg, _bbbfbd.Width)
		var (
			_bgage VectorDrawable
			_gedf  = false
		)
		switch _fffdbf := _ddgc._aafg.(type) {
		case *StyledParagraph:
			_gcgcd := _bbbfbd
			_gcgcd.Height = _fc.Floor(_faac - _fffdbf._ccddg.Top - _fffdbf._ccddg.Bottom - 0.5*_fffdbf.getTextHeight())
			_gcgcd.Width = _fagada
			_edfg, _dgeba, _abagf := _fffdbf.split(_gcgcd)
			if _abagf != nil {
				_fec.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _abagf.Error())
			}
			if _edfg != nil && _dgeba != nil {
				_cfde._efbe[_acgb]._aafg = _edfg
				_bgage = _dgeba
				_gedf = true
			}
		}
		_cfde._efbe[_acgb]._afddb = _ddgc._afddb
		_bbbfbd.Height = _bbbfbd.PageHeight - _bbbfbd.Margins.Top - _bbbfbd.Margins.Bottom
		_ceeg := _ddgc.cloneProps(nil)
		if _gedf {
			_ceeg._aafg = _bgage
		}
		_ceeg._afddb = _aedbe
		_ceeg._deef = _geee + 1
		_ceeg._bacg = _ddgc._bacg
		if _ceeg._deef+_ceeg._afddb-1 > _cfde._begg {
			for _faag := _cfde._begg; _faag < _ceeg._deef+_ceeg._afddb-1; _faag++ {
				_cfde._begg++
				_cfde._gecdg = append(_cfde._gecdg, _cfde._edefd)
			}
		}
		_cfde._efbe = append(_cfde._efbe[:_dagfb+1], append([]*TableCell{_ceeg}, _cfde._efbe[_dagfb+1:]...)...)
		return _dagfb + 1
	}
	_beeff := func(_bdabb *TableCell, _gfae int, _ecbb int, _bdfcec float64) (_ddaeb int) {
		_cabce := _bdabb.width(_cfde._cddfg, _bbbfbd.Width)
		_decc := _bdfcec
		_dbac := 1
		_ccfdf := _bbbfbd.Height
		if _ccfdf > 0 {
			for _decc > _ccfdf {
				_decc -= _bbbfbd.Height
				_ccfdf = _bbbfbd.PageHeight - _bbbfbd.Margins.Top - _bbbfbd.Margins.Bottom
				_dbac++
			}
		}
		var (
			_baccb VectorDrawable
			_bcebf = false
		)
		switch _bdee := _bdabb._aafg.(type) {
		case *StyledParagraph:
			_daag := _bbbfbd
			_daag.Height = _fc.Floor(_bbbfbd.Height - _bdee._ccddg.Top - _bdee._ccddg.Bottom - 0.5*_bdee.getTextHeight())
			_daag.Width = _cabce
			_dfgg, _fbga, _bfdd := _bdee.split(_daag)
			if _bfdd != nil {
				_fec.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _bfdd.Error())
			}
			if _dfgg != nil && _fbga != nil {
				_cfde._efbe[_gfae]._aafg = _dfgg
				_baccb = _fbga
				_bcebf = true
			}
		}
		if _dbac < 2 {
			return -1
		}
		if _cfde._efbe[_gfae]._deef+_dbac-1 > _cfde._begg {
			for _gcfab := 0; _gcfab < _dbac; _gcfab++ {
				_cfde._begg++
				_cfde._gecdg = append(_cfde._gecdg, _cfde._edefd)
			}
		}
		_eeda := _bdfcec / float64(_dbac)
		for _gccf := 0; _gccf < _dbac; _gccf++ {
			_cfde._gecdg[_ecbb+_gccf-1] = _eeda
		}
		_bbbfbd.Height = _bbbfbd.PageHeight - _bbbfbd.Margins.Top - _bbbfbd.Margins.Bottom
		_gbfddf := _bdabb.cloneProps(nil)
		if _bcebf {
			_gbfddf._aafg = _baccb
		}
		_gbfddf._afddb = 1
		_gbfddf._deef = _ecbb + _dbac - 1
		_gbfddf._bacg = _bdabb._bacg
		_cfde._efbe = append(_cfde._efbe, _gbfddf)
		return len(_cfde._efbe)
	}
	_bbfff := 1
	_bgfba := -1
	for _gdgb := 0; _gdgb < len(_cfde._efbe); _gdgb++ {
		_bcagg := _cfde._efbe[_gdgb]
		if _bgfba == _gdgb {
			_bbfff = _bcagg._deef
		}
		if _bcagg._afddb < 2 {
			if _fadc := _cfde._gecdg[_bcagg._deef-1]; _fadc > _bbbfbd.Height {
				_bgfba = _beeff(_bcagg, _gdgb, _bcagg._deef, _fadc)
				continue
			}
			continue
		}
		_bcafc := float64(0)
		for _fcee := 0; _fcee < _bcagg._afddb; _fcee++ {
			_bcafc += _cfde._gecdg[_bcagg._deef+_fcee-1]
		}
		_fdbcgc := float64(0)
		for _dfcbe := _bbfff - 1; _dfcbe < _bcagg._deef-1; _dfcbe++ {
			_fdbcgc += _cfde._gecdg[_dfcbe]
		}
		if _bcafc <= (_bbbfbd.Height - _fdbcgc) {
			continue
		}
		_gdcd := float64(0.0)
		_agbe := _bcagg._afddb
		_bfgfe := -1
		_ebaf := 1
		for _efegc := 1; _efegc <= _bcagg._afddb; _efegc++ {
			if (_gdcd + _cfde._gecdg[_bcagg._deef+_efegc-2]) > (_bbbfbd.Height - _fdbcgc) {
				_ebaf--
				break
			}
			_bfgfe = _bcagg._deef + _efegc - 1
			_agbe = _bcagg._afddb - _efegc
			_gdcd += _cfde._gecdg[_bcagg._deef+_efegc-2]
			_ebaf++
		}
		if _bcagg._afddb == _agbe {
			_bbbfbd.Height = _bbbfbd.PageHeight - _bbbfbd.Margins.Top - _bbbfbd.Margins.Bottom
			_bbfff = _bcagg._deef
			_gdgb--
			continue
		}
		if _agbe > 0 && _bcagg._afddb > _ebaf {
			_bcagg._afddb = _ebaf
			_bgfba = _dgdab(_bcagg, _agbe, _gdgb, _bfgfe)
			if _gdgb+1 == _bgfba {
				_gdgb--
			}
		}
		_bbfff = _bcagg._deef
	}
	_cfde.sortCells()
	return nil
}

type grayColor struct{ _dbdf float64 }

// CurvePolygon represents a curve polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type CurvePolygon struct {
	_dgc   *_gga.CurvePolygon
	_fafc  float64
	_bcefc float64
	_eeea  Color
	_becd  *int64
}

// AddSection adds a new content section at the end of the invoice.
func (_aage *Invoice) AddSection(title, content string) {
	_aage._fcfd = append(_aage._fcfd, [2]string{title, content})
}

// NewTextStyle creates a new text style object which can be used to style
// chunks of text.
// Default attributes:
// Font: Helvetica
// Font size: 10
// Encoding: WinAnsiEncoding
// Text color: black
func (_aece *Creator) NewTextStyle() TextStyle { return _ddaea(_aece._dff) }

// TOCLine represents a line in a table of contents.
// The component can be used both in the context of a
// table of contents component and as a standalone component.
// The representation of a table of contents line is as follows:
/*
         [number] [title]      [separator] [page]
   e.g.: Chapter1 Introduction ........... 1
*/
type TOCLine struct {
	_afgdd *StyledParagraph

	// Holds the text and style of the number part of the TOC line.
	Number TextChunk

	// Holds the text and style of the title part of the TOC line.
	Title TextChunk

	// Holds the text and style of the separator part of the TOC line.
	Separator TextChunk

	// Holds the text and style of the page part of the TOC line.
	Page    TextChunk
	_afage  float64
	_gfec   uint
	_ebgcdc float64
	_afdaf  Positioning
	_beed   float64
	_cdbff  float64
	_ebaab  int64
}

// HorizontalAlignment represents the horizontal alignment of components
// within a page.
type HorizontalAlignment int

// SetIncludeInTOC sets a flag to indicate whether or not to include in tOC.
func (_bbe *Chapter) SetIncludeInTOC(includeInTOC bool) { _bbe._aga = includeInTOC }
