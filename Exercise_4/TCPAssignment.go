package main

import (
	"fmt"
	"net"
	//"time"
)

func main(){
	con := connectTCP()
	message := "This is a message!"
	defer con.Close()
	receiveMessage(con)
	sendMessage(con, message)
	listener := localServerConnecton()
	defer listener.Close()
	serverConnection, err := listener.Accept()
	if err != nil{
		fmt.Println("could not connect to server: ", err)
		return
	}
	defer serverConnection.Close()
	receiveMessage(serverConnection)
}


func connectTCP() net.Conn{
	serverAddress := net.TCPAddr{
		Port: 34933,
		IP: net.ParseIP("10.100.23.204"),
	}
	connection, err := net.Dial("tcp", serverAddress.String())
	if err != nil{
		fmt.Println("Unable to connect to server: ", err)
		return nil
	}
	fmt.Println("Connection was good")
	return connection
}


func receiveMessage(con net.Conn){ //rename serverStatus to something that make sence
	buffer := make([]byte, 1024)
	numberOfBytes, err := con.Read(buffer)
	if err != nil{
		fmt.Println("Unable to read from server: ", err)
		return
	}
	fmt.Println("Message from server: ", string(buffer[:numberOfBytes]))
}

func sendMessage(con net.Conn, message string){
	_, err := con.Write([]byte (message))
	if err != nil{
		fmt.Println("Unable to write to server: ", err)
		return
	}
	fmt.Println("Message sendt")
}

func localServerConnecton() net.Listener {
	localAddress := net.TCPAddr{
		Port: 20002,
		IP: net.ParseIP("10.100.23.12"),
	}
	listener, err := net.Listen("tcp", localAddress.String())
	if err != nil{
		fmt.Println("Unable to connect to local server: ", err)
		return nil
	}
	fmt.Println("Connection succesfull bitch!")
	return listener
}