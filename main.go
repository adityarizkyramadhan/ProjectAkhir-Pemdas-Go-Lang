package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	tampilanUtama()
}

func formatMoney(uang int) {
	p := message.NewPrinter(language.Indonesian)
	p.Printf("Rp %d,00\n", uang)

}

func tampilanUtama() {
	fmt.Println("=========== Mesin Kasir ===========")
	fmt.Println("1. Masuk")
	fmt.Println("2. Keluar")
	var pil int
	fmt.Scanln(&pil)
	switch pil {
	case 1:
		menuMasuk(login())
	case 2:
		os.Exit(0)
	}

}

func login() bool {
	var username, password string
	fmt.Println("===== Login =====")
	fmt.Print("Masukkan username anda : ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&password)
	if username != "adityacafe" || password != "adityacafe" {
		return false
	}
	return true
}

func menuMasuk(login bool) {
	if !login {
		fmt.Println("Password salah")
		tampilanUtama()
	}
	fmt.Println("Menu Utama")
	fmt.Println("1. Beli makanan")
	var pilih int
	fmt.Print("Masukkan pilihan : ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		beliMakanan()
	case 2:
	}
}

func beliMakanan() {
	makanans := []string{"Martabak Manis", "Nasi Goreng", "Nasi Gudeg", "Nasi Padang"}
	makanansi := map[string]int{}
	makanansi["martabak manis"] = 20000
	makanansi["nasi goreng"] = 10000
	makanansi["nasi gudeg"] = 15000
	makanansi["nasi padang"] = 12000
	nomer := 1
	for k, v := range makanansi {
		fmt.Println("Makanan ke -", nomer)
		fmt.Println("Makanan :", k)
		fmt.Println("Harga   :", v)
		fmt.Println("=========================")
		nomer++
	}
	fmt.Println("Untuk keluar pilih angka 5")
	penyimpanbeli := make([]string, 5)
	penyimpanJumlah := make([]int, 5)
	for i := 0; i < len(penyimpanbeli); i++ {
		fmt.Print("Masukkan makanan yang ingin dibeli (Masukkan angka) : ")
		var pemilih int
		fmt.Scanln(&pemilih)
		if pemilih == 1 {
			penyimpanbeli[i] = makanans[0]
		} else if pemilih == 2 {
			penyimpanbeli[i] = makanans[1]
		} else if pemilih == 3 {
			penyimpanbeli[i] = makanans[2]
		} else if pemilih == 4 {
			penyimpanbeli[i] = makanans[3]
		} else if pemilih == 5 {
			penyimpanbeli[i] = "stop"
		}
		if penyimpanbeli[i] == "stop" {
			break
		}
		fmt.Print("Masukkan jumlah pembelian : ")
		fmt.Scanln(&penyimpanJumlah[i])
		penyimpanbeli[i] = strings.ToLower(penyimpanbeli[i])

	}
	fmt.Println(penyimpanbeli)
	harga := 0
	fmt.Println("Menu yang anda pesan : ")
	for i := 0; i < len(penyimpanbeli); i++ {
		if penyimpanbeli[i] == "stop" {
			break
		}
		for j := 0; j < len(makanans); j++ {
			temp := strings.ToLower(makanans[j])
			if penyimpanbeli[i] == temp {
				harga = harga + makanansi[penyimpanbeli[i]]*penyimpanJumlah[i]
				tempHar := harga
				fmt.Print(penyimpanbeli[i], ": ")
				formatMoney(makanansi[penyimpanbeli[i]])
				fmt.Println("Jumlah beli :", penyimpanJumlah[i])
				fmt.Print("Harga : ")
				formatMoney(tempHar)
				tempHar = 0
			}
		}
	}
	fmt.Print("Harga total : ")
	formatMoney(harga)

}

// func getInput() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	clean := strings.TrimSpace(input)
// 	return clean

// }

// func filterMenu(array []string, fungsi func(string) bool) (result []string) {
// 	for _, nama := range array {s
// 		if fungsi(nama) {
// 			result = append(result, nama)
// 		}
// 	}
// 	return result
// }
