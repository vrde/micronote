package main

import (
	"flag"
	"fmt"
	m "github.com/vrde/micronote"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/.micronote/notes")
	if err != nil {
		log.Fatal(err)
	}

	p := m.NewParser()
	notes, err := p.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	date := flag.String("date", "", "a date")
	flag.Parse()

	fmt.Println("Search for: " + *date)
	r := m.Search(notes, *date)
	fmt.Printf("%s", &r)
}
