package tutors

type RegisterUserInput struct {
	Nama       string `json:"nama" binding:"required"`
	Email      string `json:"email" binding:"required,email" gorm:"unique"`
	Password   string `json:"password" binding:"required"`
	RegisterAs string `json:"register_as" binding:"required"`
}

type LogisUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	LoginAs  string `json:"login_as" binding:"required"`
}

type CreateTutorInput struct {
	Nama          string `json:"nama" binding:"required"`
	MasaKerja     int    `json:"masa_kerja" binding:"required"`
	SitusWeb      string `json:"situs_web" binding:"required"`
	Kompetensi    string `json:"kompetensi" binding:"required"`
	Pekerjaan     string `json:"pekerjaan" binding:"required"`
	TopikDiminati string `json:"topik_diminati" binding:"required"`
}

type GetTutorInput struct {
	ID int `uri:"id" binding:"required"`
}
