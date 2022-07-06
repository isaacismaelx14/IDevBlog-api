package schemas

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement;index;size:10" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
