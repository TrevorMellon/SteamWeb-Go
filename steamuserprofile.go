package main

import (
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/Philipp15b/go-steamapi"

	steamfriendssince "github.com/TrevorMellon/SteamWebFriendSince-Go"

	SteamApi "github.com/Philipp15b/go-steamapi"

	"github.com/TrevorMellon/SteamWebCommon-Go"
	"github.com/TrevorMellon/SteamWebMySql-Go"
)

func GetUserPlayerProfile(profile *SteamCommon.SteamUserProfile) {
	b3 := appSettings.UsingDatabase
	var b bool
	if b3 {
		b, p := SteamMySql.MysqlCheckUser(profile.SteamID)
		if b {
			*profile = p
		}
	} else {
		b = false
	}
	//*profile = p
	apikey := appSettings.SteamApiKey
	b2 := CheckUserPlayerProfileParsed(*profile)

	if !b || !b2 || !b3 {
		id := profile.SteamID
		myid := make([]uint64, 1)
		myid[0] = uint64(id)
		summ, _ := SteamApi.GetPlayerSummaries(myid, apikey)
		SteamPlayerSummaryToSteamUserProfile(profile, summ[0])

		f := steamfriendssince.GetMySteamFriends(apikey, uint64(id))
		ids := make([]uint64, len(f.Friends))
		for k, v := range f.Friends {
			ids[k] = v.SteamID
		}
		summ, _ = SteamApi.GetPlayerSummaries(ids, apikey)
		SteamPlayerSummariesToSteamUserProfiles(profile.Friends, summ)

		SteamFriendSinceToSteamUserProfile(profile, f.Friends)
		if b3 {
			SteamMySql.MysqlUpsertFriends(*profile)
			SteamMySql.MySqlUpsertUser(*profile)
			SteamMySql.MysqlGetFriends(profile)
		}

	} else {
		SteamMySql.MysqlGetFriends(profile)
	}
	CreateSteamUserProfileDataString(profile)
}

/*
Returns false if profile hasn't been parsed
*/
func CheckUserPlayerProfileParsed(profile SteamCommon.SteamUserProfile) bool {

	now := time.Now()
	then := profile.SummaryParsed

	diff := now.Sub(then)

	hf := diff.Minutes()

	if hf > 60 {
		return false
	}
	return true
}

func CreateSteamUserProfileDataString(profile *SteamCommon.SteamUserProfile) {
	//f := &profile.Friends
	for k, v := range profile.Friends {
		profile.Friends[k].FriendSinceStr = v.FriendSince.Format("2006 Jan 2")
		//log.Println(v.FriendSinceStr)
		//log.Println(v.FriendSince)
	}

}

func SteamPlayerSummariesToSteamUserProfiles(profiles []SteamCommon.SteamUserProfile, summaries []SteamApi.PlayerSummary) {
	for k, _ := range profiles {
		SteamPlayerSummaryToSteamUserProfile(&profiles[k], summaries[k])
	}
}

func SteamPlayerSummaryToSteamUserProfile(profile *SteamCommon.SteamUserProfile, summary SteamApi.PlayerSummary) {
	name := summary.PersonaName

	if b := utf8.ValidString(name); !b {
		r := []rune(name)
		name = ""
		for _, v := range r {
			name += strconv.QuoteRuneToASCII(v)
		}
	}

	real := summary.RealName

	if b := utf8.ValidString(real); !b {
		r := []rune(real)
		real = ""
		for _, v := range r {
			real += strconv.QuoteRuneToASCII(v)
		}
	}

	profile.SteamID = SteamCommon.SteamID(summary.SteamID)
	profile.PersonaName = name
	profile.DisplayPersona = summary.PersonaName
	profile.Realname = real
	profile.DisplayRealname = summary.RealName
	profile.Avatar.Small = summary.SmallAvatarURL
	profile.Avatar.Medium = summary.MediumAvatarURL
	profile.Avatar.Large = summary.LargeAvatarURL
	profile.Url = summary.ProfileURL

	switch summary.CommunityVisibilityState {
	case steamapi.Public:
		profile.ProfileType = SteamCommon.ProfilePublic
		break
	case steamapi.Private:
		profile.ProfileType = SteamCommon.ProfilePrivate
		break
	case steamapi.FriendsOnly:
		profile.ProfileType = SteamCommon.ProfileSemi
		break
	default:
		profile.ProfileType = SteamCommon.ProfileUnknown
		break
	}

	switch summary.PersonaState {
	case steamapi.Online:
		profile.Status = SteamCommon.StatusOnline
		break
	case steamapi.Offline:
		profile.Status = SteamCommon.StatusOffline
		break
	case steamapi.Away:
		profile.Status = SteamCommon.StatusAway
		break
	case steamapi.Busy:
		profile.Status = SteamCommon.StatusBusy
		break
	default:
		profile.Status = SteamCommon.StatusUnknown
	}
}

func SteamFriendSinceToSteamUserProfile(profile *SteamCommon.SteamUserProfile, friends []steamfriendssince.FriendSince) {
	profile.Friends = make([]SteamCommon.SteamUserProfile, len(friends))
	f := profile.Friends
	for k, v := range friends {
		f[k].PersonaName = v.PersonaName
		f[k].DisplayPersona = v.PersonaName
		f[k].Realname = v.RealName
		f[k].DisplayRealname = v.RealName
		f[k].Avatar.Small = v.Summary.SmallAvatarURL
		f[k].Avatar.Medium = v.Summary.MediumAvatarURL
		f[k].Avatar.Large = v.Summary.LargeAvatarURL
		f[k].SteamID = SteamCommon.SteamID(v.SteamID)
		f[k].FriendSince = time.Unix(v.FriendSince.UnixTime, 0)
		f[k].Url = v.Summary.ProfileURL
		switch v.Summary.CommunityVisibilityState {
		case steamapi.Public:
			f[k].ProfileType = SteamCommon.ProfilePublic
			break
		case steamapi.Private:
			f[k].ProfileType = SteamCommon.ProfilePrivate
			break
		case steamapi.FriendsOnly:
			f[k].ProfileType = SteamCommon.ProfileSemi
			break
		default:
			f[k].ProfileType = SteamCommon.ProfileUnknown
			break
		}

		switch v.Summary.PersonaState {
		case steamapi.Online:
			f[k].Status = SteamCommon.StatusOnline
			break
		case steamapi.Offline:
			f[k].Status = SteamCommon.StatusOffline
			break
		case steamapi.Away:
			f[k].Status = SteamCommon.StatusAway
			break
		case steamapi.Busy:
			f[k].Status = SteamCommon.StatusBusy
			break
		default:
			f[k].Status = SteamCommon.StatusUnknown
			break
		}
	}
	profile.Friends = f
}
