package usecase

import "content-updater/domain/repository"

type ArticleUsecase struct {
	ArticleRepositories []repository.ArticleRepository
}

func NewArticleUsecase(
	ars ...repository.ArticleRepository,
) ArticleUsecase {
	return ArticleUsecase{
		ArticleRepositories: ars,
	}
}

func (a ArticleUsecase) UpdateArticleList(NumOfEntry int) error {
	for _, articleRepository := range a.ArticleRepositories {
		articles := articleRepository.GetArticles()
		for i := 0; i < NumOfEntry; i++ {
			articles[i].Title
		}
	}
	//記事のタイトルとURLを取得
	//JSONとして出力

	return nil
}
