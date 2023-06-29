package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Greeting struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/greetings", handleGreetings)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGreetings(w http.ResponseWriter, r *http.Request) {
    greeting := Greeting{
        Message: "Hello, Go!",
    }

    jsonBytes, err := json.Marshal(greeting)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}
