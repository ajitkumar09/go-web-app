package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		method      string
		path        string
		wantStatus  int
		wantContent string
		wantType    string
	}{
		{method: http.MethodGet, path: "/", wantStatus: http.StatusOK, wantContent: "Learn DevOps from Basics", wantType: "text/html; charset=utf-8"},
		{method: http.MethodGet, path: "/home", wantStatus: http.StatusOK, wantContent: "Learn DevOps from Basics", wantType: "text/html; charset=utf-8"},
		{method: http.MethodGet, path: "/courses", wantStatus: http.StatusOK, wantContent: "Main Tutorials", wantType: "text/html; charset=utf-8"},
		{method: http.MethodGet, path: "/about", wantStatus: http.StatusOK, wantContent: "About the Author", wantType: "text/html; charset=utf-8"},
		{method: http.MethodGet, path: "/contact", wantStatus: http.StatusOK, wantContent: "Contact Us", wantType: "text/html; charset=utf-8"},
		{method: http.MethodGet, path: "/healthz", wantStatus: http.StatusOK, wantContent: "ok", wantType: "text/plain; charset=utf-8"},
		{method: http.MethodGet, path: "/home/unknown", wantStatus: http.StatusNotFound},
		{method: http.MethodGet, path: "/missing", wantStatus: http.StatusNotFound},
		{method: http.MethodPost, path: "/courses", wantStatus: http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			newRouter().ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Fatalf("status = %v, want %v", rr.Code, tt.wantStatus)
			}
			if tt.wantContent != "" && !strings.Contains(rr.Body.String(), tt.wantContent) {
				t.Errorf("response body does not contain %q", tt.wantContent)
			}
			if tt.wantType != "" && rr.Header().Get("Content-Type") != tt.wantType {
				t.Errorf("Content-Type = %q, want %q", rr.Header().Get("Content-Type"), tt.wantType)
			}
		})
	}
}
