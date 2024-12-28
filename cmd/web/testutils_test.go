package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leoromanini/medication_api/internal/models/mocks"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		medications: &mocks.MedicationModel{},
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) request(t *testing.T, method string, urlPath string, inputBody io.Reader) (int, http.Header, string) {

	req, err := http.NewRequest(method, ts.URL+urlPath, inputBody)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := ts.Client()
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)

	return resp.StatusCode, resp.Header, string(body)
}
