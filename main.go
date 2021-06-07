package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"os"
	"strings"
)
type Req struct {
	reqUrl,contentType string
}
func main() {
	basePath := "/root/go/src/go_static_server"
	address := ":9000"
	listen, _ := net.Listen("tcp",address )
	log.Printf("listen=%#v\n", address)
	for  {
		conn, _ := listen.Accept()
		go func() {
			defer conn.Close()
			log.Printf("NEW CLIENT:%s\n", conn.RemoteAddr().String())
			//get request url
			req := getReq(conn)
			path := basePath+req.reqUrl
			file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
			if err!=nil {
				return
			}
			//write back header
			writeHeader(conn,req.contentType)
			//write back file
			writeFile(conn,file)
		}()
	}
}
func getReq(conn net.Conn) *Req {
	reader := bufio.NewReader(conn)
	line, _, _ := reader.ReadLine()
	splits := strings.Split(string(line), " ")
	reqUrl := splits[1]
	paths := strings.Split(reqUrl, "/")
	contentType := paths[len(paths)-1]
	return &Req{reqUrl: reqUrl,contentType: contentType}
}
func writeHeader(conn net.Conn,contentType string)  {
	buffer := bytes.Buffer{}
	buffer.WriteString("HTTP/1.1 200 OK\r\n")
	buffer.WriteString("Content-Type:")
	buffer.WriteString(contentType)
	buffer.WriteString(";charset=UTF-8\r\n\r\n")
	conn.Write(buffer.Bytes())
}
func writeFile(conn net.Conn,file *os.File)  {
	buf := make([]byte,1024)
	for  {
		readN, _ := file.Read(buf)
		if readN <= 0  {
			break
		}
		conn.Write(buf)
	}
}