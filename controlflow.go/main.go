package main

import (
	"fmt"
	"time"
)
type Server struct {
	quitch chan struct{} // 0 bytes or else in others is some bytes 
	msgch chan string 
}

//whenever we make a server struct , always need to make a constructor for that server

func newServer () *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch: make(chan string, 128), // buffered channel
}
}
func (s *Server) start (){
	fmt.Println("server starting")
	s.loop() // block
}
func (s *Server) loop (){
	for {
		select {
		case <- s.quitch:
			//do some stuff when need to quit
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
}
func (s *Server) sendMessage(msg string){
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string ){
	fmt.Println("we received message", msg)
}
func main(){
	server := newServer()
	// we need to add go here as if we dont then the server will keep running , so we made it asynchronous so that as soon as the main program stops the server also stops 
	go server.start()
	server.sendMessage("Heya there ")
	// server.start()
	time.Sleep(time.Second * 2)
}