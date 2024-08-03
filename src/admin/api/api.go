package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	"log"
	"log/slog"
	"net/http"
)

type ApiRest struct {
	listen    string
	chanTask  chan Task.Task
	taskslist Task.TaskList
}

func NewApiRest(listen string, ch chan Task.Task, list *Task.TaskList) *ApiRest {
	return &ApiRest{
		listen:    listen,
		chanTask:  ch,
		taskslist: *list,
	}
}

func (a *ApiRest) Run(ctx context.Context) error {
	router := http.NewServeMux()

	router.HandleFunc("POST /task", func(writer http.ResponseWriter, request *http.Request) {
		var task Task.Task
		task.TaskID = uuid.NewString()

		if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
			log.Println("error parsing json")
		} else {
			log.Printf("Task : %v \n", task)
			a.chanTask <- task
			writer.Write([]byte("taskID:: " + task.TaskID + " " + "task:: " + task.Command + " "))
		}
	})

	router.HandleFunc("GET /task/{id}/", func(writer http.ResponseWriter, request *http.Request) {
		task, ok := a.taskslist.Tasks[request.PathValue("id")]
		//No existe el ID:
		if !ok {
			http.Error(writer, "Invalid ID", http.StatusNotFound)
			return
		}
		js, err := json.Marshal(task)
		if err != nil {
			slog.Error("Falla lista de tareas.")
			return
		}
		writer.Write([]byte(js))
	})

	server := http.Server{
		Addr:    a.listen,
		Handler: LoggerMiddleware(router),
	}
	err := server.ListenAndServe()
	return err
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		log.Printf("method: %s, path: %s \n", request.Method, request.URL.Path) //TODO: necesitaremos metricas??

		next.ServeHTTP(writer, request)
	}
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO: autenticacion de api, permisos de ejecucion, accesos a implants, etc
		next.ServeHTTP(writer, request)
	}
}
