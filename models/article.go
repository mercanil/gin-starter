package models

import "github.com/pkg/errors"

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetArticles() (articles []Article) {
	articles = []Article{
		Article{
			Id:      0,
			Title:   "Mat",
			Content: "Bildigin Math",
		},
		Article{
			Id:      1,
			Title:   "Turkce",
			Content: "Bildigin Turkce",
		},
	}
	return
}

func GetArticleById(id int) (article *Article, err error) {
	articles := GetArticles()
	for _, article := range articles {
		if article.Id == id {
			return &article, nil
		}
	}
	return nil, errors.New("Not Found")

}
