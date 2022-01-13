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

type commit struct {
	Message string `json:"message"`
}

type responseCommit struct {
	Commit commit `json:"commit"`
}

var baseUrl string = "https://api.github.com"
var input string

func main() {

	config := ReadConfig()
	body := requestApiGithub(config, baseUrl+"/user")
	user := users{}
	err := json.Unmarshal(body, &user)
	PanicIfError(err)
	fmt.Println("Welcome To Generate Timesheet " + user.Name)
	PanicIfError(err)
	fmt.Println("Masukkan repo :")
	fmt.Scanf("%s", &input)
	repo := requestApiGithub(config, baseUrl+"/repos/"+user.Login+"/"+input+"/commits")
	commit := []responseCommit{}
	err = json.Unmarshal(repo, &commit)
	PanicIfError(err)
	fmt.Println(commit)
}

func requestApiGithub(config Config, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token "+config.AccessToken)
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
