package utils

import (
	"sync"

	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

var once sync.Once

func init() {
	purego.RegisterLibFunc(&steamUtils_init, SteamAPIDLL, flatAPI_SteamUtils)
	purego.RegisterLibFunc(&iSteamUtils_GetSecondsSinceAppActive, SteamAPIDLL, flatAPI_ISteamUtils_GetSecondsSinceAppActive)
	purego.RegisterLibFunc(&iSteamUtils_GetSecondsSinceComputerActive, SteamAPIDLL, flatAPI_ISteamUtils_GetSecondsSinceComputerActive)
	purego.RegisterLibFunc(&iSteamUtils_GetConnectedUniverse, SteamAPIDLL, flatAPI_ISteamUtils_GetConnectedUniverse)
	purego.RegisterLibFunc(&iSteamUtils_GetServerRealTime, SteamAPIDLL, flatAPI_ISteamUtils_GetServerRealTime)
	purego.RegisterLibFunc(&iSteamUtils_GetIPCountry, SteamAPIDLL, flatAPI_ISteamUtils_GetIPCountry)
	purego.RegisterLibFunc(&iSteamUtils_GetImageSize, SteamAPIDLL, flatAPI_ISteamUtils_GetImageSize)
	purego.RegisterLibFunc(&iSteamUtils_GetImageRGBA, SteamAPIDLL, flatAPI_ISteamUtils_GetImageRGBA)
	purego.RegisterLibFunc(&iSteamUtils_GetCurrentBatteryPower, SteamAPIDLL, flatAPI_ISteamUtils_GetCurrentBatteryPower)
	purego.RegisterLibFunc(&iSteamUtils_GetAppID, SteamAPIDLL, flatAPI_ISteamUtils_GetAppID)
	purego.RegisterLibFunc(&iSteamUtils_SetOverlayNotificationPosition, SteamAPIDLL, flatAPI_ISteamUtils_SetOverlayNotificationPosition)
	purego.RegisterLibFunc(&iSteamUtils_IsAPICallCompleted, SteamAPIDLL, flatAPI_ISteamUtils_IsAPICallCompleted)
	purego.RegisterLibFunc(&iSteamUtils_GetAPICallFailureReason, SteamAPIDLL, flatAPI_ISteamUtils_GetAPICallFailureReason)
	purego.RegisterLibFunc(&iSteamUtils_GetAPICallResult, SteamAPIDLL, flatAPI_ISteamUtils_GetAPICallResult)
	purego.RegisterLibFunc(&iSteamUtils_GetIPCCallCount, SteamAPIDLL, flatAPI_ISteamUtils_GetIPCCallCount)
	purego.RegisterLibFunc(&iSteamUtils_SetWarningMessageHook, SteamAPIDLL, flatAPI_ISteamUtils_SetWarningMessageHook)
	purego.RegisterLibFunc(&iSteamUtils_IsOverlayEnabled, SteamAPIDLL, flatAPI_ISteamUtils_IsOverlayEnabled)
	purego.RegisterLibFunc(&iSteamUtils_BOverlayNeedsPresent, SteamAPIDLL, flatAPI_ISteamUtils_BOverlayNeedsPresent)
	purego.RegisterLibFunc(&iSteamUtils_CheckFileSignature, SteamAPIDLL, flatAPI_ISteamUtils_CheckFileSignature)
	purego.RegisterLibFunc(&iSteamUtils_ShowGamepadTextInput, SteamAPIDLL, flatAPI_ISteamUtils_ShowGamepadTextInput)
	purego.RegisterLibFunc(&iSteamUtils_GetEnteredGamepadTextLength, SteamAPIDLL, flatAPI_ISteamUtils_GetEnteredGamepadTextLength)
	purego.RegisterLibFunc(&iSteamUtils_GetEnteredGamepadTextInput, SteamAPIDLL, flatAPI_ISteamUtils_GetEnteredGamepadTextInput)
	purego.RegisterLibFunc(&iSteamUtils_GetSteamUILanguage, SteamAPIDLL, flatAPI_ISteamUtils_GetSteamUILanguage)
	purego.RegisterLibFunc(&iSteamUtils_IsSteamRunningInVR, SteamAPIDLL, flatAPI_ISteamUtils_IsSteamRunningInVR)
	purego.RegisterLibFunc(&iSteamUtils_SetOverlayNotificationInset, SteamAPIDLL, flatAPI_ISteamUtils_SetOverlayNotificationInset)
	purego.RegisterLibFunc(&iSteamUtils_IsSteamInBigPictureMode, SteamAPIDLL, flatAPI_ISteamUtils_IsSteamInBigPictureMode)
	purego.RegisterLibFunc(&iSteamUtils_StartVRDashboard, SteamAPIDLL, flatAPI_ISteamUtils_StartVRDashboard)
	purego.RegisterLibFunc(&iSteamUtils_IsVRHeadsetStreamingEnabled, SteamAPIDLL, flatAPI_ISteamUtils_IsVRHeadsetStreamingEnabled)
	purego.RegisterLibFunc(&iSteamUtils_SetVRHeadsetStreamingEnabled, SteamAPIDLL, flatAPI_ISteamUtils_SetVRHeadsetStreamingEnabled)
	purego.RegisterLibFunc(&iSteamUtils_IsSteamChinaLauncher, SteamAPIDLL, flatAPI_ISteamUtils_IsSteamChinaLauncher)
	purego.RegisterLibFunc(&iSteamUtils_InitFilterText, SteamAPIDLL, flatAPI_ISteamUtils_InitFilterText)
	purego.RegisterLibFunc(&iSteamUtils_FilterText, SteamAPIDLL, flatAPI_ISteamUtils_FilterText)
	purego.RegisterLibFunc(&iSteamUtils_GetIPv6ConnectivityState, SteamAPIDLL, flatAPI_ISteamUtils_GetIPv6ConnectivityState)
	purego.RegisterLibFunc(&iSteamUtils_IsSteamRunningOnSteamDeck, SteamAPIDLL, flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck)
	purego.RegisterLibFunc(&iSteamUtils_ShowFloatingGamepadTextInput, SteamAPIDLL, flatAPI_ISteamUtils_ShowFloatingGamepadTextInput)
	purego.RegisterLibFunc(&iSteamUtils_SetGameLauncherMode, SteamAPIDLL, flatAPI_ISteamUtils_SetGameLauncherMode)
	purego.RegisterLibFunc(&iSteamUtils_DismissFloatingGamepadTextInput, SteamAPIDLL, flatAPI_ISteamUtils_DismissFloatingGamepadTextInput)
	purego.RegisterLibFunc(&iSteamUtils_DismissGamepadTextInput, SteamAPIDLL, flatAPI_ISteamUtils_DismissGamepadTextInput)

	steamUtils = steamUtils_init()
}
