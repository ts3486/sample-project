package model

import "time"

type PlanType string

const (
	PlanNext  PlanType = "next"  // ネクスト: 通院+入院+手術 (no daily limits)
	PlanLight PlanType = "light" // ライト: 通院+入院+手術 (with daily limits)
	PlanMini  PlanType = "mini"  // ミニ: 手術のみ (internet-only)
	PlanVIP   PlanType = "vip"   // VIP: 通院+入院+手術 (¥5M annual cap)
)

// DailyLimits represents per-day and per-use limits (日額制限)
type DailyLimits struct {
	OutpatientDailyLimit  int64 `json:"outpatient_daily_limit"`   // 通院1日あたり上限 (yen)
	OutpatientDaysPerYear int   `json:"outpatient_days_per_year"` // 通院年間日数上限
	HospitalDailyLimit    int64 `json:"hospital_daily_limit"`     // 入院1日あたり上限 (yen)
	HospitalDaysPerYear   int   `json:"hospital_days_per_year"`   // 入院年間日数上限
	SurgeryPerCaseLimit   int64 `json:"surgery_per_case_limit"`   // 手術1回あたり上限 (yen)
	SurgeryCasesPerYear   int   `json:"surgery_cases_per_year"`   // 手術年間回数上限
}

type Policy struct {
	ID                int               `json:"id" db:"id"`
	CustomerID        int               `json:"customer_id" db:"customer_id"`
	PetID             int               `json:"pet_id" db:"pet_id"`
	PolicyNumber      string            `json:"policy_number" db:"policy_number"` // e.g. "PET-2026-00042"
	Status            PolicyStatus      `json:"status" db:"status"`
	PlanType          PlanType          `json:"plan_type" db:"plan_type"`
	CoverageRatio     int               `json:"coverage_ratio" db:"coverage_ratio"`
	MonthlyPremium    int64             `json:"monthly_premium" db:"monthly_premium"` // yen
	AnnualPremium     int64             `json:"annual_premium" db:"annual_premium"`   // yen
	ReimbursementType ReimbursementType `json:"reimbursement_type" db:"reimbursement_type"`
	EffectiveDate     time.Time         `json:"effective_date" db:"effective_date"`
	ExpirationDate    time.Time         `json:"expiration_date" db:"expiration_date"`
	InjuryWaitEnd     time.Time         `json:"injury_wait_end" db:"injury_wait_end"`   // ケガ待機終了
	IllnessWaitEnd    time.Time         `json:"illness_wait_end" db:"illness_wait_end"` // 病気待機終了
	CancerWaitEnd     time.Time         `json:"cancer_wait_end" db:"cancer_wait_end"`   // がん待機終了
	CreatedAt         time.Time         `json:"created_at" db:"created_at"`
}

type PolicyStatus string

const (
	PolicyStatusPending   PolicyStatus = "pending"   // 申込中
	PolicyStatusWaiting   PolicyStatus = "waiting"   // 待機期間中
	PolicyStatusActive    PolicyStatus = "active"    // 有効
	PolicyStatusCancelled PolicyStatus = "cancelled" // 解約済み
	PolicyStatusExpired   PolicyStatus = "expired"   // 満了
	PolicyStatusSuspended PolicyStatus = "suspended" // 停止中 (未払い)
)

type ReimbursementType string

const (
	ReimbursementCounter   ReimbursementType = "counter"    // 窓口精算
	ReimbursementPostVisit ReimbursementType = "post_visit" // 後日精算
)
