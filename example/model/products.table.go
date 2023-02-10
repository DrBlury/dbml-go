package model

import "time"

// Product is generated type for table 'products'
type Product struct {
	ID         int            `db:"id" json:"id" mapstructure:"id"`
	Name       string         `db:"name" json:"name" mapstructure:"name"`
	MerchantID int            `db:"merchant_id" json:"merchant_id" mapstructure:"merchant_id"`
	Price      int            `db:"price" json:"price" mapstructure:"price"`
	Status     ProductsStatus `db:"status" json:"status" mapstructure:"status"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at" mapstructure:"created_at"`
}

// table 'products' columns list struct
type __tbl_products_columns struct {
	ID         string
	Name       string
	MerchantID string
	Price      string
	Status     string
	CreatedAt  string
}

// table 'products' metadata struct
type __tbl_products struct {
	Name    string
	Columns __tbl_products_columns
}

// table 'products' metadata info
var _tbl_products = __tbl_products{
	Columns: __tbl_products_columns{
		CreatedAt:  "created_at",
		ID:         "id",
		MerchantID: "merchant_id",
		Name:       "name",
		Price:      "price",
		Status:     "status",
	},
	Name: "products",
}

// GetColumns return list columns name for table 'products'
func (*__tbl_products) GetColumns() []string {
	return []string{"id", "name", "merchant_id", "price", "status", "created_at"}
}

// T return metadata info for table 'products'
func (*Product) T() *__tbl_products {
	return &_tbl_products
}

// TableName return table name
func (Product) TableName() string {
	return "products"
}

/*
Generated by dbml-go
    version: v1.0.0
    timestamp: 2023-02-10 14:57:53.358122001 +0000 UTC m=+0.016102334
*/
