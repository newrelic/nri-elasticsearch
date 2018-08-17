package main

import (
	"fmt"
	"testing"
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
		client, err := NewClient(nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		} else {
			if client.BaseURL != tc.wantURL {
				t.Errorf("Expected BaseURL '%s' got '%s'", tc.wantURL, client.BaseURL)
			}
		}
	}
}
