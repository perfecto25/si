package main

import (
	"fmt"
	"log"
	"os"

	"github.com/perfecto25/si2/filters"
	"github.com/urfave/cli"
)

func initData() map[string]map[string]string {
	// creates overall Data map that includes all filters

	data := make(map[string]map[string]string)
	data["operating system"] = map[string]string{}
	data["network"] = map[string]string{}

	return data
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{

		&cli.BoolFlag{
			Name:  "json",
			Usage: "json parsing",
		},
	}

	app.Action = func(c *cli.Context) error {

		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.Bool("json") == true {
			fmt.Println("JSON", name)

		} else {

			data := initData()
			//fmt.Println(data)

			fmt.Println(filters.ReturnOS2(data))

		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
