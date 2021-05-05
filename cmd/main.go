package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/timopattikawa/kamoro/grader"
)

func main() {
	var excelStudentsGrade []grader.ExcelGrade

	builder := grader.BuilderMachine{}
	builder.SetUpFile()

	grader := grader.GraderMachine{Builder: builder}

	for _, student := range builder.Students {
		log.Println(student.Nim)
		for _, submission := range student.Submissions {
			studentExcelGrade := submission.Grade(&grader)
			studentExcelGrade.NIM = student.Nim
			excelStudentsGrade = append(excelStudentsGrade, studentExcelGrade)
		}
	}

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet Nilai"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "NIM")
	xlsx.SetCellValue(sheet1Name, "B1", "Problem")
	xlsx.SetCellValue(sheet1Name, "C1", "Language")
	xlsx.SetCellValue(sheet1Name, "D1", "Status")
	xlsx.SetCellValue(sheet1Name, "E1", "Time")

	err := xlsx.AutoFilter(sheet1Name, "A1", "E1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for idx, each := range excelStudentsGrade {

		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", idx+2), each.NIM)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", idx+2), each.Problem)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", idx+2), each.Language)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", idx+2), each.Status)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", idx+2), each.Time.Seconds())
	}

	err = xlsx.SaveAs(builder.Path + "/result.xlsx")
	if err != nil {
		log.Println(err)
	}
}
