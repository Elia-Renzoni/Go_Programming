//
//	@author Elia Renzoni
//	@date 05/08/2023
//	@bief house robber
//

package main

import "fmt"

const (
	noSpoils  = -1
	increment = 2
)

func main() {
	var (
		moneyHouse []int16 = []int16{12, 3, 4, 5, 7}
		spoils     int16
	)
	if rob(moneyHouse, &spoils); spoils != noSpoils {
		fmt.Printf("No Allarm, Spoils : %d", spoils)
	} else {
		fmt.Printf("No Spoils, Allarm !")
	}
}

func rob(moneyHouse []int16, spoils *int16) {
	i := 0
	for i <= len(moneyHouse)-2 {
		*spoils += moneyHouse[i]
		i += 2
	}
}
