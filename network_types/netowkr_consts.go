package networktypes

const (
	HSteamNetConnection_Invalid                            HSteamNetConnection  = 0
	HSteamListenSocket_Invalid                             HSteamListenSocket   = 0
	HSteamNetPollGroup_Invalid                             HSteamNetPollGroup   = 0
	MaxSteamNetworkingErrMsg                               int32                = 1024
	SteamNetworkingMaxConnectionCloseReason                int32                = 128
	SteamNetworkingMaxConnectionDescription                int32                = 128
	SteamNetworkingMaxConnectionAppName                    int32                = 32
	SteamNetworkConnectionInfoFlags_Unauthenticated        int32                = 1
	SteamNetworkConnectionInfoFlags_Unencrypted            int32                = 2
	SteamNetworkConnectionInfoFlags_LoopbackBuffers        int32                = 4
	SteamNetworkConnectionInfoFlags_Fast                   int32                = 8
	SteamNetworkConnectionInfoFlags_Relayed                int32                = 16
	SteamNetworkConnectionInfoFlags_DualWifi               int32                = 32
	MaxSteamNetworkingSocketsMessageSizeSend               int32                = 512 * 1024
	SteamNetworkingSend_Unreliable                         int32                = 0
	SteamNetworkingSend_NoNagle                            int32                = 1
	SteamNetworkingSend_UnreliableNoNagle                  int32                = SteamNetworkingSend_Unreliable | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_NoDelay                            int32                = 4
	SteamNetworkingSend_UnreliableNoDelay                  int32                = SteamNetworkingSend_Unreliable | SteamNetworkingSend_NoDelay | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_Reliable                           int32                = 8
	SteamNetworkingSend_ReliableNoNagle                    int32                = SteamNetworkingSend_Reliable | SteamNetworkingSend_NoNagle
	SteamNetworkingSend_UseCurrentThread                   int32                = 16
	SteamNetworkingSend_AutoRestartBrokenSession           int32                = 32
	MaxSteamNetworkingPingLocationString                   int32                = 1024
	SteamNetworkingPing_Failed                             int32                = -1
	SteamNetworkingPing_Unknown                            int32                = -2
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Default int32                = -1
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Disable int32                = 0
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Relay   int32                = 1
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Private int32                = 2
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_Public  int32                = 4
	SteamNetworkingConfig_P2P_Transport_ICE_Enable_All     int32                = 0x7fffffff
	SteamDatagramPOPID_dev                                 SteamNetworkingPOPID = (SteamNetworkingPOPID('d') << 16) | (SteamNetworkingPOPID('e') << 8) | SteamNetworkingPOPID('v')
	SteamDatagramMaxSerializedTicket                       uint32               = 512
	MaxSteamDatagramGameCoordinatorServerLoginAppData      uint32               = 2048
	MaxSteamDatagramGameCoordinatorServerLoginSerialized   uint32               = 4096
	SteamNetworkingSocketsFakeUDPPortRecommendedMTU        int32                = 1200
	SteamNetworkingSocketsFakeUDPPortMaxMessageSize        int32                = 4096
)
