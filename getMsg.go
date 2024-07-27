package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// [
//   {
//     "msisdn": "+27220000000",
//     "message": {
//       "id": "i_66EE5680-5F25-42D2-AC45-72CBDF427747",
//       "type": "text",
//       "text": "hello"
//     }
//   }
// ]

type MessageBlock struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type msgRes struct {
	Msisdn       string       `json:"msisdn"`
	MessageBlock MessageBlock `json:"message"`
}

func getText() {
	url := "https://api.ayoba.me/v1/business/message"

	access, err := Login()

	if err != nil {
		fmt.Print(err)
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Print("Error fetching text : ", err)
		return
	}

	bearerToken := access.Token

	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response SStatus:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	// fmt.Println(string(body))

	var data []msgRes

	error := json.Unmarshal(body, &data)
	if error != nil {
		fmt.Print(error)
		return
	}

	fmt.Println("data: ", data)

}
