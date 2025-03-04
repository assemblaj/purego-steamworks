package matchmaking

import (
	. "github.com/assemblaj/purego-steamworks"
)

const (
	MaxGameServerGameDir         int32  = 32
	MaxGameServerMapName         int32  = 32
	MaxGameServerGameDescription int32  = 64
	MaxGameServerName            int32  = 64
	MaxGameServerTags            int32  = 128
	MaxGameServerGameData        int32  = 2048
	HSERVERQUERY_INVALID         int32  = -1
	FavoriteFlagNone             uint32 = 0x00
	FavoriteFlagFavorite         uint32 = 0x01
	FavoriteFlagHistory          uint32 = 0x02
)

type EChatEntryType int32

const (
	EChatEntryTypeInvalid          EChatEntryType = 0
	EChatEntryTypeChatMsg          EChatEntryType = 1
	EChatEntryTypeTyping           EChatEntryType = 2
	EChatEntryTypeInviteGame       EChatEntryType = 3
	EChatEntryTypeEmote            EChatEntryType = 4
	EChatEntryTypeLeftConversation EChatEntryType = 6
	EChatEntryTypeEntered          EChatEntryType = 7
	EChatEntryTypeWasKicked        EChatEntryType = 8
	EChatEntryTypeWasBanned        EChatEntryType = 9
	EChatEntryTypeDisconnected     EChatEntryType = 10
	EChatEntryTypeHistoricalChat   EChatEntryType = 11
	EChatEntryTypeLinkBlocked      EChatEntryType = 14
)

type EMatchMakingServerResponse int32

const (
	EServerResponded               EMatchMakingServerResponse = 0
	EServerFailedToRespond         EMatchMakingServerResponse = 1
	ENoServersListedOnMasterServer EMatchMakingServerResponse = 2
)

type ELobbyType int32

const (
	ELobbyTypePrivate       ELobbyType = 0
	ELobbyTypeFriendsOnly   ELobbyType = 1
	ELobbyTypePublic        ELobbyType = 2
	ELobbyTypeInvisible     ELobbyType = 3
	ELobbyTypePrivateUnique ELobbyType = 4
)

type ELobbyComparison int32

const (
	ELobbyComparisonEqualToOrLessThan    ELobbyComparison = -2
	ELobbyComparisonLessThan             ELobbyComparison = -1
	ELobbyComparisonEqual                ELobbyComparison = 0
	ELobbyComparisonGreaterThan          ELobbyComparison = 1
	ELobbyComparisonEqualToOrGreaterThan ELobbyComparison = 2
	ELobbyComparisonNotEqual             ELobbyComparison = 3
)

type ELobbyDistanceFilter int32

const (
	ELobbyDistanceFilterClose     ELobbyDistanceFilter = 0
	ELobbyDistanceFilterDefault   ELobbyDistanceFilter = 1
	ELobbyDistanceFilterFar       ELobbyDistanceFilter = 2
	ELobbyDistanceFilterWorldwide ELobbyDistanceFilter = 3
)

type EChatMemberStateChange int32

const (
	EChatMemberStateChangeEntered      EChatMemberStateChange = 1
	EChatMemberStateChangeLeft         EChatMemberStateChange = 2
	EChatMemberStateChangeDisconnected EChatMemberStateChange = 4
	EChatMemberStateChangeKicked       EChatMemberStateChange = 8
	EChatMemberStateChangeBanned       EChatMemberStateChange = 16
)

const (
	flatAPI_SteamMatchmaking                                             = "SteamAPI_SteamMatchmaking_v009"
	flatAPI_ISteamMatchmaking_GetFavoriteGameCount                       = "SteamAPI_ISteamMatchmaking_GetFavoriteGameCount"
	flatAPI_ISteamMatchmaking_GetFavoriteGame                            = "SteamAPI_ISteamMatchmaking_GetFavoriteGame"
	flatAPI_ISteamMatchmaking_AddFavoriteGame                            = "SteamAPI_ISteamMatchmaking_AddFavoriteGame"
	flatAPI_ISteamMatchmaking_RemoveFavoriteGame                         = "SteamAPI_ISteamMatchmaking_RemoveFavoriteGame"
	flatAPI_ISteamMatchmaking_RequestLobbyList                           = "SteamAPI_ISteamMatchmaking_RequestLobbyList"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter            = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter         = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter         = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable    = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter          = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter       = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter"
	flatAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter = "SteamAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter"
	flatAPI_ISteamMatchmaking_GetLobbyByIndex                            = "SteamAPI_ISteamMatchmaking_GetLobbyByIndex"
	flatAPI_ISteamMatchmaking_CreateLobby                                = "SteamAPI_ISteamMatchmaking_CreateLobby"
	flatAPI_ISteamMatchmaking_JoinLobby                                  = "SteamAPI_ISteamMatchmaking_JoinLobby"
	flatAPI_ISteamMatchmaking_LeaveLobby                                 = "SteamAPI_ISteamMatchmaking_LeaveLobby"
	flatAPI_ISteamMatchmaking_InviteUserToLobby                          = "SteamAPI_ISteamMatchmaking_InviteUserToLobby"
	flatAPI_ISteamMatchmaking_GetNumLobbyMembers                         = "SteamAPI_ISteamMatchmaking_GetNumLobbyMembers"
	flatAPI_ISteamMatchmaking_GetLobbyMemberByIndex                      = "SteamAPI_ISteamMatchmaking_GetLobbyMemberByIndex"
	flatAPI_ISteamMatchmaking_GetLobbyData                               = "SteamAPI_ISteamMatchmaking_GetLobbyData"
	flatAPI_ISteamMatchmaking_SetLobbyData                               = "SteamAPI_ISteamMatchmaking_SetLobbyData"
	flatAPI_ISteamMatchmaking_GetLobbyDataCount                          = "SteamAPI_ISteamMatchmaking_GetLobbyDataCount"
	flatAPI_ISteamMatchmaking_GetLobbyDataByIndex                        = "SteamAPI_ISteamMatchmaking_GetLobbyDataByIndex"
	flatAPI_ISteamMatchmaking_DeleteLobbyData                            = "SteamAPI_ISteamMatchmaking_DeleteLobbyData"
	flatAPI_ISteamMatchmaking_GetLobbyMemberData                         = "SteamAPI_ISteamMatchmaking_GetLobbyMemberData"
	flatAPI_ISteamMatchmaking_SetLobbyMemberData                         = "SteamAPI_ISteamMatchmaking_SetLobbyMemberData"
	flatAPI_ISteamMatchmaking_SendLobbyChatMsg                           = "SteamAPI_ISteamMatchmaking_SendLobbyChatMsg"
	flatAPI_ISteamMatchmaking_GetLobbyChatEntry                          = "SteamAPI_ISteamMatchmaking_GetLobbyChatEntry"
	flatAPI_ISteamMatchmaking_RequestLobbyData                           = "SteamAPI_ISteamMatchmaking_RequestLobbyData"
	flatAPI_ISteamMatchmaking_SetLobbyGameServer                         = "SteamAPI_ISteamMatchmaking_SetLobbyGameServer"
	flatAPI_ISteamMatchmaking_GetLobbyGameServer                         = "SteamAPI_ISteamMatchmaking_GetLobbyGameServer"
	flatAPI_ISteamMatchmaking_SetLobbyMemberLimit                        = "SteamAPI_ISteamMatchmaking_SetLobbyMemberLimit"
	flatAPI_ISteamMatchmaking_GetLobbyMemberLimit                        = "SteamAPI_ISteamMatchmaking_GetLobbyMemberLimit"
	flatAPI_ISteamMatchmaking_SetLobbyType                               = "SteamAPI_ISteamMatchmaking_SetLobbyType"
	flatAPI_ISteamMatchmaking_SetLobbyJoinable                           = "SteamAPI_ISteamMatchmaking_SetLobbyJoinable"
	flatAPI_ISteamMatchmaking_GetLobbyOwner                              = "SteamAPI_ISteamMatchmaking_GetLobbyOwner"
	flatAPI_ISteamMatchmaking_SetLobbyOwner                              = "SteamAPI_ISteamMatchmaking_SetLobbyOwner"
	flatAPI_ISteamMatchmaking_SetLinkedLobby                             = "SteamAPI_ISteamMatchmaking_SetLinkedLobby"
)

var (
	steamMatchmaking_init                                        func() uintptr
	iSteamMatchmaking_GetFavoriteGameCount                       func(steamMatchmaking uintptr) int32
	iSteamMatchmaking_GetFavoriteGame                            func(steamMatchmaking uintptr, iGame int32, pnAppID *AppId, pnIP *uint32, pnConnPort *uint16, pnQueryPort *uint16, punFlags *uint32, pRTime32LastPlayedOnServer *uint32) bool
	iSteamMatchmaking_AddFavoriteGame                            func(steamMatchmaking uintptr, nAppID AppId, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32, rTime32LastPlayedOnServer uint32) int32
	iSteamMatchmaking_RemoveFavoriteGame                         func(steamMatchmaking uintptr, nAppID AppId, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32) bool
	iSteamMatchmaking_RequestLobbyList                           func(steamMatchmaking uintptr) SteamAPICall
	iSteamMatchmaking_AddRequestLobbyListStringFilter            func(steamMatchmaking uintptr, pchKeyToMatch string, pchValueToMatch string, eComparisonType ELobbyComparison)
	iSteamMatchmaking_AddRequestLobbyListNumericalFilter         func(steamMatchmaking uintptr, pchKeyToMatch string, nValueToMatch int32, eComparisonType ELobbyComparison)
	iSteamMatchmaking_AddRequestLobbyListNearValueFilter         func(steamMatchmaking uintptr, pchKeyToMatch string, nValueToBeCloseTo int32)
	iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable    func(steamMatchmaking uintptr, nSlotsAvailable int32)
	iSteamMatchmaking_AddRequestLobbyListDistanceFilter          func(steamMatchmaking uintptr, eLobbyDistanceFilter ELobbyDistanceFilter)
	iSteamMatchmaking_AddRequestLobbyListResultCountFilter       func(steamMatchmaking uintptr, cMaxResults int32)
	iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID)
	iSteamMatchmaking_GetLobbyByIndex                            func(steamMatchmaking uintptr, iLobby int32) Uint64SteamID
	iSteamMatchmaking_CreateLobby                                func(steamMatchmaking uintptr, eLobbyType ELobbyType, cMaxMembers int32) SteamAPICall
	iSteamMatchmaking_JoinLobby                                  func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) SteamAPICall
	iSteamMatchmaking_LeaveLobby                                 func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID)
	iSteamMatchmaking_InviteUserToLobby                          func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDInvitee Uint64SteamID) bool
	iSteamMatchmaking_GetNumLobbyMembers                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32
	iSteamMatchmaking_GetLobbyMemberByIndex                      func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iMember int32) Uint64SteamID
	iSteamMatchmaking_GetLobbyData                               func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string) []byte
	iSteamMatchmaking_SetLobbyData                               func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string, pchValue string) bool
	iSteamMatchmaking_GetLobbyDataCount                          func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32
	iSteamMatchmaking_GetLobbyDataByIndex                        func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iLobbyData int32, pchKey *[]byte, cchKeyBufferSize int32, pchValue *[]byte, cchValueBufferSize int32) bool
	iSteamMatchmaking_DeleteLobbyData                            func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string) bool
	iSteamMatchmaking_GetLobbyMemberData                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDUser Uint64SteamID, pchKey string) []byte
	iSteamMatchmaking_SetLobbyMemberData                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string, pchValue string)
	iSteamMatchmaking_SendLobbyChatMsg                           func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pvMsgBody []byte, cubMsgBody int32) bool
	iSteamMatchmaking_GetLobbyChatEntry                          func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iChatID int32, pSteamIDUser *CSteamID, pvData *[]byte, cubData int32, peChatEntryType *EChatEntryType) int32
	iSteamMatchmaking_RequestLobbyData                           func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) bool
	iSteamMatchmaking_SetLobbyGameServer                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, unGameServerIP uint32, unGameServerPort uint16, steamIDGameServer Uint64SteamID)
	iSteamMatchmaking_GetLobbyGameServer                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, punGameServerIP *uint32, punGameServerPort *uint16, psteamIDGameServer *CSteamID) bool
	iSteamMatchmaking_SetLobbyMemberLimit                        func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, cMaxMembers int32) bool
	iSteamMatchmaking_GetLobbyMemberLimit                        func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32
	iSteamMatchmaking_SetLobbyType                               func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, eLobbyType ELobbyType) bool
	iSteamMatchmaking_SetLobbyJoinable                           func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, bLobbyJoinable bool) bool
	iSteamMatchmaking_GetLobbyOwner                              func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) Uint64SteamID
	iSteamMatchmaking_SetLobbyOwner                              func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDNewOwner Uint64SteamID) bool
	iSteamMatchmaking_SetLinkedLobby                             func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDLobbyDependent Uint64SteamID) bool
)

var steamMatchmaking uintptr

func matchmaking() uintptr {
	if steamMatchmaking == 0 {
		steamMatchmaking = steamMatchmaking_init()
		return steamMatchmaking
	}
	return steamMatchmaking
}

func GetFavoriteGameCount() int32 {
	return iSteamMatchmaking_GetFavoriteGameCount(matchmaking())
}

func GetFavoriteGame(gameIndex int32) (appID AppId, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32, success bool) {
	success = iSteamMatchmaking_GetFavoriteGame(matchmaking(), gameIndex, &appID, &IP, &connPort, &queryPort, &flags, &lastPlayedOnServerTime)
	return appID, IP, connPort, queryPort, flags, lastPlayedOnServerTime, success
}

func AddFavoriteGame(appID AppId, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32) int32 {
	return iSteamMatchmaking_AddFavoriteGame(matchmaking(), appID, IP, connPort, queryPort, flags, lastPlayedOnServerTime)
}

func RemoveFavoriteGame(appID AppId, IP uint32, connPort uint16, queryPort uint16, flags uint32) bool {
	return iSteamMatchmaking_RemoveFavoriteGame(matchmaking(), appID, IP, connPort, queryPort, flags)
}

func RequestLobbyList() CallResult[LobbyMatchList] {
	handle := iSteamMatchmaking_RequestLobbyList(matchmaking())
	return CallResult[LobbyMatchList]{Handle: handle}
}

func AddRequestLobbyListStringFilter(keyToMatch string, valueToMatch string, comparisonType ELobbyComparison) {
	iSteamMatchmaking_AddRequestLobbyListStringFilter(matchmaking(), keyToMatch, valueToMatch, comparisonType)
}

func AddRequestLobbyListNumericalFilter(keyToMatch string, valueToMatch int32, comparisonType ELobbyComparison) {
	iSteamMatchmaking_AddRequestLobbyListNumericalFilter(matchmaking(), keyToMatch, valueToMatch, comparisonType)
}

func AddRequestLobbyListNearValueFilter(keyToMatch string, valueToBeCloseTo int32) {
	iSteamMatchmaking_AddRequestLobbyListNearValueFilter(matchmaking(), keyToMatch, valueToBeCloseTo)
}

func AddRequestLobbyListFilterSlotsAvailable(slotsAvailable int32) {
	iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable(matchmaking(), slotsAvailable)
}

func AddRequestLobbyListDistanceFilter(lobbyDistanceFilter ELobbyDistanceFilter) {
	iSteamMatchmaking_AddRequestLobbyListDistanceFilter(matchmaking(), lobbyDistanceFilter)
}

func AddRequestLobbyListResultCountFilter(cMaxResults int32) {
	iSteamMatchmaking_AddRequestLobbyListResultCountFilter(matchmaking(), cMaxResults)
}

func AddRequestLobbyListCompatibleMembersFilter(lobbySteamID Uint64SteamID) {
	iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter(matchmaking(), lobbySteamID)
}

func GetLobbyByIndex(iLobby int32) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyByIndex(matchmaking(), iLobby)
}

func CreateLobby(lobbyType ELobbyType, maxMembers int32) CallResult[LobbyCreated] {
	handle := iSteamMatchmaking_CreateLobby(matchmaking(), lobbyType, maxMembers)
	return CallResult[LobbyCreated]{Handle: handle}
}

func JoinLobby(lobbySteamID Uint64SteamID) CallResult[LobbyEnter] {
	handle := iSteamMatchmaking_JoinLobby(matchmaking(), lobbySteamID)
	return CallResult[LobbyEnter]{Handle: handle}
}

func LeaveLobby(lobbySteamID Uint64SteamID) {
	iSteamMatchmaking_LeaveLobby(matchmaking(), lobbySteamID)
}

func InviteUserToLobby(lobbySteamID Uint64SteamID, inviteeSteamID Uint64SteamID) bool {
	return iSteamMatchmaking_InviteUserToLobby(matchmaking(), lobbySteamID, inviteeSteamID)
}

func GetNumLobbyMembers(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetNumLobbyMembers(matchmaking(), lobbySteamID)
}

func GetLobbyMemberByIndex(lobbySteamID Uint64SteamID, memberIndex int32) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyMemberByIndex(matchmaking(), lobbySteamID, memberIndex)
}

func GetLobbyData(lobbySteamID Uint64SteamID, key string) []byte {
	return iSteamMatchmaking_GetLobbyData(matchmaking(), lobbySteamID, key)
}

func SetLobbyData(lobbySteamID Uint64SteamID, key string, value string) bool {
	return iSteamMatchmaking_SetLobbyData(matchmaking(), lobbySteamID, key, value)
}

func GetLobbyDataCount(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetLobbyDataCount(matchmaking(), lobbySteamID)
}

func GetLobbyDataByIndex(lobbySteamID Uint64SteamID, lobbyDataIndex int32, keyBufferSize int32, valueBufferSize int32) (key []byte, value []byte, success bool) {
	key = make([]byte, 0, keyBufferSize)
	value = make([]byte, 0, valueBufferSize)
	success = iSteamMatchmaking_GetLobbyDataByIndex(matchmaking(), lobbySteamID, lobbyDataIndex, &key, keyBufferSize, &value, valueBufferSize)
	return key, value, success
}

func DeleteLobbyData(lobbySteamID Uint64SteamID, key string) bool {
	return iSteamMatchmaking_DeleteLobbyData(matchmaking(), lobbySteamID, key)
}

func GetLobbyMemberData(lobbySteamID Uint64SteamID, userSteamID Uint64SteamID, key string) []byte {
	return iSteamMatchmaking_GetLobbyMemberData(matchmaking(), lobbySteamID, userSteamID, key)
}

func SetLobbyMemberData(lobbySteamID Uint64SteamID, key string, value string) {
	iSteamMatchmaking_SetLobbyMemberData(matchmaking(), lobbySteamID, key, value)
}

func SendLobbyChatMsg(lobbySteamID Uint64SteamID, msgBody []byte) bool {
	return iSteamMatchmaking_SendLobbyChatMsg(matchmaking(), lobbySteamID, msgBody, int32(len(msgBody)))
}

func GetLobbyChatEntry(lobbySteamID Uint64SteamID, chatID int32, dataSize int32) (userSteamID CSteamID, data []byte, chatEntryType EChatEntryType) {
	data = make([]byte, 0, dataSize)
	result := iSteamMatchmaking_GetLobbyChatEntry(matchmaking(), lobbySteamID, chatID, &userSteamID, &data, dataSize, &chatEntryType)
	return userSteamID, data[:result], chatEntryType
}

func RequestLobbyData(lobbySteamID Uint64SteamID) bool {
	return iSteamMatchmaking_RequestLobbyData(matchmaking(), lobbySteamID)
}

func SetLobbyGameServer(lobbySteamID Uint64SteamID, gameServerIP uint32, gameServerPort uint16, gameServerSteamID Uint64SteamID) {
	iSteamMatchmaking_SetLobbyGameServer(matchmaking(), lobbySteamID, gameServerIP, gameServerPort, gameServerSteamID)
}

func GetLobbyGameServer(lobbySteamID Uint64SteamID) (gameServerIP uint32, gameServerPort uint16, gameServerSteamID CSteamID, success bool) {
	success = iSteamMatchmaking_GetLobbyGameServer(matchmaking(), lobbySteamID, &gameServerIP, &gameServerPort, &gameServerSteamID)
	return gameServerIP, gameServerPort, gameServerSteamID, success
}

func SetLobbyMemberLimit(lobbySteamID Uint64SteamID, cMaxMembers int32) bool {
	return iSteamMatchmaking_SetLobbyMemberLimit(matchmaking(), lobbySteamID, cMaxMembers)
}

func GetLobbyMemberLimit(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetLobbyMemberLimit(matchmaking(), lobbySteamID)
}

func SetLobbyType(lobbySteamID Uint64SteamID, lobbyType ELobbyType) bool {
	return iSteamMatchmaking_SetLobbyType(matchmaking(), lobbySteamID, lobbyType)
}

func SetLobbyJoinable(lobbySteamID Uint64SteamID, joinableLobby bool) bool {
	return iSteamMatchmaking_SetLobbyJoinable(matchmaking(), lobbySteamID, joinableLobby)
}

func GetLobbyOwner(lobbySteamID Uint64SteamID) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyOwner(matchmaking(), lobbySteamID)
}

func SetLobbyOwner(lobbySteamID Uint64SteamID, newUserSteamID Uint64SteamID) bool {
	return iSteamMatchmaking_SetLobbyOwner(matchmaking(), lobbySteamID, newUserSteamID)
}

func SetLinkedLobby(lobbySteamID Uint64SteamID, dependentLobbySteamID Uint64SteamID) bool {
	return iSteamMatchmaking_SetLinkedLobby(matchmaking(), lobbySteamID, dependentLobbySteamID)
}
