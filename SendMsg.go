package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// {
// 	"msisdns": [
// 	  "+27220000000",
// 	  "+27220000001",
// 	  "+27220000002"
// 	],
// 	"message": {
// 	  "type": "text",
// 	  "text": "hello"
// 	}
//   }

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type RequestData struct {
	Msisdns []string `json:"msisdns"`
	Message Message  `json:"message"`
}

func sendMsg(c *gin.Context) {
	var myparam RequestData
	access, err := Login()

	if err != nil {
		fmt.Print(err)
	}

	bearerToken := access.Token
	url := "https://api.ayoba.me/v1/business/message"

	// data := RequestData{
	// 	Msisdns: []string{"+27823235496"},
	// 	Message: Message{
	// 		Type: "text",
	// 		Text: "Yah neh",
	// 	},
	// }

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("Error marshalling data:", err)
	// 	return
	// }

	if err := c.BindJSON(&myparam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(myparam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling data"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Handle the response here
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(string(body))

}
