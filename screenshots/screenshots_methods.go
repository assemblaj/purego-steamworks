package screenshots

import (
	. "github.com/assemblaj/purego-steamworks/remote_storage_types"

	. "github.com/assemblaj/purego-steamworks"
)

const (
	ScreenshotMaxTaggedUsers          uint32 = 32
	ScreenshotMaxTaggedPublishedFiles uint32 = 32
	UFSTagTypeMax                     int32  = 255
	UFSTagValueMax                    int32  = 255
	ScreenshotThumbWidth              int32  = 200
)

const (
	flatAPI_SteamScreenshots                           = "SteamAPI_SteamScreenshots_v003"
	flatAPI_ISteamScreenshots_WriteScreenshot          = "SteamAPI_ISteamScreenshots_WriteScreenshot"
	flatAPI_ISteamScreenshots_AddScreenshotToLibrary   = "SteamAPI_ISteamScreenshots_AddScreenshotToLibrary"
	flatAPI_ISteamScreenshots_TriggerScreenshot        = "SteamAPI_ISteamScreenshots_TriggerScreenshot"
	flatAPI_ISteamScreenshots_HookScreenshots          = "SteamAPI_ISteamScreenshots_HookScreenshots"
	flatAPI_ISteamScreenshots_SetLocation              = "SteamAPI_ISteamScreenshots_SetLocation"
	flatAPI_ISteamScreenshots_TagUser                  = "SteamAPI_ISteamScreenshots_TagUser"
	flatAPI_ISteamScreenshots_TagPublishedFile         = "SteamAPI_ISteamScreenshots_TagPublishedFile"
	flatAPI_ISteamScreenshots_IsScreenshotsHooked      = "SteamAPI_ISteamScreenshots_IsScreenshotsHooked"
	flatAPI_ISteamScreenshots_AddVRScreenshotToLibrary = "SteamAPI_ISteamScreenshots_AddVRScreenshotToLibrary"
)

type ScreenshotHandle uint

type EVRScreenshotType int32

const (
	EVRScreenshotTypeNone           EVRScreenshotType = 0
	EVRScreenshotTypeMono           EVRScreenshotType = 1
	EVRScreenshotTypeStereo         EVRScreenshotType = 2
	EVRScreenshotTypeMonoCubemap    EVRScreenshotType = 3
	EVRScreenshotTypeMonoPanorama   EVRScreenshotType = 4
	EVRScreenshotTypeStereoPanorama EVRScreenshotType = 5
)

var (
	steamScreenshots_init                      func() uintptr
	iSteamScreenshots_WriteScreenshot          func(steamScreenshots uintptr, pubRGB []byte, cubRGB uint32, nWidth int32, nHeight int32) ScreenshotHandle
	iSteamScreenshots_AddScreenshotToLibrary   func(steamScreenshots uintptr, pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle
	iSteamScreenshots_TriggerScreenshot        func(steamScreenshots uintptr)
	iSteamScreenshots_HookScreenshots          func(steamScreenshots uintptr, bHook bool)
	iSteamScreenshots_SetLocation              func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, pchLocation string) bool
	iSteamScreenshots_TagUser                  func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool
	iSteamScreenshots_TagPublishedFile         func(steamScreenshots uintptr, hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool
	iSteamScreenshots_IsScreenshotsHooked      func(steamScreenshots uintptr) bool
	iSteamScreenshots_AddVRScreenshotToLibrary func(steamScreenshots uintptr, eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle
)

var steamScreenshots uintptr

func screenshots() uintptr {
	if steamScreenshots == 0 {
		steamScreenshots = steamScreenshots_init()
		return steamScreenshots
	}
	return steamScreenshots
}

func WriteScreenshot(pubRGB []byte, nWidth int32, nHeight int32) ScreenshotHandle {
	return iSteamScreenshots_WriteScreenshot(screenshots(), pubRGB, uint32(len(pubRGB)), nWidth, nHeight)
}

func AddScreenshotToLibrary(pchFilename string, pchThumbnailFilename string, nWidth int32, nHeight int32) ScreenshotHandle {
	return iSteamScreenshots_AddScreenshotToLibrary(screenshots(), pchFilename, pchThumbnailFilename, nWidth, nHeight)
}

func TriggerScreenshot() {
	iSteamScreenshots_TriggerScreenshot(screenshots())
}

func HookScreenshots(bHook bool) {
	iSteamScreenshots_HookScreenshots(screenshots(), bHook)
}

func SetLocation(hScreenshot ScreenshotHandle, pchLocation string) bool {
	return iSteamScreenshots_SetLocation(screenshots(), hScreenshot, pchLocation)
}

func TagUser(hScreenshot ScreenshotHandle, steamID Uint64SteamID) bool {
	return iSteamScreenshots_TagUser(screenshots(), hScreenshot, steamID)
}

func TagPublishedFile(hScreenshot ScreenshotHandle, unPublishedFileID PublishedFileId) bool {
	return iSteamScreenshots_TagPublishedFile(screenshots(), hScreenshot, unPublishedFileID)
}

func IsScreenshotsHooked() bool {
	return iSteamScreenshots_IsScreenshotsHooked(screenshots())
}

func AddVRScreenshotToLibrary(eType EVRScreenshotType, pchFilename string, pchVRFilename string) ScreenshotHandle {
	return iSteamScreenshots_AddVRScreenshotToLibrary(screenshots(), eType, pchFilename, pchVRFilename)
}
