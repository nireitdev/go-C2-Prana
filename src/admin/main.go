package main

import (
	"context"
	"github.dev/nireitdev/go-C2-Prana/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.Println("Connectando con el servidor...")

	opts := grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalln("Imposible conectar con el servidor. Bye!")
	}
	defer conn.Close()

	srv := proto.NewServerClient(conn)

	req := proto.AdminRequest{Request: "Hola!"}
	resp, err := srv.AdminCmd(context.Background(), &req)
	if err != nil {
		log.Fatalf("Fallo intercambio de msgs: %v \n", err)
	}

	log.Println("Recibido:", resp.Result)

}
