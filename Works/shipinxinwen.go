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

func Shipinxinwen_DoWork(rootUrl string,crawlerChan chan string,collectionName string){

	collector := colly.NewCollector()

	//爬取下一页链接
	collector.OnHTML("div#pages", func(e *colly.HTMLElement) {
		GetNextPage_Shipinxinwen(collector,e)
	})

	//爬取每页的公告url
	collector.OnHTML("ul.photo-list.picbig>li", func(e *colly.HTMLElement) {
		//log.Printf("爬取每页的公告url")
		GetDetailUrl_Shipinxinwen(e,collector, collectionName)
	})

	//获取详细信息
	collector.OnHTML("div#Article", func(e *colly.HTMLElement){
		//log.Printf("获取详细信息")
		GetDetailInfo_Shipinxinwen(e,collectionName)
	})

	collector.Visit(rootUrl)
	crawlerChan <- "视频新闻爬取结束"

}


//获取下一页链接
func GetNextPage_Shipinxinwen(collector *colly.Collector, element *colly.HTMLElement) {
	index := element.ChildText("span")
	indexInt,err  := strconv.Atoi(index)
	if err != nil{
		return
	}
	nextIndex := indexInt + 1
	url := "http://www.sqxw.gov.cn/html/shipin/shipinxinwen/"+strconv.Itoa(nextIndex)+".html"
	collector.Visit(url)
}

//爬取每页的公告url
func GetDetailUrl_Shipinxinwen( e *colly.HTMLElement, collector *colly.Collector,collectionName string) {

	url := e.ChildAttr("div.img-wrap>a","href")
	imgShow := e.ChildAttr("a>img","src")
	regThumb := regexp.MustCompile(`thumb_[0-9]+_[0-9]+_`)
	thumb := regThumb.FindString(imgShow)
	imgShow = strings.Replace(imgShow,thumb,"",-1)

	//如果数据库中已经有该数据
	if utils.CheckWebsite(collectionName,url) == false{
		info := models.Shipinxinwen{
			Website:url,
			ImgShow:imgShow,
		}
		session := utils.Session.Copy()
		collection := session.DB("ShiqianNews").C(collectionName)
		err := collection.Insert(info)
		if err != nil {
			log.Fatalf("[main.go] Work : 出现错误")
			return
		}
	}

	if len(url) > 0{
		collector.Visit(url)
	}

	if len(url) > 0 {
		collector.Visit(url)
	}
}
//获取详细信息
func GetDetailInfo_Shipinxinwen(e *colly.HTMLElement, collectionName string) {

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
	saveTitle := utils.TitleRg2.FindAllString(title,-1)
	title = ""
	for i := 0; i < len(saveTitle) ; i++  {
		title += saveTitle[i] + "";
	}

	//处理timer
	timer =  utils.TimerReg.FindString(timer)

	//获取cotent
	var contentList []string
	e.ForEach("div.content>p", func(i int, element *colly.HTMLElement) {
		p := element.Text
		span := element.ChildText("span")

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
		}
	})

	//获取edit
	edit := e.ChildText("p.text-right")
	edit = utils.ReplaceString(edit)
	saveEdit := utils.EditReg.FindAllString(edit,-1)
	edit = ""
	for i := 0; i < len(saveEdit) ; i++ {
		edit += saveEdit[i] + " "
	}
	edit = strings.Replace(edit,"聽","",-1)

	//获取视频url
	vedioUrl := e.ChildAttr("div.content>p>iframe","src")

	//如果没有imglist 则认为是脏数据，不要爬取
	if title == "" || timer == ""||len(contentList) <= 0 || vedioUrl == "" || len(vedioUrl) <= 0{
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
		"edit":edit,
		"vediourl":vedioUrl,
		"contents":contentList,
	}})

	if err != nil {
		//log.Fatalf("[main.go] Work : 出现错误 %s ",err)
		return
	}
}


