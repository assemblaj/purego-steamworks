package puregosteamworks

import (
	"syscall"
)

func OpenLibrary(name string) (uintptr, error) {
	handle, err := syscall.LoadLibrary(name)
	return uintptr(handle), err
}

func CloseLibrary(lib uintptr) error {
	return syscall.FreeLibrary(syscall.Handle(lib))
}
