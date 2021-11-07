package main 

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"log"
	"regexp"
	"strings"
)

// base URL for movie info API
const movieAPI string = "https://filmow.com/async/tooltip/movie/?movie_pk="

/*
	@description: creates movie object from given ID
	@param id: movie ID from filmow
	@return: created movie structure
*/
func Parse(id int) *Movie {

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

		newMovie := CreateMovie(id, GetTitle(movieKey), GetOriginalTitle(movieKey), GetDirector(htmlKey), GetYear(htmlKey), GetRate(htmlKey))

		// add to movies array
		AddMovie(newMovie)

		return newMovie
	}

	return nil
}

func GetTitle(json map[string]interface{}) string {
	title := json["title"].(string)
	return strings.Replace(title, ",", " ", -1)
}

func GetOriginalTitle(json map[string]interface{}) string {
	title := json["title_orig"].(string)
	return strings.Replace(title, ",", " ", -1)
}

func GetRate(html string) string {
	exp, _ := regexp.Compile("Nota: (.*?) estrela")
	result := exp.FindString(html)
	//log.Println("Match for rate regex:", result, result[6:9])

	return strings.Replace(result[6:9], ",", ".", -1)
}

func GetYear(html string) string {
	exp, _ := regexp.Compile("Lançamento Mundial: </b>[0-9]*")
	result := exp.FindString(html)
	//log.Println("Match for year regex:", result, result[25:29])

	if len(result) < 30 {
		return ""
	} else {
		return result[25:29]
	}
}

func GetDirector(html string) string {
	exp, _ := regexp.Compile("Diretor:</b> (<a.*\">.*</a>)*")
	temp := exp.FindString(html)

	names, _ := regexp.Compile("\">.*?<")
	result := names.FindString(temp)

	// as goland regex does not support lookbehind, we need to do a static replacement
	result = strings.Replace(result, "<", ";", -1)
	result = strings.Replace(result, "\">", "", -1)

	//log.Println(result)

	return result
}