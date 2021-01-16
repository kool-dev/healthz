package cmd

import (
	"fmt"
	"kool-dev/healthz/cmd/checks"
	"os"

	"github.com/spf13/cobra"
)

var version string = "0.0.0-dev"

var (
	jsonInput string

	rootCmd = &cobra.Command{
		Use:     "healthz",
		Version: version,
		Short:   "Healthz checks if your application is healthy!",
		Long:    `A Fast simple and reliable tool to health check your application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if jsonInput == "" {
				fmt.Println("missing --input parameter")
				os.Exit(2)
			}

			t, err := checks.InitChecks(jsonInput)

			if err != nil {
				fmt.Printf("[%s check error] %s\n", t, err)
				os.Exit(1)
			}

			fmt.Println("\nAll checks ran successfully.\nApplication is healthy.")
			os.Exit(0)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&jsonInput, "input", "i", "", "Json input to perform health checks")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
