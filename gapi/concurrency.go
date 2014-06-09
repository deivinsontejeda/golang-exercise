package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	GITHUB_ENDPOINT = "https://api.github.com/users"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Login string `json."login"`
	Name  string `json."name"`
}

// Perform a call to GITHUB API and parser the response
func GetJSON(u string) UserResponse {
	url := fmt.Sprintf("%s/%s", GITHUB_ENDPOINT, u)
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	b, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println("Error2", resp.Body)
	}
	i := &UserResponse{}
	e = json.Unmarshal(b, i)
	return *i

}

func GetProfile(users <-chan string, ur chan<- UserResponse) {
	for u := range users {
		ur <- GetJSON(u)
	}
}

func main() {
	// figure out how works size
	usersChan := make(chan string, 4)
	ur := make(chan UserResponse, 4)

	// taken args each one is a user
	args := os.Args[1:]

	// spanw up to  3 workers
	for w := 1; w <= 3; w++ {
		go GetProfile(usersChan, ur)
	}

	// Intereting:
	// Workers are blocking until we did not assign job to processing
	// in this case the `jobs` are store in `usersChan` channel

	//Send how many `job` as user we collected from CLI
	for _, u := range args {
		usersChan <- u
	}

	defer close(usersChan)

	// Collect all the results of work
	for f := 1; f <= len(args); f++ {
		fmt.Println(<-ur)
	}

}
