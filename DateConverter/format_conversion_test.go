package DateConverter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatConversion_Convert_With_Format(t *testing.T) {
	fc := FormatConversion{FormatString: "2006"}

	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	result := fc.Convert(dt)

	assert.Equal(t, "2023", result, "Conversion should produce 2023")

}

func TestFormatConversion_Convert_With_UnknownFormat(t *testing.T) {
	fc := FormatConversion{FormatString: "junk"}

	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	result := fc.Convert(dt)

	assert.Equal(t, "junk", result, "Conversion should produce junk")

}

func TestFormatConversion_Convert_With_Custom(t *testing.T) {
	fc := FormatConversion{
		CustomConverter: func(t2 time.Time) string {
			return fmt.Sprintf("-%d-", t2.Year())
		}}

	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	result := fc.Convert(dt)

	assert.Equal(t, "-2023-", result, "Conversion should produce -2023-")

}

func TestNewConstantConverterFunc(t *testing.T) {
	fn := NewConstantConverterFunc("test")
	dt := time.Date(2023, 12, 17, 11, 0, 0, 0, time.UTC)

	assert.Equal(t, "test", fn(dt))
}
