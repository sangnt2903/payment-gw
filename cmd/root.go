package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "payment-gw",
    Short: "A payment gateway service",
    Long:  `A payment gateway service that handles payment processing using hexagonal architecture.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}