package main

import (
	"flag"
	"log"
	"time"

	"github.com/kautsarady/mangrove/query"
)

func main() {
	total := flag.Int("ti", 9443, "desired total product to fetch")
	mload := flag.Int("ml", 2000, "default=500, depend on host machine, higher is faster")
	output := flag.String("o", "data.json", "default=data.json, output file name, must be a JSON file")
	flag.Parse()

	q := query.Make(*mload, *total)

	t := time.Now()
	stream := q.FetchToStream()

	for i := 1; i <= int(*total); i++ {
		query.AppendJSON(*output, <-stream)
		// book := <-stream
		// log.Printf("Got %s", book.Title)
		log.Printf("Finished writing %d data", i)
	}

	log.Printf("%v", time.Since(t))
	// res, err := query.ReadJSON("data.json")
	// if err != nil {
	// 	fmt.Errorf("asd")
	// }

	// fmt.Println(len(res))
}
