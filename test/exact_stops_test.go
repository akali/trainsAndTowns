package test

import (
	"fmt"
	"testing"
	"transAndTowns/solution"
)

func exactStopsNaive(s *solution.Solution, v rune, to rune, stops int) (int, error) {
	if stops < 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	result := 0
	found := false
	if v == to {
		if stops == 0 {
			result += 1
		}
		found = true
	}
	for _, edge := range s.Graph[v] {
		u := edge.U
		if val, err := exactStopsNaive(s, u, to, stops-1); err == nil {
			result += val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	return result, nil
}

func routesLessStopsNumberNaive(s *solution.Solution, v, to rune, stops int) (int, error) {
	if stops < 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	result := 0
	found := false
	if v == to {
		result += 1
		found = true
	}
	for _, edge := range s.Graph[v] {
		u := edge.U
		if val, err := routesLessStopsNumberNaive(s, u, to, stops-1); err == nil {
			result += val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	return result, nil
}

func setUp() *solution.Solution {
	s := solution.NewSolution()

	s.AddEdge('A', 'B', 5)
	s.AddEdge('B', 'C', 4)
	s.AddEdge('C', 'D', 8)
	s.AddEdge('D', 'C', 8)
	s.AddEdge('D', 'E', 6)
	s.AddEdge('A', 'D', 5)
	s.AddEdge('C', 'E', 2)
	s.AddEdge('E', 'B', 3)
	s.AddEdge('A', 'E', 7)

	return s
}

func TestExactStops(t *testing.T) {
	s := setUp()

	result, err := s.RoutesExactStopsNumber('A', 'C', 4)
	actual_result, actual_err := exactStopsNaive(s, 'A', 'C', 4)
	if err != actual_err {
		t.Errorf("Expected %v, got %v\n", actual_err, err)
	}
	if result != actual_result {
		t.Errorf("Expected %d, got %d\n", actual_result, result)
	}
}

func TestRoutesLessStopsNumber(t *testing.T) {
	s := setUp()

	result, err := s.RoutesLessStopsNumber('C', 'C', 3)
	actual_result, actual_err := routesLessStopsNumberNaive(s, 'C', 'C', 3)
	if err != actual_err {
		t.Errorf("Expected %v, got %v\n", actual_err, err)
	}
	actual_result--
	if result != actual_result {
		t.Errorf("Expected %d, got %d\n", actual_result, result)
	}
}
