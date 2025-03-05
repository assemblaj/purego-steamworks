package puregosteamworks

import (
	"unsafe"
)

type CSteamID uint64
type CGameID uint64
type Uint64SteamID uint64
type Uint64GameID uint64

var steamAPIDLL uintptr

type SteamAPIWarningMessageHook string
type SteamCallbackID uint64

type lint64 int64
type ulint64 uint64
type intp int64
type uintp uint64
type AppId_t uint32
type DepotId uint
type RTime32 uint
type SteamAPICall uint64
type AccountID uint
type PFNPreMinidumpCallback uintptr
type HSteamPipe int32
type HSteamUser int32
type SteamErrMsg [1024]byte
type UGCHandle uint64

const (
	AppIdInvalid             AppId_t      = 0x0
	DepotIdInvalid           DepotId      = 0x0
	APICallInvalid           SteamAPICall = 0x0
	AccountIdInvalid         AccountID    = 0
	SteamAccountIDMask       uint32       = 0xFFFFFFFF
	SteamAccountInstanceMask uint32       = 0x000FFFFF
	SteamUserDefaultInstance uint32       = 1
	GameExtraInfoMax         int32        = 64
	MaxSteamErrMsg           int32        = 1024
)

type EUniverse int32

const (
	EUniverseInvalid  EUniverse = 0
	EUniversePublic   EUniverse = 1
	EUniverseBeta     EUniverse = 2
	EUniverseInternal EUniverse = 3
	EUniverseDev      EUniverse = 4
	EUniverseMax      EUniverse = 5
)

type EResult int32

const (
	EResultNone                                    EResult = 0
	EResultOK                                      EResult = 1
	EResultFail                                    EResult = 2
	EResultNoConnection                            EResult = 3
	EResultInvalidPassword                         EResult = 5
	EResultLoggedInElsewhere                       EResult = 6
	EResultInvalidProtocolVer                      EResult = 7
	EResultInvalidParam                            EResult = 8
	EResultFileNotFound                            EResult = 9
	EResultBusy                                    EResult = 10
	EResultInvalidState                            EResult = 11
	EResultInvalidName                             EResult = 12
	EResultInvalidEmail                            EResult = 13
	EResultDuplicateName                           EResult = 14
	EResultAccessDenied                            EResult = 15
	EResultTimeout                                 EResult = 16
	EResultBanned                                  EResult = 17
	EResultAccountNotFound                         EResult = 18
	EResultInvalidSteamID                          EResult = 19
	EResultServiceUnavailable                      EResult = 20
	EResultNotLoggedOn                             EResult = 21
	EResultPending                                 EResult = 22
	EResultEncryptionFailure                       EResult = 23
	EResultInsufficientPrivilege                   EResult = 24
	EResultLimitExceeded                           EResult = 25
	EResultRevoked                                 EResult = 26
	EResultExpired                                 EResult = 27
	EResultAlreadyRedeemed                         EResult = 28
	EResultDuplicateRequest                        EResult = 29
	EResultAlreadyOwned                            EResult = 30
	EResultIPNotFound                              EResult = 31
	EResultPersistFailed                           EResult = 32
	EResultLockingFailed                           EResult = 33
	EResultLogonSessionReplaced                    EResult = 34
	EResultConnectFailed                           EResult = 35
	EResultHandshakeFailed                         EResult = 36
	EResultIOFailure                               EResult = 37
	EResultRemoteDisconnect                        EResult = 38
	EResultShoppingCartNotFound                    EResult = 39
	EResultBlocked                                 EResult = 40
	EResultIgnored                                 EResult = 41
	EResultNoMatch                                 EResult = 42
	EResultAccountDisabled                         EResult = 43
	EResultServiceReadOnly                         EResult = 44
	EResultAccountNotFeatured                      EResult = 45
	EResultAdministratorOK                         EResult = 46
	EResultContentVersion                          EResult = 47
	EResultTryAnotherCM                            EResult = 48
	EResultPasswordRequiredToKickSession           EResult = 49
	EResultAlreadyLoggedInElsewhere                EResult = 50
	EResultSuspended                               EResult = 51
	EResultCancelled                               EResult = 52
	EResultDataCorruption                          EResult = 53
	EResultDiskFull                                EResult = 54
	EResultRemoteCallFailed                        EResult = 55
	EResultPasswordUnset                           EResult = 56
	EResultExternalAccountUnlinked                 EResult = 57
	EResultPSNTicketInvalid                        EResult = 58
	EResultExternalAccountAlreadyLinked            EResult = 59
	EResultRemoteFileConflict                      EResult = 60
	EResultIllegalPassword                         EResult = 61
	EResultSameAsPreviousValue                     EResult = 62
	EResultAccountLogonDenied                      EResult = 63
	EResultCannotUseOldPassword                    EResult = 64
	EResultInvalidLoginAuthCode                    EResult = 65
	EResultAccountLogonDeniedNoMail                EResult = 66
	EResultHardwareNotCapableOfIPT                 EResult = 67
	EResultIPTInitError                            EResult = 68
	EResultParentalControlRestricted               EResult = 69
	EResultFacebookQueryError                      EResult = 70
	EResultExpiredLoginAuthCode                    EResult = 71
	EResultIPLoginRestrictionFailed                EResult = 72
	EResultAccountLockedDown                       EResult = 73
	EResultAccountLogonDeniedVerifiedEmailRequired EResult = 74
	EResultNoMatchingURL                           EResult = 75
	EResultBadResponse                             EResult = 76
	EResultRequirePasswordReEntry                  EResult = 77
	EResultValueOutOfRange                         EResult = 78
	EResultUnexpectedError                         EResult = 79
	EResultDisabled                                EResult = 80
	EResultInvalidCEGSubmission                    EResult = 81
	EResultRestrictedDevice                        EResult = 82
	EResultRegionLocked                            EResult = 83
	EResultRateLimitExceeded                       EResult = 84
	EResultAccountLoginDeniedNeedTwoFactor         EResult = 85
	EResultItemDeleted                             EResult = 86
	EResultAccountLoginDeniedThrottle              EResult = 87
	EResultTwoFactorCodeMismatch                   EResult = 88
	EResultTwoFactorActivationCodeMismatch         EResult = 89
	EResultAccountAssociatedToMultiplePartners     EResult = 90
	EResultNotModified                             EResult = 91
	EResultNoMobileDevice                          EResult = 92
	EResultTimeNotSynced                           EResult = 93
	EResultSmsCodeFailed                           EResult = 94
	EResultAccountLimitExceeded                    EResult = 95
	EResultAccountActivityLimitExceeded            EResult = 96
	EResultPhoneActivityLimitExceeded              EResult = 97
	EResultRefundToWallet                          EResult = 98
	EResultEmailSendFailure                        EResult = 99
	EResultNotSettled                              EResult = 100
	EResultNeedCaptcha                             EResult = 101
	EResultGSLTDenied                              EResult = 102
	EResultGSOwnerDenied                           EResult = 103
	EResultInvalidItemType                         EResult = 104
	EResultIPBanned                                EResult = 105
	EResultGSLTExpired                             EResult = 106
	EResultInsufficientFunds                       EResult = 107
	EResultTooManyPending                          EResult = 108
	EResultNoSiteLicensesFound                     EResult = 109
	EResultWGNetworkSendExceeded                   EResult = 110
	EResultAccountNotFriends                       EResult = 111
	EResultLimitedUserAccount                      EResult = 112
	EResultCantRemoveItem                          EResult = 113
	EResultAccountDeleted                          EResult = 114
	EResultExistingUserCancelledLicense            EResult = 115
	EResultCommunityCooldown                       EResult = 116
	EResultNoLauncherSpecified                     EResult = 117
	EResultMustAgreeToSSA                          EResult = 118
	EResultLauncherMigrated                        EResult = 119
	EResultSteamRealmMismatch                      EResult = 120
	EResultInvalidSignature                        EResult = 121
	EResultParseFailure                            EResult = 122
	EResultNoVerifiedPhone                         EResult = 123
	EResultInsufficientBattery                     EResult = 124
	EResultChargerRequired                         EResult = 125
	EResultCachedCredentialInvalid                 EResult = 126
	EResultPhoneNumberIsVOIP                       EResult = 127
	EResultNotSupported                            EResult = 128
	EResultFamilySizeLimitExceeded                 EResult = 129
	EResultOfflineAppCacheInvalid                  EResult = 130
)

type EDenyReason int32

const (
	EDenyInvalid                 EDenyReason = 0
	EDenyInvalidVersion          EDenyReason = 1
	EDenyGeneric                 EDenyReason = 2
	EDenyNotLoggedOn             EDenyReason = 3
	EDenyNoLicense               EDenyReason = 4
	EDenyCheater                 EDenyReason = 5
	EDenyLoggedInElseWhere       EDenyReason = 6
	EDenyUnknownText             EDenyReason = 7
	EDenyIncompatibleAnticheat   EDenyReason = 8
	EDenyMemoryCorruption        EDenyReason = 9
	EDenyIncompatibleSoftware    EDenyReason = 10
	EDenySteamConnectionLost     EDenyReason = 11
	EDenySteamConnectionError    EDenyReason = 12
	EDenySteamResponseTimedOut   EDenyReason = 13
	EDenySteamValidationStalled  EDenyReason = 14
	EDenySteamOwnerLeftGuestUser EDenyReason = 15
)

type EChatSteamIDInstanceFlags int32

const (
	EChatAccountInstanceMask  EChatSteamIDInstanceFlags = 4095
	EChatInstanceFlagClan     EChatSteamIDInstanceFlags = 524288
	EChatInstanceFlagLobby    EChatSteamIDInstanceFlags = 262144
	EChatInstanceFlagMMSLobby EChatSteamIDInstanceFlags = 131072
)

type EBetaBranchFlags int32

const (
	EBetaBranchNone      EBetaBranchFlags = 0
	EBetaBranchDefault   EBetaBranchFlags = 1
	EBetaBranchAvailable EBetaBranchFlags = 2
	EBetaBranchPrivate   EBetaBranchFlags = 4
	EBetaBranchSelected  EBetaBranchFlags = 8
	EBetaBranchInstalled EBetaBranchFlags = 16
)

type ESteamAPICallFailure int32

const (
	ESteamAPICallFailureNone               ESteamAPICallFailure = -1
	ESteamAPICallFailureSteamGone          ESteamAPICallFailure = 0
	ESteamAPICallFailureNetworkFailure     ESteamAPICallFailure = 1
	ESteamAPICallFailureInvalidHandle      ESteamAPICallFailure = 2
	ESteamAPICallFailureMismatchedCallback ESteamAPICallFailure = 3
)

type ESteamAPIInitResult int32

const (
	ESteamAPIInitResultOK              ESteamAPIInitResult = 0
	ESteamAPIInitResultFailedGeneric   ESteamAPIInitResult = 1
	ESteamAPIInitResultNoSteamClient   ESteamAPIInitResult = 2
	ESteamAPIInitResultVersionMismatch ESteamAPIInitResult = 3
)

type SteamAPICallCompleted struct {
	AsyncCall SteamAPICall
	Callback  int32
	Param     uint32
}

const SteamAPICallCompletedID SteamCallbackID = 703

func uintptrToStruct[T any](ptr uintptr) *T {
	if ptr == 0 {
		return nil
	}
	return (*T)(unsafe.Pointer(ptr))
}

func structToUintptr[T any](s *T) uintptr {
	if s == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(s))
}

func bytesToStruct[T any](data []byte) *T {
	return (*T)(unsafe.Pointer(&data[0]))
}
