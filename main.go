package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// todo: func to build header

func get(url string) {

	response, err := http.Get(url)
	
	if err != nil {
		fmt.Println(err)
	}

	
	defer response.Body.Close() // closing until function end
	
	if response.StatusCode == 200 { // ok
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", bodyText)
	}
	
}

func main() {
    fmt.Print("Enter username: ")

    var input string

    fmt.Scanln(&input)

    fmt.Println(input)
	
	get("http://google.com")
}
