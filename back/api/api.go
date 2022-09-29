package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type MovieApi struct {
	apiKey string
	url    string
}

var api = MovieApi{
	apiKey: os.Getenv("API_KEY"),
	url:    "http://www.omdbapi.com",
}

func GetMovieByTitleFromApi(title string) map[string]interface{} {
	resp, err := http.Get(api.url + "/?t=" + url.QueryEscape(title) + "&apikey=" + api.apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatalln(err)
	}

	return m
}

func DecodeElseEmptyString(m map[string]interface{}, key string) interface{} {
	val, ok := m[key]
	if ok && val != "N/A" {
		return val
	}

	return ""
}

func GetMovieStatsFromApi(title string) map[string]interface{} {
	m := GetMovieByTitleFromApi(title)

	stats := make(map[string]interface{})

	stats["BoxOffice"] = DecodeElseEmptyString(m, "BoxOffice")
	stats["Ratings"] = DecodeElseEmptyString(m, "Ratings")
	stats["Metascore"] = DecodeElseEmptyString(m, "Metascore")
	stats["imdbRating"] = DecodeElseEmptyString(m, "imdbRating")
	stats["imdbVotes"] = DecodeElseEmptyString(m, "imdbVotes")
	stats["Awards"] = DecodeElseEmptyString(m, "Awards")

	return stats
}
