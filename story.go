package gypsy

import (
	"encoding/json"
	"io"
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

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
