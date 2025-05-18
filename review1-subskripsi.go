package main

import (
	"fmt"
	"sort"
	"time"
)

type Subscription struct {
	Name          string
	Category      string
	MonthlyCost   float64
	PaymentDate   time.Time
	PaymentMethod string
	Active        bool
}

var subscriptions []Subscription

func main() {
	var choice int
	
	for {
		fmt.Println("\n=== Aplikasi Manajemen Subskripsi ===")
		fmt.Println("1. Tambah Langganan")
		fmt.Println("2. Lihat Semua Langganan")
		fmt.Println("3. Cari Langganan (Sequential)")
		fmt.Println("4. Cari Langganan (Binary)")
		fmt.Println("5. Urutkan by Biaya (Selection Sort)")
		fmt.Println("6. Urutkan by Tanggal (Insertion Sort)")
		fmt.Println("7. Total Pengeluaran & Rekomendasi")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		
		fmt.Scanln(&choice)
		
		if choice == 1 {
			addSubscription()
		} else if choice == 2 {
			viewSubscriptions()
		} else if choice == 3 {
			searchSequential()
		} else if choice == 4 {
			searchBinary()
		} else if choice == 5 {
			sortByCostSelection()
		} else if choice == 6 {
			sortByDateInsertion()
		} else if choice == 7 {
			calculateTotalAndRecommend()
		} else if choice == 8 {
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func addSubscription() {
	var name, category, method, dateStr string
	var cost float64
	
	fmt.Print("Nama Layanan: ")
	fmt.Scanln(&name)
	
	fmt.Print("Kategori: ")
	fmt.Scanln(&category)
	
	fmt.Print("Biaya Bulanan: ")
	fmt.Scanln(&cost)
	
	fmt.Print("Tanggal Pembayaran (YYYY-MM-DD): ")
	fmt.Scanln(&dateStr)
	
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Format tanggal salah!")
		return
	}
	
	fmt.Print("Metode Pembayaran: ")
	fmt.Scanln(&method)
	
	subscriptions = append(subscriptions, Subscription{
		Name:          name,
		Category:      category,
		MonthlyCost:   cost,
		PaymentDate:   date,
		PaymentMethod: method,
		Active:        true,
	})
	
	fmt.Println("Langganan berhasil ditambahkan!")
}

func viewSubscriptions() {
	if len(subscriptions) == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}
	
	for i, sub := range subscriptions {
		fmt.Printf("\n%d. %s (%s)\n", i+1, sub.Name, sub.Category)
		fmt.Printf("   Biaya: Rp%.2f/bulan\n", sub.MonthlyCost)
		fmt.Printf("   Pembayaran berikutnya: %s via %s\n", 
			sub.PaymentDate.Format("2006-01-02"), sub.PaymentMethod)
		fmt.Printf("   Status: %v\n", sub.Active)
	}
}

func searchSequential() {
	fmt.Print("Cari berdasarkan nama: ")
	var query string
	fmt.Scanln(&query)
	
	found := false
	i := 0
	for i < len(subscriptions) && !found {
		if subscriptions[i].Name == query {
			printSubscription(subscriptions[i])
			found = true
		}
		i++
	}
	
	if !found {
		fmt.Println("Langganan tidak ditemukan.")
	}
}

func searchBinary() {
	if len(subscriptions) == 0 {
		fmt.Println("Belum ada langganan.")
		return
	}
	
	sort.Slice(subscriptions, func(i, j int) bool {
		return subscriptions[i].Name < subscriptions[j].Name
	})
	
	fmt.Print("Cari berdasarkan nama: ")
	var query string
	fmt.Scanln(&query)
	
	low := 0
	high := len(subscriptions) - 1
	found := false
	
	for low <= high && !found {
		mid := (low + high) / 2
		
		if subscriptions[mid].Name == query {
			printSubscription(subscriptions[mid])
			found = true
		} else if subscriptions[mid].Name < query {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	if !found {
		fmt.Println("Langganan tidak ditemukan.")
	}
}

func sortByCostSelection() {
	n := len(subscriptions)
	
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if subscriptions[j].MonthlyCost > subscriptions[minIdx].MonthlyCost {
				minIdx = j
			}
		}
		subscriptions[i], subscriptions[minIdx] = subscriptions[minIdx], subscriptions[i]
	}
	
	fmt.Println("Langganan telah diurutkan berdasarkan biaya tertinggi:")
	viewSubscriptions()
}

func sortByDateInsertion() {
	n := len(subscriptions)
	
	for i := 1; i < n; i++ {
		key := subscriptions[i]
		j := i - 1
		
		for j >= 0 && subscriptions[j].PaymentDate.After(key.PaymentDate) {
			subscriptions[j+1] = subscriptions[j]
			j = j - 1
		}
		subscriptions[j+1] = key
	}
	
	fmt.Println("Langganan telah diurutkan berdasarkan tanggal pembayaran:")
	viewSubscriptions()
}

func calculateTotalAndRecommend() {
	var total float64
	var mostExpensive Subscription
	maxCost := 0.0
	
	for _, sub := range subscriptions {
		if sub.Active {
			total += sub.MonthlyCost
			if sub.MonthlyCost > maxCost {
				maxCost = sub.MonthlyCost
				mostExpensive = sub
			}
		}
	}
	
	fmt.Printf("\nTotal pengeluaran bulanan: Rp%.2f\n", total)
	
	if len(subscriptions) > 3 {
		fmt.Println("\nRekomendasi untuk menghemat:")
		fmt.Printf("Pertimbangkan untuk berhenti berlangganan %s (Rp%.2f/bulan)\n", 
			mostExpensive.Name, mostExpensive.MonthlyCost)
	}
}

func printSubscription(sub Subscription) {
	fmt.Println("\nDetail Langganan:")
	fmt.Printf("Nama: %s\n", sub.Name)
	fmt.Printf("Kategori: %s\n", sub.Category)
	fmt.Printf("Biaya: Rp%.2f/bulan\n", sub.MonthlyCost)
	fmt.Printf("Pembayaran berikutnya: %s via %s\n", 
		sub.PaymentDate.Format("2006-01-02"), sub.PaymentMethod)
	fmt.Printf("Status: %v\n", sub.Active)
}