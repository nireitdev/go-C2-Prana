package Task

import (
	"log/slog"
	"time"
)

const (
	WAITINGTORUN = iota
	RUNNING
	RUNNEDOK
	RUNNEDFAIL
	SENDOK
)

type Task struct {
	TaskID    string `json:"taskid"`
	ImplantID string `json:"implantid"`
	Command   string `json:"command"`
	//RunTime time
	//IsRepeatble bool
	StatusRun int
	Result    string `json:"result"` //solo un resultado por ahora
}

func NewTask(taskID string, implantID string, cmd string) *Task {
	return &Task{
		TaskID:    taskID,
		ImplantID: implantID,
		Command:   cmd,
		StatusRun: WAITINGTORUN,
	}
}

type TaskList struct {
	Tasks map[string]Task
}

func NewTaskList() *TaskList {
	tl := make(map[string]Task)
	return &TaskList{Tasks: tl}
}

// Print debug information about the list of task
func (tl TaskList) PrintTaskListInfo() {
	for {
		time.Sleep(10 * time.Second)
		totaltask := len(tl.Tasks)
		totaldone := 0
		totalwait := 0
		for _, task := range tl.Tasks {
			if task.StatusRun == WAITINGTORUN {
				totalwait++
			}
			if task.StatusRun == RUNNEDOK {
				totaldone++
			}
		}
		slog.Info("Tasks:", "total", totaltask, "done", totaldone, "wait", totalwait)
	}
}
