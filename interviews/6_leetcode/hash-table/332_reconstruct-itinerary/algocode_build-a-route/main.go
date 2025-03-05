package main

import "fmt"

func route(tickets [][]string) []string {
	dist := make(map[string]struct{})
	mapping := make(map[string]string)
	res := make([]string, 0)
	for _, ticket := range tickets {
		dist[ticket[1]] = struct{}{}
		mapping[ticket[0]] = ticket[1]
	}
	for _, ticket := range tickets { // ищем первый город
		if _, found := dist[ticket[0]]; !found {
			res = append(res, ticket[0])
			break
		}
	}
	var src string
	if len(res) > 0 {
		src = res[0]
	}
	for {
		if dist, found := mapping[src]; found {
			res = append(res, dist)
			src = dist
		} else {
			break
		}
	}
	return res
}

func main() {
	tickets := [][]string{{"B", "C"},{"A", "B"},{"C", "D"}}
	fmt.Println(route(tickets))
}
