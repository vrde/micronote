package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	n "github.com/vrde/micronote/note"
	u "github.com/vrde/micronote/utils"
)

func main() {
	// f, err := os.Open(os.Getenv("HOME") + "/.micronote/notes")
	var date, tags string
	flag.StringVar(&date, "date", "", "a date")
	flag.StringVar(&date, "d", "", "a date")
	flag.StringVar(&tags, "tags", "", "Comma separated list of tags")
	flag.StringVar(&tags, "t", "", "Comma separated list of tags")
	flag.Parse()

	f, err := os.OpenFile(os.Getenv("HOME")+"/.micronote/notes", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if flag.NFlag() == 0 && len(os.Args[1:]) == 0 {
		flag.PrintDefaults()
	} else if flag.NFlag() == 0 && len(os.Args[1:]) > 0 {
		if _, err := f.Seek(0, 2); err != nil {
			log.Fatal(err)
		}
		note := n.Note{
			Date: u.NewDate(""),
			Tags: []string{},
			Text: strings.Join(os.Args[1:], " "),
		}
		f.WriteString(note.String())
	} else {
		p := n.NewParser()
		notes, err := p.Parse(f)
		if err != nil {
			log.Fatal(err)
		}
		r := n.Search(notes, u.NewDate(date))
		fmt.Printf("%s", &r)
	}
}
