// @title Payment Gateway API
// @version 1.0
// @description Payment Gateway Service with Momo integration
// @host localhost:8080
// @BasePath /
package cmd

import (
	"fmt"
	"log"
	"payment-gw/internal/adapters/primary/http"
	"payment-gw/internal/adapters/secondary/momo"
	"payment-gw/internal/adapters/secondary/postgres"
	"payment-gw/internal/core/services"
	"payment-gw/pkg/conf"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the payment gateway server",
	Long:  `Start the payment gateway server that handles payment processing requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := conf.Load(); err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		log.Printf("Running in %s environment", conf.GetEnv())
		return runServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServer() error {
	log.Println("Starting payment gateway service...")

	// Initialize database
	log.Println("Connecting to database...")
	db, err := postgres.NewConnection(postgres.LoadConfig())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Initialize repositories and providers
	paymentRepo := postgres.NewPaymentRepository(db)
	momoProvider := momo.NewProvider()

	// Initialize services
	paymentService := services.NewPaymentService(paymentRepo, momoProvider)

	// Initialize HTTP server with dependencies
	server := http.NewServer(paymentService)

	port := conf.GetStringDefault("server", "port", "9000")
	return server.Run(fmt.Sprintf(":%s", port))
}
