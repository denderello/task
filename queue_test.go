package task

import "testing"

func TestFullySuccessfulQueue(t *testing.T) {
	ts := []Task{
		NewTaskRecorder("task-1", true, true),
		NewTaskRecorder("task-2", true, true),
		NewTaskRecorder("task-3", true, true),
	}
	q := NewQueue(ts)

	err := q.Start()
	if err != nil {
		t.Fatalf("Queue should have been executed without errors. Got error instead: %v", err)
	}

	for _, task := range ts {
		taskRecorder := task.(*TaskRecorder)
		if !taskRecorder.RunExecuted {
			t.Fatalf("Expected all tasks to be executed but %s was not.", taskRecorder.Name())
		}
	}
}

func TestUnsuccessfulQueue(t *testing.T) {
	ts := []Task{
		NewTaskRecorder("task-1", true, true),
		NewTaskRecorder("task-2", true, true),
		NewTaskRecorder("task-3", true, true),
		NewTaskRecorder("task-4", false, true),
		NewTaskRecorder("task-5", true, true),
	}
	q := NewQueue(ts)

	err := q.Start()
	if err == nil {
		t.Fatal("Queue should have been executed with errors")
	}

	for _, task := range ts[0:4] {
		taskRecorder := task.(*TaskRecorder)
		if !taskRecorder.RunExecuted {
			t.Fatalf("Expected all tasks to be executed but %s was not.", taskRecorder.Name())
		}

		if !taskRecorder.RollbackExecuted {
			t.Fatalf("Expected all tasks to be rolled back but %s was not.", taskRecorder.Name())
		}
	}

	skippedTask := ts[4].(*TaskRecorder)
	if skippedTask.RunExecuted {
		t.Fatalf("Expected task %s to be skipped but it was executed.", skippedTask.Name())
	}
}

func TestStopOnRollbackFailure(t *testing.T) {
	ts := []Task{
		NewTaskRecorder("task-1", true, true),
		NewTaskRecorder("task-2", true, false),
		NewTaskRecorder("task-3", true, true),
		NewTaskRecorder("task-4", false, true),
		NewTaskRecorder("task-5", true, true),
	}
	q := NewQueue(ts)

	err := q.Start()
	if err == nil {
		t.Fatal("Queue should have been executed with errors")
	}

	for _, task := range ts[1:4] {
		taskRecorder := task.(*TaskRecorder)
		if !taskRecorder.RollbackExecuted {
			t.Fatalf("Expected all tasks to be rolled back but %s was not.", taskRecorder.Name())
		}
	}

	skippedTask := ts[0].(*TaskRecorder)
	if skippedTask.RollbackExecuted {
		t.Fatalf("Expected task %s to be skipped but it was rolled back.", skippedTask.Name())
	}
}
