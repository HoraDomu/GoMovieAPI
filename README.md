# GoMovieAPI
Testing beta 
Step 6 Full Function
func GetMovies(genre string) []Movie {
    genre = strings.ToLower(genre)         // ①
    genreID, ok := GenreIDs[genre]         // ②
    if !ok {                               // ③
        return []Movie{}                   // ④
    }

    url := fmt.Sprintf(                     // ⑤
        "https://api.themoviedb.org/3/discover/movie?api_key=%s&with_genres=%d&sort_by=release_date.desc",
        API_KEY, genreID,
    )

    resp, err := http.Get(url)             // ⑥
    if err != nil {                        // ⑦
        fmt.Println("Error fetching movies:", err) // ⑧
        return []Movie{}                   // ⑨
    }
    defer resp.Body.Close()                // ⑩

    body, err := ioutil.ReadAll(resp.Body) // ⑪
    if err != nil {                        // ⑫
        fmt.Println("Error reading response:", err) // ⑬
        return []Movie{}                   // ⑭
    }

    var tmdbResp TMDBResponse
    err = json.Unmarshal(body, &tmdbResp)  // ⑮
    if err != nil {                        // ⑯
        fmt.Println("Error parsing JSON:", err) // ⑰
        return []Movie{}                   // ⑱
    }

    movies := []Movie{}
    for i, m := range tmdbResp.Results {   // ⑲
        if i >= 5 {                        // ⑳
            break                           // ㉑
        }
        movies = append(movies, Movie{Title: fmt.Sprintf("%s (%s)", m.Title, m.ReleaseDate[:4])}) // ㉒
    }

    return movies                           // ㉓
}

Line by line explanation
① genre = strings.ToLower(genre)

Converts user input to all lowercase.

Example: "Action" → "action".

This helps us match the map key because Go maps are case-sensitive.

② genreID, ok := GenreIDs[genre]

GenreIDs is a map: map[string]int.

GenreIDs[genre] tries to look up the number for the genre.

ok is a boolean that is true if the key exists, false if it doesn’t.

Example:

GenreIDs := map[string]int{"action": 28}
id, ok := GenreIDs["action"] // id=28, ok=true
id, ok := GenreIDs["comedy"] // id=0, ok=false

③ if !ok {

!ok means “not ok”, i.e., the genre wasn’t found in the map.

! = “not” in Go (logical NOT).

④ return []Movie{}

If the genre doesn’t exist, we return an empty slice of Movie.

This prevents the program from crashing.

⑤ url := fmt.Sprintf(...)

fmt.Sprintf is like Python’s f-strings.

%s → insert a string (API_KEY)

%d → insert an integer (genreID)

Example:

API_KEY = "abc123"
genreID = 28
fmt.Sprintf("key=%s&id=%d", API_KEY, genreID)
// Output: "key=abc123&id=28"


We build the URL that TMDB expects.

⑥ resp, err := http.Get(url)

http.Get(url) makes a GET request to the internet.

resp = the response from the server

err = any error that happened while trying to fetch

⑦ if err != nil {

Checks if there was an error (e.g., no internet, invalid URL).

err != nil = “there is an error”.

⑧ fmt.Println("Error fetching movies:", err)

Prints a helpful error message.

⑨ return []Movie{}

If the request failed, return an empty slice to avoid crashing.

⑩ defer resp.Body.Close()

Ensures we close the connection after reading, even if an error happens later.

defer = “run this line when the function ends”.

⑪ body, err := ioutil.ReadAll(resp.Body)

Reads the entire HTTP response into a byte slice (body).

⑫ if err != nil {

Checks if reading the body failed.

⑬ fmt.Println("Error reading response:", err)

Prints an error if reading failed.

⑭ return []Movie{}

Return empty slice if reading fails.

⑮ err = json.Unmarshal(body, &tmdbResp)

json.Unmarshal = convert JSON into Go structs

&tmdbResp = we pass a pointer so the function can fill the struct with data

⑯ if err != nil {

Checks if parsing the JSON failed

⑰ fmt.Println("Error parsing JSON:", err)

Print a helpful message if the JSON is invalid

⑱ return []Movie{}

Return empty slice on JSON parse failure

⑲ for i, m := range tmdbResp.Results {

Loop over all movies returned by TMDB

i = index (0,1,2…)

m = single movie (TMDBMovie)

⑳ if i >= 5 {

Stop after top 5 movies

㉑ break

Exit the loop once we have 5 movies

㉒ movies = append(movies, Movie{Title: fmt.Sprintf("%s (%s)", m.Title, m.ReleaseDate[:4])})

Add the movie to our slice

fmt.Sprintf("%s (%s)", m.Title, m.ReleaseDate[:4]) → formats like: "Die Hard (1988)"

append() → adds a new item to the slice

㉓ return movies

Return the slice of top 5 movies to main.go

✅ Summary:

!ok → checks if the map key doesn’t exist

fmt.Sprintf → builds a string with variables inside

Each step checks for errors, handles them, then returns the data
