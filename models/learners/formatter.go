package learners

type LearnerFormatter struct {
	ID    int    `json:"id_learner"`
	Nama  string `json:"nama_learner"`
	Email string `json:"email_learner"`
	Token string `json:"token"`
}

func Formatlearner(learner Learner, token string) LearnerFormatter {
	formatter := LearnerFormatter{
		ID:    learner.ID,
		Nama:  learner.Nama,
		Email: learner.Email,
		Token: token,
	}

	return formatter

}

type InfoLearnerFormatter struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	Usia          int    `json:"usia" `
	Pekerjaan     string `json:"pekerjaan" `
	TopikDiminati string `json:"topik_diminati" `
}

func FormatInfoLearner(learner Learner) *InfoLearnerFormatter {
	formatter := InfoLearnerFormatter{
		ID:            learner.ID,
		Nama:          learner.Nama,
		Email:         learner.Email,
		Usia:          learner.Usia,
		Pekerjaan:     learner.Pekerjaan,
		TopikDiminati: learner.TopikDiminati,
	}

	return &formatter

}
