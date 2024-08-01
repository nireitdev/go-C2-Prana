
BINDIR := bin
BINARIOS := admin implant server

.PHONY: $(BINARIOS) proto

all: $(BINARIOS)

$(BINDIR):
	mkdir -p $(BINDIR)

$(BINARIOS): |$(BINDIR)
	go build -C src/  -o ../$(BINDIR)/$@ $@/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/internal/proto/grpc.proto


