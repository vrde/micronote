package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	// "os/exec"
	n "github.com/vrde/micronote/note"
	u "github.com/vrde/micronote/utils"
	"strings"
)

func main() {
	// f, err := os.Open(os.Getenv("HOME") + "/.micronote/notes")
	var date, tags, home string
	flag.StringVar(&date, "date", "all", "a date")
	flag.StringVar(&date, "d", "all", "a date")
	flag.StringVar(&tags, "tags", "", "Comma separated list of tags")
	flag.StringVar(&tags, "t", "", "Comma separated list of tags")
	flag.StringVar(&home, "h", os.Getenv("HOME")+"/.micronote/notes", "Home")
	flag.Parse()

	f, err := os.OpenFile(home, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	text := strings.Join(os.Args[len(os.Args)-flag.NArg():], " ")

	if flag.NFlag() == 0 && len(os.Args[1:]) == 0 {
		flag.PrintDefaults()

		// cmd := "/usr/bin/nvim"
		// args := []string{
		// 	"-c", `"%"`,
		// 	"-c", `"normal O"`,
		// 	"-c", `".!date --iso-8601=seconds"`,
		// 	"-c", `"normal O"`,
		// 	"-c", `"startinsert"`,
		// 	home}

		// if err := exec.Command(cmd, args...).Run(); err != nil {
		// if err := exec.Command("vim", "ciao").Start(); err != nil {
		// 	log.Fatal(err)
		// }
	} else if text != "" {
		if _, err := f.Seek(0, 2); err != nil {
			log.Fatal(err)
		}
		note := n.Note{
			Date: u.NewDate("now"),
			Tags: strings.Split(tags, " "),
			Text: text,
		}
		f.WriteString(note.String())
	} else {
		p := n.NewParser()
		notes, err := p.Parse(f)
		if err != nil {
			log.Fatal(err)
		}
		r := n.Search(notes, u.NewDate(date), tags)
		fmt.Printf("%s", &r)
	}
}
