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

type Pagination struct {
	Items        []User
	CurrentPage  int
	ItemsPerPage int
	TotalItems   int
	TotalPages   int
}

var initialPage int = 1
var CurrentLeaderboardPage *int = &initialPage

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
		page := CurrentLeaderboardPage
		if *page == 0 {
			*page = 1
		}

		itemsPerPage := 5
		totalItems := 100
		totalPages := 20

		if *page > totalPages {
			*page = totalPages
		}

		start := (*page - 1) * itemsPerPage
		end := start + itemsPerPage
		if end > totalItems {
			end = totalItems
		}

		pagination := &Pagination{
			Items:        userList[start:end],
			CurrentPage:  *page,
			ItemsPerPage: itemsPerPage,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		}

		temp.ExecuteTemplate(w, "leaderboard", pagination)
	})

	http.HandleFunc("/leaderboard/subtractPage", func(w http.ResponseWriter, r *http.Request) {
		*CurrentLeaderboardPage--
		http.Redirect(w, r, "/leaderboard", http.StatusSeeOther)
	})

	http.HandleFunc("/leaderboard/addPage", func(w http.ResponseWriter, r *http.Request) {
		*CurrentLeaderboardPage++
		http.Redirect(w, r, "/leaderboard", http.StatusSeeOther)
	})

	RunServer()
}

func RunServer() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./templates"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./templates/img"))))
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./templates/video"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./templates/font"))))

	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
