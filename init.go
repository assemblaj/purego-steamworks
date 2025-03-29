package puregosteamworks

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

func init() {
	initApi()
	initCallbacks()
	initApps()
	initClient()
	initFriends()
	initGameSearch()
	initGameServer()
	initGameServerStats()
	initInput()
	initInventory()
	initMatchmaking()
	initMatchmakingServer()
	initNetworkTypes()
	initNetworkMessages()
	initNetworkSockets()
	initNetworkUtils()
	initParties()
	initRemotePlay()
	initRemoteStorage()
	initScreenshots()
	initUGC()
	initUser()
	initUserStats()
	initUtils()
}

func getSteamLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "libsteam_api.dylib"
	case "linux":
		return "libsteam_api.so"
	case "windows":
		return "steam_api64.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func initApi() {
	var err error
	steamAPILib, err = openLibrary(getSteamLibrary())
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&restartAppIfNecessary, steamAPILib, flatAPI_RestartAppIfNecessary)
	purego.RegisterLibFunc(&initFlat, steamAPILib, flatAPI_InitFlat)
	purego.RegisterLibFunc(&Shutdown, steamAPILib, flatAPI_Shutdown)
	purego.RegisterLibFunc(&WriteMiniDump, steamAPILib, flatAPI_WriteMiniDump)
	purego.RegisterLibFunc(&SetMiniDumpComment, steamAPILib, flatAPI_SetMiniDumpComment)
	purego.RegisterLibFunc(&IsSteamRunning, steamAPILib, flatAPI_IsSteamRunning)
	purego.RegisterLibFunc(&ReleaseCurrentThreadMemory, steamAPILib, flatAPI_ReleaseCurrentThreadMemory)

	purego.RegisterLibFunc(&gameServerInit, steamAPILib, flatAPI_SteamGameServer_Init)
	purego.RegisterLibFunc(&GameServerShutdown, steamAPILib, flatAPI_SteamGameServer_Shutdown)
	purego.RegisterLibFunc(&GameServerBSecure, steamAPILib, flatAPI_SteamGameServer_BSecure)
	purego.RegisterLibFunc(&GameServerGetSteamID, steamAPILib, flatAPI_SteamGameServer_GetSteamID)

}

var (
	addr_getHSteamPipe                   uintptr
	addr_manualdispatch_init             uintptr
	addr_manualdispatch_runFrame         uintptr
	addr_manualdispatch_getNextCallback  uintptr
	addr_manualdispatch_freeLastCallback uintptr
	addr_manualdispatch_getApiCallResult uintptr
)

func initCallbacks() {
	var err error

	//purego.RegisterLibFunc(&getHSteamPipe, steamAPILib, flatAPI_GetHSteamPipe)
	addr_getHSteamPipe, err = openSymbol(steamAPILib, flatAPI_GetHSteamPipe)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_GetHSteamPipe)
	}

	// purego.RegisterLibFunc(&manualdispatch_init, steamAPILib, flatAPI_ManualDispatch_Init)
	addr_manualdispatch_init, err = openSymbol(steamAPILib, flatAPI_ManualDispatch_Init)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_ManualDispatch_Init)
	}

	//purego.RegisterLibFunc(&manualdispatch_runFrame, steamAPILib, flatAPI_ManualDispatch_RunFrame)
	addr_manualdispatch_runFrame, err = openSymbol(steamAPILib, flatAPI_ManualDispatch_RunFrame)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_ManualDispatch_RunFrame)
	}

	// purego.RegisterLibFunc(&manualdispatch_getNextCallback, steamAPILib, flatAPI_ManualDispatch_GetNextCallback)
	addr_manualdispatch_getNextCallback, err = openSymbol(steamAPILib, flatAPI_ManualDispatch_GetNextCallback)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_ManualDispatch_GetNextCallback)
	}

	//purego.RegisterLibFunc(&manualdispatch_freeLastCallback, steamAPILib, flatAPI_ManualDispatch_FreeLastCallback)
	addr_manualdispatch_freeLastCallback, err = openSymbol(steamAPILib, flatAPI_ManualDispatch_FreeLastCallback)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_ManualDispatch_FreeLastCallback)
	}

	//purego.RegisterLibFunc(&manualdispatch_getApiCallResult, steamAPILib, flatAPI_ManualDispatch_GetAPICallResult)
	addr_manualdispatch_getApiCallResult, err = openSymbol(steamAPILib, flatAPI_ManualDispatch_GetAPICallResult)
	if err != nil {
		panic("cannot openSymbol: " + flatAPI_ManualDispatch_GetAPICallResult)
	}

	getHSteamPipe = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_getHSteamPipe)
		return _r0
	}

	manualdispatch_init = func() {
		purego.SyscallN(addr_manualdispatch_init)
	}

	manualdispatch_runFrame = func(pipe HSteamPipe) {
		purego.SyscallN(addr_manualdispatch_runFrame, uintptr(pipe))
	}

	manualdispatch_getNextCallback = func(pipe HSteamPipe, msg *callbackMsg) bool {
		_r0, _, _ := purego.SyscallN(addr_manualdispatch_getNextCallback, uintptr(pipe), uintptr(unsafe.Pointer(msg)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(msg)
		return __r0
	}

	manualdispatch_freeLastCallback = func(pipe HSteamPipe) {
		purego.SyscallN(addr_manualdispatch_freeLastCallback, uintptr(pipe))
	}

	manualdispatch_getApiCallResult = func(hSteamPipe HSteamPipe, hSteamAPICall SteamAPICall, pCallback []byte, cubCallback int, iCallbackExpected int, pbFailed *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_manualdispatch_getApiCallResult, uintptr(hSteamPipe), uintptr(hSteamAPICall),
			uintptr(unsafe.Pointer(unsafe.SliceData(pCallback))), uintptr(cubCallback), uintptr(iCallbackExpected), uintptr(unsafe.Pointer(pbFailed)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pCallback)
		runtime.KeepAlive(pbFailed)
		return __r0
	}
}

var (
	addr_steamApps_get                             uintptr
	addr_iSteamApps_BIsSubscribed                  uintptr
	addr_iSteamApps_BIsLowViolence                 uintptr
	addr_iSteamApps_BIsCybercafe                   uintptr
	addr_iSteamApps_BIsVACBanned                   uintptr
	addr_iSteamApps_GetCurrentGameLanguage         uintptr
	addr_iSteamApps_GetAvailableGameLanguages      uintptr
	addr_iSteamApps_BIsSubscribedApp               uintptr
	addr_iSteamApps_BIsDlcInstalled                uintptr
	addr_iSteamApps_GetEarliestPurchaseUnixTime    uintptr
	addr_iSteamApps_BIsSubscribedFromFreeWeekend   uintptr
	addr_iSteamApps_GetDLCCount                    uintptr
	addr_iSteamApps_BGetDLCDataByIndex             uintptr
	addr_iSteamApps_InstallDLC                     uintptr
	addr_iSteamApps_UninstallDLC                   uintptr
	addr_iSteamApps_RequestAppProofOfPurchaseKey   uintptr
	addr_iSteamApps_GetCurrentBetaName             uintptr
	addr_iSteamApps_MarkContentCorrupt             uintptr
	addr_iSteamApps_GetInstalledDepots             uintptr
	addr_iSteamApps_GetAppInstallDir               uintptr
	addr_iSteamApps_BIsAppInstalled                uintptr
	addr_iSteamApps_GetAppOwner                    uintptr
	addr_iSteamApps_GetLaunchQueryParam            uintptr
	addr_iSteamApps_GetDlcDownloadProgress         uintptr
	addr_iSteamApps_GetAppBuildId                  uintptr
	addr_iSteamApps_RequestAllProofOfPurchaseKeys  uintptr
	addr_iSteamApps_GetFileDetails                 uintptr
	addr_iSteamApps_GetLaunchCommandLine           uintptr
	addr_iSteamApps_BIsSubscribedFromFamilySharing uintptr
	addr_iSteamApps_BIsTimedTrial                  uintptr
	addr_iSteamApps_SetDlcContext                  uintptr
	addr_iSteamApps_GetNumBetas                    uintptr
	addr_iSteamApps_GetBetaInfo                    uintptr
	addr_iSteamApps_SetActiveBeta                  uintptr
)

func initApps() {
	var err error

	addr_steamApps_get, err = openSymbol(steamAPILib, flatAPI_SteamApps)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamApps)
	}
	addr_iSteamApps_BIsSubscribed, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsSubscribed)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsSubscribed")
	}

	addr_iSteamApps_BIsLowViolence, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsLowViolence)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsLowViolence")
	}

	addr_iSteamApps_BIsCybercafe, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsCybercafe)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsCybercafe")
	}
	addr_iSteamApps_BIsVACBanned, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsVACBanned)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsVACBanned")
	}
	addr_iSteamApps_GetCurrentGameLanguage, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetCurrentGameLanguage)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetCurrentGameLanguage")
	}
	addr_iSteamApps_GetAvailableGameLanguages, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetAvailableGameLanguages)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetAvailableGameLanguages")
	}
	addr_iSteamApps_BIsSubscribedApp, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsSubscribedApp)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsSubscribedApp")
	}
	addr_iSteamApps_BIsDlcInstalled, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsDlcInstalled)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsDlcInstalled")
	}
	addr_iSteamApps_GetEarliestPurchaseUnixTime, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetEarliestPurchaseUnixTime)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetEarliestPurchaseUnixTime")
	}
	addr_iSteamApps_BIsSubscribedFromFreeWeekend, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsSubscribedFromFreeWeekend)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsSubscribedFromFreeWeekend")
	}
	addr_iSteamApps_GetDLCCount, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetDLCCount)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetDLCCount")
	}
	addr_iSteamApps_BGetDLCDataByIndex, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BGetDLCDataByIndex)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BGetDLCDataByIndex")
	}
	addr_iSteamApps_InstallDLC, err = openSymbol(steamAPILib, flatAPI_ISteamApps_InstallDLC)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_InstallDLC")
	}

	addr_iSteamApps_UninstallDLC, err = openSymbol(steamAPILib, flatAPI_ISteamApps_UninstallDLC)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_UninstallDLC")
	}

	addr_iSteamApps_RequestAppProofOfPurchaseKey, err = openSymbol(steamAPILib, flatAPI_ISteamApps_RequestAppProofOfPurchaseKey)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_RequestAppProofOfPurchaseKey")
	}

	addr_iSteamApps_GetCurrentBetaName, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetCurrentBetaName)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps__ISteamApps_GetCurrentBetaName")
	}

	addr_iSteamApps_MarkContentCorrupt, err = openSymbol(steamAPILib, flatAPI_ISteamApps_MarkContentCorrupt)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_MarkContentCorrupt")
	}
	addr_iSteamApps_GetInstalledDepots, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetInstalledDepots)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetInstalledDepots")
	}
	addr_iSteamApps_GetAppInstallDir, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetAppInstallDir)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetAppInstallDir")
	}
	addr_iSteamApps_BIsAppInstalled, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsAppInstalled)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsAppInstalled")
	}
	addr_iSteamApps_GetAppOwner, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetAppOwner)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetAppOwner")
	}
	addr_iSteamApps_GetLaunchQueryParam, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetLaunchQueryParam)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetLaunchQueryParam")
	}
	addr_iSteamApps_GetDlcDownloadProgress, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetDlcDownloadProgress)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetDlcDownloadProgress")
	}
	addr_iSteamApps_GetAppBuildId, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetAppBuildId)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetAppBuildId")
	}
	addr_iSteamApps_RequestAllProofOfPurchaseKeys, err = openSymbol(steamAPILib, flatAPI_ISteamApps_RequestAllProofOfPurchaseKeys)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_RequestAllProofOfPurchaseKeys")
	}

	addr_iSteamApps_GetFileDetails, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetFileDetails)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetFileDetails")
	}

	addr_iSteamApps_GetLaunchCommandLine, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetLaunchCommandLine)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetLaunchCommandLine")
	}
	addr_iSteamApps_BIsSubscribedFromFamilySharing, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsSubscribedFromFamilySharing)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsSubscribedFromFamilySharing")
	}
	addr_iSteamApps_BIsTimedTrial, err = openSymbol(steamAPILib, flatAPI_ISteamApps_BIsTimedTrial)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_BIsTimedTrial")
	}
	addr_iSteamApps_SetDlcContext, err = openSymbol(steamAPILib, flatAPI_ISteamApps_SetDlcContext)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_SetDlcContext")
	}
	addr_iSteamApps_GetNumBetas, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetNumBetas)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetNumBetas")
	}
	addr_iSteamApps_GetBetaInfo, err = openSymbol(steamAPILib, flatAPI_ISteamApps_GetBetaInfo)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_GetBetaInfo")
	}
	addr_iSteamApps_SetActiveBeta, err = openSymbol(steamAPILib, flatAPI_ISteamApps_SetActiveBeta)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamApps_SetActiveBeta")
	}

	// 	purego.RegisterLibFunc(&steamApps_get, steamAPILib, flatAPI_SteamApps)
	steamApps_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamApps_get)
		return _r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsSubscribed, steamAPILib, flatAPI_ISteamApps_BIsSubscribed)
	iSteamApps_BIsSubscribed = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsSubscribed, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsLowViolence, steamAPILib, flatAPI_ISteamApps_BIsLowViolence)
	iSteamApps_BIsLowViolence = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsLowViolence, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsCybercafe, steamAPILib, flatAPI_ISteamApps_BIsCybercafe)
	iSteamApps_BIsCybercafe = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsCybercafe, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsVACBanned, steamAPILib, flatAPI_ISteamApps_BIsVACBanned)
	iSteamApps_BIsVACBanned = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsVACBanned, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetCurrentGameLanguage, steamAPILib, flatAPI_ISteamApps_GetCurrentGameLanguage)
	iSteamApps_GetCurrentGameLanguage = func(steamApps uintptr) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetCurrentGameLanguage, uintptr(steamApps))
		__r0 := goString(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetAvailableGameLanguages, steamAPILib, flatAPI_ISteamApps_GetAvailableGameLanguages)
	iSteamApps_GetAvailableGameLanguages = func(steamApps uintptr) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetAvailableGameLanguages, uintptr(steamApps))
		__r0 := goString(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedApp, steamAPILib, flatAPI_ISteamApps_BIsSubscribedApp)
	iSteamApps_BIsSubscribedApp = func(steamApps uintptr, appID AppId_t) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsSubscribedApp, uintptr(steamApps), uintptr(appID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamApps_BIsDlcInstalled, steamAPILib, flatAPI_ISteamApps_BIsDlcInstalled)
	iSteamApps_BIsDlcInstalled = func(steamApps uintptr, appID AppId_t) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsDlcInstalled, uintptr(steamApps), uintptr(appID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetEarliestPurchaseUnixTime, steamAPILib, flatAPI_ISteamApps_GetEarliestPurchaseUnixTime)
	iSteamApps_GetEarliestPurchaseUnixTime = func(steamApps uintptr, appID AppId_t) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetEarliestPurchaseUnixTime, uintptr(steamApps), uintptr(appID))
		__r0 := uint32(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedFromFreeWeekend, steamAPILib, flatAPI_ISteamApps_BIsSubscribedFromFreeWeekend)
	iSteamApps_BIsSubscribedFromFreeWeekend = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsSubscribedFromFreeWeekend, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetDLCCount, steamAPILib, flatAPI_ISteamApps_GetDLCCount)
	iSteamApps_GetDLCCount = func(steamApps uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetDLCCount, uintptr(steamApps))
		__r0 := int32(_r0)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamApps_BGetDLCDataByIndex, steamAPILib, flatAPI_ISteamApps_BGetDLCDataByIndex)
	iSteamApps_BGetDLCDataByIndex = func(steamApps uintptr, dlcIndex int32, appID *AppId_t, available *bool, name []byte, nameBufferSize int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BGetDLCDataByIndex, uintptr(steamApps), uintptr(dlcIndex), uintptr(unsafe.Pointer(appID)), uintptr(unsafe.Pointer(available)), uintptr(unsafe.Pointer(unsafe.SliceData(name))), uintptr(nameBufferSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(appID)
		runtime.KeepAlive(available)
		runtime.KeepAlive(name)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamApps_InstallDLC, steamAPILib, flatAPI_ISteamApps_InstallDLC)
	iSteamApps_InstallDLC = func(steamApps uintptr, appID AppId_t) {
		purego.SyscallN(addr_iSteamApps_InstallDLC, uintptr(steamApps), uintptr(appID))
	}

	// 	purego.RegisterLibFunc(&iSteamApps_UninstallDLC, steamAPILib, flatAPI_ISteamApps_UninstallDLC)
	iSteamApps_UninstallDLC = func(steamApps uintptr, appID AppId_t) {
		purego.SyscallN(addr_iSteamApps_UninstallDLC, uintptr(steamApps), uintptr(appID))
	}

	// 	purego.RegisterLibFunc(&iSteamApps_RequestAppProofOfPurchaseKey, steamAPILib, flatAPI_ISteamApps_RequestAppProofOfPurchaseKey)
	iSteamApps_RequestAppProofOfPurchaseKey = func(steamApps uintptr, appID AppId_t) {
		purego.SyscallN(addr_iSteamApps_RequestAppProofOfPurchaseKey, uintptr(steamApps), uintptr(appID))
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetCurrentBetaName, steamAPILib, flatAPI_ISteamApps_GetCurrentBetaName)
	iSteamApps_GetCurrentBetaName = func(steamApps uintptr, name []byte, nameBufferSize int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetCurrentBetaName, uintptr(steamApps), uintptr(unsafe.Pointer(unsafe.SliceData(name))), uintptr(nameBufferSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_MarkContentCorrupt, steamAPILib, flatAPI_ISteamApps_MarkContentCorrupt)
	iSteamApps_MarkContentCorrupt = func(steamApps uintptr, missingFilesOnly bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_MarkContentCorrupt, uintptr(steamApps), boolToUintptr(missingFilesOnly))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetInstalledDepots, steamAPILib, flatAPI_ISteamApps_GetInstalledDepots)
	iSteamApps_GetInstalledDepots = func(steamApps uintptr, appID AppId_t, depots []DepotId, maxDepots uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetInstalledDepots, uintptr(steamApps), uintptr(appID), uintptr(unsafe.Pointer(unsafe.SliceData(depots))), uintptr(maxDepots))
		__r0 := uint32(_r0)
		runtime.KeepAlive(depots)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetAppInstallDir, steamAPILib, flatAPI_ISteamApps_GetAppInstallDir)
	iSteamApps_GetAppInstallDir = func(steamApps uintptr, appID AppId_t, folder []byte, folderBufferSize uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetAppInstallDir, uintptr(steamApps), uintptr(appID), uintptr(unsafe.Pointer(unsafe.SliceData(folder))), uintptr(folderBufferSize))
		__r0 := uint32(_r0)
		runtime.KeepAlive(folder)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsAppInstalled, steamAPILib, flatAPI_ISteamApps_BIsAppInstalled)
	iSteamApps_BIsAppInstalled = func(steamApps uintptr, appID AppId_t) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsAppInstalled, uintptr(steamApps), uintptr(appID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetAppOwner, steamAPILib, flatAPI_ISteamApps_GetAppOwner)
	iSteamApps_GetAppOwner = func(steamApps uintptr) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetAppOwner, uintptr(steamApps))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetLaunchQueryParam, steamAPILib, flatAPI_ISteamApps_GetLaunchQueryParam)
	iSteamApps_GetLaunchQueryParam = func(steamApps uintptr, key string) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetLaunchQueryParam, uintptr(steamApps), uintptr(unsafe.Pointer(bytePtrFromString(key))))
		__r0 := goString(_r0)
		runtime.KeepAlive(key)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetDlcDownloadProgress, steamAPILib, flatAPI_ISteamApps_GetDlcDownloadProgress)
	iSteamApps_GetDlcDownloadProgress = func(steamApps uintptr, appID AppId_t, bytesDownloaded *uint64, bytesTotal *uint64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetDlcDownloadProgress, uintptr(steamApps), uintptr(appID), uintptr(unsafe.Pointer(bytesDownloaded)), uintptr(unsafe.Pointer(bytesTotal)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(bytesDownloaded)
		runtime.KeepAlive(bytesTotal)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetAppBuildId, steamAPILib, flatAPI_ISteamApps_GetAppBuildId)
	iSteamApps_GetAppBuildId = func(steamApps uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetAppBuildId, uintptr(steamApps))
		__r0 := int32(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_RequestAllProofOfPurchaseKeys, steamAPILib, flatAPI_ISteamApps_RequestAllProofOfPurchaseKeys)
	iSteamApps_RequestAllProofOfPurchaseKeys = func(steamApps uintptr) {
		purego.SyscallN(addr_iSteamApps_RequestAllProofOfPurchaseKeys, uintptr(steamApps))
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetFileDetails, steamAPILib, flatAPI_ISteamApps_GetFileDetails)
	iSteamApps_GetFileDetails = func(steamApps uintptr, pszFileName string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetFileDetails, uintptr(steamApps), uintptr(unsafe.Pointer(bytePtrFromString(pszFileName))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pszFileName)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetLaunchCommandLine, steamAPILib, flatAPI_ISteamApps_GetLaunchCommandLine)
	iSteamApps_GetLaunchCommandLine = func(steamApps uintptr, pszCommandLine []byte, commandLineStringSize int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetLaunchCommandLine, uintptr(steamApps), uintptr(unsafe.Pointer(unsafe.SliceData(pszCommandLine))), uintptr(commandLineStringSize))
		__r0 := int32(_r0)
		runtime.KeepAlive(pszCommandLine)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsSubscribedFromFamilySharing, steamAPILib, flatAPI_ISteamApps_BIsSubscribedFromFamilySharing)
	iSteamApps_BIsSubscribedFromFamilySharing = func(steamApps uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsSubscribedFromFamilySharing, uintptr(steamApps))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_BIsTimedTrial, steamAPILib, flatAPI_ISteamApps_BIsTimedTrial)
	iSteamApps_BIsTimedTrial = func(steamApps uintptr, secondsAllowed *uint32, secondsPlayed *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_BIsTimedTrial, uintptr(steamApps), uintptr(unsafe.Pointer(secondsAllowed)), uintptr(unsafe.Pointer(secondsPlayed)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(secondsAllowed)
		runtime.KeepAlive(secondsPlayed)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_SetDlcContext, steamAPILib, flatAPI_ISteamApps_SetDlcContext)
	iSteamApps_SetDlcContext = func(steamApps uintptr, appID AppId_t) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_SetDlcContext, uintptr(steamApps), uintptr(appID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetNumBetas, steamAPILib, flatAPI_ISteamApps_GetNumBetas)
	iSteamApps_GetNumBetas = func(steamApps uintptr, available *int32, private *int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetNumBetas, uintptr(steamApps), uintptr(unsafe.Pointer(available)), uintptr(unsafe.Pointer(private)))
		__r0 := int32(_r0)
		runtime.KeepAlive(available)
		runtime.KeepAlive(private)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_GetBetaInfo, steamAPILib, flatAPI_ISteamApps_GetBetaInfo)
	iSteamApps_GetBetaInfo = func(steamApps uintptr, betaIndex int32, punFlags *uint32, punBuildID *uint32, pchBetaName []byte, cchBetaName int32, pchDescription []byte, cchDescription int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_GetBetaInfo, uintptr(steamApps), uintptr(betaIndex), uintptr(unsafe.Pointer(punFlags)), uintptr(unsafe.Pointer(punBuildID)), uintptr(unsafe.Pointer(unsafe.SliceData(pchBetaName))), uintptr(cchBetaName), uintptr(unsafe.Pointer(unsafe.SliceData(pchDescription))), uintptr(cchDescription))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(punFlags)
		runtime.KeepAlive(punBuildID)
		runtime.KeepAlive(pchBetaName)
		runtime.KeepAlive(pchDescription)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamApps_SetActiveBeta, steamAPILib, flatAPI_ISteamApps_SetActiveBeta)
	iSteamApps_SetActiveBeta = func(steamApps uintptr, pchBetaName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamApps_SetActiveBeta, uintptr(steamApps), uintptr(unsafe.Pointer(bytePtrFromString(pchBetaName))))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}

var (
	addr_iSteamClient_CreateSteamPipe             uintptr
	addr_iSteamClient_BReleaseSteamPipe           uintptr
	addr_iSteamClient_ConnectToGlobalUser         uintptr
	addr_iSteamClient_CreateLocalUser             uintptr
	addr_iSteamClient_ReleaseUser                 uintptr
	addr_iSteamClient_GetISteamUser               uintptr
	addr_iSteamClient_GetISteamGameServer         uintptr
	addr_iSteamClient_SetLocalIPBinding           uintptr
	addr_iSteamClient_GetISteamUtils              uintptr
	addr_iSteamClient_GetISteamMatchmaking        uintptr
	addr_iSteamClient_GetISteamMatchmakingServers uintptr
	addr_iSteamClient_GetISteamGenericInterface   uintptr
	addr_iSteamClient_GetISteamUserStats          uintptr
	addr_iSteamClient_GetISteamGameServerStats    uintptr
	addr_iSteamClient_GetISteamApps               uintptr
	addr_iSteamClient_GetISteamNetworking         uintptr
	addr_iSteamClient_GetISteamRemoteStorage      uintptr
	addr_iSteamClient_GetISteamScreenshots        uintptr
	addr_iSteamClient_GetISteamGameSearch         uintptr
	addr_iSteamClient_GetIPCCallCount             uintptr
	addr_iSteamClient_SetWarningMessageHook       uintptr
	addr_iSteamClient_BShutdownIfAllPipesClosed   uintptr
	addr_iSteamClient_GetISteamUGC                uintptr
	addr_iSteamClient_GetISteamInventory          uintptr
	addr_iSteamClient_GetISteamInput              uintptr
	addr_iSteamClient_GetISteamParties            uintptr
	addr_iSteamClient_GetISteamRemotePlay         uintptr
)

func initClient() {
	var err error
	addr_iSteamClient_CreateSteamPipe, err = openSymbol(steamAPILib, flatAPI_ISteamClient_CreateSteamPipe)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_CreateSteamPipe")
	}
	addr_iSteamClient_BReleaseSteamPipe, err = openSymbol(steamAPILib, flatAPI_ISteamClient_BReleaseSteamPipe)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_BReleaseSteamPipe")
	}
	addr_iSteamClient_ConnectToGlobalUser, err = openSymbol(steamAPILib, flatAPI_ISteamClient_ConnectToGlobalUser)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_ConnectToGlobalUser")
	}
	addr_iSteamClient_CreateLocalUser, err = openSymbol(steamAPILib, flatAPI_ISteamClient_CreateLocalUser)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_CreateLocalUser")
	}
	addr_iSteamClient_ReleaseUser, err = openSymbol(steamAPILib, flatAPI_ISteamClient_ReleaseUser)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_ReleaseUser")
	}
	addr_iSteamClient_GetISteamUser, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamUser)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamUser")
	}
	addr_iSteamClient_GetISteamGameServer, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamGameServer)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamGameServer")
	}
	addr_iSteamClient_SetLocalIPBinding, err = openSymbol(steamAPILib, flatAPI_ISteamClient_SetLocalIPBinding)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_SetLocalIPBinding")
	}
	addr_iSteamClient_GetISteamUtils, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamUtils)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamUtils")
	}
	addr_iSteamClient_GetISteamMatchmaking, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamMatchmaking)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamMatchmaking")
	}
	addr_iSteamClient_GetISteamMatchmakingServers, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamMatchmakingServers)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamMatchmakingServers")
	}
	addr_iSteamClient_GetISteamGenericInterface, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamGenericInterface)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamGenericInterface")
	}
	addr_iSteamClient_GetISteamUserStats, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamUserStats)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamUserStats")
	}
	addr_iSteamClient_GetISteamGameServerStats, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamGameServerStats)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamGameServerStats")
	}
	addr_iSteamClient_GetISteamApps, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamApps)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamApps")
	}
	addr_iSteamClient_GetISteamNetworking, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamNetworking)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamNetworking")
	}
	addr_iSteamClient_GetISteamRemoteStorage, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamRemoteStorage)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamRemoteStorage")
	}
	addr_iSteamClient_GetISteamScreenshots, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamScreenshots)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamScreenshots")
	}
	addr_iSteamClient_GetISteamGameSearch, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamGameSearch)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamGameSearch")
	}
	addr_iSteamClient_GetIPCCallCount, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetIPCCallCount)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetIPCCallCount")
	}
	addr_iSteamClient_SetWarningMessageHook, err = openSymbol(steamAPILib, flatAPI_ISteamClient_SetWarningMessageHook)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_SetWarningMessageHook")
	}
	addr_iSteamClient_BShutdownIfAllPipesClosed, err = openSymbol(steamAPILib, flatAPI_ISteamClient_BShutdownIfAllPipesClosed)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_BShutdownIfAllPipesClosed")
	}

	addr_iSteamClient_GetISteamUGC, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamUGC)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamUGC")
	}
	addr_iSteamClient_GetISteamInventory, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamInventory)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamInventory")
	}
	addr_iSteamClient_GetISteamInput, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamInput)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamInput")
	}
	addr_iSteamClient_GetISteamParties, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamParties)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamParties")
	}
	addr_iSteamClient_GetISteamRemotePlay, err = openSymbol(steamAPILib, flatAPI_ISteamClient_GetISteamRemotePlay)
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamClient_GetISteamRemotePlay")
	}

	// 	purego.RegisterLibFunc(&iSteamClient_CreateSteamPipe, steamAPILib, flatAPI_ISteamClient_CreateSteamPipe)
	iSteamClient_CreateSteamPipe = func() HSteamPipe {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_CreateSteamPipe)
		__r0 := HSteamPipe(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_BReleaseSteamPipe, steamAPILib, flatAPI_ISteamClient_BReleaseSteamPipe)
	iSteamClient_BReleaseSteamPipe = func(hSteamPipe HSteamPipe) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_BReleaseSteamPipe, uintptr(hSteamPipe))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamClient_ConnectToGlobalUser, steamAPILib, flatAPI_ISteamClient_ConnectToGlobalUser)
	iSteamClient_ConnectToGlobalUser = func(hSteamPipe HSteamPipe) HSteamUser {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_ConnectToGlobalUser, uintptr(hSteamPipe))
		__r0 := HSteamUser(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_CreateLocalUser, steamAPILib, flatAPI_ISteamClient_CreateLocalUser)
	iSteamClient_CreateLocalUser = func(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_CreateLocalUser, uintptr(unsafe.Pointer(phSteamPipe)), uintptr(eAccountType))
		__r0 := HSteamUser(_r0)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamClient_ReleaseUser, steamAPILib, flatAPI_ISteamClient_ReleaseUser)
	iSteamClient_ReleaseUser = func(hSteamPipe HSteamPipe, hUser HSteamUser) {
		purego.SyscallN(addr_iSteamClient_ReleaseUser, uintptr(hSteamPipe), uintptr(hUser))
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamUser, steamAPILib, flatAPI_ISteamClient_GetISteamUser)
	iSteamClient_GetISteamUser = func(hUser HSteamUser, hSteamPipe HSteamPipe, pcVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamUser, uintptr(hSteamPipe), uintptr(hUser))
		__r0 := uintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamGameServer, steamAPILib, flatAPI_ISteamClient_GetISteamGameServer)
	iSteamClient_GetISteamGameServer = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamGameServer, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_SetLocalIPBinding, steamAPILib, flatAPI_ISteamClient_SetLocalIPBinding)
	iSteamClient_SetLocalIPBinding = func(unIP *SteamIPAddress, usPort uint16) {
		purego.SyscallN(addr_iSteamClient_SetLocalIPBinding, uintptr(unsafe.Pointer(unIP)), uintptr(usPort))
		runtime.KeepAlive(unIP)
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamFriends, steamAPILib, flatAPI_ISteamClient_GetISteamFriends)
	iSteamClient_GetISteamFriends = func(hUser HSteamUser, hSteamPipe HSteamPipe, pcVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamUser, uintptr(hSteamPipe), uintptr(hUser))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pcVersion)
		return __r0
	}

	//purego.RegisterLibFunc(&iSteamClient_GetISteamUtils, steamAPILib, flatAPI_ISteamClient_GetISteamUtils)
	iSteamClient_GetISteamUtils = func(hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamUtils, uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamMatchmaking, steamAPILib, flatAPI_ISteamClient_GetISteamMatchmaking)
	iSteamClient_GetISteamMatchmaking = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamMatchmaking, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamMatchmakingServers, steamAPILib, flatAPI_ISteamClient_GetISteamMatchmakingServers)
	iSteamClient_GetISteamMatchmakingServers = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamMatchmakingServers, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamGenericInterface, steamAPILib, flatAPI_ISteamClient_GetISteamGenericInterface)
	iSteamClient_GetISteamGenericInterface = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamGenericInterface, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := unsafe.Pointer(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamUserStats, steamAPILib, flatAPI_ISteamClient_GetISteamUserStats)
	iSteamClient_GetISteamUserStats = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamUserStats, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}
	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamGameServerStats, steamAPILib, flatAPI_ISteamClient_GetISteamGameServerStats)
	iSteamClient_GetISteamGameServerStats = func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamGameServerStats, uintptr(hSteamuser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamClient_GetISteamApps, steamAPILib, flatAPI_ISteamClient_GetISteamApps)
	iSteamClient_GetISteamApps = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamApps, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}
	// purego.RegisterLibFunc(&iSteamClient_GetISteamNetworking, steamAPILib, flatAPI_ISteamClient_GetISteamNetworking)
	iSteamClient_GetISteamNetworking = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamNetworking, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamClient_GetISteamRemoteStorage, steamAPILib, flatAPI_ISteamClient_GetISteamRemoteStorage)
	iSteamClient_GetISteamRemoteStorage = func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamRemoteStorage, uintptr(hSteamuser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamScreenshots, steamAPILib, flatAPI_ISteamClient_GetISteamScreenshots)
	iSteamClient_GetISteamScreenshots = func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamScreenshots, uintptr(hSteamuser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamGameSearch, steamAPILib, flatAPI_ISteamClient_GetISteamGameSearch)
	iSteamClient_GetISteamGameSearch = func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamGameSearch, uintptr(hSteamuser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetIPCCallCount, steamAPILib, flatAPI_ISteamClient_GetIPCCallCount)
	iSteamClient_GetIPCCallCount = func() uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetIPCCallCount)
		__r0 := uint32(_r0)
		return __r0

	}

	// 	purego.RegisterLibFunc(&iSteamClient_SetWarningMessageHook, steamAPILib, flatAPI_ISteamClient_SetWarningMessageHook)
	iSteamClient_SetWarningMessageHook = func(pFunction SteamAPIWarningMessageHook) {
		wrapper := func(severity int32, msg uintptr) {
			convertedString := goString(msg)
			pFunction(severity, convertedString)
		}
		cb := purego.NewCallback(wrapper)
		purego.SyscallN(addr_iSteamClient_SetWarningMessageHook, uintptr(cb))
	}

	//  purego.RegisterLibFunc(&iSteamClient_BShutdownIfAllPipesClosed, steamAPILib, flatAPI_ISteamClient_BShutdownIfAllPipesClosed)
	iSteamClient_BShutdownIfAllPipesClosed = func() bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_BShutdownIfAllPipesClosed)
		__r0 := boolFromUintptr(_r0)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamUGC, steamAPILib, flatAPI_ISteamClient_GetISteamUGC)
	iSteamClient_GetISteamUGC = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamUGC, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamInventory, steamAPILib, flatAPI_ISteamClient_GetISteamInventory)
	iSteamClient_GetISteamInventory = func(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamInventory, uintptr(hSteamuser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamInput, steamAPILib, flatAPI_ISteamClient_GetISteamInput)
	iSteamClient_GetISteamInput = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamInput, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// 	purego.RegisterLibFunc(&iSteamClient_GetISteamParties, steamAPILib, flatAPI_ISteamClient_GetISteamParties)
	iSteamClient_GetISteamParties = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamParties, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

	// purego.RegisterLibFunc(&iSteamClient_GetISteamRemotePlay, steamAPILib, flatAPI_ISteamClient_GetISteamRemotePlay)
	iSteamClient_GetISteamRemotePlay = func(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamClient_GetISteamRemotePlay, uintptr(hSteamUser), uintptr(hSteamPipe), uintptr(unsafe.Pointer(bytePtrFromString(pchVersion))))
		__r0 := uintptr(_r0)
		runtime.KeepAlive(pchVersion)
		return __r0
	}

}

var (
	addr_steamFriends_get                                                uintptr
	addr_iSteamFriends_GetPersonaName                                    uintptr
	addr_iSteamFriends_SetPersonaName                                    uintptr
	addr_iSteamFriends_GetPersonaState                                   uintptr
	addr_iSteamFriends_GetFriendCount                                    uintptr
	addr_iSteamFriends_GetFriendByIndex                                  uintptr
	addr_iSteamFriends_GetFriendRelationship                             uintptr
	addr_iSteamFriends_GetFriendPersonaState                             uintptr
	addr_iSteamFriends_GetFriendPersonaName                              uintptr
	addr_iSteamFriends_GetFriendGamePlayed                               uintptr
	addr_iSteamFriends_GetFriendPersonaNameHistory                       uintptr
	addr_iSteamFriends_GetFriendSteamLevel                               uintptr
	addr_iSteamFriends_GetFriendsGroupCount                              uintptr
	addr_iSteamFriends_GetFriendsGroupIDByIndex                          uintptr
	addr_iSteamFriends_GetFriendsGroupName                               uintptr
	addr_iSteamFriends_GetFriendsGroupMembersCount                       uintptr
	addr_iSteamFriends_GetFriendsGroupMembersList                        uintptr
	addr_iSteamFriends_HasFriend                                         uintptr
	addr_iSteamFriends_GetClanCount                                      uintptr
	addr_iSteamFriends_GetClanByIndex                                    uintptr
	addr_iSteamFriends_GetClanName                                       uintptr
	addr_iSteamFriends_GetClanTag                                        uintptr
	addr_iSteamFriends_GetClanActivityCounts                             uintptr
	addr_iSteamFriends_DownloadClanActivityCounts                        uintptr
	addr_iSteamFriends_GetFriendCountFromSource                          uintptr
	addr_iSteamFriends_GetFriendFromSourceByIndex                        uintptr
	addr_iSteamFriends_IsUserInSource                                    uintptr
	addr_iSteamFriends_SetInGameVoiceSpeaking                            uintptr
	addr_iSteamFriends_ActivateGameOverlay                               uintptr
	addr_iSteamFriends_ActivateGameOverlayToUser                         uintptr
	addr_iSteamFriends_ActivateGameOverlayToWebPage                      uintptr
	addr_iSteamFriends_ActivateGameOverlayToStore                        uintptr
	addr_iSteamFriends_SetPlayedWith                                     uintptr
	addr_iSteamFriends_ActivateGameOverlayInviteDialog                   uintptr
	addr_iSteamFriends_GetSmallFriendAvatar                              uintptr
	addr_iSteamFriends_GetMediumFriendAvatar                             uintptr
	addr_iSteamFriends_GetLargeFriendAvatar                              uintptr
	addr_iSteamFriends_RequestUserInformation                            uintptr
	addr_iSteamFriends_RequestClanOfficerList                            uintptr
	addr_iSteamFriends_GetClanOwner                                      uintptr
	addr_iSteamFriends_GetClanOfficerCount                               uintptr
	addr_iSteamFriends_GetClanOfficerByIndex                             uintptr
	addr_iSteamFriends_GetUserRestrictions                               uintptr
	addr_iSteamFriends_SetRichPresence                                   uintptr
	addr_iSteamFriends_ClearRichPresence                                 uintptr
	addr_iSteamFriends_GetFriendRichPresence                             uintptr
	addr_iSteamFriends_GetFriendRichPresenceKeyCount                     uintptr
	addr_iSteamFriends_GetFriendRichPresenceKeyByIndex                   uintptr
	addr_iSteamFriends_RequestFriendRichPresence                         uintptr
	addr_iSteamFriends_InviteUserToGame                                  uintptr
	addr_iSteamFriends_GetCoplayFriendCount                              uintptr
	addr_iSteamFriends_GetCoplayFriend                                   uintptr
	addr_iSteamFriends_GetFriendCoplayTime                               uintptr
	addr_iSteamFriends_GetFriendCoplayGame                               uintptr
	addr_iSteamFriends_JoinClanChatRoom                                  uintptr
	addr_iSteamFriends_LeaveClanChatRoom                                 uintptr
	addr_iSteamFriends_GetClanChatMemberCount                            uintptr
	addr_iSteamFriends_GetChatMemberByIndex                              uintptr
	addr_iSteamFriends_SendClanChatMessage                               uintptr
	addr_iSteamFriends_GetClanChatMessage                                uintptr
	addr_iSteamFriends_IsClanChatAdmin                                   uintptr
	addr_iSteamFriends_IsClanChatWindowOpenInSteam                       uintptr
	addr_iSteamFriends_OpenClanChatWindowInSteam                         uintptr
	addr_iSteamFriends_CloseClanChatWindowInSteam                        uintptr
	addr_iSteamFriends_SetListenForFriendsMessages                       uintptr
	addr_iSteamFriends_ReplyToFriendMessage                              uintptr
	addr_iSteamFriends_GetFriendMessage                                  uintptr
	addr_iSteamFriends_GetFollowerCount                                  uintptr
	addr_iSteamFriends_IsFollowing                                       uintptr
	addr_iSteamFriends_EnumerateFollowingList                            uintptr
	addr_iSteamFriends_IsClanPublic                                      uintptr
	addr_iSteamFriends_IsClanOfficialGameGroup                           uintptr
	addr_iSteamFriends_GetNumChatsWithUnreadPriorityMessages             uintptr
	addr_iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog uintptr
	addr_iSteamFriends_RegisterProtocolInOverlayBrowser                  uintptr
	addr_iSteamFriends_ActivateGameOverlayInviteDialogConnectString      uintptr
	addr_iSteamFriends_RequestEquippedProfileItems                       uintptr
	addr_iSteamFriends_BHasEquippedProfileItem                           uintptr
	addr_iSteamFriends_GetProfileItemPropertyString                      uintptr
	addr_iSteamFriends_GetProfileItemPropertyUint                        uintptr
)

func initFriends() {
	var err error
	addr_steamFriends_get, err = openSymbol(steamAPILib, flatAPI_SteamFriends)
	if err != nil {
		panic("cannot OpenSymbol:" + flatAPI_SteamFriends)
	}

	addr_iSteamFriends_GetPersonaName, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetPersonaName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetPersonaName")
	}
	addr_iSteamFriends_SetPersonaName, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SetPersonaName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SetPersonaName")
	}
	addr_iSteamFriends_GetPersonaState, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetPersonaState")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetPersonaState")
	}
	addr_iSteamFriends_GetFriendCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendCount")
	}
	addr_iSteamFriends_GetFriendByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendByIndex")
	}
	addr_iSteamFriends_GetFriendRelationship, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendRelationship")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendRelationship")
	}
	addr_iSteamFriends_GetFriendPersonaState, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendPersonaState")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendPersonaState")
	}
	addr_iSteamFriends_GetFriendPersonaName, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendPersonaName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendPersonaName")
	}
	addr_iSteamFriends_GetFriendGamePlayed, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendGamePlayed")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendGamePlayed")
	}
	addr_iSteamFriends_GetFriendPersonaNameHistory, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendPersonaNameHistory")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendPersonaNameHistory")
	}
	addr_iSteamFriends_GetFriendSteamLevel, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendSteamLevel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendSteamLevel")
	}
	addr_iSteamFriends_GetFriendsGroupCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendsGroupCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendsGroupCount")
	}
	addr_iSteamFriends_GetFriendsGroupIDByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendsGroupIDByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendsGroupIDByIndex")
	}
	addr_iSteamFriends_GetFriendsGroupName, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendsGroupName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendsGroupName")
	}
	addr_iSteamFriends_GetFriendsGroupMembersCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendsGroupMembersCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendsGroupMembersCount")
	}
	addr_iSteamFriends_GetFriendsGroupMembersList, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendsGroupMembersList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendsGroupMembersList")
	}
	addr_iSteamFriends_HasFriend, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_HasFriend")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_HasFriend")
	}
	addr_iSteamFriends_GetClanCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanCount")
	}
	addr_iSteamFriends_GetClanByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanByIndex")
	}
	addr_iSteamFriends_GetClanName, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanName")
	}
	addr_iSteamFriends_GetClanTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanTag")
	}
	addr_iSteamFriends_GetClanActivityCounts, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanActivityCounts")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanActivityCounts")
	}
	addr_iSteamFriends_DownloadClanActivityCounts, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_DownloadClanActivityCounts")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_DownloadClanActivityCounts")
	}
	addr_iSteamFriends_GetFriendCountFromSource, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendCountFromSource")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendCountFromSource")
	}
	addr_iSteamFriends_GetFriendFromSourceByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendFromSourceByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendFromSourceByIndex")
	}
	addr_iSteamFriends_IsUserInSource, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsUserInSource")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsUserInSource")
	}
	addr_iSteamFriends_SetInGameVoiceSpeaking, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SetInGameVoiceSpeaking")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SetInGameVoiceSpeaking")
	}
	addr_iSteamFriends_ActivateGameOverlay, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlay")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlay")
	}
	addr_iSteamFriends_ActivateGameOverlayToUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayToUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayToUser")
	}
	addr_iSteamFriends_ActivateGameOverlayToWebPage, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayToWebPage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayToWebPage")
	}
	addr_iSteamFriends_ActivateGameOverlayToStore, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayToStore")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayToStore")
	}
	addr_iSteamFriends_SetPlayedWith, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SetPlayedWith")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SetPlayedWith")
	}
	addr_iSteamFriends_ActivateGameOverlayInviteDialog, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialog")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialog")
	}
	addr_iSteamFriends_GetSmallFriendAvatar, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetSmallFriendAvatar")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetSmallFriendAvatar")
	}
	addr_iSteamFriends_GetMediumFriendAvatar, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetMediumFriendAvatar")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetMediumFriendAvatar")
	}
	addr_iSteamFriends_GetLargeFriendAvatar, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetLargeFriendAvatar")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetLargeFriendAvatar")
	}
	addr_iSteamFriends_RequestUserInformation, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_RequestUserInformation")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_RequestUserInformation")
	}
	addr_iSteamFriends_RequestClanOfficerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_RequestClanOfficerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_RequestClanOfficerList")
	}
	addr_iSteamFriends_GetClanOwner, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanOwner")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanOwner")
	}
	addr_iSteamFriends_GetClanOfficerCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanOfficerCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanOfficerCount")
	}
	addr_iSteamFriends_GetClanOfficerByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanOfficerByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanOfficerByIndex")
	}
	addr_iSteamFriends_GetUserRestrictions, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetUserRestrictions")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetUserRestrictions")
	}
	addr_iSteamFriends_SetRichPresence, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SetRichPresence")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SetRichPresence")
	}
	addr_iSteamFriends_ClearRichPresence, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ClearRichPresence")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ClearRichPresence")
	}
	addr_iSteamFriends_GetFriendRichPresence, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendRichPresence")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendRichPresence")
	}
	addr_iSteamFriends_GetFriendRichPresenceKeyCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendRichPresenceKeyCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendRichPresenceKeyCount")
	}
	addr_iSteamFriends_GetFriendRichPresenceKeyByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendRichPresenceKeyByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendRichPresenceKeyByIndex")
	}
	addr_iSteamFriends_RequestFriendRichPresence, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_RequestFriendRichPresence")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_RequestFriendRichPresence")
	}
	addr_iSteamFriends_InviteUserToGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_InviteUserToGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_InviteUserToGame")
	}
	addr_iSteamFriends_GetCoplayFriendCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetCoplayFriendCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetCoplayFriendCount")
	}
	addr_iSteamFriends_GetCoplayFriend, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetCoplayFriend")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetCoplayFriend")
	}
	addr_iSteamFriends_GetFriendCoplayTime, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendCoplayTime")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendCoplayTime")
	}
	addr_iSteamFriends_GetFriendCoplayGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendCoplayGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendCoplayGame")
	}
	addr_iSteamFriends_JoinClanChatRoom, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_JoinClanChatRoom")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_JoinClanChatRoom")
	}
	addr_iSteamFriends_LeaveClanChatRoom, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_LeaveClanChatRoom")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_LeaveClanChatRoom")
	}
	addr_iSteamFriends_GetClanChatMemberCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanChatMemberCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanChatMemberCount")
	}
	addr_iSteamFriends_GetChatMemberByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetChatMemberByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetChatMemberByIndex")
	}
	addr_iSteamFriends_SendClanChatMessage, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SendClanChatMessage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SendClanChatMessage")
	}
	addr_iSteamFriends_GetClanChatMessage, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetClanChatMessage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetClanChatMessage")
	}
	addr_iSteamFriends_IsClanChatAdmin, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsClanChatAdmin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsClanChatAdmin")
	}
	addr_iSteamFriends_IsClanChatWindowOpenInSteam, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsClanChatWindowOpenInSteam")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsClanChatWindowOpenInSteam")
	}
	addr_iSteamFriends_OpenClanChatWindowInSteam, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_OpenClanChatWindowInSteam")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_OpenClanChatWindowInSteam")
	}
	addr_iSteamFriends_CloseClanChatWindowInSteam, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_CloseClanChatWindowInSteam")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_CloseClanChatWindowInSteam")
	}
	addr_iSteamFriends_SetListenForFriendsMessages, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_SetListenForFriendsMessages")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_SetListenForFriendsMessages")
	}
	addr_iSteamFriends_ReplyToFriendMessage, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ReplyToFriendMessage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ReplyToFriendMessage")
	}
	addr_iSteamFriends_GetFriendMessage, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFriendMessage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFriendMessage")
	}
	addr_iSteamFriends_GetFollowerCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetFollowerCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetFollowerCount")
	}
	addr_iSteamFriends_IsFollowing, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsFollowing")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsFollowing")
	}
	addr_iSteamFriends_EnumerateFollowingList, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_EnumerateFollowingList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_EnumerateFollowingList")
	}
	addr_iSteamFriends_IsClanPublic, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsClanPublic")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsClanPublic")
	}
	addr_iSteamFriends_IsClanOfficialGameGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_IsClanOfficialGameGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_IsClanOfficialGameGroup")
	}
	addr_iSteamFriends_GetNumChatsWithUnreadPriorityMessages, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetNumChatsWithUnreadPriorityMessages")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetNumChatsWithUnreadPriorityMessages")
	}
	addr_iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog")
	}
	addr_iSteamFriends_RegisterProtocolInOverlayBrowser, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_RegisterProtocolInOverlayBrowser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_RegisterProtocolInOverlayBrowser")
	}
	addr_iSteamFriends_ActivateGameOverlayInviteDialogConnectString, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialogConnectString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_ActivateGameOverlayInviteDialogConnectString")
	}
	addr_iSteamFriends_RequestEquippedProfileItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_RequestEquippedProfileItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_RequestEquippedProfileItems")
	}
	addr_iSteamFriends_BHasEquippedProfileItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_BHasEquippedProfileItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_BHasEquippedProfileItem")
	}
	addr_iSteamFriends_GetProfileItemPropertyString, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetProfileItemPropertyString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetProfileItemPropertyString")
	}
	addr_iSteamFriends_GetProfileItemPropertyUint, err = openSymbol(steamAPILib, "SteamAPI_ISteamFriends_GetProfileItemPropertyUint")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamFriends_GetProfileItemPropertyUint")
	}

	steamFriends_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamFriends_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamFriends_GetPersonaName = func(steamFriends uintptr) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetPersonaName, uintptr(steamFriends))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_SetPersonaName = func(steamFriends uintptr, pchPersonaName string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_SetPersonaName, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchPersonaName))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchPersonaName)
		return __r0
	}
	iSteamFriends_GetPersonaState = func(steamFriends uintptr) EPersonaState {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetPersonaState, uintptr(steamFriends))
		__r0 := EPersonaState(_r0)
		return __r0
	}
	iSteamFriends_GetFriendCount = func(steamFriends uintptr, iFriendFlags int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendCount, uintptr(steamFriends), uintptr(iFriendFlags))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendByIndex = func(steamFriends uintptr, iFriend int32, iFriendFlags int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendByIndex, uintptr(steamFriends), uintptr(iFriend), uintptr(iFriendFlags))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_GetFriendRelationship = func(steamFriends uintptr, friendSteamID Uint64SteamID) EFriendRelationship {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendRelationship, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := EFriendRelationship(_r0)
		return __r0
	}
	iSteamFriends_GetFriendPersonaState = func(steamFriends uintptr, friendSteamID Uint64SteamID) EPersonaState {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendPersonaState, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := EPersonaState(_r0)
		return __r0
	}
	iSteamFriends_GetFriendPersonaName = func(steamFriends uintptr, friendSteamID Uint64SteamID) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendPersonaName, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetFriendGamePlayed = func(steamFriends uintptr, friendSteamID Uint64SteamID, pFriendGameInfo *FriendGameInfo) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendGamePlayed, uintptr(steamFriends), uintptr(friendSteamID), uintptr(unsafe.Pointer(pFriendGameInfo)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pFriendGameInfo)
		return __r0
	}
	iSteamFriends_GetFriendPersonaNameHistory = func(steamFriends uintptr, friendSteamID Uint64SteamID, iPersonaName int32) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendPersonaNameHistory, uintptr(steamFriends), uintptr(friendSteamID), uintptr(iPersonaName))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetFriendSteamLevel = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendSteamLevel, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendsGroupCount = func(steamFriends uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendsGroupCount, uintptr(steamFriends))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendsGroupIDByIndex = func(steamFriends uintptr, iFG int32) FriendsGroupID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendsGroupIDByIndex, uintptr(steamFriends), uintptr(iFG))
		__r0 := FriendsGroupID(_r0)
		return __r0
	}
	iSteamFriends_GetFriendsGroupName = func(steamFriends uintptr, friendsGroupID FriendsGroupID) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendsGroupName, uintptr(steamFriends), uintptr(friendsGroupID))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetFriendsGroupMembersCount = func(steamFriends uintptr, friendsGroupID FriendsGroupID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendsGroupMembersCount, uintptr(steamFriends), uintptr(friendsGroupID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendsGroupMembersList = func(steamFriends uintptr, friendsGroupID FriendsGroupID, pOutSteamIDMembers []CSteamID, nMembersCount int32) {
		purego.SyscallN(addr_iSteamFriends_GetFriendsGroupMembersList, uintptr(steamFriends), uintptr(friendsGroupID), uintptr(unsafe.Pointer(unsafe.SliceData(pOutSteamIDMembers))), uintptr(nMembersCount))
		runtime.KeepAlive(pOutSteamIDMembers)
	}
	iSteamFriends_HasFriend = func(steamFriends uintptr, friendSteamID Uint64SteamID, iFriendFlags int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_HasFriend, uintptr(steamFriends), uintptr(friendSteamID), uintptr(iFriendFlags))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_GetClanCount = func(steamFriends uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanCount, uintptr(steamFriends))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetClanByIndex = func(steamFriends uintptr, iClan int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanByIndex, uintptr(steamFriends), uintptr(iClan))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_GetClanName = func(steamFriends uintptr, clanSteamID Uint64SteamID) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanName, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetClanTag = func(steamFriends uintptr, clanSteamID Uint64SteamID) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanTag, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetClanActivityCounts = func(steamFriends uintptr, clanSteamID Uint64SteamID, pnOnline *int32, pnInGame *int32, pnChatting *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanActivityCounts, uintptr(steamFriends), uintptr(clanSteamID), uintptr(unsafe.Pointer(pnOnline)), uintptr(unsafe.Pointer(pnInGame)), uintptr(unsafe.Pointer(pnChatting)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnOnline)
		runtime.KeepAlive(pnInGame)
		runtime.KeepAlive(pnChatting)
		return __r0
	}
	iSteamFriends_DownloadClanActivityCounts = func(steamFriends uintptr, pclanSteamIDs []CSteamID, cClansToRequest int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_DownloadClanActivityCounts, uintptr(steamFriends), uintptr(unsafe.Pointer(unsafe.SliceData(pclanSteamIDs))), uintptr(cClansToRequest))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pclanSteamIDs)
		return __r0
	}
	iSteamFriends_GetFriendCountFromSource = func(steamFriends uintptr, sourceSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendCountFromSource, uintptr(steamFriends), uintptr(sourceSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendFromSourceByIndex = func(steamFriends uintptr, sourceSteamID Uint64SteamID, iFriend int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendFromSourceByIndex, uintptr(steamFriends), uintptr(sourceSteamID), uintptr(iFriend))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_IsUserInSource = func(steamFriends uintptr, userSteamID Uint64SteamID, sourceSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsUserInSource, uintptr(steamFriends), uintptr(userSteamID), uintptr(sourceSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_SetInGameVoiceSpeaking = func(steamFriends uintptr, userSteamID Uint64SteamID, bSpeaking bool) {
		purego.SyscallN(addr_iSteamFriends_SetInGameVoiceSpeaking, uintptr(steamFriends), uintptr(userSteamID), boolToUintptr(bSpeaking))
	}
	iSteamFriends_ActivateGameOverlay = func(steamFriends uintptr, pchDialog string) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlay, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchDialog))))
		runtime.KeepAlive(pchDialog)
	}
	iSteamFriends_ActivateGameOverlayToUser = func(steamFriends uintptr, pchDialog string, steamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayToUser, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchDialog))), uintptr(steamID))
		runtime.KeepAlive(pchDialog)
	}
	iSteamFriends_ActivateGameOverlayToWebPage = func(steamFriends uintptr, pchURL string, eMode EActivateGameOverlayToWebPageMode) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayToWebPage, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchURL))), uintptr(eMode))
		runtime.KeepAlive(pchURL)
	}
	iSteamFriends_ActivateGameOverlayToStore = func(steamFriends uintptr, nAppID AppId_t, eFlag EOverlayToStoreFlag) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayToStore, uintptr(steamFriends), uintptr(nAppID), uintptr(eFlag))
	}
	iSteamFriends_SetPlayedWith = func(steamFriends uintptr, userSteamIDPlayedWith Uint64SteamID) {
		purego.SyscallN(addr_iSteamFriends_SetPlayedWith, uintptr(steamFriends), uintptr(userSteamIDPlayedWith))
	}
	iSteamFriends_ActivateGameOverlayInviteDialog = func(steamFriends uintptr, lobbySteamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayInviteDialog, uintptr(steamFriends), uintptr(lobbySteamID))
	}
	iSteamFriends_GetSmallFriendAvatar = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetSmallFriendAvatar, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetMediumFriendAvatar = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetMediumFriendAvatar, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetLargeFriendAvatar = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetLargeFriendAvatar, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_RequestUserInformation = func(steamFriends uintptr, userSteamID Uint64SteamID, bRequireNameOnly bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_RequestUserInformation, uintptr(steamFriends), uintptr(userSteamID), boolToUintptr(bRequireNameOnly))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_RequestClanOfficerList = func(steamFriends uintptr, clanSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_RequestClanOfficerList, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_GetClanOwner = func(steamFriends uintptr, clanSteamID Uint64SteamID) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanOwner, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_GetClanOfficerCount = func(steamFriends uintptr, clanSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanOfficerCount, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetClanOfficerByIndex = func(steamFriends uintptr, clanSteamID Uint64SteamID, iOfficer int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanOfficerByIndex, uintptr(steamFriends), uintptr(clanSteamID), uintptr(iOfficer))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_GetUserRestrictions = func(steamFriends uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetUserRestrictions, uintptr(steamFriends))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamFriends_SetRichPresence = func(steamFriends uintptr, pchKey string, pchValue string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_SetRichPresence, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))), uintptr(unsafe.Pointer(bytePtrFromString(pchValue))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamFriends_ClearRichPresence = func(steamFriends uintptr) {
		purego.SyscallN(addr_iSteamFriends_ClearRichPresence, uintptr(steamFriends))
	}
	iSteamFriends_GetFriendRichPresenceKeyCount = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendRichPresenceKeyCount, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_RequestFriendRichPresence = func(steamFriends uintptr, friendSteamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamFriends_RequestFriendRichPresence, uintptr(steamFriends), uintptr(friendSteamID))
	}
	iSteamFriends_InviteUserToGame = func(steamFriends uintptr, friendSteamID Uint64SteamID, pchConnectString string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_InviteUserToGame, uintptr(steamFriends), uintptr(friendSteamID), uintptr(unsafe.Pointer(bytePtrFromString(pchConnectString))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchConnectString)
		return __r0
	}
	iSteamFriends_GetCoplayFriendCount = func(steamFriends uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetCoplayFriendCount, uintptr(steamFriends))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetCoplayFriend = func(steamFriends uintptr, iCoplayFriend int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetCoplayFriend, uintptr(steamFriends), uintptr(iCoplayFriend))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_GetFriendCoplayTime = func(steamFriends uintptr, friendSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendCoplayTime, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetFriendCoplayGame = func(steamFriends uintptr, friendSteamID Uint64SteamID) AppId_t {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendCoplayGame, uintptr(steamFriends), uintptr(friendSteamID))
		__r0 := AppId_t(_r0)
		return __r0
	}
	iSteamFriends_JoinClanChatRoom = func(steamFriends uintptr, clanSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_JoinClanChatRoom, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_LeaveClanChatRoom = func(steamFriends uintptr, clanSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_LeaveClanChatRoom, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_GetClanChatMemberCount = func(steamFriends uintptr, clanSteamID Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanChatMemberCount, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_GetChatMemberByIndex = func(steamFriends uintptr, clanSteamID Uint64SteamID, iUser int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetChatMemberByIndex, uintptr(steamFriends), uintptr(clanSteamID), uintptr(iUser))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamFriends_SendClanChatMessage = func(steamFriends uintptr, clanChatSteamID Uint64SteamID, pchText string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_SendClanChatMessage, uintptr(steamFriends), uintptr(clanChatSteamID), uintptr(unsafe.Pointer(bytePtrFromString(pchText))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchText)
		return __r0
	}
	iSteamFriends_GetClanChatMessage = func(steamFriends uintptr, clanChatSteamID Uint64SteamID, iMessage int32, prgchText []byte, cchTextMax int32, peChatEntryType *EChatEntryType, psteamidChatter *CSteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetClanChatMessage, uintptr(steamFriends), uintptr(clanChatSteamID), uintptr(iMessage), uintptr(unsafe.Pointer(unsafe.SliceData(prgchText))), uintptr(cchTextMax), uintptr(unsafe.Pointer(peChatEntryType)), uintptr(unsafe.Pointer(psteamidChatter)))
		__r0 := int32(_r0)
		runtime.KeepAlive(prgchText)
		runtime.KeepAlive(peChatEntryType)
		runtime.KeepAlive(psteamidChatter)
		return __r0
	}
	iSteamFriends_IsClanChatAdmin = func(steamFriends uintptr, clanChatSteamID Uint64SteamID, userSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsClanChatAdmin, uintptr(steamFriends), uintptr(clanChatSteamID), uintptr(userSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_IsClanChatWindowOpenInSteam = func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsClanChatWindowOpenInSteam, uintptr(steamFriends), uintptr(clanChatSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_OpenClanChatWindowInSteam = func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_OpenClanChatWindowInSteam, uintptr(steamFriends), uintptr(clanChatSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_CloseClanChatWindowInSteam = func(steamFriends uintptr, clanChatSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_CloseClanChatWindowInSteam, uintptr(steamFriends), uintptr(clanChatSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_SetListenForFriendsMessages = func(steamFriends uintptr, bInterceptEnabled bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_SetListenForFriendsMessages, uintptr(steamFriends), boolToUintptr(bInterceptEnabled))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_ReplyToFriendMessage = func(steamFriends uintptr, friendSteamID Uint64SteamID, pchMsgToSend string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_ReplyToFriendMessage, uintptr(steamFriends), uintptr(friendSteamID), uintptr(unsafe.Pointer(bytePtrFromString(pchMsgToSend))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchMsgToSend)
		return __r0
	}
	iSteamFriends_GetFriendMessage = func(steamFriends uintptr, friendSteamID Uint64SteamID, iMessageID int32, pvData []byte, cubData int32, peChatEntryType *EChatEntryType) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFriendMessage, uintptr(steamFriends), uintptr(friendSteamID), uintptr(iMessageID), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubData), uintptr(unsafe.Pointer(peChatEntryType)))
		__r0 := int32(_r0)
		runtime.KeepAlive(pvData)
		runtime.KeepAlive(peChatEntryType)
		return __r0
	}
	iSteamFriends_GetFollowerCount = func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetFollowerCount, uintptr(steamFriends), uintptr(steamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_IsFollowing = func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsFollowing, uintptr(steamFriends), uintptr(steamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_EnumerateFollowingList = func(steamFriends uintptr, unStartIndex uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_EnumerateFollowingList, uintptr(steamFriends), uintptr(unStartIndex))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_IsClanPublic = func(steamFriends uintptr, clanSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsClanPublic, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_IsClanOfficialGameGroup = func(steamFriends uintptr, clanSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_IsClanOfficialGameGroup, uintptr(steamFriends), uintptr(clanSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_GetNumChatsWithUnreadPriorityMessages = func(steamFriends uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetNumChatsWithUnreadPriorityMessages, uintptr(steamFriends))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog = func(steamFriends uintptr, lobbySteamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayRemotePlayTogetherInviteDialog, uintptr(steamFriends), uintptr(lobbySteamID))
	}
	iSteamFriends_RegisterProtocolInOverlayBrowser = func(steamFriends uintptr, pchProtocol string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_RegisterProtocolInOverlayBrowser, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchProtocol))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchProtocol)
		return __r0
	}
	iSteamFriends_ActivateGameOverlayInviteDialogConnectString = func(steamFriends uintptr, pchConnectString string) {
		purego.SyscallN(addr_iSteamFriends_ActivateGameOverlayInviteDialogConnectString, uintptr(steamFriends), uintptr(unsafe.Pointer(bytePtrFromString(pchConnectString))))
		runtime.KeepAlive(pchConnectString)
	}
	iSteamFriends_RequestEquippedProfileItems = func(steamFriends uintptr, steamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_RequestEquippedProfileItems, uintptr(steamFriends), uintptr(steamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamFriends_BHasEquippedProfileItem = func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_BHasEquippedProfileItem, uintptr(steamFriends), uintptr(steamID), uintptr(itemType))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamFriends_GetProfileItemPropertyString = func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetProfileItemPropertyString, uintptr(steamFriends), uintptr(steamID), uintptr(itemType), uintptr(prop))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamFriends_GetProfileItemPropertyUint = func(steamFriends uintptr, steamID Uint64SteamID, itemType ECommunityProfileItemType, prop ECommunityProfileItemProperty) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamFriends_GetProfileItemPropertyUint, uintptr(steamFriends), uintptr(steamID), uintptr(itemType), uintptr(prop))
		__r0 := uint32(_r0)
		return __r0
	}

}

var (
	addr_steamGameSearch_get                          uintptr
	addr_iSteamGameSearch_AddGameSearchParams         uintptr
	addr_iSteamGameSearch_SearchForGameWithLobby      uintptr
	addr_iSteamGameSearch_SearchForGameSolo           uintptr
	addr_iSteamGameSearch_AcceptGame                  uintptr
	addr_iSteamGameSearch_DeclineGame                 uintptr
	addr_iSteamGameSearch_RetrieveConnectionDetails   uintptr
	addr_iSteamGameSearch_EndGameSearch               uintptr
	addr_iSteamGameSearch_SetGameHostParams           uintptr
	addr_iSteamGameSearch_SetConnectionDetails        uintptr
	addr_iSteamGameSearch_RequestPlayersForGame       uintptr
	addr_iSteamGameSearch_HostConfirmGameStart        uintptr
	addr_iSteamGameSearch_CancelRequestPlayersForGame uintptr
	addr_iSteamGameSearch_SubmitPlayerResult          uintptr
	addr_iSteamGameSearch_EndGame                     uintptr
)

func initGameSearch() {
	var err error
	addr_steamGameSearch_get, err = openSymbol(steamAPILib, flatAPI_SteamGameSearch)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamGameSearch)
	}
	addr_iSteamGameSearch_AddGameSearchParams, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_AddGameSearchParams")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_AddGameSearchParams")
	}
	addr_iSteamGameSearch_SearchForGameWithLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_SearchForGameWithLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_SearchForGameWithLobby")
	}
	addr_iSteamGameSearch_SearchForGameSolo, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_SearchForGameSolo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_SearchForGameSolo")
	}
	addr_iSteamGameSearch_AcceptGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_AcceptGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_AcceptGame")
	}
	addr_iSteamGameSearch_DeclineGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_DeclineGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_DeclineGame")
	}
	addr_iSteamGameSearch_RetrieveConnectionDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_RetrieveConnectionDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_RetrieveConnectionDetails")
	}
	addr_iSteamGameSearch_EndGameSearch, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_EndGameSearch")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_EndGameSearch")
	}
	addr_iSteamGameSearch_SetGameHostParams, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_SetGameHostParams")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_SetGameHostParams")
	}
	addr_iSteamGameSearch_SetConnectionDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_SetConnectionDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_SetConnectionDetails")
	}
	addr_iSteamGameSearch_RequestPlayersForGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_RequestPlayersForGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_RequestPlayersForGame")
	}
	addr_iSteamGameSearch_HostConfirmGameStart, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_HostConfirmGameStart")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_HostConfirmGameStart")
	}
	addr_iSteamGameSearch_CancelRequestPlayersForGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_CancelRequestPlayersForGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_CancelRequestPlayersForGame")
	}
	addr_iSteamGameSearch_SubmitPlayerResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_SubmitPlayerResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_SubmitPlayerResult")
	}
	addr_iSteamGameSearch_EndGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameSearch_EndGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameSearch_EndGame")
	}

	steamGameSearch_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamGameSearch_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamGameSearch_AddGameSearchParams = func(steamGameSearch uintptr, keyToFind string, valuesToFind string) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_AddGameSearchParams, uintptr(steamGameSearch), uintptr(unsafe.Pointer(bytePtrFromString(keyToFind))), uintptr(unsafe.Pointer(bytePtrFromString(valuesToFind))))
		__r0 := EGameSearchErrorCode(_r0)
		runtime.KeepAlive(keyToFind)
		runtime.KeepAlive(valuesToFind)
		return __r0
	}
	iSteamGameSearch_SearchForGameWithLobby = func(steamGameSearch uintptr, lobbySteamID Uint64SteamID, playerMin int32, playerMax int32) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_SearchForGameWithLobby, uintptr(steamGameSearch), uintptr(lobbySteamID), uintptr(playerMin), uintptr(playerMax))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_SearchForGameSolo = func(steamGameSearch uintptr, playerMin int32, playerMax int32) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_SearchForGameSolo, uintptr(steamGameSearch), uintptr(playerMin), uintptr(playerMax))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_AcceptGame = func(steamGameSearch uintptr) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_AcceptGame, uintptr(steamGameSearch))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_DeclineGame = func(steamGameSearch uintptr) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_DeclineGame, uintptr(steamGameSearch))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_RetrieveConnectionDetails = func(steamGameSearch uintptr, hostSteamID Uint64SteamID, cpchConnectionDetails []byte, connectionDetailsSize int32) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_RetrieveConnectionDetails, uintptr(steamGameSearch), uintptr(hostSteamID), uintptr(unsafe.Pointer(unsafe.SliceData(cpchConnectionDetails))), uintptr(connectionDetailsSize))
		__r0 := EGameSearchErrorCode(_r0)
		runtime.KeepAlive(cpchConnectionDetails)
		return __r0
	}
	iSteamGameSearch_EndGameSearch = func(steamGameSearch uintptr) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_EndGameSearch, uintptr(steamGameSearch))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_SetGameHostParams = func(steamGameSearch uintptr, key string, value string) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_SetGameHostParams, uintptr(steamGameSearch), uintptr(unsafe.Pointer(bytePtrFromString(key))), uintptr(unsafe.Pointer(bytePtrFromString(value))))
		__r0 := EGameSearchErrorCode(_r0)
		runtime.KeepAlive(key)
		runtime.KeepAlive(value)
		return __r0
	}
	iSteamGameSearch_SetConnectionDetails = func(steamGameSearch uintptr, cpchConnectionDetails string, connectionDetailsSize int32) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_SetConnectionDetails, uintptr(steamGameSearch), uintptr(unsafe.Pointer(bytePtrFromString(cpchConnectionDetails))), uintptr(connectionDetailsSize))
		__r0 := EGameSearchErrorCode(_r0)
		runtime.KeepAlive(cpchConnectionDetails)
		return __r0
	}
	iSteamGameSearch_RequestPlayersForGame = func(steamGameSearch uintptr, playerMin int32, playerMax int32, maxTeamSize int32) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_RequestPlayersForGame, uintptr(steamGameSearch), uintptr(playerMin), uintptr(playerMax), uintptr(maxTeamSize))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_HostConfirmGameStart = func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_HostConfirmGameStart, uintptr(steamGameSearch), uintptr(uniqueGameID))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_CancelRequestPlayersForGame = func(steamGameSearch uintptr) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_CancelRequestPlayersForGame, uintptr(steamGameSearch))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_SubmitPlayerResult = func(steamGameSearch uintptr, uniqueGameID uint64, playerSteamID Uint64SteamID, EPlayerResult EPlayerResult) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_SubmitPlayerResult, uintptr(steamGameSearch), uintptr(uniqueGameID), uintptr(playerSteamID), uintptr(EPlayerResult))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
	iSteamGameSearch_EndGame = func(steamGameSearch uintptr, uniqueGameID uint64) EGameSearchErrorCode {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameSearch_EndGame, uintptr(steamGameSearch), uintptr(uniqueGameID))
		__r0 := EGameSearchErrorCode(_r0)
		return __r0
	}
}

var (
	addr_steamGameServer_get                                  uintptr
	addr_iSteamGameServer_SetProduct                          uintptr
	addr_iSteamGameServer_SetGameDescription                  uintptr
	addr_iSteamGameServer_SetModDir                           uintptr
	addr_iSteamGameServer_SetDedicatedServer                  uintptr
	addr_iSteamGameServer_LogOn                               uintptr
	addr_iSteamGameServer_LogOnAnonymous                      uintptr
	addr_iSteamGameServer_LogOff                              uintptr
	addr_iSteamGameServer_BLoggedOn                           uintptr
	addr_iSteamGameServer_BSecure                             uintptr
	addr_iSteamGameServer_GetSteamID                          uintptr
	addr_iSteamGameServer_WasRestartRequested                 uintptr
	addr_iSteamGameServer_SetMaxPlayerCount                   uintptr
	addr_iSteamGameServer_SetBotPlayerCount                   uintptr
	addr_iSteamGameServer_SetServerName                       uintptr
	addr_iSteamGameServer_SetMapName                          uintptr
	addr_iSteamGameServer_SetPasswordProtected                uintptr
	addr_iSteamGameServer_SetSpectatorPort                    uintptr
	addr_iSteamGameServer_SetSpectatorServerName              uintptr
	addr_iSteamGameServer_ClearAllKeyValues                   uintptr
	addr_iSteamGameServer_SetKeyValue                         uintptr
	addr_iSteamGameServer_SetGameTags                         uintptr
	addr_iSteamGameServer_SetGameData                         uintptr
	addr_iSteamGameServer_SetRegion                           uintptr
	addr_iSteamGameServer_SetAdvertiseServerActive            uintptr
	addr_iSteamGameServer_GetAuthSessionTicket                uintptr
	addr_iSteamGameServer_BeginAuthSession                    uintptr
	addr_iSteamGameServer_EndAuthSession                      uintptr
	addr_iSteamGameServer_CancelAuthTicket                    uintptr
	addr_iSteamGameServer_UserHasLicenseForApp                uintptr
	addr_iSteamGameServer_RequestUserGroupStatus              uintptr
	addr_iSteamGameServer_GetPublicIP                         uintptr
	addr_iSteamGameServer_HandleIncomingPacket                uintptr
	addr_iSteamGameServer_GetNextOutgoingPacket               uintptr
	addr_iSteamGameServer_AssociateWithClan                   uintptr
	addr_iSteamGameServer_ComputeNewPlayerCompatibility       uintptr
	addr_iSteamGameServer_CreateUnauthenticatedUserConnection uintptr
	addr_iSteamGameServer_BUpdateUserData                     uintptr
)

func initGameServer() {
	var err error
	addr_steamGameServer_get, err = openSymbol(steamAPILib, flatAPI_SteamGameServer)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamGameServer)
	}
	addr_iSteamGameServer_SetProduct, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetProduct")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetProduct")
	}
	addr_iSteamGameServer_SetGameDescription, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetGameDescription")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetGameDescription")
	}
	addr_iSteamGameServer_SetModDir, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetModDir")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetModDir")
	}
	addr_iSteamGameServer_SetDedicatedServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetDedicatedServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetDedicatedServer")
	}
	addr_iSteamGameServer_LogOn, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_LogOn")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_LogOn")
	}
	addr_iSteamGameServer_LogOnAnonymous, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_LogOnAnonymous")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_LogOnAnonymous")
	}
	addr_iSteamGameServer_LogOff, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_LogOff")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_LogOff")
	}
	addr_iSteamGameServer_BLoggedOn, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_BLoggedOn")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_BLoggedOn")
	}
	addr_iSteamGameServer_BSecure, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_BSecure")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_BSecure")
	}
	addr_iSteamGameServer_GetSteamID, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_GetSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_GetSteamID")
	}
	addr_iSteamGameServer_WasRestartRequested, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_WasRestartRequested")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_WasRestartRequested")
	}
	addr_iSteamGameServer_SetMaxPlayerCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetMaxPlayerCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetMaxPlayerCount")
	}
	addr_iSteamGameServer_SetBotPlayerCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetBotPlayerCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetBotPlayerCount")
	}
	addr_iSteamGameServer_SetServerName, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetServerName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetServerName")
	}
	addr_iSteamGameServer_SetMapName, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetMapName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetMapName")
	}
	addr_iSteamGameServer_SetPasswordProtected, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetPasswordProtected")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetPasswordProtected")
	}
	addr_iSteamGameServer_SetSpectatorPort, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetSpectatorPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetSpectatorPort")
	}
	addr_iSteamGameServer_SetSpectatorServerName, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetSpectatorServerName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetSpectatorServerName")
	}
	addr_iSteamGameServer_ClearAllKeyValues, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_ClearAllKeyValues")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_ClearAllKeyValues")
	}
	addr_iSteamGameServer_SetKeyValue, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetKeyValue")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetKeyValue")
	}
	addr_iSteamGameServer_SetGameTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetGameTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetGameTags")
	}
	addr_iSteamGameServer_SetGameData, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetGameData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetGameData")
	}
	addr_iSteamGameServer_SetRegion, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetRegion")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetRegion")
	}
	addr_iSteamGameServer_SetAdvertiseServerActive, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_SetAdvertiseServerActive")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_SetAdvertiseServerActive")
	}
	addr_iSteamGameServer_GetAuthSessionTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_GetAuthSessionTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_GetAuthSessionTicket")
	}
	addr_iSteamGameServer_BeginAuthSession, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_BeginAuthSession")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_BeginAuthSession")
	}
	addr_iSteamGameServer_EndAuthSession, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_EndAuthSession")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_EndAuthSession")
	}
	addr_iSteamGameServer_CancelAuthTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_CancelAuthTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_CancelAuthTicket")
	}
	addr_iSteamGameServer_UserHasLicenseForApp, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_UserHasLicenseForApp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_UserHasLicenseForApp")
	}
	addr_iSteamGameServer_RequestUserGroupStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_RequestUserGroupStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_RequestUserGroupStatus")
	}
	addr_iSteamGameServer_GetPublicIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_GetPublicIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_GetPublicIP")
	}
	addr_iSteamGameServer_HandleIncomingPacket, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_HandleIncomingPacket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_HandleIncomingPacket")
	}
	addr_iSteamGameServer_GetNextOutgoingPacket, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_GetNextOutgoingPacket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_GetNextOutgoingPacket")
	}
	addr_iSteamGameServer_AssociateWithClan, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_AssociateWithClan")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_AssociateWithClan")
	}
	addr_iSteamGameServer_ComputeNewPlayerCompatibility, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_ComputeNewPlayerCompatibility")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_ComputeNewPlayerCompatibility")
	}
	addr_iSteamGameServer_CreateUnauthenticatedUserConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_CreateUnauthenticatedUserConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_CreateUnauthenticatedUserConnection")
	}
	addr_iSteamGameServer_BUpdateUserData, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServer_BUpdateUserData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServer_BUpdateUserData")
	}

	steamGameServer_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamGameServer_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamGameServer_SetProduct = func(steamGameServer uintptr, pszProduct string) {
		purego.SyscallN(addr_iSteamGameServer_SetProduct, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(pszProduct))))
		runtime.KeepAlive(pszProduct)
	}
	iSteamGameServer_SetGameDescription = func(steamGameServer uintptr, pszGameDescription string) {
		purego.SyscallN(addr_iSteamGameServer_SetGameDescription, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(pszGameDescription))))
		runtime.KeepAlive(pszGameDescription)
	}
	iSteamGameServer_SetModDir = func(steamGameServer uintptr, modDir string) {
		purego.SyscallN(addr_iSteamGameServer_SetModDir, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(modDir))))
		runtime.KeepAlive(modDir)
	}
	iSteamGameServer_SetDedicatedServer = func(steamGameServer uintptr, dedicated bool) {
		purego.SyscallN(addr_iSteamGameServer_SetDedicatedServer, uintptr(steamGameServer), boolToUintptr(dedicated))
	}
	iSteamGameServer_LogOn = func(steamGameServer uintptr, token string) {
		purego.SyscallN(addr_iSteamGameServer_LogOn, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(token))))
		runtime.KeepAlive(token)
	}
	iSteamGameServer_LogOnAnonymous = func(steamGameServer uintptr) {
		purego.SyscallN(addr_iSteamGameServer_LogOnAnonymous, uintptr(steamGameServer))
	}
	iSteamGameServer_LogOff = func(steamGameServer uintptr) {
		purego.SyscallN(addr_iSteamGameServer_LogOff, uintptr(steamGameServer))
	}
	iSteamGameServer_BLoggedOn = func(steamGameServer uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_BLoggedOn, uintptr(steamGameServer))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamGameServer_BSecure = func(steamGameServer uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_BSecure, uintptr(steamGameServer))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamGameServer_GetSteamID = func(steamGameServer uintptr) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_GetSteamID, uintptr(steamGameServer))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamGameServer_WasRestartRequested = func(steamGameServer uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_WasRestartRequested, uintptr(steamGameServer))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamGameServer_SetMaxPlayerCount = func(steamGameServer uintptr, playersMax int32) {
		purego.SyscallN(addr_iSteamGameServer_SetMaxPlayerCount, uintptr(steamGameServer), uintptr(playersMax))
	}
	iSteamGameServer_SetBotPlayerCount = func(steamGameServer uintptr, botplayers int32) {
		purego.SyscallN(addr_iSteamGameServer_SetBotPlayerCount, uintptr(steamGameServer), uintptr(botplayers))
	}
	iSteamGameServer_SetServerName = func(steamGameServer uintptr, serverName string) {
		purego.SyscallN(addr_iSteamGameServer_SetServerName, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(serverName))))
		runtime.KeepAlive(serverName)
	}
	iSteamGameServer_SetMapName = func(steamGameServer uintptr, mapName string) {
		purego.SyscallN(addr_iSteamGameServer_SetMapName, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(mapName))))
		runtime.KeepAlive(mapName)
	}
	iSteamGameServer_SetPasswordProtected = func(steamGameServer uintptr, passwordProtected bool) {
		purego.SyscallN(addr_iSteamGameServer_SetPasswordProtected, uintptr(steamGameServer), boolToUintptr(passwordProtected))
	}
	iSteamGameServer_SetSpectatorPort = func(steamGameServer uintptr, spectatorPort uint16) {
		purego.SyscallN(addr_iSteamGameServer_SetSpectatorPort, uintptr(steamGameServer), uintptr(spectatorPort))
	}
	iSteamGameServer_SetSpectatorServerName = func(steamGameServer uintptr, spectatorServerName string) {
		purego.SyscallN(addr_iSteamGameServer_SetSpectatorServerName, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(spectatorServerName))))
		runtime.KeepAlive(spectatorServerName)
	}
	iSteamGameServer_ClearAllKeyValues = func(steamGameServer uintptr) {
		purego.SyscallN(addr_iSteamGameServer_ClearAllKeyValues, uintptr(steamGameServer))
	}
	iSteamGameServer_SetKeyValue = func(steamGameServer uintptr, key string, value string) {
		purego.SyscallN(addr_iSteamGameServer_SetKeyValue, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(key))), uintptr(unsafe.Pointer(bytePtrFromString(value))))
		runtime.KeepAlive(key)
		runtime.KeepAlive(value)
	}
	iSteamGameServer_SetGameTags = func(steamGameServer uintptr, gameTags string) {
		purego.SyscallN(addr_iSteamGameServer_SetGameTags, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(gameTags))))
		runtime.KeepAlive(gameTags)
	}
	iSteamGameServer_SetGameData = func(steamGameServer uintptr, gameData string) {
		purego.SyscallN(addr_iSteamGameServer_SetGameData, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(gameData))))
		runtime.KeepAlive(gameData)
	}
	iSteamGameServer_SetRegion = func(steamGameServer uintptr, region string) {
		purego.SyscallN(addr_iSteamGameServer_SetRegion, uintptr(steamGameServer), uintptr(unsafe.Pointer(bytePtrFromString(region))))
		runtime.KeepAlive(region)
	}
	iSteamGameServer_SetAdvertiseServerActive = func(steamGameServer uintptr, active bool) {
		purego.SyscallN(addr_iSteamGameServer_SetAdvertiseServerActive, uintptr(steamGameServer), boolToUintptr(active))
	}
	iSteamGameServer_GetAuthSessionTicket = func(steamGameServer uintptr, pTicket []byte, maxTicket int32, pcbTicket *uint32, snid *SteamNetworkingIdentity) HAuthTicket {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_GetAuthSessionTicket, uintptr(steamGameServer), uintptr(unsafe.Pointer(unsafe.SliceData(pTicket))), uintptr(maxTicket), uintptr(unsafe.Pointer(pcbTicket)), uintptr(unsafe.Pointer(snid)))
		__r0 := HAuthTicket(_r0)
		runtime.KeepAlive(pTicket)
		runtime.KeepAlive(pcbTicket)
		runtime.KeepAlive(snid)
		return __r0
	}
	iSteamGameServer_BeginAuthSession = func(steamGameServer uintptr, authTicket []byte, cbAuthTicket int32, steamID Uint64SteamID) EBeginAuthSessionResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_BeginAuthSession, uintptr(steamGameServer), uintptr(unsafe.Pointer(unsafe.SliceData(authTicket))), uintptr(cbAuthTicket), uintptr(steamID))
		__r0 := EBeginAuthSessionResult(_r0)
		runtime.KeepAlive(authTicket)
		return __r0
	}
	iSteamGameServer_EndAuthSession = func(steamGameServer uintptr, steamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamGameServer_EndAuthSession, uintptr(steamGameServer), uintptr(steamID))
	}
	iSteamGameServer_CancelAuthTicket = func(steamGameServer uintptr, hAuthTicket HAuthTicket) {
		purego.SyscallN(addr_iSteamGameServer_CancelAuthTicket, uintptr(steamGameServer), uintptr(hAuthTicket))
	}
	iSteamGameServer_UserHasLicenseForApp = func(steamGameServer uintptr, steamID Uint64SteamID, AppId AppId_t) EUserHasLicenseForAppResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_UserHasLicenseForApp, uintptr(steamGameServer), uintptr(steamID), uintptr(AppId))
		__r0 := EUserHasLicenseForAppResult(_r0)
		return __r0
	}
	iSteamGameServer_RequestUserGroupStatus = func(steamGameServer uintptr, user Uint64SteamID, groupSteamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_RequestUserGroupStatus, uintptr(steamGameServer), uintptr(user), uintptr(groupSteamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamGameServer_GetPublicIP = func(steamGameServer uintptr) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_GetPublicIP, uintptr(steamGameServer))
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamGameServer_HandleIncomingPacket = func(steamGameServer uintptr, pData []byte, data int32, srcIP uint32, srcPort uint16) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_HandleIncomingPacket, uintptr(steamGameServer), uintptr(unsafe.Pointer(unsafe.SliceData(pData))), uintptr(data), uintptr(srcIP), uintptr(srcPort))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamGameServer_GetNextOutgoingPacket = func(steamGameServer uintptr, pOut []byte, maxOut int32, pNetAdr *uint32, pPort *uint16) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_GetNextOutgoingPacket, uintptr(steamGameServer), uintptr(unsafe.Pointer(unsafe.SliceData(pOut))), uintptr(maxOut), uintptr(unsafe.Pointer(pNetAdr)), uintptr(unsafe.Pointer(pPort)))
		__r0 := int32(_r0)
		runtime.KeepAlive(pOut)
		runtime.KeepAlive(pNetAdr)
		runtime.KeepAlive(pPort)
		return __r0
	}
	iSteamGameServer_AssociateWithClan = func(steamGameServer uintptr, clanSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_AssociateWithClan, uintptr(steamGameServer), uintptr(clanSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamGameServer_ComputeNewPlayerCompatibility = func(steamGameServer uintptr, newPlayerSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_ComputeNewPlayerCompatibility, uintptr(steamGameServer), uintptr(newPlayerSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamGameServer_CreateUnauthenticatedUserConnection = func(steamGameServer uintptr) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_CreateUnauthenticatedUserConnection, uintptr(steamGameServer))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamGameServer_BUpdateUserData = func(steamGameServer uintptr, user Uint64SteamID, playerName string, score uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServer_BUpdateUserData, uintptr(steamGameServer), uintptr(user), uintptr(unsafe.Pointer(bytePtrFromString(playerName))), uintptr(score))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(playerName)
		return __r0
	}

}

var (
	addr_steamGameServerStats_get                   uintptr
	addr_iSteamGameServerStats_RequestUserStats     uintptr
	addr_iSteamGameServerStats_GetUserStatInt32     uintptr
	addr_iSteamGameServerStats_GetUserAchievement   uintptr
	addr_iSteamGameServerStats_SetUserStatInt32     uintptr
	addr_iSteamGameServerStats_SetUserAchievement   uintptr
	addr_iSteamGameServerStats_ClearUserAchievement uintptr
	addr_iSteamGameServerStats_StoreUserStats       uintptr
)

func initGameServerStats() {
	var err error
	addr_steamGameServerStats_get, err = openSymbol(steamAPILib, flatAPI_SteamGameServerStats)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamGameServerStats)
	}
	addr_iSteamGameServerStats_RequestUserStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_RequestUserStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_RequestUserStats")
	}
	addr_iSteamGameServerStats_GetUserStatInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_GetUserStatInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_GetUserStatInt32")
	}
	addr_iSteamGameServerStats_GetUserAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_GetUserAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_GetUserAchievement")
	}
	addr_iSteamGameServerStats_SetUserStatInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_SetUserStatInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_SetUserStatInt32")
	}
	addr_iSteamGameServerStats_SetUserAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_SetUserAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_SetUserAchievement")
	}
	addr_iSteamGameServerStats_ClearUserAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_ClearUserAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_ClearUserAchievement")
	}
	addr_iSteamGameServerStats_StoreUserStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamGameServerStats_StoreUserStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamGameServerStats_StoreUserStats")
	}

	steamGameServerStats_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamGameServerStats_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamGameServerStats_RequestUserStats = func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_RequestUserStats, uintptr(steamGameServerStats), uintptr(userSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamGameServerStats_GetUserStatInt32 = func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pData *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_GetUserStatInt32, uintptr(steamGameServerStats), uintptr(userSteamID), uintptr(unsafe.Pointer(bytePtrFromString(name))), uintptr(unsafe.Pointer(pData)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamGameServerStats_GetUserAchievement = func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, pbAchieved *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_GetUserAchievement, uintptr(steamGameServerStats), uintptr(userSteamID), uintptr(unsafe.Pointer(bytePtrFromString(name))), uintptr(unsafe.Pointer(pbAchieved)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		runtime.KeepAlive(pbAchieved)
		return __r0
	}
	iSteamGameServerStats_SetUserStatInt32 = func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string, nData int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_SetUserStatInt32, uintptr(steamGameServerStats), uintptr(userSteamID), uintptr(unsafe.Pointer(bytePtrFromString(name))), uintptr(nData))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		return __r0
	}
	iSteamGameServerStats_SetUserAchievement = func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_SetUserAchievement, uintptr(steamGameServerStats), uintptr(userSteamID), uintptr(unsafe.Pointer(bytePtrFromString(name))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		return __r0
	}
	iSteamGameServerStats_ClearUserAchievement = func(steamGameServerStats uintptr, userSteamID Uint64SteamID, name string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_ClearUserAchievement, uintptr(steamGameServerStats), uintptr(userSteamID), uintptr(unsafe.Pointer(bytePtrFromString(name))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(name)
		return __r0
	}
	iSteamGameServerStats_StoreUserStats = func(steamGameServerStats uintptr, userSteamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamGameServerStats_StoreUserStats, uintptr(steamGameServerStats), uintptr(userSteamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}

	purego.RegisterLibFunc(&iSteamGameServerStats_GetUserStatFloat, steamAPILib, flatAPI_ISteamGameServerStats_GetUserStatFloat)
	purego.RegisterLibFunc(&iSteamGameServerStats_SetUserStatFloat, steamAPILib, flatAPI_ISteamGameServerStats_SetUserStatFloat)
	purego.RegisterLibFunc(&iSteamGameServerStats_UpdateUserAvgRateStat, steamAPILib, flatAPI_ISteamGameServerStats_UpdateUserAvgRateStat)

}

var (
	addr_steamInput_get                                   uintptr
	addr_iSteamInput_Init                                 uintptr
	addr_iSteamInput_Shutdown                             uintptr
	addr_iSteamInput_SetInputActionManifestFilePath       uintptr
	addr_iSteamInput_RunFrame                             uintptr
	addr_iSteamInput_BWaitForData                         uintptr
	addr_iSteamInput_BNewDataAvailable                    uintptr
	addr_iSteamInput_GetConnectedControllers              uintptr
	addr_iSteamInput_EnableDeviceCallbacks                uintptr
	addr_iSteamInput_EnableActionEventCallbacks           uintptr
	addr_iSteamInput_GetActionSetHandle                   uintptr
	addr_iSteamInput_ActivateActionSet                    uintptr
	addr_iSteamInput_GetCurrentActionSet                  uintptr
	addr_iSteamInput_ActivateActionSetLayer               uintptr
	addr_iSteamInput_DeactivateActionSetLayer             uintptr
	addr_iSteamInput_DeactivateAllActionSetLayers         uintptr
	addr_iSteamInput_GetActiveActionSetLayers             uintptr
	addr_iSteamInput_GetDigitalActionHandle               uintptr
	addr_iSteamInput_GetDigitalActionData                 uintptr
	addr_iSteamInput_GetDigitalActionOrigins              uintptr
	addr_iSteamInput_GetStringForDigitalActionName        uintptr
	addr_iSteamInput_GetAnalogActionHandle                uintptr
	addr_iSteamInput_GetAnalogActionData                  uintptr
	addr_iSteamInput_GetAnalogActionOrigins               uintptr
	addr_iSteamInput_GetGlyphPNGForActionOrigin           uintptr
	addr_iSteamInput_GetGlyphSVGForActionOrigin           uintptr
	addr_iSteamInput_GetGlyphForActionOrigin_Legacy       uintptr
	addr_iSteamInput_GetStringForActionOrigin             uintptr
	addr_iSteamInput_GetStringForAnalogActionName         uintptr
	addr_iSteamInput_StopAnalogActionMomentum             uintptr
	addr_iSteamInput_GetMotionData                        uintptr
	addr_iSteamInput_TriggerVibration                     uintptr
	addr_iSteamInput_TriggerVibrationExtended             uintptr
	addr_iSteamInput_TriggerSimpleHapticEvent             uintptr
	addr_iSteamInput_SetLEDColor                          uintptr
	addr_iSteamInput_Legacy_TriggerHapticPulse            uintptr
	addr_iSteamInput_Legacy_TriggerRepeatedHapticPulse    uintptr
	addr_iSteamInput_ShowBindingPanel                     uintptr
	addr_iSteamInput_GetInputTypeForHandle                uintptr
	addr_iSteamInput_GetControllerForGamepadIndex         uintptr
	addr_iSteamInput_GetGamepadIndexForController         uintptr
	addr_iSteamInput_GetStringForXboxOrigin               uintptr
	addr_iSteamInput_GetGlyphForXboxOrigin                uintptr
	addr_iSteamInput_GetActionOriginFromXboxOrigin        uintptr
	addr_iSteamInput_TranslateActionOrigin                uintptr
	addr_iSteamInput_GetDeviceBindingRevision             uintptr
	addr_iSteamInput_GetRemotePlaySessionID               uintptr
	addr_iSteamInput_GetSessionInputConfigurationSettings uintptr
	addr_iSteamInput_SetDualSenseTriggerEffect            uintptr
)

func initInput() {
	var err error
	addr_steamInput_get, err = openSymbol(steamAPILib, flatAPI_SteamInput)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamInput)
	}
	addr_iSteamInput_Init, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_Init")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_Init")
	}
	addr_iSteamInput_Shutdown, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_Shutdown")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_Shutdown")
	}
	addr_iSteamInput_SetInputActionManifestFilePath, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_SetInputActionManifestFilePath")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_SetInputActionManifestFilePath")
	}
	addr_iSteamInput_RunFrame, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_RunFrame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_RunFrame")
	}
	addr_iSteamInput_BWaitForData, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_BWaitForData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_BWaitForData")
	}
	addr_iSteamInput_BNewDataAvailable, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_BNewDataAvailable")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_BNewDataAvailable")
	}
	addr_iSteamInput_GetConnectedControllers, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetConnectedControllers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetConnectedControllers")
	}
	addr_iSteamInput_EnableDeviceCallbacks, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_EnableDeviceCallbacks")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_EnableDeviceCallbacks")
	}
	addr_iSteamInput_EnableActionEventCallbacks, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_EnableActionEventCallbacks")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_EnableActionEventCallbacks")
	}
	addr_iSteamInput_GetActionSetHandle, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetActionSetHandle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetActionSetHandle")
	}
	addr_iSteamInput_ActivateActionSet, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_ActivateActionSet")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_ActivateActionSet")
	}
	addr_iSteamInput_GetCurrentActionSet, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetCurrentActionSet")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetCurrentActionSet")
	}
	addr_iSteamInput_ActivateActionSetLayer, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_ActivateActionSetLayer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_ActivateActionSetLayer")
	}
	addr_iSteamInput_DeactivateActionSetLayer, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_DeactivateActionSetLayer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_DeactivateActionSetLayer")
	}
	addr_iSteamInput_DeactivateAllActionSetLayers, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_DeactivateAllActionSetLayers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_DeactivateAllActionSetLayers")
	}
	addr_iSteamInput_GetActiveActionSetLayers, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetActiveActionSetLayers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetActiveActionSetLayers")
	}
	addr_iSteamInput_GetDigitalActionHandle, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetDigitalActionHandle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetDigitalActionHandle")
	}
	addr_iSteamInput_GetDigitalActionData, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetDigitalActionData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetDigitalActionData")
	}
	addr_iSteamInput_GetDigitalActionOrigins, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetDigitalActionOrigins")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetDigitalActionOrigins")
	}
	addr_iSteamInput_GetStringForDigitalActionName, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetStringForDigitalActionName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetStringForDigitalActionName")
	}
	addr_iSteamInput_GetAnalogActionHandle, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetAnalogActionHandle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetAnalogActionHandle")
	}
	addr_iSteamInput_GetAnalogActionData, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetAnalogActionData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetAnalogActionData")
	}
	addr_iSteamInput_GetAnalogActionOrigins, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetAnalogActionOrigins")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetAnalogActionOrigins")
	}
	addr_iSteamInput_GetGlyphPNGForActionOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetGlyphPNGForActionOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetGlyphPNGForActionOrigin")
	}
	addr_iSteamInput_GetGlyphSVGForActionOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetGlyphSVGForActionOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetGlyphSVGForActionOrigin")
	}
	addr_iSteamInput_GetGlyphForActionOrigin_Legacy, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetGlyphForActionOrigin_Legacy")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetGlyphForActionOrigin_Legacy")
	}
	addr_iSteamInput_GetStringForActionOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetStringForActionOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetStringForActionOrigin")
	}
	addr_iSteamInput_GetStringForAnalogActionName, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetStringForAnalogActionName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetStringForAnalogActionName")
	}
	addr_iSteamInput_StopAnalogActionMomentum, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_StopAnalogActionMomentum")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_StopAnalogActionMomentum")
	}
	addr_iSteamInput_GetMotionData, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetMotionData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetMotionData")
	}
	addr_iSteamInput_TriggerVibration, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_TriggerVibration")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_TriggerVibration")
	}
	addr_iSteamInput_TriggerVibrationExtended, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_TriggerVibrationExtended")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_TriggerVibrationExtended")
	}
	addr_iSteamInput_TriggerSimpleHapticEvent, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_TriggerSimpleHapticEvent")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_TriggerSimpleHapticEvent")
	}
	addr_iSteamInput_SetLEDColor, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_SetLEDColor")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_SetLEDColor")
	}
	addr_iSteamInput_Legacy_TriggerHapticPulse, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_Legacy_TriggerHapticPulse")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_Legacy_TriggerHapticPulse")
	}
	addr_iSteamInput_Legacy_TriggerRepeatedHapticPulse, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse")
	}
	addr_iSteamInput_ShowBindingPanel, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_ShowBindingPanel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_ShowBindingPanel")
	}
	addr_iSteamInput_GetInputTypeForHandle, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetInputTypeForHandle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetInputTypeForHandle")
	}
	addr_iSteamInput_GetControllerForGamepadIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetControllerForGamepadIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetControllerForGamepadIndex")
	}
	addr_iSteamInput_GetGamepadIndexForController, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetGamepadIndexForController")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetGamepadIndexForController")
	}
	addr_iSteamInput_GetStringForXboxOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetStringForXboxOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetStringForXboxOrigin")
	}
	addr_iSteamInput_GetGlyphForXboxOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetGlyphForXboxOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetGlyphForXboxOrigin")
	}
	addr_iSteamInput_GetActionOriginFromXboxOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetActionOriginFromXboxOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetActionOriginFromXboxOrigin")
	}
	addr_iSteamInput_TranslateActionOrigin, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_TranslateActionOrigin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_TranslateActionOrigin")
	}
	addr_iSteamInput_GetDeviceBindingRevision, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetDeviceBindingRevision")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetDeviceBindingRevision")
	}
	addr_iSteamInput_GetRemotePlaySessionID, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetRemotePlaySessionID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetRemotePlaySessionID")
	}
	addr_iSteamInput_GetSessionInputConfigurationSettings, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_GetSessionInputConfigurationSettings")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_GetSessionInputConfigurationSettings")
	}
	addr_iSteamInput_SetDualSenseTriggerEffect, err = openSymbol(steamAPILib, "SteamAPI_ISteamInput_SetDualSenseTriggerEffect")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInput_SetDualSenseTriggerEffect")
	}

	steamInput_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamInput_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamInput_Init = func(steamInput uintptr, explicitlyCallRunFrame bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_Init, uintptr(steamInput), boolToUintptr(explicitlyCallRunFrame))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInput_Shutdown = func(steamInput uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_Shutdown, uintptr(steamInput))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInput_SetInputActionManifestFilePath = func(steamInput uintptr, inputActionManifestAbsolutePath string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_SetInputActionManifestFilePath, uintptr(steamInput), uintptr(unsafe.Pointer(bytePtrFromString(inputActionManifestAbsolutePath))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(inputActionManifestAbsolutePath)
		return __r0
	}
	iSteamInput_RunFrame = func(steamInput uintptr, reservedValue bool) {
		purego.SyscallN(addr_iSteamInput_RunFrame, uintptr(steamInput), boolToUintptr(reservedValue))
	}
	iSteamInput_BWaitForData = func(steamInput uintptr, waitForever bool, timeout uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_BWaitForData, uintptr(steamInput), boolToUintptr(waitForever), uintptr(timeout))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInput_BNewDataAvailable = func(steamInput uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_BNewDataAvailable, uintptr(steamInput))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInput_GetConnectedControllers = func(steamInput uintptr, handlesOut *[_STEAM_INPUT_MAX_COUNT]InputHandle_t) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetConnectedControllers, uintptr(steamInput), uintptr(unsafe.Pointer(handlesOut)))
		__r0 := int32(_r0)
		runtime.KeepAlive(handlesOut)
		return __r0
	}
	iSteamInput_EnableDeviceCallbacks = func(steamInput uintptr) {
		purego.SyscallN(addr_iSteamInput_EnableDeviceCallbacks, uintptr(steamInput))
	}
	iSteamInput_EnableActionEventCallbacks = func(steamInput uintptr, pCallback SteamInputActionEventCallbackPointer) {
		purego.SyscallN(addr_iSteamInput_EnableActionEventCallbacks, uintptr(steamInput), uintptr(pCallback))
	}
	iSteamInput_GetActionSetHandle = func(steamInput uintptr, pszActionSetName string) InputActionSetHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetActionSetHandle, uintptr(steamInput), uintptr(unsafe.Pointer(bytePtrFromString(pszActionSetName))))
		__r0 := InputActionSetHandle(_r0)
		runtime.KeepAlive(pszActionSetName)
		return __r0
	}
	iSteamInput_ActivateActionSet = func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle) {
		purego.SyscallN(addr_iSteamInput_ActivateActionSet, uintptr(steamInput), uintptr(inputHandle), uintptr(actionSetHandle))
	}
	iSteamInput_GetCurrentActionSet = func(steamInput uintptr, inputHandle InputHandle_t) InputActionSetHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetCurrentActionSet, uintptr(steamInput), uintptr(inputHandle))
		__r0 := InputActionSetHandle(_r0)
		return __r0
	}
	iSteamInput_ActivateActionSetLayer = func(steamInput uintptr, inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle) {
		purego.SyscallN(addr_iSteamInput_ActivateActionSetLayer, uintptr(steamInput), uintptr(inputHandle), uintptr(actionSetLayerHandle))
	}
	iSteamInput_DeactivateActionSetLayer = func(steamInput uintptr, inputHandle InputHandle_t, actionSetLayerHandle InputActionSetHandle) {
		purego.SyscallN(addr_iSteamInput_DeactivateActionSetLayer, uintptr(steamInput), uintptr(inputHandle), uintptr(actionSetLayerHandle))
	}
	iSteamInput_DeactivateAllActionSetLayers = func(steamInput uintptr, inputHandle InputHandle_t) {
		purego.SyscallN(addr_iSteamInput_DeactivateAllActionSetLayers, uintptr(steamInput), uintptr(inputHandle))
	}
	iSteamInput_GetActiveActionSetLayers = func(steamInput uintptr, inputHandle InputHandle_t, handlesOut *[_STEAM_INPUT_MAX_ACTIVE_LAYERS]InputActionSetHandle) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetActiveActionSetLayers, uintptr(steamInput), uintptr(inputHandle), uintptr(unsafe.Pointer(handlesOut)))
		__r0 := int32(_r0)
		runtime.KeepAlive(handlesOut)
		return __r0
	}
	iSteamInput_GetDigitalActionHandle = func(steamInput uintptr, actionName string) InputDigitalActionHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetDigitalActionHandle, uintptr(steamInput), uintptr(unsafe.Pointer(bytePtrFromString(actionName))))
		__r0 := InputDigitalActionHandle(_r0)
		runtime.KeepAlive(actionName)
		return __r0
	}
	iSteamInput_GetDigitalActionData = func(steamInput uintptr, inputHandle InputHandle_t, digitalActionHandle InputDigitalActionHandle) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetDigitalActionData, uintptr(steamInput), uintptr(inputHandle), uintptr(digitalActionHandle))
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamInput_GetDigitalActionOrigins = func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle, originsOut *[_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetDigitalActionOrigins, uintptr(steamInput), uintptr(inputHandle), uintptr(actionSetHandle), uintptr(digitalActionHandle), uintptr(unsafe.Pointer(originsOut)))
		__r0 := int32(_r0)
		runtime.KeepAlive(originsOut)
		return __r0
	}
	iSteamInput_GetStringForDigitalActionName = func(steamInput uintptr, actionHandle InputDigitalActionHandle) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetStringForDigitalActionName, uintptr(steamInput), uintptr(actionHandle))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamInput_GetAnalogActionHandle = func(steamInput uintptr, actionName string) InputAnalogActionHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetAnalogActionHandle, uintptr(steamInput), uintptr(unsafe.Pointer(bytePtrFromString(actionName))))
		__r0 := InputAnalogActionHandle(_r0)
		runtime.KeepAlive(actionName)
		return __r0
	}
	iSteamInput_GetAnalogActionData = func(steamInput uintptr, inputHandle InputHandle_t, analogActionHandle InputAnalogActionHandle) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetAnalogActionData, uintptr(steamInput), uintptr(inputHandle), uintptr(analogActionHandle))
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamInput_GetAnalogActionOrigins = func(steamInput uintptr, inputHandle InputHandle_t, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle, originsOut *[_STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetAnalogActionOrigins, uintptr(steamInput), uintptr(inputHandle), uintptr(actionSetHandle), uintptr(analogActionHandle), uintptr(unsafe.Pointer(originsOut)))
		__r0 := int32(_r0)
		runtime.KeepAlive(originsOut)
		return __r0
	}
	iSteamInput_GetStringForActionOrigin = func(steamInput uintptr, eOrigin EInputActionOrigin) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetStringForActionOrigin, uintptr(steamInput), uintptr(eOrigin))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamInput_GetStringForAnalogActionName = func(steamInput uintptr, actionHandle InputAnalogActionHandle) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetStringForAnalogActionName, uintptr(steamInput), uintptr(actionHandle))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamInput_StopAnalogActionMomentum = func(steamInput uintptr, inputHandle InputHandle_t, action InputAnalogActionHandle) {
		purego.SyscallN(addr_iSteamInput_StopAnalogActionMomentum, uintptr(steamInput), uintptr(inputHandle), uintptr(action))
	}
	iSteamInput_GetMotionData = func(steamInput uintptr, inputHandle InputHandle_t) uintptr {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetMotionData, uintptr(steamInput), uintptr(inputHandle))
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamInput_TriggerVibration = func(steamInput uintptr, inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16) {
		purego.SyscallN(addr_iSteamInput_TriggerVibration, uintptr(steamInput), uintptr(inputHandle), uintptr(leftSpeed), uintptr(rightSpeed))
	}
	iSteamInput_TriggerVibrationExtended = func(steamInput uintptr, inputHandle InputHandle_t, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16) {
		purego.SyscallN(addr_iSteamInput_TriggerVibrationExtended, uintptr(steamInput), uintptr(inputHandle), uintptr(leftSpeed), uintptr(rightSpeed), uintptr(leftTriggerSpeed), uintptr(rightTriggerSpeed))
	}
	iSteamInput_TriggerSimpleHapticEvent = func(steamInput uintptr, inputHandle InputHandle_t, eHapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte) {
		purego.SyscallN(addr_iSteamInput_TriggerSimpleHapticEvent, uintptr(steamInput), uintptr(inputHandle), uintptr(eHapticLocation), uintptr(intensity), uintptr(gainDB), uintptr(otherIntensity), uintptr(otherGainDB))
	}
	iSteamInput_SetLEDColor = func(steamInput uintptr, inputHandle InputHandle_t, colorR uint8, colorG uint8, colorB uint8, flags uint) {
		purego.SyscallN(addr_iSteamInput_SetLEDColor, uintptr(steamInput), uintptr(inputHandle), uintptr(colorR), uintptr(colorG), uintptr(colorB), uintptr(flags))
	}
	iSteamInput_Legacy_TriggerHapticPulse = func(steamInput uintptr, inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16) {
		purego.SyscallN(addr_iSteamInput_Legacy_TriggerHapticPulse, uintptr(steamInput), uintptr(inputHandle), uintptr(targetPad), uintptr(durationMicroSec))
	}
	iSteamInput_Legacy_TriggerRepeatedHapticPulse = func(steamInput uintptr, inputHandle InputHandle_t, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint) {
		purego.SyscallN(addr_iSteamInput_Legacy_TriggerRepeatedHapticPulse, uintptr(steamInput), uintptr(inputHandle), uintptr(targetPad), uintptr(durationMicroSec), uintptr(offMicroSec), uintptr(repeat), uintptr(flags))
	}
	iSteamInput_ShowBindingPanel = func(steamInput uintptr, inputHandle InputHandle_t) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_ShowBindingPanel, uintptr(steamInput), uintptr(inputHandle))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInput_GetInputTypeForHandle = func(steamInput uintptr, inputHandle InputHandle_t) ESteamInputType {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetInputTypeForHandle, uintptr(steamInput), uintptr(inputHandle))
		__r0 := ESteamInputType(_r0)
		return __r0
	}
	iSteamInput_GetControllerForGamepadIndex = func(steamInput uintptr, index int32) InputHandle_t {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetControllerForGamepadIndex, uintptr(steamInput), uintptr(index))
		__r0 := InputHandle_t(_r0)
		return __r0
	}
	iSteamInput_GetGamepadIndexForController = func(steamInput uintptr, inputHandle InputHandle_t) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetGamepadIndexForController, uintptr(steamInput), uintptr(inputHandle))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamInput_GetStringForXboxOrigin = func(steamInput uintptr, eOrigin EXboxOrigin) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetStringForXboxOrigin, uintptr(steamInput), uintptr(eOrigin))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamInput_GetActionOriginFromXboxOrigin = func(steamInput uintptr, inputHandle InputHandle_t, eOrigin EXboxOrigin) EInputActionOrigin {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetActionOriginFromXboxOrigin, uintptr(steamInput), uintptr(inputHandle), uintptr(eOrigin))
		__r0 := EInputActionOrigin(_r0)
		return __r0
	}
	iSteamInput_TranslateActionOrigin = func(steamInput uintptr, destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_TranslateActionOrigin, uintptr(steamInput), uintptr(destinationInputType), uintptr(sourceOrigin))
		__r0 := EInputActionOrigin(_r0)
		return __r0
	}
	iSteamInput_GetDeviceBindingRevision = func(steamInput uintptr, inputHandle InputHandle_t, major *int32, minor *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetDeviceBindingRevision, uintptr(steamInput), uintptr(inputHandle), uintptr(unsafe.Pointer(major)), uintptr(unsafe.Pointer(minor)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(major)
		runtime.KeepAlive(minor)
		return __r0
	}
	iSteamInput_GetRemotePlaySessionID = func(steamInput uintptr, inputHandle InputHandle_t) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetRemotePlaySessionID, uintptr(steamInput), uintptr(inputHandle))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamInput_GetSessionInputConfigurationSettings = func(steamInput uintptr) uint16 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInput_GetSessionInputConfigurationSettings, uintptr(steamInput))
		__r0 := uint16(_r0)
		return __r0
	}
	iSteamInput_SetDualSenseTriggerEffect = func(steamInput uintptr, inputHandle InputHandle_t, param *ScePadTriggerEffectParam) {
		purego.SyscallN(addr_iSteamInput_SetDualSenseTriggerEffect, uintptr(steamInput), uintptr(inputHandle), uintptr(unsafe.Pointer(param)))
		runtime.KeepAlive(param)
	}
}

var (
	addr_steamInventory_get                                     uintptr
	addr_iSteamInventory_GetResultStatus                        uintptr
	addr_iSteamInventory_GetResultItems                         uintptr
	addr_iSteamInventory_GetResultItemProperty                  uintptr
	addr_iSteamInventory_GetResultTimestamp                     uintptr
	addr_iSteamInventory_CheckResultSteamID                     uintptr
	addr_iSteamInventory_DestroyResult                          uintptr
	addr_iSteamInventory_GetAllItems                            uintptr
	addr_iSteamInventory_GetItemsByID                           uintptr
	addr_iSteamInventory_SerializeResult                        uintptr
	addr_iSteamInventory_DeserializeResult                      uintptr
	addr_iSteamInventory_GenerateItems                          uintptr
	addr_iSteamInventory_GrantPromoItems                        uintptr
	addr_iSteamInventory_AddPromoItem                           uintptr
	addr_iSteamInventory_AddPromoItems                          uintptr
	addr_iSteamInventory_ConsumeItem                            uintptr
	addr_iSteamInventory_ExchangeItems                          uintptr
	addr_iSteamInventory_TransferItemQuantity                   uintptr
	addr_iSteamInventory_TriggerItemDrop                        uintptr
	addr_iSteamInventory_LoadItemDefinitions                    uintptr
	addr_iSteamInventory_GetItemDefinitionIDs                   uintptr
	addr_iSteamInventory_GetItemDefinitionProperty              uintptr
	addr_iSteamInventory_RequestEligiblePromoItemDefinitionsIDs uintptr
	addr_iSteamInventory_GetEligiblePromoItemDefinitionIDs      uintptr
	addr_iSteamInventory_StartPurchase                          uintptr
	addr_iSteamInventory_RequestPrices                          uintptr
	addr_iSteamInventory_GetNumItemsWithPrices                  uintptr
	addr_iSteamInventory_GetItemsWithPrices                     uintptr
	addr_iSteamInventory_GetItemPrice                           uintptr
	addr_iSteamInventory_StartUpdateProperties                  uintptr
	addr_iSteamInventory_RemoveProperty                         uintptr
	addr_iSteamInventory_SetPropertyString                      uintptr
	addr_iSteamInventory_SetPropertyBool                        uintptr
	addr_iSteamInventory_SetPropertyInt64                       uintptr
	addr_iSteamInventory_SetPropertyFloat                       uintptr
	addr_iSteamInventory_SubmitUpdateProperties                 uintptr
	addr_iSteamInventory_InspectItem                            uintptr
)

func initInventory() {
	var err error
	addr_steamInventory_get, err = openSymbol(steamAPILib, flatAPI_SteamInventory)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamInventory)
	}
	addr_iSteamInventory_GetResultStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetResultStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetResultStatus")
	}
	addr_iSteamInventory_GetResultItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetResultItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetResultItems")
	}
	addr_iSteamInventory_GetResultItemProperty, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetResultItemProperty")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetResultItemProperty")
	}
	addr_iSteamInventory_GetResultTimestamp, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetResultTimestamp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetResultTimestamp")
	}
	addr_iSteamInventory_CheckResultSteamID, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_CheckResultSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_CheckResultSteamID")
	}
	addr_iSteamInventory_DestroyResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_DestroyResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_DestroyResult")
	}
	addr_iSteamInventory_GetAllItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetAllItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetAllItems")
	}
	addr_iSteamInventory_GetItemsByID, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetItemsByID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetItemsByID")
	}
	addr_iSteamInventory_SerializeResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SerializeResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SerializeResult")
	}
	addr_iSteamInventory_DeserializeResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_DeserializeResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_DeserializeResult")
	}
	addr_iSteamInventory_GenerateItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GenerateItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GenerateItems")
	}
	addr_iSteamInventory_GrantPromoItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GrantPromoItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GrantPromoItems")
	}
	addr_iSteamInventory_AddPromoItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_AddPromoItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_AddPromoItem")
	}
	addr_iSteamInventory_AddPromoItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_AddPromoItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_AddPromoItems")
	}
	addr_iSteamInventory_ConsumeItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_ConsumeItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_ConsumeItem")
	}
	addr_iSteamInventory_ExchangeItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_ExchangeItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_ExchangeItems")
	}
	addr_iSteamInventory_TransferItemQuantity, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_TransferItemQuantity")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_TransferItemQuantity")
	}
	addr_iSteamInventory_TriggerItemDrop, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_TriggerItemDrop")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_TriggerItemDrop")
	}
	addr_iSteamInventory_LoadItemDefinitions, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_LoadItemDefinitions")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_LoadItemDefinitions")
	}
	addr_iSteamInventory_GetItemDefinitionIDs, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetItemDefinitionIDs")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetItemDefinitionIDs")
	}
	addr_iSteamInventory_GetItemDefinitionProperty, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetItemDefinitionProperty")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetItemDefinitionProperty")
	}
	addr_iSteamInventory_RequestEligiblePromoItemDefinitionsIDs, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs")
	}
	addr_iSteamInventory_GetEligiblePromoItemDefinitionIDs, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs")
	}
	addr_iSteamInventory_StartPurchase, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_StartPurchase")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_StartPurchase")
	}
	addr_iSteamInventory_RequestPrices, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_RequestPrices")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_RequestPrices")
	}
	addr_iSteamInventory_GetNumItemsWithPrices, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetNumItemsWithPrices")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetNumItemsWithPrices")
	}
	addr_iSteamInventory_GetItemsWithPrices, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetItemsWithPrices")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetItemsWithPrices")
	}
	addr_iSteamInventory_GetItemPrice, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_GetItemPrice")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_GetItemPrice")
	}
	addr_iSteamInventory_StartUpdateProperties, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_StartUpdateProperties")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_StartUpdateProperties")
	}
	addr_iSteamInventory_RemoveProperty, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_RemoveProperty")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_RemoveProperty")
	}
	addr_iSteamInventory_SetPropertyString, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SetPropertyString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SetPropertyString")
	}
	addr_iSteamInventory_SetPropertyBool, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SetPropertyBool")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SetPropertyBool")
	}
	addr_iSteamInventory_SetPropertyInt64, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SetPropertyInt64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SetPropertyInt64")
	}
	addr_iSteamInventory_SetPropertyFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SetPropertyFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SetPropertyFloat")
	}
	addr_iSteamInventory_SubmitUpdateProperties, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_SubmitUpdateProperties")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_SubmitUpdateProperties")
	}
	addr_iSteamInventory_InspectItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamInventory_InspectItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamInventory_InspectItem")
	}

	steamInventory_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamInventory_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamInventory_GetResultStatus = func(steamInventory uintptr, resultHandle SteamInventoryResult) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetResultStatus, uintptr(steamInventory), uintptr(resultHandle))
		__r0 := EResult(_r0)
		return __r0
	}
	iSteamInventory_GetResultItems = func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutItemsArray []SteamItemDetails, punOutItemsArraySize *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetResultItems, uintptr(steamInventory), uintptr(resultHandle), uintptr(unsafe.Pointer(unsafe.SliceData(pOutItemsArray))), uintptr(unsafe.Pointer(punOutItemsArraySize)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pOutItemsArray)
		runtime.KeepAlive(punOutItemsArraySize)
		return __r0
	}
	iSteamInventory_GetResultItemProperty = func(steamInventory uintptr, resultHandle SteamInventoryResult, unItemIndex uint32, pchPropertyName string, pchValueBuffer []byte, punValueBufferSizeOut *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetResultItemProperty, uintptr(steamInventory), uintptr(resultHandle), uintptr(unItemIndex), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))), uintptr(unsafe.Pointer(unsafe.SliceData(pchValueBuffer))), uintptr(unsafe.Pointer(punValueBufferSizeOut)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		runtime.KeepAlive(pchValueBuffer)
		runtime.KeepAlive(punValueBufferSizeOut)
		return __r0
	}
	iSteamInventory_GetResultTimestamp = func(steamInventory uintptr, resultHandle SteamInventoryResult) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetResultTimestamp, uintptr(steamInventory), uintptr(resultHandle))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamInventory_CheckResultSteamID = func(steamInventory uintptr, resultHandle SteamInventoryResult, steamIDExpected Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_CheckResultSteamID, uintptr(steamInventory), uintptr(resultHandle), uintptr(steamIDExpected))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInventory_DestroyResult = func(steamInventory uintptr, resultHandle SteamInventoryResult) {
		purego.SyscallN(addr_iSteamInventory_DestroyResult, uintptr(steamInventory), uintptr(resultHandle))
	}
	iSteamInventory_GetAllItems = func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetAllItems, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_GetItemsByID = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pInstanceIDs []SteamItemInstanceID, unCountInstanceIDs uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetItemsByID, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(unsafe.Pointer(unsafe.SliceData(pInstanceIDs))), uintptr(unCountInstanceIDs))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		runtime.KeepAlive(pInstanceIDs)
		return __r0
	}
	iSteamInventory_SerializeResult = func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutBuffer []byte, punOutBufferSize *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_SerializeResult, uintptr(steamInventory), uintptr(resultHandle), uintptr(unsafe.Pointer(unsafe.SliceData(pOutBuffer))), uintptr(unsafe.Pointer(punOutBufferSize)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pOutBuffer)
		runtime.KeepAlive(punOutBufferSize)
		return __r0
	}
	iSteamInventory_DeserializeResult = func(steamInventory uintptr, pOutResultHandle *SteamInventoryResult, pBuffer []byte, unBufferSize uint32, bRESERVEDMUSTBEFALSE bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_DeserializeResult, uintptr(steamInventory), uintptr(unsafe.Pointer(pOutResultHandle)), uintptr(unsafe.Pointer(unsafe.SliceData(pBuffer))), uintptr(unBufferSize), boolToUintptr(bRESERVEDMUSTBEFALSE))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pOutResultHandle)
		runtime.KeepAlive(pBuffer)
		return __r0
	}
	iSteamInventory_GenerateItems = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GenerateItems, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayItemDefs))), uintptr(unsafe.Pointer(unsafe.SliceData(punArrayQuantity))), uintptr(unArrayLength))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		runtime.KeepAlive(pArrayItemDefs)
		runtime.KeepAlive(punArrayQuantity)
		return __r0
	}
	iSteamInventory_GrantPromoItems = func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GrantPromoItems, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_AddPromoItem = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemDef SteamItemDef) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_AddPromoItem, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(itemDef))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_AddPromoItems = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, unArrayLength uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_AddPromoItems, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayItemDefs))), uintptr(unArrayLength))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		runtime.KeepAlive(pArrayItemDefs)
		return __r0
	}
	iSteamInventory_ConsumeItem = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemConsume SteamItemInstanceID, unQuantity uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_ConsumeItem, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(itemConsume), uintptr(unQuantity))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_ExchangeItems = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayGenerate []SteamItemDef, punArrayGenerateQuantity []uint32, unArrayGenerateLength uint32, pArrayDestroy []SteamItemInstanceID, punArrayDestroyQuantity []uint32, unArrayDestroyLength uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_ExchangeItems, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayGenerate))), uintptr(unsafe.Pointer(unsafe.SliceData(punArrayGenerateQuantity))), uintptr(unArrayGenerateLength), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayDestroy))), uintptr(unsafe.Pointer(unsafe.SliceData(punArrayDestroyQuantity))), uintptr(unArrayDestroyLength))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		runtime.KeepAlive(pArrayGenerate)
		runtime.KeepAlive(punArrayGenerateQuantity)
		runtime.KeepAlive(pArrayDestroy)
		runtime.KeepAlive(punArrayDestroyQuantity)
		return __r0
	}
	iSteamInventory_TransferItemQuantity = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemIdSource SteamItemInstanceID, unQuantity uint32, itemIdDest SteamItemInstanceID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_TransferItemQuantity, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(itemIdSource), uintptr(unQuantity), uintptr(itemIdDest))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_TriggerItemDrop = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, dropListDefinition SteamItemDef) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_TriggerItemDrop, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(dropListDefinition))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_LoadItemDefinitions = func(steamInventory uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_LoadItemDefinitions, uintptr(steamInventory))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamInventory_GetItemDefinitionIDs = func(steamInventory uintptr, pItemDefIDs []SteamItemDef, punItemDefIDsArraySize *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetItemDefinitionIDs, uintptr(steamInventory), uintptr(unsafe.Pointer(unsafe.SliceData(pItemDefIDs))), uintptr(unsafe.Pointer(punItemDefIDsArraySize)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pItemDefIDs)
		runtime.KeepAlive(punItemDefIDsArraySize)
		return __r0
	}
	iSteamInventory_GetItemDefinitionProperty = func(steamInventory uintptr, iDefinition SteamItemDef, pchPropertyName string, pchValueBuffer []byte, punValueBufferSizeOut *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetItemDefinitionProperty, uintptr(steamInventory), uintptr(iDefinition), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))), uintptr(unsafe.Pointer(unsafe.SliceData(pchValueBuffer))), uintptr(unsafe.Pointer(punValueBufferSizeOut)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		runtime.KeepAlive(pchValueBuffer)
		runtime.KeepAlive(punValueBufferSizeOut)
		return __r0
	}
	iSteamInventory_RequestEligiblePromoItemDefinitionsIDs = func(steamInventory uintptr, steamID Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_RequestEligiblePromoItemDefinitionsIDs, uintptr(steamInventory), uintptr(steamID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamInventory_GetEligiblePromoItemDefinitionIDs = func(steamInventory uintptr, steamID Uint64SteamID, pItemDefIDs []SteamItemDef, punItemDefIDsArraySize *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetEligiblePromoItemDefinitionIDs, uintptr(steamInventory), uintptr(steamID), uintptr(unsafe.Pointer(unsafe.SliceData(pItemDefIDs))), uintptr(unsafe.Pointer(punItemDefIDsArraySize)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pItemDefIDs)
		runtime.KeepAlive(punItemDefIDsArraySize)
		return __r0
	}
	iSteamInventory_StartPurchase = func(steamInventory uintptr, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_StartPurchase, uintptr(steamInventory), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayItemDefs))), uintptr(unsafe.Pointer(unsafe.SliceData(punArrayQuantity))), uintptr(unArrayLength))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pArrayItemDefs)
		runtime.KeepAlive(punArrayQuantity)
		return __r0
	}
	iSteamInventory_RequestPrices = func(steamInventory uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_RequestPrices, uintptr(steamInventory))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamInventory_GetNumItemsWithPrices = func(steamInventory uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetNumItemsWithPrices, uintptr(steamInventory))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamInventory_GetItemsWithPrices = func(steamInventory uintptr, pArrayItemDefs []SteamItemDef, pCurrentPrices []uint64, pBasePrices []uint64, unArrayLength uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetItemsWithPrices, uintptr(steamInventory), uintptr(unsafe.Pointer(unsafe.SliceData(pArrayItemDefs))), uintptr(unsafe.Pointer(unsafe.SliceData(pCurrentPrices))), uintptr(unsafe.Pointer(unsafe.SliceData(pBasePrices))), uintptr(unArrayLength))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pArrayItemDefs)
		runtime.KeepAlive(pCurrentPrices)
		runtime.KeepAlive(pBasePrices)
		return __r0
	}
	iSteamInventory_GetItemPrice = func(steamInventory uintptr, iDefinition SteamItemDef, pCurrentPrice *uint64, pBasePrice *uint64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_GetItemPrice, uintptr(steamInventory), uintptr(iDefinition), uintptr(unsafe.Pointer(pCurrentPrice)), uintptr(unsafe.Pointer(pBasePrice)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pCurrentPrice)
		runtime.KeepAlive(pBasePrice)
		return __r0
	}
	iSteamInventory_StartUpdateProperties = func(steamInventory uintptr) SteamInventoryUpdateHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_StartUpdateProperties, uintptr(steamInventory))
		__r0 := SteamInventoryUpdateHandle(_r0)
		return __r0
	}
	iSteamInventory_RemoveProperty = func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_RemoveProperty, uintptr(steamInventory), uintptr(handle), uintptr(nItemID), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		return __r0
	}
	iSteamInventory_SetPropertyString = func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, pchPropertyValue string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_SetPropertyString, uintptr(steamInventory), uintptr(handle), uintptr(nItemID), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyValue))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		runtime.KeepAlive(pchPropertyValue)
		return __r0
	}
	iSteamInventory_SetPropertyBool = func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, bValue bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_SetPropertyBool, uintptr(steamInventory), uintptr(handle), uintptr(nItemID), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))), boolToUintptr(bValue))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		return __r0
	}
	iSteamInventory_SetPropertyInt64 = func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, nValue int64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_SetPropertyInt64, uintptr(steamInventory), uintptr(handle), uintptr(nItemID), uintptr(unsafe.Pointer(bytePtrFromString(pchPropertyName))), uintptr(nValue))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPropertyName)
		return __r0
	}
	iSteamInventory_SubmitUpdateProperties = func(steamInventory uintptr, handle SteamInventoryUpdateHandle, pResultHandle *SteamInventoryResult) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_SubmitUpdateProperties, uintptr(steamInventory), uintptr(handle), uintptr(unsafe.Pointer(pResultHandle)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		return __r0
	}
	iSteamInventory_InspectItem = func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pchItemToken string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamInventory_InspectItem, uintptr(steamInventory), uintptr(unsafe.Pointer(pResultHandle)), uintptr(unsafe.Pointer(bytePtrFromString(pchItemToken))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pResultHandle)
		runtime.KeepAlive(pchItemToken)
		return __r0
	}

	purego.RegisterLibFunc(&iSteamInventory_SetPropertyFloat, steamAPILib, flatAPI_ISteamInventory_SetPropertyFloat)

}

var (
	addr_steamMatchmaking_get                                         uintptr
	addr_iSteamMatchmaking_GetFavoriteGameCount                       uintptr
	addr_iSteamMatchmaking_GetFavoriteGame                            uintptr
	addr_iSteamMatchmaking_AddFavoriteGame                            uintptr
	addr_iSteamMatchmaking_RemoveFavoriteGame                         uintptr
	addr_iSteamMatchmaking_RequestLobbyList                           uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListStringFilter            uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListNumericalFilter         uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListNearValueFilter         uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable    uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListDistanceFilter          uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListResultCountFilter       uintptr
	addr_iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter uintptr
	addr_iSteamMatchmaking_GetLobbyByIndex                            uintptr
	addr_iSteamMatchmaking_CreateLobby                                uintptr
	addr_iSteamMatchmaking_JoinLobby                                  uintptr
	addr_iSteamMatchmaking_LeaveLobby                                 uintptr
	addr_iSteamMatchmaking_InviteUserToLobby                          uintptr
	addr_iSteamMatchmaking_GetNumLobbyMembers                         uintptr
	addr_iSteamMatchmaking_GetLobbyMemberByIndex                      uintptr
	addr_iSteamMatchmaking_GetLobbyData                               uintptr
	addr_iSteamMatchmaking_SetLobbyData                               uintptr
	addr_iSteamMatchmaking_GetLobbyDataCount                          uintptr
	addr_iSteamMatchmaking_GetLobbyDataByIndex                        uintptr
	addr_iSteamMatchmaking_DeleteLobbyData                            uintptr
	addr_iSteamMatchmaking_GetLobbyMemberData                         uintptr
	addr_iSteamMatchmaking_SetLobbyMemberData                         uintptr
	addr_iSteamMatchmaking_SendLobbyChatMsg                           uintptr
	addr_iSteamMatchmaking_GetLobbyChatEntry                          uintptr
	addr_iSteamMatchmaking_RequestLobbyData                           uintptr
	addr_iSteamMatchmaking_SetLobbyGameServer                         uintptr
	addr_iSteamMatchmaking_GetLobbyGameServer                         uintptr
	addr_iSteamMatchmaking_SetLobbyMemberLimit                        uintptr
	addr_iSteamMatchmaking_GetLobbyMemberLimit                        uintptr
	addr_iSteamMatchmaking_SetLobbyType                               uintptr
	addr_iSteamMatchmaking_SetLobbyJoinable                           uintptr
	addr_iSteamMatchmaking_GetLobbyOwner                              uintptr
	addr_iSteamMatchmaking_SetLobbyOwner                              uintptr
	addr_iSteamMatchmaking_SetLinkedLobby                             uintptr
)

func initMatchmaking() {
	var err error
	addr_steamMatchmaking_get, err = openSymbol(steamAPILib, flatAPI_SteamMatchmaking)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamMatchmaking)
	}
	addr_iSteamMatchmaking_GetFavoriteGameCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetFavoriteGameCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetFavoriteGameCount")
	}
	addr_iSteamMatchmaking_GetFavoriteGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetFavoriteGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetFavoriteGame")
	}
	addr_iSteamMatchmaking_AddFavoriteGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddFavoriteGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddFavoriteGame")
	}
	addr_iSteamMatchmaking_RemoveFavoriteGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_RemoveFavoriteGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_RemoveFavoriteGame")
	}
	addr_iSteamMatchmaking_RequestLobbyList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_RequestLobbyList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_RequestLobbyList")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListStringFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListNumericalFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListNearValueFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListDistanceFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListResultCountFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter")
	}
	addr_iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter")
	}
	addr_iSteamMatchmaking_GetLobbyByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyByIndex")
	}
	addr_iSteamMatchmaking_CreateLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_CreateLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_CreateLobby")
	}
	addr_iSteamMatchmaking_JoinLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_JoinLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_JoinLobby")
	}
	addr_iSteamMatchmaking_LeaveLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_LeaveLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_LeaveLobby")
	}
	addr_iSteamMatchmaking_InviteUserToLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_InviteUserToLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_InviteUserToLobby")
	}
	addr_iSteamMatchmaking_GetNumLobbyMembers, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetNumLobbyMembers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetNumLobbyMembers")
	}
	addr_iSteamMatchmaking_GetLobbyMemberByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyMemberByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyMemberByIndex")
	}
	addr_iSteamMatchmaking_GetLobbyData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyData")
	}
	addr_iSteamMatchmaking_SetLobbyData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyData")
	}
	addr_iSteamMatchmaking_GetLobbyDataCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyDataCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyDataCount")
	}
	addr_iSteamMatchmaking_GetLobbyDataByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyDataByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyDataByIndex")
	}
	addr_iSteamMatchmaking_DeleteLobbyData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_DeleteLobbyData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_DeleteLobbyData")
	}
	addr_iSteamMatchmaking_GetLobbyMemberData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyMemberData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyMemberData")
	}
	addr_iSteamMatchmaking_SetLobbyMemberData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyMemberData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyMemberData")
	}
	addr_iSteamMatchmaking_SendLobbyChatMsg, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SendLobbyChatMsg")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SendLobbyChatMsg")
	}
	addr_iSteamMatchmaking_GetLobbyChatEntry, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyChatEntry")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyChatEntry")
	}
	addr_iSteamMatchmaking_RequestLobbyData, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_RequestLobbyData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_RequestLobbyData")
	}
	addr_iSteamMatchmaking_SetLobbyGameServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyGameServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyGameServer")
	}
	addr_iSteamMatchmaking_GetLobbyGameServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyGameServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyGameServer")
	}
	addr_iSteamMatchmaking_SetLobbyMemberLimit, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyMemberLimit")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyMemberLimit")
	}
	addr_iSteamMatchmaking_GetLobbyMemberLimit, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyMemberLimit")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyMemberLimit")
	}
	addr_iSteamMatchmaking_SetLobbyType, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyType")
	}
	addr_iSteamMatchmaking_SetLobbyJoinable, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyJoinable")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyJoinable")
	}
	addr_iSteamMatchmaking_GetLobbyOwner, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_GetLobbyOwner")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_GetLobbyOwner")
	}
	addr_iSteamMatchmaking_SetLobbyOwner, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLobbyOwner")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLobbyOwner")
	}
	addr_iSteamMatchmaking_SetLinkedLobby, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmaking_SetLinkedLobby")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmaking_SetLinkedLobby")
	}

	steamMatchmaking_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamMatchmaking_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_GetFavoriteGameCount = func(steamMatchmaking uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetFavoriteGameCount, uintptr(steamMatchmaking))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmaking_GetFavoriteGame = func(steamMatchmaking uintptr, iGame int32, pnAppID *AppId_t, pnIP *uint32, pnConnPort *uint16, pnQueryPort *uint16, punFlags *uint32, pRTime32LastPlayedOnServer *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetFavoriteGame, uintptr(steamMatchmaking), uintptr(iGame), uintptr(unsafe.Pointer(pnAppID)), uintptr(unsafe.Pointer(pnIP)), uintptr(unsafe.Pointer(pnConnPort)), uintptr(unsafe.Pointer(pnQueryPort)), uintptr(unsafe.Pointer(punFlags)), uintptr(unsafe.Pointer(pRTime32LastPlayedOnServer)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnAppID)
		runtime.KeepAlive(pnIP)
		runtime.KeepAlive(pnConnPort)
		runtime.KeepAlive(pnQueryPort)
		runtime.KeepAlive(punFlags)
		runtime.KeepAlive(pRTime32LastPlayedOnServer)
		return __r0
	}
	iSteamMatchmaking_AddFavoriteGame = func(steamMatchmaking uintptr, nAppID AppId_t, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32, rTime32LastPlayedOnServer uint32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_AddFavoriteGame, uintptr(steamMatchmaking), uintptr(nAppID), uintptr(nIP), uintptr(nConnPort), uintptr(nQueryPort), uintptr(unFlags), uintptr(rTime32LastPlayedOnServer))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmaking_RemoveFavoriteGame = func(steamMatchmaking uintptr, nAppID AppId_t, nIP uint32, nConnPort uint16, nQueryPort uint16, unFlags uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_RemoveFavoriteGame, uintptr(steamMatchmaking), uintptr(nAppID), uintptr(nIP), uintptr(nConnPort), uintptr(nQueryPort), uintptr(unFlags))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_RequestLobbyList = func(steamMatchmaking uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_RequestLobbyList, uintptr(steamMatchmaking))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamMatchmaking_AddRequestLobbyListStringFilter = func(steamMatchmaking uintptr, pchKeyToMatch string, pchValueToMatch string, eComparisonType ELobbyComparison) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListStringFilter, uintptr(steamMatchmaking), uintptr(unsafe.Pointer(bytePtrFromString(pchKeyToMatch))), uintptr(unsafe.Pointer(bytePtrFromString(pchValueToMatch))), uintptr(eComparisonType))
		runtime.KeepAlive(pchKeyToMatch)
		runtime.KeepAlive(pchValueToMatch)
	}
	iSteamMatchmaking_AddRequestLobbyListNumericalFilter = func(steamMatchmaking uintptr, pchKeyToMatch string, nValueToMatch int32, eComparisonType ELobbyComparison) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListNumericalFilter, uintptr(steamMatchmaking), uintptr(unsafe.Pointer(bytePtrFromString(pchKeyToMatch))), uintptr(nValueToMatch), uintptr(eComparisonType))
		runtime.KeepAlive(pchKeyToMatch)
	}
	iSteamMatchmaking_AddRequestLobbyListNearValueFilter = func(steamMatchmaking uintptr, pchKeyToMatch string, nValueToBeCloseTo int32) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListNearValueFilter, uintptr(steamMatchmaking), uintptr(unsafe.Pointer(bytePtrFromString(pchKeyToMatch))), uintptr(nValueToBeCloseTo))
		runtime.KeepAlive(pchKeyToMatch)
	}
	iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable = func(steamMatchmaking uintptr, nSlotsAvailable int32) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable, uintptr(steamMatchmaking), uintptr(nSlotsAvailable))
	}
	iSteamMatchmaking_AddRequestLobbyListDistanceFilter = func(steamMatchmaking uintptr, eLobbyDistanceFilter ELobbyDistanceFilter) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListDistanceFilter, uintptr(steamMatchmaking), uintptr(eLobbyDistanceFilter))
	}
	iSteamMatchmaking_AddRequestLobbyListResultCountFilter = func(steamMatchmaking uintptr, cMaxResults int32) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListResultCountFilter, uintptr(steamMatchmaking), uintptr(cMaxResults))
	}
	iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) {
		purego.SyscallN(addr_iSteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter, uintptr(steamMatchmaking), uintptr(steamIDLobby))
	}
	iSteamMatchmaking_GetLobbyByIndex = func(steamMatchmaking uintptr, iLobby int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyByIndex, uintptr(steamMatchmaking), uintptr(iLobby))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamMatchmaking_CreateLobby = func(steamMatchmaking uintptr, eLobbyType ELobbyType, cMaxMembers int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_CreateLobby, uintptr(steamMatchmaking), uintptr(eLobbyType), uintptr(cMaxMembers))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamMatchmaking_JoinLobby = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_JoinLobby, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamMatchmaking_LeaveLobby = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) {
		purego.SyscallN(addr_iSteamMatchmaking_LeaveLobby, uintptr(steamMatchmaking), uintptr(steamIDLobby))
	}
	iSteamMatchmaking_InviteUserToLobby = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDInvitee Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_InviteUserToLobby, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(steamIDInvitee))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_GetNumLobbyMembers = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetNumLobbyMembers, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmaking_GetLobbyMemberByIndex = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iMember int32) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyMemberByIndex, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(iMember))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLobbyData = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string, pchValue string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLobbyData, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))), uintptr(unsafe.Pointer(bytePtrFromString(pchValue))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamMatchmaking_GetLobbyDataCount = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyDataCount, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmaking_GetLobbyDataByIndex = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iLobbyData int32, pchKey []byte, cchKeyBufferSize int32, pchValue []byte, cchValueBufferSize int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyDataByIndex, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(iLobbyData), uintptr(unsafe.Pointer(unsafe.SliceData(pchKey))), uintptr(cchKeyBufferSize), uintptr(unsafe.Pointer(unsafe.SliceData(pchValue))), uintptr(cchValueBufferSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamMatchmaking_DeleteLobbyData = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_DeleteLobbyData, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		return __r0
	}
	iSteamMatchmaking_SetLobbyMemberData = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pchKey string, pchValue string) {
		purego.SyscallN(addr_iSteamMatchmaking_SetLobbyMemberData, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))), uintptr(unsafe.Pointer(bytePtrFromString(pchValue))))
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
	}
	iSteamMatchmaking_SendLobbyChatMsg = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, pvMsgBody []byte, cubMsgBody int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SendLobbyChatMsg, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unsafe.Pointer(unsafe.SliceData(pvMsgBody))), uintptr(cubMsgBody))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pvMsgBody)
		return __r0
	}
	iSteamMatchmaking_GetLobbyChatEntry = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, iChatID int32, pSteamIDUser *CSteamID, pvData []byte, cubData int32, peChatEntryType *EChatEntryType) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyChatEntry, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(iChatID), uintptr(unsafe.Pointer(pSteamIDUser)), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubData), uintptr(unsafe.Pointer(peChatEntryType)))
		__r0 := int32(_r0)
		runtime.KeepAlive(pSteamIDUser)
		runtime.KeepAlive(pvData)
		runtime.KeepAlive(peChatEntryType)
		return __r0
	}
	iSteamMatchmaking_RequestLobbyData = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_RequestLobbyData, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLobbyGameServer = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, unGameServerIP uint32, unGameServerPort uint16, steamIDGameServer Uint64SteamID) {
		purego.SyscallN(addr_iSteamMatchmaking_SetLobbyGameServer, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unGameServerIP), uintptr(unGameServerPort), uintptr(steamIDGameServer))
	}
	iSteamMatchmaking_GetLobbyGameServer = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, punGameServerIP *uint32, punGameServerPort *uint16, psteamIDGameServer *CSteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyGameServer, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(unsafe.Pointer(punGameServerIP)), uintptr(unsafe.Pointer(punGameServerPort)), uintptr(unsafe.Pointer(psteamIDGameServer)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(punGameServerIP)
		runtime.KeepAlive(punGameServerPort)
		runtime.KeepAlive(psteamIDGameServer)
		return __r0
	}
	iSteamMatchmaking_SetLobbyMemberLimit = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, cMaxMembers int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLobbyMemberLimit, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(cMaxMembers))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_GetLobbyMemberLimit = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyMemberLimit, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLobbyType = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, eLobbyType ELobbyType) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLobbyType, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(eLobbyType))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLobbyJoinable = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, bLobbyJoinable bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLobbyJoinable, uintptr(steamMatchmaking), uintptr(steamIDLobby), boolToUintptr(bLobbyJoinable))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_GetLobbyOwner = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_GetLobbyOwner, uintptr(steamMatchmaking), uintptr(steamIDLobby))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLobbyOwner = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDNewOwner Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLobbyOwner, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(steamIDNewOwner))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmaking_SetLinkedLobby = func(steamMatchmaking uintptr, steamIDLobby Uint64SteamID, steamIDLobbyDependent Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmaking_SetLinkedLobby, uintptr(steamMatchmaking), uintptr(steamIDLobby), uintptr(steamIDLobbyDependent))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}

var (
	addr_matchMakingKeyValuePair_t_Construct                       uintptr
	addr_servernetadr_t_Construct                                  uintptr
	addr_servernetadr_t_Init                                       uintptr
	addr_servernetadr_t_GetQueryPort                               uintptr
	addr_servernetadr_t_SetQueryPort                               uintptr
	addr_servernetadr_t_GetConnectionPort                          uintptr
	addr_servernetadr_t_SetConnectionPort                          uintptr
	addr_servernetadr_t_GetIP                                      uintptr
	addr_servernetadr_t_SetIP                                      uintptr
	addr_servernetadr_t_GetConnectionAddressString                 uintptr
	addr_servernetadr_t_GetQueryAddressString                      uintptr
	addr_servernetadr_t_IsLessThan                                 uintptr
	addr_servernetadr_t_Assign                                     uintptr
	addr_gameserveritem_t_Construct                                uintptr
	addr_gameserveritem_t_GetName                                  uintptr
	addr_gameserveritem_t_SetName                                  uintptr
	addr_steamMatchmakingServers_get                               uintptr
	addr_iSteamMatchmakingServers_RequestInternetServerList        uintptr
	addr_iSteamMatchmakingServers_RequestLANServerList             uintptr
	addr_iSteamMatchmakingServers_RequestFriendsServerList         uintptr
	addr_iSteamMatchmakingServers_RequestFavoritesServerList       uintptr
	addr_iSteamMatchmakingServers_RequestHistoryServerList         uintptr
	addr_iSteamMatchmakingServers_RequestSpectatorServerList       uintptr
	addr_iSteamMatchmakingServers_ReleaseRequest                   uintptr
	addr_iSteamMatchmakingServers_GetServerDetails                 uintptr
	addr_iSteamMatchmakingServers_CancelQuery                      uintptr
	addr_iSteamMatchmakingServers_RefreshQuery                     uintptr
	addr_iSteamMatchmakingServers_IsRefreshing                     uintptr
	addr_iSteamMatchmakingServers_GetServerCount                   uintptr
	addr_iSteamMatchmakingServers_RefreshServer                    uintptr
	addr_iSteamMatchmakingServers_PingServer                       uintptr
	addr_iSteamMatchmakingServers_PlayerDetails                    uintptr
	addr_iSteamMatchmakingServers_ServerRules                      uintptr
	addr_iSteamMatchmakingServers_CancelServerQuery                uintptr
	addr_iSteamMatchmakingServerListResponse_ServerResponded       uintptr
	addr_iSteamMatchmakingServerListResponse_ServerFailedToRespond uintptr
	addr_iSteamMatchmakingServerListResponse_RefreshComplete       uintptr
	addr_iSteamMatchmakingPingResponse_ServerResponded             uintptr
	addr_iSteamMatchmakingPingResponse_ServerFailedToRespond       uintptr
	addr_iSteamMatchmakingPlayersResponse_AddPlayerToList          uintptr
	addr_iSteamMatchmakingPlayersResponse_PlayersFailedToRespond   uintptr
	addr_iSteamMatchmakingPlayersResponse_PlayersRefreshComplete   uintptr
	addr_iSteamMatchmakingRulesResponse_RulesResponded             uintptr
	addr_iSteamMatchmakingRulesResponse_RulesFailedToRespond       uintptr
	addr_iSteamMatchmakingRulesResponse_RulesRefreshComplete       uintptr
)

func initMatchmakingServer() {

	var err error
	addr_matchMakingKeyValuePair_t_Construct, err = openSymbol(steamAPILib, "SteamAPI_MatchMakingKeyValuePair_t_Construct")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_MatchMakingKeyValuePair_t_Construct")
	}
	addr_servernetadr_t_Construct, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_Construct")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_Construct")
	}
	addr_servernetadr_t_Init, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_Init")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_Init")
	}
	addr_servernetadr_t_GetQueryPort, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_GetQueryPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_GetQueryPort")
	}
	addr_servernetadr_t_SetQueryPort, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_SetQueryPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_SetQueryPort")
	}
	addr_servernetadr_t_GetConnectionPort, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_GetConnectionPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_GetConnectionPort")
	}
	addr_servernetadr_t_SetConnectionPort, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_SetConnectionPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_SetConnectionPort")
	}
	addr_servernetadr_t_GetIP, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_GetIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_GetIP")
	}
	addr_servernetadr_t_SetIP, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_SetIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_SetIP")
	}
	addr_servernetadr_t_GetConnectionAddressString, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_GetConnectionAddressString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_GetConnectionAddressString")
	}
	addr_servernetadr_t_GetQueryAddressString, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_GetQueryAddressString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_GetQueryAddressString")
	}
	addr_servernetadr_t_IsLessThan, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_IsLessThan")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_IsLessThan")
	}
	addr_servernetadr_t_Assign, err = openSymbol(steamAPILib, "SteamAPI_servernetadr_t_Assign")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_servernetadr_t_Assign")
	}
	addr_gameserveritem_t_Construct, err = openSymbol(steamAPILib, "SteamAPI_gameserveritem_t_Construct")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_gameserveritem_t_Construct")
	}
	addr_gameserveritem_t_GetName, err = openSymbol(steamAPILib, "SteamAPI_gameserveritem_t_GetName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_gameserveritem_t_GetName")
	}
	addr_gameserveritem_t_SetName, err = openSymbol(steamAPILib, "SteamAPI_gameserveritem_t_SetName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_gameserveritem_t_SetName")
	}
	addr_steamMatchmakingServers_get, err = openSymbol(steamAPILib, flatAPI_SteamMatchmakingServers)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamMatchmakingServers)
	}
	addr_iSteamMatchmakingServers_RequestInternetServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestInternetServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestInternetServerList")
	}
	addr_iSteamMatchmakingServers_RequestLANServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestLANServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestLANServerList")
	}
	addr_iSteamMatchmakingServers_RequestFriendsServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestFriendsServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestFriendsServerList")
	}
	addr_iSteamMatchmakingServers_RequestFavoritesServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestFavoritesServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestFavoritesServerList")
	}
	addr_iSteamMatchmakingServers_RequestHistoryServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestHistoryServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestHistoryServerList")
	}
	addr_iSteamMatchmakingServers_RequestSpectatorServerList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RequestSpectatorServerList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RequestSpectatorServerList")
	}
	addr_iSteamMatchmakingServers_ReleaseRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_ReleaseRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_ReleaseRequest")
	}
	addr_iSteamMatchmakingServers_GetServerDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_GetServerDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_GetServerDetails")
	}
	addr_iSteamMatchmakingServers_CancelQuery, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_CancelQuery")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_CancelQuery")
	}
	addr_iSteamMatchmakingServers_RefreshQuery, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RefreshQuery")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RefreshQuery")
	}
	addr_iSteamMatchmakingServers_IsRefreshing, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_IsRefreshing")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_IsRefreshing")
	}
	addr_iSteamMatchmakingServers_GetServerCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_GetServerCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_GetServerCount")
	}
	addr_iSteamMatchmakingServers_RefreshServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_RefreshServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_RefreshServer")
	}
	addr_iSteamMatchmakingServers_PingServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_PingServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_PingServer")
	}
	addr_iSteamMatchmakingServers_PlayerDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_PlayerDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_PlayerDetails")
	}
	addr_iSteamMatchmakingServers_ServerRules, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_ServerRules")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_ServerRules")
	}
	addr_iSteamMatchmakingServers_CancelServerQuery, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServers_CancelServerQuery")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServers_CancelServerQuery")
	}
	addr_iSteamMatchmakingServerListResponse_ServerResponded, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServerListResponse_ServerResponded")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServerListResponse_ServerResponded")
	}
	addr_iSteamMatchmakingServerListResponse_ServerFailedToRespond, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond")
	}
	addr_iSteamMatchmakingServerListResponse_RefreshComplete, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingServerListResponse_RefreshComplete")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingServerListResponse_RefreshComplete")
	}
	addr_iSteamMatchmakingPingResponse_ServerResponded, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingPingResponse_ServerResponded")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingPingResponse_ServerResponded")
	}
	addr_iSteamMatchmakingPingResponse_ServerFailedToRespond, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond")
	}
	addr_iSteamMatchmakingPlayersResponse_AddPlayerToList, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList")
	}
	addr_iSteamMatchmakingPlayersResponse_PlayersFailedToRespond, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond")
	}
	addr_iSteamMatchmakingPlayersResponse_PlayersRefreshComplete, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete")
	}
	addr_iSteamMatchmakingRulesResponse_RulesResponded, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingRulesResponse_RulesResponded")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingRulesResponse_RulesResponded")
	}
	addr_iSteamMatchmakingRulesResponse_RulesFailedToRespond, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond")
	}
	addr_iSteamMatchmakingRulesResponse_RulesRefreshComplete, err = openSymbol(steamAPILib, "SteamAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete")
	}

	matchMakingKeyValuePair_t_Construct = func(self *MatchMakingKeyValuePair) {
		purego.SyscallN(addr_matchMakingKeyValuePair_t_Construct, structToUintptr[MatchMakingKeyValuePair](self))
	}
	servernetadr_t_Construct = func(self *ServerNetAdr) {
		purego.SyscallN(addr_servernetadr_t_Construct, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	servernetadr_t_Init = func(self *ServerNetAdr, ip uint, QueryPort uint16, ConnectionPort uint16) {
		purego.SyscallN(addr_servernetadr_t_Init, uintptr(unsafe.Pointer(self)), uintptr(ip), uintptr(QueryPort), uintptr(ConnectionPort))
		runtime.KeepAlive(self)
	}
	servernetadr_t_GetQueryPort = func(self *ServerNetAdr) uint16 {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_GetQueryPort, uintptr(unsafe.Pointer(self)))
		__r0 := uint16(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	servernetadr_t_SetQueryPort = func(self *ServerNetAdr, Port uint16) {
		purego.SyscallN(addr_servernetadr_t_SetQueryPort, uintptr(unsafe.Pointer(self)), uintptr(Port))
		runtime.KeepAlive(self)
	}
	servernetadr_t_GetConnectionPort = func(self *ServerNetAdr) uint16 {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_GetConnectionPort, uintptr(unsafe.Pointer(self)))
		__r0 := uint16(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	servernetadr_t_SetConnectionPort = func(self *ServerNetAdr, Port uint16) {
		purego.SyscallN(addr_servernetadr_t_SetConnectionPort, uintptr(unsafe.Pointer(self)), uintptr(Port))
		runtime.KeepAlive(self)
	}
	servernetadr_t_GetIP = func(self *ServerNetAdr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_GetIP, uintptr(unsafe.Pointer(self)))
		__r0 := uint32(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	servernetadr_t_SetIP = func(self *ServerNetAdr, IP uint32) {
		purego.SyscallN(addr_servernetadr_t_SetIP, uintptr(unsafe.Pointer(self)), uintptr(IP))
		runtime.KeepAlive(self)
	}
	servernetadr_t_GetConnectionAddressString = func(self *ServerNetAdr) string {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_GetConnectionAddressString, uintptr(unsafe.Pointer(self)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	servernetadr_t_GetQueryAddressString = func(self *ServerNetAdr) string {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_GetQueryAddressString, uintptr(unsafe.Pointer(self)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	servernetadr_t_IsLessThan = func(self *ServerNetAdr, netadr *ServerNetAdr) bool {
		_r0, _, _ := purego.SyscallN(addr_servernetadr_t_IsLessThan, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(netadr)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(netadr)
		return __r0
	}
	servernetadr_t_Assign = func(self *ServerNetAdr, that *ServerNetAdr) {
		purego.SyscallN(addr_servernetadr_t_Assign, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(that)))
		runtime.KeepAlive(self)
		runtime.KeepAlive(that)
	}
	gameserveritem_t_Construct = func(self *GameServerItem) {
		purego.SyscallN(addr_gameserveritem_t_Construct, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	gameserveritem_t_GetName = func(self *GameServerItem) string {
		_r0, _, _ := purego.SyscallN(addr_gameserveritem_t_GetName, uintptr(unsafe.Pointer(self)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	gameserveritem_t_SetName = func(self *GameServerItem, Name string) {
		purego.SyscallN(addr_gameserveritem_t_SetName, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(Name))))
		runtime.KeepAlive(self)
		runtime.KeepAlive(Name)
	}
	steamMatchmakingServers_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamMatchmakingServers_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamMatchmakingServers_RequestInternetServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestInternetServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(ppchFilters)), uintptr(nFilters), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(ppchFilters)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_RequestLANServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestLANServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_RequestFriendsServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestFriendsServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(ppchFilters)), uintptr(nFilters), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(ppchFilters)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_RequestFavoritesServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestFavoritesServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(ppchFilters)), uintptr(nFilters), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(ppchFilters)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_RequestHistoryServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestHistoryServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(ppchFilters)), uintptr(nFilters), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(ppchFilters)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_RequestSpectatorServerList = func(steamMatchmakingServers uintptr, iApp AppId_t, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_RequestSpectatorServerList, uintptr(steamMatchmakingServers), uintptr(iApp), uintptr(unsafe.Pointer(ppchFilters)), uintptr(nFilters), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerListRequest(_r0)
		runtime.KeepAlive(ppchFilters)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_ReleaseRequest = func(steamMatchmakingServers uintptr, hServerListRequest HServerListRequest) {
		purego.SyscallN(addr_iSteamMatchmakingServers_ReleaseRequest, uintptr(steamMatchmakingServers), uintptr(hServerListRequest))
	}
	iSteamMatchmakingServers_GetServerDetails = func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32) *GameServerItem {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_GetServerDetails, uintptr(steamMatchmakingServers), uintptr(hRequest), uintptr(iServer))
		__r0 := uintptrToStruct[GameServerItem](_r0)
		return __r0
	}
	iSteamMatchmakingServers_CancelQuery = func(steamMatchmakingServers uintptr, hRequest HServerListRequest) {
		purego.SyscallN(addr_iSteamMatchmakingServers_CancelQuery, uintptr(steamMatchmakingServers), uintptr(hRequest))
	}
	iSteamMatchmakingServers_RefreshQuery = func(steamMatchmakingServers uintptr, hRequest HServerListRequest) {
		purego.SyscallN(addr_iSteamMatchmakingServers_RefreshQuery, uintptr(steamMatchmakingServers), uintptr(hRequest))
	}
	iSteamMatchmakingServers_IsRefreshing = func(steamMatchmakingServers uintptr, hRequest HServerListRequest) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_IsRefreshing, uintptr(steamMatchmakingServers), uintptr(hRequest))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamMatchmakingServers_GetServerCount = func(steamMatchmakingServers uintptr, hRequest HServerListRequest) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_GetServerCount, uintptr(steamMatchmakingServers), uintptr(hRequest))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamMatchmakingServers_RefreshServer = func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32) {
		purego.SyscallN(addr_iSteamMatchmakingServers_RefreshServer, uintptr(steamMatchmakingServers), uintptr(hRequest), uintptr(iServer))
	}
	iSteamMatchmakingServers_PingServer = func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *ISteamMatchmakingPingResponse) HServerQuery {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_PingServer, uintptr(steamMatchmakingServers), uintptr(unIP), uintptr(usPort), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerQuery(_r0)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_PlayerDetails = func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingPlayersResponse) HServerQuery {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_PlayerDetails, uintptr(steamMatchmakingServers), uintptr(unIP), uintptr(usPort), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerQuery(_r0)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_ServerRules = func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingRulesResponse) HServerQuery {
		_r0, _, _ := purego.SyscallN(addr_iSteamMatchmakingServers_ServerRules, uintptr(steamMatchmakingServers), uintptr(unIP), uintptr(usPort), uintptr(unsafe.Pointer(pRequestServersResponse)))
		__r0 := HServerQuery(_r0)
		runtime.KeepAlive(pRequestServersResponse)
		return __r0
	}
	iSteamMatchmakingServers_CancelServerQuery = func(steamMatchmakingServers uintptr, hServerQuery HServerQuery) {
		purego.SyscallN(addr_iSteamMatchmakingServers_CancelServerQuery, uintptr(steamMatchmakingServers), uintptr(hServerQuery))
	}
	iSteamMatchmakingServerListResponse_ServerResponded = func(hRequest HServerListRequest, iServer int32) {
		purego.SyscallN(addr_iSteamMatchmakingServerListResponse_ServerResponded, uintptr(hRequest), uintptr(iServer))
	}
	iSteamMatchmakingServerListResponse_ServerFailedToRespond = func(hRequest HServerListRequest, iServer int32) {
		purego.SyscallN(addr_iSteamMatchmakingServerListResponse_ServerFailedToRespond, uintptr(hRequest), uintptr(iServer))
	}
	iSteamMatchmakingServerListResponse_RefreshComplete = func(hRequest HServerListRequest, response EMatchMakingServerResponse) {
		purego.SyscallN(addr_iSteamMatchmakingServerListResponse_RefreshComplete, uintptr(hRequest), uintptr(response))
	}
	iSteamMatchmakingPingResponse_ServerResponded = func(server *GameServerItem) {
		purego.SyscallN(addr_iSteamMatchmakingPingResponse_ServerResponded, uintptr(unsafe.Pointer(server)))
		runtime.KeepAlive(server)
	}
	iSteamMatchmakingPingResponse_ServerFailedToRespond = func() {
		purego.SyscallN(addr_iSteamMatchmakingPingResponse_ServerFailedToRespond)
	}
	iSteamMatchmakingPlayersResponse_PlayersFailedToRespond = func() {
		purego.SyscallN(addr_iSteamMatchmakingPlayersResponse_PlayersFailedToRespond)
	}
	iSteamMatchmakingPlayersResponse_PlayersRefreshComplete = func() {
		purego.SyscallN(addr_iSteamMatchmakingPlayersResponse_PlayersRefreshComplete)
	}
	iSteamMatchmakingRulesResponse_RulesResponded = func(pchRule string, pchValue string) {
		purego.SyscallN(addr_iSteamMatchmakingRulesResponse_RulesResponded, uintptr(unsafe.Pointer(bytePtrFromString(pchRule))), uintptr(unsafe.Pointer(bytePtrFromString(pchValue))))
		runtime.KeepAlive(pchRule)
		runtime.KeepAlive(pchValue)
	}
	iSteamMatchmakingRulesResponse_RulesFailedToRespond = func() {
		purego.SyscallN(addr_iSteamMatchmakingRulesResponse_RulesFailedToRespond)
	}
	iSteamMatchmakingRulesResponse_RulesRefreshComplete = func() {
		purego.SyscallN(addr_iSteamMatchmakingRulesResponse_RulesRefreshComplete)
	}

	purego.RegisterLibFunc(&iSteamMatchmakingPlayersResponse_AddPlayerToList, steamAPILib, flatAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList)
}

var (
	addr_steamIPAddress_IsSet                            uintptr
	addr_steamNetworkingIPAddr_Clear                     uintptr
	addr_steamNetworkingIPAddr_IsIPv6AllZeros            uintptr
	addr_steamNetworkingIPAddr_SetIPv6                   uintptr
	addr_steamNetworkingIPAddr_SetIPv4                   uintptr
	addr_steamNetworkingIPAddr_IsIPv4                    uintptr
	addr_steamNetworkingIPAddr_GetIPv4                   uintptr
	addr_steamNetworkingIPAddr_SetIPv6LocalHost          uintptr
	addr_steamNetworkingIPAddr_IsLocalHost               uintptr
	addr_steamNetworkingIPAddr_ToString                  uintptr
	addr_steamNetworkingIPAddr_ParseString               uintptr
	addr_steamNetworkingIPAddr_IsEqualTo                 uintptr
	addr_steamNetworkingIPAddr_GetFakeIPType             uintptr
	addr_steamNetworkingIPAddr_IsFakeIP                  uintptr
	addr_steamNetworkingIdentity_Clear                   uintptr
	addr_steamNetworkingIdentity_IsInvalid               uintptr
	addr_steamNetworkingIdentity_SetSteamID              uintptr
	addr_steamNetworkingIdentity_GetSteamID              uintptr
	addr_steamNetworkingIdentity_SetSteamID64            uintptr
	addr_steamNetworkingIdentity_GetSteamID64            uintptr
	addr_steamNetworkingIdentity_SetXboxPairwiseID       uintptr
	addr_steamNetworkingIdentity_GetXboxPairwiseID       uintptr
	addr_steamNetworkingIdentity_SetPSNID                uintptr
	addr_steamNetworkingIdentity_GetPSNID                uintptr
	addr_steamNetworkingIdentity_SetIPAddr               uintptr
	addr_steamNetworkingIdentity_GetIPAddr               uintptr
	addr_steamNetworkingIdentity_SetIPv4Addr             uintptr
	addr_steamNetworkingIdentity_GetIPv4                 uintptr
	addr_steamNetworkingIdentity_GetFakeIPType           uintptr
	addr_steamNetworkingIdentity_IsFakeIP                uintptr
	addr_steamNetworkingIdentity_SetLocalHost            uintptr
	addr_steamNetworkingIdentity_IsLocalHost             uintptr
	addr_steamNetworkingIdentity_SetGenericString        uintptr
	addr_steamNetworkingIdentity_GetGenericString        uintptr
	addr_steamNetworkingIdentity_SetGenericBytes         uintptr
	addr_steamNetworkingIdentity_GetGenericBytes         uintptr
	addr_steamNetworkingIdentity_IsEqualTo               uintptr
	addr_steamNetworkingIdentity_ToString                uintptr
	addr_steamNetworkingIdentity_ParseString             uintptr
	addr_steamNetworkingConfigValue_t_SetInt32           uintptr
	addr_steamNetworkingConfigValue_t_SetInt64           uintptr
	addr_steamNetworkingConfigValue_t_SetFloat           uintptr
	addr_steamNetworkingConfigValue_t_SetPtr             uintptr
	addr_steamNetworkingConfigValue_t_SetString          uintptr
	addr_steamDatagramHostedAddress_Clear                uintptr
	addr_steamDatagramHostedAddress_GetPopID             uintptr
	addr_steamDatagramHostedAddress_SetDevAddress        uintptr
	addr_iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort  uintptr
	addr_iSteamNetworkingFakeUDPPort_SendMessageToFakeIP uintptr
	addr_iSteamNetworkingFakeUDPPort_ReceiveMessages     uintptr
	addr_iSteamNetworkingFakeUDPPort_ScheduleCleanup     uintptr
)

func initNetworkTypes() {
	var err error
	addr_steamIPAddress_IsSet, err = openSymbol(steamAPILib, "SteamAPI_SteamIPAddress_t_IsSet")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamIPAddress_t_IsSet")
	}
	addr_steamNetworkingIPAddr_Clear, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_Clear")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_Clear")
	}
	addr_steamNetworkingIPAddr_IsIPv6AllZeros, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_IsIPv6AllZeros")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_IsIPv6AllZeros")
	}
	addr_steamNetworkingIPAddr_SetIPv6, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_SetIPv6")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_SetIPv6")
	}
	addr_steamNetworkingIPAddr_SetIPv4, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_SetIPv4")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_SetIPv4")
	}
	addr_steamNetworkingIPAddr_IsIPv4, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_IsIPv4")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_IsIPv4")
	}
	addr_steamNetworkingIPAddr_GetIPv4, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_GetIPv4")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_GetIPv4")
	}
	addr_steamNetworkingIPAddr_SetIPv6LocalHost, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_SetIPv6LocalHost")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_SetIPv6LocalHost")
	}
	addr_steamNetworkingIPAddr_IsLocalHost, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_IsLocalHost")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_IsLocalHost")
	}
	addr_steamNetworkingIPAddr_ToString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_ToString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_ToString")
	}
	addr_steamNetworkingIPAddr_ParseString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_ParseString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_ParseString")
	}
	addr_steamNetworkingIPAddr_IsEqualTo, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_IsEqualTo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_IsEqualTo")
	}
	addr_steamNetworkingIPAddr_GetFakeIPType, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_GetFakeIPType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_GetFakeIPType")
	}
	addr_steamNetworkingIPAddr_IsFakeIP, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIPAddr_IsFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIPAddr_IsFakeIP")
	}
	addr_steamNetworkingIdentity_Clear, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_Clear")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_Clear")
	}
	addr_steamNetworkingIdentity_IsInvalid, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_IsInvalid")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_IsInvalid")
	}
	addr_steamNetworkingIdentity_SetSteamID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetSteamID")
	}
	addr_steamNetworkingIdentity_GetSteamID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetSteamID")
	}
	addr_steamNetworkingIdentity_SetSteamID64, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetSteamID64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetSteamID64")
	}
	addr_steamNetworkingIdentity_GetSteamID64, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetSteamID64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetSteamID64")
	}
	addr_steamNetworkingIdentity_SetXboxPairwiseID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetXboxPairwiseID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetXboxPairwiseID")
	}
	addr_steamNetworkingIdentity_GetXboxPairwiseID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetXboxPairwiseID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetXboxPairwiseID")
	}
	addr_steamNetworkingIdentity_SetPSNID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetPSNID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetPSNID")
	}
	addr_steamNetworkingIdentity_GetPSNID, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetPSNID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetPSNID")
	}
	addr_steamNetworkingIdentity_SetIPAddr, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetIPAddr")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetIPAddr")
	}
	addr_steamNetworkingIdentity_GetIPAddr, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetIPAddr")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetIPAddr")
	}
	addr_steamNetworkingIdentity_SetIPv4Addr, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetIPv4Addr")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetIPv4Addr")
	}
	addr_steamNetworkingIdentity_GetIPv4, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetIPv4")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetIPv4")
	}
	addr_steamNetworkingIdentity_GetFakeIPType, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetFakeIPType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetFakeIPType")
	}
	addr_steamNetworkingIdentity_IsFakeIP, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_IsFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_IsFakeIP")
	}
	addr_steamNetworkingIdentity_SetLocalHost, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetLocalHost")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetLocalHost")
	}
	addr_steamNetworkingIdentity_IsLocalHost, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_IsLocalHost")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_IsLocalHost")
	}
	addr_steamNetworkingIdentity_SetGenericString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetGenericString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetGenericString")
	}
	addr_steamNetworkingIdentity_GetGenericString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetGenericString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetGenericString")
	}
	addr_steamNetworkingIdentity_SetGenericBytes, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_SetGenericBytes")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_SetGenericBytes")
	}
	addr_steamNetworkingIdentity_GetGenericBytes, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_GetGenericBytes")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_GetGenericBytes")
	}
	addr_steamNetworkingIdentity_IsEqualTo, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_IsEqualTo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_IsEqualTo")
	}
	addr_steamNetworkingIdentity_ToString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_ToString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_ToString")
	}
	addr_steamNetworkingIdentity_ParseString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingIdentity_ParseString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingIdentity_ParseString")
	}
	addr_steamNetworkingConfigValue_t_SetInt32, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingConfigValue_t_SetInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingConfigValue_t_SetInt32")
	}
	addr_steamNetworkingConfigValue_t_SetInt64, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingConfigValue_t_SetInt64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingConfigValue_t_SetInt64")
	}
	addr_steamNetworkingConfigValue_t_SetFloat, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingConfigValue_t_SetFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingConfigValue_t_SetFloat")
	}
	addr_steamNetworkingConfigValue_t_SetPtr, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingConfigValue_t_SetPtr")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingConfigValue_t_SetPtr")
	}
	addr_steamNetworkingConfigValue_t_SetString, err = openSymbol(steamAPILib, "SteamAPI_SteamNetworkingConfigValue_t_SetString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamNetworkingConfigValue_t_SetString")
	}
	addr_steamDatagramHostedAddress_Clear, err = openSymbol(steamAPILib, "SteamAPI_SteamDatagramHostedAddress_Clear")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamDatagramHostedAddress_Clear")
	}
	addr_steamDatagramHostedAddress_GetPopID, err = openSymbol(steamAPILib, "SteamAPI_SteamDatagramHostedAddress_GetPopID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamDatagramHostedAddress_GetPopID")
	}
	addr_steamDatagramHostedAddress_SetDevAddress, err = openSymbol(steamAPILib, "SteamAPI_SteamDatagramHostedAddress_SetDevAddress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_SteamDatagramHostedAddress_SetDevAddress")
	}
	addr_iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort")
	}
	addr_iSteamNetworkingFakeUDPPort_SendMessageToFakeIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP")
	}
	addr_iSteamNetworkingFakeUDPPort_ReceiveMessages, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages")
	}
	addr_iSteamNetworkingFakeUDPPort_ScheduleCleanup, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup")
	}

	steamIPAddress_IsSet = func(self *SteamIPAddress) bool {
		_r0, _, _ := purego.SyscallN(addr_steamIPAddress_IsSet, structToUintptr[SteamIPAddress](self))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	steamNetworkingIPAddr_Clear = func(self *SteamNetworkingIPAddr) {
		purego.SyscallN(addr_steamNetworkingIPAddr_Clear, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	steamNetworkingIPAddr_IsIPv6AllZeros = func(self *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_IsIPv6AllZeros, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIPAddr_SetIPv6 = func(self *SteamNetworkingIPAddr, ipv6 string, Port uint16) {
		purego.SyscallN(addr_steamNetworkingIPAddr_SetIPv6, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(ipv6))), uintptr(Port))
		runtime.KeepAlive(self)
		runtime.KeepAlive(ipv6)
	}
	steamNetworkingIPAddr_SetIPv4 = func(self *SteamNetworkingIPAddr, IP uint16, Port uint16) {
		purego.SyscallN(addr_steamNetworkingIPAddr_SetIPv4, uintptr(unsafe.Pointer(self)), uintptr(IP), uintptr(Port))
		runtime.KeepAlive(self)
	}
	steamNetworkingIPAddr_IsIPv4 = func(self *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_IsIPv4, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIPAddr_GetIPv4 = func(self *SteamNetworkingIPAddr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_GetIPv4, uintptr(unsafe.Pointer(self)))
		__r0 := uint32(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIPAddr_SetIPv6LocalHost = func(self *SteamNetworkingIPAddr, Port uint16) {
		purego.SyscallN(addr_steamNetworkingIPAddr_SetIPv6LocalHost, uintptr(unsafe.Pointer(self)), uintptr(Port))
		runtime.KeepAlive(self)
	}
	steamNetworkingIPAddr_IsLocalHost = func(self *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_IsLocalHost, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIPAddr_ToString = func(self *SteamNetworkingIPAddr, buf string, cbBuf uint32, bWithPort bool) {
		purego.SyscallN(addr_steamNetworkingIPAddr_ToString, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(buf))), uintptr(cbBuf), boolToUintptr(bWithPort))
		runtime.KeepAlive(self)
		runtime.KeepAlive(buf)
	}
	steamNetworkingIPAddr_ParseString = func(self *SteamNetworkingIPAddr, pszStr string) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_ParseString, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(pszStr))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(pszStr)
		return __r0
	}
	steamNetworkingIPAddr_IsEqualTo = func(self *SteamNetworkingIPAddr, x *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_IsEqualTo, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(x)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(x)
		return __r0
	}
	steamNetworkingIPAddr_GetFakeIPType = func(self *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_GetFakeIPType, uintptr(unsafe.Pointer(self)))
		__r0 := ESteamNetworkingFakeIPType(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIPAddr_IsFakeIP = func(self *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIPAddr_IsFakeIP, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_Clear = func(self *SteamNetworkingIdentity) {
		purego.SyscallN(addr_steamNetworkingIdentity_Clear, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_IsInvalid = func(self *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_IsInvalid, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetSteamID = func(self *SteamNetworkingIdentity, steamID Uint64SteamID) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetSteamID, uintptr(unsafe.Pointer(self)), uintptr(steamID))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_GetSteamID = func(self *SteamNetworkingIdentity) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetSteamID, uintptr(unsafe.Pointer(self)))
		__r0 := Uint64SteamID(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetSteamID64 = func(self *SteamNetworkingIdentity, steamID uint64) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetSteamID64, uintptr(unsafe.Pointer(self)), uintptr(steamID))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_GetSteamID64 = func(self *SteamNetworkingIdentity) uint64 {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetSteamID64, uintptr(unsafe.Pointer(self)))
		__r0 := uint64(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetXboxPairwiseID = func(self *SteamNetworkingIdentity, pszString string) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_SetXboxPairwiseID, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(pszString))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(pszString)
		return __r0
	}
	steamNetworkingIdentity_GetXboxPairwiseID = func(self *SteamNetworkingIdentity) string {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetXboxPairwiseID, uintptr(unsafe.Pointer(self)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetPSNID = func(self *SteamNetworkingIdentity, id uint64) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetPSNID, uintptr(unsafe.Pointer(self)), uintptr(id))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_GetPSNID = func(self *SteamNetworkingIdentity) uint64 {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetPSNID, uintptr(unsafe.Pointer(self)))
		__r0 := uint64(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetIPAddr = func(self *SteamNetworkingIdentity, addr *SteamNetworkingIPAddr) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetIPAddr, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(addr)))
		runtime.KeepAlive(self)
		runtime.KeepAlive(addr)
	}
	steamNetworkingIdentity_GetIPAddr = func(self *SteamNetworkingIdentity) *SteamNetworkingIPAddr {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetIPAddr, uintptr(unsafe.Pointer(self)))
		__r0 := uintptrToStruct[SteamNetworkingIPAddr](_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetIPv4Addr = func(self *SteamNetworkingIdentity, nIPv4u int32, nPort uint16) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetIPv4Addr, uintptr(unsafe.Pointer(self)), uintptr(nIPv4u), uintptr(nPort))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_GetIPv4 = func(self *SteamNetworkingIdentity) uint32 {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetIPv4, uintptr(unsafe.Pointer(self)))
		__r0 := uint32(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_GetFakeIPType = func(self *SteamNetworkingIdentity) ESteamNetworkingFakeIPType {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetFakeIPType, uintptr(unsafe.Pointer(self)))
		__r0 := ESteamNetworkingFakeIPType(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_IsFakeIP = func(self *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_IsFakeIP, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetLocalHost = func(self *SteamNetworkingIdentity) {
		purego.SyscallN(addr_steamNetworkingIdentity_SetLocalHost, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	steamNetworkingIdentity_IsLocalHost = func(self *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_IsLocalHost, uintptr(unsafe.Pointer(self)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetGenericString = func(self *SteamNetworkingIdentity, pszString string) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_SetGenericString, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(pszString))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(pszString)
		return __r0
	}
	steamNetworkingIdentity_GetGenericString = func(self *SteamNetworkingIdentity) string {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetGenericString, uintptr(unsafe.Pointer(self)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamNetworkingIdentity_SetGenericBytes = func(self *SteamNetworkingIdentity, data []byte, cbLen uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_SetGenericBytes, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uintptr(cbLen))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(data)
		return __r0
	}
	steamNetworkingIdentity_GetGenericBytes = func(self *SteamNetworkingIdentity, cbLen *int) string {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_GetGenericBytes, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(cbLen)))
		__r0 := goString(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(cbLen)
		return __r0
	}
	steamNetworkingIdentity_IsEqualTo = func(self *SteamNetworkingIdentity, x *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_IsEqualTo, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(x)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(x)
		return __r0
	}
	steamNetworkingIdentity_ToString = func(self *SteamNetworkingIdentity, buf string, cbBuf uint32) {
		purego.SyscallN(addr_steamNetworkingIdentity_ToString, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(buf))), uintptr(cbBuf))
		runtime.KeepAlive(self)
		runtime.KeepAlive(buf)
	}
	steamNetworkingIdentity_ParseString = func(self *SteamNetworkingIdentity, pszStr string) bool {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingIdentity_ParseString, uintptr(unsafe.Pointer(self)), uintptr(unsafe.Pointer(bytePtrFromString(pszStr))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(self)
		runtime.KeepAlive(pszStr)
		return __r0
	}
	steamNetworkingConfigValue_t_SetInt32 = func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int32) {
		purego.SyscallN(addr_steamNetworkingConfigValue_t_SetInt32, uintptr(unsafe.Pointer(self)), uintptr(eVal), uintptr(data))
		runtime.KeepAlive(self)
	}
	steamNetworkingConfigValue_t_SetInt64 = func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int64) {
		purego.SyscallN(addr_steamNetworkingConfigValue_t_SetInt64, uintptr(unsafe.Pointer(self)), uintptr(eVal), uintptr(data))
		runtime.KeepAlive(self)
	}
	steamNetworkingConfigValue_t_SetPtr = func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data []byte) {
		purego.SyscallN(addr_steamNetworkingConfigValue_t_SetPtr, uintptr(unsafe.Pointer(self)), uintptr(eVal), uintptr(unsafe.Pointer(unsafe.SliceData(data))))
		runtime.KeepAlive(self)
		runtime.KeepAlive(data)
	}
	steamNetworkingConfigValue_t_SetString = func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data string) {
		purego.SyscallN(addr_steamNetworkingConfigValue_t_SetString, uintptr(unsafe.Pointer(self)), uintptr(eVal), uintptr(unsafe.Pointer(bytePtrFromString(data))))
		runtime.KeepAlive(self)
		runtime.KeepAlive(data)
	}
	steamDatagramHostedAddress_Clear = func(self *SteamDatagramHostedAddress) {
		purego.SyscallN(addr_steamDatagramHostedAddress_Clear, uintptr(unsafe.Pointer(self)))
		runtime.KeepAlive(self)
	}
	steamDatagramHostedAddress_GetPopID = func(self *SteamDatagramHostedAddress) SteamNetworkingPOPID {
		_r0, _, _ := purego.SyscallN(addr_steamDatagramHostedAddress_GetPopID, uintptr(unsafe.Pointer(self)))
		__r0 := SteamNetworkingPOPID(_r0)
		runtime.KeepAlive(self)
		return __r0
	}
	steamDatagramHostedAddress_SetDevAddress = func(self *SteamDatagramHostedAddress, nIP uint32, nPort uint16, popid SteamNetworkingPOPID) {
		purego.SyscallN(addr_steamDatagramHostedAddress_SetDevAddress, uintptr(unsafe.Pointer(self)), uintptr(nIP), uintptr(nPort), uintptr(popid))
		runtime.KeepAlive(self)
	}
	iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort = func() {
		purego.SyscallN(addr_iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort)
	}
	iSteamNetworkingFakeUDPPort_SendMessageToFakeIP = func(remoteAddress *SteamNetworkingIPAddr, pData uintptr, cbData uint32, nSendFlags int32) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingFakeUDPPort_SendMessageToFakeIP, uintptr(unsafe.Pointer(remoteAddress)), uintptr(pData), uintptr(cbData), uintptr(nSendFlags))
		__r0 := EResult(_r0)
		runtime.KeepAlive(remoteAddress)
		return __r0
	}
	iSteamNetworkingFakeUDPPort_ReceiveMessages = func(ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingFakeUDPPort_ReceiveMessages, uintptr(unsafe.Pointer(ppOutMessages)), uintptr(nMaxMessages))
		__r0 := int32(_r0)
		runtime.KeepAlive(ppOutMessages)
		return __r0
	}
	iSteamNetworkingFakeUDPPort_ScheduleCleanup = func(remoteAddress *SteamNetworkingIPAddr) {
		purego.SyscallN(addr_iSteamNetworkingFakeUDPPort_ScheduleCleanup, uintptr(unsafe.Pointer(remoteAddress)))
		runtime.KeepAlive(remoteAddress)
	}

	purego.RegisterLibFunc(&steamNetworkingConfigValue_t_SetFloat, steamAPILib, flatAPI_SteamNetworkingConfigValue_t_SetFloat)
}

var (
	addr_steamNetworkingMessages_get                       uintptr
	addr_iSteamNetworkingMessages_SendMessageToUser        uintptr
	addr_iSteamNetworkingMessages_ReceiveMessagesOnChannel uintptr
	addr_iSteamNetworkingMessages_AcceptSessionWithUser    uintptr
	addr_iSteamNetworkingMessages_CloseSessionWithUser     uintptr
	addr_iSteamNetworkingMessages_CloseChannelWithUser     uintptr
	addr_iSteamNetworkingMessages_GetSessionConnectionInfo uintptr
)

func initNetworkMessages() {
	var err error
	addr_steamNetworkingMessages_get, err = openSymbol(steamAPILib, flatAPI_SteamNetworkingMessages)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamNetworkingMessages)
	}
	addr_iSteamNetworkingMessages_SendMessageToUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_SendMessageToUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_SendMessageToUser")
	}
	addr_iSteamNetworkingMessages_ReceiveMessagesOnChannel, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel")
	}
	addr_iSteamNetworkingMessages_AcceptSessionWithUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_AcceptSessionWithUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_AcceptSessionWithUser")
	}
	addr_iSteamNetworkingMessages_CloseSessionWithUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_CloseSessionWithUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_CloseSessionWithUser")
	}
	addr_iSteamNetworkingMessages_CloseChannelWithUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_CloseChannelWithUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_CloseChannelWithUser")
	}
	addr_iSteamNetworkingMessages_GetSessionConnectionInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingMessages_GetSessionConnectionInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingMessages_GetSessionConnectionInfo")
	}

	steamNetworkingMessages_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingMessages_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamNetworkingMessages_SendMessageToUser = func(steamNetworkingMessages uintptr, identityRemote uintptr, pubData []byte, cubData uint32, nSendFlags int32, nRemoteChannel int32) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_SendMessageToUser, uintptr(steamNetworkingMessages), uintptr(identityRemote), uintptr(unsafe.Pointer(unsafe.SliceData(pubData))), uintptr(cubData), uintptr(nSendFlags), uintptr(nRemoteChannel))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pubData)
		return __r0
	}
	iSteamNetworkingMessages_ReceiveMessagesOnChannel = func(steamNetworkingMessages uintptr, nLocalChannel int32, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_ReceiveMessagesOnChannel, uintptr(steamNetworkingMessages), uintptr(nLocalChannel), uintptr(unsafe.Pointer(ppOutMessages)), uintptr(nMaxMessages))
		__r0 := int32(_r0)
		runtime.KeepAlive(ppOutMessages)
		return __r0
	}
	iSteamNetworkingMessages_AcceptSessionWithUser = func(steamNetworkingMessages uintptr, identityRemote uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_AcceptSessionWithUser, uintptr(steamNetworkingMessages), uintptr(identityRemote))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingMessages_CloseSessionWithUser = func(steamNetworkingMessages uintptr, identityRemote uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_CloseSessionWithUser, uintptr(steamNetworkingMessages), uintptr(identityRemote))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingMessages_CloseChannelWithUser = func(steamNetworkingMessages uintptr, identityRemote uintptr, nLocalChannel int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_CloseChannelWithUser, uintptr(steamNetworkingMessages), uintptr(identityRemote), uintptr(nLocalChannel))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingMessages_GetSessionConnectionInfo = func(steamNetworkingMessages uintptr, identityRemote uintptr, pConnectionInfo *SteamNetConnectionInfo, pQuickStatus *SteamNetConnectionRealTimeStatus) ESteamNetworkingConnectionState {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingMessages_GetSessionConnectionInfo, uintptr(steamNetworkingMessages), uintptr(identityRemote), uintptr(unsafe.Pointer(pConnectionInfo)), uintptr(unsafe.Pointer(pQuickStatus)))
		__r0 := ESteamNetworkingConnectionState(_r0)
		runtime.KeepAlive(pConnectionInfo)
		runtime.KeepAlive(pQuickStatus)
		return __r0
	}
}

var (
	addr_steamNetworkingSockets_get                                      uintptr
	addr_iSteamNetworkingSockets_CreateListenSocketIP                    uintptr
	addr_iSteamNetworkingSockets_ConnectByIPAddress                      uintptr
	addr_iSteamNetworkingSockets_CreateListenSocketP2P                   uintptr
	addr_iSteamNetworkingSockets_ConnectP2P                              uintptr
	addr_iSteamNetworkingSockets_AcceptConnection                        uintptr
	addr_iSteamNetworkingSockets_CloseConnection                         uintptr
	addr_iSteamNetworkingSockets_CloseListenSocket                       uintptr
	addr_iSteamNetworkingSockets_SetConnectionUserData                   uintptr
	addr_iSteamNetworkingSockets_GetConnectionUserData                   uintptr
	addr_iSteamNetworkingSockets_SetConnectionName                       uintptr
	addr_iSteamNetworkingSockets_GetConnectionName                       uintptr
	addr_iSteamNetworkingSockets_SendMessageToConnection                 uintptr
	addr_iSteamNetworkingSockets_SendMessages                            uintptr
	addr_iSteamNetworkingSockets_FlushMessagesOnConnection               uintptr
	addr_iSteamNetworkingSockets_ReceiveMessagesOnConnection             uintptr
	addr_iSteamNetworkingSockets_GetConnectionInfo                       uintptr
	addr_iSteamNetworkingSockets_GetConnectionRealTimeStatus             uintptr
	addr_iSteamNetworkingSockets_GetDetailedConnectionStatus             uintptr
	addr_iSteamNetworkingSockets_GetListenSocketAddress                  uintptr
	addr_iSteamNetworkingSockets_CreateSocketPair                        uintptr
	addr_iSteamNetworkingSockets_ConfigureConnectionLanes                uintptr
	addr_iSteamNetworkingSockets_GetIdentity                             uintptr
	addr_iSteamNetworkingSockets_InitAuthentication                      uintptr
	addr_iSteamNetworkingSockets_GetAuthenticationStatus                 uintptr
	addr_iSteamNetworkingSockets_CreatePollGroup                         uintptr
	addr_iSteamNetworkingSockets_DestroyPollGroup                        uintptr
	addr_iSteamNetworkingSockets_SetConnectionPollGroup                  uintptr
	addr_iSteamNetworkingSockets_ReceiveMessagesOnPollGroup              uintptr
	addr_iSteamNetworkingSockets_ReceivedRelayAuthTicket                 uintptr
	addr_iSteamNetworkingSockets_FindRelayAuthTicketForServer            uintptr
	addr_iSteamNetworkingSockets_ConnectToHostedDedicatedServer          uintptr
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerPort            uintptr
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerPOPID           uintptr
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerAddress         uintptr
	addr_iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket uintptr
	addr_iSteamNetworkingSockets_GetGameCoordinatorServerLogin           uintptr
	addr_iSteamNetworkingSockets_ConnectP2PCustomSignaling               uintptr
	addr_iSteamNetworkingSockets_ReceivedP2PCustomSignal                 uintptr
	addr_iSteamNetworkingSockets_GetCertificateRequest                   uintptr
	addr_iSteamNetworkingSockets_SetCertificate                          uintptr
	addr_iSteamNetworkingSockets_ResetIdentity                           uintptr
	addr_iSteamNetworkingSockets_RunCallbacks                            uintptr
	addr_iSteamNetworkingSockets_BeginAsyncRequestFakeIP                 uintptr
	addr_iSteamNetworkingSockets_GetFakeIP                               uintptr
	addr_iSteamNetworkingSockets_CreateListenSocketP2PFakeIP             uintptr
	addr_iSteamNetworkingSockets_GetRemoteFakeIPForConnection            uintptr
	addr_iSteamNetworkingSockets_CreateFakeUDPPort                       uintptr
)

func initNetworkSockets() {
	var err error
	addr_steamNetworkingSockets_get, err = openSymbol(steamAPILib, flatAPI_SteamNetworkingSockets)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamNetworkingSockets)
	}
	addr_iSteamNetworkingSockets_CreateListenSocketIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateListenSocketIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateListenSocketIP")
	}
	addr_iSteamNetworkingSockets_ConnectByIPAddress, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ConnectByIPAddress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ConnectByIPAddress")
	}
	addr_iSteamNetworkingSockets_CreateListenSocketP2P, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2P")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2P")
	}
	addr_iSteamNetworkingSockets_ConnectP2P, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ConnectP2P")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ConnectP2P")
	}
	addr_iSteamNetworkingSockets_AcceptConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_AcceptConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_AcceptConnection")
	}
	addr_iSteamNetworkingSockets_CloseConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CloseConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CloseConnection")
	}
	addr_iSteamNetworkingSockets_CloseListenSocket, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CloseListenSocket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CloseListenSocket")
	}
	addr_iSteamNetworkingSockets_SetConnectionUserData, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SetConnectionUserData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SetConnectionUserData")
	}
	addr_iSteamNetworkingSockets_GetConnectionUserData, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetConnectionUserData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetConnectionUserData")
	}
	addr_iSteamNetworkingSockets_SetConnectionName, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SetConnectionName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SetConnectionName")
	}
	addr_iSteamNetworkingSockets_GetConnectionName, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetConnectionName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetConnectionName")
	}
	addr_iSteamNetworkingSockets_SendMessageToConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SendMessageToConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SendMessageToConnection")
	}
	addr_iSteamNetworkingSockets_SendMessages, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SendMessages")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SendMessages")
	}
	addr_iSteamNetworkingSockets_FlushMessagesOnConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_FlushMessagesOnConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_FlushMessagesOnConnection")
	}
	addr_iSteamNetworkingSockets_ReceiveMessagesOnConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection")
	}
	addr_iSteamNetworkingSockets_GetConnectionInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetConnectionInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetConnectionInfo")
	}
	addr_iSteamNetworkingSockets_GetConnectionRealTimeStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus")
	}
	addr_iSteamNetworkingSockets_GetDetailedConnectionStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus")
	}
	addr_iSteamNetworkingSockets_GetListenSocketAddress, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetListenSocketAddress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetListenSocketAddress")
	}
	addr_iSteamNetworkingSockets_CreateSocketPair, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateSocketPair")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateSocketPair")
	}
	addr_iSteamNetworkingSockets_ConfigureConnectionLanes, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ConfigureConnectionLanes")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ConfigureConnectionLanes")
	}
	addr_iSteamNetworkingSockets_GetIdentity, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetIdentity")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetIdentity")
	}
	addr_iSteamNetworkingSockets_InitAuthentication, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_InitAuthentication")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_InitAuthentication")
	}
	addr_iSteamNetworkingSockets_GetAuthenticationStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetAuthenticationStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetAuthenticationStatus")
	}
	addr_iSteamNetworkingSockets_CreatePollGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreatePollGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreatePollGroup")
	}
	addr_iSteamNetworkingSockets_DestroyPollGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_DestroyPollGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_DestroyPollGroup")
	}
	addr_iSteamNetworkingSockets_SetConnectionPollGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SetConnectionPollGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SetConnectionPollGroup")
	}
	addr_iSteamNetworkingSockets_ReceiveMessagesOnPollGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup")
	}
	addr_iSteamNetworkingSockets_ReceivedRelayAuthTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket")
	}
	addr_iSteamNetworkingSockets_FindRelayAuthTicketForServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer")
	}
	addr_iSteamNetworkingSockets_ConnectToHostedDedicatedServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer")
	}
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerPort, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort")
	}
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerPOPID, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID")
	}
	addr_iSteamNetworkingSockets_GetHostedDedicatedServerAddress, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress")
	}
	addr_iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket")
	}
	addr_iSteamNetworkingSockets_GetGameCoordinatorServerLogin, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin")
	}
	addr_iSteamNetworkingSockets_ConnectP2PCustomSignaling, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling")
	}
	addr_iSteamNetworkingSockets_ReceivedP2PCustomSignal, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal")
	}
	addr_iSteamNetworkingSockets_GetCertificateRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetCertificateRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetCertificateRequest")
	}
	addr_iSteamNetworkingSockets_SetCertificate, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_SetCertificate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_SetCertificate")
	}
	addr_iSteamNetworkingSockets_ResetIdentity, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_ResetIdentity")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_ResetIdentity")
	}
	addr_iSteamNetworkingSockets_RunCallbacks, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_RunCallbacks")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_RunCallbacks")
	}
	addr_iSteamNetworkingSockets_BeginAsyncRequestFakeIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP")
	}
	addr_iSteamNetworkingSockets_GetFakeIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetFakeIP")
	}
	addr_iSteamNetworkingSockets_CreateListenSocketP2PFakeIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP")
	}
	addr_iSteamNetworkingSockets_GetRemoteFakeIPForConnection, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection")
	}
	addr_iSteamNetworkingSockets_CreateFakeUDPPort, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingSockets_CreateFakeUDPPort")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingSockets_CreateFakeUDPPort")
	}

	steamNetworkingSockets_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingSockets_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_CreateListenSocketIP = func(steamNetworkingSockets uintptr, localAddress uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateListenSocketIP, uintptr(steamNetworkingSockets), uintptr(localAddress), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamListenSocket(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_ConnectByIPAddress = func(steamNetworkingSockets uintptr, address uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ConnectByIPAddress, uintptr(steamNetworkingSockets), uintptr(address), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamNetConnection(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_CreateListenSocketP2P = func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateListenSocketP2P, uintptr(steamNetworkingSockets), uintptr(nLocalVirtualPort), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamListenSocket(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_ConnectP2P = func(steamNetworkingSockets uintptr, identityRemote uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ConnectP2P, uintptr(steamNetworkingSockets), uintptr(identityRemote), uintptr(nRemoteVirtualPort), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamNetConnection(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_AcceptConnection = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_AcceptConnection, uintptr(steamNetworkingSockets), uintptr(hConn))
		__r0 := EResult(_r0)
		return __r0
	}
	iSteamNetworkingSockets_CloseConnection = func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nReason int32, pszDebug string, bEnableLinger bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CloseConnection, uintptr(steamNetworkingSockets), uintptr(hPeer), uintptr(nReason), uintptr(unsafe.Pointer(bytePtrFromString(pszDebug))), boolToUintptr(bEnableLinger))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszDebug)
		return __r0
	}
	iSteamNetworkingSockets_CloseListenSocket = func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CloseListenSocket, uintptr(steamNetworkingSockets), uintptr(hSocket))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_SetConnectionUserData = func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nUserData int64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_SetConnectionUserData, uintptr(steamNetworkingSockets), uintptr(hPeer), uintptr(nUserData))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_GetConnectionUserData = func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection) int64 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetConnectionUserData, uintptr(steamNetworkingSockets), uintptr(hPeer))
		__r0 := int64(_r0)
		return __r0
	}
	iSteamNetworkingSockets_SetConnectionName = func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName string) {
		purego.SyscallN(addr_iSteamNetworkingSockets_SetConnectionName, uintptr(steamNetworkingSockets), uintptr(hPeer), uintptr(unsafe.Pointer(bytePtrFromString(pszName))))
		runtime.KeepAlive(pszName)
	}
	iSteamNetworkingSockets_GetConnectionName = func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName []byte, nMaxLen int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetConnectionName, uintptr(steamNetworkingSockets), uintptr(hPeer), uintptr(unsafe.Pointer(unsafe.SliceData(pszName))), uintptr(nMaxLen))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszName)
		return __r0
	}
	iSteamNetworkingSockets_SendMessageToConnection = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pData []byte, cbData uint32, nSendFlags int32, pOutMessageNumber *int64) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_SendMessageToConnection, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(unsafe.SliceData(pData))), uintptr(cbData), uintptr(nSendFlags), uintptr(unsafe.Pointer(pOutMessageNumber)))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pData)
		runtime.KeepAlive(pOutMessageNumber)
		return __r0
	}
	iSteamNetworkingSockets_SendMessages = func(steamNetworkingSockets uintptr, nMessages int32, pMessages []SteamNetworkingMessage, pOutMessageNumberOrResult []int64) {
		purego.SyscallN(addr_iSteamNetworkingSockets_SendMessages, uintptr(steamNetworkingSockets), uintptr(nMessages), uintptr(unsafe.Pointer(unsafe.SliceData(pMessages))), uintptr(unsafe.Pointer(unsafe.SliceData(pOutMessageNumberOrResult))))
		runtime.KeepAlive(pMessages)
		runtime.KeepAlive(pOutMessageNumberOrResult)
	}
	iSteamNetworkingSockets_FlushMessagesOnConnection = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_FlushMessagesOnConnection, uintptr(steamNetworkingSockets), uintptr(hConn))
		__r0 := EResult(_r0)
		return __r0
	}
	iSteamNetworkingSockets_ReceiveMessagesOnConnection = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ReceiveMessagesOnConnection, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(ppOutMessages)), uintptr(nMaxMessages))
		__r0 := int32(_r0)
		runtime.KeepAlive(ppOutMessages)
		return __r0
	}
	iSteamNetworkingSockets_GetConnectionInfo = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pInfo *SteamNetConnectionInfo) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetConnectionInfo, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(pInfo)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pInfo)
		return __r0
	}
	iSteamNetworkingSockets_GetConnectionRealTimeStatus = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pStatus *SteamNetConnectionRealTimeStatus, nLanes int32, pLanes []SteamNetConnectionRealTimeLaneStatus) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetConnectionRealTimeStatus, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(pStatus)), uintptr(nLanes), uintptr(unsafe.Pointer(unsafe.SliceData(pLanes))))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pStatus)
		runtime.KeepAlive(pLanes)
		return __r0
	}
	iSteamNetworkingSockets_GetDetailedConnectionStatus = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pszBuf []byte, cbBuf int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetDetailedConnectionStatus, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(unsafe.SliceData(pszBuf))), uintptr(cbBuf))
		__r0 := int32(_r0)
		runtime.KeepAlive(pszBuf)
		return __r0
	}
	iSteamNetworkingSockets_GetListenSocketAddress = func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket, address *SteamNetworkingIPAddr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetListenSocketAddress, uintptr(steamNetworkingSockets), uintptr(hSocket), uintptr(unsafe.Pointer(address)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(address)
		return __r0
	}
	iSteamNetworkingSockets_CreateSocketPair = func(steamNetworkingSockets uintptr, pOutConnection1 *HSteamNetConnection, pOutConnection2 *HSteamNetConnection, bUseNetworkLoopback bool, pIdentity1 *SteamNetworkingIdentity, pIdentity2 *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateSocketPair, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pOutConnection1)), uintptr(unsafe.Pointer(pOutConnection2)), boolToUintptr(bUseNetworkLoopback), uintptr(unsafe.Pointer(pIdentity1)), uintptr(unsafe.Pointer(pIdentity2)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pOutConnection1)
		runtime.KeepAlive(pOutConnection2)
		runtime.KeepAlive(pIdentity1)
		runtime.KeepAlive(pIdentity2)
		return __r0
	}
	iSteamNetworkingSockets_ConfigureConnectionLanes = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, nNumLanes int32, pLanePriorities []int32, pLaneWeights []uint16) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ConfigureConnectionLanes, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(nNumLanes), uintptr(unsafe.Pointer(unsafe.SliceData(pLanePriorities))), uintptr(unsafe.Pointer(unsafe.SliceData(pLaneWeights))))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pLanePriorities)
		runtime.KeepAlive(pLaneWeights)
		return __r0
	}
	iSteamNetworkingSockets_GetIdentity = func(steamNetworkingSockets uintptr, pIdentity *SteamNetworkingIdentity) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetIdentity, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pIdentity)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pIdentity)
		return __r0
	}
	iSteamNetworkingSockets_InitAuthentication = func(steamNetworkingSockets uintptr) ESteamNetworkingAvailability {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_InitAuthentication, uintptr(steamNetworkingSockets))
		__r0 := ESteamNetworkingAvailability(_r0)
		return __r0
	}
	iSteamNetworkingSockets_GetAuthenticationStatus = func(steamNetworkingSockets uintptr, pDetails *SteamNetAuthenticationStatus) ESteamNetworkingAvailability {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetAuthenticationStatus, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pDetails)))
		__r0 := ESteamNetworkingAvailability(_r0)
		runtime.KeepAlive(pDetails)
		return __r0
	}
	iSteamNetworkingSockets_CreatePollGroup = func(steamNetworkingSockets uintptr) HSteamNetPollGroup {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreatePollGroup, uintptr(steamNetworkingSockets))
		__r0 := HSteamNetPollGroup(_r0)
		return __r0
	}
	iSteamNetworkingSockets_DestroyPollGroup = func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_DestroyPollGroup, uintptr(steamNetworkingSockets), uintptr(hPollGroup))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_SetConnectionPollGroup = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, hPollGroup HSteamNetPollGroup) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_SetConnectionPollGroup, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(hPollGroup))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_ReceiveMessagesOnPollGroup = func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ReceiveMessagesOnPollGroup, uintptr(steamNetworkingSockets), uintptr(hPollGroup), uintptr(unsafe.Pointer(ppOutMessages)), uintptr(nMaxMessages))
		__r0 := int32(_r0)
		runtime.KeepAlive(ppOutMessages)
		return __r0
	}
	iSteamNetworkingSockets_ReceivedRelayAuthTicket = func(steamNetworkingSockets uintptr, pvTicket []byte, cbTicket int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ReceivedRelayAuthTicket, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(unsafe.SliceData(pvTicket))), uintptr(cbTicket), uintptr(unsafe.Pointer(pOutParsedTicket)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pvTicket)
		runtime.KeepAlive(pOutParsedTicket)
		return __r0
	}
	iSteamNetworkingSockets_FindRelayAuthTicketForServer = func(steamNetworkingSockets uintptr, identityGameServer uintptr, nRemoteVirtualPort int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_FindRelayAuthTicketForServer, uintptr(steamNetworkingSockets), uintptr(identityGameServer), uintptr(nRemoteVirtualPort), uintptr(unsafe.Pointer(pOutParsedTicket)))
		__r0 := int32(_r0)
		runtime.KeepAlive(pOutParsedTicket)
		return __r0
	}
	iSteamNetworkingSockets_ConnectToHostedDedicatedServer = func(steamNetworkingSockets uintptr, identityTarget uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ConnectToHostedDedicatedServer, uintptr(steamNetworkingSockets), uintptr(identityTarget), uintptr(nRemoteVirtualPort), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamNetConnection(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_GetHostedDedicatedServerPort = func(steamNetworkingSockets uintptr) uint16 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetHostedDedicatedServerPort, uintptr(steamNetworkingSockets))
		__r0 := uint16(_r0)
		return __r0
	}
	iSteamNetworkingSockets_GetHostedDedicatedServerPOPID = func(steamNetworkingSockets uintptr) SteamNetworkingPOPID {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetHostedDedicatedServerPOPID, uintptr(steamNetworkingSockets))
		__r0 := SteamNetworkingPOPID(_r0)
		return __r0
	}
	iSteamNetworkingSockets_GetHostedDedicatedServerAddress = func(steamNetworkingSockets uintptr, pRouting *SteamDatagramHostedAddress) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetHostedDedicatedServerAddress, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pRouting)))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pRouting)
		return __r0
	}
	iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket = func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket, uintptr(steamNetworkingSockets), uintptr(nLocalVirtualPort), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamListenSocket(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_GetGameCoordinatorServerLogin = func(steamNetworkingSockets uintptr, pLoginInfo *SteamDatagramGameCoordinatorServerLogin, pcbSignedBlob *int32, pBlob []byte) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetGameCoordinatorServerLogin, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pLoginInfo)), uintptr(unsafe.Pointer(pcbSignedBlob)), uintptr(unsafe.Pointer(unsafe.SliceData(pBlob))))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pLoginInfo)
		runtime.KeepAlive(pcbSignedBlob)
		runtime.KeepAlive(pBlob)
		return __r0
	}
	iSteamNetworkingSockets_ConnectP2PCustomSignaling = func(steamNetworkingSockets uintptr, pSignaling *ISteamNetworkingConnectionSignaling, pPeerIdentity *SteamNetworkingIdentity, nRemoteVirtualPort int32, nOptions int32, pOptions *SteamNetworkingConfigValue) HSteamNetConnection {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ConnectP2PCustomSignaling, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pSignaling)), uintptr(unsafe.Pointer(pPeerIdentity)), uintptr(nRemoteVirtualPort), uintptr(nOptions), uintptr(unsafe.Pointer(pOptions)))
		__r0 := HSteamNetConnection(_r0)
		runtime.KeepAlive(pSignaling)
		runtime.KeepAlive(pPeerIdentity)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_ReceivedP2PCustomSignal = func(steamNetworkingSockets uintptr, pMsg []byte, cbMsg int32, pContext *ISteamNetworkingSignalingRecvContext) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_ReceivedP2PCustomSignal, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(unsafe.SliceData(pMsg))), uintptr(cbMsg), uintptr(unsafe.Pointer(pContext)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pMsg)
		runtime.KeepAlive(pContext)
		return __r0
	}
	iSteamNetworkingSockets_GetCertificateRequest = func(steamNetworkingSockets uintptr, pcbBlob *int32, pBlob []byte, errMsg *SteamNetworkingErrMsg) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetCertificateRequest, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(pcbBlob)), uintptr(unsafe.Pointer(unsafe.SliceData(pBlob))), uintptr(unsafe.Pointer(errMsg)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pcbBlob)
		runtime.KeepAlive(pBlob)
		runtime.KeepAlive(errMsg)
		return __r0
	}
	iSteamNetworkingSockets_SetCertificate = func(steamNetworkingSockets uintptr, pCertificate []byte, cbCertificate int32, errMsg *SteamNetworkingErrMsg) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_SetCertificate, uintptr(steamNetworkingSockets), uintptr(unsafe.Pointer(unsafe.SliceData(pCertificate))), uintptr(cbCertificate), uintptr(unsafe.Pointer(errMsg)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pCertificate)
		runtime.KeepAlive(errMsg)
		return __r0
	}
	iSteamNetworkingSockets_ResetIdentity = func(steamNetworkingSockets uintptr, pIdentity uintptr) {
		purego.SyscallN(addr_iSteamNetworkingSockets_ResetIdentity, uintptr(steamNetworkingSockets), uintptr(pIdentity))
	}
	iSteamNetworkingSockets_RunCallbacks = func(steamNetworkingSockets uintptr) {
		purego.SyscallN(addr_iSteamNetworkingSockets_RunCallbacks, uintptr(steamNetworkingSockets))
	}
	iSteamNetworkingSockets_BeginAsyncRequestFakeIP = func(steamNetworkingSockets uintptr, nNumPorts int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_BeginAsyncRequestFakeIP, uintptr(steamNetworkingSockets), uintptr(nNumPorts))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingSockets_GetFakeIP = func(steamNetworkingSockets uintptr, idxFirstPort int32, pInfo *SteamNetworkingFakeIPResult) {
		purego.SyscallN(addr_iSteamNetworkingSockets_GetFakeIP, uintptr(steamNetworkingSockets), uintptr(idxFirstPort), uintptr(unsafe.Pointer(pInfo)))
		runtime.KeepAlive(pInfo)
	}
	iSteamNetworkingSockets_CreateListenSocketP2PFakeIP = func(steamNetworkingSockets uintptr, idxFakePort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateListenSocketP2PFakeIP, uintptr(steamNetworkingSockets), uintptr(idxFakePort), uintptr(nOptions), uintptr(unsafe.Pointer(unsafe.SliceData(pOptions))))
		__r0 := HSteamListenSocket(_r0)
		runtime.KeepAlive(pOptions)
		return __r0
	}
	iSteamNetworkingSockets_GetRemoteFakeIPForConnection = func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pOutAddr *SteamNetworkingIPAddr) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_GetRemoteFakeIPForConnection, uintptr(steamNetworkingSockets), uintptr(hConn), uintptr(unsafe.Pointer(pOutAddr)))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pOutAddr)
		return __r0
	}
	iSteamNetworkingSockets_CreateFakeUDPPort = func(steamNetworkingSockets uintptr, idxFakeServerPort int32) *SteamNetworkingFakeUDPPort {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingSockets_CreateFakeUDPPort, uintptr(steamNetworkingSockets), uintptr(idxFakeServerPort))
		__r0 := uintptrToStruct[SteamNetworkingFakeUDPPort](_r0)
		return __r0
	}

}

var (
	addr_steamNetworkingUtils_get                                                    uintptr
	addr_iSteamNetworkingUtils_AllocateMessage                                       uintptr
	addr_iSteamNetworkingUtils_InitRelayNetworkAccess                                uintptr
	addr_iSteamNetworkingUtils_GetRelayNetworkStatus                                 uintptr
	addr_iSteamNetworkingUtils_GetLocalPingLocation                                  uintptr
	addr_iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations                   uintptr
	addr_iSteamNetworkingUtils_EstimatePingTimeFromLocalHost                         uintptr
	addr_iSteamNetworkingUtils_ConvertPingLocationToString                           uintptr
	addr_iSteamNetworkingUtils_ParsePingLocationString                               uintptr
	addr_iSteamNetworkingUtils_CheckPingDataUpToDate                                 uintptr
	addr_iSteamNetworkingUtils_GetPingToDataCenter                                   uintptr
	addr_iSteamNetworkingUtils_GetDirectPingToPOP                                    uintptr
	addr_iSteamNetworkingUtils_GetPOPCount                                           uintptr
	addr_iSteamNetworkingUtils_GetPOPList                                            uintptr
	addr_iSteamNetworkingUtils_GetLocalTimestamp                                     uintptr
	addr_iSteamNetworkingUtils_SetDebugOutputFunction                                uintptr
	addr_iSteamNetworkingUtils_IsFakeIPv4                                            uintptr
	addr_iSteamNetworkingUtils_GetIPv4FakeIPType                                     uintptr
	addr_iSteamNetworkingUtils_GetRealIdentityForFakeIP                              uintptr
	addr_iSteamNetworkingUtils_SetGlobalConfigValueInt32                             uintptr
	addr_iSteamNetworkingUtils_SetGlobalConfigValueFloat                             uintptr
	addr_iSteamNetworkingUtils_SetGlobalConfigValueString                            uintptr
	addr_iSteamNetworkingUtils_SetGlobalConfigValuePtr                               uintptr
	addr_iSteamNetworkingUtils_SetConnectionConfigValueInt32                         uintptr
	addr_iSteamNetworkingUtils_SetConnectionConfigValueFloat                         uintptr
	addr_iSteamNetworkingUtils_SetConnectionConfigValueString                        uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged     uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged      uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult                        uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest              uintptr
	addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed               uintptr
	addr_iSteamNetworkingUtils_SetConfigValue                                        uintptr
	addr_iSteamNetworkingUtils_SetConfigValueStruct                                  uintptr
	addr_iSteamNetworkingUtils_GetConfigValue                                        uintptr
	addr_iSteamNetworkingUtils_GetConfigValueInfo                                    uintptr
	addr_iSteamNetworkingUtils_IterateGenericEditableConfigValues                    uintptr
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString                        uintptr
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString                     uintptr
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType                   uintptr
	addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ToString                      uintptr
	addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString                   uintptr
)

func initNetworkUtils() {
	var err error
	addr_steamNetworkingUtils_get, err = openSymbol(steamAPILib, flatAPI_SteamNetworkingUtils)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamNetworkingUtils)
	}
	addr_iSteamNetworkingUtils_AllocateMessage, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_AllocateMessage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_AllocateMessage")
	}
	addr_iSteamNetworkingUtils_InitRelayNetworkAccess, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_InitRelayNetworkAccess")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_InitRelayNetworkAccess")
	}
	addr_iSteamNetworkingUtils_GetRelayNetworkStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetRelayNetworkStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetRelayNetworkStatus")
	}
	addr_iSteamNetworkingUtils_GetLocalPingLocation, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetLocalPingLocation")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetLocalPingLocation")
	}
	addr_iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations")
	}
	addr_iSteamNetworkingUtils_EstimatePingTimeFromLocalHost, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost")
	}
	addr_iSteamNetworkingUtils_ConvertPingLocationToString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_ConvertPingLocationToString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_ConvertPingLocationToString")
	}
	addr_iSteamNetworkingUtils_ParsePingLocationString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_ParsePingLocationString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_ParsePingLocationString")
	}
	addr_iSteamNetworkingUtils_CheckPingDataUpToDate, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_CheckPingDataUpToDate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_CheckPingDataUpToDate")
	}
	addr_iSteamNetworkingUtils_GetPingToDataCenter, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetPingToDataCenter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetPingToDataCenter")
	}
	addr_iSteamNetworkingUtils_GetDirectPingToPOP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetDirectPingToPOP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetDirectPingToPOP")
	}
	addr_iSteamNetworkingUtils_GetPOPCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetPOPCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetPOPCount")
	}
	addr_iSteamNetworkingUtils_GetPOPList, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetPOPList")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetPOPList")
	}
	addr_iSteamNetworkingUtils_GetLocalTimestamp, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetLocalTimestamp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetLocalTimestamp")
	}
	addr_iSteamNetworkingUtils_SetDebugOutputFunction, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetDebugOutputFunction")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetDebugOutputFunction")
	}
	addr_iSteamNetworkingUtils_IsFakeIPv4, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_IsFakeIPv4")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_IsFakeIPv4")
	}
	addr_iSteamNetworkingUtils_GetIPv4FakeIPType, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetIPv4FakeIPType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetIPv4FakeIPType")
	}
	addr_iSteamNetworkingUtils_GetRealIdentityForFakeIP, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP")
	}
	addr_iSteamNetworkingUtils_SetGlobalConfigValueInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32")
	}
	addr_iSteamNetworkingUtils_SetGlobalConfigValueFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat")
	}
	addr_iSteamNetworkingUtils_SetGlobalConfigValueString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueString")
	}
	addr_iSteamNetworkingUtils_SetGlobalConfigValuePtr, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr")
	}
	addr_iSteamNetworkingUtils_SetConnectionConfigValueInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32")
	}
	addr_iSteamNetworkingUtils_SetConnectionConfigValueFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat")
	}
	addr_iSteamNetworkingUtils_SetConnectionConfigValueString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueString")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest")
	}
	addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed")
	}
	addr_iSteamNetworkingUtils_SetConfigValue, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetConfigValue")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetConfigValue")
	}
	addr_iSteamNetworkingUtils_SetConfigValueStruct, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SetConfigValueStruct")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SetConfigValueStruct")
	}
	addr_iSteamNetworkingUtils_GetConfigValue, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetConfigValue")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetConfigValue")
	}
	addr_iSteamNetworkingUtils_GetConfigValueInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_GetConfigValueInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_GetConfigValueInfo")
	}
	addr_iSteamNetworkingUtils_IterateGenericEditableConfigValues, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues")
	}
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString")
	}
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString")
	}
	addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType")
	}
	addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ToString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString")
	}
	addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString, err = openSymbol(steamAPILib, "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString")
	}

	steamNetworkingUtils_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamNetworkingUtils_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_AllocateMessage = func(steamNetworkingUtils uintptr, cbAllocateBuffer int32) *SteamNetworkingMessage {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_AllocateMessage, uintptr(steamNetworkingUtils), uintptr(cbAllocateBuffer))
		__r0 := uintptrToStruct[SteamNetworkingMessage](_r0)
		return __r0
	}
	iSteamNetworkingUtils_InitRelayNetworkAccess = func(steamNetworkingUtils uintptr) {
		purego.SyscallN(addr_iSteamNetworkingUtils_InitRelayNetworkAccess, uintptr(steamNetworkingUtils))
	}
	iSteamNetworkingUtils_GetRelayNetworkStatus = func(steamNetworkingUtils uintptr, pDetails *SteamRelayNetworkStatus) ESteamNetworkingAvailability {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetRelayNetworkStatus, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(pDetails)))
		__r0 := ESteamNetworkingAvailability(_r0)
		runtime.KeepAlive(pDetails)
		return __r0
	}
	iSteamNetworkingUtils_GetLocalPingLocation = func(steamNetworkingUtils uintptr, result *SteamNetworkPingLocation) float32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetLocalPingLocation, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(result)))
		__r0 := float32(_r0)
		runtime.KeepAlive(result)
		return __r0
	}
	iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations = func(steamNetworkingUtils uintptr, location1 uintptr, location2 uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations, uintptr(steamNetworkingUtils), uintptr(location1), uintptr(location2))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamNetworkingUtils_EstimatePingTimeFromLocalHost = func(steamNetworkingUtils uintptr, remoteLocation uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_EstimatePingTimeFromLocalHost, uintptr(steamNetworkingUtils), uintptr(remoteLocation))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamNetworkingUtils_ConvertPingLocationToString = func(steamNetworkingUtils uintptr, location uintptr, pszBuf []byte, cchBufSize int32) {
		purego.SyscallN(addr_iSteamNetworkingUtils_ConvertPingLocationToString, uintptr(steamNetworkingUtils), uintptr(location), uintptr(unsafe.Pointer(unsafe.SliceData(pszBuf))), uintptr(cchBufSize))
		runtime.KeepAlive(pszBuf)
	}
	iSteamNetworkingUtils_ParsePingLocationString = func(steamNetworkingUtils uintptr, pszString string, result *SteamNetworkPingLocation) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_ParsePingLocationString, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(bytePtrFromString(pszString))), uintptr(unsafe.Pointer(result)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszString)
		runtime.KeepAlive(result)
		return __r0
	}
	iSteamNetworkingUtils_GetPingToDataCenter = func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID, pViaRelayPoP *SteamNetworkingPOPID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetPingToDataCenter, uintptr(steamNetworkingUtils), uintptr(popID), uintptr(unsafe.Pointer(pViaRelayPoP)))
		__r0 := int32(_r0)
		runtime.KeepAlive(pViaRelayPoP)
		return __r0
	}
	iSteamNetworkingUtils_GetDirectPingToPOP = func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetDirectPingToPOP, uintptr(steamNetworkingUtils), uintptr(popID))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamNetworkingUtils_GetPOPCount = func(steamNetworkingUtils uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetPOPCount, uintptr(steamNetworkingUtils))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamNetworkingUtils_GetPOPList = func(steamNetworkingUtils uintptr, list []SteamNetworkingPOPID, nListSz int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetPOPList, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(unsafe.SliceData(list))), uintptr(nListSz))
		__r0 := int32(_r0)
		runtime.KeepAlive(list)
		return __r0
	}
	iSteamNetworkingUtils_GetLocalTimestamp = func(steamNetworkingUtils uintptr) SteamNetworkingMicroseconds {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetLocalTimestamp, uintptr(steamNetworkingUtils))
		__r0 := SteamNetworkingMicroseconds(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetDebugOutputFunction = func(steamNetworkingUtils uintptr, eDetailLevel ESteamNetworkingSocketsDebugOutputType, pfnFunc FSteamNetworkingSocketsDebugOutput) {
		purego.SyscallN(addr_iSteamNetworkingUtils_SetDebugOutputFunction, uintptr(steamNetworkingUtils), uintptr(eDetailLevel), uintptr(pfnFunc))
	}
	iSteamNetworkingUtils_IsFakeIPv4 = func(steamNetworkingUtils uintptr, nIPv4 uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_IsFakeIPv4, uintptr(steamNetworkingUtils), uintptr(nIPv4))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_GetIPv4FakeIPType = func(steamNetworkingUtils uintptr, nIPv4 uint32) ESteamNetworkingFakeIPType {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetIPv4FakeIPType, uintptr(steamNetworkingUtils), uintptr(nIPv4))
		__r0 := ESteamNetworkingFakeIPType(_r0)
		return __r0
	}
	iSteamNetworkingUtils_GetRealIdentityForFakeIP = func(steamNetworkingUtils uintptr, fakeIP uintptr, pOutRealIdentity *SteamNetworkingIdentity) EResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetRealIdentityForFakeIP, uintptr(steamNetworkingUtils), uintptr(fakeIP), uintptr(unsafe.Pointer(pOutRealIdentity)))
		__r0 := EResult(_r0)
		runtime.KeepAlive(pOutRealIdentity)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalConfigValueInt32 = func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalConfigValueInt32, uintptr(steamNetworkingUtils), uintptr(eValue), uintptr(val))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalConfigValueString = func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalConfigValueString, uintptr(steamNetworkingUtils), uintptr(eValue), uintptr(unsafe.Pointer(bytePtrFromString(val))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(val)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalConfigValuePtr = func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val []byte) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalConfigValuePtr, uintptr(steamNetworkingUtils), uintptr(eValue), uintptr(unsafe.Pointer(unsafe.SliceData(val))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(val)
		return __r0
	}
	iSteamNetworkingUtils_SetConnectionConfigValueInt32 = func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetConnectionConfigValueInt32, uintptr(steamNetworkingUtils), uintptr(hConn), uintptr(eValue), uintptr(val))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetConnectionConfigValueString = func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetConnectionConfigValueString, uintptr(steamNetworkingUtils), uintptr(hConn), uintptr(eValue), uintptr(unsafe.Pointer(bytePtrFromString(val))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(val)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged = func(steamNetworkingUtils uintptr, fnCallback FnSteamNetConnectionStatusChanged) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamNetConnectionStatusChangedCallback](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged = func(steamNetworkingUtils uintptr, fnCallback FnSteamNetAuthenticationStatusChanged) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamNetAuthenticationStatus](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged = func(steamNetworkingUtils uintptr, fnCallback FnSteamRelayNetworkStatusChanged) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamRelayNetworkStatus](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult = func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingFakeIPResult) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamNetworkingFakeIPResult](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest = func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionRequest) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamNetworkingMessagesSessionRequest](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed = func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionFailed) bool {
		wrapper := func(callbackData uintptr) {
			convertedCallback := uintptrToStruct[SteamNetworkingMessagesSessionFailed](callbackData)
			fnCallback(convertedCallback)
		}
		cb := purego.NewCallback(wrapper)
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed, uintptr(steamNetworkingUtils), uintptr(cb))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SetConfigValue = func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetConfigValue, uintptr(steamNetworkingUtils), uintptr(eValue), uintptr(eScopeType), uintptr(scopeObj), uintptr(eDataType), uintptr(unsafe.Pointer(unsafe.SliceData(pArg))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pArg)
		return __r0
	}
	iSteamNetworkingUtils_SetConfigValueStruct = func(steamNetworkingUtils uintptr, opt uintptr, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SetConfigValueStruct, uintptr(steamNetworkingUtils), uintptr(opt), uintptr(eScopeType), uintptr(scopeObj))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamNetworkingUtils_GetConfigValue = func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, pOutDataType *ESteamNetworkingConfigDataType, pResult []byte, cbResult *size) ESteamNetworkingGetConfigValueResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_GetConfigValue, uintptr(steamNetworkingUtils), uintptr(eValue), uintptr(eScopeType), uintptr(scopeObj), uintptr(unsafe.Pointer(pOutDataType)), uintptr(unsafe.Pointer(unsafe.SliceData(pResult))), uintptr(unsafe.Pointer(cbResult)))
		__r0 := ESteamNetworkingGetConfigValueResult(_r0)
		runtime.KeepAlive(pOutDataType)
		runtime.KeepAlive(pResult)
		runtime.KeepAlive(cbResult)
		return __r0
	}
	iSteamNetworkingUtils_IterateGenericEditableConfigValues = func(steamNetworkingUtils uintptr, eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_IterateGenericEditableConfigValues, uintptr(steamNetworkingUtils), uintptr(eCurrent), boolToUintptr(bEnumerateDevVars))
		__r0 := ESteamNetworkingConfigValue(_r0)
		return __r0
	}
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString = func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr, buf []byte, cbBuf uint32, bWithPort bool) {
		purego.SyscallN(addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(addr)), uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(cbBuf), boolToUintptr(bWithPort))
		runtime.KeepAlive(addr)
		runtime.KeepAlive(buf)
	}
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString = func(steamNetworkingUtils uintptr, pAddr *SteamNetworkingIPAddr, pszStr string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(pAddr)), uintptr(unsafe.Pointer(bytePtrFromString(pszStr))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pAddr)
		runtime.KeepAlive(pszStr)
		return __r0
	}
	iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType = func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(addr)))
		__r0 := ESteamNetworkingFakeIPType(_r0)
		runtime.KeepAlive(addr)
		return __r0
	}
	iSteamNetworkingUtils_SteamNetworkingIdentity_ToString = func(steamNetworkingUtils uintptr, identity *SteamNetworkingIdentity, buf []byte, cbBuf uint32) {
		purego.SyscallN(addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ToString, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(identity)), uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(cbBuf))
		runtime.KeepAlive(identity)
		runtime.KeepAlive(buf)
	}
	iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString = func(steamNetworkingUtils uintptr, pIdentity *SteamNetworkingIdentity, pszStr string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString, uintptr(steamNetworkingUtils), uintptr(unsafe.Pointer(pIdentity)), uintptr(unsafe.Pointer(bytePtrFromString(pszStr))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pIdentity)
		runtime.KeepAlive(pszStr)
		return __r0
	}

	purego.RegisterLibFunc(&iSteamNetworkingUtils_GetLocalPingLocation, steamAPILib, flatAPI_ISteamNetworkingUtils_GetLocalPingLocation)
	purego.RegisterLibFunc(&iSteamNetworkingUtils_CheckPingDataUpToDate, steamAPILib, flatAPI_ISteamNetworkingUtils_CheckPingDataUpToDate)
	purego.RegisterLibFunc(&iSteamNetworkingUtils_SetGlobalConfigValueFloat, steamAPILib, flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat)
	purego.RegisterLibFunc(&iSteamNetworkingUtils_SetConnectionConfigValueFloat, steamAPILib, flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat)
}

var (
	addr_steamParties_get                             uintptr
	addr_iSteamParties_GetNumActiveBeacons            uintptr
	addr_iSteamParties_GetBeaconByIndex               uintptr
	addr_iSteamParties_GetBeaconDetails               uintptr
	addr_iSteamParties_JoinParty                      uintptr
	addr_iSteamParties_GetNumAvailableBeaconLocations uintptr
	addr_iSteamParties_GetAvailableBeaconLocations    uintptr
	addr_iSteamParties_CreateBeacon                   uintptr
	addr_iSteamParties_OnReservationCompleted         uintptr
	addr_iSteamParties_CancelReservation              uintptr
	addr_iSteamParties_ChangeNumOpenSlots             uintptr
	addr_iSteamParties_DestroyBeacon                  uintptr
	addr_iSteamParties_GetBeaconLocationData          uintptr
)

func initParties() {
	var err error
	addr_steamParties_get, err = openSymbol(steamAPILib, flatAPI_SteamParties)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamParties)
	}
	addr_iSteamParties_GetNumActiveBeacons, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetNumActiveBeacons")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetNumActiveBeacons")
	}
	addr_iSteamParties_GetBeaconByIndex, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetBeaconByIndex")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetBeaconByIndex")
	}
	addr_iSteamParties_GetBeaconDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetBeaconDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetBeaconDetails")
	}
	addr_iSteamParties_JoinParty, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_JoinParty")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_JoinParty")
	}
	addr_iSteamParties_GetNumAvailableBeaconLocations, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetNumAvailableBeaconLocations")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetNumAvailableBeaconLocations")
	}
	addr_iSteamParties_GetAvailableBeaconLocations, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetAvailableBeaconLocations")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetAvailableBeaconLocations")
	}
	addr_iSteamParties_CreateBeacon, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_CreateBeacon")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_CreateBeacon")
	}
	addr_iSteamParties_OnReservationCompleted, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_OnReservationCompleted")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_OnReservationCompleted")
	}
	addr_iSteamParties_CancelReservation, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_CancelReservation")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_CancelReservation")
	}
	addr_iSteamParties_ChangeNumOpenSlots, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_ChangeNumOpenSlots")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_ChangeNumOpenSlots")
	}
	addr_iSteamParties_DestroyBeacon, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_DestroyBeacon")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_DestroyBeacon")
	}
	addr_iSteamParties_GetBeaconLocationData, err = openSymbol(steamAPILib, "SteamAPI_ISteamParties_GetBeaconLocationData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamParties_GetBeaconLocationData")
	}

	steamParties_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamParties_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamParties_GetNumActiveBeacons = func(steamParties uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetNumActiveBeacons, uintptr(steamParties))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamParties_GetBeaconByIndex = func(steamParties uintptr, unIndex uint32) PartyBeaconID {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetBeaconByIndex, uintptr(steamParties), uintptr(unIndex))
		__r0 := PartyBeaconID(_r0)
		return __r0
	}
	iSteamParties_GetBeaconDetails = func(steamParties uintptr, ulBeaconID PartyBeaconID, pSteamIDBeaconOwner *CSteamID, pLocation *SteamPartyBeaconLocation, pchMetadata []byte, cchMetadata int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetBeaconDetails, uintptr(steamParties), uintptr(ulBeaconID), uintptr(unsafe.Pointer(pSteamIDBeaconOwner)), uintptr(unsafe.Pointer(pLocation)), uintptr(unsafe.Pointer(unsafe.SliceData(pchMetadata))), uintptr(cchMetadata))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pSteamIDBeaconOwner)
		runtime.KeepAlive(pLocation)
		runtime.KeepAlive(pchMetadata)
		return __r0
	}
	iSteamParties_JoinParty = func(steamParties uintptr, ulBeaconID PartyBeaconID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_JoinParty, uintptr(steamParties), uintptr(ulBeaconID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamParties_GetNumAvailableBeaconLocations = func(steamParties uintptr, puNumLocations *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetNumAvailableBeaconLocations, uintptr(steamParties), uintptr(unsafe.Pointer(puNumLocations)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(puNumLocations)
		return __r0
	}
	iSteamParties_GetAvailableBeaconLocations = func(steamParties uintptr, pLocationList []SteamPartyBeaconLocation, uMaxNumLocations uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetAvailableBeaconLocations, uintptr(steamParties), uintptr(unsafe.Pointer(unsafe.SliceData(pLocationList))), uintptr(uMaxNumLocations))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pLocationList)
		return __r0
	}
	iSteamParties_CreateBeacon = func(steamParties uintptr, unOpenSlots uint32, pBeaconLocation uintptr, pchConnectString string, pchMetadata string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_CreateBeacon, uintptr(steamParties), uintptr(unOpenSlots), uintptr(pBeaconLocation), uintptr(unsafe.Pointer(bytePtrFromString(pchConnectString))), uintptr(unsafe.Pointer(bytePtrFromString(pchMetadata))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchConnectString)
		runtime.KeepAlive(pchMetadata)
		return __r0
	}
	iSteamParties_OnReservationCompleted = func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamParties_OnReservationCompleted, uintptr(steamParties), uintptr(ulBeacon), uintptr(userSteamID))
	}
	iSteamParties_CancelReservation = func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamParties_CancelReservation, uintptr(steamParties), uintptr(ulBeacon), uintptr(userSteamID))
	}
	iSteamParties_ChangeNumOpenSlots = func(steamParties uintptr, ulBeacon PartyBeaconID, unOpenSlots uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_ChangeNumOpenSlots, uintptr(steamParties), uintptr(ulBeacon), uintptr(unOpenSlots))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamParties_DestroyBeacon = func(steamParties uintptr, ulBeacon PartyBeaconID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_DestroyBeacon, uintptr(steamParties), uintptr(ulBeacon))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamParties_GetBeaconLocationData = func(steamParties uintptr, BeaconLocation uintptr, eData ESteamPartyBeaconLocationData, pchDataStringOut []byte, cchDataStringOut int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamParties_GetBeaconLocationData, uintptr(steamParties), uintptr(BeaconLocation), uintptr(eData), uintptr(unsafe.Pointer(unsafe.SliceData(pchDataStringOut))), uintptr(cchDataStringOut))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchDataStringOut)
		return __r0
	}
}

var (
	addr_steamRemotePlay_get                            uintptr
	addr_iSteamRemotePlay_GetSessionCount               uintptr
	addr_iSteamRemotePlay_GetSessionID                  uintptr
	addr_iSteamRemotePlay_GetSessionSteamID             uintptr
	addr_iSteamRemotePlay_GetSessionClientName          uintptr
	addr_iSteamRemotePlay_GetSessionClientFormFactor    uintptr
	addr_iSteamRemotePlay_BGetSessionClientResolution   uintptr
	addr_iSteamRemotePlay_BStartRemotePlayTogether      uintptr
	addr_iSteamRemotePlay_BSendRemotePlayTogetherInvite uintptr
)

func initRemotePlay() {
	var err error
	addr_steamRemotePlay_get, err = openSymbol(steamAPILib, flatAPI_SteamRemotePlay)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamRemotePlay)
	}
	addr_iSteamRemotePlay_GetSessionCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_GetSessionCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_GetSessionCount")
	}
	addr_iSteamRemotePlay_GetSessionID, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_GetSessionID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_GetSessionID")
	}
	addr_iSteamRemotePlay_GetSessionSteamID, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_GetSessionSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_GetSessionSteamID")
	}
	addr_iSteamRemotePlay_GetSessionClientName, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_GetSessionClientName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_GetSessionClientName")
	}
	addr_iSteamRemotePlay_GetSessionClientFormFactor, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_GetSessionClientFormFactor")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_GetSessionClientFormFactor")
	}
	addr_iSteamRemotePlay_BGetSessionClientResolution, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_BGetSessionClientResolution")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_BGetSessionClientResolution")
	}
	addr_iSteamRemotePlay_BStartRemotePlayTogether, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_BStartRemotePlayTogether")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_BStartRemotePlayTogether")
	}
	addr_iSteamRemotePlay_BSendRemotePlayTogetherInvite, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemotePlay_BSendRemotePlayTogetherInvite")
	}

	steamRemotePlay_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamRemotePlay_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamRemotePlay_GetSessionCount = func(steamRemotePlay uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_GetSessionCount, uintptr(steamRemotePlay))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamRemotePlay_GetSessionID = func(steamRemotePlay uintptr, iSessionIndex int32) RemotePlaySessionID {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_GetSessionID, uintptr(steamRemotePlay), uintptr(iSessionIndex))
		__r0 := RemotePlaySessionID(_r0)
		return __r0
	}
	iSteamRemotePlay_GetSessionSteamID = func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) Uint64SteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_GetSessionSteamID, uintptr(steamRemotePlay), uintptr(unSessionID))
		__r0 := Uint64SteamID(_r0)
		return __r0
	}
	iSteamRemotePlay_GetSessionClientName = func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_GetSessionClientName, uintptr(steamRemotePlay), uintptr(unSessionID))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamRemotePlay_GetSessionClientFormFactor = func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID) ESteamDeviceFormFactor {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_GetSessionClientFormFactor, uintptr(steamRemotePlay), uintptr(unSessionID))
		__r0 := ESteamDeviceFormFactor(_r0)
		return __r0
	}
	iSteamRemotePlay_BGetSessionClientResolution = func(steamRemotePlay uintptr, unSessionID RemotePlaySessionID, pnResolutionX *int32, pnResolutionY *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_BGetSessionClientResolution, uintptr(steamRemotePlay), uintptr(unSessionID), uintptr(unsafe.Pointer(pnResolutionX)), uintptr(unsafe.Pointer(pnResolutionY)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnResolutionX)
		runtime.KeepAlive(pnResolutionY)
		return __r0
	}
	iSteamRemotePlay_BStartRemotePlayTogether = func(steamRemotePlay uintptr, bShowOverlay bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_BStartRemotePlayTogether, uintptr(steamRemotePlay), boolToUintptr(bShowOverlay))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemotePlay_BSendRemotePlayTogetherInvite = func(steamRemotePlay uintptr, steamIDFriend Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemotePlay_BSendRemotePlayTogetherInvite, uintptr(steamRemotePlay), uintptr(steamIDFriend))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}

var (
	addr_steamRemoteStorage_get                                      uintptr
	addr_iSteamRemoteStorage_FileWrite                               uintptr
	addr_iSteamRemoteStorage_FileRead                                uintptr
	addr_iSteamRemoteStorage_FileWriteAsync                          uintptr
	addr_iSteamRemoteStorage_FileReadAsync                           uintptr
	addr_iSteamRemoteStorage_FileReadAsyncComplete                   uintptr
	addr_iSteamRemoteStorage_FileForget                              uintptr
	addr_iSteamRemoteStorage_FileDelete                              uintptr
	addr_iSteamRemoteStorage_FileShare                               uintptr
	addr_iSteamRemoteStorage_SetSyncPlatforms                        uintptr
	addr_iSteamRemoteStorage_FileWriteStreamOpen                     uintptr
	addr_iSteamRemoteStorage_FileWriteStreamWriteChunk               uintptr
	addr_iSteamRemoteStorage_FileWriteStreamClose                    uintptr
	addr_iSteamRemoteStorage_FileWriteStreamCancel                   uintptr
	addr_iSteamRemoteStorage_FileExists                              uintptr
	addr_iSteamRemoteStorage_FilePersisted                           uintptr
	addr_iSteamRemoteStorage_GetFileSize                             uintptr
	addr_iSteamRemoteStorage_GetFileTimestamp                        uintptr
	addr_iSteamRemoteStorage_GetSyncPlatforms                        uintptr
	addr_iSteamRemoteStorage_GetFileCount                            uintptr
	addr_iSteamRemoteStorage_GetFileNameAndSize                      uintptr
	addr_iSteamRemoteStorage_GetQuota                                uintptr
	addr_iSteamRemoteStorage_IsCloudEnabledForAccount                uintptr
	addr_iSteamRemoteStorage_IsCloudEnabledForApp                    uintptr
	addr_iSteamRemoteStorage_SetCloudEnabledForApp                   uintptr
	addr_iSteamRemoteStorage_UGCDownload                             uintptr
	addr_iSteamRemoteStorage_GetUGCDownloadProgress                  uintptr
	addr_iSteamRemoteStorage_GetUGCDetails                           uintptr
	addr_iSteamRemoteStorage_UGCRead                                 uintptr
	addr_iSteamRemoteStorage_GetCachedUGCCount                       uintptr
	addr_iSteamRemoteStorage_GetCachedUGCHandle                      uintptr
	addr_iSteamRemoteStorage_PublishWorkshopFile                     uintptr
	addr_iSteamRemoteStorage_CreatePublishedFileUpdateRequest        uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileFile                 uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFilePreviewFile          uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileTitle                uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileDescription          uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileVisibility           uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileTags                 uintptr
	addr_iSteamRemoteStorage_CommitPublishedFileUpdate               uintptr
	addr_iSteamRemoteStorage_GetPublishedFileDetails                 uintptr
	addr_iSteamRemoteStorage_DeletePublishedFile                     uintptr
	addr_iSteamRemoteStorage_EnumerateUserPublishedFiles             uintptr
	addr_iSteamRemoteStorage_SubscribePublishedFile                  uintptr
	addr_iSteamRemoteStorage_EnumerateUserSubscribedFiles            uintptr
	addr_iSteamRemoteStorage_UnsubscribePublishedFile                uintptr
	addr_iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription uintptr
	addr_iSteamRemoteStorage_GetPublishedItemVoteDetails             uintptr
	addr_iSteamRemoteStorage_UpdateUserPublishedItemVote             uintptr
	addr_iSteamRemoteStorage_GetUserPublishedItemVoteDetails         uintptr
	addr_iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles        uintptr
	addr_iSteamRemoteStorage_PublishVideo                            uintptr
	addr_iSteamRemoteStorage_SetUserPublishedFileAction              uintptr
	addr_iSteamRemoteStorage_EnumeratePublishedFilesByUserAction     uintptr
	addr_iSteamRemoteStorage_EnumeratePublishedWorkshopFiles         uintptr
	addr_iSteamRemoteStorage_UGCDownloadToLocation                   uintptr
	addr_iSteamRemoteStorage_GetLocalFileChangeCount                 uintptr
	addr_iSteamRemoteStorage_GetLocalFileChange                      uintptr
	addr_iSteamRemoteStorage_BeginFileWriteBatch                     uintptr
	addr_iSteamRemoteStorage_EndFileWriteBatch                       uintptr
)

func initRemoteStorage() {
	var err error
	addr_steamRemoteStorage_get, err = openSymbol(steamAPILib, flatAPI_SteamRemoteStorage)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamRemoteStorage)
	}
	addr_iSteamRemoteStorage_FileWrite, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWrite")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWrite")
	}
	addr_iSteamRemoteStorage_FileRead, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileRead")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileRead")
	}
	addr_iSteamRemoteStorage_FileWriteAsync, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWriteAsync")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWriteAsync")
	}
	addr_iSteamRemoteStorage_FileReadAsync, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileReadAsync")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileReadAsync")
	}
	addr_iSteamRemoteStorage_FileReadAsyncComplete, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileReadAsyncComplete")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileReadAsyncComplete")
	}
	addr_iSteamRemoteStorage_FileForget, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileForget")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileForget")
	}
	addr_iSteamRemoteStorage_FileDelete, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileDelete")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileDelete")
	}
	addr_iSteamRemoteStorage_FileShare, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileShare")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileShare")
	}
	addr_iSteamRemoteStorage_SetSyncPlatforms, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_SetSyncPlatforms")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_SetSyncPlatforms")
	}
	addr_iSteamRemoteStorage_FileWriteStreamOpen, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWriteStreamOpen")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWriteStreamOpen")
	}
	addr_iSteamRemoteStorage_FileWriteStreamWriteChunk, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWriteStreamWriteChunk")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWriteStreamWriteChunk")
	}
	addr_iSteamRemoteStorage_FileWriteStreamClose, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWriteStreamClose")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWriteStreamClose")
	}
	addr_iSteamRemoteStorage_FileWriteStreamCancel, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileWriteStreamCancel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileWriteStreamCancel")
	}
	addr_iSteamRemoteStorage_FileExists, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FileExists")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FileExists")
	}
	addr_iSteamRemoteStorage_FilePersisted, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_FilePersisted")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_FilePersisted")
	}
	addr_iSteamRemoteStorage_GetFileSize, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetFileSize")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetFileSize")
	}
	addr_iSteamRemoteStorage_GetFileTimestamp, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetFileTimestamp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetFileTimestamp")
	}
	addr_iSteamRemoteStorage_GetSyncPlatforms, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetSyncPlatforms")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetSyncPlatforms")
	}
	addr_iSteamRemoteStorage_GetFileCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetFileCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetFileCount")
	}
	addr_iSteamRemoteStorage_GetFileNameAndSize, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetFileNameAndSize")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetFileNameAndSize")
	}
	addr_iSteamRemoteStorage_GetQuota, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetQuota")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetQuota")
	}
	addr_iSteamRemoteStorage_IsCloudEnabledForAccount, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_IsCloudEnabledForAccount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_IsCloudEnabledForAccount")
	}
	addr_iSteamRemoteStorage_IsCloudEnabledForApp, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_IsCloudEnabledForApp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_IsCloudEnabledForApp")
	}
	addr_iSteamRemoteStorage_SetCloudEnabledForApp, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_SetCloudEnabledForApp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_SetCloudEnabledForApp")
	}
	addr_iSteamRemoteStorage_UGCDownload, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UGCDownload")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UGCDownload")
	}
	addr_iSteamRemoteStorage_GetUGCDownloadProgress, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetUGCDownloadProgress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetUGCDownloadProgress")
	}
	addr_iSteamRemoteStorage_GetUGCDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetUGCDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetUGCDetails")
	}
	addr_iSteamRemoteStorage_UGCRead, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UGCRead")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UGCRead")
	}
	addr_iSteamRemoteStorage_GetCachedUGCCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetCachedUGCCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetCachedUGCCount")
	}
	addr_iSteamRemoteStorage_GetCachedUGCHandle, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetCachedUGCHandle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetCachedUGCHandle")
	}
	addr_iSteamRemoteStorage_PublishWorkshopFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_PublishWorkshopFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_PublishWorkshopFile")
	}
	addr_iSteamRemoteStorage_CreatePublishedFileUpdateRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_CreatePublishedFileUpdateRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_CreatePublishedFileUpdateRequest")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileFile")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFilePreviewFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFilePreviewFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFilePreviewFile")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileTitle, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTitle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTitle")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileDescription, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileDescription")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileDescription")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileVisibility, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileVisibility")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileVisibility")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileTags")
	}
	addr_iSteamRemoteStorage_CommitPublishedFileUpdate, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_CommitPublishedFileUpdate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_CommitPublishedFileUpdate")
	}
	addr_iSteamRemoteStorage_GetPublishedFileDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetPublishedFileDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetPublishedFileDetails")
	}
	addr_iSteamRemoteStorage_DeletePublishedFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_DeletePublishedFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_DeletePublishedFile")
	}
	addr_iSteamRemoteStorage_EnumerateUserPublishedFiles, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EnumerateUserPublishedFiles")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EnumerateUserPublishedFiles")
	}
	addr_iSteamRemoteStorage_SubscribePublishedFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_SubscribePublishedFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_SubscribePublishedFile")
	}
	addr_iSteamRemoteStorage_EnumerateUserSubscribedFiles, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EnumerateUserSubscribedFiles")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EnumerateUserSubscribedFiles")
	}
	addr_iSteamRemoteStorage_UnsubscribePublishedFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UnsubscribePublishedFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UnsubscribePublishedFile")
	}
	addr_iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdatePublishedFileSetChangeDescription")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdatePublishedFileSetChangeDescription")
	}
	addr_iSteamRemoteStorage_GetPublishedItemVoteDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetPublishedItemVoteDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetPublishedItemVoteDetails")
	}
	addr_iSteamRemoteStorage_UpdateUserPublishedItemVote, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UpdateUserPublishedItemVote")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UpdateUserPublishedItemVote")
	}
	addr_iSteamRemoteStorage_GetUserPublishedItemVoteDetails, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetUserPublishedItemVoteDetails")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetUserPublishedItemVoteDetails")
	}
	addr_iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EnumerateUserSharedWorkshopFiles")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EnumerateUserSharedWorkshopFiles")
	}
	addr_iSteamRemoteStorage_PublishVideo, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_PublishVideo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_PublishVideo")
	}
	addr_iSteamRemoteStorage_SetUserPublishedFileAction, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_SetUserPublishedFileAction")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_SetUserPublishedFileAction")
	}
	addr_iSteamRemoteStorage_EnumeratePublishedFilesByUserAction, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EnumeratePublishedFilesByUserAction")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EnumeratePublishedFilesByUserAction")
	}
	addr_iSteamRemoteStorage_EnumeratePublishedWorkshopFiles, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EnumeratePublishedWorkshopFiles")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EnumeratePublishedWorkshopFiles")
	}
	addr_iSteamRemoteStorage_UGCDownloadToLocation, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_UGCDownloadToLocation")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_UGCDownloadToLocation")
	}
	addr_iSteamRemoteStorage_GetLocalFileChangeCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetLocalFileChangeCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetLocalFileChangeCount")
	}
	addr_iSteamRemoteStorage_GetLocalFileChange, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_GetLocalFileChange")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_GetLocalFileChange")
	}
	addr_iSteamRemoteStorage_BeginFileWriteBatch, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_BeginFileWriteBatch")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_BeginFileWriteBatch")
	}
	addr_iSteamRemoteStorage_EndFileWriteBatch, err = openSymbol(steamAPILib, "SteamAPI_ISteamRemoteStorage_EndFileWriteBatch")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamRemoteStorage_EndFileWriteBatch")
	}

	steamRemoteStorage_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamRemoteStorage_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_FileWrite = func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubData int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWrite, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubData))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		runtime.KeepAlive(pvData)
		return __r0
	}
	iSteamRemoteStorage_FileRead = func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubDataToRead int32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileRead, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubDataToRead))
		__r0 := int32(_r0)
		runtime.KeepAlive(pchFile)
		runtime.KeepAlive(pvData)
		return __r0
	}
	iSteamRemoteStorage_FileWriteAsync = func(steamRemoteStorage uintptr, pchFile string, pvData []byte, cubData uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWriteAsync, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubData))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchFile)
		runtime.KeepAlive(pvData)
		return __r0
	}
	iSteamRemoteStorage_FileReadAsync = func(steamRemoteStorage uintptr, pchFile string, nOffset uint32, cubToRead uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileReadAsync, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(nOffset), uintptr(cubToRead))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FileReadAsyncComplete = func(steamRemoteStorage uintptr, hReadCall SteamAPICall, pvBuffer []byte, cubToRead uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileReadAsyncComplete, uintptr(steamRemoteStorage), uintptr(hReadCall), uintptr(unsafe.Pointer(unsafe.SliceData(pvBuffer))), uintptr(cubToRead))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pvBuffer)
		return __r0
	}
	iSteamRemoteStorage_FileForget = func(steamRemoteStorage uintptr, pchFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileForget, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FileDelete = func(steamRemoteStorage uintptr, pchFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileDelete, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FileShare = func(steamRemoteStorage uintptr, pchFile string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileShare, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_SetSyncPlatforms = func(steamRemoteStorage uintptr, pchFile string, eRemoteStoragePlatform ERemoteStoragePlatform) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_SetSyncPlatforms, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(eRemoteStoragePlatform))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FileWriteStreamOpen = func(steamRemoteStorage uintptr, pchFile string) UGCFileWriteStreamHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWriteStreamOpen, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := UGCFileWriteStreamHandle(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FileWriteStreamWriteChunk = func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle, pvData []byte, cubData int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWriteStreamWriteChunk, uintptr(steamRemoteStorage), uintptr(writeHandle), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubData))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pvData)
		return __r0
	}
	iSteamRemoteStorage_FileWriteStreamClose = func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWriteStreamClose, uintptr(steamRemoteStorage), uintptr(writeHandle))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_FileWriteStreamCancel = func(steamRemoteStorage uintptr, writeHandle UGCFileWriteStreamHandle) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileWriteStreamCancel, uintptr(steamRemoteStorage), uintptr(writeHandle))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_FileExists = func(steamRemoteStorage uintptr, pchFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FileExists, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_FilePersisted = func(steamRemoteStorage uintptr, pchFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_FilePersisted, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_GetFileSize = func(steamRemoteStorage uintptr, pchFile string) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetFileSize, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := int32(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_GetFileTimestamp = func(steamRemoteStorage uintptr, pchFile string) int64 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetFileTimestamp, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := int64(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_GetSyncPlatforms = func(steamRemoteStorage uintptr, pchFile string) ERemoteStoragePlatform {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetSyncPlatforms, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := ERemoteStoragePlatform(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_GetFileCount = func(steamRemoteStorage uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetFileCount, uintptr(steamRemoteStorage))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamRemoteStorage_GetQuota = func(steamRemoteStorage uintptr, pnTotalBytes *uint64, puAvailableBytes *uint64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetQuota, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(pnTotalBytes)), uintptr(unsafe.Pointer(puAvailableBytes)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnTotalBytes)
		runtime.KeepAlive(puAvailableBytes)
		return __r0
	}
	iSteamRemoteStorage_IsCloudEnabledForAccount = func(steamRemoteStorage uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_IsCloudEnabledForAccount, uintptr(steamRemoteStorage))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_IsCloudEnabledForApp = func(steamRemoteStorage uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_IsCloudEnabledForApp, uintptr(steamRemoteStorage))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_SetCloudEnabledForApp = func(steamRemoteStorage uintptr, bEnabled bool) {
		purego.SyscallN(addr_iSteamRemoteStorage_SetCloudEnabledForApp, uintptr(steamRemoteStorage), boolToUintptr(bEnabled))
	}
	iSteamRemoteStorage_UGCDownload = func(steamRemoteStorage uintptr, hContent UGCHandle, unPriority uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UGCDownload, uintptr(steamRemoteStorage), uintptr(hContent), uintptr(unPriority))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_GetUGCDownloadProgress = func(steamRemoteStorage uintptr, hContent UGCHandle, pnBytesDownloaded *int32, pnBytesExpected *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetUGCDownloadProgress, uintptr(steamRemoteStorage), uintptr(hContent), uintptr(unsafe.Pointer(pnBytesDownloaded)), uintptr(unsafe.Pointer(pnBytesExpected)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnBytesDownloaded)
		runtime.KeepAlive(pnBytesExpected)
		return __r0
	}
	iSteamRemoteStorage_GetUGCDetails = func(steamRemoteStorage uintptr, hContent UGCHandle, pnAppID *AppId_t, ppchName *[]byte, pnFileSizeInBytes *int32, pSteamIDOwner *CSteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetUGCDetails, uintptr(steamRemoteStorage), uintptr(hContent), uintptr(unsafe.Pointer(pnAppID)), uintptr(unsafe.Pointer(ppchName)), uintptr(unsafe.Pointer(pnFileSizeInBytes)), uintptr(unsafe.Pointer(pSteamIDOwner)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnAppID)
		runtime.KeepAlive(ppchName)
		runtime.KeepAlive(pnFileSizeInBytes)
		runtime.KeepAlive(pSteamIDOwner)
		return __r0
	}
	iSteamRemoteStorage_UGCRead = func(steamRemoteStorage uintptr, hContent UGCHandle, pvData []byte, cubDataToRead int32, cOffset uint32, eAction EUGCReadAction) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UGCRead, uintptr(steamRemoteStorage), uintptr(hContent), uintptr(unsafe.Pointer(unsafe.SliceData(pvData))), uintptr(cubDataToRead), uintptr(cOffset), uintptr(eAction))
		__r0 := int32(_r0)
		runtime.KeepAlive(pvData)
		return __r0
	}
	iSteamRemoteStorage_GetCachedUGCCount = func(steamRemoteStorage uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetCachedUGCCount, uintptr(steamRemoteStorage))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamRemoteStorage_GetCachedUGCHandle = func(steamRemoteStorage uintptr, iCachedContent int32) UGCHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetCachedUGCHandle, uintptr(steamRemoteStorage), uintptr(iCachedContent))
		__r0 := UGCHandle(_r0)
		return __r0
	}
	iSteamRemoteStorage_PublishWorkshopFile = func(steamRemoteStorage uintptr, pchFile string, pchPreviewFile string, nConsumerAppId AppId_t, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr, eWorkshopFileType EWorkshopFileType) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_PublishWorkshopFile, uintptr(steamRemoteStorage), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))), uintptr(unsafe.Pointer(bytePtrFromString(pchPreviewFile))), uintptr(nConsumerAppId), uintptr(unsafe.Pointer(bytePtrFromString(pchTitle))), uintptr(unsafe.Pointer(bytePtrFromString(pchDescription))), uintptr(eVisibility), uintptr(pTags), uintptr(eWorkshopFileType))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchFile)
		runtime.KeepAlive(pchPreviewFile)
		runtime.KeepAlive(pchTitle)
		runtime.KeepAlive(pchDescription)
		return __r0
	}
	iSteamRemoteStorage_CreatePublishedFileUpdateRequest = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) PublishedFileUpdateHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_CreatePublishedFileUpdateRequest, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := PublishedFileUpdateHandle(_r0)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileFile = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileFile, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(unsafe.Pointer(bytePtrFromString(pchFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchFile)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFilePreviewFile = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchPreviewFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFilePreviewFile, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(unsafe.Pointer(bytePtrFromString(pchPreviewFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchPreviewFile)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileTitle = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchTitle string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileTitle, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(unsafe.Pointer(bytePtrFromString(pchTitle))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchTitle)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileDescription = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchDescription string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileDescription, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(unsafe.Pointer(bytePtrFromString(pchDescription))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchDescription)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileVisibility = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileVisibility, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(eVisibility))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileTags = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pTags uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileTags, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(pTags))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_CommitPublishedFileUpdate = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_CommitPublishedFileUpdate, uintptr(steamRemoteStorage), uintptr(updateHandle))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_GetPublishedFileDetails = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, unMaxSecondsOld uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetPublishedFileDetails, uintptr(steamRemoteStorage), uintptr(unPublishedFileId), uintptr(unMaxSecondsOld))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_DeletePublishedFile = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_DeletePublishedFile, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_EnumerateUserPublishedFiles = func(steamRemoteStorage uintptr, unStartIndex uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EnumerateUserPublishedFiles, uintptr(steamRemoteStorage), uintptr(unStartIndex))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_SubscribePublishedFile = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_SubscribePublishedFile, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_EnumerateUserSubscribedFiles = func(steamRemoteStorage uintptr, unStartIndex uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EnumerateUserSubscribedFiles, uintptr(steamRemoteStorage), uintptr(unStartIndex))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_UnsubscribePublishedFile = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UnsubscribePublishedFile, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription = func(steamRemoteStorage uintptr, updateHandle PublishedFileUpdateHandle, pchChangeDescription string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdatePublishedFileSetChangeDescription, uintptr(steamRemoteStorage), uintptr(updateHandle), uintptr(unsafe.Pointer(bytePtrFromString(pchChangeDescription))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchChangeDescription)
		return __r0
	}
	iSteamRemoteStorage_GetPublishedItemVoteDetails = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetPublishedItemVoteDetails, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_UpdateUserPublishedItemVote = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, bVoteUp bool) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UpdateUserPublishedItemVote, uintptr(steamRemoteStorage), uintptr(unPublishedFileId), boolToUintptr(bVoteUp))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_GetUserPublishedItemVoteDetails = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetUserPublishedItemVoteDetails, uintptr(steamRemoteStorage), uintptr(unPublishedFileId))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles = func(steamRemoteStorage uintptr, steamId Uint64SteamID, unStartIndex uint32, pRequiredTags uintptr, pExcludedTags uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EnumerateUserSharedWorkshopFiles, uintptr(steamRemoteStorage), uintptr(steamId), uintptr(unStartIndex), uintptr(pRequiredTags), uintptr(pExcludedTags))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_PublishVideo = func(steamRemoteStorage uintptr, eVideoProvider EWorkshopVideoProvider, pchVideoAccount string, pchVideoIdentifier string, pchPreviewFile string, nConsumerAppId AppId_t, pchTitle string, pchDescription string, eVisibility ERemoteStoragePublishedFileVisibility, pTags uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_PublishVideo, uintptr(steamRemoteStorage), uintptr(eVideoProvider), uintptr(unsafe.Pointer(bytePtrFromString(pchVideoAccount))), uintptr(unsafe.Pointer(bytePtrFromString(pchVideoIdentifier))), uintptr(unsafe.Pointer(bytePtrFromString(pchPreviewFile))), uintptr(nConsumerAppId), uintptr(unsafe.Pointer(bytePtrFromString(pchTitle))), uintptr(unsafe.Pointer(bytePtrFromString(pchDescription))), uintptr(eVisibility), uintptr(pTags))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchVideoAccount)
		runtime.KeepAlive(pchVideoIdentifier)
		runtime.KeepAlive(pchPreviewFile)
		runtime.KeepAlive(pchTitle)
		runtime.KeepAlive(pchDescription)
		return __r0
	}
	iSteamRemoteStorage_SetUserPublishedFileAction = func(steamRemoteStorage uintptr, unPublishedFileId PublishedFileId, eAction EWorkshopFileAction) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_SetUserPublishedFileAction, uintptr(steamRemoteStorage), uintptr(unPublishedFileId), uintptr(eAction))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_EnumeratePublishedFilesByUserAction = func(steamRemoteStorage uintptr, eAction EWorkshopFileAction, unStartIndex uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EnumeratePublishedFilesByUserAction, uintptr(steamRemoteStorage), uintptr(eAction), uintptr(unStartIndex))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamRemoteStorage_EnumeratePublishedWorkshopFiles = func(steamRemoteStorage uintptr, eEnumerationType EWorkshopEnumerationType, unStartIndex uint32, unCount uint32, unDays uint32, pTags *SteamParamStringArray, pUserTags *SteamParamStringArray) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EnumeratePublishedWorkshopFiles, uintptr(steamRemoteStorage), uintptr(eEnumerationType), uintptr(unStartIndex), uintptr(unCount), uintptr(unDays), uintptr(unsafe.Pointer(pTags)), uintptr(unsafe.Pointer(pUserTags)))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pTags)
		runtime.KeepAlive(pUserTags)
		return __r0
	}
	iSteamRemoteStorage_UGCDownloadToLocation = func(steamRemoteStorage uintptr, hContent UGCHandle, pchLocation string, unPriority uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_UGCDownloadToLocation, uintptr(steamRemoteStorage), uintptr(hContent), uintptr(unsafe.Pointer(bytePtrFromString(pchLocation))), uintptr(unPriority))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchLocation)
		return __r0
	}
	iSteamRemoteStorage_GetLocalFileChangeCount = func(steamRemoteStorage uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_GetLocalFileChangeCount, uintptr(steamRemoteStorage))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamRemoteStorage_BeginFileWriteBatch = func(steamRemoteStorage uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_BeginFileWriteBatch, uintptr(steamRemoteStorage))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamRemoteStorage_EndFileWriteBatch = func(steamRemoteStorage uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamRemoteStorage_EndFileWriteBatch, uintptr(steamRemoteStorage))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}

var (
	addr_steamScreenshots_get                       uintptr
	addr_iSteamScreenshots_WriteScreenshot          uintptr
	addr_iSteamScreenshots_AddScreenshotToLibrary   uintptr
	addr_iSteamScreenshots_TriggerScreenshot        uintptr
	addr_iSteamScreenshots_HookScreenshots          uintptr
	addr_iSteamScreenshots_SetLocation              uintptr
	addr_iSteamScreenshots_TagUser                  uintptr
	addr_iSteamScreenshots_TagPublishedFile         uintptr
	addr_iSteamScreenshots_IsScreenshotsHooked      uintptr
	addr_iSteamScreenshots_AddVRScreenshotToLibrary uintptr
)

func initScreenshots() {
	var err error
	addr_steamScreenshots_get, err = openSymbol(steamAPILib, flatAPI_SteamScreenshots)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamScreenshots)
	}
	addr_iSteamScreenshots_WriteScreenshot, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_WriteScreenshot")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_WriteScreenshot")
	}
	addr_iSteamScreenshots_AddScreenshotToLibrary, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_AddScreenshotToLibrary")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_AddScreenshotToLibrary")
	}
	addr_iSteamScreenshots_TriggerScreenshot, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_TriggerScreenshot")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_TriggerScreenshot")
	}
	addr_iSteamScreenshots_HookScreenshots, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_HookScreenshots")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_HookScreenshots")
	}
	addr_iSteamScreenshots_SetLocation, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_SetLocation")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_SetLocation")
	}
	addr_iSteamScreenshots_TagUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_TagUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_TagUser")
	}
	addr_iSteamScreenshots_TagPublishedFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_TagPublishedFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_TagPublishedFile")
	}
	addr_iSteamScreenshots_IsScreenshotsHooked, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_IsScreenshotsHooked")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_IsScreenshotsHooked")
	}
	addr_iSteamScreenshots_AddVRScreenshotToLibrary, err = openSymbol(steamAPILib, "SteamAPI_ISteamScreenshots_AddVRScreenshotToLibrary")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamScreenshots_AddVRScreenshotToLibrary")
	}

	steamScreenshots_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamScreenshots_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamScreenshots_WriteScreenshot = func(steamScreenshots uintptr, pubRGB []byte, cubRGB uint32, nWidth int32, nHeight int32) ScreenshotHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_WriteScreenshot, uintptr(steamScreenshots), uintptr(unsafe.Pointer(unsafe.SliceData(pubRGB))), uintptr(cubRGB), uintptr(nWidth), uintptr(nHeight))
		__r0 := ScreenshotHandle(_r0)
		runtime.KeepAlive(pubRGB)
		return __r0
	}
	iSteamScreenshots_AddScreenshotToLibrary = func(steamScreenshots uintptr, pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_AddScreenshotToLibrary, uintptr(steamScreenshots), uintptr(unsafe.Pointer(bytePtrFromString(pchFilename))), uintptr(unsafe.Pointer(bytePtrFromString(pchThumbnailFilename))), uintptr(nWidth), uintptr(nHeight))
		__r0 := ScreenshotHandle(_r0)
		runtime.KeepAlive(pchFilename)
		runtime.KeepAlive(pchThumbnailFilename)
		return __r0
	}
	iSteamScreenshots_TriggerScreenshot = func(steamScreenshots uintptr) {
		purego.SyscallN(addr_iSteamScreenshots_TriggerScreenshot, uintptr(steamScreenshots))
	}
	iSteamScreenshots_HookScreenshots = func(steamScreenshots uintptr, bHook bool) {
		purego.SyscallN(addr_iSteamScreenshots_HookScreenshots, uintptr(steamScreenshots), boolToUintptr(bHook))
	}
	iSteamScreenshots_SetLocation = func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, pchLocation string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_SetLocation, uintptr(steamScreenshots), uintptr(hScreenshot), uintptr(unsafe.Pointer(bytePtrFromString(pchLocation))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchLocation)
		return __r0
	}
	iSteamScreenshots_TagUser = func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_TagUser, uintptr(steamScreenshots), uintptr(hScreenshot), uintptr(steamID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamScreenshots_TagPublishedFile = func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_TagPublishedFile, uintptr(steamScreenshots), uintptr(hScreenshot), uintptr(unPublishedFileID))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamScreenshots_IsScreenshotsHooked = func(steamScreenshots uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_IsScreenshotsHooked, uintptr(steamScreenshots))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamScreenshots_AddVRScreenshotToLibrary = func(steamScreenshots uintptr, eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamScreenshots_AddVRScreenshotToLibrary, uintptr(steamScreenshots), uintptr(eType), uintptr(unsafe.Pointer(bytePtrFromString(pchFilename))), uintptr(unsafe.Pointer(bytePtrFromString(pchVRFilename))))
		__r0 := ScreenshotHandle(_r0)
		runtime.KeepAlive(pchFilename)
		runtime.KeepAlive(pchVRFilename)
		return __r0
	}
}

var (
	addr_steamUGC_get                                  uintptr
	addr_iSteamUGC_CreateQueryUserUGCRequest           uintptr
	addr_iSteamUGC_CreateQueryAllUGCRequestPage        uintptr
	addr_iSteamUGC_CreateQueryAllUGCRequestCursor      uintptr
	addr_iSteamUGC_CreateQueryUGCDetailsRequest        uintptr
	addr_iSteamUGC_SendQueryUGCRequest                 uintptr
	addr_iSteamUGC_GetQueryUGCResult                   uintptr
	addr_iSteamUGC_GetQueryUGCNumTags                  uintptr
	addr_iSteamUGC_GetQueryUGCTag                      uintptr
	addr_iSteamUGC_GetQueryUGCTagDisplayName           uintptr
	addr_iSteamUGC_GetQueryUGCPreviewURL               uintptr
	addr_iSteamUGC_GetQueryUGCMetadata                 uintptr
	addr_iSteamUGC_GetQueryUGCChildren                 uintptr
	addr_iSteamUGC_GetQueryUGCStatistic                uintptr
	addr_iSteamUGC_GetQueryUGCNumAdditionalPreviews    uintptr
	addr_iSteamUGC_GetQueryUGCAdditionalPreview        uintptr
	addr_iSteamUGC_GetQueryUGCNumKeyValueTags          uintptr
	addr_iSteamUGC_GetQueryUGCKeyValueTag              uintptr
	addr_iSteamUGC_GetQueryFirstUGCKeyValueTag         uintptr
	addr_iSteamUGC_GetNumSupportedGameVersions         uintptr
	addr_iSteamUGC_GetSupportedGameVersionData         uintptr
	addr_iSteamUGC_GetQueryUGCContentDescriptors       uintptr
	addr_iSteamUGC_ReleaseQueryUGCRequest              uintptr
	addr_iSteamUGC_AddRequiredTag                      uintptr
	addr_iSteamUGC_AddRequiredTagGroup                 uintptr
	addr_iSteamUGC_AddExcludedTag                      uintptr
	addr_iSteamUGC_SetReturnOnlyIDs                    uintptr
	addr_iSteamUGC_SetReturnKeyValueTags               uintptr
	addr_iSteamUGC_SetReturnLongDescription            uintptr
	addr_iSteamUGC_SetReturnMetadata                   uintptr
	addr_iSteamUGC_SetReturnChildren                   uintptr
	addr_iSteamUGC_SetReturnAdditionalPreviews         uintptr
	addr_iSteamUGC_SetReturnTotalOnly                  uintptr
	addr_iSteamUGC_SetReturnPlaytimeStats              uintptr
	addr_iSteamUGC_SetLanguage                         uintptr
	addr_iSteamUGC_SetAllowCachedResponse              uintptr
	addr_iSteamUGC_SetAdminQuery                       uintptr
	addr_iSteamUGC_SetCloudFileNameFilter              uintptr
	addr_iSteamUGC_SetMatchAnyTag                      uintptr
	addr_iSteamUGC_SetSearchText                       uintptr
	addr_iSteamUGC_SetRankedByTrendDays                uintptr
	addr_iSteamUGC_SetTimeCreatedDateRange             uintptr
	addr_iSteamUGC_SetTimeUpdatedDateRange             uintptr
	addr_iSteamUGC_AddRequiredKeyValueTag              uintptr
	addr_iSteamUGC_CreateItem                          uintptr
	addr_iSteamUGC_StartItemUpdate                     uintptr
	addr_iSteamUGC_SetItemTitle                        uintptr
	addr_iSteamUGC_SetItemDescription                  uintptr
	addr_iSteamUGC_SetItemUpdateLanguage               uintptr
	addr_iSteamUGC_SetItemMetadata                     uintptr
	addr_iSteamUGC_SetItemVisibility                   uintptr
	addr_iSteamUGC_SetItemTags                         uintptr
	addr_iSteamUGC_SetItemContent                      uintptr
	addr_iSteamUGC_SetItemPreview                      uintptr
	addr_iSteamUGC_SetAllowLegacyUpload                uintptr
	addr_iSteamUGC_RemoveAllItemKeyValueTags           uintptr
	addr_iSteamUGC_RemoveItemKeyValueTags              uintptr
	addr_iSteamUGC_AddItemKeyValueTag                  uintptr
	addr_iSteamUGC_AddItemPreviewFile                  uintptr
	addr_iSteamUGC_AddItemPreviewVideo                 uintptr
	addr_iSteamUGC_UpdateItemPreviewFile               uintptr
	addr_iSteamUGC_UpdateItemPreviewVideo              uintptr
	addr_iSteamUGC_RemoveItemPreview                   uintptr
	addr_iSteamUGC_AddContentDescriptor                uintptr
	addr_iSteamUGC_RemoveContentDescriptor             uintptr
	addr_iSteamUGC_SetRequiredGameVersions             uintptr
	addr_iSteamUGC_SubmitItemUpdate                    uintptr
	addr_iSteamUGC_GetItemUpdateProgress               uintptr
	addr_iSteamUGC_SetUserItemVote                     uintptr
	addr_iSteamUGC_GetUserItemVote                     uintptr
	addr_iSteamUGC_AddItemToFavorites                  uintptr
	addr_iSteamUGC_RemoveItemFromFavorites             uintptr
	addr_iSteamUGC_SubscribeItem                       uintptr
	addr_iSteamUGC_UnsubscribeItem                     uintptr
	addr_iSteamUGC_GetNumSubscribedItems               uintptr
	addr_iSteamUGC_GetSubscribedItems                  uintptr
	addr_iSteamUGC_GetItemState                        uintptr
	addr_iSteamUGC_GetItemInstallInfo                  uintptr
	addr_iSteamUGC_GetItemDownloadInfo                 uintptr
	addr_iSteamUGC_DownloadItem                        uintptr
	addr_iSteamUGC_BInitWorkshopForGameServer          uintptr
	addr_iSteamUGC_SuspendDownloads                    uintptr
	addr_iSteamUGC_StartPlaytimeTracking               uintptr
	addr_iSteamUGC_StopPlaytimeTracking                uintptr
	addr_iSteamUGC_StopPlaytimeTrackingForAllItems     uintptr
	addr_iSteamUGC_AddDependency                       uintptr
	addr_iSteamUGC_RemoveDependency                    uintptr
	addr_iSteamUGC_AddAppDependency                    uintptr
	addr_iSteamUGC_RemoveAppDependency                 uintptr
	addr_iSteamUGC_GetAppDependencies                  uintptr
	addr_iSteamUGC_DeleteItem                          uintptr
	addr_iSteamUGC_ShowWorkshopEULA                    uintptr
	addr_iSteamUGC_GetWorkshopEULAStatus               uintptr
	addr_iSteamUGC_GetUserContentDescriptorPreferences uintptr
)

func initUGC() {
	var err error
	addr_steamUGC_get, err = openSymbol(steamAPILib, flatAPI_SteamUGC)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamUGC)
	}
	addr_iSteamUGC_CreateQueryUserUGCRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_CreateQueryUserUGCRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_CreateQueryUserUGCRequest")
	}
	addr_iSteamUGC_CreateQueryAllUGCRequestPage, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_CreateQueryAllUGCRequestPage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_CreateQueryAllUGCRequestPage")
	}
	addr_iSteamUGC_CreateQueryAllUGCRequestCursor, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_CreateQueryAllUGCRequestCursor")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_CreateQueryAllUGCRequestCursor")
	}
	addr_iSteamUGC_CreateQueryUGCDetailsRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_CreateQueryUGCDetailsRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_CreateQueryUGCDetailsRequest")
	}
	addr_iSteamUGC_SendQueryUGCRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SendQueryUGCRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SendQueryUGCRequest")
	}
	addr_iSteamUGC_GetQueryUGCResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCResult")
	}
	addr_iSteamUGC_GetQueryUGCNumTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCNumTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCNumTags")
	}
	addr_iSteamUGC_GetQueryUGCTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCTag")
	}
	addr_iSteamUGC_GetQueryUGCTagDisplayName, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCTagDisplayName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCTagDisplayName")
	}
	addr_iSteamUGC_GetQueryUGCPreviewURL, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCPreviewURL")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCPreviewURL")
	}
	addr_iSteamUGC_GetQueryUGCMetadata, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCMetadata")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCMetadata")
	}
	addr_iSteamUGC_GetQueryUGCChildren, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCChildren")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCChildren")
	}
	addr_iSteamUGC_GetQueryUGCStatistic, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCStatistic")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCStatistic")
	}
	addr_iSteamUGC_GetQueryUGCNumAdditionalPreviews, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCNumAdditionalPreviews")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCNumAdditionalPreviews")
	}
	addr_iSteamUGC_GetQueryUGCAdditionalPreview, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCAdditionalPreview")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCAdditionalPreview")
	}
	addr_iSteamUGC_GetQueryUGCNumKeyValueTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCNumKeyValueTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCNumKeyValueTags")
	}
	addr_iSteamUGC_GetQueryUGCKeyValueTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCKeyValueTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCKeyValueTag")
	}
	addr_iSteamUGC_GetQueryFirstUGCKeyValueTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryFirstUGCKeyValueTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryFirstUGCKeyValueTag")
	}
	addr_iSteamUGC_GetNumSupportedGameVersions, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetNumSupportedGameVersions")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetNumSupportedGameVersions")
	}
	addr_iSteamUGC_GetSupportedGameVersionData, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetSupportedGameVersionData")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetSupportedGameVersionData")
	}
	addr_iSteamUGC_GetQueryUGCContentDescriptors, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetQueryUGCContentDescriptors")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetQueryUGCContentDescriptors")
	}
	addr_iSteamUGC_ReleaseQueryUGCRequest, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_ReleaseQueryUGCRequest")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_ReleaseQueryUGCRequest")
	}
	addr_iSteamUGC_AddRequiredTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddRequiredTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddRequiredTag")
	}
	addr_iSteamUGC_AddRequiredTagGroup, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddRequiredTagGroup")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddRequiredTagGroup")
	}
	addr_iSteamUGC_AddExcludedTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddExcludedTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddExcludedTag")
	}
	addr_iSteamUGC_SetReturnOnlyIDs, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnOnlyIDs")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnOnlyIDs")
	}
	addr_iSteamUGC_SetReturnKeyValueTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnKeyValueTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnKeyValueTags")
	}
	addr_iSteamUGC_SetReturnLongDescription, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnLongDescription")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnLongDescription")
	}
	addr_iSteamUGC_SetReturnMetadata, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnMetadata")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnMetadata")
	}
	addr_iSteamUGC_SetReturnChildren, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnChildren")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnChildren")
	}
	addr_iSteamUGC_SetReturnAdditionalPreviews, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnAdditionalPreviews")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnAdditionalPreviews")
	}
	addr_iSteamUGC_SetReturnTotalOnly, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnTotalOnly")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnTotalOnly")
	}
	addr_iSteamUGC_SetReturnPlaytimeStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetReturnPlaytimeStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetReturnPlaytimeStats")
	}
	addr_iSteamUGC_SetLanguage, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetLanguage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetLanguage")
	}
	addr_iSteamUGC_SetAllowCachedResponse, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetAllowCachedResponse")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetAllowCachedResponse")
	}
	addr_iSteamUGC_SetAdminQuery, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetAdminQuery")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetAdminQuery")
	}
	addr_iSteamUGC_SetCloudFileNameFilter, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetCloudFileNameFilter")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetCloudFileNameFilter")
	}
	addr_iSteamUGC_SetMatchAnyTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetMatchAnyTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetMatchAnyTag")
	}
	addr_iSteamUGC_SetSearchText, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetSearchText")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetSearchText")
	}
	addr_iSteamUGC_SetRankedByTrendDays, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetRankedByTrendDays")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetRankedByTrendDays")
	}
	addr_iSteamUGC_SetTimeCreatedDateRange, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetTimeCreatedDateRange")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetTimeCreatedDateRange")
	}
	addr_iSteamUGC_SetTimeUpdatedDateRange, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetTimeUpdatedDateRange")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetTimeUpdatedDateRange")
	}
	addr_iSteamUGC_AddRequiredKeyValueTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddRequiredKeyValueTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddRequiredKeyValueTag")
	}
	addr_iSteamUGC_CreateItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_CreateItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_CreateItem")
	}
	addr_iSteamUGC_StartItemUpdate, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_StartItemUpdate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_StartItemUpdate")
	}
	addr_iSteamUGC_SetItemTitle, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemTitle")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemTitle")
	}
	addr_iSteamUGC_SetItemDescription, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemDescription")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemDescription")
	}
	addr_iSteamUGC_SetItemUpdateLanguage, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemUpdateLanguage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemUpdateLanguage")
	}
	addr_iSteamUGC_SetItemMetadata, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemMetadata")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemMetadata")
	}
	addr_iSteamUGC_SetItemVisibility, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemVisibility")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemVisibility")
	}
	addr_iSteamUGC_SetItemTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemTags")
	}
	addr_iSteamUGC_SetItemContent, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemContent")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemContent")
	}
	addr_iSteamUGC_SetItemPreview, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetItemPreview")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetItemPreview")
	}
	addr_iSteamUGC_SetAllowLegacyUpload, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetAllowLegacyUpload")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetAllowLegacyUpload")
	}
	addr_iSteamUGC_RemoveAllItemKeyValueTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveAllItemKeyValueTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveAllItemKeyValueTags")
	}
	addr_iSteamUGC_RemoveItemKeyValueTags, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveItemKeyValueTags")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveItemKeyValueTags")
	}
	addr_iSteamUGC_AddItemKeyValueTag, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddItemKeyValueTag")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddItemKeyValueTag")
	}
	addr_iSteamUGC_AddItemPreviewFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddItemPreviewFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddItemPreviewFile")
	}
	addr_iSteamUGC_AddItemPreviewVideo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddItemPreviewVideo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddItemPreviewVideo")
	}
	addr_iSteamUGC_UpdateItemPreviewFile, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_UpdateItemPreviewFile")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_UpdateItemPreviewFile")
	}
	addr_iSteamUGC_UpdateItemPreviewVideo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_UpdateItemPreviewVideo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_UpdateItemPreviewVideo")
	}
	addr_iSteamUGC_RemoveItemPreview, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveItemPreview")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveItemPreview")
	}
	addr_iSteamUGC_AddContentDescriptor, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddContentDescriptor")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddContentDescriptor")
	}
	addr_iSteamUGC_RemoveContentDescriptor, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveContentDescriptor")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveContentDescriptor")
	}
	addr_iSteamUGC_SetRequiredGameVersions, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetRequiredGameVersions")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetRequiredGameVersions")
	}
	addr_iSteamUGC_SubmitItemUpdate, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SubmitItemUpdate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SubmitItemUpdate")
	}
	addr_iSteamUGC_GetItemUpdateProgress, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetItemUpdateProgress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetItemUpdateProgress")
	}
	addr_iSteamUGC_SetUserItemVote, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SetUserItemVote")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SetUserItemVote")
	}
	addr_iSteamUGC_GetUserItemVote, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetUserItemVote")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetUserItemVote")
	}
	addr_iSteamUGC_AddItemToFavorites, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddItemToFavorites")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddItemToFavorites")
	}
	addr_iSteamUGC_RemoveItemFromFavorites, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveItemFromFavorites")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveItemFromFavorites")
	}
	addr_iSteamUGC_SubscribeItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SubscribeItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SubscribeItem")
	}
	addr_iSteamUGC_UnsubscribeItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_UnsubscribeItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_UnsubscribeItem")
	}
	addr_iSteamUGC_GetNumSubscribedItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetNumSubscribedItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetNumSubscribedItems")
	}
	addr_iSteamUGC_GetSubscribedItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetSubscribedItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetSubscribedItems")
	}
	addr_iSteamUGC_GetItemState, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetItemState")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetItemState")
	}
	addr_iSteamUGC_GetItemInstallInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetItemInstallInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetItemInstallInfo")
	}
	addr_iSteamUGC_GetItemDownloadInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetItemDownloadInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetItemDownloadInfo")
	}
	addr_iSteamUGC_DownloadItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_DownloadItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_DownloadItem")
	}
	addr_iSteamUGC_BInitWorkshopForGameServer, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_BInitWorkshopForGameServer")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_BInitWorkshopForGameServer")
	}
	addr_iSteamUGC_SuspendDownloads, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_SuspendDownloads")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_SuspendDownloads")
	}
	addr_iSteamUGC_StartPlaytimeTracking, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_StartPlaytimeTracking")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_StartPlaytimeTracking")
	}
	addr_iSteamUGC_StopPlaytimeTracking, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_StopPlaytimeTracking")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_StopPlaytimeTracking")
	}
	addr_iSteamUGC_StopPlaytimeTrackingForAllItems, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_StopPlaytimeTrackingForAllItems")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_StopPlaytimeTrackingForAllItems")
	}
	addr_iSteamUGC_AddDependency, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddDependency")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddDependency")
	}
	addr_iSteamUGC_RemoveDependency, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveDependency")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveDependency")
	}
	addr_iSteamUGC_AddAppDependency, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_AddAppDependency")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_AddAppDependency")
	}
	addr_iSteamUGC_RemoveAppDependency, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_RemoveAppDependency")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_RemoveAppDependency")
	}
	addr_iSteamUGC_GetAppDependencies, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetAppDependencies")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetAppDependencies")
	}
	addr_iSteamUGC_DeleteItem, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_DeleteItem")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_DeleteItem")
	}
	addr_iSteamUGC_ShowWorkshopEULA, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_ShowWorkshopEULA")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_ShowWorkshopEULA")
	}
	addr_iSteamUGC_GetWorkshopEULAStatus, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetWorkshopEULAStatus")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetWorkshopEULAStatus")
	}
	addr_iSteamUGC_GetUserContentDescriptorPreferences, err = openSymbol(steamAPILib, "SteamAPI_ISteamUGC_GetUserContentDescriptorPreferences")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUGC_GetUserContentDescriptorPreferences")
	}

	steamUGC_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamUGC_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamUGC_CreateQueryUserUGCRequest = func(steamUGC uintptr, unAccountID AccountID, eListType EUserUGCList, eMatchingUGCType EUGCMatchingUGCType, eSortOrder EUserUGCListSortOrder, nCreatorAppID AppId_t, nConsumerAppID AppId_t, unPage uint32) UGCQueryHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_CreateQueryUserUGCRequest, uintptr(steamUGC), uintptr(unAccountID), uintptr(eListType), uintptr(eMatchingUGCType), uintptr(eSortOrder), uintptr(nCreatorAppID), uintptr(nConsumerAppID), uintptr(unPage))
		__r0 := UGCQueryHandle(_r0)
		return __r0
	}
	iSteamUGC_CreateQueryAllUGCRequestPage = func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId_t, nConsumerAppID AppId_t, unPage uint32) UGCQueryHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_CreateQueryAllUGCRequestPage, uintptr(steamUGC), uintptr(eQueryType), uintptr(eMatchingeMatchingUGCTypeFileType), uintptr(nCreatorAppID), uintptr(nConsumerAppID), uintptr(unPage))
		__r0 := UGCQueryHandle(_r0)
		return __r0
	}
	iSteamUGC_CreateQueryAllUGCRequestCursor = func(steamUGC uintptr, eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId_t, nConsumerAppID AppId_t, pchCursor string) UGCQueryHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_CreateQueryAllUGCRequestCursor, uintptr(steamUGC), uintptr(eQueryType), uintptr(eMatchingeMatchingUGCTypeFileType), uintptr(nCreatorAppID), uintptr(nConsumerAppID), uintptr(unsafe.Pointer(bytePtrFromString(pchCursor))))
		__r0 := UGCQueryHandle(_r0)
		runtime.KeepAlive(pchCursor)
		return __r0
	}
	iSteamUGC_CreateQueryUGCDetailsRequest = func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) UGCQueryHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_CreateQueryUGCDetailsRequest, uintptr(steamUGC), uintptr(unsafe.Pointer(unsafe.SliceData(pvecPublishedFileID))), uintptr(unNumPublishedFileIDs))
		__r0 := UGCQueryHandle(_r0)
		runtime.KeepAlive(pvecPublishedFileID)
		return __r0
	}
	iSteamUGC_SendQueryUGCRequest = func(steamUGC uintptr, handle UGCQueryHandle) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SendQueryUGCRequest, uintptr(steamUGC), uintptr(handle))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_GetQueryUGCResult = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pDetails *SteamUGCDetails) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCResult, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(pDetails)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pDetails)
		return __r0
	}
	iSteamUGC_GetQueryUGCNumTags = func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCNumTags, uintptr(steamUGC), uintptr(handle), uintptr(index))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetQueryUGCTag = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue []byte, cchValueSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCTag, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(indexTag), uintptr(unsafe.Pointer(unsafe.SliceData(pchValue))), uintptr(cchValueSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamUGC_GetQueryUGCTagDisplayName = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, indexTag uint32, pchValue []byte, cchValueSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCTagDisplayName, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(indexTag), uintptr(unsafe.Pointer(unsafe.SliceData(pchValue))), uintptr(cchValueSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamUGC_GetQueryUGCPreviewURL = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchURL []byte, cchURLSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCPreviewURL, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(unsafe.SliceData(pchURL))), uintptr(cchURLSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchURL)
		return __r0
	}
	iSteamUGC_GetQueryUGCMetadata = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchMetadata []byte, cchMetadatasize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCMetadata, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(unsafe.SliceData(pchMetadata))), uintptr(cchMetadatasize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchMetadata)
		return __r0
	}
	iSteamUGC_GetQueryUGCChildren = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecPublishedFileID []PublishedFileId, cMaxEntries uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCChildren, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(unsafe.SliceData(pvecPublishedFileID))), uintptr(cMaxEntries))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pvecPublishedFileID)
		return __r0
	}
	iSteamUGC_GetQueryUGCStatistic = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, eStatType EItemStatistic, pStatValue *uint64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCStatistic, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(eStatType), uintptr(unsafe.Pointer(pStatValue)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pStatValue)
		return __r0
	}
	iSteamUGC_GetQueryUGCNumAdditionalPreviews = func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCNumAdditionalPreviews, uintptr(steamUGC), uintptr(handle), uintptr(index))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetQueryUGCAdditionalPreview = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, previewIndex uint32, pchURLOrVideoID []byte, cchURLSize uint32, pchOriginalFileName []byte, cchOriginalFileNameSize uint32, pPreviewType *EItemPreviewType) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCAdditionalPreview, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(previewIndex), uintptr(unsafe.Pointer(unsafe.SliceData(pchURLOrVideoID))), uintptr(cchURLSize), uintptr(unsafe.Pointer(unsafe.SliceData(pchOriginalFileName))), uintptr(cchOriginalFileNameSize), uintptr(unsafe.Pointer(pPreviewType)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchURLOrVideoID)
		runtime.KeepAlive(pchOriginalFileName)
		runtime.KeepAlive(pPreviewType)
		return __r0
	}
	iSteamUGC_GetQueryUGCNumKeyValueTags = func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCNumKeyValueTags, uintptr(steamUGC), uintptr(handle), uintptr(index))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetQueryUGCKeyValueTag = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, pchKey []byte, cchKeySize uint32, pchValue []byte, cchValueSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCKeyValueTag, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(keyValueTagIndex), uintptr(unsafe.Pointer(unsafe.SliceData(pchKey))), uintptr(cchKeySize), uintptr(unsafe.Pointer(unsafe.SliceData(pchValue))), uintptr(cchValueSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamUGC_GetQueryFirstUGCKeyValueTag = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pchKey string, pchValue []byte, cchValueSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryFirstUGCKeyValueTag, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))), uintptr(unsafe.Pointer(unsafe.SliceData(pchValue))), uintptr(cchValueSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamUGC_GetNumSupportedGameVersions = func(steamUGC uintptr, handle UGCQueryHandle, index uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetNumSupportedGameVersions, uintptr(steamUGC), uintptr(handle), uintptr(index))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetSupportedGameVersionData = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, versionIndex uint32, pchGameBranchMin []byte, pchGameBranchMax []byte, cchGameBranchSize uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetSupportedGameVersionData, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(versionIndex), uintptr(unsafe.Pointer(unsafe.SliceData(pchGameBranchMin))), uintptr(unsafe.Pointer(unsafe.SliceData(pchGameBranchMax))), uintptr(cchGameBranchSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchGameBranchMin)
		runtime.KeepAlive(pchGameBranchMax)
		return __r0
	}
	iSteamUGC_GetQueryUGCContentDescriptors = func(steamUGC uintptr, handle UGCQueryHandle, index uint32, pvecDescriptors []EUGCContentDescriptorID, cMaxEntries uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetQueryUGCContentDescriptors, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(unsafe.SliceData(pvecDescriptors))), uintptr(cMaxEntries))
		__r0 := uint32(_r0)
		runtime.KeepAlive(pvecDescriptors)
		return __r0
	}
	iSteamUGC_ReleaseQueryUGCRequest = func(steamUGC uintptr, handle UGCQueryHandle) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_ReleaseQueryUGCRequest, uintptr(steamUGC), uintptr(handle))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_AddRequiredTag = func(steamUGC uintptr, handle UGCQueryHandle, pTagName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddRequiredTag, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pTagName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pTagName)
		return __r0
	}
	iSteamUGC_AddRequiredTagGroup = func(steamUGC uintptr, handle UGCQueryHandle, pTagGroups []SteamParamStringArray) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddRequiredTagGroup, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(unsafe.SliceData(pTagGroups))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pTagGroups)
		return __r0
	}
	iSteamUGC_AddExcludedTag = func(steamUGC uintptr, handle UGCQueryHandle, pTagName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddExcludedTag, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pTagName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pTagName)
		return __r0
	}
	iSteamUGC_SetReturnOnlyIDs = func(steamUGC uintptr, handle UGCQueryHandle, bReturnOnlyIDs bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnOnlyIDs, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnOnlyIDs))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnKeyValueTags = func(steamUGC uintptr, handle UGCQueryHandle, bReturnKeyValueTags bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnKeyValueTags, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnKeyValueTags))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnLongDescription = func(steamUGC uintptr, handle UGCQueryHandle, bReturnLongDescription bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnLongDescription, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnLongDescription))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnMetadata = func(steamUGC uintptr, handle UGCQueryHandle, bReturnMetadata bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnMetadata, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnMetadata))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnChildren = func(steamUGC uintptr, handle UGCQueryHandle, bReturnChildren bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnChildren, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnChildren))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnAdditionalPreviews = func(steamUGC uintptr, handle UGCQueryHandle, bReturnAdditionalPreviews bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnAdditionalPreviews, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnAdditionalPreviews))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnTotalOnly = func(steamUGC uintptr, handle UGCQueryHandle, bReturnTotalOnly bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnTotalOnly, uintptr(steamUGC), uintptr(handle), boolToUintptr(bReturnTotalOnly))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetReturnPlaytimeStats = func(steamUGC uintptr, handle UGCQueryHandle, unDays uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetReturnPlaytimeStats, uintptr(steamUGC), uintptr(handle), uintptr(unDays))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetLanguage = func(steamUGC uintptr, handle UGCQueryHandle, pchLanguage string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetLanguage, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchLanguage))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchLanguage)
		return __r0
	}
	iSteamUGC_SetAllowCachedResponse = func(steamUGC uintptr, handle UGCQueryHandle, unMaxAgeSeconds uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetAllowCachedResponse, uintptr(steamUGC), uintptr(handle), uintptr(unMaxAgeSeconds))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetAdminQuery = func(steamUGC uintptr, handle UGCUpdateHandle, bAdminQuery bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetAdminQuery, uintptr(steamUGC), uintptr(handle), boolToUintptr(bAdminQuery))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetCloudFileNameFilter = func(steamUGC uintptr, handle UGCQueryHandle, pMatchCloudFileName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetCloudFileNameFilter, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pMatchCloudFileName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pMatchCloudFileName)
		return __r0
	}
	iSteamUGC_SetMatchAnyTag = func(steamUGC uintptr, handle UGCQueryHandle, bMatchAnyTag bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetMatchAnyTag, uintptr(steamUGC), uintptr(handle), boolToUintptr(bMatchAnyTag))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetSearchText = func(steamUGC uintptr, handle UGCQueryHandle, pSearchText string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetSearchText, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pSearchText))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pSearchText)
		return __r0
	}
	iSteamUGC_SetRankedByTrendDays = func(steamUGC uintptr, handle UGCQueryHandle, unDays uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetRankedByTrendDays, uintptr(steamUGC), uintptr(handle), uintptr(unDays))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetTimeCreatedDateRange = func(steamUGC uintptr, handle UGCQueryHandle, rtStart RTime32, rtEnd RTime32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetTimeCreatedDateRange, uintptr(steamUGC), uintptr(handle), uintptr(rtStart), uintptr(rtEnd))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetTimeUpdatedDateRange = func(steamUGC uintptr, handle UGCQueryHandle, rtStart RTime32, rtEnd RTime32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetTimeUpdatedDateRange, uintptr(steamUGC), uintptr(handle), uintptr(rtStart), uintptr(rtEnd))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_AddRequiredKeyValueTag = func(steamUGC uintptr, handle UGCQueryHandle, pKey string, pValue string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddRequiredKeyValueTag, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pKey))), uintptr(unsafe.Pointer(bytePtrFromString(pValue))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pKey)
		runtime.KeepAlive(pValue)
		return __r0
	}
	iSteamUGC_CreateItem = func(steamUGC uintptr, nConsumerAppId AppId_t, eFileType EWorkshopFileType) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_CreateItem, uintptr(steamUGC), uintptr(nConsumerAppId), uintptr(eFileType))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_StartItemUpdate = func(steamUGC uintptr, nConsumerAppId AppId_t, nPublishedFileID PublishedFileId) UGCUpdateHandle {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_StartItemUpdate, uintptr(steamUGC), uintptr(nConsumerAppId), uintptr(nPublishedFileID))
		__r0 := UGCUpdateHandle(_r0)
		return __r0
	}
	iSteamUGC_SetItemTitle = func(steamUGC uintptr, handle UGCUpdateHandle, pchTitle string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemTitle, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchTitle))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchTitle)
		return __r0
	}
	iSteamUGC_SetItemDescription = func(steamUGC uintptr, handle UGCUpdateHandle, pchDescription string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemDescription, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchDescription))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchDescription)
		return __r0
	}
	iSteamUGC_SetItemUpdateLanguage = func(steamUGC uintptr, handle UGCUpdateHandle, pchLanguage string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemUpdateLanguage, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchLanguage))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchLanguage)
		return __r0
	}
	iSteamUGC_SetItemMetadata = func(steamUGC uintptr, handle UGCUpdateHandle, pchMetaData string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemMetadata, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchMetaData))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchMetaData)
		return __r0
	}
	iSteamUGC_SetItemVisibility = func(steamUGC uintptr, handle UGCUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemVisibility, uintptr(steamUGC), uintptr(handle), uintptr(eVisibility))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetItemTags = func(steamUGC uintptr, updateHandle UGCUpdateHandle, pTags []SteamParamStringArray, bAllowAdminTags bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemTags, uintptr(steamUGC), uintptr(updateHandle), uintptr(unsafe.Pointer(unsafe.SliceData(pTags))), boolToUintptr(bAllowAdminTags))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pTags)
		return __r0
	}
	iSteamUGC_SetItemContent = func(steamUGC uintptr, handle UGCUpdateHandle, pszContentFolder string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemContent, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pszContentFolder))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszContentFolder)
		return __r0
	}
	iSteamUGC_SetItemPreview = func(steamUGC uintptr, handle UGCUpdateHandle, pszPreviewFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetItemPreview, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pszPreviewFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszPreviewFile)
		return __r0
	}
	iSteamUGC_SetAllowLegacyUpload = func(steamUGC uintptr, handle UGCUpdateHandle, bAllowLegacyUpload bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetAllowLegacyUpload, uintptr(steamUGC), uintptr(handle), boolToUintptr(bAllowLegacyUpload))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_RemoveAllItemKeyValueTags = func(steamUGC uintptr, handle UGCUpdateHandle) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveAllItemKeyValueTags, uintptr(steamUGC), uintptr(handle))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_RemoveItemKeyValueTags = func(steamUGC uintptr, handle UGCUpdateHandle, pchKey string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveItemKeyValueTags, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		return __r0
	}
	iSteamUGC_AddItemKeyValueTag = func(steamUGC uintptr, handle UGCUpdateHandle, pchKey string, pchValue string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddItemKeyValueTag, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))), uintptr(unsafe.Pointer(bytePtrFromString(pchValue))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchKey)
		runtime.KeepAlive(pchValue)
		return __r0
	}
	iSteamUGC_AddItemPreviewFile = func(steamUGC uintptr, handle UGCUpdateHandle, pszPreviewFile string, Type EItemPreviewType) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddItemPreviewFile, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pszPreviewFile))), uintptr(Type))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszPreviewFile)
		return __r0
	}
	iSteamUGC_AddItemPreviewVideo = func(steamUGC uintptr, handle UGCUpdateHandle, pszVideoID string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddItemPreviewVideo, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pszVideoID))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszVideoID)
		return __r0
	}
	iSteamUGC_UpdateItemPreviewFile = func(steamUGC uintptr, handle UGCUpdateHandle, index uint32, pszPreviewFile string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_UpdateItemPreviewFile, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(bytePtrFromString(pszPreviewFile))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszPreviewFile)
		return __r0
	}
	iSteamUGC_UpdateItemPreviewVideo = func(steamUGC uintptr, handle UGCUpdateHandle, index uint32, pszVideoID string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_UpdateItemPreviewVideo, uintptr(steamUGC), uintptr(handle), uintptr(index), uintptr(unsafe.Pointer(bytePtrFromString(pszVideoID))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszVideoID)
		return __r0
	}
	iSteamUGC_RemoveItemPreview = func(steamUGC uintptr, handle UGCUpdateHandle, index uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveItemPreview, uintptr(steamUGC), uintptr(handle), uintptr(index))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_AddContentDescriptor = func(steamUGC uintptr, handle UGCUpdateHandle, descid EUGCContentDescriptorID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddContentDescriptor, uintptr(steamUGC), uintptr(handle), uintptr(descid))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_RemoveContentDescriptor = func(steamUGC uintptr, handle UGCUpdateHandle, descid EUGCContentDescriptorID) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveContentDescriptor, uintptr(steamUGC), uintptr(handle), uintptr(descid))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_SetRequiredGameVersions = func(steamUGC uintptr, handle UGCUpdateHandle, pszGameBranchMin string, pszGameBranchMax string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetRequiredGameVersions, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pszGameBranchMin))), uintptr(unsafe.Pointer(bytePtrFromString(pszGameBranchMax))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszGameBranchMin)
		runtime.KeepAlive(pszGameBranchMax)
		return __r0
	}
	iSteamUGC_SubmitItemUpdate = func(steamUGC uintptr, handle UGCUpdateHandle, pchChangeNote string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SubmitItemUpdate, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(bytePtrFromString(pchChangeNote))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchChangeNote)
		return __r0
	}
	iSteamUGC_GetItemUpdateProgress = func(steamUGC uintptr, handle UGCUpdateHandle, punBytesProcessed *uint64, punBytesTotal *uint64) EItemUpdateStatus {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetItemUpdateProgress, uintptr(steamUGC), uintptr(handle), uintptr(unsafe.Pointer(punBytesProcessed)), uintptr(unsafe.Pointer(punBytesTotal)))
		__r0 := EItemUpdateStatus(_r0)
		runtime.KeepAlive(punBytesProcessed)
		runtime.KeepAlive(punBytesTotal)
		return __r0
	}
	iSteamUGC_SetUserItemVote = func(steamUGC uintptr, nPublishedFileID PublishedFileId, bVoteUp bool) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SetUserItemVote, uintptr(steamUGC), uintptr(nPublishedFileID), boolToUintptr(bVoteUp))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_GetUserItemVote = func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetUserItemVote, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_AddItemToFavorites = func(steamUGC uintptr, nAppId AppId_t, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddItemToFavorites, uintptr(steamUGC), uintptr(nAppId), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_RemoveItemFromFavorites = func(steamUGC uintptr, nAppId AppId_t, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveItemFromFavorites, uintptr(steamUGC), uintptr(nAppId), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_SubscribeItem = func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_SubscribeItem, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_UnsubscribeItem = func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_UnsubscribeItem, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_GetNumSubscribedItems = func(steamUGC uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetNumSubscribedItems, uintptr(steamUGC))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetSubscribedItems = func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, cMaxEntries uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetSubscribedItems, uintptr(steamUGC), uintptr(unsafe.Pointer(unsafe.SliceData(pvecPublishedFileID))), uintptr(cMaxEntries))
		__r0 := uint32(_r0)
		runtime.KeepAlive(pvecPublishedFileID)
		return __r0
	}
	iSteamUGC_GetItemState = func(steamUGC uintptr, nPublishedFileID PublishedFileId) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetItemState, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUGC_GetItemInstallInfo = func(steamUGC uintptr, nPublishedFileID PublishedFileId, punSizeOnDisk *uint64, pchFolder []byte, cchFolderSize uint32, punTimeStamp *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetItemInstallInfo, uintptr(steamUGC), uintptr(nPublishedFileID), uintptr(unsafe.Pointer(punSizeOnDisk)), uintptr(unsafe.Pointer(unsafe.SliceData(pchFolder))), uintptr(cchFolderSize), uintptr(unsafe.Pointer(punTimeStamp)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(punSizeOnDisk)
		runtime.KeepAlive(pchFolder)
		runtime.KeepAlive(punTimeStamp)
		return __r0
	}
	iSteamUGC_GetItemDownloadInfo = func(steamUGC uintptr, nPublishedFileID PublishedFileId, punBytesDownloaded *uint64, punBytesTotal *uint64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetItemDownloadInfo, uintptr(steamUGC), uintptr(nPublishedFileID), uintptr(unsafe.Pointer(punBytesDownloaded)), uintptr(unsafe.Pointer(punBytesTotal)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(punBytesDownloaded)
		runtime.KeepAlive(punBytesTotal)
		return __r0
	}
	iSteamUGC_DownloadItem = func(steamUGC uintptr, nPublishedFileID PublishedFileId, bHighPriority bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_DownloadItem, uintptr(steamUGC), uintptr(nPublishedFileID), boolToUintptr(bHighPriority))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_BInitWorkshopForGameServer = func(steamUGC uintptr, unWorkshopDepotID DepotId, pszFolder string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_BInitWorkshopForGameServer, uintptr(steamUGC), uintptr(unWorkshopDepotID), uintptr(unsafe.Pointer(bytePtrFromString(pszFolder))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pszFolder)
		return __r0
	}
	iSteamUGC_SuspendDownloads = func(steamUGC uintptr, bSuspend bool) {
		purego.SyscallN(addr_iSteamUGC_SuspendDownloads, uintptr(steamUGC), boolToUintptr(bSuspend))
	}
	iSteamUGC_StartPlaytimeTracking = func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_StartPlaytimeTracking, uintptr(steamUGC), uintptr(unsafe.Pointer(unsafe.SliceData(pvecPublishedFileID))), uintptr(unNumPublishedFileIDs))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pvecPublishedFileID)
		return __r0
	}
	iSteamUGC_StopPlaytimeTracking = func(steamUGC uintptr, pvecPublishedFileID []PublishedFileId, unNumPublishedFileIDs uint32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_StopPlaytimeTracking, uintptr(steamUGC), uintptr(unsafe.Pointer(unsafe.SliceData(pvecPublishedFileID))), uintptr(unNumPublishedFileIDs))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pvecPublishedFileID)
		return __r0
	}
	iSteamUGC_StopPlaytimeTrackingForAllItems = func(steamUGC uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_StopPlaytimeTrackingForAllItems, uintptr(steamUGC))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_AddDependency = func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddDependency, uintptr(steamUGC), uintptr(nParentPublishedFileID), uintptr(nChildPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_RemoveDependency = func(steamUGC uintptr, nParentPublishedFileID PublishedFileId, nChildPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveDependency, uintptr(steamUGC), uintptr(nParentPublishedFileID), uintptr(nChildPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_AddAppDependency = func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId_t) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_AddAppDependency, uintptr(steamUGC), uintptr(nPublishedFileID), uintptr(nAppID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_RemoveAppDependency = func(steamUGC uintptr, nPublishedFileID PublishedFileId, nAppID AppId_t) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_RemoveAppDependency, uintptr(steamUGC), uintptr(nPublishedFileID), uintptr(nAppID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_GetAppDependencies = func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetAppDependencies, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_DeleteItem = func(steamUGC uintptr, nPublishedFileID PublishedFileId) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_DeleteItem, uintptr(steamUGC), uintptr(nPublishedFileID))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_ShowWorkshopEULA = func(steamUGC uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_ShowWorkshopEULA, uintptr(steamUGC))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUGC_GetWorkshopEULAStatus = func(steamUGC uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetWorkshopEULAStatus, uintptr(steamUGC))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUGC_GetUserContentDescriptorPreferences = func(steamUGC uintptr, pvecDescriptors []EUGCContentDescriptorID, cMaxEntries uint32) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUGC_GetUserContentDescriptorPreferences, uintptr(steamUGC), uintptr(unsafe.Pointer(unsafe.SliceData(pvecDescriptors))), uintptr(cMaxEntries))
		__r0 := uint32(_r0)
		runtime.KeepAlive(pvecDescriptors)
		return __r0
	}
}

var (
	addr_steamUser_get                             uintptr
	addr_iSteamUser_GetHSteamUser                  uintptr
	addr_iSteamUser_BLoggedOn                      uintptr
	addr_iSteamUser_GetSteamID                     uintptr
	addr_iSteamUser_TrackAppUsageEvent             uintptr
	addr_iSteamUser_GetUserDataFolder              uintptr
	addr_iSteamUser_StartVoiceRecording            uintptr
	addr_iSteamUser_StopVoiceRecording             uintptr
	addr_iSteamUser_GetAvailableVoice              uintptr
	addr_iSteamUser_GetVoice                       uintptr
	addr_iSteamUser_DecompressVoice                uintptr
	addr_iSteamUser_GetVoiceOptimalSampleRate      uintptr
	addr_iSteamUser_GetAuthSessionTicket           uintptr
	addr_iSteamUser_GetAuthTicketForWebApi         uintptr
	addr_iSteamUser_BeginAuthSession               uintptr
	addr_iSteamUser_EndAuthSession                 uintptr
	addr_iSteamUser_CancelAuthTicket               uintptr
	addr_iSteamUser_UserHasLicenseForApp           uintptr
	addr_iSteamUser_BIsBehindNAT                   uintptr
	addr_iSteamUser_AdvertiseGame                  uintptr
	addr_iSteamUser_RequestEncryptedAppTicket      uintptr
	addr_iSteamUser_GetEncryptedAppTicket          uintptr
	addr_iSteamUser_GetGameBadgeLevel              uintptr
	addr_iSteamUser_GetPlayerSteamLevel            uintptr
	addr_iSteamUser_RequestStoreAuthURL            uintptr
	addr_iSteamUser_BIsPhoneVerified               uintptr
	addr_iSteamUser_BIsTwoFactorEnabled            uintptr
	addr_iSteamUser_BIsPhoneIdentifying            uintptr
	addr_iSteamUser_BIsPhoneRequiringVerification  uintptr
	addr_iSteamUser_GetMarketEligibility           uintptr
	addr_iSteamUser_GetDurationControl             uintptr
	addr_iSteamUser_BSetDurationControlOnlineState uintptr
)

func initUser() {
	var err error
	addr_steamUser_get, err = openSymbol(steamAPILib, flatAPI_SteamUser)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamUser)
	}
	addr_iSteamUser_GetHSteamUser, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetHSteamUser")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetHSteamUser")
	}
	addr_iSteamUser_BLoggedOn, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BLoggedOn")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BLoggedOn")
	}
	addr_iSteamUser_GetSteamID, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetSteamID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetSteamID")
	}
	addr_iSteamUser_TrackAppUsageEvent, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_TrackAppUsageEvent")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_TrackAppUsageEvent")
	}
	addr_iSteamUser_GetUserDataFolder, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetUserDataFolder")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetUserDataFolder")
	}
	addr_iSteamUser_StartVoiceRecording, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_StartVoiceRecording")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_StartVoiceRecording")
	}
	addr_iSteamUser_StopVoiceRecording, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_StopVoiceRecording")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_StopVoiceRecording")
	}
	addr_iSteamUser_GetAvailableVoice, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetAvailableVoice")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetAvailableVoice")
	}
	addr_iSteamUser_GetVoice, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetVoice")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetVoice")
	}
	addr_iSteamUser_DecompressVoice, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_DecompressVoice")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_DecompressVoice")
	}
	addr_iSteamUser_GetVoiceOptimalSampleRate, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetVoiceOptimalSampleRate")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetVoiceOptimalSampleRate")
	}
	addr_iSteamUser_GetAuthSessionTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetAuthSessionTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetAuthSessionTicket")
	}
	addr_iSteamUser_GetAuthTicketForWebApi, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetAuthTicketForWebApi")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetAuthTicketForWebApi")
	}
	addr_iSteamUser_BeginAuthSession, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BeginAuthSession")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BeginAuthSession")
	}
	addr_iSteamUser_EndAuthSession, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_EndAuthSession")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_EndAuthSession")
	}
	addr_iSteamUser_CancelAuthTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_CancelAuthTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_CancelAuthTicket")
	}
	addr_iSteamUser_UserHasLicenseForApp, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_UserHasLicenseForApp")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_UserHasLicenseForApp")
	}
	addr_iSteamUser_BIsBehindNAT, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BIsBehindNAT")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BIsBehindNAT")
	}
	addr_iSteamUser_AdvertiseGame, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_AdvertiseGame")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_AdvertiseGame")
	}
	addr_iSteamUser_RequestEncryptedAppTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_RequestEncryptedAppTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_RequestEncryptedAppTicket")
	}
	addr_iSteamUser_GetEncryptedAppTicket, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetEncryptedAppTicket")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetEncryptedAppTicket")
	}
	addr_iSteamUser_GetGameBadgeLevel, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetGameBadgeLevel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetGameBadgeLevel")
	}
	addr_iSteamUser_GetPlayerSteamLevel, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetPlayerSteamLevel")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetPlayerSteamLevel")
	}
	addr_iSteamUser_RequestStoreAuthURL, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_RequestStoreAuthURL")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_RequestStoreAuthURL")
	}
	addr_iSteamUser_BIsPhoneVerified, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BIsPhoneVerified")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BIsPhoneVerified")
	}
	addr_iSteamUser_BIsTwoFactorEnabled, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BIsTwoFactorEnabled")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BIsTwoFactorEnabled")
	}
	addr_iSteamUser_BIsPhoneIdentifying, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BIsPhoneIdentifying")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BIsPhoneIdentifying")
	}
	addr_iSteamUser_BIsPhoneRequiringVerification, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BIsPhoneRequiringVerification")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BIsPhoneRequiringVerification")
	}
	addr_iSteamUser_GetMarketEligibility, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetMarketEligibility")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetMarketEligibility")
	}
	addr_iSteamUser_GetDurationControl, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_GetDurationControl")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_GetDurationControl")
	}
	addr_iSteamUser_BSetDurationControlOnlineState, err = openSymbol(steamAPILib, "SteamAPI_ISteamUser_BSetDurationControlOnlineState")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUser_BSetDurationControlOnlineState")
	}

	steamUser_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamUser_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamUser_GetHSteamUser = func(steamUser uintptr) HSteamUser {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetHSteamUser, uintptr(steamUser))
		__r0 := HSteamUser(_r0)
		return __r0
	}
	iSteamUser_BLoggedOn = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BLoggedOn, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_GetSteamID = func(steamUser uintptr) CSteamID {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetSteamID, uintptr(steamUser))
		__r0 := CSteamID(_r0)
		return __r0
	}
	iSteamUser_TrackAppUsageEvent = func(steamUser uintptr, gameID Uint64SteamID, eAppUsageEvent int32, pchExtraInfo string) {
		purego.SyscallN(addr_iSteamUser_TrackAppUsageEvent, uintptr(steamUser), uintptr(gameID), uintptr(eAppUsageEvent), uintptr(unsafe.Pointer(bytePtrFromString(pchExtraInfo))))
		runtime.KeepAlive(pchExtraInfo)
	}
	iSteamUser_GetUserDataFolder = func(steamUser uintptr, pchBuffer []byte, cubBuffer int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetUserDataFolder, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pchBuffer))), uintptr(cubBuffer))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchBuffer)
		return __r0
	}
	iSteamUser_StartVoiceRecording = func(steamUser uintptr) {
		purego.SyscallN(addr_iSteamUser_StartVoiceRecording, uintptr(steamUser))
	}
	iSteamUser_StopVoiceRecording = func(steamUser uintptr) {
		purego.SyscallN(addr_iSteamUser_StopVoiceRecording, uintptr(steamUser))
	}
	iSteamUser_GetAvailableVoice = func(steamUser uintptr, pcbCompressed *uint32, pcbUncompressedDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetAvailableVoice, uintptr(steamUser), uintptr(unsafe.Pointer(pcbCompressed)), uintptr(unsafe.Pointer(pcbUncompressedDeprecated)), uintptr(nUncompressedVoiceDesiredSampleRateDeprecated))
		__r0 := EVoiceResult(_r0)
		runtime.KeepAlive(pcbCompressed)
		runtime.KeepAlive(pcbUncompressedDeprecated)
		return __r0
	}
	iSteamUser_GetVoice = func(steamUser uintptr, bWantCompressed bool, pDestBuffer []byte, cbDestBufferSize uint32, nBytesWritten *uint32, bWantUncompressedDeprecated bool, pUncompressedDestBufferDeprecated []byte, cbUncompressedDestBufferSizeDeprecated uint32, nUncompressBytesWrittenDeprecated *uint32, nUncompressedVoiceDesiredSampleRateDeprecated uint32) EVoiceResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetVoice, uintptr(steamUser), boolToUintptr(bWantCompressed), uintptr(unsafe.Pointer(unsafe.SliceData(pDestBuffer))), uintptr(cbDestBufferSize), uintptr(unsafe.Pointer(nBytesWritten)), boolToUintptr(bWantUncompressedDeprecated), uintptr(unsafe.Pointer(unsafe.SliceData(pUncompressedDestBufferDeprecated))), uintptr(cbUncompressedDestBufferSizeDeprecated), uintptr(unsafe.Pointer(nUncompressBytesWrittenDeprecated)), uintptr(nUncompressedVoiceDesiredSampleRateDeprecated))
		__r0 := EVoiceResult(_r0)
		runtime.KeepAlive(pDestBuffer)
		runtime.KeepAlive(nBytesWritten)
		runtime.KeepAlive(pUncompressedDestBufferDeprecated)
		runtime.KeepAlive(nUncompressBytesWrittenDeprecated)
		return __r0
	}
	iSteamUser_DecompressVoice = func(steamUser uintptr, pCompressed []byte, cbCompressed uint32, pDestBuffer []byte, cbDestBufferSize uint32, nBytesWritten *uint32, nDesiredSampleRate uint32) EVoiceResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_DecompressVoice, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pCompressed))), uintptr(cbCompressed), uintptr(unsafe.Pointer(unsafe.SliceData(pDestBuffer))), uintptr(cbDestBufferSize), uintptr(unsafe.Pointer(nBytesWritten)), uintptr(nDesiredSampleRate))
		__r0 := EVoiceResult(_r0)
		runtime.KeepAlive(pCompressed)
		runtime.KeepAlive(pDestBuffer)
		runtime.KeepAlive(nBytesWritten)
		return __r0
	}
	iSteamUser_GetVoiceOptimalSampleRate = func(steamUser uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetVoiceOptimalSampleRate, uintptr(steamUser))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUser_GetAuthSessionTicket = func(steamUser uintptr, pTicket []byte, cbMaxTicket int32, pcbTicket *uint32, pSteamNetworkingIdentity *SteamNetworkingIdentity) HAuthTicket {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetAuthSessionTicket, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pTicket))), uintptr(cbMaxTicket), uintptr(unsafe.Pointer(pcbTicket)), uintptr(unsafe.Pointer(pSteamNetworkingIdentity)))
		__r0 := HAuthTicket(_r0)
		runtime.KeepAlive(pTicket)
		runtime.KeepAlive(pcbTicket)
		runtime.KeepAlive(pSteamNetworkingIdentity)
		return __r0
	}
	iSteamUser_GetAuthTicketForWebApi = func(steamUser uintptr, pchIdentity string) HAuthTicket {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetAuthTicketForWebApi, uintptr(steamUser), uintptr(unsafe.Pointer(bytePtrFromString(pchIdentity))))
		__r0 := HAuthTicket(_r0)
		runtime.KeepAlive(pchIdentity)
		return __r0
	}
	iSteamUser_BeginAuthSession = func(steamUser uintptr, pAuthTicket []byte, cbAuthTicket int32, steamID Uint64SteamID) EBeginAuthSessionResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BeginAuthSession, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pAuthTicket))), uintptr(cbAuthTicket), uintptr(steamID))
		__r0 := EBeginAuthSessionResult(_r0)
		runtime.KeepAlive(pAuthTicket)
		return __r0
	}
	iSteamUser_EndAuthSession = func(steamUser uintptr, steamID Uint64SteamID) {
		purego.SyscallN(addr_iSteamUser_EndAuthSession, uintptr(steamUser), uintptr(steamID))
	}
	iSteamUser_CancelAuthTicket = func(steamUser uintptr, hAuthTicket HAuthTicket) {
		purego.SyscallN(addr_iSteamUser_CancelAuthTicket, uintptr(steamUser), uintptr(hAuthTicket))
	}
	iSteamUser_UserHasLicenseForApp = func(steamUser uintptr, steamID Uint64SteamID, appID AppId_t) EUserHasLicenseForAppResult {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_UserHasLicenseForApp, uintptr(steamUser), uintptr(steamID), uintptr(appID))
		__r0 := EUserHasLicenseForAppResult(_r0)
		return __r0
	}
	iSteamUser_BIsBehindNAT = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BIsBehindNAT, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_AdvertiseGame = func(steamUser uintptr, steamIDGameServer Uint64SteamID, unIPServer uint32, usPortServer uint16) {
		purego.SyscallN(addr_iSteamUser_AdvertiseGame, uintptr(steamUser), uintptr(steamIDGameServer), uintptr(unIPServer), uintptr(usPortServer))
	}
	iSteamUser_RequestEncryptedAppTicket = func(steamUser uintptr, pDataToInclude []byte, cbDataToInclude int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_RequestEncryptedAppTicket, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pDataToInclude))), uintptr(cbDataToInclude))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pDataToInclude)
		return __r0
	}
	iSteamUser_GetEncryptedAppTicket = func(steamUser uintptr, pTicket []byte, cbMaxTicket int32, pcbTicket *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetEncryptedAppTicket, uintptr(steamUser), uintptr(unsafe.Pointer(unsafe.SliceData(pTicket))), uintptr(cbMaxTicket), uintptr(unsafe.Pointer(pcbTicket)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pTicket)
		runtime.KeepAlive(pcbTicket)
		return __r0
	}
	iSteamUser_GetGameBadgeLevel = func(steamUser uintptr, nSeries int32, bFoil bool) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetGameBadgeLevel, uintptr(steamUser), uintptr(nSeries), boolToUintptr(bFoil))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamUser_GetPlayerSteamLevel = func(steamUser uintptr) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetPlayerSteamLevel, uintptr(steamUser))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamUser_RequestStoreAuthURL = func(steamUser uintptr, pchRedirectURL string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_RequestStoreAuthURL, uintptr(steamUser), uintptr(unsafe.Pointer(bytePtrFromString(pchRedirectURL))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchRedirectURL)
		return __r0
	}
	iSteamUser_BIsPhoneVerified = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BIsPhoneVerified, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_BIsTwoFactorEnabled = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BIsTwoFactorEnabled, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_BIsPhoneIdentifying = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BIsPhoneIdentifying, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_BIsPhoneRequiringVerification = func(steamUser uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BIsPhoneRequiringVerification, uintptr(steamUser))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUser_GetMarketEligibility = func(steamUser uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetMarketEligibility, uintptr(steamUser))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUser_GetDurationControl = func(steamUser uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_GetDurationControl, uintptr(steamUser))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUser_BSetDurationControlOnlineState = func(steamUser uintptr, eNewState EDurationControlOnlineState) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUser_BSetDurationControlOnlineState, uintptr(steamUser), uintptr(eNewState))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}

var (
	addr_steamUserStats_get                                  uintptr
	addr_iSteamUserStats_GetStatInt32                        uintptr
	addr_iSteamUserStats_GetStatFloat                        uintptr
	addr_iSteamUserStats_SetStatInt32                        uintptr
	addr_iSteamUserStats_SetStatFloat                        uintptr
	addr_iSteamUserStats_UpdateAvgRateStat                   uintptr
	addr_iSteamUserStats_GetAchievement                      uintptr
	addr_iSteamUserStats_SetAchievement                      uintptr
	addr_iSteamUserStats_ClearAchievement                    uintptr
	addr_iSteamUserStats_GetAchievementAndUnlockTime         uintptr
	addr_iSteamUserStats_StoreStats                          uintptr
	addr_iSteamUserStats_GetAchievementIcon                  uintptr
	addr_iSteamUserStats_GetAchievementDisplayAttribute      uintptr
	addr_iSteamUserStats_IndicateAchievementProgress         uintptr
	addr_iSteamUserStats_GetNumAchievements                  uintptr
	addr_iSteamUserStats_GetAchievementName                  uintptr
	addr_iSteamUserStats_RequestUserStats                    uintptr
	addr_iSteamUserStats_GetUserStatInt32                    uintptr
	addr_iSteamUserStats_GetUserStatFloat                    uintptr
	addr_iSteamUserStats_GetUserAchievement                  uintptr
	addr_iSteamUserStats_GetUserAchievementAndUnlockTime     uintptr
	addr_iSteamUserStats_ResetAllStats                       uintptr
	addr_iSteamUserStats_FindOrCreateLeaderboard             uintptr
	addr_iSteamUserStats_FindLeaderboard                     uintptr
	addr_iSteamUserStats_GetLeaderboardName                  uintptr
	addr_iSteamUserStats_GetLeaderboardEntryCount            uintptr
	addr_iSteamUserStats_GetLeaderboardSortMethod            uintptr
	addr_iSteamUserStats_GetLeaderboardDisplayType           uintptr
	addr_iSteamUserStats_DownloadLeaderboardEntries          uintptr
	addr_iSteamUserStats_DownloadLeaderboardEntriesForUsers  uintptr
	addr_iSteamUserStats_GetDownloadedLeaderboardEntry       uintptr
	addr_iSteamUserStats_UploadLeaderboardScore              uintptr
	addr_iSteamUserStats_AttachLeaderboardUGC                uintptr
	addr_iSteamUserStats_GetNumberOfCurrentPlayers           uintptr
	addr_iSteamUserStats_RequestGlobalAchievementPercentages uintptr
	addr_iSteamUserStats_GetMostAchievedAchievementInfo      uintptr
	addr_iSteamUserStats_GetNextMostAchievedAchievementInfo  uintptr
	addr_iSteamUserStats_GetAchievementAchievedPercent       uintptr
	addr_iSteamUserStats_RequestGlobalStats                  uintptr
	addr_iSteamUserStats_GetGlobalStatInt64                  uintptr
	addr_iSteamUserStats_GetGlobalStatDouble                 uintptr
	addr_iSteamUserStats_GetGlobalStatHistoryInt64           uintptr
	addr_iSteamUserStats_GetGlobalStatHistoryDouble          uintptr
	addr_iSteamUserStats_GetAchievementProgressLimitsInt32   uintptr
	addr_iSteamUserStats_GetAchievementProgressLimitsFloat   uintptr
)

func initUserStats() {
	var err error
	addr_steamUserStats_get, err = openSymbol(steamAPILib, flatAPI_SteamUserStats)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamUserStats)
	}
	addr_iSteamUserStats_GetStatInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetStatInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetStatInt32")
	}
	addr_iSteamUserStats_GetStatFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetStatFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetStatFloat")
	}
	addr_iSteamUserStats_SetStatInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_SetStatInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_SetStatInt32")
	}
	addr_iSteamUserStats_SetStatFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_SetStatFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_SetStatFloat")
	}
	addr_iSteamUserStats_UpdateAvgRateStat, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_UpdateAvgRateStat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_UpdateAvgRateStat")
	}
	addr_iSteamUserStats_GetAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievement")
	}
	addr_iSteamUserStats_SetAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_SetAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_SetAchievement")
	}
	addr_iSteamUserStats_ClearAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_ClearAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_ClearAchievement")
	}
	addr_iSteamUserStats_GetAchievementAndUnlockTime, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementAndUnlockTime")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementAndUnlockTime")
	}
	addr_iSteamUserStats_StoreStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_StoreStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_StoreStats")
	}
	addr_iSteamUserStats_GetAchievementIcon, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementIcon")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementIcon")
	}
	addr_iSteamUserStats_GetAchievementDisplayAttribute, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementDisplayAttribute")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementDisplayAttribute")
	}
	addr_iSteamUserStats_IndicateAchievementProgress, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_IndicateAchievementProgress")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_IndicateAchievementProgress")
	}
	addr_iSteamUserStats_GetNumAchievements, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetNumAchievements")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetNumAchievements")
	}
	addr_iSteamUserStats_GetAchievementName, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementName")
	}
	addr_iSteamUserStats_RequestUserStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_RequestUserStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_RequestUserStats")
	}
	addr_iSteamUserStats_GetUserStatInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetUserStatInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetUserStatInt32")
	}
	addr_iSteamUserStats_GetUserStatFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetUserStatFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetUserStatFloat")
	}
	addr_iSteamUserStats_GetUserAchievement, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetUserAchievement")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetUserAchievement")
	}
	addr_iSteamUserStats_GetUserAchievementAndUnlockTime, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetUserAchievementAndUnlockTime")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetUserAchievementAndUnlockTime")
	}
	addr_iSteamUserStats_ResetAllStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_ResetAllStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_ResetAllStats")
	}
	addr_iSteamUserStats_FindOrCreateLeaderboard, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_FindOrCreateLeaderboard")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_FindOrCreateLeaderboard")
	}
	addr_iSteamUserStats_FindLeaderboard, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_FindLeaderboard")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_FindLeaderboard")
	}
	addr_iSteamUserStats_GetLeaderboardName, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetLeaderboardName")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetLeaderboardName")
	}
	addr_iSteamUserStats_GetLeaderboardEntryCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetLeaderboardEntryCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetLeaderboardEntryCount")
	}
	addr_iSteamUserStats_GetLeaderboardSortMethod, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetLeaderboardSortMethod")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetLeaderboardSortMethod")
	}
	addr_iSteamUserStats_GetLeaderboardDisplayType, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetLeaderboardDisplayType")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetLeaderboardDisplayType")
	}
	addr_iSteamUserStats_DownloadLeaderboardEntries, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_DownloadLeaderboardEntries")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_DownloadLeaderboardEntries")
	}
	addr_iSteamUserStats_DownloadLeaderboardEntriesForUsers, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers")
	}
	addr_iSteamUserStats_GetDownloadedLeaderboardEntry, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetDownloadedLeaderboardEntry")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetDownloadedLeaderboardEntry")
	}
	addr_iSteamUserStats_UploadLeaderboardScore, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_UploadLeaderboardScore")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_UploadLeaderboardScore")
	}
	addr_iSteamUserStats_AttachLeaderboardUGC, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_AttachLeaderboardUGC")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_AttachLeaderboardUGC")
	}
	addr_iSteamUserStats_GetNumberOfCurrentPlayers, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetNumberOfCurrentPlayers")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetNumberOfCurrentPlayers")
	}
	addr_iSteamUserStats_RequestGlobalAchievementPercentages, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_RequestGlobalAchievementPercentages")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_RequestGlobalAchievementPercentages")
	}
	addr_iSteamUserStats_GetMostAchievedAchievementInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetMostAchievedAchievementInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetMostAchievedAchievementInfo")
	}
	addr_iSteamUserStats_GetNextMostAchievedAchievementInfo, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo")
	}
	addr_iSteamUserStats_GetAchievementAchievedPercent, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementAchievedPercent")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementAchievedPercent")
	}
	addr_iSteamUserStats_RequestGlobalStats, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_RequestGlobalStats")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_RequestGlobalStats")
	}
	addr_iSteamUserStats_GetGlobalStatInt64, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetGlobalStatInt64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetGlobalStatInt64")
	}
	addr_iSteamUserStats_GetGlobalStatDouble, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetGlobalStatDouble")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetGlobalStatDouble")
	}
	addr_iSteamUserStats_GetGlobalStatHistoryInt64, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetGlobalStatHistoryInt64")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetGlobalStatHistoryInt64")
	}
	addr_iSteamUserStats_GetGlobalStatHistoryDouble, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetGlobalStatHistoryDouble")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetGlobalStatHistoryDouble")
	}
	addr_iSteamUserStats_GetAchievementProgressLimitsInt32, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsInt32")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementProgressLimitsInt32")
	}
	addr_iSteamUserStats_GetAchievementProgressLimitsFloat, err = openSymbol(steamAPILib, "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsFloat")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUserStats_GetAchievementProgressLimitsFloat")
	}

	steamUserStats_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamUserStats_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamUserStats_GetStatInt32 = func(steamUserStats uintptr, pchName string, pData *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetStatInt32, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pData)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamUserStats_SetStatInt32 = func(steamUserStats uintptr, pchName string, nData int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_SetStatInt32, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(nData))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		return __r0
	}
	iSteamUserStats_GetAchievement = func(steamUserStats uintptr, pchName string, pbAchieved *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievement, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pbAchieved)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pbAchieved)
		return __r0
	}
	iSteamUserStats_SetAchievement = func(steamUserStats uintptr, pchName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_SetAchievement, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		return __r0
	}
	iSteamUserStats_ClearAchievement = func(steamUserStats uintptr, pchName string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_ClearAchievement, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		return __r0
	}
	iSteamUserStats_GetAchievementAndUnlockTime = func(steamUserStats uintptr, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievementAndUnlockTime, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pbAchieved)), uintptr(unsafe.Pointer(punUnlockTime)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pbAchieved)
		runtime.KeepAlive(punUnlockTime)
		return __r0
	}
	iSteamUserStats_StoreStats = func(steamUserStats uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_StoreStats, uintptr(steamUserStats))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUserStats_GetAchievementIcon = func(steamUserStats uintptr, pchName string) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievementIcon, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))))
		__r0 := int32(_r0)
		runtime.KeepAlive(pchName)
		return __r0
	}
	iSteamUserStats_GetAchievementDisplayAttribute = func(steamUserStats uintptr, pchName string, pchKey string) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievementDisplayAttribute, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(bytePtrFromString(pchKey))))
		__r0 := goString(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pchKey)
		return __r0
	}
	iSteamUserStats_IndicateAchievementProgress = func(steamUserStats uintptr, pchName string, nCurProgress uint32, nMaxProgress uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_IndicateAchievementProgress, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(nCurProgress), uintptr(nMaxProgress))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		return __r0
	}
	iSteamUserStats_GetNumAchievements = func(steamUserStats uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetNumAchievements, uintptr(steamUserStats))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUserStats_GetAchievementName = func(steamUserStats uintptr, iAchievement uint32) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievementName, uintptr(steamUserStats), uintptr(iAchievement))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamUserStats_RequestUserStats = func(steamUserStats uintptr, steamIDUser Uint64SteamID) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_RequestUserStats, uintptr(steamUserStats), uintptr(steamIDUser))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_GetUserStatInt32 = func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pData *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetUserStatInt32, uintptr(steamUserStats), uintptr(steamIDUser), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pData)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamUserStats_GetUserAchievement = func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetUserAchievement, uintptr(steamUserStats), uintptr(steamIDUser), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pbAchieved)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pbAchieved)
		return __r0
	}
	iSteamUserStats_GetUserAchievementAndUnlockTime = func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetUserAchievementAndUnlockTime, uintptr(steamUserStats), uintptr(steamIDUser), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pbAchieved)), uintptr(unsafe.Pointer(punUnlockTime)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pbAchieved)
		runtime.KeepAlive(punUnlockTime)
		return __r0
	}
	iSteamUserStats_ResetAllStats = func(steamUserStats uintptr, bAchievementsToo bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_ResetAllStats, uintptr(steamUserStats), boolToUintptr(bAchievementsToo))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUserStats_FindOrCreateLeaderboard = func(steamUserStats uintptr, pchLeaderboardName string, eLeaderboardSortMethod ELeaderboardSortMethod, eLeaderboardDisplayType ELeaderboardDisplayType) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_FindOrCreateLeaderboard, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchLeaderboardName))), uintptr(eLeaderboardSortMethod), uintptr(eLeaderboardDisplayType))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchLeaderboardName)
		return __r0
	}
	iSteamUserStats_FindLeaderboard = func(steamUserStats uintptr, pchLeaderboardName string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_FindLeaderboard, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchLeaderboardName))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pchLeaderboardName)
		return __r0
	}
	iSteamUserStats_GetLeaderboardName = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetLeaderboardName, uintptr(steamUserStats), uintptr(hSteamLeaderboard))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamUserStats_GetLeaderboardEntryCount = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetLeaderboardEntryCount, uintptr(steamUserStats), uintptr(hSteamLeaderboard))
		__r0 := int32(_r0)
		return __r0
	}
	iSteamUserStats_GetLeaderboardSortMethod = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardSortMethod {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetLeaderboardSortMethod, uintptr(steamUserStats), uintptr(hSteamLeaderboard))
		__r0 := ELeaderboardSortMethod(_r0)
		return __r0
	}
	iSteamUserStats_GetLeaderboardDisplayType = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardDisplayType {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetLeaderboardDisplayType, uintptr(steamUserStats), uintptr(hSteamLeaderboard))
		__r0 := ELeaderboardDisplayType(_r0)
		return __r0
	}
	iSteamUserStats_DownloadLeaderboardEntries = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_DownloadLeaderboardEntries, uintptr(steamUserStats), uintptr(hSteamLeaderboard), uintptr(eLeaderboardDataRequest), uintptr(nRangeStart), uintptr(nRangeEnd))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_DownloadLeaderboardEntriesForUsers = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, prgUsers []CSteamID, cUsers int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_DownloadLeaderboardEntriesForUsers, uintptr(steamUserStats), uintptr(hSteamLeaderboard), uintptr(unsafe.Pointer(unsafe.SliceData(prgUsers))), uintptr(cUsers))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(prgUsers)
		return __r0
	}
	iSteamUserStats_GetDownloadedLeaderboardEntry = func(steamUserStats uintptr, hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, pLeaderboardEntry *LeaderboardEntry, pDetails []int32, cDetailsMax int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetDownloadedLeaderboardEntry, uintptr(steamUserStats), uintptr(hSteamLeaderboardEntries), uintptr(index), uintptr(unsafe.Pointer(pLeaderboardEntry)), uintptr(unsafe.Pointer(unsafe.SliceData(pDetails))), uintptr(cDetailsMax))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pLeaderboardEntry)
		runtime.KeepAlive(pDetails)
		return __r0
	}
	iSteamUserStats_UploadLeaderboardScore = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32, cScoreDetailsCount int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_UploadLeaderboardScore, uintptr(steamUserStats), uintptr(hSteamLeaderboard), uintptr(eLeaderboardUploadScoreMethod), uintptr(nScore), uintptr(unsafe.Pointer(unsafe.SliceData(pScoreDetails))), uintptr(cScoreDetailsCount))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(pScoreDetails)
		return __r0
	}
	iSteamUserStats_AttachLeaderboardUGC = func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, hUGC UGCHandle) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_AttachLeaderboardUGC, uintptr(steamUserStats), uintptr(hSteamLeaderboard), uintptr(hUGC))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_GetNumberOfCurrentPlayers = func(steamUserStats uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetNumberOfCurrentPlayers, uintptr(steamUserStats))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_RequestGlobalAchievementPercentages = func(steamUserStats uintptr) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_RequestGlobalAchievementPercentages, uintptr(steamUserStats))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_RequestGlobalStats = func(steamUserStats uintptr, nHistoryDays int32) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_RequestGlobalStats, uintptr(steamUserStats), uintptr(nHistoryDays))
		__r0 := SteamAPICall(_r0)
		return __r0
	}
	iSteamUserStats_GetGlobalStatInt64 = func(steamUserStats uintptr, pchStatName string, pData *int64) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetGlobalStatInt64, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchStatName))), uintptr(unsafe.Pointer(pData)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchStatName)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamUserStats_GetGlobalStatHistoryInt64 = func(steamUserStats uintptr, pchStatName string, pData []int64, cubData uint32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetGlobalStatHistoryInt64, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchStatName))), uintptr(unsafe.Pointer(unsafe.SliceData(pData))), uintptr(cubData))
		__r0 := int32(_r0)
		runtime.KeepAlive(pchStatName)
		runtime.KeepAlive(pData)
		return __r0
	}
	iSteamUserStats_GetAchievementProgressLimitsInt32 = func(steamUserStats uintptr, pchName string, pnMinProgress *int32, pnMaxProgress *int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUserStats_GetAchievementProgressLimitsInt32, uintptr(steamUserStats), uintptr(unsafe.Pointer(bytePtrFromString(pchName))), uintptr(unsafe.Pointer(pnMinProgress)), uintptr(unsafe.Pointer(pnMaxProgress)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchName)
		runtime.KeepAlive(pnMinProgress)
		runtime.KeepAlive(pnMaxProgress)
		return __r0
	}

	purego.RegisterLibFunc(&iSteamUserStats_GetStatFloat, steamAPILib, flatAPI_ISteamUserStats_GetStatFloat)
	purego.RegisterLibFunc(&iSteamUserStats_SetStatFloat, steamAPILib, flatAPI_ISteamUserStats_SetStatFloat)
	purego.RegisterLibFunc(&iSteamUserStats_UpdateAvgRateStat, steamAPILib, flatAPI_ISteamUserStats_UpdateAvgRateStat)
	purego.RegisterLibFunc(&iSteamUserStats_GetUserStatFloat, steamAPILib, flatAPI_ISteamUserStats_GetUserStatFloat)
	purego.RegisterLibFunc(&iSteamUserStats_GetMostAchievedAchievementInfo, steamAPILib, flatAPI_ISteamUserStats_GetMostAchievedAchievementInfo)
	purego.RegisterLibFunc(&iSteamUserStats_GetNextMostAchievedAchievementInfo, steamAPILib, flatAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo)
	purego.RegisterLibFunc(&iSteamUserStats_GetAchievementAchievedPercent, steamAPILib, flatAPI_ISteamUserStats_GetAchievementAchievedPercent)
	purego.RegisterLibFunc(&iSteamUserStats_GetAchievementProgressLimitsFloat, steamAPILib, flatAPI_ISteamUserStats_GetAchievementProgressLimitsFloat)
}

var (
	addr_steamUtils_get                              uintptr
	addr_iSteamUtils_GetSecondsSinceAppActive        uintptr
	addr_iSteamUtils_GetSecondsSinceComputerActive   uintptr
	addr_iSteamUtils_GetConnectedUniverse            uintptr
	addr_iSteamUtils_GetServerRealTime               uintptr
	addr_iSteamUtils_GetIPCountry                    uintptr
	addr_iSteamUtils_GetImageSize                    uintptr
	addr_iSteamUtils_GetImageRGBA                    uintptr
	addr_iSteamUtils_GetCurrentBatteryPower          uintptr
	addr_iSteamUtils_GetAppID                        uintptr
	addr_iSteamUtils_SetOverlayNotificationPosition  uintptr
	addr_iSteamUtils_IsAPICallCompleted              uintptr
	addr_iSteamUtils_GetAPICallFailureReason         uintptr
	addr_iSteamUtils_GetAPICallResult                uintptr
	addr_iSteamUtils_GetIPCCallCount                 uintptr
	addr_iSteamUtils_SetWarningMessageHook           uintptr
	addr_iSteamUtils_IsOverlayEnabled                uintptr
	addr_iSteamUtils_BOverlayNeedsPresent            uintptr
	addr_iSteamUtils_CheckFileSignature              uintptr
	addr_iSteamUtils_ShowGamepadTextInput            uintptr
	addr_iSteamUtils_GetEnteredGamepadTextLength     uintptr
	addr_iSteamUtils_GetEnteredGamepadTextInput      uintptr
	addr_iSteamUtils_GetSteamUILanguage              uintptr
	addr_iSteamUtils_IsSteamRunningInVR              uintptr
	addr_iSteamUtils_SetOverlayNotificationInset     uintptr
	addr_iSteamUtils_IsSteamInBigPictureMode         uintptr
	addr_iSteamUtils_StartVRDashboard                uintptr
	addr_iSteamUtils_IsVRHeadsetStreamingEnabled     uintptr
	addr_iSteamUtils_SetVRHeadsetStreamingEnabled    uintptr
	addr_iSteamUtils_IsSteamChinaLauncher            uintptr
	addr_iSteamUtils_InitFilterText                  uintptr
	addr_iSteamUtils_FilterText                      uintptr
	addr_iSteamUtils_GetIPv6ConnectivityState        uintptr
	addr_iSteamUtils_IsSteamRunningOnSteamDeck       uintptr
	addr_iSteamUtils_ShowFloatingGamepadTextInput    uintptr
	addr_iSteamUtils_SetGameLauncherMode             uintptr
	addr_iSteamUtils_DismissFloatingGamepadTextInput uintptr
	addr_iSteamUtils_DismissGamepadTextInput         uintptr
)

func initUtils() {
	var err error
	addr_steamUtils_get, err = openSymbol(steamAPILib, flatAPI_SteamUtils)
	if err != nil {
		panic("cannot OpenSymbol: " + flatAPI_SteamUtils)
	}
	addr_iSteamUtils_GetSecondsSinceAppActive, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetSecondsSinceAppActive")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetSecondsSinceAppActive")
	}
	addr_iSteamUtils_GetSecondsSinceComputerActive, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetSecondsSinceComputerActive")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetSecondsSinceComputerActive")
	}
	addr_iSteamUtils_GetConnectedUniverse, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetConnectedUniverse")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetConnectedUniverse")
	}
	addr_iSteamUtils_GetServerRealTime, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetServerRealTime")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetServerRealTime")
	}
	addr_iSteamUtils_GetIPCountry, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetIPCountry")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetIPCountry")
	}
	addr_iSteamUtils_GetImageSize, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetImageSize")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetImageSize")
	}
	addr_iSteamUtils_GetImageRGBA, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetImageRGBA")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetImageRGBA")
	}
	addr_iSteamUtils_GetCurrentBatteryPower, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetCurrentBatteryPower")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetCurrentBatteryPower")
	}
	addr_iSteamUtils_GetAppID, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetAppID")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetAppID")
	}
	addr_iSteamUtils_SetOverlayNotificationPosition, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_SetOverlayNotificationPosition")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_SetOverlayNotificationPosition")
	}
	addr_iSteamUtils_IsAPICallCompleted, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsAPICallCompleted")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsAPICallCompleted")
	}
	addr_iSteamUtils_GetAPICallFailureReason, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetAPICallFailureReason")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetAPICallFailureReason")
	}
	addr_iSteamUtils_GetAPICallResult, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetAPICallResult")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetAPICallResult")
	}
	addr_iSteamUtils_GetIPCCallCount, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetIPCCallCount")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetIPCCallCount")
	}
	addr_iSteamUtils_SetWarningMessageHook, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_SetWarningMessageHook")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_SetWarningMessageHook")
	}
	addr_iSteamUtils_IsOverlayEnabled, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsOverlayEnabled")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsOverlayEnabled")
	}
	addr_iSteamUtils_BOverlayNeedsPresent, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_BOverlayNeedsPresent")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_BOverlayNeedsPresent")
	}
	addr_iSteamUtils_CheckFileSignature, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_CheckFileSignature")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_CheckFileSignature")
	}
	addr_iSteamUtils_ShowGamepadTextInput, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_ShowGamepadTextInput")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_ShowGamepadTextInput")
	}
	addr_iSteamUtils_GetEnteredGamepadTextLength, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetEnteredGamepadTextLength")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetEnteredGamepadTextLength")
	}
	addr_iSteamUtils_GetEnteredGamepadTextInput, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetEnteredGamepadTextInput")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetEnteredGamepadTextInput")
	}
	addr_iSteamUtils_GetSteamUILanguage, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetSteamUILanguage")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetSteamUILanguage")
	}
	addr_iSteamUtils_IsSteamRunningInVR, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsSteamRunningInVR")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsSteamRunningInVR")
	}
	addr_iSteamUtils_SetOverlayNotificationInset, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_SetOverlayNotificationInset")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_SetOverlayNotificationInset")
	}
	addr_iSteamUtils_IsSteamInBigPictureMode, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsSteamInBigPictureMode")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsSteamInBigPictureMode")
	}
	addr_iSteamUtils_StartVRDashboard, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_StartVRDashboard")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_StartVRDashboard")
	}
	addr_iSteamUtils_IsVRHeadsetStreamingEnabled, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled")
	}
	addr_iSteamUtils_SetVRHeadsetStreamingEnabled, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled")
	}
	addr_iSteamUtils_IsSteamChinaLauncher, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsSteamChinaLauncher")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsSteamChinaLauncher")
	}
	addr_iSteamUtils_InitFilterText, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_InitFilterText")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_InitFilterText")
	}
	addr_iSteamUtils_FilterText, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_FilterText")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_FilterText")
	}
	addr_iSteamUtils_GetIPv6ConnectivityState, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_GetIPv6ConnectivityState")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_GetIPv6ConnectivityState")
	}
	addr_iSteamUtils_IsSteamRunningOnSteamDeck, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck")
	}
	addr_iSteamUtils_ShowFloatingGamepadTextInput, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput")
	}
	addr_iSteamUtils_SetGameLauncherMode, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_SetGameLauncherMode")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_SetGameLauncherMode")
	}
	addr_iSteamUtils_DismissFloatingGamepadTextInput, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_DismissFloatingGamepadTextInput")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_DismissFloatingGamepadTextInput")
	}
	addr_iSteamUtils_DismissGamepadTextInput, err = openSymbol(steamAPILib, "SteamAPI_ISteamUtils_DismissGamepadTextInput")
	if err != nil {
		panic("cannot OpenSymbol: SteamAPI_ISteamUtils_DismissGamepadTextInput")
	}

	steamUtils_get = func() uintptr {
		_r0, _, _ := purego.SyscallN(addr_steamUtils_get)
		__r0 := uintptr(_r0)
		return __r0
	}
	iSteamUtils_GetSecondsSinceAppActive = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetSecondsSinceAppActive, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_GetSecondsSinceComputerActive = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetSecondsSinceComputerActive, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_GetConnectedUniverse = func(steamUtils uintptr) EUniverse {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetConnectedUniverse, uintptr(steamUtils))
		__r0 := EUniverse(_r0)
		return __r0
	}
	iSteamUtils_GetServerRealTime = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetServerRealTime, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_GetIPCountry = func(steamUtils uintptr) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetIPCountry, uintptr(steamUtils))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamUtils_GetImageSize = func(steamUtils uintptr, iImage int32, pnWidth *uint32, pnHeight *uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetImageSize, uintptr(steamUtils), uintptr(iImage), uintptr(unsafe.Pointer(pnWidth)), uintptr(unsafe.Pointer(pnHeight)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pnWidth)
		runtime.KeepAlive(pnHeight)
		return __r0
	}
	iSteamUtils_GetImageRGBA = func(steamUtils uintptr, iImage int32, pubDest []byte, nDestBufferSize int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetImageRGBA, uintptr(steamUtils), uintptr(iImage), uintptr(unsafe.Pointer(unsafe.SliceData(pubDest))), uintptr(nDestBufferSize))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pubDest)
		return __r0
	}
	iSteamUtils_GetCurrentBatteryPower = func(steamUtils uintptr) uint8 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetCurrentBatteryPower, uintptr(steamUtils))
		__r0 := uint8(_r0)
		return __r0
	}
	iSteamUtils_GetAppID = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetAppID, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_SetOverlayNotificationPosition = func(steamUtils uintptr, eNotificationPosition ENotificationPosition) {
		purego.SyscallN(addr_iSteamUtils_SetOverlayNotificationPosition, uintptr(steamUtils), uintptr(eNotificationPosition))
	}
	iSteamUtils_IsAPICallCompleted = func(steamUtils uintptr, hSteamAPICall SteamAPICall, pbFailed *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsAPICallCompleted, uintptr(steamUtils), uintptr(hSteamAPICall), uintptr(unsafe.Pointer(pbFailed)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pbFailed)
		return __r0
	}
	iSteamUtils_GetAPICallFailureReason = func(steamUtils uintptr, hSteamAPICall SteamAPICall) ESteamAPICallFailure {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetAPICallFailureReason, uintptr(steamUtils), uintptr(hSteamAPICall))
		__r0 := ESteamAPICallFailure(_r0)
		return __r0
	}
	iSteamUtils_GetAPICallResult = func(steamUtils uintptr, hSteamAPICall SteamAPICall, pCallback []byte, cubCallback int32, iCallbackExpected int32, pbFailed *bool) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetAPICallResult, uintptr(steamUtils), uintptr(hSteamAPICall), uintptr(unsafe.Pointer(unsafe.SliceData(pCallback))), uintptr(cubCallback), uintptr(iCallbackExpected), uintptr(unsafe.Pointer(pbFailed)))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pCallback)
		runtime.KeepAlive(pbFailed)
		return __r0
	}
	iSteamUtils_GetIPCCallCount = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetIPCCallCount, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_SetWarningMessageHook = func(steamUtils uintptr, pFunction SteamAPIWarningMessageHook) {
		wrapper := func(severity int32, msg uintptr) {
			convertedString := goString(msg)
			pFunction(severity, convertedString)
		}
		cb := purego.NewCallback(wrapper)
		purego.SyscallN(addr_iSteamUtils_SetWarningMessageHook, uintptr(steamUtils), uintptr(cb))
	}
	iSteamUtils_IsOverlayEnabled = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsOverlayEnabled, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_BOverlayNeedsPresent = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_BOverlayNeedsPresent, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_CheckFileSignature = func(steamUtils uintptr, szFileName string) SteamAPICall {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_CheckFileSignature, uintptr(steamUtils), uintptr(unsafe.Pointer(bytePtrFromString(szFileName))))
		__r0 := SteamAPICall(_r0)
		runtime.KeepAlive(szFileName)
		return __r0
	}
	iSteamUtils_ShowGamepadTextInput = func(steamUtils uintptr, eInputMode EGamepadTextInputMode, eLineInputMode EGamepadTextInputLineMode, pchDescription string, unCharMax uint32, pchExistingText string) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_ShowGamepadTextInput, uintptr(steamUtils), uintptr(eInputMode), uintptr(eLineInputMode), uintptr(unsafe.Pointer(bytePtrFromString(pchDescription))), uintptr(unCharMax), uintptr(unsafe.Pointer(bytePtrFromString(pchExistingText))))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchDescription)
		runtime.KeepAlive(pchExistingText)
		return __r0
	}
	iSteamUtils_GetEnteredGamepadTextLength = func(steamUtils uintptr) uint32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetEnteredGamepadTextLength, uintptr(steamUtils))
		__r0 := uint32(_r0)
		return __r0
	}
	iSteamUtils_GetEnteredGamepadTextInput = func(steamUtils uintptr, pchText []byte, cchText uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetEnteredGamepadTextInput, uintptr(steamUtils), uintptr(unsafe.Pointer(unsafe.SliceData(pchText))), uintptr(cchText))
		__r0 := boolFromUintptr(_r0)
		runtime.KeepAlive(pchText)
		return __r0
	}
	iSteamUtils_GetSteamUILanguage = func(steamUtils uintptr) string {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetSteamUILanguage, uintptr(steamUtils))
		__r0 := goString(_r0)
		return __r0
	}
	iSteamUtils_IsSteamRunningInVR = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsSteamRunningInVR, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_SetOverlayNotificationInset = func(steamUtils uintptr, nHorizontalInset int32, nVerticalInset int32) {
		purego.SyscallN(addr_iSteamUtils_SetOverlayNotificationInset, uintptr(steamUtils), uintptr(nHorizontalInset), uintptr(nVerticalInset))
	}
	iSteamUtils_IsSteamInBigPictureMode = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsSteamInBigPictureMode, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_StartVRDashboard = func(steamUtils uintptr) {
		purego.SyscallN(addr_iSteamUtils_StartVRDashboard, uintptr(steamUtils))
	}
	iSteamUtils_IsVRHeadsetStreamingEnabled = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsVRHeadsetStreamingEnabled, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_SetVRHeadsetStreamingEnabled = func(steamUtils uintptr, bEnabled bool) {
		purego.SyscallN(addr_iSteamUtils_SetVRHeadsetStreamingEnabled, uintptr(steamUtils), boolToUintptr(bEnabled))
	}
	iSteamUtils_IsSteamChinaLauncher = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsSteamChinaLauncher, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_InitFilterText = func(steamUtils uintptr, unFilterOptions uint32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_InitFilterText, uintptr(steamUtils), uintptr(unFilterOptions))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_FilterText = func(steamUtils uintptr, eContext ETextFilteringContext, sourceSteamID Uint64SteamID, pchInputMessage string, pchOutFilteredText []byte, nByteSizeOutFilteredText uint32) int32 {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_FilterText, uintptr(steamUtils), uintptr(eContext), uintptr(sourceSteamID), uintptr(unsafe.Pointer(bytePtrFromString(pchInputMessage))), uintptr(unsafe.Pointer(unsafe.SliceData(pchOutFilteredText))), uintptr(nByteSizeOutFilteredText))
		__r0 := int32(_r0)
		runtime.KeepAlive(pchInputMessage)
		runtime.KeepAlive(pchOutFilteredText)
		return __r0
	}
	iSteamUtils_GetIPv6ConnectivityState = func(steamUtils uintptr, eProtocol ESteamIPv6ConnectivityProtocol) ESteamIPv6ConnectivityState {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_GetIPv6ConnectivityState, uintptr(steamUtils), uintptr(eProtocol))
		__r0 := ESteamIPv6ConnectivityState(_r0)
		return __r0
	}
	iSteamUtils_IsSteamRunningOnSteamDeck = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_IsSteamRunningOnSteamDeck, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_ShowFloatingGamepadTextInput = func(steamUtils uintptr, eKeyboardMode EFloatingGamepadTextInputMode, nTextFieldXPosition int32, nTextFieldYPosition int32, nTextFieldWidth int32, nTextFieldHeight int32) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_ShowFloatingGamepadTextInput, uintptr(steamUtils), uintptr(eKeyboardMode), uintptr(nTextFieldXPosition), uintptr(nTextFieldYPosition), uintptr(nTextFieldWidth), uintptr(nTextFieldHeight))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_SetGameLauncherMode = func(steamUtils uintptr, bLauncherMode bool) {
		purego.SyscallN(addr_iSteamUtils_SetGameLauncherMode, uintptr(steamUtils), boolToUintptr(bLauncherMode))
	}
	iSteamUtils_DismissFloatingGamepadTextInput = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_DismissFloatingGamepadTextInput, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
	iSteamUtils_DismissGamepadTextInput = func(steamUtils uintptr) bool {
		_r0, _, _ := purego.SyscallN(addr_iSteamUtils_DismissGamepadTextInput, uintptr(steamUtils))
		__r0 := boolFromUintptr(_r0)
		return __r0
	}
}
