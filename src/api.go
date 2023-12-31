package main 

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"log"
	"regexp"
	"strings"
	"sync"
)

// base URL for movie info API
const movieAPI string = "https://filmow.com/async/tooltip/movie/?movie_pk="

/*
	@description: creates movie object from given ID
	@param id: movie ID from filmow
	@return: created movie structure
*/
func Parse(id int, comments string, wg *sync.WaitGroup) *Movie {

	response, err := http.Get(movieAPI + strconv.Itoa(id))
	
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close() // closing until function end
	defer wg.Done()
	
	var jsonResponse map[string]interface{}

	if response.StatusCode == 200 { // ok
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal([]byte(bodyText), &jsonResponse)

		movieKey := jsonResponse["movie"].(map[string]interface{})
		htmlKey := jsonResponse["html"].(string)

		newMovie := CreateMovie(id, GetTitle(movieKey), GetOriginalTitle(movieKey), GetDirector(htmlKey), GetYear(htmlKey), GetRate(htmlKey), comments)

		// add to movies array
		AddMovie(newMovie)

		return newMovie
	}

	return nil
}

func SimpleParse(id int) *Movie {
	log.Println("New api call for ID", id)
	response, err := http.Get(movieAPI + strconv.Itoa(id))
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close() // closing until function end
	var jsonResponse map[string]interface{}
	if response.StatusCode == 200 { // ok
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal([]byte(bodyText), &jsonResponse)
		movieKey := jsonResponse["movie"].(map[string]interface{})
		htmlKey := jsonResponse["html"].(string)
		newMovie := CreateMovie(id, GetTitle(movieKey), GetOriginalTitle(movieKey), GetDirector(htmlKey), GetYear(htmlKey), GetRate(htmlKey), "Sem informação")
		return newMovie
	}
	return nil
}

/*

*/
func GetTitle(json map[string]interface{}) string {
	title := json["title"].(string)
	return strings.Replace(title, ",", " ", -1)
}

/*

*/
func GetOriginalTitle(json map[string]interface{}) string {
	title := json["title_orig"].(string)
	return strings.Replace(title, ",", " ", -1)
}


/*

*/
func GetRate(html string) string {
	exp, _ := regexp.Compile("Nota: (.*?) estrela")
	result := exp.FindString(html)
	//log.Println("Match for rate regex:", result, result[6:9])

	return strings.Replace(result[6:9], ",", ".", -1) // point as decimal separator to avoid CSV problems
}

/*

*/
func GetYear(html string) string {
	exp, _ := regexp.Compile("Mundial: </b>[0-9].*\n\t")
	result := exp.FindString(html)
	//log.Println("Partial match for year regex:", result)

	split1 := strings.Split(result, "</b>")
	if len(split1) < 2 {
		return "Sem informação"
	}

	split2 := strings.Split(split1[1], "\n")
	if len(split2) < 1 {
		return "Sem informação"
	}
	//log.Println("Match for year regex:", split2[0])
	return split2[0]
}

/*

*/
func GetDirector(html string) string {
	exp, _ := regexp.Compile("Diretor:</b> (<a.*\">.*</a>)*")
	temp := exp.FindString(html)

	names, _ := regexp.Compile("\">(.*?)<")
	result := names.FindAllString(temp, -1)
	//log.Println(result)

	if len(result) < 1 {
		return "Sem informação"
	}

	var corrected_list []string
	for _, director := range result {
		// as golang regex does not support lookbehind or lookahead, we need to do a static replacement
		director = strings.Replace(director, "<", "", -1)
		director = strings.Replace(director, "\">", "", -1)
		corrected_list = append(corrected_list, director)
	}

	final := strings.Join(corrected_list, " | ")
	//log.Println(final)

	return final
}
