package steamworks

import (
	"runtime"
	"unsafe"
)

const (
	flatAPI_SteamUserStats                   = "SteamAPI_SteamUserStats_v013"
	flatAPI_ISteamUserStats_GetAchievement   = "SteamAPI_ISteamUserStats_GetAchievement"
	flatAPI_ISteamUserStats_SetAchievement   = "SteamAPI_ISteamUserStats_SetAchievement"
	flatAPI_ISteamUserStats_ClearAchievement = "SteamAPI_ISteamUserStats_ClearAchievement"
	flatAPI_ISteamUserStats_StoreStats       = "SteamAPI_ISteamUserStats_StoreStats"
)

type steamUserStats uintptr

var (
	steamUserStats_init func() uintptr
	getAchievement      func(steamUserStats uintptr, name uintptr, achieved uintptr) uintptr
	setAchievement      func(steamUserState uintptr, name uintptr) uintptr
	clearAchievement    func(steamUserStats uintptr, name uintptr) uintptr
	storeStats          func(steamUserStats uintptr) uintptr
)

func (s steamUserStats) GetAchievement(name string) (achieved, success bool) {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	result := getAchievement(uintptr(s), uintptr(unsafe.Pointer(&cname[0])), uintptr(unsafe.Pointer(&achieved)))
	success = byte(result) != 0
	return
}

func (s steamUserStats) SetAchievement(name string) bool {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	result := setAchievement(uintptr(s), uintptr(unsafe.Pointer(&cname[0])))
	return byte(result) != 0
}

func (s steamUserStats) ClearAchievement(name string) bool {
	cname := append([]byte(name), 0)
	defer runtime.KeepAlive(cname)

	result := clearAchievement(uintptr(s), uintptr(unsafe.Pointer(&cname[0])))
	return byte(result) != 0
}

func (s steamUserStats) StoreStats() bool {
	result := storeStats(uintptr(s))
	return byte(result) != 0
}
