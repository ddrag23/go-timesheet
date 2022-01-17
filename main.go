package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type users struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type committer struct {
	DateCommit time.Time `json:"date"`
}
type commit struct {
	Committer committer `json:"committer"`
	Message   string    `json:"message"`
}

type responseCommit struct {
	Commit commit `json:"commit"`
}

var (
	baseUrl                        string = "https://api.github.com"
	input, sinceDate, untilDate    string
	year, day, untilYear, untilDay int
	month, untilMonth              time.Month
	resCommit                      []responseCommit
)

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
	fmt.Println("Masukkan tanggal : \nExample : 2022-2-1")
	fmt.Scanf("%s", &sinceDate)
	dateSince := strings.Split(sinceDate, "-")
	year = StrToInt(dateSince[0])
	day = StrToInt(dateSince[2])
	month = time.Month(StrToInt(dateSince[1]))
	fmt.Println("Masukkan batas tanggal : \nExample : 2022-3-1")
	fmt.Scanf("%s", &untilDate)
	dateUntil := strings.Split(untilDate, "-")
	untilYear = StrToInt(dateUntil[0])
	untilDay = StrToInt(dateUntil[2])
	untilMonth = time.Month(StrToInt(dateUntil[1]))
	parseSince := time.Date(year, time.Month(month), day-1, 00, 00, 00, 0, time.UTC).Format(time.RFC3339)
	parseUntil := time.Date(untilYear, untilMonth, untilDay-1, 00, 00, 00, 0, time.Local).Format(time.RFC3339)
	// tryTime := time.Date(2022, time.Now().Month(), , 00, 00, 00, 0, time.UTC).Format(time.RFC3339)
	// fmt.Println(tryTime)
	fmt.Println(parseSince)
	repo := requestApiGithub(config, baseUrl+"/repos/"+user.Login+"/"+input+"/commits?since="+parseSince+"&until="+parseUntil)

	err = json.Unmarshal(repo, &resCommit)
	// fmt.Println(string(repo))
	PanicIfError(err)
	for _, res := range resCommit {
		fmt.Println(res.Commit.Message)
		fmt.Println(res.Commit.Committer.DateCommit)
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
