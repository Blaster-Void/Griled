package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	g "github.com/AllenDang/giu"
)

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int  `json:"id"`
	Safe bool `json:"safe"`
}

var px string
var jokera string
var sint int32
var mint int32
var hint int32

func joker() {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)

	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)

	}
	defer resp.Body.Close()

	var joke JokeResponse
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	if joke.Error {
		fmt.Println("API returned an error")
		return
	}
	if joke.Type == "twopart" {
		fmt.Printf("Setup: %s\n", joke.Setup)
		fmt.Printf("Delivery: %s\n", joke.Delivery)

		jokera = fmt.Sprintf("JOKE:%s\n:%s", joke.Setup, joke.Delivery)

	}
}
func genjoke() {
	go joker()
}

func onImSoCute() {
	fmt.Println("")
	px = "CUTE"
	sint, mint, hint = 0, 0, 0
}

func callingtime(h, m, s int) {

	for {
		time.Sleep(time.Second * 1)
		s--
		if s < 0 {
			s = 59
			m--
		}
		if m < 0 {
			go joker()
			m = 59
			h--
		}
		if h < 0 {
			h = 23
		}
		px = fmt.Sprintf("%20v:%20v:%20v", h, m, s)
		if s == 0 && m == 0 && h == 0 {
			fmt.Println("timeup")

			break
		}

	}
}
func onClickMe() {

	go callingtime(int(hint), int(mint), int(sint))

}

func loop() {
	g.SingleWindow().Layout(

		g.Row(
			g.InputInt(&sint).Label("Sec").Size(100),
			g.Button("Click 2 start").OnClick(onClickMe),
			g.Button("ReSet All").OnClick(onImSoCute),
		),
		g.Row(
			g.InputInt(&mint).Size(100),
		),
		g.Row(
			g.InputInt(&hint).Size(100),
		),
		g.Button("Genrate Joke").OnClick(genjoke),
		g.Label(px),
		g.Label(jokera),
	)

}

func main() {
	wnd := g.NewMasterWindow("Time is Joke?", 400, 200, g.MasterWindowFlagsMaximized)
	wnd.Run(loop)
}
