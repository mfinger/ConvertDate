package DateConverter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDateConverter(t *testing.T) {
	dc, err := NewDateConverter("y4-m2-m2", "m2/d2/y4")

	assert.NotNil(t, dc)
	assert.Nil(t, err)
	assert.NotNil(t, dc.inputRegex)
	assert.Equal(t, "m2/d2/y4", dc.OutputFormat)

}

func TestDateConverter_SetInputFormat(t *testing.T) {
	dc, _ := NewDateConverter("", "m2/d2/y4")

	dc.SetInputFormat("y4-m2-d2")

	assert.Equal(t, `(\d{4})-(\d{2})-(\d{2})`, dc.inputRegex.String())

}

func TestDateConverter_GenerateNewDate(t *testing.T) {
	dc, _ := NewDateConverter("", "m2/d2/y4")

	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	result := dc.GenerateNewDate(dt)

	assert.Equal(t, "12/17/2023", result)
}

func TestDateConverter_ConvertString(t *testing.T) {
	type TestCase struct {
		Name         string
		InputFormat  string
		OutputFormat string
		InputString  string
		Expected     string
	}

	tests := []TestCase{
		{
			Name:         "Single convert",
			InputFormat:  "m2/d2/y4",
			InputString:  "Today is 12/21/2023",
			OutputFormat: "y4-m2-d2",
			Expected:     "Today is 2023-12-21",
		},
		{
			Name:         "Conplex convert",
			InputFormat:  "d2or day of ml in the year y4",
			InputString:  "Today is the 21st day of December in the year 2023",
			OutputFormat: "y4-m2-d2",
			Expected:     "Today is the 2023-12-21",
		},
		{
			Name:         "All output formats",
			InputFormat:  "m2/d2/y4",
			InputString:  "09/01/2023",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "1:01:9:09:September:Sep:st:23:2023:",
		},
		{
			Name:         "Single digits on double digits",
			InputFormat:  "m2/d2/y4",
			InputString:  "9/01/2023",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "9/01/2023",
		},
		{
			Name:         "Double digits on single digits",
			InputFormat:  "m1/d1/y4",
			InputString:  "09/01/2023",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "09/01/2023",
		},
	}

	dc, _ := NewDateConverter("", "")

	for _, test := range tests {
		dc.SetInputFormat(test.InputFormat)
		dc.OutputFormat = test.OutputFormat

		result := dc.ConvertString(test.InputString)

		assert.Equal(t, test.Expected, result, "TestCase '%s' failed.")

	}
	//input := "Today is 12/21/2023, tomorrow is December 22nd, 2023, yesterday was the 19th of December, and 2023-02-31 is not a valid date"

}
