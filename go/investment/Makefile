.DEFAULT_GOAL := build

fmt:
	golangci-lint run --disable-all --no-config -Egofmt --fix
	golangci-lint run --disable-all --no-config -Egofumpt --fix

revive: fmt
	revive -config $(HOME)/go/extras/revive.toml

lint: revive
	golangci-lint run

test:
	go test -shuffle on .

testv:
	go test -shuffle on -v .

.PHONY: fmt revive lint test testv
