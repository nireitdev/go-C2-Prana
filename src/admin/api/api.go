package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.dev/nireitdev/go-C2-Prana/internal/Task"
	"log"
	"net/http"
)

type ApiRest struct {
	listen string
}

func NewApiRest(listen string) *ApiRest {
	return &ApiRest{
		listen: listen,
	}
}

func (a *ApiRest) Run(ctx context.Context, chanTask chan Task.Task) error {
	router := http.NewServeMux()

	router.HandleFunc("POST /task", func(writer http.ResponseWriter, request *http.Request) {

		var task Task.Task
		task.TaskID = uuid.NewString()

		if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
			log.Println("error parsing json")

		} else {
			log.Printf("Task : %v \n", task)

			chanTask <- task
			writer.Write([]byte("taskID:: " + task.TaskID + " " + "task:: " + task.Command + " "))
		}

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
