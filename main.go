package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/swarndeepkumar/messaginglib-go/messaging"
	"github.com/gorilla/mux"
	"sync"
	"os"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Con_Bootstrap  string
	Con_Groupid  string
	Con_SessionTimeoutMS string
	Pro_Bootstrap string
}

var wg = sync.WaitGroup{}
var configuration = Configuration{}
/// get  configuration from file 
var err = gonfig.GetConf("config/config.dev.json", &configuration)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
        if err != nil {
		fmt.Println("configuration file could not read");
                os.Exit(500)
        }
	// consumer configuration
	var consConfig = map[string]string{"bootstrap.servers":configuration.Con_Bootstrap,"group.id":configuration.Con_Groupid,"session.timeout.ms":configuration.Con_SessionTimeoutMS}
	router.HandleFunc("/topic/{tname}/msg/{pmessage}", sendMessage)
	//router.HandleFunc("/topic/{tname}", consumeTopic)


	consumeTopic(consConfig)
	log.Fatal(http.ListenAndServe(":8080", router))
	// need for goroutne
	wg.Wait()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome messaging demo1 application !\n")
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pmessage := vars["pmessage"]
	tname := vars["tname"]
	// prodcuer configuration
	var prodConfig = map[string]string{"bootstrap.servers":configuration.Pro_Bootstrap}
	// callback method, that will be called after completing the method 
        var callback = func(err error, data string){
		if err != nil { 
			// write your success code, bec message is produced
			fmt.Println("sendMessage: error :", err);
		}else{ 
 			// write your falure code, bec message could not produced
			fmt.Println("sendMessage: success");
		}
          }
          messaging.Producemessage(prodConfig, tname, pmessage, callback);
	 fmt.Fprint(w, "Welcome messaging demo1 application !\n")
	 fmt.Fprint(w, "\n")
	 fmt.Fprint(w, "\n")
	 fmt.Fprint(w, "Message send successfull.. !\n")
	 fmt.Fprint(w, "\n")
         fmt.Fprint(w, "Topic name is- ",tname + "\n")
	 fmt.Fprint(w, "Message was  - ",pmessage + "\n")

}

func consumeTopic(consConfig map[string]string){
	var topics = []string {"topiccode1","topiccode2"}
	// callback method, that will be called after completing the method
        var callbackconsumer = func(err error , data []string){
		 if err != nil {
                        // write your success code, bec message is produced
                        fmt.Println("consumeTopic: error :", err);
                }else{
                        // write your falure code, bec message could not produced
                        fmt.Println("consumeTopic demo 1 application: success: recieved message:",data);
                }

        }
         messaging.Consumemessage(consConfig, topics, callbackconsumer)

}

