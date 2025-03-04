package remotestoragetypes

import (
	. "github.com/assemblaj/purego-steamworks"
)

type PublishedFileId uint64

type SteamParamStringArray struct {
	Strings    []string
	NumStrings int32
}
type ERemoteStoragePublishedFileVisibility int32

const (
	ERemoteStoragePublishedFileVisibilityPublic      ERemoteStoragePublishedFileVisibility = 0
	ERemoteStoragePublishedFileVisibilityFriendsOnly ERemoteStoragePublishedFileVisibility = 1
	ERemoteStoragePublishedFileVisibilityPrivate     ERemoteStoragePublishedFileVisibility = 2
	ERemoteStoragePublishedFileVisibilityUnlisted    ERemoteStoragePublishedFileVisibility = 3
)

type EWorkshopFileType int32

const (
	EWorkshopFileTypeFirst                  EWorkshopFileType = 0
	EWorkshopFileTypeCommunity              EWorkshopFileType = 0
	EWorkshopFileTypeMicrotransaction       EWorkshopFileType = 1
	EWorkshopFileTypeCollection             EWorkshopFileType = 2
	EWorkshopFileTypeArt                    EWorkshopFileType = 3
	EWorkshopFileTypeVideo                  EWorkshopFileType = 4
	EWorkshopFileTypeScreenshot             EWorkshopFileType = 5
	EWorkshopFileTypeGame                   EWorkshopFileType = 6
	EWorkshopFileTypeSoftware               EWorkshopFileType = 7
	EWorkshopFileTypeConcept                EWorkshopFileType = 8
	EWorkshopFileTypeWebGuide               EWorkshopFileType = 9
	EWorkshopFileTypeIntegratedGuide        EWorkshopFileType = 10
	EWorkshopFileTypeMerch                  EWorkshopFileType = 11
	EWorkshopFileTypeControllerBinding      EWorkshopFileType = 12
	EWorkshopFileTypeSteamworksAccessInvite EWorkshopFileType = 13
	EWorkshopFileTypeSteamVideo             EWorkshopFileType = 14
	EWorkshopFileTypeGameManagedItem        EWorkshopFileType = 15
	EWorkshopFileTypeClip                   EWorkshopFileType = 16
	EWorkshopFileTypeMax                    EWorkshopFileType = 17
)

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
	return CallbackString(cb)
}

func (cb RemoteStorageSubscribePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageSubscribePublishedFileResultID
}

func (cb RemoteStorageSubscribePublishedFileResult) Name() string {
	return "RemoteStorageSubscribePublishedFileResult"
}

func (cb RemoteStorageSubscribePublishedFileResult) String() string {
	return CallbackString(cb)
}
