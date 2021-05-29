package ezexcelize

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type File struct {
	excelize.File
}

func NewFile() *File {
	return &File{*excelize.NewFile()}
}

func (f *File) WriteRowAsString(sheet string, row int, col int, cells ...string) error {
	for offset, cell := range cells {
		axis, err := excelize.CoordinatesToCellName(col+offset, row)
		if err != nil {
			return err
		}
		if err2 := f.SetCellStr(sheet, axis, cell); err2 != nil {
			return err2
		}
	}
	return nil
}
