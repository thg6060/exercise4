package Model

import "fmt"

type User struct {
	Id         string
	Name       string
	Birth      int64
	Created    int64
	Updated_at int64
}

//connect db
var db, err = DbConn()

//1.1/ Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
func (u *User) CreateTableUrs() error {

	defer db.Close()
	err = db.CreateTables(new(User))
	err = db.Sync2(new(User))
	if err != nil {
		return err
	}

	return nil
}
//1.2/ Viết hàm: insert và update user, viết hàm list user hoặc đọc user theo id(4 hàm)
func (u *User) Insert(urs *User) error {

	_, err = db.Insert(urs)

	if err != nil {
		return err
	}
	return nil

}

func (u *User) Update(urs *User) error {

	_, err = db.Table(urs).Where("id = ?",urs.Id).Update(urs)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) ShowList() error {
	users := make([]*User,10)
	err = db.Find(&users)

	for _, item := range users {
		fmt.Println(item)
	}

	if err != nil {
		return err
	}
	return nil

}
func (u *User) UserbyID(id string) (*User, error) {
	result:=User{}
	temp :=&result

	_, err = db.Table(temp).Where("id = ?", id).Get(temp)

	if err != nil {
		return nil, err
	}
	return temp, nil

}

func (u *User) InsertwithPnt(urs *User) error {
	//1.3/ Viết hàm: sau khi tạo user thì insert user_id vào user_point với số điểm 10.
	pnt := Point{
		Points:  10,
		User_id: urs.Id,
	}
	_, err = db.Insert(u, pnt)

	if err != nil {
		return err
	}
	return nil

}

func (u *User) TransactionBirth(id string,birth int64) error {
	/*
- b2: tạo 1 transaction khi update `birth` thành công thì cộng 10 điểm vào `point`
 sau đó sửa lại `name ` thành `$name + "updated "` 
nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
*/

	session := db.NewSession()
	defer session.Close()
	p:=Point{}
	us:=User{}

	// add Begin() before any action
	_,err:=session.Cols("points").Where("user_id = ?",id).Get(&p)
	_,err=session.Cols("Name").Where("id = ?",id).Get(&us)
	if err = session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}

	_,err = session.Where("id = ?", id).Update(&User{Birth:birth,Name:u.Name+" Update"})
	_,err =session.Cols("points").Where("user_id = ?",id).Update(&Point{Points: p.Points+10})
	err = session.Commit()

	if err != nil{
		session.Rollback()
		return err
	}
	return nil

}
