package main

import (
	"fmt"
)

func channeldemo() {
	ch := make(chan string) //the channel will accept string messages

	//anonymous function. Feeding message "hello" to the channel. If you dont add go
	//it'll throw an error, because you are trying to send it but there is no receiver.
	//adding go makes it non blocking, so by the time the message is being pushed in the channel
	//recv is ready to receive it.
	go func(msg string) {
		ch <- msg
	}("hello")

	recv := <-ch

	fmt.Println(recv)
}
