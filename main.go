package main

import (
	"fmt"

	"github.com/Sandy143toce/Josephus-problem-Go/models"
)

func JosephusEliminator(n int, k int, people []models.PlayerProfile) int {
	result := make([]int, 0)
	eliminator := 0
	round := 1
	for n > 1 {
		eliminator = (eliminator + k - 1) % n
		result = append(result, people[eliminator].Id)
		fmt.Println("The person Eliminated in round", round, "is:", people[eliminator].Name)
		people = append(people[:eliminator], people[eliminator+1:]...)
		n--
		round++
	}
	result = append(result, people[0].Id)
	return result[len(result)-1]
}

func main() {
	players := []models.PlayerProfile{
		{Name: "Sandip", Age: 25, Id: 1},
		{Name: "Ram", Age: 22, Id: 2},
		{Name: "Sam", Age: 30, Id: 3},
		{Name: "Jack", Age: 28, Id: 4},
		{Name: "Elon", Age: 26, Id: 5},
	}
	n, k := len(players), 2
	result := JosephusEliminator(n, k, players)
	fmt.Println("The ID of the person to survive is:", result)
}
