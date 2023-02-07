package main

import (
	"fmt"
)

func josephus(n int, k int, people []models.PlayerProfile) int {
	result := make([]int, 0)
	idx := 0
	round := 1
	for n > 1 {
		idx = (idx + k - 1) % n
		result = append(result, people[idx].id)
		fmt.Println("The person Eliminated in round ", round, " is: ", people[idx].Name)
		people = append(people[:idx], people[idx+1:]...)
		n--
		round++
	}
	result = append(result, people[0].id)
	return result[len(result)-1]
}

func main() {
	players := []PlayerProfile{
		{Name: "John", Age: 25, id: 1},
		{Name: "Jane", Age: 22, id: 2},
		{Name: "Jim", Age: 30, id: 3},
		{Name: "Jill", Age: 28, id: 4},
		{Name: "Jack", Age: 26, id: 5},
	}
	n, k := 5, 2
	result := josephus(n, k, players)
	fmt.Println("The Last person to survive is:", result)
}
