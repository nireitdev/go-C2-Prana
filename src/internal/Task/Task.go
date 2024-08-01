package Task

import "github.com/google/uuid"

type Task struct {
	TaskID    string `json:"taskid"`
	ImplantID string `json:"implantid"`
	Command   string `json:"command"`
	//RunTime time
	//IsRepeatble bool
	Result string `json:"result"` //solo un resultado por ahora
}

func NewTask(implantID string, cmd string) *Task {
	return &Task{
		TaskID:    uuid.NewString(),
		ImplantID: implantID,
		Command:   cmd,
	}
}

type TaskList struct {
	Tasks map[string]Task
}

func NewTaskList() *TaskList {
	tl := make(map[string]Task)
	return &TaskList{Tasks: tl}
}
