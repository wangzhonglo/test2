package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) (err error) {
	if err = db.AutoMigrate(&Student{}); err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}
	if err = db.AutoMigrate(&Grade{}); err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}
	if err = db.AutoMigrate(&Stu{}); err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}
	if err = db.AutoMigrate(&Score{}); err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}
	if err = db.AutoMigrate(&Teacher{}); err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}

	return nil
}
