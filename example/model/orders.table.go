// Code generated by dbml-gen-go-model. DO NOT EDIT.
// Supported by duythinht@2020
package model

// Order is generated type for table 'orders'
type Order struct {
	ID     int    `db:"id" json:"id" mapstructure:"id"`
	UserID int    `db:"user_id" json:"user_id" mapstructure:"user_id"`
	Status string `db:"status" json:"status" mapstructure:"status"`
	// When order created
	CreatedAt string `db:"created_at" json:"created_at" mapstructure:"created_at"`
}

// table 'orders' columns list struct
type __tbl_orders_columns struct {
	ID        string
	UserID    string
	Status    string
	CreatedAt string
}

// table 'orders' metadata struct
type __tbl_orders struct {
	Name    string
	Columns __tbl_orders_columns
}

// table 'orders' metadata info
var _tbl_orders = __tbl_orders{
	Columns: __tbl_orders_columns{
		CreatedAt: "created_at",
		ID:        "id",
		Status:    "status",
		UserID:    "user_id",
	},
	Name: "orders",
}

// GetColumns return list columns name for table 'orders'
func (*__tbl_orders) GetColumns() []string {
	return []string{"id", "user_id", "status", "created_at"}
}

// T return metadata info for table 'orders'
func (*Order) T() *__tbl_orders {
	return &_tbl_orders
}
