package algorithm

import (
	"fmt"
	c "glcharge/go/src/container"
	"glcharge/go/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// ChangeGroupMaxCurrentByIdDB     func(id int, maxCurrent float64)
	// ChangeConnectorStatusByIdDB     func(id int, status string)
	// ChangeChargePointPriorityByIdDB func(id int, priority int)
	GetChargePointConnectorDB func() ([]models.ChargePointConnector, error)
	GetChargePointStatusDB    func() ([]models.ChargePointStatus, error)
	GetGroupsDB               func() ([]models.Group, error)
)

type DalMock struct {
}

// ChangeChargePointPriorityById implements IDal.
func (*DalMock) ChangeChargePointPriorityById(id int, priority int) {
	panic("unimplemented")
}

// ChangeConnectorStatusById implements IDal.
func (*DalMock) ChangeConnectorStatusById(id int, status string) {
	panic("unimplemented")
}

// ChangeGroupMaxCurrentById implements IDal.
func (*DalMock) ChangeGroupMaxCurrentById(id int, maxCurrent float64) {
	panic("unimplemented")
}

// GetChargePointConnector implements IDal.
func (*DalMock) GetChargePointConnector() ([]models.ChargePointConnector, error) {
	return GetChargePointConnectorDB()
}

// GetChargePointStatus implements IDal.
func (*DalMock) GetChargePointStatus() ([]models.ChargePointStatus, error) {
	return GetChargePointStatusDB()
}

// GetGroups implements IDal.
func (*DalMock) GetGroups() ([]models.Group, error) {
	return GetGroupsDB()
}

// InitDB implements IDal.
func (*DalMock) InitDB(conn_str string) {
	panic("unimplemented")
}

func TestAlgorithm3Charging(t *testing.T) {
	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	// groups := []models.Group{
	// 	{Id: 2, MaxCurrent: 300},
	// 	{Id: 1, MaxCurrent: 500},
	// }
	// chargePoints := []models.ChargePointStatus{
	// 	{ChargePointId: 5, Priority: 1, GroupId: 2},
	// 	{ChargePointId: 4, Priority: 3, GroupId: 2},
	// 	{ChargePointId: 1, Priority: 1, GroupId: 1},
	// 	{ChargePointId: 3, Priority: 4, GroupId: 2},
	// 	{ChargePointId: 2, Priority: 4, GroupId: 1},
	// }
	// chargePointConnectors := []models.ChargePointConnector{
	// 	{Id: 5, Status: "Charging", ChargePointId: 2},
	// 	{Id: 6, Status: "Available", ChargePointId: 2},
	// 	{Id: 7, Status: "Charging", ChargePointId: 2},
	// 	{Id: 8, Status: "Available", ChargePointId: 3},
	// 	{Id: 9, Status: "Available", ChargePointId: 3},
	// 	{Id: 10, Status: "Available", ChargePointId: 4},
	// 	{Id: 11, Status: "Charging", ChargePointId: 4},
	// 	{Id: 12, Status: "Available", ChargePointId: 4},
	// 	{Id: 13, Status: "Charging", ChargePointId: 5},
	// 	{Id: 14, Status: "Available", ChargePointId: 5},
	// 	{Id: 15, Status: "Charging", ChargePointId: 5},
	// 	{Id: 1, Status: "Available", ChargePointId: 1},
	// 	{Id: 2, Status: "Available", ChargePointId: 1},
	// 	{Id: 3, Status: "Available", ChargePointId: 1},
	// 	{Id: 4, Status: "Available", ChargePointId: 1},
	// }

	groups := []models.Group{
		{Id: 1, MaxCurrent: 100},
	}

	chargePoints := []models.ChargePointStatus{
		{ChargePointId: 1, Priority: 1, GroupId: 1},
		{ChargePointId: 2, Priority: 3, GroupId: 1},
		{ChargePointId: 3, Priority: 2, GroupId: 1},
	}

	chargePointConnectors := []models.ChargePointConnector{
		{Id: 1, Status: "Charging", ChargePointId: 1},
		{Id: 2, Status: "Charging", ChargePointId: 2},
		{Id: 3, Status: "Charging", ChargePointId: 3},
	}

	GetGroupsDB = func() ([]models.Group, error) {
		return groups, nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		return chargePoints, nil
	}

	GetChargePointConnectorDB = func() ([]models.ChargePointConnector, error) {
		return chargePointConnectors, nil
	}

	resultMap := Algorithm()
	fmt.Println(resultMap)

	expected := map[int]float64{
		1: 16.67,
		2: 50,
		3: 33.33,
	}
	assert.Equal(t, expected, resultMap)
}

func TestAlgorithm2Charging1Available(t *testing.T) {
	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	// groups := []models.Group{
	// 	{Id: 2, MaxCurrent: 300},
	// 	{Id: 1, MaxCurrent: 500},
	// }
	// chargePoints := []models.ChargePointStatus{
	// 	{ChargePointId: 5, Priority: 1, GroupId: 2},
	// 	{ChargePointId: 4, Priority: 3, GroupId: 2},
	// 	{ChargePointId: 1, Priority: 1, GroupId: 1},
	// 	{ChargePointId: 3, Priority: 4, GroupId: 2},
	// 	{ChargePointId: 2, Priority: 4, GroupId: 1},
	// }
	// chargePointConnectors := []models.ChargePointConnector{
	// 	{Id: 5, Status: "Charging", ChargePointId: 2},
	// 	{Id: 6, Status: "Available", ChargePointId: 2},
	// 	{Id: 7, Status: "Charging", ChargePointId: 2},
	// 	{Id: 8, Status: "Available", ChargePointId: 3},
	// 	{Id: 9, Status: "Available", ChargePointId: 3},
	// 	{Id: 10, Status: "Available", ChargePointId: 4},
	// 	{Id: 11, Status: "Charging", ChargePointId: 4},
	// 	{Id: 12, Status: "Available", ChargePointId: 4},
	// 	{Id: 13, Status: "Charging", ChargePointId: 5},
	// 	{Id: 14, Status: "Available", ChargePointId: 5},
	// 	{Id: 15, Status: "Charging", ChargePointId: 5},
	// 	{Id: 1, Status: "Available", ChargePointId: 1},
	// 	{Id: 2, Status: "Available", ChargePointId: 1},
	// 	{Id: 3, Status: "Available", ChargePointId: 1},
	// 	{Id: 4, Status: "Available", ChargePointId: 1},
	// }

	groups := []models.Group{
		{Id: 1, MaxCurrent: 100},
	}

	chargePoints := []models.ChargePointStatus{
		{ChargePointId: 1, Priority: 1, GroupId: 1},
		{ChargePointId: 2, Priority: 3, GroupId: 1},
		{ChargePointId: 3, Priority: 2, GroupId: 1},
	}

	chargePointConnectors := []models.ChargePointConnector{
		{Id: 1, Status: "Charging", ChargePointId: 1},
		{Id: 2, Status: "Available", ChargePointId: 2},
		{Id: 3, Status: "Charging", ChargePointId: 3},
	}

	GetGroupsDB = func() ([]models.Group, error) {
		return groups, nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		return chargePoints, nil
	}

	GetChargePointConnectorDB = func() ([]models.ChargePointConnector, error) {
		return chargePointConnectors, nil
	}

	resultMap := Algorithm()
	fmt.Println(resultMap)

	expected := map[int]float64{
		1: 33.33,
		3: 66.67,
	}
	assert.Equal(t, expected, resultMap)
}

func TestAlgorithm1Charging2Available(t *testing.T) {
	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	// groups := []models.Group{
	// 	{Id: 2, MaxCurrent: 300},
	// 	{Id: 1, MaxCurrent: 500},
	// }
	// chargePoints := []models.ChargePointStatus{
	// 	{ChargePointId: 5, Priority: 1, GroupId: 2},
	// 	{ChargePointId: 4, Priority: 3, GroupId: 2},
	// 	{ChargePointId: 1, Priority: 1, GroupId: 1},
	// 	{ChargePointId: 3, Priority: 4, GroupId: 2},
	// 	{ChargePointId: 2, Priority: 4, GroupId: 1},
	// }
	// chargePointConnectors := []models.ChargePointConnector{
	// 	{Id: 5, Status: "Charging", ChargePointId: 2},
	// 	{Id: 6, Status: "Available", ChargePointId: 2},
	// 	{Id: 7, Status: "Charging", ChargePointId: 2},
	// 	{Id: 8, Status: "Available", ChargePointId: 3},
	// 	{Id: 9, Status: "Available", ChargePointId: 3},
	// 	{Id: 10, Status: "Available", ChargePointId: 4},
	// 	{Id: 11, Status: "Charging", ChargePointId: 4},
	// 	{Id: 12, Status: "Available", ChargePointId: 4},
	// 	{Id: 13, Status: "Charging", ChargePointId: 5},
	// 	{Id: 14, Status: "Available", ChargePointId: 5},
	// 	{Id: 15, Status: "Charging", ChargePointId: 5},
	// 	{Id: 1, Status: "Available", ChargePointId: 1},
	// 	{Id: 2, Status: "Available", ChargePointId: 1},
	// 	{Id: 3, Status: "Available", ChargePointId: 1},
	// 	{Id: 4, Status: "Available", ChargePointId: 1},
	// }

	groups := []models.Group{
		{Id: 1, MaxCurrent: 100},
	}

	chargePoints := []models.ChargePointStatus{
		{ChargePointId: 1, Priority: 1, GroupId: 1},
		{ChargePointId: 2, Priority: 3, GroupId: 1},
		{ChargePointId: 3, Priority: 2, GroupId: 1},
	}

	chargePointConnectors := []models.ChargePointConnector{
		{Id: 1, Status: "Charging", ChargePointId: 1},
		{Id: 2, Status: "Available", ChargePointId: 2},
		{Id: 3, Status: "Available", ChargePointId: 3},
	}

	GetGroupsDB = func() ([]models.Group, error) {
		return groups, nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		return chargePoints, nil
	}

	GetChargePointConnectorDB = func() ([]models.ChargePointConnector, error) {
		return chargePointConnectors, nil
	}

	resultMap := Algorithm()

	expected := map[int]float64{
		1: 100.0,
	}
	assert.Equal(t, expected, resultMap)
}
