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
