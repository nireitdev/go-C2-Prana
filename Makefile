
BINDIR := bin
BINARIOS := admin implant server

.PHONY: $(BINARIOS) proto

include COMPILECFG
LDFLAGS=-ldflags "-w -s -X 'main.SERVERIP=$(SERVERIP)' -X 'main.SERVERPORT=$(SERVERPORT)' \
					-X 'main.APIRESTIP=$(APIRESTIP)' -X 'main.APIRESTPORT=$(APIRESTPORT)'"

all: $(BINARIOS)

$(BINDIR):
	mkdir -p $(BINDIR)

$(BINARIOS): |$(BINDIR)
	go build -C src/ ${LDFLAGS}  -o ../$(BINDIR)/$@ $@/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/internal/proto/grpc.proto


