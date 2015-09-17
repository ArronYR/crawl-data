package crawldata

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	s "strings"
)

func OpenDatabase() (*sql.DB, error) {
	// 连接数据库
	db, err := sql.Open("mysql", "root:mysql@tcp(xxx.xxx.xxx.xxx:3306)/indiepic?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetAllImages() (imageDatas ImageDatas, err error) {
	// 连接数据库
	db, err := OpenDatabase()
	if err != nil {
		fmt.Printf(s.Join([]string{"连接数据库失败", err.Error()}, "-->"))
		return nil, err
	}
	defer db.Close()

	// Prepare statement for inserting data
	imgOut, err := db.Query("SELECT * FROM gratisography")
	if err != nil {
		fmt.Println(s.Join([]string{"获取数据失败", err.Error()}, "-->"))
		return nil, err
	}
	defer imgOut.Close()

	var (
		id          int
		img_url     string
		type_name   string
		title       string
		width       int
		height      int
		create_time string
	)
	for imgOut.Next() {
		err := imgOut.Scan(&id, &img_url, &type_name, &title, &width, &height, &create_time)
		if err != nil {
			fmt.Println(s.Join([]string{"查询数据失败", err.Error()}, "-->"))
			return nil, err
		} else {
			imageData := ImageData{img_url, type_name, title, width, height}
			imageDatas = append(imageDatas, imageData)
		}
	}

	return imageDatas, nil
}
