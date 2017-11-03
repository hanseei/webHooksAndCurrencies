package main 

import (
	//"fmt"
	"net/http"
	//"strings"
	"os"
	"CurrencyTicker"
	"CMD/WebHookAPI"
)


func (db *CurrencyTickerDB) Get_WebHooks(w http.ResponseWriter, db    , id string) {
	webhook, ok := db.Get(id)
	if !ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(webhook)
}


func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	switch r.method {
	
	case "POST":
		var webhook WebHook{}
		err := json.NewDecoder(r.Body).Decode(weebhook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		webhook := insert_webhook() // ...

		return 

	/case "GET":
		http.Header.Add(w.Header(), "content-type", "application/json")

		return
	}
}

func main () {

 mongoDB.CurrencyData.Global_db = &mongoDB.CurrencyData.CurrencyTicker_db{
		"mongodb://localhost",
		"currencyTicker_db",
		"CurrencyData",
	}


//port := os.Getenv("PORT")

port := "localhost:8080"


//http.ListenAndServe(":"+port, nil)
http.ListenAndServe("localhost8080", nil)


}
