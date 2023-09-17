package algorithm

import (
	"fmt"
	c "glcharge/container"
	"glcharge/models"
	"math"
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
func transform(inputMap map[int][]models.ChargePointStatus) int {
	outputMap := make(map[int][]models.ChargePointStatus)
	emptyKeys := make([]int, 0)
	sum := 0
	// Iterate over the inputMap
	for key := 1; key <= len(inputMap); key++ {
		value := inputMap[key]
		if len(value) == 0 {
			emptyKeys = append(emptyKeys, key)
		} else {
			if len(emptyKeys) != 0 {
				emptyKey := emptyKeys[0]
				emptyKeys = emptyKeys[1:]
				outputMap[emptyKey] = value
				emptyKeys = append(emptyKeys, key)
			} else {
				outputMap[key] = value
			}
		}
	}

	for key := 1; key <= len(outputMap); key++ {
		sum += len(outputMap[key]) * key
	}
	return sum
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
		if groupChargePointsMap[chargePoint.GroupId] != nil {
			groupChargePointsMap[chargePoint.GroupId] =
				append(groupChargePointsMap[chargePoint.GroupId], chargePoint)
		} else {
			groupChargePointsMap[chargePoint.GroupId] =
				append(make([]models.ChargePointStatus, 0), chargePoint)
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
		sumOfPriorities := transform(priorityChargePointsMap)
		fraction := float64(100) / float64(sumOfPriorities)

		fmt.Println("group: ", group.Id, " fraction: ", fraction, " sumOfPriroities: ", sumOfPriorities)
		// if numOfChargePoints > 1 {
		// 	sort.Slice(chargePoints, func(i, j int) bool {
		// 		return chargePoints[i].Priority < chargePoints[j].Priority
		// 	})
		// }

		// The output must be a list of maps with ChargePointId and Current fields. The order does not matter
		for _, chargePoint := range chargePointsGroup {
			percentageOfMaxCurrentForChargePoint := math.Round(float64(chargePoint.Priority)*fraction) / 100
			calculatedMaxCurrentForChargePoint := maxCurrent * percentageOfMaxCurrentForChargePoint
			resultMap[chargePoint.ChargePointId] = calculatedMaxCurrentForChargePoint
		}
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
