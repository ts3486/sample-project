package model

type AppliedDiscount struct {
	Type   DiscountType `json:"type"`
	Name   string       `json:"name"`
	Amount int64        `json:"amount"` // yen saved per year
}

type DiscountType string

const (
	DiscountMultiPet     DiscountType = "multi_pet"    // 多頭割引
	DiscountMicrochip    DiscountType = "microchip"    // マイクロチップ割引
	DiscountInternet     DiscountType = "internet"     // インターネット割引
	DiscountAnnualPay    DiscountType = "annual_pay"   // 年払い割引
	DiscountNeutered     DiscountType = "neutered"     // 避妊・去勢割引
	DiscountContinuation DiscountType = "continuation" // 継続割引 (renewal)
	DiscountNoClaim      DiscountType = "no_claim"     // 無事故割引
)
