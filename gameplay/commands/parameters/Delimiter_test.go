package parameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelimiterReturnsNotFoundWhenNotMatched(t *testing.T) {
	c := "Some text without the delimiter."
	d := NewDelimiter("=")

	i := d.Find(c)

	assert.Equal(t, -1, i)
}

func TestDelimiterReturnsFirstCharacterIndexWhenMatched(t *testing.T) {
	c := "p Bridget=Hi, Bridget!"
	d := NewDelimiter("=")

	i := d.Find(c)

	assert.Equal(t, 9, i)
}

func TestDelimiterConsumesDelimitingText(t *testing.T) {
	c := "= Hi, Bridget!"
	d := NewDelimiter("=")

	i := d.Find(c)

	assert.Equal(t, 0, i)

	v, remaining := d.Consume(c)

	assert.Equal(t, v.Single(), "=")
	assert.Equal(t, remaining, "Hi, Bridget!")
}
