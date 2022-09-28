package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orang struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	daftarTeman := []Orang{
		{Nama: "Nabila Cahyani", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Latifah", Alamat: "Jogja", Pekerjaan: "Freelancer", Alasan: "menambah ilmu dan untuk mencari kerja"},
		{Nama: "Vita Tri Utami", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "menambah pengetahuan dalam service suatu website"},
		{Nama: "Kulsum Khoiriah", Alamat: "Bogor", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Komang Andira Trisna P.", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambah pengetahuan dan pemahaman bahasa Go"},
		{Nama: "Amelia Sari", Alamat: "Pamulang", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Feti Lutviyani ", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambahskill di bidang IT"},
		{Nama: "Jesica", Alamat: "Padang", Pekerjaan: "Freelance", Alasan: "ingin mengenal pemrograman golang "},
		{Nama: "Vini Jumatul Fitri", Alamat: "Medan", Pekerjaan: "Freelance", Alasan: "menambah bahasa pemrograman"},
		{Nama: "Natasya Sitorus", Alamat: "Semarang", Pekerjaan: "Freelance", Alasan: "Untuk menambah pengetahuan lebih mengenai go"},
		{Nama: "Hanafi Adhi Prasetyo", Alamat: "Cilacap", Pekerjaan: "Freelance", Alasan: "Untuk menambah pengetahuan lebih mengenai go"},
	}
	number := convert((os.Args[1]))
	lengthData := len(daftarTeman) //11
	if number <= lengthData {
		fmt.Println("Nama: ", daftarTeman[number-1].Nama)
		fmt.Println("Alamat: ", daftarTeman[number-1].Alamat)
		fmt.Println("Pekerjaan: ", daftarTeman[number-1].Pekerjaan)
		fmt.Println("Alasan Mengikuti DTS: ", daftarTeman[number-1].Alasan)
	} else {
		fmt.Printf("Biodata Dengan Absen no. %d Tidak ada, No Absen Hanya sampai %d", number, lengthData)
	}

}

func convert(C string) (angka int) {
	angka, err := strconv.Atoi(C)
	if err != nil {
		panic("Silahkan Masukan Angka")
	}
	return
}
