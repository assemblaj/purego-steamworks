package input

import common "github.com/assemblaj/purego-steamworks"

const (
	SteamInputDeviceConnectedID     common.SteamCallbackID = 2801
	SteamInputDeviceDisconnectedID  common.SteamCallbackID = 2802
	SteamInputConfigurationLoadedID common.SteamCallbackID = 2803
	SteamInputGamepadSlotChangeID   common.SteamCallbackID = 2804
)

type SteamInputDeviceConnected struct {
	ConnectedDeviceHandle InputHandle
}
type SteamInputDeviceDisconnected struct {
	DisconnectedDeviceHandle InputHandle
}
type SteamInputConfigurationLoaded struct {
	AppID             common.AppId
	DeviceHandle      InputHandle
	MappingCreator    common.CSteamID
	MajorRevision     uint32
	MinorRevision     uint32
	UsesSteamInputAPI bool
	UsesGamepadAPI    bool
	_                 [2]byte
}
type SteamInputGamepadSlotChange struct {
	AppID          common.AppId
	DeviceHandle   InputHandle
	DeviceType     ESteamInputType
	OldGamepadSlot int32
	NewGamepadSlot int32
}

func (cb SteamInputDeviceConnected) CallbackID() common.SteamCallbackID {
	return SteamInputDeviceConnectedID
}

func (cb SteamInputDeviceConnected) Name() string {
	return "SteamInputDeviceConnected"
}

func (cb SteamInputDeviceConnected) String() string {
	return common.CallbackString(cb)
}

func (cb SteamInputDeviceDisconnected) CallbackID() common.SteamCallbackID {
	return SteamInputDeviceDisconnectedID
}

func (cb SteamInputDeviceDisconnected) Name() string {
	return "SteamInputDeviceDisconnected"
}

func (cb SteamInputDeviceDisconnected) String() string {
	return common.CallbackString(cb)
}

func (cb SteamInputConfigurationLoaded) CallbackID() common.SteamCallbackID {
	return SteamInputConfigurationLoadedID
}

func (cb SteamInputConfigurationLoaded) Name() string {
	return "SteamInputConfigurationLoaded"
}

func (cb SteamInputConfigurationLoaded) String() string {
	return common.CallbackString(cb)
}

func (cb SteamInputGamepadSlotChange) CallbackID() common.SteamCallbackID {
	return SteamInputGamepadSlotChangeID
}

func (cb SteamInputGamepadSlotChange) Name() string {
	return "SteamInputGamepadSlotChange"
}

func (cb SteamInputGamepadSlotChange) String() string {
	return common.CallbackString(cb)
}
