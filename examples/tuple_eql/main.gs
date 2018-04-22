#! /usr/bin/env gs

a,(b,c),d:=1,(2,3),4
a,b,c,d = d,c,b,a

if a != 4 {
	return "tuple_eql.fail"
}
if b != 3 {
	return "tuple_eql.fail"
}
if c != 2 {
	return "tuple_eql.fail"
}
if d != 1 {
	return "tuple_eql.fail"
}

return "tuple_eql.pass"
