package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/TrevorMellon/SteamWebCommon-Go"

	"github.com/gorilla/mux"
)

type SteamUserFriendPage struct {
	Page    WebPage
	Profile SteamCommon.SteamUserProfile
}

type WebPage struct {
	Title       string
	Image       string
	Description string
}

type SteamGamePage struct {
	Page    WebPage
	Profile SteamCommon.SteamUserProfile
	Games   []SteamCommon.SteamUserGame
}

/*type myprofile struct {
	steam                    SteamUserFriend
	page                     WebPage
	FriendsOnlineStatusColor []string
}*/

/*type SteamUserFriend struct {
	SteamID     uint64
	PersonaName string
	Image       string
	Url         string
	Friends     []SteamFriendsSince.FriendSince

	FriendStats SteamFriendsSince.FriendsStatistics
}*/

//type SortDirection int

/*
const (
	SortUnset SortDirection = iota
	SortAsc   SortDirection = iota
	SortDesc  SortDirection = iota
)*/

func fromSortString(in string) SortFriendsType {
	switch in {
	case "asc":
		return SortFriendsAsc
	case "dsc":
		return SortFriendsDsc
	case "desc":
		return SortFriendsDsc
	default:
		return SortFriendsAsc
	}
}

func fromFilterString(in string) FilterType {
	switch in {
	case "online":
		return FilterOnline
	case "offline":
		return FilterOffline
	case "away":
		return FilterAway
	case "busy":
		return FilterBusy
	case "private":
		return FilterPrivate
	case "public":
		return FilterPublic
	case "friendsonly":
		return FilterSemiPrivate
	default:
		return FilterNone
	}
}

func setBorderColor(profile *SteamCommon.SteamUserProfile) {

	for k, v := range profile.Friends {
		//summary := v.Summary
		//s := summary.PersonaState

		s := v.Status

		var str string

		switch s {
		case SteamCommon.StatusOnline:
			str = "#0000FF"
			break
		case SteamCommon.StatusBusy:
			str = "#FFee42" //yellow
			break
		/*case SteamApi.Snooze:
		profile.FriendsOnlineStatusColor[k] = "#f4ee42" //yellow
		break*/
		case SteamCommon.StatusAway:
			str = "#f4ee42" //yellow
			break
		default:
			str = "#020202"
			break
		}

		if v.ProfileType == SteamCommon.ProfilePrivate {
			str = "#FF0000"
		}
		if v.ProfileType == SteamCommon.ProfileSemi {
			str = "#a87d05"
		}

		profile.Friends[k].StatusColor = str
	}
}

func checkSteamUserFilter(r *http.Request) (uint64, FilterType) {
	var reti uint64
	reti = 0

	vars := mux.Vars(r)

	if len(vars["filter"]) == 0 {
		return 0, FilterNone
	}

	reti, err := strconv.ParseUint(vars["steamid"], 10, 64)

	if err != nil {
		return 0, FilterNone
	}

	filter := fromFilterString(vars["filter"])

	return reti, filter
}

func SteamFriendsPage(w http.ResponseWriter, r *http.Request) {
	//apikey := appSettings.SteamApiKey

	vars := mux.Vars(r)

	steamid, err := strconv.ParseUint(vars["steamid"], 10, 64)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var page SteamUserFriendPage

	var profile SteamCommon.SteamUserProfile
	profile.SteamID = SteamCommon.SteamID(steamid)

	GetUserPlayerProfile(&profile)

	r.ParseForm()

	sortdir := fromSortString(r.Form.Get("sort"))

	switch sortdir {
	case SortFriendsAsc:
		SortProfileFriends(&profile, SortFriendsAsc)
		break
	case SortFriendsDsc:
		SortProfileFriends(&profile, SortFriendsDsc)
		break
	default:
		break
	}

	page.Profile = profile

	_, filter := checkSteamUserFilter(r)

	if filter != FilterNone {
		page.Profile.Friends = FilterSteamFriends(page.Profile, filter)
	}

	setBorderColor(&page.Profile)

	switch filter {
	case FilterNone:
		page.Page.Title = page.Profile.DisplayPersona + "'s Friends List"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every current steam friend of " + page.Profile.DisplayPersona
		break
	case FilterOnline:
		page.Page.Title = page.Profile.DisplayPersona + "'s Online Friends List"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that's currently online."
		break
	case FilterOffline:
		page.Page.Title = page.Profile.DisplayPersona + "'s Offline Friends List"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that's currently offline."
		break
	case FilterAway:
		page.Page.Title = page.Profile.DisplayPersona + "'s Friends List (Away)"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that's currently away from their keyboard."
		break
	case FilterBusy:
		page.Page.Title = page.Profile.DisplayPersona + "'s Busy Friends List"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that are currently busy."
		break
	case FilterPublic:
		page.Page.Title = page.Profile.DisplayPersona + "'s Friends List (Public)"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that shares their profile with the public."
		break
	case FilterPrivate:
		page.Page.Title = page.Profile.DisplayPersona + "'s Friends List (Private)"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that currently has a private only profile."
		break
	case FilterSemiPrivate:
		page.Page.Title = page.Profile.DisplayPersona + "'s Friends List (Friends Only)"
		page.Page.Image = page.Profile.Avatar.Large
		page.Page.Description = "A listing of every steam friend of " + page.Profile.DisplayPersona + " that only shares their profile with their own friends."
		break
	}
	//}

	temp, err := ioutil.ReadFile("./htmltemplates/steamfriends.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	t := template.New("Steam")

	t, err = t.Parse(string(temp))

	if err != nil {
		log.Fatal(err)
		return
	}
	t.Execute(w, &page)
}

func steamuserfriendhandler(w http.ResponseWriter, r *http.Request) {
	SteamFriendsPage(w, r)
}

func steamgamehandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	steamid, err := strconv.ParseUint(vars["steamid"], 10, 64)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var gamepage SteamGamePage
	gamepage.Profile.SteamID = SteamCommon.SteamID(steamid)

	gamelist := getPlayerGames(SteamCommon.SteamID(steamid), appSettings.SteamApiKey)

	gamepage.Games = gamelist.Games
	GetUserPlayerProfile(&gamepage.Profile)

	name := gamepage.Profile.DisplayPersona

	gamepage.Page.Title = name + "'s Game List"
	gamepage.Page.Description = "A listing of every game " + name + "has played since joining steam"
	gamepage.Page.Image = gamepage.Profile.Avatar.Large

	filter := vars["filter"]
	filter = strings.ToLower(filter)
	switch filter {
	case "unplayed":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesUnplayed)
		gamepage.Page.Title = name + "'s Unplayed Game List"
		gamepage.Page.Description = "A listing of every game " + name + " has played since joining steam"
		break
	case "60min":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesUnplayed)
		gamepage.Page.Title = name + "'s Game List of partially played games"
		gamepage.Page.Description = "A listing of every game " + name + " has started but didn't really play played since joining steam"
		break
	case "1hour":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesMoreThanOneHour)
		gamepage.Page.Title = name + "'s Game List (> 1 hours play)"
		gamepage.Page.Description = "A listing of every game " + name + " has played more than an hour since joining steam"
		break
	case "2hour":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesMoreThanTwoHours)
		gamepage.Page.Title = name + "'s Game List (> 2 hours play)"
		gamepage.Page.Description = "A listing of every game " + name + " has played more than two hours since joining steam"
		break
	case "4hour":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesMoreThanFourHours)
		gamepage.Page.Title = name + "'s Game List (> 4 hours play)"
		gamepage.Page.Description = "A listing of every game " + name + " has played more than four hours since joining steam"
		break
	case "8hour":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesMoreThanEightHours)
		gamepage.Page.Title = name + "'s Game List (> 8 hours play)"
		gamepage.Page.Description = "A listing of every game " + name + " has played more than eight hours since joining steam"
		break
	case "16hour":
		gamepage.Games = FilterGames(gamepage.Games, FilterGamesMoreThanSixteenHours)
		gamepage.Page.Title = name + "'s Game List (> 16 hours play)"
		gamepage.Page.Description = "A listing of every game " + name + " has played more than sixteen hours since joining steam"
		break
	}

	temp, err := ioutil.ReadFile("./htmltemplates/steamgames.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	t := template.New("Steam")

	t, err = t.Parse(string(temp))

	if err != nil {
		log.Fatal(err)
		return
	}
	t.Execute(w, &gamepage)
}

func steamhandler(w http.ResponseWriter, r *http.Request) {
	/*if i := checkUserUrl(r.RequestURI); i > 0 {
		SteamFriendsPage(w, r)
		return
	}*/
	/*if i, _ := checkSteamUserFilter(r); i > 0 {
		SteamFriendsPage(w, r)
		return
	}*/
	r.ParseForm()
	s := r.Form.Get("steamid")
	if s == "" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	//steamid, _ := strconv.ParseUint(s, 10, 64)

	sortdir := r.Form.Get("sort")

	if sortdir == "" {
		w.Header().Set("Location", "/steam/user/"+s+"/friends")
	} else {
		w.Header().Set("Location", "/steam/user/"+s+"/friends/?sort="+sortdir)
	}
	w.WriteHeader(301)

	//SteamFriendsPage(w, r, steamid)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		w.Write([]byte("Resource Not Found!!"))
	}
}
