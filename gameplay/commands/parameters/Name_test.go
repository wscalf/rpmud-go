package parameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameFindsSingleWordNameAtEndOfInput(t *testing.T) {
	c := "value"

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), c)
	assert.Equal(t, r, "")
}

func TestNameFindsSingleWordNameWithRemainingInput(t *testing.T) {
	c := "value more"

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "value")
	assert.Equal(t, r, "more")
}

func TestNameFindsQuotedNameAtEndOfInput(t *testing.T) {
	c := `"quoted name"`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "quoted name")
	assert.Equal(t, r, "")
}

func TestNameFindsQuotedNameWithRemainingInput(t *testing.T) {
	c := `"quoted name" more`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, v.Single(), "quoted name")
	assert.Equal(t, r, "more")
}

func TestNameReturnsNotFoundForEmptyInput(t *testing.T) {
	n := NewName("param", false)

	i := n.Find("")

	assert.Equal(t, -1, i)
}

func TestNameReturnsNotFoundForUnbalancedQuotes(t *testing.T) {
	n := NewName("param", false)

	i := n.Find(`"quoted`)

	assert.Equal(t, -1, i)
}
