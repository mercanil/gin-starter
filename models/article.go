package models

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetArticles() (articles []Article) {
	articles = []Article{
		Article{
			Id:      1,
			Title:   "Mat",
			Content: "Bildigin Math",
		},
		Article{
			Id:      2,
			Title:   "Turkce",
			Content: "Bildigin Turkce",
		},
	}
	return
}
