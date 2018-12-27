# Project-specific variables
BINARIES ?=	anonuuid
SOURCES :=	$(shell find . -name "*.go")
COMMANDS :=	$(shell go list ./... | grep -v /vendor/ | grep /cmd/)
PACKAGES :=	$(shell go list ./... | grep -v /vendor/ | grep -v /cmd/)
GOENV ?=	GO111MODULE=on
GO ?=		$(GOENV) go
USER ?=		$(shell whoami)


all:	build


.PHONY: build
build:	$(BINARIES)


$(BINARIES):	$(SOURCES)
	$(GO) install ./cmd/...


.PHONY: test
test:
	$(GO) get -t .
	$(GO) test -v .


.PHONY: clean
clean:
	rm -f $(BINARIES)


.PHONY: re
re:	clean all


.PHONY:	cover
cover:	profile.out


profile.out:	$(SOURCES)
	rm -f $@
	$(GO) test -covermode=count -coverpkg=. -coverprofile=$@ .
