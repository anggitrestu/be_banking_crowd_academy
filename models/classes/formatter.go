package classes

type ResponseClass struct {
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
