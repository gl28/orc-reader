package main

import (
	"log"

	"github.com/scritchley/orc"
)

func main() {
	r, err := orc.Open("test_data.orc")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	c := r.Select("_col0", "_col1", "_col2")

	for c.Stripes() {

		for c.Next() {
			
			log.Println(c.Row())
		}
	}

	if err := c.Err(); err != nil {
		log.Fatal(err)
	}
}