package Works

import (
	"ShiqianCrawler/models"
	"ShiqianCrawler/utils"
	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"strings"
)

func Wenquan_DoWork(rootUrl string,crawlerChan chan string,collectionName string){

	collector := colly.NewCollector()

	//爬取下一页链接
	collector.OnHTML("div#pages", func(e *colly.HTMLElement) {
		GetNextPage_Wenquan(collector,e)
	})

	//爬取每页的公告url
	collector.OnHTML("ul.list.lh24.f14>li", func(e *colly.HTMLElement) {
		//log.Printf("爬取每页的公告url")
		GetDetailUrl_Wenquan(e,collector, collectionName)
	})

	//获取详细信息
	collector.OnHTML("div#Article", func(e *colly.HTMLElement){
		//log.Printf("获取详细信息")
		GetDetailInfo_Wenquan(e,collectionName)
	})

	collector.Visit(rootUrl)
	crawlerChan <- "温泉文化爬取结束"

}


//获取下一页链接
func GetNextPage_Wenquan(collector *colly.Collector, element *colly.HTMLElement) {
	index := element.ChildText("span")
	indexInt,err  := strconv.Atoi(index)
	if err != nil{
		return
	}
	nextIndex := indexInt + 1
	url := "http://www.sqxw.gov.cn/html/wenhua/wenquanwenhua/"+strconv.Itoa(nextIndex)+".html"
	collector.Visit(url)
}

//爬取每页的公告url
func GetDetailUrl_Wenquan(element *colly.HTMLElement, collector *colly.Collector,collectionName string) {
	url := element.ChildAttr("a","href")
	if url == "" || len(url) <= 0 {
		return
	}

	if err := utils.InsertDetailUrl(collectionName, url,models.Wenquan{}); err != nil {
		log.Printf("%s",err)
	}

	if len(url) > 0 {
		collector.Visit(url)
	}
}
//获取详细信息
func GetDetailInfo_Wenquan(e *colly.HTMLElement, collectionName string) {
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

	//获取cotent
	imgShow := e.ChildAttr("div.content>img","src")
	var contentList []string
	e.ForEach("div.content>p", func(i int, element *colly.HTMLElement) {
		p := element.Text
		span := element.ChildText("span")

		pUrl := element.ChildAttr("img","src")
		var save string
		if p != "" {
			save = p
		}else if span != "" {
			save = span
		}

		if save != "" {
			save = utils.ReplaceString(save)
			save = strings.Replace(save,"聽","",-1)
			contentList = append(contentList, save)
		}else if pUrl != "" {
			if imgShow == "" {
				imgShow = pUrl
			}
			contentList = append(contentList, pUrl)
		}

	})

	//如果没有imglist 则认为是脏数据，不要爬取
	if title == "" || timer == ""||len(contentList) <= 0 || strings.Contains(title,`銆愬皬闃″湪鐜板満銆戠煶闃″幙鑱氬嚖涔`){

		utils.RemoveDataByWebsite(collectionName,website)
		return
	}

	//进行数据的更新
	session := utils.Session.Copy()
	defer session.Close()

	collection := session.DB("ShiqianNews").C(collectionName)
	err := collection.Update(bson.M{"website":website},bson.M{"$set":bson.M{
		"title":title,
		"timerid":timer,
		"imgshow":imgShow,
		"contents":contentList,
	}})

	if err != nil {
		//log.Fatalf("[main.go] Work : 出现错误 %s ",err)
		return
	}
}


