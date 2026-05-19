package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ngdangdat/compiler-in-go/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey programming language!\n", user.Username)
	fmt.Println("Type in commands!")
	repl.Start(os.Stdin, os.Stdout)
}
