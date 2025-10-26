package main

import (
  "encoding/json"
  "log"
  "net/http"
  "os"
  "time"
)

var start = time.Now()

type Health struct {
  Service string        `json:"service"`
  Status  string        `json:"status"`
  Uptime  time.Duration `json:"uptime"`
  Version string        `json:"version"`
}

func main() {
  svc := env("SERVICE_NAME", "user")
  port := env("PORT", "8081")
  version := env("VERSION", "0.1.0")

  mux := http.NewServeMux()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(Health{
      Service: svc,
      Status:  "ok",
      Uptime:  time.Since(start).Truncate(time.Millisecond),
      Version: version,
    })
  })

  srv := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
  log.Printf("%s service starting on :%s", svc, port)
  if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    log.Fatalf("server error: %v", err)
  }
}

func env(k, def string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  return def
}
