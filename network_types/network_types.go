package networktypes

import (
	"unsafe"

	. "github.com/assemblaj/purego-steamworks"
)

type ESteamIPType int32

const (
	ESteamIPTypeIPv4 ESteamIPType = 0
	ESteamIPTypeIPv6 ESteamIPType = 1
)

type SteamIPAddress struct {
	IPv6 [16]uint8
	Type ESteamIPType
	_    [2]byte
}

type HSteamNetConnection uint
type HSteamListenSocket uint
type HSteamNetPollGroup uint
type SteamDatagramRelayAuthTicket struct{}
type ISteamNetworkingConnectionSignaling struct{}
type ISteamNetworkingSignalingRecvContext struct{}
type SteamNetworkingFakeUDPPort struct{}

type SteamNetworkingErrMsg [1024]byte
type SteamDatagramHostedAddress struct {
	Size int32
	Data [128]byte
}
type SteamDatagramGameCoordinatorServerLogin struct {
	Identity SteamNetworkingIdentity
	Routing  SteamDatagramHostedAddress
	AppID    AppId
	RTime    RTime32
	AppData  int32
	Data     [2048]byte
}

type ESteamNetworkingConnectionState int32

const (
	ESteamNetworkingConnectionStateNone                   ESteamNetworkingConnectionState = 0
	ESteamNetworkingConnectionStateConnecting             ESteamNetworkingConnectionState = 1
	ESteamNetworkingConnectionStateFindingRoute           ESteamNetworkingConnectionState = 2
	ESteamNetworkingConnectionStateConnected              ESteamNetworkingConnectionState = 3
	ESteamNetworkingConnectionStateClosedByPeer           ESteamNetworkingConnectionState = 4
	ESteamNetworkingConnectionStateProblemDetectedLocally ESteamNetworkingConnectionState = 5
	ESteamNetworkingConnectionStateFinWait                ESteamNetworkingConnectionState = -1
	ESteamNetworkingConnectionStateLinger                 ESteamNetworkingConnectionState = -2
	ESteamNetworkingConnectionStateDead                   ESteamNetworkingConnectionState = -3
	ESteamNetworkingConnectionStateForce32Bit             ESteamNetworkingConnectionState = 2147483647
)

type SteamNetConnectionInfo struct {
	IdentityRemote        SteamNetworkingIdentity
	UserData              int64
	ListenSocket          HSteamListenSocket
	AddrRemote            SteamNetworkingIPAddr
	pad1                  uint16
	_                     [2]byte
	POPRemoteID           SteamNetworkingPOPID
	POPRelayID            SteamNetworkingPOPID
	State                 ESteamNetworkingConnectionState
	EndReason             int32
	EndDebug              [128]byte
	ConnectionDescription [128]byte
	Flags                 int32
	reserved              [63]uint32
}
type SteamNetConnectionRealTimeStatus struct {
	State                   ESteamNetworkingConnectionState
	Ping                    int32
	ConnectionQualityLocal  float32
	ConnectionQualityRemote float32
	OutPacketsPerSec        float32
	OutBytesPerSec          float32
	InPacketsPerSec         float32
	InBytesPerSec           float32
	SendRateBytesPerSecond  int32
	PendingUnreliable       int32
	PendingReliable         int32
	SentUnackedReliable     int32
	QueueTime               SteamNetworkingMicroseconds
	reserved                [16]uint32
}
type SteamNetConnectionRealTimeLaneStatus struct {
	PendingUnreliable   int32
	PendingReliable     int32
	SentUnackedReliable int32
	reservePad1         int32
	QueueTime           SteamNetworkingMicroseconds
	reserved            [10]uint32
}

type SteamNetworkingMessage struct {
	Data          unsafe.Pointer
	Size          int32
	Conn          HSteamNetConnection
	IdentityPeer  SteamNetworkingIdentity
	ConnUserData  int64
	TimeReceived  SteamNetworkingMicroseconds
	MessageNumber int64
	FreeData      uintptr
	Release       uintptr
	Channel       int32
	Flags         int32
	UserData      int64
	Lane          uint16
	pad1          uint16
	_             [4]byte
}

type FnSteamRelayNetworkStatusChanged uintptr

type ESteamNetworkingAvailability int32

const (
	ESteamNetworkingAvailabilityCannotTry  ESteamNetworkingAvailability = -102
	ESteamNetworkingAvailabilityFailed     ESteamNetworkingAvailability = -101
	ESteamNetworkingAvailabilityPreviously ESteamNetworkingAvailability = -100
	ESteamNetworkingAvailabilityRetrying   ESteamNetworkingAvailability = -10
	ESteamNetworkingAvailabilityNeverTried ESteamNetworkingAvailability = 1
	ESteamNetworkingAvailabilityWaiting    ESteamNetworkingAvailability = 2
	ESteamNetworkingAvailabilityAttempting ESteamNetworkingAvailability = 3
	ESteamNetworkingAvailabilityCurrent    ESteamNetworkingAvailability = 100
	ESteamNetworkingAvailabilityUnknown    ESteamNetworkingAvailability = 0
	ESteamNetworkingAvailabilityForce32bit ESteamNetworkingAvailability = 2147483647
)

type ESteamNetworkingSocketsDebugOutputType int32

const (
	ESteamNetworkingSocketsDebugOutputTypeNone       ESteamNetworkingSocketsDebugOutputType = 0
	ESteamNetworkingSocketsDebugOutputTypeBug        ESteamNetworkingSocketsDebugOutputType = 1
	ESteamNetworkingSocketsDebugOutputTypeError      ESteamNetworkingSocketsDebugOutputType = 2
	ESteamNetworkingSocketsDebugOutputTypeImportant  ESteamNetworkingSocketsDebugOutputType = 3
	ESteamNetworkingSocketsDebugOutputTypeWarning    ESteamNetworkingSocketsDebugOutputType = 4
	ESteamNetworkingSocketsDebugOutputTypeMsg        ESteamNetworkingSocketsDebugOutputType = 5
	ESteamNetworkingSocketsDebugOutputTypeVerbose    ESteamNetworkingSocketsDebugOutputType = 6
	ESteamNetworkingSocketsDebugOutputTypeDebug      ESteamNetworkingSocketsDebugOutputType = 7
	ESteamNetworkingSocketsDebugOutputTypeEverything ESteamNetworkingSocketsDebugOutputType = 8
	ESteamNetworkingSocketsDebugOutputTypeForce32Bit ESteamNetworkingSocketsDebugOutputType = 2147483647
)

type ESteamNetworkingFakeIPType int32

const (
	ESteamNetworkingFakeIPTypeInvalid    ESteamNetworkingFakeIPType = 0
	ESteamNetworkingFakeIPTypeNotFake    ESteamNetworkingFakeIPType = 1
	ESteamNetworkingFakeIPTypeGlobalIPv4 ESteamNetworkingFakeIPType = 2
	ESteamNetworkingFakeIPTypeLocalIPv4  ESteamNetworkingFakeIPType = 3
	ESteamNetworkingFakeIPTypeForce32Bit ESteamNetworkingFakeIPType = 2147483647
)

type SteamNetworkPingLocation struct {
	Data [512]uint8
}

type SteamNetworkingPOPID uint

type SteamNetworkingMicroseconds int64

type FSteamNetworkingSocketsDebugOutput uintptr

type SteamNetworkingIPAddr struct {
	IPv6 [16]uint8
	Port uint16
}
type ESteamNetworkingIdentityType int32

const (
	ESteamNetworkingIdentityTypeInvalid        ESteamNetworkingIdentityType = 0
	ESteamNetworkingIdentityTypeSteamID        ESteamNetworkingIdentityType = 16
	ESteamNetworkingIdentityTypeXboxPairwiseID ESteamNetworkingIdentityType = 17
	ESteamNetworkingIdentityTypeSonyPSN        ESteamNetworkingIdentityType = 18
	ESteamNetworkingIdentityTypeIPAddress      ESteamNetworkingIdentityType = 1
	ESteamNetworkingIdentityTypeGenericString  ESteamNetworkingIdentityType = 2
	ESteamNetworkingIdentityTypeGenericBytes   ESteamNetworkingIdentityType = 3
	ESteamNetworkingIdentityTypeUnknownType    ESteamNetworkingIdentityType = 4
	ESteamNetworkingIdentityTypeForce32bit     ESteamNetworkingIdentityType = 2147483647
)

type SteamNetworkingConfigValue struct {
	Value    ESteamNetworkingConfigValue
	DataType ESteamNetworkingConfigDataType
	Int64    int64
}
type SteamNetworkingIdentity struct {
	Type             ESteamNetworkingIdentityType
	Size             int32
	UnknownRawString [128]byte
}

type ESteamNetworkingConfigValue int32

const (
	ESteamNetworkingConfigInvalid                                       ESteamNetworkingConfigValue = 0
	ESteamNetworkingConfigTimeoutInitial                                ESteamNetworkingConfigValue = 24
	ESteamNetworkingConfigTimeoutConnected                              ESteamNetworkingConfigValue = 25
	ESteamNetworkingConfigSendBufferSize                                ESteamNetworkingConfigValue = 9
	ESteamNetworkingConfigRecvBufferSize                                ESteamNetworkingConfigValue = 47
	ESteamNetworkingConfigRecvBufferMessages                            ESteamNetworkingConfigValue = 48
	ESteamNetworkingConfigRecvMaxMessageSize                            ESteamNetworkingConfigValue = 49
	ESteamNetworkingConfigRecvMaxSegmentsPerPacket                      ESteamNetworkingConfigValue = 50
	ESteamNetworkingConfigConnectionUserData                            ESteamNetworkingConfigValue = 40
	ESteamNetworkingConfigSendRateMin                                   ESteamNetworkingConfigValue = 10
	ESteamNetworkingConfigSendRateMax                                   ESteamNetworkingConfigValue = 11
	ESteamNetworkingConfigNagleTime                                     ESteamNetworkingConfigValue = 12
	ESteamNetworkingConfigIPAllowWithoutAuth                            ESteamNetworkingConfigValue = 23
	ESteamNetworkingConfigIPLocalHostAllowWithoutAuth                   ESteamNetworkingConfigValue = 52
	ESteamNetworkingConfigMTUPacketSize                                 ESteamNetworkingConfigValue = 32
	ESteamNetworkingConfigMTUDataSize                                   ESteamNetworkingConfigValue = 33
	ESteamNetworkingConfigUnencrypted                                   ESteamNetworkingConfigValue = 34
	ESteamNetworkingConfigSymmetricConnect                              ESteamNetworkingConfigValue = 37
	ESteamNetworkingConfigLocalVirtualPort                              ESteamNetworkingConfigValue = 38
	ESteamNetworkingConfigDualWifiEnable                                ESteamNetworkingConfigValue = 39
	ESteamNetworkingConfigEnableDiagnosticsUI                           ESteamNetworkingConfigValue = 46
	ESteamNetworkingConfigSendTimeSincePreviousPacket                   ESteamNetworkingConfigValue = 59
	ESteamNetworkingConfigFakePacketLossSend                            ESteamNetworkingConfigValue = 2
	ESteamNetworkingConfigFakePacketLossRecv                            ESteamNetworkingConfigValue = 3
	ESteamNetworkingConfigFakePacketLagSend                             ESteamNetworkingConfigValue = 4
	ESteamNetworkingConfigFakePacketLagRecv                             ESteamNetworkingConfigValue = 5
	ESteamNetworkingConfigFakePacketJitterSendAvg                       ESteamNetworkingConfigValue = 53
	ESteamNetworkingConfigFakePacketJitterSendMax                       ESteamNetworkingConfigValue = 54
	ESteamNetworkingConfigFakePacketJitterSendPct                       ESteamNetworkingConfigValue = 55
	ESteamNetworkingConfigFakePacketJitterRecvAvg                       ESteamNetworkingConfigValue = 56
	ESteamNetworkingConfigFakePacketJitterRecvMax                       ESteamNetworkingConfigValue = 57
	ESteamNetworkingConfigFakePacketJitterRecvPct                       ESteamNetworkingConfigValue = 58
	ESteamNetworkingConfigFakePacketReorderSend                         ESteamNetworkingConfigValue = 6
	ESteamNetworkingConfigFakePacketReorderRecv                         ESteamNetworkingConfigValue = 7
	ESteamNetworkingConfigFakePacketReorderTime                         ESteamNetworkingConfigValue = 8
	ESteamNetworkingConfigFakePacketDupSend                             ESteamNetworkingConfigValue = 26
	ESteamNetworkingConfigFakePacketDupRecv                             ESteamNetworkingConfigValue = 27
	ESteamNetworkingConfigFakePacketDupTimeMax                          ESteamNetworkingConfigValue = 28
	ESteamNetworkingConfigPacketTraceMaxBytes                           ESteamNetworkingConfigValue = 41
	ESteamNetworkingConfigFakeRateLimitSendRate                         ESteamNetworkingConfigValue = 42
	ESteamNetworkingConfigFakeRateLimitSendBurst                        ESteamNetworkingConfigValue = 43
	ESteamNetworkingConfigFakeRateLimitRecvRate                         ESteamNetworkingConfigValue = 44
	ESteamNetworkingConfigFakeRateLimitRecvBurst                        ESteamNetworkingConfigValue = 45
	ESteamNetworkingConfigOutOfOrderCorrectionWindowMicroseconds        ESteamNetworkingConfigValue = 51
	ESteamNetworkingConfigCallbackConnectionStatusChanged               ESteamNetworkingConfigValue = 201
	ESteamNetworkingConfigCallbackAuthStatusChanged                     ESteamNetworkingConfigValue = 202
	ESteamNetworkingConfigCallbackRelayNetworkStatusChanged             ESteamNetworkingConfigValue = 203
	ESteamNetworkingConfigCallbackMessagesSessionRequest                ESteamNetworkingConfigValue = 204
	ESteamNetworkingConfigCallbackMessagesSessionFailed                 ESteamNetworkingConfigValue = 205
	ESteamNetworkingConfigCallbackCreateConnectionSignaling             ESteamNetworkingConfigValue = 206
	ESteamNetworkingConfigCallbackFakeIPResult                          ESteamNetworkingConfigValue = 207
	ESteamNetworkingConfigP2PSTUNServerList                             ESteamNetworkingConfigValue = 103
	ESteamNetworkingConfigP2PTransportICEEnable                         ESteamNetworkingConfigValue = 104
	ESteamNetworkingConfigP2PTransportICEPenalty                        ESteamNetworkingConfigValue = 105
	ESteamNetworkingConfigP2PTransportSDRPenalty                        ESteamNetworkingConfigValue = 106
	ESteamNetworkingConfigP2PTURNServerList                             ESteamNetworkingConfigValue = 107
	ESteamNetworkingConfigP2PTURNUserList                               ESteamNetworkingConfigValue = 108
	ESteamNetworkingConfigP2PTURNPassList                               ESteamNetworkingConfigValue = 109
	ESteamNetworkingConfigP2PTransportICEImplementation                 ESteamNetworkingConfigValue = 110
	ESteamNetworkingConfigSDRClientConsecutitivePingTimeoutsFailInitial ESteamNetworkingConfigValue = 19
	ESteamNetworkingConfigSDRClientConsecutitivePingTimeoutsFail        ESteamNetworkingConfigValue = 20
	ESteamNetworkingConfigSDRClientMinPingsBeforePingAccurate           ESteamNetworkingConfigValue = 21
	ESteamNetworkingConfigSDRClientSingleSocket                         ESteamNetworkingConfigValue = 22
	ESteamNetworkingConfigSDRClientForceRelayCluster                    ESteamNetworkingConfigValue = 29
	ESteamNetworkingConfigSDRClientDevTicket                            ESteamNetworkingConfigValue = 30
	ESteamNetworkingConfigSDRClientForceProxyAddr                       ESteamNetworkingConfigValue = 31
	ESteamNetworkingConfigSDRClientFakeClusterPing                      ESteamNetworkingConfigValue = 36
	ESteamNetworkingConfigSDRClientLimitPingProbesToNearestN            ESteamNetworkingConfigValue = 60
	ESteamNetworkingConfigLogLevelAckRTT                                ESteamNetworkingConfigValue = 13
	ESteamNetworkingConfigLogLevelPacketDecode                          ESteamNetworkingConfigValue = 14
	ESteamNetworkingConfigLogLevelMessage                               ESteamNetworkingConfigValue = 15
	ESteamNetworkingConfigLogLevelPacketGaps                            ESteamNetworkingConfigValue = 16
	ESteamNetworkingConfigLogLevelP2PRendezvous                         ESteamNetworkingConfigValue = 17
	ESteamNetworkingConfigLogLevelSDRRelayPings                         ESteamNetworkingConfigValue = 18
	ESteamNetworkingConfigECN                                           ESteamNetworkingConfigValue = 999
	ESteamNetworkingConfigDELETEDEnumerateDevVars                       ESteamNetworkingConfigValue = 35
	ESteamNetworkingConfigValueForce32Bit                               ESteamNetworkingConfigValue = 2147483647
)

type FnSteamNetConnectionStatusChanged uintptr
type FnSteamNetAuthenticationStatusChanged uintptr
type FnSteamNetworkingMessagesSessionRequest uintptr
type FnSteamNetworkingMessagesSessionFailed uintptr
type FnSteamNetworkingFakeIPResult uintptr

type ESteamNetworkingConfigScope int32

const (
	ESteamNetworkingConfigGlobal           ESteamNetworkingConfigScope = 1
	ESteamNetworkingConfigSocketsInterface ESteamNetworkingConfigScope = 2
	ESteamNetworkingConfigListenSocket     ESteamNetworkingConfigScope = 3
	ESteamNetworkingConfigConnection       ESteamNetworkingConfigScope = 4
	ESteamNetworkingConfigScopeForce32Bit  ESteamNetworkingConfigScope = 2147483647
)

type ESteamNetworkingConfigDataType int32

const (
	ESteamNetworkingConfigInt32              ESteamNetworkingConfigDataType = 1
	ESteamNetworkingConfigInt64              ESteamNetworkingConfigDataType = 2
	ESteamNetworkingConfigFloat              ESteamNetworkingConfigDataType = 3
	ESteamNetworkingConfigString             ESteamNetworkingConfigDataType = 4
	ESteamNetworkingConfigPtr                ESteamNetworkingConfigDataType = 5
	ESteamNetworkingConfigDataTypeForce32Bit ESteamNetworkingConfigDataType = 2147483647
)

type ESteamNetworkingGetConfigValueResult int32

const (
	ESteamNetworkingGetConfigValueBadValue         ESteamNetworkingGetConfigValueResult = -1
	ESteamNetworkingGetConfigValueBadScopeObj      ESteamNetworkingGetConfigValueResult = -2
	ESteamNetworkingGetConfigValueBufferTooSmall   ESteamNetworkingGetConfigValueResult = -3
	ESteamNetworkingGetConfigValueOK               ESteamNetworkingGetConfigValueResult = 1
	ESteamNetworkingGetConfigValueOKInherited      ESteamNetworkingGetConfigValueResult = 2
	ESteamNetworkingGetConfigValueResultForce32Bit ESteamNetworkingGetConfigValueResult = 2147483647
)

type ESteamNetConnectionEnd int32

const (
	ESteamNetConnectionEndInvalid                       ESteamNetConnectionEnd = 0
	ESteamNetConnectionEndAppMin                        ESteamNetConnectionEnd = 1000
	ESteamNetConnectionEndAppGeneric                    ESteamNetConnectionEnd = 1000
	ESteamNetConnectionEndAppMax                        ESteamNetConnectionEnd = 1999
	ESteamNetConnectionEndAppExceptionMin               ESteamNetConnectionEnd = 2000
	ESteamNetConnectionEndAppExceptionGeneric           ESteamNetConnectionEnd = 2000
	ESteamNetConnectionEndAppExceptionMax               ESteamNetConnectionEnd = 2999
	ESteamNetConnectionEndLocalMin                      ESteamNetConnectionEnd = 3000
	ESteamNetConnectionEndLocalOfflineMode              ESteamNetConnectionEnd = 3001
	ESteamNetConnectionEndLocalManyRelayConnectivity    ESteamNetConnectionEnd = 3002
	ESteamNetConnectionEndLocalHostedServerPrimaryRelay ESteamNetConnectionEnd = 3003
	ESteamNetConnectionEndLocalNetworkConfig            ESteamNetConnectionEnd = 3004
	ESteamNetConnectionEndLocalRights                   ESteamNetConnectionEnd = 3005
	ESteamNetConnectionEndLocalP2PICENoPublicAddresses  ESteamNetConnectionEnd = 3006
	ESteamNetConnectionEndLocalMax                      ESteamNetConnectionEnd = 3999
	ESteamNetConnectionEndRemoteMin                     ESteamNetConnectionEnd = 4000
	ESteamNetConnectionEndRemoteTimeout                 ESteamNetConnectionEnd = 4001
	ESteamNetConnectionEndRemoteBadCrypt                ESteamNetConnectionEnd = 4002
	ESteamNetConnectionEndRemoteBadCert                 ESteamNetConnectionEnd = 4003
	ESteamNetConnectionEndRemoteBadProtocolVersion      ESteamNetConnectionEnd = 4006
	ESteamNetConnectionEndRemoteP2PICENoPublicAddresses ESteamNetConnectionEnd = 4007
	ESteamNetConnectionEndRemoteMax                     ESteamNetConnectionEnd = 4999
	ESteamNetConnectionEndMiscMin                       ESteamNetConnectionEnd = 5000
	ESteamNetConnectionEndMiscGeneric                   ESteamNetConnectionEnd = 5001
	ESteamNetConnectionEndMiscInternalError             ESteamNetConnectionEnd = 5002
	ESteamNetConnectionEndMiscTimeout                   ESteamNetConnectionEnd = 5003
	ESteamNetConnectionEndMiscSteamConnectivity         ESteamNetConnectionEnd = 5005
	ESteamNetConnectionEndMiscNoRelaySessionsToClient   ESteamNetConnectionEnd = 5006
	ESteamNetConnectionEndMiscP2PRendezvous             ESteamNetConnectionEnd = 5008
	ESteamNetConnectionEndMiscP2PNATFirewall            ESteamNetConnectionEnd = 5009
	ESteamNetConnectionEndMiscPeerSentNoConnection      ESteamNetConnectionEnd = 5010
	ESteamNetConnectionEndMiscMax                       ESteamNetConnectionEnd = 5999
	ESteamNetConnectionEndForce32Bit                    ESteamNetConnectionEnd = 2147483647
)
