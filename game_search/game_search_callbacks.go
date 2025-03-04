package gamesearch

import . "github.com/assemblaj/purego-steamworks"

const (
	SearchForGameProgressCallbackID            SteamCallbackID = 5201
	SearchForGameResultCallbackID              SteamCallbackID = 5202
	RequestPlayersForGameProgressCallbackID    SteamCallbackID = 5211
	RequestPlayersForGameResultCallbackID      SteamCallbackID = 5212
	RequestPlayersForGameFinalResultCallbackID SteamCallbackID = 5213
	SubmitPlayerResultResultCallbackID         SteamCallbackID = 5214
	EndGameResultCallbackID                    SteamCallbackID = 5215
)

type SearchForGameProgressCallback struct {
	SearchID                 uint64
	Result                   EResult
	LobbyID                  CSteamID
	SteamIDEndedSearch       CSteamID
	SecondsRemainingEstimate int32
	PlayersSearching         int32
}
type SearchForGameResultCallback struct {
	SearchID           uint64
	Result             EResult
	CountPlayersInGame int32
	CountAcceptedGame  int32
	SteamIDHost        CSteamID
	FinalCallback      bool
	_                  [7]byte
}

type RequestPlayersForGameProgressCallback struct {
	Result   EResult
	SearchID uint64
}
type RequestPlayersForGameResultCallback struct {
	Result                   EResult
	_                        [4]byte
	SearchID                 uint64
	SteamIDPlayerFound       CSteamID
	SteamIDLobby             CSteamID
	PlayerAcceptState        PlayerAcceptState
	PlayerIndex              int32
	TotalPlayersFound        int32
	TotalPlayersAcceptedGame int32
	SuggestedTeamIndex       int32
	_                        [4]byte
	UniqueGameID             uint64
}
type PlayerAcceptState int32

const (
	EStateUnknown        PlayerAcceptState = 0
	EStatePlayerAccepted PlayerAcceptState = 1
	EStatePlayerDeclined PlayerAcceptState = 2
)

type RequestPlayersForGameFinalResultCallback struct {
	Result       EResult
	SearchID     uint64
	UniqueGameID uint64
}
type SubmitPlayerResultResultCallback struct {
	Result        EResult
	UniqueGameID  uint64
	SteamIDPlayer CSteamID
}
type EndGameResultCallback struct {
	Result       EResult
	UniqueGameID uint64
}

func (cb SearchForGameProgressCallback) CallbackID() SteamCallbackID {
	return SearchForGameProgressCallbackID
}

func (cb SearchForGameProgressCallback) Name() string {
	return "SearchForGameProgressCallback"
}

func (cb SearchForGameProgressCallback) String() string {
	return CallbackString(cb)
}

func (cb SearchForGameResultCallback) CallbackID() SteamCallbackID {
	return SearchForGameResultCallbackID
}

func (cb SearchForGameResultCallback) Name() string {
	return "SearchForGameResultCallback"
}

func (cb SearchForGameResultCallback) String() string {
	return CallbackString(cb)
}

func (cb RequestPlayersForGameProgressCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameProgressCallbackID
}

func (cb RequestPlayersForGameProgressCallback) Name() string {
	return "RequestPlayersForGameProgressCallback"
}

func (cb RequestPlayersForGameProgressCallback) String() string {
	return CallbackString(cb)
}

func (cb RequestPlayersForGameResultCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameResultCallbackID
}

func (cb RequestPlayersForGameResultCallback) Name() string {
	return "RequestPlayersForGameResultCallback"
}

func (cb RequestPlayersForGameResultCallback) String() string {
	return CallbackString(cb)
}

func (cb RequestPlayersForGameFinalResultCallback) CallbackID() SteamCallbackID {
	return RequestPlayersForGameFinalResultCallbackID
}

func (cb RequestPlayersForGameFinalResultCallback) Name() string {
	return "RequestPlayersForGameFinalResultCallback"
}

func (cb RequestPlayersForGameFinalResultCallback) String() string {
	return CallbackString(cb)
}

func (cb SubmitPlayerResultResultCallback) CallbackID() SteamCallbackID {
	return SubmitPlayerResultResultCallbackID
}

func (cb SubmitPlayerResultResultCallback) Name() string {
	return "SubmitPlayerResultResultCallback"
}

func (cb SubmitPlayerResultResultCallback) String() string {
	return CallbackString(cb)
}

func (cb EndGameResultCallback) CallbackID() SteamCallbackID {
	return EndGameResultCallbackID
}

func (cb EndGameResultCallback) Name() string {
	return "EndGameResultCallback"
}

func (cb EndGameResultCallback) String() string {
	return CallbackString(cb)
}
