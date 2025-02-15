package steamworks

import "unsafe"

func cStringToGoString(v uintptr, sizeHint int) string {
	bs := make([]byte, 0, sizeHint)
	for i := int32(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(v))
		v += unsafe.Sizeof(byte(0))
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}
