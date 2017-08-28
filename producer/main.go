package main

import (
	"fmt"
	"github.com/swarndeepkumar/rdkafka-go/producer"
)
//var broker string = "localhost:9092"
//var topic string = "topic3"
func main() {
	 var broker string = "localhost:9092"
	 var topic  string = "topiccode2"
	 var message string = "my message swarn2"
         var callback = func(err error, data string){
               // if err
		// then execute something
		// else somthing else
		fmt.Println("error messsage", err)
		fmt.Println("response message", data)
	  }
          producer.Producemessage(broker, topic, message, callback);
}
