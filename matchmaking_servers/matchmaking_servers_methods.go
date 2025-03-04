package matchmakingservers

import (
	. "github.com/assemblaj/purego-steamworks/matchmaking"

	. "github.com/assemblaj/purego-steamworks"
)

const (
	flatAPI_MatchMakingKeyValuePair_t_Construct = "SteamAPI_MatchMakingKeyValuePair_t_Construct"
)

var (
	matchMakingKeyValuePair_t_Construct func(*MatchMakingKeyValuePair)
)

type MatchMakingKeyValuePair struct {
	Key   [256]byte
	Value [256]byte
}

func (m *MatchMakingKeyValuePair) Construct() {
	matchMakingKeyValuePair_t_Construct(m)
}

type ServerNetAdr struct {
	ConnectionPort uint16
	QueryPort      uint16
	IP             uint32
}

const (
	flatAPI_servernetadr_t_Construct                  = "SteamAPI_servernetadr_t_Construct"
	flatAPI_servernetadr_t_Init                       = "SteamAPI_servernetadr_t_Initt"
	flatAPI_servernetadr_t_GetQueryPort               = "SteamAPI_servernetadr_t_GetQueryPort"
	flatAPI_servernetadr_t_SetQueryPort               = "SteamAPI_servernetadr_t_SetQueryPort"
	flatAPI_servernetadr_t_GetConnectionPort          = "SteamAPI_servernetadr_t_GetConnectionPort"
	flatAPI_servernetadr_t_SetConnectionPort          = "SteamAPI_servernetadr_t_SetConnectionPort"
	flatAPI_servernetadr_t_GetIP                      = "SteamAPI_servernetadr_t_GetIP"
	flatAPI_servernetadr_t_SetIP                      = "SteamAPI_servernetadr_t_SetIP"
	flatAPI_servernetadr_t_GetConnectionAddressString = "SteamAPI_servernetadr_t_GetConnectionAddressString"
	flatAPI_servernetadr_t_GetQueryAddressString      = "SteamAPI_servernetadr_t_GetQueryAddressString"
	flatAPI_servernetadr_t_IsLessThan                 = "SteamAPI_servernetadr_t_IsLessThan"
	flatAPI_servernetadr_t_Assign                     = "SteamAPI_servernetadr_t_Assign"
)

var (
	servernetadr_t_Construct                  func(self *ServerNetAdr)
	servernetadr_t_Init                       func(self *ServerNetAdr, ip uint, QueryPort uint16, ConnectionPort uint16)
	servernetadr_t_GetQueryPort               func(self *ServerNetAdr) uint16
	servernetadr_t_SetQueryPort               func(self *ServerNetAdr, Port uint16)
	servernetadr_t_GetConnectionPort          func(self *ServerNetAdr) uint16
	servernetadr_t_SetConnectionPort          func(self *ServerNetAdr, Port uint16)
	servernetadr_t_GetIP                      func(self *ServerNetAdr) uint32
	servernetadr_t_SetIP                      func(self *ServerNetAdr, IP uint32)
	servernetadr_t_GetConnectionAddressString func(self *ServerNetAdr) string
	servernetadr_t_GetQueryAddressString      func(self *ServerNetAdr) string
	servernetadr_t_IsLessThan                 func(self *ServerNetAdr, netadr *ServerNetAdr) bool
	servernetadr_t_Assign                     func(self *ServerNetAdr, that *ServerNetAdr)
)

func (s *ServerNetAdr) Construct() {
	servernetadr_t_Construct(s)
}

func (s *ServerNetAdr) Init(ip uint, QueryPort uint16, ConnectionPort uint16) {
	servernetadr_t_Init(s, ip, QueryPort, ConnectionPort)
}
func (s *ServerNetAdr) GetQueryPort() uint16 {
	return servernetadr_t_GetQueryPort(s)
}
func (s *ServerNetAdr) SetQueryPort(Port uint16) {
	servernetadr_t_SetQueryPort(s, Port)
}
func (s *ServerNetAdr) GetConnectionPort() uint16 {
	return servernetadr_t_GetConnectionPort(s)
}
func (s *ServerNetAdr) SetConnectionPort(Port uint16) {
	servernetadr_t_SetConnectionPort(s, Port)
}
func (s *ServerNetAdr) GetIP() uint32 {
	return servernetadr_t_GetIP(s)
}
func (s *ServerNetAdr) SetIP(IP uint32) {
	servernetadr_t_SetIP(s, IP)
}
func (s *ServerNetAdr) GetConnectionAddressString() string {
	return servernetadr_t_GetConnectionAddressString(s)
}
func (s *ServerNetAdr) GetQueryAddressString() string {
	return servernetadr_t_GetQueryAddressString(s)
}
func (s *ServerNetAdr) IsLessThan(netadr *ServerNetAdr) bool {
	return servernetadr_t_IsLessThan(s, netadr)
}
func (s *ServerNetAdr) Assign(that *ServerNetAdr) {
	servernetadr_t_Assign(s, that)
}

type GameServerItem struct {
	NetAdr                ServerNetAdr
	Ping                  int32
	HadSuccessfulResponse bool
	DoNotRefresh          bool
	_                     [2]byte
	GameDir               [32]byte
	Map                   [32]byte
	GameDescription       [64]byte
	AppID                 uint32
	Players               int32
	MaxPlayers            int32
	BotPlayers            int32
	Password              bool
	Secure                bool
	_                     [2]byte
	TimeLastPlayed        uint32
	ServerVersion         int32
	ServerName            [64]byte
	GameTags              [128]byte
	ID                    CSteamID
}

func (g *GameServerItem) Construct() {
	gameserveritem_t_Construct(g)
}

func (g *GameServerItem) GetName() string {
	return gameserveritem_t_GetName(g)
}

func (g *GameServerItem) SetName(name string) {
	gameserveritem_t_SetName(g, name)
}

const (
	flatAPI_gameserveritem_t_Construct = "SteamAPI_gameserveritem_t_Construct"
	flatAPI_gameserveritem_t_GetName   = "SteamAPI_gameserveritem_t_GetName"
	flatAPI_gameserveritem_t_SetName   = "SteamAPI_gameserveritem_t_SetName"
)

var (
	gameserveritem_t_Construct func(self *GameServerItem)
	gameserveritem_t_GetName   func(self *GameServerItem) string
	gameserveritem_t_SetName   func(self *GameServerItem, Name string)
)

type HServerListRequest uintptr
type HServerQuery int32

const (
	flatAPI_SteamMatchmakingServers                                   = "SteamAPI_SteamMatchmakingServers_v002"
	flatAPI_ISteamMatchmakingServers_RequestInternetServerList        = "SteamAPI_ISteamMatchmakingServers_RequestInternetServerList"
	flatAPI_ISteamMatchmakingServers_RequestLANServerList             = "SteamAPI_ISteamMatchmakingServers_RequestLANServerList"
	flatAPI_ISteamMatchmakingServers_RequestFriendsServerList         = "SteamAPI_ISteamMatchmakingServers_RequestFriendsServerList"
	flatAPI_ISteamMatchmakingServers_RequestFavoritesServerList       = "SteamAPI_ISteamMatchmakingServers_RequestFavoritesServerList"
	flatAPI_ISteamMatchmakingServers_RequestHistoryServerList         = "SteamAPI_ISteamMatchmakingServers_RequestHistoryServerList"
	flatAPI_ISteamMatchmakingServers_RequestSpectatorServerList       = "SteamAPI_ISteamMatchmakingServers_RequestSpectatorServerList"
	flatAPI_ISteamMatchmakingServers_ReleaseRequest                   = "SteamAPI_ISteamMatchmakingServers_ReleaseRequest"
	flatAPI_ISteamMatchmakingServers_GetServerDetails                 = "SteamAPI_ISteamMatchmakingServers_GetServerDetails"
	flatAPI_ISteamMatchmakingServers_CancelQuery                      = "SteamAPI_ISteamMatchmakingServers_CancelQuery"
	flatAPI_ISteamMatchmakingServers_RefreshQuery                     = "SteamAPI_ISteamMatchmakingServers_RefreshQuery"
	flatAPI_ISteamMatchmakingServers_IsRefreshing                     = "SteamAPI_ISteamMatchmakingServers_IsRefreshing"
	flatAPI_ISteamMatchmakingServers_GetServerCount                   = "SteamAPI_ISteamMatchmakingServers_GetServerCount"
	flatAPI_ISteamMatchmakingServers_RefreshServer                    = "SteamAPI_ISteamMatchmakingServers_RefreshServer"
	flatAPI_ISteamMatchmakingServers_PingServer                       = "SteamAPI_ISteamMatchmakingServers_PingServer"
	flatAPI_ISteamMatchmakingServers_PlayerDetails                    = "SteamAPI_ISteamMatchmakingServers_PlayerDetails"
	flatAPI_ISteamMatchmakingServers_ServerRules                      = "SteamAPI_ISteamMatchmakingServers_ServerRules"
	flatAPI_ISteamMatchmakingServers_CancelServerQuery                = "SteamAPI_ISteamMatchmakingServers_CancelServerQuery"
	flatAPI_ISteamMatchmakingServerListResponse_ServerResponded       = "SteamAPI_ISteamMatchmakingServerListResponse_ServerResponded"
	flatAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond = "SteamAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond"
	flatAPI_ISteamMatchmakingServerListResponse_RefreshComplete       = "SteamAPI_ISteamMatchmakingServerListResponse_RefreshComplete"
	flatAPI_ISteamMatchmakingPingResponse_ServerResponded             = "SteamAPI_ISteamMatchmakingPingResponse_ServerResponded"
	flatAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond       = "SteamAPI_ISteamMatchmakingPingResponse_ServerFailedToRespond"
	flatAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList          = "SteamAPI_ISteamMatchmakingPlayersResponse_AddPlayerToList"
	flatAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond   = "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersFailedToRespond"
	flatAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete   = "SteamAPI_ISteamMatchmakingPlayersResponse_PlayersRefreshComplete"
	flatAPI_ISteamMatchmakingRulesResponse_RulesResponded             = "SteamAPI_ISteamMatchmakingRulesResponse_RulesResponded"
	flatAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond       = "SteamAPI_ISteamMatchmakingRulesResponse_RulesFailedToRespond"
	flatAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete       = "SteamAPI_ISteamMatchmakingRulesResponse_RulesRefreshComplete"
)

var (
	steamMatchmakingServers_init                              func() uintptr
	iSteamMatchmakingServers_RequestInternetServerList        func(steamMatchmakingServers uintptr, iApp AppId, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestLANServerList             func(steamMatchmakingServers uintptr, iApp AppId, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestFriendsServerList         func(steamMatchmakingServers uintptr, iApp AppId, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestFavoritesServerList       func(steamMatchmakingServers uintptr, iApp AppId, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestHistoryServerList         func(steamMatchmakingServers uintptr, iApp AppId, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_RequestSpectatorServerList       func(steamMatchmakingServers uintptr, iApp AppId, ppchFilters *[]MatchMakingKeyValuePair, nFilters uint32, pRequestServersResponse *MatchmakingServerListResponse) HServerListRequest
	iSteamMatchmakingServers_ReleaseRequest                   func(steamMatchmakingServers uintptr, hServerListRequest HServerListRequest)
	iSteamMatchmakingServers_GetServerDetails                 func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32) *GameServerItem
	iSteamMatchmakingServers_CancelQuery                      func(steamMatchmakingServers uintptr, hRequest HServerListRequest)
	iSteamMatchmakingServers_RefreshQuery                     func(steamMatchmakingServers uintptr, hRequest HServerListRequest)
	iSteamMatchmakingServers_IsRefreshing                     func(steamMatchmakingServers uintptr, hRequest HServerListRequest) bool
	iSteamMatchmakingServers_GetServerCount                   func(steamMatchmakingServers uintptr, hRequest HServerListRequest) int32
	iSteamMatchmakingServers_RefreshServer                    func(steamMatchmakingServers uintptr, hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServers_PingServer                       func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *ISteamMatchmakingPingResponse) HServerQuery
	iSteamMatchmakingServers_PlayerDetails                    func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingPlayersResponse) HServerQuery
	iSteamMatchmakingServers_ServerRules                      func(steamMatchmakingServers uintptr, unIP uint32, usPort uint16, pRequestServersResponse *MatchmakingRulesResponse) HServerQuery
	iSteamMatchmakingServers_CancelServerQuery                func(steamMatchmakingServers uintptr, hServerQuery HServerQuery)
	iSteamMatchmakingServerListResponse_ServerResponded       func(hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServerListResponse_ServerFailedToRespond func(hRequest HServerListRequest, iServer int32)
	iSteamMatchmakingServerListResponse_RefreshComplete       func(hRequest HServerListRequest, response EMatchMakingServerResponse)
	iSteamMatchmakingPingResponse_ServerResponded             func(server *GameServerItem)
	iSteamMatchmakingPingResponse_ServerFailedToRespond       func()
	iSteamMatchmakingPlayersResponse_AddPlayerToList          func(pchName string, nScore int32, flTimePlayed float32)
	iSteamMatchmakingPlayersResponse_PlayersFailedToRespond   func()
	iSteamMatchmakingPlayersResponse_PlayersRefreshComplete   func()
	iSteamMatchmakingRulesResponse_RulesResponded             func(pchRule string, pchValue string)
	iSteamMatchmakingRulesResponse_RulesFailedToRespond       func()
	iSteamMatchmakingRulesResponse_RulesRefreshComplete       func()
)

var steamMatchmakingServers uintptr

func matchmakingServers() uintptr {
	if steamMatchmakingServers == 0 {
		steamMatchmakingServers = steamMatchmakingServers_init()
		return steamMatchmakingServers
	}
	return steamMatchmakingServers
}

func RequestInternetServerList(App AppId, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestInternetServerList(matchmakingServers(), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func RequestLANServerList(App AppId) (RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	request = iSteamMatchmakingServers_RequestLANServerList(matchmakingServers(), App, &RequestServersResponse)
	return RequestServersResponse, request
}

func RequestFriendsServerList(App AppId, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestFriendsServerList(matchmakingServers(), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func RequestFavoritesServerList(App AppId, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestFavoritesServerList(matchmakingServers(), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func RequestHistoryServerList(App AppId, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestHistoryServerList(matchmakingServers(), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func RequestSpectatorServerList(App AppId, FilterCount uint32) (Filters []MatchMakingKeyValuePair, RequestServersResponse MatchmakingServerListResponse, request HServerListRequest) {
	Filters = make([]MatchMakingKeyValuePair, 0, FilterCount)
	request = iSteamMatchmakingServers_RequestSpectatorServerList(matchmakingServers(), App, &Filters, FilterCount, &RequestServersResponse)
	return Filters, RequestServersResponse, request
}

func ReleaseRequest(ServerListRequest HServerListRequest) {
	iSteamMatchmakingServers_ReleaseRequest(matchmakingServers(), ServerListRequest)
}

func GetServerDetails(Request HServerListRequest, Server int32) *GameServerItem {
	return iSteamMatchmakingServers_GetServerDetails(matchmakingServers(), Request, Server)
}

func CancelQuery(Request HServerListRequest) {
	iSteamMatchmakingServers_CancelQuery(matchmakingServers(), Request)
}

func RefreshQuery(Request HServerListRequest) {
	iSteamMatchmakingServers_RefreshQuery(matchmakingServers(), Request)
}

func IsRefreshing(Request HServerListRequest) bool {
	return iSteamMatchmakingServers_IsRefreshing(matchmakingServers(), Request)
}

func GetServerCount(Request HServerListRequest) int32 {
	return iSteamMatchmakingServers_GetServerCount(matchmakingServers(), Request)
}

func RefreshServer(Request HServerListRequest, Server int32) {
	iSteamMatchmakingServers_RefreshServer(matchmakingServers(), Request, Server)
}

func PingServer(IP uint32, Port uint16) (RequestServersResponse ISteamMatchmakingPingResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_PingServer(matchmakingServers(), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func PlayerDetails(IP uint32, Port uint16) (RequestServersResponse MatchmakingPlayersResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_PlayerDetails(matchmakingServers(), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func ServerRules(IP uint32, Port uint16) (RequestServersResponse MatchmakingRulesResponse, query HServerQuery) {
	query = iSteamMatchmakingServers_ServerRules(matchmakingServers(), IP, Port, &RequestServersResponse)
	return RequestServersResponse, query
}

func CancelServerQuery(serverQuery HServerQuery) {
	iSteamMatchmakingServers_CancelServerQuery(matchmakingServers(), serverQuery)
}
