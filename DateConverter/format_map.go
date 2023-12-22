package DateConverter

import "time"

// Function to get the ordinal for day from the passed in time
func ordinalConverter(t time.Time) string {
	dayDigit := t.Day() % 10
	switch dayDigit {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// FormatMap Map of all the format conversion specifications.
var FormatMap = map[string]FormatConversion{
	"ms": {
		RegularExpression: `((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)`,
		FormatString:      "Jan",
		Description:       "Long month name (i.e January)",
	},
	"ml": {
		RegularExpression: `((?i)january|february|march|april|may|june|july|august|september|october|november|december)`,
		FormatString:      "January",
		Description:       "Short month name (i.e Jan)",
	},
	"m1": {
		RegularExpression: `(\d{2}|[1-9])`,
		FormatString:      "1",
		Description:       "Non zero padded month number (i.e. 1)",
	},
	"m2": {
		RegularExpression: `(\d{2})`,
		FormatString:      "01",
		Description:       "Zero padded month number (i.e. 01)",
	},
	"d1": {
		RegularExpression: `(\d{2}|[1-9])`,
		FormatString:      "2",
		Description:       "Non zero padded month number (i.e. 2)",
	},
	"d2": {
		RegularExpression: `(\d{2})`,
		FormatString:      "02",
		Description:       "Zero padded day number (i.e. 02)",
	},
	"or": {
		RegularExpression: `(st|nd|rd|th)`,
		FormatString:      "",
		CustomConverter:   ordinalConverter,
		Description:       "Required ordinal string (i.e. st, nd, rd, th)",
	},
	"oo": {
		RegularExpression: `(st|nd|rd|th)?`,
		FormatString:      "",
		CustomConverter:   ordinalConverter,
		Description:       "Optional rdinal string (i.e. st, nd, rd, th)",
	},
	"y4": {
		RegularExpression: `(\d{4})`,
		FormatString:      "2006",
		Description:       "Year number including century (i.e. 2006)",
	},
	"y2": {
		RegularExpression: `(\d{2})`,
		FormatString:      "06",
		Description:       "Year number without century (i.e. 06)",
	},
}
