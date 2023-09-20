package handlers

import (
	"encoding/json"
	c "glcharge/go/src/container"
	"glcharge/go/src/models"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ChangeGroupMaxCurrentByIdDB     func(id int, maxCurrent float64)
	ChangeConnectorStatusByIdDB     func(id int, status string)
	ChangeChargePointPriorityByIdDB func(id int, priority int)
	GetChargePointStatusDB          func() ([]models.ChargePointStatus, error)
)

type DalMock struct {
}

// GetGroups implements storage.IDal.
func (*DalMock) GetGroups() ([]models.Group, error) {
	panic("unimplemented")
}

// AddChargePoint implements storage.IDal.
func (*DalMock) AddChargePoint(priority int, groupId int) {
	panic("unimplemented")
}

// AddChargePointConnector implements storage.IDal.
func (*DalMock) AddChargePointConnector(status string, chargePointId int) {
	panic("unimplemented")
}

// AddGroup implements storage.IDal.
func (*DalMock) AddGroup(maxCurrent float64) {
	panic("unimplemented")
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
	return nil, nil
}

// GetChargePointStatus implements IDal.
func (*DalMock) GetChargePointStatus() ([]models.ChargePointStatus, error) {
	return GetChargePointStatusDB()
}

// InitDB implements IDal.
func (*DalMock) InitDB(conn_str string) {
	panic("unimplemented")
}

func TestChangeChargePointPriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		var chargePoints []models.ChargePointStatus
		chargePoints = append(chargePoints,
			models.ChargePointStatus{
				ChargePointId: 1,
				Priority:      1,
				GroupId:       1,
			},
			models.ChargePointStatus{
				ChargePointId: 2,
				Priority:      1,
				GroupId:       1,
			},
		)
		return chargePoints, nil
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority/1/2", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.ChargePointStatus{
			ChargePointId: 1,
			Priority:      2,
			GroupId:       1,
		}, chargePointStatus)
}

func TestChangeChargePointPriorityBadValueChargePointId(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority/a/5", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "ChargePointId is not a number", response["message"])
}

func TestChangeChargePointPriorityBadValuePriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) {
		chargePointStatus.Priority = priority
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority/1/a", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Priority is not a number", response["message"])
}

func TestChangeChargePointPriorityBadValuePriorityTooLow(t *testing.T) {
	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		var chargePoints []models.ChargePointStatus
		chargePoints = append(chargePoints,
			models.ChargePointStatus{
				ChargePointId: 1,
				Priority:      1,
				GroupId:       1,
			},
			models.ChargePointStatus{
				ChargePointId: 2,
				Priority:      1,
				GroupId:       1,
			},
		)
		return chargePoints, nil
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority/1/-4", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Priority should be between 0 and 2", response["message"])
}

func TestChangeChargePointPriorityBadValuePriorityTooHigh(t *testing.T) {
	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) {
		var chargePoints []models.ChargePointStatus
		chargePoints = append(chargePoints,
			models.ChargePointStatus{
				ChargePointId: 1,
				Priority:      1,
				GroupId:       1,
			},
			models.ChargePointStatus{
				ChargePointId: 2,
				Priority:      1,
				GroupId:       1,
			},
		)
		return chargePoints, nil
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority/1/3", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Priority should be between 0 and 2", response["message"])
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
	req := httptest.NewRequest("PUT", "/changeConnectorStatus/1/Available", nil)
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
	req := httptest.NewRequest("PUT", "/changeConnectorStatus/a/Available", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "ConnectorId is not a number", response["message"])
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
	req := httptest.NewRequest("PUT", "/changeConnectorStatus/1/asd", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "Wrong status value", response["message"])
}

func TestChangeMaxCurrentGroup(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup/1/500.0", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.Group{
			Id:         1,
			MaxCurrent: 500.0,
		}, groupMaxCurrent)
}

func TestChangeMaxCurrentGroupBadValueGroupId(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup/a/500.0", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "GroupId is not a number", response["message"])
}

func TestChangeMaxCurrentGroupBadValueMaxCurrent(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup/1/a", nil)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "MaxCurrent is not a number", response["message"])
}

func TestChangeMaxCurrentGroupBadValueLessThanZero(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) {
		groupMaxCurrent.MaxCurrent = maxCurrent
	}

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup/1/-1", nil)
	rr := httptest.NewRecorder()
	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 500, rr.Result().StatusCode)
	assert.Equal(t, "MaxCurrent must be greater or equal to 0", response["message"])
}
