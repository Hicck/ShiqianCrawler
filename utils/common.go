package utils

import (
	"ShiqianCrawler/models"
	"errors"
	"github.com/axgle/mahonia"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"
	"strings"
)

func CheckWebsite(collectionName string,website string)(reslut bool){
	info := models.Shiqianshanshui{}

	session := Session.Copy()
	defer session.Close()

	collection := session.DB("ShiqianNews").C(collectionName)
	collection.Find(bson.M{"website":website}).One(&info)

	if info.Website == "" || len(info.Website) <= 0{
		reslut = false
	}else {
		reslut = true
	}
	return
}

func InsertDetailUrl(collectionName , url string,infotype interface{})(err error){

	if CheckWebsite(collectionName,url) == true {
		return
	}
	var info interface{}
	switch infotype.(type) {
		case models.Shiqianshanshui:
			info = models.Shiqianshanshui{
				Website:url,
			}
		case models.Snack:
			info = models.Snack{
				Website:url,
			}
	case models.Minsu:
		info = models.Minsu{
			Website:url,
		}
	case models.Taitea:
		info = models.Taitea{
			Website:url,
		}
	case models.Wenquan:
		info = models.Wenquan{
			Website:url,
		}
	case models.Yelang:
		info = models.Yelang{
			Website:url,
		}
	case models.RedCulture:
		info = models.RedCulture{
			Website:url,
		}
	case models.Yaowen:
		info = models.Yaowen{
			Website:url,
		}
		default:
			log.Printf("[common] InsertDetailUrl : 传入的错误的参数类型")
	}

	session := Session.Copy()
	defer session.Close()
	collection := session.DB("ShiqianNews").C(collectionName)
	err = collection.Insert(info)
	if err != nil {
		errStr := "[common] InsertDetailUrl : 出现错误, 数据库集合 : " + collectionName + "\n"
		err = errors.New(errStr)
		return
	}
	return
}


func RemoveDataByWebsite(collectionName string,website string){
	session := Session.Copy()
	defer session.Close()
	collection := session.DB("ShiqianNews").C(collectionName)
	collection.Remove(bson.M{"website":website})
}


func GetWebsite(rUrl *url.URL)(website string){
	website = rUrl.Scheme + "://" + rUrl.Host+rUrl.RequestURI()
	return
}

func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func ReplaceString(source string)(destation string){
	destation = strings.Replace(source,"“","",-1)
	destation = strings.Replace(destation,"”","",-1)
	destation = strings.Replace(destation,"·",".",-1)
	destation = strings.Replace(destation,"、",".",-1)
	destation = strings.Replace(destation,"，",".",-1)
	destation = strings.Replace(destation,"，",".",-1)
	destation = strings.Replace(destation,"……","",-1)
	destation = strings.Replace(destation,"—","",-1)
	destation = strings.Replace(destation,"》","",-1)

	destation = convertToString(destation, "gbk", "utf-8")
	destation = strings.Replace(destation,"聽","",-1)

	return
}




