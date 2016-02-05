package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/onfoot/lumbergopher/logs"
)

func receiveLog(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Method != "PUT" {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "400 Bad request - incorrect content type", http.StatusBadRequest)
		return
	}

	var events logs.LogEvents
	body, readErr := ioutil.ReadAll(req.Body)

	if readErr != nil {
		log.Printf("Error: %v", readErr)

		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	jsonErr := json.Unmarshal(body, &events)

	if jsonErr != nil {
		log.Printf("Error: %v", jsonErr)

		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	for _, event := range events.Events {
		log.Printf("%v: [%s] %v", event.Time.Format(time.UnixDate), strings.ToUpper(event.Level.String()), event.Message)
	}
}

func main() {
	http.HandleFunc("/log", receiveLog)

	err := http.ListenAndServe(":8109", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
