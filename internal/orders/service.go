package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Pranav2259/ecomm/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrProductNotStock = errors.New("products ordered qty not in stock")
)

type Service interface {
	PlaceOrder(context.Context, createOrderParam) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParam) (repo.Order, error) {

	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("valid customerID is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("add atleast one item")
	}

	// create a transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	// create order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.Items {
		// check if product exists in db
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNotStock

		}

		// place order if exists
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		if err != nil {
			return repo.Order{}, err
		}

		// newqty := product.Quantity - item.Quantity
		// _,err := s.d
	}

	tx.Commit(ctx)

	return order, nil
}
