package main

import (
	"fmt"
	"sync"
)

//CODE SOURCE INSPIRATION :
//https://www.tutorialspoint.com/golang-program-to-compute-all-prime-numbers-up-to-a-given-number-using-concurrency

func cekKelipatan(kelipatan, input int) bool {
	if input%kelipatan == 0 {
		return true
	}
	return false
}

func cariKelipatan(mulai, akhir, kelipatan int, wg *sync.WaitGroup, hasil chan<- int) {
	defer wg.Done()
	for i := mulai; i <= akhir; i++ {
		if cekKelipatan(kelipatan, i) {
			hasil <- i
			fmt.Println("Angka", i, "adalah kelipatan", kelipatan)
		} else {
			fmt.Println("Angka", i)
		}
	}
}

func cetakHasil(hasil <-chan int, kelipatan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\n--Rangkuman Nilai Kelipatan %d--\n", kelipatan)
	for nilai := range hasil {
		fmt.Println(nilai)
	}
}

func main() {
	var wg sync.WaitGroup
	hasilKelipatan := make(chan int, 100)

	start := 1
	end := 150
	kelipatan := 17

	numberOfGoroutine := 3

	segmentSize := (end - start + 1) / numberOfGoroutine

	fmt.Printf("\noo- Mencari hasil kelipatan %d -oo\n", kelipatan)
	for i := 0; i < numberOfGoroutine; i++ {
		wg.Add(1)
		go cariKelipatan(start+i*segmentSize, start+(i+1)*segmentSize-1, kelipatan, &wg, hasilKelipatan)
	}

	wg.Wait()
	close(hasilKelipatan)

	wg.Add(1)
	go cetakHasil(hasilKelipatan, kelipatan, &wg)

	wg.Wait()
}
