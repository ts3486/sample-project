-- +goose Up
CREATE TABLE claims (
    id                SERIAL PRIMARY KEY,
    policy_id         INT NOT NULL REFERENCES policies(id),
    customer_id       INT NOT NULL REFERENCES customers(id),
    pet_id            INT NOT NULL REFERENCES pets(id),
    claim_number      VARCHAR(20) UNIQUE NOT NULL,
    category          VARCHAR(20) NOT NULL,  -- 'outpatient', 'hospital', 'surgery'
    status            VARCHAR(20) DEFAULT 'submitted',
    clinic_name       VARCHAR(200) NOT NULL,
    treatment_date    DATE NOT NULL,
    diagnosis         VARCHAR(500) NOT NULL,
    treatment_detail  TEXT,
    total_amount      BIGINT NOT NULL,       -- 治療費合計 (yen)
    covered_amount    BIGINT DEFAULT 0,      -- 補償対象額 (yen)
    reimbursed_amount BIGINT DEFAULT 0,      -- 保険金額 (yen)
    created_at        TIMESTAMPTZ DEFAULT NOW(),
    updated_at        TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_claims_policy_id ON claims(policy_id);
CREATE INDEX idx_claims_customer_id ON claims(customer_id);
CREATE INDEX idx_claims_pet_id ON claims(pet_id);

-- +goose Down
DROP TABLE claims;