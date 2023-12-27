package main //exr1

import (
	"fmt"
	"sync"
)

//const workNumb = 8

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) int {
	b.mu.Lock()
	b.balance = b.balance + amount
	fmt.Println(b.balance)
	b.mu.Unlock()
	return b.balance
}

func (b *BankAccount) Withdraw(amount int) int {
	b.mu.Lock()
	if amount > b.balance {
		fmt.Println("Withdraft больше чем ваш баланс")
	} else {
		b.balance = b.balance - amount
	}
	b.mu.Unlock()
	return b.balance
}

func main() {
	b := BankAccount{
		balance: 10000,
	}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(i)
		go func() {
			b.Deposit(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			b.Withdraw(100000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.balance)
}
