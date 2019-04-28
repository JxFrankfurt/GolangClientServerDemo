// Jackson Frankfurt
// golang client program

package main
//this is package level
import (
    "fmt"
    "net"
    "os"
    "bufio"
)

//what is const? a type?
const ( 
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Dial server
    connection, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error dialing:", err.Error())
        os.Exit(1) //error exit code
    }
    stdReader := bufio.NewReader(os.Stdin)
    connectionReader := bufio.NewReader(connection)    
    fmt.Print("What is your name?\n") // Prompt the user:
    userName,_ := stdReader.ReadString('\n')
    fmt.Print("With whom would you like to connect?\n")
    friendName,_ := stdReader.ReadString('\n')
    
    fmt.Fprintf(connection, userName + 'calls' + friendName) // "GET / HTTP/1.0\r\n\r\n"
    
    // respone from the server. '\n' is the last character to recieve.
    status, err := connectionReader.ReadString('\n') 
    fmt.Print(status) // the server's message
    connection.Close()
}
