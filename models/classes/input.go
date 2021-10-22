package classes

type CreateClassInput struct {
	TutorID   int    `form:"tutor_id" binding:"required"`
	Jenis     string `form:"jenis" binding:"required"`
	Topik     string `form:"topik" binding:"required"`
	Judul     string `form:"judul" binding:"required"`
	Jadwal    string `form:"jadwal" binding:"required"`
	LinkZoom  string `form:"link_zoom" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
}
