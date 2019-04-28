// Jackson Frankfurt
// golang client program
// Idea: 
  // Client takes command line input of own name and name of friend
  // Client sends to server own name and name of friend
  // Server connects two clients if they ask for each other // Would be more simple if I just tested that they have the same 'secret code'
  // I'll need to list the pairs of users who are connected
  // I'll need to list the users who are pending connection
  // I'll need to search the pending users whenever I add a new pending user.
  // Server forwards messages between clients
  // Client loops:
    // prompt to user for a message
    // send that message to server

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