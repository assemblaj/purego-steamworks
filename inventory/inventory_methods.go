package inventory

import (
	. "github.com/assemblaj/purego-steamworks"
)

const (
	SteamItemInstanceIDInvalid        SteamItemInstanceID        = ^SteamItemInstanceID(0)
	SteamInventoryResultInvalid       SteamInventoryResult       = -1
	SteamInventoryUpdateHandleInvalid SteamInventoryUpdateHandle = 0xffffffffffffffff
)

type SteamItemDef int32
type SteamInventoryUpdateHandle uint64
type SteamItemDetails struct {
	ItemId     SteamItemInstanceID
	Definition SteamItemDef
	Quantity   uint16
	Flags      uint16
}
type SteamItemInstanceID uint64

type SteamInventoryResult int32

type ESteamItemFlags int32

const (
	ESteamItemNoTrade  ESteamItemFlags = 1
	ESteamItemRemoved  ESteamItemFlags = 256
	ESteamItemConsumed ESteamItemFlags = 512
)

const (
	flatAPI_SteamInventory                                         = "SteamAPI_SteamInventory_v003"
	flatAPI_ISteamInventory_GetResultStatus                        = "SteamAPI_ISteamInventory_GetResultStatus"
	flatAPI_ISteamInventory_GetResultItems                         = "SteamAPI_ISteamInventory_GetResultItems"
	flatAPI_ISteamInventory_GetResultItemProperty                  = "SteamAPI_ISteamInventory_GetResultItemProperty"
	flatAPI_ISteamInventory_GetResultTimestamp                     = "SteamAPI_ISteamInventory_GetResultTimestamp"
	flatAPI_ISteamInventory_CheckResultSteamID                     = "SteamAPI_ISteamInventory_CheckResultSteamID"
	flatAPI_ISteamInventory_DestroyResult                          = "SteamAPI_ISteamInventory_DestroyResult"
	flatAPI_ISteamInventory_GetAllItems                            = "SteamAPI_ISteamInventory_GetAllItems"
	flatAPI_ISteamInventory_GetItemsByID                           = "SteamAPI_ISteamInventory_GetItemsByID"
	flatAPI_ISteamInventory_SerializeResult                        = "SteamAPI_ISteamInventory_SerializeResult"
	flatAPI_ISteamInventory_DeserializeResult                      = "SteamAPI_ISteamInventory_DeserializeResult"
	flatAPI_ISteamInventory_GenerateItems                          = "SteamAPI_ISteamInventory_GenerateItems"
	flatAPI_ISteamInventory_GrantPromoItems                        = "SteamAPI_ISteamInventory_GrantPromoItems"
	flatAPI_ISteamInventory_AddPromoItem                           = "SteamAPI_ISteamInventory_AddPromoItem"
	flatAPI_ISteamInventory_AddPromoItems                          = "SteamAPI_ISteamInventory_AddPromoItems"
	flatAPI_ISteamInventory_ConsumeItem                            = "SteamAPI_ISteamInventory_ConsumeItem"
	flatAPI_ISteamInventory_ExchangeItems                          = "SteamAPI_ISteamInventory_ExchangeItems"
	flatAPI_ISteamInventory_TransferItemQuantity                   = "SteamAPI_ISteamInventory_TransferItemQuantity"
	flatAPI_ISteamInventory_TriggerItemDrop                        = "SteamAPI_ISteamInventory_TriggerItemDrop"
	flatAPI_ISteamInventory_LoadItemDefinitions                    = "SteamAPI_ISteamInventory_LoadItemDefinitions"
	flatAPI_ISteamInventory_GetItemDefinitionIDs                   = "SteamAPI_ISteamInventory_GetItemDefinitionIDs"
	flatAPI_ISteamInventory_GetItemDefinitionProperty              = "SteamAPI_ISteamInventory_GetItemDefinitionProperty"
	flatAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs = "SteamAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs"
	flatAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs      = "SteamAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs"
	flatAPI_ISteamInventory_StartPurchase                          = "SteamAPI_ISteamInventory_StartPurchase"
	flatAPI_ISteamInventory_RequestPrices                          = "SteamAPI_ISteamInventory_RequestPrices"
	flatAPI_ISteamInventory_GetNumItemsWithPrices                  = "SteamAPI_ISteamInventory_GetNumItemsWithPrices"
	flatAPI_ISteamInventory_GetItemsWithPrices                     = "SteamAPI_ISteamInventory_GetItemsWithPrices"
	flatAPI_ISteamInventory_GetItemPrice                           = "SteamAPI_ISteamInventory_GetItemPrice"
	flatAPI_ISteamInventory_StartUpdateProperties                  = "SteamAPI_ISteamInventory_StartUpdateProperties"
	flatAPI_ISteamInventory_RemoveProperty                         = "SteamAPI_ISteamInventory_RemoveProperty"
	flatAPI_ISteamInventory_SetPropertyString                      = "SteamAPI_ISteamInventory_SetPropertyString"
	flatAPI_ISteamInventory_SetPropertyBool                        = "SteamAPI_ISteamInventory_SetPropertyBool"
	flatAPI_ISteamInventory_SetPropertyInt64                       = "SteamAPI_ISteamInventory_SetPropertyInt64"
	flatAPI_ISteamInventory_SetPropertyFloat                       = "SteamAPI_ISteamInventory_SetPropertyFloat"
	flatAPI_ISteamInventory_SubmitUpdateProperties                 = "SteamAPI_ISteamInventory_SubmitUpdateProperties"
	flatAPI_ISteamInventory_InspectItem                            = "SteamAPI_ISteamInventory_InspectItem"
)

var (
	steamInventory_init                                    func() uintptr
	iSteamInventory_GetResultStatus                        func(steamInventory uintptr, resultHandle SteamInventoryResult) EResult
	iSteamInventory_GetResultItems                         func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutItemsArray *[]SteamItemDetails, punOutItemsArraySize *uint32) bool
	iSteamInventory_GetResultItemProperty                  func(steamInventory uintptr, resultHandle SteamInventoryResult, unItemIndex uint32, pchPropertyName string, pchValueBuffer *[]byte, punValueBufferSizeOut *uint32) bool
	iSteamInventory_GetResultTimestamp                     func(steamInventory uintptr, resultHandle SteamInventoryResult) uint32
	iSteamInventory_CheckResultSteamID                     func(steamInventory uintptr, resultHandle SteamInventoryResult, steamIDExpected Uint64SteamID) bool
	iSteamInventory_DestroyResult                          func(steamInventory uintptr, resultHandle SteamInventoryResult)
	iSteamInventory_GetAllItems                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_GetItemsByID                           func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pInstanceIDs *[]SteamItemInstanceID, unCountInstanceIDs uint32) bool
	iSteamInventory_SerializeResult                        func(steamInventory uintptr, resultHandle SteamInventoryResult, pOutBuffer *[]byte, punOutBufferSize *uint32) bool
	iSteamInventory_DeserializeResult                      func(steamInventory uintptr, pOutResultHandle *SteamInventoryResult, pBuffer *[]byte, unBufferSize uint32, bRESERVEDMUSTBEFALSE bool) bool
	iSteamInventory_GenerateItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) bool
	iSteamInventory_GrantPromoItems                        func(steamInventory uintptr, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_AddPromoItem                           func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemDef SteamItemDef) bool
	iSteamInventory_AddPromoItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayItemDefs []SteamItemDef, unArrayLength uint32) bool
	iSteamInventory_ConsumeItem                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemConsume SteamItemInstanceID, unQuantity uint32) bool
	iSteamInventory_ExchangeItems                          func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pArrayGenerate []SteamItemDef, punArrayGenerateQuantity []uint32, unArrayGenerateLength uint32, pArrayDestroy []SteamItemInstanceID, punArrayDestroyQuantity []uint32, unArrayDestroyLength uint32) bool
	iSteamInventory_TransferItemQuantity                   func(steamInventory uintptr, pResultHandle *SteamInventoryResult, itemIdSource SteamItemInstanceID, unQuantity uint32, itemIdDest SteamItemInstanceID) bool
	iSteamInventory_TriggerItemDrop                        func(steamInventory uintptr, pResultHandle *SteamInventoryResult, dropListDefinition SteamItemDef) bool
	iSteamInventory_LoadItemDefinitions                    func(steamInventory uintptr) bool
	iSteamInventory_GetItemDefinitionIDs                   func(steamInventory uintptr, pItemDefIDs *[]SteamItemDef, punItemDefIDsArraySize *uint32) bool
	iSteamInventory_GetItemDefinitionProperty              func(steamInventory uintptr, iDefinition SteamItemDef, pchPropertyName string, pchValueBuffer *[]byte, punValueBufferSizeOut *uint32) bool
	iSteamInventory_RequestEligiblePromoItemDefinitionsIDs func(steamInventory uintptr, steamID Uint64SteamID) SteamAPICall
	iSteamInventory_GetEligiblePromoItemDefinitionIDs      func(steamInventory uintptr, steamID Uint64SteamID, pItemDefIDs *[]SteamItemDef, punItemDefIDsArraySize *uint32) bool
	iSteamInventory_StartPurchase                          func(steamInventory uintptr, pArrayItemDefs []SteamItemDef, punArrayQuantity []uint32, unArrayLength uint32) SteamAPICall
	iSteamInventory_RequestPrices                          func(steamInventory uintptr) SteamAPICall
	iSteamInventory_GetNumItemsWithPrices                  func(steamInventory uintptr) uint32
	iSteamInventory_GetItemsWithPrices                     func(steamInventory uintptr, pArrayItemDefs *[]SteamItemDef, pCurrentPrices *[]uint64, pBasePrices *[]uint64, unArrayLength uint32) bool
	iSteamInventory_GetItemPrice                           func(steamInventory uintptr, iDefinition SteamItemDef, pCurrentPrice *uint64, pBasePrice *uint64) bool
	iSteamInventory_StartUpdateProperties                  func(steamInventory uintptr) SteamInventoryUpdateHandle
	iSteamInventory_RemoveProperty                         func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string) bool
	iSteamInventory_SetPropertyString                      func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, pchPropertyValue string) bool
	iSteamInventory_SetPropertyBool                        func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, bValue bool) bool
	iSteamInventory_SetPropertyInt64                       func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, nValue int64) bool
	iSteamInventory_SetPropertyFloat                       func(steamInventory uintptr, handle SteamInventoryUpdateHandle, nItemID SteamItemInstanceID, pchPropertyName string, flValue float32) bool
	iSteamInventory_SubmitUpdateProperties                 func(steamInventory uintptr, handle SteamInventoryUpdateHandle, pResultHandle *SteamInventoryResult) bool
	iSteamInventory_InspectItem                            func(steamInventory uintptr, pResultHandle *SteamInventoryResult, pchItemToken string) bool
)

var steamInventory uintptr

func inventory() uintptr {
	if steamInventory == 0 {
		steamInventory = steamInventory_init()
		return steamInventory
	}
	return steamInventory
}

func GetResultStatus(resultHandle SteamInventoryResult) EResult {
	return iSteamInventory_GetResultStatus(inventory(), resultHandle)
}

func GetResultItems(resultHandle SteamInventoryResult, outItemsArraySize uint32) ([]SteamItemDetails, bool) {
	var pOutItemsArray []SteamItemDetails = make([]SteamItemDetails, outItemsArraySize)
	result := iSteamInventory_GetResultItems(inventory(), resultHandle, &pOutItemsArray, &outItemsArraySize)
	return pOutItemsArray[:outItemsArraySize], result
}

func GetResultItemProperty(resultHandle SteamInventoryResult, itemIndex uint32, propertyName string, valueBufferSizeOut uint32) ([]byte, bool) {
	var pchValueBuffer []byte = make([]byte, 0, valueBufferSizeOut)
	result := iSteamInventory_GetResultItemProperty(inventory(), resultHandle, itemIndex, propertyName, &pchValueBuffer, &valueBufferSizeOut)
	return pchValueBuffer[:valueBufferSizeOut], result
}

func GetResultTimestamp(resultHandle SteamInventoryResult) uint32 {
	return iSteamInventory_GetResultTimestamp(inventory(), resultHandle)
}

func CheckResultSteamID(resultHandle SteamInventoryResult, expectedSteamID Uint64SteamID) bool {
	return iSteamInventory_CheckResultSteamID(inventory(), resultHandle, expectedSteamID)
}

func DestroyResult(resultHandle SteamInventoryResult) {
	iSteamInventory_DestroyResult(inventory(), resultHandle)
}

func GetAllItems() (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GetAllItems(inventory(), &resultHandle)
	return resultHandle, success
}

func GetItemsByID(countInstanceIDs uint32) (resultHandle SteamInventoryResult, instanceIDs []SteamItemInstanceID, result bool) {
	instanceIDs = make([]SteamItemInstanceID, 0, countInstanceIDs)
	result = iSteamInventory_GetItemsByID(inventory(), &resultHandle, &instanceIDs, countInstanceIDs)
	return resultHandle, instanceIDs, result
}

func SerializeResult(resultHandle SteamInventoryResult, outBufferSize uint32) ([]byte, bool) {
	var pOutBuffer []byte = make([]byte, 0, outBufferSize)
	result := iSteamInventory_SerializeResult(inventory(), resultHandle, &pOutBuffer, &outBufferSize)
	return pOutBuffer[:outBufferSize], result
}

func DeserializeResult(bufferSize uint32, bRESERVEDMUSTBEFALSE bool) (resultHandle SteamInventoryResult, buffer []byte, success bool) {
	var pBuffer []byte = make([]byte, 0, bufferSize)
	success = iSteamInventory_DeserializeResult(inventory(), &resultHandle, &pBuffer, bufferSize, bRESERVEDMUSTBEFALSE)
	return resultHandle, pBuffer, success
}

func GenerateItems(arrayItemDefs []SteamItemDef, arrayQuantity []uint32, arrayLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GenerateItems(inventory(), &resultHandle, arrayItemDefs, arrayQuantity, arrayLength)
	return resultHandle, success
}

func GrantPromoItems() (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_GrantPromoItems(inventory(), &resultHandle)
	return resultHandle, success
}

func AddPromoItem(itemDef SteamItemDef) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_AddPromoItem(inventory(), &resultHandle, itemDef)
	return resultHandle, success
}

func AddPromoItems(arrayItemDefs []SteamItemDef, arrayLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_AddPromoItems(inventory(), &resultHandle, arrayItemDefs, arrayLength)
	return resultHandle, success
}

func ConsumeItem(itemConsume SteamItemInstanceID, quantity uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_ConsumeItem(inventory(), &resultHandle, itemConsume, quantity)
	return resultHandle, success
}

func ExchangeItems(arrayGenerate []SteamItemDef, arrayGenerateQuantity []uint32, arrayGenerateLength uint32, arrayDestroy []SteamItemInstanceID, arrayDestroyQuantity []uint32, arrayDestroyLength uint32) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_ExchangeItems(inventory(), &resultHandle, arrayGenerate, arrayGenerateQuantity, arrayGenerateLength, arrayDestroy, arrayDestroyQuantity, arrayDestroyLength)
	return resultHandle, success
}

func TransferItemQuantity(itemIdSource SteamItemInstanceID, quantity uint32, itemIdDest SteamItemInstanceID) (pResultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_TransferItemQuantity(inventory(), &pResultHandle, itemIdSource, quantity, itemIdDest)
	return pResultHandle, success
}

func TriggerItemDrop(dropListDefinition SteamItemDef) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_TriggerItemDrop(inventory(), &resultHandle, dropListDefinition)
	return resultHandle, success
}

func LoadItemDefinitions() bool {
	return iSteamInventory_LoadItemDefinitions(inventory())
}

func GetItemDefinitionIDs(itemDefIDsArraySize uint32) ([]SteamItemDef, bool) {
	var pItemDefIDs []SteamItemDef = make([]SteamItemDef, itemDefIDsArraySize)
	result := iSteamInventory_GetItemDefinitionIDs(inventory(), &pItemDefIDs, &itemDefIDsArraySize)
	return pItemDefIDs[:itemDefIDsArraySize], result
}

func GetItemDefinitionProperty(definition SteamItemDef, propertyName string, valueBufferSizeOut uint32) ([]byte, bool) {
	var pchValueBuffer []byte = make([]byte, 0, valueBufferSizeOut)
	result := iSteamInventory_GetItemDefinitionProperty(inventory(), definition, propertyName, &pchValueBuffer, &valueBufferSizeOut)
	return pchValueBuffer, result
}

func RequestEligiblePromoItemDefinitionsIDs(steamID Uint64SteamID) CallResult[SteamInventoryEligiblePromoItemDefIDs] {
	handle := iSteamInventory_RequestEligiblePromoItemDefinitionsIDs(inventory(), steamID)
	return CallResult[SteamInventoryEligiblePromoItemDefIDs]{Handle: handle}
}

func GetEligiblePromoItemDefinitionIDs(steamID Uint64SteamID, itemDefIDsArraySize uint32) ([]SteamItemDef, bool) {
	var pItemDefIDs []SteamItemDef = make([]SteamItemDef, itemDefIDsArraySize)
	result := iSteamInventory_GetEligiblePromoItemDefinitionIDs(inventory(), steamID, &pItemDefIDs, &itemDefIDsArraySize)
	return pItemDefIDs[:itemDefIDsArraySize], result
}

func StartPurchase(arrayItemDefs []SteamItemDef, punArrayQuantity []uint32, arrayLength uint32) CallResult[SteamInventoryStartPurchaseResult] {
	handle := iSteamInventory_StartPurchase(inventory(), arrayItemDefs, punArrayQuantity, arrayLength)
	return CallResult[SteamInventoryStartPurchaseResult]{Handle: handle}
}

func RequestPrices() CallResult[SteamInventoryRequestPricesResult] {
	handle := iSteamInventory_RequestPrices(inventory())
	return CallResult[SteamInventoryRequestPricesResult]{Handle: handle}
}

func GetNumItemsWithPrices() uint32 {
	return iSteamInventory_GetNumItemsWithPrices(inventory())
}

func GetItemsWithPrices(arrayLength uint32) ([]SteamItemDef, []uint64, []uint64, bool) {
	var pArrayItemDefs []SteamItemDef = make([]SteamItemDef, arrayLength)
	var pCurrentPrices []uint64 = make([]uint64, arrayLength)
	var pBasePrices []uint64 = make([]uint64, arrayLength)
	result := iSteamInventory_GetItemsWithPrices(inventory(), &pArrayItemDefs, &pCurrentPrices, &pBasePrices, arrayLength)
	return pArrayItemDefs, pCurrentPrices, pBasePrices, result
}

func GetItemPrice(definition SteamItemDef) (currentPrice uint64, basePrice uint64, succes bool) {
	succes = iSteamInventory_GetItemPrice(inventory(), definition, &currentPrice, &basePrice)
	return currentPrice, basePrice, succes
}

func StartUpdateProperties() SteamInventoryUpdateHandle {
	return iSteamInventory_StartUpdateProperties(inventory())
}

func RemoveProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string) bool {
	return iSteamInventory_RemoveProperty(inventory(), handle, itemID, propertyName)
}

func SetProperty(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, pchPropertyValue string) bool {
	return iSteamInventory_SetPropertyString(inventory(), handle, itemID, propertyName, pchPropertyValue)
}

func SetPropertyBool(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value bool) bool {
	return iSteamInventory_SetPropertyBool(inventory(), handle, itemID, propertyName, value)
}

func SetPropertyInt64y(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value int64) bool {
	return iSteamInventory_SetPropertyInt64(inventory(), handle, itemID, propertyName, value)
}

func SetPropertyFloat(handle SteamInventoryUpdateHandle, itemID SteamItemInstanceID, propertyName string, value float32) bool {
	return iSteamInventory_SetPropertyFloat(inventory(), handle, itemID, propertyName, value)
}

func SubmitUpdateProperties(handle SteamInventoryUpdateHandle) (resultHandle SteamInventoryResult, succes bool) {
	succes = iSteamInventory_SubmitUpdateProperties(inventory(), handle, &resultHandle)
	return resultHandle, succes
}

func InspectItem(itemToken string) (resultHandle SteamInventoryResult, success bool) {
	success = iSteamInventory_InspectItem(inventory(), &resultHandle, itemToken)
	return resultHandle, success
}
