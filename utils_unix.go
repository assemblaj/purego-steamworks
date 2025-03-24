//go:build darwin || freebsd || linux

package puregosteamworks

import "golang.org/x/sys/unix"

func bytePtrToString(p *byte) string {
	return unix.BytePtrToString(p)
}

func bytePtrFromString(s string) *byte {
	p, err := unix.BytePtrFromString(s)
	if err != nil {
		panic(err)
	}
	return p
}

func byteSliceFromString(s string) []byte {
	p, err := unix.ByteSliceFromString(s)
	if err != nil {
		panic(err)
	}
	return p
}

func byteSliceToString(s []byte) string {
	return unix.ByteSliceToString(s)
}
