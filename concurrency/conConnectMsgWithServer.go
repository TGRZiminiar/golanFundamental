package concurrency

import (
	"fmt"
	"time"
)

type Message struct {
	From string;
	Payload string;
}

type Server struct {
	msgch chan Message;
	quitch chan struct{};
}

func (s *Server) StartAndListen(){
	
	//name for loop
	free:
		for {
			select {

				//block here until someone sending a message to the channel
			case msg := <-s.msgch:
				fmt.Printf("Receive Message From:%s\tPayload:%s\n", msg.From, msg.Payload);
			case <-s.quitch:
				fmt.Println("The Server Shut Down");
				break free;
			default:
				
			}
		}

		fmt.Println("Server Shut Down");
}

func sendMessageToServer(msgch chan Message, payload string){
	
	msg := Message{
		From: "Mix",
		Payload: payload,
	}

	msgch <- msg;

	// fmt.Println("Sending Message");
}

func quitServer(quitch chan struct {}){
	close(quitch);
}

func ExampleConCurrency2(){

	s := &Server{
		msgch: make(chan Message),
		quitch: make(chan struct{}),
	};

	go s.StartAndListen();

	for i := 0; i < 5; i++ {
		go func(){
			time.Sleep(time.Millisecond * 200);
			sendMessageToServer(s.msgch, "Hello World");
		}()
	}

	go func(){
		time.Sleep(time.Millisecond * 400);
		quitServer(s.quitch);
	}()


}

