package model

import (
	"time"
)

type Stu struct {
	Id        uint64    `gorm:"column:id;autoIncrement;"`              // 自增编号
	Name      string    `gorm:"type:varchar(64);not null;default:'';"` // 名称
	StuId     string    `gorm:"type:varchar(64);not null;default:'';"`
	Age       uint64    `gorm:"type:bigint;not null;default:0;"`       // 年龄
	Sex       uint64    `gorm:"type:bigint;not null;default:0;"`       // 性别
	Grade     uint64    `gorm:"type:bigint;not null;default:0;"`       //年级
	Step      string    `gorm:"type:varchar(64);not null;default:'';"` //所在系
	CreatedAt time.Time `gorm:"type:datetime(3);autoCreateTime;"`      // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime(3);autoUpdateTime;"`      // 更新时间
	Password  string    `gorm:"type:varchar(64);not null;default:'';"`
}

func (m *Stu) TableName() string {
	return "t_stu"
}
