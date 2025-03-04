package steamgameserverstats

import (
	. "github.com/assemblaj/purego-steamworks"
)

var (
	steamGameServerStats_init                   func() uintptr
	iSteamGameServerStats_RequestUserStats      func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall
	iSteamGameServerStats_GetUserStatInt32      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pData *int32) bool
	iSteamGameServerStats_GetUserStatFloat      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pData *float32) bool
	iSteamGameServerStats_GetUserAchievement    func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pbAchieved *bool) bool
	iSteamGameServerStats_SetUserStatInt32      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, nData int32) bool
	iSteamGameServerStats_SetUserStatFloat      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, fData float32) bool
	iSteamGameServerStats_UpdateUserAvgRateStat func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, flCountThisSession float32, dSessionLength float64) bool
	iSteamGameServerStats_SetUserAchievement    func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool
	iSteamGameServerStats_ClearUserAchievement  func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool
	iSteamGameServerStats_StoreUserStats        func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall
)

const (
	flatAPI_SteamGameServerStats                        = "SteamAPI_SteamGameServerStats_v001"
	flatAPI_ISteamGameServerStats_RequestUserStats      = "SteamAPI_ISteamGameServerStats_RequestUserStats"
	flatAPI_ISteamGameServerStats_GetUserStatInt32      = "SteamAPI_ISteamGameServerStats_GetUserStatInt32"
	flatAPI_ISteamGameServerStats_GetUserStatFloat      = "SteamAPI_ISteamGameServerStats_GetUserStatFloat"
	flatAPI_ISteamGameServerStats_GetUserAchievement    = "SteamAPI_ISteamGameServerStats_GetUserAchievement"
	flatAPI_ISteamGameServerStats_SetUserStatInt32      = "SteamAPI_ISteamGameServerStats_SetUserStatInt32"
	flatAPI_ISteamGameServerStats_SetUserStatFloat      = "SteamAPI_ISteamGameServerStats_SetUserStatFloat"
	flatAPI_ISteamGameServerStats_UpdateUserAvgRateStat = "SteamAPI_ISteamGameServerStats_UpdateUserAvgRateStat"
	flatAPI_ISteamGameServerStats_SetUserAchievement    = "SteamAPI_ISteamGameServerStats_SetUserAchievement"
	flatAPI_ISteamGameServerStats_ClearUserAchievement  = "SteamAPI_ISteamGameServerStats_ClearUserAchievement"
	flatAPI_ISteamGameServerStats_StoreUserStats        = "SteamAPI_ISteamGameServerStats_StoreUserStats"
)

var steamGameServerStats uintptr

func gameServerStats() uintptr {
	if steamGameServerStats == 0 {
		steamGameServerStats = steamGameServerStats_init()
		return steamGameServerStats
	}
	return steamGameServerStats
}

func RequestUserStats(userSteamID Uint64SteamID) CallResult[GSStatsReceived] {
	handle := iSteamGameServerStats_RequestUserStats(gameServerStats(), userSteamID)
	return CallResult[GSStatsReceived]{Handle: handle}
}

func GetUserStat(userSteamID Uint64SteamID, name string) (int32, bool) {
	var pData int32
	result := iSteamGameServerStats_GetUserStatInt32(gameServerStats(), userSteamID, name, &pData)
	return pData, result
}

func GetUserStatFloat(userSteamID Uint64SteamID, name string) (float32, bool) {
	var pData float32
	result := iSteamGameServerStats_GetUserStatFloat(gameServerStats(), userSteamID, name, &pData)
	return pData, result
}

func GetUserAchievement(userSteamID Uint64SteamID, name string) (bool, bool) {
	var pbAchieved bool
	result := iSteamGameServerStats_GetUserAchievement(gameServerStats(), userSteamID, name, &pbAchieved)
	return pbAchieved, result
}

func SetUserStat(userSteamID Uint64SteamID, name string, nData int32) bool {
	return iSteamGameServerStats_SetUserStatInt32(gameServerStats(), userSteamID, name, nData)
}

func SetUserStatFloat(userSteamID Uint64SteamID, name string, fData float32) bool {
	return iSteamGameServerStats_SetUserStatFloat(gameServerStats(), userSteamID, name, fData)
}

func UpdateUserAvgRateStat(userSteamID Uint64SteamID, name string, countThisSession float32, sessionLength float64) bool {
	return iSteamGameServerStats_UpdateUserAvgRateStat(gameServerStats(), userSteamID, name, countThisSession, sessionLength)
}

func SetUserAchievement(userSteamID Uint64SteamID, name string) bool {
	return iSteamGameServerStats_SetUserAchievement(gameServerStats(), userSteamID, name)
}

func ClearUserAchievement(userSteamID Uint64SteamID, name string) bool {
	return iSteamGameServerStats_ClearUserAchievement(gameServerStats(), userSteamID, name)
}

func StoreUserStats(userSteamID Uint64SteamID) CallResult[GSStatsStored] {
	handle := iSteamGameServerStats_StoreUserStats(gameServerStats(), userSteamID)
	return CallResult[GSStatsStored]{Handle: handle}
}
