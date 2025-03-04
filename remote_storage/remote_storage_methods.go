package remotestorage

import (
	common "github.com/assemblaj/purego-steamworks"
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"
)

const (
	MaxCloudFileChunkSize uint32 = 100 * 1024 * 1024

	PublishedFileIdInvalid           PublishedFileId           = 0
	UGCHandleInvalid                 common.UGCHandle          = 0xffffffffffffffff
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
	EWorkshopVoteUnvoted EWorkshopVote = 0
	EWorkshopVoteFor     EWorkshopVote = 1
	EWorkshopVoteAgainst EWorkshopVote = 2
	EWorkshopVoteLater   EWorkshopVote = 3
)

type ERemoteStorageFilePathType int32

const (
	ERemoteStorageFilePathTypeInvalid     ERemoteStorageFilePathType = 0
	ERemoteStorageFilePathTypeAbsolute    ERemoteStorageFilePathType = 1
	ERemoteStorageFilePathTypeAPIFilename ERemoteStorageFilePathType = 2
)

type EUGCReadAction int32

const (
	EUGCReadContinueReadingUntilFinished EUGCReadAction = 0
	EUGCReadContinueReading              EUGCReadAction = 1
	EUGCReadClose                        EUGCReadAction = 2
)

type ERemoteStorageLocalFileChange int32

const (
	ERemoteStorageLocalFileChangeInvalid     ERemoteStorageLocalFileChange = 0
	ERemoteStorageLocalFileChangeFileUpdated ERemoteStorageLocalFileChange = 1
	ERemoteStorageLocalFileChangeFileDeleted ERemoteStorageLocalFileChange = 2
)

type EWorkshopVideoProvider int32

const (
	EWorkshopVideoProviderNone    EWorkshopVideoProvider = 0
	EWorkshopVideoProviderYoutube EWorkshopVideoProvider = 1
)

type EWorkshopFileAction int32

const (
	EWorkshopFileActionPlayed    EWorkshopFileAction = 0
	EWorkshopFileActionCompleted EWorkshopFileAction = 1
)

type EWorkshopEnumerationType int32

const (
	EWorkshopEnumerationTypeRankedByVote            EWorkshopEnumerationType = 0
	EWorkshopEnumerationTypeRecent                  EWorkshopEnumerationType = 1
	EWorkshopEnumerationTypeTrending                EWorkshopEnumerationType = 2
	EWorkshopEnumerationTypeFavoritesOfFriends      EWorkshopEnumerationType = 3
	EWorkshopEnumerationTypeVotedByFriends          EWorkshopEnumerationType = 4
	EWorkshopEnumerationTypeContentByFriends        EWorkshopEnumerationType = 5
	EWorkshopEnumerationTypeRecentFromFollowedUsers EWorkshopEnumerationType = 6
)

type ERemoteStoragePlatform int32

const (
	ERemoteStoragePlatformNone    ERemoteStoragePlatform = 0
	ERemoteStoragePlatformWindows ERemoteStoragePlatform = 1
	ERemoteStoragePlatformOSX     ERemoteStoragePlatform = 2
	ERemoteStoragePlatformPS3     ERemoteStoragePlatform = 4
	ERemoteStoragePlatformLinux   ERemoteStoragePlatform = 8
	ERemoteStoragePlatformSwitch  ERemoteStoragePlatform = 16
	ERemoteStoragePlatformAndroid ERemoteStoragePlatform = 32
	ERemoteStoragePlatformIOS     ERemoteStoragePlatform = 64
	ERemoteStoragePlatformAll     ERemoteStoragePlatform = -1
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
	iSteamRemoteStorage_FileRead                                func(steamRemoteStorage uintptr, pchFile string, pvData *[]byte, cubDataToRead int32) int32
	iSteamRemoteStorage_FileWriteAsync                          func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubData uint32) common.SteamAPICall
	iSteamRemoteStorage_FileReadAsync                           func(steamRemoteStorage uintptr, pchFile string, nOffset uint32, cubToRead uint32) common.SteamAPICall
	iSteamRemoteStorage_FileReadAsyncComplete                   func(steamRemoteStorage uintptr, hReadCall common.SteamAPICall, pvBuffer *[]byte, cubToRead uint32) bool
	iSteamRemoteStorage_FileForget                              func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_FileDelete                              func(steamRemoteStorage uintptr, pchFile string) bool
	iSteamRemoteStorage_FileShare                               func(steamRemoteStorage uintptr, pchFile string) common.SteamAPICall
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
	iSteamRemoteStorage_UGCDownload                             func(steamRemoteStorage uintptr, hContent common.UGCHandle, unPriority uint32) common.SteamAPICall
	iSteamRemoteStorage_GetUGCDownloadProgress                  func(steamRemoteStorage uintptr, hContent common.UGCHandle, pnBytesDownloaded *int32, pnBytesExpected *int32) bool
	iSteamRemoteStorage_GetUGCDetails                           func(steamRemoteStorage uintptr, hContent common.UGCHandle, pnAppID *common.AppId, ppchName *[]byte, pnFileSizeInBytes *int32, pSteamIDOwner *common.CSteamID) bool
	iSteamRemoteStorage_UGCRead                                 func(steamRemoteStorage uintptr, hContent common.UGCHandle, pvData *[]byte, cubDataToRead int32, cOffset uint32, eAction EUGCReadAction) int32
	iSteamRemoteStorage_GetCachedUGCCount                       func(steamRemoteStorage uintptr) int32
	iSteamRemoteStorage_GetCachedUGCHandle                      func(steamRemoteStorage uintptr, iCachedContent int32) common.UGCHandle
	iSteamRemoteStorage_PublishWorkshopFile                     func(steamRemoteStorage uintptr, pchFile string, pchPreviewFile string, nConsumerAppId common.AppId, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr, eWorkshopFileType EWorkshopFileType) common.SteamAPICall
	iSteamRemoteStorage_CreatePublishedFileUpdateRequest        func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) PublishedFileUpdateHandle
	iSteamRemoteStorage_UpdatePublishedFileFile                 func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchFile string) bool
	iSteamRemoteStorage_UpdatePublishedFilePreviewFile          func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchPreviewFile string) bool
	iSteamRemoteStorage_UpdatePublishedFileTitle                func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchTitle string) bool
	iSteamRemoteStorage_UpdatePublishedFileDescription          func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchDescription string) bool
	iSteamRemoteStorage_UpdatePublishedFileVisibility           func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool
	iSteamRemoteStorage_UpdatePublishedFileTags                 func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pTags uintptr) bool
	iSteamRemoteStorage_CommitPublishedFileUpdate               func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle) common.SteamAPICall
	iSteamRemoteStorage_GetPublishedFileDetails                 func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, unMaxSecondsOld uint32) common.SteamAPICall
	iSteamRemoteStorage_DeletePublishedFile                     func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) common.SteamAPICall
	iSteamRemoteStorage_EnumerateUserPublishedFiles             func(steamRemoteStorage uintptr, unStartIndex uint32) common.SteamAPICall
	iSteamRemoteStorage_SubscribePublishedFile                  func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) common.SteamAPICall
	iSteamRemoteStorage_EnumerateUserSubscribedFiles            func(steamRemoteStorage uintptr, unStartIndex uint32) common.SteamAPICall
	iSteamRemoteStorage_UnsubscribePublishedFile                func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) common.SteamAPICall
	iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchChangeDescription string) bool
	iSteamRemoteStorage_GetPublishedItemVoteDetails             func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) common.SteamAPICall
	iSteamRemoteStorage_UpdateUserPublishedItemVote             func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, bVoteUp bool) common.SteamAPICall
	iSteamRemoteStorage_GetUserPublishedItemVoteDetails         func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) common.SteamAPICall
	iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles        func(steamRemoteStorage uintptr, steamId common.Uint64SteamID, unStartIndex uint32, pRequiredTags uintptr, pExcludedTags uintptr) common.SteamAPICall
	iSteamRemoteStorage_PublishVideo                            func(steamRemoteStorage uintptr, eVideoProvider EWorkshopVideoProvider, pchVideoAccount string, pchVideoIdentifier string, pchPreviewFile string, nConsumerAppId common.AppId, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr) common.SteamAPICall
	iSteamRemoteStorage_SetUserPublishedFileAction              func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, eAction EWorkshopFileAction) common.SteamAPICall
	iSteamRemoteStorage_EnumeratePublishedFilesByUserAction     func(steamRemoteStorage uintptr, eAction EWorkshopFileAction, unStartIndex uint32) common.SteamAPICall
	iSteamRemoteStorage_EnumeratePublishedWorkshopFiles         func(steamRemoteStorage uintptr, eEnumerationType EWorkshopEnumerationType, unStartIndex uint32, unCount uint32, unDays uint32, pTags *SteamParamStringArray, pUserTags *SteamParamStringArray) common.SteamAPICall
	iSteamRemoteStorage_UGCDownloadToLocation                   func(steamRemoteStorage uintptr, hContent common.UGCHandle, pchLocation string, unPriority uint32) common.SteamAPICall
	iSteamRemoteStorage_GetLocalFileChangeCount                 func(steamRemoteStorage uintptr) int32
	iSteamRemoteStorage_GetLocalFileChange                      func(steamRemoteStorage uintptr, iFile int32, pEChangeType *ERemoteStorageLocalFileChange, pEFilePathType *ERemoteStorageFilePathType) []byte
	iSteamRemoteStorage_BeginFileWriteBatch                     func(steamRemoteStorage uintptr) bool
	iSteamRemoteStorage_EndFileWriteBatch                       func(steamRemoteStorage uintptr) bool
)

var steamRemoteStorage uintptr

func remoteStorage() uintptr {
	if steamRemoteStorage == 0 {
		steamRemoteStorage = steamRemoteStorage_init()
		return steamRemoteStorage
	}
	return steamRemoteStorage
}

func FileWrite(File string, Data []byte) bool {
	return iSteamRemoteStorage_FileWrite(remoteStorage(), File, Data, int32(len(Data)))
}

func FileRead(File string, DataToReadSize int32) (Data []byte) {
	Data = make([]byte, 0, DataToReadSize)
	result := iSteamRemoteStorage_FileRead(remoteStorage(), File, &Data, DataToReadSize)
	return Data[:result]
}

func FileWriteAsync(File string, Data []byte) common.CallResult[RemoteStorageFileWriteAsyncComplete] {
	handle := iSteamRemoteStorage_FileWriteAsync(remoteStorage(), File, Data, uint32(len(Data)))
	return common.CallResult[RemoteStorageFileWriteAsyncComplete]{Handle: handle}
}

func FileReadAsync(File string, Offset uint32, DataToReadSize uint32) common.CallResult[RemoteStorageFileReadAsyncComplete] {
	handle := iSteamRemoteStorage_FileReadAsync(remoteStorage(), File, Offset, DataToReadSize)
	return common.CallResult[RemoteStorageFileReadAsyncComplete]{Handle: handle}
}

func FileReadAsyncComplete(ReadCallHandle common.SteamAPICall, DataToReadSize uint32) (Buffer []byte, success bool) {
	Buffer = make([]byte, 0, DataToReadSize)
	success = iSteamRemoteStorage_FileReadAsyncComplete(remoteStorage(), ReadCallHandle, &Buffer, DataToReadSize)
	return Buffer, success
}

func FileForget(File string) bool {
	return iSteamRemoteStorage_FileForget(remoteStorage(), File)
}

func FileDelete(File string) bool {
	return iSteamRemoteStorage_FileDelete(remoteStorage(), File)
}

func FileShare(File string) common.CallResult[RemoteStorageFileShareResult] {
	handle := iSteamRemoteStorage_FileShare(remoteStorage(), File)
	return common.CallResult[RemoteStorageFileShareResult]{Handle: handle}
}

func SetSyncPlatforms(File string, RemoteStoragePlatform ERemoteStoragePlatform) bool {
	return iSteamRemoteStorage_SetSyncPlatforms(remoteStorage(), File, RemoteStoragePlatform)
}

func FileWriteStreamOpen(File string) UGCFileWriteStreamHandle {
	return iSteamRemoteStorage_FileWriteStreamOpen(remoteStorage(), File)
}

func FileWriteStreamWriteChunk(writeHandle UGCFileWriteStreamHandle, Data []byte) bool {
	return iSteamRemoteStorage_FileWriteStreamWriteChunk(remoteStorage(), writeHandle, Data, int32(len(Data)))
}

func FileWriteStreamClose(writeHandle UGCFileWriteStreamHandle) bool {
	return iSteamRemoteStorage_FileWriteStreamClose(remoteStorage(), writeHandle)
}

func FileWriteStreamCancel(writeHandle UGCFileWriteStreamHandle) bool {
	return iSteamRemoteStorage_FileWriteStreamCancel(remoteStorage(), writeHandle)
}

func FileExists(File string) bool {
	return iSteamRemoteStorage_FileExists(remoteStorage(), File)
}

func FilePersisted(File string) bool {
	return iSteamRemoteStorage_FilePersisted(remoteStorage(), File)
}

func GetFileSize(File string) int32 {
	return iSteamRemoteStorage_GetFileSize(remoteStorage(), File)
}

func GetFileTimestamp(File string) int64 {
	return iSteamRemoteStorage_GetFileTimestamp(remoteStorage(), File)
}

func GetSyncPlatforms(File string) ERemoteStoragePlatform {
	return iSteamRemoteStorage_GetSyncPlatforms(remoteStorage(), File)
}

func GetFileCount() int32 {
	return iSteamRemoteStorage_GetFileCount(remoteStorage())
}

func GetFileNameAndSize(FileIndex int32) (name string, FileSizeInBytes int32) {
	var nameBuf []byte = iSteamRemoteStorage_GetFileNameAndSize(remoteStorage(), FileIndex, &FileSizeInBytes)
	return string(nameBuf), FileSizeInBytes
}

func GetQuota() (TotalBytes uint64, AvailableBytes uint64, success bool) {
	success = iSteamRemoteStorage_GetQuota(remoteStorage(), &TotalBytes, &AvailableBytes)
	return TotalBytes, AvailableBytes, success
}

func IsCloudEnabledForAccount() bool {
	return iSteamRemoteStorage_IsCloudEnabledForAccount(remoteStorage())
}

func IsCloudEnabledForApp() bool {
	return iSteamRemoteStorage_IsCloudEnabledForApp(remoteStorage())
}

func SetCloudEnabledForApp(Enabled bool) {
	iSteamRemoteStorage_SetCloudEnabledForApp(remoteStorage(), Enabled)
}

func UGCDownload(Content common.UGCHandle, Priority uint32) common.CallResult[RemoteStorageDownloadUGCResult] {
	handle := iSteamRemoteStorage_UGCDownload(remoteStorage(), Content, Priority)
	return common.CallResult[RemoteStorageDownloadUGCResult]{Handle: handle}
}

func GetUGCDownloadProgress(Content common.UGCHandle) (BytesDownloaded int32, BytesExpected int32, success bool) {
	success = iSteamRemoteStorage_GetUGCDownloadProgress(remoteStorage(), Content, &BytesDownloaded, &BytesExpected)
	return BytesDownloaded, BytesExpected, success
}

func GetUGCDetails(Content common.UGCHandle) (AppID common.AppId, Name string, FileSizeInBytes int32, ownerSteamID common.CSteamID, success bool) {
	var ppchNameBuf []byte = make([]byte, 0, 4096)
	success = iSteamRemoteStorage_GetUGCDetails(remoteStorage(), Content, &AppID, &ppchNameBuf, &FileSizeInBytes, &ownerSteamID)
	return AppID, string(ppchNameBuf), FileSizeInBytes, ownerSteamID, success
}

func UGCRead(Content common.UGCHandle, DataToReadSize int32, Offset uint32, Action EUGCReadAction) (Data []byte) {
	Data = make([]byte, 0, DataToReadSize)
	result := iSteamRemoteStorage_UGCRead(remoteStorage(), Content, &Data, DataToReadSize, Offset, Action)
	return Data[:result]
}

func GetCachedUGCCount() int32 {
	return iSteamRemoteStorage_GetCachedUGCCount(remoteStorage())
}

func GetCachedUGCHandle(CachedContent int32) common.UGCHandle {
	return iSteamRemoteStorage_GetCachedUGCHandle(remoteStorage(), CachedContent)
}

func PublishWorkshopFile(File string, PreviewFile string, ConsumerAppId common.AppId, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray, WorkshopFileType EWorkshopFileType) common.CallResult[RemoteStoragePublishFileProgress] {
	handle := iSteamRemoteStorage_PublishWorkshopFile(remoteStorage(), File, PreviewFile, ConsumerAppId, Title, Description, Visibility, common.StructToUintptr[SteamParamStringArray](&Tags), WorkshopFileType)
	return common.CallResult[RemoteStoragePublishFileProgress]{Handle: handle}
}

func CreatePublishedFileUpdateRequest(PublishedFileId PublishedFileId) PublishedFileUpdateHandle {
	return iSteamRemoteStorage_CreatePublishedFileUpdateRequest(remoteStorage(), PublishedFileId)
}

func UpdatePublishedFileFile(updateHandle PublishedFileUpdateHandle, File string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileFile(remoteStorage(), updateHandle, File)
}

func UpdatePublishedFilePreviewFile(updateHandle PublishedFileUpdateHandle, PreviewFile string) bool {
	return iSteamRemoteStorage_UpdatePublishedFilePreviewFile(remoteStorage(), updateHandle, PreviewFile)
}

func UpdatePublishedFileTitle(updateHandle PublishedFileUpdateHandle, Title string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileTitle(remoteStorage(), updateHandle, Title)
}

func UpdatePublishedFileDescription(updateHandle PublishedFileUpdateHandle, Description string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileDescription(remoteStorage(), updateHandle, Description)
}

func UpdatePublishedFileVisibility(updateHandle PublishedFileUpdateHandle, Visibility ERemoteStoragePublishedFileVisibility) bool {
	return iSteamRemoteStorage_UpdatePublishedFileVisibility(remoteStorage(), updateHandle, Visibility)
}

func UpdatePublishedFileTags(updateHandle PublishedFileUpdateHandle, Tags SteamParamStringArray) bool {
	return iSteamRemoteStorage_UpdatePublishedFileTags(remoteStorage(), updateHandle, common.StructToUintptr[SteamParamStringArray](&Tags))
}

func CommitPublishedFileUpdate(updateHandle PublishedFileUpdateHandle) common.CallResult[RemoteStorageUpdatePublishedFileResult] {
	handle := iSteamRemoteStorage_CommitPublishedFileUpdate(remoteStorage(), updateHandle)
	return common.CallResult[RemoteStorageUpdatePublishedFileResult]{Handle: handle}
}

func GetPublishedFileDetails(PublishedFileId PublishedFileId, MaxSecondsOld uint32) common.CallResult[RemoteStorageGetPublishedFileDetailsResult] {
	handle := iSteamRemoteStorage_GetPublishedFileDetails(remoteStorage(), PublishedFileId, MaxSecondsOld)
	return common.CallResult[RemoteStorageGetPublishedFileDetailsResult]{Handle: handle}
}

func DeletePublishedFile(PublishedFileId PublishedFileId) common.CallResult[RemoteStorageDeletePublishedFileResult] {
	handle := iSteamRemoteStorage_DeletePublishedFile(remoteStorage(), PublishedFileId)
	return common.CallResult[RemoteStorageDeletePublishedFileResult]{Handle: handle}
}

func EnumerateUserPublishedFiles(StartIndex uint32) common.CallResult[RemoteStorageEnumerateUserPublishedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserPublishedFiles(remoteStorage(), StartIndex)
	return common.CallResult[RemoteStorageEnumerateUserPublishedFilesResult]{Handle: handle}
}

func SubscribePublishedFile(PublishedFileId PublishedFileId) common.CallResult[RemoteStorageSubscribePublishedFileResult] {
	handle := iSteamRemoteStorage_SubscribePublishedFile(remoteStorage(), PublishedFileId)
	return common.CallResult[RemoteStorageSubscribePublishedFileResult]{Handle: handle}
}

func EnumerateUserSubscribedFiles(StartIndex uint32) common.CallResult[RemoteStorageEnumerateUserSubscribedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserSubscribedFiles(remoteStorage(), StartIndex)
	return common.CallResult[RemoteStorageEnumerateUserSubscribedFilesResult]{Handle: handle}
}

func UnsubscribePublishedFile(PublishedFileId PublishedFileId) common.CallResult[RemoteStorageUnsubscribePublishedFileResult] {
	handle := iSteamRemoteStorage_UnsubscribePublishedFile(remoteStorage(), PublishedFileId)
	return common.CallResult[RemoteStorageUnsubscribePublishedFileResult]{Handle: handle}
}

func UpdatePublishedFileSetChangeDescription(updateHandle PublishedFileUpdateHandle, ChangeDescription string) bool {
	return iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription(remoteStorage(), updateHandle, ChangeDescription)
}

func GetPublishedItemVoteDetails(PublishedFileId PublishedFileId) common.CallResult[RemoteStorageGetPublishedItemVoteDetailsResult] {
	handle := iSteamRemoteStorage_GetPublishedItemVoteDetails(remoteStorage(), PublishedFileId)
	return common.CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]{Handle: handle}
}

func UpdateUserPublishedItemVote(PublishedFileId PublishedFileId, VoteUp bool) common.CallResult[RemoteStorageUpdateUserPublishedItemVoteResult] {
	handle := iSteamRemoteStorage_UpdateUserPublishedItemVote(remoteStorage(), PublishedFileId, VoteUp)
	return common.CallResult[RemoteStorageUpdateUserPublishedItemVoteResult]{Handle: handle}
}

func GetUserPublishedItemVoteDetails(PublishedFileId PublishedFileId) common.CallResult[RemoteStorageGetPublishedItemVoteDetailsResult] {
	handle := iSteamRemoteStorage_GetUserPublishedItemVoteDetails(remoteStorage(), PublishedFileId)
	return common.CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]{Handle: handle}
}

func EnumerateUserSharedWorkshopFiles(steamId common.Uint64SteamID, StartIndex uint32, RequiredTags SteamParamStringArray, ExcludedTags SteamParamStringArray) common.CallResult[RemoteStorageEnumerateUserPublishedFilesResult] {
	handle := iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles(remoteStorage(), steamId, StartIndex, common.StructToUintptr[SteamParamStringArray](&RequiredTags), common.StructToUintptr[SteamParamStringArray](&ExcludedTags))
	return common.CallResult[RemoteStorageEnumerateUserPublishedFilesResult]{Handle: handle}
}

func PublishVideo(VideoProvider EWorkshopVideoProvider, VideoAccount string, VideoIdentifier string, PreviewFile string, ConsumerAppId common.AppId, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray) common.CallResult[RemoteStoragePublishFileProgress] {
	handle := iSteamRemoteStorage_PublishVideo(remoteStorage(), VideoProvider, VideoAccount, VideoIdentifier, PreviewFile, ConsumerAppId, Title, Description, Visibility, common.StructToUintptr[SteamParamStringArray](&Tags))
	return common.CallResult[RemoteStoragePublishFileProgress]{Handle: handle}
}

func SetUserPublishedFileAction(PublishedFileId PublishedFileId, Action EWorkshopFileAction) common.CallResult[RemoteStorageSetUserPublishedFileActionResult] {
	handle := iSteamRemoteStorage_SetUserPublishedFileAction(remoteStorage(), PublishedFileId, Action)
	return common.CallResult[RemoteStorageSetUserPublishedFileActionResult]{Handle: handle}
}

func EnumeratePublishedFilesByUserAction(Action EWorkshopFileAction, StartIndex uint32) common.CallResult[RemoteStorageEnumeratePublishedFilesByUserActionResult] {
	handle := iSteamRemoteStorage_EnumeratePublishedFilesByUserAction(remoteStorage(), Action, StartIndex)
	return common.CallResult[RemoteStorageEnumeratePublishedFilesByUserActionResult]{Handle: handle}
}

func EnumeratePublishedWorkshopFiles(EnumerationType EWorkshopEnumerationType, StartIndex uint32, Count uint32, Days uint32) (Tags SteamParamStringArray, UserTags SteamParamStringArray, apiHandle common.CallResult[RemoteStorageEnumerateWorkshopFilesResult]) {
	handle := iSteamRemoteStorage_EnumeratePublishedWorkshopFiles(remoteStorage(), EnumerationType, StartIndex, Count, Days, &Tags, &UserTags)
	return Tags, UserTags, common.CallResult[RemoteStorageEnumerateWorkshopFilesResult]{Handle: handle}
}

func UGCDownloadToLocation(Content common.UGCHandle, Location string, Priority uint32) common.CallResult[RemoteStorageDownloadUGCResult] {
	handle := iSteamRemoteStorage_UGCDownloadToLocation(remoteStorage(), Content, Location, Priority)
	return common.CallResult[RemoteStorageDownloadUGCResult]{Handle: handle}
}

func GetLocalFileChangeCount() int32 {
	return iSteamRemoteStorage_GetLocalFileChangeCount(remoteStorage())
}

func GetLocalFileChange(FileIndex int32) (ChangeType ERemoteStorageLocalFileChange, FilePathType ERemoteStorageFilePathType, result []byte) {
	result = iSteamRemoteStorage_GetLocalFileChange(remoteStorage(), FileIndex, &ChangeType, &FilePathType)
	return ChangeType, FilePathType, result
}

func BeginFileWriteBatch() bool {
	return iSteamRemoteStorage_BeginFileWriteBatch(remoteStorage())
}

func EndFileWriteBatch() bool {
	return iSteamRemoteStorage_EndFileWriteBatch(remoteStorage())
}
