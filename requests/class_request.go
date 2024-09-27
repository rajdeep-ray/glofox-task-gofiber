package requests

type ClassRequest struct{
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Capacity  int    `json:"capacity"`
}