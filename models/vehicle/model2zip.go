package models

type ModelToZip struct {
	ZipCodeId  int	`json:"zip_code_id" db:"zip_code_id"`
	ModelId    int	`json:"model_id" db:"model_id"`
	TotalCount int	`json:"total_count" db:"id"`
}
