// chat_client.go
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    reader := bufio.NewReader(conn)

    // Show "Enter your name:" prompt from server
    serverPrompt, _ := reader.ReadString(':') // read until colon
    fmt.Print(serverPrompt)

    // Let user type name and send to server
    stdinReader := bufio.NewReader(os.Stdin)
    name, _ := stdinReader.ReadString('\n')
    fmt.Fprint(conn, name)

    // Now start chat: one goroutine to read from server
    go func() {
        for {
            msg, err := reader.ReadString('\n')
            if err != nil {
                break
            }
            fmt.Print(msg)
        }
    }()

    // Main goroutine: read user input and send to server
    for {
        text, err := stdinReader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Fprint(conn, text)
    }
}
