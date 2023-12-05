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
	FindTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByPriceRange(ctx context.Context, min, max int64) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error)
	FilterTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByOldest(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByCheap(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error)
	FiterTicketBySold(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)
}

type TicketRepository interface {
	FindAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
	SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error)
	FindTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByPriceRange(ctx context.Context, min, max int64) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error)
	FilterTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByOldest(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByCheap(ctx context.Context) ([]*entity.Ticket, error)
	FilterTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error)
	FiterTicketBySold(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)
}

type TicketService struct {
	ticketRepository TicketRepository
}

func NewTicketRepository(ticketRepository TicketRepository) *TicketService {
	return &TicketService{
		ticketRepository: ticketRepository,
	}
}

func (s *TicketService) FindAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FindAllTickets(ctx)
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.ticketRepository.CreateTicket(ctx, ticket)
}

func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.ticketRepository.UpdateTicket(ctx, ticket)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int64) error {
	return s.ticketRepository.DeleteTicket(ctx, id)
}

func (s *TicketService) SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error) {
	return s.ticketRepository.SearchTicket(ctx, name)
}

func (s *TicketService) FindTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.ticketRepository.FindTicketByID(ctx, id)
}

func (s *TicketService) FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByCategory(ctx, category)
}

func (s *TicketService) FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByLocation(ctx, location)
}

func (s *TicketService) FilterTicketByPriceRange(ctx context.Context, min, max int64) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByPriceRange(ctx, min, max)
}

func (s *TicketService) FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByPrice(ctx, price)
}

func (s *TicketService) FilterTicketByNewest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByNewest(ctx)
}

func (s *TicketService) FilterTicketByOldest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByOldest(ctx)
}

func (s *TicketService) FilterTicketByCheap(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByCheap(ctx)
}

func (s *TicketService) FilterTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FilterTicketByExpensive(ctx)
}

func (s *TicketService) FiterTicketBySold(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.FiterTicketBySold(ctx)
}

func (s *TicketService) SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error) {
	return s.ticketRepository.SortTicketByAvailable(ctx)
}
