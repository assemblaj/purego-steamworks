package inventory

import (
	"sync"

	. "github.com/assemblaj/purego-steamworks"

	"github.com/ebitengine/purego"
)

var once sync.Once

func init() {
	purego.RegisterLibFunc(&steamInventory_init, SteamAPIDLL, flatAPI_SteamInventory)
	purego.RegisterLibFunc(&iSteamInventory_GetResultStatus, SteamAPIDLL, flatAPI_ISteamInventory_GetResultStatus)
	purego.RegisterLibFunc(&iSteamInventory_GetResultItems, SteamAPIDLL, flatAPI_ISteamInventory_GetResultItems)
	purego.RegisterLibFunc(&iSteamInventory_GetResultItemProperty, SteamAPIDLL, flatAPI_ISteamInventory_GetResultItemProperty)
	purego.RegisterLibFunc(&iSteamInventory_GetResultTimestamp, SteamAPIDLL, flatAPI_ISteamInventory_GetResultTimestamp)
	purego.RegisterLibFunc(&iSteamInventory_CheckResultSteamID, SteamAPIDLL, flatAPI_ISteamInventory_CheckResultSteamID)
	purego.RegisterLibFunc(&iSteamInventory_DestroyResult, SteamAPIDLL, flatAPI_ISteamInventory_DestroyResult)
	purego.RegisterLibFunc(&iSteamInventory_GetAllItems, SteamAPIDLL, flatAPI_ISteamInventory_GetAllItems)
	purego.RegisterLibFunc(&iSteamInventory_GetItemsByID, SteamAPIDLL, flatAPI_ISteamInventory_GetItemsByID)
	purego.RegisterLibFunc(&iSteamInventory_SerializeResult, SteamAPIDLL, flatAPI_ISteamInventory_SerializeResult)
	purego.RegisterLibFunc(&iSteamInventory_DeserializeResult, SteamAPIDLL, flatAPI_ISteamInventory_DeserializeResult)
	purego.RegisterLibFunc(&iSteamInventory_GenerateItems, SteamAPIDLL, flatAPI_ISteamInventory_GenerateItems)
	purego.RegisterLibFunc(&iSteamInventory_GrantPromoItems, SteamAPIDLL, flatAPI_ISteamInventory_GrantPromoItems)
	purego.RegisterLibFunc(&iSteamInventory_AddPromoItem, SteamAPIDLL, flatAPI_ISteamInventory_AddPromoItem)
	purego.RegisterLibFunc(&iSteamInventory_AddPromoItems, SteamAPIDLL, flatAPI_ISteamInventory_AddPromoItems)
	purego.RegisterLibFunc(&iSteamInventory_ConsumeItem, SteamAPIDLL, flatAPI_ISteamInventory_ConsumeItem)
	purego.RegisterLibFunc(&iSteamInventory_ExchangeItems, SteamAPIDLL, flatAPI_ISteamInventory_ExchangeItems)
	purego.RegisterLibFunc(&iSteamInventory_TransferItemQuantity, SteamAPIDLL, flatAPI_ISteamInventory_TransferItemQuantity)
	purego.RegisterLibFunc(&iSteamInventory_TriggerItemDrop, SteamAPIDLL, flatAPI_ISteamInventory_TriggerItemDrop)
	purego.RegisterLibFunc(&iSteamInventory_LoadItemDefinitions, SteamAPIDLL, flatAPI_ISteamInventory_LoadItemDefinitions)
	purego.RegisterLibFunc(&iSteamInventory_GetItemDefinitionIDs, SteamAPIDLL, flatAPI_ISteamInventory_GetItemDefinitionIDs)
	purego.RegisterLibFunc(&iSteamInventory_GetItemDefinitionProperty, SteamAPIDLL, flatAPI_ISteamInventory_GetItemDefinitionProperty)
	purego.RegisterLibFunc(&iSteamInventory_RequestEligiblePromoItemDefinitionsIDs, SteamAPIDLL, flatAPI_ISteamInventory_RequestEligiblePromoItemDefinitionsIDs)
	purego.RegisterLibFunc(&iSteamInventory_GetEligiblePromoItemDefinitionIDs, SteamAPIDLL, flatAPI_ISteamInventory_GetEligiblePromoItemDefinitionIDs)
	purego.RegisterLibFunc(&iSteamInventory_StartPurchase, SteamAPIDLL, flatAPI_ISteamInventory_StartPurchase)
	purego.RegisterLibFunc(&iSteamInventory_RequestPrices, SteamAPIDLL, flatAPI_ISteamInventory_RequestPrices)
	purego.RegisterLibFunc(&iSteamInventory_GetNumItemsWithPrices, SteamAPIDLL, flatAPI_ISteamInventory_GetNumItemsWithPrices)
	purego.RegisterLibFunc(&iSteamInventory_GetItemsWithPrices, SteamAPIDLL, flatAPI_ISteamInventory_GetItemsWithPrices)
	purego.RegisterLibFunc(&iSteamInventory_GetItemPrice, SteamAPIDLL, flatAPI_ISteamInventory_GetItemPrice)
	purego.RegisterLibFunc(&iSteamInventory_StartUpdateProperties, SteamAPIDLL, flatAPI_ISteamInventory_StartUpdateProperties)
	purego.RegisterLibFunc(&iSteamInventory_RemoveProperty, SteamAPIDLL, flatAPI_ISteamInventory_RemoveProperty)
	purego.RegisterLibFunc(&iSteamInventory_SetPropertyString, SteamAPIDLL, flatAPI_ISteamInventory_SetPropertyString)
	purego.RegisterLibFunc(&iSteamInventory_SetPropertyBool, SteamAPIDLL, flatAPI_ISteamInventory_SetPropertyBool)
	purego.RegisterLibFunc(&iSteamInventory_SetPropertyInt64, SteamAPIDLL, flatAPI_ISteamInventory_SetPropertyInt64)
	purego.RegisterLibFunc(&iSteamInventory_SetPropertyFloat, SteamAPIDLL, flatAPI_ISteamInventory_SetPropertyFloat)
	purego.RegisterLibFunc(&iSteamInventory_SubmitUpdateProperties, SteamAPIDLL, flatAPI_ISteamInventory_SubmitUpdateProperties)
	purego.RegisterLibFunc(&iSteamInventory_InspectItem, SteamAPIDLL, flatAPI_ISteamInventory_InspectItem)
}
