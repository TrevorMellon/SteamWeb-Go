package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	SteamMySql "github.com/TrevorMellon/SteamWebMySql-Go"
)

type Settings struct {
	UsingDatabase    bool
	SteamApiKey      string
	DatabaseSettings SteamMySql.DatabaseSettings
}

var appSettings Settings

func DefaultSettings() {
	appSettings.SteamApiKey = "ADD-API-KEY-HERE"
	appSettings.UsingDatabase = false
	appSettings.DatabaseSettings.Type = SteamMySql.DatabaseNone
}

func CheckSettings() {
	settingsfile := "./config.json"
	body, _ := ioutil.ReadFile(settingsfile)
	if len(body) == 0 {
		DefaultSettings()
		b, _ := json.Marshal(appSettings)
		ioutil.WriteFile(settingsfile, b, 0644)
	} else {
		var s Settings
		json.Unmarshal(body, &s)
		appSettings = s
		SteamMySql.DBSettings = appSettings.DatabaseSettings
		if s.SteamApiKey == "ADD-API-KEY-HERE" || s.SteamApiKey == "" {
			fmt.Println("Please enter a steam api key to config.json!")
		}
	}
}

func idwebhandler(w http.ResponseWriter, r *http.Request) {
	s :=
		`
<html>
<body style="background-color:black;">
	<form action="/steam/" method="get" style="color:white;"> 
		Steam64 ID:
		<input type="text" name="steamid" /> 
		<input type="submit" value="Submit" />
		<input type="hidden" name="sort" value="asc" />
	</form>
	<a href="https://steamid.io/lookup" target="_blank" style="color:#8888FF";>Lookup ids here</a>
</body>
</html> 
`

	w.Write([]byte(s))
}

func idparser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	steamid := r.Form.Get("steamid")

	url := "/steam/user/" + steamid + "/friends/"

	//w.WriteHeader(301)
	http.Redirect(w, r, url, 301)
}

func loadFile(resource string) ([]byte, error) {

	body, err := ioutil.ReadFile(resource)
	if err != nil {
		return nil, err
	}
	return body, err
}

func imageresourcehandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	s := "./images/" + vars["resource"]

	rr, _ := loadFile(s)
	w.Write(rr)
}

func jsresourcehandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	s := "./js/" + vars["resource"]
	//fmt.Println(r.RequestURI)
	rr, _ := loadFile(s)
	w.Write(rr)
}

func main() {
	r := mux.NewRouter()
	CheckSettings()

	r.HandleFunc("/steam/user/{steamid:[0-9]+}/friends/{filter}/", steamuserfriendhandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/friends/{filter}", steamuserfriendhandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/friends/", steamuserfriendhandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/friends", steamuserfriendhandler)

	r.HandleFunc("/steam/user/{steamid:[0-9]+}/games", steamgamehandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/games/", steamgamehandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/games/{filter:[A-Za-z0-9]+}", steamgamehandler)
	r.HandleFunc("/steam/user/{steamid:[0-9]+}/games/{filter:[A-Za-z0-9]+}/", steamgamehandler)
	r.HandleFunc("/steam/", idparser)
	r.HandleFunc("/steam", idparser)

	r.HandleFunc("/js/{resource}", jsresourcehandler)
	r.HandleFunc("/images/{resource}", imageresourcehandler)
	r.HandleFunc("/", idwebhandler)

	OpenIDFunc(r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8686", r))
	//http.Handle("/", r)
}
