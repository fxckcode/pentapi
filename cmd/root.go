package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Global flag variables
var (
	verbose bool
	output  string
	proxy   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pentapi",
	Short: "PentAPI - Automated API security testing CLI tool",
	Long: `PentAPI is a fast and extensible command-line tool written in Go
for scanning and testing REST APIs for security vulnerabilities.
Use it to discover endpoints, test authorization, and fuzz parameters.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Temporary run behavior to show flag values
		fmt.Println("PentAPI CLI is running")
		fmt.Println("Verbose:", verbose)
		fmt.Println("Output:", output)
		fmt.Println("Proxy:", proxy)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "results.json", "Output file name")
	rootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "Proxy address (e.g. http://127.0.0.1:8080)")

	// Help message for the root command
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	PentAPI - Automated API Security Testing Tool

	Usage:
	pentapi [command]

	Available Commands:
	help        Show this help message
	scan        Scan an API for vulnerabilities (coming soon)

	Flags:
	-v, --verbose     Enable verbose output
	-o, --output      Output file (default "results.json")
	-p, --proxy       Proxy to use for requests (e.g. http://127.0.0.1:8080)

	Examples:
	pentapi scan https://api.example.com
	pentapi scan -v -p http://127.0.0.1:8080 https://api.example.com
		`)
	})
}
