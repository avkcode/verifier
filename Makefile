BINARY := verifier
GO := go
GOFLAGS :=
LDFLAGS := -w -s

.PHONY: all build clean release

all: build

build:
	$(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY) ./cmd/verifier

clean:
	rm -rf bin/

release:
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY)-linux-amd64 ./cmd/verifier
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY)-linux-arm64 ./cmd/verifier
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY)-darwin-amd64 ./cmd/verifier
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY)-darwin-arm64 ./cmd/verifier
