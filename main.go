package main

import (
	"github.com/mingrammer/go-todo-rest-api-example/app"
	"github.com/mingrammer/go-todo-rest-api-example/config"
	"bufio"
	"fmt"
	"os"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/foobar", d1, 0644)
}
