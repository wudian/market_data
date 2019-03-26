package main

type Ticker struct {
	Symbol string 	`json:"symbol,string"`
	Last float64      `json:"last,string"`
	Buy  float64      `json:"buy,string"`
	Sell float64      `json:"sell,string"`
	High float64      `json:"high,string"`
	Low  float64      `json:"low,string"`
	Vol  float64      `json:"vol,string"`
	Date int64       `json:"date"` // 单位:秒(second)
}