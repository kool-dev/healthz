package checks

import (
	"errors"
	"net"
	"net/http"
	"testing"
)

func TestTCP(t *testing.T) {
	t.Run("Test Error Handling", func(t *testing.T) {
		Checks.FuncDial = func(s, v string) (net.Conn, error) {
			return nil, errors.New("connection error")
		}
		err := checkSocket("tcp", "localhost:80")

		if err == nil {
			t.Error("expecting error 'connection error', got none")
		} else if err.Error() != "connection error" {
			t.Errorf("expecting error 'connection error', got %v", err)
		}
	})
}

func TestHTTP(t *testing.T) {
	t.Run("Test Successful Response", func(t *testing.T) {
		Checks.FuncDo = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       nil,
			}, nil
		}

		err := checkHTTP("http", "http://localhost:8080")

		if err != nil {
			t.Errorf("expecting no errors, got %s", err)
		}
	})

	t.Run("Test 500 Server Response", func(t *testing.T) {
		Checks.FuncDo = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 500,
				Body:       nil,
			}, nil
		}

		err := checkHTTP("http", "http://localhost:8080")

		if err == nil {
			t.Errorf("expecting error for 500 response, got none")
		}
	})

	t.Run("Test 500 Server Response", func(t *testing.T) {
		Checks.FuncDo = func(*http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("error on request")
		}

		err := checkHTTP("http", "http://localhost:8080")

		if err == nil {
			t.Error("expecting error 'error on request', got none")
		} else if err.Error() != "error on request" {
			t.Errorf("expecting error 'error on request', got %v", err)
		}
	})
}

func TestExec(t *testing.T) {
	t.Run("Valid Command Executes Successfully", func(t *testing.T) {
		testCmd := "ls -lah"

		err := checkExec("exec", testCmd)

		if err != nil {
			t.Errorf("expecting no errors, got %s", err)
		}
	})

	t.Run("Invalid Command Returns Error", func(t *testing.T) {
		testCmd := "invalid command"

		err := checkExec("exec", testCmd)

		if err == nil {
			t.Error("expecting errors, got none")
		}
	})
}
