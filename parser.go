package micronote

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"time"
)

type Note struct {
	Date time.Time
	Tags []string
	Text string
}

type Notes []Note

type Parser struct {
	current *Note
}

func NewParser() *Parser {
	return &Parser{
		current: nil,
	}
}

func (p *Parser) Feed(s string, eof bool) (*Note, error) {
	var result *Note
	if eof == true {
		return p.current, nil
	}
	tokens := strings.SplitN(s, " ", 2)
	date, err := time.Parse(time.RFC3339, tokens[0])
	if err == nil {
		result = p.current
		p.current = &Note{}
		p.current.Date = date
		p.current.Tags = strings.Split(tokens[1], " ")
	} else {
		if p.current == nil {
			return nil, errors.New("Invalid token, expected date, got " + tokens[0])
		}
		p.current.Text += s + "\n"
	}

	return result, nil
}

func (p *Parser) Parse(r io.Reader) (Notes, error) {
	notes := Notes{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		note, err := p.Feed(scanner.Text(), false)
		if err != nil {
			return nil, err
		}
		if note != nil {
			notes = append(notes, *note)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	note, err := p.Feed("", true)
	if err != nil {
		return nil, err
	}
	if note != nil {
		notes = append(notes, *note)
	}

	return notes, nil
}
