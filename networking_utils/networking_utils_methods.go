package networkingutils

import (
	. "github.com/assemblaj/purego-steamworks/network_types"

	common "github.com/assemblaj/purego-steamworks"
)

type intptr int64
type size uint64

var (
	steamNetworkingUtils_init                                                   func() uintptr
	iSteamNetworkingUtils_AllocateMessage                                       func(steamNetworkingUtils uintptr, cbAllocateBuffer int32) *SteamNetworkingMessage
	iSteamNetworkingUtils_InitRelayNetworkAccess                                func(steamNetworkingUtils uintptr)
	iSteamNetworkingUtils_GetRelayNetworkStatus                                 func(steamNetworkingUtils uintptr, pDetails *SteamRelayNetworkStatus) ESteamNetworkingAvailability
	iSteamNetworkingUtils_GetLocalPingLocation                                  func(steamNetworkingUtils uintptr, result *SteamNetworkPingLocation) float32
	iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations                   func(steamNetworkingUtils uintptr, location1 uintptr, location2 uintptr) int32
	iSteamNetworkingUtils_EstimatePingTimeFromLocalHost                         func(steamNetworkingUtils uintptr, remoteLocation uintptr) int32
	iSteamNetworkingUtils_ConvertPingLocationToString                           func(steamNetworkingUtils uintptr, location uintptr, pszBuf *[]byte, cchBufSize int32)
	iSteamNetworkingUtils_ParsePingLocationString                               func(steamNetworkingUtils uintptr, pszString string, result *SteamNetworkPingLocation) bool
	iSteamNetworkingUtils_CheckPingDataUpToDate                                 func(steamNetworkingUtils uintptr, flMaxAgeSeconds float32) bool
	iSteamNetworkingUtils_GetPingToDataCenter                                   func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID, pViaRelayPoP *SteamNetworkingPOPID) int32
	iSteamNetworkingUtils_GetDirectPingToPOP                                    func(steamNetworkingUtils uintptr, popID SteamNetworkingPOPID) int32
	iSteamNetworkingUtils_GetPOPCount                                           func(steamNetworkingUtils uintptr) int32
	iSteamNetworkingUtils_GetPOPList                                            func(steamNetworkingUtils uintptr, list *[]SteamNetworkingPOPID, nListSz int32) int32
	iSteamNetworkingUtils_GetLocalTimestamp                                     func(steamNetworkingUtils uintptr) SteamNetworkingMicroseconds
	iSteamNetworkingUtils_SetDebugOutputFunction                                func(steamNetworkingUtils uintptr, eDetailLevel ESteamNetworkingSocketsDebugOutputType, pfnFunc FSteamNetworkingSocketsDebugOutput)
	iSteamNetworkingUtils_IsFakeIPv4                                            func(steamNetworkingUtils uintptr, nIPv4 uint32) bool
	iSteamNetworkingUtils_GetIPv4FakeIPType                                     func(steamNetworkingUtils uintptr, nIPv4 uint32) ESteamNetworkingFakeIPType
	iSteamNetworkingUtils_GetRealIdentityForFakeIP                              func(steamNetworkingUtils uintptr, fakeIP uintptr, pOutRealIdentity *SteamNetworkingIdentity) common.EResult
	iSteamNetworkingUtils_SetGlobalConfigValueInt32                             func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val int32) bool
	iSteamNetworkingUtils_SetGlobalConfigValueFloat                             func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val float32) bool
	iSteamNetworkingUtils_SetGlobalConfigValueString                            func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val string) bool
	iSteamNetworkingUtils_SetGlobalConfigValuePtr                               func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, val []byte) bool
	iSteamNetworkingUtils_SetConnectionConfigValueInt32                         func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool
	iSteamNetworkingUtils_SetConnectionConfigValueFloat                         func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val float32) bool
	iSteamNetworkingUtils_SetConnectionConfigValueString                        func(steamNetworkingUtils uintptr, hConn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged     func(steamNetworkingUtils uintptr, fnCallback FnSteamNetConnectionStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged func(steamNetworkingUtils uintptr, fnCallback FnSteamNetAuthenticationStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged      func(steamNetworkingUtils uintptr, fnCallback FnSteamRelayNetworkStatusChanged) bool
	iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult                        func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingFakeIPResult) bool
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest              func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionRequest) bool
	iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed               func(steamNetworkingUtils uintptr, fnCallback FnSteamNetworkingMessagesSessionFailed) bool
	iSteamNetworkingUtils_SetConfigValue                                        func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool
	iSteamNetworkingUtils_SetConfigValueStruct                                  func(steamNetworkingUtils uintptr, opt uintptr, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool
	iSteamNetworkingUtils_GetConfigValue                                        func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, pOutDataType *ESteamNetworkingConfigDataType, pResult *[]byte, cbResult *size) ESteamNetworkingGetConfigValueResult
	iSteamNetworkingUtils_GetConfigValueInfo                                    func(steamNetworkingUtils uintptr, eValue ESteamNetworkingConfigValue, pOutDataType *ESteamNetworkingConfigDataType, pOutScope *ESteamNetworkingConfigScope) []byte
	iSteamNetworkingUtils_IterateGenericEditableConfigValues                    func(steamNetworkingUtils uintptr, eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString                        func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr, buf *[]byte, cbBuf uint32, bWithPort bool)
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString                     func(steamNetworkingUtils uintptr, pAddr *SteamNetworkingIPAddr, pszStr string) bool
	iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType                   func(steamNetworkingUtils uintptr, addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType
	iSteamNetworkingUtils_SteamNetworkingIdentity_ToString                      func(steamNetworkingUtils uintptr, identity *SteamNetworkingIdentity, buf *[]byte, cbBuf uint32)
	iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString                   func(steamNetworkingUtils uintptr, pIdentity *SteamNetworkingIdentity, pszStr string) bool
)

const (
	flatAPI_SteamNetworkingUtils                                                        = "SteamAPI_SteamNetworkingUtils_SteamAPI_v004"
	flatAPI_ISteamNetworkingUtils_AllocateMessage                                       = "SteamAPI_ISteamNetworkingUtils_AllocateMessage"
	flatAPI_ISteamNetworkingUtils_InitRelayNetworkAccess                                = "SteamAPI_ISteamNetworkingUtils_InitRelayNetworkAccess"
	flatAPI_ISteamNetworkingUtils_GetRelayNetworkStatus                                 = "SteamAPI_ISteamNetworkingUtils_GetRelayNetworkStatus"
	flatAPI_ISteamNetworkingUtils_GetLocalPingLocation                                  = "SteamAPI_ISteamNetworkingUtils_GetLocalPingLocation"
	flatAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations                   = "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations"
	flatAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost                         = "SteamAPI_ISteamNetworkingUtils_EstimatePingTimeFromLocalHost"
	flatAPI_ISteamNetworkingUtils_ConvertPingLocationToString                           = "SteamAPI_ISteamNetworkingUtils_ConvertPingLocationToString"
	flatAPI_ISteamNetworkingUtils_ParsePingLocationString                               = "SteamAPI_ISteamNetworkingUtils_ParsePingLocationString"
	flatAPI_ISteamNetworkingUtils_CheckPingDataUpToDate                                 = "SteamAPI_ISteamNetworkingUtils_CheckPingDataUpToDate"
	flatAPI_ISteamNetworkingUtils_GetPingToDataCenter                                   = "SteamAPI_ISteamNetworkingUtils_GetPingToDataCenter"
	flatAPI_ISteamNetworkingUtils_GetDirectPingToPOP                                    = "SteamAPI_ISteamNetworkingUtils_GetDirectPingToPOP"
	flatAPI_ISteamNetworkingUtils_GetPOPCount                                           = "SteamAPI_ISteamNetworkingUtils_GetPOPCount"
	flatAPI_ISteamNetworkingUtils_GetPOPList                                            = "SteamAPI_ISteamNetworkingUtils_GetPOPList"
	flatAPI_ISteamNetworkingUtils_GetLocalTimestamp                                     = "SteamAPI_ISteamNetworkingUtils_GetLocalTimestamp"
	flatAPI_ISteamNetworkingUtils_SetDebugOutputFunction                                = "SteamAPI_ISteamNetworkingUtils_SetDebugOutputFunction"
	flatAPI_ISteamNetworkingUtils_IsFakeIPv4                                            = "SteamAPI_ISteamNetworkingUtils_IsFakeIPv4"
	flatAPI_ISteamNetworkingUtils_GetIPv4FakeIPType                                     = "SteamAPI_ISteamNetworkingUtils_GetIPv4FakeIPType"
	flatAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP                              = "SteamAPI_ISteamNetworkingUtils_GetRealIdentityForFakeIP"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32                             = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueInt32"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat                             = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueFloat"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValueString                            = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValueString"
	flatAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr                               = "SteamAPI_ISteamNetworkingUtils_SetGlobalConfigValuePtr"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32                         = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueInt32"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat                         = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueFloat"
	flatAPI_ISteamNetworkingUtils_SetConnectionConfigValueString                        = "SteamAPI_ISteamNetworkingUtils_SetConnectionConfigValueString"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged     = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged      = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult                        = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_FakeIPResult"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest              = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest"
	flatAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed               = "SteamAPI_ISteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed"
	flatAPI_ISteamNetworkingUtils_SetConfigValue                                        = "SteamAPI_ISteamNetworkingUtils_SetConfigValue"
	flatAPI_ISteamNetworkingUtils_SetConfigValueStruct                                  = "SteamAPI_ISteamNetworkingUtils_SetConfigValueStruct"
	flatAPI_ISteamNetworkingUtils_GetConfigValue                                        = "SteamAPI_ISteamNetworkingUtils_GetConfigValue"
	flatAPI_ISteamNetworkingUtils_GetConfigValueInfo                                    = "SteamAPI_ISteamNetworkingUtils_GetConfigValueInfo"
	flatAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues                    = "SteamAPI_ISteamNetworkingUtils_IterateGenericEditableConfigValues"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString                        = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ToString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString                     = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_ParseString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType                   = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString                      = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ToString"
	flatAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString                   = "SteamAPI_ISteamNetworkingUtils_SteamNetworkingIdentity_ParseString"
)

var steamNetworkingUtils uintptr

func AllocateMessage(AllocateBufferSize int32) *SteamNetworkingMessage {
	return iSteamNetworkingUtils_AllocateMessage(steamNetworkingUtils, AllocateBufferSize)
}

func InitRelayNetworkAccess() {
	iSteamNetworkingUtils_InitRelayNetworkAccess(steamNetworkingUtils)
}

func GetRelayNetworkStatus() (Details SteamRelayNetworkStatus, availability ESteamNetworkingAvailability) {
	availability = iSteamNetworkingUtils_GetRelayNetworkStatus(steamNetworkingUtils, &Details)
	return Details, availability
}

func GetLocalPingLocation() (result SteamNetworkPingLocation, ageInSeconds float32) {
	ageInSeconds = iSteamNetworkingUtils_GetLocalPingLocation(steamNetworkingUtils, &result)
	return result, ageInSeconds
}

func EstimatePingTimeBetweenTwoLocations(location1 SteamNetworkPingLocation, location2 SteamNetworkPingLocation) int32 {
	return iSteamNetworkingUtils_EstimatePingTimeBetweenTwoLocations(steamNetworkingUtils, common.StructToUintptr[SteamNetworkPingLocation](&location1), common.StructToUintptr[SteamNetworkPingLocation](&location2))
}

func EstimatePingTimeFromLocalHost(remoteLocation SteamNetworkPingLocation) int32 {
	return iSteamNetworkingUtils_EstimatePingTimeFromLocalHost(steamNetworkingUtils, common.StructToUintptr[SteamNetworkPingLocation](&remoteLocation))
}

func ConvertPingLocationToString(location SteamNetworkPingLocation, BufSize int32) string {
	var pszBuf []byte = make([]byte, 0, BufSize)
	iSteamNetworkingUtils_ConvertPingLocationToString(steamNetworkingUtils, common.StructToUintptr[SteamNetworkPingLocation](&location), &pszBuf, BufSize)
	return string(pszBuf)
}

func ParsePingLocationString(String string) (result SteamNetworkPingLocation, success bool) {
	success = iSteamNetworkingUtils_ParsePingLocationString(steamNetworkingUtils, String, &result)
	return result, success
}

func CheckPingDataUpToDate(MaxAgeSeconds float32) bool {
	return iSteamNetworkingUtils_CheckPingDataUpToDate(steamNetworkingUtils, MaxAgeSeconds)
}

func GetPingToDataCenter(popID SteamNetworkingPOPID) (ViaRelayPoP SteamNetworkingPOPID, pingTime int32) {
	pingTime = iSteamNetworkingUtils_GetPingToDataCenter(steamNetworkingUtils, popID, &ViaRelayPoP)
	return ViaRelayPoP, pingTime
}

func GetDirectPingToPOP(popID SteamNetworkingPOPID) int32 {
	return iSteamNetworkingUtils_GetDirectPingToPOP(steamNetworkingUtils, popID)
}

func GetPOPCount() int32 {
	return iSteamNetworkingUtils_GetPOPCount(steamNetworkingUtils)
}

func GetPOPList(ListSz int32) (list []SteamNetworkingPOPID) {
	list = make([]SteamNetworkingPOPID, 0, ListSz)
	actual := iSteamNetworkingUtils_GetPOPList(steamNetworkingUtils, &list, ListSz)
	return list[:actual]
}

func GetLocalTimestamp() SteamNetworkingMicroseconds {
	return iSteamNetworkingUtils_GetLocalTimestamp(steamNetworkingUtils)
}

func SetDebugOutputFunction(DetailLevel ESteamNetworkingSocketsDebugOutputType, Func FSteamNetworkingSocketsDebugOutput) {
	iSteamNetworkingUtils_SetDebugOutputFunction(steamNetworkingUtils, DetailLevel, Func)
}

func IsFakeIPv4(IPv4 uint32) bool {
	return iSteamNetworkingUtils_IsFakeIPv4(steamNetworkingUtils, IPv4)
}

func GetIPv4FakeIPType(IPv4 uint32) ESteamNetworkingFakeIPType {
	return iSteamNetworkingUtils_GetIPv4FakeIPType(steamNetworkingUtils, IPv4)
}

func GetRealIdentityForFakeIP(fakeIP SteamNetworkingIPAddr) (OutRealIdentity SteamNetworkingIdentity, result common.EResult) {
	result = iSteamNetworkingUtils_GetRealIdentityForFakeIP(steamNetworkingUtils, common.StructToUintptr[SteamNetworkingIPAddr](&fakeIP), &OutRealIdentity)
	return OutRealIdentity, result
}

func SetGlobalConfigValueInt32(eValue ESteamNetworkingConfigValue, val int32) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueInt32(steamNetworkingUtils, eValue, val)
}

func SetGlobalConfigValueFloat(eValue ESteamNetworkingConfigValue, val float32) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueFloat(steamNetworkingUtils, eValue, val)
}

func SetGlobalConfigValueString(eValue ESteamNetworkingConfigValue, val string) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValueString(steamNetworkingUtils, eValue, val)
}

func SetGlobalConfigValuePtr(eValue ESteamNetworkingConfigValue, val []byte) bool {
	return iSteamNetworkingUtils_SetGlobalConfigValuePtr(steamNetworkingUtils, eValue, val)
}

func SetConnectionConfigValueInt32(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val int32) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueInt32(steamNetworkingUtils, Conn, eValue, val)
}

func SetConnectionConfigValueFloat(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val float32) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueFloat(steamNetworkingUtils, Conn, eValue, val)
}

func SetConnectionConfigValueString(Conn HSteamNetConnection, eValue ESteamNetworkingConfigValue, val string) bool {
	return iSteamNetworkingUtils_SetConnectionConfigValueString(steamNetworkingUtils, Conn, eValue, val)
}

func SetGlobalCallback_SteamNetConnectionStatusChanged(fnCallback FnSteamNetConnectionStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamNetConnectionStatusChanged(steamNetworkingUtils, fnCallback)
}

func SetGlobalCallback_SteamNetAuthenticationStatusChanged(fnCallback FnSteamNetAuthenticationStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamNetAuthenticationStatusChanged(steamNetworkingUtils, fnCallback)
}

func SetGlobalCallback_SteamRelayNetworkStatusChanged(fnCallback FnSteamRelayNetworkStatusChanged) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_SteamRelayNetworkStatusChanged(steamNetworkingUtils, fnCallback)
}

func SetGlobalCallback_FakeIPResult(fnCallback FnSteamNetworkingFakeIPResult) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_FakeIPResult(steamNetworkingUtils, fnCallback)
}

func SetGlobalCallback_MessagesSessionRequest(fnCallback FnSteamNetworkingMessagesSessionRequest) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionRequest(steamNetworkingUtils, fnCallback)
}

func SetGlobalCallback_MessagesSessionFailed(fnCallback FnSteamNetworkingMessagesSessionFailed) bool {
	return iSteamNetworkingUtils_SetGlobalCallback_MessagesSessionFailed(steamNetworkingUtils, fnCallback)
}

func SetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, eDataType ESteamNetworkingConfigDataType, pArg []byte) bool {
	return iSteamNetworkingUtils_SetConfigValue(steamNetworkingUtils, eValue, eScopeType, scopeObj, eDataType, pArg)
}

func SetConfigValueStruct(opt SteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr) bool {
	return iSteamNetworkingUtils_SetConfigValueStruct(steamNetworkingUtils, common.StructToUintptr[SteamNetworkingConfigValue](&opt), eScopeType, scopeObj)
}

func GetConfigValue(eValue ESteamNetworkingConfigValue, eScopeType ESteamNetworkingConfigScope, scopeObj intptr, cbResult size) (pOutDataType ESteamNetworkingConfigDataType, pResult []byte, result ESteamNetworkingGetConfigValueResult) {
	pResult = make([]byte, 0, cbResult)
	result = iSteamNetworkingUtils_GetConfigValue(steamNetworkingUtils, eValue, eScopeType, scopeObj, &pOutDataType, &pResult, &cbResult)
	return pOutDataType, pResult, result
}

func GetConfigValueInfo(eValue ESteamNetworkingConfigValue) (pOutDataType ESteamNetworkingConfigDataType, pOutScope ESteamNetworkingConfigScope, name string) {
	nameBuf := iSteamNetworkingUtils_GetConfigValueInfo(steamNetworkingUtils, eValue, &pOutDataType, &pOutScope)
	if nameBuf == nil {
		return pOutDataType, pOutScope, ""
	}
	return pOutDataType, pOutScope, string(nameBuf)
}

func IterateGenericEditableConfigValues(eCurrent ESteamNetworkingConfigValue, bEnumerateDevVars bool) ESteamNetworkingConfigValue {
	return iSteamNetworkingUtils_IterateGenericEditableConfigValues(steamNetworkingUtils, eCurrent, bEnumerateDevVars)
}

func SteamNetworkingIPAddr_ToString(addr *SteamNetworkingIPAddr, cbBuf uint32, bWithPort bool) string {
	var buf []byte = make([]byte, 0, cbBuf)
	iSteamNetworkingUtils_SteamNetworkingIPAddr_ToString(steamNetworkingUtils, addr, &buf, cbBuf, bWithPort)
	return string(buf)
}

func SteamNetworkingIPAddr_ParseString(pAddr *SteamNetworkingIPAddr, pszStr string) bool {
	return iSteamNetworkingUtils_SteamNetworkingIPAddr_ParseString(steamNetworkingUtils, pAddr, pszStr)
}

func SteamNetworkingIPAddr_GetFakeIPType(addr *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType {
	return iSteamNetworkingUtils_SteamNetworkingIPAddr_GetFakeIPType(steamNetworkingUtils, addr)
}

func SteamNetworkingIdentity_ToString(identity *SteamNetworkingIdentity, cbBuf uint32) string {
	var buf []byte = make([]byte, 0, cbBuf)
	iSteamNetworkingUtils_SteamNetworkingIdentity_ToString(steamNetworkingUtils, identity, &buf, cbBuf)
	return string(buf)
}

func SteamNetworkingIdentity_ParseString(pIdentity *SteamNetworkingIdentity, pszStr string) bool {
	return iSteamNetworkingUtils_SteamNetworkingIdentity_ParseString(steamNetworkingUtils, pIdentity, pszStr)
}
