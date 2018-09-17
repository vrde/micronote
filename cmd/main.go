package main

import (
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
	fmt.Printf("%v", notes)

	defer f.Close()
}
