package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    Version   = "dev"
    CommitSHA = "none"
    BuildTime = "unknown"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version information",
    Run: func(cmd *cobra.Command, args []string) {
        PrintVersion()
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}

func PrintVersion() {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Commit: %s\n", CommitSHA)
    fmt.Printf("BuildTime: %s\n", BuildTime)
}