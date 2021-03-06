package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func LogCommand() *cli.Command {
	return &cli.Command{
		Name:  "log",
		Usage: "View your entire goal log",
		Before: func(c *cli.Context) error {
			exists, err := IsDBExists()
			if !exists || err != nil {
				e := fmt.Errorf("You need to initialize the application using:\n $ gftd init\n")
				RED.Println(e)
				return e // TODO: Find a way to disable help text
			}
			fmt.Fprintf(c.App.Writer, "Fetching your goals\n") // TODO: Add a progress bar
			return nil
		},
		Action: func(c *cli.Context) error {
			if err := ViewGoals(); err != nil {
				RED.Println(err)
				return err
			}
			return nil
		},
	}
}

func ViewGoals() error {
	goals, err := ReadAllGoals()
	if err != nil {
		return err
	}
	if len(goals) == 0 {
		return fmt.Errorf("No goals have been added yet")
	}

	table := GetTableView(goals)
	fmt.Println(table)
	return nil
}
