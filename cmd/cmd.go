package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"transAndTowns/solution"
)

func Run(reader *bufio.Reader) error {
	scanner := bufio.NewScanner(reader)

	s := solution.NewSolution()

	for scanner.Scan() {
		line := scanner.Text()

		for _, token := range strings.Fields(line) {
			v, u := rune(token[0]), rune(token[1])
			var k = 1
			if token[len(token)-1] != ',' {
				k = 0
			}
			w, err := strconv.Atoi(token[2 : len(token)-k])
			if err != nil {
				return fmt.Errorf("can't parse weight in %s", token)
			}
			s.AddEdge(v, u, w)
		}
	}

	var cnt = 1

	var result interface{}
	for _, route := range []string{"ABC", "AD", "ADC", "AEBCD", "AED"} {
		if distance, err := s.RouteDistance(route); err != nil {
			result = err
		} else {
			result = distance
		}
		fmt.Printf("Output #%d: %v\n", cnt, result)
		cnt++
	}

	if number, err := s.RoutesLessStopsNumber('C', 'C', 3); err != nil {
		result = err
	} else {
		result = number
	}
	fmt.Printf("Output #%d: %v\n", cnt, result)
	cnt++

	if number, err := s.RoutesExactStopsNumber('A', 'C', 4); err != nil {
		result = err
	} else {
		result = number
	}
	fmt.Printf("Output #%d: %v\n", cnt, result)
	cnt++

	for _, route := range []string{"AC", "BB"} {
		if distance, err := s.ShortestDistance(rune(route[0]), rune(route[1])); err != nil {
			result = err
		} else {
			result = distance
		}
		fmt.Printf("Output #%d: %v\n", cnt, result)
		cnt++
	}

	if number, err := s.RoutesLessDistanceNumber('C', 'C', 30); err != nil {
		result = err
	} else {
		result = number
	}
	fmt.Printf("Output #%d: %v\n", cnt, result)
	cnt++

	return nil
}
