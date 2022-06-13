package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "Test page", Body: []byte("This is a simple page")}
	p1.save()

	p2, _ := loadPage("Test page")
	fmt.Println(string(p2.Body))
}
