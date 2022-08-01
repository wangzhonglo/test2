package svc

import (
	"greet/internal/config"
	"greet/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(conf config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(conf.MysqlDataSource), &gorm.Config{})
	if err != nil {
		logx.Errorf("connect mysql err: %v", err)
	}
	if conf.Debug {
		db = db.Debug()
	}

	if err = model.InitTable(db); err != nil {
		logx.Errorf("init table err: %v", err)
		panic(err)
	}
	return &ServiceContext{
		Config: conf,
		DB:     db,
	}
}
