// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Role string `gorm:"column:role" json:"role"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
