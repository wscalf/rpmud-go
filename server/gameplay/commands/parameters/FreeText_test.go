package parameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreeTextReturnsNotFoundForEmptyInput(t *testing.T) {
	c := ""
	f := NewFreeText("text", false)

	i := f.Find(c)
	assert.Equal(t, -1, i)
}

func TestFreeTextConsumsAndTrimsRemainingInput(t *testing.T) {
	c := "Ohai there! All of this should be consumed. "
	f := NewFreeText("text", false)

	i := f.Find(c)
	assert.Equal(t, 0, i)

	v, remaining := f.Consume(c)

	assert.Equal(t, v.Single(), "Ohai there! All of this should be consumed.")
	assert.Equal(t, remaining, "")
}
