package main

import (
	"io"
	"log"
	"net"
)

// echo - функция-обработчик, просто отражающая полученные данные
func echo(conn net.Conn) {
	defer conn.Close()

	// Создаем буфер для хранения полученных данных
	buf := make([]byte, 512)
	for {
		// Получаем данные через conn.Read в буфер
		size, err := conn.Read(buf[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(buf))

		// Отправляем данные через conn.Write
		log.Println("Writing data")
		if _, err := conn.Write(buf[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
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

// telnet localhost 20080
