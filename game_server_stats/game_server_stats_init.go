package steamgameserverstats

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamGameServerStats_init, SteamAPIDLL, flatAPI_SteamGameServerStats)
	purego.RegisterLibFunc(&iSteamGameServerStats_RequestUserStats, SteamAPIDLL, flatAPI_ISteamGameServerStats_RequestUserStats)
	purego.RegisterLibFunc(&iSteamGameServerStats_GetUserStatInt32, SteamAPIDLL, flatAPI_ISteamGameServerStats_GetUserStatInt32)
	purego.RegisterLibFunc(&iSteamGameServerStats_GetUserStatFloat, SteamAPIDLL, flatAPI_ISteamGameServerStats_GetUserStatFloat)
	purego.RegisterLibFunc(&iSteamGameServerStats_GetUserAchievement, SteamAPIDLL, flatAPI_ISteamGameServerStats_GetUserAchievement)
	purego.RegisterLibFunc(&iSteamGameServerStats_SetUserStatInt32, SteamAPIDLL, flatAPI_ISteamGameServerStats_SetUserStatInt32)
	purego.RegisterLibFunc(&iSteamGameServerStats_SetUserStatFloat, SteamAPIDLL, flatAPI_ISteamGameServerStats_SetUserStatFloat)
	purego.RegisterLibFunc(&iSteamGameServerStats_UpdateUserAvgRateStat, SteamAPIDLL, flatAPI_ISteamGameServerStats_UpdateUserAvgRateStat)
	purego.RegisterLibFunc(&iSteamGameServerStats_SetUserAchievement, SteamAPIDLL, flatAPI_ISteamGameServerStats_SetUserAchievement)
	purego.RegisterLibFunc(&iSteamGameServerStats_ClearUserAchievement, SteamAPIDLL, flatAPI_ISteamGameServerStats_ClearUserAchievement)
	purego.RegisterLibFunc(&iSteamGameServerStats_StoreUserStats, SteamAPIDLL, flatAPI_ISteamGameServerStats_StoreUserStats)
}
