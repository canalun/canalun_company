package repository

import "content-updater/domain/model"

type ArticleRepository interface {
	GetArticles() []model.Article
}

type ContentUpdater interface {
	UpdateContent()
}
