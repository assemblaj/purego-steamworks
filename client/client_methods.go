package client

import (
	"unsafe"

	. "github.com/assemblaj/purego-steamworks/network_types"

	. "github.com/assemblaj/purego-steamworks"
)

type EAccountType int32

const (
	EAccountTypeInvalid        EAccountType = 0
	EAccountTypeIndividual     EAccountType = 1
	EAccountTypeMultiseat      EAccountType = 2
	EAccountTypeGameServer     EAccountType = 3
	EAccountTypeAnonGameServer EAccountType = 4
	EAccountTypePending        EAccountType = 5
	EAccountTypeContentServer  EAccountType = 6
	EAccountTypeClan           EAccountType = 7
	EAccountTypeChat           EAccountType = 8
	EAccountTypeConsoleUser    EAccountType = 9
	EAccountTypeAnonUser       EAccountType = 10
	EAccountTypeMax            EAccountType = 11
)

var (
	iSteamClient_CreateSteamPipe             func() HSteamPipe
	iSteamClient_BReleaseSteamPipe           func(hSteamPipe HSteamPipe) bool
	iSteamClient_ConnectToGlobalUser         func(hSteamPipe HSteamPipe) HSteamUser
	iSteamClient_CreateLocalUser             func(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser
	iSteamClient_ReleaseUser                 func(hSteamPipe HSteamPipe, hUser HSteamUser)
	iSteamClient_GetISteamUser               func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameServer         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_SetLocalIPBinding           func(unIP *SteamIPAddress, usPort uint16)
	iSteamClient_GetISteamFriends            func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamUtils              func(hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamMatchmaking        func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamMatchmakingServers func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGenericInterface   func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer
	iSteamClient_GetISteamUserStats          func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameServerStats    func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamApps               func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamNetworking         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamRemoteStorage      func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamScreenshots        func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameSearch         func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetIPCCallCount             func() uint32
	iSteamClient_SetWarningMessageHook       func(pFunction SteamAPIWarningMessageHook)
	iSteamClient_BShutdownIfAllPipesClosed   func() bool
	iSteamClient_GetISteamUGC                func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamInventory          func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamInput              func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamParties            func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamRemotePlay         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
)

const (
	flatAPI_ISteamClient_CreateSteamPipe             = "SteamAPI_ISteamClient_CreateSteamPipe"
	flatAPI_ISteamClient_BReleaseSteamPipe           = "SteamAPI_ISteamClient_BReleaseSteamPipe"
	flatAPI_ISteamClient_ConnectToGlobalUser         = "SteamAPI_ISteamClient_ConnectToGlobalUser"
	flatAPI_ISteamClient_CreateLocalUser             = "SteamAPI_ISteamClient_CreateLocalUser"
	flatAPI_ISteamClient_ReleaseUser                 = "SteamAPI_ISteamClient_ReleaseUser"
	flatAPI_ISteamClient_GetISteamUser               = "SteamAPI_ISteamClient_GetISteamUser"
	flatAPI_ISteamClient_GetISteamGameServer         = "SteamAPI_ISteamClient_GetISteamGameServer"
	flatAPI_ISteamClient_SetLocalIPBinding           = "SteamAPI_ISteamClient_SetLocalIPBinding"
	flatAPI_ISteamClient_GetISteamFriends            = "SteamAPI_ISteamClient_GetISteamFriends"
	flatAPI_ISteamClient_GetISteamUtils              = "SteamAPI_ISteamClient_GetISteamUtils"
	flatAPI_ISteamClient_GetISteamMatchmaking        = "SteamAPI_ISteamClient_GetISteamMatchmaking"
	flatAPI_ISteamClient_GetISteamMatchmakingServers = "SteamAPI_ISteamClient_GetISteamMatchmakingServers"
	flatAPI_ISteamClient_GetISteamGenericInterface   = "SteamAPI_ISteamClient_GetISteamGenericInterface"
	flatAPI_ISteamClient_GetISteamUserStats          = "SteamAPI_ISteamClient_GetISteamUserStats"
	flatAPI_ISteamClient_GetISteamGameServerStats    = "SteamAPI_ISteamClient_GetISteamGameServerStats"
	flatAPI_ISteamClient_GetISteamApps               = "SteamAPI_ISteamClient_GetISteamApps"
	flatAPI_ISteamClient_GetISteamNetworking         = "SteamAPI_ISteamClient_GetISteamNetworking"
	flatAPI_ISteamClient_GetISteamRemoteStorage      = "SteamAPI_ISteamClient_GetISteamRemoteStorage"
	flatAPI_ISteamClient_GetISteamScreenshots        = "SteamAPI_ISteamClient_GetISteamScreenshots"
	flatAPI_ISteamClient_GetISteamGameSearch         = "SteamAPI_ISteamClient_GetISteamGameSearch"
	flatAPI_ISteamClient_GetIPCCallCount             = "SteamAPI_ISteamClient_GetIPCCallCount"
	flatAPI_ISteamClient_SetWarningMessageHook       = "SteamAPI_ISteamClient_SetWarningMessageHook"
	flatAPI_ISteamClient_BShutdownIfAllPipesClosed   = "SteamAPI_ISteamClient_BShutdownIfAllPipesClosed"
	flatAPI_ISteamClient_GetISteamUGC                = "SteamAPI_ISteamClient_GetISteamUGC"
	flatAPI_ISteamClient_GetISteamInventory          = "SteamAPI_ISteamClient_GetISteamInventory"
	flatAPI_ISteamClient_GetISteamInput              = "SteamAPI_ISteamClient_GetISteamInput"
	flatAPI_ISteamClient_GetISteamParties            = "SteamAPI_ISteamClient_GetISteamParties"
	flatAPI_ISteamClient_GetISteamRemotePlay         = "SteamAPI_ISteamClient_GetISteamRemotePlay"
)

func CreateSteamPipe() HSteamPipe {
	return iSteamClient_CreateSteamPipe()
}

func BReleaseSteamPipe(hSteamPipe HSteamPipe) bool {
	return iSteamClient_BReleaseSteamPipe(hSteamPipe)
}

func ConnectToGlobalUser(hSteamPipe HSteamPipe) HSteamUser {
	return iSteamClient_ConnectToGlobalUser(hSteamPipe)
}

func CreateLocalUser(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser {
	return iSteamClient_CreateLocalUser(phSteamPipe, eAccountType)
}

func ReleaseUser(hSteamPipe HSteamPipe, hUser HSteamUser) {
	iSteamClient_ReleaseUser(hSteamPipe, hUser)
}

func SetLocalIPBinding(unIP *SteamIPAddress, usPort uint16) {
	iSteamClient_SetLocalIPBinding(unIP, usPort)
}

func GetSteamGenericInterface(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer {
	return iSteamClient_GetISteamGenericInterface(hSteamUser, hSteamPipe, pchVersion)
}

func GetIPCCallCount() uint32 {
	return iSteamClient_GetIPCCallCount()
}

func SetWarningMessageHook(pFunction SteamAPIWarningMessageHook) {
	iSteamClient_SetWarningMessageHook(pFunction)
}

func BShutdownIfAllPipesClosed() bool {
	return iSteamClient_BShutdownIfAllPipesClosed()
}
