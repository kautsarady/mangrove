package main

import (
	"flag"
	"log"

	"github.com/kautsarady/mangrove/query"
)

func main() {
	total := flag.Int("ti", 1000, "desired total product to fetch")
	mload := flag.Int("ml", 1000, "default=1000, depend on host machine, higher is faster")
	output := flag.String("o", "data.json", "default=data.json, output file name, must be a JSON file")
	flag.Parse()

	stream := query.Make(*mload, *total).FetchToStream()

	for i := 1; i <= int(*total); i++ {
		query.AppendJSON(*output, <-stream)
		log.Printf("Finished writing %d data", i)
	}
}
