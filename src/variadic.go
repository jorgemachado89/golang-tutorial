package main

import (
	"fmt"
)

func main() {

	// bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 7, 5, 2)
	// fmt.Println(bestFinish)

	tailoredBestFinish := tailoredResultHoc(bestLeagueFinishes)
	fmt.Println(tailoredBestFinish)
}

type bestLeague func(...int) int

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}

func tailoredResultHoc(positionCalc bestLeague) string {
	leagueRanks := []int{13, 10, 13, 17, 14, 16, 7, 5, 2}

	return fmt.Sprint("Highest League Finish: ", positionCalc(leagueRanks...))
}