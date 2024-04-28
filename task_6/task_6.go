package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"task_6/validator"

	"github.com/brandenc40/romannumeral"
)

func main() {
	// Read File with certain filename "roman.txt"
	file := ReadFile("roman.txt")
	defer file.Close()
	result := ScanRomanFile(file, "result.txt")

	fmt.Printf(`DONE writing new file with %d lines.`, result)

}

func ReadFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func ScanRomanFile(file *os.File, outputFile string) int {
	fmt.Println("File scanned and converting...")
	scanner := bufio.NewScanner(file)

	fmt.Println("Generating new result file...")
	wFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer wFile.Close()

	lines := 0
	for scanner.Scan() {
		//validate first using function RomanValidator I've made
		//because romannumeral by brandenc40 doesn't have any solid validator
		if realRoman, err := validator.RomanValidator(scanner.Text()); realRoman {
			int_number, _ := romannumeral.StringToInt(scanner.Text())
			WriteFileString(wFile, fmt.Sprintln(scanner.Text(), "=", int_number))
			lines++
		} else {
			WriteFileString(wFile, fmt.Sprintln(scanner.Text(), "=", err.Error()))
			lines++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func WriteFileString(outputFile *os.File, data string) {
	_, err2 := outputFile.WriteString(data)
	if err2 != nil {
		log.Fatal(err2)
	}
}
