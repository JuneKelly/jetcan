package main

import (
	"github.com/codegangsta/cli"
	"fmt"
	"os"
)


func initApp() (app *cli.App) {
	app = cli.NewApp()

	app.Name = "jetcan"
	app.Version = "0.1.0"
	app.Action = func (c *cli.Context) {
		fmt.Println("this is jetcan")
	}
	return app
}

func main() {
	app := initApp()
	app.Run(os.Args)
}
