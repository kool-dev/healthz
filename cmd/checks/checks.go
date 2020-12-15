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

type dialType func(string, string) (net.Conn, error)
var dialFn dialType = net.Dial

// InitChecks runs all the checks specified in the Json input string
func InitChecks(input string) (t string, err error) {
	checks := make([]toCheck, 0)
	err = json.Unmarshal([]byte(input), &checks)
	if err != nil {
		t = "unmarshall"
		return
	}

	for _, k := range checks {
		switch k.Type {
			case "tcp", "tcp4", "tcp6":
				err = checkSocket(k.Type, k.Value)
				t = k.Type
			case "http", "https":
				err = checkHTTP(k.Type, k.Value)
				t = k.Type
			case "exec":
				err = checkExec(k.Type, k.Value)
				t = k.Type
		}

		if err != nil {
			return
		}
	}
	return
}

func checkSocket(t, u string) (err error) {
	conn, err := dialFn(t, u)

	if conn != nil {
		fmt.Printf("[%s check success] Connected to %s://%s\n", t, t, u)
	}

	return
}

func checkHTTP(t, u string) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[%s check success] Received %d from %s\n", t, resp.StatusCode, u)
	}

	return
}

func checkExec(t, c string) (err error) {
	args, err := shlex.Split(c, true)
	if err != nil {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	err = cmd.Run()
	if err != nil {
		return
	}

	fmt.Printf("[%s check success] cmd.Run() ran successfully\n", t)
	return
}
