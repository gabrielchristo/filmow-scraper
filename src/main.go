package main

import (
	_ "fmt"
	"os"
	"github.com/gocolly/colly"
	"strings"
	"strconv"
)

func UNUSED(x ...interface{}) {} // UNUSED allows unused variables to be included in Go programs

func main() {

	// IDs array
	var moviesIDs []int

	// get username from command line
    //fmt.Print("Enter username: ")
	var input string
    //fmt.Scanln(&input)
	input = os.Args[1]

	// update filmow URLs with user
	moviesURL := strings.Replace("https://filmow.com/usuario/USERNAME/filmes/ja-vi/", "USERNAME", input, -1)
	showsURL := strings.Replace("https://filmow.com/usuario/USERNAME/series/ja-vi/", "USERNAME", input, -1)

	// create collector
	c := colly.NewCollector(
		colly.AllowedDomains("filmow.com"),
		//colly.Async(true),
	)

	c.OnResponse(func(r *colly.Response) {
		//log.Println("Visited", r.Request.URL)
	})

	// callback for movies list
	c.OnHTML("#movies-list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
			id, _ := strconv.Atoi(elem.Attr("data-movie-pk"))
			moviesIDs = append(moviesIDs, id)
		})
	})

	// callback for shows list

	// callback for next movies page
	c.OnHTML(".pagination-centered", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
			newPage := elem.Attr("href")
			c.Visit(e.Request.AbsoluteURL(newPage))
		})
	})

	// callback for next shows page

	// visit main URLs
	c.Visit(moviesURL)
	c.Visit(showsURL)

	// parsing info from obtained movie IDs
	for _, ID := range moviesIDs {
		Parse(ID)
	}

	// show all movies
	ShowAllMovies()
}
