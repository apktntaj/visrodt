package excel

import (
	"fmt"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func CreateSheet() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set value of a cell
	f.SetCellValue("Sheet2", "A2", "Hello World")
	f.SetCellValue("Sheet1", "B2", "100")
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ReadFile(name string) {
	if isNotXlsx(name) {
		fmt.Println("Ubah dulu ya ke xlsx")
		return
	}

	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// sheets := f.GetSheetName(0)

	// Get value from cell by given worksheet name and cell reference.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)

	for _, sheet := range f.GetSheetList() {

		fmt.Println("=========================")
		fmt.Println(sheet)
		fmt.Println("=========================")
		// Get all the rows in the Sheet1.
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}
}

func isNotXlsx(name string) bool {

	return filepath.Ext(name) != ".xlsx"
}
