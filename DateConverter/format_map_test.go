package DateConverter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatMap_ordinalConverter(t *testing.T) {
	assert.Equal(t, "th", ordinalConverter(time.Date(2023, 12, 20, 11, 0, 0, 0, time.UTC)))
	assert.Equal(t, "st", ordinalConverter(time.Date(2023, 12, 21, 11, 0, 0, 0, time.UTC)))
	assert.Equal(t, "nd", ordinalConverter(time.Date(2023, 12, 22, 11, 0, 0, 0, time.UTC)))
	assert.Equal(t, "rd", ordinalConverter(time.Date(2023, 12, 23, 11, 0, 0, 0, time.UTC)))
	assert.Equal(t, "th", ordinalConverter(time.Date(2023, 12, 24, 11, 0, 0, 0, time.UTC)))
	assert.Equal(t, "th", ordinalConverter(time.Date(2023, 12, 25, 11, 0, 0, 0, time.UTC)))
}
