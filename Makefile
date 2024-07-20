
BINDIR := bin
BINARIOS := admin implant server

.PHONY: $(BINARIOS) proto

all: $(BINARIOS)

$(BINDIR):
	mkdir -p $(BINDIR)

$(BINARIOS): |$(BINDIR)
	go build -o $(BINDIR)/$@ ./src/$@/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/proto/grpc.proto


