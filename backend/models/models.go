package models

type Plan struct {
	contents string `json:"contents"`
	deadline string `json:"deadline"`
	step     string `json:"step"`
}
