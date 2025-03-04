package client

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&iSteamClient_CreateSteamPipe, SteamAPIDLL, flatAPI_ISteamClient_CreateSteamPipe)
	purego.RegisterLibFunc(&iSteamClient_BReleaseSteamPipe, SteamAPIDLL, flatAPI_ISteamClient_BReleaseSteamPipe)
	purego.RegisterLibFunc(&iSteamClient_ConnectToGlobalUser, SteamAPIDLL, flatAPI_ISteamClient_ConnectToGlobalUser)
	purego.RegisterLibFunc(&iSteamClient_CreateLocalUser, SteamAPIDLL, flatAPI_ISteamClient_CreateLocalUser)
	purego.RegisterLibFunc(&iSteamClient_ReleaseUser, SteamAPIDLL, flatAPI_ISteamClient_ReleaseUser)
	purego.RegisterLibFunc(&iSteamClient_GetISteamUser, SteamAPIDLL, flatAPI_ISteamClient_GetISteamUser)
	purego.RegisterLibFunc(&iSteamClient_GetISteamGameServer, SteamAPIDLL, flatAPI_ISteamClient_GetISteamGameServer)
	purego.RegisterLibFunc(&iSteamClient_SetLocalIPBinding, SteamAPIDLL, flatAPI_ISteamClient_SetLocalIPBinding)
	purego.RegisterLibFunc(&iSteamClient_GetISteamFriends, SteamAPIDLL, flatAPI_ISteamClient_GetISteamFriends)
	purego.RegisterLibFunc(&iSteamClient_GetISteamUtils, SteamAPIDLL, flatAPI_ISteamClient_GetISteamUtils)
	purego.RegisterLibFunc(&iSteamClient_GetISteamMatchmaking, SteamAPIDLL, flatAPI_ISteamClient_GetISteamMatchmaking)
	purego.RegisterLibFunc(&iSteamClient_GetISteamMatchmakingServers, SteamAPIDLL, flatAPI_ISteamClient_GetISteamMatchmakingServers)
	purego.RegisterLibFunc(&iSteamClient_GetISteamGenericInterface, SteamAPIDLL, flatAPI_ISteamClient_GetISteamGenericInterface)
	purego.RegisterLibFunc(&iSteamClient_GetISteamUserStats, SteamAPIDLL, flatAPI_ISteamClient_GetISteamUserStats)
	purego.RegisterLibFunc(&iSteamClient_GetISteamGameServerStats, SteamAPIDLL, flatAPI_ISteamClient_GetISteamGameServerStats)
	purego.RegisterLibFunc(&iSteamClient_GetISteamApps, SteamAPIDLL, flatAPI_ISteamClient_GetISteamApps)
	purego.RegisterLibFunc(&iSteamClient_GetISteamNetworking, SteamAPIDLL, flatAPI_ISteamClient_GetISteamNetworking)
	purego.RegisterLibFunc(&iSteamClient_GetISteamRemoteStorage, SteamAPIDLL, flatAPI_ISteamClient_GetISteamRemoteStorage)
	purego.RegisterLibFunc(&iSteamClient_GetISteamScreenshots, SteamAPIDLL, flatAPI_ISteamClient_GetISteamScreenshots)
	purego.RegisterLibFunc(&iSteamClient_GetISteamGameSearch, SteamAPIDLL, flatAPI_ISteamClient_GetISteamGameSearch)
	purego.RegisterLibFunc(&iSteamClient_GetIPCCallCount, SteamAPIDLL, flatAPI_ISteamClient_GetIPCCallCount)
	purego.RegisterLibFunc(&iSteamClient_SetWarningMessageHook, SteamAPIDLL, flatAPI_ISteamClient_SetWarningMessageHook)
	purego.RegisterLibFunc(&iSteamClient_BShutdownIfAllPipesClosed, SteamAPIDLL, flatAPI_ISteamClient_BShutdownIfAllPipesClosed)
	purego.RegisterLibFunc(&iSteamClient_GetISteamUGC, SteamAPIDLL, flatAPI_ISteamClient_GetISteamUGC)
	purego.RegisterLibFunc(&iSteamClient_GetISteamInventory, SteamAPIDLL, flatAPI_ISteamClient_GetISteamInventory)
	purego.RegisterLibFunc(&iSteamClient_GetISteamInput, SteamAPIDLL, flatAPI_ISteamClient_GetISteamInput)
	purego.RegisterLibFunc(&iSteamClient_GetISteamParties, SteamAPIDLL, flatAPI_ISteamClient_GetISteamParties)
	purego.RegisterLibFunc(&iSteamClient_GetISteamRemotePlay, SteamAPIDLL, flatAPI_ISteamClient_GetISteamRemotePlay)
}
