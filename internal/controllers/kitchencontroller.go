package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OrtemRepos/KitchenService/internal/domain/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KitchenController struct {
	logger         *zap.Logger
	kitchenService service.KitchenService
}

func NewKitchenController(router *gin.Engine, logger *zap.Logger, kitchenService service.KitchenService) *KitchenController {
	return &KitchenController{logger, kitchenService}
}

func (kc *KitchenController) acceptTicket(c *gin.Context) {
	ticketIDstr := c.Param("ticketID")
	ticketID, err := strconv.ParseUint(ticketIDstr, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var form ticketAcceptance
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := kc.kitchenService.Accept(context.Background(), uint(ticketID), form.ReadyBy); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}