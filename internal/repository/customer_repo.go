package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ts3486/pet-insurance-api/internal/model"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *model.Customer) error {
	query := `INSERT INTO customers (email, password_hash, first_name, last_name, phone, post_code, address)
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              RETURNING id, created_at`
	return r.db.QueryRow(query,
		customer.Email, customer.Password, customer.FirstName, customer.LastName,
		customer.Phone, customer.PostCode, customer.Address,
	).Scan(&customer.ID, &customer.CreatedAt)
}

func (r *CustomerRepository) GetByID(id int) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Get(&customer, "SELECT * FROM customers WHERE id = $1", id)
	return &customer, err
}

func (r *CustomerRepository) GetByEmail(email string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Get(&customer, "SELECT * FROM customers WHERE email = $1", email)
	return &customer, err
}
