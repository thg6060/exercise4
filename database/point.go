package Model

type Point struct{
	User_id string
	Points int64
	Max_points int64
}

////1. Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
