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

const (
	// MaxPerPage Currently gramedia limit max item load per page to 100
	MaxPerPage = 100
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
func (book Book) setDescAndSend(stream chan Book, bookWG *sync.WaitGroup) {

	defer bookWG.Done()

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

// Query struct
type Query struct {
	MaxLoad    int
	PerPage    int
	TotalItems int
}

// Make query
func Make(MaxLoad int, TotalItems int) *Query {
	pp := MaxPerPage

	if TotalItems < pp {
		pp = TotalItems
	}

	return &Query{
		MaxLoad:    MaxLoad,
		PerPage:    pp,
		TotalItems: TotalItems,
	}
}

// FetchToStream fetch all query to stream channel
func (q *Query) FetchToStream() chan Book {

	stream := make(chan Book)

	go func() {

		RangePerFetch := q.MaxLoad / q.PerPage
		TilPage := q.TotalItems / q.PerPage

		for i := 1; i <= TilPage; i += RangePerFetch {

			from, to := getRange(i, TilPage, RangePerFetch)
			log.Printf("Fetching products from page %d - %d", from, to)

			var pageWG sync.WaitGroup
			for i := from; i <= to; i++ {

				pageWG.Add(1)
				go q.FetchPage(i, stream, &pageWG)

			}
			pageWG.Wait()
		}
	}()

	return stream
}

// FetchPage fetch page data
func (q *Query) FetchPage(page int, stream chan Book, pageWG *sync.WaitGroup) {

	defer pageWG.Done()

	log.Printf("Fetching %d products at page %d", q.PerPage, page)

	res, err := http.Get(fmt.Sprintf(
		"https://www.gramedia.com/api/products?format=json&page=%d&per_page=%d",
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

	var bookWG sync.WaitGroup
	for _, book := range books {
		bookWG.Add(1)
		go book.setDescAndSend(stream, &bookWG)
	}
	bookWG.Wait()

}
