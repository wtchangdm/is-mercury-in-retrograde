package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wtchangdm/is-mercury-in-retrograde/internal/mercury"
)

type responsePayload struct {
	IsMercuryInRetrograde bool `json:"isMercuryInRetrograde"`
}

type errorPayload struct {
	Error string `json:"error"`
}

func retrogradeHandler(w http.ResponseWriter, r *http.Request) {
	var payload interface{}
	m := mercury.New()
	ok, err := m.Retrograde()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		payload = &errorPayload{
			Error: fmt.Sprintf("Error fetching status: %v", err),
		}
	} else {
		payload = &responsePayload{
			IsMercuryInRetrograde: ok,
		}
	}

	response, _ := json.Marshal(payload)
	w.Write(response)
}

func Serve() {
	http.HandleFunc("/isMercuryInRetrograde", retrogradeHandler)
	fmt.Println("Try API at http://localhost:8080/isMercuryInRetrograde.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
