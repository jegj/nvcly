install:
	go mod download
run:
	go run main.go
test:
	go test -v ./...
vet:
	go vet main.go
vet_shadow:
	go vet -vettool=$(which shadow) main.go
staticcheck:
	staticcheck ./...
audit:
	govulncheck -mode binary -show verbose nvcly
gosec:
	gosec ./...

