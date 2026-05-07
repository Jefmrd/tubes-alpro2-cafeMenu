package main
import "fmt"


const NMAX int = 999
type dataMenu [NMAX] menu
type menu struct {
	nama string
	harga int
	komposisi string
	kategori string
}

func main(){
	var listMenu dataMenu
	var tambahNamaMenu, tambahKategori, tambahKomposisiMenu string
	var tambahHargaMenu int
	var menuUtama, menuAksi int
	listMenu = dataMenu{
		{nama: "Nasi Goreng", 
		harga: 15000, 
		komposisi: "Nasi, Telur, Kecap",
		kategori: "makanan utama",
		},
		{nama:      "Matcha Latte",
        harga:     18000,
        komposisi: "Bubuk Matcha, Susu Segar",
        kategori:  "minuman",
		},
		{nama:      "Kentang Goreng",
        harga:     12000,
        komposisi: "Kentang, Garam, Kaldu",
        kategori:  "makanan ringan",
		},
	}
	
	tambahNamaMenu = ""
	tambahHargaMenu = 0
	tambahKomposisiMenu = ""
	tambahKategori = ""
	menuUtama = 1
	fmt.Println("Selamat datang diCafe-menu (^o^)")
	for menuUtama == 1 {
		//memilih aksi
		fmt.Println("Silahkan pilih angka untuk melakukan aksi yang tersedia")
		fmt.Println("1.Menampilkan menu")
		fmt.Println("2.Menambahkan menu")
		fmt.Println("3.Mengubah menu")
		fmt.Println("4.Menangurangi menu")
		fmt.Scan(&menuAksi)
		if menuAksi == 1 {
			fmt.Println("Menampilkan menu yang tersedia")
			tampilMenu(listMenu)
		}else if menuAksi == 2 {
			fmt.Println("Masukkan menu yang ingin anda tambahkan (tanpa spasi atau gunakan _ )")
			fmt.Print("Nama menu: ")
			fmt.Scan(&tambahNamaMenu)
			fmt.Print("Harga menu: Rp")
			fmt.Scan(&tambahHargaMenu)
			fmt.Print("Komposisi menu: ")
			fmt.Scan(&tambahKomposisiMenu)
			fmt.Print("Kategori menu: ")
			fmt.Scan(&tambahKategori)
			tambahMenu(tambahNamaMenu, tambahKomposisiMenu, tambahKategori, tambahHargaMenu, &listMenu)
		}else if menuAksi == 3 {
			fmt.Println("")
		}else if menuAksi == 4 {
			fmt.Println("")
		}else {
			fmt.Println("Aksi tidak tersedia (!_!)")
		}
		
		
		//Penutup
		fmt.Println("Ada lagi yang ingin anda lakukan?")
		fmt.Println("1.Kembali menu utama")
		fmt.Println("2.Mengakhiri program")
		fmt.Scan(&menuUtama)
	}
	fmt.Println("Sampai jumpa (^_^)")
}

func tampilMenu(listMenu dataMenu){
	
	i:=0
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-12s | %-30s | %-15s |\n", "No", "Nama Menu", "Harga","Komposisi", "Kategori")
	fmt.Println("------------------------------------------------------------------------------------------------")
	for listMenu[i].nama != "" {
		fmt.Printf("| %-3d | %-20s | Rp%-10d | %-30s | %-15s |\n",i+1, listMenu[i].nama, listMenu[i].harga, listMenu[i].komposisi, listMenu[i].kategori)
		i++
	}
	fmt.Println("------------------------------------------------------------------------------------------------")
}

func tambahMenu (nama, komposisi, kategori string, harga int, listMenu*dataMenu){
	var found int
	
	found = -1
	i := 0
	for found == -1 {
		if listMenu[i].nama == "" {
			listMenu[i].nama = nama
			listMenu[i].harga = harga
			listMenu[i].komposisi = komposisi
			listMenu[i].kategori = kategori
			found = 1
		}
		i++
	}
	
}