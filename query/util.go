package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// AppendJSON file with incoming data (buffering)
func AppendJSON(filepath string, book Book) {
	var books []Book

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(byteValue, &books); err != nil {
		fmt.Errorf("%v", err)
	}

	books = append(books, book)

	dataJSON, err := json.Marshal(books)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	if err := ioutil.WriteFile(filepath, dataJSON, 0644); err != nil {
		fmt.Errorf("%v", err)
	}
}

// ReadJSON parsing json file into []Book
func ReadJSON(filepath string) ([]Book, error) {
	var books []Book

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(byteValue, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// Simple algorithm fo getting next pages range to fetch
func getRange(i, MaxPage, MaxPagePerFetch int) (int, int) {
	from, to := i, i+(MaxPagePerFetch-1)

	if to > MaxPage {
		to = (MaxPage % MaxPagePerFetch) + from - 1
	}

	return from, to
}
