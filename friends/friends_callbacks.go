package friends

import . "github.com/assemblaj/purego-steamworks"

const (
	PersonaStateChangeID               SteamCallbackID = 304
	GameOverlayActivatedID             SteamCallbackID = 331
	GameServerChangeRequestedID        SteamCallbackID = 332
	GameLobbyJoinRequestedID           SteamCallbackID = 333
	AvatarImageLoadedID                SteamCallbackID = 334
	ClanOfficerListResponseID          SteamCallbackID = 335
	FriendRichPresenceUpdateID         SteamCallbackID = 336
	GameRichPresenceJoinRequestedID    SteamCallbackID = 337
	GameConnectedClanChatMsgID         SteamCallbackID = 338
	GameConnectedChatJoinID            SteamCallbackID = 339
	GameConnectedChatLeaveID           SteamCallbackID = 340
	DownloadClanActivityCountsResultID SteamCallbackID = 341
	JoinClanChatRoomCompletionResultID SteamCallbackID = 342
	GameConnectedFriendChatMsgID       SteamCallbackID = 343
	FriendsGetFollowerCountID          SteamCallbackID = 344
	FriendsIsFollowingID               SteamCallbackID = 345
	FriendsEnumerateFollowingListID    SteamCallbackID = 346
	SetPersonaNameResponseID           SteamCallbackID = 347
	UnreadChatMessagesChangedID        SteamCallbackID = 348
	OverlayBrowserProtocolNavigationID SteamCallbackID = 349
	EquippedProfileItemsChangedID      SteamCallbackID = 350
	EquippedProfileItemsID             SteamCallbackID = 351
)

type EChatRoomEnterResponse int32

const (
	EChatRoomEnterResponseSuccess           EChatRoomEnterResponse = 1
	EChatRoomEnterResponseDoesntExist       EChatRoomEnterResponse = 2
	EChatRoomEnterResponseNotAllowed        EChatRoomEnterResponse = 3
	EChatRoomEnterResponseFull              EChatRoomEnterResponse = 4
	EChatRoomEnterResponseError             EChatRoomEnterResponse = 5
	EChatRoomEnterResponseBanned            EChatRoomEnterResponse = 6
	EChatRoomEnterResponseLimited           EChatRoomEnterResponse = 7
	EChatRoomEnterResponseClanDisabled      EChatRoomEnterResponse = 8
	EChatRoomEnterResponseCommunityBan      EChatRoomEnterResponse = 9
	EChatRoomEnterResponseMemberBlockedYou  EChatRoomEnterResponse = 10
	EChatRoomEnterResponseYouBlockedMember  EChatRoomEnterResponse = 11
	EChatRoomEnterResponseRatelimitExceeded EChatRoomEnterResponse = 15
)

type PersonaStateChange struct {
	SteamID     uint64
	ChangeFlags int32
}
type GameOverlayActivated struct {
	Active        uint8
	UserInitiated bool
	_             [2]byte
	AppID         AppId
	OverlayPID    uint32
}
type GameServerChangeRequested struct {
	Server   [64]byte
	Password [64]byte
}
type GameLobbyJoinRequested struct {
	SteamIDLobby  CSteamID
	SteamIDFriend CSteamID
}
type AvatarImageLoaded struct {
	SteamID CSteamID
	Image   int32
	Wide    int32
	Tall    int32
}
type ClanOfficerListResponse struct {
	SteamIDClan CSteamID
	Officers    int32
	Success     uint8
}
type FriendRichPresenceUpdate struct {
	SteamIDFriend CSteamID
	AppID         AppId
}
type GameRichPresenceJoinRequested struct {
	SteamIDFriend CSteamID
	Connect       [256]byte
}
type GameConnectedClanChatMsg struct {
	SteamIDClanChat CSteamID
	SteamIDUser     CSteamID
	MessageID       int32
}
type GameConnectedChatJoin struct {
	SteamIDClanChat CSteamID
	SteamIDUser     CSteamID
}
type GameConnectedChatLeave struct {
	SteamIDClanChat CSteamID
	SteamIDUser     CSteamID
	Kicked          bool
	Dropped         bool
	_               [6]byte
}
type DownloadClanActivityCountsResult struct {
	Success bool
}
type JoinClanChatRoomCompletionResult struct {
	SteamIDClanChat       CSteamID
	ChatRoomEnterResponse EChatRoomEnterResponse
}
type GameConnectedFriendChatMsg struct {
	SteamIDUser CSteamID
	MessageID   int32
}
type FriendsGetFollowerCount struct {
	Result  EResult
	SteamID CSteamID
	Count   int32
}
type FriendsIsFollowing struct {
	Result      EResult
	SteamID     CSteamID
	IsFollowing bool
}
type FriendsEnumerateFollowingList struct {
	Result           EResult
	SteamID          [50]CSteamID
	ResultsReturned  int32
	TotalResultCount int32
}
type SetPersonaNameResponse struct {
	Success      bool
	LocalSuccess bool
	Result       EResult
}
type UnreadChatMessagesChanged struct {
}
type OverlayBrowserProtocolNavigation struct {
	URI [1024]byte
}
type EquippedProfileItemsChanged struct {
	SteamID CSteamID
}
type EquippedProfileItems struct {
	Result                   EResult
	SteamID                  CSteamID
	HasAnimatedAvatar        bool
	HasAvatarFrame           bool
	HasProfileModifier       bool
	HasProfileBackground     bool
	HasMiniProfileBackground bool
	FromCache                bool
	_                        [1]byte
}

func (cb PersonaStateChange) CallbackID() SteamCallbackID {
	return PersonaStateChangeID
}

func (cb PersonaStateChange) Name() string {
	return "PersonaStateChange"
}

func (cb PersonaStateChange) String() string {
	return CallbackString(cb)
}

func (cb GameOverlayActivated) CallbackID() SteamCallbackID {
	return GameOverlayActivatedID
}

func (cb GameOverlayActivated) Name() string {
	return "GameOverlayActivated"
}

func (cb GameOverlayActivated) String() string {
	return CallbackString(cb)
}

func (cb GameServerChangeRequested) CallbackID() SteamCallbackID {
	return GameServerChangeRequestedID
}

func (cb GameServerChangeRequested) Name() string {
	return "GameServerChangeRequested"
}

func (cb GameServerChangeRequested) String() string {
	return CallbackString(cb)
}

func (cb GameLobbyJoinRequested) CallbackID() SteamCallbackID {
	return GameLobbyJoinRequestedID
}

func (cb GameLobbyJoinRequested) Name() string {
	return "GameLobbyJoinRequested"
}

func (cb GameLobbyJoinRequested) String() string {
	return CallbackString(cb)
}

func (cb AvatarImageLoaded) CallbackID() SteamCallbackID {
	return AvatarImageLoadedID
}

func (cb AvatarImageLoaded) Name() string {
	return "AvatarImageLoaded"
}

func (cb AvatarImageLoaded) String() string {
	return CallbackString(cb)
}

func (cb ClanOfficerListResponse) CallbackID() SteamCallbackID {
	return ClanOfficerListResponseID
}

func (cb ClanOfficerListResponse) Name() string {
	return "ClanOfficerListResponse"
}

func (cb ClanOfficerListResponse) String() string {
	return CallbackString(cb)
}

func (cb FriendRichPresenceUpdate) CallbackID() SteamCallbackID {
	return FriendRichPresenceUpdateID
}

func (cb FriendRichPresenceUpdate) Name() string {
	return "FriendRichPresenceUpdate"
}

func (cb FriendRichPresenceUpdate) String() string {
	return CallbackString(cb)
}

func (cb GameRichPresenceJoinRequested) CallbackID() SteamCallbackID {
	return GameRichPresenceJoinRequestedID
}

func (cb GameRichPresenceJoinRequested) Name() string {
	return "GameRichPresenceJoinRequested"
}

func (cb GameRichPresenceJoinRequested) String() string {
	return CallbackString(cb)
}

func (cb GameConnectedClanChatMsg) CallbackID() SteamCallbackID {
	return GameConnectedClanChatMsgID
}

func (cb GameConnectedClanChatMsg) Name() string {
	return "GameConnectedClanChatMsg"
}

func (cb GameConnectedClanChatMsg) String() string {
	return CallbackString(cb)
}

func (cb GameConnectedChatJoin) CallbackID() SteamCallbackID {
	return GameConnectedChatJoinID
}

func (cb GameConnectedChatJoin) Name() string {
	return "GameConnectedChatJoin"
}

func (cb GameConnectedChatJoin) String() string {
	return CallbackString(cb)
}

func (cb GameConnectedChatLeave) CallbackID() SteamCallbackID {
	return GameConnectedChatLeaveID
}

func (cb GameConnectedChatLeave) Name() string {
	return "GameConnectedChatLeave"
}

func (cb GameConnectedChatLeave) String() string {
	return CallbackString(cb)
}

func (cb DownloadClanActivityCountsResult) CallbackID() SteamCallbackID {
	return DownloadClanActivityCountsResultID
}

func (cb DownloadClanActivityCountsResult) Name() string {
	return "DownloadClanActivityCountsResult"
}

func (cb DownloadClanActivityCountsResult) String() string {
	return CallbackString(cb)
}

func (cb JoinClanChatRoomCompletionResult) CallbackID() SteamCallbackID {
	return JoinClanChatRoomCompletionResultID
}

func (cb JoinClanChatRoomCompletionResult) Name() string {
	return "JoinClanChatRoomCompletionResult"
}

func (cb JoinClanChatRoomCompletionResult) String() string {
	return CallbackString(cb)
}

func (cb GameConnectedFriendChatMsg) CallbackID() SteamCallbackID {
	return GameConnectedFriendChatMsgID
}

func (cb GameConnectedFriendChatMsg) Name() string {
	return "GameConnectedFriendChatMsg"
}

func (cb GameConnectedFriendChatMsg) String() string {
	return CallbackString(cb)
}

func (cb FriendsGetFollowerCount) CallbackID() SteamCallbackID {
	return FriendsGetFollowerCountID
}

func (cb FriendsGetFollowerCount) Name() string {
	return "FriendsGetFollowerCount"
}

func (cb FriendsGetFollowerCount) String() string {
	return CallbackString(cb)
}

func (cb FriendsIsFollowing) CallbackID() SteamCallbackID {
	return FriendsIsFollowingID
}

func (cb FriendsIsFollowing) Name() string {
	return "FriendsIsFollowing"
}

func (cb FriendsIsFollowing) String() string {
	return CallbackString(cb)
}

func (cb FriendsEnumerateFollowingList) CallbackID() SteamCallbackID {
	return FriendsEnumerateFollowingListID
}

func (cb FriendsEnumerateFollowingList) Name() string {
	return "FriendsEnumerateFollowingList"
}

func (cb FriendsEnumerateFollowingList) String() string {
	return CallbackString(cb)
}

func (cb SetPersonaNameResponse) CallbackID() SteamCallbackID {
	return SetPersonaNameResponseID
}

func (cb SetPersonaNameResponse) Name() string {
	return "SetPersonaNameResponse"
}

func (cb SetPersonaNameResponse) String() string {
	return CallbackString(cb)
}

func (cb UnreadChatMessagesChanged) CallbackID() SteamCallbackID {
	return UnreadChatMessagesChangedID
}

func (cb UnreadChatMessagesChanged) Name() string {
	return "UnreadChatMessagesChanged"
}

func (cb UnreadChatMessagesChanged) String() string {
	return CallbackString(cb)
}

func (cb OverlayBrowserProtocolNavigation) CallbackID() SteamCallbackID {
	return OverlayBrowserProtocolNavigationID
}

func (cb OverlayBrowserProtocolNavigation) Name() string {
	return "OverlayBrowserProtocolNavigation"
}

func (cb OverlayBrowserProtocolNavigation) String() string {
	return CallbackString(cb)
}

func (cb EquippedProfileItemsChanged) CallbackID() SteamCallbackID {
	return EquippedProfileItemsChangedID
}

func (cb EquippedProfileItemsChanged) Name() string {
	return "EquippedProfileItemsChanged"
}

func (cb EquippedProfileItemsChanged) String() string {
	return CallbackString(cb)
}

func (cb EquippedProfileItems) CallbackID() SteamCallbackID {
	return EquippedProfileItemsID
}

func (cb EquippedProfileItems) Name() string {
	return "EquippedProfileItems"
}

func (cb EquippedProfileItems) String() string {
	return CallbackString(cb)
}
