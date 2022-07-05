package schemas

type Post struct {
	Id       uint   `gorm:"primaryKey;autoIncrement;index;size:10" json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"image_url"`
}
