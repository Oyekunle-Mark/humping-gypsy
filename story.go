package gypsy

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var defaultHanldlerTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Adventure Game</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

func init() {
	tmpl = template.Must(template.New("").Parse(defaultHanldlerTmpl))
}

var tmpl *template.Template

// NewHandler returns a type that implements htt.Handler
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tmpl.Execute(w, chapter)

		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something broke...", http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "Cannot find chapter.", http.StatusNotFound)
}

// Story the story type
type Story map[string]Chapter

// Chapter structure of the chapter
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option structure of option
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// JSONStory decodes the io reader and returns a story map
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
