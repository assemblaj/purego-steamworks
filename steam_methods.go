package puregosteamworks

import (
	"fmt"
)

const (
	ESteamAPIInitResult_OK              = 0
	ESteamAPIInitResult_FailedGeneric   = 1 // Some other failure
	ESteamAPIInitResult_NoSteamClient   = 2 // We cannot connect to Steam, steam probably isn't running
	ESteamAPIInitResult_VersionMismatch = 3 // Steam client appears to be out of date
)

type callbackMsg struct {
	SteamUser HSteamUser // Specific user to whom this callback applies.
	Callback  int32      // Callback identifier.  (Corresponds to the k_iCallback enum in the callback structure.)
	ParamData uintptr    // Points to the callback structure
	ParamSize int32      // Size of the data pointed to by m_pubParam
}

const (
	flatAPI_RestartAppIfNecessary      = "SteamAPI_RestartAppIfNecessary"
	flatAPI_InitFlat                   = "SteamAPI_InitFlat"
	flatAPI_Shutdown                   = "SteamAPI_Shutdown"
	flatAPI_SetMiniDumpComment         = "SteamAPI_SetMiniDumpComment"
	flatAPI_WriteMiniDump              = "SteamAPI_WriteMiniDump"
	flatAPI_IsSteamRunning             = "SteamAPI_IsSteamRunning"
	flatAPI_ReleaseCurrentThreadMemory = "SteamAPI_ReleaseCurrentThreadMemory"
)

type steamErrMsg [1024]byte

func (s *steamErrMsg) String() string {
	for i, b := range s {
		if b == 0 {
			return string(s[:i])
		}
	}
	return ""
}

var (
	restartAppiIfNecessary     func(int) uintptr
	initFlat                   func(*steamErrMsg) uintptr
	Shutdown                   func()
	WriteMiniDump              func(uStructuredExceptionCode uint32, pvExceptionInfo []byte, uBuildID uint32)
	SetMiniDumpComment         func(pchMsg string)
	IsSteamRunning             func() bool
	ReleaseCurrentThreadMemory func()
)

var debugMode bool = false

func SetDebugMode(debugEnabled bool) {
	debugMode = debugEnabled
}

func RestartAppiIfNecessary(appID uint32) bool {
	result := restartAppiIfNecessary(int(appID))
	return byte(result) != 0
}

func Init() error {
	var msg steamErrMsg
	result := initFlat(&msg)

	if ESteamAPIInitResult(result) != ESteamAPIInitResult_OK {
		return fmt.Errorf("steamworks: InitFlat failed: %d, %s", ESteamAPIInitResult(result), msg.String())
	}
	return nil
}
