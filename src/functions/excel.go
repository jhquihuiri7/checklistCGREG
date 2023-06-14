package functions

import (
	"checklistCGREG/src/structs"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(data structs.Data) string {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	data.ProcessData(f)
	tempFile := "temp.xlsx"
	if err := f.SaveAs(tempFile); err != nil {
		return err.Error()
	}
	//defer os.Remove(tempFile)
	return tempFile
}
