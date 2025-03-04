package puregosteamworks

// Steam Apps Callbacks
const (
	DlcInstalledID                  SteamCallbackID = 1005
	NewUrlLaunchParametersID        SteamCallbackID = 1014
	AppProofOfPurchaseKeyResponseID SteamCallbackID = 1021
	FileDetailsResultID             SteamCallbackID = 1023
	TimedTrialStatusID              SteamCallbackID = 1030
)

type DlcInstalled struct {
	AppID AppId
}
type NewUrlLaunchParameters struct {
}
type AppProofOfPurchaseKeyResponse struct {
	Result    EResult
	AppID     uint32
	KeyLength uint32
	Key       [240]byte
}

type FileDetailsResult struct {
	Result   EResult
	_        [4]byte
	FileSize uint64
	FileSHA  [20]uint8
	Flags    uint32
}

type TimedTrialStatus struct {
	AppID          AppId
	IsOffline      bool
	_              [3]byte
	SecondsAllowed uint32
	SecondsPlayed  uint32
}

func (cb DlcInstalled) CallbackID() SteamCallbackID {
	return DlcInstalledID
}

func (cb DlcInstalled) Name() string {
	return "DlcInstalled"
}

func (cb DlcInstalled) String() string {
	return callbackString(cb)
}

func (cb NewUrlLaunchParameters) CallbackID() SteamCallbackID {
	return NewUrlLaunchParametersID
}

func (cb NewUrlLaunchParameters) Name() string {
	return "NewUrlLaunchParameters"
}

func (cb NewUrlLaunchParameters) String() string {
	return callbackString(cb)
}

func (cb AppProofOfPurchaseKeyResponse) CallbackID() SteamCallbackID {
	return AppProofOfPurchaseKeyResponseID
}

func (cb AppProofOfPurchaseKeyResponse) Name() string {
	return "AppProofOfPurchaseKeyResponse"
}

func (cb AppProofOfPurchaseKeyResponse) String() string {
	return callbackString(cb)
}

func (cb FileDetailsResult) CallbackID() SteamCallbackID {
	return FileDetailsResultID
}

func (cb FileDetailsResult) Name() string {
	return "FileDetailsResult"
}

func (cb FileDetailsResult) String() string {
	return callbackString(cb)
}

func (cb TimedTrialStatus) CallbackID() SteamCallbackID {
	return TimedTrialStatusID
}

func (cb TimedTrialStatus) Name() string {
	return "TimedTrialStatus"
}

func (cb TimedTrialStatus) String() string {
	return callbackString(cb)
}

// Steam Friends Callbacks
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
	return callbackString(cb)
}

func (cb GameOverlayActivated) CallbackID() SteamCallbackID {
	return GameOverlayActivatedID
}

func (cb GameOverlayActivated) Name() string {
	return "GameOverlayActivated"
}

func (cb GameOverlayActivated) String() string {
	return callbackString(cb)
}

func (cb GameServerChangeRequested) CallbackID() SteamCallbackID {
	return GameServerChangeRequestedID
}

func (cb GameServerChangeRequested) Name() string {
	return "GameServerChangeRequested"
}

func (cb GameServerChangeRequested) String() string {
	return callbackString(cb)
}

func (cb GameLobbyJoinRequested) CallbackID() SteamCallbackID {
	return GameLobbyJoinRequestedID
}

func (cb GameLobbyJoinRequested) Name() string {
	return "GameLobbyJoinRequested"
}

func (cb GameLobbyJoinRequested) String() string {
	return callbackString(cb)
}

func (cb AvatarImageLoaded) CallbackID() SteamCallbackID {
	return AvatarImageLoadedID
}

func (cb AvatarImageLoaded) Name() string {
	return "AvatarImageLoaded"
}

func (cb AvatarImageLoaded) String() string {
	return callbackString(cb)
}

func (cb ClanOfficerListResponse) CallbackID() SteamCallbackID {
	return ClanOfficerListResponseID
}

func (cb ClanOfficerListResponse) Name() string {
	return "ClanOfficerListResponse"
}

func (cb ClanOfficerListResponse) String() string {
	return callbackString(cb)
}

func (cb FriendRichPresenceUpdate) CallbackID() SteamCallbackID {
	return FriendRichPresenceUpdateID
}

func (cb FriendRichPresenceUpdate) Name() string {
	return "FriendRichPresenceUpdate"
}

func (cb FriendRichPresenceUpdate) String() string {
	return callbackString(cb)
}

func (cb GameRichPresenceJoinRequested) CallbackID() SteamCallbackID {
	return GameRichPresenceJoinRequestedID
}

func (cb GameRichPresenceJoinRequested) Name() string {
	return "GameRichPresenceJoinRequested"
}

func (cb GameRichPresenceJoinRequested) String() string {
	return callbackString(cb)
}

func (cb GameConnectedClanChatMsg) CallbackID() SteamCallbackID {
	return GameConnectedClanChatMsgID
}

func (cb GameConnectedClanChatMsg) Name() string {
	return "GameConnectedClanChatMsg"
}

func (cb GameConnectedClanChatMsg) String() string {
	return callbackString(cb)
}

func (cb GameConnectedChatJoin) CallbackID() SteamCallbackID {
	return GameConnectedChatJoinID
}

func (cb GameConnectedChatJoin) Name() string {
	return "GameConnectedChatJoin"
}

func (cb GameConnectedChatJoin) String() string {
	return callbackString(cb)
}

func (cb GameConnectedChatLeave) CallbackID() SteamCallbackID {
	return GameConnectedChatLeaveID
}

func (cb GameConnectedChatLeave) Name() string {
	return "GameConnectedChatLeave"
}

func (cb GameConnectedChatLeave) String() string {
	return callbackString(cb)
}

func (cb DownloadClanActivityCountsResult) CallbackID() SteamCallbackID {
	return DownloadClanActivityCountsResultID
}

func (cb DownloadClanActivityCountsResult) Name() string {
	return "DownloadClanActivityCountsResult"
}

func (cb DownloadClanActivityCountsResult) String() string {
	return callbackString(cb)
}

func (cb JoinClanChatRoomCompletionResult) CallbackID() SteamCallbackID {
	return JoinClanChatRoomCompletionResultID
}

func (cb JoinClanChatRoomCompletionResult) Name() string {
	return "JoinClanChatRoomCompletionResult"
}

func (cb JoinClanChatRoomCompletionResult) String() string {
	return callbackString(cb)
}

func (cb GameConnectedFriendChatMsg) CallbackID() SteamCallbackID {
	return GameConnectedFriendChatMsgID
}

func (cb GameConnectedFriendChatMsg) Name() string {
	return "GameConnectedFriendChatMsg"
}

func (cb GameConnectedFriendChatMsg) String() string {
	return callbackString(cb)
}

func (cb FriendsGetFollowerCount) CallbackID() SteamCallbackID {
	return FriendsGetFollowerCountID
}

func (cb FriendsGetFollowerCount) Name() string {
	return "FriendsGetFollowerCount"
}

func (cb FriendsGetFollowerCount) String() string {
	return callbackString(cb)
}

func (cb FriendsIsFollowing) CallbackID() SteamCallbackID {
	return FriendsIsFollowingID
}

func (cb FriendsIsFollowing) Name() string {
	return "FriendsIsFollowing"
}

func (cb FriendsIsFollowing) String() string {
	return callbackString(cb)
}

func (cb FriendsEnumerateFollowingList) CallbackID() SteamCallbackID {
	return FriendsEnumerateFollowingListID
}

func (cb FriendsEnumerateFollowingList) Name() string {
	return "FriendsEnumerateFollowingList"
}

func (cb FriendsEnumerateFollowingList) String() string {
	return callbackString(cb)
}

func (cb SetPersonaNameResponse) CallbackID() SteamCallbackID {
	return SetPersonaNameResponseID
}

func (cb SetPersonaNameResponse) Name() string {
	return "SetPersonaNameResponse"
}

func (cb SetPersonaNameResponse) String() string {
	return callbackString(cb)
}

func (cb UnreadChatMessagesChanged) CallbackID() SteamCallbackID {
	return UnreadChatMessagesChangedID
}

func (cb UnreadChatMessagesChanged) Name() string {
	return "UnreadChatMessagesChanged"
}

func (cb UnreadChatMessagesChanged) String() string {
	return callbackString(cb)
}

func (cb OverlayBrowserProtocolNavigation) CallbackID() SteamCallbackID {
	return OverlayBrowserProtocolNavigationID
}

func (cb OverlayBrowserProtocolNavigation) Name() string {
	return "OverlayBrowserProtocolNavigation"
}

func (cb OverlayBrowserProtocolNavigation) String() string {
	return callbackString(cb)
}

func (cb EquippedProfileItemsChanged) CallbackID() SteamCallbackID {
	return EquippedProfileItemsChangedID
}

func (cb EquippedProfileItemsChanged) Name() string {
	return "EquippedProfileItemsChanged"
}

func (cb EquippedProfileItemsChanged) String() string {
	return callbackString(cb)
}

func (cb EquippedProfileItems) CallbackID() SteamCallbackID {
	return EquippedProfileItemsID
}

func (cb EquippedProfileItems) Name() string {
	return "EquippedProfileItems"
}

func (cb EquippedProfileItems) String() string {
	return callbackString(cb)
}

// Steam Game Search Callbacks
const (
	SearchForGameProgressCallbackID            SteamCallbackID = 5201
	SearchForGameResultCallbackID              SteamCallbackID = 5202
	RequestPlayersForGameProgressCallbackID    SteamCallbackID = 5211
	RequestPlayersForGameResultCallbackID      SteamCallbackID = 5212
	RequestPlayersForGameFinalResultCallbackID SteamCallbackID = 5213
	SubmitPlayerResultResultCallbackID         SteamCallbackID = 5214
	EndGameResultCallbackID                    SteamCallbackID = 5215
)

type SearchForGameProgressCallback struct {
	SearchID                 uint64
	Result                   EResult
	LobbyID                  CSteamID
	SteamIDEndedSearch       CSteamID
	SecondsRemainingEstimate int32
	PlayersSearching         int32
}
type SearchForGameResultCallback struct {
	SearchID           uint64
	Result             EResult
	CountPlayersInGame int32
	CountAcceptedGame  int32
	SteamIDHost        CSteamID
	FinalCallback      bool
	_                  [7]byte
}

type RequestPlayersForGameProgressCallback struct {
	Result   EResult
	SearchID uint64
}
type RequestPlayersForGameResultCallback struct {
	Result                   EResult
	_                        [4]byte
	SearchID                 uint64
	SteamIDPlayerFound       CSteamID
	SteamIDLobby             CSteamID
	PlayerAcceptState        PlayerAcceptState
	PlayerIndex              int32
	TotalPlayersFound        int32
	TotalPlayersAcceptedGame int32
	SuggestedTeamIndex       int32
	_                        [4]byte
	UniqueGameID             uint64
}
type PlayerAcceptState int32

const (
	EStateUnknown        PlayerAcceptState = 0
	EStatePlayerAccepted PlayerAcceptState = 1
	EStatePlayerDeclined PlayerAcceptState = 2
)

type RequestPlayersForGameFinalResultCallback struct {
	Result       EResult
	SearchID     uint64
	UniqueGameID uint64
}
type SubmitPlayerResultResultCallback struct {
	Result        EResult
	UniqueGameID  uint64
	SteamIDPlayer CSteamID
}
type EndGameResultCallback struct {
	Result       EResult
	UniqueGameID uint64
}

func (cb SearchForGameProgressCallback) CallbackID() SteamCallbackID {
	return SearchForGameProgressCallbackID
}

func (cb SearchForGameProgressCallback) Name() string {
	return "SearchForGameProgressCallback"
}

func (cb SearchForGameProgressCallback) String() string {
	return callbackString(cb)
}

func (cb SearchForGameResultCallback) CallbackID() SteamCallbackID {
	return SearchForGameResultCallbackID
}

func (cb SearchForGameResultCallback) Name() string {
	return "SearchForGameResultCallback"
}

func (cb SearchForGameResultCallback) String() string {
	return callbackString(cb)
}

func (cb RequestPlayersForGameProgressCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameProgressCallbackID
}

func (cb RequestPlayersForGameProgressCallback) Name() string {
	return "RequestPlayersForGameProgressCallback"
}

func (cb RequestPlayersForGameProgressCallback) String() string {
	return callbackString(cb)
}

func (cb RequestPlayersForGameResultCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameResultCallbackID
}

func (cb RequestPlayersForGameResultCallback) Name() string {
	return "RequestPlayersForGameResultCallback"
}

func (cb RequestPlayersForGameResultCallback) String() string {
	return callbackString(cb)
}

func (cb RequestPlayersForGameFinalResultCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameFinalResultCallbackID
}

func (cb RequestPlayersForGameFinalResultCallback) Name() string {
	return "RequestPlayersForGameFinalResultCallback"
}

func (cb RequestPlayersForGameFinalResultCallback) String() string {
	return callbackString(cb)
}

func (cb SubmitPlayerResultResultCallback) CallbackID() SteamCallbackID {
	return SubmitPlayerResultResultCallbackID
}

func (cb SubmitPlayerResultResultCallback) Name() string {
	return "SubmitPlayerResultResultCallback"
}

func (cb SubmitPlayerResultResultCallback) String() string {
	return callbackString(cb)
}

func (cb EndGameResultCallback) CallbackID() SteamCallbackID {
	return EndGameResultCallbackID
}

func (cb EndGameResultCallback) Name() string {
	return "EndGameResultCallback"
}

func (cb EndGameResultCallback) String() string {
	return callbackString(cb)
}

// Steam Game Server Callbacks

const (
	GSClientApproveID                     SteamCallbackID = 201
	GSClientDenyID                        SteamCallbackID = 202
	GSClientKickID                        SteamCallbackID = 203
	GSClientAchievementStatusID           SteamCallbackID = 206
	GSPolicyResponseID                    SteamCallbackID = 115
	GSGameplayStatsID                     SteamCallbackID = 207
	GSClientGroupStatusID                 SteamCallbackID = 208
	GSReputationID                        SteamCallbackID = 209
	AssociateWithClanResultID             SteamCallbackID = 210
	ComputeNewPlayerCompatibilityResultID SteamCallbackID = 211
)

type GSClientApprove struct {
	SteamID      CSteamID
	OwnerSteamID CSteamID
}
type GSClientDeny struct {
	SteamID      CSteamID
	DenyReason   EDenyReason
	OptionalText [128]byte
}
type GSClientKick struct {
	SteamID    CSteamID
	DenyReason EDenyReason
}
type GSClientAchievementStatus struct {
	SteamID     uint64
	Achievement [128]byte
	Unlocked    bool
}
type GSPolicyResponse struct {
	Secure uint8
}
type GSGameplayStats struct {
	Result             EResult
	Rank               int32
	TotalConnects      uint32
	TotalMinutesPlayed uint32
}
type GSClientGroupStatus struct {
	SteamIDUser  CSteamID
	SteamIDGroup CSteamID
	Member       bool
	Officer      bool
}
type GSReputation struct {
	Result          EResult
	ReputationScore uint32
	Banned          bool
	BannedIP        uint32
	BannedPort      uint16
	BannedGameID    uint64
	BanExpires      uint32
}
type AssociateWithClanResult struct {
	Result EResult
}
type ComputeNewPlayerCompatibilityResult struct {
	Result                           EResult
	PlayersThatDontLikeCandidate     int32
	PlayersThatCandidateDoesntLike   int32
	ClanPlayersThatDontLikeCandidate int32
	SteamIDCandidate                 CSteamID
}

func (cb GSClientApprove) CallbackID() SteamCallbackID {
	return GSClientApproveID
}

func (cb GSClientApprove) Name() string {
	return "GSClientApprove"
}

func (cb GSClientApprove) String() string {
	return callbackString(cb)
}

func (cb GSClientDeny) CallbackID() SteamCallbackID {
	return GSClientDenyID
}

func (cb GSClientDeny) Name() string {
	return "GSClientDeny"
}

func (cb GSClientDeny) String() string {
	return callbackString(cb)
}

func (cb GSClientKick) CallbackID() SteamCallbackID {
	return GSClientKickID
}

func (cb GSClientKick) Name() string {
	return "GSClientKick"
}

func (cb GSClientKick) String() string {
	return callbackString(cb)
}

func (cb GSClientAchievementStatus) CallbackID() SteamCallbackID {
	return GSClientAchievementStatusID
}

func (cb GSClientAchievementStatus) Name() string {
	return "GSClientAchievementStatus"
}

func (cb GSClientAchievementStatus) String() string {
	return callbackString(cb)
}

func (cb GSPolicyResponse) CallbackID() SteamCallbackID {
	return GSPolicyResponseID
}

func (cb GSPolicyResponse) Name() string {
	return "GSPolicyResponse"
}

func (cb GSPolicyResponse) String() string {
	return callbackString(cb)
}

func (cb GSGameplayStats) CallbackID() SteamCallbackID {
	return GSGameplayStatsID
}

func (cb GSGameplayStats) Name() string {
	return "GSGameplayStats"
}

func (cb GSGameplayStats) String() string {
	return callbackString(cb)
}

func (cb GSClientGroupStatus) CallbackID() SteamCallbackID {
	return GSClientGroupStatusID
}

func (cb GSClientGroupStatus) Name() string {
	return "GSClientGroupStatus"
}

func (cb GSClientGroupStatus) String() string {
	return callbackString(cb)
}

func (cb GSReputation) CallbackID() SteamCallbackID {
	return GSReputationID
}

func (cb GSReputation) Name() string {
	return "GSReputation"
}

func (cb GSReputation) String() string {
	return callbackString(cb)
}

func (cb AssociateWithClanResult) CallbackID() SteamCallbackID {
	return AssociateWithClanResultID
}

func (cb AssociateWithClanResult) Name() string {
	return "AssociateWithClanResult"
}

func (cb AssociateWithClanResult) String() string {
	return callbackString(cb)
}

func (cb ComputeNewPlayerCompatibilityResult) CallbackID() SteamCallbackID {
	return ComputeNewPlayerCompatibilityResultID
}

func (cb ComputeNewPlayerCompatibilityResult) Name() string {
	return "ComputeNewPlayerCompatibilityResult"
}

func (cb ComputeNewPlayerCompatibilityResult) String() string {
	return callbackString(cb)
}

// Steam Game Server Stats Callbacks
const (
	GSStatsReceivedID SteamCallbackID = 1800
	GSStatsStoredID   SteamCallbackID = 1801
	GSStatsUnloadedID SteamCallbackID = 1108
)

type GSStatsReceived struct {
	Result      EResult
	SteamIDUser CSteamID
}
type GSStatsStored struct {
	Result      EResult
	SteamIDUser CSteamID
}
type GSStatsUnloaded struct {
	SteamIDUser CSteamID
}

func (cb GSStatsReceived) CallbackID() SteamCallbackID {
	return GSStatsReceivedID
}

func (cb GSStatsReceived) Name() string {
	return "GSStatsReceived"
}

func (cb GSStatsReceived) String() string {
	return callbackString(cb)
}

func (cb GSStatsStored) CallbackID() SteamCallbackID {
	return GSStatsStoredID
}

func (cb GSStatsStored) Name() string {
	return "GSStatsStored"
}

func (cb GSStatsStored) String() string {
	return callbackString(cb)
}

func (cb GSStatsUnloaded) CallbackID() SteamCallbackID {
	return GSStatsUnloadedID
}

func (cb GSStatsUnloaded) Name() string {
	return "GSStatsUnloaded"
}

func (cb GSStatsUnloaded) String() string {
	return callbackString(cb)
}

// Steam Input Callbacks
const (
	SteamInputDeviceConnectedID     SteamCallbackID = 2801
	SteamInputDeviceDisconnectedID  SteamCallbackID = 2802
	SteamInputConfigurationLoadedID SteamCallbackID = 2803
	SteamInputGamepadSlotChangeID   SteamCallbackID = 2804
)

type SteamInputDeviceConnected struct {
	ConnectedDeviceHandle InputHandle
}
type SteamInputDeviceDisconnected struct {
	DisconnectedDeviceHandle InputHandle
}
type SteamInputConfigurationLoaded struct {
	AppID             AppId
	DeviceHandle      InputHandle
	MappingCreator    CSteamID
	MajorRevision     uint32
	MinorRevision     uint32
	UsesSteamInputAPI bool
	UsesGamepadAPI    bool
	_                 [2]byte
}
type SteamInputGamepadSlotChange struct {
	AppID          AppId
	DeviceHandle   InputHandle
	DeviceType     ESteamInputType
	OldGamepadSlot int32
	NewGamepadSlot int32
}

func (cb SteamInputDeviceConnected) CallbackID() SteamCallbackID {
	return SteamInputDeviceConnectedID
}

func (cb SteamInputDeviceConnected) Name() string {
	return "SteamInputDeviceConnected"
}

func (cb SteamInputDeviceConnected) String() string {
	return callbackString(cb)
}

func (cb SteamInputDeviceDisconnected) CallbackID() SteamCallbackID {
	return SteamInputDeviceDisconnectedID
}

func (cb SteamInputDeviceDisconnected) Name() string {
	return "SteamInputDeviceDisconnected"
}

func (cb SteamInputDeviceDisconnected) String() string {
	return callbackString(cb)
}

func (cb SteamInputConfigurationLoaded) CallbackID() SteamCallbackID {
	return SteamInputConfigurationLoadedID
}

func (cb SteamInputConfigurationLoaded) Name() string {
	return "SteamInputConfigurationLoaded"
}

func (cb SteamInputConfigurationLoaded) String() string {
	return callbackString(cb)
}

func (cb SteamInputGamepadSlotChange) CallbackID() SteamCallbackID {
	return SteamInputGamepadSlotChangeID
}

func (cb SteamInputGamepadSlotChange) Name() string {
	return "SteamInputGamepadSlotChange"
}

func (cb SteamInputGamepadSlotChange) String() string {
	return callbackString(cb)
}

// Steam Inventory Callbacks
const (
	SteamInventoryResultReadyID             SteamCallbackID = 4700
	SteamInventoryFullUpdateID              SteamCallbackID = 4701
	SteamInventoryDefinitionUpdateID        SteamCallbackID = 4702
	SteamInventoryEligiblePromoItemDefIDsID SteamCallbackID = 4703
	SteamInventoryStartPurchaseResultID     SteamCallbackID = 4704
	SteamInventoryRequestPricesResultID     SteamCallbackID = 4705
)

type SteamInventoryResultReady struct {
	Handle SteamInventoryResult
	Result EResult
}
type SteamInventoryFullUpdate struct {
	Handle SteamInventoryResult
}
type SteamInventoryDefinitionUpdate struct {
}
type SteamInventoryEligiblePromoItemDefIDs struct {
	Result                   EResult
	_                        [4]byte
	SteamID                  CSteamID
	NumEligiblePromoItemDefs int32
	CachedData               bool
	_                        [3]byte
}
type SteamInventoryStartPurchaseResult struct {
	Result  EResult
	_       [4]byte
	OrderID uint64
	TransID uint64
}
type SteamInventoryRequestPricesResult struct {
	Result   EResult
	Currency [4]byte
}

func (cb SteamInventoryResultReady) CallbackID() SteamCallbackID {
	return SteamInventoryResultReadyID
}

func (cb SteamInventoryResultReady) Name() string {
	return "SteamInventoryResultReady"
}

func (cb SteamInventoryResultReady) String() string {
	return callbackString(cb)
}

func (cb SteamInventoryFullUpdate) CallbackID() SteamCallbackID {
	return SteamInventoryFullUpdateID
}

func (cb SteamInventoryFullUpdate) Name() string {
	return "SteamInventoryFullUpdate"
}

func (cb SteamInventoryFullUpdate) String() string {
	return callbackString(cb)
}

func (cb SteamInventoryDefinitionUpdate) CallbackID() SteamCallbackID {
	return SteamInventoryDefinitionUpdateID
}

func (cb SteamInventoryDefinitionUpdate) Name() string {
	return "SteamInventoryDefinitionUpdate"
}

func (cb SteamInventoryDefinitionUpdate) String() string {
	return callbackString(cb)
}

func (cb SteamInventoryEligiblePromoItemDefIDs) CallbackID() SteamCallbackID {
	return SteamInventoryEligiblePromoItemDefIDsID
}

func (cb SteamInventoryEligiblePromoItemDefIDs) Name() string {
	return "SteamInventoryEligiblePromoItemDefIDs"
}

func (cb SteamInventoryEligiblePromoItemDefIDs) String() string {
	return callbackString(cb)
}

func (cb SteamInventoryStartPurchaseResult) CallbackID() SteamCallbackID {
	return SteamInventoryStartPurchaseResultID
}

func (cb SteamInventoryStartPurchaseResult) Name() string {
	return "SteamInventoryStartPurchaseResult"
}

func (cb SteamInventoryStartPurchaseResult) String() string {
	return callbackString(cb)
}

func (cb SteamInventoryRequestPricesResult) CallbackID() SteamCallbackID {
	return SteamInventoryRequestPricesResultID
}

func (cb SteamInventoryRequestPricesResult) Name() string {
	return "SteamInventoryRequestPricesResult"
}

func (cb SteamInventoryRequestPricesResult) String() string {
	return callbackString(cb)
}

// Steam Matchmaking Callbacks
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
	return callbackString(cb)
}

func (cb LobbyInvite) CallbackID() SteamCallbackID {
	return LobbyInviteID
}

func (cb LobbyInvite) Name() string {
	return "LobbyInvite"
}

func (cb LobbyInvite) String() string {
	return callbackString(cb)
}

func (cb LobbyEnter) CallbackID() SteamCallbackID {
	return LobbyEnterID
}

func (cb LobbyEnter) Name() string {
	return "LobbyEnter"
}

func (cb LobbyEnter) String() string {
	return callbackString(cb)
}

func (cb LobbyDataUpdate) CallbackID() SteamCallbackID {
	return LobbyDataUpdateID
}

func (cb LobbyDataUpdate) Name() string {
	return "LobbyDataUpdate"
}

func (cb LobbyDataUpdate) String() string {
	return callbackString(cb)
}

func (cb LobbyChatUpdate) CallbackID() SteamCallbackID {
	return LobbyChatUpdateID
}

func (cb LobbyChatUpdate) Name() string {
	return "LobbyChatUpdate"
}

func (cb LobbyChatUpdate) String() string {
	return callbackString(cb)
}

func (cb LobbyChatMsg) CallbackID() SteamCallbackID {
	return LobbyChatMsgID
}

func (cb LobbyChatMsg) Name() string {
	return "LobbyChatMsg"
}

func (cb LobbyChatMsg) String() string {
	return callbackString(cb)
}

func (cb LobbyGameCreated) CallbackID() SteamCallbackID {
	return LobbyGameCreatedID
}

func (cb LobbyGameCreated) Name() string {
	return "LobbyGameCreated"
}

func (cb LobbyGameCreated) String() string {
	return callbackString(cb)
}

func (cb LobbyMatchList) CallbackID() SteamCallbackID {
	return LobbyMatchListID
}

func (cb LobbyMatchList) Name() string {
	return "LobbyMatchList"
}

func (cb LobbyMatchList) String() string {
	return callbackString(cb)
}

func (cb LobbyKicked) CallbackID() SteamCallbackID {
	return LobbyKickedID
}

func (cb LobbyKicked) Name() string {
	return "LobbyKicked"
}

func (cb LobbyKicked) String() string {
	return callbackString(cb)
}

func (cb LobbyCreated) CallbackID() SteamCallbackID {
	return LobbyCreatedID
}

func (cb LobbyCreated) Name() string {
	return "LobbyCreated"
}

func (cb LobbyCreated) String() string {
	return callbackString(cb)
}

func (cb PSNGameBootInviteResult) CallbackID() SteamCallbackID {
	return PSNGameBootInviteResultID
}

func (cb PSNGameBootInviteResult) Name() string {
	return "PSNGameBootInviteResult"
}

func (cb PSNGameBootInviteResult) String() string {
	return callbackString(cb)
}

func (cb FavoritesListAccountsUpdated) CallbackID() SteamCallbackID {
	return FavoritesListAccountsUpdatedID
}

func (cb FavoritesListAccountsUpdated) Name() string {
	return "FavoritesListAccountsUpdated"
}

func (cb FavoritesListAccountsUpdated) String() string {
	return callbackString(cb)
}

// Steam Matchmaking Server Callbacks
type MatchmakingRulesResponse struct{}

func (s MatchmakingRulesResponse) RulesResponded(pchRule string, pchValue string) {
	iSteamMatchmakingRulesResponse_RulesResponded(pchRule, pchValue)
}

func (s MatchmakingRulesResponse) RulesFailedToRespond() {
	iSteamMatchmakingRulesResponse_RulesFailedToRespond()
}

func (s MatchmakingRulesResponse) RulesRefreshComplete() {
	iSteamMatchmakingRulesResponse_RulesRefreshComplete()
}

type ISteamMatchmakingPingResponse struct{}

func (s ISteamMatchmakingPingResponse) ServerResponded(server *GameServerItem) {
	iSteamMatchmakingPingResponse_ServerResponded(server)
}

func (s ISteamMatchmakingPingResponse) ServerFailedToRespond() {
	iSteamMatchmakingPingResponse_ServerFailedToRespond()
}

type MatchmakingPlayersResponse struct{}

func (s MatchmakingPlayersResponse) AddPlayerToList(pchName string, nScore int32, flTimePlayed float32) {
	iSteamMatchmakingPlayersResponse_AddPlayerToList(pchName, nScore, flTimePlayed)
}

func (s MatchmakingPlayersResponse) PlayersFailedToRespond() {
	iSteamMatchmakingPlayersResponse_PlayersFailedToRespond()
}

func (s MatchmakingPlayersResponse) PlayersRefreshComplete() {
	iSteamMatchmakingPlayersResponse_PlayersRefreshComplete()
}

type MatchmakingServerListResponse struct{}

func (s MatchmakingServerListResponse) ServerResponded(hRequest HServerListRequest, iServer int32) {
	iSteamMatchmakingServerListResponse_ServerResponded(hRequest, iServer)
}

func (s MatchmakingServerListResponse) ServerFailedToRespond(hRequest HServerListRequest, iServer int32) {
	iSteamMatchmakingServerListResponse_ServerFailedToRespond(hRequest, iServer)
}

func (s MatchmakingServerListResponse) RefreshComplete(hRequest HServerListRequest, response EMatchMakingServerResponse) {
	iSteamMatchmakingServerListResponse_RefreshComplete(hRequest, response)
}

// Steam Networking Messages Callbacks
const SteamNetworkingMessagesSessionRequestID SteamCallbackID = 1251

type SteamNetworkingMessagesSessionRequest struct {
	IdentityRemote SteamNetworkingIdentity
}

const SteamNetworkingMessagesSessionFailedID SteamCallbackID = 1252

type SteamNetworkingMessagesSessionFailed struct {
	Info SteamNetConnectionInfo
}

func (cb SteamNetworkingMessagesSessionRequest) CallbackID() SteamCallbackID {
	return SteamNetworkingMessagesSessionRequestID
}

func (cb SteamNetworkingMessagesSessionRequest) Name() string {
	return "SteamNetworkingMessagesSessionRequest"
}

func (cb SteamNetworkingMessagesSessionRequest) String() string {
	return callbackString(cb)
}

func (cb SteamNetworkingMessagesSessionFailed) CallbackID() SteamCallbackID {
	return SteamNetworkingMessagesSessionFailedID
}

func (cb SteamNetworkingMessagesSessionFailed) Name() string {
	return "SteamNetworkingMessagesSessionFailed"
}

func (cb SteamNetworkingMessagesSessionFailed) String() string {
	return callbackString(cb)
}

// Steam Networking Sockets Callbacks
const SteamNetAuthenticationStatusID SteamCallbackID = 1222

type SteamNetAuthenticationStatus struct {
	Avail    ESteamNetworkingAvailability
	DebugMsg [256]byte
}

const SteamNetConnectionStatusChangedCallbackID SteamCallbackID = 1221

type SteamNetConnectionStatusChangedCallback struct {
	Conn     HSteamNetConnection
	Info     SteamNetConnectionInfo
	OldState ESteamNetworkingConnectionState
}

const SteamNetworkingFakeIPResultID SteamCallbackID = 1223

type SteamNetworkingFakeIPResult struct {
	Result   EResult
	Identity SteamNetworkingIdentity
	IP       uint32
	Ports    [8]uint16
	_        [4]byte
}

func (cb SteamNetAuthenticationStatus) CallbackID() SteamCallbackID {
	return SteamNetAuthenticationStatusID
}

func (cb SteamNetAuthenticationStatus) Name() string {
	return "SteamNetAuthenticationStatus"
}

func (cb SteamNetAuthenticationStatus) String() string {
	return callbackString(cb)
}

func (cb SteamNetConnectionStatusChangedCallback) CallbackID() SteamCallbackID {
	return SteamNetConnectionStatusChangedCallbackID
}

func (cb SteamNetConnectionStatusChangedCallback) Name() string {
	return "SteamNetConnectionStatusChangedCallback"
}

func (cb SteamNetConnectionStatusChangedCallback) String() string {
	return callbackString(cb)
}

func (cb SteamNetworkingFakeIPResult) CallbackID() SteamCallbackID {
	return SteamNetworkingFakeIPResultID
}

func (cb SteamNetworkingFakeIPResult) Name() string {
	return "SteamNetworkingFakeIPResult"
}

func (cb SteamNetworkingFakeIPResult) String() string {
	return callbackString(cb)
}

// Steam Networking Utils Callbacks
type SteamRelayNetworkStatus struct {
	Avail                     ESteamNetworkingAvailability
	PingMeasurementInProgress int32
	AvailNetworkConfig        ESteamNetworkingAvailability
	AvailAnyRelay             ESteamNetworkingAvailability
	DebugMsg                  [256]byte
}

const SteamRelayNetworkStatusID SteamCallbackID = 1281

func (cb SteamRelayNetworkStatus) CallbackID() SteamCallbackID {
	return SteamRelayNetworkStatusID
}

func (cb SteamRelayNetworkStatus) Name() string {
	return "SteamRelayNetworkStatus"
}

func (cb SteamRelayNetworkStatus) String() string {
	return callbackString(cb)
}

// Steam Parties Callbacks
const (
	JoinPartyCallbackID               SteamCallbackID = 5301
	CreateBeaconCallbackID            SteamCallbackID = 5302
	ReservationNotificationCallbackID SteamCallbackID = 5303
	ChangeNumOpenSlotsCallbackID      SteamCallbackID = 5304
	AvailableBeaconLocationsUpdatedID SteamCallbackID = 5305
	ActiveBeaconsUpdatedID            SteamCallbackID = 5306
)

type JoinPartyCallback struct {
	Result             EResult
	BeaconID           PartyBeaconID
	SteamIDBeaconOwner CSteamID
	ConnectString      [256]byte
}
type CreateBeaconCallback struct {
	Result   EResult
	BeaconID PartyBeaconID
}
type ReservationNotificationCallback struct {
	BeaconID      PartyBeaconID
	SteamIDJoiner CSteamID
}
type ChangeNumOpenSlotsCallback struct {
	Result EResult
}
type AvailableBeaconLocationsUpdated struct {
}
type ActiveBeaconsUpdated struct {
}

func (cb JoinPartyCallback) CallbackID() SteamCallbackID {
	return JoinPartyCallbackID
}

func (cb JoinPartyCallback) Name() string {
	return "JoinPartyCallback"
}

func (cb JoinPartyCallback) String() string {
	return callbackString(cb)
}

func (cb CreateBeaconCallback) CallbackID() SteamCallbackID {
	return CreateBeaconCallbackID
}

func (cb CreateBeaconCallback) Name() string {
	return "CreateBeaconCallback"
}

func (cb CreateBeaconCallback) String() string {
	return callbackString(cb)
}

func (cb ReservationNotificationCallback) CallbackID() SteamCallbackID {
	return ReservationNotificationCallbackID
}

func (cb ReservationNotificationCallback) Name() string {
	return "ReservationNotificationCallback"
}

func (cb ReservationNotificationCallback) String() string {
	return callbackString(cb)
}

func (cb ChangeNumOpenSlotsCallback) CallbackID() SteamCallbackID {
	return ChangeNumOpenSlotsCallbackID
}

func (cb ChangeNumOpenSlotsCallback) Name() string {
	return "ChangeNumOpenSlotsCallback"
}

func (cb ChangeNumOpenSlotsCallback) String() string {
	return callbackString(cb)
}

func (cb AvailableBeaconLocationsUpdated) CallbackID() SteamCallbackID {
	return AvailableBeaconLocationsUpdatedID
}

func (cb AvailableBeaconLocationsUpdated) Name() string {
	return "AvailableBeaconLocationsUpdated"
}

func (cb AvailableBeaconLocationsUpdated) String() string {
	return callbackString(cb)
}

func (cb ActiveBeaconsUpdated) CallbackID() SteamCallbackID {
	return ActiveBeaconsUpdatedID
}

func (cb ActiveBeaconsUpdated) Name() string {
	return "ActiveBeaconsUpdated"
}

func (cb ActiveBeaconsUpdated) String() string {
	return callbackString(cb)
}

// Steam Remote Play Callbacks
const (
	SteamRemotePlaySessionConnectedID    SteamCallbackID = 5701
	SteamRemotePlaySessionDisconnectedID SteamCallbackID = 5702
	SteamRemotePlayTogetherGuestInviteID SteamCallbackID = 5703
)

type SteamRemotePlaySessionConnected struct {
	SessionID RemotePlaySessionID
}
type SteamRemotePlaySessionDisconnected struct {
	SessionID RemotePlaySessionID
}
type SteamRemotePlayTogetherGuestInvite struct {
	ConnectURL [1024]byte
}

func (cb SteamRemotePlaySessionConnected) CallbackID() SteamCallbackID {
	return SteamRemotePlaySessionConnectedID
}

func (cb SteamRemotePlaySessionConnected) Name() string {
	return "SteamRemotePlaySessionConnected"
}

func (cb SteamRemotePlaySessionConnected) String() string {
	return callbackString(cb)
}

func (cb SteamRemotePlaySessionDisconnected) CallbackID() SteamCallbackID {
	return SteamRemotePlaySessionDisconnectedID
}

func (cb SteamRemotePlaySessionDisconnected) Name() string {
	return "SteamRemotePlaySessionDisconnected"
}

func (cb SteamRemotePlaySessionDisconnected) String() string {
	return callbackString(cb)
}

func (cb SteamRemotePlayTogetherGuestInvite) CallbackID() SteamCallbackID {
	return SteamRemotePlayTogetherGuestInviteID
}

func (cb SteamRemotePlayTogetherGuestInvite) Name() string {
	return "SteamRemotePlayTogetherGuestInvite"
}

func (cb SteamRemotePlayTogetherGuestInvite) String() string {
	return callbackString(cb)
}

// Steam Remote Storage Callbacks

const (
	RemoteStorageFileShareResultID                           SteamCallbackID = 1307
	RemoteStoragePublishFileResultID                         SteamCallbackID = 1309
	RemoteStorageDeletePublishedFileResultID                 SteamCallbackID = 1311
	RemoteStorageEnumerateUserPublishedFilesResultID         SteamCallbackID = 1312
	RemoteStorageEnumerateUserSubscribedFilesResultID        SteamCallbackID = 1314
	RemoteStorageUpdatePublishedFileResultID                 SteamCallbackID = 1316
	RemoteStorageDownloadUGCResultID                         SteamCallbackID = 1317
	RemoteStorageGetPublishedFileDetailsResultID             SteamCallbackID = 1318
	RemoteStorageEnumerateWorkshopFilesResultID              SteamCallbackID = 1319
	RemoteStorageGetPublishedItemVoteDetailsResultID         SteamCallbackID = 1320
	RemoteStoragePublishedFileSubscribedID                   SteamCallbackID = 1321
	RemoteStoragePublishedFileUnsubscribedID                 SteamCallbackID = 1322
	RemoteStoragePublishedFileDeletedID                      SteamCallbackID = 1323
	RemoteStorageUpdateUserPublishedItemVoteResultID         SteamCallbackID = 1324
	RemoteStorageUserVoteDetailsID                           SteamCallbackID = 1325
	RemoteStorageEnumerateUserSharedWorkshopFilesResultID    SteamCallbackID = 1326
	RemoteStorageSetUserPublishedFileActionResultID          SteamCallbackID = 1327
	RemoteStorageEnumeratePublishedFilesByUserActionResultID SteamCallbackID = 1328
	RemoteStoragePublishFileProgressID                       SteamCallbackID = 1329
	RemoteStoragePublishedFileUpdatedID                      SteamCallbackID = 1330
	RemoteStorageFileWriteAsyncCompleteID                    SteamCallbackID = 1331
	RemoteStorageFileReadAsyncCompleteID                     SteamCallbackID = 1332
	RemoteStorageLocalFileChangeID                           SteamCallbackID = 1333
)

type RemoteStorageFileShareResult struct {
	Result   EResult
	File     UGCHandle
	Filename [260]byte
}
type RemoteStoragePublishFileResult struct {
	Result                                  EResult
	PublishedFileId                         PublishedFileId
	UserNeedsToAcceptWorkshopLegalAgreement bool
	_                                       [7]byte
}
type RemoteStorageDeletePublishedFileResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
}
type RemoteStorageEnumerateUserPublishedFilesResult struct {
	Result           EResult
	ResultsReturned  int32
	TotalResultCount int32
	PublishedFileId  [50]PublishedFileId
}

type RemoteStorageEnumerateUserSubscribedFilesResult struct {
	Result           EResult
	ResultsReturned  int32
	TotalResultCount int32
	PublishedFileId  [50]PublishedFileId
	TimeSubscribed   [50]uint32
}

type RemoteStorageUpdatePublishedFileResult struct {
	Result                                  EResult
	PublishedFileId                         PublishedFileId
	UserNeedsToAcceptWorkshopLegalAgreement bool
	_                                       [7]byte
}
type RemoteStorageDownloadUGCResult struct {
	Result       EResult
	File         UGCHandle
	AppID        AppId
	SizeInBytes  int32
	FileName     [260]byte
	SteamIDOwner uint64
}
type RemoteStorageGetPublishedFileDetailsResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
	CreatorAppID    AppId
	ConsumerAppID   AppId
	Title           [129]byte
	Description     [8000]byte
	File            UGCHandle
	PreviewFile     UGCHandle
	SteamIDOwner    uint64
	TimeCreated     uint32
	TimeUpdated     uint32
	Visibility      ERemoteStoragePublishedFileVisibility
	Banned          bool
	_               [3]byte
	Tags            [1025]byte
	TagsTruncated   bool
	_               [2]byte
	FileName        [260]byte
	FileSize        int32
	PreviewFileSize int32
	URL             [256]byte
	FileType        EWorkshopFileType
	AcceptedForUse  bool
	_               [3]byte
}
type RemoteStorageEnumerateWorkshopFilesResult struct {
	Result           EResult
	ResultsReturned  int32
	TotalResultCount int32
	PublishedFileId  [50]PublishedFileId
	Score            [50]float32
	AppId            AppId
	StartIndex       uint32
}
type RemoteStorageGetPublishedItemVoteDetailsResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
	VotesFor        int32
	VotesAgainst    int32
	Reports         int32
	Score           float32
}
type RemoteStoragePublishedFileSubscribed struct {
	PublishedFileId PublishedFileId
	AppID           AppId
}
type RemoteStoragePublishedFileUnsubscribed struct {
	PublishedFileId PublishedFileId
	AppID           AppId
}
type RemoteStoragePublishedFileDeleted struct {
	PublishedFileId PublishedFileId
	AppID           AppId
}
type RemoteStorageUpdateUserPublishedItemVoteResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
}
type RemoteStorageUserVoteDetails struct {
	Result          EResult
	PublishedFileId PublishedFileId
	Vote            EWorkshopVote
}
type RemoteStorageEnumerateUserSharedWorkshopFilesResult struct {
	Result           EResult
	ResultsReturned  int32
	TotalResultCount int32
	PublishedFileId  [50]PublishedFileId
}
type RemoteStorageSetUserPublishedFileActionResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
	Action          EWorkshopFileAction
}
type RemoteStorageEnumeratePublishedFilesByUserActionResult struct {
	Result           EResult
	Action           EWorkshopFileAction
	ResultsReturned  int32
	TotalResultCount int32
	PublishedFileId  [50]PublishedFileId
	TimeUpdated      [50]uint32
}
type RemoteStoragePublishFileProgress struct {
	PercentFile float64
	Preview     bool
	_           [7]byte
}
type RemoteStoragePublishedFileUpdated struct {
	PublishedFileId PublishedFileId
	AppID           AppId
	Unused          uint64
}
type RemoteStorageFileWriteAsyncComplete struct {
	Result EResult
}
type RemoteStorageFileReadAsyncComplete struct {
	FileReadAsync SteamAPICall
	Result        EResult
	Offset        uint32
	Read          uint32
}
type RemoteStorageLocalFileChange struct {
}

func (cb RemoteStorageFileShareResult) CallbackID() SteamCallbackID {
	return RemoteStorageFileShareResultID
}

func (cb RemoteStorageFileShareResult) Name() string {
	return "RemoteStorageFileShareResult"
}

func (cb RemoteStorageFileShareResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishFileResult) CallbackID() SteamCallbackID {
	return RemoteStoragePublishFileResultID
}

func (cb RemoteStoragePublishFileResult) Name() string {
	return "RemoteStoragePublishFileResult"
}

func (cb RemoteStoragePublishFileResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageDeletePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageDeletePublishedFileResultID
}

func (cb RemoteStorageDeletePublishedFileResult) Name() string {
	return "RemoteStorageDeletePublishedFileResult"
}

func (cb RemoteStorageDeletePublishedFileResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserPublishedFilesResultID
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) Name() string {
	return "RemoteStorageEnumerateUserPublishedFilesResult"
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserSubscribedFilesResultID
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) Name() string {
	return "RemoteStorageEnumerateUserSubscribedFilesResult"
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageUpdatePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageUpdatePublishedFileResultID
}

func (cb RemoteStorageUpdatePublishedFileResult) Name() string {
	return "RemoteStorageUpdatePublishedFileResult"
}

func (cb RemoteStorageUpdatePublishedFileResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageDownloadUGCResult) CallbackID() SteamCallbackID {
	return RemoteStorageDownloadUGCResultID
}

func (cb RemoteStorageDownloadUGCResult) Name() string {
	return "RemoteStorageDownloadUGCResult"
}

func (cb RemoteStorageDownloadUGCResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageGetPublishedFileDetailsResult) CallbackID() SteamCallbackID {
	return RemoteStorageGetPublishedFileDetailsResultID
}

func (cb RemoteStorageGetPublishedFileDetailsResult) Name() string {
	return "RemoteStorageGetPublishedFileDetailsResult"
}

func (cb RemoteStorageGetPublishedFileDetailsResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateWorkshopFilesResultID
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) Name() string {
	return "RemoteStorageEnumerateWorkshopFilesResult"
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) CallbackID() SteamCallbackID {
	return RemoteStorageGetPublishedItemVoteDetailsResultID
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) Name() string {
	return "RemoteStorageGetPublishedItemVoteDetailsResult"
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishedFileSubscribed) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileSubscribedID
}

func (cb RemoteStoragePublishedFileSubscribed) Name() string {
	return "RemoteStoragePublishedFileSubscribed"
}

func (cb RemoteStoragePublishedFileSubscribed) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishedFileUnsubscribed) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileUnsubscribedID
}

func (cb RemoteStoragePublishedFileUnsubscribed) Name() string {
	return "RemoteStoragePublishedFileUnsubscribed"
}

func (cb RemoteStoragePublishedFileUnsubscribed) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishedFileDeleted) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileDeletedID
}

func (cb RemoteStoragePublishedFileDeleted) Name() string {
	return "RemoteStoragePublishedFileDeleted"
}

func (cb RemoteStoragePublishedFileDeleted) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) CallbackID() SteamCallbackID {
	return RemoteStorageUpdateUserPublishedItemVoteResultID
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) Name() string {
	return "RemoteStorageUpdateUserPublishedItemVoteResult"
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageUserVoteDetails) CallbackID() SteamCallbackID {
	return RemoteStorageUserVoteDetailsID
}

func (cb RemoteStorageUserVoteDetails) Name() string {
	return "RemoteStorageUserVoteDetails"
}

func (cb RemoteStorageUserVoteDetails) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserSharedWorkshopFilesResultID
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) Name() string {
	return "RemoteStorageEnumerateUserSharedWorkshopFilesResult"
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageSetUserPublishedFileActionResult) CallbackID() SteamCallbackID {
	return RemoteStorageSetUserPublishedFileActionResultID
}

func (cb RemoteStorageSetUserPublishedFileActionResult) Name() string {
	return "RemoteStorageSetUserPublishedFileActionResult"
}

func (cb RemoteStorageSetUserPublishedFileActionResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumeratePublishedFilesByUserActionResultID
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) Name() string {
	return "RemoteStorageEnumeratePublishedFilesByUserActionResult"
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishFileProgress) CallbackID() SteamCallbackID {
	return RemoteStoragePublishFileProgressID
}

func (cb RemoteStoragePublishFileProgress) Name() string {
	return "RemoteStoragePublishFileProgress"
}

func (cb RemoteStoragePublishFileProgress) String() string {
	return callbackString(cb)
}

func (cb RemoteStoragePublishedFileUpdated) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileUpdatedID
}

func (cb RemoteStoragePublishedFileUpdated) Name() string {
	return "RemoteStoragePublishedFileUpdated"
}

func (cb RemoteStoragePublishedFileUpdated) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageFileWriteAsyncComplete) CallbackID() SteamCallbackID {
	return RemoteStorageFileWriteAsyncCompleteID
}

func (cb RemoteStorageFileWriteAsyncComplete) Name() string {
	return "RemoteStorageFileWriteAsyncComplete"
}

func (cb RemoteStorageFileWriteAsyncComplete) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageFileReadAsyncComplete) CallbackID() SteamCallbackID {
	return RemoteStorageFileReadAsyncCompleteID
}

func (cb RemoteStorageFileReadAsyncComplete) Name() string {
	return "RemoteStorageFileReadAsyncComplete"
}

func (cb RemoteStorageFileReadAsyncComplete) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageLocalFileChange) CallbackID() SteamCallbackID {
	return RemoteStorageLocalFileChangeID
}

func (cb RemoteStorageLocalFileChange) Name() string {
	return "RemoteStorageLocalFileChange"
}

func (cb RemoteStorageLocalFileChange) String() string {
	return callbackString(cb)
}

const RemoteStorageSubscribePublishedFileResultID SteamCallbackID = 1313

type RemoteStorageSubscribePublishedFileResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
}

const RemoteStorageUnsubscribePublishedFileResultID SteamCallbackID = 1315

type RemoteStorageUnsubscribePublishedFileResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
}

func (cb RemoteStorageUnsubscribePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageUnsubscribePublishedFileResultID
}

func (cb RemoteStorageUnsubscribePublishedFileResult) Name() string {
	return "RemoteStorageUnsubscribePublishedFileResult"
}

func (cb RemoteStorageUnsubscribePublishedFileResult) String() string {
	return callbackString(cb)
}

func (cb RemoteStorageSubscribePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageSubscribePublishedFileResultID
}

func (cb RemoteStorageSubscribePublishedFileResult) Name() string {
	return "RemoteStorageSubscribePublishedFileResult"
}

func (cb RemoteStorageSubscribePublishedFileResult) String() string {
	return callbackString(cb)
}

// Steam Screenshots Callbacks
const (
	ScreenshotReadyID     SteamCallbackID = 2301
	ScreenshotRequestedID SteamCallbackID = 2302
)

type ScreenshotReady struct {
	Local  ScreenshotHandle
	Result EResult
}
type ScreenshotRequested struct {
}

func (cb ScreenshotReady) CallbackID() SteamCallbackID {
	return ScreenshotReadyID
}

func (cb ScreenshotReady) Name() string {
	return "ScreenshotReady"
}

func (cb ScreenshotReady) String() string {
	return callbackString(cb)
}

func (cb ScreenshotRequested) CallbackID() SteamCallbackID {
	return ScreenshotRequestedID
}

func (cb ScreenshotRequested) Name() string {
	return "ScreenshotRequested"
}

func (cb ScreenshotRequested) String() string {
	return callbackString(cb)
}

// Steam UGC Callbacks

const (
	SteamUGCQueryCompletedID          SteamCallbackID = 3401
	SteamUGCRequestUGCDetailsResultID SteamCallbackID = 3402
	CreateItemResultID                SteamCallbackID = 3403
	SubmitItemUpdateResultID          SteamCallbackID = 3404
	ItemInstalledID                   SteamCallbackID = 3405
	DownloadItemResultID              SteamCallbackID = 3406
	UserFavoriteItemsListChangedID    SteamCallbackID = 3407
	SetUserItemVoteResultID           SteamCallbackID = 3408
	GetUserItemVoteResultID           SteamCallbackID = 3409
	StartPlaytimeTrackingResultID     SteamCallbackID = 3410
	StopPlaytimeTrackingResultID      SteamCallbackID = 3411
	AddUGCDependencyResultID          SteamCallbackID = 3412
	RemoveUGCDependencyResultID       SteamCallbackID = 3413
	AddAppDependencyResultID          SteamCallbackID = 3414
	RemoveAppDependencyResultID       SteamCallbackID = 3415
	GetAppDependenciesResultID        SteamCallbackID = 3416
	DeleteItemResultID                SteamCallbackID = 3417
	UserSubscribedItemsListChangedID  SteamCallbackID = 3418
	WorkshopEULAStatusID              SteamCallbackID = 3420
)

type SteamUGCQueryCompleted struct {
	Handle               UGCQueryHandle
	Result               EResult
	NumResultsReturned   uint32
	TotalMatchingResults uint32
	CachedData           bool
	NextCursor           [256]byte
}
type SteamUGCRequestUGCDetailsResult struct {
	Details    SteamUGCDetails
	CachedData bool
}
type CreateItemResult struct {
	Result                                  EResult
	PublishedFileId                         PublishedFileId
	UserNeedsToAcceptWorkshopLegalAgreement bool
}
type SubmitItemUpdateResult struct {
	Result                                  EResult
	UserNeedsToAcceptWorkshopLegalAgreement bool
	PublishedFileId                         PublishedFileId
}
type ItemInstalled struct {
	AppID           AppId
	PublishedFileId PublishedFileId
	LegacyContent   UGCHandle
	ManifestID      uint64
}
type DownloadItemResult struct {
	AppID           AppId
	PublishedFileId PublishedFileId
	Result          EResult
}
type UserFavoriteItemsListChanged struct {
	PublishedFileId PublishedFileId
	Result          EResult
	WasAddRequest   bool
	_               [7]byte
}
type SetUserItemVoteResult struct {
	PublishedFileId PublishedFileId
	Result          EResult
	VoteUp          bool
}
type GetUserItemVoteResult struct {
	PublishedFileId PublishedFileId
	Result          EResult
	VotedUp         bool
	VotedDown       bool
	VoteSkipped     bool
}
type StartPlaytimeTrackingResult struct {
	Result EResult
}
type StopPlaytimeTrackingResult struct {
	Result EResult
}
type AddUGCDependencyResult struct {
	Result               EResult
	PublishedFileId      PublishedFileId
	ChildPublishedFileId PublishedFileId
}
type RemoveUGCDependencyResult struct {
	Result               EResult
	PublishedFileId      PublishedFileId
	ChildPublishedFileId PublishedFileId
}
type AddAppDependencyResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
	AppID           AppId
}
type RemoveAppDependencyResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
	AppID           AppId
}
type GetAppDependenciesResult struct {
	Result                  EResult
	PublishedFileId         PublishedFileId
	AppIDs                  [32]AppId
	NumAppDependencies      uint32
	TotalNumAppDependencies uint32
}
type DeleteItemResult struct {
	Result          EResult
	PublishedFileId PublishedFileId
}
type UserSubscribedItemsListChanged struct {
	AppID AppId
}
type WorkshopEULAStatus struct {
	Result      EResult
	AppID       AppId
	Version     uint32
	Action      RTime32
	Accepted    bool
	NeedsAction bool
}

func (cb SteamUGCQueryCompleted) CallbackID() SteamCallbackID {
	return SteamUGCQueryCompletedID
}

func (cb SteamUGCQueryCompleted) Name() string {
	return "SteamUGCQueryCompleted"
}

func (cb SteamUGCQueryCompleted) String() string {
	return callbackString(cb)
}

func (cb SteamUGCRequestUGCDetailsResult) CallbackID() SteamCallbackID {
	return SteamUGCRequestUGCDetailsResultID
}

func (cb SteamUGCRequestUGCDetailsResult) Name() string {
	return "SteamUGCRequestUGCDetailsResult"
}

func (cb SteamUGCRequestUGCDetailsResult) String() string {
	return callbackString(cb)
}

func (cb CreateItemResult) CallbackID() SteamCallbackID {
	return CreateItemResultID
}

func (cb CreateItemResult) Name() string {
	return "CreateItemResult"
}

func (cb CreateItemResult) String() string {
	return callbackString(cb)
}

func (cb SubmitItemUpdateResult) CallbackID() SteamCallbackID {
	return SubmitItemUpdateResultID
}

func (cb SubmitItemUpdateResult) Name() string {
	return "SubmitItemUpdateResult"
}

func (cb SubmitItemUpdateResult) String() string {
	return callbackString(cb)
}

func (cb ItemInstalled) CallbackID() SteamCallbackID {
	return ItemInstalledID
}

func (cb ItemInstalled) Name() string {
	return "ItemInstalled"
}

func (cb ItemInstalled) String() string {
	return callbackString(cb)
}

func (cb DownloadItemResult) CallbackID() SteamCallbackID {
	return DownloadItemResultID
}

func (cb DownloadItemResult) Name() string {
	return "DownloadItemResult"
}

func (cb DownloadItemResult) String() string {
	return callbackString(cb)
}

func (cb UserFavoriteItemsListChanged) CallbackID() SteamCallbackID {
	return UserFavoriteItemsListChangedID
}

func (cb UserFavoriteItemsListChanged) Name() string {
	return "UserFavoriteItemsListChanged"
}

func (cb UserFavoriteItemsListChanged) String() string {
	return callbackString(cb)
}

func (cb SetUserItemVoteResult) CallbackID() SteamCallbackID {
	return SetUserItemVoteResultID
}

func (cb SetUserItemVoteResult) Name() string {
	return "SetUserItemVoteResult"
}

func (cb SetUserItemVoteResult) String() string {
	return callbackString(cb)
}

func (cb GetUserItemVoteResult) CallbackID() SteamCallbackID {
	return GetUserItemVoteResultID
}

func (cb GetUserItemVoteResult) Name() string {
	return "GetUserItemVoteResult"
}

func (cb GetUserItemVoteResult) String() string {
	return callbackString(cb)
}

func (cb StartPlaytimeTrackingResult) CallbackID() SteamCallbackID {
	return StartPlaytimeTrackingResultID
}

func (cb StartPlaytimeTrackingResult) Name() string {
	return "StartPlaytimeTrackingResult"
}

func (cb StartPlaytimeTrackingResult) String() string {
	return callbackString(cb)
}

func (cb StopPlaytimeTrackingResult) CallbackID() SteamCallbackID {
	return StopPlaytimeTrackingResultID
}

func (cb StopPlaytimeTrackingResult) Name() string {
	return "StopPlaytimeTrackingResult"
}

func (cb StopPlaytimeTrackingResult) String() string {
	return callbackString(cb)
}

func (cb AddUGCDependencyResult) CallbackID() SteamCallbackID {
	return AddUGCDependencyResultID
}

func (cb AddUGCDependencyResult) Name() string {
	return "AddUGCDependencyResult"
}

func (cb AddUGCDependencyResult) String() string {
	return callbackString(cb)
}

func (cb RemoveUGCDependencyResult) CallbackID() SteamCallbackID {
	return RemoveUGCDependencyResultID
}

func (cb RemoveUGCDependencyResult) Name() string {
	return "RemoveUGCDependencyResult"
}

func (cb RemoveUGCDependencyResult) String() string {
	return callbackString(cb)
}

func (cb AddAppDependencyResult) CallbackID() SteamCallbackID {
	return AddAppDependencyResultID
}

func (cb AddAppDependencyResult) Name() string {
	return "AddAppDependencyResult"
}

func (cb AddAppDependencyResult) String() string {
	return callbackString(cb)
}

func (cb RemoveAppDependencyResult) CallbackID() SteamCallbackID {
	return RemoveAppDependencyResultID
}

func (cb RemoveAppDependencyResult) Name() string {
	return "RemoveAppDependencyResult"
}

func (cb RemoveAppDependencyResult) String() string {
	return callbackString(cb)
}

func (cb GetAppDependenciesResult) CallbackID() SteamCallbackID {
	return GetAppDependenciesResultID
}

func (cb GetAppDependenciesResult) Name() string {
	return "GetAppDependenciesResult"
}

func (cb GetAppDependenciesResult) String() string {
	return callbackString(cb)
}

func (cb DeleteItemResult) CallbackID() SteamCallbackID {
	return DeleteItemResultID
}

func (cb DeleteItemResult) Name() string {
	return "DeleteItemResult"
}

func (cb DeleteItemResult) String() string {
	return callbackString(cb)
}

func (cb UserSubscribedItemsListChanged) CallbackID() SteamCallbackID {
	return UserSubscribedItemsListChangedID
}

func (cb UserSubscribedItemsListChanged) Name() string {
	return "UserSubscribedItemsListChanged"
}

func (cb UserSubscribedItemsListChanged) String() string {
	return callbackString(cb)
}

func (cb WorkshopEULAStatus) CallbackID() SteamCallbackID {
	return WorkshopEULAStatusID
}

func (cb WorkshopEULAStatus) Name() string {
	return "WorkshopEULAStatus"
}

func (cb WorkshopEULAStatus) String() string {
	return callbackString(cb)
}

// Steam User Callbacks

const (
	SteamServersConnectedID         SteamCallbackID = 101
	SteamServerConnectFailureID     SteamCallbackID = 102
	SteamServersDisconnectedID      SteamCallbackID = 103
	ClientGameServerDenyID          SteamCallbackID = 113
	IPCFailureID                    SteamCallbackID = 117
	LicensesUpdatedID               SteamCallbackID = 125
	ValidateAuthTicketResponseID    SteamCallbackID = 143
	MicroTxnAuthorizationResponseID SteamCallbackID = 152
	EncryptedAppTicketResponseID    SteamCallbackID = 154
	GetAuthSessionTicketResponseID  SteamCallbackID = 163
	GameWebCallbackID               SteamCallbackID = 164
	StoreAuthURLResponseID          SteamCallbackID = 165
	MarketEligibilityResponseID     SteamCallbackID = 166
	DurationControlID               SteamCallbackID = 167
	GetTicketForWebApiResponseID    SteamCallbackID = 168
)

type SteamServersConnected struct {
}
type SteamServerConnectFailure struct {
	Result        EResult
	StillRetrying bool
}
type SteamServersDisconnected struct {
	Result EResult
}
type ClientGameServerDeny struct {
	AppID          uint32
	GameServerIP   uint32
	GameServerPort uint16
	Secure         uint16
	Reason         uint32
}
type IPCFailure struct {
	FailureType uint8
}
type EFailureType int32

const (
	EFailureFlushedCallbackQueue EFailureType = 0
	EFailurePipeFail             EFailureType = 1
)

type LicensesUpdated struct {
}
type ValidateAuthTicketResponse struct {
	SteamID             CSteamID
	AuthSessionResponse EAuthSessionResponse
	OwnerSteamID        CSteamID
}
type MicroTxnAuthorizationResponse struct {
	AppID      uint32
	OrderID    uint64
	Authorized uint8
}
type EncryptedAppTicketResponse struct {
	Result EResult
}
type GetAuthSessionTicketResponse struct {
	AuthTicket HAuthTicket
	Result     EResult
}
type GameWebCallback struct {
	URL [256]byte
}
type StoreAuthURLResponse struct {
	URL [512]byte
}
type MarketEligibilityResponse struct {
	Allowed                bool
	NotAllowedReason       EMarketNotAllowedReasonFlags
	AllowedAtTime          RTime32
	SteamGuardRequiredDays int32
	NewDeviceCooldown      int32
}
type DurationControl struct {
	Result       EResult
	AppId        AppId
	Applicable   bool
	Last5h       int32
	Progress     EDurationControlProgress
	Notification EDurationControlNotification
	Today        int32
	Remaining    int32
}
type GetTicketForWebApiResponse struct {
	AuthTicket HAuthTicket
	Result     EResult
	TicketSize int32
	TicketData [2560]uint8
}

func (cb SteamServersConnected) CallbackID() SteamCallbackID {
	return SteamServersConnectedID
}

func (cb SteamServersConnected) Name() string {
	return "SteamServersConnected"
}

func (cb SteamServersConnected) String() string {
	return callbackString(cb)
}

func (cb SteamServerConnectFailure) CallbackID() SteamCallbackID {
	return SteamServerConnectFailureID
}

func (cb SteamServerConnectFailure) Name() string {
	return "SteamServerConnectFailure"
}

func (cb SteamServerConnectFailure) String() string {
	return callbackString(cb)
}

func (cb SteamServersDisconnected) CallbackID() SteamCallbackID {
	return SteamServersDisconnectedID
}

func (cb SteamServersDisconnected) Name() string {
	return "SteamServersDisconnected"
}

func (cb SteamServersDisconnected) String() string {
	return callbackString(cb)
}

func (cb ClientGameServerDeny) CallbackID() SteamCallbackID {
	return ClientGameServerDenyID
}

func (cb ClientGameServerDeny) Name() string {
	return "ClientGameServerDeny"
}

func (cb ClientGameServerDeny) String() string {
	return callbackString(cb)
}

func (cb IPCFailure) CallbackID() SteamCallbackID {
	return IPCFailureID
}

func (cb IPCFailure) Name() string {
	return "IPCFailure"
}

func (cb IPCFailure) String() string {
	return callbackString(cb)
}

func (cb LicensesUpdated) CallbackID() SteamCallbackID {
	return LicensesUpdatedID
}

func (cb LicensesUpdated) Name() string {
	return "LicensesUpdated"
}

func (cb LicensesUpdated) String() string {
	return callbackString(cb)
}

func (cb ValidateAuthTicketResponse) CallbackID() SteamCallbackID {
	return ValidateAuthTicketResponseID
}

func (cb ValidateAuthTicketResponse) Name() string {
	return "ValidateAuthTicketResponse"
}

func (cb ValidateAuthTicketResponse) String() string {
	return callbackString(cb)
}

func (cb MicroTxnAuthorizationResponse) CallbackID() SteamCallbackID {
	return MicroTxnAuthorizationResponseID
}

func (cb MicroTxnAuthorizationResponse) Name() string {
	return "MicroTxnAuthorizationResponse"
}

func (cb MicroTxnAuthorizationResponse) String() string {
	return callbackString(cb)
}

func (cb EncryptedAppTicketResponse) CallbackID() SteamCallbackID {
	return EncryptedAppTicketResponseID
}

func (cb EncryptedAppTicketResponse) Name() string {
	return "EncryptedAppTicketResponse"
}

func (cb EncryptedAppTicketResponse) String() string {
	return callbackString(cb)
}

func (cb GetAuthSessionTicketResponse) CallbackID() SteamCallbackID {
	return GetAuthSessionTicketResponseID
}

func (cb GetAuthSessionTicketResponse) Name() string {
	return "GetAuthSessionTicketResponse"
}

func (cb GetAuthSessionTicketResponse) String() string {
	return callbackString(cb)
}

func (cb GameWebCallback) CallbackID() SteamCallbackID {
	return GameWebCallbackID
}

func (cb GameWebCallback) Name() string {
	return "GameWebCallback"
}

func (cb GameWebCallback) String() string {
	return callbackString(cb)
}

func (cb StoreAuthURLResponse) CallbackID() SteamCallbackID {
	return StoreAuthURLResponseID
}

func (cb StoreAuthURLResponse) Name() string {
	return "StoreAuthURLResponse"
}

func (cb StoreAuthURLResponse) String() string {
	return callbackString(cb)
}

func (cb MarketEligibilityResponse) CallbackID() SteamCallbackID {
	return MarketEligibilityResponseID
}

func (cb MarketEligibilityResponse) Name() string {
	return "MarketEligibilityResponse"
}

func (cb MarketEligibilityResponse) String() string {
	return callbackString(cb)
}

func (cb DurationControl) CallbackID() SteamCallbackID {
	return DurationControlID
}

func (cb DurationControl) Name() string {
	return "DurationControl"
}

func (cb DurationControl) String() string {
	return callbackString(cb)
}

func (cb GetTicketForWebApiResponse) CallbackID() SteamCallbackID {
	return GetTicketForWebApiResponseID
}

func (cb GetTicketForWebApiResponse) Name() string {
	return "GetTicketForWebApiResponse"
}

func (cb GetTicketForWebApiResponse) String() string {
	return callbackString(cb)
}

// Steam User Stats Callbacks

const (
	UserStatsReceivedID                 SteamCallbackID = 1101
	UserStatsStoredID                   SteamCallbackID = 1102
	UserAchievementStoredID             SteamCallbackID = 1103
	LeaderboardFindResultID             SteamCallbackID = 1104
	LeaderboardScoresDownloadedID       SteamCallbackID = 1105
	LeaderboardScoreUploadedID          SteamCallbackID = 1106
	NumberOfCurrentPlayersID            SteamCallbackID = 1107
	UserStatsUnloadedID                 SteamCallbackID = 1108
	UserAchievementIconFetchedID        SteamCallbackID = 1109
	GlobalAchievementPercentagesReadyID SteamCallbackID = 1110
	LeaderboardUGCSetID                 SteamCallbackID = 1111
	GlobalStatsReceivedID               SteamCallbackID = 1112
)

type UserStatsReceived struct {
	GameID      uint64
	Result      EResult
	SteamIDUser CSteamID
}
type UserStatsStored struct {
	GameID uint64
	Result EResult
}
type UserAchievementStored struct {
	GameID           uint64
	GroupAchievement bool
	AchievementName  [128]byte
	CurProgress      uint32
	MaxProgress      uint32
}
type LeaderboardFindResult struct {
	SteamLeaderboard SteamLeaderboard
	LeaderboardFound uint8
}
type LeaderboardScoresDownloaded struct {
	SteamLeaderboard        SteamLeaderboard
	SteamLeaderboardEntries SteamLeaderboardEntries
	EntryCount              int32
}
type LeaderboardScoreUploaded struct {
	Success            uint8
	SteamLeaderboard   SteamLeaderboard
	Score              int32
	ScoreChanged       uint8
	GlobalRankNew      int32
	GlobalRankPrevious int32
}
type NumberOfCurrentPlayers struct {
	Success uint8
	Players int32
}
type UserStatsUnloaded struct {
	SteamIDUser CSteamID
}
type UserAchievementIconFetched struct {
	GameID          CGameID
	AchievementName [128]byte
	Achieved        bool
	IconHandle      int32
}
type GlobalAchievementPercentagesReady struct {
	GameID uint64
	Result EResult
}
type LeaderboardUGCSet struct {
	Result           EResult
	SteamLeaderboard SteamLeaderboard
}
type GlobalStatsReceived struct {
	GameID uint64
	Result EResult
}

func (cb UserStatsReceived) CallbackID() SteamCallbackID {
	return UserStatsReceivedID
}

func (cb UserStatsReceived) Name() string {
	return "UserStatsReceived"
}

func (cb UserStatsReceived) String() string {
	return callbackString(cb)
}

func (cb UserStatsStored) CallbackID() SteamCallbackID {
	return UserStatsStoredID
}

func (cb UserStatsStored) Name() string {
	return "UserStatsStored"
}

func (cb UserStatsStored) String() string {
	return callbackString(cb)
}

func (cb UserAchievementStored) CallbackID() SteamCallbackID {
	return UserAchievementStoredID
}

func (cb UserAchievementStored) Name() string {
	return "UserAchievementStored"
}

func (cb UserAchievementStored) String() string {
	return callbackString(cb)
}

func (cb LeaderboardFindResult) CallbackID() SteamCallbackID {
	return LeaderboardFindResultID
}

func (cb LeaderboardFindResult) Name() string {
	return "LeaderboardFindResult"
}

func (cb LeaderboardFindResult) String() string {
	return callbackString(cb)
}

func (cb LeaderboardScoresDownloaded) CallbackID() SteamCallbackID {
	return LeaderboardScoresDownloadedID
}

func (cb LeaderboardScoresDownloaded) Name() string {
	return "LeaderboardScoresDownloaded"
}

func (cb LeaderboardScoresDownloaded) String() string {
	return callbackString(cb)
}

func (cb LeaderboardScoreUploaded) CallbackID() SteamCallbackID {
	return LeaderboardScoreUploadedID
}

func (cb LeaderboardScoreUploaded) Name() string {
	return "LeaderboardScoreUploaded"
}

func (cb LeaderboardScoreUploaded) String() string {
	return callbackString(cb)
}

func (cb NumberOfCurrentPlayers) CallbackID() SteamCallbackID {
	return NumberOfCurrentPlayersID
}

func (cb NumberOfCurrentPlayers) Name() string {
	return "NumberOfCurrentPlayers"
}

func (cb NumberOfCurrentPlayers) String() string {
	return callbackString(cb)
}

func (cb UserStatsUnloaded) CallbackID() SteamCallbackID {
	return UserStatsUnloadedID
}

func (cb UserStatsUnloaded) Name() string {
	return "UserStatsUnloaded"
}

func (cb UserStatsUnloaded) String() string {
	return callbackString(cb)
}

func (cb UserAchievementIconFetched) CallbackID() SteamCallbackID {
	return UserAchievementIconFetchedID
}

func (cb UserAchievementIconFetched) Name() string {
	return "UserAchievementIconFetched"
}

func (cb UserAchievementIconFetched) String() string {
	return callbackString(cb)
}

func (cb GlobalAchievementPercentagesReady) CallbackID() SteamCallbackID {
	return GlobalAchievementPercentagesReadyID
}

func (cb GlobalAchievementPercentagesReady) Name() string {
	return "GlobalAchievementPercentagesReady"
}

func (cb GlobalAchievementPercentagesReady) String() string {
	return callbackString(cb)
}

func (cb LeaderboardUGCSet) CallbackID() SteamCallbackID {
	return LeaderboardUGCSetID
}

func (cb LeaderboardUGCSet) Name() string {
	return "LeaderboardUGCSet"
}

func (cb LeaderboardUGCSet) String() string {
	return callbackString(cb)
}

func (cb GlobalStatsReceived) CallbackID() SteamCallbackID {
	return GlobalStatsReceivedID
}

func (cb GlobalStatsReceived) Name() string {
	return "GlobalStatsReceived"
}

func (cb GlobalStatsReceived) String() string {
	return callbackString(cb)
}

// Steam Utils Callbacks

const (
	IPCountryID                         SteamCallbackID = 701
	LowBatteryPowerID                   SteamCallbackID = 702
	SteamShutdownID                     SteamCallbackID = 704
	CheckFileSignatureID                SteamCallbackID = 705
	GamepadTextInputDismissedID         SteamCallbackID = 714
	AppResumingFromSuspendID            SteamCallbackID = 736
	FloatingGamepadTextInputDismissedID SteamCallbackID = 738
	FilterTextDictionaryChangedID       SteamCallbackID = 739
)

type IPCountry struct {
}
type LowBatteryPower struct {
	MinutesBatteryLeft uint8
}

type SteamShutdown struct {
}
type ECheckFileSignature int32

const (
	ECheckFileSignatureInvalidSignature             ECheckFileSignature = 0
	ECheckFileSignatureValidSignature               ECheckFileSignature = 1
	ECheckFileSignatureFileNotFound                 ECheckFileSignature = 2
	ECheckFileSignatureNoSignaturesFoundForThisApp  ECheckFileSignature = 3
	ECheckFileSignatureNoSignaturesFoundForThisFile ECheckFileSignature = 4
)

type CheckFileSignatureResult struct {
	CheckFileSignature ECheckFileSignature
}
type GamepadTextInputDismissed struct {
	Submitted     bool
	SubmittedText uint32
	AppID         AppId
}
type AppResumingFromSuspend struct {
}
type FloatingGamepadTextInputDismissed struct {
}
type FilterTextDictionaryChanged struct {
	Language int32
}

func (cb IPCountry) CallbackID() SteamCallbackID {
	return IPCountryID
}

func (cb IPCountry) Name() string {
	return "IPCountry"
}

func (cb IPCountry) String() string {
	return callbackString(cb)
}

func (cb LowBatteryPower) CallbackID() SteamCallbackID {
	return LowBatteryPowerID
}

func (cb LowBatteryPower) Name() string {
	return "LowBatteryPower"
}

func (cb LowBatteryPower) String() string {
	return callbackString(cb)
}

func (cb SteamShutdown) CallbackID() SteamCallbackID {
	return SteamShutdownID
}

func (cb SteamShutdown) Name() string {
	return "SteamShutdown"
}

func (cb SteamShutdown) String() string {
	return callbackString(cb)
}

func (cb CheckFileSignatureResult) CallbackID() SteamCallbackID {
	return CheckFileSignatureID
}

func (cb CheckFileSignatureResult) Name() string {
	return "CheckFileSignature"
}

func (cb CheckFileSignatureResult) String() string {
	return callbackString(cb)
}

func (cb GamepadTextInputDismissed) CallbackID() SteamCallbackID {
	return GamepadTextInputDismissedID
}

func (cb GamepadTextInputDismissed) Name() string {
	return "GamepadTextInputDismissed"
}

func (cb GamepadTextInputDismissed) String() string {
	return callbackString(cb)
}

func (cb AppResumingFromSuspend) CallbackID() SteamCallbackID {
	return AppResumingFromSuspendID
}

func (cb AppResumingFromSuspend) Name() string {
	return "AppResumingFromSuspend"
}

func (cb AppResumingFromSuspend) String() string {
	return callbackString(cb)
}

func (cb FloatingGamepadTextInputDismissed) CallbackID() SteamCallbackID {
	return FloatingGamepadTextInputDismissedID
}

func (cb FloatingGamepadTextInputDismissed) Name() string {
	return "FloatingGamepadTextInputDismissed"
}

func (cb FloatingGamepadTextInputDismissed) String() string {
	return callbackString(cb)
}

func (cb FilterTextDictionaryChanged) CallbackID() SteamCallbackID {
	return FilterTextDictionaryChangedID
}

func (cb FilterTextDictionaryChanged) Name() string {
	return "FilterTextDictionaryChanged"
}

func (cb FilterTextDictionaryChanged) String() string {
	return callbackString(cb)
}
