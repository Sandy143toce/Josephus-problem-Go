package main

import (
	"testing"

	"github.com/Sandy143toce/Josephus-problem-Go/models"
)

func TestJosephus(t *testing.T) {
	players := []models.PlayerProfile{
		{Name: "Sandip", Age: 25, Id: 1},
		{Name: "Ram", Age: 22, Id: 2},
		{Name: "Sam", Age: 30, Id: 3},
		{Name: "Jack", Age: 28, Id: 4},
		{Name: "Elon", Age: 26, Id: 5},
	}
	players1 := []models.PlayerProfile{
		{Name: "Sandip", Age: 25, Id: 1},
		{Name: "Ram", Age: 22, Id: 2},
		{Name: "Sam", Age: 30, Id: 3},
		{Name: "Jack", Age: 28, Id: 4},
		{Name: "Elon", Age: 26, Id: 5},
		{Name: "Musk", Age: 26, Id: 6},
	}

	players2 := []models.PlayerProfile{
		{Name: "Sandip", Age: 25, Id: 1},
		{Name: "Ram", Age: 22, Id: 2},
		{Name: "Sam", Age: 30, Id: 3},
		{Name: "Jack", Age: 28, Id: 4},
	}

	tests := []struct {
		n, k   int
		people []models.PlayerProfile
		want   int
	}{
		{5, 2, players, 3},
		{6, 3, players1, 1},
		{4, 2, players2, 1},
	}

	for _, test := range tests {
		if got := JosephusEliminator(test.n, test.k, test.people); got != test.want {
			t.Errorf("josephus(%d, %d, %v) = %d; want %d", test.n, test.k, test.people, got, test.want)
		}
	}
}
