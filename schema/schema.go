package schema

import "strconv"

const (
	GrowthRate string = "growth_rate"
)


type Schema struct {
	Tables []Table
	Settings map[string]string
}


func NewSchema() *Schema {
	var result Schema
	result.Settings[GrowthRate]	= strconv.Itoa(1024 * 1024) // 1 mb

	return &result
}
