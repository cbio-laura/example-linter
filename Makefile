build:
	go build -mod=vendor -buildmode=plugin plugin/*.go

test:
	go test ./...

deploy: build 
	cp example.so /Users/laurahunter/cbio/repo/service/panels

mod-update:
	GOPRIVATE=github.com/cerebrae/* go mod tidy
	GOPRIVATE=github.com/cerebrae/* go mod vendor
	git add go.*
	git add vendor
	git commit -m "mod update"