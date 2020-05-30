#! /usr/bin/env node

let sum = 0
for(let i = 0; i < 1000000; i ++) {
	sum += i + 1
}
console.log(sum)
