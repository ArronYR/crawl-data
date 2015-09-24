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
