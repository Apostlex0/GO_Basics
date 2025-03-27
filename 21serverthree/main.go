package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type BotRequest struct {
	Number int `json:"number"`
}

type BotResponse struct {
	Result int `json:"result"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/process", processRequest).Methods("POST")
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)

	fmt.Println("Master server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", corsMiddleware(r)))

}

func processRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	requestText := string(body)
	botType, number, err := parseRequest(requestText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var botURL string
	if strings.ToLower(botType) == "a" {
		botURL = "http://localhost:8001/square"
	} else if strings.ToLower(botType) == "b" {
		botURL = "http://localhost:8002/cube"
	} else {
		http.Error(w, "Invalid bot type", http.StatusBadRequest)
		return
	}

	botReq := BotRequest{Number: number}
	botReqJSON, err := json.Marshal(botReq)
	if err != nil {
		http.Error(w, "Error creating bot request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(botURL, "application/json", strings.NewReader(string(botReqJSON)))
	if err != nil {
		http.Error(w, "Error communicating with bot server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error ", http.StatusInternalServerError)
		return
	}

	var botResp BotResponse
	err = json.Unmarshal(respBody, &botResp)
	if err != nil {
		http.Error(w, "error ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(botResp)
}

func parseRequest(request string) (string, int, error) {
	re := regexp.MustCompile(`hi bot ([A|B]), the number is (\d+)`)
	matches := re.FindStringSubmatch(request)

	if len(matches) != 3 {
		return "", 0, fmt.Errorf("invalid ")
	}

	botType := matches[1]
	number, err := strconv.Atoi(matches[2])
	if err != nil {
		return "", 0, fmt.Errorf("invalid")
	}

	return botType, number, nil
}
