package micronote

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	var note *Note
	var err error
	string_date := "2018-09-16T17:53:18+02:00"
	date, _ := time.Parse(time.RFC3339, string_date)
	p := NewParser()
	expected := Note{
		Date: date,
		Tags: []string{"first-tag", "second-tag"},
		Text: "This is a note.\n",
	}

	note, err = p.Feed(string_date+" first-tag second-tag", false)
	assert.Nil(t, err)
	assert.Nil(t, note)

	note, err = p.Feed("This is a note.", false)
	assert.Nil(t, err)
	assert.Nil(t, note)

	note, err = p.Feed("", true)
	assert.Nil(t, err)
	assert.Equal(t, expected, *note)
}

func TestParse(t *testing.T) {
	p := NewParser()
	l := time.FixedZone("CEST", 2*60*60)
	e := Notes{
		Note{
			Date: time.Date(2018, time.September, 16, 12, 0, 0, 0, l),
			Tags: []string{"foo", "bar"},
			Text: "Some multiline\n\ntext.\n\n",
		},
		Note{
			Date: time.Date(2018, time.September, 16, 11, 0, 0, 0, l),
			Tags: []string{"bar", "baz"},
			Text: "Other text.\n",
		},
	}
	r := strings.NewReader(`2018-09-16T12:00:00+02:00 foo bar
Some multiline

text.

2018-09-16T11:00:00+02:00 bar baz
Other text.
`)
	notes, err := p.Parse(r)
	assert.Nil(t, err)
	assert.Equal(t, e, notes)
}
