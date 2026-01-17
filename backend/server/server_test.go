package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/yhiraki/wakeonlan-webapp/backend/config"
)

// MockWoLService is a mock implementation of wol.Service
type MockWoLService struct {
	WakeFunc func(macAddr string) error
}

func (m *MockWoLService) Wake(macAddr string) error {
	if m.WakeFunc != nil {
		return m.WakeFunc(macAddr)
	}
	return nil
}

func TestHandleGetTargets(t *testing.T) {
	targets := []config.Target{
		{Name: "PC1", MAC: "00:11:22:33:44:55"},
		{Name: "PC2", MAC: "AA:BB:CC:DD:EE:FF"},
	}

	s := NewServer(targets, &MockWoLService{})

	req := httptest.NewRequest(http.MethodGet, "/api/targets", nil)
	w := httptest.NewRecorder()

	s.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GET /api/targets status = %d, want %d", w.Code, http.StatusOK)
	}

	var got []config.Target
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !reflect.DeepEqual(got, targets) {
		t.Errorf("GET /api/targets body = %v, want %v", got, targets)
	}
}

func TestHandleWake(t *testing.T) {
	tests := []struct {
		name       string
		body       map[string]string
		mockWake   func(string) error
		wantStatus int
	}{
		{
			name:       "success",
			body:       map[string]string{"mac": "00:11:22:33:44:55"},
			mockWake:   func(mac string) error { return nil },
			wantStatus: http.StatusOK,
		},
		{
			name:       "wol service error",
			body:       map[string]string{"mac": "00:11:22:33:44:55"},
			mockWake:   func(mac string) error { return errors.New("udp error") },
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "invalid method",
			body:       nil,
			mockWake:   nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockWoLService{WakeFunc: tt.mockWake}
			s := NewServer(nil, mockService)

			var req *http.Request
			if tt.body != nil {
				jsonBody, _ := json.Marshal(tt.body)
				req = httptest.NewRequest(http.MethodPost, "/api/wake", bytes.NewBuffer(jsonBody))
			} else {
				// Simulating method not allowed by hitting the endpoint with wrong method is tricky with just handler func,
				// but here we are testing the multiplexer routing effectively if we use Server.ServeHTTP
				// Note: If we test specific handler functions, we might need a different approach.
				// For this test, let's assume we are testing the ServeHTTP routing behavior too for "invalid method" case,
				// or just testing bad request body.
				// Let's change "invalid method" to "bad request body" for handler testing simplicity if we target the handler directly,
				// but here we use s.ServeHTTP so we can test routing.
				req = httptest.NewRequest(http.MethodGet, "/api/wake", nil)
			}

			w := httptest.NewRecorder()
			s.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("POST /api/wake status = %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
