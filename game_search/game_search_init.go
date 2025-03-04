package gamesearch

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamGameSearch_init, SteamAPIDLL, flatAPI_SteamGameSearch)
	purego.RegisterLibFunc(&iSteamGameSearch_AddGameSearchParams, SteamAPIDLL, flatAPI_ISteamGameSearch_AddGameSearchParams)
	purego.RegisterLibFunc(&iSteamGameSearch_SearchForGameWithLobby, SteamAPIDLL, flatAPI_ISteamGameSearch_SearchForGameWithLobby)
	purego.RegisterLibFunc(&iSteamGameSearch_SearchForGameSolo, SteamAPIDLL, flatAPI_ISteamGameSearch_SearchForGameSolo)
	purego.RegisterLibFunc(&iSteamGameSearch_AcceptGame, SteamAPIDLL, flatAPI_ISteamGameSearch_AcceptGame)
	purego.RegisterLibFunc(&iSteamGameSearch_DeclineGame, SteamAPIDLL, flatAPI_ISteamGameSearch_DeclineGame)
	purego.RegisterLibFunc(&iSteamGameSearch_RetrieveConnectionDetails, SteamAPIDLL, flatAPI_ISteamGameSearch_RetrieveConnectionDetails)
	purego.RegisterLibFunc(&iSteamGameSearch_EndGameSearch, SteamAPIDLL, flatAPI_ISteamGameSearch_EndGameSearch)
	purego.RegisterLibFunc(&iSteamGameSearch_SetGameHostParams, SteamAPIDLL, flatAPI_ISteamGameSearch_SetGameHostParams)
	purego.RegisterLibFunc(&iSteamGameSearch_SetConnectionDetails, SteamAPIDLL, flatAPI_ISteamGameSearch_SetConnectionDetails)
	purego.RegisterLibFunc(&iSteamGameSearch_RequestPlayersForGame, SteamAPIDLL, flatAPI_ISteamGameSearch_RequestPlayersForGame)
	purego.RegisterLibFunc(&iSteamGameSearch_HostConfirmGameStart, SteamAPIDLL, flatAPI_ISteamGameSearch_HostConfirmGameStart)
	purego.RegisterLibFunc(&iSteamGameSearch_CancelRequestPlayersForGame, SteamAPIDLL, flatAPI_ISteamGameSearch_CancelRequestPlayersForGame)
	purego.RegisterLibFunc(&iSteamGameSearch_SubmitPlayerResult, SteamAPIDLL, flatAPI_ISteamGameSearch_SubmitPlayerResult)
	purego.RegisterLibFunc(&iSteamGameSearch_EndGame, SteamAPIDLL, flatAPI_ISteamGameSearch_EndGame)
}
