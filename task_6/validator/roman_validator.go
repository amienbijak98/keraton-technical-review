package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func RomanValidator(input string) (bool, error) {
	//INITIATING ROMAN VALUES AS MAP
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	//VALIDATE ONLY ALPHA ASCII ALLOWED
	validate := validator.New()
	err := validate.Var(input, "alpha")

	if err != nil {
		return false, errors.New("only alpha ASCII input allowed")
	}

	//Counting consecutive character
	count := 1
	for i := 0; i < len(input); i++ {
		//VALIDATE ROMAN NUMERAL INPUT TYPE
		if _, exists := romanValues[input[i]]; !exists {
			return false, errors.New("input alphabet is not roman numeral")
		}

		//VALIDATE ROMAN NUMERAL WRITING RULES
		// If next character is larger, it should be a valid subtraction
		if i+1 < len(input) && romanValues[input[i]] < romanValues[input[i+1]] {
			// Check for valid subtraction
			if !(input[i] == 'I' && (input[i+1] == 'V' || input[i+1] == 'X') ||
				input[i] == 'X' && (input[i+1] == 'L' || input[i+1] == 'C') ||
				input[i] == 'C' && (input[i+1] == 'D' || input[i+1] == 'M')) {
				return false, errors.New("does not comply with the rules for writing Roman numerals")
			} else if (i-1 >= 0 && input[i] == input[i-1]) || (i-2 >= 0 && input[i] == input[i-2]) {
				//Check multiple lesser character before current character
				return false, errors.New("does not comply with the rules for writing Roman numerals")
			}

			i++
			count = 1 //reset count consecutive character
		} else {
			// Check if the same character repeats more than 3 times consecutively
			if i > 0 && input[i] == input[i-1] {
				count++
				if count > 3 || (input[i] == 'V' || input[i] == 'L' || input[i] == 'D') && count > 1 {
					return false, errors.New("does not comply with the rules for writing Roman numerals")
				}
			} else {
				count = 1
			}
		}
	}

	return true, nil
}
