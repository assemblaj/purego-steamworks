package puregosteamworks

import (
	"syscall"
)

func openLibrary(name string) (uintptr, error) {
	handle, err := syscall.LoadLibrary(name)
	return uintptr(handle), err
}

func openSymbol(lib uintptr, name string) (uintptr, error) {
	return syscall.GetProcAddress(syscall.Handle(lib), name)
}

func closeLibrary(lib uintptr) error {
	return syscall.FreeLibrary(syscall.Handle(lib))
}
