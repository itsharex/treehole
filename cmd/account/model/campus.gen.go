// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCampu = "campus"

// Campu mapped from table <campus>
type Campu struct {
	ID         uint32  `gorm:"column:id;primaryKey" json:"id"`
	Name       string  `gorm:"column:name;not null" json:"name"`
	Code       string  `gorm:"column:code;not null" json:"code"`
	Department string  `gorm:"column:department;not null" json:"department"`
	Location   string  `gorm:"column:location;not null" json:"location"`
	Level      string  `gorm:"column:level;not null" json:"level"`
	Remark     *string `gorm:"column:remark" json:"remark"`
}

// TableName Campu's table name
func (*Campu) TableName() string {
	return TableNameCampu
}
