package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/todo"
)

type typi struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetJson() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	var t typi

	json.Unmarshal([]byte(data), &t)
	fmt.Println(t.Title)

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	todo.Init()

}
