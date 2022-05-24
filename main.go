package main

import "fmt"

type task struct {
	name        string
	description string
	done        bool
}

func (t *task) doneTask() {
	t.done = true
}

func (t *task) updateDescription(newDescription string) {
	t.description = newDescription
}

func (t *task) updateName(newName string) {
	t.name = newName
}

func main() {
	t := task{
		name:        "Task 1",
		description: "This is task 1",
		/*done:        false,*/ //es redundante ya que el ZERO VALUE de un BOOLEAN es FALSE
	}
	fmt.Printf("%+v\n", t)
	t.doneTask()
	t.updateName("finish")
	t.updateDescription("finish task")
	fmt.Printf("%+v\n", t)
}
