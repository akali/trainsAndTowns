package solution

import (
	"fmt"
	"strconv"
)

type edgeT struct {
	U rune
	W int
}

type graphT map[rune][]edgeT

type Solution struct {
	Graph            graphT
	matrix           map[rune]map[rune]int
	shortestDistance map[rune]map[rune]int
	parent           map[rune]map[rune]rune
	calculated       bool
}

func NewSolution() *Solution {
	return &Solution{
		Graph:            make(graphT),
		matrix:           make(map[rune]map[rune]int),
		shortestDistance: make(map[rune]map[rune]int),
		parent:           make(map[rune]map[rune]rune),
		calculated:       false,
	}
}

func (s *Solution) AddEdge(v rune, u rune, w int) {
	s.Graph[v] = append(s.Graph[v], edgeT{u, w})
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

func (s *Solution) constructVertices() (vertices map[rune]bool) {
	vertices = make(map[rune]bool)
	for from, val := range s.matrix {
		vertices[from] = true
		if s.shortestDistance[from] == nil {
			s.shortestDistance[from] = make(map[rune]int)
		}
		if s.parent[from] == nil {
			s.parent[from] = make(map[rune]rune)
		}
		for to, w := range val {
			vertices[to] = true
			s.shortestDistance[from][to] = w
		}
	}
	return vertices
}

func (s *Solution) calcShortestDistance() {
	vertices := s.constructVertices()

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
					if w < val {
						s.parent[i][j] = k
						s.shortestDistance[i][j] = w
					}
				} else {
					s.shortestDistance[i][j] = w
					s.parent[i][j] = k
				}
			}
		}
	}
	s.calculated = true
}

type Path struct {
	Path     []rune
	Distance int
}

func PrintView(p interface{}) string {
	path := p.(*Path)
	result := ""
	result += strconv.Itoa(path.Distance) + " "
	for _, k := range path.Path {
		result += string(k)
	}
	return result
}

func (s *Solution) ShortestDistance(v rune, u rune) (*Path, error) {
	if !s.calculated {
		s.calcShortestDistance()
	}
	if s.shortestDistance[v] == nil {
		return nil, fmt.Errorf("NO SUCH ROUTE")
	} else if val, ok := s.shortestDistance[v][u]; !ok {
		return nil, fmt.Errorf("NO SUCH ROUTE")
	} else {
		return &Path{
			Path:     s.shortestDistancePath(v, u),
			Distance: val,
		}, nil
	}
}

func (s *Solution) shortestDistancePath(v rune, u rune) []rune {
	k := s.parent[v][u]

	if v == k {
		return []rune{v, u}
	}

	if u == k {
		return []rune{v, u}
	}

	start := s.shortestDistancePath(v, k)
	end := s.shortestDistancePath(k, u)

	var result []rune

	for i, v := range start {
		if i+1 < len(start) {
			result = append(result, v)
		}
	}
	for _, v := range end {
		result = append(result, v)
	}
	return result
}

func (s *Solution) ShortestDistancePath(v rune, u rune) ([]rune, error) {
	if v == u {
		return []rune{v}, nil
	}
	if !s.calculated {
		s.calcShortestDistance()
	}
	if s.shortestDistance[v] == nil {
		return []rune{}, fmt.Errorf("NO SUCH ROUTE")
	} else if _, ok := s.shortestDistance[v][u]; !ok {
		return []rune{}, fmt.Errorf("NO SUCH ROUTE")
	} else {
		return s.shortestDistancePath(v, u), nil
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
	for _, edge := range s.Graph[v] {
		u, w := edge.U, edge.W
		if val, err := s.routesLessDistanceNumber(u, to, dist-w); err == nil {
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
	var d = make([][]int, stops+1)
	vertices := s.constructVertices()

	maxElement := 'A'

	for k := range vertices {
		if maxElement < k {
			maxElement = k
		}
	}

	for i := 0; i <= stops; i++ {
		d[i] = make([]int, maxElement+1)
	}

	d[0][v] = 1

	for i := 1; i <= stops; i++ {
		for v := range vertices {
			for _, e := range s.Graph[v] {
				u := e.U
				d[i][u] += d[i-1][v]
			}
		}
	}

	result := 0

	for i := 1; i <= stops; i++ {
		result += d[i][to]
	}

	if result == 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	} else {
		return result, nil
	}
}

func (s *Solution) RoutesExactStopsNumber(v rune, to rune, stops int) (int, error) {
	var d = make([][]int, stops+1)
	vertices := s.constructVertices()

	maxElement := 'A'

	for k := range vertices {
		if maxElement < k {
			maxElement = k
		}
	}

	for i := 0; i <= stops; i++ {
		d[i] = make([]int, maxElement+1)
	}

	d[0][v] = 1

	for i := 1; i <= stops; i++ {
		for v := range vertices {
			for _, e := range s.Graph[v] {
				u := e.U
				d[i][u] += d[i-1][v]
			}
		}
	}

	if d[stops][to] == 0 {
		return 0, fmt.Errorf("NO SUCH ROUTE")
	} else {
		return d[stops][to], nil
	}
}
