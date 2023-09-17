package main

import (
	"fmt"
	"glcharge/models"
	"sort"
)

func transform(inputMap map[int][]models.ChargePointStatus) {
	// Create a slice to store the keys
	keys := make([]int, 0, len(inputMap))
	// Iterate over the map and extract the keys
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	sumOfPriorities := 0

	for i := 1; i <= len(keys); i++ {
		priority := keys[i-1]
		numberOfChargePointsWithPriority := len(inputMap[priority])

		sumOfPriorities += i * numberOfChargePointsWithPriority
	}

	fmt.Println(sumOfPriorities)
}

func main() {
	inputMap := map[int][]models.ChargePointStatus{
		3: {{}, {}, {}, {}},
		4: {{}, {}, {}},
		5: {{}, {}},
	}

	transform(inputMap)
	/*
		Group ima MaxCurrent 100 in 3 polnilnice:
			polnilnica A  - prioriteta 1,
			polnilnica B - prioriteta 3,
			polnilnica C - prioriteta 2

		- trenutno polni samo polnilnica A, potem dobi 100% od toka, ki je na voljo.
		- Če polnita polnilnica A in B, potem dobi polnilnica A 33% toka, polnilnica B pa 66% toka.
		- Če polnijo polnilnica A, B in C, potem bi polnilnica A morala dobiti 16% toka, polnilnica B 50% toka in polnilnica C 33% toka.

		example 1:
		chargePoint A, priority 1	-> 100% toka

		example2:
		chargePoint A, priority 1	-> 33% toka
		chargePoint B, priority 3	-> 66% toka

		example3:
		chargePoint A, priority 1	-> 16% toka
		chargePoint B, priority 3	-> 50% toka
		chargePoint C, priority 2	-> 33% toka

	*/

	/*
		dobim vse unikatne prioritete
		jih sortiram naraščajoče
		za vsako od prioritet dobim število polnilnih postaj s to prioriteto
		najdem index(+1) prioritete v seznamu unikatnih prioritet
		število polnilnih postaj množim z indeksom
		100% delimo z vsoto indexov(+1) prioritet * število postaj s to prioriteto -> sum of priroity units
		sumOfPriorityUnits potem množimo z indexom(+1) prioritete

		primer:
		polnilne postaje: A (1), B(3), C(1)
		unikatne prioritete: 1, 3
		sortirano: 1,3

		index prioritete 1 = 0 + 1
		število polnilnih postaj s prioriteto = 2
		index prioritete 3 = 1 + 1
		število polnilnih postaj s prioriteto = 1

		sumOfPriorityUnits = 1 * 2 + 2 * 1 = 4
		100 / 4 = 25

		A: 25 * 1
		B: 25 * 2
		C: 25 * 1


	*/

	/*
		inputMap := map[int][]int{
			4: {A, B, C, D},
			3: {E, F, G},
			5: {H, I},
		}

		unikatne prioritete (ključi mape): [4,3,5]
		sortirano: 	[3,4,5]
		indexi:		 1,2,3

		idx p1: 0 + 1 = 1
		idx p2: 1 + 1 = 2
		idx p3: 2 + 1 = 3

		sumOfPriorityUnits = 1 * 4 + 2 * 2 + 3 * 2 = 16
		100 / 17 = 5,88

		A: 6.25 * 1 = 6.25%
		B: 6.25 * 1 = 6.25%
		C: 6.25 * 1 = 6.25%
		D: 6.25 * 1 = 6.25%
		E: 6.25 * 2 = 12.5%
		F: 6.25 * 2 = 12.5%
		G: 6.25 * 2 = 12.5%
		H: 6.25 * 3 = 18.75%
		I: 6.25 * 3 = 18.75%

	*/

}
