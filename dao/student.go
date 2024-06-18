package dao

import (
	"context"

	"_gdc_/lib/infra"
)

type Student struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Age  int    `json:"age" gorm:"column:age"`
}

func (Student) TableName() string {
	return "student"
}

func GetStudentByID(ctx context.Context, id int) (result *Student, err error) {
	result = &Student{}
	err = infra.MysqlClient.WithContext(ctx).Model(&Student{}).Where("id = ?", id).Find(result).Error
	return
}

func UpdateAgeByID(ctx context.Context, id int, age int) error {
	return infra.MysqlClient.WithContext(ctx).Model(&Student{}).Where("id = ?", id).
		Updates(map[string]interface{}{"age": age}).Error
}
