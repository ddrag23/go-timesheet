package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

var (
	baseUrl   string = "https://api.github.com"
	input     string
	year, day int
	month     time.Month
	resCommit []responseCommit
)

func main() {
	// date := "2022-01-12"
	// parse, _ := time.Parse("2022-01-12T00:00:00Z", date)

	config := ReadConfig()
	body := requestApiGithub(config, baseUrl+"/user")
	user := users{}
	err := json.Unmarshal(body, &user)
	PanicIfError(err)
	fmt.Println("Welcome To Generate Timesheet " + user.Name)
	PanicIfError(err)
	fmt.Println("Masukkan repo :")
	fmt.Scanf("%s", &input)
	fmt.Println("Masukkan Tahun :")
	fmt.Scanf("%d", &year)
	fmt.Println("Masukkan Bulan :")
	fmt.Scanf("%d", &month)
	fmt.Println("Masukkan Tanggal :")
	fmt.Scanf("%d", &day)
	parse := time.Date(year, month, day, 24, 00, 00, 0, time.UTC).Format("2022-01-12T00:00:00Z")
	fmt.Println(parse)
	repo := requestApiGithub(config, baseUrl+"/repos/"+user.Login+"/"+input+"/commits?since="+parse)

	err = json.Unmarshal(repo, &resCommit)
	PanicIfError(err)
	for _, res := range resCommit {
		fmt.Println(res.Commit.Message)
	}
}

func requestApiGithub(config Config, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token "+config.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept", "vnd.github.v3+json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
