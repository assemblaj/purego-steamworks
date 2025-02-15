package steamworks

import "unsafe"

type InputHandle_t uint64

type ESteamInputType int32

const (
	ESteamInputType_Unknown              ESteamInputType = 0
	ESteamInputType_SteamController      ESteamInputType = 1
	ESteamInputType_XBox360Controller    ESteamInputType = 2
	ESteamInputType_XBoxOneController    ESteamInputType = 3
	ESteamInputType_GenericXInput        ESteamInputType = 4
	ESteamInputType_PS4Controller        ESteamInputType = 5
	ESteamInputType_AppleMFiController   ESteamInputType = 6 // Unused
	ESteamInputType_AndroidController    ESteamInputType = 7 // Unused
	ESteamInputType_SwitchJoyConPair     ESteamInputType = 8 // Unused
	ESteamInputType_SwitchJoyConSingle   ESteamInputType = 9 // Unused
	ESteamInputType_SwitchProController  ESteamInputType = 10
	ESteamInputType_MobileTouch          ESteamInputType = 11
	ESteamInputType_PS3Controller        ESteamInputType = 12
	ESteamInputType_PS5Controller        ESteamInputType = 13
	ESteamInputType_SteamDeckController  ESteamInputType = 14
	ESteamInputType_Count                ESteamInputType = 15
	ESteamInputType_MaximumPossibleValue ESteamInputType = 255
)

const (
	_STEAM_INPUT_MAX_COUNT = 16
)

type EFloatingGamepadTextInputMode int32

const (
	EFloatingGamepadTextInputMode_ModeSingleLine    EFloatingGamepadTextInputMode = 0
	EFloatingGamepadTextInputMode_ModeMultipleLines EFloatingGamepadTextInputMode = 1
	EFloatingGamepadTextInputMode_ModeEmail         EFloatingGamepadTextInputMode = 2
	EFloatingGamepadTextInputMode_ModeNumeric       EFloatingGamepadTextInputMode = 3
)

const (
	flatAPI_SteamInput                          = "SteamAPI_SteamInput_v006"
	flatAPI_ISteamInput_GetConnectedControllers = "SteamAPI_ISteamInput_GetConnectedControllers"
	flatAPI_ISteamInput_GetInputTypeForHandle   = "SteamAPI_ISteamInput_GetInputTypeForHandle"
	flatAPI_ISteamInput_Init                    = "SteamAPI_ISteamInput_Init"
	flatAPI_ISteamInput_RunFrame                = "SteamAPI_ISteamInput_RunFrame"
)

type steamInput uintptr

var (
	steamInput_init         func() uintptr
	getConnectedControllers func(steamInput uintptr, handles uintptr) uintptr
	getInputTypeForHandle   func(steamInput uintptr, inputHandle uintptr) uintptr
	steamInputInit          func(steamInput uintptr, callRunFrame uintptr) uintptr
	steamInputRunFrame      func(steamInput uintptr, zero uintptr) uintptr
)

func (s steamInput) GetConnectedControllers() []InputHandle_t {
	var handles [_STEAM_INPUT_MAX_COUNT]InputHandle_t
	result := getConnectedControllers(uintptr(s), uintptr(unsafe.Pointer(&handles[0])))
	return handles[:int(result)]
}

func (s steamInput) GetInputTypeForHandle(inputHandle InputHandle_t) ESteamInputType {
	result := getInputTypeForHandle(uintptr(s), uintptr(inputHandle))
	return ESteamInputType(result)
}

func (s steamInput) Init(bExplicitlyCallRunFrame bool) bool {
	var callRunFrame uintptr
	if bExplicitlyCallRunFrame {
		callRunFrame = 1
	}
	result := steamInputInit(uintptr(s), callRunFrame)
	return byte(result) != 0
}

func (s steamInput) RunFrame() {
	steamInputRunFrame(uintptr(s), uintptr(0))
}
