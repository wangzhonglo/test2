package model

import "time"

type Student struct {
	Id        uint64    `gorm:"column:id;autoIncrement;"`              // 自增编号
	Name      string    `gorm:"type:varchar(64);not null;default:'';"` // 名称
	Age       uint64    `gorm:"type:bigint;not null;default:0;"`       // 年龄
	Gender    uint64    `gorm:"type:bigint;not null;default:0;"`       // 性别
	CreatedAt time.Time `gorm:"type:datetime(3);autoCreateTime;"`      // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime(3);autoUpdateTime;"`      // 更新时间
}

func (m *Student) TableName() string {
	return "t_student"
}
