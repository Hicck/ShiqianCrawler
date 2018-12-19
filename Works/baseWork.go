package Works

type CrawlerWork struct {
	Url string
	CollectionName string
	ProjectNameChan chan string
	Work func(rootUrl string,crawlerChan chan string,collectionName string)
}
func(this *CrawlerWork) DoWork(){
	this.Work(this.Url,this.ProjectNameChan,this.CollectionName)
}


