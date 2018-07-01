package main

import "github.com/TrevorMellon/SteamWebCommon-Go"

type FilterGameType uint

const (
	FilterGamesUnplayed FilterGameType = iota
	FilterGameLessThanOneHour
	FilterGamesMoreThanOneHour
	FilterGamesMoreThanTwoHours
	FilterGamesMoreThanFourHours
	FilterGamesMoreThanEightHours
	FilterGamesMoreThanSixteenHours
)

func FilterGames(games []SteamCommon.SteamUserGame, filter FilterGameType) []SteamCommon.SteamUserGame {
	switch filter {
	case FilterGamesUnplayed:
		return FilterByGamesUnplayed(games)
	case FilterGameLessThanOneHour:
		return FilterGamesByLessThan(games, 60)
	case FilterGamesMoreThanOneHour:
		return FilterGamesByMinute(games, 60)
	case FilterGamesMoreThanTwoHours:
		return FilterGamesByMinute(games, 120)
	case FilterGamesMoreThanFourHours:
		return FilterGamesByMinute(games, 240)
	case FilterGamesMoreThanEightHours:
		return FilterGamesByMinute(games, 480)
	case FilterGamesMoreThanSixteenHours:
		return FilterGamesByMinute(games, 960)
	}
	return games
}

func FilterGamesByLessThan(games []SteamCommon.SteamUserGame, t uint64) []SteamCommon.SteamUserGame {
	var c int

	for _, v := range games {
		if v.PlaytimeForever < t && v.PlaytimeForever != 0 {
			c++
		}
	}

	a := make([]SteamCommon.SteamUserGame, c)
	var idx int

	for _, v := range games {
		if v.PlaytimeForever < t && v.PlaytimeForever != 0 {
			a[idx] = v
			idx++
		}
	}

	return a
}

func FilterByGamesUnplayed(games []SteamCommon.SteamUserGame) []SteamCommon.SteamUserGame {

	var c int

	for _, v := range games {
		if v.PlaytimeForever == 0 {
			c++
		}
	}

	a := make([]SteamCommon.SteamUserGame, c)
	var idx int

	for _, v := range games {
		if v.PlaytimeForever == 0 {
			a[idx] = v
			idx++
		}
	}

	return a
}

func FilterGamesByMinute(games []SteamCommon.SteamUserGame, t uint64) []SteamCommon.SteamUserGame {
	var c int

	for _, v := range games {
		if v.PlaytimeForever >= t {
			c++
		}
	}

	a := make([]SteamCommon.SteamUserGame, c)
	var idx int

	for _, v := range games {
		if v.PlaytimeForever >= t {
			a[idx] = v
			idx++
		}
	}

	return a
}
