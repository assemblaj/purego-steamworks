package gameserver

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamGameServer_init, SteamAPIDLL, flatAPI_SteamGameServer)
	purego.RegisterLibFunc(&iSteamGameServer_SetProduct, SteamAPIDLL, flatAPI_ISteamGameServer_SetProduct)
	purego.RegisterLibFunc(&iSteamGameServer_SetGameDescription, SteamAPIDLL, flatAPI_ISteamGameServer_SetGameDescription)
	purego.RegisterLibFunc(&iSteamGameServer_SetModDir, SteamAPIDLL, flatAPI_ISteamGameServer_SetModDir)
	purego.RegisterLibFunc(&iSteamGameServer_SetDedicatedServer, SteamAPIDLL, flatAPI_ISteamGameServer_SetDedicatedServer)
	purego.RegisterLibFunc(&iSteamGameServer_LogOn, SteamAPIDLL, flatAPI_ISteamGameServer_LogOn)
	purego.RegisterLibFunc(&iSteamGameServer_LogOnAnonymous, SteamAPIDLL, flatAPI_ISteamGameServer_LogOnAnonymous)
	purego.RegisterLibFunc(&iSteamGameServer_LogOff, SteamAPIDLL, flatAPI_ISteamGameServer_LogOff)
	purego.RegisterLibFunc(&iSteamGameServer_BLoggedOn, SteamAPIDLL, flatAPI_ISteamGameServer_BLoggedOn)
	purego.RegisterLibFunc(&iSteamGameServer_BSecure, SteamAPIDLL, flatAPI_ISteamGameServer_BSecure)
	purego.RegisterLibFunc(&iSteamGameServer_GetSteamID, SteamAPIDLL, flatAPI_ISteamGameServer_GetSteamID)
	purego.RegisterLibFunc(&iSteamGameServer_WasRestartRequested, SteamAPIDLL, flatAPI_ISteamGameServer_WasRestartRequested)
	purego.RegisterLibFunc(&iSteamGameServer_SetMaxPlayerCount, SteamAPIDLL, flatAPI_ISteamGameServer_SetMaxPlayerCount)
	purego.RegisterLibFunc(&iSteamGameServer_SetBotPlayerCount, SteamAPIDLL, flatAPI_ISteamGameServer_SetBotPlayerCount)
	purego.RegisterLibFunc(&iSteamGameServer_SetServerName, SteamAPIDLL, flatAPI_ISteamGameServer_SetServerName)
	purego.RegisterLibFunc(&iSteamGameServer_SetMapName, SteamAPIDLL, flatAPI_ISteamGameServer_SetMapName)
	purego.RegisterLibFunc(&iSteamGameServer_SetPasswordProtected, SteamAPIDLL, flatAPI_ISteamGameServer_SetPasswordProtected)
	purego.RegisterLibFunc(&iSteamGameServer_SetSpectatorPort, SteamAPIDLL, flatAPI_ISteamGameServer_SetSpectatorPort)
	purego.RegisterLibFunc(&iSteamGameServer_SetSpectatorServerName, SteamAPIDLL, flatAPI_ISteamGameServer_SetSpectatorServerName)
	purego.RegisterLibFunc(&iSteamGameServer_ClearAllKeyValues, SteamAPIDLL, flatAPI_ISteamGameServer_ClearAllKeyValues)
	purego.RegisterLibFunc(&iSteamGameServer_SetKeyValue, SteamAPIDLL, flatAPI_ISteamGameServer_SetKeyValue)
	purego.RegisterLibFunc(&iSteamGameServer_SetGameTags, SteamAPIDLL, flatAPI_ISteamGameServer_SetGameTags)
	purego.RegisterLibFunc(&iSteamGameServer_SetGameData, SteamAPIDLL, flatAPI_ISteamGameServer_SetGameData)
	purego.RegisterLibFunc(&iSteamGameServer_SetRegion, SteamAPIDLL, flatAPI_ISteamGameServer_SetRegion)
	purego.RegisterLibFunc(&iSteamGameServer_SetAdvertiseServerActive, SteamAPIDLL, flatAPI_ISteamGameServer_SetAdvertiseServerActive)
	purego.RegisterLibFunc(&iSteamGameServer_GetAuthSessionTicket, SteamAPIDLL, flatAPI_ISteamGameServer_GetAuthSessionTicket)
	purego.RegisterLibFunc(&iSteamGameServer_BeginAuthSession, SteamAPIDLL, flatAPI_ISteamGameServer_BeginAuthSession)
	purego.RegisterLibFunc(&iSteamGameServer_EndAuthSession, SteamAPIDLL, flatAPI_ISteamGameServer_EndAuthSession)
	purego.RegisterLibFunc(&iSteamGameServer_CancelAuthTicket, SteamAPIDLL, flatAPI_ISteamGameServer_CancelAuthTicket)
	purego.RegisterLibFunc(&iSteamGameServer_UserHasLicenseForApp, SteamAPIDLL, flatAPI_ISteamGameServer_UserHasLicenseForApp)
	purego.RegisterLibFunc(&iSteamGameServer_RequestUserGroupStatus, SteamAPIDLL, flatAPI_ISteamGameServer_RequestUserGroupStatus)
	purego.RegisterLibFunc(&iSteamGameServer_GetPublicIP, SteamAPIDLL, flatAPI_ISteamGameServer_GetPublicIP)
	purego.RegisterLibFunc(&iSteamGameServer_HandleIncomingPacket, SteamAPIDLL, flatAPI_ISteamGameServer_HandleIncomingPacket)
	purego.RegisterLibFunc(&iSteamGameServer_GetNextOutgoingPacket, SteamAPIDLL, flatAPI_ISteamGameServer_GetNextOutgoingPacket)
	purego.RegisterLibFunc(&iSteamGameServer_AssociateWithClan, SteamAPIDLL, flatAPI_ISteamGameServer_AssociateWithClan)
	purego.RegisterLibFunc(&iSteamGameServer_ComputeNewPlayerCompatibility, SteamAPIDLL, flatAPI_ISteamGameServer_ComputeNewPlayerCompatibility)
	purego.RegisterLibFunc(&iSteamGameServer_CreateUnauthenticatedUserConnection, SteamAPIDLL, flatAPI_ISteamGameServer_CreateUnauthenticatedUserConnection)
	purego.RegisterLibFunc(&iSteamGameServer_BUpdateUserData, SteamAPIDLL, flatAPI_ISteamGameServer_BUpdateUserData)
}
