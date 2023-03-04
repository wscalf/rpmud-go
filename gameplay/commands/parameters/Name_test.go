package parameters

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNameFindsSingleWordNameAtEndOfInput(t *testing.T) {
	c := "value"

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, i, 0)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), c)
	assert.Equal(t, r, "")
}

func TestNameFindsSingleWordNameWithRemainingInput(t *testing.T) {
	c := "value more"

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, i, 0)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "value")
	assert.Equal(t, r, "more")
}

func TestNameFindsQuotedNameAtEndOfInput(t *testing.T) {
	c := `"quoted name"`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, i, 0)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "quoted name")
	assert.Equal(t, r, "")
}

func TestNameFindsQuotedNameWithRemainingInput(t *testing.T) {
	c := `"quoted name" more`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, i, 0)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "quoted name")
	assert.Equal(t, r, "more")
}

func TestNameReturnsNotFoundForEmptyInput(t *testing.T) {
	n := NewName("param", false)

	i := n.Find("")

	assert.Equal(t, i, -1)
}

func TestNameReturnsNotFoundForUnbalancedQuotes(t *testing.T) {
	n := NewName("param", false)

	i := n.Find(`"quoted`)

	assert.Equal(t, i, -1)
}
