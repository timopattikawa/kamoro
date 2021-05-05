package grader

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type BuilderMachine struct {
	Students            []Student
	InputFiles          map[string]string
	ExpectedOutputFiles map[string]string
	Path                string
}

func (b BuilderMachine) findStudent(nim string) int {
	for index, student := range b.Students {
		if nim == student.Nim {
			return index
		}
	}
	return -1
}

func (b BuilderMachine) readFolderInPath(inputPath string) (map[string]string, error) {
	folders, err := ioutil.ReadDir(inputPath)
	if err != nil {
		return nil, err
	}

	mapFiles := make(map[string]string)
	for _, folder := range folders {
		if folder.IsDir() {
			if folder.Name() == "submissions" || folder.Name() == "input" || folder.Name() == "output" {
				mapFiles[folder.Name()] = filepath.Join(inputPath, folder.Name())
			}
		}
	}
	return mapFiles, nil
}

func (b *BuilderMachine) readFileInPath(inputPath string) (map[string]string, error) {
	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	mapFiles := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			mapFiles[file.Name()] = filepath.Join(inputPath, file.Name())
		}
	}
	return mapFiles, nil
}

func (b *BuilderMachine) searchStudentsSubmissions(dir string) error {
	var submissions map[string]string
	submissions, err := b.readFileInPath(dir)

	if err != nil {
		return err
	}

	for fname, fpath := range submissions {
		nim := strings.Split(fname, "-")[0]

		typeProblem := (strings.Split(fname, "-")[1])
		typeProblem = strings.Split(typeProblem, ".")[0]

		extension := strings.Split(fname, ".")[1]

		var submission Submission
		if extension == "py" {
			submission = &SubmissionPython{Path: fpath, TypeProblem: typeProblem}
		} else if extension == "cpp" {
			submission = &SubmissionCpp{Path: fpath, TypeProblem: typeProblem}
		} else if extension == "go" {
			submission = &SubmissionGo{Path: fpath, TypeProblem: typeProblem}
		}

		findStudent := b.findStudent(nim)
		if findStudent != -1 {
			b.Students[findStudent].Submissions = append(b.Students[findStudent].Submissions, submission)
		} else {
			var studentTmp Student
			studentTmp.Nim = nim
			studentTmp.Submissions = append(studentTmp.Submissions, submission)
			b.Students = append(b.Students, studentTmp)
		}
	}

	return nil
}

func (b *BuilderMachine) SetUpFile() {
	var sourcePath string
	fmt.Print("Input Path: ")
	fmt.Scanln(&sourcePath)
	b.Path = sourcePath
	folders, err := b.readFolderInPath(sourcePath)
	if err != nil {
		log.Println("++No such File or Directory++")
		return
	}

	err = b.searchStudentsSubmissions(folders["submissions"])
	if err != nil {
		log.Println("Folder Submissons not found")
		return
	}

	inputsFiles, err := b.readFileInPath(folders["input"])
	if err != nil {
		log.Println("Folder Input not found")
		return
	}
	b.InputFiles = inputsFiles

	outputFiles, err := b.readFileInPath(folders["output"])
	if err != nil {
		log.Println("Folder Output not found")
		return
	}
	b.ExpectedOutputFiles = outputFiles

}
