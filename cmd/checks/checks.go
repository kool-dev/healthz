package checks

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/exec"

	"github.com/anmitsu/go-shlex"
)

type toCheck struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type checker struct {
	FuncDo   func(*http.Request) (*http.Response, error)
	FuncDial func(string, string) (net.Conn, error)
}

//Checks uses the default methods for running checks
var Checks checker

func init() {
	Checks = checker{
		http.DefaultClient.Do,
		net.Dial,
	}
}

// InitChecks runs all the checks specified in the Json input string
func InitChecks(input string) (t string, err error) {
	c := make([]toCheck, 0)
	err = json.Unmarshal([]byte(input), &c)
	if err != nil {
		t = "unmarshall"
		return
	}

	for _, k := range c {
		switch k.Type {
		case "tcp", "tcp4", "tcp6":
			err = checkSocket(k.Name, k.Type, k.Value)
			t = k.Type
		case "http", "https":
			err = checkHTTP(k.Name, k.Type, k.Value)
			t = k.Type
		case "exec":
			err = checkExec(k.Name, k.Type, k.Value)
			t = k.Type
		}

		if err != nil {
			return
		}
	}
	return
}

func checkSocket(n, t, u string) (err error) {
	conn, err := Checks.FuncDial(t, u)

	if conn != nil {
		fmt.Printf("[\"%s\" check success] Connected to \"%s://%s\"\n", n, t, u)
	}

	return
}

func checkHTTP(n, t, u string) (err error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}

	resp, err := Checks.FuncDo(req)
	if err == nil && resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		fmt.Printf("[\"%s\" check success] Received %d from \"%s\"\n", n, resp.StatusCode, u)
	} else if err == nil {
		fmt.Printf("[\"%s\" check server error] Received %d from \"%s\"\n", n, resp.StatusCode, u)
		err = errors.New("Got response < 200 or >= 400 on Http Check")
	}

	return
}

func checkExec(n, t, cs string) (err error) {
	args, err := shlex.Split(cs, true)
	if err != nil {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	err = cmd.Run()
	if err != nil {
		return
	}

	fmt.Printf("[\"%s\" check success] No errors from \"%s\"\n", n, cs)
	return
}
