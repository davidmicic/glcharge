package handlers

import (
	c "glcharge/container"
	"glcharge/models"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ChangeGroupMaxCurrentByIdDB     func(id int, maxCurrent float64)
	ChangeConnectorStatusByIdDB     func(id int, status string)
	ChangeChargePointPriorityByIdDB func(id int, priority int)
)

type DalMock struct {
}

// ChangeChargePointPriorityById implements IDal.
func (*DalMock) ChangeChargePointPriorityById(id int, priority int) {
	ChangeChargePointPriorityByIdDB(id, priority)
}

// ChangeConnectorStatusById implements IDal.
func (*DalMock) ChangeConnectorStatusById(id int, status string) {
	ChangeConnectorStatusByIdDB(id, status)
}

// ChangeGroupMaxCurrentById implements IDal.
func (*DalMock) ChangeGroupMaxCurrentById(id int, maxCurrent float64) {
	ChangeGroupMaxCurrentByIdDB(id, maxCurrent)
}

// GetChargePointConnector implements IDal.
func (*DalMock) GetChargePointConnector() ([]models.ChargePointConnector, error) {
	panic("unimplemented")
}

// GetChargePointStatus implements IDal.
func (*DalMock) GetChargePointStatus() ([]models.ChargePointStatus, error) {
	panic("unimplemented")
}

// GetGroups implements IDal.
func (*DalMock) GetGroups() ([]models.Group, error) {
	panic("unimplemented")
}

// InitDB implements IDal.
func (*DalMock) InitDB(conn_str string) {
	panic("unimplemented")
}

func TestChangeChargePointPriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: "1",
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeChargePointPriority?chargePointId=1&priority=5", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.ChargePointStatus{
			ChargePointId: "1",
			Priority:      5,
			GroupId:       1,
		}, chargePointStatus)
}

func TestChangeChargePointPriorityBadValueChargePointId(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: "1",
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeChargePointPriority?chargePointId=a&priority=5", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "ChargePointId is not a number", rr.Body.String())
}

func TestChangeChargePointPriorityBadValuePriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: "1",
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeChargePointPriority?chargePointId=1&priority=a", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Priority is not a number", rr.Body.String())
}

func TestChangeConnectorStatus(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) {
		connectorStatus.Status = status
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeConnectorStatus?connectorId=1&status=Available", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.ChargePointConnector{
			Id:            1,
			Status:        "Available",
			ChargePointId: 1,
		}, connectorStatus)
}

func TestChangeConnectorStatusBadValueConnectorId(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) {
		connectorStatus.Status = status
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeConnectorStatus?connectorId=a&status=Available", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "ConnectorId is not a number", rr.Body.String())
}

func TestChangeConnectorStatusBadValueStatusMissing(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) {
		connectorStatus.Status = status
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeConnectorStatus?connectorId=1&status=", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Missing status value", rr.Body.String())
}

func TestChangeConnectorStatusBadValueStatusWrong(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) {
		connectorStatus.Status = status
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeConnectorStatus?connectorId=1&status=asd", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Wrong status value", rr.Body.String())
}

func TestChangeMaxCurrentGroup(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         "1",
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeMaxCurrentGroup?groupId=1&MaxCurrent=500.0", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.Group{
			Id:         "1",
			MaxCurrent: 500.0,
		}, groupMaxCurrent)
}

func TestChangeMaxCurrentGroupBadValueGroupId(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         "1",
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeMaxCurrentGroup?groupId=a&MaxCurrent=500.0", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "GroupId is not a number", rr.Body.String())
}

func TestChangeMaxCurrentGroupBadValueMaxCurrent(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         "1",
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("GET", "/changeMaxCurrentGroup?groupId=1&MaxCurrent=a", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "MaxCurrent is not a number", rr.Body.String())
}
