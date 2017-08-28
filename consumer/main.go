package main

import (
	"github.com/swarndeepkumar/rdkafka-go/consumer"
	"fmt"
)

func main() {
	var broker string = "localhost:9092"
	var service string = "messeging1"
	var topics = []string {"topiccode1","topiccode2","topiccode3"}
        var callback = func(err error , data string){
		fmt.Println("error messege from consumer", err);
		fmt.Println("data message from consumer", data);
	}	
          consumer.Consumemessage(broker, service, topics, callback);
}
