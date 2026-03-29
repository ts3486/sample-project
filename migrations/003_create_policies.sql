-- +goose Up
CREATE TABLE policies (
    id                 SERIAL PRIMARY KEY,
    customer_id        INT NOT NULL REFERENCES customers(id),
    pet_id             INT NOT NULL REFERENCES pets(id),
    policy_number      VARCHAR(20) UNIQUE NOT NULL,
    status             VARCHAR(20) DEFAULT 'pending',
    plan_type          VARCHAR(20) NOT NULL,
    coverage_ratio     INT NOT NULL,
    monthly_premium    BIGINT NOT NULL,
    annual_premium     BIGINT NOT NULL,
    reimbursement_type VARCHAR(20) NOT NULL, -- 'counter' or 'post_visit'
    effective_date     DATE NOT NULL,
    expiration_date    DATE NOT NULL,
    injury_wait_end    DATE NOT NULL,
    illness_wait_end   DATE NOT NULL,
    cancer_wait_end    DATE NOT NULL,
    created_at         TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE policies;