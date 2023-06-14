package structs

import (
	"fmt"
	"github.com/dslipak/pdf"
	"github.com/xuri/excelize/v2"
	"strings"
)

type Dir struct {
	Path string
	Name string
}

func (d *Dir) SetData(path, name string) {
	d.Path = path
	d.Name = name
}

type Data struct {
	Folder []Dir
	Files  []Dir
}

var dataTypes = []string{"xlsx", "xls", "doc", "docx", "pdf"}

func (d *Data) ProcessData(f *excelize.File) {
	index := 1
	for _, v := range d.Folder {
		f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), fmt.Sprintf("A%d", index), v.Path)
		index++
		for _, n := range d.Files {
			if n.Path == v.Path {
				splitedName := strings.Split(n.Name, ".")
				f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), fmt.Sprintf("A%d", index), n.Name)
				counter := 0
				if splitedName[len(splitedName)-1] == "pdf" {
					temp, err := pdf.Open(n.Path + "/" + n.Name)
					if err != nil {
						fmt.Println(err)
					}
					counter = temp.NumPage()
					fmt.Println(temp.NumPage())
				} else if splitedName[len(splitedName)-1] == "xlsx" || splitedName[len(splitedName)-1] == "xls" {
					f, err := excelize.OpenFile(n.Path + "/" + n.Name)
					if err != nil {
						fmt.Println(err)
						return
					}
					counter = len(f.GetSheetList())
					fmt.Println(len(f.GetSheetList()))
				} else {
					counter = 1
				}
				f.SetCellValue(f.GetSheetName(f.GetActiveSheetIndex()), fmt.Sprintf("B%d", index), counter)

				index++
			}
		}
	}
}
