package app

import (
	"assigment-2-golang-faaza/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(order *domain.Order) error {
    return r.db.Create(order).Error
}

func (r *OrderRepository) GetOrders() ([]domain.Order, error) {
    var orders []domain.Order
    if err := r.db.Preload("Items").Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (r *OrderRepository) UpdateOrder(order *domain.Order) error {
    return r.db.Save(order).Error
}

func (r *OrderRepository) DeleteOrder(id uint) error {
    return r.db.Delete(&domain.Order{}, id).Error
}
