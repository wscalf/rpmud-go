package parameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupReturnsNotFoundWhenNotFound(t *testing.T) {
	c := ""
	g := NewNameGroup("names", false)

	i := g.Find(c)

	assert.Equal(t, -1, i)
}

func TestGroupConsumesAllNamesFound(t *testing.T) {
	c := `foo bar "baz"`
	g := NewNameGroup("names", false)

	i := g.Find(c)

	assert.Equal(t, 0, i)

	v, remaining := g.Consume(c)
	assert.Equal(t, remaining, "")
	assert.Equal(t, []string{"foo", "bar", "baz"}, v.Multiple())
}
