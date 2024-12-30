package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	source string
	amount int
}

func main() {
	var bankBalance int
	var mx sync.Mutex
	fmt.Printf("Initial account balance : %d.00", bankBalance)

	incomes := []Income{
		{source: "Main Job", amount: 3000},
		{source: "Part Time", amount: 2400},
		{source: "Investments", amount: 3200},
		{source: "Mentor", amount: 1200},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			mx.Lock()
			for week := 1; week <= 52; week++ {
				bankBalance += income.amount

				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, income.amount, income.source)
				fmt.Println()
			}

			mx.Unlock()
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final bank balance : %d", bankBalance)
	fmt.Println()
}
