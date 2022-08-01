package model

import "time"

type Teacher struct {
	Id        uint64    `gorm:"column:id;autoIncrement;"`              // 自增编号
	Name      string    `gorm:"type:varchar(64);not null;default:'';"` // 名称
	TeacherId string    `gorm:"type:varchar(64);not null;default:'';"` //教师编号
	Password  string    `gorm:"type:varchar(64);not null;default:'';"`
	Step      string    `gorm:"type:varchar(64);not null;default:'';"` //所在系
	Course    string    `gorm:"type:varchar(64);not null;default:'';"` //所授课程
	CreatedAt time.Time `gorm:"type:datetime(3);autoCreateTime;"`      // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime(3);autoUpdateTime;"`      // 更新时间

}

func (m *Teacher) TableName() string {
	return "t_teacher"
}
