package articles

type CreateArticleInput struct {
	TutorID        int    `json:"tutor_id" binding:"required"`
	Kategori       string `json:"kategori"`
	Judul          string `json:"judul"`
	Konten         string `json:"konten"`
	TanggalPosting string `json:"tanggal_posting"`
}

type GetArticleInput struct {
	ID int `uri:"id" binding:"required"`
}
