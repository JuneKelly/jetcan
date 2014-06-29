package main

import (
	"github.com/codegangsta/cli"
	"github.com/ShaneKilkelly/jetcan/localstorage"
	"fmt"
	"os"
)

func handler(c *cli.Context) {
	fmt.Println("this is jetcan")
}

func initApp() (app *cli.App, err error) {
	app = cli.NewApp()

	app.Name = "jetcan"
	app.Version = "0.1.0"
	app.Action = handler

	err = localstorage.Initialize()
	if err != nil {
		return nil, err
	}
	return app, nil
}

func main() {
	app , err := initApp()
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		app.Run(os.Args)
	}
}
