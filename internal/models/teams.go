package models

type TeamInsights struct {
	Name 				string 		`json:"time"`
	TotalMembers 		uint64 		`json:"membros"`
	CompletedProjects 	uint64 		`json:"projetos_concluidos"`
	ActiveMembers 		float64 	`json:"membros_ativos"`
}