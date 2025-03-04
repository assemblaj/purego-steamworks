package puregosteamworks

import (
	"unsafe"
)

func dispatchCallback(callback *CallbackMsg) {
	if GameServerActive {
		switch SteamCallbackID(callback.Callback) {
		case GSStatsReceivedID:
			handleRawData[GSStatsReceived](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSStatsStoredID:
			handleRawData[GSStatsStored](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSStatsUnloadedID:
			handleRawData[GSStatsUnloaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientApproveID:
			handleRawData[GSClientApprove](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientDenyID:
			handleRawData[GSClientDeny](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientKickID:
			handleRawData[GSClientKick](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientAchievementStatusID:
			handleRawData[GSClientAchievementStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSPolicyResponseID:
			handleRawData[GSPolicyResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSGameplayStatsID:
			handleRawData[GSGameplayStats](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSClientGroupStatusID:
			handleRawData[GSClientGroupStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case GSReputationID:
			handleRawData[GSReputation](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case AssociateWithClanResultID:
			handleRawData[AssociateWithClanResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		case ComputeNewPlayerCompatibilityResultID:
			handleRawData[ComputeNewPlayerCompatibilityResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
		}
	}
	switch SteamCallbackID(callback.Callback) {
	case DlcInstalledID:
		handleRawData[DlcInstalled](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case NewUrlLaunchParametersID:
		handleRawData[NewUrlLaunchParameters](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AppProofOfPurchaseKeyResponseID:
		handleRawData[AppProofOfPurchaseKeyResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FileDetailsResultID:
		handleRawData[FileDetailsResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case TimedTrialStatusID:
		handleRawData[TimedTrialStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case PersonaStateChangeID:
		handleRawData[PersonaStateChange](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameOverlayActivatedID:
		handleRawData[GameOverlayActivated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameServerChangeRequestedID:
		handleRawData[GameServerChangeRequested](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameLobbyJoinRequestedID:
		handleRawData[GameLobbyJoinRequested](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AvatarImageLoadedID:
		handleRawData[AvatarImageLoaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ClanOfficerListResponseID:
		handleRawData[ClanOfficerListResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendRichPresenceUpdateID:
		handleRawData[FriendRichPresenceUpdate](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameRichPresenceJoinRequestedID:
		handleRawData[GameRichPresenceJoinRequested](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedClanChatMsgID:
		handleRawData[GameConnectedClanChatMsg](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedChatJoinID:
		handleRawData[GameConnectedChatJoin](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedChatLeaveID:
		handleRawData[GameConnectedChatLeave](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DownloadClanActivityCountsResultID:
		handleRawData[DownloadClanActivityCountsResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case JoinClanChatRoomCompletionResultID:
		handleRawData[JoinClanChatRoomCompletionResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameConnectedFriendChatMsgID:
		handleRawData[GameConnectedFriendChatMsg](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsGetFollowerCountID:
		handleRawData[FriendsGetFollowerCount](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsIsFollowingID:
		handleRawData[FriendsIsFollowing](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FriendsEnumerateFollowingListID:
		handleRawData[FriendsEnumerateFollowingList](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SetPersonaNameResponseID:
		handleRawData[SetPersonaNameResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UnreadChatMessagesChangedID:
		handleRawData[UnreadChatMessagesChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case OverlayBrowserProtocolNavigationID:
		handleRawData[OverlayBrowserProtocolNavigation](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EquippedProfileItemsChangedID:
		handleRawData[EquippedProfileItemsChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EquippedProfileItemsID:
		handleRawData[EquippedProfileItems](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SearchForGameProgressCallbackID:
		handleRawData[SearchForGameProgressCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SearchForGameResultCallbackID:
		handleRawData[SearchForGameResultCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameProgressCallbackID:
		handleRawData[RequestPlayersForGameProgressCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameResultCallbackID:
		handleRawData[RequestPlayersForGameResultCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RequestPlayersForGameFinalResultCallbackID:
		handleRawData[RequestPlayersForGameFinalResultCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SubmitPlayerResultResultCallbackID:
		handleRawData[SubmitPlayerResultResultCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EndGameResultCallbackID:
		handleRawData[EndGameResultCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputDeviceConnectedID:
		handleRawData[SteamInputDeviceConnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputDeviceDisconnectedID:
		handleRawData[SteamInputDeviceDisconnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputConfigurationLoadedID:
		handleRawData[SteamInputConfigurationLoaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInputGamepadSlotChangeID:
		handleRawData[SteamInputGamepadSlotChange](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryResultReadyID:
		handleRawData[SteamInventoryResultReady](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryFullUpdateID:
		handleRawData[SteamInventoryFullUpdate](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryDefinitionUpdateID:
		handleRawData[SteamInventoryDefinitionUpdate](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryEligiblePromoItemDefIDsID:
		handleRawData[SteamInventoryEligiblePromoItemDefIDs](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryStartPurchaseResultID:
		handleRawData[SteamInventoryStartPurchaseResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamInventoryRequestPricesResultID:
		handleRawData[SteamInventoryRequestPricesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FavoritesListChangedID:
		handleRawData[FavoritesListChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyInviteID:
		handleRawData[LobbyInvite](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyEnterID:
		handleRawData[LobbyEnter](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyDataUpdateID:
		handleRawData[LobbyDataUpdate](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyChatUpdateID:
		handleRawData[LobbyChatUpdate](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyChatMsgID:
		handleRawData[LobbyChatMsg](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyGameCreatedID:
		handleRawData[LobbyGameCreated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyMatchListID:
		handleRawData[LobbyMatchList](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyKickedID:
		handleRawData[LobbyKicked](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LobbyCreatedID:
		handleRawData[LobbyCreated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case PSNGameBootInviteResultID:
		handleRawData[PSNGameBootInviteResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FavoritesListAccountsUpdatedID:
		handleRawData[FavoritesListAccountsUpdated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingMessagesSessionRequestID:
		handleRawData[SteamNetworkingMessagesSessionRequest](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingMessagesSessionFailedID:
		handleRawData[SteamNetworkingMessagesSessionFailed](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetAuthenticationStatusID:
		handleRawData[SteamNetAuthenticationStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetConnectionStatusChangedCallbackID:
		handleRawData[SteamNetConnectionStatusChangedCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamNetworkingFakeIPResultID:
		handleRawData[SteamNetworkingFakeIPResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRelayNetworkStatusID:
		handleRawData[SteamRelayNetworkStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case JoinPartyCallbackID:
		handleRawData[JoinPartyCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CreateBeaconCallbackID:
		handleRawData[CreateBeaconCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ReservationNotificationCallbackID:
		handleRawData[ReservationNotificationCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ChangeNumOpenSlotsCallbackID:
		handleRawData[ChangeNumOpenSlotsCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AvailableBeaconLocationsUpdatedID:
		handleRawData[AvailableBeaconLocationsUpdated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ActiveBeaconsUpdatedID:
		handleRawData[ActiveBeaconsUpdated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlaySessionConnectedID:
		handleRawData[SteamRemotePlaySessionConnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlaySessionDisconnectedID:
		handleRawData[SteamRemotePlaySessionDisconnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamRemotePlayTogetherGuestInviteID:
		handleRawData[SteamRemotePlayTogetherGuestInvite](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileShareResultID:
		handleRawData[RemoteStorageFileShareResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishFileResultID:
		handleRawData[RemoteStoragePublishFileResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageDeletePublishedFileResultID:
		handleRawData[RemoteStorageDeletePublishedFileResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserPublishedFilesResultID:
		handleRawData[RemoteStorageEnumerateUserPublishedFilesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageSubscribePublishedFileResultID:
		handleRawData[RemoteStorageSubscribePublishedFileResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserSubscribedFilesResultID:
		handleRawData[RemoteStorageEnumerateUserSubscribedFilesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUnsubscribePublishedFileResultID:
		handleRawData[RemoteStorageUnsubscribePublishedFileResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUpdatePublishedFileResultID:
		handleRawData[RemoteStorageUpdatePublishedFileResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageDownloadUGCResultID:
		handleRawData[RemoteStorageDownloadUGCResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageGetPublishedFileDetailsResultID:
		handleRawData[RemoteStorageGetPublishedFileDetailsResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateWorkshopFilesResultID:
		handleRawData[RemoteStorageEnumerateWorkshopFilesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageGetPublishedItemVoteDetailsResultID:
		handleRawData[RemoteStorageGetPublishedItemVoteDetailsResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileSubscribedID:
		handleRawData[RemoteStoragePublishedFileSubscribed](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileUnsubscribedID:
		handleRawData[RemoteStoragePublishedFileUnsubscribed](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileDeletedID:
		handleRawData[RemoteStoragePublishedFileDeleted](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUpdateUserPublishedItemVoteResultID:
		handleRawData[RemoteStorageUpdateUserPublishedItemVoteResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageUserVoteDetailsID:
		handleRawData[RemoteStorageUserVoteDetails](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumerateUserSharedWorkshopFilesResultID:
		handleRawData[RemoteStorageEnumerateUserSharedWorkshopFilesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageSetUserPublishedFileActionResultID:
		handleRawData[RemoteStorageSetUserPublishedFileActionResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageEnumeratePublishedFilesByUserActionResultID:
		handleRawData[RemoteStorageEnumeratePublishedFilesByUserActionResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishFileProgressID:
		handleRawData[RemoteStoragePublishFileProgress](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStoragePublishedFileUpdatedID:
		handleRawData[RemoteStoragePublishedFileUpdated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileWriteAsyncCompleteID:
		handleRawData[RemoteStorageFileWriteAsyncComplete](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageFileReadAsyncCompleteID:
		handleRawData[RemoteStorageFileReadAsyncComplete](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoteStorageLocalFileChangeID:
		handleRawData[RemoteStorageLocalFileChange](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ScreenshotReadyID:
		handleRawData[ScreenshotReady](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ScreenshotRequestedID:
		handleRawData[ScreenshotRequested](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamUGCQueryCompletedID:
		handleRawData[SteamUGCQueryCompleted](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamUGCRequestUGCDetailsResultID:
		handleRawData[SteamUGCRequestUGCDetailsResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CreateItemResultID:
		handleRawData[CreateItemResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SubmitItemUpdateResultID:
		handleRawData[SubmitItemUpdateResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ItemInstalledID:
		handleRawData[ItemInstalled](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DownloadItemResultID:
		handleRawData[DownloadItemResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserFavoriteItemsListChangedID:
		handleRawData[UserFavoriteItemsListChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SetUserItemVoteResultID:
		handleRawData[SetUserItemVoteResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetUserItemVoteResultID:
		handleRawData[GetUserItemVoteResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case StartPlaytimeTrackingResultID:
		handleRawData[StartPlaytimeTrackingResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case StopPlaytimeTrackingResultID:
		handleRawData[StopPlaytimeTrackingResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AddUGCDependencyResultID:
		handleRawData[AddUGCDependencyResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoveUGCDependencyResultID:
		handleRawData[RemoveUGCDependencyResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AddAppDependencyResultID:
		handleRawData[AddAppDependencyResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case RemoveAppDependencyResultID:
		handleRawData[RemoveAppDependencyResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetAppDependenciesResultID:
		handleRawData[GetAppDependenciesResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DeleteItemResultID:
		handleRawData[DeleteItemResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserSubscribedItemsListChangedID:
		handleRawData[UserSubscribedItemsListChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case WorkshopEULAStatusID:
		handleRawData[WorkshopEULAStatus](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamServersConnectedID:
		handleRawData[SteamServersConnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamServerConnectFailureID:
		handleRawData[SteamServerConnectFailure](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamServersDisconnectedID:
		handleRawData[SteamServersDisconnected](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ClientGameServerDenyID:
		handleRawData[ClientGameServerDeny](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case IPCFailureID:
		handleRawData[IPCFailure](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LicensesUpdatedID:
		handleRawData[LicensesUpdated](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case ValidateAuthTicketResponseID:
		handleRawData[ValidateAuthTicketResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case MicroTxnAuthorizationResponseID:
		handleRawData[MicroTxnAuthorizationResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case EncryptedAppTicketResponseID:
		handleRawData[EncryptedAppTicketResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetAuthSessionTicketResponseID:
		handleRawData[GetAuthSessionTicketResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GameWebCallbackID:
		handleRawData[GameWebCallback](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case StoreAuthURLResponseID:
		handleRawData[StoreAuthURLResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case MarketEligibilityResponseID:
		handleRawData[MarketEligibilityResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case DurationControlID:
		handleRawData[DurationControl](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GetTicketForWebApiResponseID:
		handleRawData[GetTicketForWebApiResponse](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserStatsReceivedID:
		handleRawData[UserStatsReceived](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserStatsStoredID:
		handleRawData[UserStatsStored](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserAchievementStoredID:
		handleRawData[UserAchievementStored](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LeaderboardFindResultID:
		handleRawData[LeaderboardFindResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LeaderboardScoresDownloadedID:
		handleRawData[LeaderboardScoresDownloaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LeaderboardScoreUploadedID:
		handleRawData[LeaderboardScoreUploaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case NumberOfCurrentPlayersID:
		handleRawData[NumberOfCurrentPlayers](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserStatsUnloadedID:
		handleRawData[UserStatsUnloaded](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case UserAchievementIconFetchedID:
		handleRawData[UserAchievementIconFetched](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GlobalAchievementPercentagesReadyID:
		handleRawData[GlobalAchievementPercentagesReady](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LeaderboardUGCSetID:
		handleRawData[LeaderboardUGCSet](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GlobalStatsReceivedID:
		handleRawData[GlobalStatsReceived](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case IPCountryID:
		handleRawData[IPCountry](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case LowBatteryPowerID:
		handleRawData[LowBatteryPower](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case SteamShutdownID:
		handleRawData[SteamShutdown](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case CheckFileSignatureID:
		handleRawData[CheckFileSignatureResult](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case GamepadTextInputDismissedID:
		handleRawData[GamepadTextInputDismissed](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case AppResumingFromSuspendID:
		handleRawData[AppResumingFromSuspend](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FloatingGamepadTextInputDismissedID:
		handleRawData[FloatingGamepadTextInputDismissed](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	case FilterTextDictionaryChangedID:
		handleRawData[FilterTextDictionaryChanged](SteamCallbackID(callback.Callback), unsafe.Pointer(callback.ParamData))
	}
}

func dispatchCallResult(apiHandle SteamAPICall, failed bool, rawData uintptr, callbackID SteamCallbackID) {
	if GameServerActive {
		switch SteamCallbackID(callbackID) {
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
	switch SteamCallbackID(callbackID) {
	case DlcInstalledID:
		handleRawResult[DlcInstalled](apiHandle, failed, unsafe.Pointer(rawData))
	case NewUrlLaunchParametersID:
		handleRawResult[NewUrlLaunchParameters](apiHandle, failed, unsafe.Pointer(rawData))
	case AppProofOfPurchaseKeyResponseID:
		handleRawResult[AppProofOfPurchaseKeyResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case FileDetailsResultID:
		handleRawResult[FileDetailsResult](apiHandle, failed, unsafe.Pointer(rawData))
	case TimedTrialStatusID:
		handleRawResult[TimedTrialStatus](apiHandle, failed, unsafe.Pointer(rawData))
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
	case SteamServersConnectedID:
		handleRawResult[SteamServersConnected](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamServerConnectFailureID:
		handleRawResult[SteamServerConnectFailure](apiHandle, failed, unsafe.Pointer(rawData))
	case SteamServersDisconnectedID:
		handleRawResult[SteamServersDisconnected](apiHandle, failed, unsafe.Pointer(rawData))
	case ClientGameServerDenyID:
		handleRawResult[ClientGameServerDeny](apiHandle, failed, unsafe.Pointer(rawData))
	case IPCFailureID:
		handleRawResult[IPCFailure](apiHandle, failed, unsafe.Pointer(rawData))
	case LicensesUpdatedID:
		handleRawResult[LicensesUpdated](apiHandle, failed, unsafe.Pointer(rawData))
	case ValidateAuthTicketResponseID:
		handleRawResult[ValidateAuthTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case MicroTxnAuthorizationResponseID:
		handleRawResult[MicroTxnAuthorizationResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case EncryptedAppTicketResponseID:
		handleRawResult[EncryptedAppTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case GetAuthSessionTicketResponseID:
		handleRawResult[GetAuthSessionTicketResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case GameWebCallbackID:
		handleRawResult[GameWebCallback](apiHandle, failed, unsafe.Pointer(rawData))
	case StoreAuthURLResponseID:
		handleRawResult[StoreAuthURLResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case MarketEligibilityResponseID:
		handleRawResult[MarketEligibilityResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case DurationControlID:
		handleRawResult[DurationControl](apiHandle, failed, unsafe.Pointer(rawData))
	case GetTicketForWebApiResponseID:
		handleRawResult[GetTicketForWebApiResponse](apiHandle, failed, unsafe.Pointer(rawData))
	case UserStatsReceivedID:
		handleRawResult[UserStatsReceived](apiHandle, failed, unsafe.Pointer(rawData))
	case UserStatsStoredID:
		handleRawResult[UserStatsStored](apiHandle, failed, unsafe.Pointer(rawData))
	case UserAchievementStoredID:
		handleRawResult[UserAchievementStored](apiHandle, failed, unsafe.Pointer(rawData))
	case LeaderboardFindResultID:
		handleRawResult[LeaderboardFindResult](apiHandle, failed, unsafe.Pointer(rawData))
	case LeaderboardScoresDownloadedID:
		handleRawResult[LeaderboardScoresDownloaded](apiHandle, failed, unsafe.Pointer(rawData))
	case LeaderboardScoreUploadedID:
		handleRawResult[LeaderboardScoreUploaded](apiHandle, failed, unsafe.Pointer(rawData))
	case NumberOfCurrentPlayersID:
		handleRawResult[NumberOfCurrentPlayers](apiHandle, failed, unsafe.Pointer(rawData))
	case UserStatsUnloadedID:
		handleRawResult[UserStatsUnloaded](apiHandle, failed, unsafe.Pointer(rawData))
	case UserAchievementIconFetchedID:
		handleRawResult[UserAchievementIconFetched](apiHandle, failed, unsafe.Pointer(rawData))
	case GlobalAchievementPercentagesReadyID:
		handleRawResult[GlobalAchievementPercentagesReady](apiHandle, failed, unsafe.Pointer(rawData))
	case LeaderboardUGCSetID:
		handleRawResult[LeaderboardUGCSet](apiHandle, failed, unsafe.Pointer(rawData))
	case GlobalStatsReceivedID:
		handleRawResult[GlobalStatsReceived](apiHandle, failed, unsafe.Pointer(rawData))
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
