package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"

	"github.com/Liedsonfsa/Desafio-Codecon/internal/models"
)

func SaveUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Create("users.json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetSuperUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	file, err := os.Open("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var superusers []models.Users
	for _, user := range users {
		if user.Score >= 900 && user.Active {
			superusers = append(superusers, user)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&superusers)
}

func GetTopCountries(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	file, err := os.Open("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	topCountries := map[string]int{}
	
	for _, user := range users {
		if user.Score >= 900 && user.Active {
			topCountries[user.Country]++
		}
	}

	countries := map[int]string{}
	values := make([]int, 0, len(topCountries))
	for key, value := range topCountries {
		values = append(values, value)
		countries[value] = key
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	type data struct {
		Country 	string `json:"country"`
		SuperUsers 	uint64 `json:"superusers"`
	}

	values = values[:5]

	response := map[int]data{}
	for id, value := range values {
		response[id + 1] = data{Country: countries[value], SuperUsers: uint64(value)}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

func GetTeamInsights(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	file, err := os.Open("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	leaders := map[string]string{}
	times := map[string]int{}
	activeUsers := map[string]int{}
	completedprojects := map[string][]string{}
	
	for _, user := range users {
		times[user.Team.Name]++
		if user.Active {
			activeUsers[user.Team.Name]++
		}

		if user.Team.Leader {
			leaders[user.Team.Name] = user.Name
		}

		projects := []string{}

		for _, project := range user.Team.Projects {
			if project.Completed {
				projects = append(projects, project.Name)
			}
		}
		completedprojects[user.Team.Name] = projects
	}

	type data struct {
		Members 			int 		`json:"members"`
		ActiveMembers		int 		`json:"active_members"`
		Leader				string		`json:"leader"`
		PercentActiveUsers	float64 	`json:"percent_active_users"`
		CompletedProjects	[]string	`json:"completed_projects"`
	}

	response := map[string]data{}
	for team, membros := range times {
		response[team] = data{
			Members: membros,
			ActiveMembers: activeUsers[team],
			Leader: leaders[team],
			PercentActiveUsers: ((float64(activeUsers[team])) / (float64(times[team]) / 100)),
			CompletedProjects: completedprojects[team],
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

func GetLogins(w http.ResponseWriter, r *http.Request){
	var users []models.Users

	file, err := os.Open("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logins := map[string]int{}

	for _, user := range users {
		for _, log := range user.Logs {
			if log.Action == "login" {
				logins[log.Date]++
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&logins)
}