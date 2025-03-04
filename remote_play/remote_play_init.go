package remoteplay

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamRemotePlay_init, SteamAPIDLL, flatAPI_SteamRemotePlay)
	purego.RegisterLibFunc(&iSteamRemotePlay_GetSessionCount, SteamAPIDLL, flatAPI_ISteamRemotePlay_GetSessionCount)
	purego.RegisterLibFunc(&iSteamRemotePlay_GetSessionID, SteamAPIDLL, flatAPI_ISteamRemotePlay_GetSessionID)
	purego.RegisterLibFunc(&iSteamRemotePlay_GetSessionSteamID, SteamAPIDLL, flatAPI_ISteamRemotePlay_GetSessionSteamID)
	purego.RegisterLibFunc(&iSteamRemotePlay_GetSessionClientName, SteamAPIDLL, flatAPI_ISteamRemotePlay_GetSessionClientName)
	purego.RegisterLibFunc(&iSteamRemotePlay_GetSessionClientFormFactor, SteamAPIDLL, flatAPI_ISteamRemotePlay_GetSessionClientFormFactor)
	purego.RegisterLibFunc(&iSteamRemotePlay_BGetSessionClientResolution, SteamAPIDLL, flatAPI_ISteamRemotePlay_BGetSessionClientResolution)
	purego.RegisterLibFunc(&iSteamRemotePlay_BStartRemotePlayTogether, SteamAPIDLL, flatAPI_ISteamRemotePlay_BStartRemotePlayTogether)
	purego.RegisterLibFunc(&iSteamRemotePlay_BSendRemotePlayTogetherInvite, SteamAPIDLL, flatAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite)
}
