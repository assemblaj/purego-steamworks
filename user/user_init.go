package user

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamUser_init, SteamAPIDLL, flatAPI_SteamUser)
	purego.RegisterLibFunc(&iSteamUser_GetHSteamUser, SteamAPIDLL, flatAPI_ISteamUser_GetHSteamUser)
	purego.RegisterLibFunc(&iSteamUser_BLoggedOn, SteamAPIDLL, flatAPI_ISteamUser_BLoggedOn)
	purego.RegisterLibFunc(&iSteamUser_GetSteamID, SteamAPIDLL, flatAPI_ISteamUser_GetSteamID)
	purego.RegisterLibFunc(&iSteamUser_TrackAppUsageEvent, SteamAPIDLL, flatAPI_ISteamUser_TrackAppUsageEvent)
	purego.RegisterLibFunc(&iSteamUser_GetUserDataFolder, SteamAPIDLL, flatAPI_ISteamUser_GetUserDataFolder)
	purego.RegisterLibFunc(&iSteamUser_StartVoiceRecording, SteamAPIDLL, flatAPI_ISteamUser_StartVoiceRecording)
	purego.RegisterLibFunc(&iSteamUser_StopVoiceRecording, SteamAPIDLL, flatAPI_ISteamUser_StopVoiceRecording)
	purego.RegisterLibFunc(&iSteamUser_GetAvailableVoice, SteamAPIDLL, flatAPI_ISteamUser_GetAvailableVoice)
	purego.RegisterLibFunc(&iSteamUser_GetVoice, SteamAPIDLL, flatAPI_ISteamUser_GetVoice)
	purego.RegisterLibFunc(&iSteamUser_DecompressVoice, SteamAPIDLL, flatAPI_ISteamUser_DecompressVoice)
	purego.RegisterLibFunc(&iSteamUser_GetVoiceOptimalSampleRate, SteamAPIDLL, flatAPI_ISteamUser_GetVoiceOptimalSampleRate)
	purego.RegisterLibFunc(&iSteamUser_GetAuthSessionTicket, SteamAPIDLL, flatAPI_ISteamUser_GetAuthSessionTicket)
	purego.RegisterLibFunc(&iSteamUser_GetAuthTicketForWebApi, SteamAPIDLL, flatAPI_ISteamUser_GetAuthTicketForWebApi)
	purego.RegisterLibFunc(&iSteamUser_BeginAuthSession, SteamAPIDLL, flatAPI_ISteamUser_BeginAuthSession)
	purego.RegisterLibFunc(&iSteamUser_EndAuthSession, SteamAPIDLL, flatAPI_ISteamUser_EndAuthSession)
	purego.RegisterLibFunc(&iSteamUser_CancelAuthTicket, SteamAPIDLL, flatAPI_ISteamUser_CancelAuthTicket)
	purego.RegisterLibFunc(&iSteamUser_UserHasLicenseForApp, SteamAPIDLL, flatAPI_ISteamUser_UserHasLicenseForApp)
	purego.RegisterLibFunc(&iSteamUser_BIsBehindNAT, SteamAPIDLL, flatAPI_ISteamUser_BIsBehindNAT)
	purego.RegisterLibFunc(&iSteamUser_AdvertiseGame, SteamAPIDLL, flatAPI_ISteamUser_AdvertiseGame)
	purego.RegisterLibFunc(&iSteamUser_RequestEncryptedAppTicket, SteamAPIDLL, flatAPI_ISteamUser_RequestEncryptedAppTicket)
	purego.RegisterLibFunc(&iSteamUser_GetEncryptedAppTicket, SteamAPIDLL, flatAPI_ISteamUser_GetEncryptedAppTicket)
	purego.RegisterLibFunc(&iSteamUser_GetGameBadgeLevel, SteamAPIDLL, flatAPI_ISteamUser_GetGameBadgeLevel)
	purego.RegisterLibFunc(&iSteamUser_GetPlayerSteamLevel, SteamAPIDLL, flatAPI_ISteamUser_GetPlayerSteamLevel)
	purego.RegisterLibFunc(&iSteamUser_RequestStoreAuthURL, SteamAPIDLL, flatAPI_ISteamUser_RequestStoreAuthURL)
	purego.RegisterLibFunc(&iSteamUser_BIsPhoneVerified, SteamAPIDLL, flatAPI_ISteamUser_BIsPhoneVerified)
	purego.RegisterLibFunc(&iSteamUser_BIsTwoFactorEnabled, SteamAPIDLL, flatAPI_ISteamUser_BIsTwoFactorEnabled)
	purego.RegisterLibFunc(&iSteamUser_BIsPhoneIdentifying, SteamAPIDLL, flatAPI_ISteamUser_BIsPhoneIdentifying)
	purego.RegisterLibFunc(&iSteamUser_BIsPhoneRequiringVerification, SteamAPIDLL, flatAPI_ISteamUser_BIsPhoneRequiringVerification)
	purego.RegisterLibFunc(&iSteamUser_GetMarketEligibility, SteamAPIDLL, flatAPI_ISteamUser_GetMarketEligibility)
	purego.RegisterLibFunc(&iSteamUser_GetDurationControl, SteamAPIDLL, flatAPI_ISteamUser_GetDurationControl)
	purego.RegisterLibFunc(&iSteamUser_BSetDurationControlOnlineState, SteamAPIDLL, flatAPI_ISteamUser_BSetDurationControlOnlineState)

	steamUser = steamUser_init()
}
