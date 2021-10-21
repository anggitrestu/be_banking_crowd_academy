package myclasses

type MyClassFormatter struct {
	ID        int `json:"id"`
	LearnerID int `json:"learner_id"`
	ClassID   int `json:"class_id"`
}

func FormatMyClass(myclass MyClass) *MyClassFormatter {
	formatter := MyClassFormatter{
		ID:        myclass.ID,
		LearnerID: myclass.LearnerID,
		ClassID:   myclass.ClassID,
	}
	return &formatter
}

type class struct {
	ID        int    `json:"id"`
	TutorID   int    `json:"tutor_id"`
	Topik     string `json:"topik"`
	Jenis     string `json:"jenis"`
	Judul     string `json:"judul"`
	Jadwal    string `json:"jadwal"`
	LinkZoom  string `json:"link_zoom"`
	Deskripsi string `json:"deskripsi"`
	Modul     string `json:"modul"`
}

type ResponseMyClass struct {
	ID        int   `json:"id"`
	ClassID   int   `json:"class_id"`
	LearnerID int   `json:"learner_id"`
	Class     class `json:"class"`
}
