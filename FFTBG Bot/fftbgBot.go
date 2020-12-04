package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	bot "TwitchBot"
)

type Unit struct {
	Name          string `json:"name"`
	Gender        string
	Sign          string
	Brave         int
	Faith         int
	Class         string
	ActionSkill   string
	ReactionSkill string
	SupportSkill  string
	MoveSkill     string
	Mainhand      string
	Offhand       string
	Head          string
	Armor         string
	Accessorry    string
	ClassSkills   []string
	ExtraSkills   []string
}

type Team struct {
	Player string
	Name   string
	Units  []Unit
}

type TeamList struct {
	KeyBlack    Team `json:"black"`
	KeyBlue     Team `json:"blue"`
	KeyBrown    Team `json:"brown"`
	KeyChampion Team `json:"champion"`
	KeyGreen    Team `json:"green"`
	KeyPurple   Team `json:"purple"`
	KeyRed      Team `json:"red"`
	KeyWhite    Team `json:"white"`
	KeyYellow   Team `json:"yellow"`
}

type thisTournament struct {
	Type string   `json:"Type"`
	Key  TeamList `json:"Teams"`
	Maps []string
}

type latestTournament struct {
	ID int `json:"ID"`
}

func getCurrentTournament() thisTournament {
	var l []latestTournament
	var t thisTournament

	url := "https://fftbg.com/api/tournaments?limit=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &l)

	if err != nil {
		panic(err)
	}

	//fmt.Println(res)
	//fmt.Println(string(body))

	//fmt.Printf("%v\n", l[0].ID)

	url2 := "https://fftbg.com/api/tournament/" + strconv.Itoa(l[0].ID)

	req2, _ := http.NewRequest("GET", url2, nil)

	req2.Header.Add("cache-control", "no-cache")

	res2, err := http.DefaultClient.Do(req2)

	if err != nil {
		panic(err)
	}

	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res2.Body)

	err = json.Unmarshal(body2, &t)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v\n", t)

	return t
}

func main() {
	//t := getCurrentTournament()
	//fmt.Printf("%v\n", t.Key.KeyBlack.Units[0].Class)
	//fmt.Printf("%v\n", t.Key.KeyBlack.Units[1].Class)
	//fmt.Printf("%v\n", t.Key.KeyBlack.Units[2].Class)
	//fmt.Printf("%v\n", t.Key.KeyBlack.Units[3].Class)

	myBot := bot.BasicBot{
		Channel: "riley331",
		MsgRate: (20 / 30) * time.Second,
		Name: "RileyBBot",
		Port: "6667",
		PrivatePath: "oauth.json",
		Server: "irc.chat.twitch.tv",
	}

	//var msgRegex *regexp.Regexp = regexp.MustCompile(`^:(\w+)!\w+@\w+\.tmi\.twitch\.tv (PRIVMSG) #\w+(?: :(.*))?$`)
	//var cmdRegex *regexp.Regexp = regexp.MustCompile(`^!(\w+)\s?(\w+)?`)

	myBot.Start()

}
