sources = $(shell find . -path ./vendor -prune -o -name '*.go' -print)

dist:
	mkdir dist

dist/mc.exe: $(sources) | dist
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o dist/mc.exe

dist/mc-osx: $(sources) | dist
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/mc-osx

dist/mc-linux: $(sources) | dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/mc-linux

.PHONY: build run tag release clean
build: dist/mc.exe dist/mc-osx dist/mc-linux

run:
	go run main.go

tag:
	./scripts/tag.sh $(VERSION)

release: tag build
	./scripts/release.sh mc $(VERSION) dist/*

clean :
	-rm -r dist
