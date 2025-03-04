package parties

import . "github.com/assemblaj/purego-steamworks"

const (
	JoinPartyCallbackID               SteamCallbackID = 5301
	CreateBeaconCallbackID            SteamCallbackID = 5302
	ReservationNotificationCallbackID SteamCallbackID = 5303
	ChangeNumOpenSlotsCallbackID      SteamCallbackID = 5304
	AvailableBeaconLocationsUpdatedID SteamCallbackID = 5305
	ActiveBeaconsUpdatedID            SteamCallbackID = 5306
)

type JoinPartyCallback struct {
	Result             EResult
	BeaconID           PartyBeaconID
	SteamIDBeaconOwner CSteamID
	ConnectString      [256]byte
}
type CreateBeaconCallback struct {
	Result   EResult
	BeaconID PartyBeaconID
}
type ReservationNotificationCallback struct {
	BeaconID      PartyBeaconID
	SteamIDJoiner CSteamID
}
type ChangeNumOpenSlotsCallback struct {
	Result EResult
}
type AvailableBeaconLocationsUpdated struct {
}
type ActiveBeaconsUpdated struct {
}

func (cb JoinPartyCallback) CallbackID() SteamCallbackID {
	return JoinPartyCallbackID
}

func (cb JoinPartyCallback) Name() string {
	return "JoinPartyCallback"
}

func (cb JoinPartyCallback) String() string {
	return CallbackString(cb)
}

func (cb CreateBeaconCallback) CallbackID() SteamCallbackID {
	return CreateBeaconCallbackID
}

func (cb CreateBeaconCallback) Name() string {
	return "CreateBeaconCallback"
}

func (cb CreateBeaconCallback) String() string {
	return CallbackString(cb)
}

func (cb ReservationNotificationCallback) CallbackID() SteamCallbackID {
	return ReservationNotificationCallbackID
}

func (cb ReservationNotificationCallback) Name() string {
	return "ReservationNotificationCallback"
}

func (cb ReservationNotificationCallback) String() string {
	return CallbackString(cb)
}

func (cb ChangeNumOpenSlotsCallback) CallbackID() SteamCallbackID {
	return ChangeNumOpenSlotsCallbackID
}

func (cb ChangeNumOpenSlotsCallback) Name() string {
	return "ChangeNumOpenSlotsCallback"
}

func (cb ChangeNumOpenSlotsCallback) String() string {
	return CallbackString(cb)
}

func (cb AvailableBeaconLocationsUpdated) CallbackID() SteamCallbackID {
	return AvailableBeaconLocationsUpdatedID
}

func (cb AvailableBeaconLocationsUpdated) Name() string {
	return "AvailableBeaconLocationsUpdated"
}

func (cb AvailableBeaconLocationsUpdated) String() string {
	return CallbackString(cb)
}

func (cb ActiveBeaconsUpdated) CallbackID() SteamCallbackID {
	return ActiveBeaconsUpdatedID
}

func (cb ActiveBeaconsUpdated) Name() string {
	return "ActiveBeaconsUpdated"
}

func (cb ActiveBeaconsUpdated) String() string {
	return CallbackString(cb)
}
