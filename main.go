// main.go
package main

import (
	"fmt"
)

const Nmax = 100

// Struct utama untuk histori workout
type workout_activity struct {
	id              int
	tanggal         string
	jenis           string
	durasi          int
	kalori_terbakar int
	kategori        string
}

// Struct data latihan
type latihan struct {
	nama     string
	kategori string
}

// Global variables
var workout_activities [Nmax]workout_activity
var jumlahData int = 0
var daftar_latihan = [15]latihan{
	{"Dumbbell bench press", "Strength"},
	{"Squat", "Strength"},
	{"Deadlift", "Strength"},
	{"Dumbbell goblet squat", "Strength"},
	{"Dumbbell lateral raise", "Strength"},
	{"Dumbbell shoulder press", "Strength"},
	{"Leg curl", "Strength"},
	{"Bulgarian split squat", "Strength"},
	{"Dumbbell flys", "Strength"},
	{"Jumping lunge", "Cardio"},
	{"Pushups", "Strength"},
	{"Barbell curl", "Strength"},
	{"Barbell deadlift", "Strength"},
	{"Bent-over row", "Strength"},
	{"Incline dumbbell bench press", "Strength"},
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("            SELAMAT DATANG DI")
	fmt.Println("              WORKOUT TRACKER")
	fmt.Println("==================================================")
	fmt.Println("Mencatat latihan. Meningkatkan performa. Menjadi lebih kuat.")
	fmt.Println("--------------------------------------------------")

	for {
		fmt.Println("\n============================")
		fmt.Println("     WORKOUT TRACKER")
		fmt.Println("============================")
		fmt.Println("1. Tambah Workout")
		fmt.Println("2. Edit Workout")
		fmt.Println("3. Hapus Workout")
		fmt.Println("4. Cari History Workout")
		fmt.Println("5. Cari Latihan")
		fmt.Println("6. Urutkan History Workout")
		fmt.Println("7. Lihat Laporan")
		fmt.Println("8. Rekomendasi Workout")
		fmt.Println("9. Keluar")
		fmt.Print("\nPilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("[Tambah Workout] - fitur belum diimplementasikan")
		case 2:
			fmt.Println("[Edit Workout] - fitur belum diimplementasikan")
		case 3:
			fmt.Println("[Hapus Workout] - fitur belum diimplementasikan")
		case 4:
			fmt.Println("[Cari History Workout] - fitur belum diimplementasikan")
		case 5:
			fmt.Println("[Cari Latihan] - fitur belum diimplementasikan")
		case 6:
			fmt.Println("[Urutkan History Workout] - fitur belum diimplementasikan")
		case 7:
			fmt.Println("[Lihat Laporan] - fitur belum diimplementasikan")
		case 8:
			fmt.Println("[Rekomendasi Workout] - fitur belum diimplementasikan")
		case 9:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
