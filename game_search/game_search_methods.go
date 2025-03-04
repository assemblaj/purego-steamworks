package gamesearch

import . "github.com/assemblaj/purego-steamworks"

const (
	flatAPI_SteamGameSearch                              = "SteamAPI_SteamGameSearch_v001"
	flatAPI_ISteamGameSearch_AddGameSearchParams         = "SteamAPI_ISteamGameSearch_AddGameSearchParams"
	flatAPI_ISteamGameSearch_SearchForGameWithLobby      = "SteamAPI_ISteamGameSearch_SearchForGameWithLobby"
	flatAPI_ISteamGameSearch_SearchForGameSolo           = "SteamAPI_ISteamGameSearch_SearchForGameSolo"
	flatAPI_ISteamGameSearch_AcceptGame                  = "SteamAPI_ISteamGameSearch_AcceptGame"
	flatAPI_ISteamGameSearch_DeclineGame                 = "SteamAPI_ISteamGameSearch_DeclineGame"
	flatAPI_ISteamGameSearch_RetrieveConnectionDetails   = "SteamAPI_ISteamGameSearch_RetrieveConnectionDetails"
	flatAPI_ISteamGameSearch_EndGameSearch               = "SteamAPI_ISteamGameSearch_EndGameSearch"
	flatAPI_ISteamGameSearch_SetGameHostParams           = "SteamAPI_ISteamGameSearch_SetGameHostParams"
	flatAPI_ISteamGameSearch_SetConnectionDetails        = "SteamAPI_ISteamGameSearch_SetConnectionDetails"
	flatAPI_ISteamGameSearch_RequestPlayersForGame       = "SteamAPI_ISteamGameSearch_RequestPlayersForGame"
	flatAPI_ISteamGameSearch_HostConfirmGameStart        = "SteamAPI_ISteamGameSearch_HostConfirmGameStart"
	flatAPI_ISteamGameSearch_CancelRequestPlayersForGame = "SteamAPI_ISteamGameSearch_CancelRequestPlayersForGame"
	flatAPI_ISteamGameSearch_SubmitPlayerResult          = "SteamAPI_ISteamGameSearch_SubmitPlayerResult"
	flatAPI_ISteamGameSearch_EndGame                     = "SteamAPI_ISteamGameSearch_EndGame"
)

type EPlayerResult int32

const (
	EPlayerResultFailedToConnect EPlayerResult = 1
	EPlayerResultAbandoned       EPlayerResult = 2
	EPlayerResultKicked          EPlayerResult = 3
	EPlayerResultIncomplete      EPlayerResult = 4
	EPlayerResultCompleted       EPlayerResult = 5
)

type EGameSearchErrorCode int32

const (
	EGameSearchErrorCodeOK                            EGameSearchErrorCode = 1
	EGameSearchErrorCodeFailedSearchAlreadyInProgress EGameSearchErrorCode = 2
	EGameSearchErrorCodeFailedNoSearchInProgress      EGameSearchErrorCode = 3
	EGameSearchErrorCodeFailedNotLobbyLeader          EGameSearchErrorCode = 4
	EGameSearchErrorCodeFailedNoHostAvailable         EGameSearchErrorCode = 5
	EGameSearchErrorCodeFailedSearchParamsInvalid     EGameSearchErrorCode = 6
	EGameSearchErrorCodeFailedOffline                 EGameSearchErrorCode = 7
	EGameSearchErrorCodeFailedNotAuthorized           EGameSearchErrorCode = 8
	EGameSearchErrorCodeFailedUnknownError            EGameSearchErrorCode = 9
)

var (
	steamGameSearch_init                         func() uintptr
	iSteamGameSearch_AddGameSearchParams         func(steamGameSearch uintptr, keyToFind string, valuesToFind string) EGameSearchErrorCode
	iSteamGameSearch_SearchForGameWithLobby      func(steamGameSearch uintptr, lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode
	iSteamGameSearch_SearchForGameSolo           func(steamGameSearch uintptr, playerMin int32, playerMax int32) EGameSearchErrorCode
	iSteamGameSearch_AcceptGame                  func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_DeclineGame                 func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_RetrieveConnectionDetails   func(steamGameSearch uintptr, hostSteamID Uint64SteamID, cpchConnectionDetails *[]byte, connectionDetailsSize int32) EGameSearchErrorCode
	iSteamGameSearch_EndGameSearch               func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_SetGameHostParams           func(steamGameSearch uintptr, key string, value string) EGameSearchErrorCode
	iSteamGameSearch_SetConnectionDetails        func(steamGameSearch uintptr, cpchConnectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode
	iSteamGameSearch_RequestPlayersForGame       func(steamGameSearch uintptr, playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode
	iSteamGameSearch_HostConfirmGameStart        func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode
	iSteamGameSearch_CancelRequestPlayersForGame func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_SubmitPlayerResult          func(steamGameSearch uintptr, uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode
	iSteamGameSearch_EndGame                     func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode
)

var steamGameSearch uintptr

func gameSearch() uintptr {
	if steamGameSearch == 0 {
		steamGameSearch = steamGameSearch_init()
		return steamGameSearch
	}
	return steamGameSearch
}

func AddGameSearchParams(keyToFind string, valuesToFind string) EGameSearchErrorCode {
	return iSteamGameSearch_AddGameSearchParams(gameSearch(), keyToFind, valuesToFind)
}

func SearchForGameWithLobby(lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode {
	return iSteamGameSearch_SearchForGameWithLobby(gameSearch(), lobbySteamID, playerMin, playerMax)
}

func SearchForGameSolo(playerMin int32, playerMax int32) EGameSearchErrorCode {
	return iSteamGameSearch_SearchForGameSolo(gameSearch(), playerMin, playerMax)
}

func AcceptGame() EGameSearchErrorCode {
	return iSteamGameSearch_AcceptGame(gameSearch())
}

func DeclineGame() EGameSearchErrorCode {
	return iSteamGameSearch_DeclineGame(gameSearch())
}

func RetrieveConnectionDetails(hostSteamID Uint64SteamID, connectionDetailsSize int32) (string, EGameSearchErrorCode) {
	var details []byte = make([]byte, 0, connectionDetailsSize)
	code := iSteamGameSearch_RetrieveConnectionDetails(gameSearch(), hostSteamID, &details, connectionDetailsSize)
	return string(details), code
}

func EndGameSearch() EGameSearchErrorCode {
	return iSteamGameSearch_EndGameSearch(gameSearch())
}

func SetGameHostParams(key string, value string) EGameSearchErrorCode {
	return iSteamGameSearch_SetGameHostParams(gameSearch(), key, value)
}

func SetConnectionDetails(connectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode {
	return iSteamGameSearch_SetConnectionDetails(gameSearch(), connectionDetails, connectionDetailsSize)
}

func RequestPlayersForGame(playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode {
	return iSteamGameSearch_RequestPlayersForGame(gameSearch(), playerMin, playerMax, maxTeamSize)
}

func HostConfirmGameStart(uniqueGameID uint64) EGameSearchErrorCode {
	return iSteamGameSearch_HostConfirmGameStart(gameSearch(), uniqueGameID)
}

func CancelRequestPlayersForGame() EGameSearchErrorCode {
	return iSteamGameSearch_CancelRequestPlayersForGame(gameSearch())
}

func SubmitPlayerResult(uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode {
	return iSteamGameSearch_SubmitPlayerResult(gameSearch(), uniqueGameID, playerSteamID, EPlayerResult)
}

func EndGame(uniqueGameID uint64) EGameSearchErrorCode {
	return iSteamGameSearch_EndGame(gameSearch(), uniqueGameID)
}
