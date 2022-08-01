package model

import "time"

type Score struct {
	Id         uint64    `gorm:"column:id;autoIncrement;"`              // 自增编号
	Chinese    float64   `gorm:"type:float;not null;default:0;"`        // 语文
	CTeacherId string    `gorm:"type:varchar(64);not null;default:'';"` //教师编号
	Math       float64   `gorm:"type:float;not null;default:0;"`        // 数学
	MTeacherId string    `gorm:"type:varchar(64);not null;default:'';"` //教师编号
	English    float64   `gorm:"type:float;not null;default:0;"`        // 英语
	ETeacherId string    `gorm:"type:varchar(64);not null;default:'';"` //教师编号
	StuId      string    `gorm:"type:varchar(64);not null;default:'';"` // 用户编号
	Year       uint64    `gorm:"type:bigint;not null;default:0;"`       // 年份
	Term       uint64    `gorm:"type:bigint;not null;default:0;"`       // 学期
	CreatedAt  time.Time `gorm:"type:datetime(3);autoCreateTime;"`      // 创建时间
	UpdatedAt  time.Time `gorm:"type:datetime(3);autoUpdateTime;"`      // 更新时间
}

func (m *Score) TableName() string {
	return "t_score"
}
