package Model

import "errors"

type Point struct {
	User_id    string
	Points     int64
	Max_points int64
}

func (p *Point) Insert(pnt *Point) error {
	eff, err := db.Insert(pnt)
	if eff == 0 {
		return errors.New("Insert point failed")
	}
	if err != nil {
		return err
	}
	return nil
}

////1. Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
