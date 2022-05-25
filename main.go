package main

import "fmt"

type task struct {
	name        string
	description string
	done        bool
}

var taskList = []task{}

func addTask(task task) {
	taskList = append(taskList, task)
}

func deleteTask(index int) {
	taskList = append(taskList[:index], taskList[index+1:]...)
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
	t1 := task{
		name:        "Task 2",
		description: "This is task 2",
		/*done:        false,*/ //es redundante ya que el ZERO VALUE de un BOOLEAN es FALSE
	}
	//fmt.Printf("%+v\n", t)
	t.doneTask()
	t.updateName("finish")
	t.updateDescription("finish task")
	//fmt.Printf("%+v\n", t)
	fmt.Println(taskList)
	addTask(t)
	fmt.Printf("%+v\n", taskList[0])
	addTask(t1)
	fmt.Printf("%+v\n", taskList)
	deleteTask(0)
	fmt.Printf("%+v\n", taskList)

}
