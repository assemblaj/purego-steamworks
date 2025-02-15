package steamworks

import (
	"runtime"
	"unsafe"
)

const (
	flatAPI_SteamFriends                  = "SteamAPI_SteamFriends_v017"
	flatAPI_ISteamFriends_GetPersonaName  = "SteamAPI_ISteamFriends_GetPersonaName"
	flatAPI_ISteamFriends_SetRichPresence = "SteamAPI_ISteamFriends_SetRichPresence"
)

type steamFriends uintptr

var (
	steamFriends_init func() uintptr
	getPersonaName    func(steamFriends uintptr) string
	setRichPrescence  func(steamFriends uintptr, key uintptr, value uintptr) uintptr
)

func (s steamFriends) GetPersonaName() string {
	return getPersonaName(uintptr(s))
}

func (s steamFriends) SetRichPresence(key, value string) bool {
	ckey := append([]byte(key), 0)
	defer runtime.KeepAlive(ckey)
	cvalue := append([]byte(value), 0)
	defer runtime.KeepAlive(cvalue)

	result := setRichPrescence(uintptr(s), uintptr(unsafe.Pointer(&ckey[0])), uintptr(unsafe.Pointer(&cvalue[0])))
	return byte(result) != 0
}
