package ctb

import (
	"fmt"
	"github.com/gookit/color"
	"math"
	"strings"
)

type LineData struct {
	Data  string
	Color color.Color
}

type Table struct {
	contentFieldMaxLen []int
	contentLine        [][]LineData
	tab                []LineData
	spacing            int
	prefixTab          string
	prefixContent      string
	prefixDisable      bool
}

func NewTable(F ...TableAttrFunc) *Table {
	t := &Table{
		spacing:       10,
		prefixTab:     " - ",
		prefixContent: " * ",
		prefixDisable: true,
	}
	TableAttrFuncs(F).Apply(t)
	return t
}

func (t *Table) SetPrefixTab(Prefix string) *Table {
	t.prefixTab = Prefix
	return t
}

func (t *Table) SetPrefixContent(Prefix string) {
	t.prefixContent = Prefix
}

func (t *Table) SetPrefixDisable(Disable bool) {
	t.prefixDisable = Disable
}

func (t *Table) SetSpacing(spacing int) {
	t.spacing = spacing
}

func (t *Table) SetTab(tab []LineData) {
	t.tab = tab
}

func (t *Table) SetDataOne(data []LineData) {
	t.contentLine = append(t.contentLine, data)
}

func (t *Table) SetDataAll(data [][]LineData) {
	for _, v := range data {
		t.contentLine = append(t.contentLine, v)
	}
}

func (t *Table) initContentMaxLen() {
	if len(t.contentLine) == 0 {
		return
	}
	if len(t.contentLine) == 1 || len(t.contentLine[0]) > len(t.contentLine[1]) {
		t.contentFieldMaxLen = make([]int, len(t.contentLine[0]))
	} else {
		t.contentFieldMaxLen = make([]int, len(t.contentLine[1]))
	}
}

func (t *Table) printLine() {
	for k, contentSlice := range t.contentLine {
		lineStr := ""
		if t.prefixDisable == false {
			if k == 0 && len(t.tab) > 0 {
				lineStr += t.prefixTab
			} else {
				lineStr += t.prefixContent
			}
		}
		for index, val := range contentSlice {
			space := t.contentFieldMaxLen[index] - len(val.Data) + t.spacing
			var data string
			if val.Color == 0 {
				data = val.Data
			} else {
				data = val.Color.Sprintf(val.Data)
			}
			lineStr += fmt.Sprintf("%s%s",
				data,
				strings.Repeat(" ", space),
			)
		}
		fmt.Println(lineStr)
	}
}

func (t *Table) readData() {
	for _, v1 := range t.contentLine {
		for k, v2 := range v1 {
			if len(v2.Data) > t.contentFieldMaxLen[k] {
				t.contentFieldMaxLen[k] = len(v2.Data)
			}
		}
	}
}

func (t *Table) readPrefix() {
	if t.prefixDisable == true {
		return
	}
	diffLength := len(t.prefixTab) - len(t.prefixContent)
	if diffLength < 0 {
		t.prefixTab += strings.Repeat(" ", int(math.Abs(float64(diffLength))))
	} else {
		t.prefixContent += strings.Repeat(" ", diffLength)
	}
}

// print
func (t *Table) Print() {
	if len(t.tab) > 0 {
		t.contentLine = append([][]LineData{t.tab}, t.contentLine...)
	}
	t.initContentMaxLen()
	t.readData()
	t.readPrefix()
	t.printLine()
}
