package main

import (
	"github.com/TrevorMellon/SteamWebCommon-Go"
)

type FilterType int

const (
	FilterNone FilterType = iota
	FilterOnline
	FilterOffline
	FilterBusy
	FilterAway
	FilterLookingToTrade
	FilterLookingToGame
	FilterPrivate
	FilterPublic
	FilterSemiPrivate
)

func FilterSteamFriends(profile SteamCommon.SteamUserProfile, filter FilterType) []SteamCommon.SteamUserProfile {

	switch filter {
	case FilterOnline:
		return FilterSteamFriendsByOnlineStatus(profile)
	case FilterOffline:
		return FilterSteamFriendsByOfflineStatus(profile)
	case FilterAway:
		return FilterSteamFriendsByAwayStatus(profile)
	case FilterBusy:
		return FilterSteamFriendsByBusyStatus(profile)
	case FilterLookingToTrade:
		return FilterSteamFriendsByLookingToTradeStatus(profile)
	case FilterLookingToGame:
		return FilterSteamFriendsByLookingToGameStatus(profile)
	case FilterPublic:
		return FilterSteamFriendsByPublicProfile(profile)
	case FilterPrivate:
		return FilterSteamFriendsByPrivateProfile(profile)
	case FilterSemiPrivate:
		return FilterSteamFriendsBySemiPrivateProfile(profile)
	}
	return profile.Friends
}

func FilterSteamFriendsByPublicProfile(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfilePublic {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfilePublic {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByPrivateProfile(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfilePrivate {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfilePrivate {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsBySemiPrivateProfile(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfileSemi {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.ProfileType == SteamCommon.ProfileSemi {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByOnlineStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusOnline {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusOnline {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByOfflineStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusOffline {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusOffline {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByAwayStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusAway {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusAway {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByBusyStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusBusy {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusBusy {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByLookingToTradeStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusLookingToTrade {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusLookingToTrade {
			ret[idx] = v
			idx++
		}
	}

	return ret
}

func FilterSteamFriendsByLookingToGameStatus(profile SteamCommon.SteamUserProfile) []SteamCommon.SteamUserProfile {
	var onlinecount int
	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusLookingToGame {
			onlinecount++
		}
	}

	ret := make([]SteamCommon.SteamUserProfile, onlinecount)

	var idx int

	for _, v := range profile.Friends {
		if v.Status == SteamCommon.StatusLookingToGame {
			ret[idx] = v
			idx++
		}
	}

	return ret
}
