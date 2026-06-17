package main

import (
	"fmt"
	"os"
)

const NMAX int = 999
type menu struct {
	nama      string
	harga     int
	komposisi string
	kategori  string
	ketersediaan string
}
type dataMenu [NMAX]menu

func main() {
	var listMenu dataMenu
	var nAwal,posisi,percobaan int
	var password string

	//membaca data awal
	f, _ := os.Open("data.txt")
	//memproses data awal
	fmt.Fscan(f, &nAwal)
	for i := 0; i < nAwal; i++ {
		fmt.Fscan(f, &listMenu[i].nama, &listMenu[i].harga, &listMenu[i].komposisi, &listMenu[i].kategori,&listMenu[i].ketersediaan)
	}

	//memulai program
	fmt.Println("Selamat datang diCafe-menu (^o^)")
	posisi = 1
	for posisi == 1 {
		fmt.Println("Silahkan pilih posisi anda")
		fmt.Println("1.Admin")
		fmt.Println("2.Pelanggan")
		fmt.Println("3.Exit")
		fmt.Print("pilih(1/2/3):")
		fmt.Scan(&posisi)
		if posisi == 1 {
			percobaan = 1
			password = ""
			for password != "admin123" && percobaan<=3 {
				fmt.Print("Masukkan password: ")
				fmt.Scan(&password)
				if password == "admin123" {
					fmt.Println()
					menuAdmin(&nAwal, &listMenu)
				}else {
					fmt.Println("password salah")
					percobaan++
				}
				
			}
		}else if posisi == 2 {
			menuPelanggan(nAwal, &listMenu)
		}else if posisi == 3 {
			fmt.Println()
		}else {
			fmt.Println("Posisi tidak tersedia")
		}
	}
	fmt.Println("Sampai jumpa (^_^)")
}

func menuAdmin(nAwal *int, listMenu *dataMenu) {
	var menuAksi int
	var jumlahMenu int
	var nomor int

	menuAksi = 1
	for menuAksi != 6 {
		//memilih aksi
		fmt.Println("Silahkan pilih angka untuk melakukan aksi yang tersedia")
		fmt.Println("1.Menampilkan menu")
		fmt.Println("2.Menambahkan menu")
		fmt.Println("3.Mengubah menu")
		fmt.Println("4.Mengurangi menu")
		fmt.Println("5.Menampilkan statistik menu")
		fmt.Println("6.Exit")
		fmt.Print("pilih(1/2/3/4/5/6):")
		fmt.Scan(&menuAksi)
		if menuAksi == 1 {
			fmt.Println("Menampilkan menu yang tersedia")
			tampilMenu(*listMenu)
		} else if menuAksi == 2 {
			fmt.Println("Masukkan jumlah menu yang ingin ditambahkan")
			fmt.Scan(&jumlahMenu)
			//menghitung jumlah menu yang ditambahkan
			*nAwal = *nAwal + jumlahMenu
			fmt.Println("Masukkan menu yang ingin anda tambahkan (tanpa spasi atau gunakan _ )")
			fmt.Println("*Note: Kategori diisi dengan coffe atau noncoffe")
			for i := 1; i <= jumlahMenu; i++ {
				tambahMenu(listMenu)
			}
			tampilMenu(*listMenu)
		} else if menuAksi == 3 {
			fmt.Println("Pilih menu nomor berapa yang ingin anda ubah")
			fmt.Scan(&nomor)
			//mengecek apakah ada menu pada nomor tersebut
			if nomor <= 0 || nomor > *nAwal {
				fmt.Println("Menu tidak tersedia")
			} else {
				ubahMenu(nomor-1, listMenu)
				tampilMenu(*listMenu)
			}
		} else if menuAksi == 4 {
			fmt.Println("Pilih menu nomor berapa yang ingin anda hilangkan")
			fmt.Scan(&nomor)
			//mengecek apakah ada menu pada nomor tersebut
			if nomor <= 0 || nomor > *nAwal{
				fmt.Println("Menu tidak tersedia")
			} else {
				tampilMenu(*listMenu)
				hapusMenu(nomor-1, listMenu)
				*nAwal--
			}
		} else if menuAksi == 5 {
			statistikMenu(nAwal, listMenu)
		} else if menuAksi == 6 {
			fmt.Println("Kembali ke menu utama")
		}else {
			fmt.Println("Aksi tidak tersedia (!_!)")
		}
	}
}

func menuPelanggan(nAwal int, listMenu *dataMenu) {
	var menuAksi int
	var listAwal dataMenu
	
	fmt.Println("Menampilkan menu yang tersedia")
	tampilMenu(*listMenu)
	listAwal = *listMenu
	for menuAksi != 5 {
		fmt.Println("Silahkan pilih angka untuk melakukan aksi yang tersedia")
		fmt.Println("1.Mengurutkan menu dari harga termurah menggunakan Selection sort")
		fmt.Println("2.Mengurutkan menu dari harga termurah menggunakan Insertion sort")
		fmt.Println("3.Mencari menu berdasarkan kategori dengan sequential Search")
		fmt.Println("4.Mencari menu berdasarkan kategori dengan binery search")
		fmt.Println("5.Exit")
		fmt.Print("pilih(1/2/3/4/5):")
		fmt.Scan(&menuAksi)
		if menuAksi == 1 {
			mengurutkanMenuSelection(nAwal, listMenu)
			tampilMenu(*listMenu)
			*listMenu = listAwal
		}else if menuAksi == 2 {
			mengurutkanMenuInsertion(nAwal, listMenu)
			tampilMenu(*listMenu)
			*listMenu = listAwal
		}else if menuAksi == 3 {
			mencariDenganKategori(nAwal, listMenu)
		}else if menuAksi == 4 {
			fmt.Println("Mengurutkan menu berdasarkan kategori secara ascending")
			mengurutkanKategoriSelection(nAwal,listMenu)
			tampilMenu(*listMenu)
			mencariKategoriAscending(nAwal,listMenu)
			*listMenu = listAwal
		}else if menuAksi == 5 {
			fmt.Println()
		}else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func mengurutkanMenuSelection(n int, listMenu *dataMenu) {
	var i, j, select_i int
	var temp menu

	for i = 0; i < n; i++ {
		select_i = i
		for j = i + 1; j < n; j++ {
			if listMenu[j].harga < listMenu[select_i].harga {
				select_i = j
			}
		}

		temp = listMenu[i]
		listMenu[i] = listMenu[select_i]
		listMenu[select_i] = temp
	}
}

func mengurutkanMenuInsertion(n int, listMenu *dataMenu) {
	var i, j int
	var temp menu

	// Insertion sort biasanya dimulai dari indeks ke-1
	for i = 1; i < n; i++ {
		temp = listMenu[i] // Menyimpan menu saat ini sebagai patokan (elemen yang akan disisipkan)
		j = i - 1

		// Membandingkan dan menggeser elemen sebelumnya.
		// Kondisi listMenu[j].harga < temp.harga digunakan untuk descending (terbesar ke terkecil).
		// Jika ingin ascending, kondisinya diubah menjadi listMenu[j].harga > temp.harga
		for j >= 0 && listMenu[j].harga > temp.harga {
			listMenu[j+1] = listMenu[j] // Geser elemen ke kanan
			j--
		}

		// Menempatkan elemen patokan (temp) pada posisi yang tepat
		listMenu[j+1] = temp
	}
}

func mengurutkanKategoriSelection(n int, listMenu *dataMenu) {
	var i, j, select_i int
	var temp menu

	for i = 0; i < n; i++ {
		select_i = i
		for j = i + 1; j < n; j++ {
			if listMenu[j].kategori < listMenu[select_i].kategori {
				select_i = j
			}
		}

		temp = listMenu[i]
		listMenu[i] = listMenu[select_i]
		listMenu[select_i] = temp
	}
}

func mencariDenganKategori(n int, listMenu *dataMenu) {
	var kategori string
	var listPencarian dataMenu
	var i, n_listPencarian int

	fmt.Printf("Kategori: ")
	fmt.Scan(&kategori)

	n_listPencarian = 0

	for i = 0; i < n; i++ {
		if listMenu[i].kategori == kategori {
			listPencarian[n_listPencarian] = listMenu[i]
			n_listPencarian = n_listPencarian + 1
		}
	}

	tampilMenu(listPencarian)
}

func mencariKategoriAscending(n int, listMenu *dataMenu) {
	var kategori string
	var listPencarian dataMenu
	var n_listPencarian int = 0
	var kiri, kanan, tengah, ditemukan int

	fmt.Printf("Masukkan Kategori yang dicari: ")
	fmt.Scan(&kategori)

	kiri = 0
	kanan = n - 1
	ditemukan = -1

	// Langkah 1: Binary Search untuk menemukan satu kecocokan
	for kiri <= kanan && ditemukan == -1 {
		tengah = (kiri + kanan) / 2
		if listMenu[tengah].kategori == kategori {
			ditemukan = tengah
		} else if listMenu[tengah].kategori < kategori {
			kiri = tengah + 1 // Cari ke kanan
		} else {
			kanan = tengah - 1 // Cari ke kiri
		}
	}

	// Langkah 2: Jika ditemukan, kumpulkan semua menu dengan kategori yang sama di sekitarnya
	if ditemukan != -1 {
		// Cek ke arah kiri dari indeks yang ditemukan
		i := ditemukan
		for i >= 0 && listMenu[i].kategori == kategori {
			listPencarian[n_listPencarian] = listMenu[i]
			n_listPencarian++
			i--
		}

		// Cek ke arah kanan dari indeks yang ditemukan
		i = ditemukan + 1
		for i < n && listMenu[i].kategori == kategori {
			listPencarian[n_listPencarian] = listMenu[i]
			n_listPencarian++
			i++
		}
		
		tampilMenu(listPencarian)
	} else {
		fmt.Println("Menu dengan kategori", kategori, "tidak ditemukan.")
	}
}

func tampilMenu(listMenu dataMenu) {
	
	i:=0
	fmt.Println("------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-12s | %-30s | %-15s | %-15s |\n", "No", "Nama Menu", "Harga","Komposisi", "Kategori", "Ketersediaan")
	fmt.Println("------------------------------------------------------------------------------------------------------------------")
	for listMenu[i].nama != "" {
		fmt.Printf("| %-3d | %-20s | Rp%-10d | %-30s | %-15s | %-15s |\n",i+1, listMenu[i].nama, listMenu[i].harga, listMenu[i].komposisi, listMenu[i].kategori,listMenu[i].ketersediaan)
		i++
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------------")

}

func tambahMenu(listMenu *dataMenu) {
	var found int
	var nama, komposisi, kategori, ketersediaan string
	var harga int
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)
	fmt.Print("Harga menu: Rp")
	fmt.Scan(&harga)
	fmt.Print("Komposisi menu: ")
	fmt.Scan(&komposisi)
	fmt.Print("Kategori menu: ")
	fmt.Scan(&kategori)
	fmt.Print("ketersediaan menu: ")
	fmt.Scan(&ketersediaan)
	fmt.Println("")
	found = -1
	i := 0
	for found == -1 {
		if listMenu[i].nama == "" {
			listMenu[i].nama = nama
			listMenu[i].harga = harga
			listMenu[i].komposisi = komposisi
			listMenu[i].kategori = kategori
			listMenu[i].ketersediaan = ketersediaan
			found = 1
		}
		i++
	}
}

func ubahMenu(idx int, listMenu *dataMenu) {
	var komponenMenu string
	var namaBaru, komposisiBaru, kategoriBaru string
	var hargaBaru int
	fmt.Print("Komponen yang ingin anda ubah(nama/harga/komposisi/kategori): ")
	fmt.Scan(&komponenMenu)
	if komponenMenu == "nama" {
		fmt.Println("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		listMenu[idx].nama = namaBaru
	} else if komponenMenu == "harga" {
		fmt.Print("Masukkan harga baru: Rp")
		fmt.Scan(&hargaBaru)
		listMenu[idx].harga = hargaBaru
	} else if komponenMenu == "komposisi" {
		fmt.Print("Masukkan komposisi yang baru: ")
		fmt.Scan(&komposisiBaru)
		listMenu[idx].komposisi = komposisiBaru
	} else if komponenMenu == "kategori" {
		fmt.Print("Masukkan kategori yang baru: ")
		fmt.Scan(&kategoriBaru)
		listMenu[idx].kategori = kategoriBaru
	} else {
		fmt.Print("Mohon masukkan komponen yang ada dan benar")
	}

}

func hapusMenu(idx int, listMenu *dataMenu) {
	for listMenu[idx].nama != "" {
		listMenu[idx].nama = listMenu[idx+1].nama
		listMenu[idx].harga = listMenu[idx+1].harga
		listMenu[idx].komposisi = listMenu[idx+1].komposisi
		listMenu[idx].kategori = listMenu[idx+1].kategori
		listMenu[idx].ketersediaan = listMenu[idx+1].ketersediaan
		idx++
	}
	tampilMenu(*listMenu)
}

func statistikMenu(nAwal *int, listMenu *dataMenu) {
	var daftarKategori [NMAX]string
	var jumlahPerKategori [NMAX]int
	var jumlahKategoriUnik int = 0 
	
	// Pengecekan jika data menu masih kosong menggunakan struktur if-else
	if *nAwal == 0 {
		fmt.Println("Belum ada menu yang terdaftar untuk dihitung statistiknya.")
	} else {
		totalHarga := 0

		// Membaca setiap menu satu per satu
		for i := 0; i < *nAwal; i++ {
			kategori := listMenu[i].kategori
			harga := listMenu[i].harga

			totalHarga += harga

			// Cek apakah kategori dari menu saat ini sudah pernah dicatat
			ditemukan := false
			
			// Loop berjalan selama j belum habis DAN belum ditemukan
			for j := 0; j < jumlahKategoriUnik && !ditemukan; j++ {
				if daftarKategori[j] == kategori {
					// Jika kategori sudah ada, cukup tambahkan jumlahnya
					jumlahPerKategori[j]++ 
					ditemukan = true // Mengubah ke true untuk menghentikan loop j di iterasi berikutnya
				}
			}

			// Jika kategori belum pernah dicatat, tambahkan sebagai kategori baru
			if !ditemukan && kategori != "" {
				daftarKategori[jumlahKategoriUnik] = kategori
				jumlahPerKategori[jumlahKategoriUnik] = 1
				jumlahKategoriUnik++ // Tambah total kategori unik yang ditemukan
			}
		}

		// Menghitung rata-rata harga
		rataRataHarga := float64(totalHarga) / float64(*nAwal)

		// Menampilkan hasil statistik
		fmt.Println("\n========================================")
		fmt.Println("             STATISTIK MENU             ")
		fmt.Println("========================================")
		fmt.Printf("Total Menu        : %d\n", *nAwal)
		
		fmt.Println("\nJumlah Menu per Kategori:")
		for i := 0; i < jumlahKategoriUnik; i++ {
			fmt.Printf("- %-15s : %d menu\n", daftarKategori[i], jumlahPerKategori[i])
		}
		
		fmt.Printf("\nRata-rata Harga   : Rp%.2f\n", rataRataHarga)
		fmt.Println("========================================")
	}
}         
