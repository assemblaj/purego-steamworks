package puregosteamworks

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

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

func init() {
	var err error
	SteamAPIDLL, err = OpenLibrary(getSteamLibrary())
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&restartAppiIfNecessary, SteamAPIDLL, flatAPI_RestartAppIfNecessary)
	purego.RegisterLibFunc(&initFlat, SteamAPIDLL, flatAPI_InitFlat)
	purego.RegisterLibFunc(&Shutdown, SteamAPIDLL, flatAPI_Shutdown)
	purego.RegisterLibFunc(&WriteMiniDump, SteamAPIDLL, flatAPI_WriteMiniDump)
	purego.RegisterLibFunc(&SetMiniDumpComment, SteamAPIDLL, flatAPI_SetMiniDumpComment)
	purego.RegisterLibFunc(&IsSteamRunning, SteamAPIDLL, flatAPI_IsSteamRunning)
	purego.RegisterLibFunc(&ReleaseCurrentThreadMemory, SteamAPIDLL, flatAPI_ReleaseCurrentThreadMemory)

	purego.RegisterLibFunc(&gameServerInit, SteamAPIDLL, flatAPI_SteamGameServer_Init)
	purego.RegisterLibFunc(&GameServerShutdown, SteamAPIDLL, flatAPI_SteamGameServer_Shutdown)
	purego.RegisterLibFunc(&GameServerBSecure, SteamAPIDLL, flatAPI_SteamGameServer_BSecure)
	purego.RegisterLibFunc(&GameServerGetSteamID, SteamAPIDLL, flatAPI_SteamGameServer_GetSteamID)

}
