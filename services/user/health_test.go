package main

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHealth(t *testing.T) {
  mux := http.NewServeMux()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(Health{Service: "user", Status: "ok"})
  })

  req := httptest.NewRequest(http.MethodGet, "/health", nil)
  rr := httptest.NewRecorder()
  mux.ServeHTTP(rr, req)

  if rr.Code != http.StatusOK {
    t.Fatalf("expected 200, got %d", rr.Code)
  }
  var h Health
  if err := json.Unmarshal(rr.Body.Bytes(), &h); err != nil {
    t.Fatalf("invalid json: %v", err)
  }
  if h.Status != "ok" || h.Service != "user" {
    t.Fatalf("unexpected payload: %+v", h)
  }
}
