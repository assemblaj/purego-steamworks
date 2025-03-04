package user

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	. "github.com/assemblaj/purego-steamworks"
)

const (
	HAuthTicketInvalid HAuthTicket = 0
)

type EDurationControlNotification int32

const (
	EDurationControlNotificationNone          EDurationControlNotification = 0
	EDurationControlNotification1Hour         EDurationControlNotification = 1
	EDurationControlNotification3Hours        EDurationControlNotification = 2
	EDurationControlNotificationHalfProgress  EDurationControlNotification = 3
	EDurationControlNotificationNoProgress    EDurationControlNotification = 4
	EDurationControlNotificationExitSoon3h    EDurationControlNotification = 5
	EDurationControlNotificationExitSoon5h    EDurationControlNotification = 6
	EDurationControlNotificationExitSoonNight EDurationControlNotification = 7
)

type EDurationControlProgress int32

const (
	EDurationControlProgressFull  EDurationControlProgress = 0
	EDurationControlProgressHalf  EDurationControlProgress = 1
	EDurationControlProgressNone  EDurationControlProgress = 2
	EDurationControlExitSoon3h    EDurationControlProgress = 3
	EDurationControlExitSoon5h    EDurationControlProgress = 4
	EDurationControlExitSoonNight EDurationControlProgress = 5
)

type EMarketNotAllowedReasonFlags int32

const (
	EMarketNotAllowedReasonNone                             EMarketNotAllowedReasonFlags = 0
	EMarketNotAllowedReasonTemporaryFailure                 EMarketNotAllowedReasonFlags = 1
	EMarketNotAllowedReasonAccountDisabled                  EMarketNotAllowedReasonFlags = 2
	EMarketNotAllowedReasonAccountLockedDown                EMarketNotAllowedReasonFlags = 4
	EMarketNotAllowedReasonAccountLimited                   EMarketNotAllowedReasonFlags = 8
	EMarketNotAllowedReasonTradeBanned                      EMarketNotAllowedReasonFlags = 16
	EMarketNotAllowedReasonAccountNotTrusted                EMarketNotAllowedReasonFlags = 32
	EMarketNotAllowedReasonSteamGuardNotEnabled             EMarketNotAllowedReasonFlags = 64
	EMarketNotAllowedReasonSteamGuardOnlyRecentlyEnabled    EMarketNotAllowedReasonFlags = 128
	EMarketNotAllowedReasonRecentPasswordReset              EMarketNotAllowedReasonFlags = 256
	EMarketNotAllowedReasonNewPaymentMethod                 EMarketNotAllowedReasonFlags = 512
	EMarketNotAllowedReasonInvalidCookie                    EMarketNotAllowedReasonFlags = 1024
	EMarketNotAllowedReasonUsingNewDevice                   EMarketNotAllowedReasonFlags = 2048
	EMarketNotAllowedReasonRecentSelfRefund                 EMarketNotAllowedReasonFlags = 4096
	EMarketNotAllowedReasonNewPaymentMethodCannotBeVerified EMarketNotAllowedReasonFlags = 8192
	EMarketNotAllowedReasonNoRecentPurchases                EMarketNotAllowedReasonFlags = 16384
	EMarketNotAllowedReasonAcceptedWalletGift               EMarketNotAllowedReasonFlags = 32768
)

type EAuthSessionResponse int32

const (
	EAuthSessionResponseOK                               EAuthSessionResponse = 0
	EAuthSessionResponseUserNotConnectedToSteam          EAuthSessionResponse = 1
	EAuthSessionResponseNoLicenseOrExpired               EAuthSessionResponse = 2
	EAuthSessionResponseVACBanned                        EAuthSessionResponse = 3
	EAuthSessionResponseLoggedInElseWhere                EAuthSessionResponse = 4
	EAuthSessionResponseVACCheckTimedOut                 EAuthSessionResponse = 5
	EAuthSessionResponseAuthTicketCanceled               EAuthSessionResponse = 6
	EAuthSessionResponseAuthTicketInvalidAlreadyUsed     EAuthSessionResponse = 7
	EAuthSessionResponseAuthTicketInvalid                EAuthSessionResponse = 8
	EAuthSessionResponsePublisherIssuedBan               EAuthSessionResponse = 9
	EAuthSessionResponseAuthTicketNetworkIdentityFailure EAuthSessionResponse = 10
)

type HAuthTicket uint

type EVoiceResult int32

const (
	EVoiceResultOK                   EVoiceResult = 0
	EVoiceResultNotInitialized       EVoiceResult = 1
	EVoiceResultNotRecording         EVoiceResult = 2
	EVoiceResultNoData               EVoiceResult = 3
	EVoiceResultBufferTooSmall       EVoiceResult = 4
	EVoiceResultDataCorrupted        EVoiceResult = 5
	EVoiceResultRestricted           EVoiceResult = 6
	EVoiceResultUnsupportedCodec     EVoiceResult = 7
	EVoiceResultReceiverOutOfDate    EVoiceResult = 8
	EVoiceResultReceiverDidNotAnswer EVoiceResult = 9
)

type EDurationControlOnlineState int32

const (
	EDurationControlOnlineStateInvalid       EDurationControlOnlineState = 0
	EDurationControlOnlineStateOffline       EDurationControlOnlineState = 1
	EDurationControlOnlineStateOnline        EDurationControlOnlineState = 2
	EDurationControlOnlineStateOnlineHighPri EDurationControlOnlineState = 3
)

type EUserHasLicenseForAppResult int32

const (
	EUserHasLicenseResultHasLicense         EUserHasLicenseForAppResult = 0
	EUserHasLicenseResultDoesNotHaveLicense EUserHasLicenseForAppResult = 1
	EUserHasLicenseResultNoAuth             EUserHasLicenseForAppResult = 2
)

type EBeginAuthSessionResult int32

const (
	EBeginAuthSessionResultOK               EBeginAuthSessionResult = 0
	EBeginAuthSessionResultInvalidTicket    EBeginAuthSessionResult = 1
	EBeginAuthSessionResultDuplicateRequest EBeginAuthSessionResult = 2
	EBeginAuthSessionResultInvalidVersion   EBeginAuthSessionResult = 3
	EBeginAuthSessionResultGameMismatch     EBeginAuthSessionResult = 4
	EBeginAuthSessionResultExpiredTicket    EBeginAuthSessionResult = 5
)

var (
	steamUser_init                            func() uintptr
	iSteamUser_GetHSteamUser                  func(steamUser uintptr) HSteamUser
	iSteamUser_BLoggedOn                      func(steamUser uintptr) bool
	iSteamUser_GetSteamID                     func(steamUser uintptr) Uint64SteamID
	iSteamUser_TrackAppUsageEvent             func(steamUser uintptr, gameID Uint64SteamID, eAppUsageEvent int32, pchExtraInfo string)
	iSteamUser_GetUserDataFolder              func(steamUser uintptr, pchBuffer *[]byte, cubBuffer int32) bool
	iSteamUser_StartVoiceRecording            func(steamUser uintptr)
	iSteamUser_StopVoiceRecording             func(steamUser uintptr)
	iSteamUser_GetAvailableVoice              func(steamUser uintptr, pcbCompressed *uint32, pcbUncompressedDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult
	iSteamUser_GetVoice                       func(steamUser uintptr, bWantCompressed bool, pDestBuffer *[]byte, cbDestBufferSize uint32, nBytesWritten *uint32, bWantUncompressedDeprecated bool, pUncompressedDestBufferDeprecated []byte, cbUncompressedDestBufferSizeDeprecated uint32, nUncompressBytesWrittenDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult
	iSteamUser_DecompressVoice                func(steamUser uintptr, pCompressed []byte, cbCompressed uint32, pDestBuffer *[]byte, cbDestBufferSize uint32, nBytesWritten *uint32, nDesiredSampleRate uint32) EVoiceResult
	iSteamUser_GetVoiceOptimalSampleRate      func(steamUser uintptr) uint32
	iSteamUser_GetAuthSessionTicket           func(steamUser uintptr, pTicket *[]byte, cbMaxTicket int32, pcbTicket *uint32, pSteamNetworkingIdentity *SteamNetworkingIdentity) HAuthTicket
	iSteamUser_GetAuthTicketForWebApi         func(steamUser uintptr, pchIdentity string) HAuthTicket
	iSteamUser_BeginAuthSession               func(steamUser uintptr, pAuthTicket []byte, cbAuthTicket int32, steamID Uint64SteamID) EBeginAuthSessionResult
	iSteamUser_EndAuthSession                 func(steamUser uintptr, steamID Uint64SteamID)
	iSteamUser_CancelAuthTicket               func(steamUser uintptr, hAuthTicket HAuthTicket)
	iSteamUser_UserHasLicenseForApp           func(steamUser uintptr, steamID Uint64SteamID, appID AppId) EUserHasLicenseForAppResult
	iSteamUser_BIsBehindNAT                   func(steamUser uintptr) bool
	iSteamUser_AdvertiseGame                  func(steamUser uintptr, steamIDGameServer Uint64SteamID, unIPServer uint32, usPortServer uint16)
	iSteamUser_RequestEncryptedAppTicket      func(steamUser uintptr, pDataToInclude []byte, cbDataToInclude int32) SteamAPICall
	iSteamUser_GetEncryptedAppTicket          func(steamUser uintptr, pTicket *[]byte, cbMaxTicket int32, pcbTicket *uint32) bool
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

var steamUser uintptr

func user() uintptr {
	if steamUser == 0 {
		steamUser = steamUser_init()
		return steamUser
	}
	return steamUser
}

func GetHSteamUser() HSteamUser {
	return iSteamUser_GetHSteamUser(user())
}

func BLoggedOn() bool {
	return iSteamUser_BLoggedOn(user())
}

func GetSteamID() Uint64SteamID {
	return iSteamUser_GetSteamID(user())
}

func TrackAppUsageEvent(gameID Uint64SteamID, appUsageEvent int32, extraInfo string) {
	iSteamUser_TrackAppUsageEvent(user(), gameID, appUsageEvent, extraInfo)
}

func GetUserDataFolder(bufferSize int32) (folder string, success bool) {
	var pchBuffer []byte = make([]byte, 0, bufferSize)
	success = iSteamUser_GetUserDataFolder(user(), &pchBuffer, bufferSize)
	return string(pchBuffer), success
}

func StartVoiceRecording() {
	iSteamUser_StartVoiceRecording(user())
}

func StopVoiceRecording() {
	iSteamUser_StopVoiceRecording(user())
}

func GetAvailableVoice() (compressedDataSize uint32, result EVoiceResult) {
	var pcbUncompressedDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated uint32
	result = iSteamUser_GetAvailableVoice(user(), &compressedDataSize, &pcbUncompressedDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated)
	return compressedDataSize, result
}

func GetVoice(destBufferSize uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult) {
	var bWantCompressed bool = true
	destBuffer = make([]byte, 0, destBufferSize)
	var bWantUncompressedDeprecated bool
	var pUncompressedDestBufferDeprecated []byte
	var cbUncompressedDestBufferSizeDeprecated, nUncompressBytesWrittenDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated uint32
	result = iSteamUser_GetVoice(user(), bWantCompressed, &destBuffer, destBufferSize, &bytesWritten, bWantUncompressedDeprecated, pUncompressedDestBufferDeprecated, cbUncompressedDestBufferSizeDeprecated, &nUncompressBytesWrittenDeprecated, nUncompressedVoiceDesiredSampleRateDeprecated)
	return destBuffer, bytesWritten, result
}

func DecompressVoice(compressedData []byte, destBufferSize uint32, desiredSampleRate uint32) (destBuffer []byte, bytesWritten uint32, result EVoiceResult) {
	destBuffer = make([]byte, 0, destBufferSize)
	result = iSteamUser_DecompressVoice(user(), compressedData, uint32(len(compressedData)), &destBuffer, destBufferSize, &bytesWritten, desiredSampleRate)
	return destBuffer, bytesWritten, result
}

func GetVoiceOptimalSampleRate() uint32 {
	return iSteamUser_GetVoiceOptimalSampleRate(user())
}

func GetAuthSessionTicket(maxTicket int32, optionalIdetity *SteamNetworkingIdentity) (ticketData []byte, ticketHandle HAuthTicket) {
	var pcbTicket uint32
	ticketData = make([]byte, 0, maxTicket)
	ticketHandle = iSteamUser_GetAuthSessionTicket(user(), &ticketData, maxTicket, &pcbTicket, optionalIdetity)
	return ticketData[:pcbTicket], ticketHandle
}

func GetAuthTicketForWebApi(identity string) HAuthTicket {
	return iSteamUser_GetAuthTicketForWebApi(user(), identity)
}

func BeginAuthSession(authTicket []byte, steamID Uint64SteamID) EBeginAuthSessionResult {
	return iSteamUser_BeginAuthSession(user(), authTicket, int32(len(authTicket)), steamID)
}

func EndAuthSession(steamID Uint64SteamID) {
	iSteamUser_EndAuthSession(user(), steamID)
}

func CancelAuthTicket(hAuthTicket HAuthTicket) {
	iSteamUser_CancelAuthTicket(user(), hAuthTicket)
}

func UserHasLicenseForApp(steamID Uint64SteamID, appID AppId) EUserHasLicenseForAppResult {
	return iSteamUser_UserHasLicenseForApp(user(), steamID, appID)
}

func BIsBehindNAT() bool {
	return iSteamUser_BIsBehindNAT(user())
}

func AdvertiseGame(gameServerSteamID Uint64SteamID, serverIP uint32, serverPort uint16) {
	iSteamUser_AdvertiseGame(user(), gameServerSteamID, serverIP, serverPort)
}

func RequestEncryptedAppTicket(dataToInclude []byte) CallResult[EncryptedAppTicketResponse] {
	handle := iSteamUser_RequestEncryptedAppTicket(user(), dataToInclude, int32(len(dataToInclude)))
	return CallResult[EncryptedAppTicketResponse]{Handle: handle}
}

// Retrieves a finished ticket.
// Performs the API call once to get the proper size and then again to get the
// actual data.
func GetEncryptedAppTicket() (ticketData []byte, ticketAvailable bool) {
	var ticketSize uint32
	available := iSteamUser_GetEncryptedAppTicket(user(), nil, 0, &ticketSize)
	if available {
		ticketData = make([]byte, 0, ticketSize)
		iSteamUser_GetEncryptedAppTicket(user(), &ticketData, 0, &ticketSize)
	}
	return ticketData, available
}

func GetGameBadgeLevel(series int32, foil bool) int32 {
	return iSteamUser_GetGameBadgeLevel(user(), series, foil)
}

func GetPlayerSteamLevel() int32 {
	return iSteamUser_GetPlayerSteamLevel(user())
}

func RequestStoreAuthURL(redirectURL string) CallResult[StoreAuthURLResponse] {
	handle := iSteamUser_RequestStoreAuthURL(user(), redirectURL)
	return CallResult[StoreAuthURLResponse]{Handle: handle}
}

func BIsPhoneVerified() bool {
	return iSteamUser_BIsPhoneVerified(user())
}

func BIsTwoFactorEnabled() bool {
	return iSteamUser_BIsTwoFactorEnabled(user())
}

func BIsPhoneIdentifying() bool {
	return iSteamUser_BIsPhoneIdentifying(user())
}

func BIsPhoneRequiringVerification() bool {
	return iSteamUser_BIsPhoneRequiringVerification(user())
}

func GetMarketEligibility() CallResult[MarketEligibilityResponse] {
	handle := iSteamUser_GetMarketEligibility(user())
	return CallResult[MarketEligibilityResponse]{Handle: handle}
}

func GetDurationControl() CallResult[DurationControl] {
	handle := iSteamUser_GetDurationControl(user())
	return CallResult[DurationControl]{Handle: handle}
}

func BSetDurationControlOnlineState(newState EDurationControlOnlineState) bool {
	return iSteamUser_BSetDurationControlOnlineState(user(), newState)
}
