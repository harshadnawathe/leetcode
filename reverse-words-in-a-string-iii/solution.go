package reversewordsinastringiii

import "unsafe"

func reverseWords(s string) string {
	bytes := make([]byte, len(s)+1)
	for i := 0; i < len(s); i++ {
		bytes[i] = s[i]
	}
	//sentinal
	bytes[len(s)] = ' '

	b := 0
	for e := 0; e < len(bytes); e++ {
		if bytes[e] == ' ' {
			reverse(bytes[b:e])
			b = e + 1
		}
	}

	//remove sentinal
	bytes = bytes[:len(s)]

	return *((*string)(unsafe.Pointer(&bytes)))
}

func reverse(bs []byte) []byte {
	l, r := 0, len(bs)
	for l < r {
		r--
		bs[l], bs[r] = bs[r], bs[l]
		l++
	}
	return bs
}
