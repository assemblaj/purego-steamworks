package steamworks

import (
	"fmt"
	"runtime"
)

type AppId_t uint32
type CSteamID uint64

type ESteamAPIInitResult int32

const (
	ESteamAPIInitResult_OK              ESteamAPIInitResult = 0
	ESteamAPIInitResult_FailedGeneric   ESteamAPIInitResult = 1
	ESteamAPIInitResult_NoSteamClient   ESteamAPIInitResult = 2
	ESteamAPIInitResult_VersionMismatch ESteamAPIInitResult = 3
)

var steam uintptr

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
