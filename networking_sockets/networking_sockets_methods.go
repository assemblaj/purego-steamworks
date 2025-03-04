package networkingsockets

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	common "github.com/assemblaj/purego-steamworks"
)

var (
	steamNetworkingSockets_init                                     func() uintptr
	iSteamNetworkingSockets_CreateListenSocketIP                    func(steamNetworkingSockets uintptr, localAddress uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_ConnectByIPAddress                      func(steamNetworkingSockets uintptr, address uintptr, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_CreateListenSocketP2P                   func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_ConnectP2P                              func(steamNetworkingSockets uintptr, identityRemote uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_AcceptConnection                        func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) common.EResult
	iSteamNetworkingSockets_CloseConnection                         func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nReason int32, pszDebug string, bEnableLinger bool) bool
	iSteamNetworkingSockets_CloseListenSocket                       func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket) bool
	iSteamNetworkingSockets_SetConnectionUserData                   func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, nUserData int64) bool
	iSteamNetworkingSockets_GetConnectionUserData                   func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection) int64
	iSteamNetworkingSockets_SetConnectionName                       func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName string)
	iSteamNetworkingSockets_GetConnectionName                       func(steamNetworkingSockets uintptr, hPeer HSteamNetConnection, pszName *[]byte, nMaxLen int32) bool
	iSteamNetworkingSockets_SendMessageToConnection                 func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pData []byte, cbData uint32, nSendFlags int32, pOutMessageNumber *int64) common.EResult
	iSteamNetworkingSockets_SendMessages                            func(steamNetworkingSockets uintptr, nMessages int32, pMessages []SteamNetworkingMessage, pOutMessageNumberOrResult *[]int64)
	iSteamNetworkingSockets_FlushMessagesOnConnection               func(steamNetworkingSockets uintptr, hConn HSteamNetConnection) common.EResult
	iSteamNetworkingSockets_ReceiveMessagesOnConnection             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingSockets_GetConnectionInfo                       func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pInfo *SteamNetConnectionInfo) bool
	iSteamNetworkingSockets_GetConnectionRealTimeStatus             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pStatus *SteamNetConnectionRealTimeStatus, nLanes int32, pLanes *[]SteamNetConnectionRealTimeLaneStatus) common.EResult
	iSteamNetworkingSockets_GetDetailedConnectionStatus             func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pszBuf *[]byte, cbBuf int32) int32
	iSteamNetworkingSockets_GetListenSocketAddress                  func(steamNetworkingSockets uintptr, hSocket HSteamListenSocket, address *SteamNetworkingIPAddr) bool
	iSteamNetworkingSockets_CreateSocketPair                        func(steamNetworkingSockets uintptr, pOutConnection1 *HSteamNetConnection, pOutConnection2 *HSteamNetConnection, bUseNetworkLoopback bool, pIdentity1 *SteamNetworkingIdentity, pIdentity2 *SteamNetworkingIdentity) bool
	iSteamNetworkingSockets_ConfigureConnectionLanes                func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, nNumLanes int32, pLanePriorities []int32, pLaneWeights []uint16) common.EResult
	iSteamNetworkingSockets_GetIdentity                             func(steamNetworkingSockets uintptr, pIdentity *SteamNetworkingIdentity) bool
	iSteamNetworkingSockets_InitAuthentication                      func(steamNetworkingSockets uintptr) ESteamNetworkingAvailability
	iSteamNetworkingSockets_GetAuthenticationStatus                 func(steamNetworkingSockets uintptr, pDetails *SteamNetAuthenticationStatus) ESteamNetworkingAvailability
	iSteamNetworkingSockets_CreatePollGroup                         func(steamNetworkingSockets uintptr) HSteamNetPollGroup
	iSteamNetworkingSockets_DestroyPollGroup                        func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup) bool
	iSteamNetworkingSockets_SetConnectionPollGroup                  func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, hPollGroup HSteamNetPollGroup) bool
	iSteamNetworkingSockets_ReceiveMessagesOnPollGroup              func(steamNetworkingSockets uintptr, hPollGroup HSteamNetPollGroup, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingSockets_ReceivedRelayAuthTicket                 func(steamNetworkingSockets uintptr, pvTicket []byte, cbTicket int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) bool
	iSteamNetworkingSockets_FindRelayAuthTicketForServer            func(steamNetworkingSockets uintptr, identityGameServer uintptr, nRemoteVirtualPort int32, pOutParsedTicket *SteamDatagramRelayAuthTicket) int32
	iSteamNetworkingSockets_ConnectToHostedDedicatedServer          func(steamNetworkingSockets uintptr, identityTarget uintptr, nRemoteVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_GetHostedDedicatedServerPort            func(steamNetworkingSockets uintptr) uint16
	iSteamNetworkingSockets_GetHostedDedicatedServerPOPID           func(steamNetworkingSockets uintptr) SteamNetworkingPOPID
	iSteamNetworkingSockets_GetHostedDedicatedServerAddress         func(steamNetworkingSockets uintptr, pRouting *SteamDatagramHostedAddress) common.EResult
	iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket func(steamNetworkingSockets uintptr, nLocalVirtualPort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_GetGameCoordinatorServerLogin           func(steamNetworkingSockets uintptr, pLoginInfo *SteamDatagramGameCoordinatorServerLogin, pcbSignedBlob *int32, pBlob *[]byte) common.EResult
	iSteamNetworkingSockets_ConnectP2PCustomSignaling               func(steamNetworkingSockets uintptr, pSignaling *ISteamNetworkingConnectionSignaling, pPeerIdentity *SteamNetworkingIdentity, nRemoteVirtualPort int32, nOptions int32, pOptions *SteamNetworkingConfigValue) HSteamNetConnection
	iSteamNetworkingSockets_ReceivedP2PCustomSignal                 func(steamNetworkingSockets uintptr, pMsg []byte, cbMsg int32, pContext *ISteamNetworkingSignalingRecvContext) bool
	iSteamNetworkingSockets_GetCertificateRequest                   func(steamNetworkingSockets uintptr, pcbBlob *int32, pBlob *[]byte, errMsg *SteamNetworkingErrMsg) bool
	iSteamNetworkingSockets_SetCertificate                          func(steamNetworkingSockets uintptr, pCertificate []byte, cbCertificate int32, errMsg *SteamNetworkingErrMsg) bool
	iSteamNetworkingSockets_ResetIdentity                           func(steamNetworkingSockets uintptr, pIdentity uintptr)
	iSteamNetworkingSockets_RunCallbacks                            func(steamNetworkingSockets uintptr)
	iSteamNetworkingSockets_BeginAsyncRequestFakeIP                 func(steamNetworkingSockets uintptr, nNumPorts int32) bool
	iSteamNetworkingSockets_GetFakeIP                               func(steamNetworkingSockets uintptr, idxFirstPort int32, pInfo *SteamNetworkingFakeIPResult)
	iSteamNetworkingSockets_CreateListenSocketP2PFakeIP             func(steamNetworkingSockets uintptr, idxFakePort int32, nOptions int32, pOptions []SteamNetworkingConfigValue) HSteamListenSocket
	iSteamNetworkingSockets_GetRemoteFakeIPForConnection            func(steamNetworkingSockets uintptr, hConn HSteamNetConnection, pOutAddr *SteamNetworkingIPAddr) common.EResult
	iSteamNetworkingSockets_CreateFakeUDPPort                       func(steamNetworkingSockets uintptr, idxFakeServerPort int32) *SteamNetworkingFakeUDPPort
)

const (
	flatAPI_SteamNetworkingSockets                                          = "SteamAPI_SteamNetworkingSockets_SteamAPI_v012"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketIP                    = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketIP"
	flatAPI_ISteamNetworkingSockets_ConnectByIPAddress                      = "SteamAPI_ISteamNetworkingSockets_ConnectByIPAddress"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketP2P                   = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2P"
	flatAPI_ISteamNetworkingSockets_ConnectP2P                              = "SteamAPI_ISteamNetworkingSockets_ConnectP2P"
	flatAPI_ISteamNetworkingSockets_AcceptConnection                        = "SteamAPI_ISteamNetworkingSockets_AcceptConnection"
	flatAPI_ISteamNetworkingSockets_CloseConnection                         = "SteamAPI_ISteamNetworkingSockets_CloseConnection"
	flatAPI_ISteamNetworkingSockets_CloseListenSocket                       = "SteamAPI_ISteamNetworkingSockets_CloseListenSocket"
	flatAPI_ISteamNetworkingSockets_SetConnectionUserData                   = "SteamAPI_ISteamNetworkingSockets_SetConnectionUserData"
	flatAPI_ISteamNetworkingSockets_GetConnectionUserData                   = "SteamAPI_ISteamNetworkingSockets_GetConnectionUserData"
	flatAPI_ISteamNetworkingSockets_SetConnectionName                       = "SteamAPI_ISteamNetworkingSockets_SetConnectionName"
	flatAPI_ISteamNetworkingSockets_GetConnectionName                       = "SteamAPI_ISteamNetworkingSockets_GetConnectionName"
	flatAPI_ISteamNetworkingSockets_SendMessageToConnection                 = "SteamAPI_ISteamNetworkingSockets_SendMessageToConnection"
	flatAPI_ISteamNetworkingSockets_SendMessages                            = "SteamAPI_ISteamNetworkingSockets_SendMessages"
	flatAPI_ISteamNetworkingSockets_FlushMessagesOnConnection               = "SteamAPI_ISteamNetworkingSockets_FlushMessagesOnConnection"
	flatAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection             = "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnConnection"
	flatAPI_ISteamNetworkingSockets_GetConnectionInfo                       = "SteamAPI_ISteamNetworkingSockets_GetConnectionInfo"
	flatAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus             = "SteamAPI_ISteamNetworkingSockets_GetConnectionRealTimeStatus"
	flatAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus             = "SteamAPI_ISteamNetworkingSockets_GetDetailedConnectionStatus"
	flatAPI_ISteamNetworkingSockets_GetListenSocketAddress                  = "SteamAPI_ISteamNetworkingSockets_GetListenSocketAddress"
	flatAPI_ISteamNetworkingSockets_CreateSocketPair                        = "SteamAPI_ISteamNetworkingSockets_CreateSocketPair"
	flatAPI_ISteamNetworkingSockets_ConfigureConnectionLanes                = "SteamAPI_ISteamNetworkingSockets_ConfigureConnectionLanes"
	flatAPI_ISteamNetworkingSockets_GetIdentity                             = "SteamAPI_ISteamNetworkingSockets_GetIdentity"
	flatAPI_ISteamNetworkingSockets_InitAuthentication                      = "SteamAPI_ISteamNetworkingSockets_InitAuthentication"
	flatAPI_ISteamNetworkingSockets_GetAuthenticationStatus                 = "SteamAPI_ISteamNetworkingSockets_GetAuthenticationStatus"
	flatAPI_ISteamNetworkingSockets_CreatePollGroup                         = "SteamAPI_ISteamNetworkingSockets_CreatePollGroup"
	flatAPI_ISteamNetworkingSockets_DestroyPollGroup                        = "SteamAPI_ISteamNetworkingSockets_DestroyPollGroup"
	flatAPI_ISteamNetworkingSockets_SetConnectionPollGroup                  = "SteamAPI_ISteamNetworkingSockets_SetConnectionPollGroup"
	flatAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup              = "SteamAPI_ISteamNetworkingSockets_ReceiveMessagesOnPollGroup"
	flatAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket                 = "SteamAPI_ISteamNetworkingSockets_ReceivedRelayAuthTicket"
	flatAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer            = "SteamAPI_ISteamNetworkingSockets_FindRelayAuthTicketForServer"
	flatAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer          = "SteamAPI_ISteamNetworkingSockets_ConnectToHostedDedicatedServer"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort            = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPort"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID           = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerPOPID"
	flatAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress         = "SteamAPI_ISteamNetworkingSockets_GetHostedDedicatedServerAddress"
	flatAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket = "SteamAPI_ISteamNetworkingSockets_CreateHostedDedicatedServerListenSocket"
	flatAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin           = "SteamAPI_ISteamNetworkingSockets_GetGameCoordinatorServerLogin"
	flatAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling               = "SteamAPI_ISteamNetworkingSockets_ConnectP2PCustomSignaling"
	flatAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal                 = "SteamAPI_ISteamNetworkingSockets_ReceivedP2PCustomSignal"
	flatAPI_ISteamNetworkingSockets_GetCertificateRequest                   = "SteamAPI_ISteamNetworkingSockets_GetCertificateRequest"
	flatAPI_ISteamNetworkingSockets_SetCertificate                          = "SteamAPI_ISteamNetworkingSockets_SetCertificate"
	flatAPI_ISteamNetworkingSockets_ResetIdentity                           = "SteamAPI_ISteamNetworkingSockets_ResetIdentity"
	flatAPI_ISteamNetworkingSockets_RunCallbacks                            = "SteamAPI_ISteamNetworkingSockets_RunCallbacks"
	flatAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP                 = "SteamAPI_ISteamNetworkingSockets_BeginAsyncRequestFakeIP"
	flatAPI_ISteamNetworkingSockets_GetFakeIP                               = "SteamAPI_ISteamNetworkingSockets_GetFakeIP"
	flatAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP             = "SteamAPI_ISteamNetworkingSockets_CreateListenSocketP2PFakeIP"
	flatAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection            = "SteamAPI_ISteamNetworkingSockets_GetRemoteFakeIPForConnection"
	flatAPI_ISteamNetworkingSockets_CreateFakeUDPPort                       = "SteamAPI_ISteamNetworkingSockets_CreateFakeUDPPort"
)

var steamNetworkingSockets uintptr

func networkingSockets() uintptr {
	if steamNetworkingSockets == 0 {
		steamNetworkingSockets = steamNetworkingSockets_init()
		return steamNetworkingSockets
	}
	return steamNetworkingSockets
}

func CreateListenSocketIP(localAddress SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketIP(networkingSockets(), common.StructToUintptr[SteamNetworkingIPAddr](&localAddress), int32(len(Options)), Options)
}

func ConnectByIPAddress(address SteamNetworkingIPAddr, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectByIPAddress(networkingSockets(), common.StructToUintptr[SteamNetworkingIPAddr](&address), int32(len(Options)), Options)
}

func CreateListenSocketP2P(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketP2P(networkingSockets(), LocalVirtualPort, int32(len(Options)), Options)
}

func ConnectP2P(identityRemote SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectP2P(networkingSockets(), common.StructToUintptr[SteamNetworkingIdentity](&identityRemote), RemoteVirtualPort, int32(len(Options)), Options)
}

func AcceptConnection(Conn HSteamNetConnection) common.EResult {
	return iSteamNetworkingSockets_AcceptConnection(networkingSockets(), Conn)
}

func CloseConnection(Peer HSteamNetConnection, Reason int32, Debug string, EnableLinger bool) bool {
	return iSteamNetworkingSockets_CloseConnection(networkingSockets(), Peer, Reason, Debug, EnableLinger)
}

func CloseListenSocket(Socket HSteamListenSocket) bool {
	return iSteamNetworkingSockets_CloseListenSocket(networkingSockets(), Socket)
}

func SetConnectionUserData(Peer HSteamNetConnection, UserData int64) bool {
	return iSteamNetworkingSockets_SetConnectionUserData(networkingSockets(), Peer, UserData)
}

func GetConnectionUserData(Peer HSteamNetConnection) int64 {
	return iSteamNetworkingSockets_GetConnectionUserData(networkingSockets(), Peer)
}

func SetConnectionName(Peer HSteamNetConnection, Name string) {
	iSteamNetworkingSockets_SetConnectionName(networkingSockets(), Peer, Name)
}

func GetConnectionName(Peer HSteamNetConnection, MaxLen int32) (name string, success bool) {
	var pszName []byte = make([]byte, 0, MaxLen)
	success = iSteamNetworkingSockets_GetConnectionName(networkingSockets(), Peer, &pszName, MaxLen)
	return string(pszName), success
}

func SendMessageToConnection(Conn HSteamNetConnection, Data []byte, SendFlags int32) (OutMessageNumber int64, result common.EResult) {
	result = iSteamNetworkingSockets_SendMessageToConnection(networkingSockets(), Conn, Data, uint32(len(Data)), SendFlags, &OutMessageNumber)
	return OutMessageNumber, result
}

// To use this
// / function, you must first allocate a message object using
// / ISteamNetworkingUtils::AllocateMessage.  (Do not declare one
// / on the stack or allocate your own.)
func SendMessages(numMessages int32, Messages []SteamNetworkingMessage) (OutMessageNumberOrResult []int64) {
	OutMessageNumberOrResult = make([]int64, 0, len(Messages))
	iSteamNetworkingSockets_SendMessages(networkingSockets(), numMessages, Messages, &OutMessageNumberOrResult)
	return OutMessageNumberOrResult
}

func FlushMessagesOnConnection(Conn HSteamNetConnection) common.EResult {
	return iSteamNetworkingSockets_FlushMessagesOnConnection(networkingSockets(), Conn)
}

func ReceiveMessagesOnConnection(Conn HSteamNetConnection, MaxMessages int32) []SteamNetworkingMessage {
	var ppOutMessages []SteamNetworkingMessage = make([]SteamNetworkingMessage, 0, MaxMessages)
	result := iSteamNetworkingSockets_ReceiveMessagesOnConnection(networkingSockets(), Conn, &ppOutMessages, MaxMessages)
	return ppOutMessages[:result]
}

func GetConnectionInfo(Conn HSteamNetConnection) (Info SteamNetConnectionInfo, success bool) {
	success = iSteamNetworkingSockets_GetConnectionInfo(networkingSockets(), Conn, &Info)
	return Info, success
}

func GetConnectionRealTimeStatus(Conn HSteamNetConnection, numLanes int32) (Status SteamNetConnectionRealTimeStatus, Lanes []SteamNetConnectionRealTimeLaneStatus, result common.EResult) {
	Lanes = make([]SteamNetConnectionRealTimeLaneStatus, 0, numLanes)
	result = iSteamNetworkingSockets_GetConnectionRealTimeStatus(networkingSockets(), Conn, &Status, numLanes, &Lanes)
	return Status, Lanes, result
}

func GetDetailedConnectionStatus(Conn HSteamNetConnection, bufSize int32) (status string, result int32) {
	var pszBuf []byte = make([]byte, 0, bufSize)
	result = iSteamNetworkingSockets_GetDetailedConnectionStatus(networkingSockets(), Conn, &pszBuf, bufSize)
	return string(pszBuf), result
}

func GetListenSocketAddress(Socket HSteamListenSocket) (address SteamNetworkingIPAddr, success bool) {
	success = iSteamNetworkingSockets_GetListenSocketAddress(networkingSockets(), Socket, &address)
	return address, success
}

func CreateSocketPair(UseNetworkLoopback bool) (OutConnection1 HSteamNetConnection, OutConnection2 HSteamNetConnection, Identity1 SteamNetworkingIdentity, Identity2 SteamNetworkingIdentity, succes bool) {
	succes = iSteamNetworkingSockets_CreateSocketPair(networkingSockets(), &OutConnection1, &OutConnection2, UseNetworkLoopback, &Identity1, &Identity2)
	return OutConnection1, OutConnection2, Identity1, Identity2, succes
}

func ConfigureConnectionLanes(Conn HSteamNetConnection, NumLanes int32, LanePriorities []int32, LaneWeights []uint16) common.EResult {
	return iSteamNetworkingSockets_ConfigureConnectionLanes(networkingSockets(), Conn, NumLanes, LanePriorities, LaneWeights)
}

func GetIdentity() (Identity SteamNetworkingIdentity, success bool) {
	success = iSteamNetworkingSockets_GetIdentity(networkingSockets(), &Identity)
	return Identity, success
}

func InitAuthentication() ESteamNetworkingAvailability {
	return iSteamNetworkingSockets_InitAuthentication(networkingSockets())
}

func GetAuthenticationStatus() (Details SteamNetAuthenticationStatus, availability ESteamNetworkingAvailability) {
	availability = iSteamNetworkingSockets_GetAuthenticationStatus(networkingSockets(), &Details)
	return Details, availability
}

func CreatePollGroup() HSteamNetPollGroup {
	return iSteamNetworkingSockets_CreatePollGroup(networkingSockets())
}

func DestroyPollGroup(PollGroup HSteamNetPollGroup) bool {
	return iSteamNetworkingSockets_DestroyPollGroup(networkingSockets(), PollGroup)
}

func SetConnectionPollGroup(Conn HSteamNetConnection, PollGroup HSteamNetPollGroup) bool {
	return iSteamNetworkingSockets_SetConnectionPollGroup(networkingSockets(), Conn, PollGroup)
}

func ReceiveMessagesOnPollGroup(PollGroup HSteamNetPollGroup, MaxMessages int32) []SteamNetworkingMessage {
	var ppOutMessages []SteamNetworkingMessage = make([]SteamNetworkingMessage, 0, MaxMessages)
	result := iSteamNetworkingSockets_ReceiveMessagesOnPollGroup(networkingSockets(), PollGroup, &ppOutMessages, MaxMessages)
	return ppOutMessages[:result]
}

func ReceivedRelayAuthTicket(pvTicket []byte) (pOutParsedTicket SteamDatagramRelayAuthTicket, success bool) {
	success = iSteamNetworkingSockets_ReceivedRelayAuthTicket(networkingSockets(), pvTicket, int32(len(pvTicket)), &pOutParsedTicket)
	return pOutParsedTicket, success
}

func FindRelayAuthTicketForServer(gameServerIdentity SteamNetworkingIdentity, RemoteVirtualPort int32) (OutParsedTicket SteamDatagramRelayAuthTicket, secondsToExpire int32) {
	secondsToExpire = iSteamNetworkingSockets_FindRelayAuthTicketForServer(networkingSockets(), common.StructToUintptr[SteamNetworkingIdentity](&gameServerIdentity), RemoteVirtualPort, &OutParsedTicket)
	return OutParsedTicket, secondsToExpire
}

func ConnectToHostedDedicatedServer(targetIdentity SteamNetworkingIdentity, RemoteVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamNetConnection {
	return iSteamNetworkingSockets_ConnectToHostedDedicatedServer(networkingSockets(), common.StructToUintptr[SteamNetworkingIdentity](&targetIdentity), RemoteVirtualPort, int32(len(Options)), Options)
}

func GetHostedDedicatedServerPort() uint16 {
	return iSteamNetworkingSockets_GetHostedDedicatedServerPort(networkingSockets())
}

func GetHostedDedicatedServerPOPID() SteamNetworkingPOPID {
	return iSteamNetworkingSockets_GetHostedDedicatedServerPOPID(networkingSockets())
}

func GetHostedDedicatedServerAddress() (Routing SteamDatagramHostedAddress, result common.EResult) {
	result = iSteamNetworkingSockets_GetHostedDedicatedServerAddress(networkingSockets(), &Routing)
	return Routing, result
}

func CreateHostedDedicatedServerListenSocket(LocalVirtualPort int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateHostedDedicatedServerListenSocket(networkingSockets(), LocalVirtualPort, int32(len(Options)), Options)
}

func GetGameCoordinatorServerLogin(LoginInfo *SteamDatagramGameCoordinatorServerLogin, SignedBlob int32) (Blob []byte, result common.EResult) {
	Blob = make([]byte, 0, SignedBlob)
	result = iSteamNetworkingSockets_GetGameCoordinatorServerLogin(networkingSockets(), LoginInfo, &SignedBlob, &Blob)
	return Blob[:SignedBlob], result
}

// func  ConnectP2PCustomSignaling(pSignaling *ISteamNetworkingConnectionSignaling, pPeerIdentity *SteamNetworkingIdentity, nRemoteVirtualPort int32, nOptions int32, pOptions *SteamNetworkingConfigValue) HSteamNetConnection {
// 	return iSteamNetworkingSockets_ConnectP2PCustomSignaling(networkingSockets(), pSignaling, pPeerIdentity, nRemoteVirtualPort, nOptions, pOptions)
// }

// func  ReceivedP2PCustomSignal(cbMsg int32, pContext *ISteamNetworkingSignalingRecvContext) (pMsg []byte, success bool) {
// 	pMsg = make([]byte, 0, cbMsg)
// 	success = iSteamNetworkingSockets_ReceivedP2PCustomSignal(networkingSockets(), pMsg, cbMsg, pContext)
// 	return pMsg, success
// }

func GetCertificateRequest(BlobSize int32) (Blob []byte, errMsg SteamNetworkingErrMsg, success bool) {
	Blob = make([]byte, 0, BlobSize)
	success = iSteamNetworkingSockets_GetCertificateRequest(networkingSockets(), &BlobSize, &Blob, &errMsg)
	return Blob[:BlobSize], errMsg, success
}

func SetCertificate(Certificate []byte) (errMsg SteamNetworkingErrMsg, success bool) {
	success = iSteamNetworkingSockets_SetCertificate(networkingSockets(), Certificate, int32(len(Certificate)), &errMsg)
	return errMsg, success
}

/// NOTE: This function is not actually supported on Steam!  It is included
///       for use on other platforms where the active user can sign out and
///       a new user can sign in.

func ResetIdentity(Identity SteamNetworkingIdentity) {
	iSteamNetworkingSockets_ResetIdentity(networkingSockets(), common.StructToUintptr[SteamNetworkingIdentity](&Identity))
}

func RunCallbacks() {
	iSteamNetworkingSockets_RunCallbacks(networkingSockets())
}

func BeginAsyncRequestFakeIP(NumPorts int32) bool {
	return iSteamNetworkingSockets_BeginAsyncRequestFakeIP(networkingSockets(), NumPorts)
}

func GetFakeIP(FirstPortIdx int32) (Info SteamNetworkingFakeIPResult) {
	iSteamNetworkingSockets_GetFakeIP(networkingSockets(), FirstPortIdx, &Info)
	return Info
}

func CreateListenSocketP2PFakeIP(FakePortIdx int32, Options []SteamNetworkingConfigValue) HSteamListenSocket {
	return iSteamNetworkingSockets_CreateListenSocketP2PFakeIP(networkingSockets(), FakePortIdx, int32(len(Options)), Options)
}

func GetRemoteFakeIPForConnection(Conn HSteamNetConnection) (OutAddr SteamNetworkingIPAddr, result common.EResult) {
	result = iSteamNetworkingSockets_GetRemoteFakeIPForConnection(networkingSockets(), Conn, &OutAddr)
	return OutAddr, result
}

func CreateFakeUDPPort(FakeServerPortIdx int32) *SteamNetworkingFakeUDPPort {
	return iSteamNetworkingSockets_CreateFakeUDPPort(networkingSockets(), FakeServerPortIdx)
}
