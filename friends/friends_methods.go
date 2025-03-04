package friends

import (
	. "github.com/assemblaj/purego-steamworks/matchmaking"

	. "github.com/assemblaj/purego-steamworks"
)

const (
	MaxFriendsGroupName                    int32          = 64
	FriendsGroupLimit                      int32          = 100
	FriendsGroupID_Invalid                 FriendsGroupID = -1
	EnumerateFollowersMax                  int32          = 50
	FriendGameInfoQueryPort_NotInitialized uint16         = 0xFFFF
	FriendGameInfoQueryPort_Error          uint16         = 0xFFFE
	ChatMetadataMax                        uint32         = 8192
)

type EUserRestriction int32

const (
	UserRestrictionNone        EUserRestriction = 0
	UserRestrictionUnknown     EUserRestriction = 1
	UserRestrictionAnyChat     EUserRestriction = 2
	UserRestrictionVoiceChat   EUserRestriction = 4
	UserRestrictionGroupChat   EUserRestriction = 8
	UserRestrictionRating      EUserRestriction = 16
	UserRestrictionGameInvites EUserRestriction = 32
	UserRestrictionTrading     EUserRestriction = 64
)

type ECommunityProfileItemType int32

const (
	ECommunityProfileItemTypeAnimatedAvatar        ECommunityProfileItemType = 0
	ECommunityProfileItemTypeAvatarFrame           ECommunityProfileItemType = 1
	ECommunityProfileItemTypeProfileModifier       ECommunityProfileItemType = 2
	ECommunityProfileItemTypeProfileBackground     ECommunityProfileItemType = 3
	ECommunityProfileItemTypeMiniProfileBackground ECommunityProfileItemType = 4
)

type ECommunityProfileItemProperty int32

const (
	ECommunityProfileItemPropertyImageSmall     ECommunityProfileItemProperty = 0
	ECommunityProfileItemPropertyImageLarge     ECommunityProfileItemProperty = 1
	ECommunityProfileItemPropertyInternalName   ECommunityProfileItemProperty = 2
	ECommunityProfileItemPropertyTitle          ECommunityProfileItemProperty = 3
	ECommunityProfileItemPropertyDescription    ECommunityProfileItemProperty = 4
	ECommunityProfileItemPropertyAppID          ECommunityProfileItemProperty = 5
	ECommunityProfileItemPropertyTypeID         ECommunityProfileItemProperty = 6
	ECommunityProfileItemPropertyClass          ECommunityProfileItemProperty = 7
	ECommunityProfileItemPropertyMovieWebM      ECommunityProfileItemProperty = 8
	ECommunityProfileItemPropertyMovieMP4       ECommunityProfileItemProperty = 9
	ECommunityProfileItemPropertyMovieWebMSmall ECommunityProfileItemProperty = 10
	ECommunityProfileItemPropertyMovieMP4Small  ECommunityProfileItemProperty = 11
)

type EPersonaChange int32

const (
	EPersonaChangeName                EPersonaChange = 1
	EPersonaChangeStatus              EPersonaChange = 2
	EPersonaChangeComeOnline          EPersonaChange = 4
	EPersonaChangeGoneOffline         EPersonaChange = 8
	EPersonaChangeGamePlayed          EPersonaChange = 16
	EPersonaChangeGameServer          EPersonaChange = 32
	EPersonaChangeAvatar              EPersonaChange = 64
	EPersonaChangeJoinedSource        EPersonaChange = 128
	EPersonaChangeLeftSource          EPersonaChange = 256
	EPersonaChangeRelationshipChanged EPersonaChange = 512
	EPersonaChangeNameFirstSet        EPersonaChange = 1024
	EPersonaChangeBroadcast           EPersonaChange = 2048
	EPersonaChangeNickname            EPersonaChange = 4096
	EPersonaChangeSteamLevel          EPersonaChange = 8192
	EPersonaChangeRichPresence        EPersonaChange = 16384
)

type FriendsGroupID int16

type FriendGameInfo struct {
	GameID       CGameID
	GameIP       uint32
	GamePort     uint16
	QueryPort    uint16
	lobbySteamID CSteamID
}

type EActivateGameOverlayToWebPageMode int32

const (
	EActivateGameOverlayToWebPageModeDefault EActivateGameOverlayToWebPageMode = 0
	EActivateGameOverlayToWebPageModeModal   EActivateGameOverlayToWebPageMode = 1
)

type EOverlayToStoreFlag int32

const (
	EOverlayToStoreFlagNone             EOverlayToStoreFlag = 0
	EOverlayToStoreFlagAddToCart        EOverlayToStoreFlag = 1
	EOverlayToStoreFlagAddToCartAndShow EOverlayToStoreFlag = 2
)

type EFriendRelationship int32

const (
	EFriendRelationshipNone                EFriendRelationship = 0
	EFriendRelationshipBlocked             EFriendRelationship = 1
	EFriendRelationshipRequestRecipient    EFriendRelationship = 2
	EFriendRelationshipFriend              EFriendRelationship = 3
	EFriendRelationshipRequestInitiator    EFriendRelationship = 4
	EFriendRelationshipIgnored             EFriendRelationship = 5
	EFriendRelationshipIgnoredFriend       EFriendRelationship = 6
	EFriendRelationshipSuggestedDEPRECATED EFriendRelationship = 7
	EFriendRelationshipMax                 EFriendRelationship = 8
)

type EPersonaState int32

const (
	EPersonaStateOffline        EPersonaState = 0
	EPersonaStateOnline         EPersonaState = 1
	EPersonaStateBusy           EPersonaState = 2
	EPersonaStateAway           EPersonaState = 3
	EPersonaStateSnooze         EPersonaState = 4
	EPersonaStateLookingToTrade EPersonaState = 5
	EPersonaStateLookingToPlay  EPersonaState = 6
	EPersonaStateInvisible      EPersonaState = 7
	EPersonaStateMax            EPersonaState = 8
)

type EFriendFlags int32

const (
	EFriendFlagNone                 EFriendFlags = 0
	EFriendFlagBlocked              EFriendFlags = 1
	EFriendFlagFriendshipRequested  EFriendFlags = 2
	EFriendFlagImmediate            EFriendFlags = 4
	EFriendFlagClanMember           EFriendFlags = 8
	EFriendFlagOnGameServer         EFriendFlags = 16
	EFriendFlagRequestingFriendship EFriendFlags = 128
	EFriendFlagRequestingInfo       EFriendFlags = 256
	EFriendFlagIgnored              EFriendFlags = 512
	EFriendFlagIgnoredFriend        EFriendFlags = 1024
	EFriendFlagChatMember           EFriendFlags = 4096
	EFriendFlagAll                  EFriendFlags = 65535
)

const (
	flatAPI_SteamFriends                                                    = "SteamAPI_SteamFriends_v017"
	flatAPI_ISteamFriends_GetPersonaName                                    = "SteamAPI_ISteamFriends_GetPersonaName"
	flatAPI_ISteamFriends_SetPersonaName                                    = "SteamAPI_ISteamFriends_SetPersonaName"
	flatAPI_ISteamFriends_GetPersonaState                                   = "SteamAPI_ISteamFriends_GetPersonaState"
	flatAPI_ISteamFriends_GetFriendCount                                    = "SteamAPI_ISteamFriends_GetFriendCount"
	flatAPI_ISteamFriends_GetFriendByIndex                                  = "SteamAPI_ISteamFriends_GetFriendByIndex"
	flatAPI_ISteamFriends_GetFriendRelationship                             = "SteamAPI_ISteamFriends_GetFriendRelationship"
	flatAPI_ISteamFriends_GetFriendPersonaState                             = "SteamAPI_ISteamFriends_GetFriendPersonaState"
	flatAPI_ISteamFriends_GetFriendPersonaName                              = "SteamAPI_ISteamFriends_GetFriendPersonaName"
	flatAPI_ISteamFriends_GetFriendGamePlayed                               = "SteamAPI_ISteamFriends_GetFriendGamePlayed"
	flatAPI_ISteamFriends_GetFriendPersonaNameHistory                       = "SteamAPI_ISteamFriends_GetFriendPersonaNameHistory"
	flatAPI_ISteamFriends_GetFriendSteamLevel                               = "SteamAPI_ISteamFriends_GetFriendSteamLevel"
	flatAPI_ISteamFriends_GetFriendsGroupCount                              = "SteamAPI_ISteamFriends_GetFriendsGroupCount"
	flatAPI_ISteamFriends_GetFriendsGroupIDByIndex                          = "SteamAPI_ISteamFriends_GetFriendsGroupIDByIndex"
	flatAPI_ISteamFriends_GetFriendsGroupName                               = "SteamAPI_ISteamFriends_GetFriendsGroupName"
	flatAPI_ISteamFriends_GetFriendsGroupMembersCount                       = "SteamAPI_ISteamFriends_GetFriendsGroupMembersCount"
	flatAPI_ISteamFriends_GetFriendsGroupMembersList                        = "SteamAPI_ISteamFriends_GetFriendsGroupMembersList"
	flatAPI_ISteamFriends_HasFriend                                         = "SteamAPI_ISteamFriends_HasFriend"
	flatAPI_ISteamFriends_GetClanCount                                      = "SteamAPI_ISteamFriends_GetClanCount"
	flatAPI_ISteamFriends_GetClanByIndex                                    = "SteamAPI_ISteamFriends_GetClanByIndex"
	flatAPI_ISteamFriends_GetClanName                                       = "SteamAPI_ISteamFriends_GetClanName"
	flatAPI_ISteamFriends_GetClanTag                                        = "SteamAPI_ISteamFriends_GetClanTag"
	flatAPI_ISteamFriends_GetClanActivityCounts                             = "SteamAPI_ISteamFriends_GetClanActivityCounts"
	flatAPI_ISteamFriends_DownloadClanActivityCounts                        = "SteamAPI_ISteamFriends_DownloadClanActivityCounts"
	flatAPI_ISteamFriends_GetFriendCountFromSource                          = "SteamAPI_ISteamFriends_GetFriendCountFromSource"
	flatAPI_ISteamFriends_GetFriendFromSourceByIndex                        = "SteamAPI_ISteamFriends_GetFriendFromSourceByIndex"
	flatAPI_ISteamFriends_IsUserInSource                                    = "SteamAPI_ISteamFriends_IsUserInSource"
	flatAPI_ISteamFriends_SetInGameVoiceSpeaking                            = "SteamAPI_ISteamFriends_SetInGameVoiceSpeaking"
	flatAPI_ISteamFriends_ActivateGameOverlay                               = "SteamAPI_ISteamFriends_ActivateGameOverlay"
	flatAPI_ISteamFriends_ActivateGameOverlayToUser                         = "SteamAPI_ISteamFriends_ActivateGameOverlayToUser"
	flatAPI_ISteamFriends_ActivateGameOverlayToWebPage                      = "SteamAPI_ISteamFriends_ActivateGameOverlayToWebPage"
	flatAPI_ISteamFriends_ActivateGameOverlayToStore                        = "SteamAPI_ISteamFriends_ActivateGameOverlayToStore"
	flatAPI_ISteamFriends_SetPlayedWith                                     = "SteamAPI_ISteamFriends_SetPlayedWith"
	flatAPI_ISteamFriends_ActivateGameOverlayInviteDialog                   = "SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialog"
	flatAPI_ISteamFriends_GetSmallFriendAvatar                              = "SteamAPI_ISteamFriends_GetSmallFriendAvatar"
	flatAPI_ISteamFriends_GetMediumFriendAvatar                             = "SteamAPI_ISteamFriends_GetMediumFriendAvatar"
	flatAPI_ISteamFriends_GetLargeFriendAvatar                              = "SteamAPI_ISteamFriends_GetLargeFriendAvatar"
	flatAPI_ISteamFriends_RequestUserInformation                            = "SteamAPI_ISteamFriends_RequestUserInformation"
	flatAPI_ISteamFriends_RequestClanOfficerList                            = "SteamAPI_ISteamFriends_RequestClanOfficerList"
	flatAPI_ISteamFriends_GetClanOwner                                      = "SteamAPI_ISteamFriends_GetClanOwner"
	flatAPI_ISteamFriends_GetClanOfficerCount                               = "SteamAPI_ISteamFriends_GetClanOfficerCount"
	flatAPI_ISteamFriends_GetClanOfficerByIndex                             = "SteamAPI_ISteamFriends_GetClanOfficerByIndex"
	flatAPI_ISteamFriends_GetUserRestrictions                               = "SteamAPI_ISteamFriends_GetUserRestrictions"
	flatAPI_ISteamFriends_SetRichPresence                                   = "SteamAPI_ISteamFriends_SetRichPresence"
	flatAPI_ISteamFriends_ClearRichPresence                                 = "SteamAPI_ISteamFriends_ClearRichPresence"
	flatAPI_ISteamFriends_GetFriendRichPresence                             = "SteamAPI_ISteamFriends_GetFriendRichPresence"
	flatAPI_ISteamFriends_GetFriendRichPresenceKeyCount                     = "SteamAPI_ISteamFriends_GetFriendRichPresenceKeyCount"
	flatAPI_ISteamFriends_GetFriendRichPresenceKeyByIndex                   = "SteamAPI_ISteamFriends_GetFriendRichPresenceKeyByIndex"
	flatAPI_ISteamFriends_RequestFriendRichPresence                         = "SteamAPI_ISteamFriends_RequestFriendRichPresence"
	flatAPI_ISteamFriends_InviteUserToGame                                  = "SteamAPI_ISteamFriends_InviteUserToGame"
	flatAPI_ISteamFriends_GetCoplayFriendCount                              = "SteamAPI_ISteamFriends_GetCoplayFriendCount"
	flatAPI_ISteamFriends_GetCoplayFriend                                   = "SteamAPI_ISteamFriends_GetCoplayFriend"
	flatAPI_ISteamFriends_GetFriendCoplayTime                               = "SteamAPI_ISteamFriends_GetFriendCoplayTime"
	flatAPI_ISteamFriends_GetFriendCoplayGame                               = "SteamAPI_ISteamFriends_GetFriendCoplayGame"
	flatAPI_ISteamFriends_JoinClanChatRoom                                  = "SteamAPI_ISteamFriends_JoinClanChatRoom"
	flatAPI_ISteamFriends_LeaveClanChatRoom                                 = "SteamAPI_ISteamFriends_LeaveClanChatRoom"
	flatAPI_ISteamFriends_GetClanChatMemberCount                            = "SteamAPI_ISteamFriends_GetClanChatMemberCount"
	flatAPI_ISteamFriends_GetChatMemberByIndex                              = "SteamAPI_ISteamFriends_GetChatMemberByIndex"
	flatAPI_ISteamFriends_SendClanChatMessage                               = "SteamAPI_ISteamFriends_SendClanChatMessage"
	flatAPI_ISteamFriends_GetClanChatMessage                                = "SteamAPI_ISteamFriends_GetClanChatMessage"
	flatAPI_ISteamFriends_IsClanChatAdmin                                   = "SteamAPI_ISteamFriends_IsClanChatAdmin"
	flatAPI_ISteamFriends_IsClanChatWindowOpenInSteam                       = "SteamAPI_ISteamFriends_IsClanChatWindowOpenInSteam"
	flatAPI_ISteamFriends_OpenClanChatWindowInSteam                         = "SteamAPI_ISteamFriends_OpenClanChatWindowInSteam"
	flatAPI_ISteamFriends_CloseClanChatWindowInSteam                        = "SteamAPI_ISteamFriends_CloseClanChatWindowInSteam"
	flatAPI_ISteamFriends_SetListenForFriendsMessages                       = "SteamAPI_ISteamFriends_SetListenForFriendsMessages"
	flatAPI_ISteamFriends_ReplyToFriendMessage                              = "SteamAPI_ISteamFriends_ReplyToFriendMessage"
	flatAPI_ISteamFriends_GetFriendMessage                                  = "SteamAPI_ISteamFriends_GetFriendMessage"
	flatAPI_ISteamFriends_GetFollowerCount                                  = "SteamAPI_ISteamFriends_GetFollowerCount"
	flatAPI_ISteamFriends_IsFollowing                                       = "SteamAPI_ISteamFriends_IsFollowing"
	flatAPI_ISteamFriends_EnumerateFollowingList                            = "SteamAPI_ISteamFriends_EnumerateFollowingList"
	flatAPI_ISteamFriends_IsClanPublic                                      = "SteamAPI_ISteamFriends_IsClanPublic"
	flatAPI_ISteamFriends_IsClanOfficialGameGroup                           = "SteamAPI_ISteamFriends_IsClanOfficialGameGroup"
	flatAPI_ISteamFriends_GetNumChatsWithUnreadPriorityMessages             = "SteamAPI_ISteamFriends_GetNumChatsWithUnreadPriorityMessages"
	flatAPI_ISteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog = "SteamAPI_ISteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog"
	flatAPI_ISteamFriends_RegisterProtocolInOverlayBrowser                  = "SteamAPI_ISteamFriends_RegisterProtocolInOverlayBrowser"
	flatAPI_ISteamFriends_ActivateGameOverlayInviteDialogConnectString      = "SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialogConnectString"
	flatAPI_ISteamFriends_RequestEquippedProfileItems                       = "SteamAPI_ISteamFriends_RequestEquippedProfileItems"
	flatAPI_ISteamFriends_BHasEquippedProfileItem                           = "SteamAPI_ISteamFriends_BHasEquippedProfileItem"
	flatAPI_ISteamFriends_GetProfileItemPropertyString                      = "SteamAPI_ISteamFriends_GetProfileItemPropertyString"
	flatAPI_ISteamFriends_GetProfileItemPropertyUint                        = "SteamAPI_ISteamFriends_GetProfileItemPropertyUint"
)

var (
	steamFriends_init                                               func() uintptr
	iSteamFriends_GetPersonaName                                    func(steamFriends uintptr) string
	iSteamFriends_SetPersonaName                                    func(steamFriends uintptr, pchPersonaName string) SteamAPICall
	iSteamFriends_GetPersonaState                                   func(steamFriends uintptr) EPersonaState
	iSteamFriends_GetFriendCount                                    func(steamFriends uintptr, iFriendFlags int32) int32
	iSteamFriends_GetFriendByIndex                                  func(steamFriends uintptr, iFriend int32, iFriendFlags int32) Uint64SteamID
	iSteamFriends_GetFriendRelationship                             func(steamFriends uintptr, friendSteamID Uint64SteamID) EFriendRelationship
	iSteamFriends_GetFriendPersonaState                             func(steamFriends uintptr, friendSteamID Uint64SteamID) EPersonaState
	iSteamFriends_GetFriendPersonaName                              func(steamFriends uintptr, friendSteamID Uint64SteamID) string
	iSteamFriends_GetFriendGamePlayed                               func(steamFriends uintptr, friendSteamID Uint64SteamID, pFriendGameInfo *FriendGameInfo) bool
	iSteamFriends_GetFriendPersonaNameHistory                       func(steamFriends uintptr, friendSteamID Uint64SteamID, iPersonaName int32) string
	iSteamFriends_GetFriendSteamLevel                               func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_GetFriendsGroupCount                              func(steamFriends uintptr) int32
	iSteamFriends_GetFriendsGroupIDByIndex                          func(steamFriends uintptr, iFG int32) FriendsGroupID
	iSteamFriends_GetFriendsGroupName                               func(steamFriends uintptr, friendsGroupID FriendsGroupID) string
	iSteamFriends_GetFriendsGroupMembersCount                       func(steamFriends uintptr, friendsGroupID FriendsGroupID) int32
	iSteamFriends_GetFriendsGroupMembersList                        func(steamFriends uintptr, friendsGroupID FriendsGroupID, pOutSteamIDMembers *[]CSteamID, nMembersCount int32)
	iSteamFriends_HasFriend                                         func(steamFriends uintptr, friendSteamID Uint64SteamID, iFriendFlags int32) bool
	iSteamFriends_GetClanCount                                      func(steamFriends uintptr) int32
	iSteamFriends_GetClanByIndex                                    func(steamFriends uintptr, iClan int32) Uint64SteamID
	iSteamFriends_GetClanName                                       func(steamFriends uintptr, clanSteamID Uint64SteamID) string
	iSteamFriends_GetClanTag                                        func(steamFriends uintptr, clanSteamID Uint64SteamID) string
	iSteamFriends_GetClanActivityCounts                             func(steamFriends uintptr, clanSteamID Uint64SteamID, pnOnline *int32, pnInGame *int32, pnChatting *int32) bool
	iSteamFriends_DownloadClanActivityCounts                        func(steamFriends uintptr, pclanSteamIDs *[]CSteamID, cClansToRequest int32) SteamAPICall
	iSteamFriends_GetFriendCountFromSource                          func(steamFriends uintptr, sourceSteamID Uint64SteamID) int32
	iSteamFriends_GetFriendFromSourceByIndex                        func(steamFriends uintptr, sourceSteamID Uint64SteamID, iFriend int32) Uint64SteamID
	iSteamFriends_IsUserInSource                                    func(steamFriends uintptr, userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool
	iSteamFriends_SetInGameVoiceSpeaking                            func(steamFriends uintptr, userSteamID Uint64SteamID, bSpeaking bool)
	iSteamFriends_ActivateGameOverlay                               func(steamFriends uintptr, pchDialog string)
	iSteamFriends_ActivateGameOverlayToUser                         func(steamFriends uintptr, pchDialog string, steamID Uint64SteamID)
	iSteamFriends_ActivateGameOverlayToWebPage                      func(steamFriends uintptr, pchURL string, eMode EActivateGameOverlayToWebPageMode)
	iSteamFriends_ActivateGameOverlayToStore                        func(steamFriends uintptr, nAppID AppId, eFlag EOverlayToStoreFlag)
	iSteamFriends_SetPlayedWith                                     func(steamFriends uintptr, userSteamIDPlayedWith Uint64SteamID)
	iSteamFriends_ActivateGameOverlayInviteDialog                   func(steamFriends uintptr, lobbySteamID Uint64SteamID)
	iSteamFriends_GetSmallFriendAvatar                              func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_GetMediumFriendAvatar                             func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_GetLargeFriendAvatar                              func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_RequestUserInformation                            func(steamFriends uintptr, userSteamID Uint64SteamID, bRequireNameOnly bool) bool
	iSteamFriends_RequestClanOfficerList                            func(steamFriends uintptr, clanSteamID Uint64SteamID) SteamAPICall
	iSteamFriends_GetClanOwner                                      func(steamFriends uintptr, clanSteamID Uint64SteamID) Uint64SteamID
	iSteamFriends_GetClanOfficerCount                               func(steamFriends uintptr, clanSteamID Uint64SteamID) int32
	iSteamFriends_GetClanOfficerByIndex                             func(steamFriends uintptr, clanSteamID Uint64SteamID, iOfficer int32) Uint64SteamID
	iSteamFriends_GetUserRestrictions                               func(steamFriends uintptr) uint32
	iSteamFriends_SetRichPresence                                   func(steamFriends uintptr, pchKey string, pchValue string) bool
	iSteamFriends_ClearRichPresence                                 func(steamFriends uintptr)
	iSteamFriends_GetFriendRichPresence                             func(steamFriends uintptr, friendSteamID Uint64SteamID, pchKey string) []byte
	iSteamFriends_GetFriendRichPresenceKeyCount                     func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_GetFriendRichPresenceKeyByIndex                   func(steamFriends uintptr, friendSteamID Uint64SteamID, iKey int32) []byte
	iSteamFriends_RequestFriendRichPresence                         func(steamFriends uintptr, friendSteamID Uint64SteamID)
	iSteamFriends_InviteUserToGame                                  func(steamFriends uintptr, friendSteamID Uint64SteamID, pchConnectString string) bool
	iSteamFriends_GetCoplayFriendCount                              func(steamFriends uintptr) int32
	iSteamFriends_GetCoplayFriend                                   func(steamFriends uintptr, iCoplayFriend int32) Uint64SteamID
	iSteamFriends_GetFriendCoplayTime                               func(steamFriends uintptr, friendSteamID Uint64SteamID) int32
	iSteamFriends_GetFriendCoplayGame                               func(steamFriends uintptr, friendSteamID Uint64SteamID) AppId
	iSteamFriends_JoinClanChatRoom                                  func(steamFriends uintptr, clanSteamID Uint64SteamID) SteamAPICall
	iSteamFriends_LeaveClanChatRoom                                 func(steamFriends uintptr, clanSteamID Uint64SteamID) bool
	iSteamFriends_GetClanChatMemberCount                            func(steamFriends uintptr, clanSteamID Uint64SteamID) int32
	iSteamFriends_GetChatMemberByIndex                              func(steamFriends uintptr, clanSteamID Uint64SteamID, iUser int32) Uint64SteamID
	iSteamFriends_SendClanChatMessage                               func(steamFriends uintptr, clanChatSteamID Uint64SteamID, pchText string) bool
	iSteamFriends_GetClanChatMessage                                func(steamFriends uintptr, clanChatSteamID Uint64SteamID, iMessage int32, prgchText *[]byte, cchTextMax int32, peChatEntryType *EChatEntryType, psteamidChatter *CSteamID) int32
	iSteamFriends_IsClanChatAdmin                                   func(steamFriends uintptr, clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool
	iSteamFriends_IsClanChatWindowOpenInSteam                       func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_OpenClanChatWindowInSteam                         func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_CloseClanChatWindowInSteam                        func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_SetListenForFriendsMessages                       func(steamFriends uintptr, bInterceptEnabled bool) bool
	iSteamFriends_ReplyToFriendMessage                              func(steamFriends uintptr, friendSteamID Uint64SteamID, pchMsgToSend string) bool
	iSteamFriends_GetFriendMessage                                  func(steamFriends uintptr, friendSteamID Uint64SteamID, iMessageID int32, pvData *[]byte, cubData int32, peChatEntryType *EChatEntryType) int32
	iSteamFriends_GetFollowerCount                                  func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall
	iSteamFriends_IsFollowing                                       func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall
	iSteamFriends_EnumerateFollowingList                            func(steamFriends uintptr, unStartIndex uint32) SteamAPICall
	iSteamFriends_IsClanPublic                                      func(steamFriends uintptr, clanSteamID Uint64SteamID) bool
	iSteamFriends_IsClanOfficialGameGroup                           func(steamFriends uintptr, clanSteamID Uint64SteamID) bool
	iSteamFriends_GetNumChatsWithUnreadPriorityMessages             func(steamFriends uintptr) int32
	iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog func(steamFriends uintptr, lobbySteamID Uint64SteamID)
	iSteamFriends_RegisterProtocolInOverlayBrowser                  func(steamFriends uintptr, pchProtocol string) bool
	iSteamFriends_ActivateGameOverlayInviteDialogConnectString      func(steamFriends uintptr, pchConnectString string)
	iSteamFriends_RequestEquippedProfileItems                       func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall
	iSteamFriends_BHasEquippedProfileItem                           func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType) bool
	iSteamFriends_GetProfileItemPropertyString                      func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) string
	iSteamFriends_GetProfileItemPropertyUint                        func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) uint32
)

var steamFriends uintptr

func friends() uintptr {
	if steamFriends == 0 {
		steamFriends = steamFriends_init()
		return steamFriends
	}
	return steamFriends
}

func GetPersonaName() string {
	return iSteamFriends_GetPersonaName(friends())
}

func SetPersonaName(personaName string) CallResult[SetPersonaNameResponse] {
	handle := iSteamFriends_SetPersonaName(friends(), personaName)
	return CallResult[SetPersonaNameResponse]{Handle: handle}
}

func GetPersonaState() EPersonaState {
	return iSteamFriends_GetPersonaState(friends())
}

func GetFriendCount(friendFlags int32) int32 {
	return iSteamFriends_GetFriendCount(friends(), friendFlags)
}

func GetFriendByIndex(friend int32, friendFlags int32) Uint64SteamID {
	return iSteamFriends_GetFriendByIndex(friends(), friend, friendFlags)
}

func GetFriendRelationship(friendSteamID Uint64SteamID) EFriendRelationship {
	return iSteamFriends_GetFriendRelationship(friends(), friendSteamID)
}

func GetFriendPersonaState(friendSteamID Uint64SteamID) EPersonaState {
	return iSteamFriends_GetFriendPersonaState(friends(), friendSteamID)
}

func GetFriendPersonaName(friendSteamID Uint64SteamID) string {
	return iSteamFriends_GetFriendPersonaName(friends(), friendSteamID)
}

func GetFriendGamePlayed(friendSteamID Uint64SteamID) (FriendGameInfo, bool) {
	var pFriendGameInfo FriendGameInfo
	result := iSteamFriends_GetFriendGamePlayed(friends(), friendSteamID, &pFriendGameInfo)
	return pFriendGameInfo, result
}

func GetFriendPersonaNameHistory(friendSteamID Uint64SteamID, personaName int32) string {
	return iSteamFriends_GetFriendPersonaNameHistory(friends(), friendSteamID, personaName)
}

func GetFriendSteamLevel(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendSteamLevel(friends(), friendSteamID)
}

func GetFriendsGroupCount() int32 {
	return iSteamFriends_GetFriendsGroupCount(friends())
}

func GetFriendsGroupIDByIndex(index int32) FriendsGroupID {
	return iSteamFriends_GetFriendsGroupIDByIndex(friends(), index)
}

func GetFriendsGroupName(friendsGroupID FriendsGroupID) string {
	return iSteamFriends_GetFriendsGroupName(friends(), friendsGroupID)
}

func GetFriendsGroupMembersCount(friendsGroupID FriendsGroupID) int32 {
	return iSteamFriends_GetFriendsGroupMembersCount(friends(), friendsGroupID)
}

// Calls GetFriendsGroupMembersCount to get exact number of friends in group
func GetFriendsGroupMembersList(friendsGroupID FriendsGroupID) (members []CSteamID) {
	count := GetFriendsGroupMembersCount(friendsGroupID)
	members = make([]CSteamID, 0, count)
	iSteamFriends_GetFriendsGroupMembersList(friends(), friendsGroupID, &members, count)
	return members
}

func HasFriend(friendSteamID Uint64SteamID, friendFlags int32) bool {
	return iSteamFriends_HasFriend(friends(), friendSteamID, friendFlags)
}

func GetClanCount() int32 {
	return iSteamFriends_GetClanCount(friends())
}

func GetClanByIndex(index int32) Uint64SteamID {
	return iSteamFriends_GetClanByIndex(friends(), index)
}

func GetClanName(clanSteamID Uint64SteamID) string {
	return iSteamFriends_GetClanName(friends(), clanSteamID)
}

func GetClanTag(clanSteamID Uint64SteamID) string {
	return iSteamFriends_GetClanTag(friends(), clanSteamID)
}

func GetClanActivityCounts(clanSteamID Uint64SteamID) (online int32, inGame int32, chatting int32, success bool) {
	result := iSteamFriends_GetClanActivityCounts(friends(), clanSteamID, &online, &inGame, &chatting)
	return online, inGame, chatting, result
}

func DownloadClanActivityCounts(clansToRequest int32) ([]CSteamID, CallResult[DownloadClanActivityCountsResult]) {
	pclanSteamIDs := make([]CSteamID, 0, clansToRequest)
	result := iSteamFriends_DownloadClanActivityCounts(friends(), &pclanSteamIDs, clansToRequest)
	return pclanSteamIDs, CallResult[DownloadClanActivityCountsResult]{Handle: result}
}

func GetFriendCountFromSource(sourceSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendCountFromSource(friends(), sourceSteamID)
}

func GetFriendFromSourceByIndex(sourceSteamID Uint64SteamID, friend int32) Uint64SteamID {
	return iSteamFriends_GetFriendFromSourceByIndex(friends(), sourceSteamID, friend)
}

func IsUserInSource(userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool {
	return iSteamFriends_IsUserInSource(friends(), userSteamID, sourceSteamID)
}

func SetInGameVoiceSpeaking(userSteamID Uint64SteamID, speaking bool) {
	iSteamFriends_SetInGameVoiceSpeaking(friends(), userSteamID, speaking)
}

func ActivateGameOverlay(dialog string) {
	iSteamFriends_ActivateGameOverlay(friends(), dialog)
}

func ActivateGameOverlayToUser(dialog string, steamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayToUser(friends(), dialog, steamID)
}

func ActivateGameOverlayToWebPage(URL string, mode EActivateGameOverlayToWebPageMode) {
	iSteamFriends_ActivateGameOverlayToWebPage(friends(), URL, mode)
}

func ActivateGameOverlayToStore(appID AppId, flag EOverlayToStoreFlag) {
	iSteamFriends_ActivateGameOverlayToStore(friends(), appID, flag)
}

func SetPlayedWith(userPlayedWithSteamID Uint64SteamID) {
	iSteamFriends_SetPlayedWith(friends(), userPlayedWithSteamID)
}

func ActivateGameOverlayInviteDialog(lobbySteamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayInviteDialog(friends(), lobbySteamID)
}

func GetSmallFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetSmallFriendAvatar(friends(), friendSteamID)
}

func GetMediumFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetMediumFriendAvatar(friends(), friendSteamID)
}

func GetLargeFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetLargeFriendAvatar(friends(), friendSteamID)
}

func RequestUserInformation(userSteamID Uint64SteamID, requireNameOnly bool) bool {
	return iSteamFriends_RequestUserInformation(friends(), userSteamID, requireNameOnly)
}

func RequestClanOfficerList(clanSteamID Uint64SteamID) CallResult[ClanOfficerListResponse] {
	handle := iSteamFriends_RequestClanOfficerList(friends(), clanSteamID)
	return CallResult[ClanOfficerListResponse]{Handle: handle}
}

func GetClanOwner(clanSteamID Uint64SteamID) Uint64SteamID {
	return iSteamFriends_GetClanOwner(friends(), clanSteamID)
}

func GetClanOfficerCount(clanSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetClanOfficerCount(friends(), clanSteamID)
}

func GetClanOfficerByIndex(clanSteamID Uint64SteamID, officer int32) Uint64SteamID {
	return iSteamFriends_GetClanOfficerByIndex(friends(), clanSteamID, officer)
}

func GetUserRestrictions() uint32 {
	return iSteamFriends_GetUserRestrictions(friends())
}

func SetRichPresence(key string, value string) bool {
	return iSteamFriends_SetRichPresence(friends(), key, value)
}

func ClearRichPresence() {
	iSteamFriends_ClearRichPresence(friends())
}

func GetFriendRichPresence(friendSteamID Uint64SteamID, key string) []byte {
	return iSteamFriends_GetFriendRichPresence(friends(), friendSteamID, key)
}

func GetFriendRichPresenceKeyCount(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendRichPresenceKeyCount(friends(), friendSteamID)
}

func GetFriendRichPresenceKeyByIndex(friendSteamID Uint64SteamID, keyIndex int32) []byte {
	return iSteamFriends_GetFriendRichPresenceKeyByIndex(friends(), friendSteamID, keyIndex)
}

func RequestFriendRichPresence(friendSteamID Uint64SteamID) {
	iSteamFriends_RequestFriendRichPresence(friends(), friendSteamID)
}

func InviteUserToGame(friendSteamID Uint64SteamID, connectString string) bool {
	return iSteamFriends_InviteUserToGame(friends(), friendSteamID, connectString)
}

func GetCoplayFriendCount() int32 {
	return iSteamFriends_GetCoplayFriendCount(friends())
}

func GetCoplayFriend(coplayFriendIndex int32) Uint64SteamID {
	return iSteamFriends_GetCoplayFriend(friends(), coplayFriendIndex)
}

func GetFriendCoplayTime(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendCoplayTime(friends(), friendSteamID)
}

func GetFriendCoplayGame(friendSteamID Uint64SteamID) AppId {
	return iSteamFriends_GetFriendCoplayGame(friends(), friendSteamID)
}

func JoinClanChatRoom(clanSteamID Uint64SteamID) CallResult[JoinClanChatRoomCompletionResult] {
	handle := iSteamFriends_JoinClanChatRoom(friends(), clanSteamID)
	return CallResult[JoinClanChatRoomCompletionResult]{Handle: handle}

}

func LeaveClanChatRoom(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_LeaveClanChatRoom(friends(), clanSteamID)
}

func GetClanChatMemberCount(clanSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetClanChatMemberCount(friends(), clanSteamID)
}

func GetChatMemberByIndex(clanSteamID Uint64SteamID, userIndex int32) Uint64SteamID {
	return iSteamFriends_GetChatMemberByIndex(friends(), clanSteamID, userIndex)
}

func SendClanChatMessage(clanChatSteamID Uint64SteamID, text string) bool {
	return iSteamFriends_SendClanChatMessage(friends(), clanChatSteamID, text)
}

func GetClanChatMessage(clanChatSteamID Uint64SteamID, messageIndex int32, maxTextSize int32) (chatMsg []byte, chatEntryType EChatEntryType, chatter CSteamID, result int32) {
	chatMsg = make([]byte, 0, maxTextSize)
	result = iSteamFriends_GetClanChatMessage(friends(), clanChatSteamID, messageIndex, &chatMsg, maxTextSize, &chatEntryType, &chatter)
	return chatMsg, chatEntryType, chatter, result
}

func IsClanChatAdmin(clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanChatAdmin(friends(), clanChatSteamID, userSteamID)
}

func IsClanChatWindowOpenInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanChatWindowOpenInSteam(friends(), clanChatSteamID)
}

func OpenClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_OpenClanChatWindowInSteam(friends(), clanChatSteamID)
}

func CloseClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_CloseClanChatWindowInSteam(friends(), clanChatSteamID)
}

func SetListenForFriendsMessages(interceptEnabled bool) bool {
	return iSteamFriends_SetListenForFriendsMessages(friends(), interceptEnabled)
}

func ReplyToFriendMessage(friendSteamID Uint64SteamID, msgToSend string) bool {
	return iSteamFriends_ReplyToFriendMessage(friends(), friendSteamID, msgToSend)
}

func GetFriendMessage(friendSteamID Uint64SteamID, messageID int32, messageMaxSize int32) (data []byte, chatEntryType EChatEntryType, result int32) {
	data = make([]byte, 0, messageMaxSize)
	result = iSteamFriends_GetFriendMessage(friends(), friendSteamID, messageID, &data, messageMaxSize, &chatEntryType)
	return data, chatEntryType, result
}

func GetFollowerCount(steamID Uint64SteamID) CallResult[FriendsGetFollowerCount] {
	handle := iSteamFriends_GetFollowerCount(friends(), steamID)
	return CallResult[FriendsGetFollowerCount]{Handle: handle}
}

func IsFollowing(steamID Uint64SteamID) CallResult[FriendsIsFollowing] {
	handle := iSteamFriends_IsFollowing(friends(), steamID)
	return CallResult[FriendsIsFollowing]{Handle: handle}
}

func EnumerateFollowingList(startIndex uint32) CallResult[FriendsEnumerateFollowingList] {
	handle := iSteamFriends_EnumerateFollowingList(friends(), startIndex)
	return CallResult[FriendsEnumerateFollowingList]{Handle: handle}
}

func IsClanPublic(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanPublic(friends(), clanSteamID)
}

func IsClanOfficialGameGroup(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanOfficialGameGroup(friends(), clanSteamID)
}

func GetNumChatsWithUnreadPriorityMessages() int32 {
	return iSteamFriends_GetNumChatsWithUnreadPriorityMessages(friends())
}

func ActivateGameOverlayRemotePlayTogetherInviteDialog(lobbySteamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog(friends(), lobbySteamID)
}

func RegisterProtocolInOverlayBrowser(protocol string) bool {
	return iSteamFriends_RegisterProtocolInOverlayBrowser(friends(), protocol)
}

func ActivateGameOverlayInviteDialogConnectString(connectString string) {
	iSteamFriends_ActivateGameOverlayInviteDialogConnectString(friends(), connectString)
}

func RequestEquippedProfileItems(steamID Uint64SteamID) CallResult[EquippedProfileItems] {
	handle := iSteamFriends_RequestEquippedProfileItems(friends(), steamID)
	return CallResult[EquippedProfileItems]{Handle: handle}
}

func BHasEquippedProfileItem(steamID Uint64SteamID, itemType ECommunityProfileItemType) bool {
	return iSteamFriends_BHasEquippedProfileItem(friends(), steamID, itemType)
}

func GetProfileItemPropertyString(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) string {
	return iSteamFriends_GetProfileItemPropertyString(friends(), steamID, itemType, prop)
}

func GetProfileItemPropertyUint(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) uint32 {
	return iSteamFriends_GetProfileItemPropertyUint(friends(), steamID, itemType, prop)
}
