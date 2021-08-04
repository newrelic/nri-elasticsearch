package main

import (
	"bytes"
	"context"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		_, _ = res.Write([]byte("{\"ok\":true}"))
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
		_, _ = res.Write([]byte("{\"error\":{\"type\":\"exception\",\"reason\":\"this is an error\"}}"))
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

func TestClient_TLSUnsecureSkipVerify(t *testing.T) {
	setupTestArgs()

	srv := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintln(w, "Test server")
			assert.NoError(t, err)
		}))
	defer srv.Close()

	// Given test server is working
	req, err := http.NewRequestWithContext(context.Background(), "GET", srv.URL, nil)
	require.NoError(t, err)
	resp, err := srv.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Then create temp dir
	tmpDir, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer func() {
		err = os.RemoveAll(tmpDir)
		require.NoError(t, err)
	}()

	file := writeCApem(t, srv, tmpDir, "ca.pem")
	defer file.Close()

	args.UseSSL = true
	args.Timeout = 90

	client, err := NewClient(srv.URL)
	assert.NoError(t, err)
	assert.Equal(t, 90*time.Second, client.client.Timeout)

	// And http get should not work
	req, err = http.NewRequestWithContext(context.Background(), "GET", srv.URL, nil)
	require.NoError(t, err)
	resp, err = client.client.Do(req)
	require.Error(t, err)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	// Do not validate cert
	args.TLSInsecureSkipVerify = true

	clientB, err := NewClient(srv.URL)
	assert.NoError(t, err)
	assert.Equal(t, 90*time.Second, clientB.client.Timeout)

	// And http get should work
	req, err = http.NewRequestWithContext(context.Background(), "GET", srv.URL, nil)
	require.NoError(t, err)
	resp, err = clientB.client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
}

// Extract ca.pem from TLS server
func writeCApem(t *testing.T, srv *httptest.Server, tmpDir string, certName string) *os.File {
	caPEM := new(bytes.Buffer)
	err := pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: srv.Certificate().Raw,
	})
	require.NoError(t, err)

	// Then write the ca.pem to disk
	caPem, err := os.Create(filepath.Join(tmpDir, certName))
	require.NoError(t, err)
	_, err = caPem.Write(caPEM.Bytes())
	require.NoError(t, err)
	return caPem
}
