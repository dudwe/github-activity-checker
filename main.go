package main

import (
	"context"
	"fmt"
	"git-activity-checker/activity"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	// Declare a command
	cmd := &cli.Command{
		Name: "GithubActivityChecker",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "user", Usage: "get activity of user"},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			user := cmd.String("user")
			if user == "" {
				fmt.Fprintln(os.Stderr, "Error: --user flag is required")
				os.Exit(1)
			}
			resp, err := activity.GetActivity(cmd.String("user"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unhandled error: %[1]v\n", err)
				os.Exit(86)
			}
			fmt.Println(resp)
			return nil
		},
		Authors: []any{
			"Bhavin Shah: https://github.com/dudwe",
		},
	}

	// Simulate the command line arguments

	cmd.Run(context.Background(), os.Args)
}
