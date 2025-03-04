package parties

import (
	common "github.com/assemblaj/purego-steamworks"
)

const (
	PartyBeaconIdInvalid PartyBeaconID = 0
)

type ESteamPartyBeaconLocationType int32

const (
	ESteamPartyBeaconLocationTypeInvalid   ESteamPartyBeaconLocationType = 0
	ESteamPartyBeaconLocationTypeChatGroup ESteamPartyBeaconLocationType = 1
	ESteamPartyBeaconLocationTypeMax       ESteamPartyBeaconLocationType = 2
)

type ESteamPartyBeaconLocationData int32

const (
	ESteamPartyBeaconLocationDataInvalid       ESteamPartyBeaconLocationData = 0
	ESteamPartyBeaconLocationDataName          ESteamPartyBeaconLocationData = 1
	ESteamPartyBeaconLocationDataIconURLSmall  ESteamPartyBeaconLocationData = 2
	ESteamPartyBeaconLocationDataIconURLMedium ESteamPartyBeaconLocationData = 3
	ESteamPartyBeaconLocationDataIconURLLarge  ESteamPartyBeaconLocationData = 4
)

type PartyBeaconID uint64
type SteamPartyBeaconLocation struct {
	Type       ESteamPartyBeaconLocationType
	LocationID uint64
}

const (
	flatAPI_SteamParties                                 = "SteamAPI_SteamParties_v002"
	flatAPI_ISteamParties_GetNumActiveBeacons            = "SteamAPI_ISteamParties_GetNumActiveBeacons"
	flatAPI_ISteamParties_GetBeaconByIndex               = "SteamAPI_ISteamParties_GetBeaconByIndex"
	flatAPI_ISteamParties_GetBeaconDetails               = "SteamAPI_ISteamParties_GetBeaconDetails"
	flatAPI_ISteamParties_JoinParty                      = "SteamAPI_ISteamParties_JoinParty"
	flatAPI_ISteamParties_GetNumAvailableBeaconLocations = "SteamAPI_ISteamParties_GetNumAvailableBeaconLocations"
	flatAPI_ISteamParties_GetAvailableBeaconLocations    = "SteamAPI_ISteamParties_GetAvailableBeaconLocations"
	flatAPI_ISteamParties_CreateBeacon                   = "SteamAPI_ISteamParties_CreateBeacon"
	flatAPI_ISteamParties_OnReservationCompleted         = "SteamAPI_ISteamParties_OnReservationCompleted"
	flatAPI_ISteamParties_CancelReservation              = "SteamAPI_ISteamParties_CancelReservation"
	flatAPI_ISteamParties_ChangeNumOpenSlots             = "SteamAPI_ISteamParties_ChangeNumOpenSlots"
	flatAPI_ISteamParties_DestroyBeacon                  = "SteamAPI_ISteamParties_DestroyBeacon"
	flatAPI_ISteamParties_GetBeaconLocationData          = "SteamAPI_ISteamParties_GetBeaconLocationData"
)

var (
	steamParties_init                            func() uintptr
	iSteamParties_GetNumActiveBeacons            func(steamParties uintptr) uint32
	iSteamParties_GetBeaconByIndex               func(steamParties uintptr, unIndex uint32) PartyBeaconID
	iSteamParties_GetBeaconDetails               func(steamParties uintptr, ulBeaconID PartyBeaconID, pSteamIDBeaconOwner *common.CSteamID, pLocation *SteamPartyBeaconLocation, pchMetadata *[]byte, cchMetadata int32) bool
	iSteamParties_JoinParty                      func(steamParties uintptr, ulBeaconID PartyBeaconID) common.SteamAPICall
	iSteamParties_GetNumAvailableBeaconLocations func(steamParties uintptr, puNumLocations *uint32) bool
	iSteamParties_GetAvailableBeaconLocations    func(steamParties uintptr, pLocationList *[]SteamPartyBeaconLocation, uMaxNumLocations uint32) bool
	iSteamParties_CreateBeacon                   func(steamParties uintptr, unOpenSlots uint32, pBeaconLocation uintptr, pchConnectString string, pchMetadata string) common.SteamAPICall
	iSteamParties_OnReservationCompleted         func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID common.Uint64SteamID)
	iSteamParties_CancelReservation              func(steamParties uintptr, ulBeacon PartyBeaconID, userSteamID common.Uint64SteamID)
	iSteamParties_ChangeNumOpenSlots             func(steamParties uintptr, ulBeacon PartyBeaconID, unOpenSlots uint32) common.SteamAPICall
	iSteamParties_DestroyBeacon                  func(steamParties uintptr, ulBeacon PartyBeaconID) bool
	iSteamParties_GetBeaconLocationData          func(steamParties uintptr, BeaconLocation uintptr, eData ESteamPartyBeaconLocationData, pchDataStringOut *[]byte, cchDataStringOut int32) bool
)

var steamParties uintptr

func parties() uintptr {
	if steamParties == 0 {
		steamParties = steamParties_init()
		return steamParties
	}
	return steamParties
}

func GetNumActiveBeacons() uint32 {
	return iSteamParties_GetNumActiveBeacons(parties())
}

func GetBeaconByIndex(Index uint32) PartyBeaconID {
	return iSteamParties_GetBeaconByIndex(parties(), Index)
}

func GetBeaconDetails(BeaconID PartyBeaconID, MetadataSize int32) (beaconOwnerSteamID common.CSteamID, Location SteamPartyBeaconLocation, metadata string, success bool) {
	var pchMetadata []byte = make([]byte, 0, MetadataSize)
	success = iSteamParties_GetBeaconDetails(parties(), BeaconID, &beaconOwnerSteamID, &Location, &pchMetadata, MetadataSize)
	return beaconOwnerSteamID, Location, string(pchMetadata), success
}

func JoinParty(BeaconID PartyBeaconID) common.CallResult[JoinPartyCallback] {
	handle := iSteamParties_JoinParty(parties(), BeaconID)
	return common.CallResult[JoinPartyCallback]{Handle: handle}
}

func GetNumAvailableBeaconLocations() (NumLocations uint32, success bool) {
	success = iSteamParties_GetNumAvailableBeaconLocations(parties(), &NumLocations)
	return NumLocations, success
}

func GetAvailableBeaconLocations(MaxNumLocations uint32) (locationList []SteamPartyBeaconLocation, success bool) {
	pLocationList := make([]SteamPartyBeaconLocation, 0, MaxNumLocations)
	success = iSteamParties_GetAvailableBeaconLocations(parties(), &pLocationList, MaxNumLocations)
	return pLocationList, success
}

func CreateBeacon(OpenSlots uint32, BeaconLocation SteamPartyBeaconLocation, ConnectString string, Metadata string) common.CallResult[CreateBeaconCallback] {
	handle := iSteamParties_CreateBeacon(parties(), OpenSlots, common.StructToUintptr[SteamPartyBeaconLocation](&BeaconLocation), ConnectString, Metadata)
	return common.CallResult[CreateBeaconCallback]{Handle: handle}
}

func OnReservationCompleted(Beacon PartyBeaconID, userSteamID common.Uint64SteamID) {
	iSteamParties_OnReservationCompleted(parties(), Beacon, userSteamID)
}

func CancelReservation(Beacon PartyBeaconID, userSteamID common.Uint64SteamID) {
	iSteamParties_CancelReservation(parties(), Beacon, userSteamID)
}

func ChangeNumOpenSlots(Beacon PartyBeaconID, OpenSlots uint32) common.CallResult[ChangeNumOpenSlotsCallback] {
	handle := iSteamParties_ChangeNumOpenSlots(parties(), Beacon, OpenSlots)
	return common.CallResult[ChangeNumOpenSlotsCallback]{Handle: handle}
}

func DestroyBeacon(beacon PartyBeaconID) bool {
	return iSteamParties_DestroyBeacon(parties(), beacon)
}

func GetBeaconLocationData(BeaconLocation SteamPartyBeaconLocation, Data ESteamPartyBeaconLocationData, DataStringOut int32) (data string, success bool) {
	pchDataStringOut := make([]byte, 0, DataStringOut)
	success = iSteamParties_GetBeaconLocationData(parties(), common.StructToUintptr[SteamPartyBeaconLocation](&BeaconLocation), Data, &pchDataStringOut, DataStringOut)
	return string(pchDataStringOut), success
}
