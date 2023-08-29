// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePayment = "payment"

// Payment mapped from table <payment>
type Payment struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null;comment:名称" json:"name"`                      // 名称
	Channel    string    `gorm:"column:channel;not null;comment:标识 支付宝 微信 易支付 码支付" json:"channel"` // 标识 支付宝 微信 易支付 码支付
	Types      string    `gorm:"column:types;not null;comment:['ailipay','wxpay']" json:"types"`   // ['ailipay','wxpay']
	Params     string    `gorm:"column:params;comment:支付所需参数" json:"params"`                       // 支付所需参数
	Status     int32     `gorm:"column:status;not null;default:1;comment:1 正常 0异常" json:"status"`  // 1 正常 0异常
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Payment's table name
func (*Payment) TableName() string {
	return TableNamePayment
}