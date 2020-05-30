#! /usr/bin/env gs

sum := 0
for i := 0; i < 1000000; i ++ {
	sum += i + 1
}
println sum
