package ds

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntervalTree_InsertAndSearch(t *testing.T) {
	tree := &IntervalTree{}

	// Insert intervals into the tree
	intervals := []Interval{
		{30, 40},
		{5, 20},
		{12, 15},
		{10, 30},
		{17, 19},
		{15, 20},
	}
	for _, interval := range intervals {
		tree.Insert(interval)
	}

	// Search overlapping intervals
	searchIntervals := []Interval{
		{6, 7},
		{14, 16},
		{21, 23},
	}
	expectedResults := [][]Interval{
		{{5, 20}},
		{{5, 20}, {12, 15}, {10, 30}, {15, 20}},
		{{10, 30}},
	}

	for i, interval := range searchIntervals {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			results := tree.Search(interval)
			if !reflect.DeepEqual(results, expectedResults[i]) {
				t.Errorf("Search(%v) = %v, want %v", interval, results, expectedResults[i])
			}
		})
	}
}
