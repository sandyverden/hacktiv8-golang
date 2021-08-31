package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	NoAbsen   int
	Nama      string
	Alamat    string
	Pekerjaan string
	Motivasi  string
}

func (b Biodata) GetBiodata(noAbsen int) {
	fmt.Println("Biodata Teman Kelas Golang dengan No Absen: ", b.NoAbsen)
	fmt.Println("Nama:", b.Nama)
	fmt.Println("Alamat:", b.Alamat)
	fmt.Println("Pekerjaan:", b.Pekerjaan)
	fmt.Println("Alasan mengikuti kelas Golang:", b.Motivasi)
}

func main() {

	noAbsen, _ := strconv.ParseInt(os.Args[1], 10, 0)
	//fmt.Println(noAbsen)
	switch {
	case noAbsen == 1:
		{
			peserta := Biodata{
				NoAbsen:   1,
				Nama:      "Firman Aulia Insani",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 2:
		{
			peserta := Biodata{
				NoAbsen:   2,
				Nama:      "Andri Nur Hidayatulloh",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 3:
		{
			peserta := Biodata{
				NoAbsen:   3,
				Nama:      "I Komang Widnyana",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 4:
		{
			peserta := Biodata{
				NoAbsen:   4,
				Nama:      "Erico",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 5:
		{
			peserta := Biodata{
				NoAbsen:   5,
				Nama:      "Jose Yolanda Purba",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 6:
		{
			peserta := Biodata{
				NoAbsen:   6,
				Nama:      "Andri Kuwito",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 7:
		{
			peserta := Biodata{
				NoAbsen:   7,
				Nama:      "Sandy Budiman",
				Alamat:    "Indonesia",
				Pekerjaan: "DevOps Engineer",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang untuk meningkatkan karir",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 8:
		{
			peserta := Biodata{
				NoAbsen:   2,
				Nama:      "Rafli Andreansyah",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 9:
		{
			peserta := Biodata{
				NoAbsen:   9,
				Nama:      "Taupiq Hidayah",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	case noAbsen == 10:
		{
			peserta := Biodata{
				NoAbsen:   10,
				Nama:      "Melvita Sari",
				Alamat:    "Indonesia",
				Pekerjaan: "Mahasiswa",
				Motivasi:  "Ingin menguasai bahasa pemprograman Golang",
			}
			peserta.GetBiodata(peserta.NoAbsen)
		}
	default:
		{
			fmt.Println("Nomor Absen tidak ada")
		}
	}

}
