package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// getBanPrice queries the Kalium Appditto API and returns the balance_decimal as a float64
func getBanPrice(account string) (float64, error) {
	url := "https://kaliumapi.appditto.com/api"
	payload := []byte(fmt.Sprintf(`{"action":"account_balance","account":"%s"}`, account))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response: %v", err)
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Parse balance_decimal to float64
	balanceStr, ok := data["balance_decimal"]
	if !ok {
		return 0, fmt.Errorf("balance_decimal not found in response")
	}

	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting to float64: %v", err)
	}

	return balance, nil
}

func main() {
	account := "ban_18hkxaj65y3o8bwb3u4j78j5tatc36orfzhbwage8jdn8yrrzq7hdjx8ihiq"
	balance, err := getBanPrice(account)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Balance: %.2f BAN\n", balance)
}

