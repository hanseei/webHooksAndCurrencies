package WebHookFunctions 
	
	import (
	"net/http"
	"fmt"
	"net/url"
//	"io/ioutil"
//	"strings"
	"strconv"
	"Cloud/CurrencyTicker"
//	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"


	)




type WebHook struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
	WebhookURL string `json:"url"`
	Base string `json:"base"`
	Target string `json:"target"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`

}


func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 2, 64)
}
 



func Insert_Webhook(db *CurrencyTicker.CurrencyTickerDB){

discordURL := "https://discordapp.com/api/webhooks/374908902032146432/lHx9nUAyDy1jmoahR8WrWnWl_y0B1WYc61LRuMPrcS9H5g9CcoVV3KVq7DzpfASIPPzP"


session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
	fmt.Println("error creating session:",err.Error())
}
defer session.Close()

fmt.Println("Enter minTriggerValue:  ")
var min float64
_, err = fmt.Scanln( &min)
if err != nil {
	fmt.Println("error making minTriggerValue")
}
var max float64
fmt.Println("Enter maxTriggerValue:  ")
_, err = fmt.Scanln( &max)
if err !=nil {
	fmt.Println("error making maxTriggerValue")
}
var base string 
for base != "EURO" {
fmt.Println("Enter base currency:	 ")
_, err = fmt.Scanln(&base)
if base != "EURO" {
	fmt.Println("only euro is implemented")
}
}
if err != nil {
	fmt.Println("error making BaseCurrency")
}
fmt.Println("Enter TargetCurrency:  ")
var target string
_, err = fmt.Scanln(&target)
if err !=nil {
	fmt.Println("error making TargetCurrency")

}
in := WebHook{bson.NewObjectId(),discordURL,base,target,min,max}

err = session.DB(db.DatabaseName).C("webhooks").Insert(in)
	
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())

	}

}


func Invoke_Webhooks (db *CurrencyTicker.CurrencyTickerDB){
datas := make([]WebHook, 0, 10)
session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
	fmt.Println("error creating session:",err.Error())
}
	defer session.Close()

dbSize, err := session.DB(db.DatabaseName).C("webhooks").Count()
if err != nil {
     fmt.Println("error counting collection :( :", err.Error())
}



err = session.DB(db.DatabaseName).C("webhooks").Find(nil).All(&datas)
if err != nil {
    fmt.Println("error getting webhook:", err.Error())
}

	for i := 0; i < dbSize; i++ {
		_, err := http.PostForm(datas[i].WebhookURL, url.Values{"content": {"Webhook ID: " + datas[i].ID.Hex() + "		BaseCurrency  " + datas[i].Base + "    TargetCurrency  " + datas[i].Target + "   minTriggerValue	" + FloatToString(datas[i].Min) + "    maxTriggerValue		" + FloatToString(datas[i].Max) } , "username": {"IAMBOT"}})
		if err != nil {
			fmt.Println("Error when posting all webhooks at %i", i, err.Error())
		}
	}
}



// get from database
func Get_Last_Webhook(db *CurrencyTicker.CurrencyTickerDB) {
var myData WebHook 
session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
	fmt.Println("error creating session:",err.Error())
}
	defer session.Close()


dbSize, err := session.DB(db.DatabaseName).C("webhooks").Count()
if err != nil {
     fmt.Println("error counting collection :( :", err.Error())
	}

		err = session.DB(db.DatabaseName).C("webhooks").Find(nil).Skip(dbSize-1).One(&myData)
		if err != nil {
		    fmt.Println("error getting webhook:", err.Error())
		}

	// post to database
	res, err := http.PostForm(myData.WebhookURL, url.Values{"content": {"Webhook ID: " + myData.ID.Hex() + "		BaseCurrency  " + myData.Base + "    TargetCurrency  " + myData.Target + "   minTriggerValue	" + FloatToString(myData.Min) + "    maxTriggerValue		" + FloatToString(myData.Max) } , "username": {"IAMBOT"}})
		if err != nil {
			fmt.Errorf("Error doing post: %v", err.Error())
		}

	if res.StatusCode != http.StatusOK {
		fmt.Errorf("Wrong status code: %v", res.StatusCode)
	}

}


// remove from database
func Remove_Webhook_byId(db *CurrencyTicker.CurrencyTickerDB, id string) {

session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
	fmt.Println("error creating session:",err.Error())
}
defer session.Close()
var myData WebHook 


err = session.DB(db.DatabaseName).C("webhooks").Remove(bson.M{"_id":myData.ID})
count, _ := session.DB(db.DatabaseName).C("webhooks").Count()
fmt.Println("size of webhook: %i", count)
}


