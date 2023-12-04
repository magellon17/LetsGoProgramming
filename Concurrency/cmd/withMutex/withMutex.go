package main

import (
	"github.com/pkg/errors"
	"log"
	"sync"
	"time"
)

type PlayerWallet struct {
	coins int
	mu    sync.RWMutex
}

func (w *PlayerWallet) getCoins() int {
	w.mu.RLock()
	coins := w.coins
	w.mu.RUnlock()

	return coins
}

func (w *PlayerWallet) spendCoins(amount int) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.coins >= amount {

		time.Sleep(time.Millisecond)

		w.coins -= amount
		log.Printf("spent %d coins (remaining: %d)", amount, w.coins)
		return nil
	}
	return errors.Errorf("insufficient funds: %d < %d", w.coins, amount)
}

func PayForUnitsAndBuildins(w *PlayerWallet) {
	if err := w.spendCoins(3); err != nil {
		log.Println(err)
	}
}

func main() {
	var (
		//wg     sync.WaitGroup
		wallet = &PlayerWallet{coins: 5}
	)

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		PayForUnitsAndBuildins(wallet)
	}()

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		PayForUnitsAndBuildins(wallet)
	}()
	//wg.Wait()
	time.Sleep(100 * time.Millisecond)
}
