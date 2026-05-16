package main

import (
	"compiler-in-go/monkey/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey language!\n", user.Name)
	fmt.Printf("Type in command!\n")
	repl.Start(os.Stdin, os.Stdout)
}
