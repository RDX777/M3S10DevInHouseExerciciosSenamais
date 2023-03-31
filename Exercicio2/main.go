package main

import "fmt"

type Address struct {
	street  string
	city    string
	state   string
	zipCode string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {

	person := Person{
		name: "João Silva",
		age:  35,
		address: Address{
			street:  "Rua das Flores",
			city:    "São Paulo",
			state:   "SP",
			zipCode: "01234-567",
		},
	}

	fmt.Println("Nome:", person.name)
	fmt.Println("Idade:", person.age)
	fmt.Println("Endereço:")
	fmt.Println("  Rua:", person.address.street)
	fmt.Println("  Cidade:", person.address.city)
	fmt.Println("  Estado:", person.address.state)
	fmt.Println("  CEP:", person.address.zipCode)
}
