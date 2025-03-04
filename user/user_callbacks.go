package user

import . "github.com/assemblaj/purego-steamworks"

const (
	SteamServersConnectedID         SteamCallbackID = 101
	SteamServerConnectFailureID     SteamCallbackID = 102
	SteamServersDisconnectedID      SteamCallbackID = 103
	ClientGameServerDenyID          SteamCallbackID = 113
	IPCFailureID                    SteamCallbackID = 117
	LicensesUpdatedID               SteamCallbackID = 125
	ValidateAuthTicketResponseID    SteamCallbackID = 143
	MicroTxnAuthorizationResponseID SteamCallbackID = 152
	EncryptedAppTicketResponseID    SteamCallbackID = 154
	GetAuthSessionTicketResponseID  SteamCallbackID = 163
	GameWebCallbackID               SteamCallbackID = 164
	StoreAuthURLResponseID          SteamCallbackID = 165
	MarketEligibilityResponseID     SteamCallbackID = 166
	DurationControlID               SteamCallbackID = 167
	GetTicketForWebApiResponseID    SteamCallbackID = 168
)

type SteamServersConnected struct {
}
type SteamServerConnectFailure struct {
	Result        EResult
	StillRetrying bool
}
type SteamServersDisconnected struct {
	Result EResult
}
type ClientGameServerDeny struct {
	AppID          uint32
	GameServerIP   uint32
	GameServerPort uint16
	Secure         uint16
	Reason         uint32
}
type IPCFailure struct {
	FailureType uint8
}
type EFailureType int32

const (
	EFailureFlushedCallbackQueue EFailureType = 0
	EFailurePipeFail             EFailureType = 1
)

type LicensesUpdated struct {
}
type ValidateAuthTicketResponse struct {
	SteamID             CSteamID
	AuthSessionResponse EAuthSessionResponse
	OwnerSteamID        CSteamID
}
type MicroTxnAuthorizationResponse struct {
	AppID      uint32
	OrderID    uint64
	Authorized uint8
}
type EncryptedAppTicketResponse struct {
	Result EResult
}
type GetAuthSessionTicketResponse struct {
	AuthTicket HAuthTicket
	Result     EResult
}
type GameWebCallback struct {
	URL [256]byte
}
type StoreAuthURLResponse struct {
	URL [512]byte
}
type MarketEligibilityResponse struct {
	Allowed                bool
	NotAllowedReason       EMarketNotAllowedReasonFlags
	AllowedAtTime          RTime32
	SteamGuardRequiredDays int32
	NewDeviceCooldown      int32
}
type DurationControl struct {
	Result       EResult
	AppId        AppId
	Applicable   bool
	Last5h       int32
	Progress     EDurationControlProgress
	Notification EDurationControlNotification
	Today        int32
	Remaining    int32
}
type GetTicketForWebApiResponse struct {
	AuthTicket HAuthTicket
	Result     EResult
	TicketSize int32
	TicketData [2560]uint8
}

func (cb SteamServersConnected) CallbackID() SteamCallbackID {
	return SteamServersConnectedID
}

func (cb SteamServersConnected) Name() string {
	return "SteamServersConnected"
}

func (cb SteamServersConnected) String() string {
	return CallbackString(cb)
}

func (cb SteamServerConnectFailure) CallbackID() SteamCallbackID {
	return SteamServerConnectFailureID
}

func (cb SteamServerConnectFailure) Name() string {
	return "SteamServerConnectFailure"
}

func (cb SteamServerConnectFailure) String() string {
	return CallbackString(cb)
}

func (cb SteamServersDisconnected) CallbackID() SteamCallbackID {
	return SteamServersDisconnectedID
}

func (cb SteamServersDisconnected) Name() string {
	return "SteamServersDisconnected"
}

func (cb SteamServersDisconnected) String() string {
	return CallbackString(cb)
}

func (cb ClientGameServerDeny) CallbackID() SteamCallbackID {
	return ClientGameServerDenyID
}

func (cb ClientGameServerDeny) Name() string {
	return "ClientGameServerDeny"
}

func (cb ClientGameServerDeny) String() string {
	return CallbackString(cb)
}

func (cb IPCFailure) CallbackID() SteamCallbackID {
	return IPCFailureID
}

func (cb IPCFailure) Name() string {
	return "IPCFailure"
}

func (cb IPCFailure) String() string {
	return CallbackString(cb)
}

func (cb LicensesUpdated) CallbackID() SteamCallbackID {
	return LicensesUpdatedID
}

func (cb LicensesUpdated) Name() string {
	return "LicensesUpdated"
}

func (cb LicensesUpdated) String() string {
	return CallbackString(cb)
}

func (cb ValidateAuthTicketResponse) CallbackID() SteamCallbackID {
	return ValidateAuthTicketResponseID
}

func (cb ValidateAuthTicketResponse) Name() string {
	return "ValidateAuthTicketResponse"
}

func (cb ValidateAuthTicketResponse) String() string {
	return CallbackString(cb)
}

func (cb MicroTxnAuthorizationResponse) CallbackID() SteamCallbackID {
	return MicroTxnAuthorizationResponseID
}

func (cb MicroTxnAuthorizationResponse) Name() string {
	return "MicroTxnAuthorizationResponse"
}

func (cb MicroTxnAuthorizationResponse) String() string {
	return CallbackString(cb)
}

func (cb EncryptedAppTicketResponse) CallbackID() SteamCallbackID {
	return EncryptedAppTicketResponseID
}

func (cb EncryptedAppTicketResponse) Name() string {
	return "EncryptedAppTicketResponse"
}

func (cb EncryptedAppTicketResponse) String() string {
	return CallbackString(cb)
}

func (cb GetAuthSessionTicketResponse) CallbackID() SteamCallbackID {
	return GetAuthSessionTicketResponseID
}

func (cb GetAuthSessionTicketResponse) Name() string {
	return "GetAuthSessionTicketResponse"
}

func (cb GetAuthSessionTicketResponse) String() string {
	return CallbackString(cb)
}

func (cb GameWebCallback) CallbackID() SteamCallbackID {
	return GameWebCallbackID
}

func (cb GameWebCallback) Name() string {
	return "GameWebCallback"
}

func (cb GameWebCallback) String() string {
	return CallbackString(cb)
}

func (cb StoreAuthURLResponse) CallbackID() SteamCallbackID {
	return StoreAuthURLResponseID
}

func (cb StoreAuthURLResponse) Name() string {
	return "StoreAuthURLResponse"
}

func (cb StoreAuthURLResponse) String() string {
	return CallbackString(cb)
}

func (cb MarketEligibilityResponse) CallbackID() SteamCallbackID {
	return MarketEligibilityResponseID
}

func (cb MarketEligibilityResponse) Name() string {
	return "MarketEligibilityResponse"
}

func (cb MarketEligibilityResponse) String() string {
	return CallbackString(cb)
}

func (cb DurationControl) CallbackID() SteamCallbackID {
	return DurationControlID
}

func (cb DurationControl) Name() string {
	return "DurationControl"
}

func (cb DurationControl) String() string {
	return CallbackString(cb)
}

func (cb GetTicketForWebApiResponse) CallbackID() SteamCallbackID {
	return GetTicketForWebApiResponseID
}

func (cb GetTicketForWebApiResponse) Name() string {
	return "GetTicketForWebApiResponse"
}

func (cb GetTicketForWebApiResponse) String() string {
	return CallbackString(cb)
}
