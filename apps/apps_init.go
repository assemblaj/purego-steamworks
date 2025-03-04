package apps

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamApps_init, SteamAPIDLL, flatAPI_SteamApps)
	purego.RegisterLibFunc(&iSteamApps_BIsSubscribed, SteamAPIDLL, flatAPI_ISteamApps_BIsSubscribed)
	purego.RegisterLibFunc(&iSteamApps_BIsLowViolence, SteamAPIDLL, flatAPI_ISteamApps_BIsLowViolence)
	purego.RegisterLibFunc(&iSteamApps_BIsCybercafe, SteamAPIDLL, flatAPI_ISteamApps_BIsCybercafe)
	purego.RegisterLibFunc(&iSteamApps_BIsVACBanned, SteamAPIDLL, flatAPI_ISteamApps_BIsVACBanned)
	purego.RegisterLibFunc(&iSteamApps_GetCurrentGameLanguage, SteamAPIDLL, flatAPI_ISteamApps_GetCurrentGameLanguage)
	purego.RegisterLibFunc(&iSteamApps_GetAvailableGameLanguages, SteamAPIDLL, flatAPI_ISteamApps_GetAvailableGameLanguages)
	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedApp, SteamAPIDLL, flatAPI_ISteamApps_BIsSubscribedApp)
	purego.RegisterLibFunc(&iSteamApps_BIsDlcInstalled, SteamAPIDLL, flatAPI_ISteamApps_BIsDlcInstalled)
	purego.RegisterLibFunc(&iSteamApps_GetEarliestPurchaseUnixTime, SteamAPIDLL, flatAPI_ISteamApps_GetEarliestPurchaseUnixTime)
	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedFromFreeWeekend, SteamAPIDLL, flatAPI_ISteamApps_BIsSubscribedFromFreeWeekend)
	purego.RegisterLibFunc(&iSteamApps_GetDLCCount, SteamAPIDLL, flatAPI_ISteamApps_GetDLCCount)
	purego.RegisterLibFunc(&iSteamApps_BGetDLCDataByIndex, SteamAPIDLL, flatAPI_ISteamApps_BGetDLCDataByIndex)
	purego.RegisterLibFunc(&iSteamApps_InstallDLC, SteamAPIDLL, flatAPI_ISteamApps_InstallDLC)
	purego.RegisterLibFunc(&iSteamApps_UninstallDLC, SteamAPIDLL, flatAPI_ISteamApps_UninstallDLC)
	purego.RegisterLibFunc(&iSteamApps_RequestAppProofOfPurchaseKey, SteamAPIDLL, flatAPI_ISteamApps_RequestAppProofOfPurchaseKey)
	purego.RegisterLibFunc(&iSteamApps_GetCurrentBetaName, SteamAPIDLL, flatAPI_ISteamApps_GetCurrentBetaName)
	purego.RegisterLibFunc(&iSteamApps_MarkContentCorrupt, SteamAPIDLL, flatAPI_ISteamApps_MarkContentCorrupt)
	purego.RegisterLibFunc(&iSteamApps_GetInstalledDepots, SteamAPIDLL, flatAPI_ISteamApps_GetInstalledDepots)
	purego.RegisterLibFunc(&iSteamApps_GetAppInstallDir, SteamAPIDLL, flatAPI_ISteamApps_GetAppInstallDir)
	purego.RegisterLibFunc(&iSteamApps_BIsAppInstalled, SteamAPIDLL, flatAPI_ISteamApps_BIsAppInstalled)
	purego.RegisterLibFunc(&iSteamApps_GetAppOwner, SteamAPIDLL, flatAPI_ISteamApps_GetAppOwner)
	purego.RegisterLibFunc(&iSteamApps_GetLaunchQueryParam, SteamAPIDLL, flatAPI_ISteamApps_GetLaunchQueryParam)
	purego.RegisterLibFunc(&iSteamApps_GetDlcDownloadProgress, SteamAPIDLL, flatAPI_ISteamApps_GetDlcDownloadProgress)
	purego.RegisterLibFunc(&iSteamApps_GetAppBuildId, SteamAPIDLL, flatAPI_ISteamApps_GetAppBuildId)
	purego.RegisterLibFunc(&iSteamApps_RequestAllProofOfPurchaseKeys, SteamAPIDLL, flatAPI_ISteamApps_RequestAllProofOfPurchaseKeys)
	purego.RegisterLibFunc(&iSteamApps_GetFileDetails, SteamAPIDLL, flatAPI_ISteamApps_GetFileDetails)
	purego.RegisterLibFunc(&iSteamApps_GetLaunchCommandLine, SteamAPIDLL, flatAPI_ISteamApps_GetLaunchCommandLine)
	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedFromFamilySharing, SteamAPIDLL, flatAPI_ISteamApps_BIsSubscribedFromFamilySharing)
	purego.RegisterLibFunc(&iSteamApps_BIsTimedTrial, SteamAPIDLL, flatAPI_ISteamApps_BIsTimedTrial)
	purego.RegisterLibFunc(&iSteamApps_SetDlcContext, SteamAPIDLL, flatAPI_ISteamApps_SetDlcContext)
	purego.RegisterLibFunc(&iSteamApps_GetNumBetas, SteamAPIDLL, flatAPI_ISteamApps_GetNumBetas)
	purego.RegisterLibFunc(&iSteamApps_GetBetaInfo, SteamAPIDLL, flatAPI_ISteamApps_GetBetaInfo)
	purego.RegisterLibFunc(&iSteamApps_SetActiveBeta, SteamAPIDLL, flatAPI_ISteamApps_SetActiveBeta)
}
