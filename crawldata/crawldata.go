package crawldata

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	s "strings"
)

// 定义一个存储一条数据的结构体
type ImageData struct {
	Src    string
	Tp     string
	Title  string
	Width  int
	Height int
}

// 定义切片用于存储抓取的全部数据
type ImageDatas []ImageData

/*
   该函数将获取的数据存储到数据库
*/
func InsertData(datas *ImageDatas) {
	imageDatas := *datas
	// 连接数据库
	db, err := OpenDatabase()
	if err != nil {
		fmt.Printf(s.Join([]string{"连接数据库失败", err.Error()}, "-->"))
	}
	defer db.Close()

	for i := 0; i < len(imageDatas); i++ {
		imageData := imageDatas[i]
		// Prepare statement for inserting data
		imgIns, err := db.Prepare("INSERT INTO tableName (img_url, type_name, title, width, height) VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
		if err != nil {
			fmt.Println(s.Join([]string{"拼装数据格式", err.Error()}, "-->"))
		}
		defer imgIns.Close() // Close the statement when we leave main()

		img, err := imgIns.Exec(s.Join([]string{"http://www.gratisography.com", imageData.Src}, "/"), imageData.Tp, imageData.Title, imageData.Width, imageData.Height)
		if err != nil {
			fmt.Println(s.Join([]string{"插入数据失败", err.Error()}, "-->"))
		} else {
			success, _ := img.LastInsertId()
			// 数字变成字符串,success是int64型的值，需要转为int，网上说的Itoa64()在strconv包里不存在
			insertId := strconv.Itoa(int(success))
			fmt.Println(s.Join([]string{"成功插入数据：", insertId}, "\t-->\t"))
		}
	}
}

/*
   该函数用来抓取数据，并将存储的值返回到主函数
*/
func CrawlData(datas *ImageDatas) (imageDatas ImageDatas) {
	imageDatas = *datas
	// 规定抓取时匹配的元素
	var types = [...]string{
		"people",
		"objects",
		"whimsical",
		"nature",
		"urban",
		"animals"}

	doc, err := goquery.NewDocument("http://www.gratisography.com/")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, tp := range types {
		doc.Find("#container ul").Find(s.Join([]string{".", tp}, "")).Each(func(i int, s *goquery.Selection) {
			img := s.Find("img.lazy")
			src, _ := img.Attr("data-original")
			title, _ := img.Attr("alt")
			width, _ := img.Attr("width")
			height, _ := img.Attr("height")

			// 将宽度和高度的字符串类型转为数值型
			wd, error := strconv.Atoi(width)
			if error != nil {
				fmt.Println("字符串转换成整数失败")
			}
			hg, error := strconv.Atoi(height)
			if error != nil {
				fmt.Println("字符串转换成整数失败")
			}
			// fmt.Printf("Review %d: %s - %s - %s - %d - %d\n", i, src, tp, title, wd, hg)
			imageData := ImageData{src, tp, title, wd, hg}
			imageDatas = append(imageDatas, imageData)
		})
	}
	return
}

func Crawl() {
	// 定义一个切片存储所有数据
	var datas ImageDatas
	// 抓取数据
	imageDatas := CrawlData(&datas)
	InsertData(&imageDatas)

}
