package classes

type CreateClassInput struct {
	TutorID   int    `json:"tutor_id" binding:"required"`
	Jenis     string `json:"jenis"`
	Topik     string `json:"topik" binding:"required"`
	Judul     string `json:"judul" binding:"required"`
	Jadwal    string `json:"jadwal" binding:"required"`
	LinkZoom  string `json:"link_zoom" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Modul     string `json:"modul" binding:"required"`
}
