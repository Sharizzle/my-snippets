package code

type Code struct {
	ID   int `gorm:"primary_key" json:"id"`
	Date string `json:"date"`
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Content string `json:"content"`
}
