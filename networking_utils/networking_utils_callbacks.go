package networkingutils

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	. "github.com/assemblaj/purego-steamworks"
)

type SteamRelayNetworkStatus struct {
	Avail                     ESteamNetworkingAvailability
	PingMeasurementInProgress int32
	AvailNetworkConfig        ESteamNetworkingAvailability
	AvailAnyRelay             ESteamNetworkingAvailability
	DebugMsg                  [256]byte
}

const SteamRelayNetworkStatusID SteamCallbackID = 1281

func (cb SteamRelayNetworkStatus) CallbackID() SteamCallbackID {
	return SteamRelayNetworkStatusID
}

func (cb SteamRelayNetworkStatus) Name() string {
	return "SteamRelayNetworkStatus"
}

func (cb SteamRelayNetworkStatus) String() string {
	return CallbackString(cb)
}
