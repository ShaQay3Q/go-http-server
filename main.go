package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main(){
	//! without handling errors
	// set up the server
	server, err := net.Listen("tcp", "127.0.0.1:3001")
	log.Println("before accepting")

	if err != nil {
		log.Fatal("Error starting the server", err)
	}
	
	// set up connection to server
	connection, _ := server.Accept()

	// recieve the request
	var b [1024]byte
	numbOfBytes, _ :=  connection.Read(b[:])
	// parse request into string
	req := string(b[0:numbOfBytes]) // sub-slice
	log.Println(req)
	jsonRes := `{"name": "Maria", "age": 17}`
	resLength := len(jsonRes)
	// resLeng := len(jsonRes)

	if strings.HasPrefix(req, "GET /contact HTTP/1.1\r\n"){
		response := "HTTP/1.1 200 OK\r\n                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 gi" +
			"Content-Type: application/json\r\n" +
			"Content-Length: " + strconv.Itoa(resLength) + "\r\n" +
			"\r\n" + jsonRes

		connection.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\r\n" +
			"Content-Length: 0\r\n" +
			"\r\n"
		connection.Write([]byte(response))
	}

	log.Println("after accepting!")
	// c.Header()
	return
	
	

	// http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "application/json") // MIME type!
	// 	w.Write([]byte(`{"name": "leandro", "age": 17}`))
	// }) 
	log.Println("running!")
	http.ListenAndServe("127.0.0.1:3001", nil)
}