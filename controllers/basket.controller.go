package controllers

import (
	"github.com/arsalanaa44/basketball/enum"
	"net/http"
	"strconv"
	"time"

	"github.com/arsalanaa44/basketball/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BasketController struct {
	DB *gorm.DB
}

func NewBasketController(DB *gorm.DB) BasketController {
	return BasketController{DB}
}

func (bc *BasketController) CreateBasket(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.BasketRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newBasket := models.Basket{
		Data:      payload.Data,
		Status:    payload.Status,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := bc.DB.Create(&newBasket)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBasket})
}

func (bc *BasketController) UpdateBasket(ctx *gin.Context) {
	basketId := ctx.Param("basketId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.BasketRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedBasket models.Basket
	result := bc.DB.Where(`id = ? AND "user" = ?`, basketId, currentUser.ID).First(&updatedBasket)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No basket found"})
		return
	}
	if updatedBasket.Status == enum.Completed {

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "the basket status is completed"})
		return
	}

	basketToUpdate := models.Basket{
		Data:   payload.Data,
		Status: payload.Status,
	}

	bc.DB.Model(&updatedBasket).Updates(basketToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedBasket})
}

func (bc *BasketController) FindBasketById(ctx *gin.Context) {
	basketId := ctx.Param("basketId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var basket models.Basket
	result := bc.DB.First(&basket, "id = ?", basketId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No basket found"})
		return
	}
	if basket.User != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": basket})
}

func (bc *BasketController) FindBaskets(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	currentUser := ctx.MustGet("currentUser").(models.User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var baskets []models.Basket
	results := bc.DB.Where(`"user" = ? `, currentUser.ID).Limit(intLimit).Offset(offset).Find(&baskets)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(baskets), "data": baskets})
}

func (bc *BasketController) DeleteBasket(ctx *gin.Context) {
	basketId := ctx.Param("basketId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	// check authorized existence
	var basket models.Basket
	if result := bc.DB.First(&basket, "id = ?", basketId); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No basket found"})
		return
	} else if basket.User != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	result := bc.DB.Delete(&models.Basket{}, "id = ?", basketId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "error occurred"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
