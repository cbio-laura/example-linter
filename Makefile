build:
	go build -buildmode=plugin plugin/*.go

test:
	go test ./...

deploy: build 
	cp example.so /Users/laurahunter/cbio/repo/service/panels

mod-update:
	go mod tidy
	go mod vendor
	git add go.*
	git add vendor
	git commit -m "mod update"