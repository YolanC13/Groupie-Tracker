package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

type User struct {
	Username      string `json:"username"`
	Stars         int    `json:"stars"`
	Moons         int    `json:"moons"`
	Demons        int    `json:"demons"`
	Coins         int    `json:"coins"`
	UserCoins     int    `json:"userCoins"`
	Diamonds      int    `json:"diamonds"`
	CreatorPoints int    `json:"cp"`
	Rank          int    `json:"rank"`
}

func main() {
	// Récupération du leaderboard
	urlApi := "https://gdbrowser.com/api/leaderboard"

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("Error: %v\n", errReq)
		return
	}

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Printf("Error: %v\n", errRes)
		return
	}
	defer res.Body.Close()

	body, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		fmt.Printf("Error reading response body: %v\n", errRead)
		return
	}

	var userList []User
	errDecode := json.Unmarshal(body, &userList)
	if errDecode != nil {
		fmt.Printf("Error decoding JSON: %v\n", errDecode)
		return
	}

	//Debug du Leaderboard
	/*for i, user := range userList {
		fmt.Printf("%d. %s\n", i+1, user.Username)
		fmt.Printf("Stars: %d\n", user.Stars)
		fmt.Printf("Moons: %d\n", user.Moons)
		fmt.Printf("Demons: %d\n", user.Demons)
		fmt.Printf("Coins: %d\n", user.Coins)
		fmt.Printf("User coins: %d\n", user.UserCoins)
		fmt.Printf("Diamonds: %d\n", user.Diamonds)
		fmt.Printf("Creator points: %d\n", user.CreatorPoints)
		fmt.Println("-------------------------------------------------")
	}*/

	temp, errTemp := template.ParseGlob("templates/*.html")
	if errTemp != nil {
		fmt.Printf("Error parsing template: %v\n", errTemp)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/mainMenu", http.StatusSeeOther)
	})

	http.HandleFunc("/mainMenu", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "mainMenu", userList)
	})

	http.HandleFunc("/searchMenu", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "searchMenu", userList)
	})

	http.HandleFunc("/findUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			urlApi := "https://gdbrowser.com/api/profile/" + r.FormValue("username")

			var userToFind User

			req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
			if errReq != nil {
				fmt.Printf("Error: %v\n", errReq)
				return
			}

			res, errRes := httpClient.Do(req)
			if errRes != nil {
				fmt.Printf("Error: %v\n", errRes)
				return
			}
			defer res.Body.Close()

			body, errRead := io.ReadAll(res.Body)
			if errRead != nil {
				fmt.Printf("Error reading response body: %v\n", errRead)
				return
			}

			errDecode := json.Unmarshal(body, &userToFind)
			if errDecode != nil {
				fmt.Printf("Error decoding JSON: %v\n", errDecode)
				return
			}

			temp.ExecuteTemplate(w, "userInfo", userToFind)
		}
	})

	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "leaderboard", userList)
	})

	// Serveur
	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
