package code

type Code struct {
	ID   uint `json:"id" gorm:"primaryKey"`
	Date string `json:"date"`
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Content string `json:"content"`
}
