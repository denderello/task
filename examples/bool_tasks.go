package main

import (
	"errors"
	"log"

	"github.com/denderello/task"
)

type BoolTask struct {
	name               string
	RunSuccessful      bool
	RollbackSuccessful bool
}

func (t BoolTask) Name() string {
	return t.name
}

func (t BoolTask) Run() error {
	log.Printf("Running task %s", t.name)
	if !t.RunSuccessful {
		return errors.New("error")
	}
	return nil
}

func (t BoolTask) Rollback() error {
	log.Printf("Rolling back task %s", t.name)
	if !t.RollbackSuccessful {
		return errors.New("error")
	}
	return nil
}

func main() {
	ts := []task.Task{
		BoolTask{"task-1", true, true},
		BoolTask{"task-2", true, false},
		BoolTask{"task-3", true, true},
		BoolTask{"task-4", false, true},
		BoolTask{"task-5", true, true},
	}

	q := task.NewQueue(ts)
	err := q.Start()
	if err != nil {
		log.Fatalf("Task queue failed:\n%v", err)
	}
}
