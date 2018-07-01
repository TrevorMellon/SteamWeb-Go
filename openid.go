package main

import (
	
	"net/http"


	"github.com/gorilla/mux"
	"github.com/solovev/steam_go"
)

func OpenIDAuthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	opId := steam_go.NewOpenId(r)
	switch opId.Mode() {
	case "":
		http.Redirect(w, r, opId.AuthUrl(), 301)
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
	default:
		steamId, err := opId.ValidateAndGetId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Do whatever you want with steam id
		//w.Write([]byte(steamId))
		url := "/steam/user/" + steamId + "/friends"
		http.Redirect(w, r, url, 301)
	}
}



func OpenIDFunc(r *mux.Router) {
	a := http.HandlerFunc(OpenIDAuthenticatedHandler)
	r.Handle("/user", a)
}
