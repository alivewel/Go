package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		// Получаем данные через reader
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data:", err)
			return
		}
		log.Printf("Read %d bytes: %s\n", len(s), s)

		// Отправляем данные через writer
		log.Println("Writing data")
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data:", err)
			return
		}
		writer.Flush()
	}
}

func main() {
	// Привязываемся к TCP-порту 20080 во всех интерфейсах
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	defer listener.Close()
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Ожидаем соединения и при его установке создаем net.Conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Обрабатываем соединение, используя горутины для многопоточности
		go echo(conn)
	}
}

// Для отпарвки данных использовать команду: telnet localhost 20080
