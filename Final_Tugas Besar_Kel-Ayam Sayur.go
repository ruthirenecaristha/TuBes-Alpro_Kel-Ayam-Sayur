package main

import (
	"fmt"
	"time" // untuk tipe data waktu
)

const NMAX int = 100 // batas maksimal jumlah isi array

// tipe data struct untuk array
type subscription struct {
	name, category string    // nama dan kategori layanan
	monthlyCost    float64   // biaya bulanan langganan
	paymentDate    time.Time // waktu mulai pembayaran langganan
	paymentMethod  string    // metode pembayaran langganan
	active         bool      // status langganan (aktif/tidak)
}
type tabSubs [NMAX]subscription // array tabSubs dengan tipe data subscription

func main() {
	var choice, nData int
	var dataSub tabSubs

	for {
		menu()
		fmt.Print("Pilih menu (1-12): ")
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
			viewSubscriptions(dataSub, nData)
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
		case 12:
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
	fmt.Println("12. Keluar")
}

// prosedur menambah langganan
func addSubscription(S *tabSubs, n *int) {
	var name, category, method, dateStr string
	var cost float64
	var x int

	fmt.Print("Masukkan jumlah layanan yang ingin ditambah: ")
	fmt.Scan(&x)

	for i := 0; i < x; i++ {
		fmt.Print("Nama Layanan: ")
		fmt.Scan(&name)
		fmt.Print("Kategori: ")
		fmt.Scan(&category)
		fmt.Print("Biaya Bulanan: ")
		fmt.Scan(&cost)
		fmt.Print("Tanggal Pembayaran (YYYY-MM-DD): ")
		fmt.Scan(&dateStr)

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil || !dateIsValid(date.Year(), date.Month(), date.Day()) {
			fmt.Println("Format tanggal salah atau tidak valid!")
			return
		}

		fmt.Print("Metode Pembayaran: ")
		fmt.Scan(&method)

		S[*n] = subscription{
			name:          name,
			category:      category,
			monthlyCost:   cost,
			paymentDate:   date,
			paymentMethod: method,
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
	var name, newDate string
	fmt.Print("Masukkan nama langganan: ")
	fmt.Scan(&name)

	var ganti string
	fmt.Print("Apa yang ingin diedit? (nama/kategori/tanggal/biaya/metode bayar): ")
	fmt.Scan(&ganti)

	idx := seqSearch(*S, n, name)
	if idx == -1 {
		fmt.Println("Langganan tidak ditemukan.")
		return
	}

	var namaBaru string
	if ganti == "nama" {
		fmt.Print("Masukkan nama langganan baru: ")
		fmt.Scan(&namaBaru)
		S[idx].name = namaBaru
	}

	var kategoriBaru string
	if ganti == "kategori" {
		fmt.Print("Masukkan kategori langganan baru: ")
		fmt.Scan(&kategoriBaru)
		S[idx].category = kategoriBaru
	}

	if ganti == "tanggal" {
		fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
		fmt.Scan(&newDate)

		date, err := time.Parse("2006-01-02", newDate)
		if err != nil {
			fmt.Println("Format tanggal salah!")
			return
		}

		S[idx].paymentDate = date
	}

	var biayaBaru float64
	if ganti == "biaya" {
		fmt.Print("Masukkan jumlah biaya langganan baru: ")
		fmt.Scan(&biayaBaru)
		S[idx].monthlyCost = biayaBaru
	}

	var metodeBaru string
	if ganti == "metode" {
		fmt.Print("Masukkan metode pembayaran baru: ")
		fmt.Scan(&metodeBaru)
		S[idx].paymentMethod = metodeBaru
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
func viewSubscriptions(S tabSubs, n int) {
	if n == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}
	for i := 0; i < n; i++ {
		viewDetails(S[i])
	}
}

// prosedur menampilkan detail dari satu langganan
func viewDetails(subs subscription) {
	fmt.Printf("\n- %s (%s)\n", subs.name, subs.category)
	fmt.Printf("  Biaya: Rp%.2f/bulan\n", subs.monthlyCost)
	fmt.Printf("  Pembayaran: %s via %s\n", subs.paymentDate.Format("2006-01-02"), subs.paymentMethod)
	fmt.Printf("  Status: %v\n", subs.active)
}

// prosedur mencari langganan berdasarkan nama menggunakan binary search
func searchByName(S *tabSubs, n int) {
	var query string
	fmt.Print("Masukkan nama langganan: ")
	fmt.Scan(&query)

	sortByName(S, n)

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
func searchByCategory(S tabSubs, n int) {
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
	viewSubscriptions(*S, n)
}

// prosedur mengurutkan langganan berdasarkan tanggal pembayaran menggunakan insertion sort
func sortByDateInsertion(S *tabSubs, n int) {
	var pass, i int
	var temp subscription

	var order string
	fmt.Print("Urutkan berdasarkan tanggal (menaik/menurun): ")
	fmt.Scan(&order)

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = S[pass]
		for i > 0 &&
			((order == "menaik" && temp.paymentDate.Before(S[i-1].paymentDate)) ||
				(order == "menurun" && temp.paymentDate.After(S[i-1].paymentDate))) {
			S[i] = S[i-1]
			i--
		}
		S[i] = temp
		pass++
	}

	viewSubscriptions(*S, n)
}

// prosedur menampilkan total pengeluaran bulanan dan rekomendasi penghematan langganan (apabila jumlah langganan > 3)
func calculateTotalAndRecommend(S tabSubs, n int) {
	var total float64
	var maxCost subscription

	for i := 0; i < n; i++ {
		total += S[i].monthlyCost
	}

	fmt.Printf("Total pengeluaran bulanan: Rp%.2f\n", total)

	if n > 3 {
		maxCost = maxMonthlyCost(S, n)
		fmt.Println("Rekomendasi Penghematan:")
		fmt.Printf("Pertimbangkan berhenti berlangganan '%s' dengan biaya Rp%.2f\n", maxCost.name, maxCost.monthlyCost)
	}
}

// function untuk mencari biaya bulanan paling tinggi dan mengembalikan tipe data struct subscription
func maxMonthlyCost(S tabSubs, n int) subscription {
	idx := 0
	for i := 1; i < n; i++ {
		if S[i].monthlyCost > S[idx].monthlyCost {
			idx = i
		}
	}
	return S[idx]
}

// prosedur menampilkan biaya bulanan paling besar
func highestCost(S tabSubs, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data langganan.")
		return
	}
	max := maxMonthlyCost(S, n)
	fmt.Println("Langganan dengan biaya tertinggi:")
	viewDetails(max)
}

// prosedur pengingat pembayaran dalam 7 hari ke depan
func reminder(S tabSubs, n int) {
	today := time.Now()
	reminded := false

	fmt.Println("\n--- Pengingat Pembayaran Dalam 7 Hari ---")
	for i := 0; i < n; i++ {
		if S[i].active {
			nextPayment := time.Date(today.Year(), today.Month(), S[i].paymentDate.Day(), 0, 0, 0, 0, time.Local)
			if nextPayment.Before(today) {
				nextPayment = nextPayment.AddDate(0, 1, 0)
			}

			diff := nextPayment.Sub(today).Hours() / 24
			if diff >= 0 && diff <= 7 {
				reminded = true
				fmt.Printf("\n- %s (%s)\n", S[i].name, S[i].category)
				fmt.Printf("  Tanggal Pembayaran: %s\n", nextPayment.Format("2006-01-02"))
				fmt.Printf("  Metode: %s\n", S[i].paymentMethod)
				fmt.Printf("  Biaya: Rp%.2f\n", S[i].monthlyCost)
			}
		}
	}

	if !reminded {
		fmt.Println("Tidak ada langganan yang harus dibayar dalam 7 hari ke depan.")
	}
}
