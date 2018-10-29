package main

import (
	"flag"
	"log"

	"github.com/kautsarady/mangrove/query"
)

func main() {
	category := flag.String("c", "buku", "product category to fetch")
	ppage := flag.Int("pp", 20, "MAX=100, important to determine fetch range")
	total := flag.Int("ti", 40, "desired total product to fecth")
	mload := flag.Int("ml", 500, "default=500, depend on host machine, higher is better")
	flag.Parse()

	q := query.Make(*category, *ppage, *total)

	qr := query.Set(*mload, q)

	stream := qr.FetchToStream()

	// Write json asynchronously
	for i := 1; ; i++ {
		query.AppendJSON("data.json", <-stream)
		log.Printf("Current data count %d", i)
	}
}
