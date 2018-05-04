#! /usr/bin/env gs

a, ...b, c := 1, (2, 3)..., (4, 5, 6)...


if b[2] != 4 {
	return "ellipsis.fail"
}

return "ellipsis.pass"