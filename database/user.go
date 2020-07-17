package Model

import (
	"errors"
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

	eff, err := db.Insert(urs)
	if eff == 0 {
		return errors.New("Insert failed")
	}
	if err != nil {
		return err
	}
	return nil

}

func (u *User) Update(urs *User, condition *User) error {

	eff, err := db.Update(urs, condition)
	if eff == 0 {
		return errors.New("Update failed")
	}

	if err != nil {
		return err
	}
	return nil
}

func (u *User) ShowList() ([]*User, error) {
	var users []*User
	err := db.Find(&users)

	if err != nil {
		return nil, err
	}
	return users, nil

}

func (u *User) UserbyID(id string) (*User, error) {
	result := User{}
	eff, err := db.Where("id = ?", id).Get(&result)

	if eff == false {
		return nil, errors.New("not found user by id ")
	}
	if err != nil {
		return nil, err
	}

	return &result, nil

}
