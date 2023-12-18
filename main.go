package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pig_road_server/model"
)

func main() {
	http.HandleFunc("/restaurants", getRestaurants)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	log.Println("getRestaurants", r.Method, r.URL)
	restaurants := model.GetKoreanRestaurants()
	responseBody := make(map[string]interface{})
	responseBody["restaurants"] = restaurants

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(responseBody); err != nil {
		log.Println("Error in encoding restaurants:", err)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(responseBody)
	}
}
