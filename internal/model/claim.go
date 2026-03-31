package model

import "time"

type Claim struct {
	ID               int           `json:"id" db:"id"`
	PolicyID         int           `json:"policy_id" db:"policy_id"`
	CustomerID       int           `json:"customer_id" db:"customer_id"`
	PetID            int           `json:"pet_id" db:"pet_id"`
	ClaimNumber      string        `json:"claim_number" db:"claim_number"` // e.g. "CLM-2026-00123"
	Category         ClaimCategory `json:"category" db:"category"`         // 通院, 入院, 手術
	Status           ClaimStatus   `json:"status" db:"status"`
	ClinicName       string        `json:"clinic_name" db:"clinic_name"`
	TreatmentDate    time.Time     `json:"treatment_date" db:"treatment_date"`
	Diagnosis        string        `json:"diagnosis" db:"diagnosis"` // 診断名
	TreatmentDetail  string        `json:"treatment_detail" db:"treatment_detail"`
	TotalAmount      int64         `json:"total_amount" db:"total_amount"`           // 治療費合計 (yen)
	CoveredAmount    int64         `json:"covered_amount" db:"covered_amount"`       // 補償対象額 (yen)
	ReimbursedAmount int64         `json:"reimbursed_amount" db:"reimbursed_amount"` // 保険金額 (yen)
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
}

type ClaimCategory string

const (
	ClaimCategoryOutpatient ClaimCategory = "outpatient" // 通院
	ClaimCategoryHospital   ClaimCategory = "hospital"   // 入院
	ClaimCategorySurgery    ClaimCategory = "surgery"    // 手術
)

type ClaimStatus string

const (
	ClaimStatusSubmitted   ClaimStatus = "submitted"    // 請求済み
	ClaimStatusUnderReview ClaimStatus = "under_review" // 審査中
	ClaimStatusApproved    ClaimStatus = "approved"     // 承認済み
	ClaimStatusDenied      ClaimStatus = "denied"       // 却下
	ClaimStatusPaid        ClaimStatus = "paid"         // 支払済み
	ClaimStatusClosed      ClaimStatus = "closed"       // 完了
)
