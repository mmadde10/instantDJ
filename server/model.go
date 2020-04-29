package main

type UserLogin struct {
ID           string
Name         string
AccessToken  string
RefreshToken string
Email        string
}

type Image struct {
	Height int    `json: "height"`
	Url    string `json: "url"`
	Width  int    `json: "width"`
}

//Artist includes artist related fields
type Artist struct {
ExternalUrls interface{} `json: -`
Href         string      `json: "href"`
ID           string      `json: "id"`
Name         string      `json: "name"`
Type         string      `json: -`
URI          string      `json: "uri"`
}

// Album includes album related fields
type Album struct {
AlbumType            string      `json: "album_type"`
Artists              interface{} `json: -`
ExternalUrls         interface{} `json: -`
Href                 string      `json: "href"`
ID                   string      `json: "id"`
Name                 string      `json: "name"`
Images               []Image     `json: -`
releaseDate          string      `json: -`
releaseDatePrecision string      `json: -`
totalTracks          int         `json: -`
Type                 string      `json: -`
URI                  string      `json: "uri"`
}

type Track struct {
Album   Album    `json: "album"`
ID      string   `json: "id"`
Name    string   `json: "name"`
Artists []Artist `json: "artists"`
Href    string   `json: "href"`
}

type Item struct {
album        Album       `json: "album"`
artists      []Artist    `json: "artists"`
diskNumber   int         `json: "disc_number"`
durationMS   int         `json: "duration_ms"`
explicit     bool        `json: "explicit"`
ExternalUrls interface{} `json: -`
ExternalURLS interface{} `json: -`
href         string      `json: "href"`
ID           string      `json: "id"`
IsLocal      bool        `json: "is_local"`
IsPlayable   bool        `json: "is_playable"`
name         string      `json: "name"`
popularity   int         `json: -`
previewURL   string      `json: -`
uri          string      `json: "uri"`
}

type SearchResult struct {
artists []Artist
tracks  []Item
}
