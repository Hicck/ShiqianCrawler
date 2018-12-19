package main

import (
	"ShiqianCrawler/Works"
	_"ShiqianCrawler/utils"
	"log"
)

func main() {

	projectNameChan := make(chan string)
	crawlerWorks := [...]Works.CrawlerWork{
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/tupian/ziranfengguang/index.html",
			CollectionName:"shanshui",
			ProjectNameChan:projectNameChan,
			Work:Works.Shanshui_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/index.php?m=content&c=index&a=lists&catid=211",
			CollectionName:"snack",
			ProjectNameChan:projectNameChan,
			Work:Works.Snack_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/wenhua/minsuzongjiaowenhua/",
			CollectionName:"minsuculture",
			ProjectNameChan:projectNameChan,
			Work:Works.Minsu_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/wenhua/taichawenhua/",
			CollectionName:"taitea",
			ProjectNameChan:projectNameChan,
			Work:Works.Taitea_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/wenhua/wenquanwenhua/",
			CollectionName:"wenquanculture",
			ProjectNameChan:projectNameChan,
			Work:Works.Wenquan_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/wenhua/yelangwenhua/",
			CollectionName:"yelangculture",
			ProjectNameChan:projectNameChan,
			Work:Works.Yelang_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/wenhua/hongsewenhua/index.html",
			CollectionName:"redculture",
			ProjectNameChan:projectNameChan,
			Work:Works.RedCulture_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/xinwen/shiqianyaowen/index.html",
			CollectionName:"yaowen",
			ProjectNameChan:projectNameChan,
			Work:Works.Yaowen_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/shipin/shipinxinwen/",
			CollectionName:"shipinixinwen",
			ProjectNameChan:projectNameChan,
			Work:Works.Shipinxinwen_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/shipin/shihuashishuo/",
			CollectionName:"shihuashishuo",
			ProjectNameChan:projectNameChan,
			Work:Works.Shihuashishuo_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/xiangzhen/",
			CollectionName:"xiangzhenzixun",
			ProjectNameChan:projectNameChan,
			Work:Works.Xiangzhenzixun_DoWork,
		},
		Works.CrawlerWork{
			Url:"http://www.sqxw.gov.cn/html/bumen/",
			CollectionName:"bumendongtai",
			ProjectNameChan:projectNameChan,
			Work:Works.Bumendongtai_DoWork,
		},
	}
	for i := 0; i < len(crawlerWorks); i++ {
		go crawlerWorks[i].DoWork()
	}

	for i := 0; i < len(crawlerWorks); i++ {
		log.Printf("%s\n",<-projectNameChan)
	}

}
