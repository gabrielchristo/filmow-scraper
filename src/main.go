package main

import (
	_ "fmt"
	"os"
	//"flag"
	"log"
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
		log.Println("Visited", r.Request.URL)
	})

	// callback for movies list
	c.OnHTML("#movies-list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
			// parse ID
			id, _ := strconv.Atoi(elem.Attr("data-movie-pk"))
			if SliceContains(moviesIDs, id) {
				//log.Println("movie ID", id, "already in list")
				return // skip duplicates
			}
			moviesIDs = append(moviesIDs, id)
			//log.Println("Filmow ID", id)

			// parse comments number
			comments := elem.ChildText(".badge-num-comments")
			comments = strings.Replace(comments, ",", "", -1)
			comments = strings.Replace(comments, "K", "00", -1) // 1,8K to 1800
			//log.Println("Comments for", id, "=", comments)
			
			// parsing info from obtained movie ID
			wg.Add(1)
			go Parse(id, comments, wg)
		})
	})

	// callback for shows list ?

	// callback for next movies page
	c.OnHTML(".pagination-centered", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
			newPage := elem.Attr("href")
			//log.Println("New page", newPage)
			if !strings.Contains(newPage, "filmes") { // avoid visit "filmes" url workaround, colly bug?
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

	// show or save movies
	log.Println("Movie IDs Count", len(moviesIDs))
	//ShowAllMovies()
	//ShowAllMoviesInOrder(moviesIDs)
	//SaveAllMovies(input)
	SaveAllMoviesInOrder(moviesIDs, input)
}
