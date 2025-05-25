package main

import (
	"fmt"
)

const Nmax = 1000

type latihan struct {
	nama     string
	kategori string
}

type workout_activity struct {
	id              int
	tanggal         string
	jenis           string 
	durasi          int
	kalori_terbakar int
	kategori        string
}

func main() {
	const jumlah_latihan = 15

	var daftar_latihan = [jumlah_latihan]latihan{
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

	var workout_activities [Nmax]workout_activity
	var jumlah_data int = 0

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

		var pilihan_menu int
		fmt.Scanln(&pilihan_menu)

		switch pilihan_menu {
		case 1:
			tambahWorkout(&workout_activities, &jumlah_data, daftar_latihan[:], jumlah_latihan)
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

func tambahWorkout(activities *[Nmax]workout_activity, jumlah_data *int, available_latihan []latihan, num_available_latihan int) {
	var pilihan_latihan, durasi, kalori int
	var tanggal string
	var selected_latihan latihan

	if *jumlah_data >= Nmax {
		fmt.Println("History workout penuh.")
		return
	}

	fmt.Println("=== Tambah Workout ===")

	pilihan_latihan = inputPilihan(num_available_latihan, available_latihan)
	tanggal = inputTanggal()
	durasi = inputPositif("durasi (menit)")
	kalori = inputPositif("kalori terbakar")

	selected_latihan = available_latihan[pilihan_latihan-1]

	activities[*jumlah_data] = workout_activity{
		id:              *jumlah_data + 1,
		tanggal:         tanggal,
		jenis:           selected_latihan.nama,
		durasi:          durasi,
		kalori_terbakar: kalori,
		kategori:        selected_latihan.kategori,
	}

	(*jumlah_data)++
	fmt.Println("Workout berhasil ditambahkan.")
}

func inputPilihan(num_available_latihan int, available_latihan_list []latihan) int {
	var i int
	var pilihan int

	fmt.Println("Pilih jenis latihan:")
	for i < num_available_latihan {
		fmt.Printf("%2d. %-30s (%s)\n", i+1, available_latihan_list[i].nama, available_latihan_list[i].kategori)
		i++
	}
	fmt.Print("Masukkan nomor latihan: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > num_available_latihan {
		fmt.Println("Nomor latihan tidak valid.")
		return inputPilihan(num_available_latihan, available_latihan_list)
	}
	return pilihan
}

func inputTanggal() string {
	var tanggal string
	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)
	return tanggal
}

func inputPositif(label string) int {
	var val int
	fmt.Printf("Masukkan %s: ", label)
	fmt.Scan(&val)

	if val <= 0 {
		fmt.Printf("%s tidak boleh nol atau negatif.\n", label)
		return inputPositif(label)
	}
	return val
}
