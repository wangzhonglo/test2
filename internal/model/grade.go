package model

import "time"

type Grade struct {
	Id        uint64    `gorm:"column:id;autoIncrement;"`         // 自增编号
	Chinese   float64   `gorm:"type:float;not null;default:0;"`   // 语文
	Math      float64   `gorm:"type:float;not null;default:0;"`   // 数学
	English   float64   `gorm:"type:float;not null;default:0;"`   // 英语
	UserId    uint64    `gorm:"type:bigint;not null;default:0;"`  // 用户编号
	CreatedAt time.Time `gorm:"type:datetime(3);autoCreateTime;"` // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime(3);autoUpdateTime;"` // 更新时间
	Year      uint64    `gorm:"type:bigint;not null;default:0;"`  // 年份
	Term      uint64    `gorm:"type:bigint;not null;default:0;"`  // 学期
}

func (m *Grade) TableName() string {
	return "t_grade"
}
