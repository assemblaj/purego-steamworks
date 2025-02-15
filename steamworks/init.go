package steamworks

import "github.com/ebitengine/purego"

func init() {
	initApi()
	initSteamApps()
	initSteamFriends()
	initSteamInput()
	initSteamRemoteStorage()
	initSteamUser()
	initSteamUserStats()
	initSteamUtils()
}

func initApi() {
	var err error
	steam, err = openLibrary(getSteamLibrary())
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&restartAppiIfNecessary, steam, flatAPI_RestartAppIfNecessary)
	purego.RegisterLibFunc(&initFlat, steam, flatAPI_InitFlat)
	purego.RegisterLibFunc(&RunCallbacks, steam, flatAPI_RunCallbacks)
}

func SteamApps() steamapps {
	purego.RegisterLibFunc(&steamapps_init, steam, flatAPI_SteamApps)
	return steamapps(steamapps_init())
}

func initSteamApps() {
	purego.RegisterLibFunc(&bGetDLCDataByIndex, steam, flatAPI_ISteamApps_BGetDLCDataByIndex)
	purego.RegisterLibFunc(&bIsDlcInstalled, steam, flatAPI_ISteamApps_BIsDlcInstalled)
	purego.RegisterLibFunc(&getAppInstallDir, steam, flatAPI_ISteamApps_GetAppInstallDir)
	purego.RegisterLibFunc(&getCurrentGameLanguage, steam, flatAPI_ISteamApps_GetCurrentGameLanguage)
	purego.RegisterLibFunc(&getDLCCount, steam, flatAPI_ISteamApps_GetDLCCount)
}

func SteamFriends() steamFriends {
	purego.RegisterLibFunc(&steamFriends_init, steam, flatAPI_SteamFriends)
	return steamFriends(steamFriends_init())
}

func initSteamFriends() {
	purego.RegisterLibFunc(&getPersonaName, steam, flatAPI_ISteamFriends_GetPersonaName)
	purego.RegisterLibFunc(&setRichPrescence, steam, flatAPI_ISteamFriends_SetRichPresence)

}

func SteamInput() steamInput {
	purego.RegisterLibFunc(steamInput_init, steam, flatAPI_SteamInput)
	return steamInput(steamInput_init())
}

func initSteamInput() {
	purego.RegisterLibFunc(&getConnectedControllers, steam, flatAPI_ISteamInput_GetConnectedControllers)
	purego.RegisterLibFunc(&getInputTypeForHandle, steam, flatAPI_ISteamInput_GetInputTypeForHandle)
	purego.RegisterLibFunc(&steamInputInit, steam, flatAPI_ISteamInput_Init)
	purego.RegisterLibFunc(&steamInputRunFrame, steam, flatAPI_ISteamInput_RunFrame)
}

func SteamRemoteStorage() steamRemoteStorage {
	purego.RegisterLibFunc(steamRemoteStorage_init, steam, flatAPI_SteamRemoteStorage)
	return steamRemoteStorage(steamRemoteStorage_init())
}

func initSteamRemoteStorage() {
	purego.RegisterLibFunc(&fileWrite, steam, flatAPI_ISteamRemoteStorage_FileWrite)
	purego.RegisterLibFunc(&fileRead, steam, flatAPI_ISteamRemoteStorage_FileRead)
	purego.RegisterLibFunc(&fileDelete, steam, flatAPI_ISteamRemoteStorage_FileDelete)
	purego.RegisterLibFunc(&getFileSize, steam, flatAPI_ISteamRemoteStorage_GetFileSize)
}

func SteamUser() steamUser {
	purego.RegisterLibFunc(steamUser_init, steam, flatAPI_SteamUser)
	return steamUser(steamUser_init())
}

func initSteamUser() {
	purego.RegisterLibFunc(&getSteamID, steam, flatAPI_ISteamUser_GetSteamID)
}

func SteamUserStats() steamUserStats {
	purego.RegisterLibFunc(steamUserStats_init, steam, flatAPI_SteamUserStats)
	return steamUserStats(steamUserStats_init())
}

func initSteamUserStats() {
	purego.RegisterLibFunc(getAchievement, steam, flatAPI_ISteamUserStats_GetAchievement)
	purego.RegisterLibFunc(setAchievement, steam, flatAPI_ISteamUserStats_SetAchievement)
	purego.RegisterLibFunc(clearAchievement, steam, flatAPI_ISteamUserStats_ClearAchievement)
	purego.RegisterLibFunc(storeStats, steam, flatAPI_ISteamUserStats_StoreStats)
}

func SteamUtils() steamUtils {
	purego.RegisterLibFunc(steamUitils_init, steam, flatAPI_SteamUtils)
	return steamUtils(steamUitils_init())
}

func initSteamUtils() {
	purego.RegisterLibFunc(isSteamRunningOnSteamDeck, steam, flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck)
	purego.RegisterLibFunc(showFloatingGamepadTextInput, steam, flatAPI_ISteamUtils_ShowFloatingGamepadTextInput)
}
