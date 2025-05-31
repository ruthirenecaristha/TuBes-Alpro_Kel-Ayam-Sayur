package main

import (
	"fmt"
	"strconv" // Untuk konversi string ke int (untuk pilihan kategori)
	"time"    // Untuk tipe data waktu
)

const NMAX int = 100 // batas maksimal jumlah isi array

// tipe data struct untuk array
type subscription struct {
	name          string    // nama dan kategori layanan
	category      string    // kategori layanan
	monthlyCost   float64   // biaya bulanan langganan
	paymentDate   time.Time // waktu mulai pembayaran langganan
	paymentMethod string    // metode pembayaran langganan
	frequency     string    // frekuensi pembayaran (e.g., "bulanan", "tahunan", "mingguan")
	endDate       time.Time // waktu berakhir langganan (jika ada)
	notes         string    // catatan tentang langganan
	active        bool      // status langganan (aktif/tidak)
}
type tabSubs [NMAX]subscription // array tabSubs dengan tipe data subscription

// Daftar kategori yang sudah ada
var predefinedCategories = [8]string{"Hiburan", "Edukasi", "Produktivitas", "Berita", "Kesehatan & Kebugaran", "Cloud Storage", "Musik", "Video Streaming"}

func main() {
	var choice, nData int
	var dataSub tabSubs

	for {
		menu()
		fmt.Print("Pilih menu (1-14): ") // Perbarui jumlah menu
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			addSubscription(&dataSub, &nData)
		case 2:
			editSubscription(&dataSub, nData)
		case 3:
			deleteSubscription(&dataSub, &nData)
		case 4:
			viewSubscriptions(&dataSub, nData) // Panggil dengan pointer
		case 5:
			searchByName(&dataSub, nData)
		case 6:
			searchByCategory(dataSub, nData)
		case 7:
			sortByCostSelection(&dataSub, nData)
		case 8:
			sortByDateInsertion(&dataSub, nData)
		case 9:
			calculateTotalAndRecommend(dataSub, nData)
		case 10:
			highestCost(dataSub, nData)
		case 11:
			reminder(dataSub, nData)
		case 12: // Pengeluaran per Kategori
			displayCostByCategory(dataSub, nData)
		case 13: // Proyeksi Pengeluaran
			projectFutureExpenses(dataSub, nData)
		case 14: // Keluar
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan Tidak Valid.")
		}
	}
}

// menampilkan menu pilihan
func menu() {
	fmt.Println("\n=== Aplikasi Manajemen Subskripsi ===")
	fmt.Println("1. Tambah Langganan")
	fmt.Println("2. Update Langganan")
	fmt.Println("3. Hapus Langganan")
	fmt.Println("4. Lihat Semua Langganan")
	fmt.Println("5. Cari Langganan Berdasarkan Nama (Binary)")
	fmt.Println("6. Cari Langganan Berdasarkan Kategori (Sequential)")
	fmt.Println("7. Urutkan Langganan Berdasarkan Biaya (Selection Sort)")
	fmt.Println("8. Urutkan Langganan Berdasarkan Tanggal (Insertion Sort)")
	fmt.Println("9. Total Pengeluaran dan Rekomendasi")
	fmt.Println("10. Cari Biaya Tertinggi")
	fmt.Println("11. Pengingat Pembayaran")
	fmt.Println("12. Total Pengeluaran per Kategori")
	fmt.Println("13. Proyeksi Pengeluaran")
	fmt.Println("14. Keluar")
}

// prosedur menambah langganan
func addSubscription(S *tabSubs, n *int) {
	var name, category, method, frequency, notes, dateStr, endDateStr string
	var cost float64
	var x int

	fmt.Print("Masukkan jumlah layanan yang ingin ditambah: ")
	fmt.Scan(&x)

	for i := 0; i < x; i++ {
		if *n >= NMAX {
			fmt.Println("Array langganan penuh, tidak bisa menambah lagi.")
			return
		}

		fmt.Print("Nama Layanan: ")
		fmt.Scan(&name)

		// Deteksi Duplikat
		if seqSearch(*S, *n, name) != -1 {
			fmt.Printf("Langganan dengan nama '%s' sudah ada. Silakan masukkan nama lain.\n", name)
			i-- // Kurangi i agar iterasi ini diulang
			continue
		}

		fmt.Println("\nPilih Kategori:")
		for idx, cat := range predefinedCategories {
			fmt.Printf("%d. %s\n", idx+1, cat)
		}
		fmt.Printf("%d. Kategori Lainnya\n", len(predefinedCategories)+1)
		fmt.Print("Pilihan (angka) atau ketik kategori baru jika memilih 'Kategori Lainnya': ")

		var categoryInput string
		fmt.Scan(&categoryInput)

		var chosenCategory string
		var choiceNum int // Deklarasi di sini
		var err error     // Deklarasi di sini

		choiceNum, err = strconv.Atoi(categoryInput)
		if err == nil && choiceNum > 0 && choiceNum <= len(predefinedCategories) {
			chosenCategory = predefinedCategories[choiceNum-1]
		} else if err == nil && choiceNum == len(predefinedCategories)+1 {
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scan(&chosenCategory)
		} else {
			chosenCategory = categoryInput // Anggap input yang diketik langsung sebagai kategori baru
			fmt.Println("Pilihan tidak valid, menggunakan input Anda sebagai kategori baru.")
		}
		category = chosenCategory // Set kategori yang dipilih/dimasukkan

		fmt.Print("Biaya Bulanan: ")
		fmt.Scan(&cost)

		// Validasi biaya positif
		if cost <= 0 {
			fmt.Println("Biaya bulanan harus lebih dari 0.")
			i--
			continue
		}

		fmt.Print("Tanggal Pembayaran (YYYY-MM-DD): ")
		fmt.Scan(&dateStr)

		var paymentDate time.Time                                                                   // Deklarasi `paymentDate` di sini
		paymentDate, err = time.Parse("2006-01-02", dateStr)                                        // Assignment ulang `err`
		if err != nil || !dateIsValid(paymentDate.Year(), paymentDate.Month(), paymentDate.Day()) { // Gunakan `paymentDate`
			fmt.Println("Format tanggal pembayaran salah atau tidak valid! (YYYY-MM-DD)")
			i-- // Minta input ulang untuk item yang sama
			continue
		}

		// Tanggal Akhir Langganan
		fmt.Print("Tanggal Akhir Langganan (YYYY-MM-DD, kosongkan jika tidak ada): ")
		fmt.Scan(&endDateStr)
		var endDate time.Time
		if endDateStr != "" {
			parsedDate, errParseEnd := time.Parse("2006-01-02", endDateStr)
			if errParseEnd != nil {
				fmt.Println("Format tanggal akhir salah, tanggal akhir tidak akan disimpan.")
			} else {
				endDate = parsedDate
			}
		}

		fmt.Print("Metode Pembayaran: ")
		fmt.Scan(&method)

		// Frekuensi Pembayaran
		fmt.Print("Frekuensi Pembayaran (bulanan/tahunan/mingguan): ")
		fmt.Scan(&frequency)
		if frequency != "bulanan" && frequency != "tahunan" && frequency != "mingguan" {
			fmt.Println("Frekuensi pembayaran tidak valid. Default ke 'bulanan'.")
			frequency = "bulanan"
		}

		// Catatan
		fmt.Print("Catatan (opsional, ketik satu kata atau gunakan underscore untuk spasi): ")
		fmt.Scan(&notes)

		S[*n] = subscription{
			name:          name,
			category:      category,
			monthlyCost:   cost,
			paymentDate:   paymentDate, // Gunakan `paymentDate` yang sudah dideklarasikan
			paymentMethod: method,
			frequency:     frequency,
			endDate:       endDate,
			notes:         notes,
			active:        true,
		}
		*n = *n + 1
		fmt.Println()
	}
	fmt.Printf("%d langganan berhasil ditambahkan!\n", x)
}

// fungsi untuk memastikan tanggal berlangganan valid
func dateIsValid(year int, month time.Month, day int) bool {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return t.Year() == year && t.Month() == month && t.Day() == day
}

// prosedur mengedit / mengupdate langganan
func editSubscription(S *tabSubs, n int) {
	var name, newDateStr string
	fmt.Print("Masukkan nama langganan yang ingin diupdate: ")
	fmt.Scan(&name)

	idx := seqSearch(*S, n, name)
	if idx == -1 {
		fmt.Println("Langganan tidak ditemukan.")
		return
	}

	fmt.Println("\nDetail Langganan Saat Ini:")
	viewDetails(S[idx])

	fmt.Println("\nApa yang ingin Anda ubah? (ketik 'tanggal', 'kategori', 'biaya', 'metode', 'frekuensi', 'akhir', 'catatan', atau 'semua'): ")
	var editChoice string
	fmt.Scan(&editChoice)

	// Deklarasi variabel untuk digunakan di berbagai case
	var (
		newCategoryInput  string
		chosenNewCategory string
		choiceNum         int
		err               error // Deklarasi err di sini untuk seluruh switch statement
		newCost           float64
		newMethod         string
		newFrequency      string
		newEndDateStr     string
		newNotes          string
	)

	switch editChoice {
	case "tanggal":
		fmt.Print("Masukkan tanggal pembayaran baru (YYYY-MM-DD): ")
		fmt.Scan(&newDateStr)
		newDate, errDateParse := time.Parse("2006-01-02", newDateStr) // Gunakan variabel error baru
		if errDateParse != nil {
			fmt.Println("Format tanggal salah!")
			return
		}
		S[idx].paymentDate = newDate
	case "kategori":
		fmt.Printf("Kategori saat ini: %s\n", S[idx].category)
		fmt.Println("\nPilih Kategori Baru:")
		for i, cat := range predefinedCategories {
			fmt.Printf("%d. %s\n", i+1, cat)
		}
		fmt.Printf("%d. Kategori Lainnya (atau ketik kategori baru)\n", len(predefinedCategories)+1)
		fmt.Print("Pilihan (angka), ketik kategori baru: ")

		fmt.Scan(&newCategoryInput)
		choiceNum, err = strconv.Atoi(newCategoryInput) // Assignment, bukan deklarasi ulang
		if err == nil && choiceNum > 0 && choiceNum <= len(predefinedCategories) {
			chosenNewCategory = predefinedCategories[choiceNum-1]
		} else if err == nil && choiceNum == len(predefinedCategories)+1 {
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scan(&chosenNewCategory)
		} else {
			chosenNewCategory = newCategoryInput
			fmt.Println("Pilihan tidak valid, menggunakan input Anda sebagai kategori baru.")
		}
		S[idx].category = chosenNewCategory
	case "biaya":
		fmt.Print("Masukkan biaya bulanan baru: ")
		fmt.Scan(&newCost)
		if newCost <= 0 {
			fmt.Println("Biaya harus lebih dari 0. Tidak ada perubahan.")
			return
		}
		S[idx].monthlyCost = newCost
	case "metode":
		fmt.Print("Masukkan metode pembayaran baru: ")
		fmt.Scan(&newMethod)
		S[idx].paymentMethod = newMethod
	case "frekuensi":
		fmt.Print("Masukkan frekuensi pembayaran baru (bulanan/tahunan/mingguan): ")
		fmt.Scan(&newFrequency)
		if newFrequency != "bulanan" && newFrequency != "tahunan" && newFrequency != "mingguan" {
			fmt.Println("Frekuensi pembayaran tidak valid. Tidak ada perubahan.")
			return
		}
		S[idx].frequency = newFrequency
	case "akhir":
		fmt.Print("Masukkan tanggal akhir langganan baru (YYYY-MM-DD, kosongkan untuk menghapus): ")
		fmt.Scan(&newEndDateStr)
		if newEndDateStr == "" {
			S[idx].endDate = time.Time{} // Set ke zero value untuk menghapus
		} else {
			newEndDate, errEndDateParse := time.Parse("2006-01-02", newEndDateStr) // Gunakan variabel error baru
			if errEndDateParse != nil {
				fmt.Println("Format tanggal salah! Tidak ada perubahan.")
				return
			}
			S[idx].endDate = newEndDate
		}
	case "catatan":
		fmt.Print("Masukkan catatan baru (satu kata atau gunakan underscore untuk spasi): ")
		fmt.Scan(&newNotes)
		S[idx].notes = newNotes
	case "semua":
		// Tanggal pembayaran
		fmt.Print("Masukkan tanggal pembayaran baru (YYYY-MM-DD): ")
		fmt.Scan(&newDateStr)
		newDate, errDateParse := time.Parse("2006-01-02", newDateStr) // Gunakan variabel error baru
		if errDateParse != nil {
			fmt.Println("Format tanggal salah! Aborting 'semua' update.")
			return
		}
		S[idx].paymentDate = newDate

		// Kategori
		fmt.Println("\nPilih Kategori Baru:")
		for i, cat := range predefinedCategories {
			fmt.Printf("%d. %s\n", i+1, cat)
		}
		fmt.Printf("%d. Kategori Lainnya (atau ketik kategori baru)\n", len(predefinedCategories)+1)
		fmt.Print("Pilihan (angka), ketik kategori baru: ")
		fmt.Scan(&newCategoryInput)

		choiceNum, err = strconv.Atoi(newCategoryInput) // Assignment, bukan deklarasi ulang
		if err == nil && choiceNum > 0 && choiceNum <= len(predefinedCategories) {
			chosenNewCategory = predefinedCategories[choiceNum-1]
		} else if err == nil && choiceNum == len(predefinedCategories)+1 {
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scan(&chosenNewCategory)
		} else {
			chosenNewCategory = newCategoryInput
			fmt.Println("Pilihan tidak valid, menggunakan input Anda sebagai kategori baru.")
		}
		S[idx].category = chosenNewCategory

		// Biaya
		fmt.Print("Masukkan biaya bulanan baru: ")
		fmt.Scan(&newCost)
		if newCost > 0 {
			S[idx].monthlyCost = newCost
		} else {
			fmt.Println("Biaya harus lebih dari 0. Tidak ada perubahan biaya.")
		}

		// Metode
		fmt.Print("Masukkan metode pembayaran baru: ")
		fmt.Scan(&newMethod)
		S[idx].paymentMethod = newMethod

		// Frekuensi
		fmt.Print("Masukkan frekuensi pembayaran baru (bulanan/tahunan/mingguan): ")
		fmt.Scan(&newFrequency)
		if newFrequency == "bulanan" || newFrequency == "tahunan" || newFrequency == "mingguan" {
			S[idx].frequency = newFrequency
		} else {
			fmt.Println("Frekuensi pembayaran tidak valid. Tidak ada perubahan frekuensi.")
		}

		// Tanggal Akhir
		fmt.Print("Masukkan tanggal akhir langganan baru (YYYY-MM-DD, kosongkan untuk menghapus): ")
		fmt.Scan(&newEndDateStr)
		if newEndDateStr == "" {
			S[idx].endDate = time.Time{}
		} else {
			newEndDate, errEndDateParse := time.Parse("2006-01-02", newEndDateStr) // Gunakan variabel error baru
			if errEndDateParse != nil {
				fmt.Println("Format tanggal salah! Tidak ada perubahan tanggal akhir.")
			} else {
				S[idx].endDate = newEndDate
			}
		}

		// Catatan
		fmt.Print("Masukkan catatan baru (satu kata atau gunakan underscore untuk spasi): ")
		fmt.Scan(&newNotes)
		S[idx].notes = newNotes

	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("Langganan berhasil diperbarui.")
	viewDetails(S[idx])
}

// prosedur menghapus langganan yang tidak diikuti lagi
func deleteSubscription(S *tabSubs, n *int) {
	var name string
	fmt.Print("Masukkan nama langganan yang ingin dihapus: ")
	fmt.Scan(&name)

	idx := seqSearch(*S, *n, name)
	if idx == -1 {
		fmt.Println("Langganan tidak ditemukan.")
		return
	}

	// Konfirmasi Aksi
	var confirm string
	fmt.Printf("Apakah Anda yakin ingin menghapus langganan '%s'? (y/n): ", name)
	fmt.Scan(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := idx; i < *n-1; i++ {
		S[i] = S[i+1]
	}
	*n = *n - 1
	fmt.Printf("%s berhasil dihapus.\n", name)
}

// function untuk mencari nama langganan dengan mengembalikan index array data tersebut
func seqSearch(S tabSubs, n int, x string) int {
	for i := 0; i < n; i++ {
		if S[i].name == x {
			return i
		}
	}
	return -1
}

// prosedur menampilkan seluruh langganan
// Parameter S diubah menjadi pointer (*tabSubs) agar perubahan status `active` dapat disimpan.
func viewSubscriptions(S *tabSubs, n int) {
	if n == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}
	for i := 0; i < n; i++ {
		// Otomatis nonaktifkan langganan jika endDate sudah lewat
		// Karena S sekarang adalah *tabSubs, modifikasi ini akan permanen.
		if !S[i].endDate.IsZero() && time.Now().After(S[i].endDate) && S[i].active {
			S[i].active = false // Modifikasi data asli melalui pointer
			fmt.Printf("Status langganan '%s' telah diubah menjadi TIDAK AKTIF karena sudah berakhir.\n", S[i].name)
		}
		viewDetails(S[i]) // Pass copy dari struct individual
	}
}

// prosedur menampilkan detail dari satu langganan
func viewDetails(subs subscription) {
	fmt.Printf("\n- %s (%s)\n", subs.name, subs.category)
	fmt.Printf("  Biaya: Rp%.2f/%s\n", subs.monthlyCost, subs.frequency)
	fmt.Printf("  Pembayaran: %s via %s\n", subs.paymentDate.Format("2006-01-02"), subs.paymentMethod)
	if !subs.endDate.IsZero() {
		fmt.Printf("  Berakhir Pada: %s\n", subs.endDate.Format("2006-01-02"))
	}
	if subs.notes != "" {
		fmt.Printf("  Catatan: %s\n", subs.notes)
	}
	fmt.Printf("  Status: %v\n", subs.active)
}

// prosedur mencari langganan berdasarkan nama menggunakan binary search
func searchByName(S *tabSubs, n int) {
	var query string
	fmt.Print("Masukkan nama langganan: ")
	fmt.Scan(&query)

	sortByName(S, n) // Pastikan terurut sebelum binary search

	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if S[mid].name == query {
			viewDetails(S[mid])
			return
		} else if S[mid].name < query {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Langganan tidak ditemukan.")
}

// prosedur mengurutkan langganan berdasarkan nama yang digunakan dalam prosedur searchByName
func sortByName(S *tabSubs, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if S[i].name > S[j].name {
				S[i], S[j] = S[j], S[i]
			}
		}
	}
}

// prosedur mencari langganan berdasarkan kategori menggunakan sequential search
func searchByCategory(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	var cat string
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&cat)

	found := false
	for i := 0; i < n; i++ {
		if S[i].category == cat {
			viewDetails(S[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

// prosedur mengurutkan langganan berdasarkan biaya bulanan menggunakan selection sort
func sortByCostSelection(S *tabSubs, n int) {
	var x string
	fmt.Print("Urutkan berdasarkan biaya (menaik/menurun): ")
	fmt.Scan(&x)

	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (x == "menaik" && S[j].monthlyCost < S[idx].monthlyCost) ||
				(x == "menurun" && S[j].monthlyCost > S[idx].monthlyCost) {
				idx = j
			}
		}
		S[i], S[idx] = S[idx], S[i]
	}
	viewSubscriptions(S, n) // Panggil dengan pointer (S sudah pointer)
}

// prosedur mengurutkan langganan berdasarkan tanggal pembayaran menggunakan insertion sort
func sortByDateInsertion(S *tabSubs, n int) {
	var x string
	fmt.Print("Urutkan berdasarkan tanggal (menaik/menurun): ")
	fmt.Scan(&x)

	for i := 1; i < n; i++ {
		key := S[i]
		j := i - 1
		for j >= 0 &&
			((x == "menaik" && S[j].paymentDate.After(key.paymentDate)) ||
				(x == "menurun" && S[j].paymentDate.Before(key.paymentDate))) {
			S[j+1] = S[j]
			j--
		}
		S[j+1] = key
	}
	viewSubscriptions(S, n) // Panggil dengan pointer (S sudah pointer)
}

// prosedur menampilkan total pengeluaran bulanan dan rekomendasi penghematan langganan (apabila jumlah langganan > 3)
func calculateTotalAndRecommend(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	var total float64
	var maxCost subscription

	for i := 0; i < n; i++ {
		// Perhitungkan frekuensi pembayaran untuk total bulanan
		costForCalculation := S[i].monthlyCost
		if S[i].frequency == "tahunan" {
			costForCalculation /= 12.0
		} else if S[i].frequency == "mingguan" {
			costForCalculation *= (365.25 / 7) / 12 // Perkiraan bulanan dari mingguan
		}
		total += costForCalculation
	}

	fmt.Printf("Total pengeluaran bulanan (perkiraan): Rp%.2f\n", total) // Ubah label

	if n > 3 {
		maxCost = maxMonthlyCost(S, n)
		fmt.Println("Rekomendasi Penghematan:")
		fmt.Printf("Pertimbangkan berhenti berlangganan '%s' dengan biaya Rp%.2f/bulan\n", maxCost.name, maxCost.monthlyCost)
	}
}

// function untuk mencari biaya bulanan paling tinggi dan mengembalikan tipe data struct subscription
func maxMonthlyCost(S tabSubs, n int) subscription { // S di sini tidak perlu pointer karena hanya membaca
	if n == 0 {
		return subscription{} // Kembalikan nilai kosong jika tidak ada data
	}
	idx := 0
	for i := 1; i < n; i++ {
		if S[i].monthlyCost > S[idx].monthlyCost {
			idx = i
		}
	}
	return S[idx]
}

// prosedur menampilkan biaya bulanan paling besar
func highestCost(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	if n == 0 {
		fmt.Println("Tidak ada data langganan.")
		return
	}
	max := maxMonthlyCost(S, n)
	fmt.Println("Langganan dengan biaya tertinggi:")
	viewDetails(max)
}

// prosedur pengingat pembayaran dalam rentang hari ke depan yang ditentukan pengguna
func reminder(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	today := time.Now()
	reminded := false

	// Input jumlah hari pengingat
	var days int
	fmt.Print("Ingin diingatkan berapa hari ke depan? ")
	fmt.Scan(&days)

	if days < 0 {
		fmt.Println("Jumlah hari tidak boleh negatif.")
		return
	}

	fmt.Printf("\n--- Pengingat Pembayaran Dalam %d Hari ke Depan ---\n", days)
	for i := 0; i < n; i++ {
		if S[i].active {
			nextPayment := S[i].paymentDate // Mulai dari tanggal pembayaran terakhir

			// Loop untuk mencari tanggal pembayaran berikutnya yang jatuh di masa depan
			// dan dalam rentang hari yang diminta dari hari ini
			// Penyesuaian untuk paymentDate.Day() yang bisa lebih besar dari hari ini
			for nextPayment.Before(today) || (nextPayment.Month() == today.Month() && nextPayment.Day() < today.Day() && S[i].paymentDate.Day() != nextPayment.Day()) {
				switch S[i].frequency {
				case "bulanan":
					nextPayment = nextPayment.AddDate(0, 1, 0)
				case "tahunan":
					nextPayment = nextPayment.AddDate(1, 0, 0)
				case "mingguan":
					nextPayment = nextPayment.AddDate(0, 0, 7)
				default:
					nextPayment = nextPayment.AddDate(0, 1, 0) // Default bulanan
				}
			}

			diff := nextPayment.Sub(today).Hours() / 24
			if diff >= 0 && diff <= float64(days) { // Cek dalam rentang hari yang diminta
				reminded = true
				fmt.Printf("\n- %s (%s)\n", S[i].name, S[i].category)
				fmt.Printf("  Tanggal Pembayaran: %s\n", nextPayment.Format("2006-01-02"))
				fmt.Printf("  Metode: %s\n", S[i].paymentMethod)
				fmt.Printf("  Biaya: Rp%.2f\n", S[i].monthlyCost)
			}
		}
	}

	if !reminded {
		fmt.Printf("Tidak ada langganan yang harus dibayar dalam %d hari ke depan.\n", days)
	}
}

// prosedur menampilkan total pengeluaran per kategori
func displayCostByCategory(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	if n == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}

	costByCategory := make(map[string]float64) // Map untuk menyimpan total biaya per kategori

	for i := 0; i < n; i++ {
		// Perhitungkan biaya bulanan ekuivalen untuk total per kategori
		costForCategory := S[i].monthlyCost
		if S[i].frequency == "tahunan" {
			costForCategory /= 12.0
		} else if S[i].frequency == "mingguan" {
			costForCategory *= (365.25 / 7) / 12
		}
		costByCategory[S[i].category] += costForCategory
	}

	fmt.Println("\n--- Total Pengeluaran per Kategori (Perkiraan Bulanan) ---")
	if len(costByCategory) == 0 {
		fmt.Println("Tidak ada data kategori yang ditemukan.")
		return
	}
	for category, totalCost := range costByCategory {
		fmt.Printf("Kategori '%s': Rp%.2f\n", category, totalCost)
	}
}

// prosedur memproyeksikan pengeluaran bulanan di masa depan
func projectFutureExpenses(S tabSubs, n int) { // S di sini tidak perlu pointer karena hanya membaca
	if n == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}

	var months int
	fmt.Print("Proyeksikan pengeluaran untuk berapa bulan ke depan? ")
	fmt.Scan(&months)

	if months <= 0 {
		fmt.Println("Jumlah bulan harus lebih dari 0.")
		return
	}

	totalProjectedCost := 0.0
	for i := 0; i < n; i++ {
		monthlyEquivalentCost := S[i].monthlyCost
		if S[i].frequency == "tahunan" {
			monthlyEquivalentCost /= 12.0
		} else if S[i].frequency == "mingguan" {
			monthlyEquivalentCost *= (365.25 / 7) / 12
		}
		totalProjectedCost += monthlyEquivalentCost * float64(months)
	}
	fmt.Printf("\nProyeksi pengeluaran untuk %d bulan ke depan: Rp%.2f\n", months, totalProjectedCost)
}
