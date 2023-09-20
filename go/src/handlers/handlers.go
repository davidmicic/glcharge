package handlers

import (
	"fmt"
	"glcharge/go/src/algorithm"
	"glcharge/go/src/container"
	"net/http"
	"strconv"

	_ "glcharge/go/src/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Makehttphandlers() *gin.Engine {
	r := gin.Default()
	r.PUT("/changeChargePointPriority/:chargePointId/:priority", changeChargePointPriority)
	r.PUT("/changeConnectorStatus/:connectorId/:status", changeConnectorStatus)
	r.PUT("/changeMaxCurrentGroup/:groupId/:maxCurrent", changeMaxCurrentGroup)
	r.POST("/addGroup", addGroup)
	r.POST("/addChargePoint", addChargePoint)
	r.POST("/addChargePointConnector", addChargePointConnector)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}

// @Param priority   path int true "priority"
// @Param chargePointId path int true "chargePointId"
// @Failure 500 {object} string "ChargePointId is not a number"
// @Failure 500 {object} string "Priority is not a number"
// @Success 204 {string} string "No Content"
// @Router /changeChargePointPriority/:chargePointId/:priority [put]
func changeChargePointPriority(c *gin.Context) {
	fmt.Println("called: changeChargePointPriority")

	chargePointId, err := strconv.Atoi(c.Param("chargePointId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ChargePointId is not a number",
		})
		return
	}
	priority, err := strconv.Atoi(c.Param("priority"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Priority is not a number",
		})
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()

	chargePoints, _ := storage.GetChargePointStatus()
	numberOfChargePoints := len(chargePoints)

	if priority > numberOfChargePoints || priority < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Priority should be between 0 and %d", numberOfChargePoints),
		})
		return
	}

	storage.ChangeChargePointPriorityById(chargePointId, priority)
	resultMap := algorithm.Algorithm()
	c.JSON(http.StatusOK, resultMap)
}

// @Param status   path int true "status"
// @Param connectorId path int true "connectorId"
// @Failure 500 {object} string "ConnectorId is not a number"
// @Failure 500 {object} string "Missing status value"
// @Success 204 {string} string "Wrong status value"
// @Router /changeConnectorStatus/:connectorId/:status [put]
func changeConnectorStatus(c *gin.Context) {
	connectorId, err := strconv.Atoi(c.Param("connectorId"))
	status := c.Param("status")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ConnectorId is not a number"})
		return
	}
	if status == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Missing status value"})
		return
	}
	if status != "Available" && status != "Charging" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong status value"})
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()
	storage.ChangeConnectorStatusById(connectorId, status)
}

// @Param maxCurrent   path int true "maxCurrent"
// @Param groupId path int true "groupId"
// @Failure 500 {object} string "GroupId is not a number"
// @Failure 500 {object} string "MaxCurrent is not a number"
// @Failure 500 {object} string "MaxCurrent must be greater or equal to 0"
// @Success 204 {string} string "No Content"
// @Router /changeMaxCurrentGroup/:groupId/:maxCurrent [put]
func changeMaxCurrentGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GroupId is not a number"})
		return
	}

	maxCurrent, err := strconv.ParseFloat(c.Param("maxCurrent"), 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "MaxCurrent is not a number"})
		return
	} else if maxCurrent < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "MaxCurrent must be greater or equal to 0"})
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()
	storage.ChangeGroupMaxCurrentById(groupId, maxCurrent)
}

type GroupReq struct {
	MaxCurrent float64
}

// @Param maxCurrent   path int true "maxCurrent"
// @Success 204 {string} string "No Content"
// @Router /addGroup [post]
func addGroup(c *gin.Context) {
	var req GroupReq
	c.BindJSON(&req)
	cnt := container.GetContainer()
	storage := cnt.Storage()
	storage.AddGroup(req.MaxCurrent)
}

type ChargPointReq struct {
	Priority int
	GroupId  int
}

// @Param Priority   path int true "Priority"
// @Param GroupId   path int true "GroupId"
// @Success 204 {string} string "No Content"
// @Router /addChargePoint [post]
func addChargePoint(c *gin.Context) {
	var req ChargPointReq
	c.BindJSON(&req)
	cnt := container.GetContainer()
	storage := cnt.Storage()
	storage.AddChargePoint(req.Priority, req.GroupId)
}

type ChargePointConnector struct {
	ChargePointId int
	Status        string
}

// @Param ChargePointId   path int true "ChargePointId"
// @Param Status   path int true "Status"
// @Success 204 {string} string "No Content"
// @Router /addChargePointConnector [post]
func addChargePointConnector(c *gin.Context) {
	var req ChargePointConnector
	c.BindJSON(&req)
	cnt := container.GetContainer()
	storage := cnt.Storage()
	storage.AddChargePointConnector(req.Status, req.ChargePointId)
}
