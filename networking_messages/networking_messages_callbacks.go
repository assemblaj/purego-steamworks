package networkingmessages

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	. "github.com/assemblaj/purego-steamworks"
)

const SteamNetworkingMessagesSessionRequestID SteamCallbackID = 1251

type SteamNetworkingMessagesSessionRequest struct {
	IdentityRemote SteamNetworkingIdentity
}

const SteamNetworkingMessagesSessionFailedID SteamCallbackID = 1252

type SteamNetworkingMessagesSessionFailed struct {
	Info SteamNetConnectionInfo
}

func (cb SteamNetworkingMessagesSessionRequest) CallbackID() SteamCallbackID {
	return SteamNetworkingMessagesSessionRequestID
}

func (cb SteamNetworkingMessagesSessionRequest) Name() string {
	return "SteamNetworkingMessagesSessionRequest"
}

func (cb SteamNetworkingMessagesSessionRequest) String() string {
	return CallbackString(cb)
}

func (cb SteamNetworkingMessagesSessionFailed) CallbackID() SteamCallbackID {
	return SteamNetworkingMessagesSessionFailedID
}

func (cb SteamNetworkingMessagesSessionFailed) Name() string {
	return "SteamNetworkingMessagesSessionFailed"
}

func (cb SteamNetworkingMessagesSessionFailed) String() string {
	return CallbackString(cb)
}
