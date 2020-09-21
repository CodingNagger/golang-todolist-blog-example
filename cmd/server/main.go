package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingnagger/golang-todolist-blog-example/pkg/completion"
	"github.com/codingnagger/golang-todolist-blog-example/pkg/creation"
	"github.com/codingnagger/golang-todolist-blog-example/pkg/http/rest"
	"github.com/codingnagger/golang-todolist-blog-example/pkg/listing"
	"github.com/codingnagger/golang-todolist-blog-example/pkg/storage/memory"
)

func main() {
	var creator creation.Service
	var completor completion.Service
	var lister listing.Service

	s := memory.NewStorage()

	creator = creation.NewService(s)
	completor = completion.NewService(s)
	lister = listing.NewService(s)

	// set up the HTTP server
	router := rest.Handler(creator, lister, completor)

	fmt.Println("The tasks server is running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
