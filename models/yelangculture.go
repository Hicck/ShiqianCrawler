package models

type Yelang struct {
	Website string `bson:"website"`
	Title string `bson:"title"`
	Timer string `bson:"timerid"`
	ImgShow string `bson:"imgshow"`
	Contents []string `bson:"contents"`
}
