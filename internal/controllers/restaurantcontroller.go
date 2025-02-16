package controllers

import (
	"errors"
	"net/http"

	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/OrtemRepos/KitchenService/internal/ports"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RestaurantController struct {
	logger               *zap.Logger
	restaurantRepository ports.RestaurantRepository
}

func NewRestaurantController(
		path string, router *gin.Engine,
		logger *zap.Logger, restaurantRepository ports.RestaurantRepository,
	) *RestaurantController {
		rc := &RestaurantController{logger, restaurantRepository}
		router.GET(path, rc.getRestaurant)
		return rc
	}

func (rc *RestaurantController) getRestaurant(c *gin.Context) {
	var restaurantID uint 
	if err := c.ShouldBindQuery(&restaurantID); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	restaurant, err := rc.restaurantRepository.Load(restaurantID)
	if errors.Is(err, domain.ErrNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		rc.logger.Error("error when load from the repository", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, restaurant)
}