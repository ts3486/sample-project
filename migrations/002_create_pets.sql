-- +goose Up
CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(id),
    name VARCHAR(255) NOT NULL,
    species VARCHAR(50) NOT NULL,
    breed VARCHAR(255),
    size_class VARCHAR(50) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    date_of_birth DATE NOT NULL,
    is_neutered BOOLEAN DEFAULT FALSE,
    has_microchip BOOLEAN DEFAULT FALSE,
    microchip_id VARCHAR(50),
    photo_full_body VARCHAR(500),
    photo_face VARCHAR(500),
    vet_clinic_name VARCHAR(200),
    vet_clinic_phone VARCHAR(20),
    pre_existing TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE pets;