package articles

type InfoArticleFormatter struct {
	ID             int    `json:"id"`
	TutorID        int    `json:"tutor_id"`
	Kategori       string `json:"kategori"`
	Judul          string `json:"judul"`
	Konten         string `json:"konten"`
	TanggalPosting string `json:"tanggal_posting"`
}

func FormatArticle(article Article) *InfoArticleFormatter {
	formatter := InfoArticleFormatter{
		ID:             article.ID,
		TutorID:        article.TutorID,
		Kategori:       article.Kategori,
		Judul:          article.Judul,
		Konten:         article.Konten,
		TanggalPosting: article.TanggalPosting,
	}
	return &formatter
}

func FormatArticles(articles []Article) *[]InfoArticleFormatter {
	articlesFormat := []InfoArticleFormatter{}
	for _, article := range articles {
		articleFormat := FormatArticle(article)
		articlesFormat = append(articlesFormat, *articleFormat)
	}

	return &articlesFormat
}
