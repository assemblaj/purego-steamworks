package steamworks

import (
	"fmt"
)

const (
	flatAPI_RestartAppIfNecessary = "SteamAPI_RestartAppIfNecessary"
	flatAPI_InitFlat              = "SteamAPI_InitFlat"
	flatAPI_RunCallbacks          = "SteamAPI_RunCallbacks"
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
	restartAppiIfNecessary func(int) uintptr
	initFlat               func(*steamErrMsg) uintptr
	RunCallbacks           func()
)

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
