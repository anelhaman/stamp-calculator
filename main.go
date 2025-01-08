package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// StampCalculator struct to encapsulate the logic
type StampCalculator struct {
	stampPrices []int
	results     map[string][]int
}

// NewStampCalculator creates a new StampCalculator instance
func NewStampCalculator(stampPrices []int) *StampCalculator {
	return &StampCalculator{
		stampPrices: stampPrices,
		results:     make(map[string][]int),
	}
}

// FindCombinations finds all combinations of stamps that match the target amount
func (sc *StampCalculator) FindCombinations(target int, currentCombination []int, total int) {
	// If the total sum exceeds the target, return
	if total > target {
		return
	}

	// If we match the target, save the current combination
	if total == target {
		// Sort the combination to ensure uniqueness (order doesn't matter)
		sort.Ints(currentCombination)

		key := fmt.Sprint(currentCombination)
		// Store the combination as a key in the map to ensure uniqueness
		if _, exists := sc.results[key]; !exists {
			sc.results[key] = append([]int{}, currentCombination...) // Deep copy
		}
		return
	}

	// Try each stamp price and recursively call the function
	for _, price := range sc.stampPrices {
		sc.FindCombinations(target, append(currentCombination, price), total+price)
	}
}

// PrintCombinations prints all combinations that match the target amount
func (sc *StampCalculator) PrintCombinations(target int) {
	if len(sc.results) == 0 {
		fmt.Println("No valid combinations found.")
		return
	}

	round := 0

	// Iterate through the unique results
	for _, combination := range sc.results {
		// Count each stamp's occurrence
		stampsCount := make(map[int]int)
		total := 0
		for _, price := range combination {
			stampsCount[price]++
			total += price
		}

		// Print the combination only if the total is equal to the input amount
		if total == target {
			// Increment total combinations count
			round++

			fmt.Printf("Combination #%d\n", round)
			fmt.Println("---------------------------------------------")
			for _, price := range sc.stampPrices {
				if count, found := stampsCount[price]; found {
					fmt.Printf("Stamp %d THB:\t%d\t=\t%d\tTHB\n", price, count, price*count)
				}
			}
			fmt.Println("---------------------------------------------")
			fmt.Printf("Total:\t\t\t\t%d\tTHB\n\n", total)
		}
	}

	fmt.Printf("Total combinations: %d\n", round)
}

func main() {
	// Read the amount from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide the amount.")
		return
	}
	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid amount.")
		return
	}

	// Available stamp prices
	stampPrices := []int{9, 5, 3}

	// Create a new StampCalculator instance
	calculator := NewStampCalculator(stampPrices)

	// Find combinations
	calculator.FindCombinations(amount, []int{}, 0)

	// Print the results
	calculator.PrintCombinations(amount)
}
