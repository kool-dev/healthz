package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"encoding/json"
	"net"
	"net/http"

	"github.com/anmitsu/go-shlex"
	"github.com/spf13/cobra"
)

type toCheck struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value	  string `json:"value"`
}

var (
	jsonInput string

	rootCmd = &cobra.Command{
	Use:   "healthz",
	Short: "Healthz checks if your application is healthy!",
	Long: `A Fast simple and reliable tool to health check your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checks := make([]toCheck, 0)
		json.Unmarshal([]byte(jsonInput), &checks)

		for _, k := range checks {
			switch k.Type {
				case "tcp", "tcp4", "tcp6":
					checkSocket(k.Type, k.Value)
				case "http", "https":
					checkHTTP(k.Value)
				case "exec":
					checkExec(k.Value)
			}
		}
	},
}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init()  {
	rootCmd.PersistentFlags().StringVarP(&jsonInput, "input", "i", "", "Json input to perform health checks")
}

func checkSocket(socketType, url string)  {
	conn, err := net.Dial(socketType, url)
	if err != nil {
		fmt.Printf("Problem with dial: %v.\n", err.Error())
	}
	if conn != nil {
		fmt.Printf("Connected to %s://%s\n", socketType, url)
	}
}

func checkHTTP(url string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Problem with dial: %v.\n", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Problem with request: %s.\n", err.Error())
	} else if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Received %d from %s\n", resp.StatusCode, url)
	}
}

func checkExec(cmdValue string) {
	args, err := shlex.Split(cmdValue, true)
	if err != nil {
		fmt.Printf("Exec args parse failed with %s\n", err)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
}
