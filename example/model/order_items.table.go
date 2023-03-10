package model

// OrderItem is generated type for table 'order_items'
type OrderItem struct {
	OrderID   int `db:"order_id" json:"order_id" mapstructure:"order_id"`
	ProductID int `db:"product_id" json:"product_id" mapstructure:"product_id"`
	Quantity  int `db:"quantity" json:"quantity" mapstructure:"quantity"`
}

// table 'order_items' columns list struct
type __tbl_order_items_columns struct {
	OrderID   string
	ProductID string
	Quantity  string
}

// table 'order_items' metadata struct
type __tbl_order_items struct {
	Name    string
	Columns __tbl_order_items_columns
}

// table 'order_items' metadata info
var _tbl_order_items = __tbl_order_items{
	Columns: __tbl_order_items_columns{
		OrderID:   "order_id",
		ProductID: "product_id",
		Quantity:  "quantity",
	},
	Name: "order_items",
}

// GetColumns return list columns name for table 'order_items'
func (*__tbl_order_items) GetColumns() []string {
	return []string{"order_id", "product_id", "quantity"}
}

// T return metadata info for table 'order_items'
func (*OrderItem) T() *__tbl_order_items {
	return &_tbl_order_items
}

// TableName return table name
func (OrderItem) TableName() string {
	return "order_items"
}

/*
Generated by dbml-go
    version: v1.0.0
    timestamp: 2023-02-10 14:57:53.354031668 +0000 UTC m=+0.012012001
*/
