package parties

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamParties_init, SteamAPIDLL, flatAPI_SteamParties)
	purego.RegisterLibFunc(&iSteamParties_GetNumActiveBeacons, SteamAPIDLL, flatAPI_ISteamParties_GetNumActiveBeacons)
	purego.RegisterLibFunc(&iSteamParties_GetBeaconByIndex, SteamAPIDLL, flatAPI_ISteamParties_GetBeaconByIndex)
	purego.RegisterLibFunc(&iSteamParties_GetBeaconDetails, SteamAPIDLL, flatAPI_ISteamParties_GetBeaconDetails)
	purego.RegisterLibFunc(&iSteamParties_JoinParty, SteamAPIDLL, flatAPI_ISteamParties_JoinParty)
	purego.RegisterLibFunc(&iSteamParties_GetNumAvailableBeaconLocations, SteamAPIDLL, flatAPI_ISteamParties_GetNumAvailableBeaconLocations)
	purego.RegisterLibFunc(&iSteamParties_GetAvailableBeaconLocations, SteamAPIDLL, flatAPI_ISteamParties_GetAvailableBeaconLocations)
	purego.RegisterLibFunc(&iSteamParties_CreateBeacon, SteamAPIDLL, flatAPI_ISteamParties_CreateBeacon)
	purego.RegisterLibFunc(&iSteamParties_OnReservationCompleted, SteamAPIDLL, flatAPI_ISteamParties_OnReservationCompleted)
	purego.RegisterLibFunc(&iSteamParties_CancelReservation, SteamAPIDLL, flatAPI_ISteamParties_CancelReservation)
	purego.RegisterLibFunc(&iSteamParties_ChangeNumOpenSlots, SteamAPIDLL, flatAPI_ISteamParties_ChangeNumOpenSlots)
	purego.RegisterLibFunc(&iSteamParties_DestroyBeacon, SteamAPIDLL, flatAPI_ISteamParties_DestroyBeacon)
	purego.RegisterLibFunc(&iSteamParties_GetBeaconLocationData, SteamAPIDLL, flatAPI_ISteamParties_GetBeaconLocationData)
}
