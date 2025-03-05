package puregosteamworks

import (
	"fmt"
	"strings"
	"unsafe"
)

//SteamApps

const (
	AppProofOfPurchaseKeyMax int32 = 240
)

const (
	flatAPI_SteamApps                                 = "SteamAPI_SteamApps_v008"
	flatAPI_ISteamApps_BIsSubscribed                  = "SteamAPI_ISteamApps_BIsSubscribed"
	flatAPI_ISteamApps_BIsLowViolence                 = "SteamAPI_ISteamApps_BIsLowViolence"
	flatAPI_ISteamApps_BIsCybercafe                   = "SteamAPI_ISteamApps_BIsCybercafe"
	flatAPI_ISteamApps_BIsVACBanned                   = "SteamAPI_ISteamApps_BIsVACBanned"
	flatAPI_ISteamApps_GetCurrentGameLanguage         = "SteamAPI_ISteamApps_GetCurrentGameLanguage"
	flatAPI_ISteamApps_GetAvailableGameLanguages      = "SteamAPI_ISteamApps_GetAvailableGameLanguages"
	flatAPI_ISteamApps_BIsSubscribedApp               = "SteamAPI_ISteamApps_BIsSubscribedApp"
	flatAPI_ISteamApps_BIsDlcInstalled                = "SteamAPI_ISteamApps_BIsDlcInstalled"
	flatAPI_ISteamApps_GetEarliestPurchaseUnixTime    = "SteamAPI_ISteamApps_GetEarliestPurchaseUnixTime"
	flatAPI_ISteamApps_BIsSubscribedFromFreeWeekend   = "SteamAPI_ISteamApps_BIsSubscribedFromFreeWeekend"
	flatAPI_ISteamApps_GetDLCCount                    = "SteamAPI_ISteamApps_GetDLCCount"
	flatAPI_ISteamApps_BGetDLCDataByIndex             = "SteamAPI_ISteamApps_BGetDLCDataByIndex"
	flatAPI_ISteamApps_InstallDLC                     = "SteamAPI_ISteamApps_InstallDLC"
	flatAPI_ISteamApps_UninstallDLC                   = "SteamAPI_ISteamApps_UninstallDLC"
	flatAPI_ISteamApps_RequestAppProofOfPurchaseKey   = "SteamAPI_ISteamApps_RequestAppProofOfPurchaseKey"
	flatAPI_ISteamApps_GetCurrentBetaName             = "SteamAPI_ISteamApps_GetCurrentBetaName"
	flatAPI_ISteamApps_MarkContentCorrupt             = "SteamAPI_ISteamApps_MarkContentCorrupt"
	flatAPI_ISteamApps_GetInstalledDepots             = "SteamAPI_ISteamApps_GetInstalledDepots"
	flatAPI_ISteamApps_GetAppInstallDir               = "SteamAPI_ISteamApps_GetAppInstallDir"
	flatAPI_ISteamApps_BIsAppInstalled                = "SteamAPI_ISteamApps_BIsAppInstalled"
	flatAPI_ISteamApps_GetAppOwner                    = "SteamAPI_ISteamApps_GetAppOwner"
	flatAPI_ISteamApps_GetLaunchQueryParam            = "SteamAPI_ISteamApps_GetLaunchQueryParam"
	flatAPI_ISteamApps_GetDlcDownloadProgress         = "SteamAPI_ISteamApps_GetDlcDownloadProgress"
	flatAPI_ISteamApps_GetAppBuildId                  = "SteamAPI_ISteamApps_GetAppBuildId"
	flatAPI_ISteamApps_RequestAllProofOfPurchaseKeys  = "SteamAPI_ISteamApps_RequestAllProofOfPurchaseKeys"
	flatAPI_ISteamApps_GetFileDetails                 = "SteamAPI_ISteamApps_GetFileDetails"
	flatAPI_ISteamApps_GetLaunchCommandLine           = "SteamAPI_ISteamApps_GetLaunchCommandLine"
	flatAPI_ISteamApps_BIsSubscribedFromFamilySharing = "SteamAPI_ISteamApps_BIsSubscribedFromFamilySharing"
	flatAPI_ISteamApps_BIsTimedTrial                  = "SteamAPI_ISteamApps_BIsTimedTrial"
	flatAPI_ISteamApps_SetDlcContext                  = "SteamAPI_ISteamApps_SetDlcContext"
	flatAPI_ISteamApps_GetNumBetas                    = "SteamAPI_ISteamApps_GetNumBetas"
	flatAPI_ISteamApps_GetBetaInfo                    = "SteamAPI_ISteamApps_GetBetaInfo"
	flatAPI_ISteamApps_SetActiveBeta                  = "SteamAPI_ISteamApps_SetActiveBeta"
)

var (
	steamApps_init                            func() uintptr
	iSteamApps_BIsSubscribed                  func(steamApps uintptr) bool
	iSteamApps_BIsLowViolence                 func(steamApps uintptr) bool
	iSteamApps_BIsCybercafe                   func(steamApps uintptr) bool
	iSteamApps_BIsVACBanned                   func(steamApps uintptr) bool
	iSteamApps_GetCurrentGameLanguage         func(steamApps uintptr) string
	iSteamApps_GetAvailableGameLanguages      func(steamApps uintptr) string
	iSteamApps_BIsSubscribedApp               func(steamApps uintptr, appID AppId_t) bool
	iSteamApps_BIsDlcInstalled                func(steamApps uintptr, appID AppId_t) bool
	iSteamApps_GetEarliestPurchaseUnixTime    func(steamApps uintptr, appID AppId_t) uint32
	iSteamApps_BIsSubscribedFromFreeWeekend   func(steamApps uintptr) bool
	iSteamApps_GetDLCCount                    func(steamApps uintptr) int32
	iSteamApps_BGetDLCDataByIndex             func(steamApps uintptr, dlcIndex int32, appID *AppId_t, available *bool, name []byte, nameBufferSize int32) bool
	iSteamApps_InstallDLC                     func(steamApps uintptr, appID AppId_t)
	iSteamApps_UninstallDLC                   func(steamApps uintptr, appID AppId_t)
	iSteamApps_RequestAppProofOfPurchaseKey   func(steamApps uintptr, appID AppId_t)
	iSteamApps_GetCurrentBetaName             func(steamApps uintptr, name []byte, nameBufferSize int32) bool
	iSteamApps_MarkContentCorrupt             func(steamApps uintptr, missingFilesOnly bool) bool
	iSteamApps_GetInstalledDepots             func(steamApps uintptr, appID AppId_t, depots []DepotId, maxDepots uint32) uint32
	iSteamApps_GetAppInstallDir               func(steamApps uintptr, appID AppId_t, folder []byte, folderBufferSize uint32) uint32
	iSteamApps_BIsAppInstalled                func(steamApps uintptr, appID AppId_t) bool
	iSteamApps_GetAppOwner                    func(steamApps uintptr) Uint64SteamID
	iSteamApps_GetLaunchQueryParam            func(steamApps uintptr, key string) string
	iSteamApps_GetDlcDownloadProgress         func(steamApps uintptr, appID AppId_t, bytesDownloaded *uint64, bytesTotal *uint64) bool
	iSteamApps_GetAppBuildId                  func(steamApps uintptr) int32
	iSteamApps_RequestAllProofOfPurchaseKeys  func(steamApps uintptr)
	iSteamApps_GetFileDetails                 func(steamApps uintptr, pszFileName string) SteamAPICall
	iSteamApps_GetLaunchCommandLine           func(steamApps uintptr, pszCommandLine []byte, commandLineStringSize int32) int32
	iSteamApps_BIsSubscribedFromFamilySharing func(steamApps uintptr) bool
	iSteamApps_BIsTimedTrial                  func(steamApps uintptr, secondsAllowed *uint32, secondsPlayed *uint32) bool
	iSteamApps_SetDlcContext                  func(steamApps uintptr, appID AppId_t) bool
	iSteamApps_GetNumBetas                    func(steamApps uintptr, available *int32, private *int32) int32
	iSteamApps_GetBetaInfo                    func(steamApps uintptr, betaIndex int32, punFlags *uint32, punBuildID *uint32, pchBetaName []byte, cchBetaName int32, pchDescription []byte, cchDescription int32) bool
	iSteamApps_SetActiveBeta                  func(steamApps uintptr, pchBetaName string) bool
)

type steamApps uintptr

func SteamApps() ISteamApps {
	return steamApps(steamApps_init())
}

func (s steamApps) BIsSubscribed() bool {
	return iSteamApps_BIsSubscribed(uintptr(s))
}

func (s steamApps) BIsLowViolence() bool {
	return iSteamApps_BIsLowViolence(uintptr(s))
}

func (s steamApps) BIsCybercafe() bool {
	return iSteamApps_BIsCybercafe(uintptr(s))
}

func (s steamApps) BIsVACBanned() bool {
	return iSteamApps_BIsVACBanned(uintptr(s))
}

func (s steamApps) GetCurrentGameLanguage() string {
	return iSteamApps_GetCurrentGameLanguage(uintptr(s))
}

func (s steamApps) GetAvailableGameLanguages() []string {
	langs := iSteamApps_GetAvailableGameLanguages(uintptr(s))
	return strings.Split(langs, ",")
}

func (s steamApps) BIsSubscribedApp(appID AppId_t) bool {
	return iSteamApps_BIsSubscribedApp(uintptr(s), appID)
}

func (s steamApps) BIsDlcInstalled(appID AppId_t) bool {
	return iSteamApps_BIsDlcInstalled(uintptr(s), appID)
}

func (s steamApps) GetEarliestPurchaseUnixTime(appID AppId_t) uint32 {
	return iSteamApps_GetEarliestPurchaseUnixTime(uintptr(s), appID)
}

func (s steamApps) BIsSubscribedFromFreeWeekend() bool {
	return iSteamApps_BIsSubscribedFromFreeWeekend(uintptr(s))
}

func (s steamApps) GetDLCCount() int32 {
	return iSteamApps_GetDLCCount(uintptr(s))
}

func (s steamApps) BGetDLCDataByIndex(dlcIndex int32) (appID AppId_t, available bool, name string, success bool) {
	var nameBuf []byte = make([]byte, 0, 4096)
	result := iSteamApps_BGetDLCDataByIndex(uintptr(s), dlcIndex, &appID, &available, nameBuf, int32(len(name)))
	return appID, available, string(nameBuf), result
}

func (s steamApps) InstallDLC(appID AppId_t) {
	iSteamApps_InstallDLC(uintptr(s), appID)
}

func (s steamApps) UninstallDLC(appID AppId_t) {
	iSteamApps_UninstallDLC(uintptr(s), appID)
}

func (s steamApps) RequestAppProofOfPurchaseKey(appID AppId_t) {
	iSteamApps_RequestAppProofOfPurchaseKey(uintptr(s), appID)
}

func (s steamApps) GetCurrentBetaName() (name string, success bool) {
	var nameBufferSize int32 = 2048
	nameBuf := make([]byte, 0, nameBufferSize)
	success = iSteamApps_GetCurrentBetaName(uintptr(s), nameBuf, nameBufferSize)
	return string(name), success
}

func (s steamApps) MarkContentCorrupt(missingFilesOnly bool) bool {
	return iSteamApps_MarkContentCorrupt(uintptr(s), missingFilesOnly)
}

func (s steamApps) GetInstalledDepots(appID AppId_t, maxDepots uint32) []DepotId {
	var depots []DepotId = make([]DepotId, 0, maxDepots)
	result := iSteamApps_GetInstalledDepots(uintptr(s), appID, depots, maxDepots)
	return depots[:result]
}

func (s steamApps) GetAppInstallDir(appID AppId_t) string {
	var folderBufferSize uint32 = 4096
	folder := make([]byte, 0, folderBufferSize)
	total := iSteamApps_GetAppInstallDir(uintptr(s), appID, folder, folderBufferSize)
	return string(folder[:total])
}

func (s steamApps) BIsAppInstalled(appID AppId_t) bool {
	return iSteamApps_BIsAppInstalled(uintptr(s), appID)
}

func (s steamApps) GetAppOwner() Uint64SteamID {
	return iSteamApps_GetAppOwner(uintptr(s))
}

func (s steamApps) GetLaunchQueryParam(key string) string {
	return iSteamApps_GetLaunchQueryParam(uintptr(s), key)
}

func (s steamApps) GetDlcDownloadProgress(appID AppId_t) (bytesDownloaded uint64, bytesTotal uint64, success bool) {
	success = iSteamApps_GetDlcDownloadProgress(uintptr(s), appID, &bytesDownloaded, &bytesTotal)
	return bytesDownloaded, bytesTotal, success
}

func (s steamApps) GetAppBuildId() int32 {
	return iSteamApps_GetAppBuildId(uintptr(s))
}

func (s steamApps) RequestAllProofOfPurchaseKeys() {
	iSteamApps_RequestAllProofOfPurchaseKeys(uintptr(s))
}

func (s steamApps) GetFileDetails(pszFileName string) CallResult[FileDetailsResult] {
	handle := iSteamApps_GetFileDetails(uintptr(s), pszFileName)
	return CallResult[FileDetailsResult]{Handle: handle}
}

func (s steamApps) GetLaunchCommandLine() string {
	var commandLineStringSize int32 = 256
	var pszCommandLine []byte = make([]byte, 0, commandLineStringSize)
	total := iSteamApps_GetLaunchCommandLine(uintptr(s), pszCommandLine, commandLineStringSize)
	return string(pszCommandLine[:total])
}

func (s steamApps) BIsSubscribedFromFamilySharing() bool {
	return iSteamApps_BIsSubscribedFromFamilySharing(uintptr(s))
}

func (s steamApps) BIsTimedTrial() (inTimeTrial bool, secondsAllowed uint32, secondsPlayed uint32) {
	inTimeTrial = iSteamApps_BIsTimedTrial(uintptr(s), &secondsAllowed, &secondsPlayed)
	return inTimeTrial, secondsAllowed, secondsPlayed
}

func (s steamApps) SetDlcContext(appID AppId_t) bool {
	return iSteamApps_SetDlcContext(uintptr(s), appID)
}

func (s steamApps) GetNumBetas() (available int32, private int32, totalAppBranches int32) {
	totalAppBranches = iSteamApps_GetNumBetas(uintptr(s), &available, &private)
	return available, private, totalAppBranches
}

func (s steamApps) GetBetaInfo(betaIndex int32, betaNameSize int32, descriptionSize int32) (punFlags uint32, punBuildID uint32, pchBetaName string, pchDescription string, success bool) {
	var betaName []byte = make([]byte, 0, betaNameSize)
	var description []byte = make([]byte, 0, descriptionSize)

	success = iSteamApps_GetBetaInfo(uintptr(s), betaIndex, &punFlags, &punBuildID, betaName, int32(len(betaName)), description, int32(len(description)))
	pchBetaName = string(betaName)
	pchDescription = string(description)
	return punFlags, punBuildID, pchBetaName, pchDescription, success
}

func (s steamApps) SetActiveBeta(pchBetaName string) bool {
	return iSteamApps_SetActiveBeta(uintptr(s), pchBetaName)
}

// SteamFriends
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
	UserRestriction_None        EUserRestriction = 0
	UserRestriction_Unknown     EUserRestriction = 1
	UserRestriction_AnyChat     EUserRestriction = 2
	UserRestriction_VoiceChat   EUserRestriction = 4
	UserRestriction_GroupChat   EUserRestriction = 8
	UserRestriction_Rating      EUserRestriction = 16
	UserRestriction_GameInvites EUserRestriction = 32
	UserRestriction_Trading     EUserRestriction = 64
)

type ECommunityProfileItemType int32

const (
	ECommunityProfileItemType_AnimatedAvatar        ECommunityProfileItemType = 0
	ECommunityProfileItemType_AvatarFrame           ECommunityProfileItemType = 1
	ECommunityProfileItemType_ProfileModifier       ECommunityProfileItemType = 2
	ECommunityProfileItemType_ProfileBackground     ECommunityProfileItemType = 3
	ECommunityProfileItemType_MiniProfileBackground ECommunityProfileItemType = 4
)

type ECommunityProfileItemProperty int32

const (
	ECommunityProfileItemProperty_ImageSmall     ECommunityProfileItemProperty = 0
	ECommunityProfileItemProperty_ImageLarge     ECommunityProfileItemProperty = 1
	ECommunityProfileItemProperty_InternalName   ECommunityProfileItemProperty = 2
	ECommunityProfileItemProperty_Title          ECommunityProfileItemProperty = 3
	ECommunityProfileItemProperty_Description    ECommunityProfileItemProperty = 4
	ECommunityProfileItemProperty_AppID          ECommunityProfileItemProperty = 5
	ECommunityProfileItemProperty_TypeID         ECommunityProfileItemProperty = 6
	ECommunityProfileItemProperty_Class          ECommunityProfileItemProperty = 7
	ECommunityProfileItemProperty_MovieWebM      ECommunityProfileItemProperty = 8
	ECommunityProfileItemProperty_MovieMP4       ECommunityProfileItemProperty = 9
	ECommunityProfileItemProperty_MovieWebMSmall ECommunityProfileItemProperty = 10
	ECommunityProfileItemProperty_MovieMP4Small  ECommunityProfileItemProperty = 11
)

type EPersonaChange int32

const (
	EPersonaChange_Name                EPersonaChange = 1
	EPersonaChange_Status              EPersonaChange = 2
	EPersonaChange_ComeOnline          EPersonaChange = 4
	EPersonaChange_GoneOffline         EPersonaChange = 8
	EPersonaChange_GamePlayed          EPersonaChange = 16
	EPersonaChange_GameServer          EPersonaChange = 32
	EPersonaChange_Avatar              EPersonaChange = 64
	EPersonaChange_JoinedSource        EPersonaChange = 128
	EPersonaChange_LeftSource          EPersonaChange = 256
	EPersonaChange_RelationshipChanged EPersonaChange = 512
	EPersonaChange_NameFirstSet        EPersonaChange = 1024
	EPersonaChange_Broadcast           EPersonaChange = 2048
	EPersonaChange_Nickname            EPersonaChange = 4096
	EPersonaChange_SteamLevel          EPersonaChange = 8192
	EPersonaChange_RichPresence        EPersonaChange = 16384
)

type FriendsGroupID int16

type FriendGameInfo struct {
	GameID       CGameID
	GameIP       uint32
	GamePort     uint16
	QueryPort    uint16
	LobbySteamID CSteamID
}

type EActivateGameOverlayToWebPageMode int32

const (
	EActivateGameOverlayToWebPageMode_Default EActivateGameOverlayToWebPageMode = 0
	EActivateGameOverlayToWebPageMode_Modal   EActivateGameOverlayToWebPageMode = 1
)

type EOverlayToStoreFlag int32

const (
	EOverlayToStoreFlag_None             EOverlayToStoreFlag = 0
	EOverlayToStoreFlag_AddToCart        EOverlayToStoreFlag = 1
	EOverlayToStoreFlag_AddToCartAndShow EOverlayToStoreFlag = 2
)

type EFriendRelationship int32

const (
	EFriendRelationship_None                EFriendRelationship = 0
	EFriendRelationship_Blocked             EFriendRelationship = 1
	EFriendRelationship_RequestRecipient    EFriendRelationship = 2
	EFriendRelationship_Friend              EFriendRelationship = 3
	EFriendRelationship_RequestInitiator    EFriendRelationship = 4
	EFriendRelationship_Ignored             EFriendRelationship = 5
	EFriendRelationship_IgnoredFriend       EFriendRelationship = 6
	EFriendRelationship_SuggestedDEPRECATED EFriendRelationship = 7
	EFriendRelationship_Max                 EFriendRelationship = 8
)

type EPersonaState int32

const (
	EPersonaState_Offline        EPersonaState = 0
	EPersonaState_Online         EPersonaState = 1
	EPersonaState_Busy           EPersonaState = 2
	EPersonaState_Away           EPersonaState = 3
	EPersonaState_Snooze         EPersonaState = 4
	EPersonaState_LookingToTrade EPersonaState = 5
	EPersonaState_LookingToPlay  EPersonaState = 6
	EPersonaState_Invisible      EPersonaState = 7
	EPersonaState_Max            EPersonaState = 8
)

type EFriendFlags int32

const (
	EFriendFlag_None                 EFriendFlags = 0
	EFriendFlag_Blocked              EFriendFlags = 1
	EFriendFlag_FriendshipRequested  EFriendFlags = 2
	EFriendFlag_Immediate            EFriendFlags = 4
	EFriendFlag_ClanMember           EFriendFlags = 8
	EFriendFlag_OnGameServer         EFriendFlags = 16
	EFriendFlag_RequestingFriendship EFriendFlags = 128
	EFriendFlag_RequestingInfo       EFriendFlags = 256
	EFriendFlag_Ignored              EFriendFlags = 512
	EFriendFlag_IgnoredFriend        EFriendFlags = 1024
	EFriendFlag_ChatMember           EFriendFlags = 4096
	EFriendFlag_All                  EFriendFlags = 65535
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
	iSteamFriends_GetFriendsGroupMembersList                        func(steamFriends uintptr, friendsGroupID FriendsGroupID, pOutSteamIDMembers []CSteamID, nMembersCount int32)
	iSteamFriends_HasFriend                                         func(steamFriends uintptr, friendSteamID Uint64SteamID, iFriendFlags int32) bool
	iSteamFriends_GetClanCount                                      func(steamFriends uintptr) int32
	iSteamFriends_GetClanByIndex                                    func(steamFriends uintptr, iClan int32) Uint64SteamID
	iSteamFriends_GetClanName                                       func(steamFriends uintptr, clanSteamID Uint64SteamID) string
	iSteamFriends_GetClanTag                                        func(steamFriends uintptr, clanSteamID Uint64SteamID) string
	iSteamFriends_GetClanActivityCounts                             func(steamFriends uintptr, clanSteamID Uint64SteamID, pnOnline *int32, pnInGame *int32, pnChatting *int32) bool
	iSteamFriends_DownloadClanActivityCounts                        func(steamFriends uintptr, pclanSteamIDs []CSteamID, cClansToRequest int32) SteamAPICall
	iSteamFriends_GetFriendCountFromSource                          func(steamFriends uintptr, sourceSteamID Uint64SteamID) int32
	iSteamFriends_GetFriendFromSourceByIndex                        func(steamFriends uintptr, sourceSteamID Uint64SteamID, iFriend int32) Uint64SteamID
	iSteamFriends_IsUserInSource                                    func(steamFriends uintptr, userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool
	iSteamFriends_SetInGameVoiceSpeaking                            func(steamFriends uintptr, userSteamID Uint64SteamID, bSpeaking bool)
	iSteamFriends_ActivateGameOverlay                               func(steamFriends uintptr, pchDialog string)
	iSteamFriends_ActivateGameOverlayToUser                         func(steamFriends uintptr, pchDialog string, steamID Uint64SteamID)
	iSteamFriends_ActivateGameOverlayToWebPage                      func(steamFriends uintptr, pchURL string, eMode EActivateGameOverlayToWebPageMode)
	iSteamFriends_ActivateGameOverlayToStore                        func(steamFriends uintptr, nAppID AppId_t, eFlag EOverlayToStoreFlag)
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
	iSteamFriends_GetFriendCoplayGame                               func(steamFriends uintptr, friendSteamID Uint64SteamID) AppId_t
	iSteamFriends_JoinClanChatRoom                                  func(steamFriends uintptr, clanSteamID Uint64SteamID) SteamAPICall
	iSteamFriends_LeaveClanChatRoom                                 func(steamFriends uintptr, clanSteamID Uint64SteamID) bool
	iSteamFriends_GetClanChatMemberCount                            func(steamFriends uintptr, clanSteamID Uint64SteamID) int32
	iSteamFriends_GetChatMemberByIndex                              func(steamFriends uintptr, clanSteamID Uint64SteamID, iUser int32) Uint64SteamID
	iSteamFriends_SendClanChatMessage                               func(steamFriends uintptr, clanChatSteamID Uint64SteamID, pchText string) bool
	iSteamFriends_GetClanChatMessage                                func(steamFriends uintptr, clanChatSteamID Uint64SteamID, iMessage int32, prgchText []byte, cchTextMax int32, peChatEntryType *EChatEntryType, psteamidChatter *CSteamID) int32
	iSteamFriends_IsClanChatAdmin                                   func(steamFriends uintptr, clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool
	iSteamFriends_IsClanChatWindowOpenInSteam                       func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_OpenClanChatWindowInSteam                         func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_CloseClanChatWindowInSteam                        func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool
	iSteamFriends_SetListenForFriendsMessages                       func(steamFriends uintptr, bInterceptEnabled bool) bool
	iSteamFriends_ReplyToFriendMessage                              func(steamFriends uintptr, friendSteamID Uint64SteamID, pchMsgToSend string) bool
	iSteamFriends_GetFriendMessage                                  func(steamFriends uintptr, friendSteamID Uint64SteamID, iMessageID int32, pvData []byte, cubData int32, peChatEntryType *EChatEntryType) int32
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

type steamFriends uintptr

func SteamFriends() ISteamFriends {
	return steamFriends(steamFriends_init())
}
func (s steamFriends) GetPersonaName() string {
	return iSteamFriends_GetPersonaName(uintptr(s))
}

func (s steamFriends) SetPersonaName(personaName string) CallResult[SetPersonaNameResponse] {
	handle := iSteamFriends_SetPersonaName(uintptr(s), personaName)
	return CallResult[SetPersonaNameResponse]{Handle: handle}
}

func (s steamFriends) GetPersonaState() EPersonaState {
	return iSteamFriends_GetPersonaState(uintptr(s))
}

func (s steamFriends) GetFriendCount(friendFlags int32) int32 {
	return iSteamFriends_GetFriendCount(uintptr(s), friendFlags)
}

func (s steamFriends) GetFriendByIndex(friend int32, friendFlags int32) Uint64SteamID {
	return iSteamFriends_GetFriendByIndex(uintptr(s), friend, friendFlags)
}

func (s steamFriends) GetFriendRelationship(friendSteamID Uint64SteamID) EFriendRelationship {
	return iSteamFriends_GetFriendRelationship(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendPersonaState(friendSteamID Uint64SteamID) EPersonaState {
	return iSteamFriends_GetFriendPersonaState(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendPersonaName(friendSteamID Uint64SteamID) string {
	return iSteamFriends_GetFriendPersonaName(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendGamePlayed(friendSteamID Uint64SteamID) (FriendGameInfo, bool) {
	var pFriendGameInfo FriendGameInfo
	result := iSteamFriends_GetFriendGamePlayed(uintptr(s), friendSteamID, &pFriendGameInfo)
	return pFriendGameInfo, result
}

func (s steamFriends) GetFriendPersonaNameHistory(friendSteamID Uint64SteamID, personaName int32) string {
	return iSteamFriends_GetFriendPersonaNameHistory(uintptr(s), friendSteamID, personaName)
}

func (s steamFriends) GetFriendSteamLevel(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendSteamLevel(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendsGroupCount() int32 {
	return iSteamFriends_GetFriendsGroupCount(uintptr(s))
}

func (s steamFriends) GetFriendsGroupIDByIndex(index int32) FriendsGroupID {
	return iSteamFriends_GetFriendsGroupIDByIndex(uintptr(s), index)
}

func (s steamFriends) GetFriendsGroupName(friendsGroupID FriendsGroupID) string {
	return iSteamFriends_GetFriendsGroupName(uintptr(s), friendsGroupID)
}

func (s steamFriends) GetFriendsGroupMembersCount(friendsGroupID FriendsGroupID) int32 {
	return iSteamFriends_GetFriendsGroupMembersCount(uintptr(s), friendsGroupID)
}

// Calls GetFriendsGroupMembersCount to get exact number of friends in group
func (s steamFriends) GetFriendsGroupMembersList(friendsGroupID FriendsGroupID) (members []CSteamID) {
	count := s.GetFriendsGroupMembersCount(friendsGroupID)
	members = make([]CSteamID, 0, count)
	iSteamFriends_GetFriendsGroupMembersList(uintptr(s), friendsGroupID, members, count)
	return members
}

func (s steamFriends) HasFriend(friendSteamID Uint64SteamID, friendFlags int32) bool {
	return iSteamFriends_HasFriend(uintptr(s), friendSteamID, friendFlags)
}

func (s steamFriends) GetClanCount() int32 {
	return iSteamFriends_GetClanCount(uintptr(s))
}

func (s steamFriends) GetClanByIndex(index int32) Uint64SteamID {
	return iSteamFriends_GetClanByIndex(uintptr(s), index)
}

func (s steamFriends) GetClanName(clanSteamID Uint64SteamID) string {
	return iSteamFriends_GetClanName(uintptr(s), clanSteamID)
}

func (s steamFriends) GetClanTag(clanSteamID Uint64SteamID) string {
	return iSteamFriends_GetClanTag(uintptr(s), clanSteamID)
}

func (s steamFriends) GetClanActivityCounts(clanSteamID Uint64SteamID) (online int32, inGame int32, chatting int32, success bool) {
	result := iSteamFriends_GetClanActivityCounts(uintptr(s), clanSteamID, &online, &inGame, &chatting)
	return online, inGame, chatting, result
}

func (s steamFriends) DownloadClanActivityCounts(clansToRequest int32) ([]CSteamID, CallResult[DownloadClanActivityCountsResult]) {
	pclanSteamIDs := make([]CSteamID, 0, clansToRequest)
	result := iSteamFriends_DownloadClanActivityCounts(uintptr(s), pclanSteamIDs, clansToRequest)
	return pclanSteamIDs, CallResult[DownloadClanActivityCountsResult]{Handle: result}
}

func (s steamFriends) GetFriendCountFromSource(sourceSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendCountFromSource(uintptr(s), sourceSteamID)
}

func (s steamFriends) GetFriendFromSourceByIndex(sourceSteamID Uint64SteamID, friend int32) Uint64SteamID {
	return iSteamFriends_GetFriendFromSourceByIndex(uintptr(s), sourceSteamID, friend)
}

func (s steamFriends) IsUserInSource(userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool {
	return iSteamFriends_IsUserInSource(uintptr(s), userSteamID, sourceSteamID)
}

func (s steamFriends) SetInGameVoiceSpeaking(userSteamID Uint64SteamID, speaking bool) {
	iSteamFriends_SetInGameVoiceSpeaking(uintptr(s), userSteamID, speaking)
}

func (s steamFriends) ActivateGameOverlay(dialog string) {
	iSteamFriends_ActivateGameOverlay(uintptr(s), dialog)
}

func (s steamFriends) ActivateGameOverlayToUser(dialog string, steamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayToUser(uintptr(s), dialog, steamID)
}

func (s steamFriends) ActivateGameOverlayToWebPage(URL string, mode EActivateGameOverlayToWebPageMode) {
	iSteamFriends_ActivateGameOverlayToWebPage(uintptr(s), URL, mode)
}

func (s steamFriends) ActivateGameOverlayToStore(appID AppId_t, flag EOverlayToStoreFlag) {
	iSteamFriends_ActivateGameOverlayToStore(uintptr(s), appID, flag)
}

func (s steamFriends) SetPlayedWith(userPlayedWithSteamID Uint64SteamID) {
	iSteamFriends_SetPlayedWith(uintptr(s), userPlayedWithSteamID)
}

func (s steamFriends) ActivateGameOverlayInviteDialog(lobbySteamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayInviteDialog(uintptr(s), lobbySteamID)
}

func (s steamFriends) GetSmallFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetSmallFriendAvatar(uintptr(s), friendSteamID)
}

func (s steamFriends) GetMediumFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetMediumFriendAvatar(uintptr(s), friendSteamID)
}

func (s steamFriends) GetLargeFriendAvatar(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetLargeFriendAvatar(uintptr(s), friendSteamID)
}

func (s steamFriends) RequestUserInformation(userSteamID Uint64SteamID, requireNameOnly bool) bool {
	return iSteamFriends_RequestUserInformation(uintptr(s), userSteamID, requireNameOnly)
}

func (s steamFriends) RequestClanOfficerList(clanSteamID Uint64SteamID) CallResult[ClanOfficerListResponse] {
	handle := iSteamFriends_RequestClanOfficerList(uintptr(s), clanSteamID)
	return CallResult[ClanOfficerListResponse]{Handle: handle}
}

func (s steamFriends) GetClanOwner(clanSteamID Uint64SteamID) Uint64SteamID {
	return iSteamFriends_GetClanOwner(uintptr(s), clanSteamID)
}

func (s steamFriends) GetClanOfficerCount(clanSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetClanOfficerCount(uintptr(s), clanSteamID)
}

func (s steamFriends) GetClanOfficerByIndex(clanSteamID Uint64SteamID, officer int32) Uint64SteamID {
	return iSteamFriends_GetClanOfficerByIndex(uintptr(s), clanSteamID, officer)
}

func (s steamFriends) GetUserRestrictions() uint32 {
	return iSteamFriends_GetUserRestrictions(uintptr(s))
}

func (s steamFriends) SetRichPresence(key string, value string) bool {
	return iSteamFriends_SetRichPresence(uintptr(s), key, value)
}

func (s steamFriends) ClearRichPresence() {
	iSteamFriends_ClearRichPresence(uintptr(s))
}

func (s steamFriends) GetFriendRichPresence(friendSteamID Uint64SteamID, key string) []byte {
	return iSteamFriends_GetFriendRichPresence(uintptr(s), friendSteamID, key)
}

func (s steamFriends) GetFriendRichPresenceKeyCount(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendRichPresenceKeyCount(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendRichPresenceKeyByIndex(friendSteamID Uint64SteamID, keyIndex int32) []byte {
	return iSteamFriends_GetFriendRichPresenceKeyByIndex(uintptr(s), friendSteamID, keyIndex)
}

func (s steamFriends) RequestFriendRichPresence(friendSteamID Uint64SteamID) {
	iSteamFriends_RequestFriendRichPresence(uintptr(s), friendSteamID)
}

func (s steamFriends) InviteUserToGame(friendSteamID Uint64SteamID, connectString string) bool {
	return iSteamFriends_InviteUserToGame(uintptr(s), friendSteamID, connectString)
}

func (s steamFriends) GetCoplayFriendCount() int32 {
	return iSteamFriends_GetCoplayFriendCount(uintptr(s))
}

func (s steamFriends) GetCoplayFriend(coplayFriendIndex int32) Uint64SteamID {
	return iSteamFriends_GetCoplayFriend(uintptr(s), coplayFriendIndex)
}

func (s steamFriends) GetFriendCoplayTime(friendSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetFriendCoplayTime(uintptr(s), friendSteamID)
}

func (s steamFriends) GetFriendCoplayGame(friendSteamID Uint64SteamID) AppId_t {
	return iSteamFriends_GetFriendCoplayGame(uintptr(s), friendSteamID)
}

func (s steamFriends) JoinClanChatRoom(clanSteamID Uint64SteamID) CallResult[JoinClanChatRoomCompletionResult] {
	handle := iSteamFriends_JoinClanChatRoom(uintptr(s), clanSteamID)
	return CallResult[JoinClanChatRoomCompletionResult]{Handle: handle}

}

func (s steamFriends) LeaveClanChatRoom(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_LeaveClanChatRoom(uintptr(s), clanSteamID)
}

func (s steamFriends) GetClanChatMemberCount(clanSteamID Uint64SteamID) int32 {
	return iSteamFriends_GetClanChatMemberCount(uintptr(s), clanSteamID)
}

func (s steamFriends) GetChatMemberByIndex(clanSteamID Uint64SteamID, userIndex int32) Uint64SteamID {
	return iSteamFriends_GetChatMemberByIndex(uintptr(s), clanSteamID, userIndex)
}

func (s steamFriends) SendClanChatMessage(clanChatSteamID Uint64SteamID, text string) bool {
	return iSteamFriends_SendClanChatMessage(uintptr(s), clanChatSteamID, text)
}

func (s steamFriends) GetClanChatMessage(clanChatSteamID Uint64SteamID, messageIndex int32, maxTextSize int32) (chatMsg []byte, chatEntryType EChatEntryType, chatter CSteamID, result int32) {
	chatMsg = make([]byte, 0, maxTextSize)
	result = iSteamFriends_GetClanChatMessage(uintptr(s), clanChatSteamID, messageIndex, chatMsg, maxTextSize, &chatEntryType, &chatter)
	return chatMsg, chatEntryType, chatter, result
}

func (s steamFriends) IsClanChatAdmin(clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanChatAdmin(uintptr(s), clanChatSteamID, userSteamID)
}

func (s steamFriends) IsClanChatWindowOpenInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanChatWindowOpenInSteam(uintptr(s), clanChatSteamID)
}

func (s steamFriends) OpenClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_OpenClanChatWindowInSteam(uintptr(s), clanChatSteamID)
}

func (s steamFriends) CloseClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool {
	return iSteamFriends_CloseClanChatWindowInSteam(uintptr(s), clanChatSteamID)
}

func (s steamFriends) SetListenForFriendsMessages(interceptEnabled bool) bool {
	return iSteamFriends_SetListenForFriendsMessages(uintptr(s), interceptEnabled)
}

func (s steamFriends) ReplyToFriendMessage(friendSteamID Uint64SteamID, msgToSend string) bool {
	return iSteamFriends_ReplyToFriendMessage(uintptr(s), friendSteamID, msgToSend)
}

func (s steamFriends) GetFriendMessage(friendSteamID Uint64SteamID, messageID int32, messageMaxSize int32) (data []byte, chatEntryType EChatEntryType, result int32) {
	data = make([]byte, 0, messageMaxSize)
	result = iSteamFriends_GetFriendMessage(uintptr(s), friendSteamID, messageID, data, messageMaxSize, &chatEntryType)
	return data, chatEntryType, result
}

func (s steamFriends) GetFollowerCount(steamID Uint64SteamID) CallResult[FriendsGetFollowerCount] {
	handle := iSteamFriends_GetFollowerCount(uintptr(s), steamID)
	return CallResult[FriendsGetFollowerCount]{Handle: handle}
}

func (s steamFriends) IsFollowing(steamID Uint64SteamID) CallResult[FriendsIsFollowing] {
	handle := iSteamFriends_IsFollowing(uintptr(s), steamID)
	return CallResult[FriendsIsFollowing]{Handle: handle}
}

func (s steamFriends) EnumerateFollowingList(startIndex uint32) CallResult[FriendsEnumerateFollowingList] {
	handle := iSteamFriends_EnumerateFollowingList(uintptr(s), startIndex)
	return CallResult[FriendsEnumerateFollowingList]{Handle: handle}
}

func (s steamFriends) IsClanPublic(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanPublic(uintptr(s), clanSteamID)
}

func (s steamFriends) IsClanOfficialGameGroup(clanSteamID Uint64SteamID) bool {
	return iSteamFriends_IsClanOfficialGameGroup(uintptr(s), clanSteamID)
}

func (s steamFriends) GetNumChatsWithUnreadPriorityMessages() int32 {
	return iSteamFriends_GetNumChatsWithUnreadPriorityMessages(uintptr(s))
}

func (s steamFriends) ActivateGameOverlayRemotePlayTogetherInviteDialog(lobbySteamID Uint64SteamID) {
	iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog(uintptr(s), lobbySteamID)
}

func (s steamFriends) RegisterProtocolInOverlayBrowser(protocol string) bool {
	return iSteamFriends_RegisterProtocolInOverlayBrowser(uintptr(s), protocol)
}

func (s steamFriends) ActivateGameOverlayInviteDialogConnectString(connectString string) {
	iSteamFriends_ActivateGameOverlayInviteDialogConnectString(uintptr(s), connectString)
}

func (s steamFriends) RequestEquippedProfileItems(steamID Uint64SteamID) CallResult[EquippedProfileItems] {
	handle := iSteamFriends_RequestEquippedProfileItems(uintptr(s), steamID)
	return CallResult[EquippedProfileItems]{Handle: handle}
}

func (s steamFriends) BHasEquippedProfileItem(steamID Uint64SteamID, itemType ECommunityProfileItemType) bool {
	return iSteamFriends_BHasEquippedProfileItem(uintptr(s), steamID, itemType)
}

func (s steamFriends) GetProfileItemPropertyString(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) string {
	return iSteamFriends_GetProfileItemPropertyString(uintptr(s), steamID, itemType, prop)
}

func (s steamFriends) GetProfileItemPropertyUint(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) uint32 {
	return iSteamFriends_GetProfileItemPropertyUint(uintptr(s), steamID, itemType, prop)
}

// SteamClient
type EAccountType int32

const (
	EAccountType_Invalid        EAccountType = 0
	EAccountType_Individual     EAccountType = 1
	EAccountType_Multiseat      EAccountType = 2
	EAccountType_GameServer     EAccountType = 3
	EAccountType_AnonGameServer EAccountType = 4
	EAccountType_Pending        EAccountType = 5
	EAccountType_ContentServer  EAccountType = 6
	EAccountType_Clan           EAccountType = 7
	EAccountType_Chat           EAccountType = 8
	EAccountType_ConsoleUser    EAccountType = 9
	EAccountType_AnonUser       EAccountType = 10
	EAccountType_Max            EAccountType = 11
)

var (
	iSteamClient_CreateSteamPipe             func() HSteamPipe
	iSteamClient_BReleaseSteamPipe           func(hSteamPipe HSteamPipe) bool
	iSteamClient_ConnectToGlobalUser         func(hSteamPipe HSteamPipe) HSteamUser
	iSteamClient_CreateLocalUser             func(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser
	iSteamClient_ReleaseUser                 func(hSteamPipe HSteamPipe, hUser HSteamUser)
	iSteamClient_GetISteamUser               func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameServer         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_SetLocalIPBinding           func(unIP *SteamIPAddress, usPort uint16)
	iSteamClient_GetISteamFriends            func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamUtils              func(hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamMatchmaking        func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamMatchmakingServers func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGenericInterface   func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer
	iSteamClient_GetISteamUserStats          func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameServerStats    func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamApps               func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamNetworking         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamRemoteStorage      func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamScreenshots        func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamGameSearch         func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetIPCCallCount             func() uint32
	iSteamClient_SetWarningMessageHook       func(pFunction SteamAPIWarningMessageHook)
	iSteamClient_BShutdownIfAllPipesClosed   func() bool
	iSteamClient_GetISteamUGC                func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamInventory          func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamInput              func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamParties            func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
	iSteamClient_GetISteamRemotePlay         func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr
)

const (
	flatAPI_ISteamClient_CreateSteamPipe             = "SteamAPI_ISteamClient_CreateSteamPipe"
	flatAPI_ISteamClient_BReleaseSteamPipe           = "SteamAPI_ISteamClient_BReleaseSteamPipe"
	flatAPI_ISteamClient_ConnectToGlobalUser         = "SteamAPI_ISteamClient_ConnectToGlobalUser"
	flatAPI_ISteamClient_CreateLocalUser             = "SteamAPI_ISteamClient_CreateLocalUser"
	flatAPI_ISteamClient_ReleaseUser                 = "SteamAPI_ISteamClient_ReleaseUser"
	flatAPI_ISteamClient_GetISteamUser               = "SteamAPI_ISteamClient_GetISteamUser"
	flatAPI_ISteamClient_GetISteamGameServer         = "SteamAPI_ISteamClient_GetISteamGameServer"
	flatAPI_ISteamClient_SetLocalIPBinding           = "SteamAPI_ISteamClient_SetLocalIPBinding"
	flatAPI_ISteamClient_GetISteamFriends            = "SteamAPI_ISteamClient_GetISteamFriends"
	flatAPI_ISteamClient_GetISteamUtils              = "SteamAPI_ISteamClient_GetISteamUtils"
	flatAPI_ISteamClient_GetISteamMatchmaking        = "SteamAPI_ISteamClient_GetISteamMatchmaking"
	flatAPI_ISteamClient_GetISteamMatchmakingServers = "SteamAPI_ISteamClient_GetISteamMatchmakingServers"
	flatAPI_ISteamClient_GetISteamGenericInterface   = "SteamAPI_ISteamClient_GetISteamGenericInterface"
	flatAPI_ISteamClient_GetISteamUserStats          = "SteamAPI_ISteamClient_GetISteamUserStats"
	flatAPI_ISteamClient_GetISteamGameServerStats    = "SteamAPI_ISteamClient_GetISteamGameServerStats"
	flatAPI_ISteamClient_GetISteamApps               = "SteamAPI_ISteamClient_GetISteamApps"
	flatAPI_ISteamClient_GetISteamNetworking         = "SteamAPI_ISteamClient_GetISteamNetworking"
	flatAPI_ISteamClient_GetISteamRemoteStorage      = "SteamAPI_ISteamClient_GetISteamRemoteStorage"
	flatAPI_ISteamClient_GetISteamScreenshots        = "SteamAPI_ISteamClient_GetISteamScreenshots"
	flatAPI_ISteamClient_GetISteamGameSearch         = "SteamAPI_ISteamClient_GetISteamGameSearch"
	flatAPI_ISteamClient_GetIPCCallCount             = "SteamAPI_ISteamClient_GetIPCCallCount"
	flatAPI_ISteamClient_SetWarningMessageHook       = "SteamAPI_ISteamClient_SetWarningMessageHook"
	flatAPI_ISteamClient_BShutdownIfAllPipesClosed   = "SteamAPI_ISteamClient_BShutdownIfAllPipesClosed"
	flatAPI_ISteamClient_GetISteamUGC                = "SteamAPI_ISteamClient_GetISteamUGC"
	flatAPI_ISteamClient_GetISteamInventory          = "SteamAPI_ISteamClient_GetISteamInventory"
	flatAPI_ISteamClient_GetISteamInput              = "SteamAPI_ISteamClient_GetISteamInput"
	flatAPI_ISteamClient_GetISteamParties            = "SteamAPI_ISteamClient_GetISteamParties"
	flatAPI_ISteamClient_GetISteamRemotePlay         = "SteamAPI_ISteamClient_GetISteamRemotePlay"
)

type steamClient uintptr

func SteamClient() ISteamClient {
	var Zero steamClient
	return Zero
}
func (s steamClient) CreateSteamPipe() HSteamPipe {
	return iSteamClient_CreateSteamPipe()
}

func (s steamClient) BReleaseSteamPipe(hSteamPipe HSteamPipe) bool {
	return iSteamClient_BReleaseSteamPipe(hSteamPipe)
}

func (s steamClient) ConnectToGlobalUser(hSteamPipe HSteamPipe) HSteamUser {
	return iSteamClient_ConnectToGlobalUser(hSteamPipe)
}

func (s steamClient) CreateLocalUser(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser {
	return iSteamClient_CreateLocalUser(phSteamPipe, eAccountType)
}

func (s steamClient) ReleaseUser(hSteamPipe HSteamPipe, hUser HSteamUser) {
	iSteamClient_ReleaseUser(hSteamPipe, hUser)
}

func (s steamClient) SetLocalIPBinding(unIP *SteamIPAddress, usPort uint16) {
	iSteamClient_SetLocalIPBinding(unIP, usPort)
}

func (s steamClient) GetIPCCallCount() uint32 {
	return iSteamClient_GetIPCCallCount()
}

func (s steamClient) SetWarningMessageHook(pFunction SteamAPIWarningMessageHook) {
	iSteamClient_SetWarningMessageHook(pFunction)
}

func (s steamClient) BShutdownIfAllPipesClosed() bool {
	return iSteamClient_BShutdownIfAllPipesClosed()
}

// Steam Game Search
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
	EPlayerResult_FailedToConnect EPlayerResult = 1
	EPlayerResult_Abandoned       EPlayerResult = 2
	EPlayerResult_Kicked          EPlayerResult = 3
	EPlayerResult_Incomplete      EPlayerResult = 4
	EPlayerResult_Completed       EPlayerResult = 5
)

type EGameSearchErrorCode int32

const (
	EGameSearchErrorCode_OK                            EGameSearchErrorCode = 1
	EGameSearchErrorCode_FailedSearchAlreadyInProgress EGameSearchErrorCode = 2
	EGameSearchErrorCode_FailedNoSearchInProgress      EGameSearchErrorCode = 3
	EGameSearchErrorCode_FailedNotLobbyLeader          EGameSearchErrorCode = 4
	EGameSearchErrorCode_FailedNoHostAvailable         EGameSearchErrorCode = 5
	EGameSearchErrorCode_FailedSearchParamsInvalid     EGameSearchErrorCode = 6
	EGameSearchErrorCode_FailedOffline                 EGameSearchErrorCode = 7
	EGameSearchErrorCode_FailedNotAuthorized           EGameSearchErrorCode = 8
	EGameSearchErrorCode_FailedUnknownError            EGameSearchErrorCode = 9
)

var (
	steamGameSearch_init                         func() uintptr
	iSteamGameSearch_AddGameSearchParams         func(steamGameSearch uintptr, keyToFind string, valuesToFind string) EGameSearchErrorCode
	iSteamGameSearch_SearchForGameWithLobby      func(steamGameSearch uintptr, lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode
	iSteamGameSearch_SearchForGameSolo           func(steamGameSearch uintptr, playerMin int32, playerMax int32) EGameSearchErrorCode
	iSteamGameSearch_AcceptGame                  func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_DeclineGame                 func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_RetrieveConnectionDetails   func(steamGameSearch uintptr, hostSteamID Uint64SteamID, cpchConnectionDetails []byte, connectionDetailsSize int32) EGameSearchErrorCode
	iSteamGameSearch_EndGameSearch               func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_SetGameHostParams           func(steamGameSearch uintptr, key string, value string) EGameSearchErrorCode
	iSteamGameSearch_SetConnectionDetails        func(steamGameSearch uintptr, cpchConnectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode
	iSteamGameSearch_RequestPlayersForGame       func(steamGameSearch uintptr, playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode
	iSteamGameSearch_HostConfirmGameStart        func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode
	iSteamGameSearch_CancelRequestPlayersForGame func(steamGameSearch uintptr) EGameSearchErrorCode
	iSteamGameSearch_SubmitPlayerResult          func(steamGameSearch uintptr, uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode
	iSteamGameSearch_EndGame                     func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode
)

type steamGameSearch uintptr

func SteamGameSearch() ISteamGameSearch {
	return steamGameSearch(steamGameSearch_init())
}

func (s steamGameSearch) AddGameSearchParams(keyToFind string, valuesToFind string) EGameSearchErrorCode {
	return iSteamGameSearch_AddGameSearchParams(uintptr(s), keyToFind, valuesToFind)
}

func (s steamGameSearch) SearchForGameWithLobby(lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode {
	return iSteamGameSearch_SearchForGameWithLobby(uintptr(s), lobbySteamID, playerMin, playerMax)
}

func (s steamGameSearch) SearchForGameSolo(playerMin int32, playerMax int32) EGameSearchErrorCode {
	return iSteamGameSearch_SearchForGameSolo(uintptr(s), playerMin, playerMax)
}

func (s steamGameSearch) AcceptGame() EGameSearchErrorCode {
	return iSteamGameSearch_AcceptGame(uintptr(s))
}

func (s steamGameSearch) DeclineGame() EGameSearchErrorCode {
	return iSteamGameSearch_DeclineGame(uintptr(s))
}

func (s steamGameSearch) RetrieveConnectionDetails(hostSteamID Uint64SteamID, connectionDetailsSize int32) (string, EGameSearchErrorCode) {
	var details []byte = make([]byte, 0, connectionDetailsSize)
	code := iSteamGameSearch_RetrieveConnectionDetails(uintptr(s), hostSteamID, details, connectionDetailsSize)
	return string(details), code
}

func (s steamGameSearch) EndGameSearch() EGameSearchErrorCode {
	return iSteamGameSearch_EndGameSearch(uintptr(s))
}

func (s steamGameSearch) SetGameHostParams(key string, value string) EGameSearchErrorCode {
	return iSteamGameSearch_SetGameHostParams(uintptr(s), key, value)
}

func (s steamGameSearch) SetConnectionDetails(connectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode {
	return iSteamGameSearch_SetConnectionDetails(uintptr(s), connectionDetails, connectionDetailsSize)
}

func (s steamGameSearch) RequestPlayersForGame(playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode {
	return iSteamGameSearch_RequestPlayersForGame(uintptr(s), playerMin, playerMax, maxTeamSize)
}

func (s steamGameSearch) HostConfirmGameStart(uniqueGameID uint64) EGameSearchErrorCode {
	return iSteamGameSearch_HostConfirmGameStart(uintptr(s), uniqueGameID)
}

func (s steamGameSearch) CancelRequestPlayersForGame() EGameSearchErrorCode {
	return iSteamGameSearch_CancelRequestPlayersForGame(uintptr(s))
}

func (s steamGameSearch) SubmitPlayerResult(uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode {
	return iSteamGameSearch_SubmitPlayerResult(uintptr(s), uniqueGameID, playerSteamID, EPlayerResult)
}

func (s steamGameSearch) EndGame(uniqueGameID uint64) EGameSearchErrorCode {
	return iSteamGameSearch_EndGame(uintptr(s), uniqueGameID)
}

// Steam Game Server
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
	iSteamGameServer_GetSteamID                          func(steamGameServer uintptr) Uint64SteamID
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
	iSteamGameServer_GetAuthSessionTicket                func(steamGameServer uintptr, pTicket []byte, maxTicket int32, pcbTicket *uint32, snid *SteamNetworkingIdentity) HAuthTicket
	iSteamGameServer_BeginAuthSession                    func(steamGameServer uintptr, authTicket []byte, cbAuthTicket int32, steamID Uint64SteamID) EBeginAuthSessionResult
	iSteamGameServer_EndAuthSession                      func(steamGameServer uintptr, steamID Uint64SteamID)
	iSteamGameServer_CancelAuthTicket                    func(steamGameServer uintptr, hAuthTicket HAuthTicket)
	iSteamGameServer_UserHasLicenseForApp                func(steamGameServer uintptr, steamID Uint64SteamID, AppId AppId_t) EUserHasLicenseForAppResult
	iSteamGameServer_RequestUserGroupStatus              func(steamGameServer uintptr, user Uint64SteamID, groupSteamID Uint64SteamID) bool
	iSteamGameServer_GetPublicIP                         func(steamGameServer uintptr) uintptr
	iSteamGameServer_HandleIncomingPacket                func(steamGameServer uintptr, pData []byte, data int32, srcIP uint32, srcPort uint16) bool
	iSteamGameServer_GetNextOutgoingPacket               func(steamGameServer uintptr, pOut []byte, maxOut int32, pNetAdr *uint32, pPort *uint16) int32
	iSteamGameServer_AssociateWithClan                   func(steamGameServer uintptr, clanSteamID Uint64SteamID) SteamAPICall
	iSteamGameServer_ComputeNewPlayerCompatibility       func(steamGameServer uintptr, newPlayerSteamID Uint64SteamID) SteamAPICall
	iSteamGameServer_CreateUnauthenticatedUserConnection func(steamGameServer uintptr) Uint64SteamID
	iSteamGameServer_BUpdateUserData                     func(steamGameServer uintptr, user Uint64SteamID, playerName string, score uint32) bool
)

type steamGameServer uintptr

func SteamGameServer() ISteamGameServer {
	GameServerActive = true
	return steamGameServer(steamGameServer_init())
}

func (s steamGameServer) SetProduct(product string) {
	iSteamGameServer_SetProduct(uintptr(s), product)
}

func (s steamGameServer) SetGameDescription(gameDescription string) {
	iSteamGameServer_SetGameDescription(uintptr(s), gameDescription)
}

func (s steamGameServer) SetModDir(modDir string) {
	iSteamGameServer_SetModDir(uintptr(s), modDir)
}

func (s steamGameServer) SetDedicatedServer(dedicated bool) {
	iSteamGameServer_SetDedicatedServer(uintptr(s), dedicated)
}

func (s steamGameServer) LogOn(token string) {
	iSteamGameServer_LogOn(uintptr(s), token)
}

func (s steamGameServer) LogOnAnonymous() {
	iSteamGameServer_LogOnAnonymous(uintptr(s))
}

func (s steamGameServer) LogOff() {
	iSteamGameServer_LogOff(uintptr(s))
}

func (s steamGameServer) BLoggedOn() bool {
	return iSteamGameServer_BLoggedOn(uintptr(s))
}

func (s steamGameServer) BSecure() bool {
	return iSteamGameServer_BSecure(uintptr(s))
}

func (s steamGameServer) GetSteamID() Uint64SteamID {
	return iSteamGameServer_GetSteamID(uintptr(s))
}

func (s steamGameServer) WasRestartRequested() bool {
	return iSteamGameServer_WasRestartRequested(uintptr(s))
}

func (s steamGameServer) SetMaxPlayerCount(playersMax int32) {
	iSteamGameServer_SetMaxPlayerCount(uintptr(s), playersMax)
}

func (s steamGameServer) SetBotPlayerCount(botplayers int32) {
	iSteamGameServer_SetBotPlayerCount(uintptr(s), botplayers)
}

func (s steamGameServer) SetServerName(serverName string) {
	iSteamGameServer_SetServerName(uintptr(s), serverName)
}

func (s steamGameServer) SetMapName(mapName string) {
	iSteamGameServer_SetMapName(uintptr(s), mapName)
}

func (s steamGameServer) SetPasswordProtected(passwordProtected bool) {
	iSteamGameServer_SetPasswordProtected(uintptr(s), passwordProtected)
}

func (s steamGameServer) SetSpectatorPort(spectatorPort uint16) {
	iSteamGameServer_SetSpectatorPort(uintptr(s), spectatorPort)
}

func (s steamGameServer) SetSpectatorServerName(spectatorServerName string) {
	iSteamGameServer_SetSpectatorServerName(uintptr(s), spectatorServerName)
}

func (s steamGameServer) ClearAllKeyValues() {
	iSteamGameServer_ClearAllKeyValues(uintptr(s))
}

func (s steamGameServer) SetKeyValue(key string, value string) {
	iSteamGameServer_SetKeyValue(uintptr(s), key, value)
}

func (s steamGameServer) SetGameTags(gameTags string) {
	iSteamGameServer_SetGameTags(uintptr(s), gameTags)
}

func (s steamGameServer) SetGameData(gameData string) {
	iSteamGameServer_SetGameData(uintptr(s), gameData)
}

func (s steamGameServer) SetRegion(region string) {
	iSteamGameServer_SetRegion(uintptr(s), region)
}

func (s steamGameServer) SetAdvertiseServerActive(active bool) {
	iSteamGameServer_SetAdvertiseServerActive(uintptr(s), active)
}

// snid optional
func (s steamGameServer) GetAuthSessionTicket(maxTicket int32, snid *SteamNetworkingIdentity) ([]byte, HAuthTicket) {
	var pTicket []byte = make([]byte, 0, maxTicket)
	var pcbTicket uint32
	result := iSteamGameServer_GetAuthSessionTicket(uintptr(s), pTicket, maxTicket, &pcbTicket, snid)
	return pTicket[:pcbTicket], result
}

func (s steamGameServer) BeginAuthSession(authTicket []byte, steamID Uint64SteamID) EBeginAuthSessionResult {
	return iSteamGameServer_BeginAuthSession(uintptr(s), authTicket, int32(len(authTicket)), steamID)
}

func (s steamGameServer) EndAuthSession(steamID Uint64SteamID) {
	iSteamGameServer_EndAuthSession(uintptr(s), steamID)
}

func (s steamGameServer) CancelAuthTicket(authTicket HAuthTicket) {
	iSteamGameServer_CancelAuthTicket(uintptr(s), authTicket)
}

func (s steamGameServer) UserHasLicenseForApp(steamID Uint64SteamID, AppId AppId_t) EUserHasLicenseForAppResult {
	return iSteamGameServer_UserHasLicenseForApp(uintptr(s), steamID, AppId)
}

func (s steamGameServer) RequestUserGroupStatus(userSteamID Uint64SteamID, groupSteamID Uint64SteamID) bool {
	return iSteamGameServer_RequestUserGroupStatus(uintptr(s), userSteamID, groupSteamID)
}

func (s steamGameServer) GetPublicIP() SteamIPAddress {
	return *uintptrToStruct[SteamIPAddress](iSteamGameServer_GetPublicIP(uintptr(s)))
}

func (s steamGameServer) HandleIncomingPacket(data int32, srcIP uint32, srcPort uint16) ([]byte, bool) {
	pData := make([]byte, 0, data)
	result := iSteamGameServer_HandleIncomingPacket(uintptr(s), pData, data, srcIP, srcPort)
	return pData, result
}

func (s steamGameServer) GetNextOutgoingPacket(maxOut int32) (packet []byte, addr uint32, port uint16, result int32) {
	packet = make([]byte, 0, maxOut)
	result = iSteamGameServer_GetNextOutgoingPacket(uintptr(s), packet, maxOut, &addr, &port)
	return packet, addr, port, result
}

func (s steamGameServer) AssociateWithClan(clanSteamID Uint64SteamID) CallResult[AssociateWithClanResult] {
	handle := iSteamGameServer_AssociateWithClan(uintptr(s), clanSteamID)
	return CallResult[AssociateWithClanResult]{Handle: handle}
}

func (s steamGameServer) ComputeNewPlayerCompatibility(newPlayerSteamID Uint64SteamID) CallResult[ComputeNewPlayerCompatibilityResult] {
	handle := iSteamGameServer_ComputeNewPlayerCompatibility(uintptr(s), newPlayerSteamID)
	return CallResult[ComputeNewPlayerCompatibilityResult]{Handle: handle}
}

func (s steamGameServer) CreateUnauthenticatedUserConnection() Uint64SteamID {
	return iSteamGameServer_CreateUnauthenticatedUserConnection(uintptr(s))
}

func (s steamGameServer) BUpdateUserData(user Uint64SteamID, playerName string, score uint32) bool {
	return iSteamGameServer_BUpdateUserData(uintptr(s), user, playerName, score)
}

// Steam Game Server Stats

var (
	steamGameServerStats_init                   func() uintptr
	iSteamGameServerStats_RequestUserStats      func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall
	iSteamGameServerStats_GetUserStatInt32      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pData *int32) bool
	iSteamGameServerStats_GetUserStatFloat      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pData *float32) bool
	iSteamGameServerStats_GetUserAchievement    func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pbAchieved *bool) bool
	iSteamGameServerStats_SetUserStatInt32      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, nData int32) bool
	iSteamGameServerStats_SetUserStatFloat      func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, fData float32) bool
	iSteamGameServerStats_UpdateUserAvgRateStat func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, flCountThisSession float32, dSessionLength float64) bool
	iSteamGameServerStats_SetUserAchievement    func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool
	iSteamGameServerStats_ClearUserAchievement  func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool
	iSteamGameServerStats_StoreUserStats        func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall
)

const (
	flatAPI_SteamGameServerStats                        = "SteamAPI_SteamGameServerStats_v001"
	flatAPI_ISteamGameServerStats_RequestUserStats      = "SteamAPI_ISteamGameServerStats_RequestUserStats"
	flatAPI_ISteamGameServerStats_GetUserStatInt32      = "SteamAPI_ISteamGameServerStats_GetUserStatInt32"
	flatAPI_ISteamGameServerStats_GetUserStatFloat      = "SteamAPI_ISteamGameServerStats_GetUserStatFloat"
	flatAPI_ISteamGameServerStats_GetUserAchievement    = "SteamAPI_ISteamGameServerStats_GetUserAchievement"
	flatAPI_ISteamGameServerStats_SetUserStatInt32      = "SteamAPI_ISteamGameServerStats_SetUserStatInt32"
	flatAPI_ISteamGameServerStats_SetUserStatFloat      = "SteamAPI_ISteamGameServerStats_SetUserStatFloat"
	flatAPI_ISteamGameServerStats_UpdateUserAvgRateStat = "SteamAPI_ISteamGameServerStats_UpdateUserAvgRateStat"
	flatAPI_ISteamGameServerStats_SetUserAchievement    = "SteamAPI_ISteamGameServerStats_SetUserAchievement"
	flatAPI_ISteamGameServerStats_ClearUserAchievement  = "SteamAPI_ISteamGameServerStats_ClearUserAchievement"
	flatAPI_ISteamGameServerStats_StoreUserStats        = "SteamAPI_ISteamGameServerStats_StoreUserStats"
)

type steamGameServerStats uintptr

func SteamGameServerStats() ISteamGameServerStats {
	return steamGameServerStats(steamGameServerStats_init())
}

func (s steamGameServerStats) RequestUserStats(userSteamID Uint64SteamID) CallResult[GSStatsReceived] {
	handle := iSteamGameServerStats_RequestUserStats(uintptr(s), userSteamID)
	return CallResult[GSStatsReceived]{Handle: handle}
}

func (s steamGameServerStats) GetUserStat(userSteamID Uint64SteamID, name string) (int32, bool) {
	var pData int32
	result := iSteamGameServerStats_GetUserStatInt32(uintptr(s), userSteamID, name, &pData)
	return pData, result
}

func (s steamGameServerStats) GetUserStatFloat(userSteamID Uint64SteamID, name string) (float32, bool) {
	var pData float32
	result := iSteamGameServerStats_GetUserStatFloat(uintptr(s), userSteamID, name, &pData)
	return pData, result
}

func (s steamGameServerStats) GetUserAchievement(userSteamID Uint64SteamID, name string) (bool, bool) {
	var pbAchieved bool
	result := iSteamGameServerStats_GetUserAchievement(uintptr(s), userSteamID, name, &pbAchieved)
	return pbAchieved, result
}

func (s steamGameServerStats) SetUserStat(userSteamID Uint64SteamID, name string, nData int32) bool {
	return iSteamGameServerStats_SetUserStatInt32(uintptr(s), userSteamID, name, nData)
}

func (s steamGameServerStats) SetUserStatFloat(userSteamID Uint64SteamID, name string, fData float32) bool {
	return iSteamGameServerStats_SetUserStatFloat(uintptr(s), userSteamID, name, fData)
}

func (s steamGameServerStats) UpdateUserAvgRateStat(userSteamID Uint64SteamID, name string, countThisSession float32, sessionLength float64) bool {
	return iSteamGameServerStats_UpdateUserAvgRateStat(uintptr(s), userSteamID, name, countThisSession, sessionLength)
}

func (s steamGameServerStats) SetUserAchievement(userSteamID Uint64SteamID, name string) bool {
	return iSteamGameServerStats_SetUserAchievement(uintptr(s), userSteamID, name)
}

func (s steamGameServerStats) ClearUserAchievement(userSteamID Uint64SteamID, name string) bool {
	return iSteamGameServerStats_ClearUserAchievement(uintptr(s), userSteamID, name)
}

func (s steamGameServerStats) StoreUserStats(userSteamID Uint64SteamID) CallResult[GSStatsStored] {
	handle := iSteamGameServerStats_StoreUserStats(uintptr(s), userSteamID)
	return CallResult[GSStatsStored]{Handle: handle}
}

// Steam Input
const _STEAM_INPUT_MAX_COUNT = 16

const _STEAM_INPUT_MAX_ORIGINS = 8

const _STEAM_INPUT_MAX_ACTIVE_LAYERS = 16

type ESteamControllerPad int32

const (
	ESteamControllerPadLeft  ESteamControllerPad = 0
	ESteamControllerPadRight ESteamControllerPad = 1
)

type SteamInputActionEvent struct {
	ControllerHandle InputHandle_t
	EventType        ESteamInputActionEventType
	Action           [unsafe.Sizeof(DigitalAction{})]byte
}

func (e *SteamInputActionEvent) AnalogActionPtr() *AnalogAction {
	return (*AnalogAction)(unsafe.Pointer(&e.Action))
}
func (e *SteamInputActionEvent) DigitalActionPtr() *DigitalAction {
	return (*DigitalAction)(unsafe.Pointer(&e.Action))
}

type InputMotionData struct {
	RotQuatX  float32
	RotQuatY  float32
	RotQuatZ  float32
	RotQuatW  float32
	PosAccelX float32
	PosAccelY float32
	PosAccelZ float32
	RotVelX   float32
	RotVelY   float32
	RotVelZ   float32
}
type EInputSourceMode int32

const (
	EInputSourceMode_None           EInputSourceMode = 0
	EInputSourceMode_Dpad           EInputSourceMode = 1
	EInputSourceMode_Buttons        EInputSourceMode = 2
	EInputSourceMode_FourButtons    EInputSourceMode = 3
	EInputSourceMode_AbsoluteMouse  EInputSourceMode = 4
	EInputSourceMode_RelativeMouse  EInputSourceMode = 5
	EInputSourceMode_JoystickMove   EInputSourceMode = 6
	EInputSourceMode_JoystickMouse  EInputSourceMode = 7
	EInputSourceMode_JoystickCamera EInputSourceMode = 8
	EInputSourceMode_ScrollWheel    EInputSourceMode = 9
	EInputSourceMode_Trigger        EInputSourceMode = 10
	EInputSourceMode_TouchMenu      EInputSourceMode = 11
	EInputSourceMode_MouseJoystick  EInputSourceMode = 12
	EInputSourceMode_MouseRegion    EInputSourceMode = 13
	EInputSourceMode_RadialMenu     EInputSourceMode = 14
	EInputSourceMode_SingleButton   EInputSourceMode = 15
	EInputSourceMode_Switches       EInputSourceMode = 16
)

type EXboxOrigin int32

const (
	EXboxOrigin_A                   EXboxOrigin = 0
	EXboxOrigin_B                   EXboxOrigin = 1
	EXboxOrigin_X                   EXboxOrigin = 2
	EXboxOrigin_Y                   EXboxOrigin = 3
	EXboxOrigin_LeftBumper          EXboxOrigin = 4
	EXboxOrigin_RightBumper         EXboxOrigin = 5
	EXboxOrigin_Menu                EXboxOrigin = 6
	EXboxOrigin_View                EXboxOrigin = 7
	EXboxOrigin_LeftTriggerPull     EXboxOrigin = 8
	EXboxOrigin_LeftTriggerClick    EXboxOrigin = 9
	EXboxOrigin_RightTriggerPull    EXboxOrigin = 10
	EXboxOrigin_RightTriggerClick   EXboxOrigin = 11
	EXboxOrigin_LeftStickMove       EXboxOrigin = 12
	EXboxOrigin_LeftStickClick      EXboxOrigin = 13
	EXboxOrigin_LeftStickDPadNorth  EXboxOrigin = 14
	EXboxOrigin_LeftStickDPadSouth  EXboxOrigin = 15
	EXboxOrigin_LeftStickDPadWest   EXboxOrigin = 16
	EXboxOrigin_LeftStickDPadEast   EXboxOrigin = 17
	EXboxOrigin_RightStickMove      EXboxOrigin = 18
	EXboxOrigin_RightStickClick     EXboxOrigin = 19
	EXboxOrigin_RightStickDPadNorth EXboxOrigin = 20
	EXboxOrigin_RightStickDPadSouth EXboxOrigin = 21
	EXboxOrigin_RightStickDPadWest  EXboxOrigin = 22
	EXboxOrigin_RightStickDPadEast  EXboxOrigin = 23
	EXboxOrigin_DPadNorth           EXboxOrigin = 24
	EXboxOrigin_DPadSouth           EXboxOrigin = 25
	EXboxOrigin_DPadWest            EXboxOrigin = 26
	EXboxOrigin_DPadEast            EXboxOrigin = 27
	EXboxOrigin_Count               EXboxOrigin = 28
)

type InputDigitalActionData struct {
	State  bool
	Active bool
}
type InputAnalogActionData struct {
	Mode   EInputSourceMode // int32
	X      float32
	Y      float32
	Active bool
	_      [3]byte
}
type AnalogAction struct {
	ActionHandle     InputAnalogActionHandle
	AnalogActionData InputAnalogActionData
}
type DigitalAction struct {
	ActionHandle      InputDigitalActionHandle
	DigitalActionData InputDigitalActionData
}

type EControllerHapticType int32

const (
	EControllerHapticType_Off   EControllerHapticType = 0
	EControllerHapticType_Tick  EControllerHapticType = 1
	EControllerHapticType_Click EControllerHapticType = 2
)

type ESteamInputConfigurationEnableType int32

const (
	ESteamInputConfigurationEnableType_None        ESteamInputConfigurationEnableType = 0
	ESteamInputConfigurationEnableType_Playstation ESteamInputConfigurationEnableType = 1
	ESteamInputConfigurationEnableType_Xbox        ESteamInputConfigurationEnableType = 2
	ESteamInputConfigurationEnableType_Generic     ESteamInputConfigurationEnableType = 4
	ESteamInputConfigurationEnableType_Switch      ESteamInputConfigurationEnableType = 8
)

type ESteamInputLEDFlag int32

const (
	ESteamInputLEDFlag_SetColor           ESteamInputLEDFlag = 0
	ESteamInputLEDFlag_RestoreUserDefault ESteamInputLEDFlag = 1
)

type ESteamInputGlyphStyle int32

const (
	ESteamInputGlyphStyle_Knockout         ESteamInputGlyphStyle = 0
	ESteamInputGlyphStyle_Light            ESteamInputGlyphStyle = 1
	ESteamInputGlyphStyle_Dark             ESteamInputGlyphStyle = 2
	ESteamInputGlyphStyle_NeutralColorABXY ESteamInputGlyphStyle = 16
	ESteamInputGlyphStyle_SolidABXY        ESteamInputGlyphStyle = 32
)

type ESteamInputActionEventType int32

const (
	ESteamInputActionEventType_DigitalAction ESteamInputActionEventType = 0
	ESteamInputActionEventType_AnalogAction  ESteamInputActionEventType = 1
)

type EControllerHapticLocation int32

const (
	EControllerHapticLocation_Left  EControllerHapticLocation = 1
	EControllerHapticLocation_Right EControllerHapticLocation = 2
	EControllerHapticLocation_Both  EControllerHapticLocation = 3
)

type ScePadTriggerEffectParam struct{}

type ESteamInputGlyphSize int32

const (
	ESteamInputGlyphSize_Small  ESteamInputGlyphSize = 0
	ESteamInputGlyphSize_Medium ESteamInputGlyphSize = 1
	ESteamInputGlyphSize_Large  ESteamInputGlyphSize = 2
	ESteamInputGlyphSize_Count  ESteamInputGlyphSize = 3
)

type EInputActionOrigin int32

const (
	EInputActionOrigin_None                              EInputActionOrigin = 0
	EInputActionOrigin_SteamControllerA                  EInputActionOrigin = 1
	EInputActionOrigin_SteamControllerB                  EInputActionOrigin = 2
	EInputActionOrigin_SteamControllerX                  EInputActionOrigin = 3
	EInputActionOrigin_SteamControllerY                  EInputActionOrigin = 4
	EInputActionOrigin_SteamControllerLeftBumper         EInputActionOrigin = 5
	EInputActionOrigin_SteamControllerRightBumper        EInputActionOrigin = 6
	EInputActionOrigin_SteamControllerLeftGrip           EInputActionOrigin = 7
	EInputActionOrigin_SteamControllerRightGrip          EInputActionOrigin = 8
	EInputActionOrigin_SteamControllerStart              EInputActionOrigin = 9
	EInputActionOrigin_SteamControllerBack               EInputActionOrigin = 10
	EInputActionOrigin_SteamControllerLeftPadTouch       EInputActionOrigin = 11
	EInputActionOrigin_SteamControllerLeftPadSwipe       EInputActionOrigin = 12
	EInputActionOrigin_SteamControllerLeftPadClick       EInputActionOrigin = 13
	EInputActionOrigin_SteamControllerLeftPadDPadNorth   EInputActionOrigin = 14
	EInputActionOrigin_SteamControllerLeftPadDPadSouth   EInputActionOrigin = 15
	EInputActionOrigin_SteamControllerLeftPadDPadWest    EInputActionOrigin = 16
	EInputActionOrigin_SteamControllerLeftPadDPadEast    EInputActionOrigin = 17
	EInputActionOrigin_SteamControllerRightPadTouch      EInputActionOrigin = 18
	EInputActionOrigin_SteamControllerRightPadSwipe      EInputActionOrigin = 19
	EInputActionOrigin_SteamControllerRightPadClick      EInputActionOrigin = 20
	EInputActionOrigin_SteamControllerRightPadDPadNorth  EInputActionOrigin = 21
	EInputActionOrigin_SteamControllerRightPadDPadSouth  EInputActionOrigin = 22
	EInputActionOrigin_SteamControllerRightPadDPadWest   EInputActionOrigin = 23
	EInputActionOrigin_SteamControllerRightPadDPadEast   EInputActionOrigin = 24
	EInputActionOrigin_SteamControllerLeftTriggerPull    EInputActionOrigin = 25
	EInputActionOrigin_SteamControllerLeftTriggerClick   EInputActionOrigin = 26
	EInputActionOrigin_SteamControllerRightTriggerPull   EInputActionOrigin = 27
	EInputActionOrigin_SteamControllerRightTriggerClick  EInputActionOrigin = 28
	EInputActionOrigin_SteamControllerLeftStickMove      EInputActionOrigin = 29
	EInputActionOrigin_SteamControllerLeftStickClick     EInputActionOrigin = 30
	EInputActionOrigin_SteamControllerLeftStickDPadNorth EInputActionOrigin = 31
	EInputActionOrigin_SteamControllerLeftStickDPadSouth EInputActionOrigin = 32
	EInputActionOrigin_SteamControllerLeftStickDPadWest  EInputActionOrigin = 33
	EInputActionOrigin_SteamControllerLeftStickDPadEast  EInputActionOrigin = 34
	EInputActionOrigin_SteamControllerGyroMove           EInputActionOrigin = 35
	EInputActionOrigin_SteamControllerGyroPitch          EInputActionOrigin = 36
	EInputActionOrigin_SteamControllerGyroYaw            EInputActionOrigin = 37
	EInputActionOrigin_SteamControllerGyroRoll           EInputActionOrigin = 38
	EInputActionOrigin_SteamControllerReserved0          EInputActionOrigin = 39
	EInputActionOrigin_SteamControllerReserved1          EInputActionOrigin = 40
	EInputActionOrigin_SteamControllerReserved2          EInputActionOrigin = 41
	EInputActionOrigin_SteamControllerReserved3          EInputActionOrigin = 42
	EInputActionOrigin_SteamControllerReserved4          EInputActionOrigin = 43
	EInputActionOrigin_SteamControllerReserved5          EInputActionOrigin = 44
	EInputActionOrigin_SteamControllerReserved6          EInputActionOrigin = 45
	EInputActionOrigin_SteamControllerReserved7          EInputActionOrigin = 46
	EInputActionOrigin_SteamControllerReserved8          EInputActionOrigin = 47
	EInputActionOrigin_SteamControllerReserved9          EInputActionOrigin = 48
	EInputActionOrigin_SteamControllerReserved10         EInputActionOrigin = 49
	EInputActionOrigin_PS4X                              EInputActionOrigin = 50
	EInputActionOrigin_PS4Circle                         EInputActionOrigin = 51
	EInputActionOrigin_PS4Triangle                       EInputActionOrigin = 52
	EInputActionOrigin_PS4Square                         EInputActionOrigin = 53
	EInputActionOrigin_PS4LeftBumper                     EInputActionOrigin = 54
	EInputActionOrigin_PS4RightBumper                    EInputActionOrigin = 55
	EInputActionOrigin_PS4Options                        EInputActionOrigin = 56
	EInputActionOrigin_PS4Share                          EInputActionOrigin = 57
	EInputActionOrigin_PS4LeftPadTouch                   EInputActionOrigin = 58
	EInputActionOrigin_PS4LeftPadSwipe                   EInputActionOrigin = 59
	EInputActionOrigin_PS4LeftPadClick                   EInputActionOrigin = 60
	EInputActionOrigin_PS4LeftPadDPadNorth               EInputActionOrigin = 61
	EInputActionOrigin_PS4LeftPadDPadSouth               EInputActionOrigin = 62
	EInputActionOrigin_PS4LeftPadDPadWest                EInputActionOrigin = 63
	EInputActionOrigin_PS4LeftPadDPadEast                EInputActionOrigin = 64
	EInputActionOrigin_PS4RightPadTouch                  EInputActionOrigin = 65
	EInputActionOrigin_PS4RightPadSwipe                  EInputActionOrigin = 66
	EInputActionOrigin_PS4RightPadClick                  EInputActionOrigin = 67
	EInputActionOrigin_PS4RightPadDPadNorth              EInputActionOrigin = 68
	EInputActionOrigin_PS4RightPadDPadSouth              EInputActionOrigin = 69
	EInputActionOrigin_PS4RightPadDPadWest               EInputActionOrigin = 70
	EInputActionOrigin_PS4RightPadDPadEast               EInputActionOrigin = 71
	EInputActionOrigin_PS4CenterPadTouch                 EInputActionOrigin = 72
	EInputActionOrigin_PS4CenterPadSwipe                 EInputActionOrigin = 73
	EInputActionOrigin_PS4CenterPadClick                 EInputActionOrigin = 74
	EInputActionOrigin_PS4CenterPadDPadNorth             EInputActionOrigin = 75
	EInputActionOrigin_PS4CenterPadDPadSouth             EInputActionOrigin = 76
	EInputActionOrigin_PS4CenterPadDPadWest              EInputActionOrigin = 77
	EInputActionOrigin_PS4CenterPadDPadEast              EInputActionOrigin = 78
	EInputActionOrigin_PS4LeftTriggerPull                EInputActionOrigin = 79
	EInputActionOrigin_PS4LeftTriggerClick               EInputActionOrigin = 80
	EInputActionOrigin_PS4RightTriggerPull               EInputActionOrigin = 81
	EInputActionOrigin_PS4RightTriggerClick              EInputActionOrigin = 82
	EInputActionOrigin_PS4LeftStickMove                  EInputActionOrigin = 83
	EInputActionOrigin_PS4LeftStickClick                 EInputActionOrigin = 84
	EInputActionOrigin_PS4LeftStickDPadNorth             EInputActionOrigin = 85
	EInputActionOrigin_PS4LeftStickDPadSouth             EInputActionOrigin = 86
	EInputActionOrigin_PS4LeftStickDPadWest              EInputActionOrigin = 87
	EInputActionOrigin_PS4LeftStickDPadEast              EInputActionOrigin = 88
	EInputActionOrigin_PS4RightStickMove                 EInputActionOrigin = 89
	EInputActionOrigin_PS4RightStickClick                EInputActionOrigin = 90
	EInputActionOrigin_PS4RightStickDPadNorth            EInputActionOrigin = 91
	EInputActionOrigin_PS4RightStickDPadSouth            EInputActionOrigin = 92
	EInputActionOrigin_PS4RightStickDPadWest             EInputActionOrigin = 93
	EInputActionOrigin_PS4RightStickDPadEast             EInputActionOrigin = 94
	EInputActionOrigin_PS4DPadNorth                      EInputActionOrigin = 95
	EInputActionOrigin_PS4DPadSouth                      EInputActionOrigin = 96
	EInputActionOrigin_PS4DPadWest                       EInputActionOrigin = 97
	EInputActionOrigin_PS4DPadEast                       EInputActionOrigin = 98
	EInputActionOrigin_PS4GyroMove                       EInputActionOrigin = 99
	EInputActionOrigin_PS4GyroPitch                      EInputActionOrigin = 100
	EInputActionOrigin_PS4GyroYaw                        EInputActionOrigin = 101
	EInputActionOrigin_PS4GyroRoll                       EInputActionOrigin = 102
	EInputActionOrigin_PS4DPadMove                       EInputActionOrigin = 103
	EInputActionOrigin_PS4Reserved1                      EInputActionOrigin = 104
	EInputActionOrigin_PS4Reserved2                      EInputActionOrigin = 105
	EInputActionOrigin_PS4Reserved3                      EInputActionOrigin = 106
	EInputActionOrigin_PS4Reserved4                      EInputActionOrigin = 107
	EInputActionOrigin_PS4Reserved5                      EInputActionOrigin = 108
	EInputActionOrigin_PS4Reserved6                      EInputActionOrigin = 109
	EInputActionOrigin_PS4Reserved7                      EInputActionOrigin = 110
	EInputActionOrigin_PS4Reserved8                      EInputActionOrigin = 111
	EInputActionOrigin_PS4Reserved9                      EInputActionOrigin = 112
	EInputActionOrigin_PS4Reserved10                     EInputActionOrigin = 113
	EInputActionOrigin_XBoxOneA                          EInputActionOrigin = 114
	EInputActionOrigin_XBoxOneB                          EInputActionOrigin = 115
	EInputActionOrigin_XBoxOneX                          EInputActionOrigin = 116
	EInputActionOrigin_XBoxOneY                          EInputActionOrigin = 117
	EInputActionOrigin_XBoxOneLeftBumper                 EInputActionOrigin = 118
	EInputActionOrigin_XBoxOneRightBumper                EInputActionOrigin = 119
	EInputActionOrigin_XBoxOneMenu                       EInputActionOrigin = 120
	EInputActionOrigin_XBoxOneView                       EInputActionOrigin = 121
	EInputActionOrigin_XBoxOneLeftTriggerPull            EInputActionOrigin = 122
	EInputActionOrigin_XBoxOneLeftTriggerClick           EInputActionOrigin = 123
	EInputActionOrigin_XBoxOneRightTriggerPull           EInputActionOrigin = 124
	EInputActionOrigin_XBoxOneRightTriggerClick          EInputActionOrigin = 125
	EInputActionOrigin_XBoxOneLeftStickMove              EInputActionOrigin = 126
	EInputActionOrigin_XBoxOneLeftStickClick             EInputActionOrigin = 127
	EInputActionOrigin_XBoxOneLeftStickDPadNorth         EInputActionOrigin = 128
	EInputActionOrigin_XBoxOneLeftStickDPadSouth         EInputActionOrigin = 129
	EInputActionOrigin_XBoxOneLeftStickDPadWest          EInputActionOrigin = 130
	EInputActionOrigin_XBoxOneLeftStickDPadEast          EInputActionOrigin = 131
	EInputActionOrigin_XBoxOneRightStickMove             EInputActionOrigin = 132
	EInputActionOrigin_XBoxOneRightStickClick            EInputActionOrigin = 133
	EInputActionOrigin_XBoxOneRightStickDPadNorth        EInputActionOrigin = 134
	EInputActionOrigin_XBoxOneRightStickDPadSouth        EInputActionOrigin = 135
	EInputActionOrigin_XBoxOneRightStickDPadWest         EInputActionOrigin = 136
	EInputActionOrigin_XBoxOneRightStickDPadEast         EInputActionOrigin = 137
	EInputActionOrigin_XBoxOneDPadNorth                  EInputActionOrigin = 138
	EInputActionOrigin_XBoxOneDPadSouth                  EInputActionOrigin = 139
	EInputActionOrigin_XBoxOneDPadWest                   EInputActionOrigin = 140
	EInputActionOrigin_XBoxOneDPadEast                   EInputActionOrigin = 141
	EInputActionOrigin_XBoxOneDPadMove                   EInputActionOrigin = 142
	EInputActionOrigin_XBoxOneLeftGripLower              EInputActionOrigin = 143
	EInputActionOrigin_XBoxOneLeftGripUpper              EInputActionOrigin = 144
	EInputActionOrigin_XBoxOneRightGripLower             EInputActionOrigin = 145
	EInputActionOrigin_XBoxOneRightGripUpper             EInputActionOrigin = 146
	EInputActionOrigin_XBoxOneShare                      EInputActionOrigin = 147
	EInputActionOrigin_XBoxOneReserved6                  EInputActionOrigin = 148
	EInputActionOrigin_XBoxOneReserved7                  EInputActionOrigin = 149
	EInputActionOrigin_XBoxOneReserved8                  EInputActionOrigin = 150
	EInputActionOrigin_XBoxOneReserved9                  EInputActionOrigin = 151
	EInputActionOrigin_XBoxOneReserved10                 EInputActionOrigin = 152
	EInputActionOrigin_XBox360A                          EInputActionOrigin = 153
	EInputActionOrigin_XBox360B                          EInputActionOrigin = 154
	EInputActionOrigin_XBox360X                          EInputActionOrigin = 155
	EInputActionOrigin_XBox360Y                          EInputActionOrigin = 156
	EInputActionOrigin_XBox360LeftBumper                 EInputActionOrigin = 157
	EInputActionOrigin_XBox360RightBumper                EInputActionOrigin = 158
	EInputActionOrigin_XBox360Start                      EInputActionOrigin = 159
	EInputActionOrigin_XBox360Back                       EInputActionOrigin = 160
	EInputActionOrigin_XBox360LeftTriggerPull            EInputActionOrigin = 161
	EInputActionOrigin_XBox360LeftTriggerClick           EInputActionOrigin = 162
	EInputActionOrigin_XBox360RightTriggerPull           EInputActionOrigin = 163
	EInputActionOrigin_XBox360RightTriggerClick          EInputActionOrigin = 164
	EInputActionOrigin_XBox360LeftStickMove              EInputActionOrigin = 165
	EInputActionOrigin_XBox360LeftStickClick             EInputActionOrigin = 166
	EInputActionOrigin_XBox360LeftStickDPadNorth         EInputActionOrigin = 167
	EInputActionOrigin_XBox360LeftStickDPadSouth         EInputActionOrigin = 168
	EInputActionOrigin_XBox360LeftStickDPadWest          EInputActionOrigin = 169
	EInputActionOrigin_XBox360LeftStickDPadEast          EInputActionOrigin = 170
	EInputActionOrigin_XBox360RightStickMove             EInputActionOrigin = 171
	EInputActionOrigin_XBox360RightStickClick            EInputActionOrigin = 172
	EInputActionOrigin_XBox360RightStickDPadNorth        EInputActionOrigin = 173
	EInputActionOrigin_XBox360RightStickDPadSouth        EInputActionOrigin = 174
	EInputActionOrigin_XBox360RightStickDPadWest         EInputActionOrigin = 175
	EInputActionOrigin_XBox360RightStickDPadEast         EInputActionOrigin = 176
	EInputActionOrigin_XBox360DPadNorth                  EInputActionOrigin = 177
	EInputActionOrigin_XBox360DPadSouth                  EInputActionOrigin = 178
	EInputActionOrigin_XBox360DPadWest                   EInputActionOrigin = 179
	EInputActionOrigin_XBox360DPadEast                   EInputActionOrigin = 180
	EInputActionOrigin_XBox360DPadMove                   EInputActionOrigin = 181
	EInputActionOrigin_XBox360Reserved1                  EInputActionOrigin = 182
	EInputActionOrigin_XBox360Reserved2                  EInputActionOrigin = 183
	EInputActionOrigin_XBox360Reserved3                  EInputActionOrigin = 184
	EInputActionOrigin_XBox360Reserved4                  EInputActionOrigin = 185
	EInputActionOrigin_XBox360Reserved5                  EInputActionOrigin = 186
	EInputActionOrigin_XBox360Reserved6                  EInputActionOrigin = 187
	EInputActionOrigin_XBox360Reserved7                  EInputActionOrigin = 188
	EInputActionOrigin_XBox360Reserved8                  EInputActionOrigin = 189
	EInputActionOrigin_XBox360Reserved9                  EInputActionOrigin = 190
	EInputActionOrigin_XBox360Reserved10                 EInputActionOrigin = 191
	EInputActionOrigin_SwitchA                           EInputActionOrigin = 192
	EInputActionOrigin_SwitchB                           EInputActionOrigin = 193
	EInputActionOrigin_SwitchX                           EInputActionOrigin = 194
	EInputActionOrigin_SwitchY                           EInputActionOrigin = 195
	EInputActionOrigin_SwitchLeftBumper                  EInputActionOrigin = 196
	EInputActionOrigin_SwitchRightBumper                 EInputActionOrigin = 197
	EInputActionOrigin_SwitchPlus                        EInputActionOrigin = 198
	EInputActionOrigin_SwitchMinus                       EInputActionOrigin = 199
	EInputActionOrigin_SwitchCapture                     EInputActionOrigin = 200
	EInputActionOrigin_SwitchLeftTriggerPull             EInputActionOrigin = 201
	EInputActionOrigin_SwitchLeftTriggerClick            EInputActionOrigin = 202
	EInputActionOrigin_SwitchRightTriggerPull            EInputActionOrigin = 203
	EInputActionOrigin_SwitchRightTriggerClick           EInputActionOrigin = 204
	EInputActionOrigin_SwitchLeftStickMove               EInputActionOrigin = 205
	EInputActionOrigin_SwitchLeftStickClick              EInputActionOrigin = 206
	EInputActionOrigin_SwitchLeftStickDPadNorth          EInputActionOrigin = 207
	EInputActionOrigin_SwitchLeftStickDPadSouth          EInputActionOrigin = 208
	EInputActionOrigin_SwitchLeftStickDPadWest           EInputActionOrigin = 209
	EInputActionOrigin_SwitchLeftStickDPadEast           EInputActionOrigin = 210
	EInputActionOrigin_SwitchRightStickMove              EInputActionOrigin = 211
	EInputActionOrigin_SwitchRightStickClick             EInputActionOrigin = 212
	EInputActionOrigin_SwitchRightStickDPadNorth         EInputActionOrigin = 213
	EInputActionOrigin_SwitchRightStickDPadSouth         EInputActionOrigin = 214
	EInputActionOrigin_SwitchRightStickDPadWest          EInputActionOrigin = 215
	EInputActionOrigin_SwitchRightStickDPadEast          EInputActionOrigin = 216
	EInputActionOrigin_SwitchDPadNorth                   EInputActionOrigin = 217
	EInputActionOrigin_SwitchDPadSouth                   EInputActionOrigin = 218
	EInputActionOrigin_SwitchDPadWest                    EInputActionOrigin = 219
	EInputActionOrigin_SwitchDPadEast                    EInputActionOrigin = 220
	EInputActionOrigin_SwitchProGyroMove                 EInputActionOrigin = 221
	EInputActionOrigin_SwitchProGyroPitch                EInputActionOrigin = 222
	EInputActionOrigin_SwitchProGyroYaw                  EInputActionOrigin = 223
	EInputActionOrigin_SwitchProGyroRoll                 EInputActionOrigin = 224
	EInputActionOrigin_SwitchDPadMove                    EInputActionOrigin = 225
	EInputActionOrigin_SwitchReserved1                   EInputActionOrigin = 226
	EInputActionOrigin_SwitchReserved2                   EInputActionOrigin = 227
	EInputActionOrigin_SwitchReserved3                   EInputActionOrigin = 228
	EInputActionOrigin_SwitchReserved4                   EInputActionOrigin = 229
	EInputActionOrigin_SwitchReserved5                   EInputActionOrigin = 230
	EInputActionOrigin_SwitchReserved6                   EInputActionOrigin = 231
	EInputActionOrigin_SwitchReserved7                   EInputActionOrigin = 232
	EInputActionOrigin_SwitchReserved8                   EInputActionOrigin = 233
	EInputActionOrigin_SwitchReserved9                   EInputActionOrigin = 234
	EInputActionOrigin_SwitchReserved10                  EInputActionOrigin = 235
	EInputActionOrigin_SwitchRightGyroMove               EInputActionOrigin = 236
	EInputActionOrigin_SwitchRightGyroPitch              EInputActionOrigin = 237
	EInputActionOrigin_SwitchRightGyroYaw                EInputActionOrigin = 238
	EInputActionOrigin_SwitchRightGyroRoll               EInputActionOrigin = 239
	EInputActionOrigin_SwitchLeftGyroMove                EInputActionOrigin = 240
	EInputActionOrigin_SwitchLeftGyroPitch               EInputActionOrigin = 241
	EInputActionOrigin_SwitchLeftGyroYaw                 EInputActionOrigin = 242
	EInputActionOrigin_SwitchLeftGyroRoll                EInputActionOrigin = 243
	EInputActionOrigin_SwitchLeftGripLower               EInputActionOrigin = 244
	EInputActionOrigin_SwitchLeftGripUpper               EInputActionOrigin = 245
	EInputActionOrigin_SwitchRightGripLower              EInputActionOrigin = 246
	EInputActionOrigin_SwitchRightGripUpper              EInputActionOrigin = 247
	EInputActionOrigin_SwitchJoyConButtonN               EInputActionOrigin = 248
	EInputActionOrigin_SwitchJoyConButtonE               EInputActionOrigin = 249
	EInputActionOrigin_SwitchJoyConButtonS               EInputActionOrigin = 250
	EInputActionOrigin_SwitchJoyConButtonW               EInputActionOrigin = 251
	EInputActionOrigin_SwitchReserved15                  EInputActionOrigin = 252
	EInputActionOrigin_SwitchReserved16                  EInputActionOrigin = 253
	EInputActionOrigin_SwitchReserved17                  EInputActionOrigin = 254
	EInputActionOrigin_SwitchReserved18                  EInputActionOrigin = 255
	EInputActionOrigin_SwitchReserved19                  EInputActionOrigin = 256
	EInputActionOrigin_SwitchReserved20                  EInputActionOrigin = 257
	EInputActionOrigin_PS5X                              EInputActionOrigin = 258
	EInputActionOrigin_PS5Circle                         EInputActionOrigin = 259
	EInputActionOrigin_PS5Triangle                       EInputActionOrigin = 260
	EInputActionOrigin_PS5Square                         EInputActionOrigin = 261
	EInputActionOrigin_PS5LeftBumper                     EInputActionOrigin = 262
	EInputActionOrigin_PS5RightBumper                    EInputActionOrigin = 263
	EInputActionOrigin_PS5Option                         EInputActionOrigin = 264
	EInputActionOrigin_PS5Create                         EInputActionOrigin = 265
	EInputActionOrigin_PS5Mute                           EInputActionOrigin = 266
	EInputActionOrigin_PS5LeftPadTouch                   EInputActionOrigin = 267
	EInputActionOrigin_PS5LeftPadSwipe                   EInputActionOrigin = 268
	EInputActionOrigin_PS5LeftPadClick                   EInputActionOrigin = 269
	EInputActionOrigin_PS5LeftPadDPadNorth               EInputActionOrigin = 270
	EInputActionOrigin_PS5LeftPadDPadSouth               EInputActionOrigin = 271
	EInputActionOrigin_PS5LeftPadDPadWest                EInputActionOrigin = 272
	EInputActionOrigin_PS5LeftPadDPadEast                EInputActionOrigin = 273
	EInputActionOrigin_PS5RightPadTouch                  EInputActionOrigin = 274
	EInputActionOrigin_PS5RightPadSwipe                  EInputActionOrigin = 275
	EInputActionOrigin_PS5RightPadClick                  EInputActionOrigin = 276
	EInputActionOrigin_PS5RightPadDPadNorth              EInputActionOrigin = 277
	EInputActionOrigin_PS5RightPadDPadSouth              EInputActionOrigin = 278
	EInputActionOrigin_PS5RightPadDPadWest               EInputActionOrigin = 279
	EInputActionOrigin_PS5RightPadDPadEast               EInputActionOrigin = 280
	EInputActionOrigin_PS5CenterPadTouch                 EInputActionOrigin = 281
	EInputActionOrigin_PS5CenterPadSwipe                 EInputActionOrigin = 282
	EInputActionOrigin_PS5CenterPadClick                 EInputActionOrigin = 283
	EInputActionOrigin_PS5CenterPadDPadNorth             EInputActionOrigin = 284
	EInputActionOrigin_PS5CenterPadDPadSouth             EInputActionOrigin = 285
	EInputActionOrigin_PS5CenterPadDPadWest              EInputActionOrigin = 286
	EInputActionOrigin_PS5CenterPadDPadEast              EInputActionOrigin = 287
	EInputActionOrigin_PS5LeftTriggerPull                EInputActionOrigin = 288
	EInputActionOrigin_PS5LeftTriggerClick               EInputActionOrigin = 289
	EInputActionOrigin_PS5RightTriggerPull               EInputActionOrigin = 290
	EInputActionOrigin_PS5RightTriggerClick              EInputActionOrigin = 291
	EInputActionOrigin_PS5LeftStickMove                  EInputActionOrigin = 292
	EInputActionOrigin_PS5LeftStickClick                 EInputActionOrigin = 293
	EInputActionOrigin_PS5LeftStickDPadNorth             EInputActionOrigin = 294
	EInputActionOrigin_PS5LeftStickDPadSouth             EInputActionOrigin = 295
	EInputActionOrigin_PS5LeftStickDPadWest              EInputActionOrigin = 296
	EInputActionOrigin_PS5LeftStickDPadEast              EInputActionOrigin = 297
	EInputActionOrigin_PS5RightStickMove                 EInputActionOrigin = 298
	EInputActionOrigin_PS5RightStickClick                EInputActionOrigin = 299
	EInputActionOrigin_PS5RightStickDPadNorth            EInputActionOrigin = 300
	EInputActionOrigin_PS5RightStickDPadSouth            EInputActionOrigin = 301
	EInputActionOrigin_PS5RightStickDPadWest             EInputActionOrigin = 302
	EInputActionOrigin_PS5RightStickDPadEast             EInputActionOrigin = 303
	EInputActionOrigin_PS5DPadNorth                      EInputActionOrigin = 304
	EInputActionOrigin_PS5DPadSouth                      EInputActionOrigin = 305
	EInputActionOrigin_PS5DPadWest                       EInputActionOrigin = 306
	EInputActionOrigin_PS5DPadEast                       EInputActionOrigin = 307
	EInputActionOrigin_PS5GyroMove                       EInputActionOrigin = 308
	EInputActionOrigin_PS5GyroPitch                      EInputActionOrigin = 309
	EInputActionOrigin_PS5GyroYaw                        EInputActionOrigin = 310
	EInputActionOrigin_PS5GyroRoll                       EInputActionOrigin = 311
	EInputActionOrigin_PS5DPadMove                       EInputActionOrigin = 312
	EInputActionOrigin_PS5LeftGrip                       EInputActionOrigin = 313
	EInputActionOrigin_PS5RightGrip                      EInputActionOrigin = 314
	EInputActionOrigin_PS5LeftFn                         EInputActionOrigin = 315
	EInputActionOrigin_PS5RightFn                        EInputActionOrigin = 316
	EInputActionOrigin_PS5Reserved5                      EInputActionOrigin = 317
	EInputActionOrigin_PS5Reserved6                      EInputActionOrigin = 318
	EInputActionOrigin_PS5Reserved7                      EInputActionOrigin = 319
	EInputActionOrigin_PS5Reserved8                      EInputActionOrigin = 320
	EInputActionOrigin_PS5Reserved9                      EInputActionOrigin = 321
	EInputActionOrigin_PS5Reserved10                     EInputActionOrigin = 322
	EInputActionOrigin_PS5Reserved11                     EInputActionOrigin = 323
	EInputActionOrigin_PS5Reserved12                     EInputActionOrigin = 324
	EInputActionOrigin_PS5Reserved13                     EInputActionOrigin = 325
	EInputActionOrigin_PS5Reserved14                     EInputActionOrigin = 326
	EInputActionOrigin_PS5Reserved15                     EInputActionOrigin = 327
	EInputActionOrigin_PS5Reserved16                     EInputActionOrigin = 328
	EInputActionOrigin_PS5Reserved17                     EInputActionOrigin = 329
	EInputActionOrigin_PS5Reserved18                     EInputActionOrigin = 330
	EInputActionOrigin_PS5Reserved19                     EInputActionOrigin = 331
	EInputActionOrigin_PS5Reserved20                     EInputActionOrigin = 332
	EInputActionOrigin_SteamDeckA                        EInputActionOrigin = 333
	EInputActionOrigin_SteamDeckB                        EInputActionOrigin = 334
	EInputActionOrigin_SteamDeckX                        EInputActionOrigin = 335
	EInputActionOrigin_SteamDeckY                        EInputActionOrigin = 336
	EInputActionOrigin_SteamDeckL1                       EInputActionOrigin = 337
	EInputActionOrigin_SteamDeckR1                       EInputActionOrigin = 338
	EInputActionOrigin_SteamDeckMenu                     EInputActionOrigin = 339
	EInputActionOrigin_SteamDeckView                     EInputActionOrigin = 340
	EInputActionOrigin_SteamDeckLeftPadTouch             EInputActionOrigin = 341
	EInputActionOrigin_SteamDeckLeftPadSwipe             EInputActionOrigin = 342
	EInputActionOrigin_SteamDeckLeftPadClick             EInputActionOrigin = 343
	EInputActionOrigin_SteamDeckLeftPadDPadNorth         EInputActionOrigin = 344
	EInputActionOrigin_SteamDeckLeftPadDPadSouth         EInputActionOrigin = 345
	EInputActionOrigin_SteamDeckLeftPadDPadWest          EInputActionOrigin = 346
	EInputActionOrigin_SteamDeckLeftPadDPadEast          EInputActionOrigin = 347
	EInputActionOrigin_SteamDeckRightPadTouch            EInputActionOrigin = 348
	EInputActionOrigin_SteamDeckRightPadSwipe            EInputActionOrigin = 349
	EInputActionOrigin_SteamDeckRightPadClick            EInputActionOrigin = 350
	EInputActionOrigin_SteamDeckRightPadDPadNorth        EInputActionOrigin = 351
	EInputActionOrigin_SteamDeckRightPadDPadSouth        EInputActionOrigin = 352
	EInputActionOrigin_SteamDeckRightPadDPadWest         EInputActionOrigin = 353
	EInputActionOrigin_SteamDeckRightPadDPadEast         EInputActionOrigin = 354
	EInputActionOrigin_SteamDeckL2SoftPull               EInputActionOrigin = 355
	EInputActionOrigin_SteamDeckL2                       EInputActionOrigin = 356
	EInputActionOrigin_SteamDeckR2SoftPull               EInputActionOrigin = 357
	EInputActionOrigin_SteamDeckR2                       EInputActionOrigin = 358
	EInputActionOrigin_SteamDeckLeftStickMove            EInputActionOrigin = 359
	EInputActionOrigin_SteamDeckL3                       EInputActionOrigin = 360
	EInputActionOrigin_SteamDeckLeftStickDPadNorth       EInputActionOrigin = 361
	EInputActionOrigin_SteamDeckLeftStickDPadSouth       EInputActionOrigin = 362
	EInputActionOrigin_SteamDeckLeftStickDPadWest        EInputActionOrigin = 363
	EInputActionOrigin_SteamDeckLeftStickDPadEast        EInputActionOrigin = 364
	EInputActionOrigin_SteamDeckLeftStickTouch           EInputActionOrigin = 365
	EInputActionOrigin_SteamDeckRightStickMove           EInputActionOrigin = 366
	EInputActionOrigin_SteamDeckR3                       EInputActionOrigin = 367
	EInputActionOrigin_SteamDeckRightStickDPadNorth      EInputActionOrigin = 368
	EInputActionOrigin_SteamDeckRightStickDPadSouth      EInputActionOrigin = 369
	EInputActionOrigin_SteamDeckRightStickDPadWest       EInputActionOrigin = 370
	EInputActionOrigin_SteamDeckRightStickDPadEast       EInputActionOrigin = 371
	EInputActionOrigin_SteamDeckRightStickTouch          EInputActionOrigin = 372
	EInputActionOrigin_SteamDeckL4                       EInputActionOrigin = 373
	EInputActionOrigin_SteamDeckR4                       EInputActionOrigin = 374
	EInputActionOrigin_SteamDeckL5                       EInputActionOrigin = 375
	EInputActionOrigin_SteamDeckR5                       EInputActionOrigin = 376
	EInputActionOrigin_SteamDeckDPadMove                 EInputActionOrigin = 377
	EInputActionOrigin_SteamDeckDPadNorth                EInputActionOrigin = 378
	EInputActionOrigin_SteamDeckDPadSouth                EInputActionOrigin = 379
	EInputActionOrigin_SteamDeckDPadWest                 EInputActionOrigin = 380
	EInputActionOrigin_SteamDeckDPadEast                 EInputActionOrigin = 381
	EInputActionOrigin_SteamDeckGyroMove                 EInputActionOrigin = 382
	EInputActionOrigin_SteamDeckGyroPitch                EInputActionOrigin = 383
	EInputActionOrigin_SteamDeckGyroYaw                  EInputActionOrigin = 384
	EInputActionOrigin_SteamDeckGyroRoll                 EInputActionOrigin = 385
	EInputActionOrigin_SteamDeckReserved1                EInputActionOrigin = 386
	EInputActionOrigin_SteamDeckReserved2                EInputActionOrigin = 387
	EInputActionOrigin_SteamDeckReserved3                EInputActionOrigin = 388
	EInputActionOrigin_SteamDeckReserved4                EInputActionOrigin = 389
	EInputActionOrigin_SteamDeckReserved5                EInputActionOrigin = 390
	EInputActionOrigin_SteamDeckReserved6                EInputActionOrigin = 391
	EInputActionOrigin_SteamDeckReserved7                EInputActionOrigin = 392
	EInputActionOrigin_SteamDeckReserved8                EInputActionOrigin = 393
	EInputActionOrigin_SteamDeckReserved9                EInputActionOrigin = 394
	EInputActionOrigin_SteamDeckReserved10               EInputActionOrigin = 395
	EInputActionOrigin_SteamDeckReserved11               EInputActionOrigin = 396
	EInputActionOrigin_SteamDeckReserved12               EInputActionOrigin = 397
	EInputActionOrigin_SteamDeckReserved13               EInputActionOrigin = 398
	EInputActionOrigin_SteamDeckReserved14               EInputActionOrigin = 399
	EInputActionOrigin_SteamDeckReserved15               EInputActionOrigin = 400
	EInputActionOrigin_SteamDeckReserved16               EInputActionOrigin = 401
	EInputActionOrigin_SteamDeckReserved17               EInputActionOrigin = 402
	EInputActionOrigin_SteamDeckReserved18               EInputActionOrigin = 403
	EInputActionOrigin_SteamDeckReserved19               EInputActionOrigin = 404
	EInputActionOrigin_SteamDeckReserved20               EInputActionOrigin = 405
	EInputActionOrigin_HoripadM1                         EInputActionOrigin = 406
	EInputActionOrigin_HoripadM2                         EInputActionOrigin = 407
	EInputActionOrigin_HoripadL4                         EInputActionOrigin = 408
	EInputActionOrigin_HoripadR4                         EInputActionOrigin = 409
	EInputActionOrigin_Count                             EInputActionOrigin = 410
	EInputActionOrigin_MaximumPossibleValue              EInputActionOrigin = 32767
)

type InputHandle_t uint64
type SteamInputActionEventCallbackPointer uintptr
type InputActionSetHandle uint64
type InputDigitalActionHandle uint64
type InputAnalogActionHandle uint64
type ESteamInputType int32

const (
	ESteamInputType_Unknown              ESteamInputType = 0
	ESteamInputType_SteamController      ESteamInputType = 1
	ESteamInputType_XBox360Controller    ESteamInputType = 2
	ESteamInputType_XBoxOneController    ESteamInputType = 3
	ESteamInputType_GenericGamepad       ESteamInputType = 4
	ESteamInputType_PS4Controller        ESteamInputType = 5
	ESteamInputType_AppleMFiController   ESteamInputType = 6
	ESteamInputType_AndroidController    ESteamInputType = 7
	ESteamInputType_SwitchJoyConPair     ESteamInputType = 8
	ESteamInputType_SwitchJoyConSingle   ESteamInputType = 9
	ESteamInputType_SwitchProController  ESteamInputType = 10
	ESteamInputType_MobileTouch          ESteamInputType = 11
	ESteamInputType_PS3Controller        ESteamInputType = 12
	ESteamInputType_PS5Controller        ESteamInputType = 13
	ESteamInputType_SteamDeckController  ESteamInputType = 14
	ESteamInputType_Count                ESteamInputType = 15
	ESteamInputType_MaximumPossibleValue ESteamInputType = 255
)

const (
	flatAPI_SteamInput                                       = "SteamAPI_SteamInput_v006"
	flatAPI_ISteamInput_Init                                 = "SteamAPI_ISteamInput_Init"
	flatAPI_ISteamInput_Shutdown                             = "SteamAPI_ISteamInput_Shutdown"
	flatAPI_ISteamInput_SetInputActionManifestFilePath       = "SteamAPI_ISteamInput_SetInputActionManifestFilePath"
	flatAPI_ISteamInput_RunFrame                             = "SteamAPI_ISteamInput_RunFrame"
	flatAPI_ISteamInput_BWaitForData                         = "SteamAPI_ISteamInput_BWaitForData"
	flatAPI_ISteamInput_BNewDataAvailable                    = "SteamAPI_ISteamInput_BNewDataAvailable"
	flatAPI_ISteamInput_GetConnectedControllers              = "SteamAPI_ISteamInput_GetConnectedControllers"
	flatAPI_ISteamInput_EnableDeviceCallbacks                = "SteamAPI_ISteamInput_EnableDeviceCallbacks"
	flatAPI_ISteamInput_EnableActionEventCallbacks           = "SteamAPI_ISteamInput_EnableActionEventCallbacks"
	flatAPI_ISteamInput_GetActionSetHandle                   = "SteamAPI_ISteamInput_GetActionSetHandle"
	flatAPI_ISteamInput_ActivateActionSet                    = "SteamAPI_ISteamInput_ActivateActionSet"
	flatAPI_ISteamInput_GetCurrentActionSet                  = "SteamAPI_ISteamInput_GetCurrentActionSet"
	flatAPI_ISteamInput_ActivateActionSetLayer               = "SteamAPI_ISteamInput_ActivateActionSetLayer"
	flatAPI_ISteamInput_DeactivateActionSetLayer             = "SteamAPI_ISteamInput_DeactivateActionSetLayer"
	flatAPI_ISteamInput_DeactivateAllActionSetLayers         = "SteamAPI_ISteamInput_DeactivateAllActionSetLayers"
	flatAPI_ISteamInput_GetActiveActionSetLayers             = "SteamAPI_ISteamInput_GetActiveActionSetLayers"
	flatAPI_ISteamInput_GetDigitalActionHandle               = "SteamAPI_ISteamInput_GetDigitalActionHandle"
	flatAPI_ISteamInput_GetDigitalActionData                 = "SteamAPI_ISteamInput_GetDigitalActionData"
	flatAPI_ISteamInput_GetDigitalActionOrigins              = "SteamAPI_ISteamInput_GetDigitalActionOrigins"
	flatAPI_ISteamInput_GetStringForDigitalActionName        = "SteamAPI_ISteamInput_GetStringForDigitalActionName"
	flatAPI_ISteamInput_GetAnalogActionHandle                = "SteamAPI_ISteamInput_GetAnalogActionHandle"
	flatAPI_ISteamInput_GetAnalogActionData                  = "SteamAPI_ISteamInput_GetAnalogActionData"
	flatAPI_ISteamInput_GetAnalogActionOrigins               = "SteamAPI_ISteamInput_GetAnalogActionOrigins"
	flatAPI_ISteamInput_GetGlyphPNGForActionOrigin           = "SteamAPI_ISteamInput_GetGlyphPNGForActionOrigin"
	flatAPI_ISteamInput_GetGlyphSVGForActionOrigin           = "SteamAPI_ISteamInput_GetGlyphSVGForActionOrigin"
	flatAPI_ISteamInput_GetGlyphForActionOrigin_Legacy       = "SteamAPI_ISteamInput_GetGlyphForActionOrigin_Legacy"
	flatAPI_ISteamInput_GetStringForActionOrigin             = "SteamAPI_ISteamInput_GetStringForActionOrigin"
	flatAPI_ISteamInput_GetStringForAnalogActionName         = "SteamAPI_ISteamInput_GetStringForAnalogActionName"
	flatAPI_ISteamInput_StopAnalogActionMomentum             = "SteamAPI_ISteamInput_StopAnalogActionMomentum"
	flatAPI_ISteamInput_GetMotionData                        = "SteamAPI_ISteamInput_GetMotionData"
	flatAPI_ISteamInput_TriggerVibration                     = "SteamAPI_ISteamInput_TriggerVibration"
	flatAPI_ISteamInput_TriggerVibrationExtended             = "SteamAPI_ISteamInput_TriggerVibrationExtended"
	flatAPI_ISteamInput_TriggerSimpleHapticEvent             = "SteamAPI_ISteamInput_TriggerSimpleHapticEvent"
	flatAPI_ISteamInput_SetLEDColor                          = "SteamAPI_ISteamInput_SetLEDColor"
	flatAPI_ISteamInput_Legacy_TriggerHapticPulse            = "SteamAPI_ISteamInput_Legacy_TriggerHapticPulse"
	flatAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse    = "SteamAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse"
	flatAPI_ISteamInput_ShowBindingPanel                     = "SteamAPI_ISteamInput_ShowBindingPanel"
	flatAPI_ISteamInput_GetInputTypeForHandle                = "SteamAPI_ISteamInput_GetInputTypeForHandle"
	flatAPI_ISteamInput_GetControllerForGamepadIndex         = "SteamAPI_ISteamInput_GetControllerForGamepadIndex"
	flatAPI_ISteamInput_GetGamepadIndexForController         = "SteamAPI_ISteamInput_GetGamepadIndexForController"
	flatAPI_ISteamInput_GetStringForXboxOrigin               = "SteamAPI_ISteamInput_GetStringForXboxOrigin"
	flatAPI_ISteamInput_GetGlyphForXboxOrigin                = "SteamAPI_ISteamInput_GetGlyphForXboxOrigin"
	flatAPI_ISteamInput_GetActionOriginFromXboxOrigin        = "SteamAPI_ISteamInput_GetActionOriginFromXboxOrigin"
	flatAPI_ISteamInput_TranslateActionOrigin                = "SteamAPI_ISteamInput_TranslateActionOrigin"
	flatAPI_ISteamInput_GetDeviceBindingRevision             = "SteamAPI_ISteamInput_GetDeviceBindingRevision"
	flatAPI_ISteamInput_GetRemotePlaySessionID               = "SteamAPI_ISteamInput_GetRemotePlaySessionID"
	flatAPI_ISteamInput_GetSessionInputConfigurationSettings = "SteamAPI_ISteamInput_GetSessionInputConfigurationSettings"
	flatAPI_ISteamInput_SetDualSenseTriggerEffect            = "SteamAPI_ISteamInput_SetDualSenseTriggerEffect"
)

var (
	steamInput_init                                  func() uintptr
	iSteamInput_Init                                 func(steamInput uintptr, explicitlyCallRunFrame bool) bool
	iSteamInput_Shutdown                             func(steamInput uintptr) bool
	iSteamInput_SetInputActionManifestFilePath       func(steamInput uintptr, inputActionManifestAbsolutePath string) bool
	iSteamInput_RunFrame                             func(steamInput uintptr, reservedValue bool)
	iSteamInput_BWaitForData                         func(steamInput uintptr, waitForever bool, timeout uint32) bool
	iSteamInput_BNewDataAvailable                    func(steamInput uintptr) bool
	iSteamInput_GetConnectedControllers              func(steamInput uintptr, handlesOut *[_STEAM_INPUT_MAX_COUNT]InputHandle_t) int32
	iSteamInput_EnableDeviceCallbacks                func(steamInput uintptr)
	iSteamInput_EnableActionEventCallbacks           func(steamInput uintptr, pCallback SteamInputActionEventCallbackPointer)
	iSteamInput_GetActionSetHandle                   func(steamInput uintptr, pszActionSetName string) InputActionSetHandle
	iSteamInput_ActivateActionSet                    func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle)
	iSteamInput_GetCurrentActionSet                  func(steamInput uintptr, inputHandle InputHandle_t) InputActionSetHandle
	iSteamInput_ActivateActionSetLayer               func(steamInput uintptr, inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle)
	iSteamInput_DeactivateActionSetLayer             func(steamInput uintptr, inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle)
	iSteamInput_DeactivateAllActionSetLayers         func(steamInput uintptr, inputHandle InputHandle_t)
	iSteamInput_GetActiveActionSetLayers             func(steamInput uintptr, inputHandle InputHandle_t, handlesOut *[_STEAM_INPUT_MAX_ACTIVE_LAYERS]InputActionSetHandle) int32
	iSteamInput_GetDigitalActionHandle               func(steamInput uintptr, actionName string) InputDigitalActionHandle
	iSteamInput_GetDigitalActionData                 func(steamInput uintptr, inputHandle InputHandle_t, digitalActionHandle InputDigitalActionHandle) uintptr
	iSteamInput_GetDigitalActionOrigins              func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle, originsOut *[_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32
	iSteamInput_GetStringForDigitalActionName        func(steamInput uintptr, actionHandle InputDigitalActionHandle) string
	iSteamInput_GetAnalogActionHandle                func(steamInput uintptr, actionName string) InputAnalogActionHandle
	iSteamInput_GetAnalogActionData                  func(steamInput uintptr, inputHandle InputHandle_t, analogActionHandle InputAnalogActionHandle) uintptr
	iSteamInput_GetAnalogActionOrigins               func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle, originsOut *[_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32
	iSteamInput_GetGlyphPNGForActionOrigin           func(steamInput uintptr, eOrigin EInputActionOrigin, size ESteamInputGlyphSize, flags uint32) []byte
	iSteamInput_GetGlyphSVGForActionOrigin           func(steamInput uintptr, eOrigin EInputActionOrigin, flags uint32) []byte
	iSteamInput_GetGlyphForActionOrigin_Legacy       func(steamInput uintptr, eOrigin EInputActionOrigin) []byte
	iSteamInput_GetStringForActionOrigin             func(steamInput uintptr, eOrigin EInputActionOrigin) string
	iSteamInput_GetStringForAnalogActionName         func(steamInput uintptr, actionHandle InputAnalogActionHandle) string
	iSteamInput_StopAnalogActionMomentum             func(steamInput uintptr, inputHandle InputHandle_t, action InputAnalogActionHandle)
	iSteamInput_GetMotionData                        func(steamInput uintptr, inputHandle InputHandle_t) uintptr
	iSteamInput_TriggerVibration                     func(steamInput uintptr, inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16)
	iSteamInput_TriggerVibrationExtended             func(steamInput uintptr, inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16)
	iSteamInput_TriggerSimpleHapticEvent             func(steamInput uintptr, inputHandle InputHandle_t, eHapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte)
	iSteamInput_SetLEDColor                          func(steamInput uintptr, inputHandle InputHandle_t, colorR uint8, colorG uint8, colorB uint8, flags uint)
	iSteamInput_Legacy_TriggerHapticPulse            func(steamInput uintptr, inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16)
	iSteamInput_Legacy_TriggerRepeatedHapticPulse    func(steamInput uintptr, inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint)
	iSteamInput_ShowBindingPanel                     func(steamInput uintptr, inputHandle InputHandle_t) bool
	iSteamInput_GetInputTypeForHandle                func(steamInput uintptr, inputHandle InputHandle_t) ESteamInputType
	iSteamInput_GetControllerForGamepadIndex         func(steamInput uintptr, index int32) InputHandle_t
	iSteamInput_GetGamepadIndexForController         func(steamInput uintptr, inputHandle InputHandle_t) int32
	iSteamInput_GetStringForXboxOrigin               func(steamInput uintptr, eOrigin EXboxOrigin) string
	iSteamInput_GetGlyphForXboxOrigin                func(steamInput uintptr, eOrigin EXboxOrigin) []byte
	iSteamInput_GetActionOriginFromXboxOrigin        func(steamInput uintptr, inputHandle InputHandle_t, eOrigin EXboxOrigin) EInputActionOrigin
	iSteamInput_TranslateActionOrigin                func(steamInput uintptr, destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin
	iSteamInput_GetDeviceBindingRevision             func(steamInput uintptr, inputHandle InputHandle_t, major *int32, minor *int32) bool
	iSteamInput_GetRemotePlaySessionID               func(steamInput uintptr, inputHandle InputHandle_t) uint32
	iSteamInput_GetSessionInputConfigurationSettings func(steamInput uintptr) uint16
	iSteamInput_SetDualSenseTriggerEffect            func(steamInput uintptr, inputHandle InputHandle_t, param *ScePadTriggerEffectParam)
)

type steamInput uintptr

func SteamInput() ISteamInput {
	return steamInput(steamInput_init())
}

func (s steamInput) Init(explicitlyCallRunFrame bool) bool {
	return iSteamInput_Init(uintptr(s), explicitlyCallRunFrame)
}

func (s steamInput) Shutdown() bool {
	return iSteamInput_Shutdown(uintptr(s))
}

func (s steamInput) SetInputActionManifestFilePath(inputActionManifestAbsolutePath string) bool {
	return iSteamInput_SetInputActionManifestFilePath(uintptr(s), inputActionManifestAbsolutePath)
}

func (s steamInput) RunFrame() {
	iSteamInput_RunFrame(uintptr(s), false)
}

func (s steamInput) RunFrameEx(reservedValue bool) {
	iSteamInput_RunFrame(uintptr(s), reservedValue)
}

func (s steamInput) BWaitForData(waitForever bool, timeout uint32) bool {
	return iSteamInput_BWaitForData(uintptr(s), waitForever, timeout)
}

func (s steamInput) BNewDataAvailable() bool {
	return iSteamInput_BNewDataAvailable(uintptr(s))
}

func (s steamInput) GetConnectedControllers() []InputHandle_t {
	var handlesOut [_STEAM_INPUT_MAX_COUNT]InputHandle_t
	result := iSteamInput_GetConnectedControllers(uintptr(s), &handlesOut)
	return handlesOut[:result]
}

func (s steamInput) EnableDeviceCallbacks() {
	iSteamInput_EnableDeviceCallbacks(uintptr(s))
}

func (s steamInput) EnableActionEventCallbacks(pCallback SteamInputActionEventCallbackPointer) {
	iSteamInput_EnableActionEventCallbacks(uintptr(s), pCallback)
}

func (s steamInput) GetActionSetHandle(pszActionSetName string) InputActionSetHandle {
	return iSteamInput_GetActionSetHandle(uintptr(s), pszActionSetName)
}

func (s steamInput) ActivateActionSet(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle) {
	iSteamInput_ActivateActionSet(uintptr(s), inputHandle, actionSetHandle)
}

func (s steamInput) GetCurrentActionSet(inputHandle InputHandle_t) InputActionSetHandle {
	return iSteamInput_GetCurrentActionSet(uintptr(s), inputHandle)
}

func (s steamInput) ActivateActionSetLayer(inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle) {
	iSteamInput_ActivateActionSetLayer(uintptr(s), inputHandle, actionSetLayerHandle)
}

func (s steamInput) DeactivateActionSetLayer(inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle) {
	iSteamInput_DeactivateActionSetLayer(uintptr(s), inputHandle, actionSetLayerHandle)
}

func (s steamInput) DeactivateAllActionSetLayers(inputHandle InputHandle_t) {
	iSteamInput_DeactivateAllActionSetLayers(uintptr(s), inputHandle)
}

func (s steamInput) GetActiveActionSetLayers(inputHandle InputHandle_t) []InputActionSetHandle {
	var handlesOut [_STEAM_INPUT_MAX_ACTIVE_LAYERS]InputActionSetHandle
	result := iSteamInput_GetActiveActionSetLayers(uintptr(s), inputHandle, &handlesOut)
	return handlesOut[:result]
}

func (s steamInput) GetDigitalActionHandle(actionName string) InputDigitalActionHandle {
	return iSteamInput_GetDigitalActionHandle(uintptr(s), actionName)
}

func (s steamInput) GetDigitalActionData(inputHandle InputHandle_t, digitalActionHandle InputDigitalActionHandle) InputDigitalActionData {
	return *uintptrToStruct[InputDigitalActionData](iSteamInput_GetDigitalActionData(uintptr(s), inputHandle, digitalActionHandle))
}

func (s steamInput) GetDigitalActionOrigins(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle) []EInputActionOrigin {
	var originsOut [_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin
	result := iSteamInput_GetDigitalActionOrigins(uintptr(s), inputHandle, actionSetHandle, digitalActionHandle, &originsOut)
	return originsOut[:result]
}

func (s steamInput) GetStringForDigitalActionName(actionHandle InputDigitalActionHandle) string {
	return iSteamInput_GetStringForDigitalActionName(uintptr(s), actionHandle)
}

func (s steamInput) GetAnalogActionHandle(actionName string) InputAnalogActionHandle {
	return iSteamInput_GetAnalogActionHandle(uintptr(s), actionName)
}

func (s steamInput) GetAnalogActionData(inputHandle InputHandle_t, analogActionHandle InputAnalogActionHandle) InputAnalogActionData {
	return *uintptrToStruct[InputAnalogActionData](iSteamInput_GetAnalogActionData(uintptr(s), inputHandle, analogActionHandle))
}

func (s steamInput) GetAnalogActionOrigins(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle) []EInputActionOrigin {
	var originsOut [_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin
	result := iSteamInput_GetAnalogActionOrigins(uintptr(s), inputHandle, actionSetHandle, analogActionHandle, &originsOut)
	return originsOut[:result]
}

func (s steamInput) GetGlyphPNGForActionOrigin(origin EInputActionOrigin, size ESteamInputGlyphSize, flags uint32) []byte {
	return iSteamInput_GetGlyphPNGForActionOrigin(uintptr(s), origin, size, flags)
}

func (s steamInput) GetGlyphSVGForActionOrigin(origin EInputActionOrigin, flags uint32) []byte {
	return iSteamInput_GetGlyphSVGForActionOrigin(uintptr(s), origin, flags)
}

func (s steamInput) GetGlyphForActionOrigin_Legacy(origin EInputActionOrigin) []byte {
	return iSteamInput_GetGlyphForActionOrigin_Legacy(uintptr(s), origin)
}

func (s steamInput) GetStringForActionOrigin(origin EInputActionOrigin) string {
	return iSteamInput_GetStringForActionOrigin(uintptr(s), origin)
}

func (s steamInput) GetStringForAnalogActionName(actionHandle InputAnalogActionHandle) string {
	return iSteamInput_GetStringForAnalogActionName(uintptr(s), actionHandle)
}

func (s steamInput) StopAnalogActionMomentum(inputHandle InputHandle_t, action InputAnalogActionHandle) {
	iSteamInput_StopAnalogActionMomentum(uintptr(s), inputHandle, action)
}

func (s steamInput) GetMotionData(inputHandle InputHandle_t) InputMotionData {
	return *uintptrToStruct[InputMotionData](iSteamInput_GetMotionData(uintptr(s), inputHandle))
}

func (s steamInput) TriggerVibration(inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16) {
	iSteamInput_TriggerVibration(uintptr(s), inputHandle, leftSpeed, rightSpeed)
}

func (s steamInput) TriggerVibrationExtended(inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16) {
	iSteamInput_TriggerVibrationExtended(uintptr(s), inputHandle, leftSpeed, rightSpeed, leftTriggerSpeed, rightTriggerSpeed)
}

func (s steamInput) TriggerSimpleHapticEvent(inputHandle InputHandle_t, hapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte) {
	iSteamInput_TriggerSimpleHapticEvent(uintptr(s), inputHandle, hapticLocation, intensity, gainDB, otherIntensity, otherGainDB)
}

func (s steamInput) SetLEDColor(inputHandle InputHandle_t, colorR uint8, colorG uint8, colorB uint8, flags uint) {
	iSteamInput_SetLEDColor(uintptr(s), inputHandle, colorR, colorG, colorB, flags)
}

func (s steamInput) Legacy_TriggerHapticPulse(inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16) {
	iSteamInput_Legacy_TriggerHapticPulse(uintptr(s), inputHandle, targetPad, durationMicroSec)
}

func (s steamInput) Legacy_TriggerRepeatedHapticPulse(inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint) {
	iSteamInput_Legacy_TriggerRepeatedHapticPulse(uintptr(s), inputHandle, targetPad, durationMicroSec, offMicroSec, repeat, flags)
}

func (s steamInput) ShowBindingPanel(inputHandle InputHandle_t) bool {
	return iSteamInput_ShowBindingPanel(uintptr(s), inputHandle)
}

func (s steamInput) GetInputTypeForHandle(inputHandle InputHandle_t) ESteamInputType {
	return iSteamInput_GetInputTypeForHandle(uintptr(s), inputHandle)
}

func (s steamInput) GetControllerForGamepadIndex(index int32) InputHandle_t {
	return iSteamInput_GetControllerForGamepadIndex(uintptr(s), index)
}

func (s steamInput) GetGamepadIndexForController(inputHandle InputHandle_t) int32 {
	return iSteamInput_GetGamepadIndexForController(uintptr(s), inputHandle)
}

func (s steamInput) GetStringForXboxOrigin(origin EXboxOrigin) string {
	return iSteamInput_GetStringForXboxOrigin(uintptr(s), origin)
}

func (s steamInput) GetGlyphForXboxOrigin(origin EXboxOrigin) []byte {
	return iSteamInput_GetGlyphForXboxOrigin(uintptr(s), origin)
}

func (s steamInput) GetActionOriginFromXboxOrigin(inputHandle InputHandle_t, origin EXboxOrigin) EInputActionOrigin {
	return iSteamInput_GetActionOriginFromXboxOrigin(uintptr(s), inputHandle, origin)
}

func (s steamInput) TranslateActionOrigin(destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin {
	return iSteamInput_TranslateActionOrigin(uintptr(s), destinationInputType, sourceOrigin)
}

func (s steamInput) GetDeviceBindingRevision(inputHandle InputHandle_t) (major int32, minor int32, foundController bool) {
	foundController = iSteamInput_GetDeviceBindingRevision(uintptr(s), inputHandle, &major, &minor)
	return major, minor, foundController
}

func (s steamInput) GetRemotePlaySessionID(inputHandle InputHandle_t) uint32 {
	return iSteamInput_GetRemotePlaySessionID(uintptr(s), inputHandle)
}

func (s steamInput) GetSessionInputConfigurationSettings() uint16 {
	return iSteamInput_GetSessionInputConfigurationSettings(uintptr(s))
}

func (s steamInput) SetDualSenseTriggerEffect(inputHandle InputHandle_t, param *ScePadTriggerEffectParam) {
	iSteamInput_SetDualSenseTriggerEffect(uintptr(s), inputHandle, param)
}

// Steam Inventory
const (
	SteamItemInstanceIDInvalid        SteamItemInstanceID        = ^SteamItemInstanceID(0)
	SteamInventoryResultInvalid       SteamInventoryResult       = -1
	SteamInventoryUpdateHandleInvalid SteamInventoryUpdateHandle = 0xffffffffffffffff
)

type SteamItemDef int32
type SteamInventoryUpdateHandle uint64
type SteamItemDetails struct {
	ItemId     SteamItemInstanceID
	Definition SteamItemDef
	Quantity   uint16
	Flags      uint16
}
type SteamItemInstanceID uint64

type SteamInventoryResult int32

type ESteamItemFlags int32

const (
	ESteamItem_NoTrade  ESteamItemFlags = 1
	ESteamItem_Removed  ESteamItemFlags = 256
	ESteamItem_Consumed ESteamItemFlags = 512
)

const (
	flatAPI_SteamInventory                                         = "SteamAPI_SteamInventory_v003"
	flatAPI_ISteamInventory_GetResultStatus                        = "SteamAPI_ISteamInventory_GetResultStatus"
	flatAPI_ISteamInventory_GetResultItems                         = "SteamAPI_ISteamInventory_GetResultItems"
	flatAPI_ISteamInventory_GetResultItemProperty                  = "SteamAPI_ISteamInventory_GetResultItemProperty"
	flatAPI_ISteamInventory_GetResultTimestamp                     = "SteamAPI_ISteamInventory_GetResultTimestamp"
	flatAPI_ISteamInventory_CheckResultSteamID                     = "SteamAPI_ISteamInventory_CheckResultSteamID"
	flatAPI_ISteamInventory_DestroyResult                          = "SteamAPI_ISteamInventory_DestroyResult"
	flatAPI_ISteamInventory_GetAllItems                            = "SteamAPI_ISteamInventory_GetAllItems"
	flatAPI_ISteamInventory_GetItemsByID                           = "SteamAPI_ISteamInventory_GetItemsByID"
	flatAPI_ISteamInventory_SerializeResult                        = "SteamAPI_ISteamInventory_SerializeResult"
	flatAPI_ISteamInventory_DeserializeResult                      = "SteamAPI_ISteamInventory_DeserializeResult"
	flatAPI_ISteamInventory_GenerateItems                          = "SteamAPI_ISteamInventory_GenerateItems"
	flatAPI_ISteamInventory_GrantPromoItems                        = "SteamAPI_ISteamInventory_GrantPromoItems"
	flatAPI_ISteamInventory_AddPromoItem                           = "SteamAPI_ISteamInventory_AddPromoItem"
	flatAPI_ISteamInventory_AddPromoItems                          = "SteamAPI_ISteamInventory_AddPromoItems"
	flatAPI_ISteamInventory_ConsumeItem                            = "SteamAPI_ISteamInventory_ConsumeItem"
	flatAPI_ISteamInventory_ExchangeItems                          = "SteamAPI_ISteamInventory_ExchangeItems"
	flatAPI_ISteamInventory_TransferItemQuantity                   = "SteamAPI_ISteamInventory_TransferItemQuantity"
	flatAPI_ISteamInventory_TriggerItemDrop                        = "SteamAPI_ISteamInventory_TriggerItemDrop"
	flatAPI_ISteamInventory_LoadItemDefinitions                    = "SteamAPI_ISteamInventory_LoadItemDefinitions"
	flatAPI_ISteamInventory_GetItemDefinitionIDs                   = "SteamAPI_ISteamInventory_GetItemDefinitionIDs"
	flatAPI_ISteamInventory_GetItemDefinitionProperty              = "SteamAPI_ISteamInventory_GetItemDefinitionProperty"
	flatAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs = "SteamAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs"
	flatAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs      = "SteamAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs"
	flatAPI_ISteamInventory_StartPurchase                          = "SteamAPI_ISteamInventory_StartPurchase"
	flatAPI_ISteamInventory_RequestPrices                          = "SteamAPI_ISteamInventory_RequestPrices"
	flatAPI_ISteamInventory_GetNumItemsWithPrices                  = "SteamAPI_ISteamInventory_GetNumItemsWithPrices"
	flatAPI_ISteamInventory_GetItemsWithPrices                     = "SteamAPI_ISteamInventory_GetItemsWithPrices"
	flatAPI_ISteamInventory_GetItemPrice                           = "SteamAPI_ISteamInventory_GetItemPrice"
	flatAPI_ISteamInventory_StartUpdateProperties                  = "SteamAPI_ISteamInventory_StartUpdateProperties"
	flatAPI_ISteamInventory_RemoveProperty                         = "SteamAPI_ISteamInventory_RemoveProperty"
	flatAPI_ISteamInventory_SetPropertyString                      = "SteamAPI_ISteamInventory_SetPropertyString"
	flatAPI_ISteamInventory_SetPropertyBool                        = "SteamAPI_ISteamInventory_SetPropertyBool"
	flatAPI_ISteamInventory_SetPropertyInt64                       = "SteamAPI_ISteamInventory_SetPropertyInt64"
	flatAPI_ISteamInventory_SetPropertyFloat                       = "SteamAPI_ISteamInventory_SetPropertyFloat"
	flatAPI_ISteamInventory_SubmitUpdateProperties                 = "SteamAPI_ISteamInventory_SubmitUpdateProperties"
	flatAPI_ISteamInventory_InspectItem                            = "SteamAPI_ISteamInventory_InspectItem"
)

var (
	steamInventory_init                                    func() uintptr
	iSteamInventory_GetResultStatus                        func(steamInventory uintptr, resultHandle SteamInventoryResult) EResult
	iSteamInventory_GetResultItems                         func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutItemsArray []SteamItemDetails, punOutItemsArraySize *uint32) bool
	iSteamInventory_GetResultItemProperty                  func(steamInventory uintptr, resultHandle SteamInventoryResult, unItemIndex uint32, pchPropertyName string, pchValueBuffer []byte, punValueBufferSizeOut *uint32) bool
	iSteamInventory_GetResultTimestamp                     func(steamInventory uintptr, resultHandle SteamInventoryResult) uint32
	iSteamInventory_CheckResultSteamID                     func(steamInventory uintptr, resultHandle SteamInventoryResult, steamIDExpected Uint64SteamID) bool
	iSteamInventory_DestroyResult                          func(steamInventory uintptr, resultHandle SteamInventoryResult)
	iSteamInventory_GetAllItems                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_GetItemsByID                           func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pInstanceIDs []SteamItemInstanceID, unCountInstanceIDs uint32) bool
	iSteamInventory_SerializeResult                        func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutBuffer []byte, punOutBufferSize *uint32) bool
	iSteamInventory_DeserializeResult                      func(steamInventory uintptr, pOutResultHandle *SteamInventoryResult, pBuffer []byte, unBufferSize uint32, bRESERVEDMUSTBEFALSE bool) bool
	iSteamInventory_GenerateItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) bool
	iSteamInventory_GrantPromoItems                        func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_AddPromoItem                           func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemDef SteamItemDef) bool
	iSteamInventory_AddPromoItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, unArrayLength uint32) bool
	iSteamInventory_ConsumeItem                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemConsume SteamItemInstanceID, unQuantity uint32) bool
	iSteamInventory_ExchangeItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayGenerate []SteamItemDef, punArrayGenerateQuantity []uint32, unArrayGenerateLength uint32, pArrayDestroy []SteamItemInstanceID, punArrayDestroyQuantity []uint32, unArrayDestroyLength uint32) bool
	iSteamInventory_TransferItemQuantity                   func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemIdSource SteamItemInstanceID, unQuantity uint32, itemIdDest SteamItemInstanceID) bool
	iSteamInventory_TriggerItemDrop                        func(steamInventory uintptr, pResultHandle *SteamInventoryResult, dropListDefinition SteamItemDef) bool
	iSteamInventory_LoadItemDefinitions                    func(steamInventory uintptr) bool
	iSteamInventory_GetItemDefinitionIDs                   func(steamInventory uintptr, pItemDefIDs []SteamItemDef, punItemDefIDsArraySize *uint32) bool
	iSteamInventory_GetItemDefinitionProperty              func(steamInventory uintptr, iDefinition SteamItemDef, pchPropertyName string, pchValueBuffer []byte, punValueBufferSizeOut *uint32) bool
	iSteamInventory_RequestEligiblePromoItemDefinitionsIDs func(steamInventory uintptr, steamID Uint64SteamID) SteamAPICall
	iSteamInventory_GetEligiblePromoItemDefinitionIDs      func(steamInventory uintptr, steamID Uint64SteamID, pItemDefIDs []SteamItemDef, punItemDefIDsArraySize *uint32) bool
	iSteamInventory_StartPurchase                          func(steamInventory uintptr, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) SteamAPICall
	iSteamInventory_RequestPrices                          func(steamInventory uintptr) SteamAPICall
	iSteamInventory_GetNumItemsWithPrices                  func(steamInventory uintptr) uint32
	iSteamInventory_GetItemsWithPrices                     func(steamInventory uintptr, pArrayItemDefs []SteamItemDef, pCurrentPrices []uint64, pBasePrices []uint64, unArrayLength uint32) bool
	iSteamInventory_GetItemPrice                           func(steamInventory uintptr, iDefinition SteamItemDef, pCurrentPrice *uint64, pBasePrice *uint64) bool
	iSteamInventory_StartUpdateProperties                  func(steamInventory uintptr) SteamInventoryUpdateHandle
	iSteamInventory_RemoveProperty                         func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string) bool
	iSteamInventory_SetPropertyString                      func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, pchPropertyValue string) bool
	iSteamInventory_SetPropertyBool                        func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, bValue bool) bool
	iSteamInventory_SetPropertyInt64                       func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, nValue int64) bool
	iSteamInventory_SetPropertyFloat                       func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, flValue float32) bool
	iSteamInventory_SubmitUpdateProperties                 func(steamInventory uintptr, handle SteamInventoryUpdateHandle, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_InspectItem                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pchItemToken string) bool
)

type steamInventory uintptr

func SteamInventory() ISteamInventory {
	return steamInventory(steamInventory_init())
}

func (s steamInventory) GetResultStatus(resultHandle SteamInventoryResult) EResult {
	return iSteamInventory_GetResultStatus(uintptr(s), resultHandle)
}

func (s steamInventory) GetResultItems(resultHandle SteamInventoryResult, outItemsArraySize uint32) ([]SteamItemDetails, bool) {
	var pOutItemsArray []SteamItemDetails = make([]SteamItemDetails, outItemsArraySize)
	result := iSteamInventory_GetResultItems(uintptr(s), resultHandle, pOutItemsArray, &outItemsArraySize)
	return pOutItemsArray[:outItemsArraySize], result
}

func (s steamInventory) GetResultItemProperty(resultHandle SteamInventoryResult, itemIndex uint32, propertyName string, valueBufferSizeOut uint32) ([]byte, bool) {
	var pchValueBuffer []byte = make([]byte, 0, valueBufferSizeOut)
	result := iSteamInventory_GetResultItemProperty(uintptr(s), resultHandle, itemIndex, propertyName, pchValueBuffer, &valueBufferSizeOut)
	return pchValueBuffer[:valueBufferSizeOut], result
}

func (s steamInventory) GetResultTimestamp(resultHandle SteamInventoryResult) uint32 {
	return iSteamInventory_GetResultTimestamp(uintptr(s), resultHandle)
}

func (s steamInventory) CheckResultSteamID(resultHandle SteamInventoryResult, expectedSteamID Uint64SteamID) bool {
	return iSteamInventory_CheckResultSteamID(uintptr(s), resultHandle, expectedSteamID)
}

func (s steamInventory) DestroyResult(resultHandle SteamInventoryResult) {
	iSteamInventory_DestroyResult(uintptr(s), resultHandle)
}

func (s steamInventory) GetAllItems() (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GetAllItems(uintptr(s), &resultHandle)
	return resultHandle, success
}

func (s steamInventory) GetItemsByID(countInstanceIDs uint32) (resultHandle SteamInventoryResult, instanceIDs []SteamItemInstanceID, result bool) {
	instanceIDs = make([]SteamItemInstanceID, 0, countInstanceIDs)
	result = iSteamInventory_GetItemsByID(uintptr(s), &resultHandle, instanceIDs, countInstanceIDs)
	return resultHandle, instanceIDs, result
}

func (s steamInventory) SerializeResult(resultHandle SteamInventoryResult, outBufferSize uint32) ([]byte, bool) {
	var pOutBuffer []byte = make([]byte, 0, outBufferSize)
	result := iSteamInventory_SerializeResult(uintptr(s), resultHandle, pOutBuffer, &outBufferSize)
	return pOutBuffer[:outBufferSize], result
}

func (s steamInventory) DeserializeResult(bufferSize uint32, bRESERVEDMUSTBEFALSE bool) (resultHandle SteamInventoryResult, buffer []byte, success bool) {
	var pBuffer []byte = make([]byte, 0, bufferSize)
	success = iSteamInventory_DeserializeResult(uintptr(s), &resultHandle, pBuffer, bufferSize, bRESERVEDMUSTBEFALSE)
	return resultHandle, pBuffer, success
}

func (s steamInventory) GenerateItems(arrayItemDefs []SteamItemDef, arrayQuantity []uint32, arrayLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GenerateItems(uintptr(s), &resultHandle, arrayItemDefs, arrayQuantity, arrayLength)
	return resultHandle, success
}

func (s steamInventory) GrantPromoItems() (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GrantPromoItems(uintptr(s), &resultHandle)
	return resultHandle, success
}

func (s steamInventory) AddPromoItem(itemDef SteamItemDef) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_AddPromoItem(uintptr(s), &resultHandle, itemDef)
	return resultHandle, success
}

func (s steamInventory) AddPromoItems(arrayItemDefs []SteamItemDef, arrayLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_AddPromoItems(uintptr(s), &resultHandle, arrayItemDefs, arrayLength)
	return resultHandle, success
}

func (s steamInventory) ConsumeItem(itemConsume SteamItemInstanceID, quantity uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_ConsumeItem(uintptr(s), &resultHandle, itemConsume, quantity)
	return resultHandle, success
}

func (s steamInventory) ExchangeItems(arrayGenerate []SteamItemDef, arrayGenerateQuantity []uint32, arrayGenerateLength uint32, arrayDestroy []SteamItemInstanceID, arrayDestroyQuantity []uint32, arrayDestroyLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_ExchangeItems(uintptr(s), &resultHandle, arrayGenerate, arrayGenerateQuantity, arrayGenerateLength, arrayDestroy, arrayDestroyQuantity, arrayDestroyLength)
	return resultHandle, success
}

func (s steamInventory) TransferItemQuantity(itemIdSource SteamItemInstanceID, quantity uint32, itemIdDest SteamItemInstanceID) (pResultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_TransferItemQuantity(uintptr(s), &pResultHandle, itemIdSource, quantity, itemIdDest)
	return pResultHandle, success
}

func (s steamInventory) TriggerItemDrop(dropListDefinition SteamItemDef) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_TriggerItemDrop(uintptr(s), &resultHandle, dropListDefinition)
	return resultHandle, success
}

func (s steamInventory) LoadItemDefinitions() bool {
	return iSteamInventory_LoadItemDefinitions(uintptr(s))
}

func (s steamInventory) GetItemDefinitionIDs(itemDefIDsArraySize uint32) ([]SteamItemDef, bool) {
	var pItemDefIDs []SteamItemDef = make([]SteamItemDef, itemDefIDsArraySize)
	result := iSteamInventory_GetItemDefinitionIDs(uintptr(s), pItemDefIDs, &itemDefIDsArraySize)
	return pItemDefIDs[:itemDefIDsArraySize], result
}

func (s steamInventory) GetItemDefinitionProperty(definition SteamItemDef, propertyName string, valueBufferSizeOut uint32) ([]byte, bool) {
	var pchValueBuffer []byte = make([]byte, 0, valueBufferSizeOut)
	result := iSteamInventory_GetItemDefinitionProperty(uintptr(s), definition, propertyName, pchValueBuffer, &valueBufferSizeOut)
	return pchValueBuffer, result
}

func (s steamInventory) RequestEligiblePromoItemDefinitionsIDs(steamID Uint64SteamID) CallResult[SteamInventoryEligiblePromoItemDefIDs] {
	handle := iSteamInventory_RequestEligiblePromoItemDefinitionsIDs(uintptr(s), steamID)
	return CallResult[SteamInventoryEligiblePromoItemDefIDs]{Handle: handle}
}

func (s steamInventory) GetEligiblePromoItemDefinitionIDs(steamID Uint64SteamID, itemDefIDsArraySize uint32) ([]SteamItemDef, bool) {
	var pItemDefIDs []SteamItemDef = make([]SteamItemDef, itemDefIDsArraySize)
	result := iSteamInventory_GetEligiblePromoItemDefinitionIDs(uintptr(s), steamID, pItemDefIDs, &itemDefIDsArraySize)
	return pItemDefIDs[:itemDefIDsArraySize], result
}

func (s steamInventory) StartPurchase(arrayItemDefs []SteamItemDef, punArrayQuantity []uint32, arrayLength uint32) CallResult[SteamInventoryStartPurchaseResult] {
	handle := iSteamInventory_StartPurchase(uintptr(s), arrayItemDefs, punArrayQuantity, arrayLength)
	return CallResult[SteamInventoryStartPurchaseResult]{Handle: handle}
}

func (s steamInventory) RequestPrices() CallResult[SteamInventoryRequestPricesResult] {
	handle := iSteamInventory_RequestPrices(uintptr(s))
	return CallResult[SteamInventoryRequestPricesResult]{Handle: handle}
}

func (s steamInventory) GetNumItemsWithPrices() uint32 {
	return iSteamInventory_GetNumItemsWithPrices(uintptr(s))
}

func (s steamInventory) GetItemsWithPrices(arrayLength uint32) ([]SteamItemDef, []uint64, []uint64, bool) {
	var pArrayItemDefs []SteamItemDef = make([]SteamItemDef, arrayLength)
	var pCurrentPrices []uint64 = make([]uint64, arrayLength)
	var pBasePrices []uint64 = make([]uint64, arrayLength)
	result := iSteamInventory_GetItemsWithPrices(uintptr(s), pArrayItemDefs, pCurrentPrices, pBasePrices, arrayLength)
	return pArrayItemDefs, pCurrentPrices, pBasePrices, result
}

func (s steamInventory) GetItemPrice(definition SteamItemDef) (currentPrice uint64, basePrice uint64, succes bool) {
	succes = iSteamInventory_GetItemPrice(uintptr(s), definition, &currentPrice, &basePrice)
	return currentPrice, basePrice, succes
}

func (s steamInventory) StartUpdateProperties() SteamInventoryUpdateHandle {
	return iSteamInventory_StartUpdateProperties(uintptr(s))
}

func (s steamInventory) RemoveProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string) bool {
	return iSteamInventory_RemoveProperty(uintptr(s), handle, itemID, propertyName)
}

func (s steamInventory) SetProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, pchPropertyValue string) bool {
	return iSteamInventory_SetPropertyString(uintptr(s), handle, itemID, propertyName, pchPropertyValue)
}

func (s steamInventory) SetPropertyBool(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value bool) bool {
	return iSteamInventory_SetPropertyBool(uintptr(s), handle, itemID, propertyName, value)
}

func (s steamInventory) SetPropertyInt64y(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value int64) bool {
	return iSteamInventory_SetPropertyInt64(uintptr(s), handle, itemID, propertyName, value)
}

func (s steamInventory) SetPropertyFloat(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value float32) bool {
	return iSteamInventory_SetPropertyFloat(uintptr(s), handle, itemID, propertyName, value)
}

func (s steamInventory) SubmitUpdateProperties(handle SteamInventoryUpdateHandle) (resultHandle SteamInventoryResult, succes bool) {
	succes = iSteamInventory_SubmitUpdateProperties(uintptr(s), handle, &resultHandle)
	return resultHandle, succes
}

func (s steamInventory) InspectItem(itemToken string) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_InspectItem(uintptr(s), &resultHandle, itemToken)
	return resultHandle, success
}

// Steam Matchmaking
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
	EChatEntryType_Invalid          EChatEntryType = 0
	EChatEntryType_ChatMsg          EChatEntryType = 1
	EChatEntryType_Typing           EChatEntryType = 2
	EChatEntryType_InviteGame       EChatEntryType = 3
	EChatEntryType_Emote            EChatEntryType = 4
	EChatEntryType_LeftConversation EChatEntryType = 6
	EChatEntryType_Entered          EChatEntryType = 7
	EChatEntryType_WasKicked        EChatEntryType = 8
	EChatEntryType_WasBanned        EChatEntryType = 9
	EChatEntryType_Disconnected     EChatEntryType = 10
	EChatEntryType_HistoricalChat   EChatEntryType = 11
	EChatEntryType_LinkBlocked      EChatEntryType = 14
)

type EMatchMakingServerResponse int32

const (
	EServerResponded               EMatchMakingServerResponse = 0
	EServerFailedToRespond         EMatchMakingServerResponse = 1
	ENoServersListedOnMasterServer EMatchMakingServerResponse = 2
)

type ELobbyType int32

const (
	ELobbyType_Private       ELobbyType = 0
	ELobbyType_FriendsOnly   ELobbyType = 1
	ELobbyType_Public        ELobbyType = 2
	ELobbyType_Invisible     ELobbyType = 3
	ELobbyType_PrivateUnique ELobbyType = 4
)

type ELobbyComparison int32

const (
	ELobbyComparisonE_EqualToOrLessThan    ELobbyComparison = -2
	ELobbyComparisonE_LessThan             ELobbyComparison = -1
	ELobbyComparisonE_Equal                ELobbyComparison = 0
	ELobbyComparisonE_GreaterThan          ELobbyComparison = 1
	ELobbyComparisonE_EqualToOrGreaterThan ELobbyComparison = 2
	ELobbyComparisonE_NotEqual             ELobbyComparison = 3
)

type ELobbyDistanceFilter int32

const (
	ELobbyDistanceFilter_Close     ELobbyDistanceFilter = 0
	ELobbyDistanceFilter_Default   ELobbyDistanceFilter = 1
	ELobbyDistanceFilter_Far       ELobbyDistanceFilter = 2
	ELobbyDistanceFilter_Worldwide ELobbyDistanceFilter = 3
)

type EChatMemberStateChange int32

const (
	EChatMemberStateChange_Entered      EChatMemberStateChange = 1
	EChatMemberStateChange_Left         EChatMemberStateChange = 2
	EChatMemberStateChange_Disconnected EChatMemberStateChange = 4
	EChatMemberStateChange_Kicked       EChatMemberStateChange = 8
	EChatMemberStateChange_Banned       EChatMemberStateChange = 16
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
	iSteamMatchmaking_GetFavoriteGame                            func(steamMatchmaking uintptr, iGame int32, pnAppID *AppId_t, pnIP *uint32, pnConnPort *uint16, pnQueryPort *uint16, punFlags *uint32, pRTime32LastPlayedOnServer *uint32) bool
	iSteamMatchmaking_AddFavoriteGame                            func(steamMatchmaking uintptr, nAppID AppId_t, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32, rTime32LastPlayedOnServer uint32) int32
	iSteamMatchmaking_RemoveFavoriteGame                         func(steamMatchmaking uintptr, nAppID AppId_t, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32) bool
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
	iSteamMatchmaking_GetLobbyDataByIndex                        func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iLobbyData int32, pchKey []byte, cchKeyBufferSize int32, pchValue []byte, cchValueBufferSize int32) bool
	iSteamMatchmaking_DeleteLobbyData                            func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string) bool
	iSteamMatchmaking_GetLobbyMemberData                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDUser Uint64SteamID, pchKey string) []byte
	iSteamMatchmaking_SetLobbyMemberData                         func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string, pchValue string)
	iSteamMatchmaking_SendLobbyChatMsg                           func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pvMsgBody []byte, cubMsgBody int32) bool
	iSteamMatchmaking_GetLobbyChatEntry                          func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iChatID int32, pSteamIDUser *CSteamID, pvData []byte, cubData int32, peChatEntryType *EChatEntryType) int32
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

type steamMatchmaking uintptr

func SteamMatchmaking() ISteamMatchmaking {
	return steamMatchmaking(steamMatchmaking_init())
}

func (s steamMatchmaking) GetFavoriteGameCount() int32 {
	return iSteamMatchmaking_GetFavoriteGameCount(uintptr(s))
}

func (s steamMatchmaking) GetFavoriteGame(gameIndex int32) (appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32, success bool) {
	success = iSteamMatchmaking_GetFavoriteGame(uintptr(s), gameIndex, &appID, &IP, &connPort, &queryPort, &flags, &lastPlayedOnServerTime)
	return appID, IP, connPort, queryPort, flags, lastPlayedOnServerTime, success
}

func (s steamMatchmaking) AddFavoriteGame(appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32) int32 {
	return iSteamMatchmaking_AddFavoriteGame(uintptr(s), appID, IP, connPort, queryPort, flags, lastPlayedOnServerTime)
}

func (s steamMatchmaking) RemoveFavoriteGame(appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32) bool {
	return iSteamMatchmaking_RemoveFavoriteGame(uintptr(s), appID, IP, connPort, queryPort, flags)
}

func (s steamMatchmaking) RequestLobbyList() CallResult[LobbyMatchList] {
	handle := iSteamMatchmaking_RequestLobbyList(uintptr(s))
	return CallResult[LobbyMatchList]{Handle: handle}
}

func (s steamMatchmaking) AddRequestLobbyListStringFilter(keyToMatch string, valueToMatch string, comparisonType ELobbyComparison) {
	iSteamMatchmaking_AddRequestLobbyListStringFilter(uintptr(s), keyToMatch, valueToMatch, comparisonType)
}

func (s steamMatchmaking) AddRequestLobbyListNumericalFilter(keyToMatch string, valueToMatch int32, comparisonType ELobbyComparison) {
	iSteamMatchmaking_AddRequestLobbyListNumericalFilter(uintptr(s), keyToMatch, valueToMatch, comparisonType)
}

func (s steamMatchmaking) AddRequestLobbyListNearValueFilter(keyToMatch string, valueToBeCloseTo int32) {
	iSteamMatchmaking_AddRequestLobbyListNearValueFilter(uintptr(s), keyToMatch, valueToBeCloseTo)
}

func (s steamMatchmaking) AddRequestLobbyListFilterSlotsAvailable(slotsAvailable int32) {
	iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable(uintptr(s), slotsAvailable)
}

func (s steamMatchmaking) AddRequestLobbyListDistanceFilter(lobbyDistanceFilter ELobbyDistanceFilter) {
	iSteamMatchmaking_AddRequestLobbyListDistanceFilter(uintptr(s), lobbyDistanceFilter)
}

func (s steamMatchmaking) AddRequestLobbyListResultCountFilter(cMaxResults int32) {
	iSteamMatchmaking_AddRequestLobbyListResultCountFilter(uintptr(s), cMaxResults)
}

func (s steamMatchmaking) AddRequestLobbyListCompatibleMembersFilter(lobbySteamID Uint64SteamID) {
	iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) GetLobbyByIndex(iLobby int32) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyByIndex(uintptr(s), iLobby)
}

func (s steamMatchmaking) CreateLobby(lobbyType ELobbyType, maxMembers int32) CallResult[LobbyCreated] {
	handle := iSteamMatchmaking_CreateLobby(uintptr(s), lobbyType, maxMembers)
	return CallResult[LobbyCreated]{Handle: handle}
}

func (s steamMatchmaking) JoinLobby(lobbySteamID Uint64SteamID) CallResult[LobbyEnter] {
	handle := iSteamMatchmaking_JoinLobby(uintptr(s), lobbySteamID)
	return CallResult[LobbyEnter]{Handle: handle}
}

func (s steamMatchmaking) LeaveLobby(lobbySteamID Uint64SteamID) {
	iSteamMatchmaking_LeaveLobby(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) InviteUserToLobby(lobbySteamID Uint64SteamID, inviteeSteamID Uint64SteamID) bool {
	return iSteamMatchmaking_InviteUserToLobby(uintptr(s), lobbySteamID, inviteeSteamID)
}

func (s steamMatchmaking) GetNumLobbyMembers(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetNumLobbyMembers(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) GetLobbyMemberByIndex(lobbySteamID Uint64SteamID, memberIndex int32) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyMemberByIndex(uintptr(s), lobbySteamID, memberIndex)
}

func (s steamMatchmaking) GetLobbyData(lobbySteamID Uint64SteamID, key string) []byte {
	return iSteamMatchmaking_GetLobbyData(uintptr(s), lobbySteamID, key)
}

func (s steamMatchmaking) SetLobbyData(lobbySteamID Uint64SteamID, key string, value string) bool {
	return iSteamMatchmaking_SetLobbyData(uintptr(s), lobbySteamID, key, value)
}

func (s steamMatchmaking) GetLobbyDataCount(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetLobbyDataCount(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) GetLobbyDataByIndex(lobbySteamID Uint64SteamID, lobbyDataIndex int32, keyBufferSize int32, valueBufferSize int32) (key []byte, value []byte, success bool) {
	key = make([]byte, 0, keyBufferSize)
	value = make([]byte, 0, valueBufferSize)
	success = iSteamMatchmaking_GetLobbyDataByIndex(uintptr(s), lobbySteamID, lobbyDataIndex, key, keyBufferSize, value, valueBufferSize)
	return key, value, success
}

func (s steamMatchmaking) DeleteLobbyData(lobbySteamID Uint64SteamID, key string) bool {
	return iSteamMatchmaking_DeleteLobbyData(uintptr(s), lobbySteamID, key)
}

func (s steamMatchmaking) GetLobbyMemberData(lobbySteamID Uint64SteamID, userSteamID Uint64SteamID, key string) []byte {
	return iSteamMatchmaking_GetLobbyMemberData(uintptr(s), lobbySteamID, userSteamID, key)
}

func (s steamMatchmaking) SetLobbyMemberData(lobbySteamID Uint64SteamID, key string, value string) {
	iSteamMatchmaking_SetLobbyMemberData(uintptr(s), lobbySteamID, key, value)
}

func (s steamMatchmaking) SendLobbyChatMsg(lobbySteamID Uint64SteamID, msgBody []byte) bool {
	return iSteamMatchmaking_SendLobbyChatMsg(uintptr(s), lobbySteamID, msgBody, int32(len(msgBody)))
}

func (s steamMatchmaking) GetLobbyChatEntry(lobbySteamID Uint64SteamID, chatID int32, dataSize int32) (userSteamID CSteamID, data []byte, chatEntryType EChatEntryType) {
	data = make([]byte, 0, dataSize)
	result := iSteamMatchmaking_GetLobbyChatEntry(uintptr(s), lobbySteamID, chatID, &userSteamID, data, dataSize, &chatEntryType)
	return userSteamID, data[:result], chatEntryType
}

func (s steamMatchmaking) RequestLobbyData(lobbySteamID Uint64SteamID) bool {
	return iSteamMatchmaking_RequestLobbyData(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) SetLobbyGameServer(lobbySteamID Uint64SteamID, gameServerIP uint32, gameServerPort uint16, gameServerSteamID Uint64SteamID) {
	iSteamMatchmaking_SetLobbyGameServer(uintptr(s), lobbySteamID, gameServerIP, gameServerPort, gameServerSteamID)
}

func (s steamMatchmaking) GetLobbyGameServer(lobbySteamID Uint64SteamID) (gameServerIP uint32, gameServerPort uint16, gameServerSteamID CSteamID, success bool) {
	success = iSteamMatchmaking_GetLobbyGameServer(uintptr(s), lobbySteamID, &gameServerIP, &gameServerPort, &gameServerSteamID)
	return gameServerIP, gameServerPort, gameServerSteamID, success
}

func (s steamMatchmaking) SetLobbyMemberLimit(lobbySteamID Uint64SteamID, cMaxMembers int32) bool {
	return iSteamMatchmaking_SetLobbyMemberLimit(uintptr(s), lobbySteamID, cMaxMembers)
}

func (s steamMatchmaking) GetLobbyMemberLimit(lobbySteamID Uint64SteamID) int32 {
	return iSteamMatchmaking_GetLobbyMemberLimit(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) SetLobbyType(lobbySteamID Uint64SteamID, lobbyType ELobbyType) bool {
	return iSteamMatchmaking_SetLobbyType(uintptr(s), lobbySteamID, lobbyType)
}

func (s steamMatchmaking) SetLobbyJoinable(lobbySteamID Uint64SteamID, joinableLobby bool) bool {
	return iSteamMatchmaking_SetLobbyJoinable(uintptr(s), lobbySteamID, joinableLobby)
}

func (s steamMatchmaking) GetLobbyOwner(lobbySteamID Uint64SteamID) Uint64SteamID {
	return iSteamMatchmaking_GetLobbyOwner(uintptr(s), lobbySteamID)
}

func (s steamMatchmaking) SetLobbyOwner(lobbySteamID Uint64SteamID, newUserSteamID Uint64SteamID) bool {
	return iSteamMatchmaking_SetLobbyOwner(uintptr(s), lobbySteamID, newUserSteamID)
}

func (s steamMatchmaking) SetLinkedLobby(lobbySteamID Uint64SteamID, dependentLobbySteamID Uint64SteamID) bool {
	return iSteamMatchmaking_SetLinkedLobby(uintptr(s), lobbySteamID, dependentLobbySteamID)
}

// Steam Matchmaking Servers
const (
	flatAPI_MatchMakingKeyValuePair_t_Construct = "SteamAPI_MatchMakingKeyValuePair_t_Construct"
)

var (
	matchMakingKeyValuePair_t_Construct func(*MatchMakingKeyValuePair)
)

type MatchMakingKeyValuePair struct {
	Key   [256]byte
	Value [256]byte
}

func (m *MatchMakingKeyValuePair) Construct() {
	matchMakingKeyValuePair_t_Construct(m)
}

type ServerNetAdr struct {
	ConnectionPort uint16
	QueryPort      uint16
	IP             uint32
}

const (
	flatAPI_servernetadr_t_Construct                  = "SteamAPI_servernetadr_t_Construct"
	flatAPI_servernetadr_t_Init                       = "SteamAPI_servernetadr_t_Init"
	flatAPI_servernetadr_t_GetQueryPort               = "SteamAPI_servernetadr_t_GetQueryPort"
	flatAPI_servernetadr_t_SetQueryPort               = "SteamAPI_servernetadr_t_SetQueryPort"
	flatAPI_servernetadr_t_GetConnectionPort          = "SteamAPI_servernetadr_t_GetConnectionPort"
	flatAPI_servernetadr_t_SetConnectionPort          = "SteamAPI_servernetadr_t_SetConnectionPort"
	flatAPI_servernetadr_t_GetIP                      = "SteamAPI_servernetadr_t_GetIP"
	flatAPI_servernetadr_t_SetIP                      = "SteamAPI_servernetadr_t_SetIP"
	flatAPI_servernetadr_t_GetConnectionAddressString = "SteamAPI_servernetadr_t_GetConnectionAddressString"
	flatAPI_servernetadr_t_GetQueryAddressString      = "SteamAPI_servernetadr_t_GetQueryAddressString"
	flatAPI_servernetadr_t_IsLessThan                 = "SteamAPI_servernetadr_t_IsLessThan"
	flatAPI_servernetadr_t_Assign                     = "SteamAPI_servernetadr_t_Assign"
)

var (
	servernetadr_t_Construct                  func(self *ServerNetAdr)
	servernetadr_t_Init                       func(self *ServerNetAdr, ip uint, QueryPort uint16, ConnectionPort uint16)
	servernetadr_t_GetQueryPort               func(self *ServerNetAdr) uint16
	servernetadr_t_SetQueryPort               func(self *ServerNetAdr, Port uint16)
	servernetadr_t_GetConnectionPort          func(self *ServerNetAdr) uint16
	servernetadr_t_SetConnectionPort          func(self *ServerNetAdr, Port uint16)
	servernetadr_t_GetIP                      func(self *ServerNetAdr) uint32
	servernetadr_t_SetIP                      func(self *ServerNetAdr, IP uint32)
	servernetadr_t_GetConnectionAddressString func(self *ServerNetAdr) string
	servernetadr_t_GetQueryAddressString      func(self *ServerNetAdr) string
	servernetadr_t_IsLessThan                 func(self *ServerNetAdr, netadr *ServerNetAdr) bool
	servernetadr_t_Assign                     func(self *ServerNetAdr, that *ServerNetAdr)
)

func (s *ServerNetAdr) Construct() {
	servernetadr_t_Construct(s)
}

func (s *ServerNetAdr) Init(ip uint, QueryPort uint16, ConnectionPort uint16) {
	servernetadr_t_Init(s, ip, QueryPort, ConnectionPort)
}
func (s *ServerNetAdr) GetQueryPort() uint16 {
	return servernetadr_t_GetQueryPort(s)
}
func (s *ServerNetAdr) SetQueryPort(Port uint16) {
	servernetadr_t_SetQueryPort(s, Port)
}
func (s *ServerNetAdr) GetConnectionPort() uint16 {
	return servernetadr_t_GetConnectionPort(s)
}
func (s *ServerNetAdr) SetConnectionPort(Port uint16) {
	servernetadr_t_SetConnectionPort(s, Port)
}
func (s *ServerNetAdr) GetIP() uint32 {
	return servernetadr_t_GetIP(s)
}
func (s *ServerNetAdr) SetIP(IP uint32) {
	servernetadr_t_SetIP(s, IP)
}
func (s *ServerNetAdr) GetConnectionAddressString() string {
	return servernetadr_t_GetConnectionAddressString(s)
}
func (s *ServerNetAdr) GetQueryAddressString() string {
	return servernetadr_t_GetQueryAddressString(s)
}
func (s *ServerNetAdr) IsLessThan(netadr *ServerNetAdr) bool {
	return servernetadr_t_IsLessThan(s, netadr)
}
func (s *ServerNetAdr) Assign(that *ServerNetAdr) {
	servernetadr_t_Assign(s, that)
}

type GameServerItem struct {
	NetAdr                ServerNetAdr
	Ping                  int32
	HadSuccessfulResponse bool
	DoNotRefresh          bool
	_                     [2]byte
	GameDir               [32]byte
	Map                   [32]byte
	GameDescription       [64]byte
	AppID                 uint32
	Players               int32
	MaxPlayers            int32
	BotPlayers            int32
	Password              bool
	Secure                bool
	_                     [2]byte
	TimeLastPlayed        uint32
	ServerVersion         int32
	ServerName            [64]byte
	GameTags              [128]byte
	ID                    CSteamID
}

func (g *GameServerItem) Construct() {
	gameserveritem_t_Construct(g)
}

func (g *GameServerItem) GetName() string {
	return gameserveritem_t_GetName(g)
}

func (g *GameServerItem) SetName(name string) {
	gameserveritem_t_SetName(g, name)
}

const (
	flatAPI_gameserveritem_t_Construct = "SteamAPI_gameserveritem_t_Construct"
	flatAPI_gameserveritem_t_GetName   = "SteamAPI_gameserveritem_t_GetName"
	flatAPI_gameserveritem_t_SetName   = "SteamAPI_gameserveritem_t_SetName"
)

var (
	gameserveritem_t_Construct func(self *GameServerItem)
	gameserveritem_t_GetName   func(self *GameServerItem) string
	gameserveritem_t_SetName   func(self *GameServerItem, Name string)
)

type HServerListRequest uintptr
type HServerQuery int32

const (
	flatAPI_SteamMatchmakingServers                                   = "SteamAPI_SteamMatchmakingServers_v002"
	flatAPI_ISteamMatchmakingServers_RequestInternetServerList        = "SteamAPI_ISteamMatchmakingServers_RequestInternetServerList"
	flatAPI_ISteamMatchmakingServers_RequestLANServerList             = "SteamAPI_ISteamMatchmakingServers_RequestLANServerList"
	flatAPI_ISteamMatchmakingServers_RequestFriendsServerList         = "SteamAPI_ISteamMatchmakingServers_RequestFriendsServerList"
	flatAPI_ISteamMatchmakingServers_RequestFavoritesServerList       = "SteamAPI_ISteamMatchmakingServers_RequestFavoritesServerList"
	flatAPI_ISteamMatchmakingServers_RequestHistoryServerList         = "SteamAPI_ISteamMatchmakingServers_RequestHistoryServerList"
	flatAPI_ISteamMatchmakingServers_RequestSpectatorServerList       = "SteamAPI_ISteamMatchmakingServers_RequestSpectatorServerList"
	flatAPI_ISteamMatchmakingServers_ReleaseRequest                   = "SteamAPI_ISteamMatchmakingServers_ReleaseRequest"
	flatAPI_ISteamMatchmakingServers_GetServerDetails                 = "SteamAPI_ISteamMatchmakingServers_GetServerDetails"
	flatAPI_ISteamMatchmakingServers_CancelQuery                      = "SteamAPI_ISteamMatchmakingServers_CancelQuery"
	flatAPI_ISteamMatchmakingServers_RefreshQuery                     = "SteamAPI_ISteamMatchmakingServers_RefreshQuery"
	flatAPI_ISteamMatchmakingServers_IsRefreshing                     = "SteamAPI_ISteamMatchmakingServers_IsRefreshing"
	flatAPI_ISteamMatchmakingServers_GetServerCount                   = "SteamAPI_ISteamMatchmakingServers_GetServerCount"
	flatAPI_ISteamMatchmakingServers_RefreshServer                    = "SteamAPI_ISteamMatchmakingServers_RefreshServer"
	flatAPI_ISteamMatchmakingServers_PingServer                       = "SteamAPI_ISteamMatchmakingServers_PingServer"
	flatAPI_ISteamMatchmakingServers_PlayerDetails                    = "SteamAPI_ISteamMatchmakingServers_PlayerDetails"
	flatAPI_ISteamMatchmakingServers_ServerRules                      = "SteamAPI_ISteamMatchmakingServers_ServerRules"
	flatAPI_ISteamMatchmakingServers_CancelServerQuery                = "SteamAPI_ISteamMatchmakingServers_CancelServerQuery"
	flatAPI_ISteamMatchmakingServerListResponse_ServerResponded       = "SteamAPI_ISteamMatchmakingServerListResponse_ServerResponded"
	flatAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond = "SteamAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond"
	flatAPI_ISteamMatchmakingServerListResponse_RefreshComplete       = "SteamAPI_ISteamMatchmakingServerListResponse_RefreshComplete"
	flatAPI_ISteamMatchmakingPingResponse_ServerResponded             = "SteamAPI_ISteamMatchmakingPingResponse_ServerResponded"
	flatAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond       = "SteamAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond"
	flatAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList          = "SteamAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList"
	flatAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond   = "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond"
	flatAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete   = "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete"
	flatAPI_ISteamMatchmakingRulesResponse_RulesResponded             = "SteamAPI_ISteamMatchmakingRulesResponse_RulesResponded"
	flatAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond       = "SteamAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond"
	flatAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete       = "SteamAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete"
)

var (
	steamMatchmakingServers_init                              func() uintptr
	iSteamMatchmakingServers_RequestInternetServerList        func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestLANServerList             func(steamMatchmakingServers uintptr, iApp AppId_t, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestFriendsServerList         func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestFavoritesServerList       func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestHistoryServerList         func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestSpectatorServerList       func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_ReleaseRequest                   func(steamMatchmakingServers uintptr, hServerListRequest HServerListRequest)
	iSteamMatchmakingServers_GetServerDetails                 func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32) *GameServerItem
	iSteamMatchmakingServers_CancelQuery                      func(steamMatchmakingServers uintptr, hRequest HServerListRequest)
	iSteamMatchmakingServers_RefreshQuery                     func(steamMatchmakingServers uintptr, hRequest HServerListRequest)
	iSteamMatchmakingServers_IsRefreshing                     func(steamMatchmakingServers uintptr, hRequest HServerListRequest) bool
	iSteamMatchmakingServers_GetServerCount                   func(steamMatchmakingServers uintptr, hRequest HServerListRequest) int32
	iSteamMatchmakingServers_RefreshServer                    func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServers_PingServer                       func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *ISteamMatchmakingPingResponse) HServerQuery
	iSteamMatchmakingServers_PlayerDetails                    func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingPlayersResponse) HServerQuery
	iSteamMatchmakingServers_ServerRules                      func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingRulesResponse) HServerQuery
	iSteamMatchmakingServers_CancelServerQuery                func(steamMatchmakingServers uintptr, hServerQuery HServerQuery)
	iSteamMatchmakingServerListResponse_ServerResponded       func(hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServerListResponse_ServerFailedToRespond func(hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServerListResponse_RefreshComplete       func(hRequest HServerListRequest, response EMatchMakingServerResponse)
	iSteamMatchmakingPingResponse_ServerResponded             func(server *GameServerItem)
	iSteamMatchmakingPingResponse_ServerFailedToRespond       func()
	iSteamMatchmakingPlayersResponse_AddPlayerToList          func(pchName string, nScore int32, flTimePlayed float32)
	iSteamMatchmakingPlayersResponse_PlayersFailedToRespond   func()
	iSteamMatchmakingPlayersResponse_PlayersRefreshComplete   func()
	iSteamMatchmakingRulesResponse_RulesResponded             func(pchRule string, pchValue string)
	iSteamMatchmakingRulesResponse_RulesFailedToRespond       func()
	iSteamMatchmakingRulesResponse_RulesRefreshComplete       func()
)

type steamMatchmakingServers uintptr

func SteamMatchmakingServers() ISteamMatchmakingServers {
	return steamMatchmakingServers(steamMatchmakingServers_init())
}

func (s steamMatchmakingServers) RequestInternetServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestInternetServerList(uintptr(s), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func (s steamMatchmakingServers) RequestLANServerList(App AppId_t) (RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	request = iSteamMatchmakingServers_RequestLANServerList(uintptr(s), App, &RequestServersResponse)
	return RequestServersResponse, request
}

func (s steamMatchmakingServers) RequestFriendsServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestFriendsServerList(uintptr(s), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func (s steamMatchmakingServers) RequestFavoritesServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestFavoritesServerList(uintptr(s), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func (s steamMatchmakingServers) RequestHistoryServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestHistoryServerList(uintptr(s), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func (s steamMatchmakingServers) RequestSpectatorServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestSpectatorServerList(uintptr(s), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func (s steamMatchmakingServers) ReleaseRequest(ServerListRequest HServerListRequest) {
	iSteamMatchmakingServers_ReleaseRequest(uintptr(s), ServerListRequest)
}

func (s steamMatchmakingServers) GetServerDetails(Request HServerListRequest, Server int32) *GameServerItem {
	return iSteamMatchmakingServers_GetServerDetails(uintptr(s), Request, Server)
}

func (s steamMatchmakingServers) CancelQuery(Request HServerListRequest) {
	iSteamMatchmakingServers_CancelQuery(uintptr(s), Request)
}

func (s steamMatchmakingServers) RefreshQuery(Request HServerListRequest) {
	iSteamMatchmakingServers_RefreshQuery(uintptr(s), Request)
}

func (s steamMatchmakingServers) IsRefreshing(Request HServerListRequest) bool {
	return iSteamMatchmakingServers_IsRefreshing(uintptr(s), Request)
}

func (s steamMatchmakingServers) GetServerCount(Request HServerListRequest) int32 {
	return iSteamMatchmakingServers_GetServerCount(uintptr(s), Request)
}

func (s steamMatchmakingServers) RefreshServer(Request HServerListRequest, Server int32) {
	iSteamMatchmakingServers_RefreshServer(uintptr(s), Request, Server)
}

func (s steamMatchmakingServers) PingServer(IP uint32, Port uint16) (RequestServersResponse ISteamMatchmakingPingResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_PingServer(uintptr(s), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func (s steamMatchmakingServers) PlayerDetails(IP uint32, Port uint16) (RequestServersResponse MatchmakingPlayersResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_PlayerDetails(uintptr(s), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func (s steamMatchmakingServers) ServerRules(IP uint32, Port uint16) (RequestServersResponse MatchmakingRulesResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_ServerRules(uintptr(s), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func (s steamMatchmakingServers) CancelServerQuery(serverQuery HServerQuery) {
	iSteamMatchmakingServers_CancelServerQuery(uintptr(s), serverQuery)
}

// Steam Network Types
const (
	HSteamNetConnection_Invalid                            HSteamNetConnection  = 0
	HSteamListenSocket_Invalid                             HSteamListenSocket   = 0
	HSteamNetPollGroup_Invalid                             HSteamNetPollGroup   = 0
	MaxSteamNetworkingErrMsg                               int32                = 1024
	SteamNetworkingMaxConnectionCloseReason                int32                = 128
	SteamNetworkingMaxConnectionDescription                int32                = 128
	SteamNetworkingMaxConnectionAppName                    int32                = 32
	SteamNetworkConnectionInfoFlags_Unauthenticated        int32                = 1
	SteamNetworkConnectionInfoFlags_Unencrypted            int32                = 2
	SteamNetworkConnectionInfoFlags_LoopbackBuffers        int32                = 4
	SteamNetworkConnectionInfoFlags_Fast                   int32                = 8
	SteamNetworkConnectionInfoFlags_Relayed                int32                = 16
	SteamNetworkConnectionInfoFlags_DualWifi               int32                = 32
	MaxSteamNetworkingSocketsMessageSizeSend               int32                = 512 * 1024
	SteamNetworkingSend_Unreliable                         int32                = 0
	SteamNetworkingSend_NoNagle                            int32                = 1
	SteamNetworkingSend_UnreliableNoNagle                  int32                = SteamNetworkingSend_Unreliable | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_NoDelay                            int32                = 4
	SteamNetworkingSend_UnreliableNoDelay                  int32                = SteamNetworkingSend_Unreliable | SteamNetworkingSend_NoDelay | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_Reliable                           int32                = 8
	SteamNetworkingSend_ReliableNoNagle                    int32                = SteamNetworkingSend_Reliable | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_UseCurrentThread                   int32                = 16
	SteamNetworkingSend_AutoRestartBrokenSession           int32                = 32
	MaxSteamNetworkingPingLocationString                   int32                = 1024
	SteamNetworkingPing_Failed                             int32                = -1
	SteamNetworkingPing_Unknown                            int32                = -2
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Default int32                = -1
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Disable int32                = 0
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Relay   int32                = 1
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Private int32                = 2
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Public  int32                = 4
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_All     int32                = 0x7fffffff
	SteamDatagramPOPID_dev                                 SteamNetworkingPOPID = (SteamNetworkingPOPID('d') << 16) | (SteamNetworkingPOPID('e') << 8) | SteamNetworkingPOPID('v')
	SteamDatagramMaxSerializedTicket                       uint32               = 512
	MaxSteamDatagramGameCoordinatorServerLoginAppData      uint32               = 2048
	MaxSteamDatagramGameCoordinatorServerLoginSerialized   uint32               = 4096
	SteamNetworkingSocketsFakeUDPPortRecommendedMTU        int32                = 1200
	SteamNetworkingSocketsFakeUDPPortMaxMessageSize        int32                = 4096
)

type ESteamIPType int32

const (
	ESteamIPTypeIPv4 ESteamIPType = 0
	ESteamIPTypeIPv6 ESteamIPType = 1
)

type SteamIPAddress struct {
	IPv6 [16]uint8
	Type ESteamIPType
	_    [2]byte
}

type HSteamNetConnection uint
type HSteamListenSocket uint
type HSteamNetPollGroup uint
type SteamDatagramRelayAuthTicket struct{}
type ISteamNetworkingConnectionSignaling struct{}
type ISteamNetworkingSignalingRecvContext struct{}
type SteamNetworkingFakeUDPPort struct{}

type SteamNetworkingErrMsg [1024]byte
type SteamDatagramHostedAddress struct {
	Size int32
	Data [128]byte
}
type SteamDatagramGameCoordinatorServerLogin struct {
	Identity SteamNetworkingIdentity
	Routing  SteamDatagramHostedAddress
	AppID    AppId_t
	RTime    RTime32
	AppData  int32
	Data     [2048]byte
}

type ESteamNetworkingConnectionState int32

const (
	ESteamNetworkingConnectionState_None                   ESteamNetworkingConnectionState = 0
	ESteamNetworkingConnectionState_Connecting             ESteamNetworkingConnectionState = 1
	ESteamNetworkingConnectionState_FindingRoute           ESteamNetworkingConnectionState = 2
	ESteamNetworkingConnectionState_Connected              ESteamNetworkingConnectionState = 3
	ESteamNetworkingConnectionState_ClosedByPeer           ESteamNetworkingConnectionState = 4
	ESteamNetworkingConnectionState_ProblemDetectedLocally ESteamNetworkingConnectionState = 5
	ESteamNetworkingConnectionState_FinWait                ESteamNetworkingConnectionState = -1
	ESteamNetworkingConnectionState_Linger                 ESteamNetworkingConnectionState = -2
	ESteamNetworkingConnectionState_Dead                   ESteamNetworkingConnectionState = -3
	ESteamNetworkingConnectionState_Force32Bit             ESteamNetworkingConnectionState = 2147483647
)

type SteamNetConnectionInfo struct {
	IdentityRemote        SteamNetworkingIdentity
	UserData              int64
	ListenSocket          HSteamListenSocket
	AddrRemote            SteamNetworkingIPAddr
	pad1                  uint16
	_                     [2]byte
	POPRemoteID           SteamNetworkingPOPID
	POPRelayID            SteamNetworkingPOPID
	State                 ESteamNetworkingConnectionState
	EndReason             int32
	EndDebug              [128]byte
	ConnectionDescription [128]byte
	Flags                 int32
	reserved              [63]uint32
}
type SteamNetConnectionRealTimeStatus struct {
	State                   ESteamNetworkingConnectionState
	Ping                    int32
	ConnectionQualityLocal  float32
	ConnectionQualityRemote float32
	OutPacketsPerSec        float32
	OutBytesPerSec          float32
	InPacketsPerSec         float32
	InBytesPerSec           float32
	SendRateBytesPerSecond  int32
	PendingUnreliable       int32
	PendingReliable         int32
	SentUnackedReliable     int32
	QueueTime               SteamNetworkingMicroseconds
	reserved                [16]uint32
}
type SteamNetConnectionRealTimeLaneStatus struct {
	PendingUnreliable   int32
	PendingReliable     int32
	SentUnackedReliable int32
	reservePad1         int32
	QueueTime           SteamNetworkingMicroseconds
	reserved            [10]uint32
}

type SteamNetworkingMessage struct {
	Data          unsafe.Pointer
	Size          int32
	Conn          HSteamNetConnection
	IdentityPeer  SteamNetworkingIdentity
	ConnUserData  int64
	TimeReceived  SteamNetworkingMicroseconds
	MessageNumber int64
	FreeData      uintptr
	Release       uintptr
	Channel       int32
	Flags         int32
	UserData      int64
	Lane          uint16
	pad1          uint16
	_             [4]byte
}

type FnSteamRelayNetworkStatusChanged uintptr

type ESteamNetworkingAvailability int32

const (
	ESteamNetworkingAvailability_CannotTry  ESteamNetworkingAvailability = -102
	ESteamNetworkingAvailability_Failed     ESteamNetworkingAvailability = -101
	ESteamNetworkingAvailability_Previously ESteamNetworkingAvailability = -100
	ESteamNetworkingAvailability_Retrying   ESteamNetworkingAvailability = -10
	ESteamNetworkingAvailability_NeverTried ESteamNetworkingAvailability = 1
	ESteamNetworkingAvailability_Waiting    ESteamNetworkingAvailability = 2
	ESteamNetworkingAvailability_Attempting ESteamNetworkingAvailability = 3
	ESteamNetworkingAvailability_Current    ESteamNetworkingAvailability = 100
	ESteamNetworkingAvailability_Unknown    ESteamNetworkingAvailability = 0
	ESteamNetworkingAvailability_Force32bit ESteamNetworkingAvailability = 2147483647
)

type ESteamNetworkingSocketsDebugOutputType int32

const (
	ESteamNetworkingSocketsDebugOutputType_None       ESteamNetworkingSocketsDebugOutputType = 0
	ESteamNetworkingSocketsDebugOutputType_Bug        ESteamNetworkingSocketsDebugOutputType = 1
	ESteamNetworkingSocketsDebugOutputType_Error      ESteamNetworkingSocketsDebugOutputType = 2
	ESteamNetworkingSocketsDebugOutputType_Important  ESteamNetworkingSocketsDebugOutputType = 3
	ESteamNetworkingSocketsDebugOutputType_Warning    ESteamNetworkingSocketsDebugOutputType = 4
	ESteamNetworkingSocketsDebugOutputType_Msg        ESteamNetworkingSocketsDebugOutputType = 5
	ESteamNetworkingSocketsDebugOutputType_Verbose    ESteamNetworkingSocketsDebugOutputType = 6
	ESteamNetworkingSocketsDebugOutputType_Debug      ESteamNetworkingSocketsDebugOutputType = 7
	ESteamNetworkingSocketsDebugOutputType_Everything ESteamNetworkingSocketsDebugOutputType = 8
	ESteamNetworkingSocketsDebugOutputType_Force32Bit ESteamNetworkingSocketsDebugOutputType = 2147483647
)

type ESteamNetworkingFakeIPType int32

const (
	ESteamNetworkingFakeIPType_Invalid    ESteamNetworkingFakeIPType = 0
	ESteamNetworkingFakeIPType_NotFake    ESteamNetworkingFakeIPType = 1
	ESteamNetworkingFakeIPType_GlobalIPv4 ESteamNetworkingFakeIPType = 2
	ESteamNetworkingFakeIPType_LocalIPv4  ESteamNetworkingFakeIPType = 3
	ESteamNetworkingFakeIPType_Force32Bit ESteamNetworkingFakeIPType = 2147483647
)

type SteamNetworkPingLocation struct {
	Data [512]uint8
}

type SteamNetworkingPOPID uint

type SteamNetworkingMicroseconds int64

type FSteamNetworkingSocketsDebugOutput uintptr

type SteamNetworkingIPAddr struct {
	IPv6 [16]uint8
	Port uint16
}
type ESteamNetworkingIdentityType int32

const (
	ESteamNetworkingIdentityType_Invalid        ESteamNetworkingIdentityType = 0
	ESteamNetworkingIdentityType_SteamID        ESteamNetworkingIdentityType = 16
	ESteamNetworkingIdentityType_XboxPairwiseID ESteamNetworkingIdentityType = 17
	ESteamNetworkingIdentityType_SonyPSN        ESteamNetworkingIdentityType = 18
	ESteamNetworkingIdentityType_IPAddress      ESteamNetworkingIdentityType = 1
	ESteamNetworkingIdentityType_GenericString  ESteamNetworkingIdentityType = 2
	ESteamNetworkingIdentityType_GenericBytes   ESteamNetworkingIdentityType = 3
	ESteamNetworkingIdentityType_UnknownType    ESteamNetworkingIdentityType = 4
	ESteamNetworkingIdentityType_Force32bit     ESteamNetworkingIdentityType = 2147483647
)

type SteamNetworkingConfigValue struct {
	Value    ESteamNetworkingConfigValue
	DataType ESteamNetworkingConfigDataType
	Int64    int64
}
type SteamNetworkingIdentity struct {
	Type             ESteamNetworkingIdentityType
	Size             int32
	UnknownRawString [128]byte
}

type ESteamNetworkingConfigValue int32

const (
	ESteamNetworkingConfig_Invalid                                       ESteamNetworkingConfigValue = 0
	ESteamNetworkingConfig_TimeoutInitial                                ESteamNetworkingConfigValue = 24
	ESteamNetworkingConfig_TimeoutConnected                              ESteamNetworkingConfigValue = 25
	ESteamNetworkingConfig_SendBufferSize                                ESteamNetworkingConfigValue = 9
	ESteamNetworkingConfig_RecvBufferSize                                ESteamNetworkingConfigValue = 47
	ESteamNetworkingConfig_RecvBufferMessages                            ESteamNetworkingConfigValue = 48
	ESteamNetworkingConfig_RecvMaxMessageSize                            ESteamNetworkingConfigValue = 49
	ESteamNetworkingConfig_RecvMaxSegmentsPerPacket                      ESteamNetworkingConfigValue = 50
	ESteamNetworkingConfig_ConnectionUserData                            ESteamNetworkingConfigValue = 40
	ESteamNetworkingConfig_SendRateMin                                   ESteamNetworkingConfigValue = 10
	ESteamNetworkingConfig_SendRateMax                                   ESteamNetworkingConfigValue = 11
	ESteamNetworkingConfig_NagleTime                                     ESteamNetworkingConfigValue = 12
	ESteamNetworkingConfig_IPAllowWithoutAuth                            ESteamNetworkingConfigValue = 23
	ESteamNetworkingConfig_IPLocalHostAllowWithoutAuth                   ESteamNetworkingConfigValue = 52
	ESteamNetworkingConfig_MTUPacketSize                                 ESteamNetworkingConfigValue = 32
	ESteamNetworkingConfig_MTUDataSize                                   ESteamNetworkingConfigValue = 33
	ESteamNetworkingConfig_Unencrypted                                   ESteamNetworkingConfigValue = 34
	ESteamNetworkingConfig_SymmetricConnect                              ESteamNetworkingConfigValue = 37
	ESteamNetworkingConfig_LocalVirtualPort                              ESteamNetworkingConfigValue = 38
	ESteamNetworkingConfig_DualWifiEnable                                ESteamNetworkingConfigValue = 39
	ESteamNetworkingConfig_EnableDiagnosticsUI                           ESteamNetworkingConfigValue = 46
	ESteamNetworkingConfig_SendTimeSincePreviousPacket                   ESteamNetworkingConfigValue = 59
	ESteamNetworkingConfig_FakePacketLossSend                            ESteamNetworkingConfigValue = 2
	ESteamNetworkingConfig_FakePacketLossRecv                            ESteamNetworkingConfigValue = 3
	ESteamNetworkingConfig_FakePacketLagSend                             ESteamNetworkingConfigValue = 4
	ESteamNetworkingConfig_FakePacketLagRecv                             ESteamNetworkingConfigValue = 5
	ESteamNetworkingConfig_FakePacketJitterSendAvg                       ESteamNetworkingConfigValue = 53
	ESteamNetworkingConfig_FakePacketJitterSendMax                       ESteamNetworkingConfigValue = 54
	ESteamNetworkingConfig_FakePacketJitterSendPct                       ESteamNetworkingConfigValue = 55
	ESteamNetworkingConfig_FakePacketJitterRecvAvg                       ESteamNetworkingConfigValue = 56
	ESteamNetworkingConfig_FakePacketJitterRecvMax                       ESteamNetworkingConfigValue = 57
	ESteamNetworkingConfig_FakePacketJitterRecvPct                       ESteamNetworkingConfigValue = 58
	ESteamNetworkingConfig_FakePacketReorderSend                         ESteamNetworkingConfigValue = 6
	ESteamNetworkingConfig_FakePacketReorderRecv                         ESteamNetworkingConfigValue = 7
	ESteamNetworkingConfig_FakePacketReorderTime                         ESteamNetworkingConfigValue = 8
	ESteamNetworkingConfig_FakePacketDupSend                             ESteamNetworkingConfigValue = 26
	ESteamNetworkingConfig_FakePacketDupRecv                             ESteamNetworkingConfigValue = 27
	ESteamNetworkingConfig_FakePacketDupTimeMax                          ESteamNetworkingConfigValue = 28
	ESteamNetworkingConfig_PacketTraceMaxBytes                           ESteamNetworkingConfigValue = 41
	ESteamNetworkingConfig_FakeRateLimitSendRate                         ESteamNetworkingConfigValue = 42
	ESteamNetworkingConfig_FakeRateLimitSendBurst                        ESteamNetworkingConfigValue = 43
	ESteamNetworkingConfig_FakeRateLimitRecvRate                         ESteamNetworkingConfigValue = 44
	ESteamNetworkingConfig_FakeRateLimitRecvBurst                        ESteamNetworkingConfigValue = 45
	ESteamNetworkingConfig_OutOfOrderCorrectionWindowMicroseconds        ESteamNetworkingConfigValue = 51
	ESteamNetworkingConfig_CallbackConnectionStatusChanged               ESteamNetworkingConfigValue = 201
	ESteamNetworkingConfig_CallbackAuthStatusChanged                     ESteamNetworkingConfigValue = 202
	ESteamNetworkingConfig_CallbackRelayNetworkStatusChanged             ESteamNetworkingConfigValue = 203
	ESteamNetworkingConfig_CallbackMessagesSessionRequest                ESteamNetworkingConfigValue = 204
	ESteamNetworkingConfig_CallbackMessagesSessionFailed                 ESteamNetworkingConfigValue = 205
	ESteamNetworkingConfig_CallbackCreateConnectionSignaling             ESteamNetworkingConfigValue = 206
	ESteamNetworkingConfig_CallbackFakeIPResult                          ESteamNetworkingConfigValue = 207
	ESteamNetworkingConfig_P2PSTUNServerList                             ESteamNetworkingConfigValue = 103
	ESteamNetworkingConfig_P2PTransportICEEnable                         ESteamNetworkingConfigValue = 104
	ESteamNetworkingConfig_P2PTransportICEPenalty                        ESteamNetworkingConfigValue = 105
	ESteamNetworkingConfig_P2PTransportSDRPenalty                        ESteamNetworkingConfigValue = 106
	ESteamNetworkingConfig_P2PTURNServerList                             ESteamNetworkingConfigValue = 107
	ESteamNetworkingConfig_P2PTURNUserList                               ESteamNetworkingConfigValue = 108
	ESteamNetworkingConfig_P2PTURNPassList                               ESteamNetworkingConfigValue = 109
	ESteamNetworkingConfig_P2PTransportICEImplementation                 ESteamNetworkingConfigValue = 110
	ESteamNetworkingConfig_SDRClientConsecutitivePingTimeoutsFailInitial ESteamNetworkingConfigValue = 19
	ESteamNetworkingConfig_SDRClientConsecutitivePingTimeoutsFail        ESteamNetworkingConfigValue = 20
	ESteamNetworkingConfig_SDRClientMinPingsBeforePingAccurate           ESteamNetworkingConfigValue = 21
	ESteamNetworkingConfig_SDRClientSingleSocket                         ESteamNetworkingConfigValue = 22
	ESteamNetworkingConfig_SDRClientForceRelayCluster                    ESteamNetworkingConfigValue = 29
	ESteamNetworkingConfig_SDRClientDevTicket                            ESteamNetworkingConfigValue = 30
	ESteamNetworkingConfig_SDRClientForceProxyAddr                       ESteamNetworkingConfigValue = 31
	ESteamNetworkingConfig_SDRClientFakeClusterPing                      ESteamNetworkingConfigValue = 36
	ESteamNetworkingConfig_SDRClientLimitPingProbesToNearestN            ESteamNetworkingConfigValue = 60
	ESteamNetworkingConfig_LogLevelAckRTT                                ESteamNetworkingConfigValue = 13
	ESteamNetworkingConfig_LogLevelPacketDecode                          ESteamNetworkingConfigValue = 14
	ESteamNetworkingConfig_LogLevelMessage                               ESteamNetworkingConfigValue = 15
	ESteamNetworkingConfig_LogLevelPacketGaps                            ESteamNetworkingConfigValue = 16
	ESteamNetworkingConfig_LogLevelP2PRendezvous                         ESteamNetworkingConfigValue = 17
	ESteamNetworkingConfig_LogLevelSDRRelayPings                         ESteamNetworkingConfigValue = 18
	ESteamNetworkingConfig_ECN                                           ESteamNetworkingConfigValue = 999
	ESteamNetworkingConfig_DELETEDEnumerateDevVars                       ESteamNetworkingConfigValue = 35
	ESteamNetworkingConfig_ValueForce32Bit                               ESteamNetworkingConfigValue = 2147483647
)

type FnSteamNetConnectionStatusChanged uintptr
type FnSteamNetAuthenticationStatusChanged uintptr
type FnSteamNetworkingMessagesSessionRequest uintptr
type FnSteamNetworkingMessagesSessionFailed uintptr
type FnSteamNetworkingFakeIPResult uintptr

type ESteamNetworkingConfigScope int32

const (
	ESteamNetworkingConfig_Global           ESteamNetworkingConfigScope = 1
	ESteamNetworkingConfig_SocketsInterface ESteamNetworkingConfigScope = 2
	ESteamNetworkingConfig_ListenSocket     ESteamNetworkingConfigScope = 3
	ESteamNetworkingConfig_Connection       ESteamNetworkingConfigScope = 4
	ESteamNetworkingConfig_ScopeForce32Bit  ESteamNetworkingConfigScope = 2147483647
)

type ESteamNetworkingConfigDataType int32

const (
	ESteamNetworkingConfig_Int32              ESteamNetworkingConfigDataType = 1
	ESteamNetworkingConfig_Int64              ESteamNetworkingConfigDataType = 2
	ESteamNetworkingConfig_Float              ESteamNetworkingConfigDataType = 3
	ESteamNetworkingConfig_String             ESteamNetworkingConfigDataType = 4
	ESteamNetworkingConfig_Ptr                ESteamNetworkingConfigDataType = 5
	ESteamNetworkingConfig_DataTypeForce32Bit ESteamNetworkingConfigDataType = 2147483647
)

type ESteamNetworkingGetConfigValueResult int32

const (
	ESteamNetworkingGetConfigValue_BadValue         ESteamNetworkingGetConfigValueResult = -1
	ESteamNetworkingGetConfigValue_BadScopeObj      ESteamNetworkingGetConfigValueResult = -2
	ESteamNetworkingGetConfigValue_BufferTooSmall   ESteamNetworkingGetConfigValueResult = -3
	ESteamNetworkingGetConfigValue_OK               ESteamNetworkingGetConfigValueResult = 1
	ESteamNetworkingGetConfigValue_OKInherited      ESteamNetworkingGetConfigValueResult = 2
	ESteamNetworkingGetConfigValue_ResultForce32Bit ESteamNetworkingGetConfigValueResult = 2147483647
)

type ESteamNetConnectionEnd int32

const (
	ESteamNetConnectionEnd_Invalid                       ESteamNetConnectionEnd = 0
	ESteamNetConnectionEnd_AppMin                        ESteamNetConnectionEnd = 1000
	ESteamNetConnectionEnd_AppGeneric                    ESteamNetConnectionEnd = 1000
	ESteamNetConnectionEnd_AppMax                        ESteamNetConnectionEnd = 1999
	ESteamNetConnectionEnd_AppExceptionMin               ESteamNetConnectionEnd = 2000
	ESteamNetConnectionEnd_AppExceptionGeneric           ESteamNetConnectionEnd = 2000
	ESteamNetConnectionEnd_AppExceptionMax               ESteamNetConnectionEnd = 2999
	ESteamNetConnectionEnd_LocalMin                      ESteamNetConnectionEnd = 3000
	ESteamNetConnectionEnd_LocalOfflineMode              ESteamNetConnectionEnd = 3001
	ESteamNetConnectionEnd_LocalManyRelayConnectivity    ESteamNetConnectionEnd = 3002
	ESteamNetConnectionEnd_LocalHostedServerPrimaryRelay ESteamNetConnectionEnd = 3003
	ESteamNetConnectionEnd_LocalNetworkConfig            ESteamNetConnectionEnd = 3004
	ESteamNetConnectionEnd_LocalRights                   ESteamNetConnectionEnd = 3005
	ESteamNetConnectionEnd_LocalP2PICENoPublicAddresses  ESteamNetConnectionEnd = 3006
	ESteamNetConnectionEnd_LocalMax                      ESteamNetConnectionEnd = 3999
	ESteamNetConnectionEnd_RemoteMin                     ESteamNetConnectionEnd = 4000
	ESteamNetConnectionEnd_RemoteTimeout                 ESteamNetConnectionEnd = 4001
	ESteamNetConnectionEnd_RemoteBadCrypt                ESteamNetConnectionEnd = 4002
	ESteamNetConnectionEnd_RemoteBadCert                 ESteamNetConnectionEnd = 4003
	ESteamNetConnectionEnd_RemoteBadProtocolVersion      ESteamNetConnectionEnd = 4006
	ESteamNetConnectionEnd_RemoteP2PICENoPublicAddresses ESteamNetConnectionEnd = 4007
	ESteamNetConnectionEnd_RemoteMax                     ESteamNetConnectionEnd = 4999
	ESteamNetConnectionEnd_MiscMin                       ESteamNetConnectionEnd = 5000
	ESteamNetConnectionEnd_MiscGeneric                   ESteamNetConnectionEnd = 5001
	ESteamNetConnectionEnd_MiscInternalError             ESteamNetConnectionEnd = 5002
	ESteamNetConnectionEnd_MiscTimeout                   ESteamNetConnectionEnd = 5003
	ESteamNetConnectionEnd_MiscSteamConnectivity         ESteamNetConnectionEnd = 5005
	ESteamNetConnectionEnd_MiscNoRelaySessionsToClient   ESteamNetConnectionEnd = 5006
	ESteamNetConnectionEnd_MiscP2PRendezvous             ESteamNetConnectionEnd = 5008
	ESteamNetConnectionEnd_MiscP2PNATFirewall            ESteamNetConnectionEnd = 5009
	ESteamNetConnectionEnd_MiscPeerSentNoConnection      ESteamNetConnectionEnd = 5010
	ESteamNetConnectionEnd_MiscMax                       ESteamNetConnectionEnd = 5999
	ESteamNetConnectionEnd_Force32Bit                    ESteamNetConnectionEnd = 2147483647
)

func (nm SteamNetworkingMessage) GetData() []byte {
	return unsafePointerToBytes(nm.Data, nm.Size)
}
func (nm SteamNetworkingMessage) String() string {
	return fmt.Sprintf("Data: %s, Size: %d, MessageNumber:  %d, TimeRecieved: %d\n", string(nm.GetData()), nm.Size, nm.MessageNumber, nm.TimeReceived)
}

func unsafePointerToBytes(ptr unsafe.Pointer, size int32) []byte {
	if ptr == nil || size <= 0 {
		return nil
	}

	// Using Go 1.17+ unsafe.Slice for safety and performance
	return unsafe.Slice((*byte)(ptr), size)
}

const (
	flatAPI_SteamIPAddress_t_IsSet = "SteamAPI_SteamIPAddress_t_IsSet"
)

var (
	steamIPAddress_IsSet func(*SteamIPAddress) bool
)

func (s *SteamIPAddress) IsSet() bool {
	return steamIPAddress_IsSet(s)
}

const (
	flatAPI_SteamNetworkingIPAddr_Clear            = "SteamAPI_SteamNetworkingIPAddr_Clear"
	flatAPI_SteamNetworkingIPAddr_IsIPv6AllZeros   = "SteamAPI_SteamNetworkingIPAddr_IsIPv6AllZeros"
	flatAPI_SteamNetworkingIPAddr_SetIPv6          = "SteamAPI_SteamNetworkingIPAddr_SetIPv6"
	flatAPI_SteamNetworkingIPAddr_SetIPv4          = "SteamAPI_SteamNetworkingIPAddr_SetIPv4"
	flatAPI_SteamNetworkingIPAddr_IsIPv4           = "SteamAPI_SteamNetworkingIPAddr_IsIPv4"
	flatAPI_SteamNetworkingIPAddr_GetIPv4          = "SteamAPI_SteamNetworkingIPAddr_GetIPv4"
	flatAPI_SteamNetworkingIPAddr_SetIPv6LocalHost = "SteamAPI_SteamNetworkingIPAddr_SetIPv6LocalHost"
	flatAPI_SteamNetworkingIPAddr_IsLocalHost      = "SteamAPI_SteamNetworkingIPAddr_IsLocalHost"
	flatAPI_SteamNetworkingIPAddr_ToString         = "SteamAPI_SteamNetworkingIPAddr_ToString"
	flatAPI_SteamNetworkingIPAddr_ParseString      = "SteamAPI_SteamNetworkingIPAddr_ParseString"
	flatAPI_SteamNetworkingIPAddr_IsEqualTo        = "SteamAPI_SteamNetworkingIPAddr_IsEqualTo"
	flatAPI_SteamNetworkingIPAddr_GetFakeIPType    = "SteamAPI_SteamNetworkingIPAddr_GetFakeIPType"
	flatAPI_SteamNetworkingIPAddr_IsFakeIP         = "SteamAPI_SteamNetworkingIPAddr_IsFakeIP"
)

var (
	steamNetworkingIPAddr_Clear            func(self *SteamNetworkingIPAddr)
	steamNetworkingIPAddr_IsIPv6AllZeros   func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_SetIPv6          func(self *SteamNetworkingIPAddr, ipv6 string, Port uint16)
	steamNetworkingIPAddr_SetIPv4          func(self *SteamNetworkingIPAddr, IP uint16, Port uint16)
	steamNetworkingIPAddr_IsIPv4           func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_GetIPv4          func(self *SteamNetworkingIPAddr) uint32
	steamNetworkingIPAddr_SetIPv6LocalHost func(self *SteamNetworkingIPAddr, Port uint16)
	steamNetworkingIPAddr_IsLocalHost      func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_ToString         func(self *SteamNetworkingIPAddr, buf string, cbBuf uint32, bWithPort bool)
	steamNetworkingIPAddr_ParseString      func(self *SteamNetworkingIPAddr, pszStr string) bool
	steamNetworkingIPAddr_IsEqualTo        func(self *SteamNetworkingIPAddr, x *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_GetFakeIPType    func(self *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType
	steamNetworkingIPAddr_IsFakeIP         func(self *SteamNetworkingIPAddr) bool
)

func (s *SteamNetworkingIPAddr) Clear() {
	steamNetworkingIPAddr_Clear(s)
}

func (s *SteamNetworkingIPAddr) IsIPv6AllZeros() bool {
	return steamNetworkingIPAddr_IsIPv6AllZeros(s)
}

func (s *SteamNetworkingIPAddr) SetIPv6(ipv6 string, Port uint16) {
	steamNetworkingIPAddr_SetIPv6(s, ipv6, Port)
}

func (s *SteamNetworkingIPAddr) SetIPv4(IP uint16, Port uint16) {
	steamNetworkingIPAddr_SetIPv4(s, IP, Port)
}

func (s *SteamNetworkingIPAddr) IsIPv4() bool {
	return steamNetworkingIPAddr_IsIPv4(s)
}

func (s *SteamNetworkingIPAddr) GetIPv4() uint32 {
	return steamNetworkingIPAddr_GetIPv4(s)
}

func (s *SteamNetworkingIPAddr) SetIPv6LocalHost(Port uint16) {
	steamNetworkingIPAddr_SetIPv6LocalHost(s, Port)
}

func (s *SteamNetworkingIPAddr) IsLocalHost() bool {
	return steamNetworkingIPAddr_IsLocalHost(s)
}

func (s *SteamNetworkingIPAddr) ToString(buf string, cbBuf uint32, bWithPort bool) {
	steamNetworkingIPAddr_ToString(s, buf, cbBuf, bWithPort)
}

func (s *SteamNetworkingIPAddr) ParseString(pszStr string) bool {
	return steamNetworkingIPAddr_ParseString(s, pszStr)
}

func (s *SteamNetworkingIPAddr) IsEqualTo(x *SteamNetworkingIPAddr) bool {
	return steamNetworkingIPAddr_IsEqualTo(s, x)
}

func (s *SteamNetworkingIPAddr) GetFakeIPType() ESteamNetworkingFakeIPType {
	return steamNetworkingIPAddr_GetFakeIPType(s)
}

func (s *SteamNetworkingIPAddr) IsFakeIP() bool {
	return steamNetworkingIPAddr_IsFakeIP(s)
}

const (
	flatAPI_SteamNetworkingIdentity_Clear             = "SteamAPI_SteamNetworkingIdentity_Clear"
	flatAPI_SteamNetworkingIdentity_IsInvalid         = "SteamAPI_SteamNetworkingIdentity_IsInvalid"
	flatAPI_SteamNetworkingIdentity_SetSteamID        = "SteamAPI_SteamNetworkingIdentity_SetSteamID"
	flatAPI_SteamNetworkingIdentity_GetSteamID        = "SteamAPI_SteamNetworkingIdentity_GetSteamID"
	flatAPI_SteamNetworkingIdentity_SetSteamID64      = "SteamAPI_SteamNetworkingIdentity_SetSteamID64"
	flatAPI_SteamNetworkingIdentity_GetSteamID64      = "SteamAPI_SteamNetworkingIdentity_GetSteamID64"
	flatAPI_SteamNetworkingIdentity_SetXboxPairwiseID = "SteamAPI_SteamNetworkingIdentity_SetXboxPairwiseID"
	flatAPI_SteamNetworkingIdentity_GetXboxPairwiseID = "SteamAPI_SteamNetworkingIdentity_GetXboxPairwiseID"
	flatAPI_SteamNetworkingIdentity_SetPSNID          = "SteamAPI_SteamNetworkingIdentity_SetPSNID"
	flatAPI_SteamNetworkingIdentity_GetPSNID          = "SteamAPI_SteamNetworkingIdentity_GetPSNID"
	flatAPI_SteamNetworkingIdentity_SetIPAddr         = "SteamAPI_SteamNetworkingIdentity_SetIPAddr"
	flatAPI_SteamNetworkingIdentity_GetIPAddr         = "SteamAPI_SteamNetworkingIdentity_GetIPAddr"
	flatAPI_SteamNetworkingIdentity_SetIPv4Addr       = "SteamAPI_SteamNetworkingIdentity_SetIPv4Addr"
	flatAPI_SteamNetworkingIdentity_GetIPv4           = "SteamAPI_SteamNetworkingIdentity_GetIPv4"
	flatAPI_SteamNetworkingIdentity_GetFakeIPType     = "SteamAPI_SteamNetworkingIdentity_GetFakeIPType"
	flatAPI_SteamNetworkingIdentity_IsFakeIP          = "SteamAPI_SteamNetworkingIdentity_IsFakeIP"
	flatAPI_SteamNetworkingIdentity_SetLocalHost      = "SteamAPI_SteamNetworkingIdentity_SetLocalHost"
	flatAPI_SteamNetworkingIdentity_IsLocalHost       = "SteamAPI_SteamNetworkingIdentity_IsLocalHost"
	flatAPI_SteamNetworkingIdentity_SetGenericString  = "SteamAPI_SteamNetworkingIdentity_SetGenericString"
	flatAPI_SteamNetworkingIdentity_GetGenericString  = "SteamAPI_SteamNetworkingIdentity_GetGenericString"
	flatAPI_SteamNetworkingIdentity_SetGenericBytes   = "SteamAPI_SteamNetworkingIdentity_SetGenericBytes"
	flatAPI_SteamNetworkingIdentity_GetGenericBytes   = "SteamAPI_SteamNetworkingIdentity_GetGenericBytes"
	flatAPI_SteamNetworkingIdentity_IsEqualTo         = "SteamAPI_SteamNetworkingIdentity_IsEqualTo"
	flatAPI_SteamNetworkingIdentity_ToString          = "SteamAPI_SteamNetworkingIdentity_ToString"
	flatAPI_SteamNetworkingIdentity_ParseString       = "SteamAPI_SteamNetworkingIdentity_ParseString"
)

var (
	steamNetworkingIdentity_Clear             func(self *SteamNetworkingIdentity)
	steamNetworkingIdentity_IsInvalid         func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetSteamID        func(self *SteamNetworkingIdentity, steamID Uint64SteamID)
	steamNetworkingIdentity_GetSteamID        func(self *SteamNetworkingIdentity) Uint64SteamID
	steamNetworkingIdentity_SetSteamID64      func(self *SteamNetworkingIdentity, steamID uint64)
	steamNetworkingIdentity_GetSteamID64      func(self *SteamNetworkingIdentity) uint64
	steamNetworkingIdentity_SetXboxPairwiseID func(self *SteamNetworkingIdentity, pszString string) bool
	steamNetworkingIdentity_GetXboxPairwiseID func(self *SteamNetworkingIdentity) string
	steamNetworkingIdentity_SetPSNID          func(self *SteamNetworkingIdentity, id uint64)
	steamNetworkingIdentity_GetPSNID          func(self *SteamNetworkingIdentity) uint64
	steamNetworkingIdentity_SetIPAddr         func(self *SteamNetworkingIdentity, addr *SteamNetworkingIPAddr)
	steamNetworkingIdentity_GetIPAddr         func(self *SteamNetworkingIdentity) *SteamNetworkingIPAddr
	steamNetworkingIdentity_SetIPv4Addr       func(self *SteamNetworkingIdentity, nIPv4u int32, nPort uint16)
	steamNetworkingIdentity_GetIPv4           func(self *SteamNetworkingIdentity) uint32
	steamNetworkingIdentity_GetFakeIPType     func(self *SteamNetworkingIdentity) ESteamNetworkingFakeIPType
	steamNetworkingIdentity_IsFakeIP          func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetLocalHost      func(self *SteamNetworkingIdentity)
	steamNetworkingIdentity_IsLocalHost       func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetGenericString  func(self *SteamNetworkingIdentity, pszString string) bool
	steamNetworkingIdentity_GetGenericString  func(self *SteamNetworkingIdentity) string
	steamNetworkingIdentity_SetGenericBytes   func(self *SteamNetworkingIdentity, data []byte, cbLen uint32) bool
	steamNetworkingIdentity_GetGenericBytes   func(self *SteamNetworkingIdentity, cbLen *int) string
	steamNetworkingIdentity_IsEqualTo         func(self *SteamNetworkingIdentity, x *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_ToString          func(self *SteamNetworkingIdentity, buf string, cbBuf uint32)
	steamNetworkingIdentity_ParseString       func(self *SteamNetworkingIdentity, pszStr string) bool
)

func (s *SteamNetworkingIdentity) Clear() {
	steamNetworkingIdentity_Clear(s)
}

func (s *SteamNetworkingIdentity) IsInvalid() bool {
	return steamNetworkingIdentity_IsInvalid(s)
}
func (s *SteamNetworkingIdentity) SetSteamID(steamID Uint64SteamID) {
	steamNetworkingIdentity_SetSteamID(s, steamID)
}
func (s *SteamNetworkingIdentity) GetSteamID() Uint64SteamID {
	return steamNetworkingIdentity_GetSteamID(s)
}
func (s *SteamNetworkingIdentity) SetSteamID64(steamID uint64) {
	steamNetworkingIdentity_SetSteamID64(s, steamID)
}

func (s *SteamNetworkingIdentity) GetSteamID64() uint64 {
	return steamNetworkingIdentity_GetSteamID64(s)
}
func (s *SteamNetworkingIdentity) SetXboxPairwiseID(String string) bool {
	return steamNetworkingIdentity_SetXboxPairwiseID(s, String)
}
func (s *SteamNetworkingIdentity) GetXboxPairwiseID() string {
	return steamNetworkingIdentity_GetXboxPairwiseID(s)
}
func (s *SteamNetworkingIdentity) SetPSNID(id uint64) {
	steamNetworkingIdentity_SetPSNID(s, id)
}
func (s *SteamNetworkingIdentity) GetPSNID() uint64 {
	return steamNetworkingIdentity_GetPSNID(s)
}

func (s *SteamNetworkingIdentity) SetIPAddr(addr *SteamNetworkingIPAddr) {
	steamNetworkingIdentity_SetIPAddr(s, addr)
}

func (s *SteamNetworkingIdentity) GetIPAddr() *SteamNetworkingIPAddr {
	return steamNetworkingIdentity_GetIPAddr(s)
}

func (s *SteamNetworkingIdentity) SetIPv4Addr(nIPv4u int32, nPort uint16) {
	steamNetworkingIdentity_SetIPv4Addr(s, nIPv4u, nPort)
}
func (s *SteamNetworkingIdentity) GetIPv4() uint32 {
	return steamNetworkingIdentity_GetIPv4(s)
}

func (s *SteamNetworkingIdentity) GetFakeIPType() ESteamNetworkingFakeIPType {
	return steamNetworkingIdentity_GetFakeIPType(s)
}
func (s *SteamNetworkingIdentity) IsFakeIP() bool {
	return steamNetworkingIdentity_IsFakeIP(s)
}
func (s *SteamNetworkingIdentity) SetLocalHost() {
	steamNetworkingIdentity_SetLocalHost(s)
}
func (s *SteamNetworkingIdentity) IsLocalHost() bool {
	return steamNetworkingIdentity_IsLocalHost(s)
}
func (s *SteamNetworkingIdentity) SetGenericString(pszString string) bool {
	return steamNetworkingIdentity_SetGenericString(s, pszString)
}

func (s *SteamNetworkingIdentity) GetGenericString() string {
	return steamNetworkingIdentity_GetGenericString(s)
}

func (s *SteamNetworkingIdentity) SetGenericBytes(data []byte) bool {
	return steamNetworkingIdentity_SetGenericBytes(s, data, uint32(len(data)))
}

func (s *SteamNetworkingIdentity) GetGenericBytes(cbLen *int) string {
	return steamNetworkingIdentity_GetGenericBytes(s, cbLen)
}

func (s *SteamNetworkingIdentity) IsEqualTo(x *SteamNetworkingIdentity) bool {
	return steamNetworkingIdentity_IsEqualTo(s, x)
}

func (s *SteamNetworkingIdentity) ToString(buf string, cbBuf uint32) {
	steamNetworkingIdentity_ToString(s, buf, cbBuf)
}

func (s *SteamNetworkingIdentity) ParseString(pszStr string) bool {
	return steamNetworkingIdentity_ParseString(s, pszStr)
}

const (
	flatAPI_SteamNetworkingConfigValue_t_SetInt32  = "SteamAPI_SteamNetworkingConfigValue_t_SetInt32"
	flatAPI_SteamNetworkingConfigValue_t_SetInt64  = "SteamAPI_SteamNetworkingConfigValue_t_SetInt64"
	flatAPI_SteamNetworkingConfigValue_t_SetFloat  = "SteamAPI_SteamNetworkingConfigValue_t_SetFloat"
	flatAPI_SteamNetworkingConfigValue_t_SetPtr    = "SteamAPI_SteamNetworkingConfigValue_t_SetPtr"
	flatAPI_SteamNetworkingConfigValue_t_SetString = "SteamAPI_SteamNetworkingConfigValue_t_SetString"
)

var (
	steamNetworkingConfigValue_t_SetInt32  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int32)
	steamNetworkingConfigValue_t_SetInt64  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int64)
	steamNetworkingConfigValue_t_SetFloat  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data float32)
	steamNetworkingConfigValue_t_SetPtr    func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data []byte)
	steamNetworkingConfigValue_t_SetString func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data string)
)

func (s *SteamNetworkingConfigValue) SetInt32(eVal ESteamNetworkingConfigValue, data int32) {
	steamNetworkingConfigValue_t_SetInt32(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetInt64(eVal ESteamNetworkingConfigValue, data int64) {
	steamNetworkingConfigValue_t_SetInt64(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetFloat(eVal ESteamNetworkingConfigValue, data float32) {
	steamNetworkingConfigValue_t_SetFloat(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetPtr(eVal ESteamNetworkingConfigValue, data []byte) {
	steamNetworkingConfigValue_t_SetPtr(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetString(eVal ESteamNetworkingConfigValue, data string) {
	steamNetworkingConfigValue_t_SetString(s, eVal, data)
}

const (
	flatAPI_SteamDatagramHostedAddress_Clear         = "SteamAPI_SteamDatagramHostedAddress_Clear"
	flatAPI_SteamDatagramHostedAddress_GetPopID      = "SteamAPI_SteamDatagramHostedAddress_GetPopID"
	flatAPI_SteamDatagramHostedAddress_SetDevAddress = "SteamAPI_SteamDatagramHostedAddress_SetDevAddress"
)

var (
	steamDatagramHostedAddress_Clear         func(self *SteamDatagramHostedAddress)
	steamDatagramHostedAddress_GetPopID      func(self *SteamDatagramHostedAddress) SteamNetworkingPOPID
	steamDatagramHostedAddress_SetDevAddress func(self *SteamDatagramHostedAddress, nIP uint32, nPort uint16, popid SteamNetworkingPOPID)
)

func (s *SteamDatagramHostedAddress) Clear() {
	steamDatagramHostedAddress_Clear(s)
}

func (s *SteamDatagramHostedAddress) GetPopID() SteamNetworkingPOPID {
	return steamDatagramHostedAddress_GetPopID(s)
}

func (s *SteamDatagramHostedAddress) SetDevAddress(nIP uint32, nPort uint16, popid SteamNetworkingPOPID) {
	steamDatagramHostedAddress_SetDevAddress(s, nIP, nPort, popid)
}

var (
	iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort  func()
	iSteamNetworkingFakeUDPPort_SendMessageToFakeIP func(remoteAddress *SteamNetworkingIPAddr, pData uintptr, cbData uint32, nSendFlags int32) EResult
	iSteamNetworkingFakeUDPPort_ReceiveMessages     func(ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingFakeUDPPort_ScheduleCleanup     func(remoteAddress *SteamNetworkingIPAddr)
)

const (
	flatAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort  = "SteamAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort"
	flatAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP = "SteamAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP"
	flatAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages     = "SteamAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages"
	flatAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup     = "SteamAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup"
)

func (s *SteamNetworkingFakeUDPPort) DestroyFakeUDPPort() {
	iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort()
}

func (s *SteamNetworkingFakeUDPPort) SendMessageToFakeIP(remoteAddress *SteamNetworkingIPAddr, pData uintptr, cbData uint32, nSendFlags int32) EResult {
	return iSteamNetworkingFakeUDPPort_SendMessageToFakeIP(remoteAddress, pData, cbData, nSendFlags)
}

func (s *SteamNetworkingFakeUDPPort) ReceiveMessages(nMaxMessages int32) []SteamNetworkingMessage {
	ppOutMessages := make([]SteamNetworkingMessage, 0, nMaxMessages)
	actual := iSteamNetworkingFakeUDPPort_ReceiveMessages(&ppOutMessages, nMaxMessages)
	return ppOutMessages[:actual]
}

func (s *SteamNetworkingFakeUDPPort) ScheduleCleanup(remoteAddress *SteamNetworkingIPAddr) {
	iSteamNetworkingFakeUDPPort_ScheduleCleanup(remoteAddress)
}

// Steam Networking Messages
const (
	flatAPI_SteamNetworkingMessages                           = "SteamAPI_SteamNetworkingMessages_SteamAPI_v002"
	flatAPI_ISteamNetworkingMessages_SendMessageToUser        = "SteamAPI_ISteamNetworkingMessages_SendMessageToUser"
	flatAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel = "SteamAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel"
	flatAPI_ISteamNetworkingMessages_AcceptSessionWithUser    = "SteamAPI_ISteamNetworkingMessages_AcceptSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseSessionWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseChannelWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseChannelWithUser"
	flatAPI_ISteamNetworkingMessages_GetSessionConnectionInfo = "SteamAPI_ISteamNetworkingMessages_GetSessionConnectionInfo"
)

var (
	steamNetworkingMessages_init                      func() uintptr
	iSteamNetworkingMessages_SendMessageToUser        func(steamNetworkingMessages uintptr, identityRemote uintptr, pubData []byte, cubData uint32, nSendFlags int32, nRemoteChannel int32) EResult
	iSteamNetworkingMessages_ReceiveMessagesOnChannel func(steamNetworkingMessages uintptr, nLocalChannel int32, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingMessages_AcceptSessionWithUser    func(steamNetworkingMessages uintptr, identityRemote uintptr) bool
	iSteamNetworkingMessages_CloseSessionWithUser     func(steamNetworkingMessages uintptr, identityRemote uintptr) bool
	iSteamNetworkingMessages_CloseChannelWithUser     func(steamNetworkingMessages uintptr, identityRemote uintptr, nLocalChannel int32) bool
	iSteamNetworkingMessages_GetSessionConnectionInfo func(steamNetworkingMessages uintptr, identityRemote uintptr, pConnectionInfo *SteamNetConnectionInfo, pQuickStatus *SteamNetConnectionRealTimeStatus) ESteamNetworkingConnectionState
)

type steamNetworkingMessages uintptr

func SteamNetworkingMessages() ISteamNetworkingMessages {
	return steamNetworkingMessages(steamNetworkingMessages_init())
}

func (s steamNetworkingMessages) SendMessageToUser(remoteIdentity SteamNetworkingIdentity, Data []byte, SendFlags int32, RemoteChannel int32) EResult {
	return iSteamNetworkingMessages_SendMessageToUser(uintptr(s), structToUintptr[SteamNetworkingIdentity](&remoteIdentity), Data, uint32(len(Data)), SendFlags, RemoteChannel)
}

func (s steamNetworkingMessages) ReceiveMessagesOnChannel(LocalChannel int32, MaxMessages int32) []SteamNetworkingMessage {
	ppOutMessages := make([]SteamNetworkingMessage, 0, MaxMessages)
	actual := iSteamNetworkingMessages_ReceiveMessagesOnChannel(uintptr(s), LocalChannel, &ppOutMessages, MaxMessages)
	return ppOutMessages[:actual]
}

func (s steamNetworkingMessages) AcceptSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool {
	return iSteamNetworkingMessages_AcceptSessionWithUser(uintptr(s), structToUintptr[SteamNetworkingIdentity](&remoteIdentity))
}

func (s steamNetworkingMessages) CloseSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool {
	return iSteamNetworkingMessages_CloseSessionWithUser(uintptr(s), structToUintptr[SteamNetworkingIdentity](&remoteIdentity))
}

func (s steamNetworkingMessages) CloseChannelWithUser(remoteIdentity SteamNetworkingIdentity, LocalChannel int32) bool {
	return iSteamNetworkingMessages_CloseChannelWithUser(uintptr(s), structToUintptr[SteamNetworkingIdentity](&remoteIdentity), LocalChannel)
}

func (s steamNetworkingMessages) GetSessionConnectionInfo(remoteIdentity SteamNetworkingIdentity) (ConnectionInfo SteamNetConnectionInfo, QuickStatus SteamNetConnectionRealTimeStatus, state ESteamNetworkingConnectionState) {
	state = iSteamNetworkingMessages_GetSessionConnectionInfo(uintptr(s), structToUintptr[SteamNetworkingIdentity](&remoteIdentity), &ConnectionInfo, &QuickStatus)
	return ConnectionInfo, QuickStatus, state
}

// Steam Networking Sockets

var (
	steamNetworkingSockets_init                                     func() uintptr
	iSteamNetworkingSockets_CreateListenSocketIP                    func(steamNetworkingSockets uintptr, localAddress uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_ConnectByIPAddress                      func(steamNetworkingSockets uintptr, address uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_CreateListenSocketP2P                   func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_ConnectP2P                              func(steamNetworkingSockets uintptr, identityRemote uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_AcceptConnection                        func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) EResult
	iSteamNetworkingSockets_CloseConnection                         func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nReason int32, pszDebug string, bEnableLinger bool) bool
	iSteamNetworkingSockets_CloseListenSocket                       func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket) bool
	iSteamNetworkingSockets_SetConnectionUserData                   func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nUserData int64) bool
	iSteamNetworkingSockets_GetConnectionUserData                   func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection) int64
	iSteamNetworkingSockets_SetConnectionName                       func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName string)
	iSteamNetworkingSockets_GetConnectionName                       func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName []byte, nMaxLen int32) bool
	iSteamNetworkingSockets_SendMessageToConnection                 func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pData []byte, cbData uint32, nSendFlags int32, pOutMessageNumber *int64) EResult
	iSteamNetworkingSockets_SendMessages                            func(steamNetworkingSockets uintptr, nMessages int32, pMessages []SteamNetworkingMessage, pOutMessageNumberOrResult []int64)
	iSteamNetworkingSockets_FlushMessagesOnConnection               func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) EResult
	iSteamNetworkingSockets_ReceiveMessagesOnConnection             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingSockets_GetConnectionInfo                       func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pInfo *SteamNetConnectionInfo) bool
	iSteamNetworkingSockets_GetConnectionRealTimeStatus             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pStatus *SteamNetConnectionRealTimeStatus, nLanes int32, pLanes []SteamNetConnectionRealTimeLaneStatus) EResult
	iSteamNetworkingSockets_GetDetailedConnectionStatus             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pszBuf []byte, cbBuf int32) int32
	iSteamNetworkingSockets_GetListenSocketAddress                  func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket, address *SteamNetworkingIPAddr) bool
	iSteamNetworkingSockets_CreateSocketPair                        func(steamNetworkingSockets uintptr, pOutConnection1 *HSteamNetConnection, pOutConnection2 *HSteamNetConnection, bUseNetworkLoopback bool, pIdentity1 *SteamNetworkingIdentity, pIdentity2 *SteamNetworkingIdentity) bool
	iSteamNetworkingSockets_ConfigureConnectionLanes                func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, nNumLanes int32, pLanePriorities []int32, pLaneWeights []uint16) EResult
	iSteamNetworkingSockets_GetIdentity                             func(steamNetworkingSockets uintptr, pIdentity *SteamNetworkingIdentity) bool
	iSteamNetworkingSockets_InitAuthentication                      func(steamNetworkingSockets uintptr) ESteamNetworkingAvailability
	iSteamNetworkingSockets_GetAuthenticationStatus                 func(steamNetworkingSockets uintptr, pDetails *SteamNetAuthenticationStatus) ESteamNetworkingAvailability
	iSteamNetworkingSockets_CreatePollGroup                         func(steamNetworkingSockets uintptr) HSteamNetPollGroup
	iSteamNetworkingSockets_DestroyPollGroup                        func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup) bool
	iSteamNetworkingSockets_SetConnectionPollGroup                  func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, hPollGroup HSteamNetPollGroup) bool
	iSteamNetworkingSockets_ReceiveMessagesOnPollGroup              func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingSockets_ReceivedRelayAuthTicket                 func(steamNetworkingSockets uintptr, pvTicket []byte, cbTicket int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) bool
	iSteamNetworkingSockets_FindRelayAuthTicketForServer            func(steamNetworkingSockets uintptr, identityGameServer uintptr, nRemoteVirtualPort int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) int32
	iSteamNetworkingSockets_ConnectToHostedDedicatedServer          func(steamNetworkingSockets uintptr, identityTarget uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_GetHostedDedicatedServerPort            func(steamNetworkingSockets uintptr) uint16
	iSteamNetworkingSockets_GetHostedDedicatedServerPOPID           func(steamNetworkingSockets uintptr) SteamNetworkingPOPID
	iSteamNetworkingSockets_GetHostedDedicatedServerAddress         func(steamNetworkingSockets uintptr, pRouting *SteamDatagramHostedAddress) EResult
	iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_GetGameCoordinatorServerLogin           func(steamNetworkingSockets uintptr, pLoginInfo *SteamDatagramGameCoordinatorServerLogin, pcbSignedBlob *int32, pBlob []byte) EResult
	iSteamNetworkingSockets_ConnectP2PCustomSignaling               func(steamNetworkingSockets uintptr, pSignaling *ISteamNetworkingConnectionSignaling, pPeerIdentity *SteamNetworkingIdentity, nRemoteVirtualPort int32, nOptions int32, pOptions *SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_ReceivedP2PCustomSignal                 func(steamNetworkingSockets uintptr, pMsg []byte, cbMsg int32, pContext *ISteamNetworkingSignalingRecvContext) bool
	iSteamNetworkingSockets_GetCertificateRequest                   func(steamNetworkingSockets uintptr, pcbBlob *int32, pBlob []byte, errMsg *SteamNetworkingErrMsg) bool
	iSteamNetworkingSockets_SetCertificate                          func(steamNetworkingSockets uintptr, pCertificate []byte, cbCertificate int32, errMsg *SteamNetworkingErrMsg) bool
	iSteamNetworkingSockets_ResetIdentity                           func(steamNetworkingSockets uintptr, pIdentity uintptr)
	iSteamNetworkingSockets_RunCallbacks                            func(steamNetworkingSockets uintptr)
	iSteamNetworkingSockets_BeginAsyncRequestFakeIP                 func(steamNetworkingSockets uintptr, nNumPorts int32) bool
	iSteamNetworkingSockets_GetFakeIP                               func(steamNetworkingSockets uintptr, idxFirstPort int32, pInfo *SteamNetworkingFakeIPResult)
	iSteamNetworkingSockets_CreateListenSocketP2PFakeIP             func(steamNetworkingSockets uintptr, idxFakePort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_GetRemoteFakeIPForConnection            func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pOutAddr *SteamNetworkingIPAddr) EResult
	iSteamNetworkingSockets_CreateFakeUDPPort                       func(steamNetworkingSockets uintptr, idxFakeServerPort int32) *SteamNetworkingFakeUDPPort
)

const (
	flatAPI_SteamNetworkingSockets                                          = "SteamAPI_SteamNetworkingSockets_SteamAPI_v012"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketIP                    = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketIP"
	flatAPI_ISteamNetworkingSockets_ConnectByIPAddress                      = "SteamAPI_ISteamNetworkingSockets_ConnectByIPAddress"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketP2P                   = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2P"
	flatAPI_ISteamNetworkingSockets_ConnectP2P                              = "SteamAPI_ISteamNetworkingSockets_ConnectP2P"
	flatAPI_ISteamNetworkingSockets_AcceptConnection                        = "SteamAPI_ISteamNetworkingSockets_AcceptConnection"
	flatAPI_ISteamNetworkingSockets_CloseConnection                         = "SteamAPI_ISteamNetworkingSockets_CloseConnection"
	flatAPI_ISteamNetworkingSockets_CloseListenSocket                       = "SteamAPI_ISteamNetworkingSockets_CloseListenSocket"
	flatAPI_ISteamNetworkingSockets_SetConnectionUserData                   = "SteamAPI_ISteamNetworkingSockets_SetConnectionUserData"
	flatAPI_ISteamNetworkingSockets_GetConnectionUserData                   = "SteamAPI_ISteamNetworkingSockets_GetConnectionUserData"
	flatAPI_ISteamNetworkingSockets_SetConnectionName                       = "SteamAPI_ISteamNetworkingSockets_SetConnectionName"
	flatAPI_ISteamNetworkingSockets_GetConnectionName                       = "SteamAPI_ISteamNetworkingSockets_GetConnectionName"
	flatAPI_ISteamNetworkingSockets_SendMessageToConnection                 = "SteamAPI_ISteamNetworkingSockets_SendMessageToConnection"
	flatAPI_ISteamNetworkingSockets_SendMessages                            = "SteamAPI_ISteamNetworkingSockets_SendMessages"
	flatAPI_ISteamNetworkingSockets_FlushMessagesOnConnection               = "SteamAPI_ISteamNetworkingSockets_FlushMessagesOnConnection"
	flatAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection             = "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection"
	flatAPI_ISteamNetworkingSockets_GetConnectionInfo                       = "SteamAPI_ISteamNetworkingSockets_GetConnectionInfo"
	flatAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus             = "SteamAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus"
	flatAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus             = "SteamAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus"
	flatAPI_ISteamNetworkingSockets_GetListenSocketAddress                  = "SteamAPI_ISteamNetworkingSockets_GetListenSocketAddress"
	flatAPI_ISteamNetworkingSockets_CreateSocketPair                        = "SteamAPI_ISteamNetworkingSockets_CreateSocketPair"
	flatAPI_ISteamNetworkingSockets_ConfigureConnectionLanes                = "SteamAPI_ISteamNetworkingSockets_ConfigureConnectionLanes"
	flatAPI_ISteamNetworkingSockets_GetIdentity                             = "SteamAPI_ISteamNetworkingSockets_GetIdentity"
	flatAPI_ISteamNetworkingSockets_InitAuthentication                      = "SteamAPI_ISteamNetworkingSockets_InitAuthentication"
	flatAPI_ISteamNetworkingSockets_GetAuthenticationStatus                 = "SteamAPI_ISteamNetworkingSockets_GetAuthenticationStatus"
	flatAPI_ISteamNetworkingSockets_CreatePollGroup                         = "SteamAPI_ISteamNetworkingSockets_CreatePollGroup"
	flatAPI_ISteamNetworkingSockets_DestroyPollGroup                        = "SteamAPI_ISteamNetworkingSockets_DestroyPollGroup"
	flatAPI_ISteamNetworkingSockets_SetConnectionPollGroup                  = "SteamAPI_ISteamNetworkingSockets_SetConnectionPollGroup"
	flatAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup              = "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup"
	flatAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket                 = "SteamAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket"
	flatAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer            = "SteamAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer"
	flatAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer          = "SteamAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort            = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID           = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress         = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress"
	flatAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket = "SteamAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket"
	flatAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin           = "SteamAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin"
	flatAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling               = "SteamAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling"
	flatAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal                 = "SteamAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal"
	flatAPI_ISteamNetworkingSockets_GetCertificateRequest                   = "SteamAPI_ISteamNetworkingSockets_GetCertificateRequest"
	flatAPI_ISteamNetworkingSockets_SetCertificate                          = "SteamAPI_ISteamNetworkingSockets_SetCertificate"
	flatAPI_ISteamNetworkingSockets_ResetIdentity                           = "SteamAPI_ISteamNetworkingSockets_ResetIdentity"
	flatAPI_ISteamNetworkingSockets_RunCallbacks                            = "SteamAPI_ISteamNetworkingSockets_RunCallbacks"
	flatAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP                 = "SteamAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP"
	flatAPI_ISteamNetworkingSockets_GetFakeIP                               = "SteamAPI_ISteamNetworkingSockets_GetFakeIP"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP             = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP"
	flatAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection            = "SteamAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection"
	flatAPI_ISteamNetworkingSockets_CreateFakeUDPPort                       = "SteamAPI_ISteamNetworkingSockets_CreateFakeUDPPort"
)

type steamNetworkingSockets uintptr

func SteamNetworkingSockets() ISteamNetworkingSockets {
	return steamNetworkingSockets(steamNetworkingSockets_init())
}

func (s steamNetworkingSockets) CreateListenSocketIP(localAddress SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketIP(uintptr(s), structToUintptr[SteamNetworkingIPAddr](&localAddress), int32(len(Options)), Options)
}

func (s steamNetworkingSockets) ConnectByIPAddress(address SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectByIPAddress(uintptr(s), structToUintptr[SteamNetworkingIPAddr](&address), int32(len(Options)), Options)
}

func (s steamNetworkingSockets) CreateListenSocketP2P(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketP2P(uintptr(s), LocalVirtualPort, int32(len(Options)), Options)
}

func (s steamNetworkingSockets) ConnectP2P(identityRemote SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectP2P(uintptr(s), structToUintptr[SteamNetworkingIdentity](&identityRemote), RemoteVirtualPort, int32(len(Options)), Options)
}

func (s steamNetworkingSockets) AcceptConnection(Conn HSteamNetConnection) EResult {
	return iSteamNetworkingSockets_AcceptConnection(uintptr(s), Conn)
}

func (s steamNetworkingSockets) CloseConnection(Peer HSteamNetConnection, Reason int32, Debug string, EnableLinger bool) bool {
	return iSteamNetworkingSockets_CloseConnection(uintptr(s), Peer, Reason, Debug, EnableLinger)
}

func (s steamNetworkingSockets) CloseListenSocket(Socket HSteamListenSocket) bool {
	return iSteamNetworkingSockets_CloseListenSocket(uintptr(s), Socket)
}

func (s steamNetworkingSockets) SetConnectionUserData(Peer HSteamNetConnection, UserData int64) bool {
	return iSteamNetworkingSockets_SetConnectionUserData(uintptr(s), Peer, UserData)
}

func (s steamNetworkingSockets) GetConnectionUserData(Peer HSteamNetConnection) int64 {
	return iSteamNetworkingSockets_GetConnectionUserData(uintptr(s), Peer)
}

func (s steamNetworkingSockets) SetConnectionName(Peer HSteamNetConnection, Name string) {
	iSteamNetworkingSockets_SetConnectionName(uintptr(s), Peer, Name)
}

func (s steamNetworkingSockets) GetConnectionName(Peer HSteamNetConnection, MaxLen int32) (name string, success bool) {
	var pszName []byte = make([]byte, 0, MaxLen)
	success = iSteamNetworkingSockets_GetConnectionName(uintptr(s), Peer, pszName, MaxLen)
	return string(pszName), success
}

func (s steamNetworkingSockets) SendMessageToConnection(Conn HSteamNetConnection, Data []byte, SendFlags int32) (OutMessageNumber int64, result EResult) {
	result = iSteamNetworkingSockets_SendMessageToConnection(uintptr(s), Conn, Data, uint32(len(Data)), SendFlags, &OutMessageNumber)
	return OutMessageNumber, result
}

// To use this
// / function, you must first allocate a message object using
// / ISteamNetworkingUtils::AllocateMessage.  (Do not declare one
// / on the stack or allocate your own.)
func (s steamNetworkingSockets) SendMessages(numMessages int32, Messages []SteamNetworkingMessage) (OutMessageNumberOrResult []int64) {
	OutMessageNumberOrResult = make([]int64, 0, len(Messages))
	iSteamNetworkingSockets_SendMessages(uintptr(s), numMessages, Messages, OutMessageNumberOrResult)
	return OutMessageNumberOrResult
}

func (s steamNetworkingSockets) FlushMessagesOnConnection(Conn HSteamNetConnection) EResult {
	return iSteamNetworkingSockets_FlushMessagesOnConnection(uintptr(s), Conn)
}

func (s steamNetworkingSockets) ReceiveMessagesOnConnection(Conn HSteamNetConnection, MaxMessages int32) []SteamNetworkingMessage {
	var ppOutMessages []SteamNetworkingMessage = make([]SteamNetworkingMessage, 0, MaxMessages)
	result := iSteamNetworkingSockets_ReceiveMessagesOnConnection(uintptr(s), Conn, &ppOutMessages, MaxMessages)
	return ppOutMessages[:result]
}

func (s steamNetworkingSockets) GetConnectionInfo(Conn HSteamNetConnection) (Info SteamNetConnectionInfo, success bool) {
	success = iSteamNetworkingSockets_GetConnectionInfo(uintptr(s), Conn, &Info)
	return Info, success
}

func (s steamNetworkingSockets) GetConnectionRealTimeStatus(Conn HSteamNetConnection, numLanes int32) (Status SteamNetConnectionRealTimeStatus, Lanes []SteamNetConnectionRealTimeLaneStatus, result EResult) {
	Lanes = make([]SteamNetConnectionRealTimeLaneStatus, 0, numLanes)
	result = iSteamNetworkingSockets_GetConnectionRealTimeStatus(uintptr(s), Conn, &Status, numLanes, Lanes)
	return Status, Lanes, result
}

func (s steamNetworkingSockets) GetDetailedConnectionStatus(Conn HSteamNetConnection, bufSize int32) (status string, result int32) {
	var pszBuf []byte = make([]byte, 0, bufSize)
	result = iSteamNetworkingSockets_GetDetailedConnectionStatus(uintptr(s), Conn, pszBuf, bufSize)
	return string(pszBuf), result
}

func (s steamNetworkingSockets) GetListenSocketAddress(Socket HSteamListenSocket) (address SteamNetworkingIPAddr, success bool) {
	success = iSteamNetworkingSockets_GetListenSocketAddress(uintptr(s), Socket, &address)
	return address, success
}

func (s steamNetworkingSockets) CreateSocketPair(UseNetworkLoopback bool) (OutConnection1 HSteamNetConnection, OutConnection2 HSteamNetConnection, Identity1 SteamNetworkingIdentity, Identity2 SteamNetworkingIdentity, succes bool) {
	succes = iSteamNetworkingSockets_CreateSocketPair(uintptr(s), &OutConnection1, &OutConnection2, UseNetworkLoopback, &Identity1, &Identity2)
	return OutConnection1, OutConnection2, Identity1, Identity2, succes
}

func (s steamNetworkingSockets) ConfigureConnectionLanes(Conn HSteamNetConnection, NumLanes int32, LanePriorities []int32, LaneWeights []uint16) EResult {
	return iSteamNetworkingSockets_ConfigureConnectionLanes(uintptr(s), Conn, NumLanes, LanePriorities, LaneWeights)
}

func (s steamNetworkingSockets) GetIdentity() (Identity SteamNetworkingIdentity, success bool) {
	success = iSteamNetworkingSockets_GetIdentity(uintptr(s), &Identity)
	return Identity, success
}

func (s steamNetworkingSockets) InitAuthentication() ESteamNetworkingAvailability {
	return iSteamNetworkingSockets_InitAuthentication(uintptr(s))
}

func (s steamNetworkingSockets) GetAuthenticationStatus() (Details SteamNetAuthenticationStatus, availability ESteamNetworkingAvailability) {
	availability = iSteamNetworkingSockets_GetAuthenticationStatus(uintptr(s), &Details)
	return Details, availability
}

func (s steamNetworkingSockets) CreatePollGroup() HSteamNetPollGroup {
	return iSteamNetworkingSockets_CreatePollGroup(uintptr(s))
}

func (s steamNetworkingSockets) DestroyPollGroup(PollGroup HSteamNetPollGroup) bool {
	return iSteamNetworkingSockets_DestroyPollGroup(uintptr(s), PollGroup)
}

func (s steamNetworkingSockets) SetConnectionPollGroup(Conn HSteamNetConnection, PollGroup HSteamNetPollGroup) bool {
	return iSteamNetworkingSockets_SetConnectionPollGroup(uintptr(s), Conn, PollGroup)
}

func (s steamNetworkingSockets) ReceiveMessagesOnPollGroup(PollGroup HSteamNetPollGroup, MaxMessages int32) []SteamNetworkingMessage {
	var ppOutMessages []SteamNetworkingMessage = make([]SteamNetworkingMessage, 0, MaxMessages)
	result := iSteamNetworkingSockets_ReceiveMessagesOnPollGroup(uintptr(s), PollGroup, &ppOutMessages, MaxMessages)
	return ppOutMessages[:result]
}

func (s steamNetworkingSockets) ReceivedRelayAuthTicket(pvTicket []byte) (pOutParsedTicket SteamDatagramRelayAuthTicket, success bool) {
	success = iSteamNetworkingSockets_ReceivedRelayAuthTicket(uintptr(s), pvTicket, int32(len(pvTicket)), &pOutParsedTicket)
	return pOutParsedTicket, success
}

func (s steamNetworkingSockets) FindRelayAuthTicketForServer(gameServerIdentity SteamNetworkingIdentity, RemoteVirtualPort int32) (OutParsedTicket SteamDatagramRelayAuthTicket, secondsToExpire int32) {
	secondsToExpire = iSteamNetworkingSockets_FindRelayAuthTicketForServer(uintptr(s), structToUintptr[SteamNetworkingIdentity](&gameServerIdentity), RemoteVirtualPort, &OutParsedTicket)
	return OutParsedTicket, secondsToExpire
}

func (s steamNetworkingSockets) ConnectToHostedDedicatedServer(targetIdentity SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectToHostedDedicatedServer(uintptr(s), structToUintptr[SteamNetworkingIdentity](&targetIdentity), RemoteVirtualPort, int32(len(Options)), Options)
}

func (s steamNetworkingSockets) GetHostedDedicatedServerPort() uint16 {
	return iSteamNetworkingSockets_GetHostedDedicatedServerPort(uintptr(s))
}

func (s steamNetworkingSockets) GetHostedDedicatedServerPOPID() SteamNetworkingPOPID {
	return iSteamNetworkingSockets_GetHostedDedicatedServerPOPID(uintptr(s))
}

func (s steamNetworkingSockets) GetHostedDedicatedServerAddress() (Routing SteamDatagramHostedAddress, result EResult) {
	result = iSteamNetworkingSockets_GetHostedDedicatedServerAddress(uintptr(s), &Routing)
	return Routing, result
}

func (s steamNetworkingSockets) CreateHostedDedicatedServerListenSocket(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket(uintptr(s), LocalVirtualPort, int32(len(Options)), Options)
}

func (s steamNetworkingSockets) GetGameCoordinatorServerLogin(LoginInfo *SteamDatagramGameCoordinatorServerLogin, SignedBlob int32) (Blob []byte, result EResult) {
	Blob = make([]byte, 0, SignedBlob)
	result = iSteamNetworkingSockets_GetGameCoordinatorServerLogin(uintptr(s), LoginInfo, &SignedBlob, Blob)
	return Blob[:SignedBlob], result
}

// func (s steamNetworkingSockets)  ConnectP2PCustomSignaling(pSignaling *ISteamNetworkingConnectionSignaling, pPeerIdentity *SteamNetworkingIdentity, nRemoteVirtualPort int32, nOptions int32, pOptions *SteamNetworkingConfigValue) HSteamNetConnection {
// 	return iSteamNetworkingSockets_ConnectP2PCustomSignaling(uintptr(s), pSignaling, pPeerIdentity, nRemoteVirtualPort, nOptions, pOptions)
// }

// func (s steamNetworkingSockets)  ReceivedP2PCustomSignal(cbMsg int32, pContext *ISteamNetworkingSignalingRecvContext) (pMsg []byte, success bool) {
// 	pMsg = make([]byte, 0, cbMsg)
// 	success = iSteamNetworkingSockets_ReceivedP2PCustomSignal(uintptr(s), pMsg, cbMsg, pContext)
// 	return pMsg, success
// }

func (s steamNetworkingSockets) GetCertificateRequest(BlobSize int32) (Blob []byte, errMsg SteamNetworkingErrMsg, success bool) {
	Blob = make([]byte, 0, BlobSize)
	success = iSteamNetworkingSockets_GetCertificateRequest(uintptr(s), &BlobSize, Blob, &errMsg)
	return Blob[:BlobSize], errMsg, success
}

func (s steamNetworkingSockets) SetCertificate(Certificate []byte) (errMsg SteamNetworkingErrMsg, success bool) {
	success = iSteamNetworkingSockets_SetCertificate(uintptr(s), Certificate, int32(len(Certificate)), &errMsg)
	return errMsg, success
}

/// NOTE: This function is not actually supported on Steam!  It is included
///       for use on other platforms where the active user can sign out and
///       a new user can sign in.

func (s steamNetworkingSockets) ResetIdentity(Identity SteamNetworkingIdentity) {
	iSteamNetworkingSockets_ResetIdentity(uintptr(s), structToUintptr[SteamNetworkingIdentity](&Identity))
}

func (s steamNetworkingSockets) RunCallbacks() {
	iSteamNetworkingSockets_RunCallbacks(uintptr(s))
}

func (s steamNetworkingSockets) BeginAsyncRequestFakeIP(NumPorts int32) bool {
	return iSteamNetworkingSockets_BeginAsyncRequestFakeIP(uintptr(s), NumPorts)
}

func (s steamNetworkingSockets) GetFakeIP(FirstPortIdx int32) (Info SteamNetworkingFakeIPResult) {
	iSteamNetworkingSockets_GetFakeIP(uintptr(s), FirstPortIdx, &Info)
	return Info
}

func (s steamNetworkingSockets) CreateListenSocketP2PFakeIP(FakePortIdx int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketP2PFakeIP(uintptr(s), FakePortIdx, int32(len(Options)), Options)
}

func (s steamNetworkingSockets) GetRemoteFakeIPForConnection(Conn HSteamNetConnection) (OutAddr SteamNetworkingIPAddr, result EResult) {
	result = iSteamNetworkingSockets_GetRemoteFakeIPForConnection(uintptr(s), Conn, &OutAddr)
	return OutAddr, result
}

func (s steamNetworkingSockets) CreateFakeUDPPort(FakeServerPortIdx int32) *SteamNetworkingFakeUDPPort {
	return iSteamNetworkingSockets_CreateFakeUDPPort(uintptr(s), FakeServerPortIdx)
}

// Steam Networking Utils
type intptr int64
type size uint64

var (
	steamNetworkingUtils_init                                                   func() uintptr
	iSteamNetworkingUtils_AllocateMessage                                       func(steamNetworkingUtils uintptr, cbAllocateBuffer int32) *SteamNetworkingMessage
	iSteamNetworkingUtils_InitRelayNetworkAccess                                func(steamNetworkingUtils uintptr)
	iSteamNetworkingUtils_GetRelayNetworkStatus                                 func(steamNetworkingUtils uintptr, pDetails *SteamRelayNetworkStatus) ESteamNetworkingAvailability
	iSteamNetworkingUtils_GetLocalPingLocation                                  func(steamNetworkingUtils uintptr, result *SteamNetworkPingLocation) float32
	iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations                   func(steamNetworkingUtils uintptr, location1 uintptr, location2 uintptr) int32
	iSteamNetworkingUtils_EstimatePingTimeFromLocalHost                         func(steamNetworkingUtils uintptr, remoteLocation uintptr) int32
	iSteamNetworkingUtils_ConvertPingLocationToString                           func(steamNetworkingUtils uintptr, location uintptr, pszBuf []byte, cchBufSize int32)
	iSteamNetworkingUtils_ParsePingLocationString                               func(steamNetworkingUtils uintptr, pszString string, result *SteamNetworkPingLocation) bool
	iSteamNetworkingUtils_CheckPingDataUpToDate                                 func(steamNetworkingUtils uintptr, flMaxAgeSeconds float32) bool
	iSteamNetworkingUtils_GetPingToDataCenter                                   func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID, pViaRelayPoP *SteamNetworkingPOPID) int32
	iSteamNetworkingUtils_GetDirectPingToPOP                                    func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID) int32
	iSteamNetworkingUtils_GetPOPCount                                           func(steamNetworkingUtils uintptr) int32
	iSteamNetworkingUtils_GetPOPList                                            func(steamNetworkingUtils uintptr, list []SteamNetworkingPOPID, nListSz int32) int32
	iSteamNetworkingUtils_GetLocalTimestamp                                     func(steamNetworkingUtils uintptr) SteamNetworkingMicroseconds
	iSteamNetworkingUtils_SetDebugOutputFunction                                func(steamNetworkingUtils uintptr, eDetailLevel ESteamNetworkingSocketsDebugOutputType, pfnFunc FSteamNetworkingSocketsDebugOutput)
	iSteamNetworkingUtils_IsFakeIPv4                                            func(steamNetworkingUtils uintptr, nIPv4 uint32) bool
	iSteamNetworkingUtils_GetIPv4FakeIPType                                     func(steamNetworkingUtils uintptr, nIPv4 uint32) ESteamNetworkingFakeIPType
	iSteamNetworkingUtils_GetRealIdentityForFakeIP                              func(steamNetworkingUtils uintptr, fakeIP uintptr, pOutRealIdentity *SteamNetworkingIdentity) EResult
	iSteamNetworkingUtils_SetGlobalConfigValueInt32                             func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val int32) bool
	iSteamNetworkingUtils_SetGlobalConfigValueFloat                             func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val float32) bool
	iSteamNetworkingUtils_SetGlobalConfigValueString                            func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val string) bool
	iSteamNetworkingUtils_SetGlobalConfigValuePtr                               func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val []byte) bool
	iSteamNetworkingUtils_SetConnectionConfigValueInt32                         func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool
	iSteamNetworkingUtils_SetConnectionConfigValueFloat                         func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val float32) bool
	iSteamNetworkingUtils_SetConnectionConfigValueString                        func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged     func(steamNetworkingUtils uintptr, fnCallback FnSteamNetConnectionStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged func(steamNetworkingUtils uintptr, fnCallback FnSteamNetAuthenticationStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged      func(steamNetworkingUtils uintptr, fnCallback FnSteamRelayNetworkStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult                        func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingFakeIPResult) bool
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest              func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionRequest) bool
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed               func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionFailed) bool
	iSteamNetworkingUtils_SetConfigValue                                        func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool
	iSteamNetworkingUtils_SetConfigValueStruct                                  func(steamNetworkingUtils uintptr, opt uintptr, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool
	iSteamNetworkingUtils_GetConfigValue                                        func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, pOutDataType *ESteamNetworkingConfigDataType, pResult []byte, cbResult *size) ESteamNetworkingGetConfigValueResult
	iSteamNetworkingUtils_GetConfigValueInfo                                    func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, pOutDataType *ESteamNetworkingConfigDataType, pOutScope *ESteamNetworkingConfigScope) []byte
	iSteamNetworkingUtils_IterateGenericEditableConfigValues                    func(steamNetworkingUtils uintptr, eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString                        func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr, buf []byte, cbBuf uint32, bWithPort bool)
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString                     func(steamNetworkingUtils uintptr, pAddr *SteamNetworkingIPAddr, pszStr string) bool
	iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType                   func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType
	iSteamNetworkingUtils_SteamNetworkingIdentity_ToString                      func(steamNetworkingUtils uintptr, identity *SteamNetworkingIdentity, buf []byte, cbBuf uint32)
	iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString                   func(steamNetworkingUtils uintptr, pIdentity *SteamNetworkingIdentity, pszStr string) bool
)

const (
	flatAPI_SteamNetworkingUtils                                                        = "SteamAPI_SteamNetworkingUtils_SteamAPI_v004"
	flatAPI_ISteamNetworkingUtils_AllocateMessage                                       = "SteamAPI_ISteamNetworkingUtils_AllocateMessage"
	flatAPI_ISteamNetworkingUtils_InitRelayNetworkAccess                                = "SteamAPI_ISteamNetworkingUtils_InitRelayNetworkAccess"
	flatAPI_ISteamNetworkingUtils_GetRelayNetworkStatus                                 = "SteamAPI_ISteamNetworkingUtils_GetRelayNetworkStatus"
	flatAPI_ISteamNetworkingUtils_GetLocalPingLocation                                  = "SteamAPI_ISteamNetworkingUtils_GetLocalPingLocation"
	flatAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations                   = "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations"
	flatAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost                         = "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost"
	flatAPI_ISteamNetworkingUtils_ConvertPingLocationToString                           = "SteamAPI_ISteamNetworkingUtils_ConvertPingLocationToString"
	flatAPI_ISteamNetworkingUtils_ParsePingLocationString                               = "SteamAPI_ISteamNetworkingUtils_ParsePingLocationString"
	flatAPI_ISteamNetworkingUtils_CheckPingDataUpToDate                                 = "SteamAPI_ISteamNetworkingUtils_CheckPingDataUpToDate"
	flatAPI_ISteamNetworkingUtils_GetPingToDataCenter                                   = "SteamAPI_ISteamNetworkingUtils_GetPingToDataCenter"
	flatAPI_ISteamNetworkingUtils_GetDirectPingToPOP                                    = "SteamAPI_ISteamNetworkingUtils_GetDirectPingToPOP"
	flatAPI_ISteamNetworkingUtils_GetPOPCount                                           = "SteamAPI_ISteamNetworkingUtils_GetPOPCount"
	flatAPI_ISteamNetworkingUtils_GetPOPList                                            = "SteamAPI_ISteamNetworkingUtils_GetPOPList"
	flatAPI_ISteamNetworkingUtils_GetLocalTimestamp                                     = "SteamAPI_ISteamNetworkingUtils_GetLocalTimestamp"
	flatAPI_ISteamNetworkingUtils_SetDebugOutputFunction                                = "SteamAPI_ISteamNetworkingUtils_SetDebugOutputFunction"
	flatAPI_ISteamNetworkingUtils_IsFakeIPv4                                            = "SteamAPI_ISteamNetworkingUtils_IsFakeIPv4"
	flatAPI_ISteamNetworkingUtils_GetIPv4FakeIPType                                     = "SteamAPI_ISteamNetworkingUtils_GetIPv4FakeIPType"
	flatAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP                              = "SteamAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32                             = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat                             = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueString                            = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueString"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr                               = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32                         = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat                         = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueString                        = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueString"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged     = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged      = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult                        = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest              = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed               = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed"
	flatAPI_ISteamNetworkingUtils_SetConfigValue                                        = "SteamAPI_ISteamNetworkingUtils_SetConfigValue"
	flatAPI_ISteamNetworkingUtils_SetConfigValueStruct                                  = "SteamAPI_ISteamNetworkingUtils_SetConfigValueStruct"
	flatAPI_ISteamNetworkingUtils_GetConfigValue                                        = "SteamAPI_ISteamNetworkingUtils_GetConfigValue"
	flatAPI_ISteamNetworkingUtils_GetConfigValueInfo                                    = "SteamAPI_ISteamNetworkingUtils_GetConfigValueInfo"
	flatAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues                    = "SteamAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString                        = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString                     = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType                   = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString                      = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString                   = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString"
)

type steamNetworkingUtils uintptr

func SteamNetworkingUtils() ISteamNetworkingUtils {
	return steamNetworkingUtils(steamNetworkingUtils_init())
}

func (s steamNetworkingUtils) AllocateMessage(AllocateBufferSize int32) *SteamNetworkingMessage {
	return iSteamNetworkingUtils_AllocateMessage(uintptr(s), AllocateBufferSize)
}

func (s steamNetworkingUtils) InitRelayNetworkAccess() {
	iSteamNetworkingUtils_InitRelayNetworkAccess(uintptr(s))
}

func (s steamNetworkingUtils) GetRelayNetworkStatus() (Details SteamRelayNetworkStatus, availability ESteamNetworkingAvailability) {
	availability = iSteamNetworkingUtils_GetRelayNetworkStatus(uintptr(s), &Details)
	return Details, availability
}

func (s steamNetworkingUtils) GetLocalPingLocation() (result SteamNetworkPingLocation, ageInSeconds float32) {
	ageInSeconds = iSteamNetworkingUtils_GetLocalPingLocation(uintptr(s), &result)
	return result, ageInSeconds
}

func (s steamNetworkingUtils) EstimatePingTimeBetweenTwoLocations(location1 SteamNetworkPingLocation, location2 SteamNetworkPingLocation) int32 {
	return iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations(uintptr(s), structToUintptr[SteamNetworkPingLocation](&location1), structToUintptr[SteamNetworkPingLocation](&location2))
}

func (s steamNetworkingUtils) EstimatePingTimeFromLocalHost(remoteLocation SteamNetworkPingLocation) int32 {
	return iSteamNetworkingUtils_EstimatePingTimeFromLocalHost(uintptr(s), structToUintptr[SteamNetworkPingLocation](&remoteLocation))
}

func (s steamNetworkingUtils) ConvertPingLocationToString(location SteamNetworkPingLocation, BufSize int32) string {
	var pszBuf []byte = make([]byte, 0, BufSize)
	iSteamNetworkingUtils_ConvertPingLocationToString(uintptr(s), structToUintptr[SteamNetworkPingLocation](&location), pszBuf, BufSize)
	return string(pszBuf)
}

func (s steamNetworkingUtils) ParsePingLocationString(String string) (result SteamNetworkPingLocation, success bool) {
	success = iSteamNetworkingUtils_ParsePingLocationString(uintptr(s), String, &result)
	return result, success
}

func (s steamNetworkingUtils) CheckPingDataUpToDate(MaxAgeSeconds float32) bool {
	return iSteamNetworkingUtils_CheckPingDataUpToDate(uintptr(s), MaxAgeSeconds)
}

func (s steamNetworkingUtils) GetPingToDataCenter(popID SteamNetworkingPOPID) (ViaRelayPoP SteamNetworkingPOPID, pingTime int32) {
	pingTime = iSteamNetworkingUtils_GetPingToDataCenter(uintptr(s), popID, &ViaRelayPoP)
	return ViaRelayPoP, pingTime
}

func (s steamNetworkingUtils) GetDirectPingToPOP(popID SteamNetworkingPOPID) int32 {
	return iSteamNetworkingUtils_GetDirectPingToPOP(uintptr(s), popID)
}

func (s steamNetworkingUtils) GetPOPCount() int32 {
	return iSteamNetworkingUtils_GetPOPCount(uintptr(s))
}

func (s steamNetworkingUtils) GetPOPList(ListSz int32) (list []SteamNetworkingPOPID) {
	list = make([]SteamNetworkingPOPID, 0, ListSz)
	actual := iSteamNetworkingUtils_GetPOPList(uintptr(s), list, ListSz)
	return list[:actual]
}

func (s steamNetworkingUtils) GetLocalTimestamp() SteamNetworkingMicroseconds {
	return iSteamNetworkingUtils_GetLocalTimestamp(uintptr(s))
}

func (s steamNetworkingUtils) SetDebugOutputFunction(DetailLevel ESteamNetworkingSocketsDebugOutputType, Func FSteamNetworkingSocketsDebugOutput) {
	iSteamNetworkingUtils_SetDebugOutputFunction(uintptr(s), DetailLevel, Func)
}

func (s steamNetworkingUtils) IsFakeIPv4(IPv4 uint32) bool {
	return iSteamNetworkingUtils_IsFakeIPv4(uintptr(s), IPv4)
}

func (s steamNetworkingUtils) GetIPv4FakeIPType(IPv4 uint32) ESteamNetworkingFakeIPType {
	return iSteamNetworkingUtils_GetIPv4FakeIPType(uintptr(s), IPv4)
}

func (s steamNetworkingUtils) GetRealIdentityForFakeIP(fakeIP SteamNetworkingIPAddr) (OutRealIdentity SteamNetworkingIdentity, result EResult) {
	result = iSteamNetworkingUtils_GetRealIdentityForFakeIP(uintptr(s), structToUintptr[SteamNetworkingIPAddr](&fakeIP), &OutRealIdentity)
	return OutRealIdentity, result
}

func (s steamNetworkingUtils) SetGlobalConfigValueInt32(eValue ESteamNetworkingConfigValue, val int32) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueInt32(uintptr(s), eValue, val)
}

func (s steamNetworkingUtils) SetGlobalConfigValueFloat(eValue ESteamNetworkingConfigValue, val float32) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueFloat(uintptr(s), eValue, val)
}

func (s steamNetworkingUtils) SetGlobalConfigValueString(eValue ESteamNetworkingConfigValue, val string) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueString(uintptr(s), eValue, val)
}

func (s steamNetworkingUtils) SetGlobalConfigValuePtr(eValue ESteamNetworkingConfigValue, val []byte) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValuePtr(uintptr(s), eValue, val)
}

func (s steamNetworkingUtils) SetConnectionConfigValueInt32(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueInt32(uintptr(s), Conn, eValue, val)
}

func (s steamNetworkingUtils) SetConnectionConfigValueFloat(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val float32) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueFloat(uintptr(s), Conn, eValue, val)
}

func (s steamNetworkingUtils) SetConnectionConfigValueString(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueString(uintptr(s), Conn, eValue, val)
}

func (s steamNetworkingUtils) SetGlobalCallback_SteamNetConnectionStatusChanged(fnCallback FnSteamNetConnectionStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetGlobalCallback_SteamNetAuthenticationStatusChanged(fnCallback FnSteamNetAuthenticationStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetGlobalCallback_SteamRelayNetworkStatusChanged(fnCallback FnSteamRelayNetworkStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetGlobalCallback_FakeIPResult(fnCallback FnSteamNetworkingFakeIPResult) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetGlobalCallback_MessagesSessionRequest(fnCallback FnSteamNetworkingMessagesSessionRequest) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetGlobalCallback_MessagesSessionFailed(fnCallback FnSteamNetworkingMessagesSessionFailed) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed(uintptr(s), fnCallback)
}

func (s steamNetworkingUtils) SetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool {
	return iSteamNetworkingUtils_SetConfigValue(uintptr(s), eValue, eScopeType, scopeObj, eDataType, pArg)
}

func (s steamNetworkingUtils) SetConfigValueStruct(opt SteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool {
	return iSteamNetworkingUtils_SetConfigValueStruct(uintptr(s), structToUintptr[SteamNetworkingConfigValue](&opt), eScopeType, scopeObj)
}

func (s steamNetworkingUtils) GetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, cbResult size) (pOutDataType ESteamNetworkingConfigDataType, pResult []byte, result ESteamNetworkingGetConfigValueResult) {
	pResult = make([]byte, 0, cbResult)
	result = iSteamNetworkingUtils_GetConfigValue(uintptr(s), eValue, eScopeType, scopeObj, &pOutDataType, pResult, &cbResult)
	return pOutDataType, pResult, result
}

func (s steamNetworkingUtils) GetConfigValueInfo(eValue ESteamNetworkingConfigValue) (pOutDataType ESteamNetworkingConfigDataType, pOutScope ESteamNetworkingConfigScope, name string) {
	nameBuf := iSteamNetworkingUtils_GetConfigValueInfo(uintptr(s), eValue, &pOutDataType, &pOutScope)
	if nameBuf == nil {
		return pOutDataType, pOutScope, ""
	}
	return pOutDataType, pOutScope, string(nameBuf)
}

func (s steamNetworkingUtils) IterateGenericEditableConfigValues(eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue {
	return iSteamNetworkingUtils_IterateGenericEditableConfigValues(uintptr(s), eCurrent, bEnumerateDevVars)
}

func (s steamNetworkingUtils) SteamNetworkingIPAddr_ToString(addr *SteamNetworkingIPAddr, cbBuf uint32, bWithPort bool) string {
	var buf []byte = make([]byte, 0, cbBuf)
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString(uintptr(s), addr, buf, cbBuf, bWithPort)
	return string(buf)
}

func (s steamNetworkingUtils) SteamNetworkingIPAddr_ParseString(pAddr *SteamNetworkingIPAddr, pszStr string) bool {
	return iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString(uintptr(s), pAddr, pszStr)
}

func (s steamNetworkingUtils) SteamNetworkingIPAddr_GetFakeIPType(addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType {
	return iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType(uintptr(s), addr)
}

func (s steamNetworkingUtils) SteamNetworkingIdentity_ToString(identity *SteamNetworkingIdentity, cbBuf uint32) string {
	var buf []byte = make([]byte, 0, cbBuf)
	iSteamNetworkingUtils_SteamNetworkingIdentity_ToString(uintptr(s), identity, buf, cbBuf)
	return string(buf)
}

func (s steamNetworkingUtils) SteamNetworkingIdentity_ParseString(pIdentity *SteamNetworkingIdentity, pszStr string) bool {
	return iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString(uintptr(s), pIdentity, pszStr)
}

// Steam Parties
const (
	PartyBeaconIdInvalid PartyBeaconID = 0
)

type ESteamPartyBeaconLocationType int32

const (
	ESteamPartyBeaconLocationType_Invalid   ESteamPartyBeaconLocationType = 0
	ESteamPartyBeaconLocationType_ChatGroup ESteamPartyBeaconLocationType = 1
	ESteamPartyBeaconLocationType_Max       ESteamPartyBeaconLocationType = 2
)

type ESteamPartyBeaconLocationData int32

const (
	ESteamPartyBeaconLocationData_Invalid       ESteamPartyBeaconLocationData = 0
	ESteamPartyBeaconLocationData_Name          ESteamPartyBeaconLocationData = 1
	ESteamPartyBeaconLocationData_IconURLSmall  ESteamPartyBeaconLocationData = 2
	ESteamPartyBeaconLocationData_IconURLMedium ESteamPartyBeaconLocationData = 3
	ESteamPartyBeaconLocationData_IconURLLarge  ESteamPartyBeaconLocationData = 4
)

type PartyBeaconID uint64
type SteamPartyBeaconLocation struct {
	Type       ESteamPartyBeaconLocationType
	LocationID uint64
}

const (
	flatAPI_SteamParties                                 = "SteamAPI_SteamParties_v002"
	flatAPI_ISteamParties_GetNumActiveBeacons            = "SteamAPI_ISteamParties_GetNumActiveBeacons"
	flatAPI_ISteamParties_GetBeaconByIndex               = "SteamAPI_ISteamParties_GetBeaconByIndex"
	flatAPI_ISteamParties_GetBeaconDetails               = "SteamAPI_ISteamParties_GetBeaconDetails"
	flatAPI_ISteamParties_JoinParty                      = "SteamAPI_ISteamParties_JoinParty"
	flatAPI_ISteamParties_GetNumAvailableBeaconLocations = "SteamAPI_ISteamParties_GetNumAvailableBeaconLocations"
	flatAPI_ISteamParties_GetAvailableBeaconLocations    = "SteamAPI_ISteamParties_GetAvailableBeaconLocations"
	flatAPI_ISteamParties_CreateBeacon                   = "SteamAPI_ISteamParties_CreateBeacon"
	flatAPI_ISteamParties_OnReservationCompleted         = "SteamAPI_ISteamParties_OnReservationCompleted"
	flatAPI_ISteamParties_CancelReservation              = "SteamAPI_ISteamParties_CancelReservation"
	flatAPI_ISteamParties_ChangeNumOpenSlots             = "SteamAPI_ISteamParties_ChangeNumOpenSlots"
	flatAPI_ISteamParties_DestroyBeacon                  = "SteamAPI_ISteamParties_DestroyBeacon"
	flatAPI_ISteamParties_GetBeaconLocationData          = "SteamAPI_ISteamParties_GetBeaconLocationData"
)

var (
	steamParties_init                            func() uintptr
	iSteamParties_GetNumActiveBeacons            func(steamParties uintptr) uint32
	iSteamParties_GetBeaconByIndex               func(steamParties uintptr, unIndex uint32) PartyBeaconID
	iSteamParties_GetBeaconDetails               func(steamParties uintptr, ulBeaconID PartyBeaconID, pSteamIDBeaconOwner *CSteamID, pLocation *SteamPartyBeaconLocation, pchMetadata []byte, cchMetadata int32) bool
	iSteamParties_JoinParty                      func(steamParties uintptr, ulBeaconID PartyBeaconID) SteamAPICall
	iSteamParties_GetNumAvailableBeaconLocations func(steamParties uintptr, puNumLocations *uint32) bool
	iSteamParties_GetAvailableBeaconLocations    func(steamParties uintptr, pLocationList []SteamPartyBeaconLocation, uMaxNumLocations uint32) bool
	iSteamParties_CreateBeacon                   func(steamParties uintptr, unOpenSlots uint32, pBeaconLocation uintptr, pchConnectString string, pchMetadata string) SteamAPICall
	iSteamParties_OnReservationCompleted         func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID Uint64SteamID)
	iSteamParties_CancelReservation              func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID Uint64SteamID)
	iSteamParties_ChangeNumOpenSlots             func(steamParties uintptr, ulBeacon PartyBeaconID, unOpenSlots uint32) SteamAPICall
	iSteamParties_DestroyBeacon                  func(steamParties uintptr, ulBeacon PartyBeaconID) bool
	iSteamParties_GetBeaconLocationData          func(steamParties uintptr, BeaconLocation uintptr, eData ESteamPartyBeaconLocationData, pchDataStringOut []byte, cchDataStringOut int32) bool
)

type steamParties uintptr

func SteamParties() ISteamParties {
	return steamParties(steamParties_init())
}

func (s steamParties) GetNumActiveBeacons() uint32 {
	return iSteamParties_GetNumActiveBeacons(uintptr(s))
}

func (s steamParties) GetBeaconByIndex(Index uint32) PartyBeaconID {
	return iSteamParties_GetBeaconByIndex(uintptr(s), Index)
}

func (s steamParties) GetBeaconDetails(BeaconID PartyBeaconID, MetadataSize int32) (beaconOwnerSteamID CSteamID, Location SteamPartyBeaconLocation, metadata string, success bool) {
	var pchMetadata []byte = make([]byte, 0, MetadataSize)
	success = iSteamParties_GetBeaconDetails(uintptr(s), BeaconID, &beaconOwnerSteamID, &Location, pchMetadata, MetadataSize)
	return beaconOwnerSteamID, Location, string(pchMetadata), success
}

func (s steamParties) JoinParty(BeaconID PartyBeaconID) CallResult[JoinPartyCallback] {
	handle := iSteamParties_JoinParty(uintptr(s), BeaconID)
	return CallResult[JoinPartyCallback]{Handle: handle}
}

func (s steamParties) GetNumAvailableBeaconLocations() (NumLocations uint32, success bool) {
	success = iSteamParties_GetNumAvailableBeaconLocations(uintptr(s), &NumLocations)
	return NumLocations, success
}

func (s steamParties) GetAvailableBeaconLocations(MaxNumLocations uint32) (locationList []SteamPartyBeaconLocation, success bool) {
	pLocationList := make([]SteamPartyBeaconLocation, 0, MaxNumLocations)
	success = iSteamParties_GetAvailableBeaconLocations(uintptr(s), pLocationList, MaxNumLocations)
	return pLocationList, success
}

func (s steamParties) CreateBeacon(OpenSlots uint32, BeaconLocation SteamPartyBeaconLocation, ConnectString string, Metadata string) CallResult[CreateBeaconCallback] {
	handle := iSteamParties_CreateBeacon(uintptr(s), OpenSlots, structToUintptr[SteamPartyBeaconLocation](&BeaconLocation), ConnectString, Metadata)
	return CallResult[CreateBeaconCallback]{Handle: handle}
}

func (s steamParties) OnReservationCompleted(Beacon PartyBeaconID, userSteamID Uint64SteamID) {
	iSteamParties_OnReservationCompleted(uintptr(s), Beacon, userSteamID)
}

func (s steamParties) CancelReservation(Beacon PartyBeaconID, userSteamID Uint64SteamID) {
	iSteamParties_CancelReservation(uintptr(s), Beacon, userSteamID)
}

func (s steamParties) ChangeNumOpenSlots(Beacon PartyBeaconID, OpenSlots uint32) CallResult[ChangeNumOpenSlotsCallback] {
	handle := iSteamParties_ChangeNumOpenSlots(uintptr(s), Beacon, OpenSlots)
	return CallResult[ChangeNumOpenSlotsCallback]{Handle: handle}
}

func (s steamParties) DestroyBeacon(beacon PartyBeaconID) bool {
	return iSteamParties_DestroyBeacon(uintptr(s), beacon)
}

func (s steamParties) GetBeaconLocationData(BeaconLocation SteamPartyBeaconLocation, Data ESteamPartyBeaconLocationData, DataStringOut int32) (data string, success bool) {
	pchDataStringOut := make([]byte, 0, DataStringOut)
	success = iSteamParties_GetBeaconLocationData(uintptr(s), structToUintptr[SteamPartyBeaconLocation](&BeaconLocation), Data, pchDataStringOut, DataStringOut)
	return string(pchDataStringOut), success
}

// Steam Remote Play
const (
	flatAPI_SteamRemotePlay                                = "SteamAPI_SteamRemotePlay_v002"
	flatAPI_ISteamRemotePlay_GetSessionCount               = "SteamAPI_ISteamRemotePlay_GetSessionCount"
	flatAPI_ISteamRemotePlay_GetSessionID                  = "SteamAPI_ISteamRemotePlay_GetSessionID"
	flatAPI_ISteamRemotePlay_GetSessionSteamID             = "SteamAPI_ISteamRemotePlay_GetSessionSteamID"
	flatAPI_ISteamRemotePlay_GetSessionClientName          = "SteamAPI_ISteamRemotePlay_GetSessionClientName"
	flatAPI_ISteamRemotePlay_GetSessionClientFormFactor    = "SteamAPI_ISteamRemotePlay_GetSessionClientFormFactor"
	flatAPI_ISteamRemotePlay_BGetSessionClientResolution   = "SteamAPI_ISteamRemotePlay_BGetSessionClientResolution"
	flatAPI_ISteamRemotePlay_BStartRemotePlayTogether      = "SteamAPI_ISteamRemotePlay_BStartRemotePlayTogether"
	flatAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite = "SteamAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite"
)

type RemotePlaySessionID uint

type ESteamDeviceFormFactor int32

const (
	ESteamDeviceFormFactor_Unknown   ESteamDeviceFormFactor = 0
	ESteamDeviceFormFactor_Phone     ESteamDeviceFormFactor = 1
	ESteamDeviceFormFactor_Tablet    ESteamDeviceFormFactor = 2
	ESteamDeviceFormFactor_Computer  ESteamDeviceFormFactor = 3
	ESteamDeviceFormFactor_TV        ESteamDeviceFormFactor = 4
	ESteamDeviceFormFactor_VRHeadset ESteamDeviceFormFactor = 5
)

var (
	steamRemotePlay_init                           func() uintptr
	iSteamRemotePlay_GetSessionCount               func(steamRemotePlay uintptr) uint32
	iSteamRemotePlay_GetSessionID                  func(steamRemotePlay uintptr, iSessionIndex int32) RemotePlaySessionID
	iSteamRemotePlay_GetSessionSteamID             func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) Uint64SteamID
	iSteamRemotePlay_GetSessionClientName          func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) string
	iSteamRemotePlay_GetSessionClientFormFactor    func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) ESteamDeviceFormFactor
	iSteamRemotePlay_BGetSessionClientResolution   func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID, pnResolutionX *int32, pnResolutionY *int32) bool
	iSteamRemotePlay_BStartRemotePlayTogether      func(steamRemotePlay uintptr, bShowOverlay bool) bool
	iSteamRemotePlay_BSendRemotePlayTogetherInvite func(steamRemotePlay uintptr, steamIDFriend Uint64SteamID) bool
)

type steamRemotePlay uintptr

func SteamRemotePlay() ISteamRemotePlay {
	return steamRemotePlay(steamRemotePlay_init())
}

func (s steamRemotePlay) GetSessionCount() uint32 {
	return iSteamRemotePlay_GetSessionCount(uintptr(s))
}

func (s steamRemotePlay) GetSessionID(SessionIndex int32) RemotePlaySessionID {
	return iSteamRemotePlay_GetSessionID(uintptr(s), SessionIndex)
}

func (s steamRemotePlay) GetSessionSteamID(SessionID RemotePlaySessionID) Uint64SteamID {
	return iSteamRemotePlay_GetSessionSteamID(uintptr(s), SessionID)
}

func (s steamRemotePlay) GetSessionClientName(SessionID RemotePlaySessionID) string {
	return iSteamRemotePlay_GetSessionClientName(uintptr(s), SessionID)
}

func (s steamRemotePlay) GetSessionClientFormFactor(SessionID RemotePlaySessionID) ESteamDeviceFormFactor {
	return iSteamRemotePlay_GetSessionClientFormFactor(uintptr(s), SessionID)
}

func (s steamRemotePlay) BGetSessionClientResolution(SessionID RemotePlaySessionID) (int32, int32, bool) {
	var pnResolutionX, pnResolutionY int32
	result := iSteamRemotePlay_BGetSessionClientResolution(uintptr(s), SessionID, &pnResolutionX, &pnResolutionY)
	return pnResolutionX, pnResolutionY, result
}

func (s steamRemotePlay) BStartRemotePlayTogether(ShowOverlay bool) bool {
	return iSteamRemotePlay_BStartRemotePlayTogether(uintptr(s), ShowOverlay)
}

func (s steamRemotePlay) BSendRemotePlayTogetherInvite(friendSteamID Uint64SteamID) bool {
	return iSteamRemotePlay_BSendRemotePlayTogetherInvite(uintptr(s), friendSteamID)
}

// Steam Remote Storage
type PublishedFileId uint64

type SteamParamStringArray struct {
	Strings    []string
	NumStrings int32
}
type ERemoteStoragePublishedFileVisibility int32

const (
	ERemoteStoragePublishedFileVisibility_Public      ERemoteStoragePublishedFileVisibility = 0
	ERemoteStoragePublishedFileVisibility_FriendsOnly ERemoteStoragePublishedFileVisibility = 1
	ERemoteStoragePublishedFileVisibility_Private     ERemoteStoragePublishedFileVisibility = 2
	ERemoteStoragePublishedFileVisibility_Unlisted    ERemoteStoragePublishedFileVisibility = 3
)

type EWorkshopFileType int32

const (
	EWorkshopFileType_First                  EWorkshopFileType = 0
	EWorkshopFileType_Community              EWorkshopFileType = 0
	EWorkshopFileType_Microtransaction       EWorkshopFileType = 1
	EWorkshopFileType_Collection             EWorkshopFileType = 2
	EWorkshopFileType_Art                    EWorkshopFileType = 3
	EWorkshopFileType_Video                  EWorkshopFileType = 4
	EWorkshopFileType_Screenshot             EWorkshopFileType = 5
	EWorkshopFileType_Game                   EWorkshopFileType = 6
	EWorkshopFileType_Software               EWorkshopFileType = 7
	EWorkshopFileType_Concept                EWorkshopFileType = 8
	EWorkshopFileType_WebGuide               EWorkshopFileType = 9
	EWorkshopFileType_IntegratedGuide        EWorkshopFileType = 10
	EWorkshopFileType_Merch                  EWorkshopFileType = 11
	EWorkshopFileType_ControllerBinding      EWorkshopFileType = 12
	EWorkshopFileType_SteamworksAccessInvite EWorkshopFileType = 13
	EWorkshopFileType_SteamVideo             EWorkshopFileType = 14
	EWorkshopFileType_GameManagedItem        EWorkshopFileType = 15
	EWorkshopFileType_Clip                   EWorkshopFileType = 16
	EWorkshopFileType_Max                    EWorkshopFileType = 17
)

const (
	MaxCloudFileChunkSize uint32 = 100 * 1024 * 1024

	PublishedFileIdInvalid           PublishedFileId           = 0
	UGCHandleInvalid                 UGCHandle                 = 0xffffffffffffffff
	PublishedFileUpdateHandleInvalid PublishedFileUpdateHandle = 0xffffffffffffffff
	UGCFileStreamHandleInvalid       UGCFileWriteStreamHandle  = 0xffffffffffffffff

	PublishedDocumentTitleMax             uint32 = 128 + 1
	PublishedDocumentDescriptionMax       uint32 = 8000
	PublishedDocumentChangeDescriptionMax uint32 = 8000
	EnumeratePublishedFilesMaxResults     uint32 = 50
	TagListMax                            uint32 = 1024 + 1
	FilenameMax                           uint32 = 260
	PublishedFileURLMax                   uint32 = 256
)

type EWorkshopVote int32

const (
	EWorkshopVote_Unvoted EWorkshopVote = 0
	EWorkshopVote_For     EWorkshopVote = 1
	EWorkshopVote_Against EWorkshopVote = 2
	EWorkshopVote_Later   EWorkshopVote = 3
)

type ERemoteStorageFilePathType int32

const (
	ERemoteStorageFilePathType_Invalid     ERemoteStorageFilePathType = 0
	ERemoteStorageFilePathType_Absolute    ERemoteStorageFilePathType = 1
	ERemoteStorageFilePathType_APIFilename ERemoteStorageFilePathType = 2
)

type EUGCReadAction int32

const (
	EUGCRead_ContinueReadingUntilFinished EUGCReadAction = 0
	EUGCRead_ContinueReading              EUGCReadAction = 1
	EUGCRead_Close                        EUGCReadAction = 2
)

type ERemoteStorageLocalFileChange int32

const (
	ERemoteStorageLocalFileChange_Invalid     ERemoteStorageLocalFileChange = 0
	ERemoteStorageLocalFileChange_FileUpdated ERemoteStorageLocalFileChange = 1
	ERemoteStorageLocalFileChange_FileDeleted ERemoteStorageLocalFileChange = 2
)

type EWorkshopVideoProvider int32

const (
	EWorkshopVideoProvider_None    EWorkshopVideoProvider = 0
	EWorkshopVideoProvider_Youtube EWorkshopVideoProvider = 1
)

type EWorkshopFileAction int32

const (
	EWorkshopFileAction_Played    EWorkshopFileAction = 0
	EWorkshopFileAction_Completed EWorkshopFileAction = 1
)

type EWorkshopEnumerationType int32

const (
	EWorkshopEnumerationType_RankedByVote            EWorkshopEnumerationType = 0
	EWorkshopEnumerationType_Recent                  EWorkshopEnumerationType = 1
	EWorkshopEnumerationType_Trending                EWorkshopEnumerationType = 2
	EWorkshopEnumerationType_FavoritesOfFriends      EWorkshopEnumerationType = 3
	EWorkshopEnumerationType_VotedByFriends          EWorkshopEnumerationType = 4
	EWorkshopEnumerationType_ContentByFriends        EWorkshopEnumerationType = 5
	EWorkshopEnumerationType_RecentFromFollowedUsers EWorkshopEnumerationType = 6
)

type ERemoteStoragePlatform int32

const (
	ERemoteStoragePlatform_None    ERemoteStoragePlatform = 0
	ERemoteStoragePlatform_Windows ERemoteStoragePlatform = 1
	ERemoteStoragePlatform_OSX     ERemoteStoragePlatform = 2
	ERemoteStoragePlatform_PS3     ERemoteStoragePlatform = 4
	ERemoteStoragePlatform_Linux   ERemoteStoragePlatform = 8
	ERemoteStoragePlatform_Switch  ERemoteStoragePlatform = 16
	ERemoteStoragePlatform_Android ERemoteStoragePlatform = 32
	ERemoteStoragePlatform_IOS     ERemoteStoragePlatform = 64
	ERemoteStoragePlatform_All     ERemoteStoragePlatform = -1
)

type PublishedFileUpdateHandle uint64
type UGCFileWriteStreamHandle uint64

const (
	flatAPI_SteamRemoteStorage                                          = "SteamAPI_SteamRemoteStorage_v016"
	flatAPI_ISteamRemoteStorage_FileWrite                               = "SteamAPI_ISteamRemoteStorage_FileWrite"
	flatAPI_ISteamRemoteStorage_FileRead                                = "SteamAPI_ISteamRemoteStorage_FileRead"
	flatAPI_ISteamRemoteStorage_FileWriteAsync                          = "SteamAPI_ISteamRemoteStorage_FileWriteAsync"
	flatAPI_ISteamRemoteStorage_FileReadAsync                           = "SteamAPI_ISteamRemoteStorage_FileReadAsync"
	flatAPI_ISteamRemoteStorage_FileReadAsyncComplete                   = "SteamAPI_ISteamRemoteStorage_FileReadAsyncComplete"
	flatAPI_ISteamRemoteStorage_FileForget                              = "SteamAPI_ISteamRemoteStorage_FileForget"
	flatAPI_ISteamRemoteStorage_FileDelete                              = "SteamAPI_ISteamRemoteStorage_FileDelete"
	flatAPI_ISteamRemoteStorage_FileShare                               = "SteamAPI_ISteamRemoteStorage_FileShare"
	flatAPI_ISteamRemoteStorage_SetSyncPlatforms                        = "SteamAPI_ISteamRemoteStorage_SetSyncPlatforms"
	flatAPI_ISteamRemoteStorage_FileWriteStreamOpen                     = "SteamAPI_ISteamRemoteStorage_FileWriteStreamOpen"
	flatAPI_ISteamRemoteStorage_FileWriteStreamWriteChunk               = "SteamAPI_ISteamRemoteStorage_FileWriteStreamWriteChunk"
	flatAPI_ISteamRemoteStorage_FileWriteStreamClose                    = "SteamAPI_ISteamRemoteStorage_FileWriteStreamClose"
	flatAPI_ISteamRemoteStorage_FileWriteStreamCancel                   = "SteamAPI_ISteamRemoteStorage_FileWriteStreamCancel"
	flatAPI_ISteamRemoteStorage_FileExists                              = "SteamAPI_ISteamRemoteStorage_FileExists"
	flatAPI_ISteamRemoteStorage_FilePersisted                           = "SteamAPI_ISteamRemoteStorage_FilePersisted"
	flatAPI_ISteamRemoteStorage_GetFileSize                             = "SteamAPI_ISteamRemoteStorage_GetFileSize"
	flatAPI_ISteamRemoteStorage_GetFileTimestamp                        = "SteamAPI_ISteamRemoteStorage_GetFileTimestamp"
	flatAPI_ISteamRemoteStorage_GetSyncPlatforms                        = "SteamAPI_ISteamRemoteStorage_GetSyncPlatforms"
	flatAPI_ISteamRemoteStorage_GetFileCount                            = "SteamAPI_ISteamRemoteStorage_GetFileCount"
	flatAPI_ISteamRemoteStorage_GetFileNameAndSize                      = "SteamAPI_ISteamRemoteStorage_GetFileNameAndSize"
	flatAPI_ISteamRemoteStorage_GetQuota                                = "SteamAPI_ISteamRemoteStorage_GetQuota"
	flatAPI_ISteamRemoteStorage_IsCloudEnabledForAccount                = "SteamAPI_ISteamRemoteStorage_IsCloudEnabledForAccount"
	flatAPI_ISteamRemoteStorage_IsCloudEnabledForApp                    = "SteamAPI_ISteamRemoteStorage_IsCloudEnabledForApp"
	flatAPI_ISteamRemoteStorage_SetCloudEnabledForApp                   = "SteamAPI_ISteamRemoteStorage_SetCloudEnabledForApp"
	flatAPI_ISteamRemoteStorage_UGCDownload                             = "SteamAPI_ISteamRemoteStorage_UGCDownload"
	flatAPI_ISteamRemoteStorage_GetUGCDownloadProgress                  = "SteamAPI_ISteamRemoteStorage_GetUGCDownloadProgress"
	flatAPI_ISteamRemoteStorage_GetUGCDetails                           = "SteamAPI_ISteamRemoteStorage_GetUGCDetails"
	flatAPI_ISteamRemoteStorage_UGCRead                                 = "SteamAPI_ISteamRemoteStorage_UGCRead"
	flatAPI_ISteamRemoteStorage_GetCachedUGCCount                       = "SteamAPI_ISteamRemoteStorage_GetCachedUGCCount"
	flatAPI_ISteamRemoteStorage_GetCachedUGCHandle                      = "SteamAPI_ISteamRemoteStorage_GetCachedUGCHandle"
	flatAPI_ISteamRemoteStorage_PublishWorkshopFile                     = "SteamAPI_ISteamRemoteStorage_PublishWorkshopFile"
	flatAPI_ISteamRemoteStorage_CreatePublishedFileUpdateRequest        = "SteamAPI_ISteamRemoteStorage_CreatePublishedFileUpdateRequest"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileFile                 = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileFile"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFilePreviewFile          = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFilePreviewFile"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileTitle                = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTitle"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileDescription          = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileDescription"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileVisibility           = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileVisibility"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileTags                 = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTags"
	flatAPI_ISteamRemoteStorage_CommitPublishedFileUpdate               = "SteamAPI_ISteamRemoteStorage_CommitPublishedFileUpdate"
	flatAPI_ISteamRemoteStorage_GetPublishedFileDetails                 = "SteamAPI_ISteamRemoteStorage_GetPublishedFileDetails"
	flatAPI_ISteamRemoteStorage_DeletePublishedFile                     = "SteamAPI_ISteamRemoteStorage_DeletePublishedFile"
	flatAPI_ISteamRemoteStorage_EnumerateUserPublishedFiles             = "SteamAPI_ISteamRemoteStorage_EnumerateUserPublishedFiles"
	flatAPI_ISteamRemoteStorage_SubscribePublishedFile                  = "SteamAPI_ISteamRemoteStorage_SubscribePublishedFile"
	flatAPI_ISteamRemoteStorage_EnumerateUserSubscribedFiles            = "SteamAPI_ISteamRemoteStorage_EnumerateUserSubscribedFiles"
	flatAPI_ISteamRemoteStorage_UnsubscribePublishedFile                = "SteamAPI_ISteamRemoteStorage_UnsubscribePublishedFile"
	flatAPI_ISteamRemoteStorage_UpdatePublishedFileSetChangeDescription = "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileSetChangeDescription"
	flatAPI_ISteamRemoteStorage_GetPublishedItemVoteDetails             = "SteamAPI_ISteamRemoteStorage_GetPublishedItemVoteDetails"
	flatAPI_ISteamRemoteStorage_UpdateUserPublishedItemVote             = "SteamAPI_ISteamRemoteStorage_UpdateUserPublishedItemVote"
	flatAPI_ISteamRemoteStorage_GetUserPublishedItemVoteDetails         = "SteamAPI_ISteamRemoteStorage_GetUserPublishedItemVoteDetails"
	flatAPI_ISteamRemoteStorage_EnumerateUserSharedWorkshopFiles        = "SteamAPI_ISteamRemoteStorage_EnumerateUserSharedWorkshopFiles"
	flatAPI_ISteamRemoteStorage_PublishVideo                            = "SteamAPI_ISteamRemoteStorage_PublishVideo"
	flatAPI_ISteamRemoteStorage_SetUserPublishedFileAction              = "SteamAPI_ISteamRemoteStorage_SetUserPublishedFileAction"
	flatAPI_ISteamRemoteStorage_EnumeratePublishedFilesByUserAction     = "SteamAPI_ISteamRemoteStorage_EnumeratePublishedFilesByUserAction"
	flatAPI_ISteamRemoteStorage_EnumeratePublishedWorkshopFiles         = "SteamAPI_ISteamRemoteStorage_EnumeratePublishedWorkshopFiles"
	flatAPI_ISteamRemoteStorage_UGCDownloadToLocation                   = "SteamAPI_ISteamRemoteStorage_UGCDownloadToLocation"
	flatAPI_ISteamRemoteStorage_GetLocalFileChangeCount                 = "SteamAPI_ISteamRemoteStorage_GetLocalFileChangeCount"
	flatAPI_ISteamRemoteStorage_GetLocalFileChange                      = "SteamAPI_ISteamRemoteStorage_GetLocalFileChange"
	flatAPI_ISteamRemoteStorage_BeginFileWriteBatch                     = "SteamAPI_ISteamRemoteStorage_BeginFileWriteBatch"
	flatAPI_ISteamRemoteStorage_EndFileWriteBatch                       = "SteamAPI_ISteamRemoteStorage_EndFileWriteBatch"
)

var (
	steamRemoteStorage_init                                     func() uintptr
	iSteamRemoteStorage_FileWrite                               func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubData int32) bool
	iSteamRemoteStorage_FileRead                                func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubDataToRead int32) int32
	iSteamRemoteStorage_FileWriteAsync                          func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubData uint32) SteamAPICall
	iSteamRemoteStorage_FileReadAsync                           func(steamRemoteStorage uintptr, pchFile string, nOffset uint32, cubToRead uint32) SteamAPICall
	iSteamRemoteStorage_FileReadAsyncComplete                   func(steamRemoteStorage uintptr, hReadCall SteamAPICall, pvBuffer []byte, cubToRead uint32) bool
	iSteamRemoteStorage_FileForget                              func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_FileDelete                              func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_FileShare                               func(steamRemoteStorage uintptr, pchFile string) SteamAPICall
	iSteamRemoteStorage_SetSyncPlatforms                        func(steamRemoteStorage uintptr, pchFile string, eRemoteStoragePlatform ERemoteStoragePlatform) bool
	iSteamRemoteStorage_FileWriteStreamOpen                     func(steamRemoteStorage uintptr, pchFile string) UGCFileWriteStreamHandle
	iSteamRemoteStorage_FileWriteStreamWriteChunk               func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle, pvData []byte, cubData int32) bool
	iSteamRemoteStorage_FileWriteStreamClose                    func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle) bool
	iSteamRemoteStorage_FileWriteStreamCancel                   func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle) bool
	iSteamRemoteStorage_FileExists                              func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_FilePersisted                           func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_GetFileSize                             func(steamRemoteStorage uintptr, pchFile string) int32
	iSteamRemoteStorage_GetFileTimestamp                        func(steamRemoteStorage uintptr, pchFile string) int64
	iSteamRemoteStorage_GetSyncPlatforms                        func(steamRemoteStorage uintptr, pchFile string) ERemoteStoragePlatform
	iSteamRemoteStorage_GetFileCount                            func(steamRemoteStorage uintptr) int32
	iSteamRemoteStorage_GetFileNameAndSize                      func(steamRemoteStorage uintptr, iFile int32, pnFileSizeInBytes *int32) []byte
	iSteamRemoteStorage_GetQuota                                func(steamRemoteStorage uintptr, pnTotalBytes *uint64, puAvailableBytes *uint64) bool
	iSteamRemoteStorage_IsCloudEnabledForAccount                func(steamRemoteStorage uintptr) bool
	iSteamRemoteStorage_IsCloudEnabledForApp                    func(steamRemoteStorage uintptr) bool
	iSteamRemoteStorage_SetCloudEnabledForApp                   func(steamRemoteStorage uintptr, bEnabled bool)
	iSteamRemoteStorage_UGCDownload                             func(steamRemoteStorage uintptr, hContent UGCHandle, unPriority uint32) SteamAPICall
	iSteamRemoteStorage_GetUGCDownloadProgress                  func(steamRemoteStorage uintptr, hContent UGCHandle, pnBytesDownloaded *int32, pnBytesExpected *int32) bool
	iSteamRemoteStorage_GetUGCDetails                           func(steamRemoteStorage uintptr, hContent UGCHandle, pnAppID *AppId_t, ppchName *[]byte, pnFileSizeInBytes *int32, pSteamIDOwner *CSteamID) bool
	iSteamRemoteStorage_UGCRead                                 func(steamRemoteStorage uintptr, hContent UGCHandle, pvData []byte, cubDataToRead int32, cOffset uint32, eAction EUGCReadAction) int32
	iSteamRemoteStorage_GetCachedUGCCount                       func(steamRemoteStorage uintptr) int32
	iSteamRemoteStorage_GetCachedUGCHandle                      func(steamRemoteStorage uintptr, iCachedContent int32) UGCHandle
	iSteamRemoteStorage_PublishWorkshopFile                     func(steamRemoteStorage uintptr, pchFile string, pchPreviewFile string, nConsumerAppId AppId_t, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr, eWorkshopFileType EWorkshopFileType) SteamAPICall
	iSteamRemoteStorage_CreatePublishedFileUpdateRequest        func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) PublishedFileUpdateHandle
	iSteamRemoteStorage_UpdatePublishedFileFile                 func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchFile string) bool
	iSteamRemoteStorage_UpdatePublishedFilePreviewFile          func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchPreviewFile string) bool
	iSteamRemoteStorage_UpdatePublishedFileTitle                func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchTitle string) bool
	iSteamRemoteStorage_UpdatePublishedFileDescription          func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchDescription string) bool
	iSteamRemoteStorage_UpdatePublishedFileVisibility           func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool
	iSteamRemoteStorage_UpdatePublishedFileTags                 func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pTags uintptr) bool
	iSteamRemoteStorage_CommitPublishedFileUpdate               func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle) SteamAPICall
	iSteamRemoteStorage_GetPublishedFileDetails                 func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, unMaxSecondsOld uint32) SteamAPICall
	iSteamRemoteStorage_DeletePublishedFile                     func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall
	iSteamRemoteStorage_EnumerateUserPublishedFiles             func(steamRemoteStorage uintptr, unStartIndex uint32) SteamAPICall
	iSteamRemoteStorage_SubscribePublishedFile                  func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall
	iSteamRemoteStorage_EnumerateUserSubscribedFiles            func(steamRemoteStorage uintptr, unStartIndex uint32) SteamAPICall
	iSteamRemoteStorage_UnsubscribePublishedFile                func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall
	iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchChangeDescription string) bool
	iSteamRemoteStorage_GetPublishedItemVoteDetails             func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall
	iSteamRemoteStorage_UpdateUserPublishedItemVote             func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, bVoteUp bool) SteamAPICall
	iSteamRemoteStorage_GetUserPublishedItemVoteDetails         func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall
	iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles        func(steamRemoteStorage uintptr, steamId Uint64SteamID, unStartIndex uint32, pRequiredTags uintptr, pExcludedTags uintptr) SteamAPICall
	iSteamRemoteStorage_PublishVideo                            func(steamRemoteStorage uintptr, eVideoProvider EWorkshopVideoProvider, pchVideoAccount string, pchVideoIdentifier string, pchPreviewFile string, nConsumerAppId AppId_t, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr) SteamAPICall
	iSteamRemoteStorage_SetUserPublishedFileAction              func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, eAction EWorkshopFileAction) SteamAPICall
	iSteamRemoteStorage_EnumeratePublishedFilesByUserAction     func(steamRemoteStorage uintptr, eAction EWorkshopFileAction, unStartIndex uint32) SteamAPICall
	iSteamRemoteStorage_EnumeratePublishedWorkshopFiles         func(steamRemoteStorage uintptr, eEnumerationType EWorkshopEnumerationType, unStartIndex uint32, unCount uint32, unDays uint32, pTags *SteamParamStringArray, pUserTags *SteamParamStringArray) SteamAPICall
	iSteamRemoteStorage_UGCDownloadToLocation                   func(steamRemoteStorage uintptr, hContent UGCHandle, pchLocation string, unPriority uint32) SteamAPICall
	iSteamRemoteStorage_GetLocalFileChangeCount                 func(steamRemoteStorage uintptr) int32
	iSteamRemoteStorage_GetLocalFileChange                      func(steamRemoteStorage uintptr, iFile int32, pEChangeType *ERemoteStorageLocalFileChange, pEFilePathType *ERemoteStorageFilePathType) []byte
	iSteamRemoteStorage_BeginFileWriteBatch                     func(steamRemoteStorage uintptr) bool
	iSteamRemoteStorage_EndFileWriteBatch                       func(steamRemoteStorage uintptr) bool
)

type steamRemoteStorage uintptr

func SteamRemoteStorage() ISteamRemoteStorage {
	return steamRemoteStorage(steamRemoteStorage_init())
}

func (s steamRemoteStorage) FileWrite(File string, Data []byte) bool {
	return iSteamRemoteStorage_FileWrite(uintptr(s), File, Data, int32(len(Data)))
}

func (s steamRemoteStorage) FileRead(File string, Data []byte) (bytesRead int32) {
	result := iSteamRemoteStorage_FileRead(uintptr(s), File, Data, int32(len(Data)))
	return result
}

func (s steamRemoteStorage) FileReadEx(File string, DataToReadSize int32) (Data []byte) {
	Data = make([]byte, 0, DataToReadSize)
	result := iSteamRemoteStorage_FileRead(uintptr(s), File, Data, DataToReadSize)
	return Data[:result]
}
func (s steamRemoteStorage) FileWriteAsync(File string, Data []byte) CallResult[RemoteStorageFileWriteAsyncComplete] {
	handle := iSteamRemoteStorage_FileWriteAsync(uintptr(s), File, Data, uint32(len(Data)))
	return CallResult[RemoteStorageFileWriteAsyncComplete]{Handle: handle}
}

func (s steamRemoteStorage) FileReadAsync(File string, Offset uint32, DataToReadSize uint32) CallResult[RemoteStorageFileReadAsyncComplete] {
	handle := iSteamRemoteStorage_FileReadAsync(uintptr(s), File, Offset, DataToReadSize)
	return CallResult[RemoteStorageFileReadAsyncComplete]{Handle: handle}
}

func (s steamRemoteStorage) FileReadAsyncComplete(ReadCallHandle SteamAPICall, DataToReadSize uint32) (Buffer []byte, success bool) {
	Buffer = make([]byte, 0, DataToReadSize)
	success = iSteamRemoteStorage_FileReadAsyncComplete(uintptr(s), ReadCallHandle, Buffer, DataToReadSize)
	return Buffer, success
}

func (s steamRemoteStorage) FileForget(File string) bool {
	return iSteamRemoteStorage_FileForget(uintptr(s), File)
}

func (s steamRemoteStorage) FileDelete(File string) bool {
	return iSteamRemoteStorage_FileDelete(uintptr(s), File)
}

func (s steamRemoteStorage) FileShare(File string) CallResult[RemoteStorageFileShareResult] {
	handle := iSteamRemoteStorage_FileShare(uintptr(s), File)
	return CallResult[RemoteStorageFileShareResult]{Handle: handle}
}

func (s steamRemoteStorage) SetSyncPlatforms(File string, RemoteStoragePlatform ERemoteStoragePlatform) bool {
	return iSteamRemoteStorage_SetSyncPlatforms(uintptr(s), File, RemoteStoragePlatform)
}

func (s steamRemoteStorage) FileWriteStreamOpen(File string) UGCFileWriteStreamHandle {
	return iSteamRemoteStorage_FileWriteStreamOpen(uintptr(s), File)
}

func (s steamRemoteStorage) FileWriteStreamWriteChunk(writeHandle UGCFileWriteStreamHandle, Data []byte) bool {
	return iSteamRemoteStorage_FileWriteStreamWriteChunk(uintptr(s), writeHandle, Data, int32(len(Data)))
}

func (s steamRemoteStorage) FileWriteStreamClose(writeHandle UGCFileWriteStreamHandle) bool {
	return iSteamRemoteStorage_FileWriteStreamClose(uintptr(s), writeHandle)
}

func (s steamRemoteStorage) FileWriteStreamCancel(writeHandle UGCFileWriteStreamHandle) bool {
	return iSteamRemoteStorage_FileWriteStreamCancel(uintptr(s), writeHandle)
}

func (s steamRemoteStorage) FileExists(File string) bool {
	return iSteamRemoteStorage_FileExists(uintptr(s), File)
}

func (s steamRemoteStorage) FilePersisted(File string) bool {
	return iSteamRemoteStorage_FilePersisted(uintptr(s), File)
}

func (s steamRemoteStorage) GetFileSize(File string) int32 {
	return iSteamRemoteStorage_GetFileSize(uintptr(s), File)
}

func (s steamRemoteStorage) GetFileTimestamp(File string) int64 {
	return iSteamRemoteStorage_GetFileTimestamp(uintptr(s), File)
}

func (s steamRemoteStorage) GetSyncPlatforms(File string) ERemoteStoragePlatform {
	return iSteamRemoteStorage_GetSyncPlatforms(uintptr(s), File)
}

func (s steamRemoteStorage) GetFileCount() int32 {
	return iSteamRemoteStorage_GetFileCount(uintptr(s))
}

func (s steamRemoteStorage) GetFileNameAndSize(FileIndex int32) (name string, FileSizeInBytes int32) {
	var nameBuf []byte = iSteamRemoteStorage_GetFileNameAndSize(uintptr(s), FileIndex, &FileSizeInBytes)
	return string(nameBuf), FileSizeInBytes
}

func (s steamRemoteStorage) GetQuota() (TotalBytes uint64, AvailableBytes uint64, success bool) {
	success = iSteamRemoteStorage_GetQuota(uintptr(s), &TotalBytes, &AvailableBytes)
	return TotalBytes, AvailableBytes, success
}

func (s steamRemoteStorage) IsCloudEnabledForAccount() bool {
	return iSteamRemoteStorage_IsCloudEnabledForAccount(uintptr(s))
}

func (s steamRemoteStorage) IsCloudEnabledForApp() bool {
	return iSteamRemoteStorage_IsCloudEnabledForApp(uintptr(s))
}

func (s steamRemoteStorage) SetCloudEnabledForApp(Enabled bool) {
	iSteamRemoteStorage_SetCloudEnabledForApp(uintptr(s), Enabled)
}

func (s steamRemoteStorage) UGCDownload(Content UGCHandle, Priority uint32) CallResult[RemoteStorageDownloadUGCResult] {
	handle := iSteamRemoteStorage_UGCDownload(uintptr(s), Content, Priority)
	return CallResult[RemoteStorageDownloadUGCResult]{Handle: handle}
}

func (s steamRemoteStorage) GetUGCDownloadProgress(Content UGCHandle) (BytesDownloaded int32, BytesExpected int32, success bool) {
	success = iSteamRemoteStorage_GetUGCDownloadProgress(uintptr(s), Content, &BytesDownloaded, &BytesExpected)
	return BytesDownloaded, BytesExpected, success
}

func (s steamRemoteStorage) GetUGCDetails(Content UGCHandle) (AppID AppId_t, Name string, FileSizeInBytes int32, ownerSteamID CSteamID, success bool) {
	var ppchNameBuf []byte = make([]byte, 0, 4096)
	success = iSteamRemoteStorage_GetUGCDetails(uintptr(s), Content, &AppID, &ppchNameBuf, &FileSizeInBytes, &ownerSteamID)
	return AppID, string(ppchNameBuf), FileSizeInBytes, ownerSteamID, success
}

func (s steamRemoteStorage) UGCRead(Content UGCHandle, DataToReadSize int32, Offset uint32, Action EUGCReadAction) (Data []byte) {
	Data = make([]byte, 0, DataToReadSize)
	result := iSteamRemoteStorage_UGCRead(uintptr(s), Content, Data, DataToReadSize, Offset, Action)
	return Data[:result]
}

func (s steamRemoteStorage) GetCachedUGCCount() int32 {
	return iSteamRemoteStorage_GetCachedUGCCount(uintptr(s))
}

func (s steamRemoteStorage) GetCachedUGCHandle(CachedContent int32) UGCHandle {
	return iSteamRemoteStorage_GetCachedUGCHandle(uintptr(s), CachedContent)
}

func (s steamRemoteStorage) PublishWorkshopFile(File string, PreviewFile string, ConsumerAppId AppId_t, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray, WorkshopFileType EWorkshopFileType) CallResult[RemoteStoragePublishFileProgress] {
	handle := iSteamRemoteStorage_PublishWorkshopFile(uintptr(s), File, PreviewFile, ConsumerAppId, Title, Description, Visibility, structToUintptr[SteamParamStringArray](&Tags), WorkshopFileType)
	return CallResult[RemoteStoragePublishFileProgress]{Handle: handle}
}

func (s steamRemoteStorage) CreatePublishedFileUpdateRequest(PublishedFileId PublishedFileId) PublishedFileUpdateHandle {
	return iSteamRemoteStorage_CreatePublishedFileUpdateRequest(uintptr(s), PublishedFileId)
}

func (s steamRemoteStorage) UpdatePublishedFileFile(updateHandle PublishedFileUpdateHandle, File string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileFile(uintptr(s), updateHandle, File)
}

func (s steamRemoteStorage) UpdatePublishedFilePreviewFile(updateHandle PublishedFileUpdateHandle, PreviewFile string) bool {
	return iSteamRemoteStorage_UpdatePublishedFilePreviewFile(uintptr(s), updateHandle, PreviewFile)
}

func (s steamRemoteStorage) UpdatePublishedFileTitle(updateHandle PublishedFileUpdateHandle, Title string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileTitle(uintptr(s), updateHandle, Title)
}

func (s steamRemoteStorage) UpdatePublishedFileDescription(updateHandle PublishedFileUpdateHandle, Description string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileDescription(uintptr(s), updateHandle, Description)
}

func (s steamRemoteStorage) UpdatePublishedFileVisibility(updateHandle PublishedFileUpdateHandle, Visibility ERemoteStoragePublishedFileVisibility) bool {
	return iSteamRemoteStorage_UpdatePublishedFileVisibility(uintptr(s), updateHandle, Visibility)
}

func (s steamRemoteStorage) UpdatePublishedFileTags(updateHandle PublishedFileUpdateHandle, Tags SteamParamStringArray) bool {
	return iSteamRemoteStorage_UpdatePublishedFileTags(uintptr(s), updateHandle, structToUintptr[SteamParamStringArray](&Tags))
}

func (s steamRemoteStorage) CommitPublishedFileUpdate(updateHandle PublishedFileUpdateHandle) CallResult[RemoteStorageUpdatePublishedFileResult] {
	handle := iSteamRemoteStorage_CommitPublishedFileUpdate(uintptr(s), updateHandle)
	return CallResult[RemoteStorageUpdatePublishedFileResult]{Handle: handle}
}

func (s steamRemoteStorage) GetPublishedFileDetails(PublishedFileId PublishedFileId, MaxSecondsOld uint32) CallResult[RemoteStorageGetPublishedFileDetailsResult] {
	handle := iSteamRemoteStorage_GetPublishedFileDetails(uintptr(s), PublishedFileId, MaxSecondsOld)
	return CallResult[RemoteStorageGetPublishedFileDetailsResult]{Handle: handle}
}

func (s steamRemoteStorage) DeletePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageDeletePublishedFileResult] {
	handle := iSteamRemoteStorage_DeletePublishedFile(uintptr(s), PublishedFileId)
	return CallResult[RemoteStorageDeletePublishedFileResult]{Handle: handle}
}

func (s steamRemoteStorage) EnumerateUserPublishedFiles(StartIndex uint32) CallResult[RemoteStorageEnumerateUserPublishedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserPublishedFiles(uintptr(s), StartIndex)
	return CallResult[RemoteStorageEnumerateUserPublishedFilesResult]{Handle: handle}
}

func (s steamRemoteStorage) SubscribePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageSubscribePublishedFileResult] {
	handle := iSteamRemoteStorage_SubscribePublishedFile(uintptr(s), PublishedFileId)
	return CallResult[RemoteStorageSubscribePublishedFileResult]{Handle: handle}
}

func (s steamRemoteStorage) EnumerateUserSubscribedFiles(StartIndex uint32) CallResult[RemoteStorageEnumerateUserSubscribedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserSubscribedFiles(uintptr(s), StartIndex)
	return CallResult[RemoteStorageEnumerateUserSubscribedFilesResult]{Handle: handle}
}

func (s steamRemoteStorage) UnsubscribePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageUnsubscribePublishedFileResult] {
	handle := iSteamRemoteStorage_UnsubscribePublishedFile(uintptr(s), PublishedFileId)
	return CallResult[RemoteStorageUnsubscribePublishedFileResult]{Handle: handle}
}

func (s steamRemoteStorage) UpdatePublishedFileSetChangeDescription(updateHandle PublishedFileUpdateHandle, ChangeDescription string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription(uintptr(s), updateHandle, ChangeDescription)
}

func (s steamRemoteStorage) GetPublishedItemVoteDetails(PublishedFileId PublishedFileId) CallResult[RemoteStorageGetPublishedItemVoteDetailsResult] {
	handle := iSteamRemoteStorage_GetPublishedItemVoteDetails(uintptr(s), PublishedFileId)
	return CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]{Handle: handle}
}

func (s steamRemoteStorage) UpdateUserPublishedItemVote(PublishedFileId PublishedFileId, VoteUp bool) CallResult[RemoteStorageUpdateUserPublishedItemVoteResult] {
	handle := iSteamRemoteStorage_UpdateUserPublishedItemVote(uintptr(s), PublishedFileId, VoteUp)
	return CallResult[RemoteStorageUpdateUserPublishedItemVoteResult]{Handle: handle}
}

func (s steamRemoteStorage) GetUserPublishedItemVoteDetails(PublishedFileId PublishedFileId) CallResult[RemoteStorageGetPublishedItemVoteDetailsResult] {
	handle := iSteamRemoteStorage_GetUserPublishedItemVoteDetails(uintptr(s), PublishedFileId)
	return CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]{Handle: handle}
}

func (s steamRemoteStorage) EnumerateUserSharedWorkshopFiles(steamId Uint64SteamID, StartIndex uint32, RequiredTags SteamParamStringArray, ExcludedTags SteamParamStringArray) CallResult[RemoteStorageEnumerateUserPublishedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles(uintptr(s), steamId, StartIndex, structToUintptr[SteamParamStringArray](&RequiredTags), structToUintptr[SteamParamStringArray](&ExcludedTags))
	return CallResult[RemoteStorageEnumerateUserPublishedFilesResult]{Handle: handle}
}

func (s steamRemoteStorage) PublishVideo(VideoProvider EWorkshopVideoProvider, VideoAccount string, VideoIdentifier string, PreviewFile string, ConsumerAppId AppId_t, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray) CallResult[RemoteStoragePublishFileProgress] {
	handle := iSteamRemoteStorage_PublishVideo(uintptr(s), VideoProvider, VideoAccount, VideoIdentifier, PreviewFile, ConsumerAppId, Title, Description, Visibility, structToUintptr[SteamParamStringArray](&Tags))
	return CallResult[RemoteStoragePublishFileProgress]{Handle: handle}
}

func (s steamRemoteStorage) SetUserPublishedFileAction(PublishedFileId PublishedFileId, Action EWorkshopFileAction) CallResult[RemoteStorageSetUserPublishedFileActionResult] {
	handle := iSteamRemoteStorage_SetUserPublishedFileAction(uintptr(s), PublishedFileId, Action)
	return CallResult[RemoteStorageSetUserPublishedFileActionResult]{Handle: handle}
}

func (s steamRemoteStorage) EnumeratePublishedFilesByUserAction(Action EWorkshopFileAction, StartIndex uint32) CallResult[RemoteStorageEnumeratePublishedFilesByUserActionResult] {
	handle := iSteamRemoteStorage_EnumeratePublishedFilesByUserAction(uintptr(s), Action, StartIndex)
	return CallResult[RemoteStorageEnumeratePublishedFilesByUserActionResult]{Handle: handle}
}

func (s steamRemoteStorage) EnumeratePublishedWorkshopFiles(EnumerationType EWorkshopEnumerationType, StartIndex uint32, Count uint32, Days uint32) (Tags SteamParamStringArray, UserTags SteamParamStringArray, apiHandle CallResult[RemoteStorageEnumerateWorkshopFilesResult]) {
	handle := iSteamRemoteStorage_EnumeratePublishedWorkshopFiles(uintptr(s), EnumerationType, StartIndex, Count, Days, &Tags, &UserTags)
	return Tags, UserTags, CallResult[RemoteStorageEnumerateWorkshopFilesResult]{Handle: handle}
}

func (s steamRemoteStorage) UGCDownloadToLocation(Content UGCHandle, Location string, Priority uint32) CallResult[RemoteStorageDownloadUGCResult] {
	handle := iSteamRemoteStorage_UGCDownloadToLocation(uintptr(s), Content, Location, Priority)
	return CallResult[RemoteStorageDownloadUGCResult]{Handle: handle}
}

func (s steamRemoteStorage) GetLocalFileChangeCount() int32 {
	return iSteamRemoteStorage_GetLocalFileChangeCount(uintptr(s))
}

func (s steamRemoteStorage) GetLocalFileChange(FileIndex int32) (ChangeType ERemoteStorageLocalFileChange, FilePathType ERemoteStorageFilePathType, result []byte) {
	result = iSteamRemoteStorage_GetLocalFileChange(uintptr(s), FileIndex, &ChangeType, &FilePathType)
	return ChangeType, FilePathType, result
}

func (s steamRemoteStorage) BeginFileWriteBatch() bool {
	return iSteamRemoteStorage_BeginFileWriteBatch(uintptr(s))
}

func (s steamRemoteStorage) EndFileWriteBatch() bool {
	return iSteamRemoteStorage_EndFileWriteBatch(uintptr(s))
}

// Steam Screenshots
const (
	ScreenshotMaxTaggedUsers          uint32 = 32
	ScreenshotMaxTaggedPublishedFiles uint32 = 32
	UFSTagTypeMax                     int32  = 255
	UFSTagValueMax                    int32  = 255
	ScreenshotThumbWidth              int32  = 200
)

const (
	flatAPI_SteamScreenshots                           = "SteamAPI_SteamScreenshots_v003"
	flatAPI_ISteamScreenshots_WriteScreenshot          = "SteamAPI_ISteamScreenshots_WriteScreenshot"
	flatAPI_ISteamScreenshots_AddScreenshotToLibrary   = "SteamAPI_ISteamScreenshots_AddScreenshotToLibrary"
	flatAPI_ISteamScreenshots_TriggerScreenshot        = "SteamAPI_ISteamScreenshots_TriggerScreenshot"
	flatAPI_ISteamScreenshots_HookScreenshots          = "SteamAPI_ISteamScreenshots_HookScreenshots"
	flatAPI_ISteamScreenshots_SetLocation              = "SteamAPI_ISteamScreenshots_SetLocation"
	flatAPI_ISteamScreenshots_TagUser                  = "SteamAPI_ISteamScreenshots_TagUser"
	flatAPI_ISteamScreenshots_TagPublishedFile         = "SteamAPI_ISteamScreenshots_TagPublishedFile"
	flatAPI_ISteamScreenshots_IsScreenshotsHooked      = "SteamAPI_ISteamScreenshots_IsScreenshotsHooked"
	flatAPI_ISteamScreenshots_AddVRScreenshotToLibrary = "SteamAPI_ISteamScreenshots_AddVRScreenshotToLibrary"
)

type ScreenshotHandle uint

type EVRScreenshotType int32

const (
	EVRScreenshotType_None           EVRScreenshotType = 0
	EVRScreenshotType_Mono           EVRScreenshotType = 1
	EVRScreenshotType_Stereo         EVRScreenshotType = 2
	EVRScreenshotType_MonoCubemap    EVRScreenshotType = 3
	EVRScreenshotType_MonoPanorama   EVRScreenshotType = 4
	EVRScreenshotType_StereoPanorama EVRScreenshotType = 5
)

var (
	steamScreenshots_init                      func() uintptr
	iSteamScreenshots_WriteScreenshot          func(steamScreenshots uintptr, pubRGB []byte, cubRGB uint32, nWidth int32, nHeight int32) ScreenshotHandle
	iSteamScreenshots_AddScreenshotToLibrary   func(steamScreenshots uintptr, pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle
	iSteamScreenshots_TriggerScreenshot        func(steamScreenshots uintptr)
	iSteamScreenshots_HookScreenshots          func(steamScreenshots uintptr, bHook bool)
	iSteamScreenshots_SetLocation              func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, pchLocation string) bool
	iSteamScreenshots_TagUser                  func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool
	iSteamScreenshots_TagPublishedFile         func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool
	iSteamScreenshots_IsScreenshotsHooked      func(steamScreenshots uintptr) bool
	iSteamScreenshots_AddVRScreenshotToLibrary func(steamScreenshots uintptr, eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle
)

type steamScreenshots uintptr

func SteamScreenshots() ISteamScreenshots {
	return steamScreenshots(steamScreenshots_init())
}

func (s steamScreenshots) WriteScreenshot(pubRGB []byte, nWidth int32, nHeight int32) ScreenshotHandle {
	return iSteamScreenshots_WriteScreenshot(uintptr(s), pubRGB, uint32(len(pubRGB)), nWidth, nHeight)
}

func (s steamScreenshots) AddScreenshotToLibrary(pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle {
	return iSteamScreenshots_AddScreenshotToLibrary(uintptr(s), pchFilename, pchThumbnailFilename, nWidth, nHeight)
}

func (s steamScreenshots) TriggerScreenshot() {
	iSteamScreenshots_TriggerScreenshot(uintptr(s))
}

func (s steamScreenshots) HookScreenshots(bHook bool) {
	iSteamScreenshots_HookScreenshots(uintptr(s), bHook)
}

func (s steamScreenshots) SetLocation(hScreenshot ScreenshotHandle, pchLocation string) bool {
	return iSteamScreenshots_SetLocation(uintptr(s), hScreenshot, pchLocation)
}

func (s steamScreenshots) TagUser(hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool {
	return iSteamScreenshots_TagUser(uintptr(s), hScreenshot, steamID)
}

func (s steamScreenshots) TagPublishedFile(hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool {
	return iSteamScreenshots_TagPublishedFile(uintptr(s), hScreenshot, unPublishedFileID)
}

func (s steamScreenshots) IsScreenshotsHooked() bool {
	return iSteamScreenshots_IsScreenshotsHooked(uintptr(s))
}

func (s steamScreenshots) AddVRScreenshotToLibrary(eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle {
	return iSteamScreenshots_AddVRScreenshotToLibrary(uintptr(s), eType, pchFilename, pchVRFilename)
}

// Steam UGC
const (
	UGCQueryHandleInvalid  UGCQueryHandle  = 0xffffffffffffffff
	UGCUpdateHandleInvalid UGCUpdateHandle = 0xffffffffffffffff
	NumUGCResultsPerPage   uint32          = 50
	DeveloperMetadataMax   uint32          = 5000
)

type UGCQueryHandle uint64
type UGCUpdateHandle uint64
type EUGCContentDescriptorID int32

const (
	EUGCContentDescriptor_NudityOrSexualContent   EUGCContentDescriptorID = 1
	EUGCContentDescriptor_FrequentViolenceOrGore  EUGCContentDescriptorID = 2
	EUGCContentDescriptor_AdultOnlySexualContent  EUGCContentDescriptorID = 3
	EUGCContentDescriptor_GratuitousSexualContent EUGCContentDescriptorID = 4
	EUGCContentDescriptor_AnyMatureContent        EUGCContentDescriptorID = 5
)

type EUserUGCList int32

const (
	EUserUGCList_Published     EUserUGCList = 0
	EUserUGCList_VotedOn       EUserUGCList = 1
	EUserUGCList_VotedUp       EUserUGCList = 2
	EUserUGCList_VotedDown     EUserUGCList = 3
	EUserUGCList_WillVoteLater EUserUGCList = 4
	EUserUGCList_Favorited     EUserUGCList = 5
	EUserUGCList_Subscribed    EUserUGCList = 6
	EUserUGCList_UsedOrPlayed  EUserUGCList = 7
	EUserUGCList_Followed      EUserUGCList = 8
)

type EUGCQuery int32

const (
	EUGCQuery_RankedByVote                                  EUGCQuery = 0
	EUGCQuery_RankedByPublicationDate                       EUGCQuery = 1
	EUGCQuery_AcceptedForGameRankedByAcceptanceDate         EUGCQuery = 2
	EUGCQuery_RankedByTrend                                 EUGCQuery = 3
	EUGCQuery_FavoritedByFriendsRankedByPublicationDate     EUGCQuery = 4
	EUGCQuery_CreatedByFriendsRankedByPublicationDate       EUGCQuery = 5
	EUGCQuery_RankedByNumTimesReported                      EUGCQuery = 6
	EUGCQuery_CreatedByFollowedUsersRankedByPublicationDate EUGCQuery = 7
	EUGCQuery_NotYetRated                                   EUGCQuery = 8
	EUGCQuery_RankedByTotalVotesAsc                         EUGCQuery = 9
	EUGCQuery_RankedByVotesUp                               EUGCQuery = 10
	EUGCQuery_RankedByTextSearch                            EUGCQuery = 11
	EUGCQuery_RankedByTotalUniqueSubscriptions              EUGCQuery = 12
	EUGCQuery_RankedByPlaytimeTrend                         EUGCQuery = 13
	EUGCQuery_RankedByTotalPlaytime                         EUGCQuery = 14
	EUGCQuery_RankedByAveragePlaytimeTrend                  EUGCQuery = 15
	EUGCQuery_RankedByLifetimeAveragePlaytime               EUGCQuery = 16
	EUGCQuery_RankedByPlaytimeSessionsTrend                 EUGCQuery = 17
	EUGCQuery_RankedByLifetimePlaytimeSessions              EUGCQuery = 18
	EUGCQuery_RankedByLastUpdatedDate                       EUGCQuery = 19
)

type EItemState int32

const (
	EItemState_None            EItemState = 0
	EItemState_Subscribed      EItemState = 1
	EItemState_LegacyItem      EItemState = 2
	EItemState_Installed       EItemState = 4
	EItemState_NeedsUpdate     EItemState = 8
	EItemState_Downloading     EItemState = 16
	EItemState_DownloadPending EItemState = 32
	EItemState_DisabledLocally EItemState = 64
)

type SteamUGCDetails struct {
	PublishedFileId     PublishedFileId
	Result              EResult
	FileType            EWorkshopFileType
	CreatorAppID        AppId_t
	ConsumerAppID       AppId_t
	Title               [129]byte
	Description         [8000]byte
	SteamIDOwner        uint64
	TimeCreated         uint32
	TimeUpdated         uint32
	TimeAddedToUserList uint32
	Visibility          ERemoteStoragePublishedFileVisibility
	Banned              bool
	AcceptedForUse      bool
	TagsTruncated       bool
	Tags                [1025]byte
	File                UGCHandle
	PreviewFile         UGCHandle
	FileName            [260]byte
	FileSize            int32
	PreviewFileSize     int32
	URL                 [256]byte
	VotesUp             uint32
	VotesDown           uint32
	Score               float32
	NumChildren         uint32
	TotalFilesSize      uint64
}

type EUGCMatchingUGCType int32

const (
	EUGCMatchingUGCType_Items              EUGCMatchingUGCType = 0
	EUGCMatchingUGCType_ItemsMtx           EUGCMatchingUGCType = 1
	EUGCMatchingUGCType_ItemsReadyToUse    EUGCMatchingUGCType = 2
	EUGCMatchingUGCType_Collections        EUGCMatchingUGCType = 3
	EUGCMatchingUGCType_Artwork            EUGCMatchingUGCType = 4
	EUGCMatchingUGCType_Videos             EUGCMatchingUGCType = 5
	EUGCMatchingUGCType_Screenshots        EUGCMatchingUGCType = 6
	EUGCMatchingUGCType_AllGuides          EUGCMatchingUGCType = 7
	EUGCMatchingUGCType_WebGuides          EUGCMatchingUGCType = 8
	EUGCMatchingUGCType_IntegratedGuides   EUGCMatchingUGCType = 9
	EUGCMatchingUGCType_UsableInGame       EUGCMatchingUGCType = 10
	EUGCMatchingUGCType_ControllerBindings EUGCMatchingUGCType = 11
	EUGCMatchingUGCType_GameManagedItems   EUGCMatchingUGCType = 12
	EUGCMatchingUGCType_All                EUGCMatchingUGCType = -1
)

type EUserUGCListSortOrder int32

const (
	EUserUGCList_SortOrder_CreationOrderDesc    EUserUGCListSortOrder = 0
	EUserUGCList_SortOrder_CreationOrderAsc     EUserUGCListSortOrder = 1
	EUserUGCList_SortOrder_TitleAsc             EUserUGCListSortOrder = 2
	EUserUGCList_SortOrder_LastUpdatedDesc      EUserUGCListSortOrder = 3
	EUserUGCList_SortOrder_SubscriptionDateDesc EUserUGCListSortOrder = 4
	EUserUGCList_SortOrder_VoteScoreDesc        EUserUGCListSortOrder = 5
	EUserUGCList_SortOrder_ForModeration        EUserUGCListSortOrder = 6
)

type EItemPreviewType int32

const (
	EItemPreviewType_Image                         EItemPreviewType = 0
	EItemPreviewType_YouTubeVideo                  EItemPreviewType = 1
	EItemPreviewType_Sketchfab                     EItemPreviewType = 2
	EItemPreviewType_EnvironmentMapHorizontalCross EItemPreviewType = 3
	EItemPreviewType_EnvironmentMapLatLong         EItemPreviewType = 4
	EItemPreviewType_Clip                          EItemPreviewType = 5
	EItemPreviewType_ReservedMax                   EItemPreviewType = 255
)

type EItemUpdateStatus int32

const (
	EItemUpdateStatus_Invalid              EItemUpdateStatus = 0
	EItemUpdateStatus_PreparingConfig      EItemUpdateStatus = 1
	EItemUpdateStatus_PreparingContent     EItemUpdateStatus = 2
	EItemUpdateStatus_UploadingContent     EItemUpdateStatus = 3
	EItemUpdateStatus_UploadingPreviewFile EItemUpdateStatus = 4
	EItemUpdateStatus_CommittingChanges    EItemUpdateStatus = 5
)

type EItemStatistic int32

const (
	EItemStatistic_NumSubscriptions                    EItemStatistic = 0
	EItemStatistic_NumFavorites                        EItemStatistic = 1
	EItemStatistic_NumFollowers                        EItemStatistic = 2
	EItemStatistic_NumUniqueSubscriptions              EItemStatistic = 3
	EItemStatistic_NumUniqueFavorites                  EItemStatistic = 4
	EItemStatistic_NumUniqueFollowers                  EItemStatistic = 5
	EItemStatistic_NumUniqueWebsiteViews               EItemStatistic = 6
	EItemStatistic_ReportScore                         EItemStatistic = 7
	EItemStatistic_NumSecondsPlayed                    EItemStatistic = 8
	EItemStatistic_NumPlaytimeSessions                 EItemStatistic = 9
	EItemStatistic_NumComments                         EItemStatistic = 10
	EItemStatistic_NumSecondsPlayedDuringTimePeriod    EItemStatistic = 11
	EItemStatistic_NumPlaytimeSessionsDuringTimePeriod EItemStatistic = 12
)

const (
	flatAPI_SteamUGC                                      = "SteamAPI_SteamUGC_v020"
	flatAPI_ISteamUGC_CreateQueryUserUGCRequest           = "SteamAPI_ISteamUGC_CreateQueryUserUGCRequest"
	flatAPI_ISteamUGC_CreateQueryAllUGCRequestPage        = "SteamAPI_ISteamUGC_CreateQueryAllUGCRequestPage"
	flatAPI_ISteamUGC_CreateQueryAllUGCRequestCursor      = "SteamAPI_ISteamUGC_CreateQueryAllUGCRequestCursor"
	flatAPI_ISteamUGC_CreateQueryUGCDetailsRequest        = "SteamAPI_ISteamUGC_CreateQueryUGCDetailsRequest"
	flatAPI_ISteamUGC_SendQueryUGCRequest                 = "SteamAPI_ISteamUGC_SendQueryUGCRequest"
	flatAPI_ISteamUGC_GetQueryUGCResult                   = "SteamAPI_ISteamUGC_GetQueryUGCResult"
	flatAPI_ISteamUGC_GetQueryUGCNumTags                  = "SteamAPI_ISteamUGC_GetQueryUGCNumTags"
	flatAPI_ISteamUGC_GetQueryUGCTag                      = "SteamAPI_ISteamUGC_GetQueryUGCTag"
	flatAPI_ISteamUGC_GetQueryUGCTagDisplayName           = "SteamAPI_ISteamUGC_GetQueryUGCTagDisplayName"
	flatAPI_ISteamUGC_GetQueryUGCPreviewURL               = "SteamAPI_ISteamUGC_GetQueryUGCPreviewURL"
	flatAPI_ISteamUGC_GetQueryUGCMetadata                 = "SteamAPI_ISteamUGC_GetQueryUGCMetadata"
	flatAPI_ISteamUGC_GetQueryUGCChildren                 = "SteamAPI_ISteamUGC_GetQueryUGCChildren"
	flatAPI_ISteamUGC_GetQueryUGCStatistic                = "SteamAPI_ISteamUGC_GetQueryUGCStatistic"
	flatAPI_ISteamUGC_GetQueryUGCNumAdditionalPreviews    = "SteamAPI_ISteamUGC_GetQueryUGCNumAdditionalPreviews"
	flatAPI_ISteamUGC_GetQueryUGCAdditionalPreview        = "SteamAPI_ISteamUGC_GetQueryUGCAdditionalPreview"
	flatAPI_ISteamUGC_GetQueryUGCNumKeyValueTags          = "SteamAPI_ISteamUGC_GetQueryUGCNumKeyValueTags"
	flatAPI_ISteamUGC_GetQueryUGCKeyValueTag              = "SteamAPI_ISteamUGC_GetQueryUGCKeyValueTag"
	flatAPI_ISteamUGC_GetQueryFirstUGCKeyValueTag         = "SteamAPI_ISteamUGC_GetQueryFirstUGCKeyValueTag"
	flatAPI_ISteamUGC_GetNumSupportedGameVersions         = "SteamAPI_ISteamUGC_GetNumSupportedGameVersions"
	flatAPI_ISteamUGC_GetSupportedGameVersionData         = "SteamAPI_ISteamUGC_GetSupportedGameVersionData"
	flatAPI_ISteamUGC_GetQueryUGCContentDescriptors       = "SteamAPI_ISteamUGC_GetQueryUGCContentDescriptors"
	flatAPI_ISteamUGC_ReleaseQueryUGCRequest              = "SteamAPI_ISteamUGC_ReleaseQueryUGCRequest"
	flatAPI_ISteamUGC_AddRequiredTag                      = "SteamAPI_ISteamUGC_AddRequiredTag"
	flatAPI_ISteamUGC_AddRequiredTagGroup                 = "SteamAPI_ISteamUGC_AddRequiredTagGroup"
	flatAPI_ISteamUGC_AddExcludedTag                      = "SteamAPI_ISteamUGC_AddExcludedTag"
	flatAPI_ISteamUGC_SetReturnOnlyIDs                    = "SteamAPI_ISteamUGC_SetReturnOnlyIDs"
	flatAPI_ISteamUGC_SetReturnKeyValueTags               = "SteamAPI_ISteamUGC_SetReturnKeyValueTags"
	flatAPI_ISteamUGC_SetReturnLongDescription            = "SteamAPI_ISteamUGC_SetReturnLongDescription"
	flatAPI_ISteamUGC_SetReturnMetadata                   = "SteamAPI_ISteamUGC_SetReturnMetadata"
	flatAPI_ISteamUGC_SetReturnChildren                   = "SteamAPI_ISteamUGC_SetReturnChildren"
	flatAPI_ISteamUGC_SetReturnAdditionalPreviews         = "SteamAPI_ISteamUGC_SetReturnAdditionalPreviews"
	flatAPI_ISteamUGC_SetReturnTotalOnly                  = "SteamAPI_ISteamUGC_SetReturnTotalOnly"
	flatAPI_ISteamUGC_SetReturnPlaytimeStats              = "SteamAPI_ISteamUGC_SetReturnPlaytimeStats"
	flatAPI_ISteamUGC_SetLanguage                         = "SteamAPI_ISteamUGC_SetLanguage"
	flatAPI_ISteamUGC_SetAllowCachedResponse              = "SteamAPI_ISteamUGC_SetAllowCachedResponse"
	flatAPI_ISteamUGC_SetAdminQuery                       = "SteamAPI_ISteamUGC_SetAdminQuery"
	flatAPI_ISteamUGC_SetCloudFileNameFilter              = "SteamAPI_ISteamUGC_SetCloudFileNameFilter"
	flatAPI_ISteamUGC_SetMatchAnyTag                      = "SteamAPI_ISteamUGC_SetMatchAnyTag"
	flatAPI_ISteamUGC_SetSearchText                       = "SteamAPI_ISteamUGC_SetSearchText"
	flatAPI_ISteamUGC_SetRankedByTrendDays                = "SteamAPI_ISteamUGC_SetRankedByTrendDays"
	flatAPI_ISteamUGC_SetTimeCreatedDateRange             = "SteamAPI_ISteamUGC_SetTimeCreatedDateRange"
	flatAPI_ISteamUGC_SetTimeUpdatedDateRange             = "SteamAPI_ISteamUGC_SetTimeUpdatedDateRange"
	flatAPI_ISteamUGC_AddRequiredKeyValueTag              = "SteamAPI_ISteamUGC_AddRequiredKeyValueTag"
	flatAPI_ISteamUGC_CreateItem                          = "SteamAPI_ISteamUGC_CreateItem"
	flatAPI_ISteamUGC_StartItemUpdate                     = "SteamAPI_ISteamUGC_StartItemUpdate"
	flatAPI_ISteamUGC_SetItemTitle                        = "SteamAPI_ISteamUGC_SetItemTitle"
	flatAPI_ISteamUGC_SetItemDescription                  = "SteamAPI_ISteamUGC_SetItemDescription"
	flatAPI_ISteamUGC_SetItemUpdateLanguage               = "SteamAPI_ISteamUGC_SetItemUpdateLanguage"
	flatAPI_ISteamUGC_SetItemMetadata                     = "SteamAPI_ISteamUGC_SetItemMetadata"
	flatAPI_ISteamUGC_SetItemVisibility                   = "SteamAPI_ISteamUGC_SetItemVisibility"
	flatAPI_ISteamUGC_SetItemTags                         = "SteamAPI_ISteamUGC_SetItemTags"
	flatAPI_ISteamUGC_SetItemContent                      = "SteamAPI_ISteamUGC_SetItemContent"
	flatAPI_ISteamUGC_SetItemPreview                      = "SteamAPI_ISteamUGC_SetItemPreview"
	flatAPI_ISteamUGC_SetAllowLegacyUpload                = "SteamAPI_ISteamUGC_SetAllowLegacyUpload"
	flatAPI_ISteamUGC_RemoveAllItemKeyValueTags           = "SteamAPI_ISteamUGC_RemoveAllItemKeyValueTags"
	flatAPI_ISteamUGC_RemoveItemKeyValueTags              = "SteamAPI_ISteamUGC_RemoveItemKeyValueTags"
	flatAPI_ISteamUGC_AddItemKeyValueTag                  = "SteamAPI_ISteamUGC_AddItemKeyValueTag"
	flatAPI_ISteamUGC_AddItemPreviewFile                  = "SteamAPI_ISteamUGC_AddItemPreviewFile"
	flatAPI_ISteamUGC_AddItemPreviewVideo                 = "SteamAPI_ISteamUGC_AddItemPreviewVideo"
	flatAPI_ISteamUGC_UpdateItemPreviewFile               = "SteamAPI_ISteamUGC_UpdateItemPreviewFile"
	flatAPI_ISteamUGC_UpdateItemPreviewVideo              = "SteamAPI_ISteamUGC_UpdateItemPreviewVideo"
	flatAPI_ISteamUGC_RemoveItemPreview                   = "SteamAPI_ISteamUGC_RemoveItemPreview"
	flatAPI_ISteamUGC_AddContentDescriptor                = "SteamAPI_ISteamUGC_AddContentDescriptor"
	flatAPI_ISteamUGC_RemoveContentDescriptor             = "SteamAPI_ISteamUGC_RemoveContentDescriptor"
	flatAPI_ISteamUGC_SetRequiredGameVersions             = "SteamAPI_ISteamUGC_SetRequiredGameVersions"
	flatAPI_ISteamUGC_SubmitItemUpdate                    = "SteamAPI_ISteamUGC_SubmitItemUpdate"
	flatAPI_ISteamUGC_GetItemUpdateProgress               = "SteamAPI_ISteamUGC_GetItemUpdateProgress"
	flatAPI_ISteamUGC_SetUserItemVote                     = "SteamAPI_ISteamUGC_SetUserItemVote"
	flatAPI_ISteamUGC_GetUserItemVote                     = "SteamAPI_ISteamUGC_GetUserItemVote"
	flatAPI_ISteamUGC_AddItemToFavorites                  = "SteamAPI_ISteamUGC_AddItemToFavorites"
	flatAPI_ISteamUGC_RemoveItemFromFavorites             = "SteamAPI_ISteamUGC_RemoveItemFromFavorites"
	flatAPI_ISteamUGC_SubscribeItem                       = "SteamAPI_ISteamUGC_SubscribeItem"
	flatAPI_ISteamUGC_UnsubscribeItem                     = "SteamAPI_ISteamUGC_UnsubscribeItem"
	flatAPI_ISteamUGC_GetNumSubscribedItems               = "SteamAPI_ISteamUGC_GetNumSubscribedItems"
	flatAPI_ISteamUGC_GetSubscribedItems                  = "SteamAPI_ISteamUGC_GetSubscribedItems"
	flatAPI_ISteamUGC_GetItemState                        = "SteamAPI_ISteamUGC_GetItemState"
	flatAPI_ISteamUGC_GetItemInstallInfo                  = "SteamAPI_ISteamUGC_GetItemInstallInfo"
	flatAPI_ISteamUGC_GetItemDownloadInfo                 = "SteamAPI_ISteamUGC_GetItemDownloadInfo"
	flatAPI_ISteamUGC_DownloadItem                        = "SteamAPI_ISteamUGC_DownloadItem"
	flatAPI_ISteamUGC_BInitWorkshopForGameServer          = "SteamAPI_ISteamUGC_BInitWorkshopForGameServer"
	flatAPI_ISteamUGC_SuspendDownloads                    = "SteamAPI_ISteamUGC_SuspendDownloads"
	flatAPI_ISteamUGC_StartPlaytimeTracking               = "SteamAPI_ISteamUGC_StartPlaytimeTracking"
	flatAPI_ISteamUGC_StopPlaytimeTracking                = "SteamAPI_ISteamUGC_StopPlaytimeTracking"
	flatAPI_ISteamUGC_StopPlaytimeTrackingForAllItems     = "SteamAPI_ISteamUGC_StopPlaytimeTrackingForAllItems"
	flatAPI_ISteamUGC_AddDependency                       = "SteamAPI_ISteamUGC_AddDependency"
	flatAPI_ISteamUGC_RemoveDependency                    = "SteamAPI_ISteamUGC_RemoveDependency"
	flatAPI_ISteamUGC_AddAppDependency                    = "SteamAPI_ISteamUGC_AddAppDependency"
	flatAPI_ISteamUGC_RemoveAppDependency                 = "SteamAPI_ISteamUGC_RemoveAppDependency"
	flatAPI_ISteamUGC_GetAppDependencies                  = "SteamAPI_ISteamUGC_GetAppDependencies"
	flatAPI_ISteamUGC_DeleteItem                          = "SteamAPI_ISteamUGC_DeleteItem"
	flatAPI_ISteamUGC_ShowWorkshopEULA                    = "SteamAPI_ISteamUGC_ShowWorkshopEULA"
	flatAPI_ISteamUGC_GetWorkshopEULAStatus               = "SteamAPI_ISteamUGC_GetWorkshopEULAStatus"
	flatAPI_ISteamUGC_GetUserContentDescriptorPreferences = "SteamAPI_ISteamUGC_GetUserContentDescriptorPreferences"
)

var (
	steamUGC_init                                 func() uintptr
	iSteamUGC_CreateQueryUserUGCRequest           func(steamUGC uintptr, unAccountID AccountID, eListType EUserUGCList, eMatchingUGCType EUGCMatchingUGCType, eSortOrder EUserUGCListSortOrder, nCreatorAppID AppId_t, nConsumerAppID AppId_t, unPage uint32) UGCQueryHandle
	iSteamUGC_CreateQueryAllUGCRequestPage        func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId_t, nConsumerAppID AppId_t, unPage uint32) UGCQueryHandle
	iSteamUGC_CreateQueryAllUGCRequestCursor      func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId_t, nConsumerAppID AppId_t, pchCursor string) UGCQueryHandle
	iSteamUGC_CreateQueryUGCDetailsRequest        func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) UGCQueryHandle
	iSteamUGC_SendQueryUGCRequest                 func(steamUGC uintptr, handle UGCQueryHandle) SteamAPICall
	iSteamUGC_GetQueryUGCResult                   func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pDetails *SteamUGCDetails) bool
	iSteamUGC_GetQueryUGCNumTags                  func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCTag                      func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue []byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryUGCTagDisplayName           func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue []byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryUGCPreviewURL               func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchURL []byte, cchURLSize uint32) bool
	iSteamUGC_GetQueryUGCMetadata                 func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchMetadata []byte, cchMetadatasize uint32) bool
	iSteamUGC_GetQueryUGCChildren                 func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecPublishedFileID []PublishedFileId, cMaxEntries uint32) bool
	iSteamUGC_GetQueryUGCStatistic                func(steamUGC uintptr, handle UGCQueryHandle, index uint32, eStatType EItemStatistic, pStatValue *uint64) bool
	iSteamUGC_GetQueryUGCNumAdditionalPreviews    func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCAdditionalPreview        func(steamUGC uintptr, handle UGCQueryHandle, index uint32, previewIndex uint32, pchURLOrVideoID []byte, cchURLSize uint32, pchOriginalFileName []byte, cchOriginalFileNameSize uint32, pPreviewType *EItemPreviewType) bool
	iSteamUGC_GetQueryUGCNumKeyValueTags          func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCKeyValueTag              func(steamUGC uintptr, handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, pchKey []byte, cchKeySize uint32, pchValue []byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryFirstUGCKeyValueTag         func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchKey string, pchValue []byte, cchValueSize uint32) bool
	iSteamUGC_GetNumSupportedGameVersions         func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetSupportedGameVersionData         func(steamUGC uintptr, handle UGCQueryHandle, index uint32, versionIndex uint32, pchGameBranchMin []byte, pchGameBranchMax []byte, cchGameBranchSize uint32) bool
	iSteamUGC_GetQueryUGCContentDescriptors       func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecDescriptors []EUGCContentDescriptorID, cMaxEntries uint32) uint32
	iSteamUGC_ReleaseQueryUGCRequest              func(steamUGC uintptr, handle UGCQueryHandle) bool
	iSteamUGC_AddRequiredTag                      func(steamUGC uintptr, handle UGCQueryHandle, pTagName string) bool
	iSteamUGC_AddRequiredTagGroup                 func(steamUGC uintptr, handle UGCQueryHandle, pTagGroups []SteamParamStringArray) bool
	iSteamUGC_AddExcludedTag                      func(steamUGC uintptr, handle UGCQueryHandle, pTagName string) bool
	iSteamUGC_SetReturnOnlyIDs                    func(steamUGC uintptr, handle UGCQueryHandle, bReturnOnlyIDs bool) bool
	iSteamUGC_SetReturnKeyValueTags               func(steamUGC uintptr, handle UGCQueryHandle, bReturnKeyValueTags bool) bool
	iSteamUGC_SetReturnLongDescription            func(steamUGC uintptr, handle UGCQueryHandle, bReturnLongDescription bool) bool
	iSteamUGC_SetReturnMetadata                   func(steamUGC uintptr, handle UGCQueryHandle, bReturnMetadata bool) bool
	iSteamUGC_SetReturnChildren                   func(steamUGC uintptr, handle UGCQueryHandle, bReturnChildren bool) bool
	iSteamUGC_SetReturnAdditionalPreviews         func(steamUGC uintptr, handle UGCQueryHandle, bReturnAdditionalPreviews bool) bool
	iSteamUGC_SetReturnTotalOnly                  func(steamUGC uintptr, handle UGCQueryHandle, bReturnTotalOnly bool) bool
	iSteamUGC_SetReturnPlaytimeStats              func(steamUGC uintptr, handle UGCQueryHandle, unDays uint32) bool
	iSteamUGC_SetLanguage                         func(steamUGC uintptr, handle UGCQueryHandle, pchLanguage string) bool
	iSteamUGC_SetAllowCachedResponse              func(steamUGC uintptr, handle UGCQueryHandle, unMaxAgeSeconds uint32) bool
	iSteamUGC_SetAdminQuery                       func(steamUGC uintptr, handle UGCUpdateHandle, bAdminQuery bool) bool
	iSteamUGC_SetCloudFileNameFilter              func(steamUGC uintptr, handle UGCQueryHandle, pMatchCloudFileName string) bool
	iSteamUGC_SetMatchAnyTag                      func(steamUGC uintptr, handle UGCQueryHandle, bMatchAnyTag bool) bool
	iSteamUGC_SetSearchText                       func(steamUGC uintptr, handle UGCQueryHandle, pSearchText string) bool
	iSteamUGC_SetRankedByTrendDays                func(steamUGC uintptr, handle UGCQueryHandle, unDays uint32) bool
	iSteamUGC_SetTimeCreatedDateRange             func(steamUGC uintptr, handle UGCQueryHandle, rtStart RTime32, rtEnd RTime32) bool
	iSteamUGC_SetTimeUpdatedDateRange             func(steamUGC uintptr, handle UGCQueryHandle, rtStart RTime32, rtEnd RTime32) bool
	iSteamUGC_AddRequiredKeyValueTag              func(steamUGC uintptr, handle UGCQueryHandle, pKey string, pValue string) bool
	iSteamUGC_CreateItem                          func(steamUGC uintptr, nConsumerAppId AppId_t, eFileType EWorkshopFileType) SteamAPICall
	iSteamUGC_StartItemUpdate                     func(steamUGC uintptr, nConsumerAppId AppId_t, nPublishedFileID PublishedFileId) UGCUpdateHandle
	iSteamUGC_SetItemTitle                        func(steamUGC uintptr, handle UGCUpdateHandle, pchTitle string) bool
	iSteamUGC_SetItemDescription                  func(steamUGC uintptr, handle UGCUpdateHandle, pchDescription string) bool
	iSteamUGC_SetItemUpdateLanguage               func(steamUGC uintptr, handle UGCUpdateHandle, pchLanguage string) bool
	iSteamUGC_SetItemMetadata                     func(steamUGC uintptr, handle UGCUpdateHandle, pchMetaData string) bool
	iSteamUGC_SetItemVisibility                   func(steamUGC uintptr, handle UGCUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool
	iSteamUGC_SetItemTags                         func(steamUGC uintptr, updateHandle UGCUpdateHandle, pTags []SteamParamStringArray, bAllowAdminTags bool) bool
	iSteamUGC_SetItemContent                      func(steamUGC uintptr, handle UGCUpdateHandle, pszContentFolder string) bool
	iSteamUGC_SetItemPreview                      func(steamUGC uintptr, handle UGCUpdateHandle, pszPreviewFile string) bool
	iSteamUGC_SetAllowLegacyUpload                func(steamUGC uintptr, handle UGCUpdateHandle, bAllowLegacyUpload bool) bool
	iSteamUGC_RemoveAllItemKeyValueTags           func(steamUGC uintptr, handle UGCUpdateHandle) bool
	iSteamUGC_RemoveItemKeyValueTags              func(steamUGC uintptr, handle UGCUpdateHandle, pchKey string) bool
	iSteamUGC_AddItemKeyValueTag                  func(steamUGC uintptr, handle UGCUpdateHandle, pchKey string, pchValue string) bool
	iSteamUGC_AddItemPreviewFile                  func(steamUGC uintptr, handle UGCUpdateHandle, pszPreviewFile string, Type EItemPreviewType) bool
	iSteamUGC_AddItemPreviewVideo                 func(steamUGC uintptr, handle UGCUpdateHandle, pszVideoID string) bool
	iSteamUGC_UpdateItemPreviewFile               func(steamUGC uintptr, handle UGCUpdateHandle, index uint32, pszPreviewFile string) bool
	iSteamUGC_UpdateItemPreviewVideo              func(steamUGC uintptr, handle UGCUpdateHandle, index uint32, pszVideoID string) bool
	iSteamUGC_RemoveItemPreview                   func(steamUGC uintptr, handle UGCUpdateHandle, index uint32) bool
	iSteamUGC_AddContentDescriptor                func(steamUGC uintptr, handle UGCUpdateHandle, descid EUGCContentDescriptorID) bool
	iSteamUGC_RemoveContentDescriptor             func(steamUGC uintptr, handle UGCUpdateHandle, descid EUGCContentDescriptorID) bool
	iSteamUGC_SetRequiredGameVersions             func(steamUGC uintptr, handle UGCUpdateHandle, pszGameBranchMin string, pszGameBranchMax string) bool
	iSteamUGC_SubmitItemUpdate                    func(steamUGC uintptr, handle UGCUpdateHandle, pchChangeNote string) SteamAPICall
	iSteamUGC_GetItemUpdateProgress               func(steamUGC uintptr, handle UGCUpdateHandle, punBytesProcessed *uint64, punBytesTotal *uint64) EItemUpdateStatus
	iSteamUGC_SetUserItemVote                     func(steamUGC uintptr, nPublishedFileID PublishedFileId, bVoteUp bool) SteamAPICall
	iSteamUGC_GetUserItemVote                     func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_AddItemToFavorites                  func(steamUGC uintptr, nAppId AppId_t, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_RemoveItemFromFavorites             func(steamUGC uintptr, nAppId AppId_t, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_SubscribeItem                       func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_UnsubscribeItem                     func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_GetNumSubscribedItems               func(steamUGC uintptr) uint32
	iSteamUGC_GetSubscribedItems                  func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, cMaxEntries uint32) uint32
	iSteamUGC_GetItemState                        func(steamUGC uintptr, nPublishedFileID PublishedFileId) uint32
	iSteamUGC_GetItemInstallInfo                  func(steamUGC uintptr, nPublishedFileID PublishedFileId, punSizeOnDisk *uint64, pchFolder []byte, cchFolderSize uint32, punTimeStamp *uint32) bool
	iSteamUGC_GetItemDownloadInfo                 func(steamUGC uintptr, nPublishedFileID PublishedFileId, punBytesDownloaded *uint64, punBytesTotal *uint64) bool
	iSteamUGC_DownloadItem                        func(steamUGC uintptr, nPublishedFileID PublishedFileId, bHighPriority bool) bool
	iSteamUGC_BInitWorkshopForGameServer          func(steamUGC uintptr, unWorkshopDepotID DepotId, pszFolder string) bool
	iSteamUGC_SuspendDownloads                    func(steamUGC uintptr, bSuspend bool)
	iSteamUGC_StartPlaytimeTracking               func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall
	iSteamUGC_StopPlaytimeTracking                func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall
	iSteamUGC_StopPlaytimeTrackingForAllItems     func(steamUGC uintptr) SteamAPICall
	iSteamUGC_AddDependency                       func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_RemoveDependency                    func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_AddAppDependency                    func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId_t) SteamAPICall
	iSteamUGC_RemoveAppDependency                 func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId_t) SteamAPICall
	iSteamUGC_GetAppDependencies                  func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_DeleteItem                          func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_ShowWorkshopEULA                    func(steamUGC uintptr) bool
	iSteamUGC_GetWorkshopEULAStatus               func(steamUGC uintptr) SteamAPICall
	iSteamUGC_GetUserContentDescriptorPreferences func(steamUGC uintptr, pvecDescriptors []EUGCContentDescriptorID, cMaxEntries uint32) uint32
)

type steamUGC uintptr

func SteamUGC() ISteamUGC {
	return steamUGC(steamUGC_init())
}

func (s steamUGC) CreateQueryUserUGCRequest(accountID AccountID, listType EUserUGCList, matchingUGCType EUGCMatchingUGCType, sortOrder EUserUGCListSortOrder, creatorAppID AppId_t, consumerAppID AppId_t, page uint32) UGCQueryHandle {
	return iSteamUGC_CreateQueryUserUGCRequest(uintptr(s), accountID, listType, matchingUGCType, sortOrder, creatorAppID, consumerAppID, page)
}

func (s steamUGC) CreateQueryAllUGCRequest(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId_t, consumerAppID AppId_t, page uint32) UGCQueryHandle {
	return iSteamUGC_CreateQueryAllUGCRequestPage(uintptr(s), queryType, matchingeMatchingUGCTypeFileType, creatorAppID, consumerAppID, page)
}

func (s steamUGC) CreateQueryAllUGCRequestCursor(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId_t, consumerAppID AppId_t, cursor string) UGCQueryHandle {
	return iSteamUGC_CreateQueryAllUGCRequestCursor(uintptr(s), queryType, matchingeMatchingUGCTypeFileType, creatorAppID, consumerAppID, cursor)
}

func (s steamUGC) CreateQueryUGCDetailsRequest(publishedFileIDList []PublishedFileId) UGCQueryHandle {
	return iSteamUGC_CreateQueryUGCDetailsRequest(uintptr(s), publishedFileIDList, uint32(len(publishedFileIDList)))
}

func (s steamUGC) SendQueryUGCRequest(handle UGCQueryHandle) CallResult[SteamUGCQueryCompleted] {
	apiHandle := iSteamUGC_SendQueryUGCRequest(uintptr(s), handle)
	return CallResult[SteamUGCQueryCompleted]{Handle: apiHandle}
}

func (s steamUGC) GetQueryUGCResult(handle UGCQueryHandle, index uint32) (details SteamUGCDetails, success bool) {
	success = iSteamUGC_GetQueryUGCResult(uintptr(s), handle, index, &details)
	return details, success
}

func (s steamUGC) GetQueryUGCNumTags(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumTags(uintptr(s), handle, index)
}

func (s steamUGC) GetQueryUGCTag(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCTag(uintptr(s), handle, index, indexTag, pchValue, valueSize)
	return pchValue, result
}

func (s steamUGC) GetQueryUGCTagDisplayName(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCTagDisplayName(uintptr(s), handle, index, indexTag, pchValue, valueSize)
	return pchValue, result
}

func (s steamUGC) GetQueryUGCPreviewURL(handle UGCQueryHandle, index uint32, URLSize uint32) ([]byte, bool) {
	var pchURL []byte = make([]byte, 0, URLSize)
	result := iSteamUGC_GetQueryUGCPreviewURL(uintptr(s), handle, index, pchURL, URLSize)
	return pchURL, result
}

func (s steamUGC) GetQueryUGCMetadata(handle UGCQueryHandle, index uint32, metadataSize uint32) ([]byte, bool) {
	var pchMetadata []byte = make([]byte, 0, metadataSize)
	result := iSteamUGC_GetQueryUGCMetadata(uintptr(s), handle, index, pchMetadata, metadataSize)
	return pchMetadata, result
}

func (s steamUGC) GetQueryUGCChildren(handle UGCQueryHandle, index uint32, maxEntries uint32) ([]PublishedFileId, bool) {
	pvecPublishedFileID := make([]PublishedFileId, 0, maxEntries)
	success := iSteamUGC_GetQueryUGCChildren(uintptr(s), handle, index, pvecPublishedFileID, maxEntries)
	return pvecPublishedFileID, success
}

func (s steamUGC) GetQueryUGCStatistic(handle UGCQueryHandle, index uint32, statType EItemStatistic) (statValue uint64, success bool) {
	success = iSteamUGC_GetQueryUGCStatistic(uintptr(s), handle, index, statType, &statValue)
	return statValue, success
}

func (s steamUGC) GetQueryUGCNumAdditionalPreviews(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumAdditionalPreviews(uintptr(s), handle, index)
}

func (s steamUGC) GetQueryUGCAdditionalPreview(handle UGCQueryHandle, index uint32, previewIndex uint32, URLSize uint32, originalFileNameSize uint32) ([]byte, []byte, EItemPreviewType, bool) {
	var pchURLOrVideoID []byte = make([]byte, 0, URLSize)
	var pchOriginalFileName []byte = make([]byte, 0, originalFileNameSize)
	var pPreviewType EItemPreviewType
	result := iSteamUGC_GetQueryUGCAdditionalPreview(uintptr(s), handle, index, previewIndex, pchURLOrVideoID, URLSize, pchOriginalFileName, originalFileNameSize, &pPreviewType)
	return pchURLOrVideoID, pchOriginalFileName, pPreviewType, result
}

func (s steamUGC) GetQueryUGCNumKeyValueTags(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumKeyValueTags(uintptr(s), handle, index)
}

func (s steamUGC) GetQueryUGCKeyValueTag(handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, keySize uint32, valueSize uint32) ([]byte, []byte, bool) {
	var pchKey []byte = make([]byte, 0, keySize)
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCKeyValueTag(uintptr(s), handle, index, keyValueTagIndex, pchKey, keySize, pchValue, valueSize)
	return pchKey, pchValue, result
}

func (s steamUGC) GetQueryFirstUGCKeyValueTag(handle UGCQueryHandle, index uint32, key string, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryFirstUGCKeyValueTag(uintptr(s), handle, index, key, pchValue, valueSize)
	return pchValue, result
}

func (s steamUGC) GetNumSupportedGameVersions(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetNumSupportedGameVersions(uintptr(s), handle, index)
}

func (s steamUGC) GetSupportedGameVersionData(handle UGCQueryHandle, index uint32, versionIndex uint32, gameBranchSize uint32) ([]byte, []byte, bool) {
	var pchGameBranchMin []byte = make([]byte, 0, gameBranchSize)
	var pchGameBranchMax []byte = make([]byte, 0, gameBranchSize)
	result := iSteamUGC_GetSupportedGameVersionData(uintptr(s), handle, index, versionIndex, pchGameBranchMin, pchGameBranchMax, gameBranchSize)
	return pchGameBranchMin, pchGameBranchMax, result
}

func (s steamUGC) GetQueryUGCContentDescriptors(handle UGCQueryHandle, index uint32, maxEntries uint32) (descriptorsList []EUGCContentDescriptorID) {
	descriptorsList = make([]EUGCContentDescriptorID, 0, maxEntries)
	result := iSteamUGC_GetQueryUGCContentDescriptors(uintptr(s), handle, index, descriptorsList, maxEntries)
	return descriptorsList[:result]
}

func (s steamUGC) ReleaseQueryUGCRequest(handle UGCQueryHandle) bool {
	return iSteamUGC_ReleaseQueryUGCRequest(uintptr(s), handle)
}

func (s steamUGC) AddRequiredTag(handle UGCQueryHandle, tagName string) bool {
	return iSteamUGC_AddRequiredTag(uintptr(s), handle, tagName)
}

func (s steamUGC) AddRequiredTagGroup(handle UGCQueryHandle, tagGroups []SteamParamStringArray) bool {
	return iSteamUGC_AddRequiredTagGroup(uintptr(s), handle, tagGroups)
}

func (s steamUGC) AddExcludedTag(handle UGCQueryHandle, tagName string) bool {
	return iSteamUGC_AddExcludedTag(uintptr(s), handle, tagName)
}

func (s steamUGC) SetReturnOnlyIDs(handle UGCQueryHandle, returnOnlyIDs bool) bool {
	return iSteamUGC_SetReturnOnlyIDs(uintptr(s), handle, returnOnlyIDs)
}

func (s steamUGC) SetReturnKeyValueTags(handle UGCQueryHandle, returnKeyValueTags bool) bool {
	return iSteamUGC_SetReturnKeyValueTags(uintptr(s), handle, returnKeyValueTags)
}

func (s steamUGC) SetReturnLongDescription(handle UGCQueryHandle, returnLongDescription bool) bool {
	return iSteamUGC_SetReturnLongDescription(uintptr(s), handle, returnLongDescription)
}

func (s steamUGC) SetReturnMetadata(handle UGCQueryHandle, returnMetadata bool) bool {
	return iSteamUGC_SetReturnMetadata(uintptr(s), handle, returnMetadata)
}

func (s steamUGC) SetReturnChildren(handle UGCQueryHandle, returnChildren bool) bool {
	return iSteamUGC_SetReturnChildren(uintptr(s), handle, returnChildren)
}

func (s steamUGC) SetReturnAdditionalPreviews(handle UGCQueryHandle, returnAdditionalPreviews bool) bool {
	return iSteamUGC_SetReturnAdditionalPreviews(uintptr(s), handle, returnAdditionalPreviews)
}

func (s steamUGC) SetReturnTotalOnly(handle UGCQueryHandle, returnTotalOnly bool) bool {
	return iSteamUGC_SetReturnTotalOnly(uintptr(s), handle, returnTotalOnly)
}

func (s steamUGC) SetReturnPlaytimeStats(handle UGCQueryHandle, days uint32) bool {
	return iSteamUGC_SetReturnPlaytimeStats(uintptr(s), handle, days)
}

func (s steamUGC) SetLanguage(handle UGCQueryHandle, language string) bool {
	return iSteamUGC_SetLanguage(uintptr(s), handle, language)
}

func (s steamUGC) SetAllowCachedResponse(handle UGCQueryHandle, maxAgeSeconds uint32) bool {
	return iSteamUGC_SetAllowCachedResponse(uintptr(s), handle, maxAgeSeconds)
}

func (s steamUGC) SetAdminQuery(handle UGCUpdateHandle, adminQuery bool) bool {
	return iSteamUGC_SetAdminQuery(uintptr(s), handle, adminQuery)
}

func (s steamUGC) SetCloudFileNameFilter(handle UGCQueryHandle, matchCloudFileName string) bool {
	return iSteamUGC_SetCloudFileNameFilter(uintptr(s), handle, matchCloudFileName)
}

func (s steamUGC) SetMatchAnyTag(handle UGCQueryHandle, matchAnyTag bool) bool {
	return iSteamUGC_SetMatchAnyTag(uintptr(s), handle, matchAnyTag)
}

func (s steamUGC) SetSearchText(handle UGCQueryHandle, searchText string) bool {
	return iSteamUGC_SetSearchText(uintptr(s), handle, searchText)
}

func (s steamUGC) SetRankedByTrendDays(handle UGCQueryHandle, days uint32) bool {
	return iSteamUGC_SetRankedByTrendDays(uintptr(s), handle, days)
}

func (s steamUGC) SetTimeCreatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool {
	return iSteamUGC_SetTimeCreatedDateRange(uintptr(s), handle, startTime, endTime)
}

func (s steamUGC) SetTimeUpdatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool {
	return iSteamUGC_SetTimeUpdatedDateRange(uintptr(s), handle, startTime, endTime)
}

func (s steamUGC) AddRequiredKeyValueTag(handle UGCQueryHandle, key string, value string) bool {
	return iSteamUGC_AddRequiredKeyValueTag(uintptr(s), handle, key, value)
}

func (s steamUGC) CreateItem(consumerAppId AppId_t, fileType EWorkshopFileType) CallResult[CreateItemResult] {
	handle := iSteamUGC_CreateItem(uintptr(s), consumerAppId, fileType)
	return CallResult[CreateItemResult]{Handle: handle}
}

func (s steamUGC) StartItemUpdate(consumerAppId AppId_t, publishedFileID PublishedFileId) UGCUpdateHandle {
	return iSteamUGC_StartItemUpdate(uintptr(s), consumerAppId, publishedFileID)
}

func (s steamUGC) SetItemTitle(handle UGCUpdateHandle, title string) bool {
	return iSteamUGC_SetItemTitle(uintptr(s), handle, title)
}

func (s steamUGC) SetItemDescription(handle UGCUpdateHandle, description string) bool {
	return iSteamUGC_SetItemDescription(uintptr(s), handle, description)
}

func (s steamUGC) SetItemUpdateLanguage(handle UGCUpdateHandle, language string) bool {
	return iSteamUGC_SetItemUpdateLanguage(uintptr(s), handle, language)
}

func (s steamUGC) SetItemMetadata(handle UGCUpdateHandle, metaData string) bool {
	return iSteamUGC_SetItemMetadata(uintptr(s), handle, metaData)
}

func (s steamUGC) SetItemVisibility(handle UGCUpdateHandle, visibility ERemoteStoragePublishedFileVisibility) bool {
	return iSteamUGC_SetItemVisibility(uintptr(s), handle, visibility)
}

func (s steamUGC) SetItemTags(updateHandle UGCUpdateHandle, tags []SteamParamStringArray, allowAdminTags bool) bool {
	return iSteamUGC_SetItemTags(uintptr(s), updateHandle, tags, allowAdminTags)
}

func (s steamUGC) SetItemContent(handle UGCUpdateHandle, contentFolder string) bool {
	return iSteamUGC_SetItemContent(uintptr(s), handle, contentFolder)
}

func (s steamUGC) SetItemPreview(handle UGCUpdateHandle, previewFile string) bool {
	return iSteamUGC_SetItemPreview(uintptr(s), handle, previewFile)
}

func (s steamUGC) SetAllowLegacyUpload(handle UGCUpdateHandle, allowLegacyUpload bool) bool {
	return iSteamUGC_SetAllowLegacyUpload(uintptr(s), handle, allowLegacyUpload)
}

func (s steamUGC) RemoveAllItemKeyValueTags(handle UGCUpdateHandle) bool {
	return iSteamUGC_RemoveAllItemKeyValueTags(uintptr(s), handle)
}

func (s steamUGC) RemoveItemKeyValueTags(handle UGCUpdateHandle, key string) bool {
	return iSteamUGC_RemoveItemKeyValueTags(uintptr(s), handle, key)
}

func (s steamUGC) AddItemKeyValueTag(handle UGCUpdateHandle, key string, value string) bool {
	return iSteamUGC_AddItemKeyValueTag(uintptr(s), handle, key, value)
}

func (s steamUGC) AddItemPreviewFile(handle UGCUpdateHandle, previewFile string, Type EItemPreviewType) bool {
	return iSteamUGC_AddItemPreviewFile(uintptr(s), handle, previewFile, Type)
}

func (s steamUGC) AddItemPreviewVideo(handle UGCUpdateHandle, videoID string) bool {
	return iSteamUGC_AddItemPreviewVideo(uintptr(s), handle, videoID)
}

func (s steamUGC) UpdateItemPreviewFile(handle UGCUpdateHandle, index uint32, previewFile string) bool {
	return iSteamUGC_UpdateItemPreviewFile(uintptr(s), handle, index, previewFile)
}

func (s steamUGC) UpdateItemPreviewVideo(handle UGCUpdateHandle, index uint32, videoID string) bool {
	return iSteamUGC_UpdateItemPreviewVideo(uintptr(s), handle, index, videoID)
}

func (s steamUGC) RemoveItemPreview(handle UGCUpdateHandle, index uint32) bool {
	return iSteamUGC_RemoveItemPreview(uintptr(s), handle, index)
}

func (s steamUGC) AddContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool {
	return iSteamUGC_AddContentDescriptor(uintptr(s), handle, descID)
}

func (s steamUGC) RemoveContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool {
	return iSteamUGC_RemoveContentDescriptor(uintptr(s), handle, descID)
}

func (s steamUGC) SetRequiredGameVersions(handle UGCUpdateHandle, gameBranchMin string, gameBranchMax string) bool {
	return iSteamUGC_SetRequiredGameVersions(uintptr(s), handle, gameBranchMin, gameBranchMax)
}

func (s steamUGC) SubmitItemUpdate(handle UGCUpdateHandle, changeNote string) CallResult[SubmitItemUpdateResult] {
	apiHandle := iSteamUGC_SubmitItemUpdate(uintptr(s), handle, changeNote)
	return CallResult[SubmitItemUpdateResult]{Handle: apiHandle}
}

func (s steamUGC) GetItemUpdateProgress(handle UGCUpdateHandle) (bytesProcessed uint64, bytesTotal uint64, status EItemUpdateStatus) {
	status = iSteamUGC_GetItemUpdateProgress(uintptr(s), handle, &bytesProcessed, &bytesTotal)
	return bytesProcessed, bytesTotal, status
}

func (s steamUGC) SetUserItemVote(publishedFileID PublishedFileId, voteUp bool) CallResult[SetUserItemVoteResult] {
	handle := iSteamUGC_SetUserItemVote(uintptr(s), publishedFileID, voteUp)
	return CallResult[SetUserItemVoteResult]{Handle: handle}
}

func (s steamUGC) GetUserItemVote(publishedFileID PublishedFileId) CallResult[GetUserItemVoteResult] {
	handle := iSteamUGC_GetUserItemVote(uintptr(s), publishedFileID)
	return CallResult[GetUserItemVoteResult]{Handle: handle}
}

func (s steamUGC) AddItemToFavorites(appId AppId_t, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged] {
	handle := iSteamUGC_AddItemToFavorites(uintptr(s), appId, publishedFileID)
	return CallResult[UserFavoriteItemsListChanged]{Handle: handle}
}

func (s steamUGC) RemoveItemFromFavorites(appId AppId_t, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged] {
	handle := iSteamUGC_RemoveItemFromFavorites(uintptr(s), appId, publishedFileID)
	return CallResult[UserFavoriteItemsListChanged]{Handle: handle}
}

func (s steamUGC) SubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageSubscribePublishedFileResult] {
	handle := iSteamUGC_SubscribeItem(uintptr(s), publishedFileID)
	return CallResult[RemoteStorageSubscribePublishedFileResult]{Handle: handle}
}

func (s steamUGC) UnsubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageUnsubscribePublishedFileResult] {
	handle := iSteamUGC_UnsubscribeItem(uintptr(s), publishedFileID)
	return CallResult[RemoteStorageUnsubscribePublishedFileResult]{Handle: handle}
}

func (s steamUGC) GetNumSubscribedItems() uint32 {
	return iSteamUGC_GetNumSubscribedItems(uintptr(s))
}

func (s steamUGC) GetSubscribedItems(maxEntries uint32) []PublishedFileId {
	pvecPublishedFileID := make([]PublishedFileId, 0, maxEntries)
	total := iSteamUGC_GetSubscribedItems(uintptr(s), pvecPublishedFileID, maxEntries)
	return pvecPublishedFileID[:total]
}

func (s steamUGC) GetItemState(publishedFileID PublishedFileId) uint32 {
	return iSteamUGC_GetItemState(uintptr(s), publishedFileID)
}

func (s steamUGC) GetItemInstallInfo(publishedFileID PublishedFileId, folderSize uint32) (uint64, uint32, []byte, bool) {
	var pchFolder []byte = make([]byte, 0, folderSize)
	var punSizeOnDisk uint64
	var punTimeStamp uint32
	result := iSteamUGC_GetItemInstallInfo(uintptr(s), publishedFileID, &punSizeOnDisk, pchFolder, folderSize, &punTimeStamp)
	return punSizeOnDisk, punTimeStamp, pchFolder, result
}

func (s steamUGC) GetItemDownloadInfo(publishedFileID PublishedFileId) (bytesDownloaded uint64, bytesTotal uint64, success bool) {
	success = iSteamUGC_GetItemDownloadInfo(uintptr(s), publishedFileID, &bytesDownloaded, &bytesTotal)
	return bytesDownloaded, bytesDownloaded, success
}

func (s steamUGC) DownloadItem(publishedFileID PublishedFileId, highPriority bool) bool {
	return iSteamUGC_DownloadItem(uintptr(s), publishedFileID, highPriority)
}

func (s steamUGC) BInitWorkshopForGameServer(workshopDepotID DepotId, folder string) bool {
	return iSteamUGC_BInitWorkshopForGameServer(uintptr(s), workshopDepotID, folder)
}

func (s steamUGC) SuspendDownloads(suspend bool) {
	iSteamUGC_SuspendDownloads(uintptr(s), suspend)
}

func (s steamUGC) StartPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StartPlaytimeTrackingResult] {
	handle := iSteamUGC_StartPlaytimeTracking(uintptr(s), publishedFileIDList, uint32(len(publishedFileIDList)))
	return CallResult[StartPlaytimeTrackingResult]{Handle: handle}
}

func (s steamUGC) StopPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StopPlaytimeTrackingResult] {
	handle := iSteamUGC_StopPlaytimeTracking(uintptr(s), publishedFileIDList, uint32(len(publishedFileIDList)))
	return CallResult[StopPlaytimeTrackingResult]{Handle: handle}
}

func (s steamUGC) StopPlaytimeTrackingForAllItems() CallResult[StopPlaytimeTrackingResult] {
	handle := iSteamUGC_StopPlaytimeTrackingForAllItems(uintptr(s))
	return CallResult[StopPlaytimeTrackingResult]{Handle: handle}
}

func (s steamUGC) AddDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[AddUGCDependencyResult] {
	handle := iSteamUGC_AddDependency(uintptr(s), parentPublishedFileID, childPublishedFileID)
	return CallResult[AddUGCDependencyResult]{Handle: handle}
}

func (s steamUGC) RemoveDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[RemoveUGCDependencyResult] {
	handle := iSteamUGC_RemoveDependency(uintptr(s), parentPublishedFileID, childPublishedFileID)
	return CallResult[RemoveUGCDependencyResult]{Handle: handle}
}

func (s steamUGC) AddAppDependency(publishedFileID PublishedFileId, appID AppId_t) CallResult[AddAppDependencyResult] {
	handle := iSteamUGC_AddAppDependency(uintptr(s), publishedFileID, appID)
	return CallResult[AddAppDependencyResult]{Handle: handle}
}

func (s steamUGC) RemoveAppDependency(publishedFileID PublishedFileId, appID AppId_t) CallResult[RemoveAppDependencyResult] {
	handle := iSteamUGC_RemoveAppDependency(uintptr(s), publishedFileID, appID)
	return CallResult[RemoveAppDependencyResult]{Handle: handle}
}

func (s steamUGC) GetAppDependencies(publishedFileID PublishedFileId) CallResult[GetAppDependenciesResult] {
	handle := iSteamUGC_GetAppDependencies(uintptr(s), publishedFileID)
	return CallResult[GetAppDependenciesResult]{Handle: handle}
}

func (s steamUGC) DeleteItem(publishedFileID PublishedFileId) CallResult[DeleteItemResult] {
	handle := iSteamUGC_DeleteItem(uintptr(s), publishedFileID)
	return CallResult[DeleteItemResult]{Handle: handle}
}

func (s steamUGC) ShowWorkshopEULA() bool {
	return iSteamUGC_ShowWorkshopEULA(uintptr(s))
}

func (s steamUGC) GetWorkshopEULAStatus() CallResult[WorkshopEULAStatus] {
	handle := iSteamUGC_GetWorkshopEULAStatus(uintptr(s))
	return CallResult[WorkshopEULAStatus]{Handle: handle}
}

func (s steamUGC) GetUserContentDescriptorPreferences(maxEntries uint32) []EUGCContentDescriptorID {
	var pvecDescriptors []EUGCContentDescriptorID = make([]EUGCContentDescriptorID, 0, maxEntries)
	total := iSteamUGC_GetUserContentDescriptorPreferences(uintptr(s), pvecDescriptors, maxEntries)
	return pvecDescriptors[:total]
}

// Steam User
const (
	HAuthTicketInvalid HAuthTicket = 0
)

type EDurationControlNotification int32

const (
	EDurationControl_Notification_None          EDurationControlNotification = 0
	EDurationControl_Notification_1Hour         EDurationControlNotification = 1
	EDurationControl_Notification_3Hours        EDurationControlNotification = 2
	EDurationControl_Notification_HalfProgress  EDurationControlNotification = 3
	EDurationControl_Notification_NoProgress    EDurationControlNotification = 4
	EDurationControl_Notification_ExitSoon3h    EDurationControlNotification = 5
	EDurationControl_Notification_ExitSoon5h    EDurationControlNotification = 6
	EDurationControl_Notification_ExitSoonNight EDurationControlNotification = 7
)

type EDurationControlProgress int32

const (
	EDurationControl_ProgressFull  EDurationControlProgress = 0
	EDurationControl_ProgressHalf  EDurationControlProgress = 1
	EDurationControl_ProgressNone  EDurationControlProgress = 2
	EDurationControl_ExitSoon3h    EDurationControlProgress = 3
	EDurationControl_ExitSoon5h    EDurationControlProgress = 4
	EDurationControl_ExitSoonNight EDurationControlProgress = 5
)

type EMarketNotAllowedReasonFlags int32

const (
	EMarketNotAllowedReason_None                             EMarketNotAllowedReasonFlags = 0
	EMarketNotAllowedReason_TemporaryFailure                 EMarketNotAllowedReasonFlags = 1
	EMarketNotAllowedReason_AccountDisabled                  EMarketNotAllowedReasonFlags = 2
	EMarketNotAllowedReason_AccountLockedDown                EMarketNotAllowedReasonFlags = 4
	EMarketNotAllowedReason_AccountLimited                   EMarketNotAllowedReasonFlags = 8
	EMarketNotAllowedReason_TradeBanned                      EMarketNotAllowedReasonFlags = 16
	EMarketNotAllowedReason_AccountNotTrusted                EMarketNotAllowedReasonFlags = 32
	EMarketNotAllowedReason_SteamGuardNotEnabled             EMarketNotAllowedReasonFlags = 64
	EMarketNotAllowedReason_SteamGuardOnlyRecentlyEnabled    EMarketNotAllowedReasonFlags = 128
	EMarketNotAllowedReason_RecentPasswordReset              EMarketNotAllowedReasonFlags = 256
	EMarketNotAllowedReason_NewPaymentMethod                 EMarketNotAllowedReasonFlags = 512
	EMarketNotAllowedReason_InvalidCookie                    EMarketNotAllowedReasonFlags = 1024
	EMarketNotAllowedReason_UsingNewDevice                   EMarketNotAllowedReasonFlags = 2048
	EMarketNotAllowedReason_RecentSelfRefund                 EMarketNotAllowedReasonFlags = 4096
	EMarketNotAllowedReason_NewPaymentMethodCannotBeVerified EMarketNotAllowedReasonFlags = 8192
	EMarketNotAllowedReason_NoRecentPurchases                EMarketNotAllowedReasonFlags = 16384
	EMarketNotAllowedReason_AcceptedWalletGift               EMarketNotAllowedReasonFlags = 32768
)

type EAuthSessionResponse int32

const (
	EAuthSessionResponse_OK                               EAuthSessionResponse = 0
	EAuthSessionResponse_UserNotConnectedToSteam          EAuthSessionResponse = 1
	EAuthSessionResponse_NoLicenseOrExpired               EAuthSessionResponse = 2
	EAuthSessionResponse_VACBanned                        EAuthSessionResponse = 3
	EAuthSessionResponse_LoggedInElseWhere                EAuthSessionResponse = 4
	EAuthSessionResponse_VACCheckTimedOut                 EAuthSessionResponse = 5
	EAuthSessionResponse_AuthTicketCanceled               EAuthSessionResponse = 6
	EAuthSessionResponse_AuthTicketInvalidAlreadyUsed     EAuthSessionResponse = 7
	EAuthSessionResponse_AuthTicketInvalid                EAuthSessionResponse = 8
	EAuthSessionResponse_PublisherIssuedBan               EAuthSessionResponse = 9
	EAuthSessionResponse_AuthTicketNetworkIdentityFailure EAuthSessionResponse = 10
)

type HAuthTicket uint

type EVoiceResult int32

const (
	EVoiceResult_OK                   EVoiceResult = 0
	EVoiceResult_NotInitialized       EVoiceResult = 1
	EVoiceResult_NotRecording         EVoiceResult = 2
	EVoiceResult_NoData               EVoiceResult = 3
	EVoiceResult_BufferTooSmall       EVoiceResult = 4
	EVoiceResult_DataCorrupted        EVoiceResult = 5
	EVoiceResult_Restricted           EVoiceResult = 6
	EVoiceResult_UnsupportedCodec     EVoiceResult = 7
	EVoiceResult_ReceiverOutOfDate    EVoiceResult = 8
	EVoiceResult_ReceiverDidNotAnswer EVoiceResult = 9
)

type EDurationControlOnlineState int32

const (
	EDurationControlOnlineState_Invalid       EDurationControlOnlineState = 0
	EDurationControlOnlineState_Offline       EDurationControlOnlineState = 1
	EDurationControlOnlineState_Online        EDurationControlOnlineState = 2
	EDurationControlOnlineState_OnlineHighPri EDurationControlOnlineState = 3
)

type EUserHasLicenseForAppResult int32

const (
	EUserHasLicenseResult_HasLicense         EUserHasLicenseForAppResult = 0
	EUserHasLicenseResult_DoesNotHaveLicense EUserHasLicenseForAppResult = 1
	EUserHasLicenseResult_NoAuth             EUserHasLicenseForAppResult = 2
)

type EBeginAuthSessionResult int32

const (
	EBeginAuthSessionResult_OK               EBeginAuthSessionResult = 0
	EBeginAuthSessionResult_InvalidTicket    EBeginAuthSessionResult = 1
	EBeginAuthSessionResult_DuplicateRequest EBeginAuthSessionResult = 2
	EBeginAuthSessionResult_InvalidVersion   EBeginAuthSessionResult = 3
	EBeginAuthSessionResult_GameMismatch     EBeginAuthSessionResult = 4
	EBeginAuthSessionResult_ExpiredTicket    EBeginAuthSessionResult = 5
)

var (
	steamUser_init                            func() uintptr
	iSteamUser_GetHSteamUser                  func(steamUser uintptr) HSteamUser
	iSteamUser_BLoggedOn                      func(steamUser uintptr) bool
	iSteamUser_GetSteamID                     func(steamUser uintptr) CSteamID
	iSteamUser_TrackAppUsageEvent             func(steamUser uintptr, gameID Uint64SteamID, eAppUsageEvent int32, pchExtraInfo string)
	iSteamUser_GetUserDataFolder              func(steamUser uintptr, pchBuffer []byte, cubBuffer int32) bool
	iSteamUser_StartVoiceRecording            func(steamUser uintptr)
	iSteamUser_StopVoiceRecording             func(steamUser uintptr)
	iSteamUser_GetAvailableVoice              func(steamUser uintptr, pcbCompressed *uint32, pcbUncompressedDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult
	iSteamUser_GetVoice                       func(steamUser uintptr, bWantCompressed bool, pDestBuffer []byte, cbDestBufferSize uint32, nBytesWritten *uint32, bWantUncompressedDeprecated bool, pUncompressedDestBufferDeprecated []byte, cbUncompressedDestBufferSizeDeprecated uint32, nUncompressBytesWrittenDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult
	iSteamUser_DecompressVoice                func(steamUser uintptr, pCompressed []byte, cbCompressed uint32, pDestBuffer []byte, cbDestBufferSize uint32, nBytesWritten *uint32, nDesiredSampleRate uint32) EVoiceResult
	iSteamUser_GetVoiceOptimalSampleRate      func(steamUser uintptr) uint32
	iSteamUser_GetAuthSessionTicket           func(steamUser uintptr, pTicket []byte, cbMaxTicket int32, pcbTicket *uint32, pSteamNetworkingIdentity *SteamNetworkingIdentity) HAuthTicket
	iSteamUser_GetAuthTicketForWebApi         func(steamUser uintptr, pchIdentity string) HAuthTicket
	iSteamUser_BeginAuthSession               func(steamUser uintptr, pAuthTicket []byte, cbAuthTicket int32, steamID Uint64SteamID) EBeginAuthSessionResult
	iSteamUser_EndAuthSession                 func(steamUser uintptr, steamID Uint64SteamID)
	iSteamUser_CancelAuthTicket               func(steamUser uintptr, hAuthTicket HAuthTicket)
	iSteamUser_UserHasLicenseForApp           func(steamUser uintptr, steamID Uint64SteamID, appID AppId_t) EUserHasLicenseForAppResult
	iSteamUser_BIsBehindNAT                   func(steamUser uintptr) bool
	iSteamUser_AdvertiseGame                  func(steamUser uintptr, steamIDGameServer Uint64SteamID, unIPServer uint32, usPortServer uint16)
	iSteamUser_RequestEncryptedAppTicket      func(steamUser uintptr, pDataToInclude []byte, cbDataToInclude int32) SteamAPICall
	iSteamUser_GetEncryptedAppTicket          func(steamUser uintptr, pTicket []byte, cbMaxTicket int32, pcbTicket *uint32) bool
	iSteamUser_GetGameBadgeLevel              func(steamUser uintptr, nSeries int32, bFoil bool) int32
	iSteamUser_GetPlayerSteamLevel            func(steamUser uintptr) int32
	iSteamUser_RequestStoreAuthURL            func(steamUser uintptr, pchRedirectURL string) SteamAPICall
	iSteamUser_BIsPhoneVerified               func(steamUser uintptr) bool
	iSteamUser_BIsTwoFactorEnabled            func(steamUser uintptr) bool
	iSteamUser_BIsPhoneIdentifying            func(steamUser uintptr) bool
	iSteamUser_BIsPhoneRequiringVerification  func(steamUser uintptr) bool
	iSteamUser_GetMarketEligibility           func(steamUser uintptr) SteamAPICall
	iSteamUser_GetDurationControl             func(steamUser uintptr) SteamAPICall
	iSteamUser_BSetDurationControlOnlineState func(steamUser uintptr, eNewState EDurationControlOnlineState) bool
)

const (
	flatAPI_SteamUser                                 = "SteamAPI_SteamUser_v023"
	flatAPI_ISteamUser_GetHSteamUser                  = "SteamAPI_ISteamUser_GetHSteamUser"
	flatAPI_ISteamUser_BLoggedOn                      = "SteamAPI_ISteamUser_BLoggedOn"
	flatAPI_ISteamUser_GetSteamID                     = "SteamAPI_ISteamUser_GetSteamID"
	flatAPI_ISteamUser_TrackAppUsageEvent             = "SteamAPI_ISteamUser_TrackAppUsageEvent"
	flatAPI_ISteamUser_GetUserDataFolder              = "SteamAPI_ISteamUser_GetUserDataFolder"
	flatAPI_ISteamUser_StartVoiceRecording            = "SteamAPI_ISteamUser_StartVoiceRecording"
	flatAPI_ISteamUser_StopVoiceRecording             = "SteamAPI_ISteamUser_StopVoiceRecording"
	flatAPI_ISteamUser_GetAvailableVoice              = "SteamAPI_ISteamUser_GetAvailableVoice"
	flatAPI_ISteamUser_GetVoice                       = "SteamAPI_ISteamUser_GetVoice"
	flatAPI_ISteamUser_DecompressVoice                = "SteamAPI_ISteamUser_DecompressVoice"
	flatAPI_ISteamUser_GetVoiceOptimalSampleRate      = "SteamAPI_ISteamUser_GetVoiceOptimalSampleRate"
	flatAPI_ISteamUser_GetAuthSessionTicket           = "SteamAPI_ISteamUser_GetAuthSessionTicket"
	flatAPI_ISteamUser_GetAuthTicketForWebApi         = "SteamAPI_ISteamUser_GetAuthTicketForWebApi"
	flatAPI_ISteamUser_BeginAuthSession               = "SteamAPI_ISteamUser_BeginAuthSession"
	flatAPI_ISteamUser_EndAuthSession                 = "SteamAPI_ISteamUser_EndAuthSession"
	flatAPI_ISteamUser_CancelAuthTicket               = "SteamAPI_ISteamUser_CancelAuthTicket"
	flatAPI_ISteamUser_UserHasLicenseForApp           = "SteamAPI_ISteamUser_UserHasLicenseForApp"
	flatAPI_ISteamUser_BIsBehindNAT                   = "SteamAPI_ISteamUser_BIsBehindNAT"
	flatAPI_ISteamUser_AdvertiseGame                  = "SteamAPI_ISteamUser_AdvertiseGame"
	flatAPI_ISteamUser_RequestEncryptedAppTicket      = "SteamAPI_ISteamUser_RequestEncryptedAppTicket"
	flatAPI_ISteamUser_GetEncryptedAppTicket          = "SteamAPI_ISteamUser_GetEncryptedAppTicket"
	flatAPI_ISteamUser_GetGameBadgeLevel              = "SteamAPI_ISteamUser_GetGameBadgeLevel"
	flatAPI_ISteamUser_GetPlayerSteamLevel            = "SteamAPI_ISteamUser_GetPlayerSteamLevel"
	flatAPI_ISteamUser_RequestStoreAuthURL            = "SteamAPI_ISteamUser_RequestStoreAuthURL"
	flatAPI_ISteamUser_BIsPhoneVerified               = "SteamAPI_ISteamUser_BIsPhoneVerified"
	flatAPI_ISteamUser_BIsTwoFactorEnabled            = "SteamAPI_ISteamUser_BIsTwoFactorEnabled"
	flatAPI_ISteamUser_BIsPhoneIdentifying            = "SteamAPI_ISteamUser_BIsPhoneIdentifying"
	flatAPI_ISteamUser_BIsPhoneRequiringVerification  = "SteamAPI_ISteamUser_BIsPhoneRequiringVerification"
	flatAPI_ISteamUser_GetMarketEligibility           = "SteamAPI_ISteamUser_GetMarketEligibility"
	flatAPI_ISteamUser_GetDurationControl             = "SteamAPI_ISteamUser_GetDurationControl"
	flatAPI_ISteamUser_BSetDurationControlOnlineState = "SteamAPI_ISteamUser_BSetDurationControlOnlineState"
)

type steamUser uintptr

func SteamUser() ISteamUser {
	return steamUser(steamUser_init())
}

func (s steamUser) GetHSteamUser() HSteamUser {
	return iSteamUser_GetHSteamUser(uintptr(s))
}

func (s steamUser) BLoggedOn() bool {
	return iSteamUser_BLoggedOn(uintptr(s))
}

func (s steamUser) GetSteamID() CSteamID {
	return iSteamUser_GetSteamID(uintptr(s))
}

func (s steamUser) TrackAppUsageEvent(gameID Uint64SteamID, appUsageEvent int32, extraInfo string) {
	iSteamUser_TrackAppUsageEvent(uintptr(s), gameID, appUsageEvent, extraInfo)
}

func (s steamUser) GetUserDataFolder(bufferSize int32) (folder string, success bool) {
	var pchBuffer []byte = make([]byte, 0, bufferSize)
	success = iSteamUser_GetUserDataFolder(uintptr(s), pchBuffer, bufferSize)
	return string(pchBuffer), success
}

func (s steamUser) StartVoiceRecording() {
	iSteamUser_StartVoiceRecording(uintptr(s))
}

func (s steamUser) StopVoiceRecording() {
	iSteamUser_StopVoiceRecording(uintptr(s))
}

func (s steamUser) GetAvailableVoice() (compressedDataSize uint32, result EVoiceResult) {
	var pcbUncompressedDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated uint32
	result = iSteamUser_GetAvailableVoice(uintptr(s), &compressedDataSize, &pcbUncompressedDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated)
	return compressedDataSize, result
}

func (s steamUser) GetVoice(destBufferSize uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult) {
	var bWantCompressed bool = true
	destBuffer = make([]byte, 0, destBufferSize)
	var bWantUncompressedDeprecated bool
	var pUncompressedDestBufferDeprecated []byte
	var cbUncompressedDestBufferSizeDeprecated, nUncompressBytesWrittenDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated uint32
	result = iSteamUser_GetVoice(uintptr(s), bWantCompressed, destBuffer, destBufferSize, &bytesWritten, bWantUncompressedDeprecated, pUncompressedDestBufferDeprecated, cbUncompressedDestBufferSizeDeprecated, &nUncompressBytesWrittenDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated)
	return destBuffer, bytesWritten, result
}

func (s steamUser) DecompressVoice(compressedData []byte, destBufferSize uint32, desiredSampleRate uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult) {
	destBuffer = make([]byte, 0, destBufferSize)
	result = iSteamUser_DecompressVoice(uintptr(s), compressedData, uint32(len(compressedData)), destBuffer, destBufferSize, &bytesWritten, desiredSampleRate)
	return destBuffer, bytesWritten, result
}

func (s steamUser) GetVoiceOptimalSampleRate() uint32 {
	return iSteamUser_GetVoiceOptimalSampleRate(uintptr(s))
}

func (s steamUser) GetAuthSessionTicket(maxTicket int32, optionalIdetity *SteamNetworkingIdentity) (ticketData []byte, ticketHandle HAuthTicket) {
	var pcbTicket uint32
	ticketData = make([]byte, 0, maxTicket)
	ticketHandle = iSteamUser_GetAuthSessionTicket(uintptr(s), ticketData, maxTicket, &pcbTicket, optionalIdetity)
	return ticketData[:pcbTicket], ticketHandle
}

func (s steamUser) GetAuthTicketForWebApi(identity string) HAuthTicket {
	return iSteamUser_GetAuthTicketForWebApi(uintptr(s), identity)
}

func (s steamUser) BeginAuthSession(authTicket []byte, steamID Uint64SteamID) EBeginAuthSessionResult {
	return iSteamUser_BeginAuthSession(uintptr(s), authTicket, int32(len(authTicket)), steamID)
}

func (s steamUser) EndAuthSession(steamID Uint64SteamID) {
	iSteamUser_EndAuthSession(uintptr(s), steamID)
}

func (s steamUser) CancelAuthTicket(hAuthTicket HAuthTicket) {
	iSteamUser_CancelAuthTicket(uintptr(s), hAuthTicket)
}

func (s steamUser) UserHasLicenseForApp(steamID Uint64SteamID, appID AppId_t) EUserHasLicenseForAppResult {
	return iSteamUser_UserHasLicenseForApp(uintptr(s), steamID, appID)
}

func (s steamUser) BIsBehindNAT() bool {
	return iSteamUser_BIsBehindNAT(uintptr(s))
}

func (s steamUser) AdvertiseGame(gameServerSteamID Uint64SteamID, serverIP uint32, serverPort uint16) {
	iSteamUser_AdvertiseGame(uintptr(s), gameServerSteamID, serverIP, serverPort)
}

func (s steamUser) RequestEncryptedAppTicket(dataToInclude []byte) CallResult[EncryptedAppTicketResponse] {
	handle := iSteamUser_RequestEncryptedAppTicket(uintptr(s), dataToInclude, int32(len(dataToInclude)))
	return CallResult[EncryptedAppTicketResponse]{Handle: handle}
}

// Retrieves a finished ticket.
// Performs the API call once to get the proper size and then again to get the
// actual data.
func (s steamUser) GetEncryptedAppTicket() (ticketData []byte, ticketAvailable bool) {
	var ticketSize uint32
	available := iSteamUser_GetEncryptedAppTicket(uintptr(s), nil, 0, &ticketSize)
	if available {
		ticketData = make([]byte, 0, ticketSize)
		iSteamUser_GetEncryptedAppTicket(uintptr(s), ticketData, 0, &ticketSize)
	}
	return ticketData, available
}

func (s steamUser) GetGameBadgeLevel(series int32, foil bool) int32 {
	return iSteamUser_GetGameBadgeLevel(uintptr(s), series, foil)
}

func (s steamUser) GetPlayerSteamLevel() int32 {
	return iSteamUser_GetPlayerSteamLevel(uintptr(s))
}

func (s steamUser) RequestStoreAuthURL(redirectURL string) CallResult[StoreAuthURLResponse] {
	handle := iSteamUser_RequestStoreAuthURL(uintptr(s), redirectURL)
	return CallResult[StoreAuthURLResponse]{Handle: handle}
}

func (s steamUser) BIsPhoneVerified() bool {
	return iSteamUser_BIsPhoneVerified(uintptr(s))
}

func (s steamUser) BIsTwoFactorEnabled() bool {
	return iSteamUser_BIsTwoFactorEnabled(uintptr(s))
}

func (s steamUser) BIsPhoneIdentifying() bool {
	return iSteamUser_BIsPhoneIdentifying(uintptr(s))
}

func (s steamUser) BIsPhoneRequiringVerification() bool {
	return iSteamUser_BIsPhoneRequiringVerification(uintptr(s))
}

func (s steamUser) GetMarketEligibility() CallResult[MarketEligibilityResponse] {
	handle := iSteamUser_GetMarketEligibility(uintptr(s))
	return CallResult[MarketEligibilityResponse]{Handle: handle}
}

func (s steamUser) GetDurationControl() CallResult[DurationControl] {
	handle := iSteamUser_GetDurationControl(uintptr(s))
	return CallResult[DurationControl]{Handle: handle}
}

func (s steamUser) BSetDurationControlOnlineState(newState EDurationControlOnlineState) bool {
	return iSteamUser_BSetDurationControlOnlineState(uintptr(s), newState)
}

// Steam User Stats
type LeaderboardEntry struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	UGC         UGCHandle
}
type ELeaderboardDataRequest int32

const (
	ELeaderboardDataRequest_Global           ELeaderboardDataRequest = 0
	ELeaderboardDataRequest_GlobalAroundUser ELeaderboardDataRequest = 1
	ELeaderboardDataRequest_Friends          ELeaderboardDataRequest = 2
	ELeaderboardDataRequest_Users            ELeaderboardDataRequest = 3
)

type ELeaderboardSortMethod int32

const (
	ELeaderboardSortMethod_None       ELeaderboardSortMethod = 0
	ELeaderboardSortMethod_Ascending  ELeaderboardSortMethod = 1
	ELeaderboardSortMethod_Descending ELeaderboardSortMethod = 2
)

type ELeaderboardDisplayType int32

const (
	ELeaderboardDisplayType_None             ELeaderboardDisplayType = 0
	ELeaderboardDisplayType_Numeric          ELeaderboardDisplayType = 1
	ELeaderboardDisplayType_TimeSeconds      ELeaderboardDisplayType = 2
	ELeaderboardDisplayType_TimeMilliSeconds ELeaderboardDisplayType = 3
)

type ELeaderboardUploadScoreMethod int32

const (
	ELeaderboardUploadScoreMethod_None        ELeaderboardUploadScoreMethod = 0
	ELeaderboardUploadScoreMethod_KeepBest    ELeaderboardUploadScoreMethod = 1
	ELeaderboardUploadScoreMethod_ForceUpdate ELeaderboardUploadScoreMethod = 2
)

type SteamLeaderboard uint64
type SteamLeaderboardEntries uint64

const (
	flatAPI_SteamUserStats                                      = "SteamAPI_SteamUserStats_v013"
	flatAPI_ISteamUserStats_GetStatInt32                        = "SteamAPI_ISteamUserStats_GetStatInt32"
	flatAPI_ISteamUserStats_GetStatFloat                        = "SteamAPI_ISteamUserStats_GetStatFloat"
	flatAPI_ISteamUserStats_SetStatInt32                        = "SteamAPI_ISteamUserStats_SetStatInt32"
	flatAPI_ISteamUserStats_SetStatFloat                        = "SteamAPI_ISteamUserStats_SetStatFloat"
	flatAPI_ISteamUserStats_UpdateAvgRateStat                   = "SteamAPI_ISteamUserStats_UpdateAvgRateStat"
	flatAPI_ISteamUserStats_GetAchievement                      = "SteamAPI_ISteamUserStats_GetAchievement"
	flatAPI_ISteamUserStats_SetAchievement                      = "SteamAPI_ISteamUserStats_SetAchievement"
	flatAPI_ISteamUserStats_ClearAchievement                    = "SteamAPI_ISteamUserStats_ClearAchievement"
	flatAPI_ISteamUserStats_GetAchievementAndUnlockTime         = "SteamAPI_ISteamUserStats_GetAchievementAndUnlockTime"
	flatAPI_ISteamUserStats_StoreStats                          = "SteamAPI_ISteamUserStats_StoreStats"
	flatAPI_ISteamUserStats_GetAchievementIcon                  = "SteamAPI_ISteamUserStats_GetAchievementIcon"
	flatAPI_ISteamUserStats_GetAchievementDisplayAttribute      = "SteamAPI_ISteamUserStats_GetAchievementDisplayAttribute"
	flatAPI_ISteamUserStats_IndicateAchievementProgress         = "SteamAPI_ISteamUserStats_IndicateAchievementProgress"
	flatAPI_ISteamUserStats_GetNumAchievements                  = "SteamAPI_ISteamUserStats_GetNumAchievements"
	flatAPI_ISteamUserStats_GetAchievementName                  = "SteamAPI_ISteamUserStats_GetAchievementName"
	flatAPI_ISteamUserStats_RequestUserStats                    = "SteamAPI_ISteamUserStats_RequestUserStats"
	flatAPI_ISteamUserStats_GetUserStatInt32                    = "SteamAPI_ISteamUserStats_GetUserStatInt32"
	flatAPI_ISteamUserStats_GetUserStatFloat                    = "SteamAPI_ISteamUserStats_GetUserStatFloat"
	flatAPI_ISteamUserStats_GetUserAchievement                  = "SteamAPI_ISteamUserStats_GetUserAchievement"
	flatAPI_ISteamUserStats_GetUserAchievementAndUnlockTime     = "SteamAPI_ISteamUserStats_GetUserAchievementAndUnlockTime"
	flatAPI_ISteamUserStats_ResetAllStats                       = "SteamAPI_ISteamUserStats_ResetAllStats"
	flatAPI_ISteamUserStats_FindOrCreateLeaderboard             = "SteamAPI_ISteamUserStats_FindOrCreateLeaderboard"
	flatAPI_ISteamUserStats_FindLeaderboard                     = "SteamAPI_ISteamUserStats_FindLeaderboard"
	flatAPI_ISteamUserStats_GetLeaderboardName                  = "SteamAPI_ISteamUserStats_GetLeaderboardName"
	flatAPI_ISteamUserStats_GetLeaderboardEntryCount            = "SteamAPI_ISteamUserStats_GetLeaderboardEntryCount"
	flatAPI_ISteamUserStats_GetLeaderboardSortMethod            = "SteamAPI_ISteamUserStats_GetLeaderboardSortMethod"
	flatAPI_ISteamUserStats_GetLeaderboardDisplayType           = "SteamAPI_ISteamUserStats_GetLeaderboardDisplayType"
	flatAPI_ISteamUserStats_DownloadLeaderboardEntries          = "SteamAPI_ISteamUserStats_DownloadLeaderboardEntries"
	flatAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers  = "SteamAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers"
	flatAPI_ISteamUserStats_GetDownloadedLeaderboardEntry       = "SteamAPI_ISteamUserStats_GetDownloadedLeaderboardEntry"
	flatAPI_ISteamUserStats_UploadLeaderboardScore              = "SteamAPI_ISteamUserStats_UploadLeaderboardScore"
	flatAPI_ISteamUserStats_AttachLeaderboardUGC                = "SteamAPI_ISteamUserStats_AttachLeaderboardUGC"
	flatAPI_ISteamUserStats_GetNumberOfCurrentPlayers           = "SteamAPI_ISteamUserStats_GetNumberOfCurrentPlayers"
	flatAPI_ISteamUserStats_RequestGlobalAchievementPercentages = "SteamAPI_ISteamUserStats_RequestGlobalAchievementPercentages"
	flatAPI_ISteamUserStats_GetMostAchievedAchievementInfo      = "SteamAPI_ISteamUserStats_GetMostAchievedAchievementInfo"
	flatAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo  = "SteamAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo"
	flatAPI_ISteamUserStats_GetAchievementAchievedPercent       = "SteamAPI_ISteamUserStats_GetAchievementAchievedPercent"
	flatAPI_ISteamUserStats_RequestGlobalStats                  = "SteamAPI_ISteamUserStats_RequestGlobalStats"
	flatAPI_ISteamUserStats_GetGlobalStatInt64                  = "SteamAPI_ISteamUserStats_GetGlobalStatInt64"
	flatAPI_ISteamUserStats_GetGlobalStatDouble                 = "SteamAPI_ISteamUserStats_GetGlobalStatDouble"
	flatAPI_ISteamUserStats_GetGlobalStatHistoryInt64           = "SteamAPI_ISteamUserStats_GetGlobalStatHistoryInt64"
	flatAPI_ISteamUserStats_GetGlobalStatHistoryDouble          = "SteamAPI_ISteamUserStats_GetGlobalStatHistoryDouble"
	flatAPI_ISteamUserStats_GetAchievementProgressLimitsInt32   = "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsInt32"
	flatAPI_ISteamUserStats_GetAchievementProgressLimitsFloat   = "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsFloat"
)

var (
	steamUserStats_init                                 func() uintptr
	iSteamUserStats_GetStatInt32                        func(steamUserStats uintptr, pchName string, pData *int32) bool
	iSteamUserStats_GetStatFloat                        func(steamUserStats uintptr, pchName string, pData *float32) bool
	iSteamUserStats_SetStatInt32                        func(steamUserStats uintptr, pchName string, nData int32) bool
	iSteamUserStats_SetStatFloat                        func(steamUserStats uintptr, pchName string, fData float32) bool
	iSteamUserStats_UpdateAvgRateStat                   func(steamUserStats uintptr, pchName string, flCountThisSession float32, dSessionLength float64) bool
	iSteamUserStats_GetAchievement                      func(steamUserStats uintptr, pchName string, pbAchieved *bool) bool
	iSteamUserStats_SetAchievement                      func(steamUserStats uintptr, pchName string) bool
	iSteamUserStats_ClearAchievement                    func(steamUserStats uintptr, pchName string) bool
	iSteamUserStats_GetAchievementAndUnlockTime         func(steamUserStats uintptr, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool
	iSteamUserStats_StoreStats                          func(steamUserStats uintptr) bool
	iSteamUserStats_GetAchievementIcon                  func(steamUserStats uintptr, pchName string) int32
	iSteamUserStats_GetAchievementDisplayAttribute      func(steamUserStats uintptr, pchName string, pchKey string) string
	iSteamUserStats_IndicateAchievementProgress         func(steamUserStats uintptr, pchName string, nCurProgress uint32, nMaxProgress uint32) bool
	iSteamUserStats_GetNumAchievements                  func(steamUserStats uintptr) uint32
	iSteamUserStats_GetAchievementName                  func(steamUserStats uintptr, iAchievement uint32) string
	iSteamUserStats_RequestUserStats                    func(steamUserStats uintptr, steamIDUser Uint64SteamID) SteamAPICall
	iSteamUserStats_GetUserStatInt32                    func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pData *int32) bool
	iSteamUserStats_GetUserStatFloat                    func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pData *float32) bool
	iSteamUserStats_GetUserAchievement                  func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool) bool
	iSteamUserStats_GetUserAchievementAndUnlockTime     func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool
	iSteamUserStats_ResetAllStats                       func(steamUserStats uintptr, bAchievementsToo bool) bool
	iSteamUserStats_FindOrCreateLeaderboard             func(steamUserStats uintptr, pchLeaderboardName string, eLeaderboardSortMethod ELeaderboardSortMethod, eLeaderboardDisplayType ELeaderboardDisplayType) SteamAPICall
	iSteamUserStats_FindLeaderboard                     func(steamUserStats uintptr, pchLeaderboardName string) SteamAPICall
	iSteamUserStats_GetLeaderboardName                  func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) string
	iSteamUserStats_GetLeaderboardEntryCount            func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) int32
	iSteamUserStats_GetLeaderboardSortMethod            func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardSortMethod
	iSteamUserStats_GetLeaderboardDisplayType           func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardDisplayType
	iSteamUserStats_DownloadLeaderboardEntries          func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) SteamAPICall
	iSteamUserStats_DownloadLeaderboardEntriesForUsers  func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, prgUsers []CSteamID, cUsers int32) SteamAPICall
	iSteamUserStats_GetDownloadedLeaderboardEntry       func(steamUserStats uintptr, hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, pLeaderboardEntry *LeaderboardEntry, pDetails []int32, cDetailsMax int32) bool
	iSteamUserStats_UploadLeaderboardScore              func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32, cScoreDetailsCount int32) SteamAPICall
	iSteamUserStats_AttachLeaderboardUGC                func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, hUGC UGCHandle) SteamAPICall
	iSteamUserStats_GetNumberOfCurrentPlayers           func(steamUserStats uintptr) SteamAPICall
	iSteamUserStats_RequestGlobalAchievementPercentages func(steamUserStats uintptr) SteamAPICall
	iSteamUserStats_GetMostAchievedAchievementInfo      func(steamUserStats uintptr, pchName []byte, unNameBufLen uint32, pflPercent *float32, pbAchieved *bool) int32
	iSteamUserStats_GetNextMostAchievedAchievementInfo  func(steamUserStats uintptr, iIteratorPrevious int32, pchName []byte, unNameBufLen uint32, pflPercent *float32, pbAchieved *bool) int32
	iSteamUserStats_GetAchievementAchievedPercent       func(steamUserStats uintptr, pchName string, pflPercent *float32) bool
	iSteamUserStats_RequestGlobalStats                  func(steamUserStats uintptr, nHistoryDays int32) SteamAPICall
	iSteamUserStats_GetGlobalStatInt64                  func(steamUserStats uintptr, pchStatName string, pData *int64) bool
	iSteamUserStats_GetGlobalStatDouble                 func(steamUserStats uintptr, pchStatName string, pData *float64) bool
	iSteamUserStats_GetGlobalStatHistoryInt64           func(steamUserStats uintptr, pchStatName string, pData []int64, cubData uint32) int32
	iSteamUserStats_GetGlobalStatHistoryDouble          func(steamUserStats uintptr, pchStatName string, pData []float64, cubData uint32) int32
	iSteamUserStats_GetAchievementProgressLimitsInt32   func(steamUserStats uintptr, pchName string, pnMinProgress *int32, pnMaxProgress *int32) bool
	iSteamUserStats_GetAchievementProgressLimitsFloat   func(steamUserStats uintptr, pchName string, pfMinProgress *float32, pfMaxProgress *float32) bool
)

type steamUserStats uintptr

func SteamUserStats() ISteamUserStats {
	return steamUserStats(steamUserStats_init())
}

func (s steamUserStats) GetStat(name string) (data int32, success bool) {
	success = iSteamUserStats_GetStatInt32(uintptr(s), name, &data)
	return data, success
}

func (s steamUserStats) GetStatFloat(name string) (data float32, success bool) {
	success = iSteamUserStats_GetStatFloat(uintptr(s), name, &data)
	return data, success
}

func (s steamUserStats) SetStat(name string, data int32) bool {
	return iSteamUserStats_SetStatInt32(uintptr(s), name, data)
}

func (s steamUserStats) SetStatFloat(name string, data float32) bool {
	return iSteamUserStats_SetStatFloat(uintptr(s), name, data)
}

func (s steamUserStats) UpdateAvgRateStat(name string, countThisSession float32, sessionLength float64) bool {
	return iSteamUserStats_UpdateAvgRateStat(uintptr(s), name, countThisSession, sessionLength)
}

func (s steamUserStats) GetAchievement(name string) (achieved bool, success bool) {
	success = iSteamUserStats_GetAchievement(uintptr(s), name, &achieved)
	return achieved, success
}

func (s steamUserStats) SetAchievement(name string) bool {
	return iSteamUserStats_SetAchievement(uintptr(s), name)
}

func (s steamUserStats) ClearAchievement(name string) bool {
	return iSteamUserStats_ClearAchievement(uintptr(s), name)
}

func (s steamUserStats) GetAchievementAndUnlockTime(name string) (achieved bool, unlockTime uint32, success bool) {
	success = iSteamUserStats_GetAchievementAndUnlockTime(uintptr(s), name, &achieved, &unlockTime)
	return achieved, unlockTime, success
}

func (s steamUserStats) StoreStats() bool {
	return iSteamUserStats_StoreStats(uintptr(s))
}

func (s steamUserStats) GetAchievementIcon(name string) int32 {
	return iSteamUserStats_GetAchievementIcon(uintptr(s), name)
}

func (s steamUserStats) GetAchievementDisplayAttribute(name string, key string) string {
	return iSteamUserStats_GetAchievementDisplayAttribute(uintptr(s), name, key)
}

func (s steamUserStats) IndicateAchievementProgress(name string, curProgress uint32, maxProgress uint32) bool {
	return iSteamUserStats_IndicateAchievementProgress(uintptr(s), name, curProgress, curProgress)
}

func (s steamUserStats) GetNumAchievements() uint32 {
	return iSteamUserStats_GetNumAchievements(uintptr(s))
}

func (s steamUserStats) GetAchievementName(achievement uint32) string {
	return iSteamUserStats_GetAchievementName(uintptr(s), achievement)
}

func (s steamUserStats) RequestUserStats(userSteamID Uint64SteamID) CallResult[UserStatsReceived] {
	handle := iSteamUserStats_RequestUserStats(uintptr(s), userSteamID)
	return CallResult[UserStatsReceived]{Handle: handle}
}

func (s steamUserStats) GetUserStat(userSteamID Uint64SteamID, name string) (data int32, success bool) {
	success = iSteamUserStats_GetUserStatInt32(uintptr(s), userSteamID, name, &data)
	return data, success
}

func (s steamUserStats) GetUserStatFloat(userSteamID Uint64SteamID, name string) (data float32, success bool) {
	success = iSteamUserStats_GetUserStatFloat(uintptr(s), userSteamID, name, &data)
	return data, success
}

func (s steamUserStats) GetUserAchievement(userSteamID Uint64SteamID, name string) (achieved bool, success bool) {
	success = iSteamUserStats_GetUserAchievement(uintptr(s), userSteamID, name, &achieved)
	return achieved, success
}

func (s steamUserStats) GetUserAchievementAndUnlockTime(userSteamID Uint64SteamID, name string) (achieved bool, unlockTime uint32, success bool) {
	success = iSteamUserStats_GetUserAchievementAndUnlockTime(uintptr(s), userSteamID, name, &achieved, &unlockTime)
	return achieved, unlockTime, success
}

func (s steamUserStats) ResetAllStats(achievementsToo bool) bool {
	return iSteamUserStats_ResetAllStats(uintptr(s), achievementsToo)
}

func (s steamUserStats) FindOrCreateLeaderboard(leaderboardName string, leaderboardSortMethod ELeaderboardSortMethod, leaderboardDisplayType ELeaderboardDisplayType) CallResult[LeaderboardFindResult] {
	handle := iSteamUserStats_FindOrCreateLeaderboard(uintptr(s), leaderboardName, leaderboardSortMethod, leaderboardDisplayType)
	return CallResult[LeaderboardFindResult]{Handle: handle}
}

func (s steamUserStats) FindLeaderboard(leaderboardName string) CallResult[LeaderboardFindResult] {
	handle := iSteamUserStats_FindLeaderboard(uintptr(s), leaderboardName)
	return CallResult[LeaderboardFindResult]{Handle: handle}
}

func (s steamUserStats) GetLeaderboardName(steamLeaderboard SteamLeaderboard) string {
	return iSteamUserStats_GetLeaderboardName(uintptr(s), steamLeaderboard)
}

func (s steamUserStats) GetLeaderboardEntryCount(steamLeaderboard SteamLeaderboard) int32 {
	return iSteamUserStats_GetLeaderboardEntryCount(uintptr(s), steamLeaderboard)
}

func (s steamUserStats) GetLeaderboardSortMethod(steamLeaderboard SteamLeaderboard) ELeaderboardSortMethod {
	return iSteamUserStats_GetLeaderboardSortMethod(uintptr(s), steamLeaderboard)
}

func (s steamUserStats) GetLeaderboardDisplayType(steamLeaderboard SteamLeaderboard) ELeaderboardDisplayType {
	return iSteamUserStats_GetLeaderboardDisplayType(uintptr(s), steamLeaderboard)
}

func (s steamUserStats) DownloadLeaderboardEntries(steamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) CallResult[LeaderboardScoresDownloaded] {
	handle := iSteamUserStats_DownloadLeaderboardEntries(uintptr(s), steamLeaderboard, eLeaderboardDataRequest, nRangeStart, nRangeEnd)
	return CallResult[LeaderboardScoresDownloaded]{Handle: handle}
}

func (s steamUserStats) DownloadLeaderboardEntriesForUsers(steamLeaderboard SteamLeaderboard, prgUsers []CSteamID) CallResult[LeaderboardScoresDownloaded] {
	handle := iSteamUserStats_DownloadLeaderboardEntriesForUsers(uintptr(s), steamLeaderboard, prgUsers, int32(len(prgUsers)))
	return CallResult[LeaderboardScoresDownloaded]{Handle: handle}
}

func (s steamUserStats) GetDownloadedLeaderboardEntry(hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, cDetailsMax int32) (pLeaderboardEntry LeaderboardEntry, pDetails []int32, success bool) {
	pDetails = make([]int32, 0, cDetailsMax)
	success = iSteamUserStats_GetDownloadedLeaderboardEntry(uintptr(s), hSteamLeaderboardEntries, index, &pLeaderboardEntry, pDetails, cDetailsMax)
	return pLeaderboardEntry, pDetails, success
}

func (s steamUserStats) UploadLeaderboardScore(steamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32) CallResult[LeaderboardScoreUploaded] {
	handle := iSteamUserStats_UploadLeaderboardScore(uintptr(s), steamLeaderboard, eLeaderboardUploadScoreMethod, nScore, pScoreDetails, int32(len(pScoreDetails)))
	return CallResult[LeaderboardScoreUploaded]{Handle: handle}
}

func (s steamUserStats) AttachLeaderboardUGC(steamLeaderboard SteamLeaderboard, hUGC UGCHandle) CallResult[LeaderboardUGCSet] {
	handle := iSteamUserStats_AttachLeaderboardUGC(uintptr(s), steamLeaderboard, hUGC)
	return CallResult[LeaderboardUGCSet]{Handle: handle}
}

func (s steamUserStats) GetNumberOfCurrentPlayers() CallResult[NumberOfCurrentPlayers] {
	handle := iSteamUserStats_GetNumberOfCurrentPlayers(uintptr(s))
	return CallResult[NumberOfCurrentPlayers]{Handle: handle}
}

func (s steamUserStats) RequestGlobalAchievementPercentages() CallResult[GlobalAchievementPercentagesReady] {
	handle := iSteamUserStats_RequestGlobalAchievementPercentages(uintptr(s))
	return CallResult[GlobalAchievementPercentagesReady]{Handle: handle}
}

func (s steamUserStats) GetMostAchievedAchievementInfo(nameBufLen uint32) (name string, percent float32, achieved bool, index int32) {
	buf := make([]byte, 0, nameBufLen)

	index = iSteamUserStats_GetMostAchievedAchievementInfo(uintptr(s), buf, nameBufLen, &percent, &achieved)
	return string(buf), percent, achieved, index
}

func (s steamUserStats) GetNextMostAchievedAchievementInfo(iteratorPrevious int32, nameBufLen uint32) (name string, percent float32, achieved bool, index int32) {
	buf := make([]byte, 0, nameBufLen)

	index = iSteamUserStats_GetNextMostAchievedAchievementInfo(uintptr(s), iteratorPrevious, buf, nameBufLen, &percent, &achieved)
	return string(buf), percent, achieved, index
}

func (s steamUserStats) GetAchievementAchievedPercent(name string) (percent float32, success bool) {
	success = iSteamUserStats_GetAchievementAchievedPercent(uintptr(s), name, &percent)
	return percent, success
}

func (s steamUserStats) RequestGlobalStats(historyDays int32) CallResult[GlobalStatsReceived] {
	handle := iSteamUserStats_RequestGlobalStats(uintptr(s), historyDays)
	return CallResult[GlobalStatsReceived]{Handle: handle}
}

func (s steamUserStats) GetGlobalStat(pchStatName string) (data int64, success bool) {
	success = iSteamUserStats_GetGlobalStatInt64(uintptr(s), pchStatName, &data)
	return data, success
}

func (s steamUserStats) GetGlobalStatDouble(pchStatName string) (data float64, success bool) {
	success = iSteamUserStats_GetGlobalStatDouble(uintptr(s), pchStatName, &data)
	return data, success
}

func (s steamUserStats) GetGlobalStatHistory(pchStatName string, dataSize uint32) (data []int64) {
	data = make([]int64, 0, dataSize)
	result := iSteamUserStats_GetGlobalStatHistoryInt64(uintptr(s), pchStatName, data, dataSize)
	return data[:result]
}

func (s steamUserStats) GetGlobalStatHistoryDouble(statName string, dataSize uint32) (data []float64) {
	data = make([]float64, 0, dataSize)
	result := iSteamUserStats_GetGlobalStatHistoryDouble(uintptr(s), statName, data, dataSize)
	return data[:result]
}

func (s steamUserStats) GetAchievementProgressLimits(name string) (minProgress int32, maxProgress int32, success bool) {
	success = iSteamUserStats_GetAchievementProgressLimitsInt32(uintptr(s), name, &minProgress, &maxProgress)
	return minProgress, maxProgress, success
}

func (s steamUserStats) GetAchievementProgressLimitsFloat(name string) (minProgress float32, maxProgress float32, success bool) {
	success = iSteamUserStats_GetAchievementProgressLimitsFloat(uintptr(s), name, &minProgress, &maxProgress)
	return minProgress, maxProgress, success
}

// Steam Utils
type ESteamIPv6ConnectivityState int32

const (
	ESteamIPv6ConnectivityState_Unknown ESteamIPv6ConnectivityState = 0
	ESteamIPv6ConnectivityState_Good    ESteamIPv6ConnectivityState = 1
	ESteamIPv6ConnectivityState_Bad     ESteamIPv6ConnectivityState = 2
)

type EFloatingGamepadTextInputMode int32

const (
	EFloatingGamepadTextInputMode_ModeSingleLine    EFloatingGamepadTextInputMode = 0
	EFloatingGamepadTextInputMode_ModeMultipleLines EFloatingGamepadTextInputMode = 1
	EFloatingGamepadTextInputMode_ModeEmail         EFloatingGamepadTextInputMode = 2
	EFloatingGamepadTextInputMode_ModeNumeric       EFloatingGamepadTextInputMode = 3
)

type ETextFilteringContext int32

const (
	ETextFilteringContext_Unknown     ETextFilteringContext = 0
	ETextFilteringContext_GameContent ETextFilteringContext = 1
	ETextFilteringContext_Chat        ETextFilteringContext = 2
	ETextFilteringContext_Name        ETextFilteringContext = 3
)

type ESteamIPv6ConnectivityProtocol int32

const (
	ESteamIPv6ConnectivityProtocol_Invalid ESteamIPv6ConnectivityProtocol = 0
	ESteamIPv6ConnectivityProtocol_HTTP    ESteamIPv6ConnectivityProtocol = 1
	ESteamIPv6ConnectivityProtocol_UDP     ESteamIPv6ConnectivityProtocol = 2
)

type ENotificationPosition int32

const (
	EPosition_Invalid     ENotificationPosition = -1
	EPosition_TopLeft     ENotificationPosition = 0
	EPosition_TopRight    ENotificationPosition = 1
	EPosition_BottomLeft  ENotificationPosition = 2
	EPosition_BottomRight ENotificationPosition = 3
)

type EGamepadTextInputMode int32

const (
	EGamepadTextInputMode_Normal   EGamepadTextInputMode = 0
	EGamepadTextInputMode_Password EGamepadTextInputMode = 1
)

type EGamepadTextInputLineMode int32

const (
	EGamepadTextInputLineMode_SingleLine    EGamepadTextInputLineMode = 0
	EGamepadTextInputLineMode_MultipleLines EGamepadTextInputLineMode = 1
)

var (
	steamUtils_init                             func() uintptr
	iSteamUtils_GetSecondsSinceAppActive        func(steamUtils uintptr) uint32
	iSteamUtils_GetSecondsSinceComputerActive   func(steamUtils uintptr) uint32
	iSteamUtils_GetConnectedUniverse            func(steamUtils uintptr) EUniverse
	iSteamUtils_GetServerRealTime               func(steamUtils uintptr) uint32
	iSteamUtils_GetIPCountry                    func(steamUtils uintptr) string
	iSteamUtils_GetImageSize                    func(steamUtils uintptr, iImage int32, pnWidth *uint32, pnHeight *uint32) bool
	iSteamUtils_GetImageRGBA                    func(steamUtils uintptr, iImage int32, pubDest []byte, nDestBufferSize int32) bool
	iSteamUtils_GetCurrentBatteryPower          func(steamUtils uintptr) uint8
	iSteamUtils_GetAppID                        func(steamUtils uintptr) uint32
	iSteamUtils_SetOverlayNotificationPosition  func(steamUtils uintptr, eNotificationPosition ENotificationPosition)
	iSteamUtils_IsAPICallCompleted              func(steamUtils uintptr, hSteamAPICall SteamAPICall, pbFailed *bool) bool
	iSteamUtils_GetAPICallFailureReason         func(steamUtils uintptr, hSteamAPICall SteamAPICall) ESteamAPICallFailure
	iSteamUtils_GetAPICallResult                func(steamUtils uintptr, hSteamAPICall SteamAPICall, pCallback []byte, cubCallback int32, iCallbackExpected int32, pbFailed *bool) bool
	iSteamUtils_GetIPCCallCount                 func(steamUtils uintptr) uint32
	iSteamUtils_SetWarningMessageHook           func(steamUtils uintptr, pFunction SteamAPIWarningMessageHook)
	iSteamUtils_IsOverlayEnabled                func(steamUtils uintptr) bool
	iSteamUtils_BOverlayNeedsPresent            func(steamUtils uintptr) bool
	iSteamUtils_CheckFileSignature              func(steamUtils uintptr, szFileName string) SteamAPICall
	iSteamUtils_ShowGamepadTextInput            func(steamUtils uintptr, eInputMode EGamepadTextInputMode, eLineInputMode EGamepadTextInputLineMode, pchDescription string, unCharMax uint32, pchExistingText string) bool
	iSteamUtils_GetEnteredGamepadTextLength     func(steamUtils uintptr) uint32
	iSteamUtils_GetEnteredGamepadTextInput      func(steamUtils uintptr, pchText []byte, cchText uint32) bool
	iSteamUtils_GetSteamUILanguage              func(steamUtils uintptr) string
	iSteamUtils_IsSteamRunningInVR              func(steamUtils uintptr) bool
	iSteamUtils_SetOverlayNotificationInset     func(steamUtils uintptr, nHorizontalInset int32, nVerticalInset int32)
	iSteamUtils_IsSteamInBigPictureMode         func(steamUtils uintptr) bool
	iSteamUtils_StartVRDashboard                func(steamUtils uintptr)
	iSteamUtils_IsVRHeadsetStreamingEnabled     func(steamUtils uintptr) bool
	iSteamUtils_SetVRHeadsetStreamingEnabled    func(steamUtils uintptr, bEnabled bool)
	iSteamUtils_IsSteamChinaLauncher            func(steamUtils uintptr) bool
	iSteamUtils_InitFilterText                  func(steamUtils uintptr, unFilterOptions uint32) bool
	iSteamUtils_FilterText                      func(steamUtils uintptr, eContext ETextFilteringContext, sourceSteamID Uint64SteamID, pchInputMessage string, pchOutFilteredText []byte, nByteSizeOutFilteredText uint32) int32
	iSteamUtils_GetIPv6ConnectivityState        func(steamUtils uintptr, eProtocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState
	iSteamUtils_IsSteamRunningOnSteamDeck       func(steamUtils uintptr) bool
	iSteamUtils_ShowFloatingGamepadTextInput    func(steamUtils uintptr, eKeyboardMode EFloatingGamepadTextInputMode, nTextFieldXPosition int32, nTextFieldYPosition int32, nTextFieldWidth int32, nTextFieldHeight int32) bool
	iSteamUtils_SetGameLauncherMode             func(steamUtils uintptr, bLauncherMode bool)
	iSteamUtils_DismissFloatingGamepadTextInput func(steamUtils uintptr) bool
	iSteamUtils_DismissGamepadTextInput         func(steamUtils uintptr) bool
)

const (
	flatAPI_SteamUtils                                  = "SteamAPI_SteamUtils_v010"
	flatAPI_ISteamUtils_GetSecondsSinceAppActive        = "SteamAPI_ISteamUtils_GetSecondsSinceAppActive"
	flatAPI_ISteamUtils_GetSecondsSinceComputerActive   = "SteamAPI_ISteamUtils_GetSecondsSinceComputerActive"
	flatAPI_ISteamUtils_GetConnectedUniverse            = "SteamAPI_ISteamUtils_GetConnectedUniverse"
	flatAPI_ISteamUtils_GetServerRealTime               = "SteamAPI_ISteamUtils_GetServerRealTime"
	flatAPI_ISteamUtils_GetIPCountry                    = "SteamAPI_ISteamUtils_GetIPCountry"
	flatAPI_ISteamUtils_GetImageSize                    = "SteamAPI_ISteamUtils_GetImageSize"
	flatAPI_ISteamUtils_GetImageRGBA                    = "SteamAPI_ISteamUtils_GetImageRGBA"
	flatAPI_ISteamUtils_GetCurrentBatteryPower          = "SteamAPI_ISteamUtils_GetCurrentBatteryPower"
	flatAPI_ISteamUtils_GetAppID                        = "SteamAPI_ISteamUtils_GetAppID"
	flatAPI_ISteamUtils_SetOverlayNotificationPosition  = "SteamAPI_ISteamUtils_SetOverlayNotificationPosition"
	flatAPI_ISteamUtils_IsAPICallCompleted              = "SteamAPI_ISteamUtils_IsAPICallCompleted"
	flatAPI_ISteamUtils_GetAPICallFailureReason         = "SteamAPI_ISteamUtils_GetAPICallFailureReason"
	flatAPI_ISteamUtils_GetAPICallResult                = "SteamAPI_ISteamUtils_GetAPICallResult"
	flatAPI_ISteamUtils_GetIPCCallCount                 = "SteamAPI_ISteamUtils_GetIPCCallCount"
	flatAPI_ISteamUtils_SetWarningMessageHook           = "SteamAPI_ISteamUtils_SetWarningMessageHook"
	flatAPI_ISteamUtils_IsOverlayEnabled                = "SteamAPI_ISteamUtils_IsOverlayEnabled"
	flatAPI_ISteamUtils_BOverlayNeedsPresent            = "SteamAPI_ISteamUtils_BOverlayNeedsPresent"
	flatAPI_ISteamUtils_CheckFileSignature              = "SteamAPI_ISteamUtils_CheckFileSignature"
	flatAPI_ISteamUtils_ShowGamepadTextInput            = "SteamAPI_ISteamUtils_ShowGamepadTextInput"
	flatAPI_ISteamUtils_GetEnteredGamepadTextLength     = "SteamAPI_ISteamUtils_GetEnteredGamepadTextLength"
	flatAPI_ISteamUtils_GetEnteredGamepadTextInput      = "SteamAPI_ISteamUtils_GetEnteredGamepadTextInput"
	flatAPI_ISteamUtils_GetSteamUILanguage              = "SteamAPI_ISteamUtils_GetSteamUILanguage"
	flatAPI_ISteamUtils_IsSteamRunningInVR              = "SteamAPI_ISteamUtils_IsSteamRunningInVR"
	flatAPI_ISteamUtils_SetOverlayNotificationInset     = "SteamAPI_ISteamUtils_SetOverlayNotificationInset"
	flatAPI_ISteamUtils_IsSteamInBigPictureMode         = "SteamAPI_ISteamUtils_IsSteamInBigPictureMode"
	flatAPI_ISteamUtils_StartVRDashboard                = "SteamAPI_ISteamUtils_StartVRDashboard"
	flatAPI_ISteamUtils_IsVRHeadsetStreamingEnabled     = "SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled"
	flatAPI_ISteamUtils_SetVRHeadsetStreamingEnabled    = "SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled"
	flatAPI_ISteamUtils_IsSteamChinaLauncher            = "SteamAPI_ISteamUtils_IsSteamChinaLauncher"
	flatAPI_ISteamUtils_InitFilterText                  = "SteamAPI_ISteamUtils_InitFilterText"
	flatAPI_ISteamUtils_FilterText                      = "SteamAPI_ISteamUtils_FilterText"
	flatAPI_ISteamUtils_GetIPv6ConnectivityState        = "SteamAPI_ISteamUtils_GetIPv6ConnectivityState"
	flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck       = "SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck"
	flatAPI_ISteamUtils_ShowFloatingGamepadTextInput    = "SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput"
	flatAPI_ISteamUtils_SetGameLauncherMode             = "SteamAPI_ISteamUtils_SetGameLauncherMode"
	flatAPI_ISteamUtils_DismissFloatingGamepadTextInput = "SteamAPI_ISteamUtils_DismissFloatingGamepadTextInput"
	flatAPI_ISteamUtils_DismissGamepadTextInput         = "SteamAPI_ISteamUtils_DismissGamepadTextInput"
)

type steamUtils uintptr

func SteamUtils() ISteamUtils {
	return steamUtils(steamUtils_init())
}

func (s steamUtils) GetSecondsSinceAppActive() uint32 {
	return iSteamUtils_GetSecondsSinceAppActive(uintptr(s))
}

func (s steamUtils) GetSecondsSinceComputerActive() uint32 {
	return iSteamUtils_GetSecondsSinceComputerActive(uintptr(s))
}

func (s steamUtils) GetConnectedUniverse() EUniverse {
	return iSteamUtils_GetConnectedUniverse(uintptr(s))
}

func (s steamUtils) GetServerRealTime() uint32 {
	return iSteamUtils_GetServerRealTime(uintptr(s))
}

func (s steamUtils) GetIPCountry() string {
	return iSteamUtils_GetIPCountry(uintptr(s))
}

func (s steamUtils) GetImageSize(imageIndex int32) (uint32, uint32, bool) {
	var pnWidth, pnHeight uint32
	result := iSteamUtils_GetImageSize(uintptr(s), imageIndex, &pnWidth, &pnHeight)
	return pnWidth, pnHeight, result
}

func (s steamUtils) GetImageRGBA(imageIndex int32, destBufferSize int32) ([]byte, bool) {
	var pubDest []byte = make([]byte, 0, destBufferSize)
	result := iSteamUtils_GetImageRGBA(uintptr(s), imageIndex, pubDest, destBufferSize)
	return pubDest, result
}

func (s steamUtils) GetCurrentBatteryPower() uint8 {
	return iSteamUtils_GetCurrentBatteryPower(uintptr(s))
}

func (s steamUtils) GetAppID() uint32 {
	return iSteamUtils_GetAppID(uintptr(s))
}

func (s steamUtils) SetOverlayNotificationPosition(notificationPosition ENotificationPosition) {
	iSteamUtils_SetOverlayNotificationPosition(uintptr(s), notificationPosition)
}

func (s steamUtils) IsAPICallCompleted(steamAPICallHandle SteamAPICall) (failed bool, completed bool) {
	var pbFailed bool
	result := iSteamUtils_IsAPICallCompleted(uintptr(s), steamAPICallHandle, &pbFailed)
	return pbFailed, result
}

func (s steamUtils) GetAPICallFailureReason(steamAPICallHandle SteamAPICall) ESteamAPICallFailure {
	return iSteamUtils_GetAPICallFailureReason(uintptr(s), steamAPICallHandle)
}

func (s steamUtils) GetAPICallResult(steamAPICallHandle SteamAPICall, callbackDataSize int32, callbackExpectedID int32) (callbackData []byte, failed bool, success bool) {
	callbackData = make([]byte, 0, callbackDataSize)
	success = iSteamUtils_GetAPICallResult(uintptr(s), steamAPICallHandle, callbackData, callbackDataSize, callbackExpectedID, &failed)
	return callbackData, failed, success
}

func (s steamUtils) GetIPCCallCount() uint32 {
	return iSteamUtils_GetIPCCallCount(uintptr(s))
}

func (s steamUtils) SetWarningMessageHook(function SteamAPIWarningMessageHook) {
	iSteamUtils_SetWarningMessageHook(uintptr(s), function)
}

func (s steamUtils) IsOverlayEnabled() bool {
	return iSteamUtils_IsOverlayEnabled(uintptr(s))
}

func (s steamUtils) BOverlayNeedsPresent() bool {
	return iSteamUtils_BOverlayNeedsPresent(uintptr(s))
}

func (s steamUtils) CheckFileSignature(fileName string) CallResult[CheckFileSignatureResult] {
	handle := iSteamUtils_CheckFileSignature(uintptr(s), fileName)
	return CallResult[CheckFileSignatureResult]{Handle: handle}
}

func (s steamUtils) ShowGamepadTextInput(inputMode EGamepadTextInputMode, lineInputMode EGamepadTextInputLineMode, description string, charMax uint32, existingText string) bool {
	return iSteamUtils_ShowGamepadTextInput(uintptr(s), inputMode, lineInputMode, description, charMax, existingText)
}

func (s steamUtils) GetEnteredGamepadTextLength() uint32 {
	return iSteamUtils_GetEnteredGamepadTextLength(uintptr(s))
}

func (s steamUtils) GetEnteredGamepadTextInput(text uint32) (string, bool) {
	var pchText []byte = make([]byte, 0, text)
	result := iSteamUtils_GetEnteredGamepadTextInput(uintptr(s), pchText, text)
	return string(pchText), result
}

func (s steamUtils) GetSteamUILanguage() string {
	return iSteamUtils_GetSteamUILanguage(uintptr(s))
}

func (s steamUtils) IsSteamRunningInVR() bool {
	return iSteamUtils_IsSteamRunningInVR(uintptr(s))
}

func (s steamUtils) SetOverlayNotificationInset(horizontalInset int32, verticalInset int32) {
	iSteamUtils_SetOverlayNotificationInset(uintptr(s), horizontalInset, verticalInset)
}

func (s steamUtils) IsSteamInBigPictureMode() bool {
	return iSteamUtils_IsSteamInBigPictureMode(uintptr(s))
}

func (s steamUtils) StartVRDashboard() {
	iSteamUtils_StartVRDashboard(uintptr(s))
}

func (s steamUtils) IsVRHeadsetStreamingEnabled() bool {
	return iSteamUtils_IsVRHeadsetStreamingEnabled(uintptr(s))
}

func (s steamUtils) SetVRHeadsetStreamingEnabled(enabled bool) {
	iSteamUtils_SetVRHeadsetStreamingEnabled(uintptr(s), enabled)
}

func (s steamUtils) IsSteamChinaLauncher() bool {
	return iSteamUtils_IsSteamChinaLauncher(uintptr(s))
}

func (s steamUtils) InitFilterText(filterOptions uint32) bool {
	return iSteamUtils_InitFilterText(uintptr(s), filterOptions)
}

func (s steamUtils) FilterText(context ETextFilteringContext, sourceSteamID Uint64SteamID, inputMessage string, byteSizeOutFilteredText uint32) (string, int32) {
	if int(byteSizeOutFilteredText) < len(inputMessage)+1 {
		byteSizeOutFilteredText = uint32(len(inputMessage) + 1)
	}
	var pchOutFilteredText []byte = make([]byte, 0, byteSizeOutFilteredText)
	result := iSteamUtils_FilterText(uintptr(s), context, sourceSteamID, inputMessage, pchOutFilteredText, byteSizeOutFilteredText)
	return string(pchOutFilteredText), result
}

func (s steamUtils) GetIPv6ConnectivityState(protocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState {
	return iSteamUtils_GetIPv6ConnectivityState(uintptr(s), protocol)
}

func (s steamUtils) IsSteamRunningOnSteamDeck() bool {
	return iSteamUtils_IsSteamRunningOnSteamDeck(uintptr(s))
}

func (s steamUtils) ShowFloatingGamepadTextInput(keyboardMode EFloatingGamepadTextInputMode, textFieldXPosition int32, textFieldYPosition int32, textFieldWidth int32, textFieldHeight int32) bool {
	return iSteamUtils_ShowFloatingGamepadTextInput(uintptr(s), keyboardMode, textFieldXPosition, textFieldYPosition, textFieldWidth, textFieldHeight)
}

func (s steamUtils) SetGameLauncherMode(launcherMode bool) {
	iSteamUtils_SetGameLauncherMode(uintptr(s), launcherMode)
}

func (s steamUtils) DismissFloatingGamepadTextInput() bool {
	return iSteamUtils_DismissFloatingGamepadTextInput(uintptr(s))
}

func (s steamUtils) DismissGamepadTextInput() bool {
	return iSteamUtils_DismissGamepadTextInput(uintptr(s))
}
