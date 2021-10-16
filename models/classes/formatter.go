package classes

type InfoClassFormatter struct {
	ID        int      `json:"id"`
	TutorID   int      `json:"tutor_id"`
	Topik     string   `json:"topik"`
	Jenis     string   `json:"jenis"`
	Judul     string   `json:"judul"`
	Jadwal    string   `json:"jadwal"`
	LinkZoom  string   `json:"link_zoom"`
	Deskripsi string   `json:"deskripsi"`
	Modul     string   `json:"modul"`
	Pendaftar []string `json:"pendaftar"`
}

func FormatInfoClass(class Class, pendaftar []string) *InfoClassFormatter {
	formatter := InfoClassFormatter{
		ID:        class.ID,
		TutorID:   class.TutorID,
		Topik:     class.Topik,
		Jenis:     class.Jenis,
		Judul:     class.Judul,
		Jadwal:    class.Jadwal,
		LinkZoom:  class.Jadwal,
		Deskripsi: class.Deskripsi,
		Modul:     class.Modul,
		Pendaftar: pendaftar,
	}

	return &formatter
}

func FormatInfoClasses(classes []Class, pendaftar []string) *[]InfoClassFormatter {
	classesFormatter := []InfoClassFormatter{}

	for _, class := range classes {
		classFormatter := FormatInfoClass(class, pendaftar)
		classesFormatter = append(classesFormatter, *classFormatter)
	}
	return &classesFormatter
}

// type InfoLengkapClassFormatter struct {
// 	ID        int      `json:"id"`
// 	TutorID   int      `json:"tutor_id"`
// 	Topik     string   `json:"topik"`
// 	Jenis     string   `json:"jenis"`
// 	Judul     string   `json:"judul"`
// 	Jadwal    string   `json:"jadwal"`
// 	LinkZoom  string   `json:"link_zoom"`
// 	Deskripsi string   `json:"deskripsi"`
// 	Modul     string   `json:"modul"`
// 	Pendaftar []string `json:"pendaftar"`
// }

// func InfoLengkap(classes []Class, emails []string) {
// 	infoLengkap := InfoLengkapClassFormatter{}

// 	for _, class := range classes {

// 	}
// }
