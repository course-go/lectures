package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// START OMIT

func main() {
	registerCommand := &cobra.Command{
		Use:   "register [email] [password]",
		Short: "Register new user",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("invalid number of arguments provided")
			}
			if len(args[1]) < 6 {
				return errors.New("password not long enough")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// ...
		},
	}
	echoCommand := &cobra.Command{
		Use:   "print [string]",
		Short: "Print arguments to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	// MIDDLE OMIT

	var count int
	timesCommand := &cobra.Command{
		Use:   "times [string]",
		Short: "Print arguments to stdout count times",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for range count {
				fmt.Println("Print: " + strings.Join(args, " "))
			}
		},
	}
	timesCommand.Flags().IntVarP(&count, "count", "c", 1, "times to echo the input")

	echoCommand.AddCommand(timesCommand)

	rootCommand := &cobra.Command{Use: "my-cli"}
	rootCommand.AddCommand(registerCommand, echoCommand)
	_ = rootCommand.Execute()
}

// END OMIT
