package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"task_6/validator"
	"time"

	"github.com/brandenc40/romannumeral"
)

func main() {
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
	time.Sleep(time.Second)

	fmt.Println("Generating new result file...")
	time.Sleep(time.Second)
	wFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer wFile.Close()

	lines := 0
	for scanner.Scan() {
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
