package app

import "assigment-2-golang-faaza/internal/domain"

type OrderService struct {
    orderRepo *OrderRepository
}

func NewOrderService(orderRepo *OrderRepository) *OrderService {
    return &OrderService{orderRepo}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
    // Add validation or additional logic here if needed
    return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrders() ([]domain.Order, error) {
    return s.orderRepo.GetOrders()
}

func (s *OrderService) UpdateOrder(order *domain.Order) error {
    // Add validation or additional logic here if needed
    return s.orderRepo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
    // Add validation or additional logic here if needed
    return s.orderRepo.DeleteOrder(id)
}
