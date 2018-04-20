#! /usr/bin/env gs

sum := 0
for i := 0; i < 100; i ++ {
	sum += i + 1
}

if sum != 5050 {
	return "cycle.fail"
}

return "cycle.pass"