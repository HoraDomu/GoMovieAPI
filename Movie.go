(main.go)
package main		

import (
	"fmt"
	"strings"
)



func main(){
	var name string 
	fmt.Print("What's your name?: ")
	fmt.Scanln(&name)

	
	var genre string 
	fmt.Print("What genre are we watching?:  ")
	fmt.Scanln(&genre)
	genre = strings.ToLower(genre)
	
	fmt.Println("Hello,", name+"!")
  fmt.Println("You like", genre, "movies. Great choice! Here are some of the best", genre+ " movies currently.")
	
	movies := GetMovies(genre)

	PrintMovies(movies)

}


(movies.go)
package main

import (
	 "encoding/json" // decode JSON from TMDB // 
    "fmt"           // for printing // 
    "io/ioutil"     // read HTTP response body //
    "net/http"      // make HTTP requests //
    "strings"       // handle lowercase input //
)

type Movie struct {
    Title string
}

type TMDBMovie struct{
	Title  string 'json:"title"'
} ReleaseDate string 'json:"title"'

type TMDBResponse struct{
	Results []TMDBMovie 'json:"results"'
}


var GenreIDs = map[string]int{
	"action": 28, 
	"comedy": 35, 
	"horror": 27,
}


const API_KEY = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODBlMjAwYmI3NTM0YjllM2UwNmZlMGRlZDlhM2ZhNyIsIm5iZiI6MTc1ODU5Mjc5My41ODA5OTk5LCJzdWIiOiI2OGQxZmYxOTc3ODkxZDIxYmNjZDgwYzQiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.lVnjNYNiwFzlfJ12jeRyuYMEYWJUZu4tRQLX4CHv2f4"


func GetMovies(genre string) []Movie{
	genre = strings.ToLower(genre)
	genreID, ok := GenreIDs[genre]
	if !ok{
		return []Movie{}
	}
	url := fmt.Sprintf(

	)
}

func PrintMovies(movies []Movie) {
    if len(movies) == 0 {
        fmt.Println("Sorry, no movies found for that genre.")
        return
    }
    fmt.Println("Top Movies:")
    for i, movie := range movies {
        fmt.Printf("%d. %s\n", i+1, movie.Title)
    }
}


