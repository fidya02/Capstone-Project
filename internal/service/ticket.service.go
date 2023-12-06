package service

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"
)

type TicketUseCase interface {
	FindAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
	SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error)
	SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByOldest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByCheap(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketBySold(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)
}

type TicketRepository interface {
	FindAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
	SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error)
	SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByOldest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByCheap(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketBySold(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)
}

type TicketService struct {
	Repository TicketRepository
}

func NewTicketRepository(Repository TicketRepository) *TicketService {
	return &TicketService{
		Repository: Repository,
	}
}

func (s *TicketService) FindAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.FindAllTickets(ctx)
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.CreateTicket(ctx, ticket)
}

func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.UpdateTicket(ctx, ticket)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int64) error {
	return s.Repository.DeleteTicket(ctx, id)
}

func (s *TicketService) SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error) {
	return s.Repository.SearchTicket(ctx, name)
}

func (s *TicketService) FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByRangeTime(ctx, start, end)
}

func (s *TicketService) FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByCategory(ctx, category)
}

func (s *TicketService) FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByLocation(ctx, location)
}

func (s *TicketService) FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByPrice(ctx, price)
}

func (s *TicketService) SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByNewest(ctx)
}

func (s *TicketService) SortTicketByOldest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByOldest(ctx)
}

func (s *TicketService) SortTicketByCheap(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByCheap(ctx)
}

func (s *TicketService) SortTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByExpensive(ctx)
}

func (s *TicketService) SortTicketBySold(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketBySold(ctx)
}

func (s *TicketService) SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByAvailable(ctx)
}
