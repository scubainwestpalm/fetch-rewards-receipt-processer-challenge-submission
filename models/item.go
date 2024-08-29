package models

type Item struct {
	ShortDescription string  `json:"shortDescription" binding:"required"`
	Price            float64 `json:"price,string" binding:"required"`
}
