package solution

import (
	"fmt"
)

type edgeT struct {
	u rune
	w int
}

type graphT map[rune][]edgeT

type Solution struct {
	graph            graphT
	matrix           map[rune]map[rune]int
	shortestDistance map[rune]map[rune]int
	calculated       bool
}

func NewSolution() *Solution {
	return &Solution{
		graph:            make(graphT),
		matrix:           make(map[rune]map[rune]int),
		shortestDistance: make(map[rune]map[rune]int),
		calculated:       false,
	}
}

func (s *Solution) AddEdge(v rune, u rune, w int) {
	s.graph[v] = append(s.graph[v], edgeT{u, w})
	if s.matrix[v] == nil {
		s.matrix[v] = make(map[rune]int)
	}
	s.matrix[v][u] = w
}

func (s *Solution) debugMatrix() {
	for from, val := range s.matrix {
		for to, w := range val {
			fmt.Printf("%c %c %d\n", from, to, w)
		}
	}
}

func (s *Solution) RouteDistance(route string) (int, error) {
	v := rune(route[0])
	var result = 0
	for i, x := range route {
		if i == 0 {
			continue
		}
		if w, ok := s.matrix[v][x]; ok {
			result += w
			v = x
		} else {
			return 0, fmt.Errorf("NO SUCH ROUTE")
		}
	}
	return result, nil
}

func (s *Solution) calcShortestDistance() {
	vertices := make(map[rune]bool)
	for from, val := range s.matrix {
		vertices[from] = true
		if s.shortestDistance[from] == nil {
			s.shortestDistance[from] = make(map[rune]int)
		}
		for to, w := range val {
			vertices[to] = true
			s.shortestDistance[from][to] = w
		}
	}

	for k := range vertices {
		if s.shortestDistance[k] == nil {
			continue
		}
		for i := range vertices {
			if s.shortestDistance[i] == nil {
				continue
			}
			if _, ok := s.shortestDistance[i][k]; !ok {
				continue
			}
			for j := range vertices {
				w := s.shortestDistance[i][k]
				if val, ok := s.shortestDistance[k][j]; ok {
					w += val
				} else {
					continue
				}
				if val, ok := s.shortestDistance[i][j]; ok {
					if val < w {
						w = val
					}
				}

				s.shortestDistance[i][j] = w
			}
		}
	}
	s.calculated = true
}

func (s *Solution) ShortestDistance(v rune, u rune) (int, error) {
	if !s.calculated {
		s.calcShortestDistance()
	}
	if s.shortestDistance[v] == nil {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	} else if val, ok := s.shortestDistance[v][u]; !ok {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	} else {
		return val, nil
	}
}

func (s *Solution) routesLessDistanceNumber(v rune, to rune, dist int) (int, error) {
	if dist <= 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	result := 0
	found := false
	if v == to {
		result += 1
		found = true
	}
	for _, edge := range s.graph[v] {
		u, w := edge.u, edge.w
		if val, err := s.routesLessDistanceNumber(u, to, dist - w); err == nil {
			result += val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	return result, nil
}

func (s *Solution) RoutesLessDistanceNumber(v rune, to rune, dist int) (int, error) {
	if val, err := s.routesLessDistanceNumber(v, to, dist); err == nil {
		if v == to {
			val--
		}
		return val, nil
	} else {
		return 0, err
	}
}

func (s *Solution) RoutesLessStopsNumber(v rune, to rune, stops int) (int, error) {
	if val, err := s.routesLessStopsNumber(v, to, stops); err == nil {
		if v == to {
			val--
		}
		return val, nil
	} else {
		return 0, err
	}
}

func (s *Solution) routesLessStopsNumber(v rune, to rune, stops int) (int, error) {
	if stops < 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	result := 0
	found := false
	if v == to {
		result += 1
		found = true
	}
	for _, edge := range s.graph[v] {
		u := edge.u
		if val, err := s.routesLessStopsNumber(u, to, stops - 1); err == nil {
			result += val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	return result, nil
}

func (s *Solution) RoutesExactStopsNumber(v rune, to rune, stops int) (int, error) {
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
	for _, edge := range s.graph[v] {
		u := edge.u
		if val, err := s.RoutesExactStopsNumber(u, to, stops - 1); err == nil {
			result += val
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	}
	return result, nil
}