package main 

import (
	"fmt"
	"strconv"
	"log"
)

type Movie struct {
	id int
	title string
	title_orig string
	director string
	year string
	rate string
	comments string
}

var movies []Movie

/*
	@description: Creates movie object with given data
	@return: movie object
*/
func CreateMovie(id int, title string, title_orig string, director string, year string, rate string, comments string) *Movie {

	m := Movie{}
	m.id = id
	m.title = title
	m.title_orig = title_orig
	m.director = director
	m.year = year
	m.rate = rate
	m.comments = comments

	return &m
}

/*
	@description: Add movie struct to movies array
	@param m: movie structure
*/
func AddMovie(m *Movie) {
	movies = append(movies, *m)
	//fmt.Println("Added movie", m.title)
}

/*
	@description: return movie objet at array with given ID, if exists
	@return: movie object or nil
*/
func GetMovieByID(id int) *Movie {

	for _, element := range movies {
		if element.id == id {
			return &element
		}
	}

	return nil
}

/*
	@description: convert movies array to csv WriteAll() accepted format
*/
func TransformMoviesTo2DSlice(allMovies []Movie) [][]string {
	numRows := len(allMovies)
    result := make([][]string, numRows+1)
    // add header row
    result[0] = []string{"filmow_id","title","original_title","comments","year","filmow_rate","director"}
    // add data rows
    for i := 0; i < numRows; i++ {
        result[i+1] = []string{
			strconv.FormatInt(int64(allMovies[i].id), 10),
            allMovies[i].title,
            allMovies[i].title_orig,
			allMovies[i].comments,
			allMovies[i].year,
			allMovies[i].rate,
			allMovies[i].director,
        }
    }
    return result
}

/*
	@description: 
*/
func SaveAllMovies(username string) {
	log.Println("Movie Array Count", MovieCount())
	WriteToFile(fmt.Sprintf("../output_%s.csv", username), TransformMoviesTo2DSlice(movies))
}

/*
	@description: 
*/
func SaveAllMoviesInOrder(idArray []int, username string) {
	log.Println("Movie Array Count", MovieCount())
	var new_array []Movie
	for _, id := range idArray {
		movie := GetMovieByID(id)
		if movie != nil {
			new_array = append(new_array, *movie)
		} else {
			log.Println("SaveAllMoviesInOrder() null movie with id", id)
			var new_try *Movie
			for new_try == nil {
				new_try = SimpleParse(id) // new call to api
			}
			new_array = append(new_array, *new_try)
		}
	}
	WriteToFile(fmt.Sprintf("../output_%s.csv", username), TransformMoviesTo2DSlice(new_array))
	movies = movies[:0] // clear slice
}

/*
	@description: print all movies at array
*/
func ShowAllMovies() {
	log.Println("Movie Array Count", MovieCount())
	fmt.Println("filmow_id, title, original_title, comments, year, filmow_rate, director")
	for _, movie := range movies {
		output := fmt.Sprintf("%d, %s, %s, %s, %s, %s, %s", movie.id, movie.title, movie.title_orig, movie.comments, movie.year, movie.rate, movie.director)
		fmt.Println(output)
	}
}

/*
	@description: print all movies with a given order
*/
func ShowAllMoviesInOrder(idArray []int) {
	log.Println("Movie Array Count", MovieCount())
	fmt.Println("filmow_id, title, original_title, comments, year, filmow_rate, director")
	for _, id := range idArray {
		movie := GetMovieByID(id)
		if movie != nil {
			output := fmt.Sprintf("%d, %s, %s, %s, %s, %s, %s", movie.id, movie.title, movie.title_orig, movie.comments, movie.year, movie.rate, movie.director)
			fmt.Println(output)
		} else {
			log.Println("ShowAllMoviesInOrder() null movie with id", id)
			var new_try *Movie
			for new_try == nil {
				new_try = SimpleParse(id) // new call to api
			}
			output := fmt.Sprintf("%d, %s, %s, %s, %s, %s, %s", new_try.id, new_try.title, new_try.title_orig, new_try.comments, new_try.year, new_try.rate, new_try.director)
			fmt.Println(output)
		}
	}
}

/*
	@description: return movie count in array
*/
func MovieCount() int {
	return len(movies)
}
