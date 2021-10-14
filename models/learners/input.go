package learners

type CreateLearnerInput struct {
	Nama          string `json:"nama" binding:"required"`
	Usia          int    `json:"usia" binding:"required"`
	Pekerjaan     string `json:"pekerjaan" binding:"required"`
	TopikDiminati string `json:"topik_diminati" binding:"required"`
}

type GetLearnerInput struct {
	ID int `uri:"id" binding:"required"`
}
