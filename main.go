package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

		// Convert combination to a string key for uniqueness checking
		key := strings.Join(strings.Fields(fmt.Sprint(currentCombination)), ",")
		if _, exists := sc.results[key]; !exists {
			// If the combination is unique, store it
			sc.results[key] = append([]int{}, currentCombination...) // Deep copy
		}
		return
	}

	// Try each stamp price and recursively call the function
	for _, price := range sc.stampPrices {
		// Allow a stamp to be used repeatedly, so no restrictions here
		sc.FindCombinations(target, append(currentCombination, price), total+price)
	}
}

// PrintCombinations prints all combinations that match the target amount
func (sc *StampCalculator) PrintCombinations(target int) {
	// First, include the single-stamp solutions explicitly
	for _, price := range sc.stampPrices {
		// Check if using only this stamp repeatedly reaches the target
		if target%price == 0 {
			count := target / price
			currentCombination := make([]int, count)
			for i := 0; i < count; i++ {
				currentCombination[i] = price
			}

			// Sort and add as a valid result
			sort.Ints(currentCombination)
			key := strings.Join(strings.Fields(fmt.Sprint(currentCombination)), ",")
			if _, exists := sc.results[key]; !exists {
				sc.results[key] = append([]int{}, currentCombination...)
			}
		}
	}

	// If no results, print no valid combinations
	if len(sc.results) == 0 {
		fmt.Println("No valid combinations found.")
		return
	}

	round := 0

	// Iterate through the unique results
	for _, combination := range sc.results {
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
