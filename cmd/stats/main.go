package main

import (
	"fmt"
	"github.com/ArtDark/bgo_goroutines/pkg/stats"
	"time"
)

func main() {
	////runtime.GOMAXPROCS(4)
	//
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() {
	//	if err := f.Close(); err != nil {
	//		log.Print(err)
	//	}
	//}()
	//err = trace.Start(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer trace.Stop()

	const users = 10_000_000
	const transactionsPerUser = 100
	const transactionAmount = 1_00
	transactions := make([]int64, users*transactionsPerUser)

	for index := range transactions {
		// для простоты храним только суммы
		// и считаем, что каждая транзакция = 1 руб.
		transactions[index] = transactionAmount
	}

	total := int64(0)

	start := time.Date(2016, 1, 2, 0, 0, 0, 0, time.Local)
	finish := time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)

	months := make([]int64, 0)

	next := start
	for next.Before(finish) {
		months = append(months, next.Unix())
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, finish.Unix())
	fmt.Println(len(months))

	stats.SumConcurrently(transactions, len(months))

	time.Sleep(time.Minute)
	fmt.Println("Total:", total)
}
