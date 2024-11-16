/**
 *
 */

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bamzi/pdfext/extractor"
	"github.com/bamzi/pdfext/model"
	"github.com/bamzi/pdfext/pdfutil"
	"golang.org/x/text/unicode/norm"
)

func main() {
	filePath := ""
	text, err := extractTextFromPDF(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}

func extractTextFromPDF(filePath string) (string, error) {
	reader, file, err := model.NewPdfReaderFromFile(filePath, nil)
	if err != nil {
		return "", err
	}
	defer file.Close()

	numPages, err := reader.GetNumPages()
	if err != nil {
		return "", err
	}
	result := docTables{pageTables: make(map[int][]stringTable)}
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		tables, err := extractPageTables(reader, pageNum)
		if err != nil {
			return "", fmt.Errorf("extractPageTables failed.  err=%w",
				err)
		}
		result.pageTables[pageNum] = tables
	}

	var content strings.Builder
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		for _, table := range result.pageTables[pageNum] {
			text := table.csv()
			content.WriteString(text)
			content.WriteString("\n")
		}
	}
	return content.String(), nil
}

// extractPageTables extracts the tables from (1-offset) page number `pageNum` in opened
// PdfReader `pdfReader.
func extractPageTables(pdfReader *model.PdfReader, pageNum int) ([]stringTable, error) {
	page, err := pdfReader.GetPage(pageNum)
	if err != nil {
		return nil, err
	}
	if err := pdfutil.NormalizePage(page); err != nil {
		return nil, err
	}
	ex, err := extractor.New(page)
	if err != nil {
		return nil, err
	}
	pageText, _, _, err := ex.ExtractPageText()
	if err != nil {
		return nil, err
	}
	tables := pageText.Tables()
	stringTables := make([]stringTable, len(tables))
	for i, table := range tables {
		stringTables[i] = asStringTable(table)
	}
	return stringTables, nil
}

// docTables describes the tables in a document.
type docTables struct {
	pageTables map[int][]stringTable
}

// stringTable is the strings in TextTable.
type stringTable [][]string

// wh returns the width and height of table `t`.
func (t stringTable) wh() (int, int) {
	if len(t) == 0 {
		return 0, 0
	}
	return len(t[0]), len(t)
}

// csv returns `t` in CSV format.
func (t stringTable) csv() string {
	w, h := t.wh()
	b := new(bytes.Buffer)
	csvwriter := csv.NewWriter(b)
	for y, row := range t {
		if len(row) != w {
			err := fmt.Errorf("table = %d x %d row[%d]=%d %q", w, h, y, len(row), row)
			panic(err)
		}
		csvwriter.Write(row)
	}
	csvwriter.Flush()
	return b.String()
}

func (r *docTables) String() string {
	return r.describe(1)
}

// describe returns a string describing the tables in `r`.
//
//	                            (level 0)
//	%d pages %d tables          (level 1)
//	  page %d: %d tables        (level 2)
//	    table %d: %d x %d       (level 3)
//	        contents            (level 4)
func (r *docTables) describe(level int) string {
	if level == 0 || r.numTables() == 0 {
		return "\n"
	}
	var sb strings.Builder
	pageNumbers := r.pageNumbers()
	fmt.Fprintf(&sb, "%d pages %d tables\n", len(pageNumbers), r.numTables())
	if level <= 1 {
		return sb.String()
	}
	for _, pageNum := range r.pageNumbers() {
		tables := r.pageTables[pageNum]
		if len(tables) == 0 {
			continue
		}
		fmt.Fprintf(&sb, "   page %d: %d tables\n", pageNum, len(tables))
		if level <= 2 {
			continue
		}
		for i, table := range tables {
			w, h := table.wh()
			fmt.Fprintf(&sb, "      table %d: %d x %d\n", i+1, w, h)
			if level <= 3 || len(table) == 0 {
				continue
			}
			for _, row := range table {
				cells := make([]string, len(row))
				for i, cell := range row {
					if len(cell) > 0 {
						cells[i] = fmt.Sprintf("%q", cell)
					}
				}
				fmt.Fprintf(&sb, "        [%s]\n", strings.Join(cells, ", "))
			}
		}
	}
	return sb.String()
}

func (r *docTables) pageNumbers() []int {
	pageNums := make([]int, len(r.pageTables))
	i := 0
	for pageNum := range r.pageTables {
		pageNums[i] = pageNum
		i++
	}
	sort.Ints(pageNums)
	return pageNums
}

func (r *docTables) numTables() int {
	n := 0
	for _, tables := range r.pageTables {
		n += len(tables)
	}
	return n
}

// filter returns the tables in `r` that are at least `width` cells wide and `height` cells high.
func (r docTables) filter(width, height int) docTables {
	filtered := docTables{pageTables: make(map[int][]stringTable)}
	for pageNum, tables := range r.pageTables {
		var filteredTables []stringTable
		for _, table := range tables {
			if len(table[0]) >= width && len(table) >= height {
				filteredTables = append(filteredTables, table)
			}
		}
		if len(filteredTables) > 0 {
			filtered.pageTables[pageNum] = filteredTables
		}
	}
	return filtered
}

// asStringTable returns TextTable `table` as a stringTable.
func asStringTable(table extractor.TextTable) stringTable {
	cells := make(stringTable, table.H)
	for y, row := range table.Cells {
		cells[y] = make([]string, table.W)
		for x, cell := range row {
			cells[y][x] = cell.Text
		}
	}
	return normalizeTable(cells)
}

// normalizeTable returns `cells` with each cell normalized.
func normalizeTable(cells stringTable) stringTable {
	for y, row := range cells {
		for x, cell := range row {
			cells[y][x] = normalize(cell)
		}
	}
	return cells
}

// normalize returns a version of `text` that is NFKC normalized and has reduceSpaces() applied.
func normalize(text string) string {
	return reduceSpaces(norm.NFKC.String(text))
}

// reduceSpaces returns `text` with runs of spaces of any kind (spaces, tabs, line breaks, etc)
// reduced to a single space.
func reduceSpaces(text string) string {
	text = reSpace.ReplaceAllString(text, " ")
	return strings.Trim(text, " \t\n\r\v")
}

var reSpace = regexp.MustCompile(`(?m)\s+`)

// ------
// Updated extractTextFromPDF function
// func extractTextFromPDF(filePath string) (string, error) {
// 	f, r, err := pdf.Open(filePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()
//
// 	var content strings.Builder
//
// 	numPages := r.NumPage()
// 	for i := 1; i <= numPages; i++ {
// 		page := r.Page(i)
// 		if page.V.IsNull() {
// 			continue
// 		}
// 		// rows, _ := page.GetTextByRow()
// 		pageText, err := page.GetPlainText(nil)
// 		if err != nil {
// 			return "", err
// 		}
// 		content.WriteString(pageText)
// 		content.WriteString("\n")
// 	}
//
// 	return content.String(), nil
// }

// Extract text from PDF while preserving layout
// func extractTextFromPDF(filePath string) (string, error) {
// 	pdfReader, file, err := model.NewPdfReaderFromFile(filePath, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()
//
// 	numPages, err := pdfReader.GetNumPages()
// 	if err != nil {
// 		return "", err
// 	}
//
// 	var content strings.Builder
//
// 	for i := 1; i <= numPages; i++ {
// 		page, err := pdfReader.GetPage(i)
// 		if err != nil {
// 			return "", err
// 		}
//
// 		ext, err := extractor.New(page)
// 		if err != nil {
// 			return "", err
// 		}
//
// 		pageText, err := ext.ExtractText()
// 		if err != nil {
// 			return "", err
// 		}
//
// 		content.WriteString(pageText)
// 		content.WriteString("\n")
// 	}
//
// 	return content.String(), nil
// }
