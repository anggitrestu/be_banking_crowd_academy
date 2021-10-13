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
