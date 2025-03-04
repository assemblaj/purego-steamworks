package utils

import (
	. "github.com/assemblaj/purego-steamworks"
)

type ESteamIPv6ConnectivityState int32

const (
	ESteamIPv6ConnectivityStateUnknown ESteamIPv6ConnectivityState = 0
	ESteamIPv6ConnectivityStateGood    ESteamIPv6ConnectivityState = 1
	ESteamIPv6ConnectivityStateBad     ESteamIPv6ConnectivityState = 2
)

type EFloatingGamepadTextInputMode int32

const (
	EFloatingGamepadTextInputModeModeSingleLine    EFloatingGamepadTextInputMode = 0
	EFloatingGamepadTextInputModeModeMultipleLines EFloatingGamepadTextInputMode = 1
	EFloatingGamepadTextInputModeModeEmail         EFloatingGamepadTextInputMode = 2
	EFloatingGamepadTextInputModeModeNumeric       EFloatingGamepadTextInputMode = 3
)

type ETextFilteringContext int32

const (
	ETextFilteringContextUnknown     ETextFilteringContext = 0
	ETextFilteringContextGameContent ETextFilteringContext = 1
	ETextFilteringContextChat        ETextFilteringContext = 2
	ETextFilteringContextName        ETextFilteringContext = 3
)

type ESteamIPv6ConnectivityProtocol int32

const (
	ESteamIPv6ConnectivityProtocolInvalid ESteamIPv6ConnectivityProtocol = 0
	ESteamIPv6ConnectivityProtocolHTTP    ESteamIPv6ConnectivityProtocol = 1
	ESteamIPv6ConnectivityProtocolUDP     ESteamIPv6ConnectivityProtocol = 2
)

type ENotificationPosition int32

const (
	EPositionInvalid     ENotificationPosition = -1
	EPositionTopLeft     ENotificationPosition = 0
	EPositionTopRight    ENotificationPosition = 1
	EPositionBottomLeft  ENotificationPosition = 2
	EPositionBottomRight ENotificationPosition = 3
)

type EGamepadTextInputMode int32

const (
	EGamepadTextInputModeNormal   EGamepadTextInputMode = 0
	EGamepadTextInputModePassword EGamepadTextInputMode = 1
)

type EGamepadTextInputLineMode int32

const (
	EGamepadTextInputLineModeSingleLine    EGamepadTextInputLineMode = 0
	EGamepadTextInputLineModeMultipleLines EGamepadTextInputLineMode = 1
)

var (
	steamUtils_init                             func() uintptr
	iSteamUtils_GetSecondsSinceAppActive        func(steamUtils uintptr) uint32
	iSteamUtils_GetSecondsSinceComputerActive   func(steamUtils uintptr) uint32
	iSteamUtils_GetConnectedUniverse            func(steamUtils uintptr) EUniverse
	iSteamUtils_GetServerRealTime               func(steamUtils uintptr) uint32
	iSteamUtils_GetIPCountry                    func(steamUtils uintptr) string
	iSteamUtils_GetImageSize                    func(steamUtils uintptr, iImage int32, pnWidth *uint32, pnHeight *uint32) bool
	iSteamUtils_GetImageRGBA                    func(steamUtils uintptr, iImage int32, pubDest *[]byte, nDestBufferSize int32) bool
	iSteamUtils_GetCurrentBatteryPower          func(steamUtils uintptr) uint8
	iSteamUtils_GetAppID                        func(steamUtils uintptr) uint32
	iSteamUtils_SetOverlayNotificationPosition  func(steamUtils uintptr, eNotificationPosition ENotificationPosition)
	iSteamUtils_IsAPICallCompleted              func(steamUtils uintptr, hSteamAPICall SteamAPICall, pbFailed *bool) bool
	iSteamUtils_GetAPICallFailureReason         func(steamUtils uintptr, hSteamAPICall SteamAPICall) ESteamAPICallFailure
	iSteamUtils_GetAPICallResult                func(steamUtils uintptr, hSteamAPICall SteamAPICall, pCallback *[]byte, cubCallback int32, iCallbackExpected int32, pbFailed *bool) bool
	iSteamUtils_GetIPCCallCount                 func(steamUtils uintptr) uint32
	iSteamUtils_SetWarningMessageHook           func(steamUtils uintptr, pFunction SteamAPIWarningMessageHook)
	iSteamUtils_IsOverlayEnabled                func(steamUtils uintptr) bool
	iSteamUtils_BOverlayNeedsPresent            func(steamUtils uintptr) bool
	iSteamUtils_CheckFileSignature              func(steamUtils uintptr, szFileName string) SteamAPICall
	iSteamUtils_ShowGamepadTextInput            func(steamUtils uintptr, eInputMode EGamepadTextInputMode, eLineInputMode EGamepadTextInputLineMode, pchDescription string, unCharMax uint32, pchExistingText string) bool
	iSteamUtils_GetEnteredGamepadTextLength     func(steamUtils uintptr) uint32
	iSteamUtils_GetEnteredGamepadTextInput      func(steamUtils uintptr, pchText *[]byte, cchText uint32) bool
	iSteamUtils_GetSteamUILanguage              func(steamUtils uintptr) string
	iSteamUtils_IsSteamRunningInVR              func(steamUtils uintptr) bool
	iSteamUtils_SetOverlayNotificationInset     func(steamUtils uintptr, nHorizontalInset int32, nVerticalInset int32)
	iSteamUtils_IsSteamInBigPictureMode         func(steamUtils uintptr) bool
	iSteamUtils_StartVRDashboard                func(steamUtils uintptr)
	iSteamUtils_IsVRHeadsetStreamingEnabled     func(steamUtils uintptr) bool
	iSteamUtils_SetVRHeadsetStreamingEnabled    func(steamUtils uintptr, bEnabled bool)
	iSteamUtils_IsSteamChinaLauncher            func(steamUtils uintptr) bool
	iSteamUtils_InitFilterText                  func(steamUtils uintptr, unFilterOptions uint32) bool
	iSteamUtils_FilterText                      func(steamUtils uintptr, eContext ETextFilteringContext, sourceSteamID Uint64SteamID, pchInputMessage string, pchOutFilteredText *[]byte, nByteSizeOutFilteredText uint32) int32
	iSteamUtils_GetIPv6ConnectivityState        func(steamUtils uintptr, eProtocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState
	iSteamUtils_IsSteamRunningOnSteamDeck       func(steamUtils uintptr) bool
	iSteamUtils_ShowFloatingGamepadTextInput    func(steamUtils uintptr, eKeyboardMode EFloatingGamepadTextInputMode, nTextFieldXPosition int32, nTextFieldYPosition int32, nTextFieldWidth int32, nTextFieldHeight int32) bool
	iSteamUtils_SetGameLauncherMode             func(steamUtils uintptr, bLauncherMode bool)
	iSteamUtils_DismissFloatingGamepadTextInput func(steamUtils uintptr) bool
	iSteamUtils_DismissGamepadTextInput         func(steamUtils uintptr) bool
)

const (
	flatAPI_SteamUtils                                  = "SteamAPI_SteamUtils_v010"
	flatAPI_ISteamUtils_GetSecondsSinceAppActive        = "SteamAPI_ISteamUtils_GetSecondsSinceAppActive"
	flatAPI_ISteamUtils_GetSecondsSinceComputerActive   = "SteamAPI_ISteamUtils_GetSecondsSinceComputerActive"
	flatAPI_ISteamUtils_GetConnectedUniverse            = "SteamAPI_ISteamUtils_GetConnectedUniverse"
	flatAPI_ISteamUtils_GetServerRealTime               = "SteamAPI_ISteamUtils_GetServerRealTime"
	flatAPI_ISteamUtils_GetIPCountry                    = "SteamAPI_ISteamUtils_GetIPCountry"
	flatAPI_ISteamUtils_GetImageSize                    = "SteamAPI_ISteamUtils_GetImageSize"
	flatAPI_ISteamUtils_GetImageRGBA                    = "SteamAPI_ISteamUtils_GetImageRGBA"
	flatAPI_ISteamUtils_GetCurrentBatteryPower          = "SteamAPI_ISteamUtils_GetCurrentBatteryPower"
	flatAPI_ISteamUtils_GetAppID                        = "SteamAPI_ISteamUtils_GetAppID"
	flatAPI_ISteamUtils_SetOverlayNotificationPosition  = "SteamAPI_ISteamUtils_SetOverlayNotificationPosition"
	flatAPI_ISteamUtils_IsAPICallCompleted              = "SteamAPI_ISteamUtils_IsAPICallCompleted"
	flatAPI_ISteamUtils_GetAPICallFailureReason         = "SteamAPI_ISteamUtils_GetAPICallFailureReason"
	flatAPI_ISteamUtils_GetAPICallResult                = "SteamAPI_ISteamUtils_GetAPICallResult"
	flatAPI_ISteamUtils_GetIPCCallCount                 = "SteamAPI_ISteamUtils_GetIPCCallCount"
	flatAPI_ISteamUtils_SetWarningMessageHook           = "SteamAPI_ISteamUtils_SetWarningMessageHook"
	flatAPI_ISteamUtils_IsOverlayEnabled                = "SteamAPI_ISteamUtils_IsOverlayEnabled"
	flatAPI_ISteamUtils_BOverlayNeedsPresent            = "SteamAPI_ISteamUtils_BOverlayNeedsPresent"
	flatAPI_ISteamUtils_CheckFileSignature              = "SteamAPI_ISteamUtils_CheckFileSignature"
	flatAPI_ISteamUtils_ShowGamepadTextInput            = "SteamAPI_ISteamUtils_ShowGamepadTextInput"
	flatAPI_ISteamUtils_GetEnteredGamepadTextLength     = "SteamAPI_ISteamUtils_GetEnteredGamepadTextLength"
	flatAPI_ISteamUtils_GetEnteredGamepadTextInput      = "SteamAPI_ISteamUtils_GetEnteredGamepadTextInput"
	flatAPI_ISteamUtils_GetSteamUILanguage              = "SteamAPI_ISteamUtils_GetSteamUILanguage"
	flatAPI_ISteamUtils_IsSteamRunningInVR              = "SteamAPI_ISteamUtils_IsSteamRunningInVR"
	flatAPI_ISteamUtils_SetOverlayNotificationInset     = "SteamAPI_ISteamUtils_SetOverlayNotificationInset"
	flatAPI_ISteamUtils_IsSteamInBigPictureMode         = "SteamAPI_ISteamUtils_IsSteamInBigPictureMode"
	flatAPI_ISteamUtils_StartVRDashboard                = "SteamAPI_ISteamUtils_StartVRDashboard"
	flatAPI_ISteamUtils_IsVRHeadsetStreamingEnabled     = "SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled"
	flatAPI_ISteamUtils_SetVRHeadsetStreamingEnabled    = "SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled"
	flatAPI_ISteamUtils_IsSteamChinaLauncher            = "SteamAPI_ISteamUtils_IsSteamChinaLauncher"
	flatAPI_ISteamUtils_InitFilterText                  = "SteamAPI_ISteamUtils_InitFilterText"
	flatAPI_ISteamUtils_FilterText                      = "SteamAPI_ISteamUtils_FilterText"
	flatAPI_ISteamUtils_GetIPv6ConnectivityState        = "SteamAPI_ISteamUtils_GetIPv6ConnectivityState"
	flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck       = "SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck"
	flatAPI_ISteamUtils_ShowFloatingGamepadTextInput    = "SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput"
	flatAPI_ISteamUtils_SetGameLauncherMode             = "SteamAPI_ISteamUtils_SetGameLauncherMode"
	flatAPI_ISteamUtils_DismissFloatingGamepadTextInput = "SteamAPI_ISteamUtils_DismissFloatingGamepadTextInput"
	flatAPI_ISteamUtils_DismissGamepadTextInput         = "SteamAPI_ISteamUtils_DismissGamepadTextInput"
)

var steamUtils uintptr

func utils() uintptr {
	if steamUtils == 0 {
		steamUtils = steamUtils_init()
		return steamUtils
	}
	return steamUtils
}
func GetSecondsSinceAppActive() uint32 {
	return iSteamUtils_GetSecondsSinceAppActive(utils())
}

func GetSecondsSinceComputerActive() uint32 {
	return iSteamUtils_GetSecondsSinceComputerActive(utils())
}

func GetConnectedUniverse() EUniverse {
	return iSteamUtils_GetConnectedUniverse(utils())
}

func GetServerRealTime() uint32 {
	return iSteamUtils_GetServerRealTime(utils())
}

func GetIPCountry() string {
	return iSteamUtils_GetIPCountry(utils())
}

func GetImageSize(imageIndex int32) (uint32, uint32, bool) {
	var pnWidth, pnHeight uint32
	result := iSteamUtils_GetImageSize(utils(), imageIndex, &pnWidth, &pnHeight)
	return pnWidth, pnHeight, result
}

func GetImageRGBA(imageIndex int32, destBufferSize int32) ([]byte, bool) {
	var pubDest []byte = make([]byte, 0, destBufferSize)
	result := iSteamUtils_GetImageRGBA(utils(), imageIndex, &pubDest, destBufferSize)
	return pubDest, result
}

func GetCurrentBatteryPower() uint8 {
	return iSteamUtils_GetCurrentBatteryPower(utils())
}

func GetAppID() uint32 {
	return iSteamUtils_GetAppID(utils())
}

func SetOverlayNotificationPosition(notificationPosition ENotificationPosition) {
	iSteamUtils_SetOverlayNotificationPosition(utils(), notificationPosition)
}

func IsAPICallCompleted(steamAPICallHandle SteamAPICall) (failed bool, completed bool) {
	var pbFailed bool
	result := iSteamUtils_IsAPICallCompleted(utils(), steamAPICallHandle, &pbFailed)
	return pbFailed, result
}

func GetAPICallFailureReason(steamAPICallHandle SteamAPICall) ESteamAPICallFailure {
	return iSteamUtils_GetAPICallFailureReason(utils(), steamAPICallHandle)
}

func GetAPICallResult(steamAPICallHandle SteamAPICall, callbackDataSize int32, callbackExpectedID int32) (callbackData []byte, failed bool, success bool) {
	callbackData = make([]byte, 0, callbackDataSize)
	success = iSteamUtils_GetAPICallResult(utils(), steamAPICallHandle, &callbackData, callbackDataSize, callbackExpectedID, &failed)
	return callbackData, failed, success
}

func GetIPCCallCount() uint32 {
	return iSteamUtils_GetIPCCallCount(utils())
}

func SetWarningMessageHook(function SteamAPIWarningMessageHook) {
	iSteamUtils_SetWarningMessageHook(utils(), function)
}

func IsOverlayEnabled() bool {
	return iSteamUtils_IsOverlayEnabled(utils())
}

func BOverlayNeedsPresent() bool {
	return iSteamUtils_BOverlayNeedsPresent(utils())
}

func CheckFileSignature(fileName string) CallResult[CheckFileSignatureResult] {
	handle := iSteamUtils_CheckFileSignature(utils(), fileName)
	return CallResult[CheckFileSignatureResult]{Handle: handle}
}

func ShowGamepadTextInput(inputMode EGamepadTextInputMode, lineInputMode EGamepadTextInputLineMode, description string, charMax uint32, existingText string) bool {
	return iSteamUtils_ShowGamepadTextInput(utils(), inputMode, lineInputMode, description, charMax, existingText)
}

func GetEnteredGamepadTextLength() uint32 {
	return iSteamUtils_GetEnteredGamepadTextLength(utils())
}

func GetEnteredGamepadTextInput(text uint32) (string, bool) {
	var pchText []byte = make([]byte, 0, text)
	result := iSteamUtils_GetEnteredGamepadTextInput(utils(), &pchText, text)
	return string(pchText), result
}

func GetSteamUILanguage() string {
	return iSteamUtils_GetSteamUILanguage(utils())
}

func IsSteamRunningInVR() bool {
	return iSteamUtils_IsSteamRunningInVR(utils())
}

func SetOverlayNotificationInset(horizontalInset int32, verticalInset int32) {
	iSteamUtils_SetOverlayNotificationInset(utils(), horizontalInset, verticalInset)
}

func IsSteamInBigPictureMode() bool {
	return iSteamUtils_IsSteamInBigPictureMode(utils())
}

func StartVRDashboard() {
	iSteamUtils_StartVRDashboard(utils())
}

func IsVRHeadsetStreamingEnabled() bool {
	return iSteamUtils_IsVRHeadsetStreamingEnabled(utils())
}

func SetVRHeadsetStreamingEnabled(enabled bool) {
	iSteamUtils_SetVRHeadsetStreamingEnabled(utils(), enabled)
}

func IsSteamChinaLauncher() bool {
	return iSteamUtils_IsSteamChinaLauncher(utils())
}

func InitFilterText(filterOptions uint32) bool {
	return iSteamUtils_InitFilterText(utils(), filterOptions)
}

func FilterText(context ETextFilteringContext, sourceSteamID Uint64SteamID, inputMessage string, byteSizeOutFilteredText uint32) (string, int32) {
	if int(byteSizeOutFilteredText) < len(inputMessage)+1 {
		byteSizeOutFilteredText = uint32(len(inputMessage) + 1)
	}
	var pchOutFilteredText []byte = make([]byte, 0, byteSizeOutFilteredText)
	result := iSteamUtils_FilterText(utils(), context, sourceSteamID, inputMessage, &pchOutFilteredText, byteSizeOutFilteredText)
	return string(pchOutFilteredText), result
}

func GetIPv6ConnectivityState(protocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState {
	return iSteamUtils_GetIPv6ConnectivityState(utils(), protocol)
}

func IsSteamRunningOnSteamDeck() bool {
	return iSteamUtils_IsSteamRunningOnSteamDeck(utils())
}

func ShowFloatingGamepadTextInput(keyboardMode EFloatingGamepadTextInputMode, textFieldXPosition int32, textFieldYPosition int32, textFieldWidth int32, textFieldHeight int32) bool {
	return iSteamUtils_ShowFloatingGamepadTextInput(utils(), keyboardMode, textFieldXPosition, textFieldYPosition, textFieldWidth, textFieldHeight)
}

func SetGameLauncherMode(launcherMode bool) {
	iSteamUtils_SetGameLauncherMode(utils(), launcherMode)
}

func DismissFloatingGamepadTextInput() bool {
	return iSteamUtils_DismissFloatingGamepadTextInput(utils())
}

func DismissGamepadTextInput() bool {
	return iSteamUtils_DismissGamepadTextInput(utils())
}
