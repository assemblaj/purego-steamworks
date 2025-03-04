package utils

import . "github.com/assemblaj/purego-steamworks"

const (
	IPCountryID                         SteamCallbackID = 701
	LowBatteryPowerID                   SteamCallbackID = 702
	SteamShutdownID                     SteamCallbackID = 704
	CheckFileSignatureID                SteamCallbackID = 705
	GamepadTextInputDismissedID         SteamCallbackID = 714
	AppResumingFromSuspendID            SteamCallbackID = 736
	FloatingGamepadTextInputDismissedID SteamCallbackID = 738
	FilterTextDictionaryChangedID       SteamCallbackID = 739
)

type IPCountry struct {
}
type LowBatteryPower struct {
	MinutesBatteryLeft uint8
}

type SteamShutdown struct {
}
type ECheckFileSignature int32

const (
	ECheckFileSignatureInvalidSignature             ECheckFileSignature = 0
	ECheckFileSignatureValidSignature               ECheckFileSignature = 1
	ECheckFileSignatureFileNotFound                 ECheckFileSignature = 2
	ECheckFileSignatureNoSignaturesFoundForThisApp  ECheckFileSignature = 3
	ECheckFileSignatureNoSignaturesFoundForThisFile ECheckFileSignature = 4
)

type CheckFileSignatureResult struct {
	CheckFileSignature ECheckFileSignature
}
type GamepadTextInputDismissed struct {
	Submitted     bool
	SubmittedText uint32
	AppID         AppId
}
type AppResumingFromSuspend struct {
}
type FloatingGamepadTextInputDismissed struct {
}
type FilterTextDictionaryChanged struct {
	Language int32
}

func (cb IPCountry) CallbackID() SteamCallbackID {
	return IPCountryID
}

func (cb IPCountry) Name() string {
	return "IPCountry"
}

func (cb IPCountry) String() string {
	return CallbackString(cb)
}

func (cb LowBatteryPower) CallbackID() SteamCallbackID {
	return LowBatteryPowerID
}

func (cb LowBatteryPower) Name() string {
	return "LowBatteryPower"
}

func (cb LowBatteryPower) String() string {
	return CallbackString(cb)
}

func (cb SteamShutdown) CallbackID() SteamCallbackID {
	return SteamShutdownID
}

func (cb SteamShutdown) Name() string {
	return "SteamShutdown"
}

func (cb SteamShutdown) String() string {
	return CallbackString(cb)
}

func (cb CheckFileSignatureResult) CallbackID() SteamCallbackID {
	return CheckFileSignatureID
}

func (cb CheckFileSignatureResult) Name() string {
	return "CheckFileSignature"
}

func (cb CheckFileSignatureResult) String() string {
	return CallbackString(cb)
}

func (cb GamepadTextInputDismissed) CallbackID() SteamCallbackID {
	return GamepadTextInputDismissedID
}

func (cb GamepadTextInputDismissed) Name() string {
	return "GamepadTextInputDismissed"
}

func (cb GamepadTextInputDismissed) String() string {
	return CallbackString(cb)
}

func (cb AppResumingFromSuspend) CallbackID() SteamCallbackID {
	return AppResumingFromSuspendID
}

func (cb AppResumingFromSuspend) Name() string {
	return "AppResumingFromSuspend"
}

func (cb AppResumingFromSuspend) String() string {
	return CallbackString(cb)
}

func (cb FloatingGamepadTextInputDismissed) CallbackID() SteamCallbackID {
	return FloatingGamepadTextInputDismissedID
}

func (cb FloatingGamepadTextInputDismissed) Name() string {
	return "FloatingGamepadTextInputDismissed"
}

func (cb FloatingGamepadTextInputDismissed) String() string {
	return CallbackString(cb)
}

func (cb FilterTextDictionaryChanged) CallbackID() SteamCallbackID {
	return FilterTextDictionaryChangedID
}

func (cb FilterTextDictionaryChanged) Name() string {
	return "FilterTextDictionaryChanged"
}

func (cb FilterTextDictionaryChanged) String() string {
	return CallbackString(cb)
}
