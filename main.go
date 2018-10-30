package main

import (
	"flag"
	"log"

	"github.com/kautsarady/mangrove/query"
)

func main() {
	category := flag.String("c", "buku", "product category to fetch")
	total := flag.Int("ti", 20, "desired total product to fetch")
	mload := flag.Int("ml", 500, "default=500, depend on host machine, higher is better")
	output := flag.String("o", "data.json", "default=data.json, output file name, must be a JSON file")
	flag.Parse()

	q := query.Make(*category, *total)

	qr := query.Set(*mload, q)

	stream := qr.FetchToStream()

	// Write json synchronously
	for i := 1; i <= int(*total); i++ {
		query.AppendJSON(*output, <-stream)
		log.Printf("Finished writing %d data", i)
	}
}
