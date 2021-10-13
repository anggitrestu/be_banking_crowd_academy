package tutors

type TutorFormatter struct {
	ID    int    `json:"id_tutor"`
	Nama  string `json:"nama_tutor"`
	Email string `json:"email_tutor"`
	Token string `json:"token"`
}

func FormatTutor(tutor Tutor, token string) TutorFormatter {
	formatter := TutorFormatter{
		ID:    tutor.ID,
		Nama:  tutor.Nama,
		Email: tutor.Email,
		Token: token,
	}

	return formatter

}
