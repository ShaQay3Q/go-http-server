package main

import (
	"log"
	"net"
	"net/http"
)

func main(){
	s, _ := net.Listen("tcp", "127.0.0.1:3000")
	c, _ := s.Accept()
	c.Write([]byte(`{"name": "leandro", "age": 17}`))
	log.Println("running!")
	return
	
	

	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json") // MIME type!
		w.Write([]byte(`{"name": "leandro", "age": 17}`))
	}) 
	log.Println("running!")
	http.ListenAndServe("127.0.0.1:3000", nil)
}