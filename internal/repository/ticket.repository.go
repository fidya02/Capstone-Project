package repository

import (
	"context"
	"fmt"

	"github.com/fidya02/Capstone-Project/entity"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (r *TicketRepository) FindAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tickets: %+v", tickets)
	return tickets, nil
}

func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	err := r.db.WithContext(ctx).Create(&ticket)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *TicketRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	err := r.db.WithContext(ctx).Save(&ticket).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketRepository) DeleteTicket(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&entity.Ticket{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketRepository) SearchTicket(ctx context.Context, name string) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("category = ?", category).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) FilterTicketByLocation(ctx context.Context, location string) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("location = ?", location).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("Date >= ? AND Date <= ?", start, end).Find(&tickets)
	if err.Error != nil {
		return nil, err.Error
	}
	return tickets, nil
}

func (r *TicketRepository) FilterTicketByPrice(ctx context.Context, price int64) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("price <= ?", price).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("date DESC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketByOldest(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("date ASC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketByExpensive(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("price DESC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketByCheap(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("price ASC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketBySold(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("sold DESC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Order("available DESC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
