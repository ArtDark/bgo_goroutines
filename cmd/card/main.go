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
			card.Transaction{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "4",
				Bill:   755_49,
				Time:   1592823454,
				MCC:    "5931",
				Status: "Valid",
			},
		},
	}
	//masterPointer := &master
	//
	//transaction := card.Transaction{
	//	Id:     "5",
	//	Bill:   233_43,
	//	Time:   1596773221,
	//	Status: "Valid",
	//	MCC: "5822",
	//}

	//fmt.Println("MasterCard: ", master)
	//card.AddTransaction(masterPointer, transaction)
	//fmt.Println("MasterCard: ", master)
	//
	//mcc := []string{"5411", "5812"}
	//
	//cashBack := card.SumByMCC(master.Transactions, mcc)
	//fmt.Println("Cashback sum:", cashBack)
	//
	//
	//category := card.TranslateMCC(master.Transactions[0].MCC)
	//fmt.Println(category)\

	fmt.Println("Before sort: ", master.Transactions)
	card.SortTransactions(master.Transactions)
	fmt.Println("After sort:  ", master.Transactions)

}
