package input

import (
	"unsafe"

	common "github.com/assemblaj/purego-steamworks"
)

const STEAM_INPUT_MAX_COUNT = 16

const STEAM_INPUT_MAX_ANALOG_ACTIONS = 24

const STEAM_INPUT_MAX_DIGITAL_ACTIONS = 256

const STEAM_INPUT_MAX_ORIGINS = 8

const STEAM_INPUT_MAX_ACTIVE_LAYERS = 16

type ESteamControllerPad int32

const (
	ESteamControllerPadLeft  ESteamControllerPad = 0
	ESteamControllerPadRight ESteamControllerPad = 1
)

type SteamInputActionEvent struct {
	ControllerHandle InputHandle
	EventType        ESteamInputActionEventType
	Action           [unsafe.Sizeof(DigitalAction{})]byte
}

func (e *SteamInputActionEvent) AnalogActionPtr() *AnalogAction {
	return (*AnalogAction)(unsafe.Pointer(&e.Action))
}
func (e *SteamInputActionEvent) DigitalActionPtr() *DigitalAction {
	return (*DigitalAction)(unsafe.Pointer(&e.Action))
}

type InputMotionData struct {
	RotQuatX  float32
	RotQuatY  float32
	RotQuatZ  float32
	RotQuatW  float32
	PosAccelX float32
	PosAccelY float32
	PosAccelZ float32
	RotVelX   float32
	RotVelY   float32
	RotVelZ   float32
}
type EInputSourceMode int32

const (
	EInputSourceModeNone           EInputSourceMode = 0
	EInputSourceModeDpad           EInputSourceMode = 1
	EInputSourceModeButtons        EInputSourceMode = 2
	EInputSourceModeFourButtons    EInputSourceMode = 3
	EInputSourceModeAbsoluteMouse  EInputSourceMode = 4
	EInputSourceModeRelativeMouse  EInputSourceMode = 5
	EInputSourceModeJoystickMove   EInputSourceMode = 6
	EInputSourceModeJoystickMouse  EInputSourceMode = 7
	EInputSourceModeJoystickCamera EInputSourceMode = 8
	EInputSourceModeScrollWheel    EInputSourceMode = 9
	EInputSourceModeTrigger        EInputSourceMode = 10
	EInputSourceModeTouchMenu      EInputSourceMode = 11
	EInputSourceModeMouseJoystick  EInputSourceMode = 12
	EInputSourceModeMouseRegion    EInputSourceMode = 13
	EInputSourceModeRadialMenu     EInputSourceMode = 14
	EInputSourceModeSingleButton   EInputSourceMode = 15
	EInputSourceModeSwitches       EInputSourceMode = 16
)

type EXboxOrigin int32

const (
	EXboxOriginA                   EXboxOrigin = 0
	EXboxOriginB                   EXboxOrigin = 1
	EXboxOriginX                   EXboxOrigin = 2
	EXboxOriginY                   EXboxOrigin = 3
	EXboxOriginLeftBumper          EXboxOrigin = 4
	EXboxOriginRightBumper         EXboxOrigin = 5
	EXboxOriginMenu                EXboxOrigin = 6
	EXboxOriginView                EXboxOrigin = 7
	EXboxOriginLeftTriggerPull     EXboxOrigin = 8
	EXboxOriginLeftTriggerClick    EXboxOrigin = 9
	EXboxOriginRightTriggerPull    EXboxOrigin = 10
	EXboxOriginRightTriggerClick   EXboxOrigin = 11
	EXboxOriginLeftStickMove       EXboxOrigin = 12
	EXboxOriginLeftStickClick      EXboxOrigin = 13
	EXboxOriginLeftStickDPadNorth  EXboxOrigin = 14
	EXboxOriginLeftStickDPadSouth  EXboxOrigin = 15
	EXboxOriginLeftStickDPadWest   EXboxOrigin = 16
	EXboxOriginLeftStickDPadEast   EXboxOrigin = 17
	EXboxOriginRightStickMove      EXboxOrigin = 18
	EXboxOriginRightStickClick     EXboxOrigin = 19
	EXboxOriginRightStickDPadNorth EXboxOrigin = 20
	EXboxOriginRightStickDPadSouth EXboxOrigin = 21
	EXboxOriginRightStickDPadWest  EXboxOrigin = 22
	EXboxOriginRightStickDPadEast  EXboxOrigin = 23
	EXboxOriginDPadNorth           EXboxOrigin = 24
	EXboxOriginDPadSouth           EXboxOrigin = 25
	EXboxOriginDPadWest            EXboxOrigin = 26
	EXboxOriginDPadEast            EXboxOrigin = 27
	EXboxOriginCount               EXboxOrigin = 28
)

type InputDigitalActionData struct {
	State  bool
	Active bool
}
type InputAnalogActionData struct {
	Mode   EInputSourceMode // int32
	X      float32
	Y      float32
	Active bool
	_      [3]byte
}
type AnalogAction struct {
	ActionHandle     InputAnalogActionHandle
	AnalogActionData InputAnalogActionData
}
type DigitalAction struct {
	ActionHandle      InputDigitalActionHandle
	DigitalActionData InputDigitalActionData
}

type EControllerHapticType int32

const (
	EControllerHapticTypeOff   EControllerHapticType = 0
	EControllerHapticTypeTick  EControllerHapticType = 1
	EControllerHapticTypeClick EControllerHapticType = 2
)

type ESteamInputConfigurationEnableType int32

const (
	ESteamInputConfigurationEnableTypeNone        ESteamInputConfigurationEnableType = 0
	ESteamInputConfigurationEnableTypePlaystation ESteamInputConfigurationEnableType = 1
	ESteamInputConfigurationEnableTypeXbox        ESteamInputConfigurationEnableType = 2
	ESteamInputConfigurationEnableTypeGeneric     ESteamInputConfigurationEnableType = 4
	ESteamInputConfigurationEnableTypeSwitch      ESteamInputConfigurationEnableType = 8
)

type ESteamInputLEDFlag int32

const (
	ESteamInputLEDFlagSetColor           ESteamInputLEDFlag = 0
	ESteamInputLEDFlagRestoreUserDefault ESteamInputLEDFlag = 1
)

type ESteamInputGlyphStyle int32

const (
	ESteamInputGlyphStyleKnockout         ESteamInputGlyphStyle = 0
	ESteamInputGlyphStyleLight            ESteamInputGlyphStyle = 1
	ESteamInputGlyphStyleDark             ESteamInputGlyphStyle = 2
	ESteamInputGlyphStyleNeutralColorABXY ESteamInputGlyphStyle = 16
	ESteamInputGlyphStyleSolidABXY        ESteamInputGlyphStyle = 32
)

type ESteamInputActionEventType int32

const (
	ESteamInputActionEventTypeDigitalAction ESteamInputActionEventType = 0
	ESteamInputActionEventTypeAnalogAction  ESteamInputActionEventType = 1
)

type EControllerHapticLocation int32

const (
	EControllerHapticLocationLeft  EControllerHapticLocation = 1
	EControllerHapticLocationRight EControllerHapticLocation = 2
	EControllerHapticLocationBoth  EControllerHapticLocation = 3
)

type ScePadTriggerEffectParam struct{}

type ESteamInputGlyphSize int32

const (
	ESteamInputGlyphSizeSmall  ESteamInputGlyphSize = 0
	ESteamInputGlyphSizeMedium ESteamInputGlyphSize = 1
	ESteamInputGlyphSizeLarge  ESteamInputGlyphSize = 2
	ESteamInputGlyphSizeCount  ESteamInputGlyphSize = 3
)

type EInputActionOrigin int32

const (
	EInputActionOriginNone                              EInputActionOrigin = 0
	EInputActionOriginSteamControllerA                  EInputActionOrigin = 1
	EInputActionOriginSteamControllerB                  EInputActionOrigin = 2
	EInputActionOriginSteamControllerX                  EInputActionOrigin = 3
	EInputActionOriginSteamControllerY                  EInputActionOrigin = 4
	EInputActionOriginSteamControllerLeftBumper         EInputActionOrigin = 5
	EInputActionOriginSteamControllerRightBumper        EInputActionOrigin = 6
	EInputActionOriginSteamControllerLeftGrip           EInputActionOrigin = 7
	EInputActionOriginSteamControllerRightGrip          EInputActionOrigin = 8
	EInputActionOriginSteamControllerStart              EInputActionOrigin = 9
	EInputActionOriginSteamControllerBack               EInputActionOrigin = 10
	EInputActionOriginSteamControllerLeftPadTouch       EInputActionOrigin = 11
	EInputActionOriginSteamControllerLeftPadSwipe       EInputActionOrigin = 12
	EInputActionOriginSteamControllerLeftPadClick       EInputActionOrigin = 13
	EInputActionOriginSteamControllerLeftPadDPadNorth   EInputActionOrigin = 14
	EInputActionOriginSteamControllerLeftPadDPadSouth   EInputActionOrigin = 15
	EInputActionOriginSteamControllerLeftPadDPadWest    EInputActionOrigin = 16
	EInputActionOriginSteamControllerLeftPadDPadEast    EInputActionOrigin = 17
	EInputActionOriginSteamControllerRightPadTouch      EInputActionOrigin = 18
	EInputActionOriginSteamControllerRightPadSwipe      EInputActionOrigin = 19
	EInputActionOriginSteamControllerRightPadClick      EInputActionOrigin = 20
	EInputActionOriginSteamControllerRightPadDPadNorth  EInputActionOrigin = 21
	EInputActionOriginSteamControllerRightPadDPadSouth  EInputActionOrigin = 22
	EInputActionOriginSteamControllerRightPadDPadWest   EInputActionOrigin = 23
	EInputActionOriginSteamControllerRightPadDPadEast   EInputActionOrigin = 24
	EInputActionOriginSteamControllerLeftTriggerPull    EInputActionOrigin = 25
	EInputActionOriginSteamControllerLeftTriggerClick   EInputActionOrigin = 26
	EInputActionOriginSteamControllerRightTriggerPull   EInputActionOrigin = 27
	EInputActionOriginSteamControllerRightTriggerClick  EInputActionOrigin = 28
	EInputActionOriginSteamControllerLeftStickMove      EInputActionOrigin = 29
	EInputActionOriginSteamControllerLeftStickClick     EInputActionOrigin = 30
	EInputActionOriginSteamControllerLeftStickDPadNorth EInputActionOrigin = 31
	EInputActionOriginSteamControllerLeftStickDPadSouth EInputActionOrigin = 32
	EInputActionOriginSteamControllerLeftStickDPadWest  EInputActionOrigin = 33
	EInputActionOriginSteamControllerLeftStickDPadEast  EInputActionOrigin = 34
	EInputActionOriginSteamControllerGyroMove           EInputActionOrigin = 35
	EInputActionOriginSteamControllerGyroPitch          EInputActionOrigin = 36
	EInputActionOriginSteamControllerGyroYaw            EInputActionOrigin = 37
	EInputActionOriginSteamControllerGyroRoll           EInputActionOrigin = 38
	EInputActionOriginSteamControllerReserved0          EInputActionOrigin = 39
	EInputActionOriginSteamControllerReserved1          EInputActionOrigin = 40
	EInputActionOriginSteamControllerReserved2          EInputActionOrigin = 41
	EInputActionOriginSteamControllerReserved3          EInputActionOrigin = 42
	EInputActionOriginSteamControllerReserved4          EInputActionOrigin = 43
	EInputActionOriginSteamControllerReserved5          EInputActionOrigin = 44
	EInputActionOriginSteamControllerReserved6          EInputActionOrigin = 45
	EInputActionOriginSteamControllerReserved7          EInputActionOrigin = 46
	EInputActionOriginSteamControllerReserved8          EInputActionOrigin = 47
	EInputActionOriginSteamControllerReserved9          EInputActionOrigin = 48
	EInputActionOriginSteamControllerReserved10         EInputActionOrigin = 49
	EInputActionOriginPS4X                              EInputActionOrigin = 50
	EInputActionOriginPS4Circle                         EInputActionOrigin = 51
	EInputActionOriginPS4Triangle                       EInputActionOrigin = 52
	EInputActionOriginPS4Square                         EInputActionOrigin = 53
	EInputActionOriginPS4LeftBumper                     EInputActionOrigin = 54
	EInputActionOriginPS4RightBumper                    EInputActionOrigin = 55
	EInputActionOriginPS4Options                        EInputActionOrigin = 56
	EInputActionOriginPS4Share                          EInputActionOrigin = 57
	EInputActionOriginPS4LeftPadTouch                   EInputActionOrigin = 58
	EInputActionOriginPS4LeftPadSwipe                   EInputActionOrigin = 59
	EInputActionOriginPS4LeftPadClick                   EInputActionOrigin = 60
	EInputActionOriginPS4LeftPadDPadNorth               EInputActionOrigin = 61
	EInputActionOriginPS4LeftPadDPadSouth               EInputActionOrigin = 62
	EInputActionOriginPS4LeftPadDPadWest                EInputActionOrigin = 63
	EInputActionOriginPS4LeftPadDPadEast                EInputActionOrigin = 64
	EInputActionOriginPS4RightPadTouch                  EInputActionOrigin = 65
	EInputActionOriginPS4RightPadSwipe                  EInputActionOrigin = 66
	EInputActionOriginPS4RightPadClick                  EInputActionOrigin = 67
	EInputActionOriginPS4RightPadDPadNorth              EInputActionOrigin = 68
	EInputActionOriginPS4RightPadDPadSouth              EInputActionOrigin = 69
	EInputActionOriginPS4RightPadDPadWest               EInputActionOrigin = 70
	EInputActionOriginPS4RightPadDPadEast               EInputActionOrigin = 71
	EInputActionOriginPS4CenterPadTouch                 EInputActionOrigin = 72
	EInputActionOriginPS4CenterPadSwipe                 EInputActionOrigin = 73
	EInputActionOriginPS4CenterPadClick                 EInputActionOrigin = 74
	EInputActionOriginPS4CenterPadDPadNorth             EInputActionOrigin = 75
	EInputActionOriginPS4CenterPadDPadSouth             EInputActionOrigin = 76
	EInputActionOriginPS4CenterPadDPadWest              EInputActionOrigin = 77
	EInputActionOriginPS4CenterPadDPadEast              EInputActionOrigin = 78
	EInputActionOriginPS4LeftTriggerPull                EInputActionOrigin = 79
	EInputActionOriginPS4LeftTriggerClick               EInputActionOrigin = 80
	EInputActionOriginPS4RightTriggerPull               EInputActionOrigin = 81
	EInputActionOriginPS4RightTriggerClick              EInputActionOrigin = 82
	EInputActionOriginPS4LeftStickMove                  EInputActionOrigin = 83
	EInputActionOriginPS4LeftStickClick                 EInputActionOrigin = 84
	EInputActionOriginPS4LeftStickDPadNorth             EInputActionOrigin = 85
	EInputActionOriginPS4LeftStickDPadSouth             EInputActionOrigin = 86
	EInputActionOriginPS4LeftStickDPadWest              EInputActionOrigin = 87
	EInputActionOriginPS4LeftStickDPadEast              EInputActionOrigin = 88
	EInputActionOriginPS4RightStickMove                 EInputActionOrigin = 89
	EInputActionOriginPS4RightStickClick                EInputActionOrigin = 90
	EInputActionOriginPS4RightStickDPadNorth            EInputActionOrigin = 91
	EInputActionOriginPS4RightStickDPadSouth            EInputActionOrigin = 92
	EInputActionOriginPS4RightStickDPadWest             EInputActionOrigin = 93
	EInputActionOriginPS4RightStickDPadEast             EInputActionOrigin = 94
	EInputActionOriginPS4DPadNorth                      EInputActionOrigin = 95
	EInputActionOriginPS4DPadSouth                      EInputActionOrigin = 96
	EInputActionOriginPS4DPadWest                       EInputActionOrigin = 97
	EInputActionOriginPS4DPadEast                       EInputActionOrigin = 98
	EInputActionOriginPS4GyroMove                       EInputActionOrigin = 99
	EInputActionOriginPS4GyroPitch                      EInputActionOrigin = 100
	EInputActionOriginPS4GyroYaw                        EInputActionOrigin = 101
	EInputActionOriginPS4GyroRoll                       EInputActionOrigin = 102
	EInputActionOriginPS4DPadMove                       EInputActionOrigin = 103
	EInputActionOriginPS4Reserved1                      EInputActionOrigin = 104
	EInputActionOriginPS4Reserved2                      EInputActionOrigin = 105
	EInputActionOriginPS4Reserved3                      EInputActionOrigin = 106
	EInputActionOriginPS4Reserved4                      EInputActionOrigin = 107
	EInputActionOriginPS4Reserved5                      EInputActionOrigin = 108
	EInputActionOriginPS4Reserved6                      EInputActionOrigin = 109
	EInputActionOriginPS4Reserved7                      EInputActionOrigin = 110
	EInputActionOriginPS4Reserved8                      EInputActionOrigin = 111
	EInputActionOriginPS4Reserved9                      EInputActionOrigin = 112
	EInputActionOriginPS4Reserved10                     EInputActionOrigin = 113
	EInputActionOriginXBoxOneA                          EInputActionOrigin = 114
	EInputActionOriginXBoxOneB                          EInputActionOrigin = 115
	EInputActionOriginXBoxOneX                          EInputActionOrigin = 116
	EInputActionOriginXBoxOneY                          EInputActionOrigin = 117
	EInputActionOriginXBoxOneLeftBumper                 EInputActionOrigin = 118
	EInputActionOriginXBoxOneRightBumper                EInputActionOrigin = 119
	EInputActionOriginXBoxOneMenu                       EInputActionOrigin = 120
	EInputActionOriginXBoxOneView                       EInputActionOrigin = 121
	EInputActionOriginXBoxOneLeftTriggerPull            EInputActionOrigin = 122
	EInputActionOriginXBoxOneLeftTriggerClick           EInputActionOrigin = 123
	EInputActionOriginXBoxOneRightTriggerPull           EInputActionOrigin = 124
	EInputActionOriginXBoxOneRightTriggerClick          EInputActionOrigin = 125
	EInputActionOriginXBoxOneLeftStickMove              EInputActionOrigin = 126
	EInputActionOriginXBoxOneLeftStickClick             EInputActionOrigin = 127
	EInputActionOriginXBoxOneLeftStickDPadNorth         EInputActionOrigin = 128
	EInputActionOriginXBoxOneLeftStickDPadSouth         EInputActionOrigin = 129
	EInputActionOriginXBoxOneLeftStickDPadWest          EInputActionOrigin = 130
	EInputActionOriginXBoxOneLeftStickDPadEast          EInputActionOrigin = 131
	EInputActionOriginXBoxOneRightStickMove             EInputActionOrigin = 132
	EInputActionOriginXBoxOneRightStickClick            EInputActionOrigin = 133
	EInputActionOriginXBoxOneRightStickDPadNorth        EInputActionOrigin = 134
	EInputActionOriginXBoxOneRightStickDPadSouth        EInputActionOrigin = 135
	EInputActionOriginXBoxOneRightStickDPadWest         EInputActionOrigin = 136
	EInputActionOriginXBoxOneRightStickDPadEast         EInputActionOrigin = 137
	EInputActionOriginXBoxOneDPadNorth                  EInputActionOrigin = 138
	EInputActionOriginXBoxOneDPadSouth                  EInputActionOrigin = 139
	EInputActionOriginXBoxOneDPadWest                   EInputActionOrigin = 140
	EInputActionOriginXBoxOneDPadEast                   EInputActionOrigin = 141
	EInputActionOriginXBoxOneDPadMove                   EInputActionOrigin = 142
	EInputActionOriginXBoxOneLeftGripLower              EInputActionOrigin = 143
	EInputActionOriginXBoxOneLeftGripUpper              EInputActionOrigin = 144
	EInputActionOriginXBoxOneRightGripLower             EInputActionOrigin = 145
	EInputActionOriginXBoxOneRightGripUpper             EInputActionOrigin = 146
	EInputActionOriginXBoxOneShare                      EInputActionOrigin = 147
	EInputActionOriginXBoxOneReserved6                  EInputActionOrigin = 148
	EInputActionOriginXBoxOneReserved7                  EInputActionOrigin = 149
	EInputActionOriginXBoxOneReserved8                  EInputActionOrigin = 150
	EInputActionOriginXBoxOneReserved9                  EInputActionOrigin = 151
	EInputActionOriginXBoxOneReserved10                 EInputActionOrigin = 152
	EInputActionOriginXBox360A                          EInputActionOrigin = 153
	EInputActionOriginXBox360B                          EInputActionOrigin = 154
	EInputActionOriginXBox360X                          EInputActionOrigin = 155
	EInputActionOriginXBox360Y                          EInputActionOrigin = 156
	EInputActionOriginXBox360LeftBumper                 EInputActionOrigin = 157
	EInputActionOriginXBox360RightBumper                EInputActionOrigin = 158
	EInputActionOriginXBox360Start                      EInputActionOrigin = 159
	EInputActionOriginXBox360Back                       EInputActionOrigin = 160
	EInputActionOriginXBox360LeftTriggerPull            EInputActionOrigin = 161
	EInputActionOriginXBox360LeftTriggerClick           EInputActionOrigin = 162
	EInputActionOriginXBox360RightTriggerPull           EInputActionOrigin = 163
	EInputActionOriginXBox360RightTriggerClick          EInputActionOrigin = 164
	EInputActionOriginXBox360LeftStickMove              EInputActionOrigin = 165
	EInputActionOriginXBox360LeftStickClick             EInputActionOrigin = 166
	EInputActionOriginXBox360LeftStickDPadNorth         EInputActionOrigin = 167
	EInputActionOriginXBox360LeftStickDPadSouth         EInputActionOrigin = 168
	EInputActionOriginXBox360LeftStickDPadWest          EInputActionOrigin = 169
	EInputActionOriginXBox360LeftStickDPadEast          EInputActionOrigin = 170
	EInputActionOriginXBox360RightStickMove             EInputActionOrigin = 171
	EInputActionOriginXBox360RightStickClick            EInputActionOrigin = 172
	EInputActionOriginXBox360RightStickDPadNorth        EInputActionOrigin = 173
	EInputActionOriginXBox360RightStickDPadSouth        EInputActionOrigin = 174
	EInputActionOriginXBox360RightStickDPadWest         EInputActionOrigin = 175
	EInputActionOriginXBox360RightStickDPadEast         EInputActionOrigin = 176
	EInputActionOriginXBox360DPadNorth                  EInputActionOrigin = 177
	EInputActionOriginXBox360DPadSouth                  EInputActionOrigin = 178
	EInputActionOriginXBox360DPadWest                   EInputActionOrigin = 179
	EInputActionOriginXBox360DPadEast                   EInputActionOrigin = 180
	EInputActionOriginXBox360DPadMove                   EInputActionOrigin = 181
	EInputActionOriginXBox360Reserved1                  EInputActionOrigin = 182
	EInputActionOriginXBox360Reserved2                  EInputActionOrigin = 183
	EInputActionOriginXBox360Reserved3                  EInputActionOrigin = 184
	EInputActionOriginXBox360Reserved4                  EInputActionOrigin = 185
	EInputActionOriginXBox360Reserved5                  EInputActionOrigin = 186
	EInputActionOriginXBox360Reserved6                  EInputActionOrigin = 187
	EInputActionOriginXBox360Reserved7                  EInputActionOrigin = 188
	EInputActionOriginXBox360Reserved8                  EInputActionOrigin = 189
	EInputActionOriginXBox360Reserved9                  EInputActionOrigin = 190
	EInputActionOriginXBox360Reserved10                 EInputActionOrigin = 191
	EInputActionOriginSwitchA                           EInputActionOrigin = 192
	EInputActionOriginSwitchB                           EInputActionOrigin = 193
	EInputActionOriginSwitchX                           EInputActionOrigin = 194
	EInputActionOriginSwitchY                           EInputActionOrigin = 195
	EInputActionOriginSwitchLeftBumper                  EInputActionOrigin = 196
	EInputActionOriginSwitchRightBumper                 EInputActionOrigin = 197
	EInputActionOriginSwitchPlus                        EInputActionOrigin = 198
	EInputActionOriginSwitchMinus                       EInputActionOrigin = 199
	EInputActionOriginSwitchCapture                     EInputActionOrigin = 200
	EInputActionOriginSwitchLeftTriggerPull             EInputActionOrigin = 201
	EInputActionOriginSwitchLeftTriggerClick            EInputActionOrigin = 202
	EInputActionOriginSwitchRightTriggerPull            EInputActionOrigin = 203
	EInputActionOriginSwitchRightTriggerClick           EInputActionOrigin = 204
	EInputActionOriginSwitchLeftStickMove               EInputActionOrigin = 205
	EInputActionOriginSwitchLeftStickClick              EInputActionOrigin = 206
	EInputActionOriginSwitchLeftStickDPadNorth          EInputActionOrigin = 207
	EInputActionOriginSwitchLeftStickDPadSouth          EInputActionOrigin = 208
	EInputActionOriginSwitchLeftStickDPadWest           EInputActionOrigin = 209
	EInputActionOriginSwitchLeftStickDPadEast           EInputActionOrigin = 210
	EInputActionOriginSwitchRightStickMove              EInputActionOrigin = 211
	EInputActionOriginSwitchRightStickClick             EInputActionOrigin = 212
	EInputActionOriginSwitchRightStickDPadNorth         EInputActionOrigin = 213
	EInputActionOriginSwitchRightStickDPadSouth         EInputActionOrigin = 214
	EInputActionOriginSwitchRightStickDPadWest          EInputActionOrigin = 215
	EInputActionOriginSwitchRightStickDPadEast          EInputActionOrigin = 216
	EInputActionOriginSwitchDPadNorth                   EInputActionOrigin = 217
	EInputActionOriginSwitchDPadSouth                   EInputActionOrigin = 218
	EInputActionOriginSwitchDPadWest                    EInputActionOrigin = 219
	EInputActionOriginSwitchDPadEast                    EInputActionOrigin = 220
	EInputActionOriginSwitchProGyroMove                 EInputActionOrigin = 221
	EInputActionOriginSwitchProGyroPitch                EInputActionOrigin = 222
	EInputActionOriginSwitchProGyroYaw                  EInputActionOrigin = 223
	EInputActionOriginSwitchProGyroRoll                 EInputActionOrigin = 224
	EInputActionOriginSwitchDPadMove                    EInputActionOrigin = 225
	EInputActionOriginSwitchReserved1                   EInputActionOrigin = 226
	EInputActionOriginSwitchReserved2                   EInputActionOrigin = 227
	EInputActionOriginSwitchReserved3                   EInputActionOrigin = 228
	EInputActionOriginSwitchReserved4                   EInputActionOrigin = 229
	EInputActionOriginSwitchReserved5                   EInputActionOrigin = 230
	EInputActionOriginSwitchReserved6                   EInputActionOrigin = 231
	EInputActionOriginSwitchReserved7                   EInputActionOrigin = 232
	EInputActionOriginSwitchReserved8                   EInputActionOrigin = 233
	EInputActionOriginSwitchReserved9                   EInputActionOrigin = 234
	EInputActionOriginSwitchReserved10                  EInputActionOrigin = 235
	EInputActionOriginSwitchRightGyroMove               EInputActionOrigin = 236
	EInputActionOriginSwitchRightGyroPitch              EInputActionOrigin = 237
	EInputActionOriginSwitchRightGyroYaw                EInputActionOrigin = 238
	EInputActionOriginSwitchRightGyroRoll               EInputActionOrigin = 239
	EInputActionOriginSwitchLeftGyroMove                EInputActionOrigin = 240
	EInputActionOriginSwitchLeftGyroPitch               EInputActionOrigin = 241
	EInputActionOriginSwitchLeftGyroYaw                 EInputActionOrigin = 242
	EInputActionOriginSwitchLeftGyroRoll                EInputActionOrigin = 243
	EInputActionOriginSwitchLeftGripLower               EInputActionOrigin = 244
	EInputActionOriginSwitchLeftGripUpper               EInputActionOrigin = 245
	EInputActionOriginSwitchRightGripLower              EInputActionOrigin = 246
	EInputActionOriginSwitchRightGripUpper              EInputActionOrigin = 247
	EInputActionOriginSwitchJoyConButtonN               EInputActionOrigin = 248
	EInputActionOriginSwitchJoyConButtonE               EInputActionOrigin = 249
	EInputActionOriginSwitchJoyConButtonS               EInputActionOrigin = 250
	EInputActionOriginSwitchJoyConButtonW               EInputActionOrigin = 251
	EInputActionOriginSwitchReserved15                  EInputActionOrigin = 252
	EInputActionOriginSwitchReserved16                  EInputActionOrigin = 253
	EInputActionOriginSwitchReserved17                  EInputActionOrigin = 254
	EInputActionOriginSwitchReserved18                  EInputActionOrigin = 255
	EInputActionOriginSwitchReserved19                  EInputActionOrigin = 256
	EInputActionOriginSwitchReserved20                  EInputActionOrigin = 257
	EInputActionOriginPS5X                              EInputActionOrigin = 258
	EInputActionOriginPS5Circle                         EInputActionOrigin = 259
	EInputActionOriginPS5Triangle                       EInputActionOrigin = 260
	EInputActionOriginPS5Square                         EInputActionOrigin = 261
	EInputActionOriginPS5LeftBumper                     EInputActionOrigin = 262
	EInputActionOriginPS5RightBumper                    EInputActionOrigin = 263
	EInputActionOriginPS5Option                         EInputActionOrigin = 264
	EInputActionOriginPS5Create                         EInputActionOrigin = 265
	EInputActionOriginPS5Mute                           EInputActionOrigin = 266
	EInputActionOriginPS5LeftPadTouch                   EInputActionOrigin = 267
	EInputActionOriginPS5LeftPadSwipe                   EInputActionOrigin = 268
	EInputActionOriginPS5LeftPadClick                   EInputActionOrigin = 269
	EInputActionOriginPS5LeftPadDPadNorth               EInputActionOrigin = 270
	EInputActionOriginPS5LeftPadDPadSouth               EInputActionOrigin = 271
	EInputActionOriginPS5LeftPadDPadWest                EInputActionOrigin = 272
	EInputActionOriginPS5LeftPadDPadEast                EInputActionOrigin = 273
	EInputActionOriginPS5RightPadTouch                  EInputActionOrigin = 274
	EInputActionOriginPS5RightPadSwipe                  EInputActionOrigin = 275
	EInputActionOriginPS5RightPadClick                  EInputActionOrigin = 276
	EInputActionOriginPS5RightPadDPadNorth              EInputActionOrigin = 277
	EInputActionOriginPS5RightPadDPadSouth              EInputActionOrigin = 278
	EInputActionOriginPS5RightPadDPadWest               EInputActionOrigin = 279
	EInputActionOriginPS5RightPadDPadEast               EInputActionOrigin = 280
	EInputActionOriginPS5CenterPadTouch                 EInputActionOrigin = 281
	EInputActionOriginPS5CenterPadSwipe                 EInputActionOrigin = 282
	EInputActionOriginPS5CenterPadClick                 EInputActionOrigin = 283
	EInputActionOriginPS5CenterPadDPadNorth             EInputActionOrigin = 284
	EInputActionOriginPS5CenterPadDPadSouth             EInputActionOrigin = 285
	EInputActionOriginPS5CenterPadDPadWest              EInputActionOrigin = 286
	EInputActionOriginPS5CenterPadDPadEast              EInputActionOrigin = 287
	EInputActionOriginPS5LeftTriggerPull                EInputActionOrigin = 288
	EInputActionOriginPS5LeftTriggerClick               EInputActionOrigin = 289
	EInputActionOriginPS5RightTriggerPull               EInputActionOrigin = 290
	EInputActionOriginPS5RightTriggerClick              EInputActionOrigin = 291
	EInputActionOriginPS5LeftStickMove                  EInputActionOrigin = 292
	EInputActionOriginPS5LeftStickClick                 EInputActionOrigin = 293
	EInputActionOriginPS5LeftStickDPadNorth             EInputActionOrigin = 294
	EInputActionOriginPS5LeftStickDPadSouth             EInputActionOrigin = 295
	EInputActionOriginPS5LeftStickDPadWest              EInputActionOrigin = 296
	EInputActionOriginPS5LeftStickDPadEast              EInputActionOrigin = 297
	EInputActionOriginPS5RightStickMove                 EInputActionOrigin = 298
	EInputActionOriginPS5RightStickClick                EInputActionOrigin = 299
	EInputActionOriginPS5RightStickDPadNorth            EInputActionOrigin = 300
	EInputActionOriginPS5RightStickDPadSouth            EInputActionOrigin = 301
	EInputActionOriginPS5RightStickDPadWest             EInputActionOrigin = 302
	EInputActionOriginPS5RightStickDPadEast             EInputActionOrigin = 303
	EInputActionOriginPS5DPadNorth                      EInputActionOrigin = 304
	EInputActionOriginPS5DPadSouth                      EInputActionOrigin = 305
	EInputActionOriginPS5DPadWest                       EInputActionOrigin = 306
	EInputActionOriginPS5DPadEast                       EInputActionOrigin = 307
	EInputActionOriginPS5GyroMove                       EInputActionOrigin = 308
	EInputActionOriginPS5GyroPitch                      EInputActionOrigin = 309
	EInputActionOriginPS5GyroYaw                        EInputActionOrigin = 310
	EInputActionOriginPS5GyroRoll                       EInputActionOrigin = 311
	EInputActionOriginPS5DPadMove                       EInputActionOrigin = 312
	EInputActionOriginPS5LeftGrip                       EInputActionOrigin = 313
	EInputActionOriginPS5RightGrip                      EInputActionOrigin = 314
	EInputActionOriginPS5LeftFn                         EInputActionOrigin = 315
	EInputActionOriginPS5RightFn                        EInputActionOrigin = 316
	EInputActionOriginPS5Reserved5                      EInputActionOrigin = 317
	EInputActionOriginPS5Reserved6                      EInputActionOrigin = 318
	EInputActionOriginPS5Reserved7                      EInputActionOrigin = 319
	EInputActionOriginPS5Reserved8                      EInputActionOrigin = 320
	EInputActionOriginPS5Reserved9                      EInputActionOrigin = 321
	EInputActionOriginPS5Reserved10                     EInputActionOrigin = 322
	EInputActionOriginPS5Reserved11                     EInputActionOrigin = 323
	EInputActionOriginPS5Reserved12                     EInputActionOrigin = 324
	EInputActionOriginPS5Reserved13                     EInputActionOrigin = 325
	EInputActionOriginPS5Reserved14                     EInputActionOrigin = 326
	EInputActionOriginPS5Reserved15                     EInputActionOrigin = 327
	EInputActionOriginPS5Reserved16                     EInputActionOrigin = 328
	EInputActionOriginPS5Reserved17                     EInputActionOrigin = 329
	EInputActionOriginPS5Reserved18                     EInputActionOrigin = 330
	EInputActionOriginPS5Reserved19                     EInputActionOrigin = 331
	EInputActionOriginPS5Reserved20                     EInputActionOrigin = 332
	EInputActionOriginSteamDeckA                        EInputActionOrigin = 333
	EInputActionOriginSteamDeckB                        EInputActionOrigin = 334
	EInputActionOriginSteamDeckX                        EInputActionOrigin = 335
	EInputActionOriginSteamDeckY                        EInputActionOrigin = 336
	EInputActionOriginSteamDeckL1                       EInputActionOrigin = 337
	EInputActionOriginSteamDeckR1                       EInputActionOrigin = 338
	EInputActionOriginSteamDeckMenu                     EInputActionOrigin = 339
	EInputActionOriginSteamDeckView                     EInputActionOrigin = 340
	EInputActionOriginSteamDeckLeftPadTouch             EInputActionOrigin = 341
	EInputActionOriginSteamDeckLeftPadSwipe             EInputActionOrigin = 342
	EInputActionOriginSteamDeckLeftPadClick             EInputActionOrigin = 343
	EInputActionOriginSteamDeckLeftPadDPadNorth         EInputActionOrigin = 344
	EInputActionOriginSteamDeckLeftPadDPadSouth         EInputActionOrigin = 345
	EInputActionOriginSteamDeckLeftPadDPadWest          EInputActionOrigin = 346
	EInputActionOriginSteamDeckLeftPadDPadEast          EInputActionOrigin = 347
	EInputActionOriginSteamDeckRightPadTouch            EInputActionOrigin = 348
	EInputActionOriginSteamDeckRightPadSwipe            EInputActionOrigin = 349
	EInputActionOriginSteamDeckRightPadClick            EInputActionOrigin = 350
	EInputActionOriginSteamDeckRightPadDPadNorth        EInputActionOrigin = 351
	EInputActionOriginSteamDeckRightPadDPadSouth        EInputActionOrigin = 352
	EInputActionOriginSteamDeckRightPadDPadWest         EInputActionOrigin = 353
	EInputActionOriginSteamDeckRightPadDPadEast         EInputActionOrigin = 354
	EInputActionOriginSteamDeckL2SoftPull               EInputActionOrigin = 355
	EInputActionOriginSteamDeckL2                       EInputActionOrigin = 356
	EInputActionOriginSteamDeckR2SoftPull               EInputActionOrigin = 357
	EInputActionOriginSteamDeckR2                       EInputActionOrigin = 358
	EInputActionOriginSteamDeckLeftStickMove            EInputActionOrigin = 359
	EInputActionOriginSteamDeckL3                       EInputActionOrigin = 360
	EInputActionOriginSteamDeckLeftStickDPadNorth       EInputActionOrigin = 361
	EInputActionOriginSteamDeckLeftStickDPadSouth       EInputActionOrigin = 362
	EInputActionOriginSteamDeckLeftStickDPadWest        EInputActionOrigin = 363
	EInputActionOriginSteamDeckLeftStickDPadEast        EInputActionOrigin = 364
	EInputActionOriginSteamDeckLeftStickTouch           EInputActionOrigin = 365
	EInputActionOriginSteamDeckRightStickMove           EInputActionOrigin = 366
	EInputActionOriginSteamDeckR3                       EInputActionOrigin = 367
	EInputActionOriginSteamDeckRightStickDPadNorth      EInputActionOrigin = 368
	EInputActionOriginSteamDeckRightStickDPadSouth      EInputActionOrigin = 369
	EInputActionOriginSteamDeckRightStickDPadWest       EInputActionOrigin = 370
	EInputActionOriginSteamDeckRightStickDPadEast       EInputActionOrigin = 371
	EInputActionOriginSteamDeckRightStickTouch          EInputActionOrigin = 372
	EInputActionOriginSteamDeckL4                       EInputActionOrigin = 373
	EInputActionOriginSteamDeckR4                       EInputActionOrigin = 374
	EInputActionOriginSteamDeckL5                       EInputActionOrigin = 375
	EInputActionOriginSteamDeckR5                       EInputActionOrigin = 376
	EInputActionOriginSteamDeckDPadMove                 EInputActionOrigin = 377
	EInputActionOriginSteamDeckDPadNorth                EInputActionOrigin = 378
	EInputActionOriginSteamDeckDPadSouth                EInputActionOrigin = 379
	EInputActionOriginSteamDeckDPadWest                 EInputActionOrigin = 380
	EInputActionOriginSteamDeckDPadEast                 EInputActionOrigin = 381
	EInputActionOriginSteamDeckGyroMove                 EInputActionOrigin = 382
	EInputActionOriginSteamDeckGyroPitch                EInputActionOrigin = 383
	EInputActionOriginSteamDeckGyroYaw                  EInputActionOrigin = 384
	EInputActionOriginSteamDeckGyroRoll                 EInputActionOrigin = 385
	EInputActionOriginSteamDeckReserved1                EInputActionOrigin = 386
	EInputActionOriginSteamDeckReserved2                EInputActionOrigin = 387
	EInputActionOriginSteamDeckReserved3                EInputActionOrigin = 388
	EInputActionOriginSteamDeckReserved4                EInputActionOrigin = 389
	EInputActionOriginSteamDeckReserved5                EInputActionOrigin = 390
	EInputActionOriginSteamDeckReserved6                EInputActionOrigin = 391
	EInputActionOriginSteamDeckReserved7                EInputActionOrigin = 392
	EInputActionOriginSteamDeckReserved8                EInputActionOrigin = 393
	EInputActionOriginSteamDeckReserved9                EInputActionOrigin = 394
	EInputActionOriginSteamDeckReserved10               EInputActionOrigin = 395
	EInputActionOriginSteamDeckReserved11               EInputActionOrigin = 396
	EInputActionOriginSteamDeckReserved12               EInputActionOrigin = 397
	EInputActionOriginSteamDeckReserved13               EInputActionOrigin = 398
	EInputActionOriginSteamDeckReserved14               EInputActionOrigin = 399
	EInputActionOriginSteamDeckReserved15               EInputActionOrigin = 400
	EInputActionOriginSteamDeckReserved16               EInputActionOrigin = 401
	EInputActionOriginSteamDeckReserved17               EInputActionOrigin = 402
	EInputActionOriginSteamDeckReserved18               EInputActionOrigin = 403
	EInputActionOriginSteamDeckReserved19               EInputActionOrigin = 404
	EInputActionOriginSteamDeckReserved20               EInputActionOrigin = 405
	EInputActionOriginHoripadM1                         EInputActionOrigin = 406
	EInputActionOriginHoripadM2                         EInputActionOrigin = 407
	EInputActionOriginHoripadL4                         EInputActionOrigin = 408
	EInputActionOriginHoripadR4                         EInputActionOrigin = 409
	EInputActionOriginCount                             EInputActionOrigin = 410
	EInputActionOriginMaximumPossibleValue              EInputActionOrigin = 32767
)

type InputHandle uint64
type SteamInputActionEventCallbackPointer uintptr
type InputActionSetHandle uint64
type InputDigitalActionHandle uint64
type InputAnalogActionHandle uint64
type ESteamInputType int32

const (
	ESteamInputTypeUnknown              ESteamInputType = 0
	ESteamInputTypeSteamController      ESteamInputType = 1
	ESteamInputTypeXBox360Controller    ESteamInputType = 2
	ESteamInputTypeXBoxOneController    ESteamInputType = 3
	ESteamInputTypeGenericGamepad       ESteamInputType = 4
	ESteamInputTypePS4Controller        ESteamInputType = 5
	ESteamInputTypeAppleMFiController   ESteamInputType = 6
	ESteamInputTypeAndroidController    ESteamInputType = 7
	ESteamInputTypeSwitchJoyConPair     ESteamInputType = 8
	ESteamInputTypeSwitchJoyConSingle   ESteamInputType = 9
	ESteamInputTypeSwitchProController  ESteamInputType = 10
	ESteamInputTypeMobileTouch          ESteamInputType = 11
	ESteamInputTypePS3Controller        ESteamInputType = 12
	ESteamInputTypePS5Controller        ESteamInputType = 13
	ESteamInputTypeSteamDeckController  ESteamInputType = 14
	ESteamInputTypeCount                ESteamInputType = 15
	ESteamInputTypeMaximumPossibleValue ESteamInputType = 255
)

const (
	flatAPI_SteamInput                                       = "SteamAPI_SteamInput_v006"
	flatAPI_ISteamInput_Init                                 = "SteamAPI_ISteamInput_Init"
	flatAPI_ISteamInput_Shutdown                             = "SteamAPI_ISteamInput_Shutdown"
	flatAPI_ISteamInput_SetInputActionManifestFilePath       = "SteamAPI_ISteamInput_SetInputActionManifestFilePath"
	flatAPI_ISteamInput_RunFrame                             = "SteamAPI_ISteamInput_RunFrame"
	flatAPI_ISteamInput_BWaitForData                         = "SteamAPI_ISteamInput_BWaitForData"
	flatAPI_ISteamInput_BNewDataAvailable                    = "SteamAPI_ISteamInput_BNewDataAvailable"
	flatAPI_ISteamInput_GetConnectedControllers              = "SteamAPI_ISteamInput_GetConnectedControllers"
	flatAPI_ISteamInput_EnableDeviceCallbacks                = "SteamAPI_ISteamInput_EnableDeviceCallbacks"
	flatAPI_ISteamInput_EnableActionEventCallbacks           = "SteamAPI_ISteamInput_EnableActionEventCallbacks"
	flatAPI_ISteamInput_GetActionSetHandle                   = "SteamAPI_ISteamInput_GetActionSetHandle"
	flatAPI_ISteamInput_ActivateActionSet                    = "SteamAPI_ISteamInput_ActivateActionSet"
	flatAPI_ISteamInput_GetCurrentActionSet                  = "SteamAPI_ISteamInput_GetCurrentActionSet"
	flatAPI_ISteamInput_ActivateActionSetLayer               = "SteamAPI_ISteamInput_ActivateActionSetLayer"
	flatAPI_ISteamInput_DeactivateActionSetLayer             = "SteamAPI_ISteamInput_DeactivateActionSetLayer"
	flatAPI_ISteamInput_DeactivateAllActionSetLayers         = "SteamAPI_ISteamInput_DeactivateAllActionSetLayers"
	flatAPI_ISteamInput_GetActiveActionSetLayers             = "SteamAPI_ISteamInput_GetActiveActionSetLayers"
	flatAPI_ISteamInput_GetDigitalActionHandle               = "SteamAPI_ISteamInput_GetDigitalActionHandle"
	flatAPI_ISteamInput_GetDigitalActionData                 = "SteamAPI_ISteamInput_GetDigitalActionData"
	flatAPI_ISteamInput_GetDigitalActionOrigins              = "SteamAPI_ISteamInput_GetDigitalActionOrigins"
	flatAPI_ISteamInput_GetStringForDigitalActionName        = "SteamAPI_ISteamInput_GetStringForDigitalActionName"
	flatAPI_ISteamInput_GetAnalogActionHandle                = "SteamAPI_ISteamInput_GetAnalogActionHandle"
	flatAPI_ISteamInput_GetAnalogActionData                  = "SteamAPI_ISteamInput_GetAnalogActionData"
	flatAPI_ISteamInput_GetAnalogActionOrigins               = "SteamAPI_ISteamInput_GetAnalogActionOrigins"
	flatAPI_ISteamInput_GetGlyphPNGForActionOrigin           = "SteamAPI_ISteamInput_GetGlyphPNGForActionOrigin"
	flatAPI_ISteamInput_GetGlyphSVGForActionOrigin           = "SteamAPI_ISteamInput_GetGlyphSVGForActionOrigin"
	flatAPI_ISteamInput_GetGlyphForActionOrigin_Legacy       = "SteamAPI_ISteamInput_GetGlyphForActionOrigin_Legacy"
	flatAPI_ISteamInput_GetStringForActionOrigin             = "SteamAPI_ISteamInput_GetStringForActionOrigin"
	flatAPI_ISteamInput_GetStringForAnalogActionName         = "SteamAPI_ISteamInput_GetStringForAnalogActionName"
	flatAPI_ISteamInput_StopAnalogActionMomentum             = "SteamAPI_ISteamInput_StopAnalogActionMomentum"
	flatAPI_ISteamInput_GetMotionData                        = "SteamAPI_ISteamInput_GetMotionData"
	flatAPI_ISteamInput_TriggerVibration                     = "SteamAPI_ISteamInput_TriggerVibration"
	flatAPI_ISteamInput_TriggerVibrationExtended             = "SteamAPI_ISteamInput_TriggerVibrationExtended"
	flatAPI_ISteamInput_TriggerSimpleHapticEvent             = "SteamAPI_ISteamInput_TriggerSimpleHapticEvent"
	flatAPI_ISteamInput_SetLEDColor                          = "SteamAPI_ISteamInput_SetLEDColor"
	flatAPI_ISteamInput_Legacy_TriggerHapticPulse            = "SteamAPI_ISteamInput_Legacy_TriggerHapticPulse"
	flatAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse    = "SteamAPI_ISteamInput_Legacy_TriggerRepeatedHapticPulse"
	flatAPI_ISteamInput_ShowBindingPanel                     = "SteamAPI_ISteamInput_ShowBindingPanel"
	flatAPI_ISteamInput_GetInputTypeForHandle                = "SteamAPI_ISteamInput_GetInputTypeForHandle"
	flatAPI_ISteamInput_GetControllerForGamepadIndex         = "SteamAPI_ISteamInput_GetControllerForGamepadIndex"
	flatAPI_ISteamInput_GetGamepadIndexForController         = "SteamAPI_ISteamInput_GetGamepadIndexForController"
	flatAPI_ISteamInput_GetStringForXboxOrigin               = "SteamAPI_ISteamInput_GetStringForXboxOrigin"
	flatAPI_ISteamInput_GetGlyphForXboxOrigin                = "SteamAPI_ISteamInput_GetGlyphForXboxOrigin"
	flatAPI_ISteamInput_GetActionOriginFromXboxOrigin        = "SteamAPI_ISteamInput_GetActionOriginFromXboxOrigin"
	flatAPI_ISteamInput_TranslateActionOrigin                = "SteamAPI_ISteamInput_TranslateActionOrigin"
	flatAPI_ISteamInput_GetDeviceBindingRevision             = "SteamAPI_ISteamInput_GetDeviceBindingRevision"
	flatAPI_ISteamInput_GetRemotePlaySessionID               = "SteamAPI_ISteamInput_GetRemotePlaySessionID"
	flatAPI_ISteamInput_GetSessionInputConfigurationSettings = "SteamAPI_ISteamInput_GetSessionInputConfigurationSettings"
	flatAPI_ISteamInput_SetDualSenseTriggerEffect            = "SteamAPI_ISteamInput_SetDualSenseTriggerEffect"
)

var (
	steamInput_init                                  func() uintptr
	iSteamInput_Init                                 func(steamInput uintptr, explicitlyCallRunFrame bool) bool
	iSteamInput_Shutdown                             func(steamInput uintptr) bool
	iSteamInput_SetInputActionManifestFilePath       func(steamInput uintptr, inputActionManifestAbsolutePath string) bool
	iSteamInput_RunFrame                             func(steamInput uintptr, reservedValue bool)
	iSteamInput_BWaitForData                         func(steamInput uintptr, waitForever bool, timeout uint32) bool
	iSteamInput_BNewDataAvailable                    func(steamInput uintptr) bool
	iSteamInput_GetConnectedControllers              func(steamInput uintptr, handlesOut *[STEAM_INPUT_MAX_COUNT]InputHandle) int32
	iSteamInput_EnableDeviceCallbacks                func(steamInput uintptr)
	iSteamInput_EnableActionEventCallbacks           func(steamInput uintptr, pCallback SteamInputActionEventCallbackPointer)
	iSteamInput_GetActionSetHandle                   func(steamInput uintptr, pszActionSetName string) InputActionSetHandle
	iSteamInput_ActivateActionSet                    func(steamInput uintptr, inputHandle InputHandle, actionSetHandle InputActionSetHandle)
	iSteamInput_GetCurrentActionSet                  func(steamInput uintptr, inputHandle InputHandle) InputActionSetHandle
	iSteamInput_ActivateActionSetLayer               func(steamInput uintptr, inputHandle InputHandle, actionSetLayerHandle InputActionSetHandle)
	iSteamInput_DeactivateActionSetLayer             func(steamInput uintptr, inputHandle InputHandle, actionSetLayerHandle InputActionSetHandle)
	iSteamInput_DeactivateAllActionSetLayers         func(steamInput uintptr, inputHandle InputHandle)
	iSteamInput_GetActiveActionSetLayers             func(steamInput uintptr, inputHandle InputHandle, handlesOut *[STEAM_INPUT_MAX_ACTIVE_LAYERS]InputActionSetHandle) int32
	iSteamInput_GetDigitalActionHandle               func(steamInput uintptr, actionName string) InputDigitalActionHandle
	iSteamInput_GetDigitalActionData                 func(steamInput uintptr, inputHandle InputHandle, digitalActionHandle InputDigitalActionHandle) uintptr
	iSteamInput_GetDigitalActionOrigins              func(steamInput uintptr, inputHandle InputHandle, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle, originsOut *[STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32
	iSteamInput_GetStringForDigitalActionName        func(steamInput uintptr, actionHandle InputDigitalActionHandle) string
	iSteamInput_GetAnalogActionHandle                func(steamInput uintptr, actionName string) InputAnalogActionHandle
	iSteamInput_GetAnalogActionData                  func(steamInput uintptr, inputHandle InputHandle, analogActionHandle InputAnalogActionHandle) uintptr
	iSteamInput_GetAnalogActionOrigins               func(steamInput uintptr, inputHandle InputHandle, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle, originsOut *[STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin) int32
	iSteamInput_GetGlyphPNGForActionOrigin           func(steamInput uintptr, eOrigin EInputActionOrigin, size ESteamInputGlyphSize, flags uint32) []byte
	iSteamInput_GetGlyphSVGForActionOrigin           func(steamInput uintptr, eOrigin EInputActionOrigin, flags uint32) []byte
	iSteamInput_GetGlyphForActionOrigin_Legacy       func(steamInput uintptr, eOrigin EInputActionOrigin) []byte
	iSteamInput_GetStringForActionOrigin             func(steamInput uintptr, eOrigin EInputActionOrigin) string
	iSteamInput_GetStringForAnalogActionName         func(steamInput uintptr, actionHandle InputAnalogActionHandle) string
	iSteamInput_StopAnalogActionMomentum             func(steamInput uintptr, inputHandle InputHandle, action InputAnalogActionHandle)
	iSteamInput_GetMotionData                        func(steamInput uintptr, inputHandle InputHandle) uintptr
	iSteamInput_TriggerVibration                     func(steamInput uintptr, inputHandle InputHandle, leftSpeed uint16, rightSpeed uint16)
	iSteamInput_TriggerVibrationExtended             func(steamInput uintptr, inputHandle InputHandle, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16)
	iSteamInput_TriggerSimpleHapticEvent             func(steamInput uintptr, inputHandle InputHandle, eHapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte)
	iSteamInput_SetLEDColor                          func(steamInput uintptr, inputHandle InputHandle, colorR uint8, colorG uint8, colorB uint8, flags uint)
	iSteamInput_Legacy_TriggerHapticPulse            func(steamInput uintptr, inputHandle InputHandle, targetPad ESteamControllerPad, durationMicroSec uint16)
	iSteamInput_Legacy_TriggerRepeatedHapticPulse    func(steamInput uintptr, inputHandle InputHandle, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint)
	iSteamInput_ShowBindingPanel                     func(steamInput uintptr, inputHandle InputHandle) bool
	iSteamInput_GetInputTypeForHandle                func(steamInput uintptr, inputHandle InputHandle) ESteamInputType
	iSteamInput_GetControllerForGamepadIndex         func(steamInput uintptr, index int32) InputHandle
	iSteamInput_GetGamepadIndexForController         func(steamInput uintptr, inputHandle InputHandle) int32
	iSteamInput_GetStringForXboxOrigin               func(steamInput uintptr, eOrigin EXboxOrigin) string
	iSteamInput_GetGlyphForXboxOrigin                func(steamInput uintptr, eOrigin EXboxOrigin) []byte
	iSteamInput_GetActionOriginFromXboxOrigin        func(steamInput uintptr, inputHandle InputHandle, eOrigin EXboxOrigin) EInputActionOrigin
	iSteamInput_TranslateActionOrigin                func(steamInput uintptr, destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin
	iSteamInput_GetDeviceBindingRevision             func(steamInput uintptr, inputHandle InputHandle, major *int32, minor *int32) bool
	iSteamInput_GetRemotePlaySessionID               func(steamInput uintptr, inputHandle InputHandle) uint32
	iSteamInput_GetSessionInputConfigurationSettings func(steamInput uintptr) uint16
	iSteamInput_SetDualSenseTriggerEffect            func(steamInput uintptr, inputHandle InputHandle, param *ScePadTriggerEffectParam)
)

var steamInput uintptr

func input() uintptr {
	if steamInput == 0 {
		steamInput = steamInput_init()
		return steamInput
	}
	return steamInput
}

func Init(explicitlyCallRunFrame bool) bool {
	return iSteamInput_Init(input(), explicitlyCallRunFrame)
}

func Shutdown() bool {
	return iSteamInput_Shutdown(input())
}

func SetInputActionManifestFilePath(inputActionManifestAbsolutePath string) bool {
	return iSteamInput_SetInputActionManifestFilePath(input(), inputActionManifestAbsolutePath)
}

func RunFrame(reservedValue bool) {
	iSteamInput_RunFrame(input(), reservedValue)
}

func BWaitForData(waitForever bool, timeout uint32) bool {
	return iSteamInput_BWaitForData(input(), waitForever, timeout)
}

func BNewDataAvailable() bool {
	return iSteamInput_BNewDataAvailable(input())
}

func GetConnectedControllers() []InputHandle {
	var handlesOut [STEAM_INPUT_MAX_COUNT]InputHandle
	result := iSteamInput_GetConnectedControllers(input(), &handlesOut)
	return handlesOut[:result]
}

func EnableDeviceCallbacks() {
	iSteamInput_EnableDeviceCallbacks(input())
}

func EnableActionEventCallbacks(pCallback SteamInputActionEventCallbackPointer) {
	iSteamInput_EnableActionEventCallbacks(input(), pCallback)
}

func GetActionSetHandle(pszActionSetName string) InputActionSetHandle {
	return iSteamInput_GetActionSetHandle(input(), pszActionSetName)
}

func ActivateActionSet(inputHandle InputHandle, actionSetHandle InputActionSetHandle) {
	iSteamInput_ActivateActionSet(input(), inputHandle, actionSetHandle)
}

func GetCurrentActionSet(inputHandle InputHandle) InputActionSetHandle {
	return iSteamInput_GetCurrentActionSet(input(), inputHandle)
}

func ActivateActionSetLayer(inputHandle InputHandle, actionSetLayerHandle InputActionSetHandle) {
	iSteamInput_ActivateActionSetLayer(input(), inputHandle, actionSetLayerHandle)
}

func DeactivateActionSetLayer(inputHandle InputHandle, actionSetLayerHandle InputActionSetHandle) {
	iSteamInput_DeactivateActionSetLayer(input(), inputHandle, actionSetLayerHandle)
}

func DeactivateAllActionSetLayers(inputHandle InputHandle) {
	iSteamInput_DeactivateAllActionSetLayers(input(), inputHandle)
}

func GetActiveActionSetLayers(inputHandle InputHandle) []InputActionSetHandle {
	var handlesOut [STEAM_INPUT_MAX_ACTIVE_LAYERS]InputActionSetHandle
	result := iSteamInput_GetActiveActionSetLayers(input(), inputHandle, &handlesOut)
	return handlesOut[:result]
}

func GetDigitalActionHandle(actionName string) InputDigitalActionHandle {
	return iSteamInput_GetDigitalActionHandle(input(), actionName)
}

func GetDigitalActionData(inputHandle InputHandle, digitalActionHandle InputDigitalActionHandle) InputDigitalActionData {
	return *common.UintptrToStruct[InputDigitalActionData](iSteamInput_GetDigitalActionData(input(), inputHandle, digitalActionHandle))
}

func GetDigitalActionOrigins(inputHandle InputHandle, actionSetHandle InputActionSetHandle, digitalActionHandle InputDigitalActionHandle) []EInputActionOrigin {
	var originsOut [STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin
	result := iSteamInput_GetDigitalActionOrigins(input(), inputHandle, actionSetHandle, digitalActionHandle, &originsOut)
	return originsOut[:result]
}

func GetStringForDigitalActionName(actionHandle InputDigitalActionHandle) string {
	return iSteamInput_GetStringForDigitalActionName(input(), actionHandle)
}

func GetAnalogActionHandle(actionName string) InputAnalogActionHandle {
	return iSteamInput_GetAnalogActionHandle(input(), actionName)
}

func GetAnalogActionData(inputHandle InputHandle, analogActionHandle InputAnalogActionHandle) InputAnalogActionData {
	return *common.UintptrToStruct[InputAnalogActionData](iSteamInput_GetAnalogActionData(input(), inputHandle, analogActionHandle))
}

func GetAnalogActionOrigins(inputHandle InputHandle, actionSetHandle InputActionSetHandle, analogActionHandle InputAnalogActionHandle) []EInputActionOrigin {
	var originsOut [STEAM_INPUT_MAX_ORIGINS]EInputActionOrigin
	result := iSteamInput_GetAnalogActionOrigins(input(), inputHandle, actionSetHandle, analogActionHandle, &originsOut)
	return originsOut[:result]
}

func GetGlyphPNGForActionOrigin(origin EInputActionOrigin, size ESteamInputGlyphSize, flags uint32) []byte {
	return iSteamInput_GetGlyphPNGForActionOrigin(input(), origin, size, flags)
}

func GetGlyphSVGForActionOrigin(origin EInputActionOrigin, flags uint32) []byte {
	return iSteamInput_GetGlyphSVGForActionOrigin(input(), origin, flags)
}

func GetGlyphForActionOrigin_Legacy(origin EInputActionOrigin) []byte {
	return iSteamInput_GetGlyphForActionOrigin_Legacy(input(), origin)
}

func GetStringForActionOrigin(origin EInputActionOrigin) string {
	return iSteamInput_GetStringForActionOrigin(input(), origin)
}

func GetStringForAnalogActionName(actionHandle InputAnalogActionHandle) string {
	return iSteamInput_GetStringForAnalogActionName(input(), actionHandle)
}

func StopAnalogActionMomentum(inputHandle InputHandle, action InputAnalogActionHandle) {
	iSteamInput_StopAnalogActionMomentum(input(), inputHandle, action)
}

func GetMotionData(inputHandle InputHandle) InputMotionData {
	return *common.UintptrToStruct[InputMotionData](iSteamInput_GetMotionData(input(), inputHandle))
}

func TriggerVibration(inputHandle InputHandle, leftSpeed uint16, rightSpeed uint16) {
	iSteamInput_TriggerVibration(input(), inputHandle, leftSpeed, rightSpeed)
}

func TriggerVibrationExtended(inputHandle InputHandle, leftSpeed uint16, rightSpeed uint16, leftTriggerSpeed uint16, rightTriggerSpeed uint16) {
	iSteamInput_TriggerVibrationExtended(input(), inputHandle, leftSpeed, rightSpeed, leftTriggerSpeed, rightTriggerSpeed)
}

func TriggerSimpleHapticEvent(inputHandle InputHandle, hapticLocation EControllerHapticLocation, intensity uint8, gainDB byte, otherIntensity uint8, otherGainDB byte) {
	iSteamInput_TriggerSimpleHapticEvent(input(), inputHandle, hapticLocation, intensity, gainDB, otherIntensity, otherGainDB)
}

func SetLEDColor(inputHandle InputHandle, colorR uint8, colorG uint8, colorB uint8, flags uint) {
	iSteamInput_SetLEDColor(input(), inputHandle, colorR, colorG, colorB, flags)
}

func Legacy_TriggerHapticPulse(inputHandle InputHandle, targetPad ESteamControllerPad, durationMicroSec uint16) {
	iSteamInput_Legacy_TriggerHapticPulse(input(), inputHandle, targetPad, durationMicroSec)
}

func Legacy_TriggerRepeatedHapticPulse(inputHandle InputHandle, targetPad ESteamControllerPad, durationMicroSec uint16, offMicroSec uint16, repeat uint16, flags uint) {
	iSteamInput_Legacy_TriggerRepeatedHapticPulse(input(), inputHandle, targetPad, durationMicroSec, offMicroSec, repeat, flags)
}

func ShowBindingPanel(inputHandle InputHandle) bool {
	return iSteamInput_ShowBindingPanel(input(), inputHandle)
}

func GetInputTypeForHandle(inputHandle InputHandle) ESteamInputType {
	return iSteamInput_GetInputTypeForHandle(input(), inputHandle)
}

func GetControllerForGamepadIndex(index int32) InputHandle {
	return iSteamInput_GetControllerForGamepadIndex(input(), index)
}

func GetGamepadIndexForController(inputHandle InputHandle) int32 {
	return iSteamInput_GetGamepadIndexForController(input(), inputHandle)
}

func GetStringForXboxOrigin(origin EXboxOrigin) string {
	return iSteamInput_GetStringForXboxOrigin(input(), origin)
}

func GetGlyphForXboxOrigin(origin EXboxOrigin) []byte {
	return iSteamInput_GetGlyphForXboxOrigin(input(), origin)
}

func GetActionOriginFromXboxOrigin(inputHandle InputHandle, origin EXboxOrigin) EInputActionOrigin {
	return iSteamInput_GetActionOriginFromXboxOrigin(input(), inputHandle, origin)
}

func TranslateActionOrigin(destinationInputType ESteamInputType, sourceOrigin EInputActionOrigin) EInputActionOrigin {
	return iSteamInput_TranslateActionOrigin(input(), destinationInputType, sourceOrigin)
}

func GetDeviceBindingRevision(inputHandle InputHandle) (major int32, minor int32, foundController bool) {
	foundController = iSteamInput_GetDeviceBindingRevision(input(), inputHandle, &major, &minor)
	return major, minor, foundController
}

func GetRemotePlaySessionID(inputHandle InputHandle) uint32 {
	return iSteamInput_GetRemotePlaySessionID(input(), inputHandle)
}

func GetSessionInputConfigurationSettings() uint16 {
	return iSteamInput_GetSessionInputConfigurationSettings(input())
}

func SetDualSenseTriggerEffect(inputHandle InputHandle, param *ScePadTriggerEffectParam) {
	iSteamInput_SetDualSenseTriggerEffect(input(), inputHandle, param)
}
