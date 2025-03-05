package puregosteamworks

type ISteamApps interface {
	BIsSubscribed() bool
	BIsLowViolence() bool
	BIsCybercafe() bool
	BIsVACBanned() bool
	GetCurrentGameLanguage() string
	GetAvailableGameLanguages() []string
	BIsSubscribedApp(appID AppId_t) bool
	BIsDlcInstalled(appID AppId_t) bool
	GetEarliestPurchaseUnixTime(appID AppId_t) uint32
	BIsSubscribedFromFreeWeekend() bool
	GetDLCCount() int32
	BGetDLCDataByIndex(dlcIndex int32) (appID AppId_t, available bool, name string, success bool)
	InstallDLC(appID AppId_t)
	UninstallDLC(appID AppId_t)
	RequestAppProofOfPurchaseKey(appID AppId_t)
	GetCurrentBetaName() (name string, success bool)
	MarkContentCorrupt(missingFilesOnly bool) bool
	GetInstalledDepots(appID AppId_t, maxDepots uint32) []DepotId
	GetAppInstallDir(appID AppId_t) string
	BIsAppInstalled(appID AppId_t) bool
	GetAppOwner() Uint64SteamID
	GetLaunchQueryParam(key string) string
	GetDlcDownloadProgress(appID AppId_t) (bytesDownloaded uint64, bytesTotal uint64, success bool)
	GetAppBuildId() int32
	RequestAllProofOfPurchaseKeys()
	GetFileDetails(pszFileName string) CallResult[FileDetailsResult]
	GetLaunchCommandLine() string
	BIsSubscribedFromFamilySharing() bool
	BIsTimedTrial() (inTimeTrial bool, secondsAllowed uint32, secondsPlayed uint32)
	SetDlcContext(appID AppId_t) bool
	GetNumBetas() (available int32, private int32, totalAppBranches int32)
	GetBetaInfo(betaIndex int32, betaNameSize int32, descriptionSize int32) (punFlags uint32, punBuildID uint32, pchBetaName string, pchDescription string, success bool)
	SetActiveBeta(pchBetaName string) bool
}

type ISteamFriends interface {
	GetPersonaName() string
	SetPersonaName(personaName string) CallResult[SetPersonaNameResponse]
	GetPersonaState() EPersonaState
	GetFriendCount(friendFlags int32) int32
	GetFriendByIndex(friend int32, friendFlags int32) Uint64SteamID
	GetFriendRelationship(friendSteamID Uint64SteamID) EFriendRelationship
	GetFriendPersonaState(friendSteamID Uint64SteamID) EPersonaState
	GetFriendPersonaName(friendSteamID Uint64SteamID) string
	GetFriendGamePlayed(friendSteamID Uint64SteamID) (FriendGameInfo, bool)
	GetFriendPersonaNameHistory(friendSteamID Uint64SteamID, personaName int32) string
	GetFriendSteamLevel(friendSteamID Uint64SteamID) int32
	GetFriendsGroupCount() int32
	GetFriendsGroupIDByIndex(index int32) FriendsGroupID
	GetFriendsGroupName(friendsGroupID FriendsGroupID) string
	GetFriendsGroupMembersCount(friendsGroupID FriendsGroupID) int32
	GetFriendsGroupMembersList(friendsGroupID FriendsGroupID) (members []CSteamID)
	HasFriend(friendSteamID Uint64SteamID, friendFlags int32) bool
	GetClanCount() int32
	GetClanByIndex(index int32) Uint64SteamID
	GetClanName(clanSteamID Uint64SteamID) string
	GetClanTag(clanSteamID Uint64SteamID) string
	GetClanActivityCounts(clanSteamID Uint64SteamID) (online int32, inGame int32, chatting int32, success bool)
	DownloadClanActivityCounts(clansToRequest int32) ([]CSteamID, CallResult[DownloadClanActivityCountsResult])
	GetFriendCountFromSource(sourceSteamID Uint64SteamID) int32
	GetFriendFromSourceByIndex(sourceSteamID Uint64SteamID, friend int32) Uint64SteamID
	IsUserInSource(userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool
	SetInGameVoiceSpeaking(userSteamID Uint64SteamID, speaking bool)
	ActivateGameOverlay(dialog string)
	ActivateGameOverlayToUser(dialog string, steamID Uint64SteamID)
	ActivateGameOverlayToWebPage(URL string, mode EActivateGameOverlayToWebPageMode)
	ActivateGameOverlayToStore(appID AppId_t, flag EOverlayToStoreFlag)
	SetPlayedWith(userPlayedWithSteamID Uint64SteamID)
	ActivateGameOverlayInviteDialog(lobbySteamID Uint64SteamID)
	GetSmallFriendAvatar(friendSteamID Uint64SteamID) int32
	GetMediumFriendAvatar(friendSteamID Uint64SteamID) int32
	GetLargeFriendAvatar(friendSteamID Uint64SteamID) int32
	RequestUserInformation(userSteamID Uint64SteamID, requireNameOnly bool) bool
	RequestClanOfficerList(clanSteamID Uint64SteamID) CallResult[ClanOfficerListResponse]
	GetClanOwner(clanSteamID Uint64SteamID) Uint64SteamID
	GetClanOfficerCount(clanSteamID Uint64SteamID) int32
	GetClanOfficerByIndex(clanSteamID Uint64SteamID, officer int32) Uint64SteamID
	GetUserRestrictions() uint32
	SetRichPresence(key string, value string) bool
	ClearRichPresence()
	GetFriendRichPresence(friendSteamID Uint64SteamID, key string) []byte
	GetFriendRichPresenceKeyCount(friendSteamID Uint64SteamID) int32
	GetFriendRichPresenceKeyByIndex(friendSteamID Uint64SteamID, keyIndex int32) []byte
	RequestFriendRichPresence(friendSteamID Uint64SteamID)
	InviteUserToGame(friendSteamID Uint64SteamID, connectString string) bool
	GetCoplayFriendCount() int32
	GetCoplayFriend(coplayFriendIndex int32) Uint64SteamID
	GetFriendCoplayTime(friendSteamID Uint64SteamID) int32
	GetFriendCoplayGame(friendSteamID Uint64SteamID) AppId_t
	JoinClanChatRoom(clanSteamID Uint64SteamID) CallResult[JoinClanChatRoomCompletionResult]
	LeaveClanChatRoom(clanSteamID Uint64SteamID) bool
	GetClanChatMemberCount(clanSteamID Uint64SteamID) int32
	GetChatMemberByIndex(clanSteamID Uint64SteamID, userIndex int32) Uint64SteamID
	SendClanChatMessage(clanChatSteamID Uint64SteamID, text string) bool
	GetClanChatMessage(clanChatSteamID Uint64SteamID, messageIndex int32, maxTextSize int32) (chatMsg []byte, chatEntryType EChatEntryType, chatter CSteamID, result int32)
	IsClanChatAdmin(clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool
	IsClanChatWindowOpenInSteam(clanChatSteamID Uint64SteamID) bool
	OpenClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool
	CloseClanChatWindowInSteam(clanChatSteamID Uint64SteamID) bool
	SetListenForFriendsMessages(interceptEnabled bool) bool
	ReplyToFriendMessage(friendSteamID Uint64SteamID, msgToSend string) bool
	GetFriendMessage(friendSteamID Uint64SteamID, messageID int32, messageMaxSize int32) (data []byte, chatEntryType EChatEntryType, result int32)
	GetFollowerCount(steamID Uint64SteamID) CallResult[FriendsGetFollowerCount]
	IsFollowing(steamID Uint64SteamID) CallResult[FriendsIsFollowing]
	EnumerateFollowingList(startIndex uint32) CallResult[FriendsEnumerateFollowingList]
	IsClanPublic(clanSteamID Uint64SteamID) bool
	IsClanOfficialGameGroup(clanSteamID Uint64SteamID) bool
	GetNumChatsWithUnreadPriorityMessages() int32
	ActivateGameOverlayRemotePlayTogetherInviteDialog(lobbySteamID Uint64SteamID)
	RegisterProtocolInOverlayBrowser(protocol string) bool
	ActivateGameOverlayInviteDialogConnectString(connectString string)
	RequestEquippedProfileItems(steamID Uint64SteamID) CallResult[EquippedProfileItems]
	BHasEquippedProfileItem(steamID Uint64SteamID, itemType ECommunityProfileItemType) bool
	GetProfileItemPropertyString(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) string
	GetProfileItemPropertyUint(steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) uint32
}

type ISteamClient interface {
	CreateSteamPipe() HSteamPipe
	BReleaseSteamPipe(hSteamPipe HSteamPipe) bool
	ConnectToGlobalUser(hSteamPipe HSteamPipe) HSteamUser
	CreateLocalUser(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser
	ReleaseUser(hSteamPipe HSteamPipe, hUser HSteamUser)
	SetLocalIPBinding(unIP *SteamIPAddress, usPort uint16)
	GetIPCCallCount() uint32
	SetWarningMessageHook(pFunction SteamAPIWarningMessageHook)
	BShutdownIfAllPipesClosed() bool
}

type ISteamGameSearch interface {
	AddGameSearchParams(keyToFind string, valuesToFind string) EGameSearchErrorCode
	SearchForGameWithLobby(lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode
	SearchForGameSolo(playerMin int32, playerMax int32) EGameSearchErrorCode
	AcceptGame() EGameSearchErrorCode
	DeclineGame() EGameSearchErrorCode
	RetrieveConnectionDetails(hostSteamID Uint64SteamID, connectionDetailsSize int32) (string, EGameSearchErrorCode)
	EndGameSearch() EGameSearchErrorCode
	SetGameHostParams(key string, value string) EGameSearchErrorCode
	SetConnectionDetails(connectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode
	RequestPlayersForGame(playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode
	HostConfirmGameStart(uniqueGameID uint64) EGameSearchErrorCode
	CancelRequestPlayersForGame() EGameSearchErrorCode
	SubmitPlayerResult(uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode
	EndGame(uniqueGameID uint64) EGameSearchErrorCode
}

type ISteamGameServer interface {
	SetProduct(product string)
	SetGameDescription(gameDescription string)
	SetModDir(modDir string)
	SetDedicatedServer(dedicated bool)
	LogOn(token string)
	LogOnAnonymous()
	LogOff()
	BLoggedOn() bool
	BSecure() bool
	GetSteamID() Uint64SteamID
	WasRestartRequested() bool
	SetMaxPlayerCount(playersMax int32)
	SetBotPlayerCount(botplayers int32)
	SetServerName(serverName string)
	SetMapName(mapName string)
	SetPasswordProtected(passwordProtected bool)
	SetSpectatorPort(spectatorPort uint16)
	SetSpectatorServerName(spectatorServerName string)
	ClearAllKeyValues()
	SetKeyValue(key string, value string)
	SetGameTags(gameTags string)
	SetGameData(gameData string)
	SetRegion(region string)
	SetAdvertiseServerActive(active bool)
	GetAuthSessionTicket(maxTicket int32, snid *SteamNetworkingIdentity) ([]byte, HAuthTicket)
	BeginAuthSession(authTicket []byte, steamID Uint64SteamID) EBeginAuthSessionResult
	EndAuthSession(steamID Uint64SteamID)
	CancelAuthTicket(authTicket HAuthTicket)
	UserHasLicenseForApp(steamID Uint64SteamID, AppId AppId_t) EUserHasLicenseForAppResult
	RequestUserGroupStatus(userSteamID Uint64SteamID, groupSteamID Uint64SteamID) bool
	GetPublicIP() SteamIPAddress
	HandleIncomingPacket(data int32, srcIP uint32, srcPort uint16) ([]byte, bool)
	GetNextOutgoingPacket(maxOut int32) (packet []byte, addr uint32, port uint16, result int32)
	AssociateWithClan(clanSteamID Uint64SteamID) CallResult[AssociateWithClanResult]
	ComputeNewPlayerCompatibility(newPlayerSteamID Uint64SteamID) CallResult[ComputeNewPlayerCompatibilityResult]
	CreateUnauthenticatedUserConnection() Uint64SteamID
	BUpdateUserData(user Uint64SteamID, playerName string, score uint32) bool
}

type ISteamGameServerStats interface {
	RequestUserStats(userSteamID Uint64SteamID) CallResult[GSStatsReceived]
	GetUserStat(userSteamID Uint64SteamID, name string) (int32, bool)
	GetUserStatFloat(userSteamID Uint64SteamID, name string) (float32, bool)
	GetUserAchievement(userSteamID Uint64SteamID, name string) (bool, bool)
	SetUserStat(userSteamID Uint64SteamID, name string, nData int32) bool
	SetUserStatFloat(userSteamID Uint64SteamID, name string, fData float32) bool
	UpdateUserAvgRateStat(userSteamID Uint64SteamID, name string, countThisSession float32, sessionLength float64) bool
	SetUserAchievement(userSteamID Uint64SteamID, name string) bool
	ClearUserAchievement(userSteamID Uint64SteamID, name string) bool
	StoreUserStats(userSteamID Uint64SteamID) CallResult[GSStatsStored]
}

type ISteamInput interface {
	Init(explicitlyCallRunFrame bool) bool
	Shutdown() bool
	SetInputActionManifestFilePath(inputActionManifestAbsolutePath string) bool
	RunFrame()
	RunFrameEx(reservedValue bool)
	BWaitForData(waitForever bool, timeout uint32) bool
	BNewDataAvailable() bool
	GetConnectedControllers() []InputHandle_t
	EnableDeviceCallbacks()
	EnableActionEventCallbacks(pCallback SteamInputActionEventCallbackPointer)
	GetActionSetHandle(pszActionSetName string) InputActionSetHandle
	ActivateActionSet(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle)
	GetCurrentActionSet(inputHandle InputHandle_t) InputActionSetHandle
	ActivateActionSetLayer(inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle)
	DeactivateActionSetLayer(inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle)
	DeactivateAllActionSetLayers(inputHandle InputHandle_t)
	GetActiveActionSetLayers(inputHandle InputHandle_t) []InputActionSetHandle
	GetDigitalActionHandle(actionName string) InputDigitalActionHandle
	GetDigitalActionData(inputHandle InputHandle_t, digitalActionHandle InputDigitalActionHandle) InputDigitalActionData
	GetDigitalActionOrigins(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle) []EInputActionOrigin
	GetStringForDigitalActionName(actionHandle InputDigitalActionHandle) string
	GetAnalogActionHandle(actionName string) InputAnalogActionHandle
	GetAnalogActionData(inputHandle InputHandle_t, analogActionHandle InputAnalogActionHandle) InputAnalogActionData
	GetAnalogActionOrigins(inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle) []EInputActionOrigin
	GetGlyphPNGForActionOrigin(origin EInputActionOrigin, size ESteamInputGlyphSize, flags uint32) []byte
	GetGlyphSVGForActionOrigin(origin EInputActionOrigin, flags uint32) []byte
	GetGlyphForActionOrigin_Legacy(origin EInputActionOrigin) []byte
	GetStringForActionOrigin(origin EInputActionOrigin) string
	GetStringForAnalogActionName(actionHandle InputAnalogActionHandle) string
	StopAnalogActionMomentum(inputHandle InputHandle_t, action InputAnalogActionHandle)
	GetMotionData(inputHandle InputHandle_t) InputMotionData
	TriggerVibration(inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16)
	TriggerVibrationExtended(inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16)
	TriggerSimpleHapticEvent(inputHandle InputHandle_t, hapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte)
	SetLEDColor(inputHandle InputHandle_t, colorR uint8, colorG uint8, colorB uint8, flags uint)
	Legacy_TriggerHapticPulse(inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16)
	Legacy_TriggerRepeatedHapticPulse(inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint)
	ShowBindingPanel(inputHandle InputHandle_t) bool
	GetInputTypeForHandle(inputHandle InputHandle_t) ESteamInputType
	GetControllerForGamepadIndex(index int32) InputHandle_t
	GetGamepadIndexForController(inputHandle InputHandle_t) int32
	GetStringForXboxOrigin(origin EXboxOrigin) string
	GetGlyphForXboxOrigin(origin EXboxOrigin) []byte
	GetActionOriginFromXboxOrigin(inputHandle InputHandle_t, origin EXboxOrigin) EInputActionOrigin
	TranslateActionOrigin(destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin
	GetDeviceBindingRevision(inputHandle InputHandle_t) (major int32, minor int32, foundController bool)
	GetRemotePlaySessionID(inputHandle InputHandle_t) uint32
	GetSessionInputConfigurationSettings() uint16
	SetDualSenseTriggerEffect(inputHandle InputHandle_t, param *ScePadTriggerEffectParam)
}

type ISteamInventory interface {
	GetResultStatus(resultHandle SteamInventoryResult) EResult
	GetResultItems(resultHandle SteamInventoryResult, outItemsArraySize uint32) ([]SteamItemDetails, bool)
	GetResultItemProperty(resultHandle SteamInventoryResult, itemIndex uint32, propertyName string, valueBufferSizeOut uint32) ([]byte, bool)
	GetResultTimestamp(resultHandle SteamInventoryResult) uint32
	CheckResultSteamID(resultHandle SteamInventoryResult, expectedSteamID Uint64SteamID) bool
	DestroyResult(resultHandle SteamInventoryResult)
	GetAllItems() (resultHandle SteamInventoryResult, success bool)
	GetItemsByID(countInstanceIDs uint32) (resultHandle SteamInventoryResult, instanceIDs []SteamItemInstanceID, result bool)
	SerializeResult(resultHandle SteamInventoryResult, outBufferSize uint32) ([]byte, bool)
	DeserializeResult(bufferSize uint32, bRESERVEDMUSTBEFALSE bool) (resultHandle SteamInventoryResult, buffer []byte, success bool)
	GenerateItems(arrayItemDefs []SteamItemDef, arrayQuantity []uint32, arrayLength uint32) (resultHandle SteamInventoryResult, success bool)
	GrantPromoItems() (resultHandle SteamInventoryResult, success bool)
	AddPromoItem(itemDef SteamItemDef) (resultHandle SteamInventoryResult, success bool)
	AddPromoItems(arrayItemDefs []SteamItemDef, arrayLength uint32) (resultHandle SteamInventoryResult, success bool)
	ConsumeItem(itemConsume SteamItemInstanceID, quantity uint32) (resultHandle SteamInventoryResult, success bool)
	ExchangeItems(arrayGenerate []SteamItemDef, arrayGenerateQuantity []uint32, arrayGenerateLength uint32, arrayDestroy []SteamItemInstanceID, arrayDestroyQuantity []uint32, arrayDestroyLength uint32) (resultHandle SteamInventoryResult, success bool)
	TransferItemQuantity(itemIdSource SteamItemInstanceID, quantity uint32, itemIdDest SteamItemInstanceID) (pResultHandle SteamInventoryResult, success bool)
	TriggerItemDrop(dropListDefinition SteamItemDef) (resultHandle SteamInventoryResult, success bool)
	LoadItemDefinitions() bool
	GetItemDefinitionIDs(itemDefIDsArraySize uint32) ([]SteamItemDef, bool)
	GetItemDefinitionProperty(definition SteamItemDef, propertyName string, valueBufferSizeOut uint32) ([]byte, bool)
	RequestEligiblePromoItemDefinitionsIDs(steamID Uint64SteamID) CallResult[SteamInventoryEligiblePromoItemDefIDs]
	GetEligiblePromoItemDefinitionIDs(steamID Uint64SteamID, itemDefIDsArraySize uint32) ([]SteamItemDef, bool)
	StartPurchase(arrayItemDefs []SteamItemDef, punArrayQuantity []uint32, arrayLength uint32) CallResult[SteamInventoryStartPurchaseResult]
	RequestPrices() CallResult[SteamInventoryRequestPricesResult]
	GetNumItemsWithPrices() uint32
	GetItemsWithPrices(arrayLength uint32) ([]SteamItemDef, []uint64, []uint64, bool)
	GetItemPrice(definition SteamItemDef) (currentPrice uint64, basePrice uint64, succes bool)
	StartUpdateProperties() SteamInventoryUpdateHandle
	RemoveProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string) bool
	SetProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, pchPropertyValue string) bool
	SetPropertyBool(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value bool) bool
	SetPropertyInt64y(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value int64) bool
	SetPropertyFloat(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value float32) bool
	SubmitUpdateProperties(handle SteamInventoryUpdateHandle) (resultHandle SteamInventoryResult, succes bool)
	InspectItem(itemToken string) (resultHandle SteamInventoryResult, success bool)
}

type ISteamMatchmaking interface {
	GetFavoriteGameCount() int32
	GetFavoriteGame(gameIndex int32) (appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32, success bool)
	AddFavoriteGame(appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32, lastPlayedOnServerTime uint32) int32
	RemoveFavoriteGame(appID AppId_t, IP uint32, connPort uint16, queryPort uint16, flags uint32) bool
	RequestLobbyList() CallResult[LobbyMatchList]
	AddRequestLobbyListStringFilter(keyToMatch string, valueToMatch string, comparisonType ELobbyComparison)
	AddRequestLobbyListNumericalFilter(keyToMatch string, valueToMatch int32, comparisonType ELobbyComparison)
	AddRequestLobbyListNearValueFilter(keyToMatch string, valueToBeCloseTo int32)
	AddRequestLobbyListFilterSlotsAvailable(slotsAvailable int32)
	AddRequestLobbyListDistanceFilter(lobbyDistanceFilter ELobbyDistanceFilter)
	AddRequestLobbyListResultCountFilter(cMaxResults int32)
	AddRequestLobbyListCompatibleMembersFilter(lobbySteamID Uint64SteamID)
	GetLobbyByIndex(iLobby int32) Uint64SteamID
	CreateLobby(lobbyType ELobbyType, maxMembers int32) CallResult[LobbyCreated]
	JoinLobby(lobbySteamID Uint64SteamID) CallResult[LobbyEnter]
	LeaveLobby(lobbySteamID Uint64SteamID)
	InviteUserToLobby(lobbySteamID Uint64SteamID, inviteeSteamID Uint64SteamID) bool
	GetNumLobbyMembers(lobbySteamID Uint64SteamID) int32
	GetLobbyMemberByIndex(lobbySteamID Uint64SteamID, memberIndex int32) Uint64SteamID
	GetLobbyData(lobbySteamID Uint64SteamID, key string) []byte
	SetLobbyData(lobbySteamID Uint64SteamID, key string, value string) bool
	GetLobbyDataCount(lobbySteamID Uint64SteamID) int32
	GetLobbyDataByIndex(lobbySteamID Uint64SteamID, lobbyDataIndex int32, keyBufferSize int32, valueBufferSize int32) (key []byte, value []byte, success bool)
	DeleteLobbyData(lobbySteamID Uint64SteamID, key string) bool
	GetLobbyMemberData(lobbySteamID Uint64SteamID, userSteamID Uint64SteamID, key string) []byte
	SetLobbyMemberData(lobbySteamID Uint64SteamID, key string, value string)
	SendLobbyChatMsg(lobbySteamID Uint64SteamID, msgBody []byte) bool
	GetLobbyChatEntry(lobbySteamID Uint64SteamID, chatID int32, dataSize int32) (userSteamID CSteamID, data []byte, chatEntryType EChatEntryType)
	RequestLobbyData(lobbySteamID Uint64SteamID) bool
	SetLobbyGameServer(lobbySteamID Uint64SteamID, gameServerIP uint32, gameServerPort uint16, gameServerSteamID Uint64SteamID)
	GetLobbyGameServer(lobbySteamID Uint64SteamID) (gameServerIP uint32, gameServerPort uint16, gameServerSteamID CSteamID, success bool)
	SetLobbyMemberLimit(lobbySteamID Uint64SteamID, cMaxMembers int32) bool
	GetLobbyMemberLimit(lobbySteamID Uint64SteamID) int32
	SetLobbyType(lobbySteamID Uint64SteamID, lobbyType ELobbyType) bool
	SetLobbyJoinable(lobbySteamID Uint64SteamID, joinableLobby bool) bool
	GetLobbyOwner(lobbySteamID Uint64SteamID) Uint64SteamID
	SetLobbyOwner(lobbySteamID Uint64SteamID, newUserSteamID Uint64SteamID) bool
	SetLinkedLobby(lobbySteamID Uint64SteamID, dependentLobbySteamID Uint64SteamID) bool
}

type ISteamMatchmakingServers interface {
	RequestInternetServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	RequestLANServerList(App AppId_t) (RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	RequestFriendsServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	RequestFavoritesServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	RequestHistoryServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	RequestSpectatorServerList(App AppId_t, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest)
	ReleaseRequest(ServerListRequest HServerListRequest)
	GetServerDetails(Request HServerListRequest, Server int32) *GameServerItem
	CancelQuery(Request HServerListRequest)
	RefreshQuery(Request HServerListRequest)
	IsRefreshing(Request HServerListRequest) bool
	GetServerCount(Request HServerListRequest) int32
	RefreshServer(Request HServerListRequest, Server int32)
	PingServer(IP uint32, Port uint16) (RequestServersResponse ISteamMatchmakingPingResponse, query HServerQuery)
	PlayerDetails(IP uint32, Port uint16) (RequestServersResponse MatchmakingPlayersResponse, query HServerQuery)
	ServerRules(IP uint32, Port uint16) (RequestServersResponse MatchmakingRulesResponse, query HServerQuery)
	CancelServerQuery(serverQuery HServerQuery)
}

type ISteamNetworkingMessages interface {
	SendMessageToUser(remoteIdentity SteamNetworkingIdentity, Data []byte, SendFlags int32, RemoteChannel int32) EResult
	ReceiveMessagesOnChannel(LocalChannel int32, MaxMessages int32) []SteamNetworkingMessage
	AcceptSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool
	CloseSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool
	CloseChannelWithUser(remoteIdentity SteamNetworkingIdentity, LocalChannel int32) bool
	GetSessionConnectionInfo(remoteIdentity SteamNetworkingIdentity) (ConnectionInfo SteamNetConnectionInfo, QuickStatus SteamNetConnectionRealTimeStatus, state ESteamNetworkingConnectionState)
}

type ISteamNetworkingSockets interface {
	CreateListenSocketIP(localAddress SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamListenSocket
	ConnectByIPAddress(address SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamNetConnection
	CreateListenSocketP2P(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket
	ConnectP2P(identityRemote SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection
	AcceptConnection(Conn HSteamNetConnection) EResult
	CloseConnection(Peer HSteamNetConnection, Reason int32, Debug string, EnableLinger bool) bool
	CloseListenSocket(Socket HSteamListenSocket) bool
	SetConnectionUserData(Peer HSteamNetConnection, UserData int64) bool
	GetConnectionUserData(Peer HSteamNetConnection) int64
	SetConnectionName(Peer HSteamNetConnection, Name string)
	GetConnectionName(Peer HSteamNetConnection, MaxLen int32) (name string, success bool)
	SendMessageToConnection(Conn HSteamNetConnection, Data []byte, SendFlags int32) (OutMessageNumber int64, result EResult)
	SendMessages(numMessages int32, Messages []SteamNetworkingMessage) (OutMessageNumberOrResult []int64)
	FlushMessagesOnConnection(Conn HSteamNetConnection) EResult
	ReceiveMessagesOnConnection(Conn HSteamNetConnection, MaxMessages int32) []SteamNetworkingMessage
	GetConnectionInfo(Conn HSteamNetConnection) (Info SteamNetConnectionInfo, success bool)
	GetConnectionRealTimeStatus(Conn HSteamNetConnection, numLanes int32) (Status SteamNetConnectionRealTimeStatus, Lanes []SteamNetConnectionRealTimeLaneStatus, result EResult)
	GetDetailedConnectionStatus(Conn HSteamNetConnection, bufSize int32) (status string, result int32)
	GetListenSocketAddress(Socket HSteamListenSocket) (address SteamNetworkingIPAddr, success bool)
	CreateSocketPair(UseNetworkLoopback bool) (OutConnection1 HSteamNetConnection, OutConnection2 HSteamNetConnection, Identity1 SteamNetworkingIdentity, Identity2 SteamNetworkingIdentity, succes bool)
	ConfigureConnectionLanes(Conn HSteamNetConnection, NumLanes int32, LanePriorities []int32, LaneWeights []uint16) EResult
	GetIdentity() (Identity SteamNetworkingIdentity, success bool)
	InitAuthentication() ESteamNetworkingAvailability
	GetAuthenticationStatus() (Details SteamNetAuthenticationStatus, availability ESteamNetworkingAvailability)
	CreatePollGroup() HSteamNetPollGroup
	DestroyPollGroup(PollGroup HSteamNetPollGroup) bool
	SetConnectionPollGroup(Conn HSteamNetConnection, PollGroup HSteamNetPollGroup) bool
	ReceiveMessagesOnPollGroup(PollGroup HSteamNetPollGroup, MaxMessages int32) []SteamNetworkingMessage
	ReceivedRelayAuthTicket(pvTicket []byte) (pOutParsedTicket SteamDatagramRelayAuthTicket, success bool)
	FindRelayAuthTicketForServer(gameServerIdentity SteamNetworkingIdentity, RemoteVirtualPort int32) (OutParsedTicket SteamDatagramRelayAuthTicket, secondsToExpire int32)
	ConnectToHostedDedicatedServer(targetIdentity SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection
	GetHostedDedicatedServerPort() uint16
	GetHostedDedicatedServerPOPID() SteamNetworkingPOPID
	GetHostedDedicatedServerAddress() (Routing SteamDatagramHostedAddress, result EResult)
	CreateHostedDedicatedServerListenSocket(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket
	GetGameCoordinatorServerLogin(LoginInfo *SteamDatagramGameCoordinatorServerLogin, SignedBlob int32) (Blob []byte, result EResult)
	GetCertificateRequest(BlobSize int32) (Blob []byte, errMsg SteamNetworkingErrMsg, success bool)
	SetCertificate(Certificate []byte) (errMsg SteamNetworkingErrMsg, success bool)
	ResetIdentity(Identity SteamNetworkingIdentity)
	RunCallbacks()
	BeginAsyncRequestFakeIP(NumPorts int32) bool
	GetFakeIP(FirstPortIdx int32) (Info SteamNetworkingFakeIPResult)
	CreateListenSocketP2PFakeIP(FakePortIdx int32, Options []SteamNetworkingConfigValue) HSteamListenSocket
	GetRemoteFakeIPForConnection(Conn HSteamNetConnection) (OutAddr SteamNetworkingIPAddr, result EResult)
	CreateFakeUDPPort(FakeServerPortIdx int32) *SteamNetworkingFakeUDPPort
}

type ISteamNetworkingUtils interface {
	AllocateMessage(AllocateBufferSize int32) *SteamNetworkingMessage
	InitRelayNetworkAccess()
	GetRelayNetworkStatus() (Details SteamRelayNetworkStatus, availability ESteamNetworkingAvailability)
	GetLocalPingLocation() (result SteamNetworkPingLocation, ageInSeconds float32)
	EstimatePingTimeBetweenTwoLocations(location1 SteamNetworkPingLocation, location2 SteamNetworkPingLocation) int32
	EstimatePingTimeFromLocalHost(remoteLocation SteamNetworkPingLocation) int32
	ConvertPingLocationToString(location SteamNetworkPingLocation, BufSize int32) string
	ParsePingLocationString(String string) (result SteamNetworkPingLocation, success bool)
	CheckPingDataUpToDate(MaxAgeSeconds float32) bool
	GetPingToDataCenter(popID SteamNetworkingPOPID) (ViaRelayPoP SteamNetworkingPOPID, pingTime int32)
	GetDirectPingToPOP(popID SteamNetworkingPOPID) int32
	GetPOPCount() int32
	GetPOPList(ListSz int32) (list []SteamNetworkingPOPID)
	GetLocalTimestamp() SteamNetworkingMicroseconds
	SetDebugOutputFunction(DetailLevel ESteamNetworkingSocketsDebugOutputType, Func FSteamNetworkingSocketsDebugOutput)
	IsFakeIPv4(IPv4 uint32) bool
	GetIPv4FakeIPType(IPv4 uint32) ESteamNetworkingFakeIPType
	GetRealIdentityForFakeIP(fakeIP SteamNetworkingIPAddr) (OutRealIdentity SteamNetworkingIdentity, result EResult)
	SetGlobalConfigValueInt32(eValue ESteamNetworkingConfigValue, val int32) bool
	SetGlobalConfigValueFloat(eValue ESteamNetworkingConfigValue, val float32) bool
	SetGlobalConfigValueString(eValue ESteamNetworkingConfigValue, val string) bool
	SetGlobalConfigValuePtr(eValue ESteamNetworkingConfigValue, val []byte) bool
	SetConnectionConfigValueInt32(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool
	SetConnectionConfigValueFloat(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val float32) bool
	SetConnectionConfigValueString(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool
	SetGlobalCallback_SteamNetConnectionStatusChanged(fnCallback FnSteamNetConnectionStatusChanged) bool
	SetGlobalCallback_SteamNetAuthenticationStatusChanged(fnCallback FnSteamNetAuthenticationStatusChanged) bool
	SetGlobalCallback_SteamRelayNetworkStatusChanged(fnCallback FnSteamRelayNetworkStatusChanged) bool
	SetGlobalCallback_FakeIPResult(fnCallback FnSteamNetworkingFakeIPResult) bool
	SetGlobalCallback_MessagesSessionRequest(fnCallback FnSteamNetworkingMessagesSessionRequest) bool
	SetGlobalCallback_MessagesSessionFailed(fnCallback FnSteamNetworkingMessagesSessionFailed) bool
	SetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool
	SetConfigValueStruct(opt SteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool
	GetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, cbResult size) (pOutDataType ESteamNetworkingConfigDataType, pResult []byte, result ESteamNetworkingGetConfigValueResult)
	GetConfigValueInfo(eValue ESteamNetworkingConfigValue) (pOutDataType ESteamNetworkingConfigDataType, pOutScope ESteamNetworkingConfigScope, name string)
	IterateGenericEditableConfigValues(eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue
	SteamNetworkingIPAddr_ToString(addr *SteamNetworkingIPAddr, cbBuf uint32, bWithPort bool) string
	SteamNetworkingIPAddr_ParseString(pAddr *SteamNetworkingIPAddr, pszStr string) bool
	SteamNetworkingIPAddr_GetFakeIPType(addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType
	SteamNetworkingIdentity_ToString(identity *SteamNetworkingIdentity, cbBuf uint32) string
	SteamNetworkingIdentity_ParseString(pIdentity *SteamNetworkingIdentity, pszStr string) bool
}

type ISteamParties interface {
	GetNumActiveBeacons() uint32
	GetBeaconByIndex(Index uint32) PartyBeaconID
	GetBeaconDetails(BeaconID PartyBeaconID, MetadataSize int32) (beaconOwnerSteamID CSteamID, Location SteamPartyBeaconLocation, metadata string, success bool)
	JoinParty(BeaconID PartyBeaconID) CallResult[JoinPartyCallback]
	GetNumAvailableBeaconLocations() (NumLocations uint32, success bool)
	GetAvailableBeaconLocations(MaxNumLocations uint32) (locationList []SteamPartyBeaconLocation, success bool)
	CreateBeacon(OpenSlots uint32, BeaconLocation SteamPartyBeaconLocation, ConnectString string, Metadata string) CallResult[CreateBeaconCallback]
	OnReservationCompleted(Beacon PartyBeaconID, userSteamID Uint64SteamID)
	CancelReservation(Beacon PartyBeaconID, userSteamID Uint64SteamID)
	ChangeNumOpenSlots(Beacon PartyBeaconID, OpenSlots uint32) CallResult[ChangeNumOpenSlotsCallback]
	DestroyBeacon(beacon PartyBeaconID) bool
	GetBeaconLocationData(BeaconLocation SteamPartyBeaconLocation, Data ESteamPartyBeaconLocationData, DataStringOut int32) (data string, success bool)
}

type ISteamRemotePlay interface {
	GetSessionCount() uint32
	GetSessionID(SessionIndex int32) RemotePlaySessionID
	GetSessionSteamID(SessionID RemotePlaySessionID) Uint64SteamID
	GetSessionClientName(SessionID RemotePlaySessionID) string
	GetSessionClientFormFactor(SessionID RemotePlaySessionID) ESteamDeviceFormFactor
	BGetSessionClientResolution(SessionID RemotePlaySessionID) (int32, int32, bool)
	BStartRemotePlayTogether(ShowOverlay bool) bool
	BSendRemotePlayTogetherInvite(friendSteamID Uint64SteamID) bool
}

type ISteamRemoteStorage interface {
	FileWrite(File string, Data []byte) bool
	FileRead(File string, Data []byte) (bytesRead int32)
	FileReadEx(File string, DataToReadSize int32) (Data []byte)
	FileWriteAsync(File string, Data []byte) CallResult[RemoteStorageFileWriteAsyncComplete]
	FileReadAsync(File string, Offset uint32, DataToReadSize uint32) CallResult[RemoteStorageFileReadAsyncComplete]
	FileReadAsyncComplete(ReadCallHandle SteamAPICall, DataToReadSize uint32) (Buffer []byte, success bool)
	FileForget(File string) bool
	FileDelete(File string) bool
	FileShare(File string) CallResult[RemoteStorageFileShareResult]
	SetSyncPlatforms(File string, RemoteStoragePlatform ERemoteStoragePlatform) bool
	FileWriteStreamOpen(File string) UGCFileWriteStreamHandle
	FileWriteStreamWriteChunk(writeHandle UGCFileWriteStreamHandle, Data []byte) bool
	FileWriteStreamClose(writeHandle UGCFileWriteStreamHandle) bool
	FileWriteStreamCancel(writeHandle UGCFileWriteStreamHandle) bool
	FileExists(File string) bool
	FilePersisted(File string) bool
	GetFileSize(File string) int32
	GetFileTimestamp(File string) int64
	GetSyncPlatforms(File string) ERemoteStoragePlatform
	GetFileCount() int32
	GetFileNameAndSize(FileIndex int32) (name string, FileSizeInBytes int32)
	GetQuota() (TotalBytes uint64, AvailableBytes uint64, success bool)
	IsCloudEnabledForAccount() bool
	IsCloudEnabledForApp() bool
	SetCloudEnabledForApp(Enabled bool)
	UGCDownload(Content UGCHandle, Priority uint32) CallResult[RemoteStorageDownloadUGCResult]
	GetUGCDownloadProgress(Content UGCHandle) (BytesDownloaded int32, BytesExpected int32, success bool)
	GetUGCDetails(Content UGCHandle) (AppID AppId_t, Name string, FileSizeInBytes int32, ownerSteamID CSteamID, success bool)
	UGCRead(Content UGCHandle, DataToReadSize int32, Offset uint32, Action EUGCReadAction) (Data []byte)
	GetCachedUGCCount() int32
	GetCachedUGCHandle(CachedContent int32) UGCHandle
	PublishWorkshopFile(File string, PreviewFile string, ConsumerAppId AppId_t, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray, WorkshopFileType EWorkshopFileType) CallResult[RemoteStoragePublishFileProgress]
	CreatePublishedFileUpdateRequest(PublishedFileId PublishedFileId) PublishedFileUpdateHandle
	UpdatePublishedFileFile(updateHandle PublishedFileUpdateHandle, File string) bool
	UpdatePublishedFilePreviewFile(updateHandle PublishedFileUpdateHandle, PreviewFile string) bool
	UpdatePublishedFileTitle(updateHandle PublishedFileUpdateHandle, Title string) bool
	UpdatePublishedFileDescription(updateHandle PublishedFileUpdateHandle, Description string) bool
	UpdatePublishedFileVisibility(updateHandle PublishedFileUpdateHandle, Visibility ERemoteStoragePublishedFileVisibility) bool
	UpdatePublishedFileTags(updateHandle PublishedFileUpdateHandle, Tags SteamParamStringArray) bool
	CommitPublishedFileUpdate(updateHandle PublishedFileUpdateHandle) CallResult[RemoteStorageUpdatePublishedFileResult]
	GetPublishedFileDetails(PublishedFileId PublishedFileId, MaxSecondsOld uint32) CallResult[RemoteStorageGetPublishedFileDetailsResult]
	DeletePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageDeletePublishedFileResult]
	EnumerateUserPublishedFiles(StartIndex uint32) CallResult[RemoteStorageEnumerateUserPublishedFilesResult]
	SubscribePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageSubscribePublishedFileResult]
	EnumerateUserSubscribedFiles(StartIndex uint32) CallResult[RemoteStorageEnumerateUserSubscribedFilesResult]
	UnsubscribePublishedFile(PublishedFileId PublishedFileId) CallResult[RemoteStorageUnsubscribePublishedFileResult]
	UpdatePublishedFileSetChangeDescription(updateHandle PublishedFileUpdateHandle, ChangeDescription string) bool
	GetPublishedItemVoteDetails(PublishedFileId PublishedFileId) CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]
	UpdateUserPublishedItemVote(PublishedFileId PublishedFileId, VoteUp bool) CallResult[RemoteStorageUpdateUserPublishedItemVoteResult]
	GetUserPublishedItemVoteDetails(PublishedFileId PublishedFileId) CallResult[RemoteStorageGetPublishedItemVoteDetailsResult]
	EnumerateUserSharedWorkshopFiles(steamId Uint64SteamID, StartIndex uint32, RequiredTags SteamParamStringArray, ExcludedTags SteamParamStringArray) CallResult[RemoteStorageEnumerateUserPublishedFilesResult]
	PublishVideo(VideoProvider EWorkshopVideoProvider, VideoAccount string, VideoIdentifier string, PreviewFile string, ConsumerAppId AppId_t, Title string, Description string, Visibility ERemoteStoragePublishedFileVisibility, Tags SteamParamStringArray) CallResult[RemoteStoragePublishFileProgress]
	SetUserPublishedFileAction(PublishedFileId PublishedFileId, Action EWorkshopFileAction) CallResult[RemoteStorageSetUserPublishedFileActionResult]
	EnumeratePublishedFilesByUserAction(Action EWorkshopFileAction, StartIndex uint32) CallResult[RemoteStorageEnumeratePublishedFilesByUserActionResult]
	EnumeratePublishedWorkshopFiles(EnumerationType EWorkshopEnumerationType, StartIndex uint32, Count uint32, Days uint32) (Tags SteamParamStringArray, UserTags SteamParamStringArray, apiHandle CallResult[RemoteStorageEnumerateWorkshopFilesResult])
	UGCDownloadToLocation(Content UGCHandle, Location string, Priority uint32) CallResult[RemoteStorageDownloadUGCResult]
	GetLocalFileChangeCount() int32
	GetLocalFileChange(FileIndex int32) (ChangeType ERemoteStorageLocalFileChange, FilePathType ERemoteStorageFilePathType, result []byte)
	BeginFileWriteBatch() bool
	EndFileWriteBatch() bool
}

type ISteamScreenshots interface {
	WriteScreenshot(pubRGB []byte, nWidth int32, nHeight int32) ScreenshotHandle
	AddScreenshotToLibrary(pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle
	TriggerScreenshot()
	HookScreenshots(bHook bool)
	SetLocation(hScreenshot ScreenshotHandle, pchLocation string) bool
	TagUser(hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool
	TagPublishedFile(hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool
	IsScreenshotsHooked() bool
	AddVRScreenshotToLibrary(eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle
}

type ISteamUGC interface {
	CreateQueryUserUGCRequest(accountID AccountID, listType EUserUGCList, matchingUGCType EUGCMatchingUGCType, sortOrder EUserUGCListSortOrder, creatorAppID AppId_t, consumerAppID AppId_t, page uint32) UGCQueryHandle
	CreateQueryAllUGCRequest(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId_t, consumerAppID AppId_t, page uint32) UGCQueryHandle
	CreateQueryAllUGCRequestCursor(queryType EUGCQuery, matchingeMatchingUGCTypeFileType EUGCMatchingUGCType, creatorAppID AppId_t, consumerAppID AppId_t, cursor string) UGCQueryHandle
	CreateQueryUGCDetailsRequest(publishedFileIDList []PublishedFileId) UGCQueryHandle
	SendQueryUGCRequest(handle UGCQueryHandle) CallResult[SteamUGCQueryCompleted]
	GetQueryUGCResult(handle UGCQueryHandle, index uint32) (details SteamUGCDetails, success bool)
	GetQueryUGCNumTags(handle UGCQueryHandle, index uint32) uint32
	GetQueryUGCTag(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool)
	GetQueryUGCTagDisplayName(handle UGCQueryHandle, index uint32, indexTag uint32, valueSize uint32) ([]byte, bool)
	GetQueryUGCPreviewURL(handle UGCQueryHandle, index uint32, URLSize uint32) ([]byte, bool)
	GetQueryUGCMetadata(handle UGCQueryHandle, index uint32, metadataSize uint32) ([]byte, bool)
	GetQueryUGCChildren(handle UGCQueryHandle, index uint32, maxEntries uint32) ([]PublishedFileId, bool)
	GetQueryUGCStatistic(handle UGCQueryHandle, index uint32, statType EItemStatistic) (statValue uint64, success bool)
	GetQueryUGCNumAdditionalPreviews(handle UGCQueryHandle, index uint32) uint32
	GetQueryUGCAdditionalPreview(handle UGCQueryHandle, index uint32, previewIndex uint32, URLSize uint32, originalFileNameSize uint32) ([]byte, []byte, EItemPreviewType, bool)
	GetQueryUGCNumKeyValueTags(handle UGCQueryHandle, index uint32) uint32
	GetQueryUGCKeyValueTag(handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, keySize uint32, valueSize uint32) ([]byte, []byte, bool)
	GetQueryFirstUGCKeyValueTag(handle UGCQueryHandle, index uint32, key string, valueSize uint32) ([]byte, bool)
	GetNumSupportedGameVersions(handle UGCQueryHandle, index uint32) uint32
	GetSupportedGameVersionData(handle UGCQueryHandle, index uint32, versionIndex uint32, gameBranchSize uint32) ([]byte, []byte, bool)
	GetQueryUGCContentDescriptors(handle UGCQueryHandle, index uint32, maxEntries uint32) (descriptorsList []EUGCContentDescriptorID)
	ReleaseQueryUGCRequest(handle UGCQueryHandle) bool
	AddRequiredTag(handle UGCQueryHandle, tagName string) bool
	AddRequiredTagGroup(handle UGCQueryHandle, tagGroups []SteamParamStringArray) bool
	AddExcludedTag(handle UGCQueryHandle, tagName string) bool
	SetReturnOnlyIDs(handle UGCQueryHandle, returnOnlyIDs bool) bool
	SetReturnKeyValueTags(handle UGCQueryHandle, returnKeyValueTags bool) bool
	SetReturnLongDescription(handle UGCQueryHandle, returnLongDescription bool) bool
	SetReturnMetadata(handle UGCQueryHandle, returnMetadata bool) bool
	SetReturnChildren(handle UGCQueryHandle, returnChildren bool) bool
	SetReturnAdditionalPreviews(handle UGCQueryHandle, returnAdditionalPreviews bool) bool
	SetReturnTotalOnly(handle UGCQueryHandle, returnTotalOnly bool) bool
	SetReturnPlaytimeStats(handle UGCQueryHandle, days uint32) bool
	SetLanguage(handle UGCQueryHandle, language string) bool
	SetAllowCachedResponse(handle UGCQueryHandle, maxAgeSeconds uint32) bool
	SetAdminQuery(handle UGCUpdateHandle, adminQuery bool) bool
	SetCloudFileNameFilter(handle UGCQueryHandle, matchCloudFileName string) bool
	SetMatchAnyTag(handle UGCQueryHandle, matchAnyTag bool) bool
	SetSearchText(handle UGCQueryHandle, searchText string) bool
	SetRankedByTrendDays(handle UGCQueryHandle, days uint32) bool
	SetTimeCreatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool
	SetTimeUpdatedDateRange(handle UGCQueryHandle, startTime RTime32, endTime RTime32) bool
	AddRequiredKeyValueTag(handle UGCQueryHandle, key string, value string) bool
	CreateItem(consumerAppId AppId_t, fileType EWorkshopFileType) CallResult[CreateItemResult]
	StartItemUpdate(consumerAppId AppId_t, publishedFileID PublishedFileId) UGCUpdateHandle
	SetItemTitle(handle UGCUpdateHandle, title string) bool
	SetItemDescription(handle UGCUpdateHandle, description string) bool
	SetItemUpdateLanguage(handle UGCUpdateHandle, language string) bool
	SetItemMetadata(handle UGCUpdateHandle, metaData string) bool
	SetItemVisibility(handle UGCUpdateHandle, visibility ERemoteStoragePublishedFileVisibility) bool
	SetItemTags(updateHandle UGCUpdateHandle, tags []SteamParamStringArray, allowAdminTags bool) bool
	SetItemContent(handle UGCUpdateHandle, contentFolder string) bool
	SetItemPreview(handle UGCUpdateHandle, previewFile string) bool
	SetAllowLegacyUpload(handle UGCUpdateHandle, allowLegacyUpload bool) bool
	RemoveAllItemKeyValueTags(handle UGCUpdateHandle) bool
	RemoveItemKeyValueTags(handle UGCUpdateHandle, key string) bool
	AddItemKeyValueTag(handle UGCUpdateHandle, key string, value string) bool
	AddItemPreviewFile(handle UGCUpdateHandle, previewFile string, Type EItemPreviewType) bool
	AddItemPreviewVideo(handle UGCUpdateHandle, videoID string) bool
	UpdateItemPreviewFile(handle UGCUpdateHandle, index uint32, previewFile string) bool
	UpdateItemPreviewVideo(handle UGCUpdateHandle, index uint32, videoID string) bool
	RemoveItemPreview(handle UGCUpdateHandle, index uint32) bool
	AddContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool
	RemoveContentDescriptor(handle UGCUpdateHandle, descID EUGCContentDescriptorID) bool
	SetRequiredGameVersions(handle UGCUpdateHandle, gameBranchMin string, gameBranchMax string) bool
	SubmitItemUpdate(handle UGCUpdateHandle, changeNote string) CallResult[SubmitItemUpdateResult]
	GetItemUpdateProgress(handle UGCUpdateHandle) (bytesProcessed uint64, bytesTotal uint64, status EItemUpdateStatus)
	SetUserItemVote(publishedFileID PublishedFileId, voteUp bool) CallResult[SetUserItemVoteResult]
	GetUserItemVote(publishedFileID PublishedFileId) CallResult[GetUserItemVoteResult]
	AddItemToFavorites(appId AppId_t, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged]
	RemoveItemFromFavorites(appId AppId_t, publishedFileID PublishedFileId) CallResult[UserFavoriteItemsListChanged]
	SubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageSubscribePublishedFileResult]
	UnsubscribeItem(publishedFileID PublishedFileId) CallResult[RemoteStorageUnsubscribePublishedFileResult]
	GetNumSubscribedItems() uint32
	GetSubscribedItems(maxEntries uint32) []PublishedFileId
	GetItemState(publishedFileID PublishedFileId) uint32
	GetItemInstallInfo(publishedFileID PublishedFileId, folderSize uint32) (uint64, uint32, []byte, bool)
	GetItemDownloadInfo(publishedFileID PublishedFileId) (bytesDownloaded uint64, bytesTotal uint64, success bool)
	DownloadItem(publishedFileID PublishedFileId, highPriority bool) bool
	BInitWorkshopForGameServer(workshopDepotID DepotId, folder string) bool
	SuspendDownloads(suspend bool)
	StartPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StartPlaytimeTrackingResult]
	StopPlaytimeTracking(publishedFileIDList []PublishedFileId) CallResult[StopPlaytimeTrackingResult]
	StopPlaytimeTrackingForAllItems() CallResult[StopPlaytimeTrackingResult]
	AddDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[AddUGCDependencyResult]
	RemoveDependency(parentPublishedFileID PublishedFileId, childPublishedFileID PublishedFileId) CallResult[RemoveUGCDependencyResult]
	AddAppDependency(publishedFileID PublishedFileId, appID AppId_t) CallResult[AddAppDependencyResult]
	RemoveAppDependency(publishedFileID PublishedFileId, appID AppId_t) CallResult[RemoveAppDependencyResult]
	GetAppDependencies(publishedFileID PublishedFileId) CallResult[GetAppDependenciesResult]
	DeleteItem(publishedFileID PublishedFileId) CallResult[DeleteItemResult]
	ShowWorkshopEULA() bool
	GetWorkshopEULAStatus() CallResult[WorkshopEULAStatus]
	GetUserContentDescriptorPreferences(maxEntries uint32) []EUGCContentDescriptorID
}

type ISteamUser interface {
	GetHSteamUser() HSteamUser
	BLoggedOn() bool
	GetSteamID() CSteamID
	TrackAppUsageEvent(gameID Uint64SteamID, appUsageEvent int32, extraInfo string)
	GetUserDataFolder(bufferSize int32) (folder string, success bool)
	StartVoiceRecording()
	StopVoiceRecording()
	GetAvailableVoice() (compressedDataSize uint32, result EVoiceResult)
	GetVoice(destBufferSize uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult)
	DecompressVoice(compressedData []byte, destBufferSize uint32, desiredSampleRate uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult)
	GetVoiceOptimalSampleRate() uint32
	GetAuthSessionTicket(maxTicket int32, optionalIdetity *SteamNetworkingIdentity) (ticketData []byte, ticketHandle HAuthTicket)
	GetAuthTicketForWebApi(identity string) HAuthTicket
	BeginAuthSession(authTicket []byte, steamID Uint64SteamID) EBeginAuthSessionResult
	EndAuthSession(steamID Uint64SteamID)
	CancelAuthTicket(hAuthTicket HAuthTicket)
	UserHasLicenseForApp(steamID Uint64SteamID, appID AppId_t) EUserHasLicenseForAppResult
	BIsBehindNAT() bool
	AdvertiseGame(gameServerSteamID Uint64SteamID, serverIP uint32, serverPort uint16)
	RequestEncryptedAppTicket(dataToInclude []byte) CallResult[EncryptedAppTicketResponse]
	GetEncryptedAppTicket() (ticketData []byte, ticketAvailable bool)
	GetGameBadgeLevel(series int32, foil bool) int32
	GetPlayerSteamLevel() int32
	RequestStoreAuthURL(redirectURL string) CallResult[StoreAuthURLResponse]
	BIsPhoneVerified() bool
	BIsTwoFactorEnabled() bool
	BIsPhoneIdentifying() bool
	BIsPhoneRequiringVerification() bool
	GetMarketEligibility() CallResult[MarketEligibilityResponse]
	GetDurationControl() CallResult[DurationControl]
	BSetDurationControlOnlineState(newState EDurationControlOnlineState) bool
}

type ISteamUserStats interface {
	GetStat(name string) (data int32, success bool)
	GetStatFloat(name string) (data float32, success bool)
	SetStat(name string, data int32) bool
	SetStatFloat(name string, data float32) bool
	UpdateAvgRateStat(name string, countThisSession float32, sessionLength float64) bool
	GetAchievement(name string) (achieved bool, success bool)
	SetAchievement(name string) bool
	ClearAchievement(name string) bool
	GetAchievementAndUnlockTime(name string) (achieved bool, unlockTime uint32, success bool)
	StoreStats() bool
	GetAchievementIcon(name string) int32
	GetAchievementDisplayAttribute(name string, key string) string
	IndicateAchievementProgress(name string, curProgress uint32, maxProgress uint32) bool
	GetNumAchievements() uint32
	GetAchievementName(achievement uint32) string
	RequestUserStats(userSteamID Uint64SteamID) CallResult[UserStatsReceived]
	GetUserStat(userSteamID Uint64SteamID, name string) (data int32, success bool)
	GetUserStatFloat(userSteamID Uint64SteamID, name string) (data float32, success bool)
	GetUserAchievement(userSteamID Uint64SteamID, name string) (achieved bool, success bool)
	GetUserAchievementAndUnlockTime(userSteamID Uint64SteamID, name string) (achieved bool, unlockTime uint32, success bool)
	ResetAllStats(achievementsToo bool) bool
	FindOrCreateLeaderboard(leaderboardName string, leaderboardSortMethod ELeaderboardSortMethod, leaderboardDisplayType ELeaderboardDisplayType) CallResult[LeaderboardFindResult]
	FindLeaderboard(leaderboardName string) CallResult[LeaderboardFindResult]
	GetLeaderboardName(steamLeaderboard SteamLeaderboard) string
	GetLeaderboardEntryCount(steamLeaderboard SteamLeaderboard) int32
	GetLeaderboardSortMethod(steamLeaderboard SteamLeaderboard) ELeaderboardSortMethod
	GetLeaderboardDisplayType(steamLeaderboard SteamLeaderboard) ELeaderboardDisplayType
	DownloadLeaderboardEntries(steamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) CallResult[LeaderboardScoresDownloaded]
	DownloadLeaderboardEntriesForUsers(steamLeaderboard SteamLeaderboard, prgUsers []CSteamID) CallResult[LeaderboardScoresDownloaded]
	GetDownloadedLeaderboardEntry(hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, cDetailsMax int32) (pLeaderboardEntry LeaderboardEntry, pDetails []int32, success bool)
	UploadLeaderboardScore(steamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32) CallResult[LeaderboardScoreUploaded]
	AttachLeaderboardUGC(steamLeaderboard SteamLeaderboard, hUGC UGCHandle) CallResult[LeaderboardUGCSet]
	GetNumberOfCurrentPlayers() CallResult[NumberOfCurrentPlayers]
	RequestGlobalAchievementPercentages() CallResult[GlobalAchievementPercentagesReady]
	GetMostAchievedAchievementInfo(nameBufLen uint32) (name string, percent float32, achieved bool, index int32)
	GetNextMostAchievedAchievementInfo(iteratorPrevious int32, nameBufLen uint32) (name string, percent float32, achieved bool, index int32)
	GetAchievementAchievedPercent(name string) (percent float32, success bool)
	RequestGlobalStats(historyDays int32) CallResult[GlobalStatsReceived]
	GetGlobalStat(pchStatName string) (data int64, success bool)
	GetGlobalStatDouble(pchStatName string) (data float64, success bool)
	GetGlobalStatHistory(pchStatName string, dataSize uint32) (data []int64)
	GetGlobalStatHistoryDouble(statName string, dataSize uint32) (data []float64)
	GetAchievementProgressLimits(name string) (minProgress int32, maxProgress int32, success bool)
	GetAchievementProgressLimitsFloat(name string) (minProgress float32, maxProgress float32, success bool)
}

type ISteamUtils interface {
	GetSecondsSinceAppActive() uint32
	GetSecondsSinceComputerActive() uint32
	GetConnectedUniverse() EUniverse
	GetServerRealTime() uint32
	GetIPCountry() string
	GetImageSize(imageIndex int32) (uint32, uint32, bool)
	GetImageRGBA(imageIndex int32, destBufferSize int32) ([]byte, bool)
	GetCurrentBatteryPower() uint8
	GetAppID() uint32
	SetOverlayNotificationPosition(notificationPosition ENotificationPosition)
	IsAPICallCompleted(steamAPICallHandle SteamAPICall) (failed bool, completed bool)
	GetAPICallFailureReason(steamAPICallHandle SteamAPICall) ESteamAPICallFailure
	GetAPICallResult(steamAPICallHandle SteamAPICall, callbackDataSize int32, callbackExpectedID int32) (callbackData []byte, failed bool, success bool)
	GetIPCCallCount() uint32
	SetWarningMessageHook(function SteamAPIWarningMessageHook)
	IsOverlayEnabled() bool
	BOverlayNeedsPresent() bool
	CheckFileSignature(fileName string) CallResult[CheckFileSignatureResult]
	ShowGamepadTextInput(inputMode EGamepadTextInputMode, lineInputMode EGamepadTextInputLineMode, description string, charMax uint32, existingText string) bool
	GetEnteredGamepadTextLength() uint32
	GetEnteredGamepadTextInput(text uint32) (string, bool)
	GetSteamUILanguage() string
	IsSteamRunningInVR() bool
	SetOverlayNotificationInset(horizontalInset int32, verticalInset int32)
	IsSteamInBigPictureMode() bool
	StartVRDashboard()
	IsVRHeadsetStreamingEnabled() bool
	SetVRHeadsetStreamingEnabled(enabled bool)
	IsSteamChinaLauncher() bool
	InitFilterText(filterOptions uint32) bool
	FilterText(context ETextFilteringContext, sourceSteamID Uint64SteamID, inputMessage string, byteSizeOutFilteredText uint32) (string, int32)
	GetIPv6ConnectivityState(protocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState
	IsSteamRunningOnSteamDeck() bool
	ShowFloatingGamepadTextInput(keyboardMode EFloatingGamepadTextInputMode, textFieldXPosition int32, textFieldYPosition int32, textFieldWidth int32, textFieldHeight int32) bool
	SetGameLauncherMode(launcherMode bool)
	DismissFloatingGamepadTextInput() bool
	DismissGamepadTextInput() bool
}
