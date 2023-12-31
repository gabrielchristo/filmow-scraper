package main 

import (
	"fmt"
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
	@description: print all movies at array
*/
func ShowAllMovies() {
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
	fmt.Println("filmow_id, title, original_title, comments, year, filmow_rate, director")
	for _, id := range idArray {
		movie := GetMovieByID(id)
		output := fmt.Sprintf("%d, %s, %s, %s, %s, %s, %s", movie.id, movie.title, movie.title_orig, movie.comments, movie.year, movie.rate, movie.director)
		fmt.Println(output)
	}
}

/*
	@description: return movie count in array
*/
func MovieCount() int {
	return len(movies)
}
