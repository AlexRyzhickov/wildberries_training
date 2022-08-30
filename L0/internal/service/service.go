package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"wildberries_traineeship/internal/models"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s Service) GetOrderInfo(ctx context.Context, id string) (*models.OrderData, error) {
	order := models.Order{
		Id: id,
	}
	err := s.db.Model(&models.Order{}).First(&order).Error
	if err != nil {
		return nil, err
	}
	orderData := models.OrderData{}
	err = order.OrderData.AssignTo(&orderData)
	if err != nil {
		return nil, err
	}
	return &orderData, nil
}

type NoOrderInfo struct {
	ticker string
}

func (err NoOrderInfo) Error() string {
	return fmt.Sprintf("No info about the stock with ticker %v", err.ticker)
}
