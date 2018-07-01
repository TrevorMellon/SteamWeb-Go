package main

import (
	"sort"

	"github.com/TrevorMellon/SteamWebCommon-Go"
)

type SortFriendsType int

const (
	SortFriendsAsc SortFriendsType = iota
	SortFriendsDsc
)

type SortProfileFriendsAsc []SteamCommon.SteamUserProfile
type SortProfileFriendsDsc []SteamCommon.SteamUserProfile

func (a SortProfileFriendsAsc) Len() int {
	return len(a)
}
func (a SortProfileFriendsAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortProfileFriendsAsc) Less(i, j int) bool {
	aa := a[i].FriendSince.Unix()
	bb := a[j].FriendSince.Unix()
	if aa < bb {
		return true
	}
	if aa > bb {
		return false
	}

	return a[i].PersonaName < a[j].PersonaName
}

func (a SortProfileFriendsDsc) Len() int {
	return len(a)
}
func (a SortProfileFriendsDsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortProfileFriendsDsc) Less(i, j int) bool {
	aa := a[i].FriendSince.Unix()
	bb := a[j].FriendSince.Unix()
	if aa > bb {
		return true
	}
	if aa < bb {
		return false
	}

	return a[i].PersonaName > a[j].PersonaName
}

func SortProfileFriends(profile *SteamCommon.SteamUserProfile, sortDir SortFriendsType) {
	f := profile.Friends
	if sortDir == SortFriendsAsc {
		sort.Sort(SortProfileFriendsAsc(f))
	} else {
		sort.Sort(SortProfileFriendsDsc(f))
	}
	profile.Friends = f

}
