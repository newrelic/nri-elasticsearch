package main

import (
	"fmt"
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
		client, err := NewClient("")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		} else {
			if client.baseURL != tc.wantURL {
				t.Errorf("Expected BaseURL '%s' got '%s'", tc.wantURL, client.baseURL)
			}
		}
	}
}

func TestHostnameOverride(t *testing.T) {
	hostname, port, ssl := "host", 9, false
	hostOverride := "overridden"
	expectedURL := fmt.Sprintf("http://%s:%d", hostOverride, port)

	setupTestArgs()
	args.Hostname, args.Port, args.UseSSL = hostname, port, ssl
	client, err := NewClient(hostOverride)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	} else {
		if client.baseURL != expectedURL {
			t.Errorf("Expected BaseURL '%s' got '%s'", expectedURL, client.baseURL)
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
