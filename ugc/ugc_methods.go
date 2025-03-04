package ugc

import (
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"

	. "github.com/assemblaj/purego-steamworks"
)

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
	EUGCContentDescriptorNudityOrSexualContent   EUGCContentDescriptorID = 1
	EUGCContentDescriptorFrequentViolenceOrGore  EUGCContentDescriptorID = 2
	EUGCContentDescriptorAdultOnlySexualContent  EUGCContentDescriptorID = 3
	EUGCContentDescriptorGratuitousSexualContent EUGCContentDescriptorID = 4
	EUGCContentDescriptorAnyMatureContent        EUGCContentDescriptorID = 5
)

type EUserUGCList int32

const (
	EUserUGCListPublished     EUserUGCList = 0
	EUserUGCListVotedOn       EUserUGCList = 1
	EUserUGCListVotedUp       EUserUGCList = 2
	EUserUGCListVotedDown     EUserUGCList = 3
	EUserUGCListWillVoteLater EUserUGCList = 4
	EUserUGCListFavorited     EUserUGCList = 5
	EUserUGCListSubscribed    EUserUGCList = 6
	EUserUGCListUsedOrPlayed  EUserUGCList = 7
	EUserUGCListFollowed      EUserUGCList = 8
)

type EUGCQuery int32

const (
	EUGCQueryRankedByVote                                  EUGCQuery = 0
	EUGCQueryRankedByPublicationDate                       EUGCQuery = 1
	EUGCQueryAcceptedForGameRankedByAcceptanceDate         EUGCQuery = 2
	EUGCQueryRankedByTrend                                 EUGCQuery = 3
	EUGCQueryFavoritedByFriendsRankedByPublicationDate     EUGCQuery = 4
	EUGCQueryCreatedByFriendsRankedByPublicationDate       EUGCQuery = 5
	EUGCQueryRankedByNumTimesReported                      EUGCQuery = 6
	EUGCQueryCreatedByFollowedUsersRankedByPublicationDate EUGCQuery = 7
	EUGCQueryNotYetRated                                   EUGCQuery = 8
	EUGCQueryRankedByTotalVotesAsc                         EUGCQuery = 9
	EUGCQueryRankedByVotesUp                               EUGCQuery = 10
	EUGCQueryRankedByTextSearch                            EUGCQuery = 11
	EUGCQueryRankedByTotalUniqueSubscriptions              EUGCQuery = 12
	EUGCQueryRankedByPlaytimeTrend                         EUGCQuery = 13
	EUGCQueryRankedByTotalPlaytime                         EUGCQuery = 14
	EUGCQueryRankedByAveragePlaytimeTrend                  EUGCQuery = 15
	EUGCQueryRankedByLifetimeAveragePlaytime               EUGCQuery = 16
	EUGCQueryRankedByPlaytimeSessionsTrend                 EUGCQuery = 17
	EUGCQueryRankedByLifetimePlaytimeSessions              EUGCQuery = 18
	EUGCQueryRankedByLastUpdatedDate                       EUGCQuery = 19
)

type EItemState int32

const (
	EItemStateNone            EItemState = 0
	EItemStateSubscribed      EItemState = 1
	EItemStateLegacyItem      EItemState = 2
	EItemStateInstalled       EItemState = 4
	EItemStateNeedsUpdate     EItemState = 8
	EItemStateDownloading     EItemState = 16
	EItemStateDownloadPending EItemState = 32
	EItemStateDisabledLocally EItemState = 64
)

type SteamUGCDetails struct {
	PublishedFileId     PublishedFileId
	Result              EResult
	FileType            EWorkshopFileType
	CreatorAppID        AppId
	ConsumerAppID       AppId
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
	EUGCMatchingUGCTypeItems              EUGCMatchingUGCType = 0
	EUGCMatchingUGCTypeItemsMtx           EUGCMatchingUGCType = 1
	EUGCMatchingUGCTypeItemsReadyToUse    EUGCMatchingUGCType = 2
	EUGCMatchingUGCTypeCollections        EUGCMatchingUGCType = 3
	EUGCMatchingUGCTypeArtwork            EUGCMatchingUGCType = 4
	EUGCMatchingUGCTypeVideos             EUGCMatchingUGCType = 5
	EUGCMatchingUGCTypeScreenshots        EUGCMatchingUGCType = 6
	EUGCMatchingUGCTypeAllGuides          EUGCMatchingUGCType = 7
	EUGCMatchingUGCTypeWebGuides          EUGCMatchingUGCType = 8
	EUGCMatchingUGCTypeIntegratedGuides   EUGCMatchingUGCType = 9
	EUGCMatchingUGCTypeUsableInGame       EUGCMatchingUGCType = 10
	EUGCMatchingUGCTypeControllerBindings EUGCMatchingUGCType = 11
	EUGCMatchingUGCTypeGameManagedItems   EUGCMatchingUGCType = 12
	EUGCMatchingUGCTypeAll                EUGCMatchingUGCType = -1
)

type EUserUGCListSortOrder int32

const (
	EUserUGCListSortOrderCreationOrderDesc    EUserUGCListSortOrder = 0
	EUserUGCListSortOrderCreationOrderAsc     EUserUGCListSortOrder = 1
	EUserUGCListSortOrderTitleAsc             EUserUGCListSortOrder = 2
	EUserUGCListSortOrderLastUpdatedDesc      EUserUGCListSortOrder = 3
	EUserUGCListSortOrderSubscriptionDateDesc EUserUGCListSortOrder = 4
	EUserUGCListSortOrderVoteScoreDesc        EUserUGCListSortOrder = 5
	EUserUGCListSortOrderForModeration        EUserUGCListSortOrder = 6
)

type EItemPreviewType int32

const (
	EItemPreviewTypeImage                         EItemPreviewType = 0
	EItemPreviewTypeYouTubeVideo                  EItemPreviewType = 1
	EItemPreviewTypeSketchfab                     EItemPreviewType = 2
	EItemPreviewTypeEnvironmentMapHorizontalCross EItemPreviewType = 3
	EItemPreviewTypeEnvironmentMapLatLong         EItemPreviewType = 4
	EItemPreviewTypeClip                          EItemPreviewType = 5
	EItemPreviewTypeReservedMax                   EItemPreviewType = 255
)

type EItemUpdateStatus int32

const (
	EItemUpdateStatusInvalid              EItemUpdateStatus = 0
	EItemUpdateStatusPreparingConfig      EItemUpdateStatus = 1
	EItemUpdateStatusPreparingContent     EItemUpdateStatus = 2
	EItemUpdateStatusUploadingContent     EItemUpdateStatus = 3
	EItemUpdateStatusUploadingPreviewFile EItemUpdateStatus = 4
	EItemUpdateStatusCommittingChanges    EItemUpdateStatus = 5
)

type EItemStatistic int32

const (
	EItemStatisticNumSubscriptions                    EItemStatistic = 0
	EItemStatisticNumFavorites                        EItemStatistic = 1
	EItemStatisticNumFollowers                        EItemStatistic = 2
	EItemStatisticNumUniqueSubscriptions              EItemStatistic = 3
	EItemStatisticNumUniqueFavorites                  EItemStatistic = 4
	EItemStatisticNumUniqueFollowers                  EItemStatistic = 5
	EItemStatisticNumUniqueWebsiteViews               EItemStatistic = 6
	EItemStatisticReportScore                         EItemStatistic = 7
	EItemStatisticNumSecondsPlayed                    EItemStatistic = 8
	EItemStatisticNumPlaytimeSessions                 EItemStatistic = 9
	EItemStatisticNumComments                         EItemStatistic = 10
	EItemStatisticNumSecondsPlayedDuringTimePeriod    EItemStatistic = 11
	EItemStatisticNumPlaytimeSessionsDuringTimePeriod EItemStatistic = 12
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
	iSteamUGC_CreateQueryUserUGCRequest           func(steamUGC uintptr, unAccountID AccountID, eListType EUserUGCList, eMatchingUGCType EUGCMatchingUGCType, eSortOrder EUserUGCListSortOrder, nCreatorAppID AppId, nConsumerAppID AppId, unPage uint32) UGCQueryHandle
	iSteamUGC_CreateQueryAllUGCRequestPage        func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId, nConsumerAppID AppId, unPage uint32) UGCQueryHandle
	iSteamUGC_CreateQueryAllUGCRequestCursor      func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId, nConsumerAppID AppId, pchCursor string) UGCQueryHandle
	iSteamUGC_CreateQueryUGCDetailsRequest        func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) UGCQueryHandle
	iSteamUGC_SendQueryUGCRequest                 func(steamUGC uintptr, handle UGCQueryHandle) SteamAPICall
	iSteamUGC_GetQueryUGCResult                   func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pDetails *SteamUGCDetails) bool
	iSteamUGC_GetQueryUGCNumTags                  func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCTag                      func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue *[]byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryUGCTagDisplayName           func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue *[]byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryUGCPreviewURL               func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchURL *[]byte, cchURLSize uint32) bool
	iSteamUGC_GetQueryUGCMetadata                 func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchMetadata *[]byte, cchMetadatasize uint32) bool
	iSteamUGC_GetQueryUGCChildren                 func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecPublishedFileID *[]PublishedFileId, cMaxEntries uint32) bool
	iSteamUGC_GetQueryUGCStatistic                func(steamUGC uintptr, handle UGCQueryHandle, index uint32, eStatType EItemStatistic, pStatValue *uint64) bool
	iSteamUGC_GetQueryUGCNumAdditionalPreviews    func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCAdditionalPreview        func(steamUGC uintptr, handle UGCQueryHandle, index uint32, previewIndex uint32, pchURLOrVideoID *[]byte, cchURLSize uint32, pchOriginalFileName *[]byte, cchOriginalFileNameSize uint32, pPreviewType *EItemPreviewType) bool
	iSteamUGC_GetQueryUGCNumKeyValueTags          func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetQueryUGCKeyValueTag              func(steamUGC uintptr, handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, pchKey *[]byte, cchKeySize uint32, pchValue *[]byte, cchValueSize uint32) bool
	iSteamUGC_GetQueryFirstUGCKeyValueTag         func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchKey string, pchValue *[]byte, cchValueSize uint32) bool
	iSteamUGC_GetNumSupportedGameVersions         func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32
	iSteamUGC_GetSupportedGameVersionData         func(steamUGC uintptr, handle UGCQueryHandle, index uint32, versionIndex uint32, pchGameBranchMin *[]byte, pchGameBranchMax *[]byte, cchGameBranchSize uint32) bool
	iSteamUGC_GetQueryUGCContentDescriptors       func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecDescriptors *[]EUGCContentDescriptorID, cMaxEntries uint32) uint32
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
	iSteamUGC_CreateItem                          func(steamUGC uintptr, nConsumerAppId AppId, eFileType EWorkshopFileType) SteamAPICall
	iSteamUGC_StartItemUpdate                     func(steamUGC uintptr, nConsumerAppId AppId, nPublishedFileID PublishedFileId) UGCUpdateHandle
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
	iSteamUGC_AddItemToFavorites                  func(steamUGC uintptr, nAppId AppId, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_RemoveItemFromFavorites             func(steamUGC uintptr, nAppId AppId, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_SubscribeItem                       func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_UnsubscribeItem                     func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_GetNumSubscribedItems               func(steamUGC uintptr) uint32
	iSteamUGC_GetSubscribedItems                  func(steamUGC uintptr, pvecPublishedFileID *[]PublishedFileId, cMaxEntries uint32) uint32
	iSteamUGC_GetItemState                        func(steamUGC uintptr, nPublishedFileID PublishedFileId) uint32
	iSteamUGC_GetItemInstallInfo                  func(steamUGC uintptr, nPublishedFileID PublishedFileId, punSizeOnDisk *uint64, pchFolder *[]byte, cchFolderSize uint32, punTimeStamp *uint32) bool
	iSteamUGC_GetItemDownloadInfo                 func(steamUGC uintptr, nPublishedFileID PublishedFileId, punBytesDownloaded *uint64, punBytesTotal *uint64) bool
	iSteamUGC_DownloadItem                        func(steamUGC uintptr, nPublishedFileID PublishedFileId, bHighPriority bool) bool
	iSteamUGC_BInitWorkshopForGameServer          func(steamUGC uintptr, unWorkshopDepotID DepotId, pszFolder string) bool
	iSteamUGC_SuspendDownloads                    func(steamUGC uintptr, bSuspend bool)
	iSteamUGC_StartPlaytimeTracking               func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall
	iSteamUGC_StopPlaytimeTracking                func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall
	iSteamUGC_StopPlaytimeTrackingForAllItems     func(steamUGC uintptr) SteamAPICall
	iSteamUGC_AddDependency                       func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_RemoveDependency                    func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_AddAppDependency                    func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId) SteamAPICall
	iSteamUGC_RemoveAppDependency                 func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId) SteamAPICall
	iSteamUGC_GetAppDependencies                  func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_DeleteItem                          func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall
	iSteamUGC_ShowWorkshopEULA                    func(steamUGC uintptr) bool
	iSteamUGC_GetWorkshopEULAStatus               func(steamUGC uintptr) SteamAPICall
	iSteamUGC_GetUserContentDescriptorPreferences func(steamUGC uintptr, pvecDescriptors *[]EUGCContentDescriptorID, cMaxEntries uint32) uint32
)

var steamUGC uintptr

func ugc() uintptr {
	if steamUGC == 0 {
		steamUGC = steamUGC_init()
		return steamUGC
	}
	return steamUGC
}

func CreateQueryUserUGCRequest(accountID AccountID, listType EUserUGCList, matchingUGCType EUGCMatchingUGCType, sortOrder EUserUGCListSortOrder, creatorAppID AppId, consumerAppID AppId, page uint32) UGCQueryHandle {
	return iSteamUGC_CreateQueryUserUGCRequest(ugc(), accountID, listType, matchingUGCType, sortOrder, creatorAppID, consumerAppID, page)
}

func CreateQueryAllUGCRequest(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId, consumerAppID AppId, page uint32) UGCQueryHandle {
	return iSteamUGC_CreateQueryAllUGCRequestPage(ugc(), queryType, matchingeMatchingUGCTypeFileType, creatorAppID, consumerAppID, page)
}

func CreateQueryAllUGCRequestCursor(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId, consumerAppID AppId, cursor string) UGCQueryHandle {
	return iSteamUGC_CreateQueryAllUGCRequestCursor(ugc(), queryType, matchingeMatchingUGCTypeFileType, creatorAppID, consumerAppID, cursor)
}

func CreateQueryUGCDetailsRequest(publishedFileIDList []PublishedFileId) UGCQueryHandle {
	return iSteamUGC_CreateQueryUGCDetailsRequest(ugc(), publishedFileIDList, uint32(len(publishedFileIDList)))
}

func SendQueryUGCRequest(handle UGCQueryHandle) CallResult[SteamUGCQueryCompleted] {
	apiHandle := iSteamUGC_SendQueryUGCRequest(ugc(), handle)
	return CallResult[SteamUGCQueryCompleted]{Handle: apiHandle}
}

func GetQueryUGCResult(handle UGCQueryHandle, index uint32) (details SteamUGCDetails, success bool) {
	success = iSteamUGC_GetQueryUGCResult(ugc(), handle, index, &details)
	return details, success
}

func GetQueryUGCNumTags(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumTags(ugc(), handle, index)
}

func GetQueryUGCTag(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCTag(ugc(), handle, index, indexTag, &pchValue, valueSize)
	return pchValue, result
}

func GetQueryUGCTagDisplayName(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCTagDisplayName(ugc(), handle, index, indexTag, &pchValue, valueSize)
	return pchValue, result
}

func GetQueryUGCPreviewURL(handle UGCQueryHandle, index uint32, URLSize uint32) ([]byte, bool) {
	var pchURL []byte = make([]byte, 0, URLSize)
	result := iSteamUGC_GetQueryUGCPreviewURL(ugc(), handle, index, &pchURL, URLSize)
	return pchURL, result
}

func GetQueryUGCMetadata(handle UGCQueryHandle, index uint32, metadataSize uint32) ([]byte, bool) {
	var pchMetadata []byte = make([]byte, 0, metadataSize)
	result := iSteamUGC_GetQueryUGCMetadata(ugc(), handle, index, &pchMetadata, metadataSize)
	return pchMetadata, result
}

func GetQueryUGCChildren(handle UGCQueryHandle, index uint32, maxEntries uint32) ([]PublishedFileId, bool) {
	pvecPublishedFileID := make([]PublishedFileId, 0, maxEntries)
	success := iSteamUGC_GetQueryUGCChildren(ugc(), handle, index, &pvecPublishedFileID, maxEntries)
	return pvecPublishedFileID, success
}

func GetQueryUGCStatistic(handle UGCQueryHandle, index uint32, statType EItemStatistic) (statValue uint64, success bool) {
	success = iSteamUGC_GetQueryUGCStatistic(ugc(), handle, index, statType, &statValue)
	return statValue, success
}

func GetQueryUGCNumAdditionalPreviews(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumAdditionalPreviews(ugc(), handle, index)
}

func GetQueryUGCAdditionalPreview(handle UGCQueryHandle, index uint32, previewIndex uint32, URLSize uint32, originalFileNameSize uint32) ([]byte, []byte, EItemPreviewType, bool) {
	var pchURLOrVideoID []byte = make([]byte, 0, URLSize)
	var pchOriginalFileName []byte = make([]byte, 0, originalFileNameSize)
	var pPreviewType EItemPreviewType
	result := iSteamUGC_GetQueryUGCAdditionalPreview(ugc(), handle, index, previewIndex, &pchURLOrVideoID, URLSize, &pchOriginalFileName, originalFileNameSize, &pPreviewType)
	return pchURLOrVideoID, pchOriginalFileName, pPreviewType, result
}

func GetQueryUGCNumKeyValueTags(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetQueryUGCNumKeyValueTags(ugc(), handle, index)
}

func GetQueryUGCKeyValueTag(handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, keySize uint32, valueSize uint32) ([]byte, []byte, bool) {
	var pchKey []byte = make([]byte, 0, keySize)
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryUGCKeyValueTag(ugc(), handle, index, keyValueTagIndex, &pchKey, keySize, &pchValue, valueSize)
	return pchKey, pchValue, result
}

func GetQueryFirstUGCKeyValueTag(handle UGCQueryHandle, index uint32, key string, valueSize uint32) ([]byte, bool) {
	var pchValue []byte = make([]byte, 0, valueSize)
	result := iSteamUGC_GetQueryFirstUGCKeyValueTag(ugc(), handle, index, key, &pchValue, valueSize)
	return pchValue, result
}

func GetNumSupportedGameVersions(handle UGCQueryHandle, index uint32) uint32 {
	return iSteamUGC_GetNumSupportedGameVersions(ugc(), handle, index)
}

func GetSupportedGameVersionData(handle UGCQueryHandle, index uint32, versionIndex uint32, gameBranchSize uint32) ([]byte, []byte, bool) {
	var pchGameBranchMin []byte = make([]byte, 0, gameBranchSize)
	var pchGameBranchMax []byte = make([]byte, 0, gameBranchSize)
	result := iSteamUGC_GetSupportedGameVersionData(ugc(), handle, index, versionIndex, &pchGameBranchMin, &pchGameBranchMax, gameBranchSize)
	return pchGameBranchMin, pchGameBranchMax, result
}

func GetQueryUGCContentDescriptors(handle UGCQueryHandle, index uint32, maxEntries uint32) (descriptorsList []EUGCContentDescriptorID) {
	descriptorsList = make([]EUGCContentDescriptorID, 0, maxEntries)
	result := iSteamUGC_GetQueryUGCContentDescriptors(ugc(), handle, index, &descriptorsList, maxEntries)
	return descriptorsList[:result]
}

func ReleaseQueryUGCRequest(handle UGCQueryHandle) bool {
	return iSteamUGC_ReleaseQueryUGCRequest(ugc(), handle)
}

func AddRequiredTag(handle UGCQueryHandle, tagName string) bool {
	return iSteamUGC_AddRequiredTag(ugc(), handle, tagName)
}

func AddRequiredTagGroup(handle UGCQueryHandle, tagGroups []SteamParamStringArray) bool {
	return iSteamUGC_AddRequiredTagGroup(ugc(), handle, tagGroups)
}

func AddExcludedTag(handle UGCQueryHandle, tagName string) bool {
	return iSteamUGC_AddExcludedTag(ugc(), handle, tagName)
}

func SetReturnOnlyIDs(handle UGCQueryHandle, returnOnlyIDs bool) bool {
	return iSteamUGC_SetReturnOnlyIDs(ugc(), handle, returnOnlyIDs)
}

func SetReturnKeyValueTags(handle UGCQueryHandle, returnKeyValueTags bool) bool {
	return iSteamUGC_SetReturnKeyValueTags(ugc(), handle, returnKeyValueTags)
}

func SetReturnLongDescription(handle UGCQueryHandle, returnLongDescription bool) bool {
	return iSteamUGC_SetReturnLongDescription(ugc(), handle, returnLongDescription)
}

func SetReturnMetadata(handle UGCQueryHandle, returnMetadata bool) bool {
	return iSteamUGC_SetReturnMetadata(ugc(), handle, returnMetadata)
}

func SetReturnChildren(handle UGCQueryHandle, returnChildren bool) bool {
	return iSteamUGC_SetReturnChildren(ugc(), handle, returnChildren)
}

func SetReturnAdditionalPreviews(handle UGCQueryHandle, returnAdditionalPreviews bool) bool {
	return iSteamUGC_SetReturnAdditionalPreviews(ugc(), handle, returnAdditionalPreviews)
}

func SetReturnTotalOnly(handle UGCQueryHandle, returnTotalOnly bool) bool {
	return iSteamUGC_SetReturnTotalOnly(ugc(), handle, returnTotalOnly)
}

func SetReturnPlaytimeStats(handle UGCQueryHandle, days uint32) bool {
	return iSteamUGC_SetReturnPlaytimeStats(ugc(), handle, days)
}

func SetLanguage(handle UGCQueryHandle, language string) bool {
	return iSteamUGC_SetLanguage(ugc(), handle, language)
}

func SetAllowCachedResponse(handle UGCQueryHandle, maxAgeSeconds uint32) bool {
	return iSteamUGC_SetAllowCachedResponse(ugc(), handle, maxAgeSeconds)
}

func SetAdminQuery(handle UGCUpdateHandle, adminQuery bool) bool {
	return iSteamUGC_SetAdminQuery(ugc(), handle, adminQuery)
}

func SetCloudFileNameFilter(handle UGCQueryHandle, matchCloudFileName string) bool {
	return iSteamUGC_SetCloudFileNameFilter(ugc(), handle, matchCloudFileName)
}

func SetMatchAnyTag(handle UGCQueryHandle, matchAnyTag bool) bool {
	return iSteamUGC_SetMatchAnyTag(ugc(), handle, matchAnyTag)
}

func SetSearchText(handle UGCQueryHandle, searchText string) bool {
	return iSteamUGC_SetSearchText(ugc(), handle, searchText)
}

func SetRankedByTrendDays(handle UGCQueryHandle, days uint32) bool {
	return iSteamUGC_SetRankedByTrendDays(ugc(), handle, days)
}

func SetTimeCreatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool {
	return iSteamUGC_SetTimeCreatedDateRange(ugc(), handle, startTime, endTime)
}

func SetTimeUpdatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool {
	return iSteamUGC_SetTimeUpdatedDateRange(ugc(), handle, startTime, endTime)
}

func AddRequiredKeyValueTag(handle UGCQueryHandle, key string, value string) bool {
	return iSteamUGC_AddRequiredKeyValueTag(ugc(), handle, key, value)
}

func CreateItem(consumerAppId AppId, fileType EWorkshopFileType) CallResult[CreateItemResult] {
	handle := iSteamUGC_CreateItem(ugc(), consumerAppId, fileType)
	return CallResult[CreateItemResult]{Handle: handle}
}

func StartItemUpdate(consumerAppId AppId, publishedFileID PublishedFileId) UGCUpdateHandle {
	return iSteamUGC_StartItemUpdate(ugc(), consumerAppId, publishedFileID)
}

func SetItemTitle(handle UGCUpdateHandle, title string) bool {
	return iSteamUGC_SetItemTitle(ugc(), handle, title)
}

func SetItemDescription(handle UGCUpdateHandle, description string) bool {
	return iSteamUGC_SetItemDescription(ugc(), handle, description)
}

func SetItemUpdateLanguage(handle UGCUpdateHandle, language string) bool {
	return iSteamUGC_SetItemUpdateLanguage(ugc(), handle, language)
}

func SetItemMetadata(handle UGCUpdateHandle, metaData string) bool {
	return iSteamUGC_SetItemMetadata(ugc(), handle, metaData)
}

func SetItemVisibility(handle UGCUpdateHandle, visibility ERemoteStoragePublishedFileVisibility) bool {
	return iSteamUGC_SetItemVisibility(ugc(), handle, visibility)
}

func SetItemTags(updateHandle UGCUpdateHandle, tags []SteamParamStringArray, allowAdminTags bool) bool {
	return iSteamUGC_SetItemTags(ugc(), updateHandle, tags, allowAdminTags)
}

func SetItemContent(handle UGCUpdateHandle, contentFolder string) bool {
	return iSteamUGC_SetItemContent(ugc(), handle, contentFolder)
}

func SetItemPreview(handle UGCUpdateHandle, previewFile string) bool {
	return iSteamUGC_SetItemPreview(ugc(), handle, previewFile)
}

func SetAllowLegacyUpload(handle UGCUpdateHandle, allowLegacyUpload bool) bool {
	return iSteamUGC_SetAllowLegacyUpload(ugc(), handle, allowLegacyUpload)
}

func RemoveAllItemKeyValueTags(handle UGCUpdateHandle) bool {
	return iSteamUGC_RemoveAllItemKeyValueTags(ugc(), handle)
}

func RemoveItemKeyValueTags(handle UGCUpdateHandle, key string) bool {
	return iSteamUGC_RemoveItemKeyValueTags(ugc(), handle, key)
}

func AddItemKeyValueTag(handle UGCUpdateHandle, key string, value string) bool {
	return iSteamUGC_AddItemKeyValueTag(ugc(), handle, key, value)
}

func AddItemPreviewFile(handle UGCUpdateHandle, previewFile string, Type EItemPreviewType) bool {
	return iSteamUGC_AddItemPreviewFile(ugc(), handle, previewFile, Type)
}

func AddItemPreviewVideo(handle UGCUpdateHandle, videoID string) bool {
	return iSteamUGC_AddItemPreviewVideo(ugc(), handle, videoID)
}

func UpdateItemPreviewFile(handle UGCUpdateHandle, index uint32, previewFile string) bool {
	return iSteamUGC_UpdateItemPreviewFile(ugc(), handle, index, previewFile)
}

func UpdateItemPreviewVideo(handle UGCUpdateHandle, index uint32, videoID string) bool {
	return iSteamUGC_UpdateItemPreviewVideo(ugc(), handle, index, videoID)
}

func RemoveItemPreview(handle UGCUpdateHandle, index uint32) bool {
	return iSteamUGC_RemoveItemPreview(ugc(), handle, index)
}

func AddContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool {
	return iSteamUGC_AddContentDescriptor(ugc(), handle, descID)
}

func RemoveContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool {
	return iSteamUGC_RemoveContentDescriptor(ugc(), handle, descID)
}

func SetRequiredGameVersions(handle UGCUpdateHandle, gameBranchMin string, gameBranchMax string) bool {
	return iSteamUGC_SetRequiredGameVersions(ugc(), handle, gameBranchMin, gameBranchMax)
}

func SubmitItemUpdate(handle UGCUpdateHandle, changeNote string) CallResult[SubmitItemUpdateResult] {
	apiHandle := iSteamUGC_SubmitItemUpdate(ugc(), handle, changeNote)
	return CallResult[SubmitItemUpdateResult]{Handle: apiHandle}
}

func GetItemUpdateProgress(handle UGCUpdateHandle) (bytesProcessed uint64, bytesTotal uint64, status EItemUpdateStatus) {
	status = iSteamUGC_GetItemUpdateProgress(ugc(), handle, &bytesProcessed, &bytesTotal)
	return bytesProcessed, bytesTotal, status
}

func SetUserItemVote(publishedFileID PublishedFileId, voteUp bool) CallResult[SetUserItemVoteResult] {
	handle := iSteamUGC_SetUserItemVote(ugc(), publishedFileID, voteUp)
	return CallResult[SetUserItemVoteResult]{Handle: handle}
}

func GetUserItemVote(publishedFileID PublishedFileId) CallResult[GetUserItemVoteResult] {
	handle := iSteamUGC_GetUserItemVote(ugc(), publishedFileID)
	return CallResult[GetUserItemVoteResult]{Handle: handle}
}

func AddItemToFavorites(appId AppId, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged] {
	handle := iSteamUGC_AddItemToFavorites(ugc(), appId, publishedFileID)
	return CallResult[UserFavoriteItemsListChanged]{Handle: handle}
}

func RemoveItemFromFavorites(appId AppId, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged] {
	handle := iSteamUGC_RemoveItemFromFavorites(ugc(), appId, publishedFileID)
	return CallResult[UserFavoriteItemsListChanged]{Handle: handle}
}

func SubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageSubscribePublishedFileResult] {
	handle := iSteamUGC_SubscribeItem(ugc(), publishedFileID)
	return CallResult[RemoteStorageSubscribePublishedFileResult]{Handle: handle}
}

func UnsubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageUnsubscribePublishedFileResult] {
	handle := iSteamUGC_UnsubscribeItem(ugc(), publishedFileID)
	return CallResult[RemoteStorageUnsubscribePublishedFileResult]{Handle: handle}
}

func GetNumSubscribedItems() uint32 {
	return iSteamUGC_GetNumSubscribedItems(ugc())
}

func GetSubscribedItems(maxEntries uint32) []PublishedFileId {
	pvecPublishedFileID := make([]PublishedFileId, 0, maxEntries)
	total := iSteamUGC_GetSubscribedItems(ugc(), &pvecPublishedFileID, maxEntries)
	return pvecPublishedFileID[:total]
}

func GetItemState(publishedFileID PublishedFileId) uint32 {
	return iSteamUGC_GetItemState(ugc(), publishedFileID)
}

func GetItemInstallInfo(publishedFileID PublishedFileId, folderSize uint32) (uint64, uint32, []byte, bool) {
	var pchFolder []byte = make([]byte, 0, folderSize)
	var punSizeOnDisk uint64
	var punTimeStamp uint32
	result := iSteamUGC_GetItemInstallInfo(ugc(), publishedFileID, &punSizeOnDisk, &pchFolder, folderSize, &punTimeStamp)
	return punSizeOnDisk, punTimeStamp, pchFolder, result
}

func GetItemDownloadInfo(publishedFileID PublishedFileId) (bytesDownloaded uint64, bytesTotal uint64, success bool) {
	success = iSteamUGC_GetItemDownloadInfo(ugc(), publishedFileID, &bytesDownloaded, &bytesTotal)
	return bytesDownloaded, bytesDownloaded, success
}

func DownloadItem(publishedFileID PublishedFileId, highPriority bool) bool {
	return iSteamUGC_DownloadItem(ugc(), publishedFileID, highPriority)
}

func BInitWorkshopForGameServer(workshopDepotID DepotId, folder string) bool {
	return iSteamUGC_BInitWorkshopForGameServer(ugc(), workshopDepotID, folder)
}

func SuspendDownloads(suspend bool) {
	iSteamUGC_SuspendDownloads(ugc(), suspend)
}

func StartPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StartPlaytimeTrackingResult] {
	handle := iSteamUGC_StartPlaytimeTracking(ugc(), publishedFileIDList, uint32(len(publishedFileIDList)))
	return CallResult[StartPlaytimeTrackingResult]{Handle: handle}
}

func StopPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StopPlaytimeTrackingResult] {
	handle := iSteamUGC_StopPlaytimeTracking(ugc(), publishedFileIDList, uint32(len(publishedFileIDList)))
	return CallResult[StopPlaytimeTrackingResult]{Handle: handle}
}

func StopPlaytimeTrackingForAllItems() CallResult[StopPlaytimeTrackingResult] {
	handle := iSteamUGC_StopPlaytimeTrackingForAllItems(ugc())
	return CallResult[StopPlaytimeTrackingResult]{Handle: handle}
}

func AddDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[AddUGCDependencyResult] {
	handle := iSteamUGC_AddDependency(ugc(), parentPublishedFileID, childPublishedFileID)
	return CallResult[AddUGCDependencyResult]{Handle: handle}
}

func RemoveDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[RemoveUGCDependencyResult] {
	handle := iSteamUGC_RemoveDependency(ugc(), parentPublishedFileID, childPublishedFileID)
	return CallResult[RemoveUGCDependencyResult]{Handle: handle}
}

func AddAppDependency(publishedFileID PublishedFileId, appID AppId) CallResult[AddAppDependencyResult] {
	handle := iSteamUGC_AddAppDependency(ugc(), publishedFileID, appID)
	return CallResult[AddAppDependencyResult]{Handle: handle}
}

func RemoveAppDependency(publishedFileID PublishedFileId, appID AppId) CallResult[RemoveAppDependencyResult] {
	handle := iSteamUGC_RemoveAppDependency(ugc(), publishedFileID, appID)
	return CallResult[RemoveAppDependencyResult]{Handle: handle}
}

func GetAppDependencies(publishedFileID PublishedFileId) CallResult[GetAppDependenciesResult] {
	handle := iSteamUGC_GetAppDependencies(ugc(), publishedFileID)
	return CallResult[GetAppDependenciesResult]{Handle: handle}
}

func DeleteItem(publishedFileID PublishedFileId) CallResult[DeleteItemResult] {
	handle := iSteamUGC_DeleteItem(ugc(), publishedFileID)
	return CallResult[DeleteItemResult]{Handle: handle}
}

func ShowWorkshopEULA() bool {
	return iSteamUGC_ShowWorkshopEULA(ugc())
}

func GetWorkshopEULAStatus() CallResult[WorkshopEULAStatus] {
	handle := iSteamUGC_GetWorkshopEULAStatus(ugc())
	return CallResult[WorkshopEULAStatus]{Handle: handle}
}

func GetUserContentDescriptorPreferences(maxEntries uint32) []EUGCContentDescriptorID {
	var pvecDescriptors []EUGCContentDescriptorID = make([]EUGCContentDescriptorID, 0, maxEntries)
	total := iSteamUGC_GetUserContentDescriptorPreferences(ugc(), &pvecDescriptors, maxEntries)
	return pvecDescriptors[:total]
}
