package entity

type ToDo struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	CreatedBy string `json:"createdBy"`
}
