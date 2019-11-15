package golib

func swapBool(boolean *bool) {
	if *boolean == false {
		*boolean = true
	} else {
		*boolean = false
	}
}
