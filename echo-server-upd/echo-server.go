package main

import (
	"bufio"
	"log"
	"net"
)

// echo - функция-обработчик, просто отражающая полученные данные
func echo(conn net.Conn) {
	defer conn.Close()

	// Получаем данные через reader
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s\n", len(s), s)

	// Отправляем данные через writer
	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func main() {
	// Привязываемся к TCP-порту 20080 во всех интерфейсах
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
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

// telnet localhost 20080
