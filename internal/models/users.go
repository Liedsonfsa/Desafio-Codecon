package models

type Users struct {
	ID 		string 	`json:"id"`
	Name 	string 	`json:"nome"`
	Age 	uint64 	`json:"idade"`
	Score 	uint64 	`json:"score"`
	Active 	bool 	`json:"ativo"`
	Country string 	`json:"pais"`
	Team struct {
		Name 	string 	`json:"nome"`
		Leader 	bool 	`json:"lider"`
		Projects []struct{
			Name 		string 	`json:"nome"`
			Completed 	bool 	`json:"concluido"`
		} `json:"projetos"`
	} `json:"equipe"`
	Logs []struct{
		Date 	string 	`json:"data"`
		Action 	string 	`json:"acao"`
	}
}