package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"math/rand"
	"strconv"
	"time"
)

func getDelay(params url.Values) time.Duration {
	delay := params.Get("delay")

	if len(delay) == 0 {
		return 0
	}

	value, err := strconv.ParseUint(delay, 10, 32)

	if err != nil {
		log.Print(err)
		return 0
	}

	return time.Duration(value)
}

type Site struct {
	Id int32 `json:"id"`
	Participants int32 `json:"participants"`
	Incidents int32 `json:"incidents"`
}

func generateSite() *Site {
	return &Site{
		Id: rand.Int31(),
		Participants: rand.Int31(),
		Incidents: rand.Int31(),
	}
}

type Response struct {
	Sites []*Site `json:"sites"`
}

func generateResponse() *Response {
	var response Response

	count := rand.Int31n(30)

	response.Sites = make([]*Site, count)

	for i := int32(0); i < count; i++ {
		response.Sites[i] = generateSite()
	}

	return &response
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	delay := getDelay(params)
	w.Header().Set("X-Server", "Awesome Test Server 0.1")

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	time.Sleep(delay * time.Second)

	json, err := json.Marshal(generateResponse())

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/sites", handler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
