package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type token struct {
	Token  string `json:"access_token"`
	Expire int64  `json:"expire"`
}

const (
	username = "9487e8d6912cef120ae9cd5cd1038f56806f9956"
	password = "IAGIT92PhSRwasDX5q3cKn2KBCrKnlA"
	url      = "https://api.ayoba.me/v2/login"
)

func Login() (token, error) {

	data := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Print("error marshalling data ", err)
		return token{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return token{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return token{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return token{}, err
	}

	fmt.Println(string(body))

	var Access token

	error := json.Unmarshal(body, &Access)
	if error != nil {
		return token{}, error
	}

	return Access, nil
}

// func x(){
// 	y, err := Login()

// 	if err != nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	y
// }
