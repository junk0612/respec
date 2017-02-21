package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ReSpec"
	app.Usage = "Repeats RSpec testing multiple times"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		// コマンド設定
		{
			Name:   "run",
			Usage:  "Run the tests written with RSpec.",
			Action: helloAction,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "times, t",
					Usage: "set the number of repetitions",
					Value: 3,
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		err := exec.Command("bundle", "exec", "rspec", "-v").Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Please install rspec.")
			return err
		}
		return nil
	}

	app.Run(os.Args)
}

func helloAction(c *cli.Context) {
	repeats := c.Int("times")

	paramFirst := ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First() // c.Args()[0] と同じ意味
	}

	fmt.Printf("Hello world! %s\n", paramFirst)

	for i := 0; i < repeats; i++ {
		out, _ := exec.Command("bundle", "exec", "rspec").Output()
		fmt.Printf("period %d:", i+1)
		fmt.Println(string(out))
	}
}
