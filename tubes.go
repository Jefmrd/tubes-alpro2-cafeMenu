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
	var menuUtama, menuAksi int
	var jumlahMenu int
	listMenu = dataMenu{
		{nama: "Nasi_Goreng", 
		harga: 15000, 
		komposisi: "Nasi, Telur, Kecap",
		kategori: "makanan utama",
		},
		{nama:      "Matcha_Latte",
        harga:     18000,
        komposisi: "Bubuk Matcha, Susu Segar",
        kategori:  "minuman",
		},
		{nama:      "Kentang_Goreng",
        harga:     12000,
        komposisi: "Kentang, Garam, Kaldu",
        kategori:  "makanan ringan",
		},
	}
	
	menuUtama = 1
	fmt.Println("Selamat datang diCafe-menu (^o^)")
	for menuUtama == 1 {
		//memilih aksi
		fmt.Println("Silahkan pilih angka untuk melakukan aksi yang tersedia")
		fmt.Println("1.Menampilkan menu")
		fmt.Println("2.Menambahkan menu")
		fmt.Println("3.Mengubah menu")
		fmt.Println("4.Mengurangi menu")
		fmt.Scan(&menuAksi)
		if menuAksi == 1 {
			fmt.Println("Menampilkan menu yang tersedia")
			tampilMenu(listMenu)
		}else if menuAksi == 2 {
			fmt.Println("Masukkan jumlah menu yang ingin ditambahkan")
			fmt.Scan(&jumlahMenu)
			fmt.Println("Masukkan menu yang ingin anda tambahkan (tanpa spasi atau gunakan _ )")
			fmt.Println("*Note: Kategori diisi dengan coffe atau noncoffe")
			for i:=1; i<=jumlahMenu;i++{
				tambahMenu(&listMenu)
			}
			tampilMenu(listMenu)
		}else if menuAksi == 3 {
			fmt.Println("Pilih menu nomor berapa yang ingin anda ubah")
			
		}else if menuAksi == 4 {
			fmt.Println("Pilih menu yang ingin anda hilangkan")
			tampilMenu(listMenu)
			hapusMenu(&listMenu)
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

func tambahMenu (listMenu*dataMenu){
	var found int
	var nama, komposisi, kategori string
	var harga int
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)
	fmt.Print("Harga menu: Rp")
	fmt.Scan(&harga)
	fmt.Print("Komposisi menu: ")
	fmt.Scan(&komposisi)
	fmt.Print("Kategori menu: ")
	fmt.Scan(&kategori)
	fmt.Println("")
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

func hapusMenu(listMenu*dataMenu){
	var namaMenu string
	var idxMenu int
	
	fmt.Print("Nama menu: ")
	fmt.Scan(&namaMenu)
	
	if namaMenu == "67" {
		idxMenu = -2
	}else {
		idxMenu = cariMenuDenganSequential(namaMenu, *listMenu)
	}
	
	if idxMenu == -1 {
		fmt.Println("Data menu tidak tersedia.\nSilahkan masukkan menu kembali atau ketik 67 untuk berhenti")
		hapusMenu(listMenu)
	}else if idxMenu == -2 {
		fmt.Println("Kembali ke menu utama")
	}else {
		for listMenu[idxMenu].nama != "" {
			listMenu[idxMenu].nama = listMenu[idxMenu + 1].nama
			listMenu[idxMenu].harga = listMenu[idxMenu+1].harga
			listMenu[idxMenu].komposisi = listMenu[idxMenu+1].komposisi
			listMenu[idxMenu].kategori = listMenu[idxMenu+1].kategori
			idxMenu++
		}
		tampilMenu(*listMenu)
	}
	
}

func cariMenuDenganSequential (nama string, listMenu dataMenu)int{
	var i, found int
	i = 0
	found = -1
	for found == -1 && listMenu[i].nama != "" {
		if nama == listMenu[i].nama {
			found = i
		}
		i++
	}
	return found
}