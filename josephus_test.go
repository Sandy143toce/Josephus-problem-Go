package main

import (
	"testing"

	"github.com/Sandy143toce/Josephus-problem-Go/models"
)

func TestJosephus(t *testing.T) {
	tests := []struct {
		n, k   int
		people []models.PlayerProfile
		want   string
	}{
		{5, 2, []models.PlayerProfile{
			{Name: "Sandip", Age: 25, Id: 1},
			{Name: "Ram", Age: 22, Id: 2},
			{Name: "Sam", Age: 30, Id: 3},
			{Name: "Jack", Age: 28, Id: 4},
			{Name: "Elon", Age: 26, Id: 5},
		}, "Sam"},
		{6, 3, []models.PlayerProfile{
			{Name: "Sandip", Age: 25, Id: 1},
			{Name: "Ram", Age: 22, Id: 2},
			{Name: "Sam", Age: 30, Id: 3},
			{Name: "Jack", Age: 28, Id: 4},
			{Name: "Elon", Age: 26, Id: 5},
			{Name: "Musk", Age: 26, Id: 6},
		}, "Sandip"},
		{4, 2, []models.PlayerProfile{
			{Name: "Sandip", Age: 25, Id: 1},
			{Name: "Ram", Age: 22, Id: 2},
			{Name: "Sam", Age: 30, Id: 3},
			{Name: "Jack", Age: 28, Id: 4},
		}, "Sandip"},
	}
	for _, test := range tests {
		got := JosephusEliminator(test.n, test.k, test.people)
		if got != test.want {
			t.Errorf("josephus(%d, %d, %v) = %s; want %s", test.n, test.k, test.people, got, test.want)
		}
	}
}
