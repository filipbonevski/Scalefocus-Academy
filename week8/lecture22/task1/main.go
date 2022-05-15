package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"
)

type Story struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

type StoryPageData struct {
	PageTitle string
	Stories   []Story
}

type StoryService struct {
	urlBase string
}

func NewStoryService(url string) *StoryService {
	return &StoryService{urlBase: url}
}

func (ss *StoryService) GetTopStoriesIds() []int {
	req, err := http.NewRequest("GET", ss.urlBase+"/v0/topstories.json?print=pretty", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var ids []int
	json.NewDecoder(resp.Body).Decode(&ids)
	return ids[:10]
}

func (ss *StoryService) GetStoriesByIds(ids []int) []Story {
	dataChannel := make(chan Story, len(ids))
	wg := sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			dataChannel <- ss.GetStoryById(id)
			defer wg.Done()
		}(id)
	}
	wg.Wait()
	close(dataChannel)

	result := make([]Story, 0, len(ids))
	for v := range dataChannel {
		result = append(result, v)
	}
	return result
}

func (ss *StoryService) GetStoryById(id int) Story {
	url := fmt.Sprintf("%s/v0/item/%d.json?print=pretty", ss.urlBase, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var story Story
	json.NewDecoder(resp.Body).Decode(&story)
	return story
}

func HandleTopStories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "only get methods are allowed", http.StatusBadRequest)
			return
		}

		ss := NewStoryService("https://hacker-news.firebaseio.com/")
		ids := ss.GetTopStoriesIds()
		storiesList := ss.GetStoriesByIds(ids)
		json.NewEncoder(w).Encode(storiesList)
	}
}

func ShowTopStories(tmpl template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := NewStoryService("https://hacker-news.firebaseio.com/")
		ids := ss.GetTopStoriesIds()
		storiesList := ss.GetStoriesByIds(ids)
		data := StoryPageData{
			PageTitle: "Top Stories",
			Stories:   storiesList,
		}
		tmpl.Execute(w, data)
	}
}

func main() {
	tmpl := template.Must(template.ParseFiles("topnews.html"))
	router := http.NewServeMux()
	router.Handle("/api/top", HandleTopStories())
	router.Handle("/top", ShowTopStories(*tmpl))
	log.Fatal(http.ListenAndServe(":9090", router))
}
