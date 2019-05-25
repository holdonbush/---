package dbase

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)


func init() {
	db, err := sql.Open("mysql","root:holdonbush@/test?charset=utf8")
	checkerr(err)
	_, err = db.Exec("CREATE TABLE `coininfo`(`name` VARCHAR(64) NOT NULL,`num` float DEFAULT NULL,`explain` varchar(100) DEFAULT NULL,`time` DATE NULL DEFAULT NULL);")
	if err!=nil {
		if err.Error() == "Error 1050: Table 'coininfo' already exists" {
			fmt.Println("success")
		} else {
			log.Fatal(err)
		}
	}
	db.Close();
	//_,err = stmt.Exec()
	//checkerr(err)
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(arg ...interface{})  {
	db, err := sql.Open("mysql","root:holdonbush@/test?charset=utf8")
	defer db.Close()
	checkerr(err)
	stmt, err := db.Prepare("INSERT into `coininfo` SET `name`= ?, `num` = ?, `explain` = ?, `time` = ?")
	checkerr(err)
	res, err := stmt.Exec(arg[0], arg[1], arg[2], arg[3])
	checkerr(err)
	fmt.Println(res)
}

func Select(arg ...interface{}) float64 {
	db, err := sql.Open("mysql","root:holdonbush@/test?charset=utf8")
	defer db.Close()
	checkerr(err)
	rows, err := db.Query("SELECT `num` from `coininfo` where `time` = ? and `name`=?",arg[0],arg[1])
	checkerr(err)
	nums, temp := 0.0, 0.0
	for rows.Next() {
		rows.Scan(&temp)
		nums = nums + temp
	}
	fmt.Println(nums)
	return nums
}