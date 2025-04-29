package models

type Users struct {
	ID 		string 	`json:"id"`
	Name 	string 	`json:"name"`
	Age 	uint64 	`json:"age"`
	Score 	uint64 	`json:"score"`
	Active 	bool 	`json:"active"`
	Country string 	`json:"country"`
	Team struct {
		Name string 	`json:"name"`
		Leader bool 	`json:"leader"`
		Projects []struct{
			Name string 	`json:"name"`
			Completed bool 	`json:"completed"`
		} `json:"projects"`
	} `json:"team"`
	Logs []struct{
		Date string 	`json:"date"`
		Action string 	`json:"action"`
	}
}