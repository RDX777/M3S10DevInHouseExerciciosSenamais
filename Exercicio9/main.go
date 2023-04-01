package main

import (
	"fmt"
)

type Task struct {
	Name string
}

type TaskList interface {
	addTask(task Task)
	listTasks()
	removeTask(index int)
}

type TodoList struct {
	tasks []Task
}

func (t *TodoList) addTask(task Task) {
	t.tasks = append(t.tasks, task)
	fmt.Println("Tarefa adicionada com sucesso!")
	t.listTasks()
}

func (t *TodoList) listTasks() {
	fmt.Println("Lista de Tarefas:")
	for i, task := range t.tasks {
		fmt.Printf("%d. %s\n", i+1, task.Name)
	}
}

func (t *TodoList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Tarefa removida com sucesso!")
	t.listTasks()
}

func main() {
	var option int
	todoList := &TodoList{}

	for {
		fmt.Println("1 - Adicionar Tarefa")
		fmt.Println("2 - Listar Tarefas")
		fmt.Println("3 - Remover Tarefa")
		fmt.Println("4 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var taskName string
			fmt.Print("Digite o nome da tarefa: ")
			fmt.Scan(&taskName)
			task := Task{Name: taskName}
			todoList.addTask(task)
		case 2:
			todoList.listTasks()
		case 3:
			var index int
			fmt.Print("Digite o índice da tarefa que deseja remover: ")
			fmt.Scan(&index)
			if index > 0 && index <= len(todoList.tasks) {
				todoList.removeTask(index - 1)
			} else {
				fmt.Println("Índice inválido!")
			}
		case 4:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
