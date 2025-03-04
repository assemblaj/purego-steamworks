package matchmaking

import . "github.com/assemblaj/purego-steamworks"

const (
	FavoritesListChangedID         SteamCallbackID = 502
	LobbyInviteID                  SteamCallbackID = 503
	LobbyEnterID                   SteamCallbackID = 504
	LobbyDataUpdateID              SteamCallbackID = 505
	LobbyChatUpdateID              SteamCallbackID = 506
	LobbyChatMsgID                 SteamCallbackID = 507
	LobbyGameCreatedID             SteamCallbackID = 509
	LobbyMatchListID               SteamCallbackID = 510
	LobbyKickedID                  SteamCallbackID = 512
	LobbyCreatedID                 SteamCallbackID = 513
	PSNGameBootInviteResultID      SteamCallbackID = 515
	FavoritesListAccountsUpdatedID SteamCallbackID = 516
)

type FavoritesListChanged struct {
	IP        uint32
	QueryPort uint32
	ConnPort  uint32
	AppID     uint32
	Flags     uint32
	Add       bool
	_         [3]byte
	AccountId AccountID
}
type LobbyInvite struct {
	SteamIDUser  uint64
	SteamIDLobby uint64
	GameID       uint64
}
type LobbyEnter struct {
	SteamIDLobby          uint64
	ChatPermissions       uint32
	Locked                bool
	_                     [3]byte
	ChatRoomEnterResponse uint32
}
type LobbyDataUpdate struct {
	SteamIDLobby  uint64
	SteamIDMember uint64
	Success       uint8
}
type LobbyChatUpdate struct {
	SteamIDLobby          uint64
	SteamIDUserChanged    uint64
	SteamIDMakingChange   uint64
	ChatMemberStateChange uint32
}
type LobbyChatMsg struct {
	SteamIDLobby  uint64
	SteamIDUser   uint64
	ChatEntryType uint8
	ChatID        uint32
}
type LobbyGameCreated struct {
	SteamIDLobby      uint64
	SteamIDGameServer uint64
	IP                uint32
	Port              uint16
	_                 [2]byte
}
type LobbyMatchList struct {
	LobbiesMatching uint32
}
type LobbyKicked struct {
	SteamIDLobby          uint64
	SteamIDAdmin          uint64
	KickedDueToDisconnect uint8
}
type LobbyCreated struct {
	Result       EResult
	SteamIDLobby uint64
}
type PSNGameBootInviteResult struct {
	GameBootInviteExists bool
	SteamIDLobby         CSteamID
}
type FavoritesListAccountsUpdated struct {
	Result EResult
}

func (cb FavoritesListChanged) CallbackID() SteamCallbackID {
	return FavoritesListChangedID
}

func (cb FavoritesListChanged) Name() string {
	return "FavoritesListChanged"
}

func (cb FavoritesListChanged) String() string {
	return CallbackString(cb)
}

func (cb LobbyInvite) CallbackID() SteamCallbackID {
	return LobbyInviteID
}

func (cb LobbyInvite) Name() string {
	return "LobbyInvite"
}

func (cb LobbyInvite) String() string {
	return CallbackString(cb)
}

func (cb LobbyEnter) CallbackID() SteamCallbackID {
	return LobbyEnterID
}

func (cb LobbyEnter) Name() string {
	return "LobbyEnter"
}

func (cb LobbyEnter) String() string {
	return CallbackString(cb)
}

func (cb LobbyDataUpdate) CallbackID() SteamCallbackID {
	return LobbyDataUpdateID
}

func (cb LobbyDataUpdate) Name() string {
	return "LobbyDataUpdate"
}

func (cb LobbyDataUpdate) String() string {
	return CallbackString(cb)
}

func (cb LobbyChatUpdate) CallbackID() SteamCallbackID {
	return LobbyChatUpdateID
}

func (cb LobbyChatUpdate) Name() string {
	return "LobbyChatUpdate"
}

func (cb LobbyChatUpdate) String() string {
	return CallbackString(cb)
}

func (cb LobbyChatMsg) CallbackID() SteamCallbackID {
	return LobbyChatMsgID
}

func (cb LobbyChatMsg) Name() string {
	return "LobbyChatMsg"
}

func (cb LobbyChatMsg) String() string {
	return CallbackString(cb)
}

func (cb LobbyGameCreated) CallbackID() SteamCallbackID {
	return LobbyGameCreatedID
}

func (cb LobbyGameCreated) Name() string {
	return "LobbyGameCreated"
}

func (cb LobbyGameCreated) String() string {
	return CallbackString(cb)
}

func (cb LobbyMatchList) CallbackID() SteamCallbackID {
	return LobbyMatchListID
}

func (cb LobbyMatchList) Name() string {
	return "LobbyMatchList"
}

func (cb LobbyMatchList) String() string {
	return CallbackString(cb)
}

func (cb LobbyKicked) CallbackID() SteamCallbackID {
	return LobbyKickedID
}

func (cb LobbyKicked) Name() string {
	return "LobbyKicked"
}

func (cb LobbyKicked) String() string {
	return CallbackString(cb)
}

func (cb LobbyCreated) CallbackID() SteamCallbackID {
	return LobbyCreatedID
}

func (cb LobbyCreated) Name() string {
	return "LobbyCreated"
}

func (cb LobbyCreated) String() string {
	return CallbackString(cb)
}

func (cb PSNGameBootInviteResult) CallbackID() SteamCallbackID {
	return PSNGameBootInviteResultID
}

func (cb PSNGameBootInviteResult) Name() string {
	return "PSNGameBootInviteResult"
}

func (cb PSNGameBootInviteResult) String() string {
	return CallbackString(cb)
}

func (cb FavoritesListAccountsUpdated) CallbackID() SteamCallbackID {
	return FavoritesListAccountsUpdatedID
}

func (cb FavoritesListAccountsUpdated) Name() string {
	return "FavoritesListAccountsUpdated"
}

func (cb FavoritesListAccountsUpdated) String() string {
	return CallbackString(cb)
}
