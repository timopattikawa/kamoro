package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type submission struct {
	TypeProblem string
	Path        string
	Extension   string
}

func (s *submission) setFillAttr() {
	handleSplitFile := strings.Split(s.Path, "-")
	s.TypeProblem = handleSplitFile[1]
	handleSplitExtension := strings.Split(s.TypeProblem, ".")
	s.Extension = handleSplitExtension[1]
}

func (s submission) getType() string {
	return strings.Split(s.TypeProblem, ".")[0]
}

func compile(inputPath string) (outputPath string, err error) {
	outputPath = path.Join(fmt.Sprintf("%s", strings.Split(inputPath, ".")[0]))
	args := strings.Split(fmt.Sprintf("%s -o %s", inputPath, outputPath), " ")
	cmd := exec.Command("g++", args...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fileName := strings.Split(inputPath, "/")
		fmt.Printf("\tFailed to Compile Problem : %s. Check Code\n",
			fileName[len(fileName)-1])
	}

	return
}

func run(path, input string) string {
	cmd := exec.Command(path)
	cmd.Stdin = strings.NewReader(input)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal("Fail To Run")
	}

	return strings.TrimSpace(string(outb.Bytes()))
}

func grader(resultCompilePath string, input, expectedOutput []string) string {

	in := strings.Join(input, " ")
	outRun := run(resultCompilePath, in)

	os.Remove(resultCompilePath)

	expectedOut := strings.Join(expectedOutput, " ")
	if outRun != expectedOut {
		return "FAIL"
	}
	return "PASS"
}

func readFolderInPath(inputPath string) (map[string]string, error) {
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

func searchStudentsSubmissions(dir string) (students map[string][]submission, err error) {
	var submissions map[string]string
	submissions, err = readFileInPath(dir)
	if err != nil {
		return nil, err
	}

	students = make(map[string][]submission)
	for fname, fpath := range submissions {
		studentID := strings.Split(fname, "-")[0]
		sub := submission{Path: fpath}
		sub.setFillAttr()
		students[studentID] = append(students[studentID], sub)
	}

	return
}

func readFileInPath(inputPath string) (map[string]string, error) {
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

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := make([]string, 0)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF {
			break
		}

		lines = append(lines, strings.TrimSpace(line))
	}

	return lines, err
}

func main() {
	var sourcePath string
	fmt.Print("Input Path: ")
	fmt.Scanln(&sourcePath)

	folders, err := readFolderInPath(sourcePath)
	if err != nil {
		fmt.Println("++No such File or Directory++")
		return
	}

	studentSubmissons, err := searchStudentsSubmissions(folders["submissions"])
	if err != nil {
		fmt.Print("Folder Submissons not found")
		return
	}

	inputs, err := readFileInPath(folders["input"])
	if err != nil {
		fmt.Print("Folder Input not found")
		return
	}

	inputFiles := make(map[string][]string)
	for inputFile := range inputs {
		mappingInputs, err := readFile(inputs[inputFile])
		if err != nil {
			log.Fatal(err)
		}
		inputFile = strings.Split(inputFile, ".")[0]
		inputFiles[inputFile] = mappingInputs
	}

	outputs, err := readFileInPath(folders["output"])
	if err != nil {
		fmt.Print("Folder Output not found")
		return
	}

	outputFiles := make(map[string][]string)
	for outputFile := range outputs {
		mappingInputs, err := readFile(outputs[outputFile])
		if err != nil {
			log.Fatal(err)
		}
		outputFile = strings.Split(outputFile, ".")[0]
		outputFiles[outputFile] = mappingInputs
	}

	if len(inputFiles) != len(outputFiles) {
		log.Fatal("Error can't match input and output files")
	}

	for student := range studentSubmissons {
		fmt.Printf("Grade: %s\n", student)
		studentSub := studentSubmissons[student]
		for _, sub := range studentSub {
			typeProb := sub.getType()
			compilePath, err := compile(sub.Path)
			if err != nil {
				continue
			}

			inByType := inputFiles[typeProb]
			outByType := outputFiles[typeProb]
			result := grader(compilePath, inByType, outByType)
			fmt.Printf("\tProblem Type %s: %s\n", typeProb, result)
		}
	}

}
