package apps

import . "github.com/assemblaj/purego-steamworks"

const (
	DlcInstalledID                  SteamCallbackID = 1005
	NewUrlLaunchParametersID        SteamCallbackID = 1014
	AppProofOfPurchaseKeyResponseID SteamCallbackID = 1021
	FileDetailsResultID             SteamCallbackID = 1023
	TimedTrialStatusID              SteamCallbackID = 1030
)

type DlcInstalled struct {
	AppID AppId
}
type NewUrlLaunchParameters struct {
}
type AppProofOfPurchaseKeyResponse struct {
	Result    EResult
	AppID     uint32
	KeyLength uint32
	Key       [240]byte
}

type FileDetailsResult struct {
	Result   EResult
	_        [4]byte
	FileSize uint64
	FileSHA  [20]uint8
	Flags    uint32
}

type TimedTrialStatus struct {
	AppID          AppId
	IsOffline      bool
	_              [3]byte
	SecondsAllowed uint32
	SecondsPlayed  uint32
}

func (cb DlcInstalled) CallbackID() SteamCallbackID {
	return DlcInstalledID
}

func (cb DlcInstalled) Name() string {
	return "DlcInstalled"
}

func (cb DlcInstalled) String() string {
	return CallbackString(cb)
}

func (cb NewUrlLaunchParameters) CallbackID() SteamCallbackID {
	return NewUrlLaunchParametersID
}

func (cb NewUrlLaunchParameters) Name() string {
	return "NewUrlLaunchParameters"
}

func (cb NewUrlLaunchParameters) String() string {
	return CallbackString(cb)
}

func (cb AppProofOfPurchaseKeyResponse) CallbackID() SteamCallbackID {
	return AppProofOfPurchaseKeyResponseID
}

func (cb AppProofOfPurchaseKeyResponse) Name() string {
	return "AppProofOfPurchaseKeyResponse"
}

func (cb AppProofOfPurchaseKeyResponse) String() string {
	return CallbackString(cb)
}

func (cb FileDetailsResult) CallbackID() SteamCallbackID {
	return FileDetailsResultID
}

func (cb FileDetailsResult) Name() string {
	return "FileDetailsResult"
}

func (cb FileDetailsResult) String() string {
	return CallbackString(cb)
}

func (cb TimedTrialStatus) CallbackID() SteamCallbackID {
	return TimedTrialStatusID
}

func (cb TimedTrialStatus) Name() string {
	return "TimedTrialStatus"
}

func (cb TimedTrialStatus) String() string {
	return CallbackString(cb)
}
