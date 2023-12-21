package DateConverter

import "time"

// ConverterFunc Type for a custom converter function
type ConverterFunc func(time2 time.Time) string

// FormatConversion Struct to hold all the configure format converter specifications
type FormatConversion struct {
	RegularExpression string
	FormatString      string
	Description       string
	CustomConverter   ConverterFunc
}

// Convert Function to convert a time value to the needed string
func (fc FormatConversion) Convert(t time.Time) string {

	// Use the custom converter if one is configured, otherwise use time.Format to do the work.
	if fc.CustomConverter == nil {
		return t.Format(fc.FormatString)
	} else {
		return fc.CustomConverter(t)
	}
}
