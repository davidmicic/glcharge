package models

type Group struct {
	Id         int
	MaxCurrent float64
}

type ChargePointStatus struct {
	ChargePointId int
	Priority      int
	GroupId       int
}

type ChargePointConnector struct {
	Id            int
	Status        string
	ChargePointId int
}
