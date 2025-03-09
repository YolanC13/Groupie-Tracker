package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	groupieTracker "groupieTracker/Internals"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
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
	IsPinned      bool
}

type Level struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PlayerID    string `json:"playerID"`
	Difficulty  string `json:"difficulty"`
	Downloads   int    `json:"downloads"`
	Likes       int    `json:"likes"`
	Length      string `json:"length"`
	SongName    string `json:"songName"`
}

type Pagination struct {
	Items         []User
	CurrentPage   int
	ItemsPerPage  int
	TotalItems    int
	TotalPages    int
	CurrentFilter string
}

var initialPage int = 1
var CurrentLeaderboardPage *int = &initialPage

func main() {
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
		data := struct {
			PinnedUsers   []string
			AllUsers      []User
			CurrentFilter string
		}{
			PinnedUsers:   InitializeFavorite(),
			AllUsers:      userList,
			CurrentFilter: r.URL.Query().Get("SearchFilter"),
		}
		if data.CurrentFilter == "" {
			data.CurrentFilter = "user"
		}
		temp.ExecuteTemplate(w, "searchMenu", data)
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
				if strings.Contains(errDecode.Error(), "unexpected end of JSON input") ||
					len(body) == 0 || string(body) == "-1" {
					w.WriteHeader(http.StatusNotFound)
					temp.ExecuteTemplate(w, "error", "User not found. Please check the username and try again.")
					return
				}
				return
			}

			userToFind.Username = strings.ToUpper(userToFind.Username)

			userToFind.IsPinned = CheckIsPinned(userToFind.Username)

			temp.ExecuteTemplate(w, "userInfo", userToFind)
		}
	})

	http.HandleFunc("/findLevel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			urlApi := "https://gdbrowser.com/api/level/" + r.FormValue("levelID")

			var levelToFind Level

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

			errDecode := json.Unmarshal(body, &levelToFind)
			if errDecode != nil {
				if strings.Contains(errDecode.Error(), "unexpected end of JSON input") ||
					len(body) == 0 || string(body) == "-1" {
					w.WriteHeader(http.StatusNotFound)
					temp.ExecuteTemplate(w, "error", "Level not found. Please check the ID and try again.")
					return
				}
				return
			}

			levelToFind.Name = strings.ToUpper(levelToFind.Name)
			temp.ExecuteTemplate(w, "levelInfo", levelToFind)
		}
	})

	http.HandleFunc("/pinUser", func(w http.ResponseWriter, r *http.Request) {
		PinUser(r.FormValue("username"))
		http.Redirect(w, r, "/searchMenu", http.StatusSeeOther)
	})

	http.HandleFunc("/unPinUser", func(w http.ResponseWriter, r *http.Request) {
		UnPinUser(r.FormValue("username"))
		http.Redirect(w, r, "/searchMenu", http.StatusSeeOther)
	})

	http.HandleFunc("/faqMenu", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "faqMenu", nil)
	})

	//Leaderboard

	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		page := CurrentLeaderboardPage
		if *page == 0 {
			*page = 1
		}

		filter := r.URL.Query().Get("filter")
		if filter == "" {
			filter = "stars"
		}

		sortedList := make([]User, len(userList))
		copy(sortedList, userList)

		switch filter {
		case "stars":
			sort.Slice(sortedList, func(i, j int) bool {
				return sortedList[i].Stars > sortedList[j].Stars
			})
		case "diamonds":
			sort.Slice(sortedList, func(i, j int) bool {
				return sortedList[i].Diamonds > sortedList[j].Diamonds
			})
		case "userCoins":
			sort.Slice(sortedList, func(i, j int) bool {
				return sortedList[i].UserCoins > sortedList[j].UserCoins
			})
		}

		for i := range sortedList {
			sortedList[i].Rank = i + 1
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
			Items:         sortedList[start:end],
			CurrentPage:   *page,
			ItemsPerPage:  itemsPerPage,
			TotalItems:    totalItems,
			TotalPages:    totalPages,
			CurrentFilter: filter,
		}

		temp.ExecuteTemplate(w, "leaderboard", pagination)
	})

	http.HandleFunc("/leaderboard/subtractPage", func(w http.ResponseWriter, r *http.Request) {
		(*CurrentLeaderboardPage)--
		filter := r.URL.Query().Get("filter")
		http.Redirect(w, r, "/leaderboard?filter="+filter, http.StatusSeeOther)
	})

	http.HandleFunc("/leaderboard/addPage", func(w http.ResponseWriter, r *http.Request) {
		(*CurrentLeaderboardPage)++
		filter := r.URL.Query().Get("filter")
		http.Redirect(w, r, "/leaderboard?filter="+filter, http.StatusSeeOther)
	})

	RunServer()
}

func RunServer() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./templates"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./templates/img"))))
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./templates/video"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./templates/font"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("./templates/scripts"))))

	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

func InitializeFavorite() []string {
	fichier := "favorite.txt"
	args := os.Args[1:]
	if len(args) == 1 {
		fichier = args[0]
	}

	var pinnedUsers []string

	pinnedUsers = append(pinnedUsers, groupieTracker.LoadFile(fichier)...)
	fmt.Println(pinnedUsers)
	return pinnedUsers
}

func PinUser(username string) {
	fichier := "favorite.txt"
	file, err := os.OpenFile(fichier, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(username + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func UnPinUser(username string) {
	fichier := "favorite.txt"
	file, err := os.Open(fichier)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		if scanner.Text() != username {
			lines = append(lines, scanner.Text())
		}
	}

	file, err = os.OpenFile(fichier, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}
}

func CheckIsPinned(username string) bool {
	fichier := "favorite.txt"
	file, err := os.Open(fichier)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == username {
			return true
		}
	}

	return false
}
