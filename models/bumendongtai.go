package models

type Bumendongtai struct {
	Website string `bson:"website"`
	Bumen string `bson:"bumen"`
	TimerId string `bson:"timerid"`
	Title string `bson:"title"`
	Edit string `bson:"edit"`
	ImgShow string `bson:"imgshow"`
	Contents []string `bson:"contents"`
}
