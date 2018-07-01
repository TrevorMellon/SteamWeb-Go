package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TrevorMellon/SteamWebCommon-Go"
)

type SteamUserGameListResponse struct {
	Response SteamUserGameList `json:"response"`
}

type SteamUserGameList struct {
	GameCount int                         `json:"game_count"`
	Games     []SteamCommon.SteamUserGame `json:"games"`
}

func getPlayerGames(steamid SteamCommon.SteamID, appid string) SteamUserGameList {
	baseurl := "http://api.steampowered.com"
	url := baseurl + "/IPlayerService/GetOwnedGames/v0001/"
	id := uint64(steamid)
	idstr := strconv.FormatUint(id, 10)

	url += "?key=" + appid
	url += "&steamid=" + idstr
	url += "&include_played_free_games=1"
	url += "&include_appinfo=1"

	r, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	d := json.NewDecoder(r.Body)

	var gl SteamUserGameListResponse
	gl.Response.Games = make([]SteamCommon.SteamUserGame, 0)
	err = d.Decode(&gl)
	//err = json.Unmarshal(b, &gl)

	if err != nil {
		log.Println(err)
	}
	return gl.Response
}
