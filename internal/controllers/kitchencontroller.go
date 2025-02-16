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

func NewKitchenController(
		path string, router *gin.Engine, logger *zap.Logger,
		kitchenService service.KitchenService,
	) *KitchenController {
	kc := &KitchenController{logger, kitchenService}
	router.POST(path, kc.acceptTicket)
	return kc
}

func (kc *KitchenController) acceptTicket(c *gin.Context) {
	ticketIDstr := c.Param("ticketID")
	ticketID, err := strconv.ParseUint(ticketIDstr, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var ticket TicketAcceptance
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := kc.kitchenService.Accept(context.Background(), uint(ticketID), ticket.ReadyBy); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
