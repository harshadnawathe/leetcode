package latesttimebyreplacinghiddendigits

import "unsafe"

func maximumTime(time string) string {
	bytes := []byte(time)
	if bytes[0] == '?' {
		if (bytes[1] >= '0' && bytes[1] < '4') || bytes[1] == '?' {
			bytes[0] = '2'
		} else {
			bytes[0] = '1'
		}
	}

	if bytes[1] == '?' {
		if bytes[0] == '2' {
			bytes[1] = '3'
		} else {
			bytes[1] = '9'
		}
	}

	if bytes[3] == '?' {
		bytes[3] = '5'
	}

	if bytes[4] == '?' {
		bytes[4] = '9'
	}
	return *(*string)(unsafe.Pointer(&bytes))
}
