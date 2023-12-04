package main

import (
	"github.com/pkg/errors"
	"log"
	"time"
)

type PlayerWallet struct {
	coins int
}

func (w *PlayerWallet) spendCoins(amount int) error {
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
	wallet := &PlayerWallet{coins: 5}

	go func() {
		PayForUnitsAndBuildins(wallet)
	}()

	go func() {
		PayForUnitsAndBuildins(wallet)
	}()

	time.Sleep(100 * time.Millisecond)
}
