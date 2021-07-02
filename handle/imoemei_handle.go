package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.spider/tools"
	"io"
	"log"
)

// IndexHandle 主页
type IndexHandle struct {}
type DeepHandle struct {}
type CategoryHandle struct {}

func (i *IndexHandle) Worker(body io.Reader, _ string)  {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".b2_gap .post-list-item").Each(func(i int, selection *goquery.Selection) {
		href, ifExists  := selection.Find("a").Attr("href")
		if ifExists {
			// 进一步挖掘图片
			var req = new(tools.Request)
			var deep = new(DeepHandle)
			request, err := req.NewRequest("GET", href, "", deep, nil)
			if err != nil {
				log.Fatal(err)
			}
			if err := request.Execute(); err != nil{
				log.Fatal(err)
			}
			//var ch = make(chan int)
			//ch<- i
			//go func() {
			//	for  {
			//		select {
			//		case _, ok := <-ch:
			//			if !ok {
			//				return
			//			}
			//			request.Execute()
			//		}
			//	}
			//}()
		}
	})
}

func (i *DeepHandle) Worker(body io.Reader, _ string)  {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		doc.Find(".entry-content p").Each(func(i int, selection *goquery.Selection) {
			src, ifExists  := selection.Find("img").Attr("src")
			if ifExists {
				fmt.Println(src)
				tools.SavePic(src)
			}
		})
	}()
}

func (receiver *CategoryHandle) Worker(body io.Reader, _ string)  {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".b2_gap .post-list-item").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i)
		href, ifExists  := selection.Find(" a").Attr("href")
		if ifExists {
			fmt.Println(href)
		}
	})
}
