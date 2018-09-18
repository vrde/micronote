package note

import (
	"strings"
)

func contains(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func Search(notes Notes, date, tag string) Notes {
	r := Notes{}
	for _, x := range notes {
		if strings.Index(x.Date, date) == 0 && (tag == "" || contains(tag, x.Tags)) {
			r = append(r, x)
		}
	}
	return r
}
