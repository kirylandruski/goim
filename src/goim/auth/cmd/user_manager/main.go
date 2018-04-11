package main

import (
	"flag"
	"log"
	"github.com/abiosoft/ishell"
)

var filePath = flag.String("path", "", "the path to json file containing users")

func initShell(manager *fileUserManager) *ishell.Shell {
	shell := ishell.New()

	shell.Println("Interactive shell to mange users")
	shell.Println("Use `help` to see available commands")

	shell.AddCmd(&ishell.Cmd{
		Name: "set",
		Help: "set user",
		Func: func(c *ishell.Context) {
			c.Print("Username: ")
			username := c.ReadLine()
			c.Print("Password: ")
			password := c.ReadPassword()

			err := manager.SetUser(username, password)
			if err != nil {
				log.Fatalf("error when performing operation: %v", err.Error())
			}

			if err := manager.write(); err != nil {
				log.Fatal(err)
			}
			shell.Println("Done")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Help: "list users",
		Func: func(c *ishell.Context) {
			usernames := manager.GetUsernames()
			for _, username := range usernames {
				shell.Println(username)
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "remove",
		Help: "remove user",
		Func: func(c *ishell.Context) {
			c.Print("Username: ")
			username := c.ReadLine()
			manager.RemoveUser(username)

			if err := manager.write(); err != nil {
				log.Fatal(err)
			}
			shell.Println("Done")
		},
	})

	return shell
}

func main() {
	flag.Parse()
	if len(*filePath) == 0 {
		log.Fatalf("not all required flags were passed. use -help to see the list of flags")
	}

	manager := NewFileUserManager(*filePath)
	if err := manager.read(); err != nil {
		log.Fatal(err)
	}
	shell := initShell(manager)
	shell.Run()
}
