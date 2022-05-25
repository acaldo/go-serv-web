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

func iterateTask(list []task) {
	finish := []task{}
	dont := []task{}
	for _, task := range list {
		if task.done == true {
			finish = append(finish, task)
		} else if task.done == false {
			dont = append(dont, task)
		}
	}
	fmt.Printf("Finish: %+v\n", finish)
	fmt.Printf("Dont: %+v\n", dont)
}

func (t *task) doneTask() {
	t.done = !t.done
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
		done:        false, //es redundante ya que el ZERO VALUE de un BOOLEAN es FALSE
	}
	t1 := task{
		name:        "Task 2",
		description: "This is task 2",
		done:        true,
	}
	t2 := task{
		name:        "Task 2",
		description: "This is task 2",
		done:        false,
	}
	//fmt.Printf("%+v\n", t)
	//t.doneTask()
	t.updateName("finish")
	t.updateDescription("finish task")
	//fmt.Printf("%+v\n", t)
	fmt.Println(taskList)
	addTask(t)
	addTask(t1)
	addTask(t2)
	//for
	/* 	for i := 0; i < len(taskList); i++ {
	   		fmt.Printf("%d: %+v\n", i, taskList[i])
	   	}
	   	//for range
	   	for i, task := range taskList {
	   		fmt.Printf("%d: %+v\n", i, task)
	   	} */
	iterateTask(taskList)

}
