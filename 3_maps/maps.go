package main

import (
	"fmt"
)

func work(n int) int {
	return n * n
}

func Caching() {
	var n int
	var cache = map[int]int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if _, inMap := cache[n]; !inMap {
			cache[n] = work(n)
		}
		fmt.Printf("%d ", cache[n])
	}
}

func correctStat() {
	groupCity := map[int][]string{
		10: {"aaa", "bbb", "ccc", "ddd"},
		100: {"Tver", "Tula", "Omsk", "Pskov"},
		1000: {"Moscow", "Piter", "Eburg", "Novosib"},
	}
	cityPopulation := map[string]int{
		"Tver": 400,
		"Tula": 500,
		"aaa": 45,
		"Piter": 3750,
	}

	isIn := func (sl []string, s string) bool {
		for _, el := range sl {
			if el == s {return true}
		}
		return false
	}

	var length = len(groupCity[10]) + len(groupCity[1000])
	var wrongCities = make([]string, length)
	fmt.Println(length)
	for key := range cityPopulation {
		// if value < 100 || value > 999 {
		// 	wrongCities = append(wrongCities, key)
		// }
		if isIn(groupCity[10], key) || isIn(groupCity[1000], key) {
			wrongCities = append(wrongCities, key)
		}
	}
	for _, el := range wrongCities {
		delete(cityPopulation, el)
	}
	fmt.Println(cityPopulation)
}

func attentionPlease() {
	// learning to read bad code :))
	var friends0fDima []string
	friends0fSemyon := make([]string, 3)
	friends0fDima = append(friends0fDima, "Костя", "Семён", "Таня")
	friends0fSemyon = append(friends0fSemyon, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   friends0fDima,
		"Semyon": friends0fSemyon,
		"Костя":  nil,
	}
	_, value := friends["Костя"]
	fmt.Print(value, " ")
	delete(friends, "Dima")
	fmt.Print(friends0fSemyon[3])
}

func main() {
	// Caching()
	// correctStat()
	attentionPlease()
}
