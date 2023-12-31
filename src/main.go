package main

import (
	_ "fmt"
	"os"
	//"flag"
	//"log"
	"sync"
	"github.com/gocolly/colly"
	"strings"
	"strconv"
)

func UNUSED(x ...interface{}) {} // UNUSED allows unused variables to be included in Go programs

func main() {

	// waitgroup var
	wg := &sync.WaitGroup{}

	// IDs array
	var moviesIDs []int

	// filmow username
	var input string

	// get username from command line
    //fmt.Print("Enter username: ")
    //fmt.Scanln(&input)

	// get username from args
	input = os.Args[1]

	// update filmow URLs with user
	//moviesURL := strings.Replace("https://filmow.com/usuario/USERNAME/filmes/ja-vi/", "USERNAME", input, -1)
	//showsURL := strings.Replace("https://filmow.com/usuario/USERNAME/series/ja-vi/", "USERNAME", input, -1)
	allURL := strings.Replace("https://filmow.com/usuario/USERNAME/ja-vi/", "USERNAME", input, -1)

	// create collector
	c := colly.NewCollector(
		colly.AllowedDomains("filmow.com"),
		//colly.Async(true), // faster but will lose watched order
	)

	c.OnResponse(func(r *colly.Response) {
		//log.Println("Visited", r.Request.URL)
	})

	// callback for movies list
	c.OnHTML("#movies-list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
			id, _ := strconv.Atoi(elem.Attr("data-movie-pk"))
			moviesIDs = append(moviesIDs, id)
			//log.Println("Filmow ID", id)
			wg.Add(1)
			go Parse(id, wg) // parsing info from obtained movie ID
		})
	})

	// callback for shows list ?

	// callback for next movies page
	c.OnHTML(".pagination-centered", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
			newPage := elem.Attr("href")
			if !strings.Contains(newPage, "pagina=1") { // skip first page to avoid duplicates
				c.Visit(e.Request.AbsoluteURL(newPage))
			}
		})
	})

	// callback for next shows page ?

	// visit main URLs
	//c.Visit(moviesURL)
	//c.Visit(showsURL)
	c.Visit(allURL)

	// wait for colly and api calls
	c.Wait()
	wg.Wait()

	// show all movies
	//log.Println("Movie Count", len(moviesIDs))
	ShowAllMoviesInOrder(moviesIDs)
}
