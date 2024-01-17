package main

import (
	"bufio"
	"classroomGo1/db"
	"classroomGo1/repository"
	"fmt"
	"os"
)

func main() {
	db.InitDB()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe seu nome: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Informe seu email: ")
	scanner.Scan()
	email := scanner.Text()

	user, err := repository.SaveUser(name, email)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
