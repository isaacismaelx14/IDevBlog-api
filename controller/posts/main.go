package postctrl

type Post struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"image_url"`
}
