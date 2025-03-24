package puregosteamworks

import "golang.org/x/sys/windows"

func bytePtrToString(p *byte) string {
	return windows.BytePtrToString(p)
}

func bytePtrFromString(s string) *byte {
	p, err := windows.BytePtrFromString(s)
	if err != nil {
		panic(err)
	}
	return p
}

func byteSliceFromString(s string) []byte {
	p, err := windows.ByteSliceFromString(s)
	if err != nil {
		panic(err)
	}
	return p
}

func byteSliceToString(s []byte) string {
	return windows.ByteSliceToString(s)
}
