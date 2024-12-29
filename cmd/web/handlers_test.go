package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedWantBody = `{"ID":1,"Name":"Ibuprofen","Dosage":"400 mg","Form":"Capsule"`

func TestMedicationGet(t *testing.T) {
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/v1/medications/1",
			wantCode: http.StatusOK,
			wantBody: mockedWantBody,
		},
		{
			name:     "Valid ID with trailing slash",
			urlPath:  "/v1/medications/1/",
			wantCode: http.StatusOK,
			wantBody: mockedWantBody,
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/v1/medications/10",
			wantCode: http.StatusNotFound,
			wantBody: "{\"error\": \"Medication ID Not Found\"}",
		},
		{
			name:     "Negative ID",
			urlPath:  "/v1/medications/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/v1/medications/1.2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/v1/medications/foo",
			wantCode: http.StatusNotFound,
		},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, http.MethodGet, tt.urlPath, strings.NewReader(""))
			assert.Equal(t, tt.wantCode, code)
			assert.Contains(t, body, tt.wantBody)
		})

	}
}

func TestMedicationDelete(t *testing.T) {
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid DELETE",
			urlPath:  "/v1/medications/1",
			wantCode: http.StatusOK,
			wantBody: mockedWantBody,
		},
		{
			name:     "Valid DELETE with trailing slash",
			urlPath:  "/v1/medications/1/",
			wantCode: http.StatusOK,
			wantBody: mockedWantBody,
		},
		{
			name:     "Non-existing DELETE",
			urlPath:  "/v1/medications/10",
			wantCode: http.StatusNotFound,
		},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, http.MethodDelete, tt.urlPath, strings.NewReader(""))
			assert.Equal(t, tt.wantCode, code)
			assert.Contains(t, body, tt.wantBody)
		})

	}
}

func TestMedicationList(t *testing.T) {
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid Request",
			urlPath:  "/v1/medications",
			wantCode: http.StatusOK,
			wantBody: "[" + mockedWantBody,
		},
		{
			name:     "Valid Request with trailing slash",
			urlPath:  "/v1/medications/",
			wantCode: http.StatusOK,
			wantBody: "[" + mockedWantBody,
		},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, http.MethodGet, tt.urlPath, strings.NewReader(""))
			assert.Equal(t, tt.wantCode, code)
			assert.Contains(t, body, tt.wantBody)

		})

	}
}

func TestMedicationCreate(t *testing.T) {
	tests := []struct {
		name      string
		urlPath   string
		wantCode  int
		inputBody string
		wantBody  string
	}{
		{
			name:      "Valid POST",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusCreated,
			inputBody: `{"name": "valid", "dosage": "valid", "form": "valid"}`,
			wantBody:  mockedWantBody,
		},
		{
			name:      "Invalid name",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"name": ["invalid"], "dosage": "valid", "form": "valid"}`,
			wantBody:  `{"error": "Name of type string"}`,
		},
		{
			name:      "Invalid dosage",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"name": "valid", "dosage": true, "form": "valid"}`,
			wantBody:  `{"error": "Dosage of type string"}`,
		},
		{
			name:      "Invalid form",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"name": "valid", "dosage": "valid", "form": 100}`,
			wantBody:  `{"error": "Form of type string"}`,
		},
		{
			name:      "Empty body",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusBadRequest,
			inputBody: `{}`,
		},
		{
			name:      "Required name validation",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"dosage": "valid", "form": "valid"}`,
		},
		{
			name:      "Exceed name validation",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "` + strings.Repeat("a", 101) + `"}`,
			wantBody:  "Name cannot exceed 100 characters",
		},
		{
			name:      "Exceed dosage validation",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "foo", "dosage": "` + strings.Repeat("a", 21) + `"}`,
			wantBody:  "Name cannot exceed 20 characters",
		},
		{
			name:      "Exceed form validation",
			urlPath:   "/v1/medications",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "foo", "form": "` + strings.Repeat("a", 21) + `"}`,
			wantBody:  "Name cannot exceed 20 characters",
		},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, http.MethodPost, tt.urlPath, strings.NewReader(tt.inputBody))
			assert.Equal(t, tt.wantCode, code)
			assert.Contains(t, body, tt.wantBody)

		})

	}
}

func TestMedicationPatch(t *testing.T) {
	tests := []struct {
		name      string
		urlPath   string
		wantCode  int
		inputBody string
		wantBody  string
	}{
		{
			name:      "Valid PATCH All fields",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusOK,
			inputBody: `{"name": "name-foo", "dosage": "dosage-foo", "form": "form-foo"}`,
			wantBody:  `{"ID":1,"Name":"name-foo","Dosage":"dosage-foo","Form":"form-foo`,
		},
		{
			name:      "Valid PATCH name",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusOK,
			inputBody: `{"name": "name-bar"}`,
			wantBody:  `{"ID":1,"Name":"name-bar","Dosage":"dosage-foo","Form":"form-foo`,
		},
		{
			name:      "Valid PATCH dosage",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusOK,
			inputBody: `{"dosage": "dosage-bar"}`,
			wantBody:  `{"ID":1,"Name":"name-bar","Dosage":"dosage-bar","Form":"form-foo`,
		},
		{
			name:      "Valid PATCH form",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusOK,
			inputBody: `{"form": "form-bar"}`,
			wantBody:  `{"ID":1,"Name":"name-bar","Dosage":"dosage-bar","Form":"form-bar`,
		},
		{
			name:      "Invaid PATCH atribute",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"form": ["invalid"]}`,
		},
		{
			name:      "Required name validation",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": ""}`,
		},
		{
			name:      "Exceed name validation",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "` + strings.Repeat("a", 101) + `"}`,
			wantBody:  "Name cannot exceed 100 characters",
		},
		{
			name:      "Exceed dosage validation",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "foo", "dosage": "` + strings.Repeat("a", 21) + `"}`,
			wantBody:  "Name cannot exceed 20 characters",
		},
		{
			name:      "Exceed form validation",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusUnprocessableEntity,
			inputBody: `{"name": "foo", "form": "` + strings.Repeat("a", 21) + `"}`,
			wantBody:  "Name cannot exceed 20 characters",
		},
		{
			name:      "Invalid name",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"name": 100}`,
			wantBody:  `{"error": "Name of type string"}`,
		},
		{
			name:      "Invalid dosage",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"dosage": 100}`,
			wantBody:  `{"error": "Dosage of type string"}`,
		},
		{
			name:      "Invalid form",
			urlPath:   "/v1/medications/1",
			wantCode:  http.StatusBadRequest,
			inputBody: `{"form": 100}`,
			wantBody:  `{"error": "Form of type string"}`,
		},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, http.MethodPatch, tt.urlPath, strings.NewReader(tt.inputBody))
			assert.Equal(t, tt.wantCode, code)
			assert.Contains(t, body, tt.wantBody)
		})

	}
}

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	code, _, body := ts.request(t, http.MethodGet, "/health", strings.NewReader(""))
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, body, `{"status": "ok"}`)
}

func TestHomePage(t *testing.T) {
	t.Parallel()

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	code, _, _ := ts.request(t, http.MethodGet, "/", strings.NewReader(""))
	assert.Equal(t, http.StatusOK, code)
}
