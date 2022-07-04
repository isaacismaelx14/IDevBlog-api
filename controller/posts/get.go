package postctrl

func Get() []Post {
	posts := []Post{
		{
			Id:       "1",
			Title:    "test",
			Body:     "this is a body test",
			ImageUrl: "https://via.placeholder.com/388x150",
		},
		{
			Id:       "2",
			Title:    "test2",
			Body:     "this is a body test2",
			ImageUrl: "https://via.placeholder.com/388x150",
		},
	}

	return posts
}
