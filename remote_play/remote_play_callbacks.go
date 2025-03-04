package remoteplay

import . "github.com/assemblaj/purego-steamworks"

const (
	SteamRemotePlaySessionConnectedID    SteamCallbackID = 5701
	SteamRemotePlaySessionDisconnectedID SteamCallbackID = 5702
	SteamRemotePlayTogetherGuestInviteID SteamCallbackID = 5703
)

type SteamRemotePlaySessionConnected struct {
	SessionID RemotePlaySessionID
}
type SteamRemotePlaySessionDisconnected struct {
	SessionID RemotePlaySessionID
}
type SteamRemotePlayTogetherGuestInvite struct {
	ConnectURL [1024]byte
}

func (cb SteamRemotePlaySessionConnected) CallbackID() SteamCallbackID {
	return SteamRemotePlaySessionConnectedID
}

func (cb SteamRemotePlaySessionConnected) Name() string {
	return "SteamRemotePlaySessionConnected"
}

func (cb SteamRemotePlaySessionConnected) String() string {
	return CallbackString(cb)
}

func (cb SteamRemotePlaySessionDisconnected) CallbackID() SteamCallbackID {
	return SteamRemotePlaySessionDisconnectedID
}

func (cb SteamRemotePlaySessionDisconnected) Name() string {
	return "SteamRemotePlaySessionDisconnected"
}

func (cb SteamRemotePlaySessionDisconnected) String() string {
	return CallbackString(cb)
}

func (cb SteamRemotePlayTogetherGuestInvite) CallbackID() SteamCallbackID {
	return SteamRemotePlayTogetherGuestInviteID
}

func (cb SteamRemotePlayTogetherGuestInvite) Name() string {
	return "SteamRemotePlayTogetherGuestInvite"
}

func (cb SteamRemotePlayTogetherGuestInvite) String() string {
	return CallbackString(cb)
}
