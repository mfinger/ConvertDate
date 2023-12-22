package DateConverter

import (
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

// List of special regular expressions characters that need to be protected in the matching pattern
var specialRegexCharacters = map[string]interface{}{
	"\\": nil,
	"^":  nil,
	"$":  nil,
	".":  nil,
	"|":  nil,
	"?":  nil,
	"*":  nil,
	"+":  nil,
	"(":  nil,
	")":  nil,
	"[":  nil,
	"]":  nil,
	"{":  nil,
	"}":  nil,
}

// DateConverter Struct for the date conversion being executed.
type DateConverter struct {
	inputRegex       *regexp.Regexp
	inputParseFormat []string
	OutputProcessor  []ConverterFunc
}

// NewDateConverter Constructor
func NewDateConverter(inputFormat string, outputFormat string) (*DateConverter, error) {
	dc := new(DateConverter)

	dc.SetOutputFormat(outputFormat)
	err := dc.SetInputFormat(inputFormat)

	return dc, err
}

// GenerateNewDate For a given time value, convert it to the correct output value based on the configured format
func (dc *DateConverter) GenerateNewDate(t time.Time) string {
	var output string
	for _, fn := range dc.OutputProcessor {
		output += fn(t)
	}
	return output
}

// SetOutputFormat Setter to allow for setting/changing the output format. Also separated for better testing.
func (dc *DateConverter) SetOutputFormat(outputFormat string) {
	// New output
	var buffer string

	dc.OutputProcessor = []ConverterFunc{}

	// Loop through the output format
	for i := 0; i < len(outputFormat); i++ {
		var found bool
		var format FormatConversion

		// Skip this if we are too close to the end (i.e. last char)
		if i < len(outputFormat)-1 {

			// Look to see if these two chars are a valid format
			format, found = FormatMap[outputFormat[i:i+2]]
		}

		// If this was a not valid format specifier then just add this char to the output
		// Otherwise, add the converted value to the output
		if !found {
			buffer += outputFormat[i : i+1]
		} else {
			if len(buffer) != 0 {
				dc.OutputProcessor = append(dc.OutputProcessor, NewConstantConverterFunc(buffer))
				buffer = ""
			}
			dc.OutputProcessor = append(dc.OutputProcessor, format.Convert)
			i++
		}
	}
	if len(buffer) != 0 {
		dc.OutputProcessor = append(dc.OutputProcessor, NewConstantConverterFunc(buffer))
	}
}

// SetInputFormat Setter to allow for setting/changing the input format. Also separated for better testing.
func (dc *DateConverter) SetInputFormat(inputFormat string) error {
	// Dynamic regex string and parsing format
	inputRegexString := ""
	dc.inputParseFormat = []string{}

	// Loop through the input format
	for i := 0; i < len(inputFormat); i++ {
		var found bool
		var format FormatConversion

		// Skip this if we are too close to the end (i.e. last char)
		if i < len(inputFormat)-1 {

			// Look to see if these two chars match or format tag
			format, found = FormatMap[inputFormat[i:i+2]]
		}
		if !found {
			c := inputFormat[i : i+1]
			if _, found := specialRegexCharacters[c]; found {
				inputRegexString += "\\"
			}
			// Doesn't match a format, add this char to the regex string
			inputRegexString += inputFormat[i : i+1]
		} else {

			// Matches, so add the regex to the regex string and the format string to the parsing string (and skip a character)
			inputRegexString += format.RegularExpression
			dc.inputParseFormat = append(dc.inputParseFormat, format.FormatString)
			i++
		}
	}

	// Try to compile the built regex
	var err error
	dc.inputRegex, err = regexp.Compile(inputRegexString)

	return err
}

// ConvertString Take a given string, convert the matching dates to the new format and return the new string
func (dc *DateConverter) ConvertString(inputString string) string {

	// Find all the indexes of the group matches in the string
	matches := dc.inputRegex.FindAllStringSubmatchIndex(inputString, -1)

	var output string
	var start int

	for _, match := range matches {

		// Copy in the contents up until the start of the last second of this match, then update the start
		output += inputString[start:match[0]]
		start = match[1]

		// Build the time string and the parse format for Go's time.Parse function
		// Note: Match offsets returned by the regex call is pairs if start/end in one slice
		//       i.e. match[2] = start, match[3] = end for the first submatch
		var parseFormat, parseString string
		for j, p := 2, 0; j < len(match); j, p = j+2, p+1 {

			// If there is no input parse format  (i.e. for ordinal) then we don't include it in the parsing values.
			if len(dc.inputParseFormat[p]) != 0 {
				parseFormat += dc.inputParseFormat[p] + ":"
				parseString += inputString[match[j]:match[j+1]] + ":"
			}
		}

		// See if the date parses, if it doesn't issue a warning and skip it and include the original data.
		// Otherwise, do the replacement in the string.
		if t, err := time.Parse(parseFormat, parseString); err != nil {
			log.Printf("[WARN] Could not parse date '%v', leaving alone: %v\n", inputString[match[0]:match[1]], err)
			output += inputString[match[0]:match[1]]
		} else {
			output += dc.GenerateNewDate(t)

		}
	}
	output += inputString[start:]
	return output

}

// ConvertFile Convert dates for a given file and write out a new file
func (dc *DateConverter) ConvertFile(inputFilename string, outputFilename string) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	if buff, err := io.ReadAll(inputFile); err != nil {
		return err
	} else {
		newContent := dc.ConvertString(string(buff))

		f, err := os.OpenFile(outputFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer f.Close()
		if err != nil {
			return err
		}

		if _, err := f.WriteString(newContent); err != nil {
			return err
		}
		return nil
	}
}
