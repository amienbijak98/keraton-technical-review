# Roman Numeric Converter

Coded to complete the Keraton technical review task number 6. By Bijak Amien Muttaqie
_Estimated work time 3 hours_

## How To Use The App

- Install latest golang version.
- Place your roman numeric file next to task_6.go (the main app). With the name "roman.txt"
- Open your terminal and write command `go run task_6.go`
- It will work and resulting output file named "result.txt"

## Tools or Go Package I used

- Standart Library :
  - os : for reading and writing the file
  - bufio : for scanning the file
- External Library :
  - github.com/brandenc40/romannumeral : for converting the roman to integer value.
  - github.com/go-playground/validator/v10 : for validate input type
  - github.com/stretchr/testify/assert : for assertion test
