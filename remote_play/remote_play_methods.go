package remoteplay

import . "github.com/assemblaj/purego-steamworks"

const (
	flatAPI_SteamRemotePlay                                = "SteamAPI_SteamRemotePlay_v002"
	flatAPI_ISteamRemotePlay_GetSessionCount               = "SteamAPI_ISteamRemotePlay_GetSessionCount"
	flatAPI_ISteamRemotePlay_GetSessionID                  = "SteamAPI_ISteamRemotePlay_GetSessionID"
	flatAPI_ISteamRemotePlay_GetSessionSteamID             = "SteamAPI_ISteamRemotePlay_GetSessionSteamID"
	flatAPI_ISteamRemotePlay_GetSessionClientName          = "SteamAPI_ISteamRemotePlay_GetSessionClientName"
	flatAPI_ISteamRemotePlay_GetSessionClientFormFactor    = "SteamAPI_ISteamRemotePlay_GetSessionClientFormFactor"
	flatAPI_ISteamRemotePlay_BGetSessionClientResolution   = "SteamAPI_ISteamRemotePlay_BGetSessionClientResolution"
	flatAPI_ISteamRemotePlay_BStartRemotePlayTogether      = "SteamAPI_ISteamRemotePlay_BStartRemotePlayTogether"
	flatAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite = "SteamAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite"
)

type RemotePlaySessionID uint

type ESteamDeviceFormFactor int32

const (
	ESteamDeviceFormFactorUnknown   ESteamDeviceFormFactor = 0
	ESteamDeviceFormFactorPhone     ESteamDeviceFormFactor = 1
	ESteamDeviceFormFactorTablet    ESteamDeviceFormFactor = 2
	ESteamDeviceFormFactorComputer  ESteamDeviceFormFactor = 3
	ESteamDeviceFormFactorTV        ESteamDeviceFormFactor = 4
	ESteamDeviceFormFactorVRHeadset ESteamDeviceFormFactor = 5
)

var (
	steamRemotePlay_init                           func() uintptr
	iSteamRemotePlay_GetSessionCount               func(steamRemotePlay uintptr) uint32
	iSteamRemotePlay_GetSessionID                  func(steamRemotePlay uintptr, iSessionIndex int32) RemotePlaySessionID
	iSteamRemotePlay_GetSessionSteamID             func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) Uint64SteamID
	iSteamRemotePlay_GetSessionClientName          func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) string
	iSteamRemotePlay_GetSessionClientFormFactor    func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) ESteamDeviceFormFactor
	iSteamRemotePlay_BGetSessionClientResolution   func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID, pnResolutionX *int32, pnResolutionY *int32) bool
	iSteamRemotePlay_BStartRemotePlayTogether      func(steamRemotePlay uintptr, bShowOverlay bool) bool
	iSteamRemotePlay_BSendRemotePlayTogetherInvite func(steamRemotePlay uintptr, steamIDFriend Uint64SteamID) bool
)

var steamRemotePlay uintptr

func remotePlay() uintptr {
	if steamRemotePlay == 0 {
		steamRemotePlay = steamRemotePlay_init()
		return steamRemotePlay
	}
	return steamRemotePlay
}

func GetSessionCount() uint32 {
	return iSteamRemotePlay_GetSessionCount(remotePlay())
}

func GetSessionID(SessionIndex int32) RemotePlaySessionID {
	return iSteamRemotePlay_GetSessionID(remotePlay(), SessionIndex)
}

func GetSessionSteamID(SessionID RemotePlaySessionID) Uint64SteamID {
	return iSteamRemotePlay_GetSessionSteamID(remotePlay(), SessionID)
}

func GetSessionClientName(SessionID RemotePlaySessionID) string {
	return iSteamRemotePlay_GetSessionClientName(remotePlay(), SessionID)
}

func GetSessionClientFormFactor(SessionID RemotePlaySessionID) ESteamDeviceFormFactor {
	return iSteamRemotePlay_GetSessionClientFormFactor(remotePlay(), SessionID)
}

func BGetSessionClientResolution(SessionID RemotePlaySessionID) (int32, int32, bool) {
	var pnResolutionX, pnResolutionY int32
	result := iSteamRemotePlay_BGetSessionClientResolution(remotePlay(), SessionID, &pnResolutionX, &pnResolutionY)
	return pnResolutionX, pnResolutionY, result
}

func BStartRemotePlayTogether(ShowOverlay bool) bool {
	return iSteamRemotePlay_BStartRemotePlayTogether(remotePlay(), ShowOverlay)
}

func BSendRemotePlayTogetherInvite(friendSteamID Uint64SteamID) bool {
	return iSteamRemotePlay_BSendRemotePlayTogetherInvite(remotePlay(), friendSteamID)
}
