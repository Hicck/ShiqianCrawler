package models

type Yaowen struct {
	Website string `bson:"website"`
	Title string `bson:"title"`
	Timer string `bson:"timerid"`
	ImgShow string `bson:"imgshow"`
	Contents []string `bson:"contents"`
	Edit string `bson:"edit"`
}
