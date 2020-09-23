package main

import (
	"fmt"
	"github.com/ArtDark/bgo_goroutines/pkg/card"
)

func main() {
	master := card.Card{
		Id: "0001",
		Owner: card.Owner{
			FirstName: "Artem",
			LastName:  "Balusov",
		},
		Issuer:   "Visa",
		Balance:  43534_34,
		Currency: "RUR",
		Number:   "5106 2158 3920 4837",
		Icon:     "http://...",
		Transactions: []card.Transaction{
			{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
				Status: "Valid",
			},
			{
				Id:     "4",
				Bill:   755_49,
				Time:   1592823454,
				MCC:    "5931",
				Status: "Valid",
			},
		},
	}

	fmt.Println("Before sort: ", master.Transactions)
	card.SortTransactions(master.Transactions)
	fmt.Println("After sort:  ", master.Transactions)

}
