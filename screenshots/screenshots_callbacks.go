package screenshots

import . "github.com/assemblaj/purego-steamworks"

const (
	ScreenshotReadyID     SteamCallbackID = 2301
	ScreenshotRequestedID SteamCallbackID = 2302
)

type ScreenshotReady struct {
	Local  ScreenshotHandle
	Result EResult
}
type ScreenshotRequested struct {
}

func (cb ScreenshotReady) CallbackID() SteamCallbackID {
	return ScreenshotReadyID
}

func (cb ScreenshotReady) Name() string {
	return "ScreenshotReady"
}

func (cb ScreenshotReady) String() string {
	return CallbackString(cb)
}

func (cb ScreenshotRequested) CallbackID() SteamCallbackID {
	return ScreenshotRequestedID
}

func (cb ScreenshotRequested) Name() string {
	return "ScreenshotRequested"
}

func (cb ScreenshotRequested) String() string {
	return CallbackString(cb)
}
