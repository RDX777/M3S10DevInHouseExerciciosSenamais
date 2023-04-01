package main

import (
	"fmt"
	"time"
)

type Task struct {
	name string
}

func addTask(tasks *[]Task, task Task) {
	*tasks = append(*tasks, task)
}

func removeTask(tasks *[]Task, index int) {
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}

func printTasks(tasks []Task) {
	fmt.Println("Tarefas:")
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i+1, task.name)
	}
}

func taskCounter(tasks []Task, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Printf("Quantidade de tarefas: %d\n", len(tasks))
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	var tasks []Task
	quit := make(chan bool)

	go taskCounter(tasks, quit)

	for {
		var choice int
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1) Adicionar tarefa")
		fmt.Println("2) Listar tarefas")
		fmt.Println("3) Remover tarefa")
		fmt.Println("4) Sair")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Digite o nome da tarefa: ")
			fmt.Scanln(&name)
			task := Task{name}
			addTask(&tasks, task)
			fmt.Println("Tarefa adicionada com sucesso!")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("Não há tarefas na lista.")
			} else {
				printTasks(tasks)
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("Não há tarefas na lista.")
			} else {
				printTasks(tasks)
				var index int
				fmt.Print("Digite o número da tarefa a ser removida: ")
				fmt.Scanln(&index)
				index -= 1
				if index < 0 || index >= len(tasks) {
					fmt.Println("Número inválido.")
				} else {
					removeTask(&tasks, index)
					fmt.Println("Tarefa removida com sucesso!")
				}
			}
		case 4:
			quit <- true
			fmt.Println("Encerrando programa...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
