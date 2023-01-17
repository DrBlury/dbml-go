// Code generated by dbml-gen-go-model. DO NOT EDIT.
// Supported by duythinht@2020
package model

// Merchant is generated type for table 'merchants'
type Merchant struct {
	ID           int    `db:"id" json:"id" mapstructure:"id"`
	MerchantName string `db:"merchant_name" json:"merchant_name" mapstructure:"merchant_name"`
	CountryCode  int    `db:"country_code" json:"country_code" mapstructure:"country_code"`
	CreatedAt    string `db:"created_at" json:"created_at" mapstructure:"created_at"`
	AdminID      int    `db:"admin_id" json:"admin_id" mapstructure:"admin_id"`
}

// table 'merchants' columns list struct
type __tbl_merchants_columns struct {
	ID           string
	MerchantName string
	CountryCode  string
	CreatedAt    string
	AdminID      string
}

// table 'merchants' metadata struct
type __tbl_merchants struct {
	Name    string
	Columns __tbl_merchants_columns
}

// table 'merchants' metadata info
var _tbl_merchants = __tbl_merchants{
	Columns: __tbl_merchants_columns{
		AdminID:      "admin_id",
		CountryCode:  "country_code",
		CreatedAt:    "created_at",
		ID:           "id",
		MerchantName: "merchant_name",
	},
	Name: "merchants",
}

// GetColumns return list columns name for table 'merchants'
func (*__tbl_merchants) GetColumns() []string {
	return []string{"id", "merchant_name", "country_code", "created_at", "admin_id"}
}

// T return metadata info for table 'merchants'
func (*Merchant) T() *__tbl_merchants {
	return &_tbl_merchants
}