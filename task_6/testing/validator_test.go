package testing

import (
	"task_6/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorRoman(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected bool
	}{
		//False Input
		{name: "False Input Z$!", request: "Z$!", expected: false},
		{name: "False Input C M", request: "C M", expected: false},
		{name: "False Input 123", request: "123", expected: false},
		//Correct Roman writings rules test
		{name: "Correct Roman MMIX", request: "MMIX", expected: true},
		{name: "Correct Roman MMCMLXXXIV", request: "MMCMLXXXIV", expected: true},
		{name: "Correct Roman MMMCMLXXXVII", request: "MMMCMLXXXVII", expected: true},
		{name: "Correct Roman MCXXXIII", request: "MCXXXIII", expected: true},
		//False Roman writings rules test
		{name: "False Roman VXIIDDLL", request: "VXIIDDLL", expected: false},
		{name: "False Roman IIXVXLCM", request: "IIXVXLCM", expected: false},
		{name: "False Roman CCMXXLLV", request: "CCMXXLLV", expected: false},
		{name: "False Roman LLVVDDMM", request: "LLVVDDMM", expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := validator.RomanValidator(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
