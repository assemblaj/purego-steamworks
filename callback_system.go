package puregosteamworks

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type CallResult[T Callback] struct {
	Handle SteamAPICall
}

func (cr CallResult[T]) String() string {
	var zero T
	return fmt.Sprintf("%s API HANDLE: %d", zero, cr.Handle)
}

func callbackString(callback Callback) string {
	return fmt.Sprintf("Callback Type: %s, Callback ID: %d", callback.Name(), callback.CallbackID())
}

type Callback interface {
	CallbackID() SteamCallbackID
	Name() string
	String() string
}

type callbackRegistry struct {
	callbacks   map[SteamCallbackID]any
	callResults map[SteamAPICall]any
	mu          sync.RWMutex
}
type callbackHandler[T Callback] func(T)
type callResultHandler[T Callback] func(T, bool)

var registry = &callbackRegistry{
	callbacks:   make(map[SteamCallbackID]any),
	callResults: make(map[SteamAPICall]any),
}

func RegisterCallback[T Callback](handler callbackHandler[T]) {
	var zero T
	callbackID := zero.CallbackID()

	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.callbacks[callbackID] = handler
}

func handleRawData[T Callback](callbackID SteamCallbackID, rawcallbackdata unsafe.Pointer) {
	var zero T

	registry.mu.RLock()
	handler, exists := registry.callbacks[callbackID]
	registry.mu.RUnlock()

	if !exists {
		fmt.Printf("No handler registered for %s\n", zero.String())
		return
	}

	result := *(*T)(rawcallbackdata)

	if typedHandler, ok := handler.(callbackHandler[T]); ok {
		typedHandler(result)
	} else {
		fmt.Printf("Type: %T\n", handler)

		fmt.Printf("Type mismatch for %s\n", zero.String())
	}
}

func RegisterCallResult[T Callback](cr CallResult[T], handler callResultHandler[T]) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.callResults[cr.Handle] = handler
}

func handleRawResult[T Callback](apiHandle SteamAPICall, failed bool, rawcallbackdata unsafe.Pointer) {
	var zero T

	registry.mu.Lock()
	handler, exists := registry.callResults[apiHandle]
	registry.mu.Unlock()

	if !exists {
		fmt.Printf("No handler registered for API handle %d, %s \n", apiHandle, zero.String())
		return
	}

	result := *(*T)(rawcallbackdata)

	if typedHandler, ok := handler.(callResultHandler[T]); ok {
		typedHandler(result, failed)
		registry.mu.Lock()
		delete(registry.callResults, apiHandle)
		registry.mu.Unlock()
	} else {
		fmt.Printf("Type: %T\n", handler)

		fmt.Printf("Type mismatch for apiHandle %d, %s\n", apiHandle, zero.String())
	}
}

var (
	getHSteamPipe                   func() uintptr
	manualdispatch_init             func()
	manualdispatch_runFrame         func(HSteamPipe)
	manualdispatch_getNextCallback  func(HSteamPipe, *CallbackMsg) bool
	manualdispatch_freeLastCallback func(HSteamPipe)
	manualdispatch_getApiCallResult func(hSteamPipe HSteamPipe, hSteamAPICall SteamAPICall, pCallback uintptr, cubCallback int, iCallbackExpected int, pbFailed uintptr) bool
)

const (
	flatAPI_GetHSteamPipe                   = "SteamAPI_GetHSteamPipe"
	flatAPI_ManualDispatch_Init             = "SteamAPI_ManualDispatch_Init"
	flatAPI_ManualDispatch_RunFrame         = "SteamAPI_ManualDispatch_RunFrame"
	flatAPI_ManualDispatch_GetNextCallback  = "SteamAPI_ManualDispatch_GetNextCallback"
	flatAPI_ManualDispatch_FreeLastCallback = "SteamAPI_ManualDispatch_FreeLastCallback"
	flatAPI_ManualDispatch_GetAPICallResult = "SteamAPI_ManualDispatch_GetAPICallResult"
)

var once sync.Once

func RunCallbacksForever() {
	for {
		RunCallbacks()
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
}

func RunCallbacks() {
	once.Do(func() {
		manualdispatch_init()
	})

	pipe := HSteamPipe(getHSteamPipe())
	manualdispatch_runFrame(pipe)
	var msg CallbackMsg
	for manualdispatch_getNextCallback(pipe, &msg) {
		if SteamCallbackID(msg.Callback) == SteamAPICallCompletedID {
			pCallCompleted := (*SteamAPICallCompleted)(unsafe.Pointer(msg.ParamData))
			pTmpCallResult := uintptr(unsafe.Pointer(&make([]byte, pCallCompleted.Param)[0]))
			var failed bool
			if manualdispatch_getApiCallResult(pipe, pCallCompleted.AsyncCall, pTmpCallResult, int(pCallCompleted.Param), int(pCallCompleted.Callback), uintptr(unsafe.Pointer(&failed))) {
				dispatchCallResult(pCallCompleted.AsyncCall, failed, pTmpCallResult, SteamCallbackID(pCallCompleted.Callback))
			}
			//runtime.KeepAlive(pCallCompleted)
			//runtime.KeepAlive(pTmpCallResult)
		} else {
			dispatchCallback(&msg)
			//runtime.KeepAlive(msg)
		}
		manualdispatch_freeLastCallback(pipe)
	}

}
