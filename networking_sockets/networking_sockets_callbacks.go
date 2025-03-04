package networkingsockets

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	. "github.com/assemblaj/purego-steamworks"
)

const SteamNetAuthenticationStatusID SteamCallbackID = 1222

type SteamNetAuthenticationStatus struct {
	Avail    ESteamNetworkingAvailability
	DebugMsg [256]byte
}

const SteamNetConnectionStatusChangedCallbackID SteamCallbackID = 1221

type SteamNetConnectionStatusChangedCallback struct {
	Conn     HSteamNetConnection
	Info     SteamNetConnectionInfo
	OldState ESteamNetworkingConnectionState
}

const SteamNetworkingFakeIPResultID SteamCallbackID = 1223

type SteamNetworkingFakeIPResult struct {
	Result   EResult
	Identity SteamNetworkingIdentity
	IP       uint32
	Ports    [8]uint16
	_        [4]byte
}

func (cb SteamNetAuthenticationStatus) CallbackID() SteamCallbackID {
	return SteamNetAuthenticationStatusID
}

func (cb SteamNetAuthenticationStatus) Name() string {
	return "SteamNetAuthenticationStatus"
}

func (cb SteamNetAuthenticationStatus) String() string {
	return CallbackString(cb)
}

func (cb SteamNetConnectionStatusChangedCallback) CallbackID() SteamCallbackID {
	return SteamNetConnectionStatusChangedCallbackID
}

func (cb SteamNetConnectionStatusChangedCallback) Name() string {
	return "SteamNetConnectionStatusChangedCallback"
}

func (cb SteamNetConnectionStatusChangedCallback) String() string {
	return CallbackString(cb)
}

func (cb SteamNetworkingFakeIPResult) CallbackID() SteamCallbackID {
	return SteamNetworkingFakeIPResultID
}

func (cb SteamNetworkingFakeIPResult) Name() string {
	return "SteamNetworkingFakeIPResult"
}

func (cb SteamNetworkingFakeIPResult) String() string {
	return CallbackString(cb)
}
