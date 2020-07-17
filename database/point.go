package Model

import "fmt"

type Point struct {
	User_id    string
	Points     int64
	Max_points int64
}

func (p *Point) Insert(id string, pt int64) error {
	pnt := Point{
		Points:  pt,
		User_id: id,
	}
	eff, err := db.Insert(pnt)
	if eff == 0 {
		fmt.Println("update point failed")
	}
	if err != nil {
		return err
	}
	return nil
}

////1. Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
