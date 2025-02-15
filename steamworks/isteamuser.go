package steamworks

import "unsafe"

const is32Bit = unsafe.Sizeof(int(0)) == 4

const (
	flatAPI_SteamUser             = "SteamAPI_SteamUser_v023"
	flatAPI_ISteamUser_GetSteamID = "SteamAPI_ISteamUser_GetSteamID"
)

type steamUser uintptr

var (
	steamUser_init func() uintptr
	getSteamID     func(steamUser uintptr) uintptr
)

func (s steamUser) GetSteamID() CSteamID {
	if is32Bit {
		// On 32bit machines, syscall cannot treat a returned value as 64bit.
		panic("GetSteamID is not implemented on 32bit Windows")
	}
	result := getSteamID(uintptr(s))
	return CSteamID(result)
}
