# GoChat
Simple chat server written in Go.
This project creates a server and client(s) for simple terminal-based chatting using WebSockets (golang.org/x/net/websocket).

## Requirements
1. docker

## Build and Run it
To run projects locally all you have to do is run:
```
  1. make build
  2. make run
```

```
  And make operations below (x2, or x3 or x10) ...in different terminal windows. Type a message in one client's terminal and see in sent to every client!

  1. make run-terminal
  2. cd cmd
  3. cd client
  4. go run main.go
  
```