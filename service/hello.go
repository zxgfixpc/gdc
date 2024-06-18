package service

import (
	"context"
	"fmt"

	"_gdc_/dao"
	"_gdc_/lib/log"
)

type (
	HelloRsp struct {
		Items []Item `json:"items"`
	}
	Item struct {
		Desc string `json:"desc"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func Hello(ctx context.Context) (*HelloRsp, error) {
	log.Info(ctx, "hello")
	log.Error(ctx, "test hello error ")

	// 初始化3个db，一个写，2个读，每个库初始化1条记录，查询和更新，判断是否实现读写分离
	ret := &HelloRsp{}
	items := make([]Item, 0)

	// 连续查询10次数据库，判断是否读写分离
	for i := 0; i < 10; i++ {
		studentInfo, err := dao.GetStudentByID(ctx, 1)
		if err != nil {
			return nil, err
		}

		items = append(items, Item{
			Desc: fmt.Sprintf("初始查询第%v次", i),
			ID:   studentInfo.ID,
			Name: studentInfo.Name,
			Age:  studentInfo.Age,
		})
	}

	// 更新主库
	err := dao.UpdateAgeByID(ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	// 连续查询10次数据库，判断是否读写分离
	for i := 0; i < 10; i++ {
		studentInfo, err := dao.GetStudentByID(ctx, 1)
		if err != nil {
			return nil, err
		}

		items = append(items, Item{
			Desc: fmt.Sprintf("更新后查询第%v次", i),
			ID:   studentInfo.ID,
			Name: studentInfo.Name,
			Age:  studentInfo.Age,
		})
	}

	ret.Items = items
	return ret, nil
}
