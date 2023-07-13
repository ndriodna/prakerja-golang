package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Starwars struct {
	Name   string   `json:"name"`
	Height string   `json:"height"`
	Mass   string   `json:"mass"`
	Gender string   `json:"gender"`
	Films  []string `json:"films"`
}

func main() {
	res, _ := http.Get("https://swapi.dev/api/people/1")

	response, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	Luke := Starwars{}
	json.Unmarshal(response, &Luke)

	fmt.Println(Luke.Name)
	fmt.Println(Luke.Height)
	fmt.Println(Luke.Mass)
	fmt.Println(Luke.Gender)
	fmt.Println(Luke.Films[0])
}
