package micronote

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Note struct {
	Date string
	Tags []string
	Text string
}

func (n *Note) String() string {
	return fmt.Sprintf("%s %v\n%s", n.Date, n.Tags, n.Text)
}

type Notes []Note

func (n *Notes) String() string {
	buffer := ""
	for i, note := range *n {
		buffer += fmt.Sprintf("%d. %s\n", i, &note)
	}
	return buffer
}

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
	tokens := strings.Split(s, " ")
	matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}`, tokens[0])
	if err == nil && matched == true {
		result = p.current
		p.current = &Note{}
		p.current.Date = tokens[0]
		p.current.Tags = tokens[1:]
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
