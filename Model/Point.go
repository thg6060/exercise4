package Model

type Point struct{
	User_id string
	Points int64
	Max_points int64
}

////1. Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
func (p *Point) CreateTablePoint() error{
	
	defer db.Close()
	err = db.CreateTables(new(Point))
	err  = db.Sync2(new(Point))

	if err!=nil{
		return err
	}
	
	
	return nil
}
