package checks

import (
	"testing"
	"net"
	"errors"
)

func TestTCPErrorOutput(t *testing.T)  {
	originalDialFn := dialFn

	dialFn = func(s, v string) (net.Conn, error) {
		return nil, errors.New("connection error")
	}

	defer func() {
		dialFn = originalDialFn
	}()

	err := checkSocket("tcp", "localhost:80")

	if err == nil {
		t.Error("expecting error 'connection error', got none")
	} else if err.Error() != "connection error" {
		t.Errorf("expecting error 'connection error', got %v", err)
	}
}

func TestHTTP(t *testing.T)  {
	//@todo add logic here...
}

func TestExec(t *testing.T)  {
	//@todo add logic here...
}
