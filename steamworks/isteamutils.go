package steamworks

const (
	flatAPI_SteamUtils                               = "SteamAPI_SteamUtils_v010"
	flatAPI_ISteamUtils_IsSteamRunningOnSteamDeck    = "SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck"
	flatAPI_ISteamUtils_ShowFloatingGamepadTextInput = "SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput"
)

type steamUtils uintptr

var (
	steamUitils_init             func() uintptr
	isSteamRunningOnSteamDeck    func(steamUtils uintptr) uintptr
	showFloatingGamepadTextInput func(steamUtils uintptr, keyboardMode uintptr, textFieldXPosition uintptr, textFieldYPosition uintptr, textFieldWidth uintptr, textFieldHeight uintptr) uintptr
)

func (s steamUtils) IsSteamRunningOnSteamDeck() bool {
	result := isSteamRunningOnSteamDeck(uintptr(s))
	return byte(result) != 0
}

func (s steamUtils) ShowFloatingGamepadTextInput(keyboardMode EFloatingGamepadTextInputMode, textFieldXPosition, textFieldYPosition, textFieldWidth, textFieldHeight int32) bool {
	result := showFloatingGamepadTextInput(uintptr(s), uintptr(keyboardMode), uintptr(textFieldXPosition), uintptr(textFieldYPosition), uintptr(textFieldWidth), uintptr(textFieldHeight))
	return byte(result) != 0
}
