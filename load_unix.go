//go:build darwin || freebsd || linux

package puregosteamworks

import "github.com/ebitengine/purego"

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}

func openSymbol(lib uintptr, name string) (uintptr, error) {
	return purego.Dlsym(lib, name)
}

func closeLibrary(lib uintptr) error {
	return purego.Dlclose(lib)
}
