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

	assert.Equal(t, c, v.Single())
	assert.Equal(t, "", r)
}

func TestNameFindsSingleWordNameWithRemainingInput(t *testing.T) {
	c := "value more"

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, "value", v.Single())
	assert.Equal(t, "more", r)
}

func TestNameFindsQuotedNameAtEndOfInput(t *testing.T) {
	c := `"quoted name"`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, "quoted name", v.Single())
	assert.Equal(t, "", r)
}

func TestNameFindsQuotedNameWithRemainingInput(t *testing.T) {
	c := `"quoted name" more`

	n := NewName("param", false)
	i := n.Find(c)

	assert.Equal(t, 0, i)

	v, r := n.Consume(c)

	assert.Equal(t, "quoted name", v.Single())
	assert.Equal(t, "more", r)
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
