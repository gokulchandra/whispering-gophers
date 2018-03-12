// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"github.com/gokul/whispering-gophers/talk/code/io"
)

type Message struct {
	Body string
}

func main() {
	s:= bufio.NewScanner(io.Reader(os.Stdin))
	encoder := json.NewEncoder(os.Stdout)

	for s.Scan() {
		msg := Message{s.Text()}
		err := encoder.Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
