package controller

import (
	"P2/RESTAPI/api/request"
	"P2/RESTAPI/core/entities"
	"P2/RESTAPI/core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HumansController struct {
	Humans *service.HumansService
}

func NewController(HumansService *service.HumansService)*HumansController  {
	return &HumansController{
		Humans: HumansService,
	}
}

func (u *HumansController)CreateHumans(c *gin.Context)  {
	var createHumans request.Humans

	if err := c.ShouldBindJSON(&createHumans); err != nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	err := u.Humans.InsertIntoHumans(c,&createHumans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK, "successfully")

}
func (u *HumansController)FindAll(c *gin.Context)  {
	var humans []entities.Humans

	arr ,err := u.Humans.FindAll(c,humans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK,gin.H{"data ":arr})
}
func (u *HumansController)FindById(c *gin.Context)  {

	Id := c.Param("id")
	var humans []entities.Humans
	arr ,err := u.Humans.FindById(c,humans,Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK,gin.H{"data ":arr})
}

func (u *HumansController)UpdateById(c *gin.Context)  {
	Id := c.Param("id")
	var humans entities.Humans
	if err := c.ShouldBindJSON(&humans); err != nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	err:= u.Humans.UpdateById(c,humans,Id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK,gin.H{"update":"successfully"})
}

func (u *HumansController)DeleteById(c *gin.Context)  {
	Id := c.Param("id")
	var humans entities.Humans
	err:= u.Humans.DeleteById(c,humans,Id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK,gin.H{"Delete":"successfully"})
}





