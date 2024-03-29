package mongo

import (
	"github.com/wudian/market_data/models"
	"gopkg.in/mgo.v2"
)

var (
	MgoUrl = "127.0.0.1:27017"
	DbName = "wx"
	TableName = "ticker"
)

type MgoClient struct {
	session *mgo.Session
	c *mgo.Collection
}


func NewMgoClient() (client MgoClient, err error){
	client.session, err = mgo.Dial(MgoUrl)
	if err == nil{
		// Optional. Switch the session to a monotonic behavior.
		client.session.SetMode(mgo.Monotonic, true)
		client.c = client.session.DB(DbName).C(TableName)
	}
	return
}

func (client *MgoClient)Denit()  {
	client.session.Close()
}

func (client *MgoClient)Insert(ticker *models.Ticker) {
	client.c.Insert(ticker)

}


//
//func Test()  {
//
//
//
//
//	//err = c.Insert(models.Ticker{Symbol:"ETH-USDT"})
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//result := models.Ticker{}
//	//err = c.Find(bson.M{"symbol": "ETH-USDT"}).One(&result)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//fmt.Println(result)
//
//
//}