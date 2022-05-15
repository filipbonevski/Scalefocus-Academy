package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTopStoriesIds(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	router.Handle("/v0/topstories.json?print=pretty", handleTopStories(ids))
	mockServer := httptest.NewServer(router)

	want := ids[:10]

	ss := NewStoryService(mockServer.URL)
	got := ss.GetTopStoriesIds()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got: %v, want %v", got, want)
	}
}

func TestTopStories(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{10}
	stories := []Story{
		{
			Id:    23,
			Title: "Test Title",
			Score: 11,
		},
	}
	router.Handle("/v0/item", handleGetStory(stories))
	mockServer := httptest.NewServer(router)

	want := stories
	ss := NewStoryService(mockServer.URL)
	got := ss.GetStoriesByIds(ids)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got: %v, want %v", got, want)
	}
}

func handleTopStories(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ids)
	}
}

func handleGetStory(stories []Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		var wantedStory Story
		for _, story := range stories {
			if story.Id == id {
				wantedStory = story
				break
			}
		}
		json.NewEncoder(w).Encode(wantedStory)
	}
}
