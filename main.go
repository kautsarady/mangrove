package main

import (
	"flag"
	"log"

	"github.com/kautsarady/mangrove/query"
)

func main() {
	total := flag.Int("ti", 1000, "desired total product to fetch")
	output := flag.String("o", "data.json", "default=data.json, output file name, must be a JSON file")
	flag.Parse()

	log.Printf("Fetching %d data to %s ...", *total, *output)
	
	stream := query.Make(*total).FetchToStream()

	for i := 1; i <= int(*total); i++ {
		query.AppendJSON(*output, <-stream)
		log.Printf("Finished writing %d data", i)
	}
}
