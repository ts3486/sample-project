package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ts3486/pet-insurance-api/internal/model"
)

type BreedRepository struct {
	db *sqlx.DB
}

func NewBreedRepo(db *sqlx.DB) *BreedRepository {
	return &BreedRepository{db: db}
}

func (r *BreedRepository) GetByBreed(species model.Species, breed string) (*model.BreedRate, error) {
	var rate model.BreedRate
	err := r.db.Get(&rate, "SELECT * FROM breed_rates WHERE species = $1 AND breed = $2", species, breed)
	return &rate, err
}

func (r *BreedRepository) ListBySpecies(species model.Species) ([]model.BreedRate, error) {
	var rates []model.BreedRate
	err := r.db.Select(&rates, "SELECT * FROM breed_rates WHERE species = $1", species)
	return rates, err
}
