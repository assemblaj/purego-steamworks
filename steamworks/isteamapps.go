package steamworks

import (
	"unsafe"
)

const (
	flatAPI_SteamApps                         = "SteamAPI_SteamApps_v008"
	flatAPI_ISteamApps_BGetDLCDataByIndex     = "SteamAPI_ISteamApps_BGetDLCDataByIndex"
	flatAPI_ISteamApps_BIsDlcInstalled        = "SteamAPI_ISteamApps_BIsDlcInstalled"
	flatAPI_ISteamApps_GetAppInstallDir       = "SteamAPI_ISteamApps_GetAppInstallDir"
	flatAPI_ISteamApps_GetCurrentGameLanguage = "SteamAPI_ISteamApps_GetCurrentGameLanguage"
	flatAPI_ISteamApps_GetDLCCount            = "SteamAPI_ISteamApps_GetDLCCount"
)

type steamapps uintptr

var (
	steamapps_init         func() uintptr
	bGetDLCDataByIndex     func(steamapps uintptr, iDLC int, appID uintptr, available uintptr, pchName uintptr, nameLen int) uintptr
	bIsDlcInstalled        func(steamapps uintptr, appID AppId_t) uintptr
	getAppInstallDir       func(steamapps uintptr, appID AppId_t, path uintptr, pathLen int) uintptr
	getCurrentGameLanguage func(steamapps uintptr) string
	getDLCCount            func(steamapps uintptr) uintptr
)

func (s steamapps) BGetDLCDataByIndex(iDLC int) (appID AppId_t, available bool, pchName string, success bool) {
	var name [4096]byte
	data := bGetDLCDataByIndex(uintptr(s), iDLC, uintptr(unsafe.Pointer(&appID)), uintptr(unsafe.Pointer(&available)), uintptr(unsafe.Pointer(&name[0])), len(name))
	return appID, available, cStringToGoString(data, len(name)), byte(data) != 0
}

func (s steamapps) BIsDlcInstalled(appID AppId_t) bool {
	result := bIsDlcInstalled(uintptr(s), appID)
	return byte(result) != 0
}

func (s steamapps) GetAppInstallDir(appID AppId_t) string {
	var path [4096]byte
	result := getAppInstallDir(uintptr(s), appID, uintptr(unsafe.Pointer(&path[0])), len(path))
	return string(path[:uint32(result)-1])
}

func (s steamapps) GetCurrentGameLanguage() string {
	return getCurrentGameLanguage(uintptr(s))
}

func (s steamapps) GetDLCCount() int32 {
	result := getDLCCount(uintptr(s))
	return int32(result)
}
