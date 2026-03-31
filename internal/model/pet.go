package model

import "time"

type Species string

const (
	SpeciesDog Species = "dog"
	SpeciesCat Species = "cat"
)

type SizeClass string

const (
	SizeSmall  SizeClass = "small"
	SizeMedium SizeClass = "medium"
	SizeLarge  SizeClass = "large"
)

type Pet struct {
	ID             int       `json:"id" db:"id"`
	CustomerID     int       `json:"customer_id" db:"customer_id"`
	Name           string    `json:"name" db:"name"`
	Species        Species   `json:"species" db:"species"`       // "dog" or "cat"
	Breed          string    `json:"breed" db:"breed"`           // e.g. "トイプードル"
	SizeClass      SizeClass `json:"size_class" db:"size_class"` // "small", "medium", "large" (dogs only)
	Gender         string    `json:"gender" db:"gender"`         // "male" or "female"
	DateOfBirth    time.Time `json:"date_of_birth" db:"date_of_birth"`
	IsNeutered     bool      `json:"is_neutered" db:"is_neutered"`     // 避妊・去勢済み
	HasMicrochip   bool      `json:"has_microchip" db:"has_microchip"` // マイクロチップ装着済み
	MicrochipID    string    `json:"microchip_id,omitempty" db:"microchip_id"`
	PhotoFullBody  string    `json:"photo_full_body,omitempty" db:"photo_full_body"`   // 全身写真 URL
	PhotoFace      string    `json:"photo_face,omitempty" db:"photo_face"`             // 顔写真 URL
	VetClinicName  string    `json:"vet_clinic_name,omitempty" db:"vet_clinic_name"`   // かかりつけ動物病院名
	VetClinicPhone string    `json:"vet_clinic_phone,omitempty" db:"vet_clinic_phone"` // 動物病院電話番号
	PreExisting    string    `json:"pre_existing,omitempty" db:"pre_existing"`         // 既往症
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
