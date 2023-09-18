package algorithm

import (
	c "glcharge/go/src/container"
	"glcharge/go/src/models"
	"math"
	"sort"
)

/*
-	The algorithm should be able to distribute the maximum available current (MaxCurrent)
	to charge points in a group based on the specified priority.
-	If a charge point has no connector with "Charging" status,
	0 currernt is allocated to the charge point.
-	Otherwise, the current must be distributed based on the priority field of the charge point.
*/

/*
	VPRAŠANJA

Kaj predstavlja prioriteta pri distribuciji toka na polnilni postaji?
*/
func transform(inputMap map[int][]models.ChargePointStatus, maxCurrent float64, resultMap *map[int]float64) {
	keys := make([]int, 0, len(inputMap))
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

	minimalValuePercent := (float64(100) / float64(sumOfPriorities)) / 100

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		chargePointsWithPriority := inputMap[key]

		for j := 0; j < len(chargePointsWithPriority); j++ {
			chargePoint := chargePointsWithPriority[j]
			maxCurrentForChargePoint := maxCurrent * (minimalValuePercent * float64(i+1))
			(*resultMap)[chargePoint.ChargePointId] = math.Round(maxCurrentForChargePoint*100) / 100
		}
	}
}

func Algorithm() map[int]float64 {
	container := c.GetContainer()
	storage := container.Storage()

	groups, _ := storage.GetGroups()
	chargePoints, _ := storage.GetChargePointStatus()
	chargePointConnectors, _ := storage.GetChargePointConnector()

	// groupMap := make(map[string]models.Group)
	groupChargePointsMap := make(map[int][]models.ChargePointStatus)
	chargePointConnectorsMap := make(map[int][]models.ChargePointConnector)

	// loop through groups

	for _, connector := range chargePointConnectors {
		if connector.Status == "Charging" {
			if chargePointConnectorsMap[connector.ChargePointId] != nil {
				chargePointConnectorsMap[connector.ChargePointId] =
					append(chargePointConnectorsMap[connector.ChargePointId], connector)
			} else {
				chargePointConnectorsMap[connector.ChargePointId] =
					append(make([]models.ChargePointConnector, 0), connector)
			}
		}
	}

	for _, chargePoint := range chargePoints {
		if len(chargePointConnectorsMap[chargePoint.ChargePointId]) > 0 {
			if groupChargePointsMap[chargePoint.GroupId] != nil {
				groupChargePointsMap[chargePoint.GroupId] =
					append(groupChargePointsMap[chargePoint.GroupId], chargePoint)
			} else {
				groupChargePointsMap[chargePoint.GroupId] =
					append(make([]models.ChargePointStatus, 0), chargePoint)
			}
		}
	}

	// for each group get the number of ChargePoints with status "Charging"

	// var numOfChargePoints int = 0

	// sum the priorites of all charging points with status charging
	// Σ (from i = 1 to n) i

	resultMap := make(map[int]float64)
	for _, group := range groups {
		maxCurrent := group.MaxCurrent
		chargePointsGroup := groupChargePointsMap[group.Id]
		priorityChargePointsMap := make(map[int][]models.ChargePointStatus, 0)
		for _, chargePoint := range chargePointsGroup {
			if priorityChargePointsMap[chargePoint.Priority] != nil {
				priorityChargePointsMap[chargePoint.Priority] =
					append(priorityChargePointsMap[chargePoint.Priority], chargePoint)
			} else {
				priorityChargePointsMap[chargePoint.Priority] =
					append(make([]models.ChargePointStatus, 0), chargePoint)
			}
		}

		// numOfChargePoints := len(chargePoints)
		transform(priorityChargePointsMap, maxCurrent, &resultMap)
	}

	return resultMap
}

/*

1	->	[1]
2	->	[2]				1	->	[1]
3	->	[]		-->		2	->	[2]
4	->	[2]				3	->	[2]
5	->	[]

*/
