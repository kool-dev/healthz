package checks

import (
	"os/exec"
	"net"
	"net/http"
	"github.com/anmitsu/go-shlex"
	"fmt"
	"encoding/json"

)
type toCheck struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value	  string `json:"value"`
}

// InitChecks runs all the checks specified in the Json input string
func InitChecks(input string) (err error) {
	checks := make([]toCheck, 0)
	err = json.Unmarshal([]byte(input), &checks)
	if err != nil {
		return
	}

	for _, k := range checks {
		switch k.Type {
			case "tcp", "tcp4", "tcp6":
				err = checkSocket(k.Type, k.Value)
			case "http", "https":
				err = checkHTTP(k.Value)
			case "exec":
				err = checkExec(k.Value)
		}

		if err != nil {
			return
		}
	}
	return
}

func checkSocket(socketType, url string) (err error) {
	conn, err := net.Dial(socketType, url)

	if conn != nil {
		fmt.Printf("[TCP Check Success] Connected to %s://%s\n", socketType, url)
	}

	return
}

func checkHTTP(url string) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("[HTTP Request Creation Error] Problem with dial: %v.\n", err.Error())
	}

	resp, err := client.Do(req)
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[HTTP Success] Received %d from %s\n", resp.StatusCode, url)
	}

	return
}

func checkExec(cmdValue string) (err error) {
	args, err := shlex.Split(cmdValue, true)
	if err != nil {
		fmt.Printf("[Exec Parse Error] Exec args parse failed with %s\n", err)
	}

	cmd := exec.Command(args[0], args[1:]...)

	err = cmd.Run()
	if err != nil {
		fmt.Printf("[Exec Run Error] cmd.Run() failed with %s\n", err)
		return
	}

	fmt.Println("[Exec Run Success] cmd.Run() ran successfully")
	return
}
