package steamworks

import (
	"runtime"
	"unsafe"
)

const (
	flatAPI_SteamRemoteStorage              = "SteamAPI_SteamRemoteStorage_v016"
	flatAPI_ISteamRemoteStorage_FileWrite   = "SteamAPI_ISteamRemoteStorage_FileWrite"
	flatAPI_ISteamRemoteStorage_FileRead    = "SteamAPI_ISteamRemoteStorage_FileRead"
	flatAPI_ISteamRemoteStorage_FileDelete  = "SteamAPI_ISteamRemoteStorage_FileDelete"
	flatAPI_ISteamRemoteStorage_GetFileSize = "SteamAPI_ISteamRemoteStorage_GetFileSize"
)

type steamRemoteStorage uintptr

var (
	steamRemoteStorage_init func() uintptr
	fileWrite               func(steamRemoteStorage uintptr, file uintptr, data uintptr, len uintptr) uintptr
	fileRead                func(steamRemoteStorage uintptr, file uintptr, data uintptr, len uintptr) uintptr
	fileDelete              func(steamRemoteStorage uintptr, file uintptr) uintptr
	getFileSize             func(steamRemoteStorage uintptr, file uintptr) uintptr
)

func (s steamRemoteStorage) FileWrite(file string, data []byte) bool {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	defer runtime.KeepAlive(data)

	result := fileWrite(uintptr(s), uintptr(unsafe.Pointer(&cfile[0])), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return byte(result) != 0
}

func (s steamRemoteStorage) FileRead(file string, data []byte) int32 {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	defer runtime.KeepAlive(data)

	result := fileRead(uintptr(s), uintptr(unsafe.Pointer(&cfile[0])), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return int32(result)
}

func (s steamRemoteStorage) FileDelete(file string) bool {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	result := fileDelete(uintptr(s), uintptr(unsafe.Pointer(&cfile[0])))
	return byte(result) != 0
}

func (s steamRemoteStorage) GetFileSize(file string) int32 {
	cfile := append([]byte(file), 0)
	defer runtime.KeepAlive(cfile)

	result := getFileSize(uintptr(s), uintptr(unsafe.Pointer(&cfile[0])))
	return int32(result)
}
