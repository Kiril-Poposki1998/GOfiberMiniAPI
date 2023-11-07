package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"testing"
	"time"
)

var login_info = []byte(`{"username": "kiril", "password": "kire123"}`)

var jar, err = cookiejar.New(nil)

var Url = "http://localhost:8080"

var client = http.Client{
	Timeout: time.Second * 5,
	Jar:     jar,
}

func TestRegister(t *testing.T) {
	resp, err := client.Post(Url+"/auth/register", "application/json", bytes.NewReader(login_info))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	} else if resp.StatusCode != 200 {
		fmt.Println("Got bad status from server ", resp.Body)
		os.Exit(-1)
	}
}

func TestAuthenticate(t *testing.T) {
	resp, err := client.Post(Url+"/auth/login", "application/json", bytes.NewReader(login_info))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	} else if resp.StatusCode != 200 {
		fmt.Println("Got bad status from server")
		os.Exit(-1)
	}
}
