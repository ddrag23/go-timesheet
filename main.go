package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type users struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func main() {

	fmt.Println("Hello word")
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token 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")
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	user := users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
