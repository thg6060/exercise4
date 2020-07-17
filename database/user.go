package Model

import (
	"log"
)

type User struct {
	Id         string
	Name       string
	Birth      int64
	Created    int64
	Updated_at int64
}

//connect db


//1.1/ Viết hàm: Chỉ tạo db, và tạo model(struct) ánh xạ struct thành table (CreateTable, Sync2)
//1.2/ Viết hàm: insert và update user, viết hàm list user hoặc đọc user theo id(4 hàm)
func (u *User) Insert(urs *User) error {

	_, err = db.Insert(urs)

	if err != nil {
		return err
	}
	return nil

}

func (u *User) Update(urs *User,condition *User) error {

	_, err = db.Update(urs,condition)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) ShowList() ([]*User,error){
	var users []*User
	err := db.Find(&users)

	if err != nil {
		return nil,err
	}
	return users,nil

}

func (u *User) UserbyID(id string) (*User, error) {
	result := User{}
	temp := &result

	eff, err := db.Table(temp).Where("id = ?", id).Get(temp)
	if eff == false{
		log.Println("cannot find user by id")
		return nil,err
	}
	return temp, nil

}


func (u *User) TransactionBirth(id string, birth int64) error {
	/*
		- b2: tạo 1 transaction khi update `birth` thành công thì cộng 10 điểm vào `point`
		 sau đó sửa lại `name ` thành `$name + "updated "`
		nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
	*/

	session := db.NewSession()
	defer session.Close()
	p := Point{}
	us := User{}

	
	eff1, err := session.Cols("points").Where("user_id = ?", id).Get(&p)
	if eff1 == false {
		log.Println("Get Point with Point field faild")
		session.Rollback()
		return err
	}

	eff2, err := session.Cols("Name").Where("id = ?", id).Get(&us)
	if eff2 == false {
		log.Println("Get User with Name field failed")
		session.Rollback()
		return err
	}

	eff3, err := session.Where("id = ?", id).Update(&User{Birth: birth, Name: u.Name + " Update"})
	if eff3 == 0 {
		log.Println("Update birth user failed")
		session.Rollback()
		return err
	}

	eff4, err := session.Cols("points").Where("user_id = ?", id).Update(&Point{Points: p.Points + 10})
	if eff4 == 0 {
		log.Println("Update point(+10) failed")
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		session.Rollback()
		return err
	}
	return nil

}
