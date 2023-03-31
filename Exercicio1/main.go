package main

import "fmt"

type Book struct {
	title         string
	author        string
	publishedYear int
	pages         int
}

func main() {
	// Inicializa uma variável do tipo Book com valores de exemplo
	book := Book{
		title:         "O Senhor dos Anéis",
		author:        "J.R.R. Tolkien",
		publishedYear: 1954,
		pages:         1178,
	}

	// Imprime os valores dos campos da variável book
	fmt.Println("Título:", book.title)
	fmt.Println("Autor:", book.author)
	fmt.Println("Ano de publicação:", book.publishedYear)
	fmt.Println("Número de páginas:", book.pages)
}
