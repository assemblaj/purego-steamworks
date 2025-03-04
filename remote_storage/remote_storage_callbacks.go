package remotestorage

import (
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"

	. "github.com/assemblaj/purego-steamworks"
)

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
	return CallbackString(cb)
}

func (cb RemoteStoragePublishFileResult) CallbackID() SteamCallbackID {
	return RemoteStoragePublishFileResultID
}

func (cb RemoteStoragePublishFileResult) Name() string {
	return "RemoteStoragePublishFileResult"
}

func (cb RemoteStoragePublishFileResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageDeletePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageDeletePublishedFileResultID
}

func (cb RemoteStorageDeletePublishedFileResult) Name() string {
	return "RemoteStorageDeletePublishedFileResult"
}

func (cb RemoteStorageDeletePublishedFileResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserPublishedFilesResultID
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) Name() string {
	return "RemoteStorageEnumerateUserPublishedFilesResult"
}

func (cb RemoteStorageEnumerateUserPublishedFilesResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserSubscribedFilesResultID
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) Name() string {
	return "RemoteStorageEnumerateUserSubscribedFilesResult"
}

func (cb RemoteStorageEnumerateUserSubscribedFilesResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageUpdatePublishedFileResult) CallbackID() SteamCallbackID {
	return RemoteStorageUpdatePublishedFileResultID
}

func (cb RemoteStorageUpdatePublishedFileResult) Name() string {
	return "RemoteStorageUpdatePublishedFileResult"
}

func (cb RemoteStorageUpdatePublishedFileResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageDownloadUGCResult) CallbackID() SteamCallbackID {
	return RemoteStorageDownloadUGCResultID
}

func (cb RemoteStorageDownloadUGCResult) Name() string {
	return "RemoteStorageDownloadUGCResult"
}

func (cb RemoteStorageDownloadUGCResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageGetPublishedFileDetailsResult) CallbackID() SteamCallbackID {
	return RemoteStorageGetPublishedFileDetailsResultID
}

func (cb RemoteStorageGetPublishedFileDetailsResult) Name() string {
	return "RemoteStorageGetPublishedFileDetailsResult"
}

func (cb RemoteStorageGetPublishedFileDetailsResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateWorkshopFilesResultID
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) Name() string {
	return "RemoteStorageEnumerateWorkshopFilesResult"
}

func (cb RemoteStorageEnumerateWorkshopFilesResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) CallbackID() SteamCallbackID {
	return RemoteStorageGetPublishedItemVoteDetailsResultID
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) Name() string {
	return "RemoteStorageGetPublishedItemVoteDetailsResult"
}

func (cb RemoteStorageGetPublishedItemVoteDetailsResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStoragePublishedFileSubscribed) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileSubscribedID
}

func (cb RemoteStoragePublishedFileSubscribed) Name() string {
	return "RemoteStoragePublishedFileSubscribed"
}

func (cb RemoteStoragePublishedFileSubscribed) String() string {
	return CallbackString(cb)
}

func (cb RemoteStoragePublishedFileUnsubscribed) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileUnsubscribedID
}

func (cb RemoteStoragePublishedFileUnsubscribed) Name() string {
	return "RemoteStoragePublishedFileUnsubscribed"
}

func (cb RemoteStoragePublishedFileUnsubscribed) String() string {
	return CallbackString(cb)
}

func (cb RemoteStoragePublishedFileDeleted) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileDeletedID
}

func (cb RemoteStoragePublishedFileDeleted) Name() string {
	return "RemoteStoragePublishedFileDeleted"
}

func (cb RemoteStoragePublishedFileDeleted) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) CallbackID() SteamCallbackID {
	return RemoteStorageUpdateUserPublishedItemVoteResultID
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) Name() string {
	return "RemoteStorageUpdateUserPublishedItemVoteResult"
}

func (cb RemoteStorageUpdateUserPublishedItemVoteResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageUserVoteDetails) CallbackID() SteamCallbackID {
	return RemoteStorageUserVoteDetailsID
}

func (cb RemoteStorageUserVoteDetails) Name() string {
	return "RemoteStorageUserVoteDetails"
}

func (cb RemoteStorageUserVoteDetails) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumerateUserSharedWorkshopFilesResultID
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) Name() string {
	return "RemoteStorageEnumerateUserSharedWorkshopFilesResult"
}

func (cb RemoteStorageEnumerateUserSharedWorkshopFilesResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageSetUserPublishedFileActionResult) CallbackID() SteamCallbackID {
	return RemoteStorageSetUserPublishedFileActionResultID
}

func (cb RemoteStorageSetUserPublishedFileActionResult) Name() string {
	return "RemoteStorageSetUserPublishedFileActionResult"
}

func (cb RemoteStorageSetUserPublishedFileActionResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) CallbackID() SteamCallbackID {
	return RemoteStorageEnumeratePublishedFilesByUserActionResultID
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) Name() string {
	return "RemoteStorageEnumeratePublishedFilesByUserActionResult"
}

func (cb RemoteStorageEnumeratePublishedFilesByUserActionResult) String() string {
	return CallbackString(cb)
}

func (cb RemoteStoragePublishFileProgress) CallbackID() SteamCallbackID {
	return RemoteStoragePublishFileProgressID
}

func (cb RemoteStoragePublishFileProgress) Name() string {
	return "RemoteStoragePublishFileProgress"
}

func (cb RemoteStoragePublishFileProgress) String() string {
	return CallbackString(cb)
}

func (cb RemoteStoragePublishedFileUpdated) CallbackID() SteamCallbackID {
	return RemoteStoragePublishedFileUpdatedID
}

func (cb RemoteStoragePublishedFileUpdated) Name() string {
	return "RemoteStoragePublishedFileUpdated"
}

func (cb RemoteStoragePublishedFileUpdated) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageFileWriteAsyncComplete) CallbackID() SteamCallbackID {
	return RemoteStorageFileWriteAsyncCompleteID
}

func (cb RemoteStorageFileWriteAsyncComplete) Name() string {
	return "RemoteStorageFileWriteAsyncComplete"
}

func (cb RemoteStorageFileWriteAsyncComplete) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageFileReadAsyncComplete) CallbackID() SteamCallbackID {
	return RemoteStorageFileReadAsyncCompleteID
}

func (cb RemoteStorageFileReadAsyncComplete) Name() string {
	return "RemoteStorageFileReadAsyncComplete"
}

func (cb RemoteStorageFileReadAsyncComplete) String() string {
	return CallbackString(cb)
}

func (cb RemoteStorageLocalFileChange) CallbackID() SteamCallbackID {
	return RemoteStorageLocalFileChangeID
}

func (cb RemoteStorageLocalFileChange) Name() string {
	return "RemoteStorageLocalFileChange"
}

func (cb RemoteStorageLocalFileChange) String() string {
	return CallbackString(cb)
}
