package service

import (
	"context"
	"time"

	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/OrtemRepos/KitchenService/internal/ports"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type KitchenService struct {
	logger *zap.Logger

	ticketRepository           ports.TicketRepository
	ticketDomainEventPublisher ports.DomainEventPublisher
	restaurantRepository       ports.RestaurantRepository
}

func (ks *KitchenService) CreateMenu(id uint, menu []domain.MenuItem) error {
	restaurant := domain.Restaurant{ID: id, Menu: menu}
	return ks.restaurantRepository.Save(restaurant)
}

func (ks *KitchenService) CreateTicket(ctx context.Context, tx sqlx.Tx,
	restaurantID, ticketID uint,
	lineItems []domain.TicketLineItem) (*domain.Ticket, error) {
	ticket := domain.NewTicket(restaurantID, ticketID, lineItems)

	if err := ks.ticketRepository.Save(tx, *ticket); err != nil {
		ks.logger.Error("ticket saving error", zap.Error(err))
		return nil, err
	}
	return ticket, nil
}

func (ks *KitchenService) ConfirmCreate(ctx context.Context, tx sqlx.Tx, ticketID uint) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.ConfirmCreate()
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}

func (ks *KitchenService) CancelCreate(ctx context.Context, tx sqlx.Tx, ticketID uint) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.CancelCreate()
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}

func (ks *KitchenService) Cancel(ctx context.Context, tx sqlx.Tx, ticketID uint) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.Cancel()
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}

func (ks *KitchenService) ConfirmCancel(ctx context.Context, tx sqlx.Tx, ticketID uint) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.ConfirmCancel()
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}

func (ks *KitchenService) UndoCancel(ctx context.Context, tx sqlx.Tx, ticketID uint) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.UndoCancel()
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}

func (ks *KitchenService) Accept(ctx context.Context, tx sqlx.Tx, ticketID uint, redyBy time.Time) error {
	ticket, err := ks.ticketRepository.Load(tx, ticketID)
	if err != nil {
		ks.logger.Error(
			"error when load from repository",
			zap.Uint("ticket_id", ticketID),
			zap.Error(err),
		)
		return err
	}

	event, err := ticket.Accept(redyBy)
	if err != nil{
		return err
	}

	if err := ks.ticketDomainEventPublisher.Publish(event); err != nil {
		ks.logger.Error("event publishing error", zap.Error(err))
		return err
	}
	return nil
}