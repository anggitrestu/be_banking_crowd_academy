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

// type MyClassesFormatter struct {
// 	ID        int       `json:"id"`
// 	LearnerID int       `json:"learner_id"`
// 	ClassID   int       `json:"class_id"`
// 	Classs    class     `json:"classs"`
// 	Learner   []learner `json:"learner"`
// }

// type class struct {
// 	ID        int    `json:"id"`
// 	TutorID   int    `json:"tutor_id"`
// 	Topik     string `json:"topik"`
// 	Jenis     string `json:"jenis"`
// 	Judul     string `json:"judul"`
// 	Jadwal    string `json:"jadwal"`
// 	LinkZoom  string `json:"link_zoom"`
// 	Deskripsi string `json:"deskripsi"`
// 	Modul     string `json:"modul"`
// }

// type learner struct {
// 	ID   int    `json:"id"`
// 	Nama string `json:"nama"`
// }

// func FormatMyClasses(myclass MyClass) *MyClassesFormatter {
// 	formatter := MyClassesFormatter{
// 		ID:        myclass.ID,
// 		LearnerID: myclass.LearnerID,
// 		ClassID:   myclass.ClassID,
// 		Classs: class{
// 			ID:        myclass.Class.ID,
// 			TutorID:   myclass.Class.TutorID,
// 			Topik:     myclass.Class.Topik,
// 			Jenis:     myclass.Class.Jenis,
// 			Judul:     myclass.Class.Judul,
// 			Jadwal:    myclass.Class.Jadwal,
// 			LinkZoom:  myclass.Class.LinkZoom,
// 			Deskripsi: myclass.Class.Deskripsi,
// 			Modul:     myclass.Class.Modul,
// 		},
// 	}

// 	return &formatter
// }

/*
	ID:   myclass.Learners[0].ID,
			Nama: myclass.Learners[0].Nama,
*/
