package main

import (
	"fmt"
	//"strconv"
	"sync"
	"time"

	"./Model"
	//"github.com/rs/xid"
)

type DataUser struct {
	User Model.User
	Id   int
}

var conn, err = Model.DbConn()

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

func BirthtoTimeStamp(d int, m time.Month, y int) int64 {
	t := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	result := t.UnixNano()
	return result
}

func ReadFromDb(c chan DataUser, wg *sync.WaitGroup) error {

	rows, err := conn.Rows(&Model.User{})
	// SELECT * FROM user
	defer rows.Close()

	bean := new(Model.User)
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
	u := Model.User{
		Id:         "bs87var12b2tksrnsidg",
		Name:       "Giang xyz",
		Birth:      BirthtoTimeStamp(7, 2, 1999),
		Created:    time.Now().UnixNano(),
		Updated_at: time.Now().UnixNano(),
	}
	err := u.TransactionBirth("bs87vadsar12b2t34ksrnsidg", time.Now().UnixNano())
	fmt.Println(err)

}
