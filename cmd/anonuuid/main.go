package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/moul/anonuuid"
)

// main is the entrypoint
func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul"
	app.Version = "1.0.0-dev"
	app.Usage = "Anonymize UUIDs outputs"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "hexspeak, d",
			Usage: "Generate hexspeak style fake UUIDs",
		},
	}

	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	anonuuid := anonuuid.New()

	if c.Bool("hexspeak") {
		anonuuid.Hexspeak = true
	}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(anonuuid.Sanitize(line))
	}
}
