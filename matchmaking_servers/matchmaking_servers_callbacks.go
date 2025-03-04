package matchmakingservers

import (
	. "github.com/assemblaj/purego-steamworks/matchmaking"
)

type MatchmakingRulesResponse struct{}

func (s MatchmakingRulesResponse) RulesResponded(pchRule string, pchValue string) {
	iSteamMatchmakingRulesResponse_RulesResponded(pchRule, pchValue)
}

func (s MatchmakingRulesResponse) RulesFailedToRespond() {
	iSteamMatchmakingRulesResponse_RulesFailedToRespond()
}

func (s MatchmakingRulesResponse) RulesRefreshComplete() {
	iSteamMatchmakingRulesResponse_RulesRefreshComplete()
}

type ISteamMatchmakingPingResponse struct{}

func (s ISteamMatchmakingPingResponse) ServerResponded(server *GameServerItem) {
	iSteamMatchmakingPingResponse_ServerResponded(server)
}

func (s ISteamMatchmakingPingResponse) ServerFailedToRespond() {
	iSteamMatchmakingPingResponse_ServerFailedToRespond()
}

type MatchmakingPlayersResponse struct{}

func (s MatchmakingPlayersResponse) AddPlayerToList(pchName string, nScore int32, flTimePlayed float32) {
	iSteamMatchmakingPlayersResponse_AddPlayerToList(pchName, nScore, flTimePlayed)
}

func (s MatchmakingPlayersResponse) PlayersFailedToRespond() {
	iSteamMatchmakingPlayersResponse_PlayersFailedToRespond()
}

func (s MatchmakingPlayersResponse) PlayersRefreshComplete() {
	iSteamMatchmakingPlayersResponse_PlayersRefreshComplete()
}

type MatchmakingServerListResponse struct{}

func (s MatchmakingServerListResponse) ServerResponded(hRequest HServerListRequest, iServer int32) {
	iSteamMatchmakingServerListResponse_ServerResponded(hRequest, iServer)
}

func (s MatchmakingServerListResponse) ServerFailedToRespond(hRequest HServerListRequest, iServer int32) {
	iSteamMatchmakingServerListResponse_ServerFailedToRespond(hRequest, iServer)
}

func (s MatchmakingServerListResponse) RefreshComplete(hRequest HServerListRequest, response EMatchMakingServerResponse) {
	iSteamMatchmakingServerListResponse_RefreshComplete(hRequest, response)
}
