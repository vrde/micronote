package micronote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	n := Notes{
		Note{
			Date: "2018-09-16T12:00:00+02:00",
			Tags: []string{"foo", "bar"},
			Text: "Some multiline\n\ntext.\n\n",
		},
		Note{
			Date: "2018-09-16T11:00:00+02:00",
			Tags: []string{"bar", "baz"},
			Text: "Other text.\n",
		},
	}

	r := Search(n, "2018-09-16T11")
	assert.Equal(t, Notes{n[1]}, r)
}
