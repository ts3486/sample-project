package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ts3486/pet-insurance-api/internal/model"
)

type PetRepository struct {
	db *sqlx.DB
}

func NewPetRepo(db *sqlx.DB) *PetRepository {
	return &PetRepository{db: db}
}

func (r *PetRepository) Create(pet *model.Pet) error {
	query := `INSERT INTO pets (customer_id, name, species, breed, size_class, gender,
              date_of_birth, is_neutered, has_microchip, microchip_id,
              photo_full_body, photo_face, vet_clinic_name, vet_clinic_phone, pre_existing)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
              RETURNING id, created_at`
	return r.db.QueryRow(query,
		pet.CustomerID, pet.Name, pet.Species, pet.Breed, pet.SizeClass, pet.Gender,
		pet.DateOfBirth, pet.IsNeutered, pet.HasMicrochip, pet.MicrochipID,
		pet.PhotoFullBody, pet.PhotoFace, pet.VetClinicName, pet.VetClinicPhone, pet.PreExisting,
	).Scan(&pet.ID, &pet.CreatedAt)
}

func (r *PetRepository) GetByID(id int) (*model.Pet, error) {
	var pet model.Pet
	err := r.db.Get(&pet, "SELECT * FROM pets WHERE id = $1", id)
	return &pet, err
}

func (r *PetRepository) ListByCustomer(customerID int) ([]model.Pet, error) {
	var pets []model.Pet
	err := r.db.Select(&pets, "SELECT * FROM pets WHERE customer_id = $1", customerID)
	return pets, err
}

func (r *PetRepository) Update(pet *model.Pet) error {
	query := `UPDATE pets SET name=$1, species=$2, breed=$3, size_class=$4, gender=$5,
              date_of_birth=$6, is_neutered=$7, has_microchip=$8, microchip_id=$9,
              photo_full_body=$10, photo_face=$11, vet_clinic_name=$12, vet_clinic_phone=$13, pre_existing=$14
              WHERE id=$15`
	_, err := r.db.Exec(query,
		pet.Name, pet.Species, pet.Breed, pet.SizeClass, pet.Gender,
		pet.DateOfBirth, pet.IsNeutered, pet.HasMicrochip, pet.MicrochipID,
		pet.PhotoFullBody, pet.PhotoFace, pet.VetClinicName, pet.VetClinicPhone, pet.PreExisting,
		pet.ID,
	)
	return err
}
