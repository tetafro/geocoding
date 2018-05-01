.PHONY: dep
dep:
	dep ensure -v

.PHONY: build
build:
	go build -o ./bin/geocoding ./cmd/geocoding
