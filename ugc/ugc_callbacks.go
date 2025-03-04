package ugc

import (
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"

	. "github.com/assemblaj/purego-steamworks"
)

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
	return CallbackString(cb)
}

func (cb SteamUGCRequestUGCDetailsResult) CallbackID() SteamCallbackID {
	return SteamUGCRequestUGCDetailsResultID
}

func (cb SteamUGCRequestUGCDetailsResult) Name() string {
	return "SteamUGCRequestUGCDetailsResult"
}

func (cb SteamUGCRequestUGCDetailsResult) String() string {
	return CallbackString(cb)
}

func (cb CreateItemResult) CallbackID() SteamCallbackID {
	return CreateItemResultID
}

func (cb CreateItemResult) Name() string {
	return "CreateItemResult"
}

func (cb CreateItemResult) String() string {
	return CallbackString(cb)
}

func (cb SubmitItemUpdateResult) CallbackID() SteamCallbackID {
	return SubmitItemUpdateResultID
}

func (cb SubmitItemUpdateResult) Name() string {
	return "SubmitItemUpdateResult"
}

func (cb SubmitItemUpdateResult) String() string {
	return CallbackString(cb)
}

func (cb ItemInstalled) CallbackID() SteamCallbackID {
	return ItemInstalledID
}

func (cb ItemInstalled) Name() string {
	return "ItemInstalled"
}

func (cb ItemInstalled) String() string {
	return CallbackString(cb)
}

func (cb DownloadItemResult) CallbackID() SteamCallbackID {
	return DownloadItemResultID
}

func (cb DownloadItemResult) Name() string {
	return "DownloadItemResult"
}

func (cb DownloadItemResult) String() string {
	return CallbackString(cb)
}

func (cb UserFavoriteItemsListChanged) CallbackID() SteamCallbackID {
	return UserFavoriteItemsListChangedID
}

func (cb UserFavoriteItemsListChanged) Name() string {
	return "UserFavoriteItemsListChanged"
}

func (cb UserFavoriteItemsListChanged) String() string {
	return CallbackString(cb)
}

func (cb SetUserItemVoteResult) CallbackID() SteamCallbackID {
	return SetUserItemVoteResultID
}

func (cb SetUserItemVoteResult) Name() string {
	return "SetUserItemVoteResult"
}

func (cb SetUserItemVoteResult) String() string {
	return CallbackString(cb)
}

func (cb GetUserItemVoteResult) CallbackID() SteamCallbackID {
	return GetUserItemVoteResultID
}

func (cb GetUserItemVoteResult) Name() string {
	return "GetUserItemVoteResult"
}

func (cb GetUserItemVoteResult) String() string {
	return CallbackString(cb)
}

func (cb StartPlaytimeTrackingResult) CallbackID() SteamCallbackID {
	return StartPlaytimeTrackingResultID
}

func (cb StartPlaytimeTrackingResult) Name() string {
	return "StartPlaytimeTrackingResult"
}

func (cb StartPlaytimeTrackingResult) String() string {
	return CallbackString(cb)
}

func (cb StopPlaytimeTrackingResult) CallbackID() SteamCallbackID {
	return StopPlaytimeTrackingResultID
}

func (cb StopPlaytimeTrackingResult) Name() string {
	return "StopPlaytimeTrackingResult"
}

func (cb StopPlaytimeTrackingResult) String() string {
	return CallbackString(cb)
}

func (cb AddUGCDependencyResult) CallbackID() SteamCallbackID {
	return AddUGCDependencyResultID
}

func (cb AddUGCDependencyResult) Name() string {
	return "AddUGCDependencyResult"
}

func (cb AddUGCDependencyResult) String() string {
	return CallbackString(cb)
}

func (cb RemoveUGCDependencyResult) CallbackID() SteamCallbackID {
	return RemoveUGCDependencyResultID
}

func (cb RemoveUGCDependencyResult) Name() string {
	return "RemoveUGCDependencyResult"
}

func (cb RemoveUGCDependencyResult) String() string {
	return CallbackString(cb)
}

func (cb AddAppDependencyResult) CallbackID() SteamCallbackID {
	return AddAppDependencyResultID
}

func (cb AddAppDependencyResult) Name() string {
	return "AddAppDependencyResult"
}

func (cb AddAppDependencyResult) String() string {
	return CallbackString(cb)
}

func (cb RemoveAppDependencyResult) CallbackID() SteamCallbackID {
	return RemoveAppDependencyResultID
}

func (cb RemoveAppDependencyResult) Name() string {
	return "RemoveAppDependencyResult"
}

func (cb RemoveAppDependencyResult) String() string {
	return CallbackString(cb)
}

func (cb GetAppDependenciesResult) CallbackID() SteamCallbackID {
	return GetAppDependenciesResultID
}

func (cb GetAppDependenciesResult) Name() string {
	return "GetAppDependenciesResult"
}

func (cb GetAppDependenciesResult) String() string {
	return CallbackString(cb)
}

func (cb DeleteItemResult) CallbackID() SteamCallbackID {
	return DeleteItemResultID
}

func (cb DeleteItemResult) Name() string {
	return "DeleteItemResult"
}

func (cb DeleteItemResult) String() string {
	return CallbackString(cb)
}

func (cb UserSubscribedItemsListChanged) CallbackID() SteamCallbackID {
	return UserSubscribedItemsListChangedID
}

func (cb UserSubscribedItemsListChanged) Name() string {
	return "UserSubscribedItemsListChanged"
}

func (cb UserSubscribedItemsListChanged) String() string {
	return CallbackString(cb)
}

func (cb WorkshopEULAStatus) CallbackID() SteamCallbackID {
	return WorkshopEULAStatusID
}

func (cb WorkshopEULAStatus) Name() string {
	return "WorkshopEULAStatus"
}

func (cb WorkshopEULAStatus) String() string {
	return CallbackString(cb)
}
