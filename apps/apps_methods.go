package apps

import (
	"strings"

	. "github.com/assemblaj/purego-steamworks"
)

const (
	AppProofOfPurchaseKeyMax int32 = 240
)

const (
	flatAPI_SteamApps                                 = "SteamAPI_SteamApps_v008"
	flatAPI_ISteamApps_BIsSubscribed                  = "SteamAPI_ISteamApps_BIsSubscribed"
	flatAPI_ISteamApps_BIsLowViolence                 = "SteamAPI_ISteamApps_BIsLowViolence"
	flatAPI_ISteamApps_BIsCybercafe                   = "SteamAPI_ISteamApps_BIsCybercafe"
	flatAPI_ISteamApps_BIsVACBanned                   = "SteamAPI_ISteamApps_BIsVACBanned"
	flatAPI_ISteamApps_GetCurrentGameLanguage         = "SteamAPI_ISteamApps_GetCurrentGameLanguage"
	flatAPI_ISteamApps_GetAvailableGameLanguages      = "SteamAPI_ISteamApps_GetAvailableGameLanguages"
	flatAPI_ISteamApps_BIsSubscribedApp               = "SteamAPI_ISteamApps_BIsSubscribedApp"
	flatAPI_ISteamApps_BIsDlcInstalled                = "SteamAPI_ISteamApps_BIsDlcInstalled"
	flatAPI_ISteamApps_GetEarliestPurchaseUnixTime    = "SteamAPI_ISteamApps_GetEarliestPurchaseUnixTime"
	flatAPI_ISteamApps_BIsSubscribedFromFreeWeekend   = "SteamAPI_ISteamApps_BIsSubscribedFromFreeWeekend"
	flatAPI_ISteamApps_GetDLCCount                    = "SteamAPI_ISteamApps_GetDLCCount"
	flatAPI_ISteamApps_BGetDLCDataByIndex             = "SteamAPI_ISteamApps_BGetDLCDataByIndex"
	flatAPI_ISteamApps_InstallDLC                     = "SteamAPI_ISteamApps_InstallDLC"
	flatAPI_ISteamApps_UninstallDLC                   = "SteamAPI_ISteamApps_UninstallDLC"
	flatAPI_ISteamApps_RequestAppProofOfPurchaseKey   = "SteamAPI_ISteamApps_RequestAppProofOfPurchaseKey"
	flatAPI_ISteamApps_GetCurrentBetaName             = "SteamAPI_ISteamApps_GetCurrentBetaName"
	flatAPI_ISteamApps_MarkContentCorrupt             = "SteamAPI_ISteamApps_MarkContentCorrupt"
	flatAPI_ISteamApps_GetInstalledDepots             = "SteamAPI_ISteamApps_GetInstalledDepots"
	flatAPI_ISteamApps_GetAppInstallDir               = "SteamAPI_ISteamApps_GetAppInstallDir"
	flatAPI_ISteamApps_BIsAppInstalled                = "SteamAPI_ISteamApps_BIsAppInstalled"
	flatAPI_ISteamApps_GetAppOwner                    = "SteamAPI_ISteamApps_GetAppOwner"
	flatAPI_ISteamApps_GetLaunchQueryParam            = "SteamAPI_ISteamApps_GetLaunchQueryParam"
	flatAPI_ISteamApps_GetDlcDownloadProgress         = "SteamAPI_ISteamApps_GetDlcDownloadProgress"
	flatAPI_ISteamApps_GetAppBuildId                  = "SteamAPI_ISteamApps_GetAppBuildId"
	flatAPI_ISteamApps_RequestAllProofOfPurchaseKeys  = "SteamAPI_ISteamApps_RequestAllProofOfPurchaseKeys"
	flatAPI_ISteamApps_GetFileDetails                 = "SteamAPI_ISteamApps_GetFileDetails"
	flatAPI_ISteamApps_GetLaunchCommandLine           = "SteamAPI_ISteamApps_GetLaunchCommandLine"
	flatAPI_ISteamApps_BIsSubscribedFromFamilySharing = "SteamAPI_ISteamApps_BIsSubscribedFromFamilySharing"
	flatAPI_ISteamApps_BIsTimedTrial                  = "SteamAPI_ISteamApps_BIsTimedTrial"
	flatAPI_ISteamApps_SetDlcContext                  = "SteamAPI_ISteamApps_SetDlcContext"
	flatAPI_ISteamApps_GetNumBetas                    = "SteamAPI_ISteamApps_GetNumBetas"
	flatAPI_ISteamApps_GetBetaInfo                    = "SteamAPI_ISteamApps_GetBetaInfo"
	flatAPI_ISteamApps_SetActiveBeta                  = "SteamAPI_ISteamApps_SetActiveBeta"
)

var (
	steamApps_init                            func() uintptr
	iSteamApps_BIsSubscribed                  func(steamApps uintptr) bool
	iSteamApps_BIsLowViolence                 func(steamApps uintptr) bool
	iSteamApps_BIsCybercafe                   func(steamApps uintptr) bool
	iSteamApps_BIsVACBanned                   func(steamApps uintptr) bool
	iSteamApps_GetCurrentGameLanguage         func(steamApps uintptr) string
	iSteamApps_GetAvailableGameLanguages      func(steamApps uintptr) string
	iSteamApps_BIsSubscribedApp               func(steamApps uintptr, appID AppId) bool
	iSteamApps_BIsDlcInstalled                func(steamApps uintptr, appID AppId) bool
	iSteamApps_GetEarliestPurchaseUnixTime    func(steamApps uintptr, appID AppId) uint32
	iSteamApps_BIsSubscribedFromFreeWeekend   func(steamApps uintptr) bool
	iSteamApps_GetDLCCount                    func(steamApps uintptr) int32
	iSteamApps_BGetDLCDataByIndex             func(steamApps uintptr, dlcIndex int32, appID *AppId, available *bool, name *[]byte, nameBufferSize int32) bool
	iSteamApps_InstallDLC                     func(steamApps uintptr, appID AppId)
	iSteamApps_UninstallDLC                   func(steamApps uintptr, appID AppId)
	iSteamApps_RequestAppProofOfPurchaseKey   func(steamApps uintptr, appID AppId)
	iSteamApps_GetCurrentBetaName             func(steamApps uintptr, name *[]byte, nameBufferSize int32) bool
	iSteamApps_MarkContentCorrupt             func(steamApps uintptr, missingFilesOnly bool) bool
	iSteamApps_GetInstalledDepots             func(steamApps uintptr, appID AppId, depots *[]DepotId, maxDepots uint32) uint32
	iSteamApps_GetAppInstallDir               func(steamApps uintptr, appID AppId, folder *[]byte, folderBufferSize uint32) uint32
	iSteamApps_BIsAppInstalled                func(steamApps uintptr, appID AppId) bool
	iSteamApps_GetAppOwner                    func(steamApps uintptr) Uint64SteamID
	iSteamApps_GetLaunchQueryParam            func(steamApps uintptr, key string) string
	iSteamApps_GetDlcDownloadProgress         func(steamApps uintptr, appID AppId, bytesDownloaded *uint64, bytesTotal *uint64) bool
	iSteamApps_GetAppBuildId                  func(steamApps uintptr) int32
	iSteamApps_RequestAllProofOfPurchaseKeys  func(steamApps uintptr)
	iSteamApps_GetFileDetails                 func(steamApps uintptr, pszFileName string) SteamAPICall
	iSteamApps_GetLaunchCommandLine           func(steamApps uintptr, pszCommandLine *[]byte, commandLineStringSize int32) int32
	iSteamApps_BIsSubscribedFromFamilySharing func(steamApps uintptr) bool
	iSteamApps_BIsTimedTrial                  func(steamApps uintptr, secondsAllowed *uint32, secondsPlayed *uint32) bool
	iSteamApps_SetDlcContext                  func(steamApps uintptr, appID AppId) bool
	iSteamApps_GetNumBetas                    func(steamApps uintptr, available *int32, private *int32) int32
	iSteamApps_GetBetaInfo                    func(steamApps uintptr, betaIndex int32, punFlags *uint32, punBuildID *uint32, pchBetaName *[]byte, cchBetaName int32, pchDescription *[]byte, cchDescription int32) bool
	iSteamApps_SetActiveBeta                  func(steamApps uintptr, pchBetaName string) bool
)

var steamApps uintptr

func apps() uintptr {
	if steamApps == 0 {
		steamApps = steamApps_init()
		return steamApps
	}
	return steamApps
}

func BIsSubscribed() bool {
	return iSteamApps_BIsSubscribed(apps())
}

func BIsLowViolence() bool {
	return iSteamApps_BIsLowViolence(apps())
}

func BIsCybercafe() bool {
	return iSteamApps_BIsCybercafe(apps())
}

func BIsVACBanned() bool {
	return iSteamApps_BIsVACBanned(apps())
}

func GetCurrentGameLanguage() string {
	return iSteamApps_GetCurrentGameLanguage(apps())
}

func GetAvailableGameLanguages() []string {
	langs := iSteamApps_GetAvailableGameLanguages(apps())
	return strings.Split(langs, ",")
}

func BIsSubscribedApp(appID AppId) bool {
	return iSteamApps_BIsSubscribedApp(apps(), appID)
}

func BIsDlcInstalled(appID AppId) bool {
	return iSteamApps_BIsDlcInstalled(apps(), appID)
}

func GetEarliestPurchaseUnixTime(appID AppId) uint32 {
	return iSteamApps_GetEarliestPurchaseUnixTime(apps(), appID)
}

func BIsSubscribedFromFreeWeekend() bool {
	return iSteamApps_BIsSubscribedFromFreeWeekend(apps())
}

func GetDLCCount() int32 {
	return iSteamApps_GetDLCCount(apps())
}

func BGetDLCDataByIndex(dlcIndex int32) (appID AppId, available bool, name string, success bool) {
	var nameBuf []byte = make([]byte, 0, 4096)
	result := iSteamApps_BGetDLCDataByIndex(apps(), dlcIndex, &appID, &available, &nameBuf, int32(len(name)))
	return appID, available, string(nameBuf), result
}

func InstallDLC(appID AppId) {
	iSteamApps_InstallDLC(apps(), appID)
}

func UninstallDLC(appID AppId) {
	iSteamApps_UninstallDLC(apps(), appID)
}

func RequestAppProofOfPurchaseKey(appID AppId) {
	iSteamApps_RequestAppProofOfPurchaseKey(apps(), appID)
}

func GetCurrentBetaName() (name string, success bool) {
	var nameBufferSize int32 = 2048
	nameBuf := make([]byte, 0, nameBufferSize)
	success = iSteamApps_GetCurrentBetaName(apps(), &nameBuf, nameBufferSize)
	return string(name), success
}

func MarkContentCorrupt(missingFilesOnly bool) bool {
	return iSteamApps_MarkContentCorrupt(apps(), missingFilesOnly)
}

func GetInstalledDepots(appID AppId, maxDepots uint32) []DepotId {
	var depots []DepotId = make([]DepotId, 0, maxDepots)
	result := iSteamApps_GetInstalledDepots(apps(), appID, &depots, maxDepots)
	return depots[:result]
}

func GetAppInstallDir(appID AppId) string {
	var folderBufferSize uint32 = 2048
	var folder []byte = make([]byte, 0, folderBufferSize)
	total := iSteamApps_GetAppInstallDir(apps(), appID, &folder, folderBufferSize)
	return string(folder[:total])
}

func BIsAppInstalled(appID AppId) bool {
	return iSteamApps_BIsAppInstalled(apps(), appID)
}

func GetAppOwner() Uint64SteamID {
	return iSteamApps_GetAppOwner(apps())
}

func GetLaunchQueryParam(key string) string {
	return iSteamApps_GetLaunchQueryParam(apps(), key)
}

func GetDlcDownloadProgress(appID AppId) (bytesDownloaded uint64, bytesTotal uint64, success bool) {
	success = iSteamApps_GetDlcDownloadProgress(apps(), appID, &bytesDownloaded, &bytesTotal)
	return bytesDownloaded, bytesTotal, success
}

func GetAppBuildId() int32 {
	return iSteamApps_GetAppBuildId(apps())
}

func RequestAllProofOfPurchaseKeys() {
	iSteamApps_RequestAllProofOfPurchaseKeys(apps())
}

func GetFileDetails(pszFileName string) CallResult[FileDetailsResult] {
	handle := iSteamApps_GetFileDetails(apps(), pszFileName)
	return CallResult[FileDetailsResult]{Handle: handle}
}

func GetLaunchCommandLine() string {
	var commandLineStringSize int32 = 256
	var pszCommandLine []byte = make([]byte, 0, commandLineStringSize)
	total := iSteamApps_GetLaunchCommandLine(apps(), &pszCommandLine, commandLineStringSize)
	return string(pszCommandLine[:total])
}

func BIsSubscribedFromFamilySharing() bool {
	return iSteamApps_BIsSubscribedFromFamilySharing(apps())
}

func BIsTimedTrial() (inTimeTrial bool, secondsAllowed uint32, secondsPlayed uint32) {
	inTimeTrial = iSteamApps_BIsTimedTrial(apps(), &secondsAllowed, &secondsPlayed)
	return inTimeTrial, secondsAllowed, secondsPlayed
}

func SetDlcContext(appID AppId) bool {
	return iSteamApps_SetDlcContext(apps(), appID)
}

func GetNumBetas() (available int32, private int32, totalAppBranches int32) {
	totalAppBranches = iSteamApps_GetNumBetas(apps(), &available, &private)
	return available, private, totalAppBranches
}

func GetBetaInfo(betaIndex int32, betaNameSize int32, descriptionSize int32) (punFlags uint32, punBuildID uint32, pchBetaName string, pchDescription string, success bool) {
	var betaName []byte = make([]byte, 0, betaNameSize)
	var description []byte = make([]byte, 0, descriptionSize)

	success = iSteamApps_GetBetaInfo(apps(), betaIndex, &punFlags, &punBuildID, &betaName, int32(len(betaName)), &description, int32(len(description)))
	pchBetaName = string(betaName)
	pchDescription = string(description)
	return punFlags, punBuildID, pchBetaName, pchDescription, success
}

func SetActiveBeta(pchBetaName string) bool {
	return iSteamApps_SetActiveBeta(apps(), pchBetaName)
}
