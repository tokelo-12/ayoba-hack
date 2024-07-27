package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Street     string `json:"street"`
	Region     string `json:"region"`
	Locality   string `json:"locality"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Company struct {
	NickName string  `json:"nickName"`
	Note     string  `json:"note"`
	Address  Address `json:"address"`
	Email    string  `json:"email"`
	URL      string  `json:"url"`
}

func createCard(c *gin.Context) {
	var myparam Company
	url := "https://api.ayoba.me/v1/business/card"

	access, err := Login()

	if err != nil {
		fmt.Print(err)
	}

	bearerToken := access.Token

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
