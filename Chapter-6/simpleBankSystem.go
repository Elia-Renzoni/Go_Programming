/**
*	@author Elia Renzoni
*	@date 26/11/2023
*	@brief HashTable using go. Bank Account and transactions
**/

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var (
		overallBankAccount = make(map[int]float64)
		balance            float64
		err                error
		getBalance         func() (float64, error) = setBalance
	)

	if balance, err = getBalance(); err != nil || balance < 0 {
		os.Exit(1)
	}

	for i := 0; i < 20; i++ {
		addBankAccount(overallBankAccount, balance-float64(i))
	}
	transfertMoney(overallBankAccount, 12, 55, 6000)
	depositMoney(overallBankAccount, 4, 7000.90)
	withdraw(overallBankAccount, 3, 900)
}

func addBankAccount(accounts map[int]float64, balance float64) {
	accounts[rand.Intn(200)] = balance
}

func transfertMoney(accounts map[int]float64, senderAccountId, receiverAccountId int, value float64) {
	_, okSender := accounts[senderAccountId]
	_, okReceiver := accounts[receiverAccountId]
	if okSender && okReceiver {
		accounts[receiverAccountId] = value
		fmt.Println("Operazione eseguita")
	} else {
		fmt.Printf("Impossibile eseguire l'operazione")
	}
}

func depositMoney(accounts map[int]float64, accountId int, newBalance float64) {
	var (
		commaOkIdiom bool
		currentValue float64
	)
	if currentValue, commaOkIdiom = accounts[accountId]; commaOkIdiom {
		accounts[accountId] = newBalance + currentValue
		fmt.Printf("Aggiunti nuovi soldi !")
	} else {
		fmt.Printf("Impossibile aggiungere nuovi soldi !")
	}
}

func withdraw(accounts map[int]float64, accountId int, balanceToWithdraw float64) {
	value, ok := accounts[accountId]
	if ok {
		accounts[accountId] = value - balanceToWithdraw
		fmt.Printf("Soldi scalati con successo !")
	} else {
		fmt.Printf("Impossibile effetturare il prelievo")
	}
}

func setBalance() (float64, error) {
	var (
		balance float64
		control bool = true
		err     error
	)
	for control {
		fmt.Printf("Inserisci il saldo del nuovo conto >")
		_, err = fmt.Scanf("%f", &balance)
		if err != nil {
			return -1.0, err
		} else {
			control = false
		}
	}
	return balance, nil
}
