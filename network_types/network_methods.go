package networktypes

import (
	"fmt"
	"unsafe"

	. "github.com/assemblaj/purego-steamworks"
)

func (nm SteamNetworkingMessage) GetData() []byte {
	return unsafePointerToBytes(nm.Data, nm.Size)
}
func (nm SteamNetworkingMessage) String() string {
	return fmt.Sprintf("Data: %s, Size: %d, MessageNumber:  %d, TimeRecieved: %d\n", string(nm.GetData()), nm.Size, nm.MessageNumber, nm.TimeReceived)
}

func unsafePointerToBytes(ptr unsafe.Pointer, size int32) []byte {
	if ptr == nil || size <= 0 {
		return nil
	}

	// Using Go 1.17+ unsafe.Slice for safety and performance
	return unsafe.Slice((*byte)(ptr), size)
}

const (
	flatAPI_SteamIPAddress_t_IsSet = "SteamAPI_SteamIPAddress_t_IsSet"
)

var (
	steamIPAddress_IsSet func(*SteamIPAddress) bool
)

func (s *SteamIPAddress) IsSet() bool {
	return steamIPAddress_IsSet(s)
}

const (
	flatAPI_SteamNetworkingIPAddr_Clear            = "SteamAPI_SteamNetworkingIPAddr_Clear"
	flatAPI_SteamNetworkingIPAddr_IsIPv6AllZeros   = "SteamAPI_SteamNetworkingIPAddr_IsIPv6AllZeros"
	flatAPI_SteamNetworkingIPAddr_SetIPv6          = "SteamAPI_SteamNetworkingIPAddr_SetIPv6"
	flatAPI_SteamNetworkingIPAddr_SetIPv4          = "SteamAPI_SteamNetworkingIPAddr_SetIPv4"
	flatAPI_SteamNetworkingIPAddr_IsIPv4           = "SteamAPI_SteamNetworkingIPAddr_IsIPv4"
	flatAPI_SteamNetworkingIPAddr_GetIPv4          = "SteamAPI_SteamNetworkingIPAddr_GetIPv4"
	flatAPI_SteamNetworkingIPAddr_SetIPv6LocalHost = "SteamAPI_SteamNetworkingIPAddr_SetIPv6LocalHost"
	flatAPI_SteamNetworkingIPAddr_IsLocalHost      = "SteamAPI_SteamNetworkingIPAddr_IsLocalHost"
	flatAPI_SteamNetworkingIPAddr_ToString         = "SteamAPI_SteamNetworkingIPAddr_ToString"
	flatAPI_SteamNetworkingIPAddr_ParseString      = "SteamAPI_SteamNetworkingIPAddr_ParseString"
	flatAPI_SteamNetworkingIPAddr_IsEqualTo        = "SteamAPI_SteamNetworkingIPAddr_IsEqualTo"
	flatAPI_SteamNetworkingIPAddr_GetFakeIPType    = "SteamAPI_SteamNetworkingIPAddr_GetFakeIPType"
	flatAPI_SteamNetworkingIPAddr_IsFakeIP         = "SteamAPI_SteamNetworkingIPAddr_IsFakeIP"
)

var (
	steamNetworkingIPAddr_Clear            func(self *SteamNetworkingIPAddr)
	steamNetworkingIPAddr_IsIPv6AllZeros   func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_SetIPv6          func(self *SteamNetworkingIPAddr, ipv6 string, Port uint16)
	steamNetworkingIPAddr_SetIPv4          func(self *SteamNetworkingIPAddr, IP uint16, Port uint16)
	steamNetworkingIPAddr_IsIPv4           func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_GetIPv4          func(self *SteamNetworkingIPAddr) uint32
	steamNetworkingIPAddr_SetIPv6LocalHost func(self *SteamNetworkingIPAddr, Port uint16)
	steamNetworkingIPAddr_IsLocalHost      func(self *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_ToString         func(self *SteamNetworkingIPAddr, buf string, cbBuf uint32, bWithPort bool)
	steamNetworkingIPAddr_ParseString      func(self *SteamNetworkingIPAddr, pszStr string) bool
	steamNetworkingIPAddr_IsEqualTo        func(self *SteamNetworkingIPAddr, x *SteamNetworkingIPAddr) bool
	steamNetworkingIPAddr_GetFakeIPType    func(self *SteamNetworkingIPAddr) ESteamNetworkingFakeIPType
	steamNetworkingIPAddr_IsFakeIP         func(self *SteamNetworkingIPAddr) bool
)

func (s *SteamNetworkingIPAddr) Clear() {
	steamNetworkingIPAddr_Clear(s)
}

func (s *SteamNetworkingIPAddr) IsIPv6AllZeros() bool {
	return steamNetworkingIPAddr_IsIPv6AllZeros(s)
}

func (s *SteamNetworkingIPAddr) SetIPv6(ipv6 string, Port uint16) {
	steamNetworkingIPAddr_SetIPv6(s, ipv6, Port)
}

func (s *SteamNetworkingIPAddr) SetIPv4(IP uint16, Port uint16) {
	steamNetworkingIPAddr_SetIPv4(s, IP, Port)
}

func (s *SteamNetworkingIPAddr) IsIPv4() bool {
	return steamNetworkingIPAddr_IsIPv4(s)
}

func (s *SteamNetworkingIPAddr) GetIPv4() uint32 {
	return steamNetworkingIPAddr_GetIPv4(s)
}

func (s *SteamNetworkingIPAddr) SetIPv6LocalHost(Port uint16) {
	steamNetworkingIPAddr_SetIPv6LocalHost(s, Port)
}

func (s *SteamNetworkingIPAddr) IsLocalHost() bool {
	return steamNetworkingIPAddr_IsLocalHost(s)
}

func (s *SteamNetworkingIPAddr) ToString(buf string, cbBuf uint32, bWithPort bool) {
	steamNetworkingIPAddr_ToString(s, buf, cbBuf, bWithPort)
}

func (s *SteamNetworkingIPAddr) ParseString(pszStr string) bool {
	return steamNetworkingIPAddr_ParseString(s, pszStr)
}

func (s *SteamNetworkingIPAddr) IsEqualTo(x *SteamNetworkingIPAddr) bool {
	return steamNetworkingIPAddr_IsEqualTo(s, x)
}

func (s *SteamNetworkingIPAddr) GetFakeIPType() ESteamNetworkingFakeIPType {
	return steamNetworkingIPAddr_GetFakeIPType(s)
}

func (s *SteamNetworkingIPAddr) IsFakeIP() bool {
	return steamNetworkingIPAddr_IsFakeIP(s)
}

const (
	flatAPI_SteamNetworkingIdentity_Clear             = "SteamAPI_SteamNetworkingIdentity_Clear"
	flatAPI_SteamNetworkingIdentity_IsInvalid         = "SteamAPI_SteamNetworkingIdentity_IsInvalid"
	flatAPI_SteamNetworkingIdentity_SetSteamID        = "SteamAPI_SteamNetworkingIdentity_SetSteamID"
	flatAPI_SteamNetworkingIdentity_GetSteamID        = "SteamAPI_SteamNetworkingIdentity_GetSteamID"
	flatAPI_SteamNetworkingIdentity_SetSteamID64      = "SteamAPI_SteamNetworkingIdentity_SetSteamID64"
	flatAPI_SteamNetworkingIdentity_GetSteamID64      = "SteamAPI_SteamNetworkingIdentity_GetSteamID64"
	flatAPI_SteamNetworkingIdentity_SetXboxPairwiseID = "SteamAPI_SteamNetworkingIdentity_SetXboxPairwiseID"
	flatAPI_SteamNetworkingIdentity_GetXboxPairwiseID = "SteamAPI_SteamNetworkingIdentity_GetXboxPairwiseID"
	flatAPI_SteamNetworkingIdentity_SetPSNID          = "SteamAPI_SteamNetworkingIdentity_SetPSNID"
	flatAPI_SteamNetworkingIdentity_GetPSNID          = "SteamAPI_SteamNetworkingIdentity_GetPSNID"
	flatAPI_SteamNetworkingIdentity_SetIPAddr         = "SteamAPI_SteamNetworkingIdentity_SetIPAddr"
	flatAPI_SteamNetworkingIdentity_GetIPAddr         = "SteamAPI_SteamNetworkingIdentity_GetIPAddr"
	flatAPI_SteamNetworkingIdentity_SetIPv4Addr       = "SteamAPI_SteamNetworkingIdentity_SetIPv4Addr"
	flatAPI_SteamNetworkingIdentity_GetIPv4           = "SteamAPI_SteamNetworkingIdentity_GetIPv4"
	flatAPI_SteamNetworkingIdentity_GetFakeIPType     = "SteamAPI_SteamNetworkingIdentity_GetFakeIPType"
	flatAPI_SteamNetworkingIdentity_IsFakeIP          = "SteamAPI_SteamNetworkingIdentity_IsFakeIP"
	flatAPI_SteamNetworkingIdentity_SetLocalHost      = "SteamAPI_SteamNetworkingIdentity_SetLocalHost"
	flatAPI_SteamNetworkingIdentity_IsLocalHost       = "SteamAPI_SteamNetworkingIdentity_IsLocalHost"
	flatAPI_SteamNetworkingIdentity_SetGenericString  = "SteamAPI_SteamNetworkingIdentity_SetGenericString"
	flatAPI_SteamNetworkingIdentity_GetGenericString  = "SteamAPI_SteamNetworkingIdentity_GetGenericString"
	flatAPI_SteamNetworkingIdentity_SetGenericBytes   = "SteamAPI_SteamNetworkingIdentity_SetGenericBytes"
	flatAPI_SteamNetworkingIdentity_GetGenericBytes   = "SteamAPI_SteamNetworkingIdentity_GetGenericBytes"
	flatAPI_SteamNetworkingIdentity_IsEqualTo         = "SteamAPI_SteamNetworkingIdentity_IsEqualTo"
	flatAPI_SteamNetworkingIdentity_ToString          = "SteamAPI_SteamNetworkingIdentity_ToString"
	flatAPI_SteamNetworkingIdentity_ParseString       = "SteamAPI_SteamNetworkingIdentity_ParseString"
)

var (
	steamNetworkingIdentity_Clear             func(self *SteamNetworkingIdentity)
	steamNetworkingIdentity_IsInvalid         func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetSteamID        func(self *SteamNetworkingIdentity, steamID Uint64SteamID)
	steamNetworkingIdentity_GetSteamID        func(self *SteamNetworkingIdentity) Uint64SteamID
	steamNetworkingIdentity_SetSteamID64      func(self *SteamNetworkingIdentity, steamID uint64)
	steamNetworkingIdentity_GetSteamID64      func(self *SteamNetworkingIdentity) uint64
	steamNetworkingIdentity_SetXboxPairwiseID func(self *SteamNetworkingIdentity, pszString string) bool
	steamNetworkingIdentity_GetXboxPairwiseID func(self *SteamNetworkingIdentity) string
	steamNetworkingIdentity_SetPSNID          func(self *SteamNetworkingIdentity, id uint64)
	steamNetworkingIdentity_GetPSNID          func(self *SteamNetworkingIdentity) uint64
	steamNetworkingIdentity_SetIPAddr         func(self *SteamNetworkingIdentity, addr *SteamNetworkingIPAddr)
	steamNetworkingIdentity_GetIPAddr         func(self *SteamNetworkingIdentity) *SteamNetworkingIPAddr
	steamNetworkingIdentity_SetIPv4Addr       func(self *SteamNetworkingIdentity, nIPv4u int32, nPort uint16)
	steamNetworkingIdentity_GetIPv4           func(self *SteamNetworkingIdentity) uint32
	steamNetworkingIdentity_GetFakeIPType     func(self *SteamNetworkingIdentity) ESteamNetworkingFakeIPType
	steamNetworkingIdentity_IsFakeIP          func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetLocalHost      func(self *SteamNetworkingIdentity)
	steamNetworkingIdentity_IsLocalHost       func(self *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_SetGenericString  func(self *SteamNetworkingIdentity, pszString string) bool
	steamNetworkingIdentity_GetGenericString  func(self *SteamNetworkingIdentity) string
	steamNetworkingIdentity_SetGenericBytes   func(self *SteamNetworkingIdentity, data []byte, cbLen uint32) bool
	steamNetworkingIdentity_GetGenericBytes   func(self *SteamNetworkingIdentity, cbLen *int) string
	steamNetworkingIdentity_IsEqualTo         func(self *SteamNetworkingIdentity, x *SteamNetworkingIdentity) bool
	steamNetworkingIdentity_ToString          func(self *SteamNetworkingIdentity, buf string, cbBuf uint32)
	steamNetworkingIdentity_ParseString       func(self *SteamNetworkingIdentity, pszStr string) bool
)

func (s *SteamNetworkingIdentity) Clear() {
	steamNetworkingIdentity_Clear(s)
}

func (s *SteamNetworkingIdentity) IsInvalid() bool {
	return steamNetworkingIdentity_IsInvalid(s)
}
func (s *SteamNetworkingIdentity) SetSteamID(steamID Uint64SteamID) {
	steamNetworkingIdentity_SetSteamID(s, steamID)
}
func (s *SteamNetworkingIdentity) GetSteamID() Uint64SteamID {
	return steamNetworkingIdentity_GetSteamID(s)
}
func (s *SteamNetworkingIdentity) SetSteamID64(steamID uint64) {
	steamNetworkingIdentity_SetSteamID64(s, steamID)
}

func (s *SteamNetworkingIdentity) GetSteamID64() uint64 {
	return steamNetworkingIdentity_GetSteamID64(s)
}
func (s *SteamNetworkingIdentity) SetXboxPairwiseID(String string) bool {
	return steamNetworkingIdentity_SetXboxPairwiseID(s, String)
}
func (s *SteamNetworkingIdentity) GetXboxPairwiseID() string {
	return steamNetworkingIdentity_GetXboxPairwiseID(s)
}
func (s *SteamNetworkingIdentity) SetPSNID(id uint64) {
	steamNetworkingIdentity_SetPSNID(s, id)
}
func (s *SteamNetworkingIdentity) GetPSNID() uint64 {
	return steamNetworkingIdentity_GetPSNID(s)
}

func (s *SteamNetworkingIdentity) SetIPAddr(addr *SteamNetworkingIPAddr) {
	steamNetworkingIdentity_SetIPAddr(s, addr)
}

func (s *SteamNetworkingIdentity) GetIPAddr() *SteamNetworkingIPAddr {
	return steamNetworkingIdentity_GetIPAddr(s)
}

func (s *SteamNetworkingIdentity) SetIPv4Addr(nIPv4u int32, nPort uint16) {
	steamNetworkingIdentity_SetIPv4Addr(s, nIPv4u, nPort)
}
func (s *SteamNetworkingIdentity) GetIPv4() uint32 {
	return steamNetworkingIdentity_GetIPv4(s)
}

func (s *SteamNetworkingIdentity) GetFakeIPType() ESteamNetworkingFakeIPType {
	return steamNetworkingIdentity_GetFakeIPType(s)
}
func (s *SteamNetworkingIdentity) IsFakeIP() bool {
	return steamNetworkingIdentity_IsFakeIP(s)
}
func (s *SteamNetworkingIdentity) SetLocalHost() {
	steamNetworkingIdentity_SetLocalHost(s)
}
func (s *SteamNetworkingIdentity) IsLocalHost() bool {
	return steamNetworkingIdentity_IsLocalHost(s)
}
func (s *SteamNetworkingIdentity) SetGenericString(pszString string) bool {
	return steamNetworkingIdentity_SetGenericString(s, pszString)
}

func (s *SteamNetworkingIdentity) GetGenericString() string {
	return steamNetworkingIdentity_GetGenericString(s)
}

func (s *SteamNetworkingIdentity) SetGenericBytes(data []byte) bool {
	return steamNetworkingIdentity_SetGenericBytes(s, data, uint32(len(data)))
}

func (s *SteamNetworkingIdentity) GetGenericBytes(cbLen *int) string {
	return steamNetworkingIdentity_GetGenericBytes(s, cbLen)
}

func (s *SteamNetworkingIdentity) IsEqualTo(x *SteamNetworkingIdentity) bool {
	return steamNetworkingIdentity_IsEqualTo(s, x)
}

func (s *SteamNetworkingIdentity) ToString(buf string, cbBuf uint32) {
	steamNetworkingIdentity_ToString(s, buf, cbBuf)
}

func (s *SteamNetworkingIdentity) ParseString(pszStr string) bool {
	return steamNetworkingIdentity_ParseString(s, pszStr)
}

const (
	flatAPI_SteamNetworkingConfigValue_t_SetInt32  = "SteamAPI_SteamNetworkingConfigValue_t_SetInt32"
	flatAPI_SteamNetworkingConfigValue_t_SetInt64  = "SteamAPI_SteamNetworkingConfigValue_t_SetInt64"
	flatAPI_SteamNetworkingConfigValue_t_SetFloat  = "SteamAPI_SteamNetworkingConfigValue_t_SetFloat"
	flatAPI_SteamNetworkingConfigValue_t_SetPtr    = "SteamAPI_SteamNetworkingConfigValue_t_SetPtr"
	flatAPI_SteamNetworkingConfigValue_t_SetString = "SteamAPI_SteamNetworkingConfigValue_t_SetString"
)

var (
	steamNetworkingConfigValue_t_SetInt32  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int32)
	steamNetworkingConfigValue_t_SetInt64  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data int64)
	steamNetworkingConfigValue_t_SetFloat  func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data float32)
	steamNetworkingConfigValue_t_SetPtr    func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data []byte)
	steamNetworkingConfigValue_t_SetString func(self *SteamNetworkingConfigValue, eVal ESteamNetworkingConfigValue, data string)
)

func (s *SteamNetworkingConfigValue) SetInt32(eVal ESteamNetworkingConfigValue, data int32) {
	steamNetworkingConfigValue_t_SetInt32(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetInt64(eVal ESteamNetworkingConfigValue, data int64) {
	steamNetworkingConfigValue_t_SetInt64(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetFloat(eVal ESteamNetworkingConfigValue, data float32) {
	steamNetworkingConfigValue_t_SetFloat(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetPtr(eVal ESteamNetworkingConfigValue, data []byte) {
	steamNetworkingConfigValue_t_SetPtr(s, eVal, data)
}
func (s *SteamNetworkingConfigValue) SetString(eVal ESteamNetworkingConfigValue, data string) {
	steamNetworkingConfigValue_t_SetString(s, eVal, data)
}

const (
	flatAPI_SteamDatagramHostedAddress_Clear         = "SteamAPI_SteamDatagramHostedAddress_Clear"
	flatAPI_SteamDatagramHostedAddress_GetPopID      = "SteamAPI_SteamDatagramHostedAddress_GetPopID"
	flatAPI_SteamDatagramHostedAddress_SetDevAddress = "SteamAPI_SteamDatagramHostedAddress_SetDevAddress"
)

var (
	steamDatagramHostedAddress_Clear         func(self *SteamDatagramHostedAddress)
	steamDatagramHostedAddress_GetPopID      func(self *SteamDatagramHostedAddress) SteamNetworkingPOPID
	steamDatagramHostedAddress_SetDevAddress func(self *SteamDatagramHostedAddress, nIP uint32, nPort uint16, popid SteamNetworkingPOPID)
)

func (s *SteamDatagramHostedAddress) Clear() {
	steamDatagramHostedAddress_Clear(s)
}

func (s *SteamDatagramHostedAddress) GetPopID() SteamNetworkingPOPID {
	return steamDatagramHostedAddress_GetPopID(s)
}

func (s *SteamDatagramHostedAddress) SetDevAddress(nIP uint32, nPort uint16, popid SteamNetworkingPOPID) {
	steamDatagramHostedAddress_SetDevAddress(s, nIP, nPort, popid)
}

var (
	iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort  func()
	iSteamNetworkingFakeUDPPort_SendMessageToFakeIP func(remoteAddress *SteamNetworkingIPAddr, pData uintptr, cbData uint32, nSendFlags int32) EResult
	iSteamNetworkingFakeUDPPort_ReceiveMessages     func(ppOutMessages *[]SteamNetworkingMessage, nMaxMessages int32) int32
	iSteamNetworkingFakeUDPPort_ScheduleCleanup     func(remoteAddress *SteamNetworkingIPAddr)
)

const (
	flatAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort  = "SteamAPI_ISteamNetworkingFakeUDPPort_DestroyFakeUDPPort"
	flatAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP = "SteamAPI_ISteamNetworkingFakeUDPPort_SendMessageToFakeIP"
	flatAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages     = "SteamAPI_ISteamNetworkingFakeUDPPort_ReceiveMessages"
	flatAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup     = "SteamAPI_ISteamNetworkingFakeUDPPort_ScheduleCleanup"
)

func (s *SteamNetworkingFakeUDPPort) DestroyFakeUDPPort() {
	iSteamNetworkingFakeUDPPort_DestroyFakeUDPPort()
}

func (s *SteamNetworkingFakeUDPPort) SendMessageToFakeIP(remoteAddress *SteamNetworkingIPAddr, pData uintptr, cbData uint32, nSendFlags int32) EResult {
	return iSteamNetworkingFakeUDPPort_SendMessageToFakeIP(remoteAddress, pData, cbData, nSendFlags)
}

func (s *SteamNetworkingFakeUDPPort) ReceiveMessages(nMaxMessages int32) []SteamNetworkingMessage {
	ppOutMessages := make([]SteamNetworkingMessage, 0, nMaxMessages)
	actual := iSteamNetworkingFakeUDPPort_ReceiveMessages(&ppOutMessages, nMaxMessages)
	return ppOutMessages[:actual]
}

func (s *SteamNetworkingFakeUDPPort) ScheduleCleanup(remoteAddress *SteamNetworkingIPAddr) {
	iSteamNetworkingFakeUDPPort_ScheduleCleanup(remoteAddress)
}
