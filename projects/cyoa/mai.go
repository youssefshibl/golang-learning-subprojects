package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var DefaultFileName = "gopher.json"
var stories map[string]story

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	fileName := flag.String("file", DefaultFileName, "stories file name")
	flag.Parse()
	f, err := os.Open(*fileName)
	HandleError(err)
	defer f.Close()

	err = json.NewDecoder(f).Decode(&stories)
	HandleError(err)
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleRequest(stories))

	err = http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

func HandleRequest(stories map[string]story) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		defaultStory := "intro"
		storyName := r.URL.Query().Get("story")
		if storyName != "" {
			defaultStory = r.URL.Query().Get("story")
		}
		t, err := template.ParseFiles("page.html")
		if err != nil {
			HandleError(err)
		}
		err = t.Execute(w, stories[defaultStory])
		if err != nil {
			HandleError(err)
		}

	}
}
