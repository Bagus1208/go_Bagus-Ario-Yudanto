package main

import (
	"fmt"
	"sync"
)

type wallet struct {
	m       sync.Mutex
	balance int
}

func (money *wallet) in(increase int) {
	money.m.Lock()
	money.balance += increase
	money.m.Unlock()
}

func (money *wallet) out(decrease int) {
	money.m.Lock()
	money.balance -= decrease
	money.m.Unlock()
}

func (money *wallet) checkbalance() int {
	money.m.Lock()
	balance := money.balance
	money.m.Unlock()
	return balance
}

func main() {
	var mywallet wallet
	var wait sync.WaitGroup

	// pemasukkan tiap tahun
	for i := 0; i < 12; i++ {
		wait.Add(1)
		go func() {
			for i := 0; i < 22; i++ {
				mywallet.in(300000)
			}
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Annual income : Rp.", mywallet.checkbalance())

	// pengeluaran tiap tahun
	for i := 0; i < 12; i++ {
		wait.Add(1)
		go func() {
			for i := 0; i < 30; i++ {
				mywallet.out(150000)
			}
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Annual saving : Rp.", mywallet.checkbalance())
	fmt.Println("\nDONE")
}
