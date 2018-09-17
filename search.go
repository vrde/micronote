package micronote

import (
	"strings"
)

func Search(n Notes, t string) Notes {
	r := Notes{}
	for _, x := range n {
		if strings.Index(x.Date, t) != -1 {
			r = append(r, x)
		}
	}
	return r
}
