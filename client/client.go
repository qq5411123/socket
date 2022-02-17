package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// 1. 创建建立连接
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	fmt.Println("与tcp://127.0.0.1:3333建立连接")
	defer conn.Close()

	//粘包处理（每段增加2字节的长度存储）
	for i := 0; i < 20; i++ {

		msg := "shineyork666"   // => byte[]
		msgLen := len(msg)      // int => int16
		length := int16(msgLen) // 16 位 int16=》占 2 个字节
		fmt.Println("msgLen : ", msgLen)
		fmt.Println("length : ", length)
		pkg := new(bytes.Buffer) // 是因为
		// 把 length => 二进制转化 =》 放到 pkg
		// io.Writer
		binary.Write(pkg, binary.BigEndian, length)
		// pkg.Bytes() => 进制转化后的具体数据 ，append 方法追加msg数据前
		data := append(pkg.Bytes(), []byte(msg)...) // append()
		fmt.Println("data : ", data)
		conn.Write(data)

	}

}
