package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)

	// for every `tick` that our `ticker`
	// emits, we print `tock`
	for t := range ticker.C {
		fmt.Println(t)
		fmt.Println(strings.Repeat("#", 25))
		postJsonPlaceHolder()
		randomNumber()
	}
}

func postJsonPlaceHolder() {
	randUserInt := rand.Intn(10)
	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "Jordan",
		"userId": randUserInt,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func randomNumber() {
	randIntWater, randIntWind, statusWater, statusWind := rand.Intn(100), rand.Intn(100), "aman", "aman"
	type fields struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	}

	if randIntWater >= 6 && randIntWater <= 8 {
		statusWater = "siaga"
	} else if randIntWater > 8 {
		statusWater = "bahaya"
	}

	if randIntWind >= 7 && randIntWind <= 15 {
		statusWind = "siaga"
	} else if randIntWind > 8 {
		statusWind = "bahaya"
	}
	jsonString := fields{Water: randIntWater, Wind: randIntWind}
	res, err := json.MarshalIndent(jsonString, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
	fmt.Printf("\n")
	fmt.Println("status water :", statusWater)
	fmt.Println("status wind :", statusWind)
}
