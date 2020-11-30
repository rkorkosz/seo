CGO_ENABLED=0
LDFLAGS="-w -s"

build-crwl:
	go build -ldflags=${LDFLAGS} -o target/crwl ./cmd/crwl/main.go

build-chksts:
	go build -ldflags=${LDFLAGS} -o target/chksts ./cmd/chksts/main.go

build: build-crwl build-chksts
