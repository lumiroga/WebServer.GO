package main

import (
	"fmt"
	"time"
)

func (tL *taskList) addToList(t *task) {

	tL.tasks = append(tL.tasks, t)
}
func (tL *taskList) removeTask(index int) {
	tL.tasks = append(tL.tasks[:index], tL.tasks[index+1:]...) //Error sin los elipsis
}

type task struct {
	name        string
	description string
	completed   bool
	date        time.Time
}
type taskList struct {
	tasks []*task
}

func (t *task) markCompleted() {
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {

	t := &task{
		name:        "Clase de Go",
		description: "Terminar el curso",
		date:        time.Now(),
	}
	t1 := &task{
		name:        "Clase de Python",
		description: "Terminar el curso",
		date:        time.Now(),
	}
	t2 := &task{
		name:        "Clase de React",
		description: "Terminar el curso",
		date:        time.Now(),
	}
	t3 := &task{
		name:        "Clase de Machine Learning",
		description: "Terminar el curso",
		date:        time.Now(),
	}

	list := taskList{
		tasks: []*task{
			t, t1, t2, t3,
		},
	}

	list.addToList(&task{name: "Nueva tarea"})

	for i := 0; i < len(list.tasks); i++ {
		fmt.Print(list.tasks[i])

	}

}
