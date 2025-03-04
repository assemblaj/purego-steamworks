package networkingmessages

import (
	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

func init() {
	purego.RegisterLibFunc(&steamNetworkingMessages_init, SteamAPIDLL, flatAPI_SteamNetworkingMessages)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_SendMessageToUser, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_SendMessageToUser)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_ReceiveMessagesOnChannel, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_ReceiveMessagesOnChannel)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_AcceptSessionWithUser, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_AcceptSessionWithUser)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_CloseSessionWithUser, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_CloseSessionWithUser)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_CloseChannelWithUser, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_CloseChannelWithUser)
	purego.RegisterLibFunc(&iSteamNetworkingMessages_GetSessionConnectionInfo, SteamAPIDLL, flatAPI_ISteamNetworkingMessages_GetSessionConnectionInfo)
}
