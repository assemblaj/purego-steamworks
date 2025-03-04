package steamgameserverstats

import . "github.com/assemblaj/purego-steamworks"

const (
	GSStatsReceivedID SteamCallbackID = 1800
	GSStatsStoredID   SteamCallbackID = 1801
	GSStatsUnloadedID SteamCallbackID = 1108
)

type GSStatsReceived struct {
	Result      EResult
	SteamIDUser CSteamID
}
type GSStatsStored struct {
	Result      EResult
	SteamIDUser CSteamID
}
type GSStatsUnloaded struct {
	SteamIDUser CSteamID
}

func (cb GSStatsReceived) CallbackID() SteamCallbackID {
	return GSStatsReceivedID
}

func (cb GSStatsReceived) Name() string {
	return "GSStatsReceived"
}

func (cb GSStatsReceived) String() string {
	return CallbackString(cb)
}

func (cb GSStatsStored) CallbackID() SteamCallbackID {
	return GSStatsStoredID
}

func (cb GSStatsStored) Name() string {
	return "GSStatsStored"
}

func (cb GSStatsStored) String() string {
	return CallbackString(cb)
}

func (cb GSStatsUnloaded) CallbackID() SteamCallbackID {
	return GSStatsUnloadedID
}

func (cb GSStatsUnloaded) Name() string {
	return "GSStatsUnloaded"
}

func (cb GSStatsUnloaded) String() string {
	return CallbackString(cb)
}
