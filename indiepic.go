package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"indiepic/crawldata"
	s "strings"
)

type Results struct {
	Err   int
	Msg   string
	Datas crawldata.ImageDatas
}

func main() {
	// 使用crawldata包里面的Crawl()抓取需要的数据存到数据库
	// crawldata.Crawl()

	m := martini.Classic()
	m.Use(render.Renderer())

	var (
		results Results
		err     error
	)

	m.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"Err": "10001", "Msg": "Not Found"})
	})
	// This will set the Content-Type header to "application/json; charset=UTF-8"
	m.Get("/api", func(r render.Render) {
		results.Datas, err = crawldata.GetAllImages()
		if err != nil {
			fmt.Println(s.Join([]string{"获取数据失败", err.Error()}, "-->"))
			r.JSON(200, map[string]interface{}{"Err": "10001", "Msg": "Data Error"})
		} else {
			results.Err = 10001
			results.Msg = "获取数据成功"
			if res, err := json.Marshal(results); err == nil {
				r.JSON(200, string(res))
			} else {
				r.JSON(200, map[string]interface{}{"Err": "10001", "Msg": "Data Error"})
			}
		}
	})
	m.Run()

}
