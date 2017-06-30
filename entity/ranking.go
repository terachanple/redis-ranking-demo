package entity

type Ranking struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

type Rankings []Ranking
