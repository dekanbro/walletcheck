package walletcheck

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"

    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)

func init() {
     http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

     //Read the Request Parameter "command"
     command := r.FormValue("command")

     //Ideally do other checks for tokens/username/etc

     if command == "/ethminer" {

         ctx := appengine.NewContext(r)
         client := urlfetch.Client(ctx)

         url := "https://ethermine.org/api/miner_new/a5907c5d277c23ca615a14e83152695a21a026bf"
         res, err := client.Get(url)
         if err != nil {
           log.Fatalln(err)
         }

         body, err := ioutil.ReadAll(res.Body)
         if err != nil {
           log.Fatalln(err)
         }

         // fmt.Printf("%s\n", string(body))

         js, err := simplejson.NewJson(body)
         if err != nil {
           log.Fatalln(err)
         }

         address := js.Get("address").MustString()
         if err != nil {
           log.Fatalln(err)
         }

         hashRate := js.Get("hashRate").MustString()
         if err != nil {
           log.Fatalln(err)
         }

         unpaid := js.Get("unpaid").MustInt()
         if err != nil {
           log.Fatalln(err)
         }

         w.Header().Set("Content-Type", "application/json")





         s := []string{"addr:", address, "\\n Hash Rate:", hashRate, "\\n unpaid:", strconv.Itoa(unpaid)}
         u := []string{`{ "response_type": "in_channel",`,  `"text": "`+strings.Join(s, " ")+`"}`}


         fmt.Fprint(w, strings.Join(u, " "))
     } else {
         fmt.Fprint(w,"I do not understand your command.")
     }
}
