#! /usr/bin/env gs

func T i, j {
	if j >= 2 {
		return T i*i, j-1
	}
	return i
}

if 256 != T 2, 4 {
	return "1.fail"
}

return "1.pass"