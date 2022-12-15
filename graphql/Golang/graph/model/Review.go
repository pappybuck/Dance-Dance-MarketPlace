package model

type Review struct {
	ID          string   `json:"id"`
	ProductID   string   `json:"productId"`
	UserID      string   `json:"userId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Rating      int      `json:"rating"`
	Product     *Product `json:"product"`
	User        *User    `json:"user"`
}
