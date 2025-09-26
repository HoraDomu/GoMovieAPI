package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const API_KEY = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODBlMjAwYmI3NTM0YjllM2UwNmZlMGRlZDlhM2ZhNyIsIm5iZiI6MTc1ODU5Mjc5My41ODA5OTk5LCJzdWIiOiI2OGQxZmYxOTc3ODkxZDIxYmNjZDgwYzQiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.lVnjNYNiwFzlfJ12jeRyuYMEYWJUZu4tRQLX4CHv2f4"

type Movie struct {
	Title string `json:"title"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

func main() {
	var name string
	fmt.Print("What's your name?: ")
	fmt.Scanln(&name)

	var genre string
	fmt.Print("What genre are we watching?: ")
	fmt.Scanln(&genre)
	genre = strings.ToLower(genre)

	fmt.Println("Hello,", name+"!")
	fmt.Println("You like", genre, "movies. Great choice! Here are some of the best", genre, "movies currently.")

	genreMap := map[string]string{
		"action": "28",
		"horror": "27",
		"comedy": "35",
		"drama":  "18",
		"romance":"10749",
	}

	id, ok := genreMap[genre]
	if !ok {
		fmt.Println("Genre not found!")
		return
	}

	url := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?with_genres=%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var movieResponse MovieResponse
	err = json.Unmarshal(body, &movieResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movieResponse.Results {
		fmt.Println(movie.Title)
	}
}


