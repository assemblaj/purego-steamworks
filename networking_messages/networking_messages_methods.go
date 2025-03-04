package networkingmessages

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	common "github.com/assemblaj/purego-steamworks"
)

const (
	flatAPI_SteamNetworkingMessages                           = "SteamAPI_SteamNetworkingMessages_SteamAPI_v002"
	flatAPI_ISteamNetworkingMessages_SendMessageToUser        = "SteamAPI_ISteamNetworkingMessages_SendMessageToUser"
	flatAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel = "SteamAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel"
	flatAPI_ISteamNetworkingMessages_AcceptSessionWithUser    = "SteamAPI_ISteamNetworkingMessages_AcceptSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseSessionWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseSessionWithUser"
	flatAPI_ISteamNetworkingMessages_CloseChannelWithUser     = "SteamAPI_ISteamNetworkingMessages_CloseChannelWithUser"
	flatAPI_ISteamNetworkingMessages_GetSessionConnectionInfo = "SteamAPI_ISteamNetworkingMessages_GetSessionConnectionInfo"
)

var (
	steamNetworkingMessages_init                      func() uintptr
	iSteamNetworkingMessages_SendMessageToUser        func(steamNetworkingMessages uintptr, identityRemote uintptr, pubData []byte, cubData uint32, nSendFlags int32, nRemoteChannel int32) common.EResult
	iSteamNetworkingMessages_ReceiveMessagesOnChannel func(steamNetworkingMessages uintptr, nLocalChannel int32, ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingMessages_AcceptSessionWithUser    func(steamNetworkingMessages uintptr, identityRemote uintptr) bool
	iSteamNetworkingMessages_CloseSessionWithUser     func(steamNetworkingMessages uintptr, identityRemote uintptr) bool
	iSteamNetworkingMessages_CloseChannelWithUser     func(steamNetworkingMessages uintptr, identityRemote uintptr, nLocalChannel int32) bool
	iSteamNetworkingMessages_GetSessionConnectionInfo func(steamNetworkingMessages uintptr, identityRemote uintptr, pConnectionInfo *SteamNetConnectionInfo, pQuickStatus *SteamNetConnectionRealTimeStatus) ESteamNetworkingConnectionState
)

var steamNetworkingMessages uintptr

func networkingMessages() uintptr {
	if steamNetworkingMessages == 0 {
		steamNetworkingMessages = steamNetworkingMessages_init()
		return steamNetworkingMessages
	}
	return steamNetworkingMessages
}

func SendMessageToUser(remoteIdentity SteamNetworkingIdentity, Data []byte, SendFlags int32, RemoteChannel int32) common.EResult {
	return iSteamNetworkingMessages_SendMessageToUser(networkingMessages(), common.StructToUintptr[SteamNetworkingIdentity](&remoteIdentity), Data, uint32(len(Data)), SendFlags, RemoteChannel)
}

func ReceiveMessagesOnChannel(LocalChannel int32, MaxMessages int32) []SteamNetworkingMessage {
	var ppOutMessages []SteamNetworkingMessage = make([]SteamNetworkingMessage, 0, MaxMessages)
	actual := iSteamNetworkingMessages_ReceiveMessagesOnChannel(networkingMessages(), LocalChannel, &ppOutMessages, MaxMessages)
	return ppOutMessages[:actual]
}

func AcceptSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool {
	return iSteamNetworkingMessages_AcceptSessionWithUser(networkingMessages(), common.StructToUintptr[SteamNetworkingIdentity](&remoteIdentity))
}

func CloseSessionWithUser(remoteIdentity SteamNetworkingIdentity) bool {
	return iSteamNetworkingMessages_CloseSessionWithUser(networkingMessages(), common.StructToUintptr[SteamNetworkingIdentity](&remoteIdentity))
}

func CloseChannelWithUser(remoteIdentity SteamNetworkingIdentity, LocalChannel int32) bool {
	return iSteamNetworkingMessages_CloseChannelWithUser(networkingMessages(), common.StructToUintptr[SteamNetworkingIdentity](&remoteIdentity), LocalChannel)
}

func GetSessionConnectionInfo(remoteIdentity SteamNetworkingIdentity) (ConnectionInfo SteamNetConnectionInfo, QuickStatus SteamNetConnectionRealTimeStatus, state ESteamNetworkingConnectionState) {
	state = iSteamNetworkingMessages_GetSessionConnectionInfo(networkingMessages(), common.StructToUintptr[SteamNetworkingIdentity](&remoteIdentity), &ConnectionInfo, &QuickStatus)
	return ConnectionInfo, QuickStatus, state
}
