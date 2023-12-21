package DateConverter

import (
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

// DateConverter Struct for the date conversion being executed.
type DateConverter struct {
	inputRegex       *regexp.Regexp
	inputParseFormat []string
	OutputFormat     string
}

// NewDateConverter Constructor
func NewDateConverter(inputFormat string, outputFormat string) (*DateConverter, error) {
	dc := new(DateConverter)

	dc.OutputFormat = outputFormat

	err := dc.SetInputFormat(inputFormat)

	return dc, err
}

// GenerateNewDate For a given time value, convert it to the correct output value based on the configured format
func (dc *DateConverter) GenerateNewDate(t time.Time) string {
	// New output
	var output string

	// Loop through the output format
	for i := 0; i < len(dc.OutputFormat); i++ {
		var found bool
		var format FormatConversion

		// Skip this if we are too close to the end (i.e. last char)
		if i < len(dc.OutputFormat)-1 {

			// Look to see if these two chars are a valid format
			format, found = FormatMap[dc.OutputFormat[i:i+2]]
		}

		// If this was a not valid format specifier then just add this char to the output
		// Otherwise, add the converted value to the output
		if !found {
			output += dc.OutputFormat[i : i+1]
		} else {
			output += format.Convert(t)
			i++
		}
	}
	return output
}

// SetInputFormat Setter to allow for setting/changing the input format. Also separated for better testing.
func (dc *DateConverter) SetInputFormat(inputFormat string) error {
	// Dynamic regex string and parsing format
	inputRegexString := ""
	dc.inputParseFormat = []string{}

	// Loop through the input format
	for i := 0; i < len(inputFormat); i++ {

		// Look to see if these two chars match or format tag
		format, found := FormatMap[inputFormat[i:i+2]]
		if !found {

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

	// Work backwards through the list so the offsets are still correct
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]

		// Offset where the overall match starts and ends (for substringing)
		start := match[0]
		end := match[1]

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

		// See if the date parses, if it doesn't issue a warning and skip it.
		// Otherwise, do the replacement in the string.
		if t, err := time.Parse(parseFormat, parseString); err != nil {
			log.Printf("[WARN] Could not parse date '%v', leaving alone: %v\n", inputString[start:end], err)
		} else {
			inputString = inputString[:start] + dc.GenerateNewDate(t) + inputString[end:]
		}
	}

	return inputString

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
