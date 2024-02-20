package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

type User struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func isJsonResponse(res *http.Response) bool {
	contentType := res.Header.Get("Content-Type")
	return res.StatusCode != 204 && regexp.MustCompile(`application\/json`).MatchString(contentType)
}

func handleResponse(res *http.Response) {
	defer res.Body.Close()

	if isJsonResponse(res) {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("API response as String " + string(body))

		var user User
		json.Unmarshal(body, &user)

		fmt.Printf("API response: %v\n", user)
	}
}

func get() {
	fmt.Println("Fetching from /todos")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	handleResponse(res)
}

func put() {
	fmt.Println("Put request to /todos")

	user := User{
		1,
		2,
		"Tweet, tweet, tweet!",
		false,
	}
	jsonReq, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	handleResponse(res)
}

func post() {
	fmt.Println("Posting to /todos")
	user := User{
		1,
		2,
		"How doth the little bird scream!",
		true,
	}
	req, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(req))
	if err != nil {
		log.Fatalln(err)
	}

	handleResponse(res)
}

func delete() {
	fmt.Println("Delete request to /todos")
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	handleResponse(res)
}

func main() {
	get()
	post()
	put()
	delete()
}
