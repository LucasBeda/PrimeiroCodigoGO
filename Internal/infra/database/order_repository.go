package database

import (
	"database/sql"

	"github.com/devfullcycle/go-intensivo-jul/Internal/entity"
)

type OrderRespository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRespository {
	return &OrderRespository{
		DB: db,
	}
}

func (r *OrderRespository) Save(order *entity.Order) error {
	_, err := r.DB.Exec("Insert into Orders (id, price, tax, final_price) Values (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRespository) GetTotalTransactions() (int, error) {
	var total int
	err := r.DB.QueryRow("Select COUNT(*) From Orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
