# KAMORO
Kamoro is command-line application for grading C++, Python 3, and Golang code base on given input and expected ouput 
___
## Requirements
What you need to install in your machine:
 - `go 1.15`
 - `g++`
 - `python 3`

You need 3 folders inside the directory you want to grade : 
 - `input` : Contains input file
 - `output` : Contains Expected Output file
 - `submissions` : Contains all submissions from students
### A typical directory layout

    $ pwd
    $ /home/user/example

    example #The Directory
    |
    ├── input  
    │   ├── a.in
    │   ├── b.in          
    ├── output
    │   ├── a.out
    │   ├── b.out                
    ├── submissions                  
    │   ├── 123456-a.cpp
    │   ├── 123456-b.py
    │   ├── 123457-a.go
    |   └── ...          
    └── ...

> Use lowercase for every folder name
___
## Notes
 - `The submission name must follow this: {studentID}-{typeproblem}.cpp`
 - `input and output file name inside input and output folder must follow this: a.in / a.out {typeproblem.extension (in/out) }`

___
## How to Use
 1. Clone from my GitHub repository
 2. Go ahead to the folder that you have cloned
 3. Just type `$ make init`
 4. Just type `$ make grade`
 5. Input Your path directory for the grader 
    - Use `$ pwd` to know the path
    - example input: `$ Input Path: /home/user/example`
 6. The results of the assessment will be entered into the xlsx file in the input path
> This app compitible to run on Unix OS
