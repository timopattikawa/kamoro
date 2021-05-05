package grader

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

type Submission interface {
	Grade(grader Grader) ExcelGrade
}

// Python
type SubmissionPython struct {
	Path        string
	TypeProblem string
}

func (s *SubmissionPython) Grade(grader Grader) ExcelGrade {
	execelGrade := grader.GradeForPy(s)
	return execelGrade
}

// C++/C
type SubmissionCpp struct {
	Path        string
	CompilePath string
	TypeProblem string
}

func (s *SubmissionCpp) Grade(grader Grader) ExcelGrade {
	excelGrade := grader.GradeForCpp(s)
	return excelGrade
}

func (s *SubmissionCpp) Compile() error {
	outputPath := path.Join(fmt.Sprintf("%s", strings.Split(s.Path, ".")[0]))
	args := strings.Split(fmt.Sprintf("%s -o %s", s.Path, outputPath), " ")
	cmd := exec.Command("g++", args...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fileName := strings.Split(s.Path, "/")
		fmt.Printf("\tFailed to Compile Problem : %s. Check Code\n", fileName[len(fileName)-1])
		return err
	}

	s.CompilePath = outputPath
	return nil
}

// Go
type SubmissionGo struct {
	Path        string
	CompilePath string
	TypeProblem string
}

func (s *SubmissionGo) Grade(grader Grader) ExcelGrade {
	excelGrade := grader.GradeForGo(s)
	return excelGrade
}

func (s *SubmissionGo) Compile() error {
	outputPath := path.Join(fmt.Sprintf("%s", strings.Split(s.Path, ".")[0]))
	args := strings.Split(fmt.Sprintf("build -o %s %s", outputPath, s.Path), " ")
	cmd := exec.Command("go", args...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fileName := strings.Split(s.Path, "/")
		log.Printf("\tFailed to Compile : %s. Check Code\n", fileName[len(fileName)-1])
		return err
	}

	s.CompilePath = outputPath
	return nil
}

type Grader interface {
	GradeForCpp(submission *SubmissionCpp) ExcelGrade
	GradeForPy(submission *SubmissionPython) ExcelGrade
	GradeForGo(submission *SubmissionGo) ExcelGrade
}

type GraderMachine struct {
	Builder BuilderMachine
}

func (g *GraderMachine) GradeForPy(submission *SubmissionPython) ExcelGrade {
	// todo algo grade python file
	var status string
	start := time.Now()
	cmd := exec.Command("python3", submission.Path)
	inputFile := g.Builder.InputFiles[submission.TypeProblem+".in"]
	f, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	cmd.Stdin = strings.NewReader(string(f))
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		status = "Error Run Program"
		log.Println(err.Error())
	}

	result := strings.TrimSpace(string(outb.Bytes()))

	fileOutput, err := ioutil.ReadFile(g.Builder.ExpectedOutputFiles[submission.TypeProblem+".out"])
	if err != nil {
		panic(err)
	}

	if string(fileOutput) != result && len(status) == 0 {
		status = "FAIL"
	} else if len(status) == 0 && string(fileOutput) == result {
		status = "PASS"
	}

	end := time.Since(start)
	return ExcelGrade{
		Problem:  strings.ToUpper(submission.TypeProblem),
		Language: "Python",
		Status:   status,
		Time:     end,
	}
}

func (g *GraderMachine) GradeForCpp(submission *SubmissionCpp) ExcelGrade {
	// todo algo grade CPP file
	var status string

	err := submission.Compile()
	start := time.Now()

	if err != nil {
		status = "Compile Error"
	} else {

		cmd := exec.Command(submission.CompilePath)
		f, err := ioutil.ReadFile(g.Builder.InputFiles[submission.TypeProblem+".in"])
		if err != nil {
			log.Println("Not found input file")
			panic(err)
		}

		cmd.Stdin = strings.NewReader(string(f))
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err = cmd.Run()
		if err != nil {
			status = "Error for run Code"
			log.Println("Not found file for compile")
			log.Println(err.Error())
		}

		result := strings.TrimSpace(string(outb.Bytes()))

		fileOutput, err := ioutil.ReadFile(g.Builder.ExpectedOutputFiles[submission.TypeProblem+".out"])
		if err != nil {
			log.Println("Not found output file")
			panic(err)
		}

		if string(fileOutput) != result && len(status) == 0 {
			status = "FAIL"
		} else if len(status) == 0 && string(fileOutput) == result {
			status = "PASS"
		}

		os.Remove(submission.CompilePath)
	}
	end := time.Since(start)

	return ExcelGrade{
		Problem:  strings.ToUpper(submission.TypeProblem),
		Language: "C++/C",
		Status:   status,
		Time:     end,
	}
}

func (g *GraderMachine) GradeForGo(submission *SubmissionGo) ExcelGrade {
	// todo algo grade go file
	var status string

	err := submission.Compile()
	start := time.Now()

	if err != nil {
		status = "Compile Error"
	} else {

		cmd := exec.Command(submission.CompilePath)
		f, err := ioutil.ReadFile(g.Builder.InputFiles[submission.TypeProblem+".in"])
		if err != nil {
			log.Println("Not found input file")
			panic(err)
		}

		cmd.Stdin = strings.NewReader(string(f))
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err = cmd.Run()
		if err != nil {
			status = "Error for run Code"
			log.Println("Not found file for compile")
			log.Println(err.Error())
		}

		result := strings.TrimSpace(string(outb.Bytes()))

		fileOutput, err := ioutil.ReadFile(g.Builder.ExpectedOutputFiles[submission.TypeProblem+".out"])
		if err != nil {
			log.Println("Not found output file")
			panic(err)
		}

		if string(fileOutput) != result && len(status) == 0 {
			status = "FAIL"
		} else if len(status) == 0 && string(fileOutput) == result {
			status = "PASS"
		}

		os.Remove(submission.CompilePath)
	}
	end := time.Since(start)

	return ExcelGrade{
		Problem:  strings.ToUpper(submission.TypeProblem),
		Language: "Golang",
		Status:   status,
		Time:     end,
	}
}
