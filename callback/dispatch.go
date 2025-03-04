package callback

import (
	"unsafe"

	common "github.com/assemblaj/purego-steamworks"
	"github.com/assemblaj/purego-steamworks/apps"
	. "github.com/assemblaj/purego-steamworks/friends"
	. "github.com/assemblaj/purego-steamworks/game_search"
	. "github.com/assemblaj/purego-steamworks/game_server"
	. "github.com/assemblaj/purego-steamworks/game_server_stats"
	. "github.com/assemblaj/purego-steamworks/input"
	. "github.com/assemblaj/purego-steamworks/inventory"
	. "github.com/assemblaj/purego-steamworks/matchmaking"
	. "github.com/assemblaj/purego-steamworks/networking_messages"
	. "github.com/assemblaj/purego-steamworks/networking_sockets"
	. "github.com/assemblaj/purego-steamworks/networking_utils"
	. "github.com/assemblaj/purego-steamworks/parties"
	. "github.com/assemblaj/purego-steamworks/remote_play"
	. "github.com/assemblaj/purego-steamworks/remote_storage"
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"
	. "github.com/assemblaj/purego-steamworks/screenshots"
	. "github.com/assemblaj/purego-steamworks/ugc"
	user "github.com/assemblaj/purego-steamworks/user"
	userstats "github.com/assemblaj/purego-steamworks/user_stats"
	. "github.com/assemblaj/purego-steamworks/utils"
)

func dispatchCallback(callback *common.CallbackMsg) {
	if common.GameServerActive {
		switch common.SteamCallbackID(callback.Callback) {
		case GSStatsReceivedID:
			handleRawData[GSStatsReceived](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSStatsStoredID:
			handleRawData[GSStatsStored](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSStatsUnloadedID:
			handleRawData[GSStatsUnloaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientApproveID:
			handleRawData[GSClientApprove](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientDenyID:
			handleRawData[GSClientDeny](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientKickID:
			handleRawData[GSClientKick](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientAchievementStatusID:
			handleRawData[GSClientAchievementStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSPolicyResponseID:
			handleRawData[GSPolicyResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSGameplayStatsID:
			handleRawData[GSGameplayStats](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientGroupStatusID:
			handleRawData[GSClientGroupStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSReputationID:
			handleRawData[GSReputation](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case AssociateWithClanResultID:
			handleRawData[AssociateWithClanResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case ComputeNewPlayerCompatibilityResultID:
			handleRawData[ComputeNewPlayerCompatibilityResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		}
	}
	switch common.SteamCallbackID(callback.Callback) {
	case apps.DlcInstalledID:
		handleRawData[apps.DlcInstalled](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case apps.NewUrlLaunchParametersID:
		handleRawData[apps.NewUrlLaunchParameters](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case apps.AppProofOfPurchaseKeyResponseID:
		handleRawData[apps.AppProofOfPurchaseKeyResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case apps.FileDetailsResultID:
		handleRawData[apps.FileDetailsResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case apps.TimedTrialStatusID:
		handleRawData[apps.TimedTrialStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case PersonaStateChangeID:
		handleRawData[PersonaStateChange](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameOverlayActivatedID:
		handleRawData[GameOverlayActivated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameServerChangeRequestedID:
		handleRawData[GameServerChangeRequested](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameLobbyJoinRequestedID:
		handleRawData[GameLobbyJoinRequested](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AvatarImageLoadedID:
		handleRawData[AvatarImageLoaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ClanOfficerListResponseID:
		handleRawData[ClanOfficerListResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendRichPresenceUpdateID:
		handleRawData[FriendRichPresenceUpdate](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameRichPresenceJoinRequestedID:
		handleRawData[GameRichPresenceJoinRequested](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedClanChatMsgID:
		handleRawData[GameConnectedClanChatMsg](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedChatJoinID:
		handleRawData[GameConnectedChatJoin](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedChatLeaveID:
		handleRawData[GameConnectedChatLeave](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DownloadClanActivityCountsResultID:
		handleRawData[DownloadClanActivityCountsResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case JoinClanChatRoomCompletionResultID:
		handleRawData[JoinClanChatRoomCompletionResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedFriendChatMsgID:
		handleRawData[GameConnectedFriendChatMsg](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsGetFollowerCountID:
		handleRawData[FriendsGetFollowerCount](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsIsFollowingID:
		handleRawData[FriendsIsFollowing](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsEnumerateFollowingListID:
		handleRawData[FriendsEnumerateFollowingList](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SetPersonaNameResponseID:
		handleRawData[SetPersonaNameResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UnreadChatMessagesChangedID:
		handleRawData[UnreadChatMessagesChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case OverlayBrowserProtocolNavigationID:
		handleRawData[OverlayBrowserProtocolNavigation](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EquippedProfileItemsChangedID:
		handleRawData[EquippedProfileItemsChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EquippedProfileItemsID:
		handleRawData[EquippedProfileItems](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SearchForGameProgressCallbackID:
		handleRawData[SearchForGameProgressCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SearchForGameResultCallbackID:
		handleRawData[SearchForGameResultCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameProgressCallbackID:
		handleRawData[RequestPlayersForGameProgressCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameResultCallbackID:
		handleRawData[RequestPlayersForGameResultCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameFinalResultCallbackID:
		handleRawData[RequestPlayersForGameFinalResultCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SubmitPlayerResultResultCallbackID:
		handleRawData[SubmitPlayerResultResultCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EndGameResultCallbackID:
		handleRawData[EndGameResultCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputDeviceConnectedID:
		handleRawData[SteamInputDeviceConnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputDeviceDisconnectedID:
		handleRawData[SteamInputDeviceDisconnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputConfigurationLoadedID:
		handleRawData[SteamInputConfigurationLoaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputGamepadSlotChangeID:
		handleRawData[SteamInputGamepadSlotChange](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryResultReadyID:
		handleRawData[SteamInventoryResultReady](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryFullUpdateID:
		handleRawData[SteamInventoryFullUpdate](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryDefinitionUpdateID:
		handleRawData[SteamInventoryDefinitionUpdate](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryEligiblePromoItemDefIDsID:
		handleRawData[SteamInventoryEligiblePromoItemDefIDs](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryStartPurchaseResultID:
		handleRawData[SteamInventoryStartPurchaseResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryRequestPricesResultID:
		handleRawData[SteamInventoryRequestPricesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FavoritesListChangedID:
		handleRawData[FavoritesListChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyInviteID:
		handleRawData[LobbyInvite](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyEnterID:
		handleRawData[LobbyEnter](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyDataUpdateID:
		handleRawData[LobbyDataUpdate](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyChatUpdateID:
		handleRawData[LobbyChatUpdate](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyChatMsgID:
		handleRawData[LobbyChatMsg](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyGameCreatedID:
		handleRawData[LobbyGameCreated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyMatchListID:
		handleRawData[LobbyMatchList](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyKickedID:
		handleRawData[LobbyKicked](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyCreatedID:
		handleRawData[LobbyCreated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case PSNGameBootInviteResultID:
		handleRawData[PSNGameBootInviteResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FavoritesListAccountsUpdatedID:
		handleRawData[FavoritesListAccountsUpdated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingMessagesSessionRequestID:
		handleRawData[SteamNetworkingMessagesSessionRequest](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingMessagesSessionFailedID:
		handleRawData[SteamNetworkingMessagesSessionFailed](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetAuthenticationStatusID:
		handleRawData[SteamNetAuthenticationStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetConnectionStatusChangedCallbackID:
		handleRawData[SteamNetConnectionStatusChangedCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingFakeIPResultID:
		handleRawData[SteamNetworkingFakeIPResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRelayNetworkStatusID:
		handleRawData[SteamRelayNetworkStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case JoinPartyCallbackID:
		handleRawData[JoinPartyCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CreateBeaconCallbackID:
		handleRawData[CreateBeaconCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ReservationNotificationCallbackID:
		handleRawData[ReservationNotificationCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ChangeNumOpenSlotsCallbackID:
		handleRawData[ChangeNumOpenSlotsCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AvailableBeaconLocationsUpdatedID:
		handleRawData[AvailableBeaconLocationsUpdated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ActiveBeaconsUpdatedID:
		handleRawData[ActiveBeaconsUpdated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlaySessionConnectedID:
		handleRawData[SteamRemotePlaySessionConnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlaySessionDisconnectedID:
		handleRawData[SteamRemotePlaySessionDisconnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlayTogetherGuestInviteID:
		handleRawData[SteamRemotePlayTogetherGuestInvite](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileShareResultID:
		handleRawData[RemoteStorageFileShareResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishFileResultID:
		handleRawData[RemoteStoragePublishFileResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageDeletePublishedFileResultID:
		handleRawData[RemoteStorageDeletePublishedFileResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserPublishedFilesResultID:
		handleRawData[RemoteStorageEnumerateUserPublishedFilesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageSubscribePublishedFileResultID:
		handleRawData[RemoteStorageSubscribePublishedFileResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserSubscribedFilesResultID:
		handleRawData[RemoteStorageEnumerateUserSubscribedFilesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUnsubscribePublishedFileResultID:
		handleRawData[RemoteStorageUnsubscribePublishedFileResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUpdatePublishedFileResultID:
		handleRawData[RemoteStorageUpdatePublishedFileResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageDownloadUGCResultID:
		handleRawData[RemoteStorageDownloadUGCResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageGetPublishedFileDetailsResultID:
		handleRawData[RemoteStorageGetPublishedFileDetailsResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateWorkshopFilesResultID:
		handleRawData[RemoteStorageEnumerateWorkshopFilesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageGetPublishedItemVoteDetailsResultID:
		handleRawData[RemoteStorageGetPublishedItemVoteDetailsResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileSubscribedID:
		handleRawData[RemoteStoragePublishedFileSubscribed](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileUnsubscribedID:
		handleRawData[RemoteStoragePublishedFileUnsubscribed](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileDeletedID:
		handleRawData[RemoteStoragePublishedFileDeleted](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUpdateUserPublishedItemVoteResultID:
		handleRawData[RemoteStorageUpdateUserPublishedItemVoteResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUserVoteDetailsID:
		handleRawData[RemoteStorageUserVoteDetails](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserSharedWorkshopFilesResultID:
		handleRawData[RemoteStorageEnumerateUserSharedWorkshopFilesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageSetUserPublishedFileActionResultID:
		handleRawData[RemoteStorageSetUserPublishedFileActionResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumeratePublishedFilesByUserActionResultID:
		handleRawData[RemoteStorageEnumeratePublishedFilesByUserActionResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishFileProgressID:
		handleRawData[RemoteStoragePublishFileProgress](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileUpdatedID:
		handleRawData[RemoteStoragePublishedFileUpdated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileWriteAsyncCompleteID:
		handleRawData[RemoteStorageFileWriteAsyncComplete](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileReadAsyncCompleteID:
		handleRawData[RemoteStorageFileReadAsyncComplete](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageLocalFileChangeID:
		handleRawData[RemoteStorageLocalFileChange](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ScreenshotReadyID:
		handleRawData[ScreenshotReady](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ScreenshotRequestedID:
		handleRawData[ScreenshotRequested](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamUGCQueryCompletedID:
		handleRawData[SteamUGCQueryCompleted](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamUGCRequestUGCDetailsResultID:
		handleRawData[SteamUGCRequestUGCDetailsResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CreateItemResultID:
		handleRawData[CreateItemResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SubmitItemUpdateResultID:
		handleRawData[SubmitItemUpdateResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ItemInstalledID:
		handleRawData[ItemInstalled](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DownloadItemResultID:
		handleRawData[DownloadItemResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserFavoriteItemsListChangedID:
		handleRawData[UserFavoriteItemsListChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SetUserItemVoteResultID:
		handleRawData[SetUserItemVoteResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetUserItemVoteResultID:
		handleRawData[GetUserItemVoteResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case StartPlaytimeTrackingResultID:
		handleRawData[StartPlaytimeTrackingResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case StopPlaytimeTrackingResultID:
		handleRawData[StopPlaytimeTrackingResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AddUGCDependencyResultID:
		handleRawData[AddUGCDependencyResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoveUGCDependencyResultID:
		handleRawData[RemoveUGCDependencyResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AddAppDependencyResultID:
		handleRawData[AddAppDependencyResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoveAppDependencyResultID:
		handleRawData[RemoveAppDependencyResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetAppDependenciesResultID:
		handleRawData[GetAppDependenciesResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DeleteItemResultID:
		handleRawData[DeleteItemResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserSubscribedItemsListChangedID:
		handleRawData[UserSubscribedItemsListChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case WorkshopEULAStatusID:
		handleRawData[WorkshopEULAStatus](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.SteamServersConnectedID:
		handleRawData[user.SteamServersConnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.SteamServerConnectFailureID:
		handleRawData[user.SteamServerConnectFailure](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.SteamServersDisconnectedID:
		handleRawData[user.SteamServersDisconnected](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.ClientGameServerDenyID:
		handleRawData[user.ClientGameServerDeny](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.IPCFailureID:
		handleRawData[user.IPCFailure](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.LicensesUpdatedID:
		handleRawData[user.LicensesUpdated](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.ValidateAuthTicketResponseID:
		handleRawData[user.ValidateAuthTicketResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.MicroTxnAuthorizationResponseID:
		handleRawData[user.MicroTxnAuthorizationResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.EncryptedAppTicketResponseID:
		handleRawData[user.EncryptedAppTicketResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.GetAuthSessionTicketResponseID:
		handleRawData[user.GetAuthSessionTicketResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.GameWebCallbackID:
		handleRawData[user.GameWebCallback](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.StoreAuthURLResponseID:
		handleRawData[user.StoreAuthURLResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.MarketEligibilityResponseID:
		handleRawData[user.MarketEligibilityResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.DurationControlID:
		handleRawData[user.DurationControl](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case user.GetTicketForWebApiResponseID:
		handleRawData[user.GetTicketForWebApiResponse](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.UserStatsReceivedID:
		handleRawData[userstats.UserStatsReceived](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.UserStatsStoredID:
		handleRawData[userstats.UserStatsStored](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.UserAchievementStoredID:
		handleRawData[userstats.UserAchievementStored](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.LeaderboardFindResultID:
		handleRawData[userstats.LeaderboardFindResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.LeaderboardScoresDownloadedID:
		handleRawData[userstats.LeaderboardScoresDownloaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.LeaderboardScoreUploadedID:
		handleRawData[userstats.LeaderboardScoreUploaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.NumberOfCurrentPlayersID:
		handleRawData[userstats.NumberOfCurrentPlayers](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.UserStatsUnloadedID:
		handleRawData[userstats.UserStatsUnloaded](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.UserAchievementIconFetchedID:
		handleRawData[userstats.UserAchievementIconFetched](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.GlobalAchievementPercentagesReadyID:
		handleRawData[userstats.GlobalAchievementPercentagesReady](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.LeaderboardUGCSetID:
		handleRawData[userstats.LeaderboardUGCSet](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case userstats.GlobalStatsReceivedID:
		handleRawData[userstats.GlobalStatsReceived](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case IPCountryID:
		handleRawData[IPCountry](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LowBatteryPowerID:
		handleRawData[LowBatteryPower](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamShutdownID:
		handleRawData[SteamShutdown](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CheckFileSignatureID:
		handleRawData[CheckFileSignatureResult](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GamepadTextInputDismissedID:
		handleRawData[GamepadTextInputDismissed](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AppResumingFromSuspendID:
		handleRawData[AppResumingFromSuspend](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FloatingGamepadTextInputDismissedID:
		handleRawData[FloatingGamepadTextInputDismissed](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FilterTextDictionaryChangedID:
		handleRawData[FilterTextDictionaryChanged](common.SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	}
}

func dispatchCallResult(apiHandle common.SteamAPICall, failed bool, rawData uintptr, callbackID common.SteamCallbackID) {
	if common.GameServerActive {
		switch common.SteamCallbackID(callbackID) {
		case GSStatsReceivedID:
			handleRawResult[GSStatsReceived](apiHandle, failed, unsafe.Pointer(rawData))
		case GSStatsStoredID:
			handleRawResult[GSStatsStored](apiHandle, failed, unsafe.Pointer(rawData))
		case GSStatsUnloadedID:
			handleRawResult[GSStatsUnloaded](apiHandle, failed, unsafe.Pointer(rawData))
		case GSClientApproveID:
			handleRawResult[GSClientApprove](apiHandle, failed, unsafe.Pointer(rawData))
		case GSClientDenyID:
			handleRawResult[GSClientDeny](apiHandle, failed, unsafe.Pointer(rawData))
		case GSClientKickID:
			handleRawResult[GSClientKick](apiHandle, failed, unsafe.Pointer(rawData))
		case GSClientAchievementStatusID:
			handleRawResult[GSClientAchievementStatus](apiHandle, failed, unsafe.Pointer(rawData))
		case GSPolicyResponseID:
			handleRawResult[GSPolicyResponse](apiHandle, failed, unsafe.Pointer(rawData))
		case GSGameplayStatsID:
			handleRawResult[GSGameplayStats](apiHandle, failed, unsafe.Pointer(rawData))
		case GSClientGroupStatusID:
			handleRawResult[GSClientGroupStatus](apiHandle, failed, unsafe.Pointer(rawData))
		case GSReputationID:
			handleRawResult[GSReputation](apiHandle, failed, unsafe.Pointer(rawData))
		case AssociateWithClanResultID:
			handleRawResult[AssociateWithClanResult](apiHandle, failed, unsafe.Pointer(rawData))
		case ComputeNewPlayerCompatibilityResultID:
			handleRawResult[ComputeNewPlayerCompatibilityResult](apiHandle, failed, unsafe.Pointer(rawData))
		}
	}
	switch common.SteamCallbackID(callbackID) {
	case apps.DlcInstalledID:
		handleRawResult[apps.DlcInstalled](apiHandle, failed, unsafe.Pointer(rawData))
	case apps.NewUrlLaunchParametersID:
		handleRawResult[apps.NewUrlLaunchParameters](apiHandle, failed, unsafe.Pointer(rawData))
	case apps.AppProofOfPurchaseKeyResponseID:
		handleRawResult[apps.AppProofOfPurchaseKeyResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case apps.FileDetailsResultID:
		handleRawResult[apps.FileDetailsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case apps.TimedTrialStatusID:
		handleRawResult[apps.TimedTrialStatus](apiHandle, failed, unsafe.Pointer(rawData))
	case PersonaStateChangeID:
		handleRawResult[PersonaStateChange](apiHandle, failed, unsafe.Pointer(rawData))
	case GameOverlayActivatedID:
		handleRawResult[GameOverlayActivated](apiHandle, failed, unsafe.Pointer(rawData))
	case GameServerChangeRequestedID:
		handleRawResult[GameServerChangeRequested](apiHandle, failed, unsafe.Pointer(rawData))
	case GameLobbyJoinRequestedID:
		handleRawResult[GameLobbyJoinRequested](apiHandle, failed, unsafe.Pointer(rawData))
	case AvatarImageLoadedID:
		handleRawResult[AvatarImageLoaded](apiHandle, failed, unsafe.Pointer(rawData))
	case ClanOfficerListResponseID:
		handleRawResult[ClanOfficerListResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case FriendRichPresenceUpdateID:
		handleRawResult[FriendRichPresenceUpdate](apiHandle, failed, unsafe.Pointer(rawData))
	case GameRichPresenceJoinRequestedID:
		handleRawResult[GameRichPresenceJoinRequested](apiHandle, failed, unsafe.Pointer(rawData))
	case GameConnectedClanChatMsgID:
		handleRawResult[GameConnectedClanChatMsg](apiHandle, failed, unsafe.Pointer(rawData))
	case GameConnectedChatJoinID:
		handleRawResult[GameConnectedChatJoin](apiHandle, failed, unsafe.Pointer(rawData))
	case GameConnectedChatLeaveID:
		handleRawResult[GameConnectedChatLeave](apiHandle, failed, unsafe.Pointer(rawData))
	case DownloadClanActivityCountsResultID:
		handleRawResult[DownloadClanActivityCountsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case JoinClanChatRoomCompletionResultID:
		handleRawResult[JoinClanChatRoomCompletionResult](apiHandle, failed, unsafe.Pointer(rawData))
	case GameConnectedFriendChatMsgID:
		handleRawResult[GameConnectedFriendChatMsg](apiHandle, failed, unsafe.Pointer(rawData))
	case FriendsGetFollowerCountID:
		handleRawResult[FriendsGetFollowerCount](apiHandle, failed, unsafe.Pointer(rawData))
	case FriendsIsFollowingID:
		handleRawResult[FriendsIsFollowing](apiHandle, failed, unsafe.Pointer(rawData))
	case FriendsEnumerateFollowingListID:
		handleRawResult[FriendsEnumerateFollowingList](apiHandle, failed, unsafe.Pointer(rawData))
	case SetPersonaNameResponseID:
		handleRawResult[SetPersonaNameResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case UnreadChatMessagesChangedID:
		handleRawResult[UnreadChatMessagesChanged](apiHandle, failed, unsafe.Pointer(rawData))
	case OverlayBrowserProtocolNavigationID:
		handleRawResult[OverlayBrowserProtocolNavigation](apiHandle, failed, unsafe.Pointer(rawData))
	case EquippedProfileItemsChangedID:
		handleRawResult[EquippedProfileItemsChanged](apiHandle, failed, unsafe.Pointer(rawData))
	case EquippedProfileItemsID:
		handleRawResult[EquippedProfileItems](apiHandle, failed, unsafe.Pointer(rawData))
	case SearchForGameProgressCallbackID:
		handleRawResult[SearchForGameProgressCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case SearchForGameResultCallbackID:
		handleRawResult[SearchForGameResultCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case RequestPlayersForGameProgressCallbackID:
		handleRawResult[RequestPlayersForGameProgressCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case RequestPlayersForGameResultCallbackID:
		handleRawResult[RequestPlayersForGameResultCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case RequestPlayersForGameFinalResultCallbackID:
		handleRawResult[RequestPlayersForGameFinalResultCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case SubmitPlayerResultResultCallbackID:
		handleRawResult[SubmitPlayerResultResultCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case EndGameResultCallbackID:
		handleRawResult[EndGameResultCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInputDeviceConnectedID:
		handleRawResult[SteamInputDeviceConnected](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInputDeviceDisconnectedID:
		handleRawResult[SteamInputDeviceDisconnected](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInputConfigurationLoadedID:
		handleRawResult[SteamInputConfigurationLoaded](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInputGamepadSlotChangeID:
		handleRawResult[SteamInputGamepadSlotChange](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryResultReadyID:
		handleRawResult[SteamInventoryResultReady](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryFullUpdateID:
		handleRawResult[SteamInventoryFullUpdate](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryDefinitionUpdateID:
		handleRawResult[SteamInventoryDefinitionUpdate](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryEligiblePromoItemDefIDsID:
		handleRawResult[SteamInventoryEligiblePromoItemDefIDs](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryStartPurchaseResultID:
		handleRawResult[SteamInventoryStartPurchaseResult](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamInventoryRequestPricesResultID:
		handleRawResult[SteamInventoryRequestPricesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case FavoritesListChangedID:
		handleRawResult[FavoritesListChanged](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyInviteID:
		handleRawResult[LobbyInvite](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyEnterID:
		handleRawResult[LobbyEnter](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyDataUpdateID:
		handleRawResult[LobbyDataUpdate](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyChatUpdateID:
		handleRawResult[LobbyChatUpdate](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyChatMsgID:
		handleRawResult[LobbyChatMsg](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyGameCreatedID:
		handleRawResult[LobbyGameCreated](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyMatchListID:
		handleRawResult[LobbyMatchList](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyKickedID:
		handleRawResult[LobbyKicked](apiHandle, failed, unsafe.Pointer(rawData))
	case LobbyCreatedID:
		handleRawResult[LobbyCreated](apiHandle, failed, unsafe.Pointer(rawData))
	case PSNGameBootInviteResultID:
		handleRawResult[PSNGameBootInviteResult](apiHandle, failed, unsafe.Pointer(rawData))
	case FavoritesListAccountsUpdatedID:
		handleRawResult[FavoritesListAccountsUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamNetworkingMessagesSessionRequestID:
		handleRawResult[SteamNetworkingMessagesSessionRequest](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamNetworkingMessagesSessionFailedID:
		handleRawResult[SteamNetworkingMessagesSessionFailed](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamNetAuthenticationStatusID:
		handleRawResult[SteamNetAuthenticationStatus](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamNetConnectionStatusChangedCallbackID:
		handleRawResult[SteamNetConnectionStatusChangedCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamNetworkingFakeIPResultID:
		handleRawResult[SteamNetworkingFakeIPResult](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamRelayNetworkStatusID:
		handleRawResult[SteamRelayNetworkStatus](apiHandle, failed, unsafe.Pointer(rawData))
	case JoinPartyCallbackID:
		handleRawResult[JoinPartyCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case CreateBeaconCallbackID:
		handleRawResult[CreateBeaconCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case ReservationNotificationCallbackID:
		handleRawResult[ReservationNotificationCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case ChangeNumOpenSlotsCallbackID:
		handleRawResult[ChangeNumOpenSlotsCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case AvailableBeaconLocationsUpdatedID:
		handleRawResult[AvailableBeaconLocationsUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case ActiveBeaconsUpdatedID:
		handleRawResult[ActiveBeaconsUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamRemotePlaySessionConnectedID:
		handleRawResult[SteamRemotePlaySessionConnected](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamRemotePlaySessionDisconnectedID:
		handleRawResult[SteamRemotePlaySessionDisconnected](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamRemotePlayTogetherGuestInviteID:
		handleRawResult[SteamRemotePlayTogetherGuestInvite](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageFileShareResultID:
		handleRawResult[RemoteStorageFileShareResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishFileResultID:
		handleRawResult[RemoteStoragePublishFileResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageDeletePublishedFileResultID:
		handleRawResult[RemoteStorageDeletePublishedFileResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageEnumerateUserPublishedFilesResultID:
		handleRawResult[RemoteStorageEnumerateUserPublishedFilesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageSubscribePublishedFileResultID:
		handleRawResult[RemoteStorageSubscribePublishedFileResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageEnumerateUserSubscribedFilesResultID:
		handleRawResult[RemoteStorageEnumerateUserSubscribedFilesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageUnsubscribePublishedFileResultID:
		handleRawResult[RemoteStorageUnsubscribePublishedFileResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageUpdatePublishedFileResultID:
		handleRawResult[RemoteStorageUpdatePublishedFileResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageDownloadUGCResultID:
		handleRawResult[RemoteStorageDownloadUGCResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageGetPublishedFileDetailsResultID:
		handleRawResult[RemoteStorageGetPublishedFileDetailsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageEnumerateWorkshopFilesResultID:
		handleRawResult[RemoteStorageEnumerateWorkshopFilesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageGetPublishedItemVoteDetailsResultID:
		handleRawResult[RemoteStorageGetPublishedItemVoteDetailsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishedFileSubscribedID:
		handleRawResult[RemoteStoragePublishedFileSubscribed](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishedFileUnsubscribedID:
		handleRawResult[RemoteStoragePublishedFileUnsubscribed](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishedFileDeletedID:
		handleRawResult[RemoteStoragePublishedFileDeleted](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageUpdateUserPublishedItemVoteResultID:
		handleRawResult[RemoteStorageUpdateUserPublishedItemVoteResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageUserVoteDetailsID:
		handleRawResult[RemoteStorageUserVoteDetails](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageEnumerateUserSharedWorkshopFilesResultID:
		handleRawResult[RemoteStorageEnumerateUserSharedWorkshopFilesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageSetUserPublishedFileActionResultID:
		handleRawResult[RemoteStorageSetUserPublishedFileActionResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageEnumeratePublishedFilesByUserActionResultID:
		handleRawResult[RemoteStorageEnumeratePublishedFilesByUserActionResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishFileProgressID:
		handleRawResult[RemoteStoragePublishFileProgress](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStoragePublishedFileUpdatedID:
		handleRawResult[RemoteStoragePublishedFileUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageFileWriteAsyncCompleteID:
		handleRawResult[RemoteStorageFileWriteAsyncComplete](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageFileReadAsyncCompleteID:
		handleRawResult[RemoteStorageFileReadAsyncComplete](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoteStorageLocalFileChangeID:
		handleRawResult[RemoteStorageLocalFileChange](apiHandle, failed, unsafe.Pointer(rawData))
	case ScreenshotReadyID:
		handleRawResult[ScreenshotReady](apiHandle, failed, unsafe.Pointer(rawData))
	case ScreenshotRequestedID:
		handleRawResult[ScreenshotRequested](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamUGCQueryCompletedID:
		handleRawResult[SteamUGCQueryCompleted](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamUGCRequestUGCDetailsResultID:
		handleRawResult[SteamUGCRequestUGCDetailsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case CreateItemResultID:
		handleRawResult[CreateItemResult](apiHandle, failed, unsafe.Pointer(rawData))
	case SubmitItemUpdateResultID:
		handleRawResult[SubmitItemUpdateResult](apiHandle, failed, unsafe.Pointer(rawData))
	case ItemInstalledID:
		handleRawResult[ItemInstalled](apiHandle, failed, unsafe.Pointer(rawData))
	case DownloadItemResultID:
		handleRawResult[DownloadItemResult](apiHandle, failed, unsafe.Pointer(rawData))
	case UserFavoriteItemsListChangedID:
		handleRawResult[UserFavoriteItemsListChanged](apiHandle, failed, unsafe.Pointer(rawData))
	case SetUserItemVoteResultID:
		handleRawResult[SetUserItemVoteResult](apiHandle, failed, unsafe.Pointer(rawData))
	case GetUserItemVoteResultID:
		handleRawResult[GetUserItemVoteResult](apiHandle, failed, unsafe.Pointer(rawData))
	case StartPlaytimeTrackingResultID:
		handleRawResult[StartPlaytimeTrackingResult](apiHandle, failed, unsafe.Pointer(rawData))
	case StopPlaytimeTrackingResultID:
		handleRawResult[StopPlaytimeTrackingResult](apiHandle, failed, unsafe.Pointer(rawData))
	case AddUGCDependencyResultID:
		handleRawResult[AddUGCDependencyResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoveUGCDependencyResultID:
		handleRawResult[RemoveUGCDependencyResult](apiHandle, failed, unsafe.Pointer(rawData))
	case AddAppDependencyResultID:
		handleRawResult[AddAppDependencyResult](apiHandle, failed, unsafe.Pointer(rawData))
	case RemoveAppDependencyResultID:
		handleRawResult[RemoveAppDependencyResult](apiHandle, failed, unsafe.Pointer(rawData))
	case GetAppDependenciesResultID:
		handleRawResult[GetAppDependenciesResult](apiHandle, failed, unsafe.Pointer(rawData))
	case DeleteItemResultID:
		handleRawResult[DeleteItemResult](apiHandle, failed, unsafe.Pointer(rawData))
	case UserSubscribedItemsListChangedID:
		handleRawResult[UserSubscribedItemsListChanged](apiHandle, failed, unsafe.Pointer(rawData))
	case WorkshopEULAStatusID:
		handleRawResult[WorkshopEULAStatus](apiHandle, failed, unsafe.Pointer(rawData))
	case user.SteamServersConnectedID:
		handleRawResult[user.SteamServersConnected](apiHandle, failed, unsafe.Pointer(rawData))
	case user.SteamServerConnectFailureID:
		handleRawResult[user.SteamServerConnectFailure](apiHandle, failed, unsafe.Pointer(rawData))
	case user.SteamServersDisconnectedID:
		handleRawResult[user.SteamServersDisconnected](apiHandle, failed, unsafe.Pointer(rawData))
	case user.ClientGameServerDenyID:
		handleRawResult[user.ClientGameServerDeny](apiHandle, failed, unsafe.Pointer(rawData))
	case user.IPCFailureID:
		handleRawResult[user.IPCFailure](apiHandle, failed, unsafe.Pointer(rawData))
	case user.LicensesUpdatedID:
		handleRawResult[user.LicensesUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case user.ValidateAuthTicketResponseID:
		handleRawResult[user.ValidateAuthTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.MicroTxnAuthorizationResponseID:
		handleRawResult[user.MicroTxnAuthorizationResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.EncryptedAppTicketResponseID:
		handleRawResult[user.EncryptedAppTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.GetAuthSessionTicketResponseID:
		handleRawResult[user.GetAuthSessionTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.GameWebCallbackID:
		handleRawResult[user.GameWebCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case user.StoreAuthURLResponseID:
		handleRawResult[user.StoreAuthURLResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.MarketEligibilityResponseID:
		handleRawResult[user.MarketEligibilityResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case user.DurationControlID:
		handleRawResult[user.DurationControl](apiHandle, failed, unsafe.Pointer(rawData))
	case user.GetTicketForWebApiResponseID:
		handleRawResult[user.GetTicketForWebApiResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.UserStatsReceivedID:
		handleRawResult[userstats.UserStatsReceived](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.UserStatsStoredID:
		handleRawResult[userstats.UserStatsStored](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.UserAchievementStoredID:
		handleRawResult[userstats.UserAchievementStored](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.LeaderboardFindResultID:
		handleRawResult[userstats.LeaderboardFindResult](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.LeaderboardScoresDownloadedID:
		handleRawResult[userstats.LeaderboardScoresDownloaded](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.LeaderboardScoreUploadedID:
		handleRawResult[userstats.LeaderboardScoreUploaded](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.NumberOfCurrentPlayersID:
		handleRawResult[userstats.NumberOfCurrentPlayers](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.UserStatsUnloadedID:
		handleRawResult[userstats.UserStatsUnloaded](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.UserAchievementIconFetchedID:
		handleRawResult[userstats.UserAchievementIconFetched](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.GlobalAchievementPercentagesReadyID:
		handleRawResult[userstats.GlobalAchievementPercentagesReady](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.LeaderboardUGCSetID:
		handleRawResult[userstats.LeaderboardUGCSet](apiHandle, failed, unsafe.Pointer(rawData))
	case userstats.GlobalStatsReceivedID:
		handleRawResult[userstats.GlobalStatsReceived](apiHandle, failed, unsafe.Pointer(rawData))
	case IPCountryID:
		handleRawResult[IPCountry](apiHandle, failed, unsafe.Pointer(rawData))
	case LowBatteryPowerID:
		handleRawResult[LowBatteryPower](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamShutdownID:
		handleRawResult[SteamShutdown](apiHandle, failed, unsafe.Pointer(rawData))
	case CheckFileSignatureID:
		handleRawResult[CheckFileSignatureResult](apiHandle, failed, unsafe.Pointer(rawData))
	case GamepadTextInputDismissedID:
		handleRawResult[GamepadTextInputDismissed](apiHandle, failed, unsafe.Pointer(rawData))
	case AppResumingFromSuspendID:
		handleRawResult[AppResumingFromSuspend](apiHandle, failed, unsafe.Pointer(rawData))
	case FloatingGamepadTextInputDismissedID:
		handleRawResult[FloatingGamepadTextInputDismissed](apiHandle, failed, unsafe.Pointer(rawData))
	case FilterTextDictionaryChangedID:
		handleRawResult[FilterTextDictionaryChanged](apiHandle, failed, unsafe.Pointer(rawData))
	}
}
