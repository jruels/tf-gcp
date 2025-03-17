package main

import "fmt"

func main() {
	teams := map[string]int{"basketball": 5,
		"football": 11, "baseball": 9}

	teams["hockey"] = 11

	delete(teams, "football")

	for team, players := range teams {
		fmt.Println(team, players)
	}

}
