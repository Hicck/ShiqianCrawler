package models

type Shihuashishuo struct {
	Website string `bson:"website"`
	Timer string `bson:"timerid"`
	Title string `bson:"title"`
	Edit string `bson:"edit"`
	VedioUrl string `bson:"vediourl"`
	ImgShow string `bson:"imgshow"`
	Contents []string `bson:"contents"`
}
