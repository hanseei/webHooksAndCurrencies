
package CurrencyTicker

import (
	"fmt"
	"net/http"
	//"net/url"
	"io/ioutil"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	"time"
//	"github.com/heroku/Assignment2/WebHooks"
	"encoding/json"
)


//var Global_db CurrencyData





type CurrencyTickerDB struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
}

type CurrencyData struct {
	/*Id        		bson.ObjectId `bson:"_id,omitempty"`
	Date 			string	`json:	"date"`
	CurrencyRate 	float64 	`json:	"currencyRate"`
	*/

base string	`json:  "base"`
date string	`"json: "date"`

rates struct {
AUD float64 `json:	"AUD"`
BGN	float64 `json:	"BGN"`
BRL	float64 `json:	"BRL"`
CAD	float64 `json:	"CAD"`
CHF	float64 `json:	"CHF"`
CNY	float64 `json:	"CNY"`
CZK	float64 `json:	"CZK"`
DKK	float64 `json:	"DKK"`
GBP	float64 `json:	"GBP"`
HKD	float64	`json:	"HKD"`
HRK	float64	`json:	"HRK"`
HUF	float64	`json:	"HUF"`
IDR	float64	`json:	"IDR"`
ILS	float64	`json:	"ILS"`
INR	float64	`json:	"INR"`
JPY	float64	`json:	"JPY"`
KRW	float64	`json:	"KRW"`
MXN	float64	`json:	"MXN"`
MYR	float64	`json:	"MYR"`
NOK	float64	`json:	"NOK"`
NZD	float64	`json:	"NZD"`
PHP	float64	`json:	"PHP"`
PLN	float64	`json:	"PLN"`
RON	float64	`json:	"RON"`
RUB	float64	`json:	"RUB"`
SEK	float64	`json:	"SEK"`
SGD	float64	`json:	"SGD"`
THB	float64	`json:	"THB"`
TRY	float64	`json:	"TRY"`
USD	float64	`json:	"USD"`
ZAR	float64	`json:	"ZAR"`

} `json:	"rates"`



} 

type UserData struct {
	MinValue		float64	`json: "minValue"`
	MaxValue		float64	`json: "maxValue"`
	BaseCurrency	string	`json: "baseCurrency"` 
	TargetCurrency	string	`json:	"targetCurrency"`

}



func (db *CurrencyTickerDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"Id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.CollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

/*
func (db *CurrencyTickerDB) Add_UserData(min int, max int, base string, target string) error {

session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
		panic(err)
}
defer session.Close()
	

	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(w)

	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	return nil




}



*/


func (db *CurrencyTickerDB) AddNewTicker() error {

session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
		panic(err)
}
defer session.Close()


fixerURL := "http://api.fixer.io/latest" 

	resp, _ := http.Get(fixerURL)
	bytes, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()										// make sure we close body after

	var info map[string]interface{}
	err = json.Unmarshal(bytes, &info)
	if err != nil{
		fmt.Printf("error in unmarshalling: %v", err.Error())
	}
//	ex := CurrencyData{}
//	bytes.Date := date 
//	bytes.CurrencyRate := currRate



	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(info)
	
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}
	
	
return nil

/*	url := "http://api.fixer.io/latest?BASE=EUR"

	res, _ := http.Get(url)
	data, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()	
*/
	//return data

}










func (db *CurrencyTickerDB) Get_Latest()  CurrencyData{
	


	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	dbSize := 3
	var items CurrencyData
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(nil).Skip(dbSize-1).One(items)
	if err != nil {
		panic(err)
	}

	return items



}



/*
func (db *CurrencyTickerDB) Get_Average(base string, target string){
session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
		fmt.Println("Error in stuff:", err.Error())
}
defer session.Close()
	count, err := session.DB(db.DatabaseName).C(db.CollectionName).Count()
	if err != nil	{
		fmt.Println("Error counting stuff:", err.Error())
	}


	result := make(map[string]interface{})
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(nil).Skip(count-3).All(&result)
	/*if err != nil {
		fmt.Println("Error setting up data stuff:", err.Error())
	}
		data := make([]float64, len(result))
		for i := 0; i < 3; i++ {
			data[i] = result[target].(float64)
		}
/*
	x := data[0]
	y := data[1]
	z := data[2]

	fmt.Println("data %i", data[0])
/*
	avg := x + y + z
	avg = avg / 3
	fmt.println("average of last 3 days: %i", avg)



}


*/

/*
func (db *CurrencyTickerDB) Get_Average() CurrencyData {

}
*/






/*
func main () {
//	fmt.Printf("Latest update:  %v", string(data))
	
//	text := "Heroku timer test at:  "+ time.Now().string
	delay := time.Minute*15

	time.Sleep(delay)
}

*/