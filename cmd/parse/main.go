// parse is a command-line tool to parse ISO8601 duration strings and print their Golang time.Duration equivalents.
package main

import (
	"flag"
	"fmt"
	duration "github.com/sfomuseum/iso8601duration"
	"log"
)

func main() {

	flag.Parse()

	for _, str := range flag.Args() {

		d, err := duration.FromString(str)

		if err != nil {
			log.Fatalf("Failed to parse '%s', %v", str, err)
		}

		fmt.Printf("%v\n", d.ToDuration())
	}
}
