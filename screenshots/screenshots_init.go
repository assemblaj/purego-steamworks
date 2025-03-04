package screenshots

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamScreenshots_init, SteamAPIDLL, flatAPI_SteamScreenshots)
	purego.RegisterLibFunc(&iSteamScreenshots_WriteScreenshot, SteamAPIDLL, flatAPI_ISteamScreenshots_WriteScreenshot)
	purego.RegisterLibFunc(&iSteamScreenshots_AddScreenshotToLibrary, SteamAPIDLL, flatAPI_ISteamScreenshots_AddScreenshotToLibrary)
	purego.RegisterLibFunc(&iSteamScreenshots_TriggerScreenshot, SteamAPIDLL, flatAPI_ISteamScreenshots_TriggerScreenshot)
	purego.RegisterLibFunc(&iSteamScreenshots_HookScreenshots, SteamAPIDLL, flatAPI_ISteamScreenshots_HookScreenshots)
	purego.RegisterLibFunc(&iSteamScreenshots_SetLocation, SteamAPIDLL, flatAPI_ISteamScreenshots_SetLocation)
	purego.RegisterLibFunc(&iSteamScreenshots_TagUser, SteamAPIDLL, flatAPI_ISteamScreenshots_TagUser)
	purego.RegisterLibFunc(&iSteamScreenshots_TagPublishedFile, SteamAPIDLL, flatAPI_ISteamScreenshots_TagPublishedFile)
	purego.RegisterLibFunc(&iSteamScreenshots_IsScreenshotsHooked, SteamAPIDLL, flatAPI_ISteamScreenshots_IsScreenshotsHooked)
	purego.RegisterLibFunc(&iSteamScreenshots_AddVRScreenshotToLibrary, SteamAPIDLL, flatAPI_ISteamScreenshots_AddVRScreenshotToLibrary)
}
