package main

import (
	"context"
	"fmt"
	"github.dev/nireitdev/go-C2-Prana/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	listen string

	proto.UnimplementedServerServer
}

func (s Server) AdminCmd(ctx context.Context, req *proto.AdminRequest) (*proto.AdminResult, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method ImplantCmd not implemented")
	log.Println("Recibido: ", "Hola Admin!")
	res := proto.AdminResult{Result: req.Request}
	return &res, nil
}
func (s Server) ImplantCmd(ctx context.Context, req *proto.ImplantRequest) (*proto.ImplantResult, error) {
	log.Println("Recibido desde Implant: ", req.Request)
	res := proto.ImplantResult{Result: "Hola Implant!"}
	return &res, nil
}

func main() {
	var err error
	srv := Server{}
	srv.listen = fmt.Sprintf("%s:%s", "0.0.0.0", "3000")

	listen, err := net.Listen("tcp", srv.listen)
	if err != nil {
		log.Fatalln("Imposible crear listener")
	}
	defer listen.Close()
	log.Println("Iniciando server en: ", srv.listen)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer) //TODO: habilitar solo en Development mode!

	proto.RegisterServerServer(grpcServer, srv)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error abriendo socket")
	}

}
