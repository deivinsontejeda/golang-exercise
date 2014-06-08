package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GITHUB_ENDPOINT = "https://api.github.com/users"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Login string `json."login"`
	Name  string `json."name"`
}

type ErrorResponse struct {
	message           string `json:"message"`
	documentation_url string `json:"documentation_url"`
}

func GetProfile(username string, i interface{}) (e error) {
	url := fmt.Sprintf("%s/%s", GITHUB_ENDPOINT, username)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, e := ioutil.ReadAll(resp.Body)

	//does not work :/
	if resp.StatusCode != 200 {
		gr := &ErrorResponse{}
		json.Unmarshal(b, gr)
		fmt.Println(gr)
		return fmt.Errorf(gr.message)
	}
	if e != nil {
		fmt.Println("Error2", resp.Body)
		return e
	}
	e = json.Unmarshal(b, i)
	if e != nil {
		fmt.Println("Something wrong bad :(")
		return e
	}
	return
}

func main() {
	username := "deivinsontejeda"
	ur := &UserResponse{}
	GetProfile(username, ur)
	fmt.Println("-> ", ur)
}
