package main
import "fmt"

type dataMenu [] kategori
type kategori struct {
	makananUtama menu
	makananRingan menu
	minuman menu
}
type menu struct {
	nama string
	harga int
	komposisi string
}

func main(){
	var listMenu dataMenu
	
	listMenu = dataMenu{
		{
				makananUtama: menu{nama: "Nasi Goreng", harga: 15000, komposisi: "Nasi, Telur, Kecap"},
				
		},
	}
	
	tambahNamaMenu = ""
	tambahHargaMenu = 0
	tambahKomposisiMenu = ""
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
			fmt.Println("")
		}else if menuAksi == 2 {
			
		}else if menuAksi == 3 {
			fmt.Println("Masukkan menu yang ingin anda tambahkan")
			fmt.Print("Nama menu: ")
			fmt.Scan(&tambahNamaMenu)
			fmt.Print("Harga menu: ")
			fmt.Scan(&tambahHargaMenu)
			fmt.Print("Komposisi menu: ")
			fmt.Scan(&tambahKomposisiMenu)
			tambahMenu(tambahNamaMenu, tambahHargaMenu, tambahKomposisiMenu, &)
			
		}else if menuAksi == 4 {
			
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

func tambahMenu ()