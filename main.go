package main

import (
	"github.com/mingrammer/go-todo-rest-api-example/app"
	"github.com/mingrammer/go-todo-rest-api-example/config"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/foobar", d1, 0644)
	check(err)
}
