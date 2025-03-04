package inventory

import . "github.com/assemblaj/purego-steamworks"

const (
	SteamInventoryResultReadyID             SteamCallbackID = 4700
	SteamInventoryFullUpdateID              SteamCallbackID = 4701
	SteamInventoryDefinitionUpdateID        SteamCallbackID = 4702
	SteamInventoryEligiblePromoItemDefIDsID SteamCallbackID = 4703
	SteamInventoryStartPurchaseResultID     SteamCallbackID = 4704
	SteamInventoryRequestPricesResultID     SteamCallbackID = 4705
)

type SteamInventoryResultReady struct {
	Handle SteamInventoryResult
	Result EResult
}
type SteamInventoryFullUpdate struct {
	Handle SteamInventoryResult
}
type SteamInventoryDefinitionUpdate struct {
}
type SteamInventoryEligiblePromoItemDefIDs struct {
	Result                   EResult
	_                        [4]byte
	SteamID                  CSteamID
	NumEligiblePromoItemDefs int32
	CachedData               bool
	_                        [3]byte
}
type SteamInventoryStartPurchaseResult struct {
	Result  EResult
	_       [4]byte
	OrderID uint64
	TransID uint64
}
type SteamInventoryRequestPricesResult struct {
	Result   EResult
	Currency [4]byte
}

func (cb SteamInventoryResultReady) CallbackID() SteamCallbackID {
	return SteamInventoryResultReadyID
}

func (cb SteamInventoryResultReady) Name() string {
	return "SteamInventoryResultReady"
}

func (cb SteamInventoryResultReady) String() string {
	return CallbackString(cb)
}

func (cb SteamInventoryFullUpdate) CallbackID() SteamCallbackID {
	return SteamInventoryFullUpdateID
}

func (cb SteamInventoryFullUpdate) Name() string {
	return "SteamInventoryFullUpdate"
}

func (cb SteamInventoryFullUpdate) String() string {
	return CallbackString(cb)
}

func (cb SteamInventoryDefinitionUpdate) CallbackID() SteamCallbackID {
	return SteamInventoryDefinitionUpdateID
}

func (cb SteamInventoryDefinitionUpdate) Name() string {
	return "SteamInventoryDefinitionUpdate"
}

func (cb SteamInventoryDefinitionUpdate) String() string {
	return CallbackString(cb)
}

func (cb SteamInventoryEligiblePromoItemDefIDs) CallbackID() SteamCallbackID {
	return SteamInventoryEligiblePromoItemDefIDsID
}

func (cb SteamInventoryEligiblePromoItemDefIDs) Name() string {
	return "SteamInventoryEligiblePromoItemDefIDs"
}

func (cb SteamInventoryEligiblePromoItemDefIDs) String() string {
	return CallbackString(cb)
}

func (cb SteamInventoryStartPurchaseResult) CallbackID() SteamCallbackID {
	return SteamInventoryStartPurchaseResultID
}

func (cb SteamInventoryStartPurchaseResult) Name() string {
	return "SteamInventoryStartPurchaseResult"
}

func (cb SteamInventoryStartPurchaseResult) String() string {
	return CallbackString(cb)
}

func (cb SteamInventoryRequestPricesResult) CallbackID() SteamCallbackID {
	return SteamInventoryRequestPricesResultID
}

func (cb SteamInventoryRequestPricesResult) Name() string {
	return "SteamInventoryRequestPricesResult"
}

func (cb SteamInventoryRequestPricesResult) String() string {
	return CallbackString(cb)
}
