package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// tmdb api key
const API_KEY = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODBlMjAwYmI3NTM0YjllM2UwNmZlMGRlZDlhM2ZhNyIsIm5iZiI6MTc1ODU5Mjc5My41ODA5OTk5LCJzdWIiOiI2OGQxZmYxOTc3ODkxZDIxYmNjZDgwYzQiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.lVnjNYNiwFzlfJ12jeRyuYMEYWJUZu4tRQLX4CHv2f4"

type Movie struct {
	Title string `json:"title"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

// below is webserver
func Startwebserve() {
	webserver()
}

// serve a simple HTML form for selecting a genre
func handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
  <head><meta charset="utf-8"><title>Movie Picker</title></head>
  <body>
	<h1>Pick a genre</h1>
	<form action="/movies" method="get">
	  <label for="genre">Genre:</label>
	  <select id="genre" name="genre">
		<option value="action">Action</option>
		<option value="horror">Horror</option>
		<option value="comedy">Comedy</option>
		<option value="drama">Drama</option>
		<option value="romance">Romance</option>
	  </select>
	  <button type="submit">Find movies</button>
	</form>
  </body>
</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// /movies?genre=action - returns a JSON array of movie titles for the given genre
func handleMovies(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("genre")
	genre := strings.ToLower(q)

	genreMap := map[string]string{
		"action":  "28",
		"horror":  "27",
		"comedy":  "35",
		"drama":   "18",
		"romance": "10749",
	}

	id, ok := genreMap[genre]
	if !ok || genre == "" {
		http.Error(w, "invalid or missing genre parameter", http.StatusBadRequest)
		return
	}

	// build URL with query params to be safe
	apiURL := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie")
	u, _ := url.Parse(apiURL)
	qv := u.Query()
	qv.Set("with_genres", id)
	u.RawQuery = qv.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "failed to call movie api: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read movie api response", http.StatusInternalServerError)
		return
	}

	var movieResponse MovieResponse
	err = json.Unmarshal(body, &movieResponse)
	if err != nil {
		http.Error(w, "failed to parse movie api response", http.StatusInternalServerError)
		return
	}

	// create a simple slice of titles to return
	titles := make([]string, 0, len(movieResponse.Results))
	for _, m := range movieResponse.Results {
		titles = append(titles, m.Title)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(titles)
}

func webserver() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/movies", handleMovies)
	port := ":8080"
	fmt.Printf("server listening on port %s\n", port)
	fmt.Printf("Your server can be found at http://localhost%v\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	// start webserver
	Startwebserve()
}
