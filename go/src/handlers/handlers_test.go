package handlers

import (
	"encoding/json"
	"errors"
	c "glcharge/go/src/container"
	"glcharge/go/src/models"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ChangeGroupMaxCurrentByIdDB     func(id int, maxCurrent float64) error
	ChangeConnectorStatusByIdDB     func(id int, status string) error
	ChangeChargePointPriorityByIdDB func(id int, priority int) error
	GetChargePointStatusDB          func() ([]models.ChargePointStatus, error)
	GetChargePointStatusByIdDB      func() (models.ChargePointStatus, error)
	GetChargePointStatusGroupIdDB   func() ([]models.ChargePointStatus, error)
)

type DalMock struct {
}

// GetChargePointStatusById implements storage.IDal.
func (*DalMock) GetChargePointStatusById(id int) (models.ChargePointStatus, error) {
	return GetChargePointStatusByIdDB()
}

// GetChargePointStatusGroupId implements storage.IDal.
func (*DalMock) GetChargePointStatusGroupId(id int) ([]models.ChargePointStatus, error) {
	return GetChargePointStatusGroupIdDB()
}

// AddChargePoint implements storage.IDal.
func (*DalMock) AddChargePoint(priority int, groupId int) error {
	return nil
}

// AddChargePointConnector implements storage.IDal.
func (*DalMock) AddChargePointConnector(status string, chargePointId int) error {
	return nil
}

// AddGroup implements storage.IDal.
func (*DalMock) AddGroup(maxCurrent float64) error {
	return nil
}

// ChangeChargePointPriorityById implements storage.IDal.
func (*DalMock) ChangeChargePointPriorityById(id int, priority int) error {
	return ChangeChargePointPriorityByIdDB(id, priority)
}

// ChangeConnectorStatusById implements storage.IDal.
func (*DalMock) ChangeConnectorStatusById(id int, status string) error {
	return ChangeConnectorStatusByIdDB(id, status)
}

// ChangeGroupMaxCurrentById implements storage.IDal.
func (*DalMock) ChangeGroupMaxCurrentById(id int, maxCurrent float64) error {
	return ChangeGroupMaxCurrentByIdDB(id, maxCurrent)
}

// GetChargePointConnector implements storage.IDal.
func (*DalMock) GetChargePointConnector() ([]models.ChargePointConnector, error) {
	return nil, nil
}

// GetChargePointStatus implements storage.IDal.
func (*DalMock) GetChargePointStatus() ([]models.ChargePointStatus, error) {
	return GetChargePointStatusDB()
}

// GetGroups implements storage.IDal.
func (*DalMock) GetGroups() ([]models.Group, error) {
	return nil, nil
}

// InitDB implements storage.IDal.
func (*DalMock) InitDB(conn_str string) {
	panic("unimplemented")
}

func TestChangeChargePointPriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      0,
		GroupId:       1,
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return chargePointStatus, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		chargePointStatus.Priority = priority
		return nil
	}

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": 1}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.ChargePointStatus{
			ChargePointId: 1,
			Priority:      1,
			GroupId:       1,
		}, chargePointStatus)
}

func TestChangeChargePointPriorityWrongValueChargePointId(t *testing.T) {
	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		return errors.New("")
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	body := strings.NewReader(`{"ChargePointId": 3, "Priority": 1}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Result().StatusCode)
}

func TestChangeChargePointPriorityWrongTypeChargePointId(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		chargePointStatus.Priority = priority
		return nil
	}

	body := strings.NewReader(`{"ChargePointId": "a", "Priority": 2}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Result().StatusCode)
}

func TestChangeChargePointPriorityWrongTypePriority(t *testing.T) {

	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		chargePointStatus.Priority = priority
		return nil
	}

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": "a"}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Result().StatusCode)
}

func TestChangeChargePointPriorityWrongValuePriority2(t *testing.T) {
	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
		var chargePoints []models.ChargePointStatus
		chargePoints = append(chargePoints,
			models.ChargePointStatus{
				ChargePointId: 1,
				Priority:      1,
				GroupId:       1,
			},
		)
		return chargePoints, nil
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		chargePointStatus.Priority = priority
		return nil
	}

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": 3}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "Priority should be 0", response["message"])
}

func TestChangeChargePointPriorityPriorityTooLow(t *testing.T) {
	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": -3}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "Priority should not be negative", response["message"])
}

func TestChangeChargePointPriorityWrongValuePriority(t *testing.T) {
	var chargePointStatus = models.ChargePointStatus{
		ChargePointId: 1,
		Priority:      1,
		GroupId:       1,
	}

	ChangeChargePointPriorityByIdDB = func(id int, priority int) error {
		chargePointStatus.Priority = priority
		return nil
	}

	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
		var chargePoints []models.ChargePointStatus
		chargePoints = append(chargePoints,
			models.ChargePointStatus{
				ChargePointId: 1,
				Priority:      1,
				GroupId:       1,
			},
		)
		return chargePoints, nil
	}

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": 3}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "Priority should be 0", response["message"])
}

func TestChangeChargePointPriorityPriorityTooHigh(t *testing.T) {
	GetChargePointStatusByIdDB = func() (models.ChargePointStatus, error) {
		return models.ChargePointStatus{}, nil
	}

	GetChargePointStatusGroupIdDB = func() ([]models.ChargePointStatus, error) {
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

	body := strings.NewReader(`{"ChargePointId": 1, "Priority": 3}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeChargePointPriority", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "Priority should be between 0 and 1", response["message"])
}

func TestChangeConnectorStatus(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) error {
		connectorStatus.Status = status
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "Status": "Available"}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeConnectorStatus", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.ChargePointConnector{
			Id:            1,
			Status:        "Available",
			ChargePointId: 1,
		}, connectorStatus)
}

func TestChangeConnectorStatusWrongValueConnectorId(t *testing.T) {
	ChangeConnectorStatusByIdDB = func(id int, status string) error {
		return errors.New("")
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 4, "Status": "Available"}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeConnectorStatus", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Result().StatusCode)
}

func TestChangeConnectorStatusWrongTypeConnectorId(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) error {
		connectorStatus.Status = status
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": a, "Status": "Available"}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeConnectorStatus", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Result().StatusCode)
}

func TestChangeConnectorStatusWrongValueStatus(t *testing.T) {
	var connectorStatus = models.ChargePointConnector{
		Id:            1,
		Status:        "Charging",
		ChargePointId: 1,
	}

	ChangeConnectorStatusByIdDB = func(id int, status string) error {
		connectorStatus.Status = status
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "Status": "asd"}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeConnectorStatus", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "Wrong status value", response["message"])
}

func TestChangeMaxCurrentGroup(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) error {
		groupMaxCurrent.MaxCurrent = maxCurrent
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "MaxCurrent": 400.0}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t,
		models.Group{
			Id:         1,
			MaxCurrent: 400.0,
		}, groupMaxCurrent)
}

func TestChangeMaxCurrentGroupWrongGroupId(t *testing.T) {
	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) error {
		return errors.New("")
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "MaxCurrent": 400.0}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Result().StatusCode)
}

func TestChangeMaxCurrentGroupWrongTypeGroupId(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) error {
		groupMaxCurrent.MaxCurrent = maxCurrent
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": a, "MaxCurrent": 400.0}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Result().StatusCode)
}

func TestChangeMaxCurrentGroupWrongTypeMaxCurrent(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) error {
		groupMaxCurrent.MaxCurrent = maxCurrent
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "MaxCurrent": a}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup", body)
	rr := httptest.NewRecorder()

	Makehttphandlers().ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Result().StatusCode)
}

func TestChangeMaxCurrentGroupLessThanZero(t *testing.T) {
	var groupMaxCurrent = models.Group{
		Id:         1,
		MaxCurrent: 300.0,
	}

	ChangeGroupMaxCurrentByIdDB = func(id int, maxCurrent float64) error {
		groupMaxCurrent.MaxCurrent = maxCurrent
		return nil
	}

	GetChargePointStatusDB = func() ([]models.ChargePointStatus, error) { return nil, nil }

	body := strings.NewReader(`{"Id": 1, "MaxCurrent": -1}`)

	container := c.ResetContainer()
	mockStorage := new(DalMock)
	container.SetStorage(mockStorage)
	req := httptest.NewRequest("PUT", "/changeMaxCurrentGroup", body)
	rr := httptest.NewRecorder()
	Makehttphandlers().ServeHTTP(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(t, "MaxCurrent must be greater or equal to 0", response["message"])
}
