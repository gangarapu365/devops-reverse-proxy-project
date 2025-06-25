package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        jsonResponse(w, map[string]string{
            "status":  "ok",
            "service": "2",
        })
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        jsonResponse(w, map[string]string{
            "message": "Hello from Service 2",
        })
    })

    log.Println("Service 2 listening on port 8000...")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

