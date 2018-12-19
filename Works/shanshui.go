package Works

import (
	"ShiqianCrawler/models"
	"ShiqianCrawler/utils"
	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2/bson"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Shanshui_DoWork(rootUrl string,crawlerChan chan string,collectionName string){

	collector := colly.NewCollector()

	//爬取下一页链接
	collector.OnHTML("div #pages", func(e *colly.HTMLElement) {
		GetNextPage_ShiqianShanshui(collector,e)
	})

	//爬取每页的公告url
	collector.OnHTML("ul.photo-list.picbig>li", func(e *colly.HTMLElement) {
		GetDetailUrl_ShiqianShanshui(e,collector, collectionName)
	})

	//获取详细信息
	collector.OnHTML("div#Article", func(e *colly.HTMLElement){
		GetDetailInfo_ShiqianShanshui(e,collectionName)
	})

	collector.Visit(rootUrl)
	crawlerChan <- "石阡山水爬取结束"


}


//获取下一页链接
func GetNextPage_ShiqianShanshui(collector *colly.Collector, element *colly.HTMLElement) {
	index := element.ChildText("span")
	indexInt,err  := strconv.Atoi(index)
	if err != nil{
		return
	}
	nextIndex := indexInt + 1
	url := "http://www.sqxw.gov.cn/html/tupian/ziranfengguang/"+strconv.Itoa(nextIndex)+".html"
	collector.Visit(url)
}

//爬取每页的公告url
func GetDetailUrl_ShiqianShanshui(element *colly.HTMLElement, collector *colly.Collector,collectionName string) {
	url := element.ChildAttr("div.img-wrap>a","href")
	if url == "" || len(url) <= 0 {
		return
	}

	if utils.CheckWebsite(collectionName,url) == false {
		info := models.Shiqianshanshui{
			Website:url,
		}
		session := utils.Session.Copy()
		defer session.Close()
		collection := session.DB("ShiqianNews").C(collectionName)
		err := collection.Insert(info)
		if err != nil {
			log.Printf("[Works->Shiqianshanshui] GetDetailUrl_ShiqianShanshui : 出现错误")
			return
		}
	}
	if len(url) > 0 {
		collector.Visit(url)
	}
}
//获取详细信息
func GetDetailInfo_ShiqianShanshui(e *colly.HTMLElement, collectionName string) {
	website := utils.GetWebsite(e.Request.URL)
	reslut := utils.CheckWebsite(collectionName,website)
	if reslut == false{
		log.Printf("数据不存在 : %s\n",website)
		return
	}

	title := e.ChildText("h1")
	timer := e.ChildText("h1>span")

	//处理title
	title = utils.ReplaceString(title)
	saveTitle := utils.TitleRg1.FindAllString(title,-1)
	title = ""
	for i := 0; i < len(saveTitle) ; i++  {
		title += saveTitle[i] + "";
	}

	//处理timer
	timer =  utils.TimerReg.FindString(timer)

	//获取内容图片url
	var imgList []string
	var imgShow string
	e.ForEach("div.tool>div.list-pic>div.cont>ul>li>div.img-wrap>a>img", func(i int, element *colly.HTMLElement) {
		imgUrl := element.Attr("src")
		regThumb := regexp.MustCompile(`thumb_[0-9]+_[0-9]+_`)
		thumb := regThumb.FindString(imgUrl)
		imgUrl = strings.Replace(imgUrl,thumb,"",-1)
		imgList = append(imgList,imgUrl)
		if imgShow == "" {
			imgShow = imgUrl
		}
	})

	//获取cotent
	var contentList []string
	e.ForEach("div.tool>div.content>p", func(i int, element *colly.HTMLElement) {
		p := element.Text
		p = utils.ReplaceString(p)
		contentList = append(contentList, p)
	})

	//如果没有imglist 则认为是脏数据，不要爬取
	if len(imgList) <= 0 {
		utils.RemoveDataByWebsite(collectionName,website)
		return
	}

	//进行数据的更新
	session := utils.Session.Copy()
	collection := session.DB("ShiqianNews").C(collectionName)
	err := collection.Update(bson.M{"website":website},bson.M{"$set":bson.M{
		"title":title,
		"timerid":timer,
		"imgshow":imgShow,
		"contents":contentList,
		"imglist":imgList,
	}})

	if err != nil {
		//log.Fatalf("[main.go] Work : 出现错误 %s ",err)
		return
	}
}

