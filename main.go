package main

import (
	"errors"
	"fmt"

	//"strconv"
	"sync"
	"time"

	"github.com/thg6060/exercise4/Database"
	//"github.com/rs/xid"
)

type DataUser struct {
	User Database.User
	Id   int
}

var conn, err = Database.DbConn()

func worker(done []chan bool, c chan DataUser, wg *sync.WaitGroup, id int) {
	var mu sync.Mutex
loop:
	for {
		mu.Lock()
		select {

		case result := <-c:
			fmt.Printf("%d- %s - %s \n", result.Id, result.User.Id, result.User.Name)
			wg.Done()
			mu.Unlock()
		case ok := <-done[id]:

			if ok {
				break loop
			}
		}

	}
}

func InsertwithPoint(urs *Database.User) error {
	//1.3/ Viết hàm: sau khi tạo user thì insert user_id vào user_point với số điểm 10.
	p := Database.Point{
		UserId: urs.Id,
		Points: 10,
	}
	err := p.Insert(&p)
	if err != nil {
		return err
	}
	_, err = conn.Insert(urs)

	if err != nil {
		return err
	}
	return nil

}

func BirthtoTimeStamp(d int, m time.Month, y int) int64 {
	t := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	result := t.UnixNano()
	return result
}

func ReadFromDb(c chan DataUser, wg *sync.WaitGroup) error {

	rows, err := conn.Rows(&Database.User{})
	// SELECT * FROM user
	defer rows.Close()

	bean := new(Database.User)
	i := 0
	for rows.Next() {
		i++
		wg.Add(1)
		err = rows.Scan(bean)
		dt := DataUser{User: *bean, Id: i}
		c <- dt
	}

	if err != nil {
		return err
	}
	return nil
}

func TransactionBirth(id string, birth int64) error {
	/*
		- b2: tạo 1 transaction khi update `birth` thành công thì cộng 10 điểm vào `point`
		 sau đó sửa lại `name ` thành `$name + "updated "`
		nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
	*/

	session := conn.NewSession()
	defer session.Close()
	p := Database.Point{}
	us := Database.User{}
	//
	eff1, err := session.Cols("points").Where("user_id = ?", id).Get(&p)
	if !eff1 {
		session.Rollback()
		return errors.New("Get Point with Point field failed")
	}
	if err != nil {
		session.Rollback()
		return err
	}
	//
	eff2, err := session.Cols("Name").Where("id = ?", id).Get(&us)
	if !eff2 {
		session.Rollback()
		return errors.New("Get User with Name field failed")

	}
	if err != nil {
		session.Rollback()
		return err
	}

	//
	eff3, err := session.Where("id = ?", id).Update(&Database.User{Birth: birth, Name: us.Name + " Update"})
	if eff3 == 0 {
		session.Rollback()
		return errors.New("Get User with Name field failed")

	}
	if err != nil {
		session.Rollback()
		return err
	}

	//

	eff4, err := session.Cols("points").Where("user_id = ?", id).Update(&Database.Point{Points: p.Points + 10})
	if eff4 == 0 {
		session.Rollback()
		return errors.New("Get User with Name field failed")

	}
	if err != nil {
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

func main() {
	/*
	   - b3: insert 100 bản ghi vào user:
	   sau đó viết 1 workerpool scantableuser lấy ra tên của các user inra màn hình
	   (Dùng scan theo row)
	   dùng 2 worker
	   và thiết lập bộ đếm `${counter} - ${id} - ${name}`

	   ```go
	   defer rows.Close()
	   bean := new(Struct)
	   for rows.Next() {
	       err = rows.Scan(bean)
	   }
	*/

	/*
		//Them 100 ban ghi
			for i := 0; i < 100; i++ {
				guid := xid.New()
				s := strconv.Itoa(i)
				u := Model.User{
					Id:         guid.String(),
					Name:       "Giang " + s,
					Birth:      BirthtoTimeStamp(7, 2, 1999),
					Created:    time.Now().UnixNano(),
					Updated_at: time.Now().UnixNano(),
				}
				u.Insert(&u)
			}
	*/

	/*
		numofGoroutine := 2
		var wg sync.WaitGroup
		c := make(chan DataUser)
		done := make([]chan bool, numofGoroutine)

		defer conn.Close()

		for j := 0; j < numofGoroutine; j++ {
			go worker(done, c, &wg, j)
		}

		ReadFromDb(c, &wg)

		for k := range done {
			done[k] = make(chan bool, 1)
			done[k] <- true
		}

		wg.Wait()

		fmt.Println("done !")
	*/
	//guid := xid.New()
	//u := Model.User{}
	//data,err := u.ShowList()
	u := Database.User{
		Id:        "CucCang",
		Name:      "Giang ",
		Birth:     BirthtoTimeStamp(7, 2, 1999),
		Created:   time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	data,err:=u.ShowList()
	for _,item:=range data{
		fmt.Println(item)
	}

	if err != nil {
		fmt.Println(err)
	}

}
