package handlers

import (
	"fmt"
	"glcharge/go/src/algorithm"
	"glcharge/go/src/container"
	"net/http"

	_ "glcharge/go/src/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Makehttphandlers() *gin.Engine {
	r := gin.Default()
	r.PUT("/changeChargePointPriority", changeChargePointPriority)
	r.PUT("/changeConnectorStatus", changeConnectorStatus)
	r.PUT("/changeMaxCurrentGroup", changeMaxCurrentGroup)
	r.POST("/addGroup", addGroup)
	r.POST("/addChargePoint", addChargePoint)
	r.POST("/addChargePointConnector", addChargePointConnector)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}

func checkConnectorStatus(c *gin.Context, status string) bool {
	if status != "Available" && status != "Charging" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong status value"})
		c.Abort()
		return false
	}
	return true
}

func checkMaxCurrentValue(c *gin.Context, maxCurrent float64) bool {
	if maxCurrent < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "MaxCurrent must be greater or equal to 0"})
		c.Abort()
		return false
	}
	return true
}

func return500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "Internal server error")
}

type HandlerRes struct {
	ResultMap map[int]float64 `json:"resultMap"`
}

type ChangeChargePointPriorityReq struct {
	ChargePointId int `json:"ChargePointId" example:"1"`
	Priority      int `json:"Priority" example:"1"`
}

type ChangeConnectorStatusReq struct {
	Id     int    `json:"Id" example:"1"`
	Status string `json:"Status" example:"Available"`
}

type ChangeMaxCurrentGroupReq struct {
	Id         int     `json:"Id" example:"1"`
	MaxCurrent float64 `json:"MaxCurrent" example:"100.0"`
}

type GroupReq struct {
	MaxCurrent float64 `json:"MaxCurrent" example:"100.0"`
}

type ChargePointReq struct {
	Priority int `json:"Priority" example:"1"`
	GroupId  int `json:"GroupId" example:"1"`
}

type ChargePointConnectorReq struct {
	ChargePointId int    `json:"ChargePointId" example:"1"`
	Status        string `json:"Status" example:"Available"`
}

// @Accept  json
// @Produce json
// @Param request   body handlers.ChangeChargePointPriorityReq true "ChangeChargePointPriorityReq"
// @Success 200 {object} handlers.HandlerRes
// @Failure 400
// @Failure 500
// @Router /changeChargePointPriority [put]
func changeChargePointPriority(c *gin.Context) {
	var req ChangeChargePointPriorityReq
	c.BindJSON(&req)

	chargePointId := req.ChargePointId
	priority := req.Priority

	cnt := container.GetContainer()
	storage := cnt.Storage()
	chargePoints, err := storage.GetChargePointStatus()
	if err != nil {
		return500(c)
		return
	}

	numberOfChargePoints := len(chargePoints)

	if (numberOfChargePoints == 0 || numberOfChargePoints == 1) && priority != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Priority should be 0",
		})
		c.Abort()
		return
	} else if numberOfChargePoints > 1 && priority >= numberOfChargePoints {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Priority should be between 0 and %d", numberOfChargePoints-1),
		})
		c.Abort()
		return
	} else if priority < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Priority should not be negative",
		})
		c.Abort()
		return
	}

	err = storage.ChangeChargePointPriorityById(chargePointId, priority)
	if err != nil {
		return500(c)
		return
	}

	resultMap := algorithm.Algorithm()
	handlerRes := HandlerRes{
		ResultMap: resultMap,
	}
	c.JSON(http.StatusOK, handlerRes)
}

// @Accept  json
// @Param request   body handlers.ChangeConnectorStatusReq true "ChangeConnectorStatusReq"
// @Success 200 {object} handlers.HandlerRes
// @Failure 400
// @Failure 500
// @Router /changeConnectorStatus [put]
func changeConnectorStatus(c *gin.Context) {
	var req ChangeConnectorStatusReq
	c.BindJSON(&req)

	connectorId := req.Id
	status := req.Status
	if x := checkConnectorStatus(c, status); !x {
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()

	err := storage.ChangeConnectorStatusById(connectorId, status)
	if err != nil {
		return500(c)
		return
	}

	resultMap := algorithm.Algorithm()
	handlerRes := HandlerRes{
		ResultMap: resultMap,
	}
	c.JSON(http.StatusOK, handlerRes)
}

// @Accept  json
// @Param request   body handlers.ChangeMaxCurrentGroupReq true "ChangeMaxCurrentGroupReq"
// @Success 200 {object} handlers.HandlerRes
// @Failure 400
// @Failure 500
// @Router /changeMaxCurrentGroup [put]
func changeMaxCurrentGroup(c *gin.Context) {

	var req ChangeMaxCurrentGroupReq
	c.BindJSON(&req)

	groupId := req.Id
	maxCurrent := req.MaxCurrent
	if x := checkMaxCurrentValue(c, maxCurrent); !x {
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()
	err := storage.ChangeGroupMaxCurrentById(groupId, maxCurrent)
	if err != nil {
		return500(c)
		return
	}

	resultMap := algorithm.Algorithm()
	handlerRes := HandlerRes{
		ResultMap: resultMap,
	}
	c.JSON(http.StatusOK, handlerRes)
}

// @Param request   body GroupReq true "GroupReq"
// @Success 200
// @Failure 500
// @Router /addGroup [post]
func addGroup(c *gin.Context) {
	var req GroupReq
	c.BindJSON(&req)

	maxCurrent := req.MaxCurrent
	if x := checkMaxCurrentValue(c, maxCurrent); !x {
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()
	err := storage.AddGroup(maxCurrent)
	if err != nil {
		return500(c)
		return
	}
}

// @Param request   body ChargePointReq true "ChargePointReq"
// @Success 200
// @Failure 500
// @Router /addChargePoint [post]
func addChargePoint(c *gin.Context) {
	var req ChargePointReq
	c.BindJSON(&req)
	priority := req.Priority

	cnt := container.GetContainer()
	storage := cnt.Storage()
	chargePoints, err := storage.GetChargePointStatus()
	if err != nil {
		return500(c)
		return
	}
	numberOfChargePoints := len(chargePoints)

	if numberOfChargePoints == 0 && priority != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Priority should be 0",
		})
		c.Abort()
		return
	} else if numberOfChargePoints > 0 && priority > numberOfChargePoints {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Priority should be between 0 and %d", numberOfChargePoints),
		})
		c.Abort()
		return
	} else if priority < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Priority should not be negative",
		})
		c.Abort()
		return
	}

	err = storage.AddChargePoint(req.Priority, req.GroupId)
	if err != nil {
		return500(c)
		return
	}
}

// @Param request   body ChargePointConnectorReq true "ChargePointConnectorReq"
// @Success 200
// @Failure 500
// @Router /addChargePointConnector [post]
func addChargePointConnector(c *gin.Context) {
	var req ChargePointConnectorReq
	c.BindJSON(&req)

	chargePointId := req.ChargePointId
	status := req.Status
	if x := checkConnectorStatus(c, status); !x {
		return
	}

	cnt := container.GetContainer()
	storage := cnt.Storage()
	err := storage.AddChargePointConnector(status, chargePointId)
	if err != nil {
		return500(c)
		return
	}
}
