package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "1.0.0"
	name    = "test-tool"
)

var (
	isVersion bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:           name,
		Run:           runRootCommand,
		SilenceErrors: false,
		SilenceUsage:  false,
	}
	rootCmd.Flags().BoolVar(&isVersion, "version", false, "show version")
	rootCmd.Execute()
}

func runRootCommand(rootCmd *cobra.Command, args []string) {
	//rootCmd.Flags().BoolVar(&isVersion, "version", false, "show version")
	//fmt.Println(isVersion)
	if isVersion {
		fmt.Println(version)
		return
	}
}
