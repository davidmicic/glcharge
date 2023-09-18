package handlers

/*
The user must be able to change each charge point's
priority, each connector's status or MaxCurrent for
the group.
*/

import (
	"encoding/json"
	"fmt"
	"glcharge/go/src/algorithm"
	c "glcharge/go/src/container"
	"net/http"
	"strconv"
)

func Makehttphandlers() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/changeChargePointPriority", changeChargePointPriority)
	serveMux.HandleFunc("/changeConnectorStatus", changeConnectorStatus)
	serveMux.HandleFunc("/changeMaxCurrentGroup", changeMaxCurrentGroup)
	return serveMux
}

func changeChargePointPriority(w http.ResponseWriter, r *http.Request) {
	chargePointId, err := strconv.Atoi(r.URL.Query().Get("chargePointId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "ChargePointId is not a number")
		return
	}

	priority, err := strconv.Atoi(r.URL.Query().Get("priority"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Priority is not a number")
		return
	}

	container := c.GetContainer()
	storage := container.Storage()

	storage.ChangeChargePointPriorityById(chargePointId, priority)
	resultMap := algorithm.Algorithm()

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(resultMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func changeConnectorStatus(w http.ResponseWriter, r *http.Request) {
	connectorId, err := strconv.Atoi(r.URL.Query().Get("connectorId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "ConnectorId is not a number")
		return
	}

	status := r.URL.Query().Get("status")
	if status == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Missing status value")
		return
	}

	if status != "Available" && status != "Charging" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Wrong status value")
		return
	}

	container := c.GetContainer()
	storage := container.Storage()

	storage.ChangeConnectorStatusById(connectorId, status)
}

func changeMaxCurrentGroup(w http.ResponseWriter, r *http.Request) {
	groupId, err := strconv.Atoi(r.URL.Query().Get("groupId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "GroupId is not a number")
		return
	}

	MaxCurrent, err := strconv.ParseFloat(r.URL.Query().Get("MaxCurrent"), 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "MaxCurrent is not a number")
		return
	}

	container := c.GetContainer()
	storage := container.Storage()

	storage.ChangeGroupMaxCurrentById(groupId, MaxCurrent)
}

// var ChargePointConnectorsGrp1 = []models.ChargePointConnector{
// 	models.ChargePointConnector{
// 		Id:     1,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     2,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     3,
// 		Status: "Available",
// 	},
// }
// var ChargePointConnectorsGrp2 = []models.ChargePointConnector{
// 	models.ChargePointConnector{
// 		Id:     4,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     5,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     6,
// 		Status: "Available",
// 	},
// }
// var ChargePointConnectorsGrp3 = []models.ChargePointConnector{
// 	models.ChargePointConnector{
// 		Id:     7,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     8,
// 		Status: "Available",
// 	},
// 	models.ChargePointConnector{
// 		Id:     9,
// 		Status: "Available",
// 	},
// }

// var ChargePointStatus1 = models.ChargePointStatus{
// 	ChargePointId:         "1",
// 	Priority:              1,
// 	ChargePointConnectors: ChargePointConnectorsGrp1,
// }
// var ChargePointStatus2 = models.ChargePointStatus{
// 	ChargePointId:         "2",
// 	Priority:              2,
// 	ChargePointConnectors: ChargePointConnectorsGrp2,
// }
// var ChargePointStatus3 = models.ChargePointStatus{
// 	ChargePointId:         "3",
// 	Priority:              3,
// 	ChargePointConnectors: ChargePointConnectorsGrp3,
// }

// var group1 = models.Group{
// 	Id:           "1",
// 	MaxCurrent:   100.0,
// 	ChargePoints: []models.ChargePointStatus{ChargePointStatus1, ChargePointStatus2},
// }
// var group2 = models.Group{
// 	Id:           "2",
// 	MaxCurrent:   50.0,
// 	ChargePoints: []models.ChargePointStatus{ChargePointStatus3},
// }
