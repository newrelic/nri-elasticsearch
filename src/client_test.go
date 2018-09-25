package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	hostname, port := "host", 9
	testCases := []struct {
		name    string
		useSSL  bool
		wantURL string
	}{
		{
			"No SSL",
			false,
			fmt.Sprintf("http://%s:%d", hostname, port),
		},
		{
			"SSL",
			true,
			fmt.Sprintf("https://%s:%d", hostname, port),
		},
	}

	for _, tc := range testCases {
		setupTestArgs()
		args.Hostname, args.Port, args.UseSSL = hostname, port, tc.useSSL
		client, err := NewClient(args.Hostname)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		} else {
			if client.baseURL != tc.wantURL {
				t.Errorf("Expected BaseURL '%s' got '%s'", tc.wantURL, client.baseURL)
			}
		}
	}
}

func TestBadCertFile(t *testing.T) {
	setupTestArgs()
	args.UseSSL = true
	args.CABundleDir = "thisdirectorydoesntexist"
	args.CABundleFile = "bad_file.nope"

	_, err := NewClient("")
	assert.Error(t, err)
}

func TestAuthRequest(t *testing.T) {
	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		username, password, ok := req.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, username, "testUser")
		assert.Equal(t, password, "testPass")
		res.Write([]byte("{\"ok\":true}"))
	}))
	defer func() { testServer.Close() }()

	client := &HTTPClient{
		client:   testServer.Client(),
		useAuth:  true,
		username: "testUser",
		password: "testPass",
		baseURL:  testServer.URL,
	}

	testResult := struct {
		OK *bool `json:"ok"`
	}{}

	err := client.Request("/endpoint", &testResult)
	assert.NoError(t, err)
	assert.Equal(t, true, *testResult.OK)
}

func TestBadStatusCode(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(401)
		res.Write([]byte("{\"error\":{\"type\":\"exception\",\"reason\":\"this is an error\"}}"))
	}))
	defer func() { testServer.Close() }()

	client := &HTTPClient{
		client:   testServer.Client(),
		useAuth:  true,
		username: "testUser",
		password: "testPass",
		baseURL:  testServer.URL,
	}

	err := client.Request("/endpoint", nil)
	assert.Error(t, err)
}
