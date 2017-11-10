package main

	
	import (
//	"net/http"
//	"fmt"
//	"net/url"
//	"io/ioutil"
//	"strings"
//	"strconv"
	"Cloud/CurrencyTicker"
//	"encoding/json"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"Cloud/WebHookFunctions"

	)






func main () {


db := CurrencyTicker.CurrencyTickerDB{"mongodb://Siggy:Siggy@ds145275.mlab.com:45275/currency_db", "currency_db", "currency"}

//userWebhook := PayLoad{}

//fixerURL := "http://api.fixer.io/latest?symbols=NOK"

//	resp, _ := http.Get(fixerURL)
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	webData := string(bytes)
	
//	defer resp.Body.Close()										// make sure we close body after
	
//	parts := strings.Split(webData,"\"")	
//	base := parts[3]
//	TargetCurrency := parts[11]
//	minTriggerValue := 1.5
//	maxTriggerValue := 2.55

//discordURL := "https://discordapp.com/api/webhooks/374908902032146432/lHx9nUAyDy1jmoahR8WrWnWl_y0B1WYc61LRuMPrcS9H5g9CcoVV3KVq7DzpfASIPPzP" // TODO edit this 



/*
session, err := mgo.Dial(db.DatabaseURL)
if err != nil{
	fmt.Println("error creating sesh in userWebhookAPI.go:", err.Error())
}

//bytes, _ := ioutil.ReadAll(res.Body)
//	var info (map[string]interface{})[discordURL, base, target, min, max]
		
//	err = json.Unmarshal(bytes, &info)
in := WebHook{bson.NewObjectId(),discordURL,base,target,min,max}
/*
var inInterface map[string]interface{}
inrec, _ := json.Marshal(in)
json.Unmarshal(inrec, &inInterface)
//fmt.Println(v.WebhookURL, v.base, v.target, v.min, v.max)
*//*
err = session.DB(db.DatabaseName).C("webhooks").Insert(in)
	
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())

	}
*/
db.AddNewTicker()
WebHookFunctions.Insert_Webhook(&db)
WebHookFunctions.Invoke_Webhooks(&db)
WebHookFunctions.Get_Last_Webhook(&db)
//WebHookFunctions.Remove_Webhook_byId(&db)



//GET_LATEST()

/*body ,err, := ioutil.ReadAll(res.Body)
if err != nil {
	fmt.Printf("Id:   ", res.Id) 
}*/

// delete from database


//db.Get_Average(myData.Base, myData.Target)







// invoke all webhooks

// result := models.User{}

// result := make([]models.User, 0, 10) // Here you can specify a len and a cap.


//

// post to user


}










/*

func TriggerAllWebhooks() {

/*count,err = session.DB(db.DatabaseName).C("webhooks").Count()

for(int i; i<count; i++) {

vardata = session.DB(db.DatabaseName).C("webhooks").findOne(i)
*/


//
/*

var myData WebHook 
dbSize, err := session.DB(db.DatabaseName).C("webhooks").Count()
if err != nil {
     fmt.Println("error counting collection :( :", err.Error())
}

err = session.DB(db.DatabaseName).C("webhooks").Find(nil).All(&myData)
if err != nil {
    fmt.Println("error getting webhook:", err.Error())
}



//


res, err := http.PostForm(myData[i].WebhookURL, url.Values{"content": {"Webhook ID: " + myData[i].ID.Hex() + "		BaseCurrency  " + myData[i].Base + "    TargetCurrency  " + myData[i].Target + "   minTriggerValue	" + FloatToString(myData[i].Min) + "    maxTriggerValue		" + FloatToString(myData[i].Max) } , "username": {"IAMBOT"}})
if err != nil {
// post to user
}
}
}
*/
