package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Book rep struct & JSON
type Book struct {
	Title       string     `json:"name"`
	Authors     []Author   `json:"authors"`
	Categories  []Category `json:"categories"`
	ImageURL    string     `json:"thumbnail"`
	GramedLink  string     `json:"href"`
	Tags        []Tag      `json:"tags"`
	Description string     `json:"description"`
}

// fetch and embed book description then send it through channel
func (book Book) setDescAndSend(stream chan Book) {

	res, err := http.Get(book.GramedLink)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	desc, ok := data["description"].(string)
	if !ok {
		fmt.Errorf("error getting description key in response body")
	}

	// filtering unwanted desciption text
	r := strings.NewReplacer("<p>", "", "</p>", "", "\n", "", "\r", "")
	book.Description = r.Replace(desc)

	stream <- book
}

// Author rep struct & sub-JSON
type Author struct {
	Name string `json:"title"`
}

// Category rep struct & sub-JSON
type Category struct {
	Title string `json:"title"`
}

// Tag rep struct & sub-JSON
type Tag struct {
	Title string `json:"title"`
}

// Make query
func Make(cat string, PerPage, TotalItems int) Query {
	return Query{cat, PerPage, TotalItems}
}

// Set queries
func Set(ml int, q ...Query) *Queries {
	return &Queries{Q: q, MaxLoad: ml}
}

// Queries struct
type Queries struct {
	Q       []Query
	Books   []Book
	MaxLoad int
}

// FetchToStream fetch all query to stream channel
func (qr *Queries) FetchToStream() chan Book {

	stream := make(chan Book)

	go func() {
		// range over syncrhonously queries (category)
		for _, q := range qr.Q {

			MaxPagePerFetch := qr.MaxLoad / q.PerPage
			MaxPage := q.TotalItems / q.PerPage

			// partially fetch page per range syncrhonously
			for i := 1; i <= MaxPage; i += MaxPagePerFetch {
				from, to := getRange(i, MaxPage, MaxPagePerFetch)
				q.FetchRange(from, to, stream)
			}
		}
	}()

	return stream
}

// Query struct
type Query struct {
	Category   string
	PerPage    int
	TotalItems int
}

// FetchRange per page range
func (q *Query) FetchRange(from, to int, stream chan Book) {

	log.Printf("Fetching \"%s\" books from page %d to %d...\n", q.Category, from, to)

	// wait until all page fetched then jump to the next range
	var wg sync.WaitGroup
	for i := from; i <= to; i++ {
		wg.Add(1)
		go q.FetchPage(i, stream, &wg)
	}
	wg.Wait()
}

// FetchPage fetch page data
func (q *Query) FetchPage(page int, stream chan Book, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(fmt.Sprintf(
		"https://www.gramedia.com/api/products/?category=%s&format=json&page=%d&per_page=%d",
		q.Category,
		page,
		q.PerPage,
	))
	if err != nil {
		fmt.Errorf("%v", err)
	}

	if res.StatusCode != http.StatusOK {
		stream <- Book{}
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	books := []Book{}

	err = json.Unmarshal(body, &books)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	for _, book := range books {
		go book.setDescAndSend(stream)
	}

}
