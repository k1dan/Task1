package main

import (
	"fmt"
	"strconv"
	"strings"
	"task1/sorting"
)

func main() {
	fmt.Println("Insert numbers to sort. Numbers should be separated by commas, without spaces:")

	var numbers []int
	var in string

	fmt.Scanln(&in)
	input := strings.Split(in, ",")

	for _, v := range input {
		s, _ := strconv.Atoi(v)
		numbers = append(numbers, s)
	}

	fmt.Println("Select type of sorting from list, by typing its name:")
	fmt.Println("1. BubbleSort")
	fmt.Println("2. CocktailSort")
	fmt.Println("3. InsertionSort")
	fmt.Println("4. SelectionSort")
	fmt.Println("5. MergeSort")
	fmt.Println("6. QuickSort")

	var sortingWay string
	fmt.Scanln(&sortingWay)

	switch sortingWay {
	case "BubbleSort":
		sorting.BubbleSort(&numbers)
	case "CocktailSort":
		sorting.CocktailSort(&numbers)
	case "InsertionSort":
		sorting.InsertionSort(&numbers)
	case "SelectionSort":
		sorting.SelectionSort(&numbers)
	case "MergeSort":
		sorting.MergeSort(&numbers)
	case "QuickSort":
		sorting.QuickSort(&numbers)
	}
	fmt.Println(numbers)
}
