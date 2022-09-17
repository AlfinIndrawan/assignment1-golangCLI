package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders() ([]model.Order, error)
	Create(order model.Order) (model.Order, error)
	GetOrder(id string) (model.Order, error)
	Update(order model.Order) (model.Order, error)
	Delete(id string) error
}

type OrderRepo struct {
	connection *gorm.DB
}

func NewRepository(db *gorm.DB) OrderRepository {
	return &OrderRepo{
		connection: db,
	}
}

// get all orders with items
func (r *OrderRepo) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	res := r.connection.Preload("Items").Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}

//crete order with items
func (r *OrderRepo) Create(order model.Order) (model.Order, error) {
	err := r.connection.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

// find order by id
func (r *OrderRepo) GetOrder(id string) (model.Order, error) {
	var order model.Order
	res := r.connection.Preload("Items").Where("id = ?", id).Take(&order)
	if res.Error != nil {
		return order, res.Error
	}
	return order, nil
}

func (r *OrderRepo) Update(order model.Order) (model.Order, error) {
	r.connection.Save(&order)
	r.connection.Preload("Items").Find(&order)
	return order, nil
}

func (r *OrderRepo) Delete(id string) error {
	var order model.Order
	res := r.connection.Preload("Items").Where("id = ?", id).Take(&order)
	if res.Error != nil {
		return res.Error
	}
	r.connection.Delete(&order)
	return nil
}
