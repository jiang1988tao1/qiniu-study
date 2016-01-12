package main

import (
	"../20160112/restful"
	"fmt"
	"net/http"
	"net/url"
)

type Book struct {
}

func (book Book) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}
func (book Book) Post(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}
func (book Book) Put(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	return 200, "hello world", nil
}

func main() {
	port := 3000
	book := new(Book)
	app := restful.NewAPI()
	app.Register(book, "/books", "/bookbyjt")
	fmt.Println("Listen on port: ", port)
	go app.Start(port)

	resp, err := http.Get("http://127.0.0.1:3000/books")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)
}
