package main

import (
 "encoding/json"
 "log"
 "net/http"

 "github.com/gorilla/mux"
)

type Request struct {
 Number int `json:"number"`
}

type Response struct {
 Result int `json:"result"`
}

func main() {
 r := mux.NewRouter()
 r.HandleFunc("/cube", handleCube).Methods("POST")
 log.Fatal(http.ListenAndServe(":8002", r))
}

func handleCube(w http.ResponseWriter, r *http.Request) {
 var req Request
 err := json.NewDecoder(r.Body).Decode(&req)
 if err != nil {
  http.Error(w, "Error ", http.StatusBadRequest)
  return
 }

 result := req.Number * req.Number * req.Number

 resp := Response{Result: result}
 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(resp)
}