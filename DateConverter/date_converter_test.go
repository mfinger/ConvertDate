package DateConverter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDateConverter(t *testing.T) {
	dc, err := NewDateConverter("y4-m2-m2", "m2/d2/y4")
	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	assert.NotNil(t, dc)
	assert.Nil(t, err)
	assert.NotNil(t, dc.inputRegex)
	assert.Equal(t, 5, len(dc.OutputProcessor))
	assert.Equal(t, "12", dc.OutputProcessor[0](dt))
	assert.Equal(t, "/", dc.OutputProcessor[1](dt))
	assert.Equal(t, "17", dc.OutputProcessor[2](dt))
	assert.Equal(t, "/", dc.OutputProcessor[3](dt))
	assert.Equal(t, "2023", dc.OutputProcessor[4](dt))

}

func TestDateConverter_SetInputFormat(t *testing.T) {
	dc, _ := NewDateConverter("", "m2/d2/y4")

	dc.SetInputFormat("y4-m2-d2")

	assert.Equal(t, `(\d{4})-(\d{2})-(\d{2})`, dc.inputRegex.String())

}

func TestDateConverter_SetInputFormatTrailing(t *testing.T) {
	dc, _ := NewDateConverter("", "m2/d2/y4")

	dc.SetInputFormat("y4-m2-d2-")

	assert.Equal(t, `(\d{4})-(\d{2})-(\d{2})-`, dc.inputRegex.String())

}

func TestDateConverter_SetInputFormatWithSpecialRegex(t *testing.T) {
	dc, _ := NewDateConverter("", "m2/d2/y4")

	dc.SetInputFormat("y4.m2$d2")

	assert.Equal(t, `(\d{4})\.(\d{2})\$(\d{2})`, dc.inputRegex.String())

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
			Expected:     "1:01:9:09:September:Sep:st:23:2023:",
		},
		{
			Name:         "Required ordinal with ordinal",
			InputFormat:  "y4 m2 d1or",
			InputString:  "2023 09 1st",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "1:01:9:09:September:Sep:st:23:2023:",
		},
		{
			Name:         "Required ordinal without ordinal",
			InputFormat:  "y4 m2 d1or",
			InputString:  "2023 09 1",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "2023 09 1",
		},
		{
			Name:         "Optional ordinal with ordinal",
			InputFormat:  "y4 m2 d1oo",
			InputString:  "2023 09 1st",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "1:01:9:09:September:Sep:st:23:2023:",
		},
		{
			Name:         "Optional ordinal without ordinal",
			InputFormat:  "y4 m2 d1oo",
			InputString:  "2023 09 1",
			OutputFormat: "d1:d2:m1:m2:ml:ms:or:y2:y4:",
			Expected:     "1:01:9:09:September:Sep:st:23:2023:",
		},
	}

	dc, _ := NewDateConverter("", "")

	for _, test := range tests {
		dc.SetInputFormat(test.InputFormat)
		dc.SetOutputFormat(test.OutputFormat)

		result := dc.ConvertString(test.InputString)

		assert.Equal(t, test.Expected, result, "TestCase '%s' failed.")

	}

}
