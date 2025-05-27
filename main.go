package main

import (
	"fmt"
	"strings"
	"time"
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

	fmt.Println("============================")
	fmt.Println("     SELAMAT DATANG DI")
	fmt.Println("     WORKOUT TRACKER")
	fmt.Println("============================")

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
			editWorkout(&workout_activities, jumlah_data, daftar_latihan[:], jumlah_latihan)
		case 3:
			hapusWorkout(&workout_activities, &jumlah_data)
		case 4:
      cariHistoryWorkout(workout_activities, jumlah_data)
		case 5:
			cariLatihan(daftar_latihan[:])
		case 6:
			urutkanHistoryWorkout(workout_activities, jumlah_data)
		case 7:
			lihatLaporan(workout_activities, jumlah_data)
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
		id:              *jumlah_data,
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

	fmt.Println("\nPilih jenis latihan:")
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

func listWorkout(activities *[Nmax]workout_activity, jumlah_data int) {
	var i int
	fmt.Println("========== Daftar Workout ==========")
	i = 0
	fmt.Println("ID | Tanggal     | Jenis Latihan                  | Durasi (menit) | Kalori Terbakar | Kategori")
	for i < jumlah_data {
		fmt.Printf("%2d | %-11s | %-30s | %14d | %15d | %s\n",
			activities[i].id,
			activities[i].tanggal,
			activities[i].jenis,
			activities[i].durasi,
			activities[i].kalori_terbakar,
			activities[i].kategori)
		i++
	}
}

func editWorkout(activities *[Nmax]workout_activity, jumlah_data int, available_latihan []latihan, num_available_latihan int) {
	listWorkout(&(*activities), jumlah_data)

	var id int
	fmt.Print("Masukkan ID workout yang ingin diedit: ")
	fmt.Scan(&id)
	if id < 0 || id > jumlah_data {
		fmt.Println("ID tidak valid. Silakan coba lagi.")
		return
	}
	var tanggal, jenis_latihan, kategori string
	var durasi, kalori, pilihan_latihan int

	fmt.Println("Ingin edit tanggal? [Ya/Tidak]")
	var edit_tanggal string
	fmt.Scan(&edit_tanggal)
	if edit_tanggal == "Tidak" || edit_tanggal == "tidak" {
		tanggal = activities[id].tanggal
	} else {
		tanggal = inputTanggal()
	}

	fmt.Println("Ingin edit jenis latihan? [Ya/Tidak]\n")
	var edit_latihan string
	fmt.Scan(&edit_latihan)
	if edit_latihan == "Tidak" || edit_latihan == "tidak" {
		jenis_latihan = activities[id].jenis
		kategori = activities[id].kategori
	} else {
		pilihan_latihan = inputPilihan(num_available_latihan, available_latihan)
		jenis_latihan = available_latihan[pilihan_latihan-1].nama
		kategori = available_latihan[pilihan_latihan-1].kategori
	}

	fmt.Println("Ingin edit durasi? [Ya/Tidak]")
	var edit_durasi string
	fmt.Scan(&edit_durasi)
	if edit_durasi == "Tidak" || edit_durasi == "tidak" {
		durasi = activities[id].durasi
	} else {
		durasi = inputPositif("durasi (menit)")
	}

	fmt.Println("Ingin edit kalori terbakar? [Ya/Tidak]")
	var edit_kalori string
	fmt.Scan(&edit_kalori)
	if edit_kalori == "Tidak" || edit_kalori == "tidak" {
		kalori = activities[id].kalori_terbakar
	} else {
		kalori = inputPositif("kalori terbakar")
	}

	activities[id] = workout_activity{
		id:              id,
		tanggal:         tanggal,
		jenis:           jenis_latihan,
		durasi:          durasi,
		kalori_terbakar: kalori,
		kategori:        kategori,
	}
	fmt.Println("Workout berhasil diedit.")
}

func hapusWorkout(workout_activities *[Nmax]workout_activity, jumlah_data *int) {
	var id, index int
	if *jumlah_data == 0 {
		fmt.Println("Belum ada data workout untuk dihapus.")
		return
	}
	fmt.Println("\n=== Hapus Workout ===")
	listWorkout(&(*workout_activities), *jumlah_data)

	fmt.Print("Masukkan ID workout yang ingin dihapus: ")
	fmt.Scan(&id)
	if id < 0 || id > *jumlah_data {
		fmt.Println("ID tidak valid. Silakan coba lagi.")
		return
	}

	index = binarySearchByID(id, *jumlah_data, *workout_activities)
	if index == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	fmt.Printf("Yakin ingin menghapus workout ID %d (%s)? (Ya/Tidak): ", workout_activities[index].id, workout_activities[index].jenis)
	var konfirmasi string
	fmt.Scan(&konfirmasi)
	if strings.ToLower(konfirmasi) != "ya" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := index; i < *jumlah_data-1; i++ {
		workout_activities[i] = workout_activities[i+1]
	}

	(*jumlah_data)--
	fmt.Println("Workout berhasil dihapus.")
}

func binarySearchByID(id, jumlah_data int, workout_activities [Nmax]workout_activity ) int {
	low := 0
	high := jumlah_data - 1

	for low <= high {
		mid := (low + high) / 2
		if workout_activities[mid].id == id {
			return mid
		} else if workout_activities[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariHistoryWorkout(workout_activities [Nmax]workout_activity, jumlah_data int) {
	if jumlah_data == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}

	fmt.Println("\n=== Cari History Workout ===")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Jenis")
	fmt.Println("2. Kategori")
	fmt.Print("Pilih opsi: ")

	var opsi int
	fmt.Scan(&opsi)

	var keyword string
	switch opsi {
	case 1:
		fmt.Print("Masukkan jenis latihan: ")
		fmt.Scan(&keyword)
		cariWorkoutByField("jenis", keyword, workout_activities, jumlah_data)
	case 2:
		fmt.Print("Masukkan kategori latihan: ")
		fmt.Scan(&keyword)
		cariWorkoutByField("kategori", keyword, workout_activities, jumlah_data)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariWorkoutByField(field, keyword string, workout_activities [Nmax]workout_activity, jumlah_data int) {
	var found bool
	keyword = strings.ToLower(keyword)
	found = false

	fmt.Println("\nHasil pencarian:")
	for i := 0; i < jumlah_data; i++ {
		var fieldValue string
		switch field {
		case "jenis":
			fieldValue = workout_activities[i].jenis
		case "kategori":
			fieldValue = workout_activities[i].kategori
		}
		if strings.ToLower(fieldValue) == keyword {
			w := workout_activities[i]
			fmt.Printf("[%d] %s | %s | %d menit | %d kalori\n", w.id, w.tanggal, w.jenis, w.durasi, w.kalori_terbakar)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariLatihan(daftar_latihan []latihan) {
	fmt.Println("\n=== Cari Latihan ===")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Jenis")
	fmt.Println("2. Kategori")
	fmt.Print("Pilih opsi: ")

	var opsi int
	fmt.Scan(&opsi)

	var keyword string
	switch opsi {
	case 1:
		fmt.Print("Masukkan jenis latihan: ")
		fmt.Scan(&keyword)
		cariLatihanByField("nama", keyword, daftar_latihan)
	case 2:
		fmt.Print("Masukkan kategori latihan: ")
		fmt.Scan(&keyword)
		cariLatihanByField("kategori", keyword, daftar_latihan)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariLatihanByField(field, keyword string, daftar_latihan []latihan) {
	var found bool
	keyword = strings.ToLower(keyword)
	found = false

	fmt.Println("\nHasil pencarian:")
	for _, l := range daftar_latihan {
		var fieldValue string
		switch field {
		case "nama":
			fieldValue = l.nama
		case "kategori":
			fieldValue = l.kategori
		}

		if strings.ToLower(fieldValue) == keyword {
			fmt.Printf("- %s (%s)\n", l.nama, l.kategori)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan, periksa kembali input yang dimasukkan.")
	}
}

func urutkanHistoryWorkout(workout_activities [Nmax]workout_activity, jumlah_data int) {
	if jumlah_data == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}

	fmt.Println("\n=== Urutkan History Workout ===")
	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Durasi")
	fmt.Println("2. Kalori Terbakar")
	fmt.Print("Pilih: ")

	var kriteria int
	fmt.Scan(&kriteria)

	fmt.Println("Pilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih: ")

	var urutan int
	fmt.Scan(&urutan)

	sorted_activities := workout_activities
	switch urutan {
	case 1:
		selectionSort(kriteria, &sorted_activities, jumlah_data)
	case 2:
		insertionSort(kriteria, &sorted_activities, jumlah_data)
	default:
		fmt.Println("Pilihan urutan tidak valid.")
		return
	}

	fmt.Println("Data setelah diurutkan:")
	tampilkanWorkout(sorted_activities, jumlah_data)
}

func selectionSort(kriteria int, workout_activities *[Nmax]workout_activity, jumlah_data int) {
	var i, min_idx, j int
	for i = 0; i < jumlah_data-1; i++ {
		min_idx = i
		for j = i + 1; j < jumlah_data; j++ {
			if compare(workout_activities[j], workout_activities[min_idx], kriteria, true) {
				min_idx = j
			}
		}
		workout_activities[i], workout_activities[min_idx] = workout_activities[min_idx], workout_activities[i]
	}
}

func insertionSort(kriteria int, workout_activities *[Nmax]workout_activity, jumlah_data int) {
	for i := 1; i < jumlah_data; i++ {
		key := workout_activities[i]
		j := i - 1
		for j >= 0 && compare(key, workout_activities[j], kriteria, false) {
			workout_activities[j+1] = workout_activities[j]
			j--
		}
		workout_activities[j+1] = key
	}
}

func compare(a, b workout_activity, kriteria int, asc bool) bool {
	switch kriteria {
	case 1:
		if asc {
			return a.durasi < b.durasi
		} else {
			return a.durasi > b.durasi
		}
	case 2:
		if asc {
			return a.kalori_terbakar < b.kalori_terbakar
		} else {
			return a.kalori_terbakar > b.kalori_terbakar
		}
	}
	return false
}

func tampilkanWorkout(workout_activities [Nmax]workout_activity, jumlahData int) {
	var i int
	var w workout_activity
	fmt.Println("\nDaftar Workout:")
	fmt.Println("ID | Tanggal     | Jenis Latihan                  | Durasi (menit) | Kalori Terbakar | Kategori")
	for i = 0; i < jumlahData; i++ {
		w = workout_activities[i]
		fmt.Printf("%2d | %-11s | %-30s | %14d | %15d | %s\n",
		w.id,
		w.tanggal,
		w.jenis,
		w.durasi,
		w.kalori_terbakar,
		w.kategori)
	}
}

func lihatLaporan(workout_activities [Nmax]workout_activity, jumlah_data int) {
	if jumlah_data == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}
	
	fmt.Println("\n=== Laporan Workout ===")
	fmt.Println("Pilih laporan yang ingin ditampilkan:")
	fmt.Println("1. Tampilkan 10 aktivitas workout terakhir")
	fmt.Println("2. Hitung kalori yang terbakar pada periode tertentu")
	fmt.Print("Pilih laporan: ")
	
	var pilih_laporan int
	fmt.Scan(&pilih_laporan)

	if pilih_laporan == 1 {
		tampilkan10AktivitasTerakhir(workout_activities, jumlah_data)
	} else if pilih_laporan == 2 {
		hitungKaloriDalamPeriode(workout_activities, jumlah_data)
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}
}

func tampilkan10AktivitasTerakhir(workout_activities [Nmax]workout_activity, jumlah_data int) {
	var start, end, i int
  fmt.Println("\n10 Aktivitas Terakhir:")
  start = jumlah_data - 1
  end = start - 9
  if end < 0 {
      end = 0
  }

	fmt.Println("ID | Tanggal     | Jenis Latihan                  | Durasi (menit) | Kalori Terbakar | Kategori")
  for i = start; i >= end; i-- {
    w := workout_activities[i]
    fmt.Printf("%2d | %-11s | %-30s | %14d | %15d | %s\n",
			w.id,
			w.tanggal,
			w.jenis,
			w.durasi,
			w.kalori_terbakar,
			w.kategori)
  }
}

func hitungKaloriDalamPeriode(workout_activities [Nmax]workout_activity, jumlah_data int) {
    var tanggal_awal, tanggal_akhir, layout string
		var tData, tAwal, tAkhir time.Time
		var err error

    fmt.Println("\nMasukkan tanggal awal (YYYY-MM-DD): ")
    fmt.Scan(&tanggal_awal)
    fmt.Println("Masukkan tanggal akhir (YYYY-MM-DD): ")
    fmt.Scan(&tanggal_akhir)

    layout = "2006-01-02"
    tAwal, err = time.Parse(layout, tanggal_awal)
    if err != nil {
        fmt.Println("Format tanggal awal tidak valid.")
        return
    }
    tAkhir, err = time.Parse(layout, tanggal_akhir)
    if err != nil {
        fmt.Println("Format tanggal akhir tidak valid.")
        return
    }

		var total_kalori, i int
    total_kalori = 0
    for i = 0; i < jumlah_data; i++ {
        tData, err = time.Parse(layout, workout_activities[i].tanggal)
        if err == nil && (tData.Equal(tAwal) || tData.After(tAwal)) && (tData.Equal(tAkhir) || tData.Before(tAkhir)) {
            total_kalori += workout_activities[i].kalori_terbakar
        }
    }

    fmt.Printf("\nTotal kalori terbakar dalam periode %s hingga %s adalah %d kalori.\n", tanggal_awal, tanggal_akhir, total_kalori)
}


	