package tutors

type TutorFormatter struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatTutor(tutor Tutor, token string) *TutorFormatter {
	formatter := TutorFormatter{
		ID:    tutor.ID,
		Nama:  tutor.Nama,
		Email: tutor.Email,
		Token: token,
	}

	return &formatter

}

type InfoTutorFormatter struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	MasaKerja     int    `json:"masa_kerja"`
	SitusWeb      string `json:"situs_web"`
	Kompetensi    string `json:"kompetensi"`
	Pekerjaan     string `json:"pekerjaan"`
	TopikDiminati string `json:"topik_diminati"`
}

func FormatInfoTutor(tutor Tutor) *InfoTutorFormatter {
	formatter := InfoTutorFormatter{
		ID:            tutor.ID,
		Nama:          tutor.Nama,
		Email:         tutor.Email,
		MasaKerja:     tutor.MasaKerja,
		SitusWeb:      tutor.SitusWeb,
		Kompetensi:    tutor.Kompetensi,
		Pekerjaan:     tutor.Pekerjaan,
		TopikDiminati: tutor.TopikDiminati,
	}

	return &formatter
}
