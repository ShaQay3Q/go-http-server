package main

import (
	"log"
	"net/http"
)

func main(){
	// s, _ := net.Listen("tcp", "127.0.0.1:3000")
	// log.Println("before accepting")
	// c, _ := s.Accept()
	// var b [1024]byte
	// c.Read(b [:])
	// log.Println(b)
	// c.Write([]byte(`{"name": "leandro", "age": 17}`))
	// log.Println("after accepting!")
	// return
	
	

	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json") // MIME type!
		w.Write([]byte(`{"name": "leandro", "age": 17}`))
	}) 
	log.Println("running!")
	http.ListenAndServe("127.0.0.1:3000", nil)
}