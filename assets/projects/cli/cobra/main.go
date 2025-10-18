package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// START OMIT

func main() {
	cmdPrint := &cobra.Command{
		Short: "register [username] [password]",
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
	cmdEcho := &cobra.Command{
		Use:   "print [string]",
		Short: "Print arguments to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	// MIDDLE OMIT

	var count int
	cmdTimes := &cobra.Command{
		Use:   "times [string]",
		Short: "Print arguments to stdout count times",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for range count {
				fmt.Println("Print: " + strings.Join(args, " "))
			}
		},
	}
	cmdTimes.Flags().IntVarP(&count, "count", "c", 1, "times to echo the input")

	cmdEcho.AddCommand(cmdTimes)
	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	_ = rootCmd.Execute()
}

// END OMIT
