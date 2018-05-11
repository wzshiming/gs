#! /usr/bin/env gs


mm := map {
	"in" : map {},
	1: "map",
}

mm["in"]["i"] = "pass"

println mm[1] + "." + mm["in"]["i"]

