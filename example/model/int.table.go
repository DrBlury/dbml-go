package model

// Int is generated type for table 'int'
type Int struct {
	ID int `db:"id" json:"id" mapstructure:"id"`
}

// table 'int' columns list struct
type __tbl_int_columns struct {
	ID string
}

// table 'int' metadata struct
type __tbl_int struct {
	Name    string
	Columns __tbl_int_columns
}

// table 'int' metadata info
var _tbl_int = __tbl_int{
	Columns: __tbl_int_columns{ID: "id"},
	Name:    "int",
}

// GetColumns return list columns name for table 'int'
func (*__tbl_int) GetColumns() []string {
	return []string{"id"}
}

// T return metadata info for table 'int'
func (*Int) T() *__tbl_int {
	return &_tbl_int
}

// TableName return table name
func (Int) TableName() string {
	return "int"
}

/*
Generated by dbml-go
    version: v1.0.0
    timestamp: 2023-02-01 16:21:33.965040919 +0000 UTC m=+0.013178959
*/
