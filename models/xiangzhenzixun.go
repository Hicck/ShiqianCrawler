package models

type Xiangzhenzixun struct {
	Website string `bson:"website"`
	Title string `bson:"title"`
	Xiangzhen string `bson:"xiangzhen"`
	Timer string `bson:"timerid"`
	ImgShow string `bson:"imgshow"`
	Contents []string `bson:"contents"`
	Edit string `bson:"edit"`
}
