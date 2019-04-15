package config

import (
	"encoding/xml"
	"fmt"
	"github.com/wudian/market_data/mongo"
	"io/ioutil"
	"os"
)

type SConfig struct {
	XMLName    xml.Name `xml:"config"`  // 指定最外层的标签为config

	//Sender string `xml:"sender"`

	ApiNames SApiNames `xml:"ApiNames"`
	VecSymbols SVecSymbols  `xml:"VecSymbols"`
	Mongo SMongo `xml:"Mongo"`
}

type SApiNames struct {
	Api []SApi `xml:"api"`
}

type SApi struct {
	Name string `xml:"name"`
	Use string `xml:"use"`
	Url string `xml:"url"`
}

type SVecSymbols struct {
	Symbol []string `xml:"symbol"`
}
type SMongo struct {
	MgoUrl string `xml:"mgoUrl"`
	DbName string `xml:"dbName"`
	TableName string `xml:"tableName"`
}
func readXml() (*SConfig, error) {
	file, err := os.Open("market_data.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	v := SConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}


	//fmt.Println(v.VecSymbols, v.ApiNames)

	return &v, nil
}

func setGlobal(v *SConfig)  {
	for _, api := range v.ApiNames.Api{
		if api.Use == "true"{
			global.ApiNames[api.Name] = api.Url
		}
	}
	for _,symbol := range v.VecSymbols.Symbol{
		global.VecSymbols = append(global.VecSymbols, symbol)
	}

	m := v.Mongo
	mongo.MgoUrl = m.MgoUrl
	mongo.DbName = m.DbName
	mongo.TableName = m.TableName
}