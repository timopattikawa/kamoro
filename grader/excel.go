package grader

import (
	"time"
)

type ExcelGrade struct {
	NIM      string
	Problem  string
	Language string
	Status   string
	Time     time.Duration
}
