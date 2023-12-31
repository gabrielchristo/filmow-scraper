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

/*
	@description: filmow scraper init function
*/
func search(username string, main_url string) []int {

	// waitgroup var
	wg := &sync.WaitGroup{}

	// IDs array
	var moviesIDs []int

	// update filmow URLs with user
	//moviesURL := strings.Replace("https://filmow.com/usuario/USERNAME/filmes/ja-vi/", "USERNAME", input, -1)
	//showsURL := strings.Replace("https://filmow.com/usuario/USERNAME/series/ja-vi/", "USERNAME", input, -1)
	allURL := strings.Replace(main_url, "USERNAME", username, -1)

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

	// visit main URLs
	//c.Visit(moviesURL)
	//c.Visit(showsURL)
	c.Visit(allURL)

	// wait for colly and api calls
	c.Wait()
	wg.Wait()

	return moviesIDs
}

/*
	@description: entrypoint function
*/
func main() {

	// filmow username
	var input string

	// get username from command line
	//fmt.Print("Enter username: ")
	//fmt.Scanln(&input)

	// get username from args
	input = os.Args[1]

	// show or save watched movies
	watched := search(input, "https://filmow.com/usuario/USERNAME/ja-vi/")
	log.Println("Watched Movies ID Count", len(watched))
	//ShowAllMovies()
	//ShowAllMoviesInOrder(moviesIDs)
	//SaveAllMovies(input)
	SaveAllMoviesInOrder(watched, input + "_watched")
	
	// show or save want to watch movies
	want_to_watch := search(input, "https://filmow.com/usuario/USERNAME/quero-ver/")
	log.Println("Want to Watch Movies ID Count", len(want_to_watch))
	SaveAllMoviesInOrder(want_to_watch, input + "_want_to_watch")
}
