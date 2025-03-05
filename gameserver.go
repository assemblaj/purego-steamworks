package puregosteamworks

var GameServerActive bool = false

const (
	flatAPI_SteamGameServer_Init       = "SteamInternal_GameServer_Init_V2"
	flatAPI_SteamGameServer_Shutdown   = "SteamGameServer_Shutdown"
	flatAPI_SteamGameServer_BSecure    = "SteamGameServer_BSecure"
	flatAPI_SteamGameServer_GetSteamID = "SteamGameServer_GetSteamID"
)

type EServerMode int32

const (
	ServerMode_Invalid                 EServerMode = 0 // DO NOT USE
	ServerMode_NoAuthentication        EServerMode = 1 // Don't authenticate user logins and don't list on the server list
	ServerMode_Authentication          EServerMode = 2 // Authenticate users, list on the server list, don't run VAC on clients that connect
	ServerMode_AuthenticationAndSecure EServerMode = 3 // Authenticate users, list on the server list and VAC protect clients
)

var (
	gameServerInit       func(IP uint32, gamePort uint16, queryPort uint16, serverMode EServerMode, versionString string) bool
	GameServerShutdown   func(*steamErrMsg) uintptr
	GameServerBSecure    func(uStructuredExceptionCode uint32, pvExceptionInfo []byte, uBuildID uint32)
	GameServerGetSteamID func(pchMsg string)
)

func GamseServerInit(IP uint32, gamePort uint16, queryPort uint16, serverMode EServerMode, versionString string) bool {
	success := gameServerInit(IP, gamePort, queryPort, serverMode, versionString)
	if success {
		GameServerActive = true
	}
	return success
}
