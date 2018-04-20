
GO = /usr/bin/env go
GS = /usr/bin/env gs
GITPATH = github.com/wzshiming/gs/cmd/gs

default: gs test
	
gs:
	- $(GO) get -v $(GITPATH)
	$(GO) install $(GITPATH)

test:
	- $(GS) */*.gs


