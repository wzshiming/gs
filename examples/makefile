
GO = /usr/bin/env go
GS = /usr/bin/env gs
GITPATH = github.com/wzshiming/gs/cmd/gs

.PHONY: gs test

default: gs test
	
gs:
	- $(GO) get -v $(GITPATH)

test:
	- $(GS) */*.gs


