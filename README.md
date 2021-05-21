# ErrorLinter 

This linter will scan for duplicate error codes in a repository. The types of errors that will be counted are:

```
cberr := errors.New(123, "error")
cberr = errors.NewWrapped(123, err, "error")
cberr = errorFx(123, err, "error")
```

## To develop:
- Make updates to `example.go` file
- To build, `make build`
- To test, `make test` (Note: Tests currently aren't being scored correctly, but you can skim output)
- To deploy to a repo of your choice, update the destination and then `make deploy` (generalizing this is a TODO)


## To use in a repo

Step one: custom golangci-lint
- Build custom `golangci-lint`; start by cloning from https://github.com/golangci/golangci-lint.
- Update the `CGO_ENABLED` flag in `.goreleaser.yml` to be 1.
- Run `make build`
- Move output binary to your go path

Step two: update your repo's `.golangci.yml`
- Add the following top-level item to your golangci.yml:

```
linters-settings:
  custom:
    duplicate_errors:
      path: ./example.so
      description: Linter to scan for duplicate errors. 
      original-url: https://github.com/cbio-laura/example-linter

```
- Add the following to your enabled linters (note: don't need to remove other enabled linters):
```
linters:
  enable:
    - duplicate_errors
```
