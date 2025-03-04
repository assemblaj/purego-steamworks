package gameserver

import (
	. "github.com/assemblaj/purego-steamworks/network_types"
	"github.com/assemblaj/purego-steamworks/user"

	common "github.com/assemblaj/purego-steamworks"
)

const (
	STEAMGAMESERVER_QUERY_PORT_SHARED          uint16 = 0xffff
	MASTERSERVERUPDATERPORT_USEGAMESOCKETSHARE uint16 = STEAMGAMESERVER_QUERY_PORT_SHARED
)

const (
	flatAPI_SteamGameServer                                      = "SteamAPI_SteamGameServer_v015"
	flatAPI_ISteamGameServer_SetProduct                          = "SteamAPI_ISteamGameServer_SetProduct"
	flatAPI_ISteamGameServer_SetGameDescription                  = "SteamAPI_ISteamGameServer_SetGameDescription"
	flatAPI_ISteamGameServer_SetModDir                           = "SteamAPI_ISteamGameServer_SetModDir"
	flatAPI_ISteamGameServer_SetDedicatedServer                  = "SteamAPI_ISteamGameServer_SetDedicatedServer"
	flatAPI_ISteamGameServer_LogOn                               = "SteamAPI_ISteamGameServer_LogOn"
	flatAPI_ISteamGameServer_LogOnAnonymous                      = "SteamAPI_ISteamGameServer_LogOnAnonymous"
	flatAPI_ISteamGameServer_LogOff                              = "SteamAPI_ISteamGameServer_LogOff"
	flatAPI_ISteamGameServer_BLoggedOn                           = "SteamAPI_ISteamGameServer_BLoggedOn"
	flatAPI_ISteamGameServer_BSecure                             = "SteamAPI_ISteamGameServer_BSecure"
	flatAPI_ISteamGameServer_GetSteamID                          = "SteamAPI_ISteamGameServer_GetSteamID"
	flatAPI_ISteamGameServer_WasRestartRequested                 = "SteamAPI_ISteamGameServer_WasRestartRequested"
	flatAPI_ISteamGameServer_SetMaxPlayerCount                   = "SteamAPI_ISteamGameServer_SetMaxPlayerCount"
	flatAPI_ISteamGameServer_SetBotPlayerCount                   = "SteamAPI_ISteamGameServer_SetBotPlayerCount"
	flatAPI_ISteamGameServer_SetServerName                       = "SteamAPI_ISteamGameServer_SetServerName"
	flatAPI_ISteamGameServer_SetMapName                          = "SteamAPI_ISteamGameServer_SetMapName"
	flatAPI_ISteamGameServer_SetPasswordProtected                = "SteamAPI_ISteamGameServer_SetPasswordProtected"
	flatAPI_ISteamGameServer_SetSpectatorPort                    = "SteamAPI_ISteamGameServer_SetSpectatorPort"
	flatAPI_ISteamGameServer_SetSpectatorServerName              = "SteamAPI_ISteamGameServer_SetSpectatorServerName"
	flatAPI_ISteamGameServer_ClearAllKeyValues                   = "SteamAPI_ISteamGameServer_ClearAllKeyValues"
	flatAPI_ISteamGameServer_SetKeyValue                         = "SteamAPI_ISteamGameServer_SetKeyValue"
	flatAPI_ISteamGameServer_SetGameTags                         = "SteamAPI_ISteamGameServer_SetGameTags"
	flatAPI_ISteamGameServer_SetGameData                         = "SteamAPI_ISteamGameServer_SetGameData"
	flatAPI_ISteamGameServer_SetRegion                           = "SteamAPI_ISteamGameServer_SetRegion"
	flatAPI_ISteamGameServer_SetAdvertiseServerActive            = "SteamAPI_ISteamGameServer_SetAdvertiseServerActive"
	flatAPI_ISteamGameServer_GetAuthSessionTicket                = "SteamAPI_ISteamGameServer_GetAuthSessionTicket"
	flatAPI_ISteamGameServer_BeginAuthSession                    = "SteamAPI_ISteamGameServer_BeginAuthSession"
	flatAPI_ISteamGameServer_EndAuthSession                      = "SteamAPI_ISteamGameServer_EndAuthSession"
	flatAPI_ISteamGameServer_CancelAuthTicket                    = "SteamAPI_ISteamGameServer_CancelAuthTicket"
	flatAPI_ISteamGameServer_UserHasLicenseForApp                = "SteamAPI_ISteamGameServer_UserHasLicenseForApp"
	flatAPI_ISteamGameServer_RequestUserGroupStatus              = "SteamAPI_ISteamGameServer_RequestUserGroupStatus"
	flatAPI_ISteamGameServer_GetPublicIP                         = "SteamAPI_ISteamGameServer_GetPublicIP"
	flatAPI_ISteamGameServer_HandleIncomingPacket                = "SteamAPI_ISteamGameServer_HandleIncomingPacket"
	flatAPI_ISteamGameServer_GetNextOutgoingPacket               = "SteamAPI_ISteamGameServer_GetNextOutgoingPacket"
	flatAPI_ISteamGameServer_AssociateWithClan                   = "SteamAPI_ISteamGameServer_AssociateWithClan"
	flatAPI_ISteamGameServer_ComputeNewPlayerCompatibility       = "SteamAPI_ISteamGameServer_ComputeNewPlayerCompatibility"
	flatAPI_ISteamGameServer_CreateUnauthenticatedUserConnection = "SteamAPI_ISteamGameServer_CreateUnauthenticatedUserConnection"
	flatAPI_ISteamGameServer_BUpdateUserData                     = "SteamAPI_ISteamGameServer_BUpdateUserData"
)

var (
	steamGameServer_init                                 func() uintptr
	iSteamGameServer_SetProduct                          func(steamGameServer uintptr, pszProduct string)
	iSteamGameServer_SetGameDescription                  func(steamGameServer uintptr, pszGameDescription string)
	iSteamGameServer_SetModDir                           func(steamGameServer uintptr, modDir string)
	iSteamGameServer_SetDedicatedServer                  func(steamGameServer uintptr, dedicated bool)
	iSteamGameServer_LogOn                               func(steamGameServer uintptr, token string)
	iSteamGameServer_LogOnAnonymous                      func(steamGameServer uintptr)
	iSteamGameServer_LogOff                              func(steamGameServer uintptr)
	iSteamGameServer_BLoggedOn                           func(steamGameServer uintptr) bool
	iSteamGameServer_BSecure                             func(steamGameServer uintptr) bool
	iSteamGameServer_GetSteamID                          func(steamGameServer uintptr) common.Uint64SteamID
	iSteamGameServer_WasRestartRequested                 func(steamGameServer uintptr) bool
	iSteamGameServer_SetMaxPlayerCount                   func(steamGameServer uintptr, playersMax int32)
	iSteamGameServer_SetBotPlayerCount                   func(steamGameServer uintptr, botplayers int32)
	iSteamGameServer_SetServerName                       func(steamGameServer uintptr, serverName string)
	iSteamGameServer_SetMapName                          func(steamGameServer uintptr, mapName string)
	iSteamGameServer_SetPasswordProtected                func(steamGameServer uintptr, passwordProtected bool)
	iSteamGameServer_SetSpectatorPort                    func(steamGameServer uintptr, spectatorPort uint16)
	iSteamGameServer_SetSpectatorServerName              func(steamGameServer uintptr, spectatorServerName string)
	iSteamGameServer_ClearAllKeyValues                   func(steamGameServer uintptr)
	iSteamGameServer_SetKeyValue                         func(steamGameServer uintptr, key string, value string)
	iSteamGameServer_SetGameTags                         func(steamGameServer uintptr, gameTags string)
	iSteamGameServer_SetGameData                         func(steamGameServer uintptr, gameData string)
	iSteamGameServer_SetRegion                           func(steamGameServer uintptr, region string)
	iSteamGameServer_SetAdvertiseServerActive            func(steamGameServer uintptr, active bool)
	iSteamGameServer_GetAuthSessionTicket                func(steamGameServer uintptr, pTicket *[]byte, maxTicket int32, pcbTicket *uint32, snid *SteamNetworkingIdentity) user.HAuthTicket
	iSteamGameServer_BeginAuthSession                    func(steamGameServer uintptr, authTicket []byte, cbAuthTicket int32, steamID common.Uint64SteamID) user.EBeginAuthSessionResult
	iSteamGameServer_EndAuthSession                      func(steamGameServer uintptr, steamID common.Uint64SteamID)
	iSteamGameServer_CancelAuthTicket                    func(steamGameServer uintptr, hAuthTicket user.HAuthTicket)
	iSteamGameServer_UserHasLicenseForApp                func(steamGameServer uintptr, steamID common.Uint64SteamID, AppId common.AppId) user.EUserHasLicenseForAppResult
	iSteamGameServer_RequestUserGroupStatus              func(steamGameServer uintptr, user common.Uint64SteamID, groupSteamID common.Uint64SteamID) bool
	iSteamGameServer_GetPublicIP                         func(steamGameServer uintptr) uintptr
	iSteamGameServer_HandleIncomingPacket                func(steamGameServer uintptr, pData *[]byte, data int32, srcIP uint32, srcPort uint16) bool
	iSteamGameServer_GetNextOutgoingPacket               func(steamGameServer uintptr, pOut *[]byte, maxOut int32, pNetAdr *uint32, pPort *uint16) int32
	iSteamGameServer_AssociateWithClan                   func(steamGameServer uintptr, clanSteamID common.Uint64SteamID) common.SteamAPICall
	iSteamGameServer_ComputeNewPlayerCompatibility       func(steamGameServer uintptr, newPlayerSteamID common.Uint64SteamID) common.SteamAPICall
	iSteamGameServer_CreateUnauthenticatedUserConnection func(steamGameServer uintptr) common.Uint64SteamID
	iSteamGameServer_BUpdateUserData                     func(steamGameServer uintptr, user common.Uint64SteamID, playerName string, score uint32) bool
)

var steamGameServer uintptr

func gameServer() uintptr {
	if steamGameServer == 0 {
		steamGameServer = steamGameServer_init()
		return steamGameServer
	}
	return steamGameServer
}

func SetProduct(product string) {
	iSteamGameServer_SetProduct(gameServer(), product)
}

func SetGameDescription(gameDescription string) {
	iSteamGameServer_SetGameDescription(gameServer(), gameDescription)
}

func SetModDir(modDir string) {
	iSteamGameServer_SetModDir(gameServer(), modDir)
}

func SetDedicatedServer(dedicated bool) {
	iSteamGameServer_SetDedicatedServer(gameServer(), dedicated)
}

func LogOn(token string) {
	iSteamGameServer_LogOn(gameServer(), token)
}

func LogOnAnonymous() {
	iSteamGameServer_LogOnAnonymous(gameServer())
}

func LogOff() {
	iSteamGameServer_LogOff(gameServer())
}

func BLoggedOn() bool {
	return iSteamGameServer_BLoggedOn(gameServer())
}

func BSecure() bool {
	return iSteamGameServer_BSecure(gameServer())
}

func GetSteamID() common.Uint64SteamID {
	return iSteamGameServer_GetSteamID(gameServer())
}

func WasRestartRequested() bool {
	return iSteamGameServer_WasRestartRequested(gameServer())
}

func SetMaxPlayerCount(playersMax int32) {
	iSteamGameServer_SetMaxPlayerCount(gameServer(), playersMax)
}

func SetBotPlayerCount(botplayers int32) {
	iSteamGameServer_SetBotPlayerCount(gameServer(), botplayers)
}

func SetServerName(serverName string) {
	iSteamGameServer_SetServerName(gameServer(), serverName)
}

func SetMapName(mapName string) {
	iSteamGameServer_SetMapName(gameServer(), mapName)
}

func SetPasswordProtected(passwordProtected bool) {
	iSteamGameServer_SetPasswordProtected(gameServer(), passwordProtected)
}

func SetSpectatorPort(spectatorPort uint16) {
	iSteamGameServer_SetSpectatorPort(gameServer(), spectatorPort)
}

func SetSpectatorServerName(spectatorServerName string) {
	iSteamGameServer_SetSpectatorServerName(gameServer(), spectatorServerName)
}

func ClearAllKeyValues() {
	iSteamGameServer_ClearAllKeyValues(gameServer())
}

func SetKeyValue(key string, value string) {
	iSteamGameServer_SetKeyValue(gameServer(), key, value)
}

func SetGameTags(gameTags string) {
	iSteamGameServer_SetGameTags(gameServer(), gameTags)
}

func SetGameData(gameData string) {
	iSteamGameServer_SetGameData(gameServer(), gameData)
}

func SetRegion(region string) {
	iSteamGameServer_SetRegion(gameServer(), region)
}

func SetAdvertiseServerActive(active bool) {
	iSteamGameServer_SetAdvertiseServerActive(gameServer(), active)
}

// snid optional
func GetAuthSessionTicket(maxTicket int32, snid *SteamNetworkingIdentity) ([]byte, user.HAuthTicket) {
	var pTicket []byte = make([]byte, 0, maxTicket)
	var pcbTicket uint32
	result := iSteamGameServer_GetAuthSessionTicket(gameServer(), &pTicket, maxTicket, &pcbTicket, snid)
	return pTicket[:pcbTicket], result
}

func BeginAuthSession(authTicket []byte, steamID common.Uint64SteamID) user.EBeginAuthSessionResult {
	return iSteamGameServer_BeginAuthSession(gameServer(), authTicket, int32(len(authTicket)), steamID)
}

func EndAuthSession(steamID common.Uint64SteamID) {
	iSteamGameServer_EndAuthSession(gameServer(), steamID)
}

func CancelAuthTicket(authTicket user.HAuthTicket) {
	iSteamGameServer_CancelAuthTicket(gameServer(), authTicket)
}

func UserHasLicenseForApp(steamID common.Uint64SteamID, AppId common.AppId) user.EUserHasLicenseForAppResult {
	return iSteamGameServer_UserHasLicenseForApp(gameServer(), steamID, AppId)
}

func RequestUserGroupStatus(userSteamID common.Uint64SteamID, groupSteamID common.Uint64SteamID) bool {
	return iSteamGameServer_RequestUserGroupStatus(gameServer(), userSteamID, groupSteamID)
}

func GetPublicIP() SteamIPAddress {
	return *common.UintptrToStruct[SteamIPAddress](iSteamGameServer_GetPublicIP(gameServer()))
}

func HandleIncomingPacket(data int32, srcIP uint32, srcPort uint16) ([]byte, bool) {
	pData := make([]byte, 0, data)
	result := iSteamGameServer_HandleIncomingPacket(gameServer(), &pData, data, srcIP, srcPort)
	return pData, result
}

func GetNextOutgoingPacket(maxOut int32) (packet []byte, addr uint32, port uint16, result int32) {
	packet = make([]byte, 0, maxOut)
	result = iSteamGameServer_GetNextOutgoingPacket(gameServer(), &packet, maxOut, &addr, &port)
	return packet, addr, port, result
}

func AssociateWithClan(clanSteamID common.Uint64SteamID) common.CallResult[AssociateWithClanResult] {
	handle := iSteamGameServer_AssociateWithClan(gameServer(), clanSteamID)
	return common.CallResult[AssociateWithClanResult]{Handle: handle}
}

func ComputeNewPlayerCompatibility(newPlayerSteamID common.Uint64SteamID) common.CallResult[ComputeNewPlayerCompatibilityResult] {
	handle := iSteamGameServer_ComputeNewPlayerCompatibility(gameServer(), newPlayerSteamID)
	return common.CallResult[ComputeNewPlayerCompatibilityResult]{Handle: handle}
}

func CreateUnauthenticatedUserConnection() common.Uint64SteamID {
	return iSteamGameServer_CreateUnauthenticatedUserConnection(gameServer())
}

func BUpdateUserData(user common.Uint64SteamID, playerName string, score uint32) bool {
	return iSteamGameServer_BUpdateUserData(gameServer(), user, playerName, score)
}
